function install::installer() {
    if [[ "${__BACKUP__}" == "true" && -d "${opt}/woofind" ]]; then
        (
            cd "${opt}"
            install::getinstall \
                "
                    command zip -r \
                        woofind_${bkdate}.bak.zip \
                        woofind
                " \
                "Backup: ${GG}${opt}/woofind ${DG}-> ${GG}${opt}/woofind_${bkdate}.bak.zip${N}"
            cd
        )
    fi

    if [[ -d "${opt}/woofind" ]]; then
        install::getinstall \
            "command rm -rf ${opt}/woofind" \
            "Removing: old source..."
    fi

    install::getinstall \
        "command mv ${root} ${opt}/woofind" \
        "Moving: ${GG}${root} ${DG}-> ${GG}${opt}/woofind${N}"

    (
        cd "${opt}/woofind"
        install::getinstall \
            "command go mod tidy" \
            "Retidy: ${GG}woofind${N}"

        install::getinstall \
            "command go build -v -o woofind" \
            "Building: ${GG}woofind${N}"
        cd
    )

    install::getinstall \
        "
            command ln -sf \
                ${opt}/woofind/woofind \
                ${bin}/woofind
        " \
        "Symlink: ${GG}${opt}/woofind/woofind ${DG}-> ${GG}${bin}/woofind${N}"
}; readonly -f install::installer