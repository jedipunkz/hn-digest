---
source: "https://zeltser.com/ai-security-market-field-guide"
hn_url: "https://news.ycombinator.com/item?id=49133934"
title: "A Field Guide to the AI Security Market"
article_title: "A Field Guide to the AI Security Market"
author: "882542F3884314B"
captured_at: "2026-08-01T12:58:45Z"
capture_tool: "hn-digest"
hn_id: 49133934
score: 2
comments: 0
posted_at: "2026-08-01T12:39:39Z"
tags:
  - hacker-news
  - translated
---

# A Field Guide to the AI Security Market

- HN: [49133934](https://news.ycombinator.com/item?id=49133934)
- Source: [zeltser.com](https://zeltser.com/ai-security-market-field-guide)
- Score: 2
- Comments: 0
- Posted: 2026-08-01T12:39:39Z

## Translation

タイトル: AI セキュリティ市場のフィールド ガイド
説明: AI を保護する製品の市場は、それぞれが多くの役割を担うカテゴリや製品が重複しており、混雑していて読みにくいです。 AI Defense Matrix を使用すると、適切な質問をし、製品が何のために作られたのかを伝え、製品がまだ埋めていないギャップを埋めることができます。

記事本文:
メイン コンテンツにスキップ Lenny Zeltser セキュリティ ビルダー兼リーダー プロジェクト 執筆について × AI セキュリティ市場のフィールド ガイド
AI を保護する製品の市場は、それぞれが多くの役割を担うカテゴリや製品が重複しており、混雑していてわかりにくいです。 AI Defense Matrix を使用すると、適切な質問をし、製品が何のために作られたのかを伝え、製品がまだ埋めていないギャップを埋めることができます。
Sounil Yu 氏は、一般的なセキュリティ製品の市場を、整理整頓のための通路もなく商品が山積みになっている、整理整頓されていない食料品店に例えています。 AIセキュリティ市場もそのような店です。製品カテゴリは重複しており、ベンダーは機能ではなく目標に基づいて製品に名前を付けており、同じデモで複数の売り込みをサポートできます。買い手は必要なものを見つけることができず、売り手は製品の機能を説明するのに苦労しています。
Sounil と私は、AI セキュリティ市場を理解するために AI Defense Matrix Catalog を作成しました。カタログは各製品を AI 防御マトリックスのセルにマッピングします。これは 8 つの AI 資産クラスと 6 つの NIST CSF 機能を横断しており、各セルは 1 つの資産クラスと 1 つのセキュリティ機能を組み合わせています。これは、製品が集中している場所と、カバレッジが少ない場所を示しています。私は、AI セキュリティに関する記事の形で、その分布が市場の構造について何を明らかにしているかを調査しました。このガイドは、AI を保護する製品を評価する際にベンダーに尋ねるべきクラスターごとの質問を記載した購入者の相棒です。
4 つのクラスターが混雑したセルを埋めます。
カタログの混雑したセルには 4 つの製品クラスターがあり、それぞれに購入前に尋ねるべき独自の質問があります。
AI ランタイム ガードレール: カタログ最大のクラスターに含まれる製品は、プロンプト、推論トラフィック、RAG コンテンツ、エージェント メモリを意味するランタイム AI データを保護および監視します。これらの製品は何をフィルタリングしますか

モデルを削除し、機密データを編集し、ポリシー違反にフラグを立てます。このパターンは、Web アプリケーション ファイアウォールや DLP、およびほとんどの保護と検出が 1 つの製品にバンドルされていることでよく知られています。このカタログには、ランタイム AI データをカバーする 100 以上の製品がリストされています。 Frank Wang 氏は、Frankly Speaker の中で、「業界として、どのようなガードレールを施行すべきか、あるいは内部エージェントにどのレベルの自律性を付与すべきかについては、まだ完全には把握できていない」と書いています。製品がどこでトラフィックを検査するのか、どのプロンプトやその他のデータが境界を離れるのか、ベンダーのサービスがダウンしたときに AI アプリケーションに何が起こるのかを尋ねてください。
エージェントの ID とアクセス: 2 番目の主要クラスターの製品は、AI エージェントを、資格情報、権限スコープ、ライフサイクル管理を必要とする人間以外のプリンシパルとして扱います。 SailPoint や Saviynt などの既存のアイデンティティ企業はここでプラットフォームを拡張し、数十のスタートアップがそれに加わりました。エージェントは、アイデンティティ管理というよく知られた分野における新しい種類のプリンシパルです。この製品が委任チェーンをどのように処理するか、エージェントが人間に代わって行動するか、別のエージェントを通じて行動するか、また、実行中の委任に何が起こるかを含め、不正行為を行ったエージェントの資格情報をどのくらい早く取り消すことができるかについて質問してください。
モデルのスキャンとポスチャ: 3 番目のクラスターの製品は、AI モデルとパイプラインを検査し、モデル ファイルの改ざんをスキャンし、敵対的な弱点をテストし、どのモデルがどこで実行されるかをインベントリします。これは、モデルを対象とした脆弱性管理とレッド チームと考えてください。チームが調査結果を修正に反映できるかどうか、またスキャナーがサポートするモデル形式とパイプラインを検討してください。
オーケストレーションと MCP セキュリティ: 4 番目のクラスターの製品は、MCP サーバー、プラグイン、エージェント周りのスキャフォールディングを含む AI オーケストレーション ツールを保護します。聞く

製品がエージェントとプラグイン全体で認識できるもの、制限できるもの、および範囲外で動作するエージェントを検出する方法。
いくつかの製品は空に近いセルをカバーします。
Respond または Recover 機能をカバーする製品は 12 未満であり、それらのいくつかは、ID の一時停止によるエージェントの封じ込め、モデルの回復として位置付けられるマシンのアンラーニング (Hirundo)、エージェントのアクションのロールバック (Rubrik Agent Cloud) など、AI インシデント対応ツールの初期のスケッチにすぎません。 AI セキュリティの記事で詳しく説明したように、これらの機能はテクノロジーよりも人に依存しています。 AI インシデントへの対応と回復を、従業員、ハンドブック、およびすでに所有している回復ツールの作業として扱います。
ほとんどの製品は、設計上の目的よりも多くのセルをカバーします。
AI 防御マトリックス カタログには、各適用範囲の主張が製品の目的にとってどの程度中心的であるかが記録されます。カタログに記載されている適用範囲の約 3 分の 1 は「二次的な」または「隣接した」もので、これは製品の中心的な目的ではなく、製品がサポートする領域を意味します。カタログ内のほとんどの製品は 3 つ以上のセルをカバーします。ベンダーが広範な内容を提案する場合は、その製品が何をするために作られたのか、ベンダーが途中でどの機能を追加したかを尋ね、その回答を製品のカタログ エントリと照合します。
評価チェックリストにベンダーの実行可能性も追加します。カタログに掲載されている製品の 10 分の 1 以上をすでに買収者が吸収しており、市場はまだ数年しか経っていません。混雑したセルは統合を続けることが予想されます。管理変更条件、データのポータビリティ、および別のベンダーが製品を吸収した場合の移行パスについて質問してください。
製品を評価する前に、どのセルが必要かを決定してください。
AI セキュリティ要件に照らして AI 防御マトリックスを確認し、ベンダーの議論に備えます。
mを列挙してください

使用する各 AI アセットとプログラムに必要な機能を組み合わせることで、AI デプロイメントにとって最も重要な atrix セルを構築します。
新しい製品を評価する前に、すでに実行している製品がそれらのセルをカバーしているかどうかを確認してください。
混雑したセルの場合は、いくつかのベンダーを選択し、独自の環境で評価してください。クラスター内の製品は、デモよりも操作の点で異なります。
カタログ内の各製品の AI 防御マトリックス マッピングを確認して、議論をより適切に進めてください。
カタログ内の空のセルに対処する必要がある場合は、製品を待つのではなく、指名された人が所有するプロセスを考案します。
AI 導入に必要なセルに製品を適合させ、残りには人材とプロセスを配置します。
私のブログ投稿を電子メールで受信します。
Lenny Zeltser は、深い技術的ルーツ、製品管理の経験、ビジネスの考え方を備えたサイバーセキュリティの幹部です。彼は初期段階からエンタープライズ規模までセキュリティ製品とプログラムを構築してきました。彼は SANS Institute のファカルティフェローでもあり、マルウェア分析用の人気のある Linux ツールキットである REMnux の作成者でもあります。 Lenny は、セキュリティのリーダーシップとテクノロジーに関する彼の見解を zeltser.com で共有しています。

## Original Extract

The market for products that secure AI is crowded and hard to read, with overlapping categories and products that each do many jobs. With the AI Defense Matrix, you can ask the right questions, tell what a product was built for, and staff the gaps no product fills yet.

Skip to main content Lenny Zeltser Security builder & leader Projects Writing About × A Field Guide to the AI Security Market
The market for products that secure AI is crowded and hard to read, with overlapping categories and products that each do many jobs. With the AI Defense Matrix, you can ask the right questions, tell what a product was built for, and staff the gaps no product fills yet.
Sounil Yu compares the general security product marketplace to a disorganized grocery store , where products pile up without aisles to organize them. The AI security market is such a store, too. Product categories overlap, vendors name products after ambitions rather than capabilities, and the same demo can support multiple pitches. Buyers can’t find what they need, and sellers struggle to explain what their products do.
Sounil and I built the AI Defense Matrix Catalog to make sense of the AI security marketplace. The catalog maps each product to the cells of the AI Defense Matrix , which crosses eight AI asset classes with six NIST CSF functions, so each cell pairs one asset class with one security function. It illustrates where products cluster and where coverage is light. I explored what that distribution reveals about the market’s structure in the shape of AI security article. This guide is the buyer’s companion, with cluster-by-cluster questions to ask vendors when evaluating a product that secures AI.
Four clusters fill the crowded cells.
There are four product clusters in the catalog’s crowded cells, each with its own questions to ask before you buy.
AI runtime guardrails: The products in the catalog’s largest cluster protect and monitor runtime AI data, meaning prompts, inference traffic, RAG content, and agent memory. These products filter what enters and leaves models, redact sensitive data, and flag policy violations. The pattern is familiar from web application firewalls and DLP, and most bundle protection and detection in one product. The catalog lists more than a hundred products that cover runtime AI data. Frank Wang writes in Frankly Speaking that “as an industry, we haven’t even fully figured out what guardrails we should be enforcing , or what level of autonomy we should grant an internal agent.” Ask where the product inspects traffic, which prompts and other data leave your boundary, and what happens to your AI application when the vendor’s service goes down.
Agent identity and access: Products in the second major cluster treat AI agents as non-human principals that need credentials, permission scopes, and lifecycle management. Identity incumbents such as SailPoint and Saviynt extended their platforms here, and dozens of startups joined them. Agents are a new kind of principal in the familiar discipline of identity management. Ask how the product handles delegation chains, where an agent acts for a human or through another agent, and how quickly you can revoke a misbehaving agent’s credentials, including what happens to its in-flight delegations.
Model scanning and posture: Products in the third cluster inspect AI models and pipelines, scanning model files for tampering, testing for adversarial weaknesses, and inventorying which models run where. Think of it as vulnerability management and red teaming aimed at models. Ask whether your team can turn the findings into fixes, and which model formats and pipelines the scanner supports.
Orchestration and MCP security: Products in the fourth cluster secure AI orchestration tools, including MCP servers, plugins, and the scaffolding around agents. Ask what the product can see across your agents and plugins, what it can restrict, and how it detects an agent acting outside its scope.
A few products cover the near-empty cells.
Fewer than a dozen products cover any Respond or Recover capability, and those few are an early sketch of AI incident response tooling, including agent containment through identity suspension, machine unlearning positioned as model recovery (Hirundo), and rollback of agent actions (Rubrik Agent Cloud). These functions depend on people more than on technology, as I explored in the shape of AI security article. Treat response and recovery for AI incidents as work for your people, your playbooks, and the recovery tooling you already own.
Most products cover more cells than they were built for.
The AI Defense Matrix Catalog records how central each coverage claim is to the product’s purpose. About a third of the coverage claims in the catalog are “secondary” or “adjacent,” meaning areas the product supports rather than its central purpose. Most products in the catalog cover three or more cells. When a vendor pitches broad coverage, ask what the product was built to do and which capabilities the vendor added along the way, then check the answers against the product’s catalog entry.
Add vendor viability to your evaluation checklist as well. Acquirers have already absorbed more than one in ten of the products in the catalog, and the market is only a few years old. Expect the crowded cells to keep consolidating. Ask about change-of-control terms, data portability, and the migration path if another vendor absorbs the product.
Decide which cells you need before you evaluate products.
Review the AI Defense Matrix against your AI security requirements to prepare for vendor discussions:
List the matrix cells that matter most for your AI deployments by pairing each AI asset you use with the functions your program needs for it.
Check whether the products you already run cover those cells before you evaluate new ones.
For crowded cells, select a few vendors and evaluate them in your own environment. Products within a cluster differ more in operations than in demos.
Check each product’s AI Defense Matrix mappings in the catalog to better steer those discussions.
For cells you need to address that are empty in the catalog, devise a process, owned by a named person, instead of waiting for a product.
Match products to the cells your AI deployments need, and staff the rest with people and processes.
Receive my blog posts by email.
Lenny Zeltser is a cybersecurity executive with deep technical roots, product management experience, and a business mindset. He has built security products and programs from early stage to enterprise scale. He is also a Faculty Fellow at SANS Institute and the creator of REMnux, a popular Linux toolkit for malware analysis. Lenny shares his perspectives on security leadership and technology at zeltser.com .
