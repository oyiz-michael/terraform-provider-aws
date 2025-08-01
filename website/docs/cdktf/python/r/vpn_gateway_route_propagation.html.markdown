---
subcategory: "VPN (Site-to-Site)"
layout: "aws"
page_title: "AWS: aws_vpn_gateway_route_propagation"
description: |-
  Requests automatic route propagation between a VPN gateway and a route table.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpn_gateway_route_propagation

Requests automatic route propagation between a VPN gateway and a route table.

~> **Note:** This resource should not be used with a route table that has
the `propagating_vgws` argument set. If that argument is set, any route
propagation not explicitly listed in its value will be removed.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.vpn_gateway_route_propagation import VpnGatewayRoutePropagation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        VpnGatewayRoutePropagation(self, "example",
            route_table_id=Token.as_string(aws_route_table_example.id),
            vpn_gateway_id=Token.as_string(aws_vpn_gateway_example.id)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `vpn_gateway_id` - The id of the `aws_vpn_gateway` to propagate routes from.
* `route_table_id` - The id of the `aws_route_table` to propagate routes into.

## Attribute Reference

This resource exports no additional attributes.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `2m`)
- `delete` - (Default `2m`)

<!-- cache-key: cdktf-0.20.8 input-7caecf96130f8a679376c032adf125cf28285804b2486f9dc36fc09bb045b365 -->