---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_vpc_endpoint_private_dns"
description: |-
  Terraform resource for enabling private DNS on an AWS VPC (Virtual Private Cloud) Endpoint.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_endpoint_private_dns

Terraform resource for enabling private DNS on an AWS VPC (Virtual Private Cloud) Endpoint.

~> When using this resource, the `privateDnsEnabled` argument should be omitted on the parent `aws_vpc_endpoint` resource.
Setting the value both places can lead to unintended behavior and persistent differences.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcEndpointPrivateDns } from "./.gen/providers/aws/vpc-endpoint-private-dns";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new VpcEndpointPrivateDns(this, "example", {
      privateDnsEnabled: true,
      vpcEndpointId: Token.asString(awsVpcEndpointExample.id),
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `privateDnsEnabled` - (Required) Indicates whether a private hosted zone is associated with the VPC. Only applicable for `Interface` endpoints.
* `vpcEndpointId` - (Required) VPC endpoint identifier.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import a VPC (Virtual Private Cloud) Endpoint Private DNS using the `vpcEndpointId`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcEndpointPrivateDns } from "./.gen/providers/aws/vpc-endpoint-private-dns";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    VpcEndpointPrivateDns.generateConfigForImport(
      this,
      "example",
      "vpce-abcd-1234"
    );
  }
}

```

Using `terraform import`, import a VPC (Virtual Private Cloud) Endpoint Private DNS using the `vpcEndpointId`. For example:

```console
% terraform import aws_vpc_endpoint_private_dns.example vpce-abcd-1234
```

<!-- cache-key: cdktf-0.20.8 input-fb61383c27ad6532408d0d99242fa17f2ccc7b0a5214264644e63a609400534d -->