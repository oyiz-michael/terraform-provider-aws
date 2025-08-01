---
subcategory: "EKS (Elastic Kubernetes)"
layout: "aws"
page_title: "AWS: aws_eks_clusters"
description: |-
  Retrieve EKS Clusters list
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_eks_clusters

Retrieve EKS Clusters list

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformIterator, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_eks_cluster import DataAwsEksCluster
from imports.aws.data_aws_eks_clusters import DataAwsEksClusters
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = DataAwsEksClusters(self, "example")
        # In most cases loops should be handled in the programming language context and
        #     not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
        #     you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
        #     you need to keep this like it is.
        example_for_each_iterator = TerraformIterator.from_list(
            Token.as_any(Fn.toset(example.names)))
        data_aws_eks_cluster_example = DataAwsEksCluster(self, "example_1",
            name=Token.as_string(example_for_each_iterator.value),
            for_each=example_for_each_iterator
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        data_aws_eks_cluster_example.override_logical_id("example")
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - AWS Region.
* `names` - Set of EKS clusters names

<!-- cache-key: cdktf-0.20.8 input-bbdf203f6acf79163e602b55c79f082bdc1f1127a7bd441a503ec40e16e7dc44 -->