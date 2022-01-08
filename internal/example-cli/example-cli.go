// minimal example CLI used for binary size checking

package main

import (
	"github.com/wa-lang/cli"
)

func main() {
	(&cli.App{}).Run([]string{})
}
