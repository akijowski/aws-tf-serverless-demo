output "s3_bucket_name" {
  description = "The created S3 Bucket name"
  value = aws_s3_bucket_versioning.main_versioning.bucket
}
