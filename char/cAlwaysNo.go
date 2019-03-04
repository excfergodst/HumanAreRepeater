package char

type AlwaysNo struct{}

func (c AlwaysNo) Run(s Status) int64 {

	return 0

}

func (c AlwaysNo) Name() string {

	return "AlwaysNo"
}
