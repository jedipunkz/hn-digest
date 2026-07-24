---
source: "https://github.com/githubscum/lotor"
hn_url: "https://news.ycombinator.com/item?id=49041486"
title: "Show HN: A gate for AI agents that ships a board of its own worst flaws"
article_title: "GitHub - githubscum/lotor · GitHub"
author: "hnscum"
captured_at: "2026-07-24T21:57:14Z"
capture_tool: "hn-digest"
hn_id: 49041486
score: 1
comments: 1
posted_at: "2026-07-24T20:56:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A gate for AI agents that ships a board of its own worst flaws

- HN: [49041486](https://news.ycombinator.com/item?id=49041486)
- Source: [github.com](https://github.com/githubscum/lotor)
- Score: 1
- Comments: 1
- Posted: 2026-07-24T20:56:50Z

## Translation

タイトル: Show HN: 最悪の欠陥を抱えたボードを出荷する AI エージェントのゲート
記事のタイトル: GitHub - githubscum/lotor · GitHub
説明: GitHub でアカウントを作成して、githubscum/lotor の開発に貢献します。

記事本文:
GitHub - githubscum/lotor · GitHub
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
ギツブスカム
/
ロトール
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 操作

その他のアクション メニュー フォルダーとファイル
50 コミット 50 コミット bin bin 告白 告白 src src テストデータ テストデータ テスト テスト .gitignore .gitignore .zenodo.json .zenodo.json CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md DEMO.md DEMO.md KNOWN-LIMITS.md KNOWN-LIMITS.md ライセンス ライセンス MCP-SETUP.md MCP-SETUP.md README.md README.md manifest.json manifest.json package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント セッション用のローカル ファースト、MCP ネイティブの受信レイヤー。
このツールは、セッション中にエージェントが行ったこと (実行されたアクション、アクセスされたファイル、送信されたメッセージ、発生したコスト、発生した障害) について、署名付きの改ざん防止ログを書き込みます。ログは、検査、検証、アーカイブできる形式でマシン上に保存されます。エージェントの行動が正しかったことを証明しようとするものではなく、それらを忠実に記録し、その後の改ざんを検出可能にするだけです。
現在騒がれている議論は、アクションの結果を予測できない限り、信頼できるエージェント システムを構築することはできないというものです。それは本当で、それは前半だけです。予測は信頼性の最前線です。レコードは裏側です。世界モデルは、アクションが何を行うかを推測します。領収書にはその内容が記載されています。両方必要ですが、今日手にできるのはそのうちの 1 つだけです。
ロートルは後半です。エージェントに世界モデルを与えるわけではありません。これにより、エージェントのセッションが報告した内容について、署名され、順序付けられ、改ざんが明らかな記録が得られるため、「信頼できる」という言葉は主張ではなく、確認できる記録になります。真実ではありません。あなたが保管した記録であり、後から黙って変更することはできません。
そして、その記録はあるべき場所に保管されます。同時にもう 1 つの警告は文化に関するものです。

世界の情報ダイエットは少数の独自エンジンを介して実行されており、それは地域文化の終焉であり、真に公平なシステムなど存在しません。同じ集中化が、エージェントがあなたのために行ったことの記録の 1 層下で静かに行われています。ほとんどのアカウンタビリティ ツールはクラウド オブザーバビリティ プロキシです。エージェントのアクティビティはサーバーを介して流れ、あなた自身の履歴が他の人の帳簿の目録になります。
Lotor は最初にローカルにルーティングされます。領収書はあなたのマシンに書き込まれます。他の場所には書き込まれません。アップロード パス、アカウント、サーバー コンポーネントはありません。送信するために何も構築されていないため、何も残りません。範囲指定された保管統合、つまり編集された記録のスライスを監査人、クライアント、または自社の IT に渡すための認可された方法は保留中です。詳細についてはお問い合わせください。これが全体的な違いであり、展開チェックボックスやエンタープライズ層ではありません。マシンがユーザーのために行ったことの記録は、ユーザーとともに残り、この四半期にどのベンダーをレンタルしていても存続し、保持ポリシーではなくキーに答えます。
これが、モデルが集中化してエージェントが増加した場合の領収書が意味するものです。あなたの本。まずは地元。ベンダーが動いても動かないフロア。
責任の話はどれも同じ壁にぶつかります。記録を作成した当事者だけが記録を保管し、確認することもできません。それは監査ではありません。それは会社が独自に宿題を採点することです。作業を行ったモデルのベンダーが、そのベンダーが行った作業の唯一のログも所有している場合、監査対象の当事者が自らを認証することを信頼していることになります。
チェックにはコストがかかるため、以前は信頼がデフォルトでした。その仮定は公の場で、一度に一つの事件ごとに失効しつつある。それに代わるものは、ベンダーに対するさらなる信頼ではありません。これはベンダーが保持していない記録です。
Lotor は構造によって 2 つの役割を分離します

ション。レシートは、まずローカルのマシンのキーの下に書き込まれます。記録の完全性に実際の利害関係があるのは、記録を保管する当事者であるあなたです。仕事をした人がそれを証明する人になることはできません。そのため、ここではローカルファーストが特徴ではありません。それがポイントです。
Lotor は 1 つの基本的なものであり、署名されたローカル領収書であり、「私のエージェントが実際に何をしたか」という質問が正直に答えられる必要がある場合に適用されます。
ハーネスの許可プロンプトを置き換えます。エージェント自体の確認をオフにしてエージェントを実行し、Lotor にエージェントを停止させます。ハーネス プロンプトは呼び出しごとに一時的に表示され、クリックすると閉じられます。これは、人々がクリックスルーすることを学習するコントロールの形状です。ローター ゲートは 1 つの正確なコマンドにバインドされ、エージェントが生成できないキーで署名され、1 回のみ使用され、承認または拒否が記録されます。読むのをやめたプロンプトと引き換えに、意味を伝えなければならない署名を得るのです。
長期にわたる、委任された仕事。アクションごとにターミナルへの移動が 1 回かかるゲートは、まさにゲートする価値のあるセッション中にオフになるゲートです。委任許可により、1 つの署名が列挙された一連のリクエストをカバーするため、長時間のビルドや夜間の実行は、摩擦が生じるかフロアがまったくないかの選択になるのではなく、制限されたままになります。
文脈を考慮せずに「実際に何が実行されたか」に答えます。何が起こったかを再構築するにはトークンが必要です。ファイルを再読み込みし、履歴を再取得し、すべてをウィンドウ内に保持します。領収書の照会にはほとんどコストがかからず、エージェント自身のアカウントの影響を受けることもありません。実際には、記録のほうが推論よりも速いことが多く、正しいことのほうが多いのです。かつてこのプロジェクトの作業者は、不確実性を認めるために確保されたレポートのセクションで、どのモデルで実行されているかを報告しましたが、それは間違っていました。領収書はありませんでした。
セッションの説明責任

タイ。すべての対話型セッションは、実行されたもの、触れられたもの、送信されたもの、コスト、失敗したものなどの署名済みの受領書で終了します。失敗は埋もれるのではなく表面化する。
夜間および無人の実行。朝目覚めると、決して読むことのない文字起こしの山ではなく、あなたが見ていない間にエージェントが行ったことの記録が目に入ります。
ゲート付きの一か八かのアクション。破壊的または元に戻せないアクション (削除、デプロイ、送信、支出) は、そのアクションとその正確なパラメーターにバインドされた承認に署名するまでフェールクローズされます。拒否と承認の両方が受信されます。
あなたの味方の証拠。クライアント、ベンダー、または取引相手が何が起こったのかについて異議を唱えた場合、あなたはあなた自身の署名された記録を手に入れることになります。彼らの本が、今度はあなたのものになります。
改ざんの証拠。チェーンはログの事後改ざんを検出します。これは、キャプチャ時点でイベントが完全であったこと、または真実であったことを証明するものではありません (既知の制限を参照) が、記録がそれ以来変更されていないことを証明します。
自分自身の記憶を養うこと。領収書は、ベンダーのダッシュボードのテレメトリではなく、長期にわたるエージェントのメモリへの永続的で構造化された入力です。
ベンダーチャーンの生き残り。いつでもモデルやプロバイダーを切り替えることができます。記録はマシン上に残り、読み取り可能なままになります。あなたの歴史は、レンタルしたどのエンジンにも属しません。
範囲を限定した保管統合 (保留中、詳細についてはお問い合わせください)。監査人、クライアント、または自社の IT とレコードを共有するための認可されたパス。スコープが設定され、編集され、署名され、送信先のみに送信されるビュー。 v1 には組み込まれていません。現在、エクスポート パスがまったくないため、デフォルトは選択されたものではなく絶対的なものになっています。共有するものはすべて手動で共有します。
強制的なクレジットではなく、証明可能な優先順位。これは簿記です。彼らの本が、今度はあなたのものになります。
起訴ではなく測定。レシートには、セッションの自己報告内容が記録されます。それは判断しない

サイレントエラーを意図するかキャッチします。
自認捕獲。ログは署名時に始まります。その瞬間以前に何が起こったかは、改ざん証拠チェーンの対象にはなりません。
ここで現在形で説明されているものは何も構築されていません。構築されていないものはすべて、以下のリストまたは KNOWN-LIMITS.md にあり、名前が付けられると同時に保留中としてマークされます。
9 人のゲート付きまたは警告付きのマッチャーは、誰でも頭の中で保持できる姿勢ではありません。 3 つの名前付きプリセットは、9 つのスイッチ マトリックスを 1 つの選択肢に置き換えます。
例外なく、すべてのモードでの自己モッドおよびモード変更ゲート。 Loose とは、世界に対して自由に行動できることを意味しますが、自分を止めるものを自由に書き換えることはできません。Loose のエージェントは依然として、あなたの署名がなければ、ゲート、そのポリシー、フックを編集したり、モードを切り替えたりすることはできません。また、Loose はルールを完全にオフにすることはありません。警告は表示されますが、それでもレシートが追加されます。代替 ( off ) では、一致しないルールはチェーン書き込みがまったくない高速パスを使用するため、最も危険なモードが最も証拠を残さないモードになります。
npm run mode # 現在のモードとそのルールごとの展開を出力します。
npm run mode -- herded # switch (実際の端末で承認パスフレーズが必要)
切り替えには、ゲート アクションに署名するのと同じパスフレーズが必要です。これも同じ理由です。他のすべてで必要なものを変更するには、少なくとも 1 つのことを承認するのと同じくらいのコストがかかるはずです。有効なモードはすべてのセッションの開始レシートにスタンプされるため ( npm run レシート を参照)、その後スイッチが非表示になることはありません。
この機能が出荷される前に作成された既存のpolicy.jsonは、サイレントにアップグレードされません。手動で調整されたルールを維持し、モード Custom でロードします。プリセットに名前を付けるのは、ユーザーが要求した場合にのみ行われます。
正直限界。 Lotor のモードはハーネス自体の許可モードから独立しており、この 2 つは相互に補完しません。ゆるいp

ハーネスが独自のチェックをバイパスするように設定されている場合を除いて、どちらの層でも何も停止することはありません。 Lotor はその組み合わせを検出すると警告し、セッションごとに 1 回姿勢を記録するようになりました。これは、ハーネスが呼び出しごとに独自のモードを報告し、Lotor がそれを破棄していたためです。検出は保護ではありません。床がなくなったことを知らせるものであり、床を元に戻すものではありません。 KNOWN-LIMITS.md の項目 15 を参照してください。
1 回限りの承認は 1 つのコマンドを対象とし、1 回だけ使用されます。これはデプロイでは正しいですが、作業セッションでは間違っています。 40 回のゲート付きアクションは、ターミナルへの 40 回の移動を意味し、高価なゲートのスイッチがオフになります。これは、ゲートが防ぐために存在していた障害です。
許可は、N 個の列挙されたリクエストに対する 1 つの署名であり、1 つのセッションにバインドされ、有効期限と共有アクションの上限が設定されます。
リクエストは、ゲートがすでに書き込んだファイルから送信されます。すべての拒否は、ブロックされた正確なリクエストを段階的に実行します。これにより、印刷された承認コマンドが何も入力せずに実行できるようになります。ゲートで停止するまで数回作業し、それらの拒否の ID を使用してそれらに署名します。
npm run Grant -- --session < id > --requests a1b2c3d4,e5f6a7b8 --max-actions 10 --expires-in-ms 3600000
npm rungrant -- --session < id > -

[切り捨てられた]

## Original Extract

Contribute to githubscum/lotor development by creating an account on GitHub.

GitHub - githubscum/lotor · GitHub
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
githubscum
/
lotor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
50 Commits 50 Commits bin bin confessions confessions src src test-data test-data test test .gitignore .gitignore .zenodo.json .zenodo.json CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CLAUDE.md CLAUDE.md DEMO.md DEMO.md KNOWN-LIMITS.md KNOWN-LIMITS.md LICENSE LICENSE MCP-SETUP.md MCP-SETUP.md README.md README.md manifest.json manifest.json package-lock.json package-lock.json package.json package.json View all files Repository files navigation
A local-first, MCP-native receipt layer for AI agent sessions.
This tool writes a signed, tamper-evident log of what an agent did during a session: actions performed, files touched, messages sent, costs incurred, failures encountered. The log lives on your machine, in a format you can inspect, verify, and archive. It does not attempt to prove the agent's actions were correct, only to record them faithfully and make any subsequent tampering detectable.
The argument getting loud right now is that you cannot have agentic systems that are reliable unless they can predict the consequences of their actions. That is true, and it is only the front half. Prediction is the front of reliability. The record is the back. A world model guesses what an action will do. A receipt states what it did. You need both, and only one of them is something you can hold in your hand today.
Lotor is the back half. It does not give your agent a world model. It gives you a signed, ordered, tamper-evident record of what your agent's session reported doing, so that "reliable" stops being a claim and starts being a record you can check. Not ground truth. A record, kept by you, that cannot be quietly changed afterward.
And it keeps that record where it belongs. The other warning in the same breath is about culture: if the world's information diet runs through a handful of proprietary engines, that is the end of local culture, and no system is ever truly unbiased. The same centralization is happening one layer down, quietly, to the record of what your agents do for you. Most accountability tools are cloud observability proxies. Your agent's activity flows through their servers, and your own history becomes inventory in someone else's books.
Lotor is routed local first. The receipt is written to your machine and nowhere else. There is no upload path, no account, no server component: nothing leaves because nothing is built to send it. Scoped custodial integrations, the sanctioned way to hand a redacted slice of your record to an auditor, a client, or your own IT, are pending. Reach out for details. That is the whole distinction, not a deployment checkbox and not an enterprise tier. The record of what your machines did for you lives with you, survives whichever vendor you were renting this quarter, and answers to your key, not their retention policy.
That is what a receipt means when the models centralize and the agents multiply. Your books. Local first. The floor that does not move when the vendor does.
Every accountability story hits the same wall. The party that produced the record cannot also be the only one who keeps it and checks it. That is not an audit. That is a company grading its own homework. When the vendor whose model did the work also owns the sole log of what it did, you are trusting the audited party to certify itself.
Trust used to be the default, because checking was expensive. That assumption is expiring in public, one incident at a time. What replaces it is not more faith in the vendor. It is a record the vendor does not hold.
Lotor separates the two roles by construction. The receipt is written local first, to your machine, under your key. The party with the real stake in the record's integrity, you, is the party that keeps it. The one who did the work does not get to be the one who certifies it. That is why local first is not a feature here. It is the point.
Lotor is one primitive, a signed local receipt, applied wherever "what did my agent actually do" is a question you need answered honestly.
Replacing your harness's permission prompts. Run the agent with its own confirmations off and let Lotor be the thing that stops it. A harness prompt is per-call, ephemeral, and dismissed with a click, which is the shape of a control people learn to click through. A Lotor gate is bound to one exact command, signed with a key the agent cannot produce, single use, and recorded whether you approve or deny. You trade a prompt you stop reading for a signature you have to mean.
Long-horizon and delegated work. A gate that costs one trip to a terminal per action is a gate that gets turned off during exactly the sessions worth gating. Delegation grants make one signature cover an enumerated set of requests, so a long build or an overnight run stays gated instead of becoming a choice between friction and no floor at all.
Answering "what actually ran" without spending context. Reconstructing what happened costs tokens: re-reading files, re-deriving history, holding it all in the window. Querying a receipt costs almost nothing and is not subject to the agent's account of itself. In practice the record is often faster than the reasoning, and it is right more often. A worker in this project once reported which model it was running on, in the section of its report reserved for admitting uncertainty, and was wrong. The receipt was not.
Session accountability. End every interactive session with a signed receipt: what ran, what it touched, what it sent, what it cost, what failed. Failures are surfaced, not buried.
Overnight and unattended runs. Wake up to a morning-after record of what an agent did while you were not watching, instead of a pile of transcripts you will never read.
Gated high-stakes actions. Make destructive or irreversible actions (delete, deploy, send, spend) fail closed until you sign an approval bound to that exact action and its exact parameters. The denial and the approval are both receipted.
Proof of your side. When a client, vendor, or counterparty disputes what happened, you have your own signed record. Their books, and now yours.
Tamper-evidence. The chain detects after-the-fact alteration of the log. It does not prove the events were complete or truthful at capture time (see Known limits), but it proves the record has not been changed since.
Feeding your own memory. Receipts are durable, structured input to your own long-horizon agent memory, not telemetry for a vendor's dashboard.
Vendor-churn survival. Switch models or providers whenever you want. The record stays on your machine and stays readable. Your history does not belong to whichever engine you were renting.
Scoped custodial integrations (pending, reach out for details). The sanctioned path for sharing a record with an auditor, a client, or your own IT: a scoped, redacted view, signed, sent only where you send it. Not built in v1. Today the default is absolute rather than chosen, because there is no export path at all. Anything you share, you share by hand.
Provable priority, not enforced credit. This is bookkeeping: their books, and now yours.
Measurement, not indictment. A receipt records what the session self-reports; it does not judge intent or catch silent failures.
Self-attested capture. The log begins at signing time. What happened before that moment is not covered by the tamper-evidence chain.
Nothing described in the present tense here is unbuilt. Anything not built is in the list below or in KNOWN-LIMITS.md , and it is marked pending in the same breath as it is named.
Nine gated-or-warned matchers is not a posture anyone can hold in their head. Three named presets replace the nine-switch matrix with a single choice:
self-mod and mode-change gate in every mode, with no exception. Loose means free to act on the world, not free to rewrite what stops you: an agent in Loose still cannot edit the gate, its policy, its hooks, or switch modes without your signature. Loose also never turns a rule fully off — it warns, which still appends a receipt. The alternative ( off ) would make the most dangerous mode the one that leaves the least evidence, since an unmatched rule takes a fast path with no chain write at all.
npm run mode # print the current mode and its rule-by-rule expansion
npm run mode -- herded # switch (requires your approval passphrase at a real terminal)
Switching requires the same passphrase that signs a gated action, for the same reason: changing what everything else requires should cost at least as much as approving one thing. The mode in force is stamped into every session's opening receipt (see npm run receipts ), so a switch is never invisible after the fact.
An existing policy.json from before this feature shipped is not silently upgraded. It keeps its hand-tuned rules and loads with mode custom ; naming a preset only happens when you ask for one.
Honest limit. Lotor's mode is independent of your harness's own permission mode, and the two do not compensate for one another. Loose plus a harness set to bypass its own checks is genuinely nothing stopping anything on either layer. Lotor now warns when it sees that combination and records the posture once per session, because your harness reports its own mode on every call and Lotor was throwing that away. Detection is not protection: it tells you the floor is gone, it does not put one back. See KNOWN-LIMITS.md item 15.
A single-use approval covers one exact command and is spent once. That is right for a deploy and wrong for a working session. Forty gated actions means forty trips to a terminal, and a gate that expensive gets switched off, which is the failure it existed to prevent.
A grant is one signature over N enumerated requests, bound to one session, with an expiry and a shared action ceiling.
The requests come from files the gate already writes. Every denial stages the exact request it blocked, which is what makes the printed approve command runnable with nothing left to fill in. Work until the gate has stopped you a few times, then sign them together using the ids from those denials:
npm run grant -- --session < id > --requests a1b2c3d4,e5f6a7b8 --max-actions 10 --expires-in-ms 3600000
npm run grant -- --session < id > -

[truncated]
