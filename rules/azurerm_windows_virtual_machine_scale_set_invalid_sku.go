package rules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermWindowsVirtualMachineScaleSetInvalidSkuRule checks the pattern is valid
type AzurermWindowsVirtualMachineScaleSetInvalidSkuRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermWindowsVirtualMachineScaleSetInvalidSkuRule returns new rule with default attributes
func NewAzurermWindowsVirtualMachineScaleSetInvalidSkuRule() *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule {
	return &AzurermWindowsVirtualMachineScaleSetInvalidSkuRule{
		resourceType:  "azurerm_windows_virtual_machine_scale_set",
		attributeName: "sku",
		enum: []string{
			"Basic_A0",
			"Basic_A1",
			"Basic_A2",
			"Basic_A3",
			"Basic_A4",
			"Standard_A0",
			"Standard_A1",
			"Standard_A2",
			"Standard_A3",
			"Standard_A4",
			"Standard_A5",
			"Standard_A6",
			"Standard_A7",
			"Standard_A8",
			"Standard_A9",
			"Standard_A10",
			"Standard_A11",
			"Standard_A1_v2",
			"Standard_A2_v2",
			"Standard_A4_v2",
			"Standard_A8_v2",
			"Standard_A2m_v2",
			"Standard_A4m_v2",
			"Standard_A8m_v2",
			"Standard_B1ls",
			"Standard_B1s",
			"Standard_B1ms",
			"Standard_B2s",
			"Standard_B2ms",
			"Standard_B4ms",
			"Standard_B8ms",
			"Standard_B12ms",
			"Standard_B16ms",
			"Standard_B20ms",
			"Standard_D1",
			"Standard_D2",
			"Standard_D3",
			"Standard_D4",
			"Standard_D11",
			"Standard_D12",
			"Standard_D13",
			"Standard_D14",
			"Standard_D1_v2",
			"Standard_D2_v2",
			"Standard_D3_v2",
			"Standard_D4_v2",
			"Standard_D5_v2",
			"Standard_D2_v3",
			"Standard_D4_v3",
			"Standard_D8_v3",
			"Standard_D16_v3",
			"Standard_D32_v3",
			"Standard_D48_v3",
			"Standard_D64_v3",
			"Standard_D2s_v3",
			"Standard_D4s_v3",
			"Standard_D8s_v3",
			"Standard_D16s_v3",
			"Standard_D32s_v3",
			"Standard_D48s_v3",
			"Standard_D64s_v3",
			"Standard_D2a_v4",
			"Standard_D4a_v4",
			"Standard_D8a_v4",
			"Standard_D16a_v4",
			"Standard_D32a_v4",
			"Standard_D48a_v4",
			"Standard_D64a_v4",
			"Standard_D96a_v4",
			"Standard_D2as_v4",
			"Standard_D4as_v4",
			"Standard_D8as_v4",
			"Standard_D16as_v4",
			"Standard_D32as_v4",
			"Standard_D48as_v4",
			"Standard_D64as_v4",
			"Standard_D96as_v4",
			"Standard_D2d_v4",
			"Standard_D4d_v4",
			"Standard_D8d_v4",
			"Standard_D16d_v4",
			"Standard_D32d_v4",
			"Standard_D48d_v4",
			"Standard_D64d_v4",
			"Standard_D2ds_v4",
			"Standard_D4ds_v4",
			"Standard_D8ds_v4",
			"Standard_D16ds_v4",
			"Standard_D32ds_v4",
			"Standard_D48ds_v4",
			"Standard_D64ds_v4",
			"Standard_D2_v4",
			"Standard_D4_v4",
			"Standard_D8_v4",
			"Standard_D16_v4",
			"Standard_D32_v4",
			"Standard_D48_v4",
			"Standard_D64_v4",
			"Standard_D2s_v4",
			"Standard_D4s_v4",
			"Standard_D8s_v4",
			"Standard_D16s_v4",
			"Standard_D32s_v4",
			"Standard_D48s_v4",
			"Standard_D64s_v4",
			"Standard_D11_v2",
			"Standard_D12_v2",
			"Standard_D13_v2",
			"Standard_D14_v2",
			"Standard_D15_v2",
			"Standard_DC1s_v2",
			"Standard_DC2s_v2",
			"Standard_DC4s_v2",
			"Standard_DC8_v2",
			"Standard_DS1",
			"Standard_DS2",
			"Standard_DS3",
			"Standard_DS4",
			"Standard_DS11",
			"Standard_DS12",
			"Standard_DS13",
			"Standard_DS14",
			"Standard_DS1_v2",
			"Standard_DS2_v2",
			"Standard_DS3_v2",
			"Standard_DS4_v2",
			"Standard_DS5_v2",
			"Standard_DS11_v2",
			"Standard_DS12_v2",
			"Standard_DS13_v2",
			"Standard_DS14_v2",
			"Standard_DS15_v2",
			"Standard_DS13-4_v2",
			"Standard_DS13-2_v2",
			"Standard_DS14-8_v2",
			"Standard_DS14-4_v2",
			"Standard_E2_v3",
			"Standard_E4_v3",
			"Standard_E8_v3",
			"Standard_E16_v3",
			"Standard_E20_v3",
			"Standard_E32_v3",
			"Standard_E48_v3",
			"Standard_E64_v3",
			"Standard_E64i_v3",
			"Standard_E2s_v3",
			"Standard_E4s_v3",
			"Standard_E8s_v3",
			"Standard_E16s_v3",
			"Standard_E20s_v3",
			"Standard_E32s_v3",
			"Standard_E48s_v3",
			"Standard_E64s_v3",
			"Standard_E64is_v3",
			"Standard_E32-16_v3",
			"Standard_E32-8s_v3",
			"Standard_E64-32s_v3",
			"Standard_E64-16s_v3",
			"Standard_E2a_v4",
			"Standard_E4a_v4",
			"Standard_E8a_v4",
			"Standard_E16a_v4",
			"Standard_E20a_v4",
			"Standard_E32a_v4",
			"Standard_E48a_v4",
			"Standard_E64a_v4",
			"Standard_E96a_v4",
			"Standard_E2as_v4",
			"Standard_E4as_v4",
			"Standard_E8as_v4",
			"Standard_E16as_v4",
			"Standard_E20as_v4",
			"Standard_E32as_v4",
			"Standard_E48as_v4",
			"Standard_E64as_v4",
			"Standard_E96as_v4",
			"Standard_E2d_v4",
			"Standard_E4d_v4",
			"Standard_E8d_v4",
			"Standard_E16d_v4",
			"Standard_E20d_v4",
			"Standard_E32d_v4",
			"Standard_E48d_v4",
			"Standard_E64d_v4",
			"Standard_E2ds_v4",
			"Standard_E4ds_v4",
			"Standard_E8ds_v4",
			"Standard_E16ds_v4",
			"Standard_E20ds_v4",
			"Standard_E32ds_v4",
			"Standard_E48ds_v4",
			"Standard_E64ds_v4",
			"Standard_E80ids_v4",
			"Standard_E2_v4",
			"Standard_E4_v4",
			"Standard_E8_v4",
			"Standard_E16_v4",
			"Standard_E20_v4",
			"Standard_E32_v4",
			"Standard_E48_v4",
			"Standard_E64_v4",
			"Standard_E2s_v4",
			"Standard_E4s_v4",
			"Standard_E8s_v4",
			"Standard_E16s_v4",
			"Standard_E20s_v4",
			"Standard_E32s_v4",
			"Standard_E48s_v4",
			"Standard_E64s_v4",
			"Standard_E80is_v4",
			"Standard_M8ms",
			"Standard_M16ms",
			"Standard_M32ts",
			"Standard_M32ls",
			"Standard_M32ms",
			"Standard_M64s",
			"Standard_M64ls",
			"Standard_M64ms",
			"Standard_M128s",
			"Standard_M128ms",
			"Standard_M64",
			"Standard_M64m",
			"Standard_M128",
			"Standard_M128m",
			"Standard_M208ms_v2",
			"Standard_M208s_v2",
			"Standard_M416ms_v2",
			"Standard_M416s_v2",
			"Standard_M8-2ms",
			"Standard_M8-4ms",
			"Standard_M16-4ms",
			"Standard_M16-8ms",
			"Standard_M32-8ms",
			"Standard_M32-16ms",
			"Standard_M64-32ms",
			"Standard_M64-16ms",
			"Standard_M128-64ms",
			"Standard_M128-32ms",
			"Standard_E4-2s_v3",
			"Standard_E8-4s_v3",
			"Standard_E8-2s_v3",
			"Standard_E16-8s_v3",
			"Standard_E16-4s_v3",
			"Standard_E32-16s_v3",
			"Standard_E32-8s_v3",
			"Standard_E64-32s_v3",
			"Standard_E64-16s_v3",
			"Standard_E4-2s_v4",
			"Standard_E8-4s_v4",
			"Standard_E8-2s_v4",
			"Standard_E16-8s_v4",
			"Standard_E16-4s_v4",
			"Standard_E32-16s_v4",
			"Standard_E32-8s_v4",
			"Standard_E64-32s_v4",
			"Standard_E64-16s_v4",
			"Standard_E4-2ds_v4",
			"Standard_E8-4ds_v4",
			"Standard_E8-2ds_v4",
			"Standard_E16-8ds_v4",
			"Standard_E16-4ds_v4",
			"Standard_E32-16ds_v4",
			"Standard_E32-8ds_v4",
			"Standard_E64-32ds_v4",
			"Standard_E64-16ds_v4",
			"Standard_GS4-8",
			"Standard_GS4-4",
			"Standard_GS5-16",
			"Standard_GS5-8",
			"Standard_DS11-1_v2",
			"Standard_DS12-2_v2",
			"Standard_DS12-1_v2",
			"Standard_DS13-4_v2",
			"Standard_DS13-2_v2",
			"Standard_DS14-8_v2",
			"Standard_DS14-4_v2",
			"Standard_M416-208s_v2",
			"Standard_M416-208ms_v2",
			"Standard_F1",
			"Standard_F2",
			"Standard_F4",
			"Standard_F8",
			"Standard_F16",
			"Standard_F1s",
			"Standard_F2s",
			"Standard_F4s",
			"Standard_F8s",
			"Standard_F16s",
			"Standard_F2s_v2",
			"Standard_F4s_v2",
			"Standard_F8s_v2",
			"Standard_F16s_v2",
			"Standard_F32s_v2",
			"Standard_F48s_v2",
			"Standard_F64s_v2",
			"Standard_F72s_v2",
			"Standard_FX4mds",
			"Standard_FX12mds",
			"Standard_FX24mds",
			"Standard_FX36mds",
			"Standard_FX48mds",
			"Standard_G1",
			"Standard_G2",
			"Standard_G3",
			"Standard_G4",
			"Standard_G5",
			"Standard_GS1",
			"Standard_GS2",
			"Standard_GS3",
			"Standard_GS4",
			"Standard_GS5",
			"Standard_GS4-8",
			"Standard_GS4-4",
			"Standard_GS5-16",
			"Standard_GS5-8",
			"Standard_H8",
			"Standard_H16",
			"Standard_H8m",
			"Standard_H16m",
			"Standard_H16r",
			"Standard_H16mr",
			"Standard_HB60rs",
			"Standard_HB120rs_v2",
			"Standard_HB120rs_v3",
			"Standard_HB120-96rs_v3",
			"Standard_HB120-64rs_v3",
			"Standard_HB120-32rs_v3",
			"Standard_HB120-16rs_v3",
			"Standard_HC44rs",
			"Standard_L4s",
			"Standard_L8s",
			"Standard_L16s",
			"Standard_L32s",
			"Standard_L8s_v2",
			"Standard_L16s_v2",
			"Standard_L32s_v2",
			"Standard_L48s_v2",
			"Standard_L64s_v2",
			"Standard_L80s_v2",
			"Standard_M64s",
			"Standard_M64ms",
			"Standard_M128s",
			"Standard_M128ms",
			"Standard_M64-32ms",
			"Standard_M64-16ms",
			"Standard_M128-64ms",
			"Standard_M128-32ms",
			"Standard_NC6",
			"Standard_NC12",
			"Standard_NC24",
			"Standard_NC24r",
			"Standard_NC6s_v2",
			"Standard_NC12s_v2",
			"Standard_NC24s_v2",
			"Standard_NC24rs_v2",
			"Standard_NC6s_v3",
			"Standard_NC12s_v3",
			"Standard_NC24s_v3",
			"Standard_NC24rs_v3",
			"Standard_NC4as_T4_v3",
			"Standard_NC8as_T4_v3",
			"Standard_NC16as_T4_v3",
			"Standard_NC64as_T4_v3",
			"Standard_ND6s",
			"Standard_ND12s",
			"Standard_ND24s",
			"Standard_ND24rs",
			"Standard_ND40rs_v2",
			"Standard_ND96asr_v4",
			"Standard_NV6",
			"Standard_NV12",
			"Standard_NV24",
			"Standard_NV12s_v3",
			"Standard_NV24s_v3",
			"Standard_NV48s_v3",
			"Standard_NV4as_v4",
			"Standard_NV8as_v4",
			"Standard_NV16as_v4",
			"Standard_NV32as_v4",
		},
	}
}

// Name returns the rule name
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Name() string {
	return "azurerm_windows_virtual_machine_scale_set_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermWindowsVirtualMachineScaleSetInvalidSkuRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as sku`, val),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
