---
source: "https://alexocallaghan.com/minecraft-phonics-game"
hn_url: "https://news.ycombinator.com/item?id=48537272"
title: "Show HN: I used Claude to build my son a Minecraft phonics game"
article_title: "Building my son a Minecraft phonics game with AI"
author: "aocallaghan17"
captured_at: "2026-06-15T08:10:21Z"
capture_tool: "hn-digest"
hn_id: 48537272
score: 1
comments: 0
posted_at: "2026-06-15T06:17:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I used Claude to build my son a Minecraft phonics game

- HN: [48537272](https://news.ycombinator.com/item?id=48537272)
- Source: [alexocallaghan.com](https://alexocallaghan.com/minecraft-phonics-game)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T06:17:08Z

## Translation

タイトル: Show HN: クロードを使って息子に Minecraft フォニックス ゲームを作成しました
記事のタイトル: AI を使用して息子に Minecraft フォニックス ゲームを作成する
説明: AI を使用して息子のために Minecraft フォニックス ゲームを構築し、息子の興味や学習ニーズに合わせてゲームを調整しました

記事本文:
メイン コンテンツに移動 アレックス オキャラハン このページについて このページについて
すべての執筆 AI を使用して息子に Minecraft フォニックス ゲームを構築する
4 min read · 627 Words ai-agents 私の息子は読み書きを学んでいて、フォニックスのテストが近づいています。ほとんどの子供たちと同様、フォニックスの練習は必ずしも彼のお気に入りではありません。それは大変な作業であり、ハードワークにはモチベーションが必要です。彼は、プロとしか言いようのないほどの献身的な姿勢で Minecraft を愛しています。
そこで私は、Claude AI とのたった 1 回の会話で、彼に Minecraft フォニックス ゲームを作成しました。
私はクロードと会話を始め、私が望んでいることを説明しました。Minecraft をテーマにした、子供がフォニックスを練習するのに役立つシンプルなブラウザ ゲームで、2 つのモードがあります。1 つは絵を見て文字タイルを使用して単語の綴りをするモード、もう 1 つは単語を読んで 4 つの選択肢から一致する絵を選ぶモードです。私は、英国のカリキュラムの主要段階 1 にある子供に適切な言葉を使用する必要があると説明しました。
単一の応答で返されたのは、2 つのプレイ可能なゲーム モードを備えた Minecraft のようにスタイル設定された、動作する React コンポーネントでした。
私はエンジニアなので、放っておくのは難しいので、さらに進めました。 Claude Code を使用してファイルを適切なプロジェクトに取り込みました。Vite をセットアップし、Vitest を追加し、TypeScript に切り替えました。実際の Minecraft アイテム アセットを含む npm パッケージを見つけて、絵文字プレースホルダーを実際のゲーム内スプライトと交換しました。難易度を追加して、彼の上達に合わせてワードプールがスケールするようにしました。
TypeScript はデータ構造を明示的にするので、テストを行うことで、何かが壊れていることを心配することなく機能を追加できます。実際のアセットはアイコンをより認識しやすくするのに役立ちますが、絵文字の選択の一部には少し疑問がありました。
これによりいくつかの改善が得られましたが、そこから価値を得る必要はありませんでした。最初のバージョンはプレイ可能でした、Mi

necraft をテーマにしており、適切なフォニックス コンテンツをカバーしています。エンジニアリングの知識のない親は、Claude の UI 内のプレイアブル バージョンを使用するだけで済みます。私が行った変更も、Claude Code を使用していくつかのプロンプトを表示しただけです。
私が息子にこのゲームについて説明すると、息子はすぐに興奮し、学校を出てすぐに最初の反応は、「プレイしてもいいですか？」と熱心に尋ねることでした。それ以来、彼は定期的にこの曲を演奏してほしいと頼んできた。
私が予想していなかったのは、共同設計でした。彼はほぼすぐにフィードバックを送り始めました。ネザライトは最強の鎧を作るものであり、難しい言葉になるため、本当に入れるべきだということです。彼はそれに投資し、フォニックスの練習は彼が実際にやりたいと感じたことと並行して行われていました。
私が考えていなかった副作用もありました。彼はコンピューターのマウスを使用した経験があまりなく、タッチスクリーン アプリでは得られなかった細かい運動の練習とマウスに慣れるためにゲームを操作することになりました。
EdTech 製品は平均的な子供向けに作られています。ソフトウェアの経済的観点から、できるだけ多くのユーザーにリーチする必要があるため、そうする必要があります。私の息子は普通の子供ではありません。彼は、学習の特定の瞬間に、特定のゲームに夢中になり、特定のテストを受け、特定の子供です。市販のフォニックス アプリはそのために最適化されておらず、彼専用に設計されたアプリと同じような魅力をもたらすことはありません。
AI によって変わったのは、カスタマイズされたものを構築するコストが劇的に下がったことです。これは、親が子供のためにゲームを作成する以上の意味を持ちます。家庭教師は、生徒の特定の興味に基づいて演習を作成できます。教師はカリキュラムではなく教材をクラスに適応させることができます。制約は常に時間と技術的アクセスであり、想像力ではありませんが、

AIによって変わるat。
最初のゲームのコードは、クロードとの 1 つの会話から生まれました。最終バージョンのコードは GitHub で見つけて、ここで再生できます。
ツーリング前の測定: エンジニアリングへの AI 投資はどこに向けられるべきか
理解を失うことなくエージェントと迅速に対応する
ソフトウェア開発、プラットフォーム エンジニアリング、AI エージェントがソフトウェアの構築方法をどのように変えるかについて書いています。

## Original Extract

I built a Minecraft phonics game for my son using AI, tailoring a game to his interests and learning needs

Skip to main content Alex O'Callaghan On this page On this page
All writing Building my son a Minecraft phonics game with AI
4 min read · 627 words ai-agents My son is learning to read and write, and has a phonics test coming up. Like most kids, phonics practice isn't always his favourite thing — it's hard work, and hard work requires motivation. He also loves Minecraft with a level of dedication that I can only describe as professional.
So I built him a Minecraft phonics game in just one conversation with Claude AI.
I opened a conversation with Claude and described what I wanted: a simple browser game to help a child practice phonics, Minecraft-themed, with two modes — one where he sees a picture and has to spell the word using letter tiles, one where he reads a word and picks the matching picture from four options. I described that the words should be appropriate for a child at Key Stage 1 in the UK curriculum.
What came back in a single response was a working React component styled to look like Minecraft with two playable game modes:
Because I'm an engineer and it's hard to leave things alone, I took it further. I pulled the file into a proper project using Claude Code — set up Vite, added Vitest, switched to TypeScript. I found an npm package with real Minecraft item assets and swapped out the emoji placeholders for the actual in-game sprites. I added difficulty levels so the word pool scales as he improves.
TypeScript makes the data structures explicit, and with tests I can add features without worrying I've broken something. The real assets helps make the icons much more recognisable, some of the emoji choices were a bit questionable.
While this gave some improvements, it wasn't required to get value from it. The first version was playable, Minecraft-themed, and covered the right phonics content. A parent with no engineering background could have just used the playable version within Claude's UI. Even the changes I did make were just a few prompts with Claude Code.
When I explained the game to my son he was instantly excited, his first reaction getting out of school was to eagerly ask if he could play it. He's been regularly asking to play it in the days since.
What I didn't anticipate was the co-design. He started giving feedback almost immediately - that we really should include netherite since it makes the strongest armour and would be a tricky word. He was invested in it and the phonics practice was happening alongside something he actually wanted to do and felt ownership of.
There were side effects I hadn't considered either. He hasn't had much experience using a computer mouse and navigating the game turned into fine motor practice and mouse familiarity that he wouldn't have got from a touchscreen app.
EdTech products are built for the average child. They have to be as the economics of software require reaching as many users as possible. My son is not the average child. He is a specific child who is obsessed with a specific game, sitting a specific test, at a specific moment in his learning. No commercial phonics app was going to be optimised for that and wouldn't bring the same engagement as something designed just for him .
What AI has changed is that the cost of building something tailored has dropped dramatically. This has implications beyond parents building games for their children. Tutors could build exercises around a student's specific interests. Teachers could adapt materials to a class rather than a curriculum. The constraint has always been time and technical access, not imagination, but that's changing with AI.
The code for the initial game came from a single conversation in Claude. You can find the code for the final version on GitHub and play it here .
Measuring before tooling: where AI investment in engineering should go
Moving fast with agents without losing comprehension
I write about software development, platform engineering and how AI agents are changing the way we build software.
