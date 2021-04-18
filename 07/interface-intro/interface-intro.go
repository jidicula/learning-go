package main

type LogicProvider struct{}

// Process processes the data.
func (lp LogicProvider) Process(data string) string {
	// some biz logic here
	return ""
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
	// only client knows about Logic interface, not provider LogicProvider.
	// Allows implementing a new LogicProvider in the future that can still
	// meet Client's need.
}

// Program does some stuff.
func (c Client) Program() {
	var data string
	c.L.Process(data)
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
