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
      config_name     = string
      current_version = string
      target_version  = string
      alias           = string
    }
  ))
  description = "Map of deployment group settings.  Each key will be used as the name for a group and the values are additional settings to apply"
}

variable "deploy_command_abs_path" {
  description = "The absolute path to write a script to run code deploy based on the generated appspec files"
  type        = string
}
