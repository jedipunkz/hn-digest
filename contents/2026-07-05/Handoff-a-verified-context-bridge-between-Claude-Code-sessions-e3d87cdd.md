---
source: "https://github.com/ostikwhy-blip/claude-code-handoff-skill"
hn_url: "https://news.ycombinator.com/item?id=48795956"
title: "Handoff – a verified context bridge between Claude Code sessions"
article_title: "GitHub - ostikwhy-blip/claude-code-handoff-skill: A Claude Code skill that writes a verified HANDOFF.md so a fresh session can resume long-running work without repeating dead ends. · GitHub"
author: "ostik"
captured_at: "2026-07-05T17:57:39Z"
capture_tool: "hn-digest"
hn_id: 48795956
score: 2
comments: 0
posted_at: "2026-07-05T17:18:16Z"
tags:
  - hacker-news
  - translated
---

# Handoff – a verified context bridge between Claude Code sessions

- HN: [48795956](https://news.ycombinator.com/item?id=48795956)
- Source: [github.com](https://github.com/ostikwhy-blip/claude-code-handoff-skill)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T17:18:16Z

## Translation

タイトル: ハンドオフ – クロード コード セッション間の検証済みコンテキスト ブリッジ
記事のタイトル: GitHub - ostikwhy-blip/claude-code-handoff-skill: 検証済みの HANDOFF.md を作成するクロード コード スキル。これにより、行き止まりを繰り返すことなく、新しいセッションで長時間実行される作業を再開できます。 · GitHub
説明: 検証済みの HANDOFF.md を書き込むクロード コード スキル。これにより、行き止まりを繰り返すことなく、新しいセッションで長時間実行される作業を再開できます。 - ostikwhy-blip/claude-code-handoff-skill

記事本文:
GitHub - ostikwhy-blip/claude-code-handoff-skill: 検証済みの HANDOFF.md を書き込むクロード コード スキル。これにより、行き止まりを繰り返すことなく、新しいセッションで長時間実行される作業を再開できます。 · GitHub
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
ostikwhy-blip
/
クロードコードハンドオフスキル
公共
N

通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ostikwhy-blip/claude-code-handoff-skill
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット アセット アセット ライセンス ライセンス README.md README.md SKILL.md SKILL.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code セッション間のクリーンなブリッジ。
長いクロード コードのセッションは腐敗します。コンテキストがいっぱいになり、モデルが決定を忘れ始め、すでに失敗したアプローチを再試行し、そのように見えなくなったファイルを記述します。新しいセッションを開始すると腐敗は修正されますが、古いセッションで学んだことはすべて捨てられます。
handoff は検証済みの HANDOFF.md をプロジェクトのルートに書き込むため、行き止まりを繰り返すことなく、次のセッションが中断したところから再開されます。
乱雑な 3 時間のセッション → /handoff → 1 つのクリーンなファイルから次のセッションが再開されます。
git clone https://github.com/ostikwhy-blip/claude-code-handoff-skill.git
cp -r クロードコード-ハンドオフ-スキル ~ /.claude/skills/handoff
新しいクロード コード セッションを開始します — 完了です。依存関係、外部 API、構成はありません。
cp -r claude-code-handoff-skill your-project/.claude/skills/handoff
使用する
セッションの終了時（または気分が悪くなり始めたとき）：
/ハンドオフ
Claude はリポジトリの実際の状態を確認し、 HANDOFF.md を書き込みます。
ハンドオフから再開する
クロードはファイルを読み、現在のリポジトリと照合して再チェックし、未解決の質問を明らかにし、計画をあなたに確認します。そして、前回何が失敗したかをすでに知っているので、作業を続行します。
それがワークフロー全体です。コマンドは 2 つ。
クロードに「セッションの要約」を依頼してみてはいかがでしょうか?
その概要は、劣化したモデルのメモリから、そのメモリの信頼性が最も低くなった瞬間に書き込まれるためです。自信を持ったフィクションが得られます。「合格するテスト」

s」が再実行されることはなく、20 回編集されたファイルの説明は古く、幻覚のような進行状況です。
ハンドオフは 1 つのルールに基づいて構築されています。つまり、メモリだけからファイルには何も書き込まれません。
一行書く前に、クロードはこう言いました。
git status 、 git log 、 git diff を実行して、変更されたと思われる内容ではなく、実際に変更された内容を再構築します。
ハンドオフで言及されるすべてのファイルを再読み取りします (記憶の場所にないファイルは捕らえられた幻覚であると判断され、修正されます)
テストを再実行します。「テストに合格」は、ハンドオフ自体中に生成された出力からのみ書き込まれます。
結果のファイル内のすべてのクレームにはタグが付けられます。
次のセッションが再開されると、現在のリポジトリに対して [V] クレームが再検証され、ドリフトがあれば明示的に報告されます。ブリッジの両端の確認。
HANDOFF.md の上限は 250 ～ 400 行です。重要な内容を伝えるのに十分な長さであり、新しいセッションで実際に吸収できる程度の短さです。内容は次のとおりです。
現在の目標と検証された状態 — 意図した状態ではなく、実際の状態
下された決定とその理由 - 再訴訟されないように
失敗したアプローチと失敗した正確な理由 - 最も価値のあるセクション。これは、新しいセッションだけでは再構築できない知識であり、このスキルが存在する理由です。
このセッションのコードベースで発見された既知のトラップ
次のステップを指示し、再検討することなく行動できるほど具体的である
あなただけが答えられる未解決の質問
すべてのエントリーは 1 つの品質ゲートを通過する必要があります。この行だけを持っている見知らぬ人が、私たちが犯した間違いを避けることができるでしょうか?そうでない場合は、ファイル、行番号、および理由を取得するか、カットされます。
古いハンドオフは削除されません。それらは .handoffs/ にアーカイブされ、依然として関連する行き止まりは自動的に引き継がれます。
バックグラウンドメモリではありません。アンビエント メモリ ツールは目に見えずに機能します。何が伝わるかは決してわかりません。 HANDOFF.md は pl です

ain ファイルは、読み取り、編集、コミットできます。クロードが何か間違った場合は、次のセッションに支障をきたす前に修正します。
丸太考古学ではありません。過去のセッション ログを検索するツールは、事後的に生の履歴を調査します。ハンドオフは、ハンドオーバーの瞬間に作成される、意図的に厳選され、検証された成果物です。
損失のある要約ではありません。重要だと思われる内容は要約に残しておきます。引き継ぎは、再発見に費用がかかるもの、とりわけ行き止まりを保持します。
ボーナス: 劣化を早期に発見します
クロードがセッション中に自身のコンテキストの腐敗の兆候に気付いた場合、つまり以前の決定に矛盾し、既に決定されたものを再導出する場合、一文で引き継ぎを提案し、あなたの確認を待ちます。自動的に実行されることはありません。
ハンドオフ/
§── SKILL.md # スキル定義
└── 資産/
└── HANDOFF.template.md # すべてのハンドオフが従う構造
SKILL.md を読んで、クロードが何をするように指示されているかを正確に確認してください。拒否するよう指示されている正当化のリストも含まれています (「セッションのことははっきりと覚えています。再検証する必要はありません」→ とにかく検証してください)。
問題や PR は歓迎します。特に、うまくいかなかったセッションをどのように処理したかについてのレポートです。そのために作られているのです。
行き止まりを繰り返すことなく、新しいセッションで長時間実行される作業を再開できるように、検証済みの HANDOFF.md を書き込むクロード コード スキル。
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

A Claude Code skill that writes a verified HANDOFF.md so a fresh session can resume long-running work without repeating dead ends. - ostikwhy-blip/claude-code-handoff-skill

GitHub - ostikwhy-blip/claude-code-handoff-skill: A Claude Code skill that writes a verified HANDOFF.md so a fresh session can resume long-running work without repeating dead ends. · GitHub
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
ostikwhy-blip
/
claude-code-handoff-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
ostikwhy-blip/claude-code-handoff-skill
master Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits assets assets LICENSE LICENSE README.md README.md SKILL.md SKILL.md View all files Repository files navigation
A clean bridge between Claude Code sessions.
Long Claude Code sessions rot: context fills up, the model starts forgetting decisions, re-trying approaches that already failed, and describing files that no longer look like that. Starting a fresh session fixes the rot — but throws away everything the old session learned.
handoff writes a verified HANDOFF.md at your project root so the next session picks up exactly where this one left off, without repeating your dead ends.
A messy 3-hour session → /handoff → one clean file the next session resumes from.
git clone https://github.com/ostikwhy-blip/claude-code-handoff-skill.git
cp -r claude-code-handoff-skill ~ /.claude/skills/handoff
Start a new Claude Code session — done. No dependencies, no external APIs, no configuration.
cp -r claude-code-handoff-skill your-project/.claude/skills/handoff
Use
End of a session (or when it starts feeling off):
/handoff
Claude verifies the actual state of your repo, then writes HANDOFF.md .
resume from the handoff
Claude reads the file, re-checks it against the current repo, surfaces open questions, confirms the plan with you — then continues the work, already knowing what failed last time.
That's the whole workflow. Two commands.
Why not just ask Claude to "summarize the session"?
Because that summary is written from a degraded model's memory, at exactly the moment that memory is least reliable. You get confident fiction: tests that "pass" but were never re-run, file descriptions that are twenty edits stale, hallucinated progress.
handoff is built around one rule: nothing goes in the file from memory alone.
Before writing a single line, Claude:
Runs git status , git log , git diff to reconstruct what actually changed — not what it believes changed
Re-reads every file the handoff will mention (a file that isn't where memory says is a caught hallucination, and gets corrected)
Re-runs your tests — "tests pass" is only written from output produced during the handoff itself
Every claim in the resulting file is tagged:
When the next session resumes, it re-verifies the [V] claims against the current repo and reports any drift explicitly. Verification on both ends of the bridge.
HANDOFF.md is capped at ~250–400 lines — long enough to carry what matters, short enough that a fresh session can actually absorb it. It covers:
Current goal and verified state — the actual state, not the intended one
Decisions made and why — so they don't get relitigated
Failed approaches and exactly why they failed — the single most valuable section. This is the knowledge a fresh session cannot reconstruct on its own, and the reason this skill exists
Known traps discovered in the codebase this session
Ordered next steps , concrete enough to act on without re-exploring
Open questions only you can answer
Every entry has to pass one quality gate: would a stranger with only this line avoid the mistake we made? If not, it gets a file, a line number, and a reason — or it gets cut.
Old handoffs aren't deleted — they're archived to .handoffs/ , and still-relevant dead ends are carried forward automatically.
Not a background memory. Ambient memory tools work invisibly; you never see what gets passed along. HANDOFF.md is a plain file you can read, edit, and commit. If Claude got something wrong, you fix it before it poisons the next session.
Not log archaeology. Tools that search past session logs dig through raw history after the fact. A handoff is a deliberate, curated, verified artifact written at the moment of handover.
Not a lossy summary. Summaries keep what sounds important. A handoff keeps what's expensive to rediscover — above all, the dead ends.
Bonus: it catches degradation early
If Claude notices symptoms of its own context rot mid-session — contradicting an earlier decision, re-deriving something already settled — it suggests a handoff in one sentence and waits for your confirmation. It never runs automatically.
handoff/
├── SKILL.md # the skill definition
└── assets/
└── HANDOFF.template.md # the structure every handoff follows
Read SKILL.md to see exactly what Claude is instructed to do — including the list of rationalizations it's told to refuse ("I remember the session clearly, no need to re-verify" → verify anyway).
Issues and PRs welcome — especially reports of how it handles sessions that went badly. That's the case it's built for.
A Claude Code skill that writes a verified HANDOFF.md so a fresh session can resume long-running work without repeating dead ends.
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
