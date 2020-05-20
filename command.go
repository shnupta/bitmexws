package bitmexws

// Represents a BitMEX compatible websocket message
type Command struct {
	Op   string        `json:"op"`
	Args []interface{} `json:"args"`
}

// Creates a new BitMEX compatible websocket message
func CreateCommand(op string, args ...interface{}) *Command {
	a := make([]interface{}, len(args))
	for i, arg := range args {
		a[i] = arg
	}
	return &Command{Op: op, Args: args}
}

// Adds `arg` to the list of arguments of the command
func (c *Command) AddArg(arg interface{}) {
	if c.Args == nil {
		c.Args = make([]interface{}, 1)
	}
	c.Args = append(c.Args, arg)
}
