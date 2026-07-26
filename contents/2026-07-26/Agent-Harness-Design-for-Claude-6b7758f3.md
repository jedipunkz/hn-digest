---
source: "https://claude.com/blog/harnessing-claudes-intelligence"
hn_url: "https://news.ycombinator.com/item?id=49059497"
title: "Agent Harness Design for Claude"
article_title: "Agent Harness Design: 3 Patterns for Harnessing Claude's Intelligence | Claude by Anthropic"
author: "ankitg12"
captured_at: "2026-07-26T16:48:18Z"
capture_tool: "hn-digest"
hn_id: 49059497
score: 1
comments: 0
posted_at: "2026-07-26T16:05:11Z"
tags:
  - hacker-news
  - translated
---

# Agent Harness Design for Claude

- HN: [49059497](https://news.ycombinator.com/item?id=49059497)
- Source: [claude.com](https://claude.com/blog/harnessing-claudes-intelligence)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T16:05:11Z

## Translation

タイトル: クロード用エージェント ハーネス デザイン
記事のタイトル: エージェント ハーネスの設計: クロードの知性を活用するための 3 つのパターン |クロード by Anthropic
説明: エージェント ハーネスの設計では、エージェントにどのような足場が必要で、何が不要かを決定します。クロードのように改良されたハーネスを構築するための Anthropic の 3 つのパターン。

記事本文:
エージェント ハーネスのデザイン: クロードの知性を活用するための 3 つのパターン |クロード by Anthropic
クロード製品のご紹介 クロード
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者に連絡する 営業担当者に連絡する 営業担当者に問い合わせる
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
Claude 上のプラットフォーム構築の概要
営業担当者に連絡する 営業担当者に連絡する 営業担当者に問い合わせる
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
エージェントハーネスの設計: クロードの知性を活用するための 3 つのパターン
エージェントハーネスの設計: クロードの知性を活用するための 3 つのパターン
インテリジェンス、レイテンシ、コストのバランスをとったアプリケーションを構築します。
共有 リンクをコピー https://claude.com/blog/harnessing-claudes-intelligence
Anthropic の共同創設者の 1 人、Chris Olah 氏は、Claude のような生成 AI システムは構築されるよりも成長することが多いと述べています。研究者は成長を促すための条件を設定しますが、出現する正確な構造や機能は常に予測できるわけではありません。
これにより、クロードを使用した構築に課題が生じます。エージェント ハーネスは、クロードが単独で実行できないことに関する仮定をエンコードしますが、クロードの能力が高まるにつれて、それらの仮定は古くなります。
エージェント ハーネスは、モデルの周囲にあるソフトウェアの足場です。つまり、生のインテリジェンスを動作するエージェントに変えるループ、ツール、コンテキスト管理、ガードレールです。エージェント ハーネスの設計は、その足場に何が属するのか、そしてモデルが改善されるにつれて何を取り出せるのかを決定する実践です。
この記事では、Claude の進化に歩調を合わせるアプリケーションを構築する際にチームが使用すべき 3 つのパターンを紹介します。

遅延とコストのバランスをとりながらテリジェンスを実現します。すでに知っていることを利用し、何をやめられるかを尋ね、エージェントのハーネスとの境界を慎重に設定します。
1. ハーネスではなくモデルに頼る: クロードが知っていることを使用する
クロードがよく理解しているツールを使用してアプリケーションを構築することをお勧めします。
2024 年後半、Claude 3.5 Sonnet は、ファイルの表示、作成、編集のための bash ツールとテキスト エディター ツールのみを使用して、SWE ベンチ検証済み (当時は最先端) で 49% に達しました。 Claude Code はこれらと同じツールに基づいています。 Bash はビルディング エージェント用に設計されたものではありませんが、クロードは使い方を知っており、時間の経過とともに使いこなせるようになるツールです。
私たちは、クロードがこれらの一般的なツールをさまざまな問題を解決するパターンに組み立てるのを見てきました。たとえば、エージェント スキル、プログラムによるツール呼び出し、メモリ ツールはすべて bash ツールとテキスト エディター ツールから構築されています。
2. エージェントのハーネスを剥ぎ取ります。何をやめられるかを尋ねます。
エージェント ハーネスは、クロードが単独ではできないことについての仮定をエンコードします。クロードがより有能になるにつれて、これらの仮定をテストする必要があります。
クロード自身の行動を調整させる
一般的な仮定は、すべてのツールの結果がクロードのコンテキスト ウィンドウを介して戻って、次のアクションを通知する必要があるということです。ツール結果のトークンを処理することは、次のツールに渡す必要があるだけの場合、またはクロードが出力の小さなスライスのみを気にする場合には、時間がかかり、コストがかかり、不必要になる可能性があります。
単一の列について推論するために大きなテーブルを読み取ることを検討してください。テーブル全体がコンテキスト内に配置され、クロードは必要のないすべての行に対してトークン コストを支払います。ハードコードされたフィルターを使用して、ツール設計でこれに取り組むことができます。しかし、これは、エージェント ハーネスがクロードの方が有利なオーケストレーションの決定を行っているという事実には対処していません。
クロードにコードを実行させる

n ツール ( bash ツールや言語固有の REPL など) はこれに対処します。これにより、Claude はツール呼び出しとツール呼び出し間のロジックを表現するコードを作成できます。すべてのツール呼び出し結果がトークンとして処理されることをハーネスが決定するのではなく、クロードは、コンテキスト ウィンドウに触れることなく、次の呼び出しにどの結果を通過、フィルター、またはパイプするかを決定します。コード実行の出力のみがクロードのコンテキスト ウィンドウに到達します。
オーケストレーションの決定は、ハーネスからモデルに移ります。コードはクロードがアクションを調整するための一般的な方法であるため、強力なコーディング モデルは強力な一般エージェントでもあります。 Claude は、このパターンを使用した非コーディング評価で優れたパフォーマンスを示しました。BrowseComp (エージェントの Web ブラウズ能力をテストするベンチマーク) では、Opus 4.6 に独自のツール出力をフィルタリングする機能が与えられ、精度が 45.3% から 61.6% に向上しました。
クロードに独自のコンテキストを管理させます
タスク固有のコンテキストによって、クロードは bash やテキスト エディター ツールなどの一般的なツールを使用するようになります。一般的な前提として、システム プロンプトはタスク固有の指示を使用して手動で作成する必要があるということです。問題は、指示を含むプロンプトのプリロードが多くのタスクにわたって拡張できないことです。トークンが追加されるたびにクロードの注意力の予算が枯渇し、めったに使用されない指示を含むコンテキストをプリロードするのは無駄です。
クロードにスキルにアクセスできるようにすることで、この問題に対処できます。各スキルの YAML フロントマターは、コンテキスト ウィンドウに事前に読み込まれた短い説明であり、スキルの内容の概要を提供します。完全なスキルは、タスクが必要な場合にクロードがファイル読み取りツールを呼び出すことによって徐々に公開できます。
スキルによってクロードは独自のコンテキスト ウィンドウを自由に組み立てることができますが、コンテキスト編集はその逆で、古くなったコンテキストや無関係になったコンテキスト (古いテキストなど) を選択的に削除する方法を提供します。

結果や思考ブロック。
サブエージェントを使用することで、Claude は、いつ新しいコンテキスト ウィンドウにフォークして、特定のタスクの作業を分離するかをよりよく理解できるようになりました。 Opus 4.6 では、サブエージェントを生成する機能により、BrowseComp での結果が単一エージェントの最良の実行より 2.8% 改善されました。
クロードに独自のコンテキストを永続化させる
長時間実行されるエージェントは、単一のコンテキスト ウィンドウの制限を超える可能性があります。一般的な仮定は、メモリ システムはモデルに関する検索インフラストラクチャに依存する必要があるということです。私たちの仕事の多くは、どのコンテンツを保持するかをクロード自身が選択できる簡単な方法を提供することに焦点を当ててきました。
たとえば、圧縮により、クロードは長期的なタスクの連続性を維持するために過去のコンテキストを要約できます。いくつかのリリースを経て、クロードは覚えておくべきことを選択するのが上手になりました。たとえば、エージェント検索タスクである BrowseComp では、指定した圧縮予算に関係なく、Sonnet 4.5 は 43% で横ばいのままでした。それでも、同じ設定で Opus 4.5 は 68% にスケールアップし、Opus 4.6 は 84% に達しました。
メモリ フォルダーは別のアプローチであり、クロードがファイルにコンテキストを書き込み、後で必要に応じてそれらを読み取ることができます。クロードがこれをエージェント検索に使用しているのを見てきました。 BrowseComp-Plus では、Sonnet 4.5 にメモリ フォルダーを与えると、精度が 60.4% から 67.2% に上昇しました。
ポケモンなどの長期にわたるゲームは、クロードのメモリ フォルダーの使用能力が向上した例です。 Sonnet 3.5 は記憶をトランスクリプトとして扱い、何が重要かではなくノンプレイヤー キャラクター (NPC) が言ったことを書き留めました。 14,000 歩を歩いた後でも、イモムシ ポケモンに関する 2 つのほぼ重複したファイルを含む 31 個のファイルがあり、まだ 2 番目の町にありました。
caterpie_weedle_info:
- キャタピーとビードルはどちらもイモムシのポケモンです。
- キャタピーは毒を持たないイモムシポケモンです。
- ビードルは毒を持つ芋虫のポケモンです。

- この情報は、将来の遭遇と戦闘にとって非常に重要です。
- ポケモンが毒にかかったら、ポケモンに癒しを求めるべきです
できるだけ早くセンターに。後のモデルは戦術メモを書きました。同じステップ数の Opus 4.6 には、ディレクトリに整理された 10 個のファイル、3 つのジムバッジ、および自身の失敗から抽出された学習ファイルが含まれていました。
/ゲームプレイ/ラーニング.md:
- Bellsprout Sleep+Wrap コンボ: Sleep 前に BITE で高速 KO
パウダーランド。セットアップさせないでください。
- 第 1 世代のバッグ制限: 最大 20 アイテム。不要なTMはダンジョンの前に捨ててください。
- スピンタイル迷路: 入り口の Y 位置が異なると、異なる結果が得られます
目的地。すべてのエントリを試し、複数のポケットを連鎖させます。
- B1F y= 16 壁 x= 9 -28 ですべて固体であることを確認 (ステップ 14557 ) 3. ハーネス設計で境界を慎重に設定します。
エージェント ハーネスは、UX、コスト、またはセキュリティを強化するためのクロード周りの構造を提供します。
キャッシュヒットを最大化するコンテキストを設計する
メッセージ API はステートレスです。クロードは前のターンの会話履歴を見ることができません。これは、エージェント ハーネスが、各ターンでのクロードへの過去のすべてのアクション、ツールの説明、指示とともに新しいコンテキストをパッケージ化する必要があることを意味します。
プロンプトは、設定されたブレークポイントに基づいてキャッシュできます。つまり、Claude API は、ブレークポイントまでのコンテキストをキャッシュに書き込み、コンテキストが以前のキャッシュ エントリと一致するかどうかを確認します。
キャッシュされたトークンのコストは基本入力トークンの 10% であるため、キャッシュ ヒットを最大化するのに役立つエージェント ハーネスの原則をいくつか紹介します。
UX、可観測性、またはセキュリティ境界のために宣言型ツールを使用する
クロードは、アプリケーションのセキュリティ境界や UX サーフェスを必ずしも知っているわけではありません。クロードはツール呼び出しを発行し、ハーネスによって処理されます。 bash ツールは、Claude にアクションを実行するための広範なプログラムの活用を提供しますが、ハーネスに提供するのはコマンド文字列のみです。

すべてのアクションで同じ形状。アクションを専用ツールにプロモートすると、ハーネスに、インターセプト、ゲート、レンダリング、または監査できる型付き引数を備えたアクション固有のフックが与えられます。
セキュリティ境界を必要とするアクションは、専用ツールの当然の候補です。多くの場合、可逆性は適切な基準であり、外部 API 呼び出しなどの元に戻すのが難しいアクションは、ユーザーの確認によって制限できます。 edit などの書き込みツールには、最後に読み取られてから変更されたファイルをクロードが上書きしないように、失効チェックを含めることができます。
ツールは、ユーザーにアクションを提示する必要がある場合にも役立ちます。たとえば、モーダルとしてレンダリングして、ユーザーに質問を明確に表示したり、ユーザーに複数のオプションを提供したり、ユーザーがフィードバックを提供するまでエージェント ループをブロックしたりできます。
最後に、ツールは可観測性のために役立ちます。アクションが型指定されたツールである場合、ハーネスはログ、トレース、および再生できる構造化された引数を取得します。
ツールへのアクションを促進するという決定は、継続的に再評価される必要があります。たとえば、Claude Code の自動モード (出版時はリサーチ モード) は、bash ツールの周囲にセキュリティ境界を提供します。つまり、2 番目のクロードがコマンド文字列を読み取って、それが安全かどうかを判断します。このパターンは専用ツールの必要性を制限する可能性があるため、ユーザーが一般的な方向性を信頼するタスクにのみ使用する必要があります。専用ツールは、特定のリスクの高いアクションに引き続き使用できます。
エージェント ハーネス設計の未来
クロードの知性の最前線は常に変化しています。クロードができないことについての仮定は、その能力が段階的に変化するたびに再テストする必要があります。
このパターンが繰り返されることがわかります。長期的なタスク用に構築したエージェントでは、Sonnet 4.5 はコンテキストの制限が近づいていることを感知して途中で終了していました。コンテキスト ウィンドウをクリアするためのリセットを追加しました

この「文脈不安」に対処するために。 Opus 4.5 では、そのような動作はなくなりました。それを補うために私たちが構築したコンテキスト リセットは、エージェント ハーネスの中で重荷になっていました。
クロードのパフォーマンスのボトルネックになる可能性があるため、このデッドウェイトを取り除くことが重要です。時間の経過とともに、アプリケーションの構造や境界は、「何をやめればよいか?」という疑問に基づいて切り詰められる必要があります。
ここで説明したすべてのツールとパターンを使用するには、claude-api スキルを確認してください。
クロード プラットフォーム チームの技術スタッフのメンバーであるランス マーティンによって書かれました。取り上げられたトピックについて有益な議論をしてくださった Thariq Shihipar 氏、Barry Zhang 氏、Mike Lambert 氏、David Hershey 氏、Daliang Li 氏に心より感謝いたします。編集レビューとフィードバックを提供してくださった Lydia Hallie、Lexi Ross、Katelyn Lesse、Andy Schumeister、Rebecca Hiscott、Jake Eaton、Pedram Navid、Molly Vorwerck に感謝します。
前へ 前へ 0 / 5 次へ 次へ
クロードとともに構築するチーム向けの製品ニュースとベスト プラクティスをさらに詳しくご覧ください。
クロード モデルの説明: ユースケースに最適なモデルの選択
Enterprise AI クロード モデルの説明: ユース ケースに最適なモデルの選択 クロード モデルの説明: ユース ケースに最適なモデルの選択 クロード モデルの説明: ユース ケースに最適なモデルの選択 クロード モデルの説明: choo

[切り捨てられた]

## Original Extract

Agent harness design is deciding what scaffolding your agent needs—and what it doesn't. Three patterns from Anthropic for building harnesses that improve as Claude does.

Agent Harness Design: 3 Patterns for Harnessing Claude's Intelligence | Claude by Anthropic
Meet Claude Products Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Agent Harness Design: 3 Patterns for Harnessing Claude's Intelligence
Agent Harness Design: 3 Patterns for Harnessing Claude's Intelligence
Building applications that balance intelligence, latency, and cost.
Share Copy link https://claude.com/blog/harnessing-claudes-intelligence
One of Anthropic’s co-founders, Chris Olah, says that generative AI systems like Claude are grown more than they are built. Researchers set the conditions to direct growth, but the exact structure or capabilities that emerge aren’t always predictable.
This creates a challenge for building with Claude: agent harnesses encode assumptions about what Claude can’t do on its own, but those assumptions grow stale as Claude gets more capable.
An agent harness is the software scaffolding around a model: the loop, tools, context management, and guardrails that turn raw intelligence into a working agent. Agent harness design is the practice of deciding what belongs in that scaffolding and, as models improve, what you can take out.
In this article, we share three patterns that teams should use when building applications that keep pace with Claude’s evolving intelligence while balancing latency and cost: use what it already knows, ask what you can stop doing, and carefully set boundaries with the agent harness.
1. Lean on the model, not the harness: use what Claude knows
We suggest building applications using tools that Claude understands well.
In late 2024, Claude 3.5 Sonnet reached 49% on SWE-bench Verified—then state of the art —with only a bash tool and a text editor tool for viewing, creating, and editing files. Claude Code is grounded in these same tools. Bash wasn’t designed for building agents, but it's a tool that Claude knows how to use and gets better at using over time.
We've seen Claude compose these general tools into patterns that solve different problems. For instance, Agent Skills , programmatic tool calling , and the memory tool are all built from the bash and text editor tools.
2. Strip your agent harness down: ask what you can stop doing
Agent harnesses encode assumptions about what Claude can’t do on its own. As Claude gets more capable, those assumptions should be tested.
Let Claude orchestrate its own actions
A common assumption is that every tool result should flow back through Claude’s context window to inform the next action. Processing tool results in tokens can be slow, costly, and unnecessary if it only needs to be passed to the next tool or if Claude only cares about a small slice of the output.
Consider reading a large table to reason about a single column: the whole table lands in context and Claude pays the token cost for every row it doesn't need. It’s possible to tackle this in tool design, using hard-coded filters . But this does not address the fact that the agent harness is making an orchestration decision that Claude is better positioned to make.
Giving Claude a code execution tool (e.g., bash tool or language-specific REPL ) addresses this: it allows Claude to write code to express tool calls and the logic between them. Rather than the harness deciding that every tool call result is processed as tokens, Claude decides what results to pass through, filter, or pipe into the next call without touching the context window. Only the output of code execution reaches Claude’s context window.
The orchestration decision moves from the harness to the model. Since code is a general way for Claude to orchestrate actions, a strong coding model is also a strong general agent. Claude shows strong performance on non-coding evals using this pattern: on BrowseComp, a benchmark that tests the ability of agents to browse the web, giving Opus 4.6 the ability to filter its own tool outputs brought accuracy from 45.3% to 61.6%.
Let Claude manage its own context
Task-specific context steers Claude’s use of general tools like bash and the text editor tool. A common assumption is that system prompts should be hand-crafted with task-specific instructions. The problem is that pre-loading prompts with instructions does not scale across many tasks: every token added depletes Claude’s attention budget and it is wasteful to pre-load context with rarely used instructions.
Giving Claude the ability to access skills addresses this: the YAML frontmatter of each skill is a short description pre-loaded into the context window, providing an overview of the skill contents. The full skill can be progressively disclosed by Claude calling a read file tool if a task calls for it.
While skills give Claude the freedom to assemble its own context window, context editing is the inverse, providing a way to selectively remove context that’s become stale or irrelevant, such as old tool results or thinking blocks.
With subagents , Claude is getting better at knowing when to fork into a fresh context window to isolate work on a specific task. With Opus 4.6 , the ability to spawn subagents improved results on BrowseComp by 2.8% over the best single-agent runs.
Let Claude persist its own context
Long-running agents can exceed the limit of a single context window . A common assumption is that memory systems should rely on retrieval infrastructure around the model. Much of our work has focused on giving Claude simple ways to choose for itself what content to persist.
For example, compaction lets Claude summarize its past context in order to maintain continuity on long-horizon tasks. Over several releases, Claude has gotten better at choosing what to remember. On BrowseComp , for example, an agentic search task, Sonnet 4.5 stayed flat at 43% regardless of the compaction budget we gave it. Yet Opus 4.5 scaled to 68% and Opus 4.6 reached 84% with the same setup.
A memory folder is another approach, allowing Claude to write context to files and later read them as needed. We’ve seen Claude use this for agentic search. On BrowseComp-Plus, giving Sonnet 4.5 a memory folder lifted accuracy from 60.4% to 67.2% .
Long-horizon games , such as Pokémon, are an example of Claude’s improved ability to use a memory folder. Sonnet 3.5 treated memory as a transcript, writing down what non-player characters (NPCs) said rather than what mattered. After 14,000 steps it had 31 files—including two near-duplicates about caterpillar Pokémon—and was still in the second town:
caterpie_weedle_info:
- Caterpie and Weedle are both caterpillar Pokémon.
- Caterpie is a caterpillar Pokémon that does not have poison.
- Weedle is a caterpillar Pokémon that does have poison.
- This information is crucial for future encounters and battles.
- If our Pokémon get poisoned, we should seek healing at a Pokémon
Center as soon as possible. Later models wrote tactical notes. Opus 4.6, at the same step count, had 10 files organized into directories, three gym badges, and a learnings file distilled from its own failures:
/gameplay/learnings.md:
- Bellsprout Sleep+Wrap combo: KO FAST with BITE before Sleep
Powder lands. Don't let it set up!
- Gen 1 Bag Limit: 20 items max. Toss unneeded TMs before dungeons.
- Spin tile mazes: Different entry y-positions lead to DIFFERENT
destinations. Try ALL entries and chain through multiple pockets.
- B1F y= 16 wall CONFIRMED SOLID at ALL x= 9 -28 (step 14557 ) 3. Set boundaries carefully in your harness design
Agent harnesses provide structure around Claude to enforce UX, cost, or security.
Design context to maximize cache hits
The Messages API is stateless. Claude cannot see the conversation history of prior turns. This means that the agent harness needs to package new context alongside all past actions, tool descriptions, and instructions for Claude at each turn.
Prompts can be cached based on set breakpoints . In other words, the Claude API writes context up until a breakpoint to the cache and checks whether the context matches any prior cache entries.
Since cached tokens are 10% the cost of base input tokens, here are a few principles in the agent harness help maximize cache hits:
Use declarative tools for UX, observability, or security boundaries
Claude doesn't necessarily know an application's security boundary or UX surface. Claude emits tool calls, which are handled by the harness. A bash tool gives Claude broad programmatic leverage to perform actions, but it gives the harness only a command string—the same shape for every action. Promoting actions to dedicated tools gives the harness an action-specific hook with typed arguments it can intercept, gate, render, or audit.
Actions that require a security boundary are natural candidates for dedicated tools. Reversibility is often a good criterion, and hard-to-reverse actions such as external API calls can be gated by user confirmation. Write tools like edit can include a staleness check so Claude doesn't overwrite a file that changed since it was last read.
Tools are also useful when an action needs to be presented to a user. For example, they can be rendered as a modal to display a question clearly to the user, give the user multiple options, or block the agent loop until a user provides feedback.
Finally, tools are useful for observability. When the action is a typed tool, the harness gets structured arguments it can log, trace, and replay.
The decision to promote actions to tools should be continually re-evaluated. For example, Claude Code's auto-mode (in research mode at the time of publication) provides a security boundary around the bash tool: it has a second Claude read the command string and judge whether it's safe. This pattern can limit the need for dedicated tools, and should only be used for tasks where users trust the general direction. Dedicated tools can still earn their place for certain high-stakes actions.
The future of agent harness design
The frontier of Claude’s intelligence is always changing. Assumptions about what Claude can’t do need to be re-tested with each step change in its capability.
We see this pattern repeat itself. In an agent we built for long-horizon tasks , Sonnet 4.5 would wrap up prematurely as it sensed the context limit approaching. We added resets to clear the context window in order to address this "context anxiety." With Opus 4.5, the behavior was gone. The context resets we built to compensate had become dead weight in the agent harness.
Removing this dead weight is important because it can bottleneck Claude’s performance. Over time, the structure or boundaries in our applications should be pruned based the question: what can I stop doing?
To use all tools and patterns discussed here, check out our claude-api skill .
Written by Lance Martin, member of technical staff on the Claude Platform team. Special thanks to Thariq Shihipar, Barry Zhang, Mike Lambert, David Hershey, and Daliang Li for helpful discussion on the topics covered. Thanks to Lydia Hallie, Lexi Ross, Katelyn Lesse, Andy Schumeister, Rebecca Hiscott, Jake Eaton, Pedram Navid, and Molly Vorwerck for their editorial review and feedback.
Prev Prev 0 / 5 Next Next eBook
Explore more product news and best practices for teams building with Claude.
Claude models explained: choosing the best model for your use case
Enterprise AI Claude models explained: choosing the best model for your use case Claude models explained: choosing the best model for your use case Claude models explained: choosing the best model for your use case Claude models explained: choo

[truncated]
