<!-- https://github.com/Zeronetsec/Comet -->

# Installation
`install.sh` optional option:
- `--backup`

Use `--backup` to create a backup of the existing Comet installation before replacing it.

## Termux & Linux (root)
```bash
git clone https://github.com/Zeronetsec/Comet
cd Comet
chmod +x install.sh
./install.sh
```

## Linux (user)
```bash
git clone https://github.com/Zeronetsec/Comet
cd Comet
chmod +x install.sh
sudo ./install.sh
```

## Uninstallation
```bash
export prefix="${PREFIX:-/usr}"
rm -f "${prefix}/bin/comet"
rm -rf "${prefix}/opt/comet"
```

<!-- Copyright (c) 2026 Zeronetsec -->