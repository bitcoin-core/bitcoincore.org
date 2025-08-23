class ReleasePostGenerator < Jekyll::Generator
  def generate(site)
    site.documents.each do |page|
      next unless page.data.key?("release")

      # Version numbers are either 0.a.b.c or a.b.c, and often c is 0. Drop c if it is 0
      version = page.data["release"]
      if ((version[0] == 0 and version.length == 4) or (version[0] > 0 and version.length == 3)) and version[-1] == 0
        version = version[..-2]
      end

      version_str = version.join(".")

      # Generate front matter
      page.data["title"] = "Bitcoin Core #{version_str}"
      lang = page.data.key?("lang") ? page.data["lang"] : "en"
      page.data["name"] = "release-#{version_str}"
      page.data["id"] = "#{lang}-#{page.data['name']}"
      page.data["excerpt"] = site.data["translations"][lang]["release_excerpt"].sub("%version", version_str)
      page.data["permalink"] = "/#{lang}/releases/#{version_str}/"

      # Wrap the page with the download include and githubify
      page.content = "{% include download.html %}\n{% githubify https://github.com/bitcoin/bitcoin %}\n#{page.content}\n{% endgithubify %}"
    end
  end
end
