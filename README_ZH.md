# Sub2API

<div align="center">

[![Go](https://img.shields.io/badge/Go-1.25.7-00ADD8.svg)](https://golang.org/)
[![Vue](https://img.shields.io/badge/Vue-3.4+-4FC08D.svg)](https://vuejs.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-336791.svg)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-7+-DC382D.svg)](https://redis.io/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED.svg)](https://www.docker.com/)

<a href="https://trendshift.io/repositories/21823" target="_blank"><img src="https://trendshift.io/api/badge/repositories/21823" alt="Wei-Shaw%2Fsub2api | Trendshift" width="250" height="55"/></a>

**AI API 閘道器平台 - 訂閱配額分發管理**

[English](README.md) | 中文 | [日本語](README_JA.md)

</div>

> **Sub2API 官方仅使用  `sub2api.org` 与 `pincc.ai` 两個域名。其他使用 Sub2API 名义的網站可能为第三方部署或服務，与本項目无關，請自行甄别。**
---

## 在線體驗

體驗地址：**[https://demo.sub2api.org/](https://demo.sub2api.org/)**

演示帳號（共享演示环境；自建部署不會自動創建該帳號）：

| 信箱 | 密碼 |
|------|------|
| admin@sub2api.org | admin123 |

## 項目概述

Sub2API 是一個 AI API 閘道器平台，用于分發和管理 AI 产品訂閱的 API 配額。使用者通過平台生成的 API Key 呼叫上游 AI 服務，平台负责鉴權、計費、負載平衡和請求轉發。

## 核心功能

- **多帳號管理** - 支持多种上游帳號类型（OAuth、API Key）
- **API Key 分發** - 为使用者生成和管理 API Key
- **精確計費** - Token 級别的用量追踪和成本計算
- **智慧調度** - 智慧帳號選擇，支持粘性會話
- **並行控制** - 使用者級和帳號級並行限制
- **速率限制** - 可配置的請求和 Token 速率限制
- **内置支付系統** - 支持 EasyPay 易支付、支付宝官方、微信官方、Stripe，使用者自助儲值，無需独立部署支付服務（[配置指南](docs/PAYMENT_CN.md)）
- **管理後台** - Web 界面進行監控和管理
- **外部系統集成** - 支持通過 iframe 嵌入外部系統（如工單等），扩展管理後台功能

## ❤️ 赞助商

> [想出現在這里？](mailto:support@pincc.ai)

<table>
<tr>
<td width="180" align="center" valign="middle"><a href="https://shop.pincc.ai/"><img src="assets/partners/logos/pincc-logo.png" alt="pincc" width="150"></a></td>
<td valign="middle"><b><a href="https://shop.pincc.ai/">PinCC</a></b> 是基于 Sub2API 搭建的官方中轉服務，提供 Claude Code、Codex、Gemini 等主流模型的穩定中轉，開箱即用，免去自建部署与運維烦恼。</td>
</tr>

<tr>
<td width="180"><a href="https://www.packyapi.com/register?aff=sub2api"><img src="assets/partners/logos/packycode.png" alt="PackyCode" width="150"></a></td>
<td>感谢 PackyCode 赞助了本項目！PackyCode 是一家穩定、高效的API中轉服務商，提供 Claude Code、Codex、Gemini 等多种中轉服務。PackyCode 为本软件的使用者提供了特别优惠，使用<a href="https://www.packyapi.com/register?aff=sub2api">此鏈接</a>註冊并在儲值時填寫"sub2api"优惠碼，首次儲值可以享受9折优惠！</td>
</tr>

<tr>
<td width="180"><a href="https://poixe.com/i/sub2api"><img src="assets/partners/logos/poixe.png" alt="PoixeAI" width="150"></a></td>
<td>感谢 Poixe AI 赞助了本項目！Poixe AI 提供可靠的 AI 模型接口服務，您可以使用平台提供的 LLM API 接口輕松構建 AI 产品，同時也可以成为供應商，为平台提供大模型資源以赚取收益。通過 <a href="https://poixe.com/i/sub2api">此鏈接</a> 专属鏈接註冊，儲值額外赠送 $5 美金</td>
</tr>

<tr>
<td width="180"><a href="https://ctok.ai"><img src="assets/partners/logos/ctok.png" alt="CTok" width="150"></a></td>
<td>感谢 CTok.ai 赞助了本項目！CTok.ai 致力于打造一站式 AI 編程工具服務平台。我们提供 Claude Code 专业套餐及技术社群服務，同時支持 Google Gemini 和 OpenAI Codex。通過精心設計的套餐方案和专业的技术社群，为開發者提供穩定的服務保障和持续的技术支持，让 AI 辅助編程真正成为開發者的生产力工具。點击<a href="https://ctok.ai">這里</a>註冊！</td>
</tr>

<tr>
<td width="180"><a href="https://code.silkapi.com/"><img src="assets/partners/logos/silkapi.png" alt="silkapi" width="150"></a></td>
<td>感谢 丝绸API 赞助了本項目！ <a href="https://code.silkapi.com/">丝绸API</a> 是基于 Sub2API 搭建的中轉服務，专註于提供 Codex 高速穩定API中轉。</td>
</tr>

<tr>
<td width="180"><a href="https://ylscode.com/"><img src="assets/partners/logos/ylscode.png" alt="ylscode" width="150"></a></td>
<td>感谢 伊莉思Code 赞助了本項目！ <a href="https://ylscode.com/">伊莉思Code</a> 致力于構建安全的企业級Coding Agent生产力服務，提供穩定快速的 Codex / Claude / Gemini 訂閱服務与即用即付API多种方案灵活選擇，限時註冊赠送 3 天 Codex 试用福利！</td>
</tr>

<tr>
<td width="180"><a href="https://www.aicodemirror.com/register?invitecode=KMVZQM"><img src="assets/partners/logos/AICodeMirror.jpg" alt="AICodeMirror" width="150"></a></td>
<td>感谢 AICodeMirror 赞助了本項目！AICodeMirror 提供 Claude Code / Codex / Gemini CLI 官方高穩定性中轉服務，企业級並行、快速開票、7×24 小時专属技术支持。Claude Code / Codex / Gemini 官方通道低至原价 38% / 2% / 9%，儲值更享額外折扣！AICodeMirror 为 sub2api 使用者提供专属福利：通過<a href="https://www.aicodemirror.com/register?invitecode=KMVZQM">此鏈接</a>註冊，首次儲值立享 8 折优惠，企业客戶最高可享 75 折！</td>
</tr>

<tr>
<td width="180"><a href="https://aigocode.com/invite/SUB2API"><img src="assets/partners/logos/aigocode.png" alt="AIGoCode" width="150"></a></td>
<td>感谢 AIGoCode 赞助了本項目！AIGoCode 是一站式集成 Claude Code、Codex 以及最新 Gemini 模型的综合平台，为您提供穩定、高效、高性价比的 AI 編程服務。平台提供灵活的訂閱方案，零封号风险，免 VPN 直連，响應极速。AIGoCode 为 sub2api 使用者準備了专属福利：通過<a href="https://aigocode.com/invite/SUB2API">此鏈接</a>註冊，首次儲值可額外获得 10% 赠送額度！</td>
</tr>

<tr>
<td width="180"><a href="https://shop.bmoplus.com/?utm_source=github"><img src="assets/partners/logos/bmoplus.jpg" alt="bmoplus" width="150"></a></td>
<td>感谢 BmoPlus 赞助了本項目！BmoPlus 是一家专为AI訂閱重度使用者打造的可靠 AI 帳號代充服務商，提供穩定的 ChatGPT Plus / ChatGPT Pro(全程质保) / Claude Pro / Super Grok / Gemini Pro 的官方代充&成品帳號。 通過<a href="https://shop.bmoplus.com/?utm_source=github">BmoPlus AI成品号专卖/代充</a>註冊下單的使用者，可享GPT 官網訂閱一折 的震撼价格！</td>
</tr>

<tr>
<td width="180"><a href="https://bestproxy.com/?keyword=a2e8iuol"><img src="assets/partners/logos/bestproxy.png" alt="bestproxy" width="150"></a></td>
<td>感谢 Bestproxy 赞助了本項目！<a href="https://bestproxy.com/?keyword=a2e8iuol">Bestproxy</a> 是一家提供高纯度住宅IP，支持一号一IP独享，結合真實家庭網絡与指纹隔离，可實現鏈路环境隔离，降低關联风控概率。</td>
</tr>

<tr>
<td width="180"><a href="https://pateway.ai/?ch=1tsfr51"><img src="assets/partners/logos/pateway.png" alt="pateway" width="150"></a></td>
<td>感谢 PatewayAI 赞助了本項目！PatewayAI 是一家面向重度 AI 開發者、专註官方直連的高品质模型 API 中轉服務商。提供 Claude 全系列与 Codex 系列模型，100% 官方源直供，不掺假不註水，歡迎检驗。計費透明，Token 級账單可逐笔核驗。
同時支持企业級高並行，并为企业客戶提供了专业的管理平台，企业客戶可签訂正式合同并開具發票，更多詳情進入官網获取联系方式。
現在通過 <a href="https://pateway.ai/?ch=1tsfr51">此鏈接</a> 註冊即送 $3 试用額度，使用者儲值低至 6 折，邀請好友双向赠送，邀請奖励可达 $150。</td>
</tr>

</table>

## 生态項目

圍绕 Sub2API 的社区扩展与集成項目：

| 項目 | 說明 | 功能 |
|------|------|------|
| ~~[Sub2ApiPay](https://github.com/touwaeriol/sub2apipay)~~ | ~~自助支付系統~~ | **已内置** — 支付功能已集成到 Sub2API 中，無需独立部署。詳見 [支付配置指南](docs/PAYMENT_CN.md) |
| [sub2api-mobile](https://github.com/ckken/sub2api-mobile) | 移動端管理控制台 | 跨平台應用（iOS/Android/Web），支持使用者管理、帳號管理、監控看板、多後端切換；基于 Expo + React Native 構建 |

## 技术栈

| 組件 | 技术 |
|------|------|
| 後端 | Go 1.25.7, Gin, Ent |
| 前端 | Vue 3.4+, Vite 5+, TailwindCSS |
| 數據庫 | PostgreSQL 15+ |
| 緩存/队列 | Redis 7+ |

---

## Nginx 反向代理註意事項

通過 Nginx 反向代理 Sub2API（或 CRS 服務）并搭配 Codex CLI 使用時，需要在 Nginx 配置的 `http` 塊中添加：

```nginx
underscores_in_headers on;
```

Nginx 默認會丢弃名稱中含下划線的請求头（如 `session_id`），這會導致多帳號环境下的粘性會話功能失效。

---

## 部署方式

### 方式一：指令碼安裝（推荐）

一鍵安裝指令碼，自動从 GitHub Releases 下載预編译的二進制文件。

#### 前置条件

- Linux 服務器（amd64 或 arm64）
- PostgreSQL 15+（已安裝并運行）
- Redis 7+（已安裝并運行）
- Root 權限

#### 安裝步骤

```bash
curl -sSL https://raw.githubusercontent.com/Wei-Shaw/sub2api/main/deploy/install.sh | sudo bash
```

指令碼會自動：
1. 檢測系統架構
2. 下載最新版本
3. 安裝二進制文件到 `/opt/sub2api`
4. 創建 systemd 服務
5. 配置系統使用者和權限

#### 安裝後配置

```bash
# 1. 啟動服務
sudo systemctl start sub2api

# 2. 設定開機自啟
sudo systemctl enable sub2api

# 3. 在浏览器中打開設定向導
# http://你的服務器IP:8080
```

設定向導將引導你完成：
- 數據庫配置
- Redis 配置
- 管理員帳號創建

#### 升級

可以直接在 **管理後台** 左上角點击 **檢測更新** 按钮進行在線升級。

網頁升級功能支持：
- 自動檢測新版本
- 一鍵下載并應用更新
- 支持回滾

#### 常用命令

```bash
# 查看状态
sudo systemctl status sub2api

# 查看日志
sudo journalctl -u sub2api -f

# 重啟服務
sudo systemctl restart sub2api

# 卸載
curl -sSL https://raw.githubusercontent.com/Wei-Shaw/sub2api/main/deploy/install.sh | sudo bash -s -- uninstall -y
```

---

### 方式二：Docker Compose（推荐）

使用 Docker Compose 部署，包含 PostgreSQL 和 Redis 容器。

#### 前置条件

- Docker 20.10+
- Docker Compose v2+

#### 快速開始（一鍵部署）

使用自動化部署指令碼快速搭建：

```bash
# 創建部署目錄
mkdir -p sub2api-deploy && cd sub2api-deploy

# 下載并運行部署準備指令碼
curl -sSL https://raw.githubusercontent.com/Wei-Shaw/sub2api/main/deploy/docker-deploy.sh | bash

# 啟動服務
docker compose up -d

# 查看日志
docker compose logs -f sub2api
```

**指令碼功能：**
- 下載 `docker-compose.local.yml`（本地儲存为 `docker-compose.yml`）和 `.env.example`
- 自動生成安全憑證（JWT_SECRET、TOTP_ENCRYPTION_KEY、POSTGRES_PASSWORD）
- 創建 `.env` 文件并填充自動生成的金鑰
- 創建數據目錄（使用本地目錄，便于備份和迁移）
- 显示生成的憑證供你記錄

#### 手動部署

如果你希望手動配置：

```bash
# 1. 克隆仓庫
git clone https://github.com/Wei-Shaw/sub2api.git
cd sub2api/deploy

# 2. 複製环境配置文件
cp .env.example .env

# 3. 編輯配置（生成安全密碼）
nano .env
```

**`.env` 必须配置項：**

```bash
# PostgreSQL 密碼（必需）
POSTGRES_PASSWORD=your_secure_password_here

# JWT 金鑰（推荐 - 重啟後保持使用者登入状态）
JWT_SECRET=your_jwt_secret_here

# TOTP 加密金鑰（推荐 - 重啟後保留双因素認證）
TOTP_ENCRYPTION_KEY=your_totp_key_here

# 可選：管理員帳號
ADMIN_EMAIL=admin@example.com
ADMIN_PASSWORD=your_admin_password

# 可選：自定义端口
SERVER_PORT=8080
```

**生成安全金鑰：**
```bash
# 生成 JWT_SECRET
openssl rand -hex 32

# 生成 TOTP_ENCRYPTION_KEY
openssl rand -hex 32

# 生成 POSTGRES_PASSWORD
openssl rand -hex 32
```

```bash
# 4. 創建數據目錄（本地版）
mkdir -p data postgres_data redis_data

# 5. 啟動所有服務
# 選項 A：本地目錄版（推荐 - 易于迁移）
docker compose -f docker-compose.local.yml up -d

# 選項 B：命名卷版（簡單設定）
docker compose up -d

# 6. 查看状态
docker compose -f docker-compose.local.yml ps

# 7. 查看日志
docker compose -f docker-compose.local.yml logs -f sub2api
```

#### 部署版本對比

| 版本 | 數據存储 | 迁移便利性 | 适用場景 |
|------|---------|-----------|---------|
| **docker-compose.local.yml** | 本地目錄 | ✅ 簡單（打包整個目錄） | 生产环境、频繁備份 |
| **docker-compose.yml** | 命名卷 | ⚠️ 需要 docker 命令 | 簡單設定 |

**推荐：** 使用 `docker-compose.local.yml`（指令碼部署）以便更輕松地管理數據。

#### 啟用“數據管理”功能（datamanagementd）

如需啟用管理後台“數據管理”，需要額外部署宿主機數據管理進程 `datamanagementd`。

關鍵點：

- 主進程固定探测：`/tmp/sub2api-datamanagement.sock`
- 只有該 Socket 可連通時，數據管理功能才會開啟
- Docker 場景需將宿主機 Socket 挂載到容器同路径

詳細部署步骤見：`deploy/DATAMANAGEMENTD_CN.md`

#### 訪問

在浏览器中打開 `http://你的服務器IP:8080`

如果管理員密碼是自動生成的，在日志中查找：
```bash
docker compose -f docker-compose.local.yml logs sub2api | grep "admin password"
```

#### 升級

```bash
# 拉取最新镜像并重建容器
docker compose -f docker-compose.local.yml pull
docker compose -f docker-compose.local.yml up -d
```

#### 輕松迁移（本地目錄版）

使用 `docker-compose.local.yml` 時，可以輕松迁移到新服務器：

```bash
# 源服務器
docker compose -f docker-compose.local.yml down
cd ..
tar czf sub2api-complete.tar.gz sub2api-deploy/

# 傳輸到新服務器
scp sub2api-complete.tar.gz user@new-server:/path/

# 新服務器
tar xzf sub2api-complete.tar.gz
cd sub2api-deploy/
docker compose -f docker-compose.local.yml up -d
```

#### 常用命令

```bash
# 停止所有服務
docker compose -f docker-compose.local.yml down

# 重啟
docker compose -f docker-compose.local.yml restart

# 查看所有日志
docker compose -f docker-compose.local.yml logs -f

# 删除所有數據（谨慎！）
docker compose -f docker-compose.local.yml down
rm -rf data/ postgres_data/ redis_data/
```

---

### 方式三：源碼編译

从源碼編译安裝，适合開發或定制需求。

#### 前置条件

- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Redis 7+

#### 編译步骤

```bash
# 1. 克隆仓庫
git clone https://github.com/Wei-Shaw/sub2api.git
cd sub2api

# 2. 安裝 pnpm（如果還沒有安裝）
npm install -g pnpm

# 3. 編译前端
cd frontend
pnpm install
pnpm run build
# 構建产物輸出到 ../backend/internal/web/dist/

# 4. 編译後端（嵌入前端）
cd ../backend
go build -tags embed -o sub2api ./cmd/server

# 5. 創建配置文件
cp ../deploy/config.example.yaml ./config.yaml

# 6. 編輯配置
nano config.yaml
```

> **註意：** `-tags embed` 参數會將前端嵌入到二進制文件中。不使用此参數編译的程序將不包含前端界面。

**`config.yaml` 關鍵配置：**

```yaml
server:
  host: "0.0.0.0"
  port: 8080
  mode: "release"

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "your_password"
  dbname: "sub2api"

redis:
  host: "localhost"
  port: 6379
  password: ""

jwt:
  secret: "change-this-to-a-secure-random-string"
  expire_hour: 24

default:
  user_concurrency: 5
  user_balance: 0
  api_key_prefix: "sk-"
  rate_multiplier: 1.0
```

### Sora 功能状态（暂不可用）

> ⚠️ 当前 Sora 相關功能因上游接入与媒體鏈路存在技术問題，暂時不可用。
> 現阶段請勿在生产环境依赖 Sora 能力。
> 文件中的 `gateway.sora_*` 配置仅作预留，待技术問題修複後再恢複可用。

### Sora 媒體签名 URL（功能恢複後可選）

当配置 `gateway.sora_media_signing_key` 且 `gateway.sora_media_signed_url_ttl_seconds > 0` 時，閘道器會將 Sora 輸出的媒體地址改寫为临時签名 URL（`/sora/media-signed/...`）。這样無需 API Key 即可在浏览器中直接訪問，且具備過期控制与防篡改能力（签名包含 path + query）。

```yaml
gateway:
  # /sora/media 是否强制要求 API Key（默認 false）
  sora_media_require_api_key: false
  # 媒體临時签名金鑰（为空則停用签名）
  sora_media_signing_key: "your-signing-key"
  # 临時签名 URL 有效期（秒）
  sora_media_signed_url_ttl_seconds: 900
```

> 若未配置签名金鑰，`/sora/media-signed` 將返回 503。  
> 如需更严格的訪問控制，可將 `sora_media_require_api_key` 設为 true，仅允许携带 API Key 的 `/sora/media` 訪問。

訪問策略說明：
- `/sora/media`：内部呼叫或用戶端携带 API Key 才能下載
- `/sora/media-signed`：外部可訪問，但有签名 + 過期控制

`config.yaml` 還支持以下安全相關配置：

- `cors.allowed_origins` 配置 CORS 白名單
- `security.url_allowlist` 配置上游/价格數據/CRS 主機白名單
- `security.url_allowlist.enabled` 可關閉 URL 校驗（慎用）
- `security.url_allowlist.allow_insecure_http` 關閉校驗時允许 HTTP URL
- `security.url_allowlist.allow_private_hosts` 允许私有/本地 IP 地址
- `security.response_headers.enabled` 可啟用可配置响應头過滤（關閉時使用默認白名單）
- `security.csp` 配置 Content-Security-Policy
- `billing.circuit_breaker` 計費异常時 fail-closed
- `server.trusted_proxies` 啟用可信代理解析 X-Forwarded-For
- `turnstile.required` 在 release 模式强制啟用 Turnstile

**閘道器防御纵深建議（重點）**

- `gateway.upstream_response_read_max_bytes`：限制非流式上游响應讀取大小（默認 `8MB`），用于防止异常响應導致内存放大。
- `gateway.proxy_probe_response_read_max_bytes`：限制代理探测响應讀取大小（默認 `1MB`）。
- `gateway.gemini_debug_response_headers`：默認 `false`，仅在排障時短時開啟，避免高频請求日志開销。
- `/auth/register`、`/auth/login`、`/auth/login/2fa`、`/auth/send-verify-code` 已提供服務端兜底限流（Redis 故障時 fail-close）。
- 推荐將 WAF/CDN 作为第一层防護，服務端限流与响應讀取上限作为第二层兜底；两层同時保留，避免旁路流量与误配置风险。

**⚠️ 安全警告：HTTP URL 配置**

当 `security.url_allowlist.enabled=false` 時，系統默認执行最小 URL 校驗，**拒绝 HTTP URL**，仅允许 HTTPS。要允许 HTTP URL（例如用于開發或内網测试），必须显式設定：

```yaml
security:
  url_allowlist:
    enabled: false                # 停用白名單检查
    allow_insecure_http: true     # 允许 HTTP URL（⚠️ 不安全）
```

**或通過環境變數：**

```bash
SECURITY_URL_ALLOWLIST_ENABLED=false
SECURITY_URL_ALLOWLIST_ALLOW_INSECURE_HTTP=true
```

**允许 HTTP 的风险：**
- API 金鑰和數據以**明文傳輸**（可被截获）
- 易受**中間人攻击 (MITM)**
- **不适合生产环境**

**适用場景：**
- ✅ 開發/测试环境的本地服務器（http://localhost）
- ✅ 内網可信端點
- ✅ 获取 HTTPS 前测试帳號連通性
- ❌ 生产环境（仅使用 HTTPS）

**未設定此項時的错误示例：**
```
Invalid base URL: invalid url scheme: http
```

如關閉 URL 校驗或响應头過滤，請加强網絡层防護：
- 出站訪問白名單限制上游域名/IP
- 阻断私網/回环/鏈路本地地址
- 强制仅允许 TLS 出站
- 在反向代理层移除敏感响應头

```bash
# 6. 運行應用
./sub2api
```

#### HTTP/2 (h2c) 与 HTTP/1.1 回退

後端明文端口默認支持 h2c，并保留 HTTP/1.1 回退用于 WebSocket 与旧用戶端。浏览器通常不支持 h2c，性能收益主要在反向代理或内網鏈路。

**反向代理示例（Caddy）：**

```caddyfile
transport http {
	versions h2c h1
}
```

**驗證：**

```bash
# h2c prior knowledge
curl --http2-prior-knowledge -I http://localhost:8080/health
# HTTP/1.1 回退
curl --http1.1 -I http://localhost:8080/health
# WebSocket 回退驗證（需管理員 token）
websocat -H="Sec-WebSocket-Protocol: sub2api-admin, jwt.<ADMIN_TOKEN>" ws://localhost:8080/api/v1/admin/ops/ws/qps
```

#### 開發模式

```bash
# 後端（支持热重載）
cd backend
go run ./cmd/server

# 前端（支持热重載）
cd frontend
pnpm run dev
```

#### 代碼生成

修改 `backend/ent/schema` 後，需要重新生成 Ent + Wire：

```bash
cd backend
go generate ./ent
go generate ./cmd/server
```

---

## 簡易模式

簡易模式适合個人開發者或内部团队快速使用，不依赖完整 SaaS 功能。

- 啟用方式：設定環境變數 `RUN_MODE=simple`
- 功能差异：隐藏 SaaS 相關功能，跳過計費流程
- 安全註意事項：生产环境需同時設定 `SIMPLE_MODE_CONFIRM=true` 才允许啟動

---

## Antigravity 使用說明

Sub2API 支持 [Antigravity](https://antigravity.so/) 帳號，授權後可通過专用端點訪問 Claude 和 Gemini 模型。

### 专用端點

| 端點 | 模型 |
|------|------|
| `/antigravity/v1/messages` | Claude 模型 |
| `/antigravity/v1beta/` | Gemini 模型 |

### Claude Code 配置示例

```bash
export ANTHROPIC_BASE_URL="http://localhost:8080/antigravity"
export ANTHROPIC_AUTH_TOKEN="sk-xxx"
```

### 混合調度模式

Antigravity 帳號支持可選的**混合調度**功能。開啟後，通用端點 `/v1/messages` 和 `/v1beta/` 也會調度該帳號。

> **⚠️ 註意**：Anthropic Claude 和 Antigravity Claude **不能在同一上下文中混合使用**，請通過分組功能做好隔离。


### 已知問題
在 Claude Code 中，无法自動登出Plan Mode。（正常使用原生Claude Api時，Plan 完成後，Claude Code會弹出弹出選項让使用者同意或拒绝Plan。） 
解决办法：shift + Tab，手動登出Plan mode，然後輸入内容 告訴 Claude Code 同意或拒绝 Plan
---

## 項目結構

```
sub2api/
├── backend/                  # Go 後端服務
│   ├── cmd/server/           # 應用入口
│   ├── internal/             # 内部模塊
│   │   ├── config/           # 配置管理
│   │   ├── model/            # 數據模型
│   │   ├── service/          # 业務逻輯
│   │   ├── handler/          # HTTP 處理器
│   │   └── gateway/          # API 閘道器核心
│   └── resources/            # 静态資源
│
├── frontend/                 # Vue 3 前端
│   └── src/
│       ├── api/              # API 呼叫
│       ├── stores/           # 状态管理
│       ├── views/            # 頁面組件
│       └── components/       # 通用組件
│
└── deploy/                   # 部署文件
    ├── docker-compose.yml    # Docker Compose 配置
    ├── .env.example          # Docker Compose 環境變數
    ├── config.example.yaml   # 二進制部署完整配置文件
    └── install.sh            # 一鍵安裝指令碼
```

## 免责声明

> **使用本項目前請仔細阅讀：**
>
> :rotating_light: **服務条款风险**: 使用本項目可能违反 Anthropic 的服務条款。請在使用前仔細阅讀 Anthropic 的使用者协議，使用本項目的一切风险由使用者自行承担。
>
> :book: **免责声明**: 本項目仅供技术学习和研究使用，作者不對因使用本項目導致的帳號封禁、服務中断或其他损失承担任何责任。

---

## Star History

<a href="https://star-history.com/#Wei-Shaw/sub2api&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=Wei-Shaw/sub2api&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=Wei-Shaw/sub2api&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=Wei-Shaw/sub2api&type=Date" />
 </picture>
</a>

---

## 许可證

本項目基于 [GNU 宽通用公共许可證 v3.0](LICENSE)（或更高版本）授權。

Copyright (c) 2026 Wesley Liddick

---

<div align="center">

**如果觉得有用，請给個 Star 支持一下！**

</div>
