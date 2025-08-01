---
subcategory: "ACM PCA (Certificate Manager Private Certificate Authority)"
layout: "aws"
page_title: "AWS: aws_acmpca_certificate_authority"
description: |-
  Get information on a AWS Certificate Manager Private Certificate Authority
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_acmpca_certificate_authority

Get information on a AWS Certificate Manager Private Certificate Authority (ACM PCA Certificate Authority).

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_acmpca_certificate_authority import DataAwsAcmpcaCertificateAuthority
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsAcmpcaCertificateAuthority(self, "example",
            arn="arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"
        )
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `arn` - (Required) ARN of the certificate authority.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - ARN of the certificate authority.
* `certificate` - Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
* `certificate_chain` - Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
* `certificate_signing_request` - The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
* `usage_mode` - Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly.
* `not_after` - Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
* `not_before` - Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
* `revocation_configuration` - Nested attribute containing revocation configuration.
    * `revocation_configuration.0.crl_configuration` - Nested attribute containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority.
        * `revocation_configuration.0.crl_configuration.0.custom_cname` - Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point.
        * `revocation_configuration.0.crl_configuration.0.enabled` - Boolean value that specifies whether certificate revocation lists (CRLs) are enabled.
        * `revocation_configuration.0.crl_configuration.0.expiration_in_days` - Number of days until a certificate expires.
        * `revocation_configuration.0.crl_configuration.0.s3_bucket_name` - Name of the S3 bucket that contains the CRL.
        * `revocation_configuration.0.crl_configuration.0.s3_object_acl` - Whether the CRL is publicly readable or privately held in the CRL Amazon S3 bucket.
        * `revocation_configuration.0.ocsp_configuration.0.enabled` - Boolean value that specifies whether a custom OCSP responder is enabled.
        * `revocation_configuration.0.ocsp_configuration.0.ocsp_custom_cname` - A CNAME specifying a customized OCSP domain.
* `serial` - Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
* `status` - Status of the certificate authority.
* `tags` - Key-value map of user-defined tags that are attached to the certificate authority.
* `type` - Type of the certificate authority.

<!-- cache-key: cdktf-0.20.8 input-587c4772b3b9d6f6a944d0a4befecbe537495d065124acc24954b27217a83fb1 -->