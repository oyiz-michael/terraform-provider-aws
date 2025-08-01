---
subcategory: "Connect"
layout: "aws"
page_title: "AWS: aws_connect_instance"
description: |-
  Provides details about a specific Connect Instance.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_connect_instance

Provides details about a specific Amazon Connect Instance.

## Example Usage

By instance_alias

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsConnectInstance } from "./.gen/providers/aws/data-aws-connect-instance";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsConnectInstance(this, "foo", {
      instanceAlias: "foo",
    });
  }
}

```

By instance_id

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsConnectInstance } from "./.gen/providers/aws/data-aws-connect-instance";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsConnectInstance(this, "foo", {
      instanceId: "97afc98d-101a-ba98-ab97-ae114fc115ec",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `instanceId` - (Optional) Returns information on a specific connect instance by id
* `instanceAlias` - (Optional) Returns information on a specific connect instance by alias

~> **NOTE:** One of either `instanceId` or `instanceAlias` is required.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `createdTime` - When the instance was created.
* `arn` - ARN of the instance.
* `identityManagementType` - Specifies The identity management type attached to the instance.
* `inboundCallsEnabled` - Whether inbound calls are enabled.
* `outboundCallsEnabled` - Whether outbound calls are enabled.
* `earlyMediaEnabled` - Whether early media for outbound calls is enabled .
* `contactFlowLogsEnabled` - Whether contact flow logs are enabled.
* `contactLensEnabled` - Whether contact lens is enabled.
* `auto_resolve_best_voices` - Whether auto resolve best voices is enabled.
* `multiPartyConferenceEnabled` - Whether multi-party calls/conference is enabled.
* `use_custom_tts_voices` - Whether use custom tts voices is enabled.
* `status` - State of the instance.
* `serviceRole` - Service role of the instance.
* `tags` - A map of tags to assigned to the instance.

<!-- cache-key: cdktf-0.20.8 input-c32a17ecd6c0c47a21038fc69d6955947565ac2871bd485f076d94c78e61056a -->