package char

type NoFirstIfTrustYesAlways struct {
	always bool
}

func (c NoFirstIfTrustYesAlways) Run(s Status) int64 {

	if len(s.MyStep) == 0 {
		c.always = false
		return 0
	}

	if c.always == true {
		return 1
	}

	if s.MyLastStep == 0 && s.OppeLastStep == 1 {
		c.always = true
		return 1
	}

	return 0

}

func (c NoFirstIfTrustYesAlways) Name() string {

	return "NoFirstIfTrustYesAlways"
}
