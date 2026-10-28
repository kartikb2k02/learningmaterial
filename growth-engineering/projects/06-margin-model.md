# Project 6: A margin model for a product with a model in it

**Produces** a breakeven use count, a cohort gross profit curve, and a packaging recommendation.
**Takes** two days. **Assumes**
[Growth for AI Products](../growth-for-ai-products/README.md).

## The brief

Pick a product that calls a model and charges a price you can see. Work out whether its cohorts are
profitable over three years, and what its packaging should be.

Published pricing plus published model prices plus careful use of the product gets you close enough.
The point is the shape of the answer, not its precision.

## What it has to contain

1. **The cost per use**, built from the model's published prices and your estimate of the work one
   use requires. State the estimate and its basis.
2. **The breakeven use count**: price minus fixed cost to serve, over cost per use.
3. **Where the included allowance sits** relative to that count. If it is above, say what that means.
4. **A cohort run forward three years**, with your estimate of churn and of usage growth among
   survivors, producing monthly gross profit per cohort.
5. **The month the cohort crosses into loss**, or a statement that it does not and why.
6. **The free tier priced**: segments, cost per free user, conversion probability, and net per
   segment.
7. **A packaging recommendation** with the allowance and the overage rate, and the migration plan
   for existing customers.
8. **Two technical levers costed**: what routing and caching would buy in months of headroom.

## The hard parts

Item 1 requires a judgement about how much work one use involves, and you will be wrong. Do it as a
range and carry the range through, because the conclusion is usually robust to a factor of two and
saying so is more useful than a single confident number.

Item 4's usage growth rate is the input that decides everything and is the one nobody measures. Your
estimate should be defensible from how the product works: a tool used for a recurring task grows
differently from one used for a project.

Item 7 is where the project becomes a recommendation rather than an analysis. An allowance above the
breakeven count is not a metered plan, so if your recommendation includes one, check item 3 again.

## How to know you are finished

You should be able to state, in one sentence, whether this product's growth is constrained by
acquisition or by margin. If the answer is margin, your packaging recommendation should fix it
permanently rather than buying time, and you should be able to say which.

## Where to look when stuck

[Why AI Products Grow Differently](../growth-for-ai-products/01-why-ai-products-grow-differently.md)
for items 1 to 5,
[Free Tiers When Every Use Costs Money](../growth-for-ai-products/02-free-tiers-when-every-use-costs-money.md)
for item 6,
[Usage Growth That Destroys Margin](../growth-for-ai-products/06-usage-growth-that-destroys-margin.md)
for items 7 and 8, and
[What These Loops Cost to Run](../ai-in-growth/06-what-these-loops-cost-to-run.md) for the cost
stack in item 1.
