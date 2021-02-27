variable "region" {
  type        = string
  description = "AWS region"
}

variable "create_iam_role" {
  description = "Flag to indicate whether an IAM Role should be created to grant the proper permissions for AWS Config"
  type        = bool
  default     = false
}

variable "enabled_rules" {
  type        = list(string)
  description = <<-DOC
    A list of AWS Inspector rules that should run on a periodic basis. 
    
    For a list of available rules by region, see:
    https://docs.aws.amazon.com/inspector/latest/userguide/inspector_rules-arns.html
  DOC
}
