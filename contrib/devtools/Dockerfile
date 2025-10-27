FROM ruby:3.1.2

ENV NOKOGIRI_USE_SYSTEM_LIBRARIES=true
ENV JEKYLL_ENV=production
ENV RUBYOPT="-KU -E utf-8:utf-8"

WORKDIR /site

COPY Gemfile* .

RUN gem install bundler \
    && bundle install

# Default Jekyll port
EXPOSE 4000

CMD ["bundle", "exec", "jekyll", "serve", "--future", "--drafts", "--unpublished", "--incremental", "--strict_front_matter", "--host", "0.0.0.0"]
