// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermBatchCertificateInvalidAccountNameRule checks the pattern is valid
type AzurermBatchCertificateInvalidAccountNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermBatchCertificateInvalidAccountNameRule returns new rule with default attributes
func NewAzurermBatchCertificateInvalidAccountNameRule() *AzurermBatchCertificateInvalidAccountNameRule {
	return &AzurermBatchCertificateInvalidAccountNameRule{
		resourceType:  "azurerm_batch_certificate",
		attributeName: "account_name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AzurermBatchCertificateInvalidAccountNameRule) Name() string {
	return "azurerm_batch_certificate_invalid_account_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermBatchCertificateInvalidAccountNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermBatchCertificateInvalidAccountNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermBatchCertificateInvalidAccountNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermBatchCertificateInvalidAccountNameRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9]+$`),
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
