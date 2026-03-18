# Terraform AWS Module - Scheduler Schedule Group

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This Terraform module creates an [AWS EventBridge Scheduler Schedule Group](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-schedule-groups.html). Schedule groups organize schedules for easier management and isolation.

## Pre-Commit Hooks

[.pre-commit-config.yaml](.pre-commit-config.yaml) defines pre-commit hooks for Terraform, Go, and common linting. The `commitlint` hook enforces conventional commit messages. The `detect-secrets-hook` prevents new secrets from being introduced. See [pre-commit documentation](https://pre-commit.com/) for installation. For `commitlint`, run:

```
pre-commit install --hook-type commit-msg
```

## Usage

See the [complete example](examples/complete/) for a full working configuration using the resource naming module.

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.14 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | ~> 5.14 |

## Resources

| Name | Type |
|------|------|
| [aws_scheduler_schedule_group.schedule_group](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/scheduler_schedule_group) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | Name of the schedule group. Must be unique within the account. Conflicts with name\_prefix. | `string` | `null` | no |
| <a name="input_name_prefix"></a> [name\_prefix](#input\_name\_prefix) | Creates a unique name beginning with the specified prefix. Conflicts with name. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Map of tags to assign to the schedule group. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | Create and delete timeout configurations. Defaults are 5 minutes each. | <pre>object({<br/>    create = optional(string, "5m")<br/>    delete = optional(string, "5m")<br/>  })</pre> | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the schedule group. |
| <a name="output_creation_date"></a> [creation\_date](#output\_creation\_date) | The time the schedule group was created. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the schedule group (same as the name). |
| <a name="output_last_modification_date"></a> [last\_modification\_date](#output\_last\_modification\_date) | The time the schedule group was last modified. |
| <a name="output_name"></a> [name](#output\_name) | The name of the schedule group. |
| <a name="output_state"></a> [state](#output\_state) | The state of the schedule group (ACTIVE or DELETING). |
| <a name="output_tags_all"></a> [tags\_all](#output\_tags\_all) | Map of all tags assigned to the schedule group, including those inherited from the provider. |

## Testing

Run `make configure` to set up the repository, then `make check` to run lint, validate, plan, and Terratest. Ensure AWS credentials are configured (e.g., `AWS_PROFILE` or `AWS_ACCESS_KEY_ID`).

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

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_scheduler_schedule_group.schedule_group](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/scheduler_schedule_group) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | Name of the schedule group. Must be unique within the account. Conflicts with name\_prefix. | `string` | `null` | no |
| <a name="input_name_prefix"></a> [name\_prefix](#input\_name\_prefix) | Creates a unique name beginning with the specified prefix. Conflicts with name. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Map of tags to assign to the schedule group. | `map(string)` | `{}` | no |
| <a name="input_timeouts"></a> [timeouts](#input\_timeouts) | Create and delete timeout configurations. Defaults are 5 minutes each.<br/>Example: { create = "10m", delete = "10m" } | <pre>object({<br/>    create = optional(string, "5m")<br/>    delete = optional(string, "5m")<br/>  })</pre> | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | The ID of the schedule group (same as the name). |
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the schedule group. |
| <a name="output_name"></a> [name](#output\_name) | The name of the schedule group. |
| <a name="output_state"></a> [state](#output\_state) | The state of the schedule group (ACTIVE or DELETING). |
| <a name="output_creation_date"></a> [creation\_date](#output\_creation\_date) | The time the schedule group was created. |
| <a name="output_last_modification_date"></a> [last\_modification\_date](#output\_last\_modification\_date) | The time the schedule group was last modified. |
| <a name="output_tags_all"></a> [tags\_all](#output\_tags\_all) | Map of all tags assigned to the schedule group, including those inherited from the provider. |
<!-- END_TF_DOCS -->
