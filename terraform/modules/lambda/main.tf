resource "aws_lambda_function" "main" {
  function_name = var.function_name
  role          = var.function_execution_role_arn

  s3_bucket         = var.function_s3_bucket
  s3_key            = var.function_s3_key
  s3_object_version = var.function_s3_object_version

  publish = true

  timeout     = 5
  memory_size = 128
  handler     = "hello"
  runtime     = "go1.x"

  depends_on = [
    aws_cloudwatch_log_group.main
  ]
}

resource "aws_lambda_alias" "main" {
  function_name    = aws_lambda_function.main.arn
  function_version = aws_lambda_function.main.version
  name             = var.function_alias_name
}

resource "aws_cloudwatch_log_group" "main" {
  name              = "/aws/lambda/${var.function_name}"
  retention_in_days = 14
}
