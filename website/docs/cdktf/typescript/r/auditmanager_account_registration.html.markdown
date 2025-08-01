---
subcategory: "Audit Manager"
layout: "aws"
page_title: "AWS: aws_auditmanager_account_registration"
description: |-
  Terraform resource for managing AWS Audit Manager Account Registration.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_auditmanager_account_registration

Terraform resource for managing AWS Audit Manager Account Registration.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AuditmanagerAccountRegistration } from "./.gen/providers/aws/auditmanager-account-registration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new AuditmanagerAccountRegistration(this, "example", {});
  }
}

```

### Deregister On Destroy

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AuditmanagerAccountRegistration } from "./.gen/providers/aws/auditmanager-account-registration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new AuditmanagerAccountRegistration(this, "example", {
      deregisterOnDestroy: true,
    });
  }
}

```

## Argument Reference

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `delegatedAdminAccount` - (Optional) Identifier for the delegated administrator account.
* `deregisterOnDestroy` - (Optional) Flag to deregister AuditManager in the account upon destruction. Defaults to `false` (ie. AuditManager will remain active in the account, even if this resource is removed).
* `kmsKey` - (Optional) KMS key identifier.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Unique identifier for the account registration. Since registration is applied per AWS region, this will be the active region name (ex. `us-east-1`).
* `status` - Status of the account registration request.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Audit Manager Account Registration resources using the `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AuditmanagerAccountRegistration } from "./.gen/providers/aws/auditmanager-account-registration";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    AuditmanagerAccountRegistration.generateConfigForImport(
      this,
      "example",
      "us-east-1"
    );
  }
}

```

Using `terraform import`, import Audit Manager Account Registration resources using the `id`. For example:

```console
% terraform import aws_auditmanager_account_registration.example us-east-1
```

<!-- cache-key: cdktf-0.20.8 input-899642bafeeb264894a258bb945473b5bac7104beced5d0da0fb5cc154992857 -->