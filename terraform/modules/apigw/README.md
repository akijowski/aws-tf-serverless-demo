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
| [aws_api_gateway_deployment.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_deployment) | resource |
| [aws_api_gateway_method_settings.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method_settings) | resource |
| [aws_api_gateway_rest_api.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_rest_api) | resource |
| [aws_api_gateway_stage.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_stage) | resource |
| [aws_cloudwatch_log_group.main_access](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group) | resource |
| [aws_cloudwatch_log_group.main_execution](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group) | resource |
| [aws_lambda_permission.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_api_name"></a> [api\_name](#input\_api\_name) | API Gateway name | `string` | n/a | yes |
| <a name="input_lambda_execution_object"></a> [lambda\_execution\_object](#input\_lambda\_execution\_object) | Allows API Gateway to invoke Lambda.  Map of objects where the key is the Lambda function name and the object is configuration applied to a lambda permission resource | <pre>map(object({<br>    qualifier = string<br>  }))</pre> | n/a | yes |
| <a name="input_open_api_abs_path"></a> [open\_api\_abs\_path](#input\_open\_api\_abs\_path) | Absolute file path to the OpenAPI spec for the API body | `string` | n/a | yes |
| <a name="input_open_api_variables"></a> [open\_api\_variables](#input\_open\_api\_variables) | Variable map to use when templating the OpenaPI spec | `map(string)` | n/a | yes |
| <a name="input_stage_variables"></a> [stage\_variables](#input\_stage\_variables) | Map of additional variables to add to an API Gateway stage | `map(string)` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_api_access_log_arn"></a> [api\_access\_log\_arn](#output\_api\_access\_log\_arn) | The Cloudwatch Log Group ARN for API Access Logs |
| <a name="output_api_execution_log_arn"></a> [api\_execution\_log\_arn](#output\_api\_execution\_log\_arn) | The Cloudwatch Log Group ARN for API Execution Logs |
| <a name="output_api_gateway_arn"></a> [api\_gateway\_arn](#output\_api\_gateway\_arn) | API Gateway ARN |
| <a name="output_api_gateway_deployment_id"></a> [api\_gateway\_deployment\_id](#output\_api\_gateway\_deployment\_id) | API Gateway deployment unique ID |
| <a name="output_api_gateway_id"></a> [api\_gateway\_id](#output\_api\_gateway\_id) | API Gateway unique ID |
| <a name="output_api_gateway_stage_name"></a> [api\_gateway\_stage\_name](#output\_api\_gateway\_stage\_name) | API Gateway stage name |
<!-- END_TF_DOCS -->