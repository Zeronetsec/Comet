function install::installer() {
    if [[ "${__BACKUP__}" == "true" && -d "${opt}/comet" ]]; then
        (
            cd "${opt}"
            install::getinstall \
                "
                    command zip -r \
                        comet_${bkdate}.bak.zip \
                        comet
                " \
                "Backup: ${GG}${opt}/comet ${DG}-> ${GG}${opt}/comet_${bkdate}.bak.zip${N}"
            cd
        )
    fi

    if [[ -d "${opt}/comet" ]]; then
        install::getinstall \
            "command rm -rf ${opt}/comet" \
            "Removing: old source..."
    fi

    install::getinstall \
        "command mv ${root} ${opt}/comet" \
        "Moving: ${GG}${root} ${DG}-> ${GG}${opt}/comet${N}"

    (
        cd "${opt}/comet"
        install::getinstall \
            "command go mod tidy" \
            "Retidy: ${GG}comet${N}"

        install::getinstall \
            "command go build -v -o comet" \
            "Building: ${GG}comet${N}"
        cd
    )

    install::getinstall \
        "
            command ln -sf \
                ${opt}/comet/comet \
                ${bin}/comet
        " \
        "Symlink: ${GG}${opt}/comet/comet ${DG}-> ${GG}${bin}/comet${N}"
}