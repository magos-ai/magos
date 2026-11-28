# 0015 Native apps for Apple devices only

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

- The apps are native. Nothing of the earlier Electron app is ported:
  its code and its UI rules stay behind.
- The rules and the skills for the native interface are chosen when
  the UI work starts. The earlier prototypes of the interface are kept
  as references outside this repository.
- The apps are for Apple devices: the Mac, then the iPhone and the
  iPad (#26). Linux runs only the core and the Servo, with no screen
  (#21, #25). Apps for Windows, for Linux on a desktop and for Android
  are not planned: they wait in Later (#42), and a messenger reaches
  the agents from any phone (#28).
- Later: the app icon and the drawings of the empty states, in the
  style of the agent characters (#51); the person reorders the
  workspaces and archives finished threads (#52).

## Rests on

- The owner's choice on 2026-11-26 (#40).

## Revisit when

People from outside ask for a client on another platform, or a
commercial version comes.

## Links

PDR 0014, #21, #25, #26, #28, #42, #51, #52
