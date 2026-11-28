# make check is the check of the whole repository. CI runs it on every pull
# request, on Linux and on macOS.
.PHONY: check check-magos check-mac app

check: check-magos check-mac

# gofmt -l names each file that is not formatted; a name fails the check.
check-magos:
	! gofmt -l apps/magos | grep .
	go -C apps/magos vet ./...
	go -C apps/magos test ./...

check-mac:
ifeq ($(shell uname -s),Darwin)
	swift format lint --strict --recursive apps/mac/Package.swift apps/mac/Sources
	plutil -lint apps/mac/Info.plist
	swift build --package-path apps/mac
else
	@echo "Skipped apps/mac: the Swift check runs only on macOS."
endif

# Magos.app to run on this Mac. The helper and the signature come with #12.
app:
	swift build -c release --package-path apps/mac
	rm -rf apps/mac/.build/Magos.app
	mkdir -p apps/mac/.build/Magos.app/Contents/MacOS
	cp apps/mac/Info.plist apps/mac/.build/Magos.app/Contents/
	cp apps/mac/.build/release/Magos apps/mac/.build/Magos.app/Contents/MacOS/
	@echo "Built apps/mac/.build/Magos.app."
	@echo "Open it with: open apps/mac/.build/Magos.app"
