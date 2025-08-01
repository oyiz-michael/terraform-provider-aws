---
subcategory: "SSO Admin"
layout: "aws"
page_title: "AWS: aws_ssoadmin_permission_set_inline_policy"
description: |-
  Manages an IAM inline policy for a Single Sign-On (SSO) Permission Set
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ssoadmin_permission_set_inline_policy

Provides an IAM inline policy for a Single Sign-On (SSO) Permission Set resource

~> **NOTE:** AWS Single Sign-On (SSO) only supports one IAM inline policy per [`aws_ssoadmin_permission_set`](ssoadmin_permission_set.html) resource.
Creating or updating this resource will automatically [Provision the Permission Set](https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ProvisionPermissionSet.html) to apply the corresponding updates to all assigned accounts.

~> **NOTE:** We suggest using [`jsonencode()`](https://developer.hashicorp.com/terraform/language/functions/jsonencode) or [`aws_iam_policy_document`](/docs/providers/aws/d/iam_policy_document.html) when assigning a value to `inline_policy`. They seamlessly translate Terraform language into JSON, enabling you to maintain consistency within your configuration without the need for context switches. Also, you can sidestep potential complications arising from formatting discrepancies, whitespace inconsistencies, and other nuances inherent to JSON.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_iam_policy_document import DataAwsIamPolicyDocument
from imports.aws.data_aws_ssoadmin_instances import DataAwsSsoadminInstances
from imports.aws.ssoadmin_permission_set import SsoadminPermissionSet
from imports.aws.ssoadmin_permission_set_inline_policy import SsoadminPermissionSetInlinePolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsIamPolicyDocument(self, "example",
            statement=[DataAwsIamPolicyDocumentStatement(
                actions=["s3:ListAllMyBuckets", "s3:GetBucketLocation"],
                resources=["arn:aws:s3:::*"],
                sid="1"
            )
            ]
        )
        data_aws_ssoadmin_instances_example = DataAwsSsoadminInstances(self, "example_1")
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_ssoadmin_instances_example.override_logical_id("example")
        aws_ssoadmin_permission_set_example = SsoadminPermissionSet(self, "example_2",
            instance_arn=Token.as_string(
                Fn.lookup_nested(Fn.tolist(data_aws_ssoadmin_instances_example.arns), ["0"
                ])),
            name="Example"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_ssoadmin_permission_set_example.override_logical_id("example")
        aws_ssoadmin_permission_set_inline_policy_example =
        SsoadminPermissionSetInlinePolicy(self, "example_3",
            inline_policy=Token.as_string(example.json),
            instance_arn=Token.as_string(
                Fn.lookup_nested(Fn.tolist(data_aws_ssoadmin_instances_example.arns), ["0"
                ])),
            permission_set_arn=Token.as_string(aws_ssoadmin_permission_set_example.arn)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_ssoadmin_permission_set_inline_policy_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `inline_policy` - (Required) The IAM inline policy to attach to a Permission Set.
* `instance_arn` - (Required, Forces new resource) The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.
* `permission_set_arn` - (Required, Forces new resource) The Amazon Resource Name (ARN) of the Permission Set.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The Amazon Resource Names (ARNs) of the Permission Set and SSO Instance, separated by a comma (`,`).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SSO Permission Set Inline Policies using the `permission_set_arn` and `instance_arn` separated by a comma (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ssoadmin_permission_set_inline_policy import SsoadminPermissionSetInlinePolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SsoadminPermissionSetInlinePolicy.generate_config_for_import(self, "example", "arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72")
```

Using `terraform import`, import SSO Permission Set Inline Policies using the `permission_set_arn` and `instance_arn` separated by a comma (`,`). For example:

```console
% terraform import aws_ssoadmin_permission_set_inline_policy.example arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
```

<!-- cache-key: cdktf-0.20.8 input-9d568eefae6ec48ca027ed9251548f57b9623feef95d16d3b487cfbdf3f06ef3 -->