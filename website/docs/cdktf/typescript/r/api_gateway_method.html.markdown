---
subcategory: "API Gateway"
layout: "aws"
page_title: "AWS: aws_api_gateway_method"
description: |-
  Provides a HTTP Method for an API Gateway Resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_api_gateway_method

Provides a HTTP Method for an API Gateway Resource.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApiGatewayMethod } from "./.gen/providers/aws/api-gateway-method";
import { ApiGatewayResource } from "./.gen/providers/aws/api-gateway-resource";
import { ApiGatewayRestApi } from "./.gen/providers/aws/api-gateway-rest-api";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const myDemoApi = new ApiGatewayRestApi(this, "MyDemoAPI", {
      description: "This is my API for demonstration purposes",
      name: "MyDemoAPI",
    });
    const myDemoResource = new ApiGatewayResource(this, "MyDemoResource", {
      parentId: myDemoApi.rootResourceId,
      pathPart: "mydemoresource",
      restApiId: myDemoApi.id,
    });
    new ApiGatewayMethod(this, "MyDemoMethod", {
      authorization: "NONE",
      httpMethod: "GET",
      resourceId: myDemoResource.id,
      restApiId: myDemoApi.id,
    });
  }
}

```

## Usage with Cognito User Pool Authorizer

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformVariable, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApiGatewayAuthorizer } from "./.gen/providers/aws/api-gateway-authorizer";
import { ApiGatewayMethod } from "./.gen/providers/aws/api-gateway-method";
import { ApiGatewayResource } from "./.gen/providers/aws/api-gateway-resource";
import { ApiGatewayRestApi } from "./.gen/providers/aws/api-gateway-rest-api";
import { DataAwsCognitoUserPools } from "./.gen/providers/aws/data-aws-cognito-user-pools";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
    You can read more about this at https://cdk.tf/variables*/
    const cognitoUserPoolName = new TerraformVariable(
      this,
      "cognito_user_pool_name",
      {}
    );
    const thisVar = new ApiGatewayRestApi(this, "this", {
      name: "with-authorizer",
    });
    const dataAwsCognitoUserPoolsThis = new DataAwsCognitoUserPools(
      this,
      "this_2",
      {
        name: cognitoUserPoolName.stringValue,
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataAwsCognitoUserPoolsThis.overrideLogicalId("this");
    const awsApiGatewayAuthorizerThis = new ApiGatewayAuthorizer(
      this,
      "this_3",
      {
        name: "CognitoUserPoolAuthorizer",
        providerArns: Token.asList(dataAwsCognitoUserPoolsThis.arns),
        restApiId: thisVar.id,
        type: "COGNITO_USER_POOLS",
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsApiGatewayAuthorizerThis.overrideLogicalId("this");
    const awsApiGatewayResourceThis = new ApiGatewayResource(this, "this_4", {
      parentId: thisVar.rootResourceId,
      pathPart: "{proxy+}",
      restApiId: thisVar.id,
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsApiGatewayResourceThis.overrideLogicalId("this");
    new ApiGatewayMethod(this, "any", {
      authorization: "COGNITO_USER_POOLS",
      authorizerId: Token.asString(awsApiGatewayAuthorizerThis.id),
      httpMethod: "ANY",
      requestParameters: {
        "method.request.path.proxy": true,
      },
      resourceId: Token.asString(awsApiGatewayResourceThis.id),
      restApiId: thisVar.id,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `restApiId` - (Required) ID of the associated REST API
* `resourceId` - (Required) API resource ID
* `httpMethod` - (Required) HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
* `authorization` - (Required) Type of authorization used for the method (`NONE`, `CUSTOM`, `AWS_IAM`, `COGNITO_USER_POOLS`)
* `authorizerId` - (Optional) Authorizer id to be used when the authorization is `CUSTOM` or `COGNITO_USER_POOLS`
* `authorizationScopes` - (Optional) Authorization scopes used when the authorization is `COGNITO_USER_POOLS`
* `apiKeyRequired` - (Optional) Specify if the method requires an API key
* `operationName` - (Optional) Function name that will be given to the method when generating an SDK through API Gateway. If omitted, API Gateway will generate a function name based on the resource path and HTTP verb.
* `requestModels` - (Optional) Map of the API models used for the request's content type
  where key is the content type (e.g., `application/json`)
  and value is either `Error`, `Empty` (built-in models) or `aws_api_gateway_model`'s `name`.
* `requestValidatorId` - (Optional) ID of a `aws_api_gateway_request_validator`
* `requestParameters` - (Optional) Map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (`true`) or optional (`false`).
  For example: `request_parameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true}` would define that the header `X-Some-Header` and the query string `some-query-param` must be provided in the request.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_api_gateway_method` using `REST-API-ID/RESOURCE-ID/HTTP-METHOD`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { ApiGatewayMethod } from "./.gen/providers/aws/api-gateway-method";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    ApiGatewayMethod.generateConfigForImport(
      this,
      "example",
      "12345abcde/67890fghij/GET"
    );
  }
}

```

Using `terraform import`, import `aws_api_gateway_method` using `REST-API-ID/RESOURCE-ID/HTTP-METHOD`. For example:

```console
% terraform import aws_api_gateway_method.example 12345abcde/67890fghij/GET
```

<!-- cache-key: cdktf-0.20.8 input-6c3f4021b757e16ab96211285003821b5d20d3bece66cc6c6288a63e869cfdde -->