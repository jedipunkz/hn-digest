---
source: "https://bitnoise.pl/insights/rebuilding-bitnoise-webiste-with-claude-code-and-figma-mcp"
hn_url: "https://news.ycombinator.com/item?id=48633628"
title: "AI compressed the 15% of our website rebuild that was typing, not the other 85%"
article_title: "420 hours → 78 hours: rebuilding bitnoise.pl with Claude Code and Figma MCP"
author: "lroth"
captured_at: "2026-06-22T18:22:20Z"
capture_tool: "hn-digest"
hn_id: 48633628
score: 1
comments: 0
posted_at: "2026-06-22T18:02:16Z"
tags:
  - hacker-news
  - translated
---

# AI compressed the 15% of our website rebuild that was typing, not the other 85%

- HN: [48633628](https://news.ycombinator.com/item?id=48633628)
- Source: [bitnoise.pl](https://bitnoise.pl/insights/rebuilding-bitnoise-webiste-with-claude-code-and-figma-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T18:02:16Z

## Translation

タイトル: AI は、入力していたウェブサイト再構築の 15% を圧縮しましたが、残りの 85% は圧縮しませんでした
記事タイトル: 420 時間 → 78 時間: クロード コードと Figma MCP を使用して bitnoise.pl を再構築する
説明: 420 時間 → 78 時間: bitnoise.pl を再構築
クロード・コードとFigma MCP

記事本文:
420 時間 → 78 時間: Claude Code と Figma MCP を使用した bitnoise.pl の再構築 Bitnoise _ サービス 業界 エンゲージメント ケーススタディ 当社の製品 洞察について お問い合わせ メニュー すべての洞察に戻る 公開日 2026 年 6 月 9 日
420 時間 → 78 時間: bitnoise.pl を再構築
クロード・コードとFigma MCP
スタートに戻る 戻る 目次 速度がどこから来ているか、どこから来ていないのか、そして機械には決して渡さない作業のどの部分についての正直な説明。
あなたが読んでいる Web サイトには、人間が入力したコードは 1 行もありません。
すべての .tsx ファイル、すべてのアニメーション カーブ、すべての Strapi フェッチャー、すべての SEO スキャフォールディング (120 個の React コンポーネントにわたる 24,296 行) は、Claude Code セッションから生まれました。私たちは手書きで書きました。すべて手作業でレビューしました。
それでも8週間かかりました。
2 番目の数字について話したいと思います。業界が現在言い続けているストーリーは、「AI がコードを書き、開発者はコーヒーを飲みに行く」というものです。ここで起こったことはそうではありません。何が起こったのかというと、プロジェクトの 15% の入力作業が約 5 倍に圧縮され、残りの 85% (戦略、コピー、デザイン、モーション プロトタイピング、コード レビュー、QA、SEO) が常に必要としていた注目を集めるようになりました。キーボードがボトルネックになることはなくなりました。他のものはすべて高価なままでした。
この投稿は正直な会計です。数値は git log から直接取得されます。比較ポイントは、以前の bitnoise.pl ビルドの時間レポートから得られます。このスタックを検討しているチームにとって興味深い質問は、「モデルはコード化できるか?」ではないと考えているため、両方を共有しています。 — もちろんそれは可能ですが — しかし、「モデルの出力を出荷する価値のあるものにするために、チームの残りのメンバーは何をしなければならないでしょうか?」
開発部分だけで開発時間が最大 5.4 倍短縮され、n

新しいサイトには、より多くのページ、より多くのアニメーション、ブログとケーススタディのための Strapi の完全な統合、および古いサイトにはなかった SEO/GEO スキャフォールディングが追加されました。
2026 年 4 月 8 日 (最初のコミット) から 2026 年 6 月 2 日 (起動準備完了) までの git ログから監査されたリポジトリ自体:
カレンダーの期間: 56 日 (約 8 週間)。
アクティブな開発日数: 25 日。残りの 31 日は、設計、コピー、レビュー、QA、決定待ちで、これらの作業はすべてモデルでは実行されません。
合計コミット数: 181。すべてのコミットはクロード コードによって生成され、フレッシュ コンテキスト パスでクロード コードによってレビューされ、人間によって承認されます。
追加/削除された行: 857 件のファイル変更で、47,173 行が挿入、5,676 行が削除されました。
現在のコードベースの形状: 120 の .tsx ファイル、47 の .ts ファイル、および 1 つの .css ファイル (Tailwind v4 — はい、1 つのスタイルシート) にわたる 24,296 LOC。
集中的なコーディング時間: 53 の個別セッションで最大 78 時間。 (方法: セッション内のコミット間の実測間隔を合計します。90 分を超えるギャップがあると、新しいセッションが開始されます。これは請求レベルの数値ではありませんが、正しい形状です。)
平均コミット: 約 26 分の集中力。ほとんどの差分は 1 分以内に読めるほど小さいものです。それは故意です。
Claude Code — コード生成とコード レビュー パス用。
Figma MCP — デザイン システム、コンポーネント、フレーム、変数、モーション プロトタイプは Figma 内に存在し、モデルはそれらを直接読み取ります。 「スクリーンショットからの実装」という推測は必要ありません。
スタック: TanStack Start + Vite 7 + React 19 + Motion 12 + Tailwind v4。 CMSとしてのStrapi。お問い合わせフォーム用に再送信します。
52,800 の変更された行 × 70 文字 ÷ トークンごとの 4 文字 = 出荷されたコードの出力トークンは 925k です。
コンテキストの読み取り、再試行、レビュー パスを含めたエンドツーエンドの使用量は、およそ 1,000 ～ 2,000 万トークンでした。
実際には Claude Max プランでこの費用を支払いましたが、PR を使用して従量課金制の API 料金を自分のチームに見積もっている場合は、

Opus と Sonnet の組み合わせに応じて、同等のコストは約 90 ～ 180 ドルになります。Opus を多く使用し、効果的なキャッシュを備えたものがその最上位に位置し、Sonnet 寄りの組み合わせが最下位に位置します。予算が 200 ドルであれば、ほぼ確実にその予算内で収まります。
420 時間のタイピング作業が 78 時間の判断作業に置き換えられました。判断は外注できない部分です。
このアプローチ全体が正当化された瞬間を 1 つだけ挙げるなら、それは 2026 年 5 月 20 日になるでしょう。この日は、ブログとケーススタディのセクション全体を古いサイトから新しいスタックに移植し、どちらも Strapi がサポートした日です。
以下は、コミット ログから直接抽出した実際のタイムラインです。タイムスタンプなどもすべて含まれています。
空のルートから、SEO に配慮した、Strapi に裏付けられたライブのブログとケーススタディのセクションまで: 1 営業日のうち約 8 時間半。 8 つのコミット。開発者は1人。各コミットは数分で確認できるほど小さいものです。
正直に言っておきたいことがいくつかあります。これは、投稿の中で最も過剰に販売されやすい部分だからです。
Strapi スキーマはその日の朝に設計されたわけではありません。コンテンツ モデル (コレクション、コンポーネント、関係) は、コンテンツを最もよく知っているチームによって事前に決定されました。 Claude Code は、フロントエンドを既存のスキーマに接続しました。
コンテンツ自体を移行する必要はまったくありませんでした。 Strapi はすでに古いサイトの背後にあるヘッドレス CMS であったため、すべての投稿とケーススタディはすでにそこに存在し、編集され、洗練されていました。新しいフロントエンドを同じ Strapi インスタンスに指定しただけです。コンテンツの移動、コピーペースト、再編集は一切行われませんでした。唯一変更されたのは、既存のコンテンツが新しいデザインでどのようにレンダリングされるかということでした。
ポーランドは翌週も着陸を続けました。Strapi アセットの応答性の高い画像ラダー、固定 5 に強化されたフィルター チップ セット、プルインを停止するように調整された画像サイズ

g フル解像度のオリジナルをブログ タイルに貼り付けます。移行は1日でした。完成には 1 週間の小さなコミットを要しました。
しかし、これまで「これはスプリントだ」とされていた部分は 1 日でした。そして、タイピングよりもレビューに時間を費やした一日でした。
投稿全体のアンカーとなる比較ポイント: 以前の bitnoise.pl ビルドは、3 人の開発者 (Nowak (249 時間)、Daria (66 時間)、Gąsiorek (105 時間)) で合計 420 開発時間でした。開発のみで、設計時間は含まれていません。もともとブログとケーススタディだった 420 時間の比例部分を割り当てても、これら 2 つのセクションの再構築は 1 労働日以内に収まります。それはモデルのトリックではありません。入力をやめてステアリングを開始すると、このようなことが起こります。
古い bitnoise.pl が漂流していました。決して恥ずかしいことではありません。ドリフトが騒々しいことはめったにありません。ポジショニングは変わり、ケーススタディの組み合わせは変わり、2026 年に実際に販売されたサービスは、ホームページが主導したサービスではありませんでした。このサイトは、2026 年の企業の 2023 年のスナップショットでした。
それを修正するのは創業者レベルの仕事です。プロンプトを作成する前に、次のように書きました。
ポジショニングドキュメント: 私たちが現在誰なのか、誰に販売しているのか、トップページにどのような証拠を掲載したいのか。
各ページの目的ではなく、各ページの目的を示す、すべてのルートの意図を備えたサイトマップ。
その意図に沿った、ページごとのコンテンツ概要。
そのどれもがクロード・コードから生まれたものではありません。それは部屋から出てきた。モデルは後でそれらのドキュメントをすべてコンテキストとして読み取り、頻繁にそれらを参照しましたが、それらのドキュメントはまったく書き込まれませんでしたし、私たちもそれを望んでいませんでした。これは、LLM が危険なほど流暢である層です。彼らは、調べてみると別の会社の説明である、もっともらしいポジショニングを喜んで作成します。
これを再現する人への教訓: モデルから始めないでください。モデルが読み取る必要があるドキュメントから始めます。

すべてのヒーローの見出し、すべてのサービスの宣伝文句、すべてのケーススタディの説明は、CEO/CTO によって書かれたか大幅に書き直されました。モデルはシソーラスと長さトリマーでした。フリースタイルコピーはさせませんでした。
なぜだめですか？重大度順に 3 つの理由:
声の漂流。モデルのデフォルトは他の全員のデフォルトです。ホームページを作成させるのが、他の代理店と同じように見せる最も早い方法です。
幻覚による能力主張。サービスのコピーを作成するモデルは、Google がインデックスするページ上で販売していない機能を約束することから遠ざかる悪いプロンプトの 1 つです。
一般的な SaaS の話。技術的に何も問題がない場合でも、結果は他の何千ものサイトと同じように見えます。サイトは販売資産です。言葉は私たちのものでなければなりません。
これを最も明確に確認できる場所は、ホームページのサービス セクションです。4 つのスライド (チーム強化、プロジェクト デリバリ、AI ソリューション、カスタム ソリューション) が貼り付けられたスタックで、その間に長方形モーフの遷移が表示されます。各スライドには、モデルが作成されたコピー ブロックと出荷されたコピー ブロックがあり、それらの間の差が、音声所有権の作業が実際にどのようなものであるかがわかります。
1 つのスライド (プロジェクトのデリバリーを選択します。バブルから進捗バーまで表示されるスライド) では、パターンは毎回同じでした。
このモデルの草案では、「私たちは、期日どおり、予算内でプロジェクトを遂行できるよう支援します。」とヘッジしています。
出荷されたバージョンでは、具体的な成果物に名前が付けられ、実際のエンゲージメント モデルに基づいて、「できること」や「お手伝いすること」はすべて削除されました。
単語の約 40% が生き残りました。構造の約 100% が書き直されました。
その比率はサイトのすべてのページにわたって維持されました。モデルは作家ではありません。このモデルは、著者の最初の草案ジェネレーターです。これは便利ですが、同じものではありません。
新しいサイトのビジュアル言語 - スクロール駆動のヒーロー、変形する四角形を備えたスティッキー サービス スタック、スクランブル テックス

ハイライト、カーソル反発ドット フィールド、カスタム ソリューション セクションのヘビ アニメーション、Z 字型スペーサーは、Figma フローとモーション プロトタイプとしてすべて UX/UI チームから提供されました。クロード コードはそれらを実装しました。それはそれらを発明したわけではありません。
これが、このサイトが AI 汎用に見えない最大の理由です。
ロックを解除するには、モデル コンテキスト プロトコルの Figma プラグインである Figma MCP を使用します。接続すると、Claude Code は実際の Figma ファイル (コンポーネント、フレーム、自動レイアウト、変数、デザイン トークン、正確な間隔) を読み取ることができます。実装パスは発明ではなく翻訳になります。 「高級感を与える」というプロンプトはありません。 「このフレームを、これらのトークンを使用して、このモーション参照と一致させて実装する」というものがあります。
この設定を再現する場合、重要な点は次の 3 つです。
1日目はFigma MCPを配線。スクリーンショットは不適切なプロンプトです。実際のフレームは良いものです。
デザイントークンを権威あるものにします。 Tailwind スケールと Figma 変数は同じ数値を共有します。モデルはギャップ 6 が何であるかを推測する必要はありません。
モーションをハンドオフではなく、デザイナーと開発者の間のループとして扱います。これは、「AI がコードを記述する」というストーリーでは常に省略される部分なので、単一のアニメーションがどのように構築されたかを正確に説明する価値があります。
サイト上のあらゆる重要なインタラクションに対して、作業は同じループを通過して行われます。これらのどれでもない
ステップは単独で動作するモデルです。
設計者は、Figma の両端 (開始フレームと最終フレーム) を実際の値を持つ実際の状態として設定します。
デザイナーは、タイミング、イージング、感触など、一方が他方になる様子を示すモーション プロトタイプを構築しました。
それは、アニメーションの完全な説明とその技術的要件および実装計画を含む書面による概要になり、開発者は事前にそれをレビューしました。

コード行が書き込まれました。
クロード コードは、プロトタイプと概要の 2 つのフレームに対してそれを実装しました。
開発者は、実際のブラウザと実際の携帯電話で結果をチェックしました。プロトタイプでは正しく感じられる曲線でも、ミッドレンジの Android では 60fps では間違っていると感じる可能性があるためです。
開発者は、モーションがプロトタイプと一致するまで、時にはクロード コードに戻って調整し、多くの場合は速度、継続時間、イージング、加速度、遅延などの生の数値を手動で調整しました。
デザイナーは最終的な実行効果をレビューし、承認しました。ブラウザ上の実物に対する彼らの承認が関門でした。モデルでも開発者でもありません。
モデルがキーフレームを書きました。 「正しい」ことがどのようなものかを決定するものではありませんし、アニメーションが完成したと判断するものでもありませんでした。それは、モーションを設計した人と、それが実際のデバイスで実行されているのを感じることができる人との間の会話に残りました。
サイト上の特徴的なインタラクションの一部。それぞれがそのループを通じて構築されています。
スクロール主導のヒーローの見出し (初日に掲載)。
キー上の連続した境界線の描画により、タイルが強調表示されます。
サービスビジュアルのスクロール駆動の長方形モーフ。
プロジェクトのデリバリーのバブル スライド。
AI ソリューションの進行 → スピナー → タブレット モーフ。
共有カーソル反発ドットキャンバスアクロ

[切り捨てられた]

## Original Extract

420 hours → 78 hours: rebuilding bitnoise.pl with
Claude Code and Figma MCP

420 hours → 78 hours: rebuilding bitnoise.pl with Claude Code and Figma MCP Bitnoise _ SERVICES INDUSTRIES ENGAGEMENT CASE STUDIES OUR PRODUCTS ABOUT INSIGHTS CONTACT MENU Back to all insights Published on 9 Jun 2026
420 hours → 78 hours: rebuilding bitnoise.pl with
Claude Code and Figma MCP
Back to the start Back Table of content An honest accounting of where the speed came from, where it didn't, and which parts of the work we'd never hand to a machine.
Not a single line of code on the website you're reading was typed by a human.
Every .tsx file, every animation curve, every Strapi fetcher, every piece of SEO scaffolding — 24,296 lines of it across 120 React components — came out of a Claude Code session. We wrote zero of it by hand. We reviewed all of it by hand.
And it still took us eight weeks.
That second number is the one we want to talk about. The story the industry keeps telling itself right now is "AI writes the code, the developer goes for coffee." That isn't what happened here. What happened is that the 15% of the project that used to be typing got compressed by about 5×, and the other 85% — strategy, copy, design, motion prototyping, code review, QA, SEO — got exactly the attention it always needed. The keyboard stopped being the bottleneck. Everything else stayed expensive.
This post is the honest accounting. The numbers come straight from git log . The comparison points come from time reports of the previous bitnoise.pl build. We're sharing both because we think the interesting question for any team considering this stack isn't "can the model code?" — it obviously can — but "what does the rest of the team have to do for the model's output to be worth shipping?"
~5.4× less developer time on the dev portion alone — and the new site has more pages, more animation, full Strapi integration for blog and case studies, and SEO/GEO scaffolding the old one never had.
The repo itself, audited from git log between 2026-04-08 (first commit) and 2026-06-02 (launch-ready):
Calendar span: 56 days (~8 weeks).
Active development days: 25. The other 31 were design, copy, review, QA, and waiting on decisions — all the work the model doesn't do.
Total commits: 181. Every one of them generated by Claude Code, reviewed by Claude Code in a fresh-context pass, signed off by a human.
Lines added / removed: 47,173 inserted, 5,676 deleted, across 857 file-changes.
Codebase shape today: 24,296 LOC across 120 .tsx files, 47 .ts files, and a single .css file (Tailwind v4 — yes, one stylesheet).
Focused coding time: ~78 hours across 53 discrete sessions. (Method: sum the wall-clock intervals between commits within a session; a gap longer than 90 minutes starts a new session. It's not a billing-grade figure, but it's the right shape.)
Average commit: ~26 minutes of attention. Most diffs are small enough to read in under a minute. That's deliberate.
Claude Code — for code generation and for the code review pass.
Figma MCP — the design system, components, frames, variables, and motion prototypes live in Figma and the model reads them directly. No "implement from screenshot" guesswork.
Stack: TanStack Start + Vite 7 + React 19 + Motion 12 + Tailwind v4. Strapi as the CMS. Resend for the contact form.
52,800 changed lines × 70 chars ÷ 4 chars-per-token = 925k output tokens of shipped code.
With read context, retries, and review passes, end-to-end usage was roughly ~10–20 million tokens.
We actually paid for this with a Claude Max plan, but if you're estimating for your own team on pay-as-you-go API rates with prompt caching, the equivalent cost lands roughly $90–$180 depending on your Opus/Sonnet mix — Opus-heavy with effective caching sits near the top of that, a Sonnet-leaning mix near the bottom. Budget $200 and you'll almost certainly come in under it.
420 hours of typing replaced by 78 hours of judgment. The judgment is the part you can't outsource.
If we had to pick the single moment that justified the whole approach, it would be May 20, 2026 — the day we ported the entire blog and case-studies sections off the old site and onto the new stack, with Strapi behind both.
Here's the actual timeline, straight from the commit log — timestamps and all:
From empty route to live, SEO-clean, Strapi-backed blog and case-studies sections: about eight and a half hours of one working day. Eight commits. One developer. Each commit small enough to review in a couple of minutes.
A few things we want to be honest about, because this is the part of the post that's easiest to over-sell:
The Strapi schemas weren't designed that morning. The content model — collections, components, relations — was decided ahead of time by the team that knew the content best. Claude Code wired the frontend to a schema that already existed.
The content itself didn't have to be migrated at all. Strapi was already the headless CMS behind the old site, so every post and case study already lived there, edited and polished. We just pointed the new frontend at the same Strapi instance. There was no content move, no copy- paste, no re-editing — the only thing that changed was how that existing content gets rendered in the new design.
Polish kept landing over the following week — a responsive image ladder for Strapi assets, the filter chip set tightened to a fixed five, image sizing tuned to stop pulling full-resolution originals onto blog tiles. The migration was a day. The finish took a week of small commits.
But the part that used to be "this is a sprint" was a day. And it was a day we spent reviewing more than typing.
The comparison point that anchors the whole post: the previous bitnoise.pl build came in at 420 dev hours total across three developers — Nowak (249h), Daria (66h), Gąsiorek (105h), dev-only, no design hours included. Even allocating the proportional slice of those 420 hours that was originally blog + case studies, the rebuild of those two sections fit inside a single working day. That isn't a model trick. That's what happens when you stop typing and start steering.
The old bitnoise.pl had drifted. Not in any single embarrassing way — drift is rarely loud. The positioning had moved, the case-study mix had changed, the services we actually sold in 2026 weren't the services the homepage led with. The site was a 2023 snapshot of a 2026 company.
Fixing that is a founder-level job. Before we wrote a prompt, we wrote:
A positioning doc: who we are now, who we sell to, what proof we want on the front page.
A sitemap with intent for every route — what each page is for , not what it is.
A page-by-page content brief, pinned to that intent.
None of that came out of Claude Code. It came out of a room. The model later read every one of those documents as context, and it referred back to them often, but it didn't write any of them, and we wouldn't have wanted it to. This is the layer where LLMs are dangerously fluent — they will happily produce a plausible-sounding positioning that is, on inspection, a description of a different company.
The takeaway for anyone replicating this: don't start with the model. Start with the document the model needs to read.
Every hero headline, every services blurb, every case-study lede was written or heavily rewritten by the CEO/CTO. The model was a thesaurus and a length trimmer. We did not let it freestyle copy.
Why not? Three reasons, in order of severity:
Voice drift. The model's defaults are everyone else's defaults. Letting it write the homepage is the fastest way to look like every other agency.
Hallucinated capability claims. A model writing services copy is one bad prompt away from promising a capability we don't sell, on a page Google indexes.
Generic SaaS-speak. Even when nothing is technically wrong, the result reads like a thousand other sites. The site is a sales asset. The words have to be ours.
The clearest place to see this is the services section on the homepage — the sticky stack of four slides (Team Augmentation, Project Delivery, AI Solutions, Custom Solutions) with the rectangle- morph transitions between them. Each slide had a model-drafted copy block and a shipped copy block, and the diff between them is what the work of voice ownership actually looks like.
For one slide (we'd pick Project Delivery — the bubble-to-progress-bar one), the pattern was the same every time:
The model's draft hedged: "We can help you deliver projects on time and on budget."
The shipped version named a concrete deliverable, anchored to an actual engagement model, and dropped every "we can" and "we help".
About 40% of the words survived. About 100% of the structure was rewritten.
That ratio held across every page of the site. The model is not the writer. The model is the writer's first-draft generator — which is useful, but it is not the same thing.
The visual language of the new site — the scroll-driven hero, the sticky services stack with the morphing rectangle, the scramble-text reveal on the highlights, the cursor-repulsion dots field, the snake animation on the Custom Solutions section, the Z-shape spacer — came entirely from the UX/UI team, as Figma flows and motion prototypes. Claude Code implemented them. It did not invent them.
This is the single biggest reason the site doesn't look AI-generic.
The unlock is Figma MCP — the Figma plug-in for the Model Context Protocol. With it connected, Claude Code can read the actual Figma file: components, frames, auto-layout, variables, design tokens, exact spacing. The implementation pass becomes translation , not invention . There is no "make it feel premium" prompt; there is "implement this frame, using these tokens, matching this motion reference."
If you're replicating this setup, the three things that earned their keep:
Wire Figma MCP on day one. A screenshot is a bad prompt. The actual frame is a good one.
Make the design tokens authoritative. Our Tailwind scale and our Figma variables share the same numbers. The model never has to guess what gap-6 is supposed to be.
Treat motion as a loop between designer and developer, not a hand-off. This is the part the "AI writes the code" story always skips, so it's worth spelling out exactly how a single animation got built.
For every non-trivial interaction on the site, the work moved through the same loop. None of these
steps is the model working alone:
The designers set the two ends in Figma — the starting frame and the final frame, as real states with real values.
The designers built a motion prototype showing how one becomes the other: the timing, the easing, the feel.
That got turned into a written brief — a full description of the animation with its technical requirements and an implementation plan — which the developer reviewed before a line of code was written.
Claude Code implemented it against the two frames, the prototype, and the brief.
The developer checked the result in a real browser and on a real phone — because a curve that feels right in a prototype can feel wrong at 60fps on a mid-range Android.
The developer tuned it — sometimes by going back to Claude Code, more often by hand-adjusting the raw numbers: speed, duration, easing, acceleration, delay — until the motion matched the prototype.
The designers reviewed the final, running effect and signed off. Their approval, on the real thing in the browser, was the gate. Not the model's, not the developer's.
The model wrote the keyframes. It did not decide what "right" felt like, and it was never the one who got to call an animation done. That stayed a conversation between the people who designed the motion and the person who could feel it running on a real device.
Some of the distinctive interactions on the site, each one built through that loop:
The scroll-driven hero headline (landed on day one).
The sequential border-draw on the key-highlights tiles.
The scroll-driven rectangle morph for the services visual.
The Project Delivery bubble slide.
The AI Solutions progress → spinner → tablet morph.
The shared cursor-repulsion dots canvas acro

[truncated]
