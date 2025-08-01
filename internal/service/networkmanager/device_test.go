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

func TestAccNetworkManagerDevice_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName),
					acctest.CheckResourceAttrGlobalARNFormat(ctx, resourceName, names.AttrARN, "networkmanager", "device/{global_network_id}/{id}"),
					resource.TestCheckResourceAttr(resourceName, "aws_location.#", "0"),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, ""),
					resource.TestCheckResourceAttr(resourceName, "location.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "model", ""),
					resource.TestCheckResourceAttr(resourceName, "serial_number", ""),
					resource.TestCheckResourceAttr(resourceName, "site_id", ""),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "0"),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, ""),
					resource.TestCheckResourceAttr(resourceName, "vendor", ""),
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

func TestAccNetworkManagerDevice_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfnetworkmanager.ResourceDevice(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworkManagerDevice_allAttributes(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_device.test"
	site1ResourceName := "aws_networkmanager_site.test1"
	site2ResourceName := "aws_networkmanager_site.test2"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceConfig_allAttributes(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "description1"),
					resource.TestCheckResourceAttr(resourceName, "location.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "location.0.address", "Address 1"),
					resource.TestCheckResourceAttr(resourceName, "location.0.latitude", "1.1"),
					resource.TestCheckResourceAttr(resourceName, "location.0.longitude", "-1.1"),
					resource.TestCheckResourceAttr(resourceName, "model", "model1"),
					resource.TestCheckResourceAttr(resourceName, "serial_number", "sn1"),
					resource.TestCheckResourceAttrPair(resourceName, "site_id", site1ResourceName, names.AttrID),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "type1"),
					resource.TestCheckResourceAttr(resourceName, "vendor", "vendor1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AttrImportStateIdFunc(resourceName, names.AttrARN),
				ImportStateVerify: true,
			},
			{
				Config: testAccDeviceConfig_allAttributesUpdated(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, names.AttrDescription, "description2"),
					resource.TestCheckResourceAttr(resourceName, "location.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "location.0.address", "Address 2"),
					resource.TestCheckResourceAttr(resourceName, "location.0.latitude", "22"),
					resource.TestCheckResourceAttr(resourceName, "location.0.longitude", "-22"),
					resource.TestCheckResourceAttr(resourceName, "model", "model2"),
					resource.TestCheckResourceAttr(resourceName, "serial_number", "sn2"),
					resource.TestCheckResourceAttrPair(resourceName, "site_id", site2ResourceName, names.AttrID),
					resource.TestCheckResourceAttr(resourceName, names.AttrType, "type2"),
					resource.TestCheckResourceAttr(resourceName, "vendor", "vendor2"),
				),
			},
		},
	})
}

func TestAccNetworkManagerDevice_awsLocation(t *testing.T) { // nosemgrep:ci.aws-in-func-name
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_device.test"
	subnetResourceName := "aws_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.NetworkManagerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckDeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccDeviceConfig_awsLocation(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "aws_location.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "aws_location.0.subnet_arn", subnetResourceName, names.AttrARN),
					resource.TestCheckResourceAttr(resourceName, "aws_location.0.zone", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AttrImportStateIdFunc(resourceName, names.AttrARN),
				ImportStateVerify: true,
			},
			{
				Config: testAccDeviceConfig_awsLocationUpdated(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckDeviceExists(ctx, resourceName),
					resource.TestCheckResourceAttr(resourceName, "aws_location.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "aws_location.0.subnet_arn", ""),
					resource.TestCheckResourceAttrPair(resourceName, "aws_location.0.zone", subnetResourceName, names.AttrAvailabilityZone),
				),
			},
		},
	})
}

func testAccCheckDeviceDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_networkmanager_device" {
				continue
			}

			_, err := tfnetworkmanager.FindDeviceByTwoPartKey(ctx, conn, rs.Primary.Attributes["global_network_id"], rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Network Manager Device %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckDeviceExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Network Manager Device ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerClient(ctx)

		_, err := tfnetworkmanager.FindDeviceByTwoPartKey(ctx, conn, rs.Primary.Attributes["global_network_id"], rs.Primary.ID)

		return err
	}
}

func testAccDeviceConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_networkmanager_global_network" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_device" "test" {
  global_network_id = aws_networkmanager_global_network.test.id
}
`, rName)
}

func testAccDeviceConfig_allAttributes(rName string) string {
	return fmt.Sprintf(`
resource "aws_networkmanager_global_network" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_site" "test1" {
  global_network_id = aws_networkmanager_global_network.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_site" "test2" {
  global_network_id = aws_networkmanager_global_network.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_device" "test" {
  global_network_id = aws_networkmanager_global_network.test.id

  description   = "description1"
  model         = "model1"
  serial_number = "sn1"
  site_id       = aws_networkmanager_site.test1.id
  type          = "type1"
  vendor        = "vendor1"

  location {
    address   = "Address 1"
    latitude  = "1.1"
    longitude = "-1.1"
  }

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccDeviceConfig_allAttributesUpdated(rName string) string {
	return fmt.Sprintf(`
resource "aws_networkmanager_global_network" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_site" "test1" {
  global_network_id = aws_networkmanager_global_network.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_site" "test2" {
  global_network_id = aws_networkmanager_global_network.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_device" "test" {
  global_network_id = aws_networkmanager_global_network.test.id

  description   = "description2"
  model         = "model2"
  serial_number = "sn2"
  site_id       = aws_networkmanager_site.test2.id
  type          = "type2"
  vendor        = "vendor2"

  location {
    address   = "Address 2"
    latitude  = "22"
    longitude = "-22"
  }

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccDeviceConfig_awsLocation(rName string) string { // nosempgrep:aws-in-func-name
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptInDefaultExclude(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_block        = "10.0.0.0/24"
  vpc_id            = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_global_network" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_device" "test" {
  global_network_id = aws_networkmanager_global_network.test.id

  aws_location {
    subnet_arn = aws_subnet.test.arn
  }

  tags = {
    Name = %[1]q
  }
}
`, rName))
}

func testAccDeviceConfig_awsLocationUpdated(rName string) string { // nosemgrep:ci.aws-in-func-name
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptInDefaultExclude(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_block        = "10.0.0.0/24"
  vpc_id            = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_global_network" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_networkmanager_device" "test" {
  global_network_id = aws_networkmanager_global_network.test.id

  aws_location {
    zone = aws_subnet.test.availability_zone
  }

  tags = {
    Name = %[1]q
  }
}
`, rName))
}
