// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/hashicorp/aws-sdk-go-base/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_vpc_dhcp_options", name="DHCP Options")
// @Tags(identifierAttribute="id")
// @Testing(tagsTest=false)
func resourceVPCDHCPOptions() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceVPCDHCPOptionsCreate,
		ReadWithoutTimeout:   resourceVPCDHCPOptionsRead,
		UpdateWithoutTimeout: resourceVPCDHCPOptionsUpdate,
		DeleteWithoutTimeout: resourceVPCDHCPOptionsDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		// Keep in sync with aws_default_vpc_dhcp_options' schema.
		// See notes in vpc_default_vpc_dhcp_options.go.
		Schema: map[string]*schema.Schema{
			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrDomainName: {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{names.AttrDomainName, "domain_name_servers", "ipv6_address_preferred_lease_time", "netbios_name_servers", "netbios_node_type", "ntp_servers"},
			},
			"domain_name_servers": {
				Type:         schema.TypeList,
				Optional:     true,
				ForceNew:     true,
				Elem:         &schema.Schema{Type: schema.TypeString},
				AtLeastOneOf: []string{names.AttrDomainName, "domain_name_servers", "ipv6_address_preferred_lease_time", "netbios_name_servers", "netbios_node_type", "ntp_servers"},
			},
			"ipv6_address_preferred_lease_time": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{names.AttrDomainName, "domain_name_servers", "ipv6_address_preferred_lease_time", "netbios_name_servers", "netbios_node_type", "ntp_servers"},
			},
			"netbios_name_servers": {
				Type:         schema.TypeList,
				Optional:     true,
				ForceNew:     true,
				Elem:         &schema.Schema{Type: schema.TypeString},
				AtLeastOneOf: []string{names.AttrDomainName, "domain_name_servers", "ipv6_address_preferred_lease_time", "netbios_name_servers", "netbios_node_type", "ntp_servers"},
			},
			"netbios_node_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				AtLeastOneOf: []string{names.AttrDomainName, "domain_name_servers", "ipv6_address_preferred_lease_time", "netbios_name_servers", "netbios_node_type", "ntp_servers"},
			},
			"ntp_servers": {
				Type:         schema.TypeList,
				Optional:     true,
				ForceNew:     true,
				Elem:         &schema.Schema{Type: schema.TypeString},
				AtLeastOneOf: []string{names.AttrDomainName, "domain_name_servers", "ipv6_address_preferred_lease_time", "netbios_name_servers", "netbios_node_type", "ntp_servers"},
			},
			names.AttrOwnerID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			names.AttrTags:    tftags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
		},
	}
}

var (
	optionsMap = newDHCPOptionsMap(map[string]string{
		names.AttrDomainName:                "domain-name",
		"domain_name_servers":               "domain-name-servers",
		"ipv6_address_preferred_lease_time": "ipv6-address-preferred-lease-time",
		"netbios_name_servers":              "netbios-name-servers",
		"netbios_node_type":                 "netbios-node-type",
		"ntp_servers":                       "ntp-servers",
	})
)

func resourceVPCDHCPOptionsCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	dhcpConfigurations, err := optionsMap.resourceDataToDHCPConfigurations(d)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating EC2 DHCP Options: %s", err)
	}

	input := &ec2.CreateDhcpOptionsInput{
		DhcpConfigurations: dhcpConfigurations,
		TagSpecifications:  getTagSpecificationsIn(ctx, awstypes.ResourceTypeDhcpOptions),
	}

	output, err := conn.CreateDhcpOptions(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating EC2 DHCP Options: %s", err)
	}

	d.SetId(aws.ToString(output.DhcpOptions.DhcpOptionsId))

	return append(diags, resourceVPCDHCPOptionsRead(ctx, d, meta)...)
}

func resourceVPCDHCPOptionsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	c := meta.(*conns.AWSClient)
	conn := c.EC2Client(ctx)

	opts, err := tfresource.RetryWhenNewResourceNotFound(ctx, ec2PropagationTimeout, func(ctx context.Context) (*awstypes.DhcpOptions, error) {
		return findDHCPOptionsByID(ctx, conn, d.Id())
	}, d.IsNewResource())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] EC2 DHCP Options Set %s not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 DHCP Options (%s): %s", d.Id(), err)
	}

	ownerID := aws.ToString(opts.OwnerId)
	d.Set(names.AttrARN, dhcpOptionsARN(ctx, c, ownerID, d.Id()))
	d.Set(names.AttrOwnerID, ownerID)

	err = optionsMap.dhcpConfigurationsToResourceData(opts.DhcpConfigurations, d)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 DHCP Options (%s): %s", d.Id(), err)
	}

	setTagsOut(ctx, opts.Tags)

	return diags
}

func resourceVPCDHCPOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics

	// Tags only.

	return append(diags, resourceVPCDHCPOptionsRead(ctx, d, meta)...)
}

func resourceVPCDHCPOptionsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Client(ctx)

	vpcs, err := findVPCs(ctx, conn, &ec2.DescribeVpcsInput{
		Filters: newAttributeFilterList(map[string]string{
			"dhcp-options-id": d.Id(),
		}),
	})

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 DHCP Options Set (%s) associated VPCs: %s", d.Id(), err)
	}

	for _, v := range vpcs {
		vpcID := aws.ToString(v.VpcId)

		log.Printf("[INFO] Disassociating EC2 DHCP Options Set (%s) from VPC (%s)", d.Id(), vpcID)
		input := ec2.AssociateDhcpOptionsInput{
			DhcpOptionsId: aws.String(defaultDHCPOptionsID),
			VpcId:         aws.String(vpcID),
		}
		_, err := conn.AssociateDhcpOptions(ctx, &input)

		if tfawserr.ErrCodeEquals(err, errCodeInvalidVPCIDNotFound) {
			continue
		}

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "disassociating EC2 DHCP Options Set (%s) from VPC (%s): %s", d.Id(), vpcID, err)
		}
	}

	input := &ec2.DeleteDhcpOptionsInput{
		DhcpOptionsId: aws.String(d.Id()),
	}

	log.Printf("[INFO] Deleting EC2 DHCP Options Set: %s", d.Id())
	_, err = tfresource.RetryWhenAWSErrCodeEquals(ctx, d.Timeout(schema.TimeoutDelete), func() (any, error) {
		return conn.DeleteDhcpOptions(ctx, input)
	}, errCodeDependencyViolation)

	if tfawserr.ErrCodeEquals(err, errCodeInvalidDHCPOptionsIDNotFound) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting EC2 DHCP Options Set (%s): %s", d.Id(), err)
	}

	return diags
}

// dhcpOptionsMap represents a mapping of Terraform resource attribute name to AWS API DHCP Option name.
type dhcpOptionsMap struct {
	tfToApi map[string]string
	apiToTf map[string]string
}

func newDHCPOptionsMap(tfToApi map[string]string) *dhcpOptionsMap {
	apiToTf := make(map[string]string)

	for k, v := range tfToApi {
		apiToTf[v] = k
	}

	return &dhcpOptionsMap{
		tfToApi: tfToApi,
		apiToTf: apiToTf,
	}
}

// dhcpConfigurationsToResourceData sets Terraform ResourceData from a list of AWS API DHCP configurations.
func (m *dhcpOptionsMap) dhcpConfigurationsToResourceData(dhcpConfigurations []awstypes.DhcpConfiguration, d *schema.ResourceData) error {
	// Clear existing values
	for tfName := range m.tfToApi {
		d.Set(tfName, nil)
	}

	for _, dhcpConfiguration := range dhcpConfigurations {
		apiName := aws.ToString(dhcpConfiguration.Key)
		tfName, ok := m.apiToTf[apiName]
		if !ok {
			return fmt.Errorf("unsupported DHCP option: %s", apiName)
		}

		currentValue := d.Get(tfName)

		switch currentValue.(type) {
		case string:
			d.Set(tfName, dhcpConfiguration.Values[0].Value)
		case []any:
			var values []string
			for _, v := range dhcpConfiguration.Values {
				if v.Value != nil {
					values = append(values, aws.ToString(v.Value))
				}
			}
			d.Set(tfName, values)
		default:
			return fmt.Errorf("attribute (%s) is of unsupported type: %T", tfName, currentValue)
		}
	}

	return nil
}

// resourceDataToDHCPConfigurations returns a list of AWS API DHCP configurations from Terraform ResourceData.
func (m *dhcpOptionsMap) resourceDataToDHCPConfigurations(d *schema.ResourceData) ([]awstypes.NewDhcpConfiguration, error) {
	var output []awstypes.NewDhcpConfiguration

	for tfName, apiName := range m.tfToApi {
		value := d.Get(tfName)
		switch v := value.(type) {
		case string:
			if v != "" {
				output = append(output, awstypes.NewDhcpConfiguration{
					Key:    aws.String(apiName),
					Values: []string{v},
				})
			}
		case []any:
			var values []string
			for _, item := range v {
				if str, ok := item.(string); ok && str != "" {
					values = append(values, str)
				}
			}
			if len(values) > 0 {
				output = append(output, awstypes.NewDhcpConfiguration{
					Key:    aws.String(apiName),
					Values: values,
				})
			}
		default:
			return nil, fmt.Errorf("attribute (%s) is of unsupported type: %T", tfName, value)
		}
	}

	return output, nil
}

func dhcpOptionsARN(ctx context.Context, c *conns.AWSClient, accountID, dhcpOptionsID string) string {
	return c.RegionalARNWithAccount(ctx, names.EC2, accountID, "dhcp-options/"+dhcpOptionsID)
}
