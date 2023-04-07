/**
* # Serverless Demo App
*
* This module is the `main` application and can be used as a root module for deployment.
*/
locals {
  root_dir_rel_path = "${path.module}/../../.."
  lambdas = {
    hello-world = {
      function_handler = "hello"
      abs_file_path    = abspath("${local.root_dir_rel_path}/tmp/hello.zip")
    }
  }
}

# IAM Policies
module "iam_policies" {
  source = "../iam_policies"
}

# Project S3 bucket
module "project_bucket" {
  source        = "../s3_bucket"
  bucket_prefix = "${var.app_name}-"
}

# Project API Gateway
module "api" {
  source = "../apigw"

  api_name                = var.app_name
  open_api_abs_path       = abspath("${local.root_dir_rel_path}/reference/openapi.yaml")
  lambda_execution_object = { for name, _ in local.lambdas : name => { qualifier = "Live" } }

  stage_variables = {}

  open_api_variables = {
    helloLambdaInvocationArn = module.lambda_functions["hello-world"].invoke_arn
  }

  depends_on = [
    module.lambda_functions,
    module.code_deploy
  ]
}

# Project Lambdas
module "lambda_functions" {
  for_each = try(local.lambdas, {})
  source   = "../lambda"

  abs_file_path    = each.value.abs_file_path
  bucket_name      = module.project_bucket.s3_bucket_name
  function_alias   = "Latest"
  function_handler = each.value.function_handler
  function_name    = each.key

  additional_inline_policy_json = {
    "basic_execution"   = module.iam_policies.lambda_basic_execution_json
    "xray_write_access" = module.iam_policies.lambda_xray_write_json
  }
}

# CodeDeploy
module "code_deploy" {
  source = "../deploy"

  app_prefix = var.app_name

  groups = { for name in keys(local.lambdas) : name => {
    deploy_config_name = "CodeDeployDefault.LambdaAllAtOnce"
  } }

  # The lambda names must be consistent for Terraform to accurately plan.  Map keys in a for_each meta-argument cannot be dynamic.
  lambda_names = toset(keys(local.lambdas))

  s3_bucket               = module.project_bucket.s3_bucket_name
  deploy_command_abs_path = abspath("${local.root_dir_rel_path}/${var.code_deploy_script_path}")

  depends_on = [
    module.lambda_functions
  ]
}
