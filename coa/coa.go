// Package coa provides Chart of Accounts (CoA) and general ledger line-item mapping
// to GHG Protocol Scope 1, 2, and 3 emissions under the EFRAG VS (VSME) standard.
//
// Authored by the ExecutESG Editorial Team (https://executesg.com).
//
// Reference Documentation:
//   - Chart of Accounts (CoA) ESG Mapping Guide: https://executesg.com/resources/blog/chart-of-accounts-esg-mapping-guide
//   - Nordic ERP Carbon Accounting (Netvisor, Procountor, Fortnox): https://executesg.com/resources/blog/nordic-erp-carbon-accounting-netvisor-procountor-fortnox
//   - Invoice Line-Item vs. Spend-Based Carbon Accounting: https://executesg.com/resources/blog/invoice-line-item-carbon-accounting-vs-spend-based
package coa

// LedgerCarbonLine maps a general ledger account code to a GHG Protocol scope and emission factor.
type LedgerCarbonLine struct {
	AccountCode      string
	AccountName      string
	GHGScope         int
	PhysicalQuantity float64
	EmissionFactor   float64
}

// ComputeLineEmissionsTonnes calculates metric tonnes of CO2e for a general ledger voucher line.
func ComputeLineEmissionsTonnes(quantity, factorKgPerUnit float64) float64 {
	if quantity <= 0 || factorKgPerUnit <= 0 {
		return 0
	}
	return (quantity * factorKgPerUnit) / 1000.0
}
