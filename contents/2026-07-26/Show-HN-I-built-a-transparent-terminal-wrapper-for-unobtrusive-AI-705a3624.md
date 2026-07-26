---
source: "https://github.com/emosenkis/terminai"
hn_url: "https://news.ycombinator.com/item?id=49058883"
title: "Show HN: I built a transparent terminal wrapper for unobtrusive AI"
article_title: "GitHub - emosenkis/terminai: Make your coding AI available in your shell · GitHub"
author: "emosenkis"
captured_at: "2026-07-26T15:56:04Z"
capture_tool: "hn-digest"
hn_id: 49058883
score: 1
comments: 0
posted_at: "2026-07-26T15:05:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I built a transparent terminal wrapper for unobtrusive AI

- HN: [49058883](https://news.ycombinator.com/item?id=49058883)
- Source: [github.com](https://github.com/emosenkis/terminai)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T15:05:31Z

## Translation

タイトル: Show HN: 目立たない AI 用の透明なターミナル ラッパーを構築しました
記事のタイトル: GitHub - emosenkis/terminai: コーディング AI をシェルで利用できるようにする · GitHub
説明: コーディング AI をシェルで利用できるようにします。 GitHub でアカウントを作成して、emosenkis/terminai の開発に貢献してください。
HN テキスト: 既存の AI サブスクリプションをターミナル アシスタントとして使用して、残りの時間の邪魔にならずに、シェル出力を分析したり、複雑な awk の呪文を時折支援したりする方法が必要でした。これを行う既存のものが見つからなかったので、私が構築しました。私は出発点として pvolok の優れた mproc を使用し、実装方法について広範なガイダンスを提供しながら、最も厄介な部分を除くすべてのコーディングについては Claude と Codex に大きく依存しました。コーディング エージェントがすでにシェル コマンドを実行できることは知っていますが、それでは AI が運転席に座ることになります。テルミナイでは、私が運転手になって、コデックスかクロード（または他のエージェント）は私が助けを求めるまで後部座席に静かに座っていてください。最も難しかったのは、本当に透明にすることでした。端末を埋め込むほとんどの TUI は、端末エミュレータのネイティブ スクロールバックを中断し、クリップボードへのコピーも中断することがよくあります。私は、これらの両方が Terminai なしの場合とまったく同じように動作することを確認するために懸命に働きました。途中で、これらの動作の両方を機能させるために、ratatui をフォークしてレンダリング方法を変更する必要がありました。毎日のドライバーとして Mac と Linux でテストし、Windows で十分に動作することを確認しました。クイックインストール:
brew install emosenkis/tap/terminai

記事本文:
GitHub - emosenkis/terminai: コーディング AI をシェルで利用できるようにする · GitHub
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
エモセンキス
/
終点
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション操作

ション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
640 コミット 640 コミット .claude .claude .codex/ skill/ create-release .codex/ skill/ create-release .github/ workflows .github/ workflows config configcrossterm-0.28-facadecrossterm-0.28-facade docs docs helpers/ print-key helpers/ print-key img img npm npm plan planrat-salsa @ 3173ef6 ラットサルサ @ 3173ef6 ラタトゥイ @ 10bc72a ラタトゥイ @ 10bc72a スクープ スクープ スクリプト スクリプト src src テスト テスト ベンダー ベンダー .dockerignore .dockerignore .gitignore .gitignore .gitmodules .gitmodules .stylua.toml .stylua.toml ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンスライセンス MIGRATION_STATUS.md MIGRATION_STATUS.md README.md README.md RELEASE.md RELEASE.md SALSA.md SALSA.md TESTING.md TESTING.md Taskfile.yml Taskfile.yml config.toml config.toml mprocs.yaml.example mprocs.yaml.example Output2.txt Output2.txt run_repro.sh run_repro.sh Rustfmt.toml Rustfmt.toml terminai.example.yaml terminai.example.yaml test_scrollback.sh test_scrollback.sh test_simple_scroll.sh test_simple_scroll.sh tmp.log tmp.log tmp.sh tmp.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Terminai は、実際の AI CLI をオンデマンド オーバーレイに配置する透過的なターミナル ラッパーです。通常どおりシェルを使用し、Ctrl+Space を押して Codex、Claude Code、OpenCode、またはカスタム エージェントを開きます。ライブ ターミナル コンテキストと推奨されるシェル入力への承認ゲート型アクセスが含まれます。
Terminai はアルファ版です。私はこれを毎日のドライバーとして使用していますが、フォールバックとして利用できる通常のシェル プロファイルを常に保持しておく必要があります。
Terminai は、PTY 内で 1 つのシェル (または指定したコマンド) を開始し、ホスト端末のネイティブのスクロールバックとコピーの動作を維持しながら、VT100 エミュレーションでそれをレンダリングします。

r.ラップされた端末はプライマリ インターフェイスのままです。 Terminai は、オーバーレイがアクティブになるまで邪魔にならないようにします。
オーバーレイは、エージェントの実際の CLI を実行する別の PTY ベースの端末です。 Terminai は、モデル クライアントの実装、プロバイダーの選択、またはモデル API キーの保持を行いません。認証、モデルの選択、会話状態、およびネットワーク アクセスは、選択したエージェント CLI の責任となります。
Terminai は、互換性のあるエージェントに、ローカル MCP サーバーを介してシェルへの制御されたアクセスを提供します。
構成可能なパターンベースのプライバシー フィルタリング後に、表示されている端末と最近のスクロールバックを読み取ります。
作業ディレクトリ、シェル、OS、ディメンション、マウス モード、括弧で囲まれた貼り付け状態などのセッション コンテキストを検査します。
ラップされたセッションが変更されると、コンテキストの更新を受信します。
ユーザーが確認して承認または拒否できるように、正確なシェル入力をキューに入れます。
最新の提案の状態を確認します。
ユーザーの承認なしに、提案された入力がラップされたシェルに書き込まれることはありません。
Terminai は macOS と Linux をサポートしています。
brew install emosenkis/tap/terminai
この式では事前にビルドされたリリース バイナリがインストールされるため、Rust は必要ありません。
GitHub Releases からプラットフォームのアーカイブをダウンロードして解凍し、 PATH 上のどこかに terminai を配置します。
このリポジトリは、パッチが適用された Ratatui および Rat-salsa の依存関係に Git サブモジュールを使用します。
git clone --recurse-submodules https://github.com/emosenkis/terminai.git
CD端末
カーゴインストール --path src
クイックスタート
まず、サポートされているエージェント CLI を少なくとも 1 つインストールして認証します。次に例を示します。
コーデックスログイン
# または CLI を使用して Claude Code で認証する
次に、Terminai を起動します。
終点
コマンドを使用しない場合、Terminai は設定されたシェル (または呼び出し元のシェル) を解決します。
Windows の場合)。代わりに特定のコマンドとその引数をラップするには:
終了 -- zsh -l
端末を通常どおり使用し、事前に

ss エージェントが必要な場合は Ctrl+Space。 Ctrl+Space または Esc を押してシェルに戻ります。エージェントが入力をキューに入れたら、それを確認し、y を押して承認するか、n を押して拒否します。これらのバインディングは構成可能です。
エージェント オーバーレイが開いている間に、F10 を押してターミナル コントロールを表示し、F11 を押して終了コントロールを表示します。
フルスクリーンに切り替えるか、F9 キーを押してレイアウト モードにします。レイアウトモードでは、+/-が変化します
AI の高さ、p は上下を切り替え、g はゲスト表示を循環し、f は
フルスクリーンを切り替えます。承認モード、エージェント切り替え、履歴クリアなど
レイアウト設定はメニューからも行うことができます。
端末エミュレータのワークフローの場合は、コマンドが terminai である別のプロファイルを作成し、エミュレータの通常のシェル プロファイルをフォールバックとして保持します。
Terminai は $XDG_CONFIG_HOME/terminai/terminai.yaml から YAML をロードします。または
XDG_CONFIG_HOME が設定されていない場合は、~/.config/terminai/terminai.yaml。 Windows の場合
%APPDATA%\\terminai\\terminai.yaml を使用します (ログ/キャッシュが含まれています)
%LOCALAPPDATA%\\ターミナル)。デフォルトの構成とプロンプトを生成する
テンプレートの内容:
ターミナル初期設定
Windows では、 pwsh.exe 、 powershell.exe を備えた現在の Windows ターミナルを使用します。
または cmd.exe ;認定済みの Windows サポートを参照してください。
環境とシェル選択の優先順位。
デフォルトのエージェントは Codex です。最小限の明示的な構成は次のとおりです。
インターフェース:
チャット位置 : 下
チャットの高さのパーセント: 50
ゲストディスプレイ: サイズ変更
キーバインディング:
オーバーレイのアクティブ化: Ctrl-Space
オーバーレイの非アクティブ化: Ctrl-Space
承認：はい
拒否：n
レイアウトモード: F9
コントロールパネル：F10
全画面切り替え : F11
承認モード : 常に質問する
エージェント：
プリセット：コーデックス
承認モードは、 always-ask または auto-approval にすることができます。自動承認送信
コマンドを参照せずに、エージェントのすべての提案をシェルに直接送信します
リスク分類子。 Terminai は、このモードに ⚠ AUTO-APPROVE のマークを付けます。それを有効にする
アプリ内では確認が必要です。アプリ内モード

d エージェントの変更は次の期間持続します。
現在のセッションのみ。
chat-position は、 top 、bottom 、または fullscreen を受け入れます。
chat-height-percent は分割レイアウトを制御し、20 ～ 80% に固定されます。
guest-display はサイズ変更 (ゲストを残りのスペースにリフロー) を受け入れます。
オーバーレイ (変更されていないゲストの上に AI を描画)、または移動 (変更/トリミング)
AI から離れた変更されていないゲスト）。実行時のレイアウトの変更はセッションの間持続します。
Agent.preset を変更して、別のバンドルされたプリセットに切り替えます。
エージェント：
プリセット: クロード # codex、claude、または opencode
Codex および Claude プリセットは、Terminai のローカル MCP サーバーを有効にし、レンダリングされたコンテキスト プロンプトを自動的に挿入します。 OpenCode はコンテキスト プロンプトを受け取ります。カスタム エージェント サポートは、MCP、ツール CLI、またはその両方をオプトインできます。
組み込みプリセットは、 config/codex.yaml 、 config/claude.yaml 、および config/opencode.yaml からコンパイルされます。ユーザー プリセットは、組み込みプリセットを拡張し、引数を追加できます。
エージェント：
プリセット: codex-fast
エージェントプリセット:
コーデックス高速:
拡張: コーデックス
スイッチャーで表示 : true
追加引数:
- --モデル
- gpt-5
完全にカスタムのエージェント構成では、ランタイム値をコマンドライン引数にレンダリングできます。
エージェント：
種類：カスタム
コマンド: 私のエージェント
use-mcp : true
use-tool-cli : false
引数:
- --mcp-url
- " {{ mcp_url }} "
- --コンテキスト
- " {{ context_prompt }} "
- expr : ' ["--cwd", cwd] if cwd else [] '
文字列引数は Minijinja テンプレートとしてレンダリングされます。 expr エントリは文字列の配列として評価される必要があるため、ゼロ、1 つ、または複数の CLI 引数を発行できます。使用可能な値には、 cwd 、 context_prompt 、 uses_mcp 、 uses_tool_cli 、 mcp_url 、 mcp_command 、 mcp_port 、および tools_command が含まれます。 json フィルターと toml フィルターは、ネストされた CLI 構成に安全なシリアル化を提供します。 MCP ベアラー トークンは、引数に埋め込まれるのではなく、TERMINAI_MCP_AUTH_TOKEN でエージェント プロセスに渡されます。

nts。
バンドルされているプロンプトは config/default.jinja です。 Terminai config ディレクトリ内のdefault.jinja がこれをシャドウします。また、agent.prompt-template をそのディレクトリ内の別のテンプレートに設定することもできます。
カスタム テンプレートは、バンドルされたプロンプトを拡張し、個々のブロックをオーバーライドできます。
{% "builtin/default.jinja" を拡張します %}
{% ブロックの紹介 %} カスタマイズされた紹介。 {% エンドブロック %}
生成された構成参照はすべてのフィールドを文書化し、バージョン管理された JSON スキーマは https://terminai.app/schema-v<version>.json で公開されます。
エージェント ピッカーには、バンドルされたプリセットとユーザー プリセットが含まれています。
プリセットは show-in-switcher: false を設定します。切り替えにより現在のエージェントが終了します
確認後にセッションを開始し、新しいセッションを起動します。 「AI で読み取れる明確な
「履歴」は、Terminai の内部シェルのスクロールバック (現在の画面) のみを削除します。
ターミナル エミュレータのネイティブ スクロールバックはそのまま残ります。
MCP インターフェースと安全境界
Terminai は、認証されたローカル Streamable HTTP MCP エンドポイントを、それを有効にするエージェント プリセットに提供します。エンドポイントは以下を公開します。
セキュリティ境界は意図的に狭くなっています。
選択したエージェント CLI は、資格情報、プロバイダー トラフィック、およびモデルの動作を所有します。
Terminai 自体は端末データのアップロードやモデルのリクエストを行いません。
MCP を通じて返された端末コンテンツは、構成可能なパターンベースのフィルタリングを通過します。機密情報や個人情報が削除されることを保証するものではありません。デフォルトでは、認証情報と強力な個人識別子は編集されますが、URL、IP アドレス、日付、郵便番号、技術診断は保持されます。 Privacy.patterns を、default 、カテゴリ ( credentials 、 Financial 、 Identity 、 Medical 、 crypto 、または gitleaks )、または btc-address などのエンティティ タイプで構成します。エントリを削除するには、エントリの前に - を付けます (例: [default, -btc-address])。プライバシー.s

戦略は、 replace 、 Mask 、 hash 、 encrypt 、および redact をサポートします。
エージェントが提案した入力は、シェル PTY に到達する前に承認フローに入ります。
ユーザーが提案を確認できるように、提案は安全、注意、または危険に分類されます。分類は明示的な承認に代わるものではありません。
ゼロインストール MCP セットアップは、CLI フラグまたは環境変数を介して MCP 構成をサポートするエージェントに依存します。
ホスト端末
└── 終了プロセス
§── ラップされたシェル/コマンド PTY
│ └── VT100 ステート、ネイティブ スクロールバック、入力転送
§── 認証されたローカル MCP サーバー
│ §── ターミナル/コンテキストの読み取り → プライバシー フィルター
│ └── 提案入力 → 分類 → 承認キュー
└── エージェント CLI PTY
└── Codex、Claude Code、OpenCode、またはカスタムコマンド
重要な実装領域:
src/bin/terminai.rs : アプリケーションのエントリ ポイント、イベント ループ、レンダリング、およびオーバーレイの調整。
src/agent_launcher.rs : プリセット解像度、Minijinja レンダリング、エージェント起動計画。
src/agent_terminal.rs : PTY ライフサイクルとエージェント CLI のレンダリング。
src/mcp_host/ : rmcp とストリーミング可能な HTTP トランスポートで構築された認証済み MCP サーバー。
src/agent_tools.rs : MCP から UI 承認フローに渡される提案状態。
src/command/ : 解析と安全性クラス

[切り捨てられた]

## Original Extract

Make your coding AI available in your shell. Contribute to emosenkis/terminai development by creating an account on GitHub.

I wanted a way to use my existing AI subscriptions as a terminal assistant to analyze shell output or help me with a complicated awk incantation from time to time without getting in my way the rest of the time. I couldn't find anything existing that does this do I built it! I used pvolok's excellent mprocs as a starting point and leaned heavily on Claude and Codex for coding all but the gnarliest parts while providing extensive guidance on how it should be implemented. I know that coding agents can run shell commands already but that puts the AI in the driver's seat. Terminai let's me be the driver, with codex or claude (or any other agent) sitting quietly in the backseat until I want its help. The hardest part was making it truly transparent. Most TUIs that embed a terminal break the terminal emulator's native scrollback and frequently also break copying to the clipboard. I worked hard to make sure that that both of these work exactly as they would without terminai. Along the way I had to fork ratatui to change the way it renders in order to get both of these behaviors working. Tested on Mac and Linux as my daily driver and on Windows enough to see that it works. Quick install:
brew install emosenkis/tap/terminai

GitHub - emosenkis/terminai: Make your coding AI available in your shell · GitHub
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
emosenkis
/
terminai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
640 Commits 640 Commits .claude .claude .codex/ skills/ create-release .codex/ skills/ create-release .github/ workflows .github/ workflows config config crossterm-0.28-facade crossterm-0.28-facade docs docs helpers/ print-key helpers/ print-key img img npm npm plans plans rat-salsa @ 3173ef6 rat-salsa @ 3173ef6 ratatui @ 10bc72a ratatui @ 10bc72a scoop scoop scripts scripts src src tests tests vendor vendor .dockerignore .dockerignore .gitignore .gitignore .gitmodules .gitmodules .stylua.toml .stylua.toml ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE MIGRATION_STATUS.md MIGRATION_STATUS.md README.md README.md RELEASE.md RELEASE.md SALSA.md SALSA.md TESTING.md TESTING.md Taskfile.yml Taskfile.yml config.toml config.toml mprocs.yaml.example mprocs.yaml.example output2.txt output2.txt run_repro.sh run_repro.sh rustfmt.toml rustfmt.toml terminai.example.yaml terminai.example.yaml test_scrollback.sh test_scrollback.sh test_simple_scroll.sh test_simple_scroll.sh tmp.log tmp.log tmp.sh tmp.sh View all files Repository files navigation
Terminai is a transparent terminal wrapper that puts a real AI CLI in an on-demand overlay. Use your shell normally, then press Ctrl+Space to open Codex, Claude Code, OpenCode, or a custom agent with live terminal context and approval-gated access to suggested shell input.
Terminai is in alpha. I use it as my daily driver, but you should always keep an ordinary shell profile available as a fallback.
Terminai starts one shell (or a command you provide) inside a PTY and renders it with VT100 emulation while preserving the host terminal's native scrollback and copy behavior. The wrapped terminal remains the primary interface; Terminai stays out of the way until the overlay is activated.
The overlay is another PTY-backed terminal running the agent's actual CLI. Terminai does not implement a model client, choose a provider, or hold model API keys. Authentication, model selection, conversation state, and network access remain the responsibility of the selected agent CLI.
Terminai gives compatible agents controlled access to the shell through a local MCP server:
Read the visible terminal and recent scrollback, after configurable pattern-based privacy filtering.
Inspect session context such as the working directory, shell, OS, dimensions, mouse mode, and bracketed-paste state.
Receive context updates as the wrapped session changes.
Queue exact shell input for the user to review and approve or deny.
Check the state of the most recent suggestion.
Suggested input is never written to the wrapped shell without user approval.
Terminai supports macOS and Linux.
brew install emosenkis/tap/terminai
The formula installs a prebuilt release binary, so Rust is not required.
Download the archive for your platform from GitHub Releases , unpack it, and place terminai somewhere on your PATH .
The repository uses Git submodules for its patched Ratatui and rat-salsa dependencies.
git clone --recurse-submodules https://github.com/emosenkis/terminai.git
cd terminai
cargo install --path src
Quick start
First install and authenticate at least one supported agent CLI, for example:
codex login
# or authenticate with Claude Code using its CLI
Then launch Terminai:
terminai
With no command, Terminai resolves the configured shell (or the invoking shell
on Windows). To wrap a specific command and its arguments instead:
terminai -- zsh -l
Use the terminal normally and press Ctrl+Space when you want the agent. Press Ctrl+Space or Esc to return to the shell. When an agent queues input, review it and press y to approve or n to deny; these bindings are configurable.
While the agent overlay is open, press F10 for Terminai Controls, F11 to
toggle fullscreen, or F9 for Layout Mode. In Layout Mode, + / - changes
AI height, p toggles top/bottom, g cycles the guest display, and f
toggles fullscreen. Approval mode, agent switching, history clearing, and all
layout settings are also available through the menus.
For a terminal-emulator workflow, create a separate profile whose command is terminai and keep the emulator's normal shell profile as a fallback.
Terminai loads YAML from $XDG_CONFIG_HOME/terminai/terminai.yaml , or
~/.config/terminai/terminai.yaml when XDG_CONFIG_HOME is unset. On Windows
it uses %APPDATA%\\terminai\\terminai.yaml (with logs/cache in
%LOCALAPPDATA%\\terminai ). Generate the default configuration and prompt
template with:
terminai init-config
On Windows, use a current Windows Terminal with pwsh.exe , powershell.exe ,
or cmd.exe ; see Windows support for the qualified
environment and shell-selection precedence.
The default agent is Codex. A minimal explicit configuration is:
interface :
chat-position : bottom
chat-height-percent : 50
guest-display : resize
key_bindings :
activate-overlay : Ctrl-Space
deactivate-overlay : Ctrl-Space
approve : y
deny : n
layout-mode : F9
control-panel : F10
toggle-fullscreen : F11
approval-mode : always-ask
agent :
preset : codex
approval-mode can be always-ask or auto-approval . Auto-approval sends
every agent suggestion directly to the shell without consulting the command
risk classifier. Terminai marks this mode with ⚠ AUTO-APPROVE ; enabling it
in-app requires confirmation. In-app mode and agent changes last for the
current session only.
chat-position accepts top , bottom , or fullscreen .
chat-height-percent controls split layouts and is clamped to 20–80%.
guest-display accepts resize (reflow the guest into the remaining space),
overlay (draw AI over the unchanged guest), or move (shift/crop the
unchanged guest away from AI). Runtime layout changes last for the session.
Switch to another bundled preset by changing agent.preset :
agent :
preset : claude # codex, claude, or opencode
The Codex and Claude presets enable Terminai's local MCP server and inject the rendered context prompt automatically. OpenCode receives the context prompt; custom agent support can opt into MCP, the tool CLI, or both.
Built-in presets are compiled from config/codex.yaml , config/claude.yaml , and config/opencode.yaml . User presets can extend a built-in preset and append arguments:
agent :
preset : codex-fast
agent-presets :
codex-fast :
extends : codex
show-in-switcher : true
extra-args :
- --model
- gpt-5
A fully custom agent configuration can render runtime values into its command-line arguments:
agent :
kind : custom
command : my-agent
uses-mcp : true
uses-tool-cli : false
args :
- --mcp-url
- " {{ mcp_url }} "
- --context
- " {{ context_prompt }} "
- expr : ' ["--cwd", cwd] if cwd else [] '
String arguments are rendered as Minijinja templates. An expr entry must evaluate to an array of strings and can therefore emit zero, one, or multiple CLI arguments. Available values include cwd , context_prompt , uses_mcp , uses_tool_cli , mcp_url , mcp_command , mcp_port , and tool_command ; the json and toml filters provide safe serialization for nested CLI configuration. The MCP bearer token is passed to the agent process in TERMINAI_MCP_AUTH_TOKEN rather than embedded in arguments.
The bundled prompt is config/default.jinja . A default.jinja in the Terminai config directory shadows it. You can also set agent.prompt-template to another template in that directory.
Custom templates can extend the bundled prompt and override individual blocks:
{% extends "builtin/default.jinja" %}
{% block introduction %} Your customized introduction. {% endblock %}
The generated configuration reference documents every field, and versioned JSON Schemas are published at https://terminai.app/schema-v<version>.json .
The agent picker includes bundled presets and user presets unless a user
preset sets show-in-switcher: false . Switching terminates the current agent
session after confirmation and launches a fresh one. “Clear AI-readable
history” removes only Terminai's internal shell scrollback: the current screen
and terminal emulator's native scrollback remain intact.
MCP interface and safety boundary
Terminai serves an authenticated, local Streamable HTTP MCP endpoint to agent presets that enable it. The endpoint exposes:
The security boundary is deliberately narrow:
The selected agent CLI owns credentials, provider traffic, and model behavior.
Terminai itself does not upload terminal data or make model requests.
Terminal contents returned through MCP pass through configurable, pattern-based filtering; it is not a guarantee that secrets or private information are removed. By default it redacts credentials and strong personal identifiers but retains URLs, IP addresses, dates, postal codes, and technical diagnostics. Configure privacy.patterns with default , a category ( credentials , financial , identity , medical , crypto , or gitleaks ), or an entity type such as btc-address ; prefix an entry with - to remove it, for example [default, -btc-address] . privacy.strategy supports replace , mask , hash , encrypt , and redact .
Agent-suggested input enters an approval flow before reaching the shell PTY.
Suggestions are classified as safe, caution, or dangerous to help the user review them; classification does not replace explicit approval.
Zero-install MCP setup depends on the agent supporting MCP configuration through CLI flags or environment variables.
host terminal
└── Terminai process
├── wrapped shell/command PTY
│ └── VT100 state, native scrollback, input forwarding
├── authenticated local MCP server
│ ├── terminal/context reads → privacy filter
│ └── input suggestions → classification → approval queue
└── agent CLI PTY
└── Codex, Claude Code, OpenCode, or custom command
Important implementation areas:
src/bin/terminai.rs : application entry point, event loop, rendering, and overlay coordination.
src/agent_launcher.rs : preset resolution, Minijinja rendering, and agent launch plans.
src/agent_terminal.rs : PTY lifecycle and rendering for the agent CLI.
src/mcp_host/ : authenticated MCP server built with rmcp and Streamable HTTP transport.
src/agent_tools.rs : suggestion state passed from MCP into the UI approval flow.
src/command/ : parsing and safety classif

[truncated]
