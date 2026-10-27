---
source: "https://ben.balter.com/2026/10/26/documentation-quality-is-ai-quality/"
hn_url: "https://news.ycombinator.com/item?id=50255332"
title: "Documentation quality is your AI quality now"
article_title: "Your documentation quality is your AI quality now | Ben Balter"
image: "https://ben.balter.com/og/2026/10/26/documentation-quality-is-ai-quality.png"
author: "benbalter"
captured_at: "2026-10-27T07:23:18Z"
capture_tool: "hn-digest"
hn_id: 50255332
score: 1
comments: 1
posted_at: "2026-10-27T06:30:34Z"
tags:
  - hacker-news
---

# Documentation quality is your AI quality now

- HN: [50255332](https://news.ycombinator.com/item?id=50255332)
- Source: [ben.balter.com](https://ben.balter.com/2026/10/26/documentation-quality-is-ai-quality/)
- Score: 1
- Comments: 1
- Posted: 2026-10-27T06:30:34Z

## Translation

Title: Documentation quality is your AI quality now
Article title: Your documentation quality is your AI quality now | Ben Balter
Description: AI tools are only as good as the docs you feed them. Bold text isn

Article text:
Your documentation quality is your AI quality now | Ben Balter Skip to main content Ben Balter Posts
Developer platforms and showing your work
Developer platforms and showing your work Skip to table of contents Home
Your documentation quality is your AI quality now
Your documentation quality is your AI quality now
9 min read TL;DR: Summary: AI tools need to know what your content is, not what it looks like. Use real headings and lists, keep one living copy, convert legacy Word files in with word2md.com, and export to Word only at the end with md2word.app.
Structure is what the model sees
Word files are tiny jails for content
Twelve years ago, I built a little Heroku app to convert Microsoft Word documents to Markdown, because government employees kept asking me for one. It now lives at word2md.com , and these days it’s busier than it has ever been. I’d bet most of the people using it aren’t developers. They’re people with a folder full of .docx files who’ve noticed their chatbot gives better answers when they paste in plain text.
They’ve figured out something I wrote about in Open and Async :
Point your AI tools at your team’s docs (via MCP servers, repository indexing, or custom instructions) and watch the difference. Teams with well-structured, up-to-date documentation get noticeably better AI-assisted answers because the model can find relevant context. Teams with scattered tribal knowledge get generic responses. Your documentation quality is your AI quality now.
If your institutional knowledge lives in a share drive full of Final_v3_REVISED_bb.docx , your AI tools are about as useful as a new hire who’s been handed the share drive and wished good luck. That’s a writing problem, and it was a writing problem long before the model showed up.
Markdown is the cheapest fix I know, for the reason I gave back in 2014 : it makes you write for the reader you actually have. In 2014, that reader was a browser. Now it’s often a model summarizing your proposal before your boss skims it.
Structure is what the model sees #
Word processors encourage you to describe what something looks like . Markup makes you declare what it is . In Markdown, ## Deprecating the v1 API says “this is a second-level heading.” Word can say the same thing with its Heading 2 style, but it also lets you just make the text big and bold, and most people do. The two look identical to a sighted reader. A screen reader or a language model sees something completely different. As I put it in the book, markup is a contract; formatting is a costume. (That’s also why my ebook turned out to be accessible by default . Screen readers and LLMs want the same thing.)
To see the difference, I built the same short vacation policy twice: once with Word’s heading and list styles, and once formatted by hand, with a bold 14-point heading and typed step numbers. Then I ran both through word2md . 1 The styled version came out as:
## Vacation policy
To request time off:
1. Tell your manager. 2. Add it to the team calendar. The hand-formatted version came out as:
**Vacation policy**
To request time off:
1 \. Tell your manager.
2 \. Add it to the team calendar. Same words, same look on screen. The first has a section a model can find and a list it can count. The second has a bold sentence and two paragraphs that happen to start with numbers (the backslash means “not a list”). No converter can recover structure the author never declared. Try it on one of your own documents: drop it into word2md.com and look for bold lines that should have been headings. It runs entirely in your browser, so the file never leaves your computer.
Tables fare even worse. When I ran the same file through macOS’s built-in plain-text converter ( textutil -convert txt ), the three-column table of accrual rates came out as twelve lines, one cell per line: Tenure , Days per year , Carryover , 0–2 years , 15 , 5 , and so on. Nothing ties the 15 to its column. Chatbots use better converters than that, but you can’t see what yours kept, and you won’t know until the summary comes back wrong.
Word files are tiny jails for content #
A .docx is a zipped bundle of XML built for one application to reopen. As I put it in the book:
Word processing documents are tiny jails for content. Proprietary formats blend text with the commands the software needs to recreate it, in formats built for the application rather than the reader. We put in text, one of the simplest things computers understand, and get back a mess that only specific software can open. It’s like keeping every grocery list inside a screenshot of an Instagram post: the words are there, but only one tool can read them, and you can’t edit so much as a typo without retaking the picture.
In 2014, I measured the same sample three ways: 33,621 bytes as Word, 1,359 as HTML, and 80 as Markdown. Even HTML, the fairer comparison, was seventeen times the bytes for the same words. 2 Context windows are big now, so size rarely blocks you, but a 2023 study found that models use information less reliably as inputs grow, especially when the part that matters sits in the middle ( Liu et al., “Lost in the Middle” ). Markup the model has to read past doesn’t help your odds.
Getting content in. word2md.com converts Word documents to clean Markdown entirely in your browser (nothing is uploaded), or from the command line with npx word-to-markdown input.docx . It’s a 2024 rewrite of that Heroku app ( source on GitHub ).
Getting content out. md2word.app goes the other way. Your AI assistant drafted the policy in Markdown, but legal wants a Word file they can redline, and your boss wants something they can open in Word. md2word builds one in your browser, and nothing gets uploaded. I built it this year ( source on GitHub ).
Or let your agent do it. Both tools are also MCP servers , so Claude, Copilot, or whatever agent you use can convert a .docx to Markdown and back without you opening a browser tab.
Markdown is where you and your AI tools do the work. Word is a delivery format you add at the last possible moment, for the people who need it. You don’t have to win over your whole organization to work this way. Stop doing your thinking in Word.
The round trip isn’t lossless. When legal sends the redlined .docx back, word2md drops the comments and silently accepts every tracked change. In my test, a redline changing “Full-time employees” to “All employees” came back as “All employees,” with no sign anyone had touched it. Resolve comments and review the redlines in Word first, then convert the clean copy and compare it against your source.
And if your team lives in Google Docs rather than Word files, you may not need to migrate anything. Docs has imported and exported Markdown since 2024 . It only helps if you used its real heading and list styles.
An LLM caught me being wrong #
I wrote a 99,651-word book in Markdown and Git, and over-engineered the pipeline that turns that one source into five formats. Because it was plain text, an LLM auditor could check it one chapter at a time. One pass looked only at factual claims. It stopped on a sentence saying GitHub’s monthly all-hands spent roughly half of each session on live Q&A. I was certain of that number, and it was wrong: closer to a third. No linter would have caught it. The structure made that review cheap enough to run on every chapter.
A playbook for AI-ready docs #
You don’t need a migration project.
Use real headings. In Word or Google Docs, that’s the Heading 1, 2, and 3 styles, not bold text. In Markdown, it’s ## .
Use real lists. Click the list button instead of typing numbers or dashes by hand.
Write headings that promise something. “Background” tells a reader (or a model searching for an answer) nothing. “Deprecating the v1 API” tells them exactly what follows.
Keep one living copy and share a link. One Google Doc or wiki page beats five attachments in five inboxes. If you can’t tell which version is current, neither can your AI tools. At GitHub, even the vacation policy worked this way ( non-technical teams can too ).
Convert legacy docs once, then stop editing the Word version. Run the important ones through word2md.com , check the output, and make that (or a properly styled Google Doc) the canonical copy.
Export to Word only at the end. When someone needs a .docx , generate it with md2word.app . Treat it like a printout: disposable, and never the source of truth.
Don’t wait for permission. Draft your next proposal with real structure, keep it somewhere with version history, and share a link instead of an attachment. When your reviews move faster, people will ask how.
Put your Markdown in version control. Every change becomes a diff you can review.
Keep your AI instructions in Markdown, too. AGENTS.md and CLAUDE.md are just Markdown in your repository. This blog has both.
Let your agent do the converting. Add the word2md and md2word MCP servers to your agent, and it can turn a .docx into Markdown (and back) on its own. Or run npx word-to-markdown input.docx and npx markdown-to-word input.md yourself.
Your team writes for the format you reward. Ask for a link, not an attachment, and publish your own decision docs with real structure first.
Markdown started as John Gruber’s blogging syntax in 2004 and spent most of the years since as a developer niche. Its bet was that structure outlasts formatting, and that bet is paying off for a reader nobody saw coming. Open and Async is the long version of this argument. “Documentation as a Superpower” makes the case that writing things down compounds, and “Retool your documents” is the practical half: which tools to use, and how to move a team off attachments without a migration project.
If you want the real thing rather than a two-step demo, the book’s appendix includes a complete time-off guide written this way. ↩︎
Math: 1,359 ÷ 80 ≈ 17. Bytes aren’t tokens, but the direction holds. The Word file was about 420 times the Markdown (33,621 ÷ 80), but it’s compressed, mostly fixed overhead, and not something a model reads directly. ↩︎
Liked this post? It's now a book.
Accessible by default: writing a book like software
The ebook meets the Web Content Accessibility Guidelines (WCAG 2.1 Level AA) because the format wouldn't let me fake a heading and a browser audits every build, not because anyone scheduled a week for it. The parts Markdown can't check, like color contrast, axe-core catches under Playwright.
Word versus Markdown: more than mere semantics
How we consume content has changed dramatically, yet how we author it hasn't. Markdown forces you to write for the web.
Diff (and collaborate on) Microsoft Word documents using GitHub
Word Diff automatically converts Word documents to Markdown when you commit to GitHub, so they can be diffed, versioned, and collaborated on like code.
My book has a linter, ~5,500 automated checks, and a Pandoc pipeline that rebuilds five formats (EPUB, Kindle, paperback ×2, and web) on every git push. For a book one person wrote. Here's how I built and published it the way I ship software.
Like what you read? Subscribe via RSS to get notified when new posts are published.
Prefer email? Subscribe to get new posts in your inbox. Leave this field empty Subscribe
Or follow along on Bluesky or LinkedIn .
I'm Ben Balter, and I write here about developer platforms, product and engineering leadership, open source, and showing your work. I spent a decade building developer platforms at GitHub : five years as a product manager for trust and safety, shipping 500+ features and the API endpoints behind them to 50M+ developers, plus a bottom-up rewrite of GitHub Pages and the open source library behind GitHub's license detection. My open source projects have hundreds of millions of downloads. I wrote Open and Async , the playbook for remote and distributed teams, and shipped an MCP server alongside it so agents can query it directly. At GitHub I was also Director of Hubber (Employee) Enablement, owning the internal platform

[truncated]

## Original Extract

AI tools are only as good as the docs you feed them. Bold text isn

Your documentation quality is your AI quality now | Ben Balter Skip to main content Ben Balter Posts
Developer platforms and showing your work
Developer platforms and showing your work Skip to table of contents Home
Your documentation quality is your AI quality now
Your documentation quality is your AI quality now
9 min read TL;DR: Summary: AI tools need to know what your content is, not what it looks like. Use real headings and lists, keep one living copy, convert legacy Word files in with word2md.com, and export to Word only at the end with md2word.app.
Structure is what the model sees
Word files are tiny jails for content
Twelve years ago, I built a little Heroku app to convert Microsoft Word documents to Markdown, because government employees kept asking me for one. It now lives at word2md.com , and these days it’s busier than it has ever been. I’d bet most of the people using it aren’t developers. They’re people with a folder full of .docx files who’ve noticed their chatbot gives better answers when they paste in plain text.
They’ve figured out something I wrote about in Open and Async :
Point your AI tools at your team’s docs (via MCP servers, repository indexing, or custom instructions) and watch the difference. Teams with well-structured, up-to-date documentation get noticeably better AI-assisted answers because the model can find relevant context. Teams with scattered tribal knowledge get generic responses. Your documentation quality is your AI quality now.
If your institutional knowledge lives in a share drive full of Final_v3_REVISED_bb.docx , your AI tools are about as useful as a new hire who’s been handed the share drive and wished good luck. That’s a writing problem, and it was a writing problem long before the model showed up.
Markdown is the cheapest fix I know, for the reason I gave back in 2014 : it makes you write for the reader you actually have. In 2014, that reader was a browser. Now it’s often a model summarizing your proposal before your boss skims it.
Structure is what the model sees #
Word processors encourage you to describe what something looks like . Markup makes you declare what it is . In Markdown, ## Deprecating the v1 API says “this is a second-level heading.” Word can say the same thing with its Heading 2 style, but it also lets you just make the text big and bold, and most people do. The two look identical to a sighted reader. A screen reader or a language model sees something completely different. As I put it in the book, markup is a contract; formatting is a costume. (That’s also why my ebook turned out to be accessible by default . Screen readers and LLMs want the same thing.)
To see the difference, I built the same short vacation policy twice: once with Word’s heading and list styles, and once formatted by hand, with a bold 14-point heading and typed step numbers. Then I ran both through word2md . 1 The styled version came out as:
## Vacation policy
To request time off:
1. Tell your manager. 2. Add it to the team calendar. The hand-formatted version came out as:
**Vacation policy**
To request time off:
1 \. Tell your manager.
2 \. Add it to the team calendar. Same words, same look on screen. The first has a section a model can find and a list it can count. The second has a bold sentence and two paragraphs that happen to start with numbers (the backslash means “not a list”). No converter can recover structure the author never declared. Try it on one of your own documents: drop it into word2md.com and look for bold lines that should have been headings. It runs entirely in your browser, so the file never leaves your computer.
Tables fare even worse. When I ran the same file through macOS’s built-in plain-text converter ( textutil -convert txt ), the three-column table of accrual rates came out as twelve lines, one cell per line: Tenure , Days per year , Carryover , 0–2 years , 15 , 5 , and so on. Nothing ties the 15 to its column. Chatbots use better converters than that, but you can’t see what yours kept, and you won’t know until the summary comes back wrong.
Word files are tiny jails for content #
A .docx is a zipped bundle of XML built for one application to reopen. As I put it in the book:
Word processing documents are tiny jails for content. Proprietary formats blend text with the commands the software needs to recreate it, in formats built for the application rather than the reader. We put in text, one of the simplest things computers understand, and get back a mess that only specific software can open. It’s like keeping every grocery list inside a screenshot of an Instagram post: the words are there, but only one tool can read them, and you can’t edit so much as a typo without retaking the picture.
In 2014, I measured the same sample three ways: 33,621 bytes as Word, 1,359 as HTML, and 80 as Markdown. Even HTML, the fairer comparison, was seventeen times the bytes for the same words. 2 Context windows are big now, so size rarely blocks you, but a 2023 study found that models use information less reliably as inputs grow, especially when the part that matters sits in the middle ( Liu et al., “Lost in the Middle” ). Markup the model has to read past doesn’t help your odds.
Getting content in. word2md.com converts Word documents to clean Markdown entirely in your browser (nothing is uploaded), or from the command line with npx word-to-markdown input.docx . It’s a 2024 rewrite of that Heroku app ( source on GitHub ).
Getting content out. md2word.app goes the other way. Your AI assistant drafted the policy in Markdown, but legal wants a Word file they can redline, and your boss wants something they can open in Word. md2word builds one in your browser, and nothing gets uploaded. I built it this year ( source on GitHub ).
Or let your agent do it. Both tools are also MCP servers , so Claude, Copilot, or whatever agent you use can convert a .docx to Markdown and back without you opening a browser tab.
Markdown is where you and your AI tools do the work. Word is a delivery format you add at the last possible moment, for the people who need it. You don’t have to win over your whole organization to work this way. Stop doing your thinking in Word.
The round trip isn’t lossless. When legal sends the redlined .docx back, word2md drops the comments and silently accepts every tracked change. In my test, a redline changing “Full-time employees” to “All employees” came back as “All employees,” with no sign anyone had touched it. Resolve comments and review the redlines in Word first, then convert the clean copy and compare it against your source.
And if your team lives in Google Docs rather than Word files, you may not need to migrate anything. Docs has imported and exported Markdown since 2024 . It only helps if you used its real heading and list styles.
An LLM caught me being wrong #
I wrote a 99,651-word book in Markdown and Git, and over-engineered the pipeline that turns that one source into five formats. Because it was plain text, an LLM auditor could check it one chapter at a time. One pass looked only at factual claims. It stopped on a sentence saying GitHub’s monthly all-hands spent roughly half of each session on live Q&A. I was certain of that number, and it was wrong: closer to a third. No linter would have caught it. The structure made that review cheap enough to run on every chapter.
A playbook for AI-ready docs #
You don’t need a migration project.
Use real headings. In Word or Google Docs, that’s the Heading 1, 2, and 3 styles, not bold text. In Markdown, it’s ## .
Use real lists. Click the list button instead of typing numbers or dashes by hand.
Write headings that promise something. “Background” tells a reader (or a model searching for an answer) nothing. “Deprecating the v1 API” tells them exactly what follows.
Keep one living copy and share a link. One Google Doc or wiki page beats five attachments in five inboxes. If you can’t tell which version is current, neither can your AI tools. At GitHub, even the vacation policy worked this way ( non-technical teams can too ).
Convert legacy docs once, then stop editing the Word version. Run the important ones through word2md.com , check the output, and make that (or a properly styled Google Doc) the canonical copy.
Export to Word only at the end. When someone needs a .docx , generate it with md2word.app . Treat it like a printout: disposable, and never the source of truth.
Don’t wait for permission. Draft your next proposal with real structure, keep it somewhere with version history, and share a link instead of an attachment. When your reviews move faster, people will ask how.
Put your Markdown in version control. Every change becomes a diff you can review.
Keep your AI instructions in Markdown, too. AGENTS.md and CLAUDE.md are just Markdown in your repository. This blog has both.
Let your agent do the converting. Add the word2md and md2word MCP servers to your agent, and it can turn a .docx into Markdown (and back) on its own. Or run npx word-to-markdown input.docx and npx markdown-to-word input.md yourself.
Your team writes for the format you reward. Ask for a link, not an attachment, and publish your own decision docs with real structure first.
Markdown started as John Gruber’s blogging syntax in 2004 and spent most of the years since as a developer niche. Its bet was that structure outlasts formatting, and that bet is paying off for a reader nobody saw coming. Open and Async is the long version of this argument. “Documentation as a Superpower” makes the case that writing things down compounds, and “Retool your documents” is the practical half: which tools to use, and how to move a team off attachments without a migration project.
If you want the real thing rather than a two-step demo, the book’s appendix includes a complete time-off guide written this way. ↩︎
Math: 1,359 ÷ 80 ≈ 17. Bytes aren’t tokens, but the direction holds. The Word file was about 420 times the Markdown (33,621 ÷ 80), but it’s compressed, mostly fixed overhead, and not something a model reads directly. ↩︎
Liked this post? It's now a book.
Accessible by default: writing a book like software
The ebook meets the Web Content Accessibility Guidelines (WCAG 2.1 Level AA) because the format wouldn't let me fake a heading and a browser audits every build, not because anyone scheduled a week for it. The parts Markdown can't check, like color contrast, axe-core catches under Playwright.
Word versus Markdown: more than mere semantics
How we consume content has changed dramatically, yet how we author it hasn't. Markdown forces you to write for the web.
Diff (and collaborate on) Microsoft Word documents using GitHub
Word Diff automatically converts Word documents to Markdown when you commit to GitHub, so they can be diffed, versioned, and collaborated on like code.
My book has a linter, ~5,500 automated checks, and a Pandoc pipeline that rebuilds five formats (EPUB, Kindle, paperback ×2, and web) on every git push. For a book one person wrote. Here's how I built and published it the way I ship software.
Like what you read? Subscribe via RSS to get notified when new posts are published.
Prefer email? Subscribe to get new posts in your inbox. Leave this field empty Subscribe
Or follow along on Bluesky or LinkedIn .
I'm Ben Balter, and I write here about developer platforms, product and engineering leadership, open source, and showing your work. I spent a decade building developer platforms at GitHub : five years as a product manager for trust and safety, shipping 500+ features and the API endpoints behind them to 50M+ developers, plus a bottom-up rewrite of GitHub Pages and the open source library behind GitHub's license detection. My open source projects have hundreds of millions of downloads. I wrote Open and Async , the playbook for remote and distributed teams, and shipped an MCP server alongside it so agents can query it directly. At GitHub I was also Director of Hubber (Employee) Enablement, owning the internal platform

[truncated]
