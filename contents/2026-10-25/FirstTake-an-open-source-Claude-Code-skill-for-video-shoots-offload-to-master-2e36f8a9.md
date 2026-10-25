---
source: "https://github.com/emersonsklar/firsttake"
hn_url: "https://news.ycombinator.com/item?id=50234005"
title: "FirstTake – an open-source Claude Code skill for video shoots, offload to master"
article_title: "GitHub - emersonsklar/firsttake: Free, open-source Claude Code skill for video shoots: verified card offload, footage QC, multi-camera sync with no shared timecode, and edit, grade, mix, captions and delivery. No editing software required. · GitHub"
image: "https://repository-images.githubusercontent.com/1433456984/3181fe23-4c63-4c0e-87e7-b1b8bd3700b4"
author: "emersonsklr"
captured_at: "2026-10-25T18:11:12Z"
capture_tool: "hn-digest"
hn_id: 50234005
score: 1
comments: 0
posted_at: "2026-10-25T17:40:42Z"
tags:
  - hacker-news
---

# FirstTake – an open-source Claude Code skill for video shoots, offload to master

- HN: [50234005](https://news.ycombinator.com/item?id=50234005)
- Source: [github.com](https://github.com/emersonsklar/firsttake)
- Score: 1
- Comments: 0
- Posted: 2026-10-25T17:40:42Z

## Translation

Title: FirstTake – an open-source Claude Code skill for video shoots, offload to master
Article title: GitHub - emersonsklar/firsttake: Free, open-source Claude Code skill for video shoots: verified card offload, footage QC, multi-camera sync with no shared timecode, and edit, grade, mix, captions and delivery. No editing software required. · GitHub
Description: Free, open-source Claude Code skill for video shoots: verified card offload, footage QC, multi-camera sync with no shared timecode, and edit, grade, mix, captions and delivery. No editing software required. - emersonsklar/firsttake

Article text:
GitHub - emersonsklar/firsttake: Free, open-source Claude Code skill for video shoots: verified card offload, footage QC, multi-camera sync with no shared timecode, and edit, grade, mix, captions and delivery. No editing software required. · GitHub
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
emersonsklar
firsttake Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Free, open-source Claude Code skill for video shoots: verified card offload, footage QC, multi-camera sync with no shared timecode, and edit, grade, mix, captions and delivery. No editing software required.
sklarstad.com/firsttake/ Topics
Readme Apache-2.0 license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
58 Commits 58 Commits Folders and files
.claude-plugin .claude-plugin docs docs references references scripts scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SETUP-FOR-CLAUDE.md SETUP-FOR-CLAUDE.md SKILL.md SKILL.md requirements.txt requirements.txt View all files Repository files navigation
by Sklarstad · free and open source (Apache 2.0)
A Claude Code skill for the shoot you cannot ask to do again.
Cameras that share no clock. Forty minutes with someone who will not sit down twice. Nobody on set
whose only job is to watch the levels. At the other end: a master, a reviewer's transcript, social
cuts and pull clips — assembled by script, with every number in it measured.
Profile the kit, scout the location, design the setup, build the run sheet. The output is a
plan made of literal numbers — "lantern centre 75, 46 inches out, 30° down, 100%" — because a
principle has to be re-derived on set and a number does not. Lighting, camera and audio placement,
then the checklists, call order and day sheets that fall out of the same profile. Rooms can be
scouted from photographs before anyone travels.
The reason it starts with profiling: this is not a gear database. Every hardware number belongs
to a specific body, card, mic or drive, and almost none survive a kit change. A confident unmeasured
number is worse than no number, because the whole workflow depends on trusting it in a room where
you cannot check it.
📖 references/onboarding.md · references/location-scout.md · references/templates/
Rig to the plan, record, check the footage, secure the media. The plan from Act 1 arrives as
literal numbers so nothing has to be re-derived on set, and the checklists are ordered so the
things that cannot be fixed later happen first.
The reason it is built that way: nothing may require attention after "rolling." If you are also
the person on camera you cannot catch a focus hunt or a clipped level, and you cannot fix one
without stopping someone mid-answer. Stakes are asymmetric — forty minutes with a busy executive is
not a reshoot you get to request — so redundancy is weighted heavily and cleverness at zero.
Checking the footage means checking the files , not the monitor: log-aware exposure, focus
stability, frame timing and dropped frames, 8-bit hiding in a 10-bit container, per-channel audio.
Cross-camera frame-rate mismatch is the headline check, because it is the one that ruins an edit
silently.
Securing the media — a copy is not a backup until it has been read back
offload.py reads the card once , hashes as it reads, then re-reads every destination from
disk and compares. MHL 1.1, optional generation-chained ASC MHL, a JSON manifest, and a
self-contained HTML report you can hand to anyone.
The retention rule is deliberately strict, because the failure it guards against is silent and
final: cards stay until the deliverable is live and the archive is verified. Not until the
copy finished. Not until someone said thanks.
📖 references/qc-interpretation.md · references/offload-protocol.md
The largest part of the skill, and the part that keeps teaching. Six stages take a pile of cards
that share no clock to a signed-off master:
Reproducing someone else's graphics template runs through the whole of it: measure the reference
rather than guess it — size from ink height, family from width, weight from stroke thickness.
Either NLE, and a cut that can change hands
DaVinci Resolve · Premiere Pro (early access) · one plan, two backends
The timeline is described once, in an application-neutral plan, and a backend turns it into one
application's calls. That is what lets the person who shot it and the person who finishes it use
different software — and what lets a cut be handed over as the plan rather than as a project file
somebody else has to rebuild by hand.
The part that makes it trustworthy is a capability gate : where an application cannot reproduce
something exactly, the build stops and names it . It never approximates. The expensive failure
here was never "unsupported" — it is a build that reports success having quietly dropped the ease on
a fade or flattened an animated push-in to its first value. Both of those look like a finished
timeline.
✅ If you have a choice, use Resolve. It is driven by a Python module from a shell, it has
driven every delivered programme here, and it can script the grade. ⚠️ Premiere support is early
access : it is driven by a panel inside the app, and two things are simply absent — colour and
retime. It works, and it will show you its edges.
💸 And you do not have to buy anything. assemble_ffmpeg.py is a third backend that needs no
application at all — as does everything else here: offload and verification, QC, sync, transcription,
subtitles, loudness mastering, LUT fitting, speed ramps, graphic conforming, day sheets. It produces
a full-quality master, not a proxy : a shipped 2228-frame short rebuilt from its own plan sits at
1.08 codes of 255 against the delivered master across every frame. What you give up is the
editable timeline afterwards, not the quality. One condition, and it is measured rather than cautious: the grade must
be baked into the media once , never re-applied during assembly. Two implementations of one
colour transform agreed on geometry to r ≥ 0.9989 and were still 14–18% out in the grade, which is
how eight approved review renders turned out not to show what the NLE produced.
⚠️ Resolve free probably cannot be driven from a shell — the API covers both editions, but the
preference that lets an outside script connect is documented as Studio-only. That is unverified;
scripts/resolve_connect.py --check settles it on a machine that has the free version. Either way
the ffmpeg route is there.
Each application is better at something, and the differences are not small. Resolve can script
colour; Premiere cannot script colour at all — neither an ASC CDL nor a LUT — so a graded
handover has to arrive with the grade already in the media. Premiere can keyframe any parameter and
write an exact 29.97; in Resolve the first needs a composition and the second needs a committed
preset. The gate knows every one of those, and nle_ir.py --compare tells you what the far end will
refuse before you offer it a handover.
📖 references/post-production.md · references/nle-backends.md · references/resolve-api.md ·
references/premiere-uxp.md · references/grade-match.md
Exhibit A — the clap that doesn't save you
Clap at the top of the take. You should still do it — it is free and survives everything. But a clap
gives you one point in time , and the expensive failure is not offset, it is drift : 29.97 and
30.000 are both sold as "30p" and pull apart about a frame every 33 seconds. Over forty minutes two
angles end up roughly 70 frames apart and no single sync point fixes it. Add a third body and it
is three pairs of clocks, not one.
sync_cameras.py probes in up to ten places, so the spread between probes is the drift
measurement :
It correlates onset envelopes , not waveforms — so it works between completely different
microphones, because two mics in one room disagree about spectrum and agree about when things
start .
Exhibit B — a vertical cut is a budget, not a crop
A seated subject's head swayed 588 px ; the union of head positions spanned 1051 px of a
1215 px window. No static window holds that with margin everywhere.
The answer is not to abandon a centred frame. Set an excursion budget in pixels and seconds,
then take the crop nearest the compositional ideal that stays inside it — here 48 px for up to
1.25 s, while the subject leans forward and looks down , so no gaze room is wanted at that
instant. And 1215 of 3840 px at full height is a 0.889 downscale , so sharpness never enters the
argument.
⚠️ Probe density has to scale with the span. A fixed probe count reused on a longer clip missed
the subject's furthest lean entirely — the sparse answer would have clipped a head by ~68 px.
Exhibit C — on target, and still wrong
Integrated loudness, true peak, LRA and per-turn levelling all correct — turn-to-turn spread
1.4 LU — and one phrase still sat +2.7 LU over its speaker's median word. Twelve phrases sat
at +4 LU or more , every one of them the same speaker, lifting on question intonation. Nothing in
the chain was looking at a phrase.
And the trap underneath: a pre-limiter trim on a phrase already pinned at the ceiling barely
changes its loudness. Thirteen trims of 1.75–4.00 dB delivered 0.00–2.27 dB , because trimming
ahead of a limiter mostly stops the limiter reducing. Keep it for the crest it buys, then apply the
measured remainder after the limiter, where gain is exact.
The failures are quiet. A sync 40 ms out looks fine on a laptop and wrong on a big screen. A
timeline at the wrong playback rate shows a contented green number while dropping a fifth of the
audio. A guard can report success and do nothing. A check can be written so it cannot fail.
So every claim carries CONFIRMED / LIKELY / UNVERIFIED and a note on how it was established,
and the references read like a post-mortem in places — because those are the expensive ones:
a still that ignored its own length and silently moved the clip after it · numbered PNGs that
imported as one image sequence, so 11 of 12 captions never existed · a loudness meter measured in
the wrong channel layout, 3 dB out · a cached asset keyed on a filename instead of its content ·
a derived cache of absolute frame numbers that rendered 33 frames long while the build reported
no problems
The references are written to be read when you need them. These are written to be read in order.
The easy way: open this folder in Claude Code and say "set this up for me" . Any of the ways
it runs will do — the terminal CLI , the desktop app 's Code tab, claude.ai/code in a
browser, or the VS Code / JetBrains extensions. The desktop app and the web app need no terminal
at all. It then follows

[truncated]

## Original Extract

Free, open-source Claude Code skill for video shoots: verified card offload, footage QC, multi-camera sync with no shared timecode, and edit, grade, mix, captions and delivery. No editing software required. - emersonsklar/firsttake

GitHub - emersonsklar/firsttake: Free, open-source Claude Code skill for video shoots: verified card offload, footage QC, multi-camera sync with no shared timecode, and edit, grade, mix, captions and delivery. No editing software required. · GitHub
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
emersonsklar
firsttake Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Free, open-source Claude Code skill for video shoots: verified card offload, footage QC, multi-camera sync with no shared timecode, and edit, grade, mix, captions and delivery. No editing software required.
sklarstad.com/firsttake/ Topics
Readme Apache-2.0 license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
58 Commits 58 Commits Folders and files
.claude-plugin .claude-plugin docs docs references references scripts scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SETUP-FOR-CLAUDE.md SETUP-FOR-CLAUDE.md SKILL.md SKILL.md requirements.txt requirements.txt View all files Repository files navigation
by Sklarstad · free and open source (Apache 2.0)
A Claude Code skill for the shoot you cannot ask to do again.
Cameras that share no clock. Forty minutes with someone who will not sit down twice. Nobody on set
whose only job is to watch the levels. At the other end: a master, a reviewer's transcript, social
cuts and pull clips — assembled by script, with every number in it measured.
Profile the kit, scout the location, design the setup, build the run sheet. The output is a
plan made of literal numbers — "lantern centre 75, 46 inches out, 30° down, 100%" — because a
principle has to be re-derived on set and a number does not. Lighting, camera and audio placement,
then the checklists, call order and day sheets that fall out of the same profile. Rooms can be
scouted from photographs before anyone travels.
The reason it starts with profiling: this is not a gear database. Every hardware number belongs
to a specific body, card, mic or drive, and almost none survive a kit change. A confident unmeasured
number is worse than no number, because the whole workflow depends on trusting it in a room where
you cannot check it.
📖 references/onboarding.md · references/location-scout.md · references/templates/
Rig to the plan, record, check the footage, secure the media. The plan from Act 1 arrives as
literal numbers so nothing has to be re-derived on set, and the checklists are ordered so the
things that cannot be fixed later happen first.
The reason it is built that way: nothing may require attention after "rolling." If you are also
the person on camera you cannot catch a focus hunt or a clipped level, and you cannot fix one
without stopping someone mid-answer. Stakes are asymmetric — forty minutes with a busy executive is
not a reshoot you get to request — so redundancy is weighted heavily and cleverness at zero.
Checking the footage means checking the files , not the monitor: log-aware exposure, focus
stability, frame timing and dropped frames, 8-bit hiding in a 10-bit container, per-channel audio.
Cross-camera frame-rate mismatch is the headline check, because it is the one that ruins an edit
silently.
Securing the media — a copy is not a backup until it has been read back
offload.py reads the card once , hashes as it reads, then re-reads every destination from
disk and compares. MHL 1.1, optional generation-chained ASC MHL, a JSON manifest, and a
self-contained HTML report you can hand to anyone.
The retention rule is deliberately strict, because the failure it guards against is silent and
final: cards stay until the deliverable is live and the archive is verified. Not until the
copy finished. Not until someone said thanks.
📖 references/qc-interpretation.md · references/offload-protocol.md
The largest part of the skill, and the part that keeps teaching. Six stages take a pile of cards
that share no clock to a signed-off master:
Reproducing someone else's graphics template runs through the whole of it: measure the reference
rather than guess it — size from ink height, family from width, weight from stroke thickness.
Either NLE, and a cut that can change hands
DaVinci Resolve · Premiere Pro (early access) · one plan, two backends
The timeline is described once, in an application-neutral plan, and a backend turns it into one
application's calls. That is what lets the person who shot it and the person who finishes it use
different software — and what lets a cut be handed over as the plan rather than as a project file
somebody else has to rebuild by hand.
The part that makes it trustworthy is a capability gate : where an application cannot reproduce
something exactly, the build stops and names it . It never approximates. The expensive failure
here was never "unsupported" — it is a build that reports success having quietly dropped the ease on
a fade or flattened an animated push-in to its first value. Both of those look like a finished
timeline.
✅ If you have a choice, use Resolve. It is driven by a Python module from a shell, it has
driven every delivered programme here, and it can script the grade. ⚠️ Premiere support is early
access : it is driven by a panel inside the app, and two things are simply absent — colour and
retime. It works, and it will show you its edges.
💸 And you do not have to buy anything. assemble_ffmpeg.py is a third backend that needs no
application at all — as does everything else here: offload and verification, QC, sync, transcription,
subtitles, loudness mastering, LUT fitting, speed ramps, graphic conforming, day sheets. It produces
a full-quality master, not a proxy : a shipped 2228-frame short rebuilt from its own plan sits at
1.08 codes of 255 against the delivered master across every frame. What you give up is the
editable timeline afterwards, not the quality. One condition, and it is measured rather than cautious: the grade must
be baked into the media once , never re-applied during assembly. Two implementations of one
colour transform agreed on geometry to r ≥ 0.9989 and were still 14–18% out in the grade, which is
how eight approved review renders turned out not to show what the NLE produced.
⚠️ Resolve free probably cannot be driven from a shell — the API covers both editions, but the
preference that lets an outside script connect is documented as Studio-only. That is unverified;
scripts/resolve_connect.py --check settles it on a machine that has the free version. Either way
the ffmpeg route is there.
Each application is better at something, and the differences are not small. Resolve can script
colour; Premiere cannot script colour at all — neither an ASC CDL nor a LUT — so a graded
handover has to arrive with the grade already in the media. Premiere can keyframe any parameter and
write an exact 29.97; in Resolve the first needs a composition and the second needs a committed
preset. The gate knows every one of those, and nle_ir.py --compare tells you what the far end will
refuse before you offer it a handover.
📖 references/post-production.md · references/nle-backends.md · references/resolve-api.md ·
references/premiere-uxp.md · references/grade-match.md
Exhibit A — the clap that doesn't save you
Clap at the top of the take. You should still do it — it is free and survives everything. But a clap
gives you one point in time , and the expensive failure is not offset, it is drift : 29.97 and
30.000 are both sold as "30p" and pull apart about a frame every 33 seconds. Over forty minutes two
angles end up roughly 70 frames apart and no single sync point fixes it. Add a third body and it
is three pairs of clocks, not one.
sync_cameras.py probes in up to ten places, so the spread between probes is the drift
measurement :
It correlates onset envelopes , not waveforms — so it works between completely different
microphones, because two mics in one room disagree about spectrum and agree about when things
start .
Exhibit B — a vertical cut is a budget, not a crop
A seated subject's head swayed 588 px ; the union of head positions spanned 1051 px of a
1215 px window. No static window holds that with margin everywhere.
The answer is not to abandon a centred frame. Set an excursion budget in pixels and seconds,
then take the crop nearest the compositional ideal that stays inside it — here 48 px for up to
1.25 s, while the subject leans forward and looks down , so no gaze room is wanted at that
instant. And 1215 of 3840 px at full height is a 0.889 downscale , so sharpness never enters the
argument.
⚠️ Probe density has to scale with the span. A fixed probe count reused on a longer clip missed
the subject's furthest lean entirely — the sparse answer would have clipped a head by ~68 px.
Exhibit C — on target, and still wrong
Integrated loudness, true peak, LRA and per-turn levelling all correct — turn-to-turn spread
1.4 LU — and one phrase still sat +2.7 LU over its speaker's median word. Twelve phrases sat
at +4 LU or more , every one of them the same speaker, lifting on question intonation. Nothing in
the chain was looking at a phrase.
And the trap underneath: a pre-limiter trim on a phrase already pinned at the ceiling barely
changes its loudness. Thirteen trims of 1.75–4.00 dB delivered 0.00–2.27 dB , because trimming
ahead of a limiter mostly stops the limiter reducing. Keep it for the crest it buys, then apply the
measured remainder after the limiter, where gain is exact.
The failures are quiet. A sync 40 ms out looks fine on a laptop and wrong on a big screen. A
timeline at the wrong playback rate shows a contented green number while dropping a fifth of the
audio. A guard can report success and do nothing. A check can be written so it cannot fail.
So every claim carries CONFIRMED / LIKELY / UNVERIFIED and a note on how it was established,
and the references read like a post-mortem in places — because those are the expensive ones:
a still that ignored its own length and silently moved the clip after it · numbered PNGs that
imported as one image sequence, so 11 of 12 captions never existed · a loudness meter measured in
the wrong channel layout, 3 dB out · a cached asset keyed on a filename instead of its content ·
a derived cache of absolute frame numbers that rendered 33 frames long while the build reported
no problems
The references are written to be read when you need them. These are written to be read in order.
The easy way: open this folder in Claude Code and say "set this up for me" . Any of the ways
it runs will do — the terminal CLI , the desktop app 's Code tab, claude.ai/code in a
browser, or the VS Code / JetBrains extensions. The desktop app and the web app need no terminal
at all. It then follows

[truncated]
