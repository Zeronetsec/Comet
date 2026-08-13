// https://github.com/Zeronetsec/Comet

package console

import (
    "os"
    "github.com/Zeronetsec/Comet/utils/invinput"
)

func CometConsole(input string) {
    args := os.Args
    if len(args) < 2 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    commands := map[string]Command{
        "--portscan": Portscan{},
        "--dirfuzzer": Dirfuzzer{},
        "--osint": Osint{},
        "--paramscan": Paramscan{},
        "--tracelink": Tracelink{},
        "--uwu": Uwu{},
        "--version": Version{},
        "--help": Helper{},
        "--header": Header{},
        "--subdomain": Subdomain{},
        "--hostsearch": HostSearch{},
        "--dnslookup": DNSLookup{},
        "--sqlscan": Sqlscan{},
        "--subtakeover": SubTakeover{},
        "--corsscan": CorsScan{},
    }

    if cmd, ok := commands[args[1]]; ok {
        cmd.Execute(args)
    } else {
        invinput.InvalidOption(args[1])
        os.Exit(1)
    }
}

// Copyright (c) 2026 Zeronetsec