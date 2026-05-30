---
source: "https://www.appfactor.io/blog/how-we-made-getprocinfo3-agent-readable-semantic-discovery-ai-enrichment"
hn_url: "https://news.ycombinator.com/item?id=48323441"
title: "Agents aren't the problem – Existing systems and API's were not built for AI"
article_title: "AppFactor"
author: "keithneilson"
captured_at: "2026-05-30T11:47:04Z"
capture_tool: "hn-digest"
hn_id: 48323441
score: 1
comments: 1
posted_at: "2026-05-29T14:18:45Z"
tags:
  - hacker-news
  - translated
---

# Agents aren't the problem – Existing systems and API's were not built for AI

- HN: [48323441](https://news.ycombinator.com/item?id=48323441)
- Source: [www.appfactor.io](https://www.appfactor.io/blog/how-we-made-getprocinfo3-agent-readable-semantic-discovery-ai-enrichment)
- Score: 1
- Comments: 1
- Posted: 2026-05-29T14:18:45Z

## Translation

タイトル: エージェントは問題ではありません – 既存のシステムと API は AI 向けに構築されていません
記事タイトル: AppFactor

記事本文:
アプリファクター
プラットフォームの使用例 レガシー アプリのモダナイゼーション ビジネスを中断することなく、レガシー システムをアップグレードします。発見と文書化 複雑なシステムを明確で実用的な文書に変えます。クラウドと VMware の移行 最小限のダウンタイムと完全な制御でクラウドに移行します。ビルドとデプロイ 合理化されたビルドとデプロイのワークフローにより、より迅速に出荷できます。 Agentic ソフトウェア エンジニアリング チーム AI 主導のエンジニアリングでチームを拡張します。会社概要 採用情報 パートナー クラウド変革 ブログ MCP Bridge プラットフォームのブログ MCP Bridge Company
MCP Bridge パート 3: getProcInfo3() をエージェントが読み取り可能にした方法: ハイブリッド ディスカバリ + AI エンリッチメント
前回の記事では、API サーフェスが大きい場合に MCP ツール カタログ全体を置き換える 3 つのメタツールであるコード モードについて説明しました。これら 3 つのメタツールのうちの 1 つ目は search_tools です。今日はそれを開封していきます。
search_tools は、LLM エージェントと 200 オペレーションの API サーフェスの間に立つものです。エージェントが実行したいことを自然言語で記述し、それを実際に実行できる 3 つまたは 4 つのツールを返す必要があります。これを間違えると、エージェントは無関係なツールをバタバタ使ったり、さらに悪いことに自信を持って間違ったツールに電話をかけたりすることになります。
これは MCP Bridge の簡単な部分だと考えました。そうではありませんでした。
最初に試したことと、なぜうまくいかなかったのか
最初のカットは純粋なベクトル検索でした。すべてのツールの名前と説明を OpenAI 埋め込みモデルに埋め込み、pgvector に保存し、実行時にクエリします。教科書のデータセットでは問題なく動作します。同様の埋め込みを行うが目的が異なる 2 つのツール ( get_customer vs get_customer_full vs get_customer_with_orders ) を使用した瞬間に、エンタープライズ API に影響が及びます。
次に全文検索を試してみました。同義語辞書を備えた Postgres FTS。完全一致では良いが、CAS では悪い

エージェントの意図がツールの説明と語彙を共有していない場合。
着地したのはハイブリッドです。Postgres FTS が最初のパス (BM25 スタイルのランキングによる) を実行し、pgvector が上位の FTS 候補とより広範なセマンティックのみのプールに対して 2 番目のパスを実行し、小さなリランカーが 2 つの順序付けされたリストを 1 つのスコアにまとめます。再ランカーは、エージェントの最近のコンテキストも考慮に入れます。エージェントが顧客データを処理している場合、顧客の形をしたツールが上位にランクされます。
ソースデータが良好であれば、これはうまく機能します。
エンタープライズ API は LLM にちなんで名付けられていません。これらの名前は、多くの場合 10 年前に、それを作成したエンジニアの名前にちなんで付けられていますが、その命名規則は今では忘れ去られています。実際のお客様の実例:
getProcInfo3(custId, opt) → オブジェクト
説明: 「ドキュメントを参照してください。」
署名だけではダメです。 custId はかすかなヒントであり、opt は何かを意味する可能性があり、戻り値の型オブジェクトは何も教えてくれません。その入力が生き残るセマンティック検索メソッドはありません。
一方、応答にはシグナルが含まれています。
{
"custId": "C-44218",
"billingAddress": { "..." },
"アカウントの状態": "アクティブ",
"primaryContact": { "..." },
"assignedManager": "..."
}
これがツールが実際に行うこと、つまり顧客アカウントの検索、請求および連絡先レベルの詳細です。応答の形状によって、これが何であるかがわかります。
私たちは、お客様に API 名を修正するよう伝えることを検討しました。これは小規模では現実的な解決策ですが、重要な唯一の規模では不可能な解決策です。 70 のレガシー サービスを利用しているお客様は、私たちが求めたので API の名前を変更するつもりはありません。
MCP Bridge を OpenAI 互換のチャット完了エンドポイントにポイントします。当社は社内で Azure OpenAI を使用しています。互換性のあるプロキシ経由の Anthropic は機能します。 vLLM 経由のローカル Llama は機能します。サービスとプラットフォームでエンリッチメントを有効にします。
ツールごとに集めます

■ 使用可能なシグナル: 名前、(空の場合が多い) 説明、パラメータ・スキーマ、応答タイプ定義 (OpenAPI、WSDL、または .proto (使用可能な場合) から取得)、および不透明な API の場合は、プローブ呼び出しまたはトレースされた運用トラフィックからキャプチャされたサンプル応答。
より明確な名前、ツールの動作とエージェントがいつ使用するかを示す 1 ～ 2 文の説明、および 3 ～ 5 個のタグまたはエイリアスのリストを生成するようモデルに要求する構造化プロンプトを送信します。
スキーマに対して出力を検証します (幻覚パラメータや入出力タイプのセマンティック ドリフトがないこと)。
強化されたメタデータを元のスキーマとともに保存します。ツール呼び出しでは引き続き元の名前とスキーマが使用され、検出レイヤーのみが強化されたバージョンを参照します。
最も重要な信号は応答形状です。名前と説明だけでエンリッチメントを試みましたが、モデルでは getProcInfo3 と getProcInfo4 を区別する方法がありませんでした。応答 (またはサンプル) が確認できると、関数は読みやすくなります。 SOAP サービスの場合、通常、WSDL によって応答タイプが提供されます。文書化されていない API の場合、MCP Bridge はワンタイム プローブを実行するか、トレースからプルすることができます。
重要なことは、エンリッチメントはサービスごとにオプトインされており、元のスキーマが保持されることです。エージェントには、名前が変更されたパラメータが表示されることはありません。私たちは、ソース API から静かに分岐するシステムを望んでいませんでした。
これは、実際の (匿名化された) 顧客 SOAP サービスからの getProcInfo3 です。エンリッチメント モデルでは、元の署名、空の説明、および上記の応答サンプルが確認されました。応答フィールド (custId、billingAddress、accountStatus、primaryContact、assignedManager) から次のことが推測されます。
説明: 「ドキュメントを参照してください。」
名前: get_customer_account_details
説明: 「請求先住所、アカウントのステータス、主な連絡先、およびアカウントを含む顧客のアカウント プロファイルを取得します。

割り当てられたアカウントマネージャー。サポート、請求、販売活動を開始する前にアカウントの状態を確認するのに役立ちます。」
エイリアス: customer_lookup、account_info、customer_profile、billing_account
昨日のベンチマークで使用したエージェント タスク スイートのツール選択精度は、強化後のこの顧客のサービスで 41% から 89% に向上しました。 REST サービスにはすでに適切な名前が付いていたため、SOAP スイートの下の数値はそれほど劇的ではありませんでしたが、そこでは 71% → 88% の増加でした。教訓: 充実させることは、他に何もすることが最も難しい場所において最も重要です。
エンリッチされた説明が変更されると、埋め込みが再生成されます。元の埋め込みも保持されます。取得時にそれらを組み合わせた場合 (加重、オリジナル × 0.3 + 強化 × 0.7) が、どちらか一方を単独で使用した場合よりも優れたパフォーマンスを発揮することがわかりました。
リランカーは小さいです。私たちが手作業で厳選した数百のトリプル (インテント、ツール、ラベル) でトレーニングされたクロスエンコーダー。これは、推論コストがシステムの中で最も安価な部分です。わずかなコストでモデルベースのスコアラーに置き換えることができます。
コスト管理が重要です。再強化は、変更されたサービスに対してのみ毎晩実行されます。 200 オペレーションのサービスでは通常、完全なエンリッチメント パスごとに GPT-4 クラスの推論に約 30 セントの費用がかかります。
管理 UI でヒット率メトリクスを公開します。どのツールが search_tools によって検出され、どのツールが無視されているかを確認できます。これは、どのサービスを最初に強化するかを判断するための最も有用なシグナルです。
ドキュメントは docs.mcp-bridge.ai にあります。エンリッチメントを有効にするには、管理 UI のサービスをワンクリックで切り替えることができます。リライトを公開する前にプレビューできます。
金曜日は Product Hunt を生放送します。削除された機能は応答の後処理であり、コンテキスト ウィンドウの問題の残りの半分です。また会いましょう。
API はすでに機能していますが、AI 用に構築されたものではありません。 MCPブリッジが作る

コードを変更することなく、数分で AI 対応になります。
MCP Bridge の背後にある運用上の基本要素と、認証、レート制限、再試行、可観測性を中心に形成されたエンジニアリング上の決定。
ハイブリッド フルテキスト + pgvector ツール検出に加え、getProcInfo3() を get_customer_account_details に変換する LLM 駆動のエンリッチメント。
レガシー アプリの最新化 検出とドキュメント クラウドと VMware の移行 エージェントティック ソフトウェア エンジニアリング チームの構築と展開 MCP ブリッジ レガシー アプリの最新化 検出とドキュメント クラウドと VMware の移行 エージェントティック ソフトウェア エンジニアリング チームのプラットフォームの構築と展開 会社概要 キャリア パートナー クラウド変革 ブログ デモを予約 プライバシー ポリシー EULA 最新の AppFactor ニュースレター更新を受け取ります。

## Original Extract

AppFactor
Platform Use Cases Legacy App Modernization Upgrade legacy systems without disrupting your business. Discovery & Documentation Turn complex systems into clear, actionable documentation. Cloud and VMware Migration Move to the cloud with minimal downtime and full control. Build & Deploy Ship faster with streamlined build and deployment workflows. Agentic Software Engineering Team Scale your team with AI-driven engineering. Company About us Careers Partners Cloud transformation Blog MCP Bridge Platform Blog MCP Bridge Company
MCP Bridge Part 3: How we made getProcInfo3() agent-readable: hybrid discovery + AI Enrichment
In the previous article, we walked through Code Mode, three meta-tools that replace the entire MCP tool catalog when the API surface is large. The first of those three meta-tools is search_tools. Today we're opening it up.
search_tools is what stands between an LLM agent and a 200-operation API surface. It needs to take a natural-language description of what the agent wants to do, and return the three or four tools that can actually do it. Get this wrong and the agent ends up either flailing through irrelevant tools or, worse, calling the wrong one confidently.
We thought this would be the easy part of MCP Bridge. It wasn't.
What we tried first, and why it didn't work
Our first cut was pure vector search. Embed every tool's name + description with an OpenAI embedding model, store them in pgvector, query at runtime. It works fine on a textbook dataset. It falls over on enterprise APIs the moment you have two tools with similar embeddings but different intent ( get_customer vs get_customer_full vs get_customer_with_orders ).
We tried full-text search next. Postgres FTS with synonym dictionaries. Better on exact matches, worse on the cases where the agent's intent doesn't share vocabulary with the tool description.
What landed is a hybrid: Postgres FTS does the first pass (with BM25-style ranking), pgvector does the second pass on the top FTS candidates plus a wider semantic-only pool, and a small reranker collapses the two ordered lists into a single score. The reranker also factors in the agent's recent context, if the agent has been working with customer data, customer-shaped tools rank higher.
This works well, when the source data is good.
Enterprise APIs aren't named for LLMs. They're named for the engineers who wrote them, often a decade ago, under naming conventions that have since been forgotten. A real example from a real customer:
getProcInfo3(custId, opt) → object
Description: "See documentation."
The signature alone is useless. custId is a faint hint, opt could mean anything, and the return type object tells you nothing. No semantic search method survives that input.
The response, on the other hand, is full of signal:
{
"custId": "C-44218",
"billingAddress": { "..." },
"accountStatus": "active",
"primaryContact": { "..." },
"assignedManager": "..."
}
That's what the tool actually does, a customer account lookup, billing-and-contact level detail. The shape of the response is what tells you what this is.
We considered telling customers to fix their API names. This is a real solution at the small scale and an impossible solution at the only scale that matters. The customer with 70 legacy services is not going to rename their APIs because we asked.
You point MCP Bridge at any OpenAI-compatible chat completions endpoint. We use Azure OpenAI internally; Anthropic via a compatible proxy works; local Llama via vLLM works. You enable enrichment on a service, and the platform:
For each tool, gathers whatever signal is available: the name, the (often empty) description, the parameter schema, the response type definition (pulled from OpenAPI, WSDL, or .proto when available), and, for opaque APIs, a captured sample response from a probe call or traced production traffic.
Sends a structured prompt asking the model to generate a clearer name, a 1–2 sentence description of what the tool does and when an agent would use it, and a list of 3–5 tags or aliases.
Validates the output against a schema (no hallucinated parameters, no semantic drift in input/output types).
Stores the enriched metadata alongside the original schema. Tool calls still use the original name and schema, only the discovery layer sees the enriched version.
The signal that matters most is the response shape. We tried enrichment with just name + description and the model had no way to tell getProcInfo3 from getProcInfo4. Once it can see the response (or a sample), the function becomes legible. For SOAP services, WSDLs typically give us the response type. For undocumented APIs, MCP Bridge can either run a one-time probe or pull from a trace.
Crucially: enrichment is opt-in per service and the original schema is preserved. The agent never sees a renamed parameter. We didn't want a system that quietly diverges from the source API.
Here's getProcInfo3 from the real (anonymized) customer SOAP service. The enrichment model saw the original signature, the empty description, and the response sample above. From the response fields ( custId, billingAddress, accountStatus, primaryContact, assignedManager ), it inferred:
‍ Description: "See documentation."
Name: get_customer_account_details
‍ Description: "Fetches a customer's account profile, including billing address, account status, primary contact, and assigned account manager. Useful for verifying account state before initiating support, billing, or sales actions."
Aliases: customer_lookup, account_info, customer_profile, billing_account
Tool-selection accuracy on the agent task suite we used yesterday's benchmark on improved from 41% to 89% on this customer's services after enrichment. The numbers below the SOAP suite were less dramatic because the REST services already had reasonable names, gains were 71% → 88% there. The lesson: enrichment matters most exactly where it's hardest to do anything else.
Embeddings are regenerated when the enriched description changes. The original embedding is also kept; we found that combining them at retrieval time (weighted, original × 0.3 + enriched × 0.7) outperformed using either alone.
The reranker is small. A cross-encoder trained on a few hundred (intent, tool, label) triples we hand-curated. It's the cheapest part of the system in inference cost; you could replace it with a model-based scorer at minor cost.
Cost control matters. Re-enrichment runs nightly only for services that have changed. A 200-operation service typically costs about 30 cents of GPT-4-class inference per full enrichment pass.
We expose hit-rate metrics in the admin UI. You can see which tools are being found by search_tools and which are being ignored. This is the single most useful signal for figuring out which services to enrich first.
Docs are at docs.mcp-bridge.ai. Enabling enrichment is a one-click toggle on any service in the admin UI; you can preview the rewrites before they go live.
On Friday, we're live on Product Hunt. The feature drop is Response Post-Processing, the other half of the context-window problem. See you there.
Your APIs already work: they weren't built for AI. MCP Bridge makes them AI-ready, with no code changes, in minutes.
The operational primitives behind MCP Bridge and the engineering decisions that shaped around auth, rate limits, retries, observability.
Hybrid full-text + pgvector tool discovery, plus LLM-driven enrichment that turns getProcInfo3() into get_customer_account_details.
Legacy App Modernization Discovery & Documentation Cloud and VMware Migration Build & Deploy Agentic Software Engineering Team MCP Bridge Legacy App Modernization Discovery and Documentation Cloud and VMware Migration Build & Deploy Agentic Software Engineering Team Platform About us Career Partners Cloud transformation Blog Book a demo Privacy Policy EULA Receive the latest AppFactor newsletter updates.
