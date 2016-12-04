// sacloud-balus - A tool of destroy all resource on SakuraCloud with magical spell `balus`
package main

import (
	"fmt"
	"github.com/yamamoto-febc/sacloud-balus/lib"
	"github.com/yamamoto-febc/sacloud-balus/version"
	"gopkg.in/urfave/cli.v2"
	"io"
	"os"
	"strings"
)

var (
	appName              = "sacloud-balus"
	appUsage             = "A tool of destroy all resource on SakuraCloud with magical spell `balus`"
	appCopyright         = "Copyright (C) 2016 Kazumichi Yamamoto."
	applHelpTextTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.HelpName}} [options]

REQUIRED PARAMETERS:
   {{range .RequiredFlags}}{{.}}
   {{end}}
OPTIONS:
   {{range .NormalFlags}}{{.}}
   {{end}}
************* ADVANCED OPTIONS **************

   FOR DEBUG:
      {{range .ForDeveloperFlags}}{{.}}
      {{end}}

*************************************************
VERSION:
   {{.Version}}

{{.Copyright}}
`

	requiredFlagNames = []string{
		"access-token",
		"access-secret",
		"send-token",
		"send-module-id",
	}

	forDeveloperFlagNames = []string{
		"sakuracloud-trace-mode",
		"debug",
	}
)

func main() {

	cli.AppHelpTemplate = applHelpTextTemplate
	app := &cli.App{}
	option := lib.NewOption()

	app.Name = appName
	app.Usage = appUsage
	app.HelpName = appName
	app.Copyright = appCopyright

	app.Flags = cliFlags(option)
	app.Action = cliCommand(option)
	app.Version = version.FullVersion()

	originalHelpPrinter := cli.HelpPrinter
	cli.HelpPrinter = func(w io.Writer, templ string, d interface{}) {
		app := d.(*cli.App)
		data := newHelpData(app)
		originalHelpPrinter(w, templ, data)
	}

	app.Run(os.Args)
}

type helpData struct {
	*cli.App
	RequiredFlags     []cli.Flag
	NormalFlags       []cli.Flag
	ForDeveloperFlags []cli.Flag
}

func newHelpData(app *cli.App) interface{} {
	data := &helpData{App: app}

	for _, f := range app.VisibleFlags() {
		if isExistsFlag(requiredFlagNames, f) {
			data.RequiredFlags = append(data.RequiredFlags, f)
		} else if isExistsFlag(forDeveloperFlagNames, f) {
			data.ForDeveloperFlags = append(data.ForDeveloperFlags, f)
		} else {
			data.NormalFlags = append(data.NormalFlags, f)
		}
	}

	return data
}

func cliFlags(option *lib.Option) []cli.Flag {

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "access-token",
			Aliases:     []string{"sakuracloud-access-token"},
			Usage:       "API Token of SakuraCloud",
			EnvVars:     []string{"SAKURACLOUD_ACCESS_TOKEN"},
			DefaultText: "none",
			Destination: &option.AccessToken,
		},
		&cli.StringFlag{
			Name:        "access-secret",
			Aliases:     []string{"sakuracloud-access-token-secret"},
			Usage:       "API Secret of SakuraCloud",
			EnvVars:     []string{"SAKURACLOUD_ACCESS_TOKEN_SECRET"},
			DefaultText: "none",
			Destination: &option.AccessTokenSecret,
		},
		&cli.StringSliceFlag{
			Name:    "zones",
			Aliases: []string{"sakuracloud-zones"},
			Usage:   "Target zone list of SakuraCloud",
			EnvVars: []string{"SAKURACLOUD_ZONES"},
			Value:   cli.NewStringSlice("is1b", "tk1a"),
		},
		&cli.BoolFlag{
			Name:        "sakuracloud-trace-mode",
			Usage:       "Flag of SakuraCloud debug-mode",
			EnvVars:     []string{"SAKURACLOUD_TRACE_MODE"},
			Destination: &option.TraceMode,
			Value:       false,
		},
		&cli.IntFlag{
			Name:        "receive-port",
			Usage:       "Number of web server port for receive webhook from the Sakura IoT Platform",
			EnvVars:     []string{"SACLOUD_BALUS_RECEIVE_PORT"},
			Destination: &option.ReceivePort,
			Value:       8080,
		},
		&cli.StringFlag{
			Name:        "receive-path",
			EnvVars:     []string{"SACLOUD_BALUS_RECEIVE_PATH"},
			DefaultText: "/",
			Value:       "/",
			Destination: &option.ReceivePath,
			Usage:       "Path for receive webhook from the Sakura IoT Platform",
		},
		&cli.StringFlag{
			Name:        "receive-secret",
			EnvVars:     []string{"SACLOUD_BALUS_RECEIVE_SECRET"},
			DefaultText: "",
			Destination: &option.ReceiveSecret,
			Usage:       "Secret for receive webhook from the Sakura IoT Platform",
		},
		&cli.StringFlag{
			Name:        "send-token",
			EnvVars:     []string{"SACLOUD_BALUS_SEND_TOKEN"},
			DefaultText: "",
			Destination: &option.SendToken,
			Usage:       "Token for send webhook to the Sakura IoT Platform",
		},
		&cli.StringFlag{
			Name:        "send-secret",
			EnvVars:     []string{"SACLOUD_BALUS_SEND_SECRET"},
			DefaultText: "",
			Destination: &option.SendSecret,
			Usage:       "Secret for send webhook to the Sakura IoT Platform",
		},
		&cli.StringFlag{
			Name:        "send-module-id",
			EnvVars:     []string{"SACLOUD_BALUS_SEND_MODULE_ID"},
			DefaultText: "",
			Destination: &option.TargetModuleID,
			Usage:       "ID of target module that to send webhook on the Sakura IoT Platform",
		},
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Flag of enable DEBUG log",
			EnvVars:     []string{"SACLOUD_BALUS_DEBUG"},
			Destination: &option.Debug,
			Value:       false,
		},
	}

}

func cliCommand(option *lib.Option) func(c *cli.Context) error {
	return func(c *cli.Context) error {

		option.Zones = c.StringSlice("sakuracloud-zones")
		errors := option.Validate()
		if len(errors) != 0 {
			return flattenErrors(errors)
		}

		return lib.Start(option)

	}
}

func flattenErrors(errors []error) error {
	var list = make([]string, 0)
	for _, str := range errors {
		list = append(list, str.Error())
	}
	return fmt.Errorf(strings.Join(list, "\n"))
}

func isExistsFlag(source []string, target cli.Flag) bool {
	for _, s := range source {
		if s == target.Names()[0] {
			return true
		}
	}
	return false
}
