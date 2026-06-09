---
source: "https://www.kaelio.com/blog/open-source-anthropic-internal-data-analytics-engine"
hn_url: "https://news.ycombinator.com/item?id=48463102"
title: "Show HN: Open-Source Version of Anthropic's Internal Analytics Engine"
article_title: "We Built the Open-Source Version of Anthropic's Internal Data Analytics Engine | Kaelio"
author: "lucamrtl"
captured_at: "2026-06-09T18:53:01Z"
capture_tool: "hn-digest"
hn_id: 48463102
score: 4
comments: 0
posted_at: "2026-06-09T16:17:36Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-Source Version of Anthropic's Internal Analytics Engine

- HN: [48463102](https://news.ycombinator.com/item?id=48463102)
- Source: [www.kaelio.com](https://www.kaelio.com/blog/open-source-anthropic-internal-data-analytics-engine)
- Score: 4
- Comments: 0
- Posted: 2026-06-09T16:17:36Z

## Translation

タイトル: Show HN: Anthropic の内部分析エンジンのオープンソース バージョン
記事のタイトル: Anthropic の内部データ分析エンジンのオープンソース バージョンを構築しました |カエリオ
説明: Anthropic は、データ チームがビジネス分析を 95% 以上の精度でほぼ 100% 自動化した方法を公開しました。彼らが説明する 4 層アーキテクチャは、データ エージェント用の独立した Apache-2.0 コンテキスト エンジンである ktx としてすでにオープンソース化されているものと同じです。
HN テキスト: HN さん、試してみたい場合は、ここにリポジトリがあります: https://github.com/Kaelio/ktx フィードバックをお待ちしています。

記事本文:
Anthropic の内部データ分析エンジンのオープンソース バージョンを構築しました | Kaelio 新しい ktx はオープンソースになりました。 GitHub で表示 → 製品 ktx データ エージェントのオープンソース コンテキスト レイヤー 機能 仕組み 統合 FAQ ktx クラウド データ チーム向けのホスト型コンテキスト インフラストラクチャ 機能 ktx エンジン統合 FAQ データ エージェント コンテキスト レイヤーに基づいたマネージド エージェント 機能 仕組み ユースケース FAQ 価格サービス ブログ ドキュメント はじめに デモの予約 製品 ktx の機能 仕組み 統合 FAQ ktx クラウド機能 ktx エンジン統合 FAQ データ エージェントの機能 仕組み ユースケース FAQ価格設定 サービス ブログ GitHub に関するドキュメント 始める デモを予約する ktx を試す
データエージェントのコンテキスト層。
ブログに戻る コンテキスト レイヤーの詳細 Luca Martial Kaelio CEO
Anthropic の内部データ分析エンジンのオープンソース バージョンを構築しました
Anthropic は、データ チームがビジネス分析を 95% 以上の精度でほぼ 100% 自動化した方法を公開しました。彼らが説明する 4 層アーキテクチャは、データ エージェント用の独立した Apache-2.0 コンテキスト エンジンである ktx としてすでにオープンソース化されているものと同じです。
セルフサービス分析が失敗する理由
Anthropic のエージェント分析スタック、レイヤーごと
各 ktx コンセプトが Anthropic スタックにどのようにマッピングされるか
ツールを動かすスキル
ktx 自体を最新の状態に保つ方法: 取り込みはプル リクエストです
あなたの会社でこれを設定する方法
セルフサービス分析が失敗する理由
Anthropic のエージェント分析スタック、レイヤーごと
各 ktx コンセプトが Anthropic スタックにどのようにマッピングされるか
ツールを動かすスキル
ktx 自体を最新の状態に保つ方法: 取り込みはプル リクエストです
あなたの会社でこれを設定する方法
Anthropic は、セルフサービス分析のためのデータ チームのプレイブックを公開しました: How Anthropic Enables Self-Service Data Analytics with Claude 。
の

その結果は印象的です。同社のエージェントは現在、ビジネス分析リクエストを 100% 近く自動化し、95% 以上の精度で自動化しています。
記事は宝の山です。また、これは「どうやったか」ではなく「何をしたか」であるため、自社で再構築するということは、困難な部分のほとんどをリバースエンジニアリングすることを意味します。
ktx は、Anthropic のデータ チームが説明したのと同じ 4 つのレイヤーをパッケージ化したオープンソースのコンテキスト エンジンであり、Anthropic の記事が公開される前に出荷しました。私たちは独立して同じアーキテクチャに到達しました。これにより、これらの 4 つの層が適切であることが確認されました。 ktx は Apache-2.0 で、1 つのコマンド ( github.com/Kaelio/ktx ) でインストールして実行できます。
この投稿では、Anthropic のエージェント分析システムの各コンポーネントについて説明し、ktx が同じ 4 つのレイヤーを、人間が関与して構築および保守する 1 つのインストール可能なエンジンとしてどのようにパッケージ化するかを説明します。
免責事項: 私たちは Anthropic とは提携しておらず、Anthropic のコードを見たことがありません。
セルフサービス分析が失敗する理由
これについて非常に簡単に説明します。何百回も聞いたことがあるかもしれませんが、Anthropic は、これまでに見た中で最も明確な方法の 1 つである、データはソフトウェアではないというセルフサービス分析の背後にある問題を明確に説明しています。エージェントがコードを記述するとき、ソリューション空間は無制限であり、テストは幻覚をキャッチします。分析はその逆です。 「単一の正しいソースを使用した場合、正しい答えは 1 つだけであることがよくあり」、その数値が間違っている場合、コンパイラはそれを教えてくれません。
彼らは、エージェントが間違った答えを導き出す 3 つの理由を挙げています。
概念から実体までのあいまいさ: 「エージェントは、ユーザーの質問に最もよく答える正しいフィールドを選択できません。」彼らが挙げている例は、「アクティブ ユーザーの数を測定する際に、どのような行動が「アクティブ」であるとみなされるのか? 詐欺的なユーザーは含まれていますか?

サーズ？どのようなルックバック ウィンドウを使用しますか?」
古さ: 「データ ソース、ビジネス定義、スキーマは常に変化します。資産とエージェントの知識は古くなります。」
取得の失敗: 「正しい情報が実際にはデータ モデル内に存在し、適切に注釈が付けられている可能性がありますが、エージェントが単にそれを見つけられないだけです。」
以下のすべてのコンポーネントは、これら 3 つのうちの 1 つを無効にするために存在します。
Anthropic のエージェント分析スタック、レイヤーごと
Anthropic のスタックは倉庫から上に 4 層で構成されています。これは同じスタックで、各層とそれが強制終了する障害モードを実装する ktx 部分の注釈が付けられています。
Anthropic の定義: 「データ ウェアハウス内のデータ モデル、変換、テスト、テーブルと、それらを記述するメタデータ」。彼らは、「列とテーブルの説明、正規のメトリクス定義、グレイン ドキュメント、有効な値の範囲、系統、所有権、モデルの階層化などを、変換そのものと同じ厳密さで維持すること」を望んでいます。
余談ですが、彼らのハードルは非常に高いです。私たちの経験では、データ チームの大多数はこのレベルの成熟度には程遠いです。そのため、前提条件として成熟した基盤を要求するのではなく、ウェアハウスをスキャンし、テーブルをサンプルし、関係を推測し、説明と定義を下書きし、人間が修正して所有できるようにそれらをレビュー可能なファイルに残すように ktx を設計しました。人間がループに留まる理由については次のレイヤーで詳しく説明します。Anthropic は完全自動化について鋭い警告を発しているからです。
データ エージェント用のオープンソース コンテキスト レイヤー。
ktx は、ウェアハウス スキーマ、メトリクス、および内部ドキュメントを、エージェントが正確な SQL をクエリする管理されたコンテキストに変換します。無料でオープンソース。
これらは、エージェントがナビゲートするのに役立つ参照サーフェスです: セマンティック レイヤー (管理されたメトリック定義)、系統、実際の分析方法のクエリ コーパス

ysts は、このデータとビジネス コンテキスト レイヤーをクエリしました。
ビジネスコンテキストは非常に重要です。彼らはそれを、インデックス付きドキュメント、ロードマップ、意思決定ログ、およびそれらの組織構造のナレッジ グラフとしてモデル化します。
LLM にセマンティック レイヤーを所有させないでください。彼らは「生のテーブルとクエリログからLLMにメトリクス定義を自動生成させることでセマンティックレイヤーをブートストラップする」ことを試みたが、「人間が厳選したより小さなレイヤーと比較して、私たちの評価ではネットネガティブでした」。彼らの経験則は、「ドキュメントはクロードと一緒に生成するが、定義は人間に任せる」です。
単に生の履歴をエージェントにダンプしないでください。彼らは、「エージェントに何千もの以前のクエリへの生の検索アクセスを与えることで、精度が 1 ポイント未満向上した」こと、およびその成果は、そのコーパスを精選された参考ドキュメントに抽出することによってもたらされることを発見しました。
ktx はまさにこの線を引いています。実行可能ファイルの定義は人間が所有する YAML に存在し、SQL は決定論的コンパイラー (モデルではない) によって生成され、散文はエージェントが読み取る検索可能な Wiki に存在しますが、コンパイラーは決して実行しません。レイヤー全体が git にコミットされ、コードのようにレビューされます。これは、Anthropic の「変換自体と同じ厳密さで維持される」に対する ktx の解釈です。
「エージェントがオンデマンドで読み取るマークダウンのフォルダー。」 Anthropic は、このレイヤーの精度が 21% から 95% 以上に大幅に向上したと考えています。これは、オンデマンドでドメインの詳細を読み込む「シン トップレベル ルーターとして機能する」ナレッジ スキルと、「質問を明確にし、ソースを見つけ、クエリを実行し、結果を敵対的レビュー サブエージェントを通じてループする」という「上級アナリストが従うであろうプロセスをコード化する」アンブック スキルの 2 つに分かれています。
ktx はこのレイヤーを完全に出荷します。これには、同じアナリスト プロセスをエンコードする実際のマークダウン スキル ( ktx-analytics ) が含まれており、そのスキルをスキー板の表面に表示する型指定されたツールと組み合わせます。

私はドライブします。詳細については、以下のマッピングを参照してください。
オフライン評価: 最も一般的なビジネス上の質問、ビジネスのコンテキストに関するロングテールの質問、およびその場で収集された人間による修正をカバーする Q&A ペア。
アブレーション: 体系的に部品を削除して、ボトルネックがどこにあるのか、どの変更が精度を最も高めるかを確認します。
オンライン検証: Anthropic は、回答時にライブ レイヤーを実行し、敵対的レビューのサブエージェントを使用して結果に異議を唱え、出所を明らかにし、データの品質と鮮度をチェックします。彼らは、敵対的レビューを測定し、32% 多くのトークンと 72% 高い遅延を犠牲にして 6% の精度を追加しました。
これは、ktx が今日最も薄い層です。 ktx は修正を収集し、定義がコンパイルされていることを検証しますが、回答の正しさの評価やオンラインの敵対的レビューはまだ実行していません。詳細については以下をご覧ください。
データ エージェント用のオープンソース コンテキスト レイヤー。
ktx は、ウェアハウス スキーマ、メトリクス、および内部ドキュメントを、エージェントが正確な SQL をクエリする管理されたコンテキストに変換します。無料でオープンソース。
各 ktx コンセプトが Anthropic スタックにどのようにマッピングされるか
これは、マッピング全体を 1 つのビューにまとめたものです。 Anthropic のスタック内のすべてのコンセプトは、具体的な ktx 表面に配置されます。
2 つの違いを指摘する価値があります。1 つ目は、私たちが期待していたものとは逆です。
スキルに違いはありません。私たちは、Anthropic と同じ種類のプロセス スキルを提供します。本当の違いは検索にあります。Anthropic のナレッジ スキルでは、エージェントがオンデマンドでマークダウンのフォルダーを読み取らせますが、ktx では、ランク付けされたハイブリッド検索と値辞書を通じて、厳選された Wiki とセマンティック レイヤーが提供されます。同じ仕事で、より適切に装備されたメカニズム。
検証こそが本当のギャップです。これは、Anthropic が最後にして最も困難な利益を得た場所であり、ktx がまだ構築している場所です。
ツールを動かすスキル
Anthropic のスタックには 2 つのスキルがあります。

ドメインごとの適切な参照ドキュメントにルーティングする薄い知識スキルと、分析プロセスをエンコードする予約解除スキル。これらのジョブは両方とも ktx に存在し、分割方法が異なるだけです。
ktx-analytics スキルは、MCP 接続のすぐ隣のコーディング エージェント (Claude Code、Claude Desktop、Codex、Cursor、または OpenCode) に直接インストールされます。これは予約解除スキルです。上級アナリストのループをエンコードし、各ステップでどのツールを呼び出すかをエージェントに指示します。
1. Discover は最初に Discover_data を呼び出し、参照のみを返します。
2. 上位ヒットの読み取りを検査します (wiki_read、sl_read_source、entity_details)
3.dictionary_search を解決して、「enterprise」などの値を列にマップします。
4. 粒度、メトリクス、ディメンション、フィルター、時間枠を特定する計画を立てる
5. クエリは sl_query を優先し、必要な場合にのみ sql_execution にドロップします
6. サニティチェックの合計、フィルタ、ヌル、タイムゾーン、次に状態ソースを検証します。
7.memory_ingest の永続的な学習をコンテキストに取り込みます
ナレッジ スキルのルーター ジョブは、実行時に Discover_data によって処理され、適切なコンテキストを見つけて、何を開くかをエージェントに返します。したがって、Anthropic の 2 つのスキルは、1 つの ktx スキルと検索ツールに集約されます。スキルはプロセスであり、入力されたツールは実行面であり、スキルはエージェントにどれを呼び出すかを指示します。
読み取り専用は約束されるだけでなく強制されます。すべてのクエリは、読み取りでない限りチェックされて拒否されます。また、書き込みを行う 1 つのツールは、ウェアハウスではなくコンテキスト リポジトリにアクセスします。
マークダウンのフォルダーとコンテキスト エンジンの最も明確な違いは、エージェントが物事を見つける方法です。 Anthropic の知識スキルにより、エージェントはオンデマンドでリファレンス ドキュメントを読むことができます。 ktx は、取り込み時にキュレーションされたレイヤーにインデックスを付け、実際の検索エンジンを通じて提供します。
Discover_data は 4 つの検索レーンを並行して実行し、それらに重み付けを行い、

d は結果を逆順位融合と融合します。
ディクショナリ、サンプリングされたウェアハウス値
token 、短いクエリのフォールバック
すべて (Wiki ページ、メトリクス、ディメンション、テーブル、列) を一度に検索し、1 つのランク付けされたリストを返します。
精度を高めるために最も役立つのは、値辞書です。 「エンタープライズ アカウントの数」を尋ねると、エージェントはどの列に「enterprise」が入っているかを推測する必要がありません。dictionary_search は、ウェアハウスをスキャンせずに値を正確な列に解決します。これはマークダウンのフォルダーではできないことであり、Anthropic の概念と実体間のあいまいさを正面から攻撃します。
これは、Anthropic がたどり着いた教訓と同じであり、取得時に適用されます。つまり、レイヤーをキュレートしてから、それを見つけられるようにします。 ktx は、生の歴史が答えではないことに同意します。 Wiki とセマンティック レイヤーは、リファレンス ドキュメントを抽出したものです。 ktx は、ファイルの読み取りではなく、ランク付けされたインデックスを通じてそれらを提供するだけです。
ktx が構築するコンテキストは、設計上、2 つのコミットされたサーフェスです。 1 つは実行可能で、もう 1 つは散文です。
セマンティック ソース ( semantic-layer/warehouse/orders.yaml ) には、自由形式 SQL の代わりに、承認されたメトリクス、グレイン、および結合グラフが保持されます。
名前 : 注文
テーブル : public.orders
説明:
user : 店頭を通じて行われた注文。

[切り捨てられた]

## Original Extract

Anthropic published how its data team automated close to 100% of business analytics at 95%+ accuracy. The four-layer architecture they describe is the same one we had already open-sourced as ktx, an independent, Apache-2.0 context engine for data agents.

Hey HN, here's the repo if you'd like to try it out: https://github.com/Kaelio/ktx Would love your feedback.

We Built the Open-Source Version of Anthropic's Internal Data Analytics Engine | Kaelio New ktx is now open source. View on GitHub → Products ktx The open-source context layer for data agents Features How it works Integrations FAQ ktx Cloud Hosted context infrastructure for data teams Features ktx Engine Integrations FAQ Data Agent A managed agent grounded in your context layer Features How it works Use cases FAQ Pricing Services Blog Docs About Get Started Book a Demo Products ktx Features How it works Integrations FAQ ktx Cloud Features ktx Engine Integrations FAQ Data Agent Features How it works Use cases FAQ Pricing Services Blog Docs About GitHub Get Started Book a Demo Try ktx
The context layer for data agents.
Back to blog More in Context layer Luca Martial CEO at Kaelio
We Built the Open-Source Version of Anthropic's Internal Data Analytics Engine
Anthropic published how its data team automated close to 100% of business analytics at 95%+ accuracy. The four-layer architecture they describe is the same one we had already open-sourced as ktx, an independent, Apache-2.0 context engine for data agents.
Why self-serve analytics fails
Anthropic's agentic analytics stack, layer by layer
How each ktx concept maps to the Anthropic stack
The skill that drives the tools
How ktx keeps itself current: the ingest is a pull request
How to set this up at your company
Why self-serve analytics fails
Anthropic's agentic analytics stack, layer by layer
How each ktx concept maps to the Anthropic stack
The skill that drives the tools
How ktx keeps itself current: the ingest is a pull request
How to set this up at your company
Anthropic just published their data team's playbook for self-serve analytics: How Anthropic enables self-service data analytics with Claude .
The results are impressive: their agent now automates close to 100% of business analytics requests, and does so at 95%+ accuracy.
The article is a gold mine. It's also a "what-we-did" and not a "how-you-do-it", so rebuilding it at your own company means reverse-engineering most of the hard parts.
We already built it, so you don't have to bother with that: ktx is an open-source context engine that packages the same four layers described by Anthropic's data team, and we shipped it before Anthropic's article went up. We arrived at the same architecture independently, which confirms these are the right four layers. ktx is Apache-2.0, and you can install and run it with one command: github.com/Kaelio/ktx .
This post walks through each component of Anthropic's agentic analytics systems and explains how ktx packages the same four layers as one installable engine that builds and maintains itself, with a human in the loop.
Disclaimer: we are not affiliated with Anthropic and we never saw their code.
Why self-serve analytics fails
Very briefly on this. Even though you might have heard about it hundreds of times, Anthropic articulates the problem behind self-serve analytics in one of the clearest ways we've seen: data is not software. When an agent writes code, the solution space is open-ended and tests catch hallucinations. Analytics is the opposite. "There's often only a single correct answer using a single correct source," and no compiler will tell you when the number is wrong.
They name three reasons agents produce those wrong answers:
Concept-to-entity ambiguity: "the agent is unable to choose the correct fields that best answer a user's question." The example they cite: "In measuring the number of active users: what actions constitute being 'active'? Do you include fraudulent users? What lookback window do you use?"
Staleness: "data sources, business definitions, and schemas change constantly; assets and agent knowledge go stale."
Retrieval failure: "the right information may actually be in the data model and properly annotated, but the agent simply doesn't find it."
Every component below exists to kill one of these three.
Anthropic's agentic analytics stack, layer by layer
Anthropic's stack is four layers, built from the warehouse up. Here is the same stack, annotated with the ktx piece that implements each layer and the failure mode it kills.
Anthropic's definition: "the data models, transforms, tests, and tables in a data warehouse, along with the metadata describing them." They want "column and table descriptions, canonical metric definitions, grain documentation, valid value ranges, lineage, ownership, and model tiering ... maintained with the same rigor as the transformations themselves."
As a side note: their bar is very high. In our experience, the vast majority of data teams are nowhere near this level of maturity. That's why, instead of requiring mature foundations as a precondition, we designed ktx to scan your warehouse, sample tables, infer relationships, and draft the descriptions and definitions for you, and leave them in reviewable files so a human can correct and own them. More on why the human stays in the loop in the next layer, because Anthropic has a sharp warning about fully automating it.
The open-source context layer for data agents.
ktx turns your warehouse schema, metrics, and internal docs into governed context your agents query for accurate SQL. Free and open source.
These are the reference surfaces that help the agent navigate: a semantic layer (managed metric definitions), lineage, a query corpus of how real analysts have queried this data, and a business context layer.
Business context is critical. They model it as a knowledge graph of indexed docs, roadmaps, decision logs, and their organizational structure.
Do not let an LLM own the semantic layer. They tried "bootstrapping the semantic layer by having an LLM auto-generate metric definitions from raw tables and query logs," and it "was net-negative on our evals versus a smaller, human-curated layer." Their rule of thumb: "generate the documentation with Claude, but have a human own the definition."
Do not just dump raw history at the agent. They found that "giving the agent raw retrieval access to thousands of prior queries moved accuracy by less than a point," and that the win comes from distilling that corpus into curated reference docs.
ktx draws exactly this line. The executable definitions live in human-owned YAML, the SQL is produced by a deterministic compiler (not the model), and the prose lives in a searchable wiki the agent reads but the compiler never executes. The whole layer is committed to git and reviewed like code, which is ktx 's take on Anthropic's "maintained with the same rigor as the transformations themselves."
"A folder of markdown the agent reads on demand." Anthropic credits this layer with the entire jump from 21% to 95%+ accuracy. It splits in two: a knowledge skill that "acts as a thin top-level router" loading domain detail on demand, and an unbook skill that "encodes the process a senior analyst would follow: clarify the question, find sources, run the query, and then loop the result through adversarial review sub-agents."
ktx ships this layer in full. It includes a real markdown skill ( ktx-analytics ) that encodes the same analyst process, and pairs that skill with a typed tool surface the skill drives. More on that in the mapping below.
Offline evals: Q&A pairs covering the most common business questions, plus long-tail questions about business context, plus human corrections collected on the fly.
Ablation: systematically removing pieces to see where the bottlenecks are and which changes move accuracy the most.
Online validation: Anthropic runs a live layer at answer time that challenges the result with adversarial review sub-agents, surfaces provenance, and checks data quality and freshness. They measured adversarial review adding 6% accuracy at the cost of 32% more tokens and 72% higher latency.
This is the layer ktx is thinnest on today. ktx harvests corrections and validates that definitions compile, but it does not yet run answer-correctness evals or online adversarial review. More on that below.
The open-source context layer for data agents.
ktx turns your warehouse schema, metrics, and internal docs into governed context your agents query for accurate SQL. Free and open source.
How each ktx concept maps to the Anthropic stack
Here is the whole mapping in one view. Every concept in Anthropic's stack lands on a concrete ktx surface.
Two differences are worth calling out, and the first is the opposite of what we expected.
Skills are not a difference. We ship the same kind of process skill Anthropic does. The genuine difference is in retrieval: Anthropic's knowledge skill has the agent read a folder of markdown on demand, while ktx serves a curated wiki and semantic layer through a ranked hybrid search and a value dictionary. Same job, better-instrumented mechanism.
Validation is the real gap. This is where Anthropic got their last and hardest gains, and it is where ktx is still building.
The skill that drives the tools
Anthropic's stack has two skills. A thin knowledge skill that routes to the right per-domain reference doc, and an unbook skill that encodes the analyst process. Both of those jobs live in ktx , they are just split differently.
The ktx-analytics skill is installed straight into your coding agent (Claude Code, Claude Desktop, Codex, Cursor, or OpenCode) right next to the MCP connection. It is the unbook skill: it encodes the senior-analyst loop, and it tells the agent which tool to call at each step.
1. Discover call discover_data first, which returns refs only
2. Inspect read the top hits (wiki_read, sl_read_source, entity_details)
3. Resolve dictionary_search to map a value like "enterprise" to a column
4. Plan identify grain, metrics, dimensions, filters, time window
5. Query prefer sl_query, drop to sql_execution only when needed
6. Validate sanity-check totals, filters, nulls, time zones, then state sources
7. Capture memory_ingest durable learnings back into the context
The knowledge skill's router job is handled at runtime by discover_data , which finds the right context and hands the agent back what to open. So Anthropic's two skills collapse into one ktx skill plus a search tool: the skill is the process, the typed tools are the execution surface, and the skill tells the agent which to call.
Read-only is enforced, not just promised. Every query is checked and rejected unless it is a read, and the one tool that writes touches your context repo, never your warehouse.
The cleanest difference between a folder of markdown and a context engine is how the agent finds things. Anthropic's knowledge skill has the agent read reference docs on demand. ktx indexes the curated layer at ingest time and serves it through a real retrieval engine.
discover_data runs four search lanes in parallel, weights them, and fuses the results with Reciprocal Rank Fusion:
dictionary , sampled warehouse values
token , a short-query fallback
It searches everything at once (wiki pages, metrics, dimensions, tables, and columns) and returns one ranked list.
The most useful piece for accuracy is the value dictionary . Ask "how many enterprise accounts" and the agent does not have to guess which column holds "enterprise": dictionary_search resolves the value to the exact column without scanning the warehouse. That is something a folder of markdown cannot do, and it attacks Anthropic's concept-to-entity ambiguity head on.
This is the same lesson Anthropic landed on, applied at retrieval time: curate the layer, then make it findable. ktx agrees that raw history is not the answer. The wiki and semantic layer are the distilled reference docs. ktx just serves them through a ranked index instead of file reads.
The context ktx builds is two committed surfaces, by design. One is executable, one is prose.
A semantic source ( semantic-layer/warehouse/orders.yaml ) holds approved metrics, grain, and the join graph, instead of free-form SQL:
name : orders
table : public.orders
descriptions :
user : Orders placed through the storefront.

[truncated]
