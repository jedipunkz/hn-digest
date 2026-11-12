---
source: "https://blog.kwatch.dev/kwatch-v1-2-0"
hn_url: "https://news.ycombinator.com/item?id=50468266"
title: "Kwatch v1.2.0: Kubernetes failure root cause, now with a read-only dashboard"
article_title: "kwatch v1.2.0: a dashboard and clearer root causes"
image: "https://blog.kwatch.dev/og/kwatch-v1-2-0.png"
author: "abahmed"
captured_at: "2026-11-12T15:01:17Z"
capture_tool: "hn-digest"
hn_id: 50468266
score: 1
comments: 0
posted_at: "2026-11-12T14:51:54Z"
tags:
  - hacker-news
---

# Kwatch v1.2.0: Kubernetes failure root cause, now with a read-only dashboard

- HN: [50468266](https://news.ycombinator.com/item?id=50468266)
- Source: [blog.kwatch.dev](https://blog.kwatch.dev/kwatch-v1-2-0)
- Score: 1
- Comments: 0
- Posted: 2026-11-12T14:51:54Z

## Translation

Title: Kwatch v1.2.0: Kubernetes failure root cause, now with a read-only dashboard
Article title: kwatch v1.2.0: a dashboard and clearer root causes
Description: kwatch v1.2.0 adds a read-only Pro dashboard with browser notifications, sharper root causes, quieter alerts and more provider options.

Article text:
Skip to content kwatch Search the blog Ctrl K Docs Blog Pro Get help Sign in Blog / Releases
kwatch v1.2.0: a dashboard and clearer root causes
kwatch v1.2.0 adds a read-only Pro dashboard with browser notifications, sharper root causes, quieter alerts and more provider options.
kwatch v1.2.0 adds a read-only dashboard, names more kinds of cause, and groups related alerts into one message. It follows v1.1.0.
The installer walks you through the upgrade. Choose Upgrade from its menu.
bash <( curl -fsSL https://kwatch.dev/kwatch.sh )
Already use Helm? The chart ships with v1.2.0 too.
The dashboard runs in your cluster and opens in your browser. It is read-only: it shows what kwatch sees and changes nothing. It is part of Pro, which is free during early access. Turn it on from the installer.
Active issues, with the cause and the logs behind each one.
Browser notifications tell you when an issue opens, escalates or resolves, even when the dashboard is closed.
Sign in with a password, and serve it over TLS or behind your own proxy with dashboard.trustedProxies .
The dashboard guide covers setup and sharing it with your team.
kwatch v1.2.0 runs an investigation across pods, workloads and the cluster. It follows a failure upstream and names more kinds of cause, including:
a dependency the app cannot reach, or a backend that is not ready
a previous container that failed
an image pull that was throttled
a rollout or change that landed just before the failure
Pro shows the evidence behind the cause. It can also sample the logs of running pods for bursts of errors, and investigation.logs and investigation.sampleLogs let you choose what kwatch reads. Community messages name the type of cause; Pro adds the evidence and the command to run next.
Related alerts group into one incident, with plainer wording and fewer duplicates. Two settings keep production quiet and old noise out of your pager:
app.environment: non-production marks a cluster as non-production, so nothing pages and every incident stays at Needs attention, the level below a page.
oldProblemDays (30 by default) handles problems that were already there when kwatch first saw them. Low ones open no issue, and older needs-attention ones drop to Low.
ntfy, Gotify and Pushover: resolvedPriority sets how loud a resolved message is.
SES and SNS: use a pod IAM role (IRSA or EKS Pod Identity) and leave the keys out. An endpoint override is available.
PagerDuty: a message that raises the page uses its own paging key, and a resolve closes both.
Issue trackers: GitHub, GitLab, Gitea, Jira and ClickUp close the issue when the problem recovers and reopen the right one if it returns.
Webex: set exactly one of roomId or toPersonEmail , and kwatch checks it at startup.
KWATCH_LOG_FORMAT=json writes one JSON object per log line.
The flight recording, the saved record kwatch replay-export writes, keeps scrubbed text, with names, hosts, emails, images and credentials hidden.
kwatch lint --strict checks your config and flags rules that name reasons kwatch no longer raises.
kwatch-scorecard exits with an error on empty input, and -allow-empty restores the old exit.
Upgrading from v1.1.0 is one command. The upgrading page lists everything in one place. The short version:
The /status text page is part of Pro. /health stays in every edition.
Eight reasons are no longer raised: ContainerCannotRun , Killed , DeploymentProgressingFalse , DeploymentAvailableFalse , StatefulSetConditionFailure , DaemonSetConditionFailure , ResourceHigh and ResourceCritical . Rules that name them still load, and kwatch lint --strict points them out.
The ntfy priority is unset by default, so ntfy uses its own default of 3.
With several providers, Community picks a pager first, then chat and other providers, then email.
kwatch is source available under ELv2.
One command, and it walks you through it. Community is free forever. Pro is free during early access.
Try a symptom, such as CrashLoopBackOff or OOMKilled .
Up Down to move Enter to open Esc to close

## Original Extract

kwatch v1.2.0 adds a read-only Pro dashboard with browser notifications, sharper root causes, quieter alerts and more provider options.

Skip to content kwatch Search the blog Ctrl K Docs Blog Pro Get help Sign in Blog / Releases
kwatch v1.2.0: a dashboard and clearer root causes
kwatch v1.2.0 adds a read-only Pro dashboard with browser notifications, sharper root causes, quieter alerts and more provider options.
kwatch v1.2.0 adds a read-only dashboard, names more kinds of cause, and groups related alerts into one message. It follows v1.1.0.
The installer walks you through the upgrade. Choose Upgrade from its menu.
bash <( curl -fsSL https://kwatch.dev/kwatch.sh )
Already use Helm? The chart ships with v1.2.0 too.
The dashboard runs in your cluster and opens in your browser. It is read-only: it shows what kwatch sees and changes nothing. It is part of Pro, which is free during early access. Turn it on from the installer.
Active issues, with the cause and the logs behind each one.
Browser notifications tell you when an issue opens, escalates or resolves, even when the dashboard is closed.
Sign in with a password, and serve it over TLS or behind your own proxy with dashboard.trustedProxies .
The dashboard guide covers setup and sharing it with your team.
kwatch v1.2.0 runs an investigation across pods, workloads and the cluster. It follows a failure upstream and names more kinds of cause, including:
a dependency the app cannot reach, or a backend that is not ready
a previous container that failed
an image pull that was throttled
a rollout or change that landed just before the failure
Pro shows the evidence behind the cause. It can also sample the logs of running pods for bursts of errors, and investigation.logs and investigation.sampleLogs let you choose what kwatch reads. Community messages name the type of cause; Pro adds the evidence and the command to run next.
Related alerts group into one incident, with plainer wording and fewer duplicates. Two settings keep production quiet and old noise out of your pager:
app.environment: non-production marks a cluster as non-production, so nothing pages and every incident stays at Needs attention, the level below a page.
oldProblemDays (30 by default) handles problems that were already there when kwatch first saw them. Low ones open no issue, and older needs-attention ones drop to Low.
ntfy, Gotify and Pushover: resolvedPriority sets how loud a resolved message is.
SES and SNS: use a pod IAM role (IRSA or EKS Pod Identity) and leave the keys out. An endpoint override is available.
PagerDuty: a message that raises the page uses its own paging key, and a resolve closes both.
Issue trackers: GitHub, GitLab, Gitea, Jira and ClickUp close the issue when the problem recovers and reopen the right one if it returns.
Webex: set exactly one of roomId or toPersonEmail , and kwatch checks it at startup.
KWATCH_LOG_FORMAT=json writes one JSON object per log line.
The flight recording, the saved record kwatch replay-export writes, keeps scrubbed text, with names, hosts, emails, images and credentials hidden.
kwatch lint --strict checks your config and flags rules that name reasons kwatch no longer raises.
kwatch-scorecard exits with an error on empty input, and -allow-empty restores the old exit.
Upgrading from v1.1.0 is one command. The upgrading page lists everything in one place. The short version:
The /status text page is part of Pro. /health stays in every edition.
Eight reasons are no longer raised: ContainerCannotRun , Killed , DeploymentProgressingFalse , DeploymentAvailableFalse , StatefulSetConditionFailure , DaemonSetConditionFailure , ResourceHigh and ResourceCritical . Rules that name them still load, and kwatch lint --strict points them out.
The ntfy priority is unset by default, so ntfy uses its own default of 3.
With several providers, Community picks a pager first, then chat and other providers, then email.
kwatch is source available under ELv2.
One command, and it walks you through it. Community is free forever. Pro is free during early access.
Try a symptom, such as CrashLoopBackOff or OOMKilled .
Up Down to move Enter to open Esc to close
