resource "aws_dynamodb_table" "this" {
  name           = "TerraformRegistry"
  billing_mode   = "PROVISIONED"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "resourceType"
  range_key      = "sortKey"

  attribute {
    name = "resourceType"
    type = "S"
  }

  attribute {
    name = "sortKey"
    type = "S"
  }
}
