// Comet Framework

package procinfo

type Proc struct {
    PID int32
    User string
    TTY string
    Stat string
    CPU float64
    MEM float32
    Name string
}

// Copyright (c) 2026 Zeronetsec