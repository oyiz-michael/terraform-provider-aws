---
subcategory: "Managed Grafana"
layout: "aws"
page_title: "AWS: aws_grafana_workspace_service_account"
description: |-
  Terraform resource for managing an Amazon Managed Grafana Workspace Service Account.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_grafana_workspace_service_account

-> **Note:** You cannot update a service account. If you change any attribute, Terraform
will delete the current and create a new one.

Read about Service Accounts in the [Amazon Managed Grafana user guide](https://docs.aws.amazon.com/grafana/latest/userguide/service-accounts.html).

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.grafana_workspace_service_account import GrafanaWorkspaceServiceAccount
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GrafanaWorkspaceServiceAccount(self, "example",
            grafana_role="ADMIN",
            name="example-admin",
            workspace_id=Token.as_string(aws_grafana_workspace_example.id)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) A name for the service account. The name must be unique within the workspace, as it determines the ID associated with the service account.
* `grafana_role` - (Required) The permission level to use for this service account. For more information about the roles and the permissions each has, see the [User roles](https://docs.aws.amazon.com/grafana/latest/userguide/Grafana-user-roles.html) documentation.
* `workspace_id` - (Required) The Grafana workspace with which the service account is associated.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `service_account_id` - Identifier of the service account in the given Grafana workspace

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Managed Grafana Workspace Service Account using the `workspace_id` and `service_account_id` separated by a comma (`,`). For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.grafana_workspace_service_account import GrafanaWorkspaceServiceAccount
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        GrafanaWorkspaceServiceAccount.generate_config_for_import(self, "example", "g-abc12345,1")
```

Using `terraform import`, import Managed Grafana Workspace Service Account using the `workspace_id` and `service_account_id` separated by a comma (`,`). For example:

```console
% terraform import aws_grafana_workspace_service_account.example g-abc12345,1
```

<!-- cache-key: cdktf-0.20.8 input-d017e087c9a842d3a48dcf9e01824230a2a936439f32c238e2c60c550141430f -->