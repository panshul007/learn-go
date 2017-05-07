package appliances

type Fridge struct {
  typeName string
}

// implementing the appliance interface by implementing the interface methods.

func (fr *Fridge) Start() {
  fr.typeName = " Fridge "
}

func (fr *Fridge) GetPurpose() string {
  return "I am a " + fr.typeName + " I cool stuff down!!"
}
