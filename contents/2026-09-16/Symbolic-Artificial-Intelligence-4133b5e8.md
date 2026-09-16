---
source: "https://en.wikipedia.org/wiki/Symbolic_artificial_intelligence"
hn_url: "https://news.ycombinator.com/item?id=49721959"
title: "Symbolic Artificial Intelligence"
article_title: "Symbolic artificial intelligence - Wikipedia"
image: ""
author: "ijidak"
captured_at: "2026-09-16T04:39:05Z"
capture_tool: "hn-digest"
hn_id: 49721959
score: 1
comments: 0
posted_at: "2026-09-16T04:03:31Z"
tags:
  - hacker-news
---

# Symbolic Artificial Intelligence

- HN: [49721959](https://news.ycombinator.com/item?id=49721959)
- Source: [en.wikipedia.org](https://en.wikipedia.org/wiki/Symbolic_artificial_intelligence)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T04:03:31Z

## Translation

Title: Symbolic Artificial Intelligence
Article title: Symbolic artificial intelligence - Wikipedia

Article text:
Jump to content
Main menu
Main menu
move to sidebar
hide
Navigation
Main page
1
History
Toggle History subsection
1.1
The first AI summer: irrational exuberance, 1948–1966
1.1.1
Approaches inspired by human or animal cognition or behavior
1.1.3
Early work on knowledge representation and reasoning
1.1.3.1
Modeling formal reasoning with logic: the "neats"
1.1.3.2
Modeling implicit common-sense knowledge with frames and scripts: the "scruffies"
1.2
The first AI winter: crushed dreams, 1967–1977
1.3
The second AI summer: knowledge is power, 1978–1987
1.3.1
Knowledge-based systems
1.3.2
Success with expert systems
1.3.2.1
Architecture of knowledge-based and expert systems
1.4
The second AI winter, 1988–1993
1.5
Adding in more rigorous foundations, 1993–2011
1.5.1
Uncertain reasoning
1.6
Deep learning and neuro-symbolic AI 2011–now
1.6.1
Neuro-symbolic AI: integrating neural and symbolic approaches
2
Techniques and contributions
Toggle Techniques and contributions subsection
2.1
AI programming languages
2.3
Knowledge representation and reasoning
2.3.1
Knowledge representation
2.3.2
Automatic theorem proving
2.3.3
Reasoning in knowledge-based systems
2.3.5
Constraints and constraint-based reasoning
2.5
Natural language processing
2.6
Agents and multi-agent systems
3
Controversies
Toggle Controversies subsection
3.1
The Frame Problem: knowledge representation challenges for first-order logic
3.2
Connectionist AI: philosophical challenges and sociological conflicts
3.3
Situated robotics: the world as a model
Symbolic artificial intelligence
Part of a series on Artificial intelligence (AI)
Major goals
Artificial general intelligence
superintelligence
Deepfake pornography
Taylor Swift deepfake pornography controversy
Google Gemini image generation controversy
It's the Most Terrible Time of the Year
Navier–Stokes priority controversy
Removal of Sam Altman from OpenAI
Voiceverse NFT plagiarism scandal
In artificial intelligence (AI), symbolic artificial intelligence (also known as classical artificial intelligence or logic-based artificial intelligence ) [ 1 ] [ 2 ]
is a collection of methods based on high-level symbolic ( human-readable ) representations of problems, logic , and search . [ 1 ] Symbolic AI used tools such as logic programming , production rules , semantic nets and frames . It developed applications such as knowledge-based systems (in particular, expert systems ), symbolic mathematics , automated theorem provers , ontologies , the semantic web , and automated planning and scheduling systems. The symbolic AI paradigm led to important ideas in search , symbolic programming languages, agents , multi-agent systems , the semantic web, and the strengths and limitations of formal knowledge and reasoning systems .
Symbolic AI was the dominant paradigm of AI research from the mid-1950s until the mid-1990s. [ 3 ] Researchers in the 1960s and the 1970s were convinced that symbolic approaches would eventually succeed in creating a machine with artificial general intelligence and considered this the ultimate goal of their field. [ 4 ] An early boom, with early successes such as the Logic Theorist and Samuel 's Checkers Playing Program , led to unrealistic expectations and promises and was followed by the first AI Winter as funding dried up. [ 5 ] [ 6 ] A second boom (1969–1986) occurred with the rise of expert systems, their promise of capturing corporate expertise, and an enthusiastic corporate embrace. [ 7 ] [ 8 ] That boom, and some early successes, e.g., with XCON at DEC , was followed again by later disappointment. [ 8 ] Problems with difficulties in knowledge acquisition, maintaining large knowledge bases, and brittleness in handling out-of-domain problems arose. Another, second, AI Winter (1988–2011) followed. [ 9 ] Subsequently, AI researchers focused on addressing underlying problems in handling uncertainty and in knowledge acquisition. [ 10 ] Uncertainty was addressed with formal methods such as hidden Markov models , Bayesian reasoning , and statistical relational learning . [ 11 ] [ 12 ] Symbolic machine learning addressed the knowledge acquisition problem with contributions including Version Space , Valiant 's PAC learning , Quinlan 's ID3 decision-tree learning, case-based learning , and inductive logic programming to learn relations. [ 13 ]
Neural networks , a subsymbolic [ clarification needed ] approach, had been pursued from early days and reemerged strongly in 2012. Early examples are Rosenblatt 's perceptron learning work, the backpropagation work of Rumelhart, Hinton and Williams, [ 14 ] and work in convolutional neural networks by LeCun et al. in 1989. [ 15 ] However, neural networks were not viewed as successful until about 2012: "Until Big Data became commonplace, the general consensus in the Al community was that the so-called neural-network approach was hopeless. Systems just didn't work that well, compared to other methods. ... A revolution came in 2012, when a number of people, including a team of researchers working with Hinton, worked out a way to use the power of GPUs to enormously increase the power of neural networks." [ 16 ] Over the next several years, deep learning had spectacular success in handling vision, speech recognition , speech synthesis, image generation, and machine translation, though symbolic approaches continue to be useful in a few domains such as computer algebra systems and proof assistants .
A short history of symbolic AI to the present day follows below. Time periods and titles are drawn from Henry Kautz's 2020 AAAI Robert S. Engelmore Memorial Lecture [ 17 ] and the longer Wikipedia article on the History of AI , with dates and titles differing slightly for increased clarity.
The first AI summer: irrational exuberance, 1948–1966
Success at early attempts in AI occurred in three main areas: artificial neural networks, knowledge representation, and heuristic search, contributing to high expectations. This section summarizes Kautz's reprise of early AI history.
Cybernetic approaches attempted to replicate the feedback loops between animals and their environments. A robotic turtle, with sensors, motors for driving and steering, and seven vacuum tubes for control, based on a preprogrammed neural net, was built as early as 1948. This work can be seen as an early precursor to later work in neural networks, reinforcement learning, and situated robotics. [ 18 ]
An important early symbolic AI program was the Logic theorist , written by Allen Newell , Herbert Simon and Cliff Shaw in 1955–56, as it was able to prove 38 elementary theorems from Whitehead and Russell's Principia Mathematica . Newell, Simon, and Shaw later generalized this work to create a domain-independent problem solver, GPS (General Problem Solver). GPS solved problems represented with formal operators via state-space search using means-ends analysis . [ 19 ]
During the 1960s, symbolic approaches achieved great success at simulating intelligent behavior in structured environments such as game-playing, symbolic mathematics, and theorem-proving. AI research was concentrated in four institutions in the 1960s: Carnegie Mellon University , Stanford , MIT and (later) University of Edinburgh . Each one developed its own style of research. Earlier approaches based on cybernetics or artificial neural networks were abandoned or pushed into the background.
Herbert Simon and Allen Newell studied human problem-solving skills and attempted to formalize them, and their work laid the foundations of the field of artificial intelligence, as well as cognitive science , operations research and management science . Their research team used the results of psychological experiments to develop programs that simulated the techniques that people used to solve problems. [ 20 ] [ 21 ] This tradition, centered at Carnegie Mellon University would eventually culminate in the development of the Soar architecture in the middle 1980s. [ 22 ] [ 23 ]
In addition to the highly specialized domain-specific kinds of knowledge that we will see later used in expert systems, early symbolic AI researchers discovered another more general application of knowledge. These were called heuristics, rules of thumb that guide a search in promising directions: "How can non-enumerative search be practical when the underlying problem is exponentially hard? The approach advocated by Simon and Newell is to employ heuristics : fast algorithms that may fail on some inputs or output suboptimal solutions." [ 24 ] Another important advance was to find a way to apply these heuristics that guarantees a solution will be found, if there is one, not withstanding the occasional fallibility of heuristics: "The A* algorithm provided a general frame for complete and optimal heuristically guided search. A* is used as a subroutine within practically every AI algorithm today but is still no magic bullet; its guarantee of completeness is bought at the cost of worst-case exponential time. [ 24 ]
Early work covered both applications of formal reasoning emphasizing first-order logic , along with attempts to handle common-sense reasoning in a less formal manner.
Unlike Simon and Newell, John McCarthy felt that machines did not need to simulate the exact mechanisms of human thought, but could instead try to find the essence of abstract reasoning and problem-solving with logic, [ 25 ] regardless of whether people used the same algorithms. [ a ]
His laboratory at Stanford ( SAIL ) focused on using formal logic to solve a wide variety of problems, including knowledge representation , planning and learning . [ 29 ]
Logic was also the focus of the work at the University of Edinburgh and elsewhere in Europe which led to the development of the programming language Prolog and the science of logic programming. [ 30 ] [ 31 ]
Main article: neats vs. scruffies
Researchers at MIT (such as Marvin Minsky and Seymour Papert ) [ 32 ] [ 33 ] [ 6 ] found that solving difficult problems in vision and natural language processing required ad hoc solutions—they argued that no simple and general principle (like logic ) would capture all the aspects of intelligent behavior. Roger Schank described their "anti-logic" approaches as " scruffy " (as opposed to the " neat " paradigms at CMU and Stanford). [ 34 ] [ 35 ]
Commonsense knowledge bases (such as Doug Lenat 's Cyc ) are an example of "scruffy" AI, since they must be built by hand, one complicated concept at a time. [ 36 ] [ 37 ] [ 38 ]
The first AI winter: crushed dreams, 1967–1977
The first AI winter was a shock:
During the first AI summer, many people thought that machine intelligence could be achieved in just a few years. The Defense Advance Research Projects Agency (DARPA) launched programs to support AI research to use AI to solve problems of national security; in particular, to automate the translation of Russian to English for intelligence operations and to create autonomous tanks for the battlefield. Researchers had begun to realize that achieving AI was going to be much harder than was supposed a decade earlier, but a combination of hubris and disingenuousness led many university and think-tank researchers to accept funding with promises of deliverables that they should have known they could not fulfill. By the mid-1960s neither useful natural language translation systems nor autonomous tanks had been created, and a dramatic backlash set in. New DARPA leadership canceled existing AI funding programs.
Outside of the United States, the most fertile ground for AI research was the United Kingdom. The AI winter in the United Kingdom was spurred on not so much by disappointed military leaders as by rival academics who viewed AI researchers as charlatans and a drain on research funding. A professor of applied mathematics, Sir James Lighthill, was commissioned by Parliament to evaluate the state of AI research in the nation . The report stated that all of the problems being worked on in AI would be better handled by res

[truncated]

## Original Extract

Jump to content
Main menu
Main menu
move to sidebar
hide
Navigation
Main page
1
History
Toggle History subsection
1.1
The first AI summer: irrational exuberance, 1948–1966
1.1.1
Approaches inspired by human or animal cognition or behavior
1.1.3
Early work on knowledge representation and reasoning
1.1.3.1
Modeling formal reasoning with logic: the "neats"
1.1.3.2
Modeling implicit common-sense knowledge with frames and scripts: the "scruffies"
1.2
The first AI winter: crushed dreams, 1967–1977
1.3
The second AI summer: knowledge is power, 1978–1987
1.3.1
Knowledge-based systems
1.3.2
Success with expert systems
1.3.2.1
Architecture of knowledge-based and expert systems
1.4
The second AI winter, 1988–1993
1.5
Adding in more rigorous foundations, 1993–2011
1.5.1
Uncertain reasoning
1.6
Deep learning and neuro-symbolic AI 2011–now
1.6.1
Neuro-symbolic AI: integrating neural and symbolic approaches
2
Techniques and contributions
Toggle Techniques and contributions subsection
2.1
AI programming languages
2.3
Knowledge representation and reasoning
2.3.1
Knowledge representation
2.3.2
Automatic theorem proving
2.3.3
Reasoning in knowledge-based systems
2.3.5
Constraints and constraint-based reasoning
2.5
Natural language processing
2.6
Agents and multi-agent systems
3
Controversies
Toggle Controversies subsection
3.1
The Frame Problem: knowledge representation challenges for first-order logic
3.2
Connectionist AI: philosophical challenges and sociological conflicts
3.3
Situated robotics: the world as a model
Symbolic artificial intelligence
Part of a series on Artificial intelligence (AI)
Major goals
Artificial general intelligence
superintelligence
Deepfake pornography
Taylor Swift deepfake pornography controversy
Google Gemini image generation controversy
It's the Most Terrible Time of the Year
Navier–Stokes priority controversy
Removal of Sam Altman from OpenAI
Voiceverse NFT plagiarism scandal
In artificial intelligence (AI), symbolic artificial intelligence (also known as classical artificial intelligence or logic-based artificial intelligence ) [ 1 ] [ 2 ]
is a collection of methods based on high-level symbolic ( human-readable ) representations of problems, logic , and search . [ 1 ] Symbolic AI used tools such as logic programming , production rules , semantic nets and frames . It developed applications such as knowledge-based systems (in particular, expert systems ), symbolic mathematics , automated theorem provers , ontologies , the semantic web , and automated planning and scheduling systems. The symbolic AI paradigm led to important ideas in search , symbolic programming languages, agents , multi-agent systems , the semantic web, and the strengths and limitations of formal knowledge and reasoning systems .
Symbolic AI was the dominant paradigm of AI research from the mid-1950s until the mid-1990s. [ 3 ] Researchers in the 1960s and the 1970s were convinced that symbolic approaches would eventually succeed in creating a machine with artificial general intelligence and considered this the ultimate goal of their field. [ 4 ] An early boom, with early successes such as the Logic Theorist and Samuel 's Checkers Playing Program , led to unrealistic expectations and promises and was followed by the first AI Winter as funding dried up. [ 5 ] [ 6 ] A second boom (1969–1986) occurred with the rise of expert systems, their promise of capturing corporate expertise, and an enthusiastic corporate embrace. [ 7 ] [ 8 ] That boom, and some early successes, e.g., with XCON at DEC , was followed again by later disappointment. [ 8 ] Problems with difficulties in knowledge acquisition, maintaining large knowledge bases, and brittleness in handling out-of-domain problems arose. Another, second, AI Winter (1988–2011) followed. [ 9 ] Subsequently, AI researchers focused on addressing underlying problems in handling uncertainty and in knowledge acquisition. [ 10 ] Uncertainty was addressed with formal methods such as hidden Markov models , Bayesian reasoning , and statistical relational learning . [ 11 ] [ 12 ] Symbolic machine learning addressed the knowledge acquisition problem with contributions including Version Space , Valiant 's PAC learning , Quinlan 's ID3 decision-tree learning, case-based learning , and inductive logic programming to learn relations. [ 13 ]
Neural networks , a subsymbolic [ clarification needed ] approach, had been pursued from early days and reemerged strongly in 2012. Early examples are Rosenblatt 's perceptron learning work, the backpropagation work of Rumelhart, Hinton and Williams, [ 14 ] and work in convolutional neural networks by LeCun et al. in 1989. [ 15 ] However, neural networks were not viewed as successful until about 2012: "Until Big Data became commonplace, the general consensus in the Al community was that the so-called neural-network approach was hopeless. Systems just didn't work that well, compared to other methods. ... A revolution came in 2012, when a number of people, including a team of researchers working with Hinton, worked out a way to use the power of GPUs to enormously increase the power of neural networks." [ 16 ] Over the next several years, deep learning had spectacular success in handling vision, speech recognition , speech synthesis, image generation, and machine translation, though symbolic approaches continue to be useful in a few domains such as computer algebra systems and proof assistants .
A short history of symbolic AI to the present day follows below. Time periods and titles are drawn from Henry Kautz's 2020 AAAI Robert S. Engelmore Memorial Lecture [ 17 ] and the longer Wikipedia article on the History of AI , with dates and titles differing slightly for increased clarity.
The first AI summer: irrational exuberance, 1948–1966
Success at early attempts in AI occurred in three main areas: artificial neural networks, knowledge representation, and heuristic search, contributing to high expectations. This section summarizes Kautz's reprise of early AI history.
Cybernetic approaches attempted to replicate the feedback loops between animals and their environments. A robotic turtle, with sensors, motors for driving and steering, and seven vacuum tubes for control, based on a preprogrammed neural net, was built as early as 1948. This work can be seen as an early precursor to later work in neural networks, reinforcement learning, and situated robotics. [ 18 ]
An important early symbolic AI program was the Logic theorist , written by Allen Newell , Herbert Simon and Cliff Shaw in 1955–56, as it was able to prove 38 elementary theorems from Whitehead and Russell's Principia Mathematica . Newell, Simon, and Shaw later generalized this work to create a domain-independent problem solver, GPS (General Problem Solver). GPS solved problems represented with formal operators via state-space search using means-ends analysis . [ 19 ]
During the 1960s, symbolic approaches achieved great success at simulating intelligent behavior in structured environments such as game-playing, symbolic mathematics, and theorem-proving. AI research was concentrated in four institutions in the 1960s: Carnegie Mellon University , Stanford , MIT and (later) University of Edinburgh . Each one developed its own style of research. Earlier approaches based on cybernetics or artificial neural networks were abandoned or pushed into the background.
Herbert Simon and Allen Newell studied human problem-solving skills and attempted to formalize them, and their work laid the foundations of the field of artificial intelligence, as well as cognitive science , operations research and management science . Their research team used the results of psychological experiments to develop programs that simulated the techniques that people used to solve problems. [ 20 ] [ 21 ] This tradition, centered at Carnegie Mellon University would eventually culminate in the development of the Soar architecture in the middle 1980s. [ 22 ] [ 23 ]
In addition to the highly specialized domain-specific kinds of knowledge that we will see later used in expert systems, early symbolic AI researchers discovered another more general application of knowledge. These were called heuristics, rules of thumb that guide a search in promising directions: "How can non-enumerative search be practical when the underlying problem is exponentially hard? The approach advocated by Simon and Newell is to employ heuristics : fast algorithms that may fail on some inputs or output suboptimal solutions." [ 24 ] Another important advance was to find a way to apply these heuristics that guarantees a solution will be found, if there is one, not withstanding the occasional fallibility of heuristics: "The A* algorithm provided a general frame for complete and optimal heuristically guided search. A* is used as a subroutine within practically every AI algorithm today but is still no magic bullet; its guarantee of completeness is bought at the cost of worst-case exponential time. [ 24 ]
Early work covered both applications of formal reasoning emphasizing first-order logic , along with attempts to handle common-sense reasoning in a less formal manner.
Unlike Simon and Newell, John McCarthy felt that machines did not need to simulate the exact mechanisms of human thought, but could instead try to find the essence of abstract reasoning and problem-solving with logic, [ 25 ] regardless of whether people used the same algorithms. [ a ]
His laboratory at Stanford ( SAIL ) focused on using formal logic to solve a wide variety of problems, including knowledge representation , planning and learning . [ 29 ]
Logic was also the focus of the work at the University of Edinburgh and elsewhere in Europe which led to the development of the programming language Prolog and the science of logic programming. [ 30 ] [ 31 ]
Main article: neats vs. scruffies
Researchers at MIT (such as Marvin Minsky and Seymour Papert ) [ 32 ] [ 33 ] [ 6 ] found that solving difficult problems in vision and natural language processing required ad hoc solutions—they argued that no simple and general principle (like logic ) would capture all the aspects of intelligent behavior. Roger Schank described their "anti-logic" approaches as " scruffy " (as opposed to the " neat " paradigms at CMU and Stanford). [ 34 ] [ 35 ]
Commonsense knowledge bases (such as Doug Lenat 's Cyc ) are an example of "scruffy" AI, since they must be built by hand, one complicated concept at a time. [ 36 ] [ 37 ] [ 38 ]
The first AI winter: crushed dreams, 1967–1977
The first AI winter was a shock:
During the first AI summer, many people thought that machine intelligence could be achieved in just a few years. The Defense Advance Research Projects Agency (DARPA) launched programs to support AI research to use AI to solve problems of national security; in particular, to automate the translation of Russian to English for intelligence operations and to create autonomous tanks for the battlefield. Researchers had begun to realize that achieving AI was going to be much harder than was supposed a decade earlier, but a combination of hubris and disingenuousness led many university and think-tank researchers to accept funding with promises of deliverables that they should have known they could not fulfill. By the mid-1960s neither useful natural language translation systems nor autonomous tanks had been created, and a dramatic backlash set in. New DARPA leadership canceled existing AI funding programs.
Outside of the United States, the most fertile ground for AI research was the United Kingdom. The AI winter in the United Kingdom was spurred on not so much by disappointed military leaders as by rival academics who viewed AI researchers as charlatans and a drain on research funding. A professor of applied mathematics, Sir James Lighthill, was commissioned by Parliament to evaluate the state of AI research in the nation . The report stated that all of the problems being worked on in AI would be better handled by res

[truncated]
