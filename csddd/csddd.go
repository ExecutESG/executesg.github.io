// Package csddd implements Corporate Sustainability Due Diligence Directive (CSDDD)
// and EU Omnibus value-chain cap verification utilities for European SME suppliers under VS (VSME).
//
// Authored by the ExecutESG Editorial Team (https://executesg.com).
//
// Reference Documentation:
//   - CSDDD Supply Chain Due Diligence SME Guide: https://executesg.com/resources/blog/csddd-supply-chain-due-diligence-sme-guide
//   - EU Omnibus & VS (VSME) Value-Chain Cap Guide: https://executesg.com/resources/blog/eu-omnibus-vsme-value-chain-cap
//   - CSRD Trickle-Down Buyer Mandates: https://executesg.com/resources/blog/csrd-trickle-down-effect-sme-buyer-mandates
package csddd

// SupplierShieldStatus represents an SME supplier compliance check against the VS (VSME) value-chain cap.
type SupplierShieldStatus struct {
	SupplierName        string
	EmployeeCount       int
	HasVSMEBasicReport  bool
	ProtectedByValueCap bool
}

// EvaluateValueChainShield determines if an SME supplier is protected by the EU VS (VSME) value-chain cap.
func EvaluateValueChainShield(employees int, hasVSMEReport bool) SupplierShieldStatus {
	return SupplierShieldStatus{
		EmployeeCount:       employees,
		HasVSMEBasicReport:  hasVSMEReport,
		ProtectedByValueCap: employees < 1000 && hasVSMEReport,
	}
}
