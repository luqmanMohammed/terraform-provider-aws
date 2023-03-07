// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_ecs_cluster":              DataSourceCluster,
		"aws_ecs_container_definition": DataSourceContainerDefinition,
		"aws_ecs_service":              DataSourceService,
		"aws_ecs_task_definition":      DataSourceTaskDefinition,
		"aws_ecs_task_execution":       DataSourceTaskExecution,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_ecs_account_setting_default":    ResourceAccountSettingDefault,
		"aws_ecs_capacity_provider":          ResourceCapacityProvider,
		"aws_ecs_cluster":                    ResourceCluster,
		"aws_ecs_cluster_capacity_providers": ResourceClusterCapacityProviders,
		"aws_ecs_service":                    ResourceService,
		"aws_ecs_tag":                        ResourceTag,
		"aws_ecs_task_definition":            ResourceTaskDefinition,
		"aws_ecs_task_set":                   ResourceTaskSet,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ECS
}

var ServicePackage = &servicePackage{}
