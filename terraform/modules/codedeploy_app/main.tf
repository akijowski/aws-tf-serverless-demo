resource "random_string" "random_ten" {
  length  = 10
  special = false
  lower   = false
}

resource "aws_codedeploy_app" "main" {
  compute_platform = "Lambda"
  name             = "${var.codedeploy_app_name_prefix}-${random_string.random_ten.id}"
}
