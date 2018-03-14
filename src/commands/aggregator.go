package commands

import "gopkg.in/urfave/cli.v1"

type commandsAggregator struct {
	commands []cli.Command
}

func (c *commandsAggregator) add(command cli.Command) {
	c.commands = append(c.commands, command)
}

func (c *commandsAggregator) get() []cli.Command {
	return c.commands
}
