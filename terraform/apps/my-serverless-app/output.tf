output "s3_bucket" {
  description = "The S3 bucket for this project"
  value       = module.project_bucket.s3_bucket_name
}

output "rest_api_id" {
  description = "The API Gateway ID"
  value       = module.api.api_gateway_id
}

output "rest_api_arn" {
  description = "The API Gateway ARN"
  value       = module.api.api_gateway_arn
}

output "rest_api_stage_name" {
  description = "The API Gateway Stage"
  value       = module.api.api_gateway_stage_name
}

output "rest_api_deployment_id" {
  description = "The API Gateway Deployment ID"
  value       = module.api.api_gateway_deployment_id
}

output "rest_api_access_log_arn" {
  description = "The API Gateway access log ARN"
  value       = module.api.api_access_log_arn
}

output "rest_api_execution_log_arn" {
  description = "The API Gateway execution log ARN"
  value       = module.api.api_execution_log_arn
}

output "lambdas" {
  description = "Map of the created Lambdas"
  value       = module.lambda_functions
}

# output "code_deploy_role_arn" {
#   description = "Code Deploy service role ARN"
#   value       = module.code_deploy.role_arn
# }

# output "code_deploy_app_name" {
#   description = "Code Deploy App name"
#   value       = module.code_deploy.app_name
# }

# output "code_deploy_groups" {
#   description = "Code Deploy Groups"
#   value       = module.code_deploy.groups
# }

# output "code_deploy_appspec_etags" {
#   description = "The ETAGs for the created AppSpec files"
#   value       = module.code_deploy.appspec_object_etag
# }

# output "code_deploy_cmds" {
#   value = module.code_deploy.deploy_cmds
# }
