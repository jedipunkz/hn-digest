---
source: "https://github.com/pssah4/vault-operator"
hn_url: "https://news.ycombinator.com/item?id=48351115"
title: "Karpathy LLM Wiki pattern integrated into Obsidian agenic workflow"
article_title: "GitHub - pssah4/vault-operator: Real AI agent for your vault. Coworker, Copilot & thinking partner, that maintains your memory & knowledge, adapts to your workflows, uses plugins, skills & tools with full safety controls. BYOK & MCP. · GitHub"
author: "pssah4"
captured_at: "2026-06-01T00:24:26Z"
capture_tool: "hn-digest"
hn_id: 48351115
score: 3
comments: 0
posted_at: "2026-06-01T00:03:11Z"
tags:
  - hacker-news
  - translated
---

# Karpathy LLM Wiki pattern integrated into Obsidian agenic workflow

- HN: [48351115](https://news.ycombinator.com/item?id=48351115)
- Source: [github.com](https://github.com/pssah4/vault-operator)
- Score: 3
- Comments: 0
- Posted: 2026-06-01T00:03:11Z

## Translation

タイトル: Karpathy LLM Wiki パターンを Obsidian agenic ワークフローに統合
記事のタイトル: GitHub - pssah4/vault-operator: ボールト用の実際の AI エージェント。あなたの記憶と知識を維持し、ワークフローに適応し、完全な安全制御を備えたプラグイン、スキル、ツールを使用する同僚、副操縦士、思考パートナー。 BYOKとMCP。 · GitHub
説明: ボールト用の本物の AI エージェント。あなたの記憶と知識を維持し、ワークフローに適応し、完全な安全制御を備えたプラグイン、スキル、ツールを使用する同僚、副操縦士、思考パートナー。 BYOKとMCP。 - pssah4/vault-operator

記事本文:
GitHub - pssah4/vault-operator: ボールト用の実際の AI エージェント。あなたの記憶と知識を維持し、ワークフローに適応し、完全な安全制御を備えたプラグイン、スキル、ツールを使用する同僚、副操縦士、思考パートナー。 BYOKとMCP。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードして更新します

セッション。
アラートを閉じる
{{ メッセージ }}
pssah4
/
保管庫オペレーター
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,358 コミット 1,358 コミット .dia .dia .github/ workflows .github/ workflows アセット/ テンプレート アセット/ テンプレート バンドル-スキル バンドル-スキル バンドル-テンプレート/ ノート バンドル-テンプレート/ ノート docs docs リレー リレー src src ツール/ テンプレート アナライザー ツール/ テンプレート アナライザー .gitignore .gitignore .npmrc .npmrc .prettierrc .prettierrc ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 NOTICE PRIVACY.md PRIVACY.md README.md README.md REVIEWER_NOTES.md REVIEWER_NOTES.md SECURITY.md SECURITY.md _verify_test.ts _verify_test.ts デプロイ-ローカル.sh デプロイ-ローカル.sh esbuild.config.mjs esbuild.config.mjs eslint.config.mjs eslint.config.mjs main.js main.js マニフェスト.json マニフェスト.json mcp-server-worker.js mcp-server-worker.js package-lock.json package-lock.json package.json package.json Sandbox-worker.js Sandbox-worker.jsstyles.cssstyles.css tsconfig.json tsconfig.json version.json version.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Obsidian 保管庫内の自律型 AI エージェント。
タスクを記述すると、タスクが計画、検索、読み取り、書き込みを行い、レポートを返します。あらゆるアクションが目に見えるようになる。すべての書き込みにはあなたの承認が必要です。すべての変更はワンクリックで元に戻すことができます。
無料。オープンソース。地元第一主義。クラウド モデル、既存の ChatGPT または Copilot サブスクリプション、または Ollama または LM Studio を使用して完全にオフラインで動作します。
ドキュメント | Obsidian からインストール |コミュニティページ
これがサイドバー AI チャット以上のものである理由
チャットボットがプロンプトを読み取って回答します。ヴォールト・オペラ

tor はループを実行します。ツールを選択し、それらを Vault に対して実行し、結果をモデルにフィードバックし、タスクが完了するまで継続します。そのループが違いです。
それは Vault だけではなく、Vault に対しても作用します。読み取り、編集、作成、リンク、リファクタリング。 「ここに書き込めるもの」ではなく、目の前にある実際のファイルです。
ボールトの構造を学習します。フォルダー、ウィキリンク、フロントマター、タグ、プラグイン。毎ターン最初から始めるのではなく、そこにあるものを使用します。
それはあなたを学びます。セッション全体にわたる 3 層の記憶: 短期的なセッションの概要、長期的に持続する事実、およびエージェントの書き方とエージェントの動作方法に関するプロファイル。
AI サーフェス全体で機能します。 MCP サーバーとして実行されるため、ChatGPT、Claude Desktop、または Perplexity は、Obsidian 内エージェントと同じメモリと履歴を読み取ることができます。どの AI クライアントがアイデアを捉えたかに関係なく、思考の 1 つのスレッド。
各ステップに適切なモデルが選択されます。プロバイダーを一度構成すると、プラグインはモデルをバジェット層、メイン層、フロンティア層に分類し、引き続きジョブを実行する最も安価な層に作業をルーティングします。
知識労働に役立つこと
このプラグインは、深刻な保管庫の日常的な現実を中心に構築されています。コンテキストを失うことなく新しいソースを取得し、6 か月前に書いたものを見つけ、すでに持っている資料からドキュメントを構築し、成長に合わせて全体をナビゲートできるようにします。
出所を明らかにした情報源を取得する
Knowledge Vault で最も高価な失敗モードは、結論をなぜ信頼したかを忘れてしまうことです。ソースに戻る経路のない音は減衰します。
Vault Operator は、ブロックレベルの出自でこれを解決します。 PDF をチャットにドロップし、取り込みを要求すると、エージェントはトリアージ ステップ (10 秒間、何も読み取る前にボールト、メモリ、およびチャット履歴を確認します) を実行し、クリーンなソース ノートを生成します。エヴ

ery key クレームは、ソース内の正確な段落を解決する ↗ リンクで終わります。ワンクリックで元の文言に戻ります。
/ingest を使用すると、簡単にキャプチャできます。 1 つのドロップ、1 つの承認、1 つのメモ。 3分くらいかな。
/ingest-deep 意味を理解するために。ガイド付きの 7 ステップのダイアログでは、どのトピックをどのような形式で抽出するかを尋ね、ソース段落まで遡って派生メモを作成します。実際の研究論文の場合は 5 ～ 15 分です。
センスメイキングチュートリアル |ブロックレベルの出自の概念
ファイル名ではなく意味で検索する
全文キーワード検索、ウィキリンクによるグラフ展開、およびローカル クロスエンコーダ リランカーと組み合わせた、保管庫上のローカル ベクトル インデックス。 「X について何を知っていますか?」と尋ねます。エージェントは、たとえどのメモにもあなたが使用した単語が含まれていないとしても、意味が関連するメモを見つけます。
背景分析により、ウィキリンクが存在せずに同様のトピックについて議論しているメモのペアも明らかになります。これは、ほとんどの金庫が隠された構造を明らかにする瞬間です。
Word、Excel、およびドラフト PPTX ファイルを構築する (ベータ版)
プロジェクトのメモを Word ドキュメントに、構造化データを Excel に、会議のメモを PowerPoint のドラフトデッキに変換します。 DOCX および XLSX 出力はクリーンで信頼性があります。 PPTX はベータ版です。プラグインには 3 つのデフォルト テーマと 5 つのレイアウトが付属していますが、実際の企業テンプレートの複製はこのバージョンではサポートされていません。クライアント向けデッキの場合は、生成されたファイルを開始点として扱い、手動で仕上げを完了します。
Office ドキュメント ガイド (ベータ版の詳細)
コンテナーのヘルス チェックでは、ナレッジ グラフで孤立したノート、リンク切れ、バックリンクの欠落、弱いクラスター、一貫性のないタグ、過剰接続されたハブ ノートがないか監査します。調査結果にはアクションが伴います。つまり、機械的な修正を適用する、エージェントとのディスカッションを開始する、または却下するなどです。修復するたびに、元に戻すことのできるチェックポイントが作成されます。

Vault Operator はフェールクローズされています。そのカテゴリの自動承認を選択していない限り、書き込み操作には承認が必要です。すべてのタスクは、シャドウ Git リポジトリ (独自の Git 履歴とは別) にチェックポイントを作成します。チャットで [すべての変更を元に戻す] をクリックすると、ファイルが元に戻ります。機密フォルダーは、.obsidian-agentignore ファイルを介してエージェントからロックできます。
安全性と制御ガイド |チェックポイントの概念
インストールします。 Obsidian 設定 > コミュニティ プラグイン > 参照 > 「Vault Operator」 > インストール + 有効化。
プロバイダーを追加します。 [設定] > [Vault Operator] > [プロバイダー] > [+ プロバイダーの追加]。無料の Google AI Studio キーがあれば、すべてを試すのに十分です。
サイドバーを開いて質問してください。 「最もリンクされているノートは何ですか?」どのボールトでも動作します。 First-Run ウィザードが残りの手順を案内します。
セマンティック検索と取り込みワークフローの場合は、[設定] > [埋め込み] で埋め込みモデルも構成します。クイック スタート チュートリアルでは、すべての手順が説明されています。
完全なドキュメントは pssah4.github.io/vault-operator にあります。
チュートリアル 。最初のインストールから /ingest-deep を使用して理解するまでの段階的なウォークスルー。
ガイド。日々の仕事の参考に。
参照 。ツール、プロバイダー、設定、トラブルシューティング。
コードベース ツアー 。ディレクトリのレイアウト、読み取り順序、Kilo Code の伝統。
コンセプト。エージェント ループ、ガバナンス、ナレッジ レイヤー、メモリ システム、MCP アーキテクチャ。
git clone https://github.com/pssah4/vault-operator.git
cd ボールトオペレーター
npmインストール
npm ビルドを実行する
次に、 main.js 、manifest.json 、およびstyles.cssをリポジトリのルートから <vault>/.obsidian/plugins/vault-operator/ にコピーします。開発中の監視モード + 自動デプロイの場合は、.env 内の PLUGIN_DIR をテスト ボールトに指定し、 npm run dev を実行します。
要件: Obsidian 1.4 以降 (Bases 機能については 1.8 以降)、デスクトップのみ、ビルド用の Node.js 18 以降。
ネットワークの使用状況とローカル機能
ボールトオペレーター

r はローカルファーストです。テレメトリも分析もアカウントもありません。
プラグインは 3 つの状況でネットワーク リクエストを行い、すべて制御できます。
LLM API は、構成したプロバイダー (Anthropic、OpenAI、Google、AWS Bedrock、OpenRouter、Azure、GitHub Copilot OAuth、ChatGPT OAuth、Kilo Gateway、Ollama、LM Studio、または任意の OpenAI 互換エンドポイント) を呼び出します。
Web_search ツールを使用すると、Brave または Tavily に移動する Web 検索 (オプション、デフォルトでは無効)。
明示的に接続した MCP サーバーと、ChatGPT または Claude Desktop を使用したクロスサーフェス ワークフローが必要な場合はオプションのリモート MCP リレー。
このプラグインは、標準の Obsidian API を超えるいくつかの Node.js 機能も使用します。つまり、ローカル ナレッジ データベースとオフィス ドキュメント パイプラインのためのファイル システム アクセス、チェックポイントのためのシャドウ git、evaluate_expression のためのサンドボックス プロセスの生成、およびプレゼンテーション レンダリングのためのオプションの LibreOffice 生成です。すべての書き込みは、ボールト パスまたはプラグイン データ ディレクトリの下に残ります。コマンドは、構造化された引数を持つ固定バイナリです。エージェントはチャット テキストからシェル コマンドを構築しません。
API キーは Electron のsafeStorage (macOS の OS キーチェーン、Windows の Credential Manager、Linux の libsecret) を介して暗号化されます。 safeStorage が利用できない場合、キーはプレーンなプラグイン設定に戻ります。
建築上のインスピレーションのためのキロコード。
WebAssembly の SQLite 用の sql.js はナレッジ レイヤーを強化します。
ローカル ONNX の再ランキング用の Hugging Face Transformers.js。
純粋な JS git チェックポイントの isomorphic-git。
モデル コンテキスト プロトコル用の MCP SDK。
あなたの保管庫のための本物の AI エージェント。あなたの記憶と知識を維持し、ワークフローに適応し、完全な安全制御を備えたプラグイン、スキル、ツールを使用する同僚、副操縦士、思考パートナー。 BYOKとMCP。
pssah4.github.io/vault-operator/
トピックス
Apache-2.0ライセンス
コン

貢ぐ
読み込み中にエラーが発生しました。このページをリロードしてください。
16
フォーク
レポートリポジトリ
リリース
116
2.12.8
最新の
2026 年 5 月 31 日
+ 115 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Real AI agent for your vault. Coworker, Copilot & thinking partner, that maintains your memory & knowledge, adapts to your workflows, uses plugins, skills & tools with full safety controls. BYOK & MCP. - pssah4/vault-operator

GitHub - pssah4/vault-operator: Real AI agent for your vault. Coworker, Copilot & thinking partner, that maintains your memory & knowledge, adapts to your workflows, uses plugins, skills & tools with full safety controls. BYOK & MCP. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
pssah4
/
vault-operator
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,358 Commits 1,358 Commits .dia .dia .github/ workflows .github/ workflows assets/ templates assets/ templates bundled-skills bundled-skills bundled-templates/ notes bundled-templates/ notes docs docs relay relay src src tools/ template-analyzer tools/ template-analyzer .gitignore .gitignore .npmrc .npmrc .prettierrc .prettierrc ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md REVIEWER_NOTES.md REVIEWER_NOTES.md SECURITY.md SECURITY.md _verify_test.ts _verify_test.ts deploy-local.sh deploy-local.sh esbuild.config.mjs esbuild.config.mjs eslint.config.mjs eslint.config.mjs main.js main.js manifest.json manifest.json mcp-server-worker.js mcp-server-worker.js package-lock.json package-lock.json package.json package.json sandbox-worker.js sandbox-worker.js styles.css styles.css tsconfig.json tsconfig.json versions.json versions.json vitest.config.ts vitest.config.ts View all files Repository files navigation
An autonomous AI agent inside your Obsidian vault.
You describe a task, it plans, searches, reads, writes, and reports back. Every action is visible. Every write needs your approval. Every change is undoable in one click.
Free. Open source. Local-first. Works with cloud models, with your existing ChatGPT or Copilot subscription, or fully offline with Ollama or LM Studio.
Documentation | Install from Obsidian | Community page
Why this is more than a sidebar AI chat
A chatbot reads your prompt and answers. Vault Operator runs a loop: it picks tools, executes them against your vault, feeds the results back to the model, and continues until the task is done. That loop is the difference.
It acts on your vault, not just about it. Reading, editing, creating, linking, refactoring. Not "here is what you could write", but the actual file in front of you.
It learns your vault structure. Folders, wikilinks, frontmatter, tags, plugins. It uses what is there instead of starting from scratch every turn.
It learns you. Three-tier memory across sessions: short-term session summaries, long-term durable facts, and a profile of how you write and how you like the agent to behave.
It works across your AI surfaces. Runs as an MCP server so ChatGPT, Claude Desktop, or Perplexity can read the same memory and history as the in-Obsidian agent. One thread of thinking, regardless of which AI client captured the idea.
It picks the right model for each step. Configure a provider once, the plugin sorts the models into Budget, Main, and Frontier tiers and routes work to the cheapest tier that still does the job.
What it does for knowledge work
The plugin is built around the daily reality of a serious vault: capturing new sources without losing context, finding what you wrote six months ago, building documents from material you already have, and keeping the whole thing navigable as it grows.
Capture sources with provenance
The most expensive failure mode in a knowledge vault is forgetting why you trusted a conclusion. A note without a path back to its source decays.
Vault Operator solves this with block-level provenance. Drop a PDF into the chat, ask for an ingest, and the agent runs a triage step (ten seconds, looks at vault, memory, and chat history before reading anything), then produces a clean source note. Every key claim ends with a ↗ link that resolves to the exact paragraph in the source. One click and you are back at the original wording.
/ingest for quick capture. One drop, one approval, one note. About three minutes.
/ingest-deep for sense-making. A guided seven-step dialog that asks which topics to extract and in what shape, then writes derived notes that all trace back to the source paragraph. Five to fifteen minutes for a real research paper.
Sense-making tutorial | Block-level provenance concept
Search by meaning, not by filename
A local vector index over your vault, combined with full-text keyword search, graph expansion through wikilinks, and a local cross-encoder reranker. Ask "what do I know about X?" and the agent finds notes whose meaning is related, even if none of them contain the words you used.
The background analysis also surfaces note pairs that discuss similar topics without any wikilink between them. This is the moment most vaults reveal hidden structure.
Build Word, Excel, and draft PPTX files (beta)
Turn project notes into a Word document, structured data into Excel, or meeting notes into a draft PowerPoint deck. DOCX and XLSX output is clean and reliable. PPTX is in beta: the plugin ships with three default themes and five layouts, but real corporate template cloning is not supported in this version. For client-facing decks, treat the generated file as a starting point and finish the polish manually.
Office documents guide (beta details)
The vault health check audits your knowledge graph for orphaned notes, broken links, missing backlinks, weak clusters, inconsistent tags, and over-connected hub notes. Findings come with actions: apply a mechanical fix, open a discussion with the agent, or dismiss. Every repair creates a checkpoint you can undo.
Vault Operator is fail-closed. Write operations need your approval unless you opted into auto-approve for that category. Every task creates checkpoints in a shadow git repository (separate from your own git history). Click "Undo all changes" in the chat and the files go back. Sensitive folders can be locked from the agent via a .obsidian-agentignore file.
Safety and control guide | Checkpoints concept
Install. Obsidian Settings > Community Plugins > Browse > "Vault Operator" > Install + Enable.
Add a provider. Settings > Vault Operator > Providers > "+ Add provider". A free Google AI Studio key is enough to try everything.
Open the sidebar and ask a question. "What are my most-linked notes?" works on any vault. The First-Run wizard walks you through the rest.
For semantic search and the ingest workflows, also configure an embedding model in Settings > Embeddings. The Quick Start tutorial covers every step.
Full documentation lives at pssah4.github.io/vault-operator .
Tutorials . Step-by-step walkthroughs from first install to sense-making with /ingest-deep .
Guides . Reference for daily work.
Reference . Tools, providers, settings, troubleshooting.
Codebase tour . Directory layout, reading order, Kilo Code heritage.
Concepts . Agent loop, governance, knowledge layer, memory system, MCP architecture.
git clone https://github.com/pssah4/vault-operator.git
cd vault-operator
npm install
npm run build
Then copy main.js , manifest.json , and styles.css from the repo root into <vault>/.obsidian/plugins/vault-operator/ . For watch mode + auto-deploy during development, point PLUGIN_DIR in .env at your test vault and run npm run dev .
Requirements: Obsidian 1.4+ (1.8+ for Bases features), desktop only, Node.js 18+ for building.
Network usage and local capabilities
Vault Operator is local-first. No telemetry, no analytics, no accounts.
The plugin makes network requests in three situations, all under your control:
LLM API calls to the provider you configured (Anthropic, OpenAI, Google, AWS Bedrock, OpenRouter, Azure, GitHub Copilot OAuth, ChatGPT OAuth, Kilo Gateway, Ollama, LM Studio, or any OpenAI-compatible endpoint).
Web search (optional, disabled by default) when you use the web_search tool, going to Brave or Tavily.
MCP servers you connected explicitly, plus the optional remote-MCP relay if you want cross-surface workflows with ChatGPT or Claude Desktop.
The plugin also uses a few Node.js capabilities that go beyond the standard Obsidian API: filesystem access for the local knowledge database and the office document pipeline, shadow git for checkpoints, sandbox process spawning for evaluate_expression , and optional LibreOffice spawning for presentation rendering. All writes stay under the vault path or the plugin data directory. Commands are fixed binaries with structured arguments; the agent does not construct shell commands from chat text.
API keys are encrypted via Electron's safeStorage (OS keychain on macOS, Credential Manager on Windows, libsecret on Linux). Where safeStorage is not available, keys fall back to plain plugin settings.
Kilo Code for architectural inspiration.
sql.js for SQLite in WebAssembly powering the knowledge layer.
Hugging Face Transformers.js for local ONNX reranking.
isomorphic-git for pure-JS git checkpoints.
MCP SDK for the Model Context Protocol.
Real AI agent for your vault. Coworker, Copilot & thinking partner, that maintains your memory & knowledge, adapts to your workflows, uses plugins, skills & tools with full safety controls. BYOK & MCP.
pssah4.github.io/vault-operator/
Topics
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
16
forks
Report repository
Releases
116
2.12.8
Latest
May 31, 2026
+ 115 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
