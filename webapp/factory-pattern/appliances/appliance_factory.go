package appliances

import "errors"

// main interface used to describe appliance
type Appliance interface {
  Start()
  GetPurpose() string
}

// Create enum for appliance type
const (
  STOVE = iota // initializes index value of enum as 0
  FRIDGE
)

// factory function
func CreateAppliance(myType int) (Appliance, error) {
  switch myType {
  case STOVE:
    return new(Stove), nil
  case FRIDGE:
    return new(Fridge), nil

  default:
    return nil, errors.New("Invalid Appliance Type")
  }
}

