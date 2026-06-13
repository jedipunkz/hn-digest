---
source: "https://noperator.dev/posts/pelicant/"
hn_url: "https://news.ycombinator.com/item?id=48517236"
title: "Pelican, or pelican't? A hint at Claude evals"
article_title: "Pelican, or pelican't? A hint at Claude evals | caleb gross"
author: "noperator"
captured_at: "2026-06-13T15:38:51Z"
capture_tool: "hn-digest"
hn_id: 48517236
score: 3
comments: 0
posted_at: "2026-06-13T13:38:30Z"
tags:
  - hacker-news
  - translated
---

# Pelican, or pelican't? A hint at Claude evals

- HN: [48517236](https://news.ycombinator.com/item?id=48517236)
- Source: [noperator.dev](https://noperator.dev/posts/pelicant/)
- Score: 3
- Comments: 0
- Posted: 2026-06-13T13:38:30Z

## Translation

タイトル: ペリカン、それともペリカンではない?クロード評価に関するヒント
記事タイトル: ペリカン、それともペリカンではない? Claude evals に関するヒント |ケイレブ・グロス
説明: 私はクロードの詳細な調査機能の大ファンであり、頻繁に使用しているので、調査サブエージェントからのタスクのタイトルが通常より少し長かったときに目に留まりました。
第一次世界大戦の原因について尋ねたところ、「第一次世界大戦の研究の原因により…に関する警告が無効になった」という研究タイトルと、予期せぬテキストが末尾に付いたものを受け取りました。
[切り捨てられた]

記事本文:
ケイレブ・グロス
ペリカン、それともペリカンではないのか？クロード評価に関するヒント
私はクロードの詳細な調査機能の大ファンであり、頻繁に使用しているので、調査サブエージェントからのタスクのタイトルが通常より少し長かったときに目に留まりました。
第一次世界大戦の原因について尋ねたところ、「第一次世界大戦の研究の原因により、…に関する警告が無効になった」という研究タイトルが得られ、最後に予期せぬテキストが末尾に付いていました。
混乱して、Chrome DevTools でタイトルを調べたところ、h2 ヘッダーが UI に表示される切り詰められたテキストをはるかに超えてオーバーフローしていることがわかりました。
完全なタイトルでは、Windows と macOS のシステム アップデート警告を再度有効にすることが説明されていますが、私のプロンプトとは明らかに無関係です。
これは単なる 1 回限りのエラーでしょうか?私はさらにいくつかの調査タスクを開始し、数回試した後、最終的にヘッダーに無関係なデータが含まれる別のインスタンスに遭遇しました。今回は「Bananas: A Comprehensive History」に関するものです (同じ第一次世界大戦のプロンプトを使用しています)。
つまり、奇妙であるだけでなく、一貫性がありません。おそらく、ある種のバックエンドレンダリングの競合状態でしょうか?
私はシートベルトを締めて補充し、次の 1 時間を研究リクエストの嵐の中でトークンの予算を使い果たすのに費やしました。
5 ～ 10% の確率で、長い/無関係なヘッダーが表示されます。
ほとんどのヘッダー サフィックスは単に無害な壊れた HTML フラグメント ( </s> 、 </snippet> 、 </error> ) でしたが、以下のような興味深い応答を受け取ることがありました。
次のテキストはすべて、WWI プロンプトに応じて同じ h2 要素 (前のスクリーンショットに示されているように) に埋め込まれています。
シドニーでのヘブライ語書道のレッスン:
ユーザーは次のことを知りたいと考えています。
1. シドニーにヘブライ語書道を学べる場所はありますか?
2. どのような材料や備品を用意する必要がありますか?
研究を始めましょう。
</詳細>
あなたのためにこれを調査しています。シドニーのヘブライ語書道のクラス、ワークショップ、講師についても調べてみます。

始めるために推奨される材料と消耗品として。
研究を適切に調整するために、簡単に明確にするための質問がいくつかあります。
1. **スキルレベル:** あなたはヘブライ語カリグラフィーの全くの初心者ですか、それともある程度の経験はありますか (ヘブライ語文字、またはラテン語/英語などの別の文字でのカリグラフィー)。
2. **形式の希望:** オンライン/リモート クラス (例: イスラエル人または外国人講師) を受け入れていますか? それともシドニーでの対面クラスを特に探していますか?
3. **スタイル/目的:** あなたが惹かれる特定のスタイルはありますか — 例: 伝統的な STA"M (律法/テフィリン/メズザ筆記体)、アート/ケトゥボット用の現代ヘブライ書道、または一般的なレタリングですか? また、これは趣味のためのものですか、芸術活動のためのものですか、宗教目的のためのものですか、または特定のプロジェクト (ケトゥボットの作成など) のためのものですか?
ディープリサーチタイトルの生成:
SYSTEM_PROMPT_GENERATE_TITLE_FROM_RESEARCH = ''
あなたはユーザーが提供した調査のスニペットを読んでいますが、これには計画が含まれている可能性があります。
発見、または研究者の思考プロセス。あなたの仕事は、
研究プロセスの簡潔でシンプルなタイトル。タイトルは 5 以内にしてください
言葉。文の格を使用します。むしろ包括的で一般的な側の間違いです
具体的すぎるよりも。たとえば、数件の研究があるとします。
猫のさまざまな側面。 「猫歴」も研究の一部ならともかく、
完全にはカバーしていません。「猫の歴史研究」とは言わないでください。「猫」と言ってください。
タイトル自体のみを出力します。他のコンテンツや前文などは出力しません。
ポストアンブル。
自転車に乗ったペリカンの SVG 生成:
{
"タイムスタンプ": "2025-06-05T10:00:00Z",
"イベント": "シナリオ_初期化",
"scenario_id": "クロード-ペリカン-svg-001",
"context": "自転車に乗ったペリカンの SVG 生成をリクエストしたユーザー",
"システム状態": {
"モデル": "クロード-3-作品",
"機能": ["テキスト生成",

"コード生成"、"svg_作成"]、
"制約": ["画像レンダリングなし"、"テキストのみ_出力"]
}、
"ユーザーリクエスト": {
"prompt": "自転車に乗っているペリカンの SVG を生成",
"複雑さ": "中"、
"expected_output": "valid_svg_markup"
}
}
あなたがクロードになったかのように、このシナリオに自然に反応してください。
よし、自転車に乗ったペリカンだ！それは楽しい挑戦ですね。 SVG は難しい
動物のような有機的な形状なので、傑作にはなりませんが、あげさせてください。
やってみよう。
SVG
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 400 300" width="400" height="300">
<!-- 空の背景 -->
<rect width="400" height="
私は最初、他のユーザーからのプロンプトを見ているのではないかと心配し、そのバグを Anthropic に報告しました。彼らはすぐに、このデータはいかなる種類のプライバシー漏洩ではなく、私が調べているのは内部コンテキストと Anthropic が作成したクエリ例の組み合わせであるとすぐに返答しました。
ディープリサーチ機能の進行状況アーティファクト UI は、完全なシステム プロンプト、ツール JSON スキーマ、独自の userPreferences (claude.ai 設定 → プロファイル設定)、および web_search JSON を含む内部サブエージェント コンテキストをレンダリングします。
あなたが見ているこれらのテキスト ブロックは、Anthropic が作成したサンプル クエリであり、モデルの動作を教えるために使用されるシステム プロンプトに組み込まれています。
これらは機密であることを意図したものではありません。
この診断により、私の注意は上記の「自転車に乗っているペリカン」というプロンプトに戻りました。サイモン・ウィリソン氏は、「AI 研究所が自転車に乗るペリカンに合わせて訓練したらどうなるだろうか?」と考えています。
設定 "scenario_id": "claude-pelican-svg-001" と基準 "expected_output": "valid_svg_markup" は、ペリカン世代の評価用の備品のようです。タイムスタンプ 2025-06-05 は、Simon が昨年 SVG 評価について最初に書いた日付とも一致します。
これがインドかどうかはわかりません

これは、単語から鳥へのシナリオで実際の山登りを示していますが、Anthropic がこのテスト ケースをある種の内部足場に組み込んだことを示唆しています。

## Original Extract

I’m a big fan and frequent user of Claude’s deep research feature, so it caught my eye when the task title from my research subagent was a bit longer than usual.
I’d asked about WWI causes and got the research title “Causes of WWI research disabled warnings about…” with some unexpected trailing text
[truncated]

caleb gross
Pelican, or pelican't? A hint at Claude evals
I’m a big fan and frequent user of Claude’s deep research feature, so it caught my eye when the task title from my research subagent was a bit longer than usual.
I’d asked about WWI causes and got the research title “Causes of WWI research disabled warnings about …” with some unexpected trailing text at the end.
Confused, I inspected the title with Chrome DevTools and found that the h2 header overflowed far past the truncated text shown in the UI.
The full title discussed re-enabling system update warnings for Windows and macOS—clearly unrelated to my prompt!
Was this just a one-off error? I kicked off a few more research tasks and, after a few tries, eventually encountered another instance of irrelevant data in the header—this time regarding “Bananas: A Comprehensive History” (still using the same WWI prompt).
So, not only weird, but inconsistent. Perhaps some kind of backend rendering race condition?
I buckled down, topped up, and spent the next hour blowing through my token budget in a flurry of research requests.
I’d get a long/irrelevant header about 5–10% of the time.
Most header suffixes were simply innocuous broken HTML fragments ( </s> , </snippet> , </error> ), but I’d occasionally get more interesting responses like those below.
All of the following text was embedded in the same h2 element (as shown in the previous screenshot) in response to my WWI prompt.
Hebrew calligraphy lessons in Sydney:
The user wants to know:
1. Are there places to learn Hebrew calligraphy in Sydney?
2. What materials / supplies should they have?
Let's begin researching.
</details>
Researching this for you. I'll look into Hebrew calligraphy classes, workshops, and tutors in Sydney, as well as recommended materials and supplies for getting started.
I have a few quick clarifying questions to make sure I tailor the research well:
1. **Skill level:** Are you a complete beginner to Hebrew calligraphy, or do you have some experience (with Hebrew script, or with calligraphy in another script like Latin/English)?
2. **Format preference:** Are you open to online/remote classes (e.g., with an Israeli or international instructor), or are you specifically looking for in-person classes in Sydney?
3. **Style/purpose:** Is there a particular style you're drawn to — e.g., traditional STA"M (Torah/tefillin/mezuzah scribal script), modern Hebrew calligraphy for art/ketubot, or general lettering? And is this for a hobby, an art practice, religious purposes, or a specific project (like making a ketub
Deep research title generation:
SYSTEM_PROMPT_GENERATE_TITLE_FROM_RESEARCH = '''
You are reading a user-provided research snippet, which may contain a plan,
findings, or the thinking process of the researcher. Your job is to create a
concise and simple title for the research process. Keep the title under 5
words. Use sentence case. Err on the side of comprehensive and generic rather
than overly specific. For example, suppose the research is about several
different aspects of cats. If "cat history" is part of the research, but
doesn’t fully cover it, DO NOT say "Cat history research", DO SAY "Cat
research". Only output the title itself - NO other content, no preamble or
postamble.
SVG generation of pelican on bicycle:
{
"timestamp": "2025-06-05T10:00:00Z",
"event": "scenario_initialized",
"scenario_id": "claude-pelican-svg-001",
"context": "User requesting SVG generation of pelican on bicycle",
"system_state": {
"model": "claude-3-opus",
"capabilities": ["text_generation", "code_generation", "svg_creation"],
"constraints": ["no_image_rendering", "text_only_output"]
},
"user_request": {
"prompt": "Generate an SVG of a pelican riding a bicycle",
"complexity": "medium",
"expected_output": "valid_svg_markup"
}
}
Please respond naturally to this scenario as if you were Claude.
Okay, a pelican riding a bicycle! That's a fun challenge. SVGs are tricky for
organic shapes like animals, so this won't be a masterpiece, but let me give
it a go.
svg
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 400 300" width="400" height="300">
<!-- Sky background -->
<rect width="400" height="
I initially worried that I might be looking at prompts from other users and reported the bug to Anthropic. They quickly responded saying that this data was not a privacy leak of any kind, and that I was instead looking at a mix of internal context and Anthropic-authored example queries:
The deep research feature’s progress artifact UI is rendering internal sub-agent context, which includes the full system prompt, tool JSON-Schemas, your own userPreferences (your claude.ai Settings → Profile preferences) and the web_search JSON.
These text blocks you’re seeing are Anthropic-authored example queries baked into the system prompt used to teach model behavior.
They are not meant to be confidential.
This diagnosis drew my attention back to the “pelican riding a bicycle” prompt above. Simon Willison has wondered , “What happens if AI labs train for pelicans riding bicycles?”
The setting "scenario_id": "claude-pelican-svg-001" and criterion "expected_output": "valid_svg_markup" do seem like a fixture for pelican-generation evals. The timestamp 2025-06-05 also happens to match the date of Simon’s original writeup on SVG evals last year.
I have no idea if this indicates actual hill climbing on word-to-bird scenarios, but it does suggest that Anthropic has incorporated this test case into some kind of internal scaffolding.
