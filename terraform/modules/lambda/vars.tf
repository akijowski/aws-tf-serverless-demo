variable "function_alias_name" {
  type        = string
  description = "The lambda function alias for the latest version"
}

variable "function_execution_role_arn" {
  type        = string
  description = "The lambda function execution IAM role ARN"
}

variable "function_name" {
  type        = string
  description = "The lambda function name"
}

variable "function_s3_bucket" {
  type        = string
  description = "The S3 Bucket containing the deployment package"
}

variable "function_s3_key" {
  type        = string
  description = "The S3 Object Key containing the deployment package"
}

variable "function_s3_object_version" {
  type        = string
  description = "The S3 Object Version containing the deployment package"
}
