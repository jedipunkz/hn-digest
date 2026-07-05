---
source: "https://github.com/guy-lifshitz/tacheles"
hn_url: "https://news.ycombinator.com/item?id=48796026"
title: "Show HN: Detecting AI slop with regex and Stephen King"
article_title: "GitHub - guy-lifshitz/tacheles: Make AI-assisted writing sound like you wrote it, not a model: Tacheles flags the exact AI tells and shows you how to cut them. · GitHub"
author: "shtofadhor"
captured_at: "2026-07-05T17:57:31Z"
capture_tool: "hn-digest"
hn_id: 48796026
score: 2
comments: 0
posted_at: "2026-07-05T17:26:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Detecting AI slop with regex and Stephen King

- HN: [48796026](https://news.ycombinator.com/item?id=48796026)
- Source: [github.com](https://github.com/guy-lifshitz/tacheles)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T17:26:14Z

## Translation

タイトル: HN を表示: 正規表現とスティーブン キングを使用した AI スロップの検出
記事のタイトル: GitHub - guy-lifshitz/tacheles: AI 支援によるライティングを、モデルではなく自分が書いたように聞こえるようにする: Tacheles は、AI が指示した正確なフラグを立て、それをカットする方法を示します。 · GitHub
説明: AI 支援によるライティングを、モデルではなく自分が書いたかのように聞こえるようにします。Tacheles は、AI が指示したとおりにフラグを立て、それを切り取る方法を示します。 - ガイ・リフシッツ/タチェレス

記事本文:
GitHub - guy-lifshitz/tacheles: AI 支援によるライティングを、モデルではなく自分が書いたかのように聞こえるようにします。Tacheles は、AI が指示した正確なフラグを立て、それをカットする方法を示します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ガイリフシッツ
/
タチェレス
公共
通知
c にサインインする必要があります

ハンゲの通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット bin bin docs/ リファレンス docs/ リファレンス スクリプト scripts skill/ tacheles skill/ tacheles src src testing testing .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock Index.ts Index.ts package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
音声プロファイル
テクニカルエン
タチェレス
AI ライティング リンター: 散文のモデルの指示をオフラインで正確にキャッチします。
AI 支援によるライティングを、手本ではなく自分が書いたかのように聞こえるようにします。 Tacheles は、テキスト内で AI が伝える正確なフラグを立て、それを切り取る方法を示します。
スロップは 2 つあります。肥大化: LLM は、一度に 1 つのトークンを書き込みます。それぞれが次の可能性が最も高い単語であり、短くなく流暢かつ平均的に聞こえるように構築されます。そのため、アイデアに必要以上の言葉が詰め込まれます。スタイル: マシンとして読み取るレジスタに書き込みます。全角ダッシュ、X ではありません、 Y 、同じ 50 ワードです。 Tacheles は両方に正確な行で理由を付けてフラグを立て、マシン上で実行されます。AI なし、API キーなし、何もアップロードされておらず、毎回同じ結果になります。
あなたの代わりに書き換えられるわけではありません。何を切るかが表示されます。
タチェレス (תכל׳ס、tachles ): 要点、要点。
イディッシュ語出身。ドイツ語の Tacheles reden は、飾り気のない、率直に話すことを意味します。
前に。モデルがあなたに渡す段落:
今日の状況では、ツールだけではなく、
すべての読者の心に響く、堅牢なアイデアのタペストリー。
Tacheles は、「delve」、「robust」、「tapestry」、「それは X ではなく、Y です」とフラグを立てます。
オープナーとその周りのパッド。それが指しているものを切り取ると、次のようになります。
適切なツールを選択して、何かを書きましょう

読書中。
要点は同じだが、言葉は少なく、何も伝えていない。それはあなたのために「その後」を書くわけではありません。
何を切り取るべきかが示されているので、それを行うことができます。
代替手段にはないものは次のとおりです。
スコアではなくカットリスト。 AI 検出器はテキストをサーバーに送信し、「87% AI」を返します。数字に基づいて行動することはできません。 Tacheles は、コード リンターのやり方で、正確なスパンを正確な行に、その理由とともに名前付けします。それを修正して次に進みます。
フラグだけでなく修正も必要です。すべてのチェックは、英語の場合は Stephen King、ロシア語の場合は Ильяхов と Нора Галь という現役編集者のルールに対応しています。旗は何が間違っているかを示しています。ルールにはそれを修正する方法が記載されています。
多言語対応。 2 つの完全な言語パック (英語とロシア語) と 1 つの実験版 (ヘブライ語) が同梱されています。ドイツ語とスペイン語は次のリリースに含まれます。パックの追加はコードではなくデータであるため、英語だけでなくあらゆる言語をサポートします。
あなたに合わせて。自分の文章に合わせて調整すると、一般的なルールではなく、あなたの声からフラグがドリフトします。
プライベート、無料、繰り返し可能。モデルも API キーもなく、マシンからは何も残りません。毎回同じテキスト、同じ結果。
AI スロップには 3 種類のツールが使用され、それぞれ異なる仕事を行います。
ほとんどのオープンソース オプションは最初の 2 つです。これは 3 つ目です。多言語に対応し、ユーザーに合わせて調整されたカット リストです。
リンター。 tacheles チェック (カット リスト)、tacheles 測定 (あなたの文章のスタイロメトリー)、tacheles 比較ドラフト (リライトは少なくとも 10% カットされましたか、キングのルール)。オフライン、確定的、API キーなし。
リライトスキル。ループ全体を実行する移植可能な SKILL.md: チェック、言語の伝統に反して書き直し、クリーンになるまで再チェック。また、あなたにインタビューし、あなたのボイスアンカーを構築します。これを Claude Code、Claude.ai、Agent SDK、または任意のモデルにドロップします。
4つのプロフィール。エッセイ英語 、エッセイる 、コンサルティング英語 、テクニカル英語 。 1つをコピーして調整してください

あなた自身のものです。
言語パック。英語 (Stephen King) とロシア語 (Ильяхов / Нора Галь): 一連のテルとそれぞれの書き換え手順。
ガイドさん達。 docs/reference/ の手順とキャリブレーションと音声アンカーのウォークスルーを書き直します。
npx tacheles checkdraft.md # 1 回実行、インストールなし
npm install -g tacheles # またはインストールします
または、リポジトリのクローンを作成し、 Bun : bun run bin/tacheles checkdraft.md で実行します。
タチェレスチェックdraft.md
$ tacheles チェックdraft.md
draft.md（プロフィール：エッセイjp）
HIGH ライン 3 の禁止用語の詳細を調べる
ハイライン 3 の禁止用語の堅牢性
ハイライン 7 r-reframe-opener それはあなたが選ぶツールの問題ではありません。
MEDIUM 3行目 s-gpt-scaffolding 潜ってみよう
MEDIUM — 全角ダッシュの密度 3 つの全ダッシュ / 45 単語
不合格 — 高 3、中 2
コード リンターがレポートする方法と同様に、1 行に 1 つの検出結果が正確な行に表示されます。機械可読出力を CI または別のツールにパイプするには --json を追加します。
終了コードはクリーンな場合は 0、検出結果が高い場合は 1 であるため、CI またはコミット フックで終了コードをゲートできます。 HIGH は失敗します。 MEDIUM と LOW が報告されますが、失敗しません。
tacheles Compare-drafts <old> <new> もあります。これは、書き換えが少なくとも 10% カットすることをチェックします (King のルール)。
Tell（検出器）は個別のチェックです。それぞれは、ある種類のスロップを示す名前付きパターン、正規表現、または統計です。 s-banned-vocab は、delve のような AI 単語にフラグを立てます。 r-uniform-polish は、文章のリズムが均一すぎるとフラグを立てます。アクティブなものは 43 あり (もう 1 つが計画されています)、タイプ別 (表面、リズム、簡潔) および言語ごとにグループ化されています (言語ごとに独自のパックが追加されます)。それらはすべて src/tells/registry.json 内のデータです: ID、一致方法、メッセージ、デフォルトの重大度。エンジン内にハードコードされた情報はありません。
重大度は検出結果ごとに HIGH、MEDIUM、または LOW です。 HIGH は実行に失敗します (出口 1)。 MEDIUM と LOW は報告されますが、報告されません

失敗。それが厳密さのノブです。
プロファイルは、書き込みの種類について、どのテルをどの重大度で実行するかを決定します。プロファイルは JSON ファイルです。つまり、enabled および severity を含む Tell ID のリストと、オプションの Tell ごとのパラメータ (しきい値、単語リスト、除外) です。
実行では、ファイルが読み取られ (コード ブロック、インライン コード、およびフロントマターは無視されます)、プロファイルで有効になっている各項目が実行され、結果が行番号と重大度とともに出力されます。毎回同じ入力、同じ出力。
すべてのチェックは、tell という名前のチェックです。つまり、ある種類のスロップにフラグを立てる正規表現または統計と、それを修正するためのルールとの組み合わせです。私たちがルールを発明したわけではありません。英語のチェックはスティーヴン・キングの『書くことについて』から来ています（副詞を削除し、能動態を好み、派手な単語を削除します）。 Ильяхов と Нора Галь のロシアのパック。前後にいくつか:
これは、アクティブなテル 39 件のうち 6 件です。完全なセットは src/tells/registry.json 内のデータです。プロファイルによって、どれを実行するか、およびどれくらいの負荷で実行するかが決まります。
さまざまなモデルがどのように書くかについて個別のチェックが行われるため、ドラフトに使用したものをすべてキャッチします。
クロードはリズムに頼っています。それは X ではなく、Y のオープナー、大胆な一言格言、エムダッシュです。
GPT は語彙と足場に頼っています: delve / robotics 、 let's dive in 、あなたが... であるかどうか、そして全角ダッシュはほとんどありません。
それぞれの旅行で同じアイデアが異なるチェックを実行します。
$ tacheles チェック claude-draft.md
claude-draft.md（プロフィール：エッセイ）
ハイライン 3 r-reframe-opener これは、フレームワークを選択することではなく、フレームワークを選択することです。
ハイライン 7 r-bold-aphorism 優れたアーキテクチャは構築されません。それは得です。
MEDIUM — 全角ダッシュ密度 2 全角ダッシュ / 36 ワード
不合格 — 高 2、中 1
$ tacheles チェック gpt-draft.md
gpt-draft.md (プロフィール: エッセイ)
ハイライン3 禁用語タペストリー
ハイライン 3 の禁止用語の堅牢性
MEDIUM 3行目 s-gpt-scaffolding 潜ってみよう
MEDIUM ライン 5 s-whet

彼女のオープナー あなたが経験豊富なエンジニアであろうと、
不合格 — 高 2、中 2
これらは法則ではなく傾向です。GPT ドラフトでは依然として em-ダッシュを使用でき、Claude ドラフトでは引き続き delve と言うことができます。タチェレス氏は、あなたの文章を書いたモデルの名前を挙げようとはしません。どちらにしてもスロップにフラグを立てます。
違いを具体的にするために、ここに、blader/humanizer 自身の README の「以前」のエッセイを示します。これは、テルが詰め込まれた LLM ドラフトです。
素晴らしい質問です！このテーマに関するエッセイは次のとおりです。これがお役に立てば幸いです!
AI 支援コーディングは、大規模な言語モデルの変革の可能性の永続的な証拠として機能し、ソフトウェア開発の進化における極めて重要な瞬間を示しています。今日の急速に進化する技術情勢において、研究と実践の交差点に位置するこれらの画期的なツールは、エンジニアがアイデアを出し、反復し、実現する方法を再構築し、現代のワークフローにおける重要な役割を強調しています。
その核心となる価値提案は明確です。それは、プロセスの合理化、コラボレーションの強化、連携の促進です。オートコンプリートだけではありません。それは、創造性を大規模に解き放ち、組織が機敏性を維持しながら、シームレスで直感的で強力なエクスペリエンスをユーザーに提供できるようにすることです。このツールは触媒として機能します。アシスタントはパートナーとして機能します。このシステムはイノベーションの基盤として機能します。
業界観察者らは、趣味の実験から企業全体への展開、個人の開発者から部門を超えたチームに至るまで、導入が加速していると指摘しています。このテクノロジーは、The New York Times、Wired、The Verge で取り上げられています。さらに、ドキュメント、テスト、リファクタリングを生成する機能は、AI がどのようにより良い結果に貢献できるかを示し、自動化と人間の判断の間の複雑な相互作用を強調しています。
- 💡 **速度:** コード生成

大幅に高速化され、摩擦が軽減され、開発者に力が与えられます。
- 🚀 **品質:** トレーニングの改善により出力品質が向上し、より高い基準に貢献します。
- ✅ **採用:** 業界の広範なトレンドを反映して、使用量は増加し続けています。
具体的な詳細は入手可能な情報に基づいて限定されていますが、これらのツールが何らかのプラスの効果をもたらす可能性があると主張される可能性があります。幻覚、偏見、説明責任など、新興テクノロジーに特有の課題にもかかわらず、エコシステムは繁栄し続けています。この可能性を十分に発揮するには、チームは次の点に連携する必要があります。
[切り捨てられた]
$ tacheles チェックエッセイ.md
エッセイ.md（プロフィール：エッセイjp）
ハイライン 1 s-hedge-opener 素晴らしい質問です
ハイライン 3 s-banned-vocab 変形
ハイライン 3 の禁止用語の重要な部分
HIGH ライン 3 の禁止用語の風景
ハイライン 3 k-副詞 急速に
HIGH ライン 5 の禁止用語シームレス
高行 5 r-reframe-opener これはオートコンプリートだけではありません。それは
HIGH line 7 s-ascii のみエンタープライズ全体
HIGH 行 7 s-ascii のみのクロスファンクション
ハイライン 7 禁止用語 複雑
ハイライン 7 副詞
HIGH line 7k-passiveが登場
ハイライン 9 k-副詞 かなり
HIGHライン10k-パッシブ強化
ハイライン 13 k-副詞の可能性
ハイライン 13 k-p

[切り捨てられた]

## Original Extract

Make AI-assisted writing sound like you wrote it, not a model: Tacheles flags the exact AI tells and shows you how to cut them. - guy-lifshitz/tacheles

GitHub - guy-lifshitz/tacheles: Make AI-assisted writing sound like you wrote it, not a model: Tacheles flags the exact AI tells and shows you how to cut them. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
guy-lifshitz
/
tacheles
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits bin bin docs/ reference docs/ reference scripts scripts skills/ tacheles skills/ tacheles src src tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock index.ts index.ts package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
voice-profile
technical-en
Tacheles
An AI-writing linter: it catches the model's tells in your prose, on the exact line, offline.
Make AI-assisted writing sound like you wrote it, not a model. Tacheles flags the exact AI tells in your text and shows you how to cut them.
The slop is two things. Bloat: an LLM writes one token at a time, each the most probable next word, built to sound fluent and average, not short. So it pads, more words than the idea needs. Style: it writes in a register that reads as a machine, the em-dashes, the it's not X, it's Y , the same fifty words. Tacheles flags both, at the exact line, with the reason, and it runs on your machine: no AI, no API key, nothing uploaded, the same result every time.
It does not rewrite for you. It shows you what to cut.
Tacheles (תכל׳ס, tachles ): the bottom line, the point.
From Yiddish. The German Tacheles reden means to talk straight, no fluff.
Before. A paragraph a model would hand you:
In today's landscape, it's not just about the tools, it's about delving into a
robust tapestry of ideas that truly resonate with every reader.
Tacheles flags the tells: delve , robust , tapestry , the it's not X, it's Y
opener, and the padding around them. Cut what it points at and you get:
Pick the tools that fit, then write something worth reading.
Same point, fewer words, none of the tells. It does not write the "after" for you;
it points at what to cut so you can.
What it gives you that the alternatives do not:
A cut list, not a score. AI detectors send your text to a server and hand back "87% AI". You cannot act on a number. Tacheles names the exact span, on the exact line, with the reason, the way a code linter does. You fix it and move on.
The fix, not just the flag. Every check maps to a rule from a working editor: Stephen King for English, Ильяхов and Нора Галь for Russian. The flag says what is wrong; the rule says how to fix it.
Multilingual. Ships with two full language packs (English and Russian) and one experimental (Hebrew); German and Spanish are in the next release. Adding a pack is data, not code, so it supports any language, not just English.
Tuned to you. Calibrate it to your own writing and it flags drift from your voice, not from a generic rule.
Private, free, repeatable. No model, no API key, nothing leaves your machine. Same text, same findings, every time.
Three kinds of tool get pointed at AI slop, and they do different jobs:
Most open-source options are the first two. This is the third: a cut list, multilingual, and tuned to you.
The linter. tacheles check (the cut list), tacheles measure (your writing's stylometry), tacheles compare-drafts (did the rewrite cut at least 10%, King's rule). Offline, deterministic, no API key.
The rewrite skill. A portable SKILL.md that runs the whole loop: check, rewrite against the language's tradition, re-check until clean. It also interviews you and builds your voice anchor. Drop it into Claude Code, Claude.ai, the Agent SDK, or any model.
Four profiles. essay-en , essay-ru , consulting-en-formal , technical-en . Copy one and tune your own.
Language packs. English (Stephen King) and Russian (Ильяхов / Нора Галь): a set of tells plus a rewrite procedure for each.
The guides. Rewrite procedures and the calibration and voice-anchor walkthroughs in docs/reference/ .
npx tacheles check draft.md # run once, no install
npm install -g tacheles # or install it
Or clone the repo and run it with Bun : bun run bin/tacheles check draft.md .
tacheles check draft.md
$ tacheles check draft.md
draft.md (profile: essay-en)
HIGH line 3 s-banned-vocab delve
HIGH line 3 s-banned-vocab robust
HIGH line 7 r-reframe-opener It's not about the tools you pick, it's
MEDIUM line 3 s-gpt-scaffolding Let's dive
MEDIUM — s-em-dash-density 3 em-dashes / 45 words
FAIL — 3 HIGH, 2 MEDIUM
One finding per line, on the exact line, the way a code linter reports. Add --json for machine-readable output to pipe into CI or another tool.
Exit code is 0 when clean and 1 when there is a HIGH finding, so you can gate it in CI or a commit hook. HIGH fails; MEDIUM and LOW are reported but do not fail.
There is also tacheles compare-drafts <old> <new> , which checks that a rewrite cut at least 10% (King's rule).
Tells (the detectors) are the individual checks. Each one is a named pattern, a regex or a statistic, that flags one kind of slop. s-banned-vocab flags AI words like delve ; r-uniform-polish flags overly even sentence rhythm. There are 43 active (one more is planned), grouped by type (surface, rhythm, concision) and by language (each language adds its own pack). All of them are data in src/tells/registry.json : an id, how it matches, its message, and a default severity. No tell is hard-coded in the engine.
Severity is HIGH, MEDIUM, or LOW per finding. HIGH fails the run (exit 1); MEDIUM and LOW are reported but never fail. That is the strictness knob.
Profiles decide which tells run, and at what severity, for a kind of writing. A profile is a JSON file: a list of tell ids with enabled and severity , plus optional per-tell params (thresholds, word-lists, exclusions).
A run reads the file (ignoring code blocks, inline code, and frontmatter), executes each tell the profile enables, and prints the findings with line numbers and severities. Same input, same output, every time.
Every check is one named tell: a regex or a statistic that flags one kind of slop, paired with the rule for fixing it. We did not invent the rules. The English checks come from Stephen King's On Writing (kill the adverbs, prefer the active voice, cut the fancy word); the Russian pack from Ильяхов and Нора Галь. A few, before and after:
That is 6 of 39 active tells; the full set is data in src/tells/registry.json . Profiles decide which of them run, and how hard.
It has separate checks for how different models write, so it catches the slop whatever you drafted with.
Claude leans on rhythm: it's not X, it's Y openers, bold one-liner aphorisms, em-dashes.
GPT leans on vocabulary and scaffolding: delve / robust , let's dive in , whether you're a... , and almost no em-dashes.
The same idea from each trips different checks:
$ tacheles check claude-draft.md
claude-draft.md (profile: essay-en)
HIGH line 3 r-reframe-opener It's not about the framework you choose, it's
HIGH line 7 r-bold-aphorism Good architecture isn't built. It's earned.
MEDIUM — s-em-dash-density 2 em-dashes / 36 words
FAIL — 2 HIGH, 1 MEDIUM
$ tacheles check gpt-draft.md
gpt-draft.md (profile: essay-en)
HIGH line 3 s-banned-vocab tapestry
HIGH line 3 s-banned-vocab robust
MEDIUM line 3 s-gpt-scaffolding Let's dive
MEDIUM line 5 s-whether-opener Whether you're a seasoned engineer or
FAIL — 2 HIGH, 2 MEDIUM
These are tendencies, not laws: a GPT draft can still use em-dashes, a Claude draft can still say delve . Tacheles does not try to name the model that wrote your text. It flags the slop either way.
To make the difference concrete, here is the "before" essay from blader/humanizer 's own README: an LLM draft packed with tells.
Great question! Here is an essay on this topic. I hope this helps!
AI-assisted coding serves as an enduring testament to the transformative potential of large language models, marking a pivotal moment in the evolution of software development. In today's rapidly evolving technological landscape, these groundbreaking tools—nestled at the intersection of research and practice—are reshaping how engineers ideate, iterate, and deliver, underscoring their vital role in modern workflows.
At its core, the value proposition is clear: streamlining processes, enhancing collaboration, and fostering alignment. It's not just about autocomplete; it's about unlocking creativity at scale, ensuring that organizations can remain agile while delivering seamless, intuitive, and powerful experiences to users. The tool serves as a catalyst. The assistant functions as a partner. The system stands as a foundation for innovation.
Industry observers have noted that adoption has accelerated from hobbyist experiments to enterprise-wide rollouts, from solo developers to cross-functional teams. The technology has been featured in The New York Times, Wired, and The Verge. Additionally, the ability to generate documentation, tests, and refactors showcases how AI can contribute to better outcomes, highlighting the intricate interplay between automation and human judgment.
- 💡 **Speed:** Code generation is significantly faster, reducing friction and empowering developers.
- 🚀 **Quality:** Output quality has been enhanced through improved training, contributing to higher standards.
- ✅ **Adoption:** Usage continues to grow, reflecting broader industry trends.
While specific details are limited based on available information, it could potentially be argued that these tools might have some positive effect. Despite challenges typical of emerging technologies—including hallucinations, bias, and accountability—the ecosystem continues to thrive. In order to fully realize this potential, teams must align with
[truncated]
$ tacheles check essay.md
essay.md (profile: essay-en)
HIGH line 1 s-hedge-opener Great question
HIGH line 3 s-banned-vocab transformative
HIGH line 3 s-banned-vocab pivotal
HIGH line 3 s-banned-vocab landscape
HIGH line 3 k-adverbs rapidly
HIGH line 5 s-banned-vocab seamless
HIGH line 5 r-reframe-opener It's not just about autocomplete; it's
HIGH line 7 s-ascii-only enterprise-wide
HIGH line 7 s-ascii-only cross-functional
HIGH line 7 s-banned-vocab intricate
HIGH line 7 k-adverbs Additionally
HIGH line 7 k-passive been featured
HIGH line 9 k-adverbs significantly
HIGH line 10 k-passive been enhanced
HIGH line 13 k-adverbs potentially
HIGH line 13 k-p

[truncated]
