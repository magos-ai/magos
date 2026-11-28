# 0010 A routine proposes, and a tap applies

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

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
  routine with its next run, a switch and its last result. When nobody
  opened a routine's result three times in a row, the routine asks
  whether to turn it off (M2).
- Magos sees GitHub by polling every 1 to 2 minutes with conditional
  requests. No webhooks, and no relay of a third party.

## Rests on

- The owner's choice on 2026-11-26 (#40).
- A `304` to an authorized conditional request does not count against
  the primary rate limit of GitHub
  (https://docs.github.com/en/rest/using-the-rest-api/best-practices-for-using-the-rest-api).
- A Mac has no public address, and Magos requires no relay of a third
  party.

## Revisit when

The person allows most kinds of action to run by themselves, or a core
with a public address comes (M3), so that webhooks need no relay.

## Links

PDR 0011, #41
