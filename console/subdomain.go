// https://github.com/Zeronetsec/Comet

package console

import (
    "os"
    "github.com/Zeronetsec/Comet/utils/invinput"
    "github.com/Zeronetsec/Comet/module/subdomain"
)

type Subdomain struct{}
func (c Subdomain) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    targetDomain := args[2]
    subdomain.Find(targetDomain)
}

// Copyright (c) 2026 Zeronetsec