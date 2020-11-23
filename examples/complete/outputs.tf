output "inspector_assessment_target" {
  description = "The AWS Inspector assessment target"
  value       = module.inspector.inspector_assessment_target
}

output "aws_inspector_assessment_template_id" {
  description = "The AWS Inspector assessment template ids"
  value       = module.inspector.aws_inspector_assessment_template
}

output "aws_cloudwatch_event_rule" {
  description = "The AWS Inspector event rule"
  # hardcoding this to the [0] instance for this example until the following issue is resolved:
  # https://github.com/gruntwork-io/terratest/issues/710
  value = module.inspector.aws_cloudwatch_event_rule[0]
}

output "aws_cloudwatch_event_target" {
  description = "The AWS Inspector event target"
  value       = module.inspector.aws_cloudwatch_event_target[0]
}
