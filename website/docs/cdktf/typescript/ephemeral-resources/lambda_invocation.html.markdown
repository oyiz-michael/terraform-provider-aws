---
subcategory: "Lambda"
layout: "aws"
page_title: "AWS: aws_lambda_invocation"
description: |-
  Invokes an AWS Lambda Function as an ephemeral resource.
---


<!-- Please do not edit this file, it is generated. -->
# Ephemeral: aws_lambda_invocation

Invokes an AWS Lambda Function as an ephemeral resource. Use this ephemeral resource to execute Lambda functions during Terraform operations without persisting results in state, ideal for generating sensitive data or performing lightweight operations.

The Lambda function is invoked with [RequestResponse](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax) invocation type.

~> **Note:** Ephemeral resources are a new feature and may evolve as we continue to explore their most effective uses. [Learn more](https://developer.hashicorp.com/terraform/language/resources/ephemeral).

~> **Note:** The `aws_lambda_invocation` ephemeral resource invokes the function during every `plan` and `apply` when the function is known. A common use case for this functionality is when invoking a lightweight function—where repeated invocations are acceptable—that produces sensitive information you do not want to store in the state.

~> **Note:** If you get a `KMSAccessDeniedException: Lambda was unable to decrypt the environment variables because KMS access was denied` error when invoking a Lambda function with environment variables, the IAM role associated with the function may have been deleted and recreated after the function was created. You can fix the problem two ways: 1) updating the function's role to another role and then updating it back again to the recreated role, or 2) by using Terraform to `taint` the function and `apply` your configuration again to recreate the function. (When you create a function, Lambda grants permissions on the KMS key to the function's IAM role. If the IAM role is recreated, the grant is no longer valid. Changing the function's role or recreating the function causes Lambda to update the grant.)

## Example Usage

### Generate Sensitive Configuration

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import {
  VariableType,
  TerraformVariable,
  TerraformOutput,
  Fn,
  Token,
  TerraformStack,
} from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SsmParameter } from "./.gen/providers/aws/ssm-parameter";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
    You can read more about this at https://cdk.tf/variables*/
    const environment = new TerraformVariable(this, "environment", {
      description: "The environment name (e.g., dev, prod)",
      type: VariableType.STRING,
    });
    new TerraformOutput(this, "key_generated", {
      value: "API key generated and stored in Parameter Store",
    });
    new SsmParameter(this, "api_key", {
      name: "/app/${" + environment.value + "}/api-key",
      tags: {
        Environment: environment.stringValue,
        Generated: "ephemeral-lambda",
      },
      type: "SecureString",
      value: Token.asString(
        Fn.lookupNested(
          Fn.jsondecode(awsLambdaInvocation.secretGenerator.result),
          ["api_key"]
        )
      ),
    });
  }
}

```

### Dynamic Resource Configuration

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AutoscalingGroup } from "./.gen/providers/aws/autoscaling-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const sizing = Fn.jsondecode(awsLambdaInvocation.resourceCalculator.result);
    new AutoscalingGroup(this, "example", {
      desiredCapacity: Token.asNumber(
        Fn.lookupNested(sizing, ["desired_instances"])
      ),
      healthCheckType: "ELB",
      launchTemplate: {
        id: Token.asString(awsLaunchTemplateExample.id),
        version: "$Latest",
      },
      maxSize: Token.asNumber(Fn.lookupNested(sizing, ["max_instances"])),
      minSize: Token.asNumber(Fn.lookupNested(sizing, ["min_instances"])),
      name: "optimized-asg",
      tag: [
        {
          key: "OptimizedBy",
          propagateAtLaunch: true,
          value: "ephemeral-lambda",
        },
      ],
      targetGroupArns: [Token.asString(awsLbTargetGroupExample.arn)],
      vpcZoneIdentifier: subnetIds.listValue,
    });
  }
}

```

### Validation and Compliance Checks

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import {
  VariableType,
  TerraformVariable,
  conditional,
  Token,
  TerraformCount,
  Fn,
  TerraformStack,
} from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Resource } from "./.gen/providers/null/resource";
import { Instance } from "./.gen/providers/aws/instance";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*Terraform Variables are not always the best fit for getting inputs in the context of Terraform CDK.
    You can read more about this at https://cdk.tf/variables*/
    const instanceType = new TerraformVariable(this, "instance_type", {
      description: "The EC2 instance type to use",
      type: VariableType.STRING,
    });
    const isCompliant = compliant;
    const violations = validationResultViolations;
    /*In most cases loops should be handled in the programming language context and 
    not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
    you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
    you need to keep this like it is.*/
    const exampleCount = TerraformCount.of(
      Token.asNumber(conditional(isCompliant, 1, 0))
    );
    new Instance(this, "example", {
      ami: Token.asString(dataAwsAmiExample.id),
      instanceType: instanceType.stringValue,
      rootBlockDevice: {
        encrypted: encryptStorage.booleanValue,
      },
      tags: {
        ComplianceCheck: "passed",
        Environment: environment.stringValue,
      },
      count: exampleCount,
    });
    /*In most cases loops should be handled in the programming language context and 
    not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
    you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
    you need to keep this like it is.*/
    const complianceGateCount = TerraformCount.of(
      Token.asNumber(conditional(isCompliant, 0, 1))
    );
    new Resource(this, "compliance_gate", {
      count: complianceGateCount,
      provisioners: [
        {
          type: "local-exec",
          command:
            "echo 'Compliance violations: " +
            Token.asString(Fn.join(", ", Token.asList(violations))) +
            "' && exit 1",
        },
      ],
    });
  }
}

```

### External API Integration

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsService } from "./.gen/providers/aws/ecs-service";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const externalConfig = Fn.jsondecode(
      awsLambdaInvocation.externalConfig.result
    );
    new EcsService(this, "example", {
      cluster: Token.asString(awsEcsClusterExample.id),
      deploymentConfiguration: {
        maximum_percent: Fn.lookupNested(externalConfig, ["max_percent"]),
        minimum_healthy_percent: Fn.lookupNested(externalConfig, [
          "min_healthy_percent",
        ]),
      },
      desiredCount: Token.asNumber(
        Fn.lookupNested(externalConfig, ["replica_count"])
      ),
      name: "web-app",
      tags: {
        ConfigSource: "external-api",
        Environment: environment.stringValue,
      },
      taskDefinition: Token.asString(awsEcsTaskDefinitionExample.arn),
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `functionName` - (Required) Name or ARN of the Lambda function, version, or alias. You can append a version number or alias. If you specify only the function name, it is limited to 64 characters in length.
* `payload` - (Required) JSON that you want to provide to your Lambda function as input.

The following arguments are optional:

* `clientContext` - (Optional) Up to 3583 bytes of base64-encoded data about the invoking client to pass to the function in the context object.
* `logType` - (Optional) Set to `Tail` to include the execution log in the response. Valid values: `None` and `Tail`.
* `qualifier` - (Optional) Version or alias to invoke a published version of the function. Defaults to `$LATEST`.
* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).

## Attribute Reference

This ephemeral resource exports the following attributes in addition to the arguments above:

* `executed_version` - Version of the function that executed. When you invoke a function with an alias, this shows the version the alias resolved to.
* `function_error` - If present, indicates that an error occurred during function execution. Details about the error are included in `result`.
* `log_result` - Last 4 KB of the execution log, which is base64-encoded.
* `result` - String result of the Lambda function invocation.
* `statusCode` - HTTP status code is in the 200 range for a successful request.

## Usage Notes

### Handling Sensitive Data

Since ephemeral resources are designed to not persist data in state, they are ideal for handling sensitive information:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SecretsmanagerSecretVersion } from "./.gen/providers/aws/secretsmanager-secret-version";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SecretsmanagerSecretVersion(this, "example", {
      secretId: Token.asString(awsSecretsmanagerSecretExample.id),
      secretString: awsLambdaInvocation.credentials.result,
    });
  }
}

```

### Error Handling

Always check for function errors in your configuration:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import {
  Op,
  Fn,
  Token,
  conditional,
  TerraformCount,
  TerraformStack,
} from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Resource } from "./.gen/providers/null/resource";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const hasError = Op.neq(awsLambdaInvocation.example.functionError, "null");
    const invocationResult = Fn.jsondecode(awsLambdaInvocation.example.result);
    /*In most cases loops should be handled in the programming language context and 
    not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
    you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
    you need to keep this like it is.*/
    const validationCount = TerraformCount.of(
      Token.asNumber(
        conditional(
          hasError,
          fail(
            "Lambda function error: " +
              Token.asString(
                Fn.lookupNested(invocationResult, ["errorMessage"])
              )
          ),
          0
        )
      )
    );
    new Resource(this, "validation", {
      count: validationCount,
    });
  }
}

```

### Logging

Enable detailed logging for debugging:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformOutput, Fn, TerraformStack } from "cdktf";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new TerraformOutput(this, "execution_logs", {
      value: Fn.base64decode(awsLambdaInvocation.example.logResult),
    });
  }
}

```

<!-- cache-key: cdktf-0.20.8 input-be30ef92373442009830424f096e3a034510eb688acf4a1eea5533b5570b08bc -->