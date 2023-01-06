output "arn" {
  description = "The IAM Role ARN"
  value = aws_iam_role.main.arn
}

output "name" {
  description = "The IAM Role name"
  value = aws_iam_role.main.name
}

output "unique_id" {
  description = "The IAM Role stable unique ID"
  value = aws_iam_role.main.unique_id
}
