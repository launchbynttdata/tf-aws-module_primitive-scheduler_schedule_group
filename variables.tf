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

# -----------------------------------------------------------------------------
# Schedule group naming
# -----------------------------------------------------------------------------

variable "name" {
  description = "Name of the schedule group. Must be unique within the account. Conflicts with name_prefix."
  type        = string
  default     = null

  validation {
    condition     = var.name == null || var.name_prefix == null
    error_message = "Only one of 'name' or 'name_prefix' can be set, not both."
  }

  validation {
    condition     = var.name == null ? true : (length(var.name) >= 1 && length(var.name) <= 64)
    error_message = "Name must be between 1 and 64 characters."
  }
}

variable "name_prefix" {
  description = "Creates a unique name beginning with the specified prefix. Conflicts with name."
  type        = string
  default     = null

  validation {
    condition     = var.name_prefix == null ? true : (length(var.name_prefix) >= 1 && length(var.name_prefix) <= 64)
    error_message = "Name prefix must be between 1 and 64 characters."
  }
}

# -----------------------------------------------------------------------------
# Configuration
# -----------------------------------------------------------------------------

variable "tags" {
  description = "Map of tags to assign to the schedule group."
  type        = map(string)
  default     = {}
}

# -----------------------------------------------------------------------------
# Timeouts
# -----------------------------------------------------------------------------

variable "timeouts" {
  description = <<-EOT
    Create and delete timeout configurations. Defaults are 5 minutes each.
    Example: { create = "10m", delete = "10m" }
  EOT
  type = object({
    create = optional(string, "5m")
    delete = optional(string, "5m")
  })
  default = null
}
