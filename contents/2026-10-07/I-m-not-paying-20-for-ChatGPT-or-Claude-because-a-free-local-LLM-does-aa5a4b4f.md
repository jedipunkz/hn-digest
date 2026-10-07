---
source: "https://www.xda-developers.com/im-not-paying-20-for-chatgpt-claude-or-gemini-because-a-free-local-llm-does-everything-i-need/"
hn_url: "https://news.ycombinator.com/item?id=49996713"
title: "I'm not paying $20 for ChatGPT or Claude because a free local LLM does"
article_title: "I'm not paying $20 for ChatGPT or Claude because a free local LLM does everything I need"
image: "https://static0.xdaimages.com/wordpress/wp-content/uploads/wm/2026/10/main-fi-qwen-running-on-pc.jpg?w=1600&h=900&fit=crop"
author: "hsnewman"
captured_at: "2026-10-07T18:31:32Z"
capture_tool: "hn-digest"
hn_id: 49996713
score: 4
comments: 0
posted_at: "2026-10-07T18:19:03Z"
tags:
  - hacker-news
---

# I'm not paying $20 for ChatGPT or Claude because a free local LLM does

- HN: [49996713](https://news.ycombinator.com/item?id=49996713)
- Source: [www.xda-developers.com](https://www.xda-developers.com/im-not-paying-20-for-chatgpt-claude-or-gemini-because-a-free-local-llm-does-everything-i-need/)
- Score: 4
- Comments: 0
- Posted: 2026-10-07T18:19:03Z

## Translation

Title: I'm not paying $20 for ChatGPT or Claude because a free local LLM does
Article title: I'm not paying $20 for ChatGPT or Claude because a free local LLM does everything I need
Description: The bills disappeared, but the job gets done all the same

Article text:
I'm not paying $20 for ChatGPT or Claude because a free local LLM does everything I need
Menu
Add Us
Sign in
Sign in now
Close
News
Operating Systems
Submenu
Windows
Devices
Submenu
Single-Board Computers
Entertainment
Submenu
Android Auto
Like
Follow
Followed
12
12
More Action
Sign in
Sign in now
Claude
ESP32
AI Tools
Entertainment
Forums
Close
I'm not paying $20 for ChatGPT or Claude because a free local LLM does everything I need
By
Abhinav Raj
Published Oct 6, 2026, 7:30 AM EDT
Abhinav pivoted from a career in banking to pursue his first love in writing. Even while working full-time, he continued contributing as an editor-at-large, a role he has held for more than 7 years. A lifelong tech enthusiast who has built three gaming and productivity powerhouse PCs since 2018, his passion for technology keeps him closely following the semiconductor industry, from NVIDIA and AMD to ARM. His MSc dissertation explored how artificial intelligence will reshape the future of work, reflecting his curiosity about the wider social impact of emerging technologies.
Sign in to your XDA account
Add Us
Like
Like
follow
Follow
followed
Followed
Thread
12
Log in
Here is a fact-based summary of the story contents:
Try something different:
Show me the facts
Explain it like I’m 5
Give me a lighthearted recap
I've spent some time reflecting on my LLM usage habits, and, inevitably, three names from the biggest cloud AI providers were on my list. ChatGPT and Claude, which together covered a broad chunk of my daily workflow and the menial things that I used to spend a lot of my time on. Each brought something to the table that others don't, which had me convinced for the longest time that there simply wasn't one model capable of replacing all three.
There was a price to this convenience, of course. Both the services offered $20 Pro or Plus plans, and having subscriptions to both of them was eating into my spare change. Naturally, I tried looking for a local model that could handle everything I was throwing at the cloud models, and to my surprise, I found one that could handle most, if not all of it.
Qwen 3.8-27B changed how I saw local AI
It does the job of two subscriptions and costs me nothing
Before using Qwen 3.8-27B , I'd assumed it meant settling for something smaller and possibly less intelligent than a cloud model that I was replacing. Perhaps that's why running Qwen 3.8-27B felt like it defied all expectations I had surrounding local AI. The UD-IQ4_XS build from Unsloth is a 13.3GB download that fits on my 4070 Ti Super (with 16GB of GDDR6X VRAM) with context room to spare at 16K context window, and llama.cpp serves it over localhost at speeds that stop being a complaint after a few minutes of use.
The model holds about 33.7 tokens per second across runs, and in a Pygame test where I asked it to build a working Snake game in a single file, it did exactly that in one shot. Plain text queries give no room for complaint either, as the model is excellent at the sort of ordinary summarization and analysis work that I usually go to a cloud model for.
Think of it as the Linux desktop problem, all over again
ChatGPT was the easiest to replace
Most of what I use it for doesn't need a model running in a data centre
One of the things I use ChatGPT for most often is grammar checking and tidying up my emails, and on the off chance something verbose lands in my inbox, pulling the key details out of it. On rare days, I might even use it to trim down bloated sentences when I get too carried away, or turn a rough thought filled with jargon into something I can communicate more effectively.
I also rely heavily on sentiment analysis when I'm researching products and services. I'll often go through dozens of user forums and discussion threads to understand how consumers are reacting to a particular product, service, or pricing change and then use an LLM to codify that information.
Those are all tasks that a local model can handle without breaking a sweat, and that's something I learned only after a few days of using the Qwen and Gemma family of models. To put this to test, I fed Qwen 3.8-27B a deliberately rambling email about PC component prices with a dozen figures scattered through it. The model was able to pull every number and categorize it into a clean table, and that was enough for me to be convinced that it's a model that can take down a subscription from my bank statement.
Small models are doing more than they should
Claude is brilliant, until it tells me to go away
Five-hour limits are counterproductive for something iterative like coding
Anthropic's frontier models are undoubtedly better at coding, and if you don't believe me, you'd only have to go through a dozen tests I've put them through against Google and OpenAI's models . One might even argue that the rise of Claude changed developer workflows permanently, and it'd be a reasonable argument.
I use it much the same way, particularly when I'm working on Python projects that require a lot of back-and-forth. The problem is that Anthropic's models come with a lot of limits and guardrails, even for customers like myself who would happily pay $20 a month for some extra usage and performance from its top models. It's no surprise that tweaks to get past Claude's rate limits or to mitigate them as far as possible have their own dedicated tutorials these days.
The biggest problem that anyone who relies on Claude faces is that when the five-hour limit runs out, there isn't much you can do besides use your credit card to make it work again, which is a charge on top of a subscription. Waiting for the limit to reset completely destroys the tempo of an iterative coding project, because by the time I'm back (five hours later), I have to go through the thread again and try to pick up where I left off, potentially losing ideas for refining whatever app I was working on.
That's a problem that's almost entirely fixed by running a local model. Qwen running locally has no five-hour wall, no weekly cap, and no overage bill that piles up and stares at me at the end of the month. I can leave it iterating on a project overnight if I want, and the only cost will be the 285W the GPU draws from the wall.
As an added advantage, nothing I send to Qwen leaves my PC. That goes for both my half-written apps and the benchmarking data, none of which touch a server that another organization controls the retention policy for.
llama.cpp
Llama.cpp is an open-source framework that runs large language models locally on your computer.
Qwen models are free, but they are priceless in the right workflow
Qwen 3.8-27B doesn't exactly top the benchmark charts, but for the simple things that I need an LLM for, it absolutely does not have to. Although I found the model that fits the needs of my workflow (and my GPU's VRAM) perfectly, there's a perfect model out there on Hugging Face for every card, from the Gemma family that runs comfortably on my 16GB laptop to smaller Qwen builds that would run on even older hardware. All of them cost nothing besides the hardware they run on, so there's no reason not to tinker and give them a go.
Google is updating how content is shown. Don't miss our industry-leading content, written by humans, by setting XDA as a preferred source .
See more XDA stories on Google.
Add us on Google
Today's best deals
The most important piece of my smart home puzzle is on sale right now; don't miss this deal
AMD's Ryzen 5 7600X drops to $142, a rare discount on this budget gaming powerhouse
Start your self-hosting journey with this discounted NAS, saving you even more in the long run
See More
Trending Now
Intel's cheapest CPUs do what Google's discontinued Coral accelerators used to do, and they cost less
My entire digital life runs on a suite of free tools, and I don't have to spend a dime
Google just launched a $99 Gemini smart speaker, and I've never been more glad I stayed on Home Assistant
Thread
12
Sign in to your XDA account
We want to hear from you. Share your perspective in the comments below, and please keep the conversation respectful.
Attachment(s)
Please respect our community guidelines . No links, inappropriate language, or spam.
Your comment has not been saved
Jarek
Jarek
Jarek
#VN923895
Member since 2024-09-19
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Man, you have just shit articles to write, even 2020 AI is enough for you.
davido
davido
davido
#VF614613
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Good luck leaving Qwen 3.8 running overnight when it constantly gets stuck in reasoning loops. Local models are great, but for very small coding tasks or chat bots only. They do not replace frontier models at all.
Pancho
Pancho
Pancho
#CX089787
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I've literally never had a reasoning loop with qwen3.8 since it came out. What quant are you using?
Joel
Joel
Joel
#EM767626
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Qwen is great but for deep thinking it just won't do.
Ee
Ee
Ee
#GW880897
Member since 2026-02-15
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Looks like you don't need much. To run enogh capable llm, you need so pricy computer that this 20 for a month is a small money.
Peter
Peter
Peter
#YR677049
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Typically this post is getting hate but it makes sense to me. Ill give it a try.
Pancho
Pancho
Pancho
#CX089787
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I'm using the XL version of unsloths Q4 and I did not notice any issues with setting the KV to q8 but it allows for more context (I can imagine that also Q4 works well still, but with q8 I can already use qwens ~200k context fully in VRAM so I didn't try).
I would give it a shot as 3.8 likes to think a lot an a bigger context can be really useful for that.
M
M
M
#XZ321492
Member since 2025-09-19
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Hi. Look up malware called ClosedQuorum. It utilizes Qwen and other AI chatbots without your knowledge. So be careful.
Nihar
Nihar
Nihar
#XV539776
Member since 2026-10-07
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I have a MacBook air with 16gb Unified memory. This heats up with a 8B qwen model itself
Konrad
Konrad
Konrad
#UA862417
Member since 2025-08-01
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Let's host whole internet. Nothing leaves your PC.
SilverFos
SilverFos
SilverFos
#AR782318
Member since 2024-12-21
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
If we could we definitely would
Jacob
Jacob
Jacob
#CF204682
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I switched back to Qwen 3.6 27B on my supermicro x10 configured with 4 Pascal cards. The 3.8 is a beast, but I feel it talks to much, its generating too many tokens. I Nice trick to make them behave even more like opus is using the V model approach. First plan, then review the plan, then implement the plan, then review the implementation , each done in a new separate prompt, this will give you a opus a like result.
I used my old Android phone as a Docker server for a month, and it outperformed my expectations
I paired this Uptime Kuma killer with local LLMs, and monitoring my Proxmox servers has never been easier
When GPT-6 Astra started losing at StarCraft, it just stole the winning bot instead of playing fair
Forget Google Maps and YouTube Music — these 5 apps make Android Auto miles better
XDA
Part of the Valnet Publishing Group
Subscribe
The best of XDA, directly in your inbox.
By subscribing, you agree to receive newsletter and marketing emails, and accept our Terms of Use and Privacy Policy . You can unsubscribe anytime.
Unlock Personalized Content & Exclusive Features

## Original Extract

The bills disappeared, but the job gets done all the same

I'm not paying $20 for ChatGPT or Claude because a free local LLM does everything I need
Menu
Add Us
Sign in
Sign in now
Close
News
Operating Systems
Submenu
Windows
Devices
Submenu
Single-Board Computers
Entertainment
Submenu
Android Auto
Like
Follow
Followed
12
12
More Action
Sign in
Sign in now
Claude
ESP32
AI Tools
Entertainment
Forums
Close
I'm not paying $20 for ChatGPT or Claude because a free local LLM does everything I need
By
Abhinav Raj
Published Oct 6, 2026, 7:30 AM EDT
Abhinav pivoted from a career in banking to pursue his first love in writing. Even while working full-time, he continued contributing as an editor-at-large, a role he has held for more than 7 years. A lifelong tech enthusiast who has built three gaming and productivity powerhouse PCs since 2018, his passion for technology keeps him closely following the semiconductor industry, from NVIDIA and AMD to ARM. His MSc dissertation explored how artificial intelligence will reshape the future of work, reflecting his curiosity about the wider social impact of emerging technologies.
Sign in to your XDA account
Add Us
Like
Like
follow
Follow
followed
Followed
Thread
12
Log in
Here is a fact-based summary of the story contents:
Try something different:
Show me the facts
Explain it like I’m 5
Give me a lighthearted recap
I've spent some time reflecting on my LLM usage habits, and, inevitably, three names from the biggest cloud AI providers were on my list. ChatGPT and Claude, which together covered a broad chunk of my daily workflow and the menial things that I used to spend a lot of my time on. Each brought something to the table that others don't, which had me convinced for the longest time that there simply wasn't one model capable of replacing all three.
There was a price to this convenience, of course. Both the services offered $20 Pro or Plus plans, and having subscriptions to both of them was eating into my spare change. Naturally, I tried looking for a local model that could handle everything I was throwing at the cloud models, and to my surprise, I found one that could handle most, if not all of it.
Qwen 3.8-27B changed how I saw local AI
It does the job of two subscriptions and costs me nothing
Before using Qwen 3.8-27B , I'd assumed it meant settling for something smaller and possibly less intelligent than a cloud model that I was replacing. Perhaps that's why running Qwen 3.8-27B felt like it defied all expectations I had surrounding local AI. The UD-IQ4_XS build from Unsloth is a 13.3GB download that fits on my 4070 Ti Super (with 16GB of GDDR6X VRAM) with context room to spare at 16K context window, and llama.cpp serves it over localhost at speeds that stop being a complaint after a few minutes of use.
The model holds about 33.7 tokens per second across runs, and in a Pygame test where I asked it to build a working Snake game in a single file, it did exactly that in one shot. Plain text queries give no room for complaint either, as the model is excellent at the sort of ordinary summarization and analysis work that I usually go to a cloud model for.
Think of it as the Linux desktop problem, all over again
ChatGPT was the easiest to replace
Most of what I use it for doesn't need a model running in a data centre
One of the things I use ChatGPT for most often is grammar checking and tidying up my emails, and on the off chance something verbose lands in my inbox, pulling the key details out of it. On rare days, I might even use it to trim down bloated sentences when I get too carried away, or turn a rough thought filled with jargon into something I can communicate more effectively.
I also rely heavily on sentiment analysis when I'm researching products and services. I'll often go through dozens of user forums and discussion threads to understand how consumers are reacting to a particular product, service, or pricing change and then use an LLM to codify that information.
Those are all tasks that a local model can handle without breaking a sweat, and that's something I learned only after a few days of using the Qwen and Gemma family of models. To put this to test, I fed Qwen 3.8-27B a deliberately rambling email about PC component prices with a dozen figures scattered through it. The model was able to pull every number and categorize it into a clean table, and that was enough for me to be convinced that it's a model that can take down a subscription from my bank statement.
Small models are doing more than they should
Claude is brilliant, until it tells me to go away
Five-hour limits are counterproductive for something iterative like coding
Anthropic's frontier models are undoubtedly better at coding, and if you don't believe me, you'd only have to go through a dozen tests I've put them through against Google and OpenAI's models . One might even argue that the rise of Claude changed developer workflows permanently, and it'd be a reasonable argument.
I use it much the same way, particularly when I'm working on Python projects that require a lot of back-and-forth. The problem is that Anthropic's models come with a lot of limits and guardrails, even for customers like myself who would happily pay $20 a month for some extra usage and performance from its top models. It's no surprise that tweaks to get past Claude's rate limits or to mitigate them as far as possible have their own dedicated tutorials these days.
The biggest problem that anyone who relies on Claude faces is that when the five-hour limit runs out, there isn't much you can do besides use your credit card to make it work again, which is a charge on top of a subscription. Waiting for the limit to reset completely destroys the tempo of an iterative coding project, because by the time I'm back (five hours later), I have to go through the thread again and try to pick up where I left off, potentially losing ideas for refining whatever app I was working on.
That's a problem that's almost entirely fixed by running a local model. Qwen running locally has no five-hour wall, no weekly cap, and no overage bill that piles up and stares at me at the end of the month. I can leave it iterating on a project overnight if I want, and the only cost will be the 285W the GPU draws from the wall.
As an added advantage, nothing I send to Qwen leaves my PC. That goes for both my half-written apps and the benchmarking data, none of which touch a server that another organization controls the retention policy for.
llama.cpp
Llama.cpp is an open-source framework that runs large language models locally on your computer.
Qwen models are free, but they are priceless in the right workflow
Qwen 3.8-27B doesn't exactly top the benchmark charts, but for the simple things that I need an LLM for, it absolutely does not have to. Although I found the model that fits the needs of my workflow (and my GPU's VRAM) perfectly, there's a perfect model out there on Hugging Face for every card, from the Gemma family that runs comfortably on my 16GB laptop to smaller Qwen builds that would run on even older hardware. All of them cost nothing besides the hardware they run on, so there's no reason not to tinker and give them a go.
Google is updating how content is shown. Don't miss our industry-leading content, written by humans, by setting XDA as a preferred source .
See more XDA stories on Google.
Add us on Google
Today's best deals
The most important piece of my smart home puzzle is on sale right now; don't miss this deal
AMD's Ryzen 5 7600X drops to $142, a rare discount on this budget gaming powerhouse
Start your self-hosting journey with this discounted NAS, saving you even more in the long run
See More
Trending Now
Intel's cheapest CPUs do what Google's discontinued Coral accelerators used to do, and they cost less
My entire digital life runs on a suite of free tools, and I don't have to spend a dime
Google just launched a $99 Gemini smart speaker, and I've never been more glad I stayed on Home Assistant
Thread
12
Sign in to your XDA account
We want to hear from you. Share your perspective in the comments below, and please keep the conversation respectful.
Attachment(s)
Please respect our community guidelines . No links, inappropriate language, or spam.
Your comment has not been saved
Jarek
Jarek
Jarek
#VN923895
Member since 2024-09-19
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Man, you have just shit articles to write, even 2020 AI is enough for you.
davido
davido
davido
#VF614613
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Good luck leaving Qwen 3.8 running overnight when it constantly gets stuck in reasoning loops. Local models are great, but for very small coding tasks or chat bots only. They do not replace frontier models at all.
Pancho
Pancho
Pancho
#CX089787
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I've literally never had a reasoning loop with qwen3.8 since it came out. What quant are you using?
Joel
Joel
Joel
#EM767626
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Qwen is great but for deep thinking it just won't do.
Ee
Ee
Ee
#GW880897
Member since 2026-02-15
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Looks like you don't need much. To run enogh capable llm, you need so pricy computer that this 20 for a month is a small money.
Peter
Peter
Peter
#YR677049
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Typically this post is getting hate but it makes sense to me. Ill give it a try.
Pancho
Pancho
Pancho
#CX089787
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I'm using the XL version of unsloths Q4 and I did not notice any issues with setting the KV to q8 but it allows for more context (I can imagine that also Q4 works well still, but with q8 I can already use qwens ~200k context fully in VRAM so I didn't try).
I would give it a shot as 3.8 likes to think a lot an a bigger context can be really useful for that.
M
M
M
#XZ321492
Member since 2025-09-19
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Hi. Look up malware called ClosedQuorum. It utilizes Qwen and other AI chatbots without your knowledge. So be careful.
Nihar
Nihar
Nihar
#XV539776
Member since 2026-10-07
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I have a MacBook air with 16gb Unified memory. This heats up with a 8B qwen model itself
Konrad
Konrad
Konrad
#UA862417
Member since 2025-08-01
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Let's host whole internet. Nothing leaves your PC.
SilverFos
SilverFos
SilverFos
#AR782318
Member since 2024-12-21
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
If we could we definitely would
Jacob
Jacob
Jacob
#CF204682
Member since 2026-10-06
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I switched back to Qwen 3.6 27B on my supermicro x10 configured with 4 Pascal cards. The 3.8 is a beast, but I feel it talks to much, its generating too many tokens. I Nice trick to make them behave even more like opus is using the V model approach. First plan, then review the plan, then implement the plan, then review the implementation , each done in a new separate prompt, this will give you a opus a like result.
I used my old Android phone as a Docker server for a month, and it outperformed my expectations
I paired this Uptime Kuma killer with local LLMs, and monitoring my Proxmox servers has never been easier
When GPT-6 Astra started losing at StarCraft, it just stole the winning bot instead of playing fair
Forget Google Maps and YouTube Music — these 5 apps make Android Auto miles better
XDA
Part of the Valnet Publishing Group
Subscribe
The best of XDA, directly in your inbox.
By subscribing, you agree to receive newsletter and marketing emails, and accept our Terms of Use and Privacy Policy . You can unsubscribe anytime.
Unlock Personalized Content & Exclusive Features
