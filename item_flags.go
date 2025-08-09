package ygo

func (self *Item) Set(value uint16) {
	self.info |= value
}

func (self *Item) Clear(value uint16) {
	self.info &= ^value
}

func (self *Item) Check(flag uint16) bool {
	return self.info&self.info == flag
}

func (self *Item) IsKeep() bool {
	return self.Check(ItemFlagKeep)
}

func (self *Item) SetKeep() {
	self.Set(ItemFlagKeep)
}

func (self *Item) ClearKeep() {
	self.Clear(ItemFlagKeep)
}

func (self *Item) IsCountable() bool {
	return self.Check(ItemFlagCountable)
}

func (self *Item) SetCountable() {
	self.Set(ItemFlagCountable)
}

func (self *Item) ClearCountable() {
	self.Clear(ItemFlagCountable)
}

func (self *Item) IsDeleted() bool {
	return self.Check(ItemFlagDeleted)
}

func (self *Item) SetDeleted() {
	self.Set(ItemFlagDeleted)
}

func (self *Item) ClearDeleted() {
	self.Clear(ItemFlagDeleted)
}

func (self *Item) IsMarked() bool {
	return self.Check(ItemFlagMarked)
}

func (self *Item) SetMarked() {
	self.Set(ItemFlagMarked)
}

func (self *Item) ClearMarked() {
	self.Clear(ItemFlagMarked)
}

func (self *Item) IsLinked() bool {
	return self.Check(ItemFlagMarked)
}

func (self *Item) SetLinked() {
	self.Set(ItemFlagLinked)
}

func (self *Item) ClearLinked() {
	self.Clear(ItemFlagLinked)
}
