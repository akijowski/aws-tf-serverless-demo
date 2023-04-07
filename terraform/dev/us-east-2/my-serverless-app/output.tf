output "s3_bucket" {
  description = "The S3 bucket for this project"
  value       = module.app.s3_bucket
}

output "rest_api_id" {
  description = "The API Gateway ID"
  value       = module.app.rest_api_id
}

output "rest_api_arn" {
  description = "The API Gateway ARN"
  value       = module.app.rest_api_arn
}

output "rest_api_stage_name" {
  description = "The API Gateway Stage"
  value       = module.app.rest_api_stage_name
}

output "rest_api_deployment_id" {
  description = "The API Gateway Deployment ID"
  value       = module.app.rest_api_deployment_id
}

output "lambdas" {
  description = "The created Lambda Functions"
  value       = module.app.lambdas
}

output "code_deploy_app_name" {
  value = module.app.code_deploy_app_name
}

output "code_deploy_service_role_arn" {
  value = module.app.code_deploy_role_arn
}

output "code_deploy_groups" {
  value = module.app.code_deploy_groups
}

output "code_deploy_appspec_etags" {
  value = module.app.code_deploy_appspec_etags
}

output "code_deploy_cmds" {
  value = module.app.code_deploy_cmds
}
