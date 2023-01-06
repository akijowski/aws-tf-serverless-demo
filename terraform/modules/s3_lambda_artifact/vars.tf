variable "abs_file_path" {
  type = string
  description = "absolute path to file for upload"
}

variable "s3_bucket_name" {
  type = string
  description = "The AWS S3 bucket to use for storage"
}

variable "s3_object_key" {
  type = string
  description = "The AWS S3 object key to use for storage"
}
