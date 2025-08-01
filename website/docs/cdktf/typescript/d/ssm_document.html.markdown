---
subcategory: "SSM (Systems Manager)"
layout: "aws"
page_title: "AWS: aws_ssm_document"
description: |-
  Provides a SSM Document data source
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ssm_document

Gets the contents of the specified Systems Manager document.

## Example Usage

To get the contents of the document owned by AWS.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformOutput, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsSsmDocument } from "./.gen/providers/aws/data-aws-ssm-document";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const foo = new DataAwsSsmDocument(this, "foo", {
      documentFormat: "YAML",
      name: "AWS-GatherSoftwareInventory",
    });
    new TerraformOutput(this, "content", {
      value: foo.content,
    });
  }
}

```

To get the contents of the custom document.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsSsmDocument } from "./.gen/providers/aws/data-aws-ssm-document";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsSsmDocument(this, "test", {
      documentFormat: "JSON",
      name: Token.asString(awsSsmDocumentTest.name),
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The name of the document.
* `documentFormat` - The format of the document. Valid values: `JSON`, `TEXT`, `YAML`.
* `documentVersion` - The document version.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the document. If the document is an AWS managed document, this value will be set to the name of the document instead.
* `content` - The content for the SSM document in JSON or YAML format.
* `documentType` - The type of the document.

<!-- cache-key: cdktf-0.20.8 input-eca5e31430e76e2d70d91656b5140c3fc87162021d3263a6164e23e8ec3774ef -->