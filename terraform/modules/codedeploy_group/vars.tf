variable "codedeploy_app_name" {
  type        = string
  description = "The name for the CodeDeploy Application that is attached to this Group"
}

variable "codedeploy_group_name_prefix" {
  type        = string
  description = "The prefix to use when generating the CodeDeploy Group name"
}

variable "codedeploy_group_service_role_arn" {
  type        = string
  description = "The service role ARN for this Group to use"
}
