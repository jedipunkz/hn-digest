---
source: "https://shearer.org/research/addressing-biggest-problems-in-ai/"
hn_url: "https://news.ycombinator.com/item?id=48677239"
title: "The biggest problems in using AI"
article_title: "The biggest problems in using AI | Dan Shearer"
author: "smartmic"
captured_at: "2026-06-25T18:43:03Z"
capture_tool: "hn-digest"
hn_id: 48677239
score: 1
comments: 0
posted_at: "2026-06-25T18:15:05Z"
tags:
  - hacker-news
  - translated
---

# The biggest problems in using AI

- HN: [48677239](https://news.ycombinator.com/item?id=48677239)
- Source: [shearer.org](https://shearer.org/research/addressing-biggest-problems-in-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T18:15:05Z

## Translation

タイトル: AI利用における最大の問題
記事タイトル: AI 利用における最大の問題 |ダン・シアラー
説明: 2026 年に数十億人が使用する AI には多くの問題があり、社会のあらゆるレベルで際限なく議論されています。 2025 年末から、私は倫理と信頼性という特定の問題と、大手 AI 企業すべてが採用しているアプローチがなぜ十分ではないのかに興味を持つようになりました。予測
[切り捨てられた]

記事本文:
ダン・シアラー
テーマ オリジナル ミニマリスト インダストリアル サーマル アースバウンド サイケデリック ノルド ハイ コントラスト ソラライズド ノート
AI活用における最大の問題点
AI利用の最大の問題点 幻覚とナンセンス
AIは記憶喪失で自己認識が欠如している
Agentic AI は権限を与えません
2026 年に数十億人が使用する AI には多くの問題があり、社会のあらゆるレベルで議論が絶えません。 2025年末から私はこうなりました
彼らは、倫理と信頼性という特定の問題と、大手 AI 企業のすべてが採用しているアプローチがなぜ十分ではないのかに興味を持っています。
予測可能性、または彼らが言うところの「調整」は、このタイプの AI に期待できるものではありません。
同僚がこれらの企業とはまったく異なるアプローチに取り組み始め、2026 年 2 月から私はプロトタイプ バージョンに貢献して使用しています。
人工組織の ↗
コンセプト。この記事では、人工組織が有望な新しい組織であると私が考える理由を説明します。
方向。英国がここで説明しているように、マルチエージェント Agentic AI は非常に重要です
政府 ↗
, しかし、それがうまくいくことはほとんどありません。自分で試してみたい場合は、コアリサーチを使用できます
コード ↗
、私が毎日そうしているように。
AI の使用における最大の問題 #
Perseverance 合成エンジン ↗
(PCE) は人工組織を使用して、これらの差し迫った AI の問題を解決します。 PCE は LLM の動作を改善しようとするのではなく、避けられない誤動作を検出して修正するように設計されています。コンピューター サイエンスとは関係なく、私はこれらのアイデアをイアン M バンクスの小説やマス エフェクト ビデオ ゲームで見つけました。
。
PCE は、LLM エージェントにタスクを割り当てることで機能します。LLM エージェントには、慎重に実行すべき役割が割り当てられています。エージェントは、タスクが仕様どおりに完了するか失敗するまで相互に反復します。

正直に「私にはこれはできません、その仕事は私には不可能です」と言うのです。これまでのところ、この仕組みは、自信に満ちた虚偽の主張、幻覚、危険なアドバイスなどの一般的な問題を検出して修正するのに効果的であると思われます。 PCE では、誰も AI を信頼する必要がなく、構造だけを信頼する必要があります。この構造は、何世紀にもわたって試行されてきた構造を厳密にモデル化しているため、ほとんどの人が認識できます。他の組織と同様に、人工組織には職務の分離、独立したチェック機能、そして見るべきものだけを見ることができるエージェントが存在します。かなりうまくいきます。
この設計は、通常のトレーニングや指導では完全に修正できない 3 つの障害モード、つまり幻覚、コンテキストの問題、記憶の問題に対処します。
言語モデルは確率に従ってテキストを生成し、データベースから事実を取得するのではなく、パターンに基づいて次のテキスト (「トークン」) が選択されます。モデルに、選択する関連性の高いテキスト (「コンテキスト」) のプールがない場合でも、そのようにプログラムされているため、とにかく確率的にテキストを生成します。
その結果、モデルが虚偽または誤解を招く主張をしながらも自信に満ちているように聞こえる作話が生じます。 AI が自己表現を上手にすればするほど、
これらの幻覚はより説得力のあるものになる可能性があります。
研究は結論を出し続けています ↗
トレーニングによって幻覚は解消されない、そして新しい調査では ↗
彼らは、幻覚を「[モデルの] アーキテクチャに固有の基本的な数学的必然性」である可能性があると説明しています。 AI企業はより良い指導とトレーニングを提供することでこの問題を解決しようとしているが、幻覚が本当に避けられないのであれば、これは決して信頼できるものではない。 AI をより信頼できるものにするためには、アーキテクチャを変える必要があると私は確信しています。
モデルへのコンテキスト入力は、

'前'。質の高い事前情報は、利用可能な最良のドキュメント、以前の関連する決定、密接な背景で構成され、AI はそこからはるかに優れた出力を生成します。人間の組織と同じように、人工組織も意思決定を向上させるために最高品質の入力ドキュメントを処理し、慎重にラベルを付けたり、推測を拒否したりするよう努めています。これは、幻覚に対処できる最初の構造的な方法です。
2 番目の手法もよく知られています。他の人に作業をチェックしてもらいます。 PCE には Corroborator と呼ばれるエージェントがおり、その唯一の仕事は Composer エージェントが書いたものを読み、ソース文書に対してすべての主張を検証することです。 Corroborator はその直前にソースを持っているため、Composer がクレームを発明した場合、Corroborator はそれがサポートされていないことを認識します。裏付け者は、もっともらしい作話には動じません。なぜなら、指示されている場合はインターネット上の参考文献も含め、手元にある情報源から証明できるものだけを受け入れるように指示されているからです。
特定の AI の会話またはセッション内で、AI のコンテキスト ウィンドウ ↗
一度に考慮できる情報の最大量であり、数文字のトークンで測定されます。 2026 年に利用可能になる、100 万を超えるトークン コンテキスト ウィンドウを備えた最先端の AI モデル (Llama 4 Scout、Kimi 3.0、GPT-5.4、Claude Opus 4.6 など) は、おそらく 1000 ページのテキストを一度に保持でき、これは現代の携帯電話のストレージ容量の約 * 100 万分の 1 です。 AI はそのテキストに対して、電話でできることよりもはるかに多くのことを行っていますが、この合計制限が、現時点で AI の信頼性が低い主な原因となっています。さらに悪いことに、モデルのコンテキスト ウィンドウが最大値近くまで埋まると、モデルの気が散りやすくなり、エラーが発生しやすくなります。コンテキスト ウィンドウがいっぱいになると、情報が失われ、明白なつながりが失われ、関連する内容を組み込むことができなくなります。

なぜなら、彼らは自分たちの出力に適切な事前情報を構築するために必要なものを見つけることができないからです。 AI がすでに知っていることのより良いインデックスを持っていれば、AI の出力の信頼性はさらに高まるでしょう。
これは検索の問題であり、情報が図書館やデータベースに保存されている場合、人間の組織は情報を見つけるのが得意であることがわかっています。対照的に、モデルにコンテキストを最初から再構築するよう依頼することは、モデルに推測を依頼することになります (そして推測することになります)。 PCE は永続的なナレッジ ベースでこの問題を解決します。ドキュメントは永続的に保存され、バージョン管理され、全文検索用にインデックスが付けられます。新しいタスクが到着すると、システムはモデルに最初から開始するように要求するのではなく、以前の作業を取得します。キュレーターはこの施設の記憶を担当するエージェントであり、あらゆるものをファイル化し、インデックスを付け、見つけられるようにします。モデルは、このデータベースから取得された豊富な正しいコンテキストに基づいて動作します。このようなデータベースは新しいアイデアではありません ↗
、しかし、私たちのそれの使用法は、物理的な組織の他のよく知られた概念ときちんと調和しており、図書館部門から他の部門、または一般の人々などに渡される情報のよく知られた動作をすべて継承しています。
AI は記憶喪失であり、自己認識を欠いています #
AI は、AI との会話の間に、トレーニング後に起こったことをすべて忘れてしまいます。 AI は、以前の作業に基づいて構築したり、すでに行われた決定を追跡したり、組織の歴史にわたって蓄積されたエラーを修正したりすることはできません。単なるコンピューターとして考えられる AI は 1980 年代のようなもので、電源を入れた後、AI が何を知っているか、何ができるかを指示する必要があります。 AI についても同様で、リクエストを入力する前にすべてのチャットボットで同じことが起こります。チャットボットが先週話し合った内容を覚えているようであれば、多くの作業が完了したことを意味します

舞台裏で保存された情報を AI コンテキスト ウィンドウにロードして戻すことで、シームレスにチャットを続けることができます。この起動作業の一部は、危険なことを言わないよう AI チャットボットに指示することに専念します。マルチエージェント エージェント AI では、これは各エージェントがそのタスクに関連する指示とデータを含む初期コンテキストで起動する必要があることを意味します。
さらに微妙なことに、エージェントは自分自身の状態、つまりどのようなツールが利用可能か、現在は何時か、コンテキストが最近いっぱいになって要約にまとめる必要があるかどうかを認識できません。その認識がなければ、エージェントは自分が知っていることと知らないことを効果的に推論することができません。
人工組織の硬直性は、各エージェントを起動し、指示し、監視するコンピューター コードです。
組織はさらに別のエージェントです。このコンピューターコードは、人間の組織がどのように行動するかを定義する規則や法律に似ています。
には CEO がいますが、法制度が CEO を満足させることのない規則の施行を重視しているため、CEO ができることには依然として制限があります。
Artificial Organization の永続的な知識ベースは、AI の会話間の記憶喪失というこの問題にも対処します。以前の決定、下書き、推論、検証の結果は保存され、セッション間およびエージェント間で検索可能です。各構成タスクは、将来の作業に基づいて構築できる構造化されたトレースを残すため、システムは何が決定されたかを認識し、自身の状態を推論できます。最近の一部のエージェント システム (Claude Code、Goose) では、会話全体にメモリが追加され始めていますが、違いは、PCE のメモリが構造化され、厳選され、検索可能であることです。つまり、記憶された内容の袋ではなく、意味構造を備えた組織の記録です。

事実。
Agentic AI は権限を与えません #
AI の問題に対する主流の対応策はエージェント AI です。エージェント AI は複数の AI エージェントを接続し、多くの場合ツール、Web 検索、またはドキュメント検索にアクセスします。これは、チャットボットの会話のような単一のモデル呼び出しよりも改善されています。
既存のエージェントのアプローチの制限は、エージェントがコンテキストを自由に共有するため、あるエージェントの偏見やエラーが次のエージェントに伝播することです。ソースを取得する同じエージェントは、クレームを書き込み、自身の出力を評価します。作品が独立してチェックされるという構造上の保証はないため、合成ステップでの製造はそのまま通過します。協力がデフォルトです。敵対的レビューはそうではありません。
PCE は、構造が第一、機能が二番目という異なるアプローチを採用しています。
構成アーキテクチャ #
Perseverance 合成エンジン ↗
この組織ロジックをコードで実装します。病院内のすべての従業員が医療記録を利用できることが期待されるのと同様に、コンテキストは自由には共有されません。標準の PCE ワークフローには 5 人のエージェントが関与します。タスクは、要件に応じて個々のエージェントに委任することも、完全な PCE パイプラインを通じて委任することもできます。
作曲家はソース資料からテキストを草案します。文書を読み、一貫した散文を書きますが、評価基準は見ていません。レポートは書くが、試験問題を設定しない著者と同じです。
確証者は、情報源と草稿の両方を読んで独立して事実確認を行います。その唯一の任務は、すべての主張を検証し、読者に届く前に捏造を発見することです。 Composer と Corroborator は両方ともソース文書を参照しますが、Corroborator はソースがサポートしているもののみを受け入れるように指示されます。作曲家が何かを発明した場合、確証者はそのギャップに気づきます。
批評家は品質と安全性を評価します。

草案と評価ルーブリックは読みますが、出典は読みません。これは逆向きに聞こえるかもしれませんが、査読者は情報源を見ると、情報源が主張を裏付けているように見えるため主張が正しいと思い込み、怠惰な評価をする傾向があります。学術的な査読でおなじみのブラインドレビューは、批評家に実際にページに記載されている内容を評価することを強います。
検閲官は、対象となる視聴者に対する適切性をチェックします。文書は、事実としては正しく、十分に議論されていても、受信者にとっては文脈上間違っている可能性があります。かつて私たちは、正確でよく書かれた求職書を持っていましたが、その特定の雇用主にとって不適切な民間研究プログラムについて言及していました。裏付け者はそれを通過させ（真実）、批評家はそれを通過させ（十分な議論）、そして検閲官はコンテンツと視聴者の不一致を捕らえました。
キュレーターは、機関の記憶を公開および維持し、ファイリング、インデックス付けを行い、将来のエージェントやユーザーが作品を発見できるようにします。キュレーターは組織の図書館員であり、何が行われたか、誰が決定したか、そしてなぜ行われたかについての構造化された記録を維持します。
全体像を把握できるエージェントは一人もいないため、問題を合理的に解決できるエージェントは一人もいません。 Composer は、出典を示して出典のない主張を弁解することはできません (出典を確認していません)。批評家は「情報源はこうしなければならない」とは言えません。

[切り捨てられた]

## Original Extract

There are many problems with the AI billions of people use in 2026, discussed endlessly at all levels of society. From the end of 2025 I became interested in the particular problems of ethics and reliability, and why the approaches taken by all of the large AI companies are not good enough. Predicta
[truncated]

Dan Shearer
Theme Original Minimalist Industrial Thermal Earthbound Psychedelic Nord High Contrast Solarized Notes
The biggest problems in using AI
The Biggest Problems in Using AI Hallucination and nonsense
AI is amnesiac and lacks self-awareness
Agentic AI doesn’t do permissions
There are many problems with the AI billions of people use in 2026, discussed endlessly at all levels of society. From the end of 2025 I became
interested in the particular problems of ethics and reliability, and why the approaches taken by all of the large AI companies are not good enough.
Predictability, or ‘alignment’ as they call it, is just not something we can expect from this type of AI.
A colleague started working on a very different approach from these companies, and from February 2026 I have been contributing to and using prototype versions
of the Artificial Organisations ↗
concept. This article explains why I believe Artificial Organisations are a promising new
direction. Multi-agent Agentic AI is pretty important, as described here by the UK
government ↗
, but it is rarely done well. If you want to try for yourself, you can use the core research
code ↗
, as I do daily.
The Biggest Problems in Using AI #
The Perseverance Composition Engine ↗
(PCE) uses Artificial Organisations to solve these pressing AI problems. PCE does not try to make LLMs behave better, but is designed instead so that their inevitable misbehaviour is detected and corrected. And regardless of the computer science, I found these ideas in Iain M Banks’ novels and the Mass Effect video game
.
PCE works by assigning a task to LLM agents who each have a carefully enforced role to play. The agents iterate between each other until either the task is completed to specifications, or it fails by honestly saying “I can’t do this, the task is impossible for me.” So far, this arrangement seems effective at detecting and correcting common problems such as confident false assertions, hallucinations, or dangerous advice. With PCE, nobody needs to trust an AI, only the structure. The structure is recognisable by most people, since it is closely modelled on ones tried and tested for centuries. Like any organisation, Artificial Organisations have separation of duties, independent checks, and agents who can only see what they need to see. It works rather well.
This design addresses three failure modes that the usual training and instruction cannot fully fix: hallucination, context issues, and memory issues.
Language models generate text according to probability, where the next piece of text (a ’token’) is selected based on patterns, not by retrieving facts from a database. If a model does not have a pool of highly relevant text to select from (the ‘context’) it will probabilistically generate text anyway because that it what it is programmed to do.
The result is confabulation, where the model sounds confident while making a false or misleading claim. The better the AIs become at expressing themselves, the
more convincing these hallucinations can become.
Research keeps concluding ↗
that training does not eliminate hallucination, and newer surveys ↗
describe hallucinations as potentially “fundamental mathematical inevitabilities inherent to [the model’s] architecture.” The AI companies are trying to solve this by giving better instruction and training, but if hallucination is indeed inevitable then this will never be reliable. I am persuaded the architecture needs to change for AI to become more trustworthy.
Context input to a model is called the ‘prior’. A quality prior comprises the best available documents, previous relevant decisions, germane background, and from it AI generates much better output. Just like a human organisation, Artificial Organisations strive to deal with the best quality input documents in order to improve decisionmaking, and to carefully label or even reject guesswork. This is the first structural way we can tackle hallucinations.
A second technique is also familiar: have someone else check the work. PCE has an agent called the Corroborator whose only job is to read what the Composer agent wrote, and to verify every claim against the source documents. The Corroborator has the sources right in front of it, so if the Composer invented a claim, the Corroborator will see it is unsupported. Corroborator is unmoved by plausible confabulation, because it is instructed to only accept what can be proven from the sources to hand, including references on the internet if it has been instructed to do so.
Within any given AI conversation or session, the AI’s Context Window ↗
is the largest amount of information it can consider at once, measured in tokens of a few characters. The most advanced AI models available in 2026 with million+ token context windows (eg Llama 4 Scout, Kimi 3.0, GPT-5.4, Claude Opus 4.6) can hold maybe 1000 pages of text at once, or about * one millionth the storage capacity of a modern phone. An AI is doing a lot more with that text than a phone can possibly do, but this total limit is currently a major cause of why AI is unreliable. Even worse, models become more distractable and error-prone as their context window fills up close to its maximum. When their context window is full they lose information, miss obvious connections, and fail to incorporate relevant material because they can’t find what they need to build a good prior on their output. If AIs had a better index of the things they already know, then their output would be more reliable.
This is a retrieval problem and we know human organisations are good at finding information if we put it is kept in a library or a database. In contrast, asking a model to reconstruct context from scratch is asking it to guess (and it will). PCE solves this with a persistent knowledge base: documents are stored permanently, version-controlled, and indexed for full-text search. When a new task arrives, the system retrieves prior work rather than asking the model to start from scratch. The Curator is the agent responsible for this institutional memory, and it files, indexes, and makes everything findable. The model operates from plentiful, correct context retrieved from this database. Such a database is not a new idea ↗
, but our use of it fits neatly alongside other familiar concepts from physical organisations, inheriting all the well-known behaviours of information that passes from the library department to other departments, or to members of the public etc.
AI is amnesiac and lacks self-awareness #
Between AI conversations, an AI has forgotten everything that has happened since it was trained. An AI can’t build on prior work, track decisions already made, or correct accumulated errors across an organisation’s history. Considered as merely a computer, AI is like something from the 1980s, where after power-on you need to instruct it what it knows and what it can do. Similarly with an AI, and that is what happens with every chatbot before you can even type in a request to it. If the chatbot appears to remember what you were discussing last week then that means a lot of work has been done behind the scenes to load the stored information back into the AI context window so you can continue chatting seamlessly. Part of this boot-up work is dedicated to instructing the AI chatbot to be more reliable, not to say dangerous things etc. In multi-agent agentic AI, that means each agent must be booted up with its initial context containing instructions and data relevant to its task.
More subtly, agents lack awareness of their own state — what tools are available, what time it is, whether the context recently filled up and needed to be compacted down into a summary. Without that awareness an agent can’t reason effectively about what it knows and doesn’t know.
The rigidity of Artificial Organisations is computer code that starts up, instructs and monitors each agent, even though the controlling brain of the
organisation is yet another agent. This computer code is akin to the rules and laws that define how we expect human organisations to behave, for example, an organisation
has a CEO, but there are still limits on what the CEO can do because the legal system cares about enforcing the rules not keeping CEOs happy.
The persistent knowledgebase of Artificial Organisations also addresses this problem of amnesia between AI conversations. Prior decisions, drafts, reasoning, and verification results are preserved and searchable across sessions and between agents. Each composition task leaves a structured trace that future work can build on, so the system knows what has been decided and can reason about its own state. Some recent agentic systems (Claude Code, Goose) have begun adding memory across conversations, but the difference is that PCE’s memory is structured, curated, and searchable — an institutional record with semantic structure, not a bag of remembered facts.
Agentic AI doesn’t do permissions #
The mainstream response to AI’s problems is agentic AI, connecting multiple AI agents together, often with access to tools, web search, or document retrieval. This is an improvement over a single model call like a chatbot conversation.
The limitation of existing agentic approaches is that agents share context freely, so the biases and errors of one agent propagate to the next. The same agent that retrieves sources also writes claims and evaluates its own output. There is no structural guarantee that the work is independently checked, so fabrication in the synthesis step passes straight through. Cooperation is the default; adversarial review is not.
PCE takes a different approach: structure first, capability second.
The composition architecture #
The Perseverance Composition Engine ↗
implements this organisational logic in code. Context is not shared freely any more than we expect medical records to be available to all workers in a hospital. The standard PCE workflow involves five agents. Tasks can be delegated to individual agents or through the full PCE pipeline depending on requirements.
The Composer drafts text from source materials. It reads the documents and writes coherent prose, but does not see the evaluation criteria — like an author who writes the report but doesn’t set the exam questions.
The Corroborator fact-checks independently, reading both the sources and the draft. Its single task is to verify every claim and catch fabrication before it reaches the reader. The Composer and Corroborator both see the source documents, but the Corroborator is instructed only to accept what the sources support. If the Composer invented something, the Corroborator will see the gap.
The Critic evaluates quality and safety. It reads the draft and the evaluation rubrics, but not the sources. This sounds backwards, but when reviewers see the sources they tend toward lazy evaluation, assuming claims are right because the sources seem to support them. Blind review, familiar from academic peer review, forces the Critic to evaluate what’s actually on the page.
The Censor checks appropriateness for the intended audience. A document can be factually correct and well-argued but contextually wrong for its recipient. We once had a job application letter that was accurate and well-written, but mentioned a private research programme that was inappropriate for that particular employer. The Corroborator passed it (true), the Critic passed it (well-argued), and the Censor caught the mismatch between content and audience.
The Curator publishes and maintains institutional memory, filing, indexing, and making work discoverable to future agents and users. The Curator is the librarian of the organisation, maintaining the structured record of what has been done, who decided it, and why.
No single agent can see the complete picture, so no single agent can rationalise away problems. The Composer can’t excuse unsourced claims by pointing to the sources (it doesn’t see them). The Critic can’t say “the sources must

[truncated]
