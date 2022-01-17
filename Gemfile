source 'https://rubygems.org'

## This version should match the version in Debian Stable.
## If you update the version here, also update it in .ruby-version, .travis.yml
## and README.md. Then push your branch and make sure Travis supports that
## version.
ruby '3.0.3'

## If you add a new Gem below, run `bundle install` to install it.
group :development do
  gem 'github-pages'
end

group :testing do
  gem 'html-proofer'
end

gem "webrick", "~> 1.7"

gem 'wdm', '>= 0.1.0' if Gem.win_platform?