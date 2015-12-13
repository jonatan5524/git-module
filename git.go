// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"fmt"
	"strings"
	"time"
)

const _VERSION = "0.1.1"

func Version() string {
	return _VERSION
}

var (
	// Debug enables verbose logging on everything.
	// This should be false in case Gogs starts in SSH mode.
	Debug  = false
	Prefix = "[git-shell] "
)

func log(format string, args ...interface{}) {
	if !Debug {
		return
	}

	fmt.Print(Prefix)
	if len(args) == 0 {
		fmt.Println(format)
	} else {
		fmt.Printf(format+"\n", args...)
	}
}

var gitVersion string

// Version returns current Git version from shell.
func BinVersion() (string, error) {
	if len(gitVersion) > 0 {
		return gitVersion, nil
	}

	stdout, err := NewCommand("version").Run()
	if err != nil {
		return "", err
	}

	// Handle special case on Windows.
	i := strings.Index(stdout, "windows")
	if i >= 1 {
		gitVersion = stdout[:i-1]
		return gitVersion, nil
	}

	fields := strings.Fields(stdout)
	if len(fields) < 3 {
		return "", fmt.Errorf("not enough output: %s", stdout)
	}

	gitVersion = fields[2]
	return gitVersion, nil
}

func init() {
	BinVersion()
}

// Fsck verifies the connectivity and validity of the objects in the database
func Fsck(repoPath string, timeout time.Duration, args ...string) error {
	_, err := NewCommand("fsck").AddArguments(args...).RunInDirTimeout(timeout, repoPath)
	return err
}