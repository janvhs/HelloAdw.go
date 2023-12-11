project('hello',
  version: '0.1.0',
  meson_version: '>= 0.62.0',
)

# Import needed compilers and meson modules.
go_cc = find_program('go', required: true)

# Query the system for required dependencies.
# gotk4 uses the gtk4 library, provided by pkg-config.
dependency('gtk4')

# Setup pre-compilation constants.
hello_conf = configuration_data()
hello_conf.set_quoted('VERSION', meson.project_version())

# Tell meson, where it should put the configuration data.
configure_file(
  input: 'internal' /'app' / 'conf.go.in',
  output: 'conf.go',
  configuration: hello_conf,
)

# Copy the config file to the source root, so it can compiled with the rest.
run_command(
  'cp',
  meson.project_build_root() / 'conf.go',
  meson.project_source_root() / 'internal' / 'app' / 'conf.go',
  check: true,
)

# This is the directory, where the main file is located.
# It is done in this meson.build file, because the main func is in the root dir.
hello_dir = meson.current_source_dir()

# Because meson has no Go integration, one has to use a custom target.
go_build_target = custom_target(
  build_always_stale: true, # always run go build, because only the dir is provided.
  build_by_default: true,
  command: [go_cc, 'build', '-v', '-o', '@OUTPUT@', hello_dir],
  console: true, # used for long running commands.
  output: meson.project_name(),
)
