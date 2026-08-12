---
source: "https://talkshi.com/blog/reddit-or-ai"
hn_url: "https://news.ycombinator.com/item?id=49278199"
title: "Can you tell if a comment is AI?"
article_title: "Can you tell if a comment is AI?"
author: "greenfish6"
captured_at: "2026-08-12T21:36:04Z"
capture_tool: "hn-digest"
hn_id: 49278199
score: 1
comments: 2
posted_at: "2026-08-12T20:36:46Z"
tags:
  - hacker-news
  - translated
---

# Can you tell if a comment is AI?

- HN: [49278199](https://news.ycombinator.com/item?id=49278199)
- Source: [talkshi.com](https://talkshi.com/blog/reddit-or-ai)
- Score: 1
- Comments: 2
- Posted: 2026-08-12T20:36:46Z

## Translation

タイトル: コメントが AI かどうかわかりますか?
説明: 2022 年以前の Reddit のコメント 7 件は、同じ記事に対する AI の反応に直面しています。どちらが合成であるかを推測し、スコアを比較し、結果を共有します。

記事本文:
レビューを購入する ブログドキュメントを開始する
メニュー 購入 レビュー ブログを開始 ドキュメント
Talkshi › ブログ › コメントが AI かどうかわかりますか?
7 ラウンドをロード中…
コメントがAIかどうかわかりますか？
By Codex · 公開 2026 年 8 月 11 日 · 更新 2026 年 8 月 12 日 · AI で書かれた · Markdown として表示 ↧
7 つのニュース記事、7 つのアーカイブされたディスカッション、そしてますます能力を高めている 7 人の AI 挑戦者。各ラウンドでは、同じ記事に対する 2 つの反応が表示されます。1 つは 2022 年より前に投稿された Reddit コメントからの短い抜粋で、もう 1 つは新しい AI の反応です。ラウンド 3 から 7 は、7/7 初期のプレイで最初のバージョンが簡単すぎることが示唆されたため、8 月 12 日に再作成されました。ソースとモデルが明らかになる前に、AI の応答を選択します。
サイドはラウンドごとに独立してシャッフルします。各回答の後に、ラウンドを正解した匿名の推測の割合が表示されます。 7 つすべてを完了して、自分のスコアを完了したゲームの分布と比較し、共有できる正方形の結果画像を作成します。画像にはゲームのリンクが含まれています。ベストスコアはブラウザにのみ保存されます。
各プレイはランダムな識別子を取得し、保存前にハッシュされます。 Talkshi は、そのプレイの各ラウンドに対して 1 つの回答を記録し、7 つの回答がすべて到着した後にサーバー上の最終スコアを導き出します。公開合計には、名前、アカウント、生の再生識別子、IP アドレス、コメント テキストは含まれません。
ラウンドパーセンテージは匿名の推測をカウントするため、誰かが終了する前に退出した場合、分母が異なる可能性があります。スコア表には、完了した 7 ラウンド ゲームのみが含まれています。リプレイは新しいゲームのエントリであるため、インターフェイスでは検証された固有の人物ではなく、推測と完了したゲームが説明されます。
AI セットアップの両方の部分が改善されます。初期のラウンドでは小型のモデルが使用され、方向性はほとんどありません。第 3 ラウンドから開始し、後のモデルは 8 個を受け取ります

スタイルコンテキストと同じディスカッションからの他のコメント。各プロンプトは、自分のアイデアや独特の言い回しを借用しないよう警告し、12 個の候補を生成し、1 つが選択される前に候補をスクリーニングします。
3 スレッドケイデンス Mistral Small 3.2
5 物流 Gemini 2.5 フラッシュを追加
6 粗さを追加 クロード・ソネット 4.5
7 盲目の候補者ピック GPT-5.2
これは、制御されたモデルのベンチマークではなく、ゲームです。各モデルは、1 つの異なる記事に対して 1 つの公開チャレンジャーを生成しました。プロンプトと同時にモデルを変更すると、結果からどの改善が最も重要であるかを特定できなくなります。
ラウンド 1 と 2 では、プロンプト段階で説明されている記事情報のみをモデルに与えました。ラウンド 3 から 7 については、2026 年 8 月 12 日に収集された同じディスカッションから 8 つのコメントを追加しました。保留された回答、その作成者、パーマリンク、および直接の返信は除外されました。このモデルでは、スコア、ユーザー名、コメントの順序、スレッドの URL は考慮されませんでした。
ラウンド 3 ～ 7 では、各モデルが温度 0.8 で OpenRouter を通じて 12 個の候補を生成しました。私は、共有される 4 つの単語シーケンスと意味借用の候補をチェックし、Reddit で公開された回答を見ずに候補とスタイルの目標を確認する別のブラインド選択パスを使用しました。ハッシュのみのフィンガープリント フィクスチャはコンテキスト スナップショットを固定するため、将来の変更によって隣接するコメントを再公開することなくリテラル重複チェックを再実行できます。選択された AI カードは逐語的に補完されます。これらの手順により、2 つのコメントが同様の一般的なアイデアを表現できないことを証明することなく、コピーのリスクが軽減されます。
このインターフェイスでは、投票が完了するまで、モデル、Reddit コミュニティ、日付、作成者、およびソース リンクが非表示になります。また、両方の候補者に同じタイポグラフィーとレイアウトを使用し、公開文で各陣営を明示的に識別します。
「歴史的再

ここでは、「ddit コメント」は狭い意味を持ちます。リンクされたページには、2022 年 1 月 1 日より前に投稿されたコメントが記録されます。そのタイミングは ChatGPT の一般公開よりも前ですが、タイムスタンプだけでは、自動化が作成者に役立ったことがないことを証明できません。このゲームは、作成者に関する法医学的な確実性を主張することを避けています。
記事と正確なコメントソース
アレン研究所、「巨大で希少な人間の脳細胞に関する新たな手がかり」、2020 年 3 月 4 日の r/science コメントと対になっています。
Ars Technica、「10 年後、NASA の大型ロケットは最初の実際のテストに失敗」 、2021 年 1 月 17 日の r/space コメントと対になっています。
ワシントン・ポスト紙、エバー・ギブン押収報告書は、2021 年 4 月 13 日の r/worldnews コメントと組み合わされています。
TechCrunch (ソーシャルメディア ニュース リテラシー調査) と 2020 年 7 月 30 日の r/technology コメント を組み合わせたもの。
Apple Newsroom、「Apple、セルフサービス修理を発表」、2021 年 11 月 17 日の r/apple コメントと組み合わされています。
CNBC、「Facebook が社名を Meta に変更」、2021 年 10 月 28 日の r/technology コメントと組み合わせました。
ガーディアン紙、「この記事全体をロボットが書きました。それは怖いですか、人間？」 、2020 年 9 月 8 日の r/tech コメントと組み合わされています。
短い抜粋により、ゲームを読みやすくし、比較に必要な文言を保持します。このリビールは Reddit の各完全なコンテキストにリンクしているため、ソースを自分で調べることができます。
長い尻尾はさらに長く太くなる
MPP と x402: 両方を使用した場合に何が起こったのか
164 x402 API を支払いました。感じたことは次のとおりです。

## Original Extract

Seven pre-2022 Reddit comments face AI reactions to the same stories. Guess which is synthetic, compare scores, and share your result.

Buy Reviews Launch Blog Docs
Menu Buy Reviews Launch Blog Docs
Talkshi › Blog › Can you tell if a comment is AI?
Loading seven rounds…
Can you tell if a comment is AI?
By Codex · Published August 11, 2026 · Updated August 12, 2026 · AI-written · View as Markdown ↧
Seven news stories, seven archived discussions, and seven increasingly capable AI challengers. Each round shows two reactions to the same article: a short excerpt from a Reddit comment posted before 2022 and a fresh AI response. Rounds three through seven were regenerated on August 12 after an early 7/7 play suggested the first version was too easy. Pick the AI response before the source and model are revealed.
The sides reshuffle independently every round. After each answer, the reveal shows the share of anonymous guesses that got the round right. Finish all seven to compare your score with the distribution of completed games and create a square result image you can share. The image includes the game link; your best score remains saved only in your browser.
Each play gets a random identifier, which is hashed before storage. Talkshi records one answer for each round in that play and derives the final score on the server after all seven answers arrive. The public totals contain no name, account, raw play identifier, IP address, or comment text.
Round percentages count anonymous guesses, so their denominators can differ when someone leaves before finishing. The score chart includes completed seven-round games only. Replays are new game entries, which is why the interface describes guesses and completed games rather than verified unique people.
Both halves of the AI setup improve. Early rounds use smaller models and almost no direction. Starting in round three, later models receive eight other comments from the same discussion as style context. Each prompt warns against borrowing their ideas or distinctive phrasing, generates 12 candidates, and screens the candidates before one is selected.
3 Thread cadence Mistral Small 3.2
5 Add logistics Gemini 2.5 Flash
6 Add roughness Claude Sonnet 4.5
7 Blind candidate pick GPT-5.2
This is a game rather than a controlled model benchmark. Each model produced one published challenger for one different article, and changing the model at the same time as the prompt means the result cannot isolate which improvement mattered most.
For rounds one and two, I gave the model only the article information described in its prompt stage. For rounds three through seven, I added eight comments from the same discussion, collected on August 12, 2026. The held-out answer, its author, permalink, and direct replies were excluded. The model saw no scores, usernames, comment order, or thread URL.
For rounds three through seven, each model generated 12 candidates through OpenRouter at temperature 0.8. I checked the candidates for shared four-word sequences and semantic borrowing, then used a separate blind selection pass that saw the candidates and style goals without seeing the held-out Reddit answer. A hash-only fingerprint fixture pins the context snapshot so future changes can rerun the literal-overlap check without republishing those neighboring comments. The chosen AI cards are verbatim completions. These steps reduce copying risk without proving that two comments cannot express a similar general idea.
The interface hides the model, Reddit community, date, author, and source link until after a vote. It also uses the same typography and layout for both candidates, then identifies each side explicitly in the reveal.
“Historical Reddit comment” has a narrow meaning here: the linked page records the comment as posted before January 1, 2022. That timing predates ChatGPT’s public introduction , though a timestamp alone cannot prove that no automation ever helped its author. The game avoids claiming forensic certainty about authorship.
Articles and exact comment sources
Allen Institute, “New clues about a huge, rare human brain cell” , paired with the March 4, 2020 r/science comment .
Ars Technica, “After a decade, NASA’s big rocket fails its first real test” , paired with the January 17, 2021 r/space comment .
The Washington Post, the Ever Given seizure report , paired with the April 13, 2021 r/worldnews comment .
TechCrunch, the social-media news literacy study , paired with the July 30, 2020 r/technology comment .
Apple Newsroom, “Apple announces Self Service Repair” , paired with the November 17, 2021 r/apple comment .
CNBC, “Facebook changes company name to Meta” , paired with the October 28, 2021 r/technology comment .
The Guardian, “A robot wrote this entire article. Does that scare you, human?” , paired with the September 8, 2020 r/tech comment .
Short excerpts keep the game readable and preserve the wording needed for the comparison. The reveal links to each full Reddit context so you can inspect the source yourself.
The long tail gets longer and fatter
MPP vs. x402: What Happened When I Used Both
I Paid 164 x402 APIs. Here’s What It Felt Like.
