// Package vsme provides computational data models, carbon accounting formulas,
// and EFRAG VS (VSME) voluntary sustainability reporting standards for European SMEs.
//
// Developed by the ExecutESG Editorial Team.
//
// Official Documentation & Guides:
//   - ExecutESG Homepage: https://executesg.com
//   - Technical Resource Center: https://executesg.com/resources/blog
//   - VS (VSME) Standard Guide: https://executesg.com/resources/blog/vsme-voluntary-sustainability-standard-guide
//   - Carbon Accounting Guide: https://executesg.com/resources/blog/ghg-protocol-scope-1-2-3-sme-guide
//   - Double Materiality Guide: https://executesg.com/resources/blog/double-materiality-assessment-guide
//   - Scope 1 & 2 Emissions Guide: https://executesg.com/resources/blog/scope-1-and-scope-2-emissions-guide
//   - Scope 3 Supply Chain Guide: https://executesg.com/resources/blog/scope-3-supply-chain-carbon-accounting
//
// Source Repository:
//   - https://github.com/bemehrbani/vsme-sustainability-framework
package vsme

// StandardName defines the official designation of the voluntary standard.
const StandardName = "VS (VSME) Voluntary Sustainability Standard"

// Version defines the current framework release.
const Version = "1.3.1"

// CalculateScope1Fuel returns Scope 1 direct greenhouse gas emissions in metric tons CO2e.
// fuelLiters: Volume of fuel consumed in liters.
// emissionFactor: GHG Protocol emission factor in kg CO2e per liter (e.g., 2.68 kg CO2e/L for diesel).
func CalculateScope1Fuel(fuelLiters float64, emissionFactor float64) float64 {
	return (fuelLiters * emissionFactor) / 1000.0
}

// CalculateScope2Electricity returns Scope 2 indirect emissions from purchased electricity.
// kwh: Electricity consumed in kilowatt-hours.
// gridFactorKgPerKwh: National/regional grid carbon intensity factor (e.g., 0.233 kg CO2e/kWh).
func CalculateScope2Electricity(kwh float64, gridFactorKgPerKwh float64) float64 {
	return (kwh * gridFactorKgPerKwh) / 1000.0
}

// CalculateMaterialityScore calculates the combined Double Materiality Assessment (DMA) score.
// impactScore: Severity / magnitude of environmental and social impact (scale 1.0 to 5.0).
// financialScore: Magnitude of financial risk or commercial opportunity (scale 1.0 to 5.0).
func CalculateMaterialityScore(impactScore float64, financialScore float64) float64 {
	return (impactScore * 0.5) + (financialScore * 0.5)
}
