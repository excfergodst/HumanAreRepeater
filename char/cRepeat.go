package char

type Repeat struct{}

func (c Repeat) Run(s Status) int64 {

	if len(s.MyStep) == 0 {
		return 1
	}

	return s.OppeLastStep

}

func (c Repeat) Name() string {

	return "Repeat"
}
