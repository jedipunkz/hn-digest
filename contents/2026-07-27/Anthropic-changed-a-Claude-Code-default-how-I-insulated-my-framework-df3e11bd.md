---
source: "https://mauroepce.dev/blog/trust-boundaries-claude-code"
hn_url: "https://news.ycombinator.com/item?id=49074964"
title: "Anthropic changed a Claude Code default–how I insulated my framework"
article_title: "I Built a Framework on Claude Code. Then Anthropic Changed a Default. Here's How I Defended It. — Blog · mauroepce"
author: "oalders"
captured_at: "2026-07-27T21:03:33Z"
capture_tool: "hn-digest"
hn_id: 49074964
score: 1
comments: 0
posted_at: "2026-07-27T20:12:59Z"
tags:
  - hacker-news
  - translated
---

# Anthropic changed a Claude Code default–how I insulated my framework

- HN: [49074964](https://news.ycombinator.com/item?id=49074964)
- Source: [mauroepce.dev](https://mauroepce.dev/blog/trust-boundaries-claude-code)
- Score: 1
- Comments: 0
- Posted: 2026-07-27T20:12:59Z

## Translation

タイトル: Anthropic はクロード コードのデフォルトを変更しました – 私がフレームワークをどのように隔離したか
記事のタイトル: クロード コードのフレームワークを構築しました。次に、Anthropic がデフォルトを変更しました。これが私がそれを擁護した方法です。 — ブログ · マウロエプセ
説明: Claude Code の 2026 年 7 月の AskUserQuestion の誤機能をケーススタディとして使用した、進化するプラットフォーム上に構築されたエージェント フレームワークの防御パターン。

記事本文:
クロードコードに基づいてフレームワークを構築しました。次に、Anthropic がデフォルトを変更しました。これが私がそれを擁護した方法です。 — ブログ · mauroepce mauroepce プロジェクトについて Signal ブログ お問い合わせ ← ホーム 2026 年 7 月 27 日
クロードコードに基づいてフレームワークを構築しました。次に、Anthropic がデフォルトを変更しました。これが私がそれを擁護した方法です。
Claude Code の 2026 年 7 月の AskUserQuestion の誤機能をケーススタディとして使用した、進化するプラットフォーム上に構築されたエージェント フレームワークの防御パターン。
2026 年 7 月 17 日、Olaf Alders は「Claude Code: Anatomy of a Misfeature」を出版しました。彼は、Claude Code v2.1.198 のサイレント変更を文書化しました。 AskUserQuestion ツール (エージェントが一時停止して人間の入力を待つために使用する主要なメカニズム) は、ユーザーが 60 秒間沈黙すると、ブロックする代わりにモデルの最善の推測による回答を使用して自動継続を開始しました。リリースノートはありません。オプトアウト設定はありません。文書化されていない CLAUDE_AFK_TIMEOUT_MS 環境変数のみをエスケープとして使用します。
Anthropic は 2.1.200 以降でデフォルトを元に戻しました。しかし、このエピソードは、私が自分の仕事で考慮していなかった種類のリスクを明らかにしました。
これは、これに応じて私が構築した防御パターンと、制御していない AI エージェント プラットフォーム上にフレームワークを出荷する人にそれが適用される理由をまとめたものです。
私は個人用のクロード コード ツールキットを管理しています。その中心的な規律は、すべてのコード生成ゲートでの人間参加による確認です。実装前の仕様レビュー、コミット前の明示的な確認、 main にプッシュする前の 2 回目の OK、バグを調査する前の仮説の収集。これらのゲートはすべて 1 つの依存関係を共有しています。つまり、AskUserQuestion を呼び出してエージェントを一時停止し、私の承認またはリダイレクトを待ちます。
その動作が変化すると (たとえば、60 秒後に推測で自動継続するなど)、すべてのゲートが静かに機能低下します。ユーザーは質問されていると思います。実際には、エージェントが彼らに代わって決定しました。
それが脅威のアルダースだ

」という記事が文書化されました。この記事の残りの部分は、それに応じて私が構築したものです。2 つの独立した防御スクリプトと、それらを私の状況を超えて一般化できるパターンに結び付ける推論です。ツールキット自体を見たい場合は、 GitHub にありますが、それはこの記事の要点ではありません。
リスクのクラス (このバグに限らず)
特定の誤機能は元に戻されましたが、それによって明らかになったパターンは永続的です。プラットフォーム上に構築されたフレームワークは、プラットフォームが動作を変更すると、その保証が静かに低下する可能性があります。
従来のソフトウェアには、この問題のバージョンがあります。つまり、依存関係の更新により API コントラクトが破壊されます。ただし、エージェント フレームワークには、より新しい、より卑劣なバージョンがあります。依存関係により、API だけでなく動作がサイレントに変更される可能性があります。コードがコンパイルされます。テストは合格しました。安全扉はまだ存在しているようです。しかし、内部ではセマンティクスが変化しました。
これがあらゆる規模で重要な理由
抽象化を具体化するために、AskUserQuestion のサイレント自動続行が 3 つの段階的にリスクの高いコンテキストにまたがる様子を示します。これらのどれも、Anthropic がまったく同じ誤機能を再導入する必要はありません。同等のサイレント動作の変更 (新しいデフォルトのタイムアウト、名前変更された設定、「無害に見える」セマンティックの調整) は、同じクラスの障害を再現します。
ソロ開発者。コミット確認ステップ (差分を表示し、「このメッセージでコミットしますか?」と尋ねるステップ) は、エージェントを一時停止することになっていました。あなたは電話に出るためにその場を離れました。クロードは沈黙を承認と解釈し、コミットした。コミットは、あなたが編集したであろう次善のメッセージとともに着陸しました。回復可能ですが、レビュー規律は静かに消えていきました。
チーム環境。プッシュツーメイン ゲート (出荷前に「メインへのプッシュを確認しますか?」と尋ねるもの) が最後のチェックポイントでした。 Slack スレッドの途中にいました。クロード・ウェイト

d 60 秒、あなたが「はい」と答えるだろうと判断し、押しました。 CIが始まりました。 「メインは人間による明示的なレビューによって保護される」というチームの慣例は、ルールではなく提案になりました。
規制された、または一か八かの状況。実稼働デプロイのための CI パイプラインの手動承認ゲートは、オンコール エンジニアへの AskUserQuestion プロンプトとして実装されました。ルーチン変更中、オンコールはマルチタスクでした。 60 秒のタイムアウトにより、人間による明示的な承認なしでデプロイが完了します。コンプライアンス フレームワーク (SOC 2、HIPAA、PCI-DSS、ISO 27001) で、運用上の変更に対して文書化された人間による承認が必要な場合、自動続行は単なるバグではなくなります。監査結果です。そして、それを監査人に説明するのは組織の責任であり、人間的なものではありません。
このクラスのリスクを明確に理解すると、遡及的に「Anthropic に注意を払うことを信頼する」のは間違った姿勢のように感じられました。 Anthropic が不注意だからではなく (彼らはそうではなく、誤った機能をすぐに元に戻しました)、ベンダーの信頼が防御パターンではないからです。明示的な構成です。
私は ~/.claude/settings.json に存在する 3 つの構成値に落ち着きました。それぞれが異なる障害ベクトルから防御します。
{
"askUserQuestionTimeout": "決して",
"環境": {
"CLAUDE_AFK_TIMEOUT_MS": "9999999999",
"DISABLE_AUTOUPDATER": "1"
}
}
防御 1: "askUserQuestionTimeout": "決して"
公式設定 Anthropic は 2.1.200+ で追加されました。 (現在のデフォルトに依存するのではなく) 明示的に設定するということは、将来のバージョンが新しいデフォルトを黙って出荷する場合、明示的な値が優先されることを意味します。 「never」がデフォルト以外になっても、設定は影響を受けません。
コスト: ゼロ。価値: フレームワークはバージョンが変動しても予測どおりに動作します。
防御2: CLAUDE_AFK_TIMEOUT_MS: "9999999999"
今日は守備1で冗長。ただし、Claude Code の優先順位では、env vars が設定ファイル sett をオーバーライドします。

イングス。将来のバージョンで askUserQuestionTimeout の名前が変更された場合 (実際に API サーフェスがリファクタリングされる可能性があります)、env var は引き続き境界を保持します。 9 プラス 9 プラス 9 ミリ秒はおよそ 316 年です。事実上無限です。
こちらはベルト＆サスペンダーです。 2 つの独立したメカニズムはどちらも同じことを行うため、単一の変更ベクトルが黙って保証を破ることはできません。
防御3: DISABLE_AUTOUPDATER: "1"
3 つの中で最も重要です。 Claude Code は、ほとんどの最新の開発ツールと同様、自動更新されます。そもそも 2.1.198 の誤機能がユーザーに届いたのは自動更新のためです。影響を受けるバージョンにアップデートするように求められた人は誰もいませんでした。自動アップデーターを無効にしても、更新できなくなるわけではありません。いつ決めるかは私になります。
ワークフローは次のとおりです。新しいバージョンを確認し、変更ログを読み、コミュニティで問題がないか確認し、意図的に更新し、次に検証スクリプトを実行して構成がまだ保持されていることを確認します。更新ごとに 5 分のオーバーヘッドがかかります。これにより、「既知の不正なバージョンにサイレントに更新される」障害モードが完全に防止されます。
2 つのスクリプトが規律を捉えています
ブログ投稿を読んで 3 つの設定値をコピーしてファイルに貼り付けることは、防御パターンではありません。それは減衰する一度限りの行動です。したがって、規律を自動化する必要があります。
スクリプトを 2 つ書きました。そうですね、クロードが入力しました。私がその取り組みを指揮し、功績を認めました。現在はそのようになっているようです。どちらも github.com/mauroepce/claude-workspace/tree/main/bin で公開されています。
verify-claude-config.sh (読み取り専用診断)
クロードコードのバージョン検出。 2.1.198 ～ 199 (既知の危険ゾーン) にいる場合に警告します。
settings.json または settings.local.json の askUserQuestionTimeout
インストールされているフレームワーク コマンド (16 個を予定)
インストールされているフレームワーク テンプレート (9 個を予定)
終了コード 0 はすべてのパス、1 は警告のみ、2 は重大な問題を示します。安全です

o CI で実行、開発ループで安全に実行、偏執的な場合はすべてのエージェント セッションの前に安全に実行できます。
apply-trust-defenses.sh (冪等な設定セッター)
jq を使用して 3 つの防御値を既存の ~/.claude/settings.json にマージします。重要なプロパティ: 他のフィールドはすべて保持されます。既存の 70 の権限、優先モデル、effortLevel はすべて変更されていません。追加または更新されるのは 3 つの防御のみです。
実行するたびに、書き込み前にタイムスタンプ付きのバックアップ ( settings.json.backup-YYYYMMDDTHHMMSSZ ) が作成されます。すべての実行は冪等です。防御がすでに設置されている場合は、その旨を通知し、何も触れずに 0 から終了します。 --yes (CI またはスクリプトによるインストールの場合) を指定して呼び出すと、確認プロンプトがスキップされます。
フレームワーク インストーラーは、オプトインの最終ステップとして apply-trust-defenses.sh を実行することを提案します。ユーザーの確認がなければ何も適用されません。
これは実際にはクロード・コードに関するものではありません。これは、サイレント動作の変更によって保証が低下する可能性がある、重大なベンダーへの依存関係に関するものです。
同じ防御パターンは次のように拡張されます。
決済プロバイダー: API 呼び出しで Stripe API バージョンを明示的に固定します。 「現在の安定版」に依存しないでください。現在の安定は変更される可能性があります。
データベース ドライバー : ORM のクエリ セマンティクスが正確性を重視する場合、package.json 内の ORM バージョンを範囲ではなく正確なバージョンで固定します。
OAuth プロバイダー : 起動時にプロバイダーの検出ドキュメントを読み取り、期待する形状を検証します。 Google または Auth0 がトークン形式を変更した場合は、「ユーザーがログインできない」時点ではなく、デプロイ時にわかります。
LLM 構成: モデルのバージョン ( gpt-4o-latest ではなく gpt-4o-2024-11-20 ) と温度値をバージョン管理されたファイルで明示的にピン留めします。モデルの動作は、再現性にとって重要な点でバージョン間で変動します。
根底にある原則: 沈黙が変化するあらゆるもの

e を使用すると、保証が低下し、バージョン管理するファイルに明示的に値を固定し、その固定がまだ適切であるかどうかの自動検証を追加します。ベンダーの信頼は防御パターンではありません。明示的な構成です。
守備がポイントだ。 2 つのスクリプトと 3 つの設定値だけが必要な場合、それらは単独で動作します。
現在の構成を確認します (読み取り専用、変更なし)。
bash <(curl -fsSL https://raw.githubusercontent.com/mauroepce/claude-workspace/main/bin/verify-claude-config.sh)
3 つの防御を適用します (べき等、最初に settings.json をバックアップします)。
bash <(curl -fsSL https://raw.githubusercontent.com/mauroepce/claude-workspace/main/bin/apply-trust-defenses.sh)
どちらのスクリプトも MIT ライセンスを取得しています。どちらもツールキットの残りの部分には依存しません。ディフェンスパターンだけを取り入れていきましょう。
私がスクリプトを書いた理由の背後にある規律 (仕様優先の作業、明示的な確認ゲート、コードベース規則の自動ロード) に興味がある場合は、完全なフレームワークが独自の README に文書化されています。しかし、それは防御策を採用することとは別の決定です。
『Anatomy of a Misfeature』のオリジナル記事については、Olaf Alders に全面的に感謝します。彼の釈放プロセスの法医学的分析により、パターンが明らかになりました。私の貢献は、防御をインストール可能なスクリプトとしてパッケージ化することです。これが役立つと思われた場合は、彼の記事を読んでください。何が起こったのか、ガバナンスの側面についてさらに深く掘り下げていきます。
GitHub の問題または LinkedIn 経由でのフィードバックを歓迎します。

## Original Extract

A defense pattern for agent frameworks built on evolving platforms, using Claude Code's July 2026 AskUserQuestion misfeature as the case study.

I Built a Framework on Claude Code. Then Anthropic Changed a Default. Here's How I Defended It. — Blog · mauroepce mauroepce About Projects Signal Blog Contact ← Home July 27, 2026
I Built a Framework on Claude Code. Then Anthropic Changed a Default. Here's How I Defended It.
A defense pattern for agent frameworks built on evolving platforms, using Claude Code's July 2026 AskUserQuestion misfeature as the case study.
On July 17, 2026, Olaf Alders published "Claude Code: Anatomy of a Misfeature" . He documented a silent change in Claude Code v2.1.198. The AskUserQuestion tool (the primary mechanism agents use to pause and wait for human input) started auto-continuing after 60 seconds of user silence, using the model's best-guess answer instead of blocking. No release notes. No opt-out setting. Only an undocumented CLAUDE_AFK_TIMEOUT_MS env var as escape.
Anthropic reverted the default in 2.1.200+. But the episode revealed a class of risk I hadn't accounted for in my own work.
This is a write-up of the defensive pattern I built in response, and why it applies to anyone shipping a framework on top of an AI agent platform they don't control.
I maintain a personal Claude Code toolkit whose central discipline is human-in-the-loop confirmation at every code-generation gate. Spec review before implementation, explicit confirmation before commit, a second OK before pushing to main , hypothesis capture before investigating a bug. All of those gates share one dependency: they invoke AskUserQuestion to pause the agent and wait for me to approve or redirect.
If that behavior changes (say, auto-continues with a guess after 60 seconds), every gate silently degrades. The user thinks they're being asked. In reality, the agent decided for them.
That's the threat Alders' article documented. The rest of this piece is what I built in response: two standalone defensive scripts, plus the reasoning that ties them into a pattern generalizable beyond my situation. If you want to see the toolkit itself, it's on GitHub , but it's not the point of the article.
The class of risk (not just this bug)
The specific misfeature was reverted, but the pattern it revealed is permanent: any framework built on top of a platform can have its guarantees silently degraded when the platform changes behavior .
Traditional software has a version of this problem: dependency updates break API contracts. But agent frameworks have a newer, sneakier version. The dependency can silently change behavior , not just API. Your code compiles. Your tests pass. Your safety gates appear to still exist. But under the hood, the semantics moved.
Why this matters at every scale
To make the abstraction concrete, here's what a silent auto-continue on AskUserQuestion breaks across three progressively higher-stakes contexts. None of these require Anthropic to reintroduce the exact same misfeature. Any equivalent silent behavior change (a new default timeout, a renamed setting, a semantic tweak that "seemed harmless") reproduces the same class of failure.
Solo developer. Your commit-confirmation step (the one that shows you the diff and asks "commit with this message?" ) was supposed to pause the agent. You stepped away to take a call. Claude interpreted the silence as approval and committed. The commit landed with a suboptimal message you would have edited. Recoverable, but the review discipline just quietly disappeared.
Team environment. Your push-to-main gate (the one that asks "confirm push to main?" before shipping) was the last checkpoint. You were mid-Slack thread. Claude waited 60 seconds, decided you'd have said yes, pushed. CI kicked off. Your team's convention of "main is protected by explicit human review" became a suggestion, not a rule.
Regulated or high-stakes context. Your CI pipeline's manual approval gate for production deploys was implemented as an AskUserQuestion prompt to the on-call engineer. During a routine change, the on-call was multitasking. The 60-second timeout let the deploy through without explicit human sign-off. If your compliance framework (SOC 2, HIPAA, PCI-DSS, ISO 27001) requires documented human approval for production changes, the auto-continue is no longer just a bug. It's an audit finding. And your organization owns explaining it to the auditor, not Anthropic.
Once I saw this class of risk clearly, retroactively "trusting Anthropic to be careful" felt like the wrong stance. Not because Anthropic is careless (they're not, and they reverted the misfeature quickly), but because trust-in-vendor is not a defense pattern. Explicit configuration is.
I settled on three configuration values that live in ~/.claude/settings.json . Each one defends against a distinct failure vector:
{
"askUserQuestionTimeout": "never",
"env": {
"CLAUDE_AFK_TIMEOUT_MS": "9999999999",
"DISABLE_AUTOUPDATER": "1"
}
}
Defense 1: "askUserQuestionTimeout": "never"
The official setting Anthropic added in 2.1.200+. Setting it explicitly (rather than relying on the current default) means that if a future version silently ships a new default, my explicit value wins. Even if "never" becomes non-default, my config is unaffected.
Cost: zero. Value: the framework works predictably across version drift.
Defense 2: CLAUDE_AFK_TIMEOUT_MS: "9999999999"
Redundant with Defense 1 today. But in Claude Code's precedence order, env vars override config file settings. If a future version renames askUserQuestionTimeout (a real possibility, API surfaces get refactored), the env var still holds the line. Nine-plus-nine-plus-nine milliseconds is roughly 316 years; effectively infinite.
This is belt-and-suspenders. Two independent mechanisms that both do the same thing, so no single change vector can silently break the guarantee.
Defense 3: DISABLE_AUTOUPDATER: "1"
The most important of the three. Claude Code, like most modern dev tools, auto-updates. Auto-updates are why the 2.1.198 misfeature reached users in the first place. Nobody was asked to update to the affected version. Disabling the auto-updater doesn't prevent me from ever updating; it makes me the one who decides when .
The workflow is: I see a new version, I read the changelog, I check the community for issues, I update deliberately, then I run my verify script to confirm the config still holds. It's five minutes of overhead per update. It prevents the "silently updated into a known-bad version" failure mode entirely.
Two scripts capture the discipline
Reading a blog post and copy-pasting three config values into a file is not a defense pattern. It's a one-time behavior that decays. So the discipline needs to be automated.
I wrote two scripts. Well, Claude typed them. I directed the effort and took the credit, which is apparently how this works now. Both public at github.com/mauroepce/claude-workspace/tree/main/bin :
verify-claude-config.sh (read-only diagnostic)
Claude Code version detection. Warns if you're on 2.1.198–199 (the known danger zone).
askUserQuestionTimeout in settings.json or settings.local.json
Framework commands installed (expects 16)
Framework templates installed (expects 9)
Exit code 0 for all pass, 1 for warnings only, 2 for critical issues. Safe to run in CI, safe to run in your dev loop, safe to run before every agent session if you're paranoid.
apply-trust-defenses.sh (idempotent config setter)
Uses jq to merge the three defense values into your existing ~/.claude/settings.json . Critical property: it preserves all other fields . My existing 70 permissions, my preferred model, my effortLevel : all untouched. Only the three defenses get added or updated.
Every run creates a timestamped backup ( settings.json.backup-YYYYMMDDTHHMMSSZ ) before writing. Every run is idempotent: if the defenses are already in place, it says so and exits 0 without touching anything. If you invoke it with --yes (for CI or scripted installs), it skips the confirmation prompt.
The framework installer offers to run apply-trust-defenses.sh as an opt-in final step. Nothing gets applied without user confirmation.
This isn't really about Claude Code. It's about any critical vendor dependency where a silent behavior change can degrade your guarantees .
The same defensive pattern scales to:
Payment providers : pin Stripe API versions explicitly in your API calls. Don't rely on the "current stable"; the current stable can change.
Database drivers : pin ORM versions in package.json with exact versions, not ranges, when the ORM's query semantics matter for correctness.
OAuth providers : read your provider's discovery document at startup and validate the shape you expect. If Google or Auth0 changes their token format, you find out at deploy time, not at "user can't log in" time.
LLM configuration : pin model versions ( gpt-4o-2024-11-20 , not gpt-4o-latest ) and temperature values explicitly in versioned files. Model behavior drifts across versions in ways that matter for reproducibility.
The principle underneath: for anything where a silent change would degrade your guarantees, pin the value explicitly in a file you version-control, and add automated verification that the pin is still in place . Trust-in-vendor is not a defense pattern. Explicit configuration is.
The defenses are the point. If all you want is the two scripts and the three config values, they stand alone:
Verify your current config (read-only, no changes):
bash <(curl -fsSL https://raw.githubusercontent.com/mauroepce/claude-workspace/main/bin/verify-claude-config.sh)
Apply the three defenses (idempotent, backs up your settings.json first):
bash <(curl -fsSL https://raw.githubusercontent.com/mauroepce/claude-workspace/main/bin/apply-trust-defenses.sh)
Both scripts are MIT licensed. Neither depends on the rest of the toolkit. Take just the defense pattern and go.
If the discipline behind why I wrote the scripts (spec-first work, explicit confirmation gates, auto-load of codebase conventions) is interesting to you, the full framework is documented in its own README. But that's a separate decision from adopting the defenses.
Full credit to Olaf Alders for the original Anatomy of a Misfeature article. His forensic analysis of the release process is what revealed the pattern; my contribution is packaging the defense as installable scripts. If you found this useful, read his piece. It goes deeper on the governance side of what happened.
Feedback welcome via GitHub issues or LinkedIn .
