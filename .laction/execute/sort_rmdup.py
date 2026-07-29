import os
import sys
from pathlib import Path

script_dir = Path(__file__).resolve().parent
project_root = script_dir.parent.parent

file_path = project_root / "console" / "wordlist" / "common.txt"

if file_path.exists():
    if not os.path.exists(file_path):
        print(f"\x1b[1;31m[!] \x1b[0mFile: \x1b[0;32m{input_file} \x1b[0mnot found!")
        sys.exit(1)

    content = file_path.read_text(encoding="utf-8").splitlines()
    clean_lines = {
        line.strip()
        for line in content
        if line.strip() and not line.strip().startswith("#")
    }

    sorted_lines = sorted(clean_lines)
    file_path.write_text("\n".join(sorted_lines) + "\n", encoding="utf-8")

    print(f"\x1b[0;32m[+] \x1b[0mFile: \x1b[0;32m{file_path} \x1b[0mhas been sorted and duplicates removed")