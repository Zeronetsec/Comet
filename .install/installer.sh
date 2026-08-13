function install::installer() {
    if [[ ! -d "${HOME}/.${targetins}_log" ]]; then
        install::getinstall \
            "command mkdir -p ${HOME}/.${targetins}_log" \
            "Create directory: ${GG}${HOME}/.${targetins}_log${N}"
    fi

    (
        cd "${opt}/${targetins}"
        install::getinstall \
            "command go mod tidy" \
            "Retidy: ${GG}${targetins}${N}"

        install::getinstall \
            "command go build -v -o ${targetins}" \
            "Compiling: ${GG}${targetins}${N}"
        cd
    )
}; readonly -f install::installer