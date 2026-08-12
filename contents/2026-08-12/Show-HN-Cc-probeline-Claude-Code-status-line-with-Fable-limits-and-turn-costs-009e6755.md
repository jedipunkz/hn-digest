---
source: "https://github.com/labzink/cc-probeline"
hn_url: "https://news.ycombinator.com/item?id=49272636"
title: "Show HN: Cc-probeline – Claude Code status line with Fable limits and turn costs"
article_title: "GitHub - labzink/cc-probeline: See where it leaks, stop paying for it — a live Claude Code status line that prices every turn, your subagents, cache rebuilds, plus limits, context and git. · GitHub"
author: "labzink"
captured_at: "2026-08-12T14:14:08Z"
capture_tool: "hn-digest"
hn_id: 49272636
score: 1
comments: 1
posted_at: "2026-08-12T14:07:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cc-probeline – Claude Code status line with Fable limits and turn costs

- HN: [49272636](https://news.ycombinator.com/item?id=49272636)
- Source: [github.com](https://github.com/labzink/cc-probeline)
- Score: 1
- Comments: 1
- Posted: 2026-08-12T14:07:50Z

## Translation

タイトル: HN を表示: Cc-probeline – 寓話の制限とターンコストを含むクロード コードのステータス ライン
記事のタイトル: GitHub - labzink/cc-probeline: リーク箇所を確認し、料金の支払いを停止します — 毎ターン、サブエージェント、キャッシュの再構築、制限、コンテキスト、および Git に価格を設定するライブ クロード コード ステータス ライン。 · GitHub
説明: リーク箇所を確認し、支払いを停止します。毎ターン、サブエージェント、キャッシュの再構築、制限、コンテキスト、git の価格を示すライブ クロード コード ステータス ラインです。 - labzink/cc-プローブライン

記事本文:
GitHub - labzink/cc-probeline: リーク箇所を確認し、料金の支払いを停止します。毎ターン、サブエージェント、キャッシュの再構築、制限、コンテキスト、Git の価格を示すライブ クロード コード ステータス ラインです。 · GitHub
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
ラブジンク
/
cc-プローブリン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
257 コミット 257 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows Assets ass

ets cmd/ cc-probeline cmd/ cc-probeline コマンド コマンド フック フック 内部 内部スクリプト スクリプト テスト テスト .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コードのステータス ラインは、ディスク上に既に存在するセッション ログを読み取る毎ターン (あなた、サブエージェント、およびすべてのキャッシュの再構築) に価格を設定します。アカウント、API キー、テレメトリはありません。
セッションの合計はクロード コード自体から取得されます。ここではターンごとの内訳が公開価格表に対するトークン数から計算され、全体を通して推定値としてラベル付けされます。
実行中にネットワーク呼び出しは行われません。オプションのオプトアウト バックグラウンド アクションが 2 つあります。1 つは GitHub 上の 1 つのパブリック ファイル (現在の価格と最新バージョン) の 1 日 1 回のチェック、もう 1 つは Claude Code 自身の使用状況キャッシュの更新です。その /usage 画面はヘッドレスで実行され、セッションの開始時に 1 回、その後はほとんど実行されません。また、プランに実際にモデル スコープの週次制限があるか、最新の状態を維持するために有料の追加使用量がある場合にのみ実行されます。どちらもセッションに関するものは何も送信しません。プライバシーポリシー →
目に見えない非効率に対してお金を払うのではなく、意図的に制限を使いましょう。
brew install labzink/homebrew-tap/cc-probeline # macOS
カール -fsSL https://raw.githubusercontent.com/labzink/cc-probeline/main/scripts/install.sh | sh # macOS / Linux
Windows (Scoop) および Claude Code プラグイン マーケットプレイス: すべてのインストール オプション → 。すべてのリリースには、SHA256 チェックサムと署名された SLSA ビルド来歴が付属しています。
ほとんどのステータス行は、トークン、ターン、実行中のエージェントなどをカウントします。調査はそれらに価格を設定します。以下のすべてはセッションのローカル ログから取得されます。データは Claude Code が持っていますが、決して表示されません。
毎ターン、価格が設定されています - 不透明なものは一つもありません

セッション合計: 各ステップが独自のコストで実行されるライブ テーブル。
サブエージェントが費やした金額 — サブエージェントの作業は、実行中は表示されません。プローブは、各エージェントを自分のターンの次にライブで請求書に載せます。
キャッシュの再構築 (ドル単位) - TTL (オーケストレーターの場合は 60 分、サブエージェントの場合は 5 分) を超えてアイドル状態になると、次のターンでキャッシュ全体が静かに書き換えられます。プローブはライブでエージングし (⏱ 60 m → 0 m)、ヒット時に再構築の価格を設定します。
追加使用量はパーセントではなく金額で表示されます。プランが使い果たされると、上限 ($20.40 / $120.00) に対して今月実際に支払った金額が線に表示されます。これは、見積もりではなく Anthropic 独自の数値から直接表示されます。
正確な価格を維持します。あなたのドルは、その背後にある価格表と同じくらい正確です。 cc-probeline は、ネットワーク経由で料金を更新します (オプションのオプトアウト チェックは 1 日に 1 回、レンダリング中には行われません)。そのため、Anthropic の価格が変更されると、合計は 1 日以内に反映され、再インストールする必要はありません。オフラインまたはオプトアウトすると、ビルドに焼き付けられたテーブルに戻ります。
リセット クロック付きの 5 時間 / 7 日の制限 (プランに含まれる場合は、モデル スコープの Fable キャップも追加) — 使用量を開くために立ち止まることなく、クロックがいっぱいになるのを見て、いつ解放されるかを正確に把握します。
色分けされたゾーン — 数字が警告領域や重要領域に入ると色が変わるため、必要なときにラインが目に留まります。
さらに、モデル、コンテキスト、git、セッション時間などの重要な要素が加わります。
オーケストレーターもサブエージェントも同様に、すべてのターンが独自のラインに到達し、発生したときに価格が設定されます。最後に、あなたの推論のすべてが実際にどこに使われるのかがわかります。
端末に合わせて設計されています。セグメント、色、幅が気に入らない場合は、 /cc-probeline-config ウィザードの指示に従って設定を書き込みます。TOML を手動で編集する必要はありません。 (これはダッシュボードの下部にあるヒントです

上記。）
100% を超えた瞬間にそれがわかります。追加料金はお客様の管理下にあります。
壁にぶつかった後ではなく、行動する時間がまだある間に警告が表示されます。
キャッシュの再構築はサイレントに行われなくなり、発生した瞬間に価格がわかります。
すべてが読み取られ、どれも触れられませんでした。これがプローブと呼ばれる理由です。
プローブは観察の手段であり、介入ではありません。 cc-probeline が行うことはすべて読み取りと表示だけであり、ユーザーのアカウントにアクセスしたり、ユーザーに関するレポートを作成したりすることはありません。
内容: セッションの JSONL ログ ( ~/.claude/projects/… ) とステータスライン ペイロード クロード コードが直接パイプします。
触れないもの: 認証情報、キーチェーン、OAuth トークン — テレメトリは一切ありません。レンダリングは完全にオフラインで行われます。独自の唯一のネットワーク呼び出しは、1 日 1 回のオプションのオプトアウト価格/バージョン チェックです。公開ファイルの単純なダウンロードであり、セッションについては何も送信されません。また、Claude Code にバックグラウンドで使用状況キャッシュを更新するように要求することもできます ( /usage 、ヘッドレス、マシン全体のスロットル、オプトアウト)。これにより、モデル スコープの週ごとの制限が最新の状態に保たれます。その呼び出しは、Claude Code によるものであり、資格情報が使用され、モデル トークンはかかりません。両方をオフにすると、cc-probeline は何も送信も開始も行いません。
バイナリ: 単一のコンパイル済み Go バイナリ、ランタイム依存関係なし、1 回の実行は ≈ 5 ミリ秒です。
監査可能: MIT ライセンス、オープン ソース、SHA256 チェックサムと署名付きビルド来歴 (SLSA) で公開されたすべてのリリース — すべてのダウンロードを gh attestation verify <file> --repo labzink/cc-probeline で検証します。
以下のすべてのチャネルは同じことを行います。バイナリをインストールし、それを接続します。
クロードコードのステータス行。インストール後、Claude Code を再起動すると、
完了 — 追加のコマンドはありません。 (すでにカスタム ステータス行がある場合は、そのまま残ります。
手つかずの; cc-probeline install --merge- で cc-probeline に切り替えます。

設定 --force 。)
Homebrew (macOS — これはキャスクです。Linux では以下のcurlを使用します):
brew インストール labzink/homebrew-tap/cc-probeline
curl (macOS / Linux — OS のリリース アーカイブをダウンロードし、SHA256 を検証し、バイナリをインストールします):
カール -fsSL https://raw.githubusercontent.com/labzink/cc-probeline/main/scripts/install.sh |しー
スクープ (Windows、実験的):
スクープバケット labzink を追加 https://github.com/labzink/scoop - バケット
スクープインストールCC - プローブライン
クロード コード プラグイン マーケットプレイス:
/プラグイン マーケットプレイスに labzink/cc-probeline を追加
次に、 /plugin メニュー (または /plugin install cc-probeline ) からプラグインをインストールし、Claude Code を再起動します。以下のスラッシュ コマンドは再起動後にのみ表示されます。
再起動したら、 /cc-probeline-install を実行します。OS が検出され、適切なチャネル (Homebrew / Scoop /curl) 経由でバイナリがインストールされ、ステータス ラインが表示され、何かを実行する前に確認が行われます。上記のどのチャネルでも手動でインストールできます。このプラグインでは、後でアップグレードするための /cc-probeline-update と /cc-probeline-config ウィザードも提供します。
cc-probeline --check
「インストール OK」と表示されます。
macOS、Linux、または Windows 上のクロード コード。
クォータ セグメント (5 時間 / 7 日の制限、追加使用量) の場合: クロード コード ≥ 2.1.80。ステータス行ペイロードで rate_limits を渡します。古いバージョンでは、クォータ セグメントは非表示になります。他はすべて正常に動作します。
Claude Code 内から対話型ウィザードを実行します。
/cc-プローブライン構成
プローブ、テーブル サイズ、バックグラウンド更新について説明し、TOML を作成します。または、 ~/.config/cc-probeline/config.toml を直接編集します ( cc-probeline check-config で検証します)。
【一般】
table_rows = 10 ターンごとのコスト テーブルの行数 (最大 40)
no_color = false # true = プレーンなモノクロ出力
Price_check = true # false = 価格を取得しません。テーブルを使う

ビルドに組み込まれている
use_refresh = true # false = クロード コードの /usage をバックグラウンドで実行しない
[ widget ] # セグメントのオン/オフを切り替えます
モデル = 真
努力＝真実
コスト = 真
ctx = true
クォータ = true
git = true
プロジェクト = 真
メールアドレス = true
時間 = 真
[しきい値]
cost_budget_usd = 25 # $25 を超えるコストセグメントを赤にします (0 = オフ)
# コンテキストバーの色が反転します - 黄色/オレンジ/赤。
# 厳密に増加する必要があります。不適切な値はこれらのデフォルトに戻ります。
ctx_notice_ratio = 0.50
ctx_warn_ratio = 0.70
ctx_critical_ratio = 0.90
# レート制限ウィンドウごとに同じ 3 つのフリップ。 7d ウィンドウはこれらのキーを反映します
#uota_7d_notice_ratio / _warn_ratio / _critical_ratio として。
quota_5h_notice_ratio = 0.50
quota_5h_warn_ratio = 0.70
quota_5h_critical_ratio = 0.90
構成は、CC_PROBELINE_CONFIG=/path (明示的なオーバーライド) → 現在のリポジトリの .cc-probeline.toml (プロジェクトローカル) → ~/.config/cc-probeline/config.toml (グローバル) の優先順位で読み取られます。すべてのフィールドはオプションです。欠落しているフィールドには組み込みのデフォルトが使用され、無効な値によってステータス行が中断されることはなく、デフォルトに戻ります。
完全なリファレンス: scripts/config.toml.example 。
新しいリリースがリリースされると、ステータス行が表示されます: ↑ update: vX → vY — /cc-probeline-update を実行します。 Claude Code 内でそのコマンドを実行すると、インストールしたチャネルを通じてアップグレードされます (バイナリが見つからない場合は自動的にインストールされます)。または手動で更新します。
brew upgrade labzink/homebrew-tap/cc-probeline # Homebrew
スクープ更新 cc-probeline # スクープ
カール -fsSL https://raw.githubusercontent.com/labzink/cc-probeline/main/scripts/install.sh | sh #curl (最新のものを再実行)
更新通知は、1 日に 1 回の価格/バージョンのチェックによって行われます。 Price_check = false (または /cc-probeline-config ウィザード) でオフにすると、cc-probeline はそのままになります。

オフラインです。更新すると、ステータスラインの配線がそのまま維持されます。
アンインストールすると、以前のステータス行が復元され (cc-probeline で置き換えられた場合はバイトごと)、バイナリが削除されます。インストールしたチャネルのコマンドを使用します。その後、Claude Code を再起動します。
brew uninstall cc-probeline # Homebrew — 以前のステータス行も復元します
カール -fsSL https://raw.githubusercontent.com/labzink/cc-probeline/main/scripts/install.sh | sh -s -- --アンインストール #curl
Scoop (Windows): 最初にステータス行を復元し、次にバイナリを削除します。Scoop アンインストールは復元ステップ自体を実行できません。
cc - プローブラインのアンインストール
スクープ アンインストール CC - プローブライン
バイナリを削除せずにステータス行の接続を解除するだけするには、cc-probeline uninstall を単独で実行します。
これは自分のために作りました。私はクロード コードの内部を知りたかったのです。お金と制限が実際にどこに行くのかを知りたかったのです。実用的なバージョンには数日かかりましたが、それが私にとってのみ役に立たないことは明らかでした。
その後に開発されたのは、さらなる機能ではなく、プロトタイプではなく製品にすることでした。私がコードを書いたのではなく、クロード コードがすべての行を書きました。私の興味は別のところにあり、実際に機能する AI を使用した開発プロセスを構築することでした。
実際には、これはコードの前にコンセプトと仕様を意味し、固定された契約に基づいて段階的に作業し、すべての作業を実行します。

[切り捨てられた]

## Original Extract

See where it leaks, stop paying for it — a live Claude Code status line that prices every turn, your subagents, cache rebuilds, plus limits, context and git. - labzink/cc-probeline

GitHub - labzink/cc-probeline: See where it leaks, stop paying for it — a live Claude Code status line that prices every turn, your subagents, cache rebuilds, plus limits, context and git. · GitHub
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
labzink
/
cc-probeline
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
257 Commits 257 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows assets assets cmd/ cc-probeline cmd/ cc-probeline commands commands hooks hooks internal internal scripts scripts tests tests .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
A status line for Claude Code that prices every turn — yours, your subagents', and every cache rebuild — reading the session log already on your disk. No account, no API key, no telemetry.
Session totals come from Claude Code itself; the per-turn breakdown is computed here from token counts against a public price table, and labelled as an estimate throughout.
It makes no network calls while it runs. Two optional, opt-out background actions exist: a once-a-day check of one public file on GitHub (current prices and the latest version), and a refresh of Claude Code's own usage cache — its /usage screen run headlessly, once when a session starts and rarely after that, and only while your plan actually has a model-scoped weekly limit or paid extra usage to keep current. Neither sends anything about your session. Privacy policy →
Spend your limits on purpose, instead of paying for inefficiency you can't see.
brew install labzink/homebrew-tap/cc-probeline # macOS
curl -fsSL https://raw.githubusercontent.com/labzink/cc-probeline/main/scripts/install.sh | sh # macOS / Linux
Windows (Scoop) and the Claude Code plugin marketplace: all install options → . Every release ships with SHA256 checksums and signed SLSA build provenance .
Most status lines count things — tokens, turns, running agents. The probe prices them. Everything below comes out of your session's local log: data Claude Code has, but never shows you.
Every turn, priced — not one opaque session total: a live table where each step lands with its own cost.
What your subagents spend — subagent work is invisible while it runs. The probe puts each agent on the bill, live, next to your own turns.
Cache rebuilds, in dollars — idle past the TTL (60 min for the orchestrator, 5 for subagents), and your next turn quietly rewrites the whole cache. The probe ages it live (⏱ 60m → 0m) and prices the rebuild when it hits.
Extra usage in money, not percent — once your plan is spent, the line shows what you have actually paid on top this month against your ceiling ($20.40 / $120.00), straight from Anthropic's own figure rather than an estimate.
Prices that stay correct — your dollars are only as honest as the price table behind them. cc-probeline refreshes its rates over the network — one optional, opt-out check a day, never during render — so when Anthropic changes prices your totals follow within a day, no reinstall. Offline or opted out, it falls back to the table baked into the build.
5h / 7d limits with reset clocks (plus the model-scoped Fable cap, if your plan carries one) — watch them fill, know exactly when they free up, without stopping to open /usage.
Colour-coded zones — numbers shift colour as they enter warning and critical territory, so the line catches your eye exactly when it should.
Plus the table stakes: model, context, git, session time.
Every turn lands on its own line — orchestrator and subagents alike — priced as it happens. Finally you see where every dollar of your reasoning actually goes.
Built to fit your terminal. Don't like a segment, the colours, or the width? The /cc-probeline-config wizard walks you through it and writes the config for you — no hand-editing TOML. (That's the hint at the bottom of the dashboard above.)
The moment you cross 100%, you'll see it — and the extra bill stays under your control.
You get warned while there's still time to act — not after you've hit the wall.
Cache rebuilds stop being silent — you see the price the moment they happen.
All of it read, none of it touched — that's why it's called a probe .
A probe is an instrument of observation, not intervention. Everything cc-probeline does is read and display — it never reaches into your account or reports on you.
What it reads: your session's JSONL log ( ~/.claude/projects/… ) and the status-line payload Claude Code pipes directly to it.
What it doesn't touch: credentials, keychain, OAuth tokens — no telemetry, ever. Rendering is fully offline. Its own only network call is one optional, opt-out price/version check a day — a plain download of a public file, sending nothing about your session. It can also ask Claude Code to refresh its usage cache in the background ( /usage , headless, throttled machine-wide, opt-out) so a model-scoped weekly limit stays current: that call is Claude Code's, with your credentials, and costs no model tokens. Turn both off and cc-probeline neither sends nor starts anything.
The binary: single compiled Go binary, no runtime dependencies, one run ≈ 5 ms.
Auditable: MIT license, open source, every release published with SHA256 checksums and signed build provenance (SLSA) — verify any download with gh attestation verify <file> --repo labzink/cc-probeline .
Every channel below does the same thing: install the binary and wire it into
your Claude Code status line. After installing, restart Claude Code and you're
done — no extra commands. (If you already have a custom status line, it's left
untouched; switch to cc-probeline with cc-probeline install --merge-settings --force .)
Homebrew (macOS — it's a cask; on Linux use curl below):
brew install labzink/homebrew-tap/cc-probeline
curl (macOS / Linux — downloads the release archive for your OS, verifies SHA256, installs the binary):
curl -fsSL https://raw.githubusercontent.com/labzink/cc-probeline/main/scripts/install.sh | sh
Scoop (Windows, experimental):
scoop bucket add labzink https: // github.com / labzink / scoop - bucket
scoop install cc - probeline
Claude Code plugin marketplace:
/plugin marketplace add labzink/cc-probeline
Then install the plugin from the /plugin menu (or /plugin install cc-probeline ) and restart Claude Code — the slash commands below only show up after a restart.
Once restarted, run /cc-probeline-install : it detects your OS, installs the binary through the right channel (Homebrew / Scoop / curl) and wires the status line — asking before it runs anything. You can still install manually with any channel above. The plugin also gives you /cc-probeline-update to upgrade later and the /cc-probeline-config wizard.
cc-probeline --check
Prints Installation OK .
Claude Code on macOS, Linux, or Windows.
For the quota segment (5h / 7d limits, extra usage): Claude Code ≥ 2.1.80, which passes rate_limits in the status-line payload. On older versions the quota segment is hidden; everything else works normally.
Run the interactive wizard from inside Claude Code:
/cc-probeline-config
It walks you through probes, table size and background refresh — and writes the TOML for you. Or edit ~/.config/cc-probeline/config.toml directly (validate with cc-probeline check-config ):
[ general ]
table_rows = 10 # rows in the per-turn cost table (max 40)
no_color = false # true = plain monochrome output
price_check = true # false = never fetch prices; use the table baked into the build
usage_refresh = true # false = never run Claude Code's /usage in the background
[ widgets ] # flip any segment on/off
model = true
effort = true
cost = true
ctx = true
quota = true
git = true
project = true
email = true
time = true
[ thresholds ]
cost_budget_usd = 25 # turn the cost segment red past $25 (0 = off)
# Colour flips for the context bar — yellow / orange / red.
# Must strictly increase; bad values fall back to these defaults.
ctx_notice_ratio = 0.50
ctx_warn_ratio = 0.70
ctx_critical_ratio = 0.90
# Same three flips per rate-limit window. The 7d window mirrors these keys
# as quota_7d_notice_ratio / _warn_ratio / _critical_ratio.
quota_5h_notice_ratio = 0.50
quota_5h_warn_ratio = 0.70
quota_5h_critical_ratio = 0.90
Config is read in precedence order: CC_PROBELINE_CONFIG=/path (explicit override) → .cc-probeline.toml in the current repo (project-local) → ~/.config/cc-probeline/config.toml (global). Every field is optional; missing fields use the built-in defaults, and an invalid value never breaks the status line — it falls back to the default.
Full reference: scripts/config.toml.example .
When a newer release is out, the status line surfaces it: ↑ update: vX → vY — run /cc-probeline-update . Run that command inside Claude Code and it upgrades through whichever channel you installed with (and installs it for you if the binary is missing). Or update by hand:
brew upgrade labzink/homebrew-tap/cc-probeline # Homebrew
scoop update cc-probeline # Scoop
curl -fsSL https://raw.githubusercontent.com/labzink/cc-probeline/main/scripts/install.sh | sh # curl (re-runs latest)
The update notice comes from a once-a-day price/version check; turn it off with price_check = false (or in the /cc-probeline-config wizard) and cc-probeline stays fully offline. Updating keeps your status-line wiring intact.
Uninstalling restores the status line you had before (byte-for-byte, if cc-probeline replaced one) and removes the binary. Use the command for the channel you installed with — restart Claude Code afterwards :
brew uninstall cc-probeline # Homebrew — also restores your previous status line
curl -fsSL https://raw.githubusercontent.com/labzink/cc-probeline/main/scripts/install.sh | sh -s -- --uninstall # curl
Scoop (Windows): restore the status line first, then remove the binary — scoop uninstall can't run the restore step itself:
cc - probeline uninstall
scoop uninstall cc - probeline
To only un-wire the status line without removing the binary, run cc-probeline uninstall on its own.
I built this for myself. I wanted to see under the hood of Claude Code — where the money and the limits actually go. A working version took a few days, and it was clear it wouldn't be useful only to me.
What came after went into making it a product rather than a prototype, not into more features. I didn't write the code — Claude Code wrote every line of it. My interest was elsewhere: building a development process with AI that actually holds up.
In practice that means concept and spec before code, work in phases against a fixed contract, and every c

[truncated]
