// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanager_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfnetworkmanager "github.com/hashicorp/terraform-provider-aws/internal/service/networkmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccNetworkManagerGlobalNetwork_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_global_network.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckGlobalNetworkDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGlobalNetworkConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGlobalNetworkExists(ctx, resourceName),
					acctest.MatchResourceAttrGlobalARN(ctx, resourceName, names.AttrARN, "networkmanager", regexache.MustCompile(`global-network/global-network-.+`)),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, ""),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "0"),
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

func TestAccNetworkManagerGlobalNetwork_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_global_network.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckGlobalNetworkDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGlobalNetworkConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGlobalNetworkExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfnetworkmanager.ResourceGlobalNetwork(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworkManagerGlobalNetwork_description(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_global_network.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckGlobalNetworkDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGlobalNetworkConfig_description("description1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGlobalNetworkExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "description1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGlobalNetworkConfig_description("description2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGlobalNetworkExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "description2"),
				),
			},
		},
	})
}

func testAccCheckGlobalNetworkDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_networkmanager_global_network" {
				continue
			}

			_, err := tfnetworkmanager.FindGlobalNetworkByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Network Manager Global Network %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckGlobalNetworkExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Network Manager Global Network ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerClient(ctx)

		_, err := tfnetworkmanager.FindGlobalNetworkByID(ctx, conn, rs.Primary.ID)

		return err
	}
}

func testAccGlobalNetworkConfig_basic() string {
	return `
resource "aws_networkmanager_global_network" "test" {}
`
}

func testAccGlobalNetworkConfig_description(description string) string {
	return fmt.Sprintf(`
resource "aws_networkmanager_global_network" "test" {
  description = %[1]q
}
`, description)
}
