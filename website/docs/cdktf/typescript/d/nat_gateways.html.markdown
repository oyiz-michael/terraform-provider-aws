---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_nat_gateways"
description: |-
    Get information on Amazon NAT Gateways.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_nat_gateways

This resource can be useful for getting back a list of NAT gateway ids to be referenced elsewhere.

## Example Usage

The following returns all NAT gateways in a specified VPC that are marked as available

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformCount, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsNatGateway } from "./.gen/providers/aws/data-aws-nat-gateway";
import { DataAwsNatGateways } from "./.gen/providers/aws/data-aws-nat-gateways";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const ngws = new DataAwsNatGateways(this, "ngws", {
      filter: [
        {
          name: "state",
          values: ["available"],
        },
      ],
      vpcId: vpcId.stringValue,
    });
    /*In most cases loops should be handled in the programming language context and 
    not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
    you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
    you need to keep this like it is.*/
    const ngwCount = TerraformCount.of(Token.asNumber(Fn.lengthOf(ngws.ids)));
    new DataAwsNatGateway(this, "ngw", {
      id: Token.asString(
        Fn.lookupNested(Fn.tolist(ngws.ids), [ngwCount.index])
      ),
      count: ngwCount,
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `filter` - (Optional) Custom filter block as described below.
* `vpcId` - (Optional) VPC ID that you want to filter from.
* `tags` - (Optional) Map of tags, each pair of which must exactly match
  a pair on the desired NAT Gateways.

### `filter`

More complex filters can be expressed using one or more `filter` sub-blocks, which take the following arguments:

* `name` - (Required) Name of the field to filter by, as defined by
  [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeNatGateways.html).
* `values` - (Required) Set of values that are accepted for the given field.
  A Nat Gateway will be selected if any one of the given values matches.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - AWS Region.
* `ids` - List of all the NAT gateway ids found.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20m`)

<!-- cache-key: cdktf-0.20.8 input-ea364eac573cb0e028ac0fa5362395ed5d6609787984fd3d66a32cc55af30030 -->