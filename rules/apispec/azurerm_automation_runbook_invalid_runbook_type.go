// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermAutomationRunbookInvalidRunbookTypeRule checks the pattern is valid
type AzurermAutomationRunbookInvalidRunbookTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermAutomationRunbookInvalidRunbookTypeRule returns new rule with default attributes
func NewAzurermAutomationRunbookInvalidRunbookTypeRule() *AzurermAutomationRunbookInvalidRunbookTypeRule {
	return &AzurermAutomationRunbookInvalidRunbookTypeRule{
		resourceType:  "azurerm_automation_runbook",
		attributeName: "runbook_type",
		enum: []string{
			"Script",
			"Graph",
			"PowerShellWorkflow",
			"PowerShell",
			"GraphPowerShellWorkflow",
			"GraphPowerShell",
			"Python2",
			"Python3",
		},
	}
}

// Name returns the rule name
func (r *AzurermAutomationRunbookInvalidRunbookTypeRule) Name() string {
	return "azurerm_automation_runbook_invalid_runbook_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermAutomationRunbookInvalidRunbookTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermAutomationRunbookInvalidRunbookTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermAutomationRunbookInvalidRunbookTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermAutomationRunbookInvalidRunbookTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as runbook_type`, truncateLongMessage(val)),
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
