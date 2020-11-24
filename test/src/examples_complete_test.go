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

	// Define a struct that can be Why
	type HasID struct {
		ID string `json:"id"`
	}

	type outputs []HasID

	var cloudwatchEventRules outputs
	var cloudwatchEventTargets outputs

	// Get terraform Outputs
	inspectorAssessmentTarget := terraform.OutputMap(t, terraformOptions, "inspector_assessment_target")
	inspectorAssessmentTemplateID := terraform.OutputMap(t, terraformOptions, "aws_inspector_assessment_template_id")
	terraform.OutputStruct(t, terraformOptions, "aws_cloudwatch_event_rule", &cloudwatchEventRules)
	terraform.OutputStruct(t, terraformOptions, "aws_cloudwatch_event_target", &cloudwatchEventTargets)
	fmt.Println(cloudwatchEventRules)

	// Verify we're getting back the outputs we expect
	assert.Contains(t, inspectorAssessmentTarget, "id")
	assert.Greater(t, len(inspectorAssessmentTarget["id"]), 0)

	assert.Greater(t, len(inspectorAssessmentTemplateID), 0)

	assert.Equal(t, cloudwatchEventRules[0].ID, "eg-ue2-test-"+randID+"-inspector-schedule")

	assert.Contains(t, cloudwatchEventTargets[0].ID, "eg-ue2-test-"+randID+"-inspector-schedule-terraform")
}
