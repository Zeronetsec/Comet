// https://github.com/Zeronetsec/Comet

package console

import (
    "time"
    "fmt"
    "github.com/Zeronetsec/Comet/utils/cursor"
    "github.com/Zeronetsec/Comet/module/uwu"
)

type Uwu struct{}
func (c Uwu) Execute(args []string) {
    cursor.Hide()
    uwu.Nyan(5 * time.Second)
    cursor.Visible()

    fmt.Println()
}

// Copyright (c) 2026 Zeronetsec