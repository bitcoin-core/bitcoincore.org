source 'https://rubygems.org'

## This version should match the version in Debian Stable.
## If you update the version here, also update it in .ruby-version,
## .github/workflows/ci.yml, contrib/devtools/Dockerfile, and README.md.
## Then push your branch and make sure CI supports that version.
ruby '3.3.8'

gem 'base64'
gem 'csv'

## If you add a new Gem below, run `bundle install` to install it.
group :development do
  gem 'jekyll'
  gem 'jekyll-redirect-from'
  gem 'kramdown-parser-gfm'
end

group :testing do
  gem 'html-proofer', "~> 3.19"
end
