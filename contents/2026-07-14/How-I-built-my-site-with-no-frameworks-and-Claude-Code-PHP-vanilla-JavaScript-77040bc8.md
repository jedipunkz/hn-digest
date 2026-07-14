---
source: "https://ivanmisic.net/blog/ai-tools/how-i-built-this-site-with-claude-code"
hn_url: "https://news.ycombinator.com/item?id=48910978"
title: "How I built my site with no frameworks and Claude Code (PHP, vanilla JavaScript)"
article_title: "How I Built This Site With Claude Code | Ivan Misic | Ivan Misic"
author: "imisic"
captured_at: "2026-07-14T19:03:12Z"
capture_tool: "hn-digest"
hn_id: 48910978
score: 2
comments: 1
posted_at: "2026-07-14T18:19:01Z"
tags:
  - hacker-news
  - translated
---

# How I built my site with no frameworks and Claude Code (PHP, vanilla JavaScript)

- HN: [48910978](https://news.ycombinator.com/item?id=48910978)
- Source: [ivanmisic.net](https://ivanmisic.net/blog/ai-tools/how-i-built-this-site-with-claude-code)
- Score: 2
- Comments: 1
- Posted: 2026-07-14T18:19:01Z

## Translation

タイトル: フレームワークとクロード コードを使用せずにサイトを構築した方法 (PHP、バニラ JavaScript)
記事のタイトル: クロード コードを使用してこのサイトを構築する方法 |イヴァン・ミシック |イヴァン・ミシック
説明: 完全なビルド ストーリー: カスタム PHP MVC、バニラ JS、ブルータリスト CSS、および開発パートナーとしてのクロード コード。アーキテクチャの決定、ワークフロー、そして正直な教訓。

記事本文:
メインコンテンツにスキップ
ステーション · ivanmisic.net
ログインする
IM
イヴァン・ミシチ
製品・技術・AI
⌕
アーカイブを検索
⌘K
§00
ホーム
§01
執筆
§02
ツールシェッド
§03
ガイド
§04
アトラス
§05
について
§06
お問い合わせ
★最新船
私が共有している Claude Code ツール
№37 · 8 分 · 読む →
▸ ここから始めてください
クロードコードガイド
初心者ガイド・ゼロからライブサイトへ→
ログインする
§01
AIとツール
↵ AIとツールに戻る
3 週間、フレームワークなし
クロード コードを使用してこのサイトを構築した方法
ページをマークダウンとしてコピー
マークダウンを開く
Claude.ai で開く
ChatGPTで開く
このページについて
なぜゼロから構築するのか?
このウェブサイトは、約 3 週間の実作業をかけてゼロから構築しました。休暇中の 1 週間、数回の週末、そしてその間の深夜。ワードプレスはありません。ララベルはありません。リアクションはありません。追い風なし。 PHP、バニラ JavaScript、カスタム CSS デザイン システム、そして開発パートナーとしての Claude Code だけです。
この投稿がその全容です。何を構築したのか、どのように構築したのか、何が機能したのか、何が機能しなかったのか、何を変更するのか。
フレームワークとプラットフォームに基づいて構築された製品を何年も管理してきた後、実際に内部に何があるかを理解したいと思いました。抽象化ではありません。そのものそのもの。
しかし、別の理由がありました。 AI 支援開発は今どこにでも使われていますが、私は、それを管理したり、仕事でその恩恵を受ける前に、何かを理解する必要があると強く信じています。その長所と短所を直接知る必要があります。 Claude Code を使用して実際のプロジェクトを構築することは、AI が実際にどのように機能するかを学ぶことでした。コードを書くことはもちろんですが、計画、設計、戦略立案、問題解決も行います。
バックエンドフレームワークはありません。 Laravel も Symfony もありません。ゼロからのカスタム MVC。
フロントエンドフレームワークはありません。 React も Vue も jQuery もありません。
CSS フレームワークはありません。追い風もブートストラップもありません。トークンを使用したカスタム設計システム。
きちんと構築してください

。データベースの移行、サービス層、機能フラグ、展開パイプライン。おもちゃのプロジェクトではありません。
実装にはクロードコードを使用します。アーキテクチャとデザインの決定はすべて私が行います。クロードがコードを書きます。
npmはありません。フロントエンド用の Composer パッケージはありません。自分で作成した PHP ミニファイア以外にビルド ツールはありません。
リクエスト→index.php→ブートストラップ→ルーター→コントローラー→ビュー
↓
サービス → モデル → データベース
しかし、興味深いのは、私が強制した制約です。
コントローラーは HTTP を処理します。それでおしまい。これらはリクエスト データを抽出し、認証を確認し、サービスまたはモデルを呼び出します。ビジネス ロジックが含まれることはありません。
サービスはすべての突然変異を処理します。すべての作成、更新、および削除は、入力を検証し、スラッグを生成し、キャッシュを管理し、ServiceResult オブジェクトを返すサービスを経由します。コントローラーが書き込みのためにモデルに直接触れることはありません。
モデルは永続性を処理します。すべてのモデルは BaseModel を拡張し、許可された列の $fillable ホワイトリストを定義し、プリペアド ステートメントを排他的に使用します。モデルの外には生の SQL はありません。
// コントローラー (シン - HTTP のみ)
$result = BlogPostService::create($data);
if ($result->failed()) {
return $this->error($result->errorString());
}
// サービス（ビジネスロジック）
パブリック静的関数 create(array $data): ServiceResult {
$data = self::normalizeData($data);
$errors = self::validate($data);
if (!empty($errors)) return ServiceResult::failure($errors);
// スラッグ生成、キャッシュ無効化、読み込み時間計算...
ServiceResult::success($entity) を返します。
}
このパターンにより、コードベースが予測可能になりました。何かが壊れたときに、どのレイヤーを見ればよいのかが正確にわかります。
すべてのコンテンツ タイプには、 .env に機能フラグがあります。ブログ、ツール、ガイド、ジャーニー、API 仕様。フラグが無効になっている場合、ルーターは一般の訪問者に対して「近日公開」ページを表示します。

ただし、管理パネルにはまだすべてが表示されます。これは、新しいセクションをプライベートで構築して設定し、スイッチを入れて起動できることを意味します。
FEATURE_BLOG=true
FEATURE_TOOLS=true
FEATURE_GUIDES=true
FEATURE_NEWS=false # 次にこれを構築します
データベースの移行
すべてのスキーマ変更には、UP セクションと DOWN セクションを含む移行ファイルがあります。それらを前方に実行したり、最後のバッチをロールバックしたり、適用された内容を追跡したりできます。 CLI または管理 UI 経由で動作します。
40 件の移行があり、さらに増えています。移行システムを数回繰り返して導入したため、実際のスキーマ変更の数は追跡されたものよりも多くなります。ただし、最初のブログ スキーマから OAuth テーブル、分析、コンテンツ同期まで、それ以降のすべての変更が追跡され、元に戻すことができます。
ブログ投稿は Markdown で記述され、データベースに保存され、実行時に HTML にレンダリングされます。事前レンダリングされたキャッシュはありません。コンテンツ同期システムは、すべてを JSON バンドル (ID ベースではなくスラッグベース) にエクスポートし、本番環境にインポートできます。ビュー数と分析は、同期中に上書きされることはありません。
私は「ブルータリスト・ボールド」を選びました。暗い背景、あらゆる場所に 2 ピクセルの境界線、ゼロの境界線半径、ラベルとメタデータ用の等幅フォント、アクセント カラーとしてのエレクトリック ライム (#D4FF00)。
注意喚起。あなたが今読んでいるサイトはこれではありません。 v1 ブルータリスト ビルドは、2026 年 4 月に Station のウォーム ノワール デザインに置き換えられました。完全な再設計のストーリー (4 暦日、1 人、エンドツーエンド) は、「4 暦日、1 人、1 つの完全な再設計」にあります。この投稿の残りの部分は、元の v1 ビルドのストーリーであり、現状のままです。
まず、制約により意思決定が迅速化されます。 border-radius が 0 の場合、「これは 4px にするべきか 8px にするべきか?」という議論は決して行われません。影が許可されない場合は、境界線を使用します。アクセントカラーが 1 つだけであれば、カラーパレットに時間を無駄にする必要はありません。私が追加したすべてのルールは 1 つでした

決断を下すことが少なくなります。
第二に、見た目が特徴的です。丸みを帯びた角と柔らかなグラデーションの世界の中で、シャープなエッジが際立ちます。このサイトはテンプレートではないため、テンプレートのようには見えません。
すべては tokens.css に存在します。色、間隔、フォント サイズ、境界線の幅、トランジション。コンポーネントはトークンのみを参照します。決してハードコーディングされた値ではありません。
/* トークン.css */
--color-primary: #D4FF00;
--color-bg: #0a0a0a;
--space-4: 1レム;
--space-8: 2rem;
--ボーダー幅: 2px;
--font-mono: 'JetBrains Mono'、等幅;
このシステムは、CSS カスタム プロパティを通じて、暗い/明るいテーマと複数のアクセント カラー (ライム、シアン、ローズ) をサポートしています。テーマの切り替えは、コンポーネントの書き換えではなく、トークン値の変更を意味します。
48個のCSSファイル。 15,500行以上。すべてのコンポーネントは独自のファイル内にあります。ユーティリティは最後にロードされるため、コンポーネントをオーバーライドします。これは、ほとんどの人が個人サイト用に作成するものよりも CSS が多くなりますが、トークン システムにより保守可能です。
これは誰もが疑問に思う部分です。それで、実際に何が起こったのでしょうか？
このスピードにより、ものづくりについての考え方が変わりました。丸一日かかると見積もっていた作業が、1 ～ 2 時間で完了しました。定型文だけではありません。クロードは、適切なインデックス作成、CSS アーキテクチャの決定、セキュリティ パターン (CSRF トークン、プリペアド ステートメント、レート制限)、デプロイメント パイプライン設計を使用して、複雑なデータベース クエリを処理しました。
やりたいことを平易な言葉で説明し、動作する実稼働品質のコードを返すことができました。 「検証、スラッグ生成、キャッシュ無効化を備えたブログ投稿用のサービス レイヤーを構築してください。」終わり。 「バッチによるロールバックをサポートする移行システムを作成します。」終わり。
それは、コーヒーブレイクを必要としない非常に速く、非常に忍耐強い上級開発者とのペアプログラミングのように感じました。
早い段階で、クロードがセッションをまたいでアーキテクチャ上の決定を「忘れる」ことに気づきました。新しい b が作成されます。

utton スタイルがすでに存在する場合はそれを使用するか、確立したものとは異なるエラー処理パターンを使用します。
修正は、プロジェクト ルートにある詳細な CLAUDE.md ファイルでした。デザイントークン、命名規則、アーキテクチャルール、コンポーネントインベントリ。クロードは毎回のセッションの開始時にそれを読みます。また、CSS 標準からデータベース規約まですべてをカバーする 15 個のルール ファイルを .claude/rules/ に作成しました。 (これについては、「Making Claude Code Work on Bigger Projects」で詳しく説明します。)
この永続的な記憶が転機となりました。ルールが存在すると、クロードは新しいパターンを発明するのをやめ、確立されたパターンに従い始めました。
オーバーエンジニアリング。クロードは必要以上に抽象化を加える傾向があります。 「単なる単純な関数」は、抽象基本クラス、2 つのインターフェイス、およびファクトリとともに返されます。私は「もっとシンプルに」と言うことに多くの時間を費やしました。
CSS の一貫性。厳密なルールがなければ、クロードは重複した CSS クラスを作成してしまいます。 .card-regulator がすでに存在する場合の新しいカード コンポーネント。 .btn--primary がすぐそこにあったときの新しいボタンのバリアント。ルール ファイルによってこの問題は修正されましたが、それは重複したスタイルを数回クリーンアップした後でのみでした。
デザインテイスト。クロードにはそれがありません。ブルータリズムの美学は完全に私のビジョンでした。私が何を望んでいるかを説明すると、クロードはうまく実行してくれましたが、「境界線の半径を持たない 2 ピクセルの境界線とライム グリーンのアクセントを使用しましょう」とは決して提案しませんでした。視覚的な決定はすべて私が行いました。
コンテキストのドリフト。多くの変更を伴う長いセッションでは、クロードは以前の決定を見失うことがありました。セッションを一度に 1 つのタスクに集中し続けることが役に立ちました。ルールファイルも同様です。
愚かなセッション。これは説明が難しいです。時々、クロードは...仕事をやめてしまいます。クラッシュではありません。エラーは出ません。応答し続けますが、品質は崖から落ちてしまいます。単純なタスクがありました

1時間前にうまくやっていたのはナンセンスです。すべてのセッションに従っていたルールを無視します。それは、今行ったことと矛盾する変更を加えることになるでしょう。
私の理論: コンテキストのオーバーロードです。複雑な作業を 30 ～ 45 分続けると、何かが故障します。修正は驚くほど簡単でした。すべてを閉じます。立ち去ってください。数時間後に戻ってきて、新たなセッションを開始すると、クロードは何事もなかったかのように、中断したところから正確に再開します。説明はありません。謝罪はありません。有能な状態に戻るだけです。
これがリズムになりました。 30〜45分間集中的に作業します。クロードが知っておくべきことで苦労し始めても、戦わないでください。異なる結果を期待して同じプロンプトを 5 回再試行しないでください。やめてください。後はフレッシュセッション。その瞬間はイライラしますが、混乱した AI とさらに 1 時間議論するよりも早いです。
自然言語アーキテクチャ。 「コントローラーは HTTP のみを処理し、サービスはビジネス ロジックを処理する必要がある」と言うと、Claude は作成したすべてのコントローラーとサービスにわたってその分離を一貫して実装します。
大規模なリファクタリング。 「すべてのカテゴリの CRUD を共通の特性に抽出します。」クロードは既存の実装を読み取り、共通のパターンを特定し、トレイトを作成して、両方のサービスを更新します。ワンパスで。
セキュリティパターン。クロードは思ったより警備が上手だった。列のホワイトリスト、準備されたステートメント、CSRF 検証、レート制限。私が求めていなかったセキュリティ対策を積極的に追加してくれるでしょう。
3 週間の集中的な作業により、最初のバージョンが作成されました。以下の数字は、v2 の再設計と最初のスプリントに加えて 1 年間の反復を経た、現在のサイトの状況を示しています。
個人サイトにこれほど多くのコードが含まれる理由
上の表を合理的に読み取ると、「なぜブログには約 48,700 行もの PHP が含まれるのか?」ということになります。答えは、

ブログは、約 12 ある機能のうちの 1 つです。これらの数字がすべての原動力となります。
公共側。カテゴリ、タグ、関連投稿、全文検索を備えたブログ。評価とクリック追跡を含むツール ディレクトリ。モジュールとコンテンツ ゲーティングを含むステップバイステップのガイド。記事クラスタリングによる RSS 主導のニュース集約。概要/現在/お問い合わせページ。 OAuth (Google、GitHub、LinkedIn) およびマジックリンク ログインを使用するユーザー アカウント。 SEO インフラストラクチャ: 動的サイトマップ、robots.txt、構造化データ、正規の処理。
管理人さん、氷山です。ここにはほとんどのコードが置かれており、公開サイトよりも大きいです。以下をカバーする 21 個の管理コントローラー:
サーバー側分析を備えたダッシュボード (GA なし、サードパーティ追跡なし)
ブログエディター、メディアマネージャー、カテゴリーおよびタグ管理
クリック追跡レビューを備えたツールマネージャー
モジュールレベルのオーサリングを備えたガイドエディター
ニュース管理 (情報源、ブリーフィング、マニュアルキュレーション)
LinkedIn、Twitter、Reddit の投稿をスケジュールするためのマーケティング カレンダー (ステータスの追跡とソース記事へのリンク付き)
サイト上の任意の URL にメモを付けて残せる付箋システム
リダイレクト、移行ランナー、コンテンツ同期 (ローカルと本番環境間のバンドルのエクスポート/インポート)
その下にあるインフラ。厳密な型付けを備えたカスタム MVC。シンコのサービス層

[切り捨てられた]

## Original Extract

The full build story: custom PHP MVC, vanilla JS, brutalist CSS, and Claude Code as my dev partner. Architecture decisions, workflow, and honest lessons.

Skip to main content
station · ivanmisic.net
log in
IM
IVAN MIŠIĆ
product · tech · ai
⌕
Search archive
⌘K
§00
Home
§01
Writing
§02
Toolshed
§03
Guides
§04
Atlas
§05
About
§06
Contact
★ latest ship
The Claude Code tools I'm sharing
№37 · 8 min · read →
▸ start here
Claude Code Guide
beginner guide · zero to a live site →
log in
§01
AI & Tools
↵ back to ai & tools
Three Weeks, Zero Frameworks
How I Built This Site With Claude Code
Copy page as Markdown
Open Markdown
Open in Claude.ai
Open in ChatGPT
On This Page
Why Build From Scratch?
I built this website from scratch over about three weeks of actual work. A week during holiday, a handful of weekends, and some late nights in between. No WordPress. No Laravel. No React. No Tailwind. Just PHP, vanilla JavaScript, a custom CSS design system, and Claude Code as my development partner.
This post is the full story. What I built, how I built it, what worked, what didn't, and what I'd change.
After years of managing products built on frameworks and platforms, I wanted to understand what's actually under the hood. Not the abstraction. The thing itself.
But there was another reason. AI-assisted development is everywhere now, and I'm a strong believer that you need to understand something before you can manage it or benefit from it in your work. You need to know its strengths and weaknesses firsthand. Building a real project with Claude Code was about learning how AI works in practice. Writing code, yes, but also planning, designing, strategizing, and problem-solving.
No backend frameworks. No Laravel, no Symfony. Custom MVC from scratch.
No frontend frameworks. No React, no Vue, no jQuery.
No CSS frameworks. No Tailwind, no Bootstrap. Custom design system with tokens.
Build it properly. Database migrations, service layer, feature flags, deployment pipeline. Not a toy project.
Use Claude Code for implementation. I'd make all the architecture and design decisions. Claude would write the code.
No npm. No Composer packages for the frontend. No build tools besides a PHP minifier I wrote myself.
Request → index.php → Bootstrap → Router → Controller → View
↓
Service → Model → Database
But the interesting parts are the constraints I enforced.
Controllers handle HTTP. That's it. They extract request data, check authentication, and call services or models. They never contain business logic.
Services handle all mutations. Every create, update, and delete goes through a service that validates input, generates slugs, manages cache, and returns a ServiceResult object. Controllers never touch models directly for writes.
Models handle persistence. Every model extends BaseModel , defines a $fillable whitelist for allowed columns, and uses prepared statements exclusively. No raw SQL anywhere outside of models.
// Controller (thin - HTTP concerns only)
$result = BlogPostService::create($data);
if ($result->failed()) {
return $this->error($result->errorString());
}
// Service (business logic)
public static function create(array $data): ServiceResult {
$data = self::normalizeData($data);
$errors = self::validate($data);
if (!empty($errors)) return ServiceResult::failure($errors);
// slug generation, cache invalidation, reading time calc...
return ServiceResult::success($entity);
}
This pattern made the codebase predictable. When something breaks, I know exactly which layer to look at.
Every content type has a feature flag in .env . Blog, tools, guides, journeys, API specs. When a flag is disabled, the router shows a "Coming Soon" page for public visitors, but the admin panel still shows everything. This means I can build and populate a new section privately, then flip a switch to launch it.
FEATURE_BLOG=true
FEATURE_TOOLS=true
FEATURE_GUIDES=true
FEATURE_NEWS=false # Building this next
Database Migrations
Every schema change has a migration file with an UP and DOWN section. I can run them forward, roll back the last batch, and track what's been applied. Works via CLI or admin UI.
40 migrations and counting. We introduced the migration system a few iterations in, so the actual number of schema changes is higher than the tracked ones. But from the initial blog schema through OAuth tables, analytics, and content sync, every change since has been tracked and reversible.
Blog posts are written in Markdown, stored in the database, and rendered to HTML at runtime. No pre-rendered cache. The content sync system exports everything to a JSON bundle (slug-based, not ID-based) that can be imported on production. View counts and analytics are never overwritten during sync.
I went with "Brutalist Bold." Dark backgrounds, 2px borders everywhere, zero border-radius, monospace fonts for labels and metadata, and Electric Lime (#D4FF00) as the accent color.
Heads up. The site you're reading right now isn't this. The v1 brutalist build was replaced by a Station warm-noir design in April 2026. The full redesign story (four calendar days, one person, end-to-end) is in Four Calendar Days, One Person, One Full Redesign . The rest of this post is the original v1 build story, kept as-is.
First, constraints make decisions faster. When border-radius is 0, you never debate "should this be 4px or 8px?" When shadows aren't allowed, you use borders. When you only have one accent color, you don't waste time on color palettes. Every rule I added was one fewer decision to make.
Second, it looks distinctive. In a world of rounded corners and soft gradients, sharp edges stand out. The site doesn't look like a template because it isn't one.
Everything lives in tokens.css . Colors, spacing, font sizes, border widths, transitions. Components reference tokens only. Never a hardcoded value.
/* tokens.css */
--color-primary: #D4FF00;
--color-bg: #0a0a0a;
--space-4: 1rem;
--space-8: 2rem;
--border-width: 2px;
--font-mono: 'JetBrains Mono', monospace;
The system supports dark/light themes and multiple accent colors (lime, cyan, rose) through CSS custom properties. Switching themes means changing token values, not rewriting components.
48 CSS files. 15,500+ lines. Every component in its own file. Utilities load last so they override components. It's more CSS than most people would write for a personal site, but the token system means it's maintainable.
This is the part everyone asks about. So what actually happened?
The speed changed how I think about building things. Tasks I'd estimate taking a full day were done in an hour or two. And not just boilerplate. Claude handled complex database queries with proper indexing, CSS architecture decisions, security patterns (CSRF tokens, prepared statements, rate limiting), and deployment pipeline design.
I could describe what I wanted in plain language and get working, production-quality code back. "Build me a service layer for blog posts with validation, slug generation, and cache invalidation." Done. "Create a migration system that supports rollback by batch." Done.
It felt like pair programming with a very fast, very patient senior developer who never needs a coffee break.
Early on, I discovered that Claude would "forget" architectural decisions across sessions. It would create a new button style when one already existed, or use a different pattern for error handling than what we'd established.
The fix was a detailed CLAUDE.md file at the project root. Design tokens, naming conventions, architectural rules, component inventory. Claude reads it at the start of every session. I also created 15 rules files in .claude/rules/ covering everything from CSS standards to database conventions. (I dig into this more in Making Claude Code Work on Bigger Projects .)
This persistent memory was a turning point. Once the rules existed, Claude stopped inventing new patterns and started following the established ones.
Over-engineering. Claude tends to add more abstraction than needed. "Just a simple function" would come back with an abstract base class, two interfaces, and a factory. I spent a lot of time saying "simpler."
CSS consistency. Without strict rules, Claude would create duplicate CSS classes. A new card component when .card-regular already existed. New button variants when .btn--primary was right there. The rules files fixed this, but only after I'd cleaned up several rounds of duplicate styles.
Design taste. Claude doesn't have it. The brutalist aesthetic was entirely my vision. Claude executed it well once I explained what I wanted, but it would never have suggested "let's do 2px borders with no border-radius and lime green accents." Every visual decision was mine.
Context drift. On long sessions with many changes, Claude would sometimes lose track of earlier decisions. Keeping sessions focused on one task at a time helped. So did the rules files.
The stupid sessions. This one is harder to explain. Sometimes Claude would just... stop working. Not crash. Not error out. It would keep responding, but the quality would fall off a cliff. Simple tasks it had handled fine an hour ago would produce nonsense. It would ignore rules it had been following all session. It would make changes that contradicted what it had just done.
My theory: context overload. After 30-45 minutes of complex work, something breaks down. The fix was surprisingly simple. Close everything. Walk away. Come back a few hours later, start a fresh session, and Claude would pick up exactly where it left off like nothing happened. No explanation. No apology. Just back to being competent.
This became a rhythm. Work intensely for 30-45 minutes. If Claude starts struggling with things it should know, don't fight it. Don't retry the same prompt five times hoping for a different result. Just stop. Fresh session later. It's frustrating in the moment, but it's faster than arguing with a confused AI for another hour.
Natural language architecture. I could say "the controller should only handle HTTP, the service should handle business logic" and Claude would implement that separation consistently across every controller and service it created.
Refactoring at scale. "Extract all category CRUD into a shared trait." Claude would read the existing implementations, identify the common patterns, create the trait, and update both services. In one pass.
Security patterns. Claude was better at security than I expected. Column whitelisting, prepared statements, CSRF validation, rate limiting. It would proactively add security measures I hadn't asked for.
Three weeks of focused work produced the first version. The numbers below are where the site stands today, after the v2 redesign and a year of iteration on top of that initial sprint:
Why a personal site has this much code
A reasonable read of the table above is "why does a blog have ~48,700 lines of PHP?" The answer is that the blog is one feature out of around a dozen. Those numbers power all of this.
Public side. Blog with categories, tags, related posts, and full-text search. Tools directory with ratings and click tracking. Step-by-step guides with modules and content gating. RSS-driven news aggregation with article clustering. About / Now / Contact pages. User accounts with OAuth (Google, GitHub, LinkedIn) and magic-link login. SEO infrastructure: dynamic sitemap, robots.txt, structured data, canonical handling.
Admin, the iceberg. This is where most of the code lives, and it's bigger than the public site. 21 admin controllers covering:
Dashboard with server-side analytics (no GA, no third-party tracking)
Blog editor, media manager, category and tag management
Tools manager with click-tracking review
Guides editor with module-level authoring
News management (sources, briefings, manual curation)
Marketing calendar for scheduling LinkedIn, Twitter and Reddit posts, with status tracking and links back to source articles
Sticky-note system for leaving notes attached to any URL on the site
Redirects, migrations runner, content sync (export/import bundles between local and production)
Infrastructure underneath. Custom MVC with strict typing. Service layer with thin co

[truncated]
