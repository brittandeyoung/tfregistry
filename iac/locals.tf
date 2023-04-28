locals {
  api = {
    module = {
      create = "POST /modules"
      delete = "DELETE /modules/{namespace}/{name}/{provider}"
      list   = "GET /modules/{namespace}"
      read   = "GET /modules/{namespace}/{name}/{provider}"
      update = "PATCH /modules/{namespace}/{name}/{provider}"
    }
  }
  default_tags = {
    "Name"        = "Terraform Registry API"
    "Project"     = "tfregistry-api"
    "Environment" = terraform.workspace
  }
}
