---
source: "https://github.com/exomodel-ai/exomodel"
hn_url: "https://news.ycombinator.com/item?id=48644393"
title: "ExoModel – the object calls the LLM; you just describe what you want"
article_title: "GitHub - exomodel-ai/exomodel: Give your Python objects the ability to think and act — no glue code, no prompt engineering. · GitHub"
author: "pessoaleo"
captured_at: "2026-06-23T13:51:57Z"
capture_tool: "hn-digest"
hn_id: 48644393
score: 1
comments: 0
posted_at: "2026-06-23T13:05:07Z"
tags:
  - hacker-news
  - translated
---

# ExoModel – the object calls the LLM; you just describe what you want

- HN: [48644393](https://news.ycombinator.com/item?id=48644393)
- Source: [github.com](https://github.com/exomodel-ai/exomodel)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T13:05:07Z

## Translation

タイトル: ExoModel – オブジェクトは LLM を呼び出します。あなたはただあなたが望むものを説明するだけです
記事のタイトル: GitHub - exomodel-ai/exomodel: Python オブジェクトに考えて行動する能力を与えます。グルー コードやプロンプト エンジニアリングは必要ありません。 · GitHub
説明: Python オブジェクトに考えて行動する能力を与えます。グルー コードやプロンプト エンジニアリングは必要ありません。 - GitHub - exomodel-ai/exomodel: Python オブジェクトに考えて行動する能力を与えます。グルー コードやプロンプト エンジニアリングは必要ありません。

記事本文:
GitHub - exomodel-ai/exomodel: Python オブジェクトに考えて行動する能力を与えます。グルー コードやプロンプト エンジニアリングは必要ありません。 · GitHub
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
エクソモデル-ai
/
エクソモデル
公共
通知
通知設定を変更するにはサインインする必要があります
追加

最後のナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
18 コミット 18 コミットの例 例 src src testing testing .DS_Store .DS_Store .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff LICENSE LICENSE NOTICE NOTICE README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Python オブジェクトは自律的です。
あなたが望むものを説明してください - オブジェクトが考え、行動し、反応します。
📖 公式ドキュメント: https://exomodel.ai
📦 GitHub リポジトリ: https://github.com/exomodel-ai/exomodel
📬連絡先: contact@exomodel.ai
🚀 クイック デモ (Colab): ブラウザで ExoModel を試してください
ドメイン オブジェクトは、物事の実行方法をすでに知っています。
彼らに自然言語の命令を理解させるのに必要なものはありません
アダプター、インテント ルーター、および手書きのプロンプト テンプレート。
ExoModel は自然言語の意図とオブジェクトの動作を橋渡しするため、ドメイン オブジェクトは命令を理解し、独自のメソッドを呼び出し、保証された構造で応答できます。接着コードはありません。書き換えはありません。
今すぐインストール: pip install exomodel
従来のアプローチ
ExoModelを使用する
すべての機能の前にアダプターとパーサーを作成する
オブジェクトは命令を直接理解します - 翻訳層はありません
インテントをメソッドに手動でマップします ( iftent == X )
オブジェクトは意図に基づいて独自のメソッドを検出して呼び出します。
グルーコードはドメインロジックを埋め込みます
ドメインコードはドメインコードのまま
1 つの LLM プロバイダーに関連付けられている
ビジネスロジックに触れることなくプロバイダーを交換する
🤔 LangChain を直接使用しないのはなぜですか?
LangChain は素晴らしいです — ExoModel は実際には LangChain の上に構築されています。しかし、「LLM を持っている」と「動作するビジネス アプリケーションがある」の間にはギャップがあり、ExoModel はまさにそれを埋めます。
1. LangChain はパイプラインで考えます。あなたのビジネスはオブジェクトで考える

ts.
Proposal 、 LeadContact 、または AuditReport をモデル化するときは、チェーンではなくエンティティで考えます。 LangChain を使用すると、AI の周囲にオブジェクトをラップできます。 ExoModel は、AI をオブジェクトの内部に、それが属する場所に配置します。
2. すべてのグルー コードを記述する必要があります。
生の LangChain を使用すると、プロンプト テンプレート、出力パーサー、スキーマ検証、RAG 取得、およびツール登録をエンティティごとに手動で個別に管理できます。 ExoModel は慣例に従ってこれらすべてを処理します。スキーマを定義し、ドキュメントを添付すると、オブジェクトは何をすべきかを認識します。
3. 出力は依然としてテキストだけです。
LangChain は文字列を返します。アプリケーションには、機能する構造化され、検証され、型指定されたオブジェクト、つまりデータベースに保存したり、API に送信したり、UI でレンダリングしたり、さらなるメソッドをトリガーするために使用したりできるオブジェクトが必要です。 ExoModel の出力は常にライブ Pydantic オブジェクトです。スキーマが保証されており、解析は不要で、すぐに実行できます。
🧠 自然言語更新 — 平易な英語でオブジェクトに指示を与えます — オブジェクトは、保証された構造で独自のフィールドを更新します。パーサーも手動マッピングもありません。
📚 ネイティブ RAG グラウンディング — 実際のドキュメントでオブジェクトをグラウンディングします — PDF、URL、またはテキスト ファイルと、実際のビジネス コンテキスト内でオブジェクトの理由を添付します。
🤖 ExoAgent オーケストレーション — ルーティング ロジックを記述せずにオブジェクトを調整します — ExoAgent は意図を読み取り、 if/else ではなく、どのメソッドまたはツールを呼び出すかを決定します。
🔌 スキーマ検証済みの出力 — 人による生の入力をフィードし、スキーマ検証済みの出力を取得して、API、データベース、または UI に使用できるようにします。解析ステップはありません。
⚙️ @llm_function を使用したエージェント ツール — 任意のメソッドを 1 つのデコレーターを備えたエージェント ツールとして公開します — ExoAgent はそれを自律的に検出して呼び出します。登録も配線も必要ありません。
📊 コレクション操作 — 1 回の LLM 呼び出しでコレクション全体を操作します

— 手動でループすることなく、オブジェクトのリストを生成、更新、エクスポートします。
ExoModel は LLM に依存しません。必要なプロバイダー パッケージのみをインストールします。
pip install " exomodel[google] " # Gemini (デフォルト)
pip install " exomodel[anthropic] " # クロード
pip install " exomodel[openai] " # OpenAI / Azure OpenAI
pip install " exomodel[cohere] " # Cohere
pip install " exomodel[all] " # すべてのプロバイダー
次に、プロジェクトのルートに .env ファイルを作成します。
MY_LLM_MODEL = google_genai:gemini-2.5-フラッシュ-lite
MY_EMB_MODEL = google_genai:gemini-embedding-001 # オプション — プロバイダーから自動検出
GOOGLE_API_KEY = ここのキー
🚀 クイックスタート
pip インストール「exomodel[google]」
ステップ 2 — 継承
exomodel から ExoModel をインポート
クラス提案 (ExoModel):
クライアント : str = ""
予算: 浮動小数点 = 0.0
flagged_for_legal : bool = False
def flag_for_review (self):
自分自身。 flagged_for_legal = True
@クラスメソッド
def get_rag_sources ( cls ):
return [ "proposal_rules.md" ]
ステップ 3 — 指示する
# オブジェクトは 1 つの命令からフィールドを更新し、適切なメソッドを呼び出します
p = 提案。 create (「テスラ向けに 50,000 件の提案書を作成する」)
p 。 update (「10% の安全マージンを適用し、法的審査のためにフラグを立てる」)
# → 予算が更新され、flag_for_review() が呼び出される — if/else なし、手動ルーティングなし
print ( p . to_ui ())
# オブジェクトは、proposal_rules.md に対して自身をチェックします。
print ( p . run_analysis ())
@llm_function デコレーターは、エージェントから任意のメソッドを呼び出し可能にします。それをルール文書に指定すると、オブジェクトはそれ自体を検証します。追加のコードは必要ありません。
exomodel からインポート ExoModel 、llm_function
クラス LeadContact (ExoModel):
名前: str = ""
ステータス : str = "新規"
スコア: int = 0
@llm_function
def 修飾 ( self ):
"""会社の規模と予算のシグナルに基づいて、この潜在顧客を評価します。"""
自分自身。ステータス = "資格あり"
@llm_function
デフディスク

qualify ( self 、理由 : str ):
"""このリードを失格としてマークし、その理由を記録してください。"""
自分自身。 status = f"失格: { 理由 } "
リード = LeadContact 。 create (「Acme Corp のサラ、20 万ドルの予算について言及しました」)
リード。 master_prompt ( "このリードを評価し、適切なアクションを実行します" )
# → コンテキストに基づいて自律的に呼び出されるqualify()またはdisqualify()
🛠 建築
クラス
役割
エクソモデル
インテリジェントなオブジェクト基盤 - スキーマ主導、AI 活用、RAG 対応。これをサブクラス化してエンティティを定義します。
ExoAgent
ツール呼び出しをルーティングし、LLM コンテキストを管理し、RAG ソースを処理する推論エンジン。 ExoModel によって内部的に使用されます。直接使用することも可能です。
ExoModelList[T]
単一の LLM 呼び出しで ExoModel インスタンスを一括生成、更新、エクスポートするための型付きコレクション。
@llm_function
任意のメソッドをエージェント ツールに変換するデコレーター。実行時に master_prompt 経由で ExoAgent によって検出および呼び出し可能です。
🎯 使用例
🤝 コンサルティング アプリ — 構造化オブジェクトをリアルタイムで入力することで、複雑なプロセス (保険金請求、財務計画) をユーザーにガイドする AI アドバイザーを構築します。
🔌 エージェントミドルウェア — 人間の言語と厳格なバックエンドの橋渡しをします。すべての LLM 出力は、ネットワークに接続される前に API の仕様に正確に適合します。
📊 販売と CRM の自動化 — 提案の草案を作成し、ビジネス ルールに基づいて価格を計算し、リードのステータスを自律的に更新します。
🕵️ スマートな監査とコンプライアンス — 独自のソースコントラクトを読み取るオブジェクトを作成して、手動で監視することなく監査フィールドに値を設定し、不一致にフラグを立てます。
📈 インテリジェント ダッシュボード — 生のログまたはトランスクリプトを構造化オブジェクトのリスト ( ExoModelList ) に変換し、データ視覚化の準備を整えます。
ExoModel は Python の標準ロギング モジュールを使用します。デフォルトでは出力は表示されません。
インポートログ
#

警告とエラーのみを表示します (運用環境に推奨):
ロギング。 getLogger (「exomodel」)。 setLevel (ロギング。警告)
# すべての内部トレースを表示します (デバッグに役立ちます):
ロギング。 BasicConfig ( レベル = ロギング .DEBUG )
ロギング。 getLogger (「exomodel」)。 setLevel (ロギング.DEBUG)
ExoModel を使用すると、ドメイン コードはドメイン コードのままになります。
オブジェクトは本来意図されていたことを、よりスマートに実行します。
寄付を歓迎します! ExoModel は、開発者によって開発者のために構築されています。
機能ブランチを作成します ( git checkout -b feature/AmazingFeature )。
変更をコミットします ( git commit -m 'Add some AmazingFeature' )。
ブランチにプッシュします ( git Push Origin feature/AmazingFeature )。
Apache License 2.0 に基づいて配布されます。詳細については、「ライセンス」を参照してください。
Python オブジェクトに考えて行動する能力を与えます。グルー コードやプロンプト エンジニアリングは必要ありません。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Give your Python objects the ability to think and act — no glue code, no prompt engineering. - GitHub - exomodel-ai/exomodel: Give your Python objects the ability to think and act — no glue code, no prompt engineering.

GitHub - exomodel-ai/exomodel: Give your Python objects the ability to think and act — no glue code, no prompt engineering. · GitHub
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
exomodel-ai
/
exomodel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
18 Commits 18 Commits examples examples src src tests tests .DS_Store .DS_Store .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff LICENSE LICENSE NOTICE NOTICE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Your Python objects, autonomous.
Describe what you want — the object thinks, acts, and responds.
📖 Official Documentation: https://exomodel.ai
📦 GitHub Repository: https://github.com/exomodel-ai/exomodel
📬 Contact: contact@exomodel.ai
🚀 Quick Demo (Colab): Try ExoModel in your browser
Your domain objects already know how to do things.
Making them understand a natural language instruction shouldn't require
adapters, intent routers, and prompt templates written by hand.
ExoModel bridges natural language intent and object behavior — so your domain objects can understand instructions, call their own methods, and respond with guaranteed structure. No glue code. No rewrites.
Install now: pip install exomodel
Traditional approach
With ExoModel
Write adapters and parsers before every feature
Objects understand instructions directly — no translation layer
Map intent to methods by hand ( if intent == X )
Objects discover and call their own methods based on intent
Glue code buries your domain logic
Domain code stays domain code
Tied to one LLM provider
Swap providers without touching business logic
🤔 Why not just use LangChain directly?
LangChain is great — ExoModel is actually built on top of it. But there's a gap between "I have an LLM" and "I have a working business application", and that's exactly what ExoModel fills.
1. LangChain thinks in pipelines. Your business thinks in objects.
When you model a Proposal , a LeadContact , or an AuditReport , you're thinking in entities — not chains. LangChain makes you wrap your objects around the AI. ExoModel puts the AI inside the objects, where it belongs.
2. You still have to write all the glue code.
With raw LangChain, you manage prompt templates, output parsers, schema validation, RAG retrieval, and tool registration — separately, by hand, for every entity. ExoModel handles all of that by convention. Define the schema, attach the documents, and the object knows what to do.
3. The output is still just text.
LangChain returns strings. Your application needs structured, validated, typed objects that can act — objects you can save to a database, send to an API, render in a UI, or use to trigger further methods. ExoModel's output is always a live Pydantic object — guaranteed schema, no parsing, ready to do work.
🧠 Natural Language Updates — Give objects instructions in plain English — they update their own fields with guaranteed structure. No parsers, no manual mapping.
📚 Native RAG Grounding — Ground your objects in real documents — attach PDFs, URLs, or text files and the object reasons within your actual business context.
🤖 ExoAgent Orchestration — Let objects coordinate without writing routing logic — ExoAgent reads intent and decides which method or tool to invoke, not if/else .
🔌 Schema-Validated Output — Feed raw human input, get back schema-validated output — ready for your APIs, databases, or UI. No parsing step.
⚙️ Agentic Tools with @llm_function — Expose any method as an agent tool with one decorator — ExoAgent discovers and calls it autonomously. No registration, no wiring.
📊 Collection Operations — Operate on entire collections in one LLM call — generate, update, and export lists of objects without looping manually.
ExoModel is LLM-agnostic. Install only the provider package you need:
pip install " exomodel[google] " # Gemini (default)
pip install " exomodel[anthropic] " # Claude
pip install " exomodel[openai] " # OpenAI / Azure OpenAI
pip install " exomodel[cohere] " # Cohere
pip install " exomodel[all] " # all providers
Then create a .env file at the root of your project:
MY_LLM_MODEL = google_genai:gemini-2.5-flash-lite
MY_EMB_MODEL = google_genai:gemini-embedding-001 # optional — auto-detected from provider
GOOGLE_API_KEY = your-key-here
🚀 Quick Start
pip install " exomodel[google] "
Step 2 — Inherit
from exomodel import ExoModel
class Proposal ( ExoModel ):
client : str = ""
budget : float = 0.0
flagged_for_legal : bool = False
def flag_for_review ( self ):
self . flagged_for_legal = True
@ classmethod
def get_rag_sources ( cls ):
return [ "proposal_rules.md" ]
Step 3 — Instruct
# The object updates fields and calls the right method — from one instruction
p = Proposal . create ( "Draft a 50k proposal for Tesla" )
p . update ( "apply the 10% safety margin and flag it for legal review" )
# → budget updated, flag_for_review() called — no if/else, no manual routing
print ( p . to_ui ())
# The object checks itself against proposal_rules.md
print ( p . run_analysis ())
The @llm_function decorator makes any method callable by the agent. Point it at a rules document and the object validates itself — no extra code.
from exomodel import ExoModel , llm_function
class LeadContact ( ExoModel ):
name : str = ""
status : str = "new"
score : int = 0
@ llm_function
def qualify ( self ):
"""Qualify this lead based on company size and budget signals."""
self . status = "qualified"
@ llm_function
def disqualify ( self , reason : str ):
"""Mark this lead as disqualified and record why."""
self . status = f"disqualified: { reason } "
lead = LeadContact . create ( "Sarah from Acme Corp, mentioned $200k budget" )
lead . master_prompt ( "Evaluate this lead and take the right action" )
# → qualify() or disqualify() called autonomously based on context
🛠 Architecture
Class
Role
ExoModel
The intelligent object foundation — schema-driven, AI-powered, RAG-aware. Subclass this to define your entities.
ExoAgent
The reasoning engine that routes tool calls, manages LLM context, and processes RAG sources. Used internally by ExoModel ; also available for direct use.
ExoModelList[T]
Typed collection for bulk generation, updating, and export of ExoModel instances in a single LLM call.
@llm_function
Decorator that turns any method into an agentic tool, discoverable and callable by ExoAgent at runtime via master_prompt .
🎯 Use Cases
🤝 Consultative Apps — Build AI advisors that guide users through complex processes (insurance claims, financial planning) by populating structured objects in real time.
🔌 Agentic Middleware — Bridge human language and rigid backends. Every LLM output fits your API's exact specification before it hits the wire.
📊 Sales & CRM Automation — Draft proposals, calculate pricing against business rules, and update lead status autonomously.
🕵️ Smart Auditing & Compliance — Create objects that read their own source contracts to populate audit fields and flag inconsistencies without manual oversight.
📈 Intelligent Dashboarding — Transform raw logs or transcripts into lists of structured objects ( ExoModelList ), ready for data visualization.
ExoModel uses Python's standard logging module. No output is shown by default.
import logging
# Show warnings and errors only (recommended for production):
logging . getLogger ( "exomodel" ). setLevel ( logging . WARNING )
# Show all internal traces (useful for debugging):
logging . basicConfig ( level = logging . DEBUG )
logging . getLogger ( "exomodel" ). setLevel ( logging . DEBUG )
With ExoModel, your domain code stays domain code.
Your objects do what they were always meant to do — just smarter.
We welcome contributions! ExoModel is built by developers for developers.
Create your Feature Branch ( git checkout -b feature/AmazingFeature ).
Commit your changes ( git commit -m 'Add some AmazingFeature' ).
Push to the branch ( git push origin feature/AmazingFeature ).
Distributed under the Apache License 2.0. See LICENSE for more information.
Give your Python objects the ability to think and act — no glue code, no prompt engineering.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
