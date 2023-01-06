variable "additional_inline_policy_json" {
  type        = map(string)
  description = "A map of policy name to json to add as inline policies to this role"
  default     = {}
}

variable "role_assumption_policy_json" {
  type        = string
  description = "JSON string for IAM role assumption rules"
}

variable "role_description" {
  type        = string
  description = "Description to add to the IAM Role"
  default     = ""
}

variable "role_name" {
  type        = string
  description = "The name to use for the IAM Role"
}
