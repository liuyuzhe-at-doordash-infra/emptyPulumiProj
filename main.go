package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	/* note: **NOT** use cmd import.
	pulumi import aws:elasticache/replicationGroup:ReplicationGroup imported-pulumi-stack-my-test-222-elasticache-redis-simple my-test-222-elasticache-redis-simple
	*/
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewReplicationGroup(ctx, "imported-pulumi-stack-my-test-222-elasticache-redis-simple", &elasticache.ReplicationGroupArgs{
			AtRestEncryptionEnabled:  pulumi.Bool(true),
			AutoMinorVersionUpgrade:  pulumi.Bool(true),
			AutomaticFailoverEnabled: pulumi.Bool(true),
			//ClusterMode: &elasticache.ReplicationGroupClusterModeArgs{
			//	NumNodeGroups:        pulumi.Int(1),
			//	ReplicasPerNodeGroup: pulumi.Int(1),
			//},
			Description:       pulumi.String("Originally Managed by Terraform"),
			EngineVersion:     pulumi.String("5.0.6"),
			MaintenanceWindow: pulumi.String("tue:23:30-wed:00:30"),
			NodeType:          pulumi.String("cache.r6g.large"),
			NumCacheClusters:  pulumi.Int(2),
			//NumNodeGroups:     pulumi.Int(1),
			//NumberCacheClusters:  pulumi.Int(2),
			ParameterGroupName:   pulumi.String("my-test-222-elasticache-redis-simple"),
			Port:                 pulumi.Int(6379),
			ReplicasPerNodeGroup: pulumi.Int(1),
			//ReplicationGroupDescription: pulumi.String("Managed by Terraform"),
			ReplicationGroupId: pulumi.String("my-test-222-elasticache-redis-simple"),
			SecurityGroupIds: pulumi.StringArray{
				pulumi.String("sg-0b39791f9f42c6de8"),
			},
			SnapshotRetentionLimit: pulumi.Int(30),
			SnapshotWindow:         pulumi.String("06:00-07:00"),
			SubnetGroupName:        pulumi.String("my-test-222-elasticache-redis-simple"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("my-test-222-elasticache-redis-simple"),
			},
			TransitEncryptionEnabled: pulumi.Bool(true),
			MultiAzEnabled:           pulumi.Bool(true),
		}, pulumi.Import(pulumi.ID("my-test-222-elasticache-redis-simple")),
		)
		//pulumi.Protect(true))

		if err != nil {
			return err
		}
		return nil
	})
}
