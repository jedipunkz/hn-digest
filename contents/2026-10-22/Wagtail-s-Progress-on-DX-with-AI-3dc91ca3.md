---
source: "https://wagtail.org/blog/our-progress-on-dx-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=50194806"
title: "Wagtail's Progress on DX with AI"
article_title: "Wagtail’s progress on DX with AI | Wagtail CMS"
image: "https://media.wagtail.org/images/model_table_accuracy_and_energy_sneak_peek.min-1200x630.png"
author: "ThibWeb"
captured_at: "2026-10-22T13:31:33Z"
capture_tool: "hn-digest"
hn_id: 50194806
score: 1
comments: 0
posted_at: "2026-10-22T13:07:40Z"
tags:
  - hacker-news
---

# Wagtail's Progress on DX with AI

- HN: [50194806](https://news.ycombinator.com/item?id=50194806)
- Source: [wagtail.org](https://wagtail.org/blog/our-progress-on-dx-with-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T13:07:40Z

## Translation

Title: Wagtail's Progress on DX with AI
Article title: Wagtail’s progress on DX with AI | Wagtail CMS
Description: Where we’re at with AI adoption for developer experience. Opt-in APIs, a prototype CLI, agent skills, and a benchmark of what they cost to run.

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
An overview of where we’re at with AI adoption for Developer Experience
Thibaud Colas
Wagtail core team
Previous post
Time to rethink Shift Left accessibility with agentic AI
With Wagtail Space 2026 right around the corner, we’re in a good place to review our progress on DX with AI - six months after the results of our survey back in May. TL;DR; first, here’s what we’ve shipped :
A preview v3 API to automate content operations.
A prototype CLI with a built-in API client, and other DX niceties.
Agent skills for developers , for expert Wagtail knowledge
Agentic engineering recommendations . Get the agents, sans the vendor lock-in.
If you’re brave enough, you can even trial an experimental MCP server for your CMS .
It’s all opt-in, and designed to provide enough friction in adoption that you get to decide exactly what is right for your context. There are clear synergies here but we also hope all of those capabilities work on their own, even for people who want as little as possible to do with AI.
That’s a lot to go through, it’ll be great to showcase it end-to-end at Wagtail Space. In the meantime, here’s a recorded demo. Developers already know what it’s like to do agentic engineering, so here is the Mistral Large 4 model making content changes via DeepSeek Harness (read why we picked those two ):
It all works pretty well, but there is obviously more work to do to incrementally improve all those new capabilities. And things we haven’t gotten to yet.
Here are the possible improvements we had asked for votes on back in May:
The API and CLI are here. They’re based on Django Ninja creating a machine-readable OpenAPI specification. We have the skills too. On everything else, we are at 10-30% completion.
This is one of the most effective ways to improve developer experience. Docs on Django and Python within Wagtail. Making sure we cover 100% of Wagtail’s features both as reference docs and how-tos, explaining multiple use cases. Good tutorials of course. Doing this across a wide portion of the Wagtail ecosystem, not just core. Making sure we have good inline docs directly in the AST too.
We’re now running the ty type checker on a subset of the Wagtail codebase, to gradually adopt typing. We’re going to run it over more of the code over time. The next big milestone is probably to convert an official Wagtail site’s implementation to be fully typed (our official demo , or Made with Wagtail , or perhaps the User guide or wagtail.org).
There is more to do there. We have created official package guidelines and moved a lot of those opinions directly into our Django package template . This will need more work; probably adding more scaffolding support to the CLI. Or different templates, a lot of people have been asking for a good default headless option.
It’s been really fun to experiment with MCP for the CMS , and it works decently well:
There has been a lot of demand for this so this will require further polish before we’re comfortable recommending it. Over on the docs side, we have:
WebMCP on the Wagtail user guide
A work-in-progress public MCP server for docs , to be released over the next few weeks.
Even though MCP has its issues, it feels like a great opportunity to help CMS users have more agency in how they automate things. Rather than requiring a developer to set up API integrations, they can set up the one MCP endpoint. No more time-consuming copy-pasting from word processors to the CMS!
We also asked for input on the CMS experience back in the survey:
Here, bringing CMS capabilities to agents outside the admin goes a long way. There is more to do! We’re hoping to make progress here with smaller language models, like the Kolibri models from Aleph Alpha in Germany, or Apertus from the Swiss AI initiative.
We had also asked about which considerations on the risks and negatives of AI mattered to our audience, and the results were pretty clear, they all do:
Those are all important factors in deciding what we implement. How to encourage thoughtful and responsible AI adoption for Wagtail projects, and how to make sure Wagtail works well without any of this.
Perhaps where we have made the most progress is on the carbon footprint of AI , as we can now confidently measure the energy it takes for common Wagtail tasks across models:
This is measuring the average energy use per task, across a benchmark of 20 simple Wagtail coding tasks. It’s small per task, but it does add up over time. So what matters most is your usage patterns, and model selection. We have way more work to do on evals like this, which were fundamental to understand which skills to ship and exactly what they should contain.
This now enables us to compare the footprint of live sites serving content, which our yearly carbon footprint reporting focuses on, with common CMS and developer tasks, and other related activities (warning: it’s on a log scale, different orders of magnitude):
There is a lot to unpack here, but the good news is Wagtail Space 2026 is as low-footprint as possible ( we measured ), as it’s all fully online available for free from your own computer 🎉. We have a lot on there that will be interesting to better understand the impacts:
Measuring Web Emissions For a Wagtail Site , where Adam Newman and Oliver Winks will demo their cloud platform with built-in energy measurement .
Beyond Doom: How Frugal AI tackles the real threats to our future from James Martin, author of the Frugal AI training course .
And adoption stories from organizations who pay a close attention to their environmental footprint, like WWF-US and the French government.
And there’ll be lots for DX too, well beyond the above improvements. Onwards!
Help build Wagtail enhancements and support the long-term health of the project
1000+ contributors since 2014!
Sign up for events, features and adventures on Wagtail
Wagtail is an open source project by Torchbox . © 2026
We use cookies to measure how you use the website. We want our site to be easy for you to use; understanding how you interact with it helps us know that. Find out more .

## Original Extract

Where we’re at with AI adoption for developer experience. Opt-in APIs, a prototype CLI, agent skills, and a benchmark of what they cost to run.

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
An overview of where we’re at with AI adoption for Developer Experience
Thibaud Colas
Wagtail core team
Previous post
Time to rethink Shift Left accessibility with agentic AI
With Wagtail Space 2026 right around the corner, we’re in a good place to review our progress on DX with AI - six months after the results of our survey back in May. TL;DR; first, here’s what we’ve shipped :
A preview v3 API to automate content operations.
A prototype CLI with a built-in API client, and other DX niceties.
Agent skills for developers , for expert Wagtail knowledge
Agentic engineering recommendations . Get the agents, sans the vendor lock-in.
If you’re brave enough, you can even trial an experimental MCP server for your CMS .
It’s all opt-in, and designed to provide enough friction in adoption that you get to decide exactly what is right for your context. There are clear synergies here but we also hope all of those capabilities work on their own, even for people who want as little as possible to do with AI.
That’s a lot to go through, it’ll be great to showcase it end-to-end at Wagtail Space. In the meantime, here’s a recorded demo. Developers already know what it’s like to do agentic engineering, so here is the Mistral Large 4 model making content changes via DeepSeek Harness (read why we picked those two ):
It all works pretty well, but there is obviously more work to do to incrementally improve all those new capabilities. And things we haven’t gotten to yet.
Here are the possible improvements we had asked for votes on back in May:
The API and CLI are here. They’re based on Django Ninja creating a machine-readable OpenAPI specification. We have the skills too. On everything else, we are at 10-30% completion.
This is one of the most effective ways to improve developer experience. Docs on Django and Python within Wagtail. Making sure we cover 100% of Wagtail’s features both as reference docs and how-tos, explaining multiple use cases. Good tutorials of course. Doing this across a wide portion of the Wagtail ecosystem, not just core. Making sure we have good inline docs directly in the AST too.
We’re now running the ty type checker on a subset of the Wagtail codebase, to gradually adopt typing. We’re going to run it over more of the code over time. The next big milestone is probably to convert an official Wagtail site’s implementation to be fully typed (our official demo , or Made with Wagtail , or perhaps the User guide or wagtail.org).
There is more to do there. We have created official package guidelines and moved a lot of those opinions directly into our Django package template . This will need more work; probably adding more scaffolding support to the CLI. Or different templates, a lot of people have been asking for a good default headless option.
It’s been really fun to experiment with MCP for the CMS , and it works decently well:
There has been a lot of demand for this so this will require further polish before we’re comfortable recommending it. Over on the docs side, we have:
WebMCP on the Wagtail user guide
A work-in-progress public MCP server for docs , to be released over the next few weeks.
Even though MCP has its issues, it feels like a great opportunity to help CMS users have more agency in how they automate things. Rather than requiring a developer to set up API integrations, they can set up the one MCP endpoint. No more time-consuming copy-pasting from word processors to the CMS!
We also asked for input on the CMS experience back in the survey:
Here, bringing CMS capabilities to agents outside the admin goes a long way. There is more to do! We’re hoping to make progress here with smaller language models, like the Kolibri models from Aleph Alpha in Germany, or Apertus from the Swiss AI initiative.
We had also asked about which considerations on the risks and negatives of AI mattered to our audience, and the results were pretty clear, they all do:
Those are all important factors in deciding what we implement. How to encourage thoughtful and responsible AI adoption for Wagtail projects, and how to make sure Wagtail works well without any of this.
Perhaps where we have made the most progress is on the carbon footprint of AI , as we can now confidently measure the energy it takes for common Wagtail tasks across models:
This is measuring the average energy use per task, across a benchmark of 20 simple Wagtail coding tasks. It’s small per task, but it does add up over time. So what matters most is your usage patterns, and model selection. We have way more work to do on evals like this, which were fundamental to understand which skills to ship and exactly what they should contain.
This now enables us to compare the footprint of live sites serving content, which our yearly carbon footprint reporting focuses on, with common CMS and developer tasks, and other related activities (warning: it’s on a log scale, different orders of magnitude):
There is a lot to unpack here, but the good news is Wagtail Space 2026 is as low-footprint as possible ( we measured ), as it’s all fully online available for free from your own computer 🎉. We have a lot on there that will be interesting to better understand the impacts:
Measuring Web Emissions For a Wagtail Site , where Adam Newman and Oliver Winks will demo their cloud platform with built-in energy measurement .
Beyond Doom: How Frugal AI tackles the real threats to our future from James Martin, author of the Frugal AI training course .
And adoption stories from organizations who pay a close attention to their environmental footprint, like WWF-US and the French government.
And there’ll be lots for DX too, well beyond the above improvements. Onwards!
Help build Wagtail enhancements and support the long-term health of the project
1000+ contributors since 2014!
Sign up for events, features and adventures on Wagtail
Wagtail is an open source project by Torchbox . © 2026
We use cookies to measure how you use the website. We want our site to be easy for you to use; understanding how you interact with it helps us know that. Find out more .
