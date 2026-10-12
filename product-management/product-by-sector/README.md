# Product by Sector

The same function, bent by the market it runs in. Seven sectors, each one a full product treatment
rather than a list of industry facts: who actually buys, what the cycle does to the loop, what the
product must contain that no user ever asked for, and what that does to the roadmap.

**Read [Product Archetypes](../what-a-model-changes/04-ai-product-shapes.md) first.** The shape of a product
usually determines more than the market it sells into. A fintech marketplace behaves like a
marketplace that happens to move money, and borrowing the fintech playbook for it will mislead you
about liquidity, which is the thing that will actually kill it.

| # | Lesson | The question it answers |
|---|---|---|
| 01 | [Enterprise B2B SaaS](01-enterprise-b2b-saas.md) | Why does the buyer want things the user will never see, and what does that cost? |
| 02 | [Fintech](02-fintech.md) | What changes when being wrong moves money and the regulator has a veto? |
| 03 | [Ecommerce and D2C](03-ecommerce-and-d2c.md) | Why is this a margin business wearing a funnel costume? |
| 04 | [Marketplaces](04-marketplaces.md) | How do you build for two customers whose interests oppose each other? |
| 05 | [Devtools and Infrastructure](05-devtools-and-infrastructure.md) | What happens when your users can read the source and will tell you why you are wrong? |
| 06 | [Healthtech and Regulated Care](06-healthtech-and-regulated-care.md) | How do you build around a clinician when the buyer, the payer and the patient are three people? |
| 07 | [Vertical SaaS](07-vertical-saas.md) | How do you win a small market completely, and when does services gravity eat you? |

## The comparison matrix

The five structural forces, across all seven. This table is the reason the sector block exists in
one place: every other course states its deltas against these rows and never restates them.

| | Who signs | Cycle | Regulation | Data constraint | Margin shape | Abuse surface |
|---|---|---|---|---|---|---|
| **Enterprise B2B SaaS** | A committee, with procurement and security | Two to twelve months | Contractual rather than statutory, plus certifications | Tenancy, residency, retention, audit | High gross margin, very high cost of sale | Low, but credential and permission misuse matters |
| **Fintech** | Varies, but risk and compliance hold a veto | Weeks to a year | Statutory, with licensing and supervision | Residency, retention, lineage, right to explanation | Spread or fee based, thin and volume driven | Very high: fraud, laundering, account takeover |
| **Ecommerce and D2C** | One person, in minutes | Minutes | Consumer law, tax, payments | Mostly volume and freshness | Thin contribution margin per order | Payment fraud, returns abuse, review manipulation |
| **Marketplaces** | Both sides separately | Minutes to weeks | Varies by what is traded; often employment or tax | Identity and trust signals on both sides | Take rate on gross volume | Very high: disintermediation, fake supply, scams |
| **Devtools and Infrastructure** | A developer first, finance later | Days to adopt, months to pay | Light, until you hold customer data | Logs, secrets, telemetry consent | High margin, except where you resell compute | Quota abuse, crypto mining, supply chain |
| **Healthtech and Regulated Care** | An institution, rarely the user | Six to twenty four months | Heavy and clinical, with device regimes | Strict privacy, consent, residency, retention | Varies wildly by reimbursement | Identity, prescription and claims fraud |
| **Vertical SaaS** | An owner or operator, often non technical | Weeks to months | Whatever the industry carries | Industry specific records and retention | High margin, dragged down by services | Low, but the data is often the business of record |

Treat these as orientation, not benchmarks. They are the typical shapes seen as of **12 October
2026**, and they are deliberately ranges. Real figures, with sources and dates, live in
[Benchmarks and Their Limits](../benchmarks-and-their-limits/README.md).

## How to read a sector you have never worked in

Six questions, in this order. The order matters, because each answer constrains the next.

1. **Who signs, and who uses?** If they are different people, almost everything else follows from
   that single fact.
2. **What is the cycle, end to end?** This sets your loop length, which sets how fast you can learn.
3. **What is forbidden?** Regulation and contract define the envelope before you design anything.
4. **What does the data layer permit?** Residency, retention and consent rules kill features early
   and cheaply if you know them, expensively if you do not.
5. **Where does the margin come from?** A thin margin business cannot afford the same product
   decisions as a high margin one, regardless of how good the feature is.
6. **How will it be abused?** Every sector has a dominant abuse vector, and the first version of
   your product will meet it.
