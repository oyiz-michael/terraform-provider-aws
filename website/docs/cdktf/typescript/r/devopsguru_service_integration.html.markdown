---
subcategory: "DevOps Guru"
layout: "aws"
page_title: "AWS: aws_devopsguru_service_integration"
description: |-
  Terraform resource for managing an AWS DevOps Guru Service Integration.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_devopsguru_service_integration

Terraform resource for managing an AWS DevOps Guru Service Integration.

~> To prevent unintentional deletion of account wide settings, destruction of this resource will only remove it from the Terraform state. To disable any configured settings, explicitly set the opt-in value to `DISABLED` and apply again before destroying.

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
import { DevopsguruServiceIntegration } from "./.gen/providers/aws/devopsguru-service-integration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DevopsguruServiceIntegration(this, "example", {
      kmsServerSideEncryption: [
        {
          optInStatus: "ENABLED",
          type: "AWS_OWNED_KMS_KEY",
        },
      ],
      logsAnomalyDetection: [
        {
          optInStatus: "ENABLED",
        },
      ],
      opsCenter: [
        {
          optInStatus: "ENABLED",
        },
      ],
    });
  }
}

```

### Customer Managed KMS Key

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DevopsguruServiceIntegration } from "./.gen/providers/aws/devopsguru-service-integration";
import { KmsKey } from "./.gen/providers/aws/kms-key";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DevopsguruServiceIntegration(this, "example", {
      kmsServerSideEncryption: [
        {
          kmsKeyId: test.arn,
          optInStatus: "ENABLED",
          type: "CUSTOMER_MANAGED_KEY",
        },
      ],
      logsAnomalyDetection: [
        {
          optInStatus: "DISABLED",
        },
      ],
      opsCenter: [
        {
          optInStatus: "DISABLED",
        },
      ],
    });
    const awsKmsKeyExample = new KmsKey(this, "example_1", {});
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsKmsKeyExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `kmsServerSideEncryption` - (Required) Information about whether DevOps Guru is configured to encrypt server-side data using KMS. See [`kmsServerSideEncryption`](#kms_server_side_encryption-argument-reference) below.
* `logsAnomalyDetection` - (Required) Information about whether DevOps Guru is configured to perform log anomaly detection on Amazon CloudWatch log groups. See [`logsAnomalyDetection`](#logs_anomaly_detection-argument-reference) below.
* `opsCenter` - (Required) Information about whether DevOps Guru is configured to create an OpsItem in AWS Systems Manager OpsCenter for each created insight. See [`opsCenter`](#ops_center-argument-reference) below.

### `kmsServerSideEncryption` Argument Reference

* `kmsKeyId` - (Optional) KMS key ID. This value can be a key ID, key ARN, alias name, or alias ARN.
* `optInStatus` - (Optional) Specifies whether KMS integration is enabled. Valid values are `DISABLED` and `ENABLED`.
* `type` - (Optional) Type of KMS key used. Valid values are `CUSTOMER_MANAGED_KEY` and `AWS_OWNED_KMS_KEY`.

### `logsAnomalyDetection` Argument Reference

* `optInStatus` - (Optional) Specifies if DevOps Guru is configured to perform log anomaly detection on CloudWatch log groups. Valid values are `DISABLED` and `ENABLED`.

### `opsCenter` Argument Reference

* `optInStatus` - (Optional) Specifies if DevOps Guru is enabled to create an AWS Systems Manager OpsItem for each created insight. Valid values are `DISABLED` and `ENABLED`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - AWS region.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DevOps Guru Service Integration using the region. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DevopsguruServiceIntegration } from "./.gen/providers/aws/devopsguru-service-integration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    DevopsguruServiceIntegration.generateConfigForImport(
      this,
      "example",
      "us-east-1"
    );
  }
}

```

Using `terraform import`, import DevOps Guru Service Integration using the region. For example:

```console
% terraform import aws_devopsguru_service_integration.example us-east-1
```

<!-- cache-key: cdktf-0.20.8 input-3909a0303a6e70a7e73d52d385a16ecae6b832ca6873903ce4f1fb582bee1afd -->