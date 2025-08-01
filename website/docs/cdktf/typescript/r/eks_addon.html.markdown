---
subcategory: "EKS (Elastic Kubernetes)"
layout: "aws"
page_title: "AWS: aws_eks_addon"
description: |-
  Manages an EKS add-on
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_eks_addon

Manages an EKS add-on.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EksAddon } from "./.gen/providers/aws/eks-addon";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EksAddon(this, "example", {
      addonName: "vpc-cni",
      clusterName: Token.asString(awsEksClusterExample.name),
    });
  }
}

```

## Example Update add-on usage with resolve_conflicts_on_update and PRESERVE

`resolveConflictsOnUpdate` with `PRESERVE` can be used to retain the config changes applied to the add-on with kubectl while upgrading to a newer version of the add-on.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EksAddon } from "./.gen/providers/aws/eks-addon";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EksAddon(this, "example", {
      addonName: "coredns",
      addonVersion: "v1.10.1-eksbuild.1",
      clusterName: Token.asString(awsEksClusterExample.name),
      resolveConflictsOnUpdate: "PRESERVE",
    });
  }
}

```

## Example add-on usage with custom configuration_values

Custom add-on configuration can be passed using `configurationValues` as a single JSON string while creating or updating the add-on.

~> **Note:** `configurationValues` is a single JSON string should match the valid JSON schema for each add-on with specific version.

To find the correct JSON schema for each add-on can be extracted using [describe-addon-configuration](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-configuration.html) call.
This below is an example for extracting the `configurationValues` schema for `coredns`.

```bash
 aws eks describe-addon-configuration \
 --addon-name coredns \
 --addon-version v1.10.1-eksbuild.1
```

Example to create a `coredns` managed addon with custom `configurationValues`.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EksAddon } from "./.gen/providers/aws/eks-addon";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new EksAddon(this, "example", {
      addonName: "coredns",
      addonVersion: "v1.10.1-eksbuild.1",
      clusterName: "mycluster",
      configurationValues: Token.asString(
        Fn.jsonencode({
          replicaCount: 4,
          resources: {
            limits: {
              cpu: "100m",
              memory: "150Mi",
            },
            requests: {
              cpu: "100m",
              memory: "150Mi",
            },
          },
        })
      ),
      resolveConflictsOnCreate: "OVERWRITE",
    });
  }
}

```

### Example IAM Role for EKS Addon "vpc-cni" with AWS managed policy

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsIamPolicyDocument } from "./.gen/providers/aws/data-aws-iam-policy-document";
import { EksCluster } from "./.gen/providers/aws/eks-cluster";
import { IamOpenidConnectProvider } from "./.gen/providers/aws/iam-openid-connect-provider";
import { IamRole } from "./.gen/providers/aws/iam-role";
import { IamRolePolicyAttachment } from "./.gen/providers/aws/iam-role-policy-attachment";
import { DataTlsCertificate } from "./.gen/providers/tls/data-tls-certificate";
interface MyConfig {
  name: any;
  roleArn: any;
  vpcConfig: any;
}
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string, config: MyConfig) {
    super(scope, name);
    /*The following providers are missing schema information and might need manual adjustments to synthesize correctly: tls.
    For a more precise conversion please use the --provider flag in convert.*/
    const example = new EksCluster(this, "example", {
      name: config.name,
      roleArn: config.roleArn,
      vpcConfig: config.vpcConfig,
    });
    const dataTlsCertificateExample = new DataTlsCertificate(
      this,
      "example_1",
      {
        url: Fn.lookupNested(example.identity, ["0", "oidc", "0", "issuer"]),
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    dataTlsCertificateExample.overrideLogicalId("example");
    const awsIamOpenidConnectProviderExample = new IamOpenidConnectProvider(
      this,
      "example_2",
      {
        clientIdList: ["sts.amazonaws.com"],
        thumbprintList: [
          Token.asString(
            Fn.lookupNested(dataTlsCertificateExample.certificates, [
              "0",
              "sha1_fingerprint",
            ])
          ),
        ],
        url: Token.asString(
          Fn.lookupNested(example.identity, ["0", "oidc", "0", "issuer"])
        ),
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsIamOpenidConnectProviderExample.overrideLogicalId("example");
    const exampleAssumeRolePolicy = new DataAwsIamPolicyDocument(
      this,
      "example_assume_role_policy",
      {
        statement: [
          {
            actions: ["sts:AssumeRoleWithWebIdentity"],
            condition: [
              {
                test: "StringEquals",
                values: ["system:serviceaccount:kube-system:aws-node"],
                variable:
                  Token.asString(
                    Fn.replace(
                      Token.asString(awsIamOpenidConnectProviderExample.url),
                      "https://",
                      ""
                    )
                  ) + ":sub",
              },
            ],
            effect: "Allow",
            principals: [
              {
                identifiers: [
                  Token.asString(awsIamOpenidConnectProviderExample.arn),
                ],
                type: "Federated",
              },
            ],
          },
        ],
      }
    );
    const awsIamRoleExample = new IamRole(this, "example_4", {
      assumeRolePolicy: Token.asString(exampleAssumeRolePolicy.json),
      name: "example-vpc-cni-role",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsIamRoleExample.overrideLogicalId("example");
    const awsIamRolePolicyAttachmentExample = new IamRolePolicyAttachment(
      this,
      "example_5",
      {
        policyArn: "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
        role: Token.asString(awsIamRoleExample.name),
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsIamRolePolicyAttachmentExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `addonName` - (Required) Name of the EKS add-on. The name must match one of
  the names returned by [describe-addon-versions](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-versions.html).
* `clusterName` - (Required) Name of the EKS Cluster.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `addonVersion` - (Optional) The version of the EKS add-on. The version must
  match one of the versions returned by [describe-addon-versions](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-versions.html).
* `configurationValues` - (Optional) custom configuration values for addons with single JSON string. This JSON string value must match the JSON schema derived from [describe-addon-configuration](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-configuration.html).
* `resolveConflictsOnCreate` - (Optional) How to resolve field value conflicts when migrating a self-managed add-on to an Amazon EKS add-on. Valid values are `NONE` and `OVERWRITE`. For more details see the [CreateAddon](https://docs.aws.amazon.com/eks/latest/APIReference/API_CreateAddon.html) API Documentation.
* `resolveConflictsOnUpdate` - (Optional) How to resolve field value conflicts for an Amazon EKS add-on if you've changed a value from the Amazon EKS default value. Valid values are `NONE`, `OVERWRITE`, and `PRESERVE`. For more details see the [UpdateAddon](https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateAddon.html) API Documentation.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `podIdentityAssociation` - (Optional) Configuration block with EKS Pod Identity association settings. See [`podIdentityAssociation`](#pod-identity-association) below for details.
* `preserve` - (Optional) Indicates if you want to preserve the created resources when deleting the EKS add-on.
* `serviceAccountRoleArn` - (Optional) The Amazon Resource Name (ARN) of an
  existing IAM role to bind to the add-on's service account. The role must be
  assigned the IAM permissions required by the add-on. If you don't specify
  an existing IAM role, then the add-on uses the permissions assigned to the node
  IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)
  in the Amazon EKS User Guide.

  ~> **Note:** To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC)
  provider created for your cluster. For more information, [see Enabling IAM roles
  for service accounts on your cluster](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html)
  in the Amazon EKS User Guide.

### pod_identity_association

* `roleArn` - (Required) The Amazon Resource Name (ARN) of the IAM role to associate with the service account. The EKS Pod Identity agent manages credentials to assume this role for applications in the containers in the pods that use this service account.
* `serviceAccount` - (Required) The name of the Kubernetes service account inside the cluster to associate the IAM credentials with.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the EKS add-on.
* `id` - EKS Cluster name and EKS Addon name separated by a colon (`:`).
* `status` - Status of the EKS add-on.
* `createdAt` - Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was created.
* `modifiedAt` - Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the EKS add-on was updated.
* `tagsAll` - (Optional) Key-value map of resource tags, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `20m`)
* `update` - (Default `20m`)
* `delete` - (Default `40m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import EKS add-on using the `clusterName` and `addonName` separated by a colon (`:`). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { EksAddon } from "./.gen/providers/aws/eks-addon";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    EksAddon.generateConfigForImport(
      this,
      "myEksAddon",
      "my_cluster_name:my_addon_name"
    );
  }
}

```

Using `terraform import`, import EKS add-on using the `clusterName` and `addonName` separated by a colon (`:`). For example:

```console
% terraform import aws_eks_addon.my_eks_addon my_cluster_name:my_addon_name
```

<!-- cache-key: cdktf-0.20.8 input-c73f8260cf457fff5cfd8a0b49569977fbcb4f2b70abbf416d11d4bf1c4cbc84 -->