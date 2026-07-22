<!-- https://github.com/Zeronetsec/Comet -->

# Installation
`install.sh` optional option:
- `--backup`
- └── create a backup of the existing source installation before replacing it.

### Usage
```bash
git clone https://github.com/Zeronetsec/Comet
bash Comet/install.sh <option>
```

# Uninstallation
`uninstall.sh` optional option:
- `--remove-backup`
- └── remove all backup found.

### Usage
```bash
export prefix="${PREFIX:-/usr}"
bash $prefix/opt/comet/uninstall.sh <option>
```

<!-- Copyright (c) 2026 Zeronetsec -->