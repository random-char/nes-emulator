package cpu

var lookup []*instruction

type instruction struct {
	OpName   string
	AMName   string
	AddrMode func() uint8
	Op       func() uint8
	Cycles   uint8
}

func initLookupFor(cpu *MOSTechnology6502) {
	lookup = []*instruction{
		{OpName: op_BRK, AMName: am_IMM, Op: cpu.BRK, AddrMode: cpu.IMM, Cycles: 7}, //000->0x00
		{OpName: op_ORA, AMName: am_IZX, Op: cpu.ORA, AddrMode: cpu.IZX, Cycles: 6}, //001->0x01
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //002->0x02
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //003->0x03
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 3}, //004->0x04
		{OpName: op_ORA, AMName: am_ZP0, Op: cpu.ORA, AddrMode: cpu.ZP0, Cycles: 3}, //005->0x05
		{OpName: op_ASL, AMName: am_ZP0, Op: cpu.ASL, AddrMode: cpu.ZP0, Cycles: 5}, //006->0x06
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 5}, //007->0x07
		{OpName: op_PHP, AMName: am_IMP, Op: cpu.PHP, AddrMode: cpu.IMP, Cycles: 3}, //008->0x08
		{OpName: op_ORA, AMName: am_IMM, Op: cpu.ORA, AddrMode: cpu.IMM, Cycles: 2}, //009->0x09
		{OpName: op_ASL, AMName: am_IMP, Op: cpu.ASL, AddrMode: cpu.IMP, Cycles: 2}, //010->0x0A
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //011->0x0B
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //012->0x0C
		{OpName: op_ORA, AMName: am_ABS, Op: cpu.ORA, AddrMode: cpu.ABS, Cycles: 4}, //013->0x0D
		{OpName: op_ASL, AMName: am_ABS, Op: cpu.ASL, AddrMode: cpu.ABS, Cycles: 6}, //014->0x0E
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //015->0x0F
		{OpName: op_BPL, AMName: am_REL, Op: cpu.BPL, AddrMode: cpu.REL, Cycles: 2}, //016->0x10
		{OpName: op_ORA, AMName: am_IZY, Op: cpu.ORA, AddrMode: cpu.IZY, Cycles: 5}, //017->0x11
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //018->0x12
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //019->0x13
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //020->0x14
		{OpName: op_ORA, AMName: am_ZPX, Op: cpu.ORA, AddrMode: cpu.ZPX, Cycles: 4}, //021->0x15
		{OpName: op_ASL, AMName: am_ZPX, Op: cpu.ASL, AddrMode: cpu.ZPX, Cycles: 6}, //022->0x16
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //023->0x17
		{OpName: op_CLC, AMName: am_IMP, Op: cpu.CLC, AddrMode: cpu.IMP, Cycles: 2}, //024->0x18
		{OpName: op_ORA, AMName: am_ABY, Op: cpu.ORA, AddrMode: cpu.ABY, Cycles: 4}, //025->0x19
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //026->0x1A
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //027->0x1B
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //028->0x1C
		{OpName: op_ORA, AMName: am_ABX, Op: cpu.ORA, AddrMode: cpu.ABX, Cycles: 4}, //029->0x1D
		{OpName: op_ASL, AMName: am_ABX, Op: cpu.ASL, AddrMode: cpu.ABX, Cycles: 7}, //030->0x1E
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //031->0x1F
		{OpName: op_JSR, AMName: am_ABS, Op: cpu.JSR, AddrMode: cpu.ABS, Cycles: 6}, //032->0x20
		{OpName: op_AND, AMName: am_IZX, Op: cpu.AND, AddrMode: cpu.IZX, Cycles: 6}, //033->0x21
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //034->0x22
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //035->0x23
		{OpName: op_BIT, AMName: am_ZP0, Op: cpu.BIT, AddrMode: cpu.ZP0, Cycles: 3}, //036->0x24
		{OpName: op_AND, AMName: am_ZP0, Op: cpu.AND, AddrMode: cpu.ZP0, Cycles: 3}, //037->0x25
		{OpName: op_ROL, AMName: am_ZP0, Op: cpu.ROL, AddrMode: cpu.ZP0, Cycles: 5}, //038->0x26
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 5}, //039->0x27
		{OpName: op_PLP, AMName: am_IMP, Op: cpu.PLP, AddrMode: cpu.IMP, Cycles: 4}, //040->0x28
		{OpName: op_AND, AMName: am_IMM, Op: cpu.AND, AddrMode: cpu.IMM, Cycles: 2}, //041->0x29
		{OpName: op_ROL, AMName: am_IMP, Op: cpu.ROL, AddrMode: cpu.IMP, Cycles: 2}, //042->0x2A
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //043->0x2B
		{OpName: op_BIT, AMName: am_ABS, Op: cpu.BIT, AddrMode: cpu.ABS, Cycles: 4}, //044->0x2C
		{OpName: op_AND, AMName: am_ABS, Op: cpu.AND, AddrMode: cpu.ABS, Cycles: 4}, //045->0x2D
		{OpName: op_ROL, AMName: am_ABS, Op: cpu.ROL, AddrMode: cpu.ABS, Cycles: 6}, //046->0x2E
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //047->0x2F
		{OpName: op_BMI, AMName: am_REL, Op: cpu.BMI, AddrMode: cpu.REL, Cycles: 2}, //048->0x30
		{OpName: op_AND, AMName: am_IZY, Op: cpu.AND, AddrMode: cpu.IZY, Cycles: 5}, //049->0x31
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //050->0x32
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //051->0x33
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //052->0x34
		{OpName: op_AND, AMName: am_ZPX, Op: cpu.AND, AddrMode: cpu.ZPX, Cycles: 4}, //053->0x35
		{OpName: op_ROL, AMName: am_ZPX, Op: cpu.ROL, AddrMode: cpu.ZPX, Cycles: 6}, //054->0x36
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //055->0x37
		{OpName: op_SEC, AMName: am_IMP, Op: cpu.SEC, AddrMode: cpu.IMP, Cycles: 2}, //056->0x38
		{OpName: op_AND, AMName: am_ABY, Op: cpu.AND, AddrMode: cpu.ABY, Cycles: 4}, //057->0x39
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //058->0x3A
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //059->0x3B
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //060->0x3C
		{OpName: op_AND, AMName: am_ABX, Op: cpu.AND, AddrMode: cpu.ABX, Cycles: 4}, //061->0x3D
		{OpName: op_ROL, AMName: am_ABX, Op: cpu.ROL, AddrMode: cpu.ABX, Cycles: 7}, //062->0x3E
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //063->0x3F
		{OpName: op_RTI, AMName: am_IMP, Op: cpu.RTI, AddrMode: cpu.IMP, Cycles: 6}, //064->0x40
		{OpName: op_EOR, AMName: am_IZX, Op: cpu.EOR, AddrMode: cpu.IZX, Cycles: 6}, //065->0x41
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //066->0x42
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //067->0x43
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 3}, //068->0x44
		{OpName: op_EOR, AMName: am_ZP0, Op: cpu.EOR, AddrMode: cpu.ZP0, Cycles: 3}, //069->0x45
		{OpName: op_LSR, AMName: am_ZP0, Op: cpu.LSR, AddrMode: cpu.ZP0, Cycles: 5}, //070->0x46
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 5}, //071->0x47
		{OpName: op_PHA, AMName: am_IMP, Op: cpu.PHA, AddrMode: cpu.IMP, Cycles: 3}, //072->0x48
		{OpName: op_EOR, AMName: am_IMM, Op: cpu.EOR, AddrMode: cpu.IMM, Cycles: 2}, //073->0x49
		{OpName: op_LSR, AMName: am_IMP, Op: cpu.LSR, AddrMode: cpu.IMP, Cycles: 2}, //074->0x4A
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //075->0x4B
		{OpName: op_JMP, AMName: am_ABS, Op: cpu.JMP, AddrMode: cpu.ABS, Cycles: 3}, //076->0x4C
		{OpName: op_EOR, AMName: am_ABS, Op: cpu.EOR, AddrMode: cpu.ABS, Cycles: 4}, //077->0x4D
		{OpName: op_LSR, AMName: am_ABS, Op: cpu.LSR, AddrMode: cpu.ABS, Cycles: 6}, //078->0x4E
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //079->0x4F
		{OpName: op_BVC, AMName: am_REL, Op: cpu.BVC, AddrMode: cpu.REL, Cycles: 2}, //080->0x50
		{OpName: op_EOR, AMName: am_IZY, Op: cpu.EOR, AddrMode: cpu.IZY, Cycles: 5}, //081->0x51
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //082->0x52
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //083->0x53
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //084->0x54
		{OpName: op_EOR, AMName: am_ZPX, Op: cpu.EOR, AddrMode: cpu.ZPX, Cycles: 4}, //085->0x55
		{OpName: op_LSR, AMName: am_ZPX, Op: cpu.LSR, AddrMode: cpu.ZPX, Cycles: 6}, //086->0x56
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //087->0x57
		{OpName: op_CLI, AMName: am_IMP, Op: cpu.CLI, AddrMode: cpu.IMP, Cycles: 2}, //088->0x58
		{OpName: op_EOR, AMName: am_ABY, Op: cpu.EOR, AddrMode: cpu.ABY, Cycles: 4}, //089->0x59
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //090->0x5A
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //091->0x5B
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //092->0x5C
		{OpName: op_EOR, AMName: am_ABX, Op: cpu.EOR, AddrMode: cpu.ABX, Cycles: 4}, //093->0x5D
		{OpName: op_LSR, AMName: am_ABX, Op: cpu.LSR, AddrMode: cpu.ABX, Cycles: 7}, //094->0x5E
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //095->0x5F
		{OpName: op_RTS, AMName: am_IMP, Op: cpu.RTS, AddrMode: cpu.IMP, Cycles: 6}, //096->0x60
		{OpName: op_ADC, AMName: am_IZX, Op: cpu.ADC, AddrMode: cpu.IZX, Cycles: 6}, //097->0x61
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //098->0x62
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //099->0x63
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 3}, //100->0x64
		{OpName: op_ADC, AMName: am_ZP0, Op: cpu.ADC, AddrMode: cpu.ZP0, Cycles: 3}, //101->0x65
		{OpName: op_ROR, AMName: am_ZP0, Op: cpu.ROR, AddrMode: cpu.ZP0, Cycles: 5}, //102->0x66
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 5}, //103->0x67
		{OpName: op_PLA, AMName: am_IMP, Op: cpu.PLA, AddrMode: cpu.IMP, Cycles: 4}, //104->0x68
		{OpName: op_ADC, AMName: am_IMM, Op: cpu.ADC, AddrMode: cpu.IMM, Cycles: 2}, //105->0x69
		{OpName: op_ROR, AMName: am_IMP, Op: cpu.ROR, AddrMode: cpu.IMP, Cycles: 2}, //106->0x6A
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //107->0x6B
		{OpName: op_JMP, AMName: am_IND, Op: cpu.JMP, AddrMode: cpu.IND, Cycles: 5}, //108->0x6C
		{OpName: op_ADC, AMName: am_ABS, Op: cpu.ADC, AddrMode: cpu.ABS, Cycles: 4}, //109->0x6D
		{OpName: op_ROR, AMName: am_ABS, Op: cpu.ROR, AddrMode: cpu.ABS, Cycles: 6}, //110->0x6E
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //111->0x6F
		{OpName: op_BVS, AMName: am_REL, Op: cpu.BVS, AddrMode: cpu.REL, Cycles: 2}, //112->0x70
		{OpName: op_ADC, AMName: am_IZY, Op: cpu.ADC, AddrMode: cpu.IZY, Cycles: 5}, //113->0x71
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //114->0x72
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //115->0x73
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //116->0x74
		{OpName: op_ADC, AMName: am_ZPX, Op: cpu.ADC, AddrMode: cpu.ZPX, Cycles: 4}, //117->0x75
		{OpName: op_ROR, AMName: am_ZPX, Op: cpu.ROR, AddrMode: cpu.ZPX, Cycles: 6}, //118->0x76
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //119->0x77
		{OpName: op_SEI, AMName: am_IMP, Op: cpu.SEI, AddrMode: cpu.IMP, Cycles: 2}, //120->0x78
		{OpName: op_ADC, AMName: am_ABY, Op: cpu.ADC, AddrMode: cpu.ABY, Cycles: 4}, //121->0x79
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //122->0x7A
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //123->0x7B
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //124->0x7C
		{OpName: op_ADC, AMName: am_ABX, Op: cpu.ADC, AddrMode: cpu.ABX, Cycles: 4}, //125->0x7D
		{OpName: op_ROR, AMName: am_ABX, Op: cpu.ROR, AddrMode: cpu.ABX, Cycles: 7}, //126->0x7E
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //127->0x7F
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //128->0x80
		{OpName: op_STA, AMName: am_IZX, Op: cpu.STA, AddrMode: cpu.IZX, Cycles: 6}, //129->0x81
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //130->0x82
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //131->0x83
		{OpName: op_STY, AMName: am_ZP0, Op: cpu.STY, AddrMode: cpu.ZP0, Cycles: 3}, //132->0x84
		{OpName: op_STA, AMName: am_ZP0, Op: cpu.STA, AddrMode: cpu.ZP0, Cycles: 3}, //133->0x85
		{OpName: op_STX, AMName: am_ZP0, Op: cpu.STX, AddrMode: cpu.ZP0, Cycles: 3}, //134->0x86
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 3}, //135->0x87
		{OpName: op_DEY, AMName: am_IMP, Op: cpu.DEY, AddrMode: cpu.IMP, Cycles: 2}, //136->0x88
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //137->0x89
		{OpName: op_TXA, AMName: am_IMP, Op: cpu.TXA, AddrMode: cpu.IMP, Cycles: 2}, //138->0x8A
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //139->0x8B
		{OpName: op_STY, AMName: am_ABS, Op: cpu.STY, AddrMode: cpu.ABS, Cycles: 4}, //140->0x8C
		{OpName: op_STA, AMName: am_ABS, Op: cpu.STA, AddrMode: cpu.ABS, Cycles: 4}, //141->0x8D
		{OpName: op_STX, AMName: am_ABS, Op: cpu.STX, AddrMode: cpu.ABS, Cycles: 4}, //142->0x8E
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 4}, //143->0x8F
		{OpName: op_BCC, AMName: am_REL, Op: cpu.BCC, AddrMode: cpu.REL, Cycles: 2}, //144->0x90
		{OpName: op_STA, AMName: am_IZY, Op: cpu.STA, AddrMode: cpu.IZY, Cycles: 6}, //145->0x91
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //146->0x92
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //147->0x93
		{OpName: op_STY, AMName: am_ZPX, Op: cpu.STY, AddrMode: cpu.ZPX, Cycles: 4}, //148->0x94
		{OpName: op_STA, AMName: am_ZPX, Op: cpu.STA, AddrMode: cpu.ZPX, Cycles: 4}, //149->0x95
		{OpName: op_STX, AMName: am_ZPY, Op: cpu.STX, AddrMode: cpu.ZPY, Cycles: 4}, //150->0x96
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 4}, //151->0x97
		{OpName: op_TYA, AMName: am_IMP, Op: cpu.TYA, AddrMode: cpu.IMP, Cycles: 2}, //152->0x98
		{OpName: op_STA, AMName: am_ABY, Op: cpu.STA, AddrMode: cpu.ABY, Cycles: 5}, //153->0x99
		{OpName: op_TXS, AMName: am_IMP, Op: cpu.TXS, AddrMode: cpu.IMP, Cycles: 2}, //154->0x9A
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 5}, //155->0x9B
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 5}, //156->0x9C
		{OpName: op_STA, AMName: am_ABX, Op: cpu.STA, AddrMode: cpu.ABX, Cycles: 5}, //157->0x9D
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 5}, //158->0x9E
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 5}, //159->0x9F
		{OpName: op_LDY, AMName: am_IMM, Op: cpu.LDY, AddrMode: cpu.IMM, Cycles: 2}, //160->0xA0
		{OpName: op_LDA, AMName: am_IZX, Op: cpu.LDA, AddrMode: cpu.IZX, Cycles: 6}, //161->0xA1
		{OpName: op_LDX, AMName: am_IMM, Op: cpu.LDX, AddrMode: cpu.IMM, Cycles: 2}, //162->0xA2
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //163->0xA3
		{OpName: op_LDY, AMName: am_ZP0, Op: cpu.LDY, AddrMode: cpu.ZP0, Cycles: 3}, //164->0xA4
		{OpName: op_LDA, AMName: am_ZP0, Op: cpu.LDA, AddrMode: cpu.ZP0, Cycles: 3}, //165->0xA5
		{OpName: op_LDX, AMName: am_ZP0, Op: cpu.LDX, AddrMode: cpu.ZP0, Cycles: 3}, //166->0xA6
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 3}, //167->0xA7
		{OpName: op_TAY, AMName: am_IMP, Op: cpu.TAY, AddrMode: cpu.IMP, Cycles: 2}, //168->0xA8
		{OpName: op_LDA, AMName: am_IMM, Op: cpu.LDA, AddrMode: cpu.IMM, Cycles: 2}, //169->0xA9
		{OpName: op_TAX, AMName: am_IMP, Op: cpu.TAX, AddrMode: cpu.IMP, Cycles: 2}, //170->0xAA
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //171->0xAB
		{OpName: op_LDY, AMName: am_ABS, Op: cpu.LDY, AddrMode: cpu.ABS, Cycles: 4}, //172->0xAC
		{OpName: op_LDA, AMName: am_ABS, Op: cpu.LDA, AddrMode: cpu.ABS, Cycles: 4}, //173->0xAD
		{OpName: op_LDX, AMName: am_ABS, Op: cpu.LDX, AddrMode: cpu.ABS, Cycles: 4}, //174->0xAE
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 4}, //175->0xAF
		{OpName: op_BCS, AMName: am_REL, Op: cpu.BCS, AddrMode: cpu.REL, Cycles: 2}, //176->0xB0
		{OpName: op_LDA, AMName: am_IZY, Op: cpu.LDA, AddrMode: cpu.IZY, Cycles: 5}, //177->0xB1
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //178->0xB2
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 5}, //179->0xB3
		{OpName: op_LDY, AMName: am_ZPX, Op: cpu.LDY, AddrMode: cpu.ZPX, Cycles: 4}, //180->0xB4
		{OpName: op_LDA, AMName: am_ZPX, Op: cpu.LDA, AddrMode: cpu.ZPX, Cycles: 4}, //181->0xB5
		{OpName: op_LDX, AMName: am_ZPY, Op: cpu.LDX, AddrMode: cpu.ZPY, Cycles: 4}, //182->0xB6
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 4}, //183->0xB7
		{OpName: op_CLV, AMName: am_IMP, Op: cpu.CLV, AddrMode: cpu.IMP, Cycles: 2}, //184->0xB8
		{OpName: op_LDA, AMName: am_ABY, Op: cpu.LDA, AddrMode: cpu.ABY, Cycles: 4}, //185->0xB9
		{OpName: op_TSX, AMName: am_IMP, Op: cpu.TSX, AddrMode: cpu.IMP, Cycles: 2}, //186->0xBA
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 4}, //187->0xBB
		{OpName: op_LDY, AMName: am_ABX, Op: cpu.LDY, AddrMode: cpu.ABX, Cycles: 4}, //188->0xBC
		{OpName: op_LDA, AMName: am_ABX, Op: cpu.LDA, AddrMode: cpu.ABX, Cycles: 4}, //189->0xBD
		{OpName: op_LDX, AMName: am_ABY, Op: cpu.LDX, AddrMode: cpu.ABY, Cycles: 4}, //190->0xBE
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 4}, //191->0xBF
		{OpName: op_CPY, AMName: am_IMM, Op: cpu.CPY, AddrMode: cpu.IMM, Cycles: 2}, //192->0xC0
		{OpName: op_CMP, AMName: am_IZX, Op: cpu.CMP, AddrMode: cpu.IZX, Cycles: 6}, //193->0xC1
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //194->0xC2
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //195->0xC3
		{OpName: op_CPY, AMName: am_ZP0, Op: cpu.CPY, AddrMode: cpu.ZP0, Cycles: 3}, //196->0xC4
		{OpName: op_CMP, AMName: am_ZP0, Op: cpu.CMP, AddrMode: cpu.ZP0, Cycles: 3}, //197->0xC5
		{OpName: op_DEC, AMName: am_ZP0, Op: cpu.DEC, AddrMode: cpu.ZP0, Cycles: 5}, //198->0xC6
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 5}, //199->0xC7
		{OpName: op_INY, AMName: am_IMP, Op: cpu.INY, AddrMode: cpu.IMP, Cycles: 2}, //200->0xC8
		{OpName: op_CMP, AMName: am_IMM, Op: cpu.CMP, AddrMode: cpu.IMM, Cycles: 2}, //201->0xC9
		{OpName: op_DEX, AMName: am_IMP, Op: cpu.DEX, AddrMode: cpu.IMP, Cycles: 2}, //202->0xCA
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //203->0xCB
		{OpName: op_CPY, AMName: am_ABS, Op: cpu.CPY, AddrMode: cpu.ABS, Cycles: 4}, //204->0xCC
		{OpName: op_CMP, AMName: am_ABS, Op: cpu.CMP, AddrMode: cpu.ABS, Cycles: 4}, //205->0xCD
		{OpName: op_DEC, AMName: am_ABS, Op: cpu.DEC, AddrMode: cpu.ABS, Cycles: 6}, //206->0xCE
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //207->0xCF
		{OpName: op_BNE, AMName: am_REL, Op: cpu.BNE, AddrMode: cpu.REL, Cycles: 2}, //208->0xD0
		{OpName: op_CMP, AMName: am_IZY, Op: cpu.CMP, AddrMode: cpu.IZY, Cycles: 5}, //209->0xD1
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //210->0xD2
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //211->0xD3
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //212->0xD4
		{OpName: op_CMP, AMName: am_ZPX, Op: cpu.CMP, AddrMode: cpu.ZPX, Cycles: 4}, //213->0xD5
		{OpName: op_DEC, AMName: am_ZPX, Op: cpu.DEC, AddrMode: cpu.ZPX, Cycles: 6}, //214->0xD6
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //215->0xD7
		{OpName: op_CLD, AMName: am_IMP, Op: cpu.CLD, AddrMode: cpu.IMP, Cycles: 2}, //216->0xD8
		{OpName: op_CMP, AMName: am_ABY, Op: cpu.CMP, AddrMode: cpu.ABY, Cycles: 4}, //217->0xD9
		{OpName: op_NOP, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //218->0xDA
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //219->0xDB
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //220->0xDC
		{OpName: op_CMP, AMName: am_ABX, Op: cpu.CMP, AddrMode: cpu.ABX, Cycles: 4}, //221->0xDD
		{OpName: op_DEC, AMName: am_ABX, Op: cpu.DEC, AddrMode: cpu.ABX, Cycles: 7}, //222->0xDE
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //223->0xDF
		{OpName: op_CPX, AMName: am_IMM, Op: cpu.CPX, AddrMode: cpu.IMM, Cycles: 2}, //224->0xE0
		{OpName: op_SBC, AMName: am_IZX, Op: cpu.SBC, AddrMode: cpu.IZX, Cycles: 6}, //225->0xE1
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //226->0xE2
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //227->0xE3
		{OpName: op_CPX, AMName: am_ZP0, Op: cpu.CPX, AddrMode: cpu.ZP0, Cycles: 3}, //228->0xE4
		{OpName: op_SBC, AMName: am_ZP0, Op: cpu.SBC, AddrMode: cpu.ZP0, Cycles: 3}, //229->0xE5
		{OpName: op_INC, AMName: am_ZP0, Op: cpu.INC, AddrMode: cpu.ZP0, Cycles: 5}, //230->0xE6
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 5}, //231->0xE7
		{OpName: op_INX, AMName: am_IMP, Op: cpu.INX, AddrMode: cpu.IMP, Cycles: 2}, //232->0xE8
		{OpName: op_SBC, AMName: am_IMM, Op: cpu.SBC, AddrMode: cpu.IMM, Cycles: 2}, //233->0xE9
		{OpName: op_NOP, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //234->0xEA
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.SBC, AddrMode: cpu.IMP, Cycles: 2}, //235->0xEB
		{OpName: op_CPX, AMName: am_ABS, Op: cpu.CPX, AddrMode: cpu.ABS, Cycles: 4}, //236->0xEC
		{OpName: op_SBC, AMName: am_ABS, Op: cpu.SBC, AddrMode: cpu.ABS, Cycles: 4}, //237->0xED
		{OpName: op_INC, AMName: am_ABS, Op: cpu.INC, AddrMode: cpu.ABS, Cycles: 6}, //238->0xEE
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //239->0xEF
		{OpName: op_BEQ, AMName: am_REL, Op: cpu.BEQ, AddrMode: cpu.REL, Cycles: 2}, //240->0xF0
		{OpName: op_SBC, AMName: am_IZY, Op: cpu.SBC, AddrMode: cpu.IZY, Cycles: 5}, //241->0xF1
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 2}, //242->0xF2
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 8}, //243->0xF3
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //244->0xF4
		{OpName: op_SBC, AMName: am_ZPX, Op: cpu.SBC, AddrMode: cpu.ZPX, Cycles: 4}, //245->0xF5
		{OpName: op_INC, AMName: am_ZPX, Op: cpu.INC, AddrMode: cpu.ZPX, Cycles: 6}, //246->0xF6
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 6}, //247->0xF7
		{OpName: op_SED, AMName: am_IMP, Op: cpu.SED, AddrMode: cpu.IMP, Cycles: 2}, //248->0xF8
		{OpName: op_SBC, AMName: am_ABY, Op: cpu.SBC, AddrMode: cpu.ABY, Cycles: 4}, //249->0xF9
		{OpName: op_NOP, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 2}, //250->0xFA
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //251->0xFB
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.NOP, AddrMode: cpu.IMP, Cycles: 4}, //252->0xFC
		{OpName: op_SBC, AMName: am_ABX, Op: cpu.SBC, AddrMode: cpu.ABX, Cycles: 4}, //253->0xFD
		{OpName: op_INC, AMName: am_ABX, Op: cpu.INC, AddrMode: cpu.ABX, Cycles: 7}, //254->0xFE
		{OpName: op_ILL, AMName: am_IMP, Op: cpu.ILL, AddrMode: cpu.IMP, Cycles: 7}, //255->0xFF
	}
}
