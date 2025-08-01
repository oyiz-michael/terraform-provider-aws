---
subcategory: "Redshift Serverless"
layout: "aws"
page_title: "AWS: aws_redshiftserverless_custom_domain_association"
description: |-
  Terraform resource for managing an AWS Redshift Serverless Custom Domain Association.
---

<!-- Please do not edit this file, it is generated. -->
# Resource: aws_redshiftserverless_custom_domain_association

Terraform resource for managing an AWS Redshift Serverless Custom Domain Association.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AcmCertificate } from "./.gen/providers/aws/acm-certificate";
import { RedshiftserverlessCustomDomainAssociation } from "./.gen/providers/aws/redshiftserverless-custom-domain-association";
import { RedshiftserverlessNamespace } from "./.gen/providers/aws/redshiftserverless-namespace";
import { RedshiftserverlessWorkgroup } from "./.gen/providers/aws/redshiftserverless-workgroup";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new AcmCertificate(this, "example", {
      domainName: "example.com",
    });
    const awsRedshiftserverlessNamespaceExample =
      new RedshiftserverlessNamespace(this, "example_1", {
        namespaceName: "example-namespace",
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsRedshiftserverlessNamespaceExample.overrideLogicalId("example");
    const awsRedshiftserverlessWorkgroupExample =
      new RedshiftserverlessWorkgroup(this, "example_2", {
        namespaceName: Token.asString(
          awsRedshiftserverlessNamespaceExample.namespaceName
        ),
        workgroupName: "example-workgroup",
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsRedshiftserverlessWorkgroupExample.overrideLogicalId("example");
    const awsRedshiftserverlessCustomDomainAssociationExample =
      new RedshiftserverlessCustomDomainAssociation(this, "example_3", {
        customDomainCertificateArn: example.arn,
        customDomainName: "example.com",
        workgroupName: Token.asString(
          awsRedshiftserverlessWorkgroupExample.workgroupName
        ),
      });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsRedshiftserverlessCustomDomainAssociationExample.overrideLogicalId(
      "example"
    );
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `workgroupName` - (Required) Name of the workgroup.
* `customDomainName` - (Required) Custom domain to associate with the workgroup.
* `customDomainCertificateArn` - (Required) ARN of the certificate for the custom domain association.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `customDomainCertificateExpiryTime` - Expiration time for the certificate.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Redshift Serverless Custom Domain Association using the `workgroupName` and `customDomainName`, separated by the coma. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { RedshiftserverlessCustomDomainAssociation } from "./.gen/providers/aws/redshiftserverless-custom-domain-association";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    RedshiftserverlessCustomDomainAssociation.generateConfigForImport(
      this,
      "example",
      "example-workgroup,example.com"
    );
  }
}

```

Using `terraform import`, import Redshift Serverless Custom Domain Association using the `workgroupName` and `customDomainName`, separated by the coma. For example:

```console
% terraform import aws_redshiftserverless_custom_domain_association.example example-workgroup,example.com
```

<!-- cache-key: cdktf-0.20.8 input-27ad05d47eccb883d5fc8d70fbf7a2829718fe20ba96b49d8cd0ba02d484fe21 -->