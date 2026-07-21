---
source: "https://github.com/markmatsu/claude-logkeeper"
hn_url: "https://news.ycombinator.com/item?id=48992474"
title: "Show HN: Back up Claude Code sessions from Codespaces to a private repo"
article_title: "GitHub - markmatsu/claude-logkeeper: Auto-save Claude Code sessions to a private GitHub repo. Plaintext transcripts for reading, age-encrypted jsonl for full restore. · GitHub"
author: "markmatsushima"
captured_at: "2026-07-21T14:57:58Z"
capture_tool: "hn-digest"
hn_id: 48992474
score: 2
comments: 0
posted_at: "2026-07-21T14:00:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Back up Claude Code sessions from Codespaces to a private repo

- HN: [48992474](https://news.ycombinator.com/item?id=48992474)
- Source: [github.com](https://github.com/markmatsu/claude-logkeeper)
- Score: 2
- Comments: 0
- Posted: 2026-07-21T14:00:09Z

## Translation

タイトル: HN を表示: クロード コード セッションを Codespaces からプライベート リポジトリにバックアップする
記事タイトル: GitHub - markmatsu/claude-logkeeper: クロード コード セッションをプライベート GitHub リポジトリに自動保存します。読み取り用の平文トランスクリプト、完全な復元用の経過時間暗号化された JSONL。 · GitHub
説明: クロード コード セッションをプライベート GitHub リポジトリに自動保存します。読み取り用の平文トランスクリプト、完全な復元用の経過時間暗号化された JSONL。 - マークマツ/クロード-ログキーパー
HN テキスト: クロード コードをローカル マシンではなく GitHub コードスペースで実行することを好みますが、問題が 1 つあります。リモート コードスペースが削除または再構築されると、クロード コードのログが消えてしまいます。 Logkeeper は、クロード コード セッションが停止するたびに、人間が判読できる会話ログと完全な JSONL の経過時間で暗号化されたコピーを指定されたプライベート リポジトリに保存します。プライベート リポジトリを作成し、必要な Codespaces シークレットを追加し、フックと構成ファイルの小さなセットをプロジェクトに追加すると、バックアップが自動的に実行されます。

記事本文:
GitHub - markmatsu/claude-logkeeper: クロード コード セッションをプライベート GitHub リポジトリに自動保存します。読み取り用の平文トランスクリプト、完全な復元用の経過時間暗号化された JSONL。 · GitHub
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
マークマツ
/
クロードログキープ

えーっと
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .claude .claude .devcontainer .devcontainer .gitignore .gitignore DESIGN.md DESIGN.md LICENSE LICENSE README-Japanese.md README-Japanese.md README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ログキーパー — クロード コードの会話を GitHub に自動保存
2 つのフォルダーをプロジェクトにドロップし、git add とすべてのクロード コードを追加します。
セッションは 2 つのレイヤーでプライベート GitHub リポジトリに保存されます。
Transcripts/<repo>/<date>-<session>.md — 平文、ユーザーおよびクロード
テキストのみ。ツールの結果 (ファイルの内容、コマンド出力 - 秘密が漏洩しやすい場所)
最後に）は剥がされます。安全に閲覧したり、チーム内で共有したりできます。
raw/<repo>/<session>.jsonl.age — 完全な jsonl、age-encrypted 。のみ
秘密鍵の所有者はそれを復号化できます。 ~/.claude/projects/ に復元され、
claude --resume でセッションを元に戻します。
Index.json — ビューアのセッション メタデータ。
トリガー: クロード コード フック ( Stop = すべての応答後、SessionEnd = on
セッション終了）。コードスペースがセッション中に終了した場合、最後までのすべてが
完了した応答はすでに保存されています。
背景については、なぜ 2 つの層を使用したのか、およびその原因となった失敗について説明します。
設計 — 開発ノートを参照してください。
git 、 jq 、curl — ほとんどの開発環境にプリインストールされています
age — 小規模で最新のファイル暗号化
ツール。インストール:
他のプラットフォームと事前構築済みバイナリ: を参照してください。
年齢 README 。
gh (GitHub CLI) — オプションですが推奨されます。正確なプライベートリポジトリに使用されます
検証およびクローン/プッシュ フォールバックとして
Codespaces では、これは重要ではありません。setup.sh はすべてをインストールします。
Claude Code 自体は devcontainer 機能によってインストールされます。
両方のパスが一致

結論から言うと、.claude/ と .devcontainer/ をコピーします。
このリポジトリのフォルダーをプロジェクトに追加し、ログキーパーにプライベート ログを指定します。
リポジトリと年齢キー。
⚠️ .devcontainer/devcontainer.json をすでにお持ちですか?上書きしないでマージしてください。
すでに Codespace で Claude コードを実行している場合は、ほぼ確実に
スタック用に構成された独自の devcontainer.json。で上書きします
このリポジトリのコピーは開発環境を破壊します。このレポの
devcontainer.json は、まだ持っていない人のための実用的な例です。
ドロップイン交換品。
既存のファイルに追加する必要があるのは 2 つだけです。
Claude Code と gh をインストールし、 setup.sh を呼び出します。
前 (典型的な Next.js プロジェクト):
{
"name" : " example-app Next.js " ,
"image" : " mcr.microsoft.com/devcontainers/javascript-node:1-22-bookworm " ,
"forwardPorts" : [ 3000 ],
"postCreateCommand" : " npm install " ,
「カスタマイズ」: {
"vscode" : {
"extensions" : [ " dbaeumer.vscode-eslint " 、 " esbenp.prettier-vscode " ]
}
}
}
後 (マークされた 2 つの追加):
{
"name" : " example-app Next.js " ,
"image" : " mcr.microsoft.com/devcontainers/javascript-node:1-22-bookworm " ,
"forwardPorts" : [ 3000 ],
// 追加: Claude Code CLI + VS Code 拡張機能と gh をインストールします。
「機能」: {
"ghcr.io/anthropics/devcontainer-features/claude-code:1" : {},
"ghcr.io/devcontainers/features/github-cli:1" : {}
}、
// 変更: setup.sh を追加し、既存のコマンドを維持します
"postCreateCommand" : " npm install && bash .devcontainer/setup.sh " ,
「カスタマイズ」: {
"vscode" : {
"extensions" : [ " dbaeumer.vscode-eslint " 、 " esbenp.prettier-vscode " ]
}
}
}
差分は次のとおりです。
+ "機能": {
+ "ghcr.io/anthropics/devcontainer-features/claude-code:1": {},
+ "ghcr.io/devcontainers/features/github-cli:1": {}
+ }、
-「ぽ

stCreateCommand": "npm インストール",
+ "postCreateCommand": "npm install && bash .devcontainer/setup.sh",
自分自身のイメージを維持してください。このリポジトリのベース:ubuntu に切り替えないでください。
画像がスタック用に選択されます。画像にすでに含まれている場合は、
Node.js (上記の javascript-node など) では、ノード機能は必要ありません。
クロード コード機能では、ノードが存在するだけで済みます。画像に何もない場合は、
Node.js の場合、「ghcr.io/devcontainers/features/node:1」: {} を機能リストに追加します。
.devcontainer/setup.sh は新しいファイルなので、どこにも衝突しないのでコピーします。
そのままです。
devcontainer の変更は、新しいコンテナーでのみ有効になります。コミットした後、
コマンド パレットから Codespaces: Rebuild Container を実行するか、削除します。
コードスペースを削除し、新しいコードスペースを作成します。単純な停止/開始では、新しい設定は適用されません。
特徴。 (再構築せずに現在のコードスペースで作業を開始するには、
2 つの依存関係を手動でインストールします: sudo apt-get install -y age jq && chmod +x .claude/hooks/logkeeper.sh 。)
同様に、すでに .claude/settings.json がある場合は、フック ブロックをマージします。
ファイルを置き換えるのではなく、「既存の settings.json へのマージ」を参照してください。
以下。 (.claude/settings.local.json は別のファイルなので競合しません。)
このリポジトリには .claude/ と .devcontainer/ という 2 つのフォルダーが必要です。をクリックします。
このリポジトリの上部にある緑色の「コード」ボタン → ZIP をダウンロードして解凍し、
これら 2 つのフォルダーをプロジェクトにコピーします。 (これらは非表示のドットフォルダーです。macOS では
Cmd+Shift+を押します。 Finder で確認してください。)
プロジェクトにどちらのフォルダーもまだない場合は、これで完了です。そうなった場合は、マージに関する注意事項を参照してください
上記 — 実行時にこのリポジトリには他に何も必要ありません。
GitHub コードスペース内 (推奨)
2 つのフォルダーをプロジェクトにコピーします (上記を参照)。それがファイル全体です
ステップ — chmod や手動でのインストールは必要ありません。セットアップ

h はフックをマークします
実行可能ファイルで、 age 、 jq 、および Claude Code CLI + 拡張機能をインストールします。
コードスペースが作成されます。
3 つのコードスペース シークレットを設定します (設定 → コードスペース → シークレット)。
プロジェクトのリポジトリ。これはゼロ編集パスです。構成には一切触れません。
ファイル:
見落としがちですが、各シークレットに WORKING リポジトリへのアクセスを許可します。後
シークレットを作成すると、その「リポジトリ アクセス」はデフォルトで空になります (シークレット
リストには「リポジトリが 0 個」と表示されます）。リポジトリにアクセスできないシークレットは、
どの Codespace にも注入されないため、ログキーパーは LOGKEEPER_REPO 環境変数を作成したにもかかわらず設定されていないことを報告します。 3 つそれぞれについて
シークレットを編集して、Codespaces を実行するリポジトリへのアクセスを許可します。
(あなたのプロジェクト、例: you/my-app )。
これを PAT 自体のリポジトリ アクセスと混同しないでください。 2つあります
似たような名前の異なる設定:
シークレットのログ リポジトリを選択するのはよくある間違いです。実行しないでください。
ログ リポジトリからのコードスペースなので、値はどこにも挿入されず、
echo $LOGKEEPER_GH_TOKEN は空に戻ります。
シークレットはコンテナの起動時にのみ挿入されるため、これを変更した後は停止します
そしてコードスペースを再起動します。 echo $LOGKEEPER_REPO で確認します。
コミットしてプッシュします。
git add .claude .devcontainer && git commit -m "add logkeeper" && git Push
新しいコードスペースを作成します。すべては創造に結びついています。始める
クロードとセッションが保存を開始します。
コードスペースについてはこれで終わりです。ファイル ステップは実際には単なるコピーです。最初に注意してください
ビルドには数分かかります (「ベースイメージの選択」を参照)。
2 つのフォルダーをプロジェクトにコピーします (上記の「ファイルの取得」を参照)。
( .devcontainer/ は Codespaces によってのみ使用されますが、保持しても無害です。)
要件 ( age 、 jq 、 git ) をインストールします (「要件」を参照)。
フックを実行可能にします。これは Codespaces が行う 1 つのステップです
でもロケ地

セットアップには次のものが必要です。
chmod +x .claude/hooks/logkeeper.sh
環境変数（優先）：export LOGKEEPER_REPO="owner/claude-logs" および
シェル プロファイルで LOGKEEPER_PUBKEY="age1..." をエクスポートする、または
ファイルを編集します。リポジトリを .claude/logkeeper.conf に置き、パブリックに置きます。
.claude/logkeeper.pub にキーを入力します。
(ローカルでは LOGKEEPER_GH_TOKEN は必要ありません。通常の Git 認証情報がプッシュされます。
ログ リポジトリ)
コミット: git add .claude && git commit -m "add logkeeper"
.claude/settings.json がコミットされているため、フックはすべてのチームメイトに適用されます
誰がプルするか - 彼らは自分のキーとシークレットを追加するだけです。
LOGKEEPER_GH_TOKEN (きめ細かい PAT) を段階的に作成する
デフォルトの Codespaces トークンは、起動元のリポジトリにのみ到達できるため、
別の claude-logs リポジトリにプッシュするには、独自のトークンが必要です。きめの細かいのは、
この 1 つのリポジトリだけにロックできるため、クラシック トークンよりも優先されます。
たった 1 つの許可で。 1 分ほどかかります。
github.com で、プロフィール写真 (右上) → [設定] をクリックします。
左側のサイドバーで、一番下までスクロール→「開発者」
設定。
[パーソナル アクセス トークン] → [きめ細かいトークン] → [新規生成] をクリックします。
トークン。
トークン名: 任意、例:ログキーパー。
Expiration : 期間を選択します (きめ細かいトークンを無期限にすることはできません。
デフォルトでは 90 日が適切です。カレンダーのリマインダーを再生成するように設定します)。
リソース所有者: 自分のアカウント ( claude-logs を所有するアカウント)。
リポジトリ アクセス: [リポジトリのみを選択] を選択し、次にリポジトリを選択します。
クロードログリポジトリ。 (「すべてのリポジトリ」は使用しないでください。)
権限 → リポジトリ権限を展開 → コンテンツを検索し、
これを「読み取りおよび書き込み」に設定します。必要なのはそれだけです。 (「メタデータ:
読み取り専用」が自動的に追加されます。そのままにしておきます。)
[トークンの生成] をクリックし、今すぐトークンをコピーします。GitHub にはトークンのみが表示されます。
一度。失くしたら、ジュ

別のものを生成します。
これを LOGKEEPER_GH_TOKEN という名前の Codespaces シークレットとして追加します (上記のステップ 2)。
LOGKEEPER_GH_TOKEN の代わりに、コメントを解除することもできます。
カスタマイズは .devcontainer/devcontainer.json でブロックし、ハードコードします。
ログ リポジトリ名。 2 つの注意点: そのブロックは環境変数を読み取ることができません (GitHub は
コンテナが存在する前に作成されるため、リテラルの所有者/名前が必要です)。
新しく作成された Codespace にのみ影響し、既存の Codespace には影響しません。作成について
GitHub には、ワンクリックで書き込みアクセスを承認するためのプロンプトが表示されます。上記PATルート
これはより単純で、真のゼロ編集パスです。
キーの生成 (信頼できるローカル マシン上で)
age-keygen -o ~ /.config/age/logkeeper.txt
age1... 行は公開キーです (共有/コミットしても安全です)。
LOGKEEPER_PUBKEY または .claude/logkeeper.pub )。 AGE-SECRET-KEY-... 行は次のとおりです。
秘密キー: ローカルのパスワード マネージャー内にのみ保管し、GitHub には決して保管しないでください。
または、Codespaces シークレット内にあります。それを失うと、生のレイヤーが永久になります。
解読不能。チームのセットアップについては、以下の「チームでのログキーパーの使用」を参照してください。
既存の settings.json へのマージ
プロジェクトにすでに .claude/settings.json がある場合は、上書きせずにコピーしてください。
このリポジトリの settings.json のフック ブロックをあなたのものに追加します (Stop をマージします)
および SessionEnd 配列)。他のファイル (hooks/logk) をコピーします。

[切り捨てられた]

## Original Extract

Auto-save Claude Code sessions to a private GitHub repo. Plaintext transcripts for reading, age-encrypted jsonl for full restore. - markmatsu/claude-logkeeper

I prefer to run Claude Code in GitHub Codespaces rather than on my local machine, but there is one problem: Claude Code logs disappear when the remote Codespace is deleted or rebuilt. Logkeeper saves a human-readable conversation log and an age-encrypted copy of the full JSONL to a designated private repository whenever a Claude Code session stops. After creating the private repository, adding the required Codespaces secrets, and adding a small set of hook and configuration files to the project, the backups run automatically.

GitHub - markmatsu/claude-logkeeper: Auto-save Claude Code sessions to a private GitHub repo. Plaintext transcripts for reading, age-encrypted jsonl for full restore. · GitHub
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
markmatsu
/
claude-logkeeper
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .claude .claude .devcontainer .devcontainer .gitignore .gitignore DESIGN.md DESIGN.md LICENSE LICENSE README-Japanese.md README-Japanese.md README.md README.md View all files Repository files navigation
logkeeper — auto-save Claude Code conversations to GitHub
Drop two folders into your project, git add them, and every Claude Code
session gets saved to a private GitHub repo in two layers:
transcripts/<repo>/<date>-<session>.md — plaintext , user and Claude
text only. Tool results (file contents, command output — where secrets tend
to end up) are stripped. Safe-ish to browse and share within the team.
raw/<repo>/<session>.jsonl.age — the full jsonl, age-encrypted . Only
holders of a secret key can decrypt it. Restored into ~/.claude/projects/ ,
it brings the session back in claude --resume .
index.json — session metadata for a viewer.
Trigger: Claude Code hooks ( Stop = after every response, SessionEnd = on
session end). If your Codespace dies mid-session, everything up to the last
completed response is already saved.
For the background — why the two layers, and the failures that shaped the
design — see the development notes .
git , jq , curl — preinstalled in most dev environments
age — a small, modern file encryption
tool. Install:
Other platforms and prebuilt binaries: see the
age README .
gh (GitHub CLI) — optional but recommended; used for exact private-repo
verification and as a clone/push fallback
In Codespaces , none of this matters: setup.sh installs everything, and
Claude Code itself is installed by the devcontainer feature.
Both paths come down to the same thing: copy the .claude/ and .devcontainer/
folders from this repo into your project, then point logkeeper at a private logs
repo and an age key.
⚠️ Already have a .devcontainer/devcontainer.json ? Merge, don't overwrite.
If you're already running Claude Code in a Codespace, you almost certainly have
your own devcontainer.json configured for your stack. Overwriting it with
this repo's copy will break your dev environment. This repo's
devcontainer.json is a working example for people who don't have one yet, not
a drop-in replacement.
You only need to add two things to your existing file: the features that
install Claude Code and gh , and a call to setup.sh .
Before (a typical Next.js project):
{
"name" : " example-app Next.js " ,
"image" : " mcr.microsoft.com/devcontainers/javascript-node:1-22-bookworm " ,
"forwardPorts" : [ 3000 ],
"postCreateCommand" : " npm install " ,
"customizations" : {
"vscode" : {
"extensions" : [ " dbaeumer.vscode-eslint " , " esbenp.prettier-vscode " ]
}
}
}
After (the two additions marked):
{
"name" : " example-app Next.js " ,
"image" : " mcr.microsoft.com/devcontainers/javascript-node:1-22-bookworm " ,
"forwardPorts" : [ 3000 ],
// ADDED: installs the Claude Code CLI + VS Code extension, and gh
"features" : {
"ghcr.io/anthropics/devcontainer-features/claude-code:1" : {},
"ghcr.io/devcontainers/features/github-cli:1" : {}
},
// CHANGED: append setup.sh, keep your existing command
"postCreateCommand" : " npm install && bash .devcontainer/setup.sh " ,
"customizations" : {
"vscode" : {
"extensions" : [ " dbaeumer.vscode-eslint " , " esbenp.prettier-vscode " ]
}
}
}
The diff is only:
+ "features": {
+ "ghcr.io/anthropics/devcontainer-features/claude-code:1": {},
+ "ghcr.io/devcontainers/features/github-cli:1": {}
+ },
- "postCreateCommand": "npm install",
+ "postCreateCommand": "npm install && bash .devcontainer/setup.sh",
Keep your own image . Don't switch to this repo's base:ubuntu — your
image is chosen for your stack. Note that if your image already includes
Node.js (like javascript-node above), you do not need the node feature;
the Claude Code feature just needs Node to be present. If your image has no
Node.js, add "ghcr.io/devcontainers/features/node:1": {} to the features list.
.devcontainer/setup.sh is a new file and won't collide with anything, so copy
it in as-is.
devcontainer changes only take effect in a new container. After committing,
either run Codespaces: Rebuild Container from the command palette, or delete
the Codespace and create a fresh one. A plain stop/start will not apply the new
features. (To get going in your current Codespace without rebuilding, just
install the two dependencies by hand: sudo apt-get install -y age jq && chmod +x .claude/hooks/logkeeper.sh .)
Likewise, if you already have a .claude/settings.json , merge the hooks block
rather than replacing the file — see "Merging into an existing settings.json"
below. (A .claude/settings.local.json is a separate file and won't conflict.)
You need two folders from this repo: .claude/ and .devcontainer/ . Click the
green Code button at the top of this repo → Download ZIP , unzip it, and
copy those two folders into your project. (They're hidden dotfolders; on macOS
press Cmd+Shift+. in Finder to see them.)
That's it if your project has neither folder yet. If it does, see the merge notes
above — nothing else in this repo is needed at runtime.
In GitHub Codespaces (recommended)
Copy the two folders into your project (see above). That's the whole file
step — no chmod , no installing anything by hand. setup.sh marks the hook
executable and installs age , jq , and the Claude Code CLI + extension when
the Codespace is created.
Set three Codespaces secrets (Settings → Codespaces → Secrets), scoped to
your project repo. This is the zero-edit path — you never touch the config
files:
Easy to miss — grant each secret access to your WORKING repo. After
creating a secret, its "Repository access" is empty by default (the Secrets
list shows "0 repositories" ). A secret with no repository access is
not injected into any Codespace , so logkeeper reports LOGKEEPER_REPO env var is not set even though you created it. For each of the three
secrets, edit it and grant access to the repo you run Codespaces from
(your project, e.g. you/my-app ).
Don't confuse this with the PAT's own repository access. There are two
different settings with similar names:
Selecting the logs repo on the secret is the common mistake: you never run
a Codespace from the logs repo, so the value is injected nowhere and
echo $LOGKEEPER_GH_TOKEN comes back empty.
Secrets are injected only at container start, so after changing this, stop
and restart the Codespace. Verify with echo $LOGKEEPER_REPO .
Commit and push:
git add .claude .devcontainer && git commit -m "add logkeeper" && git push
Create a new Codespace. Everything is wired up on creation; start
claude and sessions begin saving.
That's it for Codespaces — the file step really is just a copy. Note the first
build takes a couple of minutes (see "Choosing a base image").
Copy the two folders into your project (see "Getting the files" above).
( .devcontainer/ is only used by Codespaces, but it's harmless to keep.)
Install the requirements — age , jq , git (see Requirements ).
Make the hook executable. This is the one step Codespaces does for you
but a local setup needs:
chmod +x .claude/hooks/logkeeper.sh
Env vars (take precedence): export LOGKEEPER_REPO="owner/claude-logs" and
export LOGKEEPER_PUBKEY="age1..." in your shell profile, or
Edit the files: put your repo in .claude/logkeeper.conf and your public
key in .claude/logkeeper.pub .
(No LOGKEEPER_GH_TOKEN needed locally — your normal git credentials push to
the logs repo.)
Commit: git add .claude && git commit -m "add logkeeper"
Because .claude/settings.json is committed, the hooks apply to every teammate
who pulls — they only add their own key and secrets.
Creating the LOGKEEPER_GH_TOKEN (fine-grained PAT), step by step
The default Codespaces token can only reach the repo you launched from, so
pushing to a separate claude-logs repo needs your own token. Fine-grained is
preferred over a classic token because it can be locked to just this one repo
with just one permission. It takes about a minute:
On github.com, click your profile picture (top-right) → Settings .
In the left sidebar, scroll all the way to the bottom → Developer
settings .
Click Personal access tokens → Fine-grained tokens → Generate new
token .
Token name : anything, e.g. logkeeper .
Expiration : pick a duration (fine-grained tokens can't be non-expiring;
90 days is a fine default — set a calendar reminder to regenerate).
Resource owner : your own account (the one that owns claude-logs ).
Repository access : choose Only select repositories , then select your
claude-logs repo. (Don't use "All repositories".)
Permissions → expand Repository permissions → find Contents and
set it to Read and write . That's the only one you need. ("Metadata:
Read-only" gets added automatically — leave it.)
Click Generate token , then copy the token now — GitHub shows it only
once. If you lose it, just generate another.
Add it as a Codespaces secret named LOGKEEPER_GH_TOKEN (step 2 above).
Alternative to LOGKEEPER_GH_TOKEN : you can instead uncomment the
customizations block in .devcontainer/devcontainer.json and hardcode the
logs repo name. Two gotchas: that block cannot read env vars (GitHub reads
it before the container exists, so it needs a literal owner/name ), and it
only affects newly created Codespaces , not existing ones. On creation
GitHub shows a one-click prompt to authorize write access. The PAT route above
is simpler and is the true zero-edit path.
Generating keys (on a trusted local machine)
age-keygen -o ~ /.config/age/logkeeper.txt
The age1... line is the public key (safe to share/commit, and what goes in
LOGKEEPER_PUBKEY or .claude/logkeeper.pub ). The AGE-SECRET-KEY-... line is
the secret key: keep it locally and in a password manager only, never on GitHub
or in a Codespaces secret. Losing it makes the raw layer permanently
undecryptable. For team setups, see "Using logkeeper with a team" below.
Merging into an existing settings.json
If your project already has .claude/settings.json , don't overwrite it — copy
the hooks block from this repo's settings.json into yours (merge the Stop
and SessionEnd arrays). Copy the other files ( hooks/logk

[truncated]
