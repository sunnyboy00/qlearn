package learn

type Action struct {
	Name      string
	Reward    float64
	NextState State
}
