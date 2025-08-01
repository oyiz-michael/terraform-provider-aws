---
subcategory: "S3 (Simple Storage)"
layout: "aws"
page_title: "AWS: aws_s3_object"
description: |-
    Provides metadata and optionally content of an S3 object
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_s3_object

The S3 object data source allows access to the metadata and
_optionally_ (see below) content of an object stored inside S3 bucket.

~> **Note:** The content of an object (`body` field) is available only for objects which have a human-readable `Content-Type`:

* `text/*`
* `application/json`
* `application/ld+json`
* `application/x-httpd-php`
* `application/xhtml+xml`
* `application/x-csh`
* `application/x-sh`
* `application/xml`
* `application/atom+xml`
* `application/x-sql`
* `application/yaml`

This is to prevent printing unsafe characters and potentially downloading large amount of data which would be thrown away in favor of metadata.

## Example Usage

The following example retrieves a text object (which must have a `Content-Type`
value starting with `text/`) and uses it as the `user_data` for an EC2 instance:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_s3_object import DataAwsS3Object
from imports.aws.instance import Instance
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        bootstrap_script = DataAwsS3Object(self, "bootstrap_script",
            bucket="ourcorp-deploy-config",
            key="ec2-bootstrap-script.sh"
        )
        Instance(self, "example",
            ami="ami-2757f631",
            instance_type="t2.micro",
            user_data=Token.as_string(bootstrap_script.body)
        )
```

The following, more-complex example retrieves only the metadata for a zip
file stored in S3, which is then used to pass the most recent `version_id`
to AWS Lambda for use as a function implementation. More information about
Lambda functions is available in the documentation for
[`aws_lambda_function`](/docs/providers/aws/r/lambda_function.html).

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_s3_object import DataAwsS3Object
from imports.aws.lambda_function import LambdaFunction
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        lambda_ = DataAwsS3Object(self, "lambda",
            bucket="ourcorp-lambda-functions",
            key="hello-world.zip"
        )
        LambdaFunction(self, "test_lambda",
            function_name="lambda_function_name",
            handler="exports.test",
            role=iam_for_lambda.arn,
            s3_bucket=Token.as_string(lambda_.bucket),
            s3_key=Token.as_string(lambda_.key),
            s3_object_version=Token.as_string(lambda_.version_id)
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `bucket` - (Required) Name of the bucket to read the object from. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
* `checksum_mode` - (Optional) To retrieve the object's checksum, this argument must be `ENABLED`. If you enable `checksum_mode` and the object is encrypted with KMS, you must have permission to use the `kms:Decrypt` action. Valid values: `ENABLED`
* `key` - (Required) Full path to the object inside the bucket
* `version_id` - (Optional) Specific version ID of the object returned (defaults to latest version)

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the object.
* `body` - Object data (see **limitations above** to understand cases in which this field is actually available)
* `bucket_key_enabled` - (Optional) Whether or not to use [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) for SSE-KMS.
* `cache_control` - Caching behavior along the request/reply chain.
* `checksum_crc32` - The base64-encoded, 32-bit CRC32 checksum of the object.
* `checksum_crc32c` - The base64-encoded, 32-bit CRC32C checksum of the object.
* `checksum_crc64nvme` - The base64-encoded, 64-bit CRC64NVME checksum of the object.
* `checksum_sha1` - The base64-encoded, 160-bit SHA-1 digest of the object.
* `checksum_sha256` - The base64-encoded, 256-bit SHA-256 digest of the object.
* `content_disposition` - Presentational information for the object.
* `content_encoding` - What content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
* `content_language` - Language the content is in.
* `content_length` - Size of the body in bytes.
* `content_type` - Standard MIME type describing the format of the object data.
* `etag` - [ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)
* `expiration` - If the object expiration is configured (see [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.
* `expires` - Date and time at which the object is no longer cacheable.
* `last_modified` - Last modified date of the object in RFC1123 format (e.g., `Mon, 02 Jan 2006 15:04:05 MST`)
* `metadata` - Map of metadata stored with the object in S3. [Keys](https://developer.hashicorp.com/terraform/language/expressions/types#maps-objects) are always returned in lowercase.
* `object_lock_legal_hold_status` - Indicates whether this object has an active [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds). This field is only returned if you have permission to view an object's legal hold status.
* `object_lock_mode` - Object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) currently in place for this object.
* `object_lock_retain_until_date` - The date and time when this object's object lock will expire.
* `server_side_encryption` - If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.
* `sse_kms_key_id` - If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.
* `storage_class` - [Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for `Standard` storage class objects.
* `version_id` - Latest version ID of the object returned.
* `website_redirect_location` - If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.
* `tags`  - Map of tags assigned to the object.

-> **Note:** Terraform ignores all leading `/`s in the object's `key` and treats multiple `/`s in the rest of the object's `key` as a single `/`, so values of `/index.html` and `index.html` correspond to the same S3 object as do `first//second///third//` and `first/second/third/`.

<!-- cache-key: cdktf-0.20.8 input-44ca16d8e18b450ebb7c59f6e21159a0a2ac617b2ff8fcbbab0ca9f62ee14b40 -->