---
source: "https://github.com/mirko-pira/skillhealth"
hn_url: "https://news.ycombinator.com/item?id=48644959"
title: "Skillhealth: Find Claude Code skills you never use and what they cost"
article_title: "GitHub - mirko-pira/skillhealth: Audit & observability for your agent skills — usage heat, doctor with fixes, relationship graph · GitHub"
author: "mirkopira"
captured_at: "2026-06-23T13:50:30Z"
capture_tool: "hn-digest"
hn_id: 48644959
score: 1
comments: 0
posted_at: "2026-06-23T13:49:26Z"
tags:
  - hacker-news
  - translated
---

# Skillhealth: Find Claude Code skills you never use and what they cost

- HN: [48644959](https://news.ycombinator.com/item?id=48644959)
- Source: [github.com](https://github.com/mirko-pira/skillhealth)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T13:49:26Z

## Translation

タイトル: Skillhealth: 使用したことのないクロード コードのスキルとそのコストを確認する
記事のタイトル: GitHub - mirko-pira/skillhealth: エージェント スキルの監査と観察可能性 — 使用状況の熱、修正のある医師、関係グラフ · GitHub
説明: エージェント スキルの監査と可観測性 - 使用状況の熱、修正を含む医師、関係グラフ - mirko-pira/skillhealth

記事本文:
GitHub - mirko-pira/skillhealth: エージェント スキルの監査と観察可能性 — 使用量の熱、修正を含む医師、関係グラフ · GitHub
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
ミルコピラ
/
スキル健康
公共
通知
通知設定を変更するにはサインインする必要があります

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github .github claude-skill/ skillhealth claude-skill/ skillhealth crates crates docs docs .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md Deny.toml Deny.toml dist-workspace.toml dist-workspace.toml llms.txt llms.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントのスキルについては、flutter Doctor を参照してください。所有しているもの、実際に使用しているもの、壊れているもの、それらすべてがどのように接続されているかを確認してください。
デモ · 機能 · なぜ · ツールなのか · インストール · クイックスタート · 仕組み · ベンチマーク · フラグ · キーマップ · プライバシー · ロードマップ · ライセンス
すべてが 1 つの画面上にあります。 skillhealth は、インストールされている各スキルを使用状況に基づいてランク付けします。
実際のトランスクリプトから計算された熱 (暑い、暖かい、冷たい、または死んだ) の隣にあります。
トークンと 12 週間のスパークラインでのコスト。スキルまたはトランスクリプトを変更する
コックピットが開いている間、リストは自動的に再ランク付けされます。それは死者です
正規表現ビルド スキルが上のカメラで熱くなります。医師の場合は d を押してください: 何でも
壊れているということは、その理由と修正を伴う発見となり、それがあなたのファイルにコピーされます。
クリップボード。
同じデータ、スクリプト可能: パイプ、--json、ゲート CI。
どちらの録音も完全に再生可能です。
docs/demo/demo-tui.tape および
docs/demo/demo.tape (VHS) を合成ファイルよりも
フィクスチャ環境 ( docs/demo/setup-demo.sh )。
skillhealth グラフ --open は同じポートフォリオをマップとして表示します。各スキルは
ノードは使用法によってサイズ設定され、熱によって色付けされます。相互参照はエッジであり、
ドクタードロワーと並べ替え可能なテーブルビュー。ドラッグ、ズーム、/

フィルタリングするには、
ノードの使用傾向と接続を確認します。 100 ノードを超えると最初に警告が表示されます
毛玉をレンダリングする代わりに。
スキルは依存関係のように積み重なっていきます。マーケットプレイスにいくつかインストールすると、
何百ものものを運び、セッションごとにそれぞれがコンテキストウィンドウに表示されます
発火するか否か。 skillhealth を使用すると、コマンド 1 つで全体を確認できます。
パイルと壊れた部分の貼り付け可能な修正。
ステータス: v0.2 — 完全なテスト スイート グリーン、実際の 330 スキルに対して検証済み
インストールと 3.86 GB のトランスクリプト コーパス。まずクロード・コード。コーデックスと
Gemini CLI はロードマップにあります。
実際のトランスクリプトからの使用熱 - ホット/ウォーム/コールド/デッドは
セッショントランスクリプト内の実際のスキル呼び出しから計算されます
( ~/.claude/projects/**/*.jsonl )、ファイルの日付から推測されません。あなたは学びます
実際に使っているもの。
ライブコックピット — ターミナルで SkillHealth を裸で実行します: フルスクリーン TUI
熱色リスト、12 週間のスパークライン、ドクター ビュー、および自動更新
スキルやトランスクリプトの変更。パイプと CI はプレーンな出力を保持します。
スコープ ピッカー — --scope project|all|user がどのスキル ルートを適用するかを選択します
含む;リポジトリに .claude/skills ディレクトリがある場合、プロジェクトとして自動検出されます。
それ以外の場合はすべて。 TUI サイクル スコープ内の p はライブです。
プロジェクト レンズ — --lens プロジェクト|グローバル フィルターの使用熱と医師
現在のプロジェクトのトランスクリプトのみ。 TUI の L で切り替えます。
無効化されたプラグインの認識 — EnabledPlugins によってオフになったスキルは、
専用オフ状態: ディスク上に存在し、ロードされず、から除外されます。
常時オンのトークンの合計。
コスト分割 — すべてのスキルが always_on (セッションごとにロードされる) を示す vs
on_fire (呼び出し時のみロードされる) トークンは別途コストがかかります。概要
フッターには、ポートフォリオ全体の常時表示の合計が表示されます。
入力された履歴のトランスクリプトをクロスチェックします —history.jsonl (Claude Code の
自分のコム

mand log) は 2 番目の使用シグナルです。医師が W010 を発見すると、
スキルは履歴に表示されますが、トランスクリプトのヒットはありません - トランスクリプトのローテーション
または、間違った --projects-dir が、熱を静かに控えめに表現する代わりに表面化しました。
完全な検出 — ユーザー スキル ( ~/.claude/skills )、プロジェクト スキル
( .claude/skills 、ウォークアップ付き: リポジトリの任意のサブディレクトリから実行)、
マーケットプレイス プラグイン、同じ名前の場合のシャドウイング検出機能付き
2箇所に存在します。
実用的な修正を行う医師 — 10 のチェック。すべての発見物は発送されます
なぜ、そしてほとんどの場合、シェルに貼り付けることができる具体的な修正が含まれています。
ドクター自身が何かを編集することは決してありません。
関係グラフ — エッジは実際の相互参照です (1 つのスキルの
別の CLAUDE.md スキルをまとめたボディについて言及しています)、次のようにレンダリングされます
インタラクティブな HTML ダッシュボード、Mermaid、または JSON。
詳細ビュー — skillhealth <name> : 呼び出し回数、最後に使用された、
トークンのコスト分割 (常時オン vs オンファイア)、ソース、接続スキル、
その 1 つのスキルに関する発見。
高速 — 増分キャッシュを使用した並列トランスクリプト スキャン: ~3ms の起動、
上記の 3.86 GB コーパスで測定された、ウォーム ランは約 50 ミリ秒、コールド 1 GB あたり約 0.5 秒です。
スクリプト可能 — 安定したスキーマを使用する場合はどこでも --json、スキーマの場合は --md
Mermaid グラフ、セマンティック終了コードを含むマークダウン レポート
(正常 0 件、警告 1 件、エラー 2 件) なので、そのまま CI に落ちます。
またはコミット前のフック。
ローカルのみ - ネットワークなし、テレメトリなし。トランスクリプトの内容
出力には決して表示されません。呼び出しカウントとタイムスタンプのみがカウントされます。
実際の 330 スキルのインストールによる 3 つの具体的な問題:
スキルは黙って蓄積されます。プラグインは一度に数十個同梱されます。何もない
それらをアンインストールすることはありません。ほとんどの人は「自分にはどのくらいのスキルがあるのか」には答えられません。
持ってる？」 2 分の 1 以内で、ましてや「今月解雇したのは誰ですか?」
すべてのスキルにはコンテキスト、すべてのセッションが必要です。スキル

メタデータが読み込まれています
スキルが発動するかどうかに関係なく、コンテキスト ウィンドウに表示されます。デッドスキルは
純粋な税金: 上記の 330 スキルのインストールでは、一度も発動されなかったスキルのみ
単一セッションごとに最大 18,000 個のコンテキスト トークンを書き込みます。
破損は目立ちません。壊れたシンボリックリンク、無効なフロントマター、または
削除されたスキルを指す CLAUDE.md トリガー — エージェントはそれらをスキップします
静かに。決して分かりません。
フォルダーは重要な質問に答えることができないからです。
ls ~/.claude/skills は名前を表示します。スキルがどのようなものかは認識されません。
3 月以来起動していない、そのフロントマターが 2 つの解析を停止したため
更新前に CLAUDE.md が削除したものをまだ指している、または
セッションごとに、山全体のトークンにかかる費用。
skillhealth は 3 つのソースを相互参照します - ディスク上にあるもの、あなたのもの
記録によると実際に起こったこと、そしてCLAUDE.mdが約束していること - そして
すべてのデルタを貼り付け可能な修正に変換します。
どのスキルが発動し、どのスキルがアイドル状態になるかを知るのは簡単です。役立つ仕事
は 4 つのことを結び付けています。
使用法: ファイルの日付ではなく、実際のトランスクリプト呼び出しからの熱。
コスト: 常時オンとオンファイアのトークンが分割されるため、無効なスキルは次のように表示されます。
セッションごとにかかる費用。 CLAUDE.md の予算は重要です
@import - サーフェスの読み取りで欠落する展開されたコンテンツ。
健康: すべての所見に理由と貼り付け可能な修正がある医師。
それは報告します。ファイルを編集することはありません。
配線: どのスキルがどのスキルを参照するか、CLAUDE.md トリガーがどこを指すか、
影になっているもの、または孤立しているもの。
4 つすべてを 1 つのパスで、ローカル、1 つの静的バイナリ、--json およびセマンティックを使用
終了コードを使用して、同じデータが CI にドロップされるようにします。
npx skillhealth # ゼロインストール (実際のバイナリ、Bun/Node ランタイム トリックなし)
# または
curl --proto ' =https ' --tlsv1.2 -LsSf https://github.com/mirko-pira/skillhealth/releases/latest/download/skillhealth-installer.sh |

しー
# または
貨物設置スキル健康
ソースより:
git clone https://github.com/mirko-pira/skillhealth
cd スキルヘルス
カーゴインストール --path crates/skillhealth
クイックスタート
skillhealth # ライブ コックピット: すべてのスキル、実際の使用による熱
skillhealth 医師 # 診断 — すべての結果: 理由 + 貼り付け可能な修正
skillhealth グラフ --ブラウザで # インタラクティブなダッシュボードを開きます
skillhealth コードレビュー # 1 つのスキル: 使用法、傾向、接続、調査結果
どこでもスクリプト可能:
スキルヘルス --json | jq ' .skills[] | select(.温度 == "デッド") | .name '
skillhealth chart --format mermaid # 任意の Markdown ファイルに貼り付けます
skillhealth 医師 && echo "skill health" # exit 0 / 1 警告 / 2 エラー
仕組み
~/.claude/skills/ .claude/skills/ ~/.claude/plugins/
ユーザー スキル プロジェクト スキル マーケットプレイス プラグイン
│ (CWD からのウォークアップ) │
━━━━━━━━━━━━━━━━━━━━━━┘
▼
┌─────┐ CLAUDE.md (+ @import 解像度)
│ 発見 │ ◄── トリガーリファレンス、トークンバジェット
━━━┬─────┘
▼
┌─────┐ ~/.claude/projects/**/*.jsonl
│ 解析 │ トランスクリプト — 呼び出し数
└────┬──────┘ だけ、決して満足しない
▼ │
┌───────┴──────┐ ▼
▼ ▼ ┌───────┐
┌─────┐ ┌───────┐ │ 使用状況スキャン│──► 増分キャッシュ
│ グラフ │ │ドクター │◄┤ (レーヨン) │ (暖かい ≈ 50ms)
━━━┬─────┘ └────┬───┘ └──────┘
━───────┬───┘
▼
ターミナル・JSON・マーク

ダウン + マーメイド · HTML ダッシュボード
レイヤー
木箱
所有しています
コア
スキルヘルスコア
3 つのルートすべてにわたる検出 (シャドウイング + デブリ検出による)、フロントマター解析 (寛容なフォールバックを備えた厳密な YAML - 現実世界の SKILL.md ファイルは厳密なパーサーを破る)、増分キャッシュによる並列トランスクリプト スキャン、関係グラフ、ドクター チェック、レポート モデル
CLI
スキル健康
clap ベースのバイナリ、ライブ TUI コックピット (ratatui)、レンダラー (ターミナル、詳細、ドクター、JSON、Markdown/Mermaid)、バイナリに埋め込まれた自己完結型 HTML ダッシュボード — 取得するアセットなし、オフラインで動作
ベンチマーク
実際の 330 スキルのインストールと 3.86 GB に対して Apple Silicon 上で測定
トランスクリプト コーパス ( ~/.claude/projects/**/*.jsonl ):
コールド スキャンは最悪のケースです。最初の実行、またはキャッシュ ディレクトリの実行直後です。
クリアされました。定常状態は増分キャッシュに乗ります: 各実行は再読み取りのみです
最後のスキャン以降に mtime が移動したトランスクリプトは、
コーパスがどれだけ大きくなるかに関係なく、範囲は数十ミリ秒です。
hyperfine --warmup 3 ' skillhealth --json ' # ウォーム ラン
skillhealth --cache-dir " $( mktemp -d ) " --json # コールド スキャンを強制する
フラグ
旗
説明
--scope プロジェクト | すべて | ユーザー
どのスキルルートを含めるか。リポジトリに .claude/skills ディレクトリがある場合、プロジェクトとして自動検出されます。それ以外の場合はすべて。
--レンズ プロジェクト|グローバル
使用状況の熱と医師の調査結果を現在のプロジェクトのトランスクリプト ( project ) または完全なコーパス ( global 、デフォルト) にフィルターします。
--config-dir <パス>
スキル、プラグイン、CLAUDE.md、およびhistory.jsonlのルート。デフォルトは ~/.claude です。
--projects-dir <パス>
トランスクリプト スキャンのルート。 --config-dir から切り離されました。デフォルトは ~/.claude/projects です。
--cache-dir <パス>
インクリメンタル スキャン キャッシュが存在する場所。デフォルトは OS キャッシュ ディレクトリです。
TUIキーマップ
キー
アクション
↑ ↓

[切り捨てられた]

## Original Extract

Audit & observability for your agent skills — usage heat, doctor with fixes, relationship graph - mirko-pira/skillhealth

GitHub - mirko-pira/skillhealth: Audit & observability for your agent skills — usage heat, doctor with fixes, relationship graph · GitHub
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
mirko-pira
/
skillhealth
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github .github claude-skill/ skillhealth claude-skill/ skillhealth crates crates docs docs .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md deny.toml deny.toml dist-workspace.toml dist-workspace.toml llms.txt llms.txt View all files Repository files navigation
flutter doctor for your agent skills — what you have, what you actually use, what's broken, and how it all connects.
Demo · What it does · Why · Why a tool · Install · Quickstart · How it works · Benchmarks · Flags · Keymap · Privacy · Roadmap · License
Everything on one screen. skillhealth ranks each installed skill by usage
heat (hot, warm, cold, or dead), computed from your real transcripts, next to
what it costs in tokens and a 12-week sparkline. Change a skill or transcript
while the cockpit is open and the list re-ranks itself; that's the dead
regex-build skill going hot on camera above. Press d for doctor: anything
broken becomes a finding with a why and a fix that y copies to your
clipboard.
Same data, scriptable: pipe it, --json it, gate CI on it.
Both recordings are fully reproducible:
docs/demo/demo-tui.tape and
docs/demo/demo.tape (VHS) over a synthetic
fixture environment ( docs/demo/setup-demo.sh ).
skillhealth graph --open shows the same portfolio as a map: each skill is a
node sized by usage and colored by heat, cross-references are edges, with a
doctor drawer and a sortable table view. Drag, zoom, / to filter, click a
node for its usage trend and connections. Past 100 nodes it warns you first
instead of rendering a hairball.
Skills pile up like dependencies. A few marketplace installs in, you're
carrying hundreds, and each one sits in your context window every session
whether it fires or not. skillhealth gives you one command to see the whole
pile and a pasteable fix for whatever's broken.
Status: v0.2 — full test suite green, validated against a real 330-skill
install and a 3.86 GB transcript corpus. Claude Code first; Codex and
Gemini CLI are on the roadmap.
Usage heat from your real transcripts — hot / warm / cold / dead is
computed from actual Skill invocations in your session transcripts
( ~/.claude/projects/**/*.jsonl ), not guessed from file dates. You learn
what you actually use.
Live cockpit — run skillhealth bare in a terminal: full-screen TUI with
heat-colored list, 12-week sparklines, doctor view, and auto-refresh when a
skill or transcript changes. Pipes and CI keep the plain output.
Scope picker — --scope project|all|user selects which skill roots to
include; auto-detected as project when the repo has a .claude/skills dir,
otherwise all . p in the TUI cycles scopes live.
Project lens — --lens project|global filters usage heat and the doctor
to the current project's transcripts only. L in the TUI toggles it.
Disabled-plugin awareness — skills turned off via enabledPlugins get a
dedicated off state: present on disk, never loaded, excluded from the
always-on token total.
Cost split — every skill shows always_on (loaded every session) vs
on_fire (loaded only when invoked) token costs separately. The overview
footer shows the always-on total across your portfolio.
Typed history cross-checks transcripts — history.jsonl (Claude Code's
own command log) is a second usage signal. Doctor finding W010 fires when a
skill appears in history but has zero transcript hits — transcript rotation
or a wrong --projects-dir , surfaced instead of silently understating heat.
Full discovery — user skills ( ~/.claude/skills ), project skills
( .claude/skills , with walk-up: run it from any subdirectory of a repo),
and marketplace plugins, with shadowing detection when the same name
exists in two places.
Doctor with actionable fixes — ten checks; every finding ships a
why , and most carry a concrete fix you can paste into your shell.
Doctor never edits anything itself.
Relationship graph — edges are real cross-references (one skill's
body mentioning another, CLAUDE.md wiring skills together), rendered as
an interactive HTML dashboard, Mermaid, or JSON.
Detail view — skillhealth <name> : invocation count, last used, the
token cost split (always-on vs on-fire), source, connected skills, and
findings for that one skill.
Fast — parallel transcript scan with an incremental cache: ~3ms startup,
~50ms warm runs, ~0.5s per GB cold, measured on the 3.86 GB corpus above.
Scriptable — --json everywhere with a stable schema, --md for a
Markdown report with a Mermaid graph, semantic exit codes
( 0 healthy · 1 warnings · 2 errors) so it drops straight into CI
or a pre-commit hook.
Local only — zero network, zero telemetry. Your transcript content
never appears in any output — only invocation counts and timestamps.
Three concrete pains, from a real 330-skill install:
Skills accumulate silently. Plugins ship dozens at a time; nothing
ever uninstalls them. Most people cannot answer "how many skills do I
have?" within a factor of two — let alone "which ones fired this month?"
Every skill costs context, every session. Skill metadata is loaded
into the context window whether the skill fires or not. Dead skills are
pure tax: on the 330-skill install above, the never-fired skills alone
burn ~18K tokens of context in every single session.
Breakage is invisible. A broken symlink, invalid frontmatter, or a
CLAUDE.md trigger pointing at a deleted skill — the agent skips them
silently. You find out never.
Because the folder can't answer the questions that matter.
ls ~/.claude/skills shows you names — it doesn't know that a skill
hasn't fired since March, that its frontmatter stopped parsing two
updates ago, that CLAUDE.md still points at something you deleted, or
what the whole pile costs you in tokens, every single session.
skillhealth cross-references three sources — what's on disk, what your
transcripts say actually happened, and what CLAUDE.md promises — and
turns every delta into a fix you can paste.
Knowing which skills fire and which sit idle is the easy part. The useful job
is tying four things together:
Usage : heat from real transcript invocations, not file dates.
Cost : the always-on vs on-fire token split, so a dead skill shows up as
what it costs you every session. The CLAUDE.md budget counts
@import -expanded content, which a surface read misses.
Health : a doctor where every finding has a why and a pasteable fix .
It reports; it never edits your files.
Wiring : which skills reference which, where CLAUDE.md triggers point,
what's shadowed or orphaned.
All four in one pass, local, one static binary, with --json and semantic
exit codes so the same data drops into CI.
npx skillhealth # zero install (real binaries, no Bun/Node runtime tricks)
# or
curl --proto ' =https ' --tlsv1.2 -LsSf https://github.com/mirko-pira/skillhealth/releases/latest/download/skillhealth-installer.sh | sh
# or
cargo install skillhealth
From source:
git clone https://github.com/mirko-pira/skillhealth
cd skillhealth
cargo install --path crates/skillhealth
Quickstart
skillhealth # live cockpit: every skill, heat from real usage
skillhealth doctor # diagnostics — every finding: why + pasteable fix
skillhealth graph --open # interactive dashboard in the browser
skillhealth code-review # one skill: usage, trend, connections, findings
Scriptable everywhere:
skillhealth --json | jq ' .skills[] | select(.temperature == "dead") | .name '
skillhealth graph --format mermaid # paste into any Markdown file
skillhealth doctor && echo " skills healthy " # exit 0 / 1 warnings / 2 errors
How it works
~/.claude/skills/ .claude/skills/ ~/.claude/plugins/
user skills project skills marketplace plugins
│ (walk-up from cwd) │
└──────────────────────┬────────────────────────┘
▼
┌──────────┐ CLAUDE.md (+ @import resolution)
│ discover │ ◄── trigger references, token budget
└────┬─────┘
▼
┌──────────┐ ~/.claude/projects/**/*.jsonl
│ parse │ transcripts — invocation counts
└────┬─────┘ only, never content
▼ │
┌───────────┴──────┐ ▼
▼ ▼ ┌───────────┐
┌──────────┐ ┌───────┐ │ usage scan│──► incremental cache
│ graph │ │doctor │◄┤ (rayon) │ (warm ≈ 50ms)
└────┬─────┘ └───┬───┘ └───────────┘
└───────┬────────┘
▼
terminal · JSON · Markdown + Mermaid · HTML dashboard
Layer
Crate
Owns
Core
skillhealth-core
discovery across all three roots (with shadowing + debris detection), frontmatter parsing (strict YAML with a lenient fallback — real-world SKILL.md files break strict parsers), parallel transcript scan with incremental cache, relationship graph, doctor checks, report model
CLI
skillhealth
clap-based binary, live TUI cockpit (ratatui), renderers (terminal, detail, doctor, JSON, Markdown/Mermaid), self-contained HTML dashboard embedded in the binary — no assets to fetch, works offline
Benchmarks
Measured on Apple Silicon against a real 330-skill install and a 3.86 GB
transcript corpus ( ~/.claude/projects/**/*.jsonl ):
Cold scan is the worst case — a first run, or right after the cache dir is
cleared. Steady state rides the incremental cache: each run only re-reads
transcripts whose mtime moved since the last scan, so it stays in the
tens-of-milliseconds range regardless of how large the corpus grows.
hyperfine --warmup 3 ' skillhealth --json ' # warm runs
skillhealth --cache-dir " $( mktemp -d ) " --json # force a cold scan
Flags
Flag
Description
--scope project|all|user
Which skill roots to include. Auto-detected as project when the repo has a .claude/skills dir; otherwise all .
--lens project|global
Filter usage heat and doctor findings to the current project's transcripts ( project ) or the full corpus ( global , default).
--config-dir <path>
Root for skills, plugins, CLAUDE.md, and history.jsonl . Defaults to ~/.claude .
--projects-dir <path>
Root for transcript scanning. Decoupled from --config-dir ; defaults to ~/.claude/projects .
--cache-dir <path>
Where the incremental scan cache lives. Defaults to the OS cache dir.
TUI keymap
Key
Action
↑ ↓

[truncated]
