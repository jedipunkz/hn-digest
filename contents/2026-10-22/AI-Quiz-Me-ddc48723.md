---
source: "https://thoughts.kristiandupont.com/p/ai-quiz-me"
hn_url: "https://news.ycombinator.com/item?id=50198959"
title: "AI, Quiz Me!"
article_title: "AI, Quiz Me! - by Kristian Dupont - stdout"
image: "https://substackcdn.com/image/fetch/$s_!yFAs!,f_auto,q_auto:best,fl_progressive:steep/https%3A%2F%2Fkristiandupont.substack.com%2Ftwitter%2Fsubscribe-card.jpg%3Fv%3D1177353315%26version%3D9"
author: "kristiandupont"
captured_at: "2026-10-22T18:20:33Z"
capture_tool: "hn-digest"
hn_id: 50198959
score: 2
comments: 0
posted_at: "2026-10-22T18:05:09Z"
tags:
  - hacker-news
---

# AI, Quiz Me!

- HN: [50198959](https://news.ycombinator.com/item?id=50198959)
- Source: [thoughts.kristiandupont.com](https://thoughts.kristiandupont.com/p/ai-quiz-me)
- Score: 2
- Comments: 0
- Posted: 2026-10-22T18:05:09Z

## Translation

Title: AI, Quiz Me!
Article title: AI, Quiz Me! - by Kristian Dupont - stdout
Description: I’ve noticed my desire to write about programming in this age of AI deteriorating lately.

Article text:
AI, Quiz Me! - by Kristian Dupont - stdout
stdout
Subscribe Sign in AI, Quiz Me!
Kristian Dupont Oct 22, 2026 Share I’ve noticed my desire to write about programming in this age of AI deteriorating lately. Not because interesting things aren’t happening, certainly not, but because it seems that the insights I have are being had simultaneously by everybody else.
I’m going to ignore that sentiment and write something anyway. Partly because I want to be a write, not a write-not , and partly because my brain rewards me when I publish something, and I miss that feeling.
Thanks for reading stdout! Subscribe for free to receive new posts and support my work.
So let me tell you about this Claude skill that I’ve been enjoying lately. Ever since my programming process shifted to writing primarily-English-occasionally-curly-brackets rather than the other way around, I’ve had this sinking feeling that I am losing touch with whatever codebase I am working on. I’ve experienced jumps in abstraction before. I wrote x86 assembly for a good while, then C++ and then the languages that carry so much machinery around with them that it would have made assembly-me squirm. Embracing a higher level of abstraction means giving up some control but in return opening up a world of new opportunities. I get that.
But something feels different this time.
I guess first and foremost because an agent is not deterministic the way a compiler is. But also because it feels like this time, I get to decide how much control I am willing to give up. The agent is perfectly happy to accept vague and poorly thought out instructions and quite often do a reasonable job all the same. On the other hand, retaining control is fatiguing! Which means that I am now in yet another battle of will power, like I am in so many other aspects of life, like my sweet tooth. Dangerous! I need make an effort, I think, to keep an understanding of how every codebase I work with works.
So I tried to get Claude to quiz me. This has worked so well that I have put a notice in my calendar to run it every Friday, and I think it’s really valuable.
---
name: codebase-quiz
description: Quiz the user on the current repo to find where their mental model of the code is wrong. Use when the user asks to be quizzed or tested on the codebase (”quiz me”, “codebase quiz”, “Friday quiz”).
---
# Codebase quiz
Goal: find where my mental model of this codebase is wrong, not teach me. Don’t explain
anything before I answer.
## Storage
`~/.claude/codebase-quiz/<repo>/`, where `<repo>` is the repository name from
`git remote get-url origin` (not the local path, so multiple clones share one history).
- `profile.md` — tenure, areas (name, paths, my 1–3 rating, current difficulty, questions asked
so far), session length, onboarding date, top-level structure snapshot.
- `log.md` — one dated section per session.
## Setup (don’t narrate it)
### If `profile.md` is missing → onboarding
1. Survey the repo: top-level apps/packages, CLAUDE.md files, rules, docs. Propose 8–12 areas,
each with the paths it covers.
2. For each area, count my commits: `git log --author=”$(git config user.email)” --oneline -- <paths> | wc -l`.
Find my first commit date as proposed tenure.
3. Show the table (area, paths, my commits) and ask me to: correct tenure, edit the area list,
rate each area 1–3, set session length (default 8 questions).
4. Write `profile.md`, including a snapshot of the top-level directory listing.
5. This session is a calibration round: one medium question per area. At the end, record each
area’s starting difficulty from the results, and note where my self-rating and my result
disagree.
### If `profile.md` exists
1. Read `profile.md` and `log.md`. Note days since the last session.
2. If the top-level structure differs from the snapshot, or onboarding is older than 6 months,
propose area-list changes before starting.
3. Run `git log --merges --oneline --since=<last session date>`.
4. Plan <session length> questions:
- re-test past misconceptions (more of these the longer since the last session),
- ~2 on recent merges,
- the rest favouri
[truncated]
Thanks for reading stdout! Subscribe for free to receive new posts and support my work.
Share Discussion about this post Comments Restacks Top Latest Discussions No posts

## Original Extract

I’ve noticed my desire to write about programming in this age of AI deteriorating lately.

AI, Quiz Me! - by Kristian Dupont - stdout
stdout
Subscribe Sign in AI, Quiz Me!
Kristian Dupont Oct 22, 2026 Share I’ve noticed my desire to write about programming in this age of AI deteriorating lately. Not because interesting things aren’t happening, certainly not, but because it seems that the insights I have are being had simultaneously by everybody else.
I’m going to ignore that sentiment and write something anyway. Partly because I want to be a write, not a write-not , and partly because my brain rewards me when I publish something, and I miss that feeling.
Thanks for reading stdout! Subscribe for free to receive new posts and support my work.
So let me tell you about this Claude skill that I’ve been enjoying lately. Ever since my programming process shifted to writing primarily-English-occasionally-curly-brackets rather than the other way around, I’ve had this sinking feeling that I am losing touch with whatever codebase I am working on. I’ve experienced jumps in abstraction before. I wrote x86 assembly for a good while, then C++ and then the languages that carry so much machinery around with them that it would have made assembly-me squirm. Embracing a higher level of abstraction means giving up some control but in return opening up a world of new opportunities. I get that.
But something feels different this time.
I guess first and foremost because an agent is not deterministic the way a compiler is. But also because it feels like this time, I get to decide how much control I am willing to give up. The agent is perfectly happy to accept vague and poorly thought out instructions and quite often do a reasonable job all the same. On the other hand, retaining control is fatiguing! Which means that I am now in yet another battle of will power, like I am in so many other aspects of life, like my sweet tooth. Dangerous! I need make an effort, I think, to keep an understanding of how every codebase I work with works.
So I tried to get Claude to quiz me. This has worked so well that I have put a notice in my calendar to run it every Friday, and I think it’s really valuable.
---
name: codebase-quiz
description: Quiz the user on the current repo to find where their mental model of the code is wrong. Use when the user asks to be quizzed or tested on the codebase (”quiz me”, “codebase quiz”, “Friday quiz”).
---
# Codebase quiz
Goal: find where my mental model of this codebase is wrong, not teach me. Don’t explain
anything before I answer.
## Storage
`~/.claude/codebase-quiz/<repo>/`, where `<repo>` is the repository name from
`git remote get-url origin` (not the local path, so multiple clones share one history).
- `profile.md` — tenure, areas (name, paths, my 1–3 rating, current difficulty, questions asked
so far), session length, onboarding date, top-level structure snapshot.
- `log.md` — one dated section per session.
## Setup (don’t narrate it)
### If `profile.md` is missing → onboarding
1. Survey the repo: top-level apps/packages, CLAUDE.md files, rules, docs. Propose 8–12 areas,
each with the paths it covers.
2. For each area, count my commits: `git log --author=”$(git config user.email)” --oneline -- <paths> | wc -l`.
Find my first commit date as proposed tenure.
3. Show the table (area, paths, my commits) and ask me to: correct tenure, edit the area list,
rate each area 1–3, set session length (default 8 questions).
4. Write `profile.md`, including a snapshot of the top-level directory listing.
5. This session is a calibration round: one medium question per area. At the end, record each
area’s starting difficulty from the results, and note where my self-rating and my result
disagree.
### If `profile.md` exists
1. Read `profile.md` and `log.md`. Note days since the last session.
2. If the top-level structure differs from the snapshot, or onboarding is older than 6 months,
propose area-list changes before starting.
3. Run `git log --merges --oneline --since=<last session date>`.
4. Plan <session length> questions:
- re-test past misconceptions (more of these the longer since the last session),
- ~2 on recent merges,
- the rest favouri
[truncated]
Thanks for reading stdout! Subscribe for free to receive new posts and support my work.
Share Discussion about this post Comments Restacks Top Latest Discussions No posts
