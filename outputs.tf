output "inspector_assessment_target" {
  description = "The AWS Inspector assessment target"
  value       = aws_inspector_assessment_target.target
}

output "aws_inspector_assessment_template" {
  description = "The AWS Inspector assessment template"
  value       = aws_inspector_assessment_template.assessment
}

output "aws_cloudwatch_event_rule" {
  description = "The AWS Inspector event rule"
  value       = aws_cloudwatch_event_rule.schedule
}

output "aws_cloudwatch_event_target" {
  description = "The AWS Inspector event target"
  value       = aws_cloudwatch_event_target.target
}
