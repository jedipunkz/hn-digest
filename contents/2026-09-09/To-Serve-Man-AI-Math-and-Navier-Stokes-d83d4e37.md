---
source: "https://ml5885.github.io/writing/navier-stokes.html"
hn_url: "https://news.ycombinator.com/item?id=49629132"
title: "To Serve Man: AI, Math, and Navier–Stokes"
article_title: "To Serve Man - Michael Li"
image: ""
author: "bearseascape"
captured_at: "2026-09-09T17:03:52Z"
capture_tool: "hn-digest"
hn_id: 49629132
score: 2
comments: 1
posted_at: "2026-09-09T16:28:57Z"
tags:
  - hacker-news
---

# To Serve Man: AI, Math, and Navier–Stokes

- HN: [49629132](https://news.ycombinator.com/item?id=49629132)
- Source: [ml5885.github.io](https://ml5885.github.io/writing/navier-stokes.html)
- Score: 2
- Comments: 1
- Posted: 2026-09-09T16:28:57Z

## Translation

Title: To Serve Man: AI, Math, and Navier–Stokes
Article title: To Serve Man - Michael Li

Article text:
By now, most people have heard that OpenAI has likely solved a Millennium Prize Problem ; specifically, a proof
of finite-time blowup for the forced three-dimensional
Navier-Stokes equations, announced yesterday afternoon
at 1:20 p.m. ET. @OpenAI , September 8, 2026,
1:20 p.m. ET As someone who has followed the story
somewhat closely, along with the broader debate around AI
in mathematics, I wanted to write a short blog post that
lays out a timeline of events, and (hopefully) contextualizes
this result with developments in the past year.
Towards the beginning of 2026, AI in mathematics was
still treated as somewhat of a novelty. There were some
interesting results, mostly on Erdős problems, mostly
with commercially available AI models, and mostly from a
few people Primarily @AcerFur
and @Liam06972452 . on
Twitter. Mathematicians were paying attention, though
without much alarm. For example, after one such problem
was solved in January, Terence Tao the Fields Medal-winning,
MacArthur-recognized, field-defining
mathematician ,
who kept a page tracking AI-assisted mathematical
contributions, AI contributions to Erdős
problems
remarked that he was less interested in the solution
itself than in "the emerging AI-powered capability to
rapidly write and rewrite expositions of the solution." @tao , January 7,
2026.
The results being produced began to look like less of a
novelty in May, when OpenAI announced that an internal
model had disproved the Erdős unit distance conjecture, An OpenAI model has disproved a
central conjecture in discrete geometry , May 20,
2026. an open problem in discrete geometry. This
result was supplemented by a companion paper, written by
nine mathematicians, that digested the argument and
simplified it. Alon et al., Remarks on the disproof of the unit
distance conjecture , May 20, 2026. This
seemed like an encouraging model for AI-assisted
mathematics: the model produced a result, while
mathematicians worked together to make it legible and
useful to the field. The proof was
treated not as the end of the process, but as the beginning
of one. This sort of human-centered focus, however, did not
last very long.
By the summer, the character of the results being posted
had changed quite a bit. On July 19, Levent Alpöge (a
member of technical staff at Anthropic) posted a
counterexample to the Jacobian conjecture: @__alpoge__ , July 19,
2026.
hello there the jacobian conjecture is false
thanx to my close friend akhil for asking about it and
my other close friend fable for working during the world
cup final
((1+xy)^3 z + y^2 (1+xy) (4+3xy), y + 3
x (1+xy)^2 z + 3 x y^2 (4+3xy), 2 x - 3 x^2 y - x^3 z):
\C^3\to \C^3, has jacobian determinant -2, and sends (0,
0, -1/4), (1, -3/2, 13/2), and (-1, 3/2, 13/2) to (-1/4,
0, 0)
Notably, this was a black-box proof: although the
counterexample itself was easy to verify as correct, no
paper accompanied it, nor was there any explanation of how
it had been found. A better understanding emerged only
afterward, when other mathematicians (with assistance from
AI models) @davikrehalt , July 20,
2026.
worked to find one. Subsequently, Alpöge posted a resolution
of the Hadamard matrix conjecture on August 12, @__alpoge__ , August 12,
2026.
and, on August 23, a proof that $S^6$ admits a complex
structure. @__alpoge__ , August 23,
2026.
The labs followed a similar trend. On August 1, OpenAI
published ten results attributed to an internal version
of Astra, Ten advances in mathematics and
theoretical computer science , August 1,
2026.
among them the first explicit construction of a
non-sofic group. This construction (and the nine other
results) did not come with a companion paper written by
mathematicians with domain expertise, the way the unit
distance result had been. Instead, the only proofs
provided were entirely AI-written. Again,
understanding had to come from outside. Francesco
Fournier-Facio, a mathematician at Cambridge,
independently worked out the criterion underneath the
OpenAI results and used that to produce a broader family
of examples, including torsion-free ones A torsion-free non-sofic group ,
August 3, 2026. Similarly, on August 10,
Anthropic announced that an unreleased model had raised
the lower bound on the proportion of zeros of the zeta
function lying on the critical line from 41.6% to 67.2%,
after a member of technical staff asked it to "take a
real stab" at the Riemann hypothesis and then mostly
sent it words of encouragement. Learning more about Claude's
mathematical capabilities , August 10,
2026.
Like several of the other recent announcements, this one
came without any companion paper aimed at helping
mathematicians understand the argument beyond simply
assessing its correctness.
Around this time, some mathematicians started pushing
back. Their issue wasn't that the proofs were wrong; in
fact, most were accompanied by (compiling) Lean
certificates. One complaint was the writing. In a talk
at IPAM, Accelerating Math and Theoretical
Physics with AI , IPAM, UCLA. Tao
described
how AI-generated proofs would spend pages on trivial
details and then a single line on the key idea. The
larger complaint was that mathematics was being used as
a benchmark for AI companies, strip-mined @tao , September 2, 2026.
for hype moments and aura by people with no particular
interest in the field itself. One such critique, written
by Hugo Duminil-Copin, A recipient of the 2022 Fields
Medal. in a blog post sarcastically titled "Care
for a little more AI?", lamented that: Care for a little more AI? ,
Proofs and Prompts, August 30, 2026.
...the current use of AI does not empower us, it
petrifies us. These artificial discoveries risk
decapitating entire fields before they have had time to
develop to their full potential. Even worse, they nuke
the mathematical landscape, making it increasingly
difficult to inhabit after each blast.
Terence Tao (previously a vocal proponent of AI in math) How Terry Tao Became an Evangelist
for AI in Math , Quanta, June 8, 2026.
soon
began making a similar case. In a sequence of posts on
Mathstodon, written over four days from the evening of
September 2 to the afternoon of September 5, he
developed his argument.
In the first of these posts, Tao described the idea of
contamination, A term from AI evaluation, where a
problem whose solution is already public is no longer
useful as a test. cited Duminil-Copin's article,
and suggested that: @tao , September 2, 2026, 11:15 p.m.
ET.
It may become necessary to declare certain
classes of mathematical problems off-limits to automated
solvers, in order to preserve their broader value to the
mathematical ecosystem (for instance, through the
training of future mathematicians). Admittedly this can
be hard to enforce when such tools are both powerful and
widely available.
In a second post (whose importance will be obvious later)
he gave a concrete example. Using the famous
Navier-Stokes equations, Tao explained what could be
lost if an important problem like this were solved
without any insight into the process that produced the
answer: @tao , September 3, 2026, 11:37 a.m.
ET.
[T]here is now a scenario in which an autonomous
AI harness, backed by an enormous amount of
computational resources, performs this entire iteration
internally, and ends up producing the final ansatz, and
thence the solution to the Navier-Stokes regularity
problem, while the AI company running the harness keeps
the process to arrive at that ansatz almost completely
out of public view. Technically, one of the most
prominent open problems in mathematics would now be
solved; but there would be almost no value added to
mathematics as a consequence.
Later that day, he put the concern more generally: @tao , September 3, 2026, 2:53 p.m.
ET.
But the currently fashionable practice of
pointing a powerful AI tool at the task of answering a
problem $X$, unguided by any human expert in the field
$X$ resides in, has created an unprecedented divergence
between the production of answers, and the production of
insight, to the point where the two questions have
become negatively correlated .
He followed this up with a case study that had unfolded a
few days earlier. On August 31, Julia Stadlmann J.L. Doob Research Assistant
Professor at the University of Illinois. released
a preprint lowering the best known bound on gaps between
primes. arXiv:2608.31126 , August 31,
2026. Tao described how, immediately after
this: @tao , September 3, 2026, 2:53 p.m.
ET.
...we were treated to the unedifying spectacle
of no fewer than three separate AI companies racing to
announce their own improvement on the bounded gaps
between prime result that I mentioned yesterday. These
results are numerically stronger than Stadlmann's
improvement, but I am very glad that Stadlmann was able
to complete her analysis just in time before the problem
became contaminated [...]
Unbeknownst to Tao, prior to his post on Navier-Stokes,
rumors had already been circulating about an AI lab
solving two Millennium Prize Problems. Motivated by
these rumors, and Tao's post, Andrew Curran tweeted the
following: @AndrewCurran_ , September 4,
2026, 10:26 p.m. ET.
It's fun to make predictions. Here's a new one:
Anthropic has solved a Millennium Prize Problem. And
I'll be even more specific. Claude has solved
Navier-Stokes. It is out for expert review. And to give
myself a hard deadline, they will announce it before the
IPO.
The next morning, Tao added a clarification to his own
thread: @tao , September 5, 2026, 11:19 a.m.
ET.
[I]n response to recent rumors about a possible
solution to the Navier-Stokes problem: I am not aware of
any significant developments in this regard; the above
discussion is hypothetical, but not completely
implausible at the current level of development of AI
technology.
The rumors continued regardless.
Let ten thousand agents bloom!
At 11:58 p.m. ET on Monday, September 7, Tristan Buckmaster Professor of mathematics at NYU's
Courant Institute. posted three results, in a
joint collaboration with Levent Alpöge: finite-time
blowup with smooth forcing for incompressible porous
media, for Boussinesq, and for 3D incompressible Euler,
with Lean formalizations for each. @tristanbuckmaster , September 7,
2026, 11:58 p.m. ET. Alongside them, he posted a
four-page statement, statement.pdf . explaining
the following.
On September 3, with the rumors circulating, Buckmaster
emailed a mathematician at OpenAI to correct the record.
Over the following days, he was asked with increasing
urgency to meet. On a call on Sunday, September 6, he was
told
that an internal OpenAI model had proved finite-time
blowup for forced Navier-Stokes:
I asked when the first prompt was sent by them.
This question was not answered directly by OpenAI for
some time. Eventually it was agreed that it had been
sent in the past few days, after information about our
work had reached OpenAI. I asked whether the model had
been trained on, or had access to, our sessions in
Codex, into which we had been putting all our drafts for
the whole of this project. I was told the model did not
look up user data. I asked again, about training, and I
did not get an answer. OpenAI says its researchers "did not
see any of their work through any means until they
released it publicly," but adds that "while unlikely, we
cannot rule out that de-identified data derived from
their usage of our products helped improve our models."
On the Navier-Stokes Millennium
Prize Problem .
He was then offered two proposals:
Two proposals were offered to me. The first was
that we post our Euler result, and that OpenAI post its
Navier-Stokes result the next day. The second was that,
after posting Euler, I alone write a paper presenting
the Navier-Stokes result, acknowledging that an internal
OpenAI model had resolved it. Sebastien twice asserted
that he wanted Levent removed from authorship, and said
it would all be simple if only it were not the case
that, and it was so annoying that, Levent works at
Anthropic. [...] I

[truncated]

## Original Extract

By now, most people have heard that OpenAI has likely solved a Millennium Prize Problem ; specifically, a proof
of finite-time blowup for the forced three-dimensional
Navier-Stokes equations, announced yesterday afternoon
at 1:20 p.m. ET. @OpenAI , September 8, 2026,
1:20 p.m. ET As someone who has followed the story
somewhat closely, along with the broader debate around AI
in mathematics, I wanted to write a short blog post that
lays out a timeline of events, and (hopefully) contextualizes
this result with developments in the past year.
Towards the beginning of 2026, AI in mathematics was
still treated as somewhat of a novelty. There were some
interesting results, mostly on Erdős problems, mostly
with commercially available AI models, and mostly from a
few people Primarily @AcerFur
and @Liam06972452 . on
Twitter. Mathematicians were paying attention, though
without much alarm. For example, after one such problem
was solved in January, Terence Tao the Fields Medal-winning,
MacArthur-recognized, field-defining
mathematician ,
who kept a page tracking AI-assisted mathematical
contributions, AI contributions to Erdős
problems
remarked that he was less interested in the solution
itself than in "the emerging AI-powered capability to
rapidly write and rewrite expositions of the solution." @tao , January 7,
2026.
The results being produced began to look like less of a
novelty in May, when OpenAI announced that an internal
model had disproved the Erdős unit distance conjecture, An OpenAI model has disproved a
central conjecture in discrete geometry , May 20,
2026. an open problem in discrete geometry. This
result was supplemented by a companion paper, written by
nine mathematicians, that digested the argument and
simplified it. Alon et al., Remarks on the disproof of the unit
distance conjecture , May 20, 2026. This
seemed like an encouraging model for AI-assisted
mathematics: the model produced a result, while
mathematicians worked together to make it legible and
useful to the field. The proof was
treated not as the end of the process, but as the beginning
of one. This sort of human-centered focus, however, did not
last very long.
By the summer, the character of the results being posted
had changed quite a bit. On July 19, Levent Alpöge (a
member of technical staff at Anthropic) posted a
counterexample to the Jacobian conjecture: @__alpoge__ , July 19,
2026.
hello there the jacobian conjecture is false
thanx to my close friend akhil for asking about it and
my other close friend fable for working during the world
cup final
((1+xy)^3 z + y^2 (1+xy) (4+3xy), y + 3
x (1+xy)^2 z + 3 x y^2 (4+3xy), 2 x - 3 x^2 y - x^3 z):
\C^3\to \C^3, has jacobian determinant -2, and sends (0,
0, -1/4), (1, -3/2, 13/2), and (-1, 3/2, 13/2) to (-1/4,
0, 0)
Notably, this was a black-box proof: although the
counterexample itself was easy to verify as correct, no
paper accompanied it, nor was there any explanation of how
it had been found. A better understanding emerged only
afterward, when other mathematicians (with assistance from
AI models) @davikrehalt , July 20,
2026.
worked to find one. Subsequently, Alpöge posted a resolution
of the Hadamard matrix conjecture on August 12, @__alpoge__ , August 12,
2026.
and, on August 23, a proof that $S^6$ admits a complex
structure. @__alpoge__ , August 23,
2026.
The labs followed a similar trend. On August 1, OpenAI
published ten results attributed to an internal version
of Astra, Ten advances in mathematics and
theoretical computer science , August 1,
2026.
among them the first explicit construction of a
non-sofic group. This construction (and the nine other
results) did not come with a companion paper written by
mathematicians with domain expertise, the way the unit
distance result had been. Instead, the only proofs
provided were entirely AI-written. Again,
understanding had to come from outside. Francesco
Fournier-Facio, a mathematician at Cambridge,
independently worked out the criterion underneath the
OpenAI results and used that to produce a broader family
of examples, including torsion-free ones A torsion-free non-sofic group ,
August 3, 2026. Similarly, on August 10,
Anthropic announced that an unreleased model had raised
the lower bound on the proportion of zeros of the zeta
function lying on the critical line from 41.6% to 67.2%,
after a member of technical staff asked it to "take a
real stab" at the Riemann hypothesis and then mostly
sent it words of encouragement. Learning more about Claude's
mathematical capabilities , August 10,
2026.
Like several of the other recent announcements, this one
came without any companion paper aimed at helping
mathematicians understand the argument beyond simply
assessing its correctness.
Around this time, some mathematicians started pushing
back. Their issue wasn't that the proofs were wrong; in
fact, most were accompanied by (compiling) Lean
certificates. One complaint was the writing. In a talk
at IPAM, Accelerating Math and Theoretical
Physics with AI , IPAM, UCLA. Tao
described
how AI-generated proofs would spend pages on trivial
details and then a single line on the key idea. The
larger complaint was that mathematics was being used as
a benchmark for AI companies, strip-mined @tao , September 2, 2026.
for hype moments and aura by people with no particular
interest in the field itself. One such critique, written
by Hugo Duminil-Copin, A recipient of the 2022 Fields
Medal. in a blog post sarcastically titled "Care
for a little more AI?", lamented that: Care for a little more AI? ,
Proofs and Prompts, August 30, 2026.
...the current use of AI does not empower us, it
petrifies us. These artificial discoveries risk
decapitating entire fields before they have had time to
develop to their full potential. Even worse, they nuke
the mathematical landscape, making it increasingly
difficult to inhabit after each blast.
Terence Tao (previously a vocal proponent of AI in math) How Terry Tao Became an Evangelist
for AI in Math , Quanta, June 8, 2026.
soon
began making a similar case. In a sequence of posts on
Mathstodon, written over four days from the evening of
September 2 to the afternoon of September 5, he
developed his argument.
In the first of these posts, Tao described the idea of
contamination, A term from AI evaluation, where a
problem whose solution is already public is no longer
useful as a test. cited Duminil-Copin's article,
and suggested that: @tao , September 2, 2026, 11:15 p.m.
ET.
It may become necessary to declare certain
classes of mathematical problems off-limits to automated
solvers, in order to preserve their broader value to the
mathematical ecosystem (for instance, through the
training of future mathematicians). Admittedly this can
be hard to enforce when such tools are both powerful and
widely available.
In a second post (whose importance will be obvious later)
he gave a concrete example. Using the famous
Navier-Stokes equations, Tao explained what could be
lost if an important problem like this were solved
without any insight into the process that produced the
answer: @tao , September 3, 2026, 11:37 a.m.
ET.
[T]here is now a scenario in which an autonomous
AI harness, backed by an enormous amount of
computational resources, performs this entire iteration
internally, and ends up producing the final ansatz, and
thence the solution to the Navier-Stokes regularity
problem, while the AI company running the harness keeps
the process to arrive at that ansatz almost completely
out of public view. Technically, one of the most
prominent open problems in mathematics would now be
solved; but there would be almost no value added to
mathematics as a consequence.
Later that day, he put the concern more generally: @tao , September 3, 2026, 2:53 p.m.
ET.
But the currently fashionable practice of
pointing a powerful AI tool at the task of answering a
problem $X$, unguided by any human expert in the field
$X$ resides in, has created an unprecedented divergence
between the production of answers, and the production of
insight, to the point where the two questions have
become negatively correlated .
He followed this up with a case study that had unfolded a
few days earlier. On August 31, Julia Stadlmann J.L. Doob Research Assistant
Professor at the University of Illinois. released
a preprint lowering the best known bound on gaps between
primes. arXiv:2608.31126 , August 31,
2026. Tao described how, immediately after
this: @tao , September 3, 2026, 2:53 p.m.
ET.
...we were treated to the unedifying spectacle
of no fewer than three separate AI companies racing to
announce their own improvement on the bounded gaps
between prime result that I mentioned yesterday. These
results are numerically stronger than Stadlmann's
improvement, but I am very glad that Stadlmann was able
to complete her analysis just in time before the problem
became contaminated [...]
Unbeknownst to Tao, prior to his post on Navier-Stokes,
rumors had already been circulating about an AI lab
solving two Millennium Prize Problems. Motivated by
these rumors, and Tao's post, Andrew Curran tweeted the
following: @AndrewCurran_ , September 4,
2026, 10:26 p.m. ET.
It's fun to make predictions. Here's a new one:
Anthropic has solved a Millennium Prize Problem. And
I'll be even more specific. Claude has solved
Navier-Stokes. It is out for expert review. And to give
myself a hard deadline, they will announce it before the
IPO.
The next morning, Tao added a clarification to his own
thread: @tao , September 5, 2026, 11:19 a.m.
ET.
[I]n response to recent rumors about a possible
solution to the Navier-Stokes problem: I am not aware of
any significant developments in this regard; the above
discussion is hypothetical, but not completely
implausible at the current level of development of AI
technology.
The rumors continued regardless.
Let ten thousand agents bloom!
At 11:58 p.m. ET on Monday, September 7, Tristan Buckmaster Professor of mathematics at NYU's
Courant Institute. posted three results, in a
joint collaboration with Levent Alpöge: finite-time
blowup with smooth forcing for incompressible porous
media, for Boussinesq, and for 3D incompressible Euler,
with Lean formalizations for each. @tristanbuckmaster , September 7,
2026, 11:58 p.m. ET. Alongside them, he posted a
four-page statement, statement.pdf . explaining
the following.
On September 3, with the rumors circulating, Buckmaster
emailed a mathematician at OpenAI to correct the record.
Over the following days, he was asked with increasing
urgency to meet. On a call on Sunday, September 6, he was
told
that an internal OpenAI model had proved finite-time
blowup for forced Navier-Stokes:
I asked when the first prompt was sent by them.
This question was not answered directly by OpenAI for
some time. Eventually it was agreed that it had been
sent in the past few days, after information about our
work had reached OpenAI. I asked whether the model had
been trained on, or had access to, our sessions in
Codex, into which we had been putting all our drafts for
the whole of this project. I was told the model did not
look up user data. I asked again, about training, and I
did not get an answer. OpenAI says its researchers "did not
see any of their work through any means until they
released it publicly," but adds that "while unlikely, we
cannot rule out that de-identified data derived from
their usage of our products helped improve our models."
On the Navier-Stokes Millennium
Prize Problem .
He was then offered two proposals:
Two proposals were offered to me. The first was
that we post our Euler result, and that OpenAI post its
Navier-Stokes result the next day. The second was that,
after posting Euler, I alone write a paper presenting
the Navier-Stokes result, acknowledging that an internal
OpenAI model had resolved it. Sebastien twice asserted
that he wanted Levent removed from authorship, and said
it would all be simple if only it were not the case
that, and it was so annoying that, Levent works at
Anthropic. [...] I

[truncated]
