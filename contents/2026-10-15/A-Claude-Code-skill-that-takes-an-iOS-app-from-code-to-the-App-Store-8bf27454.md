---
source: "https://github.com/allexp1/app-store-kit"
hn_url: "https://news.ycombinator.com/item?id=50106794"
title: "A Claude Code skill that takes an iOS app from code to the App Store"
article_title: "GitHub - allexp1/app-store-kit: Claude Code skill: take an iOS/iPadOS app from code to the App Store and keep it updated. Code pre-check, every language, screenshots, build upload, App Store Connect through the official API. · GitHub"
image: "https://opengraph.githubassets.com/d72964d21715a693849697619536630e8014689cdbe0c59e6c3029ef0bafba4d/allexp1/app-store-kit"
author: "allexp"
captured_at: "2026-10-15T18:02:57Z"
capture_tool: "hn-digest"
hn_id: 50106794
score: 1
comments: 0
posted_at: "2026-10-15T17:08:39Z"
tags:
  - hacker-news
---

# A Claude Code skill that takes an iOS app from code to the App Store

- HN: [50106794](https://news.ycombinator.com/item?id=50106794)
- Source: [github.com](https://github.com/allexp1/app-store-kit)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T17:08:39Z

## Translation

Title: A Claude Code skill that takes an iOS app from code to the App Store
Article title: GitHub - allexp1/app-store-kit: Claude Code skill: take an iOS/iPadOS app from code to the App Store and keep it updated. Code pre-check, every language, screenshots, build upload, App Store Connect through the official API. · GitHub
Description: Claude Code skill: take an iOS/iPadOS app from code to the App Store and keep it updated. Code pre-check, every language, screenshots, build upload, App Store Connect through the official API. - allexp1/app-store-kit

Article text:
GitHub - allexp1/app-store-kit: Claude Code skill: take an iOS/iPadOS app from code to the App Store and keep it updated. Code pre-check, every language, screenshots, build upload, App Store Connect through the official API. · GitHub
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
allexp1
app-store-kit Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
Claude Code skill: take an iOS/iPadOS app from code to the App Store and keep it updated. Code pre-check, every language, screenshots, build upload, App Store Connect through the official API.
Readme MIT license Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
.github/ workflows .github/ workflows .living .living bin bin examples/ sample-campaign examples/ sample-campaign references references scripts scripts templates templates tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md install.sh install.sh package.json package.json publish.sh publish.sh requirements.txt requirements.txt test.sh test.sh View all files Repository files navigation
A skill for Claude Code that takes an iPhone/iPad app from code to the App Store, and keeps it updated.
You say one sentence. It does the rest, step by step, and stops only where Apple needs you.
Beyond caption templates, the kit art-directs a screenshot set from the app's real screens: continuous panoramas
across panels, 3D phones crossing the seams, editorial type, a dark spread, UI cards lifted out of the phone.
Make six premium App Store screenshots for this app.
Make a panoramic six-screen campaign: floating iPhones, editorial type, cinematic light.
Redesign my existing App Store screenshots. Keep the real UI and branding. Show me six different concepts first.
Make English, Hebrew and Russian versions of the chosen campaign.
Use these reference images for the quality bar and art direction; don't copy them.
How it works ( scripts/campaign.py , details in references/creative-director.md ):
Inputs. Required: the app's screenshots. Optional: the icon (palette), brand colours and font files,
reference images (read for technique, never copied), background art with its source and licence.
Concepts. campaign.py concepts renders six different directions as one contact sheet (each a full low-res
panorama with its rationale, type and palette). You pick one, combine two, or ask for others.
Copy. One benefit per panel. Numbers, rankings, ratings and awards need a source, or the check fails;
prices, URLs and Apple badges are refused.
Render. One master image across all panels (headless Chromium), cut into exact App Store sizes, RGB, with a
QA report (text across seams, safe areas, legibility at search-result size, contrast, missing glyphs).
The same file always gives the same bytes.
Localize. Add a locale and its copy; Hebrew and Arabic mirror the whole scene and start from the right.
Export. render --install KIT puts the panels where the rest of the kit uploads them (after your yes).
campaign.py canvas also makes an editable Claude Design canvas.
Replace screenshots: overwrite the files under screens: and render again; layout and copy stay. Add a composition
engine: one decorated function in scripts/studio/engines.py . A full example from a made-up hiking app ("Ridgeline",
synthetic screens) is in examples/sample-campaign/ ; its rendered contact sheet and en/he/ru sets are in
examples/sample-campaign/preview/ on GitHub. To render it yourself:
python3 examples/sample-campaign/make_screens.py # makes the synthetic screens (not shipped on npm)
python3 scripts/campaign.py concepts examples/sample-campaign/campaign.yaml
python3 scripts/campaign.py render examples/sample-campaign/campaign.yaml
Limits: phones are CSS 3D, not ray-traced; no photographed hands unless you supply the art; fonts download once
from Google Fonts (open licences) or come from your files.
What it does by itself, and where it stops
The App Store Connect API does almost everything: texts in every language, screenshots and previews, age
rating, accessibility, review details, build, submit, release, phased release.
Chrome does the few screens the API can't reach, such as App Privacy and creating the app record.
You do only what Apple reserves for the account holder: signing in, agreements, tax, banking, EU trader status.
The skill opens the exact page, outlines the button and waits for you to say "done". It never presses Agree.
Nothing reaches Apple without your yes. A plan is used once and only for the files and account state it saw. Before changing anything it shows one page of planned changes
( connect-plan.md ). Upload, submit and release each ask again.
One page shows where you are: STATUS.md , with ✅ done, ✋ your turn, 🛑 needs your yes, ⬜ to do.
It is kept current: the references cite Apple's pages with dates, and every change Apple makes to App Store Connect
that the skill finds on a live run is written back (see .living/CHANGELOG.md ).
npx app-store-kit
or npx skills add allexp1/app-store-kit , or from a clone: ./install.sh .
Then restart Claude Code. The scripts need Python 3.10+ (and, for rendering, a Chromium: npx playwright install chromium-headless-shell once):
python3 -m pip install -r ~ /.claude/skills/app-store-kit/requirements.txt
The App Store Connect key (once)
The skill asks before it uses a key, once per session. To make one: App Store Connect › Users and Access ›
Integrations › App Store Connect API › Team Keys › Generate (role App Manager ), then download the .p8 file to
~/.appstoreconnect/private_keys/ . The skill only stores the Key ID and Issuer ID in your kit. It never opens,
prints or commits the key file.
SKILL.md : the instructions Claude follows.
scripts/ :
precheck.py : code against Apple's rules.
asc_api.py and asc_sync.py : the App Store Connect API ( status , plan , apply , submit , release ).
status_board.py : writes STATUS.md .
locales.py , whats_new.py and release.py : languages, update notes, version bump, archive and upload.
render.py and capture.sh : screenshots.
campaign.py (+ studio/ ): designed screenshot campaigns: concepts, panoramas, 3D devices, exact exports.
lint_metadata.py , validate_kit.py , preview_check.py and connect_guide.py : checks and a hand-typing guide.
growth.py and asc_reports.py : growth and marketing, and measuring it.
rejection.py and reviews.py : App Review rejections; ratings and approval-gated review replies.
subscriptions.py , testflight.py and pricing.py : subscriptions, TestFlight betas and feedback, prices and
availability.
references/ : Apple's specs, review guidelines, App Store Connect fields, the API, privacy labels and copywriting.
tests/ : ./test.sh runs everything offline, without keys.
Claude Code skill: take an iOS/iPadOS app from code to the App Store and keep it updated. Code pre-check, every language, screenshots, build upload, App Store Connect through the official API.
Readme MIT license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Claude Code skill: take an iOS/iPadOS app from code to the App Store and keep it updated. Code pre-check, every language, screenshots, build upload, App Store Connect through the official API. - allexp1/app-store-kit

GitHub - allexp1/app-store-kit: Claude Code skill: take an iOS/iPadOS app from code to the App Store and keep it updated. Code pre-check, every language, screenshots, build upload, App Store Connect through the official API. · GitHub
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
allexp1
app-store-kit Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
Claude Code skill: take an iOS/iPadOS app from code to the App Store and keep it updated. Code pre-check, every language, screenshots, build upload, App Store Connect through the official API.
Readme MIT license Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
12 Commits 12 Commits Folders and files
.github/ workflows .github/ workflows .living .living bin bin examples/ sample-campaign examples/ sample-campaign references references scripts scripts templates templates tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md install.sh install.sh package.json package.json publish.sh publish.sh requirements.txt requirements.txt test.sh test.sh View all files Repository files navigation
A skill for Claude Code that takes an iPhone/iPad app from code to the App Store, and keeps it updated.
You say one sentence. It does the rest, step by step, and stops only where Apple needs you.
Beyond caption templates, the kit art-directs a screenshot set from the app's real screens: continuous panoramas
across panels, 3D phones crossing the seams, editorial type, a dark spread, UI cards lifted out of the phone.
Make six premium App Store screenshots for this app.
Make a panoramic six-screen campaign: floating iPhones, editorial type, cinematic light.
Redesign my existing App Store screenshots. Keep the real UI and branding. Show me six different concepts first.
Make English, Hebrew and Russian versions of the chosen campaign.
Use these reference images for the quality bar and art direction; don't copy them.
How it works ( scripts/campaign.py , details in references/creative-director.md ):
Inputs. Required: the app's screenshots. Optional: the icon (palette), brand colours and font files,
reference images (read for technique, never copied), background art with its source and licence.
Concepts. campaign.py concepts renders six different directions as one contact sheet (each a full low-res
panorama with its rationale, type and palette). You pick one, combine two, or ask for others.
Copy. One benefit per panel. Numbers, rankings, ratings and awards need a source, or the check fails;
prices, URLs and Apple badges are refused.
Render. One master image across all panels (headless Chromium), cut into exact App Store sizes, RGB, with a
QA report (text across seams, safe areas, legibility at search-result size, contrast, missing glyphs).
The same file always gives the same bytes.
Localize. Add a locale and its copy; Hebrew and Arabic mirror the whole scene and start from the right.
Export. render --install KIT puts the panels where the rest of the kit uploads them (after your yes).
campaign.py canvas also makes an editable Claude Design canvas.
Replace screenshots: overwrite the files under screens: and render again; layout and copy stay. Add a composition
engine: one decorated function in scripts/studio/engines.py . A full example from a made-up hiking app ("Ridgeline",
synthetic screens) is in examples/sample-campaign/ ; its rendered contact sheet and en/he/ru sets are in
examples/sample-campaign/preview/ on GitHub. To render it yourself:
python3 examples/sample-campaign/make_screens.py # makes the synthetic screens (not shipped on npm)
python3 scripts/campaign.py concepts examples/sample-campaign/campaign.yaml
python3 scripts/campaign.py render examples/sample-campaign/campaign.yaml
Limits: phones are CSS 3D, not ray-traced; no photographed hands unless you supply the art; fonts download once
from Google Fonts (open licences) or come from your files.
What it does by itself, and where it stops
The App Store Connect API does almost everything: texts in every language, screenshots and previews, age
rating, accessibility, review details, build, submit, release, phased release.
Chrome does the few screens the API can't reach, such as App Privacy and creating the app record.
You do only what Apple reserves for the account holder: signing in, agreements, tax, banking, EU trader status.
The skill opens the exact page, outlines the button and waits for you to say "done". It never presses Agree.
Nothing reaches Apple without your yes. A plan is used once and only for the files and account state it saw. Before changing anything it shows one page of planned changes
( connect-plan.md ). Upload, submit and release each ask again.
One page shows where you are: STATUS.md , with ✅ done, ✋ your turn, 🛑 needs your yes, ⬜ to do.
It is kept current: the references cite Apple's pages with dates, and every change Apple makes to App Store Connect
that the skill finds on a live run is written back (see .living/CHANGELOG.md ).
npx app-store-kit
or npx skills add allexp1/app-store-kit , or from a clone: ./install.sh .
Then restart Claude Code. The scripts need Python 3.10+ (and, for rendering, a Chromium: npx playwright install chromium-headless-shell once):
python3 -m pip install -r ~ /.claude/skills/app-store-kit/requirements.txt
The App Store Connect key (once)
The skill asks before it uses a key, once per session. To make one: App Store Connect › Users and Access ›
Integrations › App Store Connect API › Team Keys › Generate (role App Manager ), then download the .p8 file to
~/.appstoreconnect/private_keys/ . The skill only stores the Key ID and Issuer ID in your kit. It never opens,
prints or commits the key file.
SKILL.md : the instructions Claude follows.
scripts/ :
precheck.py : code against Apple's rules.
asc_api.py and asc_sync.py : the App Store Connect API ( status , plan , apply , submit , release ).
status_board.py : writes STATUS.md .
locales.py , whats_new.py and release.py : languages, update notes, version bump, archive and upload.
render.py and capture.sh : screenshots.
campaign.py (+ studio/ ): designed screenshot campaigns: concepts, panoramas, 3D devices, exact exports.
lint_metadata.py , validate_kit.py , preview_check.py and connect_guide.py : checks and a hand-typing guide.
growth.py and asc_reports.py : growth and marketing, and measuring it.
rejection.py and reviews.py : App Review rejections; ratings and approval-gated review replies.
subscriptions.py , testflight.py and pricing.py : subscriptions, TestFlight betas and feedback, prices and
availability.
references/ : Apple's specs, review guidelines, App Store Connect fields, the API, privacy labels and copywriting.
tests/ : ./test.sh runs everything offline, without keys.
Claude Code skill: take an iOS/iPadOS app from code to the App Store and keep it updated. Code pre-check, every language, screenshots, build upload, App Store Connect through the official API.
Readme MIT license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
