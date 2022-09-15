package main

import (
	"fmt"
	"github.com/pkg/errors"
	"time"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")

		//isImport := cfg.GetBool("is_import")
		isImport := true // force to test
		if isImport {
			println("### Importing ElastiCache ###")

			// import elasticache
			importRgCfg, importErr := ConfigWithImport(ctx)
			if importErr != nil {
				return errors.Wrap(importErr, "error building config from imports")
			}

			importRg, importErr := elasticache.NewReplicationGroup(
				ctx,
				"imported-pulumi-stack-my-test-222-elasticache-redis-simple",
				importRgCfg,
				pulumi.Import(pulumi.ID("my-test-222-elasticache-redis-simple")),
			)
			if importErr != nil {
				return errors.Wrap(importErr, "error importing the existing replication group")
			}

			println("importRg.ConfigurationEndpointAddress: %v", importRg.ConfigurationEndpointAddress)
			println("importRg.EngineVersionActual: %v", importRg.EngineVersionActual)
		} else {
			println("### Creating ElastiCache ###")

			// create elasticache
			rgCfg, err := ConfigWithDefaults(ctx)
			if err != nil {
				return errors.Wrap(err, "error building default config")
			}

			rg, err := elasticache.NewReplicationGroup(ctx, cfg.Require("name"), rgCfg)
			if err != nil {
				return errors.Wrap(err, "error creating replication group")
			}

			ctx.Export("configuration_endpoint_address", rg.ConfigurationEndpointAddress)
			ctx.Export("engine_actual_version", rg.EngineVersionActual)

			// Export current timestamp to force stack update
			ctx.Export("updatedAt", pulumi.StringPtr(fmt.Sprintf("%d", time.Now().Unix())))
		}

		return nil
	})
}
