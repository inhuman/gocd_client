package commands

import (
	"gopkg.in/urfave/cli.v1"
)

func Get() []cli.Command {

	var aggregator commandsAggregator

	aggregator.add(pipelinesCommand())
	aggregator.add(packagesCommand())

	return aggregator.get()
}
