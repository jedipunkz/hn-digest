---
source: "https://humansmd.dev/"
hn_url: "https://news.ycombinator.com/item?id=50570598"
title: "Humans.md: tell your AI agents how to work with you"
article_title: "HUMANS.md: tell your AI agents how to work with you"
image: "https://humansmd.dev/og.png"
author: "furkanonder"
captured_at: "2026-11-20T05:57:41Z"
capture_tool: "hn-digest"
hn_id: 50570598
score: 1
comments: 0
posted_at: "2026-11-20T05:42:30Z"
tags:
  - hacker-news
  - translated
---

# Humans.md: tell your AI agents how to work with you

- HN: [50570598](https://news.ycombinator.com/item?id=50570598)
- Source: [humansmd.dev](https://humansmd.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-11-20T05:42:30Z

## Translation

タイトル: Humans.md: AI エージェントにどのように連携するかを指示する
記事のタイトル: HUMANS.md: AI エージェントにどのように協力するかを伝える
説明: 開かれている 1 つのファイル ~/.humans/HUMANS.md。このファイルは、Claude Code、Codex、および Cursor に、あなたへの書き込み方法、いつ決定を求めるか、いつ中断するかを指示します。

記事本文:
コンテンツへスキップ HUMANS.md コードなし 仕様と AGENTS.md 実装 FAQ ダウンロード GitHub AGENTS.md はコードベース用です。 HUMANS.mdはあなたのためのものです。
コーディング エージェントがセッション開始時に読み取る 1 つのファイル ~/.humans/HUMANS.md : どのようにあなたに手紙を書くか、いつ決定を求めるか、いつ中断するか。
ドラフト0.3。コーディング エージェントのプラグインとコマンド ラインは、独自のプロファイルを読み込みます。
技術的ではないですか？同じアイデアをコードなしで説明します。
$ claude -p "私の HUMANS.md は中断について何と言っていますか?" 「ヒューマンズにとって、緊急事態については何が問題ですか?」
TL;DR: HUMANS.md には、ブロックされた場合にのみ中断するよう記載されています。他のすべてはタスクが終了するまで待機します。
- 必要なサービスの認証情報がありません → 中断します。
・変数名やテーブルのレイアウトを決める→私が決めて最後に記載します。
プロファイルの最終更新日は 2026 年 10 月 13 日で、古いというフラグは立てられていません。
実際のセッション、プラグイン 0.2.1 を使用したクロード コード 2.1.302、2026 年 10 月 17 日、仕様のサンプル プロファイルを使用。プラグインはセッション開始時にファイルを読み込みます。モデルがどの程度忠実に従うかは、ここでは測定されません。
エージェントがあなたに手紙を書く方法は設定によって決まります。彼らは決して許可を与えません。
クロード プラグイン インストール human-md --marketplace cmdai-hq/humans-md-plugins コピー インストール ガイド
別のエージェント？ Cursor、GitHub Copilot CLI、Google Antigravity、JetBrains Junie、Cline、Warp: それぞれのコマンドをインストールします。
› HUMANS.md には中断について何と記載されていますか?人間にとって何が問題なのか。インターラプションについては何かありますか?
• 作業が妨げられている場合にのみ、割り込みを行ってください。タスクの最後に使用する、妨げにならない質問と、それぞれに対する推奨される回答を保存します。
実際のセッション、Codex CLI 0.163.0 とプラグイン 0.3.0、2026 年 10 月 17 日。プラグインはセッション st でファイルをロードします。

美術;モデルがどの程度忠実に従うかは、ここでは測定されません。
エージェントがあなたに手紙を書く方法は設定によって決まります。彼らは決して許可を与えません。
コーデックス プラグイン マーケットプレイスに cmdai-hq/humans-md-plugins を追加
codex プラグイン add human-md@humans-md コピー インストール ガイド
別のエージェント？ Cursor、GitHub Copilot CLI、Google Antigravity、JetBrains Junie、Cline、Warp: それぞれのコマンドをインストールします。
$ devin -p "私の HUMANS.md には中断について何と書かれていますか?" 「人的問題は何ですか。中断についてはどうすればよいですか?」 -- respect-workspace-trust false
HUMANS.md では、問題をブロックする場合にのみ中断する必要があると指定されています。
**TL;DR**: 私がブロックされている場合にのみ割り込みを求めます。
|設定 |値 |
|----------|----------|
| | の割り込みブロックのみ |
これは、自分で解決できない障害に遭遇しない限り、自律的に作業を続ける必要があることを意味します。実際のセッション、プラグイン 0.3.0 を使用した Devin CLI 3000.11.92、2026 年 10 月 17 日、仕様のサンプル プロファイルをそのまま出力。プラグインはセッション開始時にファイルを読み込みます。モデルがどの程度忠実に従うかは、ここでは測定されません。
エージェントがあなたに手紙を書く方法は設定によって決まります。彼らは決して許可を与えません。
devin plugins install cmdai-hq/ human-md-plugins#plugins/ devin --local コピー インストール ガイド
別のエージェント？ Cursor、GitHub Copilot CLI、Google Antigravity、JetBrains Junie、Cline、Warp: それぞれのコマンドをインストールします。
$ opencode run "私の HUMANS.md には中断について何と書かれていますか?" 「ヒューマンズにとって、緊急事態については何が問題ですか?」
TLDR: HUMANS.md には、問題をブロックする場合にのみ中断すると書かれています。今は遮るものが何もないので、邪魔はしません。
ワークスペースには HUMANS.md ファイルが存在しない (空である) ため、これは環境コンテキストのプロファイル データから取得されます。
あなたの読書の好みに合わせて、私は

TLDR を主導し、149 語の制限内に収まりました。進行を妨げるものにぶつからない限り、黙って進みます。
妨げにならない質問が 1 つあります (タスクの最後に答えてください): あなたの読書設定では、回答の上限が 149 単語に制限されており、表や例が優先されています。今後は返信をコンパクトにしておくべきですか?
実際のセッション、プラグイン 0.3.0 を使用した OpenCode 2.0.32、2026 年 10 月 17 日、仕様のサンプル プロファイルを使用。 JSON 出力からの返信テキスト。まず、フォルダ内で HUMANS.md を検索しましたが、何も見つかりませんでした。値はプラグインから取得されたものです。モデルがそれらにどれだけ忠実に従っているかは、ここでは測定されません。
エージェントがあなたに手紙を書く方法は設定によって決まります。彼らは決して許可を与えません。
opencode プラグイン追加 'github:cmdai-hq/ human-md-plugins#main::path:plugins/ opencode' コピー インストール ガイド
別のエージェント？ Cursor、GitHub Copilot CLI、Google Antigravity、JetBrains Junie、Cline、Warp: それぞれのコマンドをインストールします。
エージェントは並行して実行されます。あなたの注意はそうではありません。
すべての行はオペレーターを経由していました。どの電話がいつつながるかは彼女が決めました。
それは 1912 年のパリでのグーテンベルク交換でした。エージェントも 1 つのデスク、つまりあなたのデスクに電話をかけます。 HUMANS.md は、どの電話が現在鳴るか、どの電話がタスクの終了を待つかを通知します。
1つのファイル。すべてのセッション。一人の人間。
2 つのツール、10 個のセッションが同時に実行されます。彼らが送信するものの半分には、決定または小切手が必要です。半分はそうではないレポートです。どちらも仕事の途中で同じように到着します。夕方までに: AI 疲労、頭の霧。
セッションが開始されると、~/.humans/HUMANS.md が 1 回読み取られます。つまり、どのように読み、どのように決定し、いつ中断するかということです。すべてのセッションで同じコピーが取得されます。
それを読んだセッションは、あなたが質問したとおりに書くことができます。つまり、TL;DR が最初で、質問はタスクの終了のために保存され、現在はブロックされたセッションのみがあなたに届きます。
ファイルの様子を示す図

測定ではなく使用されます
auth-tests API ドキュメント 請求 i18n 移行 ci 検索パフォーマンス メール 中断 4 回 中断 1 回: ブロックされました 中断 2 回: ブロックまたは危険 中断 4 回 タスクの終了 3 つの質問 (それぞれ選択あり) 2 つの質問 (それぞれ選択あり) 何も保存されませんでした。 38 個のテスト…レポート authHelper.ts の名前を変更しますか?あなたが必要ですか？ API に関するクイック更新…レポート どのインデックスを使用するか?あなたが必要ですか？全体的には…ステージングの信用度を報告してください。あなたが必要ですか？私がやったことは次のとおりです... Merge を今すぐ報告しますか、それとも後で報告しますか?あなたが必要ですか？ TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR ! ！ ? ? 10 セッション、1 つのケア ラベル、1 人 上部の HUMANS.md ラベル。以下は、5 つのクロード コード セッションと 5 つのコーデックス セッションで、それぞれにピンクのバッジとしてラベルのコピーが付いています。その下には、TL;DR で始まる短い返信、タスク終了スタックで待機している 3 つの質問、およびその人に届いた 1 つのブロックされたセッションが表示されます。 HUMANS.md max_words: 79 max_words: 149 max_words: 299バッチ_質問: タスク終了割り込み_for: ブロックのみ割り込み_for: ブロックまたはリスク割り込み_for: 何でも クロード コード Codex 認証テスト API ドキュメント請求 i18n 移行 CI 検索パフォーマンス メール TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR ！ ！ ? ? 1 回の中断: ブロック 2 回の中断: ブロックまたは危険 4 回の中断 タスクの終了 3 つの質問 (それぞれ選択あり) 2 つの質問 (それぞれ選択あり) 保存されたものはありません。彼らは決して許可を与えません。
いつ停止して中断されるでしょうか?
HUMANS.md を使用した場合と使用しない場合の同じ作業が報告されました。比較するにはスイッチをオフにしてください。
エージェントがあなたに手紙を書く方法は設定によって決まります。彼らは決して許可を与えません。
このページのために書き下ろしたイラストです
認証テスト · エージェントセッション · 14:40
認証テストを新しいセッション API に移行します。
移行する

d 38 個のテスト ファイル · 212 個のテストに合格
3 つのファイルにはステージング ID プロバイダーが必要です
14:12 に停止し、18 分間待機しました。「authHelper.ts の名前を変更しても大丈夫ですか?」
停止せずに authHelper.ts の名前を変更しました: 可逆的、以下に報告されています
素晴らしいニュースです！私は認証テストの移行に取り組んできましたが、その過程でかなりの数の変更部分があったため、現状について完全な最新情報を提供したいと思いました。
まず、既存の設定を調べました。古いテストでは、Cookie ミドルウェアと密接に結合された従来の SessionStore モックが使用されていました。先に進む前に、管理者テスト用にレガシー モックを保持しておきますか、それとも完全に削除しますか?
とにかく、先に進み、41 個のテスト ファイルのうち 38 個を新しいセッション API に移行しました。途中で、トークンのリフレッシュ テストが不安定であることに気づきました。これは、ハードコーディングされたクロックが原因であると思われます。偽のタイマーに切り替えるべきですか、それとも今はそのままにしておくべきですか?
また、フォルダーの残りの部分と一致するように、authHelper.ts の名前を session-helpers.ts に変更しました。それでいいのか、それとも古い名前の方がいいのか教えてください。
oauth-callback.test.ts 、 sso.test.ts 、および logout-all.test.ts の 3 つのファイルは依然として失敗します。これらはステージング ID プロバイダーに依存していますが、ここからはアクセスできません。ステージング資格情報が変更されたかどうかを確認していただけますか?
全体として、移行は良い方向に進んでいると思いますが、フィクスチャを統合したり、プロバイダー向けの契約テストを追加したりするなど、できることは間違いなくあります。どのように進めたいか、PR を開いてマージしてステージングにデプロイするかどうかをお知らせください。
TL;DR: 41 個の認証テスト ファイルのうち 38 個が新しいセッション API で実行されます。 3 つはステージング ID プロバイダーでブロックされています。
トークン更新テストでは偽のタイマーが使用されるようになりました。ハードコードされた時計のせいで不安定になってしまった。
レナム

フォルダーと一致するように authHelper.ts を session-helpers.ts に編集しました。
質問が2つあります。 「はい」は私の選択を承認します:
従来の SessionStore モックを削除しますか?私の選択: はい、今は何も使用していません。
3 つのファイルを完了できるように、ステージング IdP 認証情報をローテーションしますか?私の選択: はい。
245 単語 · テキスト全体に広がる 4 つの質問 · 1 つの中断 · 提供されるマージとデプロイ 91 単語、キャップ 149 · 最後に 2 つの質問、それぞれ選択付き · 中断 0
測定ではなく、スクリプト化されたセッションの図。プラグインはプロファイルをロードします。モデルがどの程度それに忠実に従っているかは確立されていません。
~/.humans/HUMANS.md 実際の例
---
# HUMANS.md 0.3: 私と協力する方法。フォームのみで許可は必要ありません。リポジトリのルールとハーネスが優先されます。
# 保存された状態は非アクティブです。決して適用したり、内容から推測したりしないでください。
# out_of_scope_steps: タスクの外側の次のステップ (マージ、デプロイ) を最後の行にリストするかどうか。
人間:「0.3」
ID:トリスタン
名前：トリスタン
更新日: "2026-10-13"
可視性: プライベート
読み方：
リード付き: tldr
最大ワード数: 149
ビジュアル: [表、例]
決定する:
可逆: 決定して報告する
バッチ質問: タスクの終了
デフォルトの答え: 常に
question_style: はい_確認_推奨
中断:
割り込み対象: ブロックのみ
注: 「私はよく携帯電話で本を読みます。」
---
点灯: この返信が続くフィールド。オンにすると、どのフィールドが応答を変更するかを確認できます。
セッションと両方の返信はこのページに書き込まれ、記録されません。プロファイルは、バイトごとの仕様の実際の例です。
この返信でプロフィールが変わったこと プロフィールなしだと同じ作品がこうなります
取り消し可能な選択が行われ、その後報告される
最後に質問があり、それぞれに選択肢があります
楽しいことが先、結果は後から
テキスト内に散りばめられた 4 つの質問
タスクの途中で 1 回停止、待ち時間は 18 分
マージとデプロイは要求されずに提供されます
まずはTL;DR、

150語以内。表か例を見せてください。元に戻せるなら、決めて言ってください。タスクの最後に質問をまとめて、常に選択したものを選択して、「はい」で承認されるようにします。あなたがブロックされている場合にのみ、私に割り込んでください。私はよくスマホで本を読みます。
まず、わかりやすい言葉で言います。
どの言語でも。保持したい個人情報のみを含めてください。
エージェントが各文を編集します。
現在、コンパイラーはすでに使用されているアシスタントです。 Claude Code では、プラグインの /humans-md:compile がそれを行います。他のエージェントでは、human-md-compile スキル。
可逆: 決定して報告する OK
割り込み対象: ブロックのみOK
省略または推測: 最初に表示されます。
そしてそれを読む人。可視性: プライベートを明示的に書き込みます: これが無いプロファイルは org として読み取られます。
リーダーがチェックしました。エージェントにデータとして渡される認識されたフィールド。
source.md、コメントが渡されていません
Claude Code、Codex、または Cursor では、humans-md プラグインによってロードされます。
エージェントはセッション開始時にそれを読み取ります。
これを共有するために、チーム リポジトリはコピーを human/<id>.md として保持し、同僚のエージェントがツールでそれを読み取ったときに使用できるようにします。公開されている統合はまだありません。リポジトリ ルートにある HUMANS.md はロードされません。
HUMANS.md には、どのような相談を希望するかが記載されています。許可は与えません。リポジトリの AGENTS.md に「デプロイする前に質問する」と記載されている場合、エージェントは質問します。

[切り捨てられた]

## Original Extract

One open file, ~/.humans/HUMANS.md, that tells Claude Code, Codex and Cursor how to write to you, when to ask you to decide and when to interrupt you.

Skip to content HUMANS.md Without code Spec vs AGENTS.md Implementations FAQ Downloads GitHub AGENTS.md is for your codebase. HUMANS.md is for you.
One file, ~/.humans/HUMANS.md , that your coding agents read at session start: how to write to you, when to ask you to decide, when to interrupt you.
Draft 0.3. Plugins for your coding agents and a command line load your own profile.
Not technical? The same idea, explained without code .
$ claude -p "What does my HUMANS.md say about interruptions?" " W h a t d o e s m y H U M A N S . m d s a y a b o u t i n t e r r u p t i o n s ? "
TL;DR: Your HUMANS.md says to interrupt you only when blocked. Everything else waits until the end of the task.
- Missing credentials for a required service → I interrupt you.
- Picking a variable name or a table layout → I decide and mention it at the end.
The profile was last updated on 2026-10-13 and isn’t flagged as stale.
Real session, Claude Code 2.1.302 with plugin 0.2.1, 17 Oct 2026, with the example profile from the spec. The plugin loads your file at session start; how closely a model follows it is not measured here.
Preferences shape how agents write to you. They never grant a permission.
claude plugin install humans-md --marketplace cmdai-hq/humans-md-plugins Copy Install guide
Another agent? Cursor, GitHub Copilot CLI, Google Antigravity, JetBrains Junie, Cline and Warp: install commands for each .
› What does my HUMANS.md say about interruptions? W h a t d o e s m y H U M A N S . m d s a y a b o u t i n t e r r u p t i o n s ?
• Only interrupt you when work is blocked. Save non-blocking questions for the end of the task, with a recommended answer for each.
Real session, Codex CLI 0.163.0 with plugin 0.3.0, 17 Oct 2026. The plugin loads your file at session start; how closely a model follows it is not measured here.
Preferences shape how agents write to you. They never grant a permission.
codex plugin marketplace add cmdai-hq/humans-md-plugins
codex plugin add humans-md@humans-md Copy Install guide
Another agent? Cursor, GitHub Copilot CLI, Google Antigravity, JetBrains Junie, Cline and Warp: install commands for each .
$ devin -p "What does my HUMANS.md say about interruptions?" " W h a t d o e s m y H U M A N S . m d s a y a b o u t i n t e r r u p t i o n s ? " --respect-workspace-trust false
Your HUMANS.md specifies that I should interrupt only for blocking issues.
**TL;DR**: You only want interruptions when I'm blocked.
| Setting | Value |
|---------|-------|
| Interrupt for | Blocking only |
This means I should continue working autonomously unless I encounter a blocker I cannot resolve myself. Real session, Devin CLI 3000.11.92 with plugin 0.3.0, 17 Oct 2026, with the example profile from the spec, printed as is. The plugin loads your file at session start; how closely a model follows it is not measured here.
Preferences shape how agents write to you. They never grant a permission.
devin plugins install cmdai-hq/ humans-md-plugins#plugins/ devin --local Copy Install guide
Another agent? Cursor, GitHub Copilot CLI, Google Antigravity, JetBrains Junie, Cline and Warp: install commands for each .
$ opencode run "What does my HUMANS.md say about interruptions?" " W h a t d o e s m y H U M A N S . m d s a y a b o u t i n t e r r u p t i o n s ? "
TLDR: Your HUMANS.md says interrupt only for blocking issues. Since nothing is blocking me right now, I won’t interrupt you.
No HUMANS.md file exists in your workspace (it’s empty), so this comes from the profile data in your environment context:
Per your reading preferences, I’ve led with the TLDR and stayed under your 149-word limit. I’ll proceed silently unless I hit something that blocks progress.
One non-blocking question (answer at end of task): Your reading settings cap responses at 149 words and prefer tables/examples — shall I keep replies that compact going forward?
Real session, OpenCode 2.0.32 with plugin 0.3.0, 17 Oct 2026, with the example profile from the spec; the reply text, from its JSON output. It first searched the folder for a HUMANS.md and found none: the values came from the plugin. How closely a model follows them is not measured here.
Preferences shape how agents write to you. They never grant a permission.
opencode plugin add 'github:cmdai-hq/ humans-md-plugins#main::path:plugins/ opencode' Copy Install guide
Another agent? Cursor, GitHub Copilot CLI, Google Antigravity, JetBrains Junie, Cline and Warp: install commands for each .
Your agents run in parallel. Your attention doesn’t.
Every line went through an operator. She decided which call got through, and when.
That was the Gutenberg exchange, Paris, 1912. Your agents ring one desk too: yours. HUMANS.md tells them which calls may ring now, and which wait for the end of the task.
One file. Every session. One human.
Two tools, ten sessions running at once. Half of what they send needs you, a decision or a check; half is reports that do not. Both arrive the same way, in the middle of your work. By the evening: AI fatigue, brain fog.
When a session starts, it reads ~/.humans/HUMANS.md once: how you read, how you decide, when to interrupt you. Every session gets the same copy.
Sessions that read it can write the way you asked: a TL;DR first, questions saved for the end of the task, and only a blocked session reaching you now.
An illustration of how the file is used, not a measurement
auth-tests api docs billing i18n migration ci search perf emails 4 interruptions 1 interruption: blocked 2 interruptions: blocked or risky 4 interruptions End of task 3 questions, each with a pick 2 questions, each with a pick Nothing saved Great news! 38 tests… report Rename authHelper.ts? needs you ? Quick update on API… report Which index to use? needs you ? Overall I think… report Staging creds, please? needs you ? Here is what I did… report Merge now or later? needs you ? TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR ! ! ? ? Ten sessions, one care label, one person The HUMANS.md label at the top. Below, five Claude Code sessions and five Codex sessions, each with a copy of the label as a pink badge. Below them, short replies that start with a TL;DR, three questions waiting in an end-of-task stack, and one blocked session reaching the person. HUMANS.md max_words: 79 max_words: 149 max_words: 299 batch_questions: end_of_task interrupt_for: blocking_only interrupt_for: blocking_or_risky interrupt_for: anything Claude Code Codex auth-tests api docs billing i18n migration ci search perf emails TL;DR TL;DR TL;DR TL;DR TL;DR TL;DR ! ! ? ? 1 interruption: blocked 2 interruptions: blocked or risky 4 interruptions End of task 3 questions, each with a pick 2 questions, each with a pick Nothing saved Preferences shape how agents write to you. They never grant a permission.
When may it stop and interrupt you?
The same work, reported with and without HUMANS.md. Switch it off to compare.
Preferences shape how agents write to you. They never grant a permission.
An illustration written for this page
auth-tests · agent session · 14:40
Migrate the auth tests to the new session API.
Migrated 38 test files · 212 tests pass
3 files need the staging identity provider
Stopped at 14:12 and waited 18 min: “Is it okay to rename authHelper.ts?”
Renamed authHelper.ts without stopping: reversible, reported below
Great news! I’ve been working through the auth test migration and wanted to give you a full update on where things stand, since there were quite a few moving parts along the way.
First, I looked at the existing setup. The old tests used the legacy SessionStore mock, which was tightly coupled to the cookie middleware. Before I go further: do you want me to keep the legacy mock around for the admin tests, or remove it entirely?
Anyway, I went ahead and migrated 38 of the 41 test files to the new session API. Along the way I noticed that the token refresh tests were flaky, which seems to be caused by a hard-coded clock. Should I switch them to fake timers, or would you rather I leave them as they are for now?
I also renamed authHelper.ts to session-helpers.ts to match the rest of the folder. Let me know if that’s okay or if you’d prefer the old name.
Three files still fail: oauth-callback.test.ts , sso.test.ts and logout-all.test.ts . They depend on the staging identity provider, which I can’t reach from here. Could you check whether the staging credentials have changed?
Overall I think the migration is in a good place, although there is definitely more we could do, such as consolidating the fixtures and maybe adding contract tests for the provider. Let me know how you’d like to proceed, and whether you want me to open a PR, merge it and deploy to staging!
TL;DR: 38 of 41 auth test files run on the new session API. 3 are blocked on the staging identity provider.
Token refresh tests now use fake timers; the hard-coded clock made them flaky.
Renamed authHelper.ts to session-helpers.ts to match the folder.
2 questions. “yes” approves my pick:
Delete the legacy SessionStore mock? My pick: yes, nothing uses it now.
Rotate the staging IdP credentials so I can finish the 3 files? My pick: yes.
245 words · 4 questions spread through the text · 1 interruption · merge and deploy offered 91 words, cap 149 · 2 questions, at the end, each with a pick · 0 interruptions
Illustration of a scripted session, not a measurement. The plugin loads the profile; how closely a model follows it is not established.
~/.humans/HUMANS.md the real example
---
# HUMANS.md 0.3: how to work with me. Form only, never a permission: repo rules and your harness win.
# Stored states are inactive: never apply one, and never infer one from content.
# out_of_scope_steps: whether to LIST next steps outside the task (merge, deploy) in your closing lines.
humans: "0.3"
id: tristan
name: Tristan
updated: "2026-10-13"
visibility: private
reading:
lead_with: tldr
max_words: 149
visual: [table, example]
deciding:
reversible: decide_and_report
batch_questions: end_of_task
default_answer: always
question_style: yes_confirms_recommendation
interrupting:
interrupt_for: blocking_only
notes: "I often read on my phone."
---
Lit: the fields this reply follows. Switch it on to see which fields change the reply.
The session and both replies are written for this page, not recorded. The profile is the real example from the spec, byte for byte.
What the profile changed in this reply Without the profile, the same work reads like this
Reversible choices made, then reported
Questions at the end, each with a pick
Pleasantries first, result later
Four questions scattered in the text
One stop mid-task, 18 minutes waiting
Merge and deploy offered, unasked
TL;DR first, under 150 words. Show me a table or an example. If it can be undone, decide and tell me. Batch your questions at the end of the task, always with your pick, so a yes approves it. Only interrupt me when you are blocked. I often read on my phone.
First you say it, in plain words.
Any language. Include only the personal information you want kept.
Your agent compiles each sentence.
Today the compiler is the assistant you already use. In Claude Code, the plugin’s /humans-md:compile does it; in other agents, the humans-md-compile skill.
reversible: decide_and_report ok
interrupt_for: blocking_only ok
Left out or guessed: shown to you first.
And who reads it. Write visibility: private explicitly: a profile without it is read as org .
Reader checked. Recognized fields passed to the agent as data.
not passed source.md, comments
In Claude Code, Codex or Cursor, the humans-md plugin loads it.
Agents read it at session start.
To share it, a team repository holds a copy as humans/<id>.md , for colleagues’ agents once their tools read it; no published integration does yet. A HUMANS.md at a repository root is never loaded.
HUMANS.md says how you prefer to be consulted. It grants no permission. If a repository’s AGENTS.md says “ask before you deploy”, the agent asks

[truncated]
