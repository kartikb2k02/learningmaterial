# Agent Reliability Simulator

```text
Estimated difficulty: 3/5 | Estimated time: 90 min
```

Turn per step accuracy into the number you can actually promise a customer. A loop of ten steps at
95% each succeeds 60% of the time, and that single fact decides more agent designs than any model
choice does.

## Why it matters

Agent quality is usually discussed one step at a time, because that is the level at which it can be
observed. Users experience the whole task. The gap between those two views is multiplicative, it is
unintuitive, and it is the reason agents demo well and disappoint in production.

## What to build

1. **A task as a sequence of steps**, each with its own success probability. Compute end to end
   success as the product.
2. **Retries.** Allow a failed step to be retried a fixed number of times, and recompute. This is
   the cheapest reliability lever and the program should show how much it actually buys.
3. **A checkpoint.** Allow one step to be verified before continuing, catching some share of
   failures. Show the difference between verifying early and verifying late, which is large and in
   the direction most people do not expect.
4. **Cost and latency.** Each step and each retry costs tokens and seconds. Print expected cost and
   p95 latency alongside success, because the reliability fix is usually bought with both.
5. **A sweep over step count.** Print success against number of steps, so the shape of the decay is
   visible rather than asserted.

## What good looks like

* The program shows that cutting a task from twelve steps to six does more for reliability than
  raising per step accuracy by two points.
* Retries help, and the program shows where they stop helping.
* Verification placed early beats verification placed late, and the output demonstrates why.

## Stretch

Make failures correlated rather than independent, which is closer to reality: when a step goes
wrong, the context is now poisoned and later steps are likelier to fail too. Compare the
independent and correlated models and note which one your production traces actually resemble.
