---
subcategory: "ECS (Elastic Container)"
layout: "aws"
page_title: "AWS: aws_ecs_account_setting_default"
description: |-
  Provides an ECS Default account setting.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ecs_account_setting_default

Provides an ECS default account setting for a specific ECS Resource name within a specific region. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html).

~> **NOTE:** The AWS API does not delete this resource. When you run `destroy`, the provider will attempt to disable the setting.

~> **NOTE:** Your AWS account may not support disabling `containerInstanceLongArnFormat`, `serviceLongArnFormat`, and `taskLongArnFormat`. If your account does not support disabling these, "destroying" this resource will not disable the setting nor cause a Terraform error. However, the AWS Provider will log an AWS error: `InvalidParameterException: You can no longer disable Long Arn settings`.

## Example Usage

### Enable the long task ARN format

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsAccountSettingDefault } from "./.gen/providers/aws/ecs-account-setting-default";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EcsAccountSettingDefault(this, "test", {
      name: "taskLongArnFormat",
      value: "enabled",
    });
  }
}

```

### Set the default log driver mode to non-blocking

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsAccountSettingDefault } from "./.gen/providers/aws/ecs-account-setting-default";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EcsAccountSettingDefault(this, "test", {
      name: "defaultLogDriverMode",
      value: "non-blocking",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Name of the account setting to set.
* `value` - (Required) State of the setting.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `prinicpal_arn` - ARN that identifies the account setting.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ECS Account Setting defaults using the `name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsAccountSettingDefault } from "./.gen/providers/aws/ecs-account-setting-default";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    EcsAccountSettingDefault.generateConfigForImport(
      this,
      "example",
      "taskLongArnFormat"
    );
  }
}

```

Using `terraform import`, import ECS Account Setting defaults using the `name`. For example:

```console
% terraform import aws_ecs_account_setting_default.example taskLongArnFormat
```

<!-- cache-key: cdktf-0.20.8 input-06eff1c2390f072338d83d0bbd597eb602acc3f61d38c816c1e6efbae83da258 -->