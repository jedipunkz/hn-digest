---
source: "https://wagtail.org/blog/time-to-rethink-shift-left-accessibility-with-agentic-ai/"
hn_url: "https://news.ycombinator.com/item?id=50166800"
title: "Time to rethink Shift Left accessibility with agentic AI"
article_title: "Time to rethink Shift Left accessibility with agentic AI | Wagtail CMS"
image: "https://media.wagtail.org/images/Average_website_accessibiltiy_scores_across.min-1200x630.png"
author: "ThibWeb"
captured_at: "2026-10-20T14:44:12Z"
capture_tool: "hn-digest"
hn_id: 50166800
score: 1
comments: 0
posted_at: "2026-10-20T14:14:27Z"
tags:
  - hacker-news
---

# Time to rethink Shift Left accessibility with agentic AI

- HN: [50166800](https://news.ycombinator.com/item?id=50166800)
- Source: [wagtail.org](https://wagtail.org/blog/time-to-rethink-shift-left-accessibility-with-agentic-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T14:14:27Z

## Translation

Title: Time to rethink Shift Left accessibility with agentic AI
Article title: Time to rethink Shift Left accessibility with agentic AI | Wagtail CMS
Description: Average accessibility scores up 5 points but there is more work to do and it needs a complete rethink

Article text:
Banner Text
Button text
-->
Skip to main content
Project
‹ Back
How we make Wagtail work for everyone
How we reduce Wagtail's climate impact
How we make Wagtail more secure
Explore amazing projects built with Wagtail
Power up your content with Wagtail AI
Why Wagtail is great for public institutions
Headless or traditional Wagtail? You choose!
Why Wagtail is great for mission-driven organizations
Segmented content that just works
Why Wagtail is great for museums, libraries, and more
Publish in multiple languages with Wagtail Localize
Why Wagtail is great for academic institutions
Why your agency should use Wagtail
Why Wagtail is great for businesses
Read the docs. Learn the code.
Wagtail and Django training from the experts
Stay up to date on everything Wagtail
What's New in Wagtail and more
Contribute your time to Wagtail
Help bring new features to life
Find answers to your Wagtail questions
Get expert help with your Wagtail project
The most common Wagtail questions
And support Wagtail at the same time
How we make Wagtail work for everyone
How we reduce Wagtail's climate impact
How we make Wagtail more secure
Explore amazing projects built with Wagtail
Power up your content with Wagtail AI
Why Wagtail is great for public institutions
Headless or traditional Wagtail? You choose!
Why Wagtail is great for mission-driven organizations
Segmented content that just works
Why Wagtail is great for museums, libraries, and more
Publish in multiple languages with Wagtail Localize
Why Wagtail is great for academic institutions
Why your agency should use Wagtail
Why Wagtail is great for businesses
Read the docs. Learn the code.
Wagtail and Django training from the experts
Stay up to date on everything Wagtail
What's New in Wagtail and more
Contribute your time to Wagtail
Help bring new features to life
Find answers to your Wagtail questions
Get expert help with your Wagtail project
The most common Wagtail questions
And support Wagtail at the same time
See if Wagtail is right for you
You can deploy a working Wagtail site in as little as 15 minutes.
Follow our tutorial to build a brand new Wagtail project.
Time to rethink Shift Left accessibility with agentic AI
Moving accessibility from QA to product engineering
Thibaud Colas
Wagtail core team
Previous post
Wagtail API in action: Content ops with DeepSeek Harness
Remember back in May when we shared all our website accessibility indicators were up year-over-year? That’s held up well since. The average accessibility scores of Wagtail websites went up 5 points since May . Here’s a graph of the underlying HTTPArchive data , with Wagtail in indigo:
It’s obviously great to see our accessibility efforts paying off like that, particularly in the more disappointing context that web-wide reports like 2026 WebAIM Million show a 10% increase in detected issues year-over-year (different dataset, the chart above is HTTPArchive data , but still). It doesn’t have to be like that.
From this, my hunch is we can probably take this data-driven approach further: convert issues systematically present on existing sites into product changes to Wagtail, like to our built-in accessibility checker. Do that all the way until there are no issues remaining, going as far upstream as possible (design goals / R&D / product decisions). This has been called the " Shift left " approach to accessibility for a while, and with agentic AI it’s only become more relevant.
It’s hard to express how dramatic of a shift we’re going through for accessibility. Users of assistive tech are using more advanced tools than ever, with LLMs or otherwise. AI agents of all sorts are making much more extensive use of content semantics that we’ve been refining for humans for decades. And for accessibility knowledge workers, unsurprisingly the work is changing dramatically with AI adoption. It’s a challenging shift, as a lot of the craft is driven out of care for people first and foremost, and AI adoption comes with big ethical and societal issues. It’s nonetheless crucial we adapt to that shift if we want to keep making the web better for people.
This calls for a new strategy, but with things changing so fast for now we’ll just focus on adaptation tactics.
Move from audits to product engineering
Getting this one out of the way quick, we need to stop spending so much time on one-off audits and instead redirect auditors’ precious time and expertise towards making upstream changes. Some will be remediation in code or design, some will be even farther upstream (project specs, best practice docs, web standards). WCAG 2.2 AA audits that don’t involve assistive tech can be 90% automated with LLMs today. And every form of testing that only auditors can do now, will be automated within the next year if not earlier. There’s no AI breakthrough needed here, just more time for people to adapt their work to new AI capabilities. Just take a look at computer use work from folks like H Company , or how the Jev decision model plays through the whole Pokémon Red video game . 30 hours worth of UI testing for $2.
The 10% of audit work that remains is all about trust, legal compliance, connecting the audit to the wider business context. Important aspects of the work where technology isn’t as much of a factor.
Switch from QA automation to AI R&D
Having automated QA and monitoring of production systems is now so simple to set up, that the only reason teams aren’t adopting this is sheer inertia or lack of any capacity whatsoever for fundamental QA investments. Everyone should do it, and once they do, QA automation experts’ time can now be spent on AI R&D that benefits from deep automation skills, with tasks like:
Evals of docs to check for gaps in accessibility guidance
Benchmarks of how agents browse a site to identify points of friction
Setting up cross-platform computer use for automation of VoiceOver, NVDA, Dragon, and any other assistive tech.
And also more fundamental data engineering tasks like:
RAG pipelines so accessibility tests are based on the most appropriate guidance possible.
Automating collection of the kind of usage data that leads to better product decisions.
Invest in intellectual community and ethics
Paraphrasing No Man Is an Island , we underrate intellectual community. Whether it’s internal communities of practice, W3C community groups like the ATAG CG , or meetups, we need those spaces and connections to thrive. There’s a dearth of expertise on the ethics of our work specifically, and how to work ethically on the job within competitive pressures. This is crucial to navigate:
The balance of counting on vs. discounting LLMs as agents that help their users achieve tasks.
How to reason about the improvements brought by a technology that has so many problematic aspects.
What to do, broadly-speaking, in a world where LLMs open up so many possibilities.
Our AI guiding principles are super helpful here, but they need practicing day to day, collectively. We underrate how much we can learn from one-another as peers and mentors / mentees, and addressing that is one of the simplest ways to better adapt to the paradigm shift.
Last but certainly not least, talking directly to your users is evergreen. They’re well placed to understand and express which of the fundamentals still matter and warrant the investment, and what’s coming up ahead that is worth looking at sooner than later.
With Wagtail Space 2026 right around the corner, it’s a great time for us to rethink what "Shift-Left" means in our decision making. There is a lot of Shift-Left type thinking on the talks agenda across all things content and code. The separate community day (TBA!) will be an excellent opportunity to relaunch our accessibility team. For accessibility specifically, I expect we’ll want to discuss:
The new PDF converter package .
Agent readiness. We have agent-ready publishing on our roadmap , the overlap with accessibility is very clear as we move from how Wagtail sites rank in search results to completing user tasks.
And for CMS accessibility, it will be very interesting to discuss the pros and cons of developing CMS UI vs. CMS APIs designed to work with agents .
Be sure to sign up to Wagtail Space , and come say hi on the Wagtail Slack workspace if you want to join the accessibility team!
Help build Wagtail enhancements and support the long-term health of the project
1000+ contributors since 2014!
Sign up for events, features and adventures on Wagtail
Wagtail is an open source project by Torchbox . © 2026
We use cookies to measure how you use the website. We want our site to be easy for you to use; understanding how you interact with it helps us know that. Find out more .

## Original Extract

Average accessibility scores up 5 points but there is more work to do and it needs a complete rethink

Banner Text
Button text
-->
Skip to main content
Project
‹ Back
How we make Wagtail work for everyone
How we reduce Wagtail's climate impact
How we make Wagtail more secure
Explore amazing projects built with Wagtail
Power up your content with Wagtail AI
Why Wagtail is great for public institutions
Headless or traditional Wagtail? You choose!
Why Wagtail is great for mission-driven organizations
Segmented content that just works
Why Wagtail is great for museums, libraries, and more
Publish in multiple languages with Wagtail Localize
Why Wagtail is great for academic institutions
Why your agency should use Wagtail
Why Wagtail is great for businesses
Read the docs. Learn the code.
Wagtail and Django training from the experts
Stay up to date on everything Wagtail
What's New in Wagtail and more
Contribute your time to Wagtail
Help bring new features to life
Find answers to your Wagtail questions
Get expert help with your Wagtail project
The most common Wagtail questions
And support Wagtail at the same time
How we make Wagtail work for everyone
How we reduce Wagtail's climate impact
How we make Wagtail more secure
Explore amazing projects built with Wagtail
Power up your content with Wagtail AI
Why Wagtail is great for public institutions
Headless or traditional Wagtail? You choose!
Why Wagtail is great for mission-driven organizations
Segmented content that just works
Why Wagtail is great for museums, libraries, and more
Publish in multiple languages with Wagtail Localize
Why Wagtail is great for academic institutions
Why your agency should use Wagtail
Why Wagtail is great for businesses
Read the docs. Learn the code.
Wagtail and Django training from the experts
Stay up to date on everything Wagtail
What's New in Wagtail and more
Contribute your time to Wagtail
Help bring new features to life
Find answers to your Wagtail questions
Get expert help with your Wagtail project
The most common Wagtail questions
And support Wagtail at the same time
See if Wagtail is right for you
You can deploy a working Wagtail site in as little as 15 minutes.
Follow our tutorial to build a brand new Wagtail project.
Time to rethink Shift Left accessibility with agentic AI
Moving accessibility from QA to product engineering
Thibaud Colas
Wagtail core team
Previous post
Wagtail API in action: Content ops with DeepSeek Harness
Remember back in May when we shared all our website accessibility indicators were up year-over-year? That’s held up well since. The average accessibility scores of Wagtail websites went up 5 points since May . Here’s a graph of the underlying HTTPArchive data , with Wagtail in indigo:
It’s obviously great to see our accessibility efforts paying off like that, particularly in the more disappointing context that web-wide reports like 2026 WebAIM Million show a 10% increase in detected issues year-over-year (different dataset, the chart above is HTTPArchive data , but still). It doesn’t have to be like that.
From this, my hunch is we can probably take this data-driven approach further: convert issues systematically present on existing sites into product changes to Wagtail, like to our built-in accessibility checker. Do that all the way until there are no issues remaining, going as far upstream as possible (design goals / R&D / product decisions). This has been called the " Shift left " approach to accessibility for a while, and with agentic AI it’s only become more relevant.
It’s hard to express how dramatic of a shift we’re going through for accessibility. Users of assistive tech are using more advanced tools than ever, with LLMs or otherwise. AI agents of all sorts are making much more extensive use of content semantics that we’ve been refining for humans for decades. And for accessibility knowledge workers, unsurprisingly the work is changing dramatically with AI adoption. It’s a challenging shift, as a lot of the craft is driven out of care for people first and foremost, and AI adoption comes with big ethical and societal issues. It’s nonetheless crucial we adapt to that shift if we want to keep making the web better for people.
This calls for a new strategy, but with things changing so fast for now we’ll just focus on adaptation tactics.
Move from audits to product engineering
Getting this one out of the way quick, we need to stop spending so much time on one-off audits and instead redirect auditors’ precious time and expertise towards making upstream changes. Some will be remediation in code or design, some will be even farther upstream (project specs, best practice docs, web standards). WCAG 2.2 AA audits that don’t involve assistive tech can be 90% automated with LLMs today. And every form of testing that only auditors can do now, will be automated within the next year if not earlier. There’s no AI breakthrough needed here, just more time for people to adapt their work to new AI capabilities. Just take a look at computer use work from folks like H Company , or how the Jev decision model plays through the whole Pokémon Red video game . 30 hours worth of UI testing for $2.
The 10% of audit work that remains is all about trust, legal compliance, connecting the audit to the wider business context. Important aspects of the work where technology isn’t as much of a factor.
Switch from QA automation to AI R&D
Having automated QA and monitoring of production systems is now so simple to set up, that the only reason teams aren’t adopting this is sheer inertia or lack of any capacity whatsoever for fundamental QA investments. Everyone should do it, and once they do, QA automation experts’ time can now be spent on AI R&D that benefits from deep automation skills, with tasks like:
Evals of docs to check for gaps in accessibility guidance
Benchmarks of how agents browse a site to identify points of friction
Setting up cross-platform computer use for automation of VoiceOver, NVDA, Dragon, and any other assistive tech.
And also more fundamental data engineering tasks like:
RAG pipelines so accessibility tests are based on the most appropriate guidance possible.
Automating collection of the kind of usage data that leads to better product decisions.
Invest in intellectual community and ethics
Paraphrasing No Man Is an Island , we underrate intellectual community. Whether it’s internal communities of practice, W3C community groups like the ATAG CG , or meetups, we need those spaces and connections to thrive. There’s a dearth of expertise on the ethics of our work specifically, and how to work ethically on the job within competitive pressures. This is crucial to navigate:
The balance of counting on vs. discounting LLMs as agents that help their users achieve tasks.
How to reason about the improvements brought by a technology that has so many problematic aspects.
What to do, broadly-speaking, in a world where LLMs open up so many possibilities.
Our AI guiding principles are super helpful here, but they need practicing day to day, collectively. We underrate how much we can learn from one-another as peers and mentors / mentees, and addressing that is one of the simplest ways to better adapt to the paradigm shift.
Last but certainly not least, talking directly to your users is evergreen. They’re well placed to understand and express which of the fundamentals still matter and warrant the investment, and what’s coming up ahead that is worth looking at sooner than later.
With Wagtail Space 2026 right around the corner, it’s a great time for us to rethink what "Shift-Left" means in our decision making. There is a lot of Shift-Left type thinking on the talks agenda across all things content and code. The separate community day (TBA!) will be an excellent opportunity to relaunch our accessibility team. For accessibility specifically, I expect we’ll want to discuss:
The new PDF converter package .
Agent readiness. We have agent-ready publishing on our roadmap , the overlap with accessibility is very clear as we move from how Wagtail sites rank in search results to completing user tasks.
And for CMS accessibility, it will be very interesting to discuss the pros and cons of developing CMS UI vs. CMS APIs designed to work with agents .
Be sure to sign up to Wagtail Space , and come say hi on the Wagtail Slack workspace if you want to join the accessibility team!
Help build Wagtail enhancements and support the long-term health of the project
1000+ contributors since 2014!
Sign up for events, features and adventures on Wagtail
Wagtail is an open source project by Torchbox . © 2026
We use cookies to measure how you use the website. We want our site to be easy for you to use; understanding how you interact with it helps us know that. Find out more .
