---
source: "https://www.gsb.stanford.edu/insights/blueprint-keeping-humans-control-ai"
hn_url: "https://news.ycombinator.com/item?id=49708836"
title: "A Blueprint for Keeping Humans in Control of AI"
article_title: "A Blueprint for Keeping Humans in Control of AI | Stanford Graduate School of Business"
image: "https://www.gsb.stanford.edu/sites/default/files/styles/og_image/public/2026-09/Bayati%20AI%20Agent%20Oversight%20Key%20Image.jpg.jpeg?itok=8Ejuth9q"
author: "geox"
captured_at: "2026-09-15T07:15:57Z"
capture_tool: "hn-digest"
hn_id: 49708836
score: 1
comments: 0
posted_at: "2026-09-15T07:08:38Z"
tags:
  - hacker-news
---

# A Blueprint for Keeping Humans in Control of AI

- HN: [49708836](https://news.ycombinator.com/item?id=49708836)
- Source: [www.gsb.stanford.edu](https://www.gsb.stanford.edu/insights/blueprint-keeping-humans-control-ai)
- Score: 1
- Comments: 0
- Posted: 2026-09-15T07:08:38Z

## Translation

Title: A Blueprint for Keeping Humans in Control of AI
Article title: A Blueprint for Keeping Humans in Control of AI | Stanford Graduate School of Business
Description: New research teaches independent AI agents to ask for help — and builds oversight of models we don’t fully trust.

Article text:
A Blueprint for Keeping Humans in Control of AI | Stanford Graduate School of Business
Skip to main content
Menu
Enter the terms you wish to search for.
Insights by Stanford Business
Topics
Accounting
A Blueprint for Keeping Humans in Control of AI
New research teaches independent AI agents to ask for help — and builds oversight of models we don’t fully trust.
Stanford GSB researchers are building frameworks that keep AI under human control even as they become more autonomous.
In a cooperative game, an AI agent learns when to check in with a person and when to act alone.
Another framework sets up effective oversight of an untrusted model by pools of human or AI supervisors.
William Overman began his PhD program at Stanford Graduate School of Business in an auspicious moment: just two months before ChatGPT publicly launched in November 2022, exploding the widely held understandings of what machines are capable of.
Even as Overman began enlisting AI for his research, he grew wary of where the technology was headed. “This isn’t only about the apocalyptic potential of what could happen; I’m also thinking a lot about the future of human flourishing,” Overman says. He fears that misaligned AI could overstep its bounds — not necessarily maliciously — and inflict subtle, yet real, harms on people.
“To prevent that, we need to get this right,” Overman says. “We must set up the proper interactions and training and incentives for these AI agents and models. It’s critical to think about shaping all of that now, so that these tools help make life better for us, not worse.”
To think through the AI architecture that could help to ensure this balance, Overman enlisted his technical toolkit, calling on game theory, reinforcement learning, and causal inference. He partnered up with Mohsen Bayati , his advisor and a professor of operations, information, and technology at Stanford GSB.
Overman and Bayati have coauthored two recent papers focused on a single question: As AI systems become capable of acting on their own, how do humans stay meaningfully in control?
The papers approach the question from two directions. One demonstrates how to design a relationship between a person and an AI agent from scratch, so the agent’s drive toward independence never comes at the human’s expense. The other paper offers a more practical framework that can be used with an AI system a user didn’t build and doesn’t fully trust.
“There’s a theoretical component at the core of this that is very important,” Bayati says. “We’re developing frameworks to think about two classes of problems: collaboratively building a safe AI, and controlling a potentially unsafe but very powerful one.”
In the first paper, titled “The Oversight Game,” the researchers modeled the relationship between a person and an AI agent as a game without a winner or a loser. “It’s more like cooperative party games, where you’re trying to accomplish something together,” Overman says. “It’s not like the AI is trying to beat the human, or vice versa.” Self-driving cars are a version of this idea: Each car is trying to reach its own destination, but all the vehicles on the road share an overriding interest in avoiding collisions.
In the game Overman and Bayati designed, an AI agent can choose between two moves at any moment: Act autonomously, or defer to a human who can override its action or shut the AI agent down. The human player simultaneously chooses whether to trust the AI’s judgment or step in. Intervention is possible only when the AI defers and the human chooses to oversee. The overarching goal: Train the AI agent to recognize for itself the right moments to pause and ask the human for guidance.
The researchers tested this setup in an environment called Lavaland, a grid the AI agent must cross to reach a goal. However, the AI agent was never trained to recognize the patches of lava placed in its path. Left on its own, the AI agent would choose a route straight through the lava, triggering a steep penalty.
Over repeated attempts, the AI agent and the human learn, independently and without being told to coordinate, when to defer and when to intervene. The AI agent learns to ask for guidance as it nears lava; the human learns to step in and pick a safe (if not always optimal) way around it. (Deferring and overseeing incur a small cost, discouraging either player from overrelying on the other.) When there’s no danger, the agent defaults to acting freely and the person defaults to trusting the AI.
It’s significant that intervention comes with a cost, Bayati explains. He points to a similar situation in a healthcare setting. If an AI assistant is meant to help a doctor save time by taking over rote tasks, it’s costly for it to ask for oversight too often, because it burns the doctor’s time. But it’s also costly to act autonomously too often, because that can lead to mistakes.
The shared goal for both AI and human players in the Oversight Game is to settle into an equilibrium of trusting and double-checking at the right moments. That shared stake led to a useful finding: When they play the game by the rules, the two players settle into a stable pattern where the AI’s push toward greater autonomy can never come at the human’s expense.
The payoff wasn’t just theoretical. In the researchers’ tests with frontier models for agentic coding, the AI learned on its own to ask for a check-in before taking risky actions, while the human learned when it was safe to step back.
“That’s what we set up this game to allow: the human and AI both learning when those critical points are,” Overman says. “If the AI has to ask for oversight at every moment, it might as well not be autonomous. But we do want to make sure that, at very critical points, it knows to ask.”
Overman and Bayati’s second paper takes on a messier yet more common situation: overseeing an AI that you didn’t design. “Imagine there’s a company that wants to use an open-source AI model that was built by an entity they don’t fully trust,” Bayati says. “But they want to use it, because the AI is very powerful; its performance is off the charts.”
If you can’t retrain the model, and you can’t redesign the role it’s playing, what should you do? The researchers’ response: You watch it.
Their method, which they call calibrated collective oversight, pools several weaker, imperfect human or AI overseers — none strong enough to trust alone — and uses them to collectively rein in the stronger AI system.
Critically, the researchers can mathematically guarantee the rate of bad outcomes will stay below whatever threshold a user sets.
“Using our framework, you can set a guarantee: There will be at most 5%, or 1%, unsafe decisions,” Bayati says. “You just tune the safety component.”
That guarantee matters, because most prior approaches to AI oversight have offered only empirical evidence that they seemed to work, not a formal promise they’d keep working as AI models grew more capable.
Bayati and Overman tested their approach on a modified software-engineering benchmark, where weaker overseers successfully held an adversarially misaligned coding agent in check. They also tested their framework on a game called MACHIAVELLI, which was built to test how AI systems trade off completing a task against acting ethically. In both cases, the rate of violations tracked closely with the safety target the researchers dialed in.
Of the two papers, Bayati sees the second as closer to real-world, practical use today. “It’s such a simple algorithm,” Bayati says. All it requires is a set of weaker systems you already trust — whether they’re human or smaller AI models.
The framework described in the Oversight Game, by contrast, needs practice before it’s ready for deployment. The AI and the human have to play the “game” repeatedly in the relevant setting (whether a doctor’s office, a classroom, or somewhere else) to find their footing and smooth out the mistakes that will happen along the way.
Building those testing environments well, Bayati notes, is a significant undertaking. Still, he and Overman see their frameworks as tools others can build on as researchers and AI developers work to solve the problem of AI alignment.
Mohsen Bayati teaches Business Intelligence from Big Data and AI and other courses.
Copy Link
Copied to clipboard
Share this
https://stanford.io/46S1m7M
Sign up for more insights and ideas.
For media inquiries, visit the Newsroom .
Researchers Build a Virtual World to Run Experiments Over and Over
AI Could Make These Common Jobs More Productive Without Sacrificing Quality
Designing AI That Keeps Human Decision-Makers in Mind
Subscribe to Stanford Business Email
Contact Media Relations
More
Close
655 Knight Way
Stanford, CA 94305
USA
Footer contact links
Companies, Organizations & Recruiters

## Original Extract

New research teaches independent AI agents to ask for help — and builds oversight of models we don’t fully trust.

A Blueprint for Keeping Humans in Control of AI | Stanford Graduate School of Business
Skip to main content
Menu
Enter the terms you wish to search for.
Insights by Stanford Business
Topics
Accounting
A Blueprint for Keeping Humans in Control of AI
New research teaches independent AI agents to ask for help — and builds oversight of models we don’t fully trust.
Stanford GSB researchers are building frameworks that keep AI under human control even as they become more autonomous.
In a cooperative game, an AI agent learns when to check in with a person and when to act alone.
Another framework sets up effective oversight of an untrusted model by pools of human or AI supervisors.
William Overman began his PhD program at Stanford Graduate School of Business in an auspicious moment: just two months before ChatGPT publicly launched in November 2022, exploding the widely held understandings of what machines are capable of.
Even as Overman began enlisting AI for his research, he grew wary of where the technology was headed. “This isn’t only about the apocalyptic potential of what could happen; I’m also thinking a lot about the future of human flourishing,” Overman says. He fears that misaligned AI could overstep its bounds — not necessarily maliciously — and inflict subtle, yet real, harms on people.
“To prevent that, we need to get this right,” Overman says. “We must set up the proper interactions and training and incentives for these AI agents and models. It’s critical to think about shaping all of that now, so that these tools help make life better for us, not worse.”
To think through the AI architecture that could help to ensure this balance, Overman enlisted his technical toolkit, calling on game theory, reinforcement learning, and causal inference. He partnered up with Mohsen Bayati , his advisor and a professor of operations, information, and technology at Stanford GSB.
Overman and Bayati have coauthored two recent papers focused on a single question: As AI systems become capable of acting on their own, how do humans stay meaningfully in control?
The papers approach the question from two directions. One demonstrates how to design a relationship between a person and an AI agent from scratch, so the agent’s drive toward independence never comes at the human’s expense. The other paper offers a more practical framework that can be used with an AI system a user didn’t build and doesn’t fully trust.
“There’s a theoretical component at the core of this that is very important,” Bayati says. “We’re developing frameworks to think about two classes of problems: collaboratively building a safe AI, and controlling a potentially unsafe but very powerful one.”
In the first paper, titled “The Oversight Game,” the researchers modeled the relationship between a person and an AI agent as a game without a winner or a loser. “It’s more like cooperative party games, where you’re trying to accomplish something together,” Overman says. “It’s not like the AI is trying to beat the human, or vice versa.” Self-driving cars are a version of this idea: Each car is trying to reach its own destination, but all the vehicles on the road share an overriding interest in avoiding collisions.
In the game Overman and Bayati designed, an AI agent can choose between two moves at any moment: Act autonomously, or defer to a human who can override its action or shut the AI agent down. The human player simultaneously chooses whether to trust the AI’s judgment or step in. Intervention is possible only when the AI defers and the human chooses to oversee. The overarching goal: Train the AI agent to recognize for itself the right moments to pause and ask the human for guidance.
The researchers tested this setup in an environment called Lavaland, a grid the AI agent must cross to reach a goal. However, the AI agent was never trained to recognize the patches of lava placed in its path. Left on its own, the AI agent would choose a route straight through the lava, triggering a steep penalty.
Over repeated attempts, the AI agent and the human learn, independently and without being told to coordinate, when to defer and when to intervene. The AI agent learns to ask for guidance as it nears lava; the human learns to step in and pick a safe (if not always optimal) way around it. (Deferring and overseeing incur a small cost, discouraging either player from overrelying on the other.) When there’s no danger, the agent defaults to acting freely and the person defaults to trusting the AI.
It’s significant that intervention comes with a cost, Bayati explains. He points to a similar situation in a healthcare setting. If an AI assistant is meant to help a doctor save time by taking over rote tasks, it’s costly for it to ask for oversight too often, because it burns the doctor’s time. But it’s also costly to act autonomously too often, because that can lead to mistakes.
The shared goal for both AI and human players in the Oversight Game is to settle into an equilibrium of trusting and double-checking at the right moments. That shared stake led to a useful finding: When they play the game by the rules, the two players settle into a stable pattern where the AI’s push toward greater autonomy can never come at the human’s expense.
The payoff wasn’t just theoretical. In the researchers’ tests with frontier models for agentic coding, the AI learned on its own to ask for a check-in before taking risky actions, while the human learned when it was safe to step back.
“That’s what we set up this game to allow: the human and AI both learning when those critical points are,” Overman says. “If the AI has to ask for oversight at every moment, it might as well not be autonomous. But we do want to make sure that, at very critical points, it knows to ask.”
Overman and Bayati’s second paper takes on a messier yet more common situation: overseeing an AI that you didn’t design. “Imagine there’s a company that wants to use an open-source AI model that was built by an entity they don’t fully trust,” Bayati says. “But they want to use it, because the AI is very powerful; its performance is off the charts.”
If you can’t retrain the model, and you can’t redesign the role it’s playing, what should you do? The researchers’ response: You watch it.
Their method, which they call calibrated collective oversight, pools several weaker, imperfect human or AI overseers — none strong enough to trust alone — and uses them to collectively rein in the stronger AI system.
Critically, the researchers can mathematically guarantee the rate of bad outcomes will stay below whatever threshold a user sets.
“Using our framework, you can set a guarantee: There will be at most 5%, or 1%, unsafe decisions,” Bayati says. “You just tune the safety component.”
That guarantee matters, because most prior approaches to AI oversight have offered only empirical evidence that they seemed to work, not a formal promise they’d keep working as AI models grew more capable.
Bayati and Overman tested their approach on a modified software-engineering benchmark, where weaker overseers successfully held an adversarially misaligned coding agent in check. They also tested their framework on a game called MACHIAVELLI, which was built to test how AI systems trade off completing a task against acting ethically. In both cases, the rate of violations tracked closely with the safety target the researchers dialed in.
Of the two papers, Bayati sees the second as closer to real-world, practical use today. “It’s such a simple algorithm,” Bayati says. All it requires is a set of weaker systems you already trust — whether they’re human or smaller AI models.
The framework described in the Oversight Game, by contrast, needs practice before it’s ready for deployment. The AI and the human have to play the “game” repeatedly in the relevant setting (whether a doctor’s office, a classroom, or somewhere else) to find their footing and smooth out the mistakes that will happen along the way.
Building those testing environments well, Bayati notes, is a significant undertaking. Still, he and Overman see their frameworks as tools others can build on as researchers and AI developers work to solve the problem of AI alignment.
Mohsen Bayati teaches Business Intelligence from Big Data and AI and other courses.
Copy Link
Copied to clipboard
Share this
https://stanford.io/46S1m7M
Sign up for more insights and ideas.
For media inquiries, visit the Newsroom .
Researchers Build a Virtual World to Run Experiments Over and Over
AI Could Make These Common Jobs More Productive Without Sacrificing Quality
Designing AI That Keeps Human Decision-Makers in Mind
Subscribe to Stanford Business Email
Contact Media Relations
More
Close
655 Knight Way
Stanford, CA 94305
USA
Footer contact links
Companies, Organizations & Recruiters
