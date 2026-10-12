# Model Primitives

The AI base: tokens, context, latency, cost, structured output and the choice between models.
Enough to reason about a feature before designing it, and to tell which constraint will bind first.

| # | Lesson | The question it answers | |
|---|---|---|---|
| 01 | [Tokens and the Context Window](01-tokens-and-the-context-window.md) | What is the budget you are spending, and what happens at the edges? | ✅ |
| 02 | [Cost Per Request](02-cost-per-request.md) | What is the arithmetic every AI feature needs and almost none contain? | ✅ |
| 03 | Latency That Users Feel | Time to first token, tokens per second, streaming, and why p95 is the number | |
| 04 | Structured Output and Schemas | How do you get something a program can rely on? | |
| 05 | Tool Calls, and What They Really Are | What happens when a model asks for something to be run? | |
| 06 | Choosing a Model | Frontier or small, hosted or open weights: what is the real tradeoff? | |
| 07 | Routing, Caching and Batching | What are the three cost levers, and what does each cost you in quality? | |
| 08 | The Non Determinism Problem | Temperature, seeds, reproducibility, and what it does to your test suite | |
| 09 | Multimodal Inputs | What changes when the input is an image, a document or audio? | |
| 10 | Rate Limits and Capacity | What happens at launch when everybody arrives at once? | |
| 11 | Reading a Model Card | What does a provider tell you, and what do you have to find out yourself? | |
| 12 | What Changes Every Six Months, and What Does Not | Which of your assumptions have an expiry date? | |
