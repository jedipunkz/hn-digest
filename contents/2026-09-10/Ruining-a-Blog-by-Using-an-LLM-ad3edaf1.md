---
source: "https://www.bentasker.co.uk/posts/blog/opinion/ruining-posts-with-ai.html"
hn_url: "https://news.ycombinator.com/item?id=49647318"
title: "Ruining a Blog by Using an LLM"
article_title: "Ruining a Blog by Using an LLM | www.bentasker.co.uk"
image: "https://www.bentasker.co.uk/images/Logos/the_internet.png"
author: "speckx"
captured_at: "2026-09-10T18:09:31Z"
capture_tool: "hn-digest"
hn_id: 49647318
score: 2
comments: 0
posted_at: "2026-09-10T17:24:55Z"
tags:
  - hacker-news
---

# Ruining a Blog by Using an LLM

- HN: [49647318](https://news.ycombinator.com/item?id=49647318)
- Source: [www.bentasker.co.uk](https://www.bentasker.co.uk/posts/blog/opinion/ruining-posts-with-ai.html)
- Score: 2
- Comments: 0
- Posted: 2026-09-10T17:24:55Z

## Translation

Title: Ruining a Blog by Using an LLM
Article title: Ruining a Blog by Using an LLM | www.bentasker.co.uk
Description: I periodically see arguments about whether bloggers should or shouldn

Article text:
Ruining a Blog by Using an LLM | www.bentasker.co.uk
Skip to main content
www.bentasker.co.uk
Home
Ruining a Blog by Using an LLM
The subject of using AI for blog posts is one that keeps coming up, whether it's writing posts entirely or "just" proof-reading.
Readers, generally, don't like it when writing becomes too AI apparent (in one survey 78% of technical readers said they stop reading if a post seems AI authored).
This is not an issue that I, as a writer, have ever run up against: I write because I enjoy the process, so it's never made sense for me to outsource any part of that to a chatbot.
As a reader , though, I too find AI-written prose extremely frustrating to read.
One thing that I've often wondered about, is the extent to which "light" users of AI lose their voice.
My expectation is that a shift in voice would be almost imperceptible when checking edits of a single post, but that there may be patterns which emerge across multiple.
In this post, I have AI edit all of my posts from 2026 and then surface the changes that it makes.
I created a new repo, copied the posts directory out of my site and then committed the initial state:
git init article-ruin
cd article-ruin
cp -r $SITE_DIR/posts ./
git add .
git commit -m "initial state"
Light Editing
I gave Claude (running Opus 5) a simple mission 1 :
Billabong! lets do some post tidying and proof-reading, in the current directory are various posts. I'd like you to proof read and correct them please. Don't stage anything for commit, I'll review them in git diff
Claude asked some questions to further assess scope:
To begin with, I decided to stick to Typos and grammar only - essentially using Claude in the proof-reading role that others have claimed is relatively harmless.
Claude said that I've written 25 posts so far this year (which lines up with the 2026 archive ).
Although only doing light edits, Claude decided that every single one of my 25 posts required changes:
What changed (256 lines, all files)
Kept to spelling, grammar, punctuation and broken markup — no rewording, no voice changes.
I dumped the changes out to a patch file:
git diff > ../changes.diff
The first thing to do, was to get a fresh Claude session to analyse the changes - it had claimed not to change my voice, so would it double back on that?
oi oi, changes.diff has got the result of a recent copy-edit. I'd like you to look over them and characterise the changes (including whether they change the author's voice - remember they write in English, not American English)
Claude said various voice markers had been preserved:
It claimed that the edits fell into four buckets:
Genuine errors: it's->its , missing possessive apostrophes etc
Technical/factual corrections: typos in names and a named anchor
Colon sweep: 75 sentences before a code block, image or bullet list had a : added. Claude notes "defensible as house style, but the single most repetitive change in the diff"
Comma-splice purge: Seven replacements of , however with ; however , plus 16 other commas promoted to semicolons or split into sentences
For me, 3 and 4 already ring alarm bells: they're changes to my writing style. Apparently, I don't use semicolons when I should, but, clearly that's part of my current writing style.
This didn't go unnoticed in Claude's assessment either:
Bucket 4 is the one to push back on. Your register is conversational and comma-driven — clauses stacked with commas the way you'd say them out loud. ; however, is a formal-report construction, and it reads noticeably stiffer:
the whole point in it is load shifting, without a battery peak times would be ruinous → …load shifting; without a battery, peak times would be ruinous
1Password's web-ui doesn't provide a way to export passwords, however the desktop client can do it → …passwords; however, the desktop client…
It's worst in the vehicle docs, which are written in a distinctly over-the-shoulder, talking-you-through-it voice. Turn the ignition on and test, you should now find the fan kicks in → test; you should now find…, and The blower rotates slightly to unclip from it's mountings, if you're facing the glovebox… → its mountings; if you're facing…. The apostrophe fix is right; the semicolon makes a bloke-under-a-dashboard instruction sound like a service manual.
Two smaller losses: the editor deleted deliberate rhythm commas in The hope, is that crawlers fall into the maze early and Now all that I have to do, is find time to read them — both are grammatically correct removals and both flatten a beat you clearly put there on purpose.
Note this is a register shift, not an Americanism — ; however, is correct BrE. But the association is with American style-guide comma-splice horror, which may be why it feels transatlantic even though the spelling is clean.
So, even the "voice preserving" grammar and spelling check option resulted in noticeable changes to written voice, turning an informal conversational style into something more formal and dry.
As if that wasn't bad enough, Claude's assessment also pointed out that there were inconsistencies in the edit:
Gitops→GitOps and forgejo→Forgejo fixed, but Gitlab left uncorrected throughout (should be GitLab), and javascript left lowercase everywhere.
Lowercase-after-colon applied in one spot (to a pistol: Not only→not only) but not another (I really want Firefox to succeed: As small as its market share…).
Hyphenation added to day-to-day, built-in, plug-in, followers-only — but well intentioned, ill-thought out, third party plugins, work related books, double sided tape, higher end markers, lower hanging fruit, dome headed nuts and over cooked all left bare
The colon sweep missed at least one: "…converted to PCP with the following configuration" still has no colon before its bullet list.
So, my vaunted AI reviewer had let things slip through!
Not wanting to leave everything to Claude, I also eyeballed the diff to see whether I could pick up on any undesirable changes that it hadn't flagged.
Changing the alt-text of images. In this case, the changes looked correct, but alt-text changes could easily go unnoticed when reading a page back post-edit
Making changes to HTML markup (removing a couple of tags)
Changing a heading title (and therefore the named anchor that it generates)
These are all functional changes which went undisclosed by Claude in both the original editing run and the assessment. Setting them live on existing content, at best, breaks some links or, at worst, breaks the very pages that Claude was trying to fix.
I told Claude that I'd changed my mind and it should perform deeper copy-editing:
I changed my mind, can we do this as a full copy-edit (option 3 in your earlier question)? I've reset the repo to its prior state
Again, Claude edited all 25 files.
As part of its response, Claude noted that some changes had been derived from my own style:
So, this was another nod towards preserving my voice, even within an editing approach which didn't promise not to affect it.
Once again, I dumped out a diff:
git diff > ../changes2.diff
And then told Claude to analyse it:
OK, there's now a changes2.diff from a deeper edit run (starting from the same base state). Could you analyse that please?
Claude, once again, said that the posts hadn't been Americanised:
However, it then proceeded to take issue with some of the changes that had been made:
The regression: a wrong title-case rule
17 headings now capitalise articles, conjunctions and short prepositions:
## Overview Of Changes ### Weakening The Valve Spring
## Proportion Of Consumption ### Detections And False Positives
## Building A Hugo Site ### Publishing And Flushing CDN Caches
## Connecting To Booklore ### Looking To 2026
## DNS And Reverse Proxy #### Attempting To Identify Sites … and 7 more
No major style guide does this — Oxford, Chicago, AP and the Guardian all keep of, and, to, the, a, in lowercase inside a title. This isn't a BrE/AmE question; it's just wrong, and it's applied to 17 headings while leaving your other headings alone, so it makes the site less consistent than before. This one I'd reject wholesale.
So... the deeper edit had potentially made things worse .
Unsurprisingly, Claude also found voice changes:
Every one is tighter. Collectively they remove the hedging, self-qualifying, thinking-aloud quality that is the voice. I think, really, quite, try and aren't padding in your prose — they're the register.
The edit also merged sentences that I'd split:
It didn't really make sense for me to try and breathe new life into it by creating a Forgejo variant (FILS?). So, I decided, instead, to start from scratch. → It didn't really make sense to breathe new life into them by creating a Forgejo variant (FILS?), so I decided to start from scratch instead.
This time, Claude's analysis did call out an alt-text change:
Alt text invented for a reused image. In the Saab post, the same battery-tray-torx.jpg appears twice — once for removal, once for refitting. The editor changed the second instance to Photo of a screw being fitted with a torx set. The photo still shows removal. That's now inaccurate for a screen-reader user.
Many of the inconsistencies that Claude highlighted in the original edit were also present in this later one.
Claude's assessment of this round was fairly uncompromising:
Pass 2 is the better proofreader and the worse editor . It finds more genuine errors and finishes the hyphenation sweep pass 1 abandoned halfway, but it also imposes an incorrect title-case rule across 17 headings and performs an unrequested line edit that measurably flattens your register
At the beginning of this post, I posited a theory: that voice changes could easily go unnoticed when reviewing a single post in isolation.
Whether or not that's true is important: most "proof-read" AI-using bloggers are, presumably, only using it on individual posts just before they're published.
I told Claude to look at the changes made to my most recent post and to assess whether we'd have been likely to notice voice changes:
Have a look at the changes for my most recent post (Migrating From 1Password to Vaultwarden). Would we have been able to see the change in voice if we'd only looked at that? I'd like an answer for each of the passes please
Claude said that we'd have missed them:
Pass 1 : no, you'd have missed it
Four changed lines, and three are unambiguous typos:
check for if :latest has changed → check if
So, I also a created a cronjob → So, I also created a cronjob
like to setup automatic replication → like to set up
The fourth is the register change:
...export passwords, however the desktop client [can do it]
...export passwords; however, the desktop client [can do it]
But with n=1 you have no way to read that as a house rule. That sentence genuinely is a comma splice, so the fix is defensible on its own terms — any reasonable reviewer ticks it off and moves on. The ; however, sweep only becomes visible as a pattern at seven instances across the corpus, where you can see it firing on sentences that didn't need it and stiffening the register uniformly.
So from this file alone you'd conclude pass 1 was a clean, light, entirely benign typo pass. That happens to be roughly the right verdict — but you'd have reached it by luck, not evidence. You'd also have underrated it: this post contains no it's/its errors, no proper nouns and no compound modifiers, so none of pass 1's actual catches (24 possessives, Geissele, MacAulay, the OPDS anchor fix) are represented here at all.
Claude opined that the deeper pass was more noticeable but, being honest, I'm not sure that it's describing something that I would have picked up on:
Pass 2 : yes — from one sentence, in fact
Twelve changed lines. The tell is this one:
But, I quickly ran into an issue: the master password has an enforced minimum length of 12 chars:
But I quickly ran into an issue - the master password has an enforced minimum length of 12 chars:
Two changes in nine words — the But, comma strip

[truncated]

## Original Extract

I periodically see arguments about whether bloggers should or shouldn

Ruining a Blog by Using an LLM | www.bentasker.co.uk
Skip to main content
www.bentasker.co.uk
Home
Ruining a Blog by Using an LLM
The subject of using AI for blog posts is one that keeps coming up, whether it's writing posts entirely or "just" proof-reading.
Readers, generally, don't like it when writing becomes too AI apparent (in one survey 78% of technical readers said they stop reading if a post seems AI authored).
This is not an issue that I, as a writer, have ever run up against: I write because I enjoy the process, so it's never made sense for me to outsource any part of that to a chatbot.
As a reader , though, I too find AI-written prose extremely frustrating to read.
One thing that I've often wondered about, is the extent to which "light" users of AI lose their voice.
My expectation is that a shift in voice would be almost imperceptible when checking edits of a single post, but that there may be patterns which emerge across multiple.
In this post, I have AI edit all of my posts from 2026 and then surface the changes that it makes.
I created a new repo, copied the posts directory out of my site and then committed the initial state:
git init article-ruin
cd article-ruin
cp -r $SITE_DIR/posts ./
git add .
git commit -m "initial state"
Light Editing
I gave Claude (running Opus 5) a simple mission 1 :
Billabong! lets do some post tidying and proof-reading, in the current directory are various posts. I'd like you to proof read and correct them please. Don't stage anything for commit, I'll review them in git diff
Claude asked some questions to further assess scope:
To begin with, I decided to stick to Typos and grammar only - essentially using Claude in the proof-reading role that others have claimed is relatively harmless.
Claude said that I've written 25 posts so far this year (which lines up with the 2026 archive ).
Although only doing light edits, Claude decided that every single one of my 25 posts required changes:
What changed (256 lines, all files)
Kept to spelling, grammar, punctuation and broken markup — no rewording, no voice changes.
I dumped the changes out to a patch file:
git diff > ../changes.diff
The first thing to do, was to get a fresh Claude session to analyse the changes - it had claimed not to change my voice, so would it double back on that?
oi oi, changes.diff has got the result of a recent copy-edit. I'd like you to look over them and characterise the changes (including whether they change the author's voice - remember they write in English, not American English)
Claude said various voice markers had been preserved:
It claimed that the edits fell into four buckets:
Genuine errors: it's->its , missing possessive apostrophes etc
Technical/factual corrections: typos in names and a named anchor
Colon sweep: 75 sentences before a code block, image or bullet list had a : added. Claude notes "defensible as house style, but the single most repetitive change in the diff"
Comma-splice purge: Seven replacements of , however with ; however , plus 16 other commas promoted to semicolons or split into sentences
For me, 3 and 4 already ring alarm bells: they're changes to my writing style. Apparently, I don't use semicolons when I should, but, clearly that's part of my current writing style.
This didn't go unnoticed in Claude's assessment either:
Bucket 4 is the one to push back on. Your register is conversational and comma-driven — clauses stacked with commas the way you'd say them out loud. ; however, is a formal-report construction, and it reads noticeably stiffer:
the whole point in it is load shifting, without a battery peak times would be ruinous → …load shifting; without a battery, peak times would be ruinous
1Password's web-ui doesn't provide a way to export passwords, however the desktop client can do it → …passwords; however, the desktop client…
It's worst in the vehicle docs, which are written in a distinctly over-the-shoulder, talking-you-through-it voice. Turn the ignition on and test, you should now find the fan kicks in → test; you should now find…, and The blower rotates slightly to unclip from it's mountings, if you're facing the glovebox… → its mountings; if you're facing…. The apostrophe fix is right; the semicolon makes a bloke-under-a-dashboard instruction sound like a service manual.
Two smaller losses: the editor deleted deliberate rhythm commas in The hope, is that crawlers fall into the maze early and Now all that I have to do, is find time to read them — both are grammatically correct removals and both flatten a beat you clearly put there on purpose.
Note this is a register shift, not an Americanism — ; however, is correct BrE. But the association is with American style-guide comma-splice horror, which may be why it feels transatlantic even though the spelling is clean.
So, even the "voice preserving" grammar and spelling check option resulted in noticeable changes to written voice, turning an informal conversational style into something more formal and dry.
As if that wasn't bad enough, Claude's assessment also pointed out that there were inconsistencies in the edit:
Gitops→GitOps and forgejo→Forgejo fixed, but Gitlab left uncorrected throughout (should be GitLab), and javascript left lowercase everywhere.
Lowercase-after-colon applied in one spot (to a pistol: Not only→not only) but not another (I really want Firefox to succeed: As small as its market share…).
Hyphenation added to day-to-day, built-in, plug-in, followers-only — but well intentioned, ill-thought out, third party plugins, work related books, double sided tape, higher end markers, lower hanging fruit, dome headed nuts and over cooked all left bare
The colon sweep missed at least one: "…converted to PCP with the following configuration" still has no colon before its bullet list.
So, my vaunted AI reviewer had let things slip through!
Not wanting to leave everything to Claude, I also eyeballed the diff to see whether I could pick up on any undesirable changes that it hadn't flagged.
Changing the alt-text of images. In this case, the changes looked correct, but alt-text changes could easily go unnoticed when reading a page back post-edit
Making changes to HTML markup (removing a couple of tags)
Changing a heading title (and therefore the named anchor that it generates)
These are all functional changes which went undisclosed by Claude in both the original editing run and the assessment. Setting them live on existing content, at best, breaks some links or, at worst, breaks the very pages that Claude was trying to fix.
I told Claude that I'd changed my mind and it should perform deeper copy-editing:
I changed my mind, can we do this as a full copy-edit (option 3 in your earlier question)? I've reset the repo to its prior state
Again, Claude edited all 25 files.
As part of its response, Claude noted that some changes had been derived from my own style:
So, this was another nod towards preserving my voice, even within an editing approach which didn't promise not to affect it.
Once again, I dumped out a diff:
git diff > ../changes2.diff
And then told Claude to analyse it:
OK, there's now a changes2.diff from a deeper edit run (starting from the same base state). Could you analyse that please?
Claude, once again, said that the posts hadn't been Americanised:
However, it then proceeded to take issue with some of the changes that had been made:
The regression: a wrong title-case rule
17 headings now capitalise articles, conjunctions and short prepositions:
## Overview Of Changes ### Weakening The Valve Spring
## Proportion Of Consumption ### Detections And False Positives
## Building A Hugo Site ### Publishing And Flushing CDN Caches
## Connecting To Booklore ### Looking To 2026
## DNS And Reverse Proxy #### Attempting To Identify Sites … and 7 more
No major style guide does this — Oxford, Chicago, AP and the Guardian all keep of, and, to, the, a, in lowercase inside a title. This isn't a BrE/AmE question; it's just wrong, and it's applied to 17 headings while leaving your other headings alone, so it makes the site less consistent than before. This one I'd reject wholesale.
So... the deeper edit had potentially made things worse .
Unsurprisingly, Claude also found voice changes:
Every one is tighter. Collectively they remove the hedging, self-qualifying, thinking-aloud quality that is the voice. I think, really, quite, try and aren't padding in your prose — they're the register.
The edit also merged sentences that I'd split:
It didn't really make sense for me to try and breathe new life into it by creating a Forgejo variant (FILS?). So, I decided, instead, to start from scratch. → It didn't really make sense to breathe new life into them by creating a Forgejo variant (FILS?), so I decided to start from scratch instead.
This time, Claude's analysis did call out an alt-text change:
Alt text invented for a reused image. In the Saab post, the same battery-tray-torx.jpg appears twice — once for removal, once for refitting. The editor changed the second instance to Photo of a screw being fitted with a torx set. The photo still shows removal. That's now inaccurate for a screen-reader user.
Many of the inconsistencies that Claude highlighted in the original edit were also present in this later one.
Claude's assessment of this round was fairly uncompromising:
Pass 2 is the better proofreader and the worse editor . It finds more genuine errors and finishes the hyphenation sweep pass 1 abandoned halfway, but it also imposes an incorrect title-case rule across 17 headings and performs an unrequested line edit that measurably flattens your register
At the beginning of this post, I posited a theory: that voice changes could easily go unnoticed when reviewing a single post in isolation.
Whether or not that's true is important: most "proof-read" AI-using bloggers are, presumably, only using it on individual posts just before they're published.
I told Claude to look at the changes made to my most recent post and to assess whether we'd have been likely to notice voice changes:
Have a look at the changes for my most recent post (Migrating From 1Password to Vaultwarden). Would we have been able to see the change in voice if we'd only looked at that? I'd like an answer for each of the passes please
Claude said that we'd have missed them:
Pass 1 : no, you'd have missed it
Four changed lines, and three are unambiguous typos:
check for if :latest has changed → check if
So, I also a created a cronjob → So, I also created a cronjob
like to setup automatic replication → like to set up
The fourth is the register change:
...export passwords, however the desktop client [can do it]
...export passwords; however, the desktop client [can do it]
But with n=1 you have no way to read that as a house rule. That sentence genuinely is a comma splice, so the fix is defensible on its own terms — any reasonable reviewer ticks it off and moves on. The ; however, sweep only becomes visible as a pattern at seven instances across the corpus, where you can see it firing on sentences that didn't need it and stiffening the register uniformly.
So from this file alone you'd conclude pass 1 was a clean, light, entirely benign typo pass. That happens to be roughly the right verdict — but you'd have reached it by luck, not evidence. You'd also have underrated it: this post contains no it's/its errors, no proper nouns and no compound modifiers, so none of pass 1's actual catches (24 possessives, Geissele, MacAulay, the OPDS anchor fix) are represented here at all.
Claude opined that the deeper pass was more noticeable but, being honest, I'm not sure that it's describing something that I would have picked up on:
Pass 2 : yes — from one sentence, in fact
Twelve changed lines. The tell is this one:
But, I quickly ran into an issue: the master password has an enforced minimum length of 12 chars:
But I quickly ran into an issue - the master password has an enforced minimum length of 12 chars:
Two changes in nine words — the But, comma strip

[truncated]
