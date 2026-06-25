// https://github.com/Zeronetsec/Comet

package console

import (
    "os"
    "github.com/Zeronetsec/Comet/utils/invinput"
    "github.com/Zeronetsec/Comet/module/dnslookup"
)

type DNSLookup struct{}
func (c DNSLookup) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    targetDomain := args[2]
    dnslookup.Scan(targetDomain)
}

// Copyright (c) 2026 Zeronetsec