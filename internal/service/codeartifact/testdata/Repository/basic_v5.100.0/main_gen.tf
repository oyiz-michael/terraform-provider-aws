# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "aws_codeartifact_repository" "test" {
  repository = var.rName
  domain     = aws_codeartifact_domain.test.domain
}

resource "aws_codeartifact_domain" "test" {
  domain         = var.rName
  encryption_key = aws_kms_key.test.arn
}

resource "aws_kms_key" "test" {
  description             = var.rName
  deletion_window_in_days = 7
  enable_key_rotation     = true
}

variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
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
