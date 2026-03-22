// Comet Framework

package console

import (
    "os"
    "comet/utils/invinput"
)

func CometConsole(input string) {
    args := os.Args
    if len(args) < 2 {
        invinput.Invalid()
        return
    }

    commands := map[string]Command{
        "--sysinfo": Sysinfo{},
        "--portscan": Portscan{},
        "--procinfo": Procinfo{},
        "--checkroot": Checkroot{},
        "--dumpstring": Dumpstring{},
        "--decode": Decode{},
        "--misconfind": Misconfind{},
        "--dirfuzzer": Dirfuzzer{},
        "--osint": Osint{},
        "--paramscan": Paramscan{},
        "--tracelink": Tracelink{},
        "--listcmd": Listcmd{},
        "--searchcmd": Searchcmd{},
        "--uwu": Uwu{},
        "--version": Version{},
        "--help": Helper{},
    }

    if cmd, ok := commands[args[1]]; ok {
        cmd.Execute(args)
    } else {
        invinput.Unknown(args[1])
    }
}

// Copyright (c) 2026 Zeronetsec