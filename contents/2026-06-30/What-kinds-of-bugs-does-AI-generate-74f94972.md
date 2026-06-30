---
source: "https://blog.detail.dev/posts/ai-bug-types/"
hn_url: "https://news.ycombinator.com/item?id=48735435"
title: "What kinds of bugs does AI generate?"
article_title: "What kinds of bugs does AI generate? | Detail"
author: "drob"
captured_at: "2026-06-30T17:35:29Z"
capture_tool: "hn-digest"
hn_id: 48735435
score: 4
comments: 2
posted_at: "2026-06-30T16:48:03Z"
tags:
  - hacker-news
  - translated
---

# What kinds of bugs does AI generate?

- HN: [48735435](https://news.ycombinator.com/item?id=48735435)
- Source: [blog.detail.dev](https://blog.detail.dev/posts/ai-bug-types/)
- Score: 4
- Comments: 2
- Posted: 2026-06-30T16:48:03Z

## Translation

タイトル：AIはどんなバグを生むのか？
記事タイトル：AIはどんなバグを生み出すのか？ |詳細
説明: AI がどのようなバグを作成しているかを確認するために、数千のコードベースから 1,000 のバグをクラスタリングしました。

記事本文:
AIはどのようなバグを生成するのでしょうか？ |詳細
詳細
無料でお試しください
ログイン
2026-06-28
・ピリー・タカラ AIはどのようなバグを生成するのでしょうか？
AI がどのようなバグを作成しているかを確認するために、数千のコードベースから 1,000 のバグをクラスタリングしました。
エージェントはどのようなバグを生成していますか?バイブコード化されたソフトウェアはどのように壊れるのでしょうか?自動コードレビューではどこが不足しているのでしょうか?それぞれのバグはランダムで 1 回限りのものなのでしょうか、それとも明確なパターンを特徴付けることができるのでしょうか?さまざまな会社のエンジニアが同じいくつかのバグを繰り返し作成しているのでしょうか?
これらの質問に答えるために、99 の顧客のコードベースから 1,000 のバグの例を抽出し、それらをクラスター化しました。
コードベース間で同じ間違いが発生する
私たちは 1,000 個のバグから開始し、それぞれが詳細によってフラグが付けられ、その後お客様によって修正されました。私たちは各バグを言語に依存せず、企業に依存しないバグの特定のメカニズムの説明に書き直してから、それらの説明を埋め込んでクラスター化し、バグがどのように失敗するかによってグループ化しました。
私たちが最初に気づいたのは、バグの 70% が少数のクラスターに分類されているということでした。 21 のメカニズムの短いリストで、これらのバグの 70% が説明されています。これらの同じ 21 の間違いが、さまざまなソフトウェア製品で何百回も繰り返されています。
「いつもの容疑者」がリストの非常に下の方にいたことに私たちは驚きました。たとえば、Null ポインター例外はかつて「10 億ドルの間違い」と呼ばれていました。
, しかし、それらはコーパスの 1.7% しか占めていませんでした。
これらのクラスターを低次元空間に投影すると、一般的なバグの空間を非常に大まかにマッピングできます。最も重要な 2 つの軸は「信頼性」です。バグはエンド ユーザーが情報に不正にアクセスすることに関連していますか? – そして「競合状態」 – バグが表面化するためには、明らかではない一連のイベントが必要ですか?
競合状態 → 信頼性 → 低 高 10 11 12 13 14 15 16 17 18 19 20 21

22 1 2 3 4 5 6 7 8 9 TOCTOU のみ 高-高 高信頼性 高人種性 両方 — TOCTOU なし バグの可読性
これらのカテゴリに共通しているのは、機能を個別にスモーク テストした場合には気付かないバグであるということです。コードを作成するエージェントや新機能をテストするエンジニアにとって、これらはすぐには「判読」できません。
極端な例としては、認証チェックの欠落が挙げられます。これは、コーパス内で最も一般的なカテゴリです。これらのバグは単なるバグではなく、正真正銘の脆弱性ですが、ブロックすべきアクセスを許可することで失敗し、一般的なユーザーはそれに気づきません。また、機能するはずのない認証ケースをテストしようと特に考えていない限り、エンジニアも機能を構築することはありません。これらのバグは、後でハッカーに発見されたとき、または SaaS 製品の権限が適切に機能せず、多くの顧客の信頼が一度に失われたことに顧客が気づいたときにのみ目に見えます。
競合状態、データベース状態の不完全な更新、アプリのさまざまな部分にわたる不一致も同様に判読できません。競合状態は再現が難しい場合があり、ほとんどの手動 QA はシングルプレイヤーであるため、本番環境に移行します。一貫性のない状態処理を検出するには、エンジニアが取り組んでいる 1 つの機能に特化したものではなく、広い表面積にわたる慎重なテストが必要であるため、同様に見落としがちです。
振り返ってみると、現在のソフトウェアの製造方法を考えると、これらのバグは驚くべきことではありません。バグがマージされるには、バグを共同作成したエンジニアと AI を通過し、次にリンターと関連テストを通過し、さらにコード レビューを通過する必要があります。これらのバグの多くは運用環境でも存続し、ユーザーをすり抜けたり、再現が困難すぎたりして、最終的に Detail がバグを検出し、その後顧客が修正するまで続きました。
アノを入れて

他の方法: 現在運用環境に到達しているバグはクラッシュすることはほとんどありません。むしろ、それらは、一度見ると明らかに壊れているアプリケーションの動作ですが、探さないと見つからない可能性があります。
本番稼働前に捕らえられる スモークテスト、リンター、テスト、レビュー 本番稼働のユーザーレポート、サポートチケット、インシデント、違反 より読みやすいものは大声で失敗する 読みにくいものは静かに失敗する フローの途中でアプリがクラッシュする フォームが送信されない Safari Race でサインアップが失敗する カードに二重請求が行われる 古い価格が提供される 認証チェックがスキップされる 判読できないバグを判読できるようにする
これらのカテゴリのバグを克服できます。小規模なアプリケーションの場合、手動 QA を行うときに何を確認すればよいかを知っておくだけでも役立ちます。 (または、代理店に QA を依頼する場合も同様です。) 複雑な製品全体でこのような微妙な欠陥を防ぐには、次の 2 つのことを行う必要があります。
現在のエージェントがこれらのバグをより目に見えるようにします。これはコードのアーキテクチャの問題です。ソフトウェアを設計する際に考慮すべき最も重要な問題の 1 つは、ソフトウェアが正しく動作しているかどうかをエージェントにどのようにして明らかにするかということです。
バグの発見に優れたエージェントを構築します。私たちはこのようなパターンを何千ものコードベースから見つけ出し、それらを使用して、エージェント、特にコード レビューが常に見逃しているバグの発見に非常に優れたエージェントを構築します。この問題にご興味がございましたら、ぜひお話しさせていただきます。
バグをより読みやすくするアーキテクチャ パターンの例は、「デフォルトで拒否」です。つまり、明示的な許可がないことを「アクセスなし」を意味し、忘れられた承認チェックによってリクエストが許可されるのではなくブロックされるようにします。このようにバックエンド エンドポイントを設計すると、静かにデータを漏洩するバグは、通常のテストで検出できる代わりに大声で失敗します。

打つ。
運用前に捕捉 スモーク テスト、リンター、テスト、レビュー REACHES 運用のユーザー レポート、サポート チケット、インシデント、違反 デフォルトで拒否するように変更 読みやすくなると大声で失敗する 読みにくくなると静かに失敗する 認証が見つからない → 403 認証が見つからない → 200 試してみる
自分のコードベースに何が隠されているかを確認したい場合、最初のスキャンは無料です。ここからサインアップしてください。
サンプル。過去 2 か月半の間に 99 人の顧客から寄せられた 1,000 件のバグ。単一の顧客が混合を歪めないように 1 社あたり 5% に制限され、匿名化されています。すべてのバグは詳細によって報告され、顧客のエンジニアによって修正されました。サンプル内で最も一般的な言語は TypeScript でした。ほぼすべての顧客が少なくとも 1 つの自動コード レビュー システムを使用しており、複数の自動コード レビュー システムを使用している顧客もいます。
方法。 Claude Haiku 4.5 では、各バグを、コード、企業、言語を使用しないメカニズムの説明に書き直しました。これらの記述を OpenAI の text-embedding-3-large と Google の gemini-embedding-001 の 2 つのモデルで埋め込み、UMAP で次元を削減しました。
、HDBSCAN でクラスター化されています
これにより、すべてのポイントが強制的にクラスターに組み込まれるのではなく、外れ値がグループ化されなくなります。元のバグレポートをクラスタリングすると、メカニズムではなく言語と会社が復元されます。 Claude Opus 4.8 は、バグのサンプルを読んだ後、各クラスターに名前を付けました。
検証。正規化された相互情報量
クラスターと会社、リポジトリ、または言語の間の比率は低く、言語では約 0.14 でした。 2 つの埋め込みモデルは、実質的に同じ構造を復元しました。各クラスターからブラインド サンプルを読み取って、ラベルが適合し、グループが単一のメカニズムとしてまとまっていることを確認し、最大のクラスターを手作業でレビューしました。
無料でお試しください
エンジニアに相談する
ホーム

## Original Extract

We clustered 1,000 bugs from thousands of codebases to see what bugs AI is creating.

What kinds of bugs does AI generate? | Detail
Detail
Try For Free
Log in
2026-06-28
· Pyry Takala What kinds of bugs does AI generate?
We clustered 1,000 bugs from thousands of codebases to see what bugs AI is creating.
What bugs are agents generating? How does vibecoded software break? Where is automated code review coming up short? Is each bug a random oneoff, or can we characterize clear patterns? Are engineers at different companies creating the same few bugs over and over?
To answer these questions, we took 1,000 example bugs from 99 customer codebases and clustered them.
Same mistakes across codebases
We started with 1,000 bugs, each flagged by Detail and subsequently fixed by the customer. We rewrote each bug into a language-agnostic and company-free description of the specific mechanism of the bug, then embedded and clustered those descriptions to group bugs by how they fail.
The first thing we noticed was that 70% of the bugs fell into a small number of clusters. A short list of 21 mechanisms described 70% of these bugs. These same 21 mistakes are recurring hundreds of times across a diverse set of software products.
We were surprised to find that the “usual suspects” were very low on the list. For example, null pointer exceptions used to be called the “billion-dollar mistake”
, but they made up only 1.7% of the corpus.
If we project these clusters into a low-dimensional space, we can very coarsely map out the space of common bugs. The two most important axes are “authyness” – do the bugs relate to end users improperly accessing information? – and “race-condition-ness” – do the bugs require a non-obvious sequence of events in order to surface?
race-condition-ness → authyness → low high 10 11 12 13 14 15 16 17 18 19 20 21 22 1 2 3 4 5 6 7 8 9 TOCTOU only high–high high authyness high race-ness both — TOCTOU neither Bug legibility
What these categories all have in common is that they’re bugs you wouldn’t notice when smoke-testing features in isolation. They aren’t immediately “legible” to the agent writing the code or the engineer testing out a new feature.
An extreme example would be a missing auth check, which is by far the most common category in the corpus. These bugs are bona fide vulnerabilities, not just bugs, but they fail by allowing access they should block, a typical user wouldn’t notice them. Nor would an engineer building a feature, unless they specifically thought to test an auth case that shouldn’t work. These bugs are only visible later, when a hacker finds them, or when a customer notices that the permissions in a SaaS product don’t work properly and a lot of customer trust is lost all at once.
Race conditions, incomplete updates to database state, and inconsistencies across different parts of an app are all similarly illegible. Race conditions can be tricky to reproduce, and most manual QA is single player, so they slip through to production. Detecting inconsistent state-handling requires careful testing across broad surface area, not specific to any one feature an engineer is working on, so it’s similarly easy to miss.
In retrospect, these bugs shouldn’t surprise us, given how software is produced right now. For a bug to get merged, it needs to slip past the engineer and the AI that co-created it, then past linters and any relevant tests, and then past code review. Many of these bugs persisted in production too, either slipping past users or too hard to reproduce, until Detail finally detected them and a customer subsequently fixed them.
Put another way: the bugs that reach production right now are rarely crashes; rather, they’re application behavior that’s clearly broken once you see it, but you might not see it if you don’t look for it.
CAUGHT BEFORE PRODUCTION smoke tests · linters · tests · review REACHES PRODUCTION user reports · support tickets · incidents · breaches MORE LEGIBLE fails loudly LESS LEGIBLE fails silently App crashes mid-flow Form won't submit Signup fails on Safari Race double-charges a card Serves a stale price Auth check skipped Making illegible bugs legible
We can beat these categories of bugs. For a small application, it can help just to know what to look for when doing manual QA. (Or, when asking an agent to do our QA for us.) If we’re going to prevent these kinds of subtler defects across a complicated product, we need to do two things:
Make these bugs more visible to the agents we have today. This is a question of code architecture. When designing a piece of software, one of the most important questions to ask is how to make it obvious to agents whether or not it’s working correctly.
Build agents that are better at finding bugs. We find patterns like these across thousands of codebases, and we use them to build an agent that’s extremely good at finding the bugs that agents, and especially code review, consistently miss. If you’re interested in this problem, we would love to talk to you.
An example of an architecture pattern that makes bugs much more legible is “denial by default”: make the absence of an explicit grant mean “no access” so that a forgotten authorization check blocks a request instead of allowing it. If you architect your backend endpoints this way, bugs that would leak data silently will instead fail loudly, where ordinary testing can catch it.
CAUGHT BEFORE PRODUCTION smoke tests · linters · tests · review REACHES PRODUCTION user reports · support tickets · incidents · breaches Change to deny by default MORE LEGIBLE fails loudly LESS LEGIBLE fails silently Missing auth → 403 Missing auth → 200 Try it
If you’d like to see what’s hiding in your own codebase, the first scan is free. Sign up here.
Sample. 1,000 bugs from 99 customers over the last two and a half months, capped at 5% per company so no single customer skews the mix, and anonymized. Every bug was flagged by Detail and then fixed by the customer’s engineers. The most common language in the sample was TypeScript. Almost all customers use at least one automated code review system, and some use multiple.
Method. Claude Haiku 4.5 rewrote each bug into a code-free, company-free, language-free description of its mechanism. We embedded those descriptions with two models, OpenAI’s text-embedding-3-large and Google’s gemini-embedding-001 , reduced dimensionality with UMAP
, and clustered with HDBSCAN
, which leaves outliers ungrouped rather than forcing every point into a cluster. Clustering the original bug reports instead recovers language and company, not mechanism. Claude Opus 4.8 then named each cluster after reading a sample of its bugs.
Validation. Normalized mutual information
between the clusters and company, repository, or language was low, around 0.14 for language. The two embedding models recovered substantially the same structure. We read blind samples from each cluster to check that the labels fit and the groups hold together as single mechanisms, and reviewed the largest clusters by hand.
Try For Free
Talk to an Engineer
Home
