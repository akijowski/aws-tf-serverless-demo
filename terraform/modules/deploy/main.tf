/**
* # CodeDeploy
*
* This module is a fairly complex module which configures a CodeDeploy application for AWS Lambda.
* For each Lambda function passed in `var.groups` a CodeDeploy group will be created.
* For each Lambda name passed in `var.lambda_names` the Lambda will be looked up and the previous version and current version will be used to generate an AppSpec file for deployment.
* A helper script will also be generated to use with the AWS CLI (v2).
*/
locals {
  app_name = "${var.app_prefix}-${random_string.random_ten.id}"
  lambdas = { for name, lambda in data.aws_lambda_function.main : name => {
    unqualified_arn  = lambda.arn
    latest_version   = lambda.version
    previous_version = tostring(max(parseint(lambda.version, 10) - 1, 1))
  } }
}

resource "random_string" "random_ten" {
  length  = 10
  special = false
  lower   = false
}

resource "aws_codedeploy_app" "main" {
  count = length(var.groups) > 0 ? 1 : 0

  compute_platform = "Lambda"
  name             = local.app_name
}

resource "aws_codedeploy_deployment_group" "main" {
  for_each = try(var.groups, {})

  app_name               = one(aws_codedeploy_app.main[*].name)
  deployment_group_name  = "${local.app_name}-${each.key}"
  deployment_config_name = each.value["deploy_config_name"]
  service_role_arn       = one(aws_iam_role.main[*].arn)

  auto_rollback_configuration {
    enabled = true
    events  = ["DEPLOYMENT_FAILURE", "DEPLOYMENT_STOP_ON_ALARM"]
  }

  deployment_style {
    deployment_type   = "BLUE_GREEN"
    deployment_option = "WITH_TRAFFIC_CONTROL"
  }
}

data "aws_lambda_function" "main" {
  for_each = try(var.lambda_names, [])

  function_name = each.value
}

resource "aws_lambda_alias" "main" {
  for_each = try(var.lambda_names, [])

  name             = var.managed_alias
  description      = "Managed by CodeDeploy and Terraform"
  function_name    = local.lambdas[each.value].unqualified_arn
  function_version = local.lambdas[each.value].previous_version

  # Only want to create the alias, not update
  lifecycle {
    ignore_changes = [
      function_name,
      function_version
    ]
  }
}

# AppSpec Resources
resource "aws_s3_object" "appspec" {
  for_each = try(var.lambda_names, [])

  bucket = var.s3_bucket
  key    = "${var.app_prefix}/${each.value}/appspec.json"

  content_type = "application/json"
  content = templatefile("${path.module}/files/appspec.tftpl", {
    name            = each.value
    alias           = var.managed_alias
    current_version = local.lambdas[each.value].previous_version
    target_version  = local.lambdas[each.value].latest_version
    }
  )

  tags = {
    # The service role CodeDeployRoleForLambda can look up objects with this tag
    "UseWithCodeDeploy" = "true"
  }
}

resource "local_file" "deploy_script" {
  count = length(var.groups) > 0 ? 1 : 0

  filename        = "${var.deploy_command_abs_path}/deploy.sh"
  file_permission = "0740"
  content = templatefile("${path.module}/files/awsdeploy.tftpl", {
    names          = keys(var.groups)
    app_name       = one(aws_codedeploy_app.main[*].name)
    bucket         = var.s3_bucket
    deploy_groups  = aws_codedeploy_deployment_group.main
    app_spec_files = aws_s3_object.appspec
  })
}

resource "local_file" "appspec" {
  for_each = try(var.lambda_names, [])

  filename = "${var.deploy_command_abs_path}/${each.value}.appspec.json"
  content = templatefile("${path.module}/files/appspec.tftpl", {
    name            = each.value
    alias           = var.managed_alias
    current_version = local.lambdas[each.value]["previous_version"]
    target_version  = local.lambdas[each.value]["latest_version"]
    }
  )
}
