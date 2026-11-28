# 0006 The development agent plans the steps, and Magos runs them

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

- Development is the work of an agent: the development agent of a
  workspace, or the agent of one project. The agent plans; harnesses
  are its hands, and each step of the plan runs in a harness session.
- The agent spreads the steps over the harnesses and models that it is
  given: a cheap model for a simple step, as a component or a test;
  the next step on another harness when the limit of one runs out; a
  review on another harness than the one that wrote the code.
- For a goal such as "close all bugs" or a new feature, it writes a
  plan, with the order and what runs where, and Magos asks once before
  it runs the plan.
- The agent picks the harness and the model of each step as it writes
  the plan, and names a fallback for it. A Magos MCP tool tells it
  which harnesses and models there are, how much of each limit is
  left, and what each costs.
- Magos runs the steps in their order, and the steps of one task go on
  in the same branch and worktree. When the limit of a step's harness
  runs out, Magos takes the fallback from the plan; with no fallback,
  it asks the agent. Magos counts the cost of each step, and gives the
  counts to the agent for its next plan.
- The plan comes in two milestones. In M2: the plan, its one "Run?"
  and the run of its steps on Claude Code and Codex, with the review
  on the other harness, the fallback when a session reports that a
  limit ran out, and the cost of each step (#41). In M3: GLM and the
  other harnesses, and the limits left read in advance, so the cheap
  steps go to GLM (#22, #23). Empryo, a harness with many providers
  and cheap or local models, is checked there too (#22).

## Rests on

- No other product runs development with an agent that chooses the
  harness, the model and the review of each step: a market research of
  about 425 products on 2026-11-16 (#2), and a real task run in Hermes
  (#45). See *Checked*.

## Revisit when

A product does all of it: a plan, the harness and the model of each
step chosen by an agent, a fallback when a limit runs out, and the
review on another harness. The list of the closest products is in
`PRODUCT.md`, *How Magos differs*. A real task in bb and in Paseo Hub
checks it before #41 starts (#58).

## Links

PDR 0002, PDR 0004, #22, #23, #41, #45, #56, #58

## Checked

- 2026-11-28 #45: a real task ran in Hermes, an open-source agent by
  Nous Research. Hermes did as well as the craft flow in a terminal: 2
  turns of the person and 1 h 16 min. But the flow of that run came
  from craft: the task was in the repository of craft, and Hermes read
  its rules from the files. Every step ran on one model. Hermes's board
  of tasks for agents (Kanban) can give a task its own model, but it
  took no part in the run. The check was one task; a repository
  without craft would show Hermes's own flow. The claim stands, and M1
  stays as planned (#56).
- 2026-11-28 #2: the closest products were checked again. Two come
  nearest, and each does a part. In Paseo Hub, an orchestrating agent
  picks the provider and the model of each worker, guided by notes
  that the person writes (https://paseo.sh/docs/agent-profiles.md).
  Its docs name no fallback when a limit runs out, no rule to review
  on another provider, and no cost of each worker. In bb, when the
  person asks, an orchestrating agent starts threads and picks the
  provider of each, and by default a review runs on another agent than
  the builder
  (https://getbb.app/guides/orchestrate-coding-agents). Its docs name
  no choice of the model, no fallback when a limit runs out, and no
  cost of each thread. A public set of Claude Code rules
  (https://github.com/rwspatin/agent-orchestration) maps kinds of task
  to a harness and a model and reviews across vendors, as rules that
  one session follows, with no fallback. No product found does all
  that the decision names.
