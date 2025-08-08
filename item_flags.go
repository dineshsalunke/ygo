package ygo

type ItemFlags uint16

func (self ItemFlags) Set(value ItemFlags) {
	self |= value
}

func (self ItemFlags) Clear(value ItemFlags) {
	self &= ^value
}

func (self ItemFlags) Check(flag ItemFlags) bool {
	return self&self == flag
}

func (self ItemFlags) IsKeep() bool {
	return self.Check(ItemFlagKeep)
}

func (self ItemFlags) SetKeep() {
	self.Set(ItemFlagKeep)
}

func (self ItemFlags) ClearKeep() {
	self.Clear(ItemFlagKeep)
}

func (self ItemFlags) IsCountable() bool {
	return self.Check(ItemFlagCountable)
}

func (self ItemFlags) SetCountable() {
	self.Set(ItemFlagCountable)
}

func (self ItemFlags) ClearCountable() {
	self.Clear(ItemFlagCountable)
}

func (self ItemFlags) IsDeleted() bool {
	return self.Check(ItemFlagDeleted)
}

func (self ItemFlags) SetDeleted() {
	self.Set(ItemFlagDeleted)
}

func (self ItemFlags) ClearDeleted() {
	self.Clear(ItemFlagDeleted)
}

func (self ItemFlags) IsMarked() bool {
	return self.Check(ItemFlagMarked)
}

func (self ItemFlags) SetMarked() {
	self.Set(ItemFlagMarked)
}

func (self ItemFlags) ClearMarked() {
	self.Clear(ItemFlagMarked)
}

func (self ItemFlags) IsLinked() bool {
	return self.Check(ItemFlagMarked)
}

func (self ItemFlags) SetLinked() {
	self.Set(ItemFlagLinked)
}

func (self ItemFlags) ClearLinked() {
	self.Clear(ItemFlagLinked)
}
