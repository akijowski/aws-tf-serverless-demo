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
  region = "us-east-1"
  default_tags {
    tags = {
      team = "kraken"
      project = "tf-serverless-demo"
      prod = "never"
    }
  }
}
