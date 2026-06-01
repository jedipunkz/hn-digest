---
source: "https://blog.sentry.io/fixing-javascript-observability/"
hn_url: "https://news.ycombinator.com/item?id=48348326"
title: "Fixing JavaScript observability, one library at a time"
article_title: "Fixing JavaScript observability, one library at a time | Sentry Blog"
author: "awad_dev"
captured_at: "2026-06-01T00:26:45Z"
capture_tool: "hn-digest"
hn_id: 48348326
score: 3
comments: 0
posted_at: "2026-05-31T18:32:50Z"
tags:
  - hacker-news
  - translated
---

# Fixing JavaScript observability, one library at a time

- HN: [48348326](https://news.ycombinator.com/item?id=48348326)
- Source: [blog.sentry.io](https://blog.sentry.io/fixing-javascript-observability/)
- Score: 3
- Comments: 0
- Posted: 2026-05-31T18:32:50Z

## Translation

タイトル: JavaScript の可観測性を一度に 1 つのライブラリで修正
記事のタイトル: JavaScript の可観測性を一度に 1 つのライブラリで修正する |セントリーのブログ
説明: Sentry は、アップストリームの 44 の JavaScript ライブラリに TracingChannel サポートを追加し、脆弱なモンキーパッチをすべてのランタイムで機能するネイティブの可観測性に置き換えます。

記事本文:
メインコンテンツにスキップ
メニュー
プラットフォーム
製品
エラー監視
アプリケーションパフォーマンスの監視
アプリケーションパフォーマンスの監視
隔週の Sentry デモの紹介
Sentry の動作を確認し、エラー、パフォーマンスの問題、コンテキストがどのように組み合わされてバグをより迅速に発見し、自信を持って出荷できるかを学びます。隔週木曜日にご参加ください！
JavaScript の可観測性を一度に 1 つのライブラリで修正する
過去数週間にわたり、私たちは、今日のすべての JavaScript APM ツールを強化している「モンキー パッチ適用」を、ランタイムに組み込まれたものに置き換える、エコシステム横断的な取り組みを推進してきました。その理由、方法、そして現状は次のとおりです。
これはサーバーサイド JavaScript (Node.js、Bun、Deno、Cloudflare Workers) にのみ適用されます。ブラウザーには Diagnostics_channel がなく、ポリフィルに必要な非同期コンテキスト伝播プリミティブがありません。
モンキーパッチはスケールしない
私のチームメイトの Sigrid は、モンキーパッチが失敗する理由と、TracingChannel がそれを解決する方法の詳細な内訳を書きました。
短いバージョン: Sentry を含むすべての JavaScript APM ツールは、実行時に import-in-the-middle (IITM) および require-in-the-middle (RITM) を使用して require() 呼び出しと import 呼び出しをインターセプトすることでライブラリをインストルメント化します。これは ECMAScript モジュール (ESM) と互換性がなく、非 Node ランタイムでは機能せず、バンドラーと競合し、制御できない内部実装の詳細につながります。また、SDK は、インストゥルメンテーションを実行するライブラリが何も実行しない前にロードする必要があります。そうしないと、インストルメンテーションは何も行わずに実行されます。
これは Sentry 固有の問題ではありません。 JavaScript インストルメンテーションを維持しているすべての APM ベンダーは、同じ脆弱性に対処しています。生態系は行き詰まっています。
ほとんどのライブラリ管理者は可観測性について考えていません。彼らは何を公開する必要があるのか​​分からず、OpenTelemetry のようなものを採用するということは、実装を引き受けることを意味します。

標準を追加するだけではなく、負担を軽減します。 APM は何年にもわたってこの問題を解決してきたため、ライブラリ側の誰もそれを理解する必要がありませんでした。
TracingChannels - パッチ適用なしの可観測性
2025 年後半、私たちは Pooya Parsa (Nitro、h3、および unjs エコシステムの作成者) と協力して、Nitro フレームワーク用の Sentry SDK を構築する最善の方法を検討していました。その会話中に、チームメイトの Sigrid が、Node の Diagnostics_channel モジュールの組み込み API である TracingChannel を検討することを提案しました。 Sigrid のブログ投稿では、その API について詳しく説明していますが、中心的な考え方は単純です。ライブラリが構造化イベントを TracingChannel に公開すると、APM ツールは何もパッチを適用せずにそれらのイベントをサブスクライブできます。ライブラリには「クエリが開始されました」と「クエリが終了しました」と表示されるだけで、聞いている人は誰でもそこからスパンを作成できます。
// ライブラリ側 (例: mysql2 内)
'diagnostics_channel' から {tracingChannel } をインポートします。
const queryChannel = tracingChannel ( 'mysql2:query' ) ;
クエリチャネル 。トレースプロミス ( async ( ) => {
接続を待機して戻ります。クエリ(SQL);
} , {クエリ:SQL、サーバーアドレス:ホスト、サーバーポート:ポート});
この追加コードのコストは最小限であるため、ライブラリの管理者にとっては売りやすいものです。 APM 側から見ると、そのトレース チャネルにサブスクライブするだけでイベントを取得できます。 IITM、RITM、ローダーフック、初期化順序はありません。誰も聞いていないときはオーバーヘッドがゼロ。 Node、Bun、Deno 全体で動作します。バンドラーは安全です。この API はノード 18 から利用可能になっており、dc-polyfill はそれが欠けているランタイムをカバーしており、すでにサポート範囲に一致しています。
誰もが同意し、誰も押しつけない
トレース チャネル API とそれを OpenTelemetry で動作させる方法について十分に学んだ後、2025 年 11 月に Otel JS でイシューをオープンし、TracingChann について議論しました。

エルサポート。
反応は肯定的でした。しばらくして、OTel チームの誰かが、TracingChannel を OTel SDK に統合するための API アプローチの草案を作成しました。
しかし、エコシステムの導入を促進する大きな推進力はありません。ドラフトは存在します。エコシステムの働きはそうではありません。
TracingChannel が JavaScript 可観測性の未来であることには誰もが同意しますが、ライブラリにそれを採用させる取り組みをしている人は誰もいません。当社には、TracingChannel のサポートを必要とするデータベース、Web フレームワーク、メッセージ キュー、AI プロバイダーにわたる多数のインストゥルメンテーションがあります。それは山のように上流の PR であり、それぞれがライブラリの内部を理解し、メンテナーが受け入れる提案を作成し、変更を実装し、レビューのフィードバックを繰り返す必要があります。
そこで私は、「それでは、ボールを転がしてみませんか?」と考えました。
最初のステップは、パターンが機能することを証明することでした。以前の SDK 作業の一部として、 h3 、 srvx 、 unstorage 、 db0 、および Nitro で TracingChannel サポートをすでに手動で構築していました。 unjs エコシステムは受容的で迅速に動いてくれたので、イベントの形成方法、コンテキストの伝播フロー、OTel で機能させる方法、および従うべきセマンティック規約など、参考となるサンプルとエンドツーエンドのメンタル モデルが提供されました。
また、私たちは、「TracingChannel を使用するべきです」と言うだけではだめであることも早い段階で学びました。これは単にほこりをかぶるために棚上げされることを懇願しているだけです。代わりに、ニトロの場合と同じように、「ねえ、私たちはあなたのためにそれを行い、あなたがそれを所有するのを手伝います。」と言います。コードをリポジトリに受け入れるとメンテナンスの負担が増えるため、私たちはコードを所有してライブラリの一部にすることを支援します。
それを念頭に置いて、私は pg 、 mysql2 、および redis に連絡して関心を測り、出荷されるまでこれを完全に所有し、出荷後もサポートを提供することを申し出ました。これらはトップのデータベースドライバーライブラリです

エコシステム内でのダウンロード数は合わせて 1 週間あたり 6,000 万件を超えています。 TracingChannel を取得できれば、他のライブラリも取得できます。 3 人全員が「はい」と答え、PR を受けることに前向きでした。
また、Node.js コアの Diagnostics_channel API の作成者である Stephen Belanger にも連絡を取りました。彼は現在、提案に対するフィードバックを提供し、メンテナーを説得​​するために時々必要となる権威の代弁者として、この推進を支援しています。
したがって、私たちはこれをエコシステム全体で一つずつ実現していきます。
これが全体像にどのように当てはまるかについては、次のとおりです。私のチームは SDK ランタイムに依存しないように取り組んでおり、複数のパスを並行して作業しており、そのほとんどは即時に効果があります。 TracingChannel イニシアチブの作業は長期的な取り組みとなります。ユーザーが一夜にして新しいライブラリ バージョンにアップグレードすることは期待できません。また、移行が段階的に行われるように、すべてのユーザーに同時にそれらを実装するよう説得することもおそらく不可能です。
実際の現実は次のとおりです。1 人の人間が 44 のライブラリに TracingChannel サポートを追加しようとしても、実現することはありません。私はそれらの内部のことを知りません。このプロジェクトに携わるまで、私は Redis プロトコルの実装や mysql2 のクエリ パイプラインを一度も見たことがありませんでした。
そこで、SKILLS を介してライブラリごとの重労働を処理する Claude コードを使用してフィードバック ループを構築しました。
研究して提案する。ライブラリ名を指定すると、Claude はその非同期モデル、既存の OTel インストルメンテーション、メンテナンス ステータス、および内部アーキテクチャを調査し、確立したすべてのパターンに従って提案の草案を作成します。どこかに行ってしまう前に見直して調整します。
埋め込む。承認された提案を受け取ると、Claude はテスト、tracePromise/traceCallback の選択、hasSubscribers ガード、Node 18 の互換性、および integ の処理を​​使用して実用的な実装を作成します。

Docker を介した実際のサービスに対する配給テスト。
レビューのフィードバックを取得します。 PR が上流でレビューされると、クロードはすべてのコメントを優先順位付けし、妥当性を評価し、応答を提案し、将来の提案に通知するパターンにフラグを立てます。私は何を行動するかを決定し、メンテナーとのすべてのコミュニケーションを自分で処理します。
トラッカーを更新します。クロードは、すべての上流 PR の最新ステータスを取得し、移行トラッカーを最新の状態に保ちます。
各サイクルは次のサイクルにフィードを与えます。ある図書館のレビュープロセスから学んだことは、次の図書館の提案を改善します。知識は複合化され、今後の作業の指針となる LEARNING.md ファイルにダンプされます。
人間と AI の分割を明確にするために、クロードはリサーチ、定型実装、パターンの適用を担当します。私は、アーキテクチャの決定、挿入ポイントの特定、すべてのメンテナーとのコミュニケーション、および出荷前の各行の最終レビューを処理します。重要なのは、すべてのコミットが共同作成され、AI の関与が透明化されることです。図書館の管理者は AI ではなく人間と対話します。私が特定の部分を人間主導のままにしたのは、メンテナの仕事に敬意を示しているためであり、メンテナにコードをライブラリに採用するよう説得するためにはこれが重要です。
このアプローチにより、数年に及ぶ単独の取り組みが生産ラインに変わり、私は毎日提案を出し続け、並行して実装を開始し、それらすべてから学び、その学びを保留中の仕事や将来の仕事に組み込むことができました。
私たちは 4 つのカテゴリにわたって多くの機器を追跡しています。状況は次のとおりです。
mysql2 - マージされました。 npm エコシステムで最も人気のあるデータベース ドライバーの 1 つ。
node-redis と ioredis - 両方がマージされました。 2 つの主要な Redis クライアントには、TracingChannel サポートが提供されるようになりました。
h3 、 srvx 、 unstorage - すべてマージされました。 unjs エコシステムは初期から熱心でした。これはニトロに触れます。

次に、Nuxt および他のダウンストリーム フレームワークに触れます。
また、e18e のアンブレラ問題や、ライブラリ作成者向けの TracingChannel の使用を標準化するアントレーシング仕様を通じて、エコシステムの調整を確立することも支援しました。
これにより、計測モデルが反転します。図書館は契約を所有しており、私たちはそれに加入しています。上記のすべての問題 (ESM の破損、初期化の順序付け、ランタイム ロックイン、バンドラーの競合) は解消されます。インストルメンテーション コードはよりシンプルになり、ランタイム固有のハックの保守はなくなりました。
これは Sentry だけでなく、すべての APM ツールにも利益をもたらします。これを推進することで、ライブラリのメンテナやより広範なコミュニティとの信頼が築かれるのは確かですが、何人かのメンテナは、このアプローチがすべての人を助け、特定の APM プロバイダに偏っていないため、このアプローチを高く評価していると特に呼びかけています。
ケーススタディとして、node-redis を取り上げます。 Redis チームとの協力中、彼らはすでに独自のファーストパーティ OpenTelemetry インストルメンテーションに取り組んでいました。彼らは、私たちの TracingChannel 提案がその計測器と連携し、強化されることを望んでいました。トレース チャネルを使用して、すでに出荷されているメトリクス プラグインを再実装したところ、テストを 1 つも変更せずに機能しました。現在、私たちは痕跡を残して彼らを支援しています。
mysql2 が TracingChannel サポートをリリースした直後、誰かが独自に mysql2-otel-instrumentation を構築しました。これは、OTel のモンキーパッチが適用された @opentelemetry/instrumentation-mysql2 を置き換える純粋な Diagnostics_channel サブスクライバーです。その動機はまさに私たちが解決しようとしている問題、つまり RITM が機能していなかったということでした。ライブラリにより TracingChannel サポートが追加され、サブスクライバー マニフェストが独自に追加されます。
Express、PostgreSQL (pg)、Knex、GraphQL に対するオープン PR があり、これらのライブラリでは、TracingChannel のサポートにより、何百万ものアプリケーションがその行を変更することなく可観測性が向上します。

自分のコード。 MongoDB、Mongoose、Prisma、Hono については活発な議論が行われており、Koa と Consola については草案を作成しました。リストには、Node の組み込み HTTP モジュール、Kafka クライアント、AI プロバイダー SDK など、まだ対応していないライブラリが 20 以上あります。
個々のライブラリの導入を超えて、次の層は消費者側での重複の削減です。現時点では、TracingChannel をサブスクライブするすべての APM ツールは、ライブラリ ペイロードを OpenTelemetry セマンティック規則に個別にマップする必要があります。私たちは、TracingChannel イベントを標準化されたスパンと属性に変換する、共同管理されるモジュールのセットである共有マッパー レジストリを設計しています。目標は、まずこれを Sentry 内部で構築して実証し、次にオープンソースにして、あらゆる APM ベンダーがプラグインできるようにすることです。ライブラリが TracingChannel サポートを提供し、マッパーが存在する場合、インストルメンテーションは自動になります。
長期的な全体像は、ライブラリが最優先事項としてイベントを発行し、マッパーがコミュニティによって保守され、APM ツールが依存関係にどれだけ創造的にパッチを適用できるかではなく、データをどのように扱うかで競合するエコシステムです。まだそこには達していませんが、フライホイールは回転しています。
トレース チャネルについて話し、図書館での導入を主張することで支援できます。

[切り捨てられた]

## Original Extract

Sentry is adding TracingChannel support to 44 JavaScript libraries upstream, replacing fragile monkey-patching with native observability that works across all runtimes.

Skip to main content
Menu
Platform
Products
Error Monitoring
Application Performance Monitoring
Application Performance Monitoring
Bi-weekly Intro to Sentry Demo
See Sentry in action and learn how errors, performance issues, and context fit together to help you find bugs faster and ship with confidence. Join us bi-weekly on Thursdays!
Fixing JavaScript observability, one library at a time
Over the past few weeks, we have been driving a cross-ecosystem effort to replace the “monkey-patching” that powers all JavaScript APM tools today with something built into the runtime. Here is why, how, and where it stands.
This applies to server-side JavaScript only (Node.js, Bun, Deno, Cloudflare Workers). Browsers do not have diagnostics_channel and lack the async context propagation primitives needed to polyfill it.
Monkey-patching does not scale
My teammate Sigrid wrote a detailed breakdown of why monkey-patching is failing and how TracingChannel solves it .
The short version: every JavaScript APM tool, including Sentry’s, instruments libraries by intercepting require() and import calls at runtime using import-in-the-middle (IITM) and require-in-the-middle (RITM). This breaks with ECMAScript Modules (ESM), does not work in non-Node runtimes, conflicts with bundlers, and couples us to internal implementation details we do not control. The SDK also must load before the library it instruments, or instrumentation silently does nothing.
This is not a Sentry-specific problem. Every APM vendor maintaining JavaScript instrumentation deals with the same fragility. The ecosystem is stuck.
Most library maintainers do not think about observability. They do not know what they would need to expose, and adopting something like OpenTelemetry means taking on an implementation burden, not just adding a standard. APMs managed to patch their way around this for years, so nobody on the library side ever had to figure it out.
TracingChannels - observability without patching
In late 2025, we were working with Pooya Parsa (creator of Nitro, h3, and the unjs ecosystem) on the best way to build a Sentry SDK for the Nitro framework. During that conversation, my teammate Sigrid suggested we look into TracingChannel , a built-in API from Node’s diagnostics_channel module. Sigrid’s blog post covers that API in depth, but the core idea is simple: if a library publishes structured events on a TracingChannel , any APM tool can subscribe to those events without patching anything. The library just says “a query started” and “a query ended,” and whoever is listening can create spans from that.
// Library side (e.g. inside mysql2)
import { tracingChannel } from 'diagnostics_channel' ;
const queryChannel = tracingChannel ( 'mysql2:query' ) ;
queryChannel . tracePromise ( async ( ) => {
return await connection . query ( sql ) ;
} , { query : sql , serverAddress : host , serverPort : port } ) ;
The cost of this added code is minimal, so this is an easy sell for library maintainers. From APM’s side, we just need to subscribe to that tracing channel and we get the events. No IITM, no RITM, no loader hooks, no initialization ordering. Zero overhead when nobody is listening. Works across Node, Bun, and Deno. Bundler safe. The API has been available since Node 18, and dc-polyfill covers runtimes that lack it, which already matches our support range.
Everyone agrees, nobody is pushing
After getting enough learnings about the tracing channel API and how to make it work with OpenTelemetry, I opened an issue on Otel JS in November 2025 to discuss TracingChannel support.
The response was positive. A while after, someone from the OTel team even created a draft API approach for integrating TracingChannel into the OTel SDK.
But there is no significant push to drive ecosystem adoption. The draft exists; the ecosystem work does not.
Everyone agrees that TracingChannel is the future of JavaScript observability, but nobody is doing the work of getting libraries to adopt it. We have many instrumentations across databases, web frameworks, message queues, and AI providers that need TracingChannel support. That is a mountain of upstream PRs, each requiring understanding the library’s internals, writing a proposal that maintainers will accept, implementing the changes, and iterating on review feedback.
So I thought “fine, why not just get the ball rolling?”
The first step was proving the pattern works. I had already built TracingChannel support by hand in h3 , srvx , unstorage , db0 , and Nitro as part of the earlier SDK work. The unjs ecosystem was receptive and moved fast, which gave us shipped examples to point to and an end-to-end mental model: how events should be shaped, how context propagation flows, how to make it work with OTel, and what semantic conventions to follow.
We also learned early that you can’t just say “hey you should use TracingChannel ,” which is just begging to be shelved to collect dust. Instead, like we did with Nitro, we say “Hey, we will do it for you and help you own it.” Accepting code into a repository adds a burden of maintenance, so we offer to help own it and make it part of the library.
With that in mind, I reached out to pg , mysql2 , and redis to gauge their interest, offering to fully own this ‘til it ships and provide support even after. These are the top database driver libraries in the ecosystem, accounting for over 60 million downloads per week combined. If we can get TracingChannel in them, we can get other libraries. All three said yes and were open to receiving a PR.
I also reached out to Stephen Belanger , the creator of the diagnostics_channel API in Node.js core. He is now helping push this forward, providing feedback on proposals and acting as the voice of authority which is sometimes needed to convince maintainers.
So one by one, we’re making this happen across the ecosystem.
For context on how this fits into the bigger picture: My team is working on making our SDK runtime-agnostic, we are working multiple paths in parallel, most of which have an immediate effect. The TracingChannel initiative work is the long-term play. We cannot expect users to upgrade to new library versions overnight, and we probably won’t convince everyone to implement them at the same time so the migration will be gradual.
Here is the practical reality: Being one person trying to add TracingChannel support to 44 libraries is just not going to happen. I do not know the internals of any of them. I have never looked at the Redis protocol implementation or mysql2 ’s query pipeline before this project.
So I built a feedback loop using Claude Code that handles the per-library heavy lifting via SKILLS:
Research and Propose. Given a library name, Claude researches its async model, existing OTel instrumentation, maintenance status, and internal architecture, then drafts a proposal following all the patterns we have established. I review and adjust before it goes anywhere.
Implement. Given an approved proposal, Claude produces a working implementation with tests, handling tracePromise / traceCallback selection, hasSubscribers guards, Node 18 compatibility, and integration tests against real services via Docker.
Capture Review Feedback. When a PR gets reviewed upstream, Claude triages every comment, assesses validity, suggests responses, and flags patterns that should inform future proposals. I decide what to act on and handle all communication with maintainers myself.
Update the Tracker. Claude fetches the latest status of every upstream PR and keeps the migration tracker current.
Each cycle feeds the next one. Learnings from one library’s review process improve the next library’s proposal. The knowledge compounds and is dumped into a LEARNING.md file to guide future work.
To clarify the human/AI split: Claude handles research, boilerplate implementation, and pattern application. I handle architecture decisions, insertion point identification, all maintainer communication, and final review of every line before it ships. Critically, every commit is co-authored and AI involvement is made transparent. Library maintainers interact with a human, not with an AI. I kept certain parts human-led because that shows respect to the maintainer’s work, which is critical to convincing them to adopt code into their library.
This approach turned what would be a multi-year solo effort into a production line where I can keep dishing out proposals every day, start implementations in parallel, learn from them all and integrate the learnings into pending and future work.
We are tracking many instrumentations across four categories. Here is where things stand:
mysql2 - Merged. One of the most popular database drivers in the npm ecosystem.
node-redis and ioredis - Both merged. The two dominant Redis clients now ship TracingChannel support.
h3 , srvx , unstorage - All merged. The unjs ecosystem was early and enthusiastic. This touches Nitro, which in turn touches Nuxt and other downstream frameworks.
We also helped establish ecosystem coordination through an e18e umbrella issue and the untracing spec that standardizes TracingChannel usage for library authors.
This flips the instrumentation model. Libraries own the contract, and we subscribe to it. Every problem described above (ESM breakage, init ordering, runtime lock-in, bundler conflicts) goes away. Our instrumentation code gets simpler, and we stop maintaining runtime-specific hacks.
This also benefits every APM tool, not just Sentry. Driving it builds trust with library maintainers and the broader community, sure, but several maintainers have specifically called out that they appreciate the approach because it helps everyone and is not biased towards any one APM provider.
Take node-redis as a case study. During our collaboration with the Redis team, they were already working on their own first-party OpenTelemetry instrumentation. They wanted our TracingChannel proposal to align with and power that instrumentation. We re-implemented their already shipped metrics plugin using tracing channels and it worked without changing a single test. Now, we are helping them with traces .
Shortly after mysql2 shipped TracingChannel support, someone independently built mysql2-otel-instrumentation , a pure diagnostics_channel subscriber that replaces OTel’s monkey-patched @opentelemetry/instrumentation-mysql2 . The motivation was exactly the problem we are solving: RITM was not working. A library adds TracingChannel support, and the subscribers manifest on their own.
We have open PRs against Express, PostgreSQL (pg), Knex, and GraphQL, the kind of libraries where TracingChannel support means millions of applications get better observability without changing a line of their own code. MongoDB, Mongoose, Prisma, and Hono are in active discussion, and we have drafted proposals for Koa and Consola. There are still 20+ libraries on the list we have not reached out to yet, including Node’s built-in HTTP module, Kafka clients, and AI provider SDKs.
Beyond individual library adoption, the next layer is reducing duplication on the consumer side. Right now, every APM tool that subscribes to a TracingChannel has to independently map library payloads to OpenTelemetry semantic conventions. We are designing a shared mapper registry, a set of co-maintained modules that translate TracingChannel events into standardized spans and attributes. The goal is to build and prove this internally at Sentry first, then open-source it so any APM vendor can plug in. If a library ships TracingChannel support and a mapper exists, instrumentation becomes automatic.
The long-term picture is an ecosystem where libraries emit events as a first-class concern, mappers are community-maintained, and APM tools compete on what they do with the data rather than on how creatively they can patch your dependencies. We are not there yet, but the flywheel is turning.
You can help by talking about tracing channels and advocating for their adoption in the librarie

[truncated]
