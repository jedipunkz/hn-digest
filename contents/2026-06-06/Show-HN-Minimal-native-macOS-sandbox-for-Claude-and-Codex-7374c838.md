---
source: "https://github.com/sheremetyev/sandfence"
hn_url: "https://news.ycombinator.com/item?id=48425674"
title: "Show HN: Minimal native macOS sandbox for Claude and Codex"
article_title: "GitHub - sheremetyev/sandfence: Minimal native macOS sandbox for Claude Code and Codex · GitHub"
author: "sheremetyev"
captured_at: "2026-06-06T15:31:46Z"
capture_tool: "hn-digest"
hn_id: 48425674
score: 1
comments: 0
posted_at: "2026-06-06T14:53:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Minimal native macOS sandbox for Claude and Codex

- HN: [48425674](https://news.ycombinator.com/item?id=48425674)
- Source: [github.com](https://github.com/sheremetyev/sandfence)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T14:53:04Z

## Translation

タイトル: Show HN: クロードおよびコーデックス用の最小限のネイティブ macOS サンドボックス
記事のタイトル: GitHub - sheremetyev/sandfence: クロード コードおよびコーデックス用の最小限のネイティブ macOS サンドボックス · GitHub
説明: クロード コードおよびコーデックス用の最小限のネイティブ macOS サンドボックス - sheremetyev/sandfence

記事本文:
GitHub - sheremetyev/sandfence: クロード コードおよびコーデックス用の最小限のネイティブ macOS サンドボックス · GitHub
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
シェレメーチエフ
/
砂柵
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラン

hes タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
13 コミット 13 コミット DESIGN.md DESIGN.md ライセンス ライセンス README.md README.md Sandfence.sh Sandfence.sh test.sh test.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェント (Claude Code または Codex) を独自のリポジトリで実行します。
「スキップ権限」モードでは、エージェントではなく macOS サンドボックスがその内容を強制します。
触れることができる。間違った rm -rf 、野良 git replace --hard 、pip のインストール
システム: サンドボックスはこれらをインシデントからエラーに変換します。
s claude # 現在のリポジトリに限定してクロード コードを実行します (以下の設定)
これは、macOS の Sandbox-exec (シートベルト) に関する監査可能な 1 つの短いシェル スクリプトです。
サンドボックス)。エージェントは作業コピーへの読み取り/書き込みアクセス権を取得し、必要最低限のアクセス権を取得します。
それ自体とテストを実行します。それ以外のすべてはデフォルトで拒否され、次の場合にのみ開かれます。
尋ねてください。
ネイティブ、インプレース — エージェントは実際の macOS を使用して実際の作業コピーを編集します
ツールチェーン。 Linux ゲストもコピーされたファイルも、変更を元に同期することもありません。
軽量 — マシンではなくサンドボックス化されたプロセスです。起動する VM もイメージも必要ありません。
デーモン。 1 つのスクリプトを超えてインストールするものはありません。
OS による強制 — エージェントではなくカーネルが、何にアクセスできるかを決定するため、
スキップ権限モードでは無人で実行されます。
監査可能 — 信頼する前に 1 つのスクリプトを読んでください。 --print は正確に何を示しますか
呼び出し許可、およびその仕組みについては DESIGN.md で説明しています。
スコープ内: エージェントは間違ったコマンドを実行します (rm -rf を間違った場所で実行します)。
git replace --hard 、タスク外のファイルを破壊し、システム全体にジャンクをインストールします。
Sandfence はこれらをインシデントからエラーに変換します。
範囲外: 悪意のあるエージェント、プロンプト インジェクション、または汚染された依存関係
積極的に逃走または脱出しようとします。 Sandbox-exec はカーネルとユーザーを共有します
アカウント — それはガードレールであり、

封じ込めの境界線。ネットワークはオープンであり、
作業コピーは読み取り可能なため、エージェントが実行するコード ( npm install 、ビルド フック)
リポジトリ内のシークレットを読み取り、送信します。信頼できないコードの場合は、VM または別のを使用します。
ユーザーアカウント。
Apple Silicon 上の macOS が必要です。それをクローンして、短い s ラッパーを毎日作成します
使用するのは s clude だけです:
# クローンを作成して読み取ります。それはセキュリティツールです
git clone https://github.com/sheremetyev/sandfence ~ /.config/sandfence
# 日常的に使用するツールチェーン (および追加のパス) を含む短い `s` ラッパー
cat > ~ /.local/bin/s << ' EOF '
#!/bin/sh
exec ~/.config/sandfence/sandfence.sh --rust --node --python "$@"
終了後
chmod +x ~ /.local/bin/s
調整はあなた次第です。使用するプリセットのみを保持し、 -r DIR / -w DIR を追加します。
よくたどり着く道。各プリセットと許可は、サンドボックスを実際に拡張するものです。
認証 (1 回)。各エージェントは、サンドボックスが付与するファイルに独自のトークンを保持します。
ログインキーチェーン、決してシェル環境ではありません。
クロード コード — s claude を実行し、次に /login を実行します。印刷された URL を
ブラウザ (サンドボックスではブラウザを開くことができません)。 ~/.claude/.credentials.json を書き込みます。
それ以降はリフレッシュされます。
Codex — codex ログインを 1 回 (どこでも) 実行します。そのトークンは ~/.codex/auth.json にあります。
サンドボックスはそれを読み取ります。
読み書き可能
現在のディレクトリ - 作業コピー
読み取り専用
独自の .git / .jj — サンドボックスの外側でバージョン管理を推進します。エージェントは履歴をコミットまたは書き換えることができません
拒否されました
$HOME の残り — ~/.ssh 、 ~/.aws 、gh / glab トークン、ログイン キーチェーン、~/.gitconfig 資格情報、その他のリポジトリ
明示的に拡張します: -r PATH / -w PATH ディレクトリまたはファイルを追加し、
--rust / --node / --python プリセットは、ビルド キャッシュを保持しながら読み取り/書き込みを許可します。
レジストリ トークンと PATH プラント ベクターは拒否されました。 --print は、その内容を正確に示します。
呼び出し

n 助成金。 DESIGN.md では、それぞれの助成金が存在する理由を説明しています。
Apple Silicon 上の macOS、デフォルトのツールチェーン (rustup、nvm + Stock npm、Apple
Python3)。 Homebrew、pyenv、pnpm、yarn は自動検出されません。 -r / -w を使用して許可します。
Sandbox-exec は Apple (2017) によって非推奨になりましたが、完全に機能します (CLI のみ)
非推奨です。その下のシートベルト エンジンは依然として macOS アプリ サンドボックスに電力を供給しており、
Chrome のレンダラ サンドボックスなので、なくなることはありません。将来の macOS では CLI が変更される可能性があります。
/tmp と /var/folders は読み取り/書き込みのホールセールです。そこでリポジトリが分離されることは期待できません。
jj では、 .jj は読み取り専用です。読み取り専用コマンドを jj --ignore-working-copy … として実行します。
./test.sh ~/sandfence-tests はサンドボックス内で実際のコマンドを実行し、それぞれをアサートします
許可/拒否 — つまり、単に読み取るだけでなく、何が許可されているかを確認します。実際のディレクトリ ( /tmp ではない) を使用します。
そしてそれをプレーンシェルで実行します（sandbox-execはネストできません）。
Claude Code および Codex 用の最小限のネイティブ macOS サンドボックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Minimal native macOS sandbox for Claude Code and Codex - sheremetyev/sandfence

GitHub - sheremetyev/sandfence: Minimal native macOS sandbox for Claude Code and Codex · GitHub
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
sheremetyev
/
sandfence
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
13 Commits 13 Commits DESIGN.md DESIGN.md LICENSE LICENSE README.md README.md sandfence.sh sandfence.sh test.sh test.sh View all files Repository files navigation
Run a coding agent — Claude Code or Codex — on a repo in its own
"skip-permissions" mode, while the macOS sandbox , not the agent, enforces what it
can touch. A wrong rm -rf , a stray git reset --hard , a pip install into your
system: the sandbox turns these from incidents into errors.
s claude # run Claude Code, confined to the current repo (setup below)
It's one short, auditable shell script around macOS sandbox-exec (the Seatbelt
sandbox). The agent gets read-write access to your working copy and the bare minimum to
run itself and your tests; everything else is denied by default and opened only when you
ask.
Native, in place — the agent edits your real working copy with your real macOS
toolchain. No Linux guest, no copied files, no syncing changes back.
Lightweight — a sandboxed process, not a machine: no VM to boot, no image, no
daemon; nothing to install beyond one script.
OS-enforced — the kernel, not the agent, decides what it can touch, so you can let
it run unattended in skip-permissions mode.
Auditable — read the one script before you trust it; --print shows exactly what
any invocation grants, and DESIGN.md explains how it works.
In scope: the agent runs the wrong command — rm -rf in the wrong place,
git reset --hard , clobbering files outside the task, installing junk system-wide.
sandfence turns these from incidents into errors.
Out of scope: a malicious agent, prompt injection, or a poisoned dependency
actively trying to escape or exfiltrate. sandbox-exec shares your kernel and user
account — it's a guardrail, not a containment boundary. The network is open and your
working copy is readable, so code the agent runs ( npm install , a build hook) can
read secrets in your repo and send them out. For untrusted code, use a VM or a separate
user account.
Requires macOS on Apple Silicon . Clone it, then make a short s wrapper so daily
use is just s claude :
# Clone — and read it; it's a security tool
git clone https://github.com/sheremetyev/sandfence ~ /.config/sandfence
# A short `s` wrapper with the toolchains (and extra paths) you use day to day
cat > ~ /.local/bin/s << ' EOF '
#!/bin/sh
exec ~/.config/sandfence/sandfence.sh --rust --node --python "$@"
EOF
chmod +x ~ /.local/bin/s
s is yours to tune: keep only the presets you use, and add -r DIR / -w DIR for
paths you reach for often. Each preset and grant is a real widening of the sandbox.
Auth (once). Each agent keeps its own token in a file the sandbox grants — never the
login Keychain, never your shell environment.
Claude Code — run s claude , then /login ; paste the printed URL into your
browser (the sandbox can't open one). It writes ~/.claude/.credentials.json and
refreshes it from then on.
Codex — run codex login once (anywhere); its token lives in ~/.codex/auth.json ,
which the sandbox reads.
Read-write
the current directory — your working copy
Read-only
its own .git / .jj — you drive version control outside the sandbox; the agent can't commit or rewrite history
Denied
the rest of $HOME — ~/.ssh , ~/.aws , gh / glab tokens, the login Keychain, ~/.gitconfig credentials, other repos
Widen it explicitly: -r PATH / -w PATH add a directory or file, and the
--rust / --node / --python presets grant build caches read-write while keeping
registry tokens and PATH-plant vectors denied. --print shows exactly what an
invocation grants; DESIGN.md explains why each grant is there.
macOS on Apple Silicon , default toolchains ( rustup , nvm + stock npm , Apple
python3 ). Homebrew, pyenv, pnpm, and yarn aren't auto-detected — grant them with -r / -w .
sandbox-exec is deprecated by Apple (2017) but fully functional — only the CLI
is deprecated; the Seatbelt engine under it still powers the macOS App Sandbox and
Chrome's renderer sandbox, so it isn't going away. A future macOS could change the CLI.
/tmp and /var/folders are read-write wholesale — don't expect repo isolation there.
Under jj , .jj is read-only; run read-only commands as jj --ignore-working-copy … .
./test.sh ~/sandfence-tests runs real commands inside the sandbox and asserts each
allow/deny — so you verify what's granted, not just read it. Use a real dir (not /tmp ),
and run it in a plain shell ( sandbox-exec can't nest).
Minimal native macOS sandbox for Claude Code and Codex
There was an error while loading. Please reload this page .
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
