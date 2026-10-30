# Providers

> **Status:** no provider integration exists yet. GitHub Actions is the first
> planned integration.

A *provider* is a CI system that produces jobs needing runners: GitHub
Actions, GitLab CI, and others in the future.

## Role of a provider integration

Runnex's core speaks in provider-neutral terms. A provider integration is an
adapter that translates between those terms and one CI system. Conceptually,
each integration is responsible for:

- **Discovering demand** — learning that a job is waiting for a runner with
  certain labels (via webhooks or polling, depending on the provider).
- **Registering runners** — obtaining a short-lived, single-use registration
  credential for a runner Runnex is about to start.
- **Reporting status** — mapping provider-side job state to Runnex's model.
- **Cleaning up** — ensuring a destroyed runner is also removed from the
  provider.

Scheduling, resource accounting and the runner lifecycle are *not* the
provider's concern; those live in the core and are shared by all providers.
See [Architecture overview](../architecture/overview.md#provider-agnostic).

The concrete Go interface will be designed alongside the GitHub integration
and validated against GitLab's model before it is considered stable.

## Planned providers

| Provider       | Status  | Notes |
| -------------- | ------- | ----- |
| GitHub Actions | Planned | Via a GitHub App with least-privilege permissions; ephemeral (`--ephemeral`) runners. |
| GitLab CI      | Future  | |
