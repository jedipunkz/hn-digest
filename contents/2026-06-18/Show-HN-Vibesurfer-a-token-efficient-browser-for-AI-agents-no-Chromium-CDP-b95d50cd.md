---
source: "https://github.com/frane/vibesurfer"
hn_url: "https://news.ycombinator.com/item?id=48587641"
title: "Show HN: Vibesurfer – a token-efficient browser for AI agents, no Chromium/CDP"
article_title: "GitHub - frane/vibesurfer · GitHub"
author: "frb"
captured_at: "2026-06-18T18:57:53Z"
capture_tool: "hn-digest"
hn_id: 48587641
score: 3
comments: 0
posted_at: "2026-06-18T16:13:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Vibesurfer – a token-efficient browser for AI agents, no Chromium/CDP

- HN: [48587641](https://news.ycombinator.com/item?id=48587641)
- Source: [github.com](https://github.com/frane/vibesurfer)
- Score: 3
- Comments: 0
- Posted: 2026-06-18T16:13:46Z

## Translation

タイトル: Show HN: Vibesurfer – AI エージェント用のトークン効率の高いブラウザ、Chromium/CDP なし
記事タイトル: GitHub - frane/vibesurfer · GitHub
説明: GitHub でアカウントを作成して、frane/vibesurfer の開発に貢献します。
HN テキスト: 数多くの Chrome/Playwright/Puppeteer スキルや MCP プラグインにうんざりしていましたが、私にとっては非常に重くて不安定に感じられました。そこで、エージェントが Web サイトを効率的にテストして操作できるようにするためにこれを構築しました。

記事本文:
GitHub - フレーム/バイブサーファー · GitHub
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
フレーム
/
バイブサーファー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

104 コミット 104 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows crates crates dist dist docs docs fixtures fixtures プラグイン プラグイン スクリプト スクリプト スキル スキル テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile.linux-test Dockerfile.linux-test LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md gemini-extension.json gemini-extension.json install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカル AI エージェント用の実際のブラウザー。
エージェントがブラウザ経由で Web アプリをテストできるようにしたいと考えていました。私が試したものはすべて (Playwright、Puppeteer、その他 CDP をラップするもの) 重すぎて不安定すぎました。 CDP はセッションをドロップします。 Playwright は長時間実行するとクラッシュします。 Chrome はリリースごとに太くなっています。しかし、それは実際の問題ではありません。 CDP と Chrome は、DevTools を見つめる人間のために設計されました。これらは、while ループに陥ったエージェント向けに設計されたものではありません。
エージェントはトークンごとに支払います。応答ごとにブロックします。イベント Firehose に対処できず、読み取りごとに 4kb DOM ダンプを実行すると、コンテキスト バジェットが急速に消費されてしまいます。 Playwright を介した Hacker News のトップ ページには、エージェントが何かを行う前に約 2000 の入力トークンが表示されます。バイブサーファー経由では50くらいです。
Vibesurfer は、Rust のネイティブ ブラウザ デーモンです。読み取りは、完全な DOM の代わりに状態トークンとツリー デルタを返します。トークンを書き込みチェックします。読み取りと書き込みの間に何かが移動した場合、呼び出しは失敗し、エージェントは古いページをクリックする代わりに再読み取りを行います。その下には 3 つの実際のエンジンがあります: macOS の WKWebView、Linux の WebKitGTK、Windows の WebView2。上部のプロトコルはテキストおよび行指向です。
実際にピクセルが必要な場合は、スクリーンショット用のキャプチャと、モバイルとデスクトップのレイを切り替えるためのビューポートがあります。

アウト、および vs レイアウトを使用して境界ボックスを取得します。しかし、テキストが最初にあります。
醸造タップ フラン/タップ && 醸造インストール バイブサーファー
カール:
カール -sSL https://raw.githubusercontent.com/frane/vibesurfer/main/install.sh |しー
貨物:
カーゴインストールバイブサーファー
ソースより:
git clone https://github.com/frane/vibesurfer && cd vibesurfer
カーゴインストール --path crates/vs-cli
Linux には WebKitGTK 6 が必要です。Windows には WebView2 ランタイムが必要です (Windows 11 にはすでに搭載されていますが、Windows 10 では Microsoft から入手可能です)。
2 つの統合パスがあり、それらは独立しています。次のいずれかまたは両方をインストールできます。
Skill : SKILL.md をエージェントのスキル ディレクトリにドロップします。エージェントはそれをコンテキストとして読み取り、エージェントが持つシェルを通じて vs バイナリを直接呼び出します。これは、Bash を実行するが MCP を話さないエージェントに使用します。
MCP : vs mcp を MCP サーバーとして登録します。エージェントは、JSON-RPC 経由でバイブサーファー プリミティブを MCP ツールとして呼び出します。シェルは必要ありません。これは、ネイティブ MCP サポートを持つエージェントに使用します。
自動インストーラーは、サポートされている場合は両方を実行します。 vs が PATH 上にある場合:
対スキルインストール
Claude Desktop、Claude Code、Cursor、Codex CLI、Gemini CLI、OpenClaw を検出し、SKILL.md と MCP エントリをそれぞれに書き込みます。 2 つのうち 1 つだけをサポートするエージェントは、関連する部分のみを取得します。アップグレード後に再実行します。
スキル パスの場合は、skills/vibesurfer/SKILL.md をリポジトリからエージェントのスキル ディレクトリにコピーします。クロード家のエージェントの場合、通常は ~/.claude/skills/vibesurfer/SKILL.md です。
MCP パスの場合、このブロックをエージェントの MCP 構成 (claude_desktop_config.json 、 .cursor/mcp.json など) に追加します。
{
"mcpサーバー": {
"バイブサーファー" : {
"コマンド" : " vs " 、
"args" : [ "mcp " ]
}
}
}
Codex は、 [mcp_servers.vibesurfer] で同じ形状の TOML を使用します。リポジトリからコピーしたい場合は、JSON フォームも plugin/.mcp.json にあります。
クラウ

de Code マーケットプレイスは、1 つのコマンドで両方のサーフェスをインストールします。
/プラグインインストールフレーム/vibesurfer
リポジトリのルートで .claude-plugin/marketplace.json と plugin/.claude-plugin/plugin.json を解決します。
Gemini 拡張機能は、MCP サーバーと GEMINI.md コンテキスト ファイルを接続します。
gemini 拡張機能のインストール https://github.com/frane/vibesurfer
リポジトリのルートにある gemini-extension.json を読み取ります。
最新のボット対策システム (Cloudflare、hCaptcha、reCAPTCHA、DataDome) は、まず、event.isTrusted でゲートし、次に移動のタイミングでゲートし、次に TLS/HTTP フィンガープリントでゲートします。バイブサーファーの入力ディスパッチは、3 つすべてを渡すように構築されています。
macOS では、 vs act click と座標プリミティブ ( vs click-at 、 vs hover-at 、 vs move-to 、 vsrag ) は、 WKWebView 上のネイティブ NSEvent の MouseDown / MouseUp / MouseMoved を介してルーティングされます。各イベントは、実際のカーソルクリックと同じように、JS で isTrusted = true を伝えます。クリックするたびに送出されるベジェ パスのリードインは、フィッツ則の到着タイミング (ダイグラフから導出された制御点、オプションのオーバーシュート) を再現するため、目に見える動きはテレポートではなくターゲットに到達する人間のように見えます。
v0.1.11 以降、座標プリミティブは Linux と Windows でもネイティブです。WebKitGTK では x11rb を介した XTest (X11 / Xwayland。純粋な Wayland は ENGINE_UNSUPPORTED に戻ります)、Windows では WebView2 コンポジション コントローラーで SendMouseInput を使用します。 3 つのエンジンはすべて、カーソル プリミティブのクリックに対して isTrusted = true を発行します。 Ref ベースと act のクリックは macOS でのみ信頼されます。Linux と Windows では、挿入された JS ( isTrusted = false ) を介してディスパッチされます。指紋に敏感なサイトには、そこにある座標プリミティブを使用します。
ウォーカーは、ARIA role="..." (Radix UI、Headless UI、Reach UI、すべてのカスタム div-as-button パターン) に加えて、ロールのないフォーカス可能な div/span に対する tabindex ヒューリスティックも尊重します。最新の React UI が実用的な参照として表示される

座標の回避策がありません。
すべてのプリミティブには 1 ～ 3 文字のエイリアスがあります。文書化には長い形式が存在します。エージェントの呼び出しでは、トークンを節約するために短い形式を使用する必要があります。
頻繁に使用されるフラグ: --session= / -S 、 --full / -F 、 --since= / -s 、 --limit= / -n 、 --page= / -P 、 --json / -j 。 Inspect サブコマンドにも 1 文字または 2 文字のエイリアスがあります ( i co は Inspection console 、 i n は network 、 i req は request 、 i e は eval 、 i s は storage 、 i scr は scripts 、 i src は script 、 i d は dom 、 i p はパフォーマンスを表します)。
どちらの形式もどこでも機能します。統合テストでは、短い形式からのワイヤ リクエストが長い形式からのワイヤ リクエストとバイト同一であることがアサートされます。
$ vs so # セッションオープン
@0 # 状態トークン (16 進数の文字。0 はまだ何もないことを意味します)
s_019e08a7… # セッションID
$ vs o https://example.com # URL を開きます
@0 # オープンコールにはスナップショットは含まれません
p_019e08a7… # ページ ID
$ vs v p_019e08a7… # 表示 (a11y ツリーのスナップショット)
@44d01704049d6d31 # 状態トークン
1 ドキュメント「サンプル ドメイン」 # ref 1、ドキュメント
0 el "" # 名前のないラッパー
2 hd "サンプル ドメイン" # ref 2、見出し
3 p 「このドメインは…で使用されます。」 # ref 3、段落
5 p「詳細はこちら」 # ref 5、段落
4 lnk 「詳細」をクリック、フォーカス # ref 4、リンク、サポートされている操作
スナップショットは参照のリストです。各 ref はスナップショット間で存続する整数であるため、エージェントはページ全体を再度読み取ることなく、10 ターン後に ref 4 に作用できます。 2 文字のコード ( hd 、 p 、 lnk 、 btn 、 tf 、…) は、ARIA 文字列の代わりにロールを数バイトに圧縮します。ラベルは引用符で囲みます。ラベルの後の末尾のトークンは、要素がサポートする操作と操作をリストします。合計約 20 個の役割コードが docs/PROTOCOL.md にリストされています。
$ 対 4 クリック # 動作: ref 4 をクリック
@<new-token> # 新しいトークン、ページが変更されました
?nav # 警告: ナビゲーションが発生しました
…新しい木…

# act 応答にはデルタが含まれます。
# ナビゲーション時に完全なツリーにベースラインを再設定します
vs act は唯一の変更可能なプリミティブです。これは ref と操作 ( click 、 fill 、scroll 、 key 、 submit 、 hover 、 focus ) を受け取り、最新の状態トークンが必要です。ページが読み取りと書き込みの間で変更された場合 (JS タイマーが起動した場合、WebSocket が更新をプッシュした場合など)、呼び出しは ! を返します。 STALE_TOKEN とエージェントが再読み取りします。静かな古いクリック音はありません。同じページでの操作が成功した後 (ナビゲーションなし)、応答にはデルタ (追加の場合は +ref、削除の場合は -ref、属性の変更の場合は ~ref) のみが含まれるため、ボタンを 1 つ追加するクリックには、DOM 全体ではなくワイヤ上で最大 20 バイトのコストがかかります。
$ vs st # ステータス
セッション s_019e08a7… ページ=1
ページ p_019e08a7… url=https://www.iana.org/help/example-domains token=…
すべてのプリミティブ呼び出しは、返される前に SQLite 監査ログに 1 行を書き込みます。 vs status はそのログを読み取ります。 vs log も同様です。リプレイ、デバッグ、ガバナンスはすべて、 ~/.vibesurfer/state.db に対する SQL クエリに集約されます。サブスクライブする個別のイベント ストリームはありません。
デーモンは最初の呼び出し時に自動生成されます。状態、キャプチャ、ダウンロードは ~/.vibesurfer/ の下にあります。トランスポートは、Unix では AF_UNIX ソケット ( ~/.vibesurfer/daemon.sock ) であり、Windows では Windows 名前付きパイプです。どちらの場合でも、CLI が違いを処理します。
合計 29 個のワイヤー プリミティブ — 19 個のコア プリミティブは docs/PRIMITIVES.md で指定されます。その後の追加 ( vs_inspect 、4 つのカーソル プリミティブ、プロンプト入力、保留キュー) は、バンドルされた SKILL.md と CHANGELOG に文書化されています。すべてのシジルとエッジケースを含む完全なワイヤ形式は docs/PROTOCOL.md にあります。プラットフォームごとのプリミティブごとの検証マトリックスは docs/REALITY_CHECK.md にあります。
パス/変数
目的
~/.vibesurfer/state.db
SQLite、セッションの保持、監査

、マーク、認証 BLOB
~/.vibesurfer/daemon.sock (Unix)
CLI が通信する AF_UNIX ソケット
Windows の名前付きパイプ
Windows でも同じ役割。自動的に解決される
~/.vibesurfer/captures/
vsキャプチャのスクリーンショット
VS_CAPTURES_DIR
キャプチャ ディレクトリをオーバーライドする
VS_HOME
バイブサーファーのホーム ディレクトリをオーバーライドする
VS_DISABLE_INSPECTOR=1
インスペクターフックをスキップする (テストのみ)
VS_DAEMON_BIN
デーモンの自動生成 (テスト) に使用されるバイナリをオーバーライドします。
ソースからビルドする
Rust 1.85以降が必要です。プラットフォーム固有の依存関係:
macOS (15+): 余分なものは何もなく、システム WebKit にリンクします。
Linux : libwebkitgtk-6.0-dev 、 libgtk-4-dev 、 libsoup-3.0-dev 。
Windows : ビルド時に webview2-com によってプルされた WebView2 SDK。実行時には WebView2 ランタイムが必要です。
git clone https://github.com/frane/vibesurfer && cd vibesurfer
カーゴビルド --release
テスト スイートを実行します。
カーゴテスト --workspace --lib --bins # 高速単体テスト
カーゴテスト --workspace # 統合テストを追加します (実際のエンジン)
Linux 以外のホストで Linux エンジンをテストするには、Docker コンテナを使用します。 WebKitGTK 6 のサンドボックスには特権のないユーザー名前空間が必要です。 CI Linux ジョブは、ベア ランナー上の 1 つの sysctl で AppArmor 制限を緩和しますが、Docker フォールバックでも同じことを行うには --privileged が必要です。
docker build -f ドッケ

[切り捨てられた]

## Original Extract

Contribute to frane/vibesurfer development by creating an account on GitHub.

I grew tired of numerous Chrome/Playwright/Puppeteer skills and MCP plugins, which for me felt quite heavy and unstable. So I built this to allow my agents to efficiently test and work with websites.

GitHub - frane/vibesurfer · GitHub
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
frane
/
vibesurfer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
104 Commits 104 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows crates crates dist dist docs docs fixtures fixtures plugin plugin scripts scripts skills skills tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile.linux-test Dockerfile.linux-test LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md SECURITY.md SECURITY.md gemini-extension.json gemini-extension.json install.sh install.sh View all files Repository files navigation
A real browser for your local AI agent.
I wanted agents to test web apps via the browser. Everything I tried (Playwright, Puppeteer, anything else that wraps CDP) was too heavy and too unstable. CDP drops sessions. Playwright crashes on long runs. Chrome gets fatter every release. None of that is the actual problem though. CDP and Chrome were designed for humans staring at DevTools. They were never designed for an agent stuck in a while loop.
An agent pays per token. It blocks per response. It can't deal with the event firehose, and a 4kb DOM dump on every read burns the context budget fast. The Hacker News front page through Playwright is about 2000 input tokens before the agent has done anything. Through vibesurfer it's around 50.
vibesurfer is a native browser daemon in Rust. Reads return state tokens and tree deltas instead of the full DOM. Writes check the token. If anything moved between the read and the write, the call fails and the agent re-reads instead of clicking on a stale page. There are three real engines underneath: WKWebView on macOS, WebKitGTK on Linux, WebView2 on Windows. The protocol on top is text and line-oriented.
When you actually need pixels there's vs capture for screenshots, vs viewport to switch between mobile and desktop layouts, and vs layout to get bounding boxes. But text comes first.
brew tap frane/tap && brew install vibesurfer
curl:
curl -sSL https://raw.githubusercontent.com/frane/vibesurfer/main/install.sh | sh
Cargo:
cargo install vibesurfer
From source:
git clone https://github.com/frane/vibesurfer && cd vibesurfer
cargo install --path crates/vs-cli
Linux needs WebKitGTK 6. Windows needs the WebView2 runtime (already on Windows 11, available for Windows 10 from Microsoft).
Two integration paths, and they're independent. You can install either or both:
Skill : drop SKILL.md into the agent's skills directory. The agent reads it as context and calls the vs binary directly through whatever shell it has. Use this for any agent that runs Bash but doesn't speak MCP.
MCP : register vs mcp as an MCP server. The agent calls vibesurfer primitives as MCP tools over JSON-RPC, no shell required. Use this for agents with native MCP support.
The auto-installer does both where supported. After vs is on your PATH:
vs skill install
It detects Claude Desktop, Claude Code, Cursor, Codex CLI, Gemini CLI, and OpenClaw, then writes the SKILL.md plus the MCP entry into each one. Agents that only support one of the two get only the relevant piece. Re-run after upgrading.
For the skill path , copy skills/vibesurfer/SKILL.md from the repo into the agent's skills directory. For Claude-family agents that's typically ~/.claude/skills/vibesurfer/SKILL.md .
For the MCP path , add this block to the agent's MCP config ( claude_desktop_config.json , .cursor/mcp.json , etc.):
{
"mcpServers" : {
"vibesurfer" : {
"command" : " vs " ,
"args" : [ " mcp " ]
}
}
}
Codex uses TOML with the same shape under [mcp_servers.vibesurfer] . The JSON form also sits at plugin/.mcp.json if you would rather copy it from the repo.
Claude Code marketplace installs both surfaces from one command:
/plugin install frane/vibesurfer
Resolves .claude-plugin/marketplace.json at the repo root and plugin/.claude-plugin/plugin.json .
Gemini extension wires the MCP server plus the GEMINI.md context file:
gemini extensions install https://github.com/frane/vibesurfer
Reads gemini-extension.json at the repo root.
Modern anti-bot systems (Cloudflare, hCaptcha, reCAPTCHA, DataDome) gate first on event.isTrusted , then on movement timing, then on TLS/HTTP fingerprinting. vibesurfer's input dispatch is built to pass all three.
On macOS, vs act click and the coordinate primitives ( vs click-at , vs hover-at , vs move-to , vs drag ) route through native NSEvent mouseDown / mouseUp / mouseMoved on WKWebView . Each event carries isTrusted = true in JS, same as a real cursor click. The Bezier-pathed lead-in dispatched before every click reproduces Fitts-law arrival timing (digraph-derived control points, optional overshoot) so the visible motion looks like a human reaching the target rather than a teleport.
Since v0.1.11 the coordinate primitives are native on Linux and Windows too: XTest over x11rb on WebKitGTK (X11 / Xwayland; pure Wayland falls back to ENGINE_UNSUPPORTED ), SendMouseInput on a WebView2 composition controller on Windows. All three engines emit isTrusted = true for cursor-primitive clicks. Ref-based vs act click is trusted on macOS only — on Linux and Windows it still dispatches through injected JS ( isTrusted = false ); use the coordinate primitives there for fingerprint-sensitive sites.
The walker also honors ARIA role="..." (Radix UI, Headless UI, Reach UI, every custom-div-as-button pattern), plus a tabindex heuristic for focusable divs/spans without a role. Modern React UIs surface as actionable refs without coordinate workarounds.
Every primitive has a one-to-three-letter alias. Long forms exist for documentation; agent invocations should use the short form to save tokens.
Frequent flags: --session= / -S , --full / -F , --since= / -s , --limit= / -n , --page= / -P , --json / -j . Inspect subcommands have one-or-two-letter aliases too ( i co for inspect console , i n for network , i req for request , i e for eval , i s for storage , i scr for scripts , i src for script , i d for dom , i p for performance ).
Both forms work everywhere. The integration tests assert that the wire request from a short form is byte-identical to the wire request from the long form.
$ vs so # session-open
@0 # state token (16 hex chars; 0 means none yet)
s_019e08a7… # session id
$ vs o https://example.com # open the URL
@0 # the open call doesn't carry a snapshot
p_019e08a7… # page id
$ vs v p_019e08a7… # view (snapshot the a11y tree)
@44d01704049d6d31 # state token
1 doc "Example Domain" # ref 1, document
0 el "" # nameless wrapper
2 hd "Example Domain" # ref 2, heading
3 p "This domain is for use in…" # ref 3, paragraph
5 p "Learn more" # ref 5, paragraph
4 lnk "Learn more" click,focus # ref 4, link, supported ops
A snapshot is a list of refs. Each ref is an integer that survives across snapshots, so the agent can act on ref 4 ten turns later without re-reading the whole page. The two-letter codes ( hd , p , lnk , btn , tf , …) compress the role into a few bytes instead of an ARIA string. Labels are in quotes; the trailing tokens after a label list which vs act operations the element supports. About twenty role codes total, listed in docs/PROTOCOL.md .
$ vs a 4 click # act: click ref 4
@<new-token> # new token, page mutated
?nav # warning: navigation occurred
… new tree … # the act response carries deltas;
# on navigation it re-baselines to a full tree
vs act is the only mutating primitive. It takes a ref and an operation ( click , fill , scroll , key , submit , hover , focus ) and requires the most recent state token. If the page mutated between read and write (a JS timer fired, a websocket pushed an update, anything), the call returns ! STALE_TOKEN and the agent re-reads. No silent stale clicks. After a successful act on the same page (no navigation), the response carries only the deltas ( +ref for adds, -ref for removes, ~ref for attribute changes), so a click that adds one button costs ~20 bytes on the wire instead of the whole DOM.
$ vs st # status
session s_019e08a7… pages=1
page p_019e08a7… url=https://www.iana.org/help/example-domains token=…
Every primitive call writes one row to a SQLite audit log before it returns. vs status reads that log. So does vs log . Replay, debugging, and governance all collapse to SQL queries against ~/.vibesurfer/state.db . There is no separate event stream to subscribe to.
The daemon auto-spawns on first call. State, captures, and downloads live under ~/.vibesurfer/ . The transport is an AF_UNIX socket on Unix ( ~/.vibesurfer/daemon.sock ) and a Windows named pipe on Windows; either way, the CLI handles the difference.
29 wire primitives total — the 19 core primitives are specified in docs/PRIMITIVES.md ; the later additions ( vs_inspect , the four cursor primitives, prompt-input, and the pending queue) are documented in the bundled SKILL.md and the CHANGELOG. The full wire format with every sigil and edge case is in docs/PROTOCOL.md . The per-platform per-primitive verification matrix is in docs/REALITY_CHECK.md .
Path / variable
Purpose
~/.vibesurfer/state.db
SQLite, holds sessions, audit, marks, auth blobs
~/.vibesurfer/daemon.sock (Unix)
AF_UNIX socket the CLI talks to
Windows named pipe
Same role on Windows; resolved automatically
~/.vibesurfer/captures/
Screenshots from vs capture
VS_CAPTURES_DIR
Override the capture directory
VS_HOME
Override the vibesurfer home directory
VS_DISABLE_INSPECTOR=1
Skip inspector hooks (testing only)
VS_DAEMON_BIN
Override the binary used for daemon auto-spawn (tests)
Build from source
Requires Rust 1.85+. Platform-specific dependencies:
macOS (15+): nothing extra, links against system WebKit.
Linux : libwebkitgtk-6.0-dev , libgtk-4-dev , libsoup-3.0-dev .
Windows : WebView2 SDK pulled by webview2-com at build time; the WebView2 Runtime is required at run time.
git clone https://github.com/frane/vibesurfer && cd vibesurfer
cargo build --release
Run the test suite:
cargo test --workspace --lib --bins # fast unit tests
cargo test --workspace # adds integration tests (real engine)
For Linux engine tests on a non-Linux host, use the Docker container. WebKitGTK 6's sandbox needs unprivileged user namespaces; the CI Linux job relaxes the AppArmor restriction with one sysctl on the bare runner, while the Docker fallback needs --privileged to do the same:
docker build -f Docke

[truncated]
