all: build test

preview:
	bundle exec jekyll clean
	bundle exec jekyll serve --future --drafts --unpublished --incremental

build:
	bundle exec jekyll clean
	bundle exec jekyll build --future --drafts --unpublished

test:
	bundle exec htmlproof --disable-external --url-ignore '/^\/bin/.*/' ./_site
	contrib/qa/test-binary-availability.sh && echo "SUCCESS checking URLs for binaries"
