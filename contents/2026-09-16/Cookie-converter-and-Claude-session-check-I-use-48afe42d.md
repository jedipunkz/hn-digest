---
source: "https://claudecookie.com/"
hn_url: "https://news.ycombinator.com/item?id=49733195"
title: "Cookie converter and Claude session check I use"
article_title: "Cookie Converter & Claude cookie checker — cookies.txt to JSON"
image: "https://claudecookie.com/og.png"
author: "kirill_orlov"
captured_at: "2026-09-16T21:55:36Z"
capture_tool: "hn-digest"
hn_id: 49733195
score: 1
comments: 0
posted_at: "2026-09-16T21:20:43Z"
tags:
  - hacker-news
---

# Cookie converter and Claude session check I use

- HN: [49733195](https://news.ycombinator.com/item?id=49733195)
- Source: [claudecookie.com](https://claudecookie.com/)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T21:20:43Z

## Translation

Title: Cookie converter and Claude session check I use
Article title: Cookie Converter & Claude cookie checker — cookies.txt to JSON
Description: Convert cookies between Netscape cookies.txt and JSON — Cookie-Editor, Puppeteer, key-value, header. Format auto-detected. Free, runs entirely in your browser.

Article text:
Cookie Converter & Claude cookie checker — cookies.txt to JSON Skip to content claude cookie Cookie converter Check cookie Netscape cookies.txt JSON formats Privacy English English Русский 中文 Cookie converter Check cookie Netscape cookies.txt JSON formats Privacy Language
English English Русский 中文 Cookie converter
Convert cookies between Netscape cookies.txt and JSON.
cookies.txt to JSON and back — Cookie-Editor, Puppeteer, key-value, or a raw Cookie header. Paste it and the format is worked out for you. Cookies, in whatever format you need them.
The conversion runs entirely in your browser.
Drop your cookie file cookies.txt or .json Output Netscape Cookie-Editor Puppeteer Key-Value Header The converted cookies will appear here.
Copy Download cookies.json Convert Also here
Is your Claude session cookie still alive?
Paste a Claude session export and see in seconds whether it is live — the account, the plan, and how much of the 5-hour and weekly limits are left.
Drop a cookies.txt file onto the panel, paste a JSON export, or paste a raw Cookie header. The file is read by your browser and goes no further.
A Netscape file is recognised by its tab layout, Cookie-Editor by its expirationDate and storeId fields, Puppeteer by its expires field. The opposite format is selected for the output.
Take the result to the clipboard, or download it as cookies.txt or cookies.json. Use the swap button to run the conversion the other way.
cookies.txt — curl, wget, yt-dlp
What is the Netscape cookie format?
A plain text file with one cookie per line and seven tab-separated fields: domain, include-subdomains flag, path, secure flag, expiry, name and value. It was introduced by Netscape Navigator and is still what curl, wget and yt-dlp read and write today.
Is anything I paste sent to a server?
The conversion itself runs entirely in JavaScript in your browser. The Check cookie page is different: it sends an encrypted paste to this site’s backend and to Anthropic. The site also records anonymous usage statistics — which formats are converted and how often. There are no third-party analytics or trackers.
Which JSON format should I pick?
Cookie-Editor if you are importing back into a browser extension. Puppeteer if you are feeding page.setCookie() or context.addCookies(). Key-value for a requests or axios session. Header if you just need something to paste after curl -H.
Why does a converted line start with #HttpOnly_?
That prefix is how curl marks an http-only cookie in a cookies.txt file. It looks like a comment so older readers skip it safely, while curl and yt-dlp understand it. Cookies without the flag are written normally.
What does an expiry of 0 or -1 mean?
Both mean a session cookie — one that dies when the browser closes. Netscape files write 0, Puppeteer writes -1, and Cookie-Editor drops the field and sets session to true instead. All three are read correctly and converted to whichever spelling the target format expects.
Why do the key-value and header formats lose data?
They only store names and values. Domain, path, expiry, secure and http-only have nowhere to go. That is fine when you are attaching cookies to a single request, but you cannot convert back to a complete cookies.txt afterwards without supplying a domain.
Can I take cookies from a browser extension into yt-dlp or curl?
Yes — that is the most common reason people land here. Export from Cookie-Editor as JSON, paste it in, and the output is a Netscape cookies.txt file you can pass to yt-dlp --cookies or curl -b.
claude cookie A cookie format converter that runs entirely in your browser.
Not affiliated with, endorsed by, or connected to Anthropic. Claude is a trademark of Anthropic PBC.
© 2026 claudecookie.com. All rights reserved.

## Original Extract

Convert cookies between Netscape cookies.txt and JSON — Cookie-Editor, Puppeteer, key-value, header. Format auto-detected. Free, runs entirely in your browser.

Cookie Converter & Claude cookie checker — cookies.txt to JSON Skip to content claude cookie Cookie converter Check cookie Netscape cookies.txt JSON formats Privacy English English Русский 中文 Cookie converter Check cookie Netscape cookies.txt JSON formats Privacy Language
English English Русский 中文 Cookie converter
Convert cookies between Netscape cookies.txt and JSON.
cookies.txt to JSON and back — Cookie-Editor, Puppeteer, key-value, or a raw Cookie header. Paste it and the format is worked out for you. Cookies, in whatever format you need them.
The conversion runs entirely in your browser.
Drop your cookie file cookies.txt or .json Output Netscape Cookie-Editor Puppeteer Key-Value Header The converted cookies will appear here.
Copy Download cookies.json Convert Also here
Is your Claude session cookie still alive?
Paste a Claude session export and see in seconds whether it is live — the account, the plan, and how much of the 5-hour and weekly limits are left.
Drop a cookies.txt file onto the panel, paste a JSON export, or paste a raw Cookie header. The file is read by your browser and goes no further.
A Netscape file is recognised by its tab layout, Cookie-Editor by its expirationDate and storeId fields, Puppeteer by its expires field. The opposite format is selected for the output.
Take the result to the clipboard, or download it as cookies.txt or cookies.json. Use the swap button to run the conversion the other way.
cookies.txt — curl, wget, yt-dlp
What is the Netscape cookie format?
A plain text file with one cookie per line and seven tab-separated fields: domain, include-subdomains flag, path, secure flag, expiry, name and value. It was introduced by Netscape Navigator and is still what curl, wget and yt-dlp read and write today.
Is anything I paste sent to a server?
The conversion itself runs entirely in JavaScript in your browser. The Check cookie page is different: it sends an encrypted paste to this site’s backend and to Anthropic. The site also records anonymous usage statistics — which formats are converted and how often. There are no third-party analytics or trackers.
Which JSON format should I pick?
Cookie-Editor if you are importing back into a browser extension. Puppeteer if you are feeding page.setCookie() or context.addCookies(). Key-value for a requests or axios session. Header if you just need something to paste after curl -H.
Why does a converted line start with #HttpOnly_?
That prefix is how curl marks an http-only cookie in a cookies.txt file. It looks like a comment so older readers skip it safely, while curl and yt-dlp understand it. Cookies without the flag are written normally.
What does an expiry of 0 or -1 mean?
Both mean a session cookie — one that dies when the browser closes. Netscape files write 0, Puppeteer writes -1, and Cookie-Editor drops the field and sets session to true instead. All three are read correctly and converted to whichever spelling the target format expects.
Why do the key-value and header formats lose data?
They only store names and values. Domain, path, expiry, secure and http-only have nowhere to go. That is fine when you are attaching cookies to a single request, but you cannot convert back to a complete cookies.txt afterwards without supplying a domain.
Can I take cookies from a browser extension into yt-dlp or curl?
Yes — that is the most common reason people land here. Export from Cookie-Editor as JSON, paste it in, and the output is a Netscape cookies.txt file you can pass to yt-dlp --cookies or curl -b.
claude cookie A cookie format converter that runs entirely in your browser.
Not affiliated with, endorsed by, or connected to Anthropic. Claude is a trademark of Anthropic PBC.
© 2026 claudecookie.com. All rights reserved.
