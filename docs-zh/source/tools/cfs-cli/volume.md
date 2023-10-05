# 卷管理

## 创建卷

```bash
cfs-cli volume create [VOLUME NAME] [USER ID] [flags]
```
### Flags

| Flag                       | 描述                                                                                                        |
|----------------------------|-----------------------------------------------------------------------------------------------------------|
| cache-action               | 指定低容量缓存操作 (默认 0)                                                                                          |
| cache-capacity             | 指定低容量大小[单位: GB]                                                                                           |
| cache-high-water           | (默认 80)                                                                                                   |
| cache-low-water            | (默认 60)                                                                                                   |
| cache-lru-interval         | 指定缓存LRU间隔时间[单位: 分钟] (默认 5)                                                                                |
| cache-rule-key             | 匹配此字段的内容将被写入缓存                                                                                            |
| cache-threshold            | 指定缓存阈值[单位: 字节] (默认 10485760)                                                                              |
| cache-ttl                  | 指定缓存过期时间[单位: 天] (默认 30)                                                                                   |
| capacity                   | 指定卷的容量 (默认 10)                                                                                            |
| clientIDKey                | 如果集群验证开启，则需要此字段                                                                                           |
| crossZone                  | 禁用跨区域 (默认 "false")                                                                                        |
| delete-lock-time           | 指定卷的删除锁定时间[单位: 小时]                                                                                        |
| description                | 描述                                                                                                        |
| ebs-blk-size               | 指定EBS块大小[单位: 字节] (默认 8388608)                                                                             |
| enableQuota                | 开启配额 (默认为 "false")                                                                                        |
| follower-read              | 允许从跟随者副本读取                                                                                                |
| mp-count                   | 指定初始元数据分区数量 (默认 3)                                                                                        |
| normalZonesFirst           | 首先写入正常区域 (默认 "false")                                                                                     |
| readonly-when-full         | 当卷满时设置为只读 (默认 "false")                                                                                    |
| replica-num                | 指定数据分区副本数量(默认为3对于普通卷，1对于低容量卷)                                                                             |
| size                       | 指定数据分区大小[单位: GB] (默认 120)                                                                                 |
| transaction-mask           | 为指定操作启用事务：["create&#124;mkdir&#124;remove&#124;rename&#124;mknod&#124;symlink&#124;link"] 或 "off" 或 "all" |
| transaction-timeout        | 指定事务超时时间[单位: 分钟] (默认 1)                                                                                   |
| tx-conflict-retry-Interval | 指定事务冲突重试间隔[单位: 毫秒]                                                                                        |
| tx-conflict-retry-num      | 指定事务冲突重试次数                                                                                                |
| vol-type                   | 卷的类型 (默认 0)                                                                                               |
| y, yes                     | 对所有问题回答“yes”                                                                                              |
| zone-name                  | 指定卷的区域名称                                                                                                  |


## 删除指定卷

删除指定卷[VOLUME NAME], ec卷大小为0才能删除

```bash
cfs-cli volume delete [VOLUME NAME] [flags]
```

### Flags

| 名称     | 描述                                 |
|--------|------------------------------------|
| y, yes | 跳过所有问题并设置回答为"yes"                  |

## 获取卷信息

获取卷[VOLUME NAME]的信息

```bash
cfs-cli volume info [VOLUME NAME] [flags]
```

### Flags

| Name              | Description  |
|-------------------|--------------|
| d, data-partition | 显示数据分片的详细信息  |
| m, meta-partition | 显示元数据分片的详细信息 |

## 创建并添加的数据分片至卷

创建并添加个数为[NUMBER]的数据分片至卷[VOLUME]

```bash
cfs-cli volume add-dp [VOLUME] [NUMBER]
```

## 列出所有卷信息

获取包含当前所有卷信息的列表

```bash
cfs-cli volume list
```

## 将卷转交给其他用户

将卷[VOLUME NAME]转交给其他用户[USER ID]

```bash
cfs-cli volume transfer [VOLUME NAME] [USER ID] [flags]
```

### Flags

| 名称       | 描述                |
|----------|-------------------|
| f, force | 强制转交              |
| y, yes   | 跳过所有问题并设置回答为"yes" |

## 更新卷配置

```bash
cfs-cli volume update
```

### Flags

| 名称                         | 描述                                                                                                                   |
|----------------------------|----------------------------------------------------------------------------------------------------------------------|
| cache-action               | 设置低容量的缓存动作（默认为0）                                                                                                     |
| cache-capacity             | 指定低容量的容量[单位: GB]                                                                                                     |
| cache-high-water           | 缓存高水位 (默认 80)                                                                                                        |
| cache-low-water            | 缓存低水位 (默认 60)                                                                                                        |
| cache-lru-interval         | 指定缓存清除间隔[单位: 分] (默认 5分钟)                                                                                             |
| cache-rule                 | 指定缓存规则                                                                                                               |
| cache-threshold            | 指定缓存阈值[单位: 字节] (默认 10M)                                                                                              |
| cache-ttl                  | 指定缓存的过期时间[单位: 天] (默认 30天)                                                                                            |
| capacity                   | 指定数据节点的容量[单位: GB]                                                                                                    |
| clientIDKey                | 如果集群需要身份验证，这是必需的                                                                                                     |
| delete-lock-time           | 指定卷的删除锁定时间[单位: 小时] (默认 -1小时)                                                                                         |
| description                | 卷的描述                                                                                                                 |
| ebs-blk-size               | 指定ebsBlk大小[单位: 字节]                                                                                                   |
| enableQuota                | 启用配额                                                                                                                 |
| follower-read              | 是否允许从副本跟随者读取 (默认 false)                                                                                              |
| readonly-when-full         | 当卷满时，是否只读                                                                                                            |
| replica-num                | 指定数据分区的副本数量 (普通卷默认为3，低容量卷默认为1)                                                                                       |
| transaction-force-reset    | 将交易掩码重置为"transaction-mask"的指定值                                                                                       |
| transaction-limit          | 指定交易限制[单位: 秒] (默认无限制)                                                                                                |
| transaction-mask           | 为指定操作启用交易，如: ["create&#124;mkdir&#124;remove&#124;rename&#124;mknod&#124;symlink&#124;link&#124;pause"]，或"off"或"all" |
| transaction-timeout        | 指定交易的超时时间[单位: 分] (范围 0-60)                                                                                           |
| tx-conflict-retry-Interval | 指定交易冲突的重试间隔[单位: 毫秒] (范围 10-1000)                                                                                     |
| tx-conflict-retry-num      | 指定交易冲突的重试次数 (范围 1-100)                                                                                               |
| y, yes                     | 对所有问题回答"yes"                                                                                                         |
| zone-name                  | 指定卷的区域名称                                                                                                             |
