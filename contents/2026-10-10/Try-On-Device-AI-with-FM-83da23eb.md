---
source: "https://eclecticlight.co/2026/10/10/try-on-device-ai-with-fm/"
hn_url: "https://news.ycombinator.com/item?id=50036707"
title: "Try On-Device AI with FM"
article_title: "Try on-device AI with fm – The Eclectic Light Company"
image: "https://eclecticlight.co/wp-content/uploads/2015/01/cropped-eclecticlightlogo-e1421784280911.png?w=200"
author: "frizlab"
captured_at: "2026-10-10T20:29:38Z"
capture_tool: "hn-digest"
hn_id: 50036707
score: 1
comments: 0
posted_at: "2026-10-10T20:16:39Z"
tags:
  - hacker-news
---

# Try On-Device AI with FM

- HN: [50036707](https://news.ycombinator.com/item?id=50036707)
- Source: [eclecticlight.co](https://eclecticlight.co/2026/10/10/try-on-device-ai-with-fm/)
- Score: 1
- Comments: 0
- Posted: 2026-10-10T20:16:39Z

## Translation

Title: Try On-Device AI with FM
Article title: Try on-device AI with fm – The Eclectic Light Company
Description: If you've got an Apple silicon Mac running macOS 27, have some fun playing with on-device AI using the fm command tool. Free and completely private.

Article text:
Try on-device AI with fm – The Eclectic Light Company
Skip to content
The Eclectic Light Company
If your Mac is already running Golden Gate and you’ve been following what’s going with Apple’s AI, you’ll be aware of the emphasis it places on doing as much as possible on-device, without sending your data somewhere else to handle your task. Apple has gone to great lengths to protect your privacy when off-device AI is processed by Private Cloud Compute (PCC), but at some time in the future we’ll be paying for that service. On-device AI will remain free, of course.
This might appear puzzling at present, as even the most trivial of tasks seems to be sent off to PCC to be run using its more heavyweight models including AFM 3 Cloud, as seen in Apple Intelligence Reports (AIR). This article shows how you can start exploring Apple’s on-device AFM 3 Core and its Advanced version when available on your Mac (models with an M3 chip or later and at least 12 GB memory, we’re promised). For this you don’t need Swift programming skills, Xcode or even Python, just a little time and patience in Terminal using the new fm command.
fm has its own man page that’s clear and informative but doesn’t really explain how you can get the most out of its access to local Foundation Models. There’s also a good introductory article by Blake Crosley, although that’s mostly directed at calling fm from Python, as is the presentation from WWDC 2026 he links to.
Before you use fm for the first time, Apple requires you agree to its Legal Notice and Terms by entering
sudo fm license
and entering yes or y .
You can then dip your toes in gently by exploring its limitations, with
fm respond 'What does Disk Utility do?'
or a similar question. These are sent as prompts for the default system model (AFM 3 Core) to respond to.
You’ll quickly discover it doesn’t know much. Ask it the time of day, and it can’t read the system clock. It doesn’t even know which version of macOS you’re running, or the model of Mac. If you want a quick giggle, try the canonical question
fm respond 'How many r's are there in the word "strawberry"?'
and it will answer 2. In fact, it seems to answer 2 to any word you supply, even elephant.
It does get better when asked about macOS, though. It knows a bit about Disk Utility and other common tools, but sometimes responds incorrectly, for example asserting that Storage settings reveals at the bottom of its view how much purgeable disk space there is. Maybe it’s only been trained on Apple’s limited documentation, and hasn’t yet been allowed to scrape this blog.
I found its Chat mode both more successful and valuable for the insights it can give into its ‘reasoning’. Before entering Chat mode it’s worth giving it some carefully prepared instructions, for example using
fm chat --instructions "Answer the person's question. 1. Begin your response with a plan to solve this question. 2. Follow your plan's steps and show your work. 3. Deliver the final answer in 'answer'."
You can extend those, perhaps opening by giving its persona and role, such as
You are an expert Swift programmer.
followed by brief instructions. Don’t give it a prompt much longer than a paragraph or two, as you might exceed its token limit. Remember this is a small on-device model, not ChatGPT.
Your instructions will then open Chat mode, where you can enter your questions. After my instructions to show its ‘reasoning’, I entered
How many days are in the month of February?
to which it responded
To determine the number of days in the month of February, follow these steps:
1. Recall the standard calendar structure: each month has a specific number of days.
2. Note that February varies based on leap years, but the standard count is typically 28 or 29 days.
3. For a common year, February has 28 days; for a leap year, it has 29 days.
Since the question does not specify a year, we use the standard count: February has 28 days in a non-leap year and 29 in a leap year. Generally, it is considered 28 days unless a leap year applies.
Assuming a common year:
answer: 28
When you’re in Chat mode, enter /help to see available commands, or /quit to quit chat. You’ll be informed of the command you can use to resume that conversation, such as
Resume conversation with: fm chat --resume february-days
and you can do that when you wish.
Because this is on-device, nothing about it appears in the AIR. Instead, fm maintains transcripts hidden in ~/.fm/sessions/, in JSON format. Those are lightweight in comparison with AIRs, so should be straightforward to browse using a good text or JSON editor.
Further and more technical information is given in the Foundation Models framework documentation , although that’s aimed more at developers using Swift. Maybe one day soon.
Share on X (Opens in new window)
X
Share on Facebook (Opens in new window)
Facebook
Share on Reddit (Opens in new window)
Reddit
Share on Pinterest (Opens in new window)
Pinterest
Share on Threads (Opens in new window)
Threads
Share on Mastodon (Opens in new window)
Mastodon
Share on Bluesky (Opens in new window)
Bluesky
Email a link to a friend (Opens in new window)
Email
Print (Opens in new window)
Print
1
Alan B
on October 10, 2026 at 8:22 am
Reply
Yes just a bit of fun currently! I needed to use CTRL C twice to exit completely.
2
Jonathan S
on October 10, 2026 at 8:51 am
Reply
“Maybe it’s only been trained on Apple’s limited documentation, and hasn’t yet been allowed to scrape this blog.” Priceless. :o)
3
hoakley
on October 10, 2026 at 9:30 am
Reply
Thank you for spotting my little sarcastic jibe.
Howard.
4
Scaldis
on October 10, 2026 at 10:49 am
Reply
Maybe you should change the “No AI content” statement in the header into “Just a little bit of AI generated content” lol
5
Tristan Hubsch
on October 10, 2026 at 1:30 pm
Reply
“Maybe one day soon.” — wasn’t “real soon, now” the hallmark μ$oft catchphrase?
All seriousness aside, oh what fun it is to burn through teraflops and watts to get an answer to “Are we there yet?” </s>
Meanwhile, elsewhere from the AIverse (yes, I know this compares  to 🍊s): “ 372 [math research] result families in 719 manuscripts ,” which is but a “first batch,” the serious multilayered importance of which prompted a concerned comment from the Fields medalist Terrence Tao and (not so hushed) chatter among folks who do understand at least some of those 719 manuscrips…
SilentKnight, Skint, SystHist, silnite, LockRattler & Scrub
AIRer, XProCheck, T2M2, LogUI, Ulbow, blowhole and log utilities
Mints: a multifunction utility
xattred, SpotTest, Providable, Spotcord, Metamer & xattr tools
Precize, Alifix, UTIutility, Sparsity, alisma, Taccy, Signet
Spundle, Cormorant, Stibium, DropSum, Dintch, Fintch and cintch
Virtualisation on Apple silicon
Text Utilities: Textovert, Disclipper, Nalaprop, Dystextia and others
Begin typing your search above and press return to search. Press Esc to cancel.
Subscribe
Subscribed
The Eclectic Light Company
Join 9,408 other subscribers
Sign me up
Have a WordPress.com account? Log in now.
The Eclectic Light Company Copy shortlink View post in Reader Manage subscriptions Sign up Log in Report this content Collapse this bar

## Original Extract

If you've got an Apple silicon Mac running macOS 27, have some fun playing with on-device AI using the fm command tool. Free and completely private.

Try on-device AI with fm – The Eclectic Light Company
Skip to content
The Eclectic Light Company
If your Mac is already running Golden Gate and you’ve been following what’s going with Apple’s AI, you’ll be aware of the emphasis it places on doing as much as possible on-device, without sending your data somewhere else to handle your task. Apple has gone to great lengths to protect your privacy when off-device AI is processed by Private Cloud Compute (PCC), but at some time in the future we’ll be paying for that service. On-device AI will remain free, of course.
This might appear puzzling at present, as even the most trivial of tasks seems to be sent off to PCC to be run using its more heavyweight models including AFM 3 Cloud, as seen in Apple Intelligence Reports (AIR). This article shows how you can start exploring Apple’s on-device AFM 3 Core and its Advanced version when available on your Mac (models with an M3 chip or later and at least 12 GB memory, we’re promised). For this you don’t need Swift programming skills, Xcode or even Python, just a little time and patience in Terminal using the new fm command.
fm has its own man page that’s clear and informative but doesn’t really explain how you can get the most out of its access to local Foundation Models. There’s also a good introductory article by Blake Crosley, although that’s mostly directed at calling fm from Python, as is the presentation from WWDC 2026 he links to.
Before you use fm for the first time, Apple requires you agree to its Legal Notice and Terms by entering
sudo fm license
and entering yes or y .
You can then dip your toes in gently by exploring its limitations, with
fm respond 'What does Disk Utility do?'
or a similar question. These are sent as prompts for the default system model (AFM 3 Core) to respond to.
You’ll quickly discover it doesn’t know much. Ask it the time of day, and it can’t read the system clock. It doesn’t even know which version of macOS you’re running, or the model of Mac. If you want a quick giggle, try the canonical question
fm respond 'How many r's are there in the word "strawberry"?'
and it will answer 2. In fact, it seems to answer 2 to any word you supply, even elephant.
It does get better when asked about macOS, though. It knows a bit about Disk Utility and other common tools, but sometimes responds incorrectly, for example asserting that Storage settings reveals at the bottom of its view how much purgeable disk space there is. Maybe it’s only been trained on Apple’s limited documentation, and hasn’t yet been allowed to scrape this blog.
I found its Chat mode both more successful and valuable for the insights it can give into its ‘reasoning’. Before entering Chat mode it’s worth giving it some carefully prepared instructions, for example using
fm chat --instructions "Answer the person's question. 1. Begin your response with a plan to solve this question. 2. Follow your plan's steps and show your work. 3. Deliver the final answer in 'answer'."
You can extend those, perhaps opening by giving its persona and role, such as
You are an expert Swift programmer.
followed by brief instructions. Don’t give it a prompt much longer than a paragraph or two, as you might exceed its token limit. Remember this is a small on-device model, not ChatGPT.
Your instructions will then open Chat mode, where you can enter your questions. After my instructions to show its ‘reasoning’, I entered
How many days are in the month of February?
to which it responded
To determine the number of days in the month of February, follow these steps:
1. Recall the standard calendar structure: each month has a specific number of days.
2. Note that February varies based on leap years, but the standard count is typically 28 or 29 days.
3. For a common year, February has 28 days; for a leap year, it has 29 days.
Since the question does not specify a year, we use the standard count: February has 28 days in a non-leap year and 29 in a leap year. Generally, it is considered 28 days unless a leap year applies.
Assuming a common year:
answer: 28
When you’re in Chat mode, enter /help to see available commands, or /quit to quit chat. You’ll be informed of the command you can use to resume that conversation, such as
Resume conversation with: fm chat --resume february-days
and you can do that when you wish.
Because this is on-device, nothing about it appears in the AIR. Instead, fm maintains transcripts hidden in ~/.fm/sessions/, in JSON format. Those are lightweight in comparison with AIRs, so should be straightforward to browse using a good text or JSON editor.
Further and more technical information is given in the Foundation Models framework documentation , although that’s aimed more at developers using Swift. Maybe one day soon.
Share on X (Opens in new window)
X
Share on Facebook (Opens in new window)
Facebook
Share on Reddit (Opens in new window)
Reddit
Share on Pinterest (Opens in new window)
Pinterest
Share on Threads (Opens in new window)
Threads
Share on Mastodon (Opens in new window)
Mastodon
Share on Bluesky (Opens in new window)
Bluesky
Email a link to a friend (Opens in new window)
Email
Print (Opens in new window)
Print
1
Alan B
on October 10, 2026 at 8:22 am
Reply
Yes just a bit of fun currently! I needed to use CTRL C twice to exit completely.
2
Jonathan S
on October 10, 2026 at 8:51 am
Reply
“Maybe it’s only been trained on Apple’s limited documentation, and hasn’t yet been allowed to scrape this blog.” Priceless. :o)
3
hoakley
on October 10, 2026 at 9:30 am
Reply
Thank you for spotting my little sarcastic jibe.
Howard.
4
Scaldis
on October 10, 2026 at 10:49 am
Reply
Maybe you should change the “No AI content” statement in the header into “Just a little bit of AI generated content” lol
5
Tristan Hubsch
on October 10, 2026 at 1:30 pm
Reply
“Maybe one day soon.” — wasn’t “real soon, now” the hallmark μ$oft catchphrase?
All seriousness aside, oh what fun it is to burn through teraflops and watts to get an answer to “Are we there yet?” </s>
Meanwhile, elsewhere from the AIverse (yes, I know this compares  to 🍊s): “ 372 [math research] result families in 719 manuscripts ,” which is but a “first batch,” the serious multilayered importance of which prompted a concerned comment from the Fields medalist Terrence Tao and (not so hushed) chatter among folks who do understand at least some of those 719 manuscrips…
SilentKnight, Skint, SystHist, silnite, LockRattler & Scrub
AIRer, XProCheck, T2M2, LogUI, Ulbow, blowhole and log utilities
Mints: a multifunction utility
xattred, SpotTest, Providable, Spotcord, Metamer & xattr tools
Precize, Alifix, UTIutility, Sparsity, alisma, Taccy, Signet
Spundle, Cormorant, Stibium, DropSum, Dintch, Fintch and cintch
Virtualisation on Apple silicon
Text Utilities: Textovert, Disclipper, Nalaprop, Dystextia and others
Begin typing your search above and press return to search. Press Esc to cancel.
Subscribe
Subscribed
The Eclectic Light Company
Join 9,408 other subscribers
Sign me up
Have a WordPress.com account? Log in now.
The Eclectic Light Company Copy shortlink View post in Reader Manage subscriptions Sign up Log in Report this content Collapse this bar
