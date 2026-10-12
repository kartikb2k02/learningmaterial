# Cost Per Task Model

```text
Estimated difficulty: 2/5 | Estimated time: 60 min
```

Compute what one unit of useful work actually costs, including the parts nobody counts: retries,
the judge calls, the escalations to a human, and the context you are sending every single time
without meaning to.

## Why it matters

Teams quote the price of one model call and call it the cost of the feature. The real figure is
usually two to five times higher, and the gap is entirely made of things that feel like
infrastructure rather than product: a retry policy, a safety check, a second model grading the
first, and a human cleaning up what neither caught.

## What to build

1. **A task definition**: input tokens, output tokens, and the model price per million of each.
2. **A retry model.** Some share of calls fail or produce unusable output and run again. Multiply.
3. **A quality tax.** If you run a judge on some percentage of traffic, that is another call with
   its own token count. If a share escalates to a person, that is loaded hourly cost times minutes.
4. **Context accounting.** Separate the fixed part of the prompt from the per request part. Print
   what share of your spend is the system prompt you send on every call, because it is often most
   of it and it is the easiest thing to cut.
5. **A sweep.** Print cost per task across a range of volumes, so the monthly number appears next to
   the unit number. Both are needed and only one is usually quoted.

## What good looks like

* The program prints cost per task broken into model, retries, judging and human, as four numbers.
* Changing the system prompt length by a thousand tokens visibly changes the answer.
* You can state, in one sentence, what one task costs and which component dominates.

## Stretch

Add a routing strategy: a cheap model attempts the task first and a frontier model handles the
cases it fails. Model the quality cost of that honestly, by assuming the cheap model's failures are
not all detected, and find the escalation threshold where total cost is lowest without quality
falling below your bar.
