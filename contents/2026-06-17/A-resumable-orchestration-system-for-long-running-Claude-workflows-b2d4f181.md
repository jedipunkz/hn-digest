---
source: "https://github.com/afsalali1238/Incubator"
hn_url: "https://news.ycombinator.com/item?id=48566690"
title: "A resumable orchestration system for long-running Claude workflows"
article_title: "GitHub - afsalali1238/Incubator: Idea in, company out. A Claude skill that becomes a domain-expert CEO for your new project: it interviews you, researches the market autonomously, ships a findings report, then hires a sequenced team of specialist AI agents. Installable as a Claude Code plugin. · Git\n[truncated]"
author: "afsalali1238"
captured_at: "2026-06-17T07:59:08Z"
capture_tool: "hn-digest"
hn_id: 48566690
score: 1
comments: 0
posted_at: "2026-06-17T06:45:44Z"
tags:
  - hacker-news
  - translated
---

# A resumable orchestration system for long-running Claude workflows

- HN: [48566690](https://news.ycombinator.com/item?id=48566690)
- Source: [github.com](https://github.com/afsalali1238/Incubator)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T06:45:44Z

## Translation

タイトル: 長時間実行されるクロード ワークフローのための再開可能なオーケストレーション システム
記事のタイトル: GitHub - afsalali1238/インキュベーター: アイデアは入って、会社は出る。新しいプロジェクトのドメインエキスパート CEO となるクロード スキル: あなたにインタビューし、自律的に市場を調査し、調査結果レポートを発送し、専門 AI エージェントの順序付けられたチームを雇用します。クロードコードプラグインとしてインストール可能。 · Git
[切り捨てられた]
説明: アイデアは入る、会社は出る。新しいプロジェクトのドメインエキスパート CEO となるクロード スキル: あなたにインタビューし、自律的に市場を調査し、調査結果レポートを発送し、専門 AI エージェントの順序付けられたチームを雇用します。クロードコードプラグインとしてインストール可能。 - afsalali1238/インキュベーター

記事本文:
GitHub - afsalali1238/Incubator: アイデアは入って、会社は出る。新しいプロジェクトのドメインエキスパート CEO となるクロード スキル: あなたにインタビューし、自律的に市場を調査し、調査結果レポートを発送し、専門 AI エージェントの順序付けられたチームを雇用します。クロードコードプラグインとしてインストール可能。 · GitHub
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
別のタブまたは風でアカウントを切り替えた

うわー。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アフサラリ1238
/
インキュベーター
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット .claude-plugin .claude-plugin アセット アセット コンテンツ content dist dist docs docs サンプル サンプル スクリプト scripts skill/ project-ceo skill/ project-ceo .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md release.sh release.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
インキュベーター — アイデアは入って会社は出る
新しいプロジェクトの創設 CEO として機能するクロード スキル。テスト可能な仮説についてインタビューし、市場を調査し、証拠がそうであれば方向転換するよう指示する調査結果レポートを作成し、順序付けされた専門家チームを生成します。つまり、リサーチから派生したペルソナ ブリーフとインストール可能なスキル ファイルが役割ごとに 1 つずつ作成されます。
Claude Cowork、Claude Code、claude.ai 用に構築されました。
エージェント コーディングを使用すると、気軽に構築できます。それが罠です。午後にプロトタイプを出荷し、それを検証と勘違いしてしまうのです。センスメイキングは構築に後れをとります。
インキュベータは命令を強制します。研究が存在する前にチームを雇用したり、アイデアが反証可能な仮説になる前に構築したりすることはできません。 CEO は、証拠が停止を示している場合にはそれを伝える必要があります。そして 1 週間後に戻ってくると、会社が中断したところから正確に再開されます。
⭐ これが役立つ場合は、リポジトリにスターを付けてください。他の人が見つけやすくなります。
プロジェクトの開始時に project-ceo を呼び出します。クロードは、一般的なアドバイスではなく、実際の業界の専門知識を持って、特定の業界の創設 CEO になります。 7 フェーズのシーケンスを実行します。
CEOは設立後も消えることはありません。 INを書き込みます

ワークスペースのルートにある CUBATOR.md インデックス。 「CEO チェックイン」「プロジェクトに取り組みましょう」と戻ってくるたびに、名簿が再読み込みされ、チームの健全性がチェックされ、会社が中断したところから正確に再開されます。
Examples/airbnb/ フォルダーには、Airbnb の設立に関する完全な実行からの実際の出力が表示されます。
<あなたのプロジェクト>-会社/
§── 00_charter.md # 論文、CEO の評決、最も危険な仮定、組織の概要
§── 01_findings-report.html # 研究レポート (範囲に応じて 5 または 9 パネル)
§── 02_hiring-plan.md # 組織図 + 理由付きの採用順序
§── 04_pitch-deck.md # 10スライドVCピッチデッキ（VCトラックのみ）
§── hq.html # ダークモード コマンド センター — 任意のブラウザで開きます
§── roster.md # ヘルススコアを含む LIVE チームレジストリ — セッションごとに更新
§── チーム/# 人のペルソナ ブリーフ、雇用されたエージェントごとに 1 つ
└── スキル/ インストール可能なスペシャリスト エージェント スキルの数、役割ごとに 1 つ
INCUBATOR.md # ワークスペース ルートのセッション インデックス — CEO がプロジェクトを見つける方法
インストール
/プラグイン マーケットプレイス add afsalali1238/Incubator
/プラグインインストールインキュベーター
次に、「この件の CEO になってください。私は <あなたのアイデア> を構築しています。」
クロード・コワーク / claude.ai (マニュアル)
dist/project-ceo.skill をダウンロードする — リンクをクリックし、[生ファイルのダウンロード] (GitHub の下矢印アイコン) をクリックします。
claude.ai を開く → アバター (左下) をクリック → 設定
[スキル] (左側のサイドバー) に移動し、[スキルの追加] をクリックし、ダウンロードした .skill ファイルを選択します。
スキルは即座にインストールされます。新しい会話を開始して、「新しいプロジェクトを開始します。CEO を務めます」と言います。
clude.ai の [スキル] タブはありませんか?まだスキルをサポートしていないプランを使用している可能性があります。クロード コード (上記を参照) を試すか、skills/project-ceo/SKILL.md の内容をシステム プロンプトとして直接貼り付けます。
または、skills/project-ceo/ フォルダーをクロード コードのスキルにドロップします。

ソースから作業したい場合は、irectory を使用してください。
「新しいプロジェクトを始めます。このプロジェクトの CEO になってください。」
「<X> を構築したいので、その開始を手伝ってください。」
「このビルドのポイントを創設者として実行します。」
「CEO チェックイン — ここはどこですか?」
プロジェクトに名前を付けるだけです。CEO は INCUBATOR.md を見つけて選択します。
セットアップに応じて 2 つのモード:
クロード コード / タスク ツールが利用可能: エージェントは実際のサブエージェントとして実行されます。CEO がエージェントを生成し、独立して動作し、結果を返します。完全なマルチエージェント実行。
Claude Cowork / 標準 claude.ai: エージェントはインライン ペルソナとしてアクティブ化されます。 CEO はそれらに明確なラベルを付け ([アクティブ化: 成長リーダー] ... [CEO に戻る] )、同じ会話の中でそれらを切り替えます。同じドメインの専門知識があり、並列実行はありません。
このスキルはどちらのモードでも役立ちます。違いは速度と並列処理であり、出力の品質ではありません。
インキュベーター/
§── .claude-plugin/
│ §── plugin.json
│ └── マーケットプレイス.json
§── 資産/
│ └── バナー.png
§── スキル/プロジェクトCEO/
│ §── SKILL.md
│ ━─ 参考文献/
│ §── Interview.md # デシジョンツリー + 時間予算の質問
│ §── Research.md # メソッド + 20 サーチハードキャップ + 墓地プロトコル
│ §── org-design.md # 垂直ロスターパターン + シーケンス
│ §── Agent-skill-template.md # 概要 + スキル テンプレート + クオリティ ゲート チェックリスト
│ §── starter-pack.md # 憲章、名簿、INCUBATOR.md 形式
│ §── Orchestration.md # 雇用/解雇/デリゲート/スポーン + 取締役会 + ヘルススコア
│ └─ hq-template.md # HQ ダッシュボード HTML テンプレート
§── 例/
│ §── README.md # すべての例のインデックス
│ §── airbnb/ # ★ リード例 — 7 つの出力ファイルすべて + hq.html
│ §──slack-company/ # B2B メッセージングウェッジ
│ §── uber-company/ # 配車都市

打ち上げ
│ §── Gymshark-company/#DTC フィットネスアパレル
│ §── oculus-company/ # コンシューマ VR ハードウェア
│ └── 23andme-company/ # 消費者 DNA 検査
━── 距離/
└── project-ceo.skill # パッケージ化されたインストール可能
コンパニオン スキル (オプション — これらのスキルがなくても完全にスタンドアロンで動作します)
これらはクロード・コワークの個人的なスキルです。これらをインストールしている場合、CEO はそれらを自動的に使用します。そうでない場合は、同等の動作がフォールバックとして組み込まれます。
Grille-me — 一度に 1 つずつ質問する執拗な面接 (フェーズ 1 フォールバック: 組み込みのデシジョン ツリー)
カルパシーガイドライン — 表面上の仮定、過度の複雑さを避ける（運用原則に組み込まれている）
vd — フェーズ 3 で使用されるビジュアル ドキュメント形式 (フォールバック: インラインで生成される自己完結型 HTML)
autoresearch — 自律ループ規律 (フォールバック: Search N/20: フェーズ 2 のカウンター)
生のプロンプトから賢明な回答が得られます。これにより、永続的な企業が得られます。
フェーズには順序が強制されます。研究が開始されるまでチームを雇用することはできません。 CEO は、センスメイキングの前に構築を許可しません。
研究には厳しい上限と説明責任があります。20 件の検索が大声でカウントされ、明示的な INSUFFICIENT DATA フラグが付けられます。生のプロンプトはギャップを幻覚します。
名簿は証拠に基づいて作成されます。すべての採用者はテンプレートから生成されるのではなく、特定の調査委員会に引用されます。
セッションは継続します — INCUBATOR.md は、プロジェクトについて再度説明することなく、CEO が中断したところから正確に再開することを意味します。
エージェントの品質は制限されています。生成されたすべてのスキルは、インストールする前にコールドスタートの批評家パスを通過します。
このスキルは 1 回のセッションで役立ちます。数週間が経つと、別のツールになります。
長い AI セッションでは状態のドリフトが蓄積されます。モデルはセッションごとにコンテキストからプロジェクトを再構築しますが、その再構築は毎回わずかにドリフトします。 Thr

EE エンジニアリング上の決定はこれに対処します。
フェーズのチェックポイント設定。 INCUBATOR.md は、フェーズ 1 論文が最後ではなく、ロックされた瞬間に書き込まれます。フェーズ 1 とフェーズ 4 の間のセッションのドロップアウトはディスクから回復できます。 session-state.json はフェーズ 2、3、4 の後に記述され、フェーズ、信頼性、未解決のリスク、および 20 単語の概要が含まれます。履歴書では、CEO は会話履歴ではなく、これを最初に読みます。
正規の JSON 状態。 roster.md は、エージェント チームの状態をファイルの下部に JSON ブロックとして保存します。 CEO は最初にその JSON ブロックを読み書きしてから、そこからマークダウン テーブルを再生成します。マークダウンはプレゼンテーションのみです。両者は分岐できません。
コールドスタート品質のゲート。各エージェント スキルを生成した後、CEO は一時停止し、ユーザーからの CRITIQUE トークンを待ちます。これにより、コールドスタート評価命令による新しい世代ターンがトリガーされます。モデルは、スキルを作成していないかのようにスキルを評価する必要があります。インストリーム自己批判 (生成したばかりの同じ出力パスでモデルに評価を求めること) は、本当の評価ではありません。別ターンです。
これらはプロンプトレベルの緩和策であり、実行時の強制ではありません。サブストレートの制限 (真の並列エージェントなし、ランタイム スキーマ強制なし、ライブ KPI アクセスなし) は、完全な論文の記述に文書化されています。
企業シミュレーターではなく、創業時のオペレーティング システムとして構築されました。ここで得られるのは、永続的な状態レイヤーを備えた構造化されたビルド前の調査と計画の儀式です。これは、経験豊富な創業チームがコード行を記述する前に行うような思考であり、ターンベースのセッションに圧縮されています。
モデルは、すべて同じコンテキストから複数のペルソナを演じます。それは計画上の成果物であり、実際のチームではありません。この価値は組織の幻想ではありません。フェーズによって順序が強制されるということです。調査する前に雇用することはできず、調査する前に構築することもできません。

検証済みですが、Scale ロールを取得する前に追加することはできません。構築する前にセンスメイキング。エージェント コーディングを使用すると、気軽に構築できます。それがこの罠を防ぐために設計されたものです。
slavingia/skills から採用された名簿パターン。
アイデアはイン、会社はアウト。新しいプロジェクトのドメインエキスパート CEO となるクロード スキル: あなたにインタビューし、自律的に市場を調査し、調査結果レポートを発送し、専門 AI エージェントの順序付けられたチームを雇用します。クロードコードプラグインとしてインストール可能。
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

Idea in, company out. A Claude skill that becomes a domain-expert CEO for your new project: it interviews you, researches the market autonomously, ships a findings report, then hires a sequenced team of specialist AI agents. Installable as a Claude Code plugin. - afsalali1238/Incubator

GitHub - afsalali1238/Incubator: Idea in, company out. A Claude skill that becomes a domain-expert CEO for your new project: it interviews you, researches the market autonomously, ships a findings report, then hires a sequenced team of specialist AI agents. Installable as a Claude Code plugin. · GitHub
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
afsalali1238
/
Incubator
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits .claude-plugin .claude-plugin assets assets content content dist dist docs docs examples examples scripts scripts skills/ project-ceo skills/ project-ceo .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md release.sh release.sh View all files Repository files navigation
Incubator — idea in, company out
A Claude skill that acts as the founding CEO for a new project — interviewing you into a testable hypothesis, researching the market, writing a findings report that will tell you to pivot if the evidence says so, and generating a sequenced specialist team: persona briefs and installable skill files derived from the research, one per role.
Built for Claude Cowork , Claude Code, and claude.ai.
Agentic coding makes building feel free. That's the trap — you ship a prototype in an afternoon and mistake it for validation. Sense-making falls behind building.
Incubator enforces the order. It won't let you hire a team before the research exists, or build before the idea is a falsifiable hypothesis. The CEO is required to tell you when the evidence says stop. And when you come back a week later, it picks up exactly where the company left off.
⭐ If this is useful, star the repo — it helps others find it.
Invoke project-ceo at the start of any project. Claude becomes the founding CEO for your specific vertical — with real industry expertise, not generic advice. It runs a seven-phase sequence:
The CEO doesn't disappear after founding. It writes an INCUBATOR.md index at your workspace root. Every time you return — "CEO check in", "let's work on the project" — it re-reads the roster, checks team health, and picks up exactly where the company left off.
The examples/airbnb/ folder shows real output from a full run on the founding of Airbnb:
<your-project>-company/
├── 00_charter.md # thesis, CEO verdict, riskiest assumption, org at a glance
├── 01_findings-report.html # research report (5 or 9 panels depending on scope)
├── 02_hiring-plan.md # org chart + sequenced hire order with reasons
├── 04_pitch-deck.md # 10-slide VC pitch deck (VC track only)
├── hq.html # dark-mode command centre — open in any browser
├── roster.md # LIVE team registry with health scores — updated every session
├── team/ # persona briefs, one per hired agent
└── skills/ # installable specialist-agent skills, one per role
INCUBATOR.md # session index at workspace root — how the CEO finds your project
Install
/plugin marketplace add afsalali1238/Incubator
/plugin install incubator
Then: "be the CEO for this — I'm building <your idea>"
Claude Cowork / claude.ai (manual)
Download dist/project-ceo.skill — click the link, then click Download raw file (the down-arrow icon on GitHub)
Open claude.ai → click your avatar (bottom-left) → Settings
Go to Skills (left sidebar) → click Add skill → select the .skill file you downloaded
The skill installs instantly. Start a new conversation and say: "I'm starting a new project, act as CEO"
No claude.ai Skills tab? You may be on a plan that doesn't support skills yet. Try Claude Code (see above) or paste the contents of skills/project-ceo/SKILL.md directly as a system prompt.
Or drop the skills/project-ceo/ folder into your Claude Code skills directory if you prefer working from source.
"I'm starting a new project — be the CEO for this."
"I want to build <X>, help me kick it off."
"Run point on this build as founder."
"CEO check in — where are we?"
Just name the project — the CEO finds INCUBATOR.md and picks up.
Two modes depending on your setup:
Claude Code / Task tool available: agents run as real subagents — the CEO spawns them, they work independently, and return results. Full multi-agent execution.
Claude Cowork / standard claude.ai: agents activate as inline personas. The CEO labels them clearly ( [Activating: Growth Lead] ... [Back to CEO] ) and switches between them in the same conversation. Same domain expertise, no parallel execution.
The skill is useful in both modes. The difference is speed and parallelism, not quality of output.
Incubator/
├── .claude-plugin/
│ ├── plugin.json
│ └── marketplace.json
├── assets/
│ └── banner.png
├── skills/project-ceo/
│ ├── SKILL.md
│ └── references/
│ ├── interview.md # decision tree + time-budget question
│ ├── research.md # method + 20-search hard cap + graveyard protocol
│ ├── org-design.md # vertical roster patterns + sequencing
│ ├── agent-skill-template.md # brief + skill templates + quality gate checklist
│ ├── starter-pack.md # charter, roster, INCUBATOR.md format
│ ├── orchestration.md # hire/fire/delegate/spawn + board meeting + health scores
│ └── hq-template.md # the HQ Dashboard HTML template
├── examples/
│ ├── README.md # index of all examples
│ ├── airbnb/ # ★ lead example — all 7 output files + hq.html
│ ├── slack-company/ # B2B messaging wedge
│ ├── uber-company/ # ride-hailing city launch
│ ├── gymshark-company/ # DTC fitness apparel
│ ├── oculus-company/ # consumer VR hardware
│ └── 23andme-company/ # consumer DNA testing
└── dist/
└── project-ceo.skill # packaged installable
Companion skills (optional — works fully standalone without any of these)
These are personal Claude Cowork skills. If you have them installed, the CEO uses them automatically. If not, equivalent behaviour is built in as a fallback.
grill-me — relentless one-question-at-a-time interviewing (Phase 1 fallback: built-in decision tree)
karpathy-guidelines — surface assumptions, avoid overcomplication (built into operating principles)
vd — Visual Document format used in Phase 3 (fallback: self-contained HTML generated inline)
autoresearch — autonomous-loop discipline (fallback: Search N/20: counter in Phase 2)
A raw prompt gets you a smart answer. This gets you a persistent company .
Phases enforce order — you can't hire a team before the research exists. The CEO won't let you build before sense-making.
Research has a hard cap and accountability — 20 searches, counted out loud, with explicit INSUFFICIENT DATA flags. Raw prompts hallucinate gaps.
The roster is derived from evidence — every hire is cited to a specific research panel, not generated from a template.
Sessions persist — INCUBATOR.md means the CEO picks up exactly where you left off, without you re-explaining the project.
Agent quality is gated — every generated skill goes through a cold-start critic pass before you install it.
The skill is useful in a single session. It becomes a different tool over weeks.
Long AI sessions accumulate state drift — the model reconstructs your project from context each session, and that reconstruction drifts slightly each time. Three engineering decisions address this:
Phase checkpointing. INCUBATOR.md is written the moment the Phase 1 thesis is locked — not at the end. Any session dropout between Phase 1 and Phase 4 is recoverable from disk. session-state.json is written after Phases 2, 3, and 4, carrying phase, confidence, open risks, and a 20-word summary. On resume, the CEO reads this first — not the conversation history.
Canonical JSON state. roster.md stores agent team state as a JSON block at the bottom of the file. The CEO reads and writes that JSON block first, then regenerates the markdown table from it. Markdown is presentation only. The two cannot diverge.
Cold-start quality gates. After generating each agent skill, the CEO pauses and waits for a CRITIQUE token from the user. This triggers a new generation turn with a cold-start evaluation instruction — the model must evaluate the skill as if it didn't write it. In-stream self-critique (asking the model to evaluate in the same output pass it just generated) is not real evaluation. The separate turn is.
These are prompt-level mitigations, not runtime enforcement. The substrate limits (no true parallel agents, no runtime schema enforcement, no live KPI access) are documented in the full-thesis writeup .
Built as a founding operating system , not a company simulator. What you get is a structured pre-build research-and-planning ritual with a persistent state layer — the kind of thinking an experienced founding team does before writing a line of code, compressed into a turn-based session.
The model plays multiple personas, all from the same context. That's a planning artifact, not a real team. The value isn't the illusion of an org — it's that the phases enforce order: you can't hire before you've researched, can't build before you've validated, can't add Scale roles before you've earned them. Sense-making before building. Agentic coding makes building feel free; that's the trap this is designed to prevent.
Roster pattern adapted from slavingia/skills .
Idea in, company out. A Claude skill that becomes a domain-expert CEO for your new project: it interviews you, researches the market autonomously, ships a findings report, then hires a sequenced team of specialist AI agents. Installable as a Claude Code plugin.
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
