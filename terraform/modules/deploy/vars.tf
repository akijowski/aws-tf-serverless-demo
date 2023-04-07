variable "app_prefix" {
  type        = string
  description = "Prefix for the CodeDeploy application"
}

variable "s3_bucket" {
  type        = string
  description = "The S3 bucket used to store the latest AppSpec file.  This file can be used to start a CodeDeploy Deployment"
}

variable "groups" {
  type = map(object(
    {
      deploy_config_name = string
    }
  ))
  description = "Map of deployment group settings.  Each key will be used as the name for a group and the values are additional settings to apply"
}

variable "deploy_command_abs_path" {
  description = "The absolute path to write a script to run code deploy based on the generated appspec files"
  type        = string
}

variable "lambda_names" {
  type        = set(string)
  description = "Distinct set of Lambda names to look up.  The Lambda function must already exist prior to this module being executed"
}

variable "managed_alias" {
  type        = string
  default     = "Live"
  description = "The Lambda Alias for CodeDeploy to use.  If it does not exist on a Lambda it will be created"
}
