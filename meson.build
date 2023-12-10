project('hello',
    version: '0.1.0',
    meson_version: '>= 0.62.0',
)

go_cc = find_program('go', required: true)

hello_src = meson.current_source_dir()

# gotk4 uses the gtk4 library, provided by pkg-config.
hello_deps = [
  dependency('gtk4'),
]

go_build_target = custom_target(
    build_always_stale: true,
    build_by_default: true,
    command: [go_cc, 'build', '-v', '-o', '@OUTPUT@', hello_src],
    console: true,
    output: meson.project_name(),
)
