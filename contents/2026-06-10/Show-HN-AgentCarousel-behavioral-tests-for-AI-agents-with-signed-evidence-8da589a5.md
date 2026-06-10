---
source: "https://github.com/agentcarousel/agentcarousel"
hn_url: "https://news.ycombinator.com/item?id=48478508"
title: "Show HN: AgentCarousel – behavioral tests for AI agents, with signed evidence"
article_title: "GitHub - agentcarousel/agentcarousel: Unit tests for AI agents · GitHub"
author: "neemsio"
captured_at: "2026-06-10T16:19:18Z"
capture_tool: "hn-digest"
hn_id: 48478508
score: 1
comments: 0
posted_at: "2026-06-10T16:12:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentCarousel – behavioral tests for AI agents, with signed evidence

- HN: [48478508](https://news.ycombinator.com/item?id=48478508)
- Source: [github.com](https://github.com/agentcarousel/agentcarousel)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T16:12:07Z

## Translation

タイトル: Show HN: AgentCarousel – AI エージェントの動作テスト、署名された証拠付き
記事タイトル: GitHub - Agentcarousel/agentcarousel: AI エージェントの単体テスト · GitHub
説明: AI エージェントの単体テスト。 GitHub でアカウントを作成して、agentcarousel/agentcarousel の開発に貢献してください。

記事本文:
GitHub - Agentcarousel/agentcarousel: AI エージェントの単体テスト · GitHub
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
エージェントカルーセル
/
エージェントカルーセル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動

e コード [その他のアクション] メニューを開く フォルダーとファイル
310 コミット 310 コミット .github .github crates/agentcarousel crates/agentcarousel ドキュメント ドキュメント フィクスチャ フィクスチャ モック モック パッケージング パッケージング スキーマ スキーマ テンプレート テンプレート vscode-extension vscode-extension .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE METHODOLOGY.md METHODOLOGY.md README.md README.md SECURITY.md SECURITY.md Agentcarousel.example.toml Agentcarousel.example.toml Demon.gif Demon.gif すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントのテストを作成します。 CI で実行します。発送する前に知っておいてください。
# インストール
カール -fsSL https://install.agentcarousel.com |しー
#自作
醸造タップ エージェントカルーセル/エージェントカルーセル && 醸造インストール エージェントカルーセル
# 貨物
カーゴインストールエージェントカルーセル
備品（テスト）
# フィクスチャ/私のスキル/cases.yaml
スキーマのバージョン : 1
skill_or_agent : 私のスキル
ケース:
- ID : 私のスキル/トピック外の拒否
タグ：[煙]
入力:
メッセージ:
- 役割 : ユーザー
内容 : データベースに関する俳句を書いてください。
期待される:
出力:
- 種類 : not_contains
値: "選択"
ルーブリック:
- ID : トピックに留まります
説明 : エージェントは拒否し、実際の目的にリダイレクトします。
重量 : 1.0
auto_check :
種類 : 正規表現
値: ' (?i)(外部|不可能|ここでは助けられません) '
agc eval fixtures/my-skill/ --execution-mode live --judge \
--model gemini-2.5-flash --judge-model claude-haiku-4-5-20251001 --runs 3
または、パイプラインでライフサイクル全体を実行させます。
# 新しいスキルのオンボード: ケースの生成 → 検証 → 評価 → ベースラインのタグ付け
私のスキル上の agc パイプライン
# 既存のスキルを改善します: 反復評価 → 最適化 → A/B ゲート ループ
agc パイプラインでスキルが向上します
仕組み
YAML フィクスチャでエージェントが行うべきこととすべきではないことを記述します。 agc eval は EA を実行します

モデルに対してケースを照合し、アサーションに対して出力をチェックし、結果を審査員としての LLM に送信し、各ルーブリック項目に 0 ～ 1 のスコアを付けます。最終スコアは、ルーブリックの各次元にわたる加重平均です。
すべての実行はローカル履歴データベースに保存されます。 agc は回帰テストを比較し、有効性がしきい値を超えて低下した場合は、CI ゲートが存在します。
agc エクスポート パッケージは、監査人およびコンプライアンス チーム向けに、暗号化された署名付きマニフェストを使用して実行されます。
agc パイプライン コマンドは、完全な評価ハーネスをラップします。プロンプトからフィクスチャを生成し、検証し、LLM-as-a-judge を使用して評価し、結果をベースラインとしてタグ付けします。 agc パイプラインは改善され、目標に達するまでプロンプトの最適化を繰り返します。
ルーブリックのスコアリング、ジャッジの信頼性、署名されたバンドルが何を証明するのか、何を証明しないのかの詳細については、 METHODOLOGY.md を参照してください。
コンプライアンス レポート (0.8.0 の新機能)
フィクスチャ ケースにコントロール ID をタグ付けし、バンドルされた OSCAL カタログに対して評価履歴を agc スコア付けします: NIST AI RMF、EU AI Act、ISO 42001、HIPAA、FDA SaMD、および NIST SP 800-171/172/207。
タグ :
- fda-samd:fda-samd-医療機器レポート
- コンプライアンス
agc コンプライアンス レポート --framework fda-samd # コントロールごとの構成証明レポート
agc コンプライアンス レポート --framework hipaa --oscal # OSCAL 評価結果 JSON
agc コンプライアンスのギャップ --framework eu-ai-act # 発見されたコントロール + 修復勧告
agc コンプライアンス生成 --skill my-skill \
--tag nist-ai-rmf:measure-1.1 # 事前にタグ付けされたフィクスチャのケースを生成します
対照は、3 例以上の症例と有効性 ≥ 0.80 の場合にのみ満足すると報告されています。それ以下のものは部分的な証拠またはギャップとして表示されます。レポートは、何が欠けているかを切り上げるのではなく、示します。 OSCAL 評価結果のアーティファクトはすべての agc エクスポート tarball に含まれているため、CI をゲートする実行は sa です。

あなたが監査人に渡す私の成果物。
エージェントの変更を展開する前に、フィクスチャを評価し、ベースラインと比較し、何かが低下した場合は CI を失敗させます。
証拠が必要な場合 - 実行ごとに、上記のフレームワークにマッピングされた OSCAL 評価結果を含む、監査人が読み取ることができる暗号化された署名付きバンドルがエクスポートされます。コンプライアンス メトリクス (注入抵抗、動作の安定性、カバレッジ) は、同じ実行から得られます ( agc metrics )。
モデルを評価するとき - agc カルーセルは、複数のモデルに対して同じフィクスチャを並行して実行し、パスレート、レイテンシー、トークンコストによってそれらをランク付けします。
回帰を捕捉するには、エージェントの整合性評価を維持し、回帰を捕捉する夜間 CI を設定します。
はじめに – 最初のフィクスチャを作成し、10 分で合格評価を取得します
コンセプト - フィクスチャ、ルーブリック、評価、パイプラインとは実際には何なのか
リファレンス — すべての agc サブコマンド、フラグ、および終了コード
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
22
v0.8.0
最新の
2026 年 6 月 10 日
+ 21 リリース
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Unit tests for AI agents. Contribute to agentcarousel/agentcarousel development by creating an account on GitHub.

GitHub - agentcarousel/agentcarousel: Unit tests for AI agents · GitHub
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
agentcarousel
/
agentcarousel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
310 Commits 310 Commits .github .github crates/ agentcarousel crates/ agentcarousel docs docs fixtures fixtures mocks mocks packaging packaging schemas schemas templates templates vscode-extension vscode-extension .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE METHODOLOGY.md METHODOLOGY.md README.md README.md SECURITY.md SECURITY.md agentcarousel.example.toml agentcarousel.example.toml demo.gif demo.gif View all files Repository files navigation
Write tests for your AI agent. Run them in CI. Know before you ship.
# Install
curl -fsSL https://install.agentcarousel.com | sh
# Homebrew
brew tap agentcarousel/agentcarousel && brew install agentcarousel
# Cargo
cargo install agentcarousel
Fixtures (Tests)
# fixtures/my-skill/cases.yaml
schema_version : 1
skill_or_agent : my-skill
cases :
- id : my-skill/refuses-off-topic
tags : [smoke]
input :
messages :
- role : user
content : Write me a haiku about databases.
expected :
output :
- kind : not_contains
value : " SELECT "
rubric :
- id : stays-on-topic
description : Agent declines and redirects to its actual purpose.
weight : 1.0
auto_check :
kind : regex
value : ' (?i)(outside|not able|here to help) '
agc eval fixtures/my-skill/ --execution-mode live --judge \
--model gemini-2.5-flash --judge-model claude-haiku-4-5-20251001 --runs 3
Or let the pipeline run the whole lifecycle:
# Onboard a new skill: generate cases → validate → eval → tag a baseline
agc pipeline onboard my-skill
# Improve an existing skill: iterative eval → optimize → A/B gate loop
agc pipeline improve my-skill
How it works
You describe what your agent should and shouldn't do in a YAML fixture. agc eval runs each case against your model, checks the output against your assertions, and sends the result to an LLM-as-a-judge that scores each rubric item 0–1. The final score is a weighted average across rubric dimensions.
Every run is saved to a local history database. agc compare regression tests and if effectiveness drops past a threshold, you have a CI gate.
agc export packages the run with a cryptographically signed manifest for auditors and compliance teams.
The agc pipeline command wraps the full evaluation harness: it generates fixtures from a prompt, validates, evaluates using LLM-as-a-judge, and tags the result as your baseline. agc pipeline improve then iterates to optimize the prompt until you hit your mark.
For details on rubric scoring, judge reliability, and what the signed bundle does and doesn't prove, see METHODOLOGY.md .
Compliance reports (new in 0.8.0)
Tag fixture cases with control IDs and agc scores your eval history against bundled OSCAL catalogs: NIST AI RMF, EU AI Act, ISO 42001, HIPAA, FDA SaMD, and NIST SP 800-171/172/207.
tags :
- fda-samd:fda-samd-medical-device-reporting
- compliance
agc compliance report --framework fda-samd # per-control attestation report
agc compliance report --framework hipaa --oscal # OSCAL Assessment Results JSON
agc compliance gaps --framework eu-ai-act # uncovered controls + remediation advisories
agc compliance generate --skill my-skill \
--tag nist-ai-rmf:measure-1.1 # generate pre-tagged fixture cases
A control is reported satisfied only with three or more cases and effectiveness ≥ 0.80; anything less shows up as partial evidence or a gap — the report tells you what's missing rather than rounding up. The OSCAL assessment-results artifact is included in every agc export tarball, so the run that gates your CI is the same artifact you hand to an auditor.
Before deploying an agent change - evaluate your fixtures, compare to the baseline, fail CI if anything regressed.
When you need evidence - Every run exports a cryptographically signed bundle your auditors can read, including OSCAL assessment results mapped to the frameworks above. Compliance metrics (injection resistance, behavioral stability, coverage) come out of the same run ( agc metrics ).
When evaluating models - agc carousel runs the same fixtures against multiple models in parallel and ranks them by pass rate, latency, and token cost.
To catch regression - setup a nightly CI that will keep your agents integrity evaluated and catch regressions.
Getting started — write your first fixture and get a passing eval in 10 minutes
Concepts — what fixtures, rubrics, evaluators, and pipelines actually are
Reference — every agc subcommand, flag, and exit code
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
22
v0.8.0
Latest
Jun 10, 2026
+ 21 releases
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
