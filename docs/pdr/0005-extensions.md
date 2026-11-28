# 0005 Standard extensions only, added at runtime

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

- There are two kinds of extension only: Agent Skills with harness
  plugins, and MCP servers. Magos has no format of its own for them,
  and it knows no skill by name.
- A skill or a plugin comes from a link. Magos shows what is inside,
  and after a yes pins it at the commit it was taken from. A newer
  version upstream comes in only after the person sees its diff and
  says yes (#38).
- Tools and connectors are MCP servers. They are added at runtime,
  with no rebuild of the app, and a person can write them inside
  Magos.
- Magos's own contract is its versioned MCP tools: sessions, Servos, a
  question to the person, a dashboard card, a proposed task.
- In M1, on Codex: craft, a set of Agent Skills for the flow of
  development, reaches a session as the person's own install. The
  person installs craft in Codex and trusts its hooks, and trusts them
  again after each release of craft, as craft's README says. Before a
  Codex session starts, Magos checks with `hooks/list` that each
  enabled hook of the session is trusted, and stops with a message
  when one is not. Every Codex session on that computer gets craft,
  and Magos does not pin it until #38. The other ways are in #5 (#11).

## Rests on

- The owner's choice on 2026-11-26 (#40).
- Codex loads a plugin only when it is installed in `CODEX_HOME`. Its
  hooks run only when they are trusted, or when a thread skips the
  trust check, which skips it for every other hook of the person too
  (#5).

## Revisit when

Codex can load a plugin for one session (#5), or a harness that Magos
drives takes neither Agent Skills nor MCP servers.

## Links

PDR 0004, #5, #11, #38
