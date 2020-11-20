#-----------------------------------------------------------------------------------------------------------------------
# Setup AWS Inspector
#-----------------------------------------------------------------------------------------------------------------------
module "insepector_assessment_target_label" {
  source  = "cloudposse/label/null"
  version = "0.21.0"

  attributes = concat(module.this.attributes, ["inspector", "assessment", "target"])
  context    = module.this.context
}

resource "aws_inspector_assessment_target" "target" {
  name = module.insepector_assessment_target_label.id
}

module "insepector_assessment_template_label" {
  source  = "cloudposse/label/null"
  version = "0.21.0"

  attributes = concat(module.this.attributes, ["inspector", "assessment", "template"])
  context    = module.this.context
}

resource "aws_inspector_assessment_template" "assessment" {
  name               = module.insepector_assessment_template_label.id
  target_arn         = aws_inspector_assessment_target.target.arn
  duration           = var.assessment_duration
  rules_package_arns = var.enabled_rules
}

#-----------------------------------------------------------------------------------------------------------------------
# Create a scheduled event to run inspector
#-----------------------------------------------------------------------------------------------------------------------
module "insepector_schedule_label" {
  source  = "cloudposse/label/null"
  version = "0.21.0"

  attributes = concat(module.this.attributes, ["inspector", "schedule"])
  context    = module.this.context
}

resource "aws_cloudwatch_event_rule" "schedule" {
  count               = local.create_scheduled_event ? 1 : 0
  name                = module.insepector_schedule_label.id
  description         = "Trigger an AWS Inspector Assessment"
  schedule_expression = var.schedule_expression
}

resource "aws_cloudwatch_event_target" "target" {
  count    = local.create_scheduled_event ? 1 : 0
  rule     = aws_cloudwatch_event_rule.schedule[0].name
  arn      = aws_inspector_assessment_template.assessment.arn
  role_arn = local.create_iam_role ? module.iam_role[0].arn : var.iam_role_arn
}

#-----------------------------------------------------------------------------------------------------------------------
# Optionally create an IAM Role 
#-----------------------------------------------------------------------------------------------------------------------
module "iam_role" {
  count   = local.create_iam_role ? 1 : 0
  source  = "cloudposse/iam-role/aws"
  version = "0.6.1"

  principals = {
    "Service" = ["events.amazonaws.com"]
  }

  use_fullname = true

  policy_documents = [
    data.aws_iam_policy_document.start_assessment_policy[0].json,
  ]

  policy_document_count = 1
  policy_description    = "AWS Inspector IAM policy"
  role_description      = "AWS Inspector IAM role"

  context = module.this.context
}

data "aws_iam_policy_document" "start_assessment_policy" {
  count = local.create_iam_role ? 1 : 0

  statement {
    sid       = "StartAssessment"
    effect    = "Allow"
    resources = ["*"]
    actions   = ["inspector:StartAssessmentRun"]
  }
}

#-----------------------------------------------------------------------------------------------------------------------
# Locals and Data References
#-----------------------------------------------------------------------------------------------------------------------
locals {
  create_iam_role        = module.this.enabled && var.create_iam_role
  create_scheduled_event = module.this.enabled
}
