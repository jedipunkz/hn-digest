---
source: "https://lexontech.org/claude-writes-code-i-author-the-apps"
hn_url: "https://news.ycombinator.com/item?id=49968027"
title: "Claude writes code; I author the apps – Lex on Tech"
article_title: "Claude writes code; I author the apps — Lex on Tech"
image: "https://lexontech.org/img/og/157.png?1791213488"
author: "lexfri"
captured_at: "2026-10-05T17:59:20Z"
capture_tool: "hn-digest"
hn_id: 49968027
score: 1
comments: 0
posted_at: "2026-10-05T17:47:47Z"
tags:
  - hacker-news
---

# Claude writes code; I author the apps – Lex on Tech

- HN: [49968027](https://news.ycombinator.com/item?id=49968027)
- Source: [lexontech.org](https://lexontech.org/claude-writes-code-i-author-the-apps)
- Score: 1
- Comments: 0
- Posted: 2026-10-05T17:47:47Z

## Translation

Title: Claude writes code; I author the apps – Lex on Tech
Article title: Claude writes code; I author the apps — Lex on Tech
Description: Gotta say it’s a bit suspicious how quickly a one man team is “recreating” existing games, while being “a full time consultant”… That’s from a post I received on Mastodon after launching The Universe, Controlled, my homage to the sorely-missed early iPhone…

Article text:
Skip to content
Lex on Tech
Tech, mostly Apple, from Lex Friedman.
Membership Become a Lex on Tech member for bonus perks and to support the site. See the perks →
AI
• Opinion
Claude writes code; I author the apps
Gotta say it’s a bit suspicious how quickly a one man team is “recreating” existing games, while being “a full time consultant”…
That’s from a post I received on Mastodon after launching The Universe, Controlled , my homage to the sorely-missed early iPhone game Flight Control . I pressed to better understand exactly what the poster was suspicious of . “I wonder how much your LLM is copying the existing game,” the user wrote.
As I wrote recently, not everyone’s a fan of vibe-coded apps . “Vibe-coded,” in fact, is used as an insult, the coding equivalent of “ slop .” I think the term is problematic, but I understand some of the concern: Just as people are using AI tools to write crappy articles, generate crappy posters, and make crappy fake photos, there are folks using LLMs to create crappy websites, apps, and other software.
But there are good apps out there that are made at least in part with help from LLMs. In fact, I think it’s a safe bet that the vast majority of regularly-updated software you use today is made by developers leveraging AI tools in at least some ways.
That’s a hunch. Let me talk in real terms, using me and my actual apps. But let’s take a step back first.
I was fortunate growing up: My house had computers. I learned to code initially on a Commodore 64 and then a TRS-80. My parents hired a programming tutor for me from a local Radio Shack when they saw how much I was enjoying coding. I devoured programming books, copied code from magazines, and eventually went to a summer camp where I could study programming, too. (The camp also let me pursue my other two passions: close-up magic and movie making. I was a very cool kid.)
Like many, I started with BASIC and various permutations thereof, then studied Pascal, C, and C++. I released mediocre shareware Mac apps that were not good. I built websites that were fine.
I went to college and majored in Linguistics and Cognitive Science. I decided I didn’t want to major in computer science because I didn’t want to wrestle with computers for the rest of my life. (Hahahaha.)
After college, I moved to LA with my then-fiancée (now wife), got a job at a web-hosting company, and started teaching myself web development when I was finished with my customer support queue each day — first Perl, later PHP.
I got really, really good at web development. I was hired as the first full-time PHP developer at the parent company of MySpace. (I had a cubicle next to MySpace’s Tom for a little while.)
I co-founded a diet-tracking startup in the pre-smartphone era. I was the developer, product guy, and customer service guy in one. My other two cofounders held down other full-time jobs so that they could pay me. This was in the era where you still had to roll your own servers, pre Amazon Web Services. Our stack at a server farm in Seattle would crash at lunchtime when too many people logged meals at the same time. It was a nightmare.
We got acquired by a big Internet company, and I eventually left and went to work at Macworld , and in addition to writing articles, I’d occasionally write tools for the staff there to use for article assignments. At the time, Macworld was using bug-tracking software for article assignments, and I absolutely hated that, so I built a system for us to use instead.
A lot of my software development is inspired by hating existing apps and thus making something that helps others, but also selfishly benefits me.
I share all this to say, I have a healthy coding background.
I was a web app guy for years. Even when I started my career on the business side of podcasting, when I saw needs for tools, I developed them for myself, and then eventually product managed teams that would make the web-based tools we needed to run the ad sales side of businesses like Midroll/Stitcher, or generate web-based advertising rate cards at Wondery.
But I wasn’t making Mac apps or iOS apps. I wanted to. But it seemed daunting. I read a tutorial or two on using Xcode and SwiftUI and felt out of my depth.
After I launched Lex.Games by accident — a website for daily word games — I eventually felt like I really needed an app. I really wanted to build one. I still felt… daunted.
Then my pals Casey Liss and Ben Rice McCarthy both recommended I take a look at Hacking With Swift , an incredible resource from developer and impressive human Paul Hudson.
I patiently went through his “100 Days of Swift” online course. And by patiently, I mean, without a ton of patience; I went through it in about 30 days, and then started work on the Lex.Games iOS app.
(Quick aside, for both technical and non-technical readers: Lex.Games the app needs Lex.Games the website, to fetch puzzles, sync login states, record gameplays, manage friends, etc. That meant I needed to build an API — a means by which the app could talk to the website back and forth. I’d feel brilliant as I worked on the web side, where I felt like I knew what I was doing, generating JSON payloads for the app to consume. Then I’d feel like an idiot as my Xcode project choked over and over again trying to parse and display the data it got from the app. I’d feel like an idiot trying to make the user interface reasonable. I released many, many betas.)
When I built the Lex.Games app, Xcode offered no integration with AI tools like ChatGPT and Claude. When I got stuck, I’d of course search for answers online, and sometimes I would absolutely go to ChatGPT, share an error I was getting, and ask what it meant. Sometimes I’d paste in chunks of my code. And since I was trying to learn, I’d ask ChatGPT not just how to fix it, but to explain the fixes it was suggesting. They weren’t always good fixes, or even workable ones, but they helped me along.
This wasn’t even that long ago, but it feels like the Stone Age of LLM-assisted coding work: It was a slow, manual process. But I got smarter and released an app! More people play Lex.Games in the app each day than on the web, and I’m pleased by that. Many thousands of games are played across the Lex.Games ecosystem each day, and two-thirds of those plays are in the app.
I’m sharing all this mostly defensively, to say, hey, I’m a real developer. I can do this. I’m not a SwiftUI expert, at all , but I can and have built apps by hand. (After Lex.Games , I launched my Cryptograms app, also built by hand, so I earned that plural apps .)
But Xcode’s integration with LLM agents is a game changer.
The gate keepers and the door openers
I hear from “normal” people who don’t know exactly what’s meant by “agents” when it comes to AI tools. In this context, the easiest way to explain it is this: When I use Claude Code in Xcode today, it can see the full coding project; it can generate, delete, and edit code; and it can do a lot of this surprisingly fast.
Gotta say it’s a bit suspicious how quickly a one man team is “recreating” existing games, while being “a full time consultant”…
So back to this. I feel absolutely endowed with superpowers thanks to Claude Code. But no, I didn’t ask Claude Code to “knock off Strategery” or “recreate Flight Control.”
As I said earlier, it’s certainly possible to make a crappy app with instructions like that, without knowing or writing a lick of code. That’s not what I do.
For the record, I now own Strategery after striking a deal with the original developers. Flight Control wasn’t acquirable: EA bought it and then later shuttered it back in 2015; thus, I made The Universe, Controlled an homage instead.
With Strategery, I started by mapping out what I’d want the app to do, and how I wanted it to work. Because I knew I wanted iCloud sync — even before I was sure I wanted online gameplay, which I ended up also building — I thought about how I’d want to generate the maps that game uses, so that they could be seeded and recreated algorithmically. I thought about what rules I wanted the game to have, and how I wanted gameplay to work.
And I crafted an extremely detailed prompt for Claude Code. It was several thousand words.
I explained in detail what I wanted to create, how it would look, how the game would be played, and the functionality needed. I explained the algorithmic approach I wanted to map generation and sharing, and how I wanted things architected. And then I set Claude to work, telling it I didn’t want the entire app built right away, but rather a specific subset of the game’s functionality, on which we’d iterate.
I gave specific instructions about how I wanted the project itself shaped, because I wanted various areas of gameplay mechanics extrapolated out so that they’d be easier for me to modify directly.
When I added support for online gameplay, I pointed Claude Code to my Lex.Games API that I’d built by hand. I asked it to use a similar structure and approach — so that I could maintain and understand it — and told it precisely what I wanted, down to naming conventions and database structures.
Claude did a lot of the web-code grunt work, but it’s not as simple as saying “build an API for a game” (of which this instance of Claude had no knowledge). I had to map out the exact API calls needed, the return information needed, what the app would say back, and how the two should interact, authenticate, etc — and then set Claude to work implementing it.
Some developers maybe hate this. It’s different from autocomplete and code snippets. It’s different from going to sites like Stack Exchange and potentially copying and pasting code, and I understand that.
But believe me when I say that even though I’m asking Claude to handle the implementation, I promise you, I’m still doing development work.
My Mastodon writer asked if I just told Claude to knock off another app. I did not.
My initial prompt when I started on The Universe, Controlled was about 1,000 words — more than half the length of this article to this point. And it was explicitly only for a very limited subset of the game’s functionality. I’ve since prompted Claude with tens of thousands of words of additional prompts on the project.
I’m not trying to understate Claude’s work. I literally could not create The Universe, Controlled without Claude’s help. Claude is drawing graphics in code, which I’m atrociously bad at. Claude is implementing features that I’m not good enough to implement without its guidance.
What matters to me is that the app was built with humanity. It’s me. Hi. I’m the human.
I released 44 betas before the game hit the App Store, leveraging my own feedback and intuition on what the game needed — and of course, feedback from beta testers.
I started the game in landscape mode, the way Flight Control was initially built. But I prefer using my phone in portrait mode, and so wanted to come up with a way to make the game work well there, too. I could not just say to Claude, “make this work in portrait orientation”; that’s not a prompt that would yield good results. I promise.
I had to consider the problem and think about what it meant for gameplay. Especially since I wanted the game to handle rotation. Paths you trace for spacecraft in portrait mode must navigate to the same planets, which can’t have moved, when you rotate to landscape. That’s a complicated logic and design problem to think through. Crafting the solution is twofold: It’s figuring out what the hell the right answer is , and also implementing it. I tackled the first half, which I absolutely consider hard work to get right (and I’m pleased with how it works in the game now!) Claude handled the hard work of the implementation. I’m grateful that it did. That’s absolutely challenging work, and I’ll state it again plainly: I couldn’t have done it. I needed Claude’s help.
For more than a dozen betas, The Universe, Controlled had an extremely annoying issue: When you traced paths for spacecraft from the bo

[truncated]

## Original Extract

Gotta say it’s a bit suspicious how quickly a one man team is “recreating” existing games, while being “a full time consultant”… That’s from a post I received on Mastodon after launching The Universe, Controlled, my homage to the sorely-missed early iPhone…

Skip to content
Lex on Tech
Tech, mostly Apple, from Lex Friedman.
Membership Become a Lex on Tech member for bonus perks and to support the site. See the perks →
AI
• Opinion
Claude writes code; I author the apps
Gotta say it’s a bit suspicious how quickly a one man team is “recreating” existing games, while being “a full time consultant”…
That’s from a post I received on Mastodon after launching The Universe, Controlled , my homage to the sorely-missed early iPhone game Flight Control . I pressed to better understand exactly what the poster was suspicious of . “I wonder how much your LLM is copying the existing game,” the user wrote.
As I wrote recently, not everyone’s a fan of vibe-coded apps . “Vibe-coded,” in fact, is used as an insult, the coding equivalent of “ slop .” I think the term is problematic, but I understand some of the concern: Just as people are using AI tools to write crappy articles, generate crappy posters, and make crappy fake photos, there are folks using LLMs to create crappy websites, apps, and other software.
But there are good apps out there that are made at least in part with help from LLMs. In fact, I think it’s a safe bet that the vast majority of regularly-updated software you use today is made by developers leveraging AI tools in at least some ways.
That’s a hunch. Let me talk in real terms, using me and my actual apps. But let’s take a step back first.
I was fortunate growing up: My house had computers. I learned to code initially on a Commodore 64 and then a TRS-80. My parents hired a programming tutor for me from a local Radio Shack when they saw how much I was enjoying coding. I devoured programming books, copied code from magazines, and eventually went to a summer camp where I could study programming, too. (The camp also let me pursue my other two passions: close-up magic and movie making. I was a very cool kid.)
Like many, I started with BASIC and various permutations thereof, then studied Pascal, C, and C++. I released mediocre shareware Mac apps that were not good. I built websites that were fine.
I went to college and majored in Linguistics and Cognitive Science. I decided I didn’t want to major in computer science because I didn’t want to wrestle with computers for the rest of my life. (Hahahaha.)
After college, I moved to LA with my then-fiancée (now wife), got a job at a web-hosting company, and started teaching myself web development when I was finished with my customer support queue each day — first Perl, later PHP.
I got really, really good at web development. I was hired as the first full-time PHP developer at the parent company of MySpace. (I had a cubicle next to MySpace’s Tom for a little while.)
I co-founded a diet-tracking startup in the pre-smartphone era. I was the developer, product guy, and customer service guy in one. My other two cofounders held down other full-time jobs so that they could pay me. This was in the era where you still had to roll your own servers, pre Amazon Web Services. Our stack at a server farm in Seattle would crash at lunchtime when too many people logged meals at the same time. It was a nightmare.
We got acquired by a big Internet company, and I eventually left and went to work at Macworld , and in addition to writing articles, I’d occasionally write tools for the staff there to use for article assignments. At the time, Macworld was using bug-tracking software for article assignments, and I absolutely hated that, so I built a system for us to use instead.
A lot of my software development is inspired by hating existing apps and thus making something that helps others, but also selfishly benefits me.
I share all this to say, I have a healthy coding background.
I was a web app guy for years. Even when I started my career on the business side of podcasting, when I saw needs for tools, I developed them for myself, and then eventually product managed teams that would make the web-based tools we needed to run the ad sales side of businesses like Midroll/Stitcher, or generate web-based advertising rate cards at Wondery.
But I wasn’t making Mac apps or iOS apps. I wanted to. But it seemed daunting. I read a tutorial or two on using Xcode and SwiftUI and felt out of my depth.
After I launched Lex.Games by accident — a website for daily word games — I eventually felt like I really needed an app. I really wanted to build one. I still felt… daunted.
Then my pals Casey Liss and Ben Rice McCarthy both recommended I take a look at Hacking With Swift , an incredible resource from developer and impressive human Paul Hudson.
I patiently went through his “100 Days of Swift” online course. And by patiently, I mean, without a ton of patience; I went through it in about 30 days, and then started work on the Lex.Games iOS app.
(Quick aside, for both technical and non-technical readers: Lex.Games the app needs Lex.Games the website, to fetch puzzles, sync login states, record gameplays, manage friends, etc. That meant I needed to build an API — a means by which the app could talk to the website back and forth. I’d feel brilliant as I worked on the web side, where I felt like I knew what I was doing, generating JSON payloads for the app to consume. Then I’d feel like an idiot as my Xcode project choked over and over again trying to parse and display the data it got from the app. I’d feel like an idiot trying to make the user interface reasonable. I released many, many betas.)
When I built the Lex.Games app, Xcode offered no integration with AI tools like ChatGPT and Claude. When I got stuck, I’d of course search for answers online, and sometimes I would absolutely go to ChatGPT, share an error I was getting, and ask what it meant. Sometimes I’d paste in chunks of my code. And since I was trying to learn, I’d ask ChatGPT not just how to fix it, but to explain the fixes it was suggesting. They weren’t always good fixes, or even workable ones, but they helped me along.
This wasn’t even that long ago, but it feels like the Stone Age of LLM-assisted coding work: It was a slow, manual process. But I got smarter and released an app! More people play Lex.Games in the app each day than on the web, and I’m pleased by that. Many thousands of games are played across the Lex.Games ecosystem each day, and two-thirds of those plays are in the app.
I’m sharing all this mostly defensively, to say, hey, I’m a real developer. I can do this. I’m not a SwiftUI expert, at all , but I can and have built apps by hand. (After Lex.Games , I launched my Cryptograms app, also built by hand, so I earned that plural apps .)
But Xcode’s integration with LLM agents is a game changer.
The gate keepers and the door openers
I hear from “normal” people who don’t know exactly what’s meant by “agents” when it comes to AI tools. In this context, the easiest way to explain it is this: When I use Claude Code in Xcode today, it can see the full coding project; it can generate, delete, and edit code; and it can do a lot of this surprisingly fast.
Gotta say it’s a bit suspicious how quickly a one man team is “recreating” existing games, while being “a full time consultant”…
So back to this. I feel absolutely endowed with superpowers thanks to Claude Code. But no, I didn’t ask Claude Code to “knock off Strategery” or “recreate Flight Control.”
As I said earlier, it’s certainly possible to make a crappy app with instructions like that, without knowing or writing a lick of code. That’s not what I do.
For the record, I now own Strategery after striking a deal with the original developers. Flight Control wasn’t acquirable: EA bought it and then later shuttered it back in 2015; thus, I made The Universe, Controlled an homage instead.
With Strategery, I started by mapping out what I’d want the app to do, and how I wanted it to work. Because I knew I wanted iCloud sync — even before I was sure I wanted online gameplay, which I ended up also building — I thought about how I’d want to generate the maps that game uses, so that they could be seeded and recreated algorithmically. I thought about what rules I wanted the game to have, and how I wanted gameplay to work.
And I crafted an extremely detailed prompt for Claude Code. It was several thousand words.
I explained in detail what I wanted to create, how it would look, how the game would be played, and the functionality needed. I explained the algorithmic approach I wanted to map generation and sharing, and how I wanted things architected. And then I set Claude to work, telling it I didn’t want the entire app built right away, but rather a specific subset of the game’s functionality, on which we’d iterate.
I gave specific instructions about how I wanted the project itself shaped, because I wanted various areas of gameplay mechanics extrapolated out so that they’d be easier for me to modify directly.
When I added support for online gameplay, I pointed Claude Code to my Lex.Games API that I’d built by hand. I asked it to use a similar structure and approach — so that I could maintain and understand it — and told it precisely what I wanted, down to naming conventions and database structures.
Claude did a lot of the web-code grunt work, but it’s not as simple as saying “build an API for a game” (of which this instance of Claude had no knowledge). I had to map out the exact API calls needed, the return information needed, what the app would say back, and how the two should interact, authenticate, etc — and then set Claude to work implementing it.
Some developers maybe hate this. It’s different from autocomplete and code snippets. It’s different from going to sites like Stack Exchange and potentially copying and pasting code, and I understand that.
But believe me when I say that even though I’m asking Claude to handle the implementation, I promise you, I’m still doing development work.
My Mastodon writer asked if I just told Claude to knock off another app. I did not.
My initial prompt when I started on The Universe, Controlled was about 1,000 words — more than half the length of this article to this point. And it was explicitly only for a very limited subset of the game’s functionality. I’ve since prompted Claude with tens of thousands of words of additional prompts on the project.
I’m not trying to understate Claude’s work. I literally could not create The Universe, Controlled without Claude’s help. Claude is drawing graphics in code, which I’m atrociously bad at. Claude is implementing features that I’m not good enough to implement without its guidance.
What matters to me is that the app was built with humanity. It’s me. Hi. I’m the human.
I released 44 betas before the game hit the App Store, leveraging my own feedback and intuition on what the game needed — and of course, feedback from beta testers.
I started the game in landscape mode, the way Flight Control was initially built. But I prefer using my phone in portrait mode, and so wanted to come up with a way to make the game work well there, too. I could not just say to Claude, “make this work in portrait orientation”; that’s not a prompt that would yield good results. I promise.
I had to consider the problem and think about what it meant for gameplay. Especially since I wanted the game to handle rotation. Paths you trace for spacecraft in portrait mode must navigate to the same planets, which can’t have moved, when you rotate to landscape. That’s a complicated logic and design problem to think through. Crafting the solution is twofold: It’s figuring out what the hell the right answer is , and also implementing it. I tackled the first half, which I absolutely consider hard work to get right (and I’m pleased with how it works in the game now!) Claude handled the hard work of the implementation. I’m grateful that it did. That’s absolutely challenging work, and I’ll state it again plainly: I couldn’t have done it. I needed Claude’s help.
For more than a dozen betas, The Universe, Controlled had an extremely annoying issue: When you traced paths for spacecraft from the bo

[truncated]
