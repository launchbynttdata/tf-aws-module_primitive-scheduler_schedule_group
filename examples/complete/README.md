# Complete Example - AWS Scheduler Schedule Group

This example creates an EventBridge Scheduler schedule group using the primitive module with the resource naming module for consistent naming.

## Usage

```hcl
data "aws_region" "current" {}

module "resource_names" {
  source   = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version  = "~> 2.0"

  for_each = var.resource_names_map

  logical_product_family  = var.logical_product_family
  logical_product_service = var.logical_product_service
  class_env               = var.class_env
  instance_env            = var.instance_env
  instance_resource       = var.instance_resource
  cloud_resource_type     = each.value.name
  maximum_length          = each.value.max_length

  region = join("", split("-", data.aws_region.current.name))
}

module "schedule_group" {
  source = "../.."

  name        = module.resource_names["schedule_group"].standard
  name_prefix = var.name_prefix
  tags        = var.tags
  timeouts    = var.timeouts
}
```

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_logical_product_family"></a> [logical_product_family](#input_logical_product_family) | Logical product family for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical_product_service](#input_logical_product_service) | Logical product service for resource naming. | `string` | n/a | yes |
| <a name="input_class_env"></a> [class_env](#input_class_env) | Environment class for resource naming (e.g., dev, prod). | `string` | n/a | yes |
| <a name="input_instance_env"></a> [instance_env](#input_instance_env) | Instance environment number (0-999) for resource naming. | `number` | n/a | yes |
| <a name="input_instance_resource"></a> [instance_resource](#input_instance_resource) | Instance resource number (0-100) for resource naming. | `number` | n/a | yes |
| <a name="input_resource_names_map"></a> [resource_names_map](#input_resource_names_map) | Map of key to resource_name configuration for the resource naming module. | `map(object({ name = string, max_length = optional(number, 64) }))` | n/a | yes |
| <a name="input_name_prefix"></a> [name_prefix](#input_name_prefix) | Creates a unique name beginning with the specified prefix. Conflicts with name. Set to null when using resource naming module for name. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input_tags) | Map of tags to assign to the schedule group. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input_timeouts) | Create and delete timeout configurations. | `object({ create = optional(string, "5m"), delete = optional(string, "5m") })` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_region"></a> [region](#output_region) | The AWS region where resources are created. |
| <a name="output_id"></a> [id](#output_id) | The ID of the schedule group. |
| <a name="output_arn"></a> [arn](#output_arn) | The ARN of the schedule group. |
| <a name="output_name"></a> [name](#output_name) | The name of the schedule group. |
| <a name="output_state"></a> [state](#output_state) | The state of the schedule group. |
| <a name="output_creation_date"></a> [creation_date](#output_creation_date) | The time the schedule group was created. |
| <a name="output_last_modification_date"></a> [last_modification_date](#output_last_modification_date) | The time the schedule group was last modified. |
| <a name="output_tags_all"></a> [tags_all](#output_tags_all) | Map of all tags assigned to the schedule group. |
| <a name="output_schedule_names"></a> [schedule_names](#output_schedule_names) | Names of the schedules in the schedule group. |

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.14 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.100.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.0 |
| <a name="module_schedule_group"></a> [schedule\_group](#module\_schedule\_group) | ../.. | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_iam_role.scheduler](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy.scheduler_events](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy) | resource |
| [aws_scheduler_schedule.schedule_1](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/scheduler_schedule) | resource |
| [aws_scheduler_schedule.schedule_2](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/scheduler_schedule) | resource |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | Logical product family for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | Logical product service for resource naming. | `string` | n/a | yes |
| <a name="input_class_env"></a> [class\_env](#input\_class\_env) | Environment class for resource naming (e.g., dev, prod). | `string` | n/a | yes |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Instance environment number (0-999) for resource naming. | `number` | n/a | yes |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Instance resource number (0-100) for resource naming. | `number` | n/a | yes |
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | Map of key to resource\_name configuration for the resource naming module. | <pre>map(object({<br/>    name       = string<br/>    max_length = optional(number, 64)<br/>  }))</pre> | n/a | yes |
| <a name="input_name_prefix"></a> [name\_prefix](#input\_name\_prefix) | Creates a unique name beginning with the specified prefix. Conflicts with name. Set to null when using resource naming module for name. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Map of tags to assign to the schedule group. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | Create and delete timeout configurations. | <pre>object({<br/>    create = optional(string, "5m")<br/>    delete = optional(string, "5m")<br/>  })</pre> | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_region"></a> [region](#output\_region) | The AWS region where resources are created. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the schedule group. |
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the schedule group. |
| <a name="output_name"></a> [name](#output\_name) | The name of the schedule group. |
| <a name="output_state"></a> [state](#output\_state) | The state of the schedule group. |
| <a name="output_creation_date"></a> [creation\_date](#output\_creation\_date) | The time the schedule group was created. |
| <a name="output_last_modification_date"></a> [last\_modification\_date](#output\_last\_modification\_date) | The time the schedule group was last modified. |
| <a name="output_tags_all"></a> [tags\_all](#output\_tags\_all) | Map of all tags assigned to the schedule group. |
| <a name="output_schedule_names"></a> [schedule\_names](#output\_schedule\_names) | Names of the schedules in the schedule group. |
<!-- END_TF_DOCS -->
