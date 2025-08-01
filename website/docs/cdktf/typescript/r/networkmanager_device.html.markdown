---
subcategory: "Network Manager"
layout: "aws"
page_title: "AWS: aws_networkmanager_device"
description: |-
  Manages a Network Manager Device.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_networkmanager_device

Manages a Network Manager Device.

Use this resource to create a device in a global network. If you specify both a site ID and a location, the location of the site is used for visualization in the Network Manager console.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkmanagerDevice } from "./.gen/providers/aws/networkmanager-device";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new NetworkmanagerDevice(this, "example", {
      globalNetworkId: Token.asString(awsNetworkmanagerGlobalNetworkExample.id),
      siteId: Token.asString(awsNetworkmanagerSiteExample.id),
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `globalNetworkId` - (Required) ID of the global network.

The following arguments are optional:

* `awsLocation` - (Optional) AWS location of the device. Documented below.
* `description` - (Optional) Description of the device.
* `location` - (Optional) Location of the device. Documented below.
* `model` - (Optional) Model of device.
* `serialNumber` - (Optional) Serial number of the device.
* `siteId` - (Optional) ID of the site.
* `tags` - (Optional) Key-value tags for the device. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `type` - (Optional) Type of device.
* `vendor` - (Optional) Vendor of the device.

The `awsLocation` object supports the following:

* `subnetArn` - (Optional) ARN of the subnet that the device is located in.
* `zone` - (Optional) Zone that the device is located in. Specify the ID of an Availability Zone, Local Zone, Wavelength Zone, or an Outpost.

The `location` object supports the following:

* `address` - (Optional) Physical address.
* `latitude` - (Optional) Latitude.
* `longitude` - (Optional) Longitude.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the device.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `10m`)
* `delete` - (Default `10m`)
* `update` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_networkmanager_device` using the device ARN. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { NetworkmanagerDevice } from "./.gen/providers/aws/networkmanager-device";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    NetworkmanagerDevice.generateConfigForImport(
      this,
      "example",
      "arn:aws:networkmanager::123456789012:device/global-network-0d47f6t230mz46dy4/device-07f6fd08867abc123"
    );
  }
}

```

Using `terraform import`, import `aws_networkmanager_device` using the device ARN. For example:

```console
% terraform import aws_networkmanager_device.example arn:aws:networkmanager::123456789012:device/global-network-0d47f6t230mz46dy4/device-07f6fd08867abc123
```

<!-- cache-key: cdktf-0.20.8 input-f03efda1b457622fa934fd7975411d81ce01234046e50ee34684075df069d644 -->