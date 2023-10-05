# 访问控制列表管理 (ACL)

## 添加IP到黑名单

将指定的IP添加到卷的ACL黑名单中，防止该IP访问或操作卷。

```bash
cfs-cli acl add [volName] [ipAddress]
```

## 检查IP是否在黑名单中

查询指定的IP是否已被添加到卷的ACL黑名单中。

```bash
cfs-cli acl check [volName] [ipAddress]
```

## 从黑名单中删除IP

将指定的IP从卷的ACL黑名单中删除，允许该IP再次访问或操作卷。

```bash
cfs-cli acl del [volName] [ipAddress]
```

## 列出所有在黑名单中的IP

查看卷的ACL黑名单中的所有IP地址。

```bash
cfs-cli acl list [volName]
```
