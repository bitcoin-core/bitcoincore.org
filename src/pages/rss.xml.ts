import rss from '@astrojs/rss';
import { getCollection } from 'astro:content';
import type { APIContext } from 'astro';

export async function GET(context: APIContext) {
  const posts = await getCollection('posts', (entry) => {
    return entry.data.lang === 'en' && entry.data.type === 'posts' && entry.data.date;
  });

  return rss({
    title: 'Bitcoin Core',
    description: 'Bitcoin Core news and announcements',
    site: context.site!,
    items: posts
      .sort((a, b) => (b.data.date?.getTime() || 0) - (a.data.date?.getTime() || 0))
      .map((post) => ({
        title: post.data.title,
        pubDate: post.data.date!,
        description: post.data.excerpt || '',
        link: `/posts/${post.slug}/`,
        ...(post.data.tags && { categories: post.data.tags }),
      })),
  });
}