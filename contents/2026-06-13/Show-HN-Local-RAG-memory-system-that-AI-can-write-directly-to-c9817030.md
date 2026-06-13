---
source: "https://github.com/ptobey/local-memory-mcp"
hn_url: "https://news.ycombinator.com/item?id=48520532"
title: "Show HN: Local RAG memory system that AI can write directly to"
article_title: "GitHub - ptobey/local-memory-mcp: Persistent local memory for Claude and ChatGPT. No cloud, no subscription. Self-hosted RAG over MCP. · GitHub"
author: "ptobey"
captured_at: "2026-06-13T21:32:48Z"
capture_tool: "hn-digest"
hn_id: 48520532
score: 3
comments: 1
posted_at: "2026-06-13T19:23:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Local RAG memory system that AI can write directly to

- HN: [48520532](https://news.ycombinator.com/item?id=48520532)
- Source: [github.com](https://github.com/ptobey/local-memory-mcp)
- Score: 3
- Comments: 1
- Posted: 2026-06-13T19:23:48Z

## Translation

タイトル: Show HN: AI が直接書き込みできるローカル RAG メモリ システム
記事のタイトル: GitHub - ptobey/local-memory-mcp: Claude と ChatGPT の永続的なローカル メモリ。クラウドもサブスクリプションもありません。 MCP を介した自己ホスト型 RAG。 · GitHub
説明: Claude と ChatGPT の永続的なローカル メモリ。クラウドもサブスクリプションもありません。 MCP を介した自己ホスト型 RAG。 - ptobey/ローカルメモリ-mcp
HN テキスト: 私と私の家族にとって、新しい LLM チャットを作成するたびに情報を再共有しなければならないのは本当に面倒です。したがって、LLM が RAG に対して読み書きできるようにするローカル MCP である local-memory-mcp を作成することにしました。また、最大チャンク サイズや情報更新時の非推奨など、システムの効率を維持するのに役立つシステムも組み込まれています。ぜひ皆さんも試してみて、感想を聞かせていただければ幸いです。私たちにとっては、非常に時間とトークンの節約になりました。ご協力ありがとうございます。楽しんでください!

記事本文:
GitHub - ptobey/local-memory-mcp: Claude と ChatGPT の永続的なローカル メモリ。クラウドもサブスクリプションもありません。 MCP を介した自己ホスト型 RAG。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
プトビー
/
ローカルメモリ-mcp
公共
通知
通知設定を変更するにはサインインする必要があります

イングス
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット docs docs 例 例 src src .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンスREADME.md README.md SECURITY.md SECURITY.md config.example.json config.example.jsonデモ.gif デモ.gif docker-compose.yml docker-compose.yml 要件.txt 要件.txt run_mcp_v1_http_sse.py run_mcp_v1_http_sse.py run_mcp_v1_stdio.py run_mcp_v1_stdio.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI アシスタントは会話が終わるとすべてを忘れます。これはローカルで修正されます。
雲はありません。購読はありません。アカウントがありません。データはマシン上に残ります。
local-memory-mcp は、Claude、ChatGPT、およびその他の MCP 互換アシスタントに、ローカル ベクトル検索 (ChromaDB) を利用した永続メモリ層を提供します。一度何か言ってください。セッションをまたいで記憶されます。
新しい Claude または ChatGPT セッションはすべて空白から始まります。あなたの好み、プロジェクトの背景、決断が失われます。あなたは常に自分自身を説明し直します。
local-memory-mcp は、AI アシスタントに次のことを可能にするローカル MCP サーバーです。
覚えておく価値のあるものを保存する (「私の仕事の集中ブロックは午前 6 時 30 分から 9 時です」)
新しいセッションの開始時に関連するコンテキストを取得します
状況の変化に応じて記憶をバージョンアップして置き換える
これは、AI アシスタントが内蔵すべきメモリ層ですが、内蔵されていません。
クイックスタート (Docker - 2 分)
git clone https://github.com/ptobey/local-memory-mcp.git
cd ローカルメモリ mcp
docker 構成 --build -d
次に、MCP クライアントを http://localhost:8000/mcp に指定します。終わり。
→ Claude Desktop セットアップ · ChatGPT セットアップ · 手動 Python インストール
[MCPクライアント経由のアシスタント]
|
v
[run_mcp_v1

_stdio.py | run_mcp_v1_http_sse.py]
|
v
[src/mcp_server_v1.py]
/ | \
vvv
[vector_store.py] [reconciliation.py] [health_monitor.py]
| |
v v
[ローカルChromaDB] [リコンシリエーションログ]
書き込みパス: 保存/更新によりチャンクが書き込まれます → 調整により重複/競合がチェックされます → 書き込みが危険であると思われる場合は、警告と自己修復ヒントが返されます。
読み取りパス: 検索によりセマンティック検索が実行される → ランキングにより、類似性と軽量の語彙/最新性シグナルがブレンドされる → 明示的に要求されない限り、非推奨のチャンクは非表示のままになります。
MCP ツール: store 、 search 、 update 、 delete 、 get_chunk 、 get_evolution_chain
チェーンを置き換えるバージョン管理された更新 (strategy="version" )
デフォルトではソフト削除 (履歴は保持)、オプションでハード削除
ヒューリスティックな調整と競合のログ記録
構造化された警告[] および自己修復フィールドを使用した警告優先の書き込み応答
サイズが大きすぎるチャンクと未解決の競合のヘルスチェック
永続化されたベクター DB のローカル バックアップ/復元
オプションの SSE 認証: none (ローカルのみ)、 bearer 、または oauth
その背後にある設計思想 (AIX)
AIX (AI eXperience) とは、人間がドキュメントをファイルする方法ではなく、LLM が実際にコンテキストを消費する方法を設計することを意味します。
厳格なドキュメント スキーマよりもクリア テキスト チャンクを優先する
タイムスタンプ、信頼性、リンクの置き換え、非推奨フラグなどのメタデータを最小限に抑えますが、有用です
破壊的な上書きではなく、バージョンチェーンを使用して履歴を保存します
モデルが自己修正できるように、警告を多く含むツール応答を返します。
目標は、実用的な検索品質と信頼できる AI の動作であり、人間による完璧な分類ではありません。
ツール: ストア
input: { "text": "平日のフォーカス ブロックは午前 6 時 30 分から 9 時で、現在のデフォルト スケジュールです。" }
後で取得してください:
ツール: 検索
input: { "query": "現在の詳細な作業スケジュール", "top_k": 5 }
いくつかの焦点を絞った検索を実行して新しいセッションをブートストラップし、その後、acti のみを合成します。

ve、非推奨でないチャンクを新しいモデル インスタンスの短い概要にまとめます。詳しいフローは例/ にあります。
デフォルトではローカルファーストでユーザー制御
構成された永続ディレクトリの下のローカル ChromaDB ファイルに保存されたデータ
クラウド バックエンドは必要ありません。ユーザー管理のトンネリングを介したオプションのリモート アクセス
本物のシークレットは決してコミットしないでください - ローカルの設定/環境値を使用してください
統合 (Claude Desktop + ChatGPT)
Claude と ChatGPT の永続的なローカル メモリ。クラウドもサブスクリプションもありません。 MCP を介した自己ホスト型 RAG。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Persistent local memory for Claude and ChatGPT. No cloud, no subscription. Self-hosted RAG over MCP. - ptobey/local-memory-mcp

For me and my family, it get really annoying having to reshare information each time you create a new LLM chat. Therefore, I decided to create local-memory-mcp, a local MCP that allows LLMs to read and write to a RAG. It also has built in systems that help keep the system efficient like a max chunk size and deprecation when information gets updated. I would love it if you all would try it out and let me know your thoughts. It's been a huge time and token saver for us. Thanks for the help and enjoy!

GitHub - ptobey/local-memory-mcp: Persistent local memory for Claude and ChatGPT. No cloud, no subscription. Self-hosted RAG over MCP. · GitHub
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
ptobey
/
local-memory-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits docs docs examples examples src src .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md config.example.json config.example.json demo.gif demo.gif docker-compose.yml docker-compose.yml requirements.txt requirements.txt run_mcp_v1_http_sse.py run_mcp_v1_http_sse.py run_mcp_v1_stdio.py run_mcp_v1_stdio.py View all files Repository files navigation
AI assistants forget everything when the conversation ends. This fixes that - locally.
No cloud. No subscription. No account. Your data stays on your machine.
local-memory-mcp gives Claude, ChatGPT, and other MCP-compatible assistants a persistent memory layer powered by local vector search (ChromaDB). Tell it something once. It remembers across sessions.
Every new Claude or ChatGPT session starts blank. Your preferences, your project context, your decisions - gone. You re-explain yourself constantly.
local-memory-mcp is a local MCP server that lets your AI assistant:
Store things worth remembering ("my deep work block is 6:30–9 AM")
Retrieve relevant context at the start of any new session
Version and supersede memories as your situation changes
It's the memory layer AI assistants should have built in, but don't.
Quickstart (Docker - 2 minutes)
git clone https://github.com/ptobey/local-memory-mcp.git
cd local-memory-mcp
docker compose up --build -d
Then point your MCP client at http://localhost:8000/mcp . Done.
→ Claude Desktop setup · ChatGPT setup · Manual Python install
[Assistant via MCP Client]
|
v
[run_mcp_v1_stdio.py | run_mcp_v1_http_sse.py]
|
v
[src/mcp_server_v1.py]
/ | \
v v v
[vector_store.py] [reconciliation.py] [health_monitor.py]
| |
v v
[Local ChromaDB] [Reconciliation Log]
Write path: store / update writes a chunk → reconciliation checks for overlap/conflict → returns warnings and self-heal hints when a write looks risky.
Read path: search runs semantic retrieval → ranking blends similarity with lightweight lexical/recency signals → deprecated chunks stay hidden unless explicitly requested.
MCP tools: store , search , update , delete , get_chunk , get_evolution_chain
Versioned updates ( strategy="version" ) with supersedes chains
Soft delete by default (history retained), optional hard delete
Heuristic reconciliation and conflict logging
Warning-first write responses with structured warnings[] and self-heal fields
Health checks for oversized chunks and unresolved conflicts
Local backup/restore for the persisted vector DB
Optional SSE auth: none (local-only), bearer , or oauth
The design idea behind it (AIX)
AIX (AI eXperience) means designing for how LLMs actually consume context, not how humans file documents:
Prefer clear text chunks over rigid document schemas
Keep metadata minimal but useful: timestamps, confidence, supersedes links, deprecation flags
Preserve history with version chains instead of destructive overwrites
Return warning-rich tool responses so the model can self-correct
The goal is practical retrieval quality and reliable AI behavior, not perfect human taxonomies.
tool: store
input: { "text": "Weekday focus block is 6:30-9:00 AM, current default schedule." }
Retrieve it later:
tool: search
input: { "query": "current deep work schedule", "top_k": 5 }
Bootstrap a new session by running a few focused retrievals, then synthesizing only active, non-deprecated chunks into a short brief for the new model instance. More flows in examples/ .
Local-first and user-controlled by default
Data stored in local ChromaDB files under the configured persist directory
No cloud backend required; optional remote access via user-managed tunneling
Never commit real secrets - use local config/env values
Integrations (Claude Desktop + ChatGPT)
Persistent local memory for Claude and ChatGPT. No cloud, no subscription. Self-hosted RAG over MCP.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
