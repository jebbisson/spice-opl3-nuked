// Copyright (c) 2026 Jeb Bisson.
//
// This package vendors and wraps Nuked-OPL3, which is licensed under the
// GNU Lesser General Public License v2.1 or later. See LICENSE and opl3/COPYING.

// Package nukedopl3 provides a CGo wrapper around the Nuked-OPL3 emulator.
//
// It exposes a minimal interface for initializing the OPL3 chip, writing
// to hardware registers, and generating signed 16-bit stereo PCM samples.
// The vendored C source (opl3/) is compiled directly into the Go binary
// via CGo and requires a C compiler (gcc or clang) with CGO_ENABLED=1.
package nukedopl3
