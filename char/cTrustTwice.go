package char

type TrustTwice struct {
	Nope bool
}

func (c TrustTwice) Run(s Status) int64 {

	if len(s.MyStep) == 0 {
		c.Nope = false
	}

	if c.Nope == true {
		return 0
	}

	cnt := 0

	for _, v := range s.OppoStep {

		if v == 0 {
			cnt++
		}
	}

	if cnt > 1 {
		c.Nope = true
		return 0
	} else {
		return 1
	}

}

func (c TrustTwice) Name() string {

	return "TrustTwice"
}
