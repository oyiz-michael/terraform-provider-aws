---
subcategory: "Network Firewall"
layout: "aws"
page_title: "AWS: aws_networkfirewall_resource_policy"
description: |-
  Provides an AWS Network Firewall Resource Policy resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_networkfirewall_resource_policy

Provides an AWS Network Firewall Resource Policy Resource for a rule group or firewall policy.

## Example Usage

### For a Firewall Policy resource

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkfirewall_resource_policy import NetworkfirewallResourcePolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NetworkfirewallResourcePolicy(self, "example",
            policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": ["network-firewall:ListFirewallPolicies", "network-firewall:CreateFirewall", "network-firewall:UpdateFirewall", "network-firewall:AssociateFirewallPolicy"
                        ],
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::123456789012:root"
                        },
                        "Resource": aws_networkfirewall_firewall_policy_example.arn
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            resource_arn=Token.as_string(aws_networkfirewall_firewall_policy_example.arn)
        )
```

### For a Rule Group resource

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkfirewall_resource_policy import NetworkfirewallResourcePolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NetworkfirewallResourcePolicy(self, "example",
            policy=Token.as_string(
                Fn.jsonencode({
                    "Statement": [{
                        "Action": ["network-firewall:ListRuleGroups", "network-firewall:CreateFirewallPolicy", "network-firewall:UpdateFirewallPolicy"
                        ],
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::123456789012:root"
                        },
                        "Resource": aws_networkfirewall_rule_group_example.arn
                    }
                    ],
                    "Version": "2012-10-17"
                })),
            resource_arn=Token.as_string(aws_networkfirewall_rule_group_example.arn)
        )
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `policy` - (Required) JSON formatted policy document that controls access to the Network Firewall resource. The policy must be provided **without whitespaces**.  We recommend using [jsonencode](https://www.terraform.io/docs/configuration/functions/jsonencode.html) for formatting as seen in the examples above. For more details, including available policy statement Actions, see the [Policy](https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_PutResourcePolicy.html#API_PutResourcePolicy_RequestSyntax) parameter in the AWS API documentation.
* `resource_arn` - (Required, Forces new resource) The Amazon Resource Name (ARN) of the rule group or firewall policy.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The Amazon Resource Name (ARN) of the rule group or firewall policy associated with the resource policy.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Network Firewall Resource Policies using the `resource_arn`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.networkfirewall_resource_policy import NetworkfirewallResourcePolicy
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        NetworkfirewallResourcePolicy.generate_config_for_import(self, "example", "aws_networkfirewall_rule_group.example arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example")
```

Using `terraform import`, import Network Firewall Resource Policies using the `resource_arn`. For example:

```console
% terraform import aws_networkfirewall_resource_policy.example aws_networkfirewall_rule_group.example arn:aws:network-firewall:us-west-1:123456789012:stateful-rulegroup/example
```

<!-- cache-key: cdktf-0.20.8 input-ed7682a2a66f98e5f0ecf2441d396e01b605388bf1767d88f1c329b0c65722e2 -->