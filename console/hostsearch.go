// https://github.com/Zeronetsec/Comet

package console

import (
    "os"
    "github.com/Zeronetsec/Comet/utils/invinput"
    "github.com/Zeronetsec/Comet/module/hostsearch"
)

type HostSearch struct{}
func (c HostSearch) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    targetDomain := args[2]
    hostsearch.Scan(targetDomain)
}

// Copyright (c) 2026 Zeronetsec