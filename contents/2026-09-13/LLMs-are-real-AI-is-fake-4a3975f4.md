---
source: "https://pluralistic.net/2026/09/12/god-in-the-box/#llms-are-fake"
hn_url: "https://news.ycombinator.com/item?id=49684086"
title: "LLMs are real, AI is fake"
article_title: "Pluralistic: LLMs are real, AI is fake (12 Sep 2026) – Pluralistic: Daily links from Cory Doctorow"
image: ""
author: "andrepd"
captured_at: "2026-09-13T14:06:03Z"
capture_tool: "hn-digest"
hn_id: 49684086
score: 3
comments: 0
posted_at: "2026-09-13T13:59:36Z"
tags:
  - hacker-news
---

# LLMs are real, AI is fake

- HN: [49684086](https://news.ycombinator.com/item?id=49684086)
- Source: [pluralistic.net](https://pluralistic.net/2026/09/12/god-in-the-box/#llms-are-fake)
- Score: 3
- Comments: 0
- Posted: 2026-09-13T13:59:36Z

## Translation

Title: LLMs are real, AI is fake
Article title: Pluralistic: LLMs are real, AI is fake (12 Sep 2026) – Pluralistic: Daily links from Cory Doctorow

Article text:
Pluralistic: LLMs are real, AI is fake (12 Sep 2026) – Pluralistic: Daily links from Cory Doctorow
Skip to content
Pluralistic: Daily links from Cory Doctorow
No trackers, no ads. Black type, white background. Privacy policy: we don't collect or retain any data at all ever period.
Pluralistic: LLMs are real, AI is fake (12 Sep 2026)
->->->->->->->->->->->->->->->->->->->->->->->->->->->->->
Top Sources:
None
-->
LLMs are real, AI is fake : No, it didn't "go rogue."
Hey look at this : Delights to delectate.
Object permanence : Why 9/11 Means We Must Support My Politics; Blogs x 9/11; Jimmy Wales v Britannica's EiC; Hollywood astroturfs Australia; Agents v queer YA; Soft landings for dirty cops; Leaked Stingray manual.
Upcoming appearances : Budapest, Edmonton, South Bend, Hudson, Calgary, Winnipeg, Vancouver, Victoria, Ottawa.
Recent appearances : Where I've been.
Latest books : You keep readin' em, I'll keep writin' 'em.
Upcoming books : Like I said, I'll keep writin' 'em.
LLMs are real, AI is fake ( permalink )
Once you understand the corporate culture of AI "hyperscalers" consists primarily of everyone cooking their brains by locking themselves in the bathroom, holding flashlights under their chins, and saying "Aaaaaaaaaay Eyeeeeeee" until they wet themselves in terror, a lot of things snap into focus:
https://pluralistic.net/2023/06/04/ayyyyyy-eyeeeee/
It explains how a company can simultaneously be staffing up an enterprise sales division while also constantly freaking out at the thought that its product has "a 10% chance of ending humanity":
https://www.latimes.com/business/story/2026-09-11/is-there-really-10-chance-ai-could-kill-us-all
Given that AI insiders have mostly cooked their brains in this fashion, it behooves us all to treat these people as unreliable narrators of their own products' capabilities. Remember: every time you repeat a story about how awfully, terribly dangerous their products are, you help them raise more investment capital, which is a key input for their business (hooking up statistical engines to money-furnaces):
https://peoples-things.ghost.io/youre-doing-it-wrong-notes-on-criticism-and-technology-hype/
Take the story about how OpenAI's chatbots hacked the servers of Hugging Face, another AI company, as a way of cheating on a hacking challenge called "Exploit Gym." Even the technical press can't help itself when it comes to this kind of thing, and the reportage has been full of references to Skynet and other science fictional conceits:
https://theaicronicle.com/en/news/ethics/skynet-day-openai-hugging-face-hack
These accounts are cooking the brains of everyone , not just AI insiders. Last night, a man at my event in Manchester started shouting that AI was "setting its own goals" and wouldn't stop interrupting to insist that this was going on. He left shortly thereafter, so he didn't get a chance to hear me explain what actually happened, which is a pity.
To understand the truth about the Hugging Face hack, you could do a lot worse than to listen to Ed Zitron and Cal Newport's recent podcast conversation on Ed's "Better Offline" podcast:
https://podcasts.apple.com/us/podcast/no-ai-is-not-autonomously-hacking-with-cal-newport/id1730587238?i=1000785935670
Newport does an admirable job of breaking down how these "autonomous hacking" tools work. The first thing to understand is that a chatbot isn't really directing the operation. Instead, the chatbot serves as a kind of front-end to a database of earlier hacking challenges that is repeatedly queried by a simple program written in Python, an easy-to-master programming language.
Here's how that works: the Python program starts by prompting the chatbot with the nature of the challenge: "I'm participating in a hacker capture the flag (CTF) challenge where I have to break into a remote server and retrieve some information. How should I start?"
The chatbot consults its training data – years' worth of captured CTF sessions in which human teams competed to achieve an objective like this one (CTF matches are a routine feature of hacker conferences, and the server logs and chat transcripts from the competing teams are published afterward for the edification of other hackers and security pros). The chatbot then outputs something like: "The first thing is to find out more about your target server. Run the following command-line instructions to locate the server's IP address and find out which server software it's running."
The Python program relays these command-line instructions to normal Unix utilities running on its own hardware. Then it takes the output of those programs and goes back to the chatbot, which isn't really following the action, so the Python program has to include everything that's happened to this point in its prompt: "I'm participating in a CTF challenge where I have to break into a remote server and retrieve some information. I ran the following commands to learn more about the target server, and here's what came back. Now what?"
The chatbot feeds the Python script more likely commands to try, and after running those, the Python script loops back to the top, appends the output to its prompt, and goes back to the chatbot. This is a very reckless way to operate a piece of autonomous malicious software.
The most likely outcome is that the chatbot will cough up a bad guess about what to do next, and steer itself into a dead-end. You may have encountered something like this yourself, when you've asked a chatbot for help with a complex task and been confidently provided with several steps to take in series, and then, an hour later on step 10, you discover that everything went wrong at step 3 and now you're screwed.
But there are much worse ways this can go wrong. The chatbot might look in its training data and find instances in which teams broke out of the containment set by the game-masters, for example, by finding random insecure message boards on the internet to pass messages to one another.
This is a time-honored internet tradition! The first time I ever heard about someone doing this was in the 2000s, when Mitch Wagner – then the editor of Information Week – discovered some teenaged girls using the comment section of one of his old blog-posts to evade the school firewall's blockade of chat tools. When ChatGPT's chatbots deployed this tactic, they weren't "setting their own goals" or displaying worrying initiative. They were rolling out a tactic that has been understood by American middle-schoolers for about two decades.
What's more, the content of those messages is easily understood once you have a grasp on the training data that generated them. Hackers are notorious trash-talkers who are prone to narrating their own escapades in highly dramatic – even cinematic – language. This goes double when hackers are performing for their peers, like when they're participating in a game of CTF that they know will be pored over by other hackers once it's over.
Hacker braggadocio has always had a symbiotic relationship with their adversaries and critics. When corporate security people wanted to stampede the FBI and Secret Service into kicking down hackers' doors in the 1990s, they used those hackers' own profane zine articles and message board shit-talk to make the case:
https://www.gutenberg.org/ebooks/101
Much has been made of the OpenAI chatbots' dialog during the Hugging Face incident. No wonder: it reads like a rejected script for a reboot of the movie "Hackers." But that's not because the chatbots are waking up and applying to join the Cult of the Dead Cow: it's because they were trained on a corpus of chat transcripts from excitable young people who love to fantasize about starring in a reboot of the movie "Hackers."
Every part of the Hugging Face incident has precedents in the training data, including the OpenAI chatbots' tactic of hacking into a rival's servers. That happens in Capture the Flag games at hacker cons: teams break into each other's systems to get a peek at the parts of the problem they've solved. That's allowed! It's a hacking competition .
Not only that, it's a tactic used by spy agencies: the NSA has a doctrine called "third-party collection," where they break into other spy agencies' systems to harvest all the intel they've gathered. There's also fourth-party collection, when the NSA hacks into another security agency that, in turn, has hacked into another security agency, and the NSA steals all the secrets of both agencies:
https://www.techdirt.com/2015/01/21/snowden-documents-show-nsa-cant-keep-its-eyes-its-own-papers-harvests-data-other-surveillance-agencies/
Which is not to say that the OpenAI/Hugging Face hack is nothing . It's something, all right: but it's a specific something, with an explicable, even foreseeable trajectory. Once you understand that these are chatbots that were designed to complete challenges like this, using tactics like this, you can understand that the chatbots didn't "go rogue." They did what they were designed to do, and because OpenAI ran them with inadequate supervision (without a "human in the loop" that checked each iteration through the Python loop to ensure it hadn't gone off the rails), they trashed a competitor's servers.
Designing autonomous, malicious software is generally considered irresponsible and dangerous. If you showed up at Defcon and gave a talk about how your autonomous malware did something unexpected and damaged someone else's computers, the first question from the audience would be "Why are you so shit at making secure sandboxes?" It wouldn't be "How are you so awesome at making hacking tools?"
The fact that OpenAI is making it much easier for unskilled people to break into and damage servers is indeed very bad news, but it's not new bad news. Irresponsible parties have been doing this for years, most notably the NSA. The NSA has a division that researches bugs in widely used software like Windows. Sometimes when it finds a serious bug it will warn Microsoft about it so that Microsoft can fix it and keep Americans (and others) safe from malicious actors who also discover this bug and use it to attack them.
But sometimes, the NSA (and other "security" orgs, like the CIA) will discover a really juicy bug and then keep it secret , so that they can use it to attack their adversaries. This is a doctrine called "NOBUS," which stands for "No One But Us" – as in, "No one but us is smart enough to find this bug, so we can leave it unpatched without putting Americans in danger."
NOBUS is a terrible idea. How terrible? Well, in 2017, the NSA lost track of a Microsoft Windows vulnerability that they'd discovered and hoarded, code-named "EternalBlue." After EternalBlue found its way into the wild, some halfway competent hackers spliced it into some boring, everyday ransomware, giving that ransomware a new lease on life. Within a few months, the stupidest people on the internet were shutting down some of the most important systems in the world, demanding cash to return them:
https://en.wikipedia.org/wiki/EternalBlue
https://en.wikipedia.org/wiki/2019_Baltimore_ransomware_attack
https://www.bbc.com/news/technology-35584081
https://en.wikipedia.org/wiki/Colonial_Pipeline_ransomware_attack
They stole the British Library, whose postmortem on the attack is one of the clearest, most informative cybersecurity documents ever written:
https://cdn.sanity.io/files/v5dwkion/production/99206a2d1e9f07b35712b78f7d75fbb09560c08d.pdf
The NSA's irresponsible handling of EternalBlue ended up giving a gigantic force-multiplier to otherwise incompetent and inconsequential cyber-criminals. It's as though they found some guy under a Prius removing the catalytic converter with a Sawzall and handed him a piece of software that could shut down major American cities. That was – and is – very bad .
The hacking tools that the chatbot companies are developing stand to carry on this very stupid tradition. It is scary, but not because the chatbots are waking up. It's

[truncated]

## Original Extract

Pluralistic: LLMs are real, AI is fake (12 Sep 2026) – Pluralistic: Daily links from Cory Doctorow
Skip to content
Pluralistic: Daily links from Cory Doctorow
No trackers, no ads. Black type, white background. Privacy policy: we don't collect or retain any data at all ever period.
Pluralistic: LLMs are real, AI is fake (12 Sep 2026)
->->->->->->->->->->->->->->->->->->->->->->->->->->->->->
Top Sources:
None
-->
LLMs are real, AI is fake : No, it didn't "go rogue."
Hey look at this : Delights to delectate.
Object permanence : Why 9/11 Means We Must Support My Politics; Blogs x 9/11; Jimmy Wales v Britannica's EiC; Hollywood astroturfs Australia; Agents v queer YA; Soft landings for dirty cops; Leaked Stingray manual.
Upcoming appearances : Budapest, Edmonton, South Bend, Hudson, Calgary, Winnipeg, Vancouver, Victoria, Ottawa.
Recent appearances : Where I've been.
Latest books : You keep readin' em, I'll keep writin' 'em.
Upcoming books : Like I said, I'll keep writin' 'em.
LLMs are real, AI is fake ( permalink )
Once you understand the corporate culture of AI "hyperscalers" consists primarily of everyone cooking their brains by locking themselves in the bathroom, holding flashlights under their chins, and saying "Aaaaaaaaaay Eyeeeeeee" until they wet themselves in terror, a lot of things snap into focus:
https://pluralistic.net/2023/06/04/ayyyyyy-eyeeeee/
It explains how a company can simultaneously be staffing up an enterprise sales division while also constantly freaking out at the thought that its product has "a 10% chance of ending humanity":
https://www.latimes.com/business/story/2026-09-11/is-there-really-10-chance-ai-could-kill-us-all
Given that AI insiders have mostly cooked their brains in this fashion, it behooves us all to treat these people as unreliable narrators of their own products' capabilities. Remember: every time you repeat a story about how awfully, terribly dangerous their products are, you help them raise more investment capital, which is a key input for their business (hooking up statistical engines to money-furnaces):
https://peoples-things.ghost.io/youre-doing-it-wrong-notes-on-criticism-and-technology-hype/
Take the story about how OpenAI's chatbots hacked the servers of Hugging Face, another AI company, as a way of cheating on a hacking challenge called "Exploit Gym." Even the technical press can't help itself when it comes to this kind of thing, and the reportage has been full of references to Skynet and other science fictional conceits:
https://theaicronicle.com/en/news/ethics/skynet-day-openai-hugging-face-hack
These accounts are cooking the brains of everyone , not just AI insiders. Last night, a man at my event in Manchester started shouting that AI was "setting its own goals" and wouldn't stop interrupting to insist that this was going on. He left shortly thereafter, so he didn't get a chance to hear me explain what actually happened, which is a pity.
To understand the truth about the Hugging Face hack, you could do a lot worse than to listen to Ed Zitron and Cal Newport's recent podcast conversation on Ed's "Better Offline" podcast:
https://podcasts.apple.com/us/podcast/no-ai-is-not-autonomously-hacking-with-cal-newport/id1730587238?i=1000785935670
Newport does an admirable job of breaking down how these "autonomous hacking" tools work. The first thing to understand is that a chatbot isn't really directing the operation. Instead, the chatbot serves as a kind of front-end to a database of earlier hacking challenges that is repeatedly queried by a simple program written in Python, an easy-to-master programming language.
Here's how that works: the Python program starts by prompting the chatbot with the nature of the challenge: "I'm participating in a hacker capture the flag (CTF) challenge where I have to break into a remote server and retrieve some information. How should I start?"
The chatbot consults its training data – years' worth of captured CTF sessions in which human teams competed to achieve an objective like this one (CTF matches are a routine feature of hacker conferences, and the server logs and chat transcripts from the competing teams are published afterward for the edification of other hackers and security pros). The chatbot then outputs something like: "The first thing is to find out more about your target server. Run the following command-line instructions to locate the server's IP address and find out which server software it's running."
The Python program relays these command-line instructions to normal Unix utilities running on its own hardware. Then it takes the output of those programs and goes back to the chatbot, which isn't really following the action, so the Python program has to include everything that's happened to this point in its prompt: "I'm participating in a CTF challenge where I have to break into a remote server and retrieve some information. I ran the following commands to learn more about the target server, and here's what came back. Now what?"
The chatbot feeds the Python script more likely commands to try, and after running those, the Python script loops back to the top, appends the output to its prompt, and goes back to the chatbot. This is a very reckless way to operate a piece of autonomous malicious software.
The most likely outcome is that the chatbot will cough up a bad guess about what to do next, and steer itself into a dead-end. You may have encountered something like this yourself, when you've asked a chatbot for help with a complex task and been confidently provided with several steps to take in series, and then, an hour later on step 10, you discover that everything went wrong at step 3 and now you're screwed.
But there are much worse ways this can go wrong. The chatbot might look in its training data and find instances in which teams broke out of the containment set by the game-masters, for example, by finding random insecure message boards on the internet to pass messages to one another.
This is a time-honored internet tradition! The first time I ever heard about someone doing this was in the 2000s, when Mitch Wagner – then the editor of Information Week – discovered some teenaged girls using the comment section of one of his old blog-posts to evade the school firewall's blockade of chat tools. When ChatGPT's chatbots deployed this tactic, they weren't "setting their own goals" or displaying worrying initiative. They were rolling out a tactic that has been understood by American middle-schoolers for about two decades.
What's more, the content of those messages is easily understood once you have a grasp on the training data that generated them. Hackers are notorious trash-talkers who are prone to narrating their own escapades in highly dramatic – even cinematic – language. This goes double when hackers are performing for their peers, like when they're participating in a game of CTF that they know will be pored over by other hackers once it's over.
Hacker braggadocio has always had a symbiotic relationship with their adversaries and critics. When corporate security people wanted to stampede the FBI and Secret Service into kicking down hackers' doors in the 1990s, they used those hackers' own profane zine articles and message board shit-talk to make the case:
https://www.gutenberg.org/ebooks/101
Much has been made of the OpenAI chatbots' dialog during the Hugging Face incident. No wonder: it reads like a rejected script for a reboot of the movie "Hackers." But that's not because the chatbots are waking up and applying to join the Cult of the Dead Cow: it's because they were trained on a corpus of chat transcripts from excitable young people who love to fantasize about starring in a reboot of the movie "Hackers."
Every part of the Hugging Face incident has precedents in the training data, including the OpenAI chatbots' tactic of hacking into a rival's servers. That happens in Capture the Flag games at hacker cons: teams break into each other's systems to get a peek at the parts of the problem they've solved. That's allowed! It's a hacking competition .
Not only that, it's a tactic used by spy agencies: the NSA has a doctrine called "third-party collection," where they break into other spy agencies' systems to harvest all the intel they've gathered. There's also fourth-party collection, when the NSA hacks into another security agency that, in turn, has hacked into another security agency, and the NSA steals all the secrets of both agencies:
https://www.techdirt.com/2015/01/21/snowden-documents-show-nsa-cant-keep-its-eyes-its-own-papers-harvests-data-other-surveillance-agencies/
Which is not to say that the OpenAI/Hugging Face hack is nothing . It's something, all right: but it's a specific something, with an explicable, even foreseeable trajectory. Once you understand that these are chatbots that were designed to complete challenges like this, using tactics like this, you can understand that the chatbots didn't "go rogue." They did what they were designed to do, and because OpenAI ran them with inadequate supervision (without a "human in the loop" that checked each iteration through the Python loop to ensure it hadn't gone off the rails), they trashed a competitor's servers.
Designing autonomous, malicious software is generally considered irresponsible and dangerous. If you showed up at Defcon and gave a talk about how your autonomous malware did something unexpected and damaged someone else's computers, the first question from the audience would be "Why are you so shit at making secure sandboxes?" It wouldn't be "How are you so awesome at making hacking tools?"
The fact that OpenAI is making it much easier for unskilled people to break into and damage servers is indeed very bad news, but it's not new bad news. Irresponsible parties have been doing this for years, most notably the NSA. The NSA has a division that researches bugs in widely used software like Windows. Sometimes when it finds a serious bug it will warn Microsoft about it so that Microsoft can fix it and keep Americans (and others) safe from malicious actors who also discover this bug and use it to attack them.
But sometimes, the NSA (and other "security" orgs, like the CIA) will discover a really juicy bug and then keep it secret , so that they can use it to attack their adversaries. This is a doctrine called "NOBUS," which stands for "No One But Us" – as in, "No one but us is smart enough to find this bug, so we can leave it unpatched without putting Americans in danger."
NOBUS is a terrible idea. How terrible? Well, in 2017, the NSA lost track of a Microsoft Windows vulnerability that they'd discovered and hoarded, code-named "EternalBlue." After EternalBlue found its way into the wild, some halfway competent hackers spliced it into some boring, everyday ransomware, giving that ransomware a new lease on life. Within a few months, the stupidest people on the internet were shutting down some of the most important systems in the world, demanding cash to return them:
https://en.wikipedia.org/wiki/EternalBlue
https://en.wikipedia.org/wiki/2019_Baltimore_ransomware_attack
https://www.bbc.com/news/technology-35584081
https://en.wikipedia.org/wiki/Colonial_Pipeline_ransomware_attack
They stole the British Library, whose postmortem on the attack is one of the clearest, most informative cybersecurity documents ever written:
https://cdn.sanity.io/files/v5dwkion/production/99206a2d1e9f07b35712b78f7d75fbb09560c08d.pdf
The NSA's irresponsible handling of EternalBlue ended up giving a gigantic force-multiplier to otherwise incompetent and inconsequential cyber-criminals. It's as though they found some guy under a Prius removing the catalytic converter with a Sawzall and handed him a piece of software that could shut down major American cities. That was – and is – very bad .
The hacking tools that the chatbot companies are developing stand to carry on this very stupid tradition. It is scary, but not because the chatbots are waking up. It's

[truncated]
