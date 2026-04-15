# Architecture

## Current Shape

This repository currently provides a Go package wrapper around the vendored
Nuked-OPL3 C source.

The primary exported type is `OPL3`, which provides:

- register writes
- buffered register writes
- sample generation
- channel meters
- solo-channel masking

## Planned Dynamic-Linking Track

The long-term intent is for this repository to also provide a replaceable shared
library build target with a stable ABI for downstream projects that need a more
LGPL-friendly relinking story.

An initial handle-based `c-shared` export target now exists in
`cmd/nukedopl3shared`.

That shared-library ABI is intentionally small and focused on:

- lifecycle management
- register writes
- sample generation
- solo-channel control

It should be treated as an early foundation rather than a frozen ABI.
