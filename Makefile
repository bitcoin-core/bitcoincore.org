all: build test

preview:
	bundle exec jekyll clean
	bundle exec jekyll serve --future --drafts --unpublished --incremental --strict_front_matter

build:
	bundle exec jekyll clean
	bundle exec jekyll build --future --drafts --unpublished --strict_front_matter

test: test-fast test-slow

test-slow:
	## Check for malformed HTML and broken internal links
	bash -c "set -o pipefail ; bundle exec htmlproofer --check-html --disable-external --url-ignore '/^\/bin/.*/' ./_site | cat"
	## Check that links on the /en/download page point to the separately-stored binaries in /bin
	contrib/qa/test-binary-availability.sh && echo "SUCCESS checking URLs for binaries"

test-fast:
	## Check for broken Markdown reference-style links that are displayed in text unchanged, e.g. [broken][broken link]
	! find _site/ -name '*.html' | xargs grep ']\[' | grep -v skip-test | grep .
	## Ensure that no template strings leak through liquid rendering
	! find _site/ -name '*.html' | xargs grep '\$$(.*)'
