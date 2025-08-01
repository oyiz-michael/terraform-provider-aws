---
subcategory: "Storage Gateway"
layout: "aws"
page_title: "AWS: aws_storagegateway_nfs_file_share"
description: |-
  Manages an AWS Storage Gateway NFS File Share
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_storagegateway_nfs_file_share

Manages an AWS Storage Gateway NFS File Share.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.storagegateway_nfs_file_share import StoragegatewayNfsFileShare
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        StoragegatewayNfsFileShare(self, "example",
            client_list=["0.0.0.0/0"],
            gateway_arn=Token.as_string(aws_storagegateway_gateway_example.arn),
            location_arn=Token.as_string(aws_s3_bucket_example.arn),
            role_arn=Token.as_string(aws_iam_role_example.arn)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `client_list` - (Required) The list of clients that are allowed to access the file gateway. The list must contain either valid IP addresses or valid CIDR blocks. Set to `["0.0.0.0/0"]` to not limit access. Minimum 1 item. Maximum 100 items.
* `gateway_arn` - (Required) Amazon Resource Name (ARN) of the file gateway.
* `location_arn` - (Required) The ARN of the backed storage used for storing file data.
* `vpc_endpoint_dns_name` - (Optional) The DNS name of the VPC endpoint for S3 PrivateLink.
* `bucket_region` - (Optional) The region of the S3 bucket used by the file share. Required when specifying `vpc_endpoint_dns_name`.
* `role_arn` - (Required) The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
* `audit_destination_arn` - (Optional) The Amazon Resource Name (ARN) of the storage used for audit logs.
* `default_storage_class` - (Optional) The default [storage class](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-DefaultStorageClass) for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`.
* `guess_mime_type_enabled` - (Optional) Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
* `kms_encrypted` - (Optional) Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
* `kms_key_arn` - (Optional) Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
* `nfs_file_share_defaults` - (Optional) Nested argument with file share default values. More information below. see [NFS File Share Defaults](#nfs_file_share_defaults) for more details.
* `cache_attributes` - (Optional) Refresh cache information. see [Cache Attributes](#cache_attributes) for more details.
* `object_acl` - (Optional) Access Control List permission for S3 objects. Defaults to `private`.
* `read_only` - (Optional) Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
* `requester_pays` - (Optional) Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
* `squash` - (Optional) Maps a user to anonymous user. Defaults to `RootSquash`. Valid values: `RootSquash` (only root is mapped to anonymous user), `NoSquash` (no one is mapped to anonymous user), `AllSquash` (everyone is mapped to anonymous user)
* `file_share_name` - (Optional) The name of the file share. Must be set if an S3 prefix name is set in `location_arn`.
* `notification_policy` - (Optional) The notification policy of the file share. For more information see the [AWS Documentation](https://docs.aws.amazon.com/storagegateway/latest/APIReference/API_CreateNFSFileShare.html#StorageGateway-CreateNFSFileShare-request-NotificationPolicy). Default value is `{}`.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### nfs_file_share_defaults

Files and folders stored as Amazon S3 objects in S3 buckets don't, by default, have Unix file permissions assigned to them. Upon discovery in an S3 bucket by Storage Gateway, the S3 objects that represent files and folders are assigned these default Unix permissions.

* `directory_mode` - (Optional) The Unix directory mode in the string form "nnnn". Defaults to `"0777"`.
* `file_mode` - (Optional) The Unix file mode in the string form "nnnn". Defaults to `"0666"`.
* `group_id` - (Optional) The default group ID for the file share (unless the files have another group ID specified). Defaults to `65534` (`nfsnobody`). Valid values: `0` through `4294967294`.
* `owner_id` - (Optional) The default owner ID for the file share (unless the files have another owner ID specified). Defaults to `65534` (`nfsnobody`). Valid values: `0` through `4294967294`.

### cache_attributes

* `cache_stale_timeout_in_seconds` - (Optional) Refreshes a file share's cache by using Time To Live (TTL).
 TTL is the length of time since the last refresh after which access to the directory would cause the file gateway
  to first refresh that directory's contents from the Amazon S3 bucket. Valid Values: 300 to 2,592,000 seconds (5 minutes to 30 days)

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Amazon Resource Name (ARN) of the NFS File Share.
* `arn` - Amazon Resource Name (ARN) of the NFS File Share.
* `fileshare_id` - ID of the NFS File Share.
* `path` - File share path used by the NFS client to identify the mount point.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `10m`)
* `update` - (Default `10m`)
* `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_storagegateway_nfs_file_share` using the NFS File Share Amazon Resource Name (ARN). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.storagegateway_nfs_file_share import StoragegatewayNfsFileShare
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        StoragegatewayNfsFileShare.generate_config_for_import(self, "example", "arn:aws:storagegateway:us-east-1:123456789012:share/share-12345678")
```

Using `terraform import`, import `aws_storagegateway_nfs_file_share` using the NFS File Share Amazon Resource Name (ARN). For example:

```console
% terraform import aws_storagegateway_nfs_file_share.example arn:aws:storagegateway:us-east-1:123456789012:share/share-12345678
```

<!-- cache-key: cdktf-0.20.8 input-b22acab40a7d06f5cb9e80ef5e2c3cf1307ba6734cea6cb3ac0bd602eb959b48 -->