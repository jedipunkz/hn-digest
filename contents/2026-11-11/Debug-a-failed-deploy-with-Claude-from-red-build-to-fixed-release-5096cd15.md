---
source: "https://railyard.run/blog/debug-failed-deploy-with-claude/"
hn_url: "https://news.ycombinator.com/item?id=50459640"
title: "Debug a failed deploy with Claude: from red build to fixed release"
article_title: "Debug a failed deploy with Claude: from red build to fixed release in one conversation · Railyard"
image: "https://railyard.run/covers/en/debug-failed-deploy-with-claude-5339d2656b.png"
author: "RailyardRun"
captured_at: "2026-11-11T23:20:57Z"
capture_tool: "hn-digest"
hn_id: 50459640
score: 1
comments: 0
posted_at: "2026-11-11T22:26:51Z"
tags:
  - hacker-news
---

# Debug a failed deploy with Claude: from red build to fixed release

- HN: [50459640](https://news.ycombinator.com/item?id=50459640)
- Source: [railyard.run](https://railyard.run/blog/debug-failed-deploy-with-claude/)
- Score: 1
- Comments: 0
- Posted: 2026-11-11T22:26:51Z

## Translation

Title: Debug a failed deploy with Claude: from red build to fixed release
Article title: Debug a failed deploy with Claude: from red build to fixed release in one conversation · Railyard
Description: A practical workflow for diagnosing failed deploys with an AI agent over MCP: read the deploy error and build log, correlate runtime logs and traffic, fix…

Article text:
Skip to content 🔥 Agencies & studios Run unlimited client servers for a flat $99/mo — zero per-seat or bandwidth tax. Railyard for agencies → Railyard Managed services Browse all services Hundreds of open-source apps, one click Self-hosted alternatives Replace Slack, Notion, Zapier… Deploy any Git repo or image GitHub, git push, CLI, Docker image All deploy guides Rails, Next.js, Django, Laravel, Phoenix, Go… Open source What we build in the open Features All features Deploys, databases, logs, metrics, backups How it works From repo to a live URL Security SSO, roles, audit log, ActiveProtect Compare to hosted PaaS Side by side with Heroku, Render, Fly… For agencies Client servers, one flat bill For freelancers Every client on one flat-price server For DevOps teams Terraform, API, SSO, audit log For enterprises SSO, SCIM, audit log, your cloud AI AI agents (MCP, CLI, API) Let agents deploy and operate For vibe coders Built it with AI? Ship it, no DevOps Self-host AI apps AnythingLLM, Dify and more, one click Resources Free tools railyard.yml generator, repo check, cost calculators Docs Guides and reference Tutorials Self-host, deploy and migrate Migrate to Railyard From Heroku, Render, Railway, Fly.io Alternatives Heroku, Vercel, Fly.io, Render, Railway… Hosting cost guide Hosted PaaS vs your own server For indie hackers Many apps per server, zero markup For startups Your whole tool stack, self-hosted Blog Changelog What shipped, every week Roadmap What's next Status Pricing Sign in EN English
Deploy free sandbox → Free sandbox ← Blog Guides & tutorials • 4 min read Debug a failed deploy with Claude: from red build to fixed release in one conversation
Almokhtar Bekkour , Founder, Systems Architect
Follow this on your own server. Start free, no card needed
A deploy goes red at the worst moment. The fix is usually one line, but finding the line means opening the deploy page, scrolling a 2,000-line build log, checking whether the old release is still serving, and grepping runtime logs for the first error after the switch. An agent with read access to those same sources does that sweep in seconds and comes back with the line.
This tutorial shows the workflow on Railyard: what to ask, which tools the agent uses, how to read what it finds, and when to roll back instead of fixing forward.
Connect Railyard to Claude. In Claude, go to Customize → Connectors → Railyard → Connect . In Claude Code:
claude mcp add --transport http railyard https://app.railyard.run/api/mcp
Diagnosis only needs read access. To let Claude fix the cause and redeploy too, connect with deploy , or write if variables need to change.
The fastest path: the built-in prompt
Railyard’s MCP server ships a prompt called debug_failed_deploy . In clients that show MCP prompts, pick it from the prompt menu (in Claude Code, type / and look for the Railyard prompts), give it the app name, and send. It tells Claude to:
find the most recent failed deploy ( list_deploys with status: failed )
read its recorded error ( get_deploy ) and build log ( deploy_logs )
check runtime errors around that time ( app_logs )
explain the root cause and the fix, and ask before changing anything
The rest of this post is the same workflow done by hand, so you know what each step finds.
> The storefront deploy just failed. What happened?
list_deploys (filtered to the app and status: failed ) returns the failing deploy with its trigger and commit:
#1214 failed git_push a91c3e2 started 14:02, finished 14:05
get_deploy adds the recorded error, who or what triggered it, and how long it ran. The error line often answers the question by itself:
error: "Release command failed: bin/rails db:migrate exited 1"
Step 2: Read the build and release log
deploy_logs returns the last 200 lines of the deploy’s own log by default, up to 500. That’s where the real cause lives:
== 20260729140122 AddIndexToOrdersEmail: migrating ===========================
-- add_index(:orders, :email, {:unique=>true})
rails aborted!
PG::UniqueViolation: ERROR: could not create unique index "index_orders_on_email"
DETAIL: Key (email)=(sam@example.com) is duplicated.
Claude reads that the migration added a unique index while the existing data has duplicates. That’s a data problem, not a code typo, and it changes the fix.
The failure patterns worth knowing
Step 3: Is production still up?
A failed deploy on Railyard leaves the previous release serving. Confirm that before panicking:
> Is storefront still serving traffic? Any errors in the last 15 minutes?
That’s three tools:
list_processes : are the containers healthy, restarting, or crash looping?
app_traffic with range: live : request rate, 5xx rate and latency, compared with the previous window
app_logs with level: error and range: 1h : anything new since the failed deploy
If traffic and error rates look normal, there’s no fire. You have time to fix forward.
Step 4: Fix forward, or roll back
Fix forward when the cause is small and understood, like the duplicate data above. Claude can show the duplicates with a read-only query:
> Which emails are duplicated in orders?
query_database runs read-only SQL, and needs a write token:
SELECT email, count ( * ) FROM orders GROUP BY email HAVING count ( * ) > 1 ;
You decide how to clean the data. That part stays yours. Then push a fixed migration and redeploy, or have Claude redeploy a specific commit with deploy_app and ref .
Roll back when production is actually broken, meaning a bad deploy did go out:
> Roll storefront back to the last good release.
rollback_app redeploys an earlier successful deploy, and by default restores the variables that release ran with. See roll back a bad release from chat for the details.
After the redeploy, ask Claude to follow it to the end with get_deploy and deploy_logs , then re-run the health check from Step 3. Close the loop with list_activity , which shows the deploy, who or what triggered it, and any variable changes. It’s useful when you write the incident note.
Why this works better than pasting logs into a chat
No copy-paste ceiling. The agent fetches exactly the window it needs, and the server bounds each read so it never floods the context.
Correlation across sources. Build log, runtime log, traffic and process health in one answer, matched up in time.
Actions in the same place. Once you agree on the fix, Claude can redeploy or roll back without you switching tools, and nothing destructive happens without your confirmation.
It can’t make it worse on its own. Reading needs only read . Restoring a backup or deleting anything always waits for an admin’s approval.
Related: deploy from Claude Code step by step , and how the MCP server works .
Why people pick Railyard for this
Your own cloud account. Deploy to DigitalOcean, Hetzner, Vultr, AWS, OVHcloud, or any Linux server over SSH, billed directly by your provider with no markup.
One flat price per team. Free is free forever: one server plus a database server, three apps, no card. Pro is $19 a month for up to five servers and unlimited apps. Team is $99 a month with unlimited servers. No per-app, per-deploy, per-seat or per-gigabyte charges.
One-click templates. Deploy open-source apps, like databases, analytics and dev tools, to your own server in one click.
Managed databases. Postgres, MySQL and Redis with scheduled, restore-verified backups and point-in-time recovery.
Push-to-deploy. Deploy from GitHub, GitLab or the CLI, with automatic HTTPS and zero-downtime releases.
Get this guide by email, plus 4 short emails on moving off Heroku
One email to confirm, then the guide and four notes over two weeks: hidden costs, what moving takes, how to check your app, and an offer to move it with you. Unsubscribe in one click.
The fastest path: the built-in prompt
Step 2: Read the build and release log
Step 3: Is production still up?
Step 4: Fix forward, or roll back
Why this works better than pasting logs into a chat
Test any of our 1037 open-source templates or your own Git repo on a free 7-day 4 GB sandbox server right now. Zero cloud setup.
Let Claude read the logs for you.
Connect Railyard to Claude and ask why the last deploy failed. It reads the build log, the runtime logs and the traffic, then tells you the cause.
Launch free 7-day test server See how a deploy runs Free plan, no card, and 7 days of Pro for every new team. One flat price per team.
Railyard Deploy full-stack apps to servers in your own cloud account. hello@railyard.run
Esc to close · ↵ to open a result
Theme Colors Dark Classic light Match my system × Before you go: try it on a free server.
We lend you a server for 7 days. No card, no cloud account. Deploy a repo or a one-click app and see it live.

## Original Extract

A practical workflow for diagnosing failed deploys with an AI agent over MCP: read the deploy error and build log, correlate runtime logs and traffic, fix…

Skip to content 🔥 Agencies & studios Run unlimited client servers for a flat $99/mo — zero per-seat or bandwidth tax. Railyard for agencies → Railyard Managed services Browse all services Hundreds of open-source apps, one click Self-hosted alternatives Replace Slack, Notion, Zapier… Deploy any Git repo or image GitHub, git push, CLI, Docker image All deploy guides Rails, Next.js, Django, Laravel, Phoenix, Go… Open source What we build in the open Features All features Deploys, databases, logs, metrics, backups How it works From repo to a live URL Security SSO, roles, audit log, ActiveProtect Compare to hosted PaaS Side by side with Heroku, Render, Fly… For agencies Client servers, one flat bill For freelancers Every client on one flat-price server For DevOps teams Terraform, API, SSO, audit log For enterprises SSO, SCIM, audit log, your cloud AI AI agents (MCP, CLI, API) Let agents deploy and operate For vibe coders Built it with AI? Ship it, no DevOps Self-host AI apps AnythingLLM, Dify and more, one click Resources Free tools railyard.yml generator, repo check, cost calculators Docs Guides and reference Tutorials Self-host, deploy and migrate Migrate to Railyard From Heroku, Render, Railway, Fly.io Alternatives Heroku, Vercel, Fly.io, Render, Railway… Hosting cost guide Hosted PaaS vs your own server For indie hackers Many apps per server, zero markup For startups Your whole tool stack, self-hosted Blog Changelog What shipped, every week Roadmap What's next Status Pricing Sign in EN English
Deploy free sandbox → Free sandbox ← Blog Guides & tutorials • 4 min read Debug a failed deploy with Claude: from red build to fixed release in one conversation
Almokhtar Bekkour , Founder, Systems Architect
Follow this on your own server. Start free, no card needed
A deploy goes red at the worst moment. The fix is usually one line, but finding the line means opening the deploy page, scrolling a 2,000-line build log, checking whether the old release is still serving, and grepping runtime logs for the first error after the switch. An agent with read access to those same sources does that sweep in seconds and comes back with the line.
This tutorial shows the workflow on Railyard: what to ask, which tools the agent uses, how to read what it finds, and when to roll back instead of fixing forward.
Connect Railyard to Claude. In Claude, go to Customize → Connectors → Railyard → Connect . In Claude Code:
claude mcp add --transport http railyard https://app.railyard.run/api/mcp
Diagnosis only needs read access. To let Claude fix the cause and redeploy too, connect with deploy , or write if variables need to change.
The fastest path: the built-in prompt
Railyard’s MCP server ships a prompt called debug_failed_deploy . In clients that show MCP prompts, pick it from the prompt menu (in Claude Code, type / and look for the Railyard prompts), give it the app name, and send. It tells Claude to:
find the most recent failed deploy ( list_deploys with status: failed )
read its recorded error ( get_deploy ) and build log ( deploy_logs )
check runtime errors around that time ( app_logs )
explain the root cause and the fix, and ask before changing anything
The rest of this post is the same workflow done by hand, so you know what each step finds.
> The storefront deploy just failed. What happened?
list_deploys (filtered to the app and status: failed ) returns the failing deploy with its trigger and commit:
#1214 failed git_push a91c3e2 started 14:02, finished 14:05
get_deploy adds the recorded error, who or what triggered it, and how long it ran. The error line often answers the question by itself:
error: "Release command failed: bin/rails db:migrate exited 1"
Step 2: Read the build and release log
deploy_logs returns the last 200 lines of the deploy’s own log by default, up to 500. That’s where the real cause lives:
== 20260729140122 AddIndexToOrdersEmail: migrating ===========================
-- add_index(:orders, :email, {:unique=>true})
rails aborted!
PG::UniqueViolation: ERROR: could not create unique index "index_orders_on_email"
DETAIL: Key (email)=(sam@example.com) is duplicated.
Claude reads that the migration added a unique index while the existing data has duplicates. That’s a data problem, not a code typo, and it changes the fix.
The failure patterns worth knowing
Step 3: Is production still up?
A failed deploy on Railyard leaves the previous release serving. Confirm that before panicking:
> Is storefront still serving traffic? Any errors in the last 15 minutes?
That’s three tools:
list_processes : are the containers healthy, restarting, or crash looping?
app_traffic with range: live : request rate, 5xx rate and latency, compared with the previous window
app_logs with level: error and range: 1h : anything new since the failed deploy
If traffic and error rates look normal, there’s no fire. You have time to fix forward.
Step 4: Fix forward, or roll back
Fix forward when the cause is small and understood, like the duplicate data above. Claude can show the duplicates with a read-only query:
> Which emails are duplicated in orders?
query_database runs read-only SQL, and needs a write token:
SELECT email, count ( * ) FROM orders GROUP BY email HAVING count ( * ) > 1 ;
You decide how to clean the data. That part stays yours. Then push a fixed migration and redeploy, or have Claude redeploy a specific commit with deploy_app and ref .
Roll back when production is actually broken, meaning a bad deploy did go out:
> Roll storefront back to the last good release.
rollback_app redeploys an earlier successful deploy, and by default restores the variables that release ran with. See roll back a bad release from chat for the details.
After the redeploy, ask Claude to follow it to the end with get_deploy and deploy_logs , then re-run the health check from Step 3. Close the loop with list_activity , which shows the deploy, who or what triggered it, and any variable changes. It’s useful when you write the incident note.
Why this works better than pasting logs into a chat
No copy-paste ceiling. The agent fetches exactly the window it needs, and the server bounds each read so it never floods the context.
Correlation across sources. Build log, runtime log, traffic and process health in one answer, matched up in time.
Actions in the same place. Once you agree on the fix, Claude can redeploy or roll back without you switching tools, and nothing destructive happens without your confirmation.
It can’t make it worse on its own. Reading needs only read . Restoring a backup or deleting anything always waits for an admin’s approval.
Related: deploy from Claude Code step by step , and how the MCP server works .
Why people pick Railyard for this
Your own cloud account. Deploy to DigitalOcean, Hetzner, Vultr, AWS, OVHcloud, or any Linux server over SSH, billed directly by your provider with no markup.
One flat price per team. Free is free forever: one server plus a database server, three apps, no card. Pro is $19 a month for up to five servers and unlimited apps. Team is $99 a month with unlimited servers. No per-app, per-deploy, per-seat or per-gigabyte charges.
One-click templates. Deploy open-source apps, like databases, analytics and dev tools, to your own server in one click.
Managed databases. Postgres, MySQL and Redis with scheduled, restore-verified backups and point-in-time recovery.
Push-to-deploy. Deploy from GitHub, GitLab or the CLI, with automatic HTTPS and zero-downtime releases.
Get this guide by email, plus 4 short emails on moving off Heroku
One email to confirm, then the guide and four notes over two weeks: hidden costs, what moving takes, how to check your app, and an offer to move it with you. Unsubscribe in one click.
The fastest path: the built-in prompt
Step 2: Read the build and release log
Step 3: Is production still up?
Step 4: Fix forward, or roll back
Why this works better than pasting logs into a chat
Test any of our 1037 open-source templates or your own Git repo on a free 7-day 4 GB sandbox server right now. Zero cloud setup.
Let Claude read the logs for you.
Connect Railyard to Claude and ask why the last deploy failed. It reads the build log, the runtime logs and the traffic, then tells you the cause.
Launch free 7-day test server See how a deploy runs Free plan, no card, and 7 days of Pro for every new team. One flat price per team.
Railyard Deploy full-stack apps to servers in your own cloud account. hello@railyard.run
Esc to close · ↵ to open a result
Theme Colors Dark Classic light Match my system × Before you go: try it on a free server.
We lend you a server for 7 days. No card, no cloud account. Deploy a repo or a one-click app and see it live.
