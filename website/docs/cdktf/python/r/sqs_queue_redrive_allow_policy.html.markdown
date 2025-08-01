---
subcategory: "SQS (Simple Queue)"
layout: "aws"
page_title: "AWS: aws_sqs_queue_redrive_allow_policy"
description: |-
  Provides a SQS Queue Redrive Allow Policy resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_sqs_queue_redrive_allow_policy

Provides a SQS Queue Redrive Allow Policy resource.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sqs_queue import SqsQueue
from imports.aws.sqs_queue_redrive_allow_policy import SqsQueueRedriveAllowPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = SqsQueue(self, "example",
            name="examplequeue"
        )
        src = SqsQueue(self, "src",
            name="srcqueue",
            redrive_policy=Token.as_string(
                Fn.jsonencode({
                    "dead_letter_target_arn": example.arn,
                    "max_receive_count": 4
                }))
        )
        aws_sqs_queue_redrive_allow_policy_example = SqsQueueRedriveAllowPolicy(self, "example_2",
            queue_url=example.id,
            redrive_allow_policy=Token.as_string(
                Fn.jsonencode({
                    "redrive_permission": "byQueue",
                    "source_queue_arns": [src.arn]
                }))
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_sqs_queue_redrive_allow_policy_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `queue_url` - (Required) The URL of the SQS Queue to which to attach the policy
* `redrive_allow_policy` - (Required) The JSON redrive allow policy for the SQS queue. Learn more in the [Amazon SQS dead-letter queues documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html).

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SQS Queue Redrive Allow Policies using the queue URL. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.sqs_queue_redrive_allow_policy import SqsQueueRedriveAllowPolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        SqsQueueRedriveAllowPolicy.generate_config_for_import(self, "test", "https://queue.amazonaws.com/123456789012/myqueue")
```

Using `terraform import`, import SQS Queue Redrive Allow Policies using the queue URL. For example:

```console
% terraform import aws_sqs_queue_redrive_allow_policy.test https://queue.amazonaws.com/123456789012/myqueue
```

<!-- cache-key: cdktf-0.20.8 input-e1c64e6e2458160f388b9a7ec488abd3bd04a10ae8b82f69d0d0a207b50c127d -->