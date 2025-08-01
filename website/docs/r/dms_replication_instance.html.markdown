---
subcategory: "DMS (Database Migration)"
layout: "aws"
page_title: "AWS: aws_dms_replication_instance"
description: |-
  Provides a DMS (Data Migration Service) replication instance resource.
---

# Resource: aws_dms_replication_instance

Provides a DMS (Data Migration Service) replication instance resource. DMS replication instances can be created, updated, deleted, and imported.

## Example Usage

Create required roles and then create a DMS instance, setting the depends_on to the required role policy attachments.

```terraform
# Database Migration Service requires the below IAM Roles to be created before
# replication instances can be created. See the DMS Documentation for
# additional information: https://docs.aws.amazon.com/dms/latest/userguide/security-iam.html#CHAP_Security.APIRole
#  * dms-vpc-role
#  * dms-cloudwatch-logs-role
#  * dms-access-for-endpoint

data "aws_iam_policy_document" "dms_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      identifiers = ["dms.amazonaws.com"]
      type        = "Service"
    }
  }
}

resource "aws_iam_role" "dms-access-for-endpoint" {
  assume_role_policy = data.aws_iam_policy_document.dms_assume_role.json
  name               = "dms-access-for-endpoint"
}

resource "aws_iam_role_policy_attachment" "dms-access-for-endpoint-AmazonDMSRedshiftS3Role" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role"
  role       = aws_iam_role.dms-access-for-endpoint.name
}

resource "aws_iam_role" "dms-cloudwatch-logs-role" {
  assume_role_policy = data.aws_iam_policy_document.dms_assume_role.json
  name               = "dms-cloudwatch-logs-role"
}

resource "aws_iam_role_policy_attachment" "dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole"
  role       = aws_iam_role.dms-cloudwatch-logs-role.name
}

resource "aws_iam_role" "dms-vpc-role" {
  assume_role_policy = data.aws_iam_policy_document.dms_assume_role.json
  name               = "dms-vpc-role"
}

resource "aws_iam_role_policy_attachment" "dms-vpc-role-AmazonDMSVPCManagementRole" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole"
  role       = aws_iam_role.dms-vpc-role.name
}

# Create a new replication instance
resource "aws_dms_replication_instance" "test" {
  allocated_storage            = 20
  apply_immediately            = true
  auto_minor_version_upgrade   = true
  availability_zone            = "us-west-2c"
  engine_version               = "3.1.4"
  kms_key_arn                  = "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012"
  multi_az                     = false
  preferred_maintenance_window = "sun:10:30-sun:14:30"
  publicly_accessible          = true
  replication_instance_class   = "dms.t3.micro"
  replication_instance_id      = "test-dms-replication-instance-tf"
  replication_subnet_group_id  = aws_dms_replication_subnet_group.test-dms-replication-subnet-group-tf.id

  tags = {
    Name = "test"
  }

  vpc_security_group_ids = [
    "sg-12345678",
  ]

  depends_on = [
    aws_iam_role_policy_attachment.dms-access-for-endpoint-AmazonDMSRedshiftS3Role,
    aws_iam_role_policy_attachment.dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole,
    aws_iam_role_policy_attachment.dms-vpc-role-AmazonDMSVPCManagementRole
  ]
}
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `allocated_storage` - (Optional, Default: 50, Min: 5, Max: 6144) The amount of storage (in gigabytes) to be initially allocated for the replication instance.
* `allow_major_version_upgrade` - (Optional, Default: false) Indicates that major version upgrades are allowed.
* `apply_immediately` - (Optional, Default: false) Indicates whether the changes should be applied immediately or during the next maintenance window. Only used when updating an existing resource.
* `auto_minor_version_upgrade` - (Optional, Default: false) Indicates that minor engine upgrades will be applied automatically to the replication instance during the maintenance window.
* `availability_zone` - (Optional) The EC2 Availability Zone that the replication instance will be created in.
* `dns_name_servers` - (Optional) A list of custom DNS name servers supported for the replication instance to access your on-premise source or target database. This list overrides the default name servers supported by the replication instance. You can specify a comma-separated list of internet addresses for up to four on-premise DNS name servers.
* `engine_version` - (Optional) The engine version number of the replication instance.
* `kerberos_authentication_settings` - (Optional) Configuration block for settings required for Kerberos authentication. See below.
* `kms_key_arn` - (Optional) The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
* `multi_az` - (Optional) Specifies if the replication instance is a multi-az deployment. You cannot set the `availability_zone` parameter if the `multi_az` parameter is set to `true`.
* `network_type` - (Optional) The type of IP address protocol used by a replication instance. Valid values: `IPV4`, `DUAL`.
* `preferred_maintenance_window` - (Optional) The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
* `publicly_accessible` - (Optional, Default: false) Specifies the accessibility options for the replication instance. A value of true represents an instance with a public IP address. A value of false represents an instance with a private IP address.
* `replication_instance_class` - (Required) The compute and memory capacity of the replication instance as specified by the replication instance class. See [AWS DMS User Guide](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_ReplicationInstance.Types.html) for available instance sizes and advice on which one to choose.
* `replication_instance_id` - (Required) The replication instance identifier. This parameter is stored as a lowercase string.
* `replication_subnet_group_id` - (Optional) A subnet group to associate with the replication instance.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `vpc_security_group_ids` - (Optional) A list of VPC security group IDs to be used with the replication instance. The VPC security groups must work with the VPC containing the replication instance.

## kerberos_authentication_settings

-> Additional information can be found in the [Using Kerberos Authentication with AWS Database Migration Service documentation](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.Kerberos.html).

* `key_cache_secret_iam_arn` - (Required) ARN of the IAM role that grants AWS DMS access to the secret containing key cache file for the Kerberos authentication.
* `key_cache_secret_id` - (Required) Secret ID that stores the key cache file required for Kerberos authentication.
* `krb5_file_contents` - (Required) Contents of krb5 configuration file required for Kerberos authentication.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `replication_instance_arn` - The Amazon Resource Name (ARN) of the replication instance.
* `replication_instance_private_ips` -  A list of the private IP addresses of the replication instance.
* `replication_instance_public_ips` - A list of the public IP addresses of the replication instance.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `40m`)
- `update` - (Default `30m`)
- `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import replication instances using the `replication_instance_id`. For example:

```terraform
import {
  to = aws_dms_replication_instance.test
  id = "test-dms-replication-instance-tf"
}
```

Using `terraform import`, import replication instances using the `replication_instance_id`. For example:

```console
% terraform import aws_dms_replication_instance.test test-dms-replication-instance-tf
```
