package test

import (
	"testing"

	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Test the Terraform module in examples/complete using Terratest.
func TestExamplesComplete(t *testing.T) {
	t.Parallel()

	rootFolder := "../../"
	terraformFolderRelativeToRoot := "examples/complete"
	terraformDir := test_structure.CopyTerraformFolderToTemp(t, rootFolder, terraformFolderRelativeToRoot)

	vars := map[string]interface{}{
		"vnet_cidrs": []string{"10.0.0.0/16"},
	}
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: terraformDir,
		Upgrade:      true,
		// Variables to pass to our Terraform code using -var-file options
		//VarFiles: varFiles,
		Vars: vars,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors, then run `terraform
	terraform.InitAndApplyAndIdempotent(t, terraformOptions)

	privateSubnetId := terraform.Output(t, terraformOptions, "private_subnet_id")
	assert.NotEqual(t, "", privateSubnetId)
}