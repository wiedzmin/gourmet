package impl

import "github.com/urfave/cli/v2"

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

	// TODO: abstract away
	c.app.Name = "Gourmet bookmarks manager"
	c.app.Version = "0.1" // TODO: use semver
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
