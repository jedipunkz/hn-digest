---
source: "https://github.com/Treasury-Technologies-Inc/treasurybench"
hn_url: "https://news.ycombinator.com/item?id=48676009"
title: "Show HN: TreasuryBench – an open benchmark for personal-finance AI advice"
article_title: "GitHub - Treasury-Technologies-Inc/treasurybench: Personal-finance assistant benchmark — evaluate real finance products against synthetic user personas · GitHub"
author: "juneadkhan"
captured_at: "2026-06-25T16:44:59Z"
capture_tool: "hn-digest"
hn_id: 48676009
score: 1
comments: 0
posted_at: "2026-06-25T16:43:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: TreasuryBench – an open benchmark for personal-finance AI advice

- HN: [48676009](https://news.ycombinator.com/item?id=48676009)
- Source: [github.com](https://github.com/Treasury-Technologies-Inc/treasurybench)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T16:43:08Z

## Translation

タイトル: Show HN: TreasuryBench – 個人金融 AI アドバイスのオープン ベンチマーク
記事のタイトル: GitHub - Treasury-Technologies-Inc/treasurybench: 個人金融アシスタント ベンチマーク — 合成ユーザー ペルソナに対して実際の金融商品を評価する · GitHub
説明: 個人金融アシスタント ベンチマーク — 合成ユーザー ペルソナに対して実際の金融商品を評価する - Treasury-Technologies-Inc/treasurybench

記事本文:
GitHub - Treasury-Technologies-Inc/treasurybench: 個人金融アシスタント ベンチマーク — 合成ユーザー ペルソナに対して実際の金融商品を評価する · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
財務-T

テクノロジーズ株式会社
/
財務ベンチ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
Treasury-Technologies-Inc/トレジャリーベンチ
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット アーティファクト アーティファクト docs docs fixtures fixtures src src .gitignore .gitignore LICENSE LICENSE LIMITATIONS.md LIMITATIONS.md METHODOLOGY.md METHODOLOGY.md README.md README.md SCORING.md SCORING.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
個人金融アシスタントのベンチマーク — AI を活用した金融商品とフロンティア モデルが実際のユーザー データをどのように活用して、高レバレッジの金融機会を表面化するかを評価します。
v0.1.0 · 3 つのペルソナ · 81 のタスク · 12 のドメイン · 表に基づいた事実検証による審査員一次スコアリング
プロバイダー
レーン
スコア
事実上クリーン
レイテンシの中央値
財務省
製品候補
85.5
93%
13.7秒
ChatGPT チャット-最新
フルコンテキストベースライン †
79.6
83%
8.0秒
起源
製品候補
71.0‡
86%
46.0秒
君主
製品候補
52.1
86%
100.7秒
† フルコンテキスト ベースラインは、ペルソナの取引、残高、記憶をプロンプトに直接貼り付けます。これは、実際の消費者向け製品の仕組みではありません。これは上限の見積もりであり、製品の候補ではありません。
‡ 残高インポートがサイレントに失敗した 16 個のタスクを除いた場合は 73.1。 artifacts/RUN_INTEGRITY.md を参照してください。
スコアは 0 ～ 100 で、裁判官が優先し、テーブルに基づいた事実に基づいた上限が適用されます。財務上の事実 (拠出限度額、税金規則、プログラム条件) が古い、または間違っている場合、散文の質に関係なくタスク スコアにハードキャップが設定されます。重要なエラーの上限は 65、危険なエラーの上限は 40 です。フル スコアリング アーキテクチャ: SCORING.md 。
行ごとの最高スコアは太字で表示されます。 † はフルコンテキスト b をマークします。

aseline (製品の候補ではありません)。
ペルソナ
財務省
起源
君主
ChatGPT †
マリア・チェン — シアトル、マイクロソフト、賃貸人
87
71
57
80
プリヤ・パテル — デンバー、共働き、住宅所有者
85
69
46
70
ジョーダン・リベラ — オースティン、自営業
84
73
53
89
事実の誠実さ
81 のタスクにわたる、事実に矛盾がない回答の割合。危険 = 実際の経済的損害を引き起こす可能性のある誤った事実 (例: 実用的なアドバイスとして引用されている古い拠出制限)。
ChatGPT の 12 の危険なエラーは、その評価された品質 (85) と最終スコア (79.6) の間の最大のギャップを引き起こします。理想化されたプロンプトのコンテキストの下でも、ChatGPT は一貫して、古い 2025 年の貢献制限を現行のものとして引用しています。
すべてのキャプチャ、ジャッジプロンプト、判定、採点結果はアーティファクト/ にあります。
各実行ディレクトリには、captures/ 、judge-prompts/ 、judgings/ 、results/ が含まれており、機械可読な CSV と相違レポートが含まれます。裁判官の独立性に関する警告、Origin インポート失敗の開示、および自己著作者の開示については、artifacts/RUN_INTEGRITY.md を参照してください。
財務省は、個人財務アシスタントが次のことができるかどうかを尋ねます。
取引データと残高データを正確に読み取ります。
ユーザーのコンテキストを個人の財務概念に結び付けます。
通常の財務データに隠れている高価値の機会を明らかにします。
現在の財務ルール、制限、商品条件、およびローカル プログラムを正しく使用してください。
影響を定量化し、次のステップを正確に示します。
サポートされていない仮定、古い事実、安全でない推奨事項、一般的な定型文は避けてください。
取引履歴、口座残高、保存された思い出、雇用主、場所、目標を含む 3 つの米国の合成世帯:
Maria Chen — 20 代後半、シアトル、マイクロソフト ソフトウェア エンジニア、賃貸勤務。
プリヤ・パテル — 共働き、デンバー、住宅所有者、子供 2 人。
Jordan Rivera — オースティン、自営業、ギグ/フリーランスの収入。
81 ユーザーの自然な質問

12 ドメインにわたる ns (ペルソナあたり 27)。タスクは、「どうすれば家賃を節約できますか?」という実際のユーザーの質問のように表現されています。 「シアトルの MFTE 資格を確認する」ではありません。アシスタントはペルソナの信号から機会を推測する必要があります。
LLM 判定出力が使用可能な場合は、判定プライマリ。決定論的評価者は、正確なデータの使用、演算、および植え付けられた信号の発見を捕捉します。 LLM 審査員は、総合、パーソナライゼーション、および無制限の単位を採点します。陳腐または誤った財務事実には、散文の質に関係なく厳しい制限が適用されます。
完全なアーキテクチャ: SCORING.md · 方法論: METHODOLOGY.md · 制限事項: LIMITATIONS.md · 実行整合性: artifacts/RUN_INTEGRITY.md
pnpmインストール
pnpm validate # スキーマの一貫性とスコア合計を検証します
pnpm report # コンパクトなタスク/ドメインの概要を出力します
pnpmスモーク #フィクスチャプロバイダーをエンドツーエンドで実行します
フルコンテキストベースラインを実行する
pnpm エクスポート プロンプト -- --out=runs/my-openai-run/prompts --mode=full_context_baseline
pnpm run-provider -- --provider=openai --out=runs/my-openai-run --live=true \
--model=チャット最新 --max-output-tokens=2200 --env-file=.env
pnpm Evaluate-run -- --run=runs/my-openai-run
pnpm run-judge -- --run=runs/my-openai-run --env-file=.env \
--judge-provider=gemini --model=gemini-3.1-flash-lite
pnpm スコア-run -- --run=runs/my-openai-run
.env には OPENAI_API_KEY (プロバイダー) と GOOGLE_GENERATIVE_AI_API_KEY (ジャッジ) が必要です。代わりに OpenAI で判断するには、--judge-provider=openai を OPENAI_API_KEY とともに使用します。
pnpm エクスポートペルソナデータ -- --out=runs/my-product-data
pnpm make-capture-templates -- --out=runs/my-product --provider=myproduct --mode=product_capture
# 各ペルソナを製品にシードし、自然なプロンプトを尋ね、回答を貼り付けます
# 各captures/*.jsonファイルの`response`フィールドに
pnpmevaluate-run -- --run=runs/my-product
pnpm run-judge -- --run=runs/m

y-product --env-file=.env \
--judge-provider=gemini --model=gemini-3.1-flash-lite
pnpm スコア-run -- --run=runs/my-product
完全なシードプロトコルについては、docs/product-capture-protocol.md を参照してください。
pnpm スコア-run -- --run=artifacts/treasury-full-20260609001842
既存のキャプチャから再判定するには:
pnpm run-judge -- --run=artifacts/treasury-full-20260609001842 --env-file=.env \
--judge-provider=gemini --model=gemini-3.1-flash-lite
ライセンス
個人金融アシスタントのベンチマーク — 合成ユーザー ペルソナに対して実際の金融商品を評価します
tresury.sh/ベンチマーク
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
1
トレジャリーベンチ v0.1.0
最新の
2026 年 6 月 24 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Personal-finance assistant benchmark — evaluate real finance products against synthetic user personas - Treasury-Technologies-Inc/treasurybench

GitHub - Treasury-Technologies-Inc/treasurybench: Personal-finance assistant benchmark — evaluate real finance products against synthetic user personas · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Treasury-Technologies-Inc
/
treasurybench
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Treasury-Technologies-Inc/treasurybench
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits artifacts artifacts docs docs fixtures fixtures src src .gitignore .gitignore LICENSE LICENSE LIMITATIONS.md LIMITATIONS.md METHODOLOGY.md METHODOLOGY.md README.md README.md SCORING.md SCORING.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json View all files Repository files navigation
Personal-finance assistant benchmark — evaluate how well AI-powered finance products and frontier models use real user data to surface high-leverage financial opportunities.
v0.1.0 · 3 personas · 81 tasks · 12 domains · judge-primary scoring with table-grounded factual verification
Provider
Lane
Score
Factually Clean
Median Latency
Treasury
Product contender
85.5
93%
13.7s
ChatGPT chat-latest
Full-context baseline †
79.6
83%
8.0s
Origin
Product contender
71.0 ‡
86%
46.0s
Monarch
Product contender
52.1
86%
100.7s
† Full-context baselines paste the persona's transactions, balances, and memories directly into the prompt — this is not how a real consumer product works. It is a ceiling estimate, not a product contender.
‡ 73.1 when the 16 tasks where balance import silently failed are excluded. See artifacts/RUN_INTEGRITY.md .
Scores are 0–100, judge-primary with table-grounded factual caps. Stale or wrong financial facts (contribution limits, tax rules, program terms) hard-cap the task score regardless of prose quality — material errors cap at 65, dangerous errors at 40. Full scoring architecture: SCORING.md .
Best score per row bolded. † marks the full-context baseline (not a product contender).
Persona
Treasury
Origin
Monarch
ChatGPT †
Maria Chen — Seattle, Microsoft, renter
87
71
57
80
Priya Patel — Denver, dual income, homeowner
85
69
46
70
Jordan Rivera — Austin, self-employed
84
73
53
89
Factual Integrity
Share of answers with no locked-fact contradiction across 81 tasks. Dangerous = incorrect fact that could cause real financial harm (e.g. stale contribution limit cited as actionable advice).
ChatGPT's 12 dangerous errors drive the largest gap between its judged quality (85) and final score (79.6): it consistently cites stale 2025 contribution limits as current, even under idealized in-prompt context.
All captures, judge prompts, judgments, and scored results are in artifacts/ .
Each run directory contains captures/ , judge-prompts/ , judgments/ , and results/ with machine-readable CSVs and divergence reports. See artifacts/RUN_INTEGRITY.md for the judge-independence caveat, the Origin import-failure disclosure, and the self-authorship disclosure.
TreasuryBench asks whether a personal-finance assistant can:
Read transaction and balance data accurately.
Connect user context to personal-finance concepts.
Surface high-value opportunities hidden in ordinary financial data.
Use current financial rules, limits, product terms, and local programs correctly.
Quantify impact and give exact next steps.
Avoid unsupported assumptions, stale facts, unsafe recommendations, and generic boilerplate.
Three synthetic US households with transaction history, account balances, saved memories, employer, location, and goals:
Maria Chen — late 20s, Seattle, Microsoft software engineer, renter.
Priya Patel — dual income, Denver, homeowner, two kids.
Jordan Rivera — Austin, self-employed, gig/freelance income.
81 natural user questions (27 per persona) across 12 domains. Tasks are phrased like real user questions — "How can I save money on rent?" not "Identify Seattle MFTE eligibility." The assistant must infer the opportunity from the persona's signals.
Judge-primary when LLM judge output is available. Deterministic evaluators catch exact data use, arithmetic, and planted-signal discovery. The LLM judge grades synthesis, personalization, and open-ended credit. Stale or wrong financial facts apply hard caps regardless of prose quality.
Full architecture: SCORING.md · Methodology: METHODOLOGY.md · Limitations: LIMITATIONS.md · Run integrity: artifacts/RUN_INTEGRITY.md
pnpm install
pnpm validate # verify schema consistency and scoring totals
pnpm report # print a compact task/domain summary
pnpm smoke # run the fixture provider end-to-end
Run a full-context baseline
pnpm export-prompts -- --out=runs/my-openai-run/prompts --mode=full_context_baseline
pnpm run-provider -- --provider=openai --out=runs/my-openai-run --live=true \
--model=chat-latest --max-output-tokens=2200 --env-file=.env
pnpm evaluate-run -- --run=runs/my-openai-run
pnpm run-judge -- --run=runs/my-openai-run --env-file=.env \
--judge-provider=gemini --model=gemini-3.1-flash-lite
pnpm score-run -- --run=runs/my-openai-run
.env needs OPENAI_API_KEY (provider) and GOOGLE_GENERATIVE_AI_API_KEY (judge). Use --judge-provider=openai with OPENAI_API_KEY to judge with OpenAI instead.
pnpm export-persona-data -- --out=runs/my-product-data
pnpm make-capture-templates -- --out=runs/my-product --provider=myproduct --mode=product_capture
# seed each persona into your product, ask the natural prompt, paste the answer
# into the `response` field of each captures/*.json file
pnpm evaluate-run -- --run=runs/my-product
pnpm run-judge -- --run=runs/my-product --env-file=.env \
--judge-provider=gemini --model=gemini-3.1-flash-lite
pnpm score-run -- --run=runs/my-product
See docs/product-capture-protocol.md for the full seeding protocol.
pnpm score-run -- --run=artifacts/treasury-full-20260609001842
To re-judge from existing captures:
pnpm run-judge -- --run=artifacts/treasury-full-20260609001842 --env-file=.env \
--judge-provider=gemini --model=gemini-3.1-flash-lite
License
Personal-finance assistant benchmark — evaluate real finance products against synthetic user personas
treasury.sh/benchmarks
Resources
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
TreasuryBench v0.1.0
Latest
Jun 24, 2026
Packages
0
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
