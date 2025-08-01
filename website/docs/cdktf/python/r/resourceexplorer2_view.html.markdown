---
subcategory: "Resource Explorer"
layout: "aws"
page_title: "AWS: aws_resourceexplorer2_view"
description: |-
  Provides a resource to manage a Resource Explorer view.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_resourceexplorer2_view

Provides a resource to manage a Resource Explorer view.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.resourceexplorer2_index import Resourceexplorer2Index
from imports.aws.resourceexplorer2_view import Resourceexplorer2View
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = Resourceexplorer2Index(self, "example",
            type="LOCAL"
        )
        aws_resourceexplorer2_view_example = Resourceexplorer2View(self, "example_1",
            depends_on=[example],
            filters=[Resourceexplorer2ViewFilters(
                filter_string="resourcetype:ec2:instance"
            )
            ],
            included_property=[Resourceexplorer2ViewIncludedProperty(
                name="tags"
            )
            ],
            name="exampleview"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_resourceexplorer2_view_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `default_view` - (Optional) Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
* `filters` - (Optional) Specifies which resources are included in the results of queries made using this view. See [Filters](#filters) below for more details.
* `included_property` - (Optional) Optional fields to be included in search results from this view. See [Included Properties](#included-properties) below for more details.
* `name` - (Required) The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
* `scope`- (Optional) The root ARN of the account, an organizational unit (OU), or an organization ARN. If left empty, the default is account.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### Filters

The `filters` block supports the following:

* `filter_string` - (Required) The string that contains the search keywords, prefixes, and operators to control the results that can be returned by a search operation. For more details, see [Search query syntax](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html).

### Included Properties

The `included_property` block supports the following:

* `name` - (Required) The name of the property that is included in this view. Valid values: `tags`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the Resource Explorer view.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Resource Explorer views using the `arn`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.resourceexplorer2_view import Resourceexplorer2View
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Resourceexplorer2View.generate_config_for_import(self, "example", "arn:aws:resource-explorer-2:us-west-2:123456789012:view/exampleview/e0914f6c-6c27-4b47-b5d4-6b28381a2421")
```

Using `terraform import`, import Resource Explorer views using the `arn`. For example:

```console
% terraform import aws_resourceexplorer2_view.example arn:aws:resource-explorer-2:us-west-2:123456789012:view/exampleview/e0914f6c-6c27-4b47-b5d4-6b28381a2421
```

<!-- cache-key: cdktf-0.20.8 input-f37802152093455dbddab53c970210427c4c2dd1ddb98733b92823b676e04b58 -->