resource "aws_apigatewayv2_api" "this" {
  name          = "registry"
  protocol_type = "HTTP"
}

module "api" {
  for_each      = local.api
  source        = "./api"
  api           = aws_apigatewayv2_api.this
  stage_name    = "api"
  resource_name = each.key
  table_name    = aws_dynamodb_table.this.name
  route_keys    = local.api[each.key]
  dynamodb_arn  = aws_dynamodb_table.this.arn
  cw_group_arn  = aws_cloudwatch_log_group.api.arn
}

resource "aws_cloudwatch_log_group" "api" {
  name = "/aws/api_gw/${aws_apigatewayv2_api.this.name}"

  retention_in_days = 30
}
