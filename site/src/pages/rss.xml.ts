import rss from '@astrojs/rss';
import type { APIContext } from 'astro';
import { getSummaries } from '../lib/summaries';

const escapeXML = (value: string) =>
  value
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;');

export function GET(context: APIContext) {
  const articles = getSummaries(14);

  return rss({
    title: 'HN Digest — AI · Platform · SRE',
    description:
      'Hacker News のトップ記事を日本語でサマリー（AI・Platform・SRE 分野中心、毎日 30 件）',
    site: context.site!,
    // Media RSS: how Inoreader, Feedly and friends pick up an item thumbnail.
    xmlns: { media: 'http://search.yahoo.com/mrss/' },
    items: articles.map((article) => {
      const title = article.title_ja || article.title;
      const image = article.image_url;
      return {
        title,
        pubDate: new Date(article.posted_at || article.date),
        description: article.summary_ja,
        link: article.hn_url,
        categories: ['Hacker News'],
        // Readers that render HTML show the image inline via content:encoded;
        // those that only read <description> still get the plain summary above.
        ...(image && {
          content: `<p><img src="${escapeXML(image)}" alt="${escapeXML(title)}" /></p><p>${escapeXML(
            article.summary_ja,
          )}</p>`,
          customData:
            `<media:content url="${escapeXML(image)}" medium="image" />` +
            `<media:thumbnail url="${escapeXML(image)}" />`,
        }),
      };
    }),
    customData: '<language>ja</language>',
  });
}
