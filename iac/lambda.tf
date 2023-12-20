resource "aws_lambda_function" "this" {
  for_each         = local.api_action
  architectures    = ["arm64"]
  function_name    = var.pr_id == "none" ? format("%s_%s", local.project_name, each.key) : format("%s_%s_%s", local.project_name, each.key, var.pr_id)
  filename         = format("./build/src/api/%s/%s/bootstrap.zip", split("_", each.key)[0], split("_", each.key)[1])
  handler          = "bootstrap"
  runtime          = "provided.al2"
  memory_size      = 128
  timeout          = 3
  role             = each.value.role_arn
  source_code_hash = filebase64sha256(format("./build/src/api/%s/%s/bootstrap.zip", split("_", each.key)[0], split("_", each.key)[1]))

  environment {
    variables = {
      TABLE_NAME                    = aws_dynamodb_table.this.name
      ACCESS_CONTROL_ALLOWED_HEADER = "http://localhost:3000" # temporary to allow local development
    }
  }
}
