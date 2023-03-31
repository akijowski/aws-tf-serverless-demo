# REST API
resource "aws_api_gateway_rest_api" "main" {
  name = var.api_name
  body = templatefile(var.open_api_abs_path, var.open_api_variables)

  put_rest_api_mode = "merge"
}

# Deployment
resource "aws_api_gateway_deployment" "main" {
  rest_api_id = aws_api_gateway_rest_api.main.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.main.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

# Stage
resource "aws_api_gateway_stage" "main" {
  deployment_id = aws_api_gateway_deployment.main.id
  rest_api_id   = aws_api_gateway_rest_api.main.id
  stage_name    = "live"

  xray_tracing_enabled = true

  variables = var.stage_variables

  # Access logs are set on an Account basis: https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions
  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.main_access.arn
    format = jsonencode({
      "requestId"         = "$context.requestId"
      "extendedRequestId" = "$context.extendedRequestId"
      "ip"                = "$context.identity.sourceIp"
      "caller"            = "$context.identity.caller"
      "requestTime"       = "$context.requestTime"
      "httpMethod"        = "$context.httpMethod"
      "resourcePath"      = "$context.resourcePath"
      "status"            = "$context.status"
    })
  }

  depends_on = [aws_cloudwatch_log_group.main_access]
}

# Method Settings
resource "aws_api_gateway_method_settings" "main" {
  method_path = "*/*"
  rest_api_id = aws_api_gateway_rest_api.main.id
  stage_name  = aws_api_gateway_stage.main.stage_name
  settings {
    metrics_enabled = true
    # logging_level   = "INFO"
  }
}

# Lambda execution permissions
resource "aws_lambda_permission" "main" {
  for_each = try(var.lambda_execution_object, {})

  statement_id_prefix = "api_gateway_invoke"
  action              = "lambda:InvokeFunction"
  function_name       = each.key
  qualifier           = each.value.qualifier
  principal           = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.main.execution_arn}/*/*/*"
}

# If account level API logging is enabled, you can add this
resource "aws_cloudwatch_log_group" "main_execution" {
  name              = "API-Gateway-Execution-Logs_${aws_api_gateway_rest_api.main.id}/${aws_api_gateway_stage.main.stage_name}"
  retention_in_days = 7
}

resource "aws_cloudwatch_log_group" "main_access" {
  name              = "/aws/apigw/${aws_api_gateway_rest_api.main.name}/live/access"
  retention_in_days = 7
}
