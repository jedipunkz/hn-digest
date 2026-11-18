---
source: "https://github.com/viperfoundation/Viper/blob/main/on-device-ai.md"
hn_url: "https://news.ycombinator.com/item?id=50553420"
title: "On-device AI: new knowledge without fine-tuning"
article_title: "Viper/on-device-ai.md at main · viperfoundation/Viper · GitHub"
image: "https://opengraph.githubassets.com/d46916e6b76a5f2cf9a29ef4ede581d20ba5f5c5e18a8ce2e7837a9e6178a343/viperfoundation/Viper"
author: "designervoid"
captured_at: "2026-11-18T23:17:41Z"
capture_tool: "hn-digest"
hn_id: 50553420
score: 1
comments: 1
posted_at: "2026-11-18T22:44:39Z"
tags:
  - hacker-news
---

# On-device AI: new knowledge without fine-tuning

- HN: [50553420](https://news.ycombinator.com/item?id=50553420)
- Source: [github.com](https://github.com/viperfoundation/Viper/blob/main/on-device-ai.md)
- Score: 1
- Comments: 1
- Posted: 2026-11-18T22:44:39Z

## Translation

Title: On-device AI: new knowledge without fine-tuning
Article title: Viper/on-device-ai.md at main · viperfoundation/Viper · GitHub
Description: TON based messenger. No phone or email registration – only private/public keys or mnemonic phrase. Always secret chats. - Viper/on-device-ai.md at main · viperfoundation/Viper

Article text:
Viper/on-device-ai.md at main · viperfoundation/Viper · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
viperfoundation
Copy path Blame More file actions Blame More file actions Latest commit
History History 191 lines (144 loc) · 8.71 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions On-device AI
The model is fixed. The program around it is not.
Viper has an AI chat that runs entirely on your Mac. No cloud. No AI provider
account. Your questions never leave the device.
It knows Viper. Ask how keys work, how to move to a new phone, what TON Drop
does — it answers from Viper's own guides and shows where the answer came from.
It can act, too. Ask it to message someone, create an event or save a note, and
it prepares a draft. Nothing is sent until you tap Send .
A stock open model. Viper's harness. Your device.
What you see. AI Chat in Viper on a Mac, offline. Three everyday questions:
what Viper is, how to turn on Face ID login, and what the clock and the single
check next to a message mean. Each answer comes with Sources : the in-app
guides it was taken from ( Enter via Face ID , Message marks , About Viper ).
Tap one to open that guide and check the answer yourself.
Why it matters. A messenger's help should not need the internet, an account
with an AI company or a support chat. People ask about keys, delivery marks and
privacy when something feels wrong — often offline, often in a hurry. Viper
answers on the device, from Viper's own guides, and points to the source, so
you can trust the answer or check it in one tap.
Two pieces of work this November set us off:
Sparks of Weight-Space Programming
(MIT CSAIL, Nov 2026): frontier LLMs can read and edit a neural network's
weights much like code. Editing a tiny share of a small model's weights
removed what held it back — but it did not add abilities the model lacked.
BotPsy
(Tether QVAC, Nov 2026): a 0.8B model trained to act through tool calls,
small enough to run on a phone.
Both point at the same goal: a model on your own device that knows your app
and can act in it. We tried both roads for Viper — editing weights and training
a small model on Viper. Neither taught a model Viper's facts reliably. What
worked was leaving the model alone and building the program around it.
We did not retrain the model to teach it Viper. We tried — training facts into
the weights made it worse.
Instead, Viper builds a harness around an unchanged open model: the knowledge
it needs at the moment you ask, the actions it may propose, and checks that run
before every build ships. Swap in a better model tomorrow, and everything Viper
taught it stays.
The model, the knowledge and your questions stay on your Mac.
No API keys, no usage logs on someone else's server.
The AI proposes; you decide. Every action is a draft until you confirm it.
You choose what it knows: turn Viper's built-in knowledge on or off, add your
own notes.
Viper 1.5.34, now in App Store review, ships gpt-oss-20b on Mac with Viper's
knowledge: when you ask about Viper, the app gives the model the About Viper
card and the two in-app guides that match your question, and the answer shows
those guides as Sources (the screenshot above is 1.5.34).
On a small test this worked well: 20 questions about Viper went from 4 right
(the model alone) to 19 right (the model with the guides).
Twenty questions we wrote ourselves is not proof. So we built a large blind
test:
Hundreds of new questions , written by two different authors (another
open model and us), in Russian and English, in the ways people really ask:
short, vague, with typos, slang, a situation ("my old phone broke…"), several
questions at once — and questions about things Viper does not have.
Two blind judges grade every answer against Viper's knowledge. An answer
counts as correct only when both say so, and as invented when either one
finds a made-up feature, menu or number.
Fresh questions decide. A change counts only when it also wins on
questions no version of the harness has seen.
At this scale the 1.5.34 approach fell short. On fresh situation-style
questions it was right about a third of the time, and invented something in
about a third of its answers. Reading the failures showed why: the model often
had the right guide in front of it and still added things that were not there.
Still no training — the model's weights are untouched. Only the program around
it changed:
Guides rewritten for the model. Every guide page became short numbered
steps with the exact menu path from the app's own code, plus plain lines for
what Viper does not have (no groups, no QR scanner, no cloud sync).
The question is taken apart. A multi-part question is split into
sub-questions; each finds its own guide, and the model first picks the guide
lines that answer it.
Answers come from the guide page itself. Instead of retelling the guide
in its own words, the assistant shows the guide pages it picked. When nothing
fits, it says the guides do not cover it, instead of guessing.
Checks in code , not in the prompt: the answer's language must match the
question, a refusal while a matching guide page exists falls back to a normal
answer, and actions stay drafts.
The same model in three setups, on fresh questions no version has seen
(167 Viper questions, two blind judges):
Model alone (no harness) — stock gpt-oss-20b. It knows nothing about
Viper, and it cannot: gpt-oss-20b was released in 2025, and Viper Messenger
appeared on the App Store in January 2026. The model was finished before
Viper existed in public, so everything it knows about Viper has to come from
the harness.
1.5.34 (harness) — now in App Store review: the About Viper card and two
matching guides.
1.5.35 (harness) — the patch that fixes what this test found in 1.5.34:
the changes listed above.
The model alone cannot know Viper: gpt-oss-20b was released in 2025, and
Viper Messenger appeared on the App Store in January 2026 — after the model was
made. Its few "correct" answers are generic
messenger advice that happens to match Viper (tap a photo, then Share → Save
Image) or an honest "I don't know". On 33 questions about things Viper does
not have, it invented an answer 76% of the time; 1.5.34 6%; 1.5.35 3%.
Why these numbers look low. On simple questions 1.5.34 already works, as
the screenshot shows: the answer may not be exact in every detail, but it
points the right way. This test is harder on purpose: varied questions at
scale — slang, typos, situations, several questions at once — and strict
grading: an answer counts as correct only when both judges agree that every
Viper detail in it is true.
On the set we studied the failures on (166 questions; higher, because we fixed
the guides where these answers went wrong): 1.5.34 43% correct, 28% invented;
1.5.35 71% correct, 2% invented. We publish it, but count only the fresh set.
Draft note: 1.5.35 numbers are from the test bench, not yet from the app build.
For now this needs a powerful Mac. Every number on this page was measured on
one laptop:
MacBook Pro, Apple M5 Pro (15-core CPU: 5 super + 10 performance cores;
16-core GPU), 24 GB of memory, macOS 26.5.
The model is gpt-oss-20b (MLX, MXFP4), about 12 GB, loaded fully into
memory.
Speed: the model alone writes about 57 tokens a second and answers a
question in about 2.6 s (median; 5.6 s for the slowest tenth). The new
harness asks the model 3–4 times per question (split, find, pick, answer),
so a full answer takes about 6.4 s (median; 9.9 s for the slowest tenth).
The model itself takes about 12 GB of memory while it runs (13 GB at peak;
11 GB on disk). With macOS and the app on top, 24 GB is comfortable; a 16 GB
Mac is tight. iPhone is not there yet: the largest iPhones have 12 GB in
total, less than the model alone needs.
Pick the right guide page more often — most remaining misses are a missing or
wrong page, not a wrong answer.
Ship the new harness in the next build, then measure speed on real devices.
Publish the test questions and grading rules alongside the results.
Available now on Mac with Apple silicon and 24 GB of memory.
Viper Messenger on the App Store
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

TON based messenger. No phone or email registration – only private/public keys or mnemonic phrase. Always secret chats. - Viper/on-device-ai.md at main · viperfoundation/Viper

Viper/on-device-ai.md at main · viperfoundation/Viper · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
viperfoundation
Copy path Blame More file actions Blame More file actions Latest commit
History History 191 lines (144 loc) · 8.71 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions On-device AI
The model is fixed. The program around it is not.
Viper has an AI chat that runs entirely on your Mac. No cloud. No AI provider
account. Your questions never leave the device.
It knows Viper. Ask how keys work, how to move to a new phone, what TON Drop
does — it answers from Viper's own guides and shows where the answer came from.
It can act, too. Ask it to message someone, create an event or save a note, and
it prepares a draft. Nothing is sent until you tap Send .
A stock open model. Viper's harness. Your device.
What you see. AI Chat in Viper on a Mac, offline. Three everyday questions:
what Viper is, how to turn on Face ID login, and what the clock and the single
check next to a message mean. Each answer comes with Sources : the in-app
guides it was taken from ( Enter via Face ID , Message marks , About Viper ).
Tap one to open that guide and check the answer yourself.
Why it matters. A messenger's help should not need the internet, an account
with an AI company or a support chat. People ask about keys, delivery marks and
privacy when something feels wrong — often offline, often in a hurry. Viper
answers on the device, from Viper's own guides, and points to the source, so
you can trust the answer or check it in one tap.
Two pieces of work this November set us off:
Sparks of Weight-Space Programming
(MIT CSAIL, Nov 2026): frontier LLMs can read and edit a neural network's
weights much like code. Editing a tiny share of a small model's weights
removed what held it back — but it did not add abilities the model lacked.
BotPsy
(Tether QVAC, Nov 2026): a 0.8B model trained to act through tool calls,
small enough to run on a phone.
Both point at the same goal: a model on your own device that knows your app
and can act in it. We tried both roads for Viper — editing weights and training
a small model on Viper. Neither taught a model Viper's facts reliably. What
worked was leaving the model alone and building the program around it.
We did not retrain the model to teach it Viper. We tried — training facts into
the weights made it worse.
Instead, Viper builds a harness around an unchanged open model: the knowledge
it needs at the moment you ask, the actions it may propose, and checks that run
before every build ships. Swap in a better model tomorrow, and everything Viper
taught it stays.
The model, the knowledge and your questions stay on your Mac.
No API keys, no usage logs on someone else's server.
The AI proposes; you decide. Every action is a draft until you confirm it.
You choose what it knows: turn Viper's built-in knowledge on or off, add your
own notes.
Viper 1.5.34, now in App Store review, ships gpt-oss-20b on Mac with Viper's
knowledge: when you ask about Viper, the app gives the model the About Viper
card and the two in-app guides that match your question, and the answer shows
those guides as Sources (the screenshot above is 1.5.34).
On a small test this worked well: 20 questions about Viper went from 4 right
(the model alone) to 19 right (the model with the guides).
Twenty questions we wrote ourselves is not proof. So we built a large blind
test:
Hundreds of new questions , written by two different authors (another
open model and us), in Russian and English, in the ways people really ask:
short, vague, with typos, slang, a situation ("my old phone broke…"), several
questions at once — and questions about things Viper does not have.
Two blind judges grade every answer against Viper's knowledge. An answer
counts as correct only when both say so, and as invented when either one
finds a made-up feature, menu or number.
Fresh questions decide. A change counts only when it also wins on
questions no version of the harness has seen.
At this scale the 1.5.34 approach fell short. On fresh situation-style
questions it was right about a third of the time, and invented something in
about a third of its answers. Reading the failures showed why: the model often
had the right guide in front of it and still added things that were not there.
Still no training — the model's weights are untouched. Only the program around
it changed:
Guides rewritten for the model. Every guide page became short numbered
steps with the exact menu path from the app's own code, plus plain lines for
what Viper does not have (no groups, no QR scanner, no cloud sync).
The question is taken apart. A multi-part question is split into
sub-questions; each finds its own guide, and the model first picks the guide
lines that answer it.
Answers come from the guide page itself. Instead of retelling the guide
in its own words, the assistant shows the guide pages it picked. When nothing
fits, it says the guides do not cover it, instead of guessing.
Checks in code , not in the prompt: the answer's language must match the
question, a refusal while a matching guide page exists falls back to a normal
answer, and actions stay drafts.
The same model in three setups, on fresh questions no version has seen
(167 Viper questions, two blind judges):
Model alone (no harness) — stock gpt-oss-20b. It knows nothing about
Viper, and it cannot: gpt-oss-20b was released in 2025, and Viper Messenger
appeared on the App Store in January 2026. The model was finished before
Viper existed in public, so everything it knows about Viper has to come from
the harness.
1.5.34 (harness) — now in App Store review: the About Viper card and two
matching guides.
1.5.35 (harness) — the patch that fixes what this test found in 1.5.34:
the changes listed above.
The model alone cannot know Viper: gpt-oss-20b was released in 2025, and
Viper Messenger appeared on the App Store in January 2026 — after the model was
made. Its few "correct" answers are generic
messenger advice that happens to match Viper (tap a photo, then Share → Save
Image) or an honest "I don't know". On 33 questions about things Viper does
not have, it invented an answer 76% of the time; 1.5.34 6%; 1.5.35 3%.
Why these numbers look low. On simple questions 1.5.34 already works, as
the screenshot shows: the answer may not be exact in every detail, but it
points the right way. This test is harder on purpose: varied questions at
scale — slang, typos, situations, several questions at once — and strict
grading: an answer counts as correct only when both judges agree that every
Viper detail in it is true.
On the set we studied the failures on (166 questions; higher, because we fixed
the guides where these answers went wrong): 1.5.34 43% correct, 28% invented;
1.5.35 71% correct, 2% invented. We publish it, but count only the fresh set.
Draft note: 1.5.35 numbers are from the test bench, not yet from the app build.
For now this needs a powerful Mac. Every number on this page was measured on
one laptop:
MacBook Pro, Apple M5 Pro (15-core CPU: 5 super + 10 performance cores;
16-core GPU), 24 GB of memory, macOS 26.5.
The model is gpt-oss-20b (MLX, MXFP4), about 12 GB, loaded fully into
memory.
Speed: the model alone writes about 57 tokens a second and answers a
question in about 2.6 s (median; 5.6 s for the slowest tenth). The new
harness asks the model 3–4 times per question (split, find, pick, answer),
so a full answer takes about 6.4 s (median; 9.9 s for the slowest tenth).
The model itself takes about 12 GB of memory while it runs (13 GB at peak;
11 GB on disk). With macOS and the app on top, 24 GB is comfortable; a 16 GB
Mac is tight. iPhone is not there yet: the largest iPhones have 12 GB in
total, less than the model alone needs.
Pick the right guide page more often — most remaining misses are a missing or
wrong page, not a wrong answer.
Ship the new harness in the next build, then measure speed on real devices.
Publish the test questions and grading rules alongside the results.
Available now on Mac with Apple silicon and 24 GB of memory.
Viper Messenger on the App Store
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
