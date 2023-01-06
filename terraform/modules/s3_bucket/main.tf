resource "aws_s3_bucket" "main" {
  bucket_prefix = var.bucket_prefix

  force_destroy = true
}

resource "aws_s3_bucket_versioning" "main_versioning" {
  bucket = aws_s3_bucket.main.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_acl" "main_acl" {
  bucket = aws_s3_bucket.main.id
  acl = "private"

}
