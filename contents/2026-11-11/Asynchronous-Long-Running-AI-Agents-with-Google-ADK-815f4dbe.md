---
source: "https://benjaminetienne.com/blog/async-ai-agents-in-google-adk-pause-resume-human-in-the-loop/"
hn_url: "https://news.ycombinator.com/item?id=50455012"
title: "Asynchronous, Long-Running AI Agents with Google ADK"
article_title: "Fire and Forget: Building Asynchronous, Long-Running AI Agents with Google ADK — Benjamin Etienne"
image: "https://benjaminetienne.com/images/og-base_hu_4f6aa3dac0d4170f.png"
author: "betienne"
captured_at: "2026-11-11T16:56:16Z"
capture_tool: "hn-digest"
hn_id: 50455012
score: 1
comments: 0
posted_at: "2026-11-11T16:06:28Z"
tags:
  - hacker-news
---

# Asynchronous, Long-Running AI Agents with Google ADK

- HN: [50455012](https://news.ycombinator.com/item?id=50455012)
- Source: [benjaminetienne.com](https://benjaminetienne.com/blog/async-ai-agents-in-google-adk-pause-resume-human-in-the-loop/)
- Score: 1
- Comments: 0
- Posted: 2026-11-11T16:06:28Z

## Translation

Title: Asynchronous, Long-Running AI Agents with Google ADK
Article title: Fire and Forget: Building Asynchronous, Long-Running AI Agents with Google ADK — Benjamin Etienne
Description: Learn how to build long-running, asynchronous AI agents in Google ADK using FastAPI webhooks, DatabaseSessionService, and Human-in-the-Loop (HITL) pause/resume.

Article text:
Benjamin Etienne
Writing
About
Contact
Fire and Forget: Building Asynchronous, Long-Running AI Agents with Google ADK
Why we should stop chatting with agents
Why synchronous HTTP breaks for long-running AI agents
Why is this a problem ?
The Fire-and-Forget Pattern: HTTP 202 Accepted and Webhooks
Human-in-the-Loop (HITL) in Google ADK: Pausing and Resuming Workflows
Human validation and Long-running Operations
Why the loop can be plain Python
Deep diving in what exactly happens behind the scenes
Pausing and resuming agents, human-in-the-loop, and long-running operations in Google ADK
Why we should stop chatting with agents
In most of agent use cases today there is a UI in which the user and the agent have a conversation. This works for certain use cases where questions and answers are quick, but I believe most agents simply don’t follow a “conversational” pattern.
Because let’s be honest: real agents today have a long lifetime. They can span minutes, hours, maybe days. And obviously, you don’t wait for hours in front of your chat window.
Communication with humans can follow two patterns: synchronous and asynchronous.
Synchronous = phone call . If one drops, the conversation ends. You cannot have more than 1 call at the same time.
Asynchronous = emails . You write, and you forget. You don’t just stare at your mailbox (well, you may when you wait for an important message, but you’ll quickly get bored). You go do other things, and a notification pings you when a reply arrives.
Just like communication with humans follows those 2 principles, I believe agents should be treated the same. Not like humans, but like synchronous or asynchronous systems.
the first one is, as i said above, user experience. It is exhausting having to monitor something constantly to reply and being idle the rest of the time
the second one is related to costs and infrastructure. An agent that keeps a channel open with the user (via websockets for example) is consuming bandwidth, and stays alive while waiting for the user’s response. This idle time on the agent side is terrible in terms of costs
Let’s put it simply: Agents spend most of their time waiting.
They cost RAM, they cost CPU, they cost bandwidth.
Most of real agentic use cases today require either a human validation step, or a reaction to signals (webhook triggers for example). A human validation step may take several hours, what happens to the agent during this time ? What happens if the connection breaks ?
So when should we use async agents ?
The work takes minutes to days -> Async
The answer is needed right now -> Chat
A human has to approve or give input along the way, but not immediately -> Async
The human has to approve or give input along the way immediately -> Chat
Should start with a trigger (schedule, event) -> async
In the upcoming section, i want to guide you through a concrete example of how to implement async concepts with the Agent Development Kit from Google. Most implementations I have seen from tutorials or blog posts are often very simplistic, and i wanted to go through a realistic implementation in a series of articles.
Part 1 will cover concepts specific to ADK - Runners, DatabaseSessionService, Events, FunctionCalls, RequestInput- that are used to run agents, persist memory, resume executions, etc. I will also explain how webhooks works, and how to use them to communicate with distant services. This part will be mostly local development, with servers running on localhost and SQLite databases.
Part 2 will cover the move to production of what we build. Specifically, we will cover services like Cloud Run to host your apps and that can scale to zero when not used (efficient for costs), Cloud Tasks to delegate tasks in the background with retries, and Cloud Scheduler to schedule tasks.
Let’s kickstart this first part !
Why synchronous HTTP breaks for long-running AI agents
Let’s imagine we have an agent that does a very simple task: fetching rss feeds and summarizing the latest ones to provide a daily brief.
It’s a good test subject: it does real work (API calls + LLM summary) and takes long enough that waiting for it could hurt. I currently have around 20 feeds, and it takes around 1–2 min to read all these. Ok, it’s not overwhelming, but think of scaling to 1000 sources and you get the idea.
The agent is simply a LLM and a tool. The tool itself is just a loop over sources:
import feedparser
FEEDS = [
"https://blog.google/technology/ai/rss/" ,
"https://blog.google/innovation-and-ai/infrastructure-and-cloud/google-cloud/rss/" ,
"https://cloudblog.withgoogle.com/products/ai-machine-learning/rss/" ,
"https://developers.googleblog.com/feeds/posts/default/" ,
"https://simonwillison.net/atom/everything/" ,
"https://hnrss.org/newest?q=htmx+OR+ADK+OR+ %22a gent+development%22&points=20" ,
"https://huggingface.co/blog/feed.xml" ,
"https://www.anthropic.com/rss.xml" ,
"https://read.deeplearning.ai/the-batch/rss/" ,
"https://thegradient.pub/rss/" ,
"http://export.arxiv.org/rss/cs.AI" ,
"http://export.arxiv.org/rss/cs.CL" ,
]
def fetch_rss_feeds () -> str :
"""Fetch the latest articles from the configured RSS feeds. Call this first."""
results = []
for url in FEEDS :
entries = feedparser . parse ( url )[ 'entries' ]
for entry in entries :
if entry :
title = entry . title or "(no title)"
link = entry . link or ""
published = entry . published if hasattr ( entry , "published" ) else "unknown"
summary = ( entry . summary or "" )[: 300 ] . strip ()
results . append ( f "[ { published } ] { title } \n { link } \n { summary } \n " )
return " \n --- \n " . join ( results ) if results else "No articles found."
This is a very naive implementation. In a production context, i would mostly use TaskGroups and Semaphores from the asyncio library. For now, let’s consider this as our tool.
The agent is a simple LLM with this tool:
from google.adk.agents import Agent
root_agent = Agent (
name = "daily-feed" ,
model = "gemini-4-argon" ,
instruction = """
You are a sharp tech analyst writing a daily AI & data engineering briefing.
Your first action MUST be to call the fetch_rss_feeds tool to get today's articles.
Then summarize ONLY the articles returned by that tool.
Do NOT search the web. Do NOT add information from outside the fetched articles.
For each entry, provide:
- The title and source
- A 1-2 sentence summary of why it matters
- The article link
Rules:
- Be concise and opinionated. Skip fluff.
- Prioritize novelty and impact.
- Always cite your sources with the link.
""" ,
tools = [ fetch_rss_feeds ],
)
So far, nothing hard.
Before going to the execution loop of the agent, I want to remind some concepts / objects used by ADK since they will be key in the rest of the article.
Now we have the building blocks, let’s wire them together. Proceed in steps:
Define the way the session will be stored (in memory or in a db)
Run the runner and collect the events
app_name =...
agent =...
session_service =...
async def _run ( message : str , session_id : str , ** kwargs ):
# Create a session attached to the service
await session_service . create_session (
app_name = app_name ,
user_id = "admin" ,
session_id = session_id
)
# Create a runner
runner = Runner (
agent = agent , app_name = app_name , session_service = session_service
)
# Format the message as a Content
content = types . Content ( role = "user" , parts = [ types . Part ( text = message )])
# Run the loop and collect events
async for event in runner . run_async (
user_id = "admin" , session_id = session_id , new_message = content
):
if event . content and event . content . parts :
print ( f "agent> { event . content . parts [ 0 ] . text } \n " )
return event . content . parts [ 0 ] . text
Let’s imagine our agent is deployed on a remote endpoint (for ex, as a microservice). It provides a /run endpoint which triggers the execution of the agent. The request body looks like this:
import uuid
from pydantic import BaseModel , Field
class RunRequest ( BaseModel ):
session_id : str = Field ( default_factory = lambda : str ( uuid . uuid4 ())[: 8 ])
message : str = "Start"
interrupt_id : str | None = None
A very minimal implementation of our api could be implemented like this:
# sync_agent.py
def to_sync_app (
agent ,
app_name : str ,
session_service : InMemorySessionService | DatabaseSessionService
):
async def _run ( message : str , session_id : str , ** kwargs ):
# defined above
api = FastAPI ( title = app_name )
@api.get ( "/health" )
def health ():
return { "agent" : app_name , "status" : "healthy" }
@api.post ( "/run" )
async def run ( req : RunRequest ):
t0 = time . time ()
output = await asyncio . create_task ( _run ( ** req . model_dump ()))
return { "status" : "success" , "output" : output [: 100 ], "time_elapsed" : time . time () - t0 }
return api
In this implementation, despite the async in front of the run function, we wait for the output before returning a message.
This is 100% synchronous: the HTTP request stays open until the agent responds.
Wrap the agent in the to_sync_app and start the agent with:
uv run uvicorn rss-feed-parser.sync_agent:app --port 8002
Let’s open another terminal and send a request with curl
curl -X POST http://127.0.0.1:8002/run \
-H "Content-Type: application/json" \
-d '{"message": "start"}'
You can see the agent receiving the request and starting working in the agent terminal:
INFO: Uvicorn running on http://127.0.0.1:8002 ( Press CTRL+C to quit )
agent> I 'll fetch today' s articles right away!
...
Since we are in sync mode, the curl terminal just waits and hangs. The HTTP connection is kept open, and we eventually receive after 1–2 min the reply:
{
"status" : "success" ,
"output" : "Here's your **AI & Data Engineering Briefing - November 6, 2026** ☕\n\n---\n\n## 🔥 Top Stories\n\n### 1. *" (truncated) ,
"time_elapsed" : 83.74279189109802
}
Our application has just kept an HTTP connection open for 83s, between our request and the final report. During that time, the terminal was blocked.
Timeouts because on most servers, HTTP request will timeout after some time. Could be seconds, or hours, but it will.
Resources because during all this time, the connection was kept open. Keeping a connection open means keeping the server on the client side up, which is stupid because nothing happens. Furthermore, the client cannot perform any work and just has to wait. Not efficient.
The Fire-and-Forget Pattern: HTTP 202 Accepted and Webhooks
As said before, thinking async is thinking “delegation”. You hand over a task to the agent, go do something else, and get notified when the agent is done.
Think about when you clean your clothes. There are several options:
You go to the self-service laundry, and just stay for the whole time the machine is running, in case someone steals your clothes. You stay and you stare at the machine.
Or you can go to a professional cleaner, drop your clothes, get a receipt, and go home. The laundry calls you when it’s ready.
In the latter, the receipt is the task id, and the phone call is a webhook. And what this means concretely for you is that you can live your life between the moment you drop your clothes and the moment you collect them.
Let’s now write a to_async_app wrapper that returns 202 when receiving a new request, with the task id attached (the receipt). We can reuse the previous to_sync_app and only change the /run route:
@api.post ( "/run" , status_code = 202 )
async def run ( req : RunRequest ):
asyncio . create_task ( _run ( ** req . model_dump ()))
return { "status" : "created" , "session_id" : req . session_id }
Two things changed compared to the sync route:
We no longer await the task. asyncio.create_task schedules _run to run in the background and returns immediately, so the endpoint can answer right away.
The task id is simply the session_id : the one the caller sent, or the one RunRequest generated.
If you curl the agent now, it immediately replies with
{"status":"created","session_id":"092e44f1"}
But how does the agent communicate, or “ca

[truncated]

## Original Extract

Learn how to build long-running, asynchronous AI agents in Google ADK using FastAPI webhooks, DatabaseSessionService, and Human-in-the-Loop (HITL) pause/resume.

Benjamin Etienne
Writing
About
Contact
Fire and Forget: Building Asynchronous, Long-Running AI Agents with Google ADK
Why we should stop chatting with agents
Why synchronous HTTP breaks for long-running AI agents
Why is this a problem ?
The Fire-and-Forget Pattern: HTTP 202 Accepted and Webhooks
Human-in-the-Loop (HITL) in Google ADK: Pausing and Resuming Workflows
Human validation and Long-running Operations
Why the loop can be plain Python
Deep diving in what exactly happens behind the scenes
Pausing and resuming agents, human-in-the-loop, and long-running operations in Google ADK
Why we should stop chatting with agents
In most of agent use cases today there is a UI in which the user and the agent have a conversation. This works for certain use cases where questions and answers are quick, but I believe most agents simply don’t follow a “conversational” pattern.
Because let’s be honest: real agents today have a long lifetime. They can span minutes, hours, maybe days. And obviously, you don’t wait for hours in front of your chat window.
Communication with humans can follow two patterns: synchronous and asynchronous.
Synchronous = phone call . If one drops, the conversation ends. You cannot have more than 1 call at the same time.
Asynchronous = emails . You write, and you forget. You don’t just stare at your mailbox (well, you may when you wait for an important message, but you’ll quickly get bored). You go do other things, and a notification pings you when a reply arrives.
Just like communication with humans follows those 2 principles, I believe agents should be treated the same. Not like humans, but like synchronous or asynchronous systems.
the first one is, as i said above, user experience. It is exhausting having to monitor something constantly to reply and being idle the rest of the time
the second one is related to costs and infrastructure. An agent that keeps a channel open with the user (via websockets for example) is consuming bandwidth, and stays alive while waiting for the user’s response. This idle time on the agent side is terrible in terms of costs
Let’s put it simply: Agents spend most of their time waiting.
They cost RAM, they cost CPU, they cost bandwidth.
Most of real agentic use cases today require either a human validation step, or a reaction to signals (webhook triggers for example). A human validation step may take several hours, what happens to the agent during this time ? What happens if the connection breaks ?
So when should we use async agents ?
The work takes minutes to days -> Async
The answer is needed right now -> Chat
A human has to approve or give input along the way, but not immediately -> Async
The human has to approve or give input along the way immediately -> Chat
Should start with a trigger (schedule, event) -> async
In the upcoming section, i want to guide you through a concrete example of how to implement async concepts with the Agent Development Kit from Google. Most implementations I have seen from tutorials or blog posts are often very simplistic, and i wanted to go through a realistic implementation in a series of articles.
Part 1 will cover concepts specific to ADK - Runners, DatabaseSessionService, Events, FunctionCalls, RequestInput- that are used to run agents, persist memory, resume executions, etc. I will also explain how webhooks works, and how to use them to communicate with distant services. This part will be mostly local development, with servers running on localhost and SQLite databases.
Part 2 will cover the move to production of what we build. Specifically, we will cover services like Cloud Run to host your apps and that can scale to zero when not used (efficient for costs), Cloud Tasks to delegate tasks in the background with retries, and Cloud Scheduler to schedule tasks.
Let’s kickstart this first part !
Why synchronous HTTP breaks for long-running AI agents
Let’s imagine we have an agent that does a very simple task: fetching rss feeds and summarizing the latest ones to provide a daily brief.
It’s a good test subject: it does real work (API calls + LLM summary) and takes long enough that waiting for it could hurt. I currently have around 20 feeds, and it takes around 1–2 min to read all these. Ok, it’s not overwhelming, but think of scaling to 1000 sources and you get the idea.
The agent is simply a LLM and a tool. The tool itself is just a loop over sources:
import feedparser
FEEDS = [
"https://blog.google/technology/ai/rss/" ,
"https://blog.google/innovation-and-ai/infrastructure-and-cloud/google-cloud/rss/" ,
"https://cloudblog.withgoogle.com/products/ai-machine-learning/rss/" ,
"https://developers.googleblog.com/feeds/posts/default/" ,
"https://simonwillison.net/atom/everything/" ,
"https://hnrss.org/newest?q=htmx+OR+ADK+OR+ %22a gent+development%22&points=20" ,
"https://huggingface.co/blog/feed.xml" ,
"https://www.anthropic.com/rss.xml" ,
"https://read.deeplearning.ai/the-batch/rss/" ,
"https://thegradient.pub/rss/" ,
"http://export.arxiv.org/rss/cs.AI" ,
"http://export.arxiv.org/rss/cs.CL" ,
]
def fetch_rss_feeds () -> str :
"""Fetch the latest articles from the configured RSS feeds. Call this first."""
results = []
for url in FEEDS :
entries = feedparser . parse ( url )[ 'entries' ]
for entry in entries :
if entry :
title = entry . title or "(no title)"
link = entry . link or ""
published = entry . published if hasattr ( entry , "published" ) else "unknown"
summary = ( entry . summary or "" )[: 300 ] . strip ()
results . append ( f "[ { published } ] { title } \n { link } \n { summary } \n " )
return " \n --- \n " . join ( results ) if results else "No articles found."
This is a very naive implementation. In a production context, i would mostly use TaskGroups and Semaphores from the asyncio library. For now, let’s consider this as our tool.
The agent is a simple LLM with this tool:
from google.adk.agents import Agent
root_agent = Agent (
name = "daily-feed" ,
model = "gemini-4-argon" ,
instruction = """
You are a sharp tech analyst writing a daily AI & data engineering briefing.
Your first action MUST be to call the fetch_rss_feeds tool to get today's articles.
Then summarize ONLY the articles returned by that tool.
Do NOT search the web. Do NOT add information from outside the fetched articles.
For each entry, provide:
- The title and source
- A 1-2 sentence summary of why it matters
- The article link
Rules:
- Be concise and opinionated. Skip fluff.
- Prioritize novelty and impact.
- Always cite your sources with the link.
""" ,
tools = [ fetch_rss_feeds ],
)
So far, nothing hard.
Before going to the execution loop of the agent, I want to remind some concepts / objects used by ADK since they will be key in the rest of the article.
Now we have the building blocks, let’s wire them together. Proceed in steps:
Define the way the session will be stored (in memory or in a db)
Run the runner and collect the events
app_name =...
agent =...
session_service =...
async def _run ( message : str , session_id : str , ** kwargs ):
# Create a session attached to the service
await session_service . create_session (
app_name = app_name ,
user_id = "admin" ,
session_id = session_id
)
# Create a runner
runner = Runner (
agent = agent , app_name = app_name , session_service = session_service
)
# Format the message as a Content
content = types . Content ( role = "user" , parts = [ types . Part ( text = message )])
# Run the loop and collect events
async for event in runner . run_async (
user_id = "admin" , session_id = session_id , new_message = content
):
if event . content and event . content . parts :
print ( f "agent> { event . content . parts [ 0 ] . text } \n " )
return event . content . parts [ 0 ] . text
Let’s imagine our agent is deployed on a remote endpoint (for ex, as a microservice). It provides a /run endpoint which triggers the execution of the agent. The request body looks like this:
import uuid
from pydantic import BaseModel , Field
class RunRequest ( BaseModel ):
session_id : str = Field ( default_factory = lambda : str ( uuid . uuid4 ())[: 8 ])
message : str = "Start"
interrupt_id : str | None = None
A very minimal implementation of our api could be implemented like this:
# sync_agent.py
def to_sync_app (
agent ,
app_name : str ,
session_service : InMemorySessionService | DatabaseSessionService
):
async def _run ( message : str , session_id : str , ** kwargs ):
# defined above
api = FastAPI ( title = app_name )
@api.get ( "/health" )
def health ():
return { "agent" : app_name , "status" : "healthy" }
@api.post ( "/run" )
async def run ( req : RunRequest ):
t0 = time . time ()
output = await asyncio . create_task ( _run ( ** req . model_dump ()))
return { "status" : "success" , "output" : output [: 100 ], "time_elapsed" : time . time () - t0 }
return api
In this implementation, despite the async in front of the run function, we wait for the output before returning a message.
This is 100% synchronous: the HTTP request stays open until the agent responds.
Wrap the agent in the to_sync_app and start the agent with:
uv run uvicorn rss-feed-parser.sync_agent:app --port 8002
Let’s open another terminal and send a request with curl
curl -X POST http://127.0.0.1:8002/run \
-H "Content-Type: application/json" \
-d '{"message": "start"}'
You can see the agent receiving the request and starting working in the agent terminal:
INFO: Uvicorn running on http://127.0.0.1:8002 ( Press CTRL+C to quit )
agent> I 'll fetch today' s articles right away!
...
Since we are in sync mode, the curl terminal just waits and hangs. The HTTP connection is kept open, and we eventually receive after 1–2 min the reply:
{
"status" : "success" ,
"output" : "Here's your **AI & Data Engineering Briefing - November 6, 2026** ☕\n\n---\n\n## 🔥 Top Stories\n\n### 1. *" (truncated) ,
"time_elapsed" : 83.74279189109802
}
Our application has just kept an HTTP connection open for 83s, between our request and the final report. During that time, the terminal was blocked.
Timeouts because on most servers, HTTP request will timeout after some time. Could be seconds, or hours, but it will.
Resources because during all this time, the connection was kept open. Keeping a connection open means keeping the server on the client side up, which is stupid because nothing happens. Furthermore, the client cannot perform any work and just has to wait. Not efficient.
The Fire-and-Forget Pattern: HTTP 202 Accepted and Webhooks
As said before, thinking async is thinking “delegation”. You hand over a task to the agent, go do something else, and get notified when the agent is done.
Think about when you clean your clothes. There are several options:
You go to the self-service laundry, and just stay for the whole time the machine is running, in case someone steals your clothes. You stay and you stare at the machine.
Or you can go to a professional cleaner, drop your clothes, get a receipt, and go home. The laundry calls you when it’s ready.
In the latter, the receipt is the task id, and the phone call is a webhook. And what this means concretely for you is that you can live your life between the moment you drop your clothes and the moment you collect them.
Let’s now write a to_async_app wrapper that returns 202 when receiving a new request, with the task id attached (the receipt). We can reuse the previous to_sync_app and only change the /run route:
@api.post ( "/run" , status_code = 202 )
async def run ( req : RunRequest ):
asyncio . create_task ( _run ( ** req . model_dump ()))
return { "status" : "created" , "session_id" : req . session_id }
Two things changed compared to the sync route:
We no longer await the task. asyncio.create_task schedules _run to run in the background and returns immediately, so the endpoint can answer right away.
The task id is simply the session_id : the one the caller sent, or the one RunRequest generated.
If you curl the agent now, it immediately replies with
{"status":"created","session_id":"092e44f1"}
But how does the agent communicate, or “ca

[truncated]
