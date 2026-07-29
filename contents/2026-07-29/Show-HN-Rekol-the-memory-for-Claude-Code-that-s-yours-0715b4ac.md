---
source: "https://github.com/rekol-io/rekol"
hn_url: "https://news.ycombinator.com/item?id=49099924"
title: "Show HN: Rekol, the memory for Claude Code that's yours"
article_title: "GitHub - rekol-io/rekol: Local-first memory for Claude Code — layered markdown you own, with on-device semantic search over your notes and past sessions. No API key, no cloud. · GitHub"
author: "leonkatz"
captured_at: "2026-07-29T17:04:36Z"
capture_tool: "hn-digest"
hn_id: 49099924
score: 1
comments: 2
posted_at: "2026-07-29T16:49:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Rekol, the memory for Claude Code that's yours

- HN: [49099924](https://news.ycombinator.com/item?id=49099924)
- Source: [github.com](https://github.com/rekol-io/rekol)
- Score: 1
- Comments: 2
- Posted: 2026-07-29T16:49:49Z

## Translation

タイトル: 表示 HN: Rekol、クロード・コードの思い出をあなたのものに
記事のタイトル: GitHub - rekol-io/rekol: クロード コードのローカル ファースト メモリ — 所有する階層化されたマークダウン。メモや過去のセッションに対するオンデバイスのセマンティック検索が可能です。 API キーもクラウドもありません。 · GitHub
説明: クロード コードのローカル ファースト メモリ — メモや過去のセッションに対するオンデバイスのセマンティック検索を備えた、あなたが所有する階層化されたマークダウン。 API キーもクラウドもありません。 - rekol-io/rekol

記事本文:
GitHub - rekol-io/rekol: クロード コードのローカル ファースト メモリ — メモや過去のセッションに対するオンデバイスのセマンティック検索を備えた、あなたが所有する階層化されたマークダウン。 API キーもクラウドもありません。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
がありました

ロード中にエラーが発生しました。このページをリロードしてください。
レコルイオ
/
レコル
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
329 コミット 329 コミット .github .github bin bin docs docs フック フック スクリプト スクリプト スキル スキル src/ rekol src/ rekol テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス READ-ME-CLAUDE.md READ-ME-CLAUDE.md README.md README.md SECURITY.md SECURITY.md install.sh install.sh pyproject.toml pyproject.toml要件.txt 要件.txt uninstall.sh uninstall.sh uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すでに使用している AI アシスタント用のローカルファースト メモリ。ドロップインメモリー
レイヤー — API キーなし、マシン上で実行され、アシスタントがそれを使用します
自動的に。あなたの記憶はあなたが所有するフォルダーにマークダウンされます。
REKOL は、端末、IDE (VS Code または
JetBrains)、または Claude デスクトップ アプリ。 macOS と Linux 上で動作します。
あなたはメモリ コマンドを実行しません。あなたはただ作業するだけで、アシスタントはすでにメモリ コマンドを実行しています。
コンテキストを作成し、それを使用します。 REKOLのこだわりコーナー：
API キーはなく、完全にローカルです。埋め込み (BAAI bge-small ) とベクトル検索
( sqlite-vec ) をマシン上で実行します。アカウントもキーもテレメトリもありません。
メモリ層は完全にハードウェア上で実行されます。
新しいアプリではなく、ドロップイン レイヤー。 REKOL は Claude Code にプラグインします。
すでに使用しているアシスタント — 採用するエージェントも、切り替えるツールもありません。
自ら表面化する記憶。階層化モデル (always/when/topics/Knowledge) inj

各セッションの開始時に適切なコンテキストが反映され、
アシスタントは機能するにつれてさらに多くのことを引き出します - 指示されるのではなく、ただ知っているだけです
見ること。
あなたの記憶はあなたが所有するプレーンなマークダウン (Obsidian、grep、git) と REKOL です。
過去のセッションの記録と厳選されたメモを検索します - これら
見出しではなく、テーブルの賭けがうまくいったかどうか。
macOS または Linux。 Claude Code がインストールされています (REKOL はそのメモリ層です)。
Python ≥3.11 (その sqlite3 にenable_load_extension がある) - に必要
sqlite-vec ベクトル検索。 macOS のシステム Python と python.org
インストーラーはこれを無効にして出荷します。これにより、検索がキーワード/numpy に低下します。
フォールバック。信頼できるオプション (install.sh はこれらを自動検出して優先します):
brew install uv && uv python install 3.12 — uv の Python には常に
拡張子。インストーラーが自動的にそれを取得します。
または brew install python — Homebrew のデフォルトの python@3 は拡張機能を使用して構築されています。
jq (オプション) — ~/.claude/settings.json フックの自動配線のみ。
これがないと、インストーラは手動でマージするスニペットを出力します。
適切なインタープリターが見つからない場合、install.sh は正確な内容で早期に停止するようになりました。
劣化したセットアップをインストールするのではなく、修正してください。
git clone https://github.com/rekol-io/rekol && cd rekol && ./install.sh
これがセットアップ全体です。rekol CLI をインストールし、それを Claude Code に接続し、
既存のクロード コード履歴にインデックスを付けます。そのため、アシスタントは周囲の記憶を保持し、
過去のセッションはすぐに検索できます。 (メモリフォルダーのプロンプトに応答するか、事前に設定された
REKOL_HOME — 以下のオプションを参照してください。)
2. プロジェクトを教えてください (推奨) — クロード コードで、「rekol メモリをセットアップしてください」と言います。
いつでも/いつでも/話題の思い出を歴史から抽出し、見直すことができます —
あなたは保存されたものを承認します。 REKOL はこれなしでもすでに動作します (インストールすると履歴がインデックス化されます)

y);それを実行します
いつでも。
メモリ フォルダーを最初に選択します (プロンプトはスキップします)。
import REKOL_HOME= " $HOME /rekol-memory " # 所有する任意のフォルダー
Dropbox/iCloud/git/Syncthing 経由で同期するか、ローカルに保持します。インデックスはローカル キャッシュに保存されます。
メモリ フォルダの外にあるため、メモリを同期してもインデックスは同期されません。
耐久性のあるトランスクリプト アーカイブ — rekol はセッションのローカル コピーを保存します。
クロード・コードがそのトランスクリプトをローテーションしたとしても、このファイルは存続します。 ~/.local/share/rekol/archive の下にあります。
アップロードされることはなく、デフォルトでは同期から除外されます。 --no-archive でオフにします (または
archive_enabled: rekol.config.yaml の false ); --archive-dir を使用して再配置します。除外する
exclude_paths / .rekolignore を使用した機密性の高いプロジェクト。
./install.sh フラグ (すべてオプション):
--dry-run — すべてのアクションを実行せずに出力します。
--no-hook — クロード コードの SessionStart/フックの配線 (settings.json) をスキップします。
--no-skill — rekol/memory Claude Code スキルのインストールをスキップします。
--no-shellrc — shell-rc の編集をスキップします (PATH + REKOL_HOME エクスポート) — に書き込まれます
ログインシェルごとに、zsh の場合は ~/.zshrc、bash の場合は ~/.bashrc / ~/.bash_profile。
--test-mode — --no-hook --no-skill --no-shellrc の短縮形。
--tools-home PATH — venv + tools home (デフォルト ~/.local/share/rekol ) をオーバーライドします。
--bin-dir PATH — rekol shim が存在する場所 (デフォルト ~/bin ) をオーバーライドします。
--merge — 従来の ~/.claude/projects/*/memory/ コンテンツのインポートをオプトインします。
--no-archive — 永続的なトランスクリプト アーカイブを無効にします (seeds archive_enabled: false )。
--archive-dir PATH — 永続的なアーカイブの場所を設定します (デフォルトは ~/.local/share/rekol/archive )。
--help — 使用方法を出力して終了します。
既存のクロード コードのトランスクリプトをローカルにバックフィルしてインストールします
検索可能なインデックスにより、rekol 検索が過去の作業に対して即座に機能します

イーリー。ビヨンド
それは:
最新の状態を維持します。SessionEnd フックは作業中にインデックスを再作成します。走る
rekol session-index --incremental はいつでも再同期を強制します (または --full を実行します)
すべてをやり直します）。
Notes/docs フォルダーをインポートする — rekol import ~/Documents/ObsidianVault
テキスト ファイルのツリーを検索可能なコンテンツに変換します。これは機械式です
ドキュメントを見つけやすくする変換。アシスタント (または rekol Capture を使用するあなた) がそれらを always / when / topic レイヤーにキュレーションします。
耐久性のあるメモリを蒸留する — クロード コードに「rekol メモリをセットアップしてください」と指示します。
レビューのために履歴から常に / いつ / トピックのエントリを提案します
(オプトイン; 着地を承認することになります)。
検証 — rekol で「あなたが書いたもの」を検索します。
2 つのフェーズがあり、正直に言って違います。
1 日目 — 検索可能な履歴 (リコール)。既存のクロードのインデックスをインストールします
セッションをローカルの検索可能なストアにコード化するため、インストール直後に次のことが可能になります。
あなたが取り組んでいることについて尋ねると、クロード・コードがそれを見つけます。すべてはあなたの責任です
ディスク — トランスクリプトはアップロードされません。 (アシスタント主導の「rekol をセットアップする」
「memory」インタビューはこれを補うものであり、以下の厳選された層の由来でもあります。）
時間が経つにつれて、プロジェクトがどのように考えるか (理解するか) が学習されます。見返りは
周囲の記憶: 繰り返される「いつも X をする」、「リポジトリは Y にある」、「Z を選択した」
それは、常に / いつ / トピックのレイヤーとサーフェスにキュレーションされます。
あなたが尋ねなくても、自分のものです。その一部は、作業したりキャプチャしたりすることで蓄積されます。の
rest は明示的に実行したメモリ ブートストラップ ステップから取得されます。これは再度読み込まれます。
インデックス付きトランスクリプトを作成し、レビュー用に永続的な思い出を提案します（あなた
土地を承認します)。
正直さの表れ: そのブートストラップは、あなた自身のクロード コードを読み取ることです。
トランスクリプト コーパス — バンドルされたモデルはありません。長い歴史の上でそれを実行することは、
実数

オーケンはあなたのアカウントに対して支出し、履歴の量に比例します
餌を与えます。まさにその理由から、オプトインおよびレビューゲートとなっています。次の場合にスコープを指定して開始します
あなたのコーパスは大きいです。
サブコマンドを含む単一の rekol コマンド:
rekol search "query" [--top N] [--json] — セマンティック + キーワード検索。
rekol インデックスの再構築 | update — ベクトルインデックスを(再)構築します。
rekol Capture — 新しいメモリを追加します。
rekol import <dir> — 既存のメモ/ドキュメント ツリーを検索にインポートします。
メモリは $REKOL_HOME の下に存在します ($MEMORY_HOME はフォールバックとして受け入れられます)。
always/ — 永続的なファクト、常にロードされます。
topic/ — 正規ソース レジストリ。
知識/ — 長い形式の永続的なレッスン。
値下げは真実の源です。 SQLite ベクトル インデックスは使い捨てであり、
再構築可能で、$REKOL_HOME の外のマシンのローカル キャッシュに存在します。
( ${XDG_CACHE_HOME:-~/.cache}/rekol/<id> ) そのため、派生したものは何も保存されません。
メモリーフォルダー。
$REKOL_HOME はあなたが所有するローカル フォルダーです。どのようにしてもマシン間で同期できます
たとえば、Dropbox、iCloud Drive、Git リモート、Syncthing、あるいはまったくそうではありません。インデックス
メモリフォルダの外のローカルキャッシュに存在するため、メモリを同期します
インデックスを同期しません。維持するツールごとの無視ファイルはありません。キャッシュ
トランスクリプトを逐語的に記録する session.db も保持します。それを保つ
同期ツリーの外にあるということは、貼り付けられたシークレットが同期を通じて漏洩することがないことを意味します。
検索では、「平均プーリング」/無数のフォールバック、または再現率が低いことについて警告されます。
venv の Python には sqlite3.enable_load_extension がありません。良い状態で再構築する
通訳:
rm -rf ~ /.local/share/rekol/.venv
brew インストール uv && uv python インストール 3.12
./install.sh
brew install python@3.12 "didn't take": これは Keg のみです (Python3 はありません)
パス)。 install.sh は python@3.12 / @3.11 opt-prefixes を直接調査するようになりました。
再実行すると見つかるはずです。それ以外の場合はUVパットを使用してください

上記の時間、または
インストールする前に、PATH="$(brew --prefix python@3.12)/libexec/bin:$PATH" をエクスポートしてください。
require-python >=3.11 を使用すると pip install でインストールが失敗します:
python3 は古すぎます (例: macOS の 3.9)。上記の UV パスを使用します。
Intel (x86_64) Mac — _ARRAY_API が見つかりません / 「Numpy は利用できません」
最初の検索: そこにあるインストール可能なトーチ ホイールは、NumPy 1.x に対して構築されています。
ピン numpy<2 を自動的に新規インストールします。すでに壊れたvenvは修正されています
~/.local/share/rekol/.venv/bin/pip 'numpy<2' をインストールします。
rekol Doctor --deep を使用してインストールを検証します。モデルのロードをチェックします。
意味のある方法で埋め込むと、そのリコールがエンドツーエンドで機能します。
rekolはあなたのものできれいに取り外せます。リポジトリから:
./uninstall.sh # 対話型 — 再構築可能なインデックスを削除する前に尋ねます
./uninstall.sh --dry-run # 何も触れずにすべての変更をプレビューします
./uninstall.sh --yes # 非対話型 (インデックスを保持)
アンインストーラーは、インストーラーが設定したもの (~/bin/rekol shim、
rekol / メモリスキル、~/.local/share/rekol venv、すべての rekol フック
~/.claude/settings.json (プラス env.REKOL_HOME )、および rekol PATH + env
シェルの行をエクスポートします rc ( ~/.zshrc 、または ~/.bashrc / ~/.bash_profile )
バッシュ)。 settings.json とシェル rc をタイムスタンプ付きの .bak ファイルにバックアップします。
編集する前に実行され、べき等です (安全に編集できます)。

[切り捨てられた]

## Original Extract

Local-first memory for Claude Code — layered markdown you own, with on-device semantic search over your notes and past sessions. No API key, no cloud. - rekol-io/rekol

GitHub - rekol-io/rekol: Local-first memory for Claude Code — layered markdown you own, with on-device semantic search over your notes and past sessions. No API key, no cloud. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
rekol-io
/
rekol
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
329 Commits 329 Commits .github .github bin bin docs docs hooks hooks scripts scripts skill skill src/ rekol src/ rekol tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE READ-ME-CLAUDE.md READ-ME-CLAUDE.md README.md README.md SECURITY.md SECURITY.md install.sh install.sh pyproject.toml pyproject.toml requirements.txt requirements.txt uninstall.sh uninstall.sh uv.lock uv.lock View all files Repository files navigation
Local-first memory for the AI assistant you already use. A drop-in memory
layer — no API key, runs on your machine, and your assistant uses it
automatically. Your memory is markdown in a folder you own.
REKOL works anywhere Claude Code runs — the terminal, your IDE (VS Code or
JetBrains), or the Claude Desktop app. It runs on macOS and Linux.
You don't run memory commands — you just work, and the assistant already has
your context and uses it. REKOL's specific corner:
No API key, fully local. Embeddings (BAAI bge-small ) and vector search
( sqlite-vec ) run on your machine. No account, no key, no telemetry — the
memory layer runs entirely on your hardware.
A drop-in layer , not a new app. REKOL plugs into Claude Code, the
assistant you already use — no agent to adopt, no tool to switch.
Memory that surfaces itself. A layered model ( always / when / topics / knowledge ) injects the right context at the start of each session, and the
assistant pulls in more as it works — it just knows , instead of being told
to look.
Your memory is plain markdown you own (Obsidian, grep, git), and REKOL
searches your past session transcripts alongside your curated notes — these
are table stakes done well, not the headline.
macOS or Linux; Claude Code installed (REKOL is a memory layer for it).
Python ≥3.11 whose sqlite3 has enable_load_extension — required for
sqlite-vec vector search. macOS's system python and the python.org
installer ship this disabled , which degrades search to a keyword/numpy
fallback. Reliable options ( install.sh auto-detects and prefers these):
brew install uv && uv python install 3.12 — uv's Python always has the
extension; the installer picks it up automatically.
or brew install python — Homebrew's default python@3 is built with extensions.
jq (optional) — only for automatic ~/.claude/settings.json hook wiring;
without it the installer prints the snippet to merge by hand.
If no suitable interpreter is found, install.sh now stops early with the exact
fix rather than installing a degraded setup.
git clone https://github.com/rekol-io/rekol && cd rekol && ./install.sh
That's the whole setup: it installs the rekol CLI, wires it into Claude Code, and
indexes your existing Claude Code history — so your assistant has ambient memory and your
past sessions are searchable right away . (Answer the memory-folder prompt, or pre-set
REKOL_HOME — see options below.)
2. Teach it your project (recommended) — in Claude Code, say "set up my rekol memory."
It distills durable always / when / topics memories from your history for you to review —
you approve what's kept. REKOL already works without this (install indexed your history); run it
whenever.
Choose your memory folder up front (skips the prompt):
export REKOL_HOME= " $HOME /rekol-memory " # any folder you own
Sync it via Dropbox/iCloud/git/Syncthing or keep it local — the index lives in a local cache
outside your memory folder, so syncing your memory never syncs the index.
Durable transcript archive — rekol keeps a local copy of your sessions so your memory
survives even if Claude Code rotates its transcripts. It lives under ~/.local/share/rekol/archive ,
is never uploaded, and is excluded from sync by default. Turn it off with --no-archive (or
archive_enabled: false in rekol.config.yaml ); relocate it with --archive-dir ; exclude
sensitive projects with exclude_paths / .rekolignore .
./install.sh flags (all optional):
--dry-run — print every action without executing it.
--no-hook — skip the Claude Code SessionStart/hook wiring (settings.json).
--no-skill — skip installing the rekol / memory Claude Code skill.
--no-shellrc — skip the shell-rc edits (PATH + REKOL_HOME export) — written to
~/.zshrc for zsh, or ~/.bashrc / ~/.bash_profile for bash, per your login shell.
--test-mode — shorthand for --no-hook --no-skill --no-shellrc .
--tools-home PATH — override the venv + tools home (default ~/.local/share/rekol ).
--bin-dir PATH — override where the rekol shim lives (default ~/bin ).
--migrate — opt in to importing legacy ~/.claude/projects/*/memory/ content.
--no-archive — disable the durable transcript archive (seeds archive_enabled: false ).
--archive-dir PATH — set the durable archive location (default ~/.local/share/rekol/archive ).
--help — print usage and exit.
Install already backfilled your existing Claude Code transcripts into a local
searchable index, so rekol search works over your past work immediately. Beyond
that:
Keep it current — the SessionEnd hook re-indexes as you work; run
rekol session-index --incremental any time to force a re-sync (or --full to
re-walk everything).
Import a notes/docs folder — rekol import ~/Documents/ObsidianVault
converts a tree of text files into searchable content. This is a mechanical
conversion that makes your docs findable; the assistant (or you, with rekol capture ) curates them into the always / when / topics layers.
Distill durable memory — tell Claude Code "set up my rekol memory" to
propose always / when / topics entries from your history for your review
(opt-in; you approve what lands).
Verify — rekol search "something you wrote" .
There are two phases, and they're honestly different.
Day 1 — searchable history (recall). Install indexes your existing Claude
Code sessions into a local searchable store, so right after installing you can
ask about something you worked on and Claude Code finds it. It's all on your
disk — your transcripts are never uploaded. (The assistant-led "set up my rekol
memory" interview tops this up and is also where the curated layer below comes from.)
Over time — it learns how your project thinks (understanding). The payoff is
ambient memory: the recurring "always do X", "repos live in Y", "we chose Z"
that gets curated into the always / when / topics layers and surfaces on its
own, without you asking. Some of that accumulates as you work and capture; the
rest comes from a memory-bootstrap step you explicitly run — it reads back over
your indexed transcripts and proposes durable memories for your review (you
approve what lands).
Token honesty: that bootstrap is your own Claude Code reading your own
transcript corpus — there's no bundled model. Running it over a large history is
real token spend against your account, proportional to how much history you
feed it. It's opt-in and review-gated for exactly that reason; start scoped if
your corpus is big.
A single rekol command with subcommands:
rekol search "query" [--top N] [--json] — semantic + keyword search.
rekol index rebuild | update — (re)build the vector index.
rekol capture — add a new memory.
rekol import <dir> — import an existing notes/docs tree into search.
Memory lives under $REKOL_HOME ( $MEMORY_HOME is accepted as a fallback):
always/ — permanent facts, always loaded.
topics/ — canonical-source registry.
knowledge/ — long-form durable lessons.
The markdown is the source of truth. The SQLite vector index is disposable and
rebuildable, and lives in a machine-local cache outside $REKOL_HOME
( ${XDG_CACHE_HOME:-~/.cache}/rekol/<id> ), so nothing derived sits in your
memory folder.
$REKOL_HOME is a local folder you own; sync it across machines however you
like — Dropbox, iCloud Drive, a git remote, Syncthing, or not at all. The index
lives in a local cache outside your memory folder, so syncing your memory
never syncs the index — there is no per-tool ignore file to maintain. The cache
also holds sessions.db , which records your transcripts verbatim; keeping it
out of the synced tree means a pasted secret can never leak through sync.
Searches warn about "mean pooling" / a numpy fallback, or recall is poor:
your venv's Python lacks sqlite3.enable_load_extension . Rebuild on a good
interpreter:
rm -rf ~ /.local/share/rekol/.venv
brew install uv && uv python install 3.12
./install.sh
brew install python@3.12 "didn't take": it's keg-only (no python3 on
PATH). install.sh now probes python@3.12 / @3.11 opt-prefixes directly, so
re-running it should find it; otherwise use the uv path above, or
export PATH="$(brew --prefix python@3.12)/libexec/bin:$PATH" before installing.
Install fails at pip install with requires-python >=3.11 : your
python3 is too old (e.g. macOS's 3.9). Use the uv path above.
Intel (x86_64) Mac — _ARRAY_API not found / "Numpy is not available" on the
first search: the installable torch wheel there is built against NumPy 1.x.
Fresh installs pin numpy<2 automatically; an already-broken venv is fixed with
~/.local/share/rekol/.venv/bin/pip install 'numpy<2' .
Verify any install with rekol doctor --deep — it checks the model loads,
embeds meaningfully, and that recall works end-to-end.
rekol is yours to remove cleanly. From the repo:
./uninstall.sh # interactive — asks before deleting the rebuildable index
./uninstall.sh --dry-run # preview every change without touching anything
./uninstall.sh --yes # non-interactive (keeps the index)
The uninstaller reverses what the installer set up — the ~/bin/rekol shim, the
rekol / memory skills, the ~/.local/share/rekol venv, every rekol hook in
~/.claude/settings.json (plus env.REKOL_HOME ), and the rekol PATH + env
export lines in your shell rc ( ~/.zshrc , or ~/.bashrc / ~/.bash_profile for
bash). It backs up settings.json and the shell rc to timestamped .bak files
before editing them, and is idempotent (safe to

[truncated]
