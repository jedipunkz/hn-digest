---
source: "https://www.modelfuzz.com/"
hn_url: "https://news.ycombinator.com/item?id=49060845"
title: "Show HN: ModelFuzz – Open-source runtime guardrails for AI agents"
article_title: "ModelFuzz — Runtime Guardrails for AI Agents"
author: "higagan"
captured_at: "2026-07-26T18:57:00Z"
capture_tool: "hn-digest"
hn_id: 49060845
score: 2
comments: 1
posted_at: "2026-07-26T18:22:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ModelFuzz – Open-source runtime guardrails for AI agents

- HN: [49060845](https://news.ycombinator.com/item?id=49060845)
- Source: [www.modelfuzz.com](https://www.modelfuzz.com/)
- Score: 2
- Comments: 1
- Posted: 2026-07-26T18:22:24Z

## Translation

タイトル: Show HN: ModelFuzz – AI エージェント用のオープンソース ランタイム ガードレール
記事のタイトル: ModelFuzz — AI エージェントのランタイム ガードレール
説明: ModelFuzz は、プロンプト インジェクションから LLM エージェントを保護します。脆弱性をスキャンし、1 つのデコレータを使用して実行層で安全でないツールの呼び出しをブロックします。

記事本文:
モデル・ファズ
問題
仕組み
デモ
GitHub ↗
v0.2.1 · pip インストール モデルファズ
AI エージェントのランタイム ガードレール
プロンプト インジェクションによって引き起こされる安全でないツール呼び出しをインターセプトしてブロックします。実行層でデータの漏洩を阻止します。
プロンプトインジェクションは、エージェント独自のツールを敵に回します。
単一の汚染されたドキュメント、電子メール、または Web ページが間接的なプロンプト インジェクションを通じて LLM をハイジャックし、ソースでの LLM エージェントのセキュリティを侵害する可能性があります。モデルはそれが役に立っていると考えています。そうではありません。
！ shell.run — インフラストラクチャ上での任意のコマンドの実行。
！ http.post — 攻撃者のサーバーへのシークレットのサイレント漏洩。
！ send_email — API キーと顧客データが Attacker@evil.com に漏洩しました。
！即時レベルのフィルターでは安全性は保証できません。モデルの動作は非決定的です。
穴を見つけてください。それからそれらを封印します。
ModelFuzz には、セキュリティ ループの両方の半分、つまり脆弱なエージェントを暴露するレッドチーム スキャナーと、それらを保護するデコレーターが同梱されています。
欺瞞的なプロンプトインジェクションペイロードを持つ OpenAI 互換エンドポイントをレッドチームにします。どの攻撃がエージェントを騙してツールを呼び出すかを正確に確認します。
任意のツールを 1 つのデコレータでラップします。関数が実行される前に、すべての引数がポリシーに対してチェックされます。つまり、損害が発生する前に違反が発生します。
攻撃はリアルタイムで停止されました。
プロンプトによって挿入されたエージェントは、API キーを抽出しようとします。 ModelFuzz は実行層でそれを捕捉します。
チーム向けにホストされたダッシュボードが必要ですか?
一元化されたポリシー、監査ログ、継続的なエージェント スキャン。早期アクセスの待機リストに参加してください。

## Original Extract

ModelFuzz secures LLM agents against prompt injection. Scan for vulnerabilities, then block unsafe tool calls at the execution layer with one decorator.

model · fuzz
Problem
How it works
Demo
GitHub ↗
v0.2.1 · pip install modelfuzz
Runtime Guardrails for AI Agents
Intercept and block unsafe tool calls caused by prompt injection. Stop data exfiltration at the execution layer.
Prompt injection turns your agent's own tools against you.
A single poisoned document, email, or web page can hijack an LLM through indirect prompt injection, compromising LLM agent security at the source. The model thinks it's helping. It isn't.
! shell.run — arbitrary command execution on your infrastructure.
! http.post — silent exfiltration of secrets to an attacker's server.
! send_email — API keys and customer data leaked to attacker@evil.com .
! Prompt-level filters can't guarantee safety. Model behavior is non-deterministic .
Find the holes. Then seal them.
ModelFuzz ships with both halves of the security loop — a red-team scanner to expose vulnerable agents, and a decorator to shield them.
Red-team any OpenAI-compatible endpoint with deceptive prompt-injection payloads. See exactly which attacks trick your agent into calling a tool.
Wrap any tool with one decorator. Every argument is checked against your policies before the function runs — a violation raises before damage is done.
An attack, stopped in real time.
A prompt-injected agent tries to exfiltrate an API key. ModelFuzz catches it at the execution layer.
Want a hosted dashboard for your team?
Centralized policies, audit logs, and continuous agent scanning. Join the waitlist for early access.
