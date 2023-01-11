resource "aws_api_gateway_rest_api" "main" {
  name = var.api_name
  body = file(var.open_api_abs_path)

  put_rest_api_mode = "merge"
}

resource "aws_api_gateway_deployment" "main" {
  rest_api_id = aws_api_gateway_rest_api.main.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.main.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "main" {
  deployment_id = aws_api_gateway_deployment.main.id
  rest_api_id   = aws_api_gateway_rest_api.main.id
  stage_name    = "live"

  xray_tracing_enabled = true

  # Access logs are set on an Account basis: https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions
  # access_log_settings {
  #   destination_arn = aws_cloudwatch_log_group.main_access.arn
  #   format          = jsonencode({
  #     "requestId"         = "$context.requestId"
  #     "extendedRequestId" = "$context.extendedRequestId"
  #     "ip"                = "$context.identity.sourceIp"
  #     "caller"            = "$context.identity.caller"
  #     "requestTime"       = "$context.requestTime"
  #     "httpMethod"        = "$context.httpMethod"
  #     "resourcePath"      = "$context.resourcePath"
  #     "status"            = "$context.status"
  #   })
  # }

  # depends_on = [aws_cloudwatch_log_group.main_access]
}

resource "aws_api_gateway_method_settings" "main" {
  method_path = "*/*"
  rest_api_id = aws_api_gateway_rest_api.main.id
  stage_name  = aws_api_gateway_stage.main.stage_name
  settings {
    metrics_enabled = true
    # logging_level   = "INFO"
  }
}

resource "aws_lambda_permission" "hello_lambda_api_permission" {
  count               = length(var.execution_permissions_lambda_names)
  statement_id_prefix = "api_gateway_invoke"
  action              = "lambda:InvokeFunction"
  function_name       = var.execution_permissions_lambda_names[count.index]
  principal           = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.main.execution_arn}/*/*/*"
}

# If account level API logging is enabled, you can add this
# resource "aws_cloudwatch_log_group" "main_execution" {
#   name              = "API-Gateway-Execution-Logs_${aws_api_gateway_rest_api.main.id}/${aws_api_gateway_stage.main.stage_name}"
#   retention_in_days = 7
# }

# resource "aws_cloudwatch_log_group" "main_access" {
#   name              = "/aws/apigw/${aws_api_gateway_rest_api.main.name}-access"
#   retention_in_days = 7
# }
