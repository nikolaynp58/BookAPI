// Code generated by goa v3.19.1, DO NOT EDIT.
//
// books HTTP client CLI support package
//
// Command:
// $ goa gen bookAPI/design

package cli

import (
	bookc "bookAPI/gen/http/book/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `book (create|show|update|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` book create --body '{
      "author": "In optio dolor sed quo porro.",
      "cover_url": "Natus magni laborum.",
      "published_at": "2023-01-01",
      "title": "Omnis molestiae sed."
   }'` + "\n" +
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
) (goa.Endpoint, any, error) {
	var (
		bookFlags = flag.NewFlagSet("book", flag.ContinueOnError)

		bookCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		bookCreateBodyFlag = bookCreateFlags.String("body", "REQUIRED", "")

		bookShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		bookShowIDFlag = bookShowFlags.String("id", "REQUIRED", "ID of the book")

		bookUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		bookUpdateBodyFlag = bookUpdateFlags.String("body", "REQUIRED", "")
		bookUpdateIDFlag   = bookUpdateFlags.String("id", "REQUIRED", "ID of the book")

		bookDeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		bookDeleteIDFlag = bookDeleteFlags.String("id", "REQUIRED", "ID of the book")
	)
	bookFlags.Usage = bookUsage
	bookCreateFlags.Usage = bookCreateUsage
	bookShowFlags.Usage = bookShowUsage
	bookUpdateFlags.Usage = bookUpdateUsage
	bookDeleteFlags.Usage = bookDeleteUsage

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
		case "book":
			svcf = bookFlags
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
		case "book":
			switch epn {
			case "create":
				epf = bookCreateFlags

			case "show":
				epf = bookShowFlags

			case "update":
				epf = bookUpdateFlags

			case "delete":
				epf = bookDeleteFlags

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
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "book":
			c := bookc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = bookc.BuildCreatePayload(*bookCreateBodyFlag)
			case "show":
				endpoint = c.Show()
				data, err = bookc.BuildShowPayload(*bookShowIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = bookc.BuildUpdatePayload(*bookUpdateBodyFlag, *bookUpdateIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = bookc.BuildDeletePayload(*bookDeleteIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// bookUsage displays the usage of the book command and its subcommands.
func bookUsage() {
	fmt.Fprintf(os.Stderr, `Service for managing books
Usage:
    %[1]s [globalflags] book COMMAND [flags]

COMMAND:
    create: Create implements create.
    show: Show implements show.
    update: Update implements update.
    delete: Delete implements delete.

Additional help:
    %[1]s book COMMAND --help
`, os.Args[0])
}
func bookCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] book create -body JSON

Create implements create.
    -body JSON: 

Example:
    %[1]s book create --body '{
      "author": "In optio dolor sed quo porro.",
      "cover_url": "Natus magni laborum.",
      "published_at": "2023-01-01",
      "title": "Omnis molestiae sed."
   }'
`, os.Args[0])
}

func bookShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] book show -id STRING

Show implements show.
    -id STRING: ID of the book

Example:
    %[1]s book show --id "Ut rerum aliquid."
`, os.Args[0])
}

func bookUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] book update -body JSON -id STRING

Update implements update.
    -body JSON: 
    -id STRING: ID of the book

Example:
    %[1]s book update --body '{
      "book": {
         "author": "Eos consequuntur tempore.",
         "cover_url": "Eum maiores maxime.",
         "published_at": "2023-01-01",
         "title": "At nesciunt deserunt."
      }
   }' --id "Non dolores quasi saepe sunt est dolor."
`, os.Args[0])
}

func bookDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] book delete -id STRING

Delete implements delete.
    -id STRING: ID of the book

Example:
    %[1]s book delete --id "Repellendus itaque et."
`, os.Args[0])
}
