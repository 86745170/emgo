// +build f303xe

// Peripheral: OPAMP_Periph  Operational Amplifier (OPAMP).
// Instances:
//  OPAMP   mmap.OPAMP_BASE
//  OPAMP1  mmap.OPAMP1_BASE
//  OPAMP2  mmap.OPAMP2_BASE
//  OPAMP3  mmap.OPAMP3_BASE
//  OPAMP4  mmap.OPAMP4_BASE
// Registers:
//  0x00 32  CSR Control and status register.
// Import:
//  stm32/o/f303xe/mmap
package opamp

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	OPAMPxEN    CSR_Bits = 0x01 << 0  //+ OPAMP enable.
	FORCEVP     CSR_Bits = 0x01 << 1  //+ Connect the internal references to the plus input of the OPAMPX.
	VPSEL       CSR_Bits = 0x03 << 2  //+ Non inverting input selection.
	VPSEL_0     CSR_Bits = 0x01 << 2  //  Bit 0.
	VPSEL_1     CSR_Bits = 0x02 << 2  //  Bit 1.
	VMSEL       CSR_Bits = 0x03 << 5  //+ Inverting input selection.
	VMSEL_0     CSR_Bits = 0x01 << 5  //  Bit 0.
	VMSEL_1     CSR_Bits = 0x02 << 5  //  Bit 1.
	TCMEN       CSR_Bits = 0x01 << 7  //+ Timer-Controlled Mux mode enable.
	VMSSEL      CSR_Bits = 0x01 << 8  //+ Inverting input secondary selection.
	VPSSEL      CSR_Bits = 0x03 << 9  //+ Non inverting input secondary selection.
	VPSSEL_0    CSR_Bits = 0x01 << 9  //  Bit 0.
	VPSSEL_1    CSR_Bits = 0x02 << 9  //  Bit 1.
	CALON       CSR_Bits = 0x01 << 11 //+ Calibration mode enable.
	CALSEL      CSR_Bits = 0x03 << 12 //+ Calibration selection.
	CALSEL_0    CSR_Bits = 0x01 << 12 //  Bit 0.
	CALSEL_1    CSR_Bits = 0x02 << 12 //  Bit 1.
	PGGAIN      CSR_Bits = 0x0F << 14 //+ Gain in PGA mode.
	PGGAIN_0    CSR_Bits = 0x01 << 14 //  Bit 0.
	PGGAIN_1    CSR_Bits = 0x02 << 14 //  Bit 1.
	PGGAIN_2    CSR_Bits = 0x04 << 14 //  Bit 2.
	PGGAIN_3    CSR_Bits = 0x08 << 14 //  Bit 3.
	USERTRIM    CSR_Bits = 0x01 << 18 //+ User trimming enable.
	TRIMOFFSETP CSR_Bits = 0x1F << 19 //+ Offset trimming value (PMOS).
	TRIMOFFSETN CSR_Bits = 0x1F << 24 //+ Offset trimming value (NMOS).
	TSTREF      CSR_Bits = 0x01 << 29 //+ It enables the switch to put out the internal reference.
	OUTCAL      CSR_Bits = 0x01 << 30 //+ OPAMP output status flag.
	LOCK        CSR_Bits = 0x01 << 31 //+ OPAMP lock.
)

const (
	OPAMPxENn    = 0
	FORCEVPn     = 1
	VPSELn       = 2
	VMSELn       = 5
	TCMENn       = 7
	VMSSELn      = 8
	VPSSELn      = 9
	CALONn       = 11
	CALSELn      = 12
	PGGAINn      = 14
	USERTRIMn    = 18
	TRIMOFFSETPn = 19
	TRIMOFFSETNn = 24
	TSTREFn      = 29
	OUTCALn      = 30
	LOCKn        = 31
)