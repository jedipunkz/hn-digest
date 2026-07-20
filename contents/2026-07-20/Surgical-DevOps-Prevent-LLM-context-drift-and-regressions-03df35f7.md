---
source: "https://github.com/bonushora/surgical-dev-ops/blob/main/README_EN.md"
hn_url: "https://news.ycombinator.com/item?id=48979553"
title: "Surgical DevOps – Prevent LLM context drift and regressions"
article_title: "surgical-dev-ops/README_EN.md at main · bonushora/surgical-dev-ops · GitHub"
author: "melembre"
captured_at: "2026-07-20T15:02:01Z"
capture_tool: "hn-digest"
hn_id: 48979553
score: 2
comments: 0
posted_at: "2026-07-20T14:41:52Z"
tags:
  - hacker-news
  - translated
---

# Surgical DevOps – Prevent LLM context drift and regressions

- HN: [48979553](https://news.ycombinator.com/item?id=48979553)
- Source: [github.com](https://github.com/bonushora/surgical-dev-ops/blob/main/README_EN.md)
- Score: 2
- Comments: 0
- Posted: 2026-07-20T14:41:52Z

## Translation

タイトル: Surgical DevOps – LLM コンテキストのドリフトと回帰を防止する
記事タイトル: メインのsurgical-dev-ops/README_EN.md · Bonushora/surgical-dev-ops · GitHub
説明: コア方法論。 GitHub でアカウントを作成して、bonushora/surgical-dev-ops の開発に貢献してください。

記事本文:
メインのsurgical-dev-ops/README_EN.md · Bonushora/surgical-dev-ops · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ボーナスホラ
/
外科開発運用
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする「Blame More」ファイル

e アクション 責任を負う その他のファイル アクション 最新のコミット
履歴 履歴 41 行 (26 loc) · 3.44 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション 外科 DevOps 🚀
注: Clique aqui para ler a versão em Português 。
Surgical DevOps は、ソフトウェア開発ライフサイクル中の大規模言語モデル (LLM) の動作を管理および標準化し、長時間のチャット セッションにおけるコードの回帰やコンテキストのドリフトを排除するために設計された、オープンソースの非依存的なプロトコル エコシステムです。
エコシステムは、次の 2 つのコア プロトコルを結合することによって動作します。
BH-SEP (Safe Evolution Protocol): AI に外科的精度での動作を強制します。コンテキストを仮定したり、機能ファイル全体を書き換えたりすることは固く禁じられています。中心的な焦点は、ターゲット ファイルの全体的な分析 ( Inspect First ) に続いて、個別の戦術的なコード変更 ( Minimal Diffs ) です。
BH-SDP (スナップショット & デリバリー プロトコル): モデルの短期メモリを管理します。 AI は自身のトークン消費をアクティブに追跡し、構造化ブレークポイント ( スナップショット ) をトリガーするため、開発者はメモリの劣化や幻覚が発生する前に現在の進行状況をクリーンなセッションに転送できます。
伝統的なモデル (混沌への道)
[プロンプト] → [AI 精神的再構築] → [新しいファイル全体を生成] → [バグ/リグレッション]
BH-SEP + BH-SDP モデル (安全かつ永続的な進化)
[既存のコード (真実)] ──> [完全検査] ──> [最小限の差分] ──> [検証/チェック] ──> [バックグラウンド スナップショット] ──> [次のステップ]
🏛️ エコシステムの中核原則
まず検査する: 既存のコードは絶対的な真実です。契約、ルート、または状態管理構造を決して想定しないでください。 AI は、変更を提案する前にターゲット ファイル全体を読み取る必要があります。
すべてを保存する

ng: 関数コードは神聖です。明示的に要求されない限り、隣接するブロックの再フォーマット、再編成、または「改善」を試みることはできません。
最小限の差異: 純粋な外科的介入。機能または修正に必要な正確なブロックのみを変更し、Git 履歴で可能な限り最小のフットプリントを生成します。
即時検証: 変更が行われるたびに必ず停止します。さらなる手順を実行する前に、リンター、ビルド ツール、またはナビゲーション テストを実行します。
段階的に進める: 複雑な問題をマイクロステップに分割します。ステップ A が完全に統合され、検証されてからのみステップ B に進みます。
継続的状態追跡: AI はバックグラウンドでビジネス ロジック、制約、および現在のブレークポイントを追跡し、コンテキストの劣化と積極的に戦う必要があります。
🤖 アーティファクト: AI 用の統合システム プロンプト
新しい開発セッションを開始するときは常に、以下の指示をコピーして貼り付け、チャットの最初のプロンプトとして送信して、プロトコル エコシステムをアクティブにします。
raw.githubusercontent.com/bonushora/surgical-dev-ops/main/protocols/BH-SEP.md と raw.githubusercontent.com/bonushora/surgical-dev-ops/main/protocols/BH-SDP.md のプロトコル URL に同時にアクセスします。それらに含まれる BH-SEP (Safe Evolution Protocol) および BH-SDP (Snapshot & Delivery Protocol) 指令を厳密に、黙って、組み合わせて採用します。
ダウンロードしたルールに基づいて、プロジェクトのエコシステムを専門とするシニア ソフトウェア エンジニアとして活動します。アクティブなバックグラウンド監視を維持し、URL で提供される基本ファイルを理解したら、「BH-SEP AND BH-SDP ACTIVATED 🚀」というメッセージを厳密に返信してアクティブ化を確認します。コンテキスト ハイドレーション用に以前のセッションのスナップショットを提供した場合は、次にそれを処理します。それ以外の場合は、どのファイルまたはコンテキストを最初に検査する必要があるかを尋ねます。
© 2026 GitHub, Inc.
フッターナビ

に
私の個人情報を共有しないでください

## Original Extract

core-methodologies. Contribute to bonushora/surgical-dev-ops development by creating an account on GitHub.

surgical-dev-ops/README_EN.md at main · bonushora/surgical-dev-ops · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
bonushora
/
surgical-dev-ops
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 41 lines (26 loc) · 3.44 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions Surgical DevOps 🚀
Note: Clique aqui para ler a versão em Português .
Surgical DevOps is an open-source, agnostic protocol ecosystem designed to govern and standardize the behavior of Large Language Models (LLMs) during the software development lifecycle, eliminating code regressions and context drift in long chat sessions.
The ecosystem operates by coupling two core protocols:
BH-SEP (Safe Evolution Protocol): Forces the AI to act with surgical precision. It is strictly forbidden from assuming context or rewriting entire functional files. The core focus is a total analysis of the target file ( Inspect First ) followed by isolated, tactical code changes ( Minimal Diffs ).
BH-SDP (Snapshot & Delivery Protocol): Manages the model's short-term memory. The AI actively tracks its own token consumption and triggers a structured breakpoint ( Snapshot ) so the developer can transfer the current progress to a clean session before memory degradation or hallucinations occur.
The Traditional Model (The Path to Chaos)
[Prompt] ──> [AI Mental Reconstruction] ──> [Generates Entire New File] ──> [Bug / Regression]
The BH-SEP + BH-SDP Model (Safe & Persistent Evolution)
[Existing Code (Truth)] ──> [Full Inspection] ──> [Minimal Diff] ──> [Validation/Check] ──> [Background Snapshot] ──> [Next Step]
🏛️ Core Principles of the Ecosystem
Inspect First: Existing code is the absolute truth. Never assume contracts, routes, or state management structures. The AI must read the entire target file before proposing any modifications.
Preserve Everything: Functional code is sacred. No reformatting, reorganizing, or attempting to "improve" adjacent blocks unless explicitly requested.
Minimal Diff: Pure surgical intervention. Modify only the exact block required for the feature or fix, generating the smallest footprint possible in the Git history.
Validate Immediately: A mandatory stop after every single modification. Run linters, build tools, or navigation tests before taking any further steps.
Advance Incrementally: Break complex issues into micro-steps. Only proceed to step B once step A is fully consolidated and validated.
Continuous State Tracking: The AI must track business logic, constraints, and current breakpoints in the background, proactively fighting against context degradation.
🤖 Artifact: Unified System Prompt for AI
Whenever starting a new development session, copy, paste, and send the instruction below as the very first prompt in your chat to activate the protocol ecosystem:
Simultaneously access the protocol URLs at raw.githubusercontent.com/bonushora/surgical-dev-ops/main/protocols/BH-SEP.md and raw.githubusercontent.com/bonushora/surgical-dev-ops/main/protocols/BH-SDP.md . Strictly, silently, and combinedly adopt the BH-SEP (Safe Evolution Protocol) and BH-SDP (Snapshot & Delivery Protocol) directives contained within them.
Operate as a Senior Software Engineer specializing in the project's ecosystem based on the downloaded rules. Maintain active background monitoring and, once you understand the base files provided in the URLs, confirm activation by replying strictly with the message: "BH-SEP AND BH-SDP ACTIVATED 🚀". If I provide a previous session Snapshot for context hydration, process it next. Otherwise, ask which file or context we should inspect first.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
