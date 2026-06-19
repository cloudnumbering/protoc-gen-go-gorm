# AGENTS.md

## Scope

This repo is a protoc plugin that generates GORM models and helpers. Small template changes can affect many downstream services.

## Worktree Discipline

- Do not work in the main checkout. Use an isolated worktree under `.worktrees/`.
- Base research and changes on the latest `origin/main` unless explicitly told otherwise.
- Do not revert or tidy changes made by another user or agent.

## Change Rules

- Keep generator changes narrow and deterministic.
- Do not change generated output shape without tests or a clear downstream reason.
- Update templates, options, and tests together when behavior changes.
- Prefer root-cause fixes in generation logic over downstream patches in generated files.

## Validation

- Run `git diff --check`.
- Run the focused generator tests for changed behavior.
- For template changes, regenerate the example output used by tests when applicable.
- Use the repo test path from the README when needed: `make build-example && cd test && go test`.
- Report PR/CI status with checks run and downstream compatibility notes.
