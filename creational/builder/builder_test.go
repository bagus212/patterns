package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}

	manufacturingComplex.SetBuilder(carBuilder)

	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("wheels on car must be 4 and they were %d\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("structure on a car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("Seats on Car must be 5 and they wee %d\n", car.Seats)
	}

	busBuilder := &BusBuilder{}

	manufacturingComplex.SetBuilder(busBuilder)

	manufacturingComplex.Construct()

	bus := busBuilder.GetVehicle()

	if bus.Wheels != 8 {
		t.Errorf("wheels on bus must be 4 and they were %d\n", bus.Wheels)
	}

	if bus.Structure != "Bus" {
		t.Errorf("structure on bus must be 'Bus' and was %s\n", bus.Structure)
	}

	if bus.Seats != 30 {
		t.Errorf("seats pn bus must be 30 and they were %d\n", bus.Seats)
	}
}
