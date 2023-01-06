resource "aws_iam_role" "main" {
  name               = var.role_name
  description        = var.role_description
  assume_role_policy = var.role_assumption_policy_json

  dynamic "inline_policy" {
    for_each = var.additional_inline_policy_json
    content {
      name   = inline_policy.key
      policy = inline_policy.value
    }
  }
}
