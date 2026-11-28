# 0018 What the thread shows in M1

- Status: accepted
- Level: feature
- Date: 2026-11-27

## Decision

- A question from a session shows as a card with its options, the
  Recommended mark and a free-text answer (#4, #15).
- The core keeps the full text of each thread: the messages, each tool
  call and its output. The thread shows it read-only: each tool call
  on one line, its output folded (#9, #15).
- The harness's own log of the session can be opened from the thread
  too. The core copies it when it may; when it may not, the app loads
  it from the Servo that ran the session and shows it from there (#9,
  #15).
- A message can carry images, such as a pasted screenshot (#15).
- The id of a session can be copied, and the session opened in
  Terminal in its worktree, as a way out when Magos hangs or a session
  goes wrong (#14).
- The checks and the review verdicts on the merge question come in M2.
- One macOS notification for each question that waits, and a badge on
  the Dock icon. Other events and sounds come later (#16, #36).

## Rests on

- The owner's choice on 2026-11-26 and 2026-11-27 (#40, #53).

## Revisit when

The UI of the thread is designed: the look may change then.

## Links

#4, #9, #14, #15, #16, #36, #53
