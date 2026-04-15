# spice-opl3-nuked

`spice-opl3-nuked` is a thin Go wrapper around the upstream `Nuked-OPL3`
emulator.

## Purpose

This module exists to keep the LGPL-covered OPL emulator isolated from higher
level SpiceSynth logic. It exposes a small Go API for:

- creating and resetting an OPL3 instance
- writing OPL registers
- generating stereo PCM samples
- optional per-channel meters and solo-channel masking used by tooling/debug UIs

## Usage

```go
package main

import nukedopl3 "github.com/jebbisson/spice-opl3-nuked"

func main() {
	chip := nukedopl3.New(44100)
	defer chip.Close()

	chip.WriteRegisterBuffered(0, 0x20, 0x01)
	_, _ = chip.GenerateSamples(256)
}
```

## Build Requirements

- Go with `CGO_ENABLED=1`
- a working C compiler such as `gcc` or `clang`

## Shared Library Build

This repository now includes an initial `c-shared` export target for dynamic
linking experiments and downstream relinking workflows.

Build it with:

```bash
go build -buildmode=c-shared -o libspice_nuked_opl3.so ./cmd/nukedopl3shared
```

See `cmd/nukedopl3shared/README.md` for the exported ABI.

## License

This repository is distributed under `LGPL-2.1-or-later`.

- vendored upstream source: `opl3/`
- upstream license text: `opl3/COPYING`
- repository license text: `LICENSE`
