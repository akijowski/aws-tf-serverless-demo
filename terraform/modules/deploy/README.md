<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |
| <a name="provider_local"></a> [local](#provider\_local) | n/a |
| <a name="provider_random"></a> [random](#provider\_random) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_codedeploy_app.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/codedeploy_app) | resource |
| [aws_codedeploy_deployment_group.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/codedeploy_deployment_group) | resource |
| [aws_iam_role.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy_attachment.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_s3_object.appspec](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_object) | resource |
| [local_file.deploy_script](https://registry.terraform.io/providers/hashicorp/local/latest/docs/resources/file) | resource |
| [random_string.random_ten](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string) | resource |
| [aws_iam_policy_document.assume_role](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_app_prefix"></a> [app\_prefix](#input\_app\_prefix) | Prefix for the CodeDeploy application | `string` | n/a | yes |
| <a name="input_deploy_command_abs_path"></a> [deploy\_command\_abs\_path](#input\_deploy\_command\_abs\_path) | The absolute path to write a script to run code deploy based on the generated appspec files | `string` | n/a | yes |
| <a name="input_groups"></a> [groups](#input\_groups) | Map of deployment group settings.  Each key will be used as the name for a group and the values are additional settings to apply | <pre>map(object(<br>    {<br>      config_name     = string<br>      current_version = string<br>      target_version  = string<br>      alias           = string<br>    }<br>  ))</pre> | n/a | yes |
| <a name="input_s3_bucket"></a> [s3\_bucket](#input\_s3\_bucket) | The S3 bucket used to store the latest AppSpec file.  This file can be used to start a CodeDeploy Deployment | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_app_name"></a> [app\_name](#output\_app\_name) | The created CodeDeploy App name |
| <a name="output_appspec_object_etag"></a> [appspec\_object\_etag](#output\_appspec\_object\_etag) | The S3 Object ETAG for the generated appspec file |
| <a name="output_appspec_object_version"></a> [appspec\_object\_version](#output\_appspec\_object\_version) | The S3 Object version for the generated appspec file |
| <a name="output_deploy_cmds"></a> [deploy\_cmds](#output\_deploy\_cmds) | n/a |
| <a name="output_groups"></a> [groups](#output\_groups) | The created CodeDeploy Groups |
| <a name="output_role_arn"></a> [role\_arn](#output\_role\_arn) | The IAM Role created for the groups |
<!-- END_TF_DOCS -->