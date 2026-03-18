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

output "id" {
  description = "The ID of the schedule group (same as the name)."
  value       = aws_scheduler_schedule_group.schedule_group.id
}

output "arn" {
  description = "The ARN of the schedule group."
  value       = aws_scheduler_schedule_group.schedule_group.arn
}

output "name" {
  description = "The name of the schedule group."
  value       = aws_scheduler_schedule_group.schedule_group.name
}

output "state" {
  description = "The state of the schedule group (ACTIVE or DELETING)."
  value       = aws_scheduler_schedule_group.schedule_group.state
}

output "creation_date" {
  description = "The time the schedule group was created."
  value       = aws_scheduler_schedule_group.schedule_group.creation_date
}

output "last_modification_date" {
  description = "The time the schedule group was last modified."
  value       = aws_scheduler_schedule_group.schedule_group.last_modification_date
}

output "tags_all" {
  description = "Map of all tags assigned to the schedule group, including those inherited from the provider."
  value       = aws_scheduler_schedule_group.schedule_group.tags_all
}
