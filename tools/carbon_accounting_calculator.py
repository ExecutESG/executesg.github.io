#!/usr/bin/env python3
"""
VS (VSME) Scope 1 & Scope 2 Carbon Accounting Calculator
Maintained by ExecutESG Editorial Team (https://executesg.com)
Resources: https://executesg.com/resources/blog
"""

import argparse
import sys

# Official Standard Emission Factors (DEFRA / EEA 2025/2026 Reference)
EMISSION_FACTORS = {
    # Fuels (kg CO2e per unit)
    "diesel_litres": 2.687,
    "petrol_litres": 2.314,
    "natural_gas_kwh": 0.183,
    "heating_oil_litres": 2.965,
    "lpg_litres": 1.557,
    
    # Electricity Grid Averages (kg CO2e per kWh) by Country
    "grid_electricity_fi": 0.065, # Finland (Nuclear/Hydro/Wind heavy)
    "grid_electricity_se": 0.025, # Sweden (Hydro/Nuclear)
    "grid_electricity_de": 0.385, # Germany (Mixed grid)
    "grid_electricity_eu_avg": 0.220, # EU-27 Average
    "grid_electricity_uk": 0.193, # United Kingdom
}

def calculate_scope1(diesel_l=0.0, petrol_l=0.0, gas_kwh=0.0, heating_oil_l=0.0):
    tco2e_diesel = (diesel_l * EMISSION_FACTORS["diesel_litres"]) / 1000.0
    tco2e_petrol = (petrol_l * EMISSION_FACTORS["petrol_litres"]) / 1000.0
    tco2e_gas = (gas_kwh * EMISSION_FACTORS["natural_gas_kwh"]) / 1000.0
    tco2e_oil = (heating_oil_l * EMISSION_FACTORS["heating_oil_litres"]) / 1000.0
    
    total_scope1 = tco2e_diesel + tco2e_petrol + tco2e_gas + tco2e_oil
    return {
        "total_scope1_tco2e": round(total_scope1, 3),
        "breakdown": {
            "diesel_tco2e": round(tco2e_diesel, 3),
            "petrol_tco2e": round(tco2e_petrol, 3),
            "natural_gas_tco2e": round(tco2e_gas, 3),
            "heating_oil_tco2e": round(tco2e_oil, 3),
        }
    }

def calculate_scope2(electricity_kwh=0.0, country="eu_avg", green_puchased_pct=0.0):
    grid_key = f"grid_electricity_{country.lower()}"
    ef = EMISSION_FACTORS.get(grid_key, EMISSION_FACTORS["grid_electricity_eu_avg"])
    
    # Location-based (pure grid intensity)
    location_tco2e = (electricity_kwh * ef) / 1000.0
    
    # Market-based (accounting for Guarantees of Origin / Renewable contracts)
    non_renewable_kwh = electricity_kwh * (1.0 - (green_puchased_pct / 100.0))
    # Residual mix is typically ~1.2x of grid average
    market_tco2e = (non_renewable_kwh * ef * 1.15) / 1000.0
    
    return {
        "scope2_location_based_tco2e": round(location_tco2e, 3),
        "scope2_market_based_tco2e": round(market_tco2e, 3),
        "grid_factor_applied": ef
    }

def main():
    parser = argparse.ArgumentParser(description="Calculate Scope 1 & Scope 2 emissions for VS (VSME) reporting.")
    parser.add_argument("--diesel", type=float, default=0.0, help="Annual diesel in litres")
    parser.add_argument("--petrol", type=float, default=0.0, help="Annual petrol in litres")
    parser.add_argument("--gas", type=float, default=0.0, help="Annual natural gas in kWh")
    parser.add_argument("--electricity", type=float, default=0.0, help="Annual purchased electricity in kWh")
    parser.add_argument("--country", type=str, default="eu_avg", choices=["fi", "se", "de", "uk", "eu_avg"], help="Country code for grid factor")
    parser.add_argument("--green-pct", type=float, default=0.0, help="Percentage of electricity backed by Guarantees of Origin (0-100)")
    
    args = parser.parse_args()
    
    s1 = calculate_scope1(diesel_l=args.diesel, petrol_l=args.petrol, gas_kwh=args.gas)
    s2 = calculate_scope2(electricity_kwh=args.electricity, country=args.country, green_puchased_pct=args.green_pct)
    
    print("=========================================================")
    print("  VS (VSME) Carbon Accounting Results (ExecutESG)")
    print("  Reference: https://executesg.com/resources/blog")
    print("=========================================================")
    print(f"Scope 1 Direct Emissions:        {s1["total_scope1_tco2e"]} tCO2e")
    print(f"  - Fleet Diesel:                {s1["breakdown"]["diesel_tco2e"]} tCO2e")
    print(f"  - Fleet Petrol:                {s1["breakdown"]["petrol_tco2e"]} tCO2e")
    print(f"  - Stationary Natural Gas:      {s1["breakdown"]["natural_gas_tco2e"]} tCO2e")
    print(f"Scope 2 Indirect (Location):     {s2["scope2_location_based_tco2e"]} tCO2e")
    print(f"Scope 2 Indirect (Market-based): {s2["scope2_market_based_tco2e"]} tCO2e")
    print("---------------------------------------------------------")
    print(f"Total Scope 1 + Scope 2:         {round(s1["total_scope1_tco2e"] + s2["scope2_location_based_tco2e"], 3)} tCO2e")
    print("=========================================================")

if __name__ == "__main__":
    main()
