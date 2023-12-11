# HelloAdw

A hello-world-application in Go, showcasing how to integrate
Go with GTK4 and Meson.

## Building

```sh
$ meson setup builddir
$ meson compile -C builddir
```

## TOODs

Currently, this template is not finished, because it does not fullfil the same features
as the templates provided by GNOME Builder.
Follow Builder's Rust or C template and add them.

Here are some things, that will need to be done, before it can be considered finished.

- [ ] Add support for libadwaita 1.4 via gotk4-adwaita
- [ ] Add po translation files
- [ ] Add a XML-UI file
- [ ] Add a gschema
- [ ] Add a flatpak manifest
- [ ] Add application icons
