# Project 4: A channel test design somebody would fund

**Produces** a one page test plan with a decision rule, a budget and a named stop condition.
**Takes** one day. **Assumes**
[Acquisition Channels](../acquisition-channels/README.md) and
[Growth Economics](../growth-economics/README.md).

## The brief

Pick a real product and a channel it does not currently use. Design the test that would decide
whether to adopt it, written so that a finance team would approve it without a second meeting.

## What it has to contain

1. **The bar**: the maximum acquisition cost this product can pay, derived rather than asserted,
   with the price, margin and churn it comes from.
2. **The channel's plausible cost range**, and your reasoning. A range, not a point.
3. **The test budget**, computed from the gap between the bar and the plausible cost. If the gap is
   small, say so and recommend deciding another way, which is a valid output for this project.
4. **The three part decision rule**: continue above what, stop below what, and what happens in
   between. The third clause is the one that matters.
5. **The four things held constant** for the duration, named specifically for this product.
6. **The conversion definition**, stated once, with the distinction between signups and customers
   made explicit.
7. **The stop condition**: what would end the test early, triggered on a number the funder can see.
8. **What you will do with a pass**: the marginal cost measurement you would run before scaling.

## The hard parts

Item 3 is the exercise. Most people size a test by what feels affordable, and the arithmetic
usually says either that the test is cheap and should be run immediately, or that it costs more than
the answer is worth. Both are useful answers and the second one is the surprising one.

Item 4's third clause is the other hard part. Write what happens on an ambiguous result before you
know the result, because the whole point of the rule is that ambiguity is the likely outcome.

## How to know you are finished

Hand it to somebody and ask whether they could approve or refuse it without asking you a question.
If they need to ask what happens if the result lands in the middle, item 4 is incomplete.

## Where to look when stuck

[Channel Fit](../acquisition-channels/02-channel-fit.md) for item 1,
[Testing a Channel Properly](../acquisition-channels/03-testing-a-channel-properly.md) for items 3
to 6, [Setting and Defending a Budget](../growth-economics/04-setting-and-defending-a-budget.md) for
item 7, and [The Marginal Pound](../growth-economics/03-the-marginal-pound.md) for item 8.
