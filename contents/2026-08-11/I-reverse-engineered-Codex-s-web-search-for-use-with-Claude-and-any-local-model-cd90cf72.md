---
source: "https://github.com/mateusdcc/pi-gpt-search"
hn_url: "https://news.ycombinator.com/item?id=49255837"
title: "I reverse-engineered Codex's web search for use with Claude, and any local model"
article_title: "GitHub - mateusdcc/pi-gpt-search · GitHub"
author: "mateusdcc"
captured_at: "2026-08-11T10:42:28Z"
capture_tool: "hn-digest"
hn_id: 49255837
score: 2
comments: 0
posted_at: "2026-08-11T10:17:01Z"
tags:
  - hacker-news
  - translated
---

# I reverse-engineered Codex's web search for use with Claude, and any local model

- HN: [49255837](https://news.ycombinator.com/item?id=49255837)
- Source: [github.com](https://github.com/mateusdcc/pi-gpt-search)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T10:17:01Z

## Translation

タイトル: クロードおよびローカル モデルで使用するために Codex の Web 検索をリバース エンジニアリングしました
記事タイトル: GitHub - mateusdcc/pi-gpt-search · GitHub
説明: GitHub でアカウントを作成して、mateusdcc/pi-gpt-search の開発に貢献します。

記事本文:
GitHub - mateusdcc/pi-gpt-search · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
マテウスdcc
/
pi-gpt-検索
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
48 コミット 48 コミット src src テスト テスト .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md HOW-IT-WAS-EXTRACT.md HOW-IT-WAS-EXTRACT.md HOW-IT-WORKS.md HOW-IT-WORKS.md README.md README.md package-lock.json package-lock.json package.j

Son package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenAI Codex スタンドアロン検索エンジンを使用した、Pi のネイティブでモデルに依存しない Web 検索。
pi-gpt-search は、OpenAI Codex のスタンドアロン Web 取得インフラストラクチャを再利用することにより、あらゆる Pi モデル (Gemini、Claude、ローカル モデル、OpenRouter) にリアルタイム Web 検索機能を提供します。GPT モデル推論ターン数と GPT トークンの消費はゼロです。
⚡ クイックスタート: 1 行インストール
pi インストール npm:pi-gpt-search
または、GitHub 経由でインストールします。
pi インストール https://github.com/mateusdcc/pi-gpt-search
または、現在のリポジトリに対してプロジェクトをローカルにインストールします ( -l フラグ)。
pi インストール npm:pi-gpt-search -l
または、インストールせずに単一セッションで一時的に試してください。
pi -e npm:pi-gpt-search
⚡ 主なハイライト: ZERO-GPT 推論
🚀 GPT トークンの消費ゼロ: OpenAI のバックエンド エンドポイントを介した純粋な Web 取得。 GPT/Codex LLM ターンは実行されません。つまり、0 入力トークン、0 出力トークン、0 推論クレジットが請求されます。
👑 モデル主権: アクティブな Pi モデル (例: Gemini 3.5 Flash / Gemini 3.1 Pro) が唯一の推論モデルのままです。
🛠️ スラッシュ コマンドと LLM ツール: LLM ツール ( codex-search および codex-research ) として、また直接ユーザー コマンド ( /gpt-search ) として自動的に機能します。
🔑 資格情報の再利用: 既存のコーデックス ログイン セッション ( ~/.codex/auth.json ) またはカスタム .env トークンを自動的に使用します。
🛡️ データプライバシー: デフォルトではクエリのみ。会話履歴、プロジェクト ファイル、またはシステム プロンプトを検索のために送信しません。
Pi コーディング エージェント
━── ジェミニ（または現役モデル）
§── codex-search(クエリ: "最新の Rust リリース")
│ └── Codex/OpenAI スタンドアロン検索 API (/codex/alpha/search)
│ └── 構造化された結果 (タイトル、URL、スニペット)
│ └── ジェミニは推理を続け、ユーザーに答える
│
━━

codex-research(search_query: [...]、open: [...]、find: [...])
└── マルチステップ Web リサーチ ハーネス
└── 深い文書内容、パターンマッチングと引用
🛠️ 使用法とコマンド
LLM トークンを消費せずに、いずれかの Codex ツールを自分で実行します。
/codex-search Rust 1.97 リリースノート
/codex-research OpenAI Codex GitHub リポジトリ
素早い検索には codex-search を使用し、より包括的な結果が必要な場合には codex-research を使用します。ダイレクト コマンドは、ツールのデフォルトを保持します。つまり、 codex-search の短縮形と codex-research の長い短縮形です。
/gpt-search <query> は、単純な従来のエイリアスとして引き続き使用できます。
2. 自動 LLM ツール: codex-search
任意のモデルに、現在の事実を必要とする質問をします (単一クエリ ルックアップ)。
pi --model antigravity/gemini-3.5-flash 「Rust の最新リリースと何が変更されましたか?」
モデルは、最新の情報が必要な素早い検索のためにこれを自動的に使用します。
ログ出力の例 ( PI_WEB_SEARCH_DEBUG=1 の場合):
[PI_WEB_SEARCH_DEBUG] req_id=maqk8a5 query="最新の Rust リリース バージョンと日付 2026" Provider=codex
[PI_WEB_SEARCH_DEBUG] req_id=maqk8a5 status=200 elapsed_ms=1863 results=41
3. 自動先端研究ハーネスツール: codex-research
モデルに、マルチクエリの実行、ページ コンテンツの検査、パターン検索、リンク ナビゲーションによる詳細な反復的な Web リサーチを実行するよう依頼します。モデルは、調査の手順とソースを管理します。
4. レガシーエイリアス: web (非推奨)
名前変更前のツール名 web は、下位互換性のあるエイリアスとして保持されます。 codex-research と同じ実装に委任され、すべての呼び出しの前に非推奨の通知が追加されます。新しい統合では codex-research を直接使用する必要があります。
Pi コーディング エージェント: pi CLI がインストールされています (v0.80+)。
OpenAI Codex Auth: 認証された Codex セッション (ターミナルで codex ログインを実行するか、.e で CODEX_ACCESS_TOKEN を設定します)

nv)。
⚙️ 手動インストールと環境セットアップ
pi install ではなく手動で配置したい場合:
# グローバル (すべてのプロジェクト)
mkdir -p ~ /.pi/agent/extensions
cp -r pi-gpt-search ~ /.pi/agent/extensions/
# プロジェクトローカル
mkdir -p .pi/extensions
cp -r pi-gpt-search .pi/extensions/
2. 環境変数 (オプション)
Codex アクセス トークンを明示的にオーバーライドする場合は、.env.example を .env にコピーします。
cp .env.example .env
.env を編集します。
# オプション: 設定されていない場合は、~/.codex/auth.json を自動的に読み取ります
CODEX_ACCESS_TOKEN = your_token_here
CODEX_ACCOUNT_ID = your_account_id_ここ
# デバッグログを有効にする
PI_WEB_SEARCH_DEBUG = 1
セキュリティ上の注意: .env を Git にコミットしないでください。 .env は .gitignore にリストされています。
pi-gpt-search には、4 レベルのテスト スイートが付属しています。
npmテスト
テストスイートの内訳:
単体テスト (unit.test.ts 、command.test.ts 、normalize.test.ts 、output.test.ts 、web-tool.test.ts ): スキーマ検証、DTO 正規化、エラー クラス、出力書式設定、折りたたみ可能な表示。
統合テスト ( Provider-integration.test.ts ): 200、401、403、429、500、タイムアウト、キャンセルの模擬サーバー処理。
Real Search Test ( real-search.test.ts & real-endpoint.test.ts ): OpenAI の検索エンドポイントおよびセッション継続性に対するライブ実行。
Zero-GPT 検証 ( zero-gpt.test.ts ): GPT 推論呼び出しが行われていないことを証明するネットワーク傍受テスト。
E2E Research Harness Suite ( e2e-research.test.ts ): 完全なエンドツーエンドのマルチステップ Web リサーチ テスト スイート。
HOW-IT-WORKS.md - モジュール、データ フロー、TUI レンダラー、コンテキストの分離、およびキャンセルのアーキテクチャの詳細な詳細。
HOW-IT-WAS-EXTRACT.md - スタンドアロン検索エンドポイントがどのように発見されたかを文書化したリバース エンジニアリング ガイド。
検索インデックスの範囲: 検索結果のスニペット、URL、およびドキュメント ビューを返します。完全なヘッドレス BR は含まれません

所有者の DOM レンダラー。
セッション認証: アクティブな ChatGPT/Codex ログイン セッション ( codex login ) が必要です。セッションが期限切れになった場合は、コーデックス ログインを実行して再認証する必要があります。
3 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to mateusdcc/pi-gpt-search development by creating an account on GitHub.

GitHub - mateusdcc/pi-gpt-search · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
mateusdcc
/
pi-gpt-search
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
48 Commits 48 Commits src src tests tests .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md HOW-IT-WAS-EXTRACT.md HOW-IT-WAS-EXTRACT.md HOW-IT-WORKS.md HOW-IT-WORKS.md README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Native, Model-Independent Web Search for Pi using OpenAI Codex Standalone Search Engine.
pi-gpt-search gives any Pi model (Gemini, Claude, local models, OpenRouter) real-time web search capabilities by reusing OpenAI Codex's standalone web retrieval infrastructure - with ZERO GPT Model Inference Turns and ZERO GPT Tokens Consumed .
⚡ Quick Start: 1-Line Installation
pi install npm:pi-gpt-search
Or install via GitHub:
pi install https://github.com/mateusdcc/pi-gpt-search
Or install project-locally for your current repository ( -l flag):
pi install npm:pi-gpt-search -l
Or try it temporarily in a single session without installing:
pi -e npm:pi-gpt-search
⚡ Key Highlights: ZERO-GPT INFERENCE
🚀 Zero GPT Tokens Spent: Pure web retrieval via OpenAI's backend endpoint. No GPT/Codex LLM turns are executed, meaning 0 input tokens, 0 output tokens, and 0 reasoning credits are billed .
👑 Model Sovereign: Your active Pi model (e.g., Gemini 3.5 Flash / Gemini 3.1 Pro) remains the sole reasoning model.
🛠️ Slash Command & LLM Tools: Works both automatically as LLM tools ( codex-search & codex-research ) and as a direct user command ( /gpt-search ).
🔑 Credential Reuse: Automatically uses your existing codex login session ( ~/.codex/auth.json ) or custom .env tokens.
🛡️ Data Privacy: Query-only by default. Does not send conversation history, project files, or system prompts to search.
Pi Coding Agent
└── Gemini (or active model)
├── codex-search(query: "latest Rust release")
│ └── Codex/OpenAI Standalone Search API (/codex/alpha/search)
│ └── Structured Results (Title, URL, Snippet)
│ └── Gemini continues reasoning & answers user
│
└── codex-research(search_query: [...], open: [...], find: [...])
└── Multi-Step Web Research Harness
└── Deep document content, pattern matching & citations
🛠️ Usage & Commands
Run either Codex tool yourself without spending LLM tokens:
/codex-search Rust 1.97 release notes
/codex-research OpenAI Codex GitHub repository
Use codex-search for a quick lookup and codex-research when you want more comprehensive results. The direct commands preserve their tool defaults: short for codex-search and long for codex-research .
/gpt-search <query> remains available as a simple legacy alias.
2. Automatic LLM Tool: codex-search
Ask any model a question requiring current facts (single-query lookup):
pi --model antigravity/gemini-3.5-flash " What is the latest release of Rust and what changed? "
The model uses it automatically for quick lookups that need current information.
Example Log Output (with PI_WEB_SEARCH_DEBUG=1 ):
[PI_WEB_SEARCH_DEBUG] req_id=maqk8a5 query="latest Rust release version and date 2026" provider=codex
[PI_WEB_SEARCH_DEBUG] req_id=maqk8a5 status=200 elapsed_ms=1863 results=41
3. Automatic Advanced Research Harness Tool: codex-research
Ask models to conduct deep, iterative web research with multi-query execution, page content inspection, pattern finding, and link navigation. The model manages the research steps and sources for you.
4. Legacy Alias: web (deprecated)
The pre-rename tool name web is kept as a backward-compatible alias. It delegates to the same implementation as codex-research and prepends a deprecation notice on every invocation. New integrations should use codex-research directly.
Pi Coding Agent: pi CLI installed ( v0.80+ ).
OpenAI Codex Auth: An authenticated Codex session (run codex login in terminal, or set CODEX_ACCESS_TOKEN in .env ).
⚙️ Manual Installation & Environment Setup
If you prefer manual placement instead of pi install :
# Global (All projects)
mkdir -p ~ /.pi/agent/extensions
cp -r pi-gpt-search ~ /.pi/agent/extensions/
# Project-local
mkdir -p .pi/extensions
cp -r pi-gpt-search .pi/extensions/
2. Environment Variables (Optional)
Copy .env.example to .env if you want to explicitly override your Codex access token:
cp .env.example .env
Edit .env :
# Optional: If unset, automatically reads ~/.codex/auth.json
CODEX_ACCESS_TOKEN = your_token_here
CODEX_ACCOUNT_ID = your_account_id_here
# Enable debug logging
PI_WEB_SEARCH_DEBUG = 1
Security Note: Never commit .env to Git. .env is listed in .gitignore .
pi-gpt-search comes with a 4-level test suite:
npm test
Test suite breakdown:
Unit Tests ( unit.test.ts , commands.test.ts , normalize.test.ts , output.test.ts , web-tool.test.ts ): Schema validation, DTO normalization, error classes, output formatting, collapsible display.
Integration Tests ( provider-integration.test.ts ): Mock server handling for 200, 401, 403, 429, 500, timeouts, cancellation.
Real Search Test ( real-search.test.ts & real-endpoint.test.ts ): Live execution against OpenAI's search endpoint and session continuity.
Zero-GPT Verification ( zero-gpt.test.ts ): Network interception test proving 0 GPT inference calls are made.
E2E Research Harness Suite ( e2e-research.test.ts ): Full end-to-end multi-step web research test suite.
HOW-IT-WORKS.md - Deep architectural breakdown of modules, data flow, TUI renderers, context isolation, and cancellation.
HOW-IT-WAS-EXTRACT.md - Reverse-engineering guide documenting how the standalone search endpoint was discovered.
Search Index Scope: Returns search result snippets, URLs, and document views; does not include a full headless browser DOM renderer.
Session Auth: Requires an active ChatGPT/Codex login session ( codex login ). Expired sessions require running codex login to re-authenticate.
3 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
