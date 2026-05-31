import { existsSync, readFileSync, readdirSync } from 'fs';
import { dirname, join, resolve } from 'path';
import { fileURLToPath } from 'url';

const __dirname = dirname(fileURLToPath(import.meta.url));
// site/src/lib → site/src → site → repo root → summaries/
const SUMMARIES_DIR = resolve(__dirname, '..', '..', '..', 'summaries');

export interface Article {
  rank: number;
  hn_id: number;
  title: string;
  hn_url: string;
  source_url: string;
  score: number;
  final_score: number;
  comments: number;
  posted_at: string;
  summary_ja: string;
  date: string;
}

interface DaySummary {
  date: string;
  generated_at: string;
  articles: Omit<Article, 'date'>[];
}

export function getSummaries(days = 14): Article[] {
  if (!existsSync(SUMMARIES_DIR)) return [];

  const files = readdirSync(SUMMARIES_DIR)
    .filter((f) => f.endsWith('.json'))
    .sort()
    .reverse()
    .slice(0, days);

  const articles: Article[] = [];
  for (const file of files) {
    const raw = readFileSync(join(SUMMARIES_DIR, file), 'utf-8');
    const day: DaySummary = JSON.parse(raw);
    for (const art of day.articles) {
      articles.push({ ...art, date: day.date });
    }
  }
  return articles;
}
