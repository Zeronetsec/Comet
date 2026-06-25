// https://github.com/Zeronetsec/Comet

package console

import (
    "github.com/Zeronetsec/Comet/module/helper"
)

type Helper struct{}
func (c Helper) Execute(args []string) {
    helper.CometHelper()
}

// Copyright (c) 2026 Zeronetsec