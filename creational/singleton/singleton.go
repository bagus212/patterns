package singleton

type singleton struct {
	count int
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
