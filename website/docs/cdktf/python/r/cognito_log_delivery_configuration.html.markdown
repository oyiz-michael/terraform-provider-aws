---
subcategory: "Cognito IDP (Identity Provider)"
layout: "aws"
page_title: "AWS: aws_cognito_log_delivery_configuration"
description: |-
  Manages an AWS Cognito IDP (Identity Provider) Log Delivery Configuration.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_cognito_log_delivery_configuration

Manages an AWS Cognito IDP (Identity Provider) Log Delivery Configuration.

## Example Usage

### Basic Usage with CloudWatch Logs

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_log_group import CloudwatchLogGroup
from imports.aws.cognito_log_delivery_configuration import CognitoLogDeliveryConfiguration
from imports.aws.cognito_user_pool import CognitoUserPool
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = CloudwatchLogGroup(self, "example",
            name="example"
        )
        aws_cognito_user_pool_example = CognitoUserPool(self, "example_1",
            name="example"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_cognito_user_pool_example.override_logical_id("example")
        aws_cognito_log_delivery_configuration_example =
        CognitoLogDeliveryConfiguration(self, "example_2",
            log_configurations=[CognitoLogDeliveryConfigurationLogConfigurations(
                cloud_watch_logs_configuration=[CognitoLogDeliveryConfigurationLogConfigurationsCloudWatchLogsConfiguration(
                    log_group_arn=example.arn
                )
                ],
                event_source="userNotification",
                log_level="ERROR"
            )
            ],
            user_pool_id=Token.as_string(aws_cognito_user_pool_example.id)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_cognito_log_delivery_configuration_example.override_logical_id("example")
```

### Multiple Log Configurations with Different Destinations

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cloudwatch_log_group import CloudwatchLogGroup
from imports.aws.cognito_log_delivery_configuration import CognitoLogDeliveryConfiguration
from imports.aws.cognito_user_pool import CognitoUserPool
from imports.aws.iam_role import IamRole
from imports.aws.iam_role_policy import IamRolePolicy
from imports.aws.kinesis_firehose_delivery_stream import KinesisFirehoseDeliveryStream
from imports.aws.s3_bucket import S3Bucket
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = CloudwatchLogGroup(self, "example",
            name="example"
        )
        aws_cognito_user_pool_example = CognitoUserPool(self, "example_1",
            name="example"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_cognito_user_pool_example.override_logical_id("example")
        firehose = IamRole(self, "firehose",
            assume_role_policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "firehose.amazonaws.com"
                        }
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            name="firehose-role"
        )
        aws_s3_bucket_example = S3Bucket(self, "example_3",
            bucket="example-bucket",
            force_destroy=True
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_bucket_example.override_logical_id("example")
        aws_iam_role_policy_firehose = IamRolePolicy(self, "firehose_4",
            name="firehose-policy",
            policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": ["s3:AbortMultipartUpload", "s3:GetBucketLocation", "s3:GetObject", "s3:ListBucket", "s3:ListBucketMultipartUploads", "s3:PutObject"
                        ],
                        "Effect": "Allow",
                        "Resource": [aws_s3_bucket_example.arn, "${" + aws_s3_bucket_example.arn + "}/*"
                        ]
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            role=firehose.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_iam_role_policy_firehose.override_logical_id("firehose")
        aws_kinesis_firehose_delivery_stream_example =
        KinesisFirehoseDeliveryStream(self, "example_5",
            destination="extended_s3",
            extended_s3_configuration=KinesisFirehoseDeliveryStreamExtendedS3Configuration(
                bucket_arn=Token.as_string(aws_s3_bucket_example.arn),
                role_arn=firehose.arn
            ),
            name="example-stream"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_kinesis_firehose_delivery_stream_example.override_logical_id("example")
        aws_cognito_log_delivery_configuration_example =
        CognitoLogDeliveryConfiguration(self, "example_6",
            log_configurations=[CognitoLogDeliveryConfigurationLogConfigurations(
                cloud_watch_logs_configuration=[CognitoLogDeliveryConfigurationLogConfigurationsCloudWatchLogsConfiguration(
                    log_group_arn=example.arn
                )
                ],
                event_source="userNotification",
                log_level="INFO"
            ), CognitoLogDeliveryConfigurationLogConfigurations(
                event_source="userAuthEvents",
                firehose_configuration=[CognitoLogDeliveryConfigurationLogConfigurationsFirehoseConfiguration(
                    stream_arn=Token.as_string(aws_kinesis_firehose_delivery_stream_example.arn)
                )
                ],
                log_level="ERROR"
            )
            ],
            user_pool_id=Token.as_string(aws_cognito_user_pool_example.id)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_cognito_log_delivery_configuration_example.override_logical_id("example")
```

### S3 Configuration

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cognito_log_delivery_configuration import CognitoLogDeliveryConfiguration
from imports.aws.cognito_user_pool import CognitoUserPool
from imports.aws.s3_bucket import S3Bucket
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = CognitoUserPool(self, "example",
            name="example"
        )
        aws_s3_bucket_example = S3Bucket(self, "example_1",
            bucket="example-bucket",
            force_destroy=True
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_bucket_example.override_logical_id("example")
        aws_cognito_log_delivery_configuration_example =
        CognitoLogDeliveryConfiguration(self, "example_2",
            log_configurations=[CognitoLogDeliveryConfigurationLogConfigurations(
                event_source="userNotification",
                log_level="ERROR",
                s3_configuration=[CognitoLogDeliveryConfigurationLogConfigurationsS3Configuration(
                    bucket_arn=Token.as_string(aws_s3_bucket_example.arn)
                )
                ]
            )
            ],
            user_pool_id=example.id
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_cognito_log_delivery_configuration_example.override_logical_id("example")
```

## Argument Reference

The following arguments are required:

* `user_pool_id` - (Required) The ID of the user pool for which to configure log delivery.

The following arguments are optional:

* `log_configurations` - (Optional) Configuration block for log delivery. At least one configuration block is required. See [Log Configurations](#log-configurations) below.
* `region` - (Optional) The AWS region.

### Log Configurations

The `log_configurations` block supports the following:

* `event_source` - (Required) The event source to configure logging for. Valid values are `userNotification` and `userAuthEvents`.
* `log_level` - (Required) The log level to set for the event source. Valid values are `ERROR` and `INFO`.
* `cloud_watch_logs_configuration` - (Optional) Configuration for CloudWatch Logs delivery. See [CloudWatch Logs Configuration](#cloudwatch-logs-configuration) below.
* `firehose_configuration` - (Optional) Configuration for Kinesis Data Firehose delivery. See [Firehose Configuration](#firehose-configuration) below.
* `s3_configuration` - (Optional) Configuration for S3 delivery. See [S3 Configuration](#s3-configuration) below.

~> **Note:** At least one destination configuration (`cloud_watch_logs_configuration`, `firehose_configuration`, or `s3_configuration`) must be specified for each log configuration.

#### CloudWatch Logs Configuration

The `cloud_watch_logs_configuration` block supports the following:

* `log_group_arn` - (Optional) The ARN of the CloudWatch Logs log group to which the logs should be delivered.

#### Firehose Configuration

The `firehose_configuration` block supports the following:

* `stream_arn` - (Optional) The ARN of the Kinesis Data Firehose delivery stream to which the logs should be delivered.

#### S3 Configuration

The `s3_configuration` block supports the following:

* `bucket_arn` - (Optional) The ARN of the S3 bucket to which the logs should be delivered.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Cognito IDP (Identity Provider) Log Delivery Configuration using the `user_pool_id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.cognito_log_delivery_configuration import CognitoLogDeliveryConfiguration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        CognitoLogDeliveryConfiguration.generate_config_for_import(self, "example", "us-west-2_example123")
```

Using `terraform import`, import Cognito IDP (Identity Provider) Log Delivery Configuration using the `user_pool_id`. For example:

```console
% terraform import aws_cognito_log_delivery_configuration.example us-west-2_example123
```

<!-- cache-key: cdktf-0.20.8 input-06bcd87e77c8f84d74082358026627ce29979429c9b1eb667b8e9760f15f7795 -->