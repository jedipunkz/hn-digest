---
source: "https://swiftace.org/posts/large-scale-code-refactoring-with-ai"
hn_url: "https://news.ycombinator.com/item?id=49233633"
title: "Large-Scale Code Refactoring with AI"
article_title: "Large-Scale Code Refactoring with AI - SwiftAce"
author: "aakashns"
captured_at: "2026-08-09T18:25:03Z"
capture_tool: "hn-digest"
hn_id: 49233633
score: 1
comments: 0
posted_at: "2026-08-09T17:46:52Z"
tags:
  - hacker-news
  - translated
---

# Large-Scale Code Refactoring with AI

- HN: [49233633](https://news.ycombinator.com/item?id=49233633)
- Source: [swiftace.org](https://swiftace.org/posts/large-scale-code-refactoring-with-ai)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T17:46:52Z

## Translation

タイトル: AI を使用した大規模コード リファクタリング
記事のタイトル: AI を使用した大規模コード リファクタリング - SwiftAce
説明: 現在、Jovian を改良しており、Next.js 11 フロントエンド +... から移行しています。

記事本文:
AI を使用した大規模コード リファクタリング - SwiftAce ホーム GitHub AI を使用した大規模コード リファクタリング
現在、Jovian を改良し、Next.js 11 フロントエンド + Flask (Python) バックエンドからフルスタックの TanStack Start アプリに移行すると同時に、ユーザー インターフェイスを最新化し、いくつかの機能を変更しています。改良されたアプリは https://beta.jovian.com にデプロイされます。
私は、 src/routes フォルダーにすべてのルートが含まれ、 src/lib に utils.ts 、 utils.server.ts 、 Types.ts 、 config.ts (環境変数と定数用) などのファイルが含まれ、 src/db にデータベース スキーマとクエリが含まれる、かなりフラットな構造から始めました。
src/components フォルダーにはアプリ全体で使用される共通コンポーネントが含まれており、特定のルートで使用されるコンポーネントとサーバー関数を、ルート ファイルの隣のフォルダー -components および -functions 内に配置しました (- プレフィックスの付いたフォルダーはルートとして扱われません)。
コードベースが TypeScript の 15,000 行以上に増加するにつれて、src/lib および src/db 内のファイルは数百行、中には 1,000 行を超えるものまで増加したため、ヘルパー関数と定数を、それらを使用するルート ファイルの隣に配置して、よりモジュール的な構造にリファクタリングすることにしました。
この規模のリファクタリングには、ほんの数か月前であれば、1 週間とは言わないまでも、少なくとも数日はかかっていたでしょう。ただし、Claude Opus 5 は基本的にリファクタリング全体を 20 分程度でワンショットで完了し、費用は約 25 ドルでした。 5,700 行を超えるコードが移動され、数百のインポートが更新されました。
私が使用した完全なプロンプトは次のとおりです。
-components フォルダーと -functions フォルダーがあるのと同じように、utils.ts、utils.server.ts、types.ts などを含む -lib フォルダーも作成しましょう。
ここで、環境由来の変数をトップレベルの config.ts と config.server.ts に保持したいと考えていますが、その他のものはモジュール固有の -lib/constants.ts に入れることができます。

-lib/constants.server.ts ファイル (モジュール レベルではもう構成とは呼びません)。
基本的に、コードベース全体をよりモジュール化したいと考えています。また、@src/db/read.server.ts やその他の db クエリ ファイルも長すぎるように感じます。そこで、この CRUD ベースのファイル パターンをスキップして、代わりにデータベース クエリを含むモジュール レベルの -lib/db.server.ts ファイルを作成しましょう。
そして @src/db/schema.server.ts を src/lib/schema.server.ts に移動しましょう。
最後に、@src/lib/certificate.server.ts と @src/lib/markdown.server.ts を、適切な場所の -lib/utils.server.ts ファイルに移動します。これらのファイルはランダムで、適切ではないと思います。
さらに、作業モジュールごとにルートをより適切に整理するために、ルート グループを作成しましょう。 (manage) は、既存のディレクトリ構造が十分なカプセル化を行っていない場所での管理などのためのルート グループにすることができます。
その一方で、ファイルベースのルーティング設定を再確認し、ルートを適切に整理するための正しい選択 (例: ディレクトリかフラットか) を行ったかどうかを確認してみましょう (もちろん、実際のルート パスは変更しないでください)。
そしてもちろん、この新しい一連のガイドラインに従って CLAUDE.md も更新しましょう。ただし、あまり冗長になりすぎないようにしてください (サマジュダールからイシャラ カアフィハイまで)。
そして最後に、一部のコード コメントが長すぎると思うので、テキストの大きな壁を短くし、(絶対に必要な場合を除き) 最大 2 行にとどめ、そのガイドラインも確立しましょう。
さて、私が今考えているのはこれだけです。何か大きな間違いをしていないか、何か重要なことを見落としていないか、そしてこれが正しいアプローチであるかどうか教えてください。さあ行こう。
リファクタリングを開始する前に、いくつかの説明 (複数選択の質問として提示) を求められましたが、リファクタリングの途中で型の整理について気が変わりました。
それ以外にもありました

基本的に私の介入はなく (いくつかの小さなフォローアップタスクを与えました)、リファクタリング後、アプリケーションは完全に動作しました。 AI がなければ、私はこのリファクタリングを追求することはなかったかもしれません。プログラミングは永遠に変化しており、それと戦うことはできません。

## Original Extract

I&amp;#39;m currently revamping Jovian, migrating it from a Next.js 11 frontend +...

Large-Scale Code Refactoring with AI - SwiftAce Home GitHub Large-Scale Code Refactoring with AI
I'm currently revamping Jovian , migrating it from a Next.js 11 frontend + Flask (Python) backend to a full-stack TanStack Start app, while also modernizing the user interface and changing some features. The revamped app is deployed at https://beta.jovian.com .
I started out with a fairly flat structure where the src/routes folder contained all the routes and src/lib contained files like utils.ts , utils.server.ts , types.ts , config.ts (for environment variables & constants), and src/db contained the database schema & queries.
The folder src/components contained common components used throughout the app, and I placed components & server functions used in specific routes inside folders -components and -functions next to the route files (the - prefixed folders are not treated as routes).
As the codebase grew to 15k+ lines of TypeScript, the files in src/lib and src/db grew to hundreds of lines, some even over a thousand, so I decided to refactor it into a more modular structure, colocating helper functions and constants next to the route files that use them.
A refactor of this scale would have taken at least a few days, if not a week, just a few months ago. However, Claude Opus 5 essentially one-shotted the whole refactor in ~20 minutes and cost around $25. Over 5700 lines of code were moved & hundreds of imports were updated.
Here's the full prompt I used:
Just as we have -components and -functions folders, let's also have a -lib folder which will have utils.ts, utils.server.ts, types.ts etc.
Now I want to keep any environment-derived variables in the top level config.ts and config.server.ts, but the other stuff can go into module-specific -lib/constants.ts and -lib/constants.server.ts files (let's not call it config anymore at the module level).
Basically, I want the entire codebase to now become more modular. And I feel the @src/db/read.server.ts and other db query files are also getting too long. So, let's maybe skip this CRUD-based file pattern and instead have a module-level -lib/db.server.ts files which will have the database queries.
And let's move @src/db/schema.server.ts to src/lib/schema.server.ts.
Finally, also move @src/lib/certificate.server.ts and @src/lib/markdown.server.ts to -lib/utils.server.ts files in the right place. I think these files are kinda random and don't fit.
Further, to better organize the routes by modules of work, let's create route groups e.g. (manage) can be a route group for management etc. in places where the existing directory structure isn't doing sufficient encapsulation.
And while we're at it, let's also re-look at the file-based routing setup and see if we made the right choices (e.g. directory vs flat) to organize routes properly (of course, don't change any actual route paths).
And let's of course also update the CLAUDE.md according to this new set of guidelines, but don't get too overly verbose (samajdhar to ishara kaafi hai).
And finally, I think some of the code comments are way too long, so let's shorten the big walls of text and just keep 2 lines max (unless absolutely necessary) and also establish that guideline.
Okay, that's all I'm thinking right now, tell me if I'm making any big errors or overlooking something important, and if this is the right approach. Let's go.
It did ask me for a few clarifications (presented as multiple-choice questions) before starting the refactor, and I changed my mind about organizing types midway through the refactor:
Other than that, there was basically no intervention on my part (I did give it a couple of minor follow-up tasks), and the application worked perfectly after the refactor. Without AI, I might have never pursued this refactor. Programming has changed forever, there is no fighting it.
