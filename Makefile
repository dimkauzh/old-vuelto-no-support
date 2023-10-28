VERSION = none

.PHONY: release

release:
	git add .
	git commit -m "Release $(VERSION)"
	git tag $(VERSION)
	git push origin $(VERSION)

	GOPROXY=proxy.golang.org go list -m github.com/dimkauzh/vuelto@$(VERSION)

	git push