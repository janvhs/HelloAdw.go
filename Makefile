.PHONY: compile
compile: builddir
	meson compile -C builddir

builddir:
	meson setup builddir
