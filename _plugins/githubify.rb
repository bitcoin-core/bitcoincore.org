# This file is licensed under the MIT License (MIT) available on
# http://opensource.org/licenses/MIT.

## githubify.rb automatically adds links to pull requests, issues, and
## commits using pattern matches

## Example (URL is the repository to link to):
## {% githubify https://github.com/bitcoin/bitcoin %}
## ...content...
## {% endgithubify %}

module Jekyll

require 'yaml'

  class GitHubifyBlock < Liquid::Block

    def initialize(tag_name, text, tokens)
      super
      @repository_url = text.strip()
    end

    def render(context)
      output = super

      ## Convert #1234 into URL for the pull request
      ## If #1234 links to an issue, GitHub automatically redirects
      #
      ## Main repository; forbid any lowercase alphabetical prefix to avoid conflation with
      ## other monotree repos like bitcoin-core/gui.  Require at least a
      ## double-digit number to reduce false positive matches (e.g. "fix thing,
      ## try #2")
      output.gsub!(/(^|[^a-z])#([0-9][0-9]+)/){ |s|
        $1 + '<a href="' + @repository_url + '/pull/' + $2 + '">#' + $2 + '</a>'
      }

      ## Other monotree repos; for repo "bitcoin-core/foo", PR format is "foo#123".
      ## Must start at word boundary
      output.gsub!(/\b([a-z]+)#([0-9]+)/){ |s|
        '<a href="https://github.com/bitcoin-core/' + $1 + '/pull/' + $2 + '">' + s + '</a>'
      }

      ## Convert `123abcd` into URL for the commit
      #
      ## Only operate on 7 to 10 chars to reduce false positive matches
      output.gsub!(/`([0-9abcdef]{7,10})`/){ |s|
        '<a href="' + @repository_url + '/commit/' + $1 + '">' + s + '</a>'
      }

      output
    end
  end
end


## Code to run if plugin is disabled
module Jekyll

require 'yaml'

  class GitHubifyBlockDisabled < Liquid::Block

    def initialize(tag_name, text, tokens)
      super
    end

    def render(context)
      output = super

      output
    end
  end
end

#Do nothing if plugin is disabled
plugin_name = "githubify"
if !ENV['ENABLED_PLUGINS'].nil? and ENV['ENABLED_PLUGINS'].index(plugin_name).nil?
  print plugin_name + ' disabled' + "\n"
  Liquid::Template.register_tag(plugin_name, Jekyll::GitHubifyBlockDisabled)
else
  Liquid::Template.register_tag(plugin_name, Jekyll::GitHubifyBlock)
end
