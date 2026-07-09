---
source: "https://github.com/michaelwrites67-ctrl/yogen"
hn_url: "https://news.ycombinator.com/item?id=48845208"
title: "Show HN: Yogen – 500 AI agents argue about your idea before you bet on it"
article_title: "GitHub - michaelwrites67-ctrl/yogen: 予言 Yogen — a 500-agent AI swarm that debates your idea and predicts the outcome. Self-host free with your own Anthropic key. · GitHub"
author: "michaelruocco"
captured_at: "2026-07-09T13:44:15Z"
capture_tool: "hn-digest"
hn_id: 48845208
score: 2
comments: 0
posted_at: "2026-07-09T13:06:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Yogen – 500 AI agents argue about your idea before you bet on it

- HN: [48845208](https://news.ycombinator.com/item?id=48845208)
- Source: [github.com](https://github.com/michaelwrites67-ctrl/yogen)
- Score: 2
- Comments: 0
- Posted: 2026-07-09T13:06:02Z

## Translation

タイトル: Show HN: Yogen – 500 人の AI エージェントがあなたのアイデアに賭ける前に議論します
記事のタイトル: GitHub - michaelwrites67-ctrl/yogen: 予測 Yogen — あなたのアイデアを議論し、結果を予測する 500 人のエージェント AI の群れ。独自の Anthropic キーを使用したセルフホストは無料です。 · GitHub
説明: 予測 Yogen — あなたのアイデアを議論し、結果を予測する 500 人のエージェントからなる AI の群れ。独自の Anthropic キーを使用したセルフホストは無料です。 - michaelwrites67-ctrl/yogen

記事本文:
GitHub - michaelwrites67-ctrl/yogen: 予測 Yogen — あなたのアイデアを議論し、結果を予測する 500 人のエージェント AI の群れ。独自の Anthropic キーを使用したセルフホストは無料です。 · GitHub
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
michaelwrites67-ctrl
/
yogen
公共
通知
You must be s

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット データ data docs docs public public .env.example .env.example .gitignore .gitignore ライセンス ライセンス README.md README.md Agents.js Agents.js auth.js auth.js db.js db.js Hubspot.js Hubspot.js npm-shrinkwrap.json npm-shrinkwrap.json package.json package.json 鉄道.toml 鉄道.toml ルールエンジン.js ルールエンジン.js サーバー.js サーバー.js シミュレーション.js シミュレーション.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
予測 Yogen — あなたのアイデアについて議論し、結果を予測する 500 人のエージェントの群れ
スタートアップのアイデア、価格の変更、先延ばしにし続けている決定など、あなたの考えていることを Yogen に伝えてください。指名された 12 人の AI 専門家と 488 人のエージェント群衆が参加し、彼らはそれについて 6 ラウンドにわたって議論し、ソーシャル グラフを通じて意見が連鎖し、完全な書面による報告書による評決が得られます。
1 つの意見を与えるチャットボットではありません。ポジションを獲得する模擬討論。
ホストして試してみてください (3 つの無料シム): www.yogen.uk · 完全なデモ ビデオ
自分で実行 — 無制限、永久無料
git clone https://github.com/michaelwrites67-ctrl/yogen.git
CDヨーゲン
npmインストール
cp .env.example .env # Anthropic API キーを追加します
npm start # → http://localhost:3005
それだけです。セルフホスト型 Yogen にはシミュレーション制限がありません。料金は自分の Anthropic API の使用料のみです (500 エージェントの完全なシミュレーションの費用は約 $0.20 )。
要件: ノード 18 以降、Anthropic API キー。
あなたのアイデア
│
▼
┌─────────────────────────┐
│ 1. マーケットインテリジェンス + ナレッジグラフ抽出 │
│ (あなたのアイデア + 添付する書類から) │
§───────

─────────────────────┤
│ 2. あなたのドメイン用に生成された 12 人の指名された専門家 │
│ 4 肯定的 · 4 懐疑的 · 4 否定的 │
━━━━━━━━━━━━━━━━━━━━━━━━┤
│ 3. ソーシャル グラフ内の 488 人のバックグラウンド エージェント、│
│ 完璧な分割開始 33/33/33 │
━━━━━━━━━━━━━━━━━━━━━━━━┤
│ 4. 6 回の討論。専門家はこう主張する。スタンス │
│ 回答できない場合のみ変更します。 │
│ 意見の連鎖が群衆に波紋を広げます。 │
━━━━━━━━━━━━━━━━━━━━━━━━┤
│ 5. 評決 + 予測レポート: 最も強いシグナル、│
│ 確認すべき 1 つのこと、具体的な次のステップ — │
│ それぞれの作業はソロで、ブラウザ上で 30 分以内に実行できます。 │
━━━━━━━━━━━━━━━━━━━━━━━━┘
反論することもできます。討論の途中で「群衆に話す」と入力すると、エージェントの思い込みを修正し、証拠を追加できます。討論は次のラウンドでそれを吸収します。
すべてのシミュレーションは、Oracle の評決、最終的な群の分割、および完全な書面によるレポートで終了します。これにより、賛成派と反対派の最も強力なシグナル、最も重要な未知の要素、およびブラウザ上で 30 分以内にそれぞれ単独で実行できる次のステップが示されます。
正直さを保つメカニズム
バランスの取れたスタート。群れが始まります

実際に肯定的 33% / 懐疑的 33% / 否定的 33%。評決は議論によって得られるものであり、決して焼き付けられるものではありません。
スタンスの規律。エージェントは、ピアが答えられない点を見つけない限り、自分の立場を保持します。あら探しは洞察として扱われません。
有罪判決スコア。すべての評決には、群れがどれだけ強く収束したか、何人のエージェントが本当に考えを変えたかなど、質の高いシグナルが含まれているため、それを信頼できるかどうかがわかります。
報告書には捏造された事実はありません。エージェントはディベート中に即興でシナリオを作成します (ディベーターが行うのはこれです)。最終レポートは、実際に発言した内容に限定されており、何を行って検証すべきかを示します。
プレーンな Node.js + Express、ページごとに 1 つの HTML ファイル、フロントエンド フレームワークはありません。エージェントは Claude Haiku で実行されます (高速、安価)。クロード・ソネットに関するレポート総合。ファイルベースの JSON ストレージ — データベースをセットアップする必要はありません。
server.js ルート、認証、SSE ストリーミング、クォータ
Simulation.js ディベート エンジン — エージェントの生成、ラウンド、レポート
rules-engine.js 488-エージェントクラウド: ソーシャルグラフ、カスケード、連合
Agents.js エキスパートによるペルソナ生成
public/app.html シミュレーター UI
構成
API キーを除くすべてはオプションです。 .env.example を参照してください。パブリックにデプロイしてクォータが必要な場合は、UNLIMITED_SIMS=false および DAILY_SIM_LIMIT を設定します。 Stripe および Google-OAuth 変数はホスト型課金/ログイン用に存在しますが、無視できます。
AGPL-3.0 。自由に使用し、フォークし、そこから学び、自己ホストします。変更したバージョンをサービスとして実行する場合は、変更をオープンソースにする必要があります。 yogen.uk のホスト型サービスは Runn Ltd によって運営されています。
予測 Yogen — あなたのアイデアについて議論し、結果を予測する 500 人のエージェントからなる AI の群れ。独自の Anthropic キーを使用したセルフホストは無料です。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。リロしてください

このページに広告を掲載します。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

予言 Yogen — a 500-agent AI swarm that debates your idea and predicts the outcome. Self-host free with your own Anthropic key. - michaelwrites67-ctrl/yogen

GitHub - michaelwrites67-ctrl/yogen: 予言 Yogen — a 500-agent AI swarm that debates your idea and predicts the outcome. Self-host free with your own Anthropic key. · GitHub
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
michaelwrites67-ctrl
/
yogen
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit data data docs docs public public .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md agents.js agents.js auth.js auth.js db.js db.js hubspot.js hubspot.js npm-shrinkwrap.json npm-shrinkwrap.json package.json package.json railway.toml railway.toml rule-engine.js rule-engine.js server.js server.js simulation.js simulation.js View all files Repository files navigation
予言 Yogen — a 500-agent swarm that debates your idea and predicts the outcome
Tell Yogen what's on your mind — a startup idea, a price change, a decision you keep putting off. It spins up 12 named AI experts and a 488-agent crowd , they argue about it for six rounds, opinions cascade through a social graph, and you get a verdict with a full written report.
Not a chatbot giving you one opinion. A simulated debate where positions are earned.
Try it hosted (3 free sims): www.yogen.uk · Full demo video
Run it yourself — unlimited, free forever
git clone https://github.com/michaelwrites67-ctrl/yogen.git
cd yogen
npm install
cp .env.example .env # add your Anthropic API key
npm start # → http://localhost:3005
That's it. Self-hosted Yogen has no simulation limits — you pay only your own Anthropic API usage (a full 500-agent sim costs roughly $0.20 ).
Requirements: Node 18+, an Anthropic API key .
your idea
│
▼
┌─────────────────────────────────────────────────────┐
│ 1. Market intel + knowledge-graph extraction │
│ (from your idea + any document you attach) │
├─────────────────────────────────────────────────────┤
│ 2. 12 named experts generated for YOUR domain │
│ 4 positive · 4 sceptical · 4 negative │
├─────────────────────────────────────────────────────┤
│ 3. 488 background agents in a social graph, │
│ starting perfectly split 33/33/33 │
├─────────────────────────────────────────────────────┤
│ 4. Six rounds of debate. Experts argue; stances │
│ change only when a point can't be answered. │
│ Opinion cascades ripple through the crowd. │
├─────────────────────────────────────────────────────┤
│ 5. Verdict + prediction report: strongest signals, │
│ the one thing to check, concrete next steps — │
│ each doable solo, in a browser, under 30 min. │
└─────────────────────────────────────────────────────┘
You can argue back. Mid-debate, type into "Speak to the swarm" — correct an agent's assumption, add evidence — and the debate absorbs it in the next round.
Every simulation ends with an Oracle Verdict, the final swarm split, and a full written report — strongest signals for and against, the single most important unknown, and next steps you can each do solo, in a browser, in under 30 minutes.
The mechanics that keep it honest
Balanced start. The swarm begins exactly 33% positive / 33% sceptical / 33% negative. The verdict must be earned by the debate, never baked in.
Stance discipline. Agents hold their position unless a peer lands a point they cannot answer. Finding fault is not treated as insight.
Conviction score. Every verdict ships with a quality signal — how strongly the swarm converged and how many agents genuinely changed their mind — so you know whether to trust it.
No invented facts in reports. Agents improvise scenarios while debating (that's what debaters do); the final report is constrained to what you actually said, and tells you what to go and verify.
Plain Node.js + Express, one HTML file per page, zero front-end framework. Agents run on Claude Haiku (fast, cheap); report synthesis on Claude Sonnet. File-backed JSON storage — no database to set up.
server.js routes, auth, SSE streaming, quotas
simulation.js the debate engine — agent generation, rounds, report
rule-engine.js 488-agent crowd: social graph, cascades, coalitions
agents.js expert persona generation
public/app.html the simulator UI
Configuration
Everything is optional except the API key — see .env.example . Set UNLIMITED_SIMS=false and DAILY_SIM_LIMIT if you deploy it publicly and want quotas; Stripe and Google-OAuth vars exist for hosted billing/login and can be ignored.
AGPL-3.0 . Use it, fork it, learn from it, self-host it freely. If you run a modified version as a service, you must open-source your modifications. The hosted service at yogen.uk is operated by Runn Ltd.
予言 Yogen — a 500-agent AI swarm that debates your idea and predicts the outcome. Self-host free with your own Anthropic key.
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
