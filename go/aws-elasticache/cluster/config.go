package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// type Inputs struct {
// 	ApplyImmediately       string `json:"apply_immediately"`
// 	Description            string `json:"description"` // required
// 	EngineVersion          string `json:"engine_version"`
// 	Environment            string `json:"environment"` // required, tag
// 	MaintenanceWindow      string `json:"maintenance_window"`
// 	Name                   string `json:"name"`            // required
// 	NodeType               string `json:"node_type"`       // required
// 	NumNodeGroups          string `json:"num_node_groups"` // required
// 	ParameterGroupName     string `json:"parameter_group_name"`
// 	ReplicasPerNodeGroup   string `json:"replicas_per_node_group"` // required
// 	SecurityGroupId        string `json:"security_group_id"`       // required
// 	Service                string `json:"service"`                 // required, tag
// 	SnapshotName           string `json:"snapshot_name"`
// 	SnapshotRetentionLimit string `json:"snapshot_retention_limit"` // required
// 	SnapshotWindow         string `json:"snapshot_window"`
// 	SubnetGroupName        string `json:"subnet_group_name"` // required
// }

const (
	automaticFailoverEnabled = true
	engineVersion            = "5.0.6"
	maintenanceWindow        = "wed:11:00-wed:12:00"
	multiAzEnabled           = true
	parameterGroupName       = "dd-redis-5-lfu-cluster-on"
	snapshotWindow           = "12:00-13:00"

	pulumiManagedTag = "pulumi-managed"
)

// test using enforced config before we have context passed in
func ConfigWithImport(ctx *pulumi.Context) (rgConfig *elasticache.ReplicationGroupArgs, err error) {
	//cfg := config.New(ctx, "")
	//rgConfig = defaultReplicationGroupConfig()

	rgConfig = &elasticache.ReplicationGroupArgs{}

	// TODO: see if we still need to set this when using DD's elasticache module!!!!
	rgConfig.AtRestEncryptionEnabled = pulumi.Bool(false)

	// TODO: see if we still need to set this when using DD's elasticache module!!!!
	rgConfig.AutoMinorVersionUpgrade = pulumi.Bool(false)

	rgConfig.AutomaticFailoverEnabled = pulumi.Bool(automaticFailoverEnabled)
	//rgConfig.AutomaticFailoverEnabled = pulumi.Bool(true)

	// Required
	rgConfig.Description = pulumi.String("manually created dummy elasticache replication group for testing.")

	rgConfig.NodeType = pulumi.String("cache.t4g.micro")

	rgConfig.NumNodeGroups = pulumi.Int(1)

	// TODO: see if we still need to set this when using DD's elasticache module!!!!
	rgConfig.Port = pulumi.Int(6379)

	rgConfig.ReplicasPerNodeGroup = pulumi.Int(1)

	//rgConfig.ReplicationGroupId = pulumi.String("my-test-222-elasticache-redis-simple")
	rgConfig.ReplicationGroupId = pulumi.String("test-manual-elasticache-rp-1")

	rgConfig.SecurityGroupIds = pulumi.StringArray{
		pulumi.String("sg-0b39791f9f42c6de8"),
	}

	rgConfig.SnapshotRetentionLimit = pulumi.Int(0)

	//rgConfig.SubnetGroupName = pulumi.String("my-test-222-elasticache-redis-simple")
	rgConfig.SubnetGroupName = pulumi.String("test-elasticache-subnetgroup")

	rgConfig.Tags = pulumi.StringMap{
		//"Name": pulumi.String("my-test-222-elasticache-redis-simple"),
		//"Name": pulumi.String("test-manual-elasticache-rp-1"),
	}

	// TODO: see if we still need to set this when using DD's elasticache module!!!!
	rgConfig.TransitEncryptionEnabled = pulumi.Bool(false)

	// TODO: see if we still need to set this when using DD's elasticache module!!!!
	rgConfig.MultiAzEnabled = pulumi.Bool(multiAzEnabled)
	//rgConfig.MultiAzEnabled = pulumi.Bool(true)

	// Optional

	rgConfig.EngineVersion = pulumi.String(engineVersion)
	//rgConfig.EngineVersion = pulumi.String("5.0.6")

	rgConfig.MaintenanceWindow = pulumi.String(maintenanceWindow)
	//rgConfig.MaintenanceWindow = pulumi.String("tue:23:30-wed:00:30")

	//rgConfig.ParameterGroupName = pulumi.String("my-test-222-elasticache-redis-simple")
	rgConfig.ParameterGroupName = pulumi.String("default.redis5.0")

	//rgConfig.SnapshotWindow = pulumi.String(snapshotWindow)
	rgConfig.SnapshotWindow = pulumi.String("05:30-06:30")

	//// ApplyImmediately
	//if cfg.GetBool("apply_immediately") {
	//	rgConfig.ApplyImmediately = pulumi.Bool(true)
	//}
	//
	//// SnapshotName
	//if snapshotName := cfg.Get("snapshot_name"); snapshotName != "" {
	//	rgConfig.SnapshotName = pulumi.String(snapshotName)
	//}

	return
}

func ConfigWithDefaults(ctx *pulumi.Context) (rgConfig *elasticache.ReplicationGroupArgs, err error) {
	cfg := config.New(ctx, "")
	rgConfig = defaultReplicationGroupConfig()

	// Required

	// Description
	rgConfig.Description = pulumi.String(cfg.Require("description"))

	// ReplicationGroupId
	rgConfig.ReplicationGroupId = pulumi.String(cfg.Require("name"))

	// NodeType
	rgConfig.NodeType = pulumi.String(cfg.Require("node_type"))

	// NumNodeGroups
	rgConfig.NumNodeGroups = pulumi.Int(cfg.RequireInt("num_node_groups"))

	// ReplicasPerNodeGroup
	rgConfig.ReplicasPerNodeGroup = pulumi.Int(cfg.RequireInt("replicas_per_node_group"))

	// SecurityGroupId
	rgConfig.SecurityGroupIds = pulumi.ToStringArray([]string{cfg.Require("security_group_id")})

	// SnapshotRetentionLimit
	rgConfig.SnapshotRetentionLimit = pulumi.Int(cfg.RequireInt("snapshot_retention_limit"))

	// SubnetGroupName
	rgConfig.SubnetGroupName = pulumi.String(cfg.Require("subnet_group_name"))

	// Tags
	rgConfig.Tags = pulumi.ToStringMap(map[string]string{
		"environment":    cfg.Require("environment"),
		"service":        cfg.Require("service"),
		pulumiManagedTag: "true",
	})

	// Optional

	// ApplyImmediately
	if cfg.GetBool("apply_immediately") {
		rgConfig.ApplyImmediately = pulumi.Bool(true)
	}

	// EngineVersion
	if engineVersion := cfg.Get("engine_version"); engineVersion != "" {
		rgConfig.EngineVersion = pulumi.String(engineVersion)
	}

	// MaintenanceWindow
	if maintenanceWindow := cfg.Get("maintenance_window"); maintenanceWindow != "" {
		rgConfig.MaintenanceWindow = pulumi.String(maintenanceWindow)
	}

	// ParameterGroupName
	if parameterGroupName := cfg.Get("parameter_group_name"); parameterGroupName != "" {
		rgConfig.ParameterGroupName = pulumi.String(parameterGroupName)
	}

	// SnapshotName
	if snapshotName := cfg.Get("snapshot_name"); snapshotName != "" {
		rgConfig.SnapshotName = pulumi.String(snapshotName)
	}

	// SnapshotWindow
	if snapshotWindow := cfg.Get("snapshot_window"); snapshotWindow != "" {
		rgConfig.SnapshotWindow = pulumi.String(snapshotWindow)
	}

	return
}

func defaultReplicationGroupConfig() *elasticache.ReplicationGroupArgs {
	return &elasticache.ReplicationGroupArgs{
		AutomaticFailoverEnabled: pulumi.Bool(automaticFailoverEnabled),
		EngineVersion:            pulumi.String(engineVersion),
		MaintenanceWindow:        pulumi.String(maintenanceWindow),
		MultiAzEnabled:           pulumi.Bool(multiAzEnabled),
		ParameterGroupName:       pulumi.String(parameterGroupName),
		SnapshotWindow:           pulumi.String(snapshotWindow),
	}
}
