// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule checks the pattern is valid
type AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermVirtualMachineScaleSetInvalidEvictionPolicyRule returns new rule with default attributes
func NewAzurermVirtualMachineScaleSetInvalidEvictionPolicyRule() *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule {
	return &AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule{
		resourceType:  "azurerm_virtual_machine_scale_set",
		attributeName: "eviction_policy",
		enum: []string{
			"Deallocate",
			"Delete",
		},
	}
}

// Name returns the rule name
func (r *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule) Name() string {
	return "azurerm_virtual_machine_scale_set_invalid_eviction_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualMachineScaleSetInvalidEvictionPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as eviction_policy`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
