package ppu

func (ppu *Ricoh2c02) SetVideoReceiver(receiver VideoReceiver) {
	ppu.receiver = receiver
}

func (ppu *Ricoh2c02) SetDebugReceiver(debugReceiver DebugReceiver) {
	ppu.debugReceiver = debugReceiver
}
