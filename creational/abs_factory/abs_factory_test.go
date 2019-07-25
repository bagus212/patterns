package abs_factory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	motorBikeVehicle, err := motorbikeF.NewVehicle(SportMotorBikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("motorbike vehicle has %d weels\n", motorBikeVehicle.NumWeels())

	sportBike, ok := motorBikeVehicle.(*SportMotorBike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Sport motorbike has type %d\n", sportBike.GetMotorBikeType())
}

func TestCarFactory(t *testing.T) {
	carF, err := GetVehicleFactory(CarFactoryType)
	if err != nil {
		t.Fatal(err)
	}
	carVehicle, err := carF.NewVehicle(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("car has %d weels\n", carVehicle.NumWeels())

	luxuryCar, ok := carVehicle.(*LuxuryCar)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Luxury car has  %d doors\n", luxuryCar.NumDoors())
}
