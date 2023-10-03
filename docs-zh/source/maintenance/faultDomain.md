# 故障域配置及管理

## 什么是故障域

故障域（FaultDomain）是一组区域（Zone）的集合。若启用故障域，分片的各个副本会分布在不同的zone中，以提高数据的可靠性。

在未启用状态下，分片的所有副本会分布在同一个zone中。

## 升级及配置项

### Cluster级别配置

启用故障域需要增加cluster级别的配置，否则无法区分，新增zone是故障域zone还是归属于原有cross_zone

```bash
FaultDomain               bool  // 默认false
```

### Volume级别配置

保留：

```bash
crossZone        bool  # 跨zone
```

新增：

```bash
default_priority  bool  # true优先选择原有的zone，而不是从故障域里面分配
```

### 分片的分区数量配置

```bash
faultDomainGrpBatchCnt   int # 默认为3，也可以设置2或1
```

在网络分区中断导致zone不可用的情况下，决定是否根据可用zone创建nodeset组，则使用以下配置项。

```bash
faultDomainBuildAsPossible bool  # 默认为false，即默认不创建
```

不同faultDomainGrpBatchCnt数量下的节点集分布
- 3 zone（每个 zone 1个nodeset）
- 2 zone（一个 zone 分配2个 nodeset，一个 zone 分配一个 nodeset，以空间大小为权重，剩余空间较大的构建2个 nodeset）
- 1 zone（一个 zone 分配3个 nodeset）

### 原zone（非故障域）的使用空间阈值

升级后，可以对之前卷使用的区域进行扩容，或者按照之前的方式进行操作和维护，或者
可以将故障域的空间使用，但原始空间使用比例需要达到一个阈值，即当前配置项
meta或data使用的整体空间的比例。

默认：0.90，更新api："/admin/updateZoneExcludeRatio"

### 故障域中nodeset组的使用空间阈值

超过后，nodeset组不会用于再分配dp或mp

默认：0.75，更新api："/admin/updateDomainDataRatio"

### 配置小结

1.  现有的cluster，无论是自建的，还是社区的，无论是单个zone，还是跨zone，如果需要故障域启用，需要cluster支持，master重启，配置更新，同时管控更新现有volume的策略。否则继续沿用原有策略。
2.  如果cluster支持，volume不选择使用，则继续原有volume策略，需要在原有zone中按原有策略分配。原有资源耗尽再使用新的zone资源，
3.  如果cluster不支持，volume无法自己启用的故障域策略

| Cluster:faultDomain | Vol:crossZone | Vol:normalZonesFirst | 卷使用故障域的规则                     |
|---------------------|---------------|----------------------|-------------------------------|
| N                   | N/A           | N/A                  | 不支持故障域                        |
| Y                   | N             | N/A                  | 优先选用原有zone，原有zone资源耗尽再使用故障域资源 |
| Y                   | Y             | N                    | 仅写入故障域                        |
| Y                   | Y             | Y                    | 优先选用原有zone，原有zone资源耗尽再使用故障域资源 |

## 注意事项

默认情况下，故障域是为多区域设计的。单区域故障域视为多区域的特例，配置方式一致

故障域解决多zone场景下copysets分布没有规划影响了数据的耐久性的问题，但原有数据不能自动迁移

1.启用故障域后，新区域中的所有设备都将加入故障域

2.创建的volume会优先选择原zone的资源

3.新建卷时需要根据上表添加配置项使用域资源。默认情况下，如果可用，则首先使用原始zone资源

## 管理命令

创建使用故障域的volume

```bash
curl "http://192.168.0.11:17010/admin/createVol?name=volDomain&capacity=1000&owner=cfs&crossZone=true&normalZonesFirst=false"
```

参数列表

| 参数               | 类型     | 描述      |
|------------------|--------|---------|
| crossZone        | string | 是否跨zone |
| normalZonesFirst | bool   | 非故障域优先  |

### 查看故障域是否启用

```bash
curl "http://192.168.0.11:17010/admin/getIsDomainOn"
```

### 查看故障域使用情况

```bash
curl -v  "http://192.168.0.11:17010/admin/getDomainInfo"
```

### 查看故障域copyset group的使用情况

```bash
curl "http://192.168.0.11:17010/admin/getDomainNodeSetGrpInfo?id=37"
```

### 更新非故障域数据使用上限

```bash
curl "http://192.168.0.11:17010/admin/updateZoneExcludeRatio?ratio=0.7"
```