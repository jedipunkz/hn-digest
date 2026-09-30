---
source: "https://turkeyland.net/thoughts/ai.php"
hn_url: "https://news.ycombinator.com/item?id=49913458"
title: "CS240 AI Cheating Retrospective"
article_title: "CS 240 Spring 2026 AI Retrospective"
image: ""
author: "ArchAndStarch"
captured_at: "2026-09-30T20:36:18Z"
capture_tool: "hn-digest"
hn_id: 49913458
score: 4
comments: 0
posted_at: "2026-09-30T19:54:29Z"
tags:
  - hacker-news
---

# CS240 AI Cheating Retrospective

- HN: [49913458](https://news.ycombinator.com/item?id=49913458)
- Source: [turkeyland.net](https://turkeyland.net/thoughts/ai.php)
- Score: 4
- Comments: 0
- Posted: 2026-09-30T19:54:29Z

## Translation

Title: CS240 AI Cheating Retrospective
Article title: CS 240 Spring 2026 AI Retrospective
Description: Jeff Turkstra's personal website. Contains photographs, memoirs, TI-86 & TI-89 programs/games, quotes, MIDI's, SeaQuest images, links, and more!

Article text:
Created by Magical Gnomes!
Home
CS 240 Spring 2026 AI Retrospective - September 25, 2026
Toward the end of the Spring 2026 semester, I found myself at the
center of a storm involving use of "AI" in CS 240: Programming in C. My
approach to addressing the issue should have been better and is primarily
why the individuals that ran afoul of the clearly stated course policy
ultimately incurred little to no consequence.
I am writing this post largely due to the number of students that
continue to approach me and mention how much misinformation and inaccuracy
surrounds the, apparently, ongoing discussions regarding the incident.
I am writing this post largely to clearly document the circumstances
surrounding, and resolution to, the CS 240 ordeal last semester.
What follows is a detailed account of the entire ordeal.
To begin, I wish to be abundantly clear: the Spring 2026 offering of
CS 240 had a clearly articulated prohibition with regard to using AI/LLMs
to solve any assignments in the course. You can find it in the
syllabus here. Specifically:
You may discuss assignments in a general way with other students, but
you may not consult anyone else's work. Among other ways to get an F,
you are guilty of academic dishonesty if:
...
- You utilize ChatGPT or other software to programmatically generate
solutions to any part of an assignment, quiz, or exam
Moreover, the syllabus clearly states that we may not immediately address violations of the academic integrity policy:
If we find reason to believe that a student or team has cheated on
any assignment, we may inform the student or team promptly, or we may
decide to silently accumulate evidence against the student or team on
later assignments.
Finally, it was even clearly communicated in lecture that reading and understanding the syllabus was each student's responsibility, as mentioned on the slide from Lecture 1 found below:
I have been told some people may be implying that the expectations above
were somehow a surprise to those enrolled in the course. Such a claim is
highly specious. In addition to the syllabus, I went back and checked: this
expectation was addressed repeatedly in at least five lectures. I
have included some of the slides from those lectures below. Note that
while some may not explicitly mention AI, my slides are a starting
point for what I discuss in lecture and do not encompass everything that
is said. I will add that I am quite certain that I discussed these
expectations at other times but without a slide to explicitly reference
as well.
Here are the slides from Lecture 1 on January 12, 2026:
On January 14, 2026 I discussed the policy again in Lecture 2. While
the slide itself does not mention LLMs, I did take the opportunity to
mention the policy verbally again. Lecture 3 on January 21 included yet
another reminder...
Another discussion occurred in Lecture 7 on February 4, 2026. It was
at this time we were beginning to focus on specific identifiers during
the development of Argus, the tool that would later be deployed.
So, it was always the case that it was made abundantly clear to
students that they should not be using AI/LLMs to complete assignments and
that doing so was a violation of the course academic integrity policy.
Moreover, our usage of MOSS to identify
cheating incidents proceeded as in semesters prior. Throughout the
entire semester, students identified as having highly similar code were
confronted and, when appropriate, incurred consequences resulting from
their violation of the above policy.
Another mistaken impression some seem to have formed is that we
trained our own LLM to assist with identifying students that violated
this policy. We did not. The tool developed - named Argus - utilizes
static analysis to identify indicators for which it is highly unlikely a
reasonable explanation exists with regard to their presence in a student's
source code.
Over the summer, we actually authored a paper regarding
this tool. You can find a preprint version of it on arXiv or locally here . The paper was submitted
to SIGCSE 2027. Though ultimately not selected for publication, the peer
reviews from that process are available
here .
It is important to point out that Argus served only as a starting
point. Each potential case identified was looked at by at least one
person prior to any further action. Not all cases were pursued. In fact,
we took a very conservative look at each case, only pursuing ones for
which there was clear evidence and a lack of any reasonable explanation
for the existence of that evidence. Ultimately we flagged 267 out of
584 identified students (roughly 45.7%).
Argus reached a usable state around mid-March of 2026. We began
discussing it during our weekly TA meetings around that time, and on
March 23 I ultimately decided to form an "AI Academic Integrity" (AIAI)
team to address the large number of potential AI usage cases. We began
efforts in earnest to set up the necessary processes.
Discussions continued after the formation of the AIAI team and, as we
considered division of labor and preparation for meeting with hundreds
of students, I was reminded of the approach historically taken by CS
159. In this course, instructors would email students identified as
having violated course policy and offer them an opportunity to simply
admit responsibility. One of these emails from Fall 2017 is included
below. During our April 13, 2026 meeting I decided to adopt a similar
approach utilizing a form that I would subsequently create with the hope
that a self-reporting mechanism would cut down on the number of cases
that would require further discussion.
Concerns of potential academic integrity violations have been raised
based on your submission for the fifth homework. The solution
you submitted has been measured and determined to be highly similar to
that of at least one other student in the course.
The work you submit must be your own original effort and not the result
of unacceptable, even if unintentional, collaboration. Although it may
not be intentional, extensive collaboration with others may result in
highly similar work that resembles copying. Every student is also
responsible for protecting his/her own work. If you inadvertently allow
your work to be accessed, you may still be held responsible for
facilitating an academic integrity violation.
Please take a moment to review the policies of the course found in your
syllabus related to academic integrity.
In this particular case, the high similarity between your assignment and
others was identified by computer software and upon review cannot be
dismissed by mere coincidence. As a result you will be assigned a score
of zero for the assignment according to the academic integrity policy
outlined in the course syllabus. If you are willing to accept these
findings, please reply to this email by the stated deadline (see below)
containing the following response:
"I acknowledge that my assignment submission is in violation of the
academic integrity policy of the course and accept the score of zero
for this assignment. I have no intentions of doing this again in the
future and realize that a second offense would result in my failing
the class."
After the above statement, you must include a short summary of the
actions and names of others involved that resulted in this incident.
There would then be no further repercussions as a result of this
incident and a favorable report will be sent to the Office of the Dean
of Students stating that you were cooperative and accepting of the
outcome. The case would then be closed as far as this class is
concerned, but the Dean of S
[truncated]
Based on this email (and my use of similar ones in past offerings
of CS 240 as well), I created the form found below. I will also
note that this is where my use of the favorable/unfavorable language
originated. Although some students also seemed at the time to be hung
up on the possibility of suspension or expulsion, that is not something
within an instructor's power nor is it something that typically happens
for a first time offense. I did not mention either possibility in my
communications or on the form, though I believe I briefly mentioned them
in lecture as something the Dean of Students could consider.
On April 16, 2026, I sent out the following email with the subject
"[CS 240] Academic Integrity Violation - RESPONSE REQUIRED":
Through careful analysis and manual review, we have identified what we
consider to be clear and concrete indicators in one or more of your
homework assignment solutions that it was partially or entirely
generated by an AI/LLM tool like ChatGPT, Claude, Copilot, etc. This is
a violation of the course academic integrity policy stated in the
course syllabus and reviewed during the first week of classes.
You are required to complete the form located at:
https://endor.cs.purdue.edu/~cs240ai/index.php
on or before Monday, April 20 at 5:00pm EDT.
Failure to respond will result in a grade of 'F' for the course. Note
that even if you intend to drop the course, we require a response.
Failure to respond will also result in an unfavorable letter being sent
to the Dean of Students along with further potential disciplinary
action.
Prof. Turkstra
--
Dr. Jeffrey A. Turkstra
Teaching Associate Professor
Department of Computer Science
Purdue University
Hall of DS and AI Room 1139E
475 Stadium Mall Drive
West Lafayette, IN 47907
(765) 496-3088
jeff@cs.purdue.edu
https://turkeyland.net/
A screenshot of the linked form is included below:
CS 240 had a beginning enrollment of 599 students for the Spring
2026 semester. At the time that the email was sent out, 65 students
had withdrawn from the course already. The email above was sent to 207
enrolled students in the course. An almost identical email was sent to 60
additional students that had already dropped the course, since letters
would still need to be drafted to the Dean of Students. This represents
a little under 45% of the students.
It is worth noting that the standard penalty outlined in the syllabus for an academic integrity
violation is stated below:
Academic dishonesty is a serious offense which may result in suspension or
expulsion from the University. In addition to any other action taken, such
as suspension or expulsion, a grade of F will normally be recorded on the
transcripts of students found responsible for acts of academic dishonesty.
" grade of F " is bolded in the syllabus. For first time, minor
violations on a single assignment, I typically proceed with the lesser
penalty of a score of 0 and a course letter grade deduction.
As one can see from the form above, I lessened the penalty further
in this situation partly due to the late deployment of the tool. The
consequence would simply have been a 0 for each assignment.
The initial response to this process from my end was surprisingly
positive. Before the university intervened, 117 students submitted the
form. 108 simply took responsibility for their actions. Only 9 disputed
the findings. It is an open question whether this is reflective of their
actual behavior or instead something that underscores the purported,
and unintentional, potentially coercive nature of the form. Ultimately,
144 students submitted enrollment drop requests for the course after this
process began.
I met with well over a dozen students the following day. Every single
one of them was apologetic. A handful of them even thanked me, sharing
with me that they did not truly understand the ramifications of their
actions until I had pursued these cases.
Simultaneously and without my knowledge, contact was being made by
people - some invariably parents and students - both with our department
head as well as the Dean of Students and other units on campus. I am
also told there was significant discussion on the Purdue subreddit,
though I have yet to subject myself to a gossip mill of such scale.
Eventually some of this stuff made it to me - including direct
emails from individuals. Many of the

[truncated]

## Original Extract

Jeff Turkstra's personal website. Contains photographs, memoirs, TI-86 & TI-89 programs/games, quotes, MIDI's, SeaQuest images, links, and more!

Created by Magical Gnomes!
Home
CS 240 Spring 2026 AI Retrospective - September 25, 2026
Toward the end of the Spring 2026 semester, I found myself at the
center of a storm involving use of "AI" in CS 240: Programming in C. My
approach to addressing the issue should have been better and is primarily
why the individuals that ran afoul of the clearly stated course policy
ultimately incurred little to no consequence.
I am writing this post largely due to the number of students that
continue to approach me and mention how much misinformation and inaccuracy
surrounds the, apparently, ongoing discussions regarding the incident.
I am writing this post largely to clearly document the circumstances
surrounding, and resolution to, the CS 240 ordeal last semester.
What follows is a detailed account of the entire ordeal.
To begin, I wish to be abundantly clear: the Spring 2026 offering of
CS 240 had a clearly articulated prohibition with regard to using AI/LLMs
to solve any assignments in the course. You can find it in the
syllabus here. Specifically:
You may discuss assignments in a general way with other students, but
you may not consult anyone else's work. Among other ways to get an F,
you are guilty of academic dishonesty if:
...
- You utilize ChatGPT or other software to programmatically generate
solutions to any part of an assignment, quiz, or exam
Moreover, the syllabus clearly states that we may not immediately address violations of the academic integrity policy:
If we find reason to believe that a student or team has cheated on
any assignment, we may inform the student or team promptly, or we may
decide to silently accumulate evidence against the student or team on
later assignments.
Finally, it was even clearly communicated in lecture that reading and understanding the syllabus was each student's responsibility, as mentioned on the slide from Lecture 1 found below:
I have been told some people may be implying that the expectations above
were somehow a surprise to those enrolled in the course. Such a claim is
highly specious. In addition to the syllabus, I went back and checked: this
expectation was addressed repeatedly in at least five lectures. I
have included some of the slides from those lectures below. Note that
while some may not explicitly mention AI, my slides are a starting
point for what I discuss in lecture and do not encompass everything that
is said. I will add that I am quite certain that I discussed these
expectations at other times but without a slide to explicitly reference
as well.
Here are the slides from Lecture 1 on January 12, 2026:
On January 14, 2026 I discussed the policy again in Lecture 2. While
the slide itself does not mention LLMs, I did take the opportunity to
mention the policy verbally again. Lecture 3 on January 21 included yet
another reminder...
Another discussion occurred in Lecture 7 on February 4, 2026. It was
at this time we were beginning to focus on specific identifiers during
the development of Argus, the tool that would later be deployed.
So, it was always the case that it was made abundantly clear to
students that they should not be using AI/LLMs to complete assignments and
that doing so was a violation of the course academic integrity policy.
Moreover, our usage of MOSS to identify
cheating incidents proceeded as in semesters prior. Throughout the
entire semester, students identified as having highly similar code were
confronted and, when appropriate, incurred consequences resulting from
their violation of the above policy.
Another mistaken impression some seem to have formed is that we
trained our own LLM to assist with identifying students that violated
this policy. We did not. The tool developed - named Argus - utilizes
static analysis to identify indicators for which it is highly unlikely a
reasonable explanation exists with regard to their presence in a student's
source code.
Over the summer, we actually authored a paper regarding
this tool. You can find a preprint version of it on arXiv or locally here . The paper was submitted
to SIGCSE 2027. Though ultimately not selected for publication, the peer
reviews from that process are available
here .
It is important to point out that Argus served only as a starting
point. Each potential case identified was looked at by at least one
person prior to any further action. Not all cases were pursued. In fact,
we took a very conservative look at each case, only pursuing ones for
which there was clear evidence and a lack of any reasonable explanation
for the existence of that evidence. Ultimately we flagged 267 out of
584 identified students (roughly 45.7%).
Argus reached a usable state around mid-March of 2026. We began
discussing it during our weekly TA meetings around that time, and on
March 23 I ultimately decided to form an "AI Academic Integrity" (AIAI)
team to address the large number of potential AI usage cases. We began
efforts in earnest to set up the necessary processes.
Discussions continued after the formation of the AIAI team and, as we
considered division of labor and preparation for meeting with hundreds
of students, I was reminded of the approach historically taken by CS
159. In this course, instructors would email students identified as
having violated course policy and offer them an opportunity to simply
admit responsibility. One of these emails from Fall 2017 is included
below. During our April 13, 2026 meeting I decided to adopt a similar
approach utilizing a form that I would subsequently create with the hope
that a self-reporting mechanism would cut down on the number of cases
that would require further discussion.
Concerns of potential academic integrity violations have been raised
based on your submission for the fifth homework. The solution
you submitted has been measured and determined to be highly similar to
that of at least one other student in the course.
The work you submit must be your own original effort and not the result
of unacceptable, even if unintentional, collaboration. Although it may
not be intentional, extensive collaboration with others may result in
highly similar work that resembles copying. Every student is also
responsible for protecting his/her own work. If you inadvertently allow
your work to be accessed, you may still be held responsible for
facilitating an academic integrity violation.
Please take a moment to review the policies of the course found in your
syllabus related to academic integrity.
In this particular case, the high similarity between your assignment and
others was identified by computer software and upon review cannot be
dismissed by mere coincidence. As a result you will be assigned a score
of zero for the assignment according to the academic integrity policy
outlined in the course syllabus. If you are willing to accept these
findings, please reply to this email by the stated deadline (see below)
containing the following response:
"I acknowledge that my assignment submission is in violation of the
academic integrity policy of the course and accept the score of zero
for this assignment. I have no intentions of doing this again in the
future and realize that a second offense would result in my failing
the class."
After the above statement, you must include a short summary of the
actions and names of others involved that resulted in this incident.
There would then be no further repercussions as a result of this
incident and a favorable report will be sent to the Office of the Dean
of Students stating that you were cooperative and accepting of the
outcome. The case would then be closed as far as this class is
concerned, but the Dean of S
[truncated]
Based on this email (and my use of similar ones in past offerings
of CS 240 as well), I created the form found below. I will also
note that this is where my use of the favorable/unfavorable language
originated. Although some students also seemed at the time to be hung
up on the possibility of suspension or expulsion, that is not something
within an instructor's power nor is it something that typically happens
for a first time offense. I did not mention either possibility in my
communications or on the form, though I believe I briefly mentioned them
in lecture as something the Dean of Students could consider.
On April 16, 2026, I sent out the following email with the subject
"[CS 240] Academic Integrity Violation - RESPONSE REQUIRED":
Through careful analysis and manual review, we have identified what we
consider to be clear and concrete indicators in one or more of your
homework assignment solutions that it was partially or entirely
generated by an AI/LLM tool like ChatGPT, Claude, Copilot, etc. This is
a violation of the course academic integrity policy stated in the
course syllabus and reviewed during the first week of classes.
You are required to complete the form located at:
https://endor.cs.purdue.edu/~cs240ai/index.php
on or before Monday, April 20 at 5:00pm EDT.
Failure to respond will result in a grade of 'F' for the course. Note
that even if you intend to drop the course, we require a response.
Failure to respond will also result in an unfavorable letter being sent
to the Dean of Students along with further potential disciplinary
action.
Prof. Turkstra
--
Dr. Jeffrey A. Turkstra
Teaching Associate Professor
Department of Computer Science
Purdue University
Hall of DS and AI Room 1139E
475 Stadium Mall Drive
West Lafayette, IN 47907
(765) 496-3088
jeff@cs.purdue.edu
https://turkeyland.net/
A screenshot of the linked form is included below:
CS 240 had a beginning enrollment of 599 students for the Spring
2026 semester. At the time that the email was sent out, 65 students
had withdrawn from the course already. The email above was sent to 207
enrolled students in the course. An almost identical email was sent to 60
additional students that had already dropped the course, since letters
would still need to be drafted to the Dean of Students. This represents
a little under 45% of the students.
It is worth noting that the standard penalty outlined in the syllabus for an academic integrity
violation is stated below:
Academic dishonesty is a serious offense which may result in suspension or
expulsion from the University. In addition to any other action taken, such
as suspension or expulsion, a grade of F will normally be recorded on the
transcripts of students found responsible for acts of academic dishonesty.
" grade of F " is bolded in the syllabus. For first time, minor
violations on a single assignment, I typically proceed with the lesser
penalty of a score of 0 and a course letter grade deduction.
As one can see from the form above, I lessened the penalty further
in this situation partly due to the late deployment of the tool. The
consequence would simply have been a 0 for each assignment.
The initial response to this process from my end was surprisingly
positive. Before the university intervened, 117 students submitted the
form. 108 simply took responsibility for their actions. Only 9 disputed
the findings. It is an open question whether this is reflective of their
actual behavior or instead something that underscores the purported,
and unintentional, potentially coercive nature of the form. Ultimately,
144 students submitted enrollment drop requests for the course after this
process began.
I met with well over a dozen students the following day. Every single
one of them was apologetic. A handful of them even thanked me, sharing
with me that they did not truly understand the ramifications of their
actions until I had pursued these cases.
Simultaneously and without my knowledge, contact was being made by
people - some invariably parents and students - both with our department
head as well as the Dean of Students and other units on campus. I am
also told there was significant discussion on the Purdue subreddit,
though I have yet to subject myself to a gossip mill of such scale.
Eventually some of this stuff made it to me - including direct
emails from individuals. Many of the

[truncated]
