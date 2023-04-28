variable "api" {}
variable "stage_name" {
  type = string
}

variable "route_keys" {
}

variable "dynamodb_arn" {
  type = string
}

variable "cw_group_arn" {
  type = string
}

variable "table_name" {
  type = string
}

variable "resource_name" {
  type = string
}
