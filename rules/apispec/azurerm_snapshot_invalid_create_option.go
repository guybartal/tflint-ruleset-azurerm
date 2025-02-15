// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermSnapshotInvalidCreateOptionRule checks the pattern is valid
type AzurermSnapshotInvalidCreateOptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermSnapshotInvalidCreateOptionRule returns new rule with default attributes
func NewAzurermSnapshotInvalidCreateOptionRule() *AzurermSnapshotInvalidCreateOptionRule {
	return &AzurermSnapshotInvalidCreateOptionRule{
		resourceType:  "azurerm_snapshot",
		attributeName: "create_option",
		enum: []string{
			"Empty",
			"Attach",
			"FromImage",
			"Import",
			"Copy",
			"Restore",
			"Upload",
		},
	}
}

// Name returns the rule name
func (r *AzurermSnapshotInvalidCreateOptionRule) Name() string {
	return "azurerm_snapshot_invalid_create_option"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermSnapshotInvalidCreateOptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermSnapshotInvalidCreateOptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermSnapshotInvalidCreateOptionRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermSnapshotInvalidCreateOptionRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as create_option`, truncateLongMessage(val)),
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
