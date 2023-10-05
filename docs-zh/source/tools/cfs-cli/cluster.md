# 集群管理

## 获取集群信息

包括集群名称、地址、卷数量、节点数量及使用率等

```bash
cfs-cli cluster info
```

## 获取集群状态

按区域获取元数据和数据节点的使用量、状态等

```bash
cfs-cli cluster stat
```

## 冻结/解冻集群

设置为 `true` 冻结后，当partition写满，集群不会自动分配新的partition

```bash
cfs-cli cluster freeze [true/false]
```

## 设置内存阈值

设置集群中每个MetaNode的内存阈值。当内存使用率超过该阈值时，上面的meta partition将会被设为只读。[float]应当是一个介于0和1之间的小数.

```bash
cfs-cli cluster threshold [float]
```

## 设置集群参数

```bash
cfs-cli cluster set [flags]
```

### Flags

| 名称                  | 描述                               |
|---------------------|----------------------------------|
| autoRepairRate      | 数据节点的自动修复速率                      |
| batchCount          | 元数据节点的批量删除计数                     |
| clientIDKey         | 如果集群认证打开则需要此项                    |
| dataNodeSelector    | 设置集群的数据节点选择策略                    |
| dataNodesetSelector | 设置集群的数据节点集选择策略                   |
| deleteWorkerSleepMs | 元数据节点删除工作线程的休眠时间（毫秒）。如果为0则不休眠    |
| loadFactor          | 负载因子                             |
| markDeleteRate      | 数据节点的批量标记删除速率。如果为0则没有限制          |
| maxDpCntLimit       | 每个数据节点上的最大dp数量，默认为3000，0表示设置为默认值 |
| metaNodeSelector    | 设置集群的元数据节点选择策略                   |
| metaNodesetSelector | 设置集群的元数据节点集选择策略                  |
