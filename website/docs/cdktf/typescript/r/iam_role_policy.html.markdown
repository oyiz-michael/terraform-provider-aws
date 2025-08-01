---
subcategory: "IAM (Identity & Access Management)"
layout: "aws"
page_title: "AWS: aws_iam_role_policy"
description: |-
  Provides an IAM role policy.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_iam_role_policy

Provides an IAM role inline policy.

~> **NOTE:** For a given role, this resource is incompatible with using the [`aws_iam_role` resource](/docs/providers/aws/r/iam_role.html) `inlinePolicy` argument. When using that argument and this resource, both will attempt to manage the role's inline policies and Terraform will show a permanent difference.

~> **NOTE:** We suggest using [`jsonencode()`](https://developer.hashicorp.com/terraform/language/functions/jsonencode) or [`aws_iam_policy_document`](/docs/providers/aws/d/iam_policy_document.html) when assigning a value to `policy`. They seamlessly translate Terraform language into JSON, enabling you to maintain consistency within your configuration without the need for context switches. Also, you can sidestep potential complications arising from formatting discrepancies, whitespace inconsistencies, and other nuances inherent to JSON.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { IamRole } from "./.gen/providers/aws/iam-role";
import { IamRolePolicy } from "./.gen/providers/aws/iam-role-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const testRole = new IamRole(this, "test_role", {
      assumeRolePolicy: Token.asString(
        Fn.jsonencode({
          Statement: [
            {
              Action: "sts:AssumeRole",
              Effect: "Allow",
              Principal: {
                Service: "ec2.amazonaws.com",
              },
              Sid: "",
            },
          ],
          Version: "2012-10-17",
        })
      ),
      name: "test_role",
    });
    new IamRolePolicy(this, "test_policy", {
      name: "test_policy",
      policy: Token.asString(
        Fn.jsonencode({
          Statement: [
            {
              Action: ["ec2:Describe*"],
              Effect: "Allow",
              Resource: "*",
            },
          ],
          Version: "2012-10-17",
        })
      ),
      role: testRole.id,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `name` - (Optional) The name of the role policy.
  If omitted, Terraform will assign a random, unique name.
* `namePrefix` - (Optional) Creates a unique name beginning with the specified prefix.
  Conflicts with `name`.
* `policy` - (Required) The inline policy document.
  This is a JSON formatted string.
  For more information about building IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://learn.hashicorp.com/terraform/aws/iam-policy)
* `role` - (Required) The name of the IAM role to attach to the policy.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import IAM Role Policies using the `role_name:role_policy_name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { IamRolePolicy } from "./.gen/providers/aws/iam-role-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    IamRolePolicy.generateConfigForImport(
      this,
      "mypolicy",
      "role_of_mypolicy_name:mypolicy_name"
    );
  }
}

```

Using `terraform import`, import IAM Role Policies using the `role_name:role_policy_name`. For example:

```console
% terraform import aws_iam_role_policy.mypolicy role_of_mypolicy_name:mypolicy_name
```

<!-- cache-key: cdktf-0.20.8 input-d20fb451cbf63b9a585b625c29c9fea81468054b3da6053ef0d35b42807b148b -->