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
| [aws_iam_policy.lambda_basic_execution](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy) | data source |
| [aws_iam_policy.lambda_xray_write](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy) | data source |
| [aws_iam_policy_document.lambda_assume_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |

## Inputs

No inputs.

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_lambda_assume_role_json"></a> [lambda\_assume\_role\_json](#output\_lambda\_assume\_role\_json) | n/a |
| <a name="output_lambda_basic_execution_json"></a> [lambda\_basic\_execution\_json](#output\_lambda\_basic\_execution\_json) | n/a |
| <a name="output_lambda_xray_write_json"></a> [lambda\_xray\_write\_json](#output\_lambda\_xray\_write\_json) | n/a |
<!-- END_TF_DOCS -->