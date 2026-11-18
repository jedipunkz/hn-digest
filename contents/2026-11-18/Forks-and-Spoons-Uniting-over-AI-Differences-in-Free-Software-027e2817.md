---
source: "https://www.oblomovka.com/wp/2026/11/16/forks-and-spoons-uniting-over-ai-differences-in-free-software/"
hn_url: "https://news.ycombinator.com/item?id=50553217"
title: "Forks and Spoons: Uniting over AI Differences in Free Software"
article_title: "Forks and Spoons: Uniting over AI Differences in Free Software – Oblomovka"
image: "https://www.oblomovka.com/images/wp/forks-and-spoons.jpg"
author: "dannyobrien"
captured_at: "2026-11-18T23:17:51Z"
capture_tool: "hn-digest"
hn_id: 50553217
score: 1
comments: 0
posted_at: "2026-11-18T22:28:50Z"
tags:
  - hacker-news
---

# Forks and Spoons: Uniting over AI Differences in Free Software

- HN: [50553217](https://news.ycombinator.com/item?id=50553217)
- Source: [www.oblomovka.com](https://www.oblomovka.com/wp/2026/11/16/forks-and-spoons-uniting-over-ai-differences-in-free-software/)
- Score: 1
- Comments: 0
- Posted: 2026-11-18T22:28:50Z

## Translation

Title: Forks and Spoons: Uniting over AI Differences in Free Software
Article title: Forks and Spoons: Uniting over AI Differences in Free Software – Oblomovka
Description: It’s impressive how many of the organizational techniques of large-scale software engineering have spread from its confines to the rest of society — especially given that the entire industry was an absolute trashfire of failures…

Article text:
Forks and Spoons: Uniting over AI Differences in Free Software
On the Cordon Sanitaire between Free Software and AI
Forks and Spoons: Uniting over AI Differences in Free Software
It’s impressive how many of the organizational techniques of large-scale software engineering have spread from its confines to the rest of society — especially given that the entire industry was an absolute trashfire of failures and cost overruns throughout the 70s and 80s.
For me, the queen of these social and technological innovations is the fork. God, I could write a hymn to forking, except everybody would think I was being smutty.
It’s such a simple idea: that you can take a snapshot in time of a codebase — the blueprint to a project or even an entire organization — and take it in a new direction, while leaving others to continue in their direction. It falls naturally out of the nature of digital data, but it has such revolutionary consequences.
Before the idea of a fork, every structural and coding decision felt — and usually was — expensive and distracting. After the concept and its concomitant tooling propagated, we learned that we could treat not just codebases, but the processes of engineering codebases, as items to be safely diverged from, reversed, tweaked in flight, and recombined.
Forks are a powerful abstraction for managing conflict, both as a metaphor and as a literal tool. Disagree about the direction an institution is taking? Fork it! If its processes and tooling are well documented and open source, you can copy — at undoubtedly non-zero but relatively low cost — that knowledge, and use it to take another path. You can take your toys, go home, and leave your toys for others to continue playing with.
Given the intensity of our convictions in the free software community, I’m sure we would have torn each other apart in internecine warfare long ago, if it hadn’t been for the fork. History shows how forks create options, provide a right of exit, and defuse pressure (if not hurt feelings):
In 1992, Lucid Inc., tired of waiting for GNU Emacs 19, shipped their own , and the resulting Lucid Emacs (later XEmacs) spent two decades as GNU Emacs’s rival, each stealing the other’s best ideas, until XEmacs faded and its users drifted home. In 1997, a faction disgruntled with GCC’s development forked it and ran their own competing implementation, EGCS , for two years until the original GCC’ers conceded defeat, and amicably crowned the fork the real GCC. And so on.
So are forks the solution to the AI divide in free software ?
Well, in the initial skirmishes, they’ve certainly proven to be the sharpest implement immediately at hand. The forks are already flying: Vim Classic is a rebased fork of Vim 8.x, “maintained entirely by humans”; Human Emacs is the fork-in-waiting, if the GNU Project does not adopt a sufficiently strong anti-AI policy. Opentmux , gaybar , Gram and Arcalibre are surely just the first of many such splits.
No-AI forks are firmly in the spirit of the early, rough years of free software, when RMS single-handedly stood athwart Symbolics , personally re-implementing in shareable form every feature that the company released in its proprietary fork of the Lisp Machine implementation.
That form of last-stand battle against the enclosure of the software commons proved to be exhausting and unsustainable. Your opponent is well-financed and legion: maintaining the fork against their advances is a race that you don’t even want to be competing in. Plus if you’re truly open and they are not, you cannot use their work, but they can take your code. The harshness of this path was surely part of the inspiration for the later genius of the GNU General Public License , which sought to bend things slightly in free software’s favour in a fight with the proprietary knowledge-hoarders.
Even without concerns about models incorporating and replicating GNU code in their training corpus, a modern-day fork between AI-friendly and AI-rejecting branches similarly favours the AI-friendly side — at least at the start of the split. Even if the AI-rejecting groups are better resourced with volunteers, the machine-enhanced coders can surely, pound-for-pound, ideate, prototype and crank out code, identify security flaws, and patch bugs faster than the human-only branch.
Within AI forks in FLOSS itself, this isn’t a battle between proprietary and open source codebases, but the disparity is similar. The AI-sceptical have deliberately cut themselves off from the AIers’ contributions to the code commons. AI-using coders can and will draw from the no-AI fork; the AI-sceptical, however, have sworn off adopting any code that has been polluted by the touch of a bot.
My colleagues on the anti-AI side of this debate will say, I think, that even though the battle is hard, it’s important to fight it, to defend principles — including but not limited to principles that are foundational to free software. Besides, they will say, while in the short run AI code may win in sheer lines of code, model collapse and slop will bury the AI-users under the weight of their contradictions.
I don’t agree, but I don’t know that I’m right. I’m certainly not confident enough to bet the future of free culture on either side.
The two branches rest on two different predictions of the future. I think we should try both — and, for my own sanity, in the same timeline.
Uniting Against Non-Consensual AI
Apart from the principle of the thing, there are practical reasons why so many open-source coders are disgusted by AI: its harmful, apparently inescapable effects on many of their existing practices. Projects that accept code from outsiders — the vast majority — are swamped with AI-enabled contributions. Security bugs are being found and thrown at them at inhuman (and untriaged) rates. Their servers creak and groan under the weight of a million scrapers , all apparently from eager AI model-builders feeding on their code.
It’s no consolation to any of these tired and disillusioned maintainers that the wider world’s answer to many of these AI problems seems to be — add more AI. AI builders leave instructions for agents and their humans to either stop providing code (just describe the bug or feature in prose , and our agents will tackle it), or just leave the agents, both the project’s and the potential contributor’s, to thrash out their differences themselves. Bot reports, triaged by bots. Servers guarded by bandwidth-tweaking agents, 24/7. AI: Cause of, and Solution to, the World’s Problems!
… And I see you’ve played Forkey-Spoony before
Forks are not, now, a particularly new social innovation, but they’re not exactly ancient . We’re almost certainly still in the earliest moments of novel, digitally-inspired, social technologies.
So here’s what I propose, as a brand new social experiment in our still evolving free culture.
Instead of seeing the AI/non-AI divide as leading inevitably to hostile forks and a divided community, let’s imagine a new variant of the friendly fork. Let’s call it a spoon.
In this, the AI-friendly side of the divide — whether the mainstream of a project like the Linux kernel whose governance has decided to accept AI-assisted contributions , or an AI-friendly scion of a project that has chosen to ban AI — take on both the opportunity and the burdens of current AI on behalf of the project as a whole, both sides of the divide.
They protect the AI-rejecters from the AI storm, by explicitly accepting and actively attracting the AI contributions. For distribution projects like Guix , they run a channel — and the infrastructure around it — to package AI-generated content and code. They host the mirrors that take the scrapers. They experiment with free-software-friendly AI processes, talk to the AI-using world, and preach the good news about free software to it.
Spooning next to their colleagues, they bend to cover and protect their friends on the flip side of the fork, and make sure they can continue to do their work, free of AI’s negative externalities — and hopefully better placed to take advantage of the parts that can improve the grand project of free software for everyone, under everyone’s personal control.
They also make sure that knowledge, and ideally something close to code, flows both ways. How will take experiment and negotiation between the camps. One can imagine a parallel to clean-room implementation: AI-users write specifications that carry what they’ve learned from the codebase over to the other side of the fork. Or they help build interfaces that let free software on both sides of the divide talk to each other (looking at you, OCapN ).
Specifically: I’m a huge fan of Guix, and many of the projects I support — including Spritely , one of the participants in OCapN — have fates that are intimately tied to Guix’s future success. I firmly believe that Guix and its community are the future of the GNU Project. Guix is currently in an unhappy state: its community dislikes AI, but its proposal to restrict AI failed for reasons unrelated to this general consensus. That leaves it trapped in the worst of all worlds — exposed to GenAI works with no clear rules, but deeply reluctant to deploy AI to mitigate that storm.
My hope is that we could create a spoon to both protect them and similar AI-resistant project cores from the intrusions of AI, while also experimenting with how Guix and reproducible software might prosper in a world where AI and humans work together.
It would require AI-friendly Guix developers who want the AI-sceptical core of Guix to succeed, and a certain amount of acceptance from the AI-sceptical Guixen that such work is worth co-operating with, despite their differences. I’m a little pre-occupied right now (having launched not one , but two major projects), but I’d certainly like to hear whether there are any other AI-loving Guix-fan unicorns out there willing to work on such a spoon. Mail me !
As an AI optimist, I believe the free software community would benefit from learning through the use of AI models, experimenting with them, and adapting them, the better to bend and liberate their power for universal liberatory ends.
But far more important to me than the free software community learning to use LLMs is that the principles and understanding and practices of the free software movement spread to new AI-using communities. With an explosion in code generation and use comes an opportunity to embed open culture in a new generation of software and its users.
And to do that, we need AI-using free software advocates who are willing to step up to protect the core of their human, humane project, and act as apostles of its values to our new, AI-haunted world.
This entry was posted
on Monday, November 16th, 2026 at 10:10 pm and is filed under AI , Inkhaven , Technology .
You can follow any responses to this entry through the RSS/Atom feed.
You can leave a response , or trackback from your own site.
One Response to “Forks and Spoons: Uniting over AI Differences in Free Software”
On the Cordon Sanitaire between Free Software and AI – Oblomovka Says:
November 16th, 2026 at 10:21 pm
[…] And that’s why I think that to face these challenges, advocates for free software should use the most powerful tools in their organizational cutlery drawer: the fork and spoon. […]
Mail (will not be published) (required)
XHTML: You can use these tags: <a href="" title=""> <abbr title=""> <acronym title=""> <b> <blockquote cite=""> <cite> <code class="" title="" data-url=""> <del datetime=""> <em> <i> <q cite=""> <s> <strike> <strong> <pre class="" title="" data-url=""> <span class="" title="" data-url="">

## Original Extract

It’s impressive how many of the organizational techniques of large-scale software engineering have spread from its confines to the rest of society — especially given that the entire industry was an absolute trashfire of failures…

Forks and Spoons: Uniting over AI Differences in Free Software
On the Cordon Sanitaire between Free Software and AI
Forks and Spoons: Uniting over AI Differences in Free Software
It’s impressive how many of the organizational techniques of large-scale software engineering have spread from its confines to the rest of society — especially given that the entire industry was an absolute trashfire of failures and cost overruns throughout the 70s and 80s.
For me, the queen of these social and technological innovations is the fork. God, I could write a hymn to forking, except everybody would think I was being smutty.
It’s such a simple idea: that you can take a snapshot in time of a codebase — the blueprint to a project or even an entire organization — and take it in a new direction, while leaving others to continue in their direction. It falls naturally out of the nature of digital data, but it has such revolutionary consequences.
Before the idea of a fork, every structural and coding decision felt — and usually was — expensive and distracting. After the concept and its concomitant tooling propagated, we learned that we could treat not just codebases, but the processes of engineering codebases, as items to be safely diverged from, reversed, tweaked in flight, and recombined.
Forks are a powerful abstraction for managing conflict, both as a metaphor and as a literal tool. Disagree about the direction an institution is taking? Fork it! If its processes and tooling are well documented and open source, you can copy — at undoubtedly non-zero but relatively low cost — that knowledge, and use it to take another path. You can take your toys, go home, and leave your toys for others to continue playing with.
Given the intensity of our convictions in the free software community, I’m sure we would have torn each other apart in internecine warfare long ago, if it hadn’t been for the fork. History shows how forks create options, provide a right of exit, and defuse pressure (if not hurt feelings):
In 1992, Lucid Inc., tired of waiting for GNU Emacs 19, shipped their own , and the resulting Lucid Emacs (later XEmacs) spent two decades as GNU Emacs’s rival, each stealing the other’s best ideas, until XEmacs faded and its users drifted home. In 1997, a faction disgruntled with GCC’s development forked it and ran their own competing implementation, EGCS , for two years until the original GCC’ers conceded defeat, and amicably crowned the fork the real GCC. And so on.
So are forks the solution to the AI divide in free software ?
Well, in the initial skirmishes, they’ve certainly proven to be the sharpest implement immediately at hand. The forks are already flying: Vim Classic is a rebased fork of Vim 8.x, “maintained entirely by humans”; Human Emacs is the fork-in-waiting, if the GNU Project does not adopt a sufficiently strong anti-AI policy. Opentmux , gaybar , Gram and Arcalibre are surely just the first of many such splits.
No-AI forks are firmly in the spirit of the early, rough years of free software, when RMS single-handedly stood athwart Symbolics , personally re-implementing in shareable form every feature that the company released in its proprietary fork of the Lisp Machine implementation.
That form of last-stand battle against the enclosure of the software commons proved to be exhausting and unsustainable. Your opponent is well-financed and legion: maintaining the fork against their advances is a race that you don’t even want to be competing in. Plus if you’re truly open and they are not, you cannot use their work, but they can take your code. The harshness of this path was surely part of the inspiration for the later genius of the GNU General Public License , which sought to bend things slightly in free software’s favour in a fight with the proprietary knowledge-hoarders.
Even without concerns about models incorporating and replicating GNU code in their training corpus, a modern-day fork between AI-friendly and AI-rejecting branches similarly favours the AI-friendly side — at least at the start of the split. Even if the AI-rejecting groups are better resourced with volunteers, the machine-enhanced coders can surely, pound-for-pound, ideate, prototype and crank out code, identify security flaws, and patch bugs faster than the human-only branch.
Within AI forks in FLOSS itself, this isn’t a battle between proprietary and open source codebases, but the disparity is similar. The AI-sceptical have deliberately cut themselves off from the AIers’ contributions to the code commons. AI-using coders can and will draw from the no-AI fork; the AI-sceptical, however, have sworn off adopting any code that has been polluted by the touch of a bot.
My colleagues on the anti-AI side of this debate will say, I think, that even though the battle is hard, it’s important to fight it, to defend principles — including but not limited to principles that are foundational to free software. Besides, they will say, while in the short run AI code may win in sheer lines of code, model collapse and slop will bury the AI-users under the weight of their contradictions.
I don’t agree, but I don’t know that I’m right. I’m certainly not confident enough to bet the future of free culture on either side.
The two branches rest on two different predictions of the future. I think we should try both — and, for my own sanity, in the same timeline.
Uniting Against Non-Consensual AI
Apart from the principle of the thing, there are practical reasons why so many open-source coders are disgusted by AI: its harmful, apparently inescapable effects on many of their existing practices. Projects that accept code from outsiders — the vast majority — are swamped with AI-enabled contributions. Security bugs are being found and thrown at them at inhuman (and untriaged) rates. Their servers creak and groan under the weight of a million scrapers , all apparently from eager AI model-builders feeding on their code.
It’s no consolation to any of these tired and disillusioned maintainers that the wider world’s answer to many of these AI problems seems to be — add more AI. AI builders leave instructions for agents and their humans to either stop providing code (just describe the bug or feature in prose , and our agents will tackle it), or just leave the agents, both the project’s and the potential contributor’s, to thrash out their differences themselves. Bot reports, triaged by bots. Servers guarded by bandwidth-tweaking agents, 24/7. AI: Cause of, and Solution to, the World’s Problems!
… And I see you’ve played Forkey-Spoony before
Forks are not, now, a particularly new social innovation, but they’re not exactly ancient . We’re almost certainly still in the earliest moments of novel, digitally-inspired, social technologies.
So here’s what I propose, as a brand new social experiment in our still evolving free culture.
Instead of seeing the AI/non-AI divide as leading inevitably to hostile forks and a divided community, let’s imagine a new variant of the friendly fork. Let’s call it a spoon.
In this, the AI-friendly side of the divide — whether the mainstream of a project like the Linux kernel whose governance has decided to accept AI-assisted contributions , or an AI-friendly scion of a project that has chosen to ban AI — take on both the opportunity and the burdens of current AI on behalf of the project as a whole, both sides of the divide.
They protect the AI-rejecters from the AI storm, by explicitly accepting and actively attracting the AI contributions. For distribution projects like Guix , they run a channel — and the infrastructure around it — to package AI-generated content and code. They host the mirrors that take the scrapers. They experiment with free-software-friendly AI processes, talk to the AI-using world, and preach the good news about free software to it.
Spooning next to their colleagues, they bend to cover and protect their friends on the flip side of the fork, and make sure they can continue to do their work, free of AI’s negative externalities — and hopefully better placed to take advantage of the parts that can improve the grand project of free software for everyone, under everyone’s personal control.
They also make sure that knowledge, and ideally something close to code, flows both ways. How will take experiment and negotiation between the camps. One can imagine a parallel to clean-room implementation: AI-users write specifications that carry what they’ve learned from the codebase over to the other side of the fork. Or they help build interfaces that let free software on both sides of the divide talk to each other (looking at you, OCapN ).
Specifically: I’m a huge fan of Guix, and many of the projects I support — including Spritely , one of the participants in OCapN — have fates that are intimately tied to Guix’s future success. I firmly believe that Guix and its community are the future of the GNU Project. Guix is currently in an unhappy state: its community dislikes AI, but its proposal to restrict AI failed for reasons unrelated to this general consensus. That leaves it trapped in the worst of all worlds — exposed to GenAI works with no clear rules, but deeply reluctant to deploy AI to mitigate that storm.
My hope is that we could create a spoon to both protect them and similar AI-resistant project cores from the intrusions of AI, while also experimenting with how Guix and reproducible software might prosper in a world where AI and humans work together.
It would require AI-friendly Guix developers who want the AI-sceptical core of Guix to succeed, and a certain amount of acceptance from the AI-sceptical Guixen that such work is worth co-operating with, despite their differences. I’m a little pre-occupied right now (having launched not one , but two major projects), but I’d certainly like to hear whether there are any other AI-loving Guix-fan unicorns out there willing to work on such a spoon. Mail me !
As an AI optimist, I believe the free software community would benefit from learning through the use of AI models, experimenting with them, and adapting them, the better to bend and liberate their power for universal liberatory ends.
But far more important to me than the free software community learning to use LLMs is that the principles and understanding and practices of the free software movement spread to new AI-using communities. With an explosion in code generation and use comes an opportunity to embed open culture in a new generation of software and its users.
And to do that, we need AI-using free software advocates who are willing to step up to protect the core of their human, humane project, and act as apostles of its values to our new, AI-haunted world.
This entry was posted
on Monday, November 16th, 2026 at 10:10 pm and is filed under AI , Inkhaven , Technology .
You can follow any responses to this entry through the RSS/Atom feed.
You can leave a response , or trackback from your own site.
One Response to “Forks and Spoons: Uniting over AI Differences in Free Software”
On the Cordon Sanitaire between Free Software and AI – Oblomovka Says:
November 16th, 2026 at 10:21 pm
[…] And that’s why I think that to face these challenges, advocates for free software should use the most powerful tools in their organizational cutlery drawer: the fork and spoon. […]
Mail (will not be published) (required)
XHTML: You can use these tags: <a href="" title=""> <abbr title=""> <acronym title=""> <b> <blockquote cite=""> <cite> <code class="" title="" data-url=""> <del datetime=""> <em> <i> <q cite=""> <s> <strike> <strong> <pre class="" title="" data-url=""> <span class="" title="" data-url="">
