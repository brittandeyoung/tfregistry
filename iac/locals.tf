locals {
  project_name = "opentacos"
  name         = var.pr_id == "none" ? local.project_name : format("%s-%s", local.project_name, var.pr_id)

  api_action = {
    namespace_create = {
      route_key = "POST /namespaces"
      role_arn  = aws_iam_role.full_access.arn
    }
    namespace_read = {
      route_key = "GET /namespaces/{namespace}"
      role_arn  = aws_iam_role.full_access.arn
    }
    namespace_list = {
      route_key = "GET /namespaces"
      role_arn  = aws_iam_role.full_access.arn
    }
    namespace_update = {
      route_key = "PATCH /namespaces/{namespace}"
      role_arn  = aws_iam_role.full_access.arn
    }
    namespace_delete = {
      route_key = "DELETE /namespaces/{namespace}"
      role_arn  = aws_iam_role.full_access.arn
    }
    module_create = {
      route_key = "POST /modules/{namespace}"
      role_arn  = aws_iam_role.full_access.arn
    }
    module_read = {
      route_key = "GET /modules/{namespace}/{name}/{provider}"
      role_arn  = aws_iam_role.full_access.arn
    }
    module_list = {
      route_key = "GET /modules/{namespace}"
      role_arn  = aws_iam_role.full_access.arn
    }
    module_update = {
      route_key = "PATCH /modules/{namespace}/{name}/{provider}"
      role_arn  = aws_iam_role.full_access.arn
    }
    module_delete = {
      route_key = "DELETE /modules/{namespace}/{name}/{provider}"
      role_arn  = aws_iam_role.full_access.arn
    }
    moduleVersion_create = {
      route_key = "POST /modules/{namespace}/{name}/{provider}/versions"
      role_arn  = aws_iam_role.full_access.arn
    }
    moduleVersion_read = {
      route_key = "GET /modules/{namespace}/{name}/{provider}/versions/{version}"
      role_arn  = aws_iam_role.full_access.arn
    }
    moduleVersion_list = {
      route_key = "GET /modules/{namespace}/{name}/{provider}/versions"
      role_arn  = aws_iam_role.full_access.arn
    }
    moduleVersion_delete = {
      route_key = "DELETE /modules/{namespace}/{name}/{provider}/versions/{version}"
      role_arn  = aws_iam_role.full_access.arn
    }
  }
}
