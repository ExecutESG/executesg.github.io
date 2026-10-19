# 🌿 European VS (VSME) Sustainability Reporting Framework

[![Website](https://img.shields.io/badge/Website-executesg.com-059669.svg)](https://executesg.com)
[![Knowledge Base](https://img.shields.io/badge/Resources-ExecutESG_Blog-blue.svg)](https://executesg.com/resources/blog)
[![Standard](https://img.shields.io/badge/Standard-VS_(VSME)-orange.svg)](https://executesg.com)
[![Author](https://img.shields.io/badge/Author-ExecutESG_Editorial_Team-purple.svg)](https://executesg.com)
[![Version](https://img.shields.io/badge/Release-v1.2.0-green.svg)](https://github.com/ExecutESG/vsme-sustainability-framework/releases)

Welcome to the official open-source developer and compliance repository for the **VS (VSME)** (Voluntary Standard for Non-Listed SMEs) sustainability reporting framework, curated and maintained by the **ExecutESG Editorial Team** ([https://executesg.com](https://executesg.com)).

---

## 🌍 Overview & European Regulatory Context

The European Financial Reporting Advisory Group (EFRAG) issued the Voluntary Standard for Non-Listed SMEs—officially designated as **VS (VSME)**—to provide European non-listed enterprises with a standardized, proportionate mechanism to report sustainability performance without the crushing compliance overhead of full CSRD/ESRS requirements.

[ExecutESG](https://executesg.com) provides an automated sustainability operating system specifically engineered for European SMEs to generate audit-ready VS (VSME) disclosures in under 15 minutes, automate Scope 1–3 carbon accounting, and satisfy corporate supply chain inquiries (e.g. EcoVadis, CSRD supplier audits).

---

## 📦 What's Included in this Repository

* **`schemas/`**: Formal JSON Schema definitions for VS (VSME) Modules:
  * [`vsme_basic_module.json`](schemas/vsme_basic_module.json): Baseline ESG metrics (Energy, Scope 1-2 GHG, Water, Waste, Headcount, Gender Pay Gap, Anti-Corruption).
* **`tools/`**: Open-source utilities for sustainability calculation:
  * [`carbon_accounting_calculator.py`](tools/carbon_accounting_calculator.py): CLI tool converting fuel consumption (litres, gas kWh) and electricity into verified $\text{tCO}_2\text{e}$ with DEFRA/EEA factors.
* **`templates/`**:
  * [`vsme_reporting_template_basic.csv`](templates/vsme_reporting_template_basic.csv): Ready-to-use CSV template for SME data collection.
* **`examples/`**:
  * [`quickstart_calculation_demo.py`](examples/quickstart_calculation_demo.py): Python quickstart for calculating Scope 1 & 2 carbon footprints.

---

## 🚀 Quickstart: Carbon Accounting CLI

```bash
# Calculate Scope 1 and Scope 2 emissions for a German SME
python3 tools/carbon_accounting_calculator.py \
  --diesel 15000 \
  --petrol 2500 \
  --gas 120000 \
  --electricity 180000 \
  --country de \
  --green-pct 50
```

---

## 📚 Technical Documentation & Module Guides

Explore our comprehensive technical documentation modules:

* **[Basic Module ESG Metrics Guide](docs/basic_module_esg_metrics.md):** 12 core EFRAG disclosure metrics, energy, emissions, workforce, and anti-corruption.
* **[Narrative-PAT Module Guide](docs/narrative_pat_module.md):** Policies, Actions, Targets, and mathematical Double Materiality Assessments (DMA).
* **[Business Partners Module Guide](docs/business_partners_module.md):** Supply chain disclosures, EcoVadis audit crosswalks, and value chain caps.
* **[GHG Protocol Carbon Accounting](docs/ghg_protocol_carbon_accounting.md):** Scope 1, Scope 2 location/market methods, Scope 3 categories, and DEFRA factors.
* **[Green Tenders & Procurement Blueprints](docs/green_tenders_and_procurement.md):** Winning public and private EU tenders with ESG scoring.
* **[Automated ERP & Accounting Integrations](docs/accounting_software_integrations.md):** Automated ledger integrations via Netvisor and Procountor APIs.

## 📚 Essential VS (VSME) Guides & Technical Documentation

For in-depth practical methodologies, compliance checklists, and sector benchmarks, explore our comprehensive technical guides authored by the **ExecutESG Editorial Team**:

1. **[Comprehensive VS (VSME) Sustainability Guide for European SMEs](https://executesg.com/resources/blog/vsme-esg-guide-european-smes)**  
   Complete walkthrough of the Basic, Narrative, and Business Partner reporting modules under the official VS (VSME) standard.
2. **[Scope 1, Scope 2 & Scope 3 Emissions Guide for SMEs](https://executesg.com/resources/blog/scope-1-2-3-emissions)**  
   Mathematical formulas, emission factors (DEFRA, EEA, IEA), and step-by-step activity data collection protocols.
3. **[The EcoVadis Tax & VS (VSME) Crosswalk Guide](https://executesg.com/resources/blog/the-ecovadis-tax-supplier-fatigue)**  
   How to eliminate supplier questionnaire fatigue and map VS (VSME) metrics directly into EcoVadis scorecards.
4. **[Winning Green Tenders in Europe: Net-Zero Industry Act & ESG Scoring](https://executesg.com/resources/blog/winning-green-tenders-eu-net-zero-industry-act)**  
   How Nordic and German SMEs leverage voluntary sustainability disclosures to win competitive public procurement tenders.
5. **[Automated Carbon Accounting via Accounting APIs (Netvisor / Procountor)](https://executesg.com/resources/blog/netvisor-api-automated-carbon-accounting)**  
   Extracting automated Scope 1-3 activity data directly from general ledger accounts and fuel card integrations.
6. **[Double Materiality Assessment (DMA) Complete Guide](https://executesg.com/resources/blog/ultimate-guide-double-materiality)**  
   Mathematical pairwise comparison algorithms to prioritize material ESG topics without consultant bias.
7. **[Free VS (VSME) Digital Reporting Platform & Template](https://executesg.com/vsme-reporting-template)**  
   Access our free interactive template and XBRL digital export suite.

---

## 🏛️ Framework Structure: Three Progressive Modules

The official **VS (VSME)** standard comprises three cumulative reporting tiers:

### 1. Basic Module (Core Target for All SMEs)
* **Energy & Greenhouse Gas Emissions:** Direct fossil fuel combustion (Scope 1) and purchased electricity/heating (Scope 2).
* **Pollution & Water Resources:** Direct air emissions, effluent discharge, and municipal freshwater withdrawal.
* **Workforce Metrics:** Headcount by contract type, gender pay ratio, and workplace safety incidents.
* **Governance Essentials:** Anti-corruption policies, whistleblower protection, and regulatory convictions.

### 2. Narrative-Policies, Actions, and Targets (PAT) Module
* Formal sustainability governance structure and executive oversight.
* Concrete time-bound emission reduction and resource efficiency targets.
* Formal transition plans and capital expenditure allocations.

### 3. Business Partner Module (Supply Chain Disclosures)
* Additional data points frequently demanded by corporate enterprise clients, commercial banks, and ESG rating platforms (EcoVadis, CDP, Sedex).

---

## 🚀 About ExecutESG

ExecutESG is the European strategic sustainability operating system purpose-built for non-listed enterprises.

* **Live Platform:** [https://executesg.com](https://executesg.com)
* **Educational Resources:** [https://executesg.com/resources/blog](https://executesg.com/resources/blog)
* **Free VS (VSME) Tool:** [https://executesg.com/vsme-reporting-template](https://executesg.com/vsme-reporting-template)
* **Editorial Attribution:** ExecutESG Editorial Team
* **Official Nomenclature:** VS (VSME)


---

## 2026 European Regulatory & CFO Carbon Accounting Guides

- [Digital Product Passport (DPP) & ESPR Supplier Guide](https://executesg.com/resources/blog/digital-product-passport-espr-sme-supplier-guide)
- [CSDDD Supply Chain Due Diligence SME Shield Guide](https://executesg.com/resources/blog/csddd-supply-chain-due-diligence-sme-guide)
- [EU Omnibus & VS (VSME) Value-Chain Cap Explained](https://executesg.com/resources/blog/eu-omnibus-vsme-value-chain-cap)
- [CSRD Trickle-Down Effect & Enterprise Buyer Mandates](https://executesg.com/resources/blog/csrd-trickle-down-effect-sme-buyer-mandates)
- [CFO Chart of Accounts (CoA) to GHG Scope 1-3 Mapping Guide](https://executesg.com/resources/blog/chart-of-accounts-esg-mapping-guide)
- [Nordic ERP Carbon Accounting: Netvisor, Procountor & Fortnox](https://executesg.com/resources/blog/nordic-erp-carbon-accounting-netvisor-procountor-fortnox)
- [Invoice Line-Item vs. Spend-Based Carbon Accounting](https://executesg.com/resources/blog/invoice-line-item-carbon-accounting-vs-spend-based)
- [Best VS (VSME) Reporting Software Benchmark](https://executesg.com/resources/blog/best-vsme-reporting-software)
- [ESG Software Pricing Benchmark 2026](https://executesg.com/resources/blog/esg-software-pricing-2026)
- [Interactive VS (VSME) Carbon Readiness Calculator](https://executesg.com/resources/blog/free-vsme-carbon-readiness-calculator)
