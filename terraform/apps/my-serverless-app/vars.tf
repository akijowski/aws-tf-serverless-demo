variable "app_name" {
  type        = string
  description = "The name for the project application"
  default     = "my-serverless-app"
}

variable "code_deploy_script_path" {
  type = string
  description = "Path to the directory where a generated deploy script will be created.  It will be converted to an absolute path"
  default = "../../../../tmp"
}
