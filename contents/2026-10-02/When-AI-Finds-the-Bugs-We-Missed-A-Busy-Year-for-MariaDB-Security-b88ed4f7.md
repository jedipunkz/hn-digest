---
source: "https://mariadb.org/when-ai-finds-the-bugs-we-missed-a-very-busy-year-for-mariadb-security/"
hn_url: "https://news.ycombinator.com/item?id=49936918"
title: "When AI Finds the Bugs We Missed: A Busy Year for MariaDB Security"
article_title: "When AI Finds the Bugs We Missed: A Very Busy Year for MariaDB Security - MariaDB.org"
image: "https://mariadb.org/wp-content/uploads/2026/09/mariadb_security.png"
author: "mathnode"
captured_at: "2026-10-02T19:00:57Z"
capture_tool: "hn-digest"
hn_id: 49936918
score: 1
comments: 0
posted_at: "2026-10-02T18:31:55Z"
tags:
  - hacker-news
---

# When AI Finds the Bugs We Missed: A Busy Year for MariaDB Security

- HN: [49936918](https://news.ycombinator.com/item?id=49936918)
- Source: [mariadb.org](https://mariadb.org/when-ai-finds-the-bugs-we-missed-a-very-busy-year-for-mariadb-security/)
- Score: 1
- Comments: 0
- Posted: 2026-10-02T18:31:55Z

## Translation

Title: When AI Finds the Bugs We Missed: A Busy Year for MariaDB Security
Article title: When AI Finds the Bugs We Missed: A Very Busy Year for MariaDB Security - MariaDB.org
Description: AI-assisted security research uncovered old vulnerabilities in MariaDB Server. See how researchers found them, how we fixed them all, and why MariaDB is now more secure.

Article text:
When AI Finds the Bugs We Missed: A Very Busy Year for MariaDB Security - MariaDB.org
Skip to content
Download
Newsletter • Contact • Events • Planet
When AI Finds the Bugs We Missed: A Very Busy Year for MariaDB Security
2026-10-02 2026-10-02
Leave a comment on When AI Finds the Bugs We Missed: A Very Busy Year for MariaDB Security
You may have noticed that some recent MariaDB Server releases arrived a little later than expected. Sorry about that! But there is a good reason for it.
Over the last months, the MariaDB project has experienced something we had never seen before: a very large increase in security reports. And when I say large, look at this graph:
For years, we received a few security reports every quarter. Sometimes 3, sometimes 6, 11… the highest in the graph was 16.
And then suddenly, in one quarter, we received 81 reports .
But there is an important distinction to make here:
173 security reports does not mean 173 vulnerabilities.
Every report has to be investigated. Some turn into confirmed security vulnerabilities. Others reveal valid bugs or interesting edge cases but aren’t ultimately classified the same way.
Either way, somebody needs to reproduce the problem, understand it and determine what needs to be fixed.
This is an overview of the process:
We all know that AI is changing the way developers write code. Whether we like it or not, this is happening.
But there is another side of it: AI is also changing the way security researchers look for bugs.
Modern AI-assisted tools can analyze very large codebases, follow code paths, identify suspicious patterns, generate test cases and help security researchers investigate areas that would previously have required a significant amount of manual work.
MariaDB Server is a large and mature project, with code covering a huge number of features, protocols, storage engines, platforms and unusual combinations of them.
And suddenly, researchers equipped with new tools started looking at all of it in ways that simply weren’t practical before.
The result is what you see in the graph.
Of course, when we started receiving so many reports, this was a concern. If you have spent some time on GitHub recently, you probably know what I mean. AI can generate a very convincing bug report about a bug that simply doesn’t exist.
Nice explanation, nice code, nice stack trace… but nothing behind it. This is not what happened to us.
Almost all the reports we received were good reports. Real issues, with real work behind them. And the most important part: we fixed them all!
Our engineering teams reviewed the reports, reproduced the findings, evaluated their impact, developed fixes where needed, reviewed them, tested them, and often backported those fixes to several maintained branches.
For confirmed vulnerabilities, fixes have been included in the latest maintenance releases for the affected supported Community and Enterprise Server versions.
We have also started publishing our own MariaDB Security Advisories . This lets us communicate confirmed vulnerabilities and their fixes without waiting several weeks for the CVE assignment process to complete.
At the time of writing, 27 of those advisories are already public and fixed.
Connectors also have their dedicated Security Advisories [ 1 ][ 2 ][ 3 ].
Now multiply the investigation and engineering work by the unprecedented number of reports we received in 2026. This explains why some releases took longer than expected.
More Reports Does Not Mean MariaDB Suddenly Became Less Secure
This distinction is important. The graph shows reports received , not confirmed vulnerabilities. As you can see from the process, some issues aren’t related to the server (like connectors), and some are fixed but don’t lead to advisories and/or CVEs or are duplicates. And this explains the gap between 81 reports and 27 advisories.
The sudden increase therefore should not be read as a graph showing that MariaDB suddenly became less secure in 2026.
What changed dramatically was the amount of security research being performed against the MariaDB project (like many other open source projects), and the tools researchers have available to do it.
Issues and edge cases that were previously difficult or time-consuming to discover can now be explored much more efficiently.
And when a report identifies a real problem, that gives us the opportunity to fix it. That is exactly what has been happening. Every valid security issue reported to us has been investigated and fixed.
Fixing a security issue is not just changing a few lines of code and pushing to GitHub. As explained above, we need to reproduce it, understand it, check which versions are affected, fix it, review the fix, test it and, very often, backport it to several maintained versions, sometimes write a security advisory , assign the score, request the CVE id, …
Doing this for a few reports is normal. Doing it for 81 and then 92 reports is something else!
There is another graph I really wanted to share:
These are some of the security researchers who kept us busy 😉
At the top we have fg0x0 with an impressive 51 reports !
There is an important detail here: fg0x0 has been concentrating on the MariaDB Connectors , not MariaDB Server itself.
And this work has already resulted in concrete fixes. For example, one reported issue in Connector/Node.js concerned PAM authentication and ensuring credentials could not be sent over insecure transport . Similar work covered other connectors, including R2DBC .
On the Server side, we have letchu_pkt with 13 reports , vortfu with 12 , muhammaddaffa with 10 , followed by several other researchers.
vortfu , by the way, is from Automattic .
And some of these findings have been very significant. letchu_pkt , for example, reported several issues around Galera/wsrep handling , including problems involving values passed to shell commands. Those have since been fixed in the affected maintained releases.
So a BIG thank you to all of them!
And not only the people visible at the top of the graph. Thank you to everybody who spent time looking at MariaDB, found something suspicious and reported it responsibly.
This is how open source security should work.
There is also a lesson here that I find interesting. As developers, we naturally look carefully at new code. A new feature gets reviewed. It gets tests. People try it. People break it. But old code?
Well… it has been there forever, so it must be fine, right?… Not really 😉
Code being old only tells us that nobody found the problem before.
And this is maybe one of the very good uses of AI in software development. Instead of only generating more new code, we can also use it to go back and look again at the code we already have.
MariaDB is certainly not the only project where this can be useful. There is a lot of old code in large open source projects, and I expect security researchers will continue finding interesting things with these new tools. That’s good.
So, if you are a security researcher and you use fuzzing, static analysis, AI, your own scripts, or whatever new tool you have built: keep looking!
If you find something that looks like a security issue, report it through the appropriate private security channel rather than publishing an exploit immediately. And please provide a reproducible test case whenever possible. This makes an enormous difference.
To all the researchers who contributed during these very busy months: thank you again.
To our developers who suddenly had a mountain of security reports added to their normal work: a very big thank you as well.
And to our users who sometimes wondered why a release was taking a little longer than expected… Now you know. 😉
Posted by Frédéric Descamps 2026-10-02 2026-10-02 Posted in Community , Security
Post navigation
Your email address will not be published. Required fields are marked *
Copyright @ 2009 - 2026 MariaDB Foundation.

## Original Extract

AI-assisted security research uncovered old vulnerabilities in MariaDB Server. See how researchers found them, how we fixed them all, and why MariaDB is now more secure.

When AI Finds the Bugs We Missed: A Very Busy Year for MariaDB Security - MariaDB.org
Skip to content
Download
Newsletter • Contact • Events • Planet
When AI Finds the Bugs We Missed: A Very Busy Year for MariaDB Security
2026-10-02 2026-10-02
Leave a comment on When AI Finds the Bugs We Missed: A Very Busy Year for MariaDB Security
You may have noticed that some recent MariaDB Server releases arrived a little later than expected. Sorry about that! But there is a good reason for it.
Over the last months, the MariaDB project has experienced something we had never seen before: a very large increase in security reports. And when I say large, look at this graph:
For years, we received a few security reports every quarter. Sometimes 3, sometimes 6, 11… the highest in the graph was 16.
And then suddenly, in one quarter, we received 81 reports .
But there is an important distinction to make here:
173 security reports does not mean 173 vulnerabilities.
Every report has to be investigated. Some turn into confirmed security vulnerabilities. Others reveal valid bugs or interesting edge cases but aren’t ultimately classified the same way.
Either way, somebody needs to reproduce the problem, understand it and determine what needs to be fixed.
This is an overview of the process:
We all know that AI is changing the way developers write code. Whether we like it or not, this is happening.
But there is another side of it: AI is also changing the way security researchers look for bugs.
Modern AI-assisted tools can analyze very large codebases, follow code paths, identify suspicious patterns, generate test cases and help security researchers investigate areas that would previously have required a significant amount of manual work.
MariaDB Server is a large and mature project, with code covering a huge number of features, protocols, storage engines, platforms and unusual combinations of them.
And suddenly, researchers equipped with new tools started looking at all of it in ways that simply weren’t practical before.
The result is what you see in the graph.
Of course, when we started receiving so many reports, this was a concern. If you have spent some time on GitHub recently, you probably know what I mean. AI can generate a very convincing bug report about a bug that simply doesn’t exist.
Nice explanation, nice code, nice stack trace… but nothing behind it. This is not what happened to us.
Almost all the reports we received were good reports. Real issues, with real work behind them. And the most important part: we fixed them all!
Our engineering teams reviewed the reports, reproduced the findings, evaluated their impact, developed fixes where needed, reviewed them, tested them, and often backported those fixes to several maintained branches.
For confirmed vulnerabilities, fixes have been included in the latest maintenance releases for the affected supported Community and Enterprise Server versions.
We have also started publishing our own MariaDB Security Advisories . This lets us communicate confirmed vulnerabilities and their fixes without waiting several weeks for the CVE assignment process to complete.
At the time of writing, 27 of those advisories are already public and fixed.
Connectors also have their dedicated Security Advisories [ 1 ][ 2 ][ 3 ].
Now multiply the investigation and engineering work by the unprecedented number of reports we received in 2026. This explains why some releases took longer than expected.
More Reports Does Not Mean MariaDB Suddenly Became Less Secure
This distinction is important. The graph shows reports received , not confirmed vulnerabilities. As you can see from the process, some issues aren’t related to the server (like connectors), and some are fixed but don’t lead to advisories and/or CVEs or are duplicates. And this explains the gap between 81 reports and 27 advisories.
The sudden increase therefore should not be read as a graph showing that MariaDB suddenly became less secure in 2026.
What changed dramatically was the amount of security research being performed against the MariaDB project (like many other open source projects), and the tools researchers have available to do it.
Issues and edge cases that were previously difficult or time-consuming to discover can now be explored much more efficiently.
And when a report identifies a real problem, that gives us the opportunity to fix it. That is exactly what has been happening. Every valid security issue reported to us has been investigated and fixed.
Fixing a security issue is not just changing a few lines of code and pushing to GitHub. As explained above, we need to reproduce it, understand it, check which versions are affected, fix it, review the fix, test it and, very often, backport it to several maintained versions, sometimes write a security advisory , assign the score, request the CVE id, …
Doing this for a few reports is normal. Doing it for 81 and then 92 reports is something else!
There is another graph I really wanted to share:
These are some of the security researchers who kept us busy 😉
At the top we have fg0x0 with an impressive 51 reports !
There is an important detail here: fg0x0 has been concentrating on the MariaDB Connectors , not MariaDB Server itself.
And this work has already resulted in concrete fixes. For example, one reported issue in Connector/Node.js concerned PAM authentication and ensuring credentials could not be sent over insecure transport . Similar work covered other connectors, including R2DBC .
On the Server side, we have letchu_pkt with 13 reports , vortfu with 12 , muhammaddaffa with 10 , followed by several other researchers.
vortfu , by the way, is from Automattic .
And some of these findings have been very significant. letchu_pkt , for example, reported several issues around Galera/wsrep handling , including problems involving values passed to shell commands. Those have since been fixed in the affected maintained releases.
So a BIG thank you to all of them!
And not only the people visible at the top of the graph. Thank you to everybody who spent time looking at MariaDB, found something suspicious and reported it responsibly.
This is how open source security should work.
There is also a lesson here that I find interesting. As developers, we naturally look carefully at new code. A new feature gets reviewed. It gets tests. People try it. People break it. But old code?
Well… it has been there forever, so it must be fine, right?… Not really 😉
Code being old only tells us that nobody found the problem before.
And this is maybe one of the very good uses of AI in software development. Instead of only generating more new code, we can also use it to go back and look again at the code we already have.
MariaDB is certainly not the only project where this can be useful. There is a lot of old code in large open source projects, and I expect security researchers will continue finding interesting things with these new tools. That’s good.
So, if you are a security researcher and you use fuzzing, static analysis, AI, your own scripts, or whatever new tool you have built: keep looking!
If you find something that looks like a security issue, report it through the appropriate private security channel rather than publishing an exploit immediately. And please provide a reproducible test case whenever possible. This makes an enormous difference.
To all the researchers who contributed during these very busy months: thank you again.
To our developers who suddenly had a mountain of security reports added to their normal work: a very big thank you as well.
And to our users who sometimes wondered why a release was taking a little longer than expected… Now you know. 😉
Posted by Frédéric Descamps 2026-10-02 2026-10-02 Posted in Community , Security
Post navigation
Your email address will not be published. Required fields are marked *
Copyright @ 2009 - 2026 MariaDB Foundation.
