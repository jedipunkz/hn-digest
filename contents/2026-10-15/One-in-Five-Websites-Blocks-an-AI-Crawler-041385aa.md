---
source: "https://ciphercue.com/blog/most-websites-dont-block-ai-crawlers"
hn_url: "https://news.ycombinator.com/item?id=50102566"
title: "One in Five Websites Blocks an AI Crawler"
article_title: "One in Five Websites Blocks an AI Crawler | CipherCue - CipherCue"
image: "https://ciphercue.com/img/og/articles/most-websites-dont-block-ai-crawlers.png"
author: "adulion"
captured_at: "2026-10-15T13:07:14Z"
capture_tool: "hn-digest"
hn_id: 50102566
score: 2
comments: 1
posted_at: "2026-10-15T12:15:53Z"
tags:
  - hacker-news
---

# One in Five Websites Blocks an AI Crawler

- HN: [50102566](https://news.ycombinator.com/item?id=50102566)
- Source: [ciphercue.com](https://ciphercue.com/blog/most-websites-dont-block-ai-crawlers)
- Score: 2
- Comments: 1
- Posted: 2026-10-15T12:15:53Z

## Translation

Title: One in Five Websites Blocks an AI Crawler
Article title: One in Five Websites Blocks an AI Crawler | CipherCue - CipherCue
Description: We checked 110,744 websites against 19 AI crawlers. About 1 in 5 blocks at least one. Among those, 7 in 10 block nearly all. Most-blocked: GPTBot.

Article text:
One in Five Websites Blocks an AI Crawler | CipherCue - CipherCue
CIPHERCUE
Directory
EU Vendors
Trends
Docs
Pricing
Blog
Log in
Sign up free
Directory
EU Vendors
Trends
Docs
Pricing
Blog
Log in
Sign up free
analysis
One in Five Websites Blocks an AI Crawler
(LinkedIn)
21.3%
of the 110,744 websites we read block at least one of the 19 AI crawlers we checked for in their robots.txt. The other four in five have added no rule at all. Snapshot taken 8 October 2026.
robots.txt is the one lever a website has to tell an automated crawler to stay out. It is voluntary, the crawler has to choose to honour it, but it is also the only public, machine-readable signal a site publishes about who it wants indexing and training on its content. So we read it. We pulled the robots.txt for 110,744 websites and checked each one against a fixed list of 19 AI crawlers: the training bots (GPTBot, ClaudeBot, Google-Extended, CCBot, Bytespider and others), the search-assistant bots (PerplexityBot, OAI-SearchBot, DuckAssistBot), and the on-demand user-action fetchers (ChatGPT-User, Claude-Web, Perplexity-User).
About one site in five blocks at least one of these crawlers; the other four in five block nothing. For two years the argument over whether AI companies should train on the open web has been loud, but in robots.txt most sites have left no instruction either way. The rest of this piece is about the one in five that did, and what their configurations have in common.
When a site blocks, it usually blocks the lot
When a site decides to block AI crawlers, it rarely blocks just one or two. Almost always, it blocks nearly all of them at once.
Seven in ten of the sites that block anything block 13 or more of the 19 crawlers. So the decision is rarely "we are fine with OpenAI but not with ByteDance". It is "we have decided to keep AI crawlers out", applied across the board, usually by pasting in a block list that names all the well-known crawlers in one go. Blocking only a few is rare: fewer than one blocker in five does that. For almost everyone, it is all or nothing.
The most-blocked bot is OpenAI's
Here is the ranking of which crawlers get blocked most, by the number of sites that disallow each one.
GPTBot, OpenAI's training crawler, is the most-blocked, disallowed by nearly one site in five across the whole sample. The big training crawlers from Amazon, Anthropic, Meta, Google and Apple sit just behind it, all bunched close together. This is the signature of the wholesale block list: a site that adds a rule tends to add it for the whole group of well-known training crawlers at once, so they all rise and fall together.
CCBot (Common Crawl) and Bytespider (ByteDance) sit at the bottom, blocked by far fewer sites than the headline crawlers: roughly one site in twenty each, against GPTBot's nearly one in five. A smaller survey we ran on a single day's sample last month put Bytespider on top instead; at this scale the order is different. As the next section shows, the crawlers that rank low are the ones largely absent from the block lists sites tend to reuse, CCBot and Bytespider among them.
Opting out of training, staying open to search
Most AI companies now run two separate crawlers: one that gathers text to train models, and one that fetches pages to answer questions in an AI search product. OpenAI's are GPTBot for training and OAI-SearchBot for search, and robots.txt lets a site restrict them independently. We counted which of the two each site restricts.
Of the sites that block either OpenAI crawler, four in five block both. Among the rest, restricting the training crawler while leaving the search crawler alone is far more common than the reverse: fewer than a hundred sites in the whole sample block search but allow training. Anthropic and Perplexity show the same asymmetry.
So when sites do treat the two differently, they lean one way: block the training crawler, allow the search one. In plain terms, "don't train on my content, but you can still cite it in an answer." Whether that is a deliberate choice or just that training crawlers are the ones named in the copied block lists, robots.txt alone cannot tell us.
Blocking rates vary sharply by country
Break the numbers down by country and the national averages come apart, by more than tenfold between the top and bottom. Each country below has at least several hundred sites read, most of them thousands. The percentage for each country is accurate for the sites we measured, but those sites are not a representative cross-section of that country's web: the mix of industries, company sizes and hosting providers differs from one country to the next.
Most of the sites behind these numbers are names you would not recognise, which is itself part of the finding: this is a broad-web habit, not a handful of famous publishers. But some household names sit in there too. Amazon's robots.txt blocks 15 of the 19 crawlers, and the review site Trustpilot blocks 9.
Dutch sites block an AI crawler more than two in five times, British sites almost one in three, and German sites close behind at nearly one in four, then a cluster of other continental European countries in the low teens. The United States sits well below everyone else: fewer than three sites in a hundred block anything. The country home to OpenAI, Anthropic, Google and the loudest of the training-data lawsuits has the websites least likely to carry a single rule.
We can see the gap but can't explain it. The cause could be the mix of industries in each country, the hosting providers and web agencies common there and the defaults they ship, or simply which sites we happen to observe. What we can say is narrow: among the sites we see, blocking is far more common in the Netherlands and the UK than in the US, and we have not pinned down why.
You might expect the sites blocking every crawler to be the big news publishers, the ones suing AI companies or signing licensing deals. They are not what fills this list. Instead it is a mix of ordinary businesses: German and Dutch SME sites, UK and Irish software vendors, an Italian antiques marketplace, outdoor-equipment shops and local service firms, with no shared sector or size to connect them.
When you group the blockers by exactly which crawlers they disallow, most fall into one of two patterns. Nearly two in three (63.8%) block the same 17 crawlers. Another 6.1% block all 19. Between them, that is seven in ten blocking sites landing on just two configurations.
These split into two groups. Most of these sites are not targeting AI at all. They use a single blanket rule, User-agent: * with Disallow: / , that blocks every robot, so the AI crawlers get blocked along with Google and everything else. That is usually a staging site, a private app, or a login wall. The other group, about 7,000 sites, or one blocker in three, names an AI crawler directly. These are the sites choosing to keep AI out.
The composition of that 17-crawler set also explains the ranking. It covers the big Western training crawlers and the user-facing fetchers but omits CCBot and Bytespider, which is why those two sit at the bottom of the per-crawler table: a blanket Disallow: / catches them along with the rest, but almost none of the sites that name crawlers individually bother to list Common Crawl or ByteDance. Among the minority that do name crawlers by hand, there is no single dominant list: those configurations are fragmented across more than a thousand distinct combinations, most of them small.
It does not appear to be a WordPress plugin, the usual assumption. If anything the association runs the other way: 4.0% of blocking sites run WordPress, against 21.5% of the sites that block nothing, so WordPress sites block less often, not more. Whatever is behind these blanket rules, through hosting defaults, agency boilerplate, or a widely-copied snippet, it is not tracking the most common CMS.
How to block AI crawlers from your own site
If this piece has you wondering what your own robots.txt says, the mechanics are simple. robots.txt lives at the root of your domain, at /robots.txt , and is read top to bottom as a set of groups, each one a User-agent line followed by the rules that apply to it. A crawler only obeys a site that it chooses to honour, but the well-known AI crawlers all state that they read it.
To keep a single named crawler out, name it and disallow everything:
User-agent: GPTBot
Disallow: /
To opt out of a list of them, give each its own group. These are the training and search crawlers named most often, current as of this snapshot:
User-agent: GPTBot
Disallow: /
User-agent: OAI-SearchBot
Disallow: /
User-agent: ClaudeBot
Disallow: /
User-agent: Claude-SearchBot
Disallow: /
User-agent: Google-Extended
Disallow: /
User-agent: Applebot-Extended
Disallow: /
User-agent: Amazonbot
Disallow: /
User-agent: Meta-ExternalAgent
Disallow: /
User-agent: PerplexityBot
Disallow: /
User-agent: CCBot
Disallow: /
User-agent: Bytespider
Disallow: /
A few things worth knowing before you paste anything. A blanket rule, User-agent: * with Disallow: / , keeps out every automated visitor, which includes Google and Bing search indexing, not just AI; that is the rule most of the sites in this piece are using, often without meaning to deindex themselves. If your aim is narrower, name the crawlers you want to exclude and leave the wildcard group alone. Two of the tokens above, Google-Extended and Applebot-Extended , are opt-out signals rather than crawlers in their own right: disallowing them tells Google and Apple not to use your content for AI training without affecting ordinary search. And the list changes: vendors add and rename crawlers, so a list that is complete today will miss tomorrow. robots.txt is a stated request, not an enforcement mechanism; it records your preference and relies on the crawler to respect it.
This is 110,744 sites, read one robots.txt per site. It is a large sample but not a representative one: our scanner works through an entity graph that skews toward organisations with a corporate web presence, so read this as "among the organisations we observe" rather than "the web". robots.txt is also only a stated preference: a rule blocking GPTBot tells you the site asked OpenAI's crawler to stay out, not that it did. And absence of a rule is not consent, it is silence, which four in five sites are choosing by default rather than by decision. A survey that sampled high-traffic news domains specifically would find a much higher block rate, because that population self-selects for exactly the publishers who have added rules; we are measuring a broad, mostly non-media population, which is why our figure is lower. The raw counts are all in the tables above so you can compare like for like.
Primary source: the live robots.txt of each site, fetched once per site and parsed for user-agent block rules. We checked each site against a fixed list of 19 AI-crawler tokens (rule set v1.0) covering OpenAI, Anthropic, Google, Meta, Amazon, Apple, ByteDance, Common Crawl, Perplexity, Cohere, Mistral and DuckDuckGo, across training, search and user-action purposes. A crawler is counted as "blocked" for a site when that site's robots.txt disallows the crawler's user-agent token, either explicitly or via a wildcard that covers it. A couple of the 19 tokens, Google-Extended and Applebot-Extended, are opt-out signals rather than crawlers in their own right, so blocking them means declining to be used by those vendors' AI, which is the same decision in practice. Counts are one observation per site (latest fetch, deduplicated by domain). Snapshot date: 8 October 2026. Population: 110,744 sites, skewed toward organisations with a corporate web presence; not a representative sample of the web. robots.txt is a stated preference and does not prove a crawler's behaviour.
You can browse the organisations blocking each crawler at ciphercue.com/directory/ai-crawlers .
Researching a company mentioned in this article?
S

[truncated]

## Original Extract

We checked 110,744 websites against 19 AI crawlers. About 1 in 5 blocks at least one. Among those, 7 in 10 block nearly all. Most-blocked: GPTBot.

One in Five Websites Blocks an AI Crawler | CipherCue - CipherCue
CIPHERCUE
Directory
EU Vendors
Trends
Docs
Pricing
Blog
Log in
Sign up free
Directory
EU Vendors
Trends
Docs
Pricing
Blog
Log in
Sign up free
analysis
One in Five Websites Blocks an AI Crawler
(LinkedIn)
21.3%
of the 110,744 websites we read block at least one of the 19 AI crawlers we checked for in their robots.txt. The other four in five have added no rule at all. Snapshot taken 8 October 2026.
robots.txt is the one lever a website has to tell an automated crawler to stay out. It is voluntary, the crawler has to choose to honour it, but it is also the only public, machine-readable signal a site publishes about who it wants indexing and training on its content. So we read it. We pulled the robots.txt for 110,744 websites and checked each one against a fixed list of 19 AI crawlers: the training bots (GPTBot, ClaudeBot, Google-Extended, CCBot, Bytespider and others), the search-assistant bots (PerplexityBot, OAI-SearchBot, DuckAssistBot), and the on-demand user-action fetchers (ChatGPT-User, Claude-Web, Perplexity-User).
About one site in five blocks at least one of these crawlers; the other four in five block nothing. For two years the argument over whether AI companies should train on the open web has been loud, but in robots.txt most sites have left no instruction either way. The rest of this piece is about the one in five that did, and what their configurations have in common.
When a site blocks, it usually blocks the lot
When a site decides to block AI crawlers, it rarely blocks just one or two. Almost always, it blocks nearly all of them at once.
Seven in ten of the sites that block anything block 13 or more of the 19 crawlers. So the decision is rarely "we are fine with OpenAI but not with ByteDance". It is "we have decided to keep AI crawlers out", applied across the board, usually by pasting in a block list that names all the well-known crawlers in one go. Blocking only a few is rare: fewer than one blocker in five does that. For almost everyone, it is all or nothing.
The most-blocked bot is OpenAI's
Here is the ranking of which crawlers get blocked most, by the number of sites that disallow each one.
GPTBot, OpenAI's training crawler, is the most-blocked, disallowed by nearly one site in five across the whole sample. The big training crawlers from Amazon, Anthropic, Meta, Google and Apple sit just behind it, all bunched close together. This is the signature of the wholesale block list: a site that adds a rule tends to add it for the whole group of well-known training crawlers at once, so they all rise and fall together.
CCBot (Common Crawl) and Bytespider (ByteDance) sit at the bottom, blocked by far fewer sites than the headline crawlers: roughly one site in twenty each, against GPTBot's nearly one in five. A smaller survey we ran on a single day's sample last month put Bytespider on top instead; at this scale the order is different. As the next section shows, the crawlers that rank low are the ones largely absent from the block lists sites tend to reuse, CCBot and Bytespider among them.
Opting out of training, staying open to search
Most AI companies now run two separate crawlers: one that gathers text to train models, and one that fetches pages to answer questions in an AI search product. OpenAI's are GPTBot for training and OAI-SearchBot for search, and robots.txt lets a site restrict them independently. We counted which of the two each site restricts.
Of the sites that block either OpenAI crawler, four in five block both. Among the rest, restricting the training crawler while leaving the search crawler alone is far more common than the reverse: fewer than a hundred sites in the whole sample block search but allow training. Anthropic and Perplexity show the same asymmetry.
So when sites do treat the two differently, they lean one way: block the training crawler, allow the search one. In plain terms, "don't train on my content, but you can still cite it in an answer." Whether that is a deliberate choice or just that training crawlers are the ones named in the copied block lists, robots.txt alone cannot tell us.
Blocking rates vary sharply by country
Break the numbers down by country and the national averages come apart, by more than tenfold between the top and bottom. Each country below has at least several hundred sites read, most of them thousands. The percentage for each country is accurate for the sites we measured, but those sites are not a representative cross-section of that country's web: the mix of industries, company sizes and hosting providers differs from one country to the next.
Most of the sites behind these numbers are names you would not recognise, which is itself part of the finding: this is a broad-web habit, not a handful of famous publishers. But some household names sit in there too. Amazon's robots.txt blocks 15 of the 19 crawlers, and the review site Trustpilot blocks 9.
Dutch sites block an AI crawler more than two in five times, British sites almost one in three, and German sites close behind at nearly one in four, then a cluster of other continental European countries in the low teens. The United States sits well below everyone else: fewer than three sites in a hundred block anything. The country home to OpenAI, Anthropic, Google and the loudest of the training-data lawsuits has the websites least likely to carry a single rule.
We can see the gap but can't explain it. The cause could be the mix of industries in each country, the hosting providers and web agencies common there and the defaults they ship, or simply which sites we happen to observe. What we can say is narrow: among the sites we see, blocking is far more common in the Netherlands and the UK than in the US, and we have not pinned down why.
You might expect the sites blocking every crawler to be the big news publishers, the ones suing AI companies or signing licensing deals. They are not what fills this list. Instead it is a mix of ordinary businesses: German and Dutch SME sites, UK and Irish software vendors, an Italian antiques marketplace, outdoor-equipment shops and local service firms, with no shared sector or size to connect them.
When you group the blockers by exactly which crawlers they disallow, most fall into one of two patterns. Nearly two in three (63.8%) block the same 17 crawlers. Another 6.1% block all 19. Between them, that is seven in ten blocking sites landing on just two configurations.
These split into two groups. Most of these sites are not targeting AI at all. They use a single blanket rule, User-agent: * with Disallow: / , that blocks every robot, so the AI crawlers get blocked along with Google and everything else. That is usually a staging site, a private app, or a login wall. The other group, about 7,000 sites, or one blocker in three, names an AI crawler directly. These are the sites choosing to keep AI out.
The composition of that 17-crawler set also explains the ranking. It covers the big Western training crawlers and the user-facing fetchers but omits CCBot and Bytespider, which is why those two sit at the bottom of the per-crawler table: a blanket Disallow: / catches them along with the rest, but almost none of the sites that name crawlers individually bother to list Common Crawl or ByteDance. Among the minority that do name crawlers by hand, there is no single dominant list: those configurations are fragmented across more than a thousand distinct combinations, most of them small.
It does not appear to be a WordPress plugin, the usual assumption. If anything the association runs the other way: 4.0% of blocking sites run WordPress, against 21.5% of the sites that block nothing, so WordPress sites block less often, not more. Whatever is behind these blanket rules, through hosting defaults, agency boilerplate, or a widely-copied snippet, it is not tracking the most common CMS.
How to block AI crawlers from your own site
If this piece has you wondering what your own robots.txt says, the mechanics are simple. robots.txt lives at the root of your domain, at /robots.txt , and is read top to bottom as a set of groups, each one a User-agent line followed by the rules that apply to it. A crawler only obeys a site that it chooses to honour, but the well-known AI crawlers all state that they read it.
To keep a single named crawler out, name it and disallow everything:
User-agent: GPTBot
Disallow: /
To opt out of a list of them, give each its own group. These are the training and search crawlers named most often, current as of this snapshot:
User-agent: GPTBot
Disallow: /
User-agent: OAI-SearchBot
Disallow: /
User-agent: ClaudeBot
Disallow: /
User-agent: Claude-SearchBot
Disallow: /
User-agent: Google-Extended
Disallow: /
User-agent: Applebot-Extended
Disallow: /
User-agent: Amazonbot
Disallow: /
User-agent: Meta-ExternalAgent
Disallow: /
User-agent: PerplexityBot
Disallow: /
User-agent: CCBot
Disallow: /
User-agent: Bytespider
Disallow: /
A few things worth knowing before you paste anything. A blanket rule, User-agent: * with Disallow: / , keeps out every automated visitor, which includes Google and Bing search indexing, not just AI; that is the rule most of the sites in this piece are using, often without meaning to deindex themselves. If your aim is narrower, name the crawlers you want to exclude and leave the wildcard group alone. Two of the tokens above, Google-Extended and Applebot-Extended , are opt-out signals rather than crawlers in their own right: disallowing them tells Google and Apple not to use your content for AI training without affecting ordinary search. And the list changes: vendors add and rename crawlers, so a list that is complete today will miss tomorrow. robots.txt is a stated request, not an enforcement mechanism; it records your preference and relies on the crawler to respect it.
This is 110,744 sites, read one robots.txt per site. It is a large sample but not a representative one: our scanner works through an entity graph that skews toward organisations with a corporate web presence, so read this as "among the organisations we observe" rather than "the web". robots.txt is also only a stated preference: a rule blocking GPTBot tells you the site asked OpenAI's crawler to stay out, not that it did. And absence of a rule is not consent, it is silence, which four in five sites are choosing by default rather than by decision. A survey that sampled high-traffic news domains specifically would find a much higher block rate, because that population self-selects for exactly the publishers who have added rules; we are measuring a broad, mostly non-media population, which is why our figure is lower. The raw counts are all in the tables above so you can compare like for like.
Primary source: the live robots.txt of each site, fetched once per site and parsed for user-agent block rules. We checked each site against a fixed list of 19 AI-crawler tokens (rule set v1.0) covering OpenAI, Anthropic, Google, Meta, Amazon, Apple, ByteDance, Common Crawl, Perplexity, Cohere, Mistral and DuckDuckGo, across training, search and user-action purposes. A crawler is counted as "blocked" for a site when that site's robots.txt disallows the crawler's user-agent token, either explicitly or via a wildcard that covers it. A couple of the 19 tokens, Google-Extended and Applebot-Extended, are opt-out signals rather than crawlers in their own right, so blocking them means declining to be used by those vendors' AI, which is the same decision in practice. Counts are one observation per site (latest fetch, deduplicated by domain). Snapshot date: 8 October 2026. Population: 110,744 sites, skewed toward organisations with a corporate web presence; not a representative sample of the web. robots.txt is a stated preference and does not prove a crawler's behaviour.
You can browse the organisations blocking each crawler at ciphercue.com/directory/ai-crawlers .
Researching a company mentioned in this article?
S

[truncated]
