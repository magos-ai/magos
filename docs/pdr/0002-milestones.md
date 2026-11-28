# 0002 The order of the milestones, M1 to M6, then Later

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

The work goes in this order:

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

M1 is done when the owner can run in Magos the same flow that the
owner runs in a terminal today: from `#N` in a chat to a merged pull
request, three sessions at once, nothing lost after a reboot. Five real
tasks pass the measurement against that flow (#17). The full list is
the `Done when` of the milestone M1.

Milestones have no due dates: their numbers give the order. Milestones
M3 to M6 get their `Done when` when the one before closes.

## Rests on

- The owner's choice on 2026-11-26 (#40).
- The claim that no other product runs development with an agent that
  chooses the harness, the model and the review of each step (PDR
  0006, #45).

## Revisit when

A milestone closes: the next one gets its `Done when` then. Or a check
of PDR 0006 shows a product that does all that the claim names.

## Links

PDR 0006, #17, #45, #56
