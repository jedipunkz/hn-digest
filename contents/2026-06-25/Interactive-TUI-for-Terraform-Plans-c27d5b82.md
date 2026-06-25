---
source: "https://github.com/omarismail/terraform-plan-tui"
hn_url: "https://news.ycombinator.com/item?id=48673132"
title: "Interactive TUI for Terraform Plans"
article_title: "GitHub - omarismail/terraform-plan-tui: A Terminal UI for Terraform Plans · GitHub"
author: "omar-terraform"
captured_at: "2026-06-25T13:42:10Z"
capture_tool: "hn-digest"
hn_id: 48673132
score: 1
comments: 0
posted_at: "2026-06-25T13:33:28Z"
tags:
  - hacker-news
  - translated
---

# Interactive TUI for Terraform Plans

- HN: [48673132](https://news.ycombinator.com/item?id=48673132)
- Source: [github.com](https://github.com/omarismail/terraform-plan-tui)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T13:33:28Z

## Translation

タイトル: Terraform プランの対話型 TUI
記事のタイトル: GitHub - omarismail/terraform-plan-tui: Terraform プラン用のターミナル UI · GitHub
説明: Terraform プランのターミナル UI。 GitHub でアカウントを作成して、omarismail/terraform-plan-tui の開発に貢献してください。

記事本文:
GitHub - omarismail/terraform-plan-tui: Terraform プラン用のターミナル UI · GitHub
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
オマリスメール
/
terraform-plan-tui
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店T

ags ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット cmd/ tfplantui cmd/ tfplantui docs docs 例 例 内部 内部 .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum tfplantui.go tfplantui.go すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Terraform プランを読み取るための対話型ターミナル UI。
Terraform 構成はグラフであり、計画は提案された変更です。
そのグラフ。 Terraform プランの出力の数百行をスクロールする
その構造はほとんど見えなくなります。 tfplantui は計画を次のようにレンダリングします
依存関係グラフは依存関係の深さごとに列にレイアウトされており、
矢印キーを使用して移動し、リソースを展開すると、色分けされたリソースが表示されます。
属性の前後の差分。
レイヤ0 レイヤ1 レイヤ2
╭─────────────────╮ ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓ ╭─────────────────╮
│ ・random_id.seed │ │ ± randan_string.api_key ×4 │ │ ± random_password.master │
╰─────────────────╯ ┃ キーパー： ┃ ╰───────────────╯
╭─────────────────╮ ┃ ~ 回転 = "v1" → "v2" ┃ ╭──────────────────╮
│ ·random_pet.org │ ┃ ~ result = "abc…" → (既知…) ┃ │ ±random_pet.release │
╰───────

──────────╯ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ╰───────────────╯
なぜ
terraform plan は、その答えをプレーン テキストとして提供します。
変更により、数百行の差分をスクロールする必要があるか、生の計画が変更される
自分自身を解析するための JSON。いずれにせよ、計画を再構築するのはあなたです
頭の中で形を作ります。 tfplantui はそれをあなたのために行います。それは計画を
グラフをナビゲートできるため、人間はすべてのグラフを視覚化し、推論することができます。
トランスクリプトを読む代わりに、すぐに変更します。
依存関係のレイアウト。依存関係のないリソースは左側にあります。それぞれ
右側の列はその前の列に依存します。 ← / → に移動すると、
エッジなので、何が何に依存しているかを追跡できます。
その場で展開します。任意のリソース上でスペースを押すと、その属性の差分が展開されます。
予想どおりに色分けされています: 緑の追加、赤の削除、オレンジ
インプレース変更、および (適用後に判明) 計算値の場合。
小規模および大規模なプラン向けに構築されています。 6 つのリソース計画と
400 インスタンスの 5 層プランは同じビューを使用します。列は垂直方向にスクロールし、
レイアウトは水平にスクロールし、フラットなリスト ビュー ( t ) と
変更のみのフィルター ( n ) は、大きな計画を扱いやすくします。
github.com/omarismail/tfplantui/cmd/tfplantui@latest をインストールしてください
# またはチェックアウトからビルドする
go build -o tfplantui ./cmd/tfplantui
Go 1.24 以降が必要です。 Terraform は、バイナリ プラン ファイルを渡す場合にのみ必要です
(以下を参照);読書計画 JSON には他に何も必要ありません。
tfplantui もパッケージであるため、どの Go CLI にも同じプラン エクスプローラーを埋め込むことができます
独自の旗の後ろに。プラン ファイルと残りを渡します。バイナリ プランをロードします。
terraform show -json 、解析、レイアウト、

インタラクティブ UI — のために処理されます
あなた：
「github.com/omarismail/tfplantui」をインポートします
// 独自のコマンド/フラグ ハンドラー内:
func runPlanExplorer ( planPath string ) エラー {
// planPath は、バイナリ プラン ファイル、プラン JSON、または標準入力の「-」です。
tfplantui を返します。実行 ( planPath )
}
公開面は意図的に小さくなっています。
// インタラクティブ TUI を起動し、ユーザーが終了するまでブロックします。
func Run (パス文字列、opts ... オプション) エラー
// 同じですが、プランの JSON はすでにメモリに保持されています。
func RunJSON ( planJSON [] byte , opts ... Option ) エラー
// 代わりに非インタラクティブなテキストの概要を書き込みます (CI に便利)。
func Summary ( w io. Writer , path string , opts ... Option ) エラー
// オプション: WithAltScreen、WithTerraformPath、WithPlanName、WithProgramOptions
使用法
tfplantui は、プランの JSON 形式を読み取ります ( terraform show -json )。できます
次の 3 つのいずれかに向けてください。
# 1. バイナリ プラン ファイル — tfplantui が `terraform show -json` を実行します
terraform plan -out=plan.tfplan
tfplantui plan.tfplan
#2. 自分で作成したJSONを計画する
terraform show -json plan.tfplan > plan.json
tfplantui plan.json
#3.標準入力にパイプ接続
terraform show -json plan.tfplan | tfplantui -
# 非インタラクティブなテキストの概要 (CI / クイックトリアージに最適)
tfplantui -summary plan.json
チェックアウトからすぐにお試しください。インストールは必要ありません。
./cmd/tfplantui Examples/medium/plan.json を実行します。
キー
キー
アクション
↑ ↓ / k j
依存関係レイヤー (列) 内で移動する
← → / h l
依存関係のエッジに従ってレイヤー間を移動します
スペース/入力
選択したリソースの差分を展開/折りたたむ
[ ]
count / for_each リソースのサイクルインスタンス
あなた
変更されていない属性を表示/非表示にする
n
変更のみ — 操作なしのリソースを非表示にする
×
すべてを崩壊させる
t
グラフビューとリストビューを切り替える
グラム/グラム
列またはリストの先頭/末尾にジャンプします
/
検索

リソースアドレス別 ( Enter でジャンプ、Esc でキャンセル)
?
ヘルプを切り替える
q / Ctrl+C
やめる
色の凡例
色には 1 つの要素、つまりアクションが含まれるため、一目で判読できます。
他のものはすべて意図的にニュートラルなので、アクションの色と競合しません。
カーソルは、明るい白い文字が付いた暗いカードです。色相ではなく、対照的に読み取られます。
変更されていない (操作なし) リソースは、ほぼ背景まで淡色表示されます。彼らだけが
カーソルをそれらの上に移動すると明るくなります。
依存関係はオンデマンドで描画されます。リソースを選択すると、そのリソースから矢印が描画されます。
色を変更するのではなく、依存関係をその依存関係に入れたり、依存関係から外したりする ( ─▶ )
他のノード。フッターには名前ごとにリストも表示されます。
3 つのプランは、examples/ の下でチェックインされます。すべて、
hachicorp/random プロバイダーなので、完全にローカルで再現可能です。
各 plan.json はコミットされるため、サンプルは Terraform なしで実行されます。に
それらを再生成します (これには Terraform が必要です)。
Examples/generate.sh all # または: simple |中 |大きい
各サンプル ディレクトリには以下が含まれます。
Providers.tf — required_providers ブロック (常にロードされる)、
main.tf — 後の世界 (計画の終了点。これは読み取り可能な構成です)、
before.tf.txt — オプションの before ワールド。最初にシードに適用されます。
プランに更新/置換/削除が含まれるように状態を設定します (これは .tf ファイルではないため、
Terraform は、generate.sh がそれを交換するまでそれを無視します)、
plan.json — 生成され、コミットされたアーティファクト。
大きなサンプルの構成は、examples/large/gen_config.py によって出力されます。
Parse (internal/tfplan): resource_changes と構成を読み取ります。
プラン JSON からのブロック。インスタンス ( count / for_each ) はそれぞれのグループにグループ化されます。
リソース ブロック 。これが 1 つのグラフ ノードになります。
グラフ: 依存関係エッジは、それぞれの Terraform レコードの参照から取得されます。
リソースのconfiguration.expressions。ノードはアッシです

柱にGned
最長の依存関係パス (杉山式階層化) なので、深さが読み取られます。
左→右。
Diff (internal/tfplan/diff.go): 前後を一緒に歩きます。
スレッド after_unknown ((適用後既知) ) および
before_sensitive / after_sensitive ツリー (機密値は編集されません。
print)、 keeper などのネストされたマップに再帰します。
UI (internal/ui): タピオカティー
モデル。レンダリングはウィンドウ化され、表示される列と行のみが描画されます。
大規模なプランでも応答性を維持します。
モジュール入力変数を通過するモジュール間の依存関係エッジ
(例: ルート リソースから接続された module.x.foo) は描画されません。モジュール内
エッジと直接リソース参照はそうです。リソースは常に表示されます。それだけ
エッジの 1 つのクラスが欠落している可能性があります。
バンドルされたサンプルでは、ランダム プロバイダーが使用されており、その属性はほぼすべて
強制的に新規作成するため、多数のインプレースではなく、作成/置換/削除を示します。
～の更新情報。インプレース更新は正しくレンダリングされます (テストを参照) - ツールをポイントします
実際に彼らに会う計画を立てています。
go test ./... # ユニット + ヘッドレス TUI テスト
TFPLAN_DUMP=1 go test ./internal/ui -run TestDumpVisual -v # レイアウトに注目
ライセンス
Terraform プランのターミナル UI
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

A Terminal UI for Terraform Plans. Contribute to omarismail/terraform-plan-tui development by creating an account on GitHub.

GitHub - omarismail/terraform-plan-tui: A Terminal UI for Terraform Plans · GitHub
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
omarismail
/
terraform-plan-tui
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits cmd/ tfplantui cmd/ tfplantui docs docs examples examples internal internal .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum tfplantui.go tfplantui.go View all files Repository files navigation
An interactive terminal UI for reading Terraform plans.
A Terraform configuration is a graph , and a plan is a proposed change over
that graph . Scrolling through hundrends of lines of terraform plan output
makes that structure almost impossible to see. tfplantui renders the plan as
its dependency graph — laid out in columns by dependency depth — and lets you
walk it with the arrow keys, expanding any resource to see a color-coded
before → after diff of its attributes.
Layer 0 Layer 1 Layer 2
╭──────────────────────────────╮ ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓ ╭──────────────────────────────╮
│ · random_id.seed │ ┃ ± random_string.api_key ×4 ┃ │ ± random_password.master │
╰──────────────────────────────╯ ┃ keepers: ┃ ╰──────────────────────────────╯
╭──────────────────────────────╮ ┃ ~ rotation = "v1" → "v2" ┃ ╭──────────────────────────────╮
│ · random_pet.org │ ┃ ~ result = "abc…" → (known…) ┃ │ ± random_pet.release │
╰──────────────────────────────╯ ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛ ╰──────────────────────────────╯
Why
terraform plan gives you its answer as plain text — and for any real
change that means hundreds of lines of diff to scroll through, or the raw plan
JSON to parse yourself. Either way, you are the one reconstructing the plan's
shape in your head. tfplantui does that for you: it turns the plan into a
graph you can navigate , so a human can visualize and reason through every
change at once instead of reading a transcript.
Dependency layout. Resources with no dependencies sit on the left; each
column to the right depends on the one before it. Moving ← / → follows the
edges, so you can trace what depends on what .
Expand in place. Hit space on any resource to unfold its attribute diff,
color-coded the way you'd expect: green additions, red removals, amber
in-place changes, and (known after apply) for computed values.
Built for small and large plans. A six-resource plan and a
400-instance, five-layer plan use the same view; columns scroll vertically,
the layout scrolls horizontally, and a flat list view ( t ) plus a
changes-only filter ( n ) keep big plans tractable.
go install github.com/omarismail/tfplantui/cmd/tfplantui@latest
# or build from a checkout
go build -o tfplantui ./cmd/tfplantui
Requires Go 1.24+. Terraform is only needed if you hand it a binary plan file
(see below); reading plan JSON needs nothing else.
tfplantui is also a package, so any Go CLI can embed the same plan explorer
behind its own flag. Hand it a plan file and the rest — loading binary plans via
terraform show -json , parsing, layout, and the interactive UI — is handled for
you:
import "github.com/omarismail/tfplantui"
// In your own command / flag handler:
func runPlanExplorer ( planPath string ) error {
// planPath may be a binary plan file, plan JSON, or "-" for stdin.
return tfplantui . Run ( planPath )
}
The public surface is intentionally small:
// Launch the interactive TUI, blocking until the user quits.
func Run ( path string , opts ... Option ) error
// Same, but on plan JSON you already hold in memory.
func RunJSON ( planJSON [] byte , opts ... Option ) error
// Write the non-interactive text summary instead (handy for CI).
func Summary ( w io. Writer , path string , opts ... Option ) error
// Options: WithAltScreen, WithTerraformPath, WithPlanName, WithProgramOptions
Usage
tfplantui reads the JSON form of a plan ( terraform show -json ). You can
point it at any of three things:
# 1. a binary plan file — tfplantui runs `terraform show -json` for you
terraform plan -out=plan.tfplan
tfplantui plan.tfplan
# 2. plan JSON you produced yourself
terraform show -json plan.tfplan > plan.json
tfplantui plan.json
# 3. piped on stdin
terraform show -json plan.tfplan | tfplantui -
# non-interactive text summary (great for CI / quick triage)
tfplantui -summary plan.json
Try it immediately from a checkout — no install needed:
go run ./cmd/tfplantui examples/medium/plan.json
Keys
Key
Action
↑ ↓ / k j
move within a dependency layer (column)
← → / h l
move across layers, following dependency edges
space / enter
expand / collapse the selected resource's diff
[ ]
cycle instances of a count / for_each resource
u
show / hide unchanged attributes
n
changes-only — hide no-op resources
x
collapse everything
t
switch between graph and list views
g / G
jump to top / bottom of the column or list
/
search by resource address ( Enter jumps, Esc cancels)
?
toggle help
q / Ctrl-C
quit
Color legend
Color carries one thing — the action — so it stays legible at a glance:
Everything else is deliberately neutral so it doesn't compete with the action colors:
The cursor is a dark card with bright white text — it reads by contrast, not hue.
Unchanged (no-op) resources are dimmed almost to the background; they only
brighten when you move the cursor onto them.
Dependencies are drawn on demand: selecting a resource draws arrows from its
dependencies into it and out to its dependents ( ─▶ ), instead of recoloring
other nodes. The footer also lists them by name.
Three plans are checked in under examples/ , all using only the
hashicorp/random provider so they are fully local and reproducible:
Each plan.json is committed, so the examples run without Terraform. To
regenerate them (this does require Terraform):
examples/generate.sh all # or: simple | medium | large
Each example directory holds:
providers.tf — the required_providers block (always loaded),
main.tf — the after world (what the plan ends at; this is the readable config),
before.tf.txt — the optional before world, applied first to seed prior
state so the plan contains updates/replaces/deletes (it is not a .tf file, so
Terraform ignores it until generate.sh swaps it in),
plan.json — the generated, committed artifact.
The large example's configs are emitted by examples/large/gen_config.py .
Parse ( internal/tfplan ): reads resource_changes and the configuration
block from the plan JSON. Instances ( count / for_each ) are grouped into their
resource block , which becomes one graph node.
Graph : dependency edges come from the references Terraform records in each
resource's configuration.expressions . Nodes are assigned to columns by
longest dependency path (a Sugiyama-style layering), so depth reads
left → right.
Diff ( internal/tfplan/diff.go ): walks before / after together,
threading after_unknown (for (known after apply) ) and the
before_sensitive / after_sensitive trees (sensitive values are redacted, never
printed), recursing into nested maps such as keepers .
UI ( internal/ui ): a Bubble Tea
model. Rendering is windowed — only the visible columns and rows are drawn — so
large plans stay responsive.
Cross-module dependency edges that pass through a module input variable
(e.g. module.x.foo wired from a root resource) are not drawn; intra-module
edges and direct resource references are. Resources are always shown; only that
one class of edge may be missing.
The bundled examples use the random provider, whose attributes are almost all
force-new, so they demonstrate create/replace/delete rather than many in-place
~ updates. In-place updates render correctly (see the tests) — point the tool
at a real plan to see them.
go test ./... # unit + headless-TUI tests
TFPLAN_DUMP=1 go test ./internal/ui -run TestDumpVisual -v # eyeball the layout
License
A Terminal UI for Terraform Plans
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
