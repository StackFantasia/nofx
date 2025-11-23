package trader

import (
	"log"
	"math"
	"nofx/decision"
	"nofx/logger"
	"strings"
	"time"
)

// verifyAndUpdateActualFillPrice 验证并更新实际成交价格，确保风险不超过 2%
// 在开仓后立即调用，基于实际成交价格验证风险
func (at *AutoTrader) verifyAndUpdateActualFillPrice(
	decision *decision.Decision,
	actionRecord *logger.DecisionAction,
	side string, // "long" or "short"
	estimatedPrice float64, // 开仓前的预估价格
	openTime int64, // 开仓时间（毫秒时间戳）
) error {
	const maxRetries = 3
	const retryDelay = 500 * time.Millisecond
	const maxRiskPercent = 2.0 // 最大风险 2%

	log.Printf("  🔍 验证实际成交价格和风险...")

	// 定义查询时间范围：开仓前后各 10 秒
	startTime := openTime - 10000
	endTime := openTime + 10000

	var fills []map[string]any
	var err error

	// 重试机制：交易所可能需要时间同步成交记录
	for i := 0; i < maxRetries; i++ {
		if i > 0 {
			log.Printf("  ⏳ 等待 %v 后重试获取成交记录 (尝试 %d/%d)...", retryDelay, i+1, maxRetries)
			time.Sleep(retryDelay)
		}

		fills, err = at.trader.GetRecentFills(decision.Symbol, startTime, endTime)
		if err != nil {
			log.Printf("  ⚠️ 获取成交记录失败 (尝试 %d/%d): %v", i+1, maxRetries, err)
			continue
		}

		// 如果找到成交记录，停止重试
		if len(fills) > 0 {
			break
		}
	}

	if err != nil || len(fills) == 0 {
		log.Printf("  ⚠️ 未能获取实际成交价，使用预估价格 %.2f", estimatedPrice)
		return nil // 不阻断流程，继续执行
	}

	// 过滤匹配的成交记录
	// open_long -> Buy
	// open_short -> Sell
	expectedSide := "Buy"
	if side == "short" {
		expectedSide = "Sell"
	}

	var matchedFills []map[string]any
	for _, fill := range fills {
		fillSide, _ := fill["side"].(string)
		if fillSide == expectedSide {
			matchedFills = append(matchedFills, fill)
		}
	}

	if len(matchedFills) == 0 {
		log.Printf("  ⚠️ 未找到匹配的 %s 成交记录，使用预估价格 %.2f", expectedSide, estimatedPrice)
		return nil
	}

	// 计算加权平均成交价格
	var totalValue float64
	var totalQuantity float64

	for _, fill := range matchedFills {
		price, _ := fill["price"].(float64)
		quantity, _ := fill["quantity"].(float64)

		totalValue += price * quantity
		totalQuantity += quantity
	}

	actualEntryPrice := totalValue / totalQuantity

	// 更新 actionRecord 为实际成交价
	actionRecord.Price = actualEntryPrice

	// 计算实际滑点
	slippage := actualEntryPrice - estimatedPrice
	slippagePct := (slippage / estimatedPrice) * 100

	log.Printf("  📊 成交价格: 预估 %.2f → 实际 %.2f (滑点 %+.2f, %+.2f%%) [共 %d 笔成交]",
		estimatedPrice, actualEntryPrice, slippage, slippagePct, len(matchedFills))

	// 获取账户净值用于风险计算
	balance, err := at.trader.GetBalance()
	if err != nil {
		log.Printf("  ⚠️ 获取账户余额失败: %v", err)
		return nil // 不阻断流程
	}

	totalBalance := 0.0
	if tb, ok := balance["totalBalance"].(float64); ok {
		totalBalance = tb
	} else if tb, ok := balance["balance"].(float64); ok {
		totalBalance = tb
	}

	if totalBalance <= 0 {
		log.Printf("  ⚠️ 无法获取账户净值，跳过风险验证")
		return nil
	}

	// 计算实际风险
	actualRisk := calculatePositionRisk(
		actualEntryPrice,
		decision.StopLoss,
		decision.PositionSizeUSD,
		totalBalance,
		side,
	)

	log.Printf("  💰 风险验证: %.2f%% (仓位 $%.2f, 止损 %.2f, 净值 $%.2f)",
		actualRisk.RiskPercent, decision.PositionSizeUSD, decision.StopLoss, totalBalance)

	// 如果风险超过 2%，采取保护措施
	if actualRisk.RiskPercent > maxRiskPercent {
		log.Printf("  🚨 警告：实际风险 %.2f%% 超过 %.2f%% 限制！", actualRisk.RiskPercent, maxRiskPercent)
		log.Printf("  └─ 价格风险: %.2f%% | 止损金额: $%.2f | 手续费: $%.2f",
			actualRisk.PriceRiskPercent, actualRisk.StopLossUSD, actualRisk.FeeUSD)

		// 选项1：调整止损到更安全的位置（优先）
		adjustedStopLoss := calculateMaxStopLoss(
			actualEntryPrice,
			decision.PositionSizeUSD,
			totalBalance,
			maxRiskPercent,
			side,
		)

		if adjustedStopLoss > 0 {
			log.Printf("  🛡️ 自动调整止损: %.2f → %.2f (确保风险 ≤ %.2f%%)",
				decision.StopLoss, adjustedStopLoss, maxRiskPercent)

			// 取消旧的止损单
			if err := at.trader.CancelStopLossOrders(decision.Symbol); err != nil {
				log.Printf("  ⚠️ 取消旧止损单失败: %v", err)
			}

			// 设置新的止损
			quantity := actionRecord.Quantity
			positionSide := strings.ToUpper(side)
			if err := at.trader.SetStopLoss(decision.Symbol, positionSide, quantity, adjustedStopLoss); err != nil {
				log.Printf("  ❌ 调整止损失败: %v，建议手动平仓！", err)
			} else {
				// 更新内部记录
				posKey := decision.Symbol + "_" + side
				at.positionStopLoss[posKey] = adjustedStopLoss
				log.Printf("  ✓ 止损已调整，风险已控制在 %.2f%% 以内", maxRiskPercent)
			}
		} else {
			// 选项2：无法通过调整止损控制风险，立即平仓
			log.Printf("  ⚠️ 无法通过调整止损控制风险，建议立即平仓")
			log.Printf("  ⚠️ 请在下一个决策周期中给出平仓指令")
			// 注意：这里不直接平仓，而是让AI在下一个周期决策，避免过度干预
		}
	} else {
		log.Printf("  ✓ 风险验证通过: %.2f%% ≤ %.2f%%", actualRisk.RiskPercent, maxRiskPercent)
	}

	return nil
}

// PositionRisk 持仓风险计算结果
type PositionRisk struct {
	PriceRiskPercent float64 // 价格风险百分比
	StopLossUSD      float64 // 止损金额 (USDT)
	FeeUSD           float64 // 手续费 (USDT)
	TotalRiskUSD     float64 // 总风险 (USDT)
	RiskPercent      float64 // 占账户净值的风险百分比
}

// calculatePositionRisk 计算持仓风险
func calculatePositionRisk(
	entryPrice float64,
	stopLoss float64,
	positionSizeUSD float64,
	totalBalance float64,
	side string, // "long" or "short"
) PositionRisk {
	var priceRiskPercent float64

	if side == "short" {
		// 空单：止损价 > 入场价时亏损
		priceRiskPercent = (stopLoss - entryPrice) / entryPrice
	} else {
		// 多单：止损价 < 入场价时亏损
		priceRiskPercent = (entryPrice - stopLoss) / entryPrice
	}

	// 止损金额
	stopLossUSD := positionSizeUSD * math.Abs(priceRiskPercent)

	// 手续费估算（开仓 + 平仓，Taker 费率 0.05%）
	feeUSD := positionSizeUSD * 0.0005 * 2

	// 总风险
	totalRiskUSD := stopLossUSD + feeUSD

	// 风险占比
	riskPercent := (totalRiskUSD / totalBalance) * 100

	return PositionRisk{
		PriceRiskPercent: math.Abs(priceRiskPercent) * 100,
		StopLossUSD:      stopLossUSD,
		FeeUSD:           feeUSD,
		TotalRiskUSD:     totalRiskUSD,
		RiskPercent:      riskPercent,
	}
}

// calculateMaxStopLoss 计算满足最大风险限制的止损价格
func calculateMaxStopLoss(
	entryPrice float64,
	positionSizeUSD float64,
	totalBalance float64,
	maxRiskPercent float64,
	side string, // "long" or "short"
) float64 {
	// 预留手续费
	feeUSD := positionSizeUSD * 0.0005 * 2
	maxRiskUSD := (totalBalance * maxRiskPercent / 100) - feeUSD

	if maxRiskUSD <= 0 {
		return 0 // 无法满足风险要求
	}

	// 最大价格风险百分比
	maxPriceRiskPercent := maxRiskUSD / positionSizeUSD

	var stopLoss float64
	if side == "short" {
		// 空单：止损 = 入场价 * (1 + 风险%)
		stopLoss = entryPrice * (1 + maxPriceRiskPercent)
	} else {
		// 多单：止损 = 入场价 * (1 - 风险%)
		stopLoss = entryPrice * (1 - maxPriceRiskPercent)
	}

	return stopLoss
}

// verifyAndUpdateCloseFillPrice 验证并更新平仓的真实成交价格
// 在平仓后调用，基于交易所的成交记录获取 100% 准确的成交价格
func (at *AutoTrader) verifyAndUpdateCloseFillPrice(
	decision *decision.Decision,
	actionRecord *logger.DecisionAction,
	closeTime int64, // 平仓时间（毫秒时间戳）
) error {
	const retryDelay = 500 * time.Millisecond
	const maxRetries = 3

	log.Printf("  🔍 验证平仓真实成交价格...")

	// 定义查询时间范围：平仓前后各 10 秒
	startTime := closeTime - 10000
	endTime := closeTime + 10000

	var fills []map[string]any
	var err error

	// 重试机制：交易所可能需要时间同步成交记录
	for i := 0; i < maxRetries; i++ {
		if i > 0 {
			log.Printf("  ⏳ 等待 %v 后重试获取成交记录 (尝试 %d/%d)...", retryDelay, i+1, maxRetries)
			time.Sleep(retryDelay)
		}

		fills, err = at.trader.GetRecentFills(decision.Symbol, startTime, endTime)
		if err != nil {
			log.Printf("  ⚠️ 获取成交记录失败 (尝试 %d/%d): %v", i+1, maxRetries, err)
			continue
		}

		// 如果找到成交记录，停止重试
		if len(fills) > 0 {
			break
		}
	}

	if err != nil {
		log.Printf("  ⚠️ 无法获取成交记录，保持使用平仓前的市场价格 %.2f", actionRecord.Price)
		return nil // 不阻断流程
	}

	if len(fills) == 0 {
		log.Printf("  ⚠️ 未找到成交记录，保持使用平仓前的市场价格 %.2f", actionRecord.Price)
		return nil // 不阻断流程
	}

	// 过滤匹配的成交记录
	// close_long -> Sell
	// close_short -> Buy
	expectedSide := "Sell"
	if decision.Action == "close_short" {
		expectedSide = "Buy"
	}

	var matchedFills []map[string]any
	for _, fill := range fills {
		side, _ := fill["side"].(string)
		if side == expectedSide {
			matchedFills = append(matchedFills, fill)
		}
	}

	if len(matchedFills) == 0 {
		log.Printf("  ⚠️ 未找到匹配的 %s 成交记录，保持使用平仓前的市场价格 %.2f", expectedSide, actionRecord.Price)
		return nil
	}

	// 计算加权平均成交价格
	var totalValue float64
	var totalQuantity float64

	for _, fill := range matchedFills {
		price, _ := fill["price"].(float64)
		quantity, _ := fill["quantity"].(float64)

		totalValue += price * quantity
		totalQuantity += quantity

		log.Printf("  📊 成交记录: %.8f @ %.2f", quantity, price)
	}

	weightedAvgPrice := totalValue / totalQuantity

	// 更新 actionRecord
	oldPrice := actionRecord.Price
	actionRecord.Price = weightedAvgPrice

	log.Printf("  ✓ 成交价格已矫正: %.2f -> %.2f (共 %d 笔成交)", oldPrice, weightedAvgPrice, len(matchedFills))

	return nil
}
