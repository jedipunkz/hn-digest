---
source: "https://github.com/taylorancapital/nothing-threw"
hn_url: "https://news.ycombinator.com/item?id=49936772"
title: "AI agents failed against a live business, and how each was caught"
article_title: "GitHub - taylorancapital/nothing-threw: Seventy-four ways autonomous agents failed against a live business over five months - and how each one was actually caught. Nine were caught by anything automated. · GitHub"
image: "https://opengraph.githubassets.com/73a1b9aa8d51a81449f21326b77a79fc3a03c67b797f03eccd0c3f663cadf7f6/taylorancapital/nothing-threw"
author: "taylorancapital"
captured_at: "2026-10-02T19:01:06Z"
capture_tool: "hn-digest"
hn_id: 49936772
score: 1
comments: 0
posted_at: "2026-10-02T18:19:21Z"
tags:
  - hacker-news
---

# AI agents failed against a live business, and how each was caught

- HN: [49936772](https://news.ycombinator.com/item?id=49936772)
- Source: [github.com](https://github.com/taylorancapital/nothing-threw)
- Score: 1
- Comments: 0
- Posted: 2026-10-02T18:19:21Z

## Translation

Title: AI agents failed against a live business, and how each was caught
Article title: GitHub - taylorancapital/nothing-threw: Seventy-four ways autonomous agents failed against a live business over five months - and how each one was actually caught. Nine were caught by anything automated. · GitHub
Description: Seventy-four ways autonomous agents failed against a live business over five months - and how each one was actually caught. Nine were caught by anything automated. - taylorancapital/nothing-threw

Article text:
GitHub - taylorancapital/nothing-threw: Seventy-four ways autonomous agents failed against a live business over five months - and how each one was actually caught. Nine were caught by anything automated. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
taylorancapital
/
nothing-threw
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
FIELD_NOTES_2026-05.md FIELD_NOTES_2026-05.md FIELD_NOTES_2026-06.md FIELD_NOTES_2026-06.md FIELD_NOTES_2026-07.md FIELD_NOTES_2026-07.md FIELD_NOTES_2026-08.md FIELD_NOTES_2026-08.md FIELD_NOTES_2026-09.md FIELD_NOTES_2026-09.md FIELD_NOTES_2026-10.md FIELD_NOTES_2026-10.md README.md README.md SIX_OF_FORTY_EIGHT.md SIX_OF_FORTY_EIGHT.md TAXONOMY.md TAXONOMY.md View all files Repository files navigation
Seventy-four ways autonomous agents failed while running a live business — and how each one was actually caught.
Field notes, May–October 2026. One events business: a live Meta ad account, Stripe checkout, a Firestore back end, and a nightly analytics agent that writes reports and opens pull requests unattended.
Forty-two of the first forty-eight reported success. The interesting question turned out not to be how they failed, but how anyone ever found out — and, five months in, whether that's getting any easier. It isn't yet.
Start here → Nine of Seventy-Four — the argument in one essay. (The filename predates the count; it is kept so existing links work.)
Run it on your own system → The method — detection modes, failure classes, the bar for inclusion, and a monthly pass you can copy.
How each of the 74 was detected
Detection
Means
Count
GATE
A deterministic check refused it
2
THREW
An actual error surfaced at the time
7
HUMAN
Someone distrusted a number or a claim
19
OPERATOR
Reported as "data went missing"
1
LATER
Found by an unrelated dig, days to months on
45
Nine of seventy-four were caught by anything automated — 12%. The rest were caught because a person looked at a number and thought that can't be right — or because weeks later something else went wrong and led back to it.
Across the five passes this catalogue has been through, that share has gone 11%, 9%, 15%, 13%, 12% — a narrow band with no trend. The one apparent rise was an artifact: backfilling earlier months added two same-day code bugs, exactly the loud, self-announcing kind, and lifted the numerator in a single stroke. Nothing had gotten better. As the incidents get subtler — wrong conclusions from correct data, not broken code — nothing built so far catches them.
Month
Incidents
Detail
May 2026
1
FIELD_NOTES_2026-05.md
June 2026
3
FIELD_NOTES_2026-06.md
July 2026
2
FIELD_NOTES_2026-07.md
August 2026
11
FIELD_NOTES_2026-08.md
September 2026
31
FIELD_NOTES_2026-09.md
12 September – 2 October 2026
26
FIELD_NOTES_2026-10.md
The write-ups under "The incidents" below cover incidents 01–48. Incidents 49–74 are written up in FIELD_NOTES_2026-10.md and are not repeated here.
The month files carry the full anchors — the commit, log line, pull request or API read behind each incident. Incident numbers are assigned in catalogue order, not calendar order, and are never reassigned once published, so a May incident can carry a higher number than an August one.
A note on the anchors. They reference commits in a private repository, so the click-through is unavailable. Each one carries its date, commit subject and pull request number, which is enough to identify a specific change — but you are taking the contents on trust, and you should weigh the whole catalogue accordingly. The repository was public until September 2026 and was made private after a history rewrite removed customer personal data from it; see incident 40 for what that rewrite got wrong.
A. The agent was confidently wrong
01 — Two ID spaces that never match · HUMAN
An analysis compared an ad creative's video_id against the object_id s in a video-engagement audience. The first is the source upload; the second is the delivered rendition — the platform's auto-generated crops. They can never match, even for the same video.
Cost: a false headline that the retargeting audience contained none of the running videos, an entire "the funnel was never wired up" conclusion built on top of it, and a pull request. The audience was correctly scoped the whole time.
The insights API returns different conversion counts for the same window with and without a gender breakdown. A report divided a breakdown numerator by an un-broken denominator.
Cost: published "women's share of landing-page views: 105.00%" and nobody noticed. The impossible figure was the visible tip; every other gender-split cost in that report sat on the same mismatch and looked fine.
An event post-mortem reasoned from the code alone and blamed a two-for-one ticket offer for a gender imbalance. The flag it depended on was false on every record involved; the offer had played no part at all.
Cost: a confident, circulated, wrong explanation. Rewritten only after someone queried the data directly — and the superseded reasoning was deleted rather than kept, because a wrong cause left lying around is worse than none.
An investigation into an ad account opened by listing only ACTIVE campaigns. Three dormant campaigns — the ones that had delivered the cheapest clicks the account ever bought — were invisible to every subsequent step.
Cost: three successive proposals, each blocked by a different constraint, each looking like progress. The real fix was flipping a status back. The cost was not the wrong answer; it was four review cycles of someone else's time.
A commit removing gendered ticket pricing updated the admin event-creation form and the on-page price display to a single spots/price model. It did not touch the payment path, which still read the old per-gender fields — undefined on any new-model event. Every purchase attempt on a new event was rejected as sold out; had one somehow passed, price would have resolved to $0, charging only the flat service fee.
Cost: none realised. The live payment key wasn't switched in until the day after the fix — no real card could have reached this path. Fixed with a single source of truth for both the admin and payment sides.
A guest paying with a 3-D Secure card got a confirmed ticket but no welcome email, no trial signup, no nurture lead — the entire post-purchase funnel silently skipped anyone who hit that challenge, because the enrollment calls ran before the challenge could return. Separately, a duplicate submit could bump the seat counter twice before the retry was recognised as a repeat.
Cost: the 3DS gap was live with real payments for five days — an unknown number of real guests never got their onboarding funnel. The duplicate-submit gap's worst case, per the fix: "an over-counted seat, never an oversell or double charge."
A conversion-tracking install shipped with a literal placeholder string instead of a real tracking ID — not a valid tag, and it broke the site's analytics initialization on every page that carried it.
Cost: the only telemetry the business had, eight days after go-live, was dark sitewide for roughly four hours. Fixed the same day by removing the tag rather than supplying a real ID.
Two strings used a quote style that broke on an apostrophe inside them — a fatal syntax error that broke the entire script on two new city pages. No events loaded; the pages fell back to static, hardcoded copy for the wrong city. A second, independent defect in the same file tagged fallback content with the wrong city in the page's own structured data, handed directly to search engines.
Cost: near zero — fixed the same day it shipped, consistent with a syntax error that fails immediately and visibly on load.
A nightly analysis pipeline's rotation logic skips re-analyzing an export it believes is unchanged from the night before — decided, for three consecutive cycles, by matching filenames and file modification times. Both are always identical across pulls by construction; the underlying data had in fact been deleted and freshly re-pulled. Three consecutive nightly cycles treated genuinely new data as a stale duplicate and skipped analyzing it.
Cost: at least two full nightly cycles where fresh exported data was never actually analyzed. Caught the same night by reading each file's own embedded date range instead of trusting its name.
An audit of the analytics property's coverage concluded that 245 fields were "structurally empty because every ad we run is on Meta," written from category names rather than from probing the fields. Wrong three separate ways: one field family turned out to be real and populated; two others returned constant placeholder values on every row, which looked like data but weren't; and most of the true remainder returned a different, unrelated "not set" value with a different fix.
Cost: weeks of live ad spend on another platform stayed unread because a neighboring probe (incident 20) concluded it was unmeasurable. The durable cost is the method: triaging an API's fields by category name produces confident, checkable, wrong claims.
A metrics family errored when queried alone, asking for a companion dimension to be added. Read in isolation, that error means "unavailable." It means "needs a pairing."
Cost: weeks of real ad spend invisible to every report written before the fix — some of which bought real clicks and zero attributed sessions, and landed in neither of the business's two cost-tracking paths. Every efficiency figure in every prior report was understated by that amount.
A live read of every campaign found one objective holding all of one week's purchases on similar spend to another objective holding zero — nine times the visitors, no sales — and it was presented as proof that the second objective does not sell.
Tested per dollar, it was not statistically significant. With six lifetime conversions on the account, no causal claim in either direction was supportable.
Cost: none realised. Someone asked whether it was actually causal; the finding was corrected the same day.
A delivery diagnosis read the ad platform's attributed conversions as sales and concluded the business's sales had stopped for over a week. They had not — in the exact window the report called dead, more than a dozen tickets actually sold. The platform sees roughly one in nine real ticket sales; a complete revenue join already existed in the business's own database and nobody had read it. Read correctly, gross return on ad spend for the period was better than break-even, not the fraction of a dollar the retracted report implied.
Cost: a wrong "the business is dying" headline reach

[truncated]

## Original Extract

Seventy-four ways autonomous agents failed against a live business over five months - and how each one was actually caught. Nine were caught by anything automated. - taylorancapital/nothing-threw

GitHub - taylorancapital/nothing-threw: Seventy-four ways autonomous agents failed against a live business over five months - and how each one was actually caught. Nine were caught by anything automated. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
taylorancapital
/
nothing-threw
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
FIELD_NOTES_2026-05.md FIELD_NOTES_2026-05.md FIELD_NOTES_2026-06.md FIELD_NOTES_2026-06.md FIELD_NOTES_2026-07.md FIELD_NOTES_2026-07.md FIELD_NOTES_2026-08.md FIELD_NOTES_2026-08.md FIELD_NOTES_2026-09.md FIELD_NOTES_2026-09.md FIELD_NOTES_2026-10.md FIELD_NOTES_2026-10.md README.md README.md SIX_OF_FORTY_EIGHT.md SIX_OF_FORTY_EIGHT.md TAXONOMY.md TAXONOMY.md View all files Repository files navigation
Seventy-four ways autonomous agents failed while running a live business — and how each one was actually caught.
Field notes, May–October 2026. One events business: a live Meta ad account, Stripe checkout, a Firestore back end, and a nightly analytics agent that writes reports and opens pull requests unattended.
Forty-two of the first forty-eight reported success. The interesting question turned out not to be how they failed, but how anyone ever found out — and, five months in, whether that's getting any easier. It isn't yet.
Start here → Nine of Seventy-Four — the argument in one essay. (The filename predates the count; it is kept so existing links work.)
Run it on your own system → The method — detection modes, failure classes, the bar for inclusion, and a monthly pass you can copy.
How each of the 74 was detected
Detection
Means
Count
GATE
A deterministic check refused it
2
THREW
An actual error surfaced at the time
7
HUMAN
Someone distrusted a number or a claim
19
OPERATOR
Reported as "data went missing"
1
LATER
Found by an unrelated dig, days to months on
45
Nine of seventy-four were caught by anything automated — 12%. The rest were caught because a person looked at a number and thought that can't be right — or because weeks later something else went wrong and led back to it.
Across the five passes this catalogue has been through, that share has gone 11%, 9%, 15%, 13%, 12% — a narrow band with no trend. The one apparent rise was an artifact: backfilling earlier months added two same-day code bugs, exactly the loud, self-announcing kind, and lifted the numerator in a single stroke. Nothing had gotten better. As the incidents get subtler — wrong conclusions from correct data, not broken code — nothing built so far catches them.
Month
Incidents
Detail
May 2026
1
FIELD_NOTES_2026-05.md
June 2026
3
FIELD_NOTES_2026-06.md
July 2026
2
FIELD_NOTES_2026-07.md
August 2026
11
FIELD_NOTES_2026-08.md
September 2026
31
FIELD_NOTES_2026-09.md
12 September – 2 October 2026
26
FIELD_NOTES_2026-10.md
The write-ups under "The incidents" below cover incidents 01–48. Incidents 49–74 are written up in FIELD_NOTES_2026-10.md and are not repeated here.
The month files carry the full anchors — the commit, log line, pull request or API read behind each incident. Incident numbers are assigned in catalogue order, not calendar order, and are never reassigned once published, so a May incident can carry a higher number than an August one.
A note on the anchors. They reference commits in a private repository, so the click-through is unavailable. Each one carries its date, commit subject and pull request number, which is enough to identify a specific change — but you are taking the contents on trust, and you should weigh the whole catalogue accordingly. The repository was public until September 2026 and was made private after a history rewrite removed customer personal data from it; see incident 40 for what that rewrite got wrong.
A. The agent was confidently wrong
01 — Two ID spaces that never match · HUMAN
An analysis compared an ad creative's video_id against the object_id s in a video-engagement audience. The first is the source upload; the second is the delivered rendition — the platform's auto-generated crops. They can never match, even for the same video.
Cost: a false headline that the retargeting audience contained none of the running videos, an entire "the funnel was never wired up" conclusion built on top of it, and a pull request. The audience was correctly scoped the whole time.
The insights API returns different conversion counts for the same window with and without a gender breakdown. A report divided a breakdown numerator by an un-broken denominator.
Cost: published "women's share of landing-page views: 105.00%" and nobody noticed. The impossible figure was the visible tip; every other gender-split cost in that report sat on the same mismatch and looked fine.
An event post-mortem reasoned from the code alone and blamed a two-for-one ticket offer for a gender imbalance. The flag it depended on was false on every record involved; the offer had played no part at all.
Cost: a confident, circulated, wrong explanation. Rewritten only after someone queried the data directly — and the superseded reasoning was deleted rather than kept, because a wrong cause left lying around is worse than none.
An investigation into an ad account opened by listing only ACTIVE campaigns. Three dormant campaigns — the ones that had delivered the cheapest clicks the account ever bought — were invisible to every subsequent step.
Cost: three successive proposals, each blocked by a different constraint, each looking like progress. The real fix was flipping a status back. The cost was not the wrong answer; it was four review cycles of someone else's time.
A commit removing gendered ticket pricing updated the admin event-creation form and the on-page price display to a single spots/price model. It did not touch the payment path, which still read the old per-gender fields — undefined on any new-model event. Every purchase attempt on a new event was rejected as sold out; had one somehow passed, price would have resolved to $0, charging only the flat service fee.
Cost: none realised. The live payment key wasn't switched in until the day after the fix — no real card could have reached this path. Fixed with a single source of truth for both the admin and payment sides.
A guest paying with a 3-D Secure card got a confirmed ticket but no welcome email, no trial signup, no nurture lead — the entire post-purchase funnel silently skipped anyone who hit that challenge, because the enrollment calls ran before the challenge could return. Separately, a duplicate submit could bump the seat counter twice before the retry was recognised as a repeat.
Cost: the 3DS gap was live with real payments for five days — an unknown number of real guests never got their onboarding funnel. The duplicate-submit gap's worst case, per the fix: "an over-counted seat, never an oversell or double charge."
A conversion-tracking install shipped with a literal placeholder string instead of a real tracking ID — not a valid tag, and it broke the site's analytics initialization on every page that carried it.
Cost: the only telemetry the business had, eight days after go-live, was dark sitewide for roughly four hours. Fixed the same day by removing the tag rather than supplying a real ID.
Two strings used a quote style that broke on an apostrophe inside them — a fatal syntax error that broke the entire script on two new city pages. No events loaded; the pages fell back to static, hardcoded copy for the wrong city. A second, independent defect in the same file tagged fallback content with the wrong city in the page's own structured data, handed directly to search engines.
Cost: near zero — fixed the same day it shipped, consistent with a syntax error that fails immediately and visibly on load.
A nightly analysis pipeline's rotation logic skips re-analyzing an export it believes is unchanged from the night before — decided, for three consecutive cycles, by matching filenames and file modification times. Both are always identical across pulls by construction; the underlying data had in fact been deleted and freshly re-pulled. Three consecutive nightly cycles treated genuinely new data as a stale duplicate and skipped analyzing it.
Cost: at least two full nightly cycles where fresh exported data was never actually analyzed. Caught the same night by reading each file's own embedded date range instead of trusting its name.
An audit of the analytics property's coverage concluded that 245 fields were "structurally empty because every ad we run is on Meta," written from category names rather than from probing the fields. Wrong three separate ways: one field family turned out to be real and populated; two others returned constant placeholder values on every row, which looked like data but weren't; and most of the true remainder returned a different, unrelated "not set" value with a different fix.
Cost: weeks of live ad spend on another platform stayed unread because a neighboring probe (incident 20) concluded it was unmeasurable. The durable cost is the method: triaging an API's fields by category name produces confident, checkable, wrong claims.
A metrics family errored when queried alone, asking for a companion dimension to be added. Read in isolation, that error means "unavailable." It means "needs a pairing."
Cost: weeks of real ad spend invisible to every report written before the fix — some of which bought real clicks and zero attributed sessions, and landed in neither of the business's two cost-tracking paths. Every efficiency figure in every prior report was understated by that amount.
A live read of every campaign found one objective holding all of one week's purchases on similar spend to another objective holding zero — nine times the visitors, no sales — and it was presented as proof that the second objective does not sell.
Tested per dollar, it was not statistically significant. With six lifetime conversions on the account, no causal claim in either direction was supportable.
Cost: none realised. Someone asked whether it was actually causal; the finding was corrected the same day.
A delivery diagnosis read the ad platform's attributed conversions as sales and concluded the business's sales had stopped for over a week. They had not — in the exact window the report called dead, more than a dozen tickets actually sold. The platform sees roughly one in nine real ticket sales; a complete revenue join already existed in the business's own database and nobody had read it. Read correctly, gross return on ad spend for the period was better than break-even, not the fraction of a dollar the retracted report implied.
Cost: a wrong "the business is dying" headline reach

[truncated]
