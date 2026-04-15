package main

/*
#include <stdint.h>
*/
import "C"

import (
	"sync"
	"unsafe"

	nukedopl3 "github.com/jebbisson/spice-opl3-nuked"
)

type chipHandle uint64

var (
	handlesMu  sync.Mutex
	nextHandle chipHandle = 1
	handles               = map[chipHandle]*nukedopl3.OPL3{}
)

func main() {}

func addHandle(chip *nukedopl3.OPL3) chipHandle {
	handlesMu.Lock()
	defer handlesMu.Unlock()
	h := nextHandle
	nextHandle++
	handles[h] = chip
	return h
}

func withHandle(handle C.uint64_t, fn func(*nukedopl3.OPL3) C.int) C.int {
	handlesMu.Lock()
	chip := handles[chipHandle(handle)]
	handlesMu.Unlock()
	if chip == nil {
		return -1
	}
	return fn(chip)
}

//export SpiceNukedOPL3_Create
func SpiceNukedOPL3_Create(sampleRate C.uint32_t) C.uint64_t {
	chip := nukedopl3.New(uint32(sampleRate))
	if chip == nil {
		return 0
	}
	return C.uint64_t(addHandle(chip))
}

//export SpiceNukedOPL3_Destroy
func SpiceNukedOPL3_Destroy(handle C.uint64_t) C.int {
	handlesMu.Lock()
	chip := handles[chipHandle(handle)]
	if chip != nil {
		delete(handles, chipHandle(handle))
	}
	handlesMu.Unlock()
	if chip == nil {
		return -1
	}
	chip.Close()
	return 0
}

//export SpiceNukedOPL3_Reset
func SpiceNukedOPL3_Reset(handle C.uint64_t) C.int {
	return withHandle(handle, func(chip *nukedopl3.OPL3) C.int {
		chip.Reset()
		return 0
	})
}

//export SpiceNukedOPL3_WriteRegister
func SpiceNukedOPL3_WriteRegister(handle C.uint64_t, port C.uint16_t, reg C.uint8_t, val C.uint8_t) C.int {
	return withHandle(handle, func(chip *nukedopl3.OPL3) C.int {
		chip.WriteRegister(uint16(port), uint8(reg), uint8(val))
		return 0
	})
}

//export SpiceNukedOPL3_WriteRegisterBuffered
func SpiceNukedOPL3_WriteRegisterBuffered(handle C.uint64_t, port C.uint16_t, reg C.uint8_t, val C.uint8_t) C.int {
	return withHandle(handle, func(chip *nukedopl3.OPL3) C.int {
		chip.WriteRegisterBuffered(uint16(port), uint8(reg), uint8(val))
		return 0
	})
}

//export SpiceNukedOPL3_SetSoloChannel
func SpiceNukedOPL3_SetSoloChannel(handle C.uint64_t, ch C.int32_t) C.int {
	return withHandle(handle, func(chip *nukedopl3.OPL3) C.int {
		chip.SetSoloChannel(int(ch))
		return 0
	})
}

//export SpiceNukedOPL3_GenerateSamples
func SpiceNukedOPL3_GenerateSamples(handle C.uint64_t, frames C.int32_t, samplesOut *C.int16_t, sampleCount C.int32_t, metersOut *C.uint16_t, meterCount C.int32_t) C.int {
	if frames <= 0 || samplesOut == nil {
		return -2
	}
	return withHandle(handle, func(chip *nukedopl3.OPL3) C.int {
		samples, meters, err := chip.GenerateSamplesWithMeters(int(frames))
		if err != nil {
			return -3
		}
		if sampleCount < C.int32_t(len(samples)) {
			return -4
		}
		sampleSlice := unsafe.Slice((*int16)(unsafe.Pointer(samplesOut)), int(sampleCount))
		copy(sampleSlice, samples)
		if metersOut != nil && meterCount > 0 {
			meterSlice := unsafe.Slice((*uint16)(unsafe.Pointer(metersOut)), int(meterCount))
			copy(meterSlice, meters)
		}
		return C.int(len(samples))
	})
}
