# Cohort Retention Table

```text
Estimated difficulty: 2/5 | Estimated time: 60 min
```

Build the one chart that tells the truth about a product: cohorts down the side, months across the
top, and the share of each cohort still present in each month. Then answer the only question that
matters about it, which is whether the curve flattens.

## Why it matters

Aggregate retention is an average over cohorts of different ages and is almost always wrong in a
flattering direction. A cohort table is harder to fake and it answers a question no single number
can: does some share of your users stay indefinitely, or does every cohort eventually reach zero?
The first is a business. The second is a leaking bucket that growth spending hides.

## What to build

1. **Synthetic events first.** Generate signups across twelve monthly cohorts and activity per user
   per month, driven by a survival curve you control. Getting the shape of the maths right on data
   you generated is much faster than debugging against a warehouse.
2. **The table.** Cohort by month, printed as percentages, with the cohort size in the first column
   so the reader can see which rows are too small to trust.
3. **The flattening test.** Fit or simply inspect whether the month over month decline approaches
   zero. Print the asymptote: the share of a cohort that appears to stay forever.
4. **The honest caveat.** Print the sample size next to every cell and suppress any cell below a
   threshold, because a 50% retention figure drawn from four users is noise with a percent sign.

## What good looks like

* Two generated products, one with a flattening curve and one without, produce visibly different
  tables and different asymptotes.
* Cells based on fewer than a stated number of users are blank rather than misleading.
* You can state your own product's asymptote, or say honestly that you do not have enough months of
  data to know, which is the most common correct answer.

## Stretch

Add revenue retention alongside logo retention and watch them diverge. Then split one cohort by
acquisition channel and see whether the curve differences are larger than the differences between
cohorts, which is usually the finding that changes where money goes.
