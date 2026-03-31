# NINS

**NINS Is Not a Shell**

NINS is a shell- and compositor-agnostic desktop control CLI for Linux.

It provides one command surface for desktop tasks that are often split across compositor config, shell config, utilities, and desktop services.

## What NINS does

NINS is intended to:

- wrap existing tools behind a consistent CLI
- detect available backends for features like brightness, volume, and audio switching
- manage generated config for compositor- or shell-owned settings
- provide `status` and `doctor` commands for desktop capability checks

## What NINS is not

NINS is not:

- a command shell
- a desktop shell
- a compositor
- a desktop environment

It is a control layer over the tools and systems that already exist.

## Planned command style

```bash
nins status
nins doctor
nins brightness set 50
nins volume set 30
nins audio input switch
nins profile apply work
nins display off HDMI-A-1
```

Some commands may support simple interactive prompts when no target is given.

## Project direction

NINS is built around two kinds of actions:

- runtime actions for things that can be changed immediately through tools or APIs
- config-backed actions for things owned by the compositor or shell, handled through generated config

The goal is to make fragmented Linux desktop control easier to inspect, automate, and manage from one CLI.

## Status

Early work in progress.
