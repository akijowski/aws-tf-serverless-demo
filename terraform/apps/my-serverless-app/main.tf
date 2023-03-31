locals {
  lambdas = {
    hello-world = {
      function_handler = "hello"
      abs_file_path    = abspath("../../../../tmp/hello.zip")
    }
  }
}

# IAM Policies
module "iam_policies" {
  source = "../../modules/iam_policies"
}

# Project S3 bucket
module "project_bucket" {
  source        = "../../modules/s3_bucket"
  bucket_prefix = "${var.app_name}-"
}

# Project API Gateway
module "api" {
  source = "../../modules/apigw"

  api_name                = var.app_name
  open_api_abs_path       = abspath("../../../../reference/openapi.yaml")
  lambda_execution_object = { for name, _ in local.lambdas : name => { qualifier = "Live" } }

  stage_variables = {}

  open_api_variables = {
    helloLambdaInvocationArn = module.lambda_functions["hello-world"].function_invoke_arn
  }

  depends_on = [
    module.lambda_functions
  ]
}

# Project Lambdas
module "lambda_functions" {
  for_each = try(local.lambdas, {})
  source   = "../../modules/lambda"

  abs_file_path    = each.value.abs_file_path
  bucket_name      = module.project_bucket.s3_bucket_name
  function_alias   = "Live"
  function_handler = each.value.function_handler
  function_name    = each.key

  additional_inline_policy_json = {
    "basic_execution"   = module.iam_policies.lambda_basic_execution_json
    "xray_write_access" = module.iam_policies.lambda_xray_write_json
  }
}

# CodeDeploy
module "code_deploy" {
  source = "../../modules/deploy"

  app_prefix = var.app_name
  groups = { for name, lambda in module.lambda_functions : name => {
    config_name = "CodeDeployDefault.LambdaAllAtOnce"
    alias = "Live"
    current_version = lambda.function_version > 1 ? lambda.function_version - 1 : lambda.function_version
    target_version = lambda.function_version
  } }
  s3_bucket = module.project_bucket.s3_bucket_name
  deploy_command_abs_path = abspath(var.code_deploy_script_path)

  depends_on = [
    module.lambda_functions
  ]
}