---
subcategory: "EMR"
layout: "aws"
page_title: "AWS: aws_emr_block_public_access_configuration"
description: |-
  Terraform resource for managing an AWS EMR Block Public Access Configuration.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_emr_block_public_access_configuration

Terraform resource for managing an AWS EMR block public access configuration. This region level security configuration restricts the launch of EMR clusters that have associated security groups permitting public access on unspecified ports. See the [EMR Block Public Access Configuration](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-block-public-access.html) documentation for further information.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EmrBlockPublicAccessConfiguration } from "./.gen/providers/aws/emr-block-public-access-configuration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EmrBlockPublicAccessConfiguration(this, "example", {
      blockPublicSecurityGroupRules: true,
    });
  }
}

```

### Default Configuration

By default, each AWS region is equipped with a block public access configuration that prevents EMR clusters from being launched if they have security group rules permitting public access on any port except for port 22. The default configuration can be managed using this Terraform resource.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EmrBlockPublicAccessConfiguration } from "./.gen/providers/aws/emr-block-public-access-configuration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EmrBlockPublicAccessConfiguration(this, "example", {
      blockPublicSecurityGroupRules: true,
      permittedPublicSecurityGroupRuleRange: [
        {
          maxRange: 22,
          minRange: 22,
        },
      ],
    });
  }
}

```

~> **NOTE:** If an `aws_emr_block_public_access_configuration` Terraform resource is destroyed, the configuration will reset to this default configuration.

### Multiple Permitted Public Security Group Rule Ranges

The resource permits specification of multiple `permittedPublicSecurityGroupRuleRange` blocks.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EmrBlockPublicAccessConfiguration } from "./.gen/providers/aws/emr-block-public-access-configuration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EmrBlockPublicAccessConfiguration(this, "example", {
      blockPublicSecurityGroupRules: true,
      permittedPublicSecurityGroupRuleRange: [
        {
          maxRange: 22,
          minRange: 22,
        },
        {
          maxRange: 101,
          minRange: 100,
        },
      ],
    });
  }
}

```

### Disabling Block Public Access

To permit EMR clusters to be launched in the configured region regardless of associated security group rules, the Block Public Access feature can be disabled using this Terraform resource.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EmrBlockPublicAccessConfiguration } from "./.gen/providers/aws/emr-block-public-access-configuration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EmrBlockPublicAccessConfiguration(this, "example", {
      blockPublicSecurityGroupRules: false,
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `blockPublicSecurityGroupRules` - (Required) Enable or disable EMR Block Public Access.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `permittedPublicSecurityGroupRuleRange` - (Optional) Configuration block for defining permitted public security group rule port ranges. Can be defined multiple times per resource. Only valid if `blockPublicSecurityGroupRules` is set to `true`.

### `permittedPublicSecurityGroupRuleRange`

This block is used to define a range of TCP ports that should form exceptions to the Block Public Access Configuration. If an attempt is made to launch an EMR cluster in the configured region and account, with `block_public_security_group_rules = true`, the EMR cluster will be permitted to launch even if there are security group rules permitting public access to ports in this range.

* `minRange` - (Required) The first port in the range of TCP ports.
* `maxRange` - (Required) The final port in the range of TCP ports.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the current EMR Block Public Access Configuration. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EmrBlockPublicAccessConfiguration } from "./.gen/providers/aws/emr-block-public-access-configuration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    EmrBlockPublicAccessConfiguration.generateConfigForImport(
      this,
      "example",
      "current"
    );
  }
}

```

Using `terraform import`, import the current EMR Block Public Access Configuration. For example:

```console
% terraform import aws_emr_block_public_access_configuration.example current
```

<!-- cache-key: cdktf-0.20.8 input-3e1047e8a5bd123ea9321b8ed8b56c686d9574732dec139f5477aee4e6db3311 -->