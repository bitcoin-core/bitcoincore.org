source 'https://rubygems.org'

## This version should match the version in Debian Stable.
## If you update the version here, also update it in .ruby-version, .cirrus.yml
## and README.md. Then push your branch and make sure Cirrus supports that
## version.
ruby '3.1.2'

## If you add a new Gem below, run `bundle install` to install it.
group :development do
  gem 'jekyll'
  gem 'jekyll-redirect-from'
  gem 'kramdown-parser-gfm'
end

group :testing do
  gem 'html-proofer', "~> 3.19"
end
