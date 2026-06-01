---
source: "https://agentdeck.site/"
hn_url: "https://news.ycombinator.com/item?id=48344519"
title: "Show HN: Agent Deck: Native Mac app for managing AI coding agents| powered by PI"
article_title: "Agent Deck for macOS"
author: "streetcoder"
captured_at: "2026-06-01T00:33:32Z"
capture_tool: "hn-digest"
hn_id: 48344519
score: 2
comments: 0
posted_at: "2026-05-31T10:23:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agent Deck: Native Mac app for managing AI coding agents| powered by PI

- HN: [48344519](https://news.ycombinator.com/item?id=48344519)
- Source: [agentdeck.site](https://agentdeck.site/)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T10:23:44Z

## Translation

タイトル: Show HN: Agent Deck: AI コーディング エージェントを管理するためのネイティブ Mac アプリ| PIを活用
記事のタイトル: macOS 用エージェントデッキ
説明: Agent Deck は、AI コーディング エージェントに関するすべてを管理するためのネイティブ Mac アプリです。プロジェクトごとに専門エージェントを作成し、彼らが使用するスキルをキュレーションし、GitHub の問題からセッションを実行するすべてを 1 か所で行います。

記事本文:
エージェントデッキ
に建てられました
CLI。
AI コーディング エージェント用のネイティブ Mac アプリ。
プロジェクトごとに専門エージェントを構築し、彼らが使用するスキルを厳選し、セッションを開始します
GitHub の問題から取得し、並列ワークツリーで実行します。
macOS 用のダウンロード
GitHub で見る
無料＆オープンソース
計画、差分、承認を含む色分けされたトランスクリプト
任意のリポジトリまたは URL からスキルを厳選
カスタム プロンプトとツールを使用してプロジェクトごとにエージェントを作成する
分離されたワークツリー内の並列セッションをワンクリックで結合
GitHub の問題が準備完了セッションとして起動される
各セッションから学習するプロジェクトごとのメモリ
実際に使えるMacアプリ。
プロジェクト、セッション、GitHub の問題は、1 つの端末ではなく、1 クリックの距離に留まります。エディターの隣に表示されるウィンドウに、各セッションの動作がリアルタイムで表示されます。
計画、ツール呼び出し、インライン差分を含む色分けされたストリーミング トランスクリプト
マルチウィンドウ、アイドルパーキング、ターミナルハンドオフを備えたプロジェクトおよびセッションのサイドバー
Pi CLI 設定を診断する内蔵 Doctor
エージェントが使用するための実際のライブラリ。
GitHub リポジトリまたは skill.sh URL から個々のスキルを厳選します。パッケージ全体をクローンする必要はありません。アップストリームの同期を維持します。プロジェクトごとにどのプロンプト、スキル、およびスラッシュ コマンドをアクティブにするかを決定します。
GitHub または skill.sh URL から個々のスキルをインポートする
何がどこでアクティブになっているかを示す、ビルトイン、グローバル、ライブラリ、およびプロジェクトのスコープ
AI が生成した概要により、各スキルが何を行うかを有効にする前に把握できます。
エージェントを構築し、並行して実行します。
エージェントは、独自のプロンプト、ツール、スキル、モデル、アバターを備えた第一級のオブジェクトです。分離されたワークツリーでそれらを並べて実行し、完了した作業をワンクリックでマージして戻します。
エージェントごとのカスタム プロンプト、ツール、スキル、モデル、アバター
セッションごとの Git ブランチとワークツリーの分離
完了したワークツリーをワンクリックでマージして元に戻す

ソースブランチ
問題は出発点であり、回り道ではありません。
ネイティブ ボード内のリポジトリの問題を参照して 1 つ選択し、コンテキストとして既に読み込まれているタイトル、説明、ラベル、コメントを使用してセッションを開始します。トリアージと実行は接続されたままになります。
アプリ内のネイティブ GitHub 問題ボード
セッションに自動ロードされる問題のコンテキスト
サブ課題の進捗状況を実行と同時に追跡
鍵はご自身でご用意ください。 Pi が認証を処理します。
追加のサブスクリプションはありません。 Pi を介して既存の Claude、OpenAI、またはローカル モデルのセットアップを使用します。 Agent Deck には、セッションごとのピッカー、オンデバイスのドラフト、およびランタイム ツールが追加されます。
すべての Pi プロバイダーにわたるセッションごとのモデル ピッカー
タイトル、コミットメッセージ、概要のオンデバイス基盤モデル
OpenAI 高速モードと Exa Web 検索が組み込まれています
無料のオープンソースで、macOS 26 以降向けに構築されています。

## Original Extract

Agent Deck is a native Mac app for managing everything around AI coding agents. Author specialist agents per project, curate the skills they use, and run sessions from GitHub issues, all in one place.

Agent Deck
Built on the
CLI.
A native Mac app for AI coding agents.
Build specialist agents per project, curate the skills they use, launch sessions
from GitHub issues, and run them in parallel worktrees.
Download for macOS
View on GitHub
Free & open source
Color-coded transcript with plan, diffs, and approvals
Cherry-pick skills from any repo or URL
Author per-project agents with custom prompts and tools
Parallel sessions in isolated worktrees, merged back in one click
GitHub issues launch as ready sessions
Per-project memory that learns from each session
A Mac app you can actually use.
Projects, sessions, and GitHub issues stay one click apart, not one terminal apart. You see what every session is doing, in real time, in a window that lives next to your editor.
Color-coded streaming transcript with plan, tool calls, and inline diffs
Project and session sidebar with multi-window, idle-parking, and terminal handoff
Built-in Doctor that diagnoses your Pi CLI setup
A real library for what your agents use.
Cherry-pick individual skills from any GitHub repo or skills.sh URL. No need to clone the whole package. Keep upstream in sync. Decide which prompts, skills, and slash commands are active per project.
Import individual skills from GitHub or skills.sh URLs
Builtin, Global, Library, and Project scopes for what is active where
AI-generated summaries so you know what each skill does before enabling it
Build agents, run them in parallel.
Agents are first-class objects with their own prompt, tools, skills, model, and avatar. Run them side by side in isolated worktrees, then merge completed work back with one click.
Custom prompt, tools, skills, model, and avatar per agent
Per-session git branch and worktree isolation
One-click merge of completed worktrees back to the source branch
Issues are the start, not a detour.
Browse repo issues inside a native board, pick one, and launch a session with the title, description, labels, and comments already loaded as context. Triage and execution stay connected.
Native GitHub issue board inside the app
Issue context auto-loaded into the session
Sub-issue progress tracked alongside execution
Bring your own keys. Pi handles the auth.
No extra subscription. Use your existing Claude, OpenAI, or local-model setup through Pi. Agent Deck adds a per-session picker, on-device drafting, and runtime tools.
Per-session model picker across every Pi provider
On-device Foundation Model for titles, commit messages, and summaries
OpenAI fast mode and Exa web search built in
Free, open source, built for macOS 26 and later.
