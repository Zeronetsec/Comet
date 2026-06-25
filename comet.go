// https://github.com/Zeronetsec/Comet

package main

import (
    "os"
    "strings"
    "github.com/Zeronetsec/Comet/console"
)

func main() {
    args := os.Args[1:]
    input := strings.Join(args, " ")
    console.CometConsole(input)
}

// Copyright (c) 2026 Zeronetsec