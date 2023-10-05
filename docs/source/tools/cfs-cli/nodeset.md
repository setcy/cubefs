# Nodeset Management

## List NodeSets

```bash
cfs-cli nodeset list
```

### Flags

| 名称               | 描述                                             |
|------------------|------------------------------------------------|
| zone-name        | Only display the nodeset of the specified area |

## Show NodeSet

```bash
cfs-cli nodeset info [NODESET ID]
```

## Update NodeSet

```bash
cfs-cli nodeset update [NODESET ID] [flags]
```

### Flags

| Name             | Description                                                           |
|------------------|-----------------------------------------------------------------------|
| dataNodeSelector | Set the node select policy(datanode) for specify nodeset              |
| metaNodeSelector | Set the node select policy(metanode) for specify nodeset              |
