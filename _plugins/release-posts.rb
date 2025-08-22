class ReleasePostGenerator < Jekyll::Generator
  def generate(site)
    site.documents.each do |page|
      next unless page.data["type"] == "releases"

      page.content = "{% include download.html %}\n{% githubify https://github.com/bitcoin/bitcoin %}\n#{page.content}\n{% endgithubify %}"
    end
  end
end
