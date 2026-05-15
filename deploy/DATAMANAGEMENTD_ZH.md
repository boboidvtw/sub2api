# datamanagementd 部署說明（數據管理）

本文說明如何在宿主機部署 `datamanagementd`，并与主進程联動開啟“數據管理”功能。

## 1. 關鍵约束

- 主進程固定探测路径：`/tmp/sub2api-datamanagement.sock`
- 仅当該 Unix Socket 可連通且 `Health` 成功時，後台“數據管理”才會啟用
- `datamanagementd` 使用 SQLite 持久化元數據，不依赖主庫

## 2. 宿主機構建与運行

```bash
cd /opt/sub2api-src/datamanagement
go build -o /opt/sub2api/datamanagementd ./cmd/datamanagementd

mkdir -p /var/lib/sub2api/datamanagement
chown -R sub2api:sub2api /var/lib/sub2api/datamanagement
```

手動啟動示例：

```bash
/opt/sub2api/datamanagementd \
  -socket-path /tmp/sub2api-datamanagement.sock \
  -sqlite-path /var/lib/sub2api/datamanagement/datamanagementd.db \
  -version 1.0.0
```

## 3. systemd 托管（推荐）

仓庫已提供示例服務文件：`deploy/sub2api-datamanagementd.service`

```bash
sudo cp deploy/sub2api-datamanagementd.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now sub2api-datamanagementd
sudo systemctl status sub2api-datamanagementd
```

查看日志：

```bash
sudo journalctl -u sub2api-datamanagementd -f
```

也可以使用一鍵安裝指令碼（自動安裝二進制 + 註冊 systemd）：

```bash
# 方式一：使用現成二進制
sudo ./deploy/install-datamanagementd.sh --binary /path/to/datamanagementd

# 方式二：从源碼構建後安裝
sudo ./deploy/install-datamanagementd.sh --source /path/to/sub2api
```

## 4. Docker 部署联動

若 `sub2api` 運行在 Docker 容器中，需要將宿主機 Socket 挂載到容器同路径：

```yaml
services:
  sub2api:
    volumes:
      - /tmp/sub2api-datamanagement.sock:/tmp/sub2api-datamanagement.sock
```

建議在 `docker-compose.override.yml` 中維護該挂載，避免覆盖主 compose 文件。

## 5. 依赖检查

`datamanagementd` 执行備份時依赖以下工具：

- `pg_dump`
- `redis-cli`
- `docker`（仅 `source_mode=docker_exec` 時）

缺失依赖會導致對應任務失败，并在任務詳情中體現错误信息。
