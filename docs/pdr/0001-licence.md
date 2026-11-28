# 0001 Apache-2.0 open core, with team and cloud as closed add-ons

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

Everything in this repository is under the Apache License 2.0. The
repository is the open core: everything that one person needs on their
own computers, free of charge. Functions for a team and for a hosted
cloud come later as separate, closed add-ons to one server. Their code
does not come into this repository.

The line between the two: what one person needs belongs to the open
core; a function for a team or a hosted cloud is an add-on.

## Rests on

- The owner's choice, recorded in the README, in the planning (#40)
  and in #29: "The open core stays free for one person. A team and the
  cloud come as closed add-ons to one server."
- Assumption: a team pays for the work that people share, which one
  person does not need (PDR 0016).

## Revisit when

The first add-on is planned (#29) and needs a part that one person
needs too, or a person from outside offers code for a function of a
team (#44).

## Links

README, PDR 0016, PDR 0021, #29, #44
