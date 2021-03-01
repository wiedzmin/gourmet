package impl

import (
	"github.com/urfave/cli/v2"
	"github.com/wiedzmin/gourmet/version"
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
			Name:  "import",
			Usage: "import bookmarks from various sources",
			Subcommands: []*cli.Command{
				{
					Name:   "buku",
					Usage:  "import bookmarks from Buku",
					Action: c.importBuku,
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "db-file",
							Aliases:  []string{"d"},
							Usage:    "Buku database file",
							Required: false,
						},
					},
				},
			},
		},
	}
	return c
}

func (c *CLI) importBuku(ctx *cli.Context) error {
	dbFile := ctx.String("db-file")
	if dbFile == "" {
		dbFile = getDefaultBukuDatabase()
	}
	return importBukuDB(dbFile)
}
