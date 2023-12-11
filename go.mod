// After changing this file, run "go mod tidy"
module example.com/hello

go 1.21.4

// Upgrades every instance of the moving GC detection package to its last version.
// This is done to support the newest Go version without needing upstream changes.
replace go4.org/unsafe/assume-no-moving-gc => go4.org/unsafe/assume-no-moving-gc v0.0.0-20231121144256-b99613f794b6

require github.com/diamondburned/gotk4/pkg v0.0.5

require (
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20230221090011-e4bae7ad2296 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
)
