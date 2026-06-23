---
source: "https://github.com/cloudkj/llayer"
hn_url: "https://news.ycombinator.com/item?id=48645683"
title: "Show HN: AI Agents the Unix Way – built with bash, curl, and jq"
article_title: "GitHub - cloudkj/llayer: AI agents the Unix way · GitHub"
author: "cloudkj"
captured_at: "2026-06-23T15:00:57Z"
capture_tool: "hn-digest"
hn_id: 48645683
score: 1
comments: 0
posted_at: "2026-06-23T14:38:29Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Agents the Unix Way – built with bash, curl, and jq

- HN: [48645683](https://news.ycombinator.com/item?id=48645683)
- Source: [github.com](https://github.com/cloudkj/llayer)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T14:38:29Z

## Translation

タイトル: Show HN: Unix 方式の AI エージェント – bash、curl、および jq で構築
記事のタイトル: GitHub - Cloudkj/llayer: Unix 方式の AI エージェント · GitHub
説明: Unix 方式の AI エージェント。 GitHub でアカウントを作成して、cloudkj/llayer の開発に貢献してください。
HN テキスト: 教育演習でローカル モデルをいじり、エージェントのセットアップに取り組んでいる間、私はウサギの穴に迷い込みました。コマンド ラインのビルディング ブロックのみを使用し、可能な限り依存関係を取り除き、カスタム エージェント ループをどこまで構築できるかを確認するためです。パイプ、テキスト ストリーム、追加のみのログ、および標準のコマンド ライン コンポーネントを使用すると、かなりのところまで到達できることがわかりました。これらの概念は、古典的な Unix の哲学と非常によく一致しています。エージェントは、少数の小さなプログラムで構成されるラッパーであり、エージェント ループのさまざまな段階を検査、フィルタリング、リダイレクト、監査するためのさまざまなツールを柔軟に挿入できるようにする必要があります。このプロジェクトは現状では概念実証ですが、ツール呼び出しのサポートが十分に盛り込まれており、理論上は無期限に拡張できるはずです。そうは言っても、ツール呼び出しの高度さは基盤となるモデルによって主に制限されているように見えます。これまでのところ、ツールとしてはある程度の成功を収めた軽量のローカル モデル (llama3.2:1b など) のみを実験してきました。それにも関わらず、私はこれをここで共有したかったので、他の人がこれを構築したり拡張したりするのに十分興味深いと思うかどうかを知りたいと思っています。

記事本文:
GitHub - Cloudkj/llayer: Unix 方式の AI エージェント · GitHub
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
クラウドkj
/
レイヤ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 他のアクションを開く men

u フォルダーとファイル
18 コミット 18 コミット テスト testing .gitignore .gitignore ライセンス ライセンス README.md README.md エージェント エージェント compose.yaml compose.yaml ll-context ll-context ll-dispatch ll-dispatch ll-eval ll-eval ll-print ll-print ll-read ll-read tools.json tools.json すべてのファイルを表示リポジトリ ファイルのナビゲーション
llayer - Unix 方式の AI エージェント
bash 、curl 、および jq で構築されたフレームワーク不要の AI エージェント
llayer は、Unix の哲学を大規模言語モデルのオーケストレーションに適用する実験です。すべてが処理されます
標準のテキスト パイプを介してインターフェイスする一連の小型の単一目的コマンドを介して、REPL スタイルのエージェントを生成します
ループ。
エージェント フレームワークは状態が大きく、推論が難しい場合があります。 llayer は、エージェントのライフサイクルを基本的なものに分解します。
State & Memory : 追加専用の .jsonl 履歴ファイル ( .bash_history など)。
Context Window : 機能的な jq ストリーム リデューサ。
エージェント ループ : 標準の bash while ループ。
アーキテクチャ全体がステートレス テキスト パイプラインであるため、次のような優れた機能が利用可能になります。
⏱️ タイムトラベル デバッグ : 履歴ファイルをスライスすることでエージェントのメモリを即座に巻き戻します (例: head -n 20 .llayer_history )。
🛠️ ゼロ抽象化ツール: 標準出力が直接エージェントの入力となる純粋な Bash 関数を使用してツールを追加します。
🚀 ネイティブ パイプライン: 標準の POSIX ツールを使用してエージェントのストリームをインターセプトします。 PII データがモデルにヒットする前に grep でフィルタリングするか、 pv を使用してトークン ストリーミング速度をネイティブにベンチマークします。
開始するには: このリポジトリのクローンを作成し、docker compose でローカル モデル サーバー (Ollama) を起動してから、agent を呼び出します。
% docker 構成 --detach
...
✔ コンテナ llayer-ollama-1 が開始されました
% ./エージェント
> こんにちは、おはようございます
あなたもおはようございます！一日の始まりはいかがですか？
あるいは、個別の com を試してください。

マンズ。たとえば、ステートレスでコンテキストフリーの呼び出しは、単に次のチェーンです。
コマンド:
% echo "こんにちは、世界" | ./ll-read | ./ll-context | ./ll-eval | ./ll-print
こんにちは！お会い出来て嬉しいです。何かお手伝いできることはありますか、それともお話しませんか?
ステートフル エージェント: REPL
グラフLR
ユーザー
サブグラフの読み取り
読んでみます
歴史
ll-コンテキスト
終わり
サブグラフ EVAL
ll-評価
ll-派遣
モデル
ツール
終わり
サブグラフ PRINT
イルプリント
終わり
ユーザー -->|読む|読んでみます
読みます -->|追加|歴史
歴史 -->|パイプ| ll-コンテキスト
ll-context -->|パイプ| ll-評価
ll-eval -->|追加|歴史
ll-eval -->|パイプ| ll-派遣
ll-dispatch <-->|call|ツール
ll-dispatch -->|追加|歴史
ll-eval -->|パイプ|イルプリント
ll-eval <-->|http|モデル
ll-print -->|print|ユーザー
スタイル ll-read fill:#4A90E2
スタイル ll-context 塗りつぶし:#4A90E2
スタイル ll-eval 塗りつぶし:#4A90E2
スタイル ll-ディスパッチフィル:#4A90E2
スタイル ll-print 塗りつぶし:#4A90E2
読み込み中
エージェント スクリプトは、スタンドアロン コンポーネントを組み合わせて、AI エージェントによく似た読み取り-評価-印刷ループ (REPL) を形成します。
ユーザー入力を読み取り、対応するイベントを履歴ファイルに追加します。
履歴から ll-context を構築してモデル コンテキストを生成し、モデルを ll-eval します。
必要に応じて、サポートされているツールに ll-dispatch し、結果を履歴に追加します。ステップ2を繰り返します。
モデル出力をイベントとして履歴に追加し、ll-print してユーザーにメッセージを表示します。
エージェントの特定のコンポーネントの呼び出しは簡単です。例には、モデルに渡されたコンテキストの検査と追加が含まれます。
(cat .llayer_history && echo "PI のいくつかの桁を出力" | ./ll-read) | ./ll-context | jq -c ' .[] '
または、履歴からエージェント メッセージを再生します。
猫 .llayer_history | ./ll-print --debug
または、ストリーミングされたモデル出力を直接検査します。
echo "ピン! ポンと返信するだけ" | ./ll-read | ./ll-context | ./ll-評価
または下流に配管してください

m ツールは、モデルが応答をストリーミングし返す速度を測定し、応答を出力する前にすべての出力をバッファーします。
echo "ウッドチャックはどれくらいの木材をチャッキングできますか?" | ./ll-read | ./ll-context | ./ll-eval | pv --ラインモード |スポンジ | ./ll-print
実装
追加専用の履歴ファイルには、すべての状態が保存されます。各行は、ユーザー入力、モデルによって発行されたトークン、完了したメッセージ、またはツールの呼び出し/結果のいずれかを記述するイベント JSON オブジェクトです。この設計の動機と目標:
不変性: 個々のトークンに至るまで、すべてのイベントは監査、デバッグ、および再生可能性のために保存されます。
シンプルさ: 追加専用のテキストを使用して状態を保存することは堅牢であり、ミニマリストの哲学に沿っています。
構成可能性: ダウンストリーム ツールは、状態を変更せずにイベント ストリームを消費、フィルター、変換できます。
ll-context は、正規のイベント履歴を圧縮して、各モデル呼び出しのコンテキストを構築します。その主な目的は次のとおりです。
スコープ: このコマンドは、履歴を関連イベントにフィルターし、トークンを上位レベルのメッセージにグループ化し、構成可能なヒューリスティックを適用します (例: 過去 N ターンを保持する、ツール呼び出しペイロードを削除する、トークンを単一のアシスタント メッセージに折りたたむ) ため、モデルは簡潔なコンテキストを受け取ります。
削減: 元の履歴が書き換えられたり削除されたりすることはありません。このコマンドは、ユーザー定義のウィンドウ内のイベントから派生したフィルター処理されたモデルに適したシーケンスに履歴を縮小します。
個々のコンポーネントは、各行に type 、source 、および payload が含まれる最小限の明示的な JSONL 形式に従う DSL を利用します。基本的なイベント スキーマは次のとおりです。
{ "タイプ" : " エラー " 、 "ソース" : " システム " 、 "ペイロード" : { "テスト" : " ... " }}
{ "タイプ" : " メッセージ " 、 "ソース" : " ユーザー " 、 "ペイロード" : { "テキスト" : " ... " }}
{ "タイプ" : " message_complete " 、 "ソース" :

" システム " 、 "ペイロード" : {}}
{ "タイプ" : " トークン " 、 "ソース" : " アシスタント " 、 "ペイロード" : { "テキスト" : " ... " }}
{ "タイプ" : " ツールコール " 、 "ソース" : " アシスタント " 、 "ペイロード" : {}}
{ "タイプ" : " ツール結果 " 、 "ソース" : " ツール " 、 "ペイロード" : {}}
について
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

AI agents the Unix way. Contribute to cloudkj/llayer development by creating an account on GitHub.

While working on an educational exercise tinkering with local models and trying my hand at setting up agents, I went down a rabbit hole: to see how far I could build a custom agent loop using exclusively command-line building blocks and stripping out dependencies wherever possible. It turns out you can get pretty far with pipes, text streams, append only logs, and standard command-line components - concepts pretty well aligned with classic Unix philosophy. The agent is a wrapper composed of a handful of smaller programs, which should allow for flexibly injecting various tool to inspect, filter, redirect, and audit different stages of the agent loop. This project as it stands is a proof-of-concept, but packs enough punch with tool calling support - which in theory should make it indefinitely extensible. With that said, it does appear that the sophistication of tool calling is largely limited by the underlying model, and so far I’ve only experimented with lightweight local models (e.g. llama3.2:1b) that have modest success for tools. Nevertheless, I wanted to share this here and am curious to see if others find it interesting enough to build upon or extend!

GitHub - cloudkj/llayer: AI agents the Unix way · GitHub
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
cloudkj
/
llayer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
18 Commits 18 Commits tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md agent agent compose.yaml compose.yaml ll-context ll-context ll-dispatch ll-dispatch ll-eval ll-eval ll-print ll-print ll-read ll-read tools.json tools.json View all files Repository files navigation
llayer - AI agents the Unix way
Framework-free AI agents built with bash , curl , and jq
llayer is an experiment in applying the Unix philosophy to large language model orchestration. Everything is handled
through a suite of small, single-purpose commands interfacing over standard text pipes to produce a REPL-style agent
loop.
Agent frameworks can be state-heavy and hard to reason about. llayer deconstructs the agent lifecycle into fundamentals:
State & Memory : An append-only .jsonl history file (like .bash_history ).
Context Window : A functional jq stream reducer.
Agent Loop : A standard bash while loop.
Because the entire architecture is a stateless text pipeline, it unlocks neat capabilities, such as:
⏱️ Time-Travel Debugging : instantly rewind the agent's memory by slicing the history file (e.g., head -n 20 .llayer_history ).
🛠️ Zero Abstraction Tooling : add tools using pure Bash functions where stdout directly becomes the agent's input.
🚀 Native Pipelining : intercept the agent's stream with standard POSIX tools - e.g. filter PII data with grep before it hits the model, or natively benchmark token streaming speeds using pv .
To get started: clone this repo, start a local model server (Ollama) with docker compose, then call agent :
% docker compose up --detach
...
✔ Container llayer-ollama-1 Started
% ./agent
> Hello there, good morning
Good morning to you as well ! How ' s your day starting out so far?
Alternatively, try out the individual commands. For example, a stateless, context-free call is simply a chain of
commands:
% echo " Hello, world " | ./ll-read | ./ll-context | ./ll-eval | ./ll-print
Hello ! It ' s nice to meet you. Is there something I can help you with or would you like to chat?
The Stateful Agent: a REPL
graph LR
User
subgraph READ
ll-read
History
ll-context
end
subgraph EVAL
ll-eval
ll-dispatch
Model
Tools
end
subgraph PRINT
ll-print
end
User -->|read| ll-read
ll-read -->|append| History
History -->|pipe| ll-context
ll-context -->|pipe| ll-eval
ll-eval -->|append| History
ll-eval -->|pipe| ll-dispatch
ll-dispatch <-->|call|Tools
ll-dispatch -->|append| History
ll-eval -->|pipe| ll-print
ll-eval <-->|http| Model
ll-print -->|print| User
style ll-read fill:#4A90E2
style ll-context fill:#4A90E2
style ll-eval fill:#4A90E2
style ll-dispatch fill:#4A90E2
style ll-print fill:#4A90E2
Loading
The agent script combines the standalone components to form a read-eval-print loop (REPL) that largely resembles an AI agent:
ll-read user input and append a corresponding event to the history file.
Build ll-context from history to produce the model context and ll-eval the model.
If necessary, ll-dispatch to supported tools and append the result to the history; repeat step 2.
Append model output as events to history then ll-print to display messages to the user.
Caling specific components of the agent is straightforward. Examples include inspecting and adding to the context passed into the model:
(cat .llayer_history && echo " Print some digits of PI " | ./ll-read) | ./ll-context | jq -c ' .[] '
Or replaying agent messages from history:
cat .llayer_history | ./ll-print --debug
Or directly inspect streamed model outputs:
echo " ping! just reply pong " | ./ll-read | ./ll-context | ./ll-eval
Or pipe to a downstream tool to measure how quickly the model is streaming responses back, and buffer all of the output before printing the responses:
echo " How much wood could a woodchuck chuck? " | ./ll-read | ./ll-context | ./ll-eval | pv --line-mode | sponge | ./ll-print
Implementation
An append-only history file stores all of the state. Each line is an event JSON object describing either a user input, a token emitted by the model, a completed message, or a tool call/result. Motivations and goals of this design:
Immutability: all events, down to individual tokens, are preserved for auditing, debugging, and replayability.
Simplicity: using append-only text to store state is robust and aligns with the minimalist philosophy.
Composability: downstream tools can consume, filter, and transform the event stream without modifying state.
ll-context compacts the canonical event history to build the context for each model call. Its main purposes are:
Scoping: the command filters history down to relevant events, groups tokens into higher-level messages, and applies configurable heuristics (e.g. keep last N turns, strip tool-call payloads, collapse tokens into a single assistant message) so the model receives concise context.
Reduction: the original history is never rewritten or deleted; the command reduces the history down to a filtered, model-friendly sequence derived from events that fall within a user-defined window.
The individual components utilize a DSL that follows minimal, explicit JSONL shapes where each line contains a type , source , and payload . The basic event schema is as follows:
{ "type" : " error " , "source" : " system " , "payload" : { "test" : " ... " }}
{ "type" : " message " , "source" : " user " , "payload" : { "text" : " ... " }}
{ "type" : " message_complete " , "source" : " system " , "payload" : {}}
{ "type" : " token " , "source" : " assistant " , "payload" : { "text" : " ... " }}
{ "type" : " tool_call " , "source" : " assistant " , "payload" : {}}
{ "type" : " tool_result " , "source" : " tool " , "payload" : {}}
About
There was an error while loading. Please reload this page .
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
