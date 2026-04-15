# nukedopl3shared

This command builds a shared library wrapper around `spice-opl3-nuked` using
Go's `c-shared` build mode.

## Build

```bash
go build -buildmode=c-shared -o libspice_nuked_opl3.so ./cmd/nukedopl3shared
```

Common outputs by platform:

- Linux: `libspice_nuked_opl3.so`
- macOS: `libspice_nuked_opl3.dylib`
- Windows: `spice_nuked_opl3.dll`

The build also emits a generated C header next to the library artifact.

## Exported ABI

- `SpiceNukedOPL3_Create`
- `SpiceNukedOPL3_Destroy`
- `SpiceNukedOPL3_Reset`
- `SpiceNukedOPL3_WriteRegister`
- `SpiceNukedOPL3_WriteRegisterBuffered`
- `SpiceNukedOPL3_SetSoloChannel`
- `SpiceNukedOPL3_GenerateSamples`

The API is intentionally small and handle-based so a host application can load
and replace the shared library without binding directly to Go internals.
