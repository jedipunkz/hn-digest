---
source: "https://github.com/Elpulgo/polysbx"
hn_url: "https://news.ycombinator.com/item?id=48664294"
title: "Show HN: Boostrap to try out Claude Code in Docker or sbx or microsandbox"
article_title: "GitHub - Elpulgo/polysbx: A wrapper for using/onboarding to multiple sbx solutions for AI agent harness · GitHub"
author: "elpulgo"
captured_at: "2026-06-24T19:32:14Z"
capture_tool: "hn-digest"
hn_id: 48664294
score: 1
comments: 0
posted_at: "2026-06-24T19:01:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Boostrap to try out Claude Code in Docker or sbx or microsandbox

- HN: [48664294](https://news.ycombinator.com/item?id=48664294)
- Source: [github.com](https://github.com/Elpulgo/polysbx)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T19:01:35Z

## Translation

タイトル: HN を表示: Docker、sbx、または microsandbox でクロード コードを試すためのブースストラップ
記事タイトル: GitHub - Elpulgo/polysbx: AI エージェント ハーネス用の複数の sbx ソリューションを使用/オンボーディングするためのラッパー · GitHub
説明: AI エージェント ハーネス用の複数の sbx ソリューションを使用/オンボーディングするためのラッパー - Elpulgo/polysbx

記事本文:
GitHub - Elpulgo/polysbx: AI エージェント ハーネス用の複数の sbx ソリューションを使用/オンボーディングするためのラッパー · GitHub
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
エルプルゴ
/
ポリエスビー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット bin bin image image lib lib templates templates .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
1 つのコマンドでサンドボックス内でクロード コードを実行します。隔離を選択してください
バックエンド (Docker、Docker Sandbox、または
microsandbox )、認証 (サブスクリプション / API キー /
トークン)、およびどの言語ツールチェーンがイメージに含まれるか — Polysbx がオンボードになっており、
どのリポジトリでも動作する psb コマンドを提供します。
カール -fsSL https://raw.githubusercontent.com/elpulgo/polysbx/main/install.sh |バッシュ
# または、チェックアウトから:
./install.sh
install.sh は必要なシェル スクリプトをインストールします。
次のコマンド ラインが表示されます。
OS/アーキテクチャを検出し、前提条件をチェックします (検出と指示のみ - それ
あなたの代わりに何かをインストールすることはありません）
バックエンド、認証モード、言語、Claude 設定ディレクトリ、およびオプションを要求します
Azure DevOps / GitHub の統合
選択したモジュールのみを使用してツールチェーン イメージを構築します。 .NET / GO / Python / RUST。 (ノードは常にサンドボックスにインストールされます)
ワンタイム クロード ログインを実行します (サブスクリプション モード)
psb shim を ~/.local/bin にドロップします。
cd ~ /some/プロジェクト
psb # サンドボックスでクロードを起動します (デフォルトのサブコマンド)
psb " pkg/foo で失敗したテストを修正します "
PSB Doctor # インストールを確認します
psb build # イメージを再構築します (言語を変更した後)
psb setup-auth # 再ログイン (サブスクリプション)
psb clean # イメージ + ボリューム/サンドボックスを再利用する
psb clean --all # (sbx) も停止し、実行中のインスタンスを削除します
psb update # git-pull Polysbx 自体 (後で `psb build` を再実行)
バックエンド
3 つすべてがオンボーディング、言語モジュール、構成モデル、および psb を共有します。
表面。それらは、分離境界と資格情報が存在する場所が異なります。
仕組み (短いバージョン)

)
画像 = ベース + ノード + クロード コード + 選択した言語モジュール
(.NET 8/10、Go、Python、Rust) + オプションの azure-cli/gh 。一度建てた、
UID/GID が組み込まれているため、Linux と macOS でバインド マウントが整列します。ドッカー/msbxの使用
Debian ベース。 sbx は、同じモジュールを Docker の Claude-Code テンプレートに重ねます。
2 つの州の場所が別々に保たれています:
資格情報とセッション状態を保持する管理されたホーム - あなたがログインします
一度。これは、Docker ボリューム ( docker )、キャッシュ ホーム ディレクトリ ( msbx )、または
ホストキーチェーンプロキシ ( sbx )。
設定ディレクトリ (デフォルト ~/.claude )、読み取り専用、polysbx のソース
既知のサブパスをステージングします (スキル エージェント コマンド フック ルール settings.json )
Polysbx が所有するコピーにコピーします。設定ディレクトリには決して書き込まれません。
資格情報が構成マウントに入力されることはありません。
強化: docker 、 --cap-drop=ALL 、 --read-only 、
--no-new-privileges 、 tmpfs スクラッチ、およびリソース制限。 msbx / sbxの場合、
(マイクロ)VM 境界はそれを包含します。すべてのバックエンドは動作のガードレールを維持します
(拒否リスト、保護されたブランチのプリプッシュフックをオプトイン) で Claude を実行します。
--dangerously-skip-permissions (サンドボックスが境界です)。
ネットワーク下り: NET_MODE=allowlist はサンドボックスをバックエンドごとに制限します
選択したツールチェーンから派生した許可リスト (デフォルトで開きます)。
config — バックエンド、認証モード、言語、構成ディレクトリ、統合、制限
( templates/config.example を参照)。
Secrets.env ( chmod 600 ) — ANTHROPIC_API_KEY または CLAUDE_CODE_OAUTH_TOKEN
(API/トークン モードのみ。サブスクリプションはここには何も保存されません)
AZURE_DEVOPS_EXT_PAT / GITHUB_TOKEN (これらの統合がオンになっている場合)。
いずれかを編集して psb を再実行する (起動ごとに構成が再ステージングされます) か、psb ビルドを再実行します
(言語を変更した後)。
msbx : サンドボックス (virtiofs キャッシュ) 内のファイルでの入出力エラー
msbx バックエンドでは、プロジェクトは次の場所で共有されます。

virtiofs 経由で microVM に接続します。
Git は通常、サンドボックス内で正常に動作します (コミットを含む)。でも変異したら
サンドボックスの実行中にホストから同じファイルを送信する — 例:からコミットする
ホストターミナル、または別のツール (別のエージェント、エディタ、git gc ) を実行します。
同じ .git — ゲストはファイルの古いキャッシュ エントリを保持したままになる可能性があります
ホストが置き換えたばかりの i ノード。多くの git 操作は、次の書き込みによってファイルを書き換えます。
一時ファイルを作成し、元のファイル ( .git/config 、 refs 、
インデックス)、i ノードを交換し、これをトリップします。症状は単一のファイルが実行されることです
入出力エラーはサンドボックス内の ls から消えますが、
ホストはそれを正常に読み取ります。それは破損でもマウントの欠落でもありません。
リポジトリが正しくマップされています。これは virtiofs のメタデータ キャッシュの一貫性の制限です。
microVM ランタイム (msbx がベータ版である理由の 1 つ)、msb run はキャッシュを公開しない
モードをオンにします。
回復: Claude を終了し、psb を再起動します。新しい microVM が新しいものを取得します。
virtiofs セッションを開始し、ホストからファイルを再読み取りします。
避けてください: 同じものに対して git (または .git を書き込む別のツール) を使用しないでください。
msbx セッションがライブ中にホストからリポジトリを取得します。内部からコミットする
サンドボックスまたはホスト、同時に両方を使用することはできません。
AI エージェント ハーネス用の複数の sbx ソリューションを使用/オンボーディングするためのラッパー
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

A wrapper for using/onboarding to multiple sbx solutions for AI agent harness - Elpulgo/polysbx

GitHub - Elpulgo/polysbx: A wrapper for using/onboarding to multiple sbx solutions for AI agent harness · GitHub
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
Elpulgo
/
polysbx
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits bin bin image image lib lib templates templates .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
Run Claude Code inside a sandbox with one command. Pick your isolation
backend (Docker, Docker Sandboxes , or
microsandbox ), your auth (subscription / API key /
token), and which language toolchains go in the image — polysbx onboards you and
gives you a psb command that works in any repo.
curl -fsSL https://raw.githubusercontent.com/elpulgo/polysbx/main/install.sh | bash
# or, from a checkout:
./install.sh
install.sh installs the needed shell scripts.
A command line is shown which:
detects your OS/arch and checks prerequisites ( detect-and-instruct only — it
never installs anything for you )
asks for backend, auth mode, languages, your Claude config dir, and optional
Azure DevOps / GitHub integrations
builds the toolchain image with just the modules you chose; .NET / GO / PYTHON / RUST. (Node is always installed in the sandbox)
runs a one-time Claude login (subscription mode)
drops a psb shim on ~/.local/bin .
cd ~ /some/project
psb # launch Claude in the sandbox (default subcommand)
psb " fix the failing test in pkg/foo "
psb doctor # verify the install
psb build # rebuild the image (after changing languages)
psb setup-auth # re-login (subscription)
psb clean # reclaim image + volumes/sandboxes
psb clean --all # (sbx) also stop + remove running instances
psb update # git-pull polysbx itself (re-run `psb build` after)
Backends
All three share the onboarding, language modules, config model, and psb
surface; they differ in the isolation boundary and where credentials live.
How it works (the short version)
Image = your base + Node + Claude Code + the language modules you picked
( .NET 8/10 , Go , Python , Rust ) + optional azure-cli / gh . Built once,
UID/GID baked in so bind-mounts line up on Linux and macOS. docker / msbx use
a Debian base; sbx layers the same modules onto Docker's Claude-Code template.
Two state locations, kept separate:
a managed home holding your credentials + session state — you log in
once. It's a Docker volume ( docker ), a cache-home dir ( msbx ), or the
host-keychain proxy ( sbx ).
your config dir (default ~/.claude ), read-only, from which polysbx
stages the known subpaths ( skills agents commands hooks rules settings.json )
into a polysbx-owned copy. Your config dir is never written to, and your
credentials never enter the config mount.
Hardening: for docker , --cap-drop=ALL , --read-only ,
--no-new-privileges , tmpfs scratch, and resource limits; for msbx / sbx the
(micro)VM boundary subsumes that. All backends keep the behavioural guardrails
(opt-in deny-list, protected-branch pre-push hook) and run Claude with
--dangerously-skip-permissions (the sandbox is the boundary).
Network egress: NET_MODE=allowlist restricts the sandbox to a per-backend
allow-list derived from your selected toolchains ( open by default).
config — backend, auth mode, languages, config dir, integrations, limits
(see templates/config.example ).
secrets.env ( chmod 600 ) — ANTHROPIC_API_KEY or CLAUDE_CODE_OAUTH_TOKEN
(only in api/token modes; subscription stores nothing here), plus
AZURE_DEVOPS_EXT_PAT / GITHUB_TOKEN when those integrations are on.
Edit either and re-run psb (config is re-staged each launch) or psb build
(after changing languages).
msbx : Input/output error on a file inside the sandbox (virtiofs cache)
On the msbx backend the project is shared into the microVM over virtiofs .
Git normally works fine inside the sandbox (commits included). But if you mutate
the same file from the host while the sandbox is running — e.g. commit from a
host terminal, or run another tool (another agent, an editor, git gc ) against
the same .git — the guest can be left holding a stale cache entry for a file
whose inode the host just replaced. Many git operations rewrite a file by writing
a temp file and rename() -ing it over the original ( .git/config , refs, the
index), which swaps the inode and trips this. The symptom is a single file going
Input/output error , vanishing from ls inside the sandbox, while the
host reads it fine . It is not corruption and not a missing mount — the whole
repo is mapped correctly; it's a virtiofs metadata-cache coherence limitation of
the microVM runtime (one reason msbx is beta), and msb run exposes no cache
mode to tune it away.
Recover: exit Claude and relaunch psb — a fresh microVM gets a fresh
virtiofs session and re-reads the file from the host.
Avoid: don't drive git (or another tool that writes .git ) against the same
repo from the host while an msbx session is live. Commit from inside the
sandbox or the host, not both at once.
A wrapper for using/onboarding to multiple sbx solutions for AI agent harness
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
