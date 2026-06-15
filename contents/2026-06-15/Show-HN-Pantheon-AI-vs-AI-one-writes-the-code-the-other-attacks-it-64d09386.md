---
source: "https://github.com/lolu1032/pantheon-skills"
hn_url: "https://news.ycombinator.com/item?id=48541711"
title: "Show HN: Pantheon – AI vs AI: one writes the code, the other attacks it"
article_title: "GitHub - lolu1032/pantheon-skills: Two Claude Code skills that run a coding task through a multi-agent harness — plan → N parallel implementations → adversarial verification → judge. pantheon (Claude-only), pantheon-x (GPT-5.5 cross-model verify). MIT. · GitHub"
author: "lolu1032"
captured_at: "2026-06-15T16:44:17Z"
capture_tool: "hn-digest"
hn_id: 48541711
score: 1
comments: 0
posted_at: "2026-06-15T14:20:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Pantheon – AI vs AI: one writes the code, the other attacks it

- HN: [48541711](https://news.ycombinator.com/item?id=48541711)
- Source: [github.com](https://github.com/lolu1032/pantheon-skills)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T14:20:33Z

## Translation

タイトル: 番組 HN: パンテオン – AI 対 AI: 1 人がコードを書き、もう 1 人がコードを攻撃
記事のタイトル: GitHub - lolu1032/pantheon-skills: マルチエージェント ハーネスを通じてコーディング タスクを実行する 2 つのクロード コード スキル — 計画 → N 並列実装 → 敵対的検証 → 判断。 pantheon (クロードのみ)、pantheon-x (GPT-5.5 クロスモデル検証)。マサチューセッツ工科大学· GitHub
説明: マルチエージェント ハーネスを通じてコーディング タスクを実行する 2 つのクロード コード スキル — 計画 → N 並列実装 → 敵対的検証 → 判断。 pantheon (クロードのみ)、pantheon-x (GPT-5.5 クロスモデル検証)。マサチューセッツ工科大学- lolu1032/パンテオンスキル
HN text: あなたが思いついたコードはいつも寛大に見られます。しかしパンテオンは違います。
パンテオンは複数のサブエージェントを回すことで作られており、スコアラーがそれらを1つずつスコアして捨てたとき、生き残った個人がコードのコアになります。
パンテオンには 2 つの形式があります。
GPT 5.5 をスコアラーとして使用する pantheon-x
クロード専用のパンテオン
この2つはもう終わりです。ぜひ試してみて評価してください。ありがとう。

記事本文:
GitHub - lolu1032/pantheon-skills: マルチエージェント ハーネスを通じてコーディング タスクを実行する 2 つのクロード コード スキル — 計画 → N 並列実装 → 敵対的検証 → 判断。 pantheon (クロードのみ)、pantheon-x (GPT-5.5 クロスモデル検証)。マサチューセッツ工科大学· GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します

。
アラートを閉じる
{{ メッセージ }}
ロル1032
/
パンテオンのスキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE pantheon-x pantheon-x pantheon pantheon LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
単一のモデル パスではなく、マルチエージェント ハーネスを通じてハード コーディング タスクを実行する 2 つのクロード コード スキル: 計画 → N 個の並列実装 → 敵対的検証 → 判断。重要なのは、より賢いモデルではありません。2 番目 (および 3 番目) の実装と、結果を解析することが仕事である独立したレビュー担当者が、単一パスで問題なく出荷されるバグをキャッチするということです。
これは、ベストオブ N サンプリング、ツールに統合された自己修正、LLM-as-judge/敵対的検証など、使い古されたテクニックを 1 つの /pantheon コマンドにまとめたもので、毎回手動で再構築する必要はありません。これはモデルの周りの足場であり、モデルへの変更ではありません。モデルが根本的に推論できないタスクを救済するわけではありませんが、答えをテストとして表現できるコーディング作業の正確性を確実に強化します。
ハーネスは決定論的なパイプラインを実行します。
計画 ──▶ 実装（×N並列） ──▶ 検証（敵対的×V） ──▶ 合成
│ │ それぞれが自己修正する │ それぞれをブレイクしようとする │ 審査員が勝者を選ぶ
1 人のプランナー │ 独自のテスト (T1) に対する │ グリーン ビルド │ + 最良のアイデアの移植
N ビルダーのレビュー担当者
計画 — 厳密な仕様、正確性を定義するテスト計画、および (コードの前に) N 個の個別の戦略を導き出します。
実装 — N 個のビルダーが異なる戦略を並行して実装します。それぞれが独自のテストを実行し、失敗した場合は自己修正します（

ol 統合型自己検証、最大 5 回の反復）。
検証 — 独立した敵対的レビュー担当者が、それぞれのグリーン ビルドを破壊しようとします。過半数によって反対されたビルドは削除されます。
総合 — 審査員が勝者を選出し、次点者から移植する価値のある優れたアイデアをリストします。
価値: ビルドは独自のテストに合格しても、依然として間違っている可能性があります。敵対的レイヤーは、グリーン ビルドをゴム印するのではなく、自己作成テストが見逃した欠陥をキャッチします。
スキル
敵対的検証者
要件
パンテオン
クロード自身（独立エージェント）
有料のクロード コード プラン + ワークフロー (下記を参照)
パンテオン-X
Codex プラグイン経由の GPT-5.5 (クロスモデル)
上記 + OpenAI Codex プラグイン ( codex:codex-rescue )
pantheon-x はより強力な設定です。Claude によって書かれた実装は別のモデルによって攻撃され、単一モデルの盲点が縮小されます (同じ間違いが同じモデルの検証ツールをすり抜けます)。 Codex/GPT-5.5 をお持ちでない場合は、 pantheon を使用してください。
どちらのスキルも同じハーネス ( pantheon-class.js ) を共有します。それらは、crossModelVerify フラグのみが異なります。
これらのスキルは Claude Code のワークフロー オーケストレーション エンジンを駆動するため、ストック/無料セットアップだけでは十分ではありません。
有料プランの Claude Code ≥ v2.1.154 — Pro、Max、Team、または Enterprise (Bedrock / Vertex / Foundry も)。無料枠では利用できません。
Pro では、これを一度有効にします: /config → ダイナミック ワークフローをオンにします。
pantheon-x のみ: クロスモデル検証ツールは、標準の Claude Code ではなく、OpenAI の Codex プラグインに同梱されている codex:codex-rescue サブエージェントとして実行されます。ログインした codex CLI だけでは登録されません。プラグインをインストールします。
/プラグイン マーケットプレイス add openai/codex-plugin-cc
/プラグインインストール codex@openai-codex
加えて、ChatGPT サブスクリプション (または OPENAI_API_KEY ) と PATH 上のコーデックス CLI。 codex:codex-rescue がインストールされていない場合は、代わりに pantheon を使用してください。そうでない場合は pantheon-x が使用されます。

se は黙って敵対的パスをスキップし、すべてのビルドをパスします。
スキルとサブエージェント自体はクロード コードの標準機能です。上記以外の追加のセットアップは必要ありません。
Claude Code スキル ディレクトリにクローンを作成します (個人インストール)。
git clone https://github.com/lolu1032/pantheon-skills.git
cp -R pantheon-skills/pantheon ~ /.claude/skills/pantheon
cp -R パンテオン-スキル/パンテオン-x ~ /.claude/スキル/パンテオン-x
または、単一プロジェクトの場合は、 <project>/.claude/skills/ にコピーします。
/pantheon <正確性がテスト可能な難しい実装タスク>
/pantheon-x <同じですが、GPT-5.5 は敵対的検証を行います>
例:
/pantheon 冪等キーの処理を支払いモジュールに追加して、同時リクエストが二重請求できないようにします。テスト: pnpm テスト (vitest)
Claude はパラメーター ( task 、 workdir 、 lang + test コマンド、variants 、 verifiers ) を収集し、バックグラウンド ワークフローとしてハーネスを起動し、バリアントごとのテスト結果 (敵対的なパス ブレークを構築)、および最終的な勝者とその理論的根拠と接木提案をレポートします。
引数
デフォルト
メモ
タスク
—
1 段落の要件 + 合格基準 (テストとして表現可能)
作業ディレクトリ
/tmp/pantheon-<名前>
絶対パス;実際のリポジトリまたはスクラッチ ディレクトリ
ラング
Python/単体テスト
言語 + スタックの正確なテスト コマンド
バリエーション
3
難しい問題の場合は 5 まで上げてください
検証者
2
より厳密にするために 3 に変更します (大多数の反論によりビルドが失われます)
クロスモデル検証
false (パンテオン) / 真 (パンテオン-x)
GPT-5.5/Codex への敵対的検証のルート
コストと範囲
デーモンではありません。各呼び出しは完了まで 1 回実行されて終了します。アイドル状態ではコストはゼロです。
実行には実際のトークンが消費されます。代表的な実行は、約 11 のサブエージェントと数百 K ～ 約 1M のトークンをエンドツーエンドで実行し、実時間で約 6 ～ 10 分です。重い設定 (variants=5、verifiers=3、クロスモデル) では、コストが高くなります。 Pro/M の場合

使用量クォータから抽出されます。従量制の API アクセスでは、実行ごとに数ドル以上の予算がかかります。ここでは最も困難な 10 ～ 20% のタスクのみをルーティングします。残りのタスクにはプレーンな Opus を使用します。
これにより、生のモデルのインテリジェンスではなく、テスト可能な作業の正確性が購入されます。タスクがテストとして表現できない場合、敵対的レイヤーが把握できるものはほとんどなく、オーバーヘッドがかかる価値はありません。
コーディング/エージェントの生産性のみ。安全ゲート (サイバーセキュリティ/生物学的機能制限) を回避するためのツールではありません。
これは単なるプロンプトラッパーではないでしょうか?
モデルの変更はありません。はい、オーケストレーションです。重要な部分は敵対的なステップです。独立したエージェント ( pantheon-x の別のモデル) の仕事は、ビルドを確認するのではなく破壊することです。それは、ビルダー自身のグリーンテストのゴム印で欠陥を検出するものです。値はハーネスの形状であり、秘密のプロンプトではありません。
普通の Opus と比較したベンチマークはありますか?
正式なベンチマークはまだありません。測定されたデルタではなく、メカニズムとして説明を扱います。価値は敵対的ステップにあります。ビルドは独自のテストに合格しても間違っている可能性があり、独立したレビュー担当者が自分で作成したテストのゴム印をキャッチします。直接対決するなら、ぜひ数字を見てみたい。
ランニングコストはいくらですか?
数百 K ～ 100 万トークン、デフォルト設定では約 6 ～ 10 分。バリアント = 5 / 検証者 = 3 / クロスモデルの場合はさらに多くなります。これは、毎日の編集ではなく、最も困難なタスクの 10 ～ 20% を対象としています。 「コストと範囲」を参照してください。
「ワークフローツールが見つかりません」と表示されます/何も起こりません。
無料枠を使用しているか、ワークフローを有効にしていない可能性があります。 「要件 — 有料プランが必要」を参照してください。Pro では、「/config」→「Dynamic workflows」を選択します。
なぜ GPT-5.5 または別のベンダーのモデルに検証をルーティングするのでしょうか?
同じモデルの検証者は盲点を共有しています。これは構築者が犯す間違いですが、同じモデルのレビュー者も見逃す傾向があります。あ

別のモデルを使用することは、その相関関係を打ち破る安価な方法です。これはオプションです。パンテオンはクロード オン クロードを実行しており、引き続き役立ちます。
ソロプロジェクト、現状のまま、ベストエフォート。問題や PR は歓迎しますが、メンテナンスには保証や SLA がありません。すべてに対応できるわけではありません。 MIT ライセンスを取得しているため、さらに進化させたい場合は、フォークが第一級のオプションになります。
マルチエージェント ハーネスを通じてコーディング タスクを実行する 2 つのクロード コード スキル — 計画 → N 並列実装 → 敵対的検証 → 判断。 pantheon (クロードのみ)、pantheon-x (GPT-5.5 クロスモデル検証)。マサチューセッツ工科大学
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Two Claude Code skills that run a coding task through a multi-agent harness — plan → N parallel implementations → adversarial verification → judge. pantheon (Claude-only), pantheon-x (GPT-5.5 cross-model verify). MIT. - lolu1032/pantheon-skills

There's always a generous look at the code you've come up with. But the pantheon is different.
The pantheon is made by turning multiple sub-agents, and when the scorer scores them one by one and throws them away, the surviving individual is now the core of the code.
There are two forms of pantheon.
pantheon-x with GPT 5.5 as Scorer
a Claude-only pantheon
I'm done with these two. Please try them out and evaluate them. Thank you.

GitHub - lolu1032/pantheon-skills: Two Claude Code skills that run a coding task through a multi-agent harness — plan → N parallel implementations → adversarial verification → judge. pantheon (Claude-only), pantheon-x (GPT-5.5 cross-model verify). MIT. · GitHub
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
lolu1032
/
pantheon-skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE pantheon-x pantheon-x pantheon pantheon LICENSE LICENSE README.md README.md View all files Repository files navigation
Two Claude Code skills that run a hard coding task through a multi-agent harness instead of a single model pass: plan → N parallel implementations → adversarial verification → judge . The point isn't a smarter model — it's that a second (and third) implementation, plus an independent reviewer whose job is to break the result, catches bugs a single pass ships green.
It's a packaging of well-worn techniques — best-of-N sampling, tool-integrated self-correction, and LLM-as-judge / adversarial verification — wired into one /pantheon command so you don't reassemble them by hand each time. This is scaffolding around the model, not a change to it: it won't rescue a task the model fundamentally can't reason about, but it reliably tightens correctness on coding work whose answer you can express as tests.
The harness runs a deterministic pipeline:
Plan ──▶ Implement (×N parallel) ──▶ Verify (adversarial ×V) ──▶ Synthesize
│ │ each self-corrects │ try to BREAK each │ judge picks winner
1 planner │ against its own tests (T1) │ green build │ + grafts best ideas
N builders reviewers
Plan — derive a tight spec, a test plan that defines correctness, and N distinct strategies (before any code).
Implement — N builders implement different strategies in parallel; each runs its own tests and self-corrects on failure (tool-integrated self-verification, up to 5 iterations).
Verify — independent adversarial reviewers try to break each green build; a build refuted by a majority is dropped.
Synthesize — a judge picks the winner and lists superior ideas worth grafting from the runners-up.
The value: a build can pass its own tests yet still be wrong. The adversarial layer catches defects the self-written tests miss, instead of rubber-stamping a green build.
Skill
Adversarial verifier
Requirements
pantheon
Claude itself (independent agents)
Paid Claude Code plan + Workflows (see below)
pantheon-x
GPT-5.5 via Codex plugin (cross-model)
Above + OpenAI Codex plugin ( codex:codex-rescue )
pantheon-x is the stronger setting: the implementation written by Claude is attacked by a different model, which shrinks single-model blind spots (the same mistake slipping past a same-model verifier). If you don't have Codex/GPT-5.5, use pantheon .
Both skills share the same harness ( pantheon-class.js ); they differ only in the crossModelVerify flag.
These skills drive Claude Code's Workflow orchestration engine, so a stock/Free setup is not enough:
Claude Code ≥ v2.1.154 on a paid plan — Pro, Max, Team, or Enterprise (also Bedrock / Vertex / Foundry). Not available on the Free tier.
On Pro , enable it once: /config → turn on Dynamic workflows .
pantheon-x only: the cross-model verifier runs as the codex:codex-rescue subagent, which ships in OpenAI's Codex plugin — not stock Claude Code. A logged-in codex CLI alone does not register it. Install the plugin:
/plugin marketplace add openai/codex-plugin-cc
/plugin install codex@openai-codex
plus a ChatGPT subscription (or OPENAI_API_KEY ) and the codex CLI on PATH. If codex:codex-rescue isn't installed, use pantheon instead — pantheon-x would otherwise silently skip the adversarial pass and pass every build.
Skills and subagents themselves are stock Claude Code features; no extra setup beyond the above.
Clone into your Claude Code skills directory (personal install):
git clone https://github.com/lolu1032/pantheon-skills.git
cp -R pantheon-skills/pantheon ~ /.claude/skills/pantheon
cp -R pantheon-skills/pantheon-x ~ /.claude/skills/pantheon-x
Or for a single project, copy into <project>/.claude/skills/ .
/pantheon <a hard implementation task whose correctness is testable>
/pantheon-x <same, but GPT-5.5 does the adversarial verification>
Example:
/pantheon Add idempotency-key handling to the payments module so concurrent requests can't double-charge. Tests: pnpm test (vitest)
Claude collects the parameters ( task , workdir , lang + test command, variants , verifiers ) and launches the harness as a background Workflow, then reports: per-variant test results, which builds the adversarial pass broke, and the final winner with its rationale and grafting suggestions.
arg
default
notes
task
—
one-paragraph requirement + acceptance criteria (expressible as tests)
workdir
/tmp/pantheon-<name>
absolute path; a real repo or a scratch dir
lang
Python/unittest
language + the exact test command for your stack
variants
3
bump to 5 for harder problems
verifiers
2
bump to 3 to be stricter (majority refutation drops a build)
crossModelVerify
false ( pantheon ) / true ( pantheon-x )
route adversarial verify to GPT-5.5/Codex
Cost & scope
Not a daemon. Each invocation runs once to completion and exits — zero cost when idle.
A run spends real tokens. A representative run is ~11 subagents and a few hundred K to ~1M tokens end-to-end, ~6–10 min wall-clock; heavier settings ( variants=5 , verifiers=3 , cross-model) cost more. On Pro/Max it draws from your usage quota; on metered API access, budget a few dollars per run and up. Route only the hardest 10–20% of tasks here — use plain Opus for the rest.
This buys correctness on testable work , not raw model intelligence. If a task isn't expressible as tests, the adversarial layer has little to grip and the overhead isn't worth it.
Coding/agentic productivity only. Not a tool for bypassing safety gates (cybersecurity/biology capability restrictions).
Isn't this just a prompt wrapper?
There's no model change — it's orchestration, yes. The non-trivial part is the adversarial step: an independent agent (a different model in pantheon-x ) whose job is to break a build rather than confirm it. That's what catches defects the builder's own green tests rubber-stamp. The value is the harness shape, not a secret prompt.
Do you have benchmarks vs. plain Opus?
No formal benchmark yet — treat the description as mechanism , not a measured delta. The value is in the adversarial step: a build can pass its own tests and still be wrong, and an independent reviewer catches what the self-written tests rubber-stamp. If you run a head-to-head, I'd genuinely like to see the numbers.
What does a run cost?
A few hundred K to ~1M tokens and ~6–10 min at default settings; more for variants=5 / verifiers=3 / cross-model. It's meant for the hardest 10–20% of tasks, not everyday edits. See Cost & scope .
It says "Workflow tool not found" / nothing happens.
You're likely on the Free tier, or haven't enabled workflows. See Requirements — needs a paid plan and, on Pro, /config → Dynamic workflows .
Why route verification to GPT-5.5 / another vendor's model?
Same-model verifiers share blind spots — a mistake the builder makes, a same-model reviewer tends to miss too. A different model is a cheap way to break that correlation. It's optional: pantheon runs Claude-on-Claude and still helps.
Solo project, as-is, best-effort . Issues and PRs are welcome, but maintenance comes with no guarantees or SLA — I may not get to everything. It's MIT-licensed, so forking is a first-class option if you want to take it further.
Two Claude Code skills that run a coding task through a multi-agent harness — plan → N parallel implementations → adversarial verification → judge. pantheon (Claude-only), pantheon-x (GPT-5.5 cross-model verify). MIT.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
