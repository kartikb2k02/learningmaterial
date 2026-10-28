# Glossary

The growth terms this course uses, defined the way it uses them. Where a word is contested in the
wider industry, the entry says so rather than pretending there is a settled meaning. The right hand
column names the topic that treats the term properly.

**Metric terms are not defined here.** Acquisition cost, payback, lifetime value, cohorts, windows,
rates and denominators belong to
[How to Read a Metric](../product-management/how-to-read-a-metric/README.md) and
[Revenue and Efficiency](../product-management/revenue-and-efficiency/README.md), which define them
once. This page covers only the terms that are specific to growth work.

## The function and the loops

| Term | What it means here | Topic |
|---|---|---|
| **Growth charter** | The written statement of which surfaces a growth team may change without asking and what guardrails protect them. Two questions answered once rather than renegotiated per change | The Growth Function |
| **Loop** | A cycle whose output becomes its own input, so each turn makes the next turn larger. Distinguished from a funnel, which leaks and has to be refilled from outside | Loops and Compounding |
| **Loop speed** | The time for one full turn, from input to reinvested output. Usually matters more than loop strength, because speed compounds and strength does not | Loops and Compounding |
| **Loop strength** | The multiple each turn produces. A strength above one with a slow cycle loses to a strength below one with a fast cycle over any realistic horizon | Loops and Compounding |
| **Reinvestment step** | The point where a loop's output is converted back into its input. The step most often missing from a diagram somebody calls a loop | Loops and Compounding |
| **Input metric** | A number that feeds the growth equation and that somebody can change this quarter. Distinguished from an output metric, which is the thing being predicted | The Growth Model |
| **Growth debt** | The residue of shipped experiments: live flags, losing variants, running messages, unread events, unowned surfaces, untrusted dashboards. Compounds, because it slows the work that would clear it | Growth Debt |

## Channels

| Term | What it means here | Topic |
|---|---|---|
| **Channel fit** | Whether a channel is available to a product at all, decided by price, frequency and purchase trigger rather than by execution quality | Channel Fit |
| **Fair test** | A channel test sized against the gap between the channel's plausible cost and your bar, rather than against what felt like a reasonable budget | Testing a Channel Properly |
| **Saturation** | The point at which a channel's marginal cost exceeds what you can pay, as distinct from a channel being badly run. The average cost hides it | Saturation and the Ceiling |
| **Marginal cost per customer** | What the next customer from a channel costs, which is the average cost divided by the diminishing returns exponent. The only number that answers where the next pound goes | The Marginal Pound |
| **Lifetime ceiling** | The most a channel's customers can ever repay: price times margin over churn. A channel whose ceiling is below its cost cannot pay back at any horizon | Payback by Channel |
| **Money page** | A page serving a query with purchase intent, as opposed to one serving a query that flatters a traffic chart | Intent and the Money Pages |
| **Programmatic pages** | Pages generated from a data source against a query template. A method, not a strategy, and a liability once the data is thin | Programmatic Pages |

## Activation and the product surfaces

| Term | What it means here | Topic |
|---|---|---|
| **Activation** | One action and one window, chosen for separation and coverage together. A definition you adopt rather than a stage you discover | What Activation Means Here |
| **Separation** | The retention gap in percentage points between users who took an action and users who did not. How much the action appears to matter | What Activation Means Here |
| **Coverage** | The share of a signup cohort who take an action inside the window. How much room there is to move | What Activation Means Here |
| **Aha moment** | The action that causes retention, as opposed to the action that predicts it. The two produce charts a single percentage point apart | Finding the Aha Moment |
| **Movable share** | One minus unprompted coverage. The population an intervention can actually affect, and the ceiling on what the project is worth | Finding the Aha Moment |
| **Drop out per minute** | Abandonment divided by a step's duration. The correct way to rank onboarding steps, and it inverts the ranking by duration | Time to Value |
| **Breakeven pass rate** | The share who must complete an added step for it to pay: baseline conversion over conditional conversion. Lower when judged on retained users than on activation | Friction Worth Keeping |
| **Retention multiplier** | How much better users who cleared a harder bar retain. The input that lowers the breakeven, and the one teams most often guess | Friction Worth Keeping |
| **First artefact** | Something containing the user's own material, created in the first session. The only measure of an empty state | Empty States and First Run |
| **Natural frequency** | How often the job the product does actually arises, which sets a ceiling on engagement that no notification can exceed | Natural Frequency |
| **Capture rate** | The share of occasions when the job arose on which the product was used. Comparable across any frequency, unlike daily over monthly actives | Natural Frequency |
| **Dependence ratio** | The share of returns with a send in front of them. Falling over a lifetime is a habit forming; flat and high means the sends are the habit | Triggers |
| **Reachability** | A stock, not a flow: the share of a cohort you can still contact. Spent by sends and reduced permanently by opt outs | Habits Without Dark Patterns |
| **Lead time** | Days between a churn signal firing and the departure. Under the time an intervention needs, the signal is observing churn rather than predicting it | Spotting Users About to Leave |
| **Restick rate** | The share of reactivated users still present a month later. The discount on every revival, and usually about a third | Winning Back the Dormant |

## Messaging

| Term | What it means here | Topic |
|---|---|---|
| **Message map** | One row per user per week with a count in it, built from send logs. Not a grid of segments and stages, which hides the only thing that matters | The Message Map |
| **Frequency cap** | A limit on messages per user per week, across channels rather than per channel. The optimum rather than a compromise, because the marginal send is negative | What Not to Send |
| **Materiality test** | The rule that a send goes out only when something happened that this user would want to know about. The brake that cheap generation removes | What Not to Send |
| **Inbox placement** | The share of sent mail reaching an inbox, decided at the sending domain level by how recipients have treated your mail. Unengaged recipients tax engaged ones | Email That Is Not Ignored |
| **Permission half life** | Weeks until half of granted push permissions are revoked or muted. Nine weeks for an early prompt, up to a year for one shown after a real event | Push and Notifications |
| **Disruption cost** | The value of the task a message interrupted. The third term in every in product message, and the one no placement report contains | In Product Messaging |
| **Attributed conversion** | A conversion credited to a message by the funnel report. Includes everybody who would have converted anyway, which was half of them in the worked case | Measuring a Campaign Honestly |
| **Incremental conversion** | The difference between a treated group and a random holdout. The only measure of what a message was worth | Measuring a Campaign Honestly |

## Monetisation

| Term | What it means here | Topic |
|---|---|---|
| **Reverse trial** | Full features for a period, then a downgrade to a free tier rather than a wall. Reaches value at trial rates and caps exposure at freemium rates | Trial, Freemium or Reverse Trial |
| **Crossover cost** | The marginal cost to serve a free user at which one packaging choice stops beating another. Fifty three pence a month for freemium in the worked case | Trial, Freemium or Reverse Trial |
| **Usage wall** | A limit reached through the user's own activity, so it arrives as a consequence of success. Converts at nearly twice a hard wall and keeps three quarters of the rest reachable | Paywalls |
| **Upgrade moment** | A state the user is in, not a place in the product. Conversion at a prompt is set by what just happened to them | The Upgrade Moment |
| **Prompt fatigue** | The decline in every prompt's value as the number a user sees rises. The reason the best set of upgrade moments is smaller than the set that individually work | The Upgrade Moment |
| **Invite rate** | Invitations sent per seat per month through the ordinary operation of the product. The term in expansion that compounds, and usually nobody's job | Expansion as a Loop |
| **Fair save** | A cancel flow response that addresses the stated reason. Worth 2.7 times an obstructive flow and twice a discount over twelve months | Cancel Flows and Saves |

## The machinery

| Term | What it means here | Topic |
|---|---|---|
| **Event taxonomy** | A grammar of object and action in past tense, with everything that varies in a property. Keeps a catalogue at 67 events rather than 360 | Designing an Event Taxonomy |
| **Tracking plan** | A file in the repository stating, per event, what fires it, its properties and allowed values, where it is collected, who owns it, and which question it answers | The Tracking Plan as a Contract |
| **Back joining** | Attributing a device's prior anonymous events to a person once it becomes identified. The alias call most implementations make once and then forget | Identity |
| **Assignment** | The deterministic, uniform, independent and stable map from a unit to a variant. Four properties, and every word costs something when it is missing | What an Experimentation Platform Does |
| **Exposure** | Evidence that a unit reached the thing that differed, logged in both arms at the point of divergence. Not the same population as assignment | Assignment and Exposure |
| **Dilution** | The shrinkage of a measured effect by the exposed share, when an experiment is analysed on assignment. Deterministic, not added noise | Assignment and Exposure |
| **Sample ratio mismatch** | Arm sizes departing from the intended split by more than chance allows. Evidence the arms are not comparable, so discard rather than caveat | Assignment and Exposure |
| **Peeking** | Checking a result repeatedly and stopping at the first significant reading. Turns a 5 percent false alarm rate into 28 percent after thirty looks | Guardrails and Automatic Stops |
| **Guardrail** | A metric that stops an experiment without a human deciding. Operational rather than statistical, so it should be twitchy | Guardrails and Automatic Stops |
| **Experiment velocity** | Experiments concluded with a decision recorded, per quarter. Interchangeable with win rate in the arithmetic and the only one of the two with room to move | Experiment Velocity |
| **Decision capacity** | Decisions a team can make per week. A hard constraint on throughput once building gets cheap, and set by the agenda | The Weekly Growth Meeting |
| **Surface collision** | Two concurrent experiments on the same screen. A birthday problem, and orthogonal assignment removes the bias in the averages but not the interaction | Working With Product and Engineering |

## Motion and economics

| Term | What it means here | Topic |
|---|---|---|
| **Sales floor** | The smallest annual contract that can repay a salesperson's time inside your payback limit. About 900 at eleven hours a close | What PLG Actually Requires |
| **Routing rule** | The comparison, per signup, between self serve profit and sales profit minus the cost of a touch. Decides the motion and sizes the team | Self Serve, Sales Assist, Hybrid |
| **Product qualified lead** | A signal that a touch will pay, which requires a measured lift from the touch rather than a correlation with conversion | Product Qualified Leads |
| **Crossing rate** | The chance an adopting team seeds another team in the same company per quarter. Decides an account's ceiling almost entirely | Bottom Up Inside a Company |
| **Signups per buyer** | Signups a plan requires divided by companies that could ever buy. Over 0.3 means most of the market must sign up; over 1.0 is impossible | When PLG Is the Wrong Answer |
| **Enterprise ceiling** | The revenue above what a self serve motion can reach, gated by invoicing, single sign on, a security pack and contracts. Seventy seven percent of the market in the worked case | The Enterprise Ceiling |
| **Cash trough** | The maximum you are down before accumulated cohorts exceed monthly spend. Grows faster than linearly in the spend rate and is the number finance cares about | Setting and Defending a Budget |
| **Stopping rule** | A pre agreed condition that reverts spend without a meeting, triggered on a number finance can see. The line that gets a budget approved | Setting and Defending a Budget |

## Models in the loop

| Term | What it means here | Topic |
|---|---|---|
| **Breakeven use count** | Price minus fixed cost to serve, divided by cost per use. The number that sets the free tier, the usage cap and whether engagement work helps | Why AI Products Grow Differently |
| **Correction rate** | How often users change what a model inferred or generated. A rate near zero is suspicious rather than good | Time to Value |
| **Retry rate** | The share of users who try again after a bad first output. Worth more activation than a model upgrade and free in unit terms | Activation When the First Output Is Bad |
| **Depth of use** | How much a user has built on current model behaviour. Protective against every kind of churn except a change in the thing they built on | Retention When the Model Changes |
| **Demo loop** | Sharing an output as an acquisition channel. The only loop in this course that decays on its own, because novelty feeds two of its three terms | The Demo Loop |
| **Paired metric** | A later metric that would move the wrong way if a loop were cheating its proxy. The defence against the quiet failures, which are 80 percent of the cost | How They Fail |
| **Spam equilibrium** | The state a channel reaches when sending becomes free: volume rises, reply rates collapse, and every sender's profit is unchanged | AI Outbound and the Spam Equilibrium |
