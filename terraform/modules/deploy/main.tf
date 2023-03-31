locals {
  app_name = "${var.app_prefix}-${random_string.random_ten.id}"
}

resource "random_string" "random_ten" {
  length  = 10
  special = false
  lower   = false
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    sid    = ""
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["codedeploy.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "main" {
  count = length(var.groups) > 0 ? 1 : 0

  name               = "${local.app_name}-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy_attachment" "main" {
  count = length(var.groups) > 0 ? 1 : 0

  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole"
  role       = aws_iam_role.main[count.index].name
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
  deployment_config_name = each.value["config_name"]
  service_role_arn       = one(aws_iam_role.main[*].arn)

  auto_rollback_configuration {
    enabled = true
    events  = ["DEPLOYMENT_FAILURE", "DEPLOYMENT_STOP_ON_ALARM"]
  }

  deployment_style {
    deployment_type = "BLUE_GREEN"
    deployment_option = "WITH_TRAFFIC_CONTROL"
  }
}

resource "aws_s3_object" "appspec" {
  for_each = try(var.groups, {})

  bucket = var.s3_bucket
  key    = "${var.app_prefix}/${each.key}/appspec.json"

  content_type = "application/json"
  content      = templatefile("${path.module}/files/appspec.tftpl", { name = each.key, group = each.value })
}

resource "local_file" "deploy_script" {
  count = length(var.groups) > 0 ? 1 : 0

  filename        = "${var.deploy_command_abs_path}/deploy.sh"
  file_permission = "0740"
  content         = templatefile("${path.module}/files/awsdeploy.tftpl", {
    names          = keys(var.groups)
    app_name       = one(aws_codedeploy_app.main[*].name)
    bucket         = var.s3_bucket
    deploy_groups  = aws_codedeploy_deployment_group.main
    app_spec_files = aws_s3_object.appspec
  })
}
