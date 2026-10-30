# Security Policy

Runnex manages CI/CD runners, which typically hold credentials and execute
untrusted code. Security reports are taken seriously.

## Reporting a vulnerability

**Do not report security vulnerabilities through public GitHub issues,
discussions or pull requests.**

Report them privately through
[GitHub Security Advisories](https://github.com/Kayamii/Runnex/security/advisories/new)
("Report a vulnerability" under the repository's **Security** tab).

<!-- TODO(maintainers): add a security contact email once one exists. -->

Please include:

- A description of the issue and its impact.
- Steps to reproduce, or a proof of concept.
- The affected version (`runnex version`) or commit.
- Any suggested mitigation.

## Never publish credentials

Do not include real credentials in any report, issue, log excerpt or pull
request — this includes GitHub tokens, GitHub App private keys, runner
registration tokens, SSH keys and cloud credentials. Redact them before
sharing logs. If a credential has been exposed, revoke it immediately; do not
wait for a response.

## Responsible disclosure

- You will receive an acknowledgement as soon as a maintainer is able to
  review the report. This is a small volunteer project, so no fixed response
  time is guaranteed.
- We will work with you to understand and confirm the issue, and keep you
  informed about the fix.
- Please give us a reasonable amount of time to release a fix before
  disclosing publicly. We will coordinate the disclosure date with you.
- Reporters are credited in the advisory unless they prefer otherwise.

## Supported versions

Runnex has not published a release yet. Until 1.0, only the latest release
receives security fixes.

| Version        | Supported |
| -------------- | --------- |
| Latest release | Yes       |
| Older releases | No        |
