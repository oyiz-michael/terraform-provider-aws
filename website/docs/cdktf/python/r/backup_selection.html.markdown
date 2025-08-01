---
subcategory: "Backup"
layout: "aws"
page_title: "AWS: aws_backup_selection"
description: |-
  Manages selection conditions for AWS Backup plan resources.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_backup_selection

Manages selection conditions for AWS Backup plan resources.

## Example Usage

### IAM Role

-> For more information about creating and managing IAM Roles for backups and restores, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/iam-service-roles.html).

The below example creates an IAM role with the default managed IAM Policy for allowing AWS Backup to create backups.

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.backup_selection import BackupSelection
from imports.aws.data_aws_iam_policy_document import DataAwsIamPolicyDocument
from imports.aws.iam_role import IamRole
from imports.aws.iam_role_policy_attachment import IamRolePolicyAttachment
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name, *, name, planId):
        super().__init__(scope, name)
        assume_role = DataAwsIamPolicyDocument(self, "assume_role",
            statement=[DataAwsIamPolicyDocumentStatement(
                actions=["sts:AssumeRole"],
                effect="Allow",
                principals=[DataAwsIamPolicyDocumentStatementPrincipals(
                    identifiers=["backup.amazonaws.com"],
                    type="Service"
                )
                ]
            )
            ]
        )
        example = IamRole(self, "example",
            assume_role_policy=Token.as_string(assume_role.json),
            name="example"
        )
        aws_iam_role_policy_attachment_example = IamRolePolicyAttachment(self, "example_2",
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup",
            role=example.name
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_iam_role_policy_attachment_example.override_logical_id("example")
        aws_backup_selection_example = BackupSelection(self, "example_3",
            iam_role_arn=example.arn,
            name=name,
            plan_id=plan_id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_backup_selection_example.override_logical_id("example")
```

### Selecting Backups By Tag

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.backup_selection import BackupSelection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BackupSelection(self, "example",
            iam_role_arn=Token.as_string(aws_iam_role_example.arn),
            name="tf_example_backup_selection",
            plan_id=Token.as_string(aws_backup_plan_example.id),
            selection_tag=[BackupSelectionSelectionTag(
                key="foo",
                type="STRINGEQUALS",
                value="bar"
            )
            ]
        )
```

### Selecting Backups By Conditions

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.backup_selection import BackupSelection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BackupSelection(self, "example",
            condition=[BackupSelectionCondition(
                string_equals=[BackupSelectionConditionStringEquals(
                    key="aws:ResourceTag/Component",
                    value="rds"
                )
                ],
                string_like=[BackupSelectionConditionStringLike(
                    key="aws:ResourceTag/Application",
                    value="app*"
                )
                ],
                string_not_equals=[BackupSelectionConditionStringNotEquals(
                    key="aws:ResourceTag/Backup",
                    value="false"
                )
                ],
                string_not_like=[BackupSelectionConditionStringNotLike(
                    key="aws:ResourceTag/Environment",
                    value="test*"
                )
                ]
            )
            ],
            iam_role_arn=Token.as_string(aws_iam_role_example.arn),
            name="tf_example_backup_selection",
            plan_id=Token.as_string(aws_backup_plan_example.id),
            resources=["*"]
        )
```

### Selecting Backups By Resource

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.backup_selection import BackupSelection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BackupSelection(self, "example",
            iam_role_arn=Token.as_string(aws_iam_role_example.arn),
            name="tf_example_backup_selection",
            plan_id=Token.as_string(aws_backup_plan_example.id),
            resources=[
                Token.as_string(aws_db_instance_example.arn),
                Token.as_string(aws_ebs_volume_example.arn),
                Token.as_string(aws_efs_file_system_example.arn)
            ]
        )
```

### Selecting Backups By Not Resource

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.backup_selection import BackupSelection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BackupSelection(self, "example",
            iam_role_arn=Token.as_string(aws_iam_role_example.arn),
            name="tf_example_backup_selection",
            not_resources=[
                Token.as_string(aws_db_instance_example.arn),
                Token.as_string(aws_ebs_volume_example.arn),
                Token.as_string(aws_efs_file_system_example.arn)
            ],
            plan_id=Token.as_string(aws_backup_plan_example.id)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The display name of a resource selection document.
* `plan_id` - (Required) The backup plan ID to be associated with the selection of resources.
* `iam_role_arn` - (Required) The ARN of the IAM role that AWS Backup uses to authenticate when restoring and backing up the target resource. See the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/access-control.html#managed-policies) for additional information about using AWS managed policies or creating custom policies attached to the IAM role.
* `selection_tag` - (Optional) Tag-based conditions used to specify a set of resources to assign to a backup plan. See [below](#selection_tag) for details.
* `condition` - (Optional) Condition-based filters used to specify sets of resources for a backup plan. See [below](#condition) for details.
* `resources` - (Optional) An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to assign to a backup plan.
* `not_resources` - (Optional) An array of strings that either contain Amazon Resource Names (ARNs) or match patterns of resources to exclude from a backup plan.

### selection_tag

The `selection_tag` configuration block supports the following attributes:

* `type` - (Required) An operation, such as `STRINGEQUALS`, that is applied to the key-value pair used to filter resources in a selection.
* `key` - (Required) Key for the filter.
* `value` - (Required) Value for the filter.

### condition

The `condition` configuration block supports the following attributes:

* `string_equals` - (Optional) Filters the values of your tagged resources for only those resources that you tagged with the same value. Also called "exact matching". See [below](#string_equals) for details.
* `string_not_equals` - (Optional) Filters the values of your tagged resources for only those resources that you tagged that do not have the same value. Also called "negated matching". See [below](#string_not_equals) for details.
* `string_like` - (Optional) Filters the values of your tagged resources for matching tag values with the use of a wildcard character (`*`) anywhere in the string. For example, `prod*` or `*rod*` matches the tag value `production`. See [below](#string_like) for details.
* `string_not_like` - (Optional) Filters the values of your tagged resources for non-matching tag values with the use of a wildcard character (`*`) anywhere in the string. See [below](#string_not_like) for details.

### string_equals

The `string_equals` configuration block supports the following attributes:

* `key` - (Required) Key for the filter.
* `value` - (Required) Value for the filter.

### string_not_equals

The `string_not_equals` configuration block supports the following attributes:

* `key` - (Required) Key for the filter.
* `value` - (Required) Value for the filter.

### string_like

The `string_like` configuration block supports the following attributes:

* `key` - (Required) Key for the filter.
* `value` - (Required) Value for the filter.

### string_not_like

The `string_not_like` configuration block supports the following attributes:

* `key` - (Required) Key for the filter.
* `value` - (Required)  Value for the filter.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Backup Selection identifier

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Backup selection using the role plan_id and id separated by `|`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.backup_selection import BackupSelection
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        BackupSelection.generate_config_for_import(self, "example", "plan-id|selection-id")
```

Using `terraform import`, import Backup selection using the role plan_id and id separated by `|`. For example:

```console
% terraform import aws_backup_selection.example plan-id|selection-id
```

<!-- cache-key: cdktf-0.20.8 input-789c3fb7dc6762eb5d7f23905bec030e99cfd540546ecac689a7898340e7778a -->