package composite

type Trainer interface {
	Train()
}

type Swimmer interface {
	Swim()
}

type Eater interface {
	Eat()
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

type SharkB struct {
	Eater
	Swimmer
}
