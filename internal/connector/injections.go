package connector

import (
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/common"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/connector/internal/config"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/connector/internal/handlers"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/connector/internal/routes"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/connector/internal/services"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/connector/internal/workers"
	coreWorkers "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/workers"

	coreConfig "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/config"

	"github.com/goava/di"
)

func EnvInjections() common.InjectionMap {
	return common.InjectionMap{
		"ConfigModule":    di.Provide(config.NewConnectorsConfig, di.As(new(coreConfig.ConfigModule))),
		"ServiceInjector": di.Provide(newServiceInjector),
	}
}

func newServiceInjector(container *di.Container) coreConfig.ServiceInjector {
	return serviceInjector{parent: container}
}

type serviceInjector struct {
	parent *di.Container
}

func (s serviceInjector) Injections() (common.InjectionMap, error) {
	connectorsConfig := &config.ConnectorsConfig{}
	if err := s.parent.Resolve(&connectorsConfig); err != nil {
		return nil, err
	}

	return common.InjectionMap{
		"Config":                  di.ProvideValue(connectorsConfig),
		"ConnectorsService":       di.Provide(services.NewConnectorsService, di.As(new(services.ConnectorsService))),
		"ConnectorTypesService":   di.Provide(services.NewConnectorTypesService, di.As(new(services.ConnectorTypesService))),
		"ConnectorClusterService": di.Provide(services.NewConnectorClusterService, di.As(new(services.ConnectorClusterService))),
		"ConnectorTypesHandler":   di.Provide(handlers.NewConnectorTypesHandler),
		"ConnectorsHandler":       di.Provide(handlers.NewConnectorsHandler),
		"ConnectorClusterHandler": di.Provide(handlers.NewConnectorClusterHandler),
		"RouteLoader":             di.Provide(routes.NewRouteLoader),
		"ConnectorManager":        di.Provide(workers.NewConnectorManager, di.As(new(coreWorkers.Worker))),
	}, nil
}