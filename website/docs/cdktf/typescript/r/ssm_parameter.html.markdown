---
subcategory: "SSM (Systems Manager)"
layout: "aws"
page_title: "AWS: aws_ssm_parameter"
description: |-
  Provides a SSM Parameter resource
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ssm_parameter

Provides an SSM Parameter resource.

~> **Note:** The `overwrite` argument makes it possible to overwrite an existing SSM Parameter created outside of Terraform.

-> **Note:** Write-Only argument `valueWo` is available to use in place of `value`. Write-Only arguments are supported in HashiCorp Terraform 1.11.0 and later. [Learn more](https://developer.hashicorp.com/terraform/language/resources/ephemeral#write-only-arguments).

## Example Usage

### Basic example

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmParameter } from "./.gen/providers/aws/ssm-parameter";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SsmParameter(this, "foo", {
      name: "foo",
      type: "String",
      value: "bar",
    });
  }
}

```

### Encrypted string using default SSM KMS key

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DbInstance } from "./.gen/providers/aws/db-instance";
import { SsmParameter } from "./.gen/providers/aws/ssm-parameter";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DbInstance(this, "default", {
      allocatedStorage: 10,
      dbName: "mydb",
      dbSubnetGroupName: "my_database_subnet_group",
      engine: "mysql",
      engineVersion: "5.7.16",
      instanceClass: "db.t2.micro",
      parameterGroupName: "default.mysql5.7",
      password: databaseMasterPassword.stringValue,
      storageType: "gp2",
      username: "foo",
    });
    new SsmParameter(this, "secret", {
      description: "The parameter description",
      name: "/production/database/password/master",
      tags: {
        environment: "production",
      },
      type: "SecureString",
      value: databaseMasterPassword.stringValue,
    });
  }
}

```

~> **Note:** The unencrypted value of a SecureString will be stored in the raw state as plain-text.
[Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
* `type` - (Required) Type of the parameter. Valid types are `String`, `StringList` and `SecureString`.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `allowedPattern` - (Optional) Regular expression used to validate the parameter value.
* `dataType` - (Optional) Data type of the parameter. Valid values: `text`, `aws:ssm:integration` and `aws:ec2:image` for AMI format, see the [Native parameter support for Amazon Machine Image IDs](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html).
* `description` - (Optional) Description of the parameter.
* `insecureValue` - (Optional, exactly one of `value`, `valueWo`  or `insecureValue` is required) Value of the parameter. **Use caution:** This value is _never_ marked as sensitive in the Terraform plan output. This argument is not valid with a `type` of `SecureString`.
* `keyId` - (Optional) KMS key ID or ARN for encrypting a SecureString.
* `overwrite` - (Optional) Overwrite an existing parameter. If not specified, defaults to `false` during create operations to avoid overwriting existing resources and then `true` for all subsequent operations once the resource is managed by Terraform. [Lifecycle rules](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle) should be used to manage non-standard update behavior.
* `tags` - (Optional) Map of tags to assign to the object. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `tier` - (Optional) Parameter tier to assign to the parameter. If not specified, will use the default parameter tier for the region. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. Downgrading an `Advanced` tier parameter to `Standard` will recreate the resource. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
* `value` - (Optional, exactly one of `value`, `valueWo` or `insecureValue` is required) Value of the parameter. This value is always marked as sensitive in the Terraform plan output, regardless of `type`. In Terraform CLI version 0.15 and later, this may require additional configuration handling for certain scenarios. For more information, see the [Terraform v0.15 Upgrade Guide](https://www.terraform.io/upgrade-guides/0-15.html#sensitive-output-values).
* `valueWo` - (Optional, Write-Only, exactly one of `value`, `valueWo` or `insecureValue` is required) Value of the parameter. This value is always marked as sensitive in the Terraform plan output, regardless of `type`. Additionally, `write-only` values are never stored to state. `valueWoVersion` can be used to trigger an update and is required with this argument. In Terraform CLI version 0.15 and later, this may require additional configuration handling for certain scenarios. For more information, see the [Terraform v0.15 Upgrade Guide](https://www.terraform.io/upgrade-guides/0-15.html#sensitive-output-values).
* `valueWoVersion` - (Optional) Used together with `valueWo` to trigger an update. Increment this value when an update to the `valueWo` is required.

~> **NOTE:** `aws:ssm:integration` data_type parameters must be of the type `SecureString` and the name must start with the prefix `/d9d01087-4a3f-49e0-b0b4-d568d7826553/ssm/integrations/webhook/`. See [here](https://docs.aws.amazon.com/systems-manager/latest/userguide/creating-integrations.html) for information on the usage of `aws:ssm:integration` parameters.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the parameter.
* `hasValueWo` - Indicates whether the resource has a `valueWo` set.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).
* `version` - Version of the parameter.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SSM Parameters using the parameter store `name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmParameter } from "./.gen/providers/aws/ssm-parameter";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SsmParameter.generateConfigForImport(
      this,
      "myParam",
      "/my_path/my_paramname"
    );
  }
}

```

Using `terraform import`, import SSM Parameters using the parameter store `name`. For example:

```console
% terraform import aws_ssm_parameter.my_param /my_path/my_paramname
```

<!-- cache-key: cdktf-0.20.8 input-a232bcf8a0c2cf54365938efe48323f887180b91090ca98841d9a0e0edcf7470 -->