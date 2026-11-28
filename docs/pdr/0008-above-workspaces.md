# 0008 Only the router, Home and search stand above the workspaces

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

- Only one agent stands above the workspaces: the router. It sees the
  message and the agents' descriptions, never a workspace's data, and
  hands the message to the default agent of the workspace that fits.
  The default agent answers, or hands the task to another agent of its
  own workspace. When the router is not sure which workspace fits, it
  asks the person.
- The router is the main agent of the main chat. It reads no data of a
  workspace itself, and it acts only through the agents of the
  workspaces. A link to a film goes to the agent of the personal
  workspace, which adds it to the list; a question about a work
  project goes to the agent of the work workspace; the answer shows in
  the main chat. A question about everything at once is answered from
  the summaries that the workspaces publish for Home.
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

Which model provider the router may use is open: Q-2 in `PRODUCT.md`.

## Rests on

- The owner's choice on 2026-11-26 and 2026-11-27 (#40, #53).
- A link from the web is untrusted input. One agent that holds the
  data of every workspace and can act would join all three things that
  the security model keeps apart: untrusted input, private data and
  outside effects (#35).

## Revisit when

The security model of #35 is settled, or the comparison test of
routing models runs (M4).

## Links

PDR 0003, PDR 0020, #35, #53
