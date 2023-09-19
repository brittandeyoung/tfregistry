locals {
  ddb_table_name = var.pr_id == "none" ? local.project_name : format("%s_%s", local.project_name, var.pr_id)
}

resource "aws_dynamodb_table" "this" {
  name           = local.ddb_table_name
  billing_mode   = "PROVISIONED"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "pk"
  range_key      = "sk"

  attribute {
    name = "pk"
    type = "S"
  }

  attribute {
    name = "sk"
    type = "S"
  }
}
