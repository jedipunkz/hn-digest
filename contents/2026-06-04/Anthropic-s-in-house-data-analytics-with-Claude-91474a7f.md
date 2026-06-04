---
source: "https://claude.com/blog/how-anthropic-enables-self-service-data-analytics-with-claude"
hn_url: "https://news.ycombinator.com/item?id=48395883"
title: "Anthropic's in-house data analytics with Claude"
article_title: "How Anthropic enables self-service data analytics with Claude | Claude"
author: "dmpetrov"
captured_at: "2026-06-04T10:22:42Z"
capture_tool: "hn-digest"
hn_id: 48395883
score: 2
comments: 0
posted_at: "2026-06-04T08:38:41Z"
tags:
  - hacker-news
  - translated
---

# Anthropic's in-house data analytics with Claude

- HN: [48395883](https://news.ycombinator.com/item?id=48395883)
- Source: [claude.com](https://claude.com/blog/how-anthropic-enables-self-service-data-analytics-with-claude)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T08:38:41Z

## Translation

タイトル: Claude による Anthropic の社内データ分析
記事のタイトル: Anthropic が Claude を使用してセルフサービス データ分析を実現する方法 |クロード
説明: セルフサービスのデータ洞察を推進するクロードの能力を最大化するためのヒントとアプローチ

記事本文:
Anthropic が Claude を使用してセルフサービス データ分析を実現する方法 |クロード
クロード製品のご紹介 クロード
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
Anthropic が Claude を使用してセルフサービス データ分析を実現する方法
Anthropic が Claude を使用してセルフサービス データ分析を実現する方法
リンクをコピーする https://claude.com/blog/how-anthropic-enables-self-service-data-analytics-with-claude
多くのデータ サイエンスおよびデータ エンジニアリング チームが証明しているように、セルフサービスのビジネス分析を実現することは伝統的に困難な課題でした。
幅が広く非正規化されたテーブルを使用して、技術に乏しい同僚でもデータ モデルにアクセスしやすくすると、ビジネスの規模が拡大するにつれて、ビューが重複して定義が矛盾することになることがよくあります (SQL を学習する意欲がほとんどない従業員にとっては、ギャップを埋めることはほとんどできません)。あるいは、ユーザーのためにさらにリングフェンスで囲まれた環境を作成すると、多くの場合、ビジネス上の質問のロングテールが見逃され、チームの作業がサイロ化されるにつれてメトリクスとダッシュボードの肥大化につながります。
LLM の台頭により、これらの課題を回避するセルフサービス分析のための追加の道が提供されます。ただし、クロードを倉庫に向けてエージェントに実行させると、誤った正確性が生じる可能性があります。
アドホックなリクエストから解放されるという最初の高揚感は、この設定により、関係者が、以前は管理されていた基盤となるインフラストラクチャ、ドキュメント、専門知識から切り離されてしまうことに気づくと恐怖に変わります。

慎重に厳選されたデータセットに向けて。
Anthropic では、ビジネス分析クエリの 95% が Claude によって自動化されており、合計で最大 95% の精度が得られます。この暗記的で反復的な作業をクロードに任せることで、データ サイエンス チームは因果モデリング、予測、機械学習などのより戦略的な作業に集中できるようになります。
Anthropic のクロード コードのトップ ユーザー数十人と会い、分析エージェントの無数の設計パターンを確認した後、私たちは LLM を使用する他のデータ チーム向けのベスト プラクティスをいくつか確立しました。この投稿では、セルフサービスのビジネス洞察を推進するクロードの能力を最大限に高めるためのヒントとアプローチを共有します。
分析の精度がコード生成の問題ではなく、コンテキストと検証の問題である理由。
ほとんどのエラーを引き起こす 3 つの障害モード。
これらのエラーに対処するために私たちが構築したエージェント分析スタック。
有効性を測定する方法。そして
スキルの大部分を作成する方法の基本テンプレート (付録を参照)
LLM の生成能力は両刃の剣です。複雑な問題に対する創造的な解決策を可能にするメカニズムは、誤った出力を幻覚する可能性もあります。分析エージェントの課題を完全に理解するには、分析エージェントをコーディング エージェントと比較すると役立ちます。
コーディングは、モデルの創造性に報いるオープンエンドのソリューション空間であり、文書化とテストは幻覚に対する自然なガードレールを提供します。対照的に、分析のユースケースでは、多くの場合、単一の正しいソースを使用した単一の正解しかなく、その正しさを証明する決定的な方法がありません。
セルフサービス エージェントによるビジネス分析の場合、複雑さは主にデータの曖昧さにあります。中心的な問題は、ユーザーの質問をデータ モデル内の特定の最新のエンティティにマッピングする能力にあります。

そしてそれらを扱う正しい方法を知ってください。それができれば、結果として生じる実行と SQL は簡単になります。
私たちは、不正確な応答の大部分を占めるこの問題の 3 つの属性を特定しました。
コンセプト <> エンティティのあいまいさ : データ モデルには (数百万ものフィールドの中から) 数百の実行可能なオプションがあるため、エージェントはユーザーの質問に最もよく答える正しいフィールドを選択できません。たとえば、アクティブ ユーザーの数を測定する場合、どのようなアクションが「アクティブ」を構成するのでしょうか?不正ユーザーも含まれますか?どのようなルックバック ウィンドウを使用しますか?
データの古さ : データ ソース、ビジネス定義、スキーマは常に変化します。資産やエージェントの知識が古くなり、微妙に間違った答えが返され始めます。
取得の失敗: 正しい情報が実際にはデータ モデル内に存在し、適切に注釈が付けられている可能性がありますが、検索スペースが広大であるため、エージェントは単にそれを見つけられません。
前へ 前へ 0 / 5 次へ 次へ Claude Code Desktop を入手
IRM https://claude.ai/install.ps1 | iex コマンドをクリップボードにコピーする またはドキュメントを読む クロード コードを試す クロード コードを試す クロード コードを試す 開発者ドキュメント 開発者ドキュメント 開発者ドキュメント 電子書籍
Anthropic では、これら 3 つのエラーを最小限に抑える主な方法は、エージェント データ スタックを使用することです。各層は主に、次の 1 つ以上の問題を攻撃するために存在します。
エンティティの曖昧さ : データ基盤と真実の情報源は、単一の統制された答えが見つかるまで、妥当なエンティティの領域を縮小します。
古さ: メンテナンスと検証のプロセスにより、ビジネスの変化に応じてすべてが腐ることがなくなります。
取得失敗 : スキルにより、エージェントがその回答を確実に見つけて正しく使用できるようになります。
このセクションでは、各レイヤーをどのように構築したかについて説明します。
分析エージェントの安全性を確保する上で最も重要な側面は、

正確性は、データ ウェアハウス内のデータ モデル、変換、テスト、テーブルと、それらを記述するメタデータを含む強力なデータ基盤によって実現されます。ディメンション モデリング、シフトレフト テスト、重要なパイプラインの鮮度や完全性のチェックなどの標準的なデータ エンジニアリングとデータ品質の実践はすべて引き続き適用されます (これらを再検証するつもりはありません)。
何が変わるかというと、データ モデルのエンド ユーザーがデータの専門家 (データ サイエンティストなど) ではなくなり、さまざまな程度のデータの専門知識や基盤となるインフラストラクチャの理解をユーザーに代わって行動するエージェントになることです。この変化は、エンド ユーザーが知らないという理由だけで、結果の根本的な正しさをユーザーが検証することを要求できないという点で課題を引き起こします。
データ基盤レイヤーは主に曖昧さを目的としています。たとえば、収益が 40 の妥当な候補ではなく 1 つの管理されたデータセットに解決される場合、エージェントが検索する前に問題はほとんど解消されます。正規モデルを定義する同じリポジトリは、正規モデルが最新であることを強制する自然な場所であるため、最初の失効防御が存在する場所でもあります。
私たちは、いくつかの実践が特にうまく機能することを確認しました。
正規データセットを作成する : これまでで最も一般的な失敗は、エージェントがコンセプト (「製品 X の収益」) を単一の正しいテーブル、列、メトリクス定義にマッピングできないことです。これは通常、実装が微妙に異なる妥当な候補が複数あるためです。修正方法は、数を減らし、より厳しく管理された論理モデルです。所有者が明確で、すぐに使用でき、検出可能な、正規の単一の真実の情報源である小規模なデータセットを厳選し、重複に近いデータセットを積極的に非推奨にします。物理的なロールアップとキャッシュは依然としてコストとパフォーマンスの点で重要ですが、

それらは、代替モデルとして並存するのではなく、標準モデルから機械的に派生する必要があります。目標は、エージェントが概念を検索するときに、単一の統制された答えを見つけることです。
標準を強制する : 正規モデルとメトリクス定義が、ツール (構造的にエージェントが最初にルーティングされる。詳細は後述)、CI (変更をバイパスする変更はレビューに失敗する)、および権限 (下流チームが管理されたレイヤー上に構築するか、その理由を説明する) によって強制される場合にのみ、基礎が成立することがわかりました。そうしないと、強制力のないガバナンスはすぐに複数の候補者問題に戻ってしまいます。
アーティファクトのコロケーション : 絶えず変化するデータ モデルとビジネス ロジックに対する主な防御策はコロケーションです。ほぼすべてのデータ コード (つまり、モデリング、セマンティック レイヤー、リファレンス ドキュメント、正規のダッシュボード定義) が単一のリポジトリ内に存在し、レイヤー間の整合性を保護する CI チェックが行われます。モデリングの変更によってダウンストリームのダッシュボードが破損したり、文書化されたメトリクスが無効になったりする場合、CI はその変更にフラグを立て、修正が同じ PR で提供されます。 (このメカニズムについては、以下の「スキル」セクションで説明します。)
メタデータを一流の製品として扱う : コーディング エージェントがうまく機能するのは、コードベース (README、型署名、ドキュメント文字列など) が読みやすいためでもあります。ウェアハウスも同様に読みやすくなりますが、それは列とテーブルの説明、標準的なメトリクスの定義、グレイン ドキュメント、有効な値の範囲、系統、所有権、およびモデルの階層化が変換自体と同じ厳密さで維持されている場合に限られます。新しい洞察ではありませんが、優れたガバナンスは、エージェントが適切なデータセットを選択するのに役立つ重要なコンテキストを提供します。
データ基盤がデータ ウェアハウスそのものである場合、信頼できる情報源は、エージェントがデータ ウェアハウスをナビゲートするために参照する参照面です。この層は

概念 <> エンティティのあいまいさを軽減し、利害関係者の質問における「週間アクティブ ユーザー」を、データ モデル内の特定の管理対象エンティティに変換します。おおよそ信頼度の降順に並べると、次のようになります。
セマンティック レイヤー: コンパイルされたメトリックとディメンションの定義。質問が定義された指標に明確にマッピングされている場合、エージェントは関数を呼び出して 1 つの数値を取得します。この数値は、社内の他のすべてのサーフェスが生成するものと同じです。私たちのエージェントは構造上、（スキルの指導により）最初にセマンティック層を活用する必要があります（付録を参照）。私たちが試してみましたがうまくいきませんでした。LLM に生のテーブルとクエリ ログからメトリクス定義を自動生成させることで、セマンティック レイヤーをブートストラップするというものです。それは、私たちが排除しようとしていたまさにあいまいさをエンコードした、もっともらしく見える定義を生成し、人間が厳選したより小さな層と比較した評価ではネットネガティブでした。したがって、Claude でドキュメントを生成し、定義は人間が所有することをお勧めします。
リネージと変換グラフ: セマンティック レイヤーが質問をカバーしていない場合、リネージとテーブルのランキング (参照数に基づく) により、エージェントはどの上流モデルが概念にフィードするか、どのモデルが非推奨で、どのモデルが粒度を共有するかを推論できます。これにより、「メトリクスがわからない」が「どの管理モデルから集計するかはわかっている」に変わります。これは、以下のオンライン検証で明らかになった鮮度と産地のシグナルのバックボーンでもあります。
クエリ コーパス: ダッシュボード、ノートブック、および以前の分析からの履歴 SQL。直感的に、これは価値の高いものであるはずです。これは、既に正解したすべての質問の記録です。実際に、エージェントに何千もの以前のクエリへの生の検索アクセスを与えると、精度が 1 ポイント未満向上することがわかりました (その除去については、後のセクションで説明します)。

非構造化検索では、新しい質問を適切な前例にマッピングできませんでした。機能するのは、そのコーパスを構造化されたドメインごとのリファレンス ドキュメントと、スキルで説明されている再利用可能な分析パターンに抽出することです。クエリ履歴を、エージェントが直接読み取る真実の情報源としてではなく、キュレーションのための原材料として扱います。
ビジネスの背景: ほとんどのチームがスキップする層であり、私たちが最も長く過小評価していた層です。あなたのビジネスを理解していないエージェントは、ユーザーの質問には答えますが、その意味するところは答えません。 「第 2 四半期の発売」が特定の製品を指していること、同じ用語を 2 つのチームが異なる定義で定義していること、または木曜日に取締役会があるために質問が行われていることがわかりません。エージェントが周囲の参照を解決し、より明確な質問をできるように、インデックス付きのドキュメント、ロードマップ、意思決定ログ、および組織構造で構成される社内ナレッジ グラフをパイプします。
4 つすべてに共通する失敗パターンは、データ基盤層からのものと同じです。つまり、ドキュメントが貧弱または古いです。 Claude はギャップを埋めるのに非常に役立ちます (列の説明の草案、クエリ パターンからのメトリクス ドキュメントの提案、CI で文書化されていないモデルのフラグ付け) ですが、キュレーションと所有権は人間によって管理されます。
次の 2 つのセクションでは、その作成方法について説明します。

[切り捨てられた]

## Original Extract

Tips and approaches to maximizing Claude’s ability to drive self-serve data insights

How Anthropic enables self-service data analytics with Claude | Claude
Meet Claude Products Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
How Anthropic enables self-service data analytics with Claude
How Anthropic enables self-service data analytics with Claude
Share Copy link https://claude.com/blog/how-anthropic-enables-self-service-data-analytics-with-claude
As many data science and data engineering teams can attest, enabling self-service business analytics has traditionally been a slog.
Making the data model more accessible to less technical coworkers via wide and denormalized tables often leads to overlapping views with inconsistent definitions as the business scales (and does little to bridge the gap for employees with little desire to learn SQL). Alternatively, creating more ringfenced environments for users often misses the long tail of business questions and leads to metric and dashboard bloat as teams silo their work.
The rise of LLMs provides an additional path for self-service analytics that avoids those challenges. However, pointing Claude at a warehouse and letting the agents execute can create a false sense of precision.
The initial elation of liberation from ad-hoc requests turns into dread with the realization that this setup separates stakeholders from the underlying infrastructure, documentation, and expertise that previously steered them toward carefully curated datasets.
At Anthropic, 95% of business analytics queries are automated via Claude, with ~95% accuracy in aggregate. By giving this often rote, repetitive work to Claude, our data science team can focus on more strategic work like causal modeling, forecasting, and machine learning.
After meeting with dozens of Anthropic’s top Claude Code users and having seen myriad design patterns for analytics agents, we’ve cultivated some best practices for other data teams working with LLMs. In this post, we’ll share these tips and approaches to maximizing Claude’s ability to drive self-serve business insights, including:
Why analytics accuracy is a context and verification problem, not a code generation issue;
The three failure modes that cause most errors;
The agentic analytics stack we built to address these errors;
How we measure effectiveness; and
A basic template for how we create the majority of our skills (see the appendix)
LLMs' generative abilities are a double-edged sword: the mechanisms that enable creative solutions to complex problems can also hallucinate erroneous output. To fully understand the challenges with analytics agents, it’s useful to compare them to coding agents.
Coding is an open-ended solution space that rewards the models' creativity, while documentation and tests provide natural guardrails against hallucination. In contrast, for analytics use cases, there’s often only a single correct answer using a single correct source in which there’s no deterministic way of proving the correctness.
For self-service agentic business analytics, the complexity mainly lies in the ambiguity of the data. The central problem comes down to our ability to map a user’s question to specific and up-to-date entities in our data model and know the correct way of working with them . If we can do that, then the resulting execution and SQL becomes trivial.
We’ve identified three attributes of this problem that account for an overwhelming majority of inaccurate responses:
Concept <> entity ambiguity : with hundreds of viable options in a data model (out of potentially millions of fields), the agent is unable to choose the correct fields that best answer a user’s question. For example, in measuring the number of active users: what actions constitute being “active”? Do you include fraudulent users? What lookback window do you use?
Data staleness : data sources, business definitions, and schemas change constantly; assets and agent knowledge go stale and start returning subtly wrong answers.
Retrieval failure : the right information may actually be in the data model and properly annotated, but given the vastness of the search space, the agent simply doesn’t find it.
Prev Prev 0 / 5 Next Next Get Claude Code Desktop
irm https://claude.ai/install.ps1 | iex Copy command to clipboard Or read the documentation Try Claude Code Try Claude Code Try Claude Code Developer docs Developer docs Developer docs eBook
At Anthropic, the main way we minimize these three errors is via our agentic data stack. Each layer exists primarily to attack one or more of these problems:
Entity ambiguity : data foundations and sources of truth shrink the space of plausible entities until there's a single governed answer.
Staleness : maintenance and validation processes keep everything from rotting as the business changes.
Retrieval failure : skills make sure the agent reliably finds and correctly uses that answer.
In this section, we’ll discuss how we built each layer.
The most important aspect of ensuring analytics agents are accurate is via strong data foundations, which include the data models, transforms, tests, and tables in a data warehouse, along with the metadata describing them. Standard data engineering and data quality practices such as dimensional modeling , shift-left testing, freshness and completeness checks on critical pipelines all still apply (and we won't relitigate these).
What does change is that the end user of your data model is no longer a data expert (e.g. data scientist), but rather agents acting on behalf of users with varying degrees of data expertise or understanding of the underlying infrastructure. This shift presents a challenge in that the results can’t require the user to validate the underlying correctness simply because the end user doesn’t know.
The data foundations layer is aimed primarily at ambiguity: if revenue , for example, resolves to one governed dataset instead of forty plausible candidates, the problem largely disappears before the agent ever has to search. It's also where the first staleness defense lives, since the same repo that defines the canonical models is the natural place to enforce that they stay current.
We’ve seen a few practices work especially well:
Create canonical datasets : By far the most common failure is that the agent can’t map a concept (“revenue for product X”) to the single correct table, column, and metric definition, usually because there are multiple plausible candidates with subtly different implementations. The fix is fewer, more heavily governed logical models: curate a small set of canonical, single source-of-truth datasets that are clearly owned, consumption-ready, and discoverable, then aggressively deprecate the near-duplicates. Physical rollups and caches still matter for cost and performance, but they should derive mechanically from the canonical models rather than living alongside them as alternatives. The goal is that when an agent searches for a concept, it finds a single governed answer.
Enforce your standards : We’ve found the foundations only hold if the canonical models and metric definitions are enforced by tooling (the agent is structurally routed to them first; more on that below), by CI (changes that bypass them fail review), and by mandate (downstream teams build on the governed layer or explain why not). Governance without enforcement otherwise quickly decays back to the multiple candidates problem.
Colocate artifacts : Our main defense against constantly changing data models and business logic is colocation. Nearly all data code (i.e., modeling, semantic layer, reference docs, canonical dashboard definitions) lives in a single repo, with CI checks that protect cross-layer integrity. If a modeling change would break a downstream dashboard or invalidate a documented metric, CI flags it and the fix ships in the same PR. (We’ll come back to the mechanics of this in the Skills section below.)
Treat metadata as a first-class product : Coding agents perform well partly because codebases are legible : READMEs, type signatures, docstrings, etc. Your warehouse can be just as legible, but only if column and table descriptions, canonical metric definitions, grain documentation, valid value ranges, lineage, ownership, and model tiering are maintained with the same rigor as the transformations themselves. While not a new insight, good governance provides critical context that helps the agent choose the right dataset.
If data foundations are the data warehouse itself, sources of truth are the reference surfaces the agent consults to navigate it. This layer reduces concept <> entity ambiguity and turns “weekly active users” in a stakeholder’s question into a specific, governed entity in your data model. Roughly in descending order of trust:
Semantic layer: the compiled metric and dimension definitions. If a question maps cleanly to a defined metric, the agent calls a function and gets one number, the same number every other surface in the company produces. Our agents are structurally required (by skill instruction) to leverage the semantic layer first (see the appendix). One idea we tried that didn’t work: bootstrapping the semantic layer by having an LLM auto-generate metric definitions from raw tables and query logs. It produced plausible-looking definitions that encoded the very ambiguities we were trying to eliminate, and was net-negative on our evals versus a smaller, human-curated layer. Therefore we recommend generating the documentation with Claude, but having a human own the definition .
Lineage and the transformation graph: when the semantic layer doesn’t cover a question, lineage and table ranking (based on number of references) let the agent reason about which upstream models feed a concept, which are deprecated, and which share grain. This transforms “I don’t know the metric” into “I know which governed model to aggregate from.” It’s also the backbone of the freshness and provenance signals we surface in online validation below.
Query corpus: historical SQL from dashboards, notebooks, and prior analyses. Intuitively, this should be high-value: it’s a record of every question already answered correctly. In practice, we found that giving the agent raw retrieval access to thousands of prior queries moved accuracy by less than a point (we walk through that ablation in a later section below). Unstructured retrieval couldn’t map a new question to the right precedent. What does work is distilling that corpus into structured per-domain reference docs and reusable analysis patterns described in skills . Treat the query history as raw material for curation, not as a source of truth the agent reads directly.
Business context: the layer most teams skip, and the one we underrated the longest. An agent that doesn’t understand your business will answer what the user asked, but not what they meant. It won’t know that “the Q2 launch” refers to a specific product, that two teams define the same term differently, or that a question is being asked because a board meeting is on Thursday. We pipe in a company knowledge graph consisting of indexed docs, roadmaps, decision logs, and our organizational structure so the agent can resolve ambient references and ask better clarifying questions.
The common failure pattern across all four is the same one from the data foundations layer: poor or stale documentation . Claude is exceptionally useful for closing the gap (drafting column descriptions, proposing metric docs from query patterns, flagging undocumented models in CI), but the curation and ownership are managed by humans.
In the next two sections, we discuss how to mak

[truncated]
