# 0002 One repository for the core and the apps

- Status: accepted
- Date: 2026-11-28

## Context

Magos has a Go program, `magos`, with the core and the Servo, and an
app for the Mac. Later it gets an app for the iPhone and the iPad (#26)
and the schema of the control protocol (#9). One task often changes the
Go program and an app together, as a new message of the protocol. The
owner decided the layout on 2026-11-27, before the first code (#55).

## Options considered

- One repository with `apps/` and `packages/`, and one `Makefile` at
  the root: chosen.
- A repository for each part: a change of the protocol needs a pull
  request in each repository, in step.
- One repository with a tool for a monorepo, as Bun, Turborepo or Nx:
  the repository holds no JavaScript, and for Go and Swift such a tool
  can only run the checks and keep their results.

## Decision

We will keep the open core in one repository.

- One repository holds the open core: the Go program `magos`, with the
  core and the Servo; the app for the Mac; the app for the iPhone and
  the iPad (#26); and the schema of the control protocol (#9).
- `apps/` holds what is built and installed: `apps/magos` in Go,
  `apps/mac` in SwiftUI, and later `apps/ios`. `packages/` holds code
  that two directories of `apps/` or more use, as the schema of the
  protocol and the Swift code that the Mac app and the iPhone app
  share. `packages/` comes with its first package.
- Each directory in `apps/` keeps the files of its own toolchain:
  `apps/magos` has its own `go.mod`, with the module
  `github.com/magos-ai/magos/apps/magos`, and `apps/mac` has its own
  `Package.swift`.
- One `Makefile` at the root runs the check of each directory in
  `apps/`. Go and SwiftPM keep their own build caches. A tool for a
  monorepo is looked at again when a part in JavaScript or TypeScript
  comes.
- A part stays out of this repository only when it is closed or
  private, when it has its own users and releases without Magos, or
  when a tool requires a repository of its own. So the team and cloud
  add-ons (#29, PDR 0001), craft, and a person's own prompts, skills
  and data stay out.
- Everything in this repository is under Apache-2.0 (PDR 0001).

## Consequences

Good:

- One pull request changes the Go program and the apps together, and
  one `make check` covers them.
- A tag of the product, as `v2.0.0`, sets no version of the Go module,
  so the module path needs no `/v2`.

Bad:

- `go install` of a version needs tags such as `apps/magos/v0.3.0`.
  They come when people install `magos` that way.
- Each pull request runs every check, also a pull request that changes
  only documents. The workflow has no `paths:` filter, because the
  required check of a workflow that does not run stays pending, and
  then the pull request cannot merge.

## Invariants

I-6 and I-7, added.

## Confirmation

Review only.
