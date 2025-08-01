---
subcategory: "Network Firewall"
layout: "aws"
page_title: "AWS: aws_networkfirewall_resource_policy"
description: |-
  Retrieve information about a Network Firewall resource policy.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_networkfirewall_resource_policy

Retrieve information about a Network Firewall resource policy.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsNetworkfirewallResourcePolicy } from "./.gen/providers/aws/data-aws-networkfirewall-resource-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsNetworkfirewallResourcePolicy(this, "example", {
      resourceArn: resourcePolicyArn.stringValue,
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `resourceArn` - (Required) The Amazon Resource Name (ARN) that identifies the resource policy.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - The Amazon Resource Name (ARN) that identifies the resource policy.
* `policy` - The [policy][1] for the resource.

[1]: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/networkfirewall_resource_policy

<!-- cache-key: cdktf-0.20.8 input-8e7abd2d3562dce7bdec718e561d4fa2e059c5e4c4d74746219b58e33f3a2494 -->