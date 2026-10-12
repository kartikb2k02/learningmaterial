# Projects

Six things to build. Each one is a small Go program, standard library only, that produces a number
you can take into a real argument at work. They are ordered so that each builds on the last, and
every one of them exists because the corresponding lesson is easier to believe once you have run
the arithmetic on your own product.

None of them need a GPU, an API key, a network connection or an account. Expect an hour or two
each.

| # | Project | What you end up with | Topic it belongs to |
|---|---|---|---|
| 01 | [Eval Harness](01-eval-harness/README.md) | A pass rate with a confidence interval over a real case set | Evaluation and Quality |
| 02 | [Cost Per Task Model](02-cost-per-task-model/README.md) | What one unit of useful work costs, including retries and judging | AI Cost and Margin |
| 03 | [Trust Budget Calculator](03-trust-budget-calculator/README.md) | The accuracy your product shape actually requires | What a Model Changes |
| 04 | [Cohort Retention Table](04-cohort-retention-table/README.md) | A retention curve, and whether it flattens | Funnel, Retention and Experiments |
| 05 | [Agent Reliability Simulator](05-agent-reliability-simulator/README.md) | Per step accuracy turned into a number you can promise | Agent and Task Metrics |
| 06 | [Pricing and Margin Model](06-pricing-and-margin-model/README.md) | The usage level at which your pricing stops working | AI Pricing and Packaging |

Build them with your own numbers, not the ones in the briefs. The defaults are there so the program
runs on the first try; they are not benchmarks.
