package char

type AlwaysYes struct{}

func (c AlwaysYes) Run(s Status) int64 {

	return 1

}

func (c AlwaysYes) Name() string {

	return "AlwaysYes"
}
