---
subcategory: "Directory Service"
layout: "aws"
page_title: "AWS: aws_directory_service_directory"
description: |-
  Provides a directory in AWS Directory Service.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_directory_service_directory

Provides a Simple or Managed Microsoft directory in AWS Directory Service.

~> **Note:** All arguments including the password and customer username will be stored in the raw state as plain-text.
[Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).

## Example Usage

### SimpleAD

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DirectoryServiceDirectory } from "./.gen/providers/aws/directory-service-directory";
import { Subnet } from "./.gen/providers/aws/subnet";
import { Vpc } from "./.gen/providers/aws/vpc";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const main = new Vpc(this, "main", {
      cidrBlock: "10.0.0.0/16",
    });
    const bar = new Subnet(this, "bar", {
      availabilityZone: "us-west-2b",
      cidrBlock: "10.0.2.0/24",
      vpcId: main.id,
    });
    const foo = new Subnet(this, "foo", {
      availabilityZone: "us-west-2a",
      cidrBlock: "10.0.1.0/24",
      vpcId: main.id,
    });
    const awsDirectoryServiceDirectoryBar = new DirectoryServiceDirectory(
      this,
      "bar_3",
      {
        name: "corp.notexample.com",
        password: "SuperSecretPassw0rd",
        size: "Small",
        tags: {
          Project: "foo",
        },
        vpcSettings: {
          subnetIds: [foo.id, bar.id],
          vpcId: main.id,
        },
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsDirectoryServiceDirectoryBar.overrideLogicalId("bar");
  }
}

```

### Microsoft Active Directory (MicrosoftAD)

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DirectoryServiceDirectory } from "./.gen/providers/aws/directory-service-directory";
import { Subnet } from "./.gen/providers/aws/subnet";
import { Vpc } from "./.gen/providers/aws/vpc";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const main = new Vpc(this, "main", {
      cidrBlock: "10.0.0.0/16",
    });
    const bar = new Subnet(this, "bar", {
      availabilityZone: "us-west-2b",
      cidrBlock: "10.0.2.0/24",
      vpcId: main.id,
    });
    const foo = new Subnet(this, "foo", {
      availabilityZone: "us-west-2a",
      cidrBlock: "10.0.1.0/24",
      vpcId: main.id,
    });
    const awsDirectoryServiceDirectoryBar = new DirectoryServiceDirectory(
      this,
      "bar_3",
      {
        edition: "Standard",
        name: "corp.notexample.com",
        password: "SuperSecretPassw0rd",
        tags: {
          Project: "foo",
        },
        type: "MicrosoftAD",
        vpcSettings: {
          subnetIds: [foo.id, bar.id],
          vpcId: main.id,
        },
      }
    );
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsDirectoryServiceDirectoryBar.overrideLogicalId("bar");
  }
}

```

### Microsoft Active Directory Connector (ADConnector)

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DirectoryServiceDirectory } from "./.gen/providers/aws/directory-service-directory";
import { Subnet } from "./.gen/providers/aws/subnet";
import { Vpc } from "./.gen/providers/aws/vpc";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const main = new Vpc(this, "main", {
      cidrBlock: "10.0.0.0/16",
    });
    const bar = new Subnet(this, "bar", {
      availabilityZone: "us-west-2b",
      cidrBlock: "10.0.2.0/24",
      vpcId: main.id,
    });
    const foo = new Subnet(this, "foo", {
      availabilityZone: "us-west-2a",
      cidrBlock: "10.0.1.0/24",
      vpcId: main.id,
    });
    new DirectoryServiceDirectory(this, "connector", {
      connectSettings: {
        customerDnsIps: ["A.B.C.D"],
        customerUsername: "Admin",
        subnetIds: [foo.id, bar.id],
        vpcId: main.id,
      },
      name: "corp.notexample.com",
      password: "SuperSecretPassw0rd",
      size: "Small",
      type: "ADConnector",
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) The fully qualified name for the directory, such as `corp.example.com`
* `password` - (Required) The password for the directory administrator or connector user.
* `size` - (Optional) (For `SimpleAD` and `ADConnector` types) The size of the directory (`Small` or `Large` are accepted values). `Large` by default.
* `vpcSettings` - (Required for `SimpleAD` and `MicrosoftAD`) VPC related information about the directory. Fields documented below.
* `connectSettings` - (Required for `ADConnector`) Connector related information about the directory. Fields documented below.
* `alias` - (Optional) The alias for the directory (must be unique amongst all aliases in AWS). Required for `enableSso`.
* `description` - (Optional) A textual description for the directory.
* `desiredNumberOfDomainControllers` - (Optional) The number of domain controllers desired in the directory. Minimum value of `2`. Scaling of domain controllers is only supported for `MicrosoftAD` directories.
* `shortName` - (Optional) The short name of the directory, such as `CORP`.
* `enableSso` - (Optional) Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
* `type` (Optional) - The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
* `edition` - (Optional, for type `MicrosoftAD` only) The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise`.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

**vpc_settings** supports the following:

* `subnetIds` - (Required) The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
* `vpcId` - (Required) The identifier of the VPC that the directory is in.

**connect_settings** supports the following:

* `customerUsername` - (Required) The username corresponding to the password provided.
* `customerDnsIps` - (Required) The DNS IP addresses of the domain to connect to.
* `subnetIds` - (Required) The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
* `vpcId` - (Required) The identifier of the VPC that the directory is in.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The directory identifier.
* `accessUrl` - The access URL for the directory, such as `http://alias.awsapps.com`.
* `dnsIpAddresses` - A list of IP addresses of the DNS servers for the directory or connector.
* `securityGroupId` - The ID of the security group created by the directory.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

`connectSettings` (for `ADConnector`) is also exported with the following attributes:

* `connectIps` - The IP addresses of the AD Connector servers.

## Timeouts

`aws_directory_service_directory` provides the following [Timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) configuration options:

- `create` - (Default `60 minutes`) Used for directory creation
- `update` - (Default `60 minutes`) Used for directory update
- `delete` - (Default `60 minutes`) Used for directory deletion

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DirectoryService directories using the directory `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DirectoryServiceDirectory } from "./.gen/providers/aws/directory-service-directory";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    DirectoryServiceDirectory.generateConfigForImport(
      this,
      "sample",
      "d-926724cf57"
    );
  }
}

```

Using `terraform import`, import DirectoryService directories using the directory `id`. For example:

```console
% terraform import aws_directory_service_directory.sample d-926724cf57
```

<!-- cache-key: cdktf-0.20.8 input-c2aecef0a364e7156d19cfe8a226a33940e1e3c04ceae1a86c7841b7e7b69f47 -->