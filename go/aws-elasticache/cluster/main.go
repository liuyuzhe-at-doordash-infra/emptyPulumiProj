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
	Myimport(string) pulumi.ResourceOption
}

type MyElasticacheInterface interface {
	MyCreateReplicationGroup(*pulumi.Context, *config.Config, *elasticache.ReplicationGroupArgs, []pulumi.ResourceOption) (*elasticache.ReplicationGroup, error)
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

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")

		rgCfg, err := ConfigWithDefaults(ctx)
		if err != nil {
			return errors.Wrap(err, "error building config")
		}

		myPulumiInterface := &MyRealPulumi{}
		myElasticacheInterface := &MyRealElasticache{}

		rg, err := helper(ctx, cfg, rgCfg, myPulumiInterface, myElasticacheInterface)
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

type MyRealPulumi struct{}

// MyRealPulumi struct implmenets MyPulumiInterface interface
func (myRealPulumi *MyRealPulumi) Myimport(s string) pulumi.ResourceOption {
	return pulumi.Import(pulumi.ID(s))
}

type MyRealElasticache struct{}

// MyRealElasticache struct implmenets MyElasticacheInterface interface
func (myRealElasticache *MyRealElasticache) MyCreateReplicationGroup(ctx *pulumi.Context, cfg *config.Config, rgCfg *elasticache.ReplicationGroupArgs, opts []pulumi.ResourceOption) (*elasticache.ReplicationGroup, error) {
	rg, err := elasticache.NewReplicationGroup(ctx, cfg.Require(ResourceNameKey), rgCfg, opts...)
	return rg, err
}

func playFunc(ctx *pulumi.Context) error {

	cfg := config.New(ctx, "")

	//rgCfg, err := ConfigWithDefaults(ctx)
	rgCfg, err := ConfigWithImport(ctx)
	if err != nil {
		return errors.Wrap(err, "error building config")
	}

	myPulumiInterface := &MyRealPulumi{}
	myElasticacheInterface := &MyRealElasticache{}

	//opts := make([]pulumi.ResourceOption, 0)
	//
	//isImport := cfg.GetBool(ResourceImportKey)
	//
	//if isImport {
	//	//opts = append(opts, pulumi.Import(pulumi.ID(cfg.Require(ResourceImportIdKey))))
	//	opts = append(opts, myPulumiInterface.Myimport("myID"))
	//}
	//
	////rg, err := elasticache.NewReplicationGroup(ctx, cfg.Require(ResourceNameKey), rgCfg, opts...)
	//rg, err := myElasticacheInterface.MyCreateReplicationGroup("name", ctx, cfg)

	rg, err := helper(ctx, cfg, rgCfg, myPulumiInterface, myElasticacheInterface)

	if err != nil {
		return errors.Wrap(err, "error creating or importing replication group")
	}

	ctx.Export("configuration_endpoint_address", rg.ConfigurationEndpointAddress)
	ctx.Export("engine_actual_version", rg.EngineVersionActual)

	// Export current timestamp to force stack update
	ctx.Export("updatedAt", pulumi.StringPtr(fmt.Sprintf("%d", time.Now().Unix())))

	return nil
}

func helper(ctx *pulumi.Context,
	cfg *config.Config,
	rgCfg *elasticache.ReplicationGroupArgs,
	myPulumiInterface MyPulumiInterface,
	myElasticacheInterface MyElasticacheInterface) (*elasticache.ReplicationGroup, error) {

	opts := make([]pulumi.ResourceOption, 0)
	isImport := cfg.GetBool(ResourceImportKey)

	if isImport {
		//opts = append(opts, pulumi.Import(pulumi.ID(cfg.Require(ResourceImportIdKey))))
		opts = append(opts, myPulumiInterface.Myimport("myID"))
	}

	rg, err := myElasticacheInterface.MyCreateReplicationGroup(ctx, cfg, rgCfg, opts)

	return rg, err
}
