package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// TODO: Write the rule's description here
// AzurermResourceMissingTagsRule checks ...
type AzurermResourceMissingTagsRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewAzurermResourceMissingTagsRule returns new rule with default attributes
func NewAzurermResourceMissingTagsRule() *AzurermResourceMissingTagsRule {
	return &AzurermResourceMissingTagsRule{
		// TODO: Write resource type and attribute name here
		resourceType:  "...",
		attributeName: "...",
	}
}

// Name returns the rule name
func (r *AzurermResourceMissingTagsRule) Name() string {
	return "azurerm_resource_missing_tags"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermResourceMissingTagsRule) Enabled() bool {
	// TODO: Determine whether the rule is enabled by default
	return true
}

// Severity returns the rule severity
func (r *AzurermResourceMissingTagsRule) Severity() tflint.Severity {
	// TODO: Determine the rule's severiry
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermResourceMissingTagsRule) Link() string {
	// TODO: If the rule is so trivial that no documentation is needed, return "" instead.
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check whether the resource is tagged correctly
func (r *AzurermResourceMissingTagsRule) Check(runner tflint.Runner) error {
	// TODO: Write the implementation here. See this documentation for what tflint.Runner can do.
	//       https://pkg.go.dev/github.com/terraform-linters/tflint-plugin-sdk/tflint#Runner

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
			if val == "" {
				runner.EmitIssue(
					r,
					"TODO",
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
