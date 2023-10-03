# Fault Domain Configuration and Management

## What is a fault domain?

A fault domain is a collection of zones. If fault domains are enabled, each replica of the partition will be distributed in different zones to improve data reliability.

When not enabled, all replicas of the partition will be distributed in the same zone.

## Upgrade and Configuration Items

### Cluster-level Configuration

Enabling fault domains requires adding cluster-level configuration, otherwise it cannot be distinguished. Adding a new zone is a fault domain zone or belongs to the original cross_zone.

```bash
FaultDomain               bool  // Default is false
```

### Volume-level Configuration

Reserved:

```bash
crossZone        bool  # Cross-zone
```

New:

```bash
default_priority  bool  # True means prefer the original zone rather than allocating from the fault domain
```

### Zone count to build domain

```bash
faultDomainGrpBatchCnt   int # default: 3，can also set 2 or 1
```

If zone is unavaliable caused by network partition interruption，the following configuration items are used to decide whether to create a nodeset group based on the available zones.

```bash
faultDomainBuildAsPossible bool  # default: false，means don't build nodeset group
```

The distribution of nodesets under the number of different faultDomainGrpBatchCnt
3 zone（1 nodeset per zone）
2 zone（2 nodesets zone，1 nodeset zone，Take the size of the space as the weight, and build 2 nodeset with the larger space remaining）
1 zone（3 nodeset in 1 zone）

### The use space threshold of the non-fault domain（origin zone）

After the upgrade, the zone used by the previous volume can be expanded, or operated and maintained in the previous way, or the space of
the fault domain can be used, but the original space usage ratio needs to reach a threshold, that is, the current configuration item
The proportion of the overall space used by meta or data

Default：0.90, UpdateAPI："/admin/updateZoneExcludeRatio"

### the use space threshold of the nodeset group in the fault domain

Nodeset group will not be used in dp or mp allocation

Default：0.90, UpdateAPI："/admin/updateDomainDataRatio"

### Configuration Summary

1. For existing clusters, whether self-built or community-built, whether single zone or cross-zone, if fault domains need to be enabled, the cluster needs to support it, the master needs to be restarted, the configuration needs to be updated, and the policy for updating existing volumes needs to be managed. Otherwise, continue to use the original policy.
2. If the cluster supports it but the volume does not choose to use it, continue to use the original volume policy and allocate resources according to the original policy in the original zone. Use new zone resources when existing resources are exhausted.
3. If the cluster does not support it, the volume cannot enable its own fault domain policy.

| Cluster:faultDomain | Vol:crossZone | Vol:normalZonesFirst | Rules for volume to use domain                                                |
|---------------------|---------------|----------------------|-------------------------------------------------------------------------------|
| N                   | N/A           | N/A                  | Do not support domain                                                         |
| Y                   | N             | N/A                  | Write origin resources first before fault domain until origin reach threshold |
| Y                   | Y             | N                    | Write fault domain only                                                       |
| Y                   | Y             | Y                    | Write origin resources first before fault domain until origin reach threshold |

## Notes

the fault domain is designed for cross zone by default. The fault domain of a single zone is considered as a special case of cross zone, and the options are consistent.

Fault domains solve the problem of copyset distribution without planning in multi-zone scenarios, which affects the durability of data, but existing data cannot be automatically migrated.

1. After enabling fault domains, all devices in the new region will be added to the fault domain.
2. The volume created will prefer the resources of the original zone.
3. When creating a new volume, use the domain resources according to the configuration items in the table above. By default, if available, the original zone resources are used first.

## Management Commands

Create a volume using fault domains

```bash
curl "http://192.168.0.11:17010/admin/createVol?name=volDomain&capacity=1000&owner=cfs&crossZone=true&normalZonesFirst=false"
```

Parameter List

| Parameter        | Type   | Description               |
|------------------|--------|---------------------------|
| crossZone        | string | Whether to cross zones    |
| normalZonesFirst | bool   | Non-fault domain priority |

### Check if Fault Domain is Enabled

```bash
curl "http://192.168.0.11:17010/admin/getIsDomainOn"
```

### Check Fault Domain Usage

```bash
curl -v  "http://192.168.0.11:17010/admin/getDomainInfo"
```

### Check the usage of fault domain copyset groups

```bash
curl "http://192.168.0.11:17010/admin/getDomainNodeSetGrpInfo?id=37"
```

### Update the upper limit of non-fault domain data usage

```bash
curl "http://192.168.0.11:17010/admin/updateZoneExcludeRatio?ratio=0.7"
```