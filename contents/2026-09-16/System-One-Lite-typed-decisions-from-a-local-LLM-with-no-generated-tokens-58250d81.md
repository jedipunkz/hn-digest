---
source: "https://github.com/snellingio/system-one"
hn_url: "https://news.ycombinator.com/item?id=49727643"
title: "System One Lite – typed decisions from a local LLM, with no generated tokens"
article_title: "GitHub - snellingio/system-one · GitHub"
image: "https://opengraph.githubassets.com/d22d5f20386a6b266a21408c8ebdf3ffe001b23d829068a98d83b9466542695b/snellingio/system-one"
author: "samsnelling"
captured_at: "2026-09-16T14:39:05Z"
capture_tool: "hn-digest"
hn_id: 49727643
score: 1
comments: 0
posted_at: "2026-09-16T14:31:30Z"
tags:
  - hacker-news
---

# System One Lite – typed decisions from a local LLM, with no generated tokens

- HN: [49727643](https://news.ycombinator.com/item?id=49727643)
- Source: [github.com](https://github.com/snellingio/system-one)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T14:31:30Z

## Translation

Title: System One Lite – typed decisions from a local LLM, with no generated tokens
Article title: GitHub - snellingio/system-one · GitHub
Description: Contribute to snellingio/system-one development by creating an account on GitHub.

Article text:
GitHub - snellingio/system-one · GitHub
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
snellingio
/
system-one
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.github/ workflows .github/ workflows docs docs sdks sdks server server .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
A tiny project that turns a normal local LLM into a typed decision engine. It
needs no fine-tuning, text generation, or parser. Read the docs .
Stop asking language models to write. Start making them decide.
Language models are brilliant at producing text. Software does not want text.
It wants a route, a score, a yes or no, and an honest signal when the answer is
unclear.
So why are we still asking models to write tiny essays? We parse those essays
back into data, validate the data, retry failures, and hope nothing goes off the
rails.
System One Lite deletes the essay.
Send one unstructured state plus up to 64 typed questions. A local model scores
only the answers you allow and returns a probability distribution for every
question.
Unstructured state in. Typed probabilities out. Zero generated tokens.
No free-form response. No JSON repair loop. No invented option that your code
has never heard of.
This is an independent proof of concept.
It uses the System One Models interface.
It is not a new foundation model. It runs a stock open-weight model with MLX.
The project tests simpler AI software where the model can only decide.
A normal LLM feature
System One Lite
Input
Unstructured text
Unstructured text or JSON
Output
A generated string
A typed answer over declared options
Uncertainty
A guess written in prose
The full probability distribution
Validation
Parse, validate, retry
Constrained by construction
Output tokens
One token at a time
Zero
Failure mode
Malformed data or invented values
A valid answer that may still be wrong
Deployment
Usually a hosted model
A fixed local model on Apple silicon
The final row matters. System One Lite does not make a small model infallible.
It makes the limit between the model and your code brutally clear. The
model can choose the wrong declared answer. It cannot create a new one.
That is the difference between asking AI to behave like an API and giving it
an interface it cannot break.
Three primitives. A ridiculous number of decisions.
Type
Ask it to
Get back
Choice
Pick from a closed set
Winner, probabilities, and confidence
Score
Judge a position on an ordered scale
Weighted score, level probabilities, and confidence
Noul
Make a yes or no judgment
Probability of yes
Route support tickets. Rank leads. Gate a workflow. Score risk. Flag content.
Decide whether a human needs to look. Combine several small judgments into a
larger rule that stays in ordinary code.
Each question is independent. One answer cannot leak into the next. Your code,
not a hidden chain of thought, decides what happens after the probabilities
arrive.
System One Lite needs an Apple silicon Mac, Python 3.12 or newer, and
uv . Start the server:
cd server
uv sync
uv run uvicorn system_one_lite.api:app --port 8010
The first start loads mlx-community/Qwen3-1.7B-4bit and compiles the
Metal kernels. Then send a request:
curl -s http://127.0.0.1:8010/evaluate \
-H " Content-Type: application/json " \
-d @- << ' EOF '
{
"state": "My order was due Friday, but it is still in transit.",
"questions": {
"team": {
"type": "choice",
"instructions": "Which team should handle this message?",
"criteria": {
"deliveries": "Late, missing, or damaged orders",
"billing": "Charges, refunds, or payment methods",
"account": "Login, profile, or app problems"
}
},
"needs_reply": {
"type": "noul",
"instructions": "Does the customer need a reply?"
}
}
}
EOF
One request comes back ready for code:
{
"model" : " mlx-community/Qwen3-1.7B-4bit " ,
"answers" : {
"team" : {
"type" : " choice " ,
"choice" : " deliveries " ,
"probabilities" : {
"deliveries" : 0.71 ,
"billing" : 0.18 ,
"account" : 0.11
},
"confidence" : 0.565
},
"needs_reply" : {
"type" : " noul " ,
"noul" : 0.88
}
},
"usage" : {
"input_tokens" : 94 ,
"output_tokens" : 0
}
}
The numbers show the response shape. Exact values depend on the input and
model. The shape does not.
For a longer example with all three question types, open the
quickstart .
The 1.7B model is the default profile. The 4B Instruct model is the
larger profile. Download either model before a run:
uv run python -m tools.download_model default
uv run python -m tools.download_model larger
The demo, benchmark, and eval tools accept either profile. For example:
uv run python -m tools.evals --model larger --limit 20
The server uses default unless SYSTEM_ONE_MODEL selects another profile:
SYSTEM_ONE_MODEL=larger uv run uvicorn system_one_lite.api:app --port 8010
Both model repositories are pinned to exact commits and have checked-in
answer-code registries. Change models only after you run the same accuracy and
option-order checks on both.
The trick is almost offensively simple
For every question, the server:
Writes the state, question, and allowed answers into a prompt.
Assigns each answer a token code such as A , B , or C .
Runs the model up to the answer slot without decoding any text.
Throws away every logit except the valid answer codes.
Applies softmax and maps the probabilities back to your labels.
The model never gets the chance to ramble. It reaches the exact point where
an answer must appear, and System One Lite reads the scores directly.
Choice returns the winning label and the full distribution. Score returns the
probability-weighted level. Noul returns the probability of yes . Every extra
question gets its own full prompt, so questions cannot affect one another.
Read How it works for token alignment, option limits,
and the reason this safe path uses one model pass per question.
Extraordinary claims, meet a local eval
There is no benchmark confetti here. The repository has an eval runner for
JSONL files that use the documented dataset envelope:
cd server
uv run python -m tools.evals --datasets /path/to/jsonl-directory --limit 20
The report shows accuracy by question type. It also rotates Choice options and
checks whether changing their order changes the winner. The public dataset is
the next release step and is not in Git yet. Local files under datasets/ are
ignored, so they cannot be published by accident.
Better yet, add examples from your own traffic. A decision system earns trust
on the states it will actually see, not on a launch graphic.
The repository includes two local SDKs:
Python : sync and async clients with no runtime
dependencies. Python 3.9 or newer.
JavaScript : a typed client for Node 22.18 or
newer.
Both clients use http://127.0.0.1:8010 by default. Set SYSTEM_BASE_URL to
change it.
The part most launch posts bury
System One Lite is an experiment, not a production decision service.
The default 1.7B model is small. It will not match a frontier model on hard
judgments.
The returned probabilities are model scores. They are not calibrated odds
of being correct .
Synthetic eval data does not stand in for real production traffic.
Every question repeats the state and runs separately. Cost grows with the
number and length of questions.
The server handles one inference request at a time and returns 503 while
the engine is busy.
The current server requires Apple silicon because it uses MLX.
Use confidence to route uncertain cases. Set thresholds from labeled data that
matches your traffic. Read Confidence before you let a
score trigger anything expensive, sensitive, or hard to undo.
cd server
uv run ruff check src tools tests
uv run ruff format --check src tools tests
uv run pytest
cd ../sdks/python
uv run --with pytest pytest -q
cd ../javascript
npm ci
npm run typecheck
npm test
Project map
Path
What is inside
server/
FastAPI service, MLX engine, evals, and tests
sdks/python/
Python client
sdks/javascript/
TypeScript client
docs/
Guides and API reference
Start with the introduction . Then read the
API reference for the complete request shape, limits, and error
responses.
MIT . The supported MLX model repositories use Apache 2.0.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to snellingio/system-one development by creating an account on GitHub.

GitHub - snellingio/system-one · GitHub
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
snellingio
/
system-one
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.github/ workflows .github/ workflows docs docs sdks sdks server server .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
A tiny project that turns a normal local LLM into a typed decision engine. It
needs no fine-tuning, text generation, or parser. Read the docs .
Stop asking language models to write. Start making them decide.
Language models are brilliant at producing text. Software does not want text.
It wants a route, a score, a yes or no, and an honest signal when the answer is
unclear.
So why are we still asking models to write tiny essays? We parse those essays
back into data, validate the data, retry failures, and hope nothing goes off the
rails.
System One Lite deletes the essay.
Send one unstructured state plus up to 64 typed questions. A local model scores
only the answers you allow and returns a probability distribution for every
question.
Unstructured state in. Typed probabilities out. Zero generated tokens.
No free-form response. No JSON repair loop. No invented option that your code
has never heard of.
This is an independent proof of concept.
It uses the System One Models interface.
It is not a new foundation model. It runs a stock open-weight model with MLX.
The project tests simpler AI software where the model can only decide.
A normal LLM feature
System One Lite
Input
Unstructured text
Unstructured text or JSON
Output
A generated string
A typed answer over declared options
Uncertainty
A guess written in prose
The full probability distribution
Validation
Parse, validate, retry
Constrained by construction
Output tokens
One token at a time
Zero
Failure mode
Malformed data or invented values
A valid answer that may still be wrong
Deployment
Usually a hosted model
A fixed local model on Apple silicon
The final row matters. System One Lite does not make a small model infallible.
It makes the limit between the model and your code brutally clear. The
model can choose the wrong declared answer. It cannot create a new one.
That is the difference between asking AI to behave like an API and giving it
an interface it cannot break.
Three primitives. A ridiculous number of decisions.
Type
Ask it to
Get back
Choice
Pick from a closed set
Winner, probabilities, and confidence
Score
Judge a position on an ordered scale
Weighted score, level probabilities, and confidence
Noul
Make a yes or no judgment
Probability of yes
Route support tickets. Rank leads. Gate a workflow. Score risk. Flag content.
Decide whether a human needs to look. Combine several small judgments into a
larger rule that stays in ordinary code.
Each question is independent. One answer cannot leak into the next. Your code,
not a hidden chain of thought, decides what happens after the probabilities
arrive.
System One Lite needs an Apple silicon Mac, Python 3.12 or newer, and
uv . Start the server:
cd server
uv sync
uv run uvicorn system_one_lite.api:app --port 8010
The first start loads mlx-community/Qwen3-1.7B-4bit and compiles the
Metal kernels. Then send a request:
curl -s http://127.0.0.1:8010/evaluate \
-H " Content-Type: application/json " \
-d @- << ' EOF '
{
"state": "My order was due Friday, but it is still in transit.",
"questions": {
"team": {
"type": "choice",
"instructions": "Which team should handle this message?",
"criteria": {
"deliveries": "Late, missing, or damaged orders",
"billing": "Charges, refunds, or payment methods",
"account": "Login, profile, or app problems"
}
},
"needs_reply": {
"type": "noul",
"instructions": "Does the customer need a reply?"
}
}
}
EOF
One request comes back ready for code:
{
"model" : " mlx-community/Qwen3-1.7B-4bit " ,
"answers" : {
"team" : {
"type" : " choice " ,
"choice" : " deliveries " ,
"probabilities" : {
"deliveries" : 0.71 ,
"billing" : 0.18 ,
"account" : 0.11
},
"confidence" : 0.565
},
"needs_reply" : {
"type" : " noul " ,
"noul" : 0.88
}
},
"usage" : {
"input_tokens" : 94 ,
"output_tokens" : 0
}
}
The numbers show the response shape. Exact values depend on the input and
model. The shape does not.
For a longer example with all three question types, open the
quickstart .
The 1.7B model is the default profile. The 4B Instruct model is the
larger profile. Download either model before a run:
uv run python -m tools.download_model default
uv run python -m tools.download_model larger
The demo, benchmark, and eval tools accept either profile. For example:
uv run python -m tools.evals --model larger --limit 20
The server uses default unless SYSTEM_ONE_MODEL selects another profile:
SYSTEM_ONE_MODEL=larger uv run uvicorn system_one_lite.api:app --port 8010
Both model repositories are pinned to exact commits and have checked-in
answer-code registries. Change models only after you run the same accuracy and
option-order checks on both.
The trick is almost offensively simple
For every question, the server:
Writes the state, question, and allowed answers into a prompt.
Assigns each answer a token code such as A , B , or C .
Runs the model up to the answer slot without decoding any text.
Throws away every logit except the valid answer codes.
Applies softmax and maps the probabilities back to your labels.
The model never gets the chance to ramble. It reaches the exact point where
an answer must appear, and System One Lite reads the scores directly.
Choice returns the winning label and the full distribution. Score returns the
probability-weighted level. Noul returns the probability of yes . Every extra
question gets its own full prompt, so questions cannot affect one another.
Read How it works for token alignment, option limits,
and the reason this safe path uses one model pass per question.
Extraordinary claims, meet a local eval
There is no benchmark confetti here. The repository has an eval runner for
JSONL files that use the documented dataset envelope:
cd server
uv run python -m tools.evals --datasets /path/to/jsonl-directory --limit 20
The report shows accuracy by question type. It also rotates Choice options and
checks whether changing their order changes the winner. The public dataset is
the next release step and is not in Git yet. Local files under datasets/ are
ignored, so they cannot be published by accident.
Better yet, add examples from your own traffic. A decision system earns trust
on the states it will actually see, not on a launch graphic.
The repository includes two local SDKs:
Python : sync and async clients with no runtime
dependencies. Python 3.9 or newer.
JavaScript : a typed client for Node 22.18 or
newer.
Both clients use http://127.0.0.1:8010 by default. Set SYSTEM_BASE_URL to
change it.
The part most launch posts bury
System One Lite is an experiment, not a production decision service.
The default 1.7B model is small. It will not match a frontier model on hard
judgments.
The returned probabilities are model scores. They are not calibrated odds
of being correct .
Synthetic eval data does not stand in for real production traffic.
Every question repeats the state and runs separately. Cost grows with the
number and length of questions.
The server handles one inference request at a time and returns 503 while
the engine is busy.
The current server requires Apple silicon because it uses MLX.
Use confidence to route uncertain cases. Set thresholds from labeled data that
matches your traffic. Read Confidence before you let a
score trigger anything expensive, sensitive, or hard to undo.
cd server
uv run ruff check src tools tests
uv run ruff format --check src tools tests
uv run pytest
cd ../sdks/python
uv run --with pytest pytest -q
cd ../javascript
npm ci
npm run typecheck
npm test
Project map
Path
What is inside
server/
FastAPI service, MLX engine, evals, and tests
sdks/python/
Python client
sdks/javascript/
TypeScript client
docs/
Guides and API reference
Start with the introduction . Then read the
API reference for the complete request shape, limits, and error
responses.
MIT . The supported MLX model repositories use Apache 2.0.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
