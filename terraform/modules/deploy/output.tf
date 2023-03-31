output "app_name" {
  description = "The created CodeDeploy App name"
  value       = aws_codedeploy_app.main[0].name
}

output "role_arn" {
  description = "The IAM Role created for the groups"
  value       = aws_iam_role.main[0].arn
}

output "groups" {
  description = "The created CodeDeploy Groups"
  value = { for name, group in aws_codedeploy_deployment_group.main : name => {
    arn                    = group.arn
    group_id               = group.deployment_group_id
    deployment_config_name = group.deployment_config_name
    tags_all               = group.tags_all
  } }
}

output "appspec_object_version" {
  description = "The S3 Object version for the generated appspec file"
  value       = { for name, obj in aws_s3_object.appspec : name => obj.version_id }
}

output "appspec_object_etag" {
  description = "The S3 Object ETAG for the generated appspec file"
  value       = { for name, obj in aws_s3_object.appspec : name => obj.etag }
}

output "deploy_cmds" {
  value = one(local_file.deploy_script[*].filename)
}
