package main

type Service struct {
	description string
	durationMonths int
	monthlyfee float64
}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(recur bool) float64 {
	if recur{
		return s.monthlyfee * float64(s.durationMonths)
	} else {
		return s.monthlyfee
	}
		
}