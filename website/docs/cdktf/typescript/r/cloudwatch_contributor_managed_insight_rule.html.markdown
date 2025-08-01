---
subcategory: "CloudWatch"
layout: "aws"
page_title: "AWS: aws_cloudwatch_contributor_managed_insight_rule"
description: |-
  Terraform resource for managing an AWS CloudWatch Contributor Managed Insight Rule.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_cloudwatch_contributor_managed_insight_rule

Terraform resource for managing an AWS CloudWatch Contributor Managed Insight Rule.

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
import { CloudwatchContributorManagedInsightRule } from "./.gen/providers/aws/cloudwatch-contributor-managed-insight-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new CloudwatchContributorManagedInsightRule(this, "example", {
      resourceArn: test.arn,
      rule_state: "DISABLED",
      templateName: "VpcEndpointService-BytesByEndpointId-v1",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `resourceArn` - (Required) ARN of an Amazon Web Services resource that has managed Contributor Insights rules.
* `templateName` - (Required) Template name for the managed Contributor Insights rule, as returned by ListManagedInsightRules.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `ruleState` - (Optional) State of the rule. Valid values are `ENABLED` and `DISABLED`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Contributor Managed Insight Rule.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import CloudWatch Contributor Managed Insight Rule using the `resourceArn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { CloudwatchContributorManagedInsightRule } from "./.gen/providers/aws/cloudwatch-contributor-managed-insight-rule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    CloudwatchContributorManagedInsightRule.generateConfigForImport(
      this,
      "example",
      "contributor_managed_insight_rule-id-12345678"
    );
  }
}

```

Using `terraform import`, import CloudWatch Contributor Managed Insight Rule using the `resourceArn`. For example:

```console
% terraform import aws_cloudwatch_contributor_managed_insight_rule.example contributor_managed_insight_rule-id-12345678
```

<!-- cache-key: cdktf-0.20.8 input-dd8594cfcddda462fedfe0005c0797493ddbbbd189d12113d239e7550d0ca4a8 -->