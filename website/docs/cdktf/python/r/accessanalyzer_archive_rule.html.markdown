---
subcategory: "IAM Access Analyzer"
layout: "aws"
page_title: "AWS: aws_accessanalyzer_archive_rule"
description: |-
  Terraform resource for managing an AWS AccessAnalyzer Archive Rule.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_accessanalyzer_archive_rule

Terraform resource for managing an AWS AccessAnalyzer Archive Rule.

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
from imports.aws.accessanalyzer_archive_rule import AccessanalyzerArchiveRule
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AccessanalyzerArchiveRule(self, "example",
            analyzer_name="example-analyzer",
            filter=[AccessanalyzerArchiveRuleFilter(
                criteria="condition.aws:UserId",
                eq=["userid"]
            ), AccessanalyzerArchiveRuleFilter(
                criteria="error",
                exists=Token.as_string(True)
            ), AccessanalyzerArchiveRuleFilter(
                criteria="isPublic",
                eq=["false"]
            )
            ],
            rule_name="example-rule"
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `analyzer_name` - (Required) Analyzer name.
* `filter` - (Required) Filter criteria for the archive rule. See [Filter](#filter) for more details.
* `rule_name` - (Required) Rule name.

### Filter

**Note** One comparator must be included with each filter.

* `criteria` - (Required) Filter criteria.
* `contains` - (Optional) Contains comparator.
* `eq` - (Optional) Equals comparator.
* `exists` - (Optional) Boolean comparator.
* `neq` - (Optional) Not Equals comparator.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Resource ID in the format: `analyzer_name/rule_name`.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import AccessAnalyzer ArchiveRule using the `analyzer_name/rule_name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.accessanalyzer_archive_rule import AccessanalyzerArchiveRule
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AccessanalyzerArchiveRule.generate_config_for_import(self, "example", "example-analyzer/example-rule")
```

Using `terraform import`, import AccessAnalyzer ArchiveRule using the `analyzer_name/rule_name`. For example:

```console
% terraform import aws_accessanalyzer_archive_rule.example example-analyzer/example-rule
```

<!-- cache-key: cdktf-0.20.8 input-c77620a5d28ca89b026b70785648f355942dd3687d273a07273452ee0b4832ac -->