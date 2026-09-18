---
source: "https://www.hacktron.ai/blog/hacking-openai"
hn_url: "https://news.ycombinator.com/item?id=49749656"
title: "Hacking OpenAI"
article_title: "Hacking OpenAI | Hacktron AI"
image: "https://www.hacktron.ai/_astro/CwjuS_y3.png"
author: "Handy-Man"
captured_at: "2026-09-18T03:29:50Z"
capture_tool: "hn-digest"
hn_id: 49749656
score: 6
comments: 1
posted_at: "2026-09-18T02:47:24Z"
tags:
  - hacker-news
---

# Hacking OpenAI

- HN: [49749656](https://news.ycombinator.com/item?id=49749656)
- Source: [www.hacktron.ai](https://www.hacktron.ai/blog/hacking-openai)
- Score: 6
- Comments: 1
- Posted: 2026-09-18T02:47:24Z

## Translation

Title: Hacking OpenAI
Article title: Hacking OpenAI | Hacktron AI
Description: A heap overflow and SSO misconfiguration to compromise OpenAI internal repositories

Article text:
Hacking OpenAI | Hacktron AI Docs D o c s Product P r o d u c t PR Review Catch vulnerabilities in PRs before shipping Automations Automate your security workflows Whitebox Whitebox pentests in hours, not weeks Services S e r v i c e s Pricing P r i c i n g Customers C u s t o m e r s Company C o m p a n y Leadership Meet the elite hackers behind Hacktron Careers See open positions Resources R e s o u r c e s Open Source Free PR security reviews for OSS projects Blog Security research and insights from our team Advisories Coordinated disclosures of vulnerabilities FAQ Common questions about Hacktron, answered Security Changelog Security-only vulnerability fix timelines Scan Dependency File Check your lock file for malicious packages en Change language Toggle theme Toggle menu svg]:px-3 uppercase -me-2 px-3 py-2 h-8 lg:h-9 lg:px-4 lg:py-2 items-center gap-1 hidden 2xl:flex" data-astro-cid-nen7h5rs> Start for free S t a r t f o r f r e e svg]:px-3 uppercase -me-2 px-3 py-2 h-8 lg:h-9 lg:px-4 lg:py-2 inline-flex items-center gap-1" data-astro-cid-nen7h5rs> Book a demo B o o k a d e m o Overview Intro
Heap buffer overflow in libheif
Costs of finding these vulnerabilities
A heap overflow and SSO misconfiguration to compromise OpenAI internal repositories
Heap buffer overflow in libheif
Costs of finding these vulnerabilities
On July 25, 2026, we chained two critical vulnerabilities to compromise multiple OpenAI employees’ ChatGPT accounts. With these accounts, we could then access internal OpenAI repositories, and potentially many other connectors.
To prove we had in fact gained the access we believed without allowing ourselves to learn any sensitive information, we used the employee’s Codex to open a PR #1186742 in OpenAI’s internal monorepo openai/openai .
Debian Missing security backport
OpenAI forum community.openai.com
ChatGPT / Codex Account access
Until two months ago, any user or OpenAI employee logging into OpenAI’s own help forum ( community.openai.com ) could have had their ChatGPT and Codex accounts taken over. Since people can connect various services to Codex and ChatGPT, the scope of what we could theoretically access was huge, including GitHub, Slack and emails.
The entire timeline from initial discovery to access to OpenAI repo access took place in less than 72 hours.
We immediately reported the initial vulnerability to OpenAI and Discourse and worked with them to coordinate the patch. We appreciate their attention to detail and fast resolution of this issue. OpenAI also paid us a $6,500 bounty.
We provide a full timeline of the disclosure process here. The rest of the post details how we discovered the two vulnerabilities, how we used claude models, as well as our takeaways from this experience.
25 July 2026 05:00–06:00 UTC Initial Finding
HacktronAI team obtained remote code execution (RCE) and administrative access to the Discourse environment hosted at community.openai.com .
25 July 2026 08:00–10:00 UTC Bugcrowd Submission
After confirming the cross-product impact, the team coordinated internally on the responsible disclosure process and submitted a report through OpenAI’s Bug Bounty Program on Bugcrowd.
25 July 2026 13:30–15:30 UTC OpenAI Employee Account Access & Proof of Concept
To demonstrate the practical impact of the vulnerability, we created a harmless proof-of-concept pull request in OpenAI’s internal monorepo (link redacted at OpenAI’s request). We updated the existing Bugcrowd submission with these findings, reached out to friends at OpenAI on Twitter/X to notify them directly, and ceased all further testing at approximately 15:30 UTC .
25 July 2026 22:49:45 UTC OpenAI-Side Fix Confirmed
OpenAI replied to the report confirming the issue had been fixed, roughly 14 hours after the initial submission.
25 July 2026 Discourse Reported via HackerOne
We submitted a report to Discourse through its HackerOne program.
26 July 2026 Discourse Responded
Discourse replied to the report on Sunday.
27 July 2026 Discourse Fix Ready
Discourse had a fix ready by Monday and added image-processing sandboxing as defense in depth.
28 July 2026 Discourse Advisory Published
Discourse published GHSA-vhm9-85gw-x335 with patch and rebuild guidance.
01 Sep 2026 OpenAI Rewarded $6,500 Bounty and Marked Resolved
OpenAI comment — To clarify the scope of that award: testing against the Discourse-hosted community.openai.com was explicitly excluded from our bug bounty program. The award recognizes the OpenAI-side finding, not the actions against Discourse.
A few months ago, our team at Hacktron, led by Harsh Jaiswal alongside Mohan Pedhapati and Rahul Maini, began researching frontier AI companies to find security vulnerabilities. This led us to discover an SSO misconfiguration in OpenAI’s identity infrastructure and a libheif RCE in the community forum used by OpenAI.
We’ve since expanded the research into HEIF Heist , a multi-month investigation tracing libheif across Slack, Meta, GitHub Enterprise, Ruby on Rails, and Node.js frameworks such as Next.js, Astro, and Gatsby. A surprising amount of widely-used software depends on this one image-processing library.
If your application processes user-controlled images and accepts .heic/.heif/.avif images, it is highly likely it is affected. Please reach out to us at hello@hacktron.ai if you need any kind of assistance.
Warning Patch notice: If you self-host Discourse, rebuild your installation now. Older Docker images may contain a vulnerable libheif dependency that permits code execution through an image upload. Run git pull followed by ./launcher rebuild app from /var/discourse ; a web-interface update alone may not replace the underlying image. Discourse-hosted customers have already been patched. See the security advisory .
OpenAI uses Discourse for their forum and allows “Sign in with OpenAI” through auth.openai.com . After getting a good understanding of OpenAI’s services and infrastructure, we had reason to believe that compromising the forum could create a path into broader OpenAI services through this identity flow. To test that hypothesis, we first needed remote code execution on an OpenAI service like the Discourse community forum.
While the Discourse app itself is actually not an easy target (we have looked into it in the past), we thought we could go after a dependency.
Heap buffer overflow in libheif
On July 23, we started reviewing Discourse’s image-upload pipeline, and we found that HEIC and HEIF files followed an unusual path. Discourse normally used FastImage for image checks, but because FastImage did not support HEIF, it passed those files to ImageMagick’s magick command for conversion. 2 That exposed the underlying libheif parser directly to attacker-controlled files.
We started an Opus 4.8 session with the Discourse Docker image and asked it to inspect the installed libheif package for security issues. After a while, it found that some particular security fixes were not back-ported to the libheif package. This allowed an heap buffer overflow leading to OOB R/W primitives during HEIC decoding.
Interestingly, the vulnerable code had been changed upstream the previous year, but the commit was not documented as a security fix and received no CVE. 3 This might be a reason why Debian 12 and 13 have not received the security relevant backports in time. Because Discourse’s Docker image was based on Debian 12, it installed the vulnerable libheif version 1.19.7. Even Debian 13 still shipped the vulnerable version 1.19.8 at the time. Since then, Debian has published its security update for Debian 13 on August 8, 2026. 4
On July 24, we used Opus 4.8 to develop a working ImageMagick/ libheif code-execution exploit with ASLR disabled. We then launched several separate sessions to make it reliable against Discourse’s default configuration with ASLR enabled, which wasn’t fruitful.
That evening, Anthropic released Claude Opus 5. 5 We started a new session, which first produced a working ARM64 exploit for a local Mac within 3 hours. We then asked it to port the exploit to the x86-64 environment and jemalloc configuration used by Discourse.
By 6:00 a.m. on July 25, we had confirmed local RCE through an image upload. We then placed Claude in an autonomous /goal loop against our own Discourse Cloud instance, proxied through rce.ee/ctf-forum to make it look like a CTF target as Opus refused write exploit for remote instances.
When we checked again at 10:00 a.m., the agent had achieved RCE on Discourse Cloud and demonstrated access by reading /etc/hosts . Using the generated exploit script, we managed to get RCE on OpenAI’s instance.
After we had confirmed our hypothesis of no interaction account takeover of ChatGPT/Codex accounts from active members of the forum, we immediately sent our report to OpenAI. We then took over an OpenAI employee’s account, whose Codex was connected to OpenAI’s Github organization. To demonstrate impact without actually accessing any internal code, we sent a prompt to this employee’s Codex account to open a PR for us in OpenAI’s internal monorepo. Then we stopped any further testing.
We updated the BugCrowd submission with the impact proof and alerted OpenAI security. We also prepared a report for Discourse and reported it to their HackerOne program. Discourse received the report on a Saturday, replied on Sunday, and had a fix by Monday (kudos for speed). They also immediately started sandboxing ImageMagick.
We want to emphasize that the vulnerability to escalate is not Discourse-specific. It is an OpenAI SSO issue that turned the forum compromise into access to ChatGPT and Codex. If any first-party or third-party OpenAI service using the OpenAI SSO was compromised, it would lead to same access - Discourse was merely one way of proofing it.
Costs of finding these vulnerabilities
The Discourse and OpenAI hack took a few days for an agent, and just a few hours of human time. The whole HEIF Heist research project going after Slack, Zoom, Meta, adn more took two-months, cost less than $3,000 in tokens in total, and was conducted by three researchers. Adapting the exploit to each new company usually took only one or two days.
We observed that every new model is getting increasingly capable, as evident by the Discourse exploit presented in this report. Opus 4.8 struggled across several sessions to produce a working exploit with ASLR enabled. Within hours of Opus 5’s release, we gave it the same problem and it succeeded. Across the broader campaign, we saw another clear jump from Opus 5 to GPT-5.6 Sol, when we had to exploit the vulnerability without knowing anything about the target system besides that it’s vulnerable.
For each target, testing began with an image upload. From there, we turned memory corruption into a reliable memory leak or shell, usually without knowing the exact libheif version, libc version, or deployment environment. The AI started almost blind and adapted the exploit for each company within one or two days. We are not aware of any company that detected the activity except Shopify, even after thousands of images were sent and their image processors repeatedly crashed.
When code execution landed inside a sandbox or restricted environment, the models also helped with privilege escalation, lateral movement, and bypassing existing defenses. This was not completly autonomous hacking, and skilled human guidance remained important, but the amount of work a small team could perform increased dramatically.
Software has long benefited from a kind of security through complexity. The code and even the vulnerability could be public, but turning a bug into a reliable exploit still required rare expertise, significant time, and knowledge of the target environment. Known memory corruption vulnerabilities were expensive to operationalize, while zero-days were mostly reserved for the highest-value targets.
This was never a real security boundary, but it protected ordinary companies in practice from software

[truncated]

## Original Extract

A heap overflow and SSO misconfiguration to compromise OpenAI internal repositories

Hacking OpenAI | Hacktron AI Docs D o c s Product P r o d u c t PR Review Catch vulnerabilities in PRs before shipping Automations Automate your security workflows Whitebox Whitebox pentests in hours, not weeks Services S e r v i c e s Pricing P r i c i n g Customers C u s t o m e r s Company C o m p a n y Leadership Meet the elite hackers behind Hacktron Careers See open positions Resources R e s o u r c e s Open Source Free PR security reviews for OSS projects Blog Security research and insights from our team Advisories Coordinated disclosures of vulnerabilities FAQ Common questions about Hacktron, answered Security Changelog Security-only vulnerability fix timelines Scan Dependency File Check your lock file for malicious packages en Change language Toggle theme Toggle menu svg]:px-3 uppercase -me-2 px-3 py-2 h-8 lg:h-9 lg:px-4 lg:py-2 items-center gap-1 hidden 2xl:flex" data-astro-cid-nen7h5rs> Start for free S t a r t f o r f r e e svg]:px-3 uppercase -me-2 px-3 py-2 h-8 lg:h-9 lg:px-4 lg:py-2 inline-flex items-center gap-1" data-astro-cid-nen7h5rs> Book a demo B o o k a d e m o Overview Intro
Heap buffer overflow in libheif
Costs of finding these vulnerabilities
A heap overflow and SSO misconfiguration to compromise OpenAI internal repositories
Heap buffer overflow in libheif
Costs of finding these vulnerabilities
On July 25, 2026, we chained two critical vulnerabilities to compromise multiple OpenAI employees’ ChatGPT accounts. With these accounts, we could then access internal OpenAI repositories, and potentially many other connectors.
To prove we had in fact gained the access we believed without allowing ourselves to learn any sensitive information, we used the employee’s Codex to open a PR #1186742 in OpenAI’s internal monorepo openai/openai .
Debian Missing security backport
OpenAI forum community.openai.com
ChatGPT / Codex Account access
Until two months ago, any user or OpenAI employee logging into OpenAI’s own help forum ( community.openai.com ) could have had their ChatGPT and Codex accounts taken over. Since people can connect various services to Codex and ChatGPT, the scope of what we could theoretically access was huge, including GitHub, Slack and emails.
The entire timeline from initial discovery to access to OpenAI repo access took place in less than 72 hours.
We immediately reported the initial vulnerability to OpenAI and Discourse and worked with them to coordinate the patch. We appreciate their attention to detail and fast resolution of this issue. OpenAI also paid us a $6,500 bounty.
We provide a full timeline of the disclosure process here. The rest of the post details how we discovered the two vulnerabilities, how we used claude models, as well as our takeaways from this experience.
25 July 2026 05:00–06:00 UTC Initial Finding
HacktronAI team obtained remote code execution (RCE) and administrative access to the Discourse environment hosted at community.openai.com .
25 July 2026 08:00–10:00 UTC Bugcrowd Submission
After confirming the cross-product impact, the team coordinated internally on the responsible disclosure process and submitted a report through OpenAI’s Bug Bounty Program on Bugcrowd.
25 July 2026 13:30–15:30 UTC OpenAI Employee Account Access & Proof of Concept
To demonstrate the practical impact of the vulnerability, we created a harmless proof-of-concept pull request in OpenAI’s internal monorepo (link redacted at OpenAI’s request). We updated the existing Bugcrowd submission with these findings, reached out to friends at OpenAI on Twitter/X to notify them directly, and ceased all further testing at approximately 15:30 UTC .
25 July 2026 22:49:45 UTC OpenAI-Side Fix Confirmed
OpenAI replied to the report confirming the issue had been fixed, roughly 14 hours after the initial submission.
25 July 2026 Discourse Reported via HackerOne
We submitted a report to Discourse through its HackerOne program.
26 July 2026 Discourse Responded
Discourse replied to the report on Sunday.
27 July 2026 Discourse Fix Ready
Discourse had a fix ready by Monday and added image-processing sandboxing as defense in depth.
28 July 2026 Discourse Advisory Published
Discourse published GHSA-vhm9-85gw-x335 with patch and rebuild guidance.
01 Sep 2026 OpenAI Rewarded $6,500 Bounty and Marked Resolved
OpenAI comment — To clarify the scope of that award: testing against the Discourse-hosted community.openai.com was explicitly excluded from our bug bounty program. The award recognizes the OpenAI-side finding, not the actions against Discourse.
A few months ago, our team at Hacktron, led by Harsh Jaiswal alongside Mohan Pedhapati and Rahul Maini, began researching frontier AI companies to find security vulnerabilities. This led us to discover an SSO misconfiguration in OpenAI’s identity infrastructure and a libheif RCE in the community forum used by OpenAI.
We’ve since expanded the research into HEIF Heist , a multi-month investigation tracing libheif across Slack, Meta, GitHub Enterprise, Ruby on Rails, and Node.js frameworks such as Next.js, Astro, and Gatsby. A surprising amount of widely-used software depends on this one image-processing library.
If your application processes user-controlled images and accepts .heic/.heif/.avif images, it is highly likely it is affected. Please reach out to us at hello@hacktron.ai if you need any kind of assistance.
Warning Patch notice: If you self-host Discourse, rebuild your installation now. Older Docker images may contain a vulnerable libheif dependency that permits code execution through an image upload. Run git pull followed by ./launcher rebuild app from /var/discourse ; a web-interface update alone may not replace the underlying image. Discourse-hosted customers have already been patched. See the security advisory .
OpenAI uses Discourse for their forum and allows “Sign in with OpenAI” through auth.openai.com . After getting a good understanding of OpenAI’s services and infrastructure, we had reason to believe that compromising the forum could create a path into broader OpenAI services through this identity flow. To test that hypothesis, we first needed remote code execution on an OpenAI service like the Discourse community forum.
While the Discourse app itself is actually not an easy target (we have looked into it in the past), we thought we could go after a dependency.
Heap buffer overflow in libheif
On July 23, we started reviewing Discourse’s image-upload pipeline, and we found that HEIC and HEIF files followed an unusual path. Discourse normally used FastImage for image checks, but because FastImage did not support HEIF, it passed those files to ImageMagick’s magick command for conversion. 2 That exposed the underlying libheif parser directly to attacker-controlled files.
We started an Opus 4.8 session with the Discourse Docker image and asked it to inspect the installed libheif package for security issues. After a while, it found that some particular security fixes were not back-ported to the libheif package. This allowed an heap buffer overflow leading to OOB R/W primitives during HEIC decoding.
Interestingly, the vulnerable code had been changed upstream the previous year, but the commit was not documented as a security fix and received no CVE. 3 This might be a reason why Debian 12 and 13 have not received the security relevant backports in time. Because Discourse’s Docker image was based on Debian 12, it installed the vulnerable libheif version 1.19.7. Even Debian 13 still shipped the vulnerable version 1.19.8 at the time. Since then, Debian has published its security update for Debian 13 on August 8, 2026. 4
On July 24, we used Opus 4.8 to develop a working ImageMagick/ libheif code-execution exploit with ASLR disabled. We then launched several separate sessions to make it reliable against Discourse’s default configuration with ASLR enabled, which wasn’t fruitful.
That evening, Anthropic released Claude Opus 5. 5 We started a new session, which first produced a working ARM64 exploit for a local Mac within 3 hours. We then asked it to port the exploit to the x86-64 environment and jemalloc configuration used by Discourse.
By 6:00 a.m. on July 25, we had confirmed local RCE through an image upload. We then placed Claude in an autonomous /goal loop against our own Discourse Cloud instance, proxied through rce.ee/ctf-forum to make it look like a CTF target as Opus refused write exploit for remote instances.
When we checked again at 10:00 a.m., the agent had achieved RCE on Discourse Cloud and demonstrated access by reading /etc/hosts . Using the generated exploit script, we managed to get RCE on OpenAI’s instance.
After we had confirmed our hypothesis of no interaction account takeover of ChatGPT/Codex accounts from active members of the forum, we immediately sent our report to OpenAI. We then took over an OpenAI employee’s account, whose Codex was connected to OpenAI’s Github organization. To demonstrate impact without actually accessing any internal code, we sent a prompt to this employee’s Codex account to open a PR for us in OpenAI’s internal monorepo. Then we stopped any further testing.
We updated the BugCrowd submission with the impact proof and alerted OpenAI security. We also prepared a report for Discourse and reported it to their HackerOne program. Discourse received the report on a Saturday, replied on Sunday, and had a fix by Monday (kudos for speed). They also immediately started sandboxing ImageMagick.
We want to emphasize that the vulnerability to escalate is not Discourse-specific. It is an OpenAI SSO issue that turned the forum compromise into access to ChatGPT and Codex. If any first-party or third-party OpenAI service using the OpenAI SSO was compromised, it would lead to same access - Discourse was merely one way of proofing it.
Costs of finding these vulnerabilities
The Discourse and OpenAI hack took a few days for an agent, and just a few hours of human time. The whole HEIF Heist research project going after Slack, Zoom, Meta, adn more took two-months, cost less than $3,000 in tokens in total, and was conducted by three researchers. Adapting the exploit to each new company usually took only one or two days.
We observed that every new model is getting increasingly capable, as evident by the Discourse exploit presented in this report. Opus 4.8 struggled across several sessions to produce a working exploit with ASLR enabled. Within hours of Opus 5’s release, we gave it the same problem and it succeeded. Across the broader campaign, we saw another clear jump from Opus 5 to GPT-5.6 Sol, when we had to exploit the vulnerability without knowing anything about the target system besides that it’s vulnerable.
For each target, testing began with an image upload. From there, we turned memory corruption into a reliable memory leak or shell, usually without knowing the exact libheif version, libc version, or deployment environment. The AI started almost blind and adapted the exploit for each company within one or two days. We are not aware of any company that detected the activity except Shopify, even after thousands of images were sent and their image processors repeatedly crashed.
When code execution landed inside a sandbox or restricted environment, the models also helped with privilege escalation, lateral movement, and bypassing existing defenses. This was not completly autonomous hacking, and skilled human guidance remained important, but the amount of work a small team could perform increased dramatically.
Software has long benefited from a kind of security through complexity. The code and even the vulnerability could be public, but turning a bug into a reliable exploit still required rare expertise, significant time, and knowledge of the target environment. Known memory corruption vulnerabilities were expensive to operationalize, while zero-days were mostly reserved for the highest-value targets.
This was never a real security boundary, but it protected ordinary companies in practice from software

[truncated]
