<!-- Comet Framework -->

[![version](https://img.shields.io/badge/Comet-Version%201.0-blue.svg?maxAge=259200)]()
[![gover](https://img.shields.io/badge/Go-Version%201.26.1-blue.svg)]()
[![os](https://img.shields.io/badge/Supported%20OS-Linux-blue.svg)]()
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

# Comet Framework
Comet is a security framework focused on reconnaissance, built for high performance. <br>
It covers a wide range of recon techniques, including fuzzing, OSINT, crawling, and more.

## Features
- Massive username enumeration across 700+ domains
- High-performance directory fuzzing
- TCP port scanning
- HTML link crawling (hyperlink extraction)
- Parameter discovery via Wayback Machine
- And more

## Disclaimer
This project is provided for educational and authorized security research purposes only. <br>
The author is not responsible for any misuse or damage caused by this tool. <br>
Please read the
[DISCLAIMER](https://github.com/Zeronetsec/Comet/blob/main/DISCLAIMER.md)
file for full terms.

## Installation
```bash
git clone https://github.com/Zeronetsec/Comet.git
cd Comet
chmod +x install.sh
./install.sh

# for backup
./install.sh --backup
```

## Usage
```bash
comet --dirfuzzer <url> [--wordlist <wordlist>|--threads <threads>|--timeout <timeout>|--recursive]
comet --paramscan <url> [--threads <threads>|--timeout <timeout>|--fuzz]
comet --portscan <ip> [<start_port>:<end_port>]
comet --osint <username>
comet --searchcmd <keyword>
comet --listcmd
```
And more commands.

## License
This project is licensed under the MIT License. <br>

<!-- Copyright (c) 2026 Zeronetsec -->