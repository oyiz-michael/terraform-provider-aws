---
subcategory: "Audit Manager"
layout: "aws"
page_title: "AWS: aws_auditmanager_assessment_delegation"
description: |-
  Terraform resource for managing an AWS Audit Manager Assessment Delegation.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_auditmanager_assessment_delegation

Terraform resource for managing an AWS Audit Manager Assessment Delegation.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.auditmanager_assessment_delegation import AuditmanagerAssessmentDelegation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AuditmanagerAssessmentDelegation(self, "example",
            assessment_id=Token.as_string(aws_auditmanager_assessment_example.id),
            control_set_id="example",
            role_arn=Token.as_string(aws_iam_role_example.arn),
            role_type="RESOURCE_OWNER"
        )
```

## Argument Reference

The following arguments are required:

* `assessment_id` - (Required) Identifier for the assessment.
* `control_set_id` - (Required) Assessment control set name. This value is the control set name used during assessment creation (not the AWS-generated ID). The `_id` suffix on this attribute has been preserved to be consistent with the underlying AWS API.
* `role_arn` - (Required) Amazon Resource Name (ARN) of the IAM role.
* `role_type` - (Required) Type of customer persona. For assessment delegation, type must always be `RESOURCE_OWNER`.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `comment` - (Optional) Comment describing the delegation request.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `delegation_id` - Unique identifier for the delegation.
* `id` - Unique identifier for the resource. This is a comma-separated string containing `assessment_id`, `role_arn`, and `control_set_id`.
* `status` - Status of the delegation.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Audit Manager Assessment Delegation using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.auditmanager_assessment_delegation import AuditmanagerAssessmentDelegation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AuditmanagerAssessmentDelegation.generate_config_for_import(self, "example", "abcdef-123456,arn:aws:iam::123456789012:role/example,example")
```

Using `terraform import`, import Audit Manager Assessment Delegation using the `id`. For example:

```console
% terraform import aws_auditmanager_assessment_delegation.example abcdef-123456,arn:aws:iam::123456789012:role/example,example
```

<!-- cache-key: cdktf-0.20.8 input-a79c875ad4fd03d89be9d5dd9ec74ff4deed50ddd73a04bd26e198f2d314cd32 -->