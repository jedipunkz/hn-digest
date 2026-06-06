---
source: "https://marketplace.visualstudio.com/items?itemName=christianalfoni.stackoverflow-ai"
hn_url: "https://news.ycombinator.com/item?id=48422687"
title: "Let us replace community with AI"
article_title: "Stack Overflow AI - Visual Studio Marketplace"
author: "christianalfoni"
captured_at: "2026-06-06T09:45:17Z"
capture_tool: "hn-digest"
hn_id: 48422687
score: 2
comments: 2
posted_at: "2026-06-06T08:24:55Z"
tags:
  - hacker-news
  - translated
---

# Let us replace community with AI

- HN: [48422687](https://news.ycombinator.com/item?id=48422687)
- Source: [marketplace.visualstudio.com](https://marketplace.visualstudio.com/items?itemName=christianalfoni.stackoverflow-ai)
- Score: 2
- Comments: 2
- Posted: 2026-06-06T08:24:55Z

## Translation

タイトル: コミュニティを AI に置き換えましょう
記事のタイトル: Stack Overflow AI - Visual Studio Marketplace
説明: Visual Studio Code の拡張機能 - インラインでコーディングの質問をし、スタック オーバーフロー スレッドの形式で回答を取得します。凝縮された、複数の競合する回答のうち 1 つが受け入れられます。

記事本文:
Stack Overflow AI - Visual Studio マーケットプレイス
コンテンツにスキップ
|マーケットプレイス
サインイン
Visual Studio Code > 教育 > Stack Overflow AI Visual Studio Code は初めてですか?今すぐ入手してください。スタックオーバーフローAI
チャットボットの散文の壁ではなく、スタック オーバーフロー スレッドの形式と精神 (簡潔で複数の競合する回答、1 つが受け入れられる) でコーディングの質問に答える VS Code 拡張機能。
考え方: SO フォーマットは簡潔さを強制する機能です。 AI にそのフォーマットを実行させると、無料で規律を得ることができます。複数の回答があると、自信を持って 1 回回答することで失った判断ステップが回復します。
これは何かを解決しようとする拡張機能ではありません。これはむしろ芸術作品です。私たちの工芸品に起こっていることに対する不快感を表現できる唯一の方法です。
私たちは皆、目先の生産性向上のために AI を使用してコードを生成するというプレッシャーを感じています。それに伴い、認知的負債が増大し、コンテキストが絶え間なく切り替わり、冗長な計画は誰も実際には読まなくなります。それは、人々の協力が減り、ものづくりの喜びが減り、支援するコミュニティが減った世界を生み出します。
私たちはかつて、自分たちの作成したコードと、前進を助けてくれるコミュニティの間を行き来するリズムを持っていました。 AI が生成した TikTok をドゥームスクロールしながら AI が生成したコードを待つのは、とても暗いことですが、それはほぼ現実になりました。
行き詰まった場所にカーソルを置き (またはコードを選択)、Stack Overflow AI: 選択について尋ねる ( cmd+alt+a / ctrl+alt+a 、または右クリック) を実行します。
スタック オーバーフロー スレッドのようなスタイルのサイド パネルが開きます。ライブ調査ログには、エージェントがファイルを読み取り、タイプを grep し、Web を検索していることが示されます。
スレッドが表示されます: 質問、次に投票数を含むいくつかの競合する回答、および 1 つが承認されました ✓。コード ブロックにカーソルを合わせると、コード ブロックがコピーされます。
ボンネットの下でクロードを実行します

読み取り専用ツールセット (Read、Grep、Glob、WebSearch、WebFetch - 他のツールは拒否されます) と SDK のネイティブ json_schema 構造化出力を備えたエージェント SDK なので、答えは実際のコードに基づいており、常にきれいに解析されます。
クロードコードをインストールする必要があります。この拡張機能は、エージェント SDK の JavaScript のみを同梱し、(プラットフォームごとに 225 MB のバイナリをバンドルするのではなく) 既存の claude 実行可能ファイルを駆動します。これは、PATH および一般的なインストール場所で自動的に検出されます。必要に応じて、stackoverflowAI.claudeCodePath でオーバーライドします。認証には既存のクロード コード ログインが使用されます。
npmインストール
npm実行コンパイル
次に、VS Code で F5 キーを押して拡張機能開発ホストを起動し、任意のプロジェクトを開いてコマンドをトリガーします。
VS Codeを起動せずに反復する
スモーク.mjs は、偽のコンテキストに対してエージェント ドライバーを直接実行し、スレッドを端末に出力します。これは、プロンプトを調整する迅速な方法です。
src/agent.ts — Agent SDK を駆動します: SO 形式のシステム プロンプト、読み取り専用ツール ゲート、構造化出力、進行状況イベント。
src/extension.ts — コマンド。エディターのコンテキストをキャプチャし、Web ビューを管理し、クリップボードへのコピーを処理します。
src/webview.ts — Webview HTML シェル (CSP + リソース URI) を構築します。
media/styles.css — Stack Overflow スタイルのテーマ対応 UI。
media/main.js — WebView ロジック: メッセージ処理、小さな Markdown レンダラー、軽量の JS/TS 構文ハイライター。

## Original Extract

Extension for Visual Studio Code - Ask coding questions inline and get answers in the format of a Stack Overflow thread — condensed, multiple competing answers, one accepted.

Stack Overflow AI - Visual Studio Marketplace
Skip to content
| Marketplace
Sign in
Visual Studio Code > Education > Stack Overflow AI New to Visual Studio Code? Get it now. Stack Overflow AI
A VS Code extension that answers your coding questions in the format and spirit of a Stack Overflow thread — terse, multiple competing answers, one accepted — instead of a wall of chatbot prose.
The idea: the SO format is a forcing function for brevity. Make the AI perform that format and you get the discipline for free. Multiple answers restore the judgment step you lose with a single confident reply.
This is not an extension trying to solve anything. It is more of a piece of art — the only way I can express my discomfort with what is happening to our craft.
We are all feeling the pressure of generating code with AI for short-sighted productivity gains. And with that, the increase of cognitive debt, constant context switching, and verbose plans nobody really reads. It creates a world with less human collaboration, less joy of crafting, and fewer supporting communities.
We used to have a rhythm of moving between our crafted code and a community that helped us move forward. Waiting for AI-generated code while doom-scrolling AI-generated TikTok is as dark as it gets — but it has very much become a reality.
Put your cursor (or select code) where you're stuck and run Stack Overflow AI: Ask about selection ( cmd+alt+a / ctrl+alt+a , or right-click).
A side panel opens, styled like a Stack Overflow thread. A live research log shows the agent reading your files, grepping types, and searching the web.
You get a thread: the question, then several competing answers with vote counts and one accepted ✓. Hover any code block to Copy it.
Under the hood it runs the Claude Agent SDK with a read-only toolset (Read, Grep, Glob, WebSearch, WebFetch — any other tool is denied) and the SDK's native json_schema structured output, so the answers are grounded in your actual code and always parse cleanly.
Claude Code must be installed. The extension ships only the Agent SDK's JavaScript and drives the claude executable you already have (rather than bundling a 225 MB per-platform binary). It's auto-discovered on your PATH and in common install locations; override with stackoverflowAI.claudeCodePath if needed. Authentication uses your existing Claude Code login.
npm install
npm run compile
Then press F5 in VS Code to launch the Extension Development Host, open any project, and trigger the command.
Iterating without launching VS Code
smoke.mjs runs the agent driver directly against a fake context and prints the thread to the terminal — fast way to tune the prompt:
src/agent.ts — drives the Agent SDK: SO-format system prompt, read-only tool gate, structured output, progress events.
src/extension.ts — command, captures editor context, manages the webview, handles copy-to-clipboard.
src/webview.ts — builds the webview HTML shell (CSP + resource URIs).
media/styles.css — the Stack Overflow-styled, theme-aware UI.
media/main.js — webview logic: message handling, tiny Markdown renderer, lightweight JS/TS syntax highlighter.
