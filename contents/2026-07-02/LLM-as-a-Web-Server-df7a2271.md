---
source: "https://jin.codes/writings/llm-as-a-web-server/"
hn_url: "https://news.ycombinator.com/item?id=48758131"
title: "LLM as a Web Server"
article_title: "LLM as a Web Server - jin codes"
author: "jinthagerman"
captured_at: "2026-07-02T08:42:10Z"
capture_tool: "hn-digest"
hn_id: 48758131
score: 2
comments: 1
posted_at: "2026-07-02T08:11:48Z"
tags:
  - hacker-news
  - translated
---

# LLM as a Web Server

- HN: [48758131](https://news.ycombinator.com/item?id=48758131)
- Source: [jin.codes](https://jin.codes/writings/llm-as-a-web-server/)
- Score: 2
- Comments: 1
- Posted: 2026-07-02T08:11:48Z

## Translation

タイトル: Web サーバーとしての LLM
記事のタイトル: Web サーバーとしての LLM - jin code
説明: すべての新しいページがテンプレートではなく LLM によってライブで生成されるサイトを構築しました。ここ

記事本文:
Web サーバーとしての LLM - jin code ホーム Web サーバーとしての LLM についての記事
6 月の Prompt Poets Society のミートアップで、JD Trask の講演の 2 行が私の心に残りました。
Web サーバーはテキスト入力、テキスト出力です。
では、LLM が Web サーバーの場合はどうなるでしょうか? Web サーバー用のコードを記述するモデルではなく、すべての新しいページ自体を提供するモデル、つまりその場で発明された新しい HTML です。
明らかな理由から、これはひどいアイデアです。クリックするたびに実際のお金がかかり、ミリ秒ではなく数秒かかり、毎回少しずつ異なる結果が返されます。でも気になったので作ってみました。
この実験は 🔥 token burner 🔥 と呼ばれます。これは、LLM が実質的にビュー層である Web アプリです。
チャット ボックスを通じてモデルに話しかけてテキストを返すのではなく、Web ページを通じてモデルに話しかけると、別の Web ページが返されます。ボタン、フォーム、チャート、ゲーム、コントロール、レイアウトはすべて、会話のさまざまな形にすぎません。ボタンをクリックすることは、通常の意味での「アプリの使用」ではありません。次のメッセージを送信中です。
つまり、表面的にはバイブコーディングツールのように見えても、それはバイブコーディングツールではありません。生成された HTML は最終製品ではありません。それは会話です。
実際の様子は次のとおりです。プレーンテキストのプロンプトを使用して会話をシードします。
そして、返信は段落ではなくページとして返されます。
ページ上のすべてが新たなターンにかかるわけではありません。ここで惑星を選択すると、通常のクライアント側 JS でその詳細が拡張されます。しかし、「Deep Dive into Earth」のようなボタンは会話内の次のメッセージを送信し、モデルはそれに応じて新しいページをレンダリングします。モデルは、各ページを作成する際に、すでに存在するコンテンツで何が処理できるか、そして何が全く新しい方向に向かうに値するかという線を自ら描きます。
その 1 つのエクスプローラー ページから、火星、木星、地球の順に詳しく調べました。 Th

リーは、同じ親から分岐した 3 人の兄弟が次のように答えます。
これらのページをデザインした人はいません。それぞれが新鮮な世代であり、モデルは各部門に独自のアートディレクションを与えました。火星は黄土色になり、木星は豪華な金色になり、地球は親しみやすい青色になりました。これらのリンクはライブ ターンなので、この正確な会話を拾って、私が行ったことのない場所に分岐させることができます。
リクエストが実際にページになる仕組み
仕組みは思っているよりも単純です。すべての対話がモデルと往復するわけではありません。多くのページにはアニメーションやローカル UI 状態用の通常のクライアント側 JS がありますが、モデルが新しいターンを実行したい要素には data-prompt 属性 (またはフォームの場合は data-prompt-template) が含まれており、挿入された小さなスクリプトによってクリックが変換され、それらの要素に対する POST がそのプロンプトを本文として /c/:convoId/:turnId/act に送信されます。生成されたボタンは文字通り、送信を待つ次のメッセージです。
<button data-prompt="火星へのドリル">
火星を探検する →
</ボタン>
サーバー上では、次の 3 つのことが起こります。
履歴をロードします。再帰 SQL クエリは、parent_turn_id を介して会話ツリーをたどるので、分岐した会話から兄弟の会話がコンテキストに漏れることはありません。
強制ツールでモデルを呼び出します。新しいプロンプトが追加され、単一ツール render_page({ body_html, summary, style_hints }) にロックされた toolsChoice を使用してモデルが呼び出されます。構造化 HTML を返還する以外に選択肢はありません。
レンダリングして応答します。 body_html は共有テンプレート (Tailwind、D3、Chart.js、Three.js、Tone.js はすべてプリロードされています) にドロップされ、 GET /c/:convoId/:turnId に対する実際の HTTP 応答として返されます。
私にとってこの全体がピンと来た 1 つの詳細は、モデルが以前のターンで生成した生の HTML を決して参照しないことです。毎ターン自分自身について書き込む概要文字列のみが取得されます

アシスタント履歴として再生されました。コンテキストは小さいままです。このモデルは出力トークンを意図的に無駄にしています (それが名前の由来です) が、フィードバックされるものについては倹約しています。
本格的なエンジニアリングが必要な部品
たった 1 つの楽しい完成品を完成させるだけでなくなれば、おもちゃのバージョンでは十分ではなくなります。
生成は非同期です。ターンは status: "pending" として挿入され、LLM 呼び出しがバックグラウンドで実行され、ページは約 1 秒ごとに /status エンドポイントをポーリングし、完了するまでスピナーを表示します。
分岐と重複排除。ターンは、parent_turn_id によってキー設定されたツリーであるため、同じ親ターンから起動される同じプロンプトは、新しい世代を書き込むのではなく、既存の生成されたページを再利用します。
システムプロンプトは、偽装できない内容を伝える必要があります。私は、モデルがクライアント側で「バックエンド機能」を偽装することを明示的に禁止します。実際のロジックが必要な場合は、代わりに別のターンを往復する必要があります。
ナビゲーションはブラウザの仕事です。モデルは戻るボタンのレンダリングを禁止されています。生成された「戻る」ボタンは実際にはどこにも戻ることができません。以前のページを模倣して新しいターンを書き込むだけです。ブラウザの実際の「戻る」ボタンはすでに機能しており、ツリーを上って戻って新しい場所に分岐するだけで、半分は終わりです。
(そして、はい、会話リンクは共有可能であるため、他の人のブランチを開くということは、プロンプトでモデルに書き込まれたあらゆる JS を実行することを意味します。これはおもちゃです。おもちゃのように扱ってください。)
実際に使えるものなのでしょうか？それは当然以上のものです。フロンティア モデルでは、ページが世代を重ねるごとに感じられ、スピナーは歓迎され続けます。 Gemini 3.1 Flash Lite のような小規模で高速なモデルは、驚くほど通常のページ読み込み領域に近づきます。
そして、それらの各ページにかかる費用は次のとおりです。
claude-sonnet-4.6 はデフォルトのモデルであるため、その平均ははるかに基づいて構築されています
他のものよりページ数が多い。の

残りは小さくてノイズの多いサンプルです。
スタート ページにはモデル ピッカーがあるため、よりスマートなページをより高速で安価なページと交換することができます。
興味深いのは、「LLM は HTML を書くことができる」ということではありません。私たちはそれが可能であり、それがうまくできることを知っています。興味深いのは、実際の Web アプリの幻想がどこで壊れるか、そしてそれぞれの壊れた部分を修正するには何が必要かを正確に観察することです。
国家は自由ではない。コンテキスト内で明示的に再生されない限り、何も持続しません。 HTML ではなく要約というトリックが、これが 10 ターン目までに自身のトークンコストで崩壊しない理由のすべてです。
一貫性を保つには、モデルを大きくするだけでなく、配管も必要です。ブランチの重複除去、ターン ツリー、ツールの強制呼び出し、これらはいずれもスマートなモデルではありません。モデルが過去のターンと矛盾しないようにするのは足場です。
それがこのようなものを構築することの本当の成果です。目新しい Web サーバーではありませんが、モデルが説得力を持って即興で実現できるものと、実際にその背後に記録システムが必要なものとの間の境界線を感じるための非常に迅速な方法です。
トークンバーナーは閉ループです。モデルの唯一の「バックエンド」はそれ自体の想像力であるため、レンダリングされるものは重要な意味で間違っている可能性があります。当然の次のステップは、モデルに実際のバックエンドを与え、モデルを実際のツールまたは MCP サーバーに接続して、生成されるページがもっともらしい推測ではなく実際の状態を反映するようにすることです。
私が何度も戻ってくるユースケースは、旅行の計画です。特定のフライトを予約するのではなく、それは解決された手続き上の問題ですが、その前のやっかいな部分があります。「どこか暖かくて、おいしい食べ物があり、観光客向けではないけど、私のパートナーは長時間のフライトが嫌いで、私たちには幼児がいます。」誰もそのフォームに記入する人はいませんし、まさに決定論的な流れに変えられることに抵抗する類の推論です。ツールを使用した実際の空き状況により、旅程の根拠が維持されます

は実際には予約可能ですが、値は API 呼び出しではありません。これは、次にスローされる例外を中心にそれ自体を再構築できる UI です。
ただし、そこでもすべてをゼロから生成する必要はありません。トークンバーナーはすでにその出力を共有テンプレートにラップしているため、私が言いたいのはクロムとコンテンツの分割ではありません。これは、毎回同じ外観と動作をする、再利用可能な最上級のコンポーネント (予約フォーム、日付ピッカー、旅程カード) であり、モデルは、生成するページ内および生成するページ間でドロップできます。場合によっては、決して変更されないフロー用に事前にデザインされたページ全体を使用することもできます。モデルは、旅程、トレードオフ、最新の例外に応じて設定されたオプションなど、真に流動的な部分を即興で提供し続けます。信頼できるパーツは、毎ターン再発明するのではなく、ライブラリから構成されます。この組み合わせは、「Web アプリとしての LLM」のより正直なバージョンのように感じられます。 UI レイヤー全体を置き換えるモデルではなく、事前にテンプレート化するには流動的すぎる 1 つの部分を処理するモデルです。
それまでの間、クレジットが残っているうちに遊んでください。
執筆に戻る
© 2026 ジン・ブデルマン

## Original Extract

I built a site where every new page is generated live by an LLM instead of a template. Here

LLM as a Web Server - jin codes Home Writings About LLM as a Web Server
At June’s Prompt Poets Society meetup , two lines from JD Trask’s talk stuck with me.
Web servers are text in, text out.
So what happens if the LLM is the web server? Not a model writing code for a web server, but the model serving every new page itself, fresh HTML invented on the spot.
It’s a terrible idea for all the obvious reasons: every click costs real money, takes seconds instead of milliseconds, and comes back a little different every time. But I was curious, so I built it.
The experiment is called 🔥 token burner 🔥 : a web app where the LLM is effectively the view layer.
Instead of talking to a model through a chat box and getting text back, you talk to it through a web page and get another web page back. Buttons, forms, charts, games, controls, and layouts are all just different shapes a conversation can take. Clicking a button isn’t “using the app” in the normal sense. It’s sending the next message.
So it’s not a vibe coding tool, even if it looks like one on the surface. The generated HTML is not the end product. It is the conversation.
Here’s what that looks like in practice. You seed a conversation with a plain-text prompt:
And the reply comes back as a page, not a paragraph :
Not everything on the page costs a new turn. Selecting a planet here expands its details with ordinary client-side JS. But a button like “Deep Dive into Earth” sends the next message in the conversation, and the model renders a fresh page in response. The model draws that line itself while writing each page: what it can handle with the content already on it, and what deserves a whole new turn.
From that one explorer page I drilled into Mars , then Jupiter , then Earth . Three replies, three siblings branching off the same parent turn:
Nobody designed these pages. Each one is a fresh generation, and the model gave each branch its own art direction. Mars went ochre, Jupiter went regal gold, and Earth got the friendly blue treatment. Those links are the live turns, so you can pick up this exact conversation and branch it somewhere I never went.
How a request actually becomes a page
The mechanics are simpler than they sound. Not every interaction round-trips to the model. Plenty of pages have ordinary client-side JS for animations or local UI state, but any element the model wants to make a new turn carries a data-prompt attribute (or a data-prompt-template for forms), and a small injected script turns clicks and submits on those elements into a POST to /c/:convoId/:turnId/act with that prompt as the body. A generated button is literally the next message, waiting to be sent:
<button data-prompt="Drill into Mars">
Explore Mars →
</button>
On the server, three things happen:
Load history. A recursive SQL query walks up the conversation tree via parent_turn_id , so branched conversations don’t leak sibling turns into context.
Call the model with a forced tool. The new prompt gets appended, and the model is called with toolChoice locked to a single tool, render_page({ body_html, summary, style_hints }) . It has no choice but to hand back structured HTML.
Render and respond. The body_html gets dropped into a shared template (Tailwind, D3, Chart.js, Three.js, Tone.js all preloaded) and served back as the actual HTTP response for GET /c/:convoId/:turnId .
The one detail that made the whole thing click for me is that the model never sees the raw HTML it generated on previous turns. Only the summary string it writes about itself each turn gets replayed as assistant history. Context stays small. The model is deliberately wasteful with output tokens (hence the name) but frugal about what gets fed back in.
The parts that needed real engineering
Once you’re not just doing a single fun completion, the toy version stops being enough:
Generation is async. A turn gets inserted as status: "pending" , the LLM call runs in the background, and the page polls a /status endpoint every second or so, showing a spinner until it’s done.
Branching and dedup. Because turns are a tree keyed by parent_turn_id , the same prompt fired from the same parent turn reuses the existing generated page instead of burning a fresh generation.
The system prompt has to say what it can’t fake. I explicitly forbid the model from faking “backend functionality” client-side. Anything that needs real logic has to round-trip through another turn instead.
Navigation is the browser’s job. The model is banned from rendering back buttons. A generated “back” button can’t actually return anywhere; it would just burn a fresh turn imitating an earlier page. The browser’s real back button already works, and walking back up the tree to branch off somewhere new is half the point.
(And yes, since conversation links are shareable, opening someone else’s branch means running whatever JS their prompts talked the model into writing. It’s a toy; treat it like one.)
Is it actually usable? More than it deserves to be. With a frontier model you feel every page generation, and the spinner overstays its welcome. A small fast model like Gemini 3.1 Flash Lite gets surprisingly close to ordinary page-load territory:
And what each of those pages costs:
claude-sonnet-4.6 is the default model, so its averages are built on far
more pages than the others. The rest are smaller, noisier samples.
There’s a model picker on the start page, so you can trade smarter pages for faster, cheaper ones yourself.
The interesting part isn’t “look, the LLM can write HTML.” We know it can do that and do it well. The interesting part is watching exactly where the illusion of a real web app breaks down, and what it takes to patch each break.
State isn’t free. Nothing persists unless it’s explicitly replayed into context. The summary-not-HTML trick is the whole reason this doesn’t collapse under its own token cost by turn ten.
Consistency requires plumbing, not just a bigger model. Branch dedup, the turn tree, forced tool calls: none of that is the model being smart. It’s scaffolding that keeps the model from contradicting its past turns.
That’s the real payoff of building something like this. Not a novelty web server, but a very fast way to feel out the line between what a model can convincingly improvise and what actually needs a system of record behind it.
token burner is a closed loop. The model’s only “backend” is its own imagination, so nothing it renders can be wrong in a way that matters. The natural next step is to give it a real backend, wiring the model up to actual tools or MCP servers so the pages it generates reflect real state instead of a plausible guess at it.
The use case I keep coming back to is trip planning. Not booking a specific flight, which is a solved, procedural problem, but the messy part before it: “somewhere warm, good food, not touristy, but my partner hates long flights and we have a toddler.” Nobody fills out a form for that, and it’s exactly the kind of reasoning that resists being turned into a deterministic flow. Real availability via tools keeps the itinerary grounded in what’s actually bookable, but the value isn’t the API calls. It’s a UI that can restructure itself around whatever exception you throw at it next.
Even there, though, not everything should be generated from scratch. token burner already wraps its output in a shared template, so the split I mean isn’t chrome versus content. It’s first-class, reusable components (a booking form, a date picker, an itinerary card) that look and behave the same every time, and that the model can drop into and between the pages it generates. Maybe even whole predesigned pages for the flows that never change. The model keeps improvising the genuinely fluid part: the itinerary, the tradeoffs, the options laid out in response to your latest exception. The dependable parts it composes from a library instead of reinventing them each turn. That mix feels like the more honest version of “LLM as a web app.” Not the model replacing the whole UI layer, but the model handling the one part that’s too fluid to template in advance.
In the meantime, have a play while I still have credit.
Back to Writings
© 2026 Jin Budelmann
