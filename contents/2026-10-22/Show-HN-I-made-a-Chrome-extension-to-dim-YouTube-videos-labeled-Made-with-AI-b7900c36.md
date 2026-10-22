---
source: "https://github.com/Nova-47/ai-label-filter-for-youtube"
hn_url: "https://news.ycombinator.com/item?id=50194592"
title: "Show HN: I made a Chrome extension to dim YouTube videos labeled \"Made with AI\""
article_title: "GitHub - Nova-47/ai-label-filter-for-youtube: Chrome extension that dims or hides YouTube videos YouTube labels as made with AI, and shows why. No server. · GitHub"
image: "https://opengraph.githubassets.com/e99555e8c499a0e4f18ff921a68a8ca5c30790fa1e64a2b497594da95aa3b6ee/Nova-47/ai-label-filter-for-youtube"
author: "nova-47"
captured_at: "2026-10-22T13:31:41Z"
capture_tool: "hn-digest"
hn_id: 50194592
score: 1
comments: 0
posted_at: "2026-10-22T12:45:52Z"
tags:
  - hacker-news
---

# Show HN: I made a Chrome extension to dim YouTube videos labeled "Made with AI"

- HN: [50194592](https://news.ycombinator.com/item?id=50194592)
- Source: [github.com](https://github.com/Nova-47/ai-label-filter-for-youtube)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T12:45:52Z

## Translation

Title: Show HN: I made a Chrome extension to dim YouTube videos labeled "Made with AI"
Article title: GitHub - Nova-47/ai-label-filter-for-youtube: Chrome extension that dims or hides YouTube videos YouTube labels as made with AI, and shows why. No server. · GitHub
Description: Chrome extension that dims or hides YouTube videos YouTube labels as made with AI, and shows why. No server. - Nova-47/ai-label-filter-for-youtube

Article text:
GitHub - Nova-47/ai-label-filter-for-youtube: Chrome extension that dims or hides YouTube videos YouTube labels as made with AI, and shows why. No server. · GitHub
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
Nova-47
ai-label-filter-for-youtube Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Chrome extension that dims or hides YouTube videos YouTube labels as made with AI, and shows why. No server.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
50 Commits 50 Commits Folders and files
docs docs icons icons scripts scripts store store tests tests .gitignore .gitignore LICENSE LICENSE PRIVACY.md PRIVACY.md README.kr.md README.kr.md README.md README.md background.js background.js content.css content.css content.js content.js manifest.json manifest.json popup.html popup.html popup.js popup.js shared.js shared.js View all files Repository files navigation
A Chrome extension that dims or hides videos YouTube labels as made with AI and shows why. It uses YouTube's label. Its only inferred rule identifies channels where most recently checked videos have labels; the chip shows the count, and you can disable it. No server or account.
Install it from the Chrome Web Store , or download the zip from Releases and load it unpacked ( steps ). Try it in 60 seconds .
It catches only videos YouTube labels, and most AI videos are not labelled. For more coverage, turn on the community list AiSList or "Check every visible video" (both are off by default).
A video counts as AI only if YouTube shows its AI badge under the title. It uses no model, title heuristic, or thumbnail heuristic. Filtered cards are dimmed by default and show a chip with the reason, Show , and, when the card links its channel, Allow channel . Select Hide in the popup to remove them.
There is no server, dependency, or build step. The extension requires storage and a content script on www.youtube.com . Every network request is made in one of two functions, listed below .
Install it from the Chrome Web Store .
Download ai-label-filter-for-youtube-<version>.zip from Releases and unzip it into a folder. (Or clone the repository: it loads as is, there is nothing to build.)
Open chrome://extensions and turn on Developer mode (top right).
Click Load unpacked and select the folder.
By default, feeds change only for videos the extension already knows about, so a fresh install shows nothing. To test it:
Open this video , which YouTube labels as made with AI. After a few seconds a banner appears in the bottom-left corner: "YouTube labels this video as made with AI", with Hide this channel and Dismiss .
Go to the channel's page (@eigi_and_ai) or search for the video. Wherever its card appears, it is now dimmed with the chip "YouTube AI label", because the result is saved on your device.
To filter feeds more widely, open the popup and turn on Hide channels listed in AiSList (a community list of about 29,000 channels; Chrome asks for permission to download it) or Check every visible video (asks YouTube about each video on screen, one request at a time).
YouTube shows an "AI" badge under the title of videos it labels as made with AI. The badge is in the data YouTube's web app loads for a watch page, but not in the data for home, search or sidebar cards. To check a video's label, the extension sends one small request to /youtubei/v1/next , the endpoint the web app uses to load a watch page.
A video counts as AI only if the reply contains that badge: a metadataBadgeRenderer with style BADGE_STYLE_TYPE_SIMPLE and icon INFO . Auto-dubbed videos get a badge with the same style but a different icon, and do not count. The icon is checked, not the badge text, so YouTube's interface language does not matter.
Where
Requests
What happens
Watch page ( /watch?v=… )
1 for the video you opened, unless its result is already saved
If the video is labelled, a banner offers Hide this channel and Dismiss .
Feeds (default)
none
Cards are filtered only from what is already on your device: saved results, your channel lists, channels where most checked videos are labelled, and AiSList if it is on.
Feeds with Check every visible video on
1 per card on screen that has no saved result
Labelled videos are filtered as they are found.
Cards are recognised on home, subscriptions, search, channel pages, playlists, the watch-page sidebar and Shorts shelves.
Always shown channels : a channel on this list is never filtered.
Hidden channels : you added the channel with Hide this channel . Chip: "Your blocklist".
YouTube AI label : the video itself is labelled. Chip: "YouTube AI label".
Channels where most checked videos are labelled (on by default; turn it off in the popup): at least 5 of the last 20 different videos from the channel were checked on your device, and at least 60% are labelled. Checking a video again counts it once, moves it to the newest spot, and uses the latest result. Chip: "Channel: N of last M checked videos AI-labelled", where M is the sample size (at most 20). These cards are always dimmed with their reason, never hidden, even in Hide mode, because this is an inference rather than a label on the video. The popup switch Also dim channels whose recent videos are mostly AI-labelled turns this rule off. Checked videos are still counted, so it applies as soon as you turn it back on. The popup lists channels currently dimmed by this rule with "N of last M" and an Allow button. Allow adds the channel to Always shown channels, which overrides this rule.
Listed in AiSList : only if you turned AiSList on. Chip: "Listed in AiSList".
Cards that match no rule are shown. Rules 1, 2, 4 and 5 need a channel link on the card (a /@handle or /channel/UC… link).
Results are saved for 30 days for a video without the label (YouTube may add one later) and 180 days for a labelled one.
Dim + reason (default): the card is faded and a chip shows the reason, with Show (shows that video normally in this tab until you reload) and Allow channel (adds the channel to Always shown channels; only offered when the card has a channel link).
Hide : filtered cards are removed from view, except cards filtered only because most checked videos of their channel are labelled, which stay dimmed. To see what was hidden, open the popup and press Reveal under "This page": that tab switches to the dimmed view, with the same chips, until you press Hide again or reload.
Watch page banner : shown when the open video is labelled, in either mode. Hide this channel adds the channel to Hidden channels (the button is left out if the channel is unknown or already on that list). No banner appears for a channel on Always shown channels. The video itself is never hidden.
If anything fails (network error, timeout, unexpected reply, removed video), the video is shown and nothing is saved.
The watch page, card checks and the daily self-check all send the same request, from the youtube.com page itself (the content script):
POST https://www.youtube.com/youtubei/v1/next?prettyPrint=false&fields=<mask>
Body: {"context":{"client":{"clientName":"WEB","clientVersion":"…","hl":"en","gl":"US"}},"videoId":"…"} . Only the video ID changes. The client version is read from the YouTube page; until it is found there, a fixed recent version is sent.
Your cookies are left out ( credentials: 'omit' ), so the request is not tied to your YouTube or Google account, and YouTube answers as it would a logged-out visitor.
The field mask asks for two things only: the badges under the title, and the channel's ID and @handle. The reply is about 0.4 KB instead of about 400 KB.
One at a time, browser-wide. All youtube.com tabs share one lock, and each request starts at least 1 second after the previous reply. Card checks add a random 0–1 s on top.
Each request times out after 10 seconds.
There is no retry without the mask.
When YouTube does not answer normally, requests pause in every tab:
While requests are paused, cards are still filtered from what is already saved.
With default settings the extension sends one request per watch page whose result is not saved yet, plus the self-check's one or two requests a day.
Once every 24 hours per browser (not per tab), when a YouTube page loads or you navigate, the extension asks about two videos that carried the label on 2026-10-13: yHG-pivADos (@eigi_and_ai) and, only if the first does not show the label, Xpg-JIuJ_cU (@foreverharmony_official). It runs even when "Check every visible video" is off, so it normally sends one or two requests a day. It does not run while filtering is off or requests are paused; if requests are paused before it has sent anything, it runs once they are allowed again. If the page closes or reloads before the check finishes, it runs again on a page opened 5 minutes or more later. Its results are not saved and do not count towards channel flags.
There is no server, and nothing is sent to the developer. These are the extension's only requests:
As with any web request, YouTube and GitHub see your IP address and browser user agent.
Everything is stored in chrome.storage.local on your device. Nothing is stored in chrome.storage.sync or elsewhere:
Hidden channels and Always shown channels (channel ID and/or @handle, and when it was added)
results per video: video ID, labelled or not, the day it was checked and the day its card was last seen; for labelled videos also the channel ID and @handle. At most 20,000 videos.
per channel seen in a reply: its ID, its @handle, the IDs of the last 20 different videos checked from it, each with whether it is labelled and the time it was checked, and the day of its last result. At most 5,000 channels; above that, channels with fewer labelled videos and older results are dropped first. Dimmed channels and channels on your lists are always kept.
if AiSList is on: the downloaded list and its download details (entry count, ETag, times, last error)
whether requests are paused, until when and why, and when the last request was sent
when the self-check last ran and its result
The number of filtered cards is not stored: the popup asks the open YouTube tab each time. Uninstalling the extension deletes all of it.
It uses no analytics, telemetry, YouTube Data API, or account. Full policy: PRIVACY.md .
Permission
Why it is needed
storage
Keeps the data listed above on your device.
Content script on *://www.youtube.com/*
Reads video and channel links on YouTube pages, dims or hides cards, shows the watch-page banner, and sends the label request from the YouTube page itself. Because the request comes from www.youtube.com , no host permission is needed.
Optional: https://raw.githubusercontent.com/*

[truncated]

## Original Extract

Chrome extension that dims or hides YouTube videos YouTube labels as made with AI, and shows why. No server. - Nova-47/ai-label-filter-for-youtube

GitHub - Nova-47/ai-label-filter-for-youtube: Chrome extension that dims or hides YouTube videos YouTube labels as made with AI, and shows why. No server. · GitHub
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
Nova-47
ai-label-filter-for-youtube Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Chrome extension that dims or hides YouTube videos YouTube labels as made with AI, and shows why. No server.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
50 Commits 50 Commits Folders and files
docs docs icons icons scripts scripts store store tests tests .gitignore .gitignore LICENSE LICENSE PRIVACY.md PRIVACY.md README.kr.md README.kr.md README.md README.md background.js background.js content.css content.css content.js content.js manifest.json manifest.json popup.html popup.html popup.js popup.js shared.js shared.js View all files Repository files navigation
A Chrome extension that dims or hides videos YouTube labels as made with AI and shows why. It uses YouTube's label. Its only inferred rule identifies channels where most recently checked videos have labels; the chip shows the count, and you can disable it. No server or account.
Install it from the Chrome Web Store , or download the zip from Releases and load it unpacked ( steps ). Try it in 60 seconds .
It catches only videos YouTube labels, and most AI videos are not labelled. For more coverage, turn on the community list AiSList or "Check every visible video" (both are off by default).
A video counts as AI only if YouTube shows its AI badge under the title. It uses no model, title heuristic, or thumbnail heuristic. Filtered cards are dimmed by default and show a chip with the reason, Show , and, when the card links its channel, Allow channel . Select Hide in the popup to remove them.
There is no server, dependency, or build step. The extension requires storage and a content script on www.youtube.com . Every network request is made in one of two functions, listed below .
Install it from the Chrome Web Store .
Download ai-label-filter-for-youtube-<version>.zip from Releases and unzip it into a folder. (Or clone the repository: it loads as is, there is nothing to build.)
Open chrome://extensions and turn on Developer mode (top right).
Click Load unpacked and select the folder.
By default, feeds change only for videos the extension already knows about, so a fresh install shows nothing. To test it:
Open this video , which YouTube labels as made with AI. After a few seconds a banner appears in the bottom-left corner: "YouTube labels this video as made with AI", with Hide this channel and Dismiss .
Go to the channel's page (@eigi_and_ai) or search for the video. Wherever its card appears, it is now dimmed with the chip "YouTube AI label", because the result is saved on your device.
To filter feeds more widely, open the popup and turn on Hide channels listed in AiSList (a community list of about 29,000 channels; Chrome asks for permission to download it) or Check every visible video (asks YouTube about each video on screen, one request at a time).
YouTube shows an "AI" badge under the title of videos it labels as made with AI. The badge is in the data YouTube's web app loads for a watch page, but not in the data for home, search or sidebar cards. To check a video's label, the extension sends one small request to /youtubei/v1/next , the endpoint the web app uses to load a watch page.
A video counts as AI only if the reply contains that badge: a metadataBadgeRenderer with style BADGE_STYLE_TYPE_SIMPLE and icon INFO . Auto-dubbed videos get a badge with the same style but a different icon, and do not count. The icon is checked, not the badge text, so YouTube's interface language does not matter.
Where
Requests
What happens
Watch page ( /watch?v=… )
1 for the video you opened, unless its result is already saved
If the video is labelled, a banner offers Hide this channel and Dismiss .
Feeds (default)
none
Cards are filtered only from what is already on your device: saved results, your channel lists, channels where most checked videos are labelled, and AiSList if it is on.
Feeds with Check every visible video on
1 per card on screen that has no saved result
Labelled videos are filtered as they are found.
Cards are recognised on home, subscriptions, search, channel pages, playlists, the watch-page sidebar and Shorts shelves.
Always shown channels : a channel on this list is never filtered.
Hidden channels : you added the channel with Hide this channel . Chip: "Your blocklist".
YouTube AI label : the video itself is labelled. Chip: "YouTube AI label".
Channels where most checked videos are labelled (on by default; turn it off in the popup): at least 5 of the last 20 different videos from the channel were checked on your device, and at least 60% are labelled. Checking a video again counts it once, moves it to the newest spot, and uses the latest result. Chip: "Channel: N of last M checked videos AI-labelled", where M is the sample size (at most 20). These cards are always dimmed with their reason, never hidden, even in Hide mode, because this is an inference rather than a label on the video. The popup switch Also dim channels whose recent videos are mostly AI-labelled turns this rule off. Checked videos are still counted, so it applies as soon as you turn it back on. The popup lists channels currently dimmed by this rule with "N of last M" and an Allow button. Allow adds the channel to Always shown channels, which overrides this rule.
Listed in AiSList : only if you turned AiSList on. Chip: "Listed in AiSList".
Cards that match no rule are shown. Rules 1, 2, 4 and 5 need a channel link on the card (a /@handle or /channel/UC… link).
Results are saved for 30 days for a video without the label (YouTube may add one later) and 180 days for a labelled one.
Dim + reason (default): the card is faded and a chip shows the reason, with Show (shows that video normally in this tab until you reload) and Allow channel (adds the channel to Always shown channels; only offered when the card has a channel link).
Hide : filtered cards are removed from view, except cards filtered only because most checked videos of their channel are labelled, which stay dimmed. To see what was hidden, open the popup and press Reveal under "This page": that tab switches to the dimmed view, with the same chips, until you press Hide again or reload.
Watch page banner : shown when the open video is labelled, in either mode. Hide this channel adds the channel to Hidden channels (the button is left out if the channel is unknown or already on that list). No banner appears for a channel on Always shown channels. The video itself is never hidden.
If anything fails (network error, timeout, unexpected reply, removed video), the video is shown and nothing is saved.
The watch page, card checks and the daily self-check all send the same request, from the youtube.com page itself (the content script):
POST https://www.youtube.com/youtubei/v1/next?prettyPrint=false&fields=<mask>
Body: {"context":{"client":{"clientName":"WEB","clientVersion":"…","hl":"en","gl":"US"}},"videoId":"…"} . Only the video ID changes. The client version is read from the YouTube page; until it is found there, a fixed recent version is sent.
Your cookies are left out ( credentials: 'omit' ), so the request is not tied to your YouTube or Google account, and YouTube answers as it would a logged-out visitor.
The field mask asks for two things only: the badges under the title, and the channel's ID and @handle. The reply is about 0.4 KB instead of about 400 KB.
One at a time, browser-wide. All youtube.com tabs share one lock, and each request starts at least 1 second after the previous reply. Card checks add a random 0–1 s on top.
Each request times out after 10 seconds.
There is no retry without the mask.
When YouTube does not answer normally, requests pause in every tab:
While requests are paused, cards are still filtered from what is already saved.
With default settings the extension sends one request per watch page whose result is not saved yet, plus the self-check's one or two requests a day.
Once every 24 hours per browser (not per tab), when a YouTube page loads or you navigate, the extension asks about two videos that carried the label on 2026-10-13: yHG-pivADos (@eigi_and_ai) and, only if the first does not show the label, Xpg-JIuJ_cU (@foreverharmony_official). It runs even when "Check every visible video" is off, so it normally sends one or two requests a day. It does not run while filtering is off or requests are paused; if requests are paused before it has sent anything, it runs once they are allowed again. If the page closes or reloads before the check finishes, it runs again on a page opened 5 minutes or more later. Its results are not saved and do not count towards channel flags.
There is no server, and nothing is sent to the developer. These are the extension's only requests:
As with any web request, YouTube and GitHub see your IP address and browser user agent.
Everything is stored in chrome.storage.local on your device. Nothing is stored in chrome.storage.sync or elsewhere:
Hidden channels and Always shown channels (channel ID and/or @handle, and when it was added)
results per video: video ID, labelled or not, the day it was checked and the day its card was last seen; for labelled videos also the channel ID and @handle. At most 20,000 videos.
per channel seen in a reply: its ID, its @handle, the IDs of the last 20 different videos checked from it, each with whether it is labelled and the time it was checked, and the day of its last result. At most 5,000 channels; above that, channels with fewer labelled videos and older results are dropped first. Dimmed channels and channels on your lists are always kept.
if AiSList is on: the downloaded list and its download details (entry count, ETag, times, last error)
whether requests are paused, until when and why, and when the last request was sent
when the self-check last ran and its result
The number of filtered cards is not stored: the popup asks the open YouTube tab each time. Uninstalling the extension deletes all of it.
It uses no analytics, telemetry, YouTube Data API, or account. Full policy: PRIVACY.md .
Permission
Why it is needed
storage
Keeps the data listed above on your device.
Content script on *://www.youtube.com/*
Reads video and channel links on YouTube pages, dims or hides cards, shows the watch-page banner, and sends the label request from the YouTube page itself. Because the request comes from www.youtube.com , no host permission is needed.
Optional: https://raw.githubusercontent.com/*

[truncated]
