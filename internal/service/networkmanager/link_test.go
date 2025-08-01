// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanager_test

import (
	"context"
	"fmt"
	"testing"

	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfnetworkmanager "github.com/hashicorp/terraform-provider-aws/internal/service/networkmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccNetworkManagerLink_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_link.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckLinkDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccLinkConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinkExists(ctx, resourceName),
					acctest.CheckResourceAttrGlobalARNFormat(ctx, resourceName, names.AttrARN, "networkmanager", "link/{global_network_id}/{id}"),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.0.download_speed", "50"),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.0.upload_speed", "10"),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, ""),
					resource.TestCheckResourceAttr(resourceName, names.AttrProviderName, ""),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "0"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AttrImportStateIdFunc(resourceName, names.AttrARN),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetworkManagerLink_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_link.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckLinkDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccLinkConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinkExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfnetworkmanager.ResourceLink(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworkManagerLink_allAttributes(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_link.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckLinkDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccLinkConfig_allAttributes(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckLinkExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.0.download_speed", "50"),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.0.upload_speed", "10"),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "description1"),
					resource.TestCheckResourceAttr(resourceName, names.AttrProviderName, "provider1"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "type1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AttrImportStateIdFunc(resourceName, names.AttrARN),
				ImportStateVerify: true,
			},
			{
				Config: testAccLinkConfig_allAttributesUpdated(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckLinkExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.0.download_speed", "75"),
					resource.TestCheckResourceAttr(resourceName, "bandwidth.0.upload_speed", "20"),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "description2"),
					resource.TestCheckResourceAttr(resourceName, names.AttrProviderName, "provider2"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "type2"),
				),
			},
		},
	})
}

func testAccCheckLinkDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_networkmanager_link" {
				continue
			}

			_, err := tfnetworkmanager.FindLinkByTwoPartKey(ctx, conn, rs.Primary.Attributes["global_network_id"], rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Network Manager Link %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckLinkExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Network Manager Link ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerClient(ctx)

		_, err := tfnetworkmanager.FindLinkByTwoPartKey(ctx, conn, rs.Primary.Attributes["global_network_id"], rs.Primary.ID)

		return err
	}
}

func testAccLinkConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_networkmanager_global_network" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_site" "test" {
  global_network_id = aws_networkmanager_global_network.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_link" "test" {
  global_network_id = aws_networkmanager_global_network.test.id
  site_id           = aws_networkmanager_site.test.id

  bandwidth {
    download_speed = 50
    upload_speed   = 10
  }
}
`, rName)
}

func testAccLinkConfig_allAttributes(rName string) string {
	return fmt.Sprintf(`
resource "aws_networkmanager_global_network" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_site" "test" {
  global_network_id = aws_networkmanager_global_network.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_link" "test" {
  global_network_id = aws_networkmanager_global_network.test.id
  site_id           = aws_networkmanager_site.test.id

  bandwidth {
    download_speed = 50
    upload_speed   = 10
  }

  description   = "description1"
  provider_name = "provider1"
  type          = "type1"

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccLinkConfig_allAttributesUpdated(rName string) string {
	return fmt.Sprintf(`
resource "aws_networkmanager_global_network" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_site" "test" {
  global_network_id = aws_networkmanager_global_network.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_link" "test" {
  global_network_id = aws_networkmanager_global_network.test.id
  site_id           = aws_networkmanager_site.test.id

  bandwidth {
    download_speed = 75
    upload_speed   = 20
  }

  description   = "description2"
  provider_name = "provider2"
  type          = "type2"

  tags = {
    Name = %[1]q
  }
}
`, rName)
}
