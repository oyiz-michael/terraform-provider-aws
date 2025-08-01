---
subcategory: "EC2 Image Builder"
layout: "aws"
page_title: "AWS: aws_imagebuilder_container_recipes"
description: |-
    Get information on Image Builder Container Recipes.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_imagebuilder_container_recipes

Use this data source to get the ARNs and names of Image Builder Container Recipes matching the specified criteria.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsImagebuilderContainerRecipes } from "./.gen/providers/aws/data-aws-imagebuilder-container-recipes";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsImagebuilderContainerRecipes(this, "example", {
      filter: [
        {
          name: "platform",
          values: ["Linux"],
        },
      ],
      owner: "Self",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `owner` - (Optional) Owner of the container recipes. Valid values are `Self`, `Shared`, `Amazon` and `ThirdParty`. Defaults to `Self`.
* `filter` - (Optional) Configuration block(s) for filtering. Detailed below.

### filter Configuration Block

The `filter` configuration block supports the following arguments:

* `name` - (Required) Name of the filter field. Valid values can be found in the [Image Builder ListContainerRecipes API Reference](https://docs.aws.amazon.com/imagebuilder/latest/APIReference/API_ListContainerRecipes.html).
* `values` - (Required) Set of values that are accepted for the given filter field. Results will be selected if any given value matches.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arns` - Set of ARNs of the matched Image Builder Container Recipes.
* `names` - Set of names of the matched Image Builder Container Recipes.

<!-- cache-key: cdktf-0.20.8 input-8436e9ff70a1c87108c884abc3c2e207f53cac78f120e54d10d43d9ac01c1743 -->