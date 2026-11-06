// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Help implements the 'help' command.
func Help(w io.Writer, args []string) {
	cmd := RootCommand
Args:
	for i, arg := range args {
		for _, sub := range cmd.Commands {
			if sub.Name() == arg {
				cmd = sub
				continue Args
			}
		}

		// helpSuccess is the help command using as many args as possible that would succeed.
		helpSuccess := CommandEnv.Exec + " help"
		if i > 0 {
			helpSuccess += " " + strings.Join(args[:i], " ")
		}
		fmt.Fprintf(os.Stderr, "%s help %s: unknown help topic. Run '%s'.\n", CommandEnv.Exec, strings.Join(args, " "), helpSuccess)
		SetExitStatus(2)
		Exit()
	}

	if len(cmd.Commands) > 0 {
		PrintUsage(os.Stdout, cmd)
	} else {
		if cmd.Runnable() {
			fmt.Fprintf(w, "usage: %s\n\n", cmd.UsageLine)
		}
		fmt.Fprintln(w, strings.TrimSpace(cmd.Long))
	}
}

// PrintUsage prints usage of cmd to w
func PrintUsage(w io.Writer, cmd *Command) {
	buildCommandText(cmd)
	bw := bufio.NewWriter(w)
	fmt.Fprintln(bw, strings.TrimSpace(cmd.Long))
	fmt.Fprintf(bw, "\n\nUsage:\n\n\t%s <command> [arguments]\n\nThe commands are:\n", cmd.UsageLine)
	wid := CommandEnv.CommandsWidth
	for _, c := range cmd.Commands {
		if c.Short != "" && (c.Runnable() || len(c.Commands) > 0) {
			fmt.Fprintf(bw, "\t%s %s\n", width(wid, c.Name()), c.Short)
		}
	}
	ln := cmd.LongName()
	if ln != "" {
		fmt.Fprintf(bw, "\nUse \"svc help %s <command>\" for more information about a command.\n", ln)
	} else {
		fmt.Fprintf(bw, "\nUse \"svc help <command>\" for more information about a command.\n")
	}
	bw.Flush()
}

// buildCommandText is a no-op: all template placeholders were flattened to
// static text at build-source time, so no text/template reflection is needed.
func buildCommandText(cmd *Command) {
	_ = cmd
}

func width(width int, value string) string {
	format := fmt.Sprintf("%%-%ds", width)
	return fmt.Sprintf(format, value)
}
