// Code generated by goa v3.2.5, DO NOT EDIT.
//
// members HTTP client CLI support package
//
// Command:
// $ goa gen members/design

package cli

import (
	"flag"
	"fmt"
	membersc "members/gen/http/members/client"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `members add
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` members add --a 7137870548529157984 --b 3412226314788822956` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		membersFlags = flag.NewFlagSet("members", flag.ContinueOnError)

		membersAddFlags = flag.NewFlagSet("add", flag.ExitOnError)
		membersAddAFlag = membersAddFlags.String("a", "REQUIRED", "Left operand")
		membersAddBFlag = membersAddFlags.String("b", "REQUIRED", "Right operand")
	)
	membersFlags.Usage = membersUsage
	membersAddFlags.Usage = membersAddUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "members":
			svcf = membersFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "members":
			switch epn {
			case "add":
				epf = membersAddFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "members":
			c := membersc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = membersc.BuildAddPayload(*membersAddAFlag, *membersAddBFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// membersUsage displays the usage of the members command and its subcommands.
func membersUsage() {
	fmt.Fprintf(os.Stderr, `メンバーの一覧を取得する
Usage:
    %s [globalflags] members COMMAND [flags]

COMMAND:
    add: Add implements add.

Additional help:
    %s members COMMAND --help
`, os.Args[0], os.Args[0])
}
func membersAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] members add -a INT -b INT

Add implements add.
    -a INT: Left operand
    -b INT: Right operand

Example:
    `+os.Args[0]+` members add --a 7137870548529157984 --b 3412226314788822956
`, os.Args[0])
}