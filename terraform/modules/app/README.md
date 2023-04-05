<!-- BEGIN_TF_DOCS -->
# Serverless Demo App

This module is the `main` application and can be used as a root module for deployment.

## Requirements

No requirements.

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_api"></a> [api](#module\_api) | ../apigw | n/a |
| <a name="module_iam_policies"></a> [iam\_policies](#module\_iam\_policies) | ../iam_policies | n/a |
| <a name="module_lambda_functions"></a> [lambda\_functions](#module\_lambda\_functions) | ../lambda | n/a |
| <a name="module_project_bucket"></a> [project\_bucket](#module\_project\_bucket) | ../s3_bucket | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_app_name"></a> [app\_name](#input\_app\_name) | The name for the project application | `string` | `"my-serverless-app"` | no |
| <a name="input_code_deploy_script_path"></a> [code\_deploy\_script\_path](#input\_code\_deploy\_script\_path) | Path to the directory where a generated deploy script will be created.  It will be converted to an absolute path | `string` | `"../../../../tmp"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_lambdas"></a> [lambdas](#output\_lambdas) | Map of the created Lambdas |
| <a name="output_rest_api_access_log_arn"></a> [rest\_api\_access\_log\_arn](#output\_rest\_api\_access\_log\_arn) | The API Gateway access log ARN |
| <a name="output_rest_api_arn"></a> [rest\_api\_arn](#output\_rest\_api\_arn) | The API Gateway ARN |
| <a name="output_rest_api_deployment_id"></a> [rest\_api\_deployment\_id](#output\_rest\_api\_deployment\_id) | The API Gateway Deployment ID |
| <a name="output_rest_api_execution_log_arn"></a> [rest\_api\_execution\_log\_arn](#output\_rest\_api\_execution\_log\_arn) | The API Gateway execution log ARN |
| <a name="output_rest_api_id"></a> [rest\_api\_id](#output\_rest\_api\_id) | The API Gateway ID |
| <a name="output_rest_api_stage_name"></a> [rest\_api\_stage\_name](#output\_rest\_api\_stage\_name) | The API Gateway Stage |
| <a name="output_s3_bucket"></a> [s3\_bucket](#output\_s3\_bucket) | The S3 bucket for this project |
<!-- END_TF_DOCS -->