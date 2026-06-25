// https://github.com/Zeronetsec/Comet

package console

import (
    "github.com/Zeronetsec/Comet/module/version"
)

type Version struct{}
func (c Version) Execute(args []string) {
    version.CometVersion()
}

// Copyright (c) 2026 Zeronetsec