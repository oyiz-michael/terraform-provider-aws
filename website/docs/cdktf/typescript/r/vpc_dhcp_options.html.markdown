---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_vpc_dhcp_options"
description: |-
  Provides a VPC DHCP Options resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_dhcp_options

Provides a VPC DHCP Options resource.

## Example Usage

Basic usage:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcDhcpOptions } from "./.gen/providers/aws/vpc-dhcp-options";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new VpcDhcpOptions(this, "dns_resolver", {
      domainNameServers: ["8.8.8.8", "8.8.4.4"],
    });
  }
}

```

Full usage:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcDhcpOptions } from "./.gen/providers/aws/vpc-dhcp-options";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new VpcDhcpOptions(this, "foo", {
      domainName: "service.consul",
      domainNameServers: ["127.0.0.1", "10.0.0.2"],
      ipv6AddressPreferredLeaseTime: Token.asString(1440),
      netbiosNameServers: ["127.0.0.1"],
      netbiosNodeType: Token.asString(2),
      ntpServers: ["127.0.0.1"],
      tags: {
        Name: "foo-name",
      },
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `domainName` - (Optional) the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
* `domainNameServers` - (Optional) List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
* `ipv6AddressPreferredLeaseTime` - (Optional) How frequently, in seconds, a running instance with an IPv6 assigned to it goes through DHCPv6 lease renewal. Acceptable values are between 140 and 2147483647 (approximately 68 years). If no value is entered, the default lease time is 140 seconds. If you use long-term addressing for EC2 instances, you can increase the lease time and avoid frequent lease renewal requests. Lease renewal typically occurs when half of the lease time has elapsed.
* `ntpServers` - (Optional) List of NTP servers to configure.
* `netbiosNameServers` - (Optional) List of NETBIOS name servers.
* `netbiosNodeType` - (Optional) The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Remarks

* Notice that all arguments are optional but you have to specify at least one argument.
* `domainNameServers`, `netbiosNameServers`, `ntpServers` are limited by AWS to maximum four servers only.
* To actually use the DHCP Options Set you need to associate it to a VPC using [`aws_vpc_dhcp_options_association`](/docs/providers/aws/r/vpc_dhcp_options_association.html).
* If you delete a DHCP Options Set, all VPCs using it will be associated to AWS's `default` DHCP Option Set.
* In most cases unless you're configuring your own DNS you'll want to set `domainNameServers` to `AmazonProvidedDNS`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the DHCP Options Set.
* `arn` - The ARN of the DHCP Options Set.
* `ownerId` - The ID of the AWS account that owns the DHCP options set.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

You can find more technical documentation about DHCP Options Set in the
official [AWS User Guide](https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_DHCP_Options.html).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import VPC DHCP Options using the DHCP Options `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpcDhcpOptions } from "./.gen/providers/aws/vpc-dhcp-options";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    VpcDhcpOptions.generateConfigForImport(this, "myOptions", "dopt-d9070ebb");
  }
}

```

Using `terraform import`, import VPC DHCP Options using the DHCP Options `id`. For example:

```console
% terraform import aws_vpc_dhcp_options.my_options dopt-d9070ebb
```

<!-- cache-key: cdktf-0.20.8 input-362bd8ac8e3243385484da78fcc2f52431555ed5f258290977fe3f6ed405cd4a -->