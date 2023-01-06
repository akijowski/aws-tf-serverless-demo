variable "api_gateway_name" {
  type        = string
  description = "The name to use for the API Gateway REST API"
}

variable "lambda_function_alias" {
  type        = string
  description = "The name to use for the latest Lambda function version alias"
  default     = "Live"
}

variable "hello_lambda_function_name" {
  type        = string
  description = "The name of the Hello Lambda function"
}
