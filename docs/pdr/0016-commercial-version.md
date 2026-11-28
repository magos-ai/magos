# 0016 Shared workspaces lead the commercial version

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

- Shared workspaces are the main business feature of a commercial
  version: a person invites other people, and people and agents work
  on projects together (#29).
- Each person's tasks run on that person's own Servo: a shared
  workspace offers a task, and the person's Servo takes it.
- When the app connects, the core names its installed modules, and the
  app shows only those parts: a core with no team add-on shows no team
  chat. The app skips a name that it does not know. The core still
  checks each call itself. This comes with the first add-on (#29).

## Rests on

- The owner's choice in an earlier session, recorded in #2, #29 and
  #40.
- Nobody starts another person's harness (PDR 0004).

## Revisit when

The add-ons are planned (#29).

## Links

PDR 0001, PDR 0004, #29
