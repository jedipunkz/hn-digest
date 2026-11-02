---
source: "https://cloudopsbench.ai/blog/failure-taxonomy"
hn_url: "https://news.ycombinator.com/item?id=50340235"
title: "How Well Do Agents Perform on DevOps Tasks?"
article_title: "CloudOpsBench | How Well Do Agents Perform on DevOps Tasks?"
image: ""
author: "matt_d"
captured_at: "2026-11-02T17:53:22Z"
capture_tool: "hn-digest"
hn_id: 50340235
score: 1
comments: 0
posted_at: "2026-11-02T17:24:32Z"
tags:
  - hacker-news
---

# How Well Do Agents Perform on DevOps Tasks?

- HN: [50340235](https://news.ycombinator.com/item?id=50340235)
- Source: [cloudopsbench.ai](https://cloudopsbench.ai/blog/failure-taxonomy)
- Score: 1
- Comments: 0
- Posted: 2026-11-02T17:24:32Z

## Translation

Title: How Well Do Agents Perform on DevOps Tasks?
Article title: CloudOpsBench | How Well Do Agents Perform on DevOps Tasks?
Description: Frontier agents on CloudOpsBench have the knowledge cloud tasks require and fail mainly at judging when a task is done. A three-step taxonomy of where they fail.

Article text:
CloudOpsBench | How Well Do Agents Perform on DevOps Tasks? Skip to content CloudOps Bench . Leaderboard About Blog Contributors GitHub Blog How Well Do Agents Perform on DevOps Tasks?
CloudOpsBench is a new benchmark that measures how well AI agents handle DevOps tasks in the cloud: repairing failures, migrating services, rotating cryptographic keys, rolling out updates, and decommissioning projects. Unlike typical software engineering (SWE) tasks, DevOps tasks are safety-critical, hard to grade with a test suite, and require a deep understanding of the environment.
CloudOpsBench v1.0 consists of 45 DevOps tasks on AWS and has two distinguishing features: (1) It is a frontier benchmark. Even the strongest agents leave many tasks unsolved: Claude Fable 5.1 completed 29 of the 45 tasks, and GPT-6 Astra completed 26. (2) It runs identically in the real AWS cloud and in a high-fidelity emulator, with the same results in both. Agents can be evaluated quickly and safely, with no real cloud account involved.
The tasks span a range of representative DevOps workflows: decommissioning projects (13 tasks), securing access (12), migrating and deploying services (7), configuring and monitoring services (6), and repairing failures (7).
In order to accomplish a DevOps task, there are three components that are important for DevOps engineers—whether a human or an agent.
Knowing. The agent must have a deep understanding of cloud services. As a simple example, if the agent is prompted to delete every queue from a project, it must know which cloud services may have queues. In AWS, for instance, although SQS is the most common queue service, there are other queues in other services too, such as the video-processing service MediaConvert.
Reasoning. The agent must leverage that knowledge well to reason about which resources the task involves, survey the environment, and plan the operations. In our queue task example, this means that the agent must determine that other queues beyond the standard SQS queues exist (e.g., MediaConvert queues) and are in scope of the task, and plan a clean sweep.
Acting. The agent then must act accordingly and determine when the actions have succeeded. This is not easy because cloud operations often have side effects, which can be irreversible. For example, before deleting a MediaConvert queue, it must ensure that it is safe to do so.
Based on these steps, we classified the failures into three categories:
1 Knowing. The agent exhibits a knowledge gap about some services and fails because it did not know how the service behaves.
2 Reasoning. The agent knew how the service behaves but failed because of reasoning gaps. For instance, it may have surveyed only part of the environment, or seen a resource and treated it as outside the task.
3 Acting. The agent invoked the wrong tools and broke something or left the task unfinished.
The chart below shows the failure mode of each task for each model.
We found that knowledge gaps are rare: no task failed because the agent misunderstood a service. We found that every model we tested knows the cloud service behavior the tasks rely on. Thus, knowledge was never the bottleneck, and more documentation or better retrieval would likely not have changed these results.
In contrast, reasoning gaps are the most common: they accounted for about 80% of all failed tasks. For instance, we observed that agents stopped surveying before they had covered every service a DevOps engineer typically checks, and declared the task complete. In most of these attempts the agent never looked for additional resources in other services or resource types, and nothing in its output showed that anything was missing.
Wrong actions account for about 20% of failed tasks. For instance, in a bucket migration task, agents that found an in-progress multipart upload, a large file still being uploaded in parts, aborted it as leftover data. That destroyed the data, and the upload cannot be recovered. Because cloud infrastructure is critical, such mistakes have irreversible consequences for the environment.
The chart below shows the failure modes by task type.
The figure below shows a typical attempt at the queue task, from a Claude Fable 5.1 run: every command succeeded, and the agent's own re-check came back clean. Every model knew that MediaConvert can have queues of its own, yet we found that not one attempt by any model listed them.
aws sqs list-queues --queue-name-prefix vera2- 2 queues: vera2-ingest, vera2-worker aws sqs list-queues only these 2 queues in the region aws sqs list-queues --region $r (11 other regions) no queues in any other region aws sqs delete-queue (×2) both deleted aws sqs list-queues --queue-name-prefix vera2- 0 queues Agent's report “No vera2- queue remains in the account. … Eleven other regions were scanned and contained no queues at all, so nothing vera2-related exists outside us-east-1.” never run: aws mediaconvert list-queues vera2-transcode queue still exists Verifier: fail. A vera2- queue remains in the environment. In this article

## Original Extract

Frontier agents on CloudOpsBench have the knowledge cloud tasks require and fail mainly at judging when a task is done. A three-step taxonomy of where they fail.

CloudOpsBench | How Well Do Agents Perform on DevOps Tasks? Skip to content CloudOps Bench . Leaderboard About Blog Contributors GitHub Blog How Well Do Agents Perform on DevOps Tasks?
CloudOpsBench is a new benchmark that measures how well AI agents handle DevOps tasks in the cloud: repairing failures, migrating services, rotating cryptographic keys, rolling out updates, and decommissioning projects. Unlike typical software engineering (SWE) tasks, DevOps tasks are safety-critical, hard to grade with a test suite, and require a deep understanding of the environment.
CloudOpsBench v1.0 consists of 45 DevOps tasks on AWS and has two distinguishing features: (1) It is a frontier benchmark. Even the strongest agents leave many tasks unsolved: Claude Fable 5.1 completed 29 of the 45 tasks, and GPT-6 Astra completed 26. (2) It runs identically in the real AWS cloud and in a high-fidelity emulator, with the same results in both. Agents can be evaluated quickly and safely, with no real cloud account involved.
The tasks span a range of representative DevOps workflows: decommissioning projects (13 tasks), securing access (12), migrating and deploying services (7), configuring and monitoring services (6), and repairing failures (7).
In order to accomplish a DevOps task, there are three components that are important for DevOps engineers—whether a human or an agent.
Knowing. The agent must have a deep understanding of cloud services. As a simple example, if the agent is prompted to delete every queue from a project, it must know which cloud services may have queues. In AWS, for instance, although SQS is the most common queue service, there are other queues in other services too, such as the video-processing service MediaConvert.
Reasoning. The agent must leverage that knowledge well to reason about which resources the task involves, survey the environment, and plan the operations. In our queue task example, this means that the agent must determine that other queues beyond the standard SQS queues exist (e.g., MediaConvert queues) and are in scope of the task, and plan a clean sweep.
Acting. The agent then must act accordingly and determine when the actions have succeeded. This is not easy because cloud operations often have side effects, which can be irreversible. For example, before deleting a MediaConvert queue, it must ensure that it is safe to do so.
Based on these steps, we classified the failures into three categories:
1 Knowing. The agent exhibits a knowledge gap about some services and fails because it did not know how the service behaves.
2 Reasoning. The agent knew how the service behaves but failed because of reasoning gaps. For instance, it may have surveyed only part of the environment, or seen a resource and treated it as outside the task.
3 Acting. The agent invoked the wrong tools and broke something or left the task unfinished.
The chart below shows the failure mode of each task for each model.
We found that knowledge gaps are rare: no task failed because the agent misunderstood a service. We found that every model we tested knows the cloud service behavior the tasks rely on. Thus, knowledge was never the bottleneck, and more documentation or better retrieval would likely not have changed these results.
In contrast, reasoning gaps are the most common: they accounted for about 80% of all failed tasks. For instance, we observed that agents stopped surveying before they had covered every service a DevOps engineer typically checks, and declared the task complete. In most of these attempts the agent never looked for additional resources in other services or resource types, and nothing in its output showed that anything was missing.
Wrong actions account for about 20% of failed tasks. For instance, in a bucket migration task, agents that found an in-progress multipart upload, a large file still being uploaded in parts, aborted it as leftover data. That destroyed the data, and the upload cannot be recovered. Because cloud infrastructure is critical, such mistakes have irreversible consequences for the environment.
The chart below shows the failure modes by task type.
The figure below shows a typical attempt at the queue task, from a Claude Fable 5.1 run: every command succeeded, and the agent's own re-check came back clean. Every model knew that MediaConvert can have queues of its own, yet we found that not one attempt by any model listed them.
aws sqs list-queues --queue-name-prefix vera2- 2 queues: vera2-ingest, vera2-worker aws sqs list-queues only these 2 queues in the region aws sqs list-queues --region $r (11 other regions) no queues in any other region aws sqs delete-queue (×2) both deleted aws sqs list-queues --queue-name-prefix vera2- 0 queues Agent's report “No vera2- queue remains in the account. … Eleven other regions were scanned and contained no queues at all, so nothing vera2-related exists outside us-east-1.” never run: aws mediaconvert list-queues vera2-transcode queue still exists Verifier: fail. A vera2- queue remains in the environment. In this article
