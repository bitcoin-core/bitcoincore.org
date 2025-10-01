import { defineCollection, z } from 'astro:content';

const posts = defineCollection({
  type: 'content',
  schema: z.object({
    title: z.string(),
    name: z.string(),
    id: z.string().optional(),
    lang: z.enum(['en', 'es', 'ja', 'zh_CN', 'zh_TW']).default('en'),
    permalink: z.string().optional(),
    type: z.enum(['posts', 'pages', 'meetings', 'advisory', 'releases']).default('posts'),
    version: z.number().optional(),
    announcement: z.number().optional(),
    excerpt: z.string().optional(),
    redirect_from: z.array(z.string()).optional(),
    tags: z.array(z.string()).optional(),
    translated: z.boolean().optional(),
    author: z.string().optional(),
    date: z.date().optional()
  }),
});

const releases = defineCollection({
  type: 'content',
  schema: z.object({
    title: z.string(),
    id: z.string().optional(),
    name: z.string(),
    permalink: z.string().optional(),
    excerpt: z.string().optional(),
    date: z.date().optional(),
    release: z.array(z.number()).optional(),
    optional_magnetlink: z.string().optional(),
    optional_date: z.date().optional(),
    download: z.string().optional(),
    type: z.string().optional(),
    lang: z.enum(['en', 'es', 'ja', 'zh_CN', 'zh_TW']).optional()
  }),
});

const pages = defineCollection({
  type: 'content',
  schema: z.object({
    title: z.string(),
    name: z.string(),
    id: z.string().optional(),
    lang: z.enum(['en', 'es', 'ja', 'zh_CN', 'zh_TW']).default('en'),
    permalink: z.string().optional(),
    type: z.enum(['posts', 'pages', 'meetings', 'advisory', 'releases']).default('pages'),
    version: z.number().optional(),
    announcement: z.number().optional(),
    excerpt: z.string().optional(),
    redirect_from: z.array(z.string()).optional(),
    tags: z.array(z.string()).optional(),
    translated: z.boolean().optional(),
    author: z.string().optional(),
    date: z.date().optional()
  }),
});

const docs = defineCollection({
  type: 'content',
  schema: z.object({
    name: z.string(),
    btcversion: z.string().optional(),
    btcgroup: z.string().optional(),
    permalink: z.string().optional(),
    title: z.string().optional()
  }),
});

export const collections = {
  posts,
  pages,
  releases,
  docs,
};