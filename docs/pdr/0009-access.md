# 0009 Access is granted per resource

- Status: accepted
- Level: project
- Date: 2026-11-27

## Decision

- Access is granted per resource, never to a whole app: one
  repository, one mailbox, one calendar. Each grant has a hint for the
  model on what it is. One workspace can hold two accounts of one
  service.
- A grant goes to a workspace or to one of its projects. In M1 a
  resource is a GitHub repository or a GitHub organization, which
  covers all of its repositories, and the person connects them in the
  app (#8, #37). Other connectors join the same place later.

## Rests on

- The owner's choice on 2026-11-26 (#40).

## Revisit when

A connector's service cannot grant access to one resource only.

## Links

PDR 0003, #8, #37
