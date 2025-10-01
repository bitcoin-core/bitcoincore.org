import rss from '@astrojs/rss';
import { getCollection } from 'astro:content';
import type { APIContext } from 'astro';

export async function getStaticPaths() {
  return [
    { params: { lang: 'es' } },
    { params: { lang: 'ja' } },
    { params: { lang: 'zh-cn' } },
    { params: { lang: 'zh-tw' } },
  ];
}

export async function GET(context: APIContext) {
  const { lang } = context.params;
  
  const posts = await getCollection('posts', (entry) => {
    return entry.data.lang === lang && entry.data.type === 'posts' && entry.data.date;
  });

  const langNames: Record<string, string> = {
    'es': 'Español',
    'ja': '日本語', 
    'zh-cn': '简体中文',
    'zh-tw': '繁體中文'
  };

  return rss({
    title: `Bitcoin Core - ${langNames[lang!]}`,
    description: 'Bitcoin Core news and announcements',
    site: context.site!,
    items: posts
      .sort((a, b) => (b.data.date?.getTime() || 0) - (a.data.date?.getTime() || 0))
      .map((post) => ({
        title: post.data.title,
        pubDate: post.data.date!,
        description: post.data.excerpt || '',
        link: `/${lang}/posts/${post.slug}/`,
        ...(post.data.tags && { categories: post.data.tags }),
      })),
  });
}