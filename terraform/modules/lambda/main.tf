data "aws_iam_policy_document" "lambda_assume_role" {
  version = "2012-10-17"
  statement {
    sid     = ""
    effect  = "Allow"
    actions = ["sts:AssumeRole"]
    principals {
      identifiers = ["lambda.amazonaws.com"]
      type        = "Service"
    }
  }
}

# IAM Execution Role
resource "aws_iam_role" "main" {
  name               = "${var.function_name}-execution-role"
  description        = "Lambda Execution role for ${var.function_name}"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json

  dynamic "inline_policy" {
    for_each = var.additional_inline_policy_json
    content {
      name   = inline_policy.key
      policy = inline_policy.value
    }
  }
}

# S3 Artifact
resource "aws_s3_object" "artifact" {
  bucket = var.bucket_name
  key    = "${var.function_name}/${var.function_name}.zip"

  source      = var.abs_file_path
  source_hash = filemd5(var.abs_file_path)
}

# Lambda Function
resource "aws_lambda_function" "main" {
  function_name = var.function_name
  role          = aws_iam_role.main.arn

  s3_bucket         = var.bucket_name
  s3_key            = aws_s3_object.artifact.key
  s3_object_version = aws_s3_object.artifact.version_id

  publish = true

  timeout     = 5
  memory_size = 128
  handler     = var.function_handler
  runtime     = "go1.x"

  depends_on = [
    aws_cloudwatch_log_group.main
  ]
}

# Lambda Alias
resource "aws_lambda_alias" "main" {
  function_name    = aws_lambda_function.main.arn
  function_version = aws_lambda_function.main.version
  name             = var.function_alias
}

# Cloudwatch Logs
resource "aws_cloudwatch_log_group" "main" {
  name              = "/aws/lambda/${var.function_name}"
  retention_in_days = 14
}
