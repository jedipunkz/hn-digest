---
source: "https://poyo.co/note/20260525T094605/"
hn_url: "https://news.ycombinator.com/item?id=48354521"
title: "Two LLM UI Patterns That Aren't Chat"
article_title: "Two LLM UI Patterns That Aren't Chat"
author: "minikomi"
captured_at: "2026-06-02T00:43:13Z"
capture_tool: "hn-digest"
hn_id: 48354521
score: 5
comments: 0
posted_at: "2026-06-01T09:29:56Z"
tags:
  - hacker-news
  - translated
---

# Two LLM UI Patterns That Aren't Chat

- HN: [48354521](https://news.ycombinator.com/item?id=48354521)
- Source: [poyo.co](https://poyo.co/note/20260525T094605/)
- Score: 5
- Comments: 0
- Posted: 2026-06-01T09:29:56Z

## Translation

タイトル: チャットではない 2 つの LLM UI パターン
説明: イントロ チャットは依然としてデフォルトの LLM インターフェイスであり、ほとんどの場合はそれで問題ありません。エージェント ハーネスは依然として、単一の直線的な会話を中心に構築されています...

記事本文:
チャットではない 2 つの LLM UI パターン
▾
チャットではない 2 つの LLM UI パターン
チャットは依然としてデフォルトの LLM インターフェイスであり、ほとんどの場合はそれで問題ありません。エージェント ハーネスは依然として、単一の直線的な会話を中心に構築されています。一部の LLM タスクは、メッセージとしてよりも構造化コンテキストとして表現した方が適切です。この投稿では、テーブルとしての比較とツリーとしての探索の 2 つのパターンを見ていきます。
ここでは、私が調べた内容と、 exe.dev のshelleyを使用してまとめたライブのサンプルアプリをいくつか紹介します。
LLM にチャットで比較するよう依頼すると、すぐに面倒になってしまいます。最初は適切な表が得られますが、フォローアップの質問をすると、有用な情報が複数の回答に分割されたり、LLM が毎回表を再描画しようとしたりします。別のアイテムを追加すると、再生の問題がさらに悪化します。
項目が行、質問が列、回答がセルになった表がすでに完成しています。
新しい質問によって新しい列が作成され、項目を追加すると新しい行が作成される、ハイブリッド チャット/テーブル インターフェイスを検討しました。実際にはスプレッドシートを使ってチャットしているような感じです。
トピックを入力すると、アプリは関連するアイテムを検索し、そのページを取得して、テーブルの最初のバージョンに入力します。 「超軽量の 1 人用テント」の場合、次のようになります。
+---------------------+----------+----------+---------------------+
|アイテム |重量 |価格 |壁タイプ |
+---------------------+----------+----------+---------------------+
| Zpacks Plex ソロ | 405g | $599 |単層DCF |
|ビッグ アグネス フライ クリーク | 879g | $350 |二重壁 |
|タープテント プロトレイル Li | 425g | $399 |単層DCF |
|ニモ ホーネット エリート | 765g | $450 |二重壁OSMO |
+---------------------+----------+----------+---------------------+
アプリは現在、アイテムを収集するためにKagi検索を使用していますが、同じアプローチですw

ショッピング サイト、マーケットプレイス、社内製品データベース、採用ツール、または行がすでに存在するその他の場所に埋め込んでもうまく機能します。
最初の列も最初の検索結果から生成されます。アプリは取得したアイテムを見て、いくつかの有用なディメンションを自動的に選択するため、ユーザーが何かを尋ねる前に、テーブル上にすでに何かが用意されています。
テーブルが存在したら、質問を入力すると、テーブルに新しい列が追加されます。各質問は既存の構造に 1 つの比較ディメンションを追加し、結果は特定の着地点を持ちます。
モデル呼び出しは通常のものです。これらの項目の概要が与えられた場合、行ごとにこの質問に答えます。何が変わるかというと、その結果がすでに構築している構造に反映されるということです。
特に各行の応答を返すツール呼び出しとして LLM 呼び出しを作成することも、現在の世代のモデルによく適合します。
ユーザー プロンプトは大まかに次のようになります。
比較表の列に入力しています。
質問:「自立式ですか?」
ビッグ・アグネス・フライ・クリーク: [概要]
tarptent-protrail-li: [概要]
提供された概要を使用してください。情報が入手できない場合は「不明」と回答してください。
{
"名前": "fill_column",
"description": "表内の各項目についての比較の質問に答えてください。",
"入力スキーマ": {
"タイプ": "オブジェクト",
"プロパティ": {
「答え」: {
"タイプ": "配列",
「アイテム」: {
"タイプ": "オブジェクト",
"プロパティ": {
"row_id": { "タイプ": "文字列" },
"値": { "タイプ": "文字列" }
}、
"必須": ["行ID", "値"]
}
}
}、
"必須": ["回答"]
}
}
重要な制約は単純です。行 ID ごとに 1 つのセル値を返すということです。モデルはテーブルの特定の部分を占めます。
モデルがセルを埋めるため、列は構造化された仕様を超える可能性があります。
翻訳には当然の副作用が伴います。商品ページが日本語で、表を英語で表示したい場合、

セルには別の翻訳モードを使用せずに英語が入力されます。
単位の正規化も同様です。重さ（g）と聞くと？ソース ページに十分な情報があると仮定すると、モデルは通常、すべてを同じユニットに取り込むことができます。
ソフトジャッジメントコールも予想以上にうまく機能します。 「初心者でも大丈夫ですか？」は、唯一の正解はなく、実際の購入決定のほとんどは同様ですが、比較するのに役立つ質問です。
より新しいコンテキストが必要な質問については、アプリは列に入力する前に再検索できます。 「レディットはこれについてどう思いますか？」ライブ検索に基づいた列になります。
製品比較のケースは明白ですが、共有の一連の質問に対して一連のものを評価する場合には、同じ対話パターンがどこにでも当てはまります。
一部のタスクでは、コンテキストのさまざまな分岐を探索する必要があります。直線的なトピックとして始まったものが、それぞれが独自の集中的な探索を必要とするサブトピックを生成することがよくあり、単一のチャット インターフェイスではこれらが混ざり合ってコンテキストが汚染され、1 つのスレッドについて深く掘り下げることが困難になります。
この種の作業の自然な形状はツリーです。各ブランチは親コンテキストを継承しますが、兄弟からは独立しています。これは、私が Emacs で gptel を使用するときに以前に使用したものです。オプションを使用すると、組織モードの見出し内の会話を制限できます。
各ノードをモデルによって拡張できる小さなアウトライナーとタスク プランナーを検討しました。各ブランチには独自の焦点を当てた拡張が行われるため、他のスレッドを邪魔することなく 1 つのスレッドを深く掘り下げることができます。
各ノードには展開ボタンがあり、現在の見出しと周囲のコンテキストをモデルに送信し、モデルは 3 つの形状のいずれかで応答します。
ブレークダウンは、子として挿入されるいくつかの具体的なサブ項目です。
質問は、詳細が欠けていると状況が変わることを意味します。

答えてください、それでモデルはそれを尋ねます。 「小規模な微調整モデル用のモデル ホスティングの選択」の場合、モデルはパラメーター サイズと予想されるリクエスト量について質問し、提案される選択肢またはフリーテキスト フォールバックを示します。
オプション応答は、相互に排他的なパスをカバーします。 1 つを選択すると、現在のノード ラベルが置き換えられますが、元のゴールはコンテキストのために親チェーンに残ります。
Breakdowner は、利用可能な 1 つのツール search を使用して、1 つのユーザー メッセージをモデルに送信します。プロンプトでは、次の 3 つの形状のいずれかの JSON オブジェクトを 1 つだけ要求されます。
{"タイプ": "内訳"、"アイテム": ["..."]}
{"タイプ": "質問", "質問": "...", "選択肢": ["..."]}
{"タイプ": "オプション", "オプション": ["..."]}
違いは、並列具体的な次の動きの内訳、情報が欠落していると答えが変わる場合の質問、および選択したノードがノードを置き換える相互排他的なパスのオプションです。
このプロンプトは、一般的な計画言語に強く反対します。不可: 「オプションの調査」、「環境のセットアップ」、「計画の作成」、「テストと反復」、「評価」。モデルに与えられた肯定的な例:
Reddit で X について苦情を言った 10 人を見つける
3 店舗での価格と配送での価格を比較する
あいまいな子ノードは、次の展開時にさらに多くの作業を作成します。具体的なものについては、アクション、検索、委任、またはさらに細分化することができます。
モデルは、現在の見出し、親チェーン、およびブランチ内の以前の質問ノードから収集された以前の回答を確認します。展開されているノードの周囲のフォーカスされたコンテキストのみが表示されます。
コンテキスト:
- 親
- 子供
→ 現在の見出し
以前の回答:
...
ハンドル: 現在の見出し
親チェーンは、トップレベルのゴールに接続されたディープノードを維持します。事前の回答は、予算が厳しいこと、これは東京で行う必要があること、またはユーザーがこれまでそのことを行ったことがないことをモデルが忘れていないことを意味します。
兄弟ブランチはコンフィギュレーションを共有しません

文章。ツリー構造はモデルが何を認識するかを決定し、最終的には線形トランスクリプトよりも明確なスコープ設定になります。
これらのツールは両方とも小さく、内部でのモデルの作業は普通です。実際に重要なのはコンテキスト、つまりコンテキストがどのように構築され、どのように使用されるか、そして UI がそれをサポートするように形作られているかどうかです。
チャットはデフォルトのインターフェイスとして確立されており、それは問題ありません。しかし、ほとんどの LLM ツールはコンテキストを結果論として扱います。コンテキストは一方向に拡大し、UI はコンテキストを形成するためにほとんど何もせず、ほとんどがユーザーから隠されています。彼らは入力して答えを受け取りますが、モデルが実際に見たものは決して表面化することはありません。
モデル呼び出しを中心とした構造と、特定の種類のコンテキストを自然で視覚的に感じさせる UI には、興味深い未開発の設計作業のほとんどが存在します。

## Original Extract

Intro Chat is still the default LLM interface, and for most cases that&#39;s fine. Agentic harnesses are still built around a single linear conversation at...

Two LLM UI Patterns That Aren't Chat
▾
Two LLM UI Patterns That Aren't Chat
Chat is still the default LLM interface, and for most cases that's fine. Agentic harnesses are still built around a single linear conversation at their core. Some LLM tasks are better represented as structured context than as messages. This post looks at two patterns: comparison as a table, and exploration as a tree.
I've included my explorations here, along with some live example apps I put together using shelley on exe.dev .
Asking an LLM to compare things in chat quickly becomes tedious. At first you get a decent table, but as you ask follow-up questions, the useful information gets split across multiple answers, or the LLM tries to redraw the table every time. Adding another item makes the regeneration problem worse.
The table is already the thing you want, with items as rows, questions as columns, and answers as cells.
I explored a hybrid chat/table interface where new questions create new columns and added items create new rows. In practice it feels like chatting with a spreadsheet.
After entering a topic, the app searches for relevant items, fetches their pages, and fills the first version of the table. For "ultralight 1-person tents" that might look like:
+----------------------+----------+-------+-------------------+
| Item | Weight | Price | Wall Type |
+----------------------+----------+-------+-------------------+
| Zpacks Plex Solo | 405 g | $599 | Single-wall DCF |
| Big Agnes Fly Creek | 879 g | $350 | Double-wall |
| Tarptent ProTrail Li | 425 g | $399 | Single-wall DCF |
| Nemo Hornet Elite | 765 g | $450 | Double-wall OSMO |
+----------------------+----------+-------+-------------------+
The app currently uses Kagi search to gather items, but the same approach would work well embedded in a shopping site, marketplace, internal product database, recruiting tool, or anywhere else the rows already exist.
The initial columns are also generated from the first search results. The app looks at the retrieved items and picks a few useful dimensions automatically, so there is already something on the table before the user asks anything.
Once the table exists, typing a question adds a new column to it. Each question adds one comparison dimension to the existing structure, and the result has a specific place to land.
The model call is ordinary: given these item summaries, answer this question for each row. What changes is that the result lands in the structure you are already building.
Creating the LLM call as a tool call which specifically returns an answer for each row also fits the current generation of models well.
The user prompt looks roughly like:
You are filling in a column in a comparison table.
Question: "is it freestanding?"
big-agnes-fly-creek: [summary]
tarptent-protrail-li: [summary]
Use the provided summaries. Answer "Unknown" if the information is not available.
{
"name": "fill_column",
"description": "Answer a comparison question for each item in the table.",
"input_schema": {
"type": "object",
"properties": {
"answers": {
"type": "array",
"items": {
"type": "object",
"properties": {
"row_id": { "type": "string" },
"value": { "type": "string" }
},
"required": ["row_id", "value"]
}
}
},
"required": ["answers"]
}
}
The important constraint is simple: return exactly one cell value for each row id. The model fills a specific part of the table.
Because a model is filling the cells, the columns can go beyond structured specs.
Translation is a natural side effect. If an item's page is in Japanese and I want the table in English, the cell gets filled in English without a separate translation mode.
Unit normalization is similar. If I ask Weight (g)? the model can usually get everything into the same unit, assuming the source pages have enough information.
Soft judgement calls also work better than I expected. "Is this good for a beginner?" is a useful comparison question even though it has no single correct answer, and most real buying decisions are similar.
For questions that need fresher context, the app can search again before filling the column. "What does Reddit think about this?" becomes a column backed by live search.
The product comparison case is the obvious one, but the same interaction pattern could apply anywhere you are evaluating a set of things against a shared set of questions.
Some tasks require exploring many different branches of context. What begins as a linear topic often spawns sub-topics that each warrant their own focused exploration, and in a single chat interface these bleed together and poison the context, making it hard to go deep on any one thread.
The natural shape for this kind of work is a tree, where each branch inherits its parent context but is independent of its siblings. It's one I've used before when working with gptel in Emacs . An option allows you to limit the conversation within an org-mode heading.
I explored a small outliner and task planner where each node can be expanded by the model. Each branch gets its own focused expansion, so you can go deep on one thread without disturbing the others.
Each node has an expand button that sends the current heading and some surrounding context to the model, which replies in one of three shapes.
A breakdown is a handful of concrete sub-items, inserted as children.
A question means one missing detail would change the answer, so the model asks for it. For "Choose model hosting for small finetuned model", the model asks about parameter size and expected request volume, with suggested choices or a free-text fallback.
An options response covers mutually exclusive paths. Picking one replaces the current node label while the original goal stays in the parent chain for context.
Breakdowner sends one user message to the model with a single available tool: search . The prompt asks for exactly one JSON object in one of three shapes:
{"type": "breakdown", "items": ["..."]}
{"type": "question", "question": "...", "choices": ["..."]}
{"type": "options", "options": ["..."]}
The distinction is: breakdown for parallel concrete next moves, question when missing info would change the answer, and options for mutually exclusive paths where picking one replaces the node.
The prompt pushes hard against generic planning language. NOT: "Research options", "Set up environment", "Create plan", "Test and iterate", "Evaluate". Positive examples given to the model:
Find 10 people who complained about X on Reddit
Compare prices at 3 stores vs delivery
A vague child node creates more work at the next expansion; a concrete one can be acted on, searched, delegated, or broken down further.
The model sees the current heading, the parent chain, and any prior answers collected from question nodes earlier in the branch: only the focused context around the node being expanded.
Context:
- parent
- child
→ current heading
Prior answers:
...
Handle: current heading
The parent chain keeps a deep node connected to the top-level goal. Prior answers mean the model does not forget that the budget is tight, or that this needs to happen in Tokyo, or that the user has never done the thing before.
Sibling branches don't share context. The tree structure determines what the model sees, which ends up being a cleaner scoping than a linear transcript can provide.
Both of these tools are small and the model work inside them is ordinary. What they are really about is context: how it gets built, how it gets used, and whether the UI is shaped to support that.
Chat has established itself as the default interface, and that is fine. But most LLM tooling treats context as an afterthought: a transcript that grows in one direction, with the UI doing very little to shape it and mostly hidden from the user. They type, they get an answer, and what the model actually saw is never surfaced.
The structure around the model call, and the UI that makes a particular kind of context feel natural and visible, is where most of the interesting and unexplored design work lives.
