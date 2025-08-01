---
subcategory: "IAM Access Analyzer"
layout: "aws"
page_title: "AWS: aws_accessanalyzer_analyzer"
description: |-
  Manages an Access Analyzer Analyzer
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_accessanalyzer_analyzer

Manages an Access Analyzer Analyzer. More information can be found in the [Access Analyzer User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html).

## Example Usage

### Account Analyzer

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AccessanalyzerAnalyzer } from "./.gen/providers/aws/accessanalyzer-analyzer";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new AccessanalyzerAnalyzer(this, "example", {
      analyzerName: "example",
    });
  }
}

```

### Organization Analyzer

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AccessanalyzerAnalyzer } from "./.gen/providers/aws/accessanalyzer-analyzer";
import { OrganizationsOrganization } from "./.gen/providers/aws/organizations-organization";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new OrganizationsOrganization(this, "example", {
      awsServiceAccessPrincipals: ["access-analyzer.amazonaws.com"],
    });
    const awsAccessanalyzerAnalyzerExample = new AccessanalyzerAnalyzer(
      this,
      "example_1",
      {
        analyzerName: "example",
        dependsOn: [example],
        type: "ORGANIZATION",
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsAccessanalyzerAnalyzerExample.overrideLogicalId("example");
  }
}

```

### Organization Unused Access Analyzer With Analysis Rule

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AccessanalyzerAnalyzer } from "./.gen/providers/aws/accessanalyzer-analyzer";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new AccessanalyzerAnalyzer(this, "example", {
      analyzerName: "example",
      configuration: {
        unusedAccess: {
          analysisRule: {
            exclusion: [
              {
                accountIds: ["123456789012", "234567890123"],
              },
              {
                resourceTags: [
                  {
                    key1: "value1",
                  },
                  {
                    key2: "value2",
                  },
                ],
              },
            ],
          },
          unusedAccessAge: 180,
        },
      },
      type: "ORGANIZATION_UNUSED_ACCESS",
    });
  }
}

```

### Account Internal Access Analyzer by Resource Types

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AccessanalyzerAnalyzer } from "./.gen/providers/aws/accessanalyzer-analyzer";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new AccessanalyzerAnalyzer(this, "test", {
      analyzerName: "example",
      configuration: {
        internalAccess: {
          analysisRule: {
            inclusion: [
              {
                resourceTypes: [
                  "AWS::S3::Bucket",
                  "AWS::RDS::DBSnapshot",
                  "AWS::DynamoDB::Table",
                ],
              },
            ],
          },
        },
      },
      type: "ORGANIZATION_INTERNAL_ACCESS",
    });
  }
}

```

### Organization Internal Access Analyzer by Account ID and Resource ARN

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AccessanalyzerAnalyzer } from "./.gen/providers/aws/accessanalyzer-analyzer";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new AccessanalyzerAnalyzer(this, "test", {
      analyzerName: "example",
      configuration: {
        internalAccess: {
          analysisRule: {
            inclusion: [
              {
                accountIds: ["123456789012"],
                resourceArns: ["arn:aws:s3:::my-example-bucket"],
              },
            ],
          },
        },
      },
      type: "ORGANIZATION_INTERNAL_ACCESS",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `analyzerName` - (Required) Name of the Analyzer.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `configuration` - (Optional) A block that specifies the configuration of the analyzer. See [`configuration` Block](#configuration-block) for details.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `type` - (Optional) Type that represents the zone of trust or scope for the analyzer. Valid values are `ACCOUNT`, `ACCOUNT_INTERNAL_ACCESS`, `ACCOUNT_UNUSED_ACCESS`, `ORGANIZATION`, `ORGANIZATION_INTERNAL_ACCESS`, `ORGANIZATION_UNUSED_ACCESS`. Defaults to `ACCOUNT`.

### `configuration` Block

The `configuration` configuration block supports the following arguments:

* `internalAccess` - (Optional) Specifies the configuration of an internal access analyzer for an AWS organization or account. This configuration determines how the analyzer evaluates access within your AWS environment. See [`internalAccess` Block](#internal_access-block) for details.
* `unusedAccess` - (Optional) Specifies the configuration of an unused access analyzer for an AWS organization or account. See [`unusedAccess` Block](#unused_access-block) for details.

### `internalAccess` Block

The `internalAccess` configuration block supports the following arguments:

* `analysisRule` - (Optional) Information about analysis rules for the internal access analyzer. These rules determine which resources and access patterns will be analyzed. See [`analysisRule` Block for Internal Access Analyzer](#analysis_rule-block-for-internal-access-analyzer) for details.

### `analysisRule` Block for Internal Access Analyzer

The `analysisRule` configuration block for internal access analyzer supports the following arguments:

* `inclusion` - (Optional) List of rules for the internal access analyzer containing criteria to include in analysis. Only resources that meet the rule criteria will generate findings. See [`inclusion` Block](#inclusion-block) for details.

### `inclusion` Block

The `inclusion` configuration block supports the following arguments:

* `accountIds` - (Optional) List of AWS account IDs to apply to the internal access analysis rule criteria. Account IDs can only be applied to the analysis rule criteria for organization-level analyzers.
* `resourceArns` - (Optional) List of resource ARNs to apply to the internal access analysis rule criteria. The analyzer will only generate findings for resources that match these ARNs.
* `resourceTypes` - (Optional) List of resource types to apply to the internal access analysis rule criteria. The analyzer will only generate findings for resources of these types. Refer to [InternalAccessAnalysisRuleCriteria](https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_InternalAccessAnalysisRuleCriteria.html) in the AWS IAM Access Analyzer API Reference for valid values.

### `unusedAccess` Block

The `unusedAccess` configuration block supports the following arguments:

* `unusedAccessAge` - (Optional) Specified access age in days for which to generate findings for unused access.
* `analysisRule` - (Optional) Information about analysis rules for the analyzer. Analysis rules determine which entities will generate findings based on the criteria you define when you create the rule. See [`analysisRule` Block for Unused Access Analyzer](#analysis_rule-block-for-unused-access-analyzer) for details.

### `analysisRule` Block for Unused Access Analyzer

The `analysisRule` configuration block for unused access analyzer supports the following arguments:

* `exclusion` - (Optional) List of rules for the analyzer containing criteria to exclude from analysis. Entities that meet the rule criteria will not generate findings. See [`exclusion` Block](#exclusion-block) for details.

### `exclusion` Block

The `exclusion` configuration block supports the following arguments:

* `accountIds` - (Optional) List of AWS account IDs to apply to the analysis rule criteria. The accounts cannot include the organization analyzer owner account. Account IDs can only be applied to the analysis rule criteria for organization-level analyzers.
* `resourceTags` - (Optional) List of key-value pairs for resource tags to exclude from the analysis.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Analyzer.
* `id` - Name of the analyzer.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Access Analyzer Analyzers using the `analyzerName`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AccessanalyzerAnalyzer } from "./.gen/providers/aws/accessanalyzer-analyzer";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    AccessanalyzerAnalyzer.generateConfigForImport(this, "example", "example");
  }
}

```

Using `terraform import`, import Access Analyzer Analyzers using the `analyzerName`. For example:

```console
% terraform import aws_accessanalyzer_analyzer.example example
```

<!-- cache-key: cdktf-0.20.8 input-8e411e6881f8db0084679fad8b37a5364a7999eee5f9248192be7d773f1f3e35 -->