---
source: "https://www.openclacky.com/engineering/cache-invalidation-surface"
hn_url: "https://news.ycombinator.com/item?id=48426050"
title: "Every AI Agent Feature Is a Cache Invalidation Surface"
article_title: "Every AI Agent Feature Is a Cache Invalidation Surface | OpenClacky"
author: "gemHunter"
captured_at: "2026-06-06T18:33:39Z"
capture_tool: "hn-digest"
hn_id: 48426050
score: 2
comments: 0
posted_at: "2026-06-06T15:38:12Z"
tags:
  - hacker-news
  - translated
---

# Every AI Agent Feature Is a Cache Invalidation Surface

- HN: [48426050](https://news.ycombinator.com/item?id=48426050)
- Source: [www.openclacky.com](https://www.openclacky.com/engineering/cache-invalidation-surface)
- Score: 2
- Comments: 0
- Posted: 2026-06-06T15:38:12Z

## Translation

タイトル: すべての AI エージェント機能はキャッシュ無効化サーフェイスです
記事のタイトル: すべての AI エージェント機能はキャッシュ無効化の表面である |オープンクラッキー
説明: OpenClucky の背後にある 7 つのエンジニアリング上の決定

記事本文:
オープンクラッキー
GitHub
ロケール#スイッチ"
data-locale-locale-param="zh"
data-locale-target="ボタン"
style="color: hsl(var(--color-ink-2));">中
ロケール#スイッチ"
data-locale-locale-param="en"
data-locale-target="ボタン"
style="color: hsl(var(--color-ink-2));">EN
特長
価格設定
AIキー
クリエイター
エンタープライズ
ドキュメント
サインイン
サインアップ
ドロップダウン#トグル"
class="md:hidden p-2rounded-lg 遷移色"
style="color: hsl(var(--color-ink-2));">
ドロップダウン#閉じる">
ドロップダウン#閉じる"
class="mb-4 p-2rounded-lgtransition-colors">
ドロップダウン#閉じる">
GitHub
言語:
ロケール#スイッチ"
data-locale-locale-param="zh"
data-locale-target="ボタン"
style="color: hsl(var(--color-ink-2));">中
ロケール#スイッチ"
data-locale-locale-param="en"
data-locale-target="ボタン"
style="color: hsl(var(--color-ink-2));">EN
クリエイター
特長
価格設定
AIキー
エンタープライズ
ドキュメント
サインイン
サインアップ
オープンクラッキー
GitHub
2026 年 5 月 18 日
すべての AI エージェント機能はキャッシュ無効化サーフェイスです
Yafei Lee / OpenClucky 創設者
私は Yafei Lee です。Ruby で書かれたオープンソース AI エージェントである OpenClucky の創設者です。私たちは、スキル、メモリ、サブエージェント、ブラウザ自動化、動的なモデル切り替え、および長時間実行セッションを備えたエージェントを必要としていました。これらの機能はそれぞれ、別の方法でプロンプト キャッシュを悪化させました。
それが本当のアーキテクチャの問題でした。 LLM を呼び出す方法、別のツールを追加する方法、より多くのエージェントを調整する方法ではなく、製品が変化し続ける間キャッシュ プレフィックスを安定に保つ方法ではありません。
すべてのエージェント機能はキャッシュ無効化サーフェイスでもあります。スキルは新しいシステム コンテキストを読み込みます。ピア エージェント ワークフローはプレフィックスをフォークします。ブラウザの自動化により、揮発性のツール出力が追加されます。圧縮は歴史を書き換えます。モデル固有の状態がシステム プロンプトに表示されない限り、モデルの切り替えによりキャッシュ名前空間が断片化される可能性があります。有能なエージェントを構築している場合、

キャッシュ ヒット率が予想よりもはるかに低いのは、おそらくこれが理由です。
2 年と 3 世代のアーキテクチャ (最初の 2 世代は失敗) を経て、私たちは 7 つのエンジニアリング上の決定に集中し、これらの機能をすべてそのまま維持しながら、実際のタスク全体で 90% 以上のキャッシュ レートを達成できるようにしました。以下に、何が壊れたのか、何を試したのか、そして実際に何がうまくいったのか、その完全なストーリーを示します。
第 1 世代: RAG Everything (2024 ～ 2025 年初頭)
私たちの最初のエージェントは教科書的な RAG システムでした。ユーザーのコードベース、ドキュメント、会話履歴をベクター ストアに埋め込みました。 LLM が何かを認識する前に、すべてのクエリはハイブリッド検索、再ランキング、およびクエリの書き換えを経ます。
コストの上昇は止まらず、データは常に古いままでした。コードベースを更新するたびに再埋め込みが必要でした。リアルタイム同期は信頼性が低いため、ベクター ストアは実際のコードより遅れていました。私たちは、ますます間違っているインデックスを検索するために、ますます多額の費用を支払っていました。
そして、90% の再現率では十分ではありません。 10 回の検索で間違ったコンテキストが返されました。複数のステップを連鎖させるエージェントの場合、そのエラーは急速に悪化します。ステップ 2 でファイルが間違っているということは、ステップ 3 で編集が間違っているということは、ステップ 4 で無駄な再試行が行われることを意味します。私たちは、97% の再現率がエージェントがネットポジティブであるための最低値であると推定しましたが、それには遠く及ばませんでした。それに加えて、ベクター データベースは、クラッシュ、遅延、またはガベージを返す可能性のあるもう 1 つのコンポーネントでした。ユーザーと LLM の間にある余分な部分はすべて、遅延が隠れ、エラーが複合する場所です。
ローカル リポジトリ上で動作するコーディング エージェントについては、RAG を完全に停止しました。埋め込み、ベクトル ストア、取得パイプラインはありません。エージェントがコンテキストを必要とする場合、ファイルを直接読み取るか、 grep を使用して検索します。エージェントがドキュメントにアクセスできるようにする必要がある場合は、Web サイト上でドキュメントを読めるようにしてください。埋め込みに細断しないでください。
ゲ

世代 2: マルチエージェント オーケストレーション (2025 年中頃)
次のアイデアは、SWEBench リーダーボード プレイブックから直接得られたものです。プランナー エージェント、コーダー エージェント、レビューアー エージェント、およびテスター エージェントはすべて、役割固有のプロンプトを備えたメッセージ バスを通じて調整されます。
SWEBench ではまともなスコアが得られました。製品はひどいものでした。
エージェントのハンドオフはすべてキャッシュミスでした。各サブエージェントには独自のシステム プロンプトとキャッシュ名前空間がありました。エージェント間でコンテキストを渡すことは、状態をメッセージにシリアル化することを意味し、ハンドオフのたびに受信側エージェントのキャッシュ プレフィックスが消去されます。問題はキャッシュミスだけではありませんでした。ハンドオフのたびに、リッチな状態をより小さなメッセージにシリアル化する必要があり、有用なコンテキストが境界で失われました。
エージェント 1 人では 4 分で完了できるタスクが、エージェント 4 人では 14 分かかりました。調整のオーバーヘッドは現実のものでした。エージェントは互いに待機し、前のエージェントがすでに処理したコンテキストを再読み取りし、場合によっては互いの決定に矛盾しました。
コストは6倍高かった。 4 つの個別のキャッシュ名前空間、4 つのシステム プロンプト、定数シリアル化。人間のチームでは機能する「専門家間で仕事を分割する」という直感は、LLM には伝わりません。単一のフロンティアモデルはすでにジェネラリストです。分業しているわけではありません。オーバーヘッドが増大しています。
デバッグは悪夢でした。最終出力が間違っていた場合、その原因はどのエージェントにありましたか?プランナーは曖昧な指示を出しましたか？コーダーがそれらを誤解したのでしょうか？レビュー担当者はバグを見逃したのでしょうか?私たちは、元のタスクに費やした時間よりも、パイプラインを通じて障害を追跡することに多くの時間を費やしました。少なくとも 1 人のエージェントの場合、何か問題が発生した場合、会話を 1 つ読んで間違いを見つけます。
SWEBench スコアはユーザーの満足度を予測しませんでした。特定のベンチマークに合格するようにマルチエージェント パイプラインを調整することはできますが、実際のユーザーを悩ませる障害モード (

反復回数の少なさ、ハンドオフ間でのコンテキストの喪失、一貫性のないコード スタイルなど）はベンチマークの測定対象ではありませんでした。
役割ベースのマルチエージェント オーケストレーションを廃止しました。 1 つのメイン エージェント、1 つの会話、1 つのキャッシュ名前空間。サブエージェントは、単一の安定したツールを通じて呼び出される、分離されたスキル実行コンテキストとしてのみ存続しました。
2 世代にわたって同じ結論: このモデルはすでに十分に賢くなっています。必要なのはモデルの追加ではなく、より優れたハーネスです。
第 3 世代は、単一エージェントのキャッシュ ヒット率を中心にすべてを最適化したらどうなるかという疑問から始まりました。コストハックとしてではなく、アーキテクチャ上の原則として。キャッシュ ヒット数が多いということは、モデルが一貫したコンテキストを認識し、より高速に応答し、コストが低いことを意味します。以下のすべての決定はその目標に役立ちます。
(コードはオープンソースです。各決定を実装する正確なファイルへのリンクは、この投稿の最後にあります。)
決定 1: 歴史の成長によりプレフィックス マッチングが中断される → ダブル キャッシュ マーカー
プロンプト キャッシュはプレフィックス マッチングによって機能します。 LLM プロバイダーは、メッセージ プレフィックスのハッシュを保存します。次のリクエストがそのプレフィックスを共有する場合、キャッシュされたレートが得られます (プロバイダーによっては、キャッシュされたトークンの価格が通常の入力トークンの数分の一に設定されます)。プロバイダーにキャッシュする場所を指示する方法は、特定のメッセージにcache_control マーカーを配置することです。
単純なアプローチは、最後のメッセージの 1 つのマーカーです。それは次の 3 つの方法で壊れます。
歴史は単調に成長します。メッセージ N をマークします。次のターン、メッセージ N+1 が追加されます。古いマーカーの位置のコンテンツが変更されているため、履歴全体のキャッシュ ミスです。
ツール呼び出しが再試行されます。モデルの最後のツール呼び出しでエラーが発生するか、ユーザーが Ctrl-C を押します。 「最後のメッセージ」は破棄され、マーカーも一緒に消えます。
セッション中にモデルが切り替わります。ユーザーは Sonnet から Opus に切り替えます。できるだけ多くのプレフィックスを共有したい場合

ロスモデル。不必要なマーカーの移動はキャッシュ ミス イベントになります。
まず問題 (1) に取り組みます。修正の進行状況は git ログに表示されます。
8ff66cc 修正: キャッシュ
6ea99fe 修正: プロンプト キャッシュ
e9a3602 特技: プロンプト キャッシュが正常に動作する
7734c97 特技: 2 ポイント キャッシュを試す
最初の 3 つのコミットは増分パッチでした。最後は構造的な修正で、マーカーを 1 つではなく 2 つにしました。
毎ターン、1 つではなく 2 つの連続したメッセージをマークします。
ターン N: [..., msg_A, msg_B(*), msg_C(*)]
↑ ↑
マーカー1 マーカー2
ターン N+1: [..., msg_A, msg_B(*), msg_C(*), msg_D(*)]
↑ ↑ ↑
(まだあります) (まだあります) 新しいマーカー
ターン N+1 で、プロバイダーは msg_C のマーカーと一致しようとし、その前のすべて (システム プロンプト + ツール + 完全な履歴から最後のメッセージを除く) をヒットします。次のターンのために msg_D に新しいマーカーを置きます。
これはローリング ダブル バッファです。常に 2 つのブレークポイントを保持します。1 つは (前のターンから) 「読み取り」、もう 1 つは (現在の末尾で) 「書き込み」中です。次のターンでは、古い「書き込み」が新しい「読み取り」になり、新しい末尾に新しいものを書き込みます。両方のバッファーが同時に無効になることはありません。
マーカーを追加するたびに、書き込み層の価格でキャッシュ書き込みのコストがかかります。カバーする必要がある唯一の障害境界は「古いテール/新しいテール」エッジであり、2 つのマーカーはまさにその最小値です。 3 番目のマーカーはプレフィックスのさらに後方に配置され、独立して読み取られることのないセグメントを書き込みます。 2 は境界をカバーします。 3は冗長です。
これが 2 番目の利点であり、 commit 7734c97 の背後にある実際の動機です。モデルがツール呼び出しを再試行すると (エラー、Ctrl-C、ストリームの破損)、最後のメッセージは破棄されます。マーカーが 1 つある場合は、即時にキャッシュ ミスになります。二重マーカーの場合、通常は最後から 2 番目のマーカーが存続するため、シングルステップのロールバックは依然として有効です。

そのキャッシュ。 3 つのマーカーは 2 段階のロールバックでも生き残ることができますが、コストがエッジ ケースに見合ったものではありません。
マーカー選択ロジックには、 system_injected: true のタグが付いたメッセージをスキップするという 1 つの厳しいルールがあります。これらは一時的なメッセージ (セッション コンテキスト ブロック、圧縮命令) であり、次のターンには同じ形式では存在しません。それらのマーカーは、決して読み戻されない書き込みです。セレクターは末尾から逆方向に進み、system_injected メッセージをスキップし、実際の会話メッセージが 2 つある時点で停止します。
決定 2: 動的セッション状態によりシステム プロンプトが中断される → システム プロンプトが凍結される
エンジニアリング分野: エージェントのシステム プロンプトはセッション開始時に一度作成され、その後バイト フリーズされます。システム プロンプトに動的情報を入力するという要件は、他の場所にリダイレクトされます。
これはキャッシュ戦略全体の基礎です。システム プロンプトが変更されると、それ以降のキャッシュ エントリはすべて無効になります。部分的な修正はありません。
ただし、少なくとも 4 種類の情報がシステム プロンプトに当然「存在」する必要があります。
現在の日付、作業ディレクトリ、OS - モデルは正しいコマンドを実行するためにこれらを必要とします。
現在のモデル ID — 自己適応動作に役立ちます。
新しくインストールされたスキル — モデルは、それらを呼び出すためにスキル名を確認する必要があります。
更新されたユーザー設定 (USER.md / SOUL.md) — エージェントの性格とユーザー コンテキスト。
4 つすべてがセッション中に変更できます。これらのいずれかがシステム プロンプトにある場合は、1 回の変更ですべてが無効になります。
システム プロンプトの代わりに、この情報を通常のユーザー メッセージとして会話履歴に挿入します。
[セッションの内容: 今日は 2026 年 5 月 13 日、火曜日です。現在のモデル: クロード-ソネット-4-6。
OS：macOS。作業ディレクトリ: /Users/.../project]
このメッセージには system_injected: true というタグが付けられています。キャッシュ マーカーによって選択されず (決定 1)、としてカウントされません。

実際のユーザーターンであり、圧縮中に破棄されます。注入は日付によって制御されます。1 日あたり 1 回、さらにモデル切り替え時に 1 回です。ほとんどのセッションでは 1 つだけが表示されます。
inject_session_context の最初の実装は熱心でした。これは、システム プロンプトが構築される前に、エージェントの構築中に発生しました。これは @history.empty という意味ですか? false を返したため、 run() はシステム プロンプトの構築を完全にスキップしました。エージェントは、「今日は火曜日です」というメッセージを含む最初のリクエストを送信しましたが、システム プロンプトはありませんでした。追跡するまで 1 日間、動作が微妙に崩れていました。
修正は「システム プロンプトが構築された後に挿入する」という 1 行でした。生き残ったコードコメント:
# 重要: システム プロンプトがまだ構築されていない場合は、インジェクションをスキップします。
# それ以外の場合、空の履歴にユーザーメッセージを追加すると、
# @history.empty? false の場合、run() はビルドをスキップします。
# システムプロンプト全体。
組み立て順序は内容よりも重要です。プレフィックスの各部分の設計には何週間も費やすことができますが、アセンブリ順序が 1 ステップでも間違っている場合、キャッシュ戦略全体が無効になります。
スキルはセッション開始時にシステム プロンプトに表示され、その後フリーズされます。セッション中にインストールされたスキルは、次のセッションまで表示されません。私たちはこの摩擦を受け入れます。すべてのスキーでシステム プロンプトを再レンダリングする

[切り捨てられた]

## Original Extract

Seven engineering decisions behind OpenClacky

OpenClacky
GitHub
locale#switch"
data-locale-locale-param="zh"
data-locale-target="button"
style="color: hsl(var(--color-ink-2));">中
locale#switch"
data-locale-locale-param="en"
data-locale-target="button"
style="color: hsl(var(--color-ink-2));">EN
Features
Pricing
AI Keys
Creators
Enterprise
Docs
Sign In
Sign Up
dropdown#toggle"
class="md:hidden p-2 rounded-lg transition-colors"
style="color: hsl(var(--color-ink-2));">
dropdown#close">
dropdown#close"
class="mb-4 p-2 rounded-lg transition-colors">
dropdown#close">
GitHub
Language:
locale#switch"
data-locale-locale-param="zh"
data-locale-target="button"
style="color: hsl(var(--color-ink-2));">中
locale#switch"
data-locale-locale-param="en"
data-locale-target="button"
style="color: hsl(var(--color-ink-2));">EN
Creators
Features
Pricing
AI Keys
Enterprise
Docs
Sign In
Sign Up
OpenClacky
GitHub
May 18, 2026
Every AI Agent Feature Is a Cache Invalidation Surface
Yafei Lee / Founder of OpenClacky
I'm Yafei Lee, founder of OpenClacky , an open-source AI agent written in Ruby. We wanted an agent with skills, memory, sub-agents, browser automation, dynamic model switching, and long-running sessions. Each of those features made prompt caching worse in a different way.
That was the real architecture problem. Not how to call an LLM, not how to add another tool, not how to orchestrate more agents — how to keep the cache prefix stable while the product keeps changing.
Every agent feature is also a cache invalidation surface. Skills load new system context. Peer-agent workflows fork the prefix. Browser automation adds volatile tool output. Compression rewrites history. Model switching can fragment the cache namespace unless model-specific state stays out of the system prompt. If you're building a capable agent and your cache hit rate is much lower than expected, this is probably why.
Over two years and three architecture generations (the first two failed), we converged on seven engineering decisions that let us hit 90%+ cache rates across real tasks — while keeping all those features intact. What follows is the complete story: what broke, what we tried, and what actually worked.
Generation 1: RAG Everything (2024 – early 2025)
Our first agent was a textbook RAG system. We embedded the user's codebase, docs, and conversation history into a vector store. Every query went through hybrid retrieval, re-ranking, and query rewriting before the LLM saw anything.
The costs never stopped climbing, and the data was always stale. Every codebase update required re-embedding. Real-time sync was unreliable, so the vector store lagged behind the actual code. We were paying more and more to search an index that was increasingly wrong.
And 90% recall is not good enough. One in ten retrievals returned the wrong context. For an agent that chains multiple steps, that error compounds fast. A wrong file in step 2 means a wrong edit in step 3 means a wasted retry in step 4. We estimated that 97% recall might be the bare minimum for an agent to be net-positive, and we were nowhere close. On top of that, the vector database was one more component that could crash, lag, or return garbage. Every extra piece between the user and the LLM is a place where latency hides and errors compound.
For coding agents working over local repos, we killed RAG entirely. No embeddings, no vector store, no retrieval pipeline. If the agent needs context, it reads files directly or searches with grep . If your documentation needs to be accessible to an agent, make it readable on a website. Don't shred it into embeddings.
Generation 2: Multi-Agent Orchestration (mid-2025)
The next idea was straight from the SWEBench leaderboard playbook: a Planner agent, a Coder agent, a Reviewer agent, and a Tester agent, all coordinated through a message bus with role-specific prompts.
We got decent SWEBench scores. The product was terrible.
Every agent handoff was a cache miss. Each sub-agent had its own system prompt and cache namespace. Passing context between agents meant serializing state into messages, and every handoff wiped the receiving agent's cache prefix. The problem was not just cache misses. Each handoff forced us to serialize rich state into a smaller message, and useful context was lost at the boundary.
A task that one agent could finish in 4 minutes took 14 minutes with four. The coordination overhead was real: agents waited for each other, re-read context the previous agent had already processed, and occasionally contradicted each other's decisions.
Cost was 6× higher. Four separate cache namespaces, four system prompts, constant serialization. The "divide work among specialists" intuition that works for human teams doesn't transfer to LLMs. A single frontier model is already a generalist. You're not dividing labor; you're multiplying overhead.
Debugging was a nightmare. When the final output was wrong, which agent caused it? The Planner gave ambiguous instructions? The Coder misinterpreted them? The Reviewer missed the bug? We spent more time tracing failures through the pipeline than we spent on the original task. At least with a single agent, when something goes wrong, you read one conversation and find the mistake.
SWEBench scores didn't predict user satisfaction. We could tune the multi-agent pipeline to pass specific benchmarks, but the modes of failure that annoyed real users (slow iteration, losing context across handoffs, inconsistent code style) weren't what benchmarks measured.
We killed role-based multi-agent orchestration. One main agent, one conversation, one cache namespace. Sub-agents survived only as isolated skill execution contexts, invoked through a single stable tool.
Two generations, same conclusion: the model is already smart enough. What it needs isn't more models, it's a better harness.
Generation 3 started from a question: what if we optimized everything around a single agent's cache hit rate? Not as a cost hack, but as an architectural principle. High cache hits mean the model sees consistent context, responds faster, and costs less. Every decision below serves that goal.
(The code is open source. Links to the exact files implementing each decision are at the end of this post.)
Decision 1: History Growth Breaks Prefix Matching → Double Cache Markers
Prompt caching works by prefix matching. The LLM provider stores a hash of the message prefix; if your next request shares that prefix, you get the cached rate (depending on the provider, cached tokens are priced at a fraction of normal input tokens). The way you tell the provider where to cache is by placing cache_control markers on specific messages.
The naive approach is one marker on the last message. It breaks in three ways:
History grows monotonically. You mark message N. Next turn, message N+1 is appended. The content at the position of your old marker has changed, so it's a cache miss on the entire history.
Tool call retries. The model's last tool call errors out, or the user hits Ctrl-C. The "last message" gets discarded, and your marker vanishes with it.
Mid-session model switches. The user switches from Sonnet to Opus. You want to share as much prefix as possible across models. Any unnecessary marker movement becomes a cache miss event.
We hit problem (1) first. The fix progression is visible in our git log:
8ff66cc fix: cache
6ea99fe fix: prompt cache
e9a3602 feat: prompt cache works fine
7734c97 feat: try 2 point cache
The first three commits were incremental patches. The last one was the structural fix: two markers instead of one.
Every turn, we mark two consecutive messages, not one:
Turn N: [..., msg_A, msg_B(*), msg_C(*)]
↑ ↑
marker 1 marker 2
Turn N+1: [..., msg_A, msg_B(*), msg_C(*), msg_D(*)]
↑ ↑ ↑
(still there) (still there) new marker
On turn N+1, the provider tries to match the marker on msg_C and hits everything before it (system prompt + tools + full history minus the last message). We place a new marker on msg_D for the next turn.
This is a rolling double buffer : at any moment we hold two breakpoints — one being "read" (from the previous turn) and one being "written" (at the current tail). Next turn, the old "write" becomes the new "read," and we write a fresh one at the new tail. There's never a moment where both buffers are invalid simultaneously.
Each additional marker costs a cache write at write-tier pricing. The only failure boundary we need to cover is the "old tail / new tail" edge, and two markers is exactly the minimum for that. A third marker lands further back in the prefix, writing a segment that will never be read independently. 2 covers the boundary. 3 is redundant.
This is the second benefit, and the actual motivation behind commit 7734c97 . When the model retries a tool call (error, Ctrl-C, broken stream), the last message gets discarded. With a single marker, that's an immediate cache miss. With double markers, the second-to-last marker usually survives, so single-step rollback still hits cache. Three markers would survive two-step rollbacks, but the cost doesn't justify the edge case.
Our marker selection logic has one hard rule: skip any message tagged system_injected: true . These are ephemeral messages (session context blocks, compression instructions) that won't exist in the same form next turn. A marker on them is a write that will never be read back. The selector walks backward from the tail, skips system_injected messages, and stops when it has two real conversation messages.
Decision 2: Dynamic Session State Breaks System Prompts → Frozen System Prompt
Engineering discipline: our agent's system prompt is built once at session start, then byte-frozen. Any requirement to put dynamic information in the system prompt gets redirected elsewhere.
This is the foundation of the entire cache strategy. If the system prompt changes, every subsequent cache entry is invalidated. There is no partial fix.
But at least four kinds of information naturally "want" to live in the system prompt:
Current date, working directory, OS — the model needs these for correct commands.
Current model ID — helpful for self-adaptive behavior.
Newly installed skills — the model needs to see skill names to invoke them.
Updated user preferences (USER.md / SOUL.md) — the agent's personality and user context.
All four can change mid-session. If any of them is in the system prompt, a single change invalidates everything.
Instead of the system prompt, we inject this information as a regular user message in the conversation history:
[Session context: Today is 2026-05-13, Tuesday. Current model: claude-sonnet-4-6.
OS: macOS. Working directory: /Users/.../project]
This message is tagged system_injected: true . It won't be selected by cache markers (Decision 1), won't count as a real user turn, and gets discarded during compression. Injection is date-gated: one per day, plus one on model switch. Most sessions see exactly one.
Our first implementation of inject_session_context was eager. It fired during agent construction, before the system prompt was built. This meant @history.empty? returned false , so run() skipped system prompt construction entirely. The agent sent its first request with a "today is Tuesday" message but no system prompt. Behavior was subtly broken for a day before we traced it.
The fix was one line: inject after the system prompt is built. The code comment that survived:
# IMPORTANT: Skip injection when the system prompt hasn't been built yet.
# Otherwise, appending a user message to an empty history makes
# @history.empty? false, which causes run() to skip building the
# system prompt entirely.
Assembly order matters more than content. You can spend weeks designing each piece of the prefix, but if the assembly sequence is wrong by one step, the entire cache strategy is void.
Skills are rendered into the system prompt at session start, then frozen. A skill installed mid-session won't appear until the next session. We accept this friction. Re-rendering the system prompt on every ski

[truncated]
