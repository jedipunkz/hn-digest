---
source: "https://github.com/inklate/social-skills"
hn_url: "https://news.ycombinator.com/item?id=48957301"
title: "Show HN: Open-source skills that make any AI agent write native social posts"
article_title: "GitHub - inklate/social-skills: Social media skills for your AI agents · GitHub"
author: "missingstack"
captured_at: "2026-07-18T12:53:22Z"
capture_tool: "hn-digest"
hn_id: 48957301
score: 2
comments: 0
posted_at: "2026-07-18T12:03:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source skills that make any AI agent write native social posts

- HN: [48957301](https://news.ycombinator.com/item?id=48957301)
- Source: [github.com](https://github.com/inklate/social-skills)
- Score: 2
- Comments: 0
- Posted: 2026-07-18T12:03:50Z

## Translation

タイトル: HN を表示: AI エージェントにネイティブのソーシャル投稿を作成させるオープンソース スキル
記事のタイトル: GitHub - inklate/social-skills: AI エージェントのためのソーシャル メディア スキル · GitHub
説明: AI エージェント向けのソーシャル メディア スキル。 GitHub でアカウントを作成して、インクレート/ソーシャル スキルの開発に貢献します。

記事本文:
GitHub - inklate/social-skills: AI エージェントのためのソーシャル メディア スキル · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
インクレート
/
ソーシャルスキル
公共
通知
通知セットを変更するにはサインインする必要があります

ひずみ
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .changeset .changeset .claude-plugin .claude-plugin .github/ workflows .github/ workflows アセット アセット スクリプト スクリプト スキル スキル .gitignore .gitignore .prettierignore .prettierignore .prettierrc .prettierrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md RELEASING.md RELEASING.md bun.lock bun.lock package.json package.json repo.config.json repo.config.json skill-index.json skill-index.json すべてのファイルを表示 リポジトリ ファイルナビゲーション
AI エージェントのためのソーシャル メディア スキル
戦略、声、プラットフォームネイティブの形式 (投稿、スレッド、カルーセル、リール対応スクリプト、広告コピー、カレンダー) を学習する AI エージェントのスキルのコレクション。
エージェントはすでに書き込みを行うことができます。スクロールを止めるフックを備えた、フィードに合わせたサイズの LinkedIn 投稿を、あなたらしく書くことはできません。これらのスキルは、戦略、音声、プラットフォーム ネイティブの形式 (投稿、スレッド、カルーセル、リール対応スクリプト、広告コピー、カレンダー) を、1 回限りのプロンプトではなく完全なコンテンツ システムとして教えます。どのエージェントでもアカウントや API は機能しません。
npx スキルは、inklate/ソーシャル スキルを追加します
ヒント: エージェントがこれを非対話的に実行すると、 .agents/skills/ にのみインストールされる可能性がありますが、クロード コードはこれを読み取れません ( pass -a claude-code )。 --skill Cross-post を使用して単一のスキルをインストールします。
または、自動更新の Claude Code プラグインとして:
/プラグインマーケットプレイスにinklate/ソーシャルスキルを追加
/plugin install social-skills@inklate
次に、「ソーシャル コンテキストを設定してください」と言って開始します。ソーシャル コンテキスト スキルは、あなたに 1 回インタビューし、他のすべてのスキルが視聴者、プラットフォーム、柱、声のために読み取るファイル social-context.md を作成します。
スキー

ll
カテゴリ
何をするのか
社会的文脈
財団
ユーザーに 10 個以下の質問でインタビューし、ユーザーが誰なのか、誰と話しているのか、何について投稿しているのか、何を投稿を拒否しているのかを記録した social-context.md ファイルを作成します。
声
財団
ユーザーが提供する 3 ～ 10 個の文章サンプル (LinkedIn 投稿、X (Twitter) スレッド、電子メール、ブログ投稿) を分析し、文の長さ、リズム、語彙、句読点と絵文字の習慣、オープニング、サインオフ、ユーモアの表現など、ユーザーが実際にどのように書いているかを抽出します。
カルーセル
作成
投稿、記事、アイデアを、フック カバー、スライドごとに 1 つのアイデアのコンテンツ スライド、要約、CTA スライドを備えた、Instagram または LinkedIn 用の完全なスライドごとのカルーセル スクリプトに変換します。
クロスポスト
作成
1 つのストーリー、ブログ投稿、お知らせ、またはアイデアを、複数のプラットフォーム (LinkedIn、X (Twitter)、Instagram、Facebook、Threads、Bluesky、Mastodon) 向けのネイティブ感覚の下書きに一度に変換します。トーン、長さ、構造、ハッシュタグ、フォーマットをプラットフォームごとに調整します。1 つのキャプションをどこにでもコピーアンドペーストするのではありません。
フック
作成
ソーシャル投稿のドラフトやトピックに対して 8 つのスコア付きフック オプションを生成し、好奇心のギャップや大胆な主張から、逆張りの解釈やストーリーのコールド オープンまで、実証済みのパターンを網羅します。
リンクされた投稿
作成
LinkedIn フィードが実際にどのように機能するかを考慮して作成された LinkedIn 投稿の下書きを作成します。最初の 200 文字で「…もっと見る」クリックを獲得するフック、空白のリズム、投稿ごとに 1 つのアイデア、特定の行動喚起、ハッシュタグ スープはありません。
返信ライター
作成
LinkedIn、X、Instagram、Facebook、またはスレッドから貼り付けられたコメント、メンション、DM のバッチを優先順位付けし、必要に応じて音声返信の下書きを作成します。
再利用する
作成
ブログ投稿、ビデオ トランスクリプト、ポッドキャスト エピソード、または講演など、1 つの長い形式のソースから最も強力なアトミックなアイデアを見つけて提出します。

LinkedIn、X (Twitter)、Instagram のプラットフォーム形式のドラフトを丸 1 週間にわたって作成し、さらに短いビデオ スクリプトを作成します。
Xスレッド
作成
断片的なブログ投稿ではなく、実際のアーキテクチャを使用して X (Twitter) スレッドを構築します。
[切り捨てられた]
┌───────┐ ┌───────┐
│ ソーシャルコンテキスト │ + │ 音声 │ 基盤 — これらを最初に実行します。
└───────┬───┘ └───┬───┘ 彼らは social-context.md を書きます
━───────┬───┘
┌───────────┼───────────┐
▼ ▼ ▼
計画の作成チェック
linkedin-投稿コンテンツ-カレンダー事後チェック
x-スレッドのソーシャル監査
フックADS
カルーセル広告ライター
クロスポスト
再利用する
返信ライター
ただ言ってください
スキルは自然なフレーズで発動します。覚えるコマンドはありません。
「このブログ投稿を LinkedIn と X の投稿に変える」 → クロスポスト
「当社の立ち上げについて LinkedIn 投稿を書く」 → linkedin-post
「このドラフト用のより良いフックを教えてください」 → フック
「今後 2 週間のコンテンツを計画する」 → content-calendar
「これはインスタグラムでも使えますか？」 →事後チェック
「このオファーの Facebook 広告バリエーションを作成します」 → 広告ライター
「これらのコメントに返信するのを手伝ってください」 → Reply-writer
「これらの投稿から私の声を学びましょう」 → 声
これらは標準のエージェント スキル (スキルごとに 1 つのフォルダー、それぞれが SKILL.md) であるため、この形式を読み取るエージェントはどれもそれらを使用できます。スキル CLI は、使用するエージェントにスキルをインストールします (40 以上をサポート)。 pass -a <agent> を指定して明示的にターゲットにするか、省略してインストールしたエージェントを自動検出します。
完全なエージェントのリストは skill.sh にあります。
所定の場所にコピーします: git clone https://github.com/inklate/social-skills && cp -r social-skills/skills/* ~/.claude/skills/ (swa

p エージェントのスキル ディレクトリのターゲット)。
サブモジュール: git submodule add https://github.com/inklate/social-skills .agents/social-skills
ZIP: クロード デスクトップ/Cowork アップロード用の GitHub リリース上のスキルごとの zip。
フォークしてください。これらのスキルはカスタマイズすることを目的としており、リポジトリをフォークし、ブランドに合わせて音声ルール、プラットフォームのデフォルト、品質バーを編集します。
これらのスキルにより草案や計画が作成されます。次に登場するのは、エージェントが LinkedIn、X、Instagram、Facebook 全体で実際に公開、スケジュール、測定できるようにするコンパニオン スキル パックです。この組織を見て、いつ着地するかを聞いてください。
新しい一般的なスキル、プラットフォーム制約の更新 (文字制限は常に変化します)、および音声/形式の改善はすべて歓迎されます。 CONTRIBUTING.md を参照してください。 1 つの厳しいルール: ここでのスキルはベンダー中立性を保ち、エージェント自身の能力のみで完全な約束を実現します。
MITライセンス。 Inklate によって構築されました。
AI エージェントのためのソーシャル メディア スキル
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
1
v0.0.1
最新の
2026 年 7 月 18 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Social media skills for your AI agents. Contribute to inklate/social-skills development by creating an account on GitHub.

GitHub - inklate/social-skills: Social media skills for your AI agents · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
inklate
/
social-skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .changeset .changeset .claude-plugin .claude-plugin .github/ workflows .github/ workflows assets assets scripts scripts skills skills .gitignore .gitignore .prettierignore .prettierignore .prettierrc .prettierrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md bun.lock bun.lock package.json package.json repo.config.json repo.config.json skills-index.json skills-index.json View all files Repository files navigation
Social media skills for AI Agents
A collection of AI agent skills that learn your strategy, voice, and platform-native format — posts, threads, carousels, reels-ready scripts, ad copy, calendars.
Your agent can already write. It can't write a LinkedIn post that sounds like you, sized for the feed, with a hook that stops the scroll. These skills teach it strategy, voice, and platform-native format — posts, threads, carousels, reels-ready scripts, ad copy, calendars — a full content system, not one-off prompts. No accounts, no APIs, works in every agent.
npx skills add inklate/social-skills
Tip: when an agent runs this non-interactively it may install only to .agents/skills/ , which Claude Code doesn't read — pass -a claude-code . Install a single skill with --skill cross-post .
Or as an auto-updating Claude Code plugin:
/plugin marketplace add inklate/social-skills
/plugin install social-skills@inklate
Then say "set up my social context" to begin — the social-context skill interviews you once and writes social-context.md , the file every other skill reads for your audience, platforms, pillars, and voice.
Skill
Category
What it does
social-context
Foundation
Interview the user in ten questions or fewer and produce a social-context.md file that captures who they are, who they are talking to, what they post about, and what they refuse to post.
voice
Foundation
Analyze 3–10 writing samples the user provides — LinkedIn posts, X (Twitter) threads, emails, blog posts — and distill how they actually write: sentence length, rhythm, vocabulary, punctuation and emoji habits, openers, sign-offs, and humor register.
carousel
Create
Turn a post, article, or idea into a complete slide-by-slide carousel script for Instagram or LinkedIn, with a hook cover, one-idea-per-slide content slides, a recap, and a CTA slide.
cross-post
Create
Turn one story, blog post, announcement, or idea into native-feeling drafts for multiple platforms at once — LinkedIn, X (Twitter), Instagram, Facebook, Threads, Bluesky, and Mastodon — adapting tone, length, structure, hashtags, and format per platform instead of copy-pasting one caption everywhere.
hooks
Create
Generate eight scored hook options for any social post draft or topic, spanning proven patterns from curiosity gaps and bold claims to contrarian takes and story cold-opens.
linkedin-post
Create
Draft a LinkedIn post built for how the LinkedIn feed actually works: a hook that earns the "…see more" click in the first 200 characters, white-space rhythm, one idea per post, a specific call to action, and no hashtag soup.
reply-writer
Create
Triage a pasted batch of comments, mentions, and DMs from LinkedIn, X, Instagram, Facebook, or Threads, and draft on-voice replies where they're warranted.
repurpose
Create
Mine one long-form source — a blog post, video transcript, podcast episode, or talk — for its strongest atomic ideas and turn it into a full week of platform-shaped drafts for LinkedIn, X (Twitter), and Instagram, plus a short-video script.
x-thread
Create
Build an X (Twitter) thread with real architecture instead of a chopped-up blog post: a
[truncated]
┌────────────────┐ ┌───────┐
│ social-context │ + │ voice │ Foundation — run these first;
└───────┬────────┘ └───┬───┘ they write social-context.md
└───────┬────────┘
┌───────────────────┼────────────────────┐
▼ ▼ ▼
CREATE PLAN CHECK
linkedin-post content-calendar post-check
x-thread social-audit
hooks ADS
carousel ad-writer
cross-post
repurpose
reply-writer
Just say it
Skills trigger on natural phrases — no commands to memorize:
"turn this blog post into posts for LinkedIn and X" → cross-post
"write a LinkedIn post about our launch" → linkedin-post
"give me better hooks for this draft" → hooks
"plan my next two weeks of content" → content-calendar
"will this work on Instagram?" → post-check
"write Facebook ad variants for this offer" → ad-writer
"help me reply to these comments" → reply-writer
"learn my voice from these posts" → voice
These are standard Agent Skills — one folder per skill, each a SKILL.md — so any agent that reads the format can use them. The skills CLI installs them into whichever agent you use (40+ supported); pass -a <agent> to target one explicitly, or omit it to auto-detect the agents you have installed.
The full agent list lives at skills.sh .
Copy into place: git clone https://github.com/inklate/social-skills && cp -r social-skills/skills/* ~/.claude/skills/ (swap the target for your agent's skills directory).
Submodule: git submodule add https://github.com/inklate/social-skills .agents/social-skills
Zips: per-skill zips on GitHub Releases for Claude Desktop / Cowork upload.
Fork it. These skills are meant to be customized — fork the repo and edit the voice rules, platform defaults, and quality bars to match your brand.
These skills produce drafts and plans. Coming next: a companion skill pack that lets your agent actually publish, schedule, and measure across LinkedIn, X, Instagram, and Facebook. Watch this org to hear when it lands.
New general skills, platform-constraint updates (char limits change constantly), and voice/format improvements are all welcome — see CONTRIBUTING.md . One hard rule: skills here stay vendor-neutral and deliver their complete promise with only the agent's own abilities.
MIT LICENSE . Built by Inklate .
Social media skills for your AI agents
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
v0.0.1
Latest
Jul 18, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
