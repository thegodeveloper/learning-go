package interfaces

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return data + " processed"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	data := "this data comes from somewhere"

	c.L.Process(data)
}

func implicit(show bool) {
	if show {
		c := Client{
			L: LogicProvider{},
		}
		c.Program()
	}
}
