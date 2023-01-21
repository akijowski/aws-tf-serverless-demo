output "lambda_assume_role_json" {
  value = data.aws_iam_policy_document.lambda_assume_role.json
}

output "lambda_basic_execution_json" {
  value = data.aws_iam_policy.lambda_basic_execution.policy
}

output "lambda_xray_write_json" {
  value = data.aws_iam_policy.lambda_xray_write.policy
}

output "codedeploy_assume_role_json" {
  value = data.aws_iam_policy_document.codedeploy_assume_role.json
}
