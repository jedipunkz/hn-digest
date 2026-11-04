---
source: "https://theainewsreport.com/2026-11-04-gemini-keep-activity-training-human-review-explained.html"
hn_url: "https://news.ycombinator.com/item?id=50371848"
title: "What Gemini does with saved chats, and what turning off Keep Activity costs"
article_title: "Gemini privacy explained: what Keep Activity sends to AI…"
image: "https://theainewsreport.com/og-2026-11-04-gemini-keep-activity-training-human-review-explained.png"
author: "titusblair"
captured_at: "2026-11-04T22:27:49Z"
capture_tool: "hn-digest"
hn_id: 50371848
score: 1
comments: 0
posted_at: "2026-11-04T21:30:45Z"
tags:
  - hacker-news
---

# What Gemini does with saved chats, and what turning off Keep Activity costs

- HN: [50371848](https://news.ycombinator.com/item?id=50371848)
- Source: [theainewsreport.com](https://theainewsreport.com/2026-11-04-gemini-keep-activity-training-human-review-explained.html)
- Score: 1
- Comments: 0
- Posted: 2026-11-04T21:30:45Z

## Translation

Title: What Gemini does with saved chats, and what turning off Keep Activity costs
Article title: Gemini privacy explained: what Keep Activity sends to AI…
Description: Google's Gemini Privacy Hub says saved chats can train its AI models and that reviewed chats stay for up to three years, even after you delete them. We put the whole data path on one page, marked what the page does not say about p

Article text:
The Morning Wire AI NEWS REPORT
Gemini privacy explained: what Keep Activity sends to AI training and human reviewers, how long reviewed chats stay, and how it compares with Claude's settings
Google's Gemini Privacy Hub says saved chats can train its AI models and that reviewed chats stay for up to three years, even after you delete them. We put the whole data path on one page, marked what the page does not say about paid plans, and lined it up with Claude's documented settings.
With Keep Activity on, Gemini saves your chats and what you share (uploads, screens), and Google uses that activity to improve its services, including training generative AI models.
A subset of chats is read by human reviewers, including Google's service providers. Reviewed chats are kept for up to three years and are not deleted when you delete your activity.
Turning Keep Activity off is the switch. Your chats are then kept with your account for 72 hours, so you lose long-term history. Temporary chats are not used for training.
The hub names no different handling for paid personal Google AI plans. Work and school accounts may fall under different terms.
Claude's consumer plans train only when Model Improvement is on, with 5-year retention for that data versus 30 days when it is off (per Anthropic's Claude Code docs). Commercial plans are not trained on by default.
A Hacker News post on November 4 titled "Gemini trains on chat history, even for subscribers (personal accounts)" reached the front page with over 500 points. The poster quoted Google's own Gemini Apps Privacy Hub and said they had assumed a paid subscription meant their chats stayed out of training.
We read the hub page ourselves. Its privacy questions section is dated October 26, 2026. Everything below quotes or paraphrases that page and one linked help page, unless it says otherwise.
The path of one chat with Keep Activity on
Keep Activity is the setting that decides what happens to your Gemini chats. You manage it in Gemini Apps Activity. With it on, this is the path:
Saved: "Your chats and what you share with Gemini (like uploads and screens) will be saved in your Activity."
Used for training: Google uses this activity to improve its services, "including training generative AI models." The linked help page says Gemini trains on your prompts, its responses, and information about how those responses were made.
Read by people: "A subset of chats are reviewed by human reviewers (including Google's trained service providers)." The page does not say how big that subset is.
Separated from you, partly: "Chats are disconnected from your account before being sent to service providers."
Kept: reviewed chats are kept for up to three years, and they "are not deleted when you delete your activity."
Auto-delete: regular activity is kept for 18 months by default. You can change it to 3 months, 36 months, or keep it forever.
The key point is the gap between the last two lines. Deleting your history removes it from your view and from the 18-month store. It does not pull back a chat that a reviewer already has.
What turning Keep Activity off does, and what it costs
Turning Keep Activity off stops future chats from being used to improve Google AI. The help page says so directly: turn it off "to stop your future chats from being used to improve Google AI."
The cost is your history. With Keep Activity off, and for temporary chats, Google keeps your chats with your account for 72 hours. The hub says temporary chats are not used to train Google's AI models. Google still uses chats to respond to you and to protect users, "including with help from human reviewers."
So Gemini ties the training switch to the history feature. You cannot keep a long history of personal chats and also keep them out of training. That is the trade the Hacker News poster objected to.
Note the word future. Turning it off now does not reach back. Chats already reviewed stay for their three years.
Voice and video: "Your audio and Gemini Live videos and screenshares aren't used to improve Google services by default." If you turn on the separate audio setting, some of them may be reviewed by people.
Connected apps: the help page says Gemini does not train generative AI models directly on your Gmail inbox, Drive, Contacts, Calendar, other Workspace apps, Google Wallet, or Google Photos imagery or audio. What it does train on is the conversation: your prompt and its answer, which can include text pulled from those apps.
Screen automation is the exception the other way: screenshots taken when Gemini acts on your screen "are reviewed by trained reviewers and used to improve Google services if Keep Activity is on."
Gemini in Chrome: page content is logged to your account temporarily and does not appear in Gemini Apps Activity.
The hub mentions paid plans in one place: if you pay for a Google AI plan, Google collects "subscription related information." It describes no other difference for paid personal accounts. We found no line that says a paid plan changes training or human review.
Work and school accounts are different on paper: "If you have a work or school Google Account, your use of Gemini Apps may be subject to different data handling terms." The hub points to a separate Workspace privacy hub. We did not read that page today, so we make no claim about its terms.
For an IT team or an MSP, the practical line is simple. Client and company data belongs in a managed work account under terms your organization agreed to, not in someone's personal Gemini, paid or not.
How Claude's documented settings compare
We checked Anthropic's own pages for the same questions. We could not load OpenAI's help pages today (they returned an error to our fetcher), so ChatGPT is left out rather than guessed.
Question Gemini (personal) Claude (Free, Pro, Max, Ultra)
Training switch Keep Activity Model Improvement
Trains when switch is on Yes Yes
Private mode excluded Temporary chats not used Incognito chats not used
Retention, switch on 18 months default activity; 5 years (Claude Code docs)
reviewed chats up to 3 years
Retention, switch off 72 hours 30 days (Claude Code docs)
Work / commercial accounts "may be subject to different Team, Enterprise, API: not trained
data handling terms" unless the customer opts in Two cautions on this table. The Claude retention figures come from Anthropic's Claude Code data usage page, which describes Claude Code data; we did not find the same numbers stated for the Claude apps on the page we read. And Anthropic's privacy article does not state a default for Model Improvement.
The structural difference is where the switch lives. Gemini's training switch is also its history switch. Anthropic's pages tie the switch to how long data is kept, not to whether you see your history.
Open Gemini Apps Activity on every personal Google account that touches work and decide on Keep Activity on purpose.
If you keep it on, set auto-delete to 3 months. It does not reach reviewed chats, but it shrinks everything else.
Use temporary chats for anything you would not want a reviewer to read. Google's own advice is not to enter confidential information you would not want a reviewer to see.
Move client work to managed work accounts, and read your Workspace terms there.
Do the same check in Claude's Privacy Settings and any other chat app your team uses.
Decide on Keep Activity for every personal Google account used near work data.
Set Gemini auto-delete to 3 months if you leave Keep Activity on.
Use temporary chats for anything sensitive.
Keep client data in managed work accounts.
Check Model Improvement in Claude's Privacy Settings too.
How large the human-reviewed subset is. The hub does not give a share or a count.
Whether any paid tier has separate contract terms elsewhere. The hub is silent, and we did not find another Google page that says so.
What the Workspace privacy hub says for work accounts. We did not read it today.
ChatGPT's current settings. OpenAI's help pages returned an error to our fetcher, so we did not compare them.
Whether Claude app chats follow the same 5-year and 30-day figures. The figures we quote are from Anthropic's Claude Code page.
Whether a reviewed chat, once disconnected from an account, could still be linked back to a person by its content. The hub does not discuss this.
Google, Gemini Apps Privacy Hub (support.google.com/gemini/answer/13594961, page dated October 26, 2026), raised on Hacker News on November 4, 2026 — the original report
Google, Gemini help: connected apps and Keep Activity
Anthropic, Claude Code data usage (training and retention by plan)
Anthropic Privacy Center, Is my data used for model training?
Hacker News, Tell HN: Gemini trains on chat history, even for subscribers (item 50363461)

## Original Extract

Google's Gemini Privacy Hub says saved chats can train its AI models and that reviewed chats stay for up to three years, even after you delete them. We put the whole data path on one page, marked what the page does not say about p

The Morning Wire AI NEWS REPORT
Gemini privacy explained: what Keep Activity sends to AI training and human reviewers, how long reviewed chats stay, and how it compares with Claude's settings
Google's Gemini Privacy Hub says saved chats can train its AI models and that reviewed chats stay for up to three years, even after you delete them. We put the whole data path on one page, marked what the page does not say about paid plans, and lined it up with Claude's documented settings.
With Keep Activity on, Gemini saves your chats and what you share (uploads, screens), and Google uses that activity to improve its services, including training generative AI models.
A subset of chats is read by human reviewers, including Google's service providers. Reviewed chats are kept for up to three years and are not deleted when you delete your activity.
Turning Keep Activity off is the switch. Your chats are then kept with your account for 72 hours, so you lose long-term history. Temporary chats are not used for training.
The hub names no different handling for paid personal Google AI plans. Work and school accounts may fall under different terms.
Claude's consumer plans train only when Model Improvement is on, with 5-year retention for that data versus 30 days when it is off (per Anthropic's Claude Code docs). Commercial plans are not trained on by default.
A Hacker News post on November 4 titled "Gemini trains on chat history, even for subscribers (personal accounts)" reached the front page with over 500 points. The poster quoted Google's own Gemini Apps Privacy Hub and said they had assumed a paid subscription meant their chats stayed out of training.
We read the hub page ourselves. Its privacy questions section is dated October 26, 2026. Everything below quotes or paraphrases that page and one linked help page, unless it says otherwise.
The path of one chat with Keep Activity on
Keep Activity is the setting that decides what happens to your Gemini chats. You manage it in Gemini Apps Activity. With it on, this is the path:
Saved: "Your chats and what you share with Gemini (like uploads and screens) will be saved in your Activity."
Used for training: Google uses this activity to improve its services, "including training generative AI models." The linked help page says Gemini trains on your prompts, its responses, and information about how those responses were made.
Read by people: "A subset of chats are reviewed by human reviewers (including Google's trained service providers)." The page does not say how big that subset is.
Separated from you, partly: "Chats are disconnected from your account before being sent to service providers."
Kept: reviewed chats are kept for up to three years, and they "are not deleted when you delete your activity."
Auto-delete: regular activity is kept for 18 months by default. You can change it to 3 months, 36 months, or keep it forever.
The key point is the gap between the last two lines. Deleting your history removes it from your view and from the 18-month store. It does not pull back a chat that a reviewer already has.
What turning Keep Activity off does, and what it costs
Turning Keep Activity off stops future chats from being used to improve Google AI. The help page says so directly: turn it off "to stop your future chats from being used to improve Google AI."
The cost is your history. With Keep Activity off, and for temporary chats, Google keeps your chats with your account for 72 hours. The hub says temporary chats are not used to train Google's AI models. Google still uses chats to respond to you and to protect users, "including with help from human reviewers."
So Gemini ties the training switch to the history feature. You cannot keep a long history of personal chats and also keep them out of training. That is the trade the Hacker News poster objected to.
Note the word future. Turning it off now does not reach back. Chats already reviewed stay for their three years.
Voice and video: "Your audio and Gemini Live videos and screenshares aren't used to improve Google services by default." If you turn on the separate audio setting, some of them may be reviewed by people.
Connected apps: the help page says Gemini does not train generative AI models directly on your Gmail inbox, Drive, Contacts, Calendar, other Workspace apps, Google Wallet, or Google Photos imagery or audio. What it does train on is the conversation: your prompt and its answer, which can include text pulled from those apps.
Screen automation is the exception the other way: screenshots taken when Gemini acts on your screen "are reviewed by trained reviewers and used to improve Google services if Keep Activity is on."
Gemini in Chrome: page content is logged to your account temporarily and does not appear in Gemini Apps Activity.
The hub mentions paid plans in one place: if you pay for a Google AI plan, Google collects "subscription related information." It describes no other difference for paid personal accounts. We found no line that says a paid plan changes training or human review.
Work and school accounts are different on paper: "If you have a work or school Google Account, your use of Gemini Apps may be subject to different data handling terms." The hub points to a separate Workspace privacy hub. We did not read that page today, so we make no claim about its terms.
For an IT team or an MSP, the practical line is simple. Client and company data belongs in a managed work account under terms your organization agreed to, not in someone's personal Gemini, paid or not.
How Claude's documented settings compare
We checked Anthropic's own pages for the same questions. We could not load OpenAI's help pages today (they returned an error to our fetcher), so ChatGPT is left out rather than guessed.
Question Gemini (personal) Claude (Free, Pro, Max, Ultra)
Training switch Keep Activity Model Improvement
Trains when switch is on Yes Yes
Private mode excluded Temporary chats not used Incognito chats not used
Retention, switch on 18 months default activity; 5 years (Claude Code docs)
reviewed chats up to 3 years
Retention, switch off 72 hours 30 days (Claude Code docs)
Work / commercial accounts "may be subject to different Team, Enterprise, API: not trained
data handling terms" unless the customer opts in Two cautions on this table. The Claude retention figures come from Anthropic's Claude Code data usage page, which describes Claude Code data; we did not find the same numbers stated for the Claude apps on the page we read. And Anthropic's privacy article does not state a default for Model Improvement.
The structural difference is where the switch lives. Gemini's training switch is also its history switch. Anthropic's pages tie the switch to how long data is kept, not to whether you see your history.
Open Gemini Apps Activity on every personal Google account that touches work and decide on Keep Activity on purpose.
If you keep it on, set auto-delete to 3 months. It does not reach reviewed chats, but it shrinks everything else.
Use temporary chats for anything you would not want a reviewer to read. Google's own advice is not to enter confidential information you would not want a reviewer to see.
Move client work to managed work accounts, and read your Workspace terms there.
Do the same check in Claude's Privacy Settings and any other chat app your team uses.
Decide on Keep Activity for every personal Google account used near work data.
Set Gemini auto-delete to 3 months if you leave Keep Activity on.
Use temporary chats for anything sensitive.
Keep client data in managed work accounts.
Check Model Improvement in Claude's Privacy Settings too.
How large the human-reviewed subset is. The hub does not give a share or a count.
Whether any paid tier has separate contract terms elsewhere. The hub is silent, and we did not find another Google page that says so.
What the Workspace privacy hub says for work accounts. We did not read it today.
ChatGPT's current settings. OpenAI's help pages returned an error to our fetcher, so we did not compare them.
Whether Claude app chats follow the same 5-year and 30-day figures. The figures we quote are from Anthropic's Claude Code page.
Whether a reviewed chat, once disconnected from an account, could still be linked back to a person by its content. The hub does not discuss this.
Google, Gemini Apps Privacy Hub (support.google.com/gemini/answer/13594961, page dated October 26, 2026), raised on Hacker News on November 4, 2026 — the original report
Google, Gemini help: connected apps and Keep Activity
Anthropic, Claude Code data usage (training and retention by plan)
Anthropic Privacy Center, Is my data used for model training?
Hacker News, Tell HN: Gemini trains on chat history, even for subscribers (item 50363461)
