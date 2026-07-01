---
source: "https://github.com/Toadoum/ai-research-skill"
hn_url: "https://news.ycombinator.com/item?id=48747615"
title: "Recursive AI Research Skill for Claude Code / OpenClaw / Codex"
article_title: "GitHub - Toadoum/ai-research-skill · GitHub"
author: "Toadoum"
captured_at: "2026-07-01T14:58:23Z"
capture_tool: "hn-digest"
hn_id: 48747615
score: 2
comments: 0
posted_at: "2026-07-01T14:40:29Z"
tags:
  - hacker-news
  - translated
---

# Recursive AI Research Skill for Claude Code / OpenClaw / Codex

- HN: [48747615](https://news.ycombinator.com/item?id=48747615)
- Source: [github.com](https://github.com/Toadoum/ai-research-skill)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T14:40:29Z

## Translation

タイトル: Claude Code / OpenClaw / Codex の再帰的 AI 研究スキル
記事タイトル: GitHub - Toadoum/ai-research-skill · GitHub
説明: GitHub でアカウントを作成して、Toadoum/ai-research-skill の開発に貢献します。

記事本文:
GitHub - Toadoum/ai-research-skill · GitHub
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
キノコ
/
AI研究スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く

フォルダーとファイル
1 コミット 1 コミット 参照 参照 スクリプト スクリプト AGENTS.md AGENTS.md LESSONS.md LESSONS.md LICENSE LICENSE README.md README.md SKILL.md SKILL.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
🔬 AI Research Skill — クロード、コーデックス、OpenClaw の自己改善型リサーチ エージェント
SKILL.md を使用すると、あらゆるコーディング エージェントが ML/AI 研究ループ全体を迅速に進めることができ、ミスをするたびに改善されます。
仮説→文献レビュー→ベースラインの再現→漏れのない実験→正直な分析→論文。 Claude / Claude Code、OpenAI Codex、および OpenClaw で動作します。再帰的な設計: 自身の間違いをルールとして記録するので、同じ間違いを繰り返さないようにします。
⭐ これにより、ラベルの結果が 1 つ漏洩しないようにするには、リポジトリにスターを付けてください。他の研究者がリポジトリを見つけるのに役立ちます。
AI エージェントの「研究」ヘルプのほとんどは、存在しない論文を引用しながら自信満々に話すチャットボットです。実際の調査は、具体的で退屈で高価な方法で失敗します。実行の代わりに引用したベースライン、ラベルを密かに漏洩しているメトリクス、本当に幸運の種である「利益」、記憶からでっち上げられた引用などです。
このスキルは、これらの失敗を 1 か月かかる前に検出する規律をエンコードし、ポータブル SKILL.md としてエージェントが自動的に読み取ります。そしてそれは再帰的です。エージェントが間違いを犯して修正すると、レッスンを LESSONS.md に書き込み、今後のすべてのタスクの開始時にそのファイルを読み取り、繰り返しを停止します。 3 か月目に使用するスキルは、インストールしたスキルよりも鋭くなっています。
ステージ
エージェントが行うこと
ガードレールを施行します
フレーム
漠然としたアイデアを検証可能な仮説 + 以前の研究と比較した明白なデルタに変換します
主張が一文になるまでは実験しない
レビュー
重要な 5 ～ 15 の論文を見つけて、比較マトリックスを構築します
実際に関連する論文のみを引用する

広告 — 決して記憶からではありません
再現する
最初にセットアップで最も強力なベースラインを実行します
ゲインを測定する前に定規が必要です
デザイン
シードを設定し、分割を修正し、完全な漏洩監査を実行します
疑わしいほど良好 ≠ 画期的 — 漏洩ではないことを証明する
走る
スキャフォールド構成により、すべての数値が再現可能になります
設定は唯一の信頼できる情報源です
分析する
3 つ以上のシードにわたる平均 ± 標準偏差でベースラインと比較します
シングルシードでの勝利は発見ではなく物語である
書く
すべての主張を数字で裏付けます。再現性チェックリストを同梱
ストーリーを損なうシード/データセットを決してドロップしないでください
クイックスタート
クローンを作成し、エージェントがスキルを探す場所にドロップします。
git clone https://github.com/ < ユーザー名 > /ai-research-skill.git
クロード / クロード・コード / クロード・コワーク
フォルダー (またはパッケージ化された .skill バンドル) をスキル ディレクトリにインストールします。クロードは常に名前と説明をコンテキスト内に保持し、タスクが AI/ML 研究のような場合には完全なスキルを読み込みます。あとは普通に作業するだけです。「この論文のベースラインを再現するのを手伝ってください」「F1 が疑わしいほど高いのはなぜですか?」 —そしてそれが始まります。
cp -r ai-research-skill ~ /.openclaw/workspace/skills/ai-research
OpenClaw は同じ SKILL.md フロントマッター + 本体を読み取り、挿入された AGENTS.md パスは同じ場所に配置されます。スキルをインストールする前に検査します。コミュニティ スキルを見知らぬ人からの npm パッケージのように扱います。
AGENTS.md をリポジトリのルートに保持します (これは SKILL.md への薄いポインタです)。 Codex は AGENTS.md を読み取り、そこからスキルに従います。
自己改善のループ（興味深い部分）
タスクを開始する ─▶ LESSONS.md を読む ─ 調べます ─▶ 間違いを犯しましたか？
▲ │ はい
│ ▼
└─── LESSONS.md に新しいルールが追加されました ◀── log_lesson.py
エージェントがエラーを検出すると、次のように実行されます。
Python スクリプト/log_lesson.py \
--trigger " が報告されました

1 回のトレーニング実行から得られる利益 " \
--mistake " 単一のシードから「ベースラインを上回る」と主張しました " \
--fix " 3 つのシードを再実行しました。ゲインはノイズ帯域内にありました " \
--rule " 3 つ以上のシードにわたる平均 ± 標準値のない比較主張は禁止します" \
--タグ「シード、再現性」
スクリプトは同様のルールの重複を排除し、繰り返しをカウントし、ルールにフラグを立てます。
3 回表示されると SKILL.md に昇格し、ログがいつ必要になるかを通知します。
剪定。組み込まれたガードレールは、研究を弱体化させるあらゆる「教訓」を拒否します
整合性 (結果の捏造、隠蔽、または厳選)。ループにより、
より賢いスキルを身につければ、不誠実になることはありません。
ai-研究スキル/
§── SKILL.md # スキル (真実の源、クロード + OpenClaw)
§── AGENTS.md # Codex / OpenClaw 用のシン ポインタ
§── LESSONS.md # 再帰的記憶 — 間違い→ルール
§── スクリプト/
│ §── log_lesson.py # 重複排除され、カウントされたレッスンを追加します
│ └── new_experiment.py # 再現可能な実験ディレクトリを足場にします
━── 参考文献/
§── Literature-review.md # 先行技術を迅速に検索します。マトリックスを構築する
§── Experiment-design.md # 再現、シーズ、漏洩監査
└── Paper-writing.md # 主張→証拠 + 再現性チェックリスト
貢献する
PR 歓迎 - 特に実際の研究ミスからの新しい LESSONS.md エントリ
(それが重要です)、新しいリファレンス プレイブック、およびより多くのエージェント用のアダプターです。
ヒットした障害モードとそれを修正するルールに関する問題を開きます。
AI研究エージェント・機械学習研究助手・クロードスキル・クロード
コードスキル · SKILL.md · Codex AGENTS.md · OpenClaw スキル · 自己改善エージェント ·
再帰的 AI エージェント · 再現可能な ML · データ漏洩検出 · 実験
追跡 · アブレーション研究 · 文献レビューの自動化 · 論文作成 · LLM
研究ワークフロー・NLPリサーチ

アーチ・研究の再現性。
MIT — 使って、フォークして、出荷してください。
レビューよりも初日にリークを発見したい研究者向けに構築されています。
⭐役に立ったらスターを付けてください。
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

Contribute to Toadoum/ai-research-skill development by creating an account on GitHub.

GitHub - Toadoum/ai-research-skill · GitHub
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
Toadoum
/
ai-research-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit references references scripts scripts AGENTS.md AGENTS.md LESSONS.md LESSONS.md LICENSE LICENSE README.md README.md SKILL.md SKILL.md View all files Repository files navigation
🔬 AI Research Skill — a self-improving research agent for Claude, Codex & OpenClaw
One SKILL.md that makes any coding agent move fast through the full ML/AI research loop — and get better every time it makes a mistake.
Hypothesis → literature review → reproduce baseline → leak-free experiments → honest analysis → paper. Works with Claude / Claude Code, OpenAI Codex , and OpenClaw . Recursive by design: it logs its own mistakes as rules so it never repeats them.
⭐ If this saves you from one leaked-label result, please star the repo — it helps other researchers find it.
Most AI-agent "research" help is a chatbot that sounds confident and cites papers that don't exist. Real research fails in specific, boring, expensive ways: a baseline you quoted instead of ran, a metric that's secretly leaking the label, a "gain" that's really one lucky seed, a citation invented from memory.
This skill encodes the discipline that catches those failures before they cost you a month — as a portable SKILL.md your agent reads automatically. And it's recursive : when the agent makes a mistake and fixes it, it writes the lesson to LESSONS.md , reads that file at the start of every future task, and stops repeating itself. The skill you use in month three is sharper than the one you installed.
Stage
What the agent does
Guardrail it enforces
Frame
Turns a vague idea into a testable hypothesis + a stated delta vs prior work
No experiment until the claim is one sentence
Review
Finds the 5–15 papers that matter, builds a comparison matrix
Cite only papers actually read — never from memory
Reproduce
Runs the strongest baseline on your setup first
You need a ruler before you measure a gain
Design
Sets seeds, fixes splits, runs a full leakage audit
Suspiciously-good ≠ breakthrough — prove it's not leakage
Run
Scaffolds configs so every number is reproducible
The config is the single source of truth
Analyze
Compares vs baseline with mean ± std over ≥3 seeds
A single-seed win is a story, not a finding
Write
Backs every claim with a number; ships a reproducibility checklist
Never drop the seed/dataset that hurt the story
Quickstart
Clone it, then drop it where your agent looks for skills:
git clone https://github.com/ < your-username > /ai-research-skill.git
Claude / Claude Code / Claude Cowork
Install the folder (or a packaged .skill bundle) into your skills directory. Claude keeps the name + description in context always, and loads the full skill when your task looks like AI/ML research. Then just work normally — "help me reproduce this paper's baseline" , "why is my F1 suspiciously high?" — and it kicks in.
cp -r ai-research-skill ~ /.openclaw/workspace/skills/ai-research
OpenClaw reads the same SKILL.md frontmatter + body, and its injected AGENTS.md path lands in the same place. Inspect any skill before installing it — treat community skills like npm packages from strangers.
Keep AGENTS.md at your repo root (it's a thin pointer to SKILL.md ). Codex reads AGENTS.md and follows the skill from there.
The self-improving loop (the interesting part)
start task ─▶ read LESSONS.md ─▶ do research ─▶ made a mistake?
▲ │ yes
│ ▼
└────────── LESSONS.md now has a new rule ◀── log_lesson.py
When the agent catches an error, it runs:
python scripts/log_lesson.py \
--trigger " reported a gain from one training run " \
--mistake " claimed 'beats baseline' from a single seed " \
--fix " re-ran 3 seeds; gain was inside the noise band " \
--rule " no comparison claim without mean ± std over >=3 seeds " \
--tags " seeds,reproducibility "
The script dedupes similar rules, counts repeats, flags a rule for
promotion into SKILL.md once it's seen 3×, and tells you when the log needs
pruning . A built-in guardrail refuses any "lesson" that would weaken research
integrity (fabricating, hiding, or cherry-picking results). The loop makes the
skill smarter — it can't make it dishonest.
ai-research-skill/
├── SKILL.md # the skill (source of truth, Claude + OpenClaw)
├── AGENTS.md # thin pointer for Codex / OpenClaw
├── LESSONS.md # recursive memory — mistakes → rules
├── scripts/
│ ├── log_lesson.py # append a deduped, counted lesson
│ └── new_experiment.py # scaffold a reproducible experiment dir
└── references/
├── literature-review.md # find prior art fast; build the matrix
├── experiment-design.md # reproduce, seeds, the leakage audit
└── paper-writing.md # claim→evidence + reproducibility checklist
Contributing
PRs welcome — especially new LESSONS.md entries from real research mistakes
(that's the whole point), new reference playbooks, and adapters for more agents.
Open an issue with the failure mode you hit and the rule that fixes it.
AI research agent · machine learning research assistant · Claude skill · Claude
Code skill · SKILL.md · Codex AGENTS.md · OpenClaw skill · self-improving agent ·
recursive AI agent · reproducible ML · data leakage detection · experiment
tracking · ablation study · literature review automation · paper writing · LLM
research workflow · NLP research · research reproducibility.
MIT — use it, fork it, ship it.
Built for researchers who'd rather find the leak on day one than in review.
⭐ Star it if it helps.
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
