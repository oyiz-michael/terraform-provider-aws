---
subcategory: "Transit Gateway"
layout: "aws"
page_title: "AWS: aws_ec2_transit_gateway_connect_peer"
description: |-
  Manages an EC2 Transit Gateway Connect Peer
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ec2_transit_gateway_connect_peer

Manages an EC2 Transit Gateway Connect Peer.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ec2_transit_gateway_connect import Ec2TransitGatewayConnect
from imports.aws.ec2_transit_gateway_connect_peer import Ec2TransitGatewayConnectPeer
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = Ec2TransitGatewayConnect(self, "example",
            transit_gateway_id=Token.as_string(aws_ec2_transit_gateway_example.id),
            transport_attachment_id=Token.as_string(aws_ec2_transit_gateway_vpc_attachment_example.id)
        )
        aws_ec2_transit_gateway_connect_peer_example =
        Ec2TransitGatewayConnectPeer(self, "example_1",
            inside_cidr_blocks=["169.254.100.0/29"],
            peer_address="10.1.2.3",
            transit_gateway_attachment_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_ec2_transit_gateway_connect_peer_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `bgp_asn` - (Optional) The BGP ASN number assigned customer device. If not provided, it will use the same BGP ASN as is associated with Transit Gateway.
* `inside_cidr_blocks` - (Required) The CIDR block that will be used for addressing within the tunnel. It must contain exactly one IPv4 CIDR block and up to one IPv6 CIDR block. The IPv4 CIDR block must be /29 size and must be within 169.254.0.0/16 range, with exception of: 169.254.0.0/29, 169.254.1.0/29, 169.254.2.0/29, 169.254.3.0/29, 169.254.4.0/29, 169.254.5.0/29, 169.254.169.248/29. The IPv6 CIDR block must be /125 size and must be within fd00::/8. The first IP from each CIDR block is assigned for customer gateway, the second and third is for Transit Gateway (An example: from range 169.254.100.0/29, .1 is assigned to customer gateway and .2 and .3 are assigned to Transit Gateway)
* `peer_address` - (Required) The IP addressed assigned to customer device, which will be used as tunnel endpoint. It can be IPv4 or IPv6 address, but must be the same address family as `transit_gateway_address`
* `tags` - (Optional) Key-value tags for the EC2 Transit Gateway Connect Peer. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `transit_gateway_address` - (Optional) The IP address assigned to Transit Gateway, which will be used as tunnel endpoint. This address must be from associated Transit Gateway CIDR block. The address must be from the same address family as `peer_address`. If not set explicitly, it will be selected from associated Transit Gateway CIDR blocks
* `transit_gateway_attachment_id` - (Required) The Transit Gateway Connect

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - EC2 Transit Gateway Connect Peer identifier
* `arn` - EC2 Transit Gateway Connect Peer ARN
* `bgp_peer_address` - The IP address assigned to customer device, which is used as BGP IP address.
* `bgp_transit_gateway_addresses` - The IP addresses assigned to Transit Gateway, which are used as BGP IP addresses.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_ec2_transit_gateway_connect_peer` using the EC2 Transit Gateway Connect Peer identifier. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ec2_transit_gateway_connect_peer import Ec2TransitGatewayConnectPeer
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Ec2TransitGatewayConnectPeer.generate_config_for_import(self, "example", "tgw-connect-peer-12345678")
```

Using `terraform import`, import `aws_ec2_transit_gateway_connect_peer` using the EC2 Transit Gateway Connect Peer identifier. For example:

```console
% terraform import aws_ec2_transit_gateway_connect_peer.example tgw-connect-peer-12345678
```

<!-- cache-key: cdktf-0.20.8 input-c1b56749987c62ffee7edef8043db9ebfb4afd1baec3259d5198480fb51af7eb -->