terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "<= 5.0.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
  default_tags {
    tags = {
      "Name"        = "Terraform Registry API"
      "Project"     = "tfregistry-api"
      "Environment" = terraform.workspace
    }
  }
}


