---
source: "https://github.com/1ay1/agentty"
hn_url: "https://news.ycombinator.com/item?id=48928002"
title: "Agentty – A drop-in alternative to claude-code, written in C++26. 11.0 MB binary"
article_title: "GitHub - 1ay1/agentty: AI pair programming in your terminal — one static binary, sub-ms startup, any model · GitHub"
author: "tfeayush"
captured_at: "2026-07-15T22:49:35Z"
capture_tool: "hn-digest"
hn_id: 48928002
score: 1
comments: 0
posted_at: "2026-07-15T22:30:24Z"
tags:
  - hacker-news
  - translated
---

# Agentty – A drop-in alternative to claude-code, written in C++26. 11.0 MB binary

- HN: [48928002](https://news.ycombinator.com/item?id=48928002)
- Source: [github.com](https://github.com/1ay1/agentty)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T22:30:24Z

## Translation

タイトル: Agentty – C++26 で書かれた、クロード コードのドロップイン代替品。 11.0MBバイナリ
記事のタイトル: GitHub - 1ay1/agentty: 端末での AI ペア プログラミング — 1 つの静的バイナリ、サブミリ秒の起動、任意のモデル · GitHub
説明: 端末での AI ペア プログラミング - 1 つの静的バイナリ、サブミリ秒の起動、任意のモデル - 1ay1/agentty

記事本文:
GitHub - 1ay1/agentty: ターミナルでの AI ペア プログラミング — 1 つの静的バイナリ、サブミリ秒の起動、任意のモデル · GitHub
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
1ay1
/
代理店
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コ

デ
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,003 コミット 1,003 コミット .github/ workflows .github/ workflows acp-cpp @ d8b8008 acp-cpp @ d8b8008 docs docs include/agentty include/agentty Maya @ 0452aef Maya @ 0452aef mcp-cpp @ f90f71c mcp-cpp @ f90f71c パッケージ化 パッケージング スクリプト スクリプト src src テスト テスト .gitignore .gitignore .gitmodules .gitmodules CHANGELOG.md CHANGELOG.md CMakeLists.txt CMakeLists.txt ライセンス ライセンス README.md README.md Agentty.gif Agentty.gif install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
端末内でAIペアプログラミング
1 つの静的バイナリ。ミリ秒未満の起動。どのモデルでも。
カール -fsSL https://raw.githubusercontent.com/1ay1/agentty/master/install.sh |しー
プロジェクトの cd
代理店
最初の起動で認証が開きます — OAuth (Claude Pro/Max サブスクリプションを使用) または API キーを貼り付けます。参加すると、最初のウェルカム カードで試してみるべきことがいくつか提案されます。と入力して Enter キーを押すだけです。
コールドスタートは1ms未満。ノードもPythonもnpmインストールもありません。ただの静的バイナリです。
Claude、GPT、Groq、OpenRouter、Ollama、または任意の OpenAI 互換エンドポイント。 ^P でライブに切り替えます。
すべてのシェル呼び出しは、bwrap (Linux) / Sandbox-exec (macOS) 内で実行されます。ファイル ツールは、ワークスペース外のパスを拒否します。
インターネットのないボックス上で実行します。ラップトップはエンドツーエンドで固定された TLS を使用して SSH 経由でバイトを中継します。
読み取り、書き込み、編集、bash、grep、glob、git、web、search_docs、task — それぞれに専用のウィジェットが付いています。
エージェント スキル + 記憶を覚える/忘れる。一度教えれば、毎回のセッションであなたの習慣がわかります。
エージェント # クロード (デフォルト)
Agentty --provider openai -m gpt-4o # GPT
Agentty --provider groq -m llama-3.3-70b # Groq
Agentty --provider ollama -m qwen2.5-coder # ローカル モデル
Agentty --provider openrouter # OpenRouter 経由の任意のモデル
--pr

オーバーバイダーは持続します。 ^P を使用してアプリ内ライブを切り替えます。
キー
アクション
キー
アクション
入力してください
送信
^K
コマンドパレット
Esc
キャンセル/拒否
^J
スレッド一覧
Sタブ
サイクルプロファイル
^P
モデルピッカー
Alt+Enter
改行
^N
新しいスレッド
^G
コードブロックを実行する
^←/→ または Alt+←/→
スレッドを循環させる
もっと見る
# Debian / Ubuntu
カール -fsSLO https://github.com/1ay1/agentty/releases/latest/download/agentty_amd64.deb
sudo dpkg -i Agentty_amd64.deb
# Fedora / RHEL / CentOS
sudo dnf インストール https://github.com/1ay1/agentty/releases/latest/download/agentty-x86_64.rpm
#openSUSE
sudo zypper インストール https://github.com/1ay1/agentty/releases/latest/download/agentty-x86_64.rpm
# アーチ (AUR)
やあ -S エージェント ビン
#アルパイン
カール -fsSLO https://github.com/1ay1/agentty/releases/latest/download/agentty-x86_64.apk
sudo apk add --allow-untrusted Agentty-x86_64.apk
macOS
brew Tap 1ay1/tap && brew install Agentty
窓
スクープバケット追加 1ay1 https://github.com/1ay1/スクープ - バケット;スクープインストールエージェント
# または
winget インストールagentty.agentty
どこでも (パッケージマネージャーなし)
カール -fsSL https://raw.githubusercontent.com/1ay1/agentty/master/install.sh |しー
ソースから (C++26 ツールチェーンが必要 — GCC 14+ / 最近の Clang / MSVC)
git clone --recursive git@github.com:1ay1/agentty.git
cd Agentty && cmake -B build && cmake --build build -j
すべてのバイナリは単一の完全に静的な実行可能ファイルです (Linux では x86_64 + aarch64、macOS では Intel + Apple Silicon)。パッケージの詳細: パッケージ/README.md 。
Agentty airgap --setup user@host # 初回: 資格情報をコピーします
Agentty airgap user@host # 以降は毎回
ラップトップは SOCKS5-over-SSH 経由で中継します。実際のアップストリーム上の TLS ピン - 間にあるネットワークは MITM できません。
エージェントは、エージェント クライアント プロトコル (Zed がクロード コードに使用するものと同じプロトコル) を話します。 Zed の設定に追加します。
{
"エージェントサーバー": {
"エージェント" : {
「コム」

mand" : " エージェント " ,
"args" : [ "acp " ]
}
}
}
返信からコード ブロックを実行 (Ctrl+G)
AI は、囲われたコマンドのブロックを渡します。これをコピーアンドペーストしないでください。 ^G リスト
最後の返信からのブロック。 Enter (または数字) を押すと対話的に実行されます
実際の端末上: TUI が一時停止し、sudo パスワード プロンプトが機能し、出力が表示されます。
ストリームがライブの場合、Ctrl+C はコマンドを強制終了します (エージェントではありません)。終了すると、
結果カードを使用すると、キャプチャした出力をコンポーザーに添付できます。
折りたたまれたチップ ( a )、コピー ( y )、または破棄 ( Esc ) — したがって、「失敗しました」
X" は、何も再入力しなくてもモデルに到達します。
すべての OS でブロックごとに適切なシェルを実行します: sh / bash はブロックを介して実行します。
Linux/macOS の /bin/sh、powershell / pwsh および cmd / but はブロックされます。
Windows の PowerShell / cmd.exe。プロンプトの $ マーカーが削除され、ブロックになります
プラットフォームでは代わりに編集/コピーを実行できず、キャプチャの上限は次のとおりです。
2MB。詳細: docs/RUN_CODE_BLOCK.md
SKILL.md を .agentty/skills/ または ~/.agentty/skills/ の下のどこかにドロップします。これは次のターンに有効になります。 Claude Code の .claude/skills/ 形式と互換性があります。
内部 DSL または部族規約を備えたコードベースでは、厳選されたスキルによりエージェントの精度が最大 20% から最大 85% に跳ね上がります ( 研究 )。
純粋関数的な更新ループ: (Model, Msg) -> (Model, Cmd) 。ビューは Model -> Element で、Maya によってレンダリングされます。 posix_spawn +poll(2) によるプロセス管理。ファイルの書き込みはアトミックです ( write + fsync + rename )。
詳細: docs/ARCHITECTURE.md · docs/RENDERING.md
リリースをカットするのは 1 つのコマンドです。
scripts/cut-release.sh X.Y.Z # POSIX / macOS / Linux / Git-Bash
スクリプト \c ut-release.cmd X.Y.Z # Windows cmd.exe
CMakeLists.txt (単一ソース) のプロジェクト (agentty VERSION …) をバンプします。
すべてのマニフェストが派生する真実の)、 CHANGELOG.md の [未リリース] を促進します
[X.Y.Z] へのセクション、コミット、タグ vX.Y.Z 、およびプッシュ

s.タグプッシュが発火する
GitHub Actions、すべてのバイナリ + OS パッケージ (Linux x86_64/aarch64) をビルドします。
ネイティブ ランナー、macOS Intel/ARM、Windows .exe / .msi 上) に自動送信されます。
winget、Homebrew、Scoop、および AUR — nix/snap/gentoo マニフェストが添付されています
リリースまで。 --何も書かずにプレビューをドライランします。
端末での AI ペア プログラミング - 1 つの静的バイナリ、サブミリ秒の起動、任意のモデル
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
10
v0.2.7
最新の
2026 年 7 月 12 日
+ 9 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI pair programming in your terminal — one static binary, sub-ms startup, any model - 1ay1/agentty

GitHub - 1ay1/agentty: AI pair programming in your terminal — one static binary, sub-ms startup, any model · GitHub
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
1ay1
/
agentty
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
1,003 Commits 1,003 Commits .github/ workflows .github/ workflows acp-cpp @ d8b8008 acp-cpp @ d8b8008 docs docs include/ agentty include/ agentty maya @ 0452aef maya @ 0452aef mcp-cpp @ f90f71c mcp-cpp @ f90f71c packaging packaging scripts scripts src src tests tests .gitignore .gitignore .gitmodules .gitmodules CHANGELOG.md CHANGELOG.md CMakeLists.txt CMakeLists.txt LICENSE LICENSE README.md README.md agentty.gif agentty.gif install.sh install.sh View all files Repository files navigation
AI pair programming in your terminal
One static binary. Sub-millisecond startup. Any model.
curl -fsSL https://raw.githubusercontent.com/1ay1/agentty/master/install.sh | sh
cd your-project
agentty
First launch opens auth — OAuth (uses your Claude Pro/Max subscription) or paste an API key. Once you're in, a first-run welcome card suggests a few things to try; just type and hit Enter.
Cold start under 1ms. No Node, no Python, no npm install. Just a static binary.
Claude, GPT, Groq, OpenRouter, Ollama, or any OpenAI-compatible endpoint. Switch live with ^P .
Every shell call runs inside bwrap (Linux) / sandbox-exec (macOS). File tools refuse paths outside your workspace.
Run on a box with no internet. Your laptop relays the bytes over SSH with TLS pinned end-to-end.
read · write · edit · bash · grep · glob · git · web · search_docs · task — each with a purpose-built widget.
Agent Skills + remember/forget memory. Teach it once, every session knows your conventions.
agentty # Claude (default)
agentty --provider openai -m gpt-4o # GPT
agentty --provider groq -m llama-3.3-70b # Groq
agentty --provider ollama -m qwen2.5-coder # Local model
agentty --provider openrouter # Any model via OpenRouter
--provider persists. Switch live in-app with ^P .
Key
Action
Key
Action
Enter
Send
^K
Command palette
Esc
Cancel / reject
^J
Thread list
S-Tab
Cycle profile
^P
Model picker
Alt+Enter
Newline
^N
New thread
^G
Run code block
^←/→ or Alt+←/→
Cycle threads
More
# Debian / Ubuntu
curl -fsSLO https://github.com/1ay1/agentty/releases/latest/download/agentty_amd64.deb
sudo dpkg -i agentty_amd64.deb
# Fedora / RHEL / CentOS
sudo dnf install https://github.com/1ay1/agentty/releases/latest/download/agentty-x86_64.rpm
# openSUSE
sudo zypper install https://github.com/1ay1/agentty/releases/latest/download/agentty-x86_64.rpm
# Arch (AUR)
yay -S agentty-bin
# Alpine
curl -fsSLO https://github.com/1ay1/agentty/releases/latest/download/agentty-x86_64.apk
sudo apk add --allow-untrusted agentty-x86_64.apk
macOS
brew tap 1ay1/tap && brew install agentty
Windows
scoop bucket add 1ay1 https: // github.com / 1ay1 / scoop - bucket; scoop install agentty
# or
winget install agentty.agentty
Anywhere (no package manager)
curl -fsSL https://raw.githubusercontent.com/1ay1/agentty/master/install.sh | sh
From source (needs a C++26 toolchain — GCC 14+ / recent Clang / MSVC)
git clone --recursive git@github.com:1ay1/agentty.git
cd agentty && cmake -B build && cmake --build build -j
All binaries are a single fully-static executable (x86_64 + aarch64 on Linux, Intel + Apple Silicon on macOS). Packaging details: packaging/README.md .
agentty airgap --setup user@host # first time: copies credentials
agentty airgap user@host # every time after
Your laptop relays via SOCKS5-over-SSH. TLS pins on real upstreams — the network in between can't MITM you.
agentty speaks the Agent Client Protocol — the same protocol Zed uses for Claude Code. Add to Zed's settings:
{
"agent_servers" : {
"agentty" : {
"command" : " agentty " ,
"args" : [ " acp " ]
}
}
}
Run code blocks from replies (Ctrl+G)
The AI hands you a fenced block of commands — don't copy-paste it. ^G lists
the blocks from the last reply; Enter (or a digit) runs one interactively
on your real terminal : the TUI suspends, sudo password prompts work, output
streams live, Ctrl+C kills the command (not agentty). When it exits, a
result card lets you attach the captured output to the composer as a
collapsed chip ( a ), copy it ( y ), or discard ( Esc ) — so "it failed with
X" reaches the model without you re-typing anything.
Runs the right shell per block on every OS: sh / bash blocks through
/bin/sh on Linux/macOS, powershell / pwsh and cmd / bat blocks through
PowerShell / cmd.exe on Windows. Prompt $ markers are stripped, a block
your platform can't run offers edit/copy instead, and capture is capped at
2 MB. Details: docs/RUN_CODE_BLOCK.md
Drop a SKILL.md anywhere under .agentty/skills/ or ~/.agentty/skills/ — it's live next turn. Compatible with Claude Code's .claude/skills/ format.
On codebases with internal DSLs or tribal conventions, agent accuracy jumps from ~20% to ~85% with curated skills ( research ).
Pure-functional update loop: (Model, Msg) -> (Model, Cmd) . View is Model -> Element , rendered by maya . Process management via posix_spawn + poll(2) . File writes are atomic ( write + fsync + rename ).
Deep dive: docs/ARCHITECTURE.md · docs/RENDERING.md
Cutting a release is one command:
scripts/cut-release.sh X.Y.Z # POSIX / macOS / Linux / Git-Bash
scripts \c ut-release.cmd X.Y.Z # Windows cmd.exe
It bumps project(agentty VERSION …) in CMakeLists.txt (the single source
of truth every manifest derives from), promotes CHANGELOG.md 's [Unreleased]
section to [X.Y.Z] , commits, tags vX.Y.Z , and pushes. The tag push fires
GitHub Actions, which builds every binary + OS package (Linux x86_64/aarch64
on native runners, macOS Intel/ARM, Windows .exe / .msi ) and auto-submits to
winget, Homebrew, Scoop, and the AUR — nix/snap/gentoo manifests are attached
to the release. --dry-run previews without writing anything.
AI pair programming in your terminal — one static binary, sub-ms startup, any model
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
10
v0.2.7
Latest
Jul 12, 2026
+ 9 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
