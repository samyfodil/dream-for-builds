package kill

import (
	"github.com/taubyte/dreamland/cli/command"
	client "github.com/taubyte/dreamland/service"
	"github.com/taubyte/go-specs/common"
	"github.com/urfave/cli/v2"
)

func service(multiverse *client.Client) []*cli.Command {
	validServices := common.P2PStreamProtocols
	commands := make([]*cli.Command, len(validServices))

	for idx, _service := range validServices {
		c := &cli.Command{
			Name:   _service,
			Action: killService(_service, multiverse),
		}
		command.Universe0(c)
		commands[idx] = c
	}

	return commands
}

func killService(name string, multiverse *client.Client) cli.ActionFunc {
	return func(c *cli.Context) (err error) {
		universe := multiverse.Universe(c.String("universe"))
		return universe.KillService(name)
	}
}
