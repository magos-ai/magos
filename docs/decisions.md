# Decisions

The decisions taken while this repository was planned, on 2026-11-26
and 2026-11-27. Each one was put to the owner as a question with
options, one at a time.

This file holds them until `docs/PRODUCT.md` with its PDRs (#2) and
`docs/ARCHITECTURE.md` with its ADRs (#6) exist. Those issues move
each decision into its record, the rules of the tracker go into
`CLAUDE.md` (#7), and then this file goes. Until then, a new decision
is added here in the same way.

Each entry says what was decided. When an issue or a milestone carries
it, the entry names it.

"craft" in this file is a set of Agent Skills for the flow of
development, from taking an issue to a merged pull request, that the
owner runs in a terminal today.

## The product

### What Magos is

- Magos is one platform of AI agents for every area of one person's
  life: work, family and home, personal projects, career. Each area is
  a workspace with its own agents, connectors, memory and data.
  Nothing goes from one workspace into another.
- Most agents do not write code. Development is one kind of work
  among others: a session may run with no repository, no worktree, no
  commit and no pull request. Its work is data that it reads and
  writes through connectors, and the person sees it on dashboards, in
  chats and in notifications (#39).
- A to-do app is not needed beside Magos: its dashboards, chats and
  notifications cover that work.
- The value is in specialised agents and skills that work with the
  person's own data, not in a general chat (#2).
- Settings are a chat: a pinned summary first, then cards with live
  controls, for agents, workspaces and connectors alike (#2).
- A feature reaches every client at once, so no client waits for the
  slowest one (#2, #20).
- The Constructor is a separate helper that creates and edits agents
  and workspaces through chat (#2, M6).
- How Magos differs: simple on top, with a chat, connectors and an
  agent that writes its own skills; no lock to one vendor, with any
  harness and model on the person's own subscriptions; and a way to
  go deep, with the full log, Terminal and skills as files. The
  closest products are named in #2, and checked again when
  `docs/PRODUCT.md` is written.
- Before the first code, one real task runs in Hermes and is compared
  with the craft flow: M1 rests on the claim that no other product
  runs development this way (#45).

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
three sessions at once, nothing lost after a reboot. Five real tasks
pass the measurement against that flow (#17). The full list is the
`Done when` of the milestone M1.

Milestones have no due dates: their numbers give the order.
Milestones M3 to M6 get their `Done when` when the one before closes.

### Extensions are standard

- There are two kinds of extension only: Agent Skills with harness
  plugins, and MCP servers. Magos has no format of its own for them.
- Magos runs no skill itself. It gives the skills and the MCP servers
  of a workspace, and of the project when the session has one, to each
  harness session that it starts, and it knows no skill by name.
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
  with the order and what runs where, and Magos asks once before it
  runs the plan.
- The agent picks the harness and the model of each step as it writes
  the plan, and names a fallback for it. A Magos MCP tool tells it
  which harnesses and models there are, how much of each limit is
  left, and what each costs. Magos runs the steps in their order, and
  the steps of one task go on in the same branch and worktree. When
  the limit of a step's harness runs out, Magos takes the fallback
  from the plan; with no fallback, it asks the agent. Magos counts the
  cost of each step, and gives the counts to the agent for its next
  plan.
- The plan comes in two milestones. In M2, the plan, its one "Run?"
  and the run of its steps on Claude Code and Codex, with the review
  on the other harness, the fallback when a session reports that a
  limit ran out, and the cost of each step (#41). In M3, GLM and the
  other harnesses, and the limits left read in advance, so the cheap
  steps go to GLM (#22, #23). Empryo, a harness with many providers
  and cheap or local models, is checked there too (#22).
- A running session can send a message to another through Magos. The
  message is data, not an instruction, and a limit stops a loop (#48,
  Later).
- An agent outside Magos, as a harness with plain craft on another
  computer, can work through Magos's MCP tools. What it may read and
  write is settled in the research of agents (#39).

### Above the workspaces

- Only a router stands above the workspaces. It sees the message and
  the agents' descriptions, never a workspace's data, and hands the
  message to the default agent of the workspace that fits. When the
  router is not sure which workspace fits, it asks the person. That
  agent answers, or hands the task to another agent of its own
  workspace. No agent sees across workspaces.
- The router is the main agent of the main chat. It reads no data of a
  workspace itself, and it acts only through the agents of the
  workspaces. A link to a film goes to the agent of the personal
  workspace, which adds it to the list; a question about a work project
  goes to the agent of the work workspace; the answer shows in the main
  chat. A question about everything at once is answered from the
  summaries that the workspaces publish for Home. Why: a link from the
  web is untrusted input, and one agent that holds the data of every
  workspace and can act would join all three things that the security
  model keeps apart (#35).
- Home, and a morning brief across all areas, are built from short
  summaries that the agent of each workspace publishes for Home. Each
  workspace decides what it publishes.
- The routing model is chosen by a comparison test of small models on
  labelled real messages: accuracy, latency and cost (M4).
- Search over threads and memory is an activity above the workspaces,
  beside the main agent and Home: a dialogue in the main chat, not a
  screen of its own. The main agent turns the person's words into a
  query; the core searches the threads of every workspace with the
  full-text search of SQLite, and shows the threads it found to the
  person as cards, by workspace. No model reads the text of another
  workspace: a question about a found thread goes to the agent of its
  workspace (M4).

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
  requests: a `304` to an authorized request costs no rate limit. No
  webhooks: a Mac has no public address, and Magos requires no relay
  of a third party.

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
- On every document, question and task, the person can read a summary,
  listen to it, and talk about it by voice. For the documents that a
  session writes, the changes between versions are told in words too
  (#49, Later).
- Artefacts are a strength of Magos, because development needs them:
  mockups, images and sites, with versions and notes pinned to a
  place. A research in M2 names the kinds and their order (#50).
- A workspace's instructions name the language its agents write in.
  This needs no code of its own (M4).

### The thread in M1

- A question from a session shows as a card with its options, the
  Recommended mark and a free-text answer (#4, #15).
- The core keeps the full text of each thread: the messages, each
  tool call and its output. The thread shows it read-only: each tool
  call on one line, its output folded (#9, #15).
- The harness's own log of the session can be opened from the thread
  too. The core copies it when it may; when it may not, the app loads
  it from the Servo that ran the session and shows it from there (#9,
  #15).
- A message can carry images, such as a pasted screenshot (#15).
- The id of a session can be copied, and the session opened in
  Terminal in its worktree, as a way out when Magos hangs or a session
  goes wrong (#14).
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
- Each workspace keeps a list of the model providers that may see its
  data. The agent's plan and the router pick only from that list, and
  Magos starts no session with a provider outside it. The rule is
  written with the security model (#35); the list for Health is
  chosen in M5. Until then, the workspaces hold development work, and
  every provider may see it. Open, for #35: the router reads each
  message before a workspace is known, so which provider its model
  may use.
- An agent acts on a website on the person's behalf through a browser
  on a Servo. The sign-in comes through a password manager's channel
  for agents, so the model never sees the password, and Magos asks
  before an irreversible step (M6).

### Harness terms

- Each Servo starts the unmodified harness CLI under the login of the
  person who owns that computer. Magos never holds a harness login,
  and nobody starts another person's harness.
- The terms of every provider are read again, and shown to a lawyer,
  before the first release for other people (#32).
- Later, beside the subscription: Claude Code on an API key that the
  person gives from their secret store (#47), and the interactive
  Claude Code in a terminal if `claude -p` stops working on a
  subscription (#46).

### A commercial version

- Shared workspaces are the main business feature of a commercial
  version: a person invites other people, and people and agents work
  on projects together (#29).
- Each person's tasks run on that person's own Servo: a shared
  workspace offers a task, and the person's Servo takes it.
- The repository is the open core under Apache-2.0, as the README
  says. Functions for a team and a hosted cloud are separate add-ons.
- When the app connects, the core names its installed modules, and
  the app shows only those parts: a core with no team add-on shows no
  team chat. The app skips a name that it does not know. The core
  still checks each call itself. This comes with the first add-on
  (#29).

### People from outside

- Magos is built in public to bring in people. Until the rules of #44
  are written, code from outside is not taken, and the README says so;
  ideas and bugs are welcome in issues.
- `CONTRIBUTING.md`, `SECURITY.md` and the rule for outside changes
  come when the first person from outside comes, or before the first
  release for other people (#44). The rule: what one person needs
  belongs to the open core; a function for a team or a hosted cloud
  is an add-on, and its code does not come into this repository.

### The interface

- The apps are native. Nothing of the earlier Electron app is ported:
  its code and its UI rules stay behind.
- The rules and the skills for the native interface are chosen when
  the UI work starts. The earlier prototypes of the interface are kept
  as references outside this repository.
- The apps are for Apple devices: the Mac, then the iPhone and the
  iPad (#26). Linux runs only the core and the Servo, with no screen
  (#21, #25). Apps for Windows, for Linux on a desktop and for Android are
  not planned: they wait in Later (#42), and a messenger reaches the
  agents from any phone (#28). This is looked at again when people
  from outside ask for such a client, or when a commercial version
  comes.
- Later: the app icon and the drawings of the empty states, in the
  style of the agent characters (#51); the person reorders the
  workspaces and archives finished threads (#52).

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
- After the signature, Magos ships as one download: one DMG with the
  app and its helper, updates that reach the app by themselves, and a
  first start with no setup, which starts the core, finds the
  harnesses and offers the first workspace (#43).
- The tests of the harness adapters run on sessions recorded on this
  repository with short neutral tasks. A script replaces the home path
  and the account ids before the commit, and the harness version is in
  the file name.

## The tracker

- GitHub's default labels, and the labels that craft makes when it
  first needs them, as `in progress` and `debt`. A research issue
  starts its title with "Research:". No area labels.
- The board is the organization's project "magos roadmap".
