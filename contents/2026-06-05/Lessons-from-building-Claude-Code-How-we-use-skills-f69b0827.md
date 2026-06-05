---
source: "https://claude.com/blog/lessons-from-building-claude-code-how-we-use-skills"
hn_url: "https://news.ycombinator.com/item?id=48406389"
title: "Lessons from building Claude Code: How we use skills"
article_title: "Lessons from building Claude Code: How we use skills | Claude"
author: "geoffbp"
captured_at: "2026-06-05T00:57:31Z"
capture_tool: "hn-digest"
hn_id: 48406389
score: 1
comments: 0
posted_at: "2026-06-05T00:09:01Z"
tags:
  - hacker-news
  - translated
---

# Lessons from building Claude Code: How we use skills

- HN: [48406389](https://news.ycombinator.com/item?id=48406389)
- Source: [claude.com](https://claude.com/blog/lessons-from-building-claude-code-how-we-use-skills)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T00:09:01Z

## Translation

タイトル: クロードコード構築から得た教訓: スキルの使い方
記事のタイトル: クロード コードの構築から得た教訓: スキルの使い方 |クロード
説明: Anthropic の内部で何百ものスキルを構築し拡張することで学んだこと。

記事本文:
クロード コードの構築から得た教訓: スキルの使い方 |クロード
クロード製品のご紹介 クロード
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者に連絡する 営業担当者に連絡する 営業担当者に問い合わせる
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
クロード コードの構築から得た教訓: スキルの使い方
クロード コードの構築から得た教訓: スキルの使い方
Anthropic の社内で何百ものスキルを構築し拡張することで学んだこと。
共有 リンクをコピー https://claude.com/blog/lessons-from-building-claude-code-how-we-use-skills
スキルは、Claude Code で最もよく使用される拡張ポイントの 1 つです。柔軟性があり、作成も配布も簡単です。
しかし、この柔軟性により、何が最適かを判断することが難しくなります。どのような種類のスキルを身につける価値がありますか?スキルをどのように構成しますか?いつ他の人と共有しますか?
Anthropic ではクロード コードのスキルを広範囲に使用しており、その数百がアクティブに使用されています。これらは、開発を加速するためのスキルの活用について私たちが学んだ教訓です。
スキルは、エージェントが作業をより正確かつ効率的に行うために見つけて使用できる、指示、スクリプト、およびリソースのフォルダーです。このブログ投稿は、スキルの基本を理解していることを前提としています。初めての方は、Skilljar のエージェント スキル入門コースから始めてください。
スキルに関してよく聞かれる誤解は、スキルは「単なるマークダウン ファイル」であるというものです。これらは実際には、エージェントが検出、探索、管理できるスクリプト、アセット、データなどを含めることができるフォルダーです。

膿疱。
Claude Code では、スキルには動的フックの登録など、さまざまな構成オプションもあります。
クロード コードの最も効果的なスキルの一部は、これらの構成オプションとフォルダー構造を効果的に使用していることがわかりました。
前へ 前へ 0 / 5 次へ 次へ Claude Code Desktop を入手
IRM https://claude.ai/install.ps1 | iex コマンドをクリップボードにコピーする またはドキュメントを読む クロード コードを試す クロード コードを試す クロード コードを試す 開発者ドキュメント 開発者ドキュメント 開発者ドキュメント 電子書籍
Anthropic の社内スキルをすべてカタログ化した結果、それらが 9 つのカテゴリに分類されていることがわかりました。最高のスキルは 1 つにきれいに収まります。複数の機能にまたがってあまりにも多くのことを実行しようとするものは、エージェントを混乱させます。これは決定的なリストではありませんが、自分のスキル ライブラリのギャップを特定するのに役立つフレームワークです。
これらは、ライブラリ、CLI、または SDK を正しく使用する方法を説明するスキルです。これらは、Claude Code が時々処理するのに苦労する内部ライブラリまたは共通ライブラリの両方である可能性があります。これらのスキルには、リファレンス コード スニペットのフォルダーと、クロードがスクリプトを作成する際に避けるべき注意点のリストが含まれることがよくありました。
billing-lib — 内部課金ライブラリ: エッジ ケース、フットガンなど。
Internal-platform-cli — 内部 CLI ラッパーのすべてのサブコマンドと、それらを使用する場合の例。
Sandbox-proxy — 開発作業用に組織の Egress ゲートウェイを設定します。アクセス可能なホスト、「接続が拒否されました」エラーをデバッグする方法、ホワイトリスト エントリを追加する方法。
これらは、コードが機能することをテストまたは検証する方法を説明するスキルです。多くの場合、検証のために Playwright、tmux、またはその他の外部ツールと組み合わせられます。
検証スキルは、Claude の内部の出力品質に最も目に見える影響を与えました。エンジニアに 1 週​​間費やしてもらう価値は十分にあります。

あなたの検証スキルは優れています。
クロードに出力のビデオを記録して、テスト内容を正確に確認できるようにしたり、各ステップで状態に対してプログラムによるアサーションを強制したりするなどの手法を検討してください。これらは多くの場合、スキルにさまざまなスクリプトを含めることによって行われます。
Signup-flow-driver — サインアップ→電子メール検証→ヘッドレスブラウザでのオンボーディングまで実行され、各ステップで状態をアサートするためのフックが使用されます。
checkout-verifier — Stripe テスト カードを使用してチェックアウト UI を駆動し、請求書が実際に正しい状態にあることを検証します
tmux-cli-driver — 検証対象に TTY が必要な対話型 CLI テスト用
これらは、データと監視スタックに接続するスキルです。これらのスキルには、資格情報や特定のダッシュボード ID などを使用してデータを取得するためのライブラリに加え、一般的なワークフローやデータを取得する方法に関する手順が含まれる場合があります。
funnel-query — 「サインアップ→アクティベーション→支払いを確認するにはどのイベントに参加すればよいか」と、実際に正規の user_id を持つテーブル
cohort-compare — 2 つのコホートの保持率またはコンバージョンを比較し、統計的に有意なデルタにフラグを立て、セグメント定義にリンクします
grafana — データソース UID、クラスター名、問題 → ダッシュボード ルックアップ テーブル
datadog — フィールド参照 (@request_id と trace_id)、サービス リスト、メトリック プレフィックス規則
4. ビジネスプロセスとチームの自動化
これらは、反復的なワークフローを 1 つのコマンドに自動化するスキルです。これらのスキルは通常、かなり単純な命令ですが、他のスキルや MCP とのより複雑な依存関係がある場合があります。これらのスキルの場合、以前の結果をログ ファイルに保存すると、モデルの一貫性を維持し、ワークフローの以前の実行を反映するのに役立ちます。
standup-post — チケット トラッカー、GitHub アクティビティ、および以前の Slack を集約します → フォーマットされたスタンドアップ、デルト

aのみ
create-<ticket-system>-ticket — スキーマ (有効な列挙値、必須フィールド) と作成後のワークフロー (ping レビューアー、Slack 内のリンク) を適用します。
Weekly-recap — マージされた PR + クローズされたチケット + デプロイ → フォーマットされた要約投稿
5. コードのスキャフォールディングとテンプレート
これらは、コードベース内の特定の関数のフレームワーク定型文を生成するスキルです。これらのスキルを、作成可能なスクリプトと組み合わせることができます。これらは、コードだけではカバーできない自然言語要件がスキャフォールディングにある場合に特に役立ちます。
new-<framework>-workflow — 注釈を使用して新しいサービス/ワークフロー/ハンドラーを足場にします
new-migration — 移行ファイルのテンプレートと一般的な注意事項
create-app — 認証、ロギング、デプロイ設定が事前に設定された新しい内部アプリ
これらは、組織内でコードの品質を強化し、コードのレビューに役立つスキルです。これらには、堅牢性を最大限に高めるための決定論的なスクリプトやツールを含めることができます。これらのスキルをフックの一部として、または GitHub アクション内で自動的に実行したい場合があります。
adversarial-review — 批評する新鮮なサブエージェントを生成し、修正を実装し、結果が要点に達するまで反復します
code-style — コード スタイル、特に Claude がデフォルトでうまく機能しないスタイルを強制します。
testing-practices — テストの書き方と何をテストするかについての指示。
これらは、コードベース内でコードをフェッチ、プッシュ、デプロイするのに役立つスキルです。これらのスキルは、データを収集するために他のスキルを参照する場合があります。
babysit-pr — PR を監視 → 不安定な CI を再試行 → マージ競合を解決 → 自動マージを有効にする
deploy-<service> — ビルド → スモーク テスト → エラー率比較による段階的なトラフィック ロールアウト → 回帰時の自動ロールバック
Cherry-pick-prod — 分離されたワークツリー → チェリーピック → 競合解決 → PR

テンプレート付き
これらは、症状 (Slack スレッド、アラート、エラー署名など) を取得し、マルチツール調査を実行し、構造化されたレポートを作成するスキルです。
<service>-debugging — 最もトラフィックの多いサービスの症状 → ツール → クエリ パターンをマップします。
oncall-runner — アラートを取得 → 通常の容疑者をチェック → 結果をフォーマットする
log-correlator — リクエスト ID を指定すると、それにアクセスした可能性のあるすべてのシステムから一致するログを取得します
これらは日常的なメンテナンスと運用手順を実行するスキルであり、その一部にはガードレールの恩恵を受ける破壊的な行為が含まれます。これらにより、エンジニアは重要な運用においてベスト プラクティスに従うことが容易になります。
<resource>-orphans — 孤立したポッド/ボリュームを検索 → Slack に投稿 → ソーク期間 → ユーザーが確認 → カスケードクリーンアップ
dependency-management — 組織の依存関係承認ワークフロー
コスト調査 - 特定のバケットとクエリ パターンを使用した「ストレージ/下り料金がなぜ高騰したのか」
作るスキルが決まったら、どうやって書くのですか？これらは、スキルを作成するための Claude Code チームのベスト プラクティス、ヒント、コツの一部です。
クロードはすでにコーディング方法を知っており、コードベースを読むことができます。クロードがデフォルトで行うことを言い換えるスキルは、価値を追加することなくコンテキストを追加します。主に知識に関するスキルを公開する場合は、クロードを通常の考え方から押し出す情報に焦点を当ててください。
フロントエンドの設計スキルはその好例です。これは、Anthropic のエンジニアによって、Inter フォントや紫のグラデーションなどの古典的なパターンを避け、Claude のデザインの好みを改善するために顧客と繰り返し検討して構築されました。
どのスキルでも最もシグナルの高いコンテンツは「落とし穴」セクションです。これらのセクションは、一般的な障害点から構築する必要があります。

スキルを使用中にクロードが遭遇する帽子。理想的には、これらの課題を解決するために、時間をかけてスキルを更新してください。
「サブスクリプション テーブルは追加専用です。必要な行は、最新の created_at ではなく、最も高いバージョンを持つ行です。」 「このフィールドは、API ゲートウェイでは @request_id と呼ばれ、課金サービスでは Trace_id と呼ばれます。これらは同じ値です。」 「Stripe Webhook が実際に処理しなかった場合でも、ステージングは​​ 200 を返します。実際の状態についてはpayment_events を確認してください。」
ファイルシステムと段階的な開示を使用する
前に述べたように、スキルは単なるマークダウン ファイルではなくフォルダーです。ファイル システム全体をコンテキスト エンジニアリングと漸進的開示の一形態として考える必要があります。スキル内にどのようなファイルがあるかをクロードに伝えると、適切なタイミングでそれらのファイルが読み取られます。
プログレッシブ開示の最も単純な形式は、クロードが使用できる他のマークダウン ファイルを指定することです。たとえば、詳細な関数シグネチャと使用例を References/api.md に分割できます。
別の例: 最終出力がマークダウン ファイルの場合、コピーして使用するテンプレート ファイルをassets/に含めることができます。
クロードの作業をより効率的に行うのに役立つ、参考資料、スクリプト、サンプルなどのフォルダーを作成できます。
クロードは通常、あなたの指示に従おうとしますが、スキルは再利用できるため、指示が具体的になりすぎることに注意する必要があります。クロードに必要な情報を与えますが、状況に適応できる柔軟性も与えてください。
一部のスキルは、ユーザーからのコンテキストを使用して設定する必要がある場合があります。たとえば、Slack にスタンドアップを投稿するスキルを作成している場合、どの Slack チャネルに投稿するかをクロードに尋ねることができます。
これを行うための良いパターンは、上記の例のように、この設定情報をスキル ディレクトリの config.json ファイルに保存することです。もし

config が設定されていない場合、エージェントはユーザーに情報を要求することができます。
エージェントに構造化された多肢選択の質問を提示させたい場合は、AskUserQuestion ツールを使用するようにクロードに指示できます。
人間ではなくモデルの説明を書く
Claude Code はセッションを開始すると、利用可能なすべてのスキルとその説明のリストを作成します。このリストは、クロードが「このリクエストに対応するスキルがあるかどうか」を判断するためにスキャンしたものです。つまり、説明フィールドは概要ではなく、このスキルをいつトリガーするかの説明です。
一部のスキルには、スキル内にデータを保存することにより、一種の記憶を含めることができます。データは、追加のみのテキスト ログ ファイルや JSON ファイルのような単純なファイルにも、SQLite データベースのような複雑なファイルにも保存できます。
たとえば、standup-post スキルは、投稿が作成されるたびにstandups.log を保存する可能性があります。これは、次回実行するときに、Claude が自身の履歴を読み取り、昨日から何が変更されたかを知ることができることを意味します。
環境変数 ${CLAUDE_PLUGIN_DATA} を使用して、データを保存できる安定したディレクトリを取得できます。永続化データの詳細については、こちらのスキルを参照してください: https://code.claude.com/docs/en/plugins-reference#persistent-data-directory 。
スクリプトを保存してコードを生成する
クロードに提供できる最も強力なツールの 1 つはコードです。クロードにスクリプトとライブラリを与えると、クロードは自分のターンを作曲に費やすことができ、何をするかを決めることができます。

[切り捨てられた]

## Original Extract

What we learned building and scaling hundreds of skills internally at Anthropic.

Lessons from building Claude Code: How we use skills | Claude
Meet Claude Products Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Lessons from building Claude Code: How we use skills
Lessons from building Claude Code: How we use skills
What we learned building and scaling hundreds of skills internally at Anthropic.
Share Copy link https://claude.com/blog/lessons-from-building-claude-code-how-we-use-skills
Skills have become one of the most used extension points in Claude Code. They’re flexible, easy to make, and easy to distribute.
But this flexibility also makes it hard to know what works best. What type of skills are worth making? How do you structure a skill? When do you share them with others?
We've been using skills in Claude Code extensively at Anthropic with hundreds of them in active use. These are the lessons we've learned about using skills to accelerate our development.
Skills are folders of instructions, scripts, and resources that agents can discover and use to do things more accurately and efficiently. This blog post assumes familiarity with skills basics; if you’re new, start with our Introduction to agent skills course on Skilljar .
A common misconception we hear about skills is that they are “just markdown files.” They’re actually folders that can include scripts, assets, data, etc. that the agent can discover, explore and manipulate.
In Claude Code, skills also have a wide variety of configuration options including registering dynamic hooks.
We’ve found that some of the most effective skills in Claude Code use these configuration options and folder structure effectively.
Prev Prev 0 / 5 Next Next Get Claude Code Desktop
irm https://claude.ai/install.ps1 | iex Copy command to clipboard Or read the documentation Try Claude Code Try Claude Code Try Claude Code Developer docs Developer docs Developer docs eBook
After cataloging all of our internal skills at Anthropic, we noticed they cluster into nine categories. The best skills fit cleanly into one; the ones that try to do too much straddle several and confuse the agent. This isn't a definitive list, but it is a useful framework for identifying gaps in your own skills library.
These are skills that explain how to correctly use a library, CLI, or SDKs. They could be both for internal libraries or common libraries that Claude Code sometimes struggles to handle. These skills often included a folder of reference code snippets and a list of gotchas for Claude to avoid when writing a script.
billing-lib — your internal billing library: edge cases, footguns, etc.
internal-platform-cli — every subcommand of your internal CLI wrapper with examples on when to use them.
sandbox-proxy — configuring your org's egress gateway for dev work: which hosts are reachable, how to debug "connection refused" errors, how to add an allowlist entry.
These are skills that describe how to test or verify that your code is working. They are often paired with playwright, tmux, or other external tools for verification.
Verification skills have had the most measurable impact on Claude’s output quality internally. It can be worth having an engineer spend a week just making your verification skills excellent.
Consider techniques like having Claude record a video of its output so you can see exactly what it tested, or enforcing programmatic assertions on state at each step. These are often done by including a variety of scripts in the skill.
signup-flow-driver — runs through signup → email verify → onboarding in a headless browser, with hooks for asserting state at each step
checkout-verifier — drives the checkout UI with Stripe test cards, verifies the invoice actually lands in the right state
tmux-cli-driver — for interactive CLI testing where the thing you're verifying needs a TTY
These are skills that connect to your data and monitoring stacks. These skills might include libraries to fetch your data with credentials, specific dashboard ids, etc., as well as instructions on common workflows or ways to get data.
funnel-query — "which events do I join to see signup → activation → paid" plus the table that actually has the canonical user_id
cohort-compare — compare two cohorts' retention or conversion, flag statistically significant deltas, link to the segment definitions
grafana — datasource UIDs, cluster names, problem → dashboard lookup table
datadog — field reference (@request_id vs trace_id), service list, metric prefix conventions
4. Business process and team automation
These are skills that automate repetitive workflows into one command. These skills are usually fairly simple instructions but might have more complicated dependencies on other skills or MCPs. For these skills, saving previous results in log files can help the model stay consistent and reflect on previous executions of the workflow.
standup-post — aggregates your ticket tracker, GitHub activity, and prior Slack → formatted standup, delta-only
create-<ticket-system>-ticket — enforces schema (valid enum values, required fields) plus post-creation workflow (ping reviewer, link in Slack)
weekly-recap — merged PRs + closed tickets + deploys → formatted recap post
5. Code scaffolding and templates
These are skills that generate framework boilerplates for a specific function in a codebase. You might combine these skills with scripts that can be composed. They are especially useful when your scaffolding has natural language requirements that can’t be purely covered by code.
new-<framework>-workflow — scaffolds a new service/workflow/handler with your annotations
new-migration — your migration file template plus common gotchas
create-app — new internal app with your auth, logging, and deploy config pre-wired
These are skills that enforce code quality inside of your org and help review code. These can include deterministic scripts or tools for maximum robustness. You may want to run these skills automatically as part of hooks or inside of a GitHub Action.
adversarial-review — spawns a fresh-eyes subagent to critique, implements fixes, iterates until findings degrade to nitpicks
code-style — enforces code style, especially styles that Claude does not do well by default.
testing-practices — instructions on how to write tests and what to test.
These are skills that help you fetch, push, and deploy code inside of your codebase. These skills may reference other skills to collect data.
babysit-pr — monitors a PR → retries flaky CI → resolves merge conflicts → enables auto-merge
deploy-<service> — build → smoke test → gradual traffic rollout with error-rate comparison → auto-rollback on regression
cherry-pick-prod — isolated worktree → cherry-pick → conflict resolution → PR with template
These are skills that take a symptom (such as a Slack thread, alert, or error signature), walk through a multi-tool investigation, and produce a structured report.
<service>-debugging — maps symptoms → tools → query patterns for your highest-traffic services
oncall-runner — fetches the alert → checks the usual suspects → formats a finding
log-correlator — given a request ID, pulls matching logs from every system that might have touched it
These are skills that perform routine maintenance and operational procedures, some of which involve destructive actions that benefit from guardrails. These make it easier for engineers to follow best practices in critical operations.
<resource>-orphans — finds orphaned pods/volumes → posts to Slack → soak period → user confirms → cascading cleanup
dependency-management — your org's dependency approval workflow
cost-investigation — "why did our storage/egress bill spike" with the specific buckets and query patterns
Once you've decided on the skill to make, how do you write it? These are some of the Claude Code team’s best practices, tips, and tricks for making skills
Claude already knows how to code and can read your codebase. A skill that restates what Claude would do by default adds context without adding value. If you’re publishing a skill that is primarily about knowledge, focus on information that pushes Claude out of its normal way of thinking.
The frontend design skill is a great example; it was built by an engineer at Anthropic by iterating with customers on improving Claude’s design taste, avoiding classic patterns like the Inter font and purple gradients.
The highest-signal content in any skill is the Gotchas section. These sections should be built up from common failure points that Claude runs into when using your skill. Ideally, you will update your skill over time to capture these gotchas.
"The subscriptions table is append-only. The row you want is the one with the highest version, not the most recent created_at ." "This field is called @request_id in the API gateway and trace_id in the billing service. They're the same value." "Staging returns 200 even when the Stripe webhook didn't actually process. Check payment_events for the real state."
Use the file system and progressive disclosure
Like we said earlier, a skill is a folder, not just a markdown file. You should think of the entire file system as a form of context engineering and progressive disclosure. Tell Claude what files are in your skill, and it will read them at appropriate times.
The simplest form of progressive disclosure is to point to other markdown files for Claude to use. For example, you may split detailed function signatures and usage examples into references/api.md .
Another example: if your end output is a markdown file, you might include a template file for it in assets/ to copy and use.
You can have folders of references, scripts, examples, etc., which help Claude work more effectively.
Claude will generally try to stick to your instructions, and because skills are so reusable you’ll want to be careful of being too specific in your instructions. Give Claude the information it needs, but give it the flexibility to adapt to the situation.
Some skills may need to be set up with context from the user. For example, if you are making a skill that posts your standup to Slack, you may want Claude to ask which Slack channel to post it in.
A good pattern to do this is to store this setup information in a config.json file in the skill directory like the above example. If the config is not set up, the agent can then ask the user for information.
If you want the agent to present structured, multiple choice questions you can instruct Claude to use the AskUserQuestion tool.
Write descriptions for the model, not for humans
When Claude Code starts a session, it builds a listing of every available skill with its description. This listing is what Claude scans to decide "is there a skill for this request?" Which means the description field is not a summary, it's a description of when to trigger this skill.
Some skills can include a form of memory by storing data within them. You could store data in anything as simple as an append only text log file or JSON files, or as complicated as a SQLite database.
For example, a standup-post skill might keep a standups.log with every post it's written, which means the next time you run it, Claude reads its own history and can tell what's changed since yesterday.
You can use the env variable ${CLAUDE_PLUGIN_DATA} to get a stable directory where you can store data, read more persisting data in skills here: https://code.claude.com/docs/en/plugins-reference#persistent-data-directory .
Store scripts and generate code
One of the most powerful tools you can give Claude is code. Giving Claude scripts and libraries lets Claude spend its turns on composition, deciding wha

[truncated]
