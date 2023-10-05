# 配额管理

::: warning 注意
目录配额管理为v3.3.0版本新增feature
:::

## 创建配额

创建quota需要指定卷名一个或多个path目录。

注意：path之间不能重复，以及嵌套。

```bash
cfs-cli quota create [volname] [fullpath1,fullpath2] [flags]
```

### Flags

| 名称         | 描述                                 |
|------------|------------------------------------|
| maxBytes   | 指定配额最大字节数（默认 18446744073709551615） |
| maxFiles   | 指定配额最大文件数（默认 18446744073709551615） |


## 应用配额

apply quota需要指定卷名以及quotaId，这个接口在创建quota后执行，目的是让quota目录下（包括quota目录自身）的已有文件和目录该quotaId生效。整个创建quota的流程先执行quota
create，然后执行quota apply命令。

注意：如果quota目录下的文件数量很多，则该接口返回时间会比较长

```bash
cfs-cli quota apply [volname] [quotaId] [flags]
```

### Flags

| 名称                  | 描述                          |
|---------------------|-----------------------------|
| maxConcurrencyInode | 同时设置 inodes 的最大并发数（默认 1000） |


## 取消应用配额

revoke quota需要指定卷名以及quotaId，这个接口在准备删除quota的时候执行，目的是让quota目录下的（包括quota目录自身）的已有文件和目录该quotaId失效。整个删除quota的流程先执行quota
revoke，然后通过quota list查询确认USEDFILES和USEDBYTES的值为0，再进行quota delete操作。

```bash
cfs-cli quota revoke [volname] [quotaId] [flags]
```

### Flags

| 名称                  | 描述                        |
|---------------------|---------------------------|
| forceInode          | 强制撤销配额 inode              |
| maxConcurrencyInode | 同时删除 inode 的最大并发数（默认1000） |


## 删除配额

delete quota需要指定卷名以及quotaId

```bash
cfs-cli quota delete [volname] [quotaId] [flags]
```
### Flags

| 名称     | 描述                |
|--------|-------------------|
| y, yes | 跳过所有问题并设置回答为"yes" |

## 更新配额

update quota需要指定卷名以及quotaId，目前可以更新的值只有maxBytes和maxFiles

```bash
cfs-cli quota update [volname] [quotaId] [flags]
```

### Flags

| 名称       | 描述        |
|----------|-----------|
| maxBytes | 指定配额最大字节数 |
| maxFiles | 指定配额最大文件数 |


## 列出卷配额信息

list quota需要指定卷名，遍历出所有该卷的quota信息

``` bash
cfs-cli quota list [volname] [flags]
```

## 列出所有卷的配额信息

不带任何参数，遍历出所有带quota的卷信息

```bash
cfs-cli quota listAll [flags]
```

## 查看某个inode的配额信息

查看具体的某个inode是否带有quota信息

``` bash
cfs-cli quota getInode [volname] [inode] [flags]
```