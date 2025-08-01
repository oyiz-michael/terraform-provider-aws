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

func TestAccNetworkManagerConnectPeer_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v awstypes.ConnectPeer
	resourceName := "aws_networkmanager_connect_peer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	insideCidrBlocksv4 := "169.254.10.0/29"
	peerAddress := "1.1.1.1"
	protocol := "GRE"
	asn := "65501"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckConnectPeerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccConnectPeerConfig_basic(rName, insideCidrBlocksv4, peerAddress, asn, protocol),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckConnectPeerExists(ctx, resourceName, &v),
					acctest.MatchResourceAttrGlobalARN(ctx, resourceName, names.AttrARN, "networkmanager", regexache.MustCompile(`connect-peer/.+`)),
					resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.core_network_address"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.inside_cidr_blocks.0", insideCidrBlocksv4),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.peer_address", peerAddress),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.protocol", "GRE"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.bgp_configurations.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "connect_attachment_id"),
					resource.TestCheckResourceAttr(resourceName, "inside_cidr_blocks.0", insideCidrBlocksv4),
					resource.TestCheckResourceAttr(resourceName, "peer_address", peerAddress),
					resource.TestCheckResourceAttr(resourceName, "edge_location", acctest.Region()),
					resource.TestCheckResourceAttrSet(resourceName, "connect_attachment_id"),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrState),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetworkManagerConnectPeer_noDependsOn(t *testing.T) {
	ctx := acctest.Context(t)
	var v awstypes.ConnectPeer
	resourceName := "aws_networkmanager_connect_peer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	insideCidrBlocksv4 := "169.254.10.0/29"
	peerAddress := "1.1.1.1"
	protocol := "GRE"
	asn := "65501"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckConnectPeerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccConnectPeerConfig_noDependsOn(rName, insideCidrBlocksv4, peerAddress, asn, protocol),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckConnectPeerExists(ctx, resourceName, &v),
					acctest.MatchResourceAttrGlobalARN(ctx, resourceName, names.AttrARN, "networkmanager", regexache.MustCompile(`connect-peer/.+`)),
					resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "configuration.0.core_network_address"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.inside_cidr_blocks.0", insideCidrBlocksv4),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.peer_address", peerAddress),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.protocol", "GRE"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.bgp_configurations.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "connect_attachment_id"),
					resource.TestCheckResourceAttr(resourceName, "inside_cidr_blocks.0", insideCidrBlocksv4),
					resource.TestCheckResourceAttr(resourceName, "peer_address", peerAddress),
					resource.TestCheckResourceAttr(resourceName, "edge_location", acctest.Region()),
					resource.TestCheckResourceAttrSet(resourceName, "connect_attachment_id"),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrState),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetworkManagerConnectPeer_subnetARN(t *testing.T) {
	ctx := acctest.Context(t)
	var v awstypes.ConnectPeer
	resourceName := "aws_networkmanager_connect_peer.test"
	subnetResourceName := "aws_subnet.test2"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	peerAddress := "10.0.2.100" // Must be an address inside the subnet CIDR range.
	protocol := "NO_ENCAP"
	asn := "65501"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckConnectPeerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccConnectPeerConfig_subnetARN(rName, peerAddress, asn, protocol),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckConnectPeerExists(ctx, resourceName, &v),
					acctest.MatchResourceAttrGlobalARN(ctx, resourceName, names.AttrARN, "networkmanager", regexache.MustCompile(`connect-peer/.+`)),
					resource.TestCheckResourceAttr(resourceName, "configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.peer_address", peerAddress),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.protocol", "NO_ENCAP"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.bgp_configurations.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "connect_attachment_id"),
					resource.TestCheckResourceAttr(resourceName, "peer_address", peerAddress),
					resource.TestCheckResourceAttr(resourceName, "edge_location", acctest.Region()),
					resource.TestCheckResourceAttrSet(resourceName, "connect_attachment_id"),
					resource.TestCheckResourceAttrPair(resourceName, "subnet_arn", subnetResourceName, names.AttrARN),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrState),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckConnectPeerExists(ctx context.Context, n string, v *awstypes.ConnectPeer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Network Manager Connect Peer ID is set")
		}
		conn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerClient(ctx)

		output, err := tfnetworkmanager.FindConnectPeerByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckConnectPeerDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_networkmanager_connect_peer" {
				continue
			}

			_, err := tfnetworkmanager.FindConnectPeerByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Network Manager Connect Peer %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccConnectPeerConfig_base(rName string, protocol string) string {
	return acctest.ConfigCompose(
		acctest.ConfigVPCWithSubnetsIPv6(rName, 2),
		fmt.Sprintf(`
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
    vpn_ecmp_support   = false
    asn_ranges         = ["64512-64555"]
    inside_cidr_blocks = ["172.16.0.0/16"]
    edge_locations {
      location           = data.aws_region.current.region
      asn                = 64512
      inside_cidr_blocks = ["172.16.0.0/18"]
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

resource "aws_networkmanager_vpc_attachment" "test" {
  subnet_arns     = aws_subnet.test[*].arn
  core_network_id = aws_networkmanager_core_network_policy_attachment.test.core_network_id
  vpc_arn         = aws_vpc.test.arn
  tags = {
    segment = "shared"
  }
}

resource "aws_networkmanager_attachment_accepter" "test" {
  attachment_id   = aws_networkmanager_vpc_attachment.test.id
  attachment_type = aws_networkmanager_vpc_attachment.test.attachment_type
}

resource "aws_networkmanager_connect_attachment" "test" {
  core_network_id         = aws_networkmanager_core_network.test.id
  transport_attachment_id = aws_networkmanager_vpc_attachment.test.id
  edge_location           = aws_networkmanager_vpc_attachment.test.edge_location
  options {
    protocol = %[2]q
  }
  tags = {
    segment = "shared"
  }
  depends_on = [
    "aws_networkmanager_attachment_accepter.test"
  ]
}

resource "aws_networkmanager_attachment_accepter" "test2" {
  attachment_id   = aws_networkmanager_connect_attachment.test.id
  attachment_type = aws_networkmanager_connect_attachment.test.attachment_type
}
`, rName, protocol))
}

func testAccConnectPeerConfig_basic(rName string, insideCidrBlocks string, peerAddress string, asn string, protocol string) string {
	return acctest.ConfigCompose(testAccConnectPeerConfig_base(rName, protocol), fmt.Sprintf(`
resource "aws_networkmanager_connect_peer" "test" {
  connect_attachment_id = aws_networkmanager_connect_attachment.test.id
  peer_address          = %[3]q
  bgp_options {
    peer_asn = %[4]q
  }
  inside_cidr_blocks = [
	%[2]q
  ]
  tags = {
    Name = %[1]q
  }
  depends_on = [
    "aws_networkmanager_attachment_accepter.test"
  ]
}
`, rName, insideCidrBlocks, peerAddress, asn))
}

func testAccConnectPeerConfig_noDependsOn(rName string, insideCidrBlocks string, peerAddress string, asn string, protocol string) string {
	return acctest.ConfigCompose(testAccConnectPeerConfig_base(rName, protocol), fmt.Sprintf(`
resource "aws_networkmanager_connect_peer" "test" {
  connect_attachment_id = aws_networkmanager_connect_attachment.test.id
  peer_address          = %[3]q
  bgp_options {
    peer_asn = %[4]q
  }
  inside_cidr_blocks = [
	%[2]q
  ]
  tags = {
    Name = %[1]q
  }
}
`, rName, insideCidrBlocks, peerAddress, asn))
}

func testAccConnectPeerConfig_subnetARN(rName string, peerAddress string, asn string, protocol string) string {
	return acctest.ConfigCompose(testAccConnectPeerConfig_base(rName, protocol), fmt.Sprintf(`
resource "aws_networkmanager_connect_peer" "test" {
  connect_attachment_id = aws_networkmanager_connect_attachment.test.id
  peer_address          = %[2]q
  bgp_options {
    peer_asn = %[3]q
  }
  subnet_arn = aws_subnet.test2.arn
  tags = {
    Name = %[1]q
  }
  depends_on = [
    "aws_networkmanager_attachment_accepter.test"
  ]
}

resource "aws_subnet" "test2" {
  vpc_id     = aws_vpc.test.id
  cidr_block = cidrsubnet(aws_vpc.test.cidr_block, 8, 2)
}
`, rName, peerAddress, asn))
}
