---
source: "https://github.com/dhkts1/teamclaude-rs"
hn_url: "https://news.ycombinator.com/item?id=48955755"
title: "Teamclaude rewrite in rust, run multiple Claude accounts in parllel"
article_title: "GitHub - dhkts1/teamclaude-rs: Lean single-user rotating Anthropic proxy with a live TUI — Rust rewrite of teamclaude · GitHub"
author: "dhkts1"
captured_at: "2026-07-18T06:53:40Z"
capture_tool: "hn-digest"
hn_id: 48955755
score: 1
comments: 0
posted_at: "2026-07-18T06:36:13Z"
tags:
  - hacker-news
  - translated
---

# Teamclaude rewrite in rust, run multiple Claude accounts in parllel

- HN: [48955755](https://news.ycombinator.com/item?id=48955755)
- Source: [github.com](https://github.com/dhkts1/teamclaude-rs)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T06:36:13Z

## Translation

タイトル: Rust で Teamclaude を書き換え、複数の Claude アカウントを並列で実行
記事のタイトル: GitHub - dhkts1/teamclaude-rs: ライブ TUI を使用したリーン シングルユーザー回転 Anthropic プロキシ — Teamclaude の Rust 書き換え · GitHub
説明: ライブ TUI を使用した無駄のないシングルユーザー回転 Anthropic プロキシ — Teamclaude の Rust 書き換え - dhkts1/teamclaude-rs

記事本文:
GitHub - dhkts1/teamclaude-rs: ライブ TUI を使用したリーン シングルユーザー回転 Anthropic プロキシ — Teamclaude の Rust 書き換え · GitHub
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
dhkts1
/
チームクロード-RS
公共
通知
通知設定を変更するにはサインインする必要があります
追加のna

違反オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .githooks .githooks .github .github アセット アセット src src .gitignore .gitignore .gitleaks.toml .gitleaks.toml Cargo.lock Cargo.lock Cargo.toml Cargo.toml DESIGN.md DESIGN.md ライセンス ライセンス MITM-DESIGN.md MITM-DESIGN.md 通知 通知 README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Rust の無駄のないシングルユーザー回転 Anthropic プロキシ。クロード コードを指定します (または
任意の Anthropic API クライアント) でリクエストを複数のクロードに分散します
アカウントを更新し、OAuth トークンを自動的に更新し、ライブ TUI を表示します。
アカウントごとのクォータとリクエスト数。
これは、Node プロキシ チームクロードを Rust で最初から書き直したものです。同じディスク上の構成、同じ証明書、
したがって、同じポート上のドロップインです。
上のダッシュボードは tcr デモ、つまり偽のアカウントでレンダリングされた本物の TUI です。 tcr デモを使用して自分で実行してください。
アカウント間の負荷分散: 最も最近選択されたものをローテーションするため、
単一のアカウントが打撃を受けます。優先順位が尊重されます (「枕」を維持します)
アカウントは最後の手段として使用してください）。
1 つのポートで 2 つのエントリ モード:
ベース URL: ANTHROPIC_BASE_URL=http://127.0.0.1:3456 を設定します。
フォワードプロキシ (MITM): HTTPS_PROXY=http://127.0.0.1:3456 + を設定します。
NODE_EXTRA_CA_CERTS=<ca.pem> 。 api.anthropic.com のみが傍受され、
トークン注入。他のすべてのホストはブラインド トンネルになっているため、Claude Code の他のホストは
エンドポイントは動作し続けます。
ゼロ支出クォータ プローブ: OAuth 使用状況から各アカウントの使用状況を読み取ります。
エンドポイント (メッセージ クォータが消費されない) なので、アイドル状態でもバーが最新の状態に保たれます。
正直なライブ TUI: アカウントごとのステータス、5 時間 / 7 日のクォータ バー、プローブの健全性、
リクエスト数と最近のリクエストログ。制限に近いアカウントは次のように読み取られます
「ニア」/「フル」 — 決して誤った「エラー」ではありません。
Localhost のみ: 127.0.0.1 をバインドするため、loc

すべてのクライアントには API キーは必要ありません。
カーゴビルド --release
ln -sfn " $PWD /target/release/tcr " ~ /.local/bin/tcr # `tcr` を PATH に置きます
設定する
構成は ~/.config/teamclaude.json にあります。
{
"プロキシ" : { "ポート" : 3456 },
"スイッチしきい値" : 0.90 、
"quotaProbeSeconds" : 75 、
「アカウント」: [
{
"名前" : " you@example.com " ,
"type" : " oauth " ,
"accessToken" : " sk-ant-oat01-... " ,
"refreshToken" : " sk-ant-ort01-... " ,
"有効期限": 1893456000000 、
「優先度」：0
}
]
}
優先順位は最も低い勝ちです (デフォルトは 0)。バックアップアカウントに大きな番号を与えると、
これらは、プライマリーが上限に近い場合にのみ使用されます。
アカウントの追加: tcr ログインを実行します。Anthropic の OAuth について説明します。
ブラウザ フロー (PKCE) を実行し、結果のトークンを構成に直接書き込みます。
既存の OAuth トークンを手動で貼り付けることもできます。いずれにしても、ファイルは
0600と書かれています。決してコミットしないでください。
tcr # ライブ TUI でプロキシを開始します (終了するには q)
tcr サーバー --headless # バックグラウンドで実行し、標準出力にログを記録します
tcr run -- < args > # すでにプロキシを指している `claude` を起動します
次に、上記のように ANTHROPIC_BASE_URL / HTTPS_PROXY をエクスポートするか、 tcr run を使用します。
元のノードで実行されていたもの (および実行されなかったものもいくつか) はすべて実装されています: OAuth
login 、モデルごとの (Fable-aware) ルーティング、アカウント CLI ( accounts/remove/
優先度 / 有効 / 無効 / ステータス )、更新、保温、およびセッション
親和性。 2 つは ~/.config/teamclaude.json 経由でオプトインされており、デフォルトではオフになっています。
"sessionAffinity": true — クライアント セッションを 1 つのアカウントに固定します
一生。 Anthropic のプロンプト キャッシュはアカウントごとであるため、リクエストごとにローテーションします
毎ターンコールドキャッシュを与えます。アフィニティは、セッションのキャッシュをその上でウォームに保ちます。
アカウント間で異なるセッションが依然として分散している間、
"warmupSeconds": <n> — アイドル状態のアカウントを定期的にウォームアップして、5 時間のウィンドウを維持します
アクティブなままです。

これは実際のクォータを消費するため、意図的に有効にしてください。
127.0.0.1 のみをバインドします。フォワードプロキシ トンネルには、
ローカル ユーザー (シェル自体よりも幅の広いユーザーではありません)。
api.anthropic.com のみが TLS で終了します。他のすべてのホストはパススルーです
バイトトンネル (復号化されない)。
Config とリーフ キーは 0600 です。このリポジトリにはシークレットは含まれません。
シークレット スキャン: .githooks/pre-commit フックが実行される gitleaks git --staged
(config: .gitleaks.toml ) したがって、シークレットはコミットに到達しません。後で有効にします
git config core.hooksPath .githooks を使用してクローンを作成します (PATH に gitleaks が必要です)。
MIT — 「ライセンス」を参照してください。これはRustで書き直したものです
KarpelesLab/チームクロード (MIT);の
オリジナルの著作権とライセンスは NOTICE に保持されます。
ライブ TUI を使用した無駄のないシングルユーザー回転 Anthropic プロキシ — Teamclaude の Rust 書き換え
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Lean single-user rotating Anthropic proxy with a live TUI — Rust rewrite of teamclaude - dhkts1/teamclaude-rs

GitHub - dhkts1/teamclaude-rs: Lean single-user rotating Anthropic proxy with a live TUI — Rust rewrite of teamclaude · GitHub
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
dhkts1
/
teamclaude-rs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .githooks .githooks .github .github assets assets src src .gitignore .gitignore .gitleaks.toml .gitleaks.toml Cargo.lock Cargo.lock Cargo.toml Cargo.toml DESIGN.md DESIGN.md LICENSE LICENSE MITM-DESIGN.md MITM-DESIGN.md NOTICE NOTICE README.md README.md View all files Repository files navigation
A lean, single-user rotating Anthropic proxy in Rust. Point your Claude Code (or
any Anthropic API client) at it and it spreads requests across several Claude
accounts, refreshes their OAuth tokens automatically, and shows a live TUI with
per-account quota and request counts.
It's a from-scratch Rust rewrite of the Node proxy teamclaude — same on-disk config, same certs,
so it's a drop-in on the same port.
The dashboard above is tcr demo — the real TUI rendered with fake accounts. Run it yourself with tcr demo .
Load-balances across your accounts: least-recently-selected rotation, so no
single account gets hammered. Priority tiers are respected (keep "pillow"
accounts as last resort).
Two entry modes on one port:
Base-URL: set ANTHROPIC_BASE_URL=http://127.0.0.1:3456 .
Forward-proxy (MITM): set HTTPS_PROXY=http://127.0.0.1:3456 +
NODE_EXTRA_CA_CERTS=<ca.pem> . Only api.anthropic.com is intercepted and
token-injected; every other host is blind-tunneled, so Claude Code's other
endpoints keep working.
Zero-spend quota probe: reads each account's usage from the OAuth usage
endpoint (no message quota spent) so the bars stay fresh even when idle.
Honest live TUI: per-account status, 5h / 7d quota bars, probe health,
request counts, and a recent-request log. Near-limit accounts read as
"near"/"full" — never a false "error".
Localhost only: binds 127.0.0.1 , so local clients need no API key.
cargo build --release
ln -sfn " $PWD /target/release/tcr " ~ /.local/bin/tcr # put `tcr` on PATH
Configure
Config lives at ~/.config/teamclaude.json :
{
"proxy" : { "port" : 3456 },
"switchThreshold" : 0.90 ,
"quotaProbeSeconds" : 75 ,
"accounts" : [
{
"name" : " you@example.com " ,
"type" : " oauth " ,
"accessToken" : " sk-ant-oat01-... " ,
"refreshToken" : " sk-ant-ort01-... " ,
"expiresAt" : 1893456000000 ,
"priority" : 0
}
]
}
priority is lowest-wins (default 0); give backup accounts a higher number so
they're used only when the primaries are near their cap.
Adding accounts: run tcr login — it walks you through Anthropic's OAuth
browser flow (PKCE) and writes the resulting tokens straight into the config.
You can also paste existing OAuth tokens in by hand. Either way the file is
written 0600 ; never commit it.
tcr # start the proxy with the live TUI (q to quit)
tcr server --headless # run in the background, log to stdout
tcr run -- < args > # launch `claude` already pointed at the proxy
Then either export ANTHROPIC_BASE_URL / HTTPS_PROXY as above, or use tcr run .
Everything the Node original did (and a couple it didn't) is implemented: OAuth
login , per-model (Fable-aware) routing, the account CLI ( accounts / remove /
priority / enable / disable / status ), update , keep-warm, and session
affinity. Two are opt-in via ~/.config/teamclaude.json , off by default:
"sessionAffinity": true — pin a client session to one account for its
lifetime. Anthropic's prompt cache is per-account , so per-request rotation
gives every turn a cold cache; affinity keeps a session's cache warm on its
account while different sessions still spread across accounts.
"warmupSeconds": <n> — periodically warm idle accounts so their 5-hour window
stays active. This one spends real quota , so enable it deliberately.
Binds 127.0.0.1 only; the forward-proxy tunnel is reachable solely by the
local user (no wider than the shell itself).
Only api.anthropic.com is TLS-terminated; all other hosts are pass-through
byte tunnels (never decrypted).
Config and leaf key are 0600 . No secrets belong in this repo.
Secret scanning: a .githooks/pre-commit hook runs gitleaks git --staged
(config: .gitleaks.toml ) so no secret reaches a commit. Enable it after
cloning with git config core.hooksPath .githooks (needs gitleaks on PATH).
MIT — see LICENSE . This is a Rust rewrite of
KarpelesLab/teamclaude (MIT); the
original's copyright and license are preserved in NOTICE .
Lean single-user rotating Anthropic proxy with a live TUI — Rust rewrite of teamclaude
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
