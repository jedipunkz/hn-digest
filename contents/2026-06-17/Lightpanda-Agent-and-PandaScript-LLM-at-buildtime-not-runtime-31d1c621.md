---
source: "https://lightpanda.io/blog/posts/introducing-lightpanda-agent-and-pandascript"
hn_url: "https://news.ycombinator.com/item?id=48570753"
title: "Lightpanda Agent and PandaScript – LLM at buildtime, not runtime"
article_title: "Introducing Lightpanda Agent and PandaScript: LLM at buildtime not runtime - Blog | Lightpanda"
author: "fbouvier"
captured_at: "2026-06-17T16:23:26Z"
capture_tool: "hn-digest"
hn_id: 48570753
score: 8
comments: 0
posted_at: "2026-06-17T14:05:29Z"
tags:
  - hacker-news
  - translated
---

# Lightpanda Agent and PandaScript – LLM at buildtime, not runtime

- HN: [48570753](https://news.ycombinator.com/item?id=48570753)
- Source: [lightpanda.io](https://lightpanda.io/blog/posts/introducing-lightpanda-agent-and-pandascript)
- Score: 8
- Comments: 0
- Posted: 2026-06-17T14:05:29Z

## Translation

タイトル: Lightpanda Agent と PandaScript – 実行時ではなくビルド時の LLM
記事のタイトル: Lightpanda Agent と PandaScript の紹介: 実行時ではなくビルド時の LLM - ブログ |ライトパンダ
説明: Lightpanda にエージェントが組み込まれました。自然言語で話しかけると、信頼性の高いスクリプトが返されます。修正が必要な場合にのみ、LLM を呼び出すだけで、そのスクリプトを永久に実行します。

記事本文:
Lightpanda Agent と PandaScript の紹介: 実行時ではなくビルド時の LLM - ブログ |ライトパンダのロゴアイコンDiscord
ナビゲーションの切り替え 戻る Lightpanda エージェントと PandaScript の紹介: 実行時ではなくビルド時の LLM
IconLinkedin アナウンス自動化エージェント 2026 年 6 月 17 日水曜日 TL;DR
現在のブラウザ エージェントは、Chrome、CDP という 4 つの個別のツールが相互に接続されています。
ライブラリ、LLM、およびエージェント フレームワーク。私たちのネイティブエージェントはその4つを崩壊させます
レイヤーを 1 つのバイナリにまとめます。
自然言語で話しかけると、実際のブラウザに対して動作します。
「この Web サイトにアクセスする」、「ログインする」、「このデータを抽出する」、「それを教えて」: 言語は
ブラウザの新しいネイティブ インターフェイス。もう一度再生したい場合は、
セッションでは、再現可能なスクリプトが渡されます。 CDP も重いブラウザ サーバーも必要ありません。
複雑なセットアップや実行時の LLM は必要ありません。すべてが Lightpanda に組み込まれています。
ブラウザ エージェントは 4 つのツールです。それは 1 つである必要があります。
今すぐブラウザ エージェントを実行するには、Chrome をインストールし、CDP 経由でそれを駆動します。
人形遣いまたは劇作家、LLM に接続して意思決定を行い、全体をラップします。
エージェント フレームワーク内で呼び出しを調整するもの。それは動作します。そうですね、次のようなものです。
よくある「ページをロードしていくつか読む」という動作には、多くの可動部分が含まれます。
要素がオフになっています。」
これらの各レイヤーは、人間が直面するブラウザとブラウザの間を変換するために存在します。
それを運転したいマシン。 CDP (Chrome DevTools プロトコル) は次のように構築されました。
開発者は実行中のブラウザを外部から検査できますが、プログラムは検査できません。
1つを操作できます。また、LLM を接続するには、MCP から CDP へのレイヤーが必要です。エージェント
モデルによって駆動されることを意図されていないツールの周囲にモデルをラップします。あなた
すべてのレイヤーで翻訳税を支払っています。
私たちが 4 年前に Lightpanda を始めたとき、次のような疑問がありました。
ブラウザは if y のように見えます

人間が見るのではなく、機械のためにそれを構築しました。
画面？」 lightpanda エージェントは、エージェント スタックに適用されるベットです。それは一つです
ブラウザ、ランタイム、エージェントを含むバイナリ。
ブラウザは、 lightpandaserve と lightpanda fetch の背後にある同じエンジンです。Web ページをロードし、JavaScript を実行し、DOM を処理します。
ランタイムは、ネイティブ ツールの小さなセット ( goto 、 click 、 fill 、
extract 、evaluate、search ) を使用すると、ブラウザー (スラッシュ コマンド) を操作できるようになります。の
リクエストを読み取り、ツールを選択する LLM はオプションです。それは反対します
Anthropic、OpenAI、Gemini、Hugging Face、またはローカル Ollama、またはキーがまったくない場合
スラッシュのみモード。
LLM は実行時ではなくビルド時に実行されます
これは、他のすべてのものはぶら下がっているという考えです。他のすべてのブラウザ エージェントは
ステップごとにモデルを呼び出すブラック ボックス。 lightpanda エージェントを使用すると、モデルは
タスクを特定し、その作業をコードとしてキャプチャし、スクリプトを生成します（
これを PandaScript と呼びます)。これは再現可能で決定的です。リプレイすると
そうすれば、モデルは必要なくなり、LLM がループから外されます。
再生時に次のアクションとの間にモデル呼び出しが存在することはありません。
ホストする Chrome プロセスもセットアップする NodeJS/Python 環境も必要ありません。
劇作家/人形師が記述するコード。スクリプトを Lightpanda に渡すだけです
バイナリなので、実行はエンジンがページを駆動する速度によってのみ制限されます。そして
再生時には非決定的なものは何も実行されないため、同じスクリプトが
毎回同じ結果。ファイルを書き込むためにモデルに 1 回支払います。その後、
それはあなたが所有するプレーンな JavaScript です。
PandaScript は次のようになります。このスクリプトはトップ 3 ハッカーを獲得します
ニュース記事を表示し、各スレッドにアクセスしてトップのコメントを探します。
const HN_ORIGIN = "https://news.ycombinator.com" ;
const page = 新しいページ ();
待つ

ページ。 goto ( HN_ORIGIN );
const ストーリー = ページ。抽出([{
セレクター: "tr.athing" 、制限: 3 、
フィールド: {
id: { attr: "id" }、rank: ".rank" 、title: ".titleline > a" 、
URL: { セレクター: ".titleline > a" 、属性: "href" }
}
}]);
for ( const ストーリーオブストーリー) {
ページを待ちます。 goto ( HN_ORIGIN + "/item?id=" + story.id);
{を試してください
ストーリー.コメント = ページ。抽出([{
セレクター: "tr.athing.comtr:has(.commtext)" 、制限: 3 、
フィールド: {
作成者: ".hnuser" 、
テキスト: ".commtext"
}
}]);
} catch { story.comments = []; }
}
ストーリーを返す。
ループ、マップ、フィルター、try/catch を備えた、読み取り可能なバニラ JavaScript。そしていくつか
ネイティブ ツールの組み込みプリミティブ。それでおしまい。最後のトップレベル
式は JSON として自動出力されます。
もちろん、PandaScript を手動で作成または編集したり、AI に依頼したりすることもできます。
必要に応じて、コーディングアシスタントを使用してください。
PandaScript は仕様上 CDP を使用しません
Lightpanda ブラウザは依然として lightpandaserve と CDP を通信しており、積極的に取り組んでいます。
それを開発しています。ここでの決定はより狭いものであり、CDP を内部に入れないことを選択しました。
エージェント。
従来のヘッドレス自動化は、CDP 全体のあらゆるアクションを数百ものプロセスでマーシャリングします。
別のプロセスでブラウザに対して実行されるメソッドの数。エージェントはスキップします
それ。 Lightpanda のエンジンに対してインプロセスで実行され、小さなセットを呼び出します。
ネイティブ コマンドを直接使用できます。
ネイティブのインプロセス呼び出しがワイヤ プロトコル全体のすべてのクリックとフィルのマーシャリングを置き換えるため、シリアル化のオーバーヘッドはありません。
1 つのバイナリが自動的に動作するため、セットアップが簡単になります。起動する別のブラウザ、接続するデバッグ ポート、セットアップする NodeJS/Python 環境、管理する CDP クライアントがありません。
おまけに、ネイティブ コマンドは、運転中でも同じツール サーフェスになります。
自然言語命令、コマンド、または外部コマンドを使用してそれらを実行します。
lightpanda mcp を介した LLM。

これは、ブラウザを最初から開発するのではなく、ブラウザを開発することの (多くの) 利点の 1 つです。
Chromium をフォークする: 新しい AI 機能をネイティブに構築できるようになります。
API キーを指定するか、何も指定せずに実行します。
import ANTHROPIC_API_KEY=sk-ant-... # または OPENAI_API_KEY / GOOGLE_API_KEY / HF_TOKEN;
# ローカル LLM の場合は何もしません
lightpanda エージェント # インタラクティブ REPL
lightpanda エージェント -- タスク「news.ycombinator.com のトップ記事?」 #ワンショット
lightpanda エージェント --no-llm # スラッシュ コマンドのみ、LLM なし
lightpanda エージェント hn.js # 保存されたスクリプトを再生します。LLM もキーもありません
REPL では、英語または /goto 、 /extract などを使用して探索します。できます
/save <my_script>.js (アルファ機能) を使用して現在のセッションから再現可能な PandaScript を生成し、次に /save hn.js を使用してそれを再生します。
lightpanda エージェント <my_script>.js 。
エンドツーエンドのウォークスルーを試す
エージェント
チュートリアル
Hacker News にログインし、記事を抽出し、保存し、
そしてリプレイ。完全なリファレンスが必要な場合は、エージェントのドキュメントですべての内容がカバーされています。
プロバイダー、スラッシュ コマンド、フラグ、およびスクリプト形式のドキュメントで説明されています。
JavaScript API。
これは、あなたの言語を翻訳することによってヘッドレスブラウザを駆動する組み込みエージェントです。
自然言語またはスラッシュ コマンドとしてネイティブ ブラウザー アクションにリクエストを送信します。
Lightpanda 独自のエンジンに対してインプロセスで実行され、からのモデルを使用できます。
Anthropic、OpenAI、Gemini、Hugging Face、または地元の Ollama、またはモデルがまったくありません。
PandaScript は次のスクリプトです。
ブラウザのアクションとワークフローを自動化し、Puppeteer や
劇作家。 Lightpanda バイナリ上で直接実行できます。
NodeJS/Python を使用した個別のクライアント セットアップ。これは、小さな機能を備えたバニラの JavaScript です。
ネイティブプリミティブのセット。
いいえ、それは意図的なものです。エージェントは代わりにネイティブのインプロセス コマンドを呼び出します。
Lightpanda は現在も CDP の開発を積極的に行っています

lightpanda 用のサーブなので、既存の
Puppeteer と Playwright のワークフローは影響を受けません。
LLM または API キーなしで実行できますか?
はい。スラッシュコマンドのみモードの場合は --no-llm を使用し、 /goto 、 /click 、と入力します。
および /extract を直接実行します。保存されたスクリプトの再生にはモデルや
API キーにより、記録されたセッションはトークンなしで確定的に実行されます。
セッション中、状態が変化するすべての呼び出しが記録されます。 LLM を接続すると、
/save は、セッション インテントから再現可能な PandaScript を生成します (これは
アルファ機能については、積極的に開発および改善を行っています)。 LLM がなければ、
記録された /commands 呼び出しを逐語的に転写します。
シークレットは、 $LP_HN_PASSWORD のように、 $LP_* プレースホルダーとして書き込まれます。彼らは
Lightpanda 内で実行時に解決されるため、モデルのコンテキストには決して入りません。
または保存されたスクリプト ファイル。
代わりに外部エージェントから Lightpanda を駆動できますか?
はい。ネイティブ MCP サーバーは、
クロード コードなどの MCP 対応クライアントに対してモデルなしで同じツール サーフェイスを提供
Lightpanda内で実行されています。セットアップについては、MCP サーバーのガイドを参照してください。
IconLinkedin Francis は以前、2020 年に ChannelAdvisor に買収された e コマース分析プラットフォームである BlueBoard を共同設立しました。彼は大規模な自動化システムを実行しているときに、この種の作業には既存のブラウザーがいかに制限されているかを認識しました。 Lightpanda は、Web を自動化するためのより速く、より信頼性の高い方法を開発者に提供したいという彼の願いから生まれました。
ブラウザ エージェントは 4 つのツールです。それは 1 つである必要があります。
LLM は実行時ではなくビルド時に実行されます
PandaScript は仕様上 CDP を使用しません
エンドツーエンドのウォークスルーを試す
LLM または API キーなしで実行できますか?
代わりに外部エージェントから Lightpanda を駆動できますか?

## Original Extract

Lightpanda now has a built-in agent. Talk to it in natural language, get a reliable script back. Run that script forever only calling an LLM when it needs fixing.

Introducing Lightpanda Agent and PandaScript: LLM at buildtime not runtime - Blog | Lightpanda logo IconDiscord
Toggle navigation Back Introducing Lightpanda Agent and PandaScript: LLM at buildtime not runtime
IconLinkedin announcement automation agents Wednesday, June 17, 2026 TL;DR
A browser agent today is four separate tools wired together: Chrome, a CDP
library, an LLM, and an agent framework. Our native agent collapses those four
layers into one binary.
You talk to it in natural language and it does the work against a real browser.
“Go to this website”, “login”, “extract this data”, “tell me that”: language is
the new native interface for the browser. And if you want to replay your
session, it hands you a reproducible script. No CDP, no heavy browser server,
no complex setup, no LLM at runtime: everything is built in into Lightpanda.
Your browser agent is four tools. It should be one.
To run a browser agent today, you install Chrome, drive it over CDP with
Puppeteer or Playwright, wire in an LLM to make decisions, then wrap the whole
thing in an agent framework to orchestrate the calls. It works. Well, kind of:
it’s a lot of moving parts for what’s often, “load a page and read some
elements off it.”
Each of those layers exists to translate between a human-facing browser and a
machine that wants to drive it. CDP (the Chrome DevTools Protocol) was built so
a developer could inspect a running browser from the outside, not so a program
could operate one. And to wire an LLM, you need an MCP-to-CDP layer. The agent
wraps a model around a tool that was never meant to be driven by a model. You
are paying a translation tax at every layer.
When we started Lightpanda four years ago, the question was: “what would a
browser look like if you built it for machines instead of people looking at a
screen?” lightpanda agent is that bet applied to the agentic stack. It is one
binary that contains the browser, the runtime, and the agent.
The browser is the same engine behind lightpanda serve and lightpanda fetch : it loads webpages, runs JavaScript, and handles the DOM.
The runtime consists of a small set of native tools ( goto , click , fill ,
extract , evaluate , search ), that let you drive the browser (slash commands). The
LLM that reads your request and picks tools is optional. It runs against
Anthropic, OpenAI, Gemini, Hugging Face or local Ollama, or with no key at all
in slash-only mode.
The LLM runs at buildtime, not runtime
This is the idea that everything else hangs off. Every other browser agent is a
black box that calls a model on every step. With lightpanda agent , the model
figures out the task, we capture that work as code, and generate a script (we
call it PandaScript) that is reproducible and deterministic. When you replay
it, you don’t need a model and the LLM is gone from the loop.
There’s no model call sitting between you and the next action at replay time,
no Chrome process to host, no NodeJS/Python environment to setup, and no
Playwright/Puppeteer code to write. You just pass the script to Lightpanda
binary, so a run is bound only by how fast the engine drives the page. And
because nothing non-deterministic runs at replay, the same script produces the
same result every time. You pay the model once to write the file. After that,
it’s plain JavaScript that you own.
Here’s what that PandaScript looks like. This script grabs the top 3 Hacker
News stories, then visits each thread for its top comments:
const HN_ORIGIN = "https://news.ycombinator.com" ;
const page = new Page ();
await page. goto ( HN_ORIGIN );
const stories = page. extract ([{
selector: "tr.athing" , limit: 3 ,
fields: {
id: { attr: "id" }, rank: ".rank" , title: ".titleline > a" ,
url: { selector: ".titleline > a" , attr: "href" }
}
}]);
for ( const story of stories) {
await page. goto ( HN_ORIGIN + "/item?id=" + story.id);
try {
story.comments = page. extract ([{
selector: "tr.athing.comtr:has(.commtext)" , limit: 3 ,
fields: {
author: ".hnuser" ,
text: ".commtext"
}
}]);
} catch { story.comments = []; }
}
return stories;
Readable vanilla JavaScript with loops, map, filter, and try/catch. And a few
built-in primitives for our native tools. That’s it. The last top-level
expression auto-prints as JSON.
And of course you can also write or edit a PandaScript manually, or ask your AI
coding assistant to do so if you prefer.
PandaScript doesn’t use CDP, by design
Lightpanda browser still speaks CDP with lightpanda serve , and we are actively
developing it. The decision here is narrower: we chose not to put CDP inside
the agent.
Traditional headless automation marshals every action across CDP, with hundreds
of methods running against a browser in a separate process.. Our agent skips
that. It runs in-process against Lightpanda’s engine and calls a small set of
native commands directly.
There is no serialization overhead , because a native in-process call replaces marshalling every click and fill across a wire protocol.
Setup is easier , because one binary drives itself: there is no separate browser to launch, no debugging port to wire up, no NodeJS/Python environment to setup, and no CDP client to manage.
As a bonus, the native commands are the same tool surface whether you drive
them with natural language instructions, with commands, or with an external
LLM through lightpanda mcp .
That’s one (of many) advantages of developing a browser from scratch instead of
forking Chromium: it allows us to build new AI features natively.
Point an API key at it, or run it with none:
export ANTHROPIC_API_KEY=sk-ant-... # or OPENAI_API_KEY / GOOGLE_API_KEY / HF_TOKEN;
# or none for local LLM
lightpanda agent # interactive REPL
lightpanda agent --task "top story on news.ycombinator.com?" # one-shot
lightpanda agent --no-llm # slash commands only, no LLM
lightpanda agent hn.js # replay a saved script, no LLM, no key
In the REPL, explore in English or with /goto , /extract , and the rest. You can
generate a reproducible PandaScript from the current session with /save <my_script>.js (alpha feature), and then /save hn.js , then replay it with
lightpanda agent <my_script>.js .
Try the end-to-end walkthrough
The agent
tutorial
takes you through the whole loop: log in to Hacker News, extract stories, save,
and replay. If you want the full reference, the agent docs cover every
provider, slash command, and flag, and the script format docs explain the
JavaScript API.
It is a built-in agent that drives a headless browser by translating your
requests into native browser actions, in natural language or as slash commands.
It runs in-process against Lightpanda’s own engine and can use a model from
Anthropic, OpenAI, Gemini, Hugging Face or local Ollama, or no model at all.
PandaScript is a script to
automate browser actions and workflows, designed to replace Puppeteer or
Playwright. It can be run directly on the Lightpanda binary without needing a
separate client setup with NodeJS/Python. It’s vanilla JavaScript with a small
set of native primitives.
No, and that is deliberate. The agent calls native in-process commands instead.
Lightpanda is still actively developing CDP for lightpanda serve , so existing
Puppeteer and Playwright workflows are unaffected.
Can I run it without an LLM or API key?
Yes. Use --no-llm for slash-commands-only mode, where you type /goto , /click ,
and /extract directly. Replaying a saved script doesn’t require a model or an
API key, so recorded sessions run token-free and deterministically.
During a session, every state-changing call is recorded. With an LLM connected,
/save generates a reproducible PandaScript from the session intent (this is an
alpha feature, we are actively developing and improving it). Without an LLM, it
transcribes the recorded /commands calls verbatim.
Secrets are written as $LP_* placeholders, like $LP_HN_PASSWORD . They are
resolved at runtime inside Lightpanda, so they never enter the model context
or the saved script file.
Can I drive Lightpanda from an external agent instead?
Yes. The native MCP server exposes the
same tool surface to any MCP-aware client, like Claude Code, with no model
running inside Lightpanda. See the MCP server guide for setup.
IconLinkedin Francis previously cofounded BlueBoard, an ecommerce analytics platform acquired by ChannelAdvisor in 2020. While running large automation systems he saw how limited existing browsers were for this kind of work. Lightpanda grew from his wish to give developers a faster and more reliable way to automate the web.
Your browser agent is four tools. It should be one.
The LLM runs at buildtime, not runtime
PandaScript doesn’t use CDP, by design
Try the end-to-end walkthrough
Can I run it without an LLM or API key?
Can I drive Lightpanda from an external agent instead?
