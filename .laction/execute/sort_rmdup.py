import sys
from pathlib import Path

script_dir = Path(__file__).resolve().parent
project_root = script_dir.parent.parent

target_files = [
    project_root / "console" / "wordlist" / "dirfuzzer" / "common.txt",
    project_root / "module" / "osint" / "sites" / "osint_sites.txt",
]

def clean_and_sort_file(file_path: Path):
    if not file_path.exists():
        print(f"\x1b[1;31m[!] \x1b[0mFile: \x1b[0;32m{file_path} \x1b[0mnot found!")
        sys.exit(1)

    content = file_path.read_text(encoding="utf-8").splitlines()
    clean_lines = {
        line.strip()
        for line in content
        if line.strip() and not line.strip().startswith("#")
    }

    sorted_lines = sorted(clean_lines)
    file_path.write_text(
        "\n".join(sorted_lines) + "\n",
        encoding="utf-8",
    )

    print(f"\x1b[0;32m[+] \x1b[0mFile: \x1b[0;32m{file_path.name} \x1b[0mhas been sorted and duplicates removed")

for file_path in target_files:
    clean_and_sort_file(file_path)