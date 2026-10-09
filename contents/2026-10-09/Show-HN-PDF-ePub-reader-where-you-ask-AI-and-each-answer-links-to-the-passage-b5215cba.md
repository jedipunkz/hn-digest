---
source: "https://marginal-ia.app/en"
hn_url: "https://news.ycombinator.com/item?id=50024097"
title: "Show HN: PDF/ePub reader where you ask AI and each answer links to the passage"
article_title: "Marginalia: PDF and EPUB reader with an AI that cites the book"
image: "https://marginal-ia.app/en/opengraph-image-1xsgjv?d804e275bfda1eed"
author: "folio_j"
captured_at: "2026-10-09T18:01:59Z"
capture_tool: "hn-digest"
hn_id: 50024097
score: 2
comments: 0
posted_at: "2026-10-09T17:42:49Z"
tags:
  - hacker-news
---

# Show HN: PDF/ePub reader where you ask AI and each answer links to the passage

- HN: [50024097](https://news.ycombinator.com/item?id=50024097)
- Source: [marginal-ia.app](https://marginal-ia.app/en)
- Score: 2
- Comments: 0
- Posted: 2026-10-09T17:42:49Z

## Translation

Title: Show HN: PDF/ePub reader where you ask AI and each answer links to the passage
Article title: Marginalia: PDF and EPUB reader with an AI that cites the book
Description: PDF and EPUB reader with AI. Open the book in your browser, read it without zooming, even on your phone, ask what you do not understand, and every answer cites the exact passage. No account, nothing to install, the file never uploads.

Article text:
Marginalia: PDF and EPUB reader with an AI that cites the book Loading… — marginalia
Español Sign in Try it free AI reader for PDF and EPUB
Ask your book. The answer takes you to the page.
Open a PDF or an EPUB in your browser, read it without zooming, even on your phone, and ask what you do not understand. Every answer cites the exact passage and one tap takes you there. The file never uploads anywhere.
One physics chapter, 6 MB, a few seconds. The first three questions need no account. Got your own book? Open your own PDF or EPUB no account, nothing to install PDF and EPUB Phone and desktop English and Spanish Free University Physics · 13.7 Einstein's Theory of Gravity Original Text view replay Book Answer Chapter 13 · Gravitation 13.7 Einstein's Theory of Gravity The prediction is that if an object is sufficiently dense, it will collapse in upon itself and be surrounded by an event horizon from which nothing can escape. Karl Schwarzschild was the first person to note this phenomenon in 1916, but at that time, it was considered mostly to be a mathematical curiosity.
The Schwarzschild radius is also called the event horizon of a black hole. We noted that both space and time are stretched near massive objects, such as black holes.
However, if a neutron star gains additional mass, it would eventually collapse, shrinking beyond the Schwarzschild radius. Once that happens, the entire mass would be pulled, inevitably, to a singularity.
In the diagram, space is stretched to infinity. Time is also stretched to infinity. As objects fall toward the event horizon, we see them approaching ever more slowly, but never reaching the event horizon. As outside observers, we never see objects pass through the event horizon: effectively, time is stretched to a stop.
University Physics, OpenStax · CC BY-NC-SA Answers the app gave on this chapter. How it works From the book to the answer, and back to the book.
It is not a chat next to a PDF. It is a reader with the AI in the margin: you read, you ask, you check , all without leaving the page.
Open a PDF or an EPUB from your device. It is saved in your browser and never uploaded anywhere. Nothing to install.
Ask what you do not understand
About any part of the book or the paragraph you select, in your language. The AI answers with the book's passages in front of it and tells you when something is not in it.
Every answer carries its citations. Tap a number and the book jumps to the passage, highlighted. If the AI gets it wrong, you see it right there.
A PDF that finally reads well on your phone.
No zooming, no dragging from side to side, no words cut in half. Text view reflows the PDF so the text fits your screen and the font size you choose. The original page is one tap away.
Two columns, tables and figures where they belong
Formulas and code blocks preserved
Light, sepia or dark theme, and the typeface you prefer
A table of contents even when the PDF has none
Try it: tap “Text view” and watch how the very same page adapts to your screen.
Original Text view Size − 16 + Theme a a a Chapter 4 · Processes
The scheduler decides which process gets the CPU at every instant. The simplest policy serves in order of arrival, and its flaw is well known: a long process holds the CPU while short ones wait behind it, and the average response time explodes.
Fig. 4.3. Circular queue: the process returns to the back after its quantum. The classic alternative splits time into quanta. Each process gets a slice and goes back to the queue. The cost is the context switch, which is not free: save registers, invalidate the cache, restore the memory map. Measuring that cost precisely requires hardware counters, because the system clock is far too coarse for slices of a few microseconds. A real sched -
uler combines both ideas with priority queues that age, so no process is left without a turn indefinitely. The core loop fits in a few lines:
while (!queue_empty(&rq)) {
task_t *t = dequeue(&rq);
run_for(t, QUANTUM_US);
if (t->state == RUNNABLE)
enqueue(&rq, age(t));
} With aging, a process's effective priority grows with the time it has been waiting, which bounds the maximum latency to a multiple of the quantum.
Remember that each file in your working directory can be in one of two states: tracked or untracked. Tracked files are files that were in the last snapshot; they can be unmodified, modified, or staged. Untracked files are everything else.
When you first clone a repository, all of your files will be tracked and unmodified because Git just checked them out and you haven't edited anything.
Marginalia: notes in the margin, in the book itself.
Highlight in four colors, write a note next to it and ask from the highlight : the conversation stays linked to the passage. With an account, progress, highlights, notes and conversations wait for you on your phone.
Change the font size or the theme and every highlight stays on its sentence.
Your file never uploads. Ever.
The file and the index for searching inside it stay in your browser. When you ask, only your question and the passages needed to answer it travel, and the answer comes back with its citations. With an account, the book's text also passes through the server to build the search index, and it is not stored.
The trade-off: with an account, on another device you will see the book in your library with your notes inside, but you need to open the file there once to keep reading.
Your browser anatomy.pdf stays
question + passages answer + citations the file, never Server The AI reads the passages only those
Answers with citations without storing the book
A two-column paper, a statute, a programming manual or a textbook. Pick your field and try it with a real document like the ones you will open.
A textbook or a two-column paper, readable on your phone. Ask about a clinical picture and the answer takes you to the passage that explains it.
Open the statute or the syllabus and ask about an article: the answer cites the text and takes you to it so you can check.
Specifications, RFCs and technical books with syntax colors. Select a block of code and ask what it does.
Formulas are preserved in text view and, with an account, you can tap one to ask about it.
swipe Questions Before you open your first book.
No. You open the file and read. You can ask three questions without an account; to go on, create a free one with just your email. Is it free? Yes, for now. Each account has a daily and monthly question limit, generous for normal use. Any change will be announced here before it takes effect. What kinds of files can I open? EPUB and PDF. If the PDF is a scan (a photo of each page), you can view it as it is, but you can't switch it to text view or ask it questions. Is my book uploaded to a server? No. The file stays on your device. When you ask, your question and the passages of the book needed to answer it are sent. With an account, the book's text also passes through the server to build the search index, without being stored, and your account keeps your highlights with their quotes and your conversations, never the book itself. Do you record what I do on screen? The shape of the session is recorded, without your name: where you tap, which panels you open, how long things take to load. Not the text: the book shows up as a grey block, and your questions, the answers, your notes and the titles are covered. It shows where people get stuck. Turn it off with “Don't record my session” at the foot of the page or in the reader settings, and nothing is recorded if your browser asks not to be tracked. Can I keep on reading from my phone? Yes. With an account, your highlights, notes, conversations and the page you were on are waiting on your phone. Open the file there once and pick up where you left off. Can the AI get it wrong? Yes, it can make mistakes even with the book in front of it. That is why every answer comes with its citations: tap a number to check the original passage. Why does Marginalia exist? It started with a technical book and a phone beside it: a photo of the page, the photo sent to an AI and, from there, a conversation that truly helped. But the book and the conversation lived in separate places, the AI couldn't quote, and coming back meant hunting for the pa
[truncated]
Start now, no sign-up. If Marginalia becomes part of your routine, install it as an app and your books wait for you in the library.

## Original Extract

PDF and EPUB reader with AI. Open the book in your browser, read it without zooming, even on your phone, ask what you do not understand, and every answer cites the exact passage. No account, nothing to install, the file never uploads.

Marginalia: PDF and EPUB reader with an AI that cites the book Loading… — marginalia
Español Sign in Try it free AI reader for PDF and EPUB
Ask your book. The answer takes you to the page.
Open a PDF or an EPUB in your browser, read it without zooming, even on your phone, and ask what you do not understand. Every answer cites the exact passage and one tap takes you there. The file never uploads anywhere.
One physics chapter, 6 MB, a few seconds. The first three questions need no account. Got your own book? Open your own PDF or EPUB no account, nothing to install PDF and EPUB Phone and desktop English and Spanish Free University Physics · 13.7 Einstein's Theory of Gravity Original Text view replay Book Answer Chapter 13 · Gravitation 13.7 Einstein's Theory of Gravity The prediction is that if an object is sufficiently dense, it will collapse in upon itself and be surrounded by an event horizon from which nothing can escape. Karl Schwarzschild was the first person to note this phenomenon in 1916, but at that time, it was considered mostly to be a mathematical curiosity.
The Schwarzschild radius is also called the event horizon of a black hole. We noted that both space and time are stretched near massive objects, such as black holes.
However, if a neutron star gains additional mass, it would eventually collapse, shrinking beyond the Schwarzschild radius. Once that happens, the entire mass would be pulled, inevitably, to a singularity.
In the diagram, space is stretched to infinity. Time is also stretched to infinity. As objects fall toward the event horizon, we see them approaching ever more slowly, but never reaching the event horizon. As outside observers, we never see objects pass through the event horizon: effectively, time is stretched to a stop.
University Physics, OpenStax · CC BY-NC-SA Answers the app gave on this chapter. How it works From the book to the answer, and back to the book.
It is not a chat next to a PDF. It is a reader with the AI in the margin: you read, you ask, you check , all without leaving the page.
Open a PDF or an EPUB from your device. It is saved in your browser and never uploaded anywhere. Nothing to install.
Ask what you do not understand
About any part of the book or the paragraph you select, in your language. The AI answers with the book's passages in front of it and tells you when something is not in it.
Every answer carries its citations. Tap a number and the book jumps to the passage, highlighted. If the AI gets it wrong, you see it right there.
A PDF that finally reads well on your phone.
No zooming, no dragging from side to side, no words cut in half. Text view reflows the PDF so the text fits your screen and the font size you choose. The original page is one tap away.
Two columns, tables and figures where they belong
Formulas and code blocks preserved
Light, sepia or dark theme, and the typeface you prefer
A table of contents even when the PDF has none
Try it: tap “Text view” and watch how the very same page adapts to your screen.
Original Text view Size − 16 + Theme a a a Chapter 4 · Processes
The scheduler decides which process gets the CPU at every instant. The simplest policy serves in order of arrival, and its flaw is well known: a long process holds the CPU while short ones wait behind it, and the average response time explodes.
Fig. 4.3. Circular queue: the process returns to the back after its quantum. The classic alternative splits time into quanta. Each process gets a slice and goes back to the queue. The cost is the context switch, which is not free: save registers, invalidate the cache, restore the memory map. Measuring that cost precisely requires hardware counters, because the system clock is far too coarse for slices of a few microseconds. A real sched -
uler combines both ideas with priority queues that age, so no process is left without a turn indefinitely. The core loop fits in a few lines:
while (!queue_empty(&rq)) {
task_t *t = dequeue(&rq);
run_for(t, QUANTUM_US);
if (t->state == RUNNABLE)
enqueue(&rq, age(t));
} With aging, a process's effective priority grows with the time it has been waiting, which bounds the maximum latency to a multiple of the quantum.
Remember that each file in your working directory can be in one of two states: tracked or untracked. Tracked files are files that were in the last snapshot; they can be unmodified, modified, or staged. Untracked files are everything else.
When you first clone a repository, all of your files will be tracked and unmodified because Git just checked them out and you haven't edited anything.
Marginalia: notes in the margin, in the book itself.
Highlight in four colors, write a note next to it and ask from the highlight : the conversation stays linked to the passage. With an account, progress, highlights, notes and conversations wait for you on your phone.
Change the font size or the theme and every highlight stays on its sentence.
Your file never uploads. Ever.
The file and the index for searching inside it stay in your browser. When you ask, only your question and the passages needed to answer it travel, and the answer comes back with its citations. With an account, the book's text also passes through the server to build the search index, and it is not stored.
The trade-off: with an account, on another device you will see the book in your library with your notes inside, but you need to open the file there once to keep reading.
Your browser anatomy.pdf stays
question + passages answer + citations the file, never Server The AI reads the passages only those
Answers with citations without storing the book
A two-column paper, a statute, a programming manual or a textbook. Pick your field and try it with a real document like the ones you will open.
A textbook or a two-column paper, readable on your phone. Ask about a clinical picture and the answer takes you to the passage that explains it.
Open the statute or the syllabus and ask about an article: the answer cites the text and takes you to it so you can check.
Specifications, RFCs and technical books with syntax colors. Select a block of code and ask what it does.
Formulas are preserved in text view and, with an account, you can tap one to ask about it.
swipe Questions Before you open your first book.
No. You open the file and read. You can ask three questions without an account; to go on, create a free one with just your email. Is it free? Yes, for now. Each account has a daily and monthly question limit, generous for normal use. Any change will be announced here before it takes effect. What kinds of files can I open? EPUB and PDF. If the PDF is a scan (a photo of each page), you can view it as it is, but you can't switch it to text view or ask it questions. Is my book uploaded to a server? No. The file stays on your device. When you ask, your question and the passages of the book needed to answer it are sent. With an account, the book's text also passes through the server to build the search index, without being stored, and your account keeps your highlights with their quotes and your conversations, never the book itself. Do you record what I do on screen? The shape of the session is recorded, without your name: where you tap, which panels you open, how long things take to load. Not the text: the book shows up as a grey block, and your questions, the answers, your notes and the titles are covered. It shows where people get stuck. Turn it off with “Don't record my session” at the foot of the page or in the reader settings, and nothing is recorded if your browser asks not to be tracked. Can I keep on reading from my phone? Yes. With an account, your highlights, notes, conversations and the page you were on are waiting on your phone. Open the file there once and pick up where you left off. Can the AI get it wrong? Yes, it can make mistakes even with the book in front of it. That is why every answer comes with its citations: tap a number to check the original passage. Why does Marginalia exist? It started with a technical book and a phone beside it: a photo of the page, the photo sent to an AI and, from there, a conversation that truly helped. But the book and the conversation lived in separate places, the AI couldn't quote, and coming back meant hunting for the pa
[truncated]
Start now, no sign-up. If Marginalia becomes part of your routine, install it as an app and your books wait for you in the library.
