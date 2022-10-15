package main

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

const (
	ResourceImportKey   = "is_import"
	ResourceImportIdKey = "resource_import_id"
	ResourceNameKey     = "name"
)

type MyPulumiInterface interface {
	Myimport() pulumi.ResourceOption
}

type MyElasticacheInterface interface {
	MyCreateReplicationGroup([]pulumi.ResourceOption) (*elasticache.ReplicationGroup, error)
}

type MyRealPulumi struct {
	ID string
}

// MyRealPulumi struct implmenets MyPulumiInterface interface
func (myRealPulumi *MyRealPulumi) Myimport() pulumi.ResourceOption {
	return pulumi.Import(pulumi.ID(myRealPulumi.ID))
}

type MyRealElasticache struct {
	ctx *pulumi.Context
	//cfg   *config.Config  // TODO: re-enable
	rgCfg *elasticache.ReplicationGroupArgs
	//opts  []pulumi.ResourceOption // TODO: can be delete I think
}

// MyRealElasticache struct implmenets MyElasticacheInterface interface
func (myRealElasticache *MyRealElasticache) MyCreateReplicationGroup(opts []pulumi.ResourceOption) (*elasticache.ReplicationGroup, error) {
	rg, err := elasticache.NewReplicationGroup(
		myRealElasticache.ctx,
		//myRealElasticache.cfg.Require(ResourceNameKey),  // TODO: re-enable
		//"imported-pulumi-stack-my-test-222-elasticache-redis-simple", // TODO: for test only
		"imported-test-1-pulumi", // TODO: for test only
		myRealElasticache.rgCfg,
		opts...)
	return rg, err
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")

		//rgCfg, err := ConfigWithDefaults(ctx)
		rgCfg, err := ConfigWithImport(ctx)
		if err != nil {
			return errors.Wrap(err, "error building config")
		}

		// define instance of two structs that implement interfaces
		//myPulumiStruct := &MyRealPulumi{ID: "my-test-222-elasticache-redis-simple"}
		myPulumiStruct := &MyRealPulumi{ID: "test-manual-elasticache-rp-1"} // pre-existing elasticache resource name
		myElasticacheStruct := &MyRealElasticache{
			ctx:   ctx,
			rgCfg: rgCfg,
		}

		rg, err := helper(ctx, cfg, rgCfg, myPulumiStruct, myElasticacheStruct)
		if err != nil {
			return errors.Wrap(err, "error creating or importing replication group")
		}

		ctx.Export("configuration_endpoint_address", rg.ConfigurationEndpointAddress)
		ctx.Export("engine_actual_version", rg.EngineVersionActual)

		// Export current timestamp to force stack update
		ctx.Export("updatedAt", pulumi.StringPtr(fmt.Sprintf("%d", time.Now().Unix())))

		return nil
	})
}

func helper(ctx *pulumi.Context,
	cfg *config.Config,
	rgCfg *elasticache.ReplicationGroupArgs, // TODO: can be removed I think
	myPulumiInterface MyPulumiInterface,
	myElasticacheInterface MyElasticacheInterface) (*elasticache.ReplicationGroup, error) {

	opts := make([]pulumi.ResourceOption, 0)
	//isImport := cfg.GetBool(ResourceImportKey) // TODO: re-enable

	// ##############################################
	//isImport := true // TODO: for test only
	isImport := false // TODO: for test only
	// ##############################################

	if isImport {
		//opts = append(opts, pulumi.Import(pulumi.ID(cfg.Require(ResourceImportIdKey))))
		opts = append(opts, myPulumiInterface.Myimport())
	}

	rg, err := myElasticacheInterface.MyCreateReplicationGroup(opts)

	return rg, err
}

//func main() {
//	pulumi.Run(func(ctx *pulumi.Context) error {
//		cfg := config.New(ctx, "")
//
//		rgCfg, err := ConfigWithDefaults(ctx)
//		if err != nil {
//			return errors.Wrap(err, "error building config")
//		}
//
//		opts := make([]pulumi.ResourceOption, 0)
//
//		isImport := cfg.GetBool(ResourceImportKey)
//		if isImport {
//			opts = append(opts, pulumi.Import(pulumi.ID(cfg.Require(ResourceImportIdKey))))
//		}
//
//		rg, err := elasticache.NewReplicationGroup(ctx, cfg.Require(ResourceNameKey), rgCfg, opts...)
//		if err != nil {
//			return errors.Wrap(err, "error creating or importing replication group")
//		}
//
//		ctx.Export("configuration_endpoint_address", rg.ConfigurationEndpointAddress)
//		ctx.Export("engine_actual_version", rg.EngineVersionActual)
//
//		// Export current timestamp to force stack update
//		ctx.Export("updatedAt", pulumi.StringPtr(fmt.Sprintf("%d", time.Now().Unix())))
//
//		return nil
//	})
//}

//func playFunc(ctx *pulumi.Context) error {
//
//	cfg := config.New(ctx, "")
//
//	rgCfg, err := ConfigWithDefaults(ctx)
//	if err != nil {
//		return errors.Wrap(err, "error building config")
//	}
//
//	myPulumiInterface := &MyRealPulumi{}
//	myElasticacheInterface := &MyRealElasticache{}
//
//	//opts := make([]pulumi.ResourceOption, 0)
//	//
//	//isImport := cfg.GetBool(ResourceImportKey)
//	//
//	//if isImport {
//	//	//opts = append(opts, pulumi.Import(pulumi.ID(cfg.Require(ResourceImportIdKey))))
//	//	opts = append(opts, myPulumiInterface.Myimport("myID"))
//	//}
//	//
//	////rg, err := elasticache.NewReplicationGroup(ctx, cfg.Require(ResourceNameKey), rgCfg, opts...)
//	//rg, err := myElasticacheInterface.MyCreateReplicationGroup("name", ctx, cfg)
//
//	rg, err := helper(ctx, cfg, rgCfg, myPulumiInterface, myElasticacheInterface)
//
//	if err != nil {
//		return errors.Wrap(err, "error creating or importing replication group")
//	}
//
//	ctx.Export("configuration_endpoint_address", rg.ConfigurationEndpointAddress)
//	ctx.Export("engine_actual_version", rg.EngineVersionActual)
//
//	// Export current timestamp to force stack update
//	ctx.Export("updatedAt", pulumi.StringPtr(fmt.Sprintf("%d", time.Now().Unix())))
//
//	return nil
//}
