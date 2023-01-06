output "artifact_object_key" {
  value = aws_s3_object.main.id
}

output "artifact_object_version_id" {
  value = aws_s3_object.main.version_id
}
