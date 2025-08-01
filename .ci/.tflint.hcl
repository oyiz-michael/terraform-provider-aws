plugin "aws" {
  enabled = true
  version = "0.39.0"
  source  = "github.com/terraform-linters/tflint-ruleset-aws"
}

# Terraform Rules
# https://github.com/terraform-linters/tflint-ruleset-terraform/blob/main/docs/rules/README.md

rule "terraform_comment_syntax" {
  enabled = true
}

rule "terraform_required_providers" {
  enabled = false
}

rule "terraform_required_version" {
  enabled = false
}

# AWS Rules
# https://github.com/terraform-linters/tflint-ruleset-aws/blob/master/docs/rules/README.md

rule "aws_acm_certificate_lifecycle" {
  enabled = false
}

rule "aws_accessanalyzer_analyzer_invalid_type" {
  enabled = false
}

# Avoids errant findings related to directory paths in generated configuration files
rule "aws_iam_saml_provider_invalid_saml_metadata_document" {
  enabled = false
}
