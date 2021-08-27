/*
Package flag implements command-line flag parsing.

Usage

Define flags using flag.String(), Bool(), Int(), etc.

This declares an integer flag, -n, stored in the pointer nFlag, with type *int:
	import "flag"
	var nFlag = flag.Int("n", 1234, "help message for flag n")
if you like, you can bind the flag to variable using the Var functions.
	var flagvar int
	func init() {
		flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	}
Or you can create custom flags that satisfy the Value interface (with
pointer receivers) and couple them to flag parsing by
	flag.Var(&flagVal, "name", "help message for name")
For such flags, the default value is just the initial value of the variable.

After all flags are defined, call
	flag.Parse()
to parse the command line into the defined flags.

Flags may then be used directly. If you're using the flags themselves,
they are all pointers; if you bind to variables, they're values.
	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)

After parsing, the arguments following the flags are available sa the
slice flag.Args() or individually as flag.Arg(i).
The arguments are indexed from 0 through flag.NArg()-1.

Command line flag syntax

The following forms are permitted:

	-flag
	-flag=x
	-flag x // non-boolean flags only
One or two minus signs may be used; they are equivalent.
The last form is not permitted for boolean flags because the
meaning of the command
	cmd -x *
where * is a Unix shell wildcard, will change if there is a file
called 0, false, etc. You must use the -flag=false form to turn
off a boolean flag.

Flag parsing stops just before the first non-flag argument
("-" is a non-flag argument) or after the terminator "--".

Integer flags accept 1234, 0664, 0x1234 and my be negative.
Boolean flags may be:
	1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
Duration flags accept any input valid for time.ParseDuration.

The default set of command-line flags is controlled by
top-level functions. The FlagSet type allows one to define
independent sets of flags, such as ti implement subcommands
in a command-line interface. The methods of FlagSet are
analogous to the top-level functions for the command-line
flag set.
*/

/*
# core variables
// CommandLine is the default set of command-line flags, parsed from os.Args.
// The top-level functions such as BoolVar, Arg, and so on are wrappers for the
// methods of CommandLine.
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)


# core struct
// A FlagSet represents a set of defined flags. The zero value of a FlagSet
// has no name and has ContinueOnError error handling.
//
// Flag names must be unique within a FlagSet. An attempt to define a flag whose
// name is already in use will cause a panic.
type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler. What happens after Usage is called depends
	// on the ErrorHandling setting; for the command line, this defaults
	// to ExitOnError, which exits the program after calling Usage.
	Usage func()

	name          string
	parsed        bool
	actual        map[string]*Flag
	formal        map[string]*Flag
	args          []string // arguments after flags
	errorHandling ErrorHandling
	output        io.Writer // nil means stderr; use Output() accessor
}
*/
package main

import (
	"flag"
	"fmt"
)

type MyString string

func newMyString(val string, p *string) *MyString {
	*p = val
	return (*MyString)(p)
}

func (m *MyString) Set(val string) error {
	*m = MyString(val)
	return nil
}

func (m *MyString) String() string {
	return string(*m)
}

func main() {
	var name string
	var name1 MyString
	var boolFlag bool
	flag.StringVar(&name, "name", "", "help message for name")
	flag.Var(&name1, "name1", "help message for name1")
	flag.BoolVar(&boolFlag, "boolflag", false, "help message for boolflag")
	flag.Parse()

	fmt.Println(name1)
	fmt.Printf("%t", boolFlag)
}
