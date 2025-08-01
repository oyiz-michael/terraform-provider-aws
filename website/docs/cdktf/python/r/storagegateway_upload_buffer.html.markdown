---
subcategory: "Storage Gateway"
layout: "aws"
page_title: "AWS: aws_storagegateway_upload_buffer"
description: |-
  Manages an AWS Storage Gateway upload buffer
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_storagegateway_upload_buffer

Manages an AWS Storage Gateway upload buffer.

~> **NOTE:** The Storage Gateway API provides no method to remove an upload buffer disk. Destroying this Terraform resource does not perform any Storage Gateway actions.

## Example Usage

### Cached and VTL Gateway Type

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_storagegateway_local_disk import DataAwsStoragegatewayLocalDisk
from imports.aws.storagegateway_upload_buffer import StoragegatewayUploadBuffer
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        test = DataAwsStoragegatewayLocalDisk(self, "test",
            disk_node=Token.as_string(aws_volume_attachment_test.device_name),
            gateway_arn=Token.as_string(aws_storagegateway_gateway_test.arn)
        )
        aws_storagegateway_upload_buffer_test = StoragegatewayUploadBuffer(self, "test_1",
            disk_path=Token.as_string(test.disk_path),
            gateway_arn=Token.as_string(aws_storagegateway_gateway_test.arn)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_storagegateway_upload_buffer_test.override_logical_id("test")
```

### Stored Gateway Type

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_storagegateway_local_disk import DataAwsStoragegatewayLocalDisk
from imports.aws.storagegateway_upload_buffer import StoragegatewayUploadBuffer
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        StoragegatewayUploadBuffer(self, "example",
            disk_id=Token.as_string(data_aws_storagegateway_local_disk_example.id),
            gateway_arn=Token.as_string(aws_storagegateway_gateway_example.arn)
        )
        DataAwsStoragegatewayLocalDisk(self, "test",
            disk_node=Token.as_string(aws_volume_attachment_test.device_name),
            gateway_arn=Token.as_string(aws_storagegateway_gateway_test.arn)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `disk_id` - (Optional) Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
* `disk_path` - (Optional) Local disk path. For example, `/dev/nvme1n1`.
* `gateway_arn` - (Required) The Amazon Resource Name (ARN) of the gateway.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Combined gateway Amazon Resource Name (ARN) and local disk identifier.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_storagegateway_upload_buffer` using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.storagegateway_upload_buffer import StoragegatewayUploadBuffer
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        StoragegatewayUploadBuffer.generate_config_for_import(self, "example", "arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0")
```

Using `terraform import`, import `aws_storagegateway_upload_buffer` using the gateway Amazon Resource Name (ARN) and local disk identifier separated with a colon (`:`). For example:

```console
% terraform import aws_storagegateway_upload_buffer.example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678:pci-0000:03:00.0-scsi-0:0:0:0
```

<!-- cache-key: cdktf-0.20.8 input-b0af0afe875976ec7840026599cb6afae2ba25fa1e765fc7e2c96e4eaea75db7 -->