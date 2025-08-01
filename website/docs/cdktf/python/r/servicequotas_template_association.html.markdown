---
subcategory: "Service Quotas"
layout: "aws"
page_title: "AWS: aws_servicequotas_template_association"
description: |-
  Terraform resource for managing an AWS Service Quotas Template Association.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_servicequotas_template_association

Terraform resource for managing an AWS Service Quotas Template Association.

-> Only the management account of an organization can associate Service Quota templates, and this must be done from the `us-east-1` region.

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
from imports.aws.servicequotas_template_association import ServicequotasTemplateAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ServicequotasTemplateAssociation(self, "example")
```

## Argument Reference

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `skip_destroy` - (Optional) Skip disassociating the quota increase template upon destruction. This will remove the resource from Terraform state, but leave the remote association in place.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - AWS account ID.
* `status` - Association status. Creating this resource will result in an `ASSOCIATED` status, and quota increase requests in the template are automatically applied to new AWS accounts in the organization.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Service Quotas Template Association using the `id`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.servicequotas_template_association import ServicequotasTemplateAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        ServicequotasTemplateAssociation.generate_config_for_import(self, "example", "123456789012")
```

Using `terraform import`, import Service Quotas Template Association using the `id`. For example:

```console
% terraform import aws_servicequotas_template_association.example 123456789012 
```

<!-- cache-key: cdktf-0.20.8 input-93047f31b30ffcdf122dea5bd950050889eb1befb004af3feda76bccd44be519 -->