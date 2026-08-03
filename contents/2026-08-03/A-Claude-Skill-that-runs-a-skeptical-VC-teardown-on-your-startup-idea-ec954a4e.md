---
source: "https://github.com/zszendro/vc-teardown"
hn_url: "https://news.ycombinator.com/item?id=49161394"
title: "A Claude Skill that runs a skeptical-VC teardown on your startup idea"
article_title: "GitHub - zszendro/vc-teardown: A Claude Skill that stress-tests your startup idea as a skeptical VC, then rebuilds it into a business plan with a defensible moat, MVP scope, GTM and kill criteria. · GitHub"
author: "zszendro"
captured_at: "2026-08-03T21:59:27Z"
capture_tool: "hn-digest"
hn_id: 49161394
score: 2
comments: 0
posted_at: "2026-08-03T21:02:44Z"
tags:
  - hacker-news
  - translated
---

# A Claude Skill that runs a skeptical-VC teardown on your startup idea

- HN: [49161394](https://news.ycombinator.com/item?id=49161394)
- Source: [github.com](https://github.com/zszendro/vc-teardown)
- Score: 2
- Comments: 0
- Posted: 2026-08-03T21:02:44Z

## Translation

タイトル: あなたのスタートアップのアイデアに対して懐疑的なVCの分解を実行するクロード・スキル
記事のタイトル: GitHub - zszendro/vc-teardown: 懐疑的な VC としてスタートアップのアイデアをストレス テストし、それを防御可能な堀、MVP スコープ、GTM、キル基準を備えたビジネス プランに再構築するクロード スキル。 · GitHub
説明: 懐疑的な VC としてスタートアップのアイデアをストレス テストし、防御可能な堀、MVP スコープ、GTM、およびキル基準を備えたビジネス プランに再構築するクロード スキル。 - zszendro/vc-teardown

記事本文:
GitHub - zszendro/vc-teardown: 懐疑的な VC としてスタートアップのアイデアをストレス テストし、防御可能な堀、MVP スコープ、GTM、キル基準を備えたビジネス プランに再構築するクロード スキル。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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

ツェンドロ
/
VC 分解
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット vc-teardown vc-teardown .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード スキルは、懐疑的な VC としてあなたのスタートアップのアイデアをストレス テストし、防御可能な堀を備えた計画に再構築します。
ビジネスアイデアを 1 行または 10 ページで説明します。このスキルは、すでにそれを構築した人を調査し、本物の投資家が提起するであろう反対意見で攻撃し、既存の企業が構造的に受け入れられないくさびを見つけて、ビジネスプラン、MVP の範囲、市場投入とキルの基準を作成します。
ほとんどのアイデアが失敗するのは、アイデアが悪いからではなく、明白なバージョンがすでに構築されているからです。このスキルは、6 か月ではなく午後 1 日でそれを見つけられるように設計されています。
1. 摂取 – あらゆる詳細情報を摂取します。アンケートはありません。質問は最大 3 つまでで、分析を変える質問のみです。
2. 調査 — 直接の既存企業、無料の代替品、購入者がすでに支払っている隣接ソフトウェア、および痛みが本物である証拠を検索します。このステップは交渉の余地のないものです。このステップをスキップすると、一般的な起動アドバイスが生成されます。
3. 分解 — 16 の攻撃対象領域のライブラリから抽出された 8 ～ 12 の番号付きの課題。すべての課題は具体的な修正で終わります。計画を変更しない批判は意味がありません。
4. 堀 — 再フレーム。防御可能性を見つけるための 9 つのパターン (ほとんどの作業を行う 2 つを含む): データを集計する代わりに生成すること、および既存企業の指標の不一致を見つけること。
5. 計画 - 市場の現実、階層化されたユースケース、収益ラインがいつオンになるかを順序付け、実際の演算によるユニットエコノミクス、MVP

コープアンドスタック、GTM、リスク登録、および改ざん可能な殺害基準。
6. 納品 — 書面による計画と、元のアイデアのすべての要素をその着地点にマッピングした付録。
他のVCスキルとの違い
この空間にはクロードのスキルがいくつかあります。それらのほとんどは、提案資料、資金調達の物語、評価している企業の名前など、すでに作成したものを採点します。特定の投資家の哲学をシミュレートして、既存のストーリーを評価するものもあります。
こちらの方が早く始まり、終わり方が異なります。生のアイデアを採用し、その市場ですでに構築されているものを確立し、失われたバージョンを破棄して、別の堀の周りにアイデアを再構築します。分解は成果物ではなく、再構築が成果物です。
デッキを持っていてそれを評価したい場合は、他のデッキのいずれかを使用してください。アイデアがあり、その明白なバージョンがすでに消滅しているかどうかを知りたい場合は、これを使用してください。
3 つの制約が機能します。
それは数字を実行します。 1 つの完全に浸透した橋頭堡が 18,000 ドルの収益を生み出した場合、その数字が記録されます。特にその時は。
正直な規模の評決を述べています。わかりやすい言葉で言えば、ライフスタイル ビジネス、ブートストラップ SaaS、またはベンチャー規模です。 SMB SaaS をベンチャー市場として装うことは、勤勉さを怠り、1 年を無駄にしてしまいます。
それは最も安価なテストを指します。最大の仮定のリスクを回避する実験には、通常、お金ではなく時間がかかります。そして、そのテストに合格するまで何を構築すべきではないかが明確に示されています。
また、スポーツ向けに逆張りにならないように調整されています。課題に良い答えがある場合は、そう言って次に進みます。でっち上げられた異論は、実際の異論に対する信頼性を損ないます。
クロード アプリ — リリースから vc-teardown.skill をダウンロードし、会話内のファイル カードで [スキルを保存] をクリックします。
Claude Code — スキル ディレクトリにクローンを作成します。
git clone https://github.com/zszendro/vc-teardown.git ~

/.claude/skills/vc-teardown
手動 — セットアップがスキルをロードする場所に vc-teardown/ フォルダーをコピーします。
アイデアを説明するだけです。スキルは自動的に発動します。
フリーランスの CAD デザイナーのためのマーケットプレイスを構築したいと考えています。
これに穴を開けるのは、水族館愛好家のための定期購入ボックスです。
これが私の提案資料の概要です。VC の役割を果たして、何が問題なのか教えてください。
自治体の許可申請書を作成する AI ツールに堀はありますか?
また、「このアイデアを検証してください」、「競争相手はどのようなものか」、「堀を見つけてください」、「ビジネスプランを書いてください」といった場合にもトリガーされます。
分析が変わるため、前もって 2 つのことを伝えておく価値があります。それは、 をどこで起動するか、そして、 をブートストラップするかレイズするかです。
vc-分解/
§── SKILL.md # ワークフロー + トーンキャリブレーション
━── 参考文献/
§──Challenge-library.md # 16 攻撃対象領域
§── moat-patterns.md #9 再フレーム
└── plan-template.md # 出力構造体
参照ファイルは必要な場合にのみロードされるため、スキルは実際に動作するまでコンテキスト内で安価なままになります。
競争 · 無料の代替品 · 収益化 · 支払い意欲 · コールドスタート · 流通 · 購入者のインセンティブ · 販売サイクル · データの可用性 · 規制と責任 · 創業者と市場の適合性 · 資本効率 · 防御性 · 企業ではない特徴 · 保持 · 規模の誠実さ
コモディティを格下げする · データを集計せずに生成する · マーケティング上の露出ではなく、運用上の救済を販売する · 既存企業の指標の不一致 · 付与された分配が獲得した分配を上回る · サービスを受けていない 80% · 制度の切り替えコスト · 水平方向の範囲よりも垂直方向の深さ · 堀を順番に並べる
投資アドバイスではなく、顧客との対話の代わり、市場を知るオペレーターの代わりでもありません。構造化されたセコです

コードを書く前に得られる意見。すでに構築されているものや、明らかな計画が崩れている箇所を見つけるのが得意ですが、確信の源としては役に立ちません。
出力は、キル基準が付加された仮説です。ポイントはテストに行くことです。
参照ファイルは興味深い部分です。新しい攻撃対象領域、新しい堀のパターン、実際の分解で実際に動作した例はすべて歓迎されます。特に、現在のライブラリがうまく扱っていないカテゴリ (ハードウェア、バイオテクノロジー、規制された金融、ディープテクノロジー) からのものです。
クロード スキルは、懐疑的な VC としてスタートアップのアイデアをストレス テストし、それを防御可能な堀、MVP スコープ、GTM、キル基準を備えたビジネス プランに再構築します。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Claude Skill that stress-tests your startup idea as a skeptical VC, then rebuilds it into a business plan with a defensible moat, MVP scope, GTM and kill criteria. - zszendro/vc-teardown

GitHub - zszendro/vc-teardown: A Claude Skill that stress-tests your startup idea as a skeptical VC, then rebuilds it into a business plan with a defensible moat, MVP scope, GTM and kill criteria. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
zszendro
/
vc-teardown
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit vc-teardown vc-teardown .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
A Claude Skill that stress-tests your startup idea as a skeptical VC — then rebuilds it into a plan with a defensible moat.
Describe a business idea in one line or ten pages. The skill researches who already built it, attacks it with the objections a real investor would raise, finds the wedge the incumbent structurally won't take, and writes the business plan, MVP scope, go-to-market and kill criteria.
Most ideas fail not because they're bad, but because the obvious version is already built. This skill is designed to find that out in an afternoon instead of six months.
1. Intake — takes whatever detail you have. No questionnaire; at most three questions, and only ones that change the analysis.
2. Research — searches for direct incumbents, free substitutes, adjacent software your buyer already pays for, and evidence the pain is real. This step is non-negotiable: skipping it produces generic startup advice.
3. Teardown — 8–12 numbered challenges drawn from a library of 16 attack surfaces. Every challenge ends in a concrete revision. Criticism that doesn't change the plan doesn't count.
4. Moat — the reframe. Nine patterns for finding defensibility, including the two that do most of the work: generate the data instead of aggregating it , and find the incumbent's metric mismatch .
5. Plan — market reality, tiered use cases, revenue lines ordered by when they turn on, unit economics with the actual arithmetic, MVP scope and stack, GTM, risk register, and falsifiable kill criteria.
6. Delivery — a written plan, plus an appendix mapping every element of your original idea to where it landed.
How it differs from other VC skills
There are several Claude skills in this space. Most of them score something you've already built — a pitch deck, a fundraising narrative, a named company you're evaluating. Some simulate specific investors' philosophies to grade your existing story.
This one starts earlier and ends differently. It takes a raw idea, establishes what's already been built in that market, kills the version that loses, and rebuilds the idea around a different moat . The teardown isn't the deliverable — the reframe is.
If you have a deck and want it graded, use one of the others. If you have an idea and want to know whether the obvious version of it is already dead, use this.
Three constraints do the work:
It runs the numbers. If one fully-penetrated beachhead produces $18k of revenue, it writes that number. Especially then.
It states the honest scale verdict. Lifestyle business, bootstrap SaaS, or venture-scale — in plain words. Dressing an SMB SaaS as a venture marketplace fails diligence and wastes a year.
It names the cheapest test. The experiment that de-risks the biggest assumption usually costs time, not money — and it says explicitly what shouldn't be built until that test passes.
It's also calibrated not to be contrarian for sport. If a challenge has a good answer, it says so and moves on. Manufactured objections cost credibility on the real ones.
Claude apps — download vc-teardown.skill from Releases and click Save skill on the file card in a conversation.
Claude Code — clone into your skills directory:
git clone https://github.com/zszendro/vc-teardown.git ~ /.claude/skills/vc-teardown
Manual — copy the vc-teardown/ folder into wherever your setup loads skills from.
Just describe the idea. The skill triggers on its own:
I'm thinking about building a marketplace for freelance CAD designers.
Poke holes in this: a subscription box for aquarium hobbyists.
Here's my pitch deck outline — act as a VC and tell me what's wrong with it.
Is there a moat in an AI tool that drafts municipal permit applications?
It also triggers on "validate this idea," "what does the competition look like," "find the moat," and "write me a business plan."
Two things worth telling it up front, since they change the analysis: where you'd launch , and whether you're bootstrapping or raising .
vc-teardown/
├── SKILL.md # workflow + tone calibration
└── references/
├── challenge-library.md # 16 attack surfaces
├── moat-patterns.md # 9 reframes
└── plan-template.md # output structure
Reference files load only when needed, so the skill stays cheap in context until it's actually working.
Competition · Free substitutes · Monetization · Willingness to pay · Cold start · Distribution · Buyer incentive · Sales cycle · Data availability · Regulatory & liability · Founder-market fit · Capital efficiency · Defensibility · Feature-not-a-company · Retention · Scale honesty
Demote the commodity · Generate the data, don't aggregate it · Sell operational relief, not marketing exposure · The incumbent's metric mismatch · Granted distribution beats won distribution · The unserved 80% · Institutional switching cost · Vertical depth over horizontal reach · Sequence the moat
Not investment advice, not a substitute for talking to customers, and not a replacement for an operator who knows your market. It's a structured second opinion that arrives before you write code — good at finding what's already built and where the obvious plan breaks, and useless as a source of conviction.
The output is a hypothesis with kill criteria attached. The point is to go test it.
The reference files are the interesting part. New attack surfaces, new moat patterns, and worked examples from real teardowns are all welcome — especially from categories the current library handles badly (hardware, biotech, regulated finance, deep tech).
A Claude Skill that stress-tests your startup idea as a skeptical VC, then rebuilds it into a business plan with a defensible moat, MVP scope, GTM and kill criteria.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
