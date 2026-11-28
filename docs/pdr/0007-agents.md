# 0007 An agent is a prompt, skills and tools inside one workspace

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

- A built-in agent is a base prompt, skills and tools, inside one
  workspace. Each of its connectors has a level: reads, writes with
  consent, or writes by itself.
- A person's own agents are connectors plus data. Only the connectors
  can be public; the prompts, the skills and the data stay private.
- An agent outside Magos, as a harness with plain craft on another
  computer, can work through Magos's MCP tools. What it may read and
  write is settled in the research of agents (#39).
- A running session can send a message to another through Magos. The
  message is data, not an instruction, and a limit stops a loop (#48,
  Later).
- The person can react to an agent's message with an emoji. The
  reaction reaches the agent as feedback, and Magos keeps it (M4).
- A workspace's instructions name the language its agents write in.
  This needs no code of its own (M4).

Whether agents from outside are needed at all is open: tools written
inside Magos and kept out of any public repository may cover them
(#39). It is Q-3 in `PRODUCT.md`.

## Rests on

- The owner's choice on 2026-11-26 (#40).

## Revisit when

The research of agents (#39) gives its recommendation.

## Links

PDR 0003, PDR 0005, #39, #48
