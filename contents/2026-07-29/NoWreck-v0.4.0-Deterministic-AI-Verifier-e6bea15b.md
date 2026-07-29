---
source: "https://github.com/AstralXVoid/NoWreck"
hn_url: "https://news.ycombinator.com/item?id=49098384"
title: "NoWreck v0.4.0 – – Deterministic AI Verifier"
article_title: "GitHub - AstralXVoid/NoWreck: A CLI tool that verifies AI coding assistant claims against actual structural changes — catching hallucinated functions, fake calls, and missed modifications before they ship. · GitHub"
author: "AstralXVoid"
captured_at: "2026-07-29T15:05:47Z"
capture_tool: "hn-digest"
hn_id: 49098384
score: 1
comments: 0
posted_at: "2026-07-29T14:56:06Z"
tags:
  - hacker-news
  - translated
---

# NoWreck v0.4.0 – – Deterministic AI Verifier

- HN: [49098384](https://news.ycombinator.com/item?id=49098384)
- Source: [github.com](https://github.com/AstralXVoid/NoWreck)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T14:56:06Z

## Translation

タイトル: NoWreck v0.4.0 – – Deterministic AI Verifier
記事のタイトル: GitHub - AstralXVoid/NoWreck: AI コーディング アシスタントの主張を実際の構造変更に対して検証する CLI ツール - 出荷前に幻覚機能、偽の呼び出し、見逃した変更をキャッチします。 · GitHub
説明: AI コーディング アシスタントの主張を実際の構造変更に対して検証する CLI ツールです。出荷前に幻覚機能、偽の呼び出し、見逃した変更を検出します。 - アストラルXヴォイド/ノーレック

記事本文:
GitHub - AstralXVoid/NoWreck: AI コーディング アシスタントの主張を実際の構造的変更に対して検証する CLI ツールで、幻覚機能、偽の呼び出し、変更の欠落を出荷前に検出します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
{{

メッセージ }}
アストラルXヴォイド
/
ナウレック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット docs docs nowreck nowreck test_js_samples test_js_samples test_milestone1 test_milestone1 テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json use.md use.md ビューすべてのファイル リポジトリ ファイルのナビゲーション
+-------------------------------------------------+
|ノーレック v0.4.0 |
|決定論的 AI 検証者 |
+-------------------------------------------------+
NoWreck は、AI コーディング アシスタントのための決定論的検証ツールです。 AIのとき
コードを変更してそれが何をしたかを説明すると、NoWreck は
説明は現実と一致しています - 別の分析ではなく、構造 AST 分析を使用しています
AIさんの意見です。
$ nowreck 修正「auth.py に電子メール検証を追加」
概要
──────────
● 合計 3 件のクレーム
●2個確認済み
● 1 矛盾しています
確認済み
✓ ADD_FUNCTION validate_email → auth.py
証拠: 関数「validate_email」が auth.py に追加されました
矛盾しています
✗ CALLS_FUNCTION validate_email → sanitize_input
証拠: validate_email の本文で sanitize_input への呼び出しは検出されませんでした
何が引っかかるのか
幻覚を起こした関数やクラス — AI は何かを追加したと主張する
それはありません
偽の内部 API 呼び出し — AI は、呼び出していない関数を呼び出したと主張します
説明と差分の不一致 — AI は、次のような変更を説明します。
実際の差分と一致しません
説明のつかない変化 — AIが言及しなかった実際の変更
NoWreck は、AI の説明は何と一致するかという 1 つの質問に正確に答えます。
実際にリポジトリで変更されましたか?それ以上でもそれ以下でもありません。
pipx をインストールします。
（から

このリポジトリのクローン コピー — PyPI の公開は後で行います)
Python 3.10以降が必要です。 nowreck をシステム全体のコマンドとしてインストールします。
NoWreck は、OpenAI 互換モデルのエンドポイント (Groq、
DeepSeek、Ollama、LM Studio、OpenRouter、または OpenAI 自体。
nowreck 構成セットのbase_url https://api.groq.com/openai/v1
nowreck config set api_key < your-api-key >
nowreck 構成セット モデル llama-3.3-70b-versatile
または、NOWRECK_API_KEY 環境変数を保存する代わりに設定します。
構成。
# インタラクティブピッカー — メニュー主導のインターフェース (新規ユーザーに最適)
ナウレック --インタラクティブ
# プロンプト モード — NoWreck はモデルを呼び出し、クレームを取得し、それらを検証します
nowreck 修正「レート制限デコレータを api/client.py に追加する」
# Pre/Post モード — 詳細: 2 つのスナップショットを手動でスキャンします
nowreck 修正 --pre ./repo-v1 --post ./repo-v2
# クレームを含む事前/事後 — 特定のクレームを差分と照合して検証します
nowreck fix --pre ./before --post ./after --claims ' {"claims": [...]} '
# CI パイプラインの JSON 出力
nowreck 修正「auth.py に検証を追加」 --json
# 構成を表示または変更する
ナウレック構成ショー
nowreck 構成セットのbase_url https://api.openai.com/v1
インタラクティブモード (v0.2.0)
コマンドを覚えずにオプションを探索したいユーザー向けのメニュー主導のインターフェイス:
ナウレック --インタラクティブ
これによりターミナル ピッカーが起動し、次のことが可能になります。
プロンプト、事前/事後、またはクレーム モードから選択します
オプションを対話的に構成する
フォーマットされた出力で検証結果を確認する
複雑なコマンドを入力する必要のない 1 回限りの検証
ツールの機能を学ぶ
コマンド
説明
ナウレック
ヘルプ/使用方法を表示する
ナウレック --バージョン
バージョンを表示
ナウレック --インタラクティブ
インタラクティブなターミナル ピッカーを起動します。すべての操作をメニュー形式で行うインターフェイスです。
ナウレック修正「<プロンプト>」
プロンプト モード — 自然言語での変化を説明します。今

reck は、構成されたモデルを呼び出し、diff + クレームを取得し、それらを自動的に検証します。
nowreck fix --pre PATH --post PATH
事前/事後モード — 2 つのディレクトリのスナップショットをスキャンし、構造の変更を検出します。 --claims JSON を追加して、検出された変更に対して特定のクレームを検証します[...]
ナウレック修正 --json
色付きのターミナル テキストの代わりに構造化された JSON を出力します (CI 用)。プロンプトモードと事前/事後モードの両方で動作します。
ナウレック修正 -- 色なし
カラー端子出力を無効にします。
ナウレック構成ショー
現在の構成を表示します。
nowreck config set <キー> <値>
設定値を設定します。キー: api_key 、model 、base_url 、温度 、max_retries 。
仕組み
┌─────────────────────────┐
│ プロンプトモード │
│ │
│ プロンプト ──► AI モデル ──► diff + クレーム │
│ │ │
│ ▼ │
│ プリスキャン ──► シンボルインデックス ──► 変化検出器 │
│ ポストスキャン ──► シンボルインデックス ─┘ │
│ │ │
│ クレーム ──► クレーム検証者 ◄───────┘ │
│ │ 純粋な比較 — AI の判断なし │
│ ▼ │
│ 検証報告書 │
│ ✓ 確認しました ✗ 矛盾していますか?検証不能 │
━━━━━━━━━━━━━━━━━━━┘
NoWreck の検証パイプラインには 3 つのステージがあります。
スキャン — 両方の .py ファイルと .js ファイルを再帰的に検出します
スナップショットを作成し、適切なパーサーでそれぞれを解析し、
すべての関数、クラス、メソッドのシンボル インデックス。
Python ファイル → Python の組み込み ast モジュールで解析
JavaScript ファイル → 解析

ツリーシッターとのd
(tree-sitter-javascript 文法)
どちらのパーサーも同じ Symbol / SymbolType データ形状を生成するため、
パイプラインの残りの部分はどの言語を使用するかを決して知りません (または気にしません)
データを作成しました。
検出 — 前後のシンボルのインデックスを比較して検出します。
構造上の変更: 追加/削除された関数、クラス、ファイル、および新規
関数呼び出し。これにより、唯一の真実の情報源が生成されます。
検証者が排他的に参照する list[DetectedChange]。
検証 — AI モデルからのクレームごとに、検証者は次のことを探します。
一致する DetectedChange 。同じタイプのものが存在する場合、
ID フィールド → CONFIRMED 。矛盾する変更が存在する場合
(例: クレームでは「追加」と表示されますが、検出では「削除」と表示されます) →
矛盾しています。何も一致しない場合 → UNVERIFIABLE 。
ベリファイアは AST を解析したり、シンボル インデックスをクエリしたりすることはありません。
AI判定を適用します。その決定は純粋にフィールドベースの比較です。
請求の種類
それが何を意味するか
検証者
機能の追加
機能が追加されました
構造物存在確認
REMOVE_FUNCTION
機能が削除されました
構造物存在確認
ADD_CLASS
クラスが追加されました
構造物存在確認
REMOVE_CLASS
クラスが削除されました
構造物存在確認
ファイル_作成済み
新しいファイルが現れました
ファイルリストの差分
ファイル_削除済み
ファイルが削除されました
ファイルリストの差分
コール_ファンクション
関数が別の関数を呼び出すようになりました
AST コールサイト検出
7 つすべてが直接的な構造的事実によって検証されています - キーワードはありません
推測であり、意味解釈はありません。 NoWreck が何かを判断できない場合
推測ではなく、確実に「検証不可能」と報告します。
すべての結果には信頼スコアが含まれます。これはNoWreckの確実性を反映しています
決定論的チェックにおいては、基礎となるコードが次のとおりであるという主張ではありません。
バグなし。
100% 確認済み — 構造的事実が見つかり、一致しました
矛盾しています

100% — 反対の構造的事実が発見されました
50% では検証不可能 — どちらにしても一致する事実は存在しない
一致する呼び出しが見つからない CALLS_FUNCTION チェックは、
一致するものを見つけるもの。直接検査によって確認された欠席は、不許可ではありません。
存在よりも弱い事実。
これは、NoWreck が確実であるという意味ではありません。静的解析には実際の、
文書化された制限 — 以下の制限を参照してください。
ツール
何をするのか
NoWreckと重複しますか？
カーソル / クロード コード / 副操縦士
コードの生成と編集
いいえ — NoWreck は検証しますが、生成しません
CodeRabbit / Qodo / グレプタイル
AIによる差分品質のレビュー
いいえ — 決定論的な事実確認ではなく、主観的な AI 判断
エージェントベリファイア (aurite-ai)
コード品質/セキュリティのための AI エージェント スキル
いいえ — 真実性を主張するのではなく、コードの品質をチェックします
スロップチェック / スロップスキャン
サードパーティのパッケージ名をレジストリと照合する
いいえ — 異なる幻覚カテゴリー
ESLint / ラフ / ブラック
リンティングとフォーマット
いいえ — 構造検証ではなく、構文/コード スタイル
NoWreck は、AI 主張の決定論的検証という独自のニッチ分野を占めています
コードの変更について。これを行うツールは他にありません。
Python + JavaScript — NoWreck は両方の言語をサポートします。 Python ファイル
組み込みの ast モジュールで解析されます。 JavaScript ファイルの使用
Tree-sitter-JavaScript 文法を使用した Tree-sitter。
動的動作を見抜けない - exec() 、 eval() 、dynamic
インポート、動的引数を使用した getattr() / setattr()、メタクラス、
モンキーパッチとリフレクションはすべて検証不能になります
単純な呼び出しのみ — obj.method() や obj.method() ではなく name() 呼び出しを検出します。
連鎖通話
直接の名前一致以外のファイル間の解決はありません
セマンティック分析はありません - 意図ではなく構造を検証します
TypeScript のサポートなし — TypeScript 構文はまだサポートされていません (ファイル
スキャン中にスキップされます）。
対話型端末

非 CLI ユーザー向けのアルピッカー ✅ (v0.2.0 で実装)
JavaScript コアのサポート (ツリーシッター スキャナー + シンボル インデックス) ✅ (v0.3.0 で完了)
JavaScript の洗練 (ジェネレーター、エクスポートのデフォルト、IIFE) ✅ (v0.4.0 で完了)
--verbose モードでは、主張ごとに完全な決定論的な証拠を表示します。
追加のモデルプロバイダー (Anthropic、Gemini)
大規模なリポジトリのキャッシュ
# リポジトリのクローンを作成します
git クローン https://github.com/AstralXVoid/NoWreck.git
cd ノーレック
# pipx を使用してシステム全体にインストールします (推奨)
pipx をインストールします。
# または pip を使用して
pip install -e 。
# または仮想環境内
python3 -m venv .venv
ソース .venv/bin/activate
pip install -e 。
インストールの確認
ナウレック --バージョン
# → ナウレック0.3.0
ナウレック
# → バナー + 使用法を表示
アンインストール
pipx アンインストール ナウレック
# または
pip アンインストール ナウレック
構成
NoWreck は、現在の設定の下に .nowreck/config.json に設定を保存します。
作業ディレクトリ。
プロンプトモードに必要な設定
nowreck fix "<prompt>" を使用する前に、API キーを設定する必要があります
およびモデルプロバイダー:
# API キーを設定します (または、代わりに NOWRECK_API_KEY 環境変数を使用します)
nowreck config set api_key gsk_your_key_here
# API ベース URL を設定します (デフォルトは https://api.openai.com/v1)
nowreck 構成セットのbase_url https://api.groq.com/openai/v1
# モデルを設定します (デフォルトは gpt-4o)
nowreck 構成セット モデル llama-3.3-70b-versatile
代替: 環境変数
キーを構成に保存する代わりに NOWRECK_API_KEY を設定します。
import NOWRECK_API_KEY= " gsk_your_key_here "
オプション設定
# 温度 (0.0 = 決定的、デフォルト)
ナウレック設定温度 0.0
# 解析失敗時の最大再試行回数 (デフォルト: 1)
nowreck 構成セット max_retries 2
構成を表示する
ナウレック構成ショー
# → api_key = gsk_...
# →base_url = https://api.groq.com/openai/v1
# →

[切り捨てられた]

## Original Extract

A CLI tool that verifies AI coding assistant claims against actual structural changes — catching hallucinated functions, fake calls, and missed modifications before they ship. - AstralXVoid/NoWreck

GitHub - AstralXVoid/NoWreck: A CLI tool that verifies AI coding assistant claims against actual structural changes — catching hallucinated functions, fake calls, and missed modifications before they ship. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
AstralXVoid
/
NoWreck
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits docs docs nowreck nowreck test_js_samples test_js_samples test_milestone1 test_milestone1 tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json use.md use.md View all files Repository files navigation
+------------------------------------+
| NoWreck v0.4.0 |
| Deterministic AI Verifier |
+------------------------------------+
NoWreck is a deterministic verifier for AI coding assistants. When an AI
changes your code and explains what it did, NoWreck checks whether the
explanation matches reality — using structural AST analysis, not another
AI's opinion.
$ nowreck fix "Add email validation to auth.py"
Summary
────────────────────
● 3 claims total
● 2 confirmed
● 1 contradicted
CONFIRMED
✓ ADD_FUNCTION validate_email → auth.py
Evidence: Function 'validate_email' was added in auth.py
CONTRADICTED
✗ CALLS_FUNCTION validate_email → sanitize_input
Evidence: No call to sanitize_input detected in validate_email's body
What it catches
Hallucinated functions or classes — the AI claims it added something
that isn't there
Fake internal API calls — the AI says it called a function it didn't
Explanation-vs-diff mismatches — the AI describes a change that
doesn't match the actual diff
Unexplained changes — real modifications the AI never mentioned
NoWreck answers exactly one question: does the AI's explanation match what
actually changed in the repository? Nothing more, nothing less.
pipx install .
(from a cloned copy of this repo — PyPI publishing coming later)
Requires Python 3.10+. Installs nowreck as a system-wide command.
NoWreck works with any OpenAI-compatible model endpoint — Groq,
DeepSeek, Ollama, LM Studio, OpenRouter, or OpenAI itself.
nowreck config set base_url https://api.groq.com/openai/v1
nowreck config set api_key < your-api-key >
nowreck config set model llama-3.3-70b-versatile
Or set the NOWRECK_API_KEY environment variable instead of storing it in
config.
# Interactive picker — menu-driven interface (great for new users)
nowreck --interactive
# Prompt mode — NoWreck calls the model, gets claims, verifies them
nowreck fix " Add a rate-limiting decorator to api/client.py "
# Pre/Post mode — advanced: scan two snapshots manually
nowreck fix --pre ./repo-v1 --post ./repo-v2
# Pre/Post with claims — verify specific claims against a diff
nowreck fix --pre ./before --post ./after --claims ' {"claims": [...]} '
# JSON output for CI pipelines
nowreck fix " Add validation to auth.py " --json
# View or change configuration
nowreck config show
nowreck config set base_url https://api.openai.com/v1
Interactive Mode (in v0.2.0)
Menu-driven interface for users who prefer exploring options without memorizing commands:
nowreck --interactive
This launches a terminal picker where you can:
Choose from Prompt, Pre/Post, or Claims modes
Configure options interactively
See verification results with formatted output
One-off verifications without typing complex commands
Learning the tool's capabilities
Command
Description
nowreck
Show help / usage
nowreck --version
Show version
nowreck --interactive
Launch the interactive terminal picker — menu-driven interface for all operations
nowreck fix "<prompt>"
Prompt mode — describe a change in natural language. NoWreck calls the configured model, gets a diff + claims, and verifies them automatically.
nowreck fix --pre PATH --post PATH
Pre/Post mode — scan two directory snapshots and detect structural changes. Add --claims JSON to verify specific claims against the detected chang[...]
nowreck fix --json
Output structured JSON instead of coloured terminal text (for CI). Works with both prompt and pre/post modes.
nowreck fix --no-colour
Disable coloured terminal output.
nowreck config show
Display current configuration.
nowreck config set <key> <value>
Set a configuration value. Keys: api_key , model , base_url , temperature , max_retries .
How it works
┌─────────────────────────────────────────────────────────┐
│ Prompt mode │
│ │
│ Your prompt ──► AI model ──► diff + claims │
│ │ │
│ ▼ │
│ Pre-scan ──► Symbol index ──► Change Detector │
│ Post-scan ──► Symbol index ────────┘ │
│ │ │
│ Claims ──► Claim Verifier ◄────────────┘ │
│ │ pure comparison — no AI judgment │
│ ▼ │
│ Verification Report │
│ ✓ CONFIRMED ✗ CONTRADICTED ? UNVERIFIABLE │
└─────────────────────────────────────────────────────────┘
NoWreck's verification pipeline has three stages:
Scan — recursively discovers .py and .js files in both
snapshots, parses each with the appropriate parser, and builds a
symbol index of every function, class, and method.
Python files → parsed with Python's built-in ast module
JavaScript files → parsed with Tree-sitter
(the tree-sitter-javascript grammar)
Both parsers produce the same Symbol / SymbolType data shapes, so
the rest of the pipeline never knows (or cares) which language
produced the data.
Detect — compares the pre and post symbol indices to find
structural changes: added/removed functions, classes, files, and new
function calls. This produces the single source of truth — a
list[DetectedChange] that the verifier references exclusively.
Verify — for each claim from the AI model, the verifier looks for
a matching DetectedChange . If one exists with the same type and
identity fields → CONFIRMED . If a contradicting change exists
(e.g., claim says "added" but detection shows "removed") →
CONTRADICTED . If nothing matches → UNVERIFIABLE .
The verifier never parses AST, never queries the symbol index, and never
applies AI judgment. Its decisions are purely field-based comparison.
Claim type
What it means
Verified by
ADD_FUNCTION
A function was added
Structural existence check
REMOVE_FUNCTION
A function was removed
Structural existence check
ADD_CLASS
A class was added
Structural existence check
REMOVE_CLASS
A class was removed
Structural existence check
FILE_CREATED
A new file appeared
File-list diff
FILE_DELETED
A file was removed
File-list diff
CALLS_FUNCTION
A function now calls another
AST call-site detection
All seven are verified through direct structural facts — no keyword
guessing, no semantic interpretation. If NoWreck can't determine something
with certainty, it reports UNVERIFIABLE rather than guessing.
Every result includes a confidence score. This reflects NoWreck's certainty
in its deterministic check — not a claim that the underlying code is
bug-free.
CONFIRMED at 100% — the structural fact was found and matched
CONTRADICTED at 100% — the opposite structural fact was found
UNVERIFIABLE at 50% — no matching fact exists either way
A CALLS_FUNCTION check that finds no matching call is just as certain as
one that finds a match. An absence, confirmed by direct inspection, is not a
weaker fact than a presence.
This does not mean NoWreck is infallible. Static analysis has real,
documented limits — see Limitations below.
Tool
What it does
Overlaps with NoWreck?
Cursor / Claude Code / Copilot
Generate and edit code
No — NoWreck verifies, doesn't generate
CodeRabbit / Qodo / Greptile
AI review of a diff's quality
No — subjective AI judgment, not deterministic fact-checking
Agent Verifier (aurite-ai)
AI agent skill for code quality/security
No — checks code quality, not claim truthfulness
slopcheck / slop-scan
Check third-party package names against registries
No — different hallucination category
ESLint / Ruff / Black
Linting and formatting
No — syntax/code style, not structural verification
NoWreck occupies a unique niche: deterministic verification of AI claims
about code changes. No other tool does this.
Python + JavaScript — NoWreck supports both languages. Python files
are parsed with the built-in ast module; JavaScript files use
Tree-sitter with the tree-sitter-javascript grammar.
Cannot see through dynamic behavior — exec() , eval() , dynamic
imports, getattr() / setattr() with dynamic arguments, metaclasses,
monkey-patching, and reflection will all yield UNVERIFIABLE
Simple calls only — detects name() calls, not obj.method() or
chained calls
No cross-file resolution beyond direct name matching
No semantic analysis — it verifies structure, not intent
No TypeScript support — TypeScript syntax is not yet supported (files
will be skipped during scanning).
Interactive terminal picker for non-CLI users ✅ (done in v0.2.0)
JavaScript core support (Tree-sitter scanner + symbol index) ✅ (done in v0.3.0)
JavaScript polish (generators, export default, IIFEs) ✅ (done in v0.4.0)
--verbose mode showing full deterministic evidence per claim
Additional model providers (Anthropic, Gemini)
Caching for large repositories
# Clone the repository
git clone https://github.com/AstralXVoid/NoWreck.git
cd NoWreck
# Install system-wide with pipx (recommended)
pipx install .
# Or with pip
pip install -e .
# Or inside a virtual environment
python3 -m venv .venv
source .venv/bin/activate
pip install -e .
Verify installation
nowreck --version
# → nowreck 0.3.0
nowreck
# → shows banner + usage
Uninstall
pipx uninstall nowreck
# or
pip uninstall nowreck
Configuration
NoWreck stores configuration in .nowreck/config.json under the current
working directory.
Required settings for Prompt mode
Before using nowreck fix "<prompt>" , you need to configure an API key
and model provider:
# Set your API key (or use the NOWRECK_API_KEY env var instead)
nowreck config set api_key gsk_your_key_here
# Set the API base URL (defaults to https://api.openai.com/v1)
nowreck config set base_url https://api.groq.com/openai/v1
# Set the model (defaults to gpt-4o)
nowreck config set model llama-3.3-70b-versatile
Alternative: Environment variable
Set NOWRECK_API_KEY instead of storing the key in config:
export NOWRECK_API_KEY= " gsk_your_key_here "
Optional settings
# Temperature (0.0 = deterministic, default)
nowreck config set temperature 0.0
# Max retries on parse failure (default: 1)
nowreck config set max_retries 2
View configuration
nowreck config show
# → api_key = gsk_...
# → base_url = https://api.groq.com/openai/v1
# →

[truncated]
