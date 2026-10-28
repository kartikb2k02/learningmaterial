# Project 2: A tracking plan for an app you use

**Produces** a one page grammar, twenty specified events, an identity design and a monitoring plan.
**Takes** three days. **Assumes**
[Instrumentation and the Event Model](../instrumentation-and-the-event-model/README.md).

## The brief

Pick an app you use daily and that you did not build. Write the tracking plan its growth team should
have, from the outside, by using the product carefully and reasoning about what must exist behind
it.

Doing this from the outside is the exercise. You cannot look up the answer, so you have to decide
what matters, which is the decision a real tracking plan makes and usually skips.

## What it has to contain

1. **The grammar on one page**: the object list, the action list, the property vocabulary and the
   allowed values for each property.
2. **Twenty events**, each with: the name under your grammar, what user action fires it, its
   properties with types and allowed values, client or server, and the question it exists to answer.
   The last field is the one that makes the plan maintainable.
3. **The identity design**: how anonymous, user and account relate, what joins a device to a person,
   and where the account identifier lives.
4. **Five questions** the plan answers, written as questions, with the events each one needs. If a
   question needs more than four events, simplify the question or the plan.
5. **The monitoring plan**: which alarm on which event, and the threshold.
6. **One deliberate overlap**: an event collected on both sides, with the ratio you would alarm on.

## The hard parts

You will want more than twenty events. The constraint is the exercise: twenty events that answer
five questions is a working plan, and two hundred events that answer none is what most products
have.

You will also find that the product's own behaviour is ambiguous in places, which is the ambiguity a
real plan has to resolve by decision rather than by observation. Write the decision down.

## How to know you are finished

Hand somebody your five questions and your twenty events and ask them to write the query for one
question using only your event list. If they have to ask you what an event means, the plan has the
ambiguity failure rather than a naming problem.

## Where to look when stuck

[Designing an Event Taxonomy](../instrumentation-and-the-event-model/02-designing-an-event-taxonomy.md)
for the grammar, [Identity](../instrumentation-and-the-event-model/03-identity.md) for item 3,
[Client, Server or Both](../instrumentation-and-the-event-model/05-client-server-or-both.md) for
items 2 and 6, and
[Data Quality and the Slow Rot](../instrumentation-and-the-event-model/06-data-quality-and-the-slow-rot.md)
for item 5.
