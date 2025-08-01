# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "aws_backup_region_settings" "test" {
  resource_type_opt_in_preference = {
    "Aurora"                 = true
    "CloudFormation"         = true
    "DocumentDB"             = true
    "DSQL"                   = true
    "DynamoDB"               = true
    "EBS"                    = true
    "EC2"                    = true
    "EFS"                    = true
    "FSx"                    = true
    "Neptune"                = true
    "RDS"                    = true
    "Redshift"               = true
    "Redshift Serverless"    = true
    "S3"                     = true
    "SAP HANA on Amazon EC2" = true
    "Storage Gateway"        = true
    "Timestream"             = true
    "VirtualMachine"         = true
  }
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.100.0"
    }
  }
}

provider "aws" {}
