// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

module "resource_names" {
  source  = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version = "~> 2.0"

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

  name        = coalesce(var.name, module.resource_names["schedule_group"].standard)
  name_prefix = var.name != null ? null : var.name_prefix
  tags        = var.tags
  timeouts    = var.timeouts
}

# Default event bus ARN for schedule targets
locals {
  event_bus_arn = "arn:aws:events:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:event-bus/default"
}

# IAM role for EventBridge Scheduler to put events
resource "aws_iam_role" "scheduler" {
  name = module.resource_names["scheduler_role"].standard

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "scheduler.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role_policy" "scheduler_events" {
  name = "SchedulerEventsPolicy"
  role = aws_iam_role.scheduler.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = "events:PutEvents"
        Resource = local.event_bus_arn
      }
    ]
  })
}

resource "aws_scheduler_schedule" "schedule_1" {
  name       = module.resource_names["schedule_1"].standard
  group_name = module.schedule_group.name

  flexible_time_window {
    mode = "OFF"
  }

  schedule_expression = "rate(1 day)"

  target {
    arn      = local.event_bus_arn
    role_arn = aws_iam_role.scheduler.arn
  }
}

resource "aws_scheduler_schedule" "schedule_2" {
  name       = module.resource_names["schedule_2"].standard
  group_name = module.schedule_group.name

  flexible_time_window {
    mode = "OFF"
  }

  schedule_expression = "rate(1 hour)"

  target {
    arn      = local.event_bus_arn
    role_arn = aws_iam_role.scheduler.arn
  }
}
