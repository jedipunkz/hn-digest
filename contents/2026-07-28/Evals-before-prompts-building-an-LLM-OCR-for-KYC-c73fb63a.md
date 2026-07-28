---
source: "https://blog.nilenso.com/blog/2026/05/18/evals-before-prompts-building-an-llm-ocr-for-kyc/"
hn_url: "https://news.ycombinator.com/item?id=49080308"
title: "Evals before prompts: building an LLM OCR for KYC"
article_title: "Evals before prompts: building an LLM OCR for KYC - nilenso blog"
author: "priyangapkini"
captured_at: "2026-07-28T07:32:50Z"
capture_tool: "hn-digest"
hn_id: 49080308
score: 1
comments: 0
posted_at: "2026-07-28T06:58:01Z"
tags:
  - hacker-news
  - translated
---

# Evals before prompts: building an LLM OCR for KYC

- HN: [49080308](https://news.ycombinator.com/item?id=49080308)
- Source: [blog.nilenso.com](https://blog.nilenso.com/blog/2026/05/18/evals-before-prompts-building-an-llm-ocr-for-kyc/)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T06:58:01Z

## Translation

タイトル: プロンプト前の評価: KYC 用の LLM OCR の構築
記事のタイトル: プロンプト前の評価: KYC 用の LLM OCR の構築 - nilenso ブログ
説明: 実稼働グレードのドキュメント抽出のための評価フレームワークを構築します。

記事本文:
プロンプト前の評価: KYC 用の LLM OCR の構築
KYC (Know Your Customer) は、銀行、フィンテック、そしてますます多くの企業が、個人が本人であることを確認する方法です。その中心には、スキャンした PAN カードまたは Aadhaar のフィールドをフォームに入力するという、非常に地味なタスクがあります。人間はこれが苦手です。それは退屈でエラーが発生しやすく、大規模になると費用がかかります。代わりにLLMがそれを行うことができますか?
去年の今頃、私はクライアントの一人に対してその質問に答えようとしていた。彼らはこの抽出を行うためにサードパーティのサービスに料金を支払っていましたが、請求額が高額だったので、社内で作業を行いたいと考えていました。問題は、LLM が KYC に要求される精度を犠牲にすることなく、文書検証のコストを削減できるかどうかでした。
プロンプトを作成する前に、まず問題を理解する必要がありました。実稼働環境でどのドキュメントが表示されるのか、各ドキュメントからどのフィールドを抽出するのか、何が正しいとみなされるのか。それは、何か対策するものを構築することを意味しました。
クライアントの既存のデータを確認した後、ドメインの専門家によってすでに正常に検証および精査されているドキュメントを選択しました。これは、システムを評価するためのベンチマークとして機能する、入力と「グラウンド トゥルース」(理想的な) 出力の高度に厳選され、人間によって検証されたコレクションであるゴールデン データセットになりました。小さくしましたが、実際の生産分布を代表するものにしました。
ここから始めることが重要でした。それは、クライアントに対して弁明できる 1 つの数字を与えてくれたからです。「実際に送信された文書の種類に応じて、システムは X% の確率で正しいです。」モデルの比較、プロンプト調整、回帰チェックなどのより複雑な機能は、ベースラインの設定の下流にありました。
書類の種類は 4 つあります: PAN (インドの所得税 ID)、Aadhaar (インドの国民 ID カード)、DL (インド)

運転免許証）、および RC（インド車両登録証明書）。それぞれに独自のレイアウトと故障モードがあり、モデルが学習した内容を別のモデルに転送することはできません。 4 つすべての単一の混合精度数値があれば、本当の問題は解決されます。ある文書タイプを完全に読み取る LLM が、別の文書タイプでは失敗する可能性があるということです。そこで、ゴールデン データセットをドキュメント タイプごとに 1 つずつ、計 4 つのスライスに分割し、それぞれを独自のベースラインと反復の余地を持つ別個の問題として扱いました。
入力には、フラットベッド スキャンと携帯電話の写真、グレアとスキュー、州ごとの DL および RC レイアウトのロングテール、および英語と併記された地域言語のテキストを含むドキュメントを組み合わせて含めました。これにより、現実世界のデータを反映しながら、各評価を扱いやすくすることができました。
ゴールデン データセットはどのくらいの大きさにすべきでしょうか?本番環境を代表するのに十分な大きさと、専門家がすべての実行をエンドツーエンドでレビューできるほど十分に小さい必要がありました。私はドキュメント タイプごとに 50 ドキュメントに落ち着きました。各スライス内の現実的なミックスをカバーするにはこれで十分でした。イテレーションは速かったです。完全な実行を数時間でレビューでき、プロンプトを変更したりモデルを交換したりするたびに再度実行することができました。
スキップしがちで、スキップすべきではないものの 1 つは、ルーブリックです。ルーブリックはデータセットの半分です。 「LLM はこれを正しく読みましたか?」たとえ単一の PAN カードであっても、「はい/いいえ」の質問ではありません。私はドメインの専門家と話し合い、何が許容できるかを分野ごとに特定しました。
ID 番号 (PAN、Aadhaar、DL 番号、登録番号): 文字ごとに完全一致。 1 文字間違うと失敗、ピリオドになります。フォーマットが有効な幻覚は最悪の種類の失敗であり、正直な誤読とは別にフラグが立てられます。
名前: 私たちはアニル・クマール対アニル・クマールを容認しますか?アニル対アニル・クマール（姓は省略）。アニック vs アニル (キャラクターの混乱)。アニル K vs A

ニル・クマール（略称）。 Seshadri 対 Sesadri (音訳の違い)。
生年月日: 比較する前に正規形式に正規化されるため、01/04/1990 と 1990-04-01 は異なるものとして罰せられません。
幻覚と省略: 幻覚フィールドは常に null または [読み取り不能] よりも悪いスコアになります。ルーブリックにはそのように明確に書かれています。なぜなら、KYC では「わかりません」は回復可能ですが、「自信を持って間違った答えがあります」は回復できないからです。
ルーブリックを手にして、人間によるレビューに移りました。身分証明書の場合、訓練を受けた目による自動チェックは不可能です。スキャンが判読できない場合、LLM は「わかりません」とは言いません。彼らは、もっともらしい ID を幻覚します。つまり、正しい AAAAA9999A の形の 10 文字の文字列、チェックサムを通過する 12 桁の Aadhaar、父親の名前スロットに読み取られる名前です。各抽出をソース画像と並べて確認すると、いくつかの同じ形状で故障モードが繰り返され始めていることがわかりました。これらのパターンは、次に何を修正すべきかを教えてくれる唯一のものでした。
データセット、ルーブリック、レビュー担当者を配置して、評価ループを作成しました。スライス内の 50 個のドキュメントすべてに対してプロンプトを実行し、抽出を収集し、ルーブリックに対してそれぞれをスコア付けして、そのスライスの番号と各失敗が発生した理由のリストという 2 つの情報を取得して終了します。次の変更を促すのは、数字ではなくリストでした。その後、再度実行します。
私の最初のプロンプトは一行でした。 「この画像から次のフィールドを抽出します」に近いもの。私はこれを Gemini 2.0 Flash に対して実行していました。クライアントは Gemini ファミリを好み、Flash はその中で最も安価なビジョン モデルでした。私はまず、最もシンプルなレイアウトでフィールドが最も少ない PAN を選択しました。それがうまくいくようになり、その後、Aadhaar、DL、RC に移りました。そこから、それぞれの反復

プロンプトは 3 つの大まかな段階に分かれています。
フェーズ 1: 構造。最初の変更は、ドキュメントをより良く読むこととはまったく関係がなく、プロンプトがどのように構成されているかのみでした。各文書タイプに独自のプロンプトを与えました。レイアウト、重要なフィールド、障害モードがあまりにも異なっていたため、4 つすべてを処理するものを書こうとするのは、負け戦でした。指示と入力が絡まらないように、システムとユーザーのプロンプトを分離しました。モデルに決定させるのではなく、出力スキーマを明示的に特定しました。気温をゼロ近くまで下げました。そして、説明を短く直接的なものに書き直しました。あらゆる「お願い」と「親切」が出てきました。ほとんどが標準的なプロンプトエンジニアリングの動きです。彼らは実際に仕事をしてくれましたが、途中までしかできませんでした。
フェーズ 2: ドメイン コンテキスト。次のラウンドは、モデルが見ているものをモデルに教えることでした。ドキュメントに複数の番号が記載されている場合、裸の「DL 番号を抽出する」は役に立ちませんでした。役に立ったのは、フィールド自体のルールを詳しく説明することでした。 DL の場合、「インドの運転免許証番号は、KA-XX-YYYYNNNNNNNN のような州固有の形式に従っています」。各 ID タイプに同様のコンテキストを追加しました。 Aadhaar はちょうど 12 桁です。 PAN は AAAAA9999A の後に続きます。私は、クライアント自身の注釈から抽出したいくつかのショットの例でプロンプトを補足しました。モデルが間違っていることは少なくなり、間違っていたとしても、その間違いを見つけるのが容易になりました。
フェーズ 3: 分解。最も重要な変更は、作業の分割方法でした。 1 つのプロンプトでドキュメント タイプのすべてのフィールドを一度に返すのではなく、抽出をフィールドごとに 1 つずつ、複数のフォーカスされた呼び出しに分割しました。単一呼び出しバージョンでは、レイアウト、言語、形式、フィールド セマンティクスの調整が多すぎて、あるフィールドのエラーが他のフィールドに影響を及ぼしていました。フィールドごとの呼び出しが遅くなり、さらに多くなりました

高価ではありましたが、各呼び出しは十分に簡単だったので、モデルがつまずくことはなくなりました。
これらの反復の終了までに、システムはサードパーティ サービスのコストの 10% でゴールデン セットに対して 92% の精度を達成しました。
出荷時には、モデル自体が十分に改善されていたため、分解を削除してドキュメントごとに 1 つの呼び出しに戻すことができました。この回避策は、その有用性を静かに過ぎてしまいました。
1 年後に振り返ると、2 つのことが変わります。
プロンプトに思考の連鎖を追加します。最初にドキュメント内に表示される内容を記述してからそれを抽出するようにモデルに依頼すると、プロンプトの書き換えによって追跡していたエラーの多くが捕捉されたでしょう。
また、LLM を判定として使用して反復ループを高速化することも検討します。ルーブリックと数百のラベル付きの例を用意することで、裁判官は各抽出をグラウンド トゥルースおよび表面上の欠陥と比較することができ、すべての文書ではなく、フラグが付けられたものだけをレビューできるようになりました。注意: 裁判官も間違っている可能性があるため、その呼び出しには独自の評価が必要になります。これについては別の投稿で説明します。
私が変更しないのはエンジンです。つまり、ゴールデン データセット、ルーブリック、評価ループです。プロンプトが調整されただけです。指示を際限なく書き直したり、モデルを交換したり、温度を微調整したりすることもできたでしょう。代わりに、ベースラインを構築し、正しい意味を明確にし、変更のたびに同じ文書を同じルーブリックに基づいて実行しました。地味な作品です。しかし、実際の作業はここから始まります。
一時的なアンチパターン: 予想される失敗を例外として扱わないでください。
⌂ アーカイブ
» RLVR と GRPO を使用して小規模モデルをトレーニングして、より適切な OCaml を作成する
ソフトウェアを構築するシステムの構築
メガステネスを使用した eval へのアプローチ
RLVR と GRPO を使用して小規模モデルをトレーニングして、より適切な OCaml を作成する
プロンプト前の評価: LLM OCR の構築

KYC用
nilenso は、インドのバンガロールに拠点を置く従業員所有のソフトウェア協同組合です。
私たちはテスト駆動開発と継続的デリバリーを実践しており、Clojure と Ruby on Rails を使用するのが大好きです。
このブログは、コンサルティング会社、製品会社、そして一般的に好奇心旺盛な存在としての当社の成長を紹介するものです。
hello@nilenso.com までご連絡ください。

## Original Extract

Building an evaluation framework for production-grade document extraction.

Evals before prompts: building an LLM OCR for KYC
KYC, or Know Your Customer, is how banks, fintechs, and a growing number of businesses verify that a person is who they claim to be. At the heart of it sits a deeply unglamorous task: typing the fields off a scanned PAN card or Aadhaar into a form. Humans are bad at this. It’s boring, error-prone, and, at scale, expensive. Can an LLM do it instead?
Around this time last year, I was trying to answer that question for one of our clients. They were paying a third-party service to do this extraction for them, and the bill was steep enough that they wanted to bring the work in-house. The question was whether LLMs could reduce the cost of document verification without sacrificing the accuracy KYC demands.
Before writing any prompts, I needed to understand the problem first. Which documents we’d see in production, which fields to pull off each one, and what counted as correct. That meant building something to measure against.
After reviewing the client’s existing data, I selected the documents that had already been successfully verified and vetted by a domain expert. This became the Golden dataset , a highly curated, human-verified collection of inputs and “ ground truth ” (ideal) outputs that served as the benchmark for evaluating the system. I kept it small but representative of real production distribution.
Starting here mattered because it gave me one number I could defend to the client: “on the kind of documents you actually send us, the system is right X% of the time.” Anything fancier, like model comparisons, prompt tweaks, and regression checks, was downstream of having that baseline.
There were four document types: PAN (India’s Income Tax ID), Aadhaar (Indian National ID card), DL (Indian Driver’s License), and RC (Indian Vehicle Registration Certificate). Each has its own layout and failure modes, and a model can’t transfer what it learned from one to another. A single blended accuracy number across all four would have buried the real problem: an LLM that reads one document type perfectly might fail on another. So I split the golden dataset into four slices, one per document type, and treated each as a separate problem with its own baseline and room to iterate.
I included a mix of inputs: flatbed scans and phone photos, glare and skew, the long tail of state-wise DL and RC layouts, and documents with regional-language text alongside English. This kept each evaluation tractable while still reflecting real-world data.
How big should the golden dataset be? It needed to be big enough to be representative of production, small enough that the expert could review every run end-to-end. I settled on 50 documents per document type. That was enough to cover the realistic mix within each slice. Iteration was fast. I could review a full run in hours and do it again whenever I changed a prompt or swapped a model.
One thing that’s easy to skip and shouldn’t be: the rubric. The rubric is half the dataset. “Did the LLM read this correctly?” is not a yes/no question even for a single PAN card. I sat with the domain expert and pinned down, field by field, what counted as acceptable:
ID numbers (PAN, Aadhaar, DL number, registration number): exact match, character for character. One wrong character is a failure, full stop. Format-valid hallucinations are the worst kind of failure and are flagged separately from honest misreads.
Names: do we tolerate ANIL KUMAR vs Anil Kumar? Anil vs Anil Kumar (last name trimmed). Anik vs Anil (character confusion). Anil K vs Anil Kumar (abbreviation). Seshadri vs Sesadri (transliteration variance).
Date of birth: normalised to a canonical format before comparison, so 01/04/1990 and 1990-04-01 aren’t punished as different.
Hallucinations vs omissions: a hallucinated field is always scored worse than a null or [unreadable]. The rubric says so explicitly, because in KYC, “I don’t know” is recoverable and “here’s a confident wrong answer” is not.
With the rubric in hand, I moved to human review. For identity documents, no automated check catches what a trained eye does. When a scan is unreadable, LLMs don’t say “I don’t know”. They hallucinate plausible-looking IDs: a 10-character string in the right AAAAA9999A shape, a 12-digit Aadhaar that passes the checksum, a name read into the father’s name slot. Reviewing each extraction side by side with the source image, I could see the failure modes started to repeat in the same handful of shapes. These patterns were the only things that told me what to fix next.
With the dataset, the rubric, and a reviewer in place, I had an eval loop. Run a prompt against all 50 documents in a slice, collect the extractions, score each one against the rubric, and walk away with two things: a number for that slice and a list of why each failure happened. The list, not the number, was what drove the next prompt change. Then run it again.
My first prompt was a one-liner. Something close to “Extract the following fields from this image.” I was running this against Gemini 2.0 Flash. The client had a preference for the Gemini family, and Flash was the cheapest vision model in it. I picked PAN to start with, the simplest layout and fewest fields. I got that working, and only then moved on to Aadhaar, DL, and RC. From there, the iterations on each prompt fell into three rough phases.
Phase 1: Structure . The first round of changes had nothing to do with reading documents better, only with how the prompts were organised. I gave each document type its own prompt. Trying to write one that handled all four had been a losing battle, since the layouts, critical fields, and failure modes were too different. I separated system and user prompts so instructions and inputs weren’t tangled. I pinned down the output schema explicitly instead of letting the model decide. I dropped the temperature close to zero. And I rewrote the instructions to be short and direct. Every “please” and “kindly” came out. Standard prompt-engineering moves, mostly. They did real work, but only got me partway.
Phase 2: Domain context . The next round was about teaching the model what it was looking at. A naked “extract the DL number” didn’t help when the document had multiple numbers on it. What helped was spelling out the rule for the field itself. For DLs, “Indian Driving Licence numbers follow state-specific formats like KA-XX-YYYYNNNNNNN”. I added similar context for each ID type. Aadhaar is exactly 12 digits; PAN follows AAAAA9999A, and so on. I supplemented the prompt with a few-shot examples drawn from the client’s own annotations. The model was now wrong less often, and when it was, the wrongness was easier to spot.
Phase 3: Decomposition . The change that mattered most was in how the work was split up. Instead of having one prompt return all fields for a document type at once, I split the extraction into multiple focused calls, one per field. The single-call version was juggling too much: layout, language, format, and field semantics, with errors in one field bleeding into others. Per-field calls were slower and more expensive, but each call was simple enough that the model stopped tripping over itself.
By the end of these iterations, the system achieved 92% accuracy on the golden set at 10% of the third-party service’s cost .
When we shipped, the model itself had improved enough that I could drop the decomposition and go back to a single call per document. The workaround had quietly outlived its usefulness.
Looking back a year later, two things would change.
I’d add chain-of-thought to the prompt. Asking the model to first describe what it sees in the document and then extract it would have caught many of the errors I was chasing through prompt rewrites.
I’d also explore using an LLM-as-judge to speed up the iteration loop. With the rubric and a few hundred labelled examples in place, a judge could compare each extraction against the ground truth and surface failures, letting me review only what got flagged instead of every document. The caveat: the judge can be wrong too, so its calls would need their own evaluation, a topic for another post.
What I wouldn’t change is the engine: the golden dataset , the rubric , the eval loop . The prompts just tuned it. I could have rewritten instructions endlessly, swapped models, or tweaked the temperature. Instead, I built a baseline, nailed down what correct meant, and ran the same documents through the same rubric after every change. It’s unglamorous work. But it’s where the actual work happened.
Temporal anti-pattern: Don't treat expected failures as exceptions
⌂ Archives
» Training a small model to write better OCaml with RLVR and GRPO
Building the systems that build the software
Our approach to evals with megasthenes
Training a small model to write better OCaml with RLVR and GRPO
Evals before prompts: building an LLM OCR for KYC
nilenso is an employee-owned software cooperative based out of Bangalore, India.
We practice test driven development and continuous delivery, and love working with Clojure and Ruby on Rails.
This blog is a showcase of our growth as a consultancy, a product company and generally curious beings.
Get in touch with us at hello@nilenso.com .
