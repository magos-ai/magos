# CLAUDE.md

Before a spec, a plan or the first edit of a new feature, read
`docs/ARCHITECTURE.md`. At the start of a new feature, and for a
decision, load `craft:architecture`.

## Commands

- `make check`: the check of the whole repository, the same as CI.
  - `apps/magos`: `gofmt -l`, `go vet` and `go test`.
  - `apps/mac`: `swift format lint --strict`, `plutil -lint` of
    `Info.plist`, and `swift build`. On Linux, `make check` skips this
    part and says so.
- `make check-magos` or `make check-mac`: one part of the check.
- `make app`: builds `apps/mac/.build/Magos.app`. Open it with
  `open apps/mac/.build/Magos.app`.
- `go -C apps/magos run . serve` runs the core, and
  `go -C apps/magos run . servo` runs the Servo. Ctrl-C stops each one.
- `gofmt -w apps/magos` and
  `swift format --in-place --recursive apps/mac/Package.swift apps/mac/Sources`
  format the code.

The tools are Go, at the version in `apps/magos/go.mod`, and on macOS
the Swift toolchain of the Xcode Command Line Tools, Swift 6.4 or
later. The format and the lint come with the toolchains and with
macOS: `gofmt`, `go vet`, `swift format` and `plutil`. Install no
other tool for them.

## CI

`.github/workflows/check.yml` runs `make check` on every pull request
and on each push to `main`, on Linux (`ubuntu-latest`) and on macOS
(`xcode-27`). Its job `check`
is the required check of the `main` ruleset. That job runs always, and
it fails unless both runs of `make check` succeeded: a job that is
skipped after a failed job in its `needs` reports success, also as a
required check. The workflow has no `paths:` filter, because the
required check of a workflow that does not run stays pending, and then
the pull request cannot merge.

## The tracker

- The labels are GitHub's default labels, and the labels that craft
  makes when it first needs them, as `in progress` and `debt`. There
  are no area labels.
- A research issue starts its title with "Research:".
- The board is the organization's project "magos roadmap".
