package fpu

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"
)


func fpu(n uint) *mmio.U32 {
	return &(*[3]mmio.U32)(unsafe.Pointer(uintptr(0xe000ef34)))[n]
}


type FPCCR_Bits uint32

func (m FPCCR_Bits) Set()           { fpu(0).SetBits(uint32(m)) }
func (m FPCCR_Bits) Clear()         { fpu(0).ClearBits(uint32(m)) }
func (m FPCCR_Bits) Load() uint32   { return fpu(0).Bits(uint32(m)) }
func (m FPCCR_Bits) Store(b uint32) { fpu(0).StoreBits(uint32(m), b) }
func (m FPCCR_Bits) LoadVal() int   { return fpu(0).Field(uint32(m)) }
func (m FPCCR_Bits) StoreVal(v int) { fpu(0).SetField(uint32(m), v) }

func FPCCR_Load() FPCCR_Bits   { return FPCCR_Bits(fpu(0).Load()) }
func FPCCR_Store(b FPCCR_Bits) { fpu(0).Store(uint32(b)) }

func (b FPCCR_Bits) Field(mask FPCCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_FPCCR(v int, mask FPCCR_Bits) FPCCR_Bits {
	return FPCCR_Bits(bits.Make32(v, uint32(mask)))
}


type FPCAR_Bits uint32

func (m FPCAR_Bits) Set()           { fpu(1).SetBits(uint32(m)) }
func (m FPCAR_Bits) Clear()         { fpu(1).ClearBits(uint32(m)) }
func (m FPCAR_Bits) Load() uint32   { return fpu(1).Bits(uint32(m)) }
func (m FPCAR_Bits) Store(b uint32) { fpu(1).StoreBits(uint32(m), b) }
func (m FPCAR_Bits) LoadVal() int   { return fpu(1).Field(uint32(m)) }
func (m FPCAR_Bits) StoreVal(v int) { fpu(1).SetField(uint32(m), v) }

func FPCAR_Load() FPCAR_Bits   { return FPCAR_Bits(fpu(1).Load()) }
func FPCAR_Store(b FPCAR_Bits) { fpu(1).Store(uint32(b)) }

func (b FPCAR_Bits) Field(mask FPCAR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_FPCAR(v int, mask FPCAR_Bits) FPCAR_Bits {
	return FPCAR_Bits(bits.Make32(v, uint32(mask)))
}


type FPDSCR_Bits uint32

func (m FPDSCR_Bits) Set()           { fpu(2).SetBits(uint32(m)) }
func (m FPDSCR_Bits) Clear()         { fpu(2).ClearBits(uint32(m)) }
func (m FPDSCR_Bits) Load() uint32   { return fpu(2).Bits(uint32(m)) }
func (m FPDSCR_Bits) Store(b uint32) { fpu(2).StoreBits(uint32(m), b) }
func (m FPDSCR_Bits) LoadVal() int   { return fpu(2).Field(uint32(m)) }
func (m FPDSCR_Bits) StoreVal(v int) { fpu(2).SetField(uint32(m), v) }

func FPDSCR_Load() FPDSCR_Bits   { return FPDSCR_Bits(fpu(2).Load()) }
func FPDSCR_Store(b FPDSCR_Bits) { fpu(2).Store(uint32(b)) }

func (b FPDSCR_Bits) Field(mask FPDSCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func Make_FPDSCR(v int, mask FPDSCR_Bits) FPDSCR_Bits {
	return FPDSCR_Bits(bits.Make32(v, uint32(mask)))
}
