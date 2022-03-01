package upgrade

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_helper "github.com/lonegunmanb/terraform-module-test-helper"
)

func TestExamplesUpgrade(t *testing.T) {
	test_helper.ModuleUpgradeTest(t, "lonegunmanb", "terraform-module-azure-subnet", "examples/complete", "/src", terraform.Options{
		Upgrade:  true,
		VarFiles: []string{fmt.Sprintf("%s/examples/complete/fixtures.us-east.auto.tfvars", "/src")},
	}, 0)
}
