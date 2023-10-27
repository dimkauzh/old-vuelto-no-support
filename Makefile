VERSION = none

.PHONY: release

release:
	GOPROXY=proxy.golang.org go list -m github.com/dimkauzh/vuelto@$(VERSION)