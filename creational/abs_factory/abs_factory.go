package abs_factory

import (
	"errors"
	"fmt"
)

type Vehicle interface {
	NumWeels() int
	NumSeats() int
}

type Car interface {
	NumDoors() int
}

type Motorbike interface {
	GetMotorBikeType() int
}

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

type CarFactory struct{}

func (cf *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of Type %d not recognized", v))
	}
}

type LuxuryCar struct{}

func (*LuxuryCar) NumWeels() int {
	return 4
}

func (*LuxuryCar) NumSeats() int {
	return 4
}

func (*LuxuryCar) NumDoors() int {
	return 5
}

type FamilyCar struct{}

func (*FamilyCar) NumWeels() int {
	return 5
}

func (*FamilyCar) NumSeats() int {
	return 4
}

func (*FamilyCar) NumDoors() int {
	return 5
}

const (
	SportMotorBikeType  = 1
	CruiseMotorBikeType = 2
)

type MotorbikeFactory struct{}

func (mf *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorBikeType:
		return new(SportMotorBike), nil
	case CruiseMotorBikeType:
		return new(CruiseMotorBike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of Type %d not recognized", v))
	}
}

type SportMotorBike struct{}

func (*SportMotorBike) NumWeels() int {
	return 2
}

func (*SportMotorBike) NumSeats() int {
	return 1
}

func (*SportMotorBike) GetMotorBikeType() int {
	return SportMotorBikeType
}

type CruiseMotorBike struct{}

func (*CruiseMotorBike) NumWeels() int {
	return 2
}

func (*CruiseMotorBike) NumSeats() int {
	return 2
}

func (*CruiseMotorBike) GetMotorBikeType() int {
	return CruiseMotorBikeType
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle Factory of Type %d not recognized", f))
	}
}
