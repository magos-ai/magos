# 0004 Magos drives unmodified harnesses on the person's own logins

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

Magos orchestrates harnesses: coding-agent CLIs such as Claude Code and
Codex. It runs no skill itself. It gives each harness session that it
starts the skills and the MCP servers of the workspace, and of the
project when the session has one (PDR 0005).

Each Servo, the part of Magos that runs sessions on a computer, starts
the unmodified harness CLI under the login of the person who owns that
computer. Magos never holds a harness login, and nobody starts another
person's harness. Sessions run on the person's own subscriptions, with
any harness and any model.

Later, beside the subscription: Claude Code on an API key that the
person gives from their secret store (#47), and the interactive Claude
Code in a terminal if `claude -p` stops working on a subscription
(#46).

## Rests on

- The owner's choice on 2026-11-26 (#40).
- The terms of the providers allow this use. Not checked yet: Q-1 in
  `PRODUCT.md` (#32).

## Revisit when

The terms of a provider forbid this use (#32), or `claude -p` stops
working on a subscription (#46).

## Links

PDR 0005, PDR 0006, #32, #46, #47
