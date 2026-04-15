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

That ABI is not implemented yet in this repository, but this package layout is
intended to be the source of truth for that future work.
