---
source: "https://github.com/grandimam/suki"
hn_url: "https://news.ycombinator.com/item?id=49337613"
title: "Show HN: Personal AI tutor that builds and probes your understanding"
article_title: "GitHub - grandimam/suki: Build expertise. Validate it. Own it. · GitHub"
image: "https://opengraph.githubassets.com/4117f73576d64cb24d6692ffc44da4dd33209e6c7784c8f904d759da7ac1b4c7/grandimam/suki"
author: "grandimam"
captured_at: "2026-08-17T21:17:03Z"
capture_tool: "hn-digest"
hn_id: 49337613
score: 1
comments: 0
posted_at: "2026-08-17T21:04:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Personal AI tutor that builds and probes your understanding

- HN: [49337613](https://news.ycombinator.com/item?id=49337613)
- Source: [github.com](https://github.com/grandimam/suki)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T21:04:38Z

## Translation

タイトル: HN を表示: 理解を深めて探求するパーソナル AI 家庭教師
記事のタイトル: GitHub - grandimam/suki: 専門知識を構築する。それを検証してください。それを所有してください。 · GitHub
説明: 専門知識を構築します。それを検証してください。それを所有してください。 GitHub でアカウントを作成して、grandimam/suki の開発に貢献してください。
HN text: 以前は本や講座で勉強していました。しかし、私は本当のメンタルモデルを構築していないことに気づき続けました。本で読み飛ばされた部分を補うために、常に LLM が必要でした。そこで、私は直接の真実の情報源として書籍を使用するのをやめ、LLM 自体を使用し始めました。つまり、カリキュラムを設計し、トピックを理解しているかどうかを調査するためです。 Suki - それはオープンソースです。与えられたトピックに対して、初心者から上級者までのカリキュラムを構築し、章ごとに理解を深めます。リンク - https://github.com/grandimam/suki 皆さんのご意見をお聞かせください。

記事本文:
GitHub - grandimam/suki: 専門知識を構築します。それを検証してください。それを所有してください。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
おばあちゃん
/
すき
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット ブック ブック キャリア キャリア カリキュラム カリキュラム ホーム ラーニング プローブ プローブ 履歴書 履歴書 履歴書 src/ スキ src/ スキ テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md README.md README.md pyproject.toml pyproject.toml すべて表示

ファイル リポジトリ ファイルのナビゲーション
専門知識を構築します。それを検証してください。それを所有してください。
╭─────────────────────╮
│すき│
│ 専門知識を構築します。それを検証してください。それを所有してください。 │
╰───────────────────╯
Suki は AI コーディング エージェントを信頼できない学習パートナーに変える
あなた。道徳的な意味ではありません。それはあなたが何かを知っていると信じることを拒否します
それを読んだり、コースを視聴したり、チュートリアルに沿ってうなずいたりしたからです。
学習とはコンテンツを消費することではありません。言い返さざるをえないから
自分の頭がつかまるまで。 Suki はあなたにそうさせると、それを覚えます
自分の苦手なものを、消えてしまう前に取り戻します。
🔒 ローカルファースト · 🛡️ デフォルトでプライベート · 🧩 3 ステップ、1 つのコマンド · 🐍 Python 3 + エージェント
ループ: カリキュラム → 調査 → 本
┌─────────────────────┐
│ │
▼ │
┌─────┐ ┌─────┐ ┌─────┐ │
│ カリキュラム│ ──► │ プローブ │ ──► │ 本 │ │
│ ビルドする │ │ テストする │ │ 公開する │──────┘
│ パス │ │ & 修正 │ │ それ │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
/suki カリキュラム <トピック> — 最終的な学習パスを設計します。
章、順序、トピックにとっての「マスター済み」の意味。
/suki プローブ <トピック> [ch] — 次の章を作成して理解度を確認します
章。修正を強制し、テストし、亀裂を修復する

、次にスケジュールを設定します
腐らないようにレビューします。
/suki book <トピック> — カリキュラム、作業モデル、および
歴史を調査して出版物と同等の品質の教材を作成します。
作業コンパウンド: 各ステップでアーティファクトが ~/.suki/topics/<slug>/ に書き込まれます。
次のステップでそれが読み取られます。来週もう一度調査すると、本にそれが反映されます。
/スキカリキュラムアクティブディレクトリ
→ 第 1 層: 基本 (AD とは何か、ドメインの仕組み)
→ 第 2 層: コアメカニズム (Kerberos、チケット、トラスト)
→ 階層 3: 攻撃 (ケルベロスティング、ゴールデン チケット、委任)
→ 階層 4: 防御 + 検出
→ 階層 5: エキスパートエッジ (レッドチームのトレードクラフト、現実世界の作戦)
+ 各章の「マスター済み」の意味
+ 実際に理解を構築する順序
Suki は、絶対的な基礎から最も深い専門家までの道を設計します。
フィールドに、章ごとのマークダウンと
後の各ステップで読み取られるカリキュラム.json。
/suki プローブアクティブディレクトリ 1.1
「Kerberos 認証を自分の言葉で説明してください。」
→あなたの写真を明記してください
→ スキは亀裂がないか調べます
→弱点？ターゲットを絞ったドリルがその場で修正します
→マスターした？ 3 → 10 → 30 → 90 日後に再訪問
モデルは本ではなくあなたの言葉にある必要があります。それが続かない場合
アップすると、ギャップが露出し、インラインで修復されます。あなたが苦手なことは、
記憶が消え去る前に再び浮上する。あなたに不利な点は何もありません。
すべては複合的なアーティファクトとして記録されます。
/suki book active-directory # -> ~/.suki/topics/active-directory/book/book.pdf
→ カリキュラム（各階層ごとに 1 つのパートとして）
→ 作業モデルは永続化されます
→ プローブ + 修復履歴
→ Pandoc + LaTeX によるタイプセット (KOMA scrbook、Palatino)
あなたが構築し検証したものはすべて、本物の教材になります。
他の学習者が使用できる資料、およびあなたが実際に持っているものの証拠。
1つ

コマンド /suki はステータス ダッシュボードであり、あらゆるものへのルーターです
それ以外は。どこで中断したのか、次に何をする必要があるのか​​がわかります。
╭───────────────────╮
│すき│
│ 専門知識を構築します。それを検証してください。それを所有してください。 │
╰───────────────────╯
📊 ステータス
─────────────────
学習 3 トピック · 1 つは復習期限
⏳ 審査期限中
─────────────────
• active-directory 2.1 (Kerberos 認証) は本日リリース予定です
• Python 3.1 (Asyncio) は 2D でリリース予定
次のステップ
→ /suki プローブアクティブディレクトリ 2.1
─────────────────
クイックスタート
pip インストールスキ
uki install --all # スキルをオープンコード、クロード、コーデックスにリンクします
それだけです。エージェントを再起動し、次の手順を実行します。
/suki カリキュラム python # パスを構築する
/suki プローブ Python 1.1 # ビルド + 理解を確認する
/suki book python # マスターしたものを公開する
🧠 スキル
スキル
役割
すき
ステータス ダッシュボード + ルーター (単一のエントリ ポイント)
カリキュラム
最終的な学習パスを設計する
学ぶ
学習者の視点から教材をレビューする
プローブ
構築 + 理解の検証、亀裂の修復、間隔をあけた繰り返し
本
トピックを教材としてレンダリングする
🗂️ 州
すべては ~/.suki/ の下にあり、トピックごとに 1 つのフォルダーが存在します。現在の状態は
JSON;履歴は追加専用の JSONL です。何も削除されないため、完全な
あなたの学習の弧は

回復可能。
～/.すき/
└── topic/<slug>/ #カリキュラム.json、mastery.json、プローブ、ブック/
uki CLI はスタックを管理します。
uki install [--opencode|--claude|--codex|--all] — スキルをエージェントにリンクします
スキステータス [トピック] — 進行状況 + 間隔をあけた繰り返しの期限ステータス
スキブック <トピック> — トピックを本としてレンダリングします
pip インストールスキ
スキインストール --all
それが全体のセットアップです。 pip install uki は uki コマンドを提供します。
スキインストールは 5 つのスキルすべてをリンクします (スキ、カリキュラム、学習、
プローブ、ブック）をエージェントに問い合わせます。デフォルトは --all 、またはあなたのものだけを選択してください
エージェントには --opencode 、 --claude 、または --codex を指定します。
インストール後にエージェントを再起動し、/suki を単一のエントリとして使用します。
ポイント。サブコマンドをルーティングします: /suki カリキュラム <トピック> 、
/スキ ラーニング <ドラフト> 、/スキ プローブ <トピック> [ch] 、/スキ ブック <トピック> 、
ステータス ダッシュボードの場合は /suki のみを使用します。
git clone < このリポジトリ > && cd スキ
pip install -e 。
⚙️ 要件
コア スキルには、エージェントと uki CLI (Python 3) 以外には何も必要ありません。
/suki book には、pandoc と xelatex を使用した LaTeX ディストリビューションも必要です。
醸造インストール Pandoc
brew install --cask mactex-no-gui # または任意の TeX Live インストール
🔐 プライバシー
すべてがマシン上に残ります。カリキュラム、探査の歴史、書籍
あなたが公開しても、あなたのコンピュータからは何も出ません。唯一の外部呼び出しは、
LLM/ハーネスと Web フェッチ。
まずは残酷な正直さ : モデルについての真実を知る
モデルはあなたの言葉の中にあります。あなたはそれを言い直します。それはあなたのものになります
継続して学習: 詰め込みではなく間隔をあけて繰り返し
すべてはアーティファクトです: 名前付きファイル、セッション間で複合化
デフォルトのプライバシー : データがマシンから離れることはありません
Suki: 専門知識を構築します。それを検証してください。それを所有してください。
専門知識を構築します。それを検証してください。それを所有してください。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私のものを共有しないでください

個人情報

## Original Extract

Build expertise. Validate it. Own it. Contribute to grandimam/suki development by creating an account on GitHub.

I used to learn from books and courses. But I kept realising I hadn't built a real mental model. I always needed an LLM on the side to fill in what the books skipped. So, I stopped using books as my direct source of truth and started using LLMs itself: to design the curriculum and probe whether I understood the topic. Suki - it's opensource. Given a topic it builds a curriculum from beginner to advanced, then probes your understanding chapter by chapter. Link - https://github.com/grandimam/suki Let me know what you guys think about it.

GitHub - grandimam/suki: Build expertise. Validate it. Own it. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
grandimam
/
suki
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits book book career career curriculum curriculum home home learn learn probe probe resume resume src/ suki src/ suki test test .gitignore .gitignore AGENTS.md AGENTS.md README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Build expertise. Validate it. Own it.
╭────────────────────────────────────────────╮
│ s u k i │
│ Build expertise. Validate it. Own it. │
╰────────────────────────────────────────────╯
Suki turns an AI coding agent into a learning partner that doesn't trust
you. Not in a moral sense. It refuses to believe you know something just
because you read it, watched a course, or nodded along with a tutorial.
Learning isn't consuming content. It's being forced to say it back, from
your own head, until it holds. Suki makes you do that, then it remembers
what you're weak at and brings it back before it fades.
🔒 Local-first · 🛡️ Private by default · 🧩 3 steps, one command · 🐍 Python 3 + your agent
The loop: curriculum → probe → book
┌──────────────────────────────────────────────┐
│ │
▼ │
┌──────────┐ ┌──────────┐ ┌──────────┐ │
│ curriculum│ ──► │ probe │ ──► │ book │ │
│ build the │ │ test it │ │ publish │──────┘
│ path │ │ & fix │ │ it │
└──────────┘ └──────────┘ └──────────┘
/suki curriculum <topic> — design the definitive learning path: the
chapters, the order, what "mastered" means for your topic.
/suki probe <topic> [ch] — build and verify understanding chapter by
chapter. Force a restatement, test it, repair the cracks, then schedule
reviews so it doesn't rot.
/suki book <topic> — render your curriculum, working models, and
probe history into a publication-quality teaching book.
Work compounds: each step writes an artifact to ~/.suki/topics/<slug>/
that the next step reads. Probe again next week, and the book reflects it.
/suki curriculum active-directory
→ tier 1: The basics (what AD is, how a domain works)
→ tier 2: Core mechanics (Kerberos, tickets, trusts)
→ tier 3: Attacks (Kerberoasting, golden tickets, delegation)
→ tier 4: Defenses + detection
→ tier 5: Expert edge (red-team tradecraft, real-world ops)
+ what "mastered" means at every chapter
+ the order that actually builds understanding
Suki designs the path from absolute basics to the deepest expert end of the
field, then writes it as chapter-by-chapter markdown plus a
curriculum.json that every later step reads.
/suki probe active-directory 1.1
"Explain Kerberos authentication in your own words."
→ you state your picture
→ suki probes it for cracks
→ weak spot? a targeted drill fixes it right there
→ mastered? revisit in 3 → 10 → 30 → 90 days
The model has to be in your words, not the book's. If it doesn't hold
up, the gap is exposed and repaired inline. What you're weak at is
remembered and resurfaced before it fades. Nothing is graded against you;
everything is recorded as an artifact that compounds.
/suki book active-directory # -> ~/.suki/topics/active-directory/book/book.pdf
→ your curriculum, as one part per tier
→ your working models, persisted
→ your probe + remediation history
→ typeset with pandoc + LaTeX (KOMA scrbook, Palatino)
Everything you built and validated becomes a real teaching book: the same
material other learners can use, and the proof of what you actually hold.
One command, /suki , is the status dashboard and the router to everything
else. It knows where you left off and what's due next.
╭──────────────────────────────────────────╮
│ s u k i │
│ Build expertise. Validate it. Own it. │
╰──────────────────────────────────────────╯
📊 Status
────────────────────────────────────────────
Learning 3 topics · 1 due for review
⏳ Due for review
────────────────────────────────────────────
• active-directory 2.1 (Kerberos auth) due today
• python 3.1 (Asyncio) due in 2d
Next step
→ /suki probe active-directory 2.1
────────────────────────────────────────────
Quick start
pip install suki
suki install --all # links the skills into opencode, claude, and codex
That's it. Restart your agent, then:
/suki curriculum python # build the path
/suki probe python 1.1 # build + verify understanding
/suki book python # publish what you mastered
🧠 The skills
Skill
Role
suki
Status dashboard + router (single entry point)
curriculum
Design a definitive learning path
learn
Review material from a learner's perspective
probe
Build + verify understanding, repair cracks, spaced repetition
book
Render a topic as a teaching book
🗂️ State
Everything lives under ~/.suki/ , one folder per topic. Current state is
JSON; history is append-only JSONL. Nothing is ever deleted, so the full
arc of your learning is recoverable.
~/.suki/
└── topics/<slug>/ # curriculum.json, mastery.json, probes, book/
The suki CLI manages the stack:
suki install [--opencode|--claude|--codex|--all] — link skills into your agent
suki status [topic] — progress + spaced-repetition due status
suki book <topic> — render a topic as a book
pip install suki
suki install --all
That's the whole setup. pip install suki gives you the suki command;
suki install links all five skills ( suki , curriculum , learn ,
probe , book ) into your agent. Default to --all , or pick just your
agent with --opencode , --claude , or --codex .
Restart your agent after installing, then use /suki as the single entry
point. It routes subcommands: /suki curriculum <topic> ,
/suki learn <draft> , /suki probe <topic> [ch] , /suki book <topic> ,
or /suki alone for the status dashboard.
git clone < this repo > && cd suki
pip install -e .
⚙️ Requirements
The core skills need nothing beyond the agent and the suki CLI (Python 3).
/suki book also needs pandoc and a LaTeX distribution with xelatex :
brew install pandoc
brew install --cask mactex-no-gui # or any TeX Live install
🔐 Privacy
Everything stays on your machine. Curriculum, probe history, and the books
you publish: none of it leaves your computer. The only external calls are
the LLM/harness and web fetches.
Brutal honesty first : Know the truth about your model
The model is in your words : You restate it; it becomes yours
Learn for keeps : Spaced repetition over cramming
Everything is an artifact : Named files, compounded across sessions
Privacy by default : Your data never leaves your machine
Suki: Build expertise. Validate it. Own it.
Build expertise. Validate it. Own it.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
