package char

type IfLessThenNo struct{}

func (c IfLessThenNo) Run(s Status) int64 {

	if s.MyPoint < s.OppoPoint {
		return 0
	}

	return 1

}

func (c IfLessThenNo) Name() string {

	return "IfLessThenNo"
}
