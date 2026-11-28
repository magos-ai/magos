# 0003 A workspace for each area of life, and nothing crosses between them

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

Each area of one person's life is a workspace: work, family and home,
personal projects, career. A workspace has its own agents, connectors,
memory and data. Nothing goes from one workspace into another, and no
agent sees across workspaces. Above them stand only the router, Home
and the search over threads (PDR 0008).

A project lives inside a workspace.

In M1 the wall between workspaces is in Magos's data only: a session
sees all that the macOS user sees, and the app says so. The session
sandbox comes in M3 (#24).

## Rests on

- The owner's choice on 2026-11-26 (#40).

## Revisit when

A task of the person needs the data of two workspaces at once, and the
summaries for Home (PDR 0008) do not cover it.

## Links

PDR 0008, PDR 0009, PDR 0019, #24, #34, #35
