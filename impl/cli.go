package impl

import (
	"github.com/urfave/cli/v2"
	"github.com/wiedzmin/gourmet/version"
)

const (
	_ = iota
	typeBuku
)

// CLI is the command line interface app object structure
type CLI struct {
	app *cli.App
}

// Run is the entry point to the CLI app
func (c *CLI) Run(args []string) error {
	return c.app.Run(args)
}

func CreateCLI() *CLI {
	c := &CLI{
		app: cli.NewApp(),
	}

	c.app.Name = version.Description
	c.app.Usage = version.Usage
	c.app.Version = version.Version()
	c.app.Commands = cli.Commands{
		{
			Name:   "import",
			Usage:  "import bookmarks from various sources",
			Action: c.importStub,
		},
	}
	return c
}

func (c *CLI) importStub(ctx *cli.Context) error {
	return importBookmarks(typeBuku)
}
