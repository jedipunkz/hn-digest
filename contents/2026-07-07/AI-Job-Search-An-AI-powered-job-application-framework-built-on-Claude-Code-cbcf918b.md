---
source: "https://github.com/MadsLorentzen/ai-job-search"
hn_url: "https://news.ycombinator.com/item?id=48822056"
title: "AI Job Search – An AI-powered job application framework built on Claude Code"
article_title: "GitHub - MadsLorentzen/ai-job-search: AI-powered job application framework built on Claude Code. Fork it, fill in your profile, and let Claude evaluate jobs, tailor CVs, write cover letters, and prepare you for interviews. · GitHub"
author: "anonli"
captured_at: "2026-07-07T19:44:06Z"
capture_tool: "hn-digest"
hn_id: 48822056
score: 1
comments: 0
posted_at: "2026-07-07T18:58:08Z"
tags:
  - hacker-news
  - translated
---

# AI Job Search – An AI-powered job application framework built on Claude Code

- HN: [48822056](https://news.ycombinator.com/item?id=48822056)
- Source: [github.com](https://github.com/MadsLorentzen/ai-job-search)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T18:58:08Z

## Translation

タイトル: AI Job Search – クロード コードに基づいて構築された AI を活用した求職フレームワーク
記事のタイトル: GitHub - MadsLorentzen/ai-job-search: Claude Code に基づいて構築された AI を活用した求人応募フレームワーク。それをフォークしてプロフィールを記入し、クロードに仕事を評価してもらい、履歴書を調整し、カバーレターを書き、面接の準備をしてもらいましょう。 · GitHub
説明: Claude Code に基づいて構築された、AI を活用した求人応募フレームワーク。それをフォークしてプロフィールを記入し、クロードに仕事を評価してもらい、履歴書を調整し、カバーレターを書き、面接の準備をしてもらいましょう。 - MadsLorentzen/ai-job-search

記事本文:
GitHub - MadsLorentzen/ai-job-search: Claude Code に基づいて構築された AI を活用した求人応募フレームワーク。それをフォークしてプロフィールを記入し、クロードに仕事を評価してもらい、履歴書を調整し、カバーレターを書き、面接の準備をしてもらいましょう。 · GitHub
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
マッツロア

ンツェン
/
AI求人検索
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット .agents/ スキル .agents/ スキル .claude .claude .github .github cover_letters cover_letters cv cv ドキュメント ドキュメント job_scraper job_scraper テンプレート テンプレート ツール ツール スキルアップスキル .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md SETUP.md SETUP.md claude_animation.gif claude_animation.gif job_search_tracker.csv job_search_tracker.csv給与_lookup.py給与_lookup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code に基づいて構築された AI を活用した求人応募フレームワーク。それをフォークしてプロフィールを記入し、Claude に求人情報を評価させ、履歴書を調整し、カバーレターを書き、面接の準備をさせます。
注: これは独立したオープンソース プロジェクトであり、Anthropic との提携、承認、後援、または維持は行われていません。 Anthropic コードと Claude Code は、このワークフローが使用するツールチェーンを説明するためにのみ参照されています。
Claude Code をフルスタックの求人応募アシスタントに変える構造化されたワークフロー。コア ワークフロー (セルフ プロファイリング、適合性評価、およびドラフト作成者とレビュー担当者のアプリケーション パイプライン) は、言語や国に依存しません。求人ポータルの検索スキルはデンマーク市場 (Jobindex、Jobnet、Akademikernes Jobbank など) 向けに構築されていますが、パターンは地元の求人サイトに置き換えられるように設計されています。
/setup /scrape /apply <URL>
| | |
vvv
入力してください 求人を検索します 適合性を評価します
あなたのプロフィールポータルのスコアと推奨
| | |
vvv
プロフィール 現在の履歴書とカバーレターのドラフトが一致します
適合性評価を備えたファイルの準備ができています (LaTeX、調整済み)
| |

v v
一致するものを選択 レビューアー エージェントの批評
-> /apply -> 修正 -> 最終出力
このフレームワークは、構造化された評価基準、将来を見据えたカバーレターの構成、オプションの給与ベンチマークなど、キャリア ガイダンスのベスト プラクティスをコード化しています。
Bun (デンマークの求人検索 CLI ツール用)
lualatex および xelatex を使用した LaTeX ディストリビューション: TeX Live または MiKTeX。 CV は lualatex でコンパイルされます (pdflatex は、最近の MiKTeX インストールでは、fontawesome5 フォント拡張エラーで失敗することがよくあります)。 cover.cls には fontspec が必要なため、カバーレターは xelatex でコンパイルされます。
オプション: Poppler からの pdftotext (macOS: brew install Poppler 、Debian/Ubuntu: apt install Poppler-utils 、Windows: choco install Poppler ) — コンパイルされた CV に対する /apply の ATS 解析可能性チェックによって使用されます。不足している場合、チェックは視覚的なキーワードのレビューに正常に降格します。
gh リポジトリ フォーク MadsLorentzen/ai-job-search --clone
cd ai求人検索
2. 求人検索ツールをインストールする
cd .agents/skills/jobbank-search/cli && bun install && cd ../../../..
cd .agents/skills/jobdanmark-search/cli && bun install && cd ../../../..
cd .agents/skills/jobindex-search/cli && bun install && cd ../../../..
cd .agents/skills/jobnet-search/cli && bun install && cd ../../../..
cd .agents/skills/linkedin-search/cli && bun install && cd ../../../..
linkedin-search の場合、インストールはオプションです。ランタイム依存関係はなく、プレーン bun で実行されます。 bun install は TypeScript 開発タイプのみをプルします。
クロード
# 次にクロード コード内:
/セットアップ
/setup では 3 つのパスが提供されます。documents/ フォルダーが入力されている場合は読み取り (CV PDF、LinkedIn エクスポート、卒業証書、推薦状、過去の申請書)、チャットに貼り付けられた 1 つの履歴書をインポートするか、面接をウォークスルーします。あなたが何を持っているかを自動的に検出して尋ねます。ドキュメント フォルダー モードは冪等であり、マテリアルを追加しても安全に再実行できます。

アル;レイアウトについては、documents/README.md を参照してください。
/スクレープ
これにより、複数の求人ポータルであなたのプロフィールに一致するポジションが検索され、結果の重複が削除され、適合性順に並べ替えられて表示されます。直接実行する /apply する一致を選択します。または、スクレイピングが目視したいよりも多くのジョブを返した場合は、/rank を実行して、fit フレームワークに対してすべてのジョブをバッチスコアリングし、最初にランク付けされた候補リストを取得します。
/適用 https://jobindex.dk/job/1234567
URL を取得できない場合 (一部の求人ポータルが自動アクセスをブロックしている場合)、代わりに求人説明を直接貼り付けることができます。
/apply < 完全な職務説明をここに貼り付けます >
これにより、完全なワークフローが実行されます。適合性の評価、履歴書とカバーレターの草案、2 番目のエージェントによるレビュー、修正、最終出力の提示です。
/setup 、 /scrape 、および /apply はコア ワークフローを形成します。プロファイルを作成したら、さらに 6 つのコマンドでそれを拡張します。
/outcome は、面接段階、内定、不採用、沈黙など、応募に何が起こったかを記録します。提出された履歴書、カバーレター、および投稿テキストを document/applications/<company>_<role>/ にアーカイブし、outcome.md を /setup 形式で保持します。 パス A が解析し、トラッカーを更新します。いくつかのアプリケーションが解決されると、/setup に戻り、実際にインタビューされた内容に基づいて適合フレームワークを調整します。
/rank ブリッジ /scrape および /apply : Fit フレームワークに対して新しくスクレイピングされたすべての投稿をバッチスコアリングし (並列エージェントが各投稿を取得し、5 つの評価次元をスコアリングします)、ジョブごとの正直な強みとギャップを含むランク付けされた候補リストを返します。取引違反者は拒否権を発動し、締め切りには緊急フラグが立てられ、期限切れのマークが付けられない投稿には期限切れのマークが付けられます。番号を選択すると、完全な /apply ワークフローに渡されます。
/expand は、既にリンクされている公開ソース (GitHub リポジトリ、ポートフォリオ サイト、Kaggle、Google Scholar) をスキャンして検索することで、プロフィールを充実させます。

指定されたコースと認定資格のシラバス。発見されたコンピテンシーは、ソース タグとともにプロフィールに追加されます。 /setup の直後に、文書だけでは明示されていないスキルを明らかにするのに役立ちます。
/upskill は、あなたのプロフィールと追跡された求人情報 (または /upskill <URL> を介した単一の投稿) との間のギャップを分析します。スキルギャップの優先順位付けされたヒートマップと、Web 検索の学習リソースと推定時間による学習計画を作成します。アプリケーション間のキャリア計画に役立ちます。
/add-template は、ストックされているものの代わりに、独自の LaTeX 履歴書またはカバーレターのテンプレートを登録します。テンプレートの指示 (コンパイル エンジン、フォント、スタイル ルール、ページ制限) を取得し、必須のテスト コンパイルを実行して、テンプレートを /apply に接続します。以下の LaTeX テンプレートを参照してください。
/add-portal は、市場の求人サイト用の求人ポータル検索スキルを生成します。ポータル (検索 URL パターン、結果構造、アクセス ルール) を調査し、出荷されたものと同じ構造から CLI スキルをスキャフォールディングし、登録前にライブ クエリをテスト実行します。以下の求人検索ツールをご覧ください。
/reset も使用できます。以下の「最初からやり直す」を参照してください。
AI求人検索/
§── CLAUDE.md # 主な候補者プロフィール + ワークフロールール
§── .claude/
│ §── コマンド/
│ │ §── apply.md # /apply ワークフロー (製図者レビュー担当者)
│ │ §── setup.md # /setup オンボーディング (ドキュメント フォルダー、CV インポート、または面接)
│ │ §── Expand.md # /expand ドキュメントとオンライン プレゼンスからコンピテンシーを強化
│ │ §── add-template.md # /add-template カスタム LaTeX テンプレートを登録
│ │ §── add-portal.md # /add-portal 市場向けの求人ポータル検索スキルを生成します
│ │ §── Rank.md # /rank ジョブを優先順位付けしてランク付けされた候補リストに絞り込みます
│ │ §── 結果.md # /結果レコード

申請結果、アーカイブ資料
│ │ └──reset.md # /reset プロファイル データまたはドキュメント フォルダーをワイプ
│ §── スキル/
│ │ §── job-application-assistant/ # コアアプリケーションスキル
│ │ │ §── SKILL.md # スキル定義
│ │ │ §── 01-candidate-profile.md # あなたの学歴、経験、スキル
│ │ │ §── 02-behavioral-profile.md# PI/DISC/性格評価
│ │ │ §── 03-writing-style.md # 口調、構成、すべきこと、してはいけないこと
│ │ │ §── 04-job-evaluation.md # 職務適合性のスコアリング フレームワーク
│ │ │ §── 05-cv-templates.md # LaTeX CV 構造 + 調整ルール
│ │ │ §── 06-cover-letter-templates.md # LaTeX カバーレターテンプレート
│ │ │ └─ 07-interview-prep.md # STAR 例 + 面接フレームワーク
│ │ §── job-scraper/ # 就職オーケストレーション
│ │ └─ upskill/ # /upskill スキルギャップ分析と学習計画
│ └── settings.json # クロードコードの権限 (共有、範囲指定)
§── .agents/skills/ # 求人ポータル CLI ツール
│ §── jobbank-search/ # Akademikernes Jobbank (デンマーク)
│ §── jobdanmark-search/ # Jobdanmark.dk (デンマーク)
│ §── jobindex-search/ # Jobindex.dk (デンマーク)
│ §── jobnet-search/ # Jobnet.dk (デンマーク、政府ポータル)
│ └── linkedin-search/ # LinkedIn の公開求人情報 (国に依存しない)
§── CV/
│ └── main_example.tex # moderncv LaTeX テンプレート
§── カバーレター/
│ §── cover.cls # カスタムカバーレター LaTeX クラス
│ └── OpenFonts/ # Lato + Raleway フォント
§── templates/ # /add- 経由で登録されたカスタム テンプレート
[切り捨てられた]
/apply コマンドは、必須の PDF コンパイルを伴う製図者とレビュー者のワークフローを実行します。
求人情報 (URL またはテキスト) を解析します。
fiを評価する

自分のプロフィール (スキル、経験、文化、場所、キャリアの整合性) に反するものではない
LaTeX でカスタマイズされた履歴書とカバーレターの草案を作成する
会社を調査し、草案を批評するレビュー担当者エージェントを生成します。
査読者のフィードバックに基づいて改訂する
両方の PDF をコンパイルして検査します。履歴書には lualatex、カバーレターには xelatex を使用します。クロードは、レンダリングされたページを読み取り、履歴書が孤立したエントリ タイトルのない正確に 2 ページになり、カバーレターが署名が表示され、フォントが一貫して正確に 1 ページになるまで、LaTeX を反復処理します。
ATS による履歴書のチェック: PDF のテキスト レイヤー ( pdftotext 、オプションの依存関係) を抽出し、ATS パーサーが認識する方法でそれを検証します (連絡先の詳細がリテラル テキストとして表示され、グリフの文字化けがなく、読み上げ順序が正常である)。その後、抽出に対して投稿のキーワード カバレッジをスコア付けします。プロファイルが実際にサポートするキーワードが追加されます。本物の隙間は目に見えたままになり、決して埋まることはありません。
検証チェックリストを使用して最終出力を提示します。
履歴書とカバーレター内のすべての主張は、実際のプロフィールと照合して検証されます。このシステムはスキルや経験を捏造することはありません。
このワークフローの違い
PDF 検証ループ。ほとんどの LaTeX 履歴書テンプレートは、PDF 内で中断される「.tex では正常に見える」出力を生成します。役職名が次のページに孤立し、カバーレターが 2 ページにはみ出し、箇条書きフォントが本文フォントに静かにフォールバックします。 /apply コマンドは、すべての PDF をコンパイルして視覚的に検査し、レイアウトがきれいになるまで対象の修正 ( \needspace 、 \enlargethispage 、リスト項目のフォント マッチング ラッパー) を適用します。これ

[切り捨てられた]

## Original Extract

AI-powered job application framework built on Claude Code. Fork it, fill in your profile, and let Claude evaluate jobs, tailor CVs, write cover letters, and prepare you for interviews. - MadsLorentzen/ai-job-search

GitHub - MadsLorentzen/ai-job-search: AI-powered job application framework built on Claude Code. Fork it, fill in your profile, and let Claude evaluate jobs, tailor CVs, write cover letters, and prepare you for interviews. · GitHub
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
MadsLorentzen
/
ai-job-search
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits .agents/ skills .agents/ skills .claude .claude .github .github cover_letters cover_letters cv cv documents documents job_scraper job_scraper templates templates tools tools upskill upskill .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md SETUP.md SETUP.md claude_animation.gif claude_animation.gif job_search_tracker.csv job_search_tracker.csv salary_lookup.py salary_lookup.py View all files Repository files navigation
An AI-powered job application framework built on Claude Code . Fork it, fill in your profile, and let Claude evaluate job postings, tailor your CV, write cover letters, and prepare you for interviews.
Note: This is an independent open-source project and is not affiliated with, endorsed by, sponsored by, or maintained by Anthropic. Anthropic and Claude Code are referenced only to describe the toolchain this workflow uses.
A structured workflow that turns Claude Code into a full-stack job application assistant. The core workflow (self-profiling, fit evaluation, and the drafter-reviewer application pipeline) is language- and country-agnostic . The job portal search skills are built for the Danish market (Jobindex, Jobnet, Akademikernes Jobbank, etc.), but the pattern is designed to be swapped for your local job boards.
/setup /scrape /apply <url>
| | |
v v v
Fill in Search job Evaluate fit
your profile portals Score & recommend
| | |
v v v
Profile Present matches Draft CV + Cover Letter
files ready with fit ratings (LaTeX, tailored)
| |
v v
Pick a match Reviewer agent critiques
-> /apply -> Revise -> Final output
The framework encodes career guidance best practices, including structured evaluation criteria, forward-looking cover letter framing, and optional salary benchmarking.
Bun (for Danish job search CLI tools)
LaTeX distribution with lualatex and xelatex : TeX Live or MiKTeX . The CV compiles with lualatex (pdflatex often fails on modern MiKTeX installs with fontawesome5 font-expansion errors); the cover letter compiles with xelatex because cover.cls requires fontspec .
Optional: pdftotext from poppler (macOS: brew install poppler , Debian/Ubuntu: apt install poppler-utils , Windows: choco install poppler ) — used by /apply 's ATS parseability check on the compiled CV. If missing, the check degrades gracefully to a visual keyword review.
gh repo fork MadsLorentzen/ai-job-search --clone
cd ai-job-search
2. Install job search tools
cd .agents/skills/jobbank-search/cli && bun install && cd ../../../..
cd .agents/skills/jobdanmark-search/cli && bun install && cd ../../../..
cd .agents/skills/jobindex-search/cli && bun install && cd ../../../..
cd .agents/skills/jobnet-search/cli && bun install && cd ../../../..
cd .agents/skills/linkedin-search/cli && bun install && cd ../../../..
For linkedin-search the install is optional: it has zero runtime dependencies and runs with plain bun ; bun install only pulls TypeScript dev types.
claude
# Then inside Claude Code:
/setup
/setup offers three paths: read your documents/ folder if you have one populated (CV PDF, LinkedIn export, diplomas, reference letters, past applications), import a single CV pasted in chat, or walk through an interview. It auto-detects what you have and asks. Documents-folder mode is idempotent and safe to re-run as you add more material; see documents/README.md for the layout.
/scrape
This searches multiple job portals for positions matching your profile, deduplicates results, and presents them sorted by fit. Pick a match to run /apply on it directly — or, when a scrape returns more jobs than you want to eyeball, run /rank to batch-score them all against the fit framework and get a ranked shortlist first.
/apply https://jobindex.dk/job/1234567
If the URL can't be fetched (some job portals block automated access), you can paste the job description directly instead:
/apply < paste the full job description here >
This runs the full workflow: evaluate fit, draft CV + cover letter, review with a second agent, revise, and present the final output.
/setup , /scrape , and /apply form the core workflow. Six more commands extend it once your profile is in place:
/outcome records what happened to an application - interview stages, offers, rejections, silence. It archives the submitted CV, cover letter, and posting text into documents/applications/<company>_<role>/ , keeps outcome.md in the format /setup Path A parses, and updates the tracker. Once a few applications resolve, it points you back to /setup to calibrate the fit framework from what actually got interviews.
/rank bridges /scrape and /apply : it batch-scores all newly scraped postings against the fit framework (parallel agents fetch each posting and score the five evaluation dimensions) and returns a ranked shortlist with honest per-job strengths and gaps. Deal-breakers veto, deadlines get urgency flags, dead postings get marked expired. Pick a number and it hands off to the full /apply workflow.
/expand enriches your profile by scanning public sources you've already linked in it (GitHub repos, portfolio site, Kaggle, Google Scholar) and looking up syllabi for named courses and certifications. Discovered competencies are added to your profile with a source tag. Useful right after /setup to surface skills that documents alone don't make explicit.
/upskill analyzes the gap between your profile and your tracked job postings (or a single posting via /upskill <URL> ). Produces a prioritized heatmap of skill gaps and a learning plan with web-searched study resources and time estimates. Useful for career planning between applications.
/add-template registers your own LaTeX CV or cover letter template in place of the stock ones. It captures the template's instructions (compile engine, fonts, style rules, page limit), runs a mandatory test compile, and wires the template into /apply . See LaTeX templates below.
/add-portal generates a job-portal search skill for a job board in your market. It investigates the portal (search URL pattern, result structure, access rules), scaffolds the CLI skill from the same structure as the shipped ones, and test-runs a live query before registering. See Job search tools below.
/reset is also available, see Starting over below.
ai-job-search/
├── CLAUDE.md # Main candidate profile + workflow rules
├── .claude/
│ ├── commands/
│ │ ├── apply.md # /apply workflow (drafter-reviewer)
│ │ ├── setup.md # /setup onboarding (documents folder, CV import, or interview)
│ │ ├── expand.md # /expand competency enrichment from documents and online presence
│ │ ├── add-template.md # /add-template register custom LaTeX templates
│ │ ├── add-portal.md # /add-portal generate a job-portal search skill for your market
│ │ ├── rank.md # /rank triage scraped jobs into a ranked shortlist
│ │ ├── outcome.md # /outcome record application results, archive materials
│ │ └── reset.md # /reset wipe profile data or documents folder
│ ├── skills/
│ │ ├── job-application-assistant/ # Core application skill
│ │ │ ├── SKILL.md # Skill definition
│ │ │ ├── 01-candidate-profile.md # Your education, experience, skills
│ │ │ ├── 02-behavioral-profile.md# PI/DISC/personality assessment
│ │ │ ├── 03-writing-style.md # Tone, structure, do's and don'ts
│ │ │ ├── 04-job-evaluation.md # Scoring framework for job fit
│ │ │ ├── 05-cv-templates.md # LaTeX CV structure + tailoring rules
│ │ │ ├── 06-cover-letter-templates.md # LaTeX cover letter templates
│ │ │ └── 07-interview-prep.md # STAR examples + interview framework
│ │ ├── job-scraper/ # Job search orchestration
│ │ └── upskill/ # /upskill skill gap analysis and learning plan
│ └── settings.json # Claude Code permissions (shared, scoped)
├── .agents/skills/ # Job portal CLI tools
│ ├── jobbank-search/ # Akademikernes Jobbank (Denmark)
│ ├── jobdanmark-search/ # Jobdanmark.dk (Denmark)
│ ├── jobindex-search/ # Jobindex.dk (Denmark)
│ ├── jobnet-search/ # Jobnet.dk (Denmark, government portal)
│ └── linkedin-search/ # LinkedIn public job listings (country-agnostic)
├── cv/
│ └── main_example.tex # moderncv LaTeX template
├── cover_letters/
│ ├── cover.cls # Custom cover letter LaTeX class
│ └── OpenFonts/ # Lato + Raleway fonts
├── templates/ # Custom templates registered via /add-
[truncated]
The /apply command runs a drafter-reviewer workflow with mandatory PDF compilation:
Parse the job posting (URL or text)
Evaluate fit against your profile (skills, experience, culture, location, career alignment)
Draft a tailored CV and cover letter in LaTeX
Spawn a reviewer agent that researches the company and critiques the drafts
Revise based on the reviewer's feedback
Compile and inspect both PDFs: lualatex for the CV, xelatex for the cover letter. Claude reads the rendered pages and iterates on the LaTeX until the CV is exactly 2 pages with no orphaned entry titles, and the cover letter is exactly 1 page with the signature visible and fonts consistent.
ATS-check the CV : extract the PDF's text layer ( pdftotext , optional dependency) and verify it the way an ATS parser sees it — contact details present as literal text, no garbled glyphs, sane reading order — then score the posting's keyword coverage against the extraction. Keywords the profile genuinely supports get added; genuine gaps stay visible, never stuffed.
Present the final output with a verification checklist
All claims in the CV and cover letter are verified against your actual profile. The system never fabricates skills or experience.
What makes this workflow different
PDF verification loop. Most LaTeX-resume templates produce "looks fine in the .tex" output that breaks in the PDF: job titles orphan to the next page, cover letters spill onto page 2, bullet fonts silently fall back to the body font. The /apply command compiles and visually inspects every PDF and applies targeted fixes ( \needspace , \enlargethispage , font-matching wrappers for list items) until the layout is clean. This

[truncated]
