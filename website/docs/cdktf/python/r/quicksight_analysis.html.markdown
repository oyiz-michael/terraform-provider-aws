---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_analysis"
description: |-
  Manages a QuickSight Analysis.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_quicksight_analysis

Resource for managing a QuickSight Analysis.

## Example Usage

### From Source Template

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.quicksight_analysis import QuicksightAnalysis
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        QuicksightAnalysis(self, "example",
            analysis_id="example-id",
            name="example-name",
            source_entity=QuicksightAnalysisSourceEntity(
                source_template=QuicksightAnalysisSourceEntitySourceTemplate(
                    arn=source.arn,
                    data_set_references=[QuicksightAnalysisSourceEntitySourceTemplateDataSetReferences(
                        data_set_arn=dataset.arn,
                        data_set_placeholder="1"
                    )
                    ]
                )
            )
        )
```

### With Definition

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.quicksight_analysis import QuicksightAnalysis
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        QuicksightAnalysis(self, "example",
            analysis_id="example-id",
            definition={
                "data_set_identifiers_declarations": [{
                    "data_set_arn": dataset.arn,
                    "identifier": "1"
                }
                ],
                "sheets": [{
                    "sheet_id": "Example1",
                    "title": "Example",
                    "visuals": [{
                        "line_chart_visual": {
                            "chart_configuration": {
                                "field_wells": {
                                    "line_chart_aggregated_field_wells": {
                                        "category": [{
                                            "categorical_dimension_field": {
                                                "column": {
                                                    "column_name": "Column1",
                                                    "data_set_identifier": "1"
                                                },
                                                "field_id": "1"
                                            }
                                        }
                                        ],
                                        "values": [{
                                            "categorical_measure_field": {
                                                "aggregation_function": "COUNT",
                                                "column": {
                                                    "column_name": "Column1",
                                                    "data_set_identifier": "1"
                                                },
                                                "field_id": "2"
                                            }
                                        }
                                        ]
                                    }
                                }
                            },
                            "title": {
                                "format_text": {
                                    "plain_text": "Line Chart Example"
                                }
                            },
                            "visual_id": "LineChart"
                        }
                    }
                    ]
                }
                ]
            },
            name="example-name"
        )
```

## Argument Reference

The following arguments are required:

* `analysis_id` - (Required, Forces new resource) Identifier for the analysis.
* `name` - (Required) Display name for the analysis.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `aws_account_id` - (Optional, Forces new resource) AWS account ID.
* `definition` - (Optional) A detailed analysis definition. Only one of `definition` or `source_entity` should be configured. See [definition](#definition).
* `parameters` - (Optional) The parameters for the creation of the analysis, which you want to use to override the default settings. An analysis can have any type of parameters, and some parameters might accept multiple values. See [parameters](#parameters).
* `permissions` - (Optional) A set of resource permissions on the analysis. Maximum of 64 items. See [permissions](#permissions).
* `recovery_window_in_days` - (Optional) A value that specifies the number of days that Amazon QuickSight waits before it deletes the analysis. Use `0` to force deletion without recovery. Minimum value of `7`. Maximum value of `30`. Default to `30`.
* `source_entity` - (Optional) The entity that you are using as a source when you create the analysis (template). Only one of `definition` or `source_entity` should be configured. See [source_entity](#source_entity).
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `theme_arn` - (Optional) The Amazon Resource Name (ARN) of the theme that is being used for this analysis. The theme ARN must exist in the same AWS account where you create the analysis.

### permissions

* `actions` - (Required) List of IAM actions to grant or revoke permissions on.
* `principal` - (Required) ARN of the principal. See the [ResourcePermission documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ResourcePermission.html) for the applicable ARN values.

### source_entity

* `source_template` - (Optional) The source template. See [source_template](#source_template).

### source_template

* `arn` - (Required) The Amazon Resource Name (ARN) of the resource.
* `data_set_references` - (Required) List of dataset references. See [data_set_references](#data_set_references).

### data_set_references

* `data_set_arn` - (Required) Dataset Amazon Resource Name (ARN).
* `data_set_placeholder` - (Required) Dataset placeholder.

### parameters

* `date_time_parameters` - (Optional) A list of parameters that have a data type of date-time. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DateTimeParameter.html).
* `decimal_parameters` - (Optional) A list of parameters that have a data type of decimal. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DecimalParameter.html).
* `integer_parameters` - (Optional) A list of parameters that have a data type of integer. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_IntegerParameter.html).
* `string_parameters` - (Optional) A list of parameters that have a data type of string. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_StringParameter.html).

### definition

* `data_set_identifiers_declarations` - (Required) A list dataset identifier declarations. With this mapping,you can use dataset identifiers instead of dataset Amazon Resource Names (ARNs) throughout the analysis sub-structures. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSetIdentifierDeclaration.html).
* `analysis_defaults` - (Optional) The configuration for default analysis settings. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_AnalysisDefaults.html).
* `calculated_fields` - (Optional) A list of calculated field definitions for the analysis. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CalculatedField.html).
* `column_configurations` - (Optional) A list of analysis-level column configurations. Column configurations are used to set default formatting for a column that's used throughout an analysis. See [AWS API Documentation for complete description](ttps://docs.aws.amazon.com/quicksight/latest/APIReference/API_ColumnConfiguration.html).
* `filter_groups` - (Optional) A list of filter definitions for an analysis. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_FilterGroup.html). For more information, see [Filtering Data](https://docs.aws.amazon.com/quicksight/latest/user/filtering-visual-data.html) in Amazon QuickSight User Guide.
* `parameters_declarations` - (Optional) A list of parameter declarations for an analysis. Parameters are named variables that can transfer a value for use by an action or an object. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ParameterDeclaration.html). For more information, see [Parameters in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the Amazon QuickSight User Guide.
* `sheets` - (Optional) A list of sheet definitions for an analysis. See [AWS API Documentation for complete description](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_SheetDefinition.html).

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the analysis.
* `created_time` - The time that the analysis was created.
* `id` - A comma-delimited string joining AWS account ID and analysis ID.
* `last_updated_time` - The time that the analysis was last updated.
* `status` - The analysis creation status.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `5m`)
* `update` - (Default `5m`)
* `delete` - (Default `5m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import a QuickSight Analysis using the AWS account ID and analysis ID separated by a comma (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.quicksight_analysis import QuicksightAnalysis
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        QuicksightAnalysis.generate_config_for_import(self, "example", "123456789012,example-id")
```

Using `terraform import`, import a QuickSight Analysis using the AWS account ID and analysis ID separated by a comma (`,`). For example:

```console
% terraform import aws_quicksight_analysis.example 123456789012,example-id
```

<!-- cache-key: cdktf-0.20.8 input-a0ffd8f2c1f465fa39defe05d4356307483f42d4c37162ce146d095b6f0c576a -->