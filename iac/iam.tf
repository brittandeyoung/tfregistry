locals {
  full_access_name = format("%s-%s", local.project_name, "full-access")
  read_access_name = format("%s-%s", local.project_name, "read-only")
}

resource "aws_iam_role" "full_access" {
  name = var.pr_id == "none" ? local.full_access_name : format("%s-%s", local.full_access_name, var.pr_id)
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
          "Resource" : [
            aws_dynamodb_table.this.arn,
          ]
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

resource "aws_iam_role_policy_attachment" "full_access" {
  role       = aws_iam_role.full_access.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role" "read_only" {
  name = var.pr_id == "none" ? local.read_access_name : format("%s-%s", local.read_access_name, var.pr_id)
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
          ],
          "Resource" : [
            aws_dynamodb_table.this.arn,
          ]
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

resource "aws_iam_role_policy_attachment" "read_only" {
  role       = aws_iam_role.read_only.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
