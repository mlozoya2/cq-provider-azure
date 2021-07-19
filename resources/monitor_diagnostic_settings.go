package resources

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func MonitorDiagnosticSettings() *schema.Table {
	return &schema.Table{
		Name:        "azure_monitor_diagnostic_settings",
		Description: "DiagnosticSettingsResource the diagnostic setting resource",
		Resolver:    fetchMonitorDiagnosticSettings,
		Multiplex:   client.SubscriptionMultiplex,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "storage_account_id",
				Description: "The resource ID of the storage account to which you would like to send Diagnostic Logs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiagnosticSettings.StorageAccountID"),
			},
			{
				Name:        "service_bus_rule_id",
				Description: "The service bus rule Id of the diagnostic setting This is here to maintain backwards compatibility",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiagnosticSettings.ServiceBusRuleID"),
			},
			{
				Name:        "event_hub_authorization_rule_id",
				Description: "The resource Id for the event hub authorization rule",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiagnosticSettings.EventHubAuthorizationRuleID"),
			},
			{
				Name:        "event_hub_name",
				Description: "The name of the event hub If none is specified, the default event hub will be selected",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiagnosticSettings.EventHubName"),
			},
			{
				Name:        "workspace_id",
				Description: "The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs Example: /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/MicrosoftOperationalInsights/workspaces/viruela2",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiagnosticSettings.WorkspaceID"),
			},
			{
				Name:        "log_analytics_destination_type",
				Description: "A string indicating whether the export to Log Analytics should use the default destination type, ie AzureDiagnostics, or use a destination type constructed as follows: <normalized service identity>_<normalized category name> Possible values are: Dedicated and null (null is default)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DiagnosticSettings.LogAnalyticsDestinationType"),
			},
			{
				Name:        "id",
				Description: "Azure resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Azure resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Azure resource type",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:         "azure_monitor_diagnostic_setting_metrics",
				Description:  "MetricSettings part of MultiTenantDiagnosticSettings Specifies the settings for a particular metric",
				Resolver:     fetchMonitorDiagnosticSettingMetrics,
				AlwaysDelete: true,
				Columns: []schema.Column{
					{
						Name:        "diagnostic_setting_cq_id",
						Description: "Unique ID of azure_monitor_diagnostic_settings table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "diagnostic_setting_id",
						Description: "Unique ID of azure_monitor_diagnostic_settings table (FK)",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "time_grain",
						Description: "the timegrain of the metric in ISO8601 format",
						Type:        schema.TypeString,
					},
					{
						Name:        "category",
						Description: "Name of a Diagnostic Metric category for a resource type this setting is applied to To obtain the list of Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation",
						Type:        schema.TypeString,
					},
					{
						Name:        "enabled",
						Description: "a value indicating whether this category is enabled",
						Type:        schema.TypeBool,
					},
					{
						Name:        "retention_policy_enabled",
						Description: "a value indicating whether the retention policy is enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("RetentionPolicy.Enabled"),
					},
					{
						Name:        "retention_policy_days",
						Description: "the number of days for the retention in days A value of 0 will retain the events indefinitely",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("RetentionPolicy.Days"),
					},
				},
			},
			{
				Name:         "azure_monitor_diagnostic_setting_logs",
				Description:  "LogSettings part of MultiTenantDiagnosticSettings Specifies the settings for a particular log",
				Resolver:     fetchMonitorDiagnosticSettingLogs,
				AlwaysDelete: true,
				Columns: []schema.Column{
					{
						Name:        "diagnostic_setting_cq_id",
						Description: "Unique ID of azure_monitor_diagnostic_settings table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "diagnostic_setting_id",
						Description: "Unique ID of azure_monitor_diagnostic_settings table (FK)",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "category",
						Description: "Name of a Diagnostic Log category for a resource type this setting is applied to To obtain the list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation",
						Type:        schema.TypeString,
					},
					{
						Name:        "enabled",
						Description: "a value indicating whether this log is enabled",
						Type:        schema.TypeBool,
					},
					{
						Name:        "retention_policy_enabled",
						Description: "a value indicating whether the retention policy is enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("RetentionPolicy.Enabled"),
					},
					{
						Name:        "retention_policy_days",
						Description: "the number of days for the retention in days A value of 0 will retain the events indefinitely",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("RetentionPolicy.Days"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchMonitorDiagnosticSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	resSvc := meta.(*client.Client).Services().Resources.Resources
	svc := meta.(*client.Client).Services().Monitor.DiagnosticSettings
	resources, err := resSvc.List(ctx, "", "", nil)
	if err != nil {
		return err
	}
	for _, r := range resources.Values() {
		response, err := svc.List(ctx, *r.ID)
		if err != nil {
			return err
		}
		if response.Value == nil {
			continue
		}
		res <- *response.Value
	}

	return nil
}
func fetchMonitorDiagnosticSettingMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(insights.DiagnosticSettingsResource)
	if !ok {
		return fmt.Errorf("expected insights.DiagnosticSettingsResource but got %T", parent.Item)
	}
	if p.DiagnosticSettings == nil ||
		p.DiagnosticSettings.Metrics == nil {
		return nil
	}

	res <- *p.DiagnosticSettings.Metrics
	return nil
}
func fetchMonitorDiagnosticSettingLogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(insights.DiagnosticSettingsResource)
	if !ok {
		return fmt.Errorf("expected insights.DiagnosticSettingsResource but got %T", parent.Item)
	}
	if p.DiagnosticSettings == nil ||
		p.DiagnosticSettings.Logs == nil {
		return nil
	}

	res <- *p.DiagnosticSettings.Logs
	return nil
}