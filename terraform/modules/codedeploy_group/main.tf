resource "random_string" "random_ten" {
  length  = 10
  special = false
  lower   = false
}

resource "aws_codedeploy_deployment_group" "main" {
  app_name               = var.codedeploy_app_name
  deployment_group_name  = "${var.codedeploy_group_name_prefix}-${random_string.random_ten.id}"
  service_role_arn       = var.codedeploy_group_service_role_arn
  deployment_config_name = "CodeDeployDefault.LambdaAllAtOnce"

  auto_rollback_configuration {
    enabled = true
    events  = ["DEPLOYMENT_FAILURE", "DEPLOYMENT_STOP_ON_ALARM"]
  }
}
