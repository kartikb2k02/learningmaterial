# Instrumentation and the Event Model

The machinery every other topic in this course assumes. What mediocre events actually cost, how to
name things so the catalogue stays small, the three identities and the two that are usually
conflated, how to make the tracking plan a gate rather than a document, which events belong on
which side, how a clean implementation decays, and what you are allowed to keep.

| # | Lesson | The question it answers |
|---|---|---|
| 01 | [Why Growth Work Dies Without Instrumentation](01-why-growth-work-dies-without-instrumentation.md) | What does a team lose when the events are wrong, and why does a 3% error become 14%? |
| 02 | [Designing an Event Taxonomy](02-designing-an-event-taxonomy.md) | What are the naming rules, and what does the first week cost you three years later? |
| 03 | [Identity](03-identity.md) | Anonymous, user and account: how do you stitch them without lying? |
| 04 | [The Tracking Plan as a Contract](04-the-tracking-plan-as-a-contract.md) | Who signs it, and what does a defect cost at each stage it could be caught? |
| 05 | [Client, Server or Both](05-client-server-or-both.md) | Which events go where, and why is the overlap worth paying for twice? |
| 06 | [Data Quality and the Slow Rot](06-data-quality-and-the-slow-rot.md) | How does a clean implementation decay, and which checks actually hold it? |
| 07 | [Consent and What You May Collect](07-consent-and-what-you-may-collect.md) | What may you keep, for how long, and how biased is the part you can see? |

Rates, denominators and cohort windows are defined in
[How to Read a Metric](../../product-management/how-to-read-a-metric/README.md). The experiment
machinery these events feed is built in
[The Experiment Machine](../the-experiment-machine/README.md).
