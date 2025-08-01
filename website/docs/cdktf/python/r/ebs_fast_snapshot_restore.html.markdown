---
subcategory: "EBS (EC2)"
layout: "aws"
page_title: "AWS: aws_ebs_fast_snapshot_restore"
description: |-
  Terraform resource for managing an EBS (Elastic Block Storage) Fast Snapshot Restore.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ebs_fast_snapshot_restore

Terraform resource for managing an EBS (Elastic Block Storage) Fast Snapshot Restore.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ebs_fast_snapshot_restore import EbsFastSnapshotRestore
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        EbsFastSnapshotRestore(self, "example",
            availability_zone="us-west-2a",
            snapshot_id=Token.as_string(aws_ebs_snapshot_example.id)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `availability_zone` - (Required) Availability zone in which to enable fast snapshot restores.
* `snapshot_id` - (Required) ID of the snapshot.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - A comma-delimited string concatenating `availability_zone` and `snapshot_id`.
* `state` - State of fast snapshot restores. Valid values are `enabling`, `optimizing`, `enabled`, `disabling`, `disabled`.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `10m`)
* `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import EC2 (Elastic Compute Cloud) EBS Fast Snapshot Restore using the `example_id_arg`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.ebs_fast_snapshot_restore import EbsFastSnapshotRestore
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        EbsFastSnapshotRestore.generate_config_for_import(self, "example", "us-west-2a,snap-abcdef123456")
```

Using `terraform import`, import EC2 (Elastic Compute Cloud) EBS Fast Snapshot Restore using the `id`. For example:

```console
% terraform import aws_ebs_fast_snapshot_restore.example us-west-2a,snap-abcdef123456
```

<!-- cache-key: cdktf-0.20.8 input-be289e7b8efcb204c1afe4f782d388173ecde431c84320cc65d336b067e40124 -->