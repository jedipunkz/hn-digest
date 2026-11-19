---
source: "https://github.com/AppliedIngenuity-ai/project-chaperone"
hn_url: "https://news.ycombinator.com/item?id=50567719"
title: "AI Agentic harness with explicit risk analysis"
article_title: "GitHub - AppliedIngenuity-ai/project-chaperone: Project Chaperone - AI Agentic harness that provides explicit risk analysis and extra runtime rule enforcement for enhanced safety. · GitHub"
image: "https://opengraph.githubassets.com/5612787015cfa2e12acac50996376ec190a59307d0af66c2c70cbff6c4d975a3/AppliedIngenuity-ai/project-chaperone"
author: "BraveNewCurency"
captured_at: "2026-11-19T23:25:16Z"
capture_tool: "hn-digest"
hn_id: 50567719
score: 4
comments: 1
posted_at: "2026-11-19T22:55:05Z"
tags:
  - hacker-news
---

# AI Agentic harness with explicit risk analysis

- HN: [50567719](https://news.ycombinator.com/item?id=50567719)
- Source: [github.com](https://github.com/AppliedIngenuity-ai/project-chaperone)
- Score: 4
- Comments: 1
- Posted: 2026-11-19T22:55:05Z

## Translation

Title: AI Agentic harness with explicit risk analysis
Article title: GitHub - AppliedIngenuity-ai/project-chaperone: Project Chaperone - AI Agentic harness that provides explicit risk analysis and extra runtime rule enforcement for enhanced safety. · GitHub
Description: Project Chaperone - AI Agentic harness that provides explicit risk analysis and extra runtime rule enforcement for enhanced safety. - AppliedIngenuity-ai/project-chaperone

Article text:
GitHub - AppliedIngenuity-ai/project-chaperone: Project Chaperone - AI Agentic harness that provides explicit risk analysis and extra runtime rule enforcement for enhanced safety. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
AppliedIngenuity-ai
project-chaperone Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Project Chaperone - AI Agentic harness that provides explicit risk analysis and extra runtime rule enforcement for enhanced safety.
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
chaperone chaperone configs/ llm configs/ llm docs docs examples examples private.example private.example scripts scripts tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Paper: Project Chaperone: A Rule-Enforcing Harness for Proactive Assessment and Prevention of AI Agent Misconduct
Project Chaperone is a harness that allows agentic AI tool calling, while enforcing high-level policies such as
"no personal data may be leave the network". It adds two capabilities:
audit report before running will check every route of all workflows for violations of the rules. A human
makes policy decisions about exceptions to the rules, and those approvals are recorded. They can make these
decisions via the 'shape' of the policies needed, without worrying about the actual workflow details. (I.e. "Is it
reasonable that this workflow wants to write personal details to my calendar?")
Enforcement at runtime happens through data labels ( UNTRUSTED , PERSONAL , ...) and module/tool properties
( OUTPUT , LEAVES_NETWORK , ...). All calls are subject to the rules (minus acknowledged exceptions) at runtime.
A possible violation that is not worth designing away can be accepted as a risk: the call stays forbidden, and is
refused if a run attempts it.
Cartoon: The main idea behind Project Chaperone, simplified to a single rule: PERSONAL data may not leave by e-mail.
Labels stick with the data as it passes through the LLM, and the harness blocks the call that would break the rule.
Project Chaperone fixes the problems of "Agents with tools" (flexible, but a prompt cannot enforce a rule)
and hard-coded workflows (can be analyzed, but the analysis is manual and is often never done.)
In Project Chaperone, that analysis is automated, and the rules are enforced while the system runs.
This repo ships with five shared rules (in examples/rules/default.yaml ):
no destructive actions, nothing leaves the network,
no personal or untrusted data output, and no decision made from untrusted data (such as text an LLM wrote.) Each
project adds a few explicit exceptions. An LLM is optional: a module may call one, and a state may let an LLM choose the
next state in the workflow.
This repo also ships with these examples (some mentioned in the paper). The idea is to demonstrate how a workflow
can ask for small, scoped exceptions to the shared rules:
deliveries - Intelligently add delivery notices from your mail into your calendar.
share_notes - A minimal example showing a problematic workflow.
ask_my_mail - An example of LLM interating with email
safe_rag - Ask questions about the Federalist Papers, or your own documents.
Do the following one-time setup:
git clone https://github.com/AppliedIngenuity-ai/project-chaperone.git
cd project-chaperone
# Set up your virtual environment here (see below)
pip install -e .
# (Optional): also install test libraries
pip install -e ".[test]"
There are many different ways to setup Python virtual environments, this repo is agnostic.
One example using venv ) could be: run
python -m venv .env once, and source .env/bin/activate once per shell in which you will run
the commands below. Note that pyproject.toml asks for 3.13 or newer, and versions before 3.13.5
have not been tested.
python -m pytest
Note that tests 2 are skipped because they would ask a real LLM server (see "Step 1: your LLM" below).
From the folder project-chaperone :
python -m chaperone run examples/deliveries
You will see a WARNING about the mail.send.* module not being properly configured to send mail. This can be
ignored. In fact, the code to send email is not written yet anyway.
The startup audit will run, and ask you about any exceptions (to the 5 default rules )
that the project is requesting (via [examples/deliveries/exceptions.yaml] examples/deliveries/exceptions.yaml .)
Because this is an example, you can type 1 at each prompt. At the end, you are asked for a name (just for you
to remember why you approved it). That named list of approved exceptions are added to
state/deliveries/demo/accepted.json , so identical questions won't be asked again. Note that these approvals are
scoped per example, so "this is OK for package tracking" doesn't accidentally create loopholes for unrelated workflows.
WARNING: In a real system, you would want to examine the details think carefully about the implications before
allowing an exception.
A local web server is started, printing it's address in the form http://127.0.0.1:<port>/?token=<token> .
Open that link in a browser (must be running on the same machine.)
On that web page, you will see a chat area on the left, and a list of workflows the right. Press 'Start'
beside the assistant workflow, and the workflow's output will appear in the chat area. Type anything you want
at each question. In the demo configuration, we have replaced the calls to LLM with a script that requests
pre-defined tool calls: It will search the mail, ask for addresses (you can type any address), then finish.
The workflow then allows you to approve (via checkboxes) any senders it found.
Next, press 'Start' beside update_calendar to kick off this workflow. The conversation shows each made-up
message as it is read, and each entry's values. One message is set aside on purpose: it gives no date.
The README details each step.
Hit Ctrl-C to stop the harness web server.
You can kick off multiple workflows at once, and they are interleaved in the chat output.
Future runs won't re-ask about the exceptions you approved -- until you change the rules, exceptions, workflows,
or module labels/properties: In this case, the startup audit may ask for new approvals.
All workflow state is stored in state/deliveries/demo/ (or more generally state/<project>/<setup>/ ).
You can delete this folder to forget everything the workflow has saved locally. You can also do this via
the web UI: on the State tab, there is a "Start from nothing" button. Note that deleting this folder is
generally only recommended for demo data.
You can pass in a --setup argument to switch between diffrent configurations: The default of --setup demo ,
uses made-up data, and fake tool calls. (i.e. It doesn't talk to a real mail server or LLM). All of it's
state is in the folder state/deliveries/demo , and you can easily start over by deleting that folder.
You can pass in --setup demo_llm to pass the fake emails to a real LLM. (See configuration below)
Alternately, you can use --setup real to try it on your real calendars, email and LLM. (These must be configured,
see "Use your own LLM, mail and calendar" below.) In this configuration you should avoid deleting your entire
state folder, because will create a second calendar. You can use the reset_all workflow instead,
which clears out most state, except for the stateful parts needet to find the calendar.
See docs/exceptions.md to learn about exceptions and acknowledgements.
See the ReadMe of each example for more detail. Also see "Project layout" (below) for how the files are organized.
Delivery notices from your mail, written to a calendar. An LLM chooses the next step of a conversation; addresses,
ticks and approvals are made in boxes the LLM is not passed.
Try it: python -m chaperone run examples/deliveries (See also "Run an example" above)
Setups: demo , demo_llm , real_mail , real (see "Use your own LLM, mail and calendar", below).
Workflows: assistant , a conversation that chooses whose mail is monitored; discover_links , which finds the forms
of those senders' tracking links and asks which to use; update_calendar , which reads their new mail and writes
the calendar; reset_entries , which removes the entries and keeps the senders; reset_all , which removes
everything stored.
Also: prompts.yaml (the LLM's prompts), modules/ (one file per tool), demo/ (the made-up mail, LLM answers and
calendar).
The smallest: a note, an LLM's summary of it, shown to you, then e-mailed. The start reports the e-mail as a possible
violation of rules 3 and 5: answer 4 to each (an acknowledged risk), and the run refuses the e-mail when it is made.
Try it: python -m chaperone run examples/share_notes
Setup: demo (made-up answers, from demo/ ).
Workflow: share_summary : summarizes a note, shows the summary, then tries to send it by e-mail.
A keyword and a question about your mail; the LLM's answer shown with its labels.
Try it: python -m chaperone run examples/ask_my_mail
Setup: demo (made-up mail and answers, from demo/ ).
Workflow: ask_my_mail : searches the mail for the keyword, passes the messages and the question to the LLM, shows
the answer.
Questions about the Federalist Papers, or your own documents. An answer is shown only with quotes found in the
documents.
Try it: python -m chaperone run examples/safe_rag (needs your LLM: "Step 1: your LLM", below)
Setups: local (the papers and your LLM), arxiv (arXiv abstracts), hosted (Claude).
Workflows: single_pass , an answer shown only when a quote is found word for word; multi_pass , with a second
round when none is; answer , the simpler form; typical_rag and just_ask_the_llm , for comparison, with nothing
checked.
Also: corpus/ (the 85 papers), prompts.yaml , questions.yaml and questions_compound.yaml (the measured
questions), results.md (the last measurement).
Every setup but demo needs your own LLM, mail or calendar: "Use your own LLM, mail and calendar", below, says how.
Do not edit the shipped files. Write your changes in private/local.yaml (git ignores private/ ): the keys of that
file are paths of a shipped file, mapped to the keys you need to change, nested as that file nests them.
For example, to use an LLM server at another address, put this in your private/local.yaml :
configs/llm/local.yaml :
api_base : http://localhost:8080/v1
One file holds all your changes. A longer example:
# Your LLM server, and longer answers
configs/llm/local.yaml :
api_base : http://localhost:8080/v1
max_tokens : 8192
# deliveries, setup real: 7 days of mail, your time zone, the calendar's name
examples/deliveries/setups/real.

[truncated]

## Original Extract

Project Chaperone - AI Agentic harness that provides explicit risk analysis and extra runtime rule enforcement for enhanced safety. - AppliedIngenuity-ai/project-chaperone

GitHub - AppliedIngenuity-ai/project-chaperone: Project Chaperone - AI Agentic harness that provides explicit risk analysis and extra runtime rule enforcement for enhanced safety. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
AppliedIngenuity-ai
project-chaperone Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Project Chaperone - AI Agentic harness that provides explicit risk analysis and extra runtime rule enforcement for enhanced safety.
Readme MIT license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
chaperone chaperone configs/ llm configs/ llm docs docs examples examples private.example private.example scripts scripts tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Paper: Project Chaperone: A Rule-Enforcing Harness for Proactive Assessment and Prevention of AI Agent Misconduct
Project Chaperone is a harness that allows agentic AI tool calling, while enforcing high-level policies such as
"no personal data may be leave the network". It adds two capabilities:
audit report before running will check every route of all workflows for violations of the rules. A human
makes policy decisions about exceptions to the rules, and those approvals are recorded. They can make these
decisions via the 'shape' of the policies needed, without worrying about the actual workflow details. (I.e. "Is it
reasonable that this workflow wants to write personal details to my calendar?")
Enforcement at runtime happens through data labels ( UNTRUSTED , PERSONAL , ...) and module/tool properties
( OUTPUT , LEAVES_NETWORK , ...). All calls are subject to the rules (minus acknowledged exceptions) at runtime.
A possible violation that is not worth designing away can be accepted as a risk: the call stays forbidden, and is
refused if a run attempts it.
Cartoon: The main idea behind Project Chaperone, simplified to a single rule: PERSONAL data may not leave by e-mail.
Labels stick with the data as it passes through the LLM, and the harness blocks the call that would break the rule.
Project Chaperone fixes the problems of "Agents with tools" (flexible, but a prompt cannot enforce a rule)
and hard-coded workflows (can be analyzed, but the analysis is manual and is often never done.)
In Project Chaperone, that analysis is automated, and the rules are enforced while the system runs.
This repo ships with five shared rules (in examples/rules/default.yaml ):
no destructive actions, nothing leaves the network,
no personal or untrusted data output, and no decision made from untrusted data (such as text an LLM wrote.) Each
project adds a few explicit exceptions. An LLM is optional: a module may call one, and a state may let an LLM choose the
next state in the workflow.
This repo also ships with these examples (some mentioned in the paper). The idea is to demonstrate how a workflow
can ask for small, scoped exceptions to the shared rules:
deliveries - Intelligently add delivery notices from your mail into your calendar.
share_notes - A minimal example showing a problematic workflow.
ask_my_mail - An example of LLM interating with email
safe_rag - Ask questions about the Federalist Papers, or your own documents.
Do the following one-time setup:
git clone https://github.com/AppliedIngenuity-ai/project-chaperone.git
cd project-chaperone
# Set up your virtual environment here (see below)
pip install -e .
# (Optional): also install test libraries
pip install -e ".[test]"
There are many different ways to setup Python virtual environments, this repo is agnostic.
One example using venv ) could be: run
python -m venv .env once, and source .env/bin/activate once per shell in which you will run
the commands below. Note that pyproject.toml asks for 3.13 or newer, and versions before 3.13.5
have not been tested.
python -m pytest
Note that tests 2 are skipped because they would ask a real LLM server (see "Step 1: your LLM" below).
From the folder project-chaperone :
python -m chaperone run examples/deliveries
You will see a WARNING about the mail.send.* module not being properly configured to send mail. This can be
ignored. In fact, the code to send email is not written yet anyway.
The startup audit will run, and ask you about any exceptions (to the 5 default rules )
that the project is requesting (via [examples/deliveries/exceptions.yaml] examples/deliveries/exceptions.yaml .)
Because this is an example, you can type 1 at each prompt. At the end, you are asked for a name (just for you
to remember why you approved it). That named list of approved exceptions are added to
state/deliveries/demo/accepted.json , so identical questions won't be asked again. Note that these approvals are
scoped per example, so "this is OK for package tracking" doesn't accidentally create loopholes for unrelated workflows.
WARNING: In a real system, you would want to examine the details think carefully about the implications before
allowing an exception.
A local web server is started, printing it's address in the form http://127.0.0.1:<port>/?token=<token> .
Open that link in a browser (must be running on the same machine.)
On that web page, you will see a chat area on the left, and a list of workflows the right. Press 'Start'
beside the assistant workflow, and the workflow's output will appear in the chat area. Type anything you want
at each question. In the demo configuration, we have replaced the calls to LLM with a script that requests
pre-defined tool calls: It will search the mail, ask for addresses (you can type any address), then finish.
The workflow then allows you to approve (via checkboxes) any senders it found.
Next, press 'Start' beside update_calendar to kick off this workflow. The conversation shows each made-up
message as it is read, and each entry's values. One message is set aside on purpose: it gives no date.
The README details each step.
Hit Ctrl-C to stop the harness web server.
You can kick off multiple workflows at once, and they are interleaved in the chat output.
Future runs won't re-ask about the exceptions you approved -- until you change the rules, exceptions, workflows,
or module labels/properties: In this case, the startup audit may ask for new approvals.
All workflow state is stored in state/deliveries/demo/ (or more generally state/<project>/<setup>/ ).
You can delete this folder to forget everything the workflow has saved locally. You can also do this via
the web UI: on the State tab, there is a "Start from nothing" button. Note that deleting this folder is
generally only recommended for demo data.
You can pass in a --setup argument to switch between diffrent configurations: The default of --setup demo ,
uses made-up data, and fake tool calls. (i.e. It doesn't talk to a real mail server or LLM). All of it's
state is in the folder state/deliveries/demo , and you can easily start over by deleting that folder.
You can pass in --setup demo_llm to pass the fake emails to a real LLM. (See configuration below)
Alternately, you can use --setup real to try it on your real calendars, email and LLM. (These must be configured,
see "Use your own LLM, mail and calendar" below.) In this configuration you should avoid deleting your entire
state folder, because will create a second calendar. You can use the reset_all workflow instead,
which clears out most state, except for the stateful parts needet to find the calendar.
See docs/exceptions.md to learn about exceptions and acknowledgements.
See the ReadMe of each example for more detail. Also see "Project layout" (below) for how the files are organized.
Delivery notices from your mail, written to a calendar. An LLM chooses the next step of a conversation; addresses,
ticks and approvals are made in boxes the LLM is not passed.
Try it: python -m chaperone run examples/deliveries (See also "Run an example" above)
Setups: demo , demo_llm , real_mail , real (see "Use your own LLM, mail and calendar", below).
Workflows: assistant , a conversation that chooses whose mail is monitored; discover_links , which finds the forms
of those senders' tracking links and asks which to use; update_calendar , which reads their new mail and writes
the calendar; reset_entries , which removes the entries and keeps the senders; reset_all , which removes
everything stored.
Also: prompts.yaml (the LLM's prompts), modules/ (one file per tool), demo/ (the made-up mail, LLM answers and
calendar).
The smallest: a note, an LLM's summary of it, shown to you, then e-mailed. The start reports the e-mail as a possible
violation of rules 3 and 5: answer 4 to each (an acknowledged risk), and the run refuses the e-mail when it is made.
Try it: python -m chaperone run examples/share_notes
Setup: demo (made-up answers, from demo/ ).
Workflow: share_summary : summarizes a note, shows the summary, then tries to send it by e-mail.
A keyword and a question about your mail; the LLM's answer shown with its labels.
Try it: python -m chaperone run examples/ask_my_mail
Setup: demo (made-up mail and answers, from demo/ ).
Workflow: ask_my_mail : searches the mail for the keyword, passes the messages and the question to the LLM, shows
the answer.
Questions about the Federalist Papers, or your own documents. An answer is shown only with quotes found in the
documents.
Try it: python -m chaperone run examples/safe_rag (needs your LLM: "Step 1: your LLM", below)
Setups: local (the papers and your LLM), arxiv (arXiv abstracts), hosted (Claude).
Workflows: single_pass , an answer shown only when a quote is found word for word; multi_pass , with a second
round when none is; answer , the simpler form; typical_rag and just_ask_the_llm , for comparison, with nothing
checked.
Also: corpus/ (the 85 papers), prompts.yaml , questions.yaml and questions_compound.yaml (the measured
questions), results.md (the last measurement).
Every setup but demo needs your own LLM, mail or calendar: "Use your own LLM, mail and calendar", below, says how.
Do not edit the shipped files. Write your changes in private/local.yaml (git ignores private/ ): the keys of that
file are paths of a shipped file, mapped to the keys you need to change, nested as that file nests them.
For example, to use an LLM server at another address, put this in your private/local.yaml :
configs/llm/local.yaml :
api_base : http://localhost:8080/v1
One file holds all your changes. A longer example:
# Your LLM server, and longer answers
configs/llm/local.yaml :
api_base : http://localhost:8080/v1
max_tokens : 8192
# deliveries, setup real: 7 days of mail, your time zone, the calendar's name
examples/deliveries/setups/real.

[truncated]
