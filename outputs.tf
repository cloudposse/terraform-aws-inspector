output "inspector_assessment_target" {
  description = "The AWS Inspector assessment target"
  value       = try(aws_inspector_assessment_target.target[0], {})
}

output "aws_inspector_assessment_template" {
  description = "The AWS Inspector assessment template"
  value       = try(aws_inspector_assessment_template.assessment[0], {})
}

output "aws_cloudwatch_event_rule" {
  description = "The AWS Inspector event rule"
  value       = try(aws_cloudwatch_event_rule.schedule[0], {})
}

output "aws_cloudwatch_event_target" {
  description = "The AWS Inspector event target"
  value       = try(aws_cloudwatch_event_target.target[0], {})
}
