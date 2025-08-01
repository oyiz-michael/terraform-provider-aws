---
subcategory: "API Gateway V2"
layout: "aws"
page_title: "AWS: aws_apigatewayv2_route"
description: |-
  Manages an Amazon API Gateway Version 2 route.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_apigatewayv2_route

Manages an Amazon API Gateway Version 2 route.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/welcome.html) for [WebSocket](https://docs.aws.amazon.com/apigateway/latest/developerguide/websocket-api-develop-routes.html) and [HTTP](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-routes.html) APIs.

## Example Usage

### Basic

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.apigatewayv2_api import Apigatewayv2Api
from imports.aws.apigatewayv2_route import Apigatewayv2Route
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = Apigatewayv2Api(self, "example",
            name="example-websocket-api",
            protocol_type="WEBSOCKET",
            route_selection_expression="$request.body.action"
        )
        aws_apigatewayv2_route_example = Apigatewayv2Route(self, "example_1",
            api_id=example.id,
            route_key="$default"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_apigatewayv2_route_example.override_logical_id("example")
```

### HTTP Proxy Integration

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.apigatewayv2_api import Apigatewayv2Api
from imports.aws.apigatewayv2_integration import Apigatewayv2Integration
from imports.aws.apigatewayv2_route import Apigatewayv2Route
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = Apigatewayv2Api(self, "example",
            name="example-http-api",
            protocol_type="HTTP"
        )
        aws_apigatewayv2_integration_example = Apigatewayv2Integration(self, "example_1",
            api_id=example.id,
            integration_method="ANY",
            integration_type="HTTP_PROXY",
            integration_uri="https://example.com/{proxy}"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_apigatewayv2_integration_example.override_logical_id("example")
        aws_apigatewayv2_route_example = Apigatewayv2Route(self, "example_2",
            api_id=example.id,
            route_key="ANY /example/{proxy+}",
            target="integrations/${" + aws_apigatewayv2_integration_example.id + "}"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_apigatewayv2_route_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `api_id` - (Required) API identifier.
* `route_key` - (Required) Route key for the route. For HTTP APIs, the route key can be either `$default`, or a combination of an HTTP method and resource path, for example, `GET /pets`.
* `api_key_required` - (Optional) Boolean whether an API key is required for the route. Defaults to `false`. Supported only for WebSocket APIs.
* `authorization_scopes` - (Optional) Authorization scopes supported by this route. The scopes are used with a JWT authorizer to authorize the method invocation.
* `authorization_type` - (Optional) Authorization type for the route.
For WebSocket APIs, valid values are `NONE` for open access, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
For HTTP APIs, valid values are `NONE` for open access, `JWT` for using JSON Web Tokens, `AWS_IAM` for using AWS IAM permissions, and `CUSTOM` for using a Lambda authorizer.
Defaults to `NONE`.
* `authorizer_id` - (Optional) Identifier of the [`aws_apigatewayv2_authorizer`](apigatewayv2_authorizer.html) resource to be associated with this route.
* `model_selection_expression` - (Optional) The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route. Supported only for WebSocket APIs.
* `operation_name` - (Optional) Operation name for the route. Must be between 1 and 64 characters in length.
* `request_models` - (Optional) Request models for the route. Supported only for WebSocket APIs.
* `request_parameter` - (Optional) Request parameters for the route. Supported only for WebSocket APIs.
* `route_response_selection_expression` - (Optional) The [route response selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-route-response-selection-expressions) for the route. Supported only for WebSocket APIs.
* `target` - (Optional) Target for the route, of the form `integrations/`*`IntegrationID`*, where *`IntegrationID`* is the identifier of an [`aws_apigatewayv2_integration`](apigatewayv2_integration.html) resource.

The `request_parameter` object supports the following:

* `request_parameter_key` - (Required) Request parameter key. This is a [request data mapping parameter](https://docs.aws.amazon.com/apigateway/latest/developerguide/websocket-api-data-mapping.html#websocket-mapping-request-parameters).
* `required` - (Required) Boolean whether or not the parameter is required.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Route identifier.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_apigatewayv2_route` using the API identifier and route identifier. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.apigatewayv2_route import Apigatewayv2Route
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        Apigatewayv2Route.generate_config_for_import(self, "example", "aabbccddee/1122334")
```

Using `terraform import`, import `aws_apigatewayv2_route` using the API identifier and route identifier. For example:

```console
% terraform import aws_apigatewayv2_route.example aabbccddee/1122334
```

-> **Note:** The API Gateway managed route created as part of [_quick_create_](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-basic-concept.html#apigateway-definition-quick-create) cannot be imported.

<!-- cache-key: cdktf-0.20.8 input-8f0637a4b041eb65141ad0c40e0e6e6bd1052a535207afbc86d54073bfc472b5 -->