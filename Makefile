.PHONY: all clean test

all: terraform-provider-nsot .git/hooks/pre-commit

install: terraform-provider-nsot
	cp -f terraform-provider-nsot $$(dirname $$(which terraform))

terraform-provider-nsot: main.go nsot/*.go
	go build .

fmt:
	go fmt ./...

test: .git/hooks/pre-commit
	cd nsot ; go test -v .

clean:
	rm -f terraform-provider-nsot
	make -C yelppack clean

.git/hooks/pre-commit:
	    if [ ! -f .git/hooks/pre-commit ]; then ln -s ../../git-hooks/pre-commit .git/hooks/pre-commit; fi

itest_%:
	make -C yelppack $@

package: itest_trusty
