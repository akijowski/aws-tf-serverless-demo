output "name" {
  description = "The CodeDeploy Application name"
  value       = aws_codedeploy_app.main.name
}

output "arn" {
  description = "The CodeDeploy Application ARN"
  value       = aws_codedeploy_app.main.arn
}
