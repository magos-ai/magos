# 0001 One Go program for the core and the Servo, and native apps

- Status: accepted
- Date: 2026-11-28

## Context

Magos runs the harness sessions of one person, as Claude Code and
Codex, on that person's own computers, and shows the work in native
apps. The owner set the shape of the system while the repository was
planned, on 2026-11-26 and 2026-11-27, one question with options at a
time. Then a throwaway spike checked the shape (#3): on one Mac, with
macOS 27.0.1 on Apple silicon, Go 1.27.4 and Swift 6.4, and in a Linux
VM, with the Ubuntu kernel 6.8.0 on aarch64. The numbers below come
from its result.

The earlier private prototype of Magos was a desktop app in Electron,
with a server. Its desktop app took 159 MB of memory and 289 MB of
disk, and its server took 21 MB of memory.

## Options considered

- One Go program `magos` with two roles, and native apps in SwiftUI:
  chosen.
- The same program in Rust: weighed; Go was chosen.
- A desktop app in Electron, as in the earlier prototype: its code and
  its rules for the interface stay behind.
- The harnesses through ACP, and the agents through A2A: not taken.
  Each harness is driven through its own interface.

## Decision

We will build Magos as one Go program and native apps.

- **One program, two roles.** `magos serve` is the core. It keeps the
  workspaces, the projects, the grants, the threads and the runs in
  SQLite. `magos servo` is a Servo. It runs the harness sessions and
  works with files. The core and the Servo run as two processes with
  separate data folders, also in M1, when both run on one Mac.
- **Native apps** (PDR 0015). The Mac app is SwiftUI. It gives the
  macOS capabilities to the core as MCP tools, so the harness adapters
  are written once, in Go. Its bundle id is `dev.kalinichenko.magos`:
  it comes from a domain of the owner, and it names no company. It
  needs macOS 27 or later: M1 runs only on the owner's Mac, and the
  spike checked only macOS 27. The oldest macOS for other people is
  decided before the first release for them (#43).
- **The harnesses, driven directly.** Claude Code runs through
  `claude -p` with stream-json, and Codex through `codex app-server`,
  both in M1. Each Servo starts the unmodified CLI of a harness under
  the login of the person who owns that computer. Magos never holds the
  login of a harness (PDR 0004).
- **One control protocol.** The core, the Servo and the app talk
  through a small versioned control protocol: a local socket first,
  then pairing and HTTPS for a remote core in M3. Events are numbered
  and acknowledged. The Go side and the Swift side share the message
  types, generated from one schema (#9).
- **A helper inside the app.** `magos` ships inside `Magos.app` as a
  helper, registered as a LaunchAgent with `SMAppService`, so launchd
  restarts it. Scheduled work goes on when the app is closed. The core
  opens the app in the background when it needs a macOS tool (#12).
- **A run resumes.** After a reboot or a crash, each run resumes from
  SQLite by its harness session id (#14).
- **The wall between workspaces** is in the data of Magos only, in M1:
  a session sees all that the macOS user sees, and the app says so.
  The sandbox of a session comes in M3 (#24). The tokens of Magos
  itself are in the Keychain.
- **The signature.** M1 is signed with a free Apple Development
  certificate and built from source. Developer ID and notarization come
  when the app goes to other people (#33). After that, Magos ships as
  one download (PDR 0022, #43).
- **Recorded sessions for the tests.** The tests of the harness
  adapters run on sessions recorded on this repository with short
  neutral tasks. A script replaces the home path and the account ids
  before the commit, and the harness version is in the file name.

## Consequences

Good:

- The spike ran three harness sessions at once from one Servo: two
  Claude Code sessions and one Codex session. While they ran, the app
  answered an MCP call of the core in 0.500 ms and 0.577 ms, in two
  runs. The whole tool call of a session, through the core and the
  app, took 8 ms.
- After a `kill -9` of the core in the middle of a run, the Servo kept
  its sessions and connected to a new core. Then it sent the 17 events
  that the sessions made while it had no core. The new core held each
  of the 83 numbers of the run once, with no gap. A cancel through the
  new core ended a turn 41 ms and 1651 ms after it, in the two runs.
- `magos` with SQLite is 7.3 MiB, stripped, on the Mac. The app took
  20.0–21.0 MiB of footprint, the core 4.3–4.5 MiB and the Servo
  4.2 MiB: 28.7–29.5 MiB in all, and 7.4 MiB on disk.
- The core opened its socket 7.2–10.5 ms after its start, as medians.
  The Servo sent its first message 6.3 ms and 8.3 ms after its start.
  The app listed its tools at the core 83 ms and 127 ms after its
  start, in the two full runs.
- `SMAppService` registered the two agents, signed with a free Apple
  Development certificate, with no step of approval, and launchd ran
  them with the app closed.
- `CGO_ENABLED=0` built `magos` with `modernc.org/sqlite` for
  darwin/arm64, linux/arm64 and linux/amd64. The linux/arm64 build ran
  as the core and as the Servo.

Bad:

- Two languages: the message types must be the same in Go and in
  Swift. So they are generated from one schema, and the check proves
  that the generated code matches it (#9).
- An ad-hoc signature does not work for daily use. After a rebuild,
  launchd does not run the agents until a new registration, and macOS
  asks again for each permission. So a daily build needs the
  certificate. The certificate needs the intermediate "Apple Worldwide
  Developer Relations Certification Authority" G3 (#12).
- launchd gives the agents `PATH=/usr/bin:/bin:/usr/sbin:/sbin`, so a
  Servo under launchd does not find `claude` or `codex` by name (#13).
- A permission dialog of macOS names the app, also when the helper or
  its child process reads the folder. The dialog stops the read until
  the person answers, and the stream of the session shows nothing in
  that time.
- The launch constraint of launchd and the permission grants of macOS
  are tied to the signing identifier, which comes from the bundle id. A
  change of the bundle id after the first release for other people
  costs each user their grants. So the bundle id is looked at again
  before #43.
- The spike did not check four parts of this decision: the message
  types generated from one schema (#9), the core that opens the app for
  a macOS tool (#12), the resume of a run by its harness session id
  (#14), and a test that the core and the Servo share no files (#9).

## Invariants

I-1, I-2, I-3, I-4 and I-5, added.

## Confirmation

I-1 and I-2 have `Check:` lines in `docs/ARCHITECTURE.md`. I-3 is
review only until the test of #9 runs the core and the Servo as two
processes and proves that they share nothing but the protocol. I-4 and
I-5: review only.
