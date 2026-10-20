---
source: "https://airflow.apache.org/blog/common-ai-provider-1.0/"
hn_url: "https://news.ycombinator.com/item?id=50170494"
title: "Apache Airflow AI Provider 1.0"
article_title: "Airflow AI Provider 1.0: Agents, Sandboxes, Human Approval and Smarter Retries | Apache Airflow"
image: "https://airflow.apache.org/blog/common-ai-provider-1.0/images/common-ai-1.0.png"
author: "jonbaer"
captured_at: "2026-10-20T18:57:54Z"
capture_tool: "hn-digest"
hn_id: 50170494
score: 1
comments: 0
posted_at: "2026-10-20T18:23:24Z"
tags:
  - hacker-news
---

# Apache Airflow AI Provider 1.0

- HN: [50170494](https://news.ycombinator.com/item?id=50170494)
- Source: [airflow.apache.org](https://airflow.apache.org/blog/common-ai-provider-1.0/)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T18:23:24Z

## Translation

Title: Apache Airflow AI Provider 1.0
Article title: Airflow AI Provider 1.0: Agents, Sandboxes, Human Approval and Smarter Retries | Apache Airflow
Description: Apache Airflow's Common AI provider reaches 1.0 and installs on Airflow 2.11. On Airflow 3.3, agents can wait for a person to approve a tool call, and a model can decide whether a failed task retries.

Article text:
Airflow AI Provider 1.0: Agents, Sandboxes, Human Approval and Smarter Retries | Apache Airflow
Skip to main content
Community
Meetups
Documentation
Registry
Use Cases
Announcements
Blog
Airflow Survey 2025
Ecosystem
Community
Meetups
Documentation
Registry
Use Cases
Announcements
Blog
Airflow Survey 2025
Ecosystem
Community
Release
Mon, Oct 19, 2026
Airflow AI Provider 1.0: Agents, Sandboxes, Human Approval and Smarter Retries
Apache Airflow's Common AI provider reaches 1.0 and installs on Airflow 2.11. On Airflow 3.3, agents can wait for a person to approve a tool call, and a model can decide whether a failed task retries.
Common AI ( apache-airflow-providers-common-ai ), the Apache Airflow provider for LLMs and AI agents, is now 1.0. An agent in your pipeline can query your warehouse through read-only tools, run the code it writes in a sandbox rather than in the process that runs the task, and, on Airflow 3.3, wait for a person to approve a tool call before it acts.
Airflow is the open-source platform that schedules and runs data pipelines written in Python. A pipeline is a Dag, each step in it is a task, and a provider is a package that adds integrations to Airflow, in this case with model vendors and agent frameworks. An agent is a model working in a loop: it calls tools, the functions the Dag author allows it to use, such as a SQL query, reads each result and picks its next step until it has an answer. The April post introduced the provider at 0.1.0.
1.0 is also the first release that installs on Airflow 2.11. Airflow 2 is past its end of life, but teams still on it can start using AI in their pipelines now and upgrade on their own schedule. The examples below use Airflow 3 imports, and the 2.11 section lists what to change.
uv pip install "apache-airflow-providers-common-ai[anthropic,sql]"
Pick the extra for your model vendor: openai , anthropic , google or bedrock . The sql extra brings what SQLToolset and the common.sql operators need, and most of the examples below use it. The provider needs Airflow 2.11 or later and Python 3.11 or later. Every example below except the Anthropic SDK one calls the model through an Airflow connection of type pydanticai with the id pydanticai_default , which holds the API key and names the model, such as anthropic:claude-opus-5-5 . The quickstart creates it with one environment variable and runs a first task.
An agent that asks before it posts to Slack
This Dag explains a drop in revenue, and needs Airflow 3.3 for the pause before the Slack call. Its one task is a Python function decorated with @task.agent : the function returns the request, system_prompt holds the standing instructions, and the decorator runs the agent and pushes its final answer to XCom, Airflow’s store for small values passed between tasks. The agent gets two toolsets, each a group of tools: SQL over two tables of a warehouse, and a single Slack method that waits for a person every time the agent calls it.
from __future__ import annotations
from datetime import datetime , timedelta
from airflow.providers.common.ai.toolsets.hook import HookToolset
from airflow.providers.common.ai.toolsets.sql import SQLToolset
from airflow.providers.slack.hooks.slack_webhook import SlackWebhookHook
from airflow.sdk import dag , task
# Read-only SQL over two tables. Queries that touch any other table are refused.
warehouse = SQLToolset ( db_conn_id = "warehouse" , allowed_tables = [ "orders" , "refunds" ])
# One Slack method as a tool, and every call to it waits for a person's yes.
slack = HookToolset (
SlackWebhookHook ( slack_webhook_conn_id = "slack_revenue" ),
allowed_methods = [ "send_text" ],
tool_prefix = "slack" ,
) . approval_required ()
@dag ( schedule = "@daily" , start_date = datetime ( 2026 , 10 , 1 ), catchup = False )
def revenue_drop ():
@task.agent (
llm_conn_id = "pydanticai_default" ,
system_prompt = (
"You are a revenue analyst. Use SQL to find what changed, then post a short "
"summary to Slack: the size of the drop, the cause, and the evidence."
),
toolsets = [ warehouse , slack ],
tool_approval_timeout = timedelta ( hours = 4 ),
)
def explain_revenue_drop ( ds = None ):
return f "Net revenue on { ds } fell compared with the week before. Find out why."
explain_revenue_drop ()
revenue_drop ()
SQLToolset refuses write statements by default. With allowed_tables set, the agent can list, describe and query only orders and refunds : each query is parsed before it runs, and one that reaches another table, including through a subquery, a CTE, a JOIN or information_schema , is refused and handed back to the agent as an error. COPY and inline comments are refused too, and so is any function the SQL parser does not recognize, unless you list it in allowed_functions . The seeded warehouse also has a customers table with names and email addresses. With the toolset above, list_tables returned only ["orders", "refunds"] , and a direct call of the query tool with SELECT name, emai
[truncated]
The query tool failed: Query references tables that are not in the allowed tables list: customers. Use list_tables to see the allowed tables.
The security docs call this an application-level guardrail rather than a security boundary, and ask you to point the connection at a database role that can read only those tables. HookToolset turns the send_text method of Airflow’s Slack webhook hook into a tool, and .approval_required() marks every call to it. The Dag also needs the Slack provider, apache-airflow-providers-slack , and this SQLite warehouse needs apache-airflow-providers-sqlite .
We ran the Dag on Airflow 3.3.2 against a small SQLite warehouse that we seeded by hand: eight days of orders, and on 19 October a burst of “damaged in transit” refunds on EU orders shipped by one carrier, ParcelOne. The model was Claude Opus 5.5. We ran the Dag six times, and each run found the same cause with three to five queries. In the run shown here, the agent listed the tables, read both schemas, ran five queries and asked to post. The task stopped and the request appeared on the task’s Required Actions tab:
The box holds the arguments of the call. These are two of the four bullets in the message the agent wanted to post, unedited:
*Net revenue drop: 2026-10-19 vs 2026-10-12*
• *Size:* Net revenue fell from $22,104.62 to $20,585.59, down *$1,519 (−6.9%)*.
• *Cause:* A spike in *"damaged in transit" refunds on ParcelOne EU orders*. Gross sales did not fall.
We checked the message’s figures against the seeded tables, and all but one match: it puts refunds for other reasons at about $500, where they came to $425. While it waits, the task is in the awaiting_input state, holds no worker slot, and keeps the conversation so far in the task state store. Approve runs the call. Reject does not fail the task: the agent is told the call was denied and carries on. A reason the reviewer types reaches the agent verbatim, in place of the tool’s result, so write it as a note the agent can act on. We rejected the post with “Rejected by the reviewer: we post a cause only after the carrier confirms it. Return the summary instead.” As the reason asked, the agent returned the summary instead of posting, and its answer opened like this:
I didn't post the summary to Slack. The reviewer rejected it because a cause only goes out after the carrier confirms it, and ParcelOne hasn't confirmed anything yet. Here's what I found.
[…]
**Likely cause (not yet confirmed by the carrier):** A spike in "damaged in transit" refunds on ParcelOne orders in the EU. Sales held up; refunds caused the drop.
As in April, the task log keeps every tool call in its own collapsible group, now with the SQL each query ran, and a summary line after each part of the run: model requests, tool calls, tokens, cached tokens and an estimated cost. For this run the estimate was $0.07 across eight model requests, with most of the input tokens read from the prompt cache. The estimate comes from a price table and is not the vendor’s bill.
Tool approval that pauses the task needs Airflow 3.3. On Airflow 2.11 to 3.2, a call to a marked tool fails the task instead. It also cannot be combined with durable=True , enable_hitl_review or a SandboxToolset that starts its own sandbox: with any of them, a call to a marked tool fails the task. Under a CodeMode capability, a marked tool called from run_code does not run, and the model is told so and carries on. On 3.3 a task instance can ask for approval once per Dag run, retries and clears included, and a second request fails the task without retrying. In the upcoming Airflow 3.4, each step that calls a marked tool can pause the task again, and a retry asks again even for calls a reviewer approved in an earlier attempt, so keep gated tools idempotent. tool_approval_timeout fails the task if nobody answers within four hours, or rejects the call with on_tool_approval_timeout="deny" . The tool approval guide covers who may decide.
The upcoming Airflow 3.4 adds task loops, which repeat a group of tasks until a condition holds. In this recording, which Ash Berlin-Taylor made on a development build, an @task.llm task writes a draft, a person reviews it through HITLOperator , and the group runs again with the feedback until the reviewer approves, up to five times. The reviewer asked for emoji, and the second draft had them. The video is cropped and plays at twice the speed.
Let a model decide whether a failed task retries
On Airflow 3.3, every operator takes a retry_policy : after a failure, it decides whether the task retries, after how long, or fails at once. The provider’s two policies hand that decision to a model, so they work on any task, not only AI ones. This Dag has two ordinary tasks that fail on purpose, one with a revoked key and one with a rate limit on its first try:
from __future__ import annotations
from airflow.providers.common.ai.policies.retry import LLMRetryPolicy
from airflow.sdk import dag , get_current_context , task
# A model reads each failure, names its category, and decides: retry after a delay, or fail now.
policy = LLMRetryPolicy ( llm_conn_id = "pydanticai_default" )
@dag ( schedule = None )
def nightly_export ():
@task ( retries = 3 , retry_policy = policy )
def upload_with_revoked_key ():
raise PermissionError ( "403 Forbidden: the access key for bucket exports-eu was revoked" )
@task ( retries = 3 , retry_policy = policy )
def upload_during_throttling ():
if get_current_context ()[ "ti" ] . try_number == 1 :
raise RuntimeError ( "429 Too Many Requests: upload quota exceeded, slow down" )
return "uploaded"
upload_with_revoked_key ()
upload_during_throttling ()
nightly_export ()
LLMRetryPolicy sends the exception’s class and message to the model, with any secrets Airflow has registered masked first, and the model names a category, decides and says why. On Airflow 3.3.2 with Claude Opus 5.5, the two task logs recorded these decisions:
Retry policy decision action=fail reason='auth: A 403 Forbidden caused by a revoked access key is an authentication failure. Retrying will not help because the credentials must be rotated or replaced, so this should not be retried.'
Retry policy decision action=retry reason="rate_limit: HTTP 429 Too Many Requests with an 'upload quota exceeded, slow down' message is API throttling. The request should succeed once the quota window resets, so retry after a 60-second delay. This is attempt 1 of 4, so retries remain."
The task with the revoked key failed on its first try instead of using its three retries, and the throttled task ran again 60 seconds later and succeeded. A decision to retry is also written to the task instance’s retry_reason . If the model call itself fails, the policy falls back to its fallback_rules , exception-type rules that need no model, or to the task’s normal retry behavior when none are set.
LLMRetryPolicy , shown above, lets a text model reason from your instructions. ClassifierRetryPolicy has the model pick one of your categories instead, and e

[truncated]

## Original Extract

Apache Airflow's Common AI provider reaches 1.0 and installs on Airflow 2.11. On Airflow 3.3, agents can wait for a person to approve a tool call, and a model can decide whether a failed task retries.

Airflow AI Provider 1.0: Agents, Sandboxes, Human Approval and Smarter Retries | Apache Airflow
Skip to main content
Community
Meetups
Documentation
Registry
Use Cases
Announcements
Blog
Airflow Survey 2025
Ecosystem
Community
Meetups
Documentation
Registry
Use Cases
Announcements
Blog
Airflow Survey 2025
Ecosystem
Community
Release
Mon, Oct 19, 2026
Airflow AI Provider 1.0: Agents, Sandboxes, Human Approval and Smarter Retries
Apache Airflow's Common AI provider reaches 1.0 and installs on Airflow 2.11. On Airflow 3.3, agents can wait for a person to approve a tool call, and a model can decide whether a failed task retries.
Common AI ( apache-airflow-providers-common-ai ), the Apache Airflow provider for LLMs and AI agents, is now 1.0. An agent in your pipeline can query your warehouse through read-only tools, run the code it writes in a sandbox rather than in the process that runs the task, and, on Airflow 3.3, wait for a person to approve a tool call before it acts.
Airflow is the open-source platform that schedules and runs data pipelines written in Python. A pipeline is a Dag, each step in it is a task, and a provider is a package that adds integrations to Airflow, in this case with model vendors and agent frameworks. An agent is a model working in a loop: it calls tools, the functions the Dag author allows it to use, such as a SQL query, reads each result and picks its next step until it has an answer. The April post introduced the provider at 0.1.0.
1.0 is also the first release that installs on Airflow 2.11. Airflow 2 is past its end of life, but teams still on it can start using AI in their pipelines now and upgrade on their own schedule. The examples below use Airflow 3 imports, and the 2.11 section lists what to change.
uv pip install "apache-airflow-providers-common-ai[anthropic,sql]"
Pick the extra for your model vendor: openai , anthropic , google or bedrock . The sql extra brings what SQLToolset and the common.sql operators need, and most of the examples below use it. The provider needs Airflow 2.11 or later and Python 3.11 or later. Every example below except the Anthropic SDK one calls the model through an Airflow connection of type pydanticai with the id pydanticai_default , which holds the API key and names the model, such as anthropic:claude-opus-5-5 . The quickstart creates it with one environment variable and runs a first task.
An agent that asks before it posts to Slack
This Dag explains a drop in revenue, and needs Airflow 3.3 for the pause before the Slack call. Its one task is a Python function decorated with @task.agent : the function returns the request, system_prompt holds the standing instructions, and the decorator runs the agent and pushes its final answer to XCom, Airflow’s store for small values passed between tasks. The agent gets two toolsets, each a group of tools: SQL over two tables of a warehouse, and a single Slack method that waits for a person every time the agent calls it.
from __future__ import annotations
from datetime import datetime , timedelta
from airflow.providers.common.ai.toolsets.hook import HookToolset
from airflow.providers.common.ai.toolsets.sql import SQLToolset
from airflow.providers.slack.hooks.slack_webhook import SlackWebhookHook
from airflow.sdk import dag , task
# Read-only SQL over two tables. Queries that touch any other table are refused.
warehouse = SQLToolset ( db_conn_id = "warehouse" , allowed_tables = [ "orders" , "refunds" ])
# One Slack method as a tool, and every call to it waits for a person's yes.
slack = HookToolset (
SlackWebhookHook ( slack_webhook_conn_id = "slack_revenue" ),
allowed_methods = [ "send_text" ],
tool_prefix = "slack" ,
) . approval_required ()
@dag ( schedule = "@daily" , start_date = datetime ( 2026 , 10 , 1 ), catchup = False )
def revenue_drop ():
@task.agent (
llm_conn_id = "pydanticai_default" ,
system_prompt = (
"You are a revenue analyst. Use SQL to find what changed, then post a short "
"summary to Slack: the size of the drop, the cause, and the evidence."
),
toolsets = [ warehouse , slack ],
tool_approval_timeout = timedelta ( hours = 4 ),
)
def explain_revenue_drop ( ds = None ):
return f "Net revenue on { ds } fell compared with the week before. Find out why."
explain_revenue_drop ()
revenue_drop ()
SQLToolset refuses write statements by default. With allowed_tables set, the agent can list, describe and query only orders and refunds : each query is parsed before it runs, and one that reaches another table, including through a subquery, a CTE, a JOIN or information_schema , is refused and handed back to the agent as an error. COPY and inline comments are refused too, and so is any function the SQL parser does not recognize, unless you list it in allowed_functions . The seeded warehouse also has a customers table with names and email addresses. With the toolset above, list_tables returned only ["orders", "refunds"] , and a direct call of the query tool with SELECT name, emai
[truncated]
The query tool failed: Query references tables that are not in the allowed tables list: customers. Use list_tables to see the allowed tables.
The security docs call this an application-level guardrail rather than a security boundary, and ask you to point the connection at a database role that can read only those tables. HookToolset turns the send_text method of Airflow’s Slack webhook hook into a tool, and .approval_required() marks every call to it. The Dag also needs the Slack provider, apache-airflow-providers-slack , and this SQLite warehouse needs apache-airflow-providers-sqlite .
We ran the Dag on Airflow 3.3.2 against a small SQLite warehouse that we seeded by hand: eight days of orders, and on 19 October a burst of “damaged in transit” refunds on EU orders shipped by one carrier, ParcelOne. The model was Claude Opus 5.5. We ran the Dag six times, and each run found the same cause with three to five queries. In the run shown here, the agent listed the tables, read both schemas, ran five queries and asked to post. The task stopped and the request appeared on the task’s Required Actions tab:
The box holds the arguments of the call. These are two of the four bullets in the message the agent wanted to post, unedited:
*Net revenue drop: 2026-10-19 vs 2026-10-12*
• *Size:* Net revenue fell from $22,104.62 to $20,585.59, down *$1,519 (−6.9%)*.
• *Cause:* A spike in *"damaged in transit" refunds on ParcelOne EU orders*. Gross sales did not fall.
We checked the message’s figures against the seeded tables, and all but one match: it puts refunds for other reasons at about $500, where they came to $425. While it waits, the task is in the awaiting_input state, holds no worker slot, and keeps the conversation so far in the task state store. Approve runs the call. Reject does not fail the task: the agent is told the call was denied and carries on. A reason the reviewer types reaches the agent verbatim, in place of the tool’s result, so write it as a note the agent can act on. We rejected the post with “Rejected by the reviewer: we post a cause only after the carrier confirms it. Return the summary instead.” As the reason asked, the agent returned the summary instead of posting, and its answer opened like this:
I didn't post the summary to Slack. The reviewer rejected it because a cause only goes out after the carrier confirms it, and ParcelOne hasn't confirmed anything yet. Here's what I found.
[…]
**Likely cause (not yet confirmed by the carrier):** A spike in "damaged in transit" refunds on ParcelOne orders in the EU. Sales held up; refunds caused the drop.
As in April, the task log keeps every tool call in its own collapsible group, now with the SQL each query ran, and a summary line after each part of the run: model requests, tool calls, tokens, cached tokens and an estimated cost. For this run the estimate was $0.07 across eight model requests, with most of the input tokens read from the prompt cache. The estimate comes from a price table and is not the vendor’s bill.
Tool approval that pauses the task needs Airflow 3.3. On Airflow 2.11 to 3.2, a call to a marked tool fails the task instead. It also cannot be combined with durable=True , enable_hitl_review or a SandboxToolset that starts its own sandbox: with any of them, a call to a marked tool fails the task. Under a CodeMode capability, a marked tool called from run_code does not run, and the model is told so and carries on. On 3.3 a task instance can ask for approval once per Dag run, retries and clears included, and a second request fails the task without retrying. In the upcoming Airflow 3.4, each step that calls a marked tool can pause the task again, and a retry asks again even for calls a reviewer approved in an earlier attempt, so keep gated tools idempotent. tool_approval_timeout fails the task if nobody answers within four hours, or rejects the call with on_tool_approval_timeout="deny" . The tool approval guide covers who may decide.
The upcoming Airflow 3.4 adds task loops, which repeat a group of tasks until a condition holds. In this recording, which Ash Berlin-Taylor made on a development build, an @task.llm task writes a draft, a person reviews it through HITLOperator , and the group runs again with the feedback until the reviewer approves, up to five times. The reviewer asked for emoji, and the second draft had them. The video is cropped and plays at twice the speed.
Let a model decide whether a failed task retries
On Airflow 3.3, every operator takes a retry_policy : after a failure, it decides whether the task retries, after how long, or fails at once. The provider’s two policies hand that decision to a model, so they work on any task, not only AI ones. This Dag has two ordinary tasks that fail on purpose, one with a revoked key and one with a rate limit on its first try:
from __future__ import annotations
from airflow.providers.common.ai.policies.retry import LLMRetryPolicy
from airflow.sdk import dag , get_current_context , task
# A model reads each failure, names its category, and decides: retry after a delay, or fail now.
policy = LLMRetryPolicy ( llm_conn_id = "pydanticai_default" )
@dag ( schedule = None )
def nightly_export ():
@task ( retries = 3 , retry_policy = policy )
def upload_with_revoked_key ():
raise PermissionError ( "403 Forbidden: the access key for bucket exports-eu was revoked" )
@task ( retries = 3 , retry_policy = policy )
def upload_during_throttling ():
if get_current_context ()[ "ti" ] . try_number == 1 :
raise RuntimeError ( "429 Too Many Requests: upload quota exceeded, slow down" )
return "uploaded"
upload_with_revoked_key ()
upload_during_throttling ()
nightly_export ()
LLMRetryPolicy sends the exception’s class and message to the model, with any secrets Airflow has registered masked first, and the model names a category, decides and says why. On Airflow 3.3.2 with Claude Opus 5.5, the two task logs recorded these decisions:
Retry policy decision action=fail reason='auth: A 403 Forbidden caused by a revoked access key is an authentication failure. Retrying will not help because the credentials must be rotated or replaced, so this should not be retried.'
Retry policy decision action=retry reason="rate_limit: HTTP 429 Too Many Requests with an 'upload quota exceeded, slow down' message is API throttling. The request should succeed once the quota window resets, so retry after a 60-second delay. This is attempt 1 of 4, so retries remain."
The task with the revoked key failed on its first try instead of using its three retries, and the throttled task ran again 60 seconds later and succeeded. A decision to retry is also written to the task instance’s retry_reason . If the model call itself fails, the policy falls back to its fallback_rules , exception-type rules that need no model, or to the task’s normal retry behavior when none are set.
LLMRetryPolicy , shown above, lets a text model reason from your instructions. ClassifierRetryPolicy has the model pick one of your categories instead, and e

[truncated]
