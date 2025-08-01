---
subcategory: "Route 53 Resolver"
layout: "aws"
page_title: "AWS: aws_route53_resolver_endpoint"
description: |-
    Provides details about a specific Route 53 Resolver Endpoint
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_route53_resolver_endpoint

`aws_route53_resolver_endpoint` provides details about a specific Route53 Resolver Endpoint.

This data source allows to find a list of IPaddresses associated with a specific Route53 Resolver Endpoint.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRoute53ResolverEndpoint } from "./.gen/providers/aws/data-aws-route53-resolver-endpoint";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsRoute53ResolverEndpoint(this, "example", {
      resolverEndpointId: "rslvr-in-1abc2345ef678g91h",
    });
  }
}

```

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRoute53ResolverEndpoint } from "./.gen/providers/aws/data-aws-route53-resolver-endpoint";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsRoute53ResolverEndpoint(this, "example", {
      filter: [
        {
          name: "NAME",
          values: ["MyResolverExampleName"],
        },
      ],
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `resolverEndpointId` - (Optional) ID of the Route53 Resolver Endpoint.
* `filter` - (Optional) One or more name/value pairs to use as filters. There are
several valid keys, for a full reference, check out
[Route53resolver Filter value in the AWS API reference][1].

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - Computed ARN of the Route53 Resolver Endpoint.
* `direction` - Direction of the queries to or from the Resolver Endpoint .
* `ipAddresses` - List of IPaddresses that have been associated with the Resolver Endpoint.
* `protocols` - The protocols used by the Resolver endpoint.
* `resolverEndpointType` - The Resolver endpoint IP address type.
* `status` - Current status of the Resolver Endpoint.
* `vpcId` - ID of the Host VPC that the Resolver Endpoint resides in.

[1]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_Filter.html

<!-- cache-key: cdktf-0.20.8 input-24a314d2a96be3408d96ad3ceb80bcd44ce41d7d59384382a7b8b1d89c4d68e6 -->