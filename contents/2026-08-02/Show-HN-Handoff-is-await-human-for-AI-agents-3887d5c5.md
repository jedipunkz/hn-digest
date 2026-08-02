---
source: "https://github.com/OmegaAgent/handoff"
hn_url: "https://news.ycombinator.com/item?id=49143842"
title: "Show HN: Handoff is await human() for AI agents"
article_title: "GitHub - OmegaAgent/handoff: await human() for AI agents: when an agent hits a wall, a real phone rings and a person takes the wheel in its live browser. · GitHub"
author: "LivingGlitcher"
captured_at: "2026-08-02T12:59:10Z"
capture_tool: "hn-digest"
hn_id: 49143842
score: 1
comments: 1
posted_at: "2026-08-02T12:20:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Handoff is await human() for AI agents

- HN: [49143842](https://news.ycombinator.com/item?id=49143842)
- Source: [github.com](https://github.com/OmegaAgent/handoff)
- Score: 1
- Comments: 1
- Posted: 2026-08-02T12:20:56Z

## Translation

タイトル: HN を表示: AI エージェントのハンドオフは await human()
記事のタイトル: GitHub - OmegaAgent/handoff: AI エージェントの await human(): エージェントが壁にぶつかると、実際の電話が鳴り、ライブ ブラウザで人がハンドルを握ります。 · GitHub
説明: AI エージェントの await human(): エージェントが壁にぶつかると、実際の電話が鳴り、人がライブ ブラウザでハンドルを握ります。 - OmegaAgent/ハンドオフ

記事本文:
GitHub - OmegaAgent/handoff: await human() for AI エージェント: エージェントが壁にぶつかると、実際の電話が鳴り、ライブ ブラウザで人がハンドルを握ります。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。 PL

このページを簡単にリロードしてください。
オメガエージェント
/
引き継ぎ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
89 コミット 89 コミット .github/ workflows .github/ workflows 適合性 適合性 コア コア ドキュメント ドキュメントの例 例 マネージド マネージド スクリプト スクリプト SDK SDK 仕様 仕様 .dockerignore .dockerignore .gitignore .gitignore BACKLOG.md BACKLOG.md COTRIBUTING.md CONTRIBUTING.md DISCLOSURE.md DISCLOSURE.md GOVERNANCE.md GOVERNANCE.md ライセンス ライセンス LICENSE-MIT ライセンス-MIT 通知 通知 README.md README.md SECURITY.md SECURITY.md TRADEMARKS.md TRADEMARKS.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
プログラムが人を必要とする現時点でのオープンプロトコル。
エージェントは継続的に実行されます。ブロックを解除できる人はブロックを解除できず、その人が依然としてブロックを解除できません。
決定に対して責任がある。ハンドオフは、その間に何が起こるか、つまり人間へのリクエストがどのように行われるかを指定します。
決定が述べられ、それがどのようにして誰かに伝わり、その答えがどのように返され、どのような記録が残っているのか
その後。
設計全体がサポートするように構築されている主張は、意図的に狭いものになっています。
人間による 1 つの回答が、入力されたデータとしてエージェントに 1 回だけ配信され、1 つだけ承認されます。
効果。
そのすべての節には負荷がかかります。
人間としての一つの答え。特定の人物が行動し、誰がしたかが記録に残る。クリアランスを主張するのは、
人間であり、副作用から推測されることはありません。その違いがこのプロジェクトが存在する理由です。
以下の従来技術に関するセクションで、それがどこから来たのかを説明します。
まさに一度。回答は使用すると消費されます。保存された「はい」は一瞬も消費できません
再試行、リプレイ、または重複配信による時間。
型付きデータとして。答えは、プロンプトに含まれるフリーテキストではなく、宣言されたスキーマに対する値です。

解釈すること。レンダラーが描けないフィールドは人間が答えられないフィールドなので、未知のフィールド
型はテキスト ボックスに分解されるのではなく、拒否されます。
1 つの効果だけを許可します。答えはそれが示された特定のものにバインドされています
反対。これは、次の呼び出し、次の実行、または同様の要求に一般化されるものではありません。
プログラムの実行は再開されません。スタック スナップショットも継続キャプチャもありません。
ランタイムに到達します。コードは要求し、希望する方法で待機します (ブロッキング呼び出し、Webhook、
世論調査)、回答を受け取ります。その答えをどうするかは完全にあなた次第です。いかなる説明も
それ以外の場合は、後で書く可能性のあるものも含めて、売り過ぎであることを意味します。
また、自己ホスト型の展開では人々にリーチするのがうまくいきません。電子メール、SMS、
または音声は、送信者の評判、調整された DNS レコード、レビューされたアプリ、および通信事業者に依存します
関係。これらはどれもコンテナで出荷されるものではなく、このリポジトリもコンテナで出荷されるふりをしません。
プレリリース。プロトコルが実行されます。公開されておらず、ホストされているサービスも実行されていません
このリファレンス実装。
例の展開/night-hack/ — 4 時間のハッカソン ビルドが従来技術としてここに保存されています。
この実装ではなく、認証なしで handoff.omegas.dev で実行されます。それ
2026 年 7 月 31 日にチェックされました: ホストが回答し、Fly.io によって提供されました。 SECURITY.md には詳細が保持されており、
そのホスト名を範囲外にします。これはライブハッカソンのデモであり、このプロジェクトではなく、またすでに終了したものでもありません
リンク。 (この行は以前、ホストされたサービスが存在しないことをきっぱりと述べています。
これは間違っていた文です。)
レベル 1 は、C-1 ～ C-16、さらに C-6b および C-18 ～ C-26 の 26 件の適合ケースです。 C-17だけです
レベル 2 の場合、このサーバーは意図的にその機能を実装していません。

カバー — 以下を参照してください。の
カウントではどれがどれかを示すものではないため、列挙は裸のカウントのままではなく詳細に記述されます。
サーバーは、レベル 1 に準拠していないケースを除いてすべてのケースを渡します。 docs/review-3.md を読む
これに依存する前に、これは 3 つの独立した敵対的なレビューのうちの最新のものです
( docs/hostile-review.md 、 docs/review-2.md 、 docs/review-3.md )、それぞれに欠陥が見つかりました。
逃す前に。 BACKLOG.md は、残りのギャップを放置するのではなく、明確に記載します。
発見されること。
このツリーには何も公開されていません。 1 つの例外はそれよりも前から存在します: handoff-human 0.1.0 は
NOTICE が依存している MIT 助成金の下で、ハッカソン ビルドからすでに PyPI 上にあります。 0.2.0インチ
sdk/python は公開されておらず、クレート名と npm 名はまだ予約されていません。
この表は、最後のコミット時点での正直な状態です。この内容の他の内容と一致しない場合は、
リポジトリ、このテーブルはチェックされたものです - そして不一致は報告する価値のある欠陥です、
なぜなら、信頼されていて古くなっているステータス文書は、誰も読まないものよりも悪いからです。
spec/ は規範的であり、私たちのものを含むすべての実装の信頼できる情報源です。
仕様は、タグ名前空間 ( spec/v0.1 、
core/v0.3.1 、 sdk-py/v0.2.0 など)。バージョンポリシーについては、GOVERNANCE.mdを参照してください。
また、仕様変更のプロセスについては、コードよりも厳しいものとなります。
spec/ 標準プロトコル。 Apache-2.0。
コア/Rust ワークスペース。 Apache-2.0。 crates.io を対象としています。まだ出版されていません。
crates/handoff-protocol/ タイプ、ステート マシン、ポリシー評価。 I/Oはありません。
crates/handoff-core/ デプロイメントが実装するエンジンとポートの特性
crates/handoff-store-postgres/ リファレンス ストアと独自の移行セット
木箱/ハンドオフアダプター/配送

チャネル、チャネルごとの機能ゲート
crates/handoff-server/ 参照サーバー。バイナリ: handoffd
crates/handoff-conformance/スイート ランナー。任意のベース URL を受け取ります。
適合/宣言的なケース。ガバナンス手段であり、テストヘルパーではありません。
SDK/Python/ハンドオフヒューマン。マサチューセッツ工科大学
sdk/ts/@handoffproto/sdk。マサチューセッツ工科大学。実行時の依存関係はゼロです。
ui/responder/ 計画済み、存在しません。回答者用の独立したページ。 Apache-2.0。
例/ オリジナルのハッカソン ビルドを含む、実際に動作した例。
docs/ドキュメントのソース。
セルフホスティング
自己ホスティングは譲歩ではなく重要であるため、それが何をするのかを正確に把握する価値があります
そしてあなたを捕まえません。
現在の形状: handoffd を Postgres データベースに指定し、それに
リーチできる人々と使用する可能性のあるチャネルを確認し、実行します。それは持っています
ホストされているサービスへの依存はなく、テレフォン ホーム、ライセンス キー、無効化された機能の待機はありません。
スイッチがオンになります。 COTRIBUTING.md は、このリポジトリ内の休止中のゲートをマージ可能なブロックにします
好みの問題というよりは欠陥です。
セルフホスト展開が真に保証するもの: ステート マシンの正確さ。
厳密に 1 つの効果のセマンティクス。プロセスの再起動後も存続する待機。完全なローカル監査証跡。
完全なデータのエクスポート。また、実行時、他の誰かが生きているか、解決能力があるか、興味を持っているかに依存しません。
構造的に保証できないことは、後で発見されるのではなく、ここに記載されています。
独立した証明書。あなたが管理するキーで署名された領収書は、内部のセキュリティに十分です。
コントロールされており、あなたに不利な証拠としては価値がありません。運営者以外の当事者のみが証明できる
オペレーターのシステムが記録したもの。それはサードパーティの定義であり、当社の機能ではありません
差し控える。
配達可能性。上記を参照してください。新しい展開は、すべてのチャネルでレピュテーションがゼロから始まります

エル。
テナント間のシグナル。 「この番号は今日、12 人の異なる送信者によって 400 回ページングされました。」
他のデプロイメントを認識できない 1 つのデプロイメントでは計算できません。
約束として保持します。監査証跡はバックアップとまったく同じくらい耐久性があります。約束には必要なものがある
約束。
その間、仕様を読んで自分で実装することができます。それは支持された結果であり、
適合スイートが存在するため、私たちに尋ねることなく、正しく理解していることを証明できます。
Handoff に関する製品の一部は商用です。境界は公開されたテストによって引かれます。
便宜的にではなく：
単一テナントのハンドオフを定義、実行、独立して検証するために必要なものをすべて開きます
自分のマシン。第三者が運営することによってその価値が得られるもののみを閉じたままにしておきます: 共有
インフラストラクチャ、テナント間の知識、顧客自身のプロセスを超えて存続する約束。
実際には、ほとんどすべてがオープンです。閉じた表面は配送車両に縮小され、
クロステナントビュー、サードパーティ認証、および 1 つの企業独自の請求および組織モデル、
どれもプロトコルではありません。 3 つの条項がメンテナを拘束し、
GOVERNANCE.md : クリップルウェアは存在せず、レシート検証機能は無条件にオープンされ、
エクスポートはオープンです。変更が違反していると思われる場合は、プル リクエスト内のその条項を引用してください。そうでなければなりません
と公の場で答えた。
これがスローガンにならないようにするメカニズムは、適合スイートです。ホスト型サービス
オープン スイートには赤いビルドがあり、「静かにコアをフォークしなかった」ことを意味します。
意図を確認すれば誰でも再実行できます。
サブツリー
ライセンス
spec/ 、 core/ 、 conformance/ 、 ui/ 、 docs/ 、 リポジトリのデフォルト
Apache-2.0 (ライセンス)
SDK/**
MIT ( LICENSE-MIT 、およびディレクトリごとの LICENSE ファイル)
例/ナイトハック/
MIT 、元のリリースから変更されていません
アパッチ-2

.0 仕様と実装に関しては、その §3 が明示的に特許権を付与しているため、
プロジェクトに対する特許訴訟で彼らを終了させる。シンクライアントが取得するため、SDK 上の MIT
ライセンスが Apache-2.0 を厄介にするプロジェクトやその特許を含むあらゆるものにベンダーが参入
露出はゼロです。
Python SDK は、MIT、Copyright (c) 2026 Noureddin Bakir の下で最初に公開されました。その補助金は存続している
ここでは書き換えません。商標条件は TRADEMARKS.md に存在し、決して
ライセンスのテキスト。
まず COTRIBUTING.md を読んでください。短いバージョン:
コミットをサインオフします ( git commit -s )。私たちは開発者原産地証明書を使用します。あります
CLA は存在しませんし、今後も存在しません。なぜなら、CLA が Apache-2.0 を超えるものを購入する唯一のものだからです。
§5 はすでに付与されており、後でこれを閉じるオプションです。
仕様/変更には、その前に不合格でその後に合格する、または 2 つの適合性ケースが必要です。
独立した実装。
適合ケースのないコア動作の変更はマージされません。
最も求められているもの: 配信アダプター、他の言語への SDK ポート、適合ケース、および修正
曖昧な仕様の散文。規範文書のあいまいさは、まさに欠陥です。
セキュリティの問題は SECURITY.md のプライベート チャネルに送られ、決して公開されません。

[切り捨てられた]

## Original Extract

await human() for AI agents: when an agent hits a wall, a real phone rings and a person takes the wheel in its live browser. - OmegaAgent/handoff

GitHub - OmegaAgent/handoff: await human() for AI agents: when an agent hits a wall, a real phone rings and a person takes the wheel in its live browser. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
OmegaAgent
/
handoff
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
89 Commits 89 Commits .github/ workflows .github/ workflows conformance conformance core core docs docs examples examples managed managed scripts scripts sdk sdk spec spec .dockerignore .dockerignore .gitignore .gitignore BACKLOG.md BACKLOG.md CONTRIBUTING.md CONTRIBUTING.md DISCLOSURE.md DISCLOSURE.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE LICENSE-MIT LICENSE-MIT NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md TRADEMARKS.md TRADEMARKS.md View all files Repository files navigation
An open protocol for the moment a program needs a person.
Agents run continuously. The people who can unblock them do not, and those people are still the ones
accountable for the decision. Handoff specifies what happens in between: how a request for a human
decision is stated, how it reaches someone, how their answer comes back, and what record survives
afterwards.
The claim the whole design is built to support is deliberately narrow:
One human answer, delivered to your agent exactly once, as typed data, authorizing exactly one
effect.
Every clause of that is load-bearing:
One human answer. A specific person acted, and the record says who. Clearance is asserted by
a human, never inferred from a side effect. That distinction is the reason this project exists;
the section on prior art below explains where it came from.
Exactly once. An answer is consumed when it is used. A stored yes cannot be spent a second
time by a retry, a replay, or a duplicate delivery.
As typed data. The answer is a value against a declared schema, not free text that a prompt has
to interpret. A field a renderer cannot draw is a field a human cannot answer, so unknown field
types are rejected rather than degraded into a text box.
Authorizing exactly one effect. The answer is bound to the specific thing it was shown
against. It does not generalize to the next call, the next run, or a similar-looking request.
It does not resume your program's execution. There is no stack snapshot, no continuation capture, no
reaching into your runtime. Your code asks, waits however it prefers (a blocking call, a webhook, a
poll), and receives an answer. What it does with that answer is entirely yours. Any description that
implies otherwise is overselling, including any we might write later.
It also does not make a self-hosted deployment good at reaching people. Delivery through email, SMS,
or voice depends on sender reputation, aligned DNS records, a reviewed app, and carrier
relationships. None of that ships in a container, and this repository will not pretend it does.
Pre-release. The protocol runs; it has not been published, and there is no hosted service running
this reference implementation.
A deployment of examples/night-hack/ — the four-hour hackathon build preserved here as prior art,
and not this implementation — is running at handoff.omegas.dev with no authentication . That
was checked on 2026-07-31: the host answers, served by Fly.io. SECURITY.md holds the detail and
puts that hostname out of scope. It is a live hackathon demo, not this project, and not a dead
link. (This line previously said flatly that there is no hosted service, full stop,
which is the sentence that was wrong.)
Level 1 is 26 conformance cases: C-1 through C-16, plus C-6b and C-18 through C-26. C-17 is the only
Level 2 case, and this server deliberately does not implement what it covers — see below. The
enumeration is spelled out rather than left as a bare count, because a count does not say which
cases, and a Server passing every case but one is not Level 1 compliant. Read docs/review-3.md
before relying on any of this — it is the most recent of three independent hostile reviews
( docs/hostile-review.md , docs/review-2.md , docs/review-3.md ), each of which found defects the
one before it had missed. BACKLOG.md states the remaining gaps plainly rather than leaving them
to be discovered.
Nothing in this tree has been published. The one exception predates it: handoff-human 0.1.0 is
already on PyPI from the hackathon build, under the MIT grant NOTICE relies on. The 0.2.0 in
sdk/python is not published, and the crate and npm names are not yet reserved.
This table is the honest state as of the last commit. If it disagrees with something else in this
repository, this table is what was checked — and the disagreement is a defect worth reporting,
because a status document that is trusted and stale is worse than one nobody reads.
spec/ is normative and is the source of truth for every implementation, including ours:
The spec is versioned independently of the implementations, by tag namespace ( spec/v0.1 ,
core/v0.3.1 , sdk-py/v0.2.0 , and so on). See GOVERNANCE.md for the version policy
and for the process a spec change has to go through, which is stricter than the one for code.
spec/ The normative protocol. Apache-2.0.
core/ Rust workspace. Apache-2.0. Intended for crates.io; not yet published.
crates/handoff-protocol/ types, state machine, policy evaluation. No I/O.
crates/handoff-core/ the engine and the port traits a deployment implements
crates/handoff-store-postgres/ reference store and its own migration set
crates/handoff-adapters/ delivery channels, feature-gated per channel
crates/handoff-server/ the reference server. Binary: handoffd
crates/handoff-conformance/ the suite runner. Takes any base URL.
conformance/ Declarative cases. The governance instrument, not a test helper.
sdk/python/ handoff-human. MIT.
sdk/ts/ @handoffproto/sdk. MIT. Zero runtime dependencies.
ui/responder/ PLANNED, not present. A standalone page for the person answering. Apache-2.0.
examples/ Worked examples, including the original hackathon build.
docs/ Documentation source.
Self-hosting
Self-hosting is the point rather than a concession, so it is worth being precise about what it does
and does not get you.
The shape today: point handoffd at a Postgres database, give it a configuration file naming the
people it can reach and the channels it may use, and run it. It has
no dependency on any hosted service, no phone-home, no licence key, and no disabled feature waiting
to be switched on. CONTRIBUTING.md makes a dormant gate in this repository a mergeable-blocking
defect rather than a matter of taste.
What a self-hosted deployment genuinely guarantees: correctness of the state machine;
exactly-one-effect semantics; waits that survive a process restart; a complete local audit trail;
full data export; and no runtime dependency on anyone else being alive, solvent, or interested.
What it structurally cannot guarantee , stated here rather than discovered later:
Independent attestation. A receipt signed with a key you control is adequate for internal
control and worthless as evidence against you. Only a party who is not the operator can attest to
what an operator's system recorded. That is the definition of a third party, not a feature we are
withholding.
Deliverability. See above. A fresh deployment starts at zero reputation on every channel.
Cross-tenant signals. "This number has been paged 400 times today by twelve different senders"
is not computable by one deployment that cannot see the others.
Retention as a promise. Your audit trail is exactly as durable as your backups. A promise needs
a promisor.
Meanwhile you can read the spec and implement it yourself. That is a supported outcome, and the
conformance suite exists so you can prove you got it right without asking us.
Some of the product around Handoff is commercial. The boundary is drawn by a published test rather
than by convenience:
Open everything required to define, run, and independently verify a handoff on a single tenant's
own machine. Keep closed only what derives its value from a third party operating it: shared
infrastructure, cross-tenant knowledge, and promises that outlive the customer's own process.
In practice almost everything is open. The closed surface reduces to a delivery fleet, the
cross-tenant view, third-party attestation, and one company's own billing and organization model,
none of which is protocol. Three clauses bind the maintainer and are written into
GOVERNANCE.md : no crippleware, the receipt verifier is open unconditionally, and
export is open. If you think a change violates one, cite the clause in the pull request; it has to be
answered in public.
The mechanism that keeps this from becoming a slogan is the conformance suite. A hosted service that
cannot pass the open suite has a red build, which turns "we did not quietly fork the core" from an
intention into a check anyone can rerun.
Subtree
Licence
spec/ , core/ , conformance/ , ui/ , docs/ , repository default
Apache-2.0 ( LICENSE )
sdk/**
MIT ( LICENSE-MIT , plus per-directory LICENSE files)
examples/night-hack/
MIT , unchanged from its original release
Apache-2.0 on the spec and the implementation because its §3 grants patent rights expressly and
terminates them on patent litigation against the project. MIT on the SDKs because a thin client gets
vendored into everything, including projects whose licences make Apache-2.0 awkward, and its patent
exposure is nil.
The Python SDK was first published under MIT, Copyright (c) 2026 Noureddin Bakir . That grant stands
and is not rewritten here. Trademark conditions live in TRADEMARKS.md and never in
the licence text.
Read CONTRIBUTING.md first. The short version:
Sign off your commits ( git commit -s ). We use the Developer Certificate of Origin. There is
no CLA and there will not be one , because the only thing a CLA would buy beyond what Apache-2.0
§5 already grants is the option to close this later.
A spec/ change needs a conformance case that fails before it and passes after, or two
independent implementations.
A core behaviour change without a conformance case is not merged.
Most wanted: delivery adapters, SDK ports to other languages, conformance cases, and corrections to
ambiguous spec prose. An ambiguity in a normative document is a real defect.
Security issues go to the private channel in SECURITY.md , never a p

[truncated]
