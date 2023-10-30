VERSION = none
MESSAGE = "Release $(VERSION)"

.PHONY: release

release:
	git add .
	git commit -m "$(MESSAGE)"
	git tag $(VERSION)
	git push origin $(VERSION)

	GOPROXY=proxy.golang.org go list -m github.com/dimkauzh/vuelto@$(VERSION)

	git push