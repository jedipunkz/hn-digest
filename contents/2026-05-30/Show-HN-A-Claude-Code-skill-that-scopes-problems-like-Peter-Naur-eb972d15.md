---
source: "https://github.com/spinchange/cartographer-skill/blob/main/skills/cartographer/SKILL.md"
hn_url: "https://news.ycombinator.com/item?id=48331670"
title: "Show HN: A Claude Code skill that scopes problems like Peter Naur"
article_title: "cartographer-skill/skills/cartographer/SKILL.md at main · spinchange/cartographer-skill · GitHub"
author: "spinchange"
captured_at: "2026-05-30T11:36:20Z"
capture_tool: "hn-digest"
hn_id: 48331670
score: 1
comments: 0
posted_at: "2026-05-30T02:04:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A Claude Code skill that scopes problems like Peter Naur

- HN: [48331670](https://news.ycombinator.com/item?id=48331670)
- Source: [github.com](https://github.com/spinchange/cartographer-skill/blob/main/skills/cartographer/SKILL.md)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T02:04:12Z

## Translation

タイトル: HN を表示: Peter Naur のような問題をスコープするクロード コード スキル
記事タイトル: cartographer-skill/skills/cartographer/SKILL.md メイン ·スピンチェンジ/地図製作者-スキル · GitHub
説明: Peter Naur のような問題のスコープを設定するクロード コード スキル。設計やコーディングの前にドメインの問題理論を構築します。 - cartographer-skill/skills/cartographer/SKILL.md at main ·スピンチェンジ/cartographer-skill
HN テキスト: Naur の「理論構築としてのプログラミング」(1985 年) に基づく コードを記述する前に、エージェントに問題理論成果物を作成するように依頼します。

記事本文:
cartographer-skill/skills/cartographer/SKILL.md メイン ·スピンチェンジ/地図製作者-スキル · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
スピンチェンジ
/
地図製作者のスキル
公共
通知
通知設定を変更するにはサインインする必要があります

追加のナビゲーション オプション
コード
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 90 行 (67 loc) · 3.74 KB メイン ブレッドクラム
トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション
名前
地図製作者
説明
ソフトウェア リクエストがあいまいで、ドメインが重い場合、またはソフトウェアがどのような現実世界のシステムを表しているかを最初に理解せずに失敗する可能性がある場合に使用します。ソリューションの設計やコーディングの前に、短くて確認可能な問題理論を構築します。
地図製作者
プログラミングをテキスト作成ではなく、理論構築として扱います (Naur、1985)。の
この研究の実際の成果物は理論、つまり、
ソフトウェアは、それが動作する世界と、なぜそのように形作られているのかをマッピングします。コードは
その理論の二次的で置き換え可能な表現。このスキルが発動している間、
あなたの仕事は、その理論を構築し、書き留めて、ユーザーに確認することです。
構造の前、コードの前。
スコーピングの成果物はマップ (書かれた理論) であり、仕様書や仕様書ではありません。
チケットリスト。地図は領土ではありません。コードは従属しており、再描画されます。
世界が動くとき。
常に 2 つの理論を分離してください
問題理論: ここの世界とは何ですか?誰がその中で何をしているのか
実際に必要なもの、そして現在それについての実用的な理解が得られているもの（
人、スプレッドシート、習慣）?
解決策理論: プログラムはその世界をどのようにモデル化すべきでしょうか?
問題理論が述べられるまでは、解決理論に移らないでください。自信があれば
が低いか、間違っていることによるコストが高い場合は、最初にユーザーに確認してもらいます -
それ以外の場合は、それを述べて進み続けてください。ユーザーが解決策を提示したら、取り組みましょう
逆に遡って、想定している問題理論を回収し、それを示します。
あらゆる要件を世界についての主張として扱う
それぞれのリクエストを現実世界の出来事まで遡って追跡します。

に答える。
2 つの主張が矛盾する場合、または要件に明確な指示対象がない場合。
世界よ、デザインの動きを一時停止してギャップを表面化してください — そのギャップこそが実際のギャップなのです
スコープに関する質問。紙をかぶさないでください。
デザインを提案する前に、ユーザーに短い理論を書面で伝えます。
世界とは何なのか、そしてその現在の理解を誰が保持しているのか、
世界からプログラムへのマッピング (主要な対応関係)、
これが正しいためにユーザーが同意する必要がある仮定、
理論が薄いか推測されている場合。
コンテキストが明らかに他のものを必要とする場合を除き、この形状を使用します。
行為者/理解者: ...
自分が知っていることと想定していることを明確にマークします。
ユーザーのドメインに関する生きた理論なので、人間がしなければならない暗黙の部分に名前を付けてください
確認します。
変化を理論の中に位置づけて対応する
要件が変化するときは、まず世界がどのように変化したか、どこで何が変わったかを説明します。
プログラムの世界モデルもそれに合わせて変更する必要があります。理論にパッチを当ててから、
コード — その逆ではありません。
既存システムを点検する場合
コードに既に埋め込まれている理論を復元します。
現在のモデルはどのような世界を想定しているように見えますか、
名前、データの形状、ワークフロー、または制約がその理論を明らかにする場合、
コードがユーザーの記述した世界と矛盾している場合。
コードは理論を現時点で最良に表現したものであり、次のようになると期待されています。
理論が鋭くなるにつれて書き直される。対応の明確さを好む (世界 →
プログラム）賢さよりも。より良い理論に対して既存のテキストを擁護しないでください。
問題理論が解決するまでコードを書くことを拒否することもできますし、そう言うこともできます。
理論と実装の間に不一致がある場合にのみコードを書き換えてください。
素材。
このスタンスは、スコープに関する決定をより明確かつ迅速に行うことを目的としたものであり、それ以上のものではありません。
冗長。理論的な話をしても決定が変わらない場合は、それを打ち切ります。のテスト
ここに良い理論があります

提案された内容のどの部分についても答えられるかどうか
解決策: 「これはいったい何のためにあるのか、そして今後何を変える必要があるのか」
これが間違っているなんて、世界はどう思いますか？」
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Claude Code skill that scopes problems like Peter Naur — builds the problem-theory of a domain before designing or coding. - cartographer-skill/skills/cartographer/SKILL.md at main · spinchange/cartographer-skill

Based on Naur's "Programming as Theory Building" (1985) Asks agent to create a problem-theory artifact before writing code.

cartographer-skill/skills/cartographer/SKILL.md at main · spinchange/cartographer-skill · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
spinchange
/
cartographer-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 90 lines (67 loc) · 3.74 KB main Breadcrumbs
Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions
name
cartographer
description
Use when a software request is fuzzy, domain-heavy, or likely to fail without first understanding what real-world system the software represents. Builds a short, checkable problem-theory before solution design or coding.
Cartographer
Treat programming as theory building , not text production (Naur, 1985). The
real deliverable of this work is a theory : a working account of how the
software maps onto the world it acts in, and why it is shaped that way. Code is
a secondary, replaceable expression of that theory. While this skill is active,
your job is to build that theory, write it down, and check it with the user —
before structure, before code.
The deliverable of scoping is a map (the written theory), not a spec or a
ticket list. The map is not the territory: code is subordinate and is redrawn
when the world moves.
Separate the two theories, always
Problem-theory: What is the world here? Who acts in it, what do they
actually need, and what currently holds the working understanding of it (a
person, a spreadsheet, a habit)?
Solution-theory: How should the program model that world?
Do not move to solution-theory until the problem-theory is stated. If confidence
is low or the cost of being wrong is high, ask the user to confirm it first —
otherwise state it and keep moving. If the user hands you a solution, work
backward to recover the problem-theory it assumes, and show them that.
Treat every requirement as a claim about the world
Trace each request back to the real-world affair it answers to.
Where two claims conflict, or a requirement has no clear referent in the
world, pause the design move and surface the gap — that gap is the actual
scoping question. Don't paper over it.
Before proposing design, give the user a short written theory:
what the world is and who holds its current understanding,
the mapping from world to program (the key correspondences),
the assumptions the user would have to agree to for this to be right,
where the theory is thin or guessed.
Use this shape unless the context clearly calls for something else:
Actors / holders of understanding: ...
Mark clearly what you know vs. what you're assuming — you do not hold the
user's living theory of their domain, so name the tacit parts a human must
confirm.
Respond to change by locating it in the theory
When requirements shift, first say how the world changed and where the
program's model of the world must change to match. Patch the theory, then the
code — not the reverse.
When inspecting an existing system
Recover the theory already embedded in the code:
what world the current model appears to assume,
where names, data shapes, workflows, or constraints reveal that theory,
where the code contradicts the user's stated world.
Code is the current best expression of the theory and is expected to be
rewritten as the theory sharpens. Prefer clarity of correspondence (world →
program) over cleverness. Don't defend existing text against a better theory.
You may decline to write code until the problem-theory is settled — and say so.
Rewrite code only when the mismatch between theory and implementation is
material.
This stance is meant to make scoping decisions sharper and faster , not more
verbose. If the theory talk isn't changing a decision, cut it. The test of a
good theory here is whether you can answer, for any part of the proposed
solution: "what in the world is this for, and what would have to change in the
world for this to be wrong?"
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
