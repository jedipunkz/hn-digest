---
source: "https://github.com/virgiliojr94/book-to-skill"
hn_url: "https://news.ycombinator.com/item?id=48649087"
title: "Book-to-skill: Turn a technical book PDF into a Claude Code skill"
article_title: "GitHub - virgiliojr94/book-to-skill: Turn any technical book PDF into a Claude Code skill — ready to study, reference, and use while you work. · GitHub"
author: "ethanpil"
captured_at: "2026-06-23T18:38:51Z"
capture_tool: "hn-digest"
hn_id: 48649087
score: 1
comments: 0
posted_at: "2026-06-23T18:18:29Z"
tags:
  - hacker-news
  - translated
---

# Book-to-skill: Turn a technical book PDF into a Claude Code skill

- HN: [48649087](https://news.ycombinator.com/item?id=48649087)
- Source: [github.com](https://github.com/virgiliojr94/book-to-skill)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T18:18:29Z

## Translation

タイトル: Book-to-skill: 技術書の PDF を Claude Code スキルに変える
記事のタイトル: GitHub - virgiliojr94/book-to-skill: あらゆる技術書籍の PDF をクロード コード スキルに変換します。作業中にすぐに学習、参照、使用できます。 · GitHub
説明: あらゆる技術書籍の PDF をクロード コード スキルに変換します。作業中にすぐに学習、参照、使用できます。 - virgiliojr94/本とスキル

記事本文:
GitHub - virgiliojr94/book-to-skill: あらゆる技術書籍の PDF をクロード コード スキルに変換します。作業中にすぐに学習、参照、使用できます。 · GitHub
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
ヴァージリオjr94
/
本からスキルまで
公共
通知
通知を変更するにはサインインする必要があります

フィクション設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
73 コミット 73 コミット .github .github book_to_skill book_to_skill ドキュメント ドキュメント スクリプト スクリプト テスト テスト ツール ツール .gitignore .gitignore BACKERS.md BACKERS.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md LICENSE.md LICENSE.md README.md README.md SECURITY.md SECURITY.md SKILL.md SKILL.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あらゆる技術書籍、ドキュメント フォルダー、またはソースのコレクションを統合エージェント スキルに変換します。これにより、GitHub Copilot CLI、Amp、または Claude Code での作業中にすぐに学習、参照、使用できるようになります。
なぜ・
それが生み出すもの ·
本を超えて ·
使い方・
要件 ·
仕組み ·
ディスカバリーループ税 ·
よくある質問・
インストール・
変更履歴 ·
パフォーマンス・
建築
素晴らしい技術書を購入しました。一度は読んだことがあるでしょう。 3か月後、第7章が存在したことを思い出せなくなります。
通常の回避策は役に立ちません。
📄 「PDF を検索させてください」 → 答えではなくページのリストが表示されます
🧠 「この本についてエージェントに聞いてみます」 → 幻覚が見えるか、内容がないと言われます
📝 「読みながらメモを取ります」 → 二度と開かない 200 行のドキュメントが作成される
book-to-skill は、書籍を、エージェントがオンデマンドで読み込む構造化スキルに変換することで、この問題を解決します。
インストールしたら、「/your-book-slug replication」と入力するだけで、エージェントが適切な章を読み取って、実際のコンテンツから回答します。幻覚はありません。 PDF を調べる必要はありません。本がワークフローの一部になります。
オープンなエージェント スキル標準をサポートする任意のホストで動作します。GitHub Copilot CLI、Amp、および Claude Code はすべて同じ SKILL.md 形式を読み取ります。
/book-to-skill your-book.pdf (または fo

lder、glob、またはファイルのリスト）は、エージェントのスキル ディレクトリに完全なスキルを作成します（Copilot CLI の場合は ~/.copilot/skills/<slug>/、Amp またはクロスエージェントの場合は ~/.agents/skills/<slug>/、クロード コードの場合は ~/.claude/skills/<slug>/）。
チャプター ファイルはオンデマンドで読み込まれます。そのトピックについて質問するまでは、スキル バジェットにカウントされません。
名前には「本」とありますが、入力されるのは構造化された散文です。あなたが所有し、常に再読している知識に対しても同じ抽出が機能します。
内部ドキュメント — アーキテクチャの決定記録、ランブック、オンボーディング ガイド。 docs/ フォルダー全体を 1 つのスキルにまとめて、コーディング中にそれを尋ねます。
ブランドおよびデザイン システム — 音声ガイドライン、声のトーンに関するドキュメント、コンポーネントの原則。 60 ページの PDF をざっと読むのではなく、ブランド ブックをチームがクエリするスキルに変えます。
研究クラスター — 論文の束と自分のメモを 1 つの統合スキルに統合し、新しいマテリアルランドとして更新します (「更新 / 折りたたみ」を参照)。
仕様と標準 - RFC、API 契約、参照はするものの決して覚えていないコンプライアンス ドキュメント。
覚えておきたいと思うほど頻繁に文書を再度開く場合、それは候補です。
/book-to-skill <ドキュメントフォルダーまたはグロブへのパス>... [スキル名スラッグ]
サポートされているドキュメント形式: PDF、EPUB、DOCX、TXT、Markdown、reStructuredText、AsciiDoc、HTML、RTF、MOBI/AZW/AZW3。
# 複数のファイルをまとめて統合スキルに処理します
/book-to-skill ~ /papers/paper1.pdf ~ /notes/export.txt 統合リサーチ
# フォルダー内のサポートされているすべてのファイルをまとめて処理します
/book-to-skill ~ /workspace/project-docs/ project-knowledge
# グロブパターンに一致するファイルを処理する
/book-to-skill " ~/books/*.epub " my-library
# 新しいマテリアルを既存のスキルフォルダーに更新/折りたたみます
/book-to-skill ~ /articles/new-paper.pdf ~ /.claude/skills/project-knowledge
スキルを作成したら、それを使用します

他のエージェント スキルと同様に、次のようになります。
/designing-data-integral-apps # コアメンタルモデルをロードする
/designing-data-integral-apps replication # トピックを見つけて説明する
/designing-data-集約型アプリ ch05 # 第 5 章に進みます
/designing-data-integral-apps " どの章がありますか? "
GitHub Copilot CLI では、ファイルの書き込み後に /skills reload を実行して、新しいスキルが /skills list に表示されるようにする必要がある場合があります。クロード・コードとアンプは次のセッションでそれを拾います。
抽出プログラムはフォーマットごとにツールを順番に試し、最初に利用可能なものを使用します。何もインストールされていない場合は、実行するコマンドが表示されます。プレーン テキスト、Markdown、reStructuredText、AsciiDoc には追加の deps は必要ありません。
1 つのコマンドでセットアップをチェックします: python3 scripts/extract.py --check は、各形式にどのエクストラクターがインストールされているか、および不足しているものをインストールするための正確なコマンドを出力します。ファイルは必要ありません。
抽出が開始される前に、スキルはその本が専門的なものか、テキストが多いのかを尋ね、適切なツールを自動的に選択します。ドクリングはマークダウン テーブルとコード ブロックを保持します。散文のみの本の場合は pdftotext の方が高速です。
1 つのファイル、フォルダー、グロブ、パスのリスト
│
▼
ステップ 1.5 — 「技術書ですか、それとも文章量の多い本ですか?」
│
§── テクニカル → ドクリング (テーブル + コードブロックをマークダウンとして、~1.5 秒/ページ)
└── text → pdftotext → pypdf → pdfminer (インスタント)
│
▼
scripts/extract.py <パス…> --mode <技術|テキスト>
ソースごと: PDF → pdftotext/Docling · EPUB → ebooklib → stdlib zipfile · DOCX/HTML/RTF/…
(1 つの悪いソースは警告を出してスキップされますが、残りは引き続き処理されます)
│
§── /tmp/book_skill_work/full_text.txt (すべてのソースがマージされ、ソース マーカーが付けられます)
└── /tmp/book_skill_work/metadata.json (集計された統計 + ソースごとの配列)
│
▼
クロードは構造を分析する
(タイトル、著者、章、目次 — スパン

すべてのソースを調べます)
── または、既存のスキルをターゲットとする場合: 新しいコンテンツを (モード 4) に折りたたみます。
│
▼
章ごとの要約を生成します (それぞれ 800 ～ 1,200 トークン)
技術的な → コード例と参照表のセクションが含まれています
用語集、パターン、チートシートを生成します
コアメンタルモデルを含むマスターSKILL.mdを生成
│
▼
スキルは次のいずれかに書き込まれます。
~/.copilot/skills/<slug>/ (GitHub Copilot CLI)
~/.agents/skills/<slug>/ (Copilot CLI または Amp、クロスエージェント)
~/.claude/skills/<slug>/ (クロードコード)
/tmp/book_skill_work/ 🗑️ クリーンアップしました
抽出ベンチマーク (103 ページの技術書、CPU のみ):
実際のコンバージョン (測定: ページ、抽出されたトークン、自動検出された章、
Claude Sonnet 4.5 の推定ワンパスコストは MTok あたり $3/$15):
† 章の自動検出には、明示的な Chapter N / Capítulo N の見出しが必要です。プロギット
はセクションタイトルを使用し、Moby-Dick は章タイトル/ローマ数字を使用するため、どちらも使用しません
自動セグメント — 抽出と変換は引き続き機能しますが、セクションをポイントします。
手動で。完全なスキルを習得するには、1 冊あたり約 1 ドルの費用がかかります。再読するよりはるかに少ない
セッションごとに PDF を作成します。
完全性よりも密度 — 1,000 トークンの要約は 10,000 トークンの抜粋よりも優れています
実践者の声 — 「本では X について説明している」ではなく、「Y の場合は X を使用してください」
フロントローディング SKILL.md — 圧縮により最初の ~5,000 トークンが保持されます。最も重要なコンテンツが最初に来る
オンデマンドの章 — トピック インデックスは、どのファイルを読むべきかをクロードに伝えます。章は必要な場合にのみロードされます
決して生のテキストではなく、常にソースから信号を合成、要約、抽出します
PDF 読み取りエージェントは単に読み取るだけではなく、ナビゲーションも行います。一つ質問してみると、
目次を取得し、定義できない用語に気づき、さらにページを取得し、
バックトラック。これらのホップはすべて会話履歴に記録され、
後続のターンごとに再処理されます。

予算内に収まるように、サブエージェント
その後、読み取った内容をひどい比率で圧縮することを強制され、メインエージェントに
劣化した要約では、ソースに対して事実確認ができません。
book-to-skill は、コンパイル時にナビゲーション コストを 1 回支払います。実行時には、
アシスタントは、小さな常駐コアと、必要な 1 つのプリコンパイル済みチャプターをロードします。
検出ループや圧縮による適合はなく、抽出された完全なソースはディスク上に残ります。
検証用に。
主張するのではなく、測定する。 tools/discovery_tax.py を実行する
3 冊の本物の本 — 単一の対象を絞った質問に答えるためにコンテキストに入るトークン
(ブックからスキルまで = 常駐コア + コンパイルされた 1 章 ≈ 5,000 トークン):
利点は章のサイズに応じて変化します。コンテキストダンプに対して、一貫して利点があります。
24–51× (そしてそのコストは毎ターン繰り返されます);ワンタイム検出ループに対して
小さな章で構成される本の 2.4 倍から、大きな章の 1 つで 15.6 倍までの範囲です。
章。自分の本に再現してください:
python3 tools/discovery_tax.py --full-text /tmp/book_skill_work/full_text.txt --target-chapter 5
正直な注意点: (1) 発見の数値は 1 回限りのコストであり、モデルです。
本の実際の目次/章のサイズを使用する — よく調整されたエージェントは、最良の内容に近づきます。
場合。対照的に、コンテキストダンプコストはターンごとに繰り返されます。 (2) ツール
本を分割するには明示的な第 N 章/第 N 章の見出しが必要です。タイトルのみ
またはローマ数字の書籍 (および ebooklib なしで抽出された EPUB) はセグメント化されません。
きれいに。何度も知識に立ち返れば、本からスキルへの効果が得られます。のために
1 回限りの読み取りであれば、プレーン PDF エージェントで十分です。
「PDF/EPUB をクロード プロジェクト コンテキストにダンプすることはできないでしょうか?」
可能ですが、会話ごとにトークンの予算が前もって消費されてしまいます。 400 ページの本は約 200,000 トークンです。スキルを使用すると、質問に関連する章だけが読み込まれます - 通常

つまり、SKILL.md コア (~4K) と、質問した 1 つの章 (~1K) です。残りは必要になるまでディスク上に残ります。
経済性は規模ではなく償却です。ブックを貼り付けると、セッションのたびにトークンの請求額全額が永久に支払われます。 book-to-skill は抽出コストを 1 回支払うだけで、その後の会話では必要なスライスのみが読み込まれます。コンテキスト ウィンドウが大きいほど、これは重要になります。ウィンドウが大きいとダンプが可能になりますが、コストはかかりません。
さらに重要なことは、生のテキストの挿入は取得であるということです。スキルは推理力です。章ファイルをロードするとき、クロードはキーワードの一致を検索しません。読み取り用ではなくアプリケーション用に構造化された、事前に抽出された名前付きフレームワーク、原則、メンタル モデルを使用して作業します。
「クロードには 1M トークンのコンテキスト ウィンドウが表示されました。本全体を読み込んだままにしておくことはできませんか?」
ウィンドウが大きくなると、何がスマートかではなく、何が適合するかが変わります。代替品ではない 3 つの理由:
料金はトークンごと、通話ごとに支払います。 100 万のウィンドウではこれらのトークンが無料になるわけではありません。高額な定期的な請求が可能になります。スキルはメガバイトではなくキロバイトをロードします。
フィルを使用するとリコールが低下します。モデルは、ほぼ完全なコンテキストに埋もれた特定のファクトを取得する精度を失います (「途中で失われた」)。 1K の厳選された章は、1 つの質問に答えるための 200K の生の散文よりも優れています。
窓≠構造。コンテキスト内の完全な本は依然として生のテキストであり、モデルは毎ターン再解析する必要があります。スキルは事前に出荷されます

[切り捨てられた]

## Original Extract

Turn any technical book PDF into a Claude Code skill — ready to study, reference, and use while you work. - virgiliojr94/book-to-skill

GitHub - virgiliojr94/book-to-skill: Turn any technical book PDF into a Claude Code skill — ready to study, reference, and use while you work. · GitHub
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
virgiliojr94
/
book-to-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
73 Commits 73 Commits .github .github book_to_skill book_to_skill docs docs scripts scripts tests tests tools tools .gitignore .gitignore BACKERS.md BACKERS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE.md LICENSE.md README.md README.md SECURITY.md SECURITY.md SKILL.md SKILL.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml View all files Repository files navigation
Turn any technical book, document folder, or collection of sources into a unified agent skill — ready to study, reference, and use while you work in GitHub Copilot CLI, Amp, or Claude Code.
Why ·
What it generates ·
Beyond books ·
Usage ·
Requirements ·
How it works ·
Discovery Loop Tax ·
FAQ ·
Install ·
Changelog ·
Performance ·
Architecture
You buy a great technical book. You read it once. Three months later you can't remember chapter 7 existed.
The usual workarounds don't help:
📄 "Let me just search the PDF" → you get a list of pages, not answers
🧠 "I'll ask the agent about this book" → it either hallucinates or says it doesn't have the content
📝 "I'll take notes as I read" → you end up with a 200-line doc you never open again
book-to-skill solves this by turning the book into a structured skill your agent loads on demand.
Once installed, you just type /your-book-slug replication and the agent reads the right chapter and answers from the actual content. No hallucination. No digging through PDFs. The book becomes part of your workflow.
Works with any host that supports the open Agent Skills standard — GitHub Copilot CLI, Amp, and Claude Code all read the same SKILL.md format.
Running /book-to-skill your-book.pdf (or a folder, glob, or list of files) creates a full skill in your agent's skills directory ( ~/.copilot/skills/<slug>/ for Copilot CLI, ~/.agents/skills/<slug>/ for Amp or cross-agent, ~/.claude/skills/<slug>/ for Claude Code):
Chapter files are loaded on-demand — they don't count against the skill budget until you ask about that topic.
The name says "book", but the input is any structured prose. The same extraction works on knowledge you own and re-read constantly:
Internal documentation — architecture decision records, runbooks, onboarding guides. Fold a whole docs/ folder into one skill and ask it while you code.
Brand & design systems — voice guidelines, tone-of-voice docs, component principles. Turn a brand book into a skill your team queries instead of skimming a 60-page PDF.
Research clusters — a stack of papers plus your own notes, merged into a single unified skill and updated as new material lands (see Update / fold-in ).
Specs & standards — RFCs, API contracts, compliance docs you reference but never memorize.
If you re-open a document often enough to wish you'd memorized it, it's a candidate.
/book-to-skill <path-to-document-folder-or-glob>... [skill-name-slug]
Supported document formats: PDF, EPUB, DOCX, TXT, Markdown, reStructuredText, AsciiDoc, HTML, RTF, MOBI/AZW/AZW3.
# Process several files together into a unified skill
/book-to-skill ~ /papers/paper1.pdf ~ /notes/export.txt unified-research
# Process all supported files in a folder together
/book-to-skill ~ /workspace/project-docs/ project-knowledge
# Process files matching a glob pattern
/book-to-skill " ~/books/*.epub " my-library
# Update/fold new material into an existing skill folder
/book-to-skill ~ /articles/new-paper.pdf ~ /.claude/skills/project-knowledge
After the skill is created, use it like any other agent skill:
/designing-data-intensive-apps # load core mental models
/designing-data-intensive-apps replication # find and explain a topic
/designing-data-intensive-apps ch05 # dive into chapter 5
/designing-data-intensive-apps " what chapters do you have? "
In GitHub Copilot CLI you may need to run /skills reload after the file is written so the new skill appears in /skills list . Claude Code and Amp pick it up on the next session.
The extractor tries tools in order per format and uses the first available. If nothing is installed, it tells you which command to run. Plain text, Markdown, reStructuredText and AsciiDoc need no extra deps.
Check your setup in one command: python3 scripts/extract.py --check prints which extractors are installed for every format and the exact command to install anything missing — no file needed.
Before extraction begins, the skill asks you whether the book is technical or text-heavy and picks the right tool automatically. Docling preserves markdown tables and code blocks; pdftotext is faster for prose-only books.
One file · a folder · a glob · a list of paths
│
▼
Step 1.5 — "Technical or text-heavy book?"
│
├── technical → Docling (tables + code blocks as markdown, ~1.5s/page)
└── text → pdftotext → pypdf → pdfminer (instant)
│
▼
scripts/extract.py <paths…> --mode <technical|text>
per source: PDF → pdftotext/Docling · EPUB → ebooklib → stdlib zipfile · DOCX/HTML/RTF/…
(one bad source is skipped with a warning; the rest still process)
│
├── /tmp/book_skill_work/full_text.txt (all sources merged, with source markers)
└── /tmp/book_skill_work/metadata.json (aggregated stats + per-source array)
│
▼
Claude analyzes structure
(title, author, chapters, ToC — spanning all sources)
── or, if targeting an existing skill: folds new content in (Mode 4)
│
▼
Generates per-chapter summaries (800–1,200 tokens each)
technical → includes Code Examples + Reference Tables sections
Generates glossary, patterns, cheatsheet
Generates master SKILL.md with core mental models
│
▼
Skill written to one of:
~/.copilot/skills/<slug>/ (GitHub Copilot CLI)
~/.agents/skills/<slug>/ (Copilot CLI or Amp, cross-agent)
~/.claude/skills/<slug>/ (Claude Code)
/tmp/book_skill_work/ 🗑️ cleaned up
Extraction benchmark (103-page technical book, CPU only):
Real conversions (measured: pages, extracted tokens, chapters auto-detected,
estimated one-pass cost on Claude Sonnet 4.5 at $3/$15 per MTok):
† Chapter auto-detection needs explicit Chapter N / Capítulo N headings. Pro Git
uses section titles and Moby-Dick uses chapter titles / roman numerals, so neither
auto-segments — extraction and conversion still work, but you point at sections
manually. A full skill costs roughly $1 per book ; far less than re-reading the
PDF every session.
Density over completeness — a 1,000-token summary beats a 10,000-token excerpt
Practitioner voice — "Use X when Y", not "The book explains X"
Front-loaded SKILL.md — compaction keeps the first ~5,000 tokens; the most important content comes first
On-demand chapters — the topic index tells Claude which file to read; chapters load only when needed
Never raw text — always synthesize, summarize, extract signal from the source
A PDF-reading agent doesn't just read — it navigates . Ask it one question and it
fetches the table of contents, notices a term it can't define, pulls more pages,
backtracks. Every one of those hops lands in the conversation history and gets
re-processed on every subsequent turn . To stay inside its budget, a sub-agent
is then forced to compress what it read at brutal ratios, handing the main agent a
degraded summary it can't fact-check against the source.
book-to-skill pays the navigation cost once, at compile time . At runtime the
assistant loads a small resident core plus the one pre-compiled chapter it needs —
no discovery loop, no compress-to-fit, and the full extracted source stays on disk
for verification.
Measured, not asserted. Running tools/discovery_tax.py
on three real books — tokens entering context to answer a single targeted question
(book-to-skill = resident core + one compiled chapter ≈ 5,000 tokens):
The advantage scales with chapter size : against a context-dump it's consistently
24–51× (and that cost recurs every turn ); against a one-time discovery loop it
ranges from a modest 2.4× on a book of small chapters to 15.6× on one of large
chapters. Reproduce on your own book:
python3 tools/discovery_tax.py --full-text /tmp/book_skill_work/full_text.txt --target-chapter 5
Honest caveats: (1) the discovery figures are a one-time cost and a model
using the book's real ToC/chapter sizes — a well-tuned agent lands nearer the best
case; the context-dump cost, by contrast, recurs on every turn. (2) The tool
needs explicit Chapter N / Capítulo N headings to segment a book; titles-only
or roman-numeral books (and EPUBs extracted without ebooklib ) won't segment
cleanly. book-to-skill wins when you return to the knowledge repeatedly; for a
single one-off read, a plain PDF agent is fine.
"Can't I just dump the PDF/EPUB into my Claude project context?"
You can — but every conversation will burn that token budget upfront. A 400-page book is ~200K tokens. With a skill, only the chapters relevant to your question load — typically a SKILL.md core (~4K) plus the one chapter you asked about (~1K). The rest stays on disk until you need it.
The economics are amortization, not size. Pasting the book pays the full token bill on every turn of every session, forever . book-to-skill pays the extraction cost once and every future conversation loads only the slice it needs. The bigger your context window, the more this matters — a large window makes the dump possible , not cheap .
More importantly: raw text injection is retrieval. A skill is reasoning. When you load a chapter file, Claude isn't searching for keyword matches — it's working with pre-extracted named frameworks, principles, and mental models structured for application, not for reading.
"Claude has a 1M-token context window now — can't I just keep the whole book loaded?"
A bigger window changes what fits , not what's smart . Three reasons it isn't a substitute:
You pay per token, per call. A 1M window doesn't make those tokens free — it makes a large, recurring bill possible. The skill loads kilobytes, not megabytes.
Recall degrades with fill. Models lose precision retrieving a specific fact buried in a near-full context ("lost in the middle"). A 1K curated chapter beats 200K of raw prose for answering one question.
Window ≠ structure. A full book in context is still raw text the model must re-parse every turn. The skill ships pre-e

[truncated]
