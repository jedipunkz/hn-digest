---
source: "https://seldon-ai.com/blog/silent-epidemic-llm-tech-debt"
hn_url: "https://news.ycombinator.com/item?id=48872743"
title: "The Silent Epidemic of LLM Technical Debt"
article_title: "The Silent Epidemic of LLM Technical Debt | Seldon"
author: "nlpnerd"
captured_at: "2026-07-11T15:45:29Z"
capture_tool: "hn-digest"
hn_id: 48872743
score: 2
comments: 0
posted_at: "2026-07-11T15:12:38Z"
tags:
  - hacker-news
  - translated
---

# The Silent Epidemic of LLM Technical Debt

- HN: [48872743](https://news.ycombinator.com/item?id=48872743)
- Source: [seldon-ai.com](https://seldon-ai.com/blog/silent-epidemic-llm-tech-debt)
- Score: 2
- Comments: 0
- Posted: 2026-07-11T15:12:38Z

## Translation

タイトル: LLM 技術的負債の静かな流行
記事のタイトル: LLM 技術的負債の静かな流行 |セルダン
説明: プロンプト駆動型開発では、コア アプリケーション ロジックが確率的テキスト内に静かに埋め込まれます。ここでは、どの LLM 呼び出しが安定したワークフローに成熟し、決定論的なシステムになるかを判断する方法を示します。

記事本文:
seldon 仕組み Models Docs Blog v0 ログイン サインアップ ← ブログ July 8, 2026 10 min read Seldon LLM 技術的負債の静かな流行
スタートアップ企業は、AI 製品のプロトタイプを最速で作成する方法が、プロンプト内にコア アプリケーション ロジックを埋め込む最速の方法である可能性があることに気づき始めています。
MVP フェーズでは、通常、LLM 請求書は無害であるように見えます。
通話あたり数セント。ユーザー数は数百人。いくつかのプロンプトチェーン。おそらく、フロンティア モデルは、分類、抽出、書式設定、ルーティング、要約、およびまだ誰も形式化する時間がなかったいくつかのビジネス ロジックを処理します。
それは動作します。デモは印象的です。創設者は出荷します。
その後、製品は成長し始めます。
1 つのワークフローが 10 個のワークフローになります。 1 人のユーザーが、チームが当初 2 つと想定していたモデル呼び出しを 20 回トリガーします。サポート チケット、営業電話、文書、請求書、電子メール、トランスクリプト、ファイリング、CRM メモ、および内部記録はすべて、同じフロンティア モデルのボトルネックを通過し始めます。
請求額が上がります。レイテンシーが可視化されます。出力は非決定的なままです。一時的なものであるはずのプロンプトが実稼働インフラストラクチャになります。
このスタートアップは単に LLM を採用しただけではありません。バックエンド ロジックを確率的テキスト インターフェイス内で誤ってエンコードしてしまいました。
これは、多くの AI スタートアップ企業が現在参入しているタールピットです。
創設者と建設者が私たちに語ったもの
私たちは最近、スタートアップの創設者、経営者、技術ビルダーと、次のような単純な質問について非公式のディスカッションを行いました。
本番環境での LLM の使用量は、真にオープンエンド推論であり、反復的な抽出、分類、正規化、変換、またはワークフロー ルーティングはどの程度でしょうか?
反応は驚くほど一貫していました。
ある創設者は、分類の初期段階でフロンティア モデルを使用し、その後、それをはるかに小さな細かいファイン モデルに置き換えたと説明しました。

ローカルで実行されている uned モデル。彼らの報告によると、コスト削減は約 90% でした。最も困難な部分は移行そのものではありませんでした。 GPT ベースのバージョンがまだ「正常に動作している」ように見える間に、移行を優先するようチームを説得しました。
別のビルダーは、LLM をコールドスタート エンジンとして使用することについて説明しました。このモデルはメモをカテゴリーに分類し、十分なラベル付きデータを生成した後、より安価な分類器に置き換えました。そのワークフローでは、LLM は最終的なシステムではありませんでした。これは、最終システムのトレーニング セットの作成に役立つツールでした。
3 番目の演算子は、階層型アーキテクチャについて説明しました。予測可能な 80% には正規表現とルール、単純な大量タスクには安価なモデル、繰り返されるエンティティにはキャッシュ、そして実際に推論が必要なあいまいなケースにはフロンティア LLM のみを使用します。
10 年以上の経験を持つデータ プロフェッショナルは率直にこう言います。LLM はプロトタイプには優れていますが、規模が重要なデータ作業にはコストがかかり、時間がかかり、変動が大きいです。彼らのパターンは、LLM から始めて、安定した抽出タスクと変換タスクを Python とデータ パイプラインに移行することです。
別のコメント投稿者は、多くのチームが完全にスキップしているという質問を提起しました。ワークフロー タイプごとに実際に測定される実稼働 LLM コールの割合は何パーセントですか?インストルメンテーションがなければ、チームは、どの呼び出しがオープンエンド推論であり、どの呼び出しが決定論的抽出であり、どの呼び出しが単に高価なフォーマットであるかを判断できません。
それが中心的な問題かもしれません。
多くのチームは、AI システムがどこで終わり、データ パイプラインが始まればよいのかを知りません。
本当の問題はコストだけではない
トークンの支出はプロトタイプでは管理可能に見えますが、使用量に応じて増加します。数セントのコストがかかるワークフローでも、何十万回、何百万回も実行すると苦痛になる可能性があります。さらに悪いことに、モデルの使用量はユーザーよりも速く拡張されることがよくあります。

1 つのユーザー アクションが多数のモデル呼び出しに展開される可能性があるためです。
LLM がコール / ユーザーに費やす 時間 / 導入 → プロトタイプ / PoC 制作 図 1 — モデルはユーザーよりも早く化合物を費やす
しかし、コストは目に見える症状にすぎません。
さらに深刻な問題は、アプリケーションを推論することが難しくなることです。
従来のバックエンドには、関数、スキーマ、テスト、インターフェイス、ステート マシン、バリデータ、ログ、および型指定されたデータがあります。プロンプトチェーンにはこれらのものが何もないことがよくあります。ビジネス ルール、解析の前提、フォーマット指示、ルーティング ロジック、および検証基準が含まれる場合がありますが、そのロジックはすべて自然言語で存在します。
何かが壊れるとデバッグがおかしくなります。元データが間違っていたのでしょうか？取得に失敗しましたか?プロンプトが長すぎましたか?モデルによって動作が変わりましたか?スキーマがずれてしまったのでしょうか？新しいエッジケースは登場しましたか？ LLM は、推論すべきではない何かを推論しましたか?
あるコメント投稿者は、企業が非決定的な動作に基づいて構築された製品をどのようにサポートできるのかを尋ねました。データが破損していると顧客から苦情が来た場合はどうなりますか?ロジックはモデル内にあり、出力は異なる可能性があると思いますか?
その質問は生産上のギャップを露呈するものであるため、不快です。 LLM は曖昧さを吸収するのに優れています。実稼働システムは曖昧さを減らすことを目的としています。
LLM は新しいプロトタイプ層になりつつあります
LLM について考える最も有益な方法は、すべてのバックエンド ロジックを永続的に置き換えるものではありません。多くの場合、これらはプロトタイプ層として理解される方がよいでしょう。
LLM を使用すると、チームはワークフローを完全に理解する前にワークフローの形状を発見できます。チームがデータ モデルにコミットする前に、乱雑な入力を処理し、可能性のある構造を推測し、初期のラベルを生成し、スキーマを提案し、レコードを正規化し、有用な出力を生成できます。
それは貴重なことだ。多くの場合、それは、

正しい始め方。間違いは、プロトタイプ層が運用ランタイムのままであるべきだと想定していることです。
成功したワークフローの多くは、フロンティア モデルによって提供される探索フェーズから、モデルがフォールバックとしてのみ保持される決定論的でコンパイルされたコードによって提供される運用フェーズまで、同じ弧を描いているように見えます。
あいまいなタスク 即時プロトタイプ 繰り返し使用 安定した I/O 明示的なスキーマ コンパイル済み + フォールバック 図 2 — ワークフローは探索から本番まで成熟します
ワークフローがまだわかっていない最初の段階では、フロンティア モデルが役立ちます。ワークフローが安定すると、すべてのリクエストに対して同じフロンティア モデルを使用し続けると、一種の技術的負債になる可能性があります。
すべての AI スタートアップが尋ねるべき質問
本番環境での LLM 呼び出しごとに、次のことを尋ねます。
この呼び出しは推論を行っているのでしょうか、それともデータ エンジニアリングを行っているのでしょうか?
インテリジェントに見える呼び出しの多くは、実際にはよく知られたデータ タスクを実行しています。
この顧客をアカウントと照合します。
このレコードがどのバケットに属するかを決定します。
この文書に必要な条項が含まれているかどうかを検証します。
この高度にテンプレート化された呼び出しを固定構造に要約します。
このワークフローを次の状態にルーティングします。
これらのタスクでは、特に入力が乱雑な場合、探索中に LLM が必要になる場合があります。しかし、十分な例の後、それらの多くはより決定的な実装の候補になります。
実際的な問題は、LLM を削除すべきかどうかではありません。実際的な問題は、それらをスタック内のどこに置くべきかということです。成熟したシステムでは、LLM は、すべてのリクエストがその費用を支払うホット パスから、安価な決定論的なルートに十分な信頼がない場合にのみ到達する例外パスに移動する必要があります。
Frontier LLM 出力されるすべてのリクエストはフロンティア コストを支払う AFTER コンパイルされたパス Frontier LLM すべてのリクエストの信頼ゲート ≈ 10% Frontier LLM ≈ 90% Determinist

ic、キャッシュ、分類器、または小規模モデルの出力 図 3 — フロンティア モデルをホット パスから例外パスに移動する
目標は、LLM を排除することではありません。目標は、それらを高価な接着コードとして使用するのをやめる事です。
従来のシステムとなるべき LLM コールを特定する方法
最初のステップは計測です。 LLM の合計支出だけを見ないでください。ワークフローごとのユニットエコノミクスを見てみましょう。
有用なトレースには、プロンプト ファミリ、入力タイプ、モデル、トークン数、レイテンシー、コスト、出力スキーマ、ダウンストリーム アクション、ユーザー フィードバック、および応答が受け入れられたか、編集されたか、再試行されたか、または破棄されたかどうかが含まれている必要があります。
トレースが存在すると、エンドポイントではなくワークフローごとに呼び出しをグループ化します。たとえば、スタートアップは次のようなクラスターを検出する場合があります。
サポートチケットカテゴリの抽出
通話記録を CRM アップデートに変換する
収益報告を KPI テーブルに要約する
ユーザークエリを内部オブジェクトと照合する
次に、各クラスターが「コンパイル可能性」に関してスコア付けされる必要があります。ワークフローは、入力と出力が安定しており、出力スキーマが繰り返し行われ、タスクが大量に繰り返され、許容可能な回答を検証でき、ビジネス ロジックがルール、データ変換、小規模モデル、または既知のデータの取得として表現できる場合に、置き換えの良い候補となります。
タスクが本当に無制限である場合、出力が主観的な場合、コンテキストがリクエスト間で根本的に変化する場合、または価値が構造化された変換ではなく創造的な統合から得られる場合、ワークフローは候補としては適していません。
ワークフローをボリュームと曖昧さ別にプロットすると、優先順位が明確になります。ボリュームが多く曖昧さの少ないコーナーは、コンパイルが最初に効果を発揮する場所です。
最初にコンパイルする ルート / キャッシュ / 安価なモデル 低い ROI — とりあえず放置 フロンティア LLM を保持する フィールドを抽出する チケットを分類する エンティティを正規化する 通話を要約する フォローアップ メールの草案 GTM 戦略の草案 Open-e

nded Q&A まれな 1 回限りの解析 曖昧性が低い 曖昧性が高い → ↑ ボリューム 図 4 — 最初にコンパイルを呼び出すのはどれか
ここに実践的なルーブリックがあります。
1. 繰り返されるプロンプト ファミリを探します
同じプロンプト テンプレートが何百回または何千回も表示される場合、それが候補となります。例としては次のものが挙げられます。
「この文書から次のフィールドを抽出します…」
「このチケットを次のカテゴリのいずれかに分類してください…」
「このトランスクリプトをこの JSON スキーマに変換します…」
「このレコードがこのアカウントと一致するかどうかを確認してください…」
繰り返しは、タスクが安定したことを示す最初の信号です。
2. 構造化されたアウトプットを探す
LLM が主に JSON、テーブル、ラベル、スコア、タグ、または制約付き要約を生成している場合、フロンティア モデルは永遠に必要ない可能性があります。構造化された出力は、チームが答えがどのような形であるべきかをすでに知っていることを示しています。
形状がわかったら、次の問題は、パーサー、分類子、抽出モデル、SQL クエリ、または ETL パイプラインが同じ構造をより確実に生成できるかどうかです。
3. 差異の少ない決定を探す
一部の LLM 呼び出しでは、同じ種類の応答が何度も生成されます。たとえば:
出力空間が小さい場合、これは推論問題に見せかけた分類問題であることがよくあります。小規模な微調整モデル、埋め込み分類器、勾配ブースト モデル、またはルールさえもあれば十分な場合があります。
4. エンティティの解決が繰り返されているかどうかを確認します
エンティティ マッチングは、非表示 LLM の最も一般的な使用法の 1 つです。チームはモデルに、乱雑なテキストや内部システム全体で名前、ベンダー、顧客、会社、製品、アカウント、ティッカー、人物、またはドキュメントを照合するよう依頼します。
音量が小さい場合は便利です。大規模になると、正規 ID、エイリアス、信頼スコア、あいまい一致、埋め込み、人間によるレビュー キューを備えたエンティティ解決パイプラインになることが望まれることがよくあります。
5. 検証できる変換を探す

d
呼び出しの出力を確認できると、呼び出しを置き換えるのが簡単になります。たとえば:
JSON はスキーマと一致しますか?
抽出された日付は解析されますか?
会社は既知のアカウントにマッピングされていますか?
合計は項目の合計と一致しますか?
チケットカテゴリは許可された分類に属していますか?
検証ができれば交換もより安全になります。
6. 大量の曖昧さの少ない通話を探す
すべての高価な通話を最初に置き換える必要はありません。最適な候補は、大量で曖昧性の少ないものです。まれな戦略生成呼び出しを置き換えることは問題ではない場合があります。テキストを既知の JSON スキーマに変換する毎月 100 万回の呼び出しを置き換えることで、利益が大幅に改善される可能性があります。
有用な優先順位付けスコアは次のとおりです。
コンパイルの優先順位 =
ボリューム × コスト × 反復性 × スキーマの安定性 × 検証の容易さ
チームは最高スコアのコールから開始する必要があります。
7. 隠されたデータモデルを探す
LLM 呼び出しが繰り返される場合は、製品に実際のデータ モデルが欠落していることを示しています。ユーザーが繰り返し質問した場合:
どのアカウントが更新のリスクにさらされていますか?
どの企業がガイダンスを変更しましたか?
どのチケットがエスカレーションの可能性がありますか?
異常な条件が設定されている契約はどれですか?
フォローアップが必要な見込み客はどれですか?
答えはこれ以上のプロンプトではないかもしれません。答えは、存在すべきテーブルかもしれません。たとえば:
これは私

[切り捨てられた]

## Original Extract

Prompt-driven development quietly buries core application logic inside probabilistic text. Here is how to tell which LLM calls have matured into stable workflows that should become deterministic systems.

seldon How it works Models Docs Blog v0 Log in Sign up ← Blog July 8, 2026 10 min read Seldon The Silent Epidemic of LLM Technical Debt
Startups are discovering that the fastest way to prototype AI products may also be the fastest way to bury core application logic inside prompts.
During the MVP phase, the LLM bill usually looks harmless.
A few cents per call. A few hundred users. A few prompt chains. Maybe a frontier model handles classification, extraction, formatting, routing, summarization, and a few pieces of business logic that nobody has had time to formalize yet.
It works. The demo is impressive. The founder ships.
Then the product starts to grow.
One workflow becomes ten workflows. One user triggers twenty model calls where the team originally assumed two. Support tickets, sales calls, documents, invoices, emails, transcripts, filings, CRM notes, and internal records all start flowing through the same frontier-model bottleneck.
The bill rises. Latency becomes visible. Outputs remain non-deterministic. The prompt that was supposed to be temporary becomes production infrastructure.
The startup has not merely adopted LLMs. It has accidentally encoded its backend logic inside a probabilistic text interface.
This is the tarpit many AI startups are now entering.
What founders and builders told us
We recently ran an informal discussion with startup founders, operators, and technical builders about a simple question:
How much production LLM usage is truly open-ended reasoning, and how much is repetitive extraction, classification, normalization, transformation, or workflow routing?
The responses were strikingly consistent.
One founder described using a frontier model during the early phase for classification, then replacing it with a much smaller fine-tuned model running locally. Their reported cost reduction was around 90%. The hardest part was not the migration itself. It was convincing the team to prioritize the migration while the GPT-based version still appeared to be "working fine."
Another builder described using an LLM as a cold-start engine. The model classified notes into categories, generated enough labeled data, and was then replaced by a cheaper classifier. In that workflow, the LLM was not the final system. It was the tool that helped create the training set for the final system.
A third operator described a tiered architecture: regex and rules for the predictable 80%, cheaper models for simple high-volume tasks, caching for repeated entities, and frontier LLMs only for ambiguous cases that actually require reasoning.
A data professional with more than a decade of experience put it bluntly: LLMs are excellent for prototypes, but for scale-critical data work, they are costly, slow, and high variance. Their pattern is to start with LLMs, then migrate stable extraction and transformation tasks into Python and data pipelines.
Another commenter raised the question that many teams skip entirely: what percentage of production LLM calls are actually measured by workflow type? Without instrumentation, teams cannot tell which calls are open-ended reasoning, which are deterministic extraction, and which are simply expensive formatting.
That may be the central issue.
Many teams do not know where their AI system ends and their data pipeline should begin.
The real problem is not just cost
Token spend looks manageable in the prototype, then compounds with usage. A workflow that costs a few cents may become painful when it runs hundreds of thousands or millions of times. Worse, model usage often scales faster than user count, because one user action may fan out into many model calls.
LLM calls / spend Users Time / adoption → PROTOTYPE / PoC PRODUCTION Fig. 1 — Model spend compounds faster than users
But cost is only the visible symptom.
The deeper problem is that the application becomes difficult to reason about.
A conventional backend has functions, schemas, tests, interfaces, state machines, validators, logs, and typed data. A prompt chain often has none of those things. It may contain business rules, parsing assumptions, formatting instructions, routing logic, and validation criteria, but all of that logic lives in natural language.
When something breaks, debugging becomes strange. Was the source data wrong? Did retrieval fail? Did the prompt become too long? Did the model change behavior? Did the schema drift? Did a new edge case appear? Did the LLM infer something it should not have inferred?
One commenter asked how a company can support a product built on non-deterministic behavior. What happens when a customer complains that their data was mangled? Do you say the logic is inside a model and the output may vary?
That question is uncomfortable because it exposes the production gap. LLMs are wonderful at absorbing ambiguity. Production systems are supposed to reduce ambiguity.
LLMs are becoming the new prototype layer
The most useful way to think about LLMs is not as permanent replacements for all backend logic. They are often better understood as a prototype layer.
An LLM lets a team discover the shape of a workflow before the workflow is fully understood. It can handle messy inputs, infer likely structure, generate early labels, propose schemas, normalize records, and produce useful outputs before the team has committed to a data model.
That is valuable. In many cases, it is the right way to begin. The mistake is assuming the prototype layer should remain the production runtime.
Many successful workflows seem to follow the same arc — from an exploration phase served by the frontier model to a production phase served by deterministic, compiled code, with the model kept only as a fallback:
Ambiguous task Prompt prototype Repeated usage Stable I/O Explicit schema Compiled + fallback Fig. 2 — A workflow matures from exploration to production
The frontier model is useful in the beginning because the workflow is not yet known. Once the workflow stabilizes, continuing to use the same frontier model for every request may be a form of technical debt.
The question every AI startup should ask
For each LLM call in production, ask:
Is this call doing reasoning, or is it doing data engineering?
Many calls that appear intelligent are actually performing familiar data tasks:
Match this customer to an account.
Decide which bucket this record belongs to.
Validate whether this document has the required clauses.
Summarize this highly templated call into a fixed structure.
Route this workflow to the next state.
Those tasks may require an LLM during exploration, especially when the input is messy. But after enough examples, many of them become candidates for more deterministic implementation.
The practical question is not whether LLMs should be removed. The practical question is where they should sit in the stack. In mature systems, the LLM should move from the hot path — where every request pays for it — to the exception path, reached only when a cheap deterministic route is not confident enough:
Frontier LLM Output every request pays the frontier cost AFTER Compiled path Frontier LLM Every request confidence gate ≈10% Frontier LLM ≈90% Deterministic, cache, classifier or small model Output Fig. 3 — Move the frontier model from the hot path to the exception path
The goal is not to eliminate LLMs. The goal is to stop using them as expensive glue code.
How to identify LLM calls that should become traditional systems
The first step is instrumentation. Do not look only at total LLM spend. Look at unit economics by workflow.
A useful trace should include the prompt family, input type, model, token count, latency, cost, output schema, downstream action, user feedback, and whether the response was accepted, edited, retried, or discarded.
Once traces exist, group calls by workflow rather than by endpoint. For example, a startup may discover clusters such as:
Extract support ticket category
Convert call transcript into CRM update
Summarize earnings call into KPI table
Match user query to internal object
Each cluster should then be scored for "compilability." A workflow is a good candidate for replacement when the inputs and outputs are stable, the output schema is recurring, the task is repeated at high volume, the acceptable answer can be validated, and the business logic can be expressed as rules, data transformations, small models, or retrieval over known data.
A workflow is a poor candidate when the task is truly open-ended, the output is subjective, the context changes radically across requests, or the value comes from creative synthesis rather than structured transformation.
Plotting workflows by volume and ambiguity makes the priority obvious — the high-volume, low-ambiguity corner is where compiling pays off first:
COMPILE FIRST Route / cache / cheaper model Low ROI — leave for now Keep frontier LLM Extract fields Classify ticket Normalize entity Summarize call Draft follow-up email Draft GTM strategy Open-ended Q&A Rare one-off parse low ambiguity high ambiguity → ↑ volume Fig. 4 — Which calls to compile first
Here is a practical rubric.
1. Look for repeated prompt families
If the same prompt template appears hundreds or thousands of times, it is a candidate. Examples include:
"Extract the following fields from this document…"
"Classify this ticket into one of these categories…"
"Turn this transcript into this JSON schema…"
"Determine whether this record matches this account…"
Repetition is the first signal that the task has stabilized.
2. Look for structured outputs
If the LLM is mostly producing JSON, tables, labels, scores, tags, or constrained summaries, it may not need a frontier model forever. Structured output is a sign that the team already knows what shape the answer should have.
Once the shape is known, the next question is whether a parser, classifier, extraction model, SQL query, or ETL pipeline can produce the same structure more reliably.
3. Look for low-variance decisions
Some LLM calls produce the same kind of answer again and again. For example:
If the output space is small, this is often a classification problem disguised as a reasoning problem. A small fine-tuned model, embedding classifier, gradient-boosted model, or even rules may be sufficient.
4. Look for repeated entity resolution
Entity matching is one of the most common hidden LLM uses. Teams ask models to match names, vendors, customers, companies, products, accounts, tickers, people, or documents across messy text and internal systems.
At low volume, this is convenient. At scale, it often wants to become an entity-resolution pipeline with canonical IDs, aliases, confidence scores, fuzzy matching, embeddings, and human review queues.
5. Look for transformations that can be validated
A call is easier to replace when its output can be checked. For example:
Does the JSON match the schema?
Does the extracted date parse?
Does the company map to a known account?
Does the total equal the sum of the line items?
Does the ticket category belong to the allowed taxonomy?
If validation is possible, replacement becomes safer.
6. Look for high-volume, low-ambiguity calls
Not every expensive call should be replaced first. The best candidates are high-volume and low-ambiguity. Replacing a rare strategy-generation call may not matter; replacing a million monthly calls that convert text into a known JSON schema may materially improve margins.
A useful prioritization score is:
Compilation priority =
volume × cost × repetitiveness × schema stability × validation ease
The highest-scoring calls are where teams should start.
7. Look for hidden data models
Some repeated LLM calls indicate that the product is missing a real data model. If users repeatedly ask:
Which accounts are at renewal risk?
Which companies changed guidance?
Which tickets are likely escalations?
Which contracts have unusual terms?
Which prospects need follow-up?
the answer may not be a better prompt. The answer may be a table that should exist. For example:
This i

[truncated]
