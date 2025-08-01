// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanager_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	awstypes "github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfnetworkmanager "github.com/hashicorp/terraform-provider-aws/internal/service/networkmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccNetworkManagerSiteToSiteVPNAttachment_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v awstypes.SiteToSiteVpnAttachment
	resourceName := "aws_networkmanager_site_to_site_vpn_attachment.test"
	vpnResourceName := "aws_vpn_connection.test"
	coreNetworkResourceName := "aws_networkmanager_core_network.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	bgpASN := sdkacctest.RandIntRange(64512, 65534)
	vpnIP, err := sdkacctest.RandIpAddress("172.0.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckSiteToSiteVPNAttachmentDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccSiteToSiteVPNAttachmentConfig_basic(rName, bgpASN, vpnIP),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSiteToSiteVPNAttachmentExists(ctx, resourceName, &v),
					acctest.MatchResourceAttrGlobalARN(ctx, resourceName, names.AttrARN, "networkmanager", regexache.MustCompile(`attachment/.+`)),
					resource.TestCheckResourceAttr(resourceName, "attachment_policy_rule_number", "1"),
					resource.TestCheckResourceAttr(resourceName, "attachment_type", "SITE_TO_SITE_VPN"),
					resource.TestCheckResourceAttrPair(resourceName, "core_network_arn", coreNetworkResourceName, names.AttrARN),
					resource.TestCheckResourceAttrSet(resourceName, "core_network_id"),
					resource.TestCheckResourceAttr(resourceName, "edge_location", acctest.Region()),
					acctest.CheckResourceAttrAccountID(ctx, resourceName, names.AttrOwnerAccountID),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrResourceARN, vpnResourceName, names.AttrARN),
					resource.TestCheckResourceAttr(resourceName, "segment_name", "shared"),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrState),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttrPair(resourceName, "vpn_connection_arn", vpnResourceName, names.AttrARN),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{names.AttrState},
			},
		},
	})
}

func TestAccNetworkManagerSiteToSiteVPNAttachment_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v awstypes.SiteToSiteVpnAttachment
	resourceName := "aws_networkmanager_site_to_site_vpn_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	bgpASN := sdkacctest.RandIntRange(64512, 65534)
	vpnIP, err := sdkacctest.RandIpAddress("172.0.0.0/24")
	if err != nil {
		t.Fatal(err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckSiteToSiteVPNAttachmentDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccSiteToSiteVPNAttachmentConfig_basic(rName, bgpASN, vpnIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSiteToSiteVPNAttachmentExists(ctx, resourceName, &v),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfnetworkmanager.ResourceSiteToSiteVPNAttachment(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckSiteToSiteVPNAttachmentExists(ctx context.Context, n string, v *awstypes.SiteToSiteVpnAttachment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerClient(ctx)

		output, err := tfnetworkmanager.FindSiteToSiteVPNAttachmentByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckSiteToSiteVPNAttachmentDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_networkmanager_site_to_site_vpn_attachment" {
				continue
			}

			_, err := tfnetworkmanager.FindSiteToSiteVPNAttachmentByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Network Manager Site To Site VPN Attachment %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccSiteToSiteVPNAttachmentConfig_base(rName string, bgpASN int, vpnIP string) string {
	return fmt.Sprintf(`
resource "aws_customer_gateway" "test" {
  bgp_asn     = %[2]d
  ip_address  = %[3]q
  type        = "ipsec.1"
  device_name = %[1]q

  tags = {
    Name = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_id = aws_customer_gateway.test.id
  type                = "ipsec.1"

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_global_network" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_core_network" "test" {
  global_network_id = aws_networkmanager_global_network.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_core_network_policy_attachment" "test" {
  core_network_id = aws_networkmanager_core_network.test.id
  policy_document = data.aws_networkmanager_core_network_policy_document.test.json
}

data "aws_region" "current" {}

data "aws_networkmanager_core_network_policy_document" "test" {
  core_network_configuration {
    vpn_ecmp_support = false
    asn_ranges       = ["64512-64555"]
    edge_locations {
      location = data.aws_region.current.region
      asn      = 64512
    }
  }

  segments {
    name                          = "shared"
    description                   = "SegmentForSharedServices"
    require_attachment_acceptance = true
  }

  segment_actions {
    action     = "share"
    mode       = "attachment-route"
    segment    = "shared"
    share_with = ["*"]
  }

  attachment_policies {
    rule_number     = 1
    condition_logic = "or"

    conditions {
      type     = "tag-value"
      operator = "equals"
      key      = "segment"
      value    = "shared"
    }

    action {
      association_method = "constant"
      segment            = "shared"
    }
  }
}
`, rName, bgpASN, vpnIP)
}

func testAccSiteToSiteVPNAttachmentConfig_basic(rName string, bgpASN int, vpnIP string) string {
	return acctest.ConfigCompose(testAccSiteToSiteVPNAttachmentConfig_base(rName, bgpASN, vpnIP), `
resource "aws_networkmanager_site_to_site_vpn_attachment" "test" {
  core_network_id    = aws_networkmanager_core_network_policy_attachment.test.core_network_id
  vpn_connection_arn = aws_vpn_connection.test.arn

  tags = {
    segment = "shared"
  }
}

resource "aws_networkmanager_attachment_accepter" "test" {
  attachment_id   = aws_networkmanager_site_to_site_vpn_attachment.test.id
  attachment_type = aws_networkmanager_site_to_site_vpn_attachment.test.attachment_type
}
`)
}
