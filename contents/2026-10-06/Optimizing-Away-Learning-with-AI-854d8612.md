---
source: "https://hilalmutlu.com/blog/optimizing-away-learning-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=49982693"
title: "Optimizing Away Learning with AI"
article_title: "optimizing away learning with AI — hilal mutlu"
image: ""
author: "siberpunk"
captured_at: "2026-10-06T19:23:28Z"
capture_tool: "hn-digest"
hn_id: 49982693
score: 1
comments: 0
posted_at: "2026-10-06T19:11:01Z"
tags:
  - hacker-news
---

# Optimizing Away Learning with AI

- HN: [49982693](https://news.ycombinator.com/item?id=49982693)
- Source: [hilalmutlu.com](https://hilalmutlu.com/blog/optimizing-away-learning-with-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-10-06T19:11:01Z

## Translation

Title: Optimizing Away Learning with AI
Article title: optimizing away learning with AI — hilal mutlu
Description: What skills like abstraction, indirect learning, recall, and sitting with uncertainty we might be losing when we learn through AI.

Article text:
optimizing away learning with AI — hilal mutlu Skip to content hilal mutlu home
optimizing away learning with AI
Deep thinking and staying with uncertainty
Recently, I was doomscrolling like always and came across this tweet in which a father explains the app he vibe coded for his child. The app is basically the math version of Fruit Ninja. Apparently his son is struggling with multiplication table (oh the good old days) and his father comes up with this solution.
Like anything on the internet, this idea brought so many counter-ideas. I agree with a significant portion of them, at least for this specific scenario. One user argued that a child builds focus through classes and homework and having them exposed too many triggers (which they already suffer too much) is wrong. While I fundamentally agree with the idea, it made me think about how my learning systems changed and if they’re still working.
“Let’s go back 70 million years”
The main idea I was thinking about is that people have been able to learn things for centuries without various tools that make things easier. All the knowledge we have managed to produce and store to this day came entirely through manual human effort.
Actually there is no need to go back centuries, even just comparing the theoretical learning behaviors of a computer science student from 10 years ago with one today would give us a long list.
Let’s think about the scenario of writing a piece of code. A student 10 years ago would understand what they were trying to do, try to do it, and when they couldn’t, they would search for it on the internet or ask someone. There were even ways to ask a question properly . Today’s student, - if they are still coding at all- will directly ask an LLM for the code snippet that does the trick.
Similarly, if there was a technical book that needed to be read, I imagine that a student 10 years ago would get the book, read either all of it or the necessary sections, understand it, and form their own questions.
Now let’s examine today’s student:
She first gets the latest technical book she wants to read (Designing Data Intensive Applications Second Edition) . Then she also gets the PDF. She uploads this PDF to NotebookLM, generates a summary and an objective list for each chapter.
Then she starts reading one of the chapters. While doing so, she directly asks the LLM about a sentence that catches her attention. While doing this, does not worry about turning the question into something understandable and abstracted away from its context, because the LLM already has the entire context. She does not even try to turn the question into a meaningful human sentence. Types it using the shortest possible phrases. While waiting for the answer to load, she doesn’t sit idle, maybe reads one more sentence from the book. (Or worse, reads tweets.) She doesn’t understand that next sentence either, so asks another question. She does not understand something
[truncated]
I believe this way of working and learning by talking to AI services is taking too many skills away from us. It may even be preventing those who started directly with this revolution from ever acquiring these skills in the first place.
Computer science is a gigantic abstraction problem made up of smaller abstraction problems. The operating system abstracts the hardware, applications abstract the operating system. User interfaces abstract the complex commands given to operating systems.
Computer science requires being able to take a problem and deal with it independently from its context, and this is one of the greatest abilities of computer scientists.
The question “Why is retrieving the last 30 days of data from this table getting slower over time?” should be transformable into the problem “How can I efficiently query a growing dataset by time range?”
The question “How do I delete a task while building a todo app?” should be abstracted into “How do I remove an item from an array?”
The question “When two users add the last remaining product to their carts at the same time, one of them cannot buy it” should be abstracted into “race condition in concurrent access.”
When asking a real person or a search engine a question, we have to make the actual problem independent from its context, abstract it, and then ask it.
Because the LLM already has access to the entire context, we can give it whatever comes to mind without abstracting it first. Our ability to isolate problems, and our problem definition skill, which is the first step of problem solving in the scientific method, are becoming weaker in this way.
The problem is not only that we are starting to ask worse questions. We are using our ability to define a problem and abstract it correctly less often. Just as we started forgetting how to read maps with the spread of navigation systems, we may be forgetting this skill over time as well.
In the classical method, while researching a problem, we read documentation, browse ancient Stack Overflow pages, and read blogs. Even if all these steps are not directly related to that exact problem, they allow us to learn additional things.
For example, what we are looking for might be why a PostgreSQL query is slow, but while researching it we may also come across concepts such as execution plans, index types, or cardinality that we do not directly need or care at that moment, and we gain bits of knowledge that later make us say, “I remember seeing something about this somewhere.”
When we ask artificial intelligence for the answer, however, it takes these steps on our behalf and gives us the answer directly. We even provide extra system prompts so that it does not go into unrelated topics and gives us short and clear answers. This greatly reduces our chance of indirect learning.
Learning something is not only about seeing the answer, but also about trying to retrieve it from your mind. Now, the moment the question “I knew this, what was it again?” comes up, we turn to AI. That rarely used command we need to type in the terminal, the difference between two technical options… We are eliminating the process of trying to pull information from long-term memory and bring it back to the front.
As access to information becomes easier, keeping information in the mind starts to become unnecessary for the brain. I swear I have read blog posts about studies on this, but because I thought I could access them very easily, I did not bother keeping them in mind.
Point proven ;)
Unfortunately, knowing where to access information does not provide us with the same abilities as actually possessing that information in daily life or professional life. You experience this painful experience when a small piece of information is asked quickly in a meeting..
On Stack Overflow, we used to at least check the date of the answer before copy-pasting it. A very old answer, outdated documentation, a sentence written by a throwaway account, and so on.. We had certain criteria regarding the reliability of information, and when these bare minimum requirements were not met, skepticism kicked in and we looked for a second opinion. We tried to verify the information.
With artificial intelligence, on the other hand, we tend to accept everything it writes as correct. When I talk to my friends in technical circles, at least I hear this used in sentences: “AI says this, of course we should verify it from the official documentation,” whereas in non-technical circles, I observe an approach more like, “I learned this from AI, of course it is correct, this is the one and only truth.”
When we overlook the fact that the language models we use are, in a sense, commercial products, we may be heading toward an extremely dangerous point. Topics such as questioning the source of information and verifying its accuracy, in my opinion, deserve to be examined under a much broader heading than this relatively innocent topic of learning.
Deep thinking and staying with uncertainty
There is a 2013 drawing by Jason Heeris called “This Is Why You Shouldn’t Interrupt a Programmer.” It depicts the mental structure we build in our minds by stacking small pieces on top of one another, and how a single interruption can bring that structure down. AI may not be destroying that structure, but it is also not allowing us to build it.
We no longer allow our brains not to understand something. We do not wait for a paragraph to make sense a few pages later, for a bug to keep turning around in our heads, or for the missing pieces to come together on their own. We ask about every detail and instantly get an explanation like we were five, a real-world example, and even the chain of reasoning that we should have built ourselves. We load more information at once than our minds can handle.
In fact, staying with uncertainty was not the moment when learning failed. It was the moment when the mental model was being built. When we fill every gap instantly, we may seem to be saving time, but we may also be skipping the part where the brain builds the connections on its own, in other words, the part where it learns .
None of this means that I am arguing against using artificial intelligence while learning. As I mentioned, we have an extremely powerful teacher in our hands. Maybe the solution is to actually treat it like a teacher. To follow the steps we would normally take before asking a real teacher. Not every friction in the learning process is unnecessary; while getting rid of some of them, we may actually be getting rid of learning itself.
We are succeeding at training LLMs, but we also need to succeed at not making learning increasingly complicated and difficult in a world where LLMs exist.
Thanks for reading. If you have any feedback or would like to discuss further, I would be happy to hear from you. You can reach me through contact or leave a comment below.

## Original Extract

What skills like abstraction, indirect learning, recall, and sitting with uncertainty we might be losing when we learn through AI.

optimizing away learning with AI — hilal mutlu Skip to content hilal mutlu home
optimizing away learning with AI
Deep thinking and staying with uncertainty
Recently, I was doomscrolling like always and came across this tweet in which a father explains the app he vibe coded for his child. The app is basically the math version of Fruit Ninja. Apparently his son is struggling with multiplication table (oh the good old days) and his father comes up with this solution.
Like anything on the internet, this idea brought so many counter-ideas. I agree with a significant portion of them, at least for this specific scenario. One user argued that a child builds focus through classes and homework and having them exposed too many triggers (which they already suffer too much) is wrong. While I fundamentally agree with the idea, it made me think about how my learning systems changed and if they’re still working.
“Let’s go back 70 million years”
The main idea I was thinking about is that people have been able to learn things for centuries without various tools that make things easier. All the knowledge we have managed to produce and store to this day came entirely through manual human effort.
Actually there is no need to go back centuries, even just comparing the theoretical learning behaviors of a computer science student from 10 years ago with one today would give us a long list.
Let’s think about the scenario of writing a piece of code. A student 10 years ago would understand what they were trying to do, try to do it, and when they couldn’t, they would search for it on the internet or ask someone. There were even ways to ask a question properly . Today’s student, - if they are still coding at all- will directly ask an LLM for the code snippet that does the trick.
Similarly, if there was a technical book that needed to be read, I imagine that a student 10 years ago would get the book, read either all of it or the necessary sections, understand it, and form their own questions.
Now let’s examine today’s student:
She first gets the latest technical book she wants to read (Designing Data Intensive Applications Second Edition) . Then she also gets the PDF. She uploads this PDF to NotebookLM, generates a summary and an objective list for each chapter.
Then she starts reading one of the chapters. While doing so, she directly asks the LLM about a sentence that catches her attention. While doing this, does not worry about turning the question into something understandable and abstracted away from its context, because the LLM already has the entire context. She does not even try to turn the question into a meaningful human sentence. Types it using the shortest possible phrases. While waiting for the answer to load, she doesn’t sit idle, maybe reads one more sentence from the book. (Or worse, reads tweets.) She doesn’t understand that next sentence either, so asks another question. She does not understand something
[truncated]
I believe this way of working and learning by talking to AI services is taking too many skills away from us. It may even be preventing those who started directly with this revolution from ever acquiring these skills in the first place.
Computer science is a gigantic abstraction problem made up of smaller abstraction problems. The operating system abstracts the hardware, applications abstract the operating system. User interfaces abstract the complex commands given to operating systems.
Computer science requires being able to take a problem and deal with it independently from its context, and this is one of the greatest abilities of computer scientists.
The question “Why is retrieving the last 30 days of data from this table getting slower over time?” should be transformable into the problem “How can I efficiently query a growing dataset by time range?”
The question “How do I delete a task while building a todo app?” should be abstracted into “How do I remove an item from an array?”
The question “When two users add the last remaining product to their carts at the same time, one of them cannot buy it” should be abstracted into “race condition in concurrent access.”
When asking a real person or a search engine a question, we have to make the actual problem independent from its context, abstract it, and then ask it.
Because the LLM already has access to the entire context, we can give it whatever comes to mind without abstracting it first. Our ability to isolate problems, and our problem definition skill, which is the first step of problem solving in the scientific method, are becoming weaker in this way.
The problem is not only that we are starting to ask worse questions. We are using our ability to define a problem and abstract it correctly less often. Just as we started forgetting how to read maps with the spread of navigation systems, we may be forgetting this skill over time as well.
In the classical method, while researching a problem, we read documentation, browse ancient Stack Overflow pages, and read blogs. Even if all these steps are not directly related to that exact problem, they allow us to learn additional things.
For example, what we are looking for might be why a PostgreSQL query is slow, but while researching it we may also come across concepts such as execution plans, index types, or cardinality that we do not directly need or care at that moment, and we gain bits of knowledge that later make us say, “I remember seeing something about this somewhere.”
When we ask artificial intelligence for the answer, however, it takes these steps on our behalf and gives us the answer directly. We even provide extra system prompts so that it does not go into unrelated topics and gives us short and clear answers. This greatly reduces our chance of indirect learning.
Learning something is not only about seeing the answer, but also about trying to retrieve it from your mind. Now, the moment the question “I knew this, what was it again?” comes up, we turn to AI. That rarely used command we need to type in the terminal, the difference between two technical options… We are eliminating the process of trying to pull information from long-term memory and bring it back to the front.
As access to information becomes easier, keeping information in the mind starts to become unnecessary for the brain. I swear I have read blog posts about studies on this, but because I thought I could access them very easily, I did not bother keeping them in mind.
Point proven ;)
Unfortunately, knowing where to access information does not provide us with the same abilities as actually possessing that information in daily life or professional life. You experience this painful experience when a small piece of information is asked quickly in a meeting..
On Stack Overflow, we used to at least check the date of the answer before copy-pasting it. A very old answer, outdated documentation, a sentence written by a throwaway account, and so on.. We had certain criteria regarding the reliability of information, and when these bare minimum requirements were not met, skepticism kicked in and we looked for a second opinion. We tried to verify the information.
With artificial intelligence, on the other hand, we tend to accept everything it writes as correct. When I talk to my friends in technical circles, at least I hear this used in sentences: “AI says this, of course we should verify it from the official documentation,” whereas in non-technical circles, I observe an approach more like, “I learned this from AI, of course it is correct, this is the one and only truth.”
When we overlook the fact that the language models we use are, in a sense, commercial products, we may be heading toward an extremely dangerous point. Topics such as questioning the source of information and verifying its accuracy, in my opinion, deserve to be examined under a much broader heading than this relatively innocent topic of learning.
Deep thinking and staying with uncertainty
There is a 2013 drawing by Jason Heeris called “This Is Why You Shouldn’t Interrupt a Programmer.” It depicts the mental structure we build in our minds by stacking small pieces on top of one another, and how a single interruption can bring that structure down. AI may not be destroying that structure, but it is also not allowing us to build it.
We no longer allow our brains not to understand something. We do not wait for a paragraph to make sense a few pages later, for a bug to keep turning around in our heads, or for the missing pieces to come together on their own. We ask about every detail and instantly get an explanation like we were five, a real-world example, and even the chain of reasoning that we should have built ourselves. We load more information at once than our minds can handle.
In fact, staying with uncertainty was not the moment when learning failed. It was the moment when the mental model was being built. When we fill every gap instantly, we may seem to be saving time, but we may also be skipping the part where the brain builds the connections on its own, in other words, the part where it learns .
None of this means that I am arguing against using artificial intelligence while learning. As I mentioned, we have an extremely powerful teacher in our hands. Maybe the solution is to actually treat it like a teacher. To follow the steps we would normally take before asking a real teacher. Not every friction in the learning process is unnecessary; while getting rid of some of them, we may actually be getting rid of learning itself.
We are succeeding at training LLMs, but we also need to succeed at not making learning increasingly complicated and difficult in a world where LLMs exist.
Thanks for reading. If you have any feedback or would like to discuss further, I would be happy to hear from you. You can reach me through contact or leave a comment below.
