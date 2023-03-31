output "api_gateway_id" {
  description = "API Gateway unique ID"
  value       = aws_api_gateway_rest_api.main.id
}

output "api_gateway_arn" {
  description = "API Gateway ARN"
  value       = aws_api_gateway_rest_api.main.arn
}

output "api_gateway_stage_name" {
  description = "API Gateway stage name"
  value       = aws_api_gateway_stage.main.stage_name
}

output "api_gateway_deployment_id" {
  description = "API Gateway deployment unique ID"
  value       = aws_api_gateway_deployment.main.id
}

output "api_access_log_arn" {
  description = "The Cloudwatch Log Group ARN for API Access Logs"
  value = aws_cloudwatch_log_group.main_access.arn
}

output "api_execution_log_arn" {
  description = "The Cloudwatch Log Group ARN for API Execution Logs"
  value = aws_cloudwatch_log_group.main_execution.arn
}
