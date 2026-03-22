// Comet Framework

package main

import (
    "os"
    "strings"
    "comet/console"
)

func main() {
    args := os.Args[1:]
    input := strings.Join(args, " ")
    console.CometConsole(input)
}

// Copyright (c) 2026 Zeronetsec