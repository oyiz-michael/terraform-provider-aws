---
subcategory: "SSO Admin"
layout: "aws"
page_title: "AWS: aws_ssoadmin_application_assignments"
description: |-
  Terraform data source for managing AWS SSO Admin Application Assignments.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ssoadmin_application_assignments

Terraform data source for managing AWS SSO Admin Application Assignments.

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
from imports.aws.data_aws_ssoadmin_application_assignments import DataAwsSsoadminApplicationAssignments
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsSsoadminApplicationAssignments(self, "example",
            application_arn=Token.as_string(aws_ssoadmin_application_example.arn)
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `application_arn` - (Required) ARN of the application.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `application_assignments` - List of principals assigned to the application. See the [`application_assignments` attribute reference](#application_assignments-attribute-reference) below.

### `application_assignments` Attribute Reference

* `application_arn` - ARN of the application.
* `principal_id` - An identifier for an object in IAM Identity Center, such as a user or group.
* `principal_type` - Entity type for which the assignment will be created. Valid values are `USER` or `GROUP`.

<!-- cache-key: cdktf-0.20.8 input-2746ade34b1a6440566ec23228a499a3cea6c549aab6e20ea2724715c8d25855 -->