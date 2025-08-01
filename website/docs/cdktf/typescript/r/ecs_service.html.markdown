---
subcategory: "ECS (Elastic Container)"
layout: "aws"
page_title: "AWS: aws_ecs_service"
description: |-
  Provides an ECS service.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_ecs_service

-> **Note:** To prevent a race condition during service deletion, make sure to set `dependsOn` to the related `aws_iam_role_policy`; otherwise, the policy may be destroyed too soon and the ECS service will then get stuck in the `DRAINING` state.

Provides an ECS service - effectively a task that is expected to run until an error occurs or a user terminates it (typically a webserver or a database).

See [ECS Services section in AWS developer guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsService } from "./.gen/providers/aws/ecs-service";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EcsService(this, "mongo", {
      cluster: foo.id,
      dependsOn: [awsIamRolePolicyFoo],
      desiredCount: 3,
      iamRole: Token.asString(awsIamRoleFoo.arn),
      loadBalancer: [
        {
          containerName: "mongo",
          containerPort: 8080,
          targetGroupArn: Token.asString(awsLbTargetGroupFoo.arn),
        },
      ],
      name: "mongodb",
      orderedPlacementStrategy: [
        {
          field: "cpu",
          type: "binpack",
        },
      ],
      placementConstraints: [
        {
          expression:
            "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]",
          type: "memberOf",
        },
      ],
      taskDefinition: Token.asString(awsEcsTaskDefinitionMongo.arn),
    });
  }
}

```

### Ignoring Changes to Desired Count

You can utilize the generic Terraform resource [lifecycle configuration block](https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html) with `ignore_changes` to create an ECS service with an initial count of running instances, then ignore any changes to that count caused externally (e.g., Application Autoscaling).

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsService } from "./.gen/providers/aws/ecs-service";
interface MyConfig {
  name: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new EcsService(this, "example", {
      desiredCount: 2,
      lifecycle: {
        ignoreChanges: [desiredCount],
      },
      name: config.name,
    });
  }
}

```

### Daemon Scheduling Strategy

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsService } from "./.gen/providers/aws/ecs-service";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EcsService(this, "bar", {
      cluster: foo.id,
      name: "bar",
      schedulingStrategy: "DAEMON",
      taskDefinition: Token.asString(awsEcsTaskDefinitionBar.arn),
    });
  }
}

```

### CloudWatch Deployment Alarms

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsService } from "./.gen/providers/aws/ecs-service";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EcsService(this, "example", {
      alarms: {
        alarmNames: [Token.asString(awsCloudwatchMetricAlarmExample.alarmName)],
        enable: true,
        rollback: true,
      },
      cluster: Token.asString(awsEcsClusterExample.id),
      name: "example",
    });
  }
}

```

### External Deployment Controller

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsService } from "./.gen/providers/aws/ecs-service";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EcsService(this, "example", {
      cluster: Token.asString(awsEcsClusterExample.id),
      deploymentController: {
        type: "EXTERNAL",
      },
      name: "example",
    });
  }
}

```

### Redeploy Service On Every Apply

The key used with `triggers` is arbitrary.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsService } from "./.gen/providers/aws/ecs-service";
interface MyConfig {
  name: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    new EcsService(this, "example", {
      forceNewDeployment: true,
      triggers: {
        redeployment: Token.asString(Fn.plantimestamp()),
      },
      name: config.name,
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the service (up to 255 letters, numbers, hyphens, and underscores)

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `alarms` - (Optional) Information about the CloudWatch alarms. [See below](#alarms).
* `availabilityZoneRebalancing` - (Optional) ECS automatically redistributes tasks within a service across Availability Zones (AZs) to mitigate the risk of impaired application availability due to underlying infrastructure failures and task lifecycle activities. The valid values are `ENABLED` and `DISABLED`. Defaults to `DISABLED`.
* `capacityProviderStrategy` - (Optional) Capacity provider strategies to use for the service. Can be one or more. These can be updated without destroying and recreating the service only if `force_new_deployment = true` and not changing from 0 `capacityProviderStrategy` blocks to greater than 0, or vice versa. [See below](#capacity_provider_strategy). Conflicts with `launchType`.
* `cluster` - (Optional) ARN of an ECS cluster.
* `deploymentCircuitBreaker` - (Optional) Configuration block for deployment circuit breaker. [See below](#deployment_circuit_breaker).
* `deploymentConfiguration` - (Optional) Configuration block for deployment settings. [See below](#deployment_configuration).
* `deploymentController` - (Optional) Configuration block for deployment controller configuration. [See below](#deployment_controller).
* `deploymentMaximumPercent` - (Optional) Upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the `DAEMON` scheduling strategy.
* `deploymentMinimumHealthyPercent` - (Optional) Lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
* `desiredCount` - (Optional) Number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the `DAEMON` scheduling strategy.
* `enableEcsManagedTags` - (Optional) Whether to enable Amazon ECS managed tags for the tasks within the service.
* `enableExecuteCommand` - (Optional) Whether to enable Amazon ECS Exec for the tasks within the service.
* `forceDelete` - (Optional) Enable to delete a service even if it wasn't scaled down to zero tasks. It's only necessary to use this if the service uses the `REPLICA` scheduling strategy.
* `forceNewDeployment` - (Optional) Enable to force a new task deployment of the service. This can be used to update tasks to use a newer Docker image with same image/tag combination (e.g., `myimage:latest`), roll Fargate tasks onto a newer platform version, or immediately deploy `orderedPlacementStrategy` and `placementConstraints` updates.
* `healthCheckGracePeriodSeconds` - (Optional) Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
* `iamRole` - (Optional) ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the `awsvpc` network mode. If using `awsvpc` network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
* `launchType` - (Optional) Launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`. Conflicts with `capacityProviderStrategy`.
* `loadBalancer` - (Optional) Configuration block for load balancers. [See below](#load_balancer).
* `networkConfiguration` - (Optional) Network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. [See below](#network_configuration).
* `orderedPlacementStrategy` - (Optional) Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. The maximum number of `orderedPlacementStrategy` blocks is `5`. [See below](#ordered_placement_strategy).
* `placementConstraints` - (Optional) Rules that are taken into consideration during task placement. Updates to this configuration will take effect next task deployment unless `forceNewDeployment` is enabled. Maximum number of `placementConstraints` is `10`. [See below](#placement_constraints).
* `platformVersion` - (Optional) Platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
* `propagateTags` - (Optional) Whether to propagate the tags from the task definition or the service to the tasks. The valid values are `SERVICE` and `TASK_DEFINITION`.
* `schedulingStrategy` - (Optional) Scheduling strategy to use for the service. The valid values are `REPLICA` and `DAEMON`. Defaults to `REPLICA`. Note that [*Tasks using the Fargate launch type or the `CODE_DEPLOY` or `EXTERNAL` deployment controller types don't support the `DAEMON` scheduling strategy*](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html).
* `serviceConnectConfiguration` - (Optional) ECS Service Connect configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace. [See below](#service_connect_configuration).
* `serviceRegistries` - (Optional) Service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`. [See below](#service_registries).
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `taskDefinition` - (Optional) Family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service. Required unless using the `EXTERNAL` deployment controller. If a revision is not specified, the latest `ACTIVE` revision is used.
* `triggers` - (Optional) Map of arbitrary keys and values that, when changed, will trigger an in-place update (redeployment). Useful with `plantimestamp()`. See example above.
* `volumeConfiguration` - (Optional) Configuration for a volume specified in the task definition as a volume that is configured at launch time. Currently, the only supported volume type is an Amazon EBS volume. [See below](#volume_configuration).
* `vpcLatticeConfigurations` - (Optional) The VPC Lattice configuration for your service that allows Lattice to connect, secure, and monitor your service across multiple accounts and VPCs. [See below](#vpc_lattice_configurations).
* `waitForSteadyState` - (Optional) If `true`, Terraform will wait for the service to reach a steady state (like [`aws ecs wait services-stable`](https://docs.aws.amazon.com/cli/latest/reference/ecs/wait/services-stable.html)) before continuing. Default `false`.

### alarms

The `alarms` configuration block supports the following:

* `alarmNames` - (Required) One or more CloudWatch alarm names.
* `enable` - (Required) Whether to use the CloudWatch alarm option in the service deployment process.
* `rollback` - (Required) Whether to configure Amazon ECS to roll back the service if a service deployment fails. If rollback is used, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.

### volume_configuration

The `volumeConfiguration` configuration block supports the following:

* `name` - (Required) Name of the volume.
* `managedEbsVolume` - (Required) Configuration for the Amazon EBS volume that Amazon ECS creates and manages on your behalf. [See below](#managed_ebs_volume).

### vpc_lattice_configurations

`vpcLatticeConfigurations` supports the following:

* `roleArn` - (Required) The ARN of the IAM role to associate with this volume. This is the Amazon ECS infrastructure IAM role that is used to manage your AWS infrastructure.
* `targetGroupArn` - (Required) The full ARN of the target group or groups associated with the VPC Lattice configuration.
* `portName` - (Required) The name of the port for a target group associated with the VPC Lattice configuration.

### managed_ebs_volume

The `managedEbsVolume` configuration block supports the following:

* `roleArn` - (Required) Amazon ECS infrastructure IAM role that is used to manage your Amazon Web Services infrastructure. Recommended using the Amazon ECS-managed `AmazonECSInfrastructureRolePolicyForVolumes` IAM policy with this role.
* `encrypted` - (Optional) Whether the volume should be encrypted. Default value is `true`.
* `fileSystemType` - (Optional) Linux filesystem type for the volume. For volumes created from a snapshot, same filesystem type must be specified that the volume was using when the snapshot was created. Valid values are `ext3`, `ext4`, `xfs`. Default value is `xfs`.
* `iops` - (Optional) Number of I/O operations per second (IOPS).
* `kmsKeyId` - (Optional) Amazon Resource Name (ARN) identifier of the Amazon Web Services Key Management Service key to use for Amazon EBS encryption.
* `sizeInGb` - (Optional) Size of the volume in GiB. You must specify either a `sizeInGb` or a `snapshotId`. You can optionally specify a volume size greater than or equal to the snapshot size.
* `snapshotId` - (Optional) Snapshot that Amazon ECS uses to create the volume. You must specify either a `sizeInGb` or a `snapshotId`.
* `throughput` - (Optional) Throughput to provision for a volume, in MiB/s, with a maximum of 1,000 MiB/s.
* `volumeInitializationRate` - (Optional) Volume Initialization Rate in MiB/s. You must also specify a `snapshotId`.
* `volumeType` - (Optional) Volume type.
* `tagSpecifications` - (Optional) The tags to apply to the volume. [See below](#tag_specifications).

### capacity_provider_strategy

The `capacityProviderStrategy` configuration block supports the following:

* `base` - (Optional) Number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined.
* `capacityProvider` - (Required) Short name of the capacity provider.
* `weight` - (Required) Relative percentage of the total number of launched tasks that should use the specified capacity provider.

### deployment_configuration

The `deploymentConfiguration` configuration block supports the following:

* `strategy` - (Optional) Type of deployment strategy. Valid values: `ROLLING`, `BLUE_GREEN`. Default: `ROLLING`.
* `bakeTimeInMinutes` - (Optional) Number of minutes to wait after a new deployment is fully provisioned before terminating the old deployment. Only used when `strategy` is set to `BLUE_GREEN`.
* `lifecycleHook` - (Optional) Configuration block for lifecycle hooks that are invoked during deployments. [See below](#lifecycle_hook).

### lifecycle_hook

The `lifecycleHook` configuration block supports the following:

* `hookTargetArn` - (Required) ARN of the Lambda function to invoke for the lifecycle hook.
* `roleArn` - (Required) ARN of the IAM role that grants the service permission to invoke the Lambda function.
* `lifecycleStages` - (Required) Stages during the deployment when the hook should be invoked. Valid values: `RECONCILE_SERVICE`, `PRE_SCALE_UP`, `POST_SCALE_UP`, `TEST_TRAFFIC_SHIFT`, `POST_TEST_TRAFFIC_SHIFT`, `PRODUCTION_TRAFFIC_SHIFT`, `POST_PRODUCTION_TRAFFIC_SHIFT`.

### deployment_circuit_breaker

The `deploymentCircuitBreaker` configuration block supports the following:

* `enable` - (Required) Whether to enable the deployment circuit breaker logic for the service.
* `rollback` - (Required) Whether to enable Amazon ECS to roll back the service if a service deployment fails. If rollback is enabled, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.

### deployment_controller

The `deploymentController` configuration block supports the following:

* `type` - (Optional) Type of deployment controller. Valid values: `CODE_DEPLOY`, `ECS`, `EXTERNAL`. Default: `ECS`.

### load_balancer

`loadBalancer` supports the following:

* `elbName` - (Required for ELB Classic) Name of the ELB (Classic) to associate with the service.
* `targetGroupArn` - (Required for ALB/NLB) ARN of the Load Balancer target group to associate with the service.
* `containerName` - (Required) Name of the container to associate with the load balancer (as it appears in a container definition).
* `containerPort` - (Required) Port on the container to associate with the load balancer.
* `advancedConfiguration` - (Optional) Configuration block for Blue/Green deployment settings. Required when using `BLUE_GREEN` deployment strategy. [See below](#advanced_configuration).

-> **Version note:** Multiple `loadBalancer` configuration block support was added in Terraform AWS Provider version 2.22.0. This allows configuration of [ECS service support for multiple target groups](https://aws.amazon.com/about-aws/whats-new/2019/07/amazon-ecs-services-now-support-multiple-load-balancer-target-groups/).

### advanced_configuration

The `advancedConfiguration` configuration block supports the following:

* `alternateTargetGroupArn` - (Required) ARN of the alternate target group to use for Blue/Green deployments.
* `productionListenerRule` - (Required) ARN of the listener rule that routes production traffic.
* `roleArn` - (Required) ARN of the IAM role that allows ECS to manage the target groups.
* `testListenerRule` - (Optional) ARN of the listener rule that routes test traffic.

### network_configuration

`networkConfiguration` support the following:

* `subnets` - (Required) Subnets associated with the task or service.
* `securityGroups` - (Optional) Security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
* `assignPublicIp` - (Optional) Assign a public IP address to the ENI (Fargate launch type only). Valid values are `true` or `false`. Default `false`.

For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html)

### ordered_placement_strategy

`orderedPlacementStrategy` supports the following:

* `type` - (Required) Type of placement strategy. Must be one of: `binpack`, `random`, or `spread`
* `field` - (Optional) For the `spread` placement strategy, valid values are `instanceId` (or `host`,
 which has the same effect), or any platform or custom attribute that is applied to a container instance.
 For the `binpack` type, valid values are `memory` and `cpu`. For the `random` type, this attribute is not
 needed. For more information, see [Placement Strategy](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PlacementStrategy.html).

-> **Note:** for `spread`, `host` and `instanceId` will be normalized, by AWS, to be `instanceId`. This means the statefile will show `instanceId` but your config will differ if you use `host`.

### placement_constraints

`placementConstraints` support the following:

* `type` - (Required) Type of constraint. The only valid values at this time are `memberOf` and `distinctInstance`.
* `expression` -  (Optional) Cluster Query Language expression to apply to the constraint. Does not need to be specified for the `distinctInstance` type. For more information, see [Cluster Query Language in the Amazon EC2 Container Service Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).

### service_registries

`serviceRegistries` support the following:

* `registryArn` - (Required) ARN of the Service Registry. The currently supported service registry is Amazon Route 53 Auto Naming Service(`aws_service_discovery_service`). For more information, see [Service](https://docs.aws.amazon.com/Route53/latest/APIReference/API_autonaming_Service.html)
* `port` - (Optional) Port value used if your Service Discovery service specified an SRV record.
* `containerPort` - (Optional) Port value, already specified in the task definition, to be used for your service discovery service.
* `containerName` - (Optional) Container name value, already specified in the task definition, to be used for your service discovery service.

### service_connect_configuration

`serviceConnectConfiguration` supports the following:

* `enabled` - (Required) Whether to use Service Connect with this service.
* `logConfiguration` - (Optional) Log configuration for the container. [See below](#log_configuration).
* `namespace` - (Optional) Namespace name or ARN of the [`aws_service_discovery_http_namespace`](/docs/providers/aws/r/service_discovery_http_namespace.html) for use with Service Connect.
* `service` - (Optional) List of Service Connect service objects. [See below](#service).

### log_configuration

`logConfiguration` supports the following:

* `logDriver` - (Required) Log driver to use for the container.
* `options` - (Optional) Configuration options to send to the log driver.
* `secretOption` - (Optional) Secrets to pass to the log configuration. [See below](#secret_option).

### secret_option

`secretOption` supports the following:

* `name` - (Required) Name of the secret.
* `valueFrom` - (Required) Secret to expose to the container. The supported values are either the full ARN of the AWS Secrets Manager secret or the full ARN of the parameter in the SSM Parameter Store.

### service

`service` supports the following:

* `clientAlias` - (Optional) List of client aliases for this Service Connect service. You use these to assign names that can be used by client applications. The maximum number of client aliases that you can have in this list is 1. [See below](#client_alias).
* `discoveryName` - (Optional) Name of the new AWS Cloud Map service that Amazon ECS creates for this Amazon ECS service.
* `ingressPortOverride` - (Optional) Port number for the Service Connect proxy to listen on.
* `portName` - (Required) Name of one of the `portMappings` from all the containers in the task definition of this Amazon ECS service.
* `timeout` - (Optional) Configuration timeouts for Service Connect
* `tls` - (Optional) Configuration for enabling Transport Layer Security (TLS)

### timeout

`timeout` supports the following:

* `idleTimeoutSeconds` - (Optional) Amount of time in seconds a connection will stay active while idle. A value of 0 can be set to disable idleTimeout.
* `perRequestTimeoutSeconds` - (Optional) Amount of time in seconds for the upstream to respond with a complete response per request. A value of 0 can be set to disable perRequestTimeout. Can only be set when appProtocol isn't TCP.

### tls

`tls` supports the following:

* `issuerCertAuthority` - (Required) Details of the certificate authority which will issue the certificate.
* `kmsKey` - (Optional) KMS key used to encrypt the private key in Secrets Manager.
* `roleArn` - (Optional) ARN of the IAM Role that's associated with the Service Connect TLS.

### issuer_cert_authority

`issuerCertAuthority` supports the following:

* `awsPcaAuthorityArn` - (Optional) ARN of the [`aws_acmpca_certificate_authority`](/docs/providers/aws/r/acmpca_certificate_authority.html) used to create the TLS Certificates.

### client_alias

`clientAlias` supports the following:

* `dnsName` - (Optional) Name that you use in the applications of client tasks to connect to this service.
* `port` - (Required) Listening port number for the Service Connect proxy. This port is available inside of all of the tasks within the same namespace.
* `testTrafficRules` - (Optional) Configuration block for test traffic routing rules. [See below](#test_traffic_rules).

### test_traffic_rules

The `testTrafficRules` configuration block supports the following:

* `header` - (Optional) Configuration block for header-based routing rules. [See below](#header).

### header

The `header` configuration block supports the following:

* `name` - (Required) Name of the HTTP header to match.
* `value` - (Required) Configuration block for header value matching criteria. [See below](#value).

### value

The `value` configuration block supports the following:

* `exact` - (Required) Exact string value to match in the header.

### tag_specifications

`tagSpecifications` supports the following:

* `resourceType` - (Required) The type of volume resource. Valid values, `volume`.
* `propagateTags` - (Optional) Determines whether to propagate the tags from the task definition to the Amazon EBS volume.
* `tags` - (Optional) The tags applied to this Amazon EBS volume. `AmazonECSCreated` and `AmazonECSManaged` are reserved tags that can't be used.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN that identifies the service.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `20m`)
- `update` - (Default `20m`)
- `delete` - (Default `20m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ECS services using the `name` together with ecs cluster `name`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EcsService } from "./.gen/providers/aws/ecs-service";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    EcsService.generateConfigForImport(
      this,
      "imported",
      "cluster-name/service-name"
    );
  }
}

```

Using `terraform import`, import ECS services using the `name` together with ecs cluster `name`. For example:

```console
% terraform import aws_ecs_service.imported cluster-name/service-name
```

<!-- cache-key: cdktf-0.20.8 input-bfe9a345f98b2da5764c10d8ed1b4cd92f6438f71ede7cd0ad23009ed055d4d4 -->