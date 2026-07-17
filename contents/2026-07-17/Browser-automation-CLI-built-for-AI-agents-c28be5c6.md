---
source: "https://github.com/browser-act/skills"
hn_url: "https://news.ycombinator.com/item?id=48946421"
title: "Browser automation CLI built for AI agents"
article_title: "GitHub - browser-act/skills: Browser automation CLI built for AI agents. Break through anti-bot walls, hand off to humans across platforms when stuck. Parallel multi-task execution, independent multi-session operation, isolated multi-account browsing. · GitHub"
author: "aanthonymax"
captured_at: "2026-07-17T13:05:43Z"
capture_tool: "hn-digest"
hn_id: 48946421
score: 1
comments: 0
posted_at: "2026-07-17T12:14:02Z"
tags:
  - hacker-news
  - translated
---

# Browser automation CLI built for AI agents

- HN: [48946421](https://news.ycombinator.com/item?id=48946421)
- Source: [github.com](https://github.com/browser-act/skills)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T12:14:02Z

## Translation

タイトル: AI エージェント用に構築されたブラウザ自動化 CLI
記事のタイトル: GitHub - ブラウザー行為/スキル: AI エージェント用に構築されたブラウザー自動化 CLI。ボット対策の壁を突破し、立ち往生した場合はプラットフォームを越えて人間に引き渡します。マルチタスクの並列実行、独立したマルチセッション操作、分離されたマルチアカウントのブラウジング。 · GitHub
説明: AI エージェント用に構築されたブラウザ自動化 CLI。ボット対策の壁を突破し、立ち往生した場合はプラットフォームを越えて人間に引き渡します。マルチタスクの並列実行、独立したマルチセッション操作、分離されたマルチアカウントのブラウジング。 - ブラウザアクト/スキル

記事本文:
GitHub - browser-act/skills: AI エージェント用に構築されたブラウザ自動化 CLI。ボット対策の壁を突破し、立ち往生した場合はプラットフォームを越えて人間に引き渡します。マルチタスクの並列実行、独立したマルチセッション操作、分離されたマルチアカウントのブラウジング。 · GitHub
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
却下する

警告
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ブラウザーアクト
/
スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
364 コミット 364 コミット .github/ workflows .github/ workflows ブラウザ-act-skill-forge ブラウザ-act-skill-forge ブラウザ-アクト ドキュメント ドキュメント ソリューション ソリューション .gitignore .gitignore ライセンス ライセンス README.md README.md 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用に構築されたブラウザ自動化 CLI。ボット対策の壁を乗り越え、行き詰まったときにプラットフォームをまたがって人間に引き継ぎ、相互汚染することなく並行タスクを実行し、独立したブラウザーで複数のアカウントを隔離します。
AI エージェントに必要なブラウザは、標準ツールがアクセスできない場所に到達できること、エージェントが行き詰まったときに人間がシームレスに引き継げること、並列タスクの相互汚染を防ぐこと、そして人間が作成したスクリプトではなく LLM 推論用に設計されている必要があります。エージェント用のブラウザは 4 つのことを正しく行う必要があります。
1. ブロックを突破する - 3 つのプログレッシブ レイヤー
環境層 - ステルス指紋スプーフィング、TLS ローテーション、プロキシ切り替え。ブロックの大部分は決してトリガーされません。
実行層 —solve-captcha は CAPTCHA を自動解決します。ステルス抽出は、保護されたページを 1 つのコマンドでプルします。
人間層 — リモートアシストはライブ URL を生成します。ユーザーは任意のデバイスから引き継ぎ、エージェントは完了するとシームレスに続行します。
2. 3 つのブラウザ モード — 現実世界のシナリオによる
3. 干渉ゼロの同時実行 - すべてのエージェントが独自のレーンに存在
クロスブラウザーの並列 — 独立した Cookie、フィンガープリント、プロキシ。サイトはそれらを関連付けることができません。
同一ブラウザのマルチセッション — 共有ログイン状態、独立

インデント実行では、タスクは互いにブロックしません。
プライバシー モード — セッションごとに新しい指紋と空のプロファイルが作成され、完了すると残留物がゼロになります。
4. 人間のスクリプトではなく、エージェントの推論のために設計されています
コンパクトなテキスト出力 - インデックス付きテキスト形式で、JSON や HTML よりも数倍トークン効率が優れています。
インデックス付きインタラクション - 状態はインデックス付きリストを返します。 3をクリック / 2「...」と入力します。 DOM 解析は必要ありません。
セマンティック メモリ — すべてのブラウザには、意味によってタスクと一致する desc が含まれています。
同時実行安全 — セッション所有権 + 明示的な名前付け。マルチエージェントの操作が競合することはありません。
セキュリティ: 確認ゲート - 機密性の高い操作 (ブラウザの作成/削除、プロファイルのインポート、プロキシの変更、セキュリティとプライバシーの切り替え) には明示的なユーザーの承認が必要です。事前の承認は引き継がれません。設定の切り替えではなく、スキル層で適用されます。
より良いヘッドレス — ユーザーの邪魔をしないデフォルトのヘッドレス。検出されないステルスヘッドレス。
クロスプラットフォーム リモート ハンドオフ — どのデバイスでもリンクを開いて引き継ぎ、エージェントはシームレスに続行します。
ブラウザーアクトをインストールします。スキルのソース: https://github.com/browser-act/skills/tree/main/browser-act 。インストール後に動作を確認します。
# 保護されたページのコンテンツを抽出します (構成なし)
ブラウザアクト ステルス抽出 https://example.com
# ブラウザーの完全な自動化
browser-act --session my-task ブラウザを開く < id > https://example.com
browser-act --session my-task state # クリック可能な要素を参照
browser-act --session my-task click 3 # インデックスでクリック
browser-act --session my-task input 2 " hi " # フィールドに入力します
その他の例とワークフロー →
エージェントは各セッションの開始時に get-skill を実行します。つまり、1 回の呼び出しで環境の状態、ブラウザーのリスト、およびコマンドを取得します。
browser-act get-skills core --skill-version 2.0.2
エージェントが BrowserAct を検出して使用する方法

→
エージェント: Claude Code、Cursor、VS Code、OpenCode、OpenClaw、Codex、Gemini CLI — シェル コマンドを実行し、スキルをロードできる任意のエージェントと連携します。
ほぼすべてが無料です。支払いが必要な機能は、マネージド プロキシ (動的/静的) と最初の 5 つを超えるステルス ブラウザの 2 つだけです。
完全なドキュメントには、アンチブロッキング、ブラウザ モード、セッションと同時実行、ヘッドレス ハンドオフとリモート ハンドオフ、エージェントの設計、スキル システム、および完全なコマンド リファレンスが含まれています。
Skill Forge — あなたの専属スクレイピング エンジニア
同じ Web サイトから大規模にデータを繰り返し抽出する必要がありますか?スクレイパーを手書きしないでください。 Skill Forge は、サイトを一度探索して、その API とデータ パターンを検出し、すぐにデプロイできるスキル パッケージを生成します。その後、再探索することなく確実に実行されます (同じ安定したパスを通じて 500 または 5,000 のレコードが実行されます)。
あらゆるウェブサイト。あらゆるデータ。開始する 1 つのコマンド:
ブラウザ-act-skill-forgeをインストールします。スキルのソース: https://github.com/browser-act/skills/tree/main/browser-act-skill-forge 。インストール後に動作を確認します。
次に、エージェントに必要なことを伝えます。
「LinkedIn から求人情報 (役職、会社、給与、URL) を抽出するスキルを作成します。後で 300 個のキーワードを実行します。」
Skill Forge によってすでに生成された 30 以上のビルド済みスキルがあり、すぐにインストールして実行できます。 Amazon、Google マップ、YouTube、Reddit、WeChat、Zhihu などをカバーします。
完全なソリューション カタログを参照 →
上記で必要なものが見つかりませんか?あらゆる Web サイトのカスタム スキルを数分で生成できます。コーディングは必要ありません。必要なデータや実行するアクションを記述するだけで、残りは Skill Forge が処理します。
BrowserAct Skills は無料のオープンソースです。時間が節約できる場合は、⭐ スターを付けてください。これはプロジェクトを存続させ、より多くのスキルを提供するのに役立ちます。
🎁 ボーナス: リポジトリにスターを付けると、ディスクに参加できるようになります

#claim-500-credits チャンネルに投稿して、500 クレジットを無料で受け取りましょう!
BrowserAct チームによって ❤️ を使用して構築されました
AI エージェント用に構築されたブラウザ自動化 CLI。ボット対策の壁を突破し、立ち往生した場合はプラットフォームを越えて人間に引き渡します。マルチタスクの並列実行、独立したマルチセッション操作、分離されたマルチアカウントのブラウジング。
www.browseract.com/?co-from=github
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
211
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Browser automation CLI built for AI agents. Break through anti-bot walls, hand off to humans across platforms when stuck. Parallel multi-task execution, independent multi-session operation, isolated multi-account browsing. - browser-act/skills

GitHub - browser-act/skills: Browser automation CLI built for AI agents. Break through anti-bot walls, hand off to humans across platforms when stuck. Parallel multi-task execution, independent multi-session operation, isolated multi-account browsing. · GitHub
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
browser-act
/
skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
364 Commits 364 Commits .github/ workflows .github/ workflows browser-act-skill-forge browser-act-skill-forge browser-act browser-act docs docs solutions solutions .gitignore .gitignore LICENSE LICENSE README.md README.md requirements.txt requirements.txt View all files Repository files navigation
Browser automation CLI built for AI agents. Get past anti-bot walls, hand off to humans across platforms when stuck, run parallel tasks without cross-contamination, and isolate multiple accounts in independent browsers.
The browser an AI agent needs has to reach places standard tools can't, let a human seamlessly take over when the agent is stuck, keep parallel tasks from cross-contaminating, and be designed for LLM reasoning — not human-written scripts. A browser for agents must get four things right.
1. Break through blocks — three progressive layers
Environment layer — stealth fingerprint spoofing, TLS rotation, proxy switching. The vast majority of blocks never trigger.
Execution layer — solve-captcha auto-solves CAPTCHAs; stealth-extract pulls protected pages in one command.
Human layer — remote-assist generates a live URL; the user takes over from any device, and the agent continues seamlessly when done.
2. Three browser modes — by real-world scenario
3. Zero-interference concurrency — every agent in its own lane
Cross-browser parallel — independent cookies, fingerprints, proxies. Sites cannot correlate them.
Same-browser multi-session — shared login state, independent execution, tasks don't block each other.
Privacy mode — fresh fingerprint and empty profile per session, zero residue when done.
4. Designed for agent reasoning — not human scripts
Compact text output — indexed text format, several times more token-efficient than JSON or HTML.
Indexed interaction — state returns an indexed list; click 3 / input 2 "..." . No DOM parsing required.
Semantic memory — every browser carries a desc , matched to tasks by meaning.
Concurrency-safe — session ownership + explicit naming. Multi-agent operation never conflicts.
Security: confirmation gating — sensitive operations (browser create / delete, Profile import, proxy changes, security and privacy toggles) require explicit user approval. Prior approvals do not carry over. Enforced at the Skill layer, not a configuration toggle.
Better headless — Default headless without disrupting users; stealth headless that isn't detected.
Cross-platform remote handoff — Any device opens the link to take over, and the agent continues seamlessly.
Install browser-act. Skill source: https://github.com/browser-act/skills/tree/main/browser-act . Verify it works after installation.
# Extract protected page content (zero config)
browser-act stealth-extract https://example.com
# Full browser automation
browser-act --session my-task browser open < id > https://example.com
browser-act --session my-task state # See clickable elements
browser-act --session my-task click 3 # Click by index
browser-act --session my-task input 2 " hi " # Type into a field
More examples and workflows →
The agent runs get-skills at the start of each session — gets environment state, browser list, and commands in one call:
browser-act get-skills core --skill-version 2.0.2
How agents discover and use BrowserAct →
Agents: Claude Code · Cursor · VS Code · OpenCode · OpenClaw · Codex · Gemini CLI — works with any agent that can execute shell commands and load Skills.
Almost everything is free. Only two features require payment: managed proxies (Dynamic / Static), and stealth browsers beyond the first 5.
Full documentation covers anti-blocking, browser modes, sessions and concurrency, headless and remote handoff, agent design, the Skills system, and the complete command reference.
Skill Forge — Your Personal Scraping Engineer
Need to extract data from the same website repeatedly at scale? Don't write scrapers by hand. Skill Forge explores a site once, discovers its APIs and data patterns, generates a deploy-ready Skill package, then runs reliably without re-exploration — 500 or 5,000 records through the same stable path.
Any website. Any data. One command to start:
Install browser-act-skill-forge. Skill source: https://github.com/browser-act/skills/tree/main/browser-act-skill-forge . Verify it works after installation.
Then tell your agent what you need:
"Forge a Skill that extracts job listings from LinkedIn — title, company, salary, URL. I'll run 300 keywords later."
30+ pre-built Skills already generated by Skill Forge, ready to install and run. Covers Amazon, Google Maps, YouTube, Reddit, WeChat, Zhihu, and more.
Browse the full Solutions Catalog →
Can't find what you need above? Generate a custom Skill for any website in minutes — no coding required. Just describe what data you want or what action to perform, and Skill Forge handles the rest.
BrowserAct Skills is free and open source . If it saves you time, please give us a ⭐ Star — it keeps the project alive and helps us ship more skills.
🎁 Bonus: Once you star the repository, you can join our Discord and post in the #claim-500-credits channel to receive 500 free credits !
Built with ❤️ by the BrowserAct Team
Browser automation CLI built for AI agents. Break through anti-bot walls, hand off to humans across platforms when stuck. Parallel multi-task execution, independent multi-session operation, isolated multi-account browsing.
www.browseract.com/?co-from=github
Topics
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
211
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
