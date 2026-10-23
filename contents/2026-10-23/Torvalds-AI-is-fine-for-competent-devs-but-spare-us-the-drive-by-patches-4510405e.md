---
source: "https://www.theregister.com/columnists/2026/10/23/torvalds-ai-is-fine-for-competent-devs-but-spare-us-the-drive-by-patches/5304357"
hn_url: "https://news.ycombinator.com/item?id=50207727"
title: "Torvalds: AI is fine for competent devs, but spare us the 'drive-by' patches"
article_title: "Torvalds: AI is fine for competent devs, but spare us the 'drive-by' patches"
image: "https://image.theregister.com/5226908.jpg?imageId=5226908&x=0&y=0&cropw=100&croph=100&panox=0&panoy=0&panow=100&panoh=100&width=1200&height=683"
author: "speckx"
captured_at: "2026-10-23T13:07:59Z"
capture_tool: "hn-digest"
hn_id: 50207727
score: 9
comments: 1
posted_at: "2026-10-23T12:12:09Z"
tags:
  - hacker-news
---

# Torvalds: AI is fine for competent devs, but spare us the 'drive-by' patches

- HN: [50207727](https://news.ycombinator.com/item?id=50207727)
- Source: [www.theregister.com](https://www.theregister.com/columnists/2026/10/23/torvalds-ai-is-fine-for-competent-devs-but-spare-us-the-drive-by-patches/5304357)
- Score: 9
- Comments: 1
- Posted: 2026-10-23T12:12:09Z

## Translation

Title: Torvalds: AI is fine for competent devs, but spare us the 'drive-by' patches
Description: Kernel maintainers stop asking who used AI, as long as the code is good

Article text:
Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Torvalds: AI is fine for competent devs, but spare us the 'drive-by' patches
Kernel maintainers stop asking who used AI, as long as the code is good
Steven J. Vaughan-Nichols
Steven
J. Vaughan-Nichols
Published
fri 23 Oct 2026 // 12:31 UTC
Make us preferred on Google
READ MORE
UK threatens subsea cable saboteurs with 14 years behind bars
2 hours ago
PHP2Shell exploit bypasses protections designed to stop remote code execution
14 hours ago
Max-severity GitLab bug allows remote code execution over 12 years of releases
17 hours ago
Disgraced Indian cert issuer leaves the cert issuing business
18 hours ago
PHP2Shell puts dozens of UK government services out of service
19 hours ago
Linus Torvalds is embracing AI to write Linux kernel code, and some people are having fits. He has a message for them. So do I.
Roy Schestowitz, editor of Tux Machines, summed up his reaction to the news as " Linus Torvalds Turns to Slop . " In the podcast Pivot to AI , podcaster David Gerard agreed: "Linux kernel goes full AI slop code." Meanwhile, in the Tom's Hardware comments section, commenter "LordVile" wrote: " This is a publicity stunt from his AI overlords [The Linux Foundation] who pay his wage." They followed this by saying this is the "same reason he's flipped-flopped on AI."
I've known Linus for a few decades, and I don't take his use of AI as evidence that he's suddenly abandoned his standards. And the very idea that he'd do anything he didn't want to do with Linux on behalf of the Linux Foundation, or anyone else, is ludicrous.
I get it, though. There are reasons to be suspicious of AI. I've written about many of them. For instance, AI model collapse , which happens when AI engines learn from other AI models and their ties to the truth fray. It's a very real problem. Not, mind you, that LLMs find the grist for their answer mills from very reliable sources in the first place.
Even when AI gets good data, as it does in programming, its owners pay little to no attention to the copyrights and licenses protecting the open source code it picked up its coding tricks from. As for paying for the information it runs on, please! The very idea that Big AI must pay writers, artists, you name it, for using their work is anathema to them.
All that said, when it comes to programming, the ship has sailed. AI is becoming part of the programming toolkit, whether its critics like it or not. Why? Because while you still can't vibe code your way into delivering production-quality software, when someone who knows what they're doing, like Torvalds, uses AI, it can be very useful.
Anyone who says otherwise is just making the same mistake the Luddites did. Who now, some of you may ask? They were textile workers who attacked machinery in the early 19th century in protest against changes threatening their wages and livelihoods. You may have noticed that while people still weave cloth by hand, industrial production is mechanized. The same is happening now with AI and programming. I think we're about halfway through the process, and by 2030 or so, people programming by "hand" will be about as common as weavers doing it by hand are today.
We have two independent AI problems: we have a large codebase that AI finds real bugs in, and that adds a lot of work on already fairly overworked maintainers. And then we have people who use AI to generate patches without actually understanding the code in question.
In short, get over it. I don't love it, either. I write for a living, and my business is being crunched too. At least in programming, quality work still counts, no matter how it's done. With writing, though, true AI slop is increasingly the rule, not the exception.
Back to Linux: at the Linux Kernel Maintainers Summit in Prague the other week, Sasha Levin, a Linux kernel developer, announced that the maintainers had decided to " deprecate the Assisted-by Tag ." That means you no longer need to say if you used AI in your code submissions.
That doesn't mean you can turn in slop, though. That code won't pass muster. As Torvalds told me in a note the other day: "I think the whole 'don't use AI for commit messages' kind of argument … Maybe it's well-intentioned, but I think it's entirely misguided. I really think we want competent people to use AI as a tool."
Competent is the key word there. Torvalds also wrote: "The 'clueless people using AI' problem is different. I do think we want to discourage that. Not discourage the people who use AI and truly want to learn and truly want to become useful members of the kernel community – but discourage the drive-by patches, and discourage people who just want to feel like they sent a patch and never really wanted to become anything more. We all start clueless, but some people spend the time and effort to become something more." Exactly so.
As Torvalds said at Open Source Summit Europe : "I think [vibe-coding is] a wonderful way to find joy in programming. It makes you feel like you're doing something relevant. It allows you, as a new programmer, to do things that you would otherwise have a really hard time with." He added: "I love the concept of AI as a gateway drug" to programming.
Besides, AI has already proven very useful for improving Linux code. As long-term Linux kernel maintainer Greg Kroah-Hartman told me back in March: "Months ago, we were getting what we called 'AI slop,' AI-generated security reports that were obviously wrong or low quality." But then, "something happened a month ago, and the world switched. Now we have real reports ." It's not just Linux, he continued. "All open source projects have real reports that are made with AI, but they're good, and they're real."
True, the maintainers are getting buried in good bug reports now, but as Kroah-Hartman and discussed in Prague, AI bug finders might eventually exhaust the readily discoverable problems, easing the pressure on maintainers.
Another AI win is that you can use AI to reverse-engineer Windows drivers for Linux more easily. Christian F.K. Schaller, Red Hat's director of software engineering, created hw-harness while working on a Linux driver for a slide scanner. This open source project provides " tooling for working out how a piece of hardware behaves so that you can write a Linux driver for it."
By automating some seriously slow, tedious ways of figuring out what a binary-only driver was doing with the hardware, Schaller has made building Linux drivers from Windows drivers much easier. I know, because I did this myself back in the mid-2000s, and those manual methods are still what people use today. This could go a long way toward fixing one of those lingering annoyances of the Linux desktop: poor or entirely missing hardware drivers.
In the meantime, Torvalds told me: "Well, in all fairness, we do clearly have an 'AI problem.' Or rather, we have two independent AI problems: we have a large codebase that AI finds real bugs in, and that adds a lot of work on already fairly overworked maintainers. And then we have people who use AI to generate patches without actually understanding the code in question."
He continued: "But I'd argue that these aren't problems with AI itself; they are just problems that are exposed by AI. They are very much both very real problems, and I understand that people can really come to dislike – and even hate – AI due to them, but I do think the underlying issue isn't AI itself."
It's how people misuse AI. In the right hands, as the top Linux programmers are showing us, AI can be useful. ®
Musk disputes NASA safety panel's lunar launch estimates
Watchdog also warns retiring Falcon and Dragon could leave US access to orbit short of a backup
Torvalds: AI is fine for competent devs, but spare us the 'drive-by' patches
Kernel maintainers stop asking who used AI, as long as the code is good
Why the immediate future of liquid cooling is hybrid
PARTNER CONTENT: AI is driving liquid cooling adoption, but hybrid architectures will help operators balance density, flexibility, and operational realities
BOFH: I talked to Claude and he says you're an idiot
You know what's really sick? Getting to the pub on time on a Friday
AI’s epic cash call is matched only by its ability to strangle revenue streams
The first technology created with a self-disrupt button... But why oh why didn't they call Claude clone Fauxtoshop?
UK threatens subsea cable saboteurs with 14 years behind bars
Whitehall wants tougher penalties and new powers over telecoms operators as Russian activity raises alarm
Indian certification authority eMudhra’s private keys compromised – one of them was the number ‘4’
PWNED
US Air Force pilots used malware-laden USB drives to plan Afghanistan missions
AI bubble could burst and take global economy with it, IMF warns
OpenClaw creator says he spent $10 million in tokens last month
PHP devs publish critical bug fix as site owners warned: Pull 'em offline if you're unpatched
UPDATED
Capita set to miss legal deadline on 93,000 civil servants' pension fixes, claims union
Torvalds: AI is fine for competent devs, but spare us the 'drive-by' patches
Kernel maintainers stop asking who used AI, as long as the code is good
Nvidia woos wary Europeans with more open weights models
New embedding, parsing, and reranking models promise smarter agentic RAG workflows
Jensen Huang sends in the robot ducks to sell Nvidia's AI-everywhere vision
Containerized datacenters, androids, and telecoms all feature in Berlin keynote interrupted by environmental protesters
Nvidia's grand plan for speeding up datacenter deployments: AI trailer parks
GPUzilla boasts modular bit barns could speed up deployments from six months to just two weeks
The spec is the new code, so brush up on your prose, says Kaseya CTO
Domain expertise and clear writing will matter more than syntax, and yes, that means documentation
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Red Hat helps teach about tech writing
But learns tough lessons about AI in the process
NetBSD may have found a new role: The AI-free FOSS OS
Increasing awareness and adoption, followed by increasing donations
OpenAI dots inspire open source imitators amid technical difficulties
Persistent agents appeal to AI companies, but demand isn't obvious
Ubuntu 'Stonking Stingray' beta swims out: Deeper Rust 'oxidization' plus LLM speech-to-text
GNOME 51 is here, and Ubuntu interim releases are getting interesting again
Revive Raskin's 3 laws of software with humane work that considers a dev's duty of care
All this, and angry NeoVim users
Firefox 157: Could bold new look be just what Mozilla needs?*
Yet another new version of Firefox, with a significant reskin
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.

## Original Extract

Kernel maintainers stop asking who used AI, as long as the code is good

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Torvalds: AI is fine for competent devs, but spare us the 'drive-by' patches
Kernel maintainers stop asking who used AI, as long as the code is good
Steven J. Vaughan-Nichols
Steven
J. Vaughan-Nichols
Published
fri 23 Oct 2026 // 12:31 UTC
Make us preferred on Google
READ MORE
UK threatens subsea cable saboteurs with 14 years behind bars
2 hours ago
PHP2Shell exploit bypasses protections designed to stop remote code execution
14 hours ago
Max-severity GitLab bug allows remote code execution over 12 years of releases
17 hours ago
Disgraced Indian cert issuer leaves the cert issuing business
18 hours ago
PHP2Shell puts dozens of UK government services out of service
19 hours ago
Linus Torvalds is embracing AI to write Linux kernel code, and some people are having fits. He has a message for them. So do I.
Roy Schestowitz, editor of Tux Machines, summed up his reaction to the news as " Linus Torvalds Turns to Slop . " In the podcast Pivot to AI , podcaster David Gerard agreed: "Linux kernel goes full AI slop code." Meanwhile, in the Tom's Hardware comments section, commenter "LordVile" wrote: " This is a publicity stunt from his AI overlords [The Linux Foundation] who pay his wage." They followed this by saying this is the "same reason he's flipped-flopped on AI."
I've known Linus for a few decades, and I don't take his use of AI as evidence that he's suddenly abandoned his standards. And the very idea that he'd do anything he didn't want to do with Linux on behalf of the Linux Foundation, or anyone else, is ludicrous.
I get it, though. There are reasons to be suspicious of AI. I've written about many of them. For instance, AI model collapse , which happens when AI engines learn from other AI models and their ties to the truth fray. It's a very real problem. Not, mind you, that LLMs find the grist for their answer mills from very reliable sources in the first place.
Even when AI gets good data, as it does in programming, its owners pay little to no attention to the copyrights and licenses protecting the open source code it picked up its coding tricks from. As for paying for the information it runs on, please! The very idea that Big AI must pay writers, artists, you name it, for using their work is anathema to them.
All that said, when it comes to programming, the ship has sailed. AI is becoming part of the programming toolkit, whether its critics like it or not. Why? Because while you still can't vibe code your way into delivering production-quality software, when someone who knows what they're doing, like Torvalds, uses AI, it can be very useful.
Anyone who says otherwise is just making the same mistake the Luddites did. Who now, some of you may ask? They were textile workers who attacked machinery in the early 19th century in protest against changes threatening their wages and livelihoods. You may have noticed that while people still weave cloth by hand, industrial production is mechanized. The same is happening now with AI and programming. I think we're about halfway through the process, and by 2030 or so, people programming by "hand" will be about as common as weavers doing it by hand are today.
We have two independent AI problems: we have a large codebase that AI finds real bugs in, and that adds a lot of work on already fairly overworked maintainers. And then we have people who use AI to generate patches without actually understanding the code in question.
In short, get over it. I don't love it, either. I write for a living, and my business is being crunched too. At least in programming, quality work still counts, no matter how it's done. With writing, though, true AI slop is increasingly the rule, not the exception.
Back to Linux: at the Linux Kernel Maintainers Summit in Prague the other week, Sasha Levin, a Linux kernel developer, announced that the maintainers had decided to " deprecate the Assisted-by Tag ." That means you no longer need to say if you used AI in your code submissions.
That doesn't mean you can turn in slop, though. That code won't pass muster. As Torvalds told me in a note the other day: "I think the whole 'don't use AI for commit messages' kind of argument … Maybe it's well-intentioned, but I think it's entirely misguided. I really think we want competent people to use AI as a tool."
Competent is the key word there. Torvalds also wrote: "The 'clueless people using AI' problem is different. I do think we want to discourage that. Not discourage the people who use AI and truly want to learn and truly want to become useful members of the kernel community – but discourage the drive-by patches, and discourage people who just want to feel like they sent a patch and never really wanted to become anything more. We all start clueless, but some people spend the time and effort to become something more." Exactly so.
As Torvalds said at Open Source Summit Europe : "I think [vibe-coding is] a wonderful way to find joy in programming. It makes you feel like you're doing something relevant. It allows you, as a new programmer, to do things that you would otherwise have a really hard time with." He added: "I love the concept of AI as a gateway drug" to programming.
Besides, AI has already proven very useful for improving Linux code. As long-term Linux kernel maintainer Greg Kroah-Hartman told me back in March: "Months ago, we were getting what we called 'AI slop,' AI-generated security reports that were obviously wrong or low quality." But then, "something happened a month ago, and the world switched. Now we have real reports ." It's not just Linux, he continued. "All open source projects have real reports that are made with AI, but they're good, and they're real."
True, the maintainers are getting buried in good bug reports now, but as Kroah-Hartman and discussed in Prague, AI bug finders might eventually exhaust the readily discoverable problems, easing the pressure on maintainers.
Another AI win is that you can use AI to reverse-engineer Windows drivers for Linux more easily. Christian F.K. Schaller, Red Hat's director of software engineering, created hw-harness while working on a Linux driver for a slide scanner. This open source project provides " tooling for working out how a piece of hardware behaves so that you can write a Linux driver for it."
By automating some seriously slow, tedious ways of figuring out what a binary-only driver was doing with the hardware, Schaller has made building Linux drivers from Windows drivers much easier. I know, because I did this myself back in the mid-2000s, and those manual methods are still what people use today. This could go a long way toward fixing one of those lingering annoyances of the Linux desktop: poor or entirely missing hardware drivers.
In the meantime, Torvalds told me: "Well, in all fairness, we do clearly have an 'AI problem.' Or rather, we have two independent AI problems: we have a large codebase that AI finds real bugs in, and that adds a lot of work on already fairly overworked maintainers. And then we have people who use AI to generate patches without actually understanding the code in question."
He continued: "But I'd argue that these aren't problems with AI itself; they are just problems that are exposed by AI. They are very much both very real problems, and I understand that people can really come to dislike – and even hate – AI due to them, but I do think the underlying issue isn't AI itself."
It's how people misuse AI. In the right hands, as the top Linux programmers are showing us, AI can be useful. ®
Musk disputes NASA safety panel's lunar launch estimates
Watchdog also warns retiring Falcon and Dragon could leave US access to orbit short of a backup
Torvalds: AI is fine for competent devs, but spare us the 'drive-by' patches
Kernel maintainers stop asking who used AI, as long as the code is good
Why the immediate future of liquid cooling is hybrid
PARTNER CONTENT: AI is driving liquid cooling adoption, but hybrid architectures will help operators balance density, flexibility, and operational realities
BOFH: I talked to Claude and he says you're an idiot
You know what's really sick? Getting to the pub on time on a Friday
AI’s epic cash call is matched only by its ability to strangle revenue streams
The first technology created with a self-disrupt button... But why oh why didn't they call Claude clone Fauxtoshop?
UK threatens subsea cable saboteurs with 14 years behind bars
Whitehall wants tougher penalties and new powers over telecoms operators as Russian activity raises alarm
Indian certification authority eMudhra’s private keys compromised – one of them was the number ‘4’
PWNED
US Air Force pilots used malware-laden USB drives to plan Afghanistan missions
AI bubble could burst and take global economy with it, IMF warns
OpenClaw creator says he spent $10 million in tokens last month
PHP devs publish critical bug fix as site owners warned: Pull 'em offline if you're unpatched
UPDATED
Capita set to miss legal deadline on 93,000 civil servants' pension fixes, claims union
Torvalds: AI is fine for competent devs, but spare us the 'drive-by' patches
Kernel maintainers stop asking who used AI, as long as the code is good
Nvidia woos wary Europeans with more open weights models
New embedding, parsing, and reranking models promise smarter agentic RAG workflows
Jensen Huang sends in the robot ducks to sell Nvidia's AI-everywhere vision
Containerized datacenters, androids, and telecoms all feature in Berlin keynote interrupted by environmental protesters
Nvidia's grand plan for speeding up datacenter deployments: AI trailer parks
GPUzilla boasts modular bit barns could speed up deployments from six months to just two weeks
The spec is the new code, so brush up on your prose, says Kaseya CTO
Domain expertise and clear writing will matter more than syntax, and yes, that means documentation
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Red Hat helps teach about tech writing
But learns tough lessons about AI in the process
NetBSD may have found a new role: The AI-free FOSS OS
Increasing awareness and adoption, followed by increasing donations
OpenAI dots inspire open source imitators amid technical difficulties
Persistent agents appeal to AI companies, but demand isn't obvious
Ubuntu 'Stonking Stingray' beta swims out: Deeper Rust 'oxidization' plus LLM speech-to-text
GNOME 51 is here, and Ubuntu interim releases are getting interesting again
Revive Raskin's 3 laws of software with humane work that considers a dev's duty of care
All this, and angry NeoVim users
Firefox 157: Could bold new look be just what Mozilla needs?*
Yet another new version of Firefox, with a significant reskin
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
