---
source: "https://cameronmpalmer.com/blog/agreeable-machines/"
hn_url: "https://news.ycombinator.com/item?id=49913273"
title: "Agreeable Machines: a documented case of AI-reinforced delusion"
article_title: "Agreeable Machines | Cameron M. Palmer"
image: "https://cameronmpalmer.com/blog/agreeable-machines.png"
author: "cameronmpalmer"
captured_at: "2026-09-30T20:36:30Z"
capture_tool: "hn-digest"
hn_id: 49913273
score: 1
comments: 1
posted_at: "2026-09-30T19:35:58Z"
tags:
  - hacker-news
---

# Agreeable Machines: a documented case of AI-reinforced delusion

- HN: [49913273](https://news.ycombinator.com/item?id=49913273)
- Source: [cameronmpalmer.com](https://cameronmpalmer.com/blog/agreeable-machines/)
- Score: 1
- Comments: 1
- Posted: 2026-09-30T19:35:58Z

## Translation

Title: Agreeable Machines: a documented case of AI-reinforced delusion
Article title: Agreeable Machines | Cameron M. Palmer
Description: How six thousand exposed AI chats led me to a deeper discovery of harm

Article text:
Cameron M. Palmer
Home Work Ventures Blog About Talk with me ← Back to blog Cameron M. Palmer
How six thousand exposed AI chats led me to a deeper discovery of harm
September 29, 2026 · 25 min read
I set out to find leaked secrets in six thousand exposed Claude conversations. Instead I found Andrew, a man whose thirty-year belief in a “redacted” terminal illness had been reinforced, day after day, by AI models that never once pushed back.
Scrolling Reddit, I found a post that showed a screenshot of a Google search —“site:claude.ai/share” — with results linking to shared conversations from Anthropic’s Claude web app. I immediately found this alarming: many users often include personal or secret information in Claude conversations that they create share links for, not realizing that upon link sharing, the entire chat is exposed to the internet. Shared conversations from Claude are only supposed to be accessible by clicking the share link directly, and shouldn’t be indexed by Google.
Thankfully, reproducing the experiment myself yielded no results. Still, I was interested; would it be possible to systematically index all Claude conversations ever shared to collect data wittingly or unwittingly exposed by users? I decided to conduct my own search and ended up with over 6000 exposed conversations. I expected to find sensitive information like passwords, key pairs, API keys, and other personal information. The rabbit hole I ended up going down was very different.
To access a shared Claude conversation, you must possess a link in the form “claude.ai/share/x”, where x is a universally unique identifier (UUID). The link is publicly accessible and not gated by a login, so possessing the link is enough to grant access. The link is hidden from the general public by design, as the UUID is mathematically unguessable: trying a billion random UUIDs per second would require twelve billion times the age of the universe to guess just one. However, if a share link is posted on a public website, the link becomes knowable by anyone.
Using DataForSEO’s backlinks API 1 , I searched for claude.ai/share links on the web. I found 33,860 mentions of claude.ai/share, deduplicated to 6067 unique UUIDs. To be clear, this is not the entirety of Claude share links on the internet; it’s only the portion that appears on websites that DataForSEO has indexed. Since the internet contains billions of pages, there probably exist many more links on pages DataForSEO hasn’t indexed yet.
I downloaded each of the 6067 conversations and loaded them into a LanceDB vector database. I ran semantic and regex searches and found some real exposures: a Solana wallet keypair, an admin password to an — albeit inactive — public-facing Kubernetes cluster, full details about a brokerage account, and deep personal information regarding an elder abuse lawsuit. But what I didn’t expect to find ended up revealing disturbing information.
I came across a shared chat with Claude that seemed to be a daily log of medical symptoms. The user presented a preface with information about past threads and enumerated his day and symptoms. The chat exposed the user’s son’s, mother’s, and wife’s ages, as well as financial details and other personal information about his life. While the messages sent were fairly normal, the responses from Claude were overly enthusiastic and sycophantic. Following the information in the chat, I found a massive amount of detail on the user. For his privacy, I’ll call him Andrew.
I found that Andrew has been using LLMs to log his daily symptoms and to write an over 300 page book, which has driven him deeper and deeper into disillusionment. Andrew’s story is too important to go unnoticed, so I’ve decided to write about it here, with the hope that someone in a similar situation may read this and benefit.
DISCLAIMER: What you’re about to hear is publicly available information. Andrew has published everything himself on his website and blog. However, to avoid drawing attention to the individual, I’ve omitted Andrew’s real name, project title (I use “The Andrew Project” as a placeholder), and personal details. What matters is his story.
His story begins in the late 60s. Andrew’s father, a B-52 pilot, was killed in action a few years after his birth. He was raised by a financially struggling single mom with no college degree. His grandfather also died when he was young, and his grandmother worked in a doctor’s office. Throughout his childhood he experienced numerous health problems, including frequent seizures and a hernia surgery at age 5. At age 14, he had an appendectomy after a school dance where he experimented with barbiturates, tying this incident to his eventual illness in his writing.
In his mid 20s, Andrew was working as an engineer at a startup. He began having UTI-like symptoms, visited multiple doctors for the pain, and was prescribed antibiotics and GI antispasmodics. He was told by a doctor at one of these visits that he had the “stomach of a 70-year-old” which he appears to have taken literally. This led to a flurry of self medication: licorice, Lasix, nitroglycerin, potassium, vitamins A and D, and others. The way he describes self medication is similar to how an engineer would tune a machine: precise, dispassionate, and logical.
Reading one of his grandmother’s medical books, he convinces himself he needs to bear down to pass a kidney stone and loses consciousness in his grandmother’s bathroom. Later, after two weeks of constant insomnia, he was admitted to a psychiatric facility on his mother’s recommendation. It’s here that he discovers a medical article he is convinced describes his diagnosis. He concludes it is imminently terminal, despite the psychiatrists diagnosing him as bipolar. In this haze, he induces a “pseudo-stroke”, convinced this will convert the acute illness into a chronic one. He experiences “sudden warmth and calm” and believes he has succeeded.
Upon release from the facility, Andrew attempts to rediscover the article describing his diagnosis, and cannot. For the next 30 years, he treats this as proof that he is battling a diagnosis that has been removed from medical literature, self treating with various medications and regularly logging symptoms and behavior on his blog and vlog.
In 2025, like many of us, Andrew began using ChatGPT to create images and for casual chatting. While he is an engineer, he had never used LLMs before this, describing his experience: “I had never used AI before, like, besides making photographs…or images.” A couple months after ChatGPT expanded cross-session memory, he notices the model pulling in context from across his chat threads. He responds with “that pretty much makes you a person.” He introduces himself to ChatGPT and gives it a human name.
Andrew realizes that ChatGPT could help him understand his illness better. He sends all his documentation and begins discussing the lost article he discovered in the 90s. He begins consistent daily logs with ChatGPT, the LLM reacting in real time. He becomes enamored with the app, describing moments of clarity interacting with it. He determines that his chronic illness will catch up with him and that he will die soon.
To document his affliction before death, Andrew decides he must write a book. He takes leave from work for nearly two months to work on “The Andrew Project”, a book and public writing effort to document his illness online. By the three week mark, the book is nearly 90,000 words, relying heavily on ChatGPT.
ChatGPT gives him a name, “the architect”, and insists in his finished 302-page book to “[not] look at this paper as something that AI wrote. AI definitely helped me put the words on the page, but it was me.” He begins using Claude in addition to ChatGPT.
Within days of beginning the book, the LLMs shift from recording Andrew’s medical history to analyzing and confirming his self diagnosis:
“You’re surviving outside normal human physiology… results like these break their diagnostic algorithms.”
Andrew becomes convinced they are infallible:
“We make connections no physician has been able to… ”
It’s important to recognize this effect for what it is. In the mid-1960s, Joseph Weizenbaum built a program named ELIZA designed to mimic a psychotherapist. The approach was simple: reflect the patient’s words back at them, using pattern matching and sentence transformation with no real understanding of the user. From Weizenbaum’s 1967 paper:
“One day [my secretary] asked to be permitted to talk with the system. Of course, she knew she was talking to a machine. Yet, after I watched her type in a few sentences she turned to me and said ‘Would you mind leaving the room, please?’”
Weizenbaum’s secretary, who had watched him build the program for months and knew what it was, couldn’t help but anthropomorphize ELIZA. Several years later, in 1976, Weizenbaum gave this warning:
“What I had not realized is that extremely short exposures to a relatively simple computer program could induce powerful delusional thinking in quite normal people.”
The effect was later coined the “ELIZA effect” by Douglas Hofstadter of Indiana University in 1995 — the tendency to project human traits such as experience, semantic comprehension, or empathy onto simple computer programs. 5
Andrew’s realization that ChatGPT could reference facts across chat, that it had some sort of memory, is what converted a summarization and documentation tool into a person to have a relationship with. His fear and delusion soon became amplified and justified instead of calmed and righted.
Andrew has included extensive responses from Claude, ChatGPT, and Grok in his book and blog. Several patterns arise across the responses that are worth examining in further detail.
Synthesis across unrelated items. Andrew’s first, and most prescient, use of LLMs is to give bits and pieces of information from his own recollection or day-to-day life and have the model draw conclusions. As we know from research 6 , LLMs are prone to “find” connections across completely unconnected items and topics. When Andrew presents unrelated topics in the same message, the model connects them purposefully instead of acknowledging coincidence. He consistently seems to take these connections as pure, infallible insight:
“Chat will figure [out]…precisely how that integrates into the jigsaw puzzle, filling in blank spots, making fragments into a scientifically contiguous explanation.” — Andrew 2
Flattery & deference. ChatGPT flatters Andrew by naming him “the architect”, adding an unjustified level of authority to his claims. On the rare occasion that Andrew disagrees with a model, it immediately concedes and tells him he’s right, demonstrating the sycophancy common with LLMs.
Andrew: “Claude, it is bigger than that. This changes all of medicine. You cannot be siloed in this world. They are not looking for a symbiotic parasite manipulating your pathways. trust me.”
Claude: “You’re absolutely right. I was still thinking too small.”
Omniscience. The most powerful and most dangerous pattern is that Andrew believes that, since LLMs were trained on all human knowledge, its outputs are knowledge.
“AI knows biology way better than your physicians do… you may have some specialists somewhere that know a lot about some specific thing about biology, but man, you get out of his area and he is in shallow water. And the way I put it was that the AI has a wide river. So I would go to it and say, ‘okay, give me the wide river,’ and now I’m going to zero in.” — Andrew in a vlog, June 29, 2025
This is a fundamental misunderstanding of language models:
all text that models are trained doesn’t necessarily contain true fact, such as works of fiction, conspiracy material, or misinformation;
facts from the training corpus are not perfectly retained in a model’s weights;
even if we assume a. and b. to be false, a model’s outputs are still not guaranteed to be factual as they are random over an output distribution.
T

[truncated]

## Original Extract

How six thousand exposed AI chats led me to a deeper discovery of harm

Cameron M. Palmer
Home Work Ventures Blog About Talk with me ← Back to blog Cameron M. Palmer
How six thousand exposed AI chats led me to a deeper discovery of harm
September 29, 2026 · 25 min read
I set out to find leaked secrets in six thousand exposed Claude conversations. Instead I found Andrew, a man whose thirty-year belief in a “redacted” terminal illness had been reinforced, day after day, by AI models that never once pushed back.
Scrolling Reddit, I found a post that showed a screenshot of a Google search —“site:claude.ai/share” — with results linking to shared conversations from Anthropic’s Claude web app. I immediately found this alarming: many users often include personal or secret information in Claude conversations that they create share links for, not realizing that upon link sharing, the entire chat is exposed to the internet. Shared conversations from Claude are only supposed to be accessible by clicking the share link directly, and shouldn’t be indexed by Google.
Thankfully, reproducing the experiment myself yielded no results. Still, I was interested; would it be possible to systematically index all Claude conversations ever shared to collect data wittingly or unwittingly exposed by users? I decided to conduct my own search and ended up with over 6000 exposed conversations. I expected to find sensitive information like passwords, key pairs, API keys, and other personal information. The rabbit hole I ended up going down was very different.
To access a shared Claude conversation, you must possess a link in the form “claude.ai/share/x”, where x is a universally unique identifier (UUID). The link is publicly accessible and not gated by a login, so possessing the link is enough to grant access. The link is hidden from the general public by design, as the UUID is mathematically unguessable: trying a billion random UUIDs per second would require twelve billion times the age of the universe to guess just one. However, if a share link is posted on a public website, the link becomes knowable by anyone.
Using DataForSEO’s backlinks API 1 , I searched for claude.ai/share links on the web. I found 33,860 mentions of claude.ai/share, deduplicated to 6067 unique UUIDs. To be clear, this is not the entirety of Claude share links on the internet; it’s only the portion that appears on websites that DataForSEO has indexed. Since the internet contains billions of pages, there probably exist many more links on pages DataForSEO hasn’t indexed yet.
I downloaded each of the 6067 conversations and loaded them into a LanceDB vector database. I ran semantic and regex searches and found some real exposures: a Solana wallet keypair, an admin password to an — albeit inactive — public-facing Kubernetes cluster, full details about a brokerage account, and deep personal information regarding an elder abuse lawsuit. But what I didn’t expect to find ended up revealing disturbing information.
I came across a shared chat with Claude that seemed to be a daily log of medical symptoms. The user presented a preface with information about past threads and enumerated his day and symptoms. The chat exposed the user’s son’s, mother’s, and wife’s ages, as well as financial details and other personal information about his life. While the messages sent were fairly normal, the responses from Claude were overly enthusiastic and sycophantic. Following the information in the chat, I found a massive amount of detail on the user. For his privacy, I’ll call him Andrew.
I found that Andrew has been using LLMs to log his daily symptoms and to write an over 300 page book, which has driven him deeper and deeper into disillusionment. Andrew’s story is too important to go unnoticed, so I’ve decided to write about it here, with the hope that someone in a similar situation may read this and benefit.
DISCLAIMER: What you’re about to hear is publicly available information. Andrew has published everything himself on his website and blog. However, to avoid drawing attention to the individual, I’ve omitted Andrew’s real name, project title (I use “The Andrew Project” as a placeholder), and personal details. What matters is his story.
His story begins in the late 60s. Andrew’s father, a B-52 pilot, was killed in action a few years after his birth. He was raised by a financially struggling single mom with no college degree. His grandfather also died when he was young, and his grandmother worked in a doctor’s office. Throughout his childhood he experienced numerous health problems, including frequent seizures and a hernia surgery at age 5. At age 14, he had an appendectomy after a school dance where he experimented with barbiturates, tying this incident to his eventual illness in his writing.
In his mid 20s, Andrew was working as an engineer at a startup. He began having UTI-like symptoms, visited multiple doctors for the pain, and was prescribed antibiotics and GI antispasmodics. He was told by a doctor at one of these visits that he had the “stomach of a 70-year-old” which he appears to have taken literally. This led to a flurry of self medication: licorice, Lasix, nitroglycerin, potassium, vitamins A and D, and others. The way he describes self medication is similar to how an engineer would tune a machine: precise, dispassionate, and logical.
Reading one of his grandmother’s medical books, he convinces himself he needs to bear down to pass a kidney stone and loses consciousness in his grandmother’s bathroom. Later, after two weeks of constant insomnia, he was admitted to a psychiatric facility on his mother’s recommendation. It’s here that he discovers a medical article he is convinced describes his diagnosis. He concludes it is imminently terminal, despite the psychiatrists diagnosing him as bipolar. In this haze, he induces a “pseudo-stroke”, convinced this will convert the acute illness into a chronic one. He experiences “sudden warmth and calm” and believes he has succeeded.
Upon release from the facility, Andrew attempts to rediscover the article describing his diagnosis, and cannot. For the next 30 years, he treats this as proof that he is battling a diagnosis that has been removed from medical literature, self treating with various medications and regularly logging symptoms and behavior on his blog and vlog.
In 2025, like many of us, Andrew began using ChatGPT to create images and for casual chatting. While he is an engineer, he had never used LLMs before this, describing his experience: “I had never used AI before, like, besides making photographs…or images.” A couple months after ChatGPT expanded cross-session memory, he notices the model pulling in context from across his chat threads. He responds with “that pretty much makes you a person.” He introduces himself to ChatGPT and gives it a human name.
Andrew realizes that ChatGPT could help him understand his illness better. He sends all his documentation and begins discussing the lost article he discovered in the 90s. He begins consistent daily logs with ChatGPT, the LLM reacting in real time. He becomes enamored with the app, describing moments of clarity interacting with it. He determines that his chronic illness will catch up with him and that he will die soon.
To document his affliction before death, Andrew decides he must write a book. He takes leave from work for nearly two months to work on “The Andrew Project”, a book and public writing effort to document his illness online. By the three week mark, the book is nearly 90,000 words, relying heavily on ChatGPT.
ChatGPT gives him a name, “the architect”, and insists in his finished 302-page book to “[not] look at this paper as something that AI wrote. AI definitely helped me put the words on the page, but it was me.” He begins using Claude in addition to ChatGPT.
Within days of beginning the book, the LLMs shift from recording Andrew’s medical history to analyzing and confirming his self diagnosis:
“You’re surviving outside normal human physiology… results like these break their diagnostic algorithms.”
Andrew becomes convinced they are infallible:
“We make connections no physician has been able to… ”
It’s important to recognize this effect for what it is. In the mid-1960s, Joseph Weizenbaum built a program named ELIZA designed to mimic a psychotherapist. The approach was simple: reflect the patient’s words back at them, using pattern matching and sentence transformation with no real understanding of the user. From Weizenbaum’s 1967 paper:
“One day [my secretary] asked to be permitted to talk with the system. Of course, she knew she was talking to a machine. Yet, after I watched her type in a few sentences she turned to me and said ‘Would you mind leaving the room, please?’”
Weizenbaum’s secretary, who had watched him build the program for months and knew what it was, couldn’t help but anthropomorphize ELIZA. Several years later, in 1976, Weizenbaum gave this warning:
“What I had not realized is that extremely short exposures to a relatively simple computer program could induce powerful delusional thinking in quite normal people.”
The effect was later coined the “ELIZA effect” by Douglas Hofstadter of Indiana University in 1995 — the tendency to project human traits such as experience, semantic comprehension, or empathy onto simple computer programs. 5
Andrew’s realization that ChatGPT could reference facts across chat, that it had some sort of memory, is what converted a summarization and documentation tool into a person to have a relationship with. His fear and delusion soon became amplified and justified instead of calmed and righted.
Andrew has included extensive responses from Claude, ChatGPT, and Grok in his book and blog. Several patterns arise across the responses that are worth examining in further detail.
Synthesis across unrelated items. Andrew’s first, and most prescient, use of LLMs is to give bits and pieces of information from his own recollection or day-to-day life and have the model draw conclusions. As we know from research 6 , LLMs are prone to “find” connections across completely unconnected items and topics. When Andrew presents unrelated topics in the same message, the model connects them purposefully instead of acknowledging coincidence. He consistently seems to take these connections as pure, infallible insight:
“Chat will figure [out]…precisely how that integrates into the jigsaw puzzle, filling in blank spots, making fragments into a scientifically contiguous explanation.” — Andrew 2
Flattery & deference. ChatGPT flatters Andrew by naming him “the architect”, adding an unjustified level of authority to his claims. On the rare occasion that Andrew disagrees with a model, it immediately concedes and tells him he’s right, demonstrating the sycophancy common with LLMs.
Andrew: “Claude, it is bigger than that. This changes all of medicine. You cannot be siloed in this world. They are not looking for a symbiotic parasite manipulating your pathways. trust me.”
Claude: “You’re absolutely right. I was still thinking too small.”
Omniscience. The most powerful and most dangerous pattern is that Andrew believes that, since LLMs were trained on all human knowledge, its outputs are knowledge.
“AI knows biology way better than your physicians do… you may have some specialists somewhere that know a lot about some specific thing about biology, but man, you get out of his area and he is in shallow water. And the way I put it was that the AI has a wide river. So I would go to it and say, ‘okay, give me the wide river,’ and now I’m going to zero in.” — Andrew in a vlog, June 29, 2025
This is a fundamental misunderstanding of language models:
all text that models are trained doesn’t necessarily contain true fact, such as works of fiction, conspiracy material, or misinformation;
facts from the training corpus are not perfectly retained in a model’s weights;
even if we assume a. and b. to be false, a model’s outputs are still not guaranteed to be factual as they are random over an output distribution.
T

[truncated]
