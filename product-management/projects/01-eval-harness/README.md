# Eval Harness

```text
Estimated difficulty: 2/5 | Estimated time: 90 min
```

Build the thing that replaces your test suite. A directory of cases, three kinds of grader, several
runs per case, and a pass rate reported with a confidence interval, so that a change to a prompt
can be judged rather than argued about.

## Why this one first

Every other project here measures something. This one measures whether the product works at all,
and it is the single artifact that separates teams who can change a prompt safely from teams who
cannot. It is also much smaller than people expect: the useful version is under two hundred lines.

## What to build

1. **A case format.** Each case needs an identifier, an input, whatever context the system would
   have, and the expectation. Start with JSON on disk, one file per case, because a case set that
   lives in a database stops being reviewable.
2. **Thirty real cases**, not five thousand synthetic ones. Cover typical requests in the
   proportions you actually see, edge cases (empty, enormous, wrong language, ambiguous), and
   adversarial ones. Add every production failure you can remember.
3. **Three graders.** A code grader for things code can decide: schema validity, required
   substrings, forbidden substrings, latency, cost. A stub judge grader that takes a rubric and
   returns a score, so the interface exists before you wire a model to it. A human grader that
   writes unlabelled cases to a file for review.
4. **Repeat runs.** Each case runs several times, because a case that passes two runs in three is
   not a passing case.
5. **A reported interval.** A Wilson interval on the pass rate. Without it, a move from 80% to 84%
   on fifty cases looks like progress and is two cases of luck.
6. **A comparison mode.** Run two configurations over the same cases and print which cases flipped
   in each direction, not only the totals. The flips are where the learning is.

## What good looks like

* Running it twice on an unchanged system gives the same verdict, even though individual cases vary.
* Adding a case takes under a minute, so production failures actually get added.
* The output fits on one screen and names the failing cases.
* The interval is wide enough to be uncomfortable at thirty cases, which is the honest signal that
  you need more of them.

## Stretch

Add a held out portion that nothing is tuned against. Add tagging so you can report a pass rate per
category and see that the average was hiding one bad segment. Then wire the judge grader to a real
model and check its agreement with your human labels on a sample before you trust a single number
it produces.
