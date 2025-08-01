---
subcategory: "API Gateway"
layout: "aws"
page_title: "AWS: aws_api_gateway_resource"
description: |-
  Provides an API Gateway Resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_api_gateway_resource

Provides an API Gateway Resource.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApiGatewayResource } from "./.gen/providers/aws/api-gateway-resource";
import { ApiGatewayRestApi } from "./.gen/providers/aws/api-gateway-rest-api";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const myDemoApi = new ApiGatewayRestApi(this, "MyDemoAPI", {
      description: "This is my API for demonstration purposes",
      name: "MyDemoAPI",
    });
    new ApiGatewayResource(this, "MyDemoResource", {
      parentId: myDemoApi.rootResourceId,
      pathPart: "mydemoresource",
      restApiId: myDemoApi.id,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `restApiId` - (Required) ID of the associated REST API
* `parentId` - (Required) ID of the parent API resource
* `pathPart` - (Required) Last path segment of this API resource.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Resource's identifier.
* `path` - Complete path for this API resource, including all parent paths.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_api_gateway_resource` using `REST-API-ID/RESOURCE-ID`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApiGatewayResource } from "./.gen/providers/aws/api-gateway-resource";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ApiGatewayResource.generateConfigForImport(
      this,
      "example",
      "12345abcde/67890fghij"
    );
  }
}

```

Using `terraform import`, import `aws_api_gateway_resource` using `REST-API-ID/RESOURCE-ID`. For example:

```console
% terraform import aws_api_gateway_resource.example 12345abcde/67890fghij
```

<!-- cache-key: cdktf-0.20.8 input-992dc1ea3f06bea187f83bce1b945b7a93b0aeb75114916de486c508878e3cac -->