---
source: "https://claude.com/blog/a-field-guide-to-claude-fable-finding-your-unknowns"
hn_url: "https://news.ycombinator.com/item?id=48845045"
title: "A field guide to Claude Fable 5: Finding your unknowns"
article_title: "A field guide to Claude Fable 5: Finding your unknowns | Claude | Claude by Anthropic"
author: "maxloh"
captured_at: "2026-07-09T13:44:48Z"
capture_tool: "hn-digest"
hn_id: 48845045
score: 2
comments: 0
posted_at: "2026-07-09T12:48:28Z"
tags:
  - hacker-news
  - translated
---

# A field guide to Claude Fable 5: Finding your unknowns

- HN: [48845045](https://news.ycombinator.com/item?id=48845045)
- Source: [claude.com](https://claude.com/blog/a-field-guide-to-claude-fable-finding-your-unknowns)
- Score: 2
- Comments: 0
- Posted: 2026-07-09T12:48:28Z

## Translation

タイトル: クロード寓話のフィールドガイド 5: 未知のものを見つける
記事のタイトル: クロード寓話のフィールドガイド 5: 未知のものを見つける |クロード |クロード by Anthropic
説明: Claude Fable によるエージェント コーディングの実践的なパターン: Anthropic チームによる、実装前、実装中、実装後に未知の部分を見つける方法。

記事本文:
クロード寓話のフィールドガイド 5: 未知のものを見つける |クロード |クロード by Anthropic
クロード製品のご紹介 クロード
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
クロード寓話のフィールドガイド 5: 未知のものを見つける
クロード寓話のフィールドガイド 5: 未知のものを見つける
共有 リンクをコピー https://claude.com/blog/a-field-guide-to-claude-fable-finding-your-unknowns
クロード コードで作業していると、地図と領土の違いをよく思い出します。
やるべき仕事を表現した地図は、私のプロンプトであり、スキルであり、コンテキストであり、私がクロードに与えるものです。領域とは、作業を行う必要がある場所、コードベース、現実世界、実際の制約です。
地図と領土の違いを私は未知数と呼んでいます。クロードは未知のものに遭遇したとき、私が望んでいることの最善の推測に基づいて決定を下す必要があります。作業が増えれば増えるほど、クロードはより多くの未知のことに遭遇する可能性があります。
クロード・ファブルは、未知の部分を明らかにする私の能力によって作品の品質がボトルネックになっていると私が感じた最初のモデルです。
重要なのは、事前に計画を立てるだけでは必ずしも十分ではないということです。実装の深いところで未知の部分が見つかることもあれば、その未知の部分から、実際にはまったく別の方法で問題を解決する必要があるという事実が示されることもあります。
Fable を使ってみると、

これは、実装前、実装中、実装後に未知の部分を発見する反復的なプロセスです。
前へ 前へ 0 / 5 次へ 次へ Claude Code Desktop を入手
IRM https://claude.ai/install.ps1 | iex コマンドをクリップボードにコピーする またはドキュメントを読む クロード コードを試す クロード コードを試す クロード コードを試す 開発者ドキュメント 開発者ドキュメント 開発者ドキュメント 電子書籍
あなたの未知は何ですか?問題を抱えてクロードに相談したとき、私はそれを 4 つの方法に分類することがよくあります。
既知の既知のもの: これは基本的に私のプロンプトにあるものです。エージェントに希望を伝えるにはどうすればよいですか?
既知の未知: まだ理解していないが、理解していないことに気づいているものは何ですか?
未知の既知: とても明白なので決して書き留めることはできないが、見たらそれと分かるでしょうか?
未知の未知: 私がまったく考慮していないことは何ですか?私が気づいていない知識は何ですか?何かがどれほど良いものになるか私は知っていますか？
優秀なエージェント プログラマーには、未知の部分が比較的少ないです。ボリスやジャレッドのような人がプロンプトを出すのを見ていると、彼らが何を望んでいるのかを詳細に知っていることがわかります。これらは、コードベースとモデルの動作の両方と深く同期しています。
しかし、彼らは未知のことも想定しています。さまざまな意味で、未知の要素を減らして計画することは、エージェント コーディングのスキルです。しかし幸いなことに、これはクロードと協力することで向上できるスキルです。
クロードへの指導は微妙なバランスです。具体的すぎると、ピボットの方が適切な場合でも、クロードはあなたの指示に従います。曖昧すぎると、Claude は業界のベスト プラクティスに基づいて選択や仮定を行うことがよくあり、それがあなたのタスクには適合しない可能性があります。
未知の部分を考慮しないと、どちらの面でも失敗します。いつ道が障害物でいっぱいになるかわかりませんし、いつ道が開けられるかわかりませんが、それでもクロードに方向転換してもらいたいと考えています。
クロードは、未知のものをより早く発見するのに役立ちます。それは可能です

コードベースとインターネットを非常に迅速に検索し、平均的なトピックについてはユーザーよりもはるかに多くの知識を持っています。また、失敗からの反復処理も高速化できます。
このプロセスで最も重要な部分は、出発点についてのコンテキストをクロードに提供することです。たとえば、自分が思考プロセスのどの段階にいるのかを伝えます。問題とコードベースに関するあなたの経験を開示してください。そして、思考パートナーのようにあなたと一緒に働きましょう。
この記事では、これらの未知のことを明らかにするために私が使用するいくつかのパターンについて詳しく説明します。
仕事を始めるときに最も役立つことの 1 つは、盲点を理解することです。たとえば、コードベースの新しい部分に機能を記述している場合、または設計の反復処理などの慣れない作業を支援するために Claude を使用している場合、未知の部分がたくさんある可能性があります。
どのような質問をすればよいのか、どのような見栄えが良いのか、どのような歴史的な取り組みが行われてきたのか、どのような穴を避けるべきなのか、わからないかもしれません。
このような状況では、クロードに自分の未知の部分を見つけて説明してもらうのを手伝ってもらうことができます。私は「ブラインドスポットパス」や「未知の未知」という文字通りの言葉を使うのが好きです。クロードがあなたとコラボレーションを始めるための最良の方法を理解するためには、あなたが誰であるか、何を知っているかについてのコンテキストを与えることが通常重要です。
「新しい認証プロバイダーの追加に取り組んでいますが、このコードベースの認証モジュールについては何も知りません。関連する不明な点を見つけ出し、より適切にプロンプ​​トを表示できるようにブラインド スポット パスを実行してもらえませんか。」
「カラー グレーディングが何なのかわかりませんが、このビデオをグレーディングする必要があります。より適切にプロンプ​​トできるように、カラー グレーディングについての未知の知識を理解する方法を教えていただけますか?」
未知の既知がたくさんある領域で仕事をしているとき、それを見て初めて定義できる基準が含まれるとき、私はクロードに質問するのが好きです。

雷雨とプロトタイプを一緒に作りました。
実装中に未知の既知を見つけるのは（比較的）コストがかかる可能性があるため、プロトタイピングの早い段階で未知の既知を特定して言語化することは非常に価値があります。機能や仕様の小さな変更により、コードの実装が大きく異なる可能性があり、エージェントが以前の変更を元に戻すことがより困難になる場合があります。
たとえば、バックエンド ルートを接続したり、フロントエンドで追加の状態を維持したりすることなく、フレームに追加されたボタンがどのように見えるかを確認したい場合があります。
もう 1 つの例は、ビジュアル デザインです。私にとって、これは言葉で説明するのが難しいものですが、それを見れば、自分が何を望んでいるのかがわかります。このような場合、私は成果物に対するいくつかの設計アプローチを尋ねます。
また、私はほぼすべてのコーディング セッションを探索またはブレインストーミングのフェーズから始めます。これは、プロジェクトの範囲を定義する意図から始めるのに役立ちます。クロードは、私が見逃していたであろう価値の高いアプローチを頻繁に見つけて、木々の間から森を恋しく思うこともあります。ブレーンストーミングを行うと、範囲が狭すぎたり広すぎたりすることがなくなります。
「このデータ用のダッシュボードが欲しいのですが、視覚的なセンスがなく、何ができるかわかりません。大きく異なる 4 つのデザイン方向を備えた HTML ページを作成して、それらに対応できるようにしてください。」
「何かを接続する前に、偽のデータを使用して新しいエディターのツールバーを模擬する 1 つの HTML ファイルを作成します。実際のアプリに触れる前に、レイアウトに反応したいのです。」
「これが私の大まかな問題です。オンボーディング後にユーザーが離脱してしまうのです。コードベースを検索し、安価なものから最も野心的なものまで、介入できる 10 か所をブレインストーミングしてください。どれが共感を呼ぶか教えます。」
十分なブレーンストーミングを行っても、まだ不明な点がある可能性があります。
この場合、不明な点や曖昧な点についてクロードにインタビューするよう依頼します。クロードにインタビューを依頼するときは、次のことを試してください。

そして、質問を導くために問題に関するコンテキストを与えます。
「あいまいな点については一度に 1 つずつ質問し、私の答えによってアーキテクチャが変わる質問を優先してください。」
欲しいものを詳しく説明できないこともあります。たとえば、言語がわからない場合や、非常に複雑なのでかなり時間がかかる場合があります。
この場合、最良のアプローチはリファレンスです。図、ドキュメント、または画像を含めることもできますが、絶対的に最適な参考資料はソース コードです。
何かを特定の方法で実装するライブラリや、本当に気に入ったデザイン コンポーネントがある場合は、たとえそれが別の言語であっても、Fable にフォルダーを指定して、何を探すべきかを指示するだけです。これにより、たとえばスクリーンショットと比較して、クロードのマークアップと構造に関するより詳細な詳細が提供されます。
「ベンダー/レート リミッターのこの Rust クレートは、私が望む正確なバックオフ動作を実装しています。これを読んで、同じセマンティクスを TypeScript API クライアントに再実装してください。」
実装の準備ができたと思ったとき、私はクロードに実装計画をまとめてもらい、レビューしてもらうことが多いです。この計画は、データ モデル、型インターフェイス、UX フローなど、変更される可能性が最も高い部分に焦点を当てています。これにより、クロードは私が実際に変更する必要がある可能性のあるものを表面化することができます。
「HTML で実装計画を書きますが、データ モデルの変更、新しい型のインターフェイス、ユーザー向けのものなど、私が調整する可能性が最も高い決定を主導します。機械的なリファクタリングは最後に埋めてください。その部分についてはあなたを信頼しています。」
計画に満足したら、新しいセッションを作成し、成果物をプロンプトに渡します。これにより、Claude に新しいコンテキスト ウィンドウが表示され、計画から収集されたすべての情報が表示されます。たとえば、仕様ファイルとプロトタイプを渡して、質問するかもしれません。

それを実装してください。
しかし、どれだけ計画を立てても、未知のことが常に潜んでいるのも事実です。エージェントは、コード内で見つかったエッジ ケースにより、作業中に別のアプローチを取る必要があることに気付く場合があります。
Claude Code に対して、次の試行に備えて学習できるように、一時的な「implementation-notes.md」 (または .html) ファイルを保存して決定を追跡するように依頼します。
「implementation-notes.md ファイルを保管してください。計画からの逸脱を余儀なくされるようなエッジケースに遭遇した場合は、保守的なオプションを選択し、それを「逸脱」の下に記録し、続行してください。
何かを出荷する上で最も重要な部分の 1 つは、賛同と承認を獲得することです。最終的な文書で売り込みと説明のアーティファクトを構築すると、次のことが役立ちます。
査読者があなたと同じ未知のことから始めると理解が促進されます
専門家が予想していた未知の点やよくある失敗点を説明してもらいたい場合、承認を迅速化します。
「プロトタイプ、仕様、実装ノートを 1 つのドキュメントにパッケージ化し、Slack にドロップして同意を得ることができます。デモ GIF でリードします。」
長い作業セッションの後、クロードは私が思っていた以上に多くのことを達成したかもしれません。動作の多くは既存のコード パスに依存するため、コードの差分を読んでも、何が起こったのかを軽く理解することしかできません。
背景をたくさん教えてもらった後、変更についてクロードに質問してもらうと、何が起こったのかを理解するのに役立ちます。クイズに完璧に合格した場合にのみマージします。
「この変更で何が起こったのかをすべて理解していることを確認したいのです。コンテキスト、直感、何が行われたかなどを読んで理解できるように、変更に関する HTML レポートを提供してください。また、パスしなければならない変更についての下部のクイズも提供します。」
これがどのようにまとめられるか: Fable の起動
Fable のローンチビデオは次のとおりです。

クロードコードを使用してエンドツーエンドで編集されました。これは私にとって新しい領域であり、決して専門家ではありませんでした。
そこで私は自分が知っていることから始めました。クロードがコードを使用してビデオを編集し、文字起こしできることは知っていましたが、それが十分に正確であるかどうかはわかりませんでした。次に、Whisper のような文字起こしがどのように機能するのか、また、ffmpeg を使用して「うーん」や大きな一時停止などを正確に切り取ることができるかどうかをクロードに説明してもらいました。
私はクロードに、私が話している言葉に合わせた UI を作成してもらいたかったのですが、それが可能かどうか確信が持てなかったので、クロードに、リモーションと文字起こしを使用してプロトタイプのビデオを作成して、それが機能するかどうかを確認するように依頼しました。
最後に、ビデオ自体が少し落ち着いたように見えました。これはカラー グレーディングの結果であることはわかっていましたが、カラー グレーディングが何なのかはよく知りませんでした。私の最初のパスの試みは、クロードにいくつかのバリエーションを実行させて選択してもらうことでしたが、カラー グレーディングに関して「良い」とはどのようなものかを私は知らないことに気づきました。そこで、代わりに、自分の未知の部分を発見するために、クロードにカラーグレーディングについて教えてもらうことにしました。
地図と領土を一致させる
モデルの品質が向上すればするほど、適切なアプローチでより多くのことを達成できるようになります。長期的なタスクが間違っている場合は、未知の部分を定義したり、実行計画を作成したりするために、より多くの時間を費やす必要がある可能性があります。

[切り捨てられた]

## Original Extract

Practical patterns for agentic coding with Claude Fable: how to find your unknowns before, during, and after implementation, from the team at Anthropic.

A field guide to Claude Fable 5: Finding your unknowns | Claude | Claude by Anthropic
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
A field guide to Claude Fable 5: Finding your unknowns
A field guide to Claude Fable 5: Finding your unknowns
Share Copy link https://claude.com/blog/a-field-guide-to-claude-fable-finding-your-unknowns
When working with Claude Code, I’m often reminded of the difference between the map and the territory.
The map, a representation of the work to be done, is my prompts and skills and context, it’s what I give Claude. The territory is where the work needs to happen, the codebase, the real world, its actual constraints.
The difference between the map and the territory is what I call unknowns . When Claude runs into an unknown, it needs to make a decision based on its best guess of what I want. The more work being done, the more unknowns Claude might run into.
Claude Fable is the first model where I find the quality of the work is bottlenecked by my ability to clarify its unknowns.
Importantly, just planning ahead isn’t always enough. You can find unknowns deep in implementation, or your unknowns may point you to the fact that you should actually be solving the problem in a different way altogether.
I’ve found that working with Fable is an iterative process of discovering my unknowns before, during, and after implementation.
Prev Prev 0 / 5 Next Next Get Claude Code Desktop
irm https://claude.ai/install.ps1 | iex Copy command to clipboard Or read the documentation Try Claude Code Try Claude Code Try Claude Code Developer docs Developer docs Developer docs eBook
What are your unknowns? When I come to Claude with a problem I tend to break it down in 4 ways:
Known Knowns: This is essentially what is in my prompt. What do I tell the agent that I want?
Known Unknowns: What haven't I figured out yet, but I’m aware that I haven’t?
Unknown Knowns: What's so obvious I’d never write it down, but would recognize it if I saw it?
Unknown Unknowns: What haven't I considered at all? What knowledge am I not aware of? Do I know how good something can be?
The best agentic coders have relatively few unknowns. Watching someone like Boris or Jarred prompt, it is obvious to me that they know what they want in-detail. They are deeply in-sync with both the codebase and the model behaviors.
But they also assume unknowns. In many ways, reducing and planning for your unknowns is the skill of agentic coding. But luckily, this is a skill you can improve at, by working with Claude.
Instructing Claude is a delicate balance. If you are too specific, Claude will follow your instructions even when a pivot may be more appropriate. If you are too vague, Claude will often make choices and assumptions based on industry best practices that may not be a fit for your task.
When you don’t account for your unknowns, you fail both ways. You don't know when the path will be filled with obstacles, and you don’t know when the path will be clear, but you still want Claude to veer.
Claude can help you discover your unknowns faster. It can search through your codebase and the internet extremely quickly, and it knows much more about the average topic than you. It can also iterate from failure faster.
The most important part of this process is to give Claude context about your starting point. For example, tell it where you are in your thought process; disclose your experience with the problem and codebase; and let it work with you like a thought partner.
In this article I detail some of the patterns I use to uncover these unknowns including:
When starting work, one of the most useful things you can do is understand your blind spots. For example, if you’re writing a feature in a new part of the codebase, or using Claude to help you with unfamiliar work like iterating on a design, you’re likely to have a lot of unknown unknowns .
You may not know what questions to ask, what good looks like, what historical work has been done, or what potholes to avoid.
In these situations, you can ask Claude to help you find your unknown unknowns and explain them to you. I like to use the literal words “blind spot pass” and “unknown unknowns.” Giving it context on who you are and what you know is usually important for Claude to understand the best way to start collaborating with you.
“I'm working on adding a new auth provider but I know nothing about the auth modules in this codebase. Can you do a blind spot pass to help me figure out my relevant unknown unknowns and help me prompt you better.”
“I don’t know what color grading is but I need to grade this video. Can you teach me to understand my unknown unknowns about color grading, so that I can prompt better?”
When I’m working in an area with a lot of unknown knowns , involving criteria I only know to define when I see it, I like to ask Claude to brainstorm and prototype with me.
It’s extremely valuable to identify and verbalize unknown knowns early during prototyping, because finding them out during implementation can be (relatively) expensive. Small changes in a feature or spec can cause drastically different implementations in code, and it can be more difficult for your agent to revert previous changes.
For example, you may just want to see how a button added to a frame looks without having to wire up a backend route or maintaining additional state in the frontend.
Another example is visual design, which for me, is something that is difficult to articulate, but I know what I want when I see it. In these cases, I’ll ask for several design approaches to an artifact.
I also start almost every coding session with an exploration or brainstorming phase. This helps me start with intent to define the project’s scope. Claude often finds high-value approaches I would have missed, and sometimes misses the forest through the trees. Brainstorming prevents me from setting too narrow or too wide a scope.
"I want a dashboard for this data but I have no visual taste and don't know what's possible. Make me an HTML page with 4 wildly different design directions so I can react to them.”
“Before wiring anything up, make a single HTML file mocking the new editor toolbar with fake data. I want to react to the layout before you touch the real app."
"Here's my rough problem: users churn after onboarding. Search the codebase and brainstorm 10 places we could intervene, from cheapest to most ambitious. I'll tell you which ones resonate."
Once I’ve done sufficient brainstorming, I likely still have unknowns.
In this case, I ask Claude to interview me about any unknowns or ambiguities. When asking Claude to interview you, try and give it context about your problem to guide its questions.
"Interview me one question at a time about anything ambiguous, prioritize questions where my answer would change the architecture."
Sometimes you can’t describe what you want in detail. For example, you might not have the language or it might be so complicated that it would take you quite a while.
In this case, the best approach is a reference. While you can include diagrams, documentation or pictures, the absolute best reference is source code .
If you have a library that implements something in a certain way or a design component you really like, just point Fable at the folder and tell it what to look for, even if it’s in a different language. This provides Claude much richer detail around the markup and structure, compared to for example a screenshot.
"This Rust crate in vendor/rate-limiter implements the exact backoff behavior I want. Read it and reimplement the same semantics in our TypeScript API client."
When I think I’m ready to implement, I tend to ask Claude to put together an implementation plan for me to review. The plan focuses on the parts that might be most likely to change such as data models, type interfaces, or UX flows. This allows Claude to surface things I might actually need to alter.
"Write an implementation plan in HTML, but lead with the decisions I'm most likely to tweak with: data model changes, new type interfaces, and anything user-facing. Bury the mechanical refactoring at the bottom, I trust you on that part."
Once I am satisfied with my plan, I make a new session and pass any artifacts to the prompt. This gives Claude a fresh context window but with all of the information it compiled from your planning. For example, I might pass in a spec file and a prototype and ask an agent to implement it.
But the truth is that no matter how much planning you do, there are always unknown unknowns lurking. The agent may find during its work that it needs to take a different tack due to an edge case it found in the code.
I ask Claude Code to keep a temporary ‘implementation-notes.md’ (or .html) file where it keeps track of decisions it makes so we can learn for our next attempt.
"Keep an implementation-notes.md file. If you hit an edge case that forces you to deviate from the plan, pick the conservative option, log it under 'Deviations', and keep going."
One of the most important parts of shipping something is getting buy-in and approvals. Building pitch and explainer artifacts in the final document helps:
Accelerate understanding when reviewers start with the same unknowns you did
Accelerate approvals when experts want to see you accounted for the unknowns and common failure points they would have anticipated
"Package the prototype, the spec, and the implementation notes into a single doc I can drop in Slack to get buy-in. Lead with the demo GIF."
After a long working session, Claude might have accomplished a lot more than I realized. Reading the code diffs can only give me a light understanding of what happened, since much of the behavior will depend on existing code paths.
Asking Claude to quiz me about the change after giving me a bunch of context helps me understand what happens. I only merge after I pass the quiz perfectly.
“I want to make sure I understand everything that's happened in this change. Give me a HTML report on the changes for me to read and understand with context, intuition, what was done, etc. and a quiz at the bottom on the changes that I must pass.”
How this comes together: launching Fable
The launch video for Fable was edited end-to-end using Claude Code. This was a new domain for me and I’m by no means an expert.
So I started with what I did know. I knew that Claude could use code to edit videos and transcribe them, but I wasn’t sure if it was accurate enough. I then asked Claude to explain to me how transcription like Whisper worked, and whether I would be able to accurately cut out things like ums or large pauses using ffmpeg.
I wanted Claude to create a UI that was timed with the words I was saying, but wasn’t sure it was possible so I asked Claude to create a prototype video using Remotion and a transcription to see if it would work.
Finally, the video itself looked a bit muted, which I knew was the result of color grading but I didn’t really know what color grading was. My first pass attempt was to try and get Claude to do a few variations to pick, but I realized that I didn’t know what “good” looked like when it came to color grading. So instead, I asked Claude to teach me about color grading to discover my unknowns.
Matching the Map and Territory
The better models get, the more you can achieve with the right approach. When a long-horizon task comes back wrong, it's likely you need to spend more time defining your unknowns or creating an implementation plan that allows

[truncated]
