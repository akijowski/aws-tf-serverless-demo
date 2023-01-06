<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_s3_object.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_object) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_abs_file_path"></a> [abs\_file\_path](#input\_abs\_file\_path) | absolute path to file for upload | `string` | n/a | yes |
| <a name="input_s3_bucket_name"></a> [s3\_bucket\_name](#input\_s3\_bucket\_name) | The AWS S3 bucket to use for storage | `string` | n/a | yes |
| <a name="input_s3_object_key"></a> [s3\_object\_key](#input\_s3\_object\_key) | The AWS S3 object key to use for storage | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_artifact_object_key"></a> [artifact\_object\_key](#output\_artifact\_object\_key) | n/a |
| <a name="output_artifact_object_version_id"></a> [artifact\_object\_version\_id](#output\_artifact\_object\_version\_id) | n/a |
<!-- END_TF_DOCS -->