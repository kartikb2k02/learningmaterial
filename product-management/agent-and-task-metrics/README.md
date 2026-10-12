# Agent and Task Metrics

Measuring something that takes many steps on its own. Per step accuracy is what you can observe
and whole task success is what users experience, and the gap between those two is multiplicative,
unintuitive, and the reason agents demo well and disappoint in production.

| # | Lesson | The question it answers |
|---|---|---|
| 01 | Steps Per Task and Loop Depth | What does a rising step count tell you before quality drops? |
| 02 | Reliability as a Compound Number | How do you turn per step rates into a number you can promise? |
| 03 | Tool Error and Recovery Rate | How often do actions fail, and does the agent recover? |
| 04 | Autonomy Ratio | What share of actions happen without a human, and should it rise? |
| 05 | Time to Task Completion | Why is this the metric users actually feel? |
| 06 | Cost and Latency Percentiles | Why does the mean hide the thing that makes people leave? |
| 07 | Progress Without Completion | How do you measure an agent that is working but not finishing? |
| 08 | Correlated Failure | Why do real agents fail worse than independent probabilities predict? |
| 09 | Measuring Blast Radius | How do you quantify what a wrong action could have reached? |
| 10 | Human Time Saved, Honestly | What is the credible version of this claim? |
| 11 | Instrumenting an Agent | What has to be logged for any of this to be measurable? |
| 12 | Worked: An Agent Scorecard | One real agent, measured end to end |
