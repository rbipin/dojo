terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.51.1"
    }
  }
  cloud {
    hostname     = "app.terraform.io"
    organization = "bipin-lab"
    workspaces {
      name = "lab"
    }
  }
}

provider "aws" {
  profile = "default"
  region  = "us-east-1"
}
