<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |
| <a name="provider_random"></a> [random](#provider\_random) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_codedeploy_deployment_group.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/codedeploy_deployment_group) | resource |
| [random_string.random_ten](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_codedeploy_app_name"></a> [codedeploy\_app\_name](#input\_codedeploy\_app\_name) | The name for the CodeDeploy Application that is attached to this Group | `string` | n/a | yes |
| <a name="input_codedeploy_group_name_prefix"></a> [codedeploy\_group\_name\_prefix](#input\_codedeploy\_group\_name\_prefix) | The prefix to use when generating the CodeDeploy Group name | `string` | n/a | yes |
| <a name="input_codedeploy_group_service_role_arn"></a> [codedeploy\_group\_service\_role\_arn](#input\_codedeploy\_group\_service\_role\_arn) | The service role ARN for this Group to use | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_group_arn"></a> [group\_arn](#output\_group\_arn) | The CodeDeploy Group ARN |
| <a name="output_group_name"></a> [group\_name](#output\_group\_name) | The CodeDeploy Group name |
<!-- END_TF_DOCS -->