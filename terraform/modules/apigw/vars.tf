variable "api_name" {
  type        = string
  description = "API Gateway name"
}
variable "open_api_abs_path" {
  type        = string
  description = "Absolute file path to the OpenAPI spec for the API body"
}
variable "open_api_variables" {
  type        = map(string)
  description = "Variable map to use when templating the OpenaPI spec"
}
variable "lambda_execution_object" {
  description = "Allows API Gateway to invoke Lambda.  Map of objects where the key is the Lambda function name and the object is configuration applied to a lambda permission resource"
  type = map(object({
    qualifier = string
  }))
}
variable "stage_variables" {
  description = "Map of additional variables to add to an API Gateway stage"
  type = map(string)
}
