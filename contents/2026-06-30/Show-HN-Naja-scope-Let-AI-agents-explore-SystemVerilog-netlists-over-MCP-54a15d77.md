---
source: "https://github.com/najaeda/naja-scope"
hn_url: "https://news.ycombinator.com/item?id=48733230"
title: "Show HN: Naja-scope – Let AI agents explore SystemVerilog netlists over MCP"
article_title: "GitHub - najaeda/naja-scope · GitHub"
author: "xtofalex"
captured_at: "2026-06-30T14:50:46Z"
capture_tool: "hn-digest"
hn_id: 48733230
score: 1
comments: 0
posted_at: "2026-06-30T14:31:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Naja-scope – Let AI agents explore SystemVerilog netlists over MCP

- HN: [48733230](https://news.ycombinator.com/item?id=48733230)
- Source: [github.com](https://github.com/najaeda/naja-scope)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T14:31:03Z

## Translation

タイトル: Show HN: Naja-scope – AI エージェントが MCP 経由で SystemVerilog ネットリストを探索できるようにする
記事タイトル: GitHub - najaeda/naja-scope · GitHub
説明: GitHub でアカウントを作成して、najaeda/naja-scope の開発に貢献します。

記事本文:
GitHub - ナジャエダ/ナジャスコープ · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ナジェダ
/
ナジャスコープ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン

ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
65 コミット 65 コミット .github/ workflows .github/ workflows 例 例 src/ naja_scope src/ naja_scope テスト テスト .gitignore .gitignore LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md README_PyPI.md README_PyPI.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
チャットにソース コードを貼り付けることなく、AI アシスタントが SystemVerilog の設計を探索できるようにします。
naja-scope は AI を提供する MCP サーバーです
エージェント (クロード、および MCP 互換アシスタント) の正確で構造化されたビュー
精緻な SystemVerilog 設計。何千行ものデータをダンプする代わりに、
RTL をモデルのコンテキストに取り入れ、エージェントは的を絞った質問をします - 何が動機になっているのか
この信号は？このモジュールには何が入っているのでしょうか？このネットはどこから来たのですか？ —そして
ファイルと行の参照を含む小さくて正確な答えが返されます。
najaeda ネットリスト エンジンに基づいて構築されています。
大きなデザインはチャット ウィンドウに収まりません。 RTL の貼り付けは時間がかかり、コストがかかり、
モデルは依然として階層間の接続を確実に追跡できません。ナジャスコープ
あなたのデザインをエージェントがナビゲートできるものに変えます。
🔎 接続をトレース — 信号を駆動または負荷しているものを特定します。
モジュールの境界。
🌲 階層をたどる — オンデマンドでモジュール、インスタンス、ポートを探索します。
🎯 ソースにジャンプ — すべての答えには file:line 範囲が含まれているため、
エージェントは重要な RTL を正確に引用できます。
🧩 ロジック コーン — ファンイン/ファンアウトの組み合わせコーンをトレースします。
レジスタ境界。
💡 設計意図を復元 — enum 状態名、構造体/共用体フィールド、および
通常、設計が精緻化されると消えてしまうパラメータ式。
RTL とゲートレベルのネットリストで同様に動作します: 精緻化された SystemVerilog をロードし、
または、合成後の構造 Verilog ネットリストをロードします。

その自由とともに
標準セル ライブラリを使用し、同じ方法でゲートをナビゲートします (「
ゲートレベルの設計)。
すべての応答はトークンで区切られています。リストはページ分割され、大きな結果は次のように切り捨てられます。
クリアマーカー。コンテキストは小さいままです。あなたの答えは正確であり続けます。
私たちは CVA6 で直接対決を行いました (
プロダクション RISC-V コア): 同じ 17 の設計に関する質問に、クロードが 1 回回答
naja-scope のみを使用し、 grep /file のみを使用して 1 回読み取ります。
ソースツリー。
より多くの正解、より少ない往復のターン、そして最大 5 分の 1 のトークンの減少 —
エージェントはファイルのスクロールを停止し、構造的な答えに直接進みます。
pip install naja-scope # najaeda と MCP ランタイムを PyPI からプルします
naja-scope-mcp # stdio MCP サーバー
クロードコードに接続します
クロード mcp naja-scope を追加 -- naja-scope-mcp
または、MCP クライアントの構成に追加します。
{
"mcpサーバー": {
"ナジャスコープ" : {
"コマンド" : " naja-scope-mcp "
}
}
}
次に、アシスタントにデザインをロードして探索を開始するように依頼するだけです。
「top uart_top を使用して rtl/uart.sv から UART デザインをロードして、表示してください
tx_o を駆動するすべてのもの。」
エージェントはデザインを一度ロードすると、フォローアップの質問に即座に回答します。いいえ
ソースを再読しても、巨大なペーストはありません。
ChatGPT は、HTTP エンドポイント (カスタム コネクタ /
開発者モード)、stdio ではなく HTTP サーバーとして naja-scope を実行します。
naja-scope-mcp --transport streamable-http --host 127.0.0.1 --port 8000
これは、 http://<host>:8000/mcp で MCP を提供します。 ChatGPTがサーバーに到達するため
ネットワーク上で、その URL を ChatGPT が参照できる場所に公開します。公共の
ローカル実行用のトンネル:
# 例: ローカルサーバーへのトンネル (ngrok、cloudflared など)
ngrok http 8000 # -> https://<something>.ngrok.app → /mcp を追加
次に、ChatGPT で、[設定] → [コネクタ] を開きます (場合は開発者モードを有効にします)
必要)、カスタムを追加します

m コネクタ 、サーバーの URL を貼り付けます
( https://<your-host>/mcp )。接続したら、デザインをロードして探索するように依頼します
まさに上記の通りです。 (ChatGPT のコネクタ UI は進化します。変わらないのは、
--transport streamable-http が提供する HTTPS MCP URL。)
⚠️ HTTP サーバーには認証が組み込まれておらず、信頼されたトンネル経由でのみ公開されます。
ローカル実験には短命のトンネルを好みます。
すでに合成されていますか？構造的な Verilog ネットリストを一緒にロードします。
標準セルを定義し、同じゲートをナビゲートする Liberty ライブラリ
RTLとしての方法:
「Liberty ライブラリ pdk/stdcells.lib をロードしてから、ゲート ネットリストをロードします
build/top.v 、そして、top がどのセルから構築され、どのようなドライブが行われているかを教えてください
data_out 。」
階層、セルごとの数 ( get_module_card )、ドライバー/ロード、およびロジック コーン
すべてネットリスト上で動作します。コーンは連続するセルで停止します。ゲートネットリスト
はソース行情報を持たないため、get_source は RTL にのみ適用されます。実行可能
例は、examples/ ( stdcells.lib + counter2.v +
ゲートレベル.py)。
デザインが読み込まれると、アシスタントは次のことができるようになります。
階層パス (glob と
もしかして提案）。
パターンごとにデザイン全体のオブジェクトを検索します。
任意のモジュールの階層を表示します。
ネットのドライバー/負荷、つまり実際のエンドポイントを階層全体で取得します。
ロジック コーン (ファンイン/ファンアウト) をトレースし、レジスタ フロンティアを確認します。
ソースを取得します。オブジェクトの背後にある正確な SystemVerilog 行を取得します。
モジュール カードを入手 – ポート、カウント、クロック/リセットが一目でわかります。
設計意図の回復 — ステートマシン名、構造体フィールド、パラメーター
推敲中に失われる表現。
実行可能なエンドツーエンドのウォークスルーは、examples/ にあります。
najaeda が実行される場所であればどこでも動作します (Linux、macOS、Windows)
# チェックアウトから
python3 -m venv .venv
.venv/bin/pip install -e 。
.venv/bin/python -m

pytest -q
完全なテスト スイートは、najaeda のプレーンな pip インストールに対して実行されます (ネイティブではありません)。
ビルドが必要です。 CVA6 階層間コーン回帰
( testing/test_zzz_cone_cva6.py ) は遅く、CVA6 でない限り自動的にスキップされます。
スナップショットが存在します。
🐛 バグを見つけましたか? または機能リクエストがありますか?
GitHub で問題を開く →
📫 連絡先: contact@keplertech.io
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to najaeda/naja-scope development by creating an account on GitHub.

GitHub - najaeda/naja-scope · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
najaeda
/
naja-scope
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
65 Commits 65 Commits .github/ workflows .github/ workflows examples examples src/ naja_scope src/ naja_scope tests tests .gitignore .gitignore LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md README_PyPI.md README_PyPI.md pyproject.toml pyproject.toml View all files Repository files navigation
Let your AI assistant explore SystemVerilog designs — without pasting source code into the chat.
naja-scope is an MCP server that gives AI
agents (Claude, and any MCP-compatible assistant) a precise, structured view of
your elaborated SystemVerilog design. Instead of dumping thousands of lines of
RTL into the model's context, the agent asks targeted questions — what drives
this signal? what's inside this module? where does this net come from? — and
gets back small, exact answers with file-and-line references.
Built on the najaeda netlist engine.
Large designs don't fit in a chat window. Pasting RTL is slow, expensive, and
the model still can't reliably trace connectivity across hierarchy. naja-scope
turns your design into something an agent can navigate :
🔎 Trace connectivity — find what drives or loads any signal, across
module boundaries.
🌲 Walk the hierarchy — explore modules, instances, and ports on demand.
🎯 Jump to source — every answer comes with file:line ranges, so the
agent can quote the exact RTL that matters.
🧩 Logic cones — trace fan-in / fan-out combinational cones up to the
register boundary.
💡 Recover design intent — enum state names, struct/union fields, and
parameter formulas that normally vanish when a design is elaborated.
Works on RTL and gate-level netlists alike: load elaborated SystemVerilog,
or load a post-synthesis structural Verilog netlist together with its Liberty
standard-cell library and navigate the gates the same way (see
Gate-level designs ).
All responses are token-bounded: lists paginate, large results truncate with
clear markers. Your context stays small; your answers stay accurate.
We ran a head-to-head on CVA6 (a
production RISC-V core): the same 17 design questions, answered by Claude once
with only naja-scope and once with only grep /file reading over the
source tree.
More correct answers, fewer back-and-forth turns, and ~5× fewer tokens — the
agent stops scrolling through files and goes straight to the structural answer.
pip install naja-scope # pulls najaeda and the MCP runtime from PyPI
naja-scope-mcp # stdio MCP server
Connect it to Claude Code
claude mcp add naja-scope -- naja-scope-mcp
Or add it to any MCP client's config:
{
"mcpServers" : {
"naja-scope" : {
"command" : " naja-scope-mcp "
}
}
}
Then just ask your assistant to load a design and start exploring:
"Load my UART design from rtl/uart.sv with top uart_top , then show me
everything that drives tx_o ."
The agent loads the design once and answers follow-up questions instantly — no
re-reading source, no giant pastes.
ChatGPT connects to MCP servers over an HTTP endpoint (custom connectors /
Developer mode), so run naja-scope as an HTTP server instead of stdio:
naja-scope-mcp --transport streamable-http --host 127.0.0.1 --port 8000
This serves MCP at http://<host>:8000/mcp . Because ChatGPT reaches the server
over the network, expose that URL where ChatGPT can see it — e.g. a public
tunnel for a local run:
# example: a tunnel to your local server (ngrok, cloudflared, …)
ngrok http 8000 # -> https://<something>.ngrok.app → add /mcp
Then in ChatGPT, open Settings → Connectors (enable Developer mode if
needed), add a custom connector , and paste the server URL
( https://<your-host>/mcp ). Once connected, ask it to load a design and explore
exactly as above. (ChatGPT's connector UI evolves; the constant is: it needs an
HTTPS MCP URL, which --transport streamable-http provides.)
⚠️ The HTTP server has no built-in auth — only expose it over a trusted tunnel,
and prefer short-lived tunnels for local experiments.
Already synthesized? Load the structural Verilog netlist together with the
Liberty library that defines its standard cells, and navigate the gates the same
way as RTL:
"Load the Liberty library pdk/stdcells.lib , then the gate netlist
build/top.v , and tell me what cells top is built from and what drives
data_out ."
Hierarchy, per-cell counts ( get_module_card ), drivers/loads, and logic cones
all work on the netlist; cones stop at the sequential cells. A gate netlist
carries no source line info, so get_source applies to RTL only. A runnable
example lives in examples/ ( stdcells.lib + counter2.v +
gate_level.py ).
Once a design is loaded, your assistant can:
Resolve any signal or instance by hierarchical path (with glob and
did-you-mean suggestions).
Find objects design-wide by pattern.
Show the hierarchy of any module.
Get drivers / loads of a net — the real endpoints, across hierarchy.
Trace logic cones (fan-in / fan-out) and see the register frontier.
Get source — the exact SystemVerilog lines behind any object.
Get a module card — ports, counts, clock/reset at a glance.
Recover design intent — state-machine names, struct fields, parameter
expressions lost during elaboration.
A runnable end-to-end walkthrough lives in examples/ .
Works anywhere najaeda runs (Linux, macOS, Windows)
# from a checkout
python3 -m venv .venv
.venv/bin/pip install -e .
.venv/bin/python -m pytest -q
The full test suite runs against a plain pip install of najaeda — no native
build required. The CVA6 cross-hierarchy cone regression
( tests/test_zzz_cone_cva6.py ) is slow and skips automatically unless a CVA6
snapshot is present.
🐛 Found a bug or have a feature request?
Open an issue on GitHub →
📫 Get in touch: contact@keplertech.io
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
