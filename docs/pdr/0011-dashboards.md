# 0011 The workspace's agent assembles its dashboard from primitives

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

- Nothing on a dashboard is hard-coded: the workspace's agent
  assembles it from primitives. The view language is chosen by
  research (#20).
- The same language draws a card inside a chat answer, and a card can
  be pinned to the dashboard (#20).
- The workspace's agent sets the hints above the input field, as it
  builds the dashboard (M2).
- Artefacts are a strength of Magos, because development needs them:
  mockups, images and sites, with versions and notes pinned to a
  place. A research in M2 names the kinds and their order (#50).

## Rests on

- The owner's choice on 2026-11-26 (#40).

## Revisit when

The research of #20 chooses the view language.

## Links

PDR 0010, PDR 0014, #20, #50
