---
source: "https://hugotunius.se/2025/10/26/claude-teach-me-something.html"
hn_url: "https://news.ycombinator.com/item?id=48428097"
title: "Claude, Teach Me Something"
article_title: "Claude, Teach Me Something"
author: "dannyboland"
captured_at: "2026-06-06T21:27:25Z"
capture_tool: "hn-digest"
hn_id: 48428097
score: 2
comments: 1
posted_at: "2026-06-06T19:25:25Z"
tags:
  - hacker-news
  - translated
---

# Claude, Teach Me Something

- HN: [48428097](https://news.ycombinator.com/item?id=48428097)
- Source: [hugotunius.se](https://hugotunius.se/2025/10/26/claude-teach-me-something.html)
- Score: 2
- Comments: 1
- Posted: 2026-06-06T19:25:25Z

## Translation

タイトル: クロード、何か教えて
説明: ソクラテス式メソッドを使った破滅的なスクロールの代わりに、クロードを活用して学習します。

記事本文:
クロード、何か教えて
ホーム
私はドゥームスクロールの代替として新しいクロードワークフローを実験してきました。これは、LLM が最も得意とする非決定性とテキストを活用します。私はそれを「何かを教えてください」と呼んでいます。
アイデアは、「退屈したら、Reddit に行く代わりに、クロードに何かを教えてもらうことができる」というものです。これは最も効率的な学習方法ではないかもしれませんが、Reddit をスクロールするよりは優れています。 Claude では、これをカスタム命令を含むプロジェクトとして設定しました。現在使用しているプロンプトは次のとおりです。
プロジェクトの説明: ソクラテス教育セッション
このプロジェクトでは、単に概念を説明するのではなく、質問をして私の知識を測り、発見を導くソクラテス メソッドを使用して、私に新しいことを教えていただきます。
分野 (私の専門知識が低い順):
経済学 (行動またはその他)
あなたのアプローチ:
私が「何かを教えて」と言うと、次の手順を実行します。 「<トピック> について何か教えて」と言った場合、最初の 2 つのステップをスキップします。
繰り返しを避けるために、このプロジェクトの以前のチャットを参照してください
私の専門分野から多様なトピックを選択してください
質問を使って自分がすでに知っていることを評価する
直接説明するのではなく、対話を通じて洞察を得るように導いてください
私の回答がレッスンの方向性と深さを形作るようにしましょう
目標: ガイド付きの探求を通じて概念を発見して理解できるようにし、知っていることに基づいて、自分自身の推論を通じてギャップを埋められるようにします。
セッション全体でトピックを多様に保ちます。
セッションの終わりに、確認してさらに読むために一次資料を参照するように指示します。 Web サイト、論文、ポッドキャスト、書籍の順に優先します。
これはうまく機能します。トピックの多様性は良好で、特にクロードが私の事前知識を評価し、それに反応するため、ソクラテスのメソッドは機能します。これまでのところ、クロードは私にアレのパラドックスについて教えてくれました。

いくつか例を挙げると、協和音の物理学や料理における塩の化学などです。クロードは、プロジェクト内の以前のチャットをリストしてトピックを追跡できます。唯一の難点は、クロードが最初のユーザー インタラクションに基づいてチャットに単に「Learn something new」という名前を付けることが多いため、チャットに正しい名前を付けることです。 Claude にはチャットの名前を変更するためのツール呼び出しがありません。そのため、代わりに、最後に名前の提案を求めてから、自分でチャットの名前を変更しています。プロンプトの最後の指示により、クロードの発言を確認し、さらに深く掘り下げることができます。
当初、私はクロードにソクラテス式メソッドを使用するように指示しませんでしたが、その方がはるかに効果的です。 「情報ダンピー」が大幅に減ります。私がトピックをよく知っているとき、クロードは基本をうまくショートカットします。
これにより、非決定性とテキストという LLM の 2 つの長所が効果的に組み合わされます。トピックは多様に保たれており、私はクロードのトピックに関する膨大な知識を頼りに、興味深い論点を見つけています。クロードとすべての LLM は会話が得意で、これはソクラテス メソッドの裏返しにも当てはまります。最後に、提供されたソースは幻覚から保護し、LLM を超える次のステップを提供します。
著作権 2013 - 2025 ヒューゴ チュニウス

## Original Extract

Leveraging Claude to learn, instead of doom scrolling with a touch of the Socratic method.

Claude, Teach Me Something
Home
I’ve been experimenting with a new Claude workflow as an alternative to doom scrolling. It leverages what LLMs do best: non-determinism and text. I call it “Teach me something”.
The idea is: if I’m bored, instead of going on Reddit, I can ask Claude to teach me something. This might not be the most efficient learning method, but it beats scrolling Reddit. In Claude I’ve set this up as a project with custom instructions. The prompt I’m currently using is:
Project Instructions: Socratic Teaching Sessions
In this project you will teach me something new using the Socratic method - asking questions to gauge my knowledge and guide my discovery rather than simply explaining concepts.
Areas (in order of my decreasing expertise):
Economics (behavioral or otherwise)
Your approach:
When I say “Teach me something,” you will perform the following steps. If I say “Teach me something about <topic>” you skip the first 2 steps.
Consult previous chats in this project to avoid repetition
Choose a diverse topic from one of my areas
Use questions to assess what I already know
Guide me toward insights through dialogue rather than direct explanation
Let my responses shape the direction and depth of the lesson
Goal: Help me discover and understand concepts through guided inquiry, building on what I know and filling gaps through my own reasoning.
Keep the topics diverse across sessions.
At the end of a session direct me towards primary sources to confirm and read more. Prefer websites, papers, podcast, and books in that order.
This works nicely. The topic diversity has been good and the Socratic method works, especially because Claude gauges and responds to my prior knowledge. So far Claude has taught me about The Allais Paradox, the physics of consonance, and the chemistry of salt in cooking, to name a few. Claude can list previous chats within a project to keep track of topics. The only point of friction, is ensuring chats are named correctly as Claude will often just name them “Learn something new” based on the first user interaction. Claude lacks a tool call to rename chats, so instead I’ve been asking it to suggest a name at the end and then I rename the chat myself. The last instruction in the prompt ensures I can verify what Claude has said and dig deeper.
Initially I didn’t instruct Claude to use the Socratic method, but that works much better. It’s significantly less “information-dumpy”. When I know a topic well, Claude successfully shortcuts the basics.
This effectively combines two strengths of LLMs: non-determinism and text. The topics are kept diverse and I rely on Claude’s vast knowledge of topics to find interesting points of discussion. Claude, and all LLMs, are great at conversation and this extends to the back and forth of the Socratic method. At the end, the provided sources protect against hallucination and offer a next step beyond the LLM.
Copyright 2013 - 2025 Hugo Tunius
