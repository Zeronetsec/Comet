function install::installer() {
    if [[ ! -d "${HOME}/.${targetins}_log" ]]; then
        install::getinstall \
            "command mkdir -p ${HOME}/.${targetins}_log" \
            "Create directory: ${color_GG}${HOME}/.${targetins}_log${color_N}"
    fi

    (
        cd "${opt}/${targetins}"
        install::getinstall \
            "
                command go mod tidy
                command go build -o ${targetins}
            " \
            "Compiling: ${color_GG}${targetins}${color_N}"
    )
}; readonly -f install::installer