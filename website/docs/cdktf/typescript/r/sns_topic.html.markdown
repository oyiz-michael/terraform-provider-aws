---
subcategory: "SNS (Simple Notification)"
layout: "aws"
page_title: "AWS: aws_sns_topic"
description: |-
  Provides an SNS topic resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_sns_topic

Provides an SNS topic resource

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SnsTopic } from "./.gen/providers/aws/sns-topic";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SnsTopic(this, "user_updates", {
      name: "user-updates-topic",
    });
  }
}

```

## Example with Delivery Policy

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SnsTopic } from "./.gen/providers/aws/sns-topic";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SnsTopic(this, "user_updates", {
      deliveryPolicy:
        '{\n  "http": {\n    "defaultHealthyRetryPolicy": {\n      "minDelayTarget": 20,\n      "maxDelayTarget": 20,\n      "numRetries": 3,\n      "numMaxDelayRetries": 0,\n      "numNoDelayRetries": 0,\n      "numMinDelayRetries": 0,\n      "backoffFunction": "linear"\n    },\n    "disableSubscriptionOverrides": false,\n    "defaultThrottlePolicy": {\n      "maxReceivesPerSecond": 1\n    }\n  }\n}\n\n',
      name: "user-updates-topic",
    });
  }
}

```

## Example with Server-side encryption (SSE)

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SnsTopic } from "./.gen/providers/aws/sns-topic";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SnsTopic(this, "user_updates", {
      kmsMasterKeyId: "alias/aws/sns",
      name: "user-updates-topic",
    });
  }
}

```

## Example with First-In-First-Out (FIFO)

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SnsTopic } from "./.gen/providers/aws/sns-topic";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SnsTopic(this, "user_updates", {
      contentBasedDeduplication: true,
      fifoTopic: true,
      name: "user-updates-topic.fifo",
    });
  }
}

```

## Message Delivery Status Arguments

The `<endpoint>_success_feedback_role_arn` and `<endpoint>_failure_feedback_role_arn` arguments are used to give Amazon SNS write access to use CloudWatch Logs on your behalf. The `<endpoint>_success_feedback_sample_rate` argument is for specifying the sample rate percentage (0-100) of successfully delivered messages. After you configure the  `<endpoint>_failure_feedback_role_arn` argument, then all failed message deliveries generate CloudWatch Logs.

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Optional) The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, Terraform will assign a random, unique name. Conflicts with `namePrefix`
* `namePrefix` - (Optional) Creates a unique name beginning with the specified prefix. Conflicts with `name`
* `displayName` - (Optional) The display name for the topic
* `policy` - (Optional) The fully-formed AWS policy as JSON. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://learn.hashicorp.com/terraform/aws/iam-policy).
* `deliveryPolicy` - (Optional) The SNS delivery policy. More details in the [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html).
* `applicationSuccessFeedbackRoleArn` - (Optional) The IAM role permitted to receive success feedback for this topic
* `applicationSuccessFeedbackSampleRate` - (Optional) Percentage of success to sample
* `applicationFailureFeedbackRoleArn` - (Optional) IAM role for failure feedback
* `httpSuccessFeedbackRoleArn` - (Optional) The IAM role permitted to receive success feedback for this topic
* `httpSuccessFeedbackSampleRate` - (Optional) Percentage of success to sample
* `httpFailureFeedbackRoleArn` - (Optional) IAM role for failure feedback
* `kmsMasterKeyId` - (Optional) The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
* `signatureVersion` - (Optional) If `SignatureVersion` should be [1 (SHA1) or 2 (SHA256)](https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html). The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
* `tracingConfig` - (Optional) Tracing mode of an Amazon SNS topic. Valid values: `"PassThrough"`, `"Active"`.
* `fifoThroughputScope` - (Optional) Enables higher throughput for FIFO topics by adjusting the scope of deduplication. This attribute has two possible values, `Topic` and `MessageGroup`. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-high-throughput.html#enable-high-throughput-on-fifo-topic).
* `fifoTopic` - (Optional) Boolean indicating whether or not to create a FIFO (first-in-first-out) topic. FIFO topics can't deliver messages to customer managed endpoints, such as email addresses, mobile apps, SMS, or HTTP(S) endpoints. These endpoint types aren't guaranteed to preserve strict message ordering. Default is `false`.
* `archivePolicy` - (Optional) The message archive policy for FIFO topics. More details in the [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/message-archiving-and-replay-topic-owner.html).
* `contentBasedDeduplication` - (Optional) Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
* `lambdaSuccessFeedbackRoleArn` - (Optional) The IAM role permitted to receive success feedback for this topic
* `lambdaSuccessFeedbackSampleRate` - (Optional) Percentage of success to sample
* `lambdaFailureFeedbackRoleArn` - (Optional) IAM role for failure feedback
* `sqsSuccessFeedbackRoleArn` - (Optional) The IAM role permitted to receive success feedback for this topic
* `sqsSuccessFeedbackSampleRate` - (Optional) Percentage of success to sample
* `sqsFailureFeedbackRoleArn` - (Optional) IAM role for failure feedback
* `firehoseSuccessFeedbackRoleArn` - (Optional) The IAM role permitted to receive success feedback for this topic
* `firehoseSuccessFeedbackSampleRate` - (Optional) Percentage of success to sample
* `firehoseFailureFeedbackRoleArn` - (Optional) IAM role for failure feedback
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ARN of the SNS topic
* `arn` - The ARN of the SNS topic, as a more obvious property (clone of id)
* `beginningArchiveTime` - The oldest timestamp at which a FIFO topic subscriber can start a replay.
* `owner` - The AWS Account ID of the SNS topic owner
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import SNS Topics using the topic `arn`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SnsTopic } from "./.gen/providers/aws/sns-topic";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    SnsTopic.generateConfigForImport(
      this,
      "userUpdates",
      "arn:aws:sns:us-west-2:123456789012:my-topic"
    );
  }
}

```

Using `terraform import`, import SNS Topics using the topic `arn`. For example:

```console
% terraform import aws_sns_topic.user_updates arn:aws:sns:us-west-2:123456789012:my-topic
```

<!-- cache-key: cdktf-0.20.8 input-ed0f67557e435c2c9394ea7e885e3b460ec3c33646217efbfb1ef736bdd08a59 -->