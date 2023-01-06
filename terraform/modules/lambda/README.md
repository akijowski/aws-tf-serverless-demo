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
| [aws_lambda_alias.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_alias) | resource |
| [aws_lambda_function.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_function_alias_name"></a> [function\_alias\_name](#input\_function\_alias\_name) | The lambda function alias for the latest version | `string` | n/a | yes |
| <a name="input_function_execution_role_arn"></a> [function\_execution\_role\_arn](#input\_function\_execution\_role\_arn) | The lambda function execution IAM role ARN | `string` | n/a | yes |
| <a name="input_function_name"></a> [function\_name](#input\_function\_name) | The lambda function name | `string` | n/a | yes |
| <a name="input_function_s3_bucket"></a> [function\_s3\_bucket](#input\_function\_s3\_bucket) | The S3 Bucket containing the deployment package | `string` | n/a | yes |
| <a name="input_function_s3_key"></a> [function\_s3\_key](#input\_function\_s3\_key) | The S3 Object Key containing the deployment package | `string` | n/a | yes |
| <a name="input_function_s3_object_version"></a> [function\_s3\_object\_version](#input\_function\_s3\_object\_version) | The S3 Object Version containing the deployment package | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_function_alias_arn"></a> [function\_alias\_arn](#output\_function\_alias\_arn) | The Lambda function Alias ARN |
| <a name="output_function_arn"></a> [function\_arn](#output\_function\_arn) | The Lambda function ARN |
| <a name="output_function_invoke_arn"></a> [function\_invoke\_arn](#output\_function\_invoke\_arn) | The Lambda function invocation ARN.  This is useful for API Gateway |
| <a name="output_function_log_group_arn"></a> [function\_log\_group\_arn](#output\_function\_log\_group\_arn) | The Lambda function log group ARN |
| <a name="output_function_qualified_arn"></a> [function\_qualified\_arn](#output\_function\_qualified\_arn) | The Lambda function fully qualified ARN which includes the version or alias |
| <a name="output_function_version"></a> [function\_version](#output\_function\_version) | The Lambda function latest version |
<!-- END_TF_DOCS -->