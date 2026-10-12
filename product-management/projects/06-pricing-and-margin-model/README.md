# Pricing and Margin Model

```text
Estimated difficulty: 3/5 | Estimated time: 90 min
```

Find the usage level at which your pricing stops working. Put a seat price, a credit price and a
per task price against a cost model that scales with use, and print gross margin per customer
across the whole distribution of how much people actually use the thing.

## Why it matters

Software pricing evolved where the marginal cost of a user was zero, so a seat price was safe. With
a model in the product, your most engaged customers are your most expensive, which is the opposite
of the correlation every seat price assumes. The failure is slow: margin erodes toward your best
customers while product strategy pushes engagement up, and the two effects hide each other for
several quarters.

## What to build

1. **A usage distribution.** Not an average. Use something long tailed, because that is what usage
   is, and the average customer is not where the money or the cost lives.
2. **Three pricing models**: a flat seat price, a credit bundle with overage, and pure per task.
3. **A cost model** per customer that uses your cost per task from project 02.
4. **Margin per customer across the distribution**, printed as a table by usage decile. Find the
   decile where margin goes negative under each pricing model.
5. **The blended answer.** Print overall gross margin for each pricing model, and the share of
   customers who are unprofitable under each. These two numbers frequently point in opposite
   directions, which is the whole lesson.

## What good looks like

* You can name the usage percentile at which a seat price stops covering cost.
* Credits with overage look better on blended margin and worse on customer experience, and the
  program makes that tradeoff numeric rather than aesthetic.
* Changing cost per task by a third visibly moves which pricing model wins.

## Stretch

Add a fall in model prices over time, since token prices have fallen every year so far, and ask who
captures that: you, through margin, or the customer, through a price cut. Then add a cap on the
heaviest users and see how much of the margin problem a cap solves without changing the headline
price at all.
