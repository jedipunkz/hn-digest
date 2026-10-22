---
source: "https://shopify.engineering/remote-dev"
hn_url: "https://news.ycombinator.com/item?id=50198812"
title: "We gave AI agents a mobile dev environment"
article_title: "We gave AI agents a mobile dev environment (2026) - Shopify"
image: "https://cdn.shopify.com/b/shopify-brochure2-assets/2da4b000319a7dd77e86bd9171793eea.jpg"
author: "dayanruben"
captured_at: "2026-10-22T18:20:44Z"
capture_tool: "hn-digest"
hn_id: 50198812
score: 3
comments: 0
posted_at: "2026-10-22T17:52:31Z"
tags:
  - hacker-news
---

# We gave AI agents a mobile dev environment

- HN: [50198812](https://news.ycombinator.com/item?id=50198812)
- Source: [shopify.engineering](https://shopify.engineering/remote-dev)
- Score: 3
- Comments: 0
- Posted: 2026-10-22T17:52:31Z

## Translation

Title: We gave AI agents a mobile dev environment
Article title: We gave AI agents a mobile dev environment (2026) - Shopify
Description: We built Remote Dev so AI agents can build and test apps in parallel, without maxing out developers’ laptops.

Article text:
Skip to Content Solutions Start Start your business . Build your brand
Create your website . Online store editor
Customize your store . Store themes
Find business apps . Shopify app store
Own your site domain . Domains & hosting
Explore free business tools . Tools to run your business
Sell Sell your products . Sell online or in person
Check out customers . World-class checkout
Sell online . Grow your business online
Sell across channels . Reach millions of shoppers and boost sales
Sell globally . International sales
Sell wholesale & direct . Business-to-business (B2B)
Market Market your business . Reach & retain customers
Market across social . Social media integrations
Chat with customers . Shopify Inbox
Nurture customers . Shopify Messaging
Know your audience . Gain customer insights
Manage Manage your business . Track sales, orders & analytics
Measure your performance . Analytics and Reporting
Manage your stock & orders . Inventory & order management
Automate your business . Shopify Flow
Shopify Developers . Build with Shopify's powerful APIs
Plus . A commerce solution for growing digital brands
All Products . Explore all Shopify products & features
Resources Help and support Help and support . Get 24/7 support
Business courses . Learn from proven experts
Popular topics What is Shopify? . How our commerce platform works
Essential tools Business name generator .
What’s new Changelog . Your source for recent updates
Newsroom . All company news and press releases
blog | Mobile We gave AI agents a mobile dev environment
We built Remote Dev so AI agents can build and test apps in parallel, without maxing out developers’ laptops.
We built agentic workflows to tackle mobile development tasks in parallel, but quickly hit the limits of our development machines. Just three or four workflows were enough to max out a laptop, with compilers, simulators, and emulators competing for resources. And we had to keep it awake. If someone accidentally closed a laptop lid, the work stopped.
Our web and backend teams were already using River (our AI agent in Slack) to work without tying up their development machines. We needed the same thing for mobile. But River ran on Linux: it could edit Swift, but building and testing the iOS app still meant going back to the same laptops we’d already maxed out.
Our ability to delegate work had outgrown our ability to run it. So we built Remote Dev.
Two constraints: too many builds sharing one Mac, or missing iOS tooling on Linux.
Multiple agents, multiple tasks
With Remote Dev, we're no longer limited by how many builds and simulators our laptops can handle. Each River session gets its own environment, so we can have one agent debugging something on Android, another working on an iOS feature, and another testing a PR, all in parallel.
And since everything runs remotely, we can start a different task on each available environment, directly from Slack, without having to run Xcode or Android Studio locally!
One of the biggest advantages of Remote Dev is how much of the development workflow River can handle autonomously. It can implement changes, build the app, run tests, interact with the simulators, collect screenshots, and create a PR. If something fails, it can inspect the logs, fix the issue, and try again, all while we focus on other tasks.
We can delegate things like rebasing branches, resolving merge conflicts, checking CI failures, addressing PR feedback, triaging incoming crashes, and reviewing code. All those small tasks that used to interrupt our day can now be handled by River.
We can take this even further with automated workflows that run without anyone having to start a task manually. Some of these run overnight, automatically picking up bugs, implementing fixes, testing them, and creating PRs. So we can start the next morning with fixes already implemented and ready for review.
One of the biggest challenges with AI-generated code is knowing whether it actually works. Just because the code compiles doesn't mean the feature behaves as expected. With Remote Dev, we can go beyond builds and tests. Agents can navigate through the app, interact with elements, and capture screenshots or recordings of the result on both iOS and Android.
For bug fixes, River can reproduce the issue, record the broken behavior, implement the fix, and then repeat the same steps to show the before and after. This makes reviewing changes much easier. Instead of just looking at the code, we can actually see what changed and verify the result.
And recently, we took this even further. We added the ability for River to generate explainer videos from these recordings, showing the app running inside 3D phone renders, with animations and narration. All automatically generated, without anyone having to manually edit a video. So now River can not only implement and verify a fix, but also generate a video explaining what it did. Pretty cool!
Anyone can now contribute to mobile
One of the things we're most excited about is making mobile development accessible to people who don't have a mobile environment configured.
With River and Remote Dev, a web developer can fix a bug, a designer can iterate on a UI, and a PM can test a feature, all through Slack.
But there was still one thing missing. River could interact with the app and send screenshots, but what if you wanted to try it yourself? That's why we built Live Preview.
We stream the iOS Simulator and Android Emulator directly from River’s Remote Dev environment to a Quick website over WebSockets. The browser displays the live device stream and sends interactions back to the remote simulator or emulator, so anyone can open a link and use the actual running app in real time.
Every Remote Dev session automatically gets its own Live Preview for both platforms. Sessions are shareable, so multiple people can watch and collaborate on the same running app. We've also added support for deep linking and live streaming of the mobile CLI, making it easier to jump to specific screens and inspect what's happening under the hood.
No local setup, no builds to install, not even a Mac required.
This opens up mobile development to so many more people. Instead of sending screenshots back and forth, we can just share a link and say, "Try it yourself!"
We set up a fleet of remote Macs with our development environment and iOS simulators. Agents connect to these machines over SSH, but getting Android to work was trickier. Android’s emulator uses hardware virtualization, so running it inside a cloud VM requires nested virtualization, which macOS does not support out of the box. We got around this limitation by connecting to an Android emulator running on a separate Linux VM with nested virtualization enabled. With this setup, agents can build for both iOS and Android on a Mac and test on either platform, all within the same SSH session.
Each agent claims a ready Mac with an attached Android emulator. We replace Macs after use and periodically refresh idle ones.
We then tackled setup times. Initial app compilation and main checkout were loooooooong, taking over 30 minutes. So, instead of preparing a Mac after a request arrives, we switched to keeping a pool of warm Macs at the ready: main checked out (no small feat, as our monorepo is massive), tools installed, apps built, and caches warm. An agent claims a Mac, synchronizes its branch, and spins up a prebuilt Android VM instance to attach. The agent continues working in its own checkout, whether it’s running locally or in the cloud, and syncs the changes to the Mac. All scripts and builds then run on the Mac. Doing this, the time from requesting a Mac to inspecting code changes on a running device went down to just a few minutes.
Some background magic also had to be built to corral our pool. We destroy the Mac when the session ends and prepare a replacement. We also periodically replace idle Macs, pulling fresh main and rebuilding the apps and caches. We stagger this rotation so the pool stays available without interrupting active sessions.
We got the infra running, but River was still blind to it. It always stated that an iOS build cannot run on Linux, so it could code but not verify the results. So we created a skill to teach agents how to test their changes and prove their work. It required several iterations until that symbiotic relationship with two machines was seamless. It didn't start this way; even with the great tooling in place, River kept forgetting about the remote Mac, being silent while the communication with the remote Mac was happening, or taking a long time to come back to the user.
Iteration after iteration, we kept improving the skill, with the appropriate tooling in place: a CLI that encapsulates the complexity of the SSH flags. And instead of sleeping or setting timeouts for the Mac to finish long-running tasks, we used tmux panes that River could observe in real time.
Changing the way mobile developers work today
Even with all the constraints we had with our current laptops’ limitations, our teams were used to working locally for mobile development: having a real simulator to interact with, build the app, and test it ourselves. This was not only inefficient because of the laptop constraints, but also not scalable for the large migration task we had on our hands. Moving people to use River instead was not only a technological problem, it was also a psychological one. We had to change a habit.
Some people had tried early versions of River and the experience wasn’t the best: builds not working, too slow, and not being useful for day-to-day work. "This is not for me, I'll go back to coding locally" was the thought. We had to make it great for people to start using it more and more.
But as we all know, it’s not only build it and they'll come. Even internally. You need to show them. You need to evangelize about how good the tool is. And repeat it over and over. That’s why we used storytelling and internal ads to make people aware that River was now great on mobile.
We got the idea of converting River into a real-world person, frustrated about working on mobile so people could relate to the feeling. But then this real-world River gets a real Mac at home, enabling her to build mobile and, suddenly, all her problems are solved. People don't see all the complexity of the solution—the one we are explaining in this post—they only saw that now everything worked magically. Once we had the River character, it made the next internal campaigns easier. This is an example of one of those, all obviously AI-generated, making a parody of the Mac vs PC ad. This added not only storytelling and fun to the equation, but also nostalgia for the people that watched the original ad:
This ended up increasing River adoption. River initially created only 6% of merged mobile pull requests. This is how the number evolved over time, reversing the trend of local development and evolving to 61% of merged mobile pull requests:
Building in public, multiplayer and self-improving loops
One of the great things about using River is that it's multiplayer by default. Earlier this year, Tobi shared how this was already shifting Shopify’s culture and letting people learn from others’ way of using agents.
In our case, we've used this multiplayer feature not only to learn how others use River to code on mobile, but also to improve the experience for newer threads in the future. We can see where people get stuck, where guidance goes wrong, where access to tools gets rejected. All of it.
For this we used Jev , a really fast, really cheap classifier model. It observes all River conversations on mobile, tags them with multiple questions, and answers each message. We categorize each conversation based on how good the outcome was, and use the failures as input for an LLM to find improvements.
We run this iteratively, using River, so River can improve itself for future usage.
What started as a way to give River access to a mobile devel

[truncated]

## Original Extract

We built Remote Dev so AI agents can build and test apps in parallel, without maxing out developers’ laptops.

Skip to Content Solutions Start Start your business . Build your brand
Create your website . Online store editor
Customize your store . Store themes
Find business apps . Shopify app store
Own your site domain . Domains & hosting
Explore free business tools . Tools to run your business
Sell Sell your products . Sell online or in person
Check out customers . World-class checkout
Sell online . Grow your business online
Sell across channels . Reach millions of shoppers and boost sales
Sell globally . International sales
Sell wholesale & direct . Business-to-business (B2B)
Market Market your business . Reach & retain customers
Market across social . Social media integrations
Chat with customers . Shopify Inbox
Nurture customers . Shopify Messaging
Know your audience . Gain customer insights
Manage Manage your business . Track sales, orders & analytics
Measure your performance . Analytics and Reporting
Manage your stock & orders . Inventory & order management
Automate your business . Shopify Flow
Shopify Developers . Build with Shopify's powerful APIs
Plus . A commerce solution for growing digital brands
All Products . Explore all Shopify products & features
Resources Help and support Help and support . Get 24/7 support
Business courses . Learn from proven experts
Popular topics What is Shopify? . How our commerce platform works
Essential tools Business name generator .
What’s new Changelog . Your source for recent updates
Newsroom . All company news and press releases
blog | Mobile We gave AI agents a mobile dev environment
We built Remote Dev so AI agents can build and test apps in parallel, without maxing out developers’ laptops.
We built agentic workflows to tackle mobile development tasks in parallel, but quickly hit the limits of our development machines. Just three or four workflows were enough to max out a laptop, with compilers, simulators, and emulators competing for resources. And we had to keep it awake. If someone accidentally closed a laptop lid, the work stopped.
Our web and backend teams were already using River (our AI agent in Slack) to work without tying up their development machines. We needed the same thing for mobile. But River ran on Linux: it could edit Swift, but building and testing the iOS app still meant going back to the same laptops we’d already maxed out.
Our ability to delegate work had outgrown our ability to run it. So we built Remote Dev.
Two constraints: too many builds sharing one Mac, or missing iOS tooling on Linux.
Multiple agents, multiple tasks
With Remote Dev, we're no longer limited by how many builds and simulators our laptops can handle. Each River session gets its own environment, so we can have one agent debugging something on Android, another working on an iOS feature, and another testing a PR, all in parallel.
And since everything runs remotely, we can start a different task on each available environment, directly from Slack, without having to run Xcode or Android Studio locally!
One of the biggest advantages of Remote Dev is how much of the development workflow River can handle autonomously. It can implement changes, build the app, run tests, interact with the simulators, collect screenshots, and create a PR. If something fails, it can inspect the logs, fix the issue, and try again, all while we focus on other tasks.
We can delegate things like rebasing branches, resolving merge conflicts, checking CI failures, addressing PR feedback, triaging incoming crashes, and reviewing code. All those small tasks that used to interrupt our day can now be handled by River.
We can take this even further with automated workflows that run without anyone having to start a task manually. Some of these run overnight, automatically picking up bugs, implementing fixes, testing them, and creating PRs. So we can start the next morning with fixes already implemented and ready for review.
One of the biggest challenges with AI-generated code is knowing whether it actually works. Just because the code compiles doesn't mean the feature behaves as expected. With Remote Dev, we can go beyond builds and tests. Agents can navigate through the app, interact with elements, and capture screenshots or recordings of the result on both iOS and Android.
For bug fixes, River can reproduce the issue, record the broken behavior, implement the fix, and then repeat the same steps to show the before and after. This makes reviewing changes much easier. Instead of just looking at the code, we can actually see what changed and verify the result.
And recently, we took this even further. We added the ability for River to generate explainer videos from these recordings, showing the app running inside 3D phone renders, with animations and narration. All automatically generated, without anyone having to manually edit a video. So now River can not only implement and verify a fix, but also generate a video explaining what it did. Pretty cool!
Anyone can now contribute to mobile
One of the things we're most excited about is making mobile development accessible to people who don't have a mobile environment configured.
With River and Remote Dev, a web developer can fix a bug, a designer can iterate on a UI, and a PM can test a feature, all through Slack.
But there was still one thing missing. River could interact with the app and send screenshots, but what if you wanted to try it yourself? That's why we built Live Preview.
We stream the iOS Simulator and Android Emulator directly from River’s Remote Dev environment to a Quick website over WebSockets. The browser displays the live device stream and sends interactions back to the remote simulator or emulator, so anyone can open a link and use the actual running app in real time.
Every Remote Dev session automatically gets its own Live Preview for both platforms. Sessions are shareable, so multiple people can watch and collaborate on the same running app. We've also added support for deep linking and live streaming of the mobile CLI, making it easier to jump to specific screens and inspect what's happening under the hood.
No local setup, no builds to install, not even a Mac required.
This opens up mobile development to so many more people. Instead of sending screenshots back and forth, we can just share a link and say, "Try it yourself!"
We set up a fleet of remote Macs with our development environment and iOS simulators. Agents connect to these machines over SSH, but getting Android to work was trickier. Android’s emulator uses hardware virtualization, so running it inside a cloud VM requires nested virtualization, which macOS does not support out of the box. We got around this limitation by connecting to an Android emulator running on a separate Linux VM with nested virtualization enabled. With this setup, agents can build for both iOS and Android on a Mac and test on either platform, all within the same SSH session.
Each agent claims a ready Mac with an attached Android emulator. We replace Macs after use and periodically refresh idle ones.
We then tackled setup times. Initial app compilation and main checkout were loooooooong, taking over 30 minutes. So, instead of preparing a Mac after a request arrives, we switched to keeping a pool of warm Macs at the ready: main checked out (no small feat, as our monorepo is massive), tools installed, apps built, and caches warm. An agent claims a Mac, synchronizes its branch, and spins up a prebuilt Android VM instance to attach. The agent continues working in its own checkout, whether it’s running locally or in the cloud, and syncs the changes to the Mac. All scripts and builds then run on the Mac. Doing this, the time from requesting a Mac to inspecting code changes on a running device went down to just a few minutes.
Some background magic also had to be built to corral our pool. We destroy the Mac when the session ends and prepare a replacement. We also periodically replace idle Macs, pulling fresh main and rebuilding the apps and caches. We stagger this rotation so the pool stays available without interrupting active sessions.
We got the infra running, but River was still blind to it. It always stated that an iOS build cannot run on Linux, so it could code but not verify the results. So we created a skill to teach agents how to test their changes and prove their work. It required several iterations until that symbiotic relationship with two machines was seamless. It didn't start this way; even with the great tooling in place, River kept forgetting about the remote Mac, being silent while the communication with the remote Mac was happening, or taking a long time to come back to the user.
Iteration after iteration, we kept improving the skill, with the appropriate tooling in place: a CLI that encapsulates the complexity of the SSH flags. And instead of sleeping or setting timeouts for the Mac to finish long-running tasks, we used tmux panes that River could observe in real time.
Changing the way mobile developers work today
Even with all the constraints we had with our current laptops’ limitations, our teams were used to working locally for mobile development: having a real simulator to interact with, build the app, and test it ourselves. This was not only inefficient because of the laptop constraints, but also not scalable for the large migration task we had on our hands. Moving people to use River instead was not only a technological problem, it was also a psychological one. We had to change a habit.
Some people had tried early versions of River and the experience wasn’t the best: builds not working, too slow, and not being useful for day-to-day work. "This is not for me, I'll go back to coding locally" was the thought. We had to make it great for people to start using it more and more.
But as we all know, it’s not only build it and they'll come. Even internally. You need to show them. You need to evangelize about how good the tool is. And repeat it over and over. That’s why we used storytelling and internal ads to make people aware that River was now great on mobile.
We got the idea of converting River into a real-world person, frustrated about working on mobile so people could relate to the feeling. But then this real-world River gets a real Mac at home, enabling her to build mobile and, suddenly, all her problems are solved. People don't see all the complexity of the solution—the one we are explaining in this post—they only saw that now everything worked magically. Once we had the River character, it made the next internal campaigns easier. This is an example of one of those, all obviously AI-generated, making a parody of the Mac vs PC ad. This added not only storytelling and fun to the equation, but also nostalgia for the people that watched the original ad:
This ended up increasing River adoption. River initially created only 6% of merged mobile pull requests. This is how the number evolved over time, reversing the trend of local development and evolving to 61% of merged mobile pull requests:
Building in public, multiplayer and self-improving loops
One of the great things about using River is that it's multiplayer by default. Earlier this year, Tobi shared how this was already shifting Shopify’s culture and letting people learn from others’ way of using agents.
In our case, we've used this multiplayer feature not only to learn how others use River to code on mobile, but also to improve the experience for newer threads in the future. We can see where people get stuck, where guidance goes wrong, where access to tools gets rejected. All of it.
For this we used Jev , a really fast, really cheap classifier model. It observes all River conversations on mobile, tags them with multiple questions, and answers each message. We categorize each conversation based on how good the outcome was, and use the failures as input for an LLM to find improvements.
We run this iteratively, using River, so River can improve itself for future usage.
What started as a way to give River access to a mobile devel

[truncated]
