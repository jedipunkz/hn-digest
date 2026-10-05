---
source: "https://danunparsed.com/p/giving-code-reviewers-credit"
hn_url: "https://news.ycombinator.com/item?id=49959618"
title: "Fix Your AI Slop Problem by Giving Reviewers Credit"
article_title: "Fix Your AI Slop Problem By Giving Reviewers Credit"
image: "https://substackcdn.com/image/fetch/$s_!TDTB!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Fb216dc35-ae5a-4634-b300-e4b2f064b1b7_1376x768.jpeg"
author: "sambellll"
captured_at: "2026-10-05T01:39:47Z"
capture_tool: "hn-digest"
hn_id: 49959618
score: 1
comments: 0
posted_at: "2026-10-05T01:06:07Z"
tags:
  - hacker-news
---

# Fix Your AI Slop Problem by Giving Reviewers Credit

- HN: [49959618](https://news.ycombinator.com/item?id=49959618)
- Source: [danunparsed.com](https://danunparsed.com/p/giving-code-reviewers-credit)
- Score: 1
- Comments: 0
- Posted: 2026-10-05T01:06:07Z

## Translation

Title: Fix Your AI Slop Problem by Giving Reviewers Credit
Article title: Fix Your AI Slop Problem By Giving Reviewers Credit
Description: Reviewing isn't free labour. When AI writes and reviews your PRs, the human approval stops meaning anything. Here's how to make code review count again.

Article text:
Fix Your AI Slop Problem By Giving Reviewers Credit
Dan Unparsed
Subscribe Sign in Fix Your AI Slop Problem By Giving Reviewers Credit
Three AI passes missed a bug you could find by scrolling up
Dan Apushkinsky Oct 02, 2026 2 1 Share Approved
I was reviewing a PR recently. See if you can spot the bug:
# Branch C doesn't have a printer
user = new Person(...)
... user does some shit ...
if user not in branch_c:
user.print_document()
... 50 lines of library logic ...
### New code added in the PR:
+ staff = new Person(...)
... staff does some shit ...
+ staff.print_document() Obviously I don’t work at a library — I swapped out the real calls. But it was truly that simple. The # Branch C doesn’t have a printer isn’t added here for context either, it’s there, in the code. Any testing we'd do would pass, but the moment we deployed to branch C, we'd start seeing failures.
Maybe a better codebase wouldn’t have made this type of mistake possible. But hey, don’t blame me — production code bases are rarely perfect, and as far as ones I’ve seen go, this one isn’t too bad.
Still, good old blind copy-and-pasting would’ve prevented this. Scrolling up should’ve prevented this.
Instead, an AI coded it up, an AI model reviewed it, a human reviewer pasted it into an AI to review it, and then stamped it.
I was tagged as the second (human) reviewer. This time I happened to read the code, but I’ll admit — most times, it would’ve just pasted the PR into AI and missed the same obvious bug.
What Is the Point of a Reviewer? Actually?
I’ve always been taught that a reviewer owns the code as much as the author. And I’d always try and check a bunch of things: does it work, does it fit, will anyone understand it in a year. But there was one thing I’d never had to check before — whether the author ever read it themselves.
That came guaranteed. You couldn’t have typed staff.print_document() without hitting ctrl+f and finding an example of how it was used. Yet, today, people can ship code they’ve never seen before. The reviewer could plausibly be the first human to actually read the code.
The bug I saw didn’t require any special knowledge to catch. The fact AI even missed it was honestly pretty shocking to me. But at the same time, having a circle of AI’s read the same code, with near identical prompts, often running the same model 1 , and expecting to discover something different makes no sense.
We were mistaking a second (and third and fourth) AI pass for a human review just because it was a human copy and pasting the PR into their favourite AI harness.
At what point should you just cut the human out of the loop and save everyone some time?
I’ve talked with a few teams at various companies over the last year, and more than one has ditched the human entirely. A few would have the AI make a judgment call — is this change risky enough to warrant a human reviewer?
Personally, I’m still a fan of having human reviewers for the vast majority of work. Maybe I’m behind the curve.
Either way, this isn’t a blog post about deciding how much human you want. It’s that once you decide that number, whether it’s 0%, 100%, or something in the middle, then you should then take that direction with intention.
Today, many of us are living in no man’s land. We still have the human reviewer requirement from the past — but the human is just a proxy for yet another AI review.
At first glance a tempting fix might be to ban AI from reviews until the reviewer has at least read it once themselves. But you can't police what someone does in their own terminal, and even if you could, it’s probably not addressing the true root cause.
There are few things more infuriating than reviewing a PR or a doc where you know you’re about to spend more time reading it than the author spent writing it. They spent 5 minutes prompting AI, just for you to spend an hour reviewing it. At what point are you the one building the feature, by proxy?
And these hours that you spend reviewing code — you never get any credit. Pasting PRs into AI isn’t laziness. It’s the logical decision if you prioritize your own time.
In my experience, it’s rare for people to want to game the system. Most people want to do good, honest work and be recognized for it. They just don’t want to do invisible, uncredited labour.
The solution to this problem evaded me for a few weeks, but I actually think it’s quite simple: add reviews as subtasks on your sprint board.
Pre-AI, managing that overhead would’ve been a pain in the ass. These days AI can do it for you — hook it into your PR pipeline, have it create the ticket automatically.
And I know, I know. It’s bad practice to add tasks mid-sprint. But fuck it. Times have changed. Unless you want to slow down shipping — and no one wants that — you’re going to have to drag those tasks into the current sprint.
A review ticket does two things. First, it forces you to estimate it. Sure — it’s probably a 1 pointer, but spending 30 seconds estimating it is a signal that it’s a real thing that we value.
Second, it makes the cost visible. When a 3-point feature spawns a 3-point review, that effort finally shows up side-by-side with the feature work itself. And for the first time ever, the reviewer will no longer be punished for actually doing their due diligence.
Now you would be right to point out — how is adding reviewers to sprint boards a solution? People can still use AI to save time.
The answer is that this is a one-two punch. Giving reviewers credit sets you up to build a team culture where you’re rewarded for spending time on reviews and finding issues early. Adding reviews to sprint boards is putting up one extra wall that has to be broken down for your culture to turn to AI slop.
So give people the recognition they deserve for reviewing PRs properly, and next time someone hits approve, it’ll actually mean something.
Thanks for reading. If this was interesting, subscribe — I'm doing more of these.
Even using different models might not matter much — one 2025 study had found that when two different models get a question wrong, they pick the same wrong answer 60% of the time.
2 1 Share Discussion about this post Comments Restacks Top Latest Discussions No posts

## Original Extract

Reviewing isn't free labour. When AI writes and reviews your PRs, the human approval stops meaning anything. Here's how to make code review count again.

Fix Your AI Slop Problem By Giving Reviewers Credit
Dan Unparsed
Subscribe Sign in Fix Your AI Slop Problem By Giving Reviewers Credit
Three AI passes missed a bug you could find by scrolling up
Dan Apushkinsky Oct 02, 2026 2 1 Share Approved
I was reviewing a PR recently. See if you can spot the bug:
# Branch C doesn't have a printer
user = new Person(...)
... user does some shit ...
if user not in branch_c:
user.print_document()
... 50 lines of library logic ...
### New code added in the PR:
+ staff = new Person(...)
... staff does some shit ...
+ staff.print_document() Obviously I don’t work at a library — I swapped out the real calls. But it was truly that simple. The # Branch C doesn’t have a printer isn’t added here for context either, it’s there, in the code. Any testing we'd do would pass, but the moment we deployed to branch C, we'd start seeing failures.
Maybe a better codebase wouldn’t have made this type of mistake possible. But hey, don’t blame me — production code bases are rarely perfect, and as far as ones I’ve seen go, this one isn’t too bad.
Still, good old blind copy-and-pasting would’ve prevented this. Scrolling up should’ve prevented this.
Instead, an AI coded it up, an AI model reviewed it, a human reviewer pasted it into an AI to review it, and then stamped it.
I was tagged as the second (human) reviewer. This time I happened to read the code, but I’ll admit — most times, it would’ve just pasted the PR into AI and missed the same obvious bug.
What Is the Point of a Reviewer? Actually?
I’ve always been taught that a reviewer owns the code as much as the author. And I’d always try and check a bunch of things: does it work, does it fit, will anyone understand it in a year. But there was one thing I’d never had to check before — whether the author ever read it themselves.
That came guaranteed. You couldn’t have typed staff.print_document() without hitting ctrl+f and finding an example of how it was used. Yet, today, people can ship code they’ve never seen before. The reviewer could plausibly be the first human to actually read the code.
The bug I saw didn’t require any special knowledge to catch. The fact AI even missed it was honestly pretty shocking to me. But at the same time, having a circle of AI’s read the same code, with near identical prompts, often running the same model 1 , and expecting to discover something different makes no sense.
We were mistaking a second (and third and fourth) AI pass for a human review just because it was a human copy and pasting the PR into their favourite AI harness.
At what point should you just cut the human out of the loop and save everyone some time?
I’ve talked with a few teams at various companies over the last year, and more than one has ditched the human entirely. A few would have the AI make a judgment call — is this change risky enough to warrant a human reviewer?
Personally, I’m still a fan of having human reviewers for the vast majority of work. Maybe I’m behind the curve.
Either way, this isn’t a blog post about deciding how much human you want. It’s that once you decide that number, whether it’s 0%, 100%, or something in the middle, then you should then take that direction with intention.
Today, many of us are living in no man’s land. We still have the human reviewer requirement from the past — but the human is just a proxy for yet another AI review.
At first glance a tempting fix might be to ban AI from reviews until the reviewer has at least read it once themselves. But you can't police what someone does in their own terminal, and even if you could, it’s probably not addressing the true root cause.
There are few things more infuriating than reviewing a PR or a doc where you know you’re about to spend more time reading it than the author spent writing it. They spent 5 minutes prompting AI, just for you to spend an hour reviewing it. At what point are you the one building the feature, by proxy?
And these hours that you spend reviewing code — you never get any credit. Pasting PRs into AI isn’t laziness. It’s the logical decision if you prioritize your own time.
In my experience, it’s rare for people to want to game the system. Most people want to do good, honest work and be recognized for it. They just don’t want to do invisible, uncredited labour.
The solution to this problem evaded me for a few weeks, but I actually think it’s quite simple: add reviews as subtasks on your sprint board.
Pre-AI, managing that overhead would’ve been a pain in the ass. These days AI can do it for you — hook it into your PR pipeline, have it create the ticket automatically.
And I know, I know. It’s bad practice to add tasks mid-sprint. But fuck it. Times have changed. Unless you want to slow down shipping — and no one wants that — you’re going to have to drag those tasks into the current sprint.
A review ticket does two things. First, it forces you to estimate it. Sure — it’s probably a 1 pointer, but spending 30 seconds estimating it is a signal that it’s a real thing that we value.
Second, it makes the cost visible. When a 3-point feature spawns a 3-point review, that effort finally shows up side-by-side with the feature work itself. And for the first time ever, the reviewer will no longer be punished for actually doing their due diligence.
Now you would be right to point out — how is adding reviewers to sprint boards a solution? People can still use AI to save time.
The answer is that this is a one-two punch. Giving reviewers credit sets you up to build a team culture where you’re rewarded for spending time on reviews and finding issues early. Adding reviews to sprint boards is putting up one extra wall that has to be broken down for your culture to turn to AI slop.
So give people the recognition they deserve for reviewing PRs properly, and next time someone hits approve, it’ll actually mean something.
Thanks for reading. If this was interesting, subscribe — I'm doing more of these.
Even using different models might not matter much — one 2025 study had found that when two different models get a question wrong, they pick the same wrong answer 60% of the time.
2 1 Share Discussion about this post Comments Restacks Top Latest Discussions No posts
