variable "api_name" {
  type        = string
  description = "API Gateway name"
}
variable "open_api_abs_path" {
  type        = string
  description = "Absolute file path to the OpenAPI spec for the API body"
}

variable "hello_lambda_invocation_arn" {
  type        = string
  description = "The invocation ARN for the hello world lambda function"
}

variable "execution_permissions_lambda_names" {
  type        = list(string)
  description = "List of Lambda function names to grant execution permission for this API"
  default     = []
}

variable "execution_permissions_lambda_qualifier" {
  type        = string
  description = "The Alias or Version number to apply when granting execution permission for this API"
  default     = "Live"
}
