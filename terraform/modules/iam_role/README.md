<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_iam_role.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_additional_inline_policy_json"></a> [additional\_inline\_policy\_json](#input\_additional\_inline\_policy\_json) | A map of policy name to json to add as inline policies to this role | `map(string)` | `{}` | no |
| <a name="input_role_assumption_policy_json"></a> [role\_assumption\_policy\_json](#input\_role\_assumption\_policy\_json) | JSON string for IAM role assumption rules | `string` | n/a | yes |
| <a name="input_role_description"></a> [role\_description](#input\_role\_description) | Description to add to the IAM Role | `string` | `""` | no |
| <a name="input_role_name"></a> [role\_name](#input\_role\_name) | The name to use for the IAM Role | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The IAM Role ARN |
| <a name="output_name"></a> [name](#output\_name) | The IAM Role name |
| <a name="output_unique_id"></a> [unique\_id](#output\_unique\_id) | The IAM Role stable unique ID |
<!-- END_TF_DOCS -->