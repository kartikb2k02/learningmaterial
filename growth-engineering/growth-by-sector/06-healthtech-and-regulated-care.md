# Healthtech and Regulated Care

```text
Level: Sectors
Prerequisites: When PLG Is the Wrong Answer
Estimated difficulty: 4/5 | Estimated time: 35 min
```

## The idea in one minute

Here the limiting rate is not conversion, it is how many governance and procurement cycles complete
in a year. The worked process takes eighteen months across six stages and 12 percent of starts
finish, so with forty six processes in flight you complete 30.7 a year and win 3.5. Nothing in that
calculation is a conversion rate anybody can optimise in a quarter. The interventions that matter
shorten the calendar or raise the pass rate at the stages that reject: a pre built governance and
safety pack is worth more than two extra people, and it is a one off build rather than a permanent
cost.

For the three buyers, the clinical safety obligations and the evidence requirements, read
[Healthtech and Regulated Care](../../product-management/product-by-sector/06-healthtech-and-regulated-care.md).
This lesson is the growth delta only.

## A picture

```mermaid
flowchart LR
  C[":i-stethoscope: <b>A clinician asks</b><br/><small>62%</small>"] --> G[":i-scroll-text: <b>Governance</b><br/><small>3.5 months, 71%</small>"]
  G --> S[":i-shield-check: <b>Safety case</b><br/><small>2.5 months, 78%</small>"]
  S --> I[":i-lock: <b>Security</b><br/><small>2 months, 84%</small>"]
  I --> P[":i-stamp: <b>Procurement</b><br/><small>4 months, 69%</small>"]
  P --> B[":i-calendar: <b>Budget year</b><br/><small>5 months, 58%</small>"]
  B --> W[":i-banknote: <b>12% of starts, 18 months</b>"]
  class C neutral
  class G,S,I queue
  class P,B warn
  class W memory
```

Six gates, eighteen months, and the last two reject the most while being the least influenceable.

## How it really works

### Throughput, not conversion

Wins per year is processes in flight, divided by the cycle length in years, times the end to end
pass rate. In the worked case 46 divided by 1.5, times 12 percent, giving 3.5 wins and 396,447.

That expression has three inputs and none of them is a funnel conversion rate in the usual sense.
Two are operational capacity, one is the calendar, and all three are improvable on a two year
horizon rather than a quarterly one.

The planning consequence: a quarterly growth target in this sector is a measure of activity rather
than outcome, and the honest reporting unit is processes entered, processes advanced a stage, and
processes completed.

### Where to intervene, ranked

Pre built governance and safety documentation shortens the cycle from 18 to 13 months and raises the
pass rate, producing 5.8 wins and 649,733. Getting onto a purchasing framework gives 5.3 wins. Two
more people to run processes gives 5.0. A reference site in the same region gives 4.3.

Two observations. The documentation pack beats hiring, because it affects every process rather than
adding capacity for a few. And it is a one off build rather than a recurring cost, which is the shape
worth looking for when the calendar is the constraint.

All four together give 12.0 wins and 1,341,314, which is three and a half times the base, and that
is the realistic two year programme for a growth function in this sector.

### What the free tier is for, and what it is not

It is not for revenue. There will be none, for the reasons in
[When PLG Is the Wrong Answer](../product-led-growth/05-when-plg-is-the-wrong-answer.md): the buyer
is not the user, procurement is compulsory, and the buyer count is small.

It is for three things, each measurable. Producing a clinician champion, measured as sites with an
advocate. Shortening the safety case, because evidence of safe use in practice is the hardest part
of that submission, and measured in months saved. And supplying evidence for the business case,
measured as citations in submissions.

Two months off an eighteen month cycle raises completed cycles by 13 percent, which is 49,556 a
year of extra revenue from a product nobody paid for. State the free tier's purpose that way and it
survives a budget review; state it as a conversion funnel and it does not.

### What it changes across this course

| Topic | The delta here |
|---|---|
| [Acquisition Channels](../acquisition-channels/README.md) | One channel: clinical networks and professional bodies. Advertising is constrained by regulation rather than by cost |
| [Content and Search](../content-and-search/README.md) | Clinically reviewed content with named authors, which is a compliance requirement and now also a moat |
| [Activation and Onboarding](../activation-and-onboarding/README.md) | Access is granted by an administrator rather than chosen, so coverage is set by rollout schedules |
| [Lifecycle Messaging](../lifecycle-messaging/README.md) | Clinical messages take absolute priority and sit outside any growth budget. Document the boundary |
| [The Experiment Machine](../the-experiment-machine/README.md) | Experiments on clinical pathways need governance approval, so velocity is capped by a committee and correctly so |
| [Guardrails](../the-experiment-machine/04-guardrails-and-automatic-stops.md) | Clinical safety signals must stop automatically and cannot be traded against conversion |
| [Growth Economics](../growth-economics/README.md) | Payback periods of two to three years are normal, so the usual limits do not apply and the constraint is procurement timing |
| [Personalisation](../ai-in-growth/03-personalisation-that-is-not-creepy.md) | Inferring a health condition for marketing is the clearest prohibition in the course |

### The budget year stage is the one nobody plans for

Five months and a 58 percent pass rate, the worst in the pipeline, and it is purely a calendar
artefact: a decision made in the wrong month waits for the next financial year or dies.

Two practical responses. Know every target site's budget calendar and work backwards, so a process
starts in time to land in a cycle. And have a smaller entry point that fits inside discretionary
spend, which avoids the stage entirely for a first deployment and converts it into an expansion
conversation later.

### Reference sites are the compounding asset

A reference site in the same region raises the pass rate by 22 percent in the worked model, and the
effect compounds: each win makes the next one in that region easier, which is the only loop
available here.

So the first three customers in a region are worth far more than their contract value and should be
priced accordingly, which is the one place in this sector where
[Buying Revenue at a Loss](../growth-economics/06-buying-revenue-at-a-loss.md) genuinely applies.

## Run the numbers

Compute throughput, then price the interventions.

```go
// healthtech.go: growth when the calendar is the funnel.
package main

import "fmt"

// Here the limiting rate is not conversion, it is how many procurement and
// governance cycles complete in a year. Every growth technique that shortens
// a funnel is irrelevant; the ones that matter raise the number of cycles you
// can be in and the share of them you win.
type stage struct {
	name   string
	months float64
	pass   float64 // share that get through
}

const (
	sites    = 1_900.0 // addressable sites
	acv      = 112_000.0
	inFlight = 46.0 // processes you can be in at once, given your team
)

func main() {
	pipeline := []stage{
		{"a clinician asks for it", 1.0, 0.62},
		{"an information governance review", 3.5, 0.71},
		{"a clinical safety case", 2.5, 0.78},
		{"information security review", 2.0, 0.84},
		{"procurement and framework check", 4.0, 0.69},
		{"budget in the right financial year", 5.0, 0.58},
	}

	fmt.Printf("%-40s %10s %10s %14s\n", "stage", "months", "pass", "cumulative")
	months, cum := 0.0, 1.0
	for _, s := range pipeline {
		months += s.months
		cum *= s.pass
		fmt.Printf("%-40s %10.1f %9.0f%% %13.0f%%\n", s.name, s.months, 100*s.pass, 100*cum)
	}
	fmt.Printf("\nthe whole process takes %.1f months and %.0f%% of starts finish\n",
		months, 100*cum)

	cycles := inFlight / (months / 12)
	fmt.Printf("with %.0f processes in flight, %.1f complete a year\n", inFlight, cycles)
	fmt.Printf("of which %.1f are wins, worth %.0f\n\n", cycles*cum, cycles*cum*acv)

	fmt.Println("Nothing in that calculation is a conversion rate you can optimise")
	fmt.Println("in a quarter. The output is set by how many processes you can carry")
	fmt.Println("at once and how long each one takes, and both are operational")
	fmt.Println("capacity questions rather than funnel questions.")

	fmt.Println("\nwhat each possible improvement is worth:")
	fmt.Printf("%-44s %14s %14s %14s\n", "change", "months", "wins/year", "revenue")
	base := cycles * cum
	for _, c := range []struct {
		name string
		mf   float64 // multiplier on total months
		pf   float64 // multiplier on pass rate
		cf   float64 // multiplier on processes in flight
	}{
		{"nothing", 1, 1, 1},
		{"pre built governance and safety pack", 0.72, 1.18, 1},
		{"get onto a purchasing framework", 0.84, 1.26, 1},
		{"two more people to run processes", 1, 1, 1.4},
		{"a reference site in the same region", 1, 1.22, 1},
		{"all four together", 0.60, 1.45, 1.4},
	} {
		m := months * c.mf
		p := cum * c.pf
		if p > 1 {
			p = 1
		}
		cy := inFlight * c.cf / (m / 12)
		fmt.Printf("%-44s %14.1f %14.1f %14.0f\n", c.name, m, cy*p, cy*p*acv)
	}
	_ = base

	fmt.Println("\nThe governance and safety pack is worth more than two extra")
	fmt.Println("people, because it shortens every process and raises the pass rate")
	fmt.Println("at the two stages that reject most often. It is also a one off")
	fmt.Println("build rather than a permanent cost, which is the shape to look for")
	fmt.Println("in a sector where the calendar is the constraint.")

	fmt.Println("\nwhat the free tier is actually for:")
	fmt.Printf("%-34s %-24s %s\n", "purpose", "measured on", "sensible target")
	for _, r := range []struct{ a, b, c string }{
		{"revenue", "conversions", "none; there will be none"},
		{"producing a clinician champion", "sites with an advocate", "a third of the pipeline"},
		{"shortening the safety case", "months saved", "two to four"},
		{"evidence for the business case", "cited in a submission", "most wins"},
	} {
		fmt.Printf("%-34s %-24s %s\n", r.a, r.b, r.c)
	}
	fmt.Println("\nA free tier here converts nobody directly and shortens the")
	fmt.Println("process, which is worth more. Two months off an eighteen month cycle")
	fmt.Printf("raises completed cycles a year by 13%%, which is %.0f of extra\n",
		cycles*cum*acv*0.125)
	fmt.Println("revenue a year from a product nobody paid for.")
}
```

The pipeline table is the exercise. Stage durations and pass rates are knowable from your own
history and from asking the sites that said no at which stage they stopped. The cumulative column
is the number to show anybody who thinks this is a conversion problem.

The intervention table ranks the responses: 3.5, 5.8, 5.3, 5.0, 4.3 and 12.0 wins a year. Put your
own stage durations in; the ranking between documentation and hiring is robust across plausible
inputs.

The free tier table is for the budget conversation. Replace the purpose column's first row with a
zero and keep the other three.

## The AI layer

Three applications, and the boundary matters more here than anywhere else in the course.

Documentation and submission preparation is the largest and safest win. Governance submissions,
clinical safety cases, security questionnaires and data protection assessments are long documents
assembled from material you already hold, repeated per site with local variations. That is the
intervention worth 5.8 wins against 3.5, and it is a mapping task rather than a judgement one.

Evidence assembly is the second: collating usage and outcome data from existing sites into the form
a submission requires. It shortens the safety case stage, which is the free tier's main purpose
above.

The boundary: nothing that infers, suggests or decides anything clinical sits inside a growth
programme. Not a personalised nudge about a condition, not a predicted risk score used for
targeting, not a message whose content depends on inferred health status. Those are clinical
decisions under clinical governance, and a growth team that builds one has created a regulatory
problem rather than a growth loop. The line is not subtle and it is worth stating in writing before
anybody proposes crossing it.

## Remember this

* The limiting rate is completed governance and procurement cycles, not conversion.
* Wins a year is processes in flight, over the cycle length in years, times the end to end pass
  rate: 3.5 in the worked case.
* A quarterly growth target here measures activity. Report processes entered, advanced and
  completed.
* A pre built governance and safety pack beats two extra people, because it affects every process
  and is a one off build.
* All four interventions together take wins from 3.5 to 12.0 a year, which is the realistic two year
  programme.
* The free tier is for champions, for shortening the safety case and for evidence, not for revenue.
* Two months off an eighteen month cycle is worth 49,556 a year from a free product.
* The budget year stage is five months with the worst pass rate and is pure calendar. Know each
  site's budget cycle.
* A smaller entry point inside discretionary spend avoids that stage entirely for a first
  deployment.
* Reference sites compound within a region, which makes the first three customers there worth more
  than their contracts.
* Nothing clinical belongs in a growth programme: not inference, not targeting, not nudges.

## Try it

1. Build your pipeline table with real stage durations and pass rates, including the reasons sites
   stopped. Compute wins a year from throughput rather than from a funnel.
2. Cost a pre built governance and safety pack against two more people and present both with the
   wins a year figure attached.
3. Restate your free tier's purpose as champions, months saved and submissions cited, with targets
   for each, and remove the conversion target.

## Check yourself

1. Your sales team reports a 40 percent win rate and three deals a year. What is the growth
   constraint, and what would you change?
2. Why does a documentation pack beat hiring two more people here?
3. Somebody proposes a lifecycle campaign targeted at patients with a particular inferred condition.
   What do you say?
