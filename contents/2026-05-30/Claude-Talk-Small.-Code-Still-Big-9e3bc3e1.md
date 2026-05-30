---
source: "https://spatie.be/blog/claude-talk-small-code-still-big"
hn_url: "https://news.ycombinator.com/item?id=48327064"
title: "Claude Talk Small. Code Still Big"
article_title: "Claude Talk Small. Code Still Big. | Spatie"
author: "speckx"
captured_at: "2026-05-30T11:41:20Z"
capture_tool: "hn-digest"
hn_id: 48327064
score: 2
comments: 0
posted_at: "2026-05-29T18:13:35Z"
tags:
  - hacker-news
  - translated
---

# Claude Talk Small. Code Still Big

- HN: [48327064](https://news.ycombinator.com/item?id=48327064)
- Source: [spatie.be](https://spatie.be/blog/claude-talk-small-code-still-big)
- Score: 2
- Comments: 0
- Posted: 2026-05-29T18:13:35Z

## Translation

タイトル: クロード・トーク・スモール。コードはまだ大きい
記事タイトル: クロード・トーク・スモール。コードはまだ大きい。 |スパティ
説明: AI コーディング アシスタントに穴居人のような短い文で応答させる穴居人スキルの短い実験。いくつかのトークンを節約できましたが、予想よりも少なくなりました。

記事本文:
クロード・トーク・スモール。コードはまだ大きい。 |スパティ
= 720、
ossOpen: false、
isMobile: window.innerWidth = 720; isMobile = window.innerWidth
スパティエ
メニュー
サービス
オープンソース
私たちの理念
パッケージ
ポストカード
ガイドライン
ドキュメント
製品
ブログ
について
ログイン
私たちと一緒に働きましょう
2026 年 5 月 6 日
クロード・トーク・スモール。コードはまだ大きい。
マルセリ
code]:text-14 [&_:not(pre)>code]:text-oss-royal-blue [&_:not(pre)>code]:p-[2px] [&_>pre]:bg-oss-gray-light [&_p>code]:bg-oss-gray-light [&_>pre]:p-6 [&_>:not(pre)>code]:p-[2px] [&_p:has(video)]:mb-8 [&_p:has(img)]:mb-8 [&_>iframe]:rounded-[0.5em] sm:[&_:not(pre)>code]:text-[15.5px] sm:[&_p:has(img)]:-mx-12 [&_p>img]:rounded-[0.5em] [&_p>img]:overflow-hidden sm:[&_p:has(video)]:-mx-12 [&_p>video]:rounded-[0.5em] [&_p>video]:overflow-hidden [&_p>video]:w-full md:[&_>.insights-list-item]:-mx-12 md:[&_>.insights-list-item]:px-12 md:[&_>.insights-list-item]:my-8 md:[&_>pre]:-mx-12 md:[&_>pre]:px-12 md:[&_>pre]:-mx-12 md:[&_>pre]:px-12">
AIコーディングアシスタントはよく話します。それが役に立つこともあります。何を変更したのか、なぜ変更したのか、どのファイルが変更されたのか、まだ注意が必要なものは何か、途中でどのような仮定を行ったのかについて説明します。場合によっては、小さな変更に関する単なる言葉の多さもあります。そこで出会ったのが「穴居人」というスキル。
アイデアはシンプルです。アシスタントに穴居人のような短い文で話させるというものです。説明が減り、長い段落が減り、丁寧な回り道が減ります。理論的には、これによりアシスタントが自分の動作を説明する時間が短縮されるため、トークンが節約されるはずです。
ファイルを検査します。
私からの変更リクエストです。
私はテストを実行します。
テストは合格です。
実装計画、リスクなどに関する完全な段落の代わりに...
AI コーディング アシスタントを一日中使用していると、コード周りの通信にどれだけのトークンが費やされているかに気づき始めます。ああ、そうだね

コミュニケーションが役立つ時期です。何が変更されたのか、なぜ何かが失敗したのか、まだ注意が必要なものが説明されています。
しかし、時にはそれがたくさんあることもあります。特に、コードベースをすでに知っていて、何が必要かをすでに知っていて、その作業をアシスタントに任せるだけで十分な場合は特にそうです。そのようなときは、短いコミュニケーション スタイルが便利だと思われます。話すことを減らし、実行することを増やします。
そこで、Caveman のようなスキルによって、私の AI エージェント (最近ではクロードより Codex の方が多いと思いますが、Codex (gpt 5.5 モデル) はスリーパーであり、優れているとは言わないまでも非常に似たような出力で ROI がはるかに優れています) を実際の使用においてより安価に、またはより効率的にできるかどうかを確認したいと思いました。
はい。しかし、私が期待したり望んでいたほどではありませんでした。
理由は非常に単純です。コーディングに Claude を使用している場合、ほとんどのトークンが常に説明に含まれるわけではありません。その多くは、ファイルの読み取り、コンテキストの理解、コードの生成、変更の適用、結果の確認にあります。
したがって、たとえアシスタントが次のように言ったとしても、
問題を見つけてリクエスト クラスを更新し、ペイロードが API に正しく渡されるようにしました。
実際のコードはまだ記述する必要があります。ファイルはまだ読み取る必要があります。応答には、差分、関数名、名前空間、インポート、およびテストが含まれています。言い換えれば、Caveman は会話を短くしますが、コーディングはただ話すだけではありません。それが私が過小評価していた主な点でした。トークンの節約がもっと劇的になることを期待していましたが、高価な部分は多くの場合、作業そのものです。 Caveman のドキュメントには、トークンの約 65% を節約できると書かれており、何をしているかによってはそれが真実であると思いますが、コーディングタスクに関する私の経験では、タスクに応じて 5 ～ 20% の間である可能性が高かったです。
はい、そう思います。私の AI エージェントが何をしたかについて複数の文で説明する必要はありません。いずれにしてもコードを確認する必要があるので、このような短い説明が必要です

低いほうが、より早くコードを読み始めることができるので、私にとっては良いです。
DTO を更新します。
不足している null 許容フィールドを追加します。
私はパイントを経営しています。
大丈夫です。
また、出力をレビューするのがなんだか面白くなりました。コーディングアシスタントがまるで火事を発見したかのようにLaravelのリファクタリングについて説明するのを見るのは、とても楽しいものです。
より複雑なタスクや計画の場合 (計画モードの代わりに、Matt Pocock によるグリルミー スキルを強くお勧めします)、私は依然として通常の説明を好みます。
何かがコードベースの複数の部分に触れたり、動作が変わったり、トレードオフがある場合には、アシスタントにそれ自体を適切に説明してもらいたいと思っています。短い穴居人風の更新は面白いですが、特定のアプローチが選択された理由を理解する必要がある場合には、必ずしも十分ではありません。
不要な単語を減らすことと、有用なコンテキストを削除することの間には違いもあります。追加の説明は、間違った実装になる前に、間違った仮定を見つけるのにまさに役立つ場合があります。したがって、私はこのスキルをすべてに使用するつもりはありません。
Caveman スキルは、コミュニケーション スタイルの小さな変更が AI アシスタントとの作業の感覚にどのように影響するかを示す好例です。
いくつかのトークンは節約できましたが、私が期待していたほどの大規模なものではありませんでした。そしてそれは今では当然のことです。コーディング タスクでは、アシスタントはコード周辺の文だけでなく、実際のコードとコンテキストに多くのトークンを費やします。
これにより、小さなタスクが軽く感じられ、アシスタントの出力をレビューするのは本当に面白く、スキルが必ずしも大規模なプロジェクト固有の自動化である必要はないことがわかりました。場合によっては、スキルによってアシスタントの動作が変わるだけで、それだけでも役立つことがあります。
ちょっとした作業や楽しみのために手元に置いておくとよいでしょうか?はい。
Piper の紹介: パイプ演算子を使用した配列と文字列の操作
パイプオペレーターは素晴らしい追加でした

PHP8.5では。これは、スタンドアロン関数の構成可能性を備えたラッパー オブジェクトの人間工学をもたらします。残念ながら、PHP の標準の配列および文字列操作関数は API の一貫性について正確には知られていないため、パイプ演算子は使いにくくなっています。 Piper は、標準ライブラリをラップしてパイプ演算子との互換性を持たせる試みです。
クロードのスキルをクライアントプロジェクトに活かす
実際のクライアント プロジェクトで AI スキルを使用して、反復的な Laravel 開発をスピードアップする方法。 Saloon リクエスト クラスや DTO から Livewire コンポーネントやカスタム登録まで。
Mailcoach を利用した強力な電子メール マーケティング ツールを簡単に
成長し、つながり、変換します。
購読する
私たちは Spatie TLDR を 3 か月ごとに送信します。製品のアップデート、舞台裏で何が起こっているか、
そして興味深いリンク。また、新しい情報がある場合には、不定期に製品の更新情報も送信します。
このフォームを送信すると、当社のプライバシー ポリシーに同意したことになります。
私たちは単なる開発者ではなく、アドバイザーおよびアーキテクトとして活動します。私たちもあなたと同じようにあなたのプロジェクトを誇りに思いたいと思っています。 Laravel でのオーダーメイドの Web 開発が私たちの得意分野です。
info@spatie.be
+32 3 292 56 79
GitHub
私たちのオフィスは現在閉鎖されています。代わりにメールでお問い合わせください
一言で言えば、私たちはあなたのプロジェクトの重要な部分になりたいと思っています。

## Original Extract

A short experiment with the Caveman skill, which makes AI coding assistants respond in short caveman-like sentences. It did save some tokens, but less than expected.

Claude Talk Small. Code Still Big. | Spatie
= 720,
ossOpen: false,
isMobile: window.innerWidth = 720; isMobile = window.innerWidth
SPATIE
Menu
Services
Open source
Our philosophy
Packages
Postcards
Guidelines
Documentation
Products
Blog
About
Login
Work with us
May 06, 2026
Claude Talk Small. Code Still Big.
Marceli
code]:text-14 [&_:not(pre)>code]:text-oss-royal-blue [&_:not(pre)>code]:p-[2px] [&_>pre]:bg-oss-gray-light [&_p>code]:bg-oss-gray-light [&_>pre]:p-6 [&_>:not(pre)>code]:p-[2px] [&_p:has(video)]:mb-8 [&_p:has(img)]:mb-8 [&_>iframe]:rounded-[0.5em] sm:[&_:not(pre)>code]:text-[15.5px] sm:[&_p:has(img)]:-mx-12 [&_p>img]:rounded-[0.5em] [&_p>img]:overflow-hidden sm:[&_p:has(video)]:-mx-12 [&_p>video]:rounded-[0.5em] [&_p>video]:overflow-hidden [&_p>video]:w-full md:[&_>.insights-list-item]:-mx-12 md:[&_>.insights-list-item]:px-12 md:[&_>.insights-list-item]:my-8 md:[&_>pre]:-mx-12 md:[&_>pre]:px-12 md:[&_>pre]:-mx-12 md:[&_>pre]:px-12">
AI coding assistants talk a lot. Sometimes that is useful. They explain what they changed, why they changed it, which files were touched, what still needs attention, and which assumptions they made along the way. Sometimes it is also just a lot of words around a small change. Then I came across a skill called “Caveman” .
The idea is simple: make the assistant talk in short, caveman-like sentences. Less explanation, fewer long paragraphs, fewer polite detours. In theory, this should save tokens because the assistant spends less time explaining what it is doing.
Me inspect files.
Me change request.
Me run tests.
Tests pass.
Instead of a full paragraph about the implementation plan, risks, ...
When using AI coding assistants all day, you start noticing how many tokens are spent on communication around the code. A lot of the time that communication is useful. It explains what changed, why something failed, or what still needs attention.
But sometimes it is also a lot. Especially when you already know the codebase, already know what you want, and just need the assistant to do the work. In those moments, a shorter communication style sounds useful. Less talking, more doing.
So I wanted to see if a skill like Caveman could make my a.i agent (lately more Codex then Claude i think Codex (gpt 5.5 model) is a sleeper and has much better ROI with very similar if not better output) cheaper or more efficient in real usage.
Yes. But not as much as I expected or wanted to.
The reason is pretty simple: when you're using Claude for coding, most of the tokens are not always in the explanation. A lot of them are in reading files, understanding context, generating code, applying changes, and checking results.
So even if the assistant says:
I found the issue and updated the request class so that the payload is now correctly passed to the API.
The actual code still has to be written. The files still have to be read. The response still includes diffs, function names, namespaces, imports, and tests. In other words: Caveman makes the talking shorter, but coding is not just talking. That was the main thing I underestimated. I expected the token savings to be more dramatic, but the expensive part is often the work itself. The Caveman documentation says you can save around 65% of tokens and I believe that can be true depending on what you are doing but in my experience with coding tasks it was more likely between 5-20% depending on the task.
Yes I think so. I don't need the multi sentence explanation of what my a.i agent did. I will need to go over the code anyway so having a short description like this below is for me better since i can jump faster to reading the code.
Me update DTO.
Me add missing nullable field.
Me run Pint.
All good.
It also made reviewing the output kind of funny. There is something very entertaining about seeing a coding assistant explain a Laravel refactor like it just discovered fire.
For more complicated tasks and planning (instead of plan mode I highly recommend the grill-me skill by Matt Pocock ), I still prefer normal explanations.
If something touches multiple parts of the codebase, changes behavior, or has trade-offs, then I do want the assistant to explain itself properly. Short caveman-style updates are funny, but they are not always enough when you need to understand why a certain approach was chosen.
There is also a difference between reducing unnecessary words and removing useful context. Sometimes the extra explanation is exactly what helps you catch a wrong assumption before it turns into a bad implementation. So I would not use this skill for everything.
The Caveman skill is a cool example of how small changes in communication style can affect the feeling of working with an AI assistant.
It did save some tokens, just not in the huge way I expected. And that makes sense now. In coding tasks, the assistant spends a lot of tokens on actual code and context, not just on sentences around the code.
It made small tasks feel lighter, reviewing the assistant's output was genuinely funny, and it showed that skills do not always need to be big project-specific automations. Sometimes a skill can simply change the way the assistant behaves, and that alone can be useful.
Would I keep it around for quick tasks and for the fun of it? Yes.
Introducing Piper: array and string manipulation with the pipe operator
The pipe operator was a great addition in PHP 8.5. It brings the ergonomics of wrapper objects with the composability of standalone functions. Unfortunately, the standard array & string manipulation functions in PHP aren't exactly known for their API consistency, which makes the pipe operator awkward to use. Piper is an attempt to wrap the standard library to make it compatible with the pipe operator.
Utilizing Claude Skills in client projects
How I use AI skills in a real client project to speed up repetitive Laravel development. From Saloon request classes and DTOs to Livewire components and custom registrations.
Powered by Mailcoach , powerful email marketing tools to effortlessly
grow, connect and convert.
Subscribe
We send the Spatie TLDR every three-ish months: product updates, what's going on behind the scenes,
and interesting links. We also send occasional product updates when we have something new to share.
By submitting this form, you acknowledge our Privacy Policy .
We act as advisors and architects, not just developers. We want to be as proud of your project as you are. Tailor-made web development in Laravel is what we do best.
info@spatie.be
+32 3 292 56 79
GitHub
Our office is closed now, email us instead
In short: we'd like to be a substantial part of your project.
