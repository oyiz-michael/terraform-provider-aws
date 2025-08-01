// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package quicksight

import (
	"context"
	"unique"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	inttypes "github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/internal/vcr"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*inttypes.ServicePackageFrameworkDataSource {
	return []*inttypes.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*inttypes.ServicePackageFrameworkResource {
	return []*inttypes.ServicePackageFrameworkResource{
		{
			Factory:  newAccountSettingsResource,
			TypeName: "aws_quicksight_account_settings",
			Name:     "Account Settings",
			Region:   unique.Make(inttypes.ResourceRegionDisabled()),
		},
		{
			Factory:  newFolderMembershipResource,
			TypeName: "aws_quicksight_folder_membership",
			Name:     "Folder Membership",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newIAMPolicyAssignmentResource,
			TypeName: "aws_quicksight_iam_policy_assignment",
			Name:     "IAM Policy Assignment",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newIngestionResource,
			TypeName: "aws_quicksight_ingestion",
			Name:     "Ingestion",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newIPRestrictionResource,
			TypeName: "aws_quicksight_ip_restriction",
			Name:     "IP Restriction",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newKeyRegistrationResource,
			TypeName: "aws_quicksight_key_registration",
			Name:     "Key Registration",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newNamespaceResource,
			TypeName: "aws_quicksight_namespace",
			Name:     "Namespace",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newRefreshScheduleResource,
			TypeName: "aws_quicksight_refresh_schedule",
			Name:     "Refresh Schedule",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newRoleMembershipResource,
			TypeName: "aws_quicksight_role_membership",
			Name:     "Role Membership",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newTemplateAliasResource,
			TypeName: "aws_quicksight_template_alias",
			Name:     "Template Alias",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  newVPCConnectionResource,
			TypeName: "aws_quicksight_vpc_connection",
			Name:     "VPC Connection",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*inttypes.ServicePackageSDKDataSource {
	return []*inttypes.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceAnalysis,
			TypeName: "aws_quicksight_analysis",
			Name:     "Analysis",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceDataSet,
			TypeName: "aws_quicksight_data_set",
			Name:     "Data Set",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceGroup,
			TypeName: "aws_quicksight_group",
			Name:     "Group",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceTheme,
			TypeName: "aws_quicksight_theme",
			Name:     "Theme",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  dataSourceUser,
			TypeName: "aws_quicksight_user",
			Name:     "User",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*inttypes.ServicePackageSDKResource {
	return []*inttypes.ServicePackageSDKResource{
		{
			Factory:  resourceAccountSubscription,
			TypeName: "aws_quicksight_account_subscription",
			Name:     "Account Subscription",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceAnalysis,
			TypeName: "aws_quicksight_analysis",
			Name:     "Analysis",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceDashboard,
			TypeName: "aws_quicksight_dashboard",
			Name:     "Dashboard",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceDataSet,
			TypeName: "aws_quicksight_data_set",
			Name:     "Data Set",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceDataSource,
			TypeName: "aws_quicksight_data_source",
			Name:     "Data Source",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceFolder,
			TypeName: "aws_quicksight_folder",
			Name:     "Folder",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceGroup,
			TypeName: "aws_quicksight_group",
			Name:     "Group",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceGroupMembership,
			TypeName: "aws_quicksight_group_membership",
			Name:     "Group Membership",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceTemplate,
			TypeName: "aws_quicksight_template",
			Name:     "Template",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceTheme,
			TypeName: "aws_quicksight_theme",
			Name:     "Theme",
			Tags: unique.Make(inttypes.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			}),
			Region: unique.Make(inttypes.ResourceRegionDefault()),
		},
		{
			Factory:  resourceUser,
			TypeName: "aws_quicksight_user",
			Name:     "User",
			Region:   unique.Make(inttypes.ResourceRegionDefault()),
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.QuickSight
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*quicksight.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*quicksight.Options){
		quicksight.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *quicksight.Options) {
			if region := config[names.AttrRegion].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         p.ServicePackageName(),
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		func(o *quicksight.Options) {
			if inContext, ok := conns.FromContext(ctx); ok && inContext.VCREnabled() {
				tflog.Info(ctx, "overriding retry behavior to immediately return VCR errors")
				o.Retryer = conns.AddIsErrorRetryables(cfg.Retryer().(aws.RetryerV2), retry.IsErrorRetryableFunc(vcr.InteractionNotFoundRetryableFunc))
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return quicksight.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*quicksight.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*quicksight.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *quicksight.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*quicksight.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
