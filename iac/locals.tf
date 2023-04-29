locals {
  api = {
    module = {
      create = "POST /modules/{namespace}"
      delete = "DELETE /modules/{namespace}/{name}/{provider}"
      list   = "GET /modules/{namespace}"
      read   = "GET /modules/{namespace}/{name}/{provider}"
      update = "PATCH /modules/{namespace}/{name}/{provider}"
    }
    namespace = {
      create = "POST /namespaces"
      delete = "DELETE /namespaces/{namespace}/{name}/{provider}"
      list   = "GET /namespaces/{namespace}"
      read   = "GET /namespaces/{namespace}/{name}/{provider}"
      update = "PATCH /namespaces/{namespace}/{name}/{provider}"
    }
  }
  default_tags = {
    "Name"        = "Terraform Registry API"
    "Project"     = "tfregistry-api"
    "Environment" = terraform.workspace
  }
}
