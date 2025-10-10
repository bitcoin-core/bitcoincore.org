FROM ruby:3.1.2

ENV NOKOGIRI_USE_SYSTEM_LIBRARIES=true
ENV JEKYLL_ENV=production
ENV RUBYOPT="-KU -E utf-8:utf-8"

RUN apt update \
    && apt install -y \
    curl \
    jekyll \
    make \
    ruby \
    ruby-html-proofer \
    ruby-jekyll-redirect-from \
    ruby-kramdown-parser-gfm

WORKDIR /site

COPY . .

RUN gem install bundler \
    && bundle install

RUN make build

# Default Jekyll port
EXPOSE 4000

CMD ["bundle", "exec", "jekyll", "server", "--host", "0.0.0.0", "--future"]
