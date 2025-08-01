---
subcategory: "EC2 Image Builder"
layout: "aws"
page_title: "AWS: aws_imagebuilder_image_pipeline"
description: |-
    Manages an Image Builder Image Pipeline
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_imagebuilder_image_pipeline

Manages an Image Builder Image Pipeline.

~> **NOTE:** Starting with version `5.74.0`, lifecycle meta-argument [`replace_triggered_by`](https://developer.hashicorp.com/terraform/language/meta-arguments/lifecycle#replace_triggered_by) must be used in order to prevent a dependency error on destroy.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ImagebuilderImagePipeline } from "./.gen/providers/aws/imagebuilder-image-pipeline";
import { ImagebuilderImageRecipe } from "./.gen/providers/aws/imagebuilder-image-recipe";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new ImagebuilderImageRecipe(this, "example", {
      blockDeviceMapping: [
        {
          deviceName: "/dev/xvdb",
          ebs: {
            deleteOnTermination: Token.asString(true),
            volumeSize: 100,
            volumeType: "gp2",
          },
        },
      ],
      component: [
        {
          componentArn: Token.asString(awsImagebuilderComponentExample.arn),
          parameter: [
            {
              name: "Parameter1",
              value: "Value1",
            },
            {
              name: "Parameter2",
              value: "Value2",
            },
          ],
        },
      ],
      name: "example",
      parentImage:
        "arn:${" +
        current.partition +
        "}:imagebuilder:${" +
        dataAwsRegionCurrent.region +
        "}:aws:image/amazon-linux-2-x86/x.x.x",
      version: "1.0.0",
    });
    const awsImagebuilderImagePipelineExample = new ImagebuilderImagePipeline(
      this,
      "example_1",
      {
        imageRecipeArn: example.arn,
        infrastructureConfigurationArn: Token.asString(
          awsImagebuilderInfrastructureConfigurationExample.arn
        ),
        lifecycle: {
          replaceTriggeredBy: [example],
        },
        name: "example",
        schedule: {
          scheduleExpression: "cron(0 0 * * ? *)",
        },
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsImagebuilderImagePipelineExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

The following arguments are required:

* `infrastructureConfigurationArn` - (Required) Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
* `name` - (Required) Name of the image pipeline.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `containerRecipeArn` - (Optional) Amazon Resource Name (ARN) of the container recipe.
* `description` - (Optional) Description of the image pipeline.
* `distributionConfigurationArn` - (Optional) Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
* `enhancedImageMetadataEnabled` - (Optional) Whether additional information about the image being created is collected. Defaults to `true`.
* `executionRole` - (Optional) Amazon Resource Name (ARN) of the service-linked role to be used by Image Builder to [execute workflows](https://docs.aws.amazon.com/imagebuilder/latest/userguide/manage-image-workflows.html).
* `imageRecipeArn` - (Optional) Amazon Resource Name (ARN) of the image recipe.
* `imageScanningConfiguration` - (Optional) Configuration block with image scanning configuration. Detailed below.
* `imageTestsConfiguration` - (Optional) Configuration block with image tests configuration. Detailed below.
* `schedule` - (Optional) Configuration block with schedule settings. Detailed below.
* `status` - (Optional) Status of the image pipeline. Valid values are `DISABLED` and `ENABLED`. Defaults to `ENABLED`.
* `workflow` - (Optional) Configuration block with the workflow configuration. Detailed below.
* `tags` - (Optional) Key-value map of resource tags for the image pipeline. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

### image_scanning_configuration

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `imageScanningEnabled` - (Optional) Whether image scans are enabled. Defaults to `false`.
* `ecrConfiguration` - (Optional) Configuration block with ECR configuration for image scanning. Detailed below.

### ecr_configuration

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `container tags` - (Optional) list of tags to apply to scanned images
* `repositoryName` - (Optional) The name of the repository to scan

### image_tests_configuration

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `imageTestsEnabled` - (Optional) Whether image tests are enabled. Defaults to `true`.
* `timeoutMinutes` - (Optional) Number of minutes before image tests time out. Valid values are between `60` and `1440`. Defaults to `720`.

### schedule

The following arguments are required:

* `scheduleExpression` - (Required) Cron expression of how often the pipeline start condition is evaluated. For example, `cron(0 0 * * ? *)` is evaluated every day at midnight UTC. Configurations using the five field syntax that was previously accepted by the API, such as `cron(0 0 * * *)`, must be updated to the six field syntax. For more information, see the [Image Builder User Guide](https://docs.aws.amazon.com/imagebuilder/latest/userguide/cron-expressions.html).

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `pipelineExecutionStartCondition` - (Optional) Condition when the pipeline should trigger a new image build. Valid values are `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE` and `EXPRESSION_MATCH_ONLY`. Defaults to `EXPRESSION_MATCH_AND_DEPENDENCY_UPDATES_AVAILABLE`.

* `timezone` - (Optional) The timezone that applies to the scheduling expression. For example, "Etc/UTC", "America/Los_Angeles" in the [IANA timezone format](https://www.joda.org/joda-time/timezones.html). If not specified this defaults to UTC.

### workflow

The following arguments are required:

* `workflowArn` - (Required) Amazon Resource Name (ARN) of the Image Builder Workflow.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `onFailure` - (Optional) The action to take if the workflow fails. Must be one of `CONTINUE` or `ABORT`.
* `parallelGroup` - (Optional) The parallel group in which to run a test Workflow.
* `parameter` - (Optional) Configuration block for the workflow parameters. Detailed below.

### parameter

The following arguments are required:

* `name` - (Required) The name of the Workflow parameter.
* `value` - (Required) The value of the Workflow parameter.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the image pipeline.
* `dateCreated` - Date the image pipeline was created.
* `dateLastRun` - Date the image pipeline was last run.
* `dateNextRun` - Date the image pipeline will run next.
* `dateUpdated` - Date the image pipeline was updated.
* `platform` - Platform of the image pipeline.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_imagebuilder_image_pipeline` resources using the Amazon Resource Name (ARN). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ImagebuilderImagePipeline } from "./.gen/providers/aws/imagebuilder-image-pipeline";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ImagebuilderImagePipeline.generateConfigForImport(
      this,
      "example",
      "arn:aws:imagebuilder:us-east-1:123456789012:image-pipeline/example"
    );
  }
}

```

Using `terraform import`, import `aws_imagebuilder_image_pipeline` resources using the Amazon Resource Name (ARN). For example:

```console
% terraform import aws_imagebuilder_image_pipeline.example arn:aws:imagebuilder:us-east-1:123456789012:image-pipeline/example
```

<!-- cache-key: cdktf-0.20.8 input-8e0f255e7e531158d50c7ccc2b2e109906c6e4936b9c91ef1dc11dfc1a99c5c9 -->