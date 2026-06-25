<!-- https://github.com/Zeronetsec/Comet -->

[![version](https://img.shields.io/badge/Comet-Version%200.1-blue.svg)]()
[![os](https://img.shields.io/badge/Supported%20OS-Linux-blue.svg)]()
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

# Comet Framework
Comet is a lightweight framework specialized for high-efficiency security reconnaissance.

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
cd Comet
chmod +x install.sh
./install.sh
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

## License
This project is licensed under the MIT License.

<!-- Copyright (c) 2026 Zeronetsec -->