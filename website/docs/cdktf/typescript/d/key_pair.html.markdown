---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "AWS: aws_key_pair"
description: |-
    Provides details about a specific EC2 Key Pair.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_key_pair

Use this data source to get information about a specific EC2 Key Pair.

## Example Usage

The following example shows how to get a EC2 Key Pair including the public key material from its name.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformOutput, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsKeyPair } from "./.gen/providers/aws/data-aws-key-pair";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new DataAwsKeyPair(this, "example", {
      filter: [
        {
          name: "tag:Component",
          values: ["web"],
        },
      ],
      includePublicKey: true,
      keyName: "test",
    });
    new TerraformOutput(this, "fingerprint", {
      value: example.fingerprint,
    });
    new TerraformOutput(this, "id", {
      value: example.id,
    });
    new TerraformOutput(this, "name", {
      value: example.keyName,
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `keyPairId` - (Optional) Key Pair ID.
* `keyName` - (Optional) Key Pair name.
* `includePublicKey` - (Optional) Whether to include the public key material in the response.
* `filter` -  (Optional) Custom filter block as described below.

The arguments of this data source act as filters for querying the available
Key Pairs. The given filters must match exactly one Key Pair
whose data will be exported as attributes.

### filter Configuration Block

The `filter` configuration block supports the following arguments:

* `name` - (Required) Name of the filter field. Valid values can be found in the [EC2 DescribeKeyPairs API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeKeyPairs.html).
* `values` - (Required) Set of values that are accepted for the given filter field. Results will be selected if any given value matches.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - ID of the Key Pair.
* `arn` - ARN of the Key Pair.
* `createTime` - Timestamp for when the key pair was created in ISO 8601 format.
* `fingerprint` - SHA-1 digest of the DER encoded private key.
* `keyType` - Type of key pair.
* `publicKey` - Public key material.
* `tags` - Any tags assigned to the Key Pair.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20m`)

<!-- cache-key: cdktf-0.20.8 input-50cc3969439bfed24baba154e0c03d8e9c65f35c5b9acd7c17faf74a4e586211 -->