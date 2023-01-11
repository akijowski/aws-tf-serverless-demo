variable "api_name" {
  type        = string
  description = "API Gateway name"
}
variable "open_api_abs_path" {
  type        = string
  description = "Absolute file path to the OpenAPI spec for the API body"
}

variable "execution_permissions_lambda_names" {
  type        = list(string)
  description = "List of Lambda function names to grant execution permission for this API"
  default     = []
}
