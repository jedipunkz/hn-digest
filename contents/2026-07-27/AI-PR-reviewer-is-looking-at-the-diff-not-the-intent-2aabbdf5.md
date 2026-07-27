---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49068241"
title: "AI PR reviewer is looking at the diff, not the intent"
article_title: ""
author: "axsdrizz"
captured_at: "2026-07-27T12:52:10Z"
capture_tool: "hn-digest"
hn_id: 49068241
score: 2
comments: 1
posted_at: "2026-07-27T11:51:40Z"
tags:
  - hacker-news
  - translated
---

# AI PR reviewer is looking at the diff, not the intent

- HN: [49068241](https://news.ycombinator.com/item?id=49068241)
- Score: 2
- Comments: 1
- Posted: 2026-07-27T11:51:40Z

## Translation

タイトル: AI PR レビュー担当者は意図ではなく差分を見ている
HN テキスト: AI コード レビュー ツールは、単一ファイルまたはプル リクエスト内の構文の分析に優れています。彼らは、1マイル離れたところで、欠けているセミコロンを見つけることができます。しかし、彼らは急いでいる人間のレビュー担当者と同じ盲目に悩まされています。彼らは意図ではなく差分をレビューします。彼らは、コーディング エージェントが使用した元のプロンプトを無視し、アプリケーションのより広範なビジネス コンテキストを把握できません。コード レビューは、人間が実施するか LLM が実施するかにかかわらず、コードが要求された内容を実行しているかどうかを実際には検証しません。コードが意図した仕様と一致しているかどうか、またはエージェントの「タスクが完了した」という主張が決定論的な方法でチェックされているかどうかを確認することはできません。 LLM は、ハードなタスク、複数のステップからなるタスク、または依存関係のあるタスクを延期することで知られています。これは、PR が開かれるまでに意図自体が分岐していることを意味します。本番環境で AI コードを信頼したい場合は、差分の検証をやめて、クレームの検証を開始する必要があります。元の Jira チケット/GitHub の問題を CI の実際のコード変更にバインドする信頼できる方法を見つけた人はいますか?私は Shipmoor.dev とこの問題の解決に取り組んできましたが、他のチームが「意図のギャップ」にどのように取り組んでいるのか聞いてみたいです。

## Original Extract

AI code review tools excel at analyzing syntax within a single file or a pull request. They can catch a missing semicolon a mile away. But they suffer from the same blindness as a rushed human reviewer: they review the diff, not the intent. They overlook the original prompt used by the coding agent and fail to grasp the broader business context of the application. Code review - whether conducted by a human or an LLM - doesn't actually verify whether the code accomplishes what was prompted. It cannot confirm that the code aligns with its intended spec, or that the agent's claim of "task completed" is checked in a deterministic way. LLMs notoriously defer hard, multi-step, or dependent tasks, which means that the intent itself has diverged by the time the PR is opened. If we want to trust AI code in production, we have to stop verifying the diff and start verifying the claim. Has anyone found a reliable way to bind the original Jira ticket / GitHub issue to the actual code change in CI? I've been working on solving this problem with Shipmoor.dev, but I'd love to hear how other teams are tackling the "intent gap."

