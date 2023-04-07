output "arn" {
  description = "The Lambda function ARN"
  value       = aws_lambda_function.main.arn
}

output "invoke_arn" {
  description = "The Lambda function invocation ARN.  This is useful for API Gateway"
  value       = aws_lambda_alias.main.invoke_arn
}

output "qualified_arn" {
  description = "The Lambda function fully qualified ARN which includes the version or alias"
  value       = aws_lambda_function.main.qualified_arn
}

output "alias_arn" {
  description = "The Lambda function Alias ARN"
  value       = aws_lambda_alias.main.arn
}

output "version" {
  description = "The Lambda function latest version"
  value       = aws_lambda_function.main.version
}

output "log_group_arn" {
  description = "The Lambda function log group ARN"
  value       = aws_cloudwatch_log_group.main.arn
}

output "execution_role_arn" {
  description = "The Lambda function execution IAM Role ARN"
  value       = aws_iam_role.main.arn
}
