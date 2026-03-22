// Comet Framework

package procinfo

import (
    "fmt"
    "sort"
    "github.com/shirou/gopsutil/v3/process"
    "comet/utils/color"
)

func ShowProcInfo() {
    procs, _ := process.Processes()
    var list []Proc

    ttyMap := getTTYMap()

    for _, p := range procs {
        name, _ := p.Name()
        cpu, _ := p.CPUPercent()
        mem, _ := p.MemoryPercent()
        user, _ := p.Username()

        tty := "-"
        if t, ok := ttyMap[p.Pid]; ok {
            tty = t
        }

        status, _ := p.Status()
        stat := "-"
        if len(status) > 0 {
            stat = status[0]
        }

        list = append(list, Proc{
            PID: p.Pid,
            User: user,
            TTY: tty,
            Stat: stat,
            CPU: cpu,
            MEM: mem,
            Name: name,
        })
    }

    sort.Slice(list, func(i, j int) bool {
        return list[i].CPU > list[j].CPU
    })

    fmt.Printf(
        "%s[*] %sScanning process running\n",
        color.B, color.N,
    )

    for i, p := range list {
        if i >= 50 {
            break
        }

        fmt.Println()
        fmt.Printf(
            "%s[*] %sProcess %d%s:\n",
            color.B, color.N, i+1, color.N,
        )

        fmt.Printf(
            "    %s├── %sName: %s%s%s\n",
            color.DG, color.N, color.GG, p.Name, color.N,
        )

        fmt.Printf(
            "    %s├── %sPID: %s%d%s\n",
            color.DG, color.N, color.GG, p.PID, color.N,
        )

        fmt.Printf(
            "    %s├── %sUser: %s%s%s\n",
            color.DG, color.N, color.GG, p.User, color.N,
        )

        fmt.Printf(
            "    %s├── %sTTY: %s%s%s\n",
            color.DG, color.N, color.GG, p.TTY, color.N,
        )

        fmt.Printf(
            "    %s├── %sStat: %s%s%s\n",
            color.DG, color.N, color.GG, p.Stat, color.N,
        )

        fmt.Printf(
            "    %s├── %sCPU%%: %s%.2f%s\n",
            color.DG, color.N, color.GG, p.CPU, color.N,
        )

        fmt.Printf(
            "    %s└── %sMEM%%: %s%.2f%s\n",
            color.DG, color.N, color.GG, p.MEM, color.N,
        )
    }

    total := len(list)
    fmt.Printf(
        "\n%s[*] %sTotal process: %s%d%s\n",
        color.B, color.N, color.GG, total, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec