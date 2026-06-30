---
source: "https://www.framein.dev/"
hn_url: "https://news.ycombinator.com/item?id=48731852"
title: "Show HN: Framein – a local work-state layer that keeps AI agents in context"
article_title: "Framein - One work frame across Claude, Codex, and Gemini"
author: "BonPPa"
captured_at: "2026-06-30T13:35:01Z"
capture_tool: "hn-digest"
hn_id: 48731852
score: 1
comments: 0
posted_at: "2026-06-30T12:36:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Framein – a local work-state layer that keeps AI agents in context

- HN: [48731852](https://news.ycombinator.com/item?id=48731852)
- Source: [www.framein.dev](https://www.framein.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T12:36:49Z

## Translation

タイトル: Show HN: Framein – AI エージェントをコンテキスト内に保持するローカル作業状態レイヤー
記事のタイトル: Framein - クロード、コーデックス、ジェミニにわたる 1 つのワーク フレーム
説明: Framein は、AI コーディング エージェント ハーネスの下にあるローカルのワークステート レイヤーです: タスク コントラクト、チャレンジ ループ、モデルスイッチ カプセル、および検証ゲート。

記事本文:
GitHub
KO
AIコーディングのためのローカルワークフレーム
クロード、コーデックス、ジェミニにわたって 1 つの作業フレームを維持します。
1 人のエージェントで開始し、別のエージェントでチャレンジし、必要に応じて切り替えて、検証して作業を終了します。 Framein は、共有タスク コントラクト、意思決定証跡、リスク状態、検証結果、およびモデル切り替えカプセルを、すでに使用しているエージェント ハーネスの下に保持します。
# すでに使用しているエージェントで開始します。
$ Framein start "Google OAuth を追加"
コントラクト セット: 電子メール ログインを保持
リード・クロード
# 制限付きの反対意見については別のモデルに要求します。
$ Framein チャレンジ「セッション中の OAuth コールバック状態」 --run
査読者コーデックス
評決への挑戦
nonce/状態の検証を追加する必要があります
リードは必要な変更を受け入れます
# 必要に応じて切り替えます。検証して閉じます。
$フレーミンカプセルジェミニ
事実から作成された次の手がかり:
契約、差分、テスト、決定
$フレームミンシップ
ビルドはOK、テストは合格
リスク高: 認証/タッチ済み
=> ヒューマンゲートの準備完了
30秒のデモ
コンテキストを失うことなく、挑戦し、検証し、引き継ぎます。
タスク契約、独立したモデルチャレンジ、検証証拠、および別のエージェントが続行できるカプセルを示す支払いタスクのウォークスルー。
より良いプロンプトが役に立ちます。作品を完全な状態に保つことはできません。
優れた PRD、計画、ADR、およびスキルは、どのモデルでもより良い作業を行うのに役立ちます。痛みは、リードモデルが失速したり、別のモデルが計画に挑戦したり、割り当てやモデルの適合に切り替えを迫られたり、最終的な答えがまだ実際の検証を必要としているときに始まります。
Framein は、新しい IDE や別のエージェント ハーネスではありません。 Claude Code、Codex、Gemini、Pi、OpenCode、およびすでに使用しているスキル ワークフローの下に 1 つのローカル ワーク フレームが保持されます。
ハーネスがエンジンである場合、Framein はエンジンが書き込む共有ログブックです。
開発者メモ: Framein を構築した理由
開始します。チャレンジ。スイッチ。検証します。
Framein start はリクエストを int に変換します

o 実装が変更される前の共有タスク契約。
Framein チャレンジでは、別の査読者に構造化された評決、1 つの主要な回答、および決定概要を求めます。
Framein Capsule は、契約、差分、検証、ADR、台帳といったローカルな事実から次の手がかりを準備します。
Framein は、確定的なビルド/テスト チェックとリスク ゲートを使用して、ループを閉じて検証して出荷します。
それをスキルのように呼んでください。プロジェクトの状態と同様に保存します。
/fr:* は、エージェント セッションから同じローカル エンジンを呼び出します。
$fr-* は、非推奨のプロンプト ファイルなしで同じ動詞を公開します。
プロンプトフレームワークを維持してください。 Framein は、その下にある共有契約書、台帳、ゲートを提供します。
CLI、JSON 出力、ラッパー、MCP サーバーはすべて、同じローカル ワーク フレームの読み取りと書き込みを行います。
一度インストールしてください。フレームをエージェントの下に置いてください。
npmからインストールし、実際のプロジェクト内でFrameinを初期化します。日常的な使用では、エージェントは /fr:* または $fr-* を通じて同じワークフレーム動詞を呼び出すことができます。
npm パスは、Windows、macOS、Linux、および Node 22.5 以降の WSL で動作します。スタンドアロン実行可能ファイルは、Node と Framein をバンドルし、Windows npm shim の摩擦を回避するオプションの便利なパスとして計画されています。
npm install -g フレームイン
Framein --version
プロジェクトの cd
フレーム初期化
フレームイン統合はすべてインストールします --write
Framein start 「最小の安全な変更を完了する」
# 別のモデルをレビューまたは続行する必要がある場合
Framein チャレンジ「実装前に計画を確認する」 --run
フレーミンカプセルコーデックス
フレームライン検証
船のフレーム
よくある質問
ハーネスではありません。トークンリレーではありません。
いいえ、Claude Code、Codex、Gemini、Pi、OpenCode が作業面として残ります。 Framein は、その下のリポジトリに作業状態を保持します。
プロバイダー トークンに触れますか?
いいえ。Framein は資格情報を収集せず、モデル トラフィックをプロキシし、サブスクリプションをプールしません。プロバイダー認証は公式 CLI のままです。
よー

リポジトリ: Git 対応の JSON スナップショットとローカル SQLite キャッシュ。 FAQ の全文をお読みください。
まずは地元。資格情報リレーはありません。
タスク コントラクト、ADR、メモリ、台帳、検証結果、書き込みロックはリポジトリ内に存在します。 SQLite ストアはキャッシュです。 JSON スナップショットは git フレンドリーです。
Framein は、プロバイダーの資格情報、プールのサブスクリプション、リレー MCP ツール、または画面スクレイピングの端末 I/O (TTY) を収集しません。ユーザーが要求すると、公式 CLI をローカルで呼び出します。

## Original Extract

Framein is the local work-state layer beneath your AI coding agent harness: task contract, challenge loop, model-switch capsule, and validation gate.

GitHub
KO
Local work frame for AI coding
Keep one work frame across Claude, Codex, and Gemini.
Start with one agent, challenge it with another, switch when needed, and close the work with validation. Framein keeps a shared task contract, decision trail, risk state, validation results, and model-switch capsule beneath the agent harness you already use.
# Start in the agent you already use.
$ framein start "add Google OAuth"
contract set : preserve email login
lead claude
# Ask a different model for a bounded objection.
$ framein challenge "OAuth callback state in session" --run
reviewer codex
verdict challenge
required add nonce/state validation
lead accepts required change
# Switch when needed; close with validation.
$ framein capsule gemini
next lead prepared from facts:
contract · diff · tests · decisions
$ framein ship
build ok · tests passed
risk high: auth/ touched
=> READY WITH HUMAN GATE
30-second demo
Challenge, verify, and hand off without losing context.
A payment-task walkthrough showing a task contract, independent model challenge, validation evidence, and a capsule another agent can continue from.
Better prompts help. They do not keep the work intact.
Good PRDs, plans, ADRs, and skills help any model do better work. The pain starts when the lead model stalls, another model should challenge the plan, quota or model fit pushes you to switch, or the final answer still needs real validation.
Framein is not a new IDE or another agent harness. It keeps one local work frame under Claude Code, Codex, Gemini, Pi, OpenCode, and the skill workflows you already use.
If harnesses are the engine, Framein is the shared logbook the engines write to.
Developer note: why I built Framein
Start. Challenge. Switch. Validate.
framein start turns the request into a shared task contract before the implementation drifts.
framein challenge asks a different reviewer for a structured verdict, one lead response, and a decision brief.
framein capsule prepares the next lead from local facts: contract, diff, validation, ADRs, and ledger.
framein verify and ship close the loop with deterministic build/test checks and risk gates.
Call it like a skill. Store it like project state.
/fr:* calls the same local engine from the agent session.
$fr-* exposes the same verbs without deprecated prompt files.
Keep your prompt framework. Framein supplies the shared contract, ledger, and gates underneath.
The CLI, JSON output, wrappers, and MCP server all read and write the same local work frame.
Install once. Keep the frame under your agents.
Install from npm, then initialize Framein inside a real project. In day-to-day use, agents can call the same work-frame verbs through /fr:* or $fr-* .
The npm path works across Windows, macOS, Linux, and WSL with Node 22.5+. Standalone executables are planned as an optional convenience path that bundles Node with Framein and avoids Windows npm shim friction.
npm install -g framein
framein --version
cd your-project
framein init
framein integrations install all --write
framein start "complete the smallest safe change"
# when another model should review or continue
framein challenge "review the plan before implementation" --run
framein capsule codex
framein verify
framein ship
FAQ
Not a harness. Not a token relay.
No. Claude Code, Codex, Gemini, Pi, and OpenCode remain the working surface. Framein keeps the work state in your repo underneath them.
Does it touch provider tokens?
No. Framein collects no credentials, proxies no model traffic, and pools no subscriptions. Provider auth stays with the official CLI.
In your repo: a git-friendly JSON snapshot plus a local SQLite cache. Read the full FAQ .
Local first. No credential relay.
Task contract, ADRs, memory, ledger, validation results, and write locks live in your repo. The SQLite store is a cache; the JSON snapshot is git-friendly.
Framein does not collect provider credentials, pool subscriptions, relay MCP tools, or screen-scrape terminal I/O (TTY). It calls official CLIs locally when you ask it to.
