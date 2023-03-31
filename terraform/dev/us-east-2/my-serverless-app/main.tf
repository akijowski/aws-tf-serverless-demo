terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
    random = {
      source = "hashicorp/random"
      version = "~> 3.1.0"
    }
  }
  required_version = "~> 1.0"
}

# Configure the AWS Provider
provider "aws" {
  region = "us-east-2"
  profile = "adam"
  default_tags {
    tags = {
      project = "tf-serverless-demo"
      prod = "never"
    }
  }
}

module "app" {
    source = "../../../apps/my-serverless-app"
}
