---
source: "https://zyvop.com/is-ai-actually-cheap-enough-to-replace-developers-8i74u"
hn_url: "https://news.ycombinator.com/item?id=50323552"
title: "Is AI Cheap Enough to Replace Developers?"
article_title: "Is AI Actually Cheap Enough to Replace Developers?"
image: "https://pub-03c883138afa4f88bc40a259a9109c50.r2.dev/posts/covers/1781860432295-l4zz5yhgmd.webp"
author: "sanjays442"
captured_at: "2026-11-01T08:48:37Z"
capture_tool: "hn-digest"
hn_id: 50323552
score: 1
comments: 1
posted_at: "2026-11-01T08:04:48Z"
tags:
  - hacker-news
---

# Is AI Cheap Enough to Replace Developers?

- HN: [50323552](https://news.ycombinator.com/item?id=50323552)
- Source: [zyvop.com](https://zyvop.com/is-ai-actually-cheap-enough-to-replace-developers-8i74u)
- Score: 1
- Comments: 1
- Posted: 2026-11-01T08:04:48Z

## Translation

Title: Is AI Cheap Enough to Replace Developers?
Article title: Is AI Actually Cheap Enough to Replace Developers?
Description: What a 35-point benchmark cliff, a slowdown study nobody wanted to publish, and real 2026 payroll data say about replacing engineers with agents

Article text:
Is AI Actually Cheap Enough to Replace Developers? ZYVOP Developer Publishing Product Migrate Explore Resources Pricing Search posts, series, authors... Ctrl K Login Get Started Home Is AI Actually Cheap Enough to Replace Developers? Is AI Actually Cheap Enough to Replace Developers?
What a 35-point benchmark cliff, a slowdown study nobody wanted to publish, and real 2026 payroll data say about replacing engineers with agents
June 19, 2026 Updated on August 27, 2026 · 8 min read # AI vs Developers # AI coding tools # developer-productivity # Claude Code # Software Engineering # GitHub Copilot Markdown Report Claude Opus 4.8 clears 88.6% on SWE-bench Verified . That's the number everyone quotes to argue AI has basically solved software engineering.
Now drop it into SWE-bench Pro — the version built from private codebases the model has never seen. The best standardized score on the public leaderboard is 59.1%. On the truly private commercial set, where nobody gets to tune in advance, no model clears 47%.
That's not a rounding error. That's a 20-to-40-point cliff, and it shows up everywhere. Gemini 3.1 Pro falls from 80.6% to somewhere between 32% and 46%, depending on which cut you trust. Every lab's flagship shows the same shape of fall.
Here's a wrinkle nobody saw coming: the two highest-scoring models that exist right now — Claude Mythos 5 (95.5%) and Claude Fable 5 (95.0%) — aren't something you can actually use. Anthropic suspended public access to both under an export control directive. The "best" model on paper this week is one you can't buy.
That gap, between the benchmark you can quote and the codebase you actually own, is the entire "AI replaces developers" argument in miniature. Models are extraordinary at problems that resemble what they've already seen. They get shaky fast the moment the code is genuinely yours — your weird conventions, your one engineer's undocumented caching hack from three years ago.
Cost was never the obstacle. Reliability was.
But cost is the question people actually ask. So let's answer it — before getting into why the answer doesn't settle anything.
The subscription math, updated for June 2026
Pricing in this category moves fast enough that anything written before May is already stale. And the biggest shift just happened: GitHub Copilot blew up its entire pricing model on June 1, 2026, moving from flat-rate to usage-based credits.
Pro+ ($39/mo) includes $70 in credits; Max includes $200
$60–200/mo on Pro+/Ultra once you pick premium models manually
token-credit billing since April 2026; $100–200/mo heavy use
The fallout from the Copilot switch has been loud. Developers who used to burn 3% of their monthly allowance on a normal day are now burning that much in under an hour. One person reported a single file review — no code written — eating a fifth of their monthly cap.
The flat $10–20 sticker price that defined this market for two years is functionally gone the moment you're doing real agentic work.
A solo developer paying for tokens directly, rather than a subscription, can still land around $5–$30 a month for light use. But push an agent into a hard debugging session and a single sitting can burn 500,000+ tokens and $15 or more. Real heavy-user spend across every tool here now sits at $100–$200/month — not the number on the landing page.
Scale that to a ten-person team running agents seriously, and you're at $1,000–$2,000+/month in tooling alone. That's before anyone counts the human time spent reviewing what the agent produced.
Here's the part most "AI is basically free" posts skip: the loaded cost of a developer isn't the number on the offer letter.
Add 25–30% for payroll tax and benefits. Add 15–25% more for equity at the senior tier. A $150,000 salary becomes $220,000–$280,000 in year one once recruiting and onboarding get folded in.
Set against that, even Claude Code Max or Cursor Ultra at $2,400/year is rounding error. On subscription-cost-versus-salary-cost alone, AI tooling wins by two orders of magnitude. Every time. Nobody serious disputes that.
The actual argument is whether the AI is doing the job — or just doing the typing while a human still does the job around it.
The productivity number nobody wanted to publish — and the sequel nobody expected
In July 2025, METR ran the most rigorous test of this question anyone had attempted. Sixteen experienced open-source developers. 246 real tasks. Randomized into AI-allowed and AI-disallowed groups, working in codebases they'd maintained for years.
The expected result was a speed-up. What they found instead: developers using AI took 19% longer .
Stranger still — those same developers believed they'd been 20% faster. They were wrong about the direction of their own productivity.
METR re-ran it. Bigger cohort, less self-selected: 57 developers, 143 repos, 800+ tasks. Published February 2026. The topline slowdown shrank to roughly -4%. And among the original developers who did both rounds, the number flipped entirely — an 18% speedup, though the confidence interval was wide enough that METR called it weak evidence, not proof.
Then the story took a turn nobody saw coming.
By April 2026, METR scrapped the whole experimental design. The reason: AI use had become so universal that 30–50% of invited developers simply refused to participate without AI access. You can't run a control group when the population won't be controlled.
So METR pivoted. Their May 2026 replacement is a self-report survey of 349 technical workers. Headline number: a median 1.4–2x self-reported increase in the value of their work, with respondents forecasting 2.5x by 2027.
Here's the twist. The one subgroup best positioned to judge this rigorously — METR's own researchers — reported the lowest gains of anyone surveyed. The agency that proved developers overestimate their own speedup just published a 2026 headline number built entirely on self-report, and told readers to be skeptical of it themselves.
What never changed, across every version of this research: the most experienced developers, on the most mature codebases, got the least benefit — sometimes a negative one. One Chrome engineer put it simply: it wasn't about being unfamiliar with the tools. It was about working in a codebase he already knew cold, versus a small greenfield project where a less experienced developer would benefit far more.
Same shape as the SWE-bench Pro cliff. Just measured with a stopwatch instead of a leaderboard.
A cautionary tale from outside engineering — now with a sequel of its own
Software isn't the only function that ran this experiment.
In 2023, Klarna replaced roughly 700 customer service staff with an OpenAI-built assistant. For a while, the numbers looked unambiguous: the bot handled two-thirds of all queries. By 2025, the company was walking it back. CEO Sebastian Siemiatkowski's own words: cost became too dominant a factor, and quality dropped in a way that wasn't sustainable.
The 2026 version of that story isn't a clean "AI failed, humans win" ending. Klarna's current setup is a hybrid: AI still handles roughly two-thirds of inquiries, but instead of rehiring full-time staff for the rest, the company routes the remainder to gig-style contractors — brought on flexibly, without the overhead of full employment.
It's a third option nobody was describing in 2023 or 2025: AI for volume, on-demand humans for what AI can't close, and full-time headcount for neither end.
Customer service and software engineering aren't the same job. But the failure mode is identical: optimize a replace-the-headcount decision purely on subscription cost versus salary cost, and you forget to price in what happens when the automated version meets a case it wasn't built for. For Klarna, that was an angry customer with a nuanced refund dispute. In engineering, it's a production incident the agent introduced and didn't flag — three weeks before anyone notices.
Where the cost argument actually wins outright
None of this means the cost case is wrong everywhere. It's right in one specific place: greenfield work, low stakes per mistake, human still firmly in the loop.
Solo-founder software businesses clearing seven and eight figures with single-digit or zero employees are a real, growing category in 2026. A few hundred dollars a month in agent subscriptions against an $80,000–$120,000-a-month human team isn't an exaggeration for that segment. It's just the math.
The same crowd is also the first to flag the real risk — and it's not the AI's coding ability. It's concentration of failure. Replace fifty employees with five hundred agents, and you don't become the CEO of a lean company. You become the single point of failure for every lawsuit, every hallucination, every 3am incident. That's a cost too. It just doesn't show up on the subscription invoice.
What the job market is doing, not saying
Surveys are cheap to answer dishonestly. Payroll data isn't.
Stanford HAI's 2026 AI Index, built on ADP payroll records rather than self-reported surveys, confirms it: employment for software developers aged 22–25 has fallen nearly 20% since late 2022 — the same window AI coding tools went mainstream. Developers over 30 in the same exposed roles grew employment 6–12% over that period.
But 2026 made the picture messier, not cleaner.
Salesforce CEO Marc Benioff confirmed the company isn't hiring additional software developers this fiscal year, crediting AI agents directly. In the same breath, he announced a push to hire 1,000 new college graduates — redirected toward sales and customer-facing roles, where AI hasn't displaced headcount. Read narrowly, that's not "AI replaced the juniors." It's "AI replaced one kind of junior work, and hiring moved to wherever that work doesn't exist yet."
IBM has kept its public commitment to triple US entry-level hiring, betting junior developers shift toward judgment-heavy, customer-facing work now that AI absorbs the repetitive part. Meanwhile, some labor economists are pushing back on the AI narrative entirely — pointing at interest-rate-driven hiring freezes and a brutal grad recruiting cycle as the more plausible cause, with AI serving as a convenient, simultaneously-timed scapegoat.
A year in, the more careful read: companies aren't cleanly eliminating junior developers. They're raising what "junior" has to mean on day one. AI-fluent juniors now command a premium north of 40% over generalists doing the job the old way.
Cheap enough to replace a developer's typing? Yes. Has been for a while. Not a close call. A $20–$200/month tool against a quarter-million-dollar loaded engineer wins on cost every single time — even after this year's pricing shake-up made the entry-level numbers less honest than they used to be.
Cheap enough to replace a developer's judgment — the part that knows which benchmark cliff is about to bite you, catches the bug the agent confidently introduced, decides what shouldn't be automated at all? Not yet. The clearest evidence anyone built to answer that question collapsed under the weight of how fast adoption moved. The agency that ran it is now relying on self-reported numbers it has publicly told you to doubt.
The gap between "passes the public benchmark" and "survives contact with your actual codebase" is still tens of points wide, no matter which model you pick. And the models posting the highest scores right now aren't even available to use.
The teams getting real value out of this in 2026 aren't asking "can AI replace developers." They're asking which slice of the week is mechanical enough to hand off, and which slice is the part the salary was actually paying for.
Get that split right, and the cost question mostly answers itself.
AI is cheap enough to replace a developer's typing, but not yet cheap enough to replace their judgment — and the evidence built to settle that question kept dissolving the moment researchers tried to pin it down.
Join the discussion by logging into your account.
No comments yet. Be th

[truncated]

## Original Extract

What a 35-point benchmark cliff, a slowdown study nobody wanted to publish, and real 2026 payroll data say about replacing engineers with agents

Is AI Actually Cheap Enough to Replace Developers? ZYVOP Developer Publishing Product Migrate Explore Resources Pricing Search posts, series, authors... Ctrl K Login Get Started Home Is AI Actually Cheap Enough to Replace Developers? Is AI Actually Cheap Enough to Replace Developers?
What a 35-point benchmark cliff, a slowdown study nobody wanted to publish, and real 2026 payroll data say about replacing engineers with agents
June 19, 2026 Updated on August 27, 2026 · 8 min read # AI vs Developers # AI coding tools # developer-productivity # Claude Code # Software Engineering # GitHub Copilot Markdown Report Claude Opus 4.8 clears 88.6% on SWE-bench Verified . That's the number everyone quotes to argue AI has basically solved software engineering.
Now drop it into SWE-bench Pro — the version built from private codebases the model has never seen. The best standardized score on the public leaderboard is 59.1%. On the truly private commercial set, where nobody gets to tune in advance, no model clears 47%.
That's not a rounding error. That's a 20-to-40-point cliff, and it shows up everywhere. Gemini 3.1 Pro falls from 80.6% to somewhere between 32% and 46%, depending on which cut you trust. Every lab's flagship shows the same shape of fall.
Here's a wrinkle nobody saw coming: the two highest-scoring models that exist right now — Claude Mythos 5 (95.5%) and Claude Fable 5 (95.0%) — aren't something you can actually use. Anthropic suspended public access to both under an export control directive. The "best" model on paper this week is one you can't buy.
That gap, between the benchmark you can quote and the codebase you actually own, is the entire "AI replaces developers" argument in miniature. Models are extraordinary at problems that resemble what they've already seen. They get shaky fast the moment the code is genuinely yours — your weird conventions, your one engineer's undocumented caching hack from three years ago.
Cost was never the obstacle. Reliability was.
But cost is the question people actually ask. So let's answer it — before getting into why the answer doesn't settle anything.
The subscription math, updated for June 2026
Pricing in this category moves fast enough that anything written before May is already stale. And the biggest shift just happened: GitHub Copilot blew up its entire pricing model on June 1, 2026, moving from flat-rate to usage-based credits.
Pro+ ($39/mo) includes $70 in credits; Max includes $200
$60–200/mo on Pro+/Ultra once you pick premium models manually
token-credit billing since April 2026; $100–200/mo heavy use
The fallout from the Copilot switch has been loud. Developers who used to burn 3% of their monthly allowance on a normal day are now burning that much in under an hour. One person reported a single file review — no code written — eating a fifth of their monthly cap.
The flat $10–20 sticker price that defined this market for two years is functionally gone the moment you're doing real agentic work.
A solo developer paying for tokens directly, rather than a subscription, can still land around $5–$30 a month for light use. But push an agent into a hard debugging session and a single sitting can burn 500,000+ tokens and $15 or more. Real heavy-user spend across every tool here now sits at $100–$200/month — not the number on the landing page.
Scale that to a ten-person team running agents seriously, and you're at $1,000–$2,000+/month in tooling alone. That's before anyone counts the human time spent reviewing what the agent produced.
Here's the part most "AI is basically free" posts skip: the loaded cost of a developer isn't the number on the offer letter.
Add 25–30% for payroll tax and benefits. Add 15–25% more for equity at the senior tier. A $150,000 salary becomes $220,000–$280,000 in year one once recruiting and onboarding get folded in.
Set against that, even Claude Code Max or Cursor Ultra at $2,400/year is rounding error. On subscription-cost-versus-salary-cost alone, AI tooling wins by two orders of magnitude. Every time. Nobody serious disputes that.
The actual argument is whether the AI is doing the job — or just doing the typing while a human still does the job around it.
The productivity number nobody wanted to publish — and the sequel nobody expected
In July 2025, METR ran the most rigorous test of this question anyone had attempted. Sixteen experienced open-source developers. 246 real tasks. Randomized into AI-allowed and AI-disallowed groups, working in codebases they'd maintained for years.
The expected result was a speed-up. What they found instead: developers using AI took 19% longer .
Stranger still — those same developers believed they'd been 20% faster. They were wrong about the direction of their own productivity.
METR re-ran it. Bigger cohort, less self-selected: 57 developers, 143 repos, 800+ tasks. Published February 2026. The topline slowdown shrank to roughly -4%. And among the original developers who did both rounds, the number flipped entirely — an 18% speedup, though the confidence interval was wide enough that METR called it weak evidence, not proof.
Then the story took a turn nobody saw coming.
By April 2026, METR scrapped the whole experimental design. The reason: AI use had become so universal that 30–50% of invited developers simply refused to participate without AI access. You can't run a control group when the population won't be controlled.
So METR pivoted. Their May 2026 replacement is a self-report survey of 349 technical workers. Headline number: a median 1.4–2x self-reported increase in the value of their work, with respondents forecasting 2.5x by 2027.
Here's the twist. The one subgroup best positioned to judge this rigorously — METR's own researchers — reported the lowest gains of anyone surveyed. The agency that proved developers overestimate their own speedup just published a 2026 headline number built entirely on self-report, and told readers to be skeptical of it themselves.
What never changed, across every version of this research: the most experienced developers, on the most mature codebases, got the least benefit — sometimes a negative one. One Chrome engineer put it simply: it wasn't about being unfamiliar with the tools. It was about working in a codebase he already knew cold, versus a small greenfield project where a less experienced developer would benefit far more.
Same shape as the SWE-bench Pro cliff. Just measured with a stopwatch instead of a leaderboard.
A cautionary tale from outside engineering — now with a sequel of its own
Software isn't the only function that ran this experiment.
In 2023, Klarna replaced roughly 700 customer service staff with an OpenAI-built assistant. For a while, the numbers looked unambiguous: the bot handled two-thirds of all queries. By 2025, the company was walking it back. CEO Sebastian Siemiatkowski's own words: cost became too dominant a factor, and quality dropped in a way that wasn't sustainable.
The 2026 version of that story isn't a clean "AI failed, humans win" ending. Klarna's current setup is a hybrid: AI still handles roughly two-thirds of inquiries, but instead of rehiring full-time staff for the rest, the company routes the remainder to gig-style contractors — brought on flexibly, without the overhead of full employment.
It's a third option nobody was describing in 2023 or 2025: AI for volume, on-demand humans for what AI can't close, and full-time headcount for neither end.
Customer service and software engineering aren't the same job. But the failure mode is identical: optimize a replace-the-headcount decision purely on subscription cost versus salary cost, and you forget to price in what happens when the automated version meets a case it wasn't built for. For Klarna, that was an angry customer with a nuanced refund dispute. In engineering, it's a production incident the agent introduced and didn't flag — three weeks before anyone notices.
Where the cost argument actually wins outright
None of this means the cost case is wrong everywhere. It's right in one specific place: greenfield work, low stakes per mistake, human still firmly in the loop.
Solo-founder software businesses clearing seven and eight figures with single-digit or zero employees are a real, growing category in 2026. A few hundred dollars a month in agent subscriptions against an $80,000–$120,000-a-month human team isn't an exaggeration for that segment. It's just the math.
The same crowd is also the first to flag the real risk — and it's not the AI's coding ability. It's concentration of failure. Replace fifty employees with five hundred agents, and you don't become the CEO of a lean company. You become the single point of failure for every lawsuit, every hallucination, every 3am incident. That's a cost too. It just doesn't show up on the subscription invoice.
What the job market is doing, not saying
Surveys are cheap to answer dishonestly. Payroll data isn't.
Stanford HAI's 2026 AI Index, built on ADP payroll records rather than self-reported surveys, confirms it: employment for software developers aged 22–25 has fallen nearly 20% since late 2022 — the same window AI coding tools went mainstream. Developers over 30 in the same exposed roles grew employment 6–12% over that period.
But 2026 made the picture messier, not cleaner.
Salesforce CEO Marc Benioff confirmed the company isn't hiring additional software developers this fiscal year, crediting AI agents directly. In the same breath, he announced a push to hire 1,000 new college graduates — redirected toward sales and customer-facing roles, where AI hasn't displaced headcount. Read narrowly, that's not "AI replaced the juniors." It's "AI replaced one kind of junior work, and hiring moved to wherever that work doesn't exist yet."
IBM has kept its public commitment to triple US entry-level hiring, betting junior developers shift toward judgment-heavy, customer-facing work now that AI absorbs the repetitive part. Meanwhile, some labor economists are pushing back on the AI narrative entirely — pointing at interest-rate-driven hiring freezes and a brutal grad recruiting cycle as the more plausible cause, with AI serving as a convenient, simultaneously-timed scapegoat.
A year in, the more careful read: companies aren't cleanly eliminating junior developers. They're raising what "junior" has to mean on day one. AI-fluent juniors now command a premium north of 40% over generalists doing the job the old way.
Cheap enough to replace a developer's typing? Yes. Has been for a while. Not a close call. A $20–$200/month tool against a quarter-million-dollar loaded engineer wins on cost every single time — even after this year's pricing shake-up made the entry-level numbers less honest than they used to be.
Cheap enough to replace a developer's judgment — the part that knows which benchmark cliff is about to bite you, catches the bug the agent confidently introduced, decides what shouldn't be automated at all? Not yet. The clearest evidence anyone built to answer that question collapsed under the weight of how fast adoption moved. The agency that ran it is now relying on self-reported numbers it has publicly told you to doubt.
The gap between "passes the public benchmark" and "survives contact with your actual codebase" is still tens of points wide, no matter which model you pick. And the models posting the highest scores right now aren't even available to use.
The teams getting real value out of this in 2026 aren't asking "can AI replace developers." They're asking which slice of the week is mechanical enough to hand off, and which slice is the part the salary was actually paying for.
Get that split right, and the cost question mostly answers itself.
AI is cheap enough to replace a developer's typing, but not yet cheap enough to replace their judgment — and the evidence built to settle that question kept dissolving the moment researchers tried to pin it down.
Join the discussion by logging into your account.
No comments yet. Be th

[truncated]
