---
source: "https://norsewanderer.com/blog/blog1.html"
hn_url: "https://news.ycombinator.com/item?id=49767647"
title: "AI is just another programming tool"
article_title: "Norse Wanderer BBS - Blog Entry 1: AI Is Just Another Programming Tool"
image: ""
author: "mheilemann"
captured_at: "2026-09-19T16:16:02Z"
capture_tool: "hn-digest"
hn_id: 49767647
score: 2
comments: 0
posted_at: "2026-09-19T16:03:19Z"
tags:
  - hacker-news
---

# AI is just another programming tool

- HN: [49767647](https://news.ycombinator.com/item?id=49767647)
- Source: [norsewanderer.com](https://norsewanderer.com/blog/blog1.html)
- Score: 2
- Comments: 0
- Posted: 2026-09-19T16:03:19Z

## Translation

Title: AI is just another programming tool
Article title: Norse Wanderer BBS - Blog Entry 1: AI Is Just Another Programming Tool
Description: Norse Wanderer BBS Blog — Entry 1: AI Is Just Another Programming Tool. Notes on AI-assisted programming from someone who

Article text:
Norse Wanderer BBS - Blog Entry 1: AI Is Just Another Programming Tool
[ BLOG INDEX ]
/> AI IS JUST ANOTHER TOOL </
Michael Heilemann, The Norseman / The Viking — Notes on AI-assisted programming from someone who's been at this for almost 50 years
There is a variation of an old argument happening around AI-assisted programming: Real programmers write their own code. I've heard a few different versions of this argument before over the years.
We moved from writing machine code by hand to writing assembly language. We built libraries so we wouldn't have to repeatedly implement the same functionality. Frameworks gave us structures for building entire classes of applications. APIs let us use systems without knowing how they were implemented.
When compilers first started to be used, there were programmers who looked down on people who used them. A programmer who wrote assembly by hand could examine exactly what the processor was going to execute, instruction by instruction. Early compilers often generated code that was larger and slower than optimized hand-written assembly, and for some applications that difference mattered.
But the industry didn't stop using compilers because they weren't perfect. Compilers got better, processors got faster, and programmers discovered that spending their time thinking about the problem was usually more valuable than spending it translating every line of that problem into machine instructions. Today, nobody seriously argues that a large modern application should be written entirely in assembly simply because assembly can produce more optimal code.
There is an important difference, a compiler doesn't decide what algorithm your program should use. AI can, and that is both its strength and its weakness.
I see AI assisted coding as another programming tool that lets us work at a higher level of abstraction. There is even a similarity in how we use them. You give a compiler your code and then wait for the result. With AI, you give it your requirements and then wait as well. It feels a little like playing chess: you make your move and then sit back and think about what will come next while your opponent makes their move.
-=[ The What, More Than the How ]=-
When I use AI to write code, I tell it what I want the program to do, rather than specifying exactly how to implement it. I describe the behavior, requirements, constraints, and result I'm looking for, and let the AI work out much of the implementation.
In some ways, it is like telling a junior programmer what you want them to accomplish, except the AI can write a tremendous amount of code very quickly.
Most of the time I don't need to tell it how to do something. That's part of what makes it useful. But just like working with a junior programmer, sometimes it gets bogged down. When that happens, I need to understand the problem well enough to figure out where it is going wrong and push it in the right direction.
I've also found that the more specific I can be, the better the results are. You can't just say, "That doesn't work. Try again," and expect a fundamentally different solution. Sometimes you need to explain exactly what the data represents, what the expected behavior is, what assumptions are wrong, and even what algorithm you want it to use.
One example comes from the recreation of my old Pirate Apple II BBS, The Norse Wanderer. I originally started the board in 1983 on an Apple II+. It began as a modified GBBS Pro system, but in 1985 three of us replaced it with our own custom BBS software. I was the Sysop and handled various board modifications.
When I recreated the BBS nearly forty years later, I wasn't starting over and designing a new BBS. I had recovered disk images of the original system, including the original software and user database, and the goal was to preserve the original system while adapting it to work in a modern environment. The backend, sysop management, and display had to be rewritten for the web environment, but the original software and data provided the foundation.
As part of that work, Claude ported the code so that when someone logs on, it displays the last five users who had logged in, along with the dates and times they called. It sounds like a simple feature, and Claude happily wrote the code, except it was wrong.
When I logged in twice in succession, my previous login showed the same date and time as my current login. My account also appeared as the first entry in the "last five callers" list when it shouldn't have been there at all. On top of that, the timestamps were being handled as Universal Time when I wanted everything displayed in Pacific Time just like the original system.
The problem wasn't that the AI couldn't write the code. It could, and it did. The problem was that it did not choose an algorithm that provided the behavior it was asked to preserve.
The original BBS had already implemented this feature correctly. During the port, the AI didn't simply reproduce the original algorithm. Instead, it kept a list of recent callers and then looked up each caller's most recent login time when displaying the list.
That seems reasonable until you consider what happens when the same person logs in twice, both entries now point to the same user's latest login time. The AI had produced code that looked perfectly reasonable, but it didn't represent the historic experience.
As a test of AI, I spent about an hour trying to coax it to solve the problem in different ways, including screenshots and descriptions of what I expected to see. The AI kept trying to fix the implementation without fixing the underlying algorithm.
When I stopped explaining the desired result and specified the algorithm, it got the code right immediately, and it worked:
Create a separate database containing the actual caller history.
When a user logs in, display records 1–5, shift records 1–4 down into positions 2–5.
Insert the current caller into position 1.
Include the actual timestamp of that login.
The time should be displayed in Pacific Time.
It is a good example of why I don't think AI eliminates the need to know how to program. I had to recognize what was wrong with the approach. The AI could write the code, but I had to understand the problem well enough to tell it what the implementation needed to be.
-=[ Finding the Real Problem ]=-
I've had a very similar experience with another project I am working on now: a tool for preserving, validating, and repairing Apple II disks.
Part of the program scrapes online Apple II archives for .dsk disk images so they can be used to identify and validate disks I'm recovering. The scraper was getting caught in what appeared to be an infinite loop. It wasn't an obvious loop where the same URL was simply being visited repeatedly. It was finding additional archive content and continuing to follow it, so the activity looked legitimate.
I initially asked the AI to optimize the program because my machine was starting to page out heavily and the program was taking forever. It happily optimized the code, but the results were minimal. The reason was that it was attacking the symptom rather than the underlying architecture. The program was trying to keep too much of its state in memory, but AI did not realize that memory was the bottleneck.
I recognized that, and I told it to use a map file on disk for its storage rather than trying to maintain the entire structure in memory. That solved the problem.
Again, the AI was perfectly capable of implementing the solution, but I had to understand enough about the problem to tell it what solution was required.
And that is why I use AI for programming: it makes me a lot more productive.
I've been programming for almost 50 years, and I have a good understanding of computer hardware, architecture, operating systems, data structures, algorithms, and how software works. What I don't want to do is learn every new programming language, framework, library, and API in detail before I can use it.
I can try my ideas much faster, describe something I want to experiment with, have a working implementation generated, run it, and see what happens, make changes, and iterate. Instead of spending weeks getting the basic plumbing in place before I can evaluate an idea, I can get there in hours.
AI doesn't remove the need for programming knowledge. It changes where that knowledge is most valuable. I still need to understand the problem, the data, the algorithms, the architecture, and the constraints. What I don't want to do myself is translate every one of those decisions into thousands of lines of language-specific code.
Like tabs vs. spaces, everyone has an opinion. Let the flames begin…
Fair winds and following seas,
The Norseman

## Original Extract

Norse Wanderer BBS Blog — Entry 1: AI Is Just Another Programming Tool. Notes on AI-assisted programming from someone who

Norse Wanderer BBS - Blog Entry 1: AI Is Just Another Programming Tool
[ BLOG INDEX ]
/> AI IS JUST ANOTHER TOOL </
Michael Heilemann, The Norseman / The Viking — Notes on AI-assisted programming from someone who's been at this for almost 50 years
There is a variation of an old argument happening around AI-assisted programming: Real programmers write their own code. I've heard a few different versions of this argument before over the years.
We moved from writing machine code by hand to writing assembly language. We built libraries so we wouldn't have to repeatedly implement the same functionality. Frameworks gave us structures for building entire classes of applications. APIs let us use systems without knowing how they were implemented.
When compilers first started to be used, there were programmers who looked down on people who used them. A programmer who wrote assembly by hand could examine exactly what the processor was going to execute, instruction by instruction. Early compilers often generated code that was larger and slower than optimized hand-written assembly, and for some applications that difference mattered.
But the industry didn't stop using compilers because they weren't perfect. Compilers got better, processors got faster, and programmers discovered that spending their time thinking about the problem was usually more valuable than spending it translating every line of that problem into machine instructions. Today, nobody seriously argues that a large modern application should be written entirely in assembly simply because assembly can produce more optimal code.
There is an important difference, a compiler doesn't decide what algorithm your program should use. AI can, and that is both its strength and its weakness.
I see AI assisted coding as another programming tool that lets us work at a higher level of abstraction. There is even a similarity in how we use them. You give a compiler your code and then wait for the result. With AI, you give it your requirements and then wait as well. It feels a little like playing chess: you make your move and then sit back and think about what will come next while your opponent makes their move.
-=[ The What, More Than the How ]=-
When I use AI to write code, I tell it what I want the program to do, rather than specifying exactly how to implement it. I describe the behavior, requirements, constraints, and result I'm looking for, and let the AI work out much of the implementation.
In some ways, it is like telling a junior programmer what you want them to accomplish, except the AI can write a tremendous amount of code very quickly.
Most of the time I don't need to tell it how to do something. That's part of what makes it useful. But just like working with a junior programmer, sometimes it gets bogged down. When that happens, I need to understand the problem well enough to figure out where it is going wrong and push it in the right direction.
I've also found that the more specific I can be, the better the results are. You can't just say, "That doesn't work. Try again," and expect a fundamentally different solution. Sometimes you need to explain exactly what the data represents, what the expected behavior is, what assumptions are wrong, and even what algorithm you want it to use.
One example comes from the recreation of my old Pirate Apple II BBS, The Norse Wanderer. I originally started the board in 1983 on an Apple II+. It began as a modified GBBS Pro system, but in 1985 three of us replaced it with our own custom BBS software. I was the Sysop and handled various board modifications.
When I recreated the BBS nearly forty years later, I wasn't starting over and designing a new BBS. I had recovered disk images of the original system, including the original software and user database, and the goal was to preserve the original system while adapting it to work in a modern environment. The backend, sysop management, and display had to be rewritten for the web environment, but the original software and data provided the foundation.
As part of that work, Claude ported the code so that when someone logs on, it displays the last five users who had logged in, along with the dates and times they called. It sounds like a simple feature, and Claude happily wrote the code, except it was wrong.
When I logged in twice in succession, my previous login showed the same date and time as my current login. My account also appeared as the first entry in the "last five callers" list when it shouldn't have been there at all. On top of that, the timestamps were being handled as Universal Time when I wanted everything displayed in Pacific Time just like the original system.
The problem wasn't that the AI couldn't write the code. It could, and it did. The problem was that it did not choose an algorithm that provided the behavior it was asked to preserve.
The original BBS had already implemented this feature correctly. During the port, the AI didn't simply reproduce the original algorithm. Instead, it kept a list of recent callers and then looked up each caller's most recent login time when displaying the list.
That seems reasonable until you consider what happens when the same person logs in twice, both entries now point to the same user's latest login time. The AI had produced code that looked perfectly reasonable, but it didn't represent the historic experience.
As a test of AI, I spent about an hour trying to coax it to solve the problem in different ways, including screenshots and descriptions of what I expected to see. The AI kept trying to fix the implementation without fixing the underlying algorithm.
When I stopped explaining the desired result and specified the algorithm, it got the code right immediately, and it worked:
Create a separate database containing the actual caller history.
When a user logs in, display records 1–5, shift records 1–4 down into positions 2–5.
Insert the current caller into position 1.
Include the actual timestamp of that login.
The time should be displayed in Pacific Time.
It is a good example of why I don't think AI eliminates the need to know how to program. I had to recognize what was wrong with the approach. The AI could write the code, but I had to understand the problem well enough to tell it what the implementation needed to be.
-=[ Finding the Real Problem ]=-
I've had a very similar experience with another project I am working on now: a tool for preserving, validating, and repairing Apple II disks.
Part of the program scrapes online Apple II archives for .dsk disk images so they can be used to identify and validate disks I'm recovering. The scraper was getting caught in what appeared to be an infinite loop. It wasn't an obvious loop where the same URL was simply being visited repeatedly. It was finding additional archive content and continuing to follow it, so the activity looked legitimate.
I initially asked the AI to optimize the program because my machine was starting to page out heavily and the program was taking forever. It happily optimized the code, but the results were minimal. The reason was that it was attacking the symptom rather than the underlying architecture. The program was trying to keep too much of its state in memory, but AI did not realize that memory was the bottleneck.
I recognized that, and I told it to use a map file on disk for its storage rather than trying to maintain the entire structure in memory. That solved the problem.
Again, the AI was perfectly capable of implementing the solution, but I had to understand enough about the problem to tell it what solution was required.
And that is why I use AI for programming: it makes me a lot more productive.
I've been programming for almost 50 years, and I have a good understanding of computer hardware, architecture, operating systems, data structures, algorithms, and how software works. What I don't want to do is learn every new programming language, framework, library, and API in detail before I can use it.
I can try my ideas much faster, describe something I want to experiment with, have a working implementation generated, run it, and see what happens, make changes, and iterate. Instead of spending weeks getting the basic plumbing in place before I can evaluate an idea, I can get there in hours.
AI doesn't remove the need for programming knowledge. It changes where that knowledge is most valuable. I still need to understand the problem, the data, the algorithms, the architecture, and the constraints. What I don't want to do myself is translate every one of those decisions into thousands of lines of language-specific code.
Like tabs vs. spaces, everyone has an opinion. Let the flames begin…
Fair winds and following seas,
The Norseman
