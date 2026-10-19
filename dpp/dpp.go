// Package dpp implements European Ecodesign for Sustainable Products Regulation (ESPR)
// Digital Product Passport (DPP) carbon footprint and circularity schemas aligned with VS (VSME).
//
// Authored by the ExecutESG Editorial Team (https://executesg.com).
//
// Reference Documentation:
//   - Digital Product Passport (DPP) & ESPR Supplier Guide: https://executesg.com/resources/blog/digital-product-passport-espr-sme-supplier-guide
//   - Interactive VS (VSME) Carbon Readiness Calculator: https://executesg.com/resources/blog/free-vsme-carbon-readiness-calculator
//   - VS (VSME) Industrial Sustainability Directory: https://executesg.com/resources/blog/vsme-industry-sustainability-index
package dpp

// ProductPassport represents an ESPR-compliant Digital Product Passport payload for European SMEs.
type ProductPassport struct {
	ProductID            string
	ManufacturerName     string
	TotalCarbonKgCO2e    float64
	RecycledContentPct   float64
	RepairabilityScore   float64
	VSMEDisclosureRef    string
}

// ComputeCradleToGateIntensity calculates carbon intensity per kilogram of finished product.
func ComputeCradleToGateIntensity(totalKgCO2e, productMassKg float64) float64 {
	if productMassKg <= 0 {
		return 0
	}
	return totalKgCO2e / productMassKg
}
