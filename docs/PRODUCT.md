# Magos

The product's direction, its constraints, its open questions and the
index of its decisions. Each decision is a PDR in [`pdr/`](pdr/). The
decisions of the code are in [`decisions.md`](decisions.md) until
`ARCHITECTURE.md` and its ADRs exist (#6).

## What Magos is, and for whom

- Magos is an open-source platform of AI agents for every area of one
  person's life: work, family and home, personal projects, career.
  Each area is a workspace with its own agents, connectors, memory and
  data (PDR 0003).
- The person talks to Magos in a chat or by voice, in native apps for
  the Mac and the iPhone (PDR 0015).
- The first user is the owner, who builds Magos with Magos. Then
  people who run it on their own computers, on their own
  subscriptions. A commercial version for teams comes later (PDR
  0016).
- Most agents do not write code. Development is one kind of work among
  others: a session may run with no repository, no worktree, no commit
  and no pull request. Its work is data that it reads and writes
  through connectors, and the person sees it on dashboards, in chats
  and in notifications (#39).
- The value is in specialised agents and skills that work with the
  person's own data, not in a general chat.

## Direction and the current goal

The milestones go in order, M1 to M6, then Later (PDR 0002). The
current goal is M1 · Tasks run through Magos: the development of Magos
itself runs through Magos.app on one Mac, from `#N` in a chat to a
merged pull request, three sessions at once, nothing lost after a
reboot.

## How Magos differs

- Simple on top: a chat, connectors, and an agent that writes its own
  skills.
- No lock to one vendor: any harness and any model, on the person's
  own subscriptions (PDR 0004).
- A way to go deep: the full log of a session, a session opened in
  Terminal, skills as files (PDR 0005, PDR 0018).
- For development, the agent picks the harness, the model and the
  review of each step, with a fallback when a limit runs out and the
  cost of each step (PDR 0006). A market research of about 425
  products on 2026-11-16 found no other product that does this, and a
  real task run in Hermes on 2026-11-28 did not show it either (#45).

The closest products, checked again on 2026-11-28:

- Development:
  - [Multica](https://github.com/multica-ai/multica): a board where
    the person gives issues to coding agents on their own computers.
    The person picks the agent of each issue.
  - [bb](https://getbb.app/): a local-first IDE for coding agents. It
    comes nearest, with Paseo Hub: when the person asks, its
    orchestrating agent starts threads and picks the provider of each,
    and by default a review runs on another agent than the builder.
    Its docs name no choice of the model, no fallback when a limit
    runs out, and no cost of each thread.
  - [Paseo Hub](https://paseo.sh): starts agents on the person's
    computers from GitHub, Linear, Slack and schedules. Its
    orchestrating agent picks the provider and the model of each
    worker, guided by the person's notes. Its docs name no fallback
    when a limit runs out, no rule to review on another provider, and
    no cost of each worker.
  - [Orca](https://github.com/stablyai/orca): runs CLI agents in
    parallel worktrees on the person's own subscriptions; a
    coordinating agent splits a task and starts workers.
- Personal agents:
  - [OpenClaw](https://github.com/openclaw/openclaw): an assistant on
    the person's own devices, through chat channels and native apps.
  - [Hermes](https://github.com/NousResearch/hermes-agent): an
    open-source personal agent by Nous Research, reached through
    messengers. A real task ran in it in #45.
  - [Grok Bot](https://x.ai/blog/introducing-grok-bot): a closed agent
    in the cloud of xAI, on xAI models only, that works in apps and
    sites.

## What Magos will not do

- Be a to-do app beside its agents: dashboards, chats and
  notifications cover that work.
- Have a format of its own for extensions (PDR 0005), or run a skill
  itself (PDR 0004).
- Hold a harness login, or start another person's harness (PDR 0004).
- Let an agent see across workspaces (PDR 0003, PDR 0008). In M1 the
  wall is in Magos's data only, and the app says so (PDR 0003).
- Require a relay of a third party (PDR 0010).
- Ship apps for Windows, for Linux on a desktop or for Android, for
  now (PDR 0015, #42).
- Keep the code of a closed add-on in this repository (PDR 0001).
- Take code from outside before the rules of #44 (PDR 0021).

## Constraints

### The law

- Everything in this repository is under the Apache License 2.0 (PDR
  0001).
- Each Servo runs the unmodified harness CLI under the login of the
  person who owns the computer. The terms of every provider are read
  again, and shown to a lawyer, before the first release for other
  people (Q-1, #32).
- The rules for personal data are set before personal data comes in
  (PDR 0020, #35).

### The client

- The owner is the first client. M1 is measured against the flow that
  the owner runs in a terminal today (#17).

### The money

- On the person's own computers, sessions run on the person's own
  subscriptions, and later on their own API key too (PDR 0004, #47).
- On servers of a hosted cloud, agents run on API keys or on the
  person's own Servo. Magos keeps no subscription login there (#29).
- The open core is free for one person. Functions for a team and for a
  hosted cloud are closed add-ons (PDR 0001); their prices are open
  (Q-6).

## Open questions

### Q-1 Do the terms of each harness provider allow how Magos runs its CLI?

- Answers: lawyer
- Blocks: the first release for other people, #32, PDR 0004
- Since: 2026-11-27

### Q-2 Which model provider may the router use, as it reads a message before its workspace is known?

- Answers: person, after the research of #35
- Blocks: #35, PDR 0008, PDR 0020
- Since: 2026-11-27

### Q-3 Are agents from outside needed at all?

- Answers: person, after the research of #39
- Blocks: #39, PDR 0007
- Since: 2026-11-27

### Q-4 Which form does the memory of agents take?

- Answers: person, after the research of #34
- Blocks: #34, PDR 0019
- Since: 2026-11-27

### Q-5 Which view language draws the dashboards and the cards?

- Answers: person, after the research of #20
- Blocks: #20, PDR 0011, PDR 0014
- Since: 2026-11-27

### Q-6 What do the add-ons for a team and a hosted cloud cost?

- Answers: person
- Blocks: #29, PDR 0001
- Since: 2026-11-27

### Q-7 Which kinds of artefacts come, and in which order?

- Answers: person, after the research of #50
- Blocks: #50, PDR 0011
- Since: 2026-11-27

## Decisions

- [0001 Apache-2.0 open core, with team and cloud as closed add-ons](pdr/0001-licence.md) — accepted; revisit when: the first add-on needs a part that one person needs too (#29), or a person from outside offers code for a team (#44)
- [0002 The order of the milestones, M1 to M6, then Later](pdr/0002-milestones.md) — accepted; revisit when: a milestone closes, or a check of PDR 0006 shows a product that does all that it names
- [0003 A workspace for each area of life, and nothing crosses between them](pdr/0003-workspaces.md) — accepted; revisit when: a task needs the data of two workspaces at once, and the summaries for Home do not cover it
- [0004 Magos drives unmodified harnesses on the person's own logins](pdr/0004-harnesses.md) — accepted; revisit when: a provider's terms forbid it (#32), or `claude -p` stops working on a subscription (#46)
- [0005 Standard extensions only, added at runtime](pdr/0005-extensions.md) — accepted; revisit when: Codex loads a plugin for one session (#5), or a harness takes neither Agent Skills nor MCP servers
- [0006 The development agent plans the steps, and Magos runs them](pdr/0006-development.md) — accepted; revisit when: another product does all of it
- [0007 An agent is a prompt, skills and tools inside one workspace](pdr/0007-agents.md) — accepted; revisit when: #39 recommends
- [0008 Only the router, Home and search stand above the workspaces](pdr/0008-above-workspaces.md) — accepted; revisit when: #35 is settled, or the test of routing models runs (M4)
- [0009 Access is granted per resource](pdr/0009-access.md) — accepted; revisit when: a service cannot grant one resource only
- [0010 A routine proposes, and a tap applies](pdr/0010-routines.md) — accepted; revisit when: most kinds of action run by themselves, or a core with a public address comes (M3)
- [0011 The workspace's agent assembles its dashboard from primitives](pdr/0011-dashboards.md) — accepted; revisit when: #20 chooses the view language
- [0012 Read, listen and talk on every document, question and task](pdr/0012-voice.md) — accepted; revisit when: #49 moves from Later into M1 to M6
- [0013 Settings are a chat](pdr/0013-settings.md) — accepted; revisit when: the first place of settings is drawn (#37)
- [0014 A feature reaches every client at once](pdr/0014-client-releases.md) — accepted; revisit when: a client cannot draw the view language (#20)
- [0015 Native apps for Apple devices only](pdr/0015-clients.md) — accepted; revisit when: people from outside ask for another platform, or a commercial version comes
- [0016 Shared workspaces lead the commercial version](pdr/0016-commercial-version.md) — accepted; revisit when: the add-ons are planned (#29)
- [0017 The Constructor creates agents and workspaces through chat](pdr/0017-constructor.md) — accepted; revisit when: M6 gets its `Done when`
- [0018 What the thread shows in M1](pdr/0018-thread.md) — accepted; revisit when: the UI of the thread is designed
- [0019 Memory at three levels](pdr/0019-memory.md) — accepted; revisit when: #34 recommends
- [0020 Personal data comes in only after its rules](pdr/0020-personal-data.md) — accepted; revisit when: #35 is settled, or M5 starts
- [0021 Built in public, with no code from outside yet](pdr/0021-outside-contributions.md) — accepted; revisit when: the first person from outside offers code, or the first release for other people is planned (#44)
- [0022 One download, with no setup at the first start](pdr/0022-install.md) — accepted; revisit when: the app goes to other people (#33)
