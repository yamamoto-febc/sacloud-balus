package lib

import (
	"fmt"
	"strings"
)

type Option struct {
	ReceivePath   string
	ReceivePort   int
	ReceiveSecret string

	SendToken      string
	SendSecret     string
	TargetModuleID string

	Debug bool

	AccessToken       string
	AccessTokenSecret string
	Zones             []string
	TraceMode         bool
}

func NewOption() *Option {
	return &Option{}
}

func (o *Option) Validate() []error {
	var errors []error
	if o.AccessToken == "" {
		errors = append(errors, fmt.Errorf("[%s] is required", "sakuracloud-access-token"))
	}
	if o.AccessTokenSecret == "" {
		errors = append(errors, fmt.Errorf("[%s] is required", "sakuracloud-access-token-secret"))
	}

	if o.SendToken == "" {
		errors = append(errors, fmt.Errorf("[%s] is required", "send-token"))
	}
	if o.TargetModuleID == "" {
		errors = append(errors, fmt.Errorf("[%s] is required", "send-module-id"))
	}

	return errors
}

func (o *Option) PrintOptionValues() {
	fmt.Println("")
	fmt.Println("==================================================")
	fmt.Println("  Option values")
	fmt.Println("==================================================")
	fmt.Printf("  <Webhook receive settings>\n")
	fmt.Printf("    Path   : %s\n", o.ReceivePath)
	fmt.Printf("    Port   : %d\n", o.ReceivePort)
	fmt.Printf("    Secret : %s\n", o.ReceiveSecret)
	fmt.Println("")

	fmt.Printf("  <Webhook send settings>\n")
	fmt.Printf("    Token  : %s\n", o.SendToken)
	fmt.Printf("    Secret : %s\n", o.SendSecret)
	fmt.Printf("    Module : %s\n", o.TargetModuleID)
	fmt.Println("")

	fmt.Printf("  <SakuraCloud target settings>\n")
	fmt.Printf("    Zones  : [%s]\n", strings.Join(o.Zones, ","))
	fmt.Println("==================================================")
	fmt.Println("")
}
