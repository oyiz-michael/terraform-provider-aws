---
subcategory: "Kendra"
layout: "aws"
page_title: "AWS: aws_kendra_thesaurus"
description: |-
  Provides details about a specific Amazon Kendra Thesaurus.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_kendra_thesaurus

Provides details about a specific Amazon Kendra Thesaurus.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsKendraThesaurus } from "./.gen/providers/aws/data-aws-kendra-thesaurus";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsKendraThesaurus(this, "example", {
      indexId: "12345678-1234-1234-1234-123456789123",
      thesaurusId: "87654321-1234-4321-4321-321987654321",
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `indexId` - (Required) Identifier of the index that contains the Thesaurus.
* `thesaurusId` - (Required) Identifier of the Thesaurus.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Thesaurus.
* `createdAt` - Unix datetime that the Thesaurus was created.
* `description` - Description of the Thesaurus.
* `errorMessage` - When the `status` field value is `FAILED`, this contains a message that explains why.
* `fileSizeBytes` - Size of the Thesaurus file in bytes.
* `id` - Unique identifiers of the Thesaurus and index separated by a slash (`/`).
* `name` - Name of the Thesaurus.
* `roleArn` - ARN of a role with permission to access the S3 bucket that contains the Thesaurus. For more information, see [IAM Roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
* `sourceS3Path` - S3 location of the Thesaurus input data. Detailed below.
* `status` - Status of the Thesaurus. It is ready to use when the status is `ACTIVE`.
* `synonymRuleCount` - Number of synonym rules in the Thesaurus file.
* `termCount` - Number of unique terms in the Thesaurus file. For example, the synonyms `a,b,c` and `a=>d`, the term count would be 4.
* `updatedAt` - Date and time that the Thesaurus was last updated.
* `tags` - Metadata that helps organize the Thesaurus you create.

The `sourceS3Path` configuration block supports the following attributes:

* `bucket` - Name of the S3 bucket that contains the file.
* `key` - Name of the file.

<!-- cache-key: cdktf-0.20.8 input-b5c6241eae45023ae88c51189ce374790152219b94f3accc9790851b524815c3 -->