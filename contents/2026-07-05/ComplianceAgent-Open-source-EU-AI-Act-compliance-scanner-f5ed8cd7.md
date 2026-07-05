---
source: "https://github.com/latreon/compliance-agent"
hn_url: "https://news.ycombinator.com/item?id=48793704"
title: "ComplianceAgent: Open-source EU AI Act compliance scanner"
article_title: "GitHub - latreon/compliance-agent: Scan your codebase for EU AI Act compliance in seconds. Detects AI usage (OpenAI, Anthropic, LangChain), flags gaps, gives copy-paste fixes. · GitHub"
author: "latreon"
captured_at: "2026-07-05T13:08:56Z"
capture_tool: "hn-digest"
hn_id: 48793704
score: 1
comments: 0
posted_at: "2026-07-05T12:26:33Z"
tags:
  - hacker-news
  - translated
---

# ComplianceAgent: Open-source EU AI Act compliance scanner

- HN: [48793704](https://news.ycombinator.com/item?id=48793704)
- Source: [github.com](https://github.com/latreon/compliance-agent)
- Score: 1
- Comments: 0
- Posted: 2026-07-05T12:26:33Z

## Translation

タイトル: ComplianceAgent: オープンソースの EU AI 法コンプライアンス スキャナー
記事のタイトル: GitHub - latereon/compliance-agent: コードベースを数秒でスキャンして EU AI 法に準拠しているかどうかを確認します。 AI の使用状況 (OpenAI、Anthropic、LangChain) を検出し、ギャップにフラグを立て、コピーアンドペーストの修正を行います。 · GitHub
説明: EU AI 法に準拠しているかどうかコードベースを数秒でスキャンします。 AI の使用状況 (OpenAI、Anthropic、LangChain) を検出し、ギャップにフラグを立て、コピーアンドペーストの修正を行います。 - Latreon/コンプライアンスエージェント

記事本文:
GitHub - latreon/compliance-agent: コードベースを数秒でスキャンして、EU AI 法に準拠しているかどうかを確認します。 AI の使用状況 (OpenAI、Anthropic、LangChain) を検出し、ギャップにフラグを立て、コピーアンドペーストの修正を行います。 · GitHub
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
ラトレオン
/
コンプライアンスエージェント
公共
通知
君はね

サインインして通知設定を変更できない
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット .github .github docs docs サンプル サンプル ルール ルール src/準拠_エージェント src/準拠_エージェント テンプレート テンプレート テスト テスト .gitignore .gitignore .pre-commit-hooks.yaml .pre-commit-hooks.yaml .python-version .python-version CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI プロジェクトが EU の規則に従っているかどうかを確認してください。
EU には AI に関する新しいルールがあります。 OpenAI、Anthropic、LangChain を使用して構築している場合、
または AI フレームワークを使用する場合は、準拠しているかどうかを確認する必要があります。このツールはそれを実現します
あなた — 1 つのコマンド、約 5 秒。
30 秒スタート · 内容 · 仕組み · 例 · すべてのコマンド · FAQ
# インストール (分離された CLI ツール)
UV ツールのインストール コンプライアンス エージェント
＃紫外線がない？使用: pipx インストール コンプライアンス エージェント
# プロジェクトを確認してください
コンプライアンス エージェント スキャン。
＃ 終わり。見つかったものを読んでください。
機能 (簡易バージョン)
コードをスキャンします — AI (OpenAI、LangChain など) が使用されている場所を見つけます。
ルールをチェックします — コードを EU AI 法の要件と比較します。
何が欠けているのかを示し、修正する必要があるものを正確に示します。
コードを提供します — 各問題に対するコピー＆ペーストによる修正を提供します。
コンプライアンス エージェント スキャンを実行するとき。 、ボックス化された端末レポートを取得します。
ヘッダー、要約ストリップ、記事ごとの範囲、調査結果、およびギャップ
修正します。形状の例 (バンドルされたサンプルの実際の出力は次のとおりです)
例/EXPECTED_OUTPUT.md ):
╭─ EU AI 法遵守報告書 ───────

───────╮
│ スキャンされたファイル 3 │
│ リスク層 限定的 │
╰─────────────────── コンプライアンスエージェント ╯
╭─ スキャンの概要 ──────── スキャン概要──────────────╮
│ 3 1 6 9 │
│ ファイル AI システムのギャップ発見 │
╰───────────────────────╯
╭─ コンプライアンスのギャップ ─────────────────────╮
│ ✗ 不足 AI イベントの自動ログ記録が必要 (第 12 条) │
│ 修正: templates/art12/event_logging.py │
│ ✗ 欠落している AI の透明性開示が必要 (第 50 条) │
│ 修正: templates/art50/transparency_notice.py │
│ ✗ MISSING AI 呼び出しに関するエラー処理 (第 15 条) │
│ 修正: Try/Except + モデル呼び出しの周囲にフォールバックを追加 │
╰───────────────────────╯
次: コンプライアンスエージェントの推奨 。 --output ./fixes
ファイルの方がいいですか? scan --format pdf --output report.pdf は PDF を書き込み、
別のレポート コマンドはマークダウンまたは PDF をディスクに書き込みます。スキャンの場合、
--format マークダウンと --format json をターミナル/stdout にレンダリングします — パイプ
必要に応じて、json をファイルにコピーします (「コマンド リファレンス」を参照)。
OpenAI、Anthropic、Google、または任意の AI API を使用する
チャットボットまたは AI アシスタントを構築する
LangChain、CrewAI、AutoGen、または LangGraph を使用する
AIの導入

EU 内または EU ユーザーにサービスを提供する
罰金（ほとんどの義務について売上高の 3% あたり最大 1,500 万ユーロ）を回避したい
禁止行為に対して最大 3,500 万ユーロ / 7%)
AI は個人的なプロジェクト（ビジネスではない）にのみ使用してください
EU 内で運営したり、EU 内のユーザーにサービスを提供したりしないでください
ComplianceAgent はコマンドライン ツールであるため、これをインストールする最もクリーンな方法は次のとおりです。
ツール インストーラーを使用して、それを独自の隔離された環境に保持します。
推奨 (分離 CLI インストール)
UV ツールのインストール コンプライアンス エージェント
または、pipx を使用して:
pipx インストール コンプライアンス エージェント
uv や pipx はまだありませんか? 1 つインストールします。
brew install uv # または: brew install pipx
代替案: 仮想環境内で pip を実行する
最新の macOS/Linux では、システム Python へのベア pip インストールがブロックされます
(PEP 668、「外部管理環境」)。仮想環境を使用します。
python3 -m venv .venv
ソース .venv/bin/activate # Windows: .venv\Scripts\activate
pip インストール コンプライアンス エージェント
最新の未リリースバージョン (GitHub から)
uv ツールをインストール git+https://github.com/latreon/compliance-agent.git
# または: pipx install git+https://github.com/latreon/compliance-agent.git
機能したことを確認する
コンプライアンスエージェントのバージョン
# コンプライアンスエージェント v0.1.7
インストールまたは実行に問題がありますか?トラブルシューティング ガイドを参照してください。
スキャナーはプロジェクト ファイルを読み取り、AI 関連のパターンを探します。
import openai — OpenAI を使用しています
langchain から — LangChain を使用しています
AgentExecutor() — AI エージェントを実行しています
client.chat.completions.create() — AI API を呼び出しています
プロバイダーとフレームワークの検出では、(テキスト検索だけでなく) AST 解析が使用されます。
実際にライブラリをインポートするファイル内でのみ起動されるため、次のようなコメントが必要です。
「OpenAI」はプロバイダーの検出をトリガーしないと述べています。軽量のキーワード パターン
(チャット インターフェイスの文言など) さらにドキュメントと設定をスキャンします。
少数の発見が得られる可能性があります

.md ファイル。
ツールは、検出した内容に基づいてリスク レベルを割り当てます。
ティアの決め方。 EU AI 法に基づく高いリスクは、
技術的なものではなく、ユースケース ドメイン (付録 III)
能力。ツールを備えた自律エージェントが自動的に高リスクになるわけではありません。
履歴書審査や信用スコアリングシステムです。 ComplianceAgent が「高」に分類
Annex III ドメインインジケータを検出した場合のみ、UNACCEPTABLE の場合のみ
可能性のあるアートを検出します。 5 禁止されている行為 (例: ソーシャルスコアリング、ターゲットを絞らない行為)
顔の擦り傷）。ドメイン インジケーターはファイル パスおよびコードと照合されます
コンテンツ（ファイル名だけでなく）も含めますが、実際に AI を使用するプロジェクトのみが対象となります。
どちらもキーワードベースのヒューリスティックであり、暫定的なものです（第 6 条第 3 条でも一部を免除されています）
狭い目的のシステム）のため、一致する場合は自己評価と相談を促すことになります。
弁護士 — 法的な決定ではありません。参照
階層化の方法については docs/ARCHITECTURE.md
決まっている。
このツールは、EU AI 法の 13 の特定の条項をチェックします。
見つかった問題ごとに、ツールは次のことを行います。
どのルールに修正が必要かを示します
コピーできるコードテンプレートを提供します
どこに置くべきかを正確に教えてくれる
問題: 「AI と話しています」という通知がない
規則: EU AI 法第 50 条(1)
修正: templates/art50/transparency_notice.py をプロジェクトにコピーします
場所: チャットエンドポイントの前に追加します
実際の例
例 1: 単純なチャットボット (リスクは限定的)
# チャットボット.py
輸入オープンアイ
クライアント = openai 。オープンAI ()
デフォルトチャット ( user_input ):
クライアントを返す。チャット 。完成品。作成(
モデル = "gpt-4" 、
メッセージ = [{ "ロール" : "ユーザー" , "コンテンツ" : user_input }],
）。選択肢 [ 0 ]。メッセージ 。内容
スキャン結果 (概要の説明 - 正確な文言/カウントはスキャンから得られます):
リスク層: 制限あり — ユーザー向け AI、Annex III の高リスク ドメインが一致しない
ギャップ：アート。 12 (記録保持)、第 12 条50（透明）、Art. 15（堅牢性）
修正

: 透明性通知 + イベントログ + エラー処理を追加します。
例 2: LangChain エージェント (リスクが高い)
Web を検索して電子メールを送信できるエージェント:
# エージェント.py
ラングチェーンから。エージェントが AgentExecutor をインポートする
ラングチェーンから。ツールインポートツール
ツール = [
ツール ( name = "検索" 、 func = search_web 、説明 = "Web を検索" )、
ツール ( name = "電子メール" 、 func = send_email 、説明 = "電子メールを送信する" )、
】
executor = AgentExecutor (エージェント = エージェント、ツール = ツール)
スキャン結果 (例: コードが一致する場合にのみ、レベルが高くなります)
付属書 III 高リスクドメイン。エージェントのパターン自体が見落としを引き起こす
以下のギャップをログに記録します):
リスク: 限定的 (Annex III の高リスクドメインは検出されません)
フレームワーク: LangChain (エージェント、ツール)
問題: いくつかの問題を含む
1. ツールの使用前に人間による監視が行われない (第 14 条)
2. ツール呼び出しのログを記録しない (第 12 条)
3. API エラーに対するエラー処理の禁止 (第 15 条)
4. 「あなたはAIと話しています」という通知をしない（第50条）
修正: 人間参加型、ロギング、エラー処理、透明性を追加します。
注: ツールにアクセスすると、実際の Art が上昇します。 14 監督とアート。 12 のロギングギャップがありますが、
それ自体でシステムのリスクが高くなるわけではありません。ティアがHIGHになります
エージェントが付録 III ドメインで活動している場合のみ (例: 雇用、クレジット、
生体認証) — 階層の決定方法を参照してください。
エージェントのチームが調査し、執筆しています:
# クルー.py
from crewai import エージェント、タスク、クルー
研究者 = エージェント (役割 = "研究者" 、ツール = [検索])
ライター = エージェント (役割 = "ライター" 、ツール = [書き込み])
乗組員 = 乗組員 (
エージェント = [ 研究者 、 ライター ]、
タスク = [ タスク (説明 = "研究" 、エージェント = 研究者 )、
タスク ( 説明 = "書き込み" 、エージェント = ライター )]、
)
乗組員。キックオフ（）
スキャン結果 (例):
リスク: 限定的 (マルチエージェントだが、Annex III の高リスクドメインは検出されない)
フレームワーク: CrewAI (

紳士、乗務員、タスク)
問題: いくつかの問題を含む
1. 乗組員処刑前の監督の禁止（第 14 条）
2. エージェントのアクションをログに記録しない (第 12 条)
3. 技術文書の禁止 (第 11 条)
修正: 承認ワークフロー、ログ記録、およびドキュメントを追加します。
研究および執筆を行う乗組員は、それ自体では附属書 III の高リスクではありません。
ユースケースのため、階層は LIMITED のままです。履歴書審査で同じスタッフを指名する
またはクレジットの決定により、ティアが HIGH になります。
# フォルダーをスキャンします (. は現在のフォルダーを意味します)
コンプライアンス エージェント スキャン。
# 出力タイプ
コンプライアンス エージェント スキャン。 --format マークダウン # 読み取り用 (デフォルト)、ターミナルへ
コンプライアンス エージェント スキャン。 --コンピューター/CI の json # を標準出力にフォーマットします
コンプライアンス エージェント スキャン。 --format pdf --output report.pdf # PDF ファイル (-o エイリアス)
# 重大な問題のみを表示します (情報 | 警告 | 高 | 重大)
コンプライアンス エージェント スキャン。 --重大度高
# フォルダーをスキップ (繰り返し可能)
コンプライアンス エージェント スキャン。 --exclude " testing/* " --exclude " docs/* "
# 一致するフォルダーのみをチェックします (反復可能な許可リスト)
コンプライアンス エージェント スキャン。 --「 src/* 」を含める
# より静かでわかりやすい出力
コンプライアンス エージェント スキャン。 --quit # 概要のみで、結果ごとの詳細はありません
コンプライアンス エージェント スキャン。 --no-color # カラー出力を無効にする
コンプライアンス エージェント スキャン。 --verbose # スキャンされた内容と情報ログを表示します
コンプライアンス エージェント スキャン。 --no-update-check # PyPI バージョンチェックをスキップします
# 各問題の解決方法を示す
コンプライアンス エージェント スキャン。 --修正
# 修正テンプレートをプロジェクトにコピーする
コンプライアンス担当者が推奨します。 --output ./fixes
コンプライアンス担当者が推奨します。 --format json # 機械可読な推奨事項
# 共有可能なレポートファイルを作成する
コンプライアンスエージェントレポート。 --出力監査-2026.pdf
# CI/CD の場合: プレーン出力、重大な問題でビルドが失敗する
コンプライアンス エージェント スキャン。 --ci --フェイルオンハイ
# 最新 (または特定の) バージョンにアップグレードします
遵守する

ance エージェントのアップグレード
コンプライアンス エージェント アップグレード 0.1.2
# インストールされているバージョン (およびアップデートがあるかどうかを表示します)

[切り捨てられた]

## Original Extract

Scan your codebase for EU AI Act compliance in seconds. Detects AI usage (OpenAI, Anthropic, LangChain), flags gaps, gives copy-paste fixes. - latreon/compliance-agent

GitHub - latreon/compliance-agent: Scan your codebase for EU AI Act compliance in seconds. Detects AI usage (OpenAI, Anthropic, LangChain), flags gaps, gives copy-paste fixes. · GitHub
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
latreon
/
compliance-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits .github .github docs docs examples examples rules rules src/ compliance_agent src/ compliance_agent templates templates tests tests .gitignore .gitignore .pre-commit-hooks.yaml .pre-commit-hooks.yaml .python-version .python-version CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Check if your AI project follows EU rules.
The EU has new rules for AI. If you're building with OpenAI, Anthropic, LangChain,
or any AI framework, you need to check whether you comply. This tool does it for
you — one command, about 5 seconds.
30-Second Start · What It Does · How It Works · Examples · All Commands · FAQ
# Install (isolated CLI tool)
uv tool install compliance-agent
# no uv? use: pipx install compliance-agent
# Check your project
compliance-agent scan .
# Done. Read what it found.
What It Does (Simple Version)
Scans your code — finds where you use AI (OpenAI, LangChain, etc.).
Checks the rules — compares your code against EU AI Act requirements.
Tells you what's missing — shows exactly what you need to fix.
Gives you the code — provides copy-paste fixes for each problem.
When you run compliance-agent scan . , you get a boxed terminal report:
a header, a summary strip, per-article coverage, findings, and the gaps to
fix. Illustrative shape (real output for the bundled sample is in
examples/EXPECTED_OUTPUT.md ):
╭─ EU AI Act Compliance Report ──────────────────────────────╮
│ Files scanned 3 │
│ Risk tier LIMITED │
╰──────────────────────────────────────────── ComplianceAgent ╯
╭─ Scan Summary ─────────────────────────────────────────────╮
│ 3 1 6 9 │
│ FILES AI SYSTEMS FINDINGS GAPS │
╰────────────────────────────────────────────────────────────╯
╭─ Compliance Gaps ──────────────────────────────────────────╮
│ ✗ MISSING Automated logging of AI events required (Art.12) │
│ Fix: templates/art12/event_logging.py │
│ ✗ MISSING AI transparency disclosure required (Art. 50) │
│ Fix: templates/art50/transparency_notice.py │
│ ✗ MISSING Error handling around AI calls (Art. 15) │
│ Fix: add try/except + fallbacks around model calls │
╰────────────────────────────────────────────────────────────╯
Next: compliance-agent recommend . --output ./fixes
Prefer a file? scan --format pdf --output report.pdf writes a PDF, and the
separate report command writes markdown or PDF to disk. For scan ,
--format markdown and --format json render to the terminal/stdout — pipe
json to a file if you need one (see Command Reference ).
Use OpenAI, Anthropic, Google, or any AI API
Build chatbots or AI assistants
Use LangChain, CrewAI, AutoGen, or LangGraph
Deploy AI in the EU or serve EU users
Want to avoid fines (up to €15M / 3% of turnover for most obligations, and
up to €35M / 7% for prohibited practices)
Only use AI for personal projects (not a business)
Don't operate in, or serve users in, the EU
ComplianceAgent is a command-line tool, so the cleanest way to install it is
with a tool installer that keeps it in its own isolated environment.
Recommended (isolated CLI install)
uv tool install compliance-agent
or, with pipx:
pipx install compliance-agent
No uv or pipx yet? Install one:
brew install uv # or: brew install pipx
Alternative: pip inside a virtual environment
On modern macOS/Linux, a bare pip install into the system Python is blocked
(PEP 668, "externally-managed-environment"). Use a virtual environment:
python3 -m venv .venv
source .venv/bin/activate # Windows: .venv\Scripts\activate
pip install compliance-agent
Latest unreleased version (from GitHub)
uv tool install git+https://github.com/latreon/compliance-agent.git
# or: pipx install git+https://github.com/latreon/compliance-agent.git
Verify it worked
compliance-agent version
# ComplianceAgent v0.1.7
Trouble installing or running? See the Troubleshooting guide .
The scanner reads your project files and looks for AI-related patterns:
import openai — you're using OpenAI
from langchain — you're using LangChain
AgentExecutor() — you're running an AI agent
client.chat.completions.create() — you're calling an AI API
Provider and framework detection uses AST parsing (not just text search):
they only fire in files that actually import the library, so a comment that
mentions "OpenAI" won't trigger a provider finding. Lightweight keyword patterns
(e.g. chat-interface wording) additionally scan documentation and config, so a
small number of findings can come from .md files.
Based on what it finds, the tool assigns a risk level:
How the tier is decided. HIGH risk under the EU AI Act comes from the
use-case domain ( Annex III ), not from technical
capability. An autonomous agent with tools is not automatically high-risk — a
résumé-screening or credit-scoring system is. ComplianceAgent classifies HIGH
only when it detects Annex III domain indicators, and UNACCEPTABLE only when it
detects a likely Art. 5 prohibited practice (e.g. social scoring, untargeted
facial scraping). Domain indicators are matched against file paths and code
content (not just file names), but only for projects that actually use AI.
Both are keyword-based heuristics and provisional (Art. 6(3) also exempts some
narrow-purpose systems), so a match is a prompt to self-assess and consult
counsel — not a legal determination. See
docs/ARCHITECTURE.md for how tiers
are decided.
The tool checks 13 specific articles of the EU AI Act:
For each issue found, the tool:
Shows which rule requires the fix
Provides a code template you can copy
Tells you exactly where to put it
ISSUE: No "You're talking to AI" notice
RULE: EU AI Act Article 50(1)
FIX: Copy templates/art50/transparency_notice.py into your project
WHERE: Add it before your chat endpoint
Real Examples
Example 1: Simple chatbot (Limited risk)
# chatbot.py
import openai
client = openai . OpenAI ()
def chat ( user_input ):
return client . chat . completions . create (
model = "gpt-4" ,
messages = [{ "role" : "user" , "content" : user_input }],
). choices [ 0 ]. message . content
Scan result (illustrative summary — exact wording/counts come from the scan):
Risk tier: LIMITED — user-facing AI, no Annex III high-risk domain matched
Gaps: Art. 12 (record-keeping), Art. 50 (transparency), Art. 15 (robustness)
Fix: add a transparency notice + event logging + error handling
Example 2: LangChain agent (Higher risk)
An agent that can search the web and send emails:
# agent.py
from langchain . agents import AgentExecutor
from langchain . tools import Tool
tools = [
Tool ( name = "search" , func = search_web , description = "Search the web" ),
Tool ( name = "email" , func = send_email , description = "Send an email" ),
]
executor = AgentExecutor ( agent = agent , tools = tools )
Scan result (illustrative — the tier is HIGH only if your code also matches an
Annex III high-risk domain; agentic patterns on their own drive the oversight
and logging gaps below):
RISK: LIMITED (no Annex III high-risk domain detected)
FRAMEWORKS: LangChain (agent, tools)
ISSUES: several, including
1. No human oversight before tool use (Art. 14)
2. No logging of tool calls (Art. 12)
3. No error handling for API failures (Art. 15)
4. No "You're talking to AI" notice (Art. 50)
FIX: Add human-in-the-loop, logging, error handling, transparency.
Note: tool access raises real Art. 14 oversight and Art. 12 logging gaps, but
it does not by itself make the system HIGH risk. The tier becomes HIGH
only if the agent operates in an Annex III domain (e.g. hiring, credit,
biometrics) — see How the tier is decided .
A crew of agents researching and writing:
# crew.py
from crewai import Agent , Task , Crew
researcher = Agent ( role = "Researcher" , tools = [ search ])
writer = Agent ( role = "Writer" , tools = [ write ])
crew = Crew (
agents = [ researcher , writer ],
tasks = [ Task ( description = "Research" , agent = researcher ),
Task ( description = "Write" , agent = writer )],
)
crew . kickoff ()
Scan result (illustrative):
RISK: LIMITED (multi-agent, but no Annex III high-risk domain detected)
FRAMEWORKS: CrewAI (agent, crew, task)
ISSUES: several, including
1. No oversight before crew execution (Art. 14)
2. No logging of agent actions (Art. 12)
3. No technical documentation (Art. 11)
FIX: Add an approval workflow, logging, and documentation.
A crew researching and writing is not, by itself, an Annex III high-risk
use-case, so the tier stays LIMITED. Point the same crew at résumé screening
or credit decisions and the tier becomes HIGH.
# Scan a folder (. means the current folder)
compliance-agent scan .
# Output types
compliance-agent scan . --format markdown # for reading (default), to terminal
compliance-agent scan . --format json # for computers / CI, to stdout
compliance-agent scan . --format pdf --output report.pdf # PDF file (-o alias)
# Only show serious issues (info | warning | high | critical)
compliance-agent scan . --severity high
# Skip folders (repeatable)
compliance-agent scan . --exclude " tests/* " --exclude " docs/* "
# Only check matching folders (repeatable allow-list)
compliance-agent scan . --include " src/* "
# Quieter / plainer output
compliance-agent scan . --quiet # summary only, no per-finding detail
compliance-agent scan . --no-color # disable colored output
compliance-agent scan . --verbose # show what is scanned + info logs
compliance-agent scan . --no-update-check # skip the PyPI version check
# Show how to fix each problem
compliance-agent scan . --fix
# Copy fix templates into your project
compliance-agent recommend . --output ./fixes
compliance-agent recommend . --format json # machine-readable recommendations
# Make a shareable report file
compliance-agent report . --output audit-2026.pdf
# For CI/CD: plain output, fail the build on serious issues
compliance-agent scan . --ci --fail-on high
# Upgrade to the latest (or a specific) version
compliance-agent upgrade
compliance-agent upgrade 0.1.2
# Show the installed version (and whether an upda

[truncated]
