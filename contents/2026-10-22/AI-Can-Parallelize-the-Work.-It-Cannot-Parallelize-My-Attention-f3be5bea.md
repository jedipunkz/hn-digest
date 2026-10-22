---
source: "https://palmesi.com/blog/knowledge-work-needs-awareness/"
hn_url: "https://news.ycombinator.com/item?id=50198718"
title: "AI Can Parallelize the Work. It Cannot Parallelize My Attention"
article_title: "AI Can Parallelize the Work. It Cannot Parallelize My Attention — Pietro Palmesi"
image: ""
author: "pizzooid"
captured_at: "2026-10-22T18:20:56Z"
capture_tool: "hn-digest"
hn_id: 50198718
score: 1
comments: 1
posted_at: "2026-10-22T17:46:49Z"
tags:
  - hacker-news
---

# AI Can Parallelize the Work. It Cannot Parallelize My Attention

- HN: [50198718](https://news.ycombinator.com/item?id=50198718)
- Source: [palmesi.com](https://palmesi.com/blog/knowledge-work-needs-awareness/)
- Score: 1
- Comments: 1
- Posted: 2026-10-22T17:46:49Z

## Translation

Title: AI Can Parallelize the Work. It Cannot Parallelize My Attention
Article title: AI Can Parallelize the Work. It Cannot Parallelize My Attention — Pietro Palmesi
Description: What coding agents taught me about embodied deep work, awareness, and restraint.

Article text:
AI Can Parallelize the Work. It Cannot Parallelize My Attention — Pietro Palmesi Pietro Palmesi ← All writing AI Can Parallelize the Work. It Cannot Parallelize My Attention
What coding agents taught me about deep work
I thought I was doing the best work of my life. Agents were busy in every window, and I was flying through tickets. Then we rolled back a month of my work. How had moving so fast taken me so far from what I was trying to do?
The bug wasn’t just in the code
This was back in the good old days of programming, when code felt like magic reserved for the few who dared to learn it. Armed with a physics degree and the confidence that came with it, I felt more or less invincible.
All it took to humble me was one nasty bug in our CI pipeline. 1
The bug was difficult to reproduce and appeared only about five minutes into the Docker build. Every attempt followed the same rhythm: change one line, start the build, wait five minutes, watch it fail, repeat.
Easy to fix , I thought in the early afternoon.
I still thought so late that evening. Just one more line. One more build. Then I would find the problem, go home and sleep like a baby.
Eventually, I gave up. I went home exhausted, crawled into bed and hoped my supervisor would not come looking for the overdue fix in the morning.
Even in bed, I was still thinking through the Dockerfile. I had difficulty falling asleep. I noticed that I was tired and irritated, but during the work itself I had barely noticed my body at all. I was just inside the code.
My supervisor was waiting for it. “Are you done with that small change?”
I am usually a calm person, but I did not take the question well. I am smart. I should have figured this out.
I looked at him coldly. “The fucking pipeline doesn’t work properly. I don’t know how long it’s going to take.”
He shrugged, perhaps thinking, One of those coders again , and left.
I returned to the code. No breaks. No kind words. No time for anything except finding the bug.
That morning, I eventually had enough distance to see the problem. A software update had introduced a backwards-incompatible change. Not in one of our direct dependencies, but in a dependency of a dependency. The solution was simple: pin the older version. 2
The bug was fixed. My state was not.
I was tired and unmotivated. I began noticing colleagues watching YouTube or reading the news and judging them for slacking off. I had completed the task, but the struggle left a bitter aftertaste.
Compare this with another morning, when a different hard-to-reproduce bug appeared.
This time, I told my supervisor immediately. I explained that the bug was hard to reproduce and that finding it could take a while. With the expectation made explicit, I could approach the problem differently.
Instead of repeatedly waiting for the entire Docker build to fail, I built the image up to the step before the failure and entered the container interactively. Now I could test possible fixes within seconds instead of restarting a five-minute build each time.
I fixed the bug faster than expected, told my supervisor and moved on to a new feature with energy left to spare.
Same programmer. Same class of problem. Very different work.
No matter how intensely I concentrate, I cannot see the forest for the trees when I do not feel safe.
The difference between those two days was not concentration. One was draining and unsuccessful; the other was productive and satisfying. If anything, I concentrated harder on the first day. The difference was my awareness of the situation I was in.
On the first day, I lost sight of my exhaustion, my fear of disappointing my supervisor and the fact that my debugging strategy was not working. On the second, I recognized the uncertainty, communicated it early and retained enough perspective to change my approach.
Staying aware of my state is crucial to doing my best knowledge work. As Pranab Sachithanandan observes in his conversation with Coaches Rising , contemplative awareness often disappears when we sit down at a laptop. Attention contracts around the conceptual world on the screen, and we lose contact with the body and the wider situation.
The goal is not to cultivate a permanently calm, fuzzy feeling. Nor is it to remain in some imaginary state of peak performance. It is to notice when our capacity is narrowing and to respond before we become trapped inside it: pause, communicate, change our approach or step away.
When stress leaves a workable range, the prefrontal functions needed for working memory, flexible thinking and self-regulation can become impaired .
Awareness lets us step outside the frame created by the current problem. It allows us to examine not only the work, but also the state from which we are doing it.
Frustration is not the only thing that can take this awareness away. Speed and excitement can do it too.
Fast-forward a decade. I now live inside an agent harness. 3
I touch the code myself only for minor edits. Most of my time is spent instructing agents and reviewing their work across multiple windows. Until a few weeks ago, whenever one agent started thinking, running tests or reading the codebase, I switched to another. Waiting felt wasteful when another window was ready for a prompt.
The value of my work no longer lies primarily in being able to write a Bash script, a for loop or even a MapReduce job. It lies in specifying precisely what I want, recognizing whether the proposed implementation makes sense and verifying that the system does what I intended.
At least, that is where the value should lie.
It seems that every other week brings a faster, more capable model. Software can be produced at an accelerating pace. The question quietly shifts from Should we build this feature? to How many agents can we keep running? How many millions of tokens can we burn?
The discomfort of waiting grew stronger. So did the feeling that I had to keep up. I could feel the competition and the urge to move forward, faster and faster.
The apparent solution was obvious: start another agent.
My understanding of what was happening in each window became shallower every time I returned to it. But I barely noticed. I felt as though I was blasting through my tickets, always just one prompt away from the correct solution.
It felt like the best work I had ever done. At the same time, I began merely glancing at the diffs. I became too impatient to specify each task in the detail it required. Eventually, I stopped verifying whether the implementation matched my intention and simply trusted the agent to get it right.
Of course, I could not concentrate on everything happening in parallel. The agents were parallel; my understanding was not.
More agents were not inherently the problem. Subagents could divide one coherent task without requiring me to hold several separate goals. The cognitive load appeared when I made myself the main agent for several independent workstreams.
This time, fear alone did not take away my awareness. Excitement did. So did speed, competition and the slot-machine feeling of agentic work: pull the lever, switch windows and return to find that a new world has been created while you were not looking.
The old trap was tunnel vision. The new one is fragmentation. From inside, both can feel like extraordinary productivity.
As I grew more comfortable working with agents, I also grew more confident in my ability to develop several things at once. My changes became more audacious and less carefully broken down.
I decided to optimize part of our system while building the instrumentation needed to measure it. I knew measurement should come first, but the first improvement seemed obvious. Why wait? I could make one part faster while measuring another.
Then our LLM provider removed the cheap, fast model we were using.
We switched to a more expensive model, but it behaved differently. Given the same JSON input, it did not select the same results. Apparently, our prompt had been less general than we thought.
Suddenly, the observability work was measuring a moving target. The model had changed, its output had changed, and the results could no longer be compared directly.
I generalized the prompts, patched the application to work with the new model, expanded the observability and continued with the performance optimization, all at the same time.
After some delay, the bundle of changes went live.
The metrics showed better performance, and that performance was now more thoroughly measurable. But the customers’ verdict was simpler: it worked worse. Could we please have the old version back?
I resisted. We had better prompts now. We had observability. The performance numbers had improved. Surely we could work from there. A few more changes, a few more days, and the whole thing would come together.
The customers remained dissatisfied.
Eventually, we rolled everything back, found another provider that still offered the original model and started again from the beginning.
In theory, I had delivered three features: improved performance, better observability and support for a new model. In practice, the product had made no progress at all.
But I had produced a shit ton of code and complexity.
Had I made one change at a time, as I would have before agents made simultaneous development feel cheap, each decision would have remained easier to evaluate and reverse. I could have discarded the failed optimization without losing the observability work, or rebased one improvement after restoring the model.
Instead, the changes became entangled. Abandoning one felt like abandoning all of them. I kept believing that a much better solution was only a few days away; admitting otherwise meant throwing away a month’s apparent progress.
When producing code becomes cheap, preserving causality becomes valuable.
What I need from deep work now
These experiences point to four principles for doing knowledge work with agents.
One coherent goal at a time. I need to understand its purpose, boundaries and criteria for success, as well as how the agents’ work fits together. If I need three windows to reconstruct what is happening, coherence is already gone.
I need to notice not only the work, but the state from which I am doing it. Am I curious, anxious, excited or exhausted? Awareness keeps concentration from becoming tunnel vision and excitement from becoming fragmentation.
Changes must remain independently testable and reversible. If prompts, models, instrumentation and optimizations change together, I can no longer tell what caused what. One change at a time may look slower; it keeps learning possible.
Less is more. Agents make it almost effortless to add another feature, abstraction or option. My job is to decide which possibilities deserve to exist and what is worth building.
Deep work in the age of agents is sustained engagement with one coherent goal, guided by awareness, structured for reversibility and disciplined by restraint.
I have started with one simple experiment: when an agent begins working, I try not to open another workstream.
Instead, I stay with the task. I reread the acceptance criteria, inspect the relevant code or simply notice my posture and breathing. I ask whether I still understand what the agent is doing and whether it still serves the original goal.
The waiting time is not a delay in the work. It is part of the work.
I have been working this way for a few weeks now. It feels slower. The ticket count suggests that it may not actually be slower, although tickets do not all speak the same language. Still, the part of me that wants to move fast does not like it.
I probably send fewer prompts and produce less code. But I understand much better what I have done. At our progress meetings every other week, I can explain the changes and the decisions behind them without reconstructing everything from agent windows.
More importantly, far fewer requests for fixes come back from users. There are fewer complaints, fewer correction cycles and

[truncated]

## Original Extract

What coding agents taught me about embodied deep work, awareness, and restraint.

AI Can Parallelize the Work. It Cannot Parallelize My Attention — Pietro Palmesi Pietro Palmesi ← All writing AI Can Parallelize the Work. It Cannot Parallelize My Attention
What coding agents taught me about deep work
I thought I was doing the best work of my life. Agents were busy in every window, and I was flying through tickets. Then we rolled back a month of my work. How had moving so fast taken me so far from what I was trying to do?
The bug wasn’t just in the code
This was back in the good old days of programming, when code felt like magic reserved for the few who dared to learn it. Armed with a physics degree and the confidence that came with it, I felt more or less invincible.
All it took to humble me was one nasty bug in our CI pipeline. 1
The bug was difficult to reproduce and appeared only about five minutes into the Docker build. Every attempt followed the same rhythm: change one line, start the build, wait five minutes, watch it fail, repeat.
Easy to fix , I thought in the early afternoon.
I still thought so late that evening. Just one more line. One more build. Then I would find the problem, go home and sleep like a baby.
Eventually, I gave up. I went home exhausted, crawled into bed and hoped my supervisor would not come looking for the overdue fix in the morning.
Even in bed, I was still thinking through the Dockerfile. I had difficulty falling asleep. I noticed that I was tired and irritated, but during the work itself I had barely noticed my body at all. I was just inside the code.
My supervisor was waiting for it. “Are you done with that small change?”
I am usually a calm person, but I did not take the question well. I am smart. I should have figured this out.
I looked at him coldly. “The fucking pipeline doesn’t work properly. I don’t know how long it’s going to take.”
He shrugged, perhaps thinking, One of those coders again , and left.
I returned to the code. No breaks. No kind words. No time for anything except finding the bug.
That morning, I eventually had enough distance to see the problem. A software update had introduced a backwards-incompatible change. Not in one of our direct dependencies, but in a dependency of a dependency. The solution was simple: pin the older version. 2
The bug was fixed. My state was not.
I was tired and unmotivated. I began noticing colleagues watching YouTube or reading the news and judging them for slacking off. I had completed the task, but the struggle left a bitter aftertaste.
Compare this with another morning, when a different hard-to-reproduce bug appeared.
This time, I told my supervisor immediately. I explained that the bug was hard to reproduce and that finding it could take a while. With the expectation made explicit, I could approach the problem differently.
Instead of repeatedly waiting for the entire Docker build to fail, I built the image up to the step before the failure and entered the container interactively. Now I could test possible fixes within seconds instead of restarting a five-minute build each time.
I fixed the bug faster than expected, told my supervisor and moved on to a new feature with energy left to spare.
Same programmer. Same class of problem. Very different work.
No matter how intensely I concentrate, I cannot see the forest for the trees when I do not feel safe.
The difference between those two days was not concentration. One was draining and unsuccessful; the other was productive and satisfying. If anything, I concentrated harder on the first day. The difference was my awareness of the situation I was in.
On the first day, I lost sight of my exhaustion, my fear of disappointing my supervisor and the fact that my debugging strategy was not working. On the second, I recognized the uncertainty, communicated it early and retained enough perspective to change my approach.
Staying aware of my state is crucial to doing my best knowledge work. As Pranab Sachithanandan observes in his conversation with Coaches Rising , contemplative awareness often disappears when we sit down at a laptop. Attention contracts around the conceptual world on the screen, and we lose contact with the body and the wider situation.
The goal is not to cultivate a permanently calm, fuzzy feeling. Nor is it to remain in some imaginary state of peak performance. It is to notice when our capacity is narrowing and to respond before we become trapped inside it: pause, communicate, change our approach or step away.
When stress leaves a workable range, the prefrontal functions needed for working memory, flexible thinking and self-regulation can become impaired .
Awareness lets us step outside the frame created by the current problem. It allows us to examine not only the work, but also the state from which we are doing it.
Frustration is not the only thing that can take this awareness away. Speed and excitement can do it too.
Fast-forward a decade. I now live inside an agent harness. 3
I touch the code myself only for minor edits. Most of my time is spent instructing agents and reviewing their work across multiple windows. Until a few weeks ago, whenever one agent started thinking, running tests or reading the codebase, I switched to another. Waiting felt wasteful when another window was ready for a prompt.
The value of my work no longer lies primarily in being able to write a Bash script, a for loop or even a MapReduce job. It lies in specifying precisely what I want, recognizing whether the proposed implementation makes sense and verifying that the system does what I intended.
At least, that is where the value should lie.
It seems that every other week brings a faster, more capable model. Software can be produced at an accelerating pace. The question quietly shifts from Should we build this feature? to How many agents can we keep running? How many millions of tokens can we burn?
The discomfort of waiting grew stronger. So did the feeling that I had to keep up. I could feel the competition and the urge to move forward, faster and faster.
The apparent solution was obvious: start another agent.
My understanding of what was happening in each window became shallower every time I returned to it. But I barely noticed. I felt as though I was blasting through my tickets, always just one prompt away from the correct solution.
It felt like the best work I had ever done. At the same time, I began merely glancing at the diffs. I became too impatient to specify each task in the detail it required. Eventually, I stopped verifying whether the implementation matched my intention and simply trusted the agent to get it right.
Of course, I could not concentrate on everything happening in parallel. The agents were parallel; my understanding was not.
More agents were not inherently the problem. Subagents could divide one coherent task without requiring me to hold several separate goals. The cognitive load appeared when I made myself the main agent for several independent workstreams.
This time, fear alone did not take away my awareness. Excitement did. So did speed, competition and the slot-machine feeling of agentic work: pull the lever, switch windows and return to find that a new world has been created while you were not looking.
The old trap was tunnel vision. The new one is fragmentation. From inside, both can feel like extraordinary productivity.
As I grew more comfortable working with agents, I also grew more confident in my ability to develop several things at once. My changes became more audacious and less carefully broken down.
I decided to optimize part of our system while building the instrumentation needed to measure it. I knew measurement should come first, but the first improvement seemed obvious. Why wait? I could make one part faster while measuring another.
Then our LLM provider removed the cheap, fast model we were using.
We switched to a more expensive model, but it behaved differently. Given the same JSON input, it did not select the same results. Apparently, our prompt had been less general than we thought.
Suddenly, the observability work was measuring a moving target. The model had changed, its output had changed, and the results could no longer be compared directly.
I generalized the prompts, patched the application to work with the new model, expanded the observability and continued with the performance optimization, all at the same time.
After some delay, the bundle of changes went live.
The metrics showed better performance, and that performance was now more thoroughly measurable. But the customers’ verdict was simpler: it worked worse. Could we please have the old version back?
I resisted. We had better prompts now. We had observability. The performance numbers had improved. Surely we could work from there. A few more changes, a few more days, and the whole thing would come together.
The customers remained dissatisfied.
Eventually, we rolled everything back, found another provider that still offered the original model and started again from the beginning.
In theory, I had delivered three features: improved performance, better observability and support for a new model. In practice, the product had made no progress at all.
But I had produced a shit ton of code and complexity.
Had I made one change at a time, as I would have before agents made simultaneous development feel cheap, each decision would have remained easier to evaluate and reverse. I could have discarded the failed optimization without losing the observability work, or rebased one improvement after restoring the model.
Instead, the changes became entangled. Abandoning one felt like abandoning all of them. I kept believing that a much better solution was only a few days away; admitting otherwise meant throwing away a month’s apparent progress.
When producing code becomes cheap, preserving causality becomes valuable.
What I need from deep work now
These experiences point to four principles for doing knowledge work with agents.
One coherent goal at a time. I need to understand its purpose, boundaries and criteria for success, as well as how the agents’ work fits together. If I need three windows to reconstruct what is happening, coherence is already gone.
I need to notice not only the work, but the state from which I am doing it. Am I curious, anxious, excited or exhausted? Awareness keeps concentration from becoming tunnel vision and excitement from becoming fragmentation.
Changes must remain independently testable and reversible. If prompts, models, instrumentation and optimizations change together, I can no longer tell what caused what. One change at a time may look slower; it keeps learning possible.
Less is more. Agents make it almost effortless to add another feature, abstraction or option. My job is to decide which possibilities deserve to exist and what is worth building.
Deep work in the age of agents is sustained engagement with one coherent goal, guided by awareness, structured for reversibility and disciplined by restraint.
I have started with one simple experiment: when an agent begins working, I try not to open another workstream.
Instead, I stay with the task. I reread the acceptance criteria, inspect the relevant code or simply notice my posture and breathing. I ask whether I still understand what the agent is doing and whether it still serves the original goal.
The waiting time is not a delay in the work. It is part of the work.
I have been working this way for a few weeks now. It feels slower. The ticket count suggests that it may not actually be slower, although tickets do not all speak the same language. Still, the part of me that wants to move fast does not like it.
I probably send fewer prompts and produce less code. But I understand much better what I have done. At our progress meetings every other week, I can explain the changes and the decisions behind them without reconstructing everything from agent windows.
More importantly, far fewer requests for fixes come back from users. There are fewer complaints, fewer correction cycles and

[truncated]
