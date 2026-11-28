# 0020 Personal data comes in only after its rules

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

- The rules for personal data are set before personal data comes in,
  and are written as an ADR: at most two of untrusted input, private
  data and outside effects in one session; limits on reads, not only
  writes; a log of every tool call; a stop that ends the process
  (#35).
- Each workspace keeps a list of the model providers that may see its
  data. The agent's plan and the router pick only from that list, and
  Magos starts no session with a provider outside it. The rule is
  written with the security model (#35); the list for Health is chosen
  in M5. Until then, the workspaces hold development work, and every
  provider may see it.
- An agent acts on a website on the person's behalf through a browser
  on a Servo. The sign-in comes through a password manager's channel
  for agents, so the model never sees the password, and Magos asks
  before an irreversible step (M6).

## Rests on

- The owner's choice on 2026-11-26 (#40).

## Revisit when

The security model of #35 is settled, or M5 starts and the list of
providers for Health is chosen.

## Links

PDR 0008, #35
