<!-- https://github.com/Zeronetsec/Comet -->

<div align="center">
    <img src="https://img.shields.io/badge/Comet-Version%200.1-blue?style=square&logo=go&v=1" />
    <img src="https://img.shields.io/badge/Supported%20OS-Linux-blue?style=square&logo=linux&v=1" />
    <a href="LICENSE">
        <img src="https://img.shields.io/badge/License-GPL--3.0-blue?style=square&logo=github&v=1" />
    </a>
</div>

# Comet
Comet is a lightweight CLI tool built for high-efficiency security reconnaissance.

## Features
- Massive username enumeration across 700+ domains.
- High-performance directory fuzzing.
- HTML link crawling (hyperlink extraction).
- Parameter discovery via Wayback Machine.
- And more features.

## Disclaimer
Please read [.docs/disclaimer.md](.docs/disclaimer.md) before using this tool. </br>
Use this software at your own risk. </br>
The author is not responsible for any damage, data loss, or issues that may result from its use.

## Installation
Quick install:
```bash
git clone https://github.com/Zeronetsec/Comet
bash Comet/install.sh
```
For more detailed installation and uninstallation instructions, see [.docs/install_and_uninstall.md](.docs/install_and_uninstall.md).

## Usage Example
```bash
comet --dirfuzzer http://192.168.1.1/
comet --header http://192.168.1.1 --redirect true
comet --tracelink https://google.com --recursive
comet --dnslookup google.com
comet --paramscan http://testphp.vulnweb.com
```
And more commands.

## Credits
This project incorporates components from third-party sources. </br>
Please refer to [.docs/credits.md](.docs/credits.md) for full details and licensing information.

<!-- Copyright (c) 2026 Zeronetsec -->