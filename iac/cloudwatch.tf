resource "aws_cloudwatch_log_group" "this" {
  name = "/aws/api_gw/${aws_apigatewayv2_api.this.name}"

  retention_in_days = 30
}
