function install::installer() {
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