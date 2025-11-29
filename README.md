# 🤖 NOFX - AI交易操作系统

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![React](https://img.shields.io/badge/React-18+-61DAFB?style=flat&logo=react)](https://reactjs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.0+-3178C6?style=flat&logo=typescript)](https://www.typescriptlang.org/)
[![License](https://img.shields.io/badge/License-AGPL--3.0-blue.svg)](LICENSE)
[![Backed by Amber.ac](https://img.shields.io/badge/Backed%20by-Amber.ac-orange.svg)](https://amber.ac)

## 🚀 通用AI交易操作系统

**NOFX** 是通用架构的 **AI交易操作系统（Agentic Trading OS）**。已在加密市场打通"**多智能体决策 → 统一风控 → 低延迟执行 → 真实/纸面账户复盘**"的闭环，正按同一技术栈扩展到**股票、期货、期权、外汇等所有市场**。

### 🎯 核心特性

- **通用数据与回测层**：跨市场、跨周期、跨交易所统一表示与因子库，沉淀可迁移的"策略记忆"
- **多智能体自博弈与自进化**：策略自动对战择优，按账户级 PnL 与风险约束持续迭代
- **执行与风控一体化**：低延迟路由、滑点/风控沙箱、账户级限额，一键切换市场

---

> ⚠️ **风险提示**：本系统为实验性项目，AI自动交易存在重大风险，强烈建议仅用于学习研究或小额资金测试！

---

## 🏦 支持的交易所（DEX/CEX教程）

### CEX（中心化交易所）

| 交易所 | 状态 | 注册（手续费优惠） | API设置 |
|--------|------|-------------------|---------|
| **Binance** | ✅ 已支持 | [注册](https://www.binance.com/join?ref=NOFXCHI) | [API指南](../../getting-started/binance-api.md) |
| **Bybit** | ✅ 已支持 | [注册](https://partner.bybit.com/b/83856) | [API指南](../../getting-started/bybit-api.md) |
| **OKX** | ✅ 已支持 | [注册](https://www.okx.com/join/1865360) | [API指南](../../getting-started/okx-api.md) |

### Perp-DEX（去中心化永续交易所）

| 交易所 | 状态 | 注册（手续费优惠） | 代理钱包设置 |
|--------|------|-------------------|--------------|
| **Hyperliquid** | ✅ 已支持 | [注册](https://app.hyperliquid.xyz/join/AITRADING) | [代理钱包指南](../../getting-started/hyperliquid-agent-wallet.md) |
| **Aster DEX** | ✅ 已支持 | [注册](https://www.asterdex.com/en/referral/fdfc0e) | [API钱包指南](../../getting-started/aster-api-wallet.md) |
| **Lighter** | ✅ 已支持 | [注册](https://lighter.xyz) | [代理钱包指南](../../getting-started/lighter-agent-wallet.md) |

---

## 🤖 支持的AI模型

| AI模型 | 状态 | 获取API密钥 |
|--------|------|-------------|
| **DeepSeek** | ✅ 已支持 | [获取API密钥](https://platform.deepseek.com) |
| **Qwen (通义千问)** | ✅ 已支持 | [获取API密钥](https://dashscope.console.aliyun.com) |
| **OpenAI (GPT)** | ✅ 已支持 | [获取API密钥](https://platform.openai.com) |
| **Claude** | ✅ 已支持 | [获取API密钥](https://console.anthropic.com) |
| **Gemini** | ✅ 已支持 | [获取API密钥](https://aistudio.google.com) |
| **Grok** | ✅ 已支持 | [获取API密钥](https://console.x.ai) |
| **Kimi (月之暗面)** | ✅ 已支持 | [获取API密钥](https://platform.moonshot.cn) |

---

## ✨ 当前实现 - 加密货币市场

NOFX 目前已在**加密货币市场全面运行**，具备以下经过验证的能力：

### 🏆 多智能体竞赛框架
- **实时智能体对战**：Qwen vs DeepSeek 模型实时交易竞赛
- **独立账户管理**：每个智能体维护独立的决策日志和性能指标
- **实时性能对比**：实时 ROI 追踪、胜率统计、正面对抗分析
- **自进化循环**：智能体从历史表现中学习，持续改进

### 🧠 AI 自学习与优化
- **历史反馈系统**：每次决策前分析最近 20 个交易周期
- **智能性能分析**：
  - 识别表现最佳/最差资产
  - 计算胜率、盈亏比、以真实 USDT 计的平均盈利
  - 避免重复错误（连续亏损模式）
  - 强化成功策略（高胜率模式）
- **动态策略调整**：AI 根据回测结果自主调整交易风格

### 📊 通用市场数据层（加密货币实现）
- **多时间框架分析**：3分钟实时 + 4小时趋势数据
- **技术指标**：EMA20/50、MACD、RSI(7/14)、ATR
- **持仓量追踪**：市场情绪、资金流向分析
- **流动性过滤**：自动过滤低流动性资产（<15M USD）
- **跨交易所支持**：Binance、Hyperliquid、Aster DEX，统一数据接口

### 🎯 统一风控系统
- **仓位限制**：单资产限制（山寨币≤1.5x净值，BTC/ETH≤10x净值）
- **可配置杠杆**：根据资产类别和账户类型动态调整 1x 到 50x
- **保证金管理**：总使用率≤90%，AI 控制分配
- **风险回报强制执行**：强制≥1:2 的止损止盈比
- **防叠加保护**：防止同一资产/方向的重复仓位

### ⚡ 低延迟执行引擎
- **多交易所 API 集成**：Binance Futures、Hyperliquid DEX、Aster DEX
- **自动精度处理**：每个交易所智能订单大小和价格格式化
- **优先级执行**：先平仓现有持仓，再开新仓
- **滑点控制**：执行前验证，实时精度检查

### 🎨 专业监控界面
- **币安风格仪表板**：专业暗色主题，实时更新
- **净值曲线**：历史账户价值追踪（USD/百分比切换）
- **性能图表**：多智能体 ROI 对比，实时更新
- **完整决策日志**：每笔交易的完整思维链（CoT）推理
- **5秒数据刷新**：实时账户、持仓和盈亏更新

---

## 🏗️ 技术架构

NOFX 采用现代化的模块化架构：

- **后端：** Go + Gin 框架，SQLite 数据库
- **前端：** React 18 + TypeScript + Vite + TailwindCSS
- **多交易所支持：** Binance、Hyperliquid、Aster DEX
- **AI 集成：** DeepSeek、Qwen 及自定义 OpenAI 兼容 API
- **状态管理：** 前端 Zustand，后端数据库驱动
- **实时更新：** SWR，5-10 秒轮询间隔

**核心特性：**
- 🗄️ 数据库驱动的配置（无需编辑 JSON）
- 🔐 JWT 认证，支持可选的 2FA
- 📊 实时性能跟踪和分析
- 🤖 多 AI 竞赛模式，实时对比
- 🔌 RESTful API，完整的配置和监控

---

## 🚀 快速开始

### 🐳 方式A：Docker 一键部署（最简单 - 新手推荐！）

**⚡ 使用Docker只需3步即可开始交易 - 无需安装任何环境！**

Docker会自动处理所有依赖（Go、Node.js、TA-Lib）和环境配置，完美适合新手！

#### 步骤1：准备配置文件
```bash
# 复制配置文件模板
cp config.json.example config.json

# 编辑并填入你的API密钥
nano config.json  # 或使用其他编辑器
```

⚠️ **注意**: 基础config.json仍需要一些设置，但~~交易员配置~~现在通过Web界面进行。

#### 步骤2：一键启动
```bash
# 方式1：使用便捷脚本（推荐）
chmod +x scripts/start.sh
./scripts/start.sh start --build


# 方式2：直接使用docker compose
# 如果您还在使用旧的独立 `docker-compose`，请升级到 Docker Desktop 或 Docker 20.10+
docker compose up -d --build
```

#### 步骤3：访问控制台
在浏览器中打开：**http://localhost:3000**

**就是这么简单！🎉** 你的AI交易系统已经运行起来了！

#### 管理你的系统
```bash
./scripts/start.sh logs      # 查看日志
./scripts/start.sh status    # 检查状态
./scripts/start.sh stop      # 停止服务
./scripts/start.sh restart   # 重启服务
```

**📖 详细的Docker部署教程、故障排查和高级配置：**
- **中文**: 查看 [DOCKER_DEPLOY.md](DOCKER_DEPLOY.md)
- **English**: See [DOCKER_DEPLOY.en.md](DOCKER_DEPLOY.en.md)
- **日本語**: [DOCKER_DEPLOY.ja.md](DOCKER_DEPLOY.ja.md)を参照

---

### 📦 方式B：手动安装（开发者）

**注意**：如果你使用了上面的Docker部署，请跳过本节。手动安装仅在你需要修改代码或不想使用Docker时需要。

### 1. 环境要求

- **Go 1.21+**
- **Node.js 18+**
- **TA-Lib** 库（技术指标计算）

#### 安装 TA-Lib

**macOS:**
```bash
brew install ta-lib
```

**Ubuntu/Debian:**
```bash
sudo apt-get install libta-lib0-dev
```

**其他系统**: 参考 [TA-Lib官方文档](https://github.com/markcheno/go-talib)

### 2. 克隆项目

```bash
git clone <repository-url>
cd nofx
```

### 3. 安装依赖

**后端:**
```bash
go mod download
```

**前端:**
```bash
cd web
npm install
cd ..
```

### 4. 获取AI API密钥

在配置系统之前，您需要获取AI API密钥。请选择以下AI提供商之一：

#### 选项1：DeepSeek（推荐新手）

**为什么选择DeepSeek？**
- 💰 比GPT-4便宜（约1/10成本）
- 🚀 响应速度快
- 🎯 交易决策质量优秀
- 🌍 全球可用无需VPN

**如何获取DeepSeek API密钥：**

1. **访问**：[https://platform.deepseek.com](https://platform.deepseek.com)
2. **注册**：使用邮箱/手机号注册
3. **验证**：完成邮箱/手机验证
4. **充值**：向账户添加余额
   - 最低：约$5美元
   - 推荐：$20-50美元用于测试
5. **创建API密钥**：
   - 进入API Keys部分
   - 点击"创建新密钥"
   - 复制并保存密钥（以`sk-`开头）
   - ⚠️ **重要**：立即保存 - 之后无法再查看！

**价格**：每百万tokens约$0.14（非常便宜！）

#### 选项2：Qwen（阿里云通义千问）

**如何获取Qwen API密钥：**

1. **访问**：[https://dashscope.console.aliyun.com](https://dashscope.console.aliyun.com)
2. **注册**：使用阿里云账户注册
3. **开通服务**：激活DashScope服务
4. **创建API密钥**：
   - 进入API密钥管理
   - 创建新密钥
   - 复制并保存（以`sk-`开头）

**注意**：可能需要中国手机号注册

---

#### ⚙️ 杠杆配置 (v2.0.3+)

**什么是杠杆配置？**

杠杆设置控制AI每次交易可以使用的最大杠杆。这对于风险管理至关重要，特别是对于有杠杆限制的币安子账户。

**配置格式：**

```json
"leverage": {
  "btc_eth_leverage": 5,    // BTC和ETH的最大杠杆
  "altcoin_leverage": 5      // 所有其他币种的最大杠杆
}
```

**⚠️ 重要：币安子账户限制**

- **子账户**：币安限制为**≤5倍杠杆**
- **主账户**：可使用最高20倍（山寨币）或50倍（BTC/ETH）
- 如果您使用子账户并设置杠杆>5倍，交易将**失败**，错误信息：`Subaccounts are restricted from using leverage greater than 5x`

**推荐设置：**

| 账户类型 | BTC/ETH杠杆 | 山寨币杠杆 | 风险级别 |
|---------|------------|-----------|---------|
| **子账户** | `5` | `5` | ✅ 安全（默认） |
| **主账户（保守）** | `10` | `10` | 🟡 中等 |
| **主账户（激进）** | `20` | `15` | 🔴 高 |
| **主账户（最大）** | `50` | `20` | 🔴🔴 非常高 |

**示例：**

**安全配置（子账户或保守）：**
```json
"leverage": {
  "btc_eth_leverage": 5,
  "altcoin_leverage": 5
}
```

**激进配置（仅主账户）：**
```json
"leverage": {
  "btc_eth_leverage": 20,
  "altcoin_leverage": 15
}
```

**AI如何使用杠杆：**

- AI可以选择**从1倍到您配置的最大值之间的任何杠杆**
- 例如，当`altcoin_leverage: 20`时，AI可能根据市场情况决定使用5倍、10倍或20倍
- 配置设置的是**上限**，而不是固定值
- AI在选择杠杆时会考虑波动性、风险回报比和账户余额

---

#### ⚠️ 重要：`use_default_coins` 字段

**智能默认行为（v2.0.2+）：**

系统现在会自动默认为`use_default_coins: true`，如果：
- 您在config.json中未包含此字段，或
- 您将其设为`false`但未提供`coin_pool_api_url`

这让新手更友好！您甚至可以完全省略此字段。

**配置示例：**

✅ **选项1：显式设置（推荐以保持清晰）**
```json
"use_default_coins": true,
"coin_pool_api_url": "",
"oi_top_api_url": ""
```

✅ **选项2：省略字段（自动使用默认币种）**
```json
// 完全不包含"use_default_coins"
"coin_pool_api_url": "",
"oi_top_api_url": ""
```

⚙️ **高级：使用外部API**
```json
"use_default_coins": false,
"coin_pool_api_url": "http://your-api.com/coins",
"oi_top_api_url": "http://your-api.com/oi"
```

---

### 6. 运行系统

#### 🚀 启动系统（2个步骤）

系统有**2个部分**需要分别运行：
1. **后端**（AI交易大脑 + API）
2. **前端**（Web监控仪表板）

---

#### **步骤1：启动后端**

打开终端并运行：

```bash
# 构建程序（首次运行或代码更改后）
go build -o nofx

# 启动后端
./nofx
```

**您应该看到：**

```
🚀 启动自动交易系统...
✓ Trader [my_trader] 已初始化
✓ API服务器启动在端口 8080
📊 开始交易监控...
```

**⚠️ 如果看到错误：**

| 错误信息 | 解决方案 |
|---------|---------|
| `invalid API key` | 检查Web界面中的API密钥 |
| `TA-Lib not found` | 运行`brew install ta-lib`（macOS） |
| `port 8080 already in use` | 修改.env文件中的`API_PORT` |
| `DeepSeek API error` | 验证DeepSeek API密钥和余额 |

**✅ 后端运行正常的标志：**
- 无错误信息
- 出现"开始交易监控..."
- 系统显示账户余额
- 保持此终端窗口打开！

---

#### **步骤2：启动前端**

打开**新的终端窗口**（保持第一个运行！），然后：

```bash
cd web
npm run dev
```

**您应该看到：**

```
VITE v5.x.x  ready in xxx ms

➜  Local:   http://localhost:3000/
➜  Network: use --host to expose
```

**✅ 前端运行正常的标志：**
- "Local: http://localhost:3000/"消息
- 无错误信息
- 也保持此终端窗口打开！

---

#### **步骤3：访问仪表板**

在Web浏览器中访问：

**🌐 http://localhost:3000**

**您将看到：**
- 📊 实时账户余额
- 📈 持仓（如果有）
- 🤖 AI决策日志
- 📉 净值曲线图

**首次使用提示：**
- 首次AI决策可能需要3-5分钟
- 初始决策可能显示"观望"- 这是正常的
- AI需要先分析市场状况

---

### 7. 监控系统

**需要关注的内容：**

✅ **健康系统标志：**
- 后端终端每3-5分钟显示决策周期
- 无持续错误信息
- 账户余额更新
- Web仪表板自动刷新

⚠️ **警告标志：**
- 重复的API错误
- 10分钟以上无决策
- 余额快速下降

**检查系统状态：**

```bash
# 在新终端窗口中
curl http://localhost:8080/api/health
```

应返回：`{"status":"ok"}`

---

### 8. 停止系统

**优雅关闭（推荐）：**

1. 转到**后端终端**（第一个）
2. 按`Ctrl+C`
3. 等待"系统已停止"消息
4. 转到**前端终端**（第二个）
5. 按`Ctrl+C`

**⚠️ 重要：**
- 始终先停止后端
- 关闭终端前等待确认
- 不要强制退出（不要直接关闭终端）

---

## 📖 AI决策流程

每个决策周期（默认3分钟），系统按以下流程运行：

### 步骤1: 📊 分析历史表现（最近20个周期）
- ✓ 计算整体胜率、平均盈利、盈亏比
- ✓ 统计各币种表现（胜率、平均USDT盈亏）
- ✓ 识别最佳/最差币种
- ✓ 列出最近5笔交易详情（含准确盈亏金额）
- ✓ 计算夏普比率衡量风险调整后收益
- 📌 **新增 (v2.0.2)**: 考虑杠杆的准确USDT盈亏计算

**↓**

### 步骤2: 💰 获取账户状态
- 账户净值、可用余额、未实现盈亏
- 持仓数量、总盈亏（已实现+未实现）
- 保证金使用率（current/maximum）
- 风险评估指标

**↓**

### 步骤3: 🔍 分析现有持仓（如果有）
- 获取每个持仓的市场数据（3分钟+4小时K线）
- 计算技术指标（RSI、MACD、EMA）
- 显示持仓时长（例如"持仓时长2小时15分钟"）
- AI判断是否需要平仓（止盈、止损或调整）
- 📌 **新增 (v2.0.2)**: 追踪持仓时长帮助AI决策

**↓**

### 步骤4: 🎯 评估新机会（候选币种池）
- 获取币种池（2种模式）：
  - 🌟 **默认模式**: BTC、ETH、SOL、BNB、XRP等
  - ⚙️ **高级模式**: AI500（前20） + OI Top（前20）
- 合并去重，过滤低流动性币种（持仓量<15M USD）
- 批量获取市场数据和技术指标
- 为每个候选币种准备完整的原始数据序列

**↓**

### 步骤5: 🧠 AI综合决策
- 查看历史反馈（胜率、盈亏比、最佳/最差币种）
- 接收所有原始序列数据（K线、指标、持仓量）
- Chain of Thought 思维链分析
- 输出决策：平仓/开仓/持有/观望
- 包含杠杆、仓位、止损、止盈参数
- 📌 **新增 (v2.0.2)**: AI可自由分析原始序列，不受预定义指标限制

**↓**

### 步骤6: ⚡ 执行交易
- 优先级排序：先平仓，再开仓
- 精度自动适配（LOT_SIZE规则）
- 防止仓位叠加（同币种同方向拒绝开仓）
- 平仓后自动取消所有挂单
- 记录开仓时间用于持仓时长追踪
- 📌 追踪持仓开仓时间

**↓**

### 步骤7: 📝 记录日志
- 保存完整决策记录到 `decision_logs/`
- 包含思维链、决策JSON、账户快照、执行结果
- 存储完整持仓数据（数量、杠杆、开/平仓时间）
- 使用 `symbol_side` 键值防止多空冲突
- 📌 **新增 (v2.0.2)**: 防止多空持仓冲突，考虑数量+杠杆

**↓**

**🔄 （每3-5分钟重复一次）**

### v2.0.2的核心改进

**📌 持仓时长追踪：**
- 系统现在追踪每个持仓已持有多长时间
- 在用户提示中显示："持仓时长2小时15分钟"
- 帮助AI更好地判断何时退出仓位

**📌 准确的盈亏计算：**
- 之前：只显示百分比（100U@5% = 1000U@5% = 都显示"5.0"）
- 现在：真实USDT盈亏 = 仓位价值 × 价格变化% × 杠杆倍数
- 示例：1000 USDT × 5% × 20倍 = 1000 USDT实际盈利

**📌 增强的AI自由度：**
- AI可以自由分析所有原始序列数据
- 不再局限于预定义的指标组合
- 可以执行自己的趋势分析、支撑位/阻力位计算

**📌 改进的持仓追踪：**
- 使用`symbol_side`键值（例如"BTCUSDT_long"）
- 防止同时持有多空仓时的冲突
- 存储完整数据：数量、杠杆、开/平仓时间

---

## 🧠 AI自我学习示例

### 历史反馈（Prompt中自动添加）

```markdown
## 📊 历史表现反馈

### 整体表现
- **总交易数**: 15 笔 (盈利: 8 | 亏损: 7)
- **胜率**: 53.3%
- **平均盈利**: +3.2% | 平均亏损: -2.1%
- **盈亏比**: 1.52:1

### 最近交易
1. BTCUSDT LONG: 95000.0000 → 97500.0000 = +2.63% ✓
2. ETHUSDT SHORT: 3500.0000 → 3450.0000 = +1.43% ✓
3. SOLUSDT LONG: 185.0000 → 180.0000 = -2.70% ✗
4. BNBUSDT LONG: 610.0000 → 625.0000 = +2.46% ✓
5. ADAUSDT LONG: 0.8500 → 0.8300 = -2.35% ✗

### 币种表现
- **最佳**: BTCUSDT (胜率75%, 平均+2.5%)
- **最差**: SOLUSDT (胜率25%, 平均-1.8%)
```

### AI如何使用反馈

1. **避免连续亏损币种**: 看到SOLUSDT连续3次止损，AI会避开或更谨慎
2. **强化成功策略**: BTC突破做多胜率75%，AI会继续这个模式
3. **动态调整风格**: 胜率<40%时变保守，盈亏比>2时保持激进
4. **识别市场环境**: 连续亏损可能说明市场震荡，减少交易频率

---

## 📊 Web界面功能

### 1. 竞赛页面（Competition）

- **🏆 排行榜**: 实时收益率排名，金色边框突出显示领先者
- **📈 性能对比图**: 双AI收益率曲线对比（紫色vs蓝色）
- **⚔️ Head-to-Head**: 直接对比，显示领先差距
- **实时数据**: 总净值、盈亏%、持仓数、保证金使用率

### 2. 详情页面（Details）

- **账户净值曲线**: 历史走势图（美元/百分比切换）
- **统计信息**: 总周期、成功/失败、开仓/平仓统计
- **持仓表格**: 所有持仓详情（入场价、当前价、盈亏%、强平价）
- **AI决策日志**: 最近决策记录（可展开思维链）

### 3. 实时更新

- 系统状态、账户信息、持仓列表：**每5秒刷新**
- 决策日志、统计信息：**每10秒刷新**
- 收益率图表：**每10秒刷新**

---

## 🎛️ API接口

### 竞赛相关

```bash
GET /api/competition          # 竞赛排行榜（所有trader）
GET /api/traders              # Trader列表
```

### 单Trader相关

```bash
GET /api/status?trader_id=xxx            # 系统状态
GET /api/account?trader_id=xxx           # 账户信息
GET /api/positions?trader_id=xxx         # 持仓列表
GET /api/equity-history?trader_id=xxx    # 净值历史（图表数据）
GET /api/decisions/latest?trader_id=xxx  # 最新5条决策
GET /api/statistics?trader_id=xxx        # 统计信息
```

### 系统接口

```bash
GET /api/health                   # 健康检查
GET /api/config               # 系统配置
```

---

## 📝 决策日志格式

每次AI决策都会生成详细的JSON日志：

### 日志文件路径
```
decision_logs/
├── qwen_trader/
│   └── decision_20251028_153042_cycle15.json
└── deepseek_trader/
    └── decision_20251028_153045_cycle15.json
```

### 日志内容示例

```json
{
  "timestamp": "2025-10-28T15:30:42+08:00",
  "cycle_number": 15,
  "cot_trace": "当前持仓：ETHUSDT多头盈利+2.3%，趋势良好继续持有...",
  "decision_json": "[{\"symbol\":\"BTCUSDT\",\"action\":\"open_long\"...}]",
  "account_state": {
    "total_balance": 1045.80,
    "available_balance": 823.40,
    "position_count": 3,
    "margin_used_pct": 21.3
  },
  "positions": [...],
  "candidate_coins": ["BTCUSDT", "ETHUSDT", ...],
  "decisions": [
    {
      "action": "open_long",
      "symbol": "BTCUSDT",
      "quantity": 0.015,
      "leverage": 50,
      "price": 95800.0,
      "order_id": 123456789,
      "success": true
    }
  ],
  "execution_log": ["✓ BTCUSDT open_long 成功"],
  "success": true
}
```

---

## 🔧 风险控制详解

### 单币种仓位限制

| 币种类型 | 仓位价值上限 | 杠杆 | 保证金占用 | 示例（1000U账户） |
|---------|-------------|------|-----------|------------------|
| 山寨币  | 1.5倍净值    | 20x  | 7.5%      | 最多开1500U仓位 = 75U保证金 |
| BTC/ETH | 10倍净值     | 50x  | 20%       | 最多开10000U仓位 = 200U保证金 |

### 为什么这样设计？

1. **高杠杆 + 小仓位 = 分散风险**
   - 20倍杠杆，1500U仓位，只需75U保证金
   - 可以同时开10+个小仓位，分散单币种风险

2. **单币种风险可控**
   - 山寨币仓位≤1.5倍净值，5%反向波动 = 7.5%损失
   - BTC仓位≤10倍净值，2%反向波动 = 20%损失

3. **不限制总保证金使用率**
   - AI根据市场机会自主决策保证金使用率
   - 上限90%，但不强制满仓
   - 有好机会就开仓，没机会就观望

### 防止过度交易

- **同币种同方向不允许重复开仓**: 防止AI连续开同一个仓位导致超限
- **先平仓后开仓**: 换仓时确保先释放保证金
- **止损止盈强制检查**: 风险回报比≥1:2

---

## ⚠️ 重要风险提示

### 交易风险

1. **加密货币市场波动极大**，AI决策不保证盈利
2. **合约交易使用杠杆**，亏损可能超过本金
3. **市场极端行情**下可能出现爆仓风险
4. **资金费率**可能影响持仓成本
5. **流动性风险**：某些币种可能出现滑点

### 技术风险

1. **网络延迟**可能导致价格滑点
2. **API限流**可能影响交易执行
3. **AI API超时**可能导致决策失败
4. **系统Bug**可能引发意外行为

### 使用建议

✅ **建议做法**
- 仅使用可承受损失的资金测试
- 从小额资金开始（建议100-500 USDT）
- 定期检查系统运行状态
- 监控账户余额变化
- 分析AI决策日志，理解策略

❌ **不建议做法**
- 投入全部资金或借贷资金
- 长时间无人监控运行
- 盲目信任AI决策
- 在不理解系统的情况下使用
- 在市场极端波动时运行

---

## 📈 性能优化建议

1. **合理设置决策周期**: 建议3-5分钟，避免过度交易
2. **控制候选币种数量**: 系统默认分析AI500前20 + OI Top前20
3. **定期清理日志**: 避免占用过多磁盘空间
4. **监控API调用次数**: 避免触发Binance限流（权重限制）
5. **小额资金测试**: 先用100-500 USDT测试策略有效性

---

## 📄 开源协议

本项目采用 **GNU Affero 通用公共许可证 v3.0 (AGPL-3.0)** - 详见 [LICENSE](LICENSE) 文件

**这意味着什么：**
- ✅ 你可以使用、修改和分发此软件
- ✅ 你必须公开你修改版本的源代码
- ✅ 如果你在服务器上运行修改版本，必须向用户提供源代码
- ✅ 所有衍生作品也必须使用 AGPL-3.0 许可证

如需商业许可或有疑问,请联系维护者。
