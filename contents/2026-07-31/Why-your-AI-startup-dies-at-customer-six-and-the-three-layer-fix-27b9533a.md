---
source: "https://github.com/teal-sea/why-your-ai-startup-dies-at-customer-six"
hn_url: "https://news.ycombinator.com/item?id=49127950"
title: "Why your AI startup dies at customer six (and the three-layer fix)"
article_title: "GitHub - teal-sea/why-your-ai-startup-dies-at-customer-six: An essay on why bespoke deployments stop scaling, and the semantic / kinetic / rulebook split that fixes it. · GitHub"
author: "tlince"
captured_at: "2026-07-31T20:54:37Z"
capture_tool: "hn-digest"
hn_id: 49127950
score: 2
comments: 1
posted_at: "2026-07-31T20:04:19Z"
tags:
  - hacker-news
  - translated
---

# Why your AI startup dies at customer six (and the three-layer fix)

- HN: [49127950](https://news.ycombinator.com/item?id=49127950)
- Source: [github.com](https://github.com/teal-sea/why-your-ai-startup-dies-at-customer-six)
- Score: 2
- Comments: 1
- Posted: 2026-07-31T20:04:19Z

## Translation

タイトル: AI スタートアップが顧客 6 で消滅する理由 (および 3 層の修正)
記事のタイトル: GitHub - teal-sea/why-your-ai-startup-dies-at-customer-six: オーダーメイドのデプロイメントがスケーリングを停止する理由と、それを修正するセマンティック/キネティック/ルールブックの分割に関するエッセイ。 · GitHub
説明: オーダーメイドのデプロイメントがスケーリングを停止する理由と、それを修正するセマンティック/キネティック/ルールブックの分割についてのエッセイ。 - ティールシー/Why-your-ai-startup-dies-at-customer-6

記事本文:
GitHub - teal-sea/why-your-ai-startup-dies-at-customer-six: オーダーメイドのデプロイメントがスケーリングを停止する理由と、それを修正するセマンティック / キネティック / ルールブックの分割に関するエッセイ。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ティールシー
/
なぜあなたのAIをスタートするのか

顧客6でアップダイス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI スタートアップが顧客 6 で消滅する理由 (および 3 層の修正)
顧客 1 は簡単です。彼らのルールは製品です。つまり、彼らの預金ポリシーをハードコード化すると、
時間、彼らの癖、そしてデモは魔法のように見えます、なぜならそれは彼ら自身のビジネスのようなものだからです
彼らに戻って。
顧客 2 は 3 つの異なる点を必要としていますが、あなたは忙しいので、例外はどこに行きますか
コードはすでにあります。ワークフローの分岐、プロンプトの段落、特殊なケース
ツール。顧客 6 のそばにはクライアント名が記載されたフォルダーがあり、オンボーディングが始まります
数週間以内に開発に着手します。
遠くからこれを書いているわけではありません。私は米国のリムジン予約サービス、AI の運営を行っています。
ライブ通話で電話に応答します。私は最近、次のようなルールを探しました。
承認されたデポジットがなければ予約を確認することはできません。私のシステムでは、
状態の数秒前に決済処理業者に対してカード承認が検証されたことを意味します
変化します。このルールは 3 か所で見つかりました。1 つは予約ワークフローのバージョン、もう 1 つは予約ワークフローのバージョンです。
エージェントの予約ツール、ダッシュボードの検証に 1 つ、月ごとに 3 つずつ書き込まれます
私のさまざまな気分。彼らはほぼ同意した。まだ何も爆発していませんでした、そしてそれは
恐ろしい点: 何もチェックしていませんでした。ルールを変更すると、ミスしたコピーによってエラーがスローされなくなります。
エラーです。何か高価になるまで、実際の顧客に対して古いルールを適用し続けるだけです
が起こります。そして、テナントは私 1 人だけです。ここで、6 つのクライアント名が含まれるフォルダーを想像してください。
それ。
この誘惑は数字で説明できます。チャートモーグル

の定着率レポートには AI ネイティブ企業が含まれています
純収益維持率の中央値は 48% で、通常の B2B SaaS の 82% に対して、
興味深いのはセグメンテーションです。月額 250 ドルを超えると、AI の定着率は次のようになります。
通常の SaaS も 85% です。 ChartMogul 自身の解釈によると、その層に残っているのは次のとおりです。
ワークフローへのより深い適合とより緊密な統合、そしてそれさえも緩やかに維持します。価格も重要です。
企業バイヤーの代理、年間契約と調達の慣性、および AI ネイティブ
コホートでは消費者製品を比較に混ぜます。それにしても奥が深い
ワークフローに適合させる最も簡単な方法は次のとおりであるため、ワークフローに適合させることはまさに誘惑です。
クライアントのあらゆる癖をコードに吸収し、そこから離れることはできません。なしでやってください
その下にプラットフォームがあり、a16z のマーク・アンドルスコはすでに着陸場所に名前を付けています:「数千」
保守やアップグレードが不可能なオーダーメイドの展開。」そのとき、彼の中では、
つまり、あなたは X の Palantir ではありません。より優れたフロントエンドを備えた「X の Accenture」なのです。
診断は 1 つの文です。ルールはそれを所有するものの中に存在しません。 3
デポジットルールのコピーが漂流し、エージェントがデポジットステップなしで本に手を伸ばす
2 つの衣装で同じバグです。修正はレイヤーに名前を付けることから始まります。
セマンティック層は物事の本質です。予約、車両、料金、顧客：係員1名
それぞれ、1 セットの法的状態を形成します。名詞。私の 3 つの漂流コピーは、とりわけ、
物事、法的状態と呼ばれる 3 つの私的な定義が確認されました。
それができるのがキネティックレイヤーです。見積もり、予約、キャンセル: ビジネスの動詞
そこにはルールがあり、それ以外にテーブルへの扉はありません。デポジットルールは本の中に存在します。
すべての呼び出し元に対して 1 回、動詞ではなく状態を保護します。
予約の確認方法を学びます

、手動オーバーライド、インポート、同じルールを継承します
同じドアで。
ルールブック層では、レート、しきい値、
スケジュール、切り替え。動詞が実行時に読み取るクライアントごとのルールブックとその名前
少しズルをします。ルールは動詞の中に生きています。ルールブックには、各クライアントの答えが記載されています。
彼ら。
ギリシャ語は、Palantir のオントロジー、つまり、
エンタープライズ、その名詞と動詞、そして最初の 2 つの層も次のような匂いがするかどうか
ドメイン駆動設計、それは次のような理由からです: エヴァンスをまっすぐに貫く血統
2003 年。ルールブックはさらに目新しいものではありません。すべての SaaS アプリにはテナント設定テーブルがあり、
すべての機能フラグ サービスは、マーケティング予算を伴うルールブックです。ニュースはそうではありません
ファイリングシステム。ニュースはルールブックを正直に保つものです。
発信者を尊重し、何も強制しないというのはルールブックとは名ばかりで、ほとんどがルールブックです。
2 つのことがこの問題を現実にしており、それらがこのエッセイの残りの部分です。
偽物と、侵入するものに対する規律。
a16z がすでにこの会話に「10 年」というタイトルを付けて以来、パランティアのアークについての私の読みを簡単に説明します。
標準化されるまではサービス会社のように見えたオーダーメイドのインストール
クライアントの名詞、動詞、および動詞を宣言するためのマシンが 1 つ上のレベルであることが判明しました。
ルール。デプロイメントはオーサリングになり、インストール時間は S-1 までに 14 日に短縮されました。
あなたにとって重要なのは、垂直型オペレーターはこれまでの Palantir よりも簡単であるということです
した: 彼らは、あらゆる企業の名詞に対応できる汎用性の高いマシンを必要としていましたが、企業は
1 つのバーティカルにサービスを提供すると、名詞や動詞自体が見つかると凍結される可能性があります。
そして各クライアントにルールブックのみを渡します。
これが私が自分のシステムで学ばなければならなかった部分です: あなたは decl をしません

あなたのプリミティブです。あなた
ルールブックはそれらを見つけるための道具です。
2 つのプローブによって強制的に検出が行われます。
顧客 2 が最初のプローブです。顧客がこれまでに行ったことのないリクエストはすべて、
並べ替え: 見逃したパラメータ (ルールブック)、見逃した動詞 (キネティック)、新しい種類の
完全に（意味論的に）、または拒否。一部のリクエストは複数のバケットにピースを配置します。
仕分けはまだ作業中です。ソフトウェア製品ラインの担当者はこれを次のように教えていました。
1998 年の共通性と変動性分析。スタートアップ バージョンは次の文に収まります。
ルールブックはふるいであり、それを通過しないものはすべて製品です。
AI エージェントは 2 番目のプローブであり、何か別のものをテストします。お客様による 2 つのテスト
システムの一部は共有されます。エージェントは共有部分が実際に運ぶかどうかをテストします
自分たちのルール。ワークフローは、スクリプトに記述した順序でオペレーションを呼び出します。
不変式は、あなたが書こうと考えたシーケンスにのみ直面します。エージェントが動詞を作成します
実行時に、会話ごとに、誰もスクリプト化していない順序で、何もせずに本にアクセスします。
会話の途中で納得のいく理由で、入金手続きは午前2時に行われました。
セキュリティ担当者は、思いもよらない入力をシステムに供給するマシンを表す言葉を持っています。
テストするもの: ファザー。エージェントは怠惰なファザーです。それは空間を探索するのではなく、さまよう
そして、助手席に座った顧客と一緒に歩き回ります。すべての呼び出しでサンプリングが行われます
何もテストされていない動詞の順序。エージェントのもとでは、
動詞内に存在しないルール、またはその下の制約層内にあるルールは呼び出し元が実行できません
ルートアラウンドは、スクリプト化された発信者がたまたま共有した習慣にすぎません。その「その下」
実際の重みがあり、私がそれを所有します。一部の不変条件はデータベース自体に属します。
アプリケーションコードがありません

それらを迂回することができます。議論はドアの数に関するものであり、
彼らの高度。何であれ
任意の順序で安全であること自体はプリミティブです。それ以外はすべて振り付けであり、
エージェントはそれを守ると決して約束しなかった。
スケールしないことを内面化している場合は、ここで反対する必要があります。
顧客 1 のハードコーディングが標準的なアドバイスであり、顧客 2 でのプラットフォーム化は
時期尚早な抽象化。同意します。ふるいは何も構築しないように求めていることに注意してください。それは
プラットフォームではなく、仕分けの規律です。顧客 2 以降は、台帳で実行できます。
クライアントの違いはパラメータ、動詞の欠落、新しいオブジェクト、または拒否として記録されます。
コードは好きなだけハードコーディングされたままになります。あなたが買うのは、最終的には
プラットフォームを抽出してください。プリミティブについては推測していません。あなたは2年分の本を読んでいます
整理された証拠。時期尚早な抽象化は、sieve を実行する前にその出力を推測することです。
これは、初日から紙の上で実行しているだけです。
2000 年代のルール エンジンは「ビジネス ユーザー」として販売されていたため、今では大きな負担となっています。
ロジックを書いてください」と死んだのと同じように、ルールブックは愚かなままでなければなりません。パラメータ、しきい値、
スケジュール、動詞がすでに知っているポリシーの中から選択します。構成はテスト済みのものを選択して調整できます
行動;構成内の条件はコードであるため、独自のロジックを導入することはできません。
テスト、レビューア、デバッガを使用せずに実行される、2 番目の実行環境
あなたの本当の影の中で成長します。つまり、ルールブックに if ステートメントが必要になる日、
それは成長を求める設定ではなく、ふるいにかけることです。変化する条件
システムが許可するものはルールであり、ルールは動的であるため、ルールはルールに昇格します。
動詞が呼び出す名前付きのテスト済みポリシーとしてのキネティック レイヤー。そのノブは
ルールブックを守らないと拒否される

セド。名前付きポリシーのライブラリを構築するのが本当の設計です
価格設定だけでも棚がいっぱいになってしまいますが、それはコード内やテストで一度だけ発生します。そして動詞はルールブックを検証します
デポジットのしきい値がゼロに設定されているため、それ自体、法的範囲、構成のスキーマ。
パラメーターのみですが、それでも火傷を負います。
プロモーションには独自の失敗モードがあります。すべてのクライアントの条件付きでプロモーションを行い、書籍が成長します。
40 の旗、これはより良い衛生状態での例外的な追加です。テストも同じふるいです。
顧客 6 の条件は、顧客 7 のルールブック エントリである場合に昇進を獲得します
存在するのを待っていて、それが一般化しない場合、拒否は代償を伴う製品の決定です
その上で。この時代の偉大な生き残りである Salesforce は、両方のことを実現しました。設定を実際のものに変えました
厳しいゲートを持つ言語であるため、Apex はテスト カバレッジが 75% を下回ると本番環境にデプロイされません。
そしてとにかくフローを成長させました: 条件付きの構成、シャドウ実行環境
地球規模での配送。ある道で支払われた代償は、病気がまだ続いている
他には、あなたが持っていないリソースを持っている会社です。
AI は病気と治療の両方をスピードアップします。病気: Anthropic の創設者のハンドブックには、
属の名前、エージェントの技術的負債。種I

[切り捨てられた]

## Original Extract

An essay on why bespoke deployments stop scaling, and the semantic / kinetic / rulebook split that fixes it. - teal-sea/why-your-ai-startup-dies-at-customer-six

GitHub - teal-sea/why-your-ai-startup-dies-at-customer-six: An essay on why bespoke deployments stop scaling, and the semantic / kinetic / rulebook split that fixes it. · GitHub
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
teal-sea
/
why-your-ai-startup-dies-at-customer-six
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit README.md README.md View all files Repository files navigation
Why your AI startup dies at customer six (and the three-layer fix)
Customer one is easy. Their rules are the product: you hardcode their deposit policy, their
hours, their quirks, and the demo looks like magic because it's their own business looking
back at them.
Customer two needs three things different, and you're busy, so their exceptions go where
the code already is. A branch in the workflow, a paragraph in the prompt, a special case in
the tool. By customer six there's a folder with client names on it, and onboarding turns
into development, quoted in weeks.
I'm not writing this from a distance. I run operations for a US limo booking service, an AI
answers its phone on live calls, and I recently went looking for the rule that says a
reservation can't be confirmed without an authorized deposit. Confirmed, in my system,
means a card authorization verified against the payment processor seconds before the state
changes. I found that rule in three places: one version in the booking workflow, one in the
agent's booking tool, one in the dashboard's validation, written months apart, by three
different moods of me. They agreed, mostly. Nothing had blown up yet, and that was the
scary part: nothing was checking. Change the rule and the copy you miss doesn't throw an
error, it just keeps enforcing the old rule on real customers until something expensive
happens. And I only have one tenant: me. Now imagine the folder with six client names on
it.
The numbers explain the temptation. ChartMogul's retention report has AI-native companies
at a median net revenue retention of 48 percent against 82 for ordinary B2B SaaS, and the
segmentation is the interesting part: above 250 dollars a month, AI retention looks like
normal SaaS again, 85 percent. What retains at that tier, on ChartMogul's own reading, is
deeper workflow fit and tighter integration, and hold even that loosely: price is also a
proxy for enterprise buyers, annual contracts and procurement inertia, and the AI-native
cohort mixes consumer products into the comparison. Still, deep
workflow fit is exactly the temptation, because the fastest way to fit a workflow is to
absorb the client's every quirk into the code where they can never leave. Do it without a
platform underneath and Marc Andrusko at a16z has already named where you land: "thousands
of bespoke deployments that are impossible to maintain or upgrade." At that point, in his
words, you aren't Palantir for X. You're "Accenture for X" with a nicer front-end.
The diagnosis is one sentence: the rule doesn't live inside the thing that owns it. Three
drifting copies of a deposit rule and an agent reaching for book without the deposit step
are the same bug in two costumes. The fix starts with naming the layers.
The semantic layer is what things are. Reservation, vehicle, rate, customer: one official
shape each, one set of legal states. The nouns. My three drifting copies were, among other
things, three private definitions of the legal state called confirmed.
The kinetic layer is what can be done. Quote, book, cancel: the verbs of the business, with
the rules inside them and no other door to the tables. The deposit rule lives inside book,
once, for every caller, and it guards the state, not the verb: if a second verb ever
learns to confirm a reservation, a manual override, an import, it inherits the same rule
at the same door.
The rulebook layer is where anyone is allowed to be different: rates, thresholds,
schedules, toggles. A rulebook per client that the verbs read at runtime, and the name
cheats a little: the rules live in the verbs. The rulebook holds each client's answers to
them.
The Greek is borrowed from Palantir's ontology, the semantic and kinetic elements of an
enterprise, its nouns and verbs, and if the first two layers also smell like
domain-driven design, that's because they are: the lineage runs straight through Evans
2003. The rulebook is even less novel. Every SaaS app has a tenant settings table, and
every feature-flag service is a rulebook with a marketing budget. The news is not the
filing system. The news is what keeps a rulebook honest, because config that scripted
callers respect and nothing enforces is a rulebook in name only, and most of them are.
Two things keep it real, and they are the rest of this essay: an adversary that exposes
the fake ones, and a discipline for what gets in.
My read of Palantir's arc, briefly, since a16z already titled this conversation: a decade
of bespoke installs that looked like a services company, until the thing they standardized
turned out to be one level up, the machine for declaring a client's nouns, verbs and
rules. Deployments became authoring, and install time fell to fourteen days by the S-1.
The part that matters for you is that a vertical operator has it easier than Palantir ever
did: they needed a machine general enough for any enterprise's nouns, while a company
serving one vertical can freeze the nouns and verbs themselves, once they've been found,
and hand each client only the rulebook.
Here's the part I had to learn on my own system: you don't declare your primitives. You
find them, and the rulebook is the instrument you find them with.
Two probes force the discovery.
Customer two is the first probe. Every request they make that customer one never made gets
sorted: a parameter you missed (rulebook), a verb you're missing (kinetic), a new kind of
thing entirely (semantic), or a refusal. Some requests land pieces in more than one bucket;
the sorting is still the work. The software product-line people were teaching this as
commonality and variability analysis in 1998. The startup version fits in a sentence: the
rulebook is a sieve, and whatever won't pass through it is product work.
An AI agent is the second probe, and it tests something different. Customer two tests which
parts of the system are shared. The agent tests whether the shared parts actually carry
their own rules. A workflow calls your operations in the order you scripted, so your
invariants only ever face the sequences you thought to write. An agent composes your verbs
at runtime, per conversation, in orders nobody scripted, and it will reach for book without
the deposit step at two in the morning for reasons that made sense to it mid-conversation.
Security people have a word for a machine that feeds your system inputs you never thought
to test: a fuzzer. The agent is a lazier fuzzer. It doesn't search the space, it wanders
it, and it wanders it with your customers in the passenger seat. Every call samples an
ordering of your verbs that nothing ever tested. Under an agent,
a rule that doesn't live inside a verb, or below it in a constraint layer no caller can
route around, is just a habit your scripted callers happened to share. That "below it"
carries real weight, and I'll own it: some invariants belong in the database itself, where
no application code can route around them. The argument is about the number of doors, not
their altitude. Whatever must be
safe on its own, in any order, is a primitive. Everything else is choreography, and an
agent never promised to follow it.
If you've internalized do things that don't scale, this is where you should object:
hardcoding customer one is the canonical advice, and platformizing at customer two is
premature abstraction. Agreed, and notice the sieve asks you to build nothing. It's a
sorting discipline, not a platform. From customer two on you can run it in a ledger: every
client difference gets filed as a parameter, a missing verb, a new object, or a refusal,
while the code stays as hardcoded as you like. What you're buying is that when you finally
do extract the platform, you aren't guessing at the primitives; you're reading two years of
sorted evidence. Premature abstraction is guessing at the sieve's output before running it.
This is just running it, on paper, from day one.
Now the toll, because the rules engines of the 2000s that were sold as "business users
write the logic" died the same death: the rulebook must stay dumb. Parameters, thresholds,
schedules, picks among policies the verbs already know. Config can select and tune tested
behavior; it can never introduce logic of its own, because a conditional in config is code
that runs with no tests, no reviewer and no debugger, a second execution environment
growing in the shadow of your real one. So the day the rulebook wants an if-statement,
that isn't config asking to grow, that's the sieve firing. A conditional that changes
what the system permits is a rule, and rules are kinetic, so it gets promoted into the
kinetic layer as a named, tested policy the verbs invoke, with its knobs left in the
rulebook, or it gets refused. Building that library of named policies is real design
work, pricing alone will fill a shelf, but it happens once, in code, with tests. And the verbs validate the rulebook
itself, legal ranges, a schema for the config, because a deposit threshold set to zero is
parameters-only and still burns you.
Promotion has its own failure mode: promote every client's conditional and book grows
forty flags, which is exception accretion with better hygiene. The test is the same sieve.
Customer six's conditional earns promotion when it's customer seven's rulebook entry
waiting to exist, and when it won't generalize, refusal is a product decision with a price
on it. Salesforce, the era's great survivor, did both things. It turned config into a real
language with a hard gate, Apex won't deploy to production below 75 percent test coverage,
and then it grew Flow anyway: config with conditionals, a shadow execution environment
shipping at planetary scale. The toll paid on one path, the disease still running on the
other, at a company with resources you don't have.
AI speeds up both the disease and the cure. The disease: Anthropic's founder playbook has a
name for the genus, agentic technical debt. The species I

[truncated]
