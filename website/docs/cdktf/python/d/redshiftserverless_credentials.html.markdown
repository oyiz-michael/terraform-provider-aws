---
subcategory: "Redshift Serverless"
layout: "aws"
page_title: "AWS: aws_redshiftserverless_credentials"
description: |-
  Provides redshift serverless credentials
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_redshiftserverless_credentials

Provides redshift serverless temporary credentials for a workgroup.

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_redshiftserverless_credentials import DataAwsRedshiftserverlessCredentials
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsRedshiftserverlessCredentials(self, "example",
            workgroup_name=Token.as_string(aws_redshiftserverless_workgroup_example.workgroup_name)
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `workgroup_name` - (Required) The name of the workgroup associated with the database.
* `db_name` - (Optional) The name of the database to get temporary authorization to log on to.
* `duration_seconds` - (Optional) The number of seconds until the returned temporary password expires. The minimum is 900 seconds, and the maximum is 3600 seconds.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `db_password` - Temporary password that authorizes the user name returned by `db_user` to log on to the database `db_name`.
* `db_user` - A database user name that is authorized to log on to the database `db_name` using the password `db_password` . If the specified `db_user` exists in the database, the new user name has the same database privileges as the user named in `db_user` . By default, the user is added to PUBLIC. the user doesn't exist in the database.
* `expiration` - Date and time the password in `db_password` expires.

<!-- cache-key: cdktf-0.20.8 input-47410480b624f732e0717e1f47164e0aac1b6badeecbed2c7d0d15eb5565bcec -->