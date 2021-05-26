provider "aws" {
  region = "us-east-2"
}

module "inspector" {
  source = "../.."

  create_iam_role = var.create_iam_role
  enabled_rules   = var.enabled_rules

  context = module.this.context
}
