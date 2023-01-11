module "iam_policies" {
  source = "../../../modules/iam_policies"
}

module "hello_lambda_execution_role" {
  source = "../../../modules/iam_role"

  role_name                   = "${var.hello_lambda_function_name}-execution-role"
  role_assumption_policy_json = module.iam_policies.lambda_assume_role_json

  additional_inline_policy_json = {
    "basic_execution"   = module.iam_policies.lambda_basic_execution_json
    "xray_write_access" = module.iam_policies.lambda_xray_write_json
  }
}

resource "random_pet" "lambda_bucket_name" {
  prefix = var.hello_lambda_function_name
  length = 3
}

module "lambda_bucket" {
  source = "../../../modules/s3_bucket"

  bucket_prefix = random_pet.lambda_bucket_name.id
}

module "hello_lambda_artifact" {
  source = "../../../modules/s3_lambda_artifact"

  abs_file_path  = abspath("../../../../tmp/hello.zip")
  s3_object_key  = "hello-lambda.zip"
  s3_bucket_name = module.lambda_bucket.s3_bucket_name
}

module "hello_lambda_function" {
  source = "../../../modules/lambda"

  function_alias_name         = var.lambda_function_alias
  function_execution_role_arn = module.hello_lambda_execution_role.arn
  function_name               = var.hello_lambda_function_name

  function_s3_bucket         = module.lambda_bucket.s3_bucket_name
  function_s3_key            = module.hello_lambda_artifact.artifact_object_key
  function_s3_object_version = module.hello_lambda_artifact.artifact_object_version_id
}

module "api_gateway" {
  source = "../../../modules/apigw"

  api_name                           = var.api_gateway_name
  open_api_abs_path                  = abspath("../../../../reference/openapi.yaml")
  execution_permissions_lambda_names = [var.hello_lambda_function_name]
  hello_lambda_invocation_arn        = module.hello_lambda_function.function_invoke_arn
}
