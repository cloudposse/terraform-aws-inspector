provider "aws" {
  region = var.region
}
module "inspector" {
  source = "../.."

  create_iam_role = var.create_iam_role
  enabled_rules   = var.enabled_rules

  context = module.this.context
}
