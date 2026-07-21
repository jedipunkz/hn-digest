---
source: "https://github.com/Hedde/trigger_tree"
hn_url: "https://news.ycombinator.com/item?id=48990020"
title: "Show HN: Trigger-tree; see which project docs your AI coding agent reads"
article_title: "GitHub - Hedde/trigger_tree: Documentation-discovery telemetry for Claude Code — heat/cold maps, health grade, evidence-backed router fixes. 100% local, zero tokens. /tt · GitHub"
author: "DevAutomationCC"
captured_at: "2026-07-21T10:15:37Z"
capture_tool: "hn-digest"
hn_id: 48990020
score: 1
comments: 0
posted_at: "2026-07-21T09:34:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Trigger-tree; see which project docs your AI coding agent reads

- HN: [48990020](https://news.ycombinator.com/item?id=48990020)
- Source: [github.com](https://github.com/Hedde/trigger_tree)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T09:34:34Z

## Translation

タイトル: HN を表示: Trigger-tree; AI コーディング エージェントがどのプロジェクト ドキュメントを読むかを確認する
記事のタイトル: GitHub - Hedde/trigger_tree: クロード コードのドキュメント検出テレメトリ — ヒート/コールド マップ、健全性グレード、証拠に裏付けられたルーターの修正。 100% ローカル、トークンゼロ。 /tt · GitHub
説明: クロード コードのドキュメント検出テレメトリ — ヒート/コールド マップ、健全性グレード、証拠に裏付けられたルーターの修正。 100% ローカル、トークンゼロ。 /tt - ヘッデ/trigger_tree

記事本文:
GitHub - Hedde/trigger_tree: クロード コードのドキュメント検出テレメトリ — ヒート/コールド マップ、健全性グレード、証拠に裏付けられたルーターの修正。 100% ローカル、トークンゼロ。 /tt · GitHub
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
{{ メッセージ }}
ヘッデ
/
トリガーツリー

公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
87 コミット 87 コミット .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .codex-plugin .codex-plugin .github .github アセット アセット フック フック スクリプト スクリプト スキル スキル テスト テスト .coveragerc .coveragerc .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md SKILL.md SKILL.md Index.htmlindex.html package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml 要件-dev.txt 要件-dev.txt 要件-test.txt 要件-test.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI が実際にどのドキュメントを検出するかを確認します。
AI コーディング アシスタントはプロジェクトのドキュメントを読んで、作業方法を決定します。
トリガーツリーは、どのドキュメントを発見して使用するのか、そしてどのドキュメントを決して発見しないのかを示します
見つけます。 100%地元産。ゼロモデルトークン。クラウドも分析ベンダーもありません。
✓ ドキュメントのヒート マップとコールド マップ · ✓ ライブ パルス ダッシュボード · ✓ 証拠に裏付けられたルーターの修正
ウェブサイト · プライバシー ポリシー · 変更履歴
なぜドキュメントを測定するのでしょうか? · クイックスタート · コマンド
仕組み · 改善ループ · ダッシュボード
発見のためにドキュメントを構造化する
構成 · プライバシーとデータ
プラットフォームのサポート · FAQ · 開発
AI 支援開発では、ドキュメントはもはや人間だけのものではありません。
ステアリングもCLAUDE.md、規約、およびアーキテクチャのドキュメントには、
チームがソフトウェアを構築する方法のアシスタント: どのパターンをコピーするか、どのガードレールを作成するか
すでに行われた決定を尊重すること。とき

アシスタントは右を読みます
先生、あなたのやり方でうまくいきます。そうでない場合は、推測します。
そして、ここに不快な部分があります。決して読まれないルールは何も保護しません。
チームはドキュメントの作成に多大な投資を行ってから、機能していると思い込んでいますが、未読の状態です。
ガードレールは静かに故障します。 AI が規則を「無視」した場合にのみ、次のようなことに気づくことができます。
実際のところ、それは決して見つかりませんでした。
トリガーツリーはそのループを閉じます。どのドキュメントが実際に参照されているかを基準ごとに測定します。
タスクでは、決して存在しないタスクが表示されます (そして、その理由は、ルーティングされていない? 参照されていない? 廃止されている?)、
そして修正がうまくいったかどうかを証明します。ドキュメントは希望に満ちた成果物ではなくなる
そして、スプリントをまたいで追跡する健全性グレードを備えた、監視対象のインフラストラクチャになります。
この測定値は他には存在しません。アントロピック独自の
ベスト プラクティス ガイドでは、次のように警告しています。
CLAUDE.md が肥大化すると指示が無視されるため、パフォーマンスの高いチームがレビューを行う
容赦なく文脈を変える。しかし、それらの変更が正しかったかどうかは何もわかりません。エージェント
可観測性プラットフォーム (Langfuse、Arize、W&B Weave) はトークンとトレースを測定します。
タスクごとにどのプロジェクト ドキュメントが読まれたかを測定するものはありません。そしてちょうど同じように
Fallow は未使用コードのレビューを証拠に基づいて行いました。
トリガーツリーは、ドキュメントをレビュー、保護、再ルーティング、または救出するための証拠を提供します
大胆な裁判官。
/tt スラッシュ コマンド、Claude フック、およびオプションの Claude にはこのパスを使用します。
ステータスライン:
/プラグインマーケットプレイス追加Hedde/trigger_tree
/plugin install トリガーツリー@trigger-tree
/tt setup # プロジェクトを接続します。デフォルトでローカルの 200 文字のプロンプト プレビューが表示される
/tt ドクター # このリポジトリが有線でテレメトリを受信していることを証明します
いくつかのセッションでは正常に動作し、フックはサイレントにログに記録されます。その後、次のようになります。
/tt ステータス # スナップショット: 現在の熱、有効期間読み取り、未処理のパス
/tt 洞察 # 完全なヒート/コールド マップ レポート + HTML
コーデックス
Codex ライフサイクル フックにはこのパスを使用し、

自然言語トリガーツリーのワークフロー:
コーデックス プラグイン マーケットプレイス Hedde/trigger_tree を追加
コーデックスプラグイン追加trigger-tree@trigger-tree
新しい Codex スレッドを開始し、 /hooks でバンドルされているフックを確認して信頼してから、質問してください。
「トリガー ツリーのステータスを表示する」または「トリガー ツリーのライブ ダッシュボードを開く」。テレメトリーは
公式の Codex ライフサイクル フックから自動的に収集されます。ラッパーは必要ありません。
GitHub インストールと OpenAI Curated の比較: 上記のコマンドは、trigger-tree をインストールします
直接; OpenAI のパブリック ディレクトリには追加しません。 OpenAI が厳選したリスト
を介してスキルのみを個別に提出する必要があります。
OpenAI プラグイン ポータル、続いてレビュー、
承認、および明示的な公開ステップ。
Claude Code コマンドと Codex ワークフロー
Claude Code は、9 つのサブコマンドを含む 1 つのコマンドを公開します。 Codex で、一致を要求します
自然言語での結果。バンドルされたトリガー ツリー スキルは、同じローカル コアを実行します。
ヒントは意図的にクライアント固有のものです。クロードのアドバイスは、Anthropic のガイダンスに従います。
自動メモリを監査し、CLAUDE.md を簡潔に保ち、競合する命令を削除します。
Codex のアドバイスは、OpenAI のガイダンスに従っています。
AGENTS.md と再現可能な開発環境を維持します。
トリガーツリーは 3 つの軽量フックを登録します — 完全な透明性:
シェル側のログを $PROJECT/.trigger-tree/history.jsonl にフックします — ゼロモデル
トークン、ツール呼び出しごとに数ミリ秒、ログ記録の失敗は決して中断されません
セッション (ロガーは常に 0 で終了します)。
アグリゲーターは決定的です。すべてのカウントは次の手順で行われます。
tt-stats.py ;モデルは解釈するだけで、決して計算しません。
Discovery はモデル駆動型のままです。CLAUDE.md はルーターのままです。トリガーツリー
それを測定します。コンテキストを挿入したり、ルーティングをオーバーライドしたりすることはありません。 Bash の検索数
rg 、 grep 、または find が既存のドキュメントを明示的にターゲットにしている場合のみ
p

ああ、検索出力は読み取りとして扱われることはありません。 Bash セッションでは、軽量
リーダー ラッパーは、変数、置換、
ループとグロブが解決されます。コマンドの動作を保持し、一致するもののみを記録します。
コマンド、パターン、出力、コンテンツではなく、パスです。他のシェルは保守的なものを使用します
リテラルパスフォールバック。
サブエージェントの読み取りには属性が付けられます ( Explore 、 Plan など)。自動ロードされたコンテキスト (以下を含む)
再帰的な CLAUDE.md @import グラフ - 設計により読み取りテレメトリには表示されません
常にロードされているものとして分類されます。発動したスキルが測定されます。
読み取り数がゼロのファイルはレビューの候補となり、削除が推奨されることはありません。
保護されたコンテキスト (常にロードされるファイル、安全なパス、重要なタグ/グロブ、およびドキュメント)
多くのインリンクがある) は、可能性の高いキープとして呼び出されます。その後、ループが閉じます。
/tt の洞察は、フォルダーのヒート マップとコールド マップを表示し、ルーター ギャップにフラグを立てます。
他のドキュメントにリンクすらされていない、手付かずのファイル。
/tt の提案により、それが具体的な編集に変わります (「次の docs/README.md に X を追加します)」
Y — 手つかず、参照されていない、3 週間で 0 件の読み取り")。
適用してから、 /tt "sharpened UX router" をメモします。
傾向 (週ごとの狩猟率) は、変更が実際に機能したかどうかを示します。
推測ではなく測定されます。
/tt watch は、2 番目の端末でライブビューを開きます。読み取るたびに白く点滅し、
親フォルダーに波紋が広がり、その後フェードバックして熱色に戻ります。
⠹ トリガーツリー myproject · ライブ ドキュメントディスカバリー
docs/design/ · 1 未読
§─原則.md ███・h 3.2・12×
§─ ui-patterns.md █████ h 6.8 · 17×
└─ アクセシビリティ.md · 0
docs/database/ · 🔍 2 件の検索 · 1 件の未読
└─ 移行.md ██··· h 1.4 · 4×
33 回の読み取り · 2 回のスキャン (狩猟) · 1 回のスキル使用 · 3 回のセッション
● docs/design/ui-patterns.md · 2 秒前
🔍 ドキュメント/データベース [探索]

· 31秒前
熱と読み取りカウントは意図的に異なる信号です。読み取りは寿命です
証拠があり、決して減少しません。 Heat は現在の注目です。各タイムスタンプ付き読み取りには
体重は 0.5^(age_days / 30) なので、その寄与は 30 日ごとに半分になります (今日は 1、
30 日後には 0.5、90 日後には 0.125、1 年後には実質的にゼロになります)。新しい
read はファイルをすぐに再加熱します。 /tt の分析情報では、正確な 7-、30-、および
90 日間の読み取りウィンドウ、最終読み取り日、および有効期間読み取り。フォルダーの熱量は合計です
ファイルの熱の影響を受けます。したがって、寒いということは、現在活動が停止していることを意味し、決して時代遅れではなく、安全に使用できることを意味します。
削除します。未変更の分類と保護されたコンテキストの分類は、別個の保護手段として残ります。
フォルダー ラベルは 2 つの信号を分離します: 🔍 N 回の検索により、フォルダーが
N unread は Read イベントのないファイルをカウントします。検索中
フォルダーは、そのファイルが参照されたふりをすることはありません。 1 を読むと、未読のみが下がります。
←/→ を使用して参照すると、同じカウンタの範囲が選択したプロンプトに限定されます。
ライブ ツリーは 5 秒ごとにインベントリを更新します: 削除されたファイルとフォルダ
現在の概要からは消えますが、その証拠は引き続き利用可能です
過去のプロンプトブラウジングと集計傾向。
最近アクティブなフォルダーが一時的に静かなフォルダーの上に移動するため、ライブ作業はそのまま残ります
小さなビューポート内。 8時以降はアルファベット順に戻ります
秒。フォルダー内のファイルとプロンプト履歴ビューはアルファベット順に安定しています。
ライブビューには、アクティビティが証明されたフォルダーが最大 10 個表示され、そのまま折りたたまれます。
フォルダー/ファイルを 1 つの静かな要約にまとめます。 /tt の洞察は完全なコールド パスのままです
在庫;基礎となるテレメトリからは何も削除されません。
ライブ列は水平 5 セル ヒート バーを使用し、その熱/寿命を配置します。
列が利用可能な右端に対して配置されるため、ペインが広くなります

長いファイル名をさらに公開する
未使用のスペースを残す代わりに。並べ替えは明示的です: f は最近のフォーカスを復元します。
h は最もホットなものを最初に示し、c は最もコールドなものを最初に示します (タッチされていないファイルを含む)。
n は、A ～ Z と Z ～ A を切り替えます。常にロードされる CLAUDE.md 、 AGENTS.md 、ルール、およびスキル
風邪と誤って伝えられるのではなく、注射済みと表示されます。現在のモードは
常にフッターに印刷されます。ダッシュボード内のプロンプトプライバシーを変更するには、s を押します。
変更は今後のプロンプトに適用され、gitignored プロジェクト構成にアトミックに書き込まれます。
ライブ ダッシュボードでは、30 秒ごとに 1 つの静かなメンテナンス ヒントもローテーションします。ヒントは
リポジトリ対応およびクライアント固有: クロードにはメモリ/ルールのガイダンスが表示されますが、コーデックスには次のようなガイダンスが表示されます。
AGENTS.md および検証ガイダンス。テレメトリ ステータスライン自体は安定したままです。
コントロールは、独自の永続的な凡例行を占有します (狭い場合はコンパクトなフォームを使用)
ペイン)、プロンプト ナビゲーションやライブ ハートビートから分離されているため、キーはそのまま残ります。
右端で消えるのではなく、発見可能になります。
プロンプトごとに参照: ←を押して古いプロンプトに移動し、→を押して新しいプロンプトに移動します
1 - ツリーは、その入力に対して集計されたものを正確にフィルタリングします (読み取り、
スキャンとスキルの使用、ヘッダーにプロンプト ラベルが付いています)。タイムラインが決して折り返らない、または
端でモードを変更します。ライブ概要に戻ります。
/tt セットアップのデフォルトは TT_LOG_PROMPTS='trun

[切り捨てられた]

## Original Extract

Documentation-discovery telemetry for Claude Code — heat/cold maps, health grade, evidence-backed router fixes. 100% local, zero tokens. /tt - Hedde/trigger_tree

GitHub - Hedde/trigger_tree: Documentation-discovery telemetry for Claude Code — heat/cold maps, health grade, evidence-backed router fixes. 100% local, zero tokens. /tt · GitHub
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
Hedde
/
trigger_tree
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
87 Commits 87 Commits .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .codex-plugin .codex-plugin .github .github assets assets hooks hooks scripts scripts skills skills tests tests .coveragerc .coveragerc .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md SKILL.md SKILL.md index.html index.html package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml requirements-dev.txt requirements-dev.txt requirements-test.txt requirements-test.txt View all files Repository files navigation
See which docs your AI actually discovers.
AI coding assistants read your project documentation to decide how to work.
trigger-tree shows which docs they discover and use — and which ones they never
find. 100% local. Zero model tokens. No cloud, no analytics vendors.
✓ heat & cold maps of your documentation · ✓ live pulse dashboard · ✓ evidence-backed router fixes
Website · Privacy policy · Changelog
Why measure documentation? · Quick start · Commands
How it works · The improvement loop · The dashboard
Structuring your docs for discovery
Configuration · Privacy & data
Platform support · FAQ · Development
In AI-assisted development, documentation is no longer just for humans — it is the
steering wheel . Your CLAUDE.md, conventions, and architecture docs tell the
assistant how your team builds software: which patterns to copy, which guardrails
to respect, which decisions were already made. When the assistant reads the right
doc, it works your way. When it doesn't, it guesses.
And here is the uncomfortable part: a rule that is never read protects nothing.
Teams invest heavily in writing docs, then assume they work — but an unread
guardrail fails silently. You only notice when the AI "ignores" a convention that,
in truth, it simply never found.
trigger-tree closes that loop. It measures which docs are actually consulted per
task, surfaces the ones that never are (and why — unrouted? unreferenced? obsolete?),
and proves whether your fixes worked. Documentation stops being a hopeful artifact
and becomes monitored infrastructure — with a health grade to track sprint over sprint.
This measurement doesn't exist anywhere else. Anthropic's own
best-practices guide warns that a
bloated CLAUDE.md causes instructions to be ignored — so high-performing teams review
context ruthlessly. But nothing tells them whether those changes were right . Agent
observability platforms (Langfuse, Arize, W&B Weave) measure tokens and traces;
none measure which project docs were read per task. And just as
Fallow made unused-code review evidence-based;
trigger-tree gives you evidence to review, protect, reroute, or rescue docs nobody
dared judge.
Use this path for /tt slash commands, Claude hooks, and the optional Claude
statusline:
/plugin marketplace add Hedde/trigger_tree
/plugin install trigger-tree@trigger-tree
/tt setup # wires the project; local 200-character prompt previews by default
/tt doctor # proves this repo is wired and receiving telemetry
Work normally for a few sessions — the hooks log silently — then:
/tt status # snapshot: current heat, lifetime reads, untouched paths
/tt insights # full heat/cold map report + HTML
Codex
Use this path for Codex lifecycle hooks and natural-language trigger-tree workflows:
codex plugin marketplace add Hedde/trigger_tree
codex plugin add trigger-tree@trigger-tree
Start a new Codex thread, review and trust the bundled hooks with /hooks , then ask
“Show trigger-tree status” or “Open the trigger-tree live dashboard.” Telemetry is
collected automatically from official Codex lifecycle hooks; no wrapper is required.
GitHub install versus OpenAI Curated: the commands above install trigger-tree
directly; they do not add it to OpenAI’s public directory. An OpenAI Curated listing
requires a separate skills-only submission through the
OpenAI plugin portal , followed by review,
approval, and an explicit publish step.
Claude Code commands and Codex workflows
Claude Code exposes one command with nine subcommands. In Codex, ask for the matching
outcome in natural language; the bundled trigger-tree skill runs the same local core.
Tips are intentionally client-specific. Claude advice follows Anthropic's guidance to
audit auto memory, keep CLAUDE.md concise, and remove conflicting instructions .
Codex advice follows OpenAI's guidance to
maintain AGENTS.md and a reproducible development environment .
trigger-tree registers three lightweight hooks — full transparency:
Hooks log shell-side to $PROJECT/.trigger-tree/history.jsonl — zero model
tokens, a few milliseconds per tool call, and a logging failure can never break
your session (loggers always exit 0).
The aggregator is deterministic — all counting happens in
tt-stats.py ; the model only interprets, never computes.
Discovery stays model-driven — your CLAUDE.md remains the router. trigger-tree
measures it; it never injects context or overrides routing. Bash searches count
only when rg , grep , or find explicitly targets an existing documentation
path; search output is never treated as a read. In Bash sessions, lightweight
reader wrappers observe expanded file arguments after variables, substitutions,
loops, and globs resolve. They preserve command behavior and record only matching
paths—not commands, patterns, output, or contents. Other shells use the conservative
literal-path fallback.
Subagent reads are attributed ( Explore , Plan , …). Auto-loaded context—including
the recursive CLAUDE.md @import graph—is invisible to Read telemetry by design
and classified as always loaded ; invoked skills are measured.
Files with zero reads are review candidates , never removal recommendations.
Protected context (always-loaded files, safety paths, critical tags/globs, and docs
with many in-links) is called out as likely-keep. Then the loop closes:
/tt insights shows the folder heat & cold map and flags router gaps :
untouched files that no other doc even links to.
/tt suggestions turns that into concrete edits ("add X to docs/README.md under
Y — untouched and unreferenced, 0 reads in 3 weeks").
You apply, then /tt note "sharpened UX router" .
The trend (hunting ratio per week) shows whether the change actually worked —
measured, not guessed.
/tt watch opens a live view in a second terminal. Every read flashes white and
ripples up through its parent folders, then fades back to its heat color:
⠹ trigger-tree myproject · live doc-discovery
docs/design/ · 1 unread
├─ principles.md ███·· h 3.2 · 12×
├─ ui-patterns.md █████ h 6.8 · 17×
└─ accessibility.md · 0
docs/database/ · 🔍 2 searches · 1 unread
└─ migrations.md ██··· h 1.4 · 4×
33 reads · 2 scans (hunting) · 1 skill uses · 3 sessions
● docs/design/ui-patterns.md · 2s ago
🔍 docs/database [Explore] · 31s ago
Heat and read count are deliberately different signals. Reads are the lifetime
evidence and never decrease. Heat is current attention: each timestamped read has
weight 0.5^(age_days / 30) , so its contribution halves every 30 days (1 today,
0.5 after 30 days, 0.125 after 90 days, and effectively zero after a year). A new
read reheats the file immediately. /tt insights also shows exact 7-, 30-, and
90-day read windows, the last-read date, and lifetime reads. Folder heat is the sum
of its file heat. Cold therefore means inactive now , never obsolete or safe to
remove; untouched and protected-context classifications remain separate safeguards.
Folder labels keep two signals separate: 🔍 N searches proves the folder was
explicitly searched, while N unread counts files without a Read event. Searching
a folder never pretends its files were consulted; reading one lowers only unread .
The same counters are scoped to the selected prompt when browsing with ←/→.
The live tree refreshes its inventory every five seconds: deleted files and folders
disappear from the current overview while their evidence remains available in
historical prompt browsing and aggregate trends.
Recently active folders temporarily move above quiet folders so live work stays
inside a small viewport. They settle back into alphabetical order after eight
seconds; files within folders and prompt-history views remain alphabetically stable.
The live view shows at most ten folders with proven activity and collapses untouched
folders/files into one quiet summary. /tt insights remains the complete cold-path
inventory; nothing is removed from the underlying telemetry.
The live rows use horizontal five-cell heat bars and place their heat/lifetime
column against the available right edge, so wider panes expose more of long filenames
instead of leaving unused space. Sorting is explicit: f restores recent-focus,
h shows hottest first, c shows coldest first (including untouched files), and
n toggles A–Z and Z–A. Always-loaded CLAUDE.md , AGENTS.md , rules, and skills
are labeled injected instead of being misrepresented as cold. The current mode is
always printed in the footer. Press s to change prompt privacy inside the dashboard;
changes apply to future prompts and are written atomically to the gitignored project config.
The live dashboard also rotates one quiet maintenance tip every 30 seconds. Tips are
repository-aware and client-specific: Claude sees memory/rules guidance, while Codex sees
AGENTS.md and verification guidance. The telemetry statusline itself remains stable.
The controls occupy their own persistent legend row (with a compact form for narrow
panes), separate from prompt navigation and the live heartbeat, so the keys remain
discoverable instead of disappearing at the right edge.
Browse per prompt : press ← to move to older prompts and → to move to newer
ones — the tree filters to exactly what was aggregated for that input (its reads,
scans and skill uses, with its prompt label in the header). The timeline never wraps or
changes mode at its ends; a returns to the live overview.
/tt setup defaults to TT_LOG_PROMPTS='trun

[truncated]
