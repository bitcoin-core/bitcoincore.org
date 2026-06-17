Jekyll::Hooks.register :documents, :post_init do |post|

  # get the current post last modified time
  modification_time = `git log -1 --format="%ci" #{post.path}`

  # inject modification_time in post's datas.
  post.data['git-last-modified-date'] = modification_time

end
