# Trust Budget Calculator

```text
Estimated difficulty: 1/5 | Estimated time: 45 min
```

Turn "is the model good enough" into arithmetic. Given a product shape, the volume, and what one
mistake costs the person who absorbs it, compute the accuracy the feature actually requires, and
compare it against the accuracy you measure.

## Why it matters

The question of whether a model is good enough has no answer in the abstract and a precise answer
for a specific feature. The same model is comfortably sufficient for a suggestion surface and
nowhere near sufficient for an autonomous worker, and the difference is more than an order of
magnitude in required accuracy. Most arguments about model quality are really this calculation
being skipped.

## What to build

1. **A shape definition**: outputs per day, the share of wrong outputs a human catches before any
   damage, the minutes lost when one is not caught, and the loaded cost of that person's time.
2. **Damage per month** at your measured accuracy.
3. **Required accuracy**, by inverting that against an error budget you are willing to carry. Clamp
   it: when mistakes are cheap enough, the answer is that there is no accuracy constraint at all,
   and printing a negative percentage is a bug rather than an insight.
4. **A verdict** per shape: fits, out of reach, or unconstrained.
5. **Sensitivity.** Print what happens when the catch rate falls, which is what approval fatigue
   does to an approved action flow within about a month of launch.

## What good looks like

* Four shapes with the same model produce four different verdicts.
* Lowering the catch rate from 0.85 to 0.3 visibly moves a feature from viable to not.
* You can name the three changes that fix an out of reach feature, none of which is a better prompt:
  narrow the task, add a review step, or route the uncertain cases to a person.

## Stretch

Add a second error class. Most products have cheap mistakes and expensive ones in different
proportions, and an average error rate hides exactly the thing that matters. Model them separately
and see how often the expensive tail, not the overall rate, is the binding constraint.
