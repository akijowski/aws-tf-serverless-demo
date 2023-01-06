output "api_gateway_arn" {
  description = "API Gateway ARN"
  value = module.api_gateway.api_gateway_arn
}

output "api_gateway_id" {
  description = "API Gateway unique ID"
  value = module.api_gateway.api_gateway_id
}

output "api_gateway_deployment_id" {
  description = "API Gateway deployment ID"
  value = module.api_gateway.api_gateway_deployment_id
}

output "hello_lambda_function_arn" {
  description = "The Hello Lambda function ARN"
  value = module.hello_lambda_function.function_arn
}

output "hello_lambda_function_invoke_arn" {
  description = "The Hello Lambda function invocation ARN.  This is useful for API Gateway"
  value = module.hello_lambda_function.function_invoke_arn
}

output "hello_lambda_function_qualified_arn" {
  description = "The fully-qualified Lambda function ARN that includes the Lambda Alias or Version"
  value = module.hello_lambda_function.function_qualified_arn
}

output "hello_lambda_s3_object_key" {
  description = "The S3 Object key that stores the Hello Lambda function artifact"
  value = module.hello_lambda_artifact.artifact_object_key
}

output "hello_lambda_function_execution_role_arn" {
  description = "The Hello Lambda function execution IAM Role"
  value = module.hello_lambda_execution_role.arn
}
