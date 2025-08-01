---
subcategory: "S3 (Simple Storage)"
layout: "aws"
page_title: "AWS: aws_s3_bucket_inventory"
description: |-
  Provides a S3 bucket inventory configuration resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_s3_bucket_inventory

Provides a S3 bucket [inventory configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html) resource.

-> This resource cannot be used with S3 directory buckets.

## Example Usage

### Add inventory configuration

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.s3_bucket import S3Bucket
from imports.aws.s3_bucket_inventory import S3BucketInventory
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        inventory = S3Bucket(self, "inventory",
            bucket="my-tf-inventory-bucket"
        )
        test = S3Bucket(self, "test",
            bucket="my-tf-test-bucket"
        )
        aws_s3_bucket_inventory_test = S3BucketInventory(self, "test_2",
            bucket=test.id,
            destination=S3BucketInventoryDestination(
                bucket=S3BucketInventoryDestinationBucket(
                    bucket_arn=inventory.arn,
                    format="ORC"
                )
            ),
            included_object_versions="All",
            name="EntireBucketDaily",
            schedule=S3BucketInventorySchedule(
                frequency="Daily"
            )
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_s3_bucket_inventory_test.override_logical_id("test")
```

### Add inventory configuration with S3 object prefix

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.s3_bucket import S3Bucket
from imports.aws.s3_bucket_inventory import S3BucketInventory
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        inventory = S3Bucket(self, "inventory",
            bucket="my-tf-inventory-bucket"
        )
        test = S3Bucket(self, "test",
            bucket="my-tf-test-bucket"
        )
        S3BucketInventory(self, "test-prefix",
            bucket=test.id,
            destination=S3BucketInventoryDestination(
                bucket=S3BucketInventoryDestinationBucket(
                    bucket_arn=inventory.arn,
                    format="ORC",
                    prefix="inventory"
                )
            ),
            filter=S3BucketInventoryFilter(
                prefix="documents/"
            ),
            included_object_versions="All",
            name="DocumentsWeekly",
            schedule=S3BucketInventorySchedule(
                frequency="Daily"
            )
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `bucket` - (Required) Name of the source bucket that inventory lists the objects for.
* `name` - (Required) Unique identifier of the inventory configuration for the bucket.
* `included_object_versions` - (Required) Object versions to include in the inventory list. Valid values: `All`, `Current`.
* `schedule` - (Required) Specifies the schedule for generating inventory results (documented below).
* `destination` - (Required) Contains information about where to publish the inventory results (documented below).
* `enabled` - (Optional, Default: `true`) Specifies whether the inventory is enabled or disabled.
* `filter` - (Optional) Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
* `optional_fields` - (Optional) List of optional fields that are included in the inventory results. Please refer to the S3 [documentation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_InventoryConfiguration.html#AmazonS3-Type-InventoryConfiguration-OptionalFields) for more details.

The `filter` configuration supports the following:

* `prefix` - (Optional) Prefix that an object must have to be included in the inventory results.

The `schedule` configuration supports the following:

* `frequency` - (Required) Specifies how frequently inventory results are produced. Valid values: `Daily`, `Weekly`.

The `destination` configuration supports the following:

* `bucket` - (Required) S3 bucket configuration where inventory results are published (documented below).

The `bucket` configuration supports the following:

* `bucket_arn` - (Required) Amazon S3 bucket ARN of the destination.
* `format` - (Required) Specifies the output format of the inventory results. Can be `CSV`, [`ORC`](https://orc.apache.org/) or [`Parquet`](https://parquet.apache.org/).
* `account_id` - (Optional) ID of the account that owns the destination bucket. Recommended to be set to prevent problems if the destination bucket ownership changes.
* `prefix` - (Optional) Prefix that is prepended to all inventory results.
* `encryption` - (Optional) Contains the type of server-side encryption to use to encrypt the inventory (documented below).

The `encryption` configuration supports the following:

* `sse_kms` - (Optional) Specifies to use server-side encryption with AWS KMS-managed keys to encrypt the inventory file (documented below).
* `sse_s3` - (Optional) Specifies to use server-side encryption with Amazon S3-managed keys (SSE-S3) to encrypt the inventory file.

The `sse_kms` configuration supports the following:

* `key_id` - (Required) ARN of the KMS customer master key (CMK) used to encrypt the inventory file.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import S3 bucket inventory configurations using `bucket:inventory`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.s3_bucket_inventory import S3BucketInventory
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        S3BucketInventory.generate_config_for_import(self, "myBucketEntireBucket", "my-bucket:EntireBucket")
```

Using `terraform import`, import S3 bucket inventory configurations using `bucket:inventory`. For example:

```console
% terraform import aws_s3_bucket_inventory.my-bucket-entire-bucket my-bucket:EntireBucket
```

<!-- cache-key: cdktf-0.20.8 input-499f31ef13749d25d44eeae505cdd0f9a631b462702e522a4fe5109fce2e31b0 -->