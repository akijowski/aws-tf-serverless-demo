output "group_name" {
  description = "The CodeDeploy Group name"
  value       = aws_codedeploy_deployment_group.main.id
}

output "group_arn" {
  description = "The CodeDeploy Group ARN"
  value       = aws_codedeploy_deployment_group.main.arn
}
