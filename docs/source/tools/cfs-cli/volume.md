# Volume Management

## Create Volume

```bash
cfs-cli volume create [VOLUME NAME] [USER ID] [flags]
```

| Flag                           | Description                                                                                                                                   |
|--------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| cache-action                   | Specify low volume cacheAction (default 0)                                                                                                    |
| cache-capacity                 | Specify low volume capacity[Unit: GB]                                                                                                         |
| cache-high-water               | (default 80)                                                                                                                                  |
| cache-low-water                | (default 60)                                                                                                                                  |
| cache-lru-interval             | Specify interval expiration time[Unit: min] (default 5)                                                                                       |
| cache-rule-key                 | Anything that match this field will be written to the cache                                                                                   |
| cache-threshold                | Specify cache threshold[Unit: byte] (default 10485760)                                                                                        |
| cache-ttl                      | Specify cache expiration time[Unit: day] (default 30)                                                                                         |
| capacity                       | Specify volume capacity (default 10)                                                                                                          |
| clientIDKey                    | needed if cluster authentication is on                                                                                                        |
| crossZone                      | Disable cross zone (default "false")                                                                                                          |
| delete-lock-time               | Specify delete lock time[Unit: hour] for volume                                                                                               |
| description                    | Description                                                                                                                                   |
| ebs-blk-size                   | Specify ebsBlk Size[Unit: byte] (default 8388608)                                                                                             |
| enableQuota                    | Enable quota (default false) (default "false")                                                                                                |
| follower-read                  | Enable read form replica follower                                                                                                             |
| mp-count                       | Specify init meta partition count (default 3)                                                                                                 |
| normalZonesFirst               | Write to normal zone first (default "false")                                                                                                  |
| readonly-when-full             | Enable volume becomes read only when it is full (default "false")                                                                             |
| replica-num                    | Specify data partition replicas number(default 3 for normal volume,1 for low volume)                                                          |
| size                           | Specify data partition size[Unit: GB] (default 120)                                                                                           |
| transaction-mask               | Enable transaction for specified operation: ["create&#124;mkdir&#124;remove&#124;rename&#124;mknod&#124;symlink&#124;link"] or "off" or "all" |
| transaction-timeout            | Specify timeout[Unit: minute] for transaction [1-60] (default 1)                                                                              |
| tx-conflict-retry-Interval     | Specify retry interval[Unit: ms] for transaction conflict [10-1000]                                                                           |
| tx-conflict-retry-num          | Specify retry times for transaction conflict [1-100]                                                                                          |
| vol-type                       | Type of volume (default 0)                                                                                                                    |
| y, yes                         | Answer yes for all questions                                                                                                                  |
| zone-name                      | Specify volume zone name                                                                                                                      |

```

## Delete Volume

Delete the specified volume [VOLUME NAME]. The size of the ec volume must be 0 to be deleted.

```bash
cfs-cli volume delete [VOLUME NAME] [flags]
```

| Name   | Description                                     |
|--------|-------------------------------------------------|
| y, yes | Skip all questions and set the answer to "yes". |

## Show Volume

Get information of the volume [VOLUME NAME].

```bash
cfs-cli volume info [VOLUME NAME] [flags]
```

### Flags

| Name              | Description                                      |
|-------------------|--------------------------------------------------|
| d, data-partition | Show detailed information of the data Partition. |
| m, meta-partition | Show detailed information of the Meta Partition. |

## Create and Add Data Partitions to the Volume

Create and add [NUMBER] data partitions to the volume [VOLUME].

```bash
cfs-cli volume add-dp [VOLUME] [NUMBER OF DATA PARTITIONS]
```

## List Volumes

```bash
cfs-cli volume list
```

## Transfer Volume

Transfer the volume [VOLUME NAME] to another user [USER ID].

```bash
cfs-cli volume transfer [VOLUME NAME] [USER ID] [flags]   
```

### Flags

| Name     | Description                                     |
|----------|-------------------------------------------------|
| f, force | Force transfer.                                 |
| y, yes   | Skip all questions and set the answer to "yes". |

## Volume Configuration Setup

Update the configurations of the volume.

```bash
cfs-cli volume update
```

### Flags

| Name                       | Description                                                                                                                                              |
|----------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|
| cache-action               | Specify low volume cacheAction (default 0)                                                                                                               |
| cache-capacity             | Specify low volume capacity[Unit: GB]                                                                                                                    |
| cache-high-water           | (default 80)                                                                                                                                             |
| cache-low-water            | (default 60)                                                                                                                                             |
| cache-lru-interval         | Specify interval expiration time[Unit: min] (default 5)                                                                                                  |
| cache-rule                 | Specify cache rule                                                                                                                                       |
| cache-threshold            | Specify cache threshold[Unit: byte] (default 10M)                                                                                                        |
| cache-ttl                  | Specify cache expiration time[Unit: day] (default 30)                                                                                                    |
| capacity                   | Specify volume datanode capacity [Unit: GB]                                                                                                              |
| clientIDKey                | needed if cluster authentication is on                                                                                                                   |
| delete-lock-time           | Specify delete lock time[Unit: hour] for volume (default -1)                                                                                             |
| description                | The description of volume                                                                                                                                |
| ebs-blk-size               | Specify ebsBlk Size[Unit: byte]                                                                                                                          |
| enableQuota                | Enable quota                                                                                                                                             |
| follower-read              | Enable read form replica follower (default false)                                                                                                        |
| readonly-when-full         | Enable volume becomes read only when it is full                                                                                                          |
| replica-num                | Specify data partition replicas number(default 3 for normal volume,1 for low volume)                                                                     |
| transaction-force-reset    | Reset transaction mask to the specified value of "transaction-mask"                                                                                      |
| transaction-limit          | Specify limitation[Unit: second] for transaction(default 0 unlimited)                                                                                    |
| transaction-mask           | Enable transaction for specified operation: ["create&#124;mkdir&#124;remove&#124;rename&#124;mknod&#124;symlink&#124;link&#124;pause"] or "off" or "all" |
| transaction-timeout        | Specify timeout[Unit: minute] for transaction (0-60]                                                                                                     |
| tx-conflict-retry-Interval | Specify retry interval[Unit: ms] for transaction conflict [10-1000]                                                                                      |
| tx-conflict-retry-num      | Specify retry times for transaction conflict [1-100]                                                                                                     |
| y, yes                     | Answer yes for all questions                                                                                                                             |
| zone-name                  | Specify volume zone name                                                                                                                                 |
