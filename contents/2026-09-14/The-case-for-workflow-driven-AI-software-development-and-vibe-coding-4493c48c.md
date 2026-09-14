---
source: "https://blog.codeyam.com/p/workflows-for-the-win-for-non-developers"
hn_url: "https://news.ycombinator.com/item?id=49700741"
title: "The case for workflow-driven AI software development and vibe coding"
article_title: "Workflows for the Win (for Non-Developers)"
image: "https://substackcdn.com/image/fetch/$s_!Oo65!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Fed0833c0-f820-4b9c-8ee1-6014f6ca4434_700x600.gif"
author: "nadis"
captured_at: "2026-09-14T18:02:31Z"
capture_tool: "hn-digest"
hn_id: 49700741
score: 1
comments: 0
posted_at: "2026-09-14T17:33:23Z"
tags:
  - hacker-news
---

# The case for workflow-driven AI software development and vibe coding

- HN: [49700741](https://news.ycombinator.com/item?id=49700741)
- Source: [blog.codeyam.com](https://blog.codeyam.com/p/workflows-for-the-win-for-non-developers)
- Score: 1
- Comments: 0
- Posted: 2026-09-14T17:33:23Z

## Translation

Title: The case for workflow-driven AI software development and vibe coding
Article title: Workflows for the Win (for Non-Developers)
Description: The case for workflow-driven AI software development and vibe coding.

Article text:
Workflows for the Win (for Non-Developers)
CodeYam’s Substack
Subscribe Sign in Workflows for the Win (for Non-Developers)
The case for workflow-driven AI software development and vibe coding.
CodeYam and Jared Cosulich Jul 21, 2026 2 1 Share This blog post is written for people who are not familiar with software development. If you are familiar with software development we recommend reading this version: Workflows for the Win (for Developers) .
At CodeYam, we’ve become strong believers in AI workflows. They are a core part of the CodeYam experience. Workflows provide a robust and consistent development experience that gives us confidence that the code being written is not introducing technical debt that will make it harder and harder to make changes to the codebase in the future.
One of the biggest complaints we hear from non-developers who are vibe coding is that eventually, as the application becomes more complex, it also becomes more brittle. It becomes harder to make changes confidently. Eventually the AI is unable to make the requested change as the code is such a mess. Workflows can ensure that the coding, even when vibe coding, is done in a manner that better ensures you will be able to make changes to your application as it grows in complexity.
So what are workflows exactly? How do they differ from other AI tools or techniques such as skills?
We define a workflow as follows:
A workflow is a step-by-step process with automated audits and check-ins at the end of various steps.
This differs from AI skills in that a workflow breaks down a complex challenge into many steps and delivers the instructions for each step only after the previous step is completed.
You could do this with skills by having many skills that you apply as you move through the challenge, but a workflow seeks to automate this process and might have additional checks and audits at the end of each step. This keeps each step in a workflow small and easy to understand.
This is not complex technologically, but deconstructing a complex task into smaller sub-tasks is an often overlooked strategy when dealing with complex challenges. The book The Checklist Manifesto provides ample evidence for the value of taking complex processes, breaking them down, and creating checkpoints for each well-understood stage of the process. Steadily working through each checkpoint ensures that nothing is forgotten and frees up mental capacity for the more challenging, creative parts of the work.
Making a change to a codebase, such as building a feature, is a complex process. You don’t want to just write the code sufficiently to make the change work. You want to write tests, ensure the code is well-organized and easy to find, ensure that nothing else breaks unexpectedly from the change, etc. This is in addition to making sure that the change performs as expected and that you like how it behaves once the change is in place.
A primary goal of software development, beyond building the product, is to ensure that you can make changes to the product quickly and effectively even as the complexity of the product grows. This is no easy feat.
Software explodes in complexity as you build out additional features. Ensuring the code remains well-tested and well-organized as you quickly iterate on the product has always been a challenge for software teams.
AI introduces new benefits and challenges toward achieving this goal. AI can go very fast and allows non-developers to vibe code a product without any understanding of the underlying code. As mentioned before, though, this often results in a brittle application that becomes harder to change over time and frequently must be abandoned.
Workflows can take the AI large language models that exist today and guarantee a much more reliable process for ensuring a codebase remains maintainable regardless of the AIs ability to think through such a challenge itself. And since AIs can work so fast the extra effort to ensure the codebase remains maintainable is not a significant cost.
Unlike humans, AI has no problem following a strict workflow, regardless of its length or how tedious each step is. So in this manner AI + workflows are often more reliable than even a team of senior engineers. It is difficult and tedious for the most talented and diligent software engineer to follow a robust workflow every time. AI does not have this problem.
At CodeYam we have a workflow that, at a high level, looks like this:
The CodeYam workflow for AIs (at a high level). The goal of this workflow is simple. Start by planning the feature. Then rapidly prototype and demo the prototype to ensure that there is a shared understanding of the desired outcome. At this point, it’s easy to iterate on the code to make changes if needed.
Once the prototype is approved we can deconstruct the code to ensure all functionality is organized and test-driven into small, well-defined functions with clear boundaries that are well-tested.
The next step is to document in a glossary that records every function, every function’s usage, and how it is tested. This glossary serves as the organizational step so that every aspect of the codebase is easily discoverable, encourages reusing existing functions, and acts as an audit itself, ensuring that every entry in the glossary is recorded properly, matching the code, with tests to demonstrate its functionality.
We then break these fundamental steps down even further into a workflow that has over 25 steps (and is constantly evolving) with a number of sub-tasks at each step. Many of these steps have audit gates that help ensure that the AI has not forgotten anything and that the system remains clean, well-tested, and well-organized.
This workflow allows us to kick off any feature knowing that the AI will check in at predetermined times to allow the user to confirm the work is being done as expected, with robust steps and audit checks along the way that essentially guarantee the work is done in an organized and reliable manner. This creates a nice rhythm to the work as well as you can easily get a sense as to how far along in the workflow the AI is and when it is likely to complete the work.
Is this completely necessary? Will AIs be able to handle this on their own and reliably ensure that the codebase is always clean, well-tested, well-organized, and maintainable? Possibly. Right now, all of the top models we have tested struggle to do this effectively. Even if they were able to without a strict workflow, though, there still may be advantages to workflows.
Much as the research in The Checklist Manifesto revealed, checklists are valuable tools in any complex work, especially work that is complex and creative.
Why ask any intelligence to keep in memory dozens of to dos when it can externalize that work to a tool? And why not provide that intelligence with tools that help it double check its work or that can automatically handle tedious aspects of the workflow?
We’ve written about this before in To Tool Or Not To Tool , but the idea that an intelligence would avoid using tools that make its life easier because it is intelligent enough to do the work without the tools is ridiculous.
If anything, a sign of higher intelligence is that it uses more tooling, not less. The only situation where it makes sense to avoid a tool is when the task at hand is so easy that even engaging with the tool is a waste of time.
So the better question becomes, why wouldn’t you use workflows and tooling to make the job an AI is trying to accomplish easier and more reliable?
It’s unlikely that the challenge of maintaining a robust, clean, well-tested, and well-organized codebase will become so easy for AIs that tooling is not helpful. Tests themselves are tooling in this same manner. Tests provide a fast and reliable way to ensure that the application works properly in a wide range of situations. So you’d have to believe that AIs would be so powerful as to not need tests at all to completely avoid the benefits of tooling, which seems highly unlikely.
Either way, right now we are seeing great benefits from workflows and the accompanying tooling we have created. They work with any technology stack and at any level of complexity. They give us a lot of confidence that no matter how complex the application we are building gets, it will continue to be easy to make changes to it.
Don’t take our word for it. Try CodeYam yourself. It’s free to use and takes just a few minutes to get started. Simply follow the instructions on the CodeYam homepage .
If you’re building something more ambitious, are new to building software, or want help establishing AI-native workflows for your team, get in touch. We’d be happy to help you get started.
Our mission is to empower individuals and teams to build beautifully designed, highly maintainable software with CodeYam .
Thanks for reading! Subscribe for free to receive new posts and support our work.
2 1 Share Previous Next A guest post by Jared Cosulich Co-founder & CTO at CodeYam Subscribe to Jared Discussion about this post Comments Restacks Top Latest Discussions No posts

## Original Extract

The case for workflow-driven AI software development and vibe coding.

Workflows for the Win (for Non-Developers)
CodeYam’s Substack
Subscribe Sign in Workflows for the Win (for Non-Developers)
The case for workflow-driven AI software development and vibe coding.
CodeYam and Jared Cosulich Jul 21, 2026 2 1 Share This blog post is written for people who are not familiar with software development. If you are familiar with software development we recommend reading this version: Workflows for the Win (for Developers) .
At CodeYam, we’ve become strong believers in AI workflows. They are a core part of the CodeYam experience. Workflows provide a robust and consistent development experience that gives us confidence that the code being written is not introducing technical debt that will make it harder and harder to make changes to the codebase in the future.
One of the biggest complaints we hear from non-developers who are vibe coding is that eventually, as the application becomes more complex, it also becomes more brittle. It becomes harder to make changes confidently. Eventually the AI is unable to make the requested change as the code is such a mess. Workflows can ensure that the coding, even when vibe coding, is done in a manner that better ensures you will be able to make changes to your application as it grows in complexity.
So what are workflows exactly? How do they differ from other AI tools or techniques such as skills?
We define a workflow as follows:
A workflow is a step-by-step process with automated audits and check-ins at the end of various steps.
This differs from AI skills in that a workflow breaks down a complex challenge into many steps and delivers the instructions for each step only after the previous step is completed.
You could do this with skills by having many skills that you apply as you move through the challenge, but a workflow seeks to automate this process and might have additional checks and audits at the end of each step. This keeps each step in a workflow small and easy to understand.
This is not complex technologically, but deconstructing a complex task into smaller sub-tasks is an often overlooked strategy when dealing with complex challenges. The book The Checklist Manifesto provides ample evidence for the value of taking complex processes, breaking them down, and creating checkpoints for each well-understood stage of the process. Steadily working through each checkpoint ensures that nothing is forgotten and frees up mental capacity for the more challenging, creative parts of the work.
Making a change to a codebase, such as building a feature, is a complex process. You don’t want to just write the code sufficiently to make the change work. You want to write tests, ensure the code is well-organized and easy to find, ensure that nothing else breaks unexpectedly from the change, etc. This is in addition to making sure that the change performs as expected and that you like how it behaves once the change is in place.
A primary goal of software development, beyond building the product, is to ensure that you can make changes to the product quickly and effectively even as the complexity of the product grows. This is no easy feat.
Software explodes in complexity as you build out additional features. Ensuring the code remains well-tested and well-organized as you quickly iterate on the product has always been a challenge for software teams.
AI introduces new benefits and challenges toward achieving this goal. AI can go very fast and allows non-developers to vibe code a product without any understanding of the underlying code. As mentioned before, though, this often results in a brittle application that becomes harder to change over time and frequently must be abandoned.
Workflows can take the AI large language models that exist today and guarantee a much more reliable process for ensuring a codebase remains maintainable regardless of the AIs ability to think through such a challenge itself. And since AIs can work so fast the extra effort to ensure the codebase remains maintainable is not a significant cost.
Unlike humans, AI has no problem following a strict workflow, regardless of its length or how tedious each step is. So in this manner AI + workflows are often more reliable than even a team of senior engineers. It is difficult and tedious for the most talented and diligent software engineer to follow a robust workflow every time. AI does not have this problem.
At CodeYam we have a workflow that, at a high level, looks like this:
The CodeYam workflow for AIs (at a high level). The goal of this workflow is simple. Start by planning the feature. Then rapidly prototype and demo the prototype to ensure that there is a shared understanding of the desired outcome. At this point, it’s easy to iterate on the code to make changes if needed.
Once the prototype is approved we can deconstruct the code to ensure all functionality is organized and test-driven into small, well-defined functions with clear boundaries that are well-tested.
The next step is to document in a glossary that records every function, every function’s usage, and how it is tested. This glossary serves as the organizational step so that every aspect of the codebase is easily discoverable, encourages reusing existing functions, and acts as an audit itself, ensuring that every entry in the glossary is recorded properly, matching the code, with tests to demonstrate its functionality.
We then break these fundamental steps down even further into a workflow that has over 25 steps (and is constantly evolving) with a number of sub-tasks at each step. Many of these steps have audit gates that help ensure that the AI has not forgotten anything and that the system remains clean, well-tested, and well-organized.
This workflow allows us to kick off any feature knowing that the AI will check in at predetermined times to allow the user to confirm the work is being done as expected, with robust steps and audit checks along the way that essentially guarantee the work is done in an organized and reliable manner. This creates a nice rhythm to the work as well as you can easily get a sense as to how far along in the workflow the AI is and when it is likely to complete the work.
Is this completely necessary? Will AIs be able to handle this on their own and reliably ensure that the codebase is always clean, well-tested, well-organized, and maintainable? Possibly. Right now, all of the top models we have tested struggle to do this effectively. Even if they were able to without a strict workflow, though, there still may be advantages to workflows.
Much as the research in The Checklist Manifesto revealed, checklists are valuable tools in any complex work, especially work that is complex and creative.
Why ask any intelligence to keep in memory dozens of to dos when it can externalize that work to a tool? And why not provide that intelligence with tools that help it double check its work or that can automatically handle tedious aspects of the workflow?
We’ve written about this before in To Tool Or Not To Tool , but the idea that an intelligence would avoid using tools that make its life easier because it is intelligent enough to do the work without the tools is ridiculous.
If anything, a sign of higher intelligence is that it uses more tooling, not less. The only situation where it makes sense to avoid a tool is when the task at hand is so easy that even engaging with the tool is a waste of time.
So the better question becomes, why wouldn’t you use workflows and tooling to make the job an AI is trying to accomplish easier and more reliable?
It’s unlikely that the challenge of maintaining a robust, clean, well-tested, and well-organized codebase will become so easy for AIs that tooling is not helpful. Tests themselves are tooling in this same manner. Tests provide a fast and reliable way to ensure that the application works properly in a wide range of situations. So you’d have to believe that AIs would be so powerful as to not need tests at all to completely avoid the benefits of tooling, which seems highly unlikely.
Either way, right now we are seeing great benefits from workflows and the accompanying tooling we have created. They work with any technology stack and at any level of complexity. They give us a lot of confidence that no matter how complex the application we are building gets, it will continue to be easy to make changes to it.
Don’t take our word for it. Try CodeYam yourself. It’s free to use and takes just a few minutes to get started. Simply follow the instructions on the CodeYam homepage .
If you’re building something more ambitious, are new to building software, or want help establishing AI-native workflows for your team, get in touch. We’d be happy to help you get started.
Our mission is to empower individuals and teams to build beautifully designed, highly maintainable software with CodeYam .
Thanks for reading! Subscribe for free to receive new posts and support our work.
2 1 Share Previous Next A guest post by Jared Cosulich Co-founder & CTO at CodeYam Subscribe to Jared Discussion about this post Comments Restacks Top Latest Discussions No posts
