resource "aws_s3_object" "main" {
  bucket = var.s3_bucket_name
  key    = var.s3_object_key

  source = var.abs_file_path
  source_hash = filemd5(var.abs_file_path)
}
