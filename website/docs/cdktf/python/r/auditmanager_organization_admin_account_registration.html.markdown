---
subcategory: "Audit Manager"
layout: "aws"
page_title: "AWS: aws_auditmanager_organization_admin_account_registration"
description: |-
  Terraform resource for managing AWS Audit Manager Organization Admin Account Registration.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_auditmanager_organization_admin_account_registration

Terraform resource for managing AWS Audit Manager Organization Admin Account Registration.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.auditmanager_organization_admin_account_registration import AuditmanagerOrganizationAdminAccountRegistration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AuditmanagerOrganizationAdminAccountRegistration(self, "example",
            admin_account_id="123456789012"
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `admin_account_id` - (Required) Identifier for the organization administrator account.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Identifier for the organization administrator account.
* `organization_id` - Identifier for the organization.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Audit Manager Organization Admin Account Registration using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.auditmanager_organization_admin_account_registration import AuditmanagerOrganizationAdminAccountRegistration
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AuditmanagerOrganizationAdminAccountRegistration.generate_config_for_import(self, "example", "123456789012 ")
```

Using `terraform import`, import Audit Manager Organization Admin Account Registration using the `id`. For example:

```console
% terraform import aws_auditmanager_organization_admin_account_registration.example 123456789012 
```

<!-- cache-key: cdktf-0.20.8 input-91c7acf4c230e940acb6ae534c7176427645285d63b7be072a89ba2d53dcc2ea -->