---
source: "https://sync.buzz"
hn_url: "https://news.ycombinator.com/item?id=48535550"
title: "Sync – Quality Control and Project Management System for AI Agents"
article_title: "Sync — Take AI agents under absolute control · Sync"
author: "nikolai_evseev"
captured_at: "2026-06-15T04:37:54Z"
capture_tool: "hn-digest"
hn_id: 48535550
score: 2
comments: 0
posted_at: "2026-06-15T01:46:28Z"
tags:
  - hacker-news
  - translated
---

# Sync – Quality Control and Project Management System for AI Agents

- HN: [48535550](https://news.ycombinator.com/item?id=48535550)
- Source: [sync.buzz](https://sync.buzz)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T01:46:28Z

## Translation

タイトル: Sync – AI エージェント向けの品質管理およびプロジェクト管理システム
記事のタイトル: Sync — AI エージェントを絶対的な制御下に置く · Sync
説明: Sync は、AI エージェントが作業するプロジェクト管理および品質管理レイヤーです。仕様、ルール、意思決定、ロードマップを定義すると、すべてのエージェントがそれらに対して責任を負い、幻覚を見たり過去の間違いを繰り返すのではなく、アーキテクチャに従います。無料、オープンソース、ローカルファースト。

記事本文:
同期 — AI エージェントを絶対的な制御下に置く · 同期 同期機能の比較 はじめる ドキュメント GitHub 無料およびプライベート — GitHub でスターを付ける AI エージェントを絶対的な制御下に置く
Sync は、AI エージェントが作業するプロジェクト管理および品質管理レイヤーです。仕様、ルール、意思決定、ロードマップを一度定義します。Sync はすべてのエージェントに責任を負わせるため、幻覚を見せたり過去の間違いを繰り返すのではなく、エージェントがアーキテクチャに基づいて構築されます。コードが移動すると、古いコンテキストにフラグが立てられるため、エージェントは停止して質問する必要があります。
ダウンロード GitHub で表示 インストールから結果までワンステップで完了します。完全にあなたのマシン上で実行され、何もアップロードされません。現在はアルファ版 v0.4.5 です。チームとクラウド機能の構築 — 待機リストに参加してください。
MCPサーバー · デスクトップアプリ · 内蔵ターミナル · クロードコード · カーソル · Zed · Codex · OpenCode
状況が悪いと、エージェントとの距離が縮まってしまいます
エージェントは、読み取ったものと同じように調整されます。大きなタスクに関する古くて散在したコンテキストにそれを向けると、あなたが見て見ぬふりをしている間に、自信を持ってコースから外れてしまいます。
コードが数か月前に移行したときでも、見つけたものはすべて信頼します。
ドキュメントが欠落している場合は、もっともらしいが間違ったコンテキストが生み出されます。
すでに覆した選択を何度も再訴訟します。
必須のターミナル コマンド、git clone、長い設定ファイルはありません。すべてはデスクトップ アプリ内で、またはワンクリックで行われます。
macOS、Windows、または Linux 用の Sync をダウンロードします。バイナリは 1 つで、依存関係はありません。ただ開くだけです。
任意のローカル リポジトリでのポイント同期。 .sync/ ナレッジ ベースが自動的にセットアップされます。手動での構成は必要ありません。
0 3 お気に入りのエージェントを接続します
ようこそウィザードは、Claude Code、Cursor、Zed、Codex、または OpenCode を即座に接続し、Sync の組み込みターミナル内で直接実行できます。

設定ファイルはありません。
最新リリース v0.4.5 をダウンロード · macOS、Windows & Linux · 無料のオープンソース
API 検索… ⌘K 2 A SW + 決定有効 アップロードにキューベースの取り込みを使用する
API を処理から切り離すので、遅いワーカーがリクエストをブロックすることはありません。インラインアプローチに取って代わります。
古いアップロードは 10 MB 未満にする必要があります
説明されているファイルは最後のコミットで変更されました。エージェントが信頼する前に、Sync がレビュー用のフラグを立てます。
エージェントにテキストを送信するだけでなく、エージェントを管理するように構築されています
AI コーディング エージェントは強力ですが、それに対抗する計画がなければ、アーキテクチャのドリフト、決定の忘れ、回帰を引き起こします。メモリ ファイルと Wiki はコンテキストを保持できますが、エージェントがプロジェクトに対して責任を負うように構築されているわけではありません。同期は、エディターとエージェントの上にあるミッションコントロール層です。
同期 プロジェクト管理と品質管理 メモリ システム 例:カーソル ルール ドキュメント Confluence · Notion 主な目標 エージェントの結果を制御し、品質を強制する
単純なテキストコンテキストをエージェントに提供する
人間が読める知識の共有
エージェントが読み取り可能な厳格な制約とウィジェット
非構造化テキスト、表、画像
自動 — コード変更と密接に結びついています
ルールを尊重しなければなりません。コンテキストが古くなると承認を求める
メモは読みますが、無視することがよくあります
通常はシームレスにアクセスできない
ビジュアル プロジェクト ハブ、ウィジェット、直接監視
Sync は非常に効果的なメモリ システムとしても機能しますが、その真の力はガバナンスと品質管理です。
あなたとエージェントが共有する 1 つの真実の情報源
構造化され、バージョン管理され、常に最新です。そのため、ドリフト後に後片付けするのではなく、あなたとエージェント (およびチーム全体) が同じものを構築します。
古い知識であることを示す知識
すべてのエントリは、それが説明するコードにバインドされます。コミットがそれを超えると、同期はそれを古いものとしてマークし、正確なコミットを指します。

d ファイル — 古いデータに基づいて黙って構築するのではなく、エージェントを停止して入力を求めます。
コードにバインドされています。それを通過するコミットには古いフラグが立てられます
有効、未検証、古い、無効 — 決して推測ではありません
古いコンテキストを信頼する前に、エージェントに一時停止して質問するよう強制します
セッションは、15 分間の有効期限を持つ署名付き JWT です。
仕様、決定、制約、観察はコーディング中にエージェントによって捕捉され、後で正確にクエリできる型付きの構造になります。散文ではなく、誰も再読する必要はありません。
理由を付けた決定
エージェントが尊重しなければならない制約
システムが実際にどのように動作するかについての観察
知識 書面による決定 · d-7af213 有効 更新のたびにリフレッシュ トークンのローテーションを使用します。
生のトークンや PII をログに記録しないでください。境界で編集します。
プロジェクト ハブ — ロードマップ、仕様、マイルストーン
エージェントが取り組むロードマップ
目標→マイルストーン→仕様、メモとしてではなく制御要素として使用されます。これは、エージェントに指示する計画です。エージェントは、独自の推測ではなく、ロードマップから次の内容を引き出し、それに基づいてコードを作成します。
優先順位に従ってランク付けされた目標、マイルストーン、仕様
エージェントは独自の計画ではなく、ロードマップに基づいてコードを作成します
ブロックされた作業はキューから自動的に削除されます
アクティブ · a1b2c3d シートベースの請求 カスタム ウィジェット API 大きな仕事の最も重要な部分
エージェントがどこに向かっているのかを常に制御する
大規模なタスクの難しい部分は、コードを入力することではなく、正しい結果を目指して作業を続けることです。同期により、エージェントの方向が読みやすく、操作可能になります。
未解決の質問に答えてください - あなたの答えは標準になります
一度決定を覆すか、制約を設定します。それは保持します
自動的にコミットされるものはありません。すべてのエージェントの書き込みを確認して承認します。
エージェント Redis のサーバー側でセッションを保持することを計画しています... サーバー側セッションまたはステートレス JWT?オープンユースタ

テレス JWT、有効期限は 15 分。サーバーセッションストアはありません。エージェント調整計画→ステートレスJWT。決定を記録する。あなたが決めるのです。それは進みます。何も自動コミットされません。
エージェントを指示した場所で実行する
エージェント用の組み込みターミナル
Sync 内で直接エージェントを起動します。複数のターミナルを並べて、それぞれがプロジェクトに開かれ、内蔵 MCP サーバーを通じてコンテキストに接続されます。個別のコンソール ウィンドウやコンテキストの切り替えはありません。
複数のエージェント セッションが並行して実行されている
それぞれがプロジェクト内で起動され、そのコンテキストに関連付けられます
内蔵 MCP サーバーを利用 - 追加のセットアップは不要
inside Sync $ git-sync mcp · 接続済み → ロードマップを読む · スコープ内の 3 つの仕様 → 制約を尊重: アップロード < 10 MB ✎ 未解決の質問を提出 → 回答待ち あなたからチーム全体へ
一人で始めてチーム全体にスケールアップ
あなたはその難しい問題に一人で取り組み始めます。ナレッジは git 内に存在するため、プッシュした瞬間にプライベート コンテキストではなくなります。チーム全体とそのエージェントが同じ新しい計画を読み、その計画に沿って操作します。
新しいチームメイトまたはエージェントは、MCP 経由で数秒でオンボーディングします
誰もが同じ計画と同じ答えに取り組んでいます
部族の知識がなく、再説明も必要ありません
あなたは有効な決定を書きました · d-3c2a4f — トークン戦略
git Push AK がそれを読み取り、MR がそれを読み取る JD がそれを読み取る 1 人の開発者がそれを書き込みます。チーム全体とそのエージェントはそれを読み、それに基づいて舵を取ります。
AI エージェントのミッションコントロール
単なるメモではありません。タスク トラッカー エージェントは実際に作業しており、プロジェクトで現在何が起こっているかを正確に示すビジュアル ダッシュボードを備えています。
マイルストーンを設定し、仕様を定義すると、エージェントは独自の推測ではなく、あなたの計画から次の内容を引き出します。ブロックされた作業はキューから削除されます。リアルタイムで進捗状況を確認できます。
あなたまたはあなたのエージェントが構築するウィジェット
ウィジェットを追加してプロジェクトの健全性を視覚化し、キューを開きます

estions、古いエントリ、または任意のメトリクス。エージェントはダッシュボードを作成することもでき、Sync を生きたオペレーション コンソールに変えることができます。
共有コンテキストを正直に保つように構築
信頼できる検証状態
すべてのエントリは有効、未検証、古い、または無効です。そのため、あなたもあなたのエージェントも、ひっそりと間違ったコンテキストに基づいて構築することはありません。
エージェントがドキュメントを作成する言語を選択します。既製のプリセットに加えて、BCP-47 タグを使用できるため、仕様、決定事項、ドキュメントがチームにとって自然に読み取れます。
型付きリンク — depend_on、supersedes、references — フロントマター、手動リンク、および本文の言及から構築されます。
ディスカッションは正確なテキストに固定されています。人間による回答は、仕様の標準的な帰属部分となります。
サイド レールからプロジェクトを切り替えることも、複数のプロジェクトを厳選されたワークスペースにグループ化して 1 つとして操作することもできます。
すべてのエントリは、それが説明するファイルにバインドされます。コードが変更されると、Sync は影響を受けるナレッジを再チェックし、ドリフトした内容にフラグを立てます。手動による監査は必要ありません。
私たちはAIが決して間違ったことをしないとは言いません
プロジェクトは変化します。ファイルが移動します。 API の名前が変更されます。チェックしないままにしておくと、エージェントは密かに古くなったコンテキストを信頼し続けますが、それは通知されません。 Sync の古いエンジンは、それを確実に行う必要があります。
古いコンテキストに自動的にフラグを立てる
仕様または決定で記述されているファイルにコミットが触れると、Sync はそのファイルを古いものとしてマークし、正確なコミットとドリフトされたファイルを示します。手動監査はありません。
エージェントに停止して尋ねるよう強制します
エージェントは、古いデータに基づいて黙って構築するのではなく、古くなったデータを表面化し、それを再検証するように依頼する必要があります。つまり、一番の幻覚経路を遮断する必要があります。
何も自動コミットされません。あなたはプロジェクト ハブでエージェントの提案された決定を確認し、質問に答えます。あなたは常に運転席に座っています。
AI を使用してコードを出荷するすべての人向け
ソロで構築している場合でも、チームを率いている場合でも、S

ync はエージェントとの関係を予測可能にし、あなたが責任を持ち続けられるようにします。
ルーチンをエージェントに委任し、アーキテクトの役割を維持します。方向性を一度設定します。エージェントは推測ではなく、あなたの計画に基づいて構築します。
エージェントに指示を繰り返すのはやめてください。一度仕様を作成し、それをコードにリンクすると、すべてのセッションで正規バージョンが取得されます (推論も含まれます)。
Notion で消えてしまう散文ではなく、AI が実際に読み取って実行する仕様を作成します。承認基準をマイルストーンにリンクします。進歩が起こるのを見てください。
ソフトウェアだけではありません。エージェントがカスタム ウィジェットを作成するため、プロジェクト ハブは構築しているものすべてに適応します。
チームが着地したときに最初になる
現在、Sync は無料でオープンソースです。次にチーム機能とクラウド機能を構築しています。メールを送っていただければ、最初にお知らせします。間にノイズはありません。
今日は無料のオープンソースです。チームとクラウドの機能がリリースされたら、最初にメールでお知らせします。
Sync は、AI エージェントが作業するプロジェクト管理および品質管理レイヤーです。あなたがプロジェクト (仕様、ルール、意思決定、ロードマップ) を定義すると、Sync はすべてのエージェントにそれに対する責任を負わせるため、エージェントは幻覚を見たり過去の間違いを繰り返したりすることなく、アーキテクチャと制約に従います。すべてはリポジトリ内に存在します。クラウドもアカウントもテレメトリもありません。
Wiki や Markdown ドキュメントとどう違うのですか?
Wiki と Markdown は静かに腐敗します。ページがコードと一致しなくなっても何も通知されず、エージェントは通常、いずれにしてもそれらにシームレスにアクセスできません。 Sync は、エージェントが読み取り可能な構造化された仕様と制約を保存し、それぞれを記述したファイルに関連付け、それらのファイルが変更されると古いというフラグを立てます。そのため、エージェントが古い情報に対してパターン マッチングを行うことはありません。
これはエージェントのための単なる記憶システムではないでしょうか?
同期を非常に効率的なメモリ システムとして使用できます。

テム — しかし、それが重要ではありません。記憶ツールは生のテキストをエージェントに渡し、それが依然として真実であることを望みます。 Sync の仕事はガバナンスです。あらゆる仕様、決定、制約を記述されたコードに結び付け、エージェントにそれらを尊重させ、コードが変更されると古いというフラグを立て、古いデータを処理する前に強制的に停止して承認を求めます。力となるのはリコールだけではなく、コントロールと品質です。
エージェントを制御できるようにするにはどうすればよいですか?
エージェントが取り組む計画、つまり目標、マイルストーン、優先順位に基づいてランク付けされた仕様を設定します。動作中、エージェントは制約を読み取り、ロードマップを更新し、曖昧な点がある場合には未解決の質問をファイルします。あなたの答えは標準的になります。決定を覆し、尊重しなければならない制約を設定します。自動的にコミットされるものはありません。すべての変更を確認して承認します。方向を変えるのはあなたです。エージェントが続きます。
リポジトリ内。 Sync はプレーン ファイルを `.sync/` の下に書き込みます。SaaS バックエンドもアカウントもなく、何もアップロードされません。知識はバージョン管理され、プル リクエストでレビュー可能になり、コードとともに伝達されます。
どのエージェントとエディターが同期と連携しますか?
MCP 対応のクライアント。ようこそウィザードは、Claude Code、Cursor、Zed、OpenCode、Codex を即座に接続し、Sync の内蔵ターミナル内で直接実行できます。それ以外の場合は、「claude mcp add git-sync」を使用してクライアントにバイナリを指定します。

[切り捨てられた]

## Original Extract

Sync is the project-management and quality-control layer your AI agents work against. Define your specs, rules, decisions, and roadmap — and every agent stays accountable to them, following your architecture instead of hallucinating or repeating past mistakes. Free, open source, local-first.

Sync — Take AI agents under absolute control · Sync Sync Features Compare Get started Docs GitHub Free & private — star it on GitHub Take AI agents under absolute control
Sync is the project-management and quality-control layer your AI agents work against. Define your specs, rules, decisions, and roadmap once — Sync holds every agent accountable to them, so they build against your architecture instead of hallucinating or repeating past mistakes. When the code moves, stale context gets flagged and the agent has to stop and ask.
Download View on GitHub From install to results in one step. Runs entirely on your machine — nothing is ever uploaded. Currently in alpha · v0.4.5 . Building team & cloud features — join the waitlist .
MCP server · Desktop app · Built-in terminal · Claude Code · Cursor · Zed · Codex · OpenCode
Bad context is how you and your agent drift apart
An agent is only as aligned as what it reads. Point it at stale, scattered context on a large task and it heads off course — confidently, while you're looking the other way.
It trusts whatever it finds, even when the code moved on months ago.
Where docs are missing, it invents plausible-but-wrong context.
It re-litigates choices you already overturned, again and again.
No mandatory terminal commands, no git clone, no long configuration files. Everything happens in the desktop app or with a single click.
Download Sync for macOS, Windows, or Linux. One binary, no dependencies — just open it.
Point Sync at any local repository. It sets up the .sync/ knowledge base automatically — no manual config.
0 3 Connect your favorite agent
The Welcome wizard wires up Claude Code, Cursor, Zed, Codex, or OpenCode instantly — and you can run them right inside Sync's built-in terminal. No config files.
Download Latest release v0.4.5 · macOS, Windows & Linux · free and open source
Api Search… ⌘K 2 A S W + Decision valid Use queue-based ingestion for uploads
Decouples the API from processing so a slow worker never blocks a request. Supersedes the inline approach.
stale Uploads must stay under 10 MB
The file it describes changed in the last commit — Sync flags it for review before your agent trusts it.
Built to govern agents, not just feed them text
AI coding agents are powerful, but without a plan to work against they cause architectural drift, forgotten decisions, and regressions. Memory files and wikis can hold context — they aren't built to keep an agent accountable to your project. Sync is: a mission-control layer above your editor and agent.
Sync Project management & quality control Memory systems e.g. Cursor rules Documentation Confluence · Notion Primary goal Control agent results & enforce quality
Provide simple text context to the agent
Human-readable knowledge sharing
Strict, agent-readable constraints & widgets
Unstructured text, tables, images
Automatic — tied tightly to code changes
Must respect the rules; asks for approval when context goes stale
Reads notes, but often ignores them
Usually can't access it seamlessly
Visual Project Hub, widgets, direct oversight
Sync works as a highly effective memory system too — but its real power is governance and quality control.
One source of truth you and your agent share
Structured, versioned, and always current — so you and your agent (and the whole team) build the same thing, instead of you cleaning up after it drifts.
Knowledge that flags itself stale
Every entry is bound to the code it describes. When a commit moves past it, Sync marks it stale — pointing at the exact commit and files — and makes the agent stop and ask for your input instead of silently building on old data.
Bound to the code; flagged stale on the commit that moves past it
valid · unverified · stale · invalid — never a guess
Forces the agent to pause and ask before trusting stale context
Sessions are signed JWTs with a 15-minute expiry.
Specs, decisions, constraints, and observations get captured during coding, by the agent, in typed structure it can query precisely later — not prose nobody re-reads.
Decisions with the reasoning attached
Constraints the agent must respect
Observations about how the system really behaves
knowledge written decision · d-7af213 valid Use refresh-token rotation on every renew.
Never log raw tokens or PII — redact at the boundary.
Project Hub — roadmap, specs, milestones
A roadmap your agent works against
Goals → milestones → specs, used not as notes but as control elements. This is the plan you point the agent at — it pulls what's next from your roadmap and codes against it, instead of its own guess.
Goals, milestones, and specs ranked by priority
The agent codes against the roadmap, not its own plan
Blocked work drops out of the queue automatically
active · a1b2c3d Seat-based billing Custom widget API The most important part of big work
Stay in control of where the agent is heading
The hard part of a large task isn't typing the code — it's keeping the work pointed at the right outcome. Sync makes the agent's direction legible and steerable.
Answer its open questions — your answers become canonical
Overturn a decision or set a constraint once; it holds
Nothing auto-commits — you review and approve every agent write
agent Planning to keep sessions server-side in Redis… Server-side sessions or stateless JWT? open you Stateless JWT, 15-min expiry. No server session store. agent Adjusting plan → stateless JWT. Recording the decision. You decide; it proceeds. Nothing auto-commits.
Run agents where you steer them
A built-in terminal for your agents
Launch your agents right inside Sync — multiple terminals side by side, each opened into your project and wired to its context through the built-in MCP server. No separate console window, no context switching.
Multiple agent sessions running in parallel
Each launched into your project, wired to its context
Powered by the built-in MCP server — no extra setup
inside Sync $ git-sync mcp · connected → reading roadmap · 3 specs in scope → respecting constraint: uploads < 10 MB ✎ filed open question → awaiting your answer From you to the whole team
Start solo, scale to the whole team
You start alone on the hard problem. Because the knowledge lives in git, it stops being your private context the moment you push — the whole team and their agents read and steer with the same fresh plan.
A new teammate or agent onboards in seconds via MCP
Everyone works the same plan and the same answers
No tribal knowledge, no re-explaining
You wrote valid decision · d-3c2a4f — Token strategy
git push AK reads it MR reads it JD reads it One developer writes it. The whole team — and their agents — read and steer with it.
Your mission control for AI agents
Not just notes — a task tracker agents actually work against, with a visual dashboard that shows exactly what's happening in your project right now.
Set milestones, define specs, and the agent pulls what's next from your plan — not its own guess. Blocked work drops out of the queue. You see progress in real time.
Widgets you or your agent build
Add widgets to visualize project health, open questions, stale entries, or any metric. Your agent can create dashboards too — turning Sync into a living operations console.
Built to keep your shared context honest
A validation state you can trust
Every entry is valid, unverified, stale, or invalid — so neither you nor your agent ever builds on context that quietly went wrong.
Choose the language your agent writes docs in. Ready-made presets, plus any BCP-47 tag — so specs, decisions, and docs read naturally for your team.
Typed links — depends_on, supersedes, references — built from frontmatter, manual links, and body mentions.
Discussion pinned to exact text. A human answer becomes the canonical, attributed part of the spec.
Switch between projects from the side rail, or group several into a curated workspace and steer them as one.
Every entry is bound to the files it describes. When the code changes, Sync re-checks the affected knowledge and flags what drifted — no manual audits.
We don't pretend AI never gets it wrong
Projects change. Files move. APIs get renamed. Left unchecked, your agent keeps trusting context that quietly went stale — and it won't tell you. Sync's staleness engine ensures it has to.
Flags stale context automatically
When a commit touches files a spec or decision describes, Sync marks it stale — pointing at the exact commit and drifted files. No manual audit.
Forces the agent to stop and ask
Instead of silently building on old data, the agent has to surface what went stale and ask you to re-validate it — shutting down the #1 hallucination path.
Nothing is auto-committed. You review the agent's proposed decisions and answer its questions in the Project Hub — you're always in the driver's seat.
For anyone who ships code with AI
Whether you're building solo or leading a team, Sync makes your relationship with agents predictable — and keeps you in charge.
Delegate routine to agents, keep the architect role. Set the direction once — the agent builds against your plan, not its guess.
Stop repeating instructions to your agent. Write a spec once, link it to the code, and every session gets the canonical version — reasoning included.
Write specs the AI actually reads and executes — not prose that dies in Notion. Link acceptance criteria to milestones. Watch progress happen.
And not just software. Because your agent authors custom widgets, the Project Hub adapts to whatever you're building:
Be first when the team features land
Sync is free and open source today. We're building team and cloud features next — drop your email and we'll tell you first. No noise in between.
Free and open source today. We'll email you first when team and cloud features land.
Sync is the project-management and quality-control layer your AI agents work against. You define the project — its specs, rules, decisions, and roadmap — and Sync holds every agent accountable to it, so they follow your architecture and constraints instead of hallucinating or repeating past mistakes. Everything lives in your repository; there is no cloud, no account, and no telemetry.
How is it different from a wiki or Markdown docs?
Wikis and Markdown rot silently: nothing tells you when a page no longer matches the code, and your agent usually can't access them seamlessly anyway. Sync stores structured, agent-readable specs and constraints, ties each one to the files it describes, and flags it stale when those files change — so the agent never pattern-matches against out-of-date information.
Isn't this just a memory system for agents?
You can use Sync as a highly effective memory system — but that's not the point. Memory tools pass raw text to the agent and hope it's still true. Sync's job is governance: it ties every spec, decision, and constraint to the code it describes, makes the agent respect them, flags them stale when the code changes, and forces it to stop and ask for your approval before acting on old data. The power is control and quality, not just recall.
How does it keep me in control of the agent?
You set the plan the agent works against — goals, milestones, and specs ranked by priority. As it works, the agent reads your constraints, updates the roadmap, and files open questions when it hits ambiguity; your answers become canonical. You overturn decisions and set constraints it must respect. Nothing is auto-committed — you review and approve every change. You steer the direction; the agent follows.
In your repository. Sync writes plain files under `.sync/` — there is no SaaS backend, no account, and nothing is uploaded. Your knowledge is versioned, reviewable in pull requests, and travels with the code.
Which agents and editors work with Sync?
Any MCP-capable client. The Welcome wizard wires up Claude Code, Cursor, Zed, OpenCode, and Codex for you instantly — and you can run them right inside Sync's built-in terminal. For anything else, point your client at the binary with `claude mcp add git-sync

[truncated]
