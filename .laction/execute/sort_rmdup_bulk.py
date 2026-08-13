import sys
from pathlib import Path

script_dir = Path(__file__).resolve().parent
project_root = script_dir.parent.parent

target_ext = [
    ".txt",
]

target_folder = [
    project_root / "console" / "wordlist" / "sqlscan",
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

def find_target_files(
    folders: list[Path],
    extensions: list[str],
) -> list[Path]:
    found = []
    for folder in folders:
        if not folder.exists():
            print(f"\x1b[1;31m[!] \x1b[0mFolder: \x1b[0;32m{folder} \x1b[0mnot found!")
            continue
        if not folder.is_dir():
            print(f"\x1b[1;31m[!] \x1b[0m{folder} \x1b[0mis not a folder!")
            continue

        for ext in extensions:
            found.extend(folder.rglob(f"*{ext}"))

    return found

all_files = find_target_files(
    target_folder, target_ext,
)

if not all_files:
    print("\x1b[1;31m[!] \x1b[0mNo matching files found!")
    sys.exit(1)

for file_path in all_files:
    clean_and_sort_file(file_path)