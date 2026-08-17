---
source: "https://www.dropstone.io/blog/dropstone-sdk-1-0"
hn_url: "https://news.ycombinator.com/item?id=49329451"
title: "What happens when an AI agent stops being a disposable session?"
article_title: "Introducing Dropstone SDK 1.0: One Memory Across CLI, Chat, and Your Pipelines"
image: "https://www.dropstone.io/blog/dropstone-sdk-1-0/opengraph-image?a781c952d076c140"
author: "DarenWatson"
captured_at: "2026-08-17T12:23:00Z"
capture_tool: "hn-digest"
hn_id: 49329451
score: 2
comments: 0
posted_at: "2026-08-17T11:57:20Z"
tags:
  - hacker-news
  - translated
---

# What happens when an AI agent stops being a disposable session?

- HN: [49329451](https://news.ycombinator.com/item?id=49329451)
- Source: [www.dropstone.io](https://www.dropstone.io/blog/dropstone-sdk-1-0)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T11:57:20Z

## Translation

タイトル: AI エージェントが使い捨てセッションでなくなったらどうなりますか?
記事のタイトル: Dropstone SDK 1.0 の紹介: CLI、チャット、パイプラインにわたる 1 つのメモリ
説明: Dropstone SDK 1.0 は安定しており、現在 npm で利用可能です。これにより、あらゆる TypeScript パイプラインに CLI やチャットと同じクロスサーフェス メモリが提供されます。エージェントに一度、どこでも教えることができ、すべてのセッションがすでにそれを認識しています。 CLI、チャット、VS Code、SDK にわたる 1 つの永続的なアカウント メモリに加え、OpenAI 互換
[切り捨てられた]

記事本文:
Dropstone SDK 1.0 の紹介: CLI、チャット、パイプライン全体で 1 つのメモリを実現 Dropstone プラットフォーム
ブログ / 製品発表 Dropstone SDK 1.0 の紹介: CLI、チャット、パイプライン全体で 1 つのメモリを使用
実際にはどうなるか
パイプラインが記憶すると何が起こるか
本日、Dropstone エージェント ランタイム用の TypeScript SDK の最初の安定バージョンである Dropstone SDK 1.0 をリリースします。
SDK を使用すると、Dropstone CLI およびチャットで実行される同じエージェントにプログラムでアクセスできます。また、SDK には、アカウントごとに 1 つの永続メモリがあり、エージェントが実行されるすべてのサーフェスで共有される、私たちが最も重要だと考えている部分が含まれています。私たちはそれを継続性と呼んでいます。 CLI でエージェントに何かを教えると、SDK から開始されたセッションはすでにそれを認識しています。午前 2 時に CI パイプラインで何かを学習すると、次にチャットを開いたときにその内容が表示されます。そもそもメモリが複数存在することはなかったので、エクスポートするものや同期するものは何もありません。
この投稿では、この方法で構築した理由、パイプラインが記憶できるときに何が変わるか、構築する前に知っておくべき制限について説明します。
AI ツールを毎日使用している場合は、そのルーチンをすでに知っています。月曜日にコーディング エージェントを修正すると、火曜日に別のツールが同じ間違いを犯します。決定が行われたときにチャット ウィンドウがルームになかったため、同じアーキテクチャの概要をチャット ウィンドウに貼り付けます。これ自体は劇的なものではありませんが、積み重なると実際のコストがかかります。使用するすべてのツールは毎回ゼロから始まります。
過去 1 年間の業界の答えは、メモリ プラグインの波でした。エディター間でルールを同期するものもあります。チャット履歴をファイルにエクスポートするものもあります。すべてのツールがセッションの開始時に再読み込みするという共有メモを保持している人もいます。それらは誠実な試みであり、すべてが 1 つの前提を共有しています。

ory はツールに属しており、私たちができる最善のことはそれをコピーすることです。
私たちはその思い込みこそが実際の問題だと考えています。メモリがツール内に存在すると、ツール間に永遠に配管を構築することになります。それで、それを移動させました。
ソフトウェアは以前にもこの動きを行った。アプリケーションの状態は各アプリケーション内に存在しており、チームはコピーの一貫性を保つために無限のツールを構築していました。その後、状態はアプリの下のレイヤーであるデータベースに移行し、一貫性を保つものが何も残っていなかったため、そのツールのほとんどは静かに消滅しました。 ID もシングル サインオンと同様の変化を遂げました。
継続性は、エージェントが学習した内容に同じ考え方を適用します。アカウントごとに 1 つのメモリがあり、アプリではなくランタイムによって保持されます。 CLI はそれを読み取ります。チャットがそれを読み上げます。 VS Code がそれを読み取ります。 SDK はそれを読み取ります。すべてのサーフェスが 1 つのメモリを共有すると、クロスサーフェス メモリは構成する機能ではなくなり、依存するプロパティになります。
これが、SDK が利便性を超えて重要である理由でもあります。 SDK は、エージェントがパイプライン、スケジュールされたジョブ、内部ツール、およびまだ誰も思いつかない場所にアクセスする方法です。これらのセッションが残りの作業と何も共有しなかった場合、私たちはより多くのサイロを作成するためのより迅速な方法を構築したでしょう。これらは同じアカウント メモリを共有するため、新しいサーフェスが追加されるたびにシステム全体がより便利になります。
実際にはどうなるか
ここに示すことができる最も単純なバージョンを示します。昨日、CLI でエージェントを 1 回修正したとします。このリポジトリは npm ではなく bun を使用します。今夜、このスクリプトは CI で、クリーンなランナー上で実行されます。構成はコピーされず、コンテキストも貼り付けられません。
import { createDropstone } を「@blankline/dropstone-sdk」からコピーします
const { クライアント、サーバー } = createDropstone() を待ちます
const session = await client.session.create({
body: { title: "CI 依存関係チェック" },
})
const Reply = 待つ

client.session.prompt({
パス: { id: session.data.id }、
本文: {
Parts: [{ type : "text" , text: "このリポジトリのインストール ステップを追加します。" }]、
}、
})
// エージェントは、このリポジトリが bun を使用していることをすでに知っています。
// 昨日、CLI で一度教えました。
コンソール.ログ(返信.データ)
await server.close() パイプラインが記憶すると何が起こるか
上のデモは意図的に小さくしています。私たちが SDK に注目する理由は、同じプロパティがパイプライン スケールで何を行うかということです。パイプラインは実際には、ユーザーがスリープしている間に実行される単なる表面にすぎないからです。
過去 40 回のデプロイを記憶するリリース パイプラインには、ロールバック規則を再度伝える必要はありません。先月の事件を覚えている監視員は、次の調査を白紙から始めるわけではありません。それは今見ているものを当時見たものと比較します。前回の侵入の試みの形状を記憶しているセキュリティ監視者は、認識が記憶のためにあるため、2 回目の侵入の試みをより早く認識します。エンジニアリングが先週決定した内容を覚えているサポート エージェントは、顧客に先月の回答を提供しなくなりました。
これらのどれも、よりスマートなモデルを必要としません。過去にあった同じモデルです。また、メモリが共有されるため、学習は両方向に複雑になります。パイプラインが夜間に発見した内容は、翌朝 CLI で利用可能になり、朝に CLI に教えた内容は、その夜のパイプラインの動作を形成します。
すべてを保持するメモリは埋め立て地であるため、Continuity は 2 種類のものを保持し、それらを別々に扱います。
ルールは立っての指示です。常にアロー関数を使用してください。無断で依存関係を追加しないでください。ルールは検索されずにすべてのタスクに適用されます。これは、タスクが関連していると思われる場合にのみ表示される設定は、件名が変わると静かに適用を停止するためです。あ

事実は、現時点でたまたま真実であるものです。ステージング データベースは VPN の背後にあり、このプロジェクトはノード 22 を固定します。すべての事実をすべての会話に引き込むとノイズになるため、事実は関連する場合にのみ呼び出されます。
エージェントは、タスクの終了を待つのではなく、ユーザーがレッスンを修正したり、アプローチを拒否したり、好みを述べたりした瞬間にレッスンを記録します。何かがルールなのか事実なのか判断できない場合は、事実を保存します。これは、ルールを見逃した場合は 1 回リマインドする必要があり、間違ったルールがどこにでも付いてくるためです。学んだことをいつでも尋ねることができ、教えたときだけ忘れます。また、メモリが利用できなくなった場合でも、セッションはメモリなしで継続されます。記憶力は利点であるべきであり、決して依存すべきではありません。
CI、自動化、およびサーバーレスの場合、CLI をインストールする必要はまったくありません。ヘッドレス クライアントは API キーを使用して Dropstone API と直接通信し、API はほとんどの AI ツールがすでに使用しているチャット完了形式に従っているため、既存の統合を切り替えるのは通常、ベース URL の変更です。
API キーはアカウント内で作成され、メモリ アクセスはキーが属するアカウントに従います。パイプライン キーから開始されたセッションは、CLI およびチャットと同じアカウント メモリを共有します。これは、キーが同じアカウントへのもう 1 つのドアにすぎないためです。
import { createDropstoneApi } を「@blankline/dropstone-sdk」からコピーします
// 環境から DROPSTONE_API_KEY を読み取ります。
const dropstone = createDropstoneApi()
const resp = await Dropstone.chat.completions.create({
モデル: "dropstone-pro" , // または "dropstone-fast" / "dropstone-heavy"
メッセージ: [
{ role: "user" 、content: "夜間のテストの失敗を要約します。" }、
]、
})
console.log(resp.choices[ 0 ].message.content)
console.log( "Cost: $" + resp.usage?.cost) ライブラリとしての完全なエージェント
2番目

モードには完全なエージェントが埋め込まれます。 1 回の呼び出しで、CLI がローカル サブプロセスとして使用するのと同じランタイムが生成され、それに対する型指定されたクライアント (セッション、ファイル操作、ツール、ストリーミング、構造化出力、および CLI とチャットが使用する同じアカウント メモリ) が渡されます。 CLI を使用して一度サインインすると、すべての SDK セッションが同じメモリを継承します。
同じ実行時間であるため、安全モデルは変更されずに引き継がれます。すべてのツール呼び出し、ファイル編集、およびシェル コマンドは、実行される前に承認ゲートの背後にあります。エージェントの背後にあるモデルはサイクルごとに変化します。この境界は重みではなく実行時に存在するため、存在しません。
何かを 1.0 と呼ぶことは安定性を約束するものであり、それが何を意味するのかを以下に示します。クライアントは、サーバーの OpenAPI 仕様から生成された TypeScript 定義を使用して完全に型指定されます。すべてのリクエスト本文、クエリパラメータ、レスポンスはエディタ内でオートコンプリートされ、文書化された API に 1 対 1 でマッピングされます。
v1 サーフェスは下位互換性を維持します。新しいエンドポイントは、新しい Effect HttpApi コントラクトを反映する v2 サブパス エクスポートに配置されるため、新しいエンドポイントがよりリッチなサーフェスを取得している間、既存の統合は機能し続けます。 SDK は運用環境で使用され、速いペースで出荷されるため、依存するバージョン範囲を固定してください。
その構築方法について知っておく価値のあることが 2 つあります。 SDK のほとんどは Dropstone 自体によって書かれており、現在保守に役立っているのと同じコードベースで動作しており、製品は意図したとおりに動作しています。そしてメモリは、ランタイムの中で私たちが最も積極的に開発を続けている部分です。ここから継続性が向上します。 1.0 は API サーフェスに関する約束であり、メモリ作業が完了したという主張ではありません。
メモリにはサインインが必要です。サインインしていない場合は何も保存されず、メモリ ツールはロードされません。
継続性はあなたが説明したものを保存します

ルール、事実、修正をしっかりと教えます。これはセッションの記録ではなく、セッション終了後はセッションのコンテキストは保持されません。内容をすべて確認したい場合は、尋ねてください。何かをなくしてほしいなら、そう言ってください。
最後に、記憶はエージェントの一貫性を保つものであり、正確なものではありません。間違ったルールは、削除するまで忠実に適用されます。私たちはそれが正しい取引であると考えています。なぜなら代替手段は決して改善されないエージェントだからです。しかし、それは取引であり、目を開いてそれを行う必要があります。
SDK は npm (https://www.npmjs.com/package/@blankline/dropstone-sdk) にあります。以下を使用してインストールします。
Copy npm i @blankline/dropstone-sdk 「メモリは決してツールの機能ではありませんでした。メモリはインフラストラクチャであり、ツールの内部ではなく、ツールの下に属します。」
Dropstone SDK 1.0 は現在 npm で公開されています: npm i @blankline/dropstone-sdk。パッケージは https://www.npmjs.com/package/@blankline/dropstone-sdk にあり、完全なリファレンスは https://docs.dropstone.io/cli/sdk にあり、メモリ モデルは https://docs.dropstone.io/cli/memory に文書化されています。
モデルはレンタル中です。道具はレンタルです。エージェントが学習したことは、スタックの唯一の複合部分であり、それはあなたのものである必要があります。これがツールの下にある 1 つのメモリの意味であり、1.0 に同梱されているものです。
他の人がこの投稿を見つけられるように支援する
投稿する シェアする シェアする Dropstone Architecting Intelligence。
Dropstone は、再帰的推論モデルを安全にスケールするために必要な動的インフラストラクチャを構築します。
プラットフォーム ガバナンス: Dropstone は、Blankline によって開発および運用されている高度な自律コーディング環境です。基礎となるアーキテクチャは、反復的な自己修正を通じて複雑なエンジニアリング タスクを解決するように設計された独自の再帰的推論モデルを利用しています。すべてのモデルの重み、トレーニング データセット、およびソース インフラストラクチャは、引き続き独占的な知的財産です。

f ブランクライン。
商業利用: Dropstone へのアクセスには、エンタープライズ マスター サービス契約が適用されます。プラットフォームのコア ロジックの無断複製、リバース エンジニアリング、再配布は固く禁じられています。請求、ライセンス、およびコンプライアンスに関するすべての問い合わせは、Blankline 法務部に送信する必要があります。
モデルの動作と責任 Dropstone は確率的検索パスを採用し、ソフトウェア合成において人間のベースラインを超えるパフォーマンスを実現します。システムは自律的にエラーを解決できますが、出力は非決定的です。ユーザーは、生成されたコードのレビュー、コンパイル、展開に対する全責任を負います。 Blankline は、運用環境に導入されたダウンストリームの運用上の障害、論理エラー、またはセキュリティの脆弱性に対する責任を負いません。
安全な調整と許容される使用 このプラットフォームには、悪意のあるペイロードや難読化されたエクスプロイトの生成を防ぐように設計された調整レイヤーが統合されています。これらの安全ガードレールをバイパス (「ジェイルブレイク」) しようとしたり、サイバー攻撃活動用のモデルを操作したり、同意のないコンテンツを生成したりする試みは、利用規約の重大な違反となります。違反すると次のような結果が生じます

[切り捨てられた]

## Original Extract

Dropstone SDK 1.0 is stable and available on npm today. It gives any TypeScript pipeline the same cross-surface memory as the CLI and chat: teach the agent once, anywhere, and every session already knows it. One persistent account memory across CLI, chat, VS Code, and SDK, plus an OpenAI-compatible
[truncated]

Introducing Dropstone SDK 1.0: One Memory Across CLI, Chat, and Your Pipelines Dropstone Platform
Blog / Product Launch Introducing Dropstone SDK 1.0: One Memory Across CLI, Chat, and Your Pipelines
What This Looks Like in Practice
What Happens When a Pipeline Remembers
Today we are releasing Dropstone SDK 1.0, the first stable version of the TypeScript SDK for the Dropstone agent runtime.
The SDK gives you programmatic access to the same agent that runs in the Dropstone CLI and chat, and it carries the part we think matters most: one persistent memory per account, shared across every surface the agent runs on. We call it Continuity. Teach the agent something in the CLI and a session started from the SDK already knows it. Learn something in a CI pipeline at 2am and it is there the next time you open chat. There is nothing to export and nothing to sync, because there was never more than one memory in the first place.
This post covers why we built it this way, what changes when a pipeline can remember, and the limits you should know about before you build on it.
If you use AI tools every day, you already know the routine. You correct the coding agent on Monday, and on Tuesday a different tool makes the same mistake. You paste the same architecture summary into a chat window because the chat window was not in the room when the decisions were made. None of this is dramatic on its own, but it adds up to a real cost: every tool you use starts from zero, every time.
The industry's answer over the past year has been a wave of memory plugins. Some sync your rules between editors. Some export your chat history into files. Some keep a shared note that every tool re-reads at the start of a session. They are sincere attempts, and they all share one assumption: that memory belongs to the tool, and the best we can do is copy it around.
We think that assumption is the actual problem. If memory lives inside tools, you end up building plumbing between them forever. So we moved it.
Software has made this move before. Application state used to live inside each application, and teams built endless tooling to keep the copies consistent. Then state moved into databases, a layer below the apps, and most of that tooling quietly disappeared because there was nothing left to keep consistent. Identity went through the same shift with single sign-on.
Continuity applies the same idea to what an agent has learned. There is one memory per account, held by the runtime rather than by any app. The CLI reads it. Chat reads it. VS Code reads it. The SDK reads it. When every surface shares one memory, cross-surface memory stops being a feature you configure and becomes a property you rely on.
This is also why the SDK matters beyond convenience. An SDK is how the agent gets into pipelines, scheduled jobs, internal tools, and places nobody has thought of yet. If those sessions shared nothing with the rest of your work, we would have built a faster way to create more silos. Because they share the same account memory, every new surface makes the whole system more useful.
What This Looks Like in Practice
Here is the simplest version we can show. Suppose yesterday, in the CLI, you corrected the agent once: this repo uses bun, not npm. Tonight this script runs in CI, on a clean runner, with no configuration copied over and no context pasted in:
Copy import { createDropstone } from "@blankline/dropstone-sdk"
const { client, server } = await createDropstone()
const session = await client.session.create({
body: { title: "CI dependency check" },
})
const reply = await client.session.prompt({
path: { id: session.data.id },
body: {
parts: [{ type : "text" , text: "Add the install step for this repo." }],
},
})
// The agent already knows this repo uses bun.
// You taught it once, in the CLI, yesterday.
console.log(reply.data)
await server.close() What Happens When a Pipeline Remembers
The demo above is deliberately small. The reason we care about the SDK is what the same property does at pipeline scale, because a pipeline is really just a surface that runs while you are asleep.
A release pipeline that remembers your last forty deploys does not need to be told your rollback convention again. A monitoring agent that remembers last month's incident does not start the next investigation from a blank page; it compares what it sees now against what it saw then. A security watcher that remembers the shape of a previous intrusion attempt recognizes the second attempt sooner, because recognition is what memory is for. A support agent that remembers what engineering decided last week stops giving customers last month's answer.
None of these require a smarter model. They are the same model with a past. And because the memory is shared, the learning compounds in both directions: something the pipeline discovers overnight is available to you in the CLI the next morning, and something you teach the CLI in the morning shapes what the pipeline does that night.
A memory that keeps everything is a landfill, so Continuity keeps two kinds of things and treats them differently.
A rule is a standing instruction. Always use arrow functions. Never add a dependency without asking. Rules apply to every task without being searched for, because a preference that only shows up when the task looks related will quietly stop applying the moment the subject changes. A fact is something that happens to be true right now: the staging database is behind the VPN, this project pins Node 22. Facts are recalled only when they are relevant, since pulling every fact into every conversation would be noise.
The agent records a lesson the moment you correct it, reject an approach, or state a preference, rather than waiting for the end of a task. When it cannot tell whether something is a rule or a fact, it stores a fact, because a missed rule costs you one reminder while a wrong rule follows you everywhere. You can ask at any time what it has learned, and it forgets only when you tell it to. And if memory is ever unavailable, the session simply continues without it. Memory should be an advantage, never a dependency.
For CI, automation, and serverless, you do not need the CLI installed at all. The headless client talks directly to the Dropstone API with an API key, and the API follows the chat completions shape that most AI tooling already uses, so switching an existing integration over is usually a base URL change.
API keys are created inside your account, and memory access follows the account the key belongs to. A session started from a pipeline key shares the same account memory as your CLI and chat, because the key is just another door into the same account:
Copy import { createDropstoneApi } from "@blankline/dropstone-sdk"
// Reads DROPSTONE_API_KEY from the environment.
const dropstone = createDropstoneApi()
const resp = await dropstone.chat.completions.create({
model: "dropstone-pro" , // or "dropstone-fast" / "dropstone-heavy"
messages: [
{ role: "user" , content: "Summarize the nightly test failures." },
],
})
console.log(resp.choices[ 0 ].message.content)
console.log( "Cost: $" + resp.usage?.cost) The Full Agent, as a Library
The second mode embeds the complete agent. A single call spawns the same runtime the CLI uses as a local subprocess and hands you a typed client against it: sessions, file operations, tools, streaming, structured output, and the same account memory the CLI and chat use. Sign in once with the CLI and every SDK session inherits the same memory.
Because it is the same runtime, the safety model carries over unchanged. Every tool call, file edit, and shell command sits behind an approval gate before it executes. The model behind the agent changes from cycle to cycle. That boundary does not, because it lives in the runtime rather than in the weights.
Calling something 1.0 is a stability promise, so here is what we mean by it. The client is fully typed, with TypeScript definitions generated from the server's OpenAPI specification. Every request body, query parameter, and response autocompletes in your editor and maps one-to-one to the documented API.
The v1 surface stays backward compatible. New endpoints land on the v2 subpath export, which mirrors the newer Effect HttpApi contract, so existing integrations keep working while new ones get the richer surface. The SDK is used in production and ships on a fast cadence, so pin the version range you depend on.
Two things worth knowing about how it is built. Most of the SDK was written by Dropstone itself, working in the same codebase it now helps maintain, which is the product working as intended. And memory is the part of the runtime we are continuing to develop most actively. Continuity gets better from here; 1.0 is a promise about the API surface, not a claim that the memory work is finished.
Memory requires being signed in. If you are not signed in, nothing is stored and the memory tools are not loaded.
Continuity stores what you explicitly teach it: rules, facts, and corrections. It is not a recording of your sessions, and session context is not kept after a session ends. If you want to see everything it holds, ask. If you want something gone, say so.
Finally, memory makes an agent consistent, not correct. A wrong rule will be applied faithfully until you remove it. We think that is the right trade, because the alternative is an agent that never improves, but it is a trade, and you should make it with your eyes open.
The SDK is on npm at https://www.npmjs.com/package/@blankline/dropstone-sdk. Install it with:
Copy npm i @blankline/dropstone-sdk “ Memory was never a feature of a tool. It is infrastructure, and it belongs below the tools, not inside one of them. ”
Dropstone SDK 1.0 is live on npm today: npm i @blankline/dropstone-sdk. The package is at https://www.npmjs.com/package/@blankline/dropstone-sdk, the full reference is at https://docs.dropstone.io/cli/sdk, and the memory model is documented at https://docs.dropstone.io/cli/memory.
Models are rented. Tools are rented. What your agent has learned is the only part of the stack that compounds, and it should belong to you. That is what one memory below the tools means, and it is what 1.0 ships.
Help others discover this post
Post Share Share Dropstone Architecting Intelligence.
Dropstone builds the kinetic infrastructure required to safely scale recursive reasoning models.
PLATFORM GOVERNANCE: Dropstone is an advanced autonomous coding environment developed and operated by Blankline . The underlying architecture utilizes proprietary recursive reasoning models designed to solve complex engineering tasks through iterative self-correction. All model weights, training datasets, and source infrastructure remain the exclusive intellectual property of Blankline.
COMMERCIAL USE: Access to Dropstone is governed by our Enterprise Master Services Agreement. Any unauthorized reproduction, reverse engineering of model behavior, or redistribution of the platform's core logic is strictly prohibited. All billing, licensing, and compliance inquiries must be directed to the Blankline Legal Department.
MODEL BEHAVIOR & LIABILITY Dropstone employs stochastic search paths to achieve performance exceeding human baselines in software synthesis. While the system is capable of autonomous error resolution, output is non-deterministic. The user retains full responsibility for the review, compilation, and deployment of generated code. Blankline disclaims liability for downstream operational failures, logic errors, or security vulnerabilities introduced into production environments.
SAFETY ALIGNMENT & ACCEPTABLE USE This platform integrates alignment layers designed to prevent the generation of malicious payloads or obfuscated exploits. Any attempt to bypass these safety guardrails ("jailbreaking"), manipulate the model for cyber-offensive operations, or generate non-consensual content constitutes a material breach of the Terms of Service. Violations will result in

[truncated]
