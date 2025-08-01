---
subcategory: "Location"
layout: "aws"
page_title: "AWS: aws_location_tracker"
description: |-
    Retrieve information about a Location Service Tracker.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_location_tracker

Retrieve information about a Location Service Tracker.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_location_tracker import DataAwsLocationTracker
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsLocationTracker(self, "example",
            tracker_name="example"
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `tracker_name` - (Required) Name of the tracker resource.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `create_time` - Timestamp for when the tracker resource was created in ISO 8601 format.
* `description` - Optional description for the tracker resource.
* `kms_key_id` - Key identifier for an AWS KMS customer managed key assigned to the Amazon Location resource.
* `position_filtering` - Position filtering method of the tracker resource.
* `tags` - Key-value map of resource tags for the tracker.
* `tracker_arn` - ARN for the tracker resource. Used when you need to specify a resource across all AWS.
* `update_time` - Timestamp for when the tracker resource was last updated in ISO 8601 format.

<!-- cache-key: cdktf-0.20.8 input-4c4d39ee343ad0ecb6eff6915853cc0fe3346b0a1ef7a4755de993c32a67e00c -->