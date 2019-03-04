package char

type TrustOnce struct {
	Nope bool
}

func (c TrustOnce) Run(s Status) int64 {

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

	if cnt > 0 {
		c.Nope = true
		return 0
	} else {
		return 1
	}

}

func (c TrustOnce) Name() string {

	return "TrustOnce"
}
