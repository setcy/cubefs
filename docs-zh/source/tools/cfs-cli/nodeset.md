# nodeset管理

## 列出所有nodeset

```bash
cfs-cli nodeset list
```

### Flags

| 名称               | 描述              |
|------------------|-----------------|
| zone-name        | 仅显示指定区域的nodeset |

## 获取单个nodeset的信息

```bash
cfs-cli nodeset info [NODESET ID]
```

## 更新某个nodeset的信息

```bash
cfs-cli nodeset update [NODESET ID] [flags]
```

### Flags

| 名称               | 描述                  |
|------------------|---------------------|
| dataNodeSelector | 为指定的节点集设置数据节点的选择策略  |
| metaNodeSelector | 为指定的节点集设置元数据节点的选择策略 |
