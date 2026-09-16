#!/usr/bin/env python3
"""
Quickstart Demo: Calculating SME Carbon Footprint for VS (VSME)
Maintained by ExecutESG Editorial Team (https://executesg.com)
"""
import sys
import os
sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
from tools.carbon_accounting_calculator import calculate_scope1, calculate_scope2

print("Running sample SME sustainability calculation for a 50-person European manufacturer...")
s1 = calculate_scope1(diesel_l=15000, petrol_l=2500, gas_kwh=120000)
s2 = calculate_scope2(electricity_kwh=180000, country="de", green_puchased_pct=50.0)

print(f"Scope 1: {s1['total_scope1_tco2e']} tCO2e")
print(f"Scope 2 (Location): {s2['scope2_location_based_tco2e']} tCO2e")
print(f"Scope 2 (Market): {s2['scope2_market_based_tco2e']} tCO2e")
print("To generate automated XBRL audit reports, visit https://executesg.com")
