---
source: "https://guibai.dev/a/7656050265522913280/"
hn_url: "https://news.ycombinator.com/item?id=48775745"
title: "Stop AI from Wrecking Your Codebase with Spec-Driven Development"
article_title: "Stop AI from Wrecking Your Codebase with Spec-Driven Development — Guibai"
author: "Soarez"
captured_at: "2026-07-03T15:50:46Z"
capture_tool: "hn-digest"
hn_id: 48775745
score: 4
comments: 0
posted_at: "2026-07-03T14:49:23Z"
tags:
  - hacker-news
  - translated
---

# Stop AI from Wrecking Your Codebase with Spec-Driven Development

- HN: [48775745](https://news.ycombinator.com/item?id=48775745)
- Source: [guibai.dev](https://guibai.dev/a/7656050265522913280/)
- Score: 4
- Comments: 0
- Posted: 2026-07-03T14:49:23Z

## Translation

タイトル: 仕様駆動開発で AI がコードベースを破壊するのを阻止する
記事のタイトル: 仕様駆動開発で AI がコードベースを破壊しないようにする — Guibai
説明: 仕様駆動開発 (SDD) は、アドホック AI コーディングを構造化されたワークフローに置き換えます。このワークフローでは、コードの 1 行が書き込まれる前に、仕様、設計、およびタスクのドキュメントによって境界が定義されます。 OpenSpec などのツールは、階層化されたスキル、AGENTS.md、およびフックと組み合わせることで、ソフトな慣習をハードな制約に変えます。
[切り捨てられた]

記事本文:
跪拜 Guibai ☾ ← すべての記事 Agent Stop AI from Wrecking Your Codebase with Spec-Driven Development
構造化された仕様がなければ、AI コーディング ツールは不一致を増幅させます。要件を幻覚化し、変更範囲を拡大し、セッション間のルールを忘れてしまいます。 SDD は人間の役割を 1 行ずつレビューする者からルールを定義する者に移し、手戻りを減らし、複数のスプリント、複数の開発者が関わるプロジェクトに十分な予測可能な AI 出力を実現します。
AI コーディング アシスタントは、体系化されたコンテキストを欠いているため、不安定な出力を生成します。要件を推測し、チームの慣例を無視し、セッション間でルールを忘れてしまいます。仕様駆動開発 (SDD) は、作業をフロントローディングすることでこの問題を解決します。仕様は何を構築するか、何を触らないかを定義し、設計ドキュメントは技術的アプローチを固定し、タスク ファイルは作業を検証可能な小さな塊に分割します。その後、AI は即興で実行するのではなく、これらのガードレール内で実行します。
OpenSpec は、これらのドキュメントを生成、適用、同期、アーカイブする「/opsx」コマンド セットを備えた軽量のオンランプを提供します。 MCP サーバーは、Figma、Yapi、および内部 Wiki から直接コンテキストを取得するため、最初の仕様ドラフトに実際のビジネス データが不足することはありません。コード生成後、3 層の制約システム (コーディング習慣用のスキル、プロジェクト固有のルール用の AGENTS.md、自動適用用のフック) が、i18n を使用する必要があるハードコーディングされた文字列などの残りのドリフトを捕捉します。
完了した変更をアーカイブすると、1 回限りのセッションが成長するプロジェクトのナレッジ ベースに変わります。次の AI セッションでは、蓄積された境界、設計上の決定、既知の落とし穴が継承されるため、限界コラボレーション コストは時間の経過とともに低下します。経験則: 2 日を超えると予想されるタスクは、仕様を作成するための初期費用が正当化されます。
中心的な変化はツールではなく、人間を動かすことです

コードレビュー担当者から仕様作成者まで。 AI は協力者ではなく実行者になります。
プロンプトとドキュメントのソフト制約は、長時間のセッションでは失敗します。ルールに準拠しない出力をブロックするフックのみがルールを適用します。
アーカイブは最も過小評価されているステップです。これがなければ、新しいセッションはすべてゼロから始まり、同じ間違いが繰り返されます。
MCP 統合により、SDD が大規模に実用的になります。MCP 統合がなければ、分散したツールからコンテキストを手動で組み立てることになり、効率の向上が損なわれてしまいます。
2 日間のヒューリスティックは有用なリトマス試験紙です。そのしきい値を下回ると、仕様のオーバーヘッドが防止できる再作業を超えます。

## Original Extract

Spec-Driven Development (SDD) replaces ad-hoc AI coding with a structured workflow where Spec, Design, and Task documents define boundaries before a single line of code is written. Tools like OpenSpec, combined with layered Skills, AGENTS.md, and Hooks, turn soft conventions into hard constraints th
[truncated]

跪拜 Guibai ☾ ← All articles Agent Stop AI from Wrecking Your Codebase with Spec-Driven Development
Without structured specs, AI coding tools amplify inconsistency—they hallucinate requirements, expand change scopes, and forget rules between sessions. SDD shifts the human role from line-by-line reviewer to rule definer, cutting rework and making AI output predictable enough for multi-sprint, multi-developer projects.
AI coding assistants produce unstable output because they lack organized context—they guess requirements, ignore team conventions, and forget rules across sessions. Spec-Driven Development (SDD) fixes this by front-loading the work: a Spec defines what to build and what not to touch, a Design doc locks in the technical approach, and a Tasks file breaks the work into small, verifiable chunks. The AI then executes within these guardrails rather than improvising.
OpenSpec provides a lightweight on-ramp with a `/opsx` command set that generates, applies, syncs, and archives these documents. MCP servers pull context directly from Figma, Yapi, and internal wikis so the initial Spec draft isn't starved of real business data. After code generation, a three-layer constraint system—Skills for coding habits, AGENTS.md for project-specific rules, and Hooks for automated enforcement—catches the remaining drift, such as hardcoded strings that should use i18n.
Archiving completed changes turns one-off sessions into a growing project knowledge base. The next AI session inherits accumulated boundaries, design decisions, and known pitfalls, so marginal collaboration cost drops over time. The rule of thumb: any task expected to span more than two days justifies the upfront cost of writing a Spec.
The core shift isn't tooling—it's moving the human from code reviewer to spec author. The AI becomes an executor, not a collaborator.
Soft constraints in prompts and docs fail under long sessions; only Hooks that block non-compliant output make rules stick.
Archive is the most undervalued step. Without it, every new session starts from zero, and the same mistakes repeat.
MCP integration is what makes SDD practical at scale—without it, manually assembling context from scattered tools kills the efficiency gain.
The two-day heuristic is a useful litmus test: below that threshold, the spec overhead exceeds the rework it prevents.
