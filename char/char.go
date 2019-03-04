package char

type Status struct {
	MyStep       []int64
	OppoStep     []int64
	MyPoint      int64
	OppoPoint    int64
	MyLastStep   int64
	OppeLastStep int64
}

type Char interface {
	Run(Status) int64
	Name() string
}

func ListInit() (list []Char) {

	list = append(list, Random{})
	list = append(list, Repeat{})
	list = append(list, EightyYesTwentyNo{})
	list = append(list, AlwaysNo{})
	list = append(list, AlwaysYes{})
	list = append(list, NoFirstIfTrustYesAlways{})
	list = append(list, TrustOnce{})
	list = append(list, TrustTwice{})
	list = append(list, IfLessThenNo{})

	return list
}

func (s *Status) RecordStep(my int64, oppo int64) {

	s.MyStep = append(s.MyStep, my)
	s.OppoStep = append(s.OppoStep, oppo)

	s.MyLastStep = my
	s.OppeLastStep = oppo

}
