package jiracmd

import (
	"github.com/coryb/figtree"
	"github.com/coryb/oreo"
	"github.com/go-jira/jira"
	"github.com/go-jira/jira/jiracli"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func CmdCheckLoginRegistry() *jiracli.CommandRegistryEntry {
	opts := jiracli.CommonOptions{}
	return &jiracli.CommandRegistryEntry{
		"Check login into jira server",
		func(fig *figtree.FigTree, cmd *kingpin.CmdClause) error {
			jiracli.LoadConfigs(cmd, fig, &opts)
			return nil
		},
		func(o *oreo.Client, globals *jiracli.GlobalOptions) error {
			return CmdCheckLogin(o, globals, &opts)
		},
	}
}

// CmdCheckLogin will attempt to login into jira server
func CmdCheckLogin(o *oreo.Client, globals *jiracli.GlobalOptions, opts *jiracli.CommonOptions) error {
	if globals.AuthMethod() == "api-token" {
		log.Noticef("No need to login when using api-token authentication method")
		return nil
	}
	o2 := o.WithoutPostCallbacks()
	if c := jira.CheckSession(o2, globals.Endpoint.Value); c {
		log.Noticef("Has login")
	} else {
		log.Noticef("No login")
	}
	return nil
}
