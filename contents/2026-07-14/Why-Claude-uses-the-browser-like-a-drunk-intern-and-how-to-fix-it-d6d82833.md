---
source: "https://pluno.ai/blog/why-claude-uses-browser-like-drunk-intern"
hn_url: "https://news.ycombinator.com/item?id=48907910"
title: "Why Claude uses the browser like a drunk intern, and how to fix it"
article_title: "Why Claude Uses the Browser Like a Drunk Intern"
author: "korabs"
captured_at: "2026-07-14T15:17:09Z"
capture_tool: "hn-digest"
hn_id: 48907910
score: 2
comments: 1
posted_at: "2026-07-14T14:58:57Z"
tags:
  - hacker-news
  - translated
---

# Why Claude uses the browser like a drunk intern, and how to fix it

- HN: [48907910](https://news.ycombinator.com/item?id=48907910)
- Source: [pluno.ai](https://pluno.ai/blog/why-claude-uses-browser-like-drunk-intern)
- Score: 2
- Comments: 1
- Posted: 2026-07-14T14:58:57Z

## Translation

タイトル: クロードが酔ったインターンのようにブラウザを使用する理由とその修正方法
記事のタイトル: クロードが酔ったインターンのようにブラウザを使用する理由
説明: ブラウザ エージェントはランタイムとしてピクセルを使用するため、速度が遅くなります。この修正は、UI の下で製品コントラクトを学習し、それらを直接使用することです。

記事本文:
クロードが酔ったインターンのようにブラウザを使用する理由 コンテンツにスキップ 製品価格 お客様のリソース ログイン ブック デモ FAQ だけでなく、複雑な技術的なチケットを解決する AI サポート エージェント。 Zendesk 上の B2B チーム向けに構築されています。
Pluno は、G2 で 4.8/5 の最高評価の AI サポート エージェントであり、Reddit でも定期的に推奨されています。
著作権 © 2026 Unbrained GmbH
ブログに戻る AI エージェント ブラウザ エージェントの自動化 クロードが酔ったインターンのようにブラウザを使用する理由とその修正方法
ブラウザ エージェントを誤って構築した経緯
クロードが酔ったインターンのようにブラウザを使用する理由
ブラウザエージェントの別のベンチマーク
ブラウザ エージェントを誤って構築した経緯
私たちは過去数年間、顧客サポート用の AI エージェントの構築に費やしました。
最初のバージョンは基本的に、サポート コンテキスト (ドキュメント、チケット、以前の解決策、内部メモ) に対するエージェント検索でした。
これにより、驚くほどの作業が完了しますが、サポート作業は製品の状態にすぐに影響します。多くの場合、その答えは、構成、権限、請求状態、または 3 画面の深さに埋め込まれた管理設定など、顧客のアカウントによって異なります。
そのため、代理店は製品を検査し、場合によっては製品を変更する必要があります。
私たちは「ブラウザ エージェントを構築しましょう」というところから始めたわけではありません。
「エージェントは API に接続する必要がある」ということから始めました。
クロードが酔ったインターンのようにブラウザを使用する理由
Comet、Claude、ChatGPT Atlas、Browser Use など、これまでに少なくとも 1 つのブラウザ エージェントを使用したことがあるでしょう。
ページを探索してください。もしかしたらスクリーンショットを撮ることもあるかもしれません。可能性のあるターゲットを見つけます。マウスを動かします。クリック。待って。もう一度スクリーンショット。スクロール。ドロップダウンを試してください。モーダルから回復します。状態が変化したかどうかを推測します。
わかりました。これが私たちが持っている情報であり、公開されている機能です。しかし、それは恐ろしく遅く、非常に非効率的です。
音声エージェントにバーブの予約を頼むようなもの

別の音声エージェントに電話して予約することもできます。技術的には可能です。もしかしたら、初めては魔法さえかかるかもしれません。しかし、両方のシステムが構造化データを直接交換した可能性があることを理解すると、不合理です。
UI は、レイアウト、ラベル、プログレッシブ開示、ホバー状態、応答性の高いブレークポイント、視覚的な階層など、人間向けに最適化されています。これらはいずれも、製品の実際の実行モデルではありません。
その下では、フロントエンドがすでに構造化システムと通信しています。クリックは、リクエスト、ミューテーション、RPC 呼び出し、WebSocket メッセージ、サーバー アクション、署名されたリクエスト、永続化された GraphQL 操作、または製品固有のコマンドになります。入力、出力、エラー、権限、副作用、状態遷移があります。
これがエージェントが望むインターフェイスであり、まさにそれが MCP サーバーと CLI がある理由です。
しかし、多くの製品にはまだ MCP サーバーが搭載されていないため、エージェントは画面を見てどのボタンをクリックするかを判断する必要があります。
エージェントは、ピクセルではなく、UI の下にある構造化インターフェイスを使用する必要があります。
1 つの操作で、十数の UI ステップを置き換えることができます。小さなスクリプトで脆弱なクリック パスを置き換えることができます。
スクリーンショットやレイアウトから意図を推測するようモデルに依頼するのではなく、モデルにオブジェクト、フィールド、メソッド、応答、エラー、不変条件を与えます。
それは現在のモデルにはるかによく適合します。彼らは、コード、構造化データ、ツール呼び出しに関してすでに非常に優れています。 「このスクリーンショットを解釈し、正しいものをクリックし、UI が変化しないことを祈る」というタスクの場合、信頼性ははるかに低くなります。
ピクセルは、コード内にすでに存在する動作の非可逆トランスポートです。
したがって、基礎となる API を直接使用しないのはなぜでしょうか?という質問は簡単に思えます。
難しいのは、インターフェースが存在しないということではありません。
難しいのは、それが構造的に公開されていない、または文書化されていないことです。
通常は非公開です。安定していない可能性もあります。スプリであることが多い

REST エンドポイント、GraphQL 操作、WebSocket メッセージ、署名されたリクエスト、WASM モジュール、Web ワーカー、クライアント キャッシュ、サーバー アクション、およびフロントエンドのみの規約にまたがります。重要な動作の一部は、リクエストのシーケンス、ID 解決、検証応答、キャッシュの無効化、または権限依存のブランチに存在する場合があります。
「API を使用するだけ」では、UI が依存している実際のコントラクトをエージェントが学習する必要がある部分がスキップされます。
どの操作がそれらを作成または変更するか
サーバー側の言語がユーザー側の UI 言語にどのようにマッピングされるか
てか、ステータス ID さえユーザーの言語に依存することがある (*咳* Zendesk)
その情報は、フロントエンド コード、ネットワーク トラフィック、ランタイム状態、キャッシュの更新、エラー処理など、暗黙的にのみ存在することがよくあります。
フロントエンドがアクションを実行すると、製品コントラクトが明らかになります。オブジェクトを取得します。操作を送信します。入力を検証します。エラーを受け取ります。状態を更新します。
エージェントはソース コードをチェックし、何が起こっているかを観察し、ワークフローの形状を抽出できます。
ルート、操作、またはメッセージ パターン
必須フィールドとオブジェクトの関係
検証、認可、成功の状態
では、エージェントにそこから学習させるにはどうすればよいでしょうか?
単純なループはすでに非常にうまく機能しています。エージェントに独自のセッションを与え、何がうまくいったか、どこで苦労したか、何を新たに発見したかに基づいてスキルを更新するように指示します。
この学習はユーザー間で行われるため、機密データについては細心の注意を払う必要があります。幸いなことに、私たちは過去 3 年間、ユーザー データを損なうことなくエージェントが過去のサポート チケットから安全に学習できる方法を見つけ出すことに費やしてきました。同じ原則がここにも当てはまります。このアプローチを使用して、現在 4,300 を超える Web アプリの製品固有のスキル コレクションを蓄積しています。
これらのスキルには API パス、コードが含まれます

スニペット、ワークフロー ロジック、ドメイン知識、役立つエッジ ケース、製品が実際にどのように動作するかについてのメモ。抽出プロセスには改善すべき点がたくさんありますが、セッションから役立つスキルを生成するようにエージェントに指示するという最も単純なアプローチでも、驚くほど遠くまで到達できます。
スキルを習得すれば、ほぼ完璧な MCP サーバーを手に入れることができます。
エージェントはコードから直接 API と通信できます。複数ステップのワークフローを実行するスクリプトを作成できます。また、抽象的なツールの説明からは決してわからない、製品固有の事柄も理解できます。
これらのスキルが製品全体に存在すると、ワークフローも 1 つのタブに関連付けられなくなります。つまり、1 つのツールからエクスポートし、データを変換し、別のツールを更新し、結果を検証します。
クロード コードがファイルを直接操作するのではなく、コンピューターを使用して VS コードを使用しているところを想像してください。
これが、Pluno と他のブラウザ エージェントの速度の違いです。
ブラウザエージェントの別のベンチマーク
ベンチマークを作成しましたが、それはでたらめです。
しかし、製品を学習した効果を他に示す方法がわかりません。
HubSpot、Notion、Stripe、Linear などを含む 24 のツールの 312 の実際のタスクにわたって、Pluno を他のブラウザ エージェントと比較しました。
Pluno は以前にそれらの製品を学習していました。 API、ワークフロー、オブジェクト モデル、権限、奇妙なエッジ ケースを認識していました。他のエージェントは主に、新しい Web サイトのように、ページを検査し、次のクリックを推測し、何かを試し、失敗した場合は回復するようにアプローチしました。
それは学術的に正確な比較ではありません。
しかし、それは私たちが構築しようとしている製品エクスペリエンスです。ユーザーは、毎朝 Stripe、Notion、Linear、Salesforce、または HubSpot を再検出するエージェントを望んでいません。彼らは、その下にあるソフトウェアをすでに理解しており、そのソフトウェアのおかげで 10 倍高速なエージェントを求めています。
主な違いは速度であると予想していました

d.直接実行すると、スクリーンショット、カーソルの移動、再試行、ページが安定するまでの待機が回避されます。
しかし、信頼性にも大きな違いがありました。多くのブラウザ エージェントは単に動作が遅いだけではありませんでした。彼らは UI で行き詰まったり、行き止まりに陥ったり、状態を失ったり、諦めたりしました。インターフェイスは速度のボトルネックだけではありませんでした。それが成功のボトルネックでした。
したがって、その点に注意して数字を読んでください。これらは、Pluno がこれまで見たことのない製品の使用に優れていることを証明するものではありません。これらは、エージェントが製品を学習し、その知識を次のタスクに使用したときに何が起こるかを示します。
タスクセットを含む完全な記事を間もなく公開する予定です。
UI の使用が適切なツールであるケースはまだあります。
初めての探索では、多くの場合ブラウザが必要になります。ビジュアル製品にはビジュアルコンテキストが必要です。キャンバス ツール、デザイン ツール、ダッシュボード、マップ、およびあいまいなワークフローでは、画面を見る必要がある場合があります。一部の製品では、使用可能なバックエンド コントラクトが公開されていない場合があります。一部のタスクでは視覚的な確認が必要です。
重要なのは、エージェントが UI を決して使用すべきではないということではありません。
重要なのは、UI が、繰り返し行われる SaaS 操作のデフォルトの実行レイヤーであってはいけないということです。
ピクセルはランタイムではなくブートストラップ層である必要があります。
インターフェースがボトルネックになっている
モデルの推論が苦手なため、ブラウザ エージェントが遅いというわけではありません。
遅いのは、間違ったインターフェイスを使用しているためです。
コードと構造化データの処理能力がますます高まっているモデルを取得し、スクリーンショット、ボタン、ドロップダウン、レイアウトの変更、モーダル状態を強制的に実行します。
ソフトウェアは UI の下ですでに構造化された動作を持っています。難しい部分は、その動作をエージェントに安全に公開することです。コントラクトの学習、トレースの編集、スキーマの検証、アクセス許可の制限、副作用の確認の要求、必要に応じて目視検査に戻ることです。
未来

ボタンをクリックするのが上手になるエージェントではありません。
UI の下にあるコントラクトを学習し、それを直接使用するのはエージェントです。
試してみたい場合は、ブラウザ拡張機能としてのエージェントをご覧ください。
最新の製品アップデートをご覧ください → 10 倍高速に動作する準備はできていますか?
Pluno をブラウザに追加すると、遅い Web ワークフローが高速な API アクションに変わります。
2026 年における Zendesk 販売のベスト代替品: 移行購入者ガイド
Zendesk Sell は 2027 年 8 月 31 日に廃止されます。Zendesk Sell の最良の代替手段 (Nutshell、Pipedrive、HubSpot、Salesforce) を移行パスと実際の価格で比較します。
2026 年の顧客サービス向上のための 6 つのベスト Ada 代替案
2026 年の Ada の代替となるベスト 6 つを、学習方法、複雑なチケットの処理方法、セットアップ方法、価格などで比較しました。これにより、チームに適切なエージェントをマッチングできます。
Freshdesk と Zendesk: 2026 年にどちらのカスタマー サポート プラットフォームがより良い価値を提供しますか?
2026 年の Freshdesk と Zendesk: 迅速かつ手頃な価格のチケット発行と、大規模なオムニチャネル サポート。価格、AI 請求、セットアップ、複雑なチケットの適合性を比較します。

## Original Extract

Browser agents are slow because they use pixels as the runtime. The fix is learning the product contracts underneath the UI, then using them directly.

Why Claude Uses the Browser Like a Drunk Intern Skip to content Product Pricing Customers Resources Login Book Demo The AI support agent that resolves complex, technical tickets — not just FAQs. Built for B2B teams on Zendesk.
Pluno is a top-rated AI support agent - 4.8/5 on G2 and regularly recommended on Reddit.
Copyright © 2026 Unbrained GmbH
Back to Blog AI Agents Browser Agents Automation Why Claude uses the browser like a drunk intern, and how to fix it
How we accidentally built a browser agent
Why Claude uses the browser like a drunk intern
A different benchmark for browser agents
How we accidentally built a browser agent
We spent the last few years building AI agents for customer support.
The first version was basically agentic search over support context: docs, tickets, prior resolutions, internal notes.
That gets you surprisingly far, but support work quickly runs into product state. The answer often depends on the customer's account: configuration, permissions, billing state, or some admin setting buried three screens deep.
So the agent has to inspect the product, and sometimes change it.
We didn't start with "let's build a browser agent."
We started with "the agent needs to be connected to the API."
Why Claude uses the browser like a drunk intern
You've probably used at least one browser agent by now, whether it's Comet, Claude, ChatGPT Atlas, Browser Use, or anything else.
Explore the page. Maybe even take a screenshot. Find the likely target. Move the mouse. Click. Wait. Screenshot again. Scroll. Try a dropdown. Recover from a modal. Guess whether the state changed.
I get it - this is the information we have, and the functionality that's exposed. But it's horribly slow, and highly inefficient.
It's like asking a voice agent to book a barber appointment by calling another voice agent. Technically possible. Maybe even magical the first time. But absurd once you realize both systems could have exchanged structured data directly.
The UI is optimized for humans: layout, labels, progressive disclosure, hover states, responsive breakpoints, visual hierarchy. None of that is the product's actual execution model.
Underneath, the frontend is already speaking to a structured system. A click becomes a request, mutation, RPC call, websocket message, server action, signed request, persisted GraphQL operation, or some product-specific command. There are inputs, outputs, errors, permissions, side effects, and state transitions.
That's the interface an agent should want, and it's exactly why we have MCP servers and CLIs.
But since many products still don't have an MCP server, we're stuck with agents looking at a screen to figure out which button to click.
Agents should use the structured interface underneath the UI, not pixels.
One operation can replace a dozen UI steps. A small script can replace a fragile click path.
Instead of asking a model to infer intent from screenshots and layout, give it objects, fields, methods, responses, errors, and invariants.
That's a much better fit for current models. They're already extremely good at code, structured data, and tool calls. They're much less reliable when the task is "interpret this screenshot, click the right thing, and hope the UI didn't shift."
Pixels are a lossy transport for behavior that already exists in code.
So the question seems easy: why not use the underlying API directly?
The hard part isn't that the interface doesn't exist.
The hard part is that it is not structurally exposed or documented.
It's usually not public. It might not even be stable. It's often split across REST endpoints, GraphQL operations, websocket messages, signed requests, WASM modules, web workers, client caches, server actions, and frontend-only conventions. Some of the important behavior may live in request sequencing, id resolution, validation responses, cache invalidation, or permission-dependent branches.
"Just use the API" skips the part where the agent has to learn the actual contract the UI is relying on.
which operations create or mutate them
how server-side language maps to the user-facing UI language
heck, sometimes even that status ids depend on the user's language (*cough* Zendesk)
That information is often present only implicitly: in frontend code, network traffic, runtime state, cache updates, and error handling.
When the frontend performs an action, it reveals the product contract. It fetches objects. It sends operations. It validates inputs. It receives errors. It updates state.
An agent can check the source code, observe what happens, and extract the shape of the workflow:
routes, operations, or message patterns
required fields and object relationships
validation, authorization, and success states
So how can we make an agent learn from that?
A simple loop already works quite well: give an agent its own sessions, then tell it to update its skills based on what worked, where it struggled, and what it newly discovered.
Since this learning happens across users, we need to be extremely careful about sensitive data. Luckily, we have spent the last three years figuring out how to make agents learn securely from past support tickets without compromising user data. The same principles apply here. Using that approach, we have already accumulated product-specific skill collections for more than 4,300 web apps now.
Those skills include API paths, code snippets, workflow logic, domain knowledge, useful edge cases, and notes about how the product actually behaves. There's a lot to improve about the extraction process, but even the simplest approach, telling the agent to generate helpful skills from its sessions, gets you surprisingly far.
Once you have the skills, it's almost like having a perfect MCP server.
The agent can talk to the API directly from code. It can write scripts that execute multi-step workflows. And it understands product-specific things that you'd never know from abstract tool descriptions.
Once those skills exist across products, workflows also stop being tied to one tab: export from one tool, transform the data, update another tool, and verify the result.
Imagine Claude using VS Code via computer use, instead of Claude Code working on the files directly.
That's the speed difference between Pluno and other browser agents.
A different benchmark for browser agents
We've created a benchmark, and it's bullshit.
But I don't know how else to show the effect of learning the product.
We compared Pluno against other browser agents across 312 real-world tasks in 24 tools, including HubSpot, Notion, Stripe, Linear, and more.
Pluno had learned those products before. It knew their APIs, workflows, object models, permissions, and weird edge cases. The other agents mostly approached them like fresh websites: inspect the page, infer the next click, try something, recover if it fails.
That isn't an academically clean comparison.
But it is the product experience we're trying to build. Users don't want an agent that rediscovers Stripe, Notion, Linear, Salesforce, or HubSpot every morning. They want an agent that already knows the software underneath and is 10x faster because of it.
We expected the main gap to be speed. Direct execution avoids screenshots, cursor movement, retries, and waiting for pages to settle.
But there was also a big difference in reliability. Many browser agents did not just run slowly. They got stuck in the UI, clicked into dead ends, lost state, or gave up. The interface was not only the bottleneck for speed. It was the bottleneck for success.
So read the numbers with that caveat. They don't prove that Pluno is better at using a previously unseen product. They show what happens when the agent has learned the product, then uses that knowledge on the next task.
We'll soon publish a full writeup with the task set.
There are still cases where UI use is the right tool.
First-time exploration often needs the browser. Visual products need visual context. Canvas tools, design tools, dashboards, maps, and ambiguous workflows may require looking at the screen. Some products may not expose a usable backend contract. Some tasks need visual confirmation.
The point isn't that agents should never use UIs.
The point is that the UI shouldn't be the default execution layer for repeated SaaS operations.
Pixels should be the bootstrap layer, not the runtime.
The interface is the bottleneck
Browser agents aren't mainly slow because models are bad at reasoning.
They're slow because they're using the wrong interface.
We take models that are increasingly competent with code and structured data, then force them through screenshots, buttons, dropdowns, layout shifts, and modal state.
Software already has structured behavior underneath the UI. The hard part is exposing that behavior to the agent safely: learning the contract, redacting traces, validating schemas, bounding permissions, requiring confirmation for side effects, and falling back to visual inspection when needed.
The future isn't agents that get better at clicking buttons.
It's agents that learn the contracts underneath the UI, then use them directly.
If you'd like to try it, here's the agent as a browser extension .
See our latest product updates → Ready to work 10x faster?
Add Pluno to your browser and turn slow web workflows into fast API actions.
Best Zendesk Sell Alternatives in 2026: A Migration Buyer's Guide
Zendesk Sell retires August 31, 2027. Compare the best Zendesk Sell alternatives — Nutshell, Pipedrive, HubSpot, Salesforce — with migration paths and real pricing.
6 Best Ada Alternatives for Better Customer Service in 2026
The 6 best Ada alternatives in 2026, compared on how they learn, handle complex tickets, set up, and price — so you can match the right agent to your team.
Freshdesk vs Zendesk: Which Customer Support Platform Delivers Better Value in 2026?
Freshdesk vs Zendesk in 2026: fast, affordable ticketing versus omnichannel support at scale. Compare pricing, AI billing, setup, and complex-ticket fit.
