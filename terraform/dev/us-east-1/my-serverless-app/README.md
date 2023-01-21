<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 4.0 |
| <a name="requirement_random"></a> [random](#requirement\_random) | ~> 3.1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 4.48.0 |
| <a name="provider_random"></a> [random](#provider\_random) | 3.1.3 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_api_gateway"></a> [api\_gateway](#module\_api\_gateway) | ../../../modules/apigw | n/a |
| <a name="module_codedeploy_app"></a> [codedeploy\_app](#module\_codedeploy\_app) | ../../../modules/codedeploy_app | n/a |
| <a name="module_codedeploy_group_hello_lambda"></a> [codedeploy\_group\_hello\_lambda](#module\_codedeploy\_group\_hello\_lambda) | ../../../modules/codedeploy_group | n/a |
| <a name="module_codedeploy_role"></a> [codedeploy\_role](#module\_codedeploy\_role) | ../../../modules/iam_role | n/a |
| <a name="module_hello_lambda_artifact"></a> [hello\_lambda\_artifact](#module\_hello\_lambda\_artifact) | ../../../modules/s3_lambda_artifact | n/a |
| <a name="module_hello_lambda_execution_role"></a> [hello\_lambda\_execution\_role](#module\_hello\_lambda\_execution\_role) | ../../../modules/iam_role | n/a |
| <a name="module_hello_lambda_function"></a> [hello\_lambda\_function](#module\_hello\_lambda\_function) | ../../../modules/lambda | n/a |
| <a name="module_iam_policies"></a> [iam\_policies](#module\_iam\_policies) | ../../../modules/iam_policies | n/a |
| <a name="module_lambda_bucket"></a> [lambda\_bucket](#module\_lambda\_bucket) | ../../../modules/s3_bucket | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_iam_role_policy_attachment.main_codedeploy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [random_pet.lambda_bucket_name](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/pet) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_api_gateway_name"></a> [api\_gateway\_name](#input\_api\_gateway\_name) | The name to use for the API Gateway REST API | `string` | n/a | yes |
| <a name="input_hello_lambda_function_name"></a> [hello\_lambda\_function\_name](#input\_hello\_lambda\_function\_name) | The name of the Hello Lambda function | `string` | n/a | yes |
| <a name="input_lambda_function_alias"></a> [lambda\_function\_alias](#input\_lambda\_function\_alias) | The name to use for the latest Lambda function version alias | `string` | `"Live"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_api_gateway_arn"></a> [api\_gateway\_arn](#output\_api\_gateway\_arn) | API Gateway ARN |
| <a name="output_api_gateway_deployment_id"></a> [api\_gateway\_deployment\_id](#output\_api\_gateway\_deployment\_id) | API Gateway deployment ID |
| <a name="output_api_gateway_id"></a> [api\_gateway\_id](#output\_api\_gateway\_id) | API Gateway unique ID |
| <a name="output_hello_lambda_function_arn"></a> [hello\_lambda\_function\_arn](#output\_hello\_lambda\_function\_arn) | The Hello Lambda function ARN |
| <a name="output_hello_lambda_function_execution_role_arn"></a> [hello\_lambda\_function\_execution\_role\_arn](#output\_hello\_lambda\_function\_execution\_role\_arn) | The Hello Lambda function execution IAM Role |
| <a name="output_hello_lambda_function_invoke_arn"></a> [hello\_lambda\_function\_invoke\_arn](#output\_hello\_lambda\_function\_invoke\_arn) | The Hello Lambda function invocation ARN.  This is useful for API Gateway |
| <a name="output_hello_lambda_function_qualified_arn"></a> [hello\_lambda\_function\_qualified\_arn](#output\_hello\_lambda\_function\_qualified\_arn) | The fully-qualified Lambda function ARN that includes the Lambda Alias or Version |
| <a name="output_hello_lambda_s3_object_key"></a> [hello\_lambda\_s3\_object\_key](#output\_hello\_lambda\_s3\_object\_key) | The S3 Object key that stores the Hello Lambda function artifact |
<!-- END_TF_DOCS -->