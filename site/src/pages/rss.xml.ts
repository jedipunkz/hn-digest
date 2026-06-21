import rss from '@astrojs/rss';
import type { APIContext } from 'astro';
import { getSummaries } from '../lib/summaries';

export function GET(context: APIContext) {
  const articles = getSummaries(14);

  return rss({
    title: 'HN Digest — AI · Platform · SRE',
    description:
      'Hacker News のトップ記事を AI が日本語でサマリー（AI・Platform・SRE 分野中心、毎日 10 件）',
    site: context.site!,
    items: articles.map((article) => ({
      title: article.title_ja || article.title,
      pubDate: new Date(article.posted_at || article.date),
      description: article.summary_ja,
      link: article.hn_url,
      categories: ['Hacker News'],
    })),
    customData: '<language>ja</language>',
  });
}
