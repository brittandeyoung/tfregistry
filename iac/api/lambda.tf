resource "aws_lambda_function" "this" {
  for_each      = var.route_keys
  function_name = format("terraform-registry-%s-%s-%s", var.resource_name, var.stage_name, each.key)

  filename = format("build/src/%s/%s/main.zip", var.resource_name, each.key)
  handler  = "main"

  runtime          = "go1.x"
  memory_size      = 1024
  timeout          = 30
  source_code_hash = filebase64sha256(format("build/src/%s/%s/main.zip", var.resource_name, each.key))
  role             = aws_iam_role.this[each.key].arn
  environment {
    variables = {
      table_name = var.table_name
    }
  }
}

resource "aws_cloudwatch_log_group" "lambda" {
  for_each = var.route_keys
  name     = "/aws/lambda/${aws_lambda_function.this[each.key].function_name}"

  retention_in_days = 30
}

resource "aws_iam_role" "this" {
  for_each = var.route_keys
  name     = format("terraform-registry-%s-%s-%s", var.resource_name, var.stage_name, each.key)
  inline_policy {
    name = "db-access"
    policy = jsonencode({
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Sid" : "ListAndDescribe",
          "Effect" : "Allow",
          "Action" : [
            "dynamodb:List*",
            "dynamodb:DescribeReservedCapacity*",
            "dynamodb:DescribeLimits",
            "dynamodb:DescribeTimeToLive"
          ],
          "Resource" : "*"
        },
        {
          "Sid" : "SpecificTable",
          "Effect" : "Allow",
          "Action" : [
            "dynamodb:BatchGet*",
            "dynamodb:DescribeStream",
            "dynamodb:DescribeTable",
            "dynamodb:Get*",
            "dynamodb:Query",
            "dynamodb:Scan",
            "dynamodb:BatchWrite*",
            "dynamodb:CreateTable",
            "dynamodb:Delete*",
            "dynamodb:Update*",
            "dynamodb:PutItem"
          ],
          "Resource" : "${var.dynamodb_arn}"
        }
      ]
    })
  }

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Sid    = ""
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  for_each   = var.route_keys
  role       = aws_iam_role.this[each.key].name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
