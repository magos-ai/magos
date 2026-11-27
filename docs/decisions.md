# Decisions

The decisions taken while this repository was planned, on 2026-11-26
and 2026-11-27. Each one was put to the owner as a question with
options, one at a time.

This file holds them until `docs/PRODUCT.md` with its PDRs (#2) and
`docs/ARCHITECTURE.md` with its ADRs (#6) exist. Those issues move
each decision into its record; then this file goes. Until then, a new
decision is added here in the same way.

Each entry says what was decided, and where the plan carries it.

## The product

### What Magos is

- Magos is one platform of AI agents for every area of one person's
  life: work, family and home, personal projects, career. Each area is
  a workspace with its own agents, connectors, memory and data.
  Nothing goes from one workspace into another.
- Most agents do not write code. Development is one workspace among
  others: a session may run with no repository, no worktree, no
  commit and no pull request. Its work is data that it reads and
  writes through connectors, and the person sees it on dashboards, in
  chats and in notifications (#39).
- A to-do app is not needed beside Magos: its dashboards, chats and
  notifications cover that work.

### The order of the work

1. M1 · Tasks run through Magos: the development of Magos itself runs
   through Magos.app on one Mac.
2. M2 · Magos suggests: what to take next and what can run in
   parallel, without being asked.
3. M3 · Second computer and harnesses.
4. M4 · Apps as agents and skills: the owner's own tools run inside
   Magos as agents and skills, and one main chat routes each message.
5. M5 · Health.
6. M6 · Automations and accounts.
7. Later: what is wanted but not planned.

M1 is done when the owner can run in Magos the same flow that the owner runs
in a terminal today: from `#N` in a chat to a merged pull request,
several sessions at once, nothing lost after a reboot. Five real tasks
are measured against that flow (#17).

Milestones have no due dates: their numbers give the order.
Milestones M3 to M6 get their `Done when` when the one before closes.

### Extensions are standard

- There are two kinds of extension only: Agent Skills with harness
  plugins, and MCP servers. Magos has no format of its own for them.
- Magos runs no skill itself. It gives the skills and the MCP servers
  of a workspace to each harness session that it starts, and it knows
  no skill by name.
- A skill or a plugin comes from a link. Magos shows what is inside,
  and after a yes pins it at the commit it was taken from. A newer
  version upstream comes in only after the person sees its diff and
  says yes (#38).
- Tools and connectors are MCP servers. They are added at runtime,
  with no rebuild of the app, and a person can write them inside
  Magos.
- Magos's own contract is its versioned MCP tools: sessions, Servos, a
  question to the person, a dashboard card, a proposed task.

### Agents

- A built-in agent is a base prompt, skills and tools, inside one
  workspace. Each of its connectors has a level: reads, writes with
  consent, or writes by itself.
- Whether agents from outside are needed at all is open: tools written
  inside Magos and kept out of any public repository may cover them
  (#39).
- A person's own agents are connectors plus data. Only the connectors
  can be public; the prompts, the skills and the data stay private.
- Development is the work of an agent: the development agent of a
  workspace, or the agent of one project. The agent plans; harnesses
  are its hands, and each step of the plan runs in a harness session.
  The agent spreads the steps over the harnesses and models that it is
  given: a cheap model for a simple step, as a component or a test;
  the next step on another harness when the limit of one runs out; a
  review on another harness than the one that wrote the code. For a
  goal such as "close all bugs" or a new feature, it writes a plan,
  with the order and what runs where, and asks once before it runs the
  plan. Who picks the harness of each step, and in which milestone,
  is decided next.

### Above the workspaces

- Only a router stands above the workspaces. It sees the message and
  the agents' descriptions, never a workspace's data, and hands the
  message to the default agent of the workspace that fits. That agent
  answers, or hands the task to another agent of its own workspace.
  No agent sees across workspaces.
- Home, and a morning brief across all areas, are built from short
  summaries that the agent of each workspace publishes for Home. Each
  workspace decides what it publishes.
- The routing model is chosen by a comparison test of small models on
  labelled real messages: accuracy, latency and cost (M4).
- Search over threads and memory is a dialogue with the main agent,
  not a screen of its own (M4).

### Access

- Access is granted per resource, never to a whole app: one
  repository, one mailbox, one calendar. Each grant has a hint for the
  model on what it is. One workspace can hold two accounts of one
  service.
- A project lives inside a workspace. A grant goes to a workspace or
  to one of its projects. In M1 a resource is a GitHub repository or a
  GitHub organization, which covers all of its repositories, and the
  person connects them in the app (#8, #37). Other connectors join the
  same place later.

### Routines and suggestions

- A routine is data: when, which workspace, which harness, which
  skills, which prompt.
- A routine only proposes, as one card: "found 4 debt tasks, file
  them?". A tap starts a second session that writes only what was
  approved. A kind of action can be allowed to run by itself: one
  setting per routine and kind of action, which can be taken back.
- A routine that fills a dashboard or makes suggestions runs a cheap
  harness session with the workspace's skills, and puts its cards
  through Magos's MCP tools. What Magos knows by itself, as its
  sessions, their state and the questions that wait, it shows with no
  model.
- A routine starts at a time or on an event: a new issue, a merged
  pull request, the end of a session. A schedule view shows each
  routine with its next run, a switch and its last result. When
  nobody opened a routine's result three times in a row, the routine
  asks whether to turn it off (M2).
- Magos sees GitHub by polling every 1 to 2 minutes with conditional
  requests: a `304` costs no rate limit. No webhooks: a Mac has no
  public address, and Magos requires no relay of a third party.

### Dashboards and the chat

- Nothing on a dashboard is hard-coded: the workspace's agent
  assembles it from primitives. The view language is chosen by
  research (#20).
- The same language draws a card inside a chat answer, and a card can
  be pinned to the dashboard (#20).
- The workspace's agent sets the hints above the input field, as it
  builds the dashboard (M2).
- The person can react to an agent's message with an emoji. The
  reaction reaches the agent as feedback, and Magos keeps it (M4).
- Voice goes both ways: dictation, and a talk mode with spoken replies
  (M4).

### The thread in M1

- A question from a session shows as a card with its options, the
  Recommended mark and a free-text answer (#4, #15).
- The thread shows the session's full log, read-only: each tool call
  on one line, its output folded. The log stays on the Servo that ran
  it, because it can hold code and secrets; the core passes a page of
  it on request and keeps none (#9, #15).
- A message can carry images, such as a pasted screenshot (#15).
- A session can be copied by its id and opened in Terminal, as a way
  out when Magos hangs or a session goes wrong (#14).
- The checks and the review verdicts on the merge question come in M2.
- One macOS notification for each question that waits, and a badge on
  the Dock icon. Other events and sounds come later (#16, #36). The
  look may change when the UI is designed.

### Memory and security

- Agents need memory at three levels: one shared memory about the
  person, one per workspace, and maybe one per project. Its form is
  chosen by research (#34).
- The rules for personal data are set before personal data comes in,
  and are written as an ADR: at most two of untrusted input, private
  data and outside effects in one session; limits on reads, not only
  writes; a log of every tool call; a stop that ends the process
  (#35).
- An agent acts on a website on the person's behalf through a browser
  on a Servo. The sign-in comes through a password manager's channel
  for agents, so the model never sees the password, and Magos asks
  before an irreversible step (M6).

### Harness terms

- Each Servo starts the unmodified harness CLI under the login of the
  person who owns that computer. Magos never holds a harness login,
  and nobody starts another person's harness.
- The terms of every provider are read again before the first release
  for other people (#32).

### A commercial version

- Shared workspaces are the main business feature of a commercial
  version: a person invites other people, and people and agents work
  on projects together (#29).
- Each person's tasks run on that person's own Servo: a shared
  workspace offers a task, and the person's Servo takes it.
- The repository is the open core under Apache-2.0, as the README
  says. Functions for a team and a hosted cloud are separate add-ons.

### The interface

- The apps are native. Nothing of the earlier Electron app is ported:
  its code and its UI rules stay behind.
- The rules and the skills for the native interface are chosen when
  the UI work starts. The earlier prototypes of the interface are kept
  as references outside this repository.

## The architecture

These hold until the spike (#3) checks them; ADR 0001 records the
result (#6).

- One Go program, `magos`, with two roles: `magos serve`, the core,
  which keeps workspaces, projects, grants, threads and runs in
  SQLite; and `magos servo`, which runs harness sessions and works
  with files. Rust was weighed; Go was chosen.
- The Mac app is SwiftUI. It gives the macOS capabilities to the core
  as MCP tools, so the harness adapters are written once, in Go.
- `magos` ships inside `Magos.app` as a helper, registered as a
  LaunchAgent with `SMAppService`, so launchd restarts it. Scheduled
  work goes on when the app is closed. The core opens the app in the
  background when it needs a macOS tool.
- Harnesses are driven directly: Claude Code through `claude -p` with
  stream-json, Codex through `codex app-server`. Both are in M1. No
  ACP for harnesses, and no A2A.
- The core, the Servo and the app talk through a small versioned
  control protocol: a local socket first, pairing and HTTPS for a
  remote core in M3. Events are numbered and acknowledged. The Go and
  Swift sides share the message types, generated from one schema.
- In M1 the core and the Servo run as two processes with separate data
  folders, and a test proves that they share no files.
- After a reboot or a crash, each run resumes from SQLite by its
  harness session id.
- In M1 the wall between workspaces is in Magos's data only: a session
  sees all that the macOS user sees, and the app says so. The session
  sandbox comes in M3 (#24). Magos's own tokens are in the Keychain.
- M1 is signed with a free Apple Development certificate and built
  from source. Developer ID and notarization come when the app goes to
  other people (#33).
- The tests of the harness adapters run on sessions recorded on this
  repository with short neutral tasks. A script replaces the home path
  and the account ids before the commit, and the harness version is in
  the file name.

## The tracker

- GitHub's default labels, and the labels that craft makes when it
  first needs them, as `in progress` and `debt`. A research issue
  starts its title with "Research:". No area labels.
- The board is the organization's project "magos roadmap".
