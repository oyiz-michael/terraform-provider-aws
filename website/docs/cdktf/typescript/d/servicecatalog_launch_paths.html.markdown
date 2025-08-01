---
subcategory: "Service Catalog"
layout: "aws"
page_title: "AWS: aws_servicecatalog_launch_paths"
description: |-
  Provides information on Service Catalog Launch Paths
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_servicecatalog_launch_paths

Lists the paths to the specified product. A path is how the user has access to a specified product, and is necessary when provisioning a product. A path also determines the constraints put on the product.

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
import { DataAwsServicecatalogLaunchPaths } from "./.gen/providers/aws/data-aws-servicecatalog-launch-paths";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsServicecatalogLaunchPaths(this, "example", {
      productId: "prod-yakog5pdriver",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `productId` - (Required) Product identifier.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `acceptLanguage` - (Optional) Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `summaries` - Block with information about the launch path. See details below.

### summaries

* `constraint_summaries` - Block for constraints on the portfolio-product relationship. See details below.
* `pathId` - Identifier of the product path.
* `name` - Name of the portfolio to which the path was assigned.
* `tags` - Tags associated with this product path.

### constraint_summaries

* `description` - Description of the constraint.
* `type` - Type of constraint. Valid values are `LAUNCH`, `NOTIFICATION`, `STACKSET`, and `TEMPLATE`.

<!-- cache-key: cdktf-0.20.8 input-44f221fe998314af9e0151501a2bcbf8a9a96b758431c76f2fdfc4d5e26ac8a1 -->