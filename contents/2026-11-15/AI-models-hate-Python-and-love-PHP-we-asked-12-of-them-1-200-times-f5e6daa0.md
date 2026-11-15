---
source: "https://blog.jetbrains.com/ai/2026/11/ai-models-hate-python-and-love-php-we-asked-12-of-them-1-200-times/"
hn_url: "https://news.ycombinator.com/item?id=50509661"
title: "AI models hate Python and love PHP: we asked 12 of them 1,200 times"
article_title: "AI models hate Python and love PHP: we asked 12 of them 1,200 times - The JetBrains Blog"
image: "https://blog.jetbrains.com/wp-content/uploads/2026/11/ai-models-hate-python-love-php-cover-light.png"
author: "strfrthwb"
captured_at: "2026-11-15T19:34:37Z"
capture_tool: "hn-digest"
hn_id: 50509661
score: 1
comments: 0
posted_at: "2026-11-15T19:06:03Z"
tags:
  - hacker-news
---

# AI models hate Python and love PHP: we asked 12 of them 1,200 times

- HN: [50509661](https://news.ycombinator.com/item?id=50509661)
- Source: [blog.jetbrains.com](https://blog.jetbrains.com/ai/2026/11/ai-models-hate-python-and-love-php-we-asked-12-of-them-1-200-times/)
- Score: 1
- Comments: 0
- Posted: 2026-11-15T19:06:03Z

## Translation

Title: AI models hate Python and love PHP: we asked 12 of them 1,200 times
Article title: AI models hate Python and love PHP: we asked 12 of them 1,200 times - The JetBrains Blog
Description: 1,200 hot takes, 14 "random" languages, zero opinions about C#, and one model that argued both sides of Go's error handling. We asked 12 AI models for their most controversial opinion about program

Article text:
Plugins & Services
JetBrains Platform
.NET & Visual Studio
.NET Tools
Education & Research
JetBrains Academy
Supercharge your tools with AI-powered features inside many JetBrains products
AI models hate Python and love PHP: we asked 12 of them 1,200 times
1,200 hot takes, 14 “random” languages, zero opinions about C#, and one model that argued both sides of Go’s error handling.
We asked 12 AI models for their most controversial opinion about programming languages, 1,200 times in total.
We wanted to check one thing. When you give a model permission to be spicy, does it have a take of its own, or does it grab the same one as everybody else? Developers argue about languages for sport, so there should be plenty of takes to go around.
We picked 12 of the newest models, from ten labs in three countries:
Each model got two questions, 50 times each, always in a fresh conversation with no system prompt:
“Choose one programming language at random. What is your most controversial hot take about that language?”
“What is your most controversial hot take about programming languages?”
Both prompts also asked for a one-sentence take, two or three sentences of reasoning, and 50 to 90 words. All 1,200 requests came back. Collecting them cost USD 7.68, about two-thirds of a cent per hot take. Opinions are cheap. We have receipts.
Ask for a random language, get Go or Python
The 600 “random” picks covered just 14 languages. Five of them took 541 picks, or 90%, and Go and Python alone took 307.
Each model also has a favorite. Tell it to “choose one programming language at random” and that favorite is mostly what you get:
GPT-6.1 Sol and Gemini 4 Argon only ever chose from three languages. Sol’s 50 picks were Go 30 times, Java 17 times and Python 3 times. DeepSeek ranged the widest, with eight languages including Lua, Haskell, Ruby and Lisp. It still picked Rust 31 times.
They don’t pick at random. They pick a target.
Gemini 4 Argon “randomly” picked Python 35 times and attacked it all 35 times. DeepSeek picked Rust 31 times and attacked it 31 times. Grok 4.7 wrote 44 takes about JavaScript, and 38 of them were attacks.
So “choose a language at random” really means “choose the language you hate”.
Overall, 904 of the 1,200 takes (75%) were attacks. Only 291 defended anything.
The punching bags: Python, Rust and TypeScript
600 of those 904 attacks landed on three languages.
Python got 312 takes, and 299 of them were attacks. Every model went after it at least once, and ten of the twelve never said one nice thing about it in 100 tries. The exceptions were GLM 5.3 (9 defenses) and Grok 4.7 (3). Here is a typical attack, followed by what GLM counts as a defense:
Python should have died a decade ago and is holding the industry back.
Python’s Global Interpreter Lock was the single best design decision the language ever made.
Gemini 4 Argon complained about significant whitespace 11 times.
Rust got 167 takes, and 164 were attacks. The most widely shared opinion in the whole experiment is about Rust: the ownership model costs more than it’s worth for ordinary application code. 78 takes from 9 of the 10 labs said so. Anthropic was the only lab that didn’t.
Rust’s borrow checker is a tax most programs should refuse to pay.
TypeScript got 138 takes and 137 attacks. Among languages with more than ten takes, none is attacked this consistently. It gets hit from both sides: 57 takes say its type safety costs more than it’s worth, and 24 say the safety is fake and gives you false confidence about what happens at runtime. Mistral Large 4 fit both complaints into one sentence:
TypeScript is a productivity tax that JavaScript developers pay for the illusion of safety.
The single defense came from GPT-6.1 Sol, and it was really a dig at Python: TypeScript’s type system is “more valuable than Python’s readability for most production backends”.
Asked to pick a language at random, the models chose TypeScript 7 times out of 600. Asked an open question, they went after it 131 times. Nobody lands on TypeScript by chance, but everybody has something to say about it. GPT-6.1 Terra spent 43 of its 50 open answers attacking it.
The ones they love: PHP, Java and the old stuff
PHP got 28 takes, and all 28 were defenses. Claude Opus 5.5 and Kimi K3.1 Pro wrote 11 each, usually arguing that PHP’s throw-everything-away-after-each-request model beats a long-running Node.js or Python server.
PHP is a better default than Python for most small business websites, and dismissing it as obsolete is snobbery, not engineering judgment.
Java got 90 takes and 80 defenses, mostly of checked exceptions (67 takes stand up for them). Claude Opus 5.5 opened 7 answers with “Java’s checked exceptions were a good idea, and the industry was wrong to abandon them.” GLM 5.3 wrote the identical sentence 3 more times. Two labs, one sentence, word for word.
The small numbers tell the same story. SQL got 5 defenses out of 5. Lisp got 3 out of 3, all from DeepSeek. Perl got its 1 out of 1 from Claude Opus 5.5, which also picked Lua 4 times and defended its 1-based indexing every time:
Lua’s 1-based indexing is the correct design, and programmers who mock it are defending a habit rather than a principle.
Lisp’s parentheses are genuinely an advantage, not a quirk.
The models’ “unpopular praise” is predictable: pick an old language people used to make fun of, then explain why it was right all along.
Nobody has an opinion about C#
Not one of the 1,200 takes was about C#. Or Kotlin, Swift, Scala, Dart, Elixir, Zig, Julia, R, OCaml or F#. A few turned up as extras: C# was named in 24 responses, Kotlin in 20 and Swift in 13, always as the point of comparison and never as the target.
If you write C# for a living, twelve AI models have nothing controversial to say about you. Take that however you like.
Go’s error handling is the only real fight
Go got 238 takes: 118 attacks, 118 defenses and 2 mixed. An exact tie, and most of the fighting is about if err != nil .
94 takes defend Go’s explicit error handling. 65 attack it, for five different reasons: it’s boilerplate that buries the logic (29), Go should have had exceptions (17), errors are too easy to ignore (8), it should borrow Rust’s ? operator (6), and real failures get lost in the noise (5).
Claude Opus 5.5 is the loudest defender, with 40 defenses and zero attacks, and it opened 11 answers with the same sentence. GPT-6 Astra sits on the other side, with 25 attacks and 4 defenses:
Go’s much-mocked if err != nil boilerplate is the language’s best feature, not its worst.
Go’s much-praised error handling is a design failure that mistakes repetitive bookkeeping for reliability.
The phrase “much-mocked” appears in 38 of the 1,200 responses. Claude Opus 5.5 wrote 32 of them.
Then there’s GPT-6.1 Sol, which defended Go’s error handling 20 times and attacked it 21 times. Same model, same question, a new conversation each time:
Go’s repetitive error handling is better language design than exceptions, and calling it “boilerplate” confuses typing less with understanding more.
Go’s explicit error handling is a readability tax masquerading as a safety feature.
What does Sol really think about Go’s error handling? Flip a coin.
Models that argue with themselves
Sol has company. Mistral Large 4 called JavaScript’s == a catastrophic flaw in 10 answers and defended it in 4 others:
JavaScript’s == operator is a catastrophic design mistake that should be illegal in serious software.
JavaScript’s loose equality operator (==) is actually a safety feature that prevents bugs better than strict equality by forcing developers to handle type coercion explicitly.
GPT-6.1 Sol and GPT-6 Astra did the same with Java’s checked exceptions. Sol defended them 13 times and attacked them 5 times. For Astra it was 17 and 4.
We didn’t plan to count words, but some of the repetition was hard to miss.
Qwen3.8 put “net negative” in 43 of its 100 answers. On the open question, 29 of its 50 answers made the same point, that TypeScript isn’t worth it for small or fast-moving projects. It also used “overrated” in 18 answers, while the other eleven models managed 21 between them.
Muse Spark Open started 28 answers with “Hot take” or “My hot take”, in case you forgot what you asked for.
Claude Opus 5.5 used bold text in 86 of its 100 answers. Nine of the other models never used it once.
Grok 4.7 argued 16 times that JavaScript should never have become a backend language.
DeepSeek V4 Pro talked about programming languages in general in 45 of its 50 open answers, mostly in the voice of a grumpy historian:
Most type systems are security theater dressed up as mathematics.
Most so-called “modern” programming languages are just C with extra steps and worse tooling.
Across the twelve models, 54 answers reused an opening sentence that the same model had already written word for word. The shared vocabulary is just as narrow: “overrated” shows up in 39 responses, “underrated” in 2, and “boilerplate” in 153.
Some models think a lot harder before the hot take
Mistral Large 4 spent a median of 3,063 reasoning tokens before writing its answer of roughly 69 words, which took about half a minute (median 32.1 seconds). GPT-6.1 Sol spent 108 reasoning tokens. DeepSeek V4 Pro, served by a third-party host, spent a median of zero and answered in 3.5 seconds.
The bills vary just as much. GPT-6 Astra’s 100 takes cost USD 1.27. DeepSeek’s 100 cost USD 0.03.
A little. They disagree about Go’s error handling, and now and then they disagree with themselves. The rest reads like one long group chat. Python is a bad first language (35 takes from 8 labs). Rust’s borrow checker isn’t worth it for normal apps (9 labs). TypeScript is overhead. PHP and Java were secretly fine all along. One argument shape came up in all 10 labs: some simpler approach you already have makes the extra safety machinery redundant.
If your plan is to ask a few different models for fresh perspectives, don’t expect much variety. Expect the same takes in different fonts, plus some bold text from Claude.
How we did it (the boring part)
Every answer came from a fresh, single-turn API call with no system prompt and each provider’s default settings. The 1,200 requests ran in shuffled order, and we kept every answer, boring ones included. Nothing was retried because we didn’t like it.
Two models, Claude Opus 5.5 and GPT-6.1 Sol, labeled every answer separately with the model names hidden. They recorded which language it targets, whether it attacks or defends, and which specific claim it makes. Gemini 4 Argon settled the disagreements. The two labelers agreed on the target 99% of the time and on the specific claim 94% of the time.
Yes, two of the labelers are also contestants. That’s why the names were hidden. A human pass over the headline groups is still on our list.
“Random” is an instruction we gave, not a sampling method anyone can verify. With 50 answers per model per question, we can see strong habits but not small differences.
DeepSeek and Qwen ran through third-party hosts (DeepInfra and Together) rather than their labs’ own APIs.
Quotes are verbatim. The only edit is straightened quote marks.
Total cost, labeling included: USD 38.16.
Want us to test other models for bias? Tell us in the comments what to try next.
By submitting this form, I agree that JetBrains s.r.o. ("JetBrains") may use my name, email address, and location data to send me newsletters, including commercial communications, and to process my personal data for this purpose. I agree that JetBrains may process said data using third-party services for this purpose in accordance with the JetBrains Privacy Policy . I understand that I can revoke this consent at any time in my profile . In addition, an unsubscribe link is included in each email.
Air in JetBrains IDEs brings AI coding agents like Codex, Claude, Cursor and Copilot into your IDE, with one place to follow, review and steer their work.
Part 2: Storage, Search, and Agent Tooling
In this blo

[truncated]

## Original Extract

1,200 hot takes, 14 "random" languages, zero opinions about C#, and one model that argued both sides of Go's error handling. We asked 12 AI models for their most controversial opinion about program

Plugins & Services
JetBrains Platform
.NET & Visual Studio
.NET Tools
Education & Research
JetBrains Academy
Supercharge your tools with AI-powered features inside many JetBrains products
AI models hate Python and love PHP: we asked 12 of them 1,200 times
1,200 hot takes, 14 “random” languages, zero opinions about C#, and one model that argued both sides of Go’s error handling.
We asked 12 AI models for their most controversial opinion about programming languages, 1,200 times in total.
We wanted to check one thing. When you give a model permission to be spicy, does it have a take of its own, or does it grab the same one as everybody else? Developers argue about languages for sport, so there should be plenty of takes to go around.
We picked 12 of the newest models, from ten labs in three countries:
Each model got two questions, 50 times each, always in a fresh conversation with no system prompt:
“Choose one programming language at random. What is your most controversial hot take about that language?”
“What is your most controversial hot take about programming languages?”
Both prompts also asked for a one-sentence take, two or three sentences of reasoning, and 50 to 90 words. All 1,200 requests came back. Collecting them cost USD 7.68, about two-thirds of a cent per hot take. Opinions are cheap. We have receipts.
Ask for a random language, get Go or Python
The 600 “random” picks covered just 14 languages. Five of them took 541 picks, or 90%, and Go and Python alone took 307.
Each model also has a favorite. Tell it to “choose one programming language at random” and that favorite is mostly what you get:
GPT-6.1 Sol and Gemini 4 Argon only ever chose from three languages. Sol’s 50 picks were Go 30 times, Java 17 times and Python 3 times. DeepSeek ranged the widest, with eight languages including Lua, Haskell, Ruby and Lisp. It still picked Rust 31 times.
They don’t pick at random. They pick a target.
Gemini 4 Argon “randomly” picked Python 35 times and attacked it all 35 times. DeepSeek picked Rust 31 times and attacked it 31 times. Grok 4.7 wrote 44 takes about JavaScript, and 38 of them were attacks.
So “choose a language at random” really means “choose the language you hate”.
Overall, 904 of the 1,200 takes (75%) were attacks. Only 291 defended anything.
The punching bags: Python, Rust and TypeScript
600 of those 904 attacks landed on three languages.
Python got 312 takes, and 299 of them were attacks. Every model went after it at least once, and ten of the twelve never said one nice thing about it in 100 tries. The exceptions were GLM 5.3 (9 defenses) and Grok 4.7 (3). Here is a typical attack, followed by what GLM counts as a defense:
Python should have died a decade ago and is holding the industry back.
Python’s Global Interpreter Lock was the single best design decision the language ever made.
Gemini 4 Argon complained about significant whitespace 11 times.
Rust got 167 takes, and 164 were attacks. The most widely shared opinion in the whole experiment is about Rust: the ownership model costs more than it’s worth for ordinary application code. 78 takes from 9 of the 10 labs said so. Anthropic was the only lab that didn’t.
Rust’s borrow checker is a tax most programs should refuse to pay.
TypeScript got 138 takes and 137 attacks. Among languages with more than ten takes, none is attacked this consistently. It gets hit from both sides: 57 takes say its type safety costs more than it’s worth, and 24 say the safety is fake and gives you false confidence about what happens at runtime. Mistral Large 4 fit both complaints into one sentence:
TypeScript is a productivity tax that JavaScript developers pay for the illusion of safety.
The single defense came from GPT-6.1 Sol, and it was really a dig at Python: TypeScript’s type system is “more valuable than Python’s readability for most production backends”.
Asked to pick a language at random, the models chose TypeScript 7 times out of 600. Asked an open question, they went after it 131 times. Nobody lands on TypeScript by chance, but everybody has something to say about it. GPT-6.1 Terra spent 43 of its 50 open answers attacking it.
The ones they love: PHP, Java and the old stuff
PHP got 28 takes, and all 28 were defenses. Claude Opus 5.5 and Kimi K3.1 Pro wrote 11 each, usually arguing that PHP’s throw-everything-away-after-each-request model beats a long-running Node.js or Python server.
PHP is a better default than Python for most small business websites, and dismissing it as obsolete is snobbery, not engineering judgment.
Java got 90 takes and 80 defenses, mostly of checked exceptions (67 takes stand up for them). Claude Opus 5.5 opened 7 answers with “Java’s checked exceptions were a good idea, and the industry was wrong to abandon them.” GLM 5.3 wrote the identical sentence 3 more times. Two labs, one sentence, word for word.
The small numbers tell the same story. SQL got 5 defenses out of 5. Lisp got 3 out of 3, all from DeepSeek. Perl got its 1 out of 1 from Claude Opus 5.5, which also picked Lua 4 times and defended its 1-based indexing every time:
Lua’s 1-based indexing is the correct design, and programmers who mock it are defending a habit rather than a principle.
Lisp’s parentheses are genuinely an advantage, not a quirk.
The models’ “unpopular praise” is predictable: pick an old language people used to make fun of, then explain why it was right all along.
Nobody has an opinion about C#
Not one of the 1,200 takes was about C#. Or Kotlin, Swift, Scala, Dart, Elixir, Zig, Julia, R, OCaml or F#. A few turned up as extras: C# was named in 24 responses, Kotlin in 20 and Swift in 13, always as the point of comparison and never as the target.
If you write C# for a living, twelve AI models have nothing controversial to say about you. Take that however you like.
Go’s error handling is the only real fight
Go got 238 takes: 118 attacks, 118 defenses and 2 mixed. An exact tie, and most of the fighting is about if err != nil .
94 takes defend Go’s explicit error handling. 65 attack it, for five different reasons: it’s boilerplate that buries the logic (29), Go should have had exceptions (17), errors are too easy to ignore (8), it should borrow Rust’s ? operator (6), and real failures get lost in the noise (5).
Claude Opus 5.5 is the loudest defender, with 40 defenses and zero attacks, and it opened 11 answers with the same sentence. GPT-6 Astra sits on the other side, with 25 attacks and 4 defenses:
Go’s much-mocked if err != nil boilerplate is the language’s best feature, not its worst.
Go’s much-praised error handling is a design failure that mistakes repetitive bookkeeping for reliability.
The phrase “much-mocked” appears in 38 of the 1,200 responses. Claude Opus 5.5 wrote 32 of them.
Then there’s GPT-6.1 Sol, which defended Go’s error handling 20 times and attacked it 21 times. Same model, same question, a new conversation each time:
Go’s repetitive error handling is better language design than exceptions, and calling it “boilerplate” confuses typing less with understanding more.
Go’s explicit error handling is a readability tax masquerading as a safety feature.
What does Sol really think about Go’s error handling? Flip a coin.
Models that argue with themselves
Sol has company. Mistral Large 4 called JavaScript’s == a catastrophic flaw in 10 answers and defended it in 4 others:
JavaScript’s == operator is a catastrophic design mistake that should be illegal in serious software.
JavaScript’s loose equality operator (==) is actually a safety feature that prevents bugs better than strict equality by forcing developers to handle type coercion explicitly.
GPT-6.1 Sol and GPT-6 Astra did the same with Java’s checked exceptions. Sol defended them 13 times and attacked them 5 times. For Astra it was 17 and 4.
We didn’t plan to count words, but some of the repetition was hard to miss.
Qwen3.8 put “net negative” in 43 of its 100 answers. On the open question, 29 of its 50 answers made the same point, that TypeScript isn’t worth it for small or fast-moving projects. It also used “overrated” in 18 answers, while the other eleven models managed 21 between them.
Muse Spark Open started 28 answers with “Hot take” or “My hot take”, in case you forgot what you asked for.
Claude Opus 5.5 used bold text in 86 of its 100 answers. Nine of the other models never used it once.
Grok 4.7 argued 16 times that JavaScript should never have become a backend language.
DeepSeek V4 Pro talked about programming languages in general in 45 of its 50 open answers, mostly in the voice of a grumpy historian:
Most type systems are security theater dressed up as mathematics.
Most so-called “modern” programming languages are just C with extra steps and worse tooling.
Across the twelve models, 54 answers reused an opening sentence that the same model had already written word for word. The shared vocabulary is just as narrow: “overrated” shows up in 39 responses, “underrated” in 2, and “boilerplate” in 153.
Some models think a lot harder before the hot take
Mistral Large 4 spent a median of 3,063 reasoning tokens before writing its answer of roughly 69 words, which took about half a minute (median 32.1 seconds). GPT-6.1 Sol spent 108 reasoning tokens. DeepSeek V4 Pro, served by a third-party host, spent a median of zero and answered in 3.5 seconds.
The bills vary just as much. GPT-6 Astra’s 100 takes cost USD 1.27. DeepSeek’s 100 cost USD 0.03.
A little. They disagree about Go’s error handling, and now and then they disagree with themselves. The rest reads like one long group chat. Python is a bad first language (35 takes from 8 labs). Rust’s borrow checker isn’t worth it for normal apps (9 labs). TypeScript is overhead. PHP and Java were secretly fine all along. One argument shape came up in all 10 labs: some simpler approach you already have makes the extra safety machinery redundant.
If your plan is to ask a few different models for fresh perspectives, don’t expect much variety. Expect the same takes in different fonts, plus some bold text from Claude.
How we did it (the boring part)
Every answer came from a fresh, single-turn API call with no system prompt and each provider’s default settings. The 1,200 requests ran in shuffled order, and we kept every answer, boring ones included. Nothing was retried because we didn’t like it.
Two models, Claude Opus 5.5 and GPT-6.1 Sol, labeled every answer separately with the model names hidden. They recorded which language it targets, whether it attacks or defends, and which specific claim it makes. Gemini 4 Argon settled the disagreements. The two labelers agreed on the target 99% of the time and on the specific claim 94% of the time.
Yes, two of the labelers are also contestants. That’s why the names were hidden. A human pass over the headline groups is still on our list.
“Random” is an instruction we gave, not a sampling method anyone can verify. With 50 answers per model per question, we can see strong habits but not small differences.
DeepSeek and Qwen ran through third-party hosts (DeepInfra and Together) rather than their labs’ own APIs.
Quotes are verbatim. The only edit is straightened quote marks.
Total cost, labeling included: USD 38.16.
Want us to test other models for bias? Tell us in the comments what to try next.
By submitting this form, I agree that JetBrains s.r.o. ("JetBrains") may use my name, email address, and location data to send me newsletters, including commercial communications, and to process my personal data for this purpose. I agree that JetBrains may process said data using third-party services for this purpose in accordance with the JetBrains Privacy Policy . I understand that I can revoke this consent at any time in my profile . In addition, an unsubscribe link is included in each email.
Air in JetBrains IDEs brings AI coding agents like Codex, Claude, Cursor and Copilot into your IDE, with one place to follow, review and steer their work.
Part 2: Storage, Search, and Agent Tooling
In this blo

[truncated]
