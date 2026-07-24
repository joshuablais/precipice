(use-modules (guix packages)
             (guix profiles)
             (gnu packages golang)
             (gnu packages golang-build)
             (gnu packages golang-check)
             (gnu packages golang-xyz)
             (gnu packages web)
             (gnu packages rust-apps)
             (renaissance packages air)
             )

(specifications->manifest '("go"
                            "delve"
                            "go-tools"
                            "templ"
                            "air"
                            "gopls"))
