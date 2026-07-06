---
source: "https://github.com/forgedlogicdev/pyforge-memory"
hn_url: "https://news.ycombinator.com/item?id=48807937"
title: "Pyforge-memory – three-tier memory for AI agents that works"
article_title: "GitHub - forgedlogicdev/pyforge-memory · GitHub"
author: "ForgedLogicdev"
captured_at: "2026-07-06T17:46:45Z"
capture_tool: "hn-digest"
hn_id: 48807937
score: 1
comments: 0
posted_at: "2026-07-06T17:37:26Z"
tags:
  - hacker-news
  - translated
---

# Pyforge-memory – three-tier memory for AI agents that works

- HN: [48807937](https://news.ycombinator.com/item?id=48807937)
- Source: [github.com](https://github.com/forgedlogicdev/pyforge-memory)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T17:37:26Z

## Translation

タイトル: Pyforge-memory – AI エージェントが動作する 3 層メモリ
記事のタイトル: GitHub - forgedlogicdev/pyforge-memory · GitHub
説明: GitHub でアカウントを作成して、forgedlogicdev/pyforge-memory 開発に貢献します。

記事本文:
GitHub - forgedlogicdev/pyforge-memory · GitHub
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
鍛造ロジック開発
/
pyforge-メモリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらに開くアクション

ns メニューのフォルダーとファイル
2 コミット 2 コミット .github .github pyforge_memory pyforge_memory .gitignore .gitignore LAUNCH_POST.md LAUNCH_POST.md ライセンス ライセンス README.md README.md Index.html Index.html pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用の 3 層メモリ。コンテキストをクリーンで関連性があり、管理しやすい状態に保ちます。LLM が考える前に 3000 文字を超える日記に溺れることはもうありません。
Forged Logic の本番環境で実戦テスト済み。
ほとんどの AI 記憶ツールは、すべての会話をコンテキスト ウィンドウにダンプします。 50 回の交換の後、エージェントの ID プロンプトは表示されずに切り詰められ、一般的なアシスタントの動作に戻ります。
pyforge-memory は代わりに 3 つの層を使用します。
コンテキストの合計は約 3000 文字から約 1000 文字に減少し、すべての文字が適切な文字になります。
3 層アーキテクチャ — 逐語的な最新性、キーワードの関連性、圧縮された履歴
ポイズン フィルタリング — 一般的なアシスタント スクリプトの出力がメモリを汚染する前に検出してブロックします。
構成可能 — 独自の禁止用語、ストップワード、メモリ制限を設定します
ファイルベース — JSONL ストレージ、データベースなし、API キーなし、どこでも実行可能
ナレッジ ディレクトリ — テキスト ファイル全体でドメイン コンテキストをキーワード検索します。
フレームワークに依存しない — あらゆる LLM API (Ollama、OpenAI、Claude、ローカル) で動作します
pip インストール pyforge-memory
クイックスタート
pyforge_memoryからMemoryEngineをインポート
# ストレージパスを使用して初期化する
エンジン = メモリエンジン (
Memory_dir = "./my_agent/memory" ,
Knowledge_dir = "./my_agent/knowledge"
）
# LLM 呼び出し可能 (メッセージを受け取ってテキストを返す任意の関数)
def llm_call (メッセージ):
# ... ここで LLM を呼び出します ...
応答を返す
# 各ユーザーメッセージのコンテキストを構築する
メッセージ = エンジン 。 build_context (
user_message = "データベース スキーマについて何を話し合いましたか?" 、
システムプロンプト = "Y

あなたは役に立つコーディングアシスタントです。" ,
query_fn = llm_call 、
）
# LLM に送信する
応答 = llm_call (メッセージ)
# 交換を保存する
エンジン。 save_exchange ( "データベース スキーマについて何を話し合いましたか?" , response )
仕組み
ユーザーメッセージが到着
│
▼
┌─────────┐
│ 1. 抽出 │ 意味のあるキーワードを取り出し、ストップワードを取り除く
│ キーワード │
━━━━┬───────┘
│
▼
┌─────────┐
│ 2. 逐語的 │ 最新 N 回の生の交換、ポイズン フィルタリング、
│ レイヤー │ 重複に近い検出
━━━━┬───────┘
│
▼
┌─────────┐
│ 3. キーワード │ 記憶+知識ファイルの検索
│ レイヤー │ キーワード一致、スコア、ランク用
━━━━┬───────┘
│
▼
┌─────────┐
│ 4. ダイジェスト │ 昔の会話を圧縮してまとめたもの
│ レイヤー │ メモリ > 10 エントリの場合に自動生成
━━━━┬───────┘
│
▼
┌─────────┐
│ 5. 組み立てる │ [システム] + 逐語 + キーワード + ダイジェスト
│ コンテキスト │ + [ユーザー] → LLM の準備ができました
━━━━━━━━┘
毒検出のカスタマイズ
エンジン = メモリエンジン (
禁止用語 = [
「私はAIアシスタントです」、
"言語モデルとして" ,
# 独自のパターンを追加する
]
）
# テキストをチェックする
エンジンがあれば。 is_poisoned ( "AI アシスタントとして、私はできません..." ):
print (「捕まえた!」)
知識ベース
テキスト ファイルをナレッジ ディレクトリにドロップします。関連する場合はキーワード検索されます。
私のエージェント/知識/
§── データベーススキーマ.tx

t
§── api-docs.txt
└── コーディング規約.txt
なぜ 3 層なのでしょうか?
逐語的にのみ → 古い会話の文脈が失われる
キーワードだけ → 現在の会話の流れを失う
ダイジェストのみ → 圧縮されすぎて詳細が欠落します
3 つの階層: 最新性 + 関連性 + 圧縮。
これがプロジェクトに役立つ場合は、ヒントを歓迎します。
GyhqwvFv4ZX3co7T2tziLZTU1vMJoYjfXXS2dpG3FPXsc
Solana Explorer で見る
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

Contribute to forgedlogicdev/pyforge-memory development by creating an account on GitHub.

GitHub - forgedlogicdev/pyforge-memory · GitHub
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
forgedlogicdev
/
pyforge-memory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github .github pyforge_memory pyforge_memory .gitignore .gitignore LAUNCH_POST.md LAUNCH_POST.md LICENSE LICENSE README.md README.md index.html index.html pyproject.toml pyproject.toml View all files Repository files navigation
Three-tier memory for AI agents. Keeps context clean, relevant, and manageable — no more drowning your LLM in 3000+ characters of diary before it can think.
Battle-tested in production at Forged Logic .
Most AI memory tools dump every conversation into the context window. After 50 exchanges, your agent's identity prompt gets silently truncated and it reverts to generic assistant behavior.
pyforge-memory uses three tiers instead:
Total context drops from ~3000 characters to ~1000 — and every character is the right character.
Three-tier architecture — verbatim recency, keyword relevance, compressed history
Poison filtering — detects and blocks generic assistant-script output before it poisons memory
Configurable — set your own forbidden terms, stop words, memory limits
File-based — JSONL storage, no database, no API keys, runs anywhere
Knowledge directory — keyword-search across text files for domain context
Framework agnostic — works with any LLM API (Ollama, OpenAI, Claude, local)
pip install pyforge-memory
Quick start
from pyforge_memory import MemoryEngine
# Initialize with storage paths
engine = MemoryEngine (
memory_dir = "./my_agent/memory" ,
knowledge_dir = "./my_agent/knowledge"
)
# Your LLM callable (any function that takes messages and returns text)
def llm_call ( messages ):
# ... call your LLM here ...
return response
# Build context for each user message
messages = engine . build_context (
user_message = "what did we discuss about the database schema?" ,
system_prompt = "You are a helpful coding assistant." ,
query_fn = llm_call ,
)
# Send to your LLM
response = llm_call ( messages )
# Save the exchange
engine . save_exchange ( "what did we discuss about the database schema?" , response )
How it works
User message arrives
│
▼
┌──────────────────┐
│ 1. Extract │ Pull meaningful keywords, strip stop words
│ keywords │
└──────┬───────────┘
│
▼
┌──────────────────┐
│ 2. Verbatim │ Last N raw exchanges, poison-filtered,
│ layer │ near-duplicate detection
└──────┬───────────┘
│
▼
┌──────────────────┐
│ 3. Keyword │ Search memory + knowledge files
│ layer │ for keyword matches, score and rank
└──────┬───────────┘
│
▼
┌──────────────────┐
│ 4. Digest │ Compressed summary of old conversations
│ layer │ auto-generated when memory > 10 entries
└──────┬───────────┘
│
▼
┌──────────────────┐
│ 5. Assemble │ [system] + verbatim + keyword + digest
│ context │ + [user] → ready for LLM
└──────────────────┘
Customizing poison detection
engine = MemoryEngine (
forbidden_terms = [
"i am an ai assistant" ,
"as a language model" ,
# Add your own patterns
]
)
# Check any text
if engine . is_poisoned ( "As an AI assistant, I cannot..." ):
print ( "Caught it!" )
Knowledge base
Drop text files into your knowledge directory. They're keyword-searched when relevant:
my_agent/knowledge/
├── database-schema.txt
├── api-docs.txt
└── coding-conventions.txt
Why three tiers?
Verbatim alone → loses context from older conversations
Keyword alone → loses the flow of the current conversation
Digest alone → too compressed, misses specifics
Three tiers together: recency + relevance + compression.
If this helped your project, tips are welcome:
GyhqwvFv4ZX3co7T2tziLZTU1vMJoYjfJx2dpG3FPXsc
View on Solana Explorer
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
