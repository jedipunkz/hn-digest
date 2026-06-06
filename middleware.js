export const config = {
  matcher: '/rss.xml',
};

export default function middleware(request) {
  const token = new URL(request.url).searchParams.get('token');
  const expected = process.env.RSS_TOKEN;

  if (!expected || token !== expected) {
    return new Response('Forbidden', {
      status: 403,
      headers: { 'content-type': 'text/plain; charset=utf-8' },
    });
  }
  // token 一致時は何も返さず静的 rss.xml の配信を継続
}
