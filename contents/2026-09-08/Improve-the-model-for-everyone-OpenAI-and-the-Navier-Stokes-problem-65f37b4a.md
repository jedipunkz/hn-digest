---
source: "https://twitter.com/danimberman/status/2097379292367802672"
hn_url: "https://news.ycombinator.com/item?id=49615732"
title: "Improve the model for everyone: OpenAI and the Navier–Stokes problem"
article_title: "Daniel Imberman on X: \"What a Millennium Prize fight tells you about the AI tools you use every day https://t.co/segqKLfF3M\" / X"
image: "https://pbs.twimg.com/media/HRtjB6VXwAApJpB.jpg"
author: "dimberman"
captured_at: "2026-09-08T19:45:35Z"
capture_tool: "hn-digest"
hn_id: 49615732
score: 2
comments: 1
posted_at: "2026-09-08T19:35:46Z"
tags:
  - hacker-news
---

# Improve the model for everyone: OpenAI and the Navier–Stokes problem

- HN: [49615732](https://news.ycombinator.com/item?id=49615732)
- Source: [twitter.com](https://twitter.com/danimberman/status/2097379292367802672)
- Score: 2
- Comments: 1
- Posted: 2026-09-08T19:35:46Z

## Translation

Title: Improve the model for everyone: OpenAI and the Navier–Stokes problem
Article title: Daniel Imberman on X: "What a Millennium Prize fight tells you about the AI tools you use every day https://t.co/segqKLfF3M" / X
Description: Improve the model for everyone

Article text:
Daniel Imberman on X: "What a Millennium Prize fight tells you about the AI tools you use every day https://t.co/segqKLfF3M" / X Post
@danimberman What a Millennium Prize fight tells you about the AI tools you use every day Improve the model for everyone
What a Millennium Prize fight tells you about the AI tools you use every day
Disclosure: I'm a partner in trustedrouter.com, which exists because of exactly the problem this piece describes. More at the end.
In 2010, Business Insider released an article discussing instant messages sent by a then 19 year old Mark Zuckerberg to a friend in 2004. In these messages, Zuckerberg brags that Harvard students were giving him very sensitive information: emails, photos, addresses. They “trust me”, he said. “Dumb fucks.”
Twenty-two years later, NYU mathematician Tristan Buckmaster is working on a version of the Navier–Stokes problem. He's teamed up with Levent Alpöge of Anthropic, and they're attacking it through a side door that two mathematicians, Córdoba and Martínez-Zoroa, opened a few years ago. While the method is published, almost nobody works on it. It's not the kind of thing a model would just 'retrieve' from the problem statement.
This August, they get a proof for the Euler equations. On August 22 they had the proof verified in Lean (meaning a computer checked every step). Two weeks later, Buckmaster is on a call with OpenAI, and they're telling him their internal model has proven this version of Navier–Stokes using the same obscure methods he had been investigating.
When he pushes on when they started, the answer comes out: a few days ago. After word of his work had reached them.
Now, I'm not here to tell you that OpenAI actively dove into Buckmaster's chat logs and stole his research. Strong accusations require strong proof, and at the moment, no such proof exists. What I'm suggesting is something less nefarious, but more insidious. What happens to the world of unique research in a world where the tools you're using are training on your unique thought patterns?
So let’s assume that OpenAI’s engineers weren’t actively pulling chat logs, what if the model was? LLMs go through two major phases. Pre-training (scraping the whole world) and post-training (reinforcement-learning) where the models are tuned on curated examples (including, say, user sessions).
Two weeks is nowhere near enough time to use any of this information in pre-training, but post-training can turn around in days. Even though the proof was two weeks old, this project in one form or another had been sitting in Codex for the better part of a year. More than enough time for the sessions to get scooped into the model.
In his chats with OpenAI they explicitly denied looking up any user data, but when Buckmaster asks about training, he never gets an answer. OpenAI's own help page says content from individual ChatGPT and Codex accounts may be used to train models unless you opt out. Buckmaster was on a personal account, paying out of his own research funds.
Now think of the start-up you’re building, the research you’re doing, the book that you’re writing. How certain are you that when you use AI to co-author your work, that what you are building stays exclusively yours.
According to Buckmaster, OpenAI said that if they published, they’d say publicly that he and Alpöge deserve to receive the Clay Prize as the “closest humans to the problem.” Separately, they proposed that he write up the Navier–Stokes result alone (removing Alpöge), and credit OpenAI's model. When he said he'd go public, they replied "Why would you ruin your career?” When Buckmaster pushed back they said “If you don't want me to be nice, then I don't have to be nice."
Sébastien Bubeck, who leads OpenAI's math effort, has called the allegations "false and inflammatory" and says he came into the discussion following academic norms. Buckmaster, for his part, says plainly he's not accusing anyone of anything. He hasn't seen their proof and doesn't know if his data was used. He's stating what he was told, and when.
Ultimately, I would be extremely wary of what we give to major corporations. The systems in the US are built in a way where punishments and fines are insufficient to act as true deterrents. When Anthropic paid a 1.5 billion dollar settlement for pirating books in their training, that is a line item in the P&L.
Just because nobody looked at your data, doesn’t mean your data wasn’t used.
AI tooling has become an essential element of our workflows, yet like early Facebook users we seem to have forgotten just how valuable our data actually is.
Part of the response has been to route around the consumer products entirely. Companies are now self-hosting, or going through the API, where the terms are different. But a router is just another party that sees your prompts. So now the question isn't whether it's encrypted. It's whether it logs, how long it keeps what it logs, and whether the endpoint on the other side is one that trains.
trustedrouter.com is something my good friend Joseph Perla built to fix exactly this. It's a drop-in router. It has the same OpenAI SDK, with equal or cheaper prices than the other routers. The difference is what happens to your prompt once it hits the box.
The gateway runs inside a trusted execution environment and doesn't log prompts or outputs. Not even the people running it can read your requests. And I'm not asking you to take that on faith, because that's the whole point of this article. The code is public, and the running gateway is attested against it, so you can check that the thing serving your traffic is the thing you read. If you want the model on the other end to be zero-retention too, trustedrouter/zdr only routes to providers that commit to it.
Does this make your data "safe"? No. Nothing does. But at least you're back to the deal you thought you had.
Which is more than Buckmaster got.
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
© 2026 X Corp. Daniel Imberman @danimberman What a Millennium Prize fight tells you about the AI tools you use every day Improve the model for everyone
What a Millennium Prize fight tells you about the AI tools you use every day
Disclosure: I'm a partner in trustedrouter.com, which exists because of exactly the problem this piece describes. More at the end.
In 2010, Business Insider released an article discussing instant messages sent by a then 19 year old Mark Zuckerberg to a friend in 2004. In these messages, Zuckerberg brags that Harvard students were giving him very sensitive information: emails, photos, addresses. They “trust me”, he said. “Dumb fucks.”
Twenty-two years later, NYU mathematician Tristan Buckmaster is working on a version of the Navier–Stokes problem. He's teamed up with Levent Alpöge of Anthropic, and they're attacking it through a side door that two mathematicians, Córdoba and Martínez-Zoroa, opened a few years ago. While the method is published, almost nobody works on it. It's not the kind of thing a model would just 'retrieve' from the problem statement.
This August, they get a proof for the Euler equations. On August 22 they had the proof verified in Lean (meaning a computer checked every step). Two weeks later, Buckmaster is on a call with OpenAI, and they're telling him their internal model has proven this version of Navier–Stokes using the same obscure methods he had been investigating.
When he pushes on when they started, the answer comes out: a few days ago. After word of his work had reached them.
Now, I'm not here to tell you that OpenAI actively dove into Buckmaster's chat logs and stole his research. Strong accusations require strong proof, and at the moment, no such proof exists. What I'm suggesting is something less nefarious, but more insidious. What happens to the world of unique research in a world where the tools you're using are training on your unique thought patterns?
So let’s assume that OpenAI’s engineers weren’t actively pulling chat logs, what if the model was? LLMs go through two major phases. Pre-training (scraping the whole world) and post-training (reinforcement-learning) where the models are tuned on curated examples (including, say, user sessions).
Two weeks is nowhere near enough time to use any of this information in pre-training, but post-training can turn around in days. Even though the proof was two weeks old, this project in one form or another had been sitting in Codex for the better part of a year. More than enough time for the sessions to get scooped into the model.
In his chats with OpenAI they explicitly denied looking up any user data, but when Buckmaster asks about training, he never gets an answer. OpenAI's own help page says content from individual ChatGPT and Codex accounts may be used to train models unless you opt out. Buckmaster was on a personal account, paying out of his own research funds.
Now think of the start-up you’re building, the research you’re doing, the book that you’re writing. How certain are you that when you use AI to co-author your work, that what you are building stays exclusively yours.
According to Buckmaster, OpenAI said that if they published, they’d say publicly that he and Alpöge deserve to receive the Clay Prize as the “closest humans to the problem.” Separately, they proposed that he write up the Navier–Stokes result alone (removing Alpöge), and credit OpenAI's model. When he said he'd go public, they replied "Why would you ruin your career?” When Buckmaster pushed back they said “If you don't want me to be nice, then I don't have to be nice."
Sébastien Bubeck, who leads OpenAI's math effort, has called the allegations "false and inflammatory" and says he came into the discussion following academic norms. Buckmaster, for his part, says plainly he's not accusing anyone of anything. He hasn't seen their proof and doesn't know if his data was used. He's stating what he was told, and when.
Ultimately, I would be extremely wary of what we give to major corporations. The systems in the US are built in a way where punishments and fines are insufficient to act as true deterrents. When Anthropic paid a 1.5 billion dollar settlement for pirating books in their training, that is a line item in the P&L.
Just because nobody looked at your data, doesn’t mean your data wasn’t used.
AI tooling has become an essential element of our workflows, yet like early Facebook users we seem to have forgotten just how valuable our data actually is.
Part of the response has been to route around the consumer products entirely. Companies are now self-hosting, or going through the API, where the terms are different. But a router is just another party that sees your prompts. So now the question isn't whether it's encrypted. It's whether it logs, how long it keeps what it logs, and whether the endpoint on the other side is one that trains.
trustedrouter.com is something my good friend Joseph Perla built to fix exactly this. It's a drop-in router. It has the same OpenAI SDK, with equal or cheaper prices than the other routers. The difference is what happens to your prompt once it hits the box.
The gateway runs inside a trusted execution environment and doesn't log prompts or outputs. Not even the people running it can read your requests. And I'm not asking you to take that on faith, because that's the whole point of this article. The code is public, and the running gateway is attested against it, so you can check that the thing serving your traffic is the thing you read. If you want the model on the other end to be zero-retention too, trustedrouter/zdr only routes to providers that commit to it.
Does this make your data "safe"? No. Nothing does. But at least you're back to the deal you thought you had.
Which is more than Buckmaster got.

## Original Extract

Improve the model for everyone

Daniel Imberman on X: "What a Millennium Prize fight tells you about the AI tools you use every day https://t.co/segqKLfF3M" / X Post
@danimberman What a Millennium Prize fight tells you about the AI tools you use every day Improve the model for everyone
What a Millennium Prize fight tells you about the AI tools you use every day
Disclosure: I'm a partner in trustedrouter.com, which exists because of exactly the problem this piece describes. More at the end.
In 2010, Business Insider released an article discussing instant messages sent by a then 19 year old Mark Zuckerberg to a friend in 2004. In these messages, Zuckerberg brags that Harvard students were giving him very sensitive information: emails, photos, addresses. They “trust me”, he said. “Dumb fucks.”
Twenty-two years later, NYU mathematician Tristan Buckmaster is working on a version of the Navier–Stokes problem. He's teamed up with Levent Alpöge of Anthropic, and they're attacking it through a side door that two mathematicians, Córdoba and Martínez-Zoroa, opened a few years ago. While the method is published, almost nobody works on it. It's not the kind of thing a model would just 'retrieve' from the problem statement.
This August, they get a proof for the Euler equations. On August 22 they had the proof verified in Lean (meaning a computer checked every step). Two weeks later, Buckmaster is on a call with OpenAI, and they're telling him their internal model has proven this version of Navier–Stokes using the same obscure methods he had been investigating.
When he pushes on when they started, the answer comes out: a few days ago. After word of his work had reached them.
Now, I'm not here to tell you that OpenAI actively dove into Buckmaster's chat logs and stole his research. Strong accusations require strong proof, and at the moment, no such proof exists. What I'm suggesting is something less nefarious, but more insidious. What happens to the world of unique research in a world where the tools you're using are training on your unique thought patterns?
So let’s assume that OpenAI’s engineers weren’t actively pulling chat logs, what if the model was? LLMs go through two major phases. Pre-training (scraping the whole world) and post-training (reinforcement-learning) where the models are tuned on curated examples (including, say, user sessions).
Two weeks is nowhere near enough time to use any of this information in pre-training, but post-training can turn around in days. Even though the proof was two weeks old, this project in one form or another had been sitting in Codex for the better part of a year. More than enough time for the sessions to get scooped into the model.
In his chats with OpenAI they explicitly denied looking up any user data, but when Buckmaster asks about training, he never gets an answer. OpenAI's own help page says content from individual ChatGPT and Codex accounts may be used to train models unless you opt out. Buckmaster was on a personal account, paying out of his own research funds.
Now think of the start-up you’re building, the research you’re doing, the book that you’re writing. How certain are you that when you use AI to co-author your work, that what you are building stays exclusively yours.
According to Buckmaster, OpenAI said that if they published, they’d say publicly that he and Alpöge deserve to receive the Clay Prize as the “closest humans to the problem.” Separately, they proposed that he write up the Navier–Stokes result alone (removing Alpöge), and credit OpenAI's model. When he said he'd go public, they replied "Why would you ruin your career?” When Buckmaster pushed back they said “If you don't want me to be nice, then I don't have to be nice."
Sébastien Bubeck, who leads OpenAI's math effort, has called the allegations "false and inflammatory" and says he came into the discussion following academic norms. Buckmaster, for his part, says plainly he's not accusing anyone of anything. He hasn't seen their proof and doesn't know if his data was used. He's stating what he was told, and when.
Ultimately, I would be extremely wary of what we give to major corporations. The systems in the US are built in a way where punishments and fines are insufficient to act as true deterrents. When Anthropic paid a 1.5 billion dollar settlement for pirating books in their training, that is a line item in the P&L.
Just because nobody looked at your data, doesn’t mean your data wasn’t used.
AI tooling has become an essential element of our workflows, yet like early Facebook users we seem to have forgotten just how valuable our data actually is.
Part of the response has been to route around the consumer products entirely. Companies are now self-hosting, or going through the API, where the terms are different. But a router is just another party that sees your prompts. So now the question isn't whether it's encrypted. It's whether it logs, how long it keeps what it logs, and whether the endpoint on the other side is one that trains.
trustedrouter.com is something my good friend Joseph Perla built to fix exactly this. It's a drop-in router. It has the same OpenAI SDK, with equal or cheaper prices than the other routers. The difference is what happens to your prompt once it hits the box.
The gateway runs inside a trusted execution environment and doesn't log prompts or outputs. Not even the people running it can read your requests. And I'm not asking you to take that on faith, because that's the whole point of this article. The code is public, and the running gateway is attested against it, so you can check that the thing serving your traffic is the thing you read. If you want the model on the other end to be zero-retention too, trustedrouter/zdr only routes to providers that commit to it.
Does this make your data "safe"? No. Nothing does. But at least you're back to the deal you thought you had.
Which is more than Buckmaster got.
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
© 2026 X Corp. Daniel Imberman @danimberman What a Millennium Prize fight tells you about the AI tools you use every day Improve the model for everyone
What a Millennium Prize fight tells you about the AI tools you use every day
Disclosure: I'm a partner in trustedrouter.com, which exists because of exactly the problem this piece describes. More at the end.
In 2010, Business Insider released an article discussing instant messages sent by a then 19 year old Mark Zuckerberg to a friend in 2004. In these messages, Zuckerberg brags that Harvard students were giving him very sensitive information: emails, photos, addresses. They “trust me”, he said. “Dumb fucks.”
Twenty-two years later, NYU mathematician Tristan Buckmaster is working on a version of the Navier–Stokes problem. He's teamed up with Levent Alpöge of Anthropic, and they're attacking it through a side door that two mathematicians, Córdoba and Martínez-Zoroa, opened a few years ago. While the method is published, almost nobody works on it. It's not the kind of thing a model would just 'retrieve' from the problem statement.
This August, they get a proof for the Euler equations. On August 22 they had the proof verified in Lean (meaning a computer checked every step). Two weeks later, Buckmaster is on a call with OpenAI, and they're telling him their internal model has proven this version of Navier–Stokes using the same obscure methods he had been investigating.
When he pushes on when they started, the answer comes out: a few days ago. After word of his work had reached them.
Now, I'm not here to tell you that OpenAI actively dove into Buckmaster's chat logs and stole his research. Strong accusations require strong proof, and at the moment, no such proof exists. What I'm suggesting is something less nefarious, but more insidious. What happens to the world of unique research in a world where the tools you're using are training on your unique thought patterns?
So let’s assume that OpenAI’s engineers weren’t actively pulling chat logs, what if the model was? LLMs go through two major phases. Pre-training (scraping the whole world) and post-training (reinforcement-learning) where the models are tuned on curated examples (including, say, user sessions).
Two weeks is nowhere near enough time to use any of this information in pre-training, but post-training can turn around in days. Even though the proof was two weeks old, this project in one form or another had been sitting in Codex for the better part of a year. More than enough time for the sessions to get scooped into the model.
In his chats with OpenAI they explicitly denied looking up any user data, but when Buckmaster asks about training, he never gets an answer. OpenAI's own help page says content from individual ChatGPT and Codex accounts may be used to train models unless you opt out. Buckmaster was on a personal account, paying out of his own research funds.
Now think of the start-up you’re building, the research you’re doing, the book that you’re writing. How certain are you that when you use AI to co-author your work, that what you are building stays exclusively yours.
According to Buckmaster, OpenAI said that if they published, they’d say publicly that he and Alpöge deserve to receive the Clay Prize as the “closest humans to the problem.” Separately, they proposed that he write up the Navier–Stokes result alone (removing Alpöge), and credit OpenAI's model. When he said he'd go public, they replied "Why would you ruin your career?” When Buckmaster pushed back they said “If you don't want me to be nice, then I don't have to be nice."
Sébastien Bubeck, who leads OpenAI's math effort, has called the allegations "false and inflammatory" and says he came into the discussion following academic norms. Buckmaster, for his part, says plainly he's not accusing anyone of anything. He hasn't seen their proof and doesn't know if his data was used. He's stating what he was told, and when.
Ultimately, I would be extremely wary of what we give to major corporations. The systems in the US are built in a way where punishments and fines are insufficient to act as true deterrents. When Anthropic paid a 1.5 billion dollar settlement for pirating books in their training, that is a line item in the P&L.
Just because nobody looked at your data, doesn’t mean your data wasn’t used.
AI tooling has become an essential element of our workflows, yet like early Facebook users we seem to have forgotten just how valuable our data actually is.
Part of the response has been to route around the consumer products entirely. Companies are now self-hosting, or going through the API, where the terms are different. But a router is just another party that sees your prompts. So now the question isn't whether it's encrypted. It's whether it logs, how long it keeps what it logs, and whether the endpoint on the other side is one that trains.
trustedrouter.com is something my good friend Joseph Perla built to fix exactly this. It's a drop-in router. It has the same OpenAI SDK, with equal or cheaper prices than the other routers. The difference is what happens to your prompt once it hits the box.
The gateway runs inside a trusted execution environment and doesn't log prompts or outputs. Not even the people running it can read your requests. And I'm not asking you to take that on faith, because that's the whole point of this article. The code is public, and the running gateway is attested against it, so you can check that the thing serving your traffic is the thing you read. If you want the model on the other end to be zero-retention too, trustedrouter/zdr only routes to providers that commit to it.
Does this make your data "safe"? No. Nothing does. But at least you're back to the deal you thought you had.
Which is more than Buckmaster got.
