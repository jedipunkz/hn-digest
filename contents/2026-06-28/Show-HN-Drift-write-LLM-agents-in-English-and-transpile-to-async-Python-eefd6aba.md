---
source: "https://github.com/rileyq7/drift"
hn_url: "https://news.ycombinator.com/item?id=48711629"
title: "Show HN: Drift, write LLM agents in English and transpile to async Python"
article_title: "GitHub - rileyq7/drift: An intent-based language for agentic systems. Write agents in English. Transpile to async Python. · GitHub"
author: "rileyq12"
captured_at: "2026-06-28T21:22:20Z"
capture_tool: "hn-digest"
hn_id: 48711629
score: 2
comments: 0
posted_at: "2026-06-28T21:04:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Drift, write LLM agents in English and transpile to async Python

- HN: [48711629](https://news.ycombinator.com/item?id=48711629)
- Source: [github.com](https://github.com/rileyq7/drift)
- Score: 2
- Comments: 0
- Posted: 2026-06-28T21:04:32Z

## Translation

タイトル: Show HN: Drift、LLM エージェントを英語で記述し、非同期 Python にトランスパイルする
記事のタイトル: GitHub - rileyq7/drift: エージェント システム用のインテント ベースの言語。エージェントを英語で書きます。非同期 Python にトランスパイルします。 · GitHub
説明: エージェント システム用のインテントベースの言語。エージェントを英語で書きます。非同期 Python にトランスパイルします。 - ライリーq7/ドリフト

記事本文:
GitHub - rileyq7/drift: エージェント システム用のインテントベースの言語。エージェントを英語で書きます。非同期 Python にトランスパイルします。 · GitHub
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
ライリーq7
/
ドリフト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

イオンオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
30 コミット 30 コミット .github/ workflows .github/ workflows apps/ Web Apps/ Web docs docs ドリフト ドリフトの例 サンプル テスト テスト vscode-drift vscode-drift .env.example .env.example .gitignore .gitignore ライセンス ライセンス LLM.md LLM.md README.md README.md RELEASE_CHECKLIST.md RELEASE_CHECKLIST.mdデモ.pyデモ.pydrift-dendric-spec.htmldrift-dendric-spec.htmldrift_cli.pydrift_cli.py pyproject.toml pyproject.toml pytest.ini pytest.ini すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェント システム用のインテントベースの言語。英語の形をしたブロックでエージェントを作成し、非同期 Python として実行します。
エージェントの受信箱トリアージ {
モデル：「gpt-5.4-nano」
予算: 実行あたり 0.10 ドル
ステップトリアージ(emails: list<string>) -> list<EmailAnalysis> {
分析してみましょう = []
メール内の各メールに対して並行して {
let Analysis = 電子メールを EmailAnalysis として分類する
分析.追加(分析)
}
分析の各 a について {
if a.priority == "緊急" {
「緊急 {a.subject}: {a.summary}」と応答します
}
}
リターン分析
}
}
完全なエージェント: モデルの選択、予算、並列ファンアウト、構造化された分類、条件付き出力。トランスパイラーは、Drift のシン ランタイムで実行される非同期 Python を出力します。 OpenAI との対戦: 5 つの電子メール、1.82 秒、0.0092 ドル、5 つの型付きデータクラスが返されました。
pip インストール ドリフト言語
オプションの追加物:
pip install "drift-lang[mcp]" # MCP ツールのサポート
pip install "drift-lang[dendric] " # Dendric メモリ バックエンド
pip install "drift-lang[all]"
最初のエージェントまで 30 秒
ドリフトニューハロー
CD こんにちは
ドリフト ラン hello.drift --input ' {"name":"Riley"} '
API キーは必要ありません。 Drift はモックプロバイダーにフォールバックするため、何かがすぐに機能することがわかります。実際のモデルを使用するには、ANTHROPIC_API_KEY または OPENAI_API_KEY を .env にドロップします。
ドリフト 新しい <名前> 開始の足場

rプロジェクト
drift run <file.drift> トランスパイルして実行
ドリフトチェック <file.drift> 構文を検証します
ドリフトトランスパイル <ファイル> Python の出力 (ファイルに書き込むには -o を使用します)
ドリフトレックス/解析デバッグツール
言語には何があるのか
エージェント : 最上位ユニット。モデル、バジェット、状態、メモリ、およびステップがあります。
step : 入力されたサブプロシージャ。本体は一連の宣言文です。
インテント動詞: summary、extract、classify、translate、match、generate など。それぞれが型指定された LLM 呼び出しになります。
confident<T> : 信頼度ゲート分岐。確かな場合は安価なパスを実行し、そうでない場合はエスカレーションします。
モデル { … } : 信頼度 < 0.7 の場合はprefer、fallback、アップグレード、ストリームは「高速」、次に「低速」を使用したマルチプロバイダー ルーティング。
python|mcp|rest からのツール名: 外部ツールを宣言します。 MCP は公式 SDK に対して実行されます。
Pipeline : -> 、 => 、 ~> 、 |> 演算子を使用したコンポーザブル フロー。
xs 並列の各 x に対して: asyncio.gather の下にあります。
試行 / 回復 : 再試行、失敗、および名前付きアームを使用した構造化されたエラー処理。
メモリ : 短期スクラッチパッドまたは耐久性のあるバックエンド (Dendric)。思い出す、思い出す、既視感、忘れる。
動詞の定義 : 独自の型付き動詞を使用してインテント語彙を拡張します。
クロスエージェント呼び出し: OtherAgent.step(args) は正常に機能します。
ファイル
のために
LLM.md
コーディング エージェント (クロード、カーソル、コパイロット): ワンショット ロードの完全なリファレンス
docs/言語.md
ドリフトを学ぶ人間
docs/cookbook.md
パターンのコピー＆ペースト
docs/gotchas.md
よくある間違い
例
動作する .drift プログラムとその生成された Python については、examples/ を参照してください。
confident_demo.drift :confident<T> 分岐
Grant_checker.drift : エンドツーエンドのインテント + 構造化されたリターン
inbox_sorter.drift : 各 … 並列トリアージ
inbox_triage_live.drift : 正規の 30 行デモ (本物の LLM 検証済み: 5 通の電子メール、1.82 秒、0.0092 ドル)
Grant_checker_with_memory.drift : デンドリックバック

長期記憶
Grant_checker_compare.drift : 引用防止メモリ。 Run 2 の LLM 推論では Run 1 を名前で引用し、並べて比較します。
アルファ。言語表面は安定しており、ランタイムは動作し、352/352 テストに合格しています。厳密な JSON 出力、MCP ツール、Dendric メモリ、ソースマップされたランタイム エラーを備えた OpenAI + Anthropic プロバイダー。音声プリミティブは解析されますが、アダプターはまだ接続されていません。型システムBeyondconfident<T>はロードマップ上にあります。
エージェント システム用のインテントベースの言語。エージェントを英語で書きます。非同期 Python にトランスパイルします。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.1 — 最初のパブリックアルファ版
最新の
2026 年 6 月 28 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An intent-based language for agentic systems. Write agents in English. Transpile to async Python. - rileyq7/drift

GitHub - rileyq7/drift: An intent-based language for agentic systems. Write agents in English. Transpile to async Python. · GitHub
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
rileyq7
/
drift
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
30 Commits 30 Commits .github/ workflows .github/ workflows apps/ web apps/ web docs docs drift drift examples examples tests tests vscode-drift vscode-drift .env.example .env.example .gitignore .gitignore LICENSE LICENSE LLM.md LLM.md README.md README.md RELEASE_CHECKLIST.md RELEASE_CHECKLIST.md demo.py demo.py drift-dendric-spec.html drift-dendric-spec.html drift_cli.py drift_cli.py pyproject.toml pyproject.toml pytest.ini pytest.ini View all files Repository files navigation
An intent-based language for agentic systems. Write your agent in English-shaped blocks, run it as async Python.
agent InboxTriage {
model: "gpt-5.4-nano"
budget: $0.10 per run
step triage(emails: list<string>) -> list<EmailAnalysis> {
let analyses = []
for each email in emails parallel {
let analysis = classify email as EmailAnalysis
analyses.add(analysis)
}
for each a in analyses {
if a.priority == "urgent" {
respond "URGENT {a.subject}: {a.summary}"
}
}
return analyses
}
}
A full agent: model choice, budget, parallel fan-out, structured classification, conditional output. The transpiler emits async Python that runs on Drift's thin runtime. Live against OpenAI: 5 emails, 1.82s, $0.0092, returned 5 typed dataclasses.
pip install drift-lang
Optional extras:
pip install " drift-lang[mcp] " # MCP tool support
pip install " drift-lang[dendric] " # Dendric memory backend
pip install " drift-lang[all] "
30 seconds to your first agent
drift new hello
cd hello
drift run hello.drift --input ' {"name":"Riley"} '
No API key required. Drift falls back to a mock provider so you see something work immediately. Drop an ANTHROPIC_API_KEY or OPENAI_API_KEY into .env to use a real model.
drift new <name> Scaffold a starter project
drift run <file.drift> Transpile and execute
drift check <file.drift> Validate syntax
drift transpile <file> Emit Python (use -o to write to a file)
drift lex / parse Debug tooling
What's in the language
agent : top-level unit. Has model , budget , state , memory , and step s.
step : typed sub-procedure. Body is a sequence of declarative statements.
Intent verbs : summarize , extract , classify , translate , match , generate , etc. Each one becomes a typed LLM call.
confident<T> : confidence-gated branching. Run the cheap path when sure, escalate when not.
model { … } : multi-provider routing with prefer , fallback , upgrade when confidence < 0.7 , and stream "fast" then "slow" .
tool name from python|mcp|rest : declare external tools. MCP runs against the official SDK.
pipeline : composable flow with -> , => , ~> , |> operators.
for each x in xs parallel : asyncio.gather underneath.
attempt / recover : structured error handling with retry, fail, and named arms.
memory : short-term scratchpad or durable backend (Dendric). remember , recall , deja_vu , forget .
define verb : extend the intent vocabulary with your own typed verbs.
Cross-agent calls : OtherAgent.step(args) just works.
File
For
LLM.md
Coding agents (Claude, Cursor, Copilot): complete reference for one-shot loading
docs/language.md
Humans learning Drift
docs/cookbook.md
Copy-paste patterns
docs/gotchas.md
Common mistakes
Examples
See examples/ for working .drift programs and their generated Python:
confident_demo.drift : confident<T> branching
grant_checker.drift : end-to-end intent + structured return
inbox_sorter.drift : for each … parallel triage
inbox_triage_live.drift : the canonical 30-line demo (real-LLM verified: 5 emails, 1.82s, $0.0092)
grant_checker_with_memory.drift : Dendric-backed long-term memory
grant_checker_compare.drift : citation-proof memory. Run 2's LLM reasoning cites Run 1 by name and makes side-by-side comparisons
Alpha. Language surface is stable, runtime works, 352/352 tests passing. OpenAI + Anthropic providers with strict-JSON output, MCP tools, Dendric memory, source-mapped runtime errors. Voice primitives parse but adapters aren't wired yet. Type system beyond confident<T> is on the roadmap.
An intent-based language for agentic systems. Write agents in English. Transpile to async Python.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.1 — first public alpha
Latest
Jun 28, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
