# AWS Elasticache Cluster

This Pulumi program provisions an Elasticache Redis cluster.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| apply\_immediately | Whether any database modifications are applied immediately, or during the next maintenance window | `string` | `"false"` | no |
| description | Description of this cache and how it's used | `string` | N/A | yes |
| engine\_version | Engine version with which to deploy the cluster. See https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/supported-engine-versions.html | `string` | `"5.0.6"` | no |
| environment | Environment to which this cache will be deployed | `string` | N/A | yes |
| maintenance\_window | Weekly time range when maintenance is performed. Format is ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). Minimum maintenance window is a 60 minute period. Example: sun:05:00-sun:09:00 | `string` | `"wed:11:00-wed:12:00"` | no |
| name | The name to give the cache | `string` | N/A | yes |
| node\_type | Cache instance type to use for this cluster | `string` | N/A | yes |
| num\_node\_groups | The number of node groups (aka shards) for this cache | `string` | N/A | yes |
| parameter\_group\_name | The parameter group to use for this cluster | `string` | `"dd-redis-5-lfu-cluster-on"` | no |
| replicas\_per\_node\_group | The number of replicas in each node group (shard). This should be "0" if you don't need replicas, and typically "1" if you do | `string` | N/A | yes |
| security\_group\_id | The security group id this cache is deployed into | `string` | N/A | yes |
| service | The name of the service that owns the cache | `string` | N/A | yes |
| snapshot\_name | The name of a snapshot from which to restore data into the new node group. Adding/Changing/Removing the snapshot_name has no effect on existing cluster resource | `string` | N/A | no |
| snapshot\_retention\_limit | Number of days to retain snapshots before deleting. If set to "0", snapshots are turned off | `string` | N/A | yes |
| snapshot\_window | Daily time range (UTC) when ElastiCache will begin taking a snapshot. Minimum window is a 60 minute period. Example: "05:00-09:00" | `string` | `"12:00-13:00"` | no |
| subnet\_group\_name | Subnet group name for this cache | `string` | N/A | yes |
