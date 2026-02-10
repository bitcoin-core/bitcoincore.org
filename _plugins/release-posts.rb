class ReleasePostGenerator < Jekyll::Generator
  def generate(site)

    release_blog_templates = {}
    site.data["translations"].each do |lang, _|
      release_blog_templates[lang] = [
        File.exists?(File.join(site.source, "_includes", "posts", lang, "release-blog-major.md")),
        File.exists?(File.join(site.source, "_includes", "posts", lang, "release-blog-minor.md")),
      ]
    end

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

      # Generate the release blog post if we have the templates
      is_major = version[1] == 0
      if version[0] >= 28 and ((is_major and release_blog_templates[lang][0]) or (not is_major and release_blog_templates[lang][1]))
        release_blog_id = "#{lang}-blog-release-#{version_str}"
        if not site.posts.docs.find { |post| post.data["id"] == release_blog_id }
          date = page.data["date"]
          release_blog = Jekyll::Document.new("_posts/#{lang}/posts/#{date.strftime('%F')}-release-#{version_str}.md", site: site, collection: site.posts)
          release_blog.data["date"] = date
          release_blog.data["name"] = "blog-release-#{version_str}"
          release_blog.data["id"] = release_blog_id
          release_blog.data["lang"] = lang
          release_blog.data["type"] = "posts"
          release_blog.data["layout"] = "post"
          release_blog.data["version"] = 1
          release_blog.data["announcement"] = 1
          release_blog.data["excerpt"] = page.data["excerpt"]
          release_blog.data["release_version"] = version_str
          release_blog.data["title"] = site.data["translations"][lang]['released'].sub("%version", version_str)
          release_blog.data["permalink"] = "/#{lang}/#{date.strftime('%Y/%m/%d')}/release-#{version_str}/"

          if is_major and release_blog_templates[lang][0]
            release_blog.data["eol_version"] = "#{version[0] - 3}.x"
            release_blog.data["old_version"] = "#{version[0] - 1}.x"

            release_blog.content = "{% include posts/#{lang}/release-blog-major.md %}"
          elsif not is_major and release_blog_templates[lang][1]
            release_blog.content = "{% include posts/#{lang}/release-blog-minor.md %}"
          end

          site.posts.docs << release_blog
        end
      end
    end

  site.posts.docs.sort_by!(&:date)
  end
end
