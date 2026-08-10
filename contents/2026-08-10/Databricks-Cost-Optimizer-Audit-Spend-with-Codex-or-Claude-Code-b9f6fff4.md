---
source: "https://github.com/kylehuirevvision/databricks-cost-optimizer"
hn_url: "https://news.ycombinator.com/item?id=49242026"
title: "Databricks Cost Optimizer: Audit Spend with Codex or Claude Code"
article_title: "GitHub - kylehuirevvision/databricks-cost-optimizer: Read-only-first Databricks FinOps toolkit and Agent Skill for Codex and Claude Code · GitHub"
author: "kylehui818"
captured_at: "2026-08-10T10:57:57Z"
capture_tool: "hn-digest"
hn_id: 49242026
score: 1
comments: 0
posted_at: "2026-08-10T10:53:21Z"
tags:
  - hacker-news
  - translated
---

# Databricks Cost Optimizer: Audit Spend with Codex or Claude Code

- HN: [49242026](https://news.ycombinator.com/item?id=49242026)
- Source: [github.com](https://github.com/kylehuirevvision/databricks-cost-optimizer)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T10:53:21Z

## Translation

タイトル: Databricks Cost Optimizer: Codex または Claude Code による支出の監査
記事のタイトル: GitHub - kylehuirevvision/databricks-cost-optimizer: 読み取り専用の Databricks FinOps ツールキットと Codex および Claude Code 用のエージェント スキル · GitHub
説明: Codex および Claude Code 用の読み取り専用 Databricks FinOps ツールキットとエージェント スキル - kylehuirevvision/databricks-cost-optimizer

記事本文:
GitHub - kylehuirevvision/databricks-cost-optimizer: Codex および Claude Code 用の読み取り専用 Databricks FinOps ツールキットとエージェント スキル · GitHub
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
カイルフイレヴィジョン
/
データブリック-コスト-オプティマイザー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .agents/ skill/ optimize-databricks-costs .agents/ skill/ optimize-databricks-costs .claude/ skill/ optimize-da

tabricks-costs .claude/ skill/ optimize-databricks-costs .github/ workflows .github/ workflows docs/images docs/images 例 例 スクリプト スクリプト テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
読み取り専用の Databricks FinOps ツールキットとエージェント スキルにより、コストのスパイクを検出し、そのワークロードへの影響を説明し、明示的に承認された最適化のみを適用します。
Codex と Claude Code の両方で動作します。
Codex は .agents/skills/optimize-databricks-costs/ を検出します。
Claude Code は .claude/skills/optimize-databricks-costs/ を検出します。
どちらのコピーも同じオープン SKILL.md 形式を使用し、CI によって同一に保たれます。
ウェアハウスの自動停止を繰り返し超える SQL ポーリング間隔
サイズが大きすぎる SQL ウェアハウスまたは継続的にアクティブな SQL ウェアハウス
使用例がほとんど実証されていない、常時実行されている Databricks アプリ
失敗、キャンセル、重複、またはカナリアジョブの実行
実際のワークロードに応じて解決される高コストのユーザーとサービス プリンシパル
コスト配分タグが欠落しており、予算範囲が弱い
鮮度、待ち時間、可用性に影響を与える節約提案
デフォルトのワークフローは読み取り専用です。作業は次の 4 つの明示的な段階に分けられます。
影響を受けるすべてのワークロードを特定し、影響に対する承認を取得します。
ロールバック コマンドとライブ検証を使用して、承認された変更のみを適用します。
バンドルされた監査スクリプトは、Databricks システム テーブルとインベントリ API を使用します。ウェアハウスの編集、アプリの停止、ジョブの一時停止、権限の変更、予算の作成は行われません。
Databricks CLI 0.229 以降。ライブテストでは0.296が使用されました
構成された Databricks CLI プロファイル
実行中の SQL ウェアハウスへのアクセス
関連する system.billing 、 system.query 、 system.lake を読み取る権限

flow および system.compute テーブル
アカウント管理者アクセスでは最も広範な監査が可能ですが、ツールキットは現在の ID に表示が許可されているものをすべてレポートします。
git clone https://github.com/kylehuirevvision/databricks-cost-optimizer.git
cd databricks-cost-optimizer
コーデックス
次に、次のように呼び出します。
$optimize-databricks-costs はコストの観点から Databricks アカウントを監査します
Codex は、OpenAI スキルの公式ドキュメントに従って .agents/skills からリポジトリ スキルを読み込みます。
git clone https://github.com/kylehuirevvision/databricks-cost-optimizer.git
cd databricks-cost-optimizer
クロード
次に、次のように呼び出します。
/optimize-databricks-costs コストの観点から Databricks アカウントを監査します
Claude Code は、Claude Code の公式スキル ドキュメントに従って、プロジェクト スキルを .claude/skills からロードします。
証拠コレクターを直接実行する
Databricks ウェアハウスのリスト
読み取り専用監査を実行します。
python3 .agents/skills/optimize-databricks-costs/scripts/audit.py \
--warehouse-id YOUR_WAREHOUSE_ID \
--最近 14 \
--output-dir レポート/最新
--profile PROFILE を指定してデフォルト以外の CLI プロファイルを使用します。 --dry-run を使用して Databricks に接続せずに、すべての SQL ステートメントをプレビューします。
このコマンドは、JSON 証拠とスターター マークダウン サマリーを reports/ の下に書き込みます。通常、レポートにはアカウント ID、電子メール、クエリ メタデータ、リソース名が含まれるため、このディレクトリは Git によって無視されます。
システム テーブル クエリが失敗した場合でも、コレクターは取得した証拠を書き込みますが、ゼロ以外で終了し、サマリーに不完全のマークを付けます。証拠が欠けていてもコストがゼロとして報告されることはありません。
サニタイズされた Databricks の請求ビュー。 8 月 10 日は半日であるため、以下の比較から除外されています。
このワークフローが使用された最初のアカウントでは、完了した 1 日あたりの実効定価使用量は、8 月 5 日の 343 ドルから 8 月 9 日には 108 ドルに減少し、68.5% 削減されました。

ション。 SQL の使用量は 257 ドルから 42 ドルに減少し、83.7% 削減され、測定された改善の大部分を占めました。
これは観察された前後の結果であり、対照実験や普遍的な貯蓄主張ではありません。ワークロードの量は変動する可能性があり、最近の請求が部分的または再計算される可能性があり、交渉された請求書や基礎となるクラウド インフラストラクチャのコストが異なる場合があります。
匿名化された取り組みにより、10 分間自動停止される倉庫に対して、キャッシュされた小さなクエリが約 7.5 分ごとに実行されていることがわかりました。クエリではコンピューティングはほとんど実行されませんでしたが、その頻度により、2 週間の枠内で 1 時間ごとにウェアハウスに請求が可能でした。再発を阻止し、倉庫のサイズを縮小し、自動停止を短縮し、未使用のアプリを停止した後、翌日のコストは大幅に減少しました。
推論パターンと注意事項については、サニタイズされたケーススタディを参照してください。
python3 -m 単体テスト 検出 -s テスト -v
python3 スクリプト/sync_skill.py --check
python3 スクリプト/check_public_safety.py
Codex の組み込みスキルクリエーターを持っているメンテナは、両方のスキル ディレクトリに対してquick_validate.py を追加で実行できます。一般ユーザーにはその内部ヘルパーは必要ありません。
監査では、 system.billing.list_prices の実効定価を使用して USD コストを見積もります。これは、必ずしも交渉された請求書の金額であるとは限りません。最近の請求レコードは遅れて到着したり、修正されたりする可能性があるため、均等期間の比較では、現在の日と構成可能な取り込みラグ ウィンドウがデフォルトで除外されます。
Codex および Claude Code 用の読み取り専用 Databricks FinOps ツールキットとエージェント スキル
Readme MIT ライセンス
セキュリティポリシー アクティビティスター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Read-only-first Databricks FinOps toolkit and Agent Skill for Codex and Claude Code - kylehuirevvision/databricks-cost-optimizer

GitHub - kylehuirevvision/databricks-cost-optimizer: Read-only-first Databricks FinOps toolkit and Agent Skill for Codex and Claude Code · GitHub
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
kylehuirevvision
/
databricks-cost-optimizer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .agents/ skills/ optimize-databricks-costs .agents/ skills/ optimize-databricks-costs .claude/ skills/ optimize-databricks-costs .claude/ skills/ optimize-databricks-costs .github/ workflows .github/ workflows docs/ images docs/ images examples examples scripts scripts tests tests .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
A read-only-first Databricks FinOps toolkit and Agent Skill for finding cost spikes, explaining their workload impact, and applying only explicitly approved optimizations.
It works in both Codex and Claude Code:
Codex discovers .agents/skills/optimize-databricks-costs/ .
Claude Code discovers .claude/skills/optimize-databricks-costs/ .
Both copies use the same open SKILL.md format and are kept identical by CI.
SQL polling intervals that repeatedly beat warehouse auto-stop
Oversized or continuously active SQL warehouses
Always-running Databricks Apps with little demonstrated use
Failed, cancelled, duplicate, or canary Job runs
High-cost users and service principals, resolved to their real workloads
Missing cost-allocation tags and weak budget coverage
Savings proposals that silently create freshness, latency, or availability impacts
The default workflow is read-only. It separates work into four explicit stages:
Identify every affected workload and obtain approval for the impact.
Apply only approved changes, with rollback commands and live verification.
The bundled audit script uses Databricks system tables and inventory APIs. It does not edit warehouses, stop Apps, pause Jobs, change permissions, or create budgets.
Databricks CLI 0.229 or newer; the live test used 0.296
A configured Databricks CLI profile
Access to a running SQL warehouse
Permission to read the relevant system.billing , system.query , system.lakeflow , and system.compute tables
Account-admin access gives the broadest audit, but the toolkit reports whatever the current identity is allowed to see.
git clone https://github.com/kylehuirevvision/databricks-cost-optimizer.git
cd databricks-cost-optimizer
codex
Then invoke:
$optimize-databricks-costs audit my Databricks account from a cost perspective
Codex loads repository skills from .agents/skills according to the official OpenAI skill documentation .
git clone https://github.com/kylehuirevvision/databricks-cost-optimizer.git
cd databricks-cost-optimizer
claude
Then invoke:
/optimize-databricks-costs audit my Databricks account from a cost perspective
Claude Code loads project skills from .claude/skills according to the official Claude Code skill documentation .
Run the evidence collector directly
databricks warehouses list
Run the read-only audit:
python3 .agents/skills/optimize-databricks-costs/scripts/audit.py \
--warehouse-id YOUR_WAREHOUSE_ID \
--recent-days 14 \
--output-dir reports/latest
Use a non-default CLI profile with --profile PROFILE . Preview every SQL statement without connecting to Databricks with --dry-run .
The command writes JSON evidence and a starter Markdown summary under reports/ . That directory is ignored by Git because reports commonly contain account IDs, emails, query metadata, and resource names.
If any system-table query fails, the collector still writes the evidence it obtained but exits nonzero and marks the summary incomplete. Missing evidence is never reported as zero cost.
Sanitized Databricks billing view. August 10 is a partial day and is excluded from the comparison below.
In the first account where this workflow was used, completed daily effective-list-price usage fell from $343 on August 5 to $108 on August 9 , a 68.5% reduction . SQL usage fell from $257 to $42 , an 83.7% reduction , and accounted for most of the measured improvement.
This is an observed before-and-after result, not a controlled experiment or a universal savings claim. Workload volume can vary, recent billing can be partial or restated, and negotiated invoice or underlying cloud-infrastructure costs may differ.
An anonymized engagement found tiny cached queries running approximately every 7.5 minutes against a warehouse with a 10-minute auto-stop. The queries performed almost no compute, but their cadence kept the warehouse billable for every hour in a two-week window. After stopping the recurrence, reducing warehouse size, shortening auto-stop, and stopping unused Apps, the following day's cost dropped materially.
See the sanitized case study for the reasoning pattern and caveats.
python3 -m unittest discover -s tests -v
python3 scripts/sync_skill.py --check
python3 scripts/check_public_safety.py
Maintainers who have Codex's built-in skill-creator may additionally run its quick_validate.py against both skill directories. Public users do not need that internal helper.
The audit estimates USD cost using the effective list price in system.billing.list_prices . This is not necessarily the amount on a negotiated invoice. Recent billing records can arrive late or be restated, so equal-period comparisons exclude the current day and a configurable ingestion-lag window by default.
Read-only-first Databricks FinOps toolkit and Agent Skill for Codex and Claude Code
Readme MIT license Contributing
Security policy Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
