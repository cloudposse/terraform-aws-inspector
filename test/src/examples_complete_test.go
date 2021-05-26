package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

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

	var inspectorAssessmentTarget map[string]interface{}
	var inspectorAssessmentTemplate map[string]interface{}
	var cloudwatchEventRule map[string]interface{}
	var cloudwatchEventTarget map[string]interface{}

	// Get terraform Outputs
	terraform.OutputStruct(t, terraformOptions, "inspector_assessment_target", &inspectorAssessmentTarget)
	terraform.OutputStruct(t, terraformOptions, "aws_inspector_assessment_template", &inspectorAssessmentTemplate)
	terraform.OutputStruct(t, terraformOptions, "aws_cloudwatch_event_rule", &cloudwatchEventRule)
	terraform.OutputStruct(t, terraformOptions, "aws_cloudwatch_event_target", &cloudwatchEventTarget)
	fmt.Println(cloudwatchEventRule)

	// Verify we're getting back the outputs we expect
	assert.Contains(t, inspectorAssessmentTarget, "id")
	assert.Contains(t, inspectorAssessmentTemplate, "id")

	assert.Equal(t, cloudwatchEventRule["id"], "eg-ue2-test-"+randID+"-inspector-schedule")
	assert.Contains(t, cloudwatchEventTarget["id"], "eg-ue2-test-"+randID+"-inspector-schedule-terraform")
}
