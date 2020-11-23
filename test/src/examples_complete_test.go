package test

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func listContains(s []string, e string) bool {
	for _, a := range s {
		if a == e || strings.Contains(a, e) {
			return true
		}
	}
	return false
}

// Test the Terraform module in examples/complete using Terratest.
func TestExamplesComplete(t *testing.T) {
	t.Parallel()

	// We always include a random attribute so that parallel tests and AWS resources do not interfere with each
	// other
	rand.Seed(time.Now().UnixNano())
	randID := strconv.Itoa(rand.Intn(100000))
	attributes := []string{randID}

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../../examples/complete",
		Upgrade:      true,
		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"fixtures.us-east-2.tfvars"},
		Vars: map[string]interface{}{
			"attributes": attributes,
		},
	}
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Get terraform Outputs
	inspectorAssessmentTarget := terraform.OutputMap(t, terraformOptions, "inspector_assessment_target")
	inspectorAssessmentTemplateID := terraform.OutputMap(t, terraformOptions, "aws_inspector_assessment_template_id")
	cloudwatchEventRule := terraform.OutputMap(t, terraformOptions, "aws_cloudwatch_event_rule")
	cloudwatchEventTarget := terraform.OutputMap(t, terraformOptions, "aws_cloudwatch_event_target")

	// Verify we're getting back the outputs we expect
	assert.Contains(t, inspectorAssessmentTarget, "id")
	assert.Greater(t, len(inspectorAssessmentTarget["id"]), 0)

	assert.Greater(t, len(inspectorAssessmentTemplateID), 0)

	assert.Contains(t, cloudwatchEventRule, "id")
	assert.Greater(t, len(cloudwatchEventRule["id"]), 0)

	assert.Contains(t, cloudwatchEventTarget, "id")
	assert.Greater(t, len(cloudwatchEventTarget["id"]), 0)
}
