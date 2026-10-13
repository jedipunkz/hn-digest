---
source: "https://www.quantamagazine.org/as-ai-closed-in-on-unique-games-proof-researchers-raced-to-beat-the-machines-20261007/"
hn_url: "https://news.ycombinator.com/item?id=50074999"
title: "As AI Closed in on 'Unique Games' Proof, Researchers Raced to Beat the Machines"
article_title: "As AI Closed In on ‘Unique Games’ Proof, Researchers Raced to Beat the Machines | Quanta Magazine"
image: "https://www.quantamagazine.org/wp-content/uploads/2026/10/Unique-Games-cr.James-OBrien-Social.jpg"
author: "ibobev"
captured_at: "2026-10-13T17:58:31Z"
capture_tool: "hn-digest"
hn_id: 50074999
score: 1
comments: 0
posted_at: "2026-10-13T17:02:45Z"
tags:
  - hacker-news
---

# As AI Closed in on 'Unique Games' Proof, Researchers Raced to Beat the Machines

- HN: [50074999](https://news.ycombinator.com/item?id=50074999)
- Source: [www.quantamagazine.org](https://www.quantamagazine.org/as-ai-closed-in-on-unique-games-proof-researchers-raced-to-beat-the-machines-20261007/)
- Score: 1
- Comments: 0
- Posted: 2026-10-13T17:02:45Z

## Translation

Title: As AI Closed in on 'Unique Games' Proof, Researchers Raced to Beat the Machines
Article title: As AI Closed In on ‘Unique Games’ Proof, Researchers Raced to Beat the Machines | Quanta Magazine
Description: In the shadow of a rumored AI proof of one of the biggest problems in their field, three computer scientists rushed to publish their own milestone result.

Article text:
As AI Closed In on ‘Unique Games’ Proof, Researchers Raced to Beat the Machines | Quanta Magazine
Quanta Homepage
Physics
Newsletter
Get the latest news delivered to your inbox.
Youtube
Instagram
RSS
Newsletter
An editorially independent publication supported by the Simons Foundation.
Create a reading list by clicking the Read Later icon next to the articles you wish to save.
Type search term(s) and press enter
Home
As AI Closed In on ‘Unique Games’ Proof, Researchers Raced to Beat the Machines
Comment
Save Article
Read Later
Share
Facebook
Copied!
Copy link
Email
Pocket
Reddit
Ycombinator
Save Article
Read Later
Read Later
As AI Closed In on ‘Unique Games’ Proof, Researchers Raced to Beat the Machines
Save Article
Read Later
James O’Brien for Quanta Magazine
artificial intelligence
computational complexity
computer science
graph theory
unique games conjecture
All topics
On the morning of September 11, 2026, Dor Minzer received a text message from a friend asking whether he was close to settling one of the most famous open questions in theoretical computer science. Minzer, a professor at the Massachusetts Institute of Technology and a leading expert on the “unique games” conjecture, initially thought it was a joke. Then more messages started coming in, and the rumors began to cohere. The artificial intelligence company OpenAI, fresh from announcing a bombshell proof about the behavior of fluids that sent ripples through the math world , had allegedly discovered a proof of the unique games conjecture. They might make the result public any day. If Minzer had any results of his own to share, the messages suggested, now would be the time to do so.
The unique games conjecture is an iconic question in computational complexity theory, the study of the inherent difficulty of mathematical problems. Roughly speaking, it states that a problem about satisfying multiple constraints at the same time can be extremely hard, even if you’re willing to settle for a poor approximation of the best possible solution. A proof of the conjecture would also automatically imply that current methods for solving many seemingly unrelated problems can’t be improved. Researchers would be a big step closer to a unified theory of computational difficulty.
Minzer had not, in fact, proved the unique games conjecture. But he and his graduate students Yumou Fei and Shuo Wang had recently proved a milestone result about a closely related question. They were in the midst of writing up their findings — a process that would typically take months or longer — when the rumors began circulating. The trio rushed to finish their draft to avoid being overshadowed by an OpenAI press release. Three days later, prioritizing completeness over clarity, they posted a 95-page paper online with a disclaimer on the first page: “The current version of the manuscript is complete mathematically, but it is not in the shape we wished to share in.”
Copied!
Copy link
Email
Pocket
Reddit
Ycombinator
Newsletter
Get Quanta Magazine delivered to your inbox
Dor Minzer and his students recently rushed to post a milestone result in theoretical computer science, hoping to avoid being upstaged by a rumored AI proof.
Their decision to publish quickly was soon vindicated. On October 6, OpenAI announced a proof of the unique games conjecture, along with 376 other results across many different fields of math. Researchers greeted the deluge of AI proofs with the mixture of alarm and curiosity that has become familiar in the wake of recent advances. Their reaction to Minzer’s result was less ambivalent.
“He’s put out many great papers on this topic,” said Ryan O’Donnell , a theoretical computer scientist at Carnegie Mellon University. “This is another truly great one.”
The problems at the heart of both the unique games conjecture and the team’s new result are about simultaneously satisfying many constraints that may be in tension with each other. Similar problems arise in everyday life: filling in a sudoku puzzle, for instance, or planning seating arrangements for a wedding. (You’ll want to seat friends together, while making sure that guests who don’t get along are assigned to different tables — and you’ll have to fill all the seats at each table.)
Solving constraint satisfaction problems exactly can be hard or even impossible. Sometimes, there is a solution that satisfies every constraint, but all known methods for finding that solution are painfully slow. Or the constraints might conflict with each other, and a perfect solution simply doesn’t exist.
In these cases, researchers are often willing to settle for approximate solutions that satisfy some but not all of the constraints. Complexity theorists want to understand why it’s sometimes easy and sometimes hard to find these approximate solutions. “A lot of work has been devoted to trying to figure out where is the boundary,” O’Donnell said.
The unique games conjecture, first articulated in a 2002 paper by the complexity theorist Subhash Khot when he was still a graduate student, is intimately connected to that quest. It concerns a specific constraint satisfaction problem about mathematical objects called graphs, which are networks of nodes connected by lines called edges. You’re given a graph and asked to color its nodes using a fixed color palette. Each edge comes with a rule that determines what color the nodes at either end should be. For instance, one particular edge might require that the two nodes have the same color. Another edge might demand that if one node is red, the other must be blue, and that if one node is green, the other must be yellow.
Subhash Khot posed the unique games conjecture back when he was still in graduate school.
Béatrice de Géa for Quanta Magazine
The difficulty of this coloring problem will depend on the layout of the graph, the number of allowed colors, and the specific set of constraints. Khot was interested in cases where the best possible coloring satisfies most but not all of the constraints — say, 99% of them. Finding this near-perfect coloring is a difficult task. Intuitively it would seem easier to find a coloring that satisfies a smaller fraction of the constraints — 1% or even 0.001% of them.
The unique games conjecture says that this intuition is wrong. No matter how much you’re willing to relax your standards, there will always be cases where even these subpar colorings are hard to find.
The unique games conjecture makes a striking claim, but its real value lies in its connections to other problems. Researchers have used the conjecture to attack apparently unrelated questions about the geometry of foams and the properties of voting systems . And in 2008, the computer scientist Prasad Raghavendra proved that if the conjecture is true , one particular classic algorithm is the best strategy for every constraint satisfaction problem where no perfect solution exists. You can’t do better by exploiting the quirks of specific problems.
They solved the problem in the old-fashioned way, with their minds, and wrote it with their own fingers.
Ryan O’Donnell, Carnegie Mellon University
Yet the conjecture has a blind spot. It only applies to cases where the best possible solution satisfies most constraints, such as the 99% case above. It says nothing about cases where there’s a perfect solution — where 100% of the constraints can be satisfied.
Khot recognized this limitation from the beginning. To address those cases, he defined a variant of the unique games problem called the 2-to-1 games problem, in which the constraints are looser. In the original problem, coloring a node on one side of an edge must leave only one option for the color on the other side. In the new version, he allowed for two possible colors.
For this problem, he conjectured, there are cases where the best possible solution satisfies all constraints, but it’s still hard to find a solution that satisfies any tiny fraction. A proof of this lesser-known conjecture would also imply results about a host of other problems that don’t fall under the umbrella of the unique games conjecture.
In their paper from September, Minzer and his colleagues proved a slight variant of Khot’s second conjecture that has nearly the same consequences. For Minzer, the work was seven years in the making.
In 2018, while still a graduate student, Minzer played a leading role in a celebrated result marking the first significant advance toward proving Khot’s two conjectures. But his progress eventually stalled. Then, in 2025, he discussed the 2-to-1 problem with Fei and Wang, who had just finished their first year of graduate school and wrapped up a fruitful first collaboration.
“I gave them two options of which problem to choose,” Minzer said. “They were brave enough to choose this one.”
First Big Steps Toward Proving the Unique Games Conjecture
Save Article
Read Later
To settle the conjecture, the trio would need to build a mathematical bridge from the 2-to-1 problem to another problem whose difficulty was already well understood. In Minzer’s previous work, the first step in building that bridge involved transforming the graph-coloring problem into a problem about “error-correcting codes,” a method for encoding messages so that a recipient can detect and fix transmission errors.
After a few months, Fei and Wang identified a promising new approach to building error-correcting codes. But they couldn’t see how to combine those new codes with all the other pieces that needed to go into the bridge.
“It looked very remote,” Minzer said. “You needed miracles to happen, all of the stars to align.”
For months after that, the stars did not align. The trio repeatedly tried and failed to make the proof work, but assembling the pieces felt “like trying to fit a circle and a triangle together in a jigsaw puzzle,” Fei said. Yet they learned something from each thwarted attempt. In April 2026, they finally succeeded.
“We took five failures and managed to staple this [new code] on top of them,” Minzer said. “And then it worked.”
Technically, Minzer, Fei, and Wang proved a slightly weaker version of Khot’s 2-to-1 games conjecture, in which there are four options for each constraint rather than two. But many of the implications of Khot’s original conjecture already followed from this 4-to-1 version — most notably, a new result about the difficulty of a classic graph-coloring problem that predates Khot’s work by decades. In this famous problem, the only constraint is that adjacent nodes can’t have the same color. Complexity theorists have long known that in cases where it’s possible to color a given graph with just three colors, such a solution can be prohibitively difficult to find. Minzer, Fei, and Wang’s new result now also implies that in these cases, no matter how many extra colors you’re allowed to use, there will always be perverse graphs where it’s still hard to find a solution.
Shuo Wang (top) and Yumou Fei used an exotic error-correcting code to make their new proof work where others had failed.
“You cannot do it even with the entire Crayola box,” said Mark Braverman of Princeton University. Computer scientists had sought to prove this result for decades — in fact, it was one of the problems that originally motivated Khot to propose his graph-coloring conjectures.
OpenAI’s October release included a Lean-verified proof of the 2-to-1 conjecture, but in an AI-generated manuscript that didn’t undergo any human editing or review by independent experts. Not that the September 4-to-1 result, which was rushed to predate the OpenAI dump, is in perfect shape.
“From section 6 onward, there are literally no connecting words,” Minzer said — just an unbroken sequence of definitions and proofs of intermediate results. “We felt the need to apologize for that.”
The team plans to post a new version of their paper when they’ve finished fleshing it out. But in an age where AI-assisted writing is becoming increasingly ubiquitous, the distinctly unpolishe

[truncated]

## Original Extract

In the shadow of a rumored AI proof of one of the biggest problems in their field, three computer scientists rushed to publish their own milestone result.

As AI Closed In on ‘Unique Games’ Proof, Researchers Raced to Beat the Machines | Quanta Magazine
Quanta Homepage
Physics
Newsletter
Get the latest news delivered to your inbox.
Youtube
Instagram
RSS
Newsletter
An editorially independent publication supported by the Simons Foundation.
Create a reading list by clicking the Read Later icon next to the articles you wish to save.
Type search term(s) and press enter
Home
As AI Closed In on ‘Unique Games’ Proof, Researchers Raced to Beat the Machines
Comment
Save Article
Read Later
Share
Facebook
Copied!
Copy link
Email
Pocket
Reddit
Ycombinator
Save Article
Read Later
Read Later
As AI Closed In on ‘Unique Games’ Proof, Researchers Raced to Beat the Machines
Save Article
Read Later
James O’Brien for Quanta Magazine
artificial intelligence
computational complexity
computer science
graph theory
unique games conjecture
All topics
On the morning of September 11, 2026, Dor Minzer received a text message from a friend asking whether he was close to settling one of the most famous open questions in theoretical computer science. Minzer, a professor at the Massachusetts Institute of Technology and a leading expert on the “unique games” conjecture, initially thought it was a joke. Then more messages started coming in, and the rumors began to cohere. The artificial intelligence company OpenAI, fresh from announcing a bombshell proof about the behavior of fluids that sent ripples through the math world , had allegedly discovered a proof of the unique games conjecture. They might make the result public any day. If Minzer had any results of his own to share, the messages suggested, now would be the time to do so.
The unique games conjecture is an iconic question in computational complexity theory, the study of the inherent difficulty of mathematical problems. Roughly speaking, it states that a problem about satisfying multiple constraints at the same time can be extremely hard, even if you’re willing to settle for a poor approximation of the best possible solution. A proof of the conjecture would also automatically imply that current methods for solving many seemingly unrelated problems can’t be improved. Researchers would be a big step closer to a unified theory of computational difficulty.
Minzer had not, in fact, proved the unique games conjecture. But he and his graduate students Yumou Fei and Shuo Wang had recently proved a milestone result about a closely related question. They were in the midst of writing up their findings — a process that would typically take months or longer — when the rumors began circulating. The trio rushed to finish their draft to avoid being overshadowed by an OpenAI press release. Three days later, prioritizing completeness over clarity, they posted a 95-page paper online with a disclaimer on the first page: “The current version of the manuscript is complete mathematically, but it is not in the shape we wished to share in.”
Copied!
Copy link
Email
Pocket
Reddit
Ycombinator
Newsletter
Get Quanta Magazine delivered to your inbox
Dor Minzer and his students recently rushed to post a milestone result in theoretical computer science, hoping to avoid being upstaged by a rumored AI proof.
Their decision to publish quickly was soon vindicated. On October 6, OpenAI announced a proof of the unique games conjecture, along with 376 other results across many different fields of math. Researchers greeted the deluge of AI proofs with the mixture of alarm and curiosity that has become familiar in the wake of recent advances. Their reaction to Minzer’s result was less ambivalent.
“He’s put out many great papers on this topic,” said Ryan O’Donnell , a theoretical computer scientist at Carnegie Mellon University. “This is another truly great one.”
The problems at the heart of both the unique games conjecture and the team’s new result are about simultaneously satisfying many constraints that may be in tension with each other. Similar problems arise in everyday life: filling in a sudoku puzzle, for instance, or planning seating arrangements for a wedding. (You’ll want to seat friends together, while making sure that guests who don’t get along are assigned to different tables — and you’ll have to fill all the seats at each table.)
Solving constraint satisfaction problems exactly can be hard or even impossible. Sometimes, there is a solution that satisfies every constraint, but all known methods for finding that solution are painfully slow. Or the constraints might conflict with each other, and a perfect solution simply doesn’t exist.
In these cases, researchers are often willing to settle for approximate solutions that satisfy some but not all of the constraints. Complexity theorists want to understand why it’s sometimes easy and sometimes hard to find these approximate solutions. “A lot of work has been devoted to trying to figure out where is the boundary,” O’Donnell said.
The unique games conjecture, first articulated in a 2002 paper by the complexity theorist Subhash Khot when he was still a graduate student, is intimately connected to that quest. It concerns a specific constraint satisfaction problem about mathematical objects called graphs, which are networks of nodes connected by lines called edges. You’re given a graph and asked to color its nodes using a fixed color palette. Each edge comes with a rule that determines what color the nodes at either end should be. For instance, one particular edge might require that the two nodes have the same color. Another edge might demand that if one node is red, the other must be blue, and that if one node is green, the other must be yellow.
Subhash Khot posed the unique games conjecture back when he was still in graduate school.
Béatrice de Géa for Quanta Magazine
The difficulty of this coloring problem will depend on the layout of the graph, the number of allowed colors, and the specific set of constraints. Khot was interested in cases where the best possible coloring satisfies most but not all of the constraints — say, 99% of them. Finding this near-perfect coloring is a difficult task. Intuitively it would seem easier to find a coloring that satisfies a smaller fraction of the constraints — 1% or even 0.001% of them.
The unique games conjecture says that this intuition is wrong. No matter how much you’re willing to relax your standards, there will always be cases where even these subpar colorings are hard to find.
The unique games conjecture makes a striking claim, but its real value lies in its connections to other problems. Researchers have used the conjecture to attack apparently unrelated questions about the geometry of foams and the properties of voting systems . And in 2008, the computer scientist Prasad Raghavendra proved that if the conjecture is true , one particular classic algorithm is the best strategy for every constraint satisfaction problem where no perfect solution exists. You can’t do better by exploiting the quirks of specific problems.
They solved the problem in the old-fashioned way, with their minds, and wrote it with their own fingers.
Ryan O’Donnell, Carnegie Mellon University
Yet the conjecture has a blind spot. It only applies to cases where the best possible solution satisfies most constraints, such as the 99% case above. It says nothing about cases where there’s a perfect solution — where 100% of the constraints can be satisfied.
Khot recognized this limitation from the beginning. To address those cases, he defined a variant of the unique games problem called the 2-to-1 games problem, in which the constraints are looser. In the original problem, coloring a node on one side of an edge must leave only one option for the color on the other side. In the new version, he allowed for two possible colors.
For this problem, he conjectured, there are cases where the best possible solution satisfies all constraints, but it’s still hard to find a solution that satisfies any tiny fraction. A proof of this lesser-known conjecture would also imply results about a host of other problems that don’t fall under the umbrella of the unique games conjecture.
In their paper from September, Minzer and his colleagues proved a slight variant of Khot’s second conjecture that has nearly the same consequences. For Minzer, the work was seven years in the making.
In 2018, while still a graduate student, Minzer played a leading role in a celebrated result marking the first significant advance toward proving Khot’s two conjectures. But his progress eventually stalled. Then, in 2025, he discussed the 2-to-1 problem with Fei and Wang, who had just finished their first year of graduate school and wrapped up a fruitful first collaboration.
“I gave them two options of which problem to choose,” Minzer said. “They were brave enough to choose this one.”
First Big Steps Toward Proving the Unique Games Conjecture
Save Article
Read Later
To settle the conjecture, the trio would need to build a mathematical bridge from the 2-to-1 problem to another problem whose difficulty was already well understood. In Minzer’s previous work, the first step in building that bridge involved transforming the graph-coloring problem into a problem about “error-correcting codes,” a method for encoding messages so that a recipient can detect and fix transmission errors.
After a few months, Fei and Wang identified a promising new approach to building error-correcting codes. But they couldn’t see how to combine those new codes with all the other pieces that needed to go into the bridge.
“It looked very remote,” Minzer said. “You needed miracles to happen, all of the stars to align.”
For months after that, the stars did not align. The trio repeatedly tried and failed to make the proof work, but assembling the pieces felt “like trying to fit a circle and a triangle together in a jigsaw puzzle,” Fei said. Yet they learned something from each thwarted attempt. In April 2026, they finally succeeded.
“We took five failures and managed to staple this [new code] on top of them,” Minzer said. “And then it worked.”
Technically, Minzer, Fei, and Wang proved a slightly weaker version of Khot’s 2-to-1 games conjecture, in which there are four options for each constraint rather than two. But many of the implications of Khot’s original conjecture already followed from this 4-to-1 version — most notably, a new result about the difficulty of a classic graph-coloring problem that predates Khot’s work by decades. In this famous problem, the only constraint is that adjacent nodes can’t have the same color. Complexity theorists have long known that in cases where it’s possible to color a given graph with just three colors, such a solution can be prohibitively difficult to find. Minzer, Fei, and Wang’s new result now also implies that in these cases, no matter how many extra colors you’re allowed to use, there will always be perverse graphs where it’s still hard to find a solution.
Shuo Wang (top) and Yumou Fei used an exotic error-correcting code to make their new proof work where others had failed.
“You cannot do it even with the entire Crayola box,” said Mark Braverman of Princeton University. Computer scientists had sought to prove this result for decades — in fact, it was one of the problems that originally motivated Khot to propose his graph-coloring conjectures.
OpenAI’s October release included a Lean-verified proof of the 2-to-1 conjecture, but in an AI-generated manuscript that didn’t undergo any human editing or review by independent experts. Not that the September 4-to-1 result, which was rushed to predate the OpenAI dump, is in perfect shape.
“From section 6 onward, there are literally no connecting words,” Minzer said — just an unbroken sequence of definitions and proofs of intermediate results. “We felt the need to apologize for that.”
The team plans to post a new version of their paper when they’ve finished fleshing it out. But in an age where AI-assisted writing is becoming increasingly ubiquitous, the distinctly unpolishe

[truncated]
