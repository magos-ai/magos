# Decisions

The decisions taken while this repository was planned, on 2026-11-26
and 2026-11-27. Each one was put to the owner as a question with
options, one at a time.

The decisions of the product are now in [`PRODUCT.md`](PRODUCT.md)
and its PDRs (#2), and a new one goes there. This file holds the rest
until `docs/ARCHITECTURE.md` with its ADRs exists (#6) and the rules
of the tracker go into `CLAUDE.md` (#7); then it goes. Until then, a
new decision of the code or of the tracker is added here in the same
way.

Each entry says what was decided. When an issue or a milestone carries
it, the entry names it.

"craft" in this file is a set of Agent Skills for the flow of
development, from taking an issue to a merged pull request, that the
owner runs in a terminal today.

## The architecture

These hold until the spike (#3) checks them; ADR 0001 records the
result (#6).

- One Go program, `magos`, with two roles: `magos serve`, the core,
  which keeps workspaces, projects, grants, threads and runs in
  SQLite; and `magos servo`, which runs harness sessions and works
  with files. Rust was weighed; Go was chosen.
- The Mac app is SwiftUI. It gives the macOS capabilities to the core
  as MCP tools, so the harness adapters are written once, in Go.
- `magos` ships inside `Magos.app` as a helper, registered as a
  LaunchAgent with `SMAppService`, so launchd restarts it. Scheduled
  work goes on when the app is closed. The core opens the app in the
  background when it needs a macOS tool.
- Harnesses are driven directly: Claude Code through `claude -p` with
  stream-json, Codex through `codex app-server`. Both are in M1. No
  ACP for harnesses, and no A2A.
- The core, the Servo and the app talk through a small versioned
  control protocol: a local socket first, pairing and HTTPS for a
  remote core in M3. Events are numbered and acknowledged. The Go and
  Swift sides share the message types, generated from one schema.
- In M1 the core and the Servo run as two processes with separate data
  folders, and a test proves that they share no files.
- After a reboot or a crash, each run resumes from SQLite by its
  harness session id.
- In M1 the wall between workspaces is in Magos's data only: a session
  sees all that the macOS user sees, and the app says so. The session
  sandbox comes in M3 (#24). Magos's own tokens are in the Keychain.
- M1 is signed with a free Apple Development certificate and built
  from source. Developer ID and notarization come when the app goes to
  other people (#33).
- After the signature, Magos ships as one download: one DMG with the
  app and its helper, updates that reach the app by themselves, and a
  first start with no setup, which starts the core, finds the
  harnesses and offers the first workspace (#43).
- The tests of the harness adapters run on sessions recorded on this
  repository with short neutral tasks. A script replaces the home path
  and the account ids before the commit, and the harness version is in
  the file name.

## The repository

ADR 0002 records these in the pull request of the scaffold (#7).

- One repository holds the open core: the Go program `magos`, with
  the core and the Servo; the app for the Mac; the app for the iPhone
  and the iPad (#26); and the schema of the control protocol (#9). So
  one task can change the Go program and the apps in one pull request.
- `apps/` holds what is built and installed: `apps/magos` in Go,
  `apps/mac` in SwiftUI, and later `apps/ios`. `packages/` holds code
  that two directories of `apps/` or more use, as the schema of the
  protocol and the Swift code that the Mac app and the iPhone app
  share. `packages/` comes with its first package.
- Each directory in `apps/` keeps the files of its own toolchain:
  `apps/magos` has its own `go.mod`, with the module
  `github.com/magos-ai/magos/apps/magos`, and `apps/mac` has its own
  `Package.swift`. So a tag of the product, as `v2.0.0`, sets no
  version of the Go module, and the module path needs no `/v2`.
  `go install` of a version needs tags such as `apps/magos/v0.3.0`.
  They come when people install `magos` that way.
- No tool for a monorepo, as Bun, Turborepo or Nx. The repository
  holds no JavaScript, and for Go and Swift such a tool can only run
  the checks and keep their results. Go and SwiftPM keep their own
  build caches, and one `Makefile` at the root runs the check of each
  directory in `apps/`. This is looked at again when a part in
  JavaScript or TypeScript comes.
- A part stays out of this repository only when it is closed or
  private, when it has its own users and releases without Magos, or
  when a tool requires a repository of its own. So the team and cloud
  add-ons (#29), craft, and a person's own prompts, skills and data
  stay out.

## The tracker

- GitHub's default labels, and the labels that craft makes when it
  first needs them, as `in progress` and `debt`. A research issue
  starts its title with "Research:". No area labels.
- The board is the organization's project "magos roadmap".
