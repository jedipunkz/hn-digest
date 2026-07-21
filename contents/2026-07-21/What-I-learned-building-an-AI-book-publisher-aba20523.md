---
source: "https://edwardbenson.com/2026/07/ai-book-publishing-lessons"
hn_url: "https://news.ycombinator.com/item?id=48993790"
title: "What I learned building an AI book publisher"
article_title: "What I learned building an AI book publisher"
author: "eob"
captured_at: "2026-07-21T16:17:15Z"
capture_tool: "hn-digest"
hn_id: 48993790
score: 1
comments: 0
posted_at: "2026-07-21T15:40:12Z"
tags:
  - hacker-news
  - translated
---

# What I learned building an AI book publisher

- HN: [48993790](https://news.ycombinator.com/item?id=48993790)
- Source: [edwardbenson.com](https://edwardbenson.com/2026/07/ai-book-publishing-lessons)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T15:40:12Z

## Translation

タイトル: AI 書籍出版社の構築で学んだこと

記事本文:
AI 書籍出版社の構築で学んだこと > Ted Benson / Writing Tinkering About AI 書籍出版社の構築で学んだこと
ここ数か月間、私は印刷品質のクックブックを作成するというサイド プロジェクトをハッキングしてきました。このプロジェクトの開始時に、私は自分自身に 3 つの厳しい要件を設定しました。
生成入力は本の表紙に収まる程度のものである必要があります
生成出力はだらしないものであってはならず、印刷を正当化するのに十分なものでなければなりません。
生成はエンドツーエンドで自動化する必要がある
これまでのところ、500 回以上の試行を経て 53 冊のクックブックが作成され、数億のトークンが消費されました。
Web 上および PDF および EPUB のダウンロードとしてすべて無料です。
私の意見では、これらの本はまだ印刷する価値がありません。
しかし、レシピとレシピのキュレーションは非常に優れており、私が「スロップなし」の主張を擁護したくなるほどで​​す。
私は現在太平洋上の飛行機に乗っていますが、ロボットの作者にチェックインすることができず、正直に言って少しお腹が空いています。そこで、これまでに学んだことについていくつかメモするのに最適な機会だと考えました。
重要な場合に備えて、この投稿は AI による執筆に関するものですが、AI を使用せずに執筆および編集されています。
私は料理が大好きですが、ちょっとズルいと思って料理本を作ることにしました。
クックブックは、開発サイクルのこの時点で LLM が自律的に適切に生成できる、長い間開発されてきたコンテンツの一種に限りなく近いものです。
まず、料理本はエピソード形式です。 LLM は、長文テキストの生成において人間らしいペースと一貫性を維持するのに苦労しています。しかし、クックブックでは追跡すべきブック間の依存関係がほとんどないため、これらの問題は最小限に抑えられます。ゲーム・オブ・スローンズのフィナーレと比較して、ホームコメディのエピソードを生成するようなものです。章間およびレシピ間の依存関係は、ほとんどの場合、配布上関連性があり、過度の使用を避けます。

本全体を通して同じ比喩が何度も繰り返されます。
Second Cookbook は主にキュレーションと翻訳の行為であり、LLM は優れたキュレーターおよび翻訳者です。次の素晴らしい小説を書くには、何層もの創造的なひらめきとバランスが必要ですが、人間の料理本でさえ、本質的には既存の文化的知識を集めて、特定のシェフの料理に対する見方に翻訳したものです。チリですがビーガンです。ピザですが、ペストが付いています。フォーですが、アメリカの食材を使用しています。
したがって、私にとってクックブックは、LLM 用に作成できる最も簡単な長文の評価に似ています。
ロボット ブック クラブは、さまざまなシリーズの書籍で構成されています。
各書籍シリーズは、書籍のタイトル、裏表紙の説明、著者の声に関するいくつかのメモなど、いくつかの簡単な入力に基づいて完成した書籍を作成するプログラムです。
現在、3 つの料理本シリーズがあります。
ディアスポラ シリーズ 。アメリカの二世が祖父母から学ぶレシピが含まれています。 （例：日本、キューバ、スペイン）
有名レストランのスタイルを家庭料理にアレンジしたアットホームシリーズ。 (例: ユダヤ系デリ、中華料理のテイクアウト ..)
ダイエット シリーズ : 特定の食事を維持する人向けのレシピが含まれています (例: 30 種類の簡単な食事)
私は、最初の書籍シリーズ プログラムである「ディアスポラ シリーズ」を、多くの試行錯誤を経て手書きで書きました。そのシリーズからエージェント スキルを抽出し、それ以降のすべてのシリーズ プログラムはそのスキルを使用して AI によって作成されました。
出版物は生成されてから直接ウェブサイトに公開されますが、私はその後、書籍の PDF をランダムに選択してざっと目を通します。これにより、変更リクエストのリストが作成され、エージェントにフィードバックして、シリーズ プログラムのスキルと既存のプログラムの両方を更新します。 「版数」に基づいて、各本の再版がどのように行われたかがわかります。

」をウェブサイトに掲載しています。
台湾語は第 13 版にありますが、私が手書きしたため実際の数字はおそらく 200 程度です。一方、ベトナム語などの他の書籍の一部はプログラムの最新バージョンを使用して生成されているため、第 3 版にのみ掲載されています。
各書籍シリーズのプログラムは、エージェント自体ではなく、エージェントのワークフローです。私は単にエージェントの未来に抵抗しているだけかもしれませんが、少なくともロボット ブック クラブのようなものについては、トップレベルのコントロール プレーンを LLM に譲るべきだという考えにはまだ納得していません。各シリーズには非常に具体的なリサーチ、執筆、編集、植字、出版のプロセスがあり、そこから逸脱したくありません。そのプロセスには非常に多くのエージェントが含まれていますが、プロセス全体自体は厳格なままであるべきだと私は思います。
トップレベルの剛性を持つことにより、より制御されたエラー処理、エラー回復、ロギング、および可観測性のセットアップもサポートされます。これは、出力量のスケールアップに非常に役立ちます。
ワークフローは出力品質のスケールアップに役立ちます
ワークフローは品質を向上させる最大の要因です。 Anthropic のブランド化されたワークフロー機能を意味するのではなく、それが表す一般化されたアイデア、つまり LLM 処理の一般的なパターンをカプセル化し、単純な呼び出しでそれらを繰り返し使用することを意味します。 LLM スタイルの関数カプセル化に少し似ています。
私が持っているそのようなワークフローの 1 つは画像生成を処理します。これは、多くの地方の食べ物がトレーニング データで正確に表現されていないため、これが困難でした。
この画像の「モデル 1」は、画像生成が単なるプロンプトである場合の出力を示しています。 「モデル 2」は、現在の食品画像ワークフローで実行されているものと同じ入力です。
私が導入した別のワークフローは、一連の編集パスを使用して、章の導入全体で同じ比喩の繰り返しを減らすことを目的としています。

すでに書かれたものを検査します。
特定の分野にどのようなワークフローが役立つかを理解するには時間がかかりますが、それらを持っているので、エージェントに新しいシリーズ プログラムの作成を手伝ってもらうときに名前で参照することができます。次のように言うことができます。「章の画像には必ず地方の食べ物と盛り付けられた食べ物のワークフローを使用してください。ただし、表紙には雰囲気のあるワークフローを使用し、章のタイトルには料理人のワークフローを使用する必要があります。」
結果として得られるプログラムのバージョン 1 は、以前のシリーズ プログラムのバージョン 30 に到達することで獲得したトリックの恩恵をすでに受けています。
着実に進歩するにはキャッシュが不可欠です
各本の生成は完了するまでに何時間もかかります。生成の最終段階でエラーが発生して処理全体が破壊されたら、何も進められないことに早い段階で気づきました。 (信じてください。PDF 内のコンテンツをレイアウトしようとしたことがあるなら、ましてや AI にレイアウトしてもらうと、処理実行の最後にエラーが発生するでしょう。)
キャッシュを使用しない場合、これがサイド プロジェクトの場合、反復は 1 日あたり 1 回です。これを仕事として行う場合は、おそらく 1 日に 3 回の繰り返しになります。キャッシュを使用すると、2 回目にブックを生成するのに約 10 分かかり、一度に実行できる反復回数が大幅に変わります。
唯一の難点は、日々のキャッシュの無効化を避けるために、前もってデータを共同生成するのではなく、新しいデータを追加することを好む開発スタイルを採用することになるということです。ひどいことではありませんが、ブック シリーズのプログラムは、データ作成の追加形式のログのように見えます。
出力の安定性をエディション間で管理するのは難しい
文章では「最愛の人を殺せ」と言います。自分が気に入っている文章は喜んで削除してください。しかし、それは編集のプロには耐えられません。

セス。完全に自動化された書き込みにも同様の格言が必要になるかもしれません。「メニューの変更を許可する」のようなものです。
各書籍の生成には何千もの自律的な決定が含まれており、1 つのプロンプトやワークフローを変更すると、下流に多大な影響が及ぶ可能性があります。著者の声のトーンを微調整することで、料理本の章全体が追加されたり、削除されたりするのを見てきました。
私は書籍の版間の一貫性をある程度制御するために多くの戦略をハッキングしてきましたが、共有する価値があると思われる唯一の高レベルの戦略は、キャッシュ キーを賢く扱うことです。デフォルトの LLM キャッシュは、LLM への入力でキー設定される必要がありますが、実際のキャッシュ キーが変更された場合でも保持されるように、ブック全体または 1 つのブック内で特定の世代を固定したい場合があります。
質の高いヒルクライムには、優れたサンドボックスが不可欠です
永続的なバージョン管理されたサンドボックスは、出力をレビューし、書籍の品質を向上させるために重要です。
各ブックの生成は、ログ、LLM 入出力スナップショット、およびバージョン管理されたファイルシステムを含むカスタム サンドボックス リグ内で行われます。
つまり、特定のエージェントが起動される直前または直後に巻き戻して、ファイル システムの状態を検査し、実行された LLM 呼び出しとその出力を正確に確認できるということです。
生成ワークフロー全体では、1 冊あたり数千のノードがあり、これを合計すると、すでに 4 分の 1 テラバイト近くのログ データに達します。しかし、何が問題だったのかという疑問に対する事後的な答えを引き出すための非常に貴重な方法となっています。
生成プロセスの適切な視覚的プロットを生成できませんでした。デバッグに役立つ適切なビジュアライゼーションを見つけることは理論的には可能だと思われますが、おそらくそれ自体がサイド プロジェクトに過ぎません。
さらに役に立ったのは、最終版の本のどのような欠陥に気づいたかをメモすることです。

出力し、それらのメモを関連書籍のサンドボックス フォルダーへの参照とともにコーディング エージェントに提供します。多くの場合、数分後に、適切なデバッグ スレッドが実行されます。
自己改善ループは再帰的である必要はない
RSI が人気の頭字語であることは知っていますが、自己改善ループを再帰的に行う必要はありません。
私はバグ修正フローを、世代ごとに自動的に実行されるいくつかのスキルに変えました。
出力批評スキルは世代ごとに自動的に実行され、パフォーマンス、回復力、出力品質に関連する問題を特定し、その修正を実装できる場所を提案します。
バグ作成スキルは、出力された批評を未解決のバグ チケットの重複排除されたフォルダーにまとめます。
バグトリアージスキルは、私が取り組むべき上位 5 つのバグが何かを示唆します。
ほんの数か月前まで、私はログを手作業で読み、見つけた内容に基づいて更新を手作業でコーディングしていました。しかし、この改善ループを生成自体に組み込むことで、開発プロセスの約 95% が自動化されました。私は時々、トップ 5 の改善提案を読んで、追加のコンテキストや制約を追加して、エージェントにそれらの作業を開始してもらいます。
あなたはすぐに AI が生成したフィクションを読んで、気に入るでしょう。
フロンティアモデルは、次の素晴らしい小説を生み出すのにすでに十分に優れていると思います。必要なのは、生成プロセス全体にわたってバランス感覚と品質を維持するための適切なハーネスだけです。
LLM は、プロンプトを出してリッピングさせるだけでは、依然として長文コンテンツが苦手です。しかし、Robot Book Club に取り組んでいると、プロンプトを適切に囲むハーネスを使用すれば、出力を少なくとも 1 桁向上させることができると確信しています。おそらく2桁も大きいでしょう。
優れた作家、多作の読書家、そして優れた人物だと思います

ささやく人なら、今日のモデルを使って約 1 年間献身的にフルタイムで努力すれば、実行可能な小説執筆のハーネスを構築できるでしょう。すべての種類の本を生成できるわけではありませんが、「新興テクノロジー スリラー」や「成人向けロマンス」などの特定のカテゴリを生成できます。
..そして、読んでみたいと思います。
LinkedIn ロボット ブック クラブ 🤖 カヤ 🥥 · © 2026 Edward Benson.無断転載を禁じます。

## Original Extract

What I learned building an AI book publisher > Ted Benson / Writing Tinkering About What I learned building an AI book publisher
For the past few months, I’ve been hacking on a side project to generate print-quality cookbooks . At the start of this project, I set three hard requirements for myself:
Generation Input should be no more than what would fit on a book cover
Generation Output shouldn’t be slop — it should be good enough to justify printing.
Generation must be end-to-end automated
So far, that’s produced 53 cookbooks over more than 500 attempts consuming hundreds of millions of tokens.
All free on the web and as PDF & EPUB downloads.
The books aren't yet worth printing, in my opinion.
But the recipes, and recipe curation, are a pretty darn good, to the point I’d be wiling to defend the “no slop” claim.
I’m currently on a plane over the Pacific, unable to check in with my robot authors, and honestly a bit hungry. So I figured what better time to jot down a few notes about what I’ve learned so far.
In case it’s important to you: while this post is about AI writing, it's written and edited without AI.
I love to cook, but I chose to generate cookbooks because I think it's a bit of a cheat.
Cookbooks are as close as you can get to a type of long-from content LLMs can generate well, autonomously, at this point in their development cycle.
First, Cookbooks are episodic. LLMs struggle to maintain humanlike pace and cohesion across long-form text generation. But cookbooks minimize these issues because there are few cross-book dependencies to keep track of. It’s like generating a sitcom episode compared to the Game of Thrones finale. Cross-chapter and cross-recipe dependencies are mostly relevant distributionally — avoiding overuse of the same trope over and over throughout the book.
Second Cookbooks are mostly an act of curation and translation , and LLMs are great curators and translators. Writing the next great novel requires layers of creative spark and balance, But even human cookbooks are essentially a curation of existing cultural knowledge, translated into a particular chef’s take on the cuisine. Chili, but vegan. Pizza, but with pesto. Pho, but with American ingredients.
So cookbooks, to me, are akin to the easiest long-form writing eval that you could create for an LLM.
Robot Book Club is organized into different series of books .
Each book series is a program that produces a finished book based on a few simple inputs: the book title, the back-cover description, and a few notes on author voice.
Right now there are three cookbook series:
The Diaspora Series , which contains recipes second generation Americans might learn from their grandparents. (Examples: Japan, Cuba, Spain)
The At Home Series , which adapts famous restaurant styles into home cooking. (Examples: Jewish Deli, Chinese Takeout ..)
The Diet Series , which contains recipes for people who maintain specific diets (Examples: Whole 30 Quick Meals)
I hand wrote the very first book series program — The Diaspora Series - through a lot of trial and error. From that series, I extracted an agent skill, and every series program since has been AI-authored through the use of that skill.
While publication goes straight from generation onto website, I do skim through a random selection of book PDFs afterward. This produces a list of change requests, which I feed back into my agent to update both the series program skill and also the existing programs. You can see how regenerations of each book have happened based on the “Edition Number” on the website.
Taiwanese is on Edition 13 — though the real number is probably something like 200 since I hand-wrote that one — while some of the other books like Vietnamese are only on Edition 3, since they were generated using the most recent version of the program.
Each book series program is a workflow of agents rather than an agent itself. I might just be resisting our agentic futures, but I’m not yet sold on the idea that the top-level control plane should be yielded to an LLM, at least for something like Robot Book Club. Each series has a very specific research, writing, editing, typesetting, and publishing process that I don't want to veer from. That process includes many, many agents within it, but the overall process itself should, I think, remain rigid.
Having that top-level rigidity also supports a more controlled error handling, error recovery, logging, and observability setup — something that’s really helped scale the output volume.
Workflows help you scale output quality
Workflows have been the biggest factor in scaling quality. I don’t mean Anthropic’s branded Workflow feature, but the generalized idea it represents: encapsulating common patterns of LLM processing and then using them over and over again via a simple invocation. A bit like functional encapsulation, LLM style.
One such workflow I have handles image generation — something that’s been challenging since many regional foods aren’t exactly represented well in training data.
"Model 1" in that image is what output looks like if image generation is merely a prompt. "Model 2" is the same input being run through the current food image workflow.
Another workflow I’ve got aims to reduce repeating the same tropes across chapter introductions using a series of editorial passes that inspect what’s already been written.
It takes a while to figure out what specific workflows are helpful for a particular domain, but you have them, you can refer to them by name when asking your agent to help write a new series program. You can say something like: “make sure chapter images use the regional food and plated food workflows, but the cover should use the atmospheric workflow, and chapter titles should use the cooking chef workflow”.
Version 1 of the resulting program will already benefit from the tricks won by making it to Version 30 of your prior series program.
Caching is essential for steady progress
Each book generation takes many hours to complete, and early on I realized that I wouldn’t make any progress if an error in the final stages of generation destroyed the whole processing run. (And believe me: if you’ve ever tried to layout content in a PDF, let alone have an AI do it for you, you will hit errors at the end of your processing runs.)
Without caching, you’ve got a single iteration per day if this is your side project. And maybe three iterations per day if you’re doing this as your job. With caching, generating a book the second time around takes about ten minutes, significantly changing how many iterations you can go through in a single sitting.
The only catch is that you’ll find yourself adopting a development style that favors appending new data rather than co-generating it up front — to avoid day-over-day cache invalidations. It's not terrible.. but the book series programs end up looking a bit like an append-style log of data creation.
Output stability is hard to manage across editions
In writing they say “kill your darlings” — be willing to delete writing that you love, but that doesn’t survive the editing process. I think we may need a similar maxim for fully-automated writing.. something like “allow menu to change.”
Each book generation involves thousands of autonomous decisions, and changing any one prompt or workflow can cascade into enormous downstream effects. I’ve seen entire cookbook chapters added and destroyed from minor adjustments to the author’s tone of voice.
While I’ve hacked a bunch of strategies to gain some control over consistency across book editions, the only high level one that feels worth sharing is to get clever with your cache keys. Your default LLM cache should be keyed on the inputs to the LLM, but occasionally you want to pin a particular generation across books, or just within a single book, so that it persists even if the true cache key has changed.
Good sandboxing is essential to quality hill-climbing
Persistent versioned sandboxes are critical for reviewing output and hill-climbing book quality.
Each book generation happens within a custom sandbox rig that contains logs, LLM input/output snapshots, and a versioned filesystem.
That means you can rewind to exactly before or after a particular agent fired, inspect the state of the filesystem, and then review exactly what LLM calls got made and what their outputs were.
The whole generation workflow is thousands of nodes per book, and that's added up to almost a quarter terrabyte of logging data already. But it's been an invaluable way to extract post-mortem answers to questions about what went wrong.
I've failed to generate good visual plots of the generation process. Finding a good visualization to help debug is something that feels theoretically possible but perhaps it's a side project all by itself.
What's been more helpful is to take notes about what flaw I'm seeing in the final book output, and then provide those notes to a coding agent along with a reference to the sandbox folders of the related books. A few minutes later, some reasonable debugging threads to pull on often result.
Self-improvement loops don’t have to be recursive
I know RSI is a hot acronym, but self-improvement loops don’t have to be recursive to be awesome.
I've turned my bug-fixing flow into an handful of skills that run automatically after every generation:
An output critique skill runs automatically after every generation, identifying issues related to performance, resilience, and output quality, and then suggesting where a fix for them might be implemented.
A bug writing skill curates output critiques into a de-duplicated folder of open bug tickets.
A bug triage skill suggests what the top five bugs I should work on are.
Just a few months ago, I was hand-reading logs and hand-coding updates based on what I found. But incorporating this improvement loop into the generation itself has automated about 95% of my development process. I just jump in every once in a while to read the top five improvement suggestions, and then I add additional context or constraints and let an agent begin working on them.
You will soon read AI generated fiction.. and like it
I think frontier models are already good enough to generate the next great novel. They just need the right harness to maintain a sense of balance and quality across the generation process.
LLMs remain poor at long-form content if all you do is prompt them and let them rip. But working on Robot Book Club makes me confident that you can get at least an order of magnitude output improvement with the right harness surrounding your prompts. Maybe even two orders of magnitude.
I think a good writer, a prolific reader, and an agent whisperer, could build a viable fiction-writing harness in about a year of dedicated, full-time effort with today’s models. It wouldn’t be able to generate every kind of book, but a specific category: like “Emerging Technology Thriller” or “Coming of Age Romance”.
..And I think you’d want to read it.
LinkedIn Robot Book Club 🤖 Kaya 🥥 · © 2026 Edward Benson. All rights reserved.
