---
source: "https://github.com/demetrius-edelin/tutor"
hn_url: "https://news.ycombinator.com/item?id=50036777"
title: "Show HN: Tutor, a self-hosted LLM tutor that teaches you your books or docs"
article_title: "GitHub - demetrius-edelin/tutor: Self-hosted AI tutor (BYOK) that turns your own EPUB and PDF books into a course. It maps the concepts, tests what you know, and teaches the rest with references to the book. · GitHub"
image: "https://opengraph.githubassets.com/59a268b51e1556d92638432cab25583664189d14f946d34dce8607d83493a7c8/demetrius-edelin/tutor"
author: "edelind"
captured_at: "2026-10-10T20:29:26Z"
capture_tool: "hn-digest"
hn_id: 50036777
score: 1
comments: 0
posted_at: "2026-10-10T20:23:47Z"
tags:
  - hacker-news
---

# Show HN: Tutor, a self-hosted LLM tutor that teaches you your books or docs

- HN: [50036777](https://news.ycombinator.com/item?id=50036777)
- Source: [github.com](https://github.com/demetrius-edelin/tutor)
- Score: 1
- Comments: 0
- Posted: 2026-10-10T20:23:47Z

## Translation

Title: Show HN: Tutor, a self-hosted LLM tutor that teaches you your books or docs
Article title: GitHub - demetrius-edelin/tutor: Self-hosted AI tutor (BYOK) that turns your own EPUB and PDF books into a course. It maps the concepts, tests what you know, and teaches the rest with references to the book. · GitHub
Description: Self-hosted AI tutor (BYOK) that turns your own EPUB and PDF books into a course. It maps the concepts, tests what you know, and teaches the rest with references to the book. - demetrius-edelin/tutor
HN text: Reading does not always guarantee retention, especially for complex materials. So I created this tool that extracts the concepts and teaches them back to you (and then tests you as much as you need) so you can really learn the study material.
This is a local project (hosted option coming soon) and it works for EPUB, PDF books or online HTML pages.

Article text:
GitHub - demetrius-edelin/tutor: Self-hosted AI tutor (BYOK) that turns your own EPUB and PDF books into a course. It maps the concepts, tests what you know, and teaches the rest with references to the book. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
demetrius-edelin
tutor Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Self-hosted AI tutor (BYOK) that turns your own EPUB and PDF books into a course. It maps the concepts, tests what you know, and teaches the rest with references to the book.
Readme MIT license Activity Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
45 Commits 45 Commits Folders and files
docs docs src src test test .env.example .env.example .gitignore .gitignore DESIGN-v3.md DESIGN-v3.md LICENSE LICENSE PLAN.md PLAN.md README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.config.ts vite.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
See also my other project, Tapas Habit & Goal Tracker , a full-featured habit tracker for iPhone and Android. Also, I am looking for a full-time role as a software developer. Contact me on LinkedIn .
A personal tutor that teaches you from your books or from online documentation. You make a subject of study, for example "SQL", and add books to it. A book can be an EPUB file, a PDF file, an online book, or online documentation. The tutor finds the concepts in the books and checks which concepts you know. Then it teaches the other concepts one at a time, with references to the books, and tests each one.
The tutor runs on your computer. It uses one large language model (LLM) from Anthropic, OpenAI, or OpenRouter. You select the model.
Why not a simple chat with your books?
A chat with a book answers the questions that you think of. The tutor turns your books into a course:
Structure. The tutor puts the concepts of a subject into modules on a concept map. It teaches them one at a time, from a study queue.
All the concepts. Ingest finds the concepts in each chapter. Then it makes sure that the concepts cover the bold and italic terms of the book.
Visible progress. The concept map and the review board show the status of each concept, for example known, learning, or mastered.
Your control. You select the concepts to test, to learn, and to skip. You also set the order of the study queue. The tutor only suggests.
You need Node.js 22 or later and an API key for Anthropic, OpenAI, or OpenRouter. We test the tutor on macOS.
git clone https://github.com/demetrius-edelin/tutor.git
cd tutor
npm install
cp .env.example .env # then set the provider, the model, and the API key
npm run llm:check # optional: send 3 small test requests to the model
npm run ingest -- SQL /path/to/book.epub # add a book to the subject "SQL"
npm run fetch -- https://doc.rust-lang.org/book/ # optional: download an online book into an EPUB file
npm start # then open http://localhost:3000
In the Windows command prompt, use copy instead of cp .
The ingest command shows the number of model requests and asks before it starts. To start with some chapters of the book only, see Step 3 of the usage guide .
You use the tutor in four steps. Steps 1 to 3 are commands in the terminal. Step 4 is the app in the browser.
flowchart LR
S1["1. Set up<br/>npm install<br/>.env"] --> S2["2. Check (optional)<br/>npm run llm:check<br/>npm run parse"]
S2 --> S3["3. Add a book<br/>or some chapters<br/>npm run ingest"]
S3 --> S4["4. Study<br/>npm start"]
S3 -->|"more chapters or books"| S3
style S2 stroke-dasharray: 5 5
Loading
Set up the tutor . Do this one time.
Check the model and the book . This step is optional. It saves no data.
Add a book to a subject . Add the full book, or only the chapters that you want to study now. Only this step puts books into the tutor. For an online book, download it first .
Study in the browser . Start the app and learn. The app shows the next action at each stage.
The usage guide explains each step and each command.
In the app, you open a subject and go through these stages:
Choose: on the concept map, test, learn, or skip each concept. To do this for many concepts in one step, select their checkboxes and use the bar at the bottom of the page.
Diagnosis: the tutor asks 2 questions about each concept that you selected for a test. To start it, click "Test it" next to a concept, or select concepts and click "Test them".
Study queue: the concepts to learn, in an order that you can change.
Lesson and test: the tutor teaches one concept from your books, with references. Then it tests the concept with the questions that fit it: one question for a simple concept, at most 5 for a larger one. To pass, answer each question correctly. A pass makes the concept mastered.
Review board: a list of all concepts with their status and their stars.
The model writes the questions and grades the open answers. If you think that a grade is wrong, use "Dispute the grade". A disputed answer counts as correct.
Command
Use
Uses the model
What it writes
npm run llm:check
Optional. After a change to .env .
Yes, 3 small requests.
Nothing.
npm run parse -- <book>
Optional. To check a book and to find the chapter numbers for --chapters .
No.
A report in data/parse/<book>/ , for you to read.
npm run fetch -- <url>
To study an online book or online documentation. Then use the EPUB file as the book.
Only for a site with no llms.txt file: one small request.
An EPUB file in data/web/ .
npm run ingest -- <subject> <book> --chapters <list> --preview
Recommended. To check the concepts of some chapters before you save them.
Yes.
Preview files in the folder of the book. The database does not change.
npm run ingest -- <subject> <book> --chapters <list>
To study some chapters of a book.
Yes.
The folder of the book and the database.
npm run ingest -- <subject> <book>
To study all the chapters of a book.
Yes.
The folder of the book and the database.
npm start
Each time that you want to study.
Yes, for the diagnosis, the lessons, and the tests.
Your progress in the database.
npm run refresh -- <subject> <book>
Only after an update of the tutor that changes the parser.
Only for new images.
The section files of the book.
The folder of the book is data/subjects/<subject>/books/<book>/ . The database is data/tutor.db .
The tutor reads its configuration from the .env file. Copy .env.example to .env . The file explains each value.
The tutor has no default model. After a change to .env , run npm run llm:check .
We tested the tutor with different models. GPT-6 Luna from OpenAI, at the reasoning level high , gave the best results. It was also the cheapest model in our tests, and it is very smart for a model in its price range.
Update, October 2026: Claude Haiku 5.5 from Anthropic has now the same price as GPT-6 Luna, but it gives much better results!
LLM_PROVIDER=anthropic
LLM_MODEL=claude-haiku-5-5
LLM_REASONING=high
If you use OpenAI, use these values:
LLM_PROVIDER=openai
LLM_MODEL=gpt-6-luna
LLM_REASONING=high
Cost
The tutor uses your API key, so your provider charges you for each request.
Ingest makes the most requests, because the model reads each chapter and each image of the book. The command shows the number of requests and asks before it starts.
The tutor keeps the model results of each chapter and each image in a cache. A later run does not pay again for the chapters in the cache.
In the app, the diagnosis, the lessons, the questions about a lesson, and the tests use the model.
A high reasoning level makes each request slower.
The app uses port 3000. To use a different port, set PORT in the shell. The tutor does not read PORT from .env .
EPUB files work best. The parser reads the publisher CSS for headings, bold, and italic.
PDF files must be tagged. Word and many publishing tools make tagged PDF files. The parser stops with a clear message for an untagged or a scanned PDF.
Some books show code or tables as images. Ingest uses the model to read them. If the model does not accept images, the images stay as placeholders.
Online books and online documentation: npm run fetch downloads the pages into an EPUB file. The command cannot read pages behind a login, or pages that show their text only after JavaScript runs. See Add an online book .
Must I run parse before ingest ? No. Ingest parses the book by itself. The parse command only prints a report in the terminal, at no cost. See Check a book .
How do I find the number of a chapter for --chapters ? Run npm run parse and read the # column. This number can be different from the number in the title of the chapter.
How do I find the id of a section for --sections ? Run npm run parse and read the file names in data/parse/<book>/sections/ . The file 01-10-pointers.md is section 1.10 .
Must I ingest the full book? No. With --chapters or --sections , you can study a book one part at a time. Later, you can add more parts.
Does ingest always save to the database? Yes. Only --preview keeps the database as it is. With --chapters , the app shows the concepts of these chapters immediately.
Do I pay two times for the chapters of a preview? No. The tutor keeps the model results of each run, so a later run does not pay for these chapters again.
How do I add more chapters later? Run ingest again with a list of the new chapters or sections. See Add more chapters later .
Can I move or delete the book file after ingest? Yes. Ingest keeps a copy of the book file in the folder of the book.
Can the tutor teach from online documentation? Yes. Run npm run fetch with the address of the documentation. Then ingest the EPUB file. See Add an online book .
How do I add a second book to a subject? Run npm run ingest again with the same subject name. The tutor adds the concepts of the new book to the concept map of the subject.
How do I make a new subject? Run npm run ingest with a new subject name. The first book makes the subject.
How do I delete a subject? On the list of subjects, click "Delete" below the subject. Then click the red button in the panel that opens. The tutor deletes the books, the concepts, your progress, and the folder of the subject. You cannot undo this.
Can I delete data/parse/ ? Yes. Nothing else uses it.
The tutor is at version 0.1.0. All the commands and app stages in this README and in the usage guide work. The tutor is for one person on one computer.
Add a book from the app, with the progress of the ingest. Now you add books with a command in the terminal.
A container image, to run the tutor on your own server.
Command
What it does
npm run dev:server and npm run dev:app
Run the server and the app with live reload. Use t

[truncated]

## Original Extract

Self-hosted AI tutor (BYOK) that turns your own EPUB and PDF books into a course. It maps the concepts, tests what you know, and teaches the rest with references to the book. - demetrius-edelin/tutor

Reading does not always guarantee retention, especially for complex materials. So I created this tool that extracts the concepts and teaches them back to you (and then tests you as much as you need) so you can really learn the study material.
This is a local project (hosted option coming soon) and it works for EPUB, PDF books or online HTML pages.

GitHub - demetrius-edelin/tutor: Self-hosted AI tutor (BYOK) that turns your own EPUB and PDF books into a course. It maps the concepts, tests what you know, and teaches the rest with references to the book. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
demetrius-edelin
tutor Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Self-hosted AI tutor (BYOK) that turns your own EPUB and PDF books into a course. It maps the concepts, tests what you know, and teaches the rest with references to the book.
Readme MIT license Activity Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
45 Commits 45 Commits Folders and files
docs docs src src test test .env.example .env.example .gitignore .gitignore DESIGN-v3.md DESIGN-v3.md LICENSE LICENSE PLAN.md PLAN.md README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.config.ts vite.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
See also my other project, Tapas Habit & Goal Tracker , a full-featured habit tracker for iPhone and Android. Also, I am looking for a full-time role as a software developer. Contact me on LinkedIn .
A personal tutor that teaches you from your books or from online documentation. You make a subject of study, for example "SQL", and add books to it. A book can be an EPUB file, a PDF file, an online book, or online documentation. The tutor finds the concepts in the books and checks which concepts you know. Then it teaches the other concepts one at a time, with references to the books, and tests each one.
The tutor runs on your computer. It uses one large language model (LLM) from Anthropic, OpenAI, or OpenRouter. You select the model.
Why not a simple chat with your books?
A chat with a book answers the questions that you think of. The tutor turns your books into a course:
Structure. The tutor puts the concepts of a subject into modules on a concept map. It teaches them one at a time, from a study queue.
All the concepts. Ingest finds the concepts in each chapter. Then it makes sure that the concepts cover the bold and italic terms of the book.
Visible progress. The concept map and the review board show the status of each concept, for example known, learning, or mastered.
Your control. You select the concepts to test, to learn, and to skip. You also set the order of the study queue. The tutor only suggests.
You need Node.js 22 or later and an API key for Anthropic, OpenAI, or OpenRouter. We test the tutor on macOS.
git clone https://github.com/demetrius-edelin/tutor.git
cd tutor
npm install
cp .env.example .env # then set the provider, the model, and the API key
npm run llm:check # optional: send 3 small test requests to the model
npm run ingest -- SQL /path/to/book.epub # add a book to the subject "SQL"
npm run fetch -- https://doc.rust-lang.org/book/ # optional: download an online book into an EPUB file
npm start # then open http://localhost:3000
In the Windows command prompt, use copy instead of cp .
The ingest command shows the number of model requests and asks before it starts. To start with some chapters of the book only, see Step 3 of the usage guide .
You use the tutor in four steps. Steps 1 to 3 are commands in the terminal. Step 4 is the app in the browser.
flowchart LR
S1["1. Set up<br/>npm install<br/>.env"] --> S2["2. Check (optional)<br/>npm run llm:check<br/>npm run parse"]
S2 --> S3["3. Add a book<br/>or some chapters<br/>npm run ingest"]
S3 --> S4["4. Study<br/>npm start"]
S3 -->|"more chapters or books"| S3
style S2 stroke-dasharray: 5 5
Loading
Set up the tutor . Do this one time.
Check the model and the book . This step is optional. It saves no data.
Add a book to a subject . Add the full book, or only the chapters that you want to study now. Only this step puts books into the tutor. For an online book, download it first .
Study in the browser . Start the app and learn. The app shows the next action at each stage.
The usage guide explains each step and each command.
In the app, you open a subject and go through these stages:
Choose: on the concept map, test, learn, or skip each concept. To do this for many concepts in one step, select their checkboxes and use the bar at the bottom of the page.
Diagnosis: the tutor asks 2 questions about each concept that you selected for a test. To start it, click "Test it" next to a concept, or select concepts and click "Test them".
Study queue: the concepts to learn, in an order that you can change.
Lesson and test: the tutor teaches one concept from your books, with references. Then it tests the concept with the questions that fit it: one question for a simple concept, at most 5 for a larger one. To pass, answer each question correctly. A pass makes the concept mastered.
Review board: a list of all concepts with their status and their stars.
The model writes the questions and grades the open answers. If you think that a grade is wrong, use "Dispute the grade". A disputed answer counts as correct.
Command
Use
Uses the model
What it writes
npm run llm:check
Optional. After a change to .env .
Yes, 3 small requests.
Nothing.
npm run parse -- <book>
Optional. To check a book and to find the chapter numbers for --chapters .
No.
A report in data/parse/<book>/ , for you to read.
npm run fetch -- <url>
To study an online book or online documentation. Then use the EPUB file as the book.
Only for a site with no llms.txt file: one small request.
An EPUB file in data/web/ .
npm run ingest -- <subject> <book> --chapters <list> --preview
Recommended. To check the concepts of some chapters before you save them.
Yes.
Preview files in the folder of the book. The database does not change.
npm run ingest -- <subject> <book> --chapters <list>
To study some chapters of a book.
Yes.
The folder of the book and the database.
npm run ingest -- <subject> <book>
To study all the chapters of a book.
Yes.
The folder of the book and the database.
npm start
Each time that you want to study.
Yes, for the diagnosis, the lessons, and the tests.
Your progress in the database.
npm run refresh -- <subject> <book>
Only after an update of the tutor that changes the parser.
Only for new images.
The section files of the book.
The folder of the book is data/subjects/<subject>/books/<book>/ . The database is data/tutor.db .
The tutor reads its configuration from the .env file. Copy .env.example to .env . The file explains each value.
The tutor has no default model. After a change to .env , run npm run llm:check .
We tested the tutor with different models. GPT-6 Luna from OpenAI, at the reasoning level high , gave the best results. It was also the cheapest model in our tests, and it is very smart for a model in its price range.
Update, October 2026: Claude Haiku 5.5 from Anthropic has now the same price as GPT-6 Luna, but it gives much better results!
LLM_PROVIDER=anthropic
LLM_MODEL=claude-haiku-5-5
LLM_REASONING=high
If you use OpenAI, use these values:
LLM_PROVIDER=openai
LLM_MODEL=gpt-6-luna
LLM_REASONING=high
Cost
The tutor uses your API key, so your provider charges you for each request.
Ingest makes the most requests, because the model reads each chapter and each image of the book. The command shows the number of requests and asks before it starts.
The tutor keeps the model results of each chapter and each image in a cache. A later run does not pay again for the chapters in the cache.
In the app, the diagnosis, the lessons, the questions about a lesson, and the tests use the model.
A high reasoning level makes each request slower.
The app uses port 3000. To use a different port, set PORT in the shell. The tutor does not read PORT from .env .
EPUB files work best. The parser reads the publisher CSS for headings, bold, and italic.
PDF files must be tagged. Word and many publishing tools make tagged PDF files. The parser stops with a clear message for an untagged or a scanned PDF.
Some books show code or tables as images. Ingest uses the model to read them. If the model does not accept images, the images stay as placeholders.
Online books and online documentation: npm run fetch downloads the pages into an EPUB file. The command cannot read pages behind a login, or pages that show their text only after JavaScript runs. See Add an online book .
Must I run parse before ingest ? No. Ingest parses the book by itself. The parse command only prints a report in the terminal, at no cost. See Check a book .
How do I find the number of a chapter for --chapters ? Run npm run parse and read the # column. This number can be different from the number in the title of the chapter.
How do I find the id of a section for --sections ? Run npm run parse and read the file names in data/parse/<book>/sections/ . The file 01-10-pointers.md is section 1.10 .
Must I ingest the full book? No. With --chapters or --sections , you can study a book one part at a time. Later, you can add more parts.
Does ingest always save to the database? Yes. Only --preview keeps the database as it is. With --chapters , the app shows the concepts of these chapters immediately.
Do I pay two times for the chapters of a preview? No. The tutor keeps the model results of each run, so a later run does not pay for these chapters again.
How do I add more chapters later? Run ingest again with a list of the new chapters or sections. See Add more chapters later .
Can I move or delete the book file after ingest? Yes. Ingest keeps a copy of the book file in the folder of the book.
Can the tutor teach from online documentation? Yes. Run npm run fetch with the address of the documentation. Then ingest the EPUB file. See Add an online book .
How do I add a second book to a subject? Run npm run ingest again with the same subject name. The tutor adds the concepts of the new book to the concept map of the subject.
How do I make a new subject? Run npm run ingest with a new subject name. The first book makes the subject.
How do I delete a subject? On the list of subjects, click "Delete" below the subject. Then click the red button in the panel that opens. The tutor deletes the books, the concepts, your progress, and the folder of the subject. You cannot undo this.
Can I delete data/parse/ ? Yes. Nothing else uses it.
The tutor is at version 0.1.0. All the commands and app stages in this README and in the usage guide work. The tutor is for one person on one computer.
Add a book from the app, with the progress of the ingest. Now you add books with a command in the terminal.
A container image, to run the tutor on your own server.
Command
What it does
npm run dev:server and npm run dev:app
Run the server and the app with live reload. Use t

[truncated]
