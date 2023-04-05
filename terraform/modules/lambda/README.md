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
| [aws_cloudwatch_log_group.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group) | resource |
| [aws_iam_role.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_lambda_alias.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_alias) | resource |
| [aws_lambda_function.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function) | resource |
| [aws_s3_object.artifact](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_object) | resource |
| [aws_iam_policy_document.lambda_assume_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_abs_file_path"></a> [abs\_file\_path](#input\_abs\_file\_path) | The absolute file path to the zip archive | `string` | n/a | yes |
| <a name="input_additional_inline_policy_json"></a> [additional\_inline\_policy\_json](#input\_additional\_inline\_policy\_json) | A map of policy name to json to add as inline policies to this lambda | `map(string)` | `{}` | no |
| <a name="input_bucket_name"></a> [bucket\_name](#input\_bucket\_name) | The S3 bucket where the Lambda Zip file will be stored | `string` | n/a | yes |
| <a name="input_function_alias"></a> [function\_alias](#input\_function\_alias) | The name of the Alias that will point to the latest version | `string` | n/a | yes |
| <a name="input_function_handler"></a> [function\_handler](#input\_function\_handler) | The handler Lambda will invoke on startup | `string` | n/a | yes |
| <a name="input_function_name"></a> [function\_name](#input\_function\_name) | The name of the Lambda function | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_function_alias_arn"></a> [function\_alias\_arn](#output\_function\_alias\_arn) | The Lambda function Alias ARN |
| <a name="output_function_arn"></a> [function\_arn](#output\_function\_arn) | The Lambda function ARN |
| <a name="output_function_execution_role_arn"></a> [function\_execution\_role\_arn](#output\_function\_execution\_role\_arn) | The Lambda function execution IAM Role ARN |
| <a name="output_function_invoke_arn"></a> [function\_invoke\_arn](#output\_function\_invoke\_arn) | The Lambda function invocation ARN.  This is useful for API Gateway |
| <a name="output_function_log_group_arn"></a> [function\_log\_group\_arn](#output\_function\_log\_group\_arn) | The Lambda function log group ARN |
| <a name="output_function_qualified_arn"></a> [function\_qualified\_arn](#output\_function\_qualified\_arn) | The Lambda function fully qualified ARN which includes the version or alias |
| <a name="output_function_version"></a> [function\_version](#output\_function\_version) | The Lambda function latest version |
<!-- END_TF_DOCS -->