variable "bucket_name" {
  type        = string
  description = "The S3 bucket where the Lambda Zip file will be stored"
}

variable "function_name" {
  type        = string
  description = "The name of the Lambda function"
}

variable "function_handler" {
  type        = string
  description = "The handler Lambda will invoke on startup"
}

variable "function_alias" {
  type        = string
  description = "The name of the Alias that will point to the latest version"
}

variable "abs_file_path" {
  type        = string
  description = "The absolute file path to the zip archive"
}

variable "additional_inline_policy_json" {
  type        = map(string)
  description = "A map of policy name to json to add as inline policies to this lambda"
  default     = {}
}
