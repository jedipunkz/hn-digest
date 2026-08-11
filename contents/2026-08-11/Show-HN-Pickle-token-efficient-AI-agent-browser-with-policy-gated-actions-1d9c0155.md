---
source: "https://picklebrowser.com/"
hn_url: "https://news.ycombinator.com/item?id=49260102"
title: "Show HN: Pickle – token efficient AI agent browser with policy-gated actions"
article_title: "Pickle Browser — the agent-first browser"
author: "ismail_h"
captured_at: "2026-08-11T15:50:47Z"
capture_tool: "hn-digest"
hn_id: 49260102
score: 3
comments: 0
posted_at: "2026-08-11T15:43:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Pickle – token efficient AI agent browser with policy-gated actions

- HN: [49260102](https://news.ycombinator.com/item?id=49260102)
- Source: [picklebrowser.com](https://picklebrowser.com/)
- Score: 3
- Comments: 0
- Posted: 2026-08-11T15:43:55Z

## Translation

タイトル: Show HN: Pickle – ポリシーゲート型アクションを備えたトークン効率的な AI エージェント ブラウザー
記事のタイトル: Pickle Browser — エージェントファーストのブラウザ
説明: エージェントに、より安価でプライベートなブラウザを提供します。
HN テキスト: 私は自分の仕事の多くにエージェントを使用していますが、エージェントが私のブラウザにアクセスするたびに、HTML を読み取るために大量のトークンを無駄にすることになるため、エージェント向けに Web ページのコンテンツを簡素化するブラウザを構築してみました。 Pickle はエージェント ブラウザであり、通常のブラウザが持つすべての機能 (タブ、検索履歴、ブックマークなど) を備えていますが、トークンの使用量を削減するために、ページは生の HTML ではなくコンパクトな構造化データとして読み込まれます。興味のある方向けの概要: ポリシーゲート (ブロックされたドメイン、購入などの承認が必要なアクション) であり、すべてのアクションがログに記録されるため、エージェントの動作を常に確認でき、いつでも引き継ぐことができます。モデルの強度に応じて自動ルーティングされるため、弱い/ローカルなモデルとも互換性があります。つまり、小さなローカル モデルは、一度に 1 ステップずつアクションを選択することに制限されているため、アクションが幻覚になったり、要素 ID が構成されたりすることはありません。より強力なモデルでは、複数のステップを先取りして計画できます。トークンの圧縮に関しては、トークンの使用量が約 32 分の 1 に減少していることがわかりました。その他の便利な機能:
- あなたとあなたのエージェントが書き込むことができる共有ノートブック
- ページ要素のメモリはセッション間で保持されるため、エージェントはページを再訪問するたびにボタンやフィールドを再配置する必要がありません。
- セッション間の履歴検索
- ブラウジング モード、ユース ケース (リサーチ、学習、ショッピングなど) に応じてエージェントの動作を推進するさまざまなペルソナ セットアップに API キーは必要なく、ダウンロードして開き、初回起動時にローカル モデルを選択する (または独自のモデルを接続する) だけで済みます。これは、Claude Desktop、Cursor、VS Code、Codex CLI のワンクリック構成です。試してみて、知らせてください

あなたが思うこと！

記事本文:
">
ピクルブラウザ
デモ
モデル
Pickle ブラウザを使用する理由
比較する
接続する
ダウンロード
エージェントファーストのブラウザ
ブラウザに何をすべきかを指示します。それを見てください。
Pickle Browser は検索し、実際のページを開いて読み取り、ソースとともにレポートを返します。これはマシン上で実行され、ウィンドウで監視できるため、すべてのステップを確認し、いつでもハンドルを握ることができます。ローカルで無料で実行されるモデルを使用するか、すでに料金を支払っている Claude または ChatGPT を持ち込んでください。
無料 · アカウントなしでローカルで実行可能 · ページがデバイスから離れることはありません
観察 → 34 要素 · 610 tk
extract_markdown → 1.1k tk (vs 38k raw)
完了 · いつでも引き継ぐことができます
未編集の画面キャプチャ: 検索し、実際のページを開いて読み取り、見つけた内容を書き込みます。その一方で、ユーザーはすべてのステップを監視し、いつでも引き継ぐことができます。
本物の走り、スピードアップ。何も上演または再現されませんでした。
思考が起こる場所を選択するのはあなたです。
Pickle Browser は 1 つの AI 企業と結びついていません。給電方法は 3 つあり、設定でいつでも切り替えることができます。
Ollama エンジンはアプリ内に同梱されています。最初の実行時に、ハードウェアに合わせたサイズのオープン モデル (約 1 GB) が 1 つダウンロードされ、オフラインで動作します。
閲覧中のページを要約したり、データを取り出したり、目の前にあるものについての質問に答えたりするのが得意
モデル Qwen2.5、Llama 3.2、Phi-3、Mistral およびその他の Ollama モデル
ページのコンテンツがデバイスから離れることはありません
率直な答えは、小規模なローカル モデルは、長い複数ステップの調査では信頼性がはるかに低いです。スレッドを失い、話題から逸れてしまう可能性があります。アプリはそれらにラベルを付けて警告し、タスクの途中で切り替えることができます。
料金を支払わずにさらに多くの機能をご希望ですか?現在 2 つのルートが無料で、どちらも 1 回限りのセットアップです。
Ollama Cloud gpt-oss 20B および 120B、qwen3-coder 480B (無料サインイン後)
NVIDIA 無料の API キーを使用して 100 以上のホストされたモデルを構築
実際の複数ステップの調査に十分な速度

任意のラップトップ
これらは他の人のサーバーで実行されるため、そのタスクのページ テキストがそのサーバーに送信されます。履歴、メモ、ログインはローカルに残ります。
すでに Claude または ChatGPT に登録していますか?使ってください。新しく買うものは何もありません。
Claude、GPT、または Gemini の API キー
または、Claude Desktop、Claude Code、Cursor、または VS Code から MCP 経由で接続し、このブラウザを 72 のツールで操作します
クレジットはさらに充実します。生の HTML ではなく、クリーンなページの概要が送信されます。
接続されたエージェントは同じ承認ゲートを受け取ります。購入、送信、および取り消しできないことはすべて、最初にあなたに尋ねます。
他のすべてのエージェント ブラウザは、見えない場所で実行されます。
クラウド エージェント ブラウザは、自動化を実現する開発者向けのインフラストラクチャです。ヘッドレス、リモート、毎回空のセッションから開始されます。 Pickle Browser は逆の賭けをします。これはあなたが使用するブラウザであり、あなたが見ている間にエージェントが操作します。ランニングコストは安くなりますが、それは効率性の副作用であり、切り替える理由ではありません。その理由は、作業は目の前でハードウェア上で行われ、途中で停止できるためです。
ランニングコストが安くなる
ページあたりのトークンが最大 30 倍少なくなります
通常のエージェントは、その内容を確認するためだけに、生の HTML のページ全体 (30,000 以上のトークン) を読み取ります。 Pickle Browser は、代わりに明確で構造化された概要を送信し、各ステップで保存されたトークンを表示します。組み込みのローカル モデルでは、各実行のコストは $0 です。
デフォルトでは非公開
クラウドではなくコンピューター上で動作します
実際のブラウザとログインを使用し、ページのコンテンツがデバイスから離れることはありません。ホストされたブラウザーやアカウントを作成する必要はありません。モデルは同梱されているため、オフラインでも動作します。
モデル、ウィンドウ、電話番号
任意の MCP エージェントに接続します。モデルを選択すると、Pickle Browser は単なるブラウザーになります。開いているすべてのタブが表示され、いつでもハンドルを握ることができ、購入する前に常に質問されます。
があります

現在、エージェントにブラウザを提供する一般的な方法を教えてください。ここでは、それぞれの違いと、Pickle Browser の異なる点を示します。
✓ 組み込まれています · ~ 部分的またはセットアップに依存します · ✗ 何かを諦めなければなりません
ページを読むのに 40,000 トークンはかからないはずです。
最近の 1 つの Web ページは、生の HTML の 30 ～ 60,000 トークンであることがよくあります。 Pickle Browser は、エージェントに構造化された読み取り (見出し、リンク、安定した ID を持つインタラクティブな要素) を提供します。これはそのごく一部であり、保存された内容を正確にログに記録します。節約はスローガンではありません。それらはツールバーの数字です。
サンプルページが読み取られました。アプリのステップごとに実際の節約額が表示されます。
すでに使用しているエージェントを持参してください。
Pickle Browser は Model Context Protocol を話すため、Claude Desktop などのあらゆる MCP クライアントで実際のブラウザを駆動できます。サーバー エントリを 1 つ追加すると、エージェントは参照、読み取り、およびショッピング ツールの完全なセットを取得できます。
Pickle Browser をインストールします。ローカル ランタイムと MCP サーバーがバンドルされています。
Pickle Browser サーバーを追加します。エージェント クライアントの MCP 構成に 1 つのエントリを追加します。
エージェントにブラウズを依頼します。エージェントは、あなたが見ているタブを動かします。
アプリは、クライアントに一致するコピー＆ペーストの設定を接続画面に表示します。
Pickle Browser Search は、ストアがサポートしている場合は Universal Commerce Protocol を介して実際の製品カタログを読み取り、サポートしていない場合は Web 結果をクリーンアップします。エージェントは販売者全体で 1 つのカートを満たします。チェックアウトすると、各店舗の実際のカートに分割されます。
エージェントが必要とするすべてが 1 か所にまとめられています。
ブラウザに組み込まれたチャット ボックスではなく、あなたとあなたのエージェントのための完全なツールキットがすべてあなたのマシン上で実行されます。
エンジンが内蔵されており、最初の実行でオープンモデルを牽引します。その後: キーなし、アカウントなし、オフラインで動作、実行あたり $0。
1 つの編集可能な Markdown ファイルが、そのトーン、ルール、境界を設定します。
詳細な調査、価格の比較、フォームへの入力… が組み込まれており、

独自のものを追加することもできます。
目標を分解し、各ステップを監視し、ソースを含む構造化された成果物を取得します。
エージェントが行ったことをスクリーンショットごとに段階的に確認します。
すべてのステップで、保存されたトークンと生の HTML の送信を示します。測定されたものであり、約束されたものではありません。
UCP + ウェブ経由で買い物をすると、チェックアウトは販売者ごとに実際のカートに分割されます。
エージェントが使用できるように、Gmail、Slack、Notion、または独自の MCP サーバーを接続します。
ダウンロードしてエージェントにキーを渡します。
無料、アカウントなし。最初の実行時にワンクリックでローカル AI をセットアップするか、すでに使用しているモデルを接続します。
1 つのコードベース、3 つのプラットフォームすべて · ～120 MB · 初期ビルド
バンドルされたローカル モデルは、閲覧しているページについての要約、抽出、回答など、単一ページの迅速な作業に最適です。小規模なモデルは、長期の複数ステップの調査でははるかにエラーが発生しやすいため、精度が重要な場合は、より大きなモデルまたはクラウド モデルを接続します。
Windows ビルドのアップデートと修正については、電子メールを使用してお知らせします。次にダウンロードリンクが表示されます。スパムはありません。
ありがとう！ Windows のダウンロードの準備が完了しました。
始まりませんでしたか？上のボタンを使用してください。
あなたの電子メールは、ダウンロードとアップデートの通知を送信するためにのみ使用されます。当社のプライバシー ポリシーをご覧ください。
質問、バグ、または早期アクセスが必要ですか?
私たちはすべてを読みます。何が壊れたのか、こうなってほしいと思っていること、自動化しようとしているものを教えてください。
問題を報告する場合は、[設定] から OS とアプリのバージョンを含めてください。 API キーやページのコンテンツを貼り付けないでください。

## Original Extract

Give your agent a browser that is cheaper, private, and yours.

I use agents for a lot of my own work, but whenever they access my browser they end up wasting a lot of tokens reading HTML, so I tried building a browser that simplifies webpage content for them. Pickle is an agent browser and it has all the features a regular browser has (tabs, search history, bookmarks, etc.), but pages load as compact structured data instead of raw HTML to reduce token usage. Overview for those curious: It's policy-gated (blocked domains, actions that need approval like purchases) and every action is logged, so you're always able to see what your agent is doing and can take over at any time. It's also compatible with weaker/local models as it auto-routes by model strength. i.e. small local models are limited to picking actions one step at a time so they don't hallucinate actions or make up element IDs. Stronger models can plan multiple steps ahead. On the token compression side, I've been seeing around 32x less token usage. Other useful features:
- shared notebook that you and your agents can write to
- memory of page elements are preserved between sessions, so agents don't have to relocate buttons/fields every time they revisit a page
- cross-session history search
- Browsing modes, different personas that drive agent behaviour depending on your use case (e.g. research, study, shopping) You don't need an API key to set it up and can just download it, open it, and pick a local model on first launch (or connect your own). It's a one click config for Claude Desktop, Cursor, VS Code, and Codex CLI. Give it a go, and let me know what you think!

">
Pickle Browser
Demo
Models
Why Pickle Browser
Compare
Connect
Download
The agent-first browser
Tell your browser what to do. Watch it do it.
Pickle Browser searches, opens real pages, reads them, and reports back with its sources. It runs on your machine, in a window you can watch , so you see every step and can take the wheel at any point. Use a model that runs locally for free, or bring the Claude or ChatGPT you already pay for.
Free · runs locally with no account · your pages never leave your device
observe → 34 elements · 610 tk
extract_markdown → 1.1k tk (vs 38k raw)
done · you can take over anytime
Unedited screen capture: it searches, opens real pages, reads them, and writes up what it found, while you watch every step and can take over at any point.
Real run, sped up. Nothing staged or re-enacted.
You choose where the thinking happens.
Pickle Browser is not tied to one AI company. Three ways to power it, and you can switch at any time in Settings.
The Ollama engine ships inside the app. On first run it downloads one open model sized to your hardware (about 1 GB), then works offline.
Good at summarizing the page you're on, pulling out data, answering questions about what's in front of you
Models Qwen2.5, Llama 3.2, Phi-3, Mistral and other Ollama models
Page content never leaves your device
Straight answer: small local models are far less reliable on long multi-step research. They can lose the thread and go off topic. The app labels them and warns you, and you can switch mid-task.
Want more capability without paying? Two routes are free today, and both are one-time setup.
Ollama Cloud gpt-oss 20B and 120B, qwen3-coder 480B, after a free sign-in
NVIDIA Build 100+ hosted models with a free API key
Fast enough for real multi-step research on any laptop
These run on someone else's servers, so the page text for that task is sent to them. Your history, notes and logins still stay local.
Already subscribed to Claude or ChatGPT? Use it. Nothing new to buy.
Your API key for Claude, GPT or Gemini
Or connect over MCP from Claude Desktop, Claude Code, Cursor or VS Code and drive this browser with 72 tools
Your credits go further: it sends a clean page summary instead of raw HTML
Connected agents get the same approval gates. Buying, sending and anything irreversible still asks you first.
Every other agent browser runs somewhere you can't see.
Cloud agent browsers are infrastructure for developers shipping automation: headless, remote, and starting from a blank session every time. Pickle Browser takes the opposite bet. It's a browser you use, that an agent drives while you watch. It does cost less to run, but that's a side effect of being efficient, not the reason to switch. The reason is that the work happens in front of you, on your hardware, and you can stop it mid-step.
Cheaper to run
Up to 30× fewer tokens per page
A normal agent reads a whole page of raw HTML — 30,000+ tokens — just to see what's on it. Pickle Browser sends a clean, structured summary instead, and shows the tokens saved on every step. On the built-in local model, each run costs $0 .
Private by default
Runs on your computer, not a cloud
It uses your real browser and logins , and page content never leaves your device. No hosted browser, no account to create — and because a model ships in the box, it even works offline.
Your model, your window, your call
Connect any MCP agent — you choose the model, Pickle Browser is just the browser. You see every tab it opens, can grab the wheel any time, and it always asks before it buys .
There are four common ways to give an agent a browser today. Here's where each one leaves you — and where Pickle Browser is different.
✓ built in · ~ partial or depends on setup · ✗ not without giving something up
Reading a page shouldn't cost 40,000 tokens.
A single modern web page is often 30–60k tokens of raw HTML. Pickle Browser gives the agent a structured read — headings, links, and interactive elements with stable ids — that's a small fraction of that, and logs exactly what it saved. The savings aren't a slogan; they're a number in the toolbar.
Example page read. Actual savings shown per step in the app.
Bring the agent you already use.
Pickle Browser speaks the Model Context Protocol, so any MCP client — like Claude Desktop — can drive your real browser. Add one server entry and the agent gets a full set of browse, read, and shop tools.
Install Pickle Browser — the local runtime and MCP server come bundled.
Add the Pickle Browser server — one entry in your agent client's MCP config.
Ask your agent to browse — it drives the tab you're watching.
The app shows a copy-paste config on its Connect screen, matched to your client.
Pickle Browser Search reads real product catalogs over the Universal Commerce Protocol when a store supports it, and clean web results when it doesn't. Your agent fills one cart across merchants; checkout splits it into a real cart at each store.
Everything the agent needs, in one place.
Not a chat box bolted onto a browser — a full toolkit for you and your agent, all running on your machine.
The engine is built in and pulls an open model on first run. After that: no keys, no account, works offline, $0 per run.
One editable Markdown file sets its tone, rules, and boundaries.
Deep Research, Compare Prices, Fill Form… built in, and you can add your own.
Decompose a goal, watch each step, get a structured deliverable with sources.
Step through what the agent did, screenshot by screenshot.
Every step shows tokens saved vs sending raw HTML — measured, not promised.
Shop over UCP + the web, checkout splits into a real cart per merchant.
Plug in Gmail, Slack, Notion, or your own MCP server for the agent to use.
Download and hand your agent the keys.
Free, no account. Set up local AI in one click on first run, or connect the model you already use.
One codebase, all three platforms · ~120 MB · early build
The bundled local model is best for quick single-page work: summarizing, extracting, and answering about the page you are on. Small models are far more error-prone on long multi-step research, so connect a larger or cloud model when accuracy matters.
We use your email to tell you about updates and fixes for the Windows build. Your download link appears next. No spam.
Thanks! Your Windows download is ready.
Didn't start? Use the button above.
Your email is only used to send the download and update notices. See our privacy policy .
Questions, bugs, or want early access?
We read everything. Tell us what broke, what you wish it did, or what you are trying to automate.
If you are reporting a problem, include your OS and the app version from Settings. Please do not paste API keys or page content.
