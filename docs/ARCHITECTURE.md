# Architecture

This file is the map of the code: its parts, their borders, and the
rules that the code keeps. The ADRs in `docs/adr/` record why.

## Bird's-eye view

Magos is one Go program, `magos`, and native apps (ADR 0001). The
program has two roles. `magos serve` is the core: it will keep the
workspaces, the projects, the grants, the threads and the runs in
SQLite. `magos servo` is a Servo: it will run the harness sessions, as
Claude Code and Codex, and work with files. The Mac app shows the work
to the person, and gives the macOS capabilities to the core as MCP
tools. The core, the Servo and the app will talk through a small
versioned control protocol (#9).

So far only the frame exists: `magos serve` and `magos servo` start,
wait and stop, and the Mac app opens one window.

## Codemap

- `apps/`: what is built and installed (ADR 0002).
  - `apps/magos`: the Go program, a Go module of its own. `main.go`
    reads the command and runs one role. `internal/core` is the core,
    and `internal/servo` is the Servo.
  - `apps/mac`: the Mac app, a Swift package of its own, with the
    `Info.plist` of `Magos.app`.
- `packages/`: code that two directories of `apps/` or more use, as the
  schema of the control protocol (#9). It comes with its first package.
- `docs/`: this file, the ADRs in `docs/adr/`, and `PRODUCT.md`, with
  the decisions of the product in `docs/pdr/`.
- `Makefile`: `make check`, the check of each directory in `apps/`, and
  `make app`.

## Modules

- **The core:** the package `apps/magos/internal/core`.
- **The Servo:** the package `apps/magos/internal/servo`.
- **The Mac app:** `apps/mac`.

The API of the core and of the Servo is the exported names of its
package. The two know each other only through the control protocol, so
neither imports a package of the other (I-1, I-2).

## Layers

In `apps/magos`, `main.go` is the edge: it reads the command and gives
the output. The logic of each role is in its package under `internal/`.
The data layer comes with SQLite (#8).

## Cross-cutting concerns

- **Command output.** Each command of `magos` gives the result first,
  and the next step on its own line. With no command, `magos` prints
  its usage. A command that uses colour shows it only on a terminal,
  and never when `NO_COLOR` is set.

## Invariants

- **I-1.** The core imports no package of the Servo.
  Check: grep `magos/apps/magos/internal/servo` in `apps/magos/internal/core`.
  ADR: 0001.
- **I-2.** The Servo imports no package of the core.
  Check: grep `magos/apps/magos/internal/core` in `apps/magos/internal/servo`.
  ADR: 0001.
- **I-3.** The core and the Servo share no files: each one keeps its
  own data folder. Review only, until the test of #9.
  ADR: 0001.
- **I-4.** Magos never reads, stores or passes the login of a harness.
  ADR: 0001.
- **I-5.** Each Servo starts the unmodified CLI of a harness, under the
  login of the person who owns that computer.
  ADR: 0001.
- **I-6.** This repository holds no code of a closed add-on.
  ADR: 0002.
- **I-7.** Everything in this repository is under Apache-2.0.
  ADR: 0002.

## Limits of decomposition

None yet: the code is too small to set them from its own sizes.

## ADRs

- [0001 One Go program for the core and the Servo, and native apps](adr/0001-one-go-program-and-native-apps.md)
- [0002 One repository for the core and the apps](adr/0002-one-repository-for-the-core-and-the-apps.md)
