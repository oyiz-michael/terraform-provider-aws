---
subcategory: "VPC Lattice"
layout: "aws"
page_title: "AWS: aws_vpclattice_access_log_subscription"
description: |-
  Terraform resource for managing an AWS VPC Lattice Service Network or Services Access log subscription.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpclattice_access_log_subscription

Terraform resource for managing an AWS VPC Lattice Service Network or Service Access log subscription.

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
import { VpclatticeAccessLogSubscription } from "./.gen/providers/aws/vpclattice-access-log-subscription";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new VpclatticeAccessLogSubscription(this, "example", {
      destinationArn: Token.asString(bucket.arn),
      resourceIdentifier: Token.asString(awsVpclatticeServiceNetworkExample.id),
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `destinationArn` - (Required, Forces new resource) Amazon Resource Name (ARN) of the log destination.
* `resourceIdentifier` - (Required, Forces new resource) The ID or Amazon Resource Identifier (ARN) of the service network or service. You must use the ARN if the resources specified in the operation are in different accounts.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `serviceNetworkLogType` - (Optional, Forces new resource) Type of log that monitors your Amazon VPC Lattice service networks. Valid values are: `SERVICE`, `RESOURCE`. Defaults to `SERVICE`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - ID of the access log subscription.
* `arn` - Amazon Resource Name (ARN) of the access log subscription.
* `resourceArn` - Amazon Resource Name (ARN) of the service network or service.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import VPC Lattice Access Log Subscription using the access log subscription ID. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { VpclatticeAccessLogSubscription } from "./.gen/providers/aws/vpclattice-access-log-subscription";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    VpclatticeAccessLogSubscription.generateConfigForImport(
      this,
      "example",
      "rft-8012925589"
    );
  }
}

```

Using `terraform import`, import VPC Lattice Access Log Subscription using the access log subscription ID. For example:

```console
% terraform import aws_vpclattice_access_log_subscription.example rft-8012925589
```

<!-- cache-key: cdktf-0.20.8 input-7737d7cbf0c0d9c164911abb108046a6683e3664e7d996391270ae9c257a5ce2 -->