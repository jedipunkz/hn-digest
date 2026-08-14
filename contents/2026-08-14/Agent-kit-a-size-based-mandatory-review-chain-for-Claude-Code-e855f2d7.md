---
source: "https://github.com/horilla98/agent-kit"
hn_url: "https://news.ycombinator.com/item?id=49296844"
title: "Agent-kit – a size-based mandatory review chain for Claude Code"
article_title: "GitHub - horilla98/agent-kit · GitHub"
author: "danielhorilla"
captured_at: "2026-08-14T10:55:39Z"
capture_tool: "hn-digest"
hn_id: 49296844
score: 1
comments: 0
posted_at: "2026-08-14T10:28:43Z"
tags:
  - hacker-news
  - translated
---

# Agent-kit – a size-based mandatory review chain for Claude Code

- HN: [49296844](https://news.ycombinator.com/item?id=49296844)
- Source: [github.com](https://github.com/horilla98/agent-kit)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T10:28:43Z

## Translation

タイトル: Agent-kit – クロード コードのサイズベースの必須レビュー チェーン
記事タイトル: GitHub - horilla98/agent-kit · GitHub
説明: GitHub でアカウントを作成して、horilla98/agent-kit の開発に貢献します。

記事本文:
GitHub - horilla98/agent-kit · GitHub
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
ホリラ98
/
エージェントキット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .claude/ エージェント .claude/ エージェント .github .github ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Agent-kit (無料コア) — クロード コード用のサイズベースのワークフロー ゲート
これは、エージェント キットの無料のオープンソース コアです: eng

サイズベースの必須の背後にあるもの
レビューチェーン。チケットのサイズと、開始時に必要なチェーンステップを思い出させます。
クロード コードのセッションごとに、チェーン メンバーが持つ PR をボット コメント経由で追跡します。
実際に痕跡を残しました。
これは完全なエージェント キットではありません。フルパッケージ (11 のエージェントチャーター、7 つのスラッシュコマンド、
完全な 13 ステップの開発ライフサイクル、バイリンガル テンプレート ツリー、インストーラーは有料です。以下を参照してください。
ここに含まれるもの (エージェントキット v2.0 ソースに対して検証済み)
8 つのスクリプト ( .github/scripts/ ): lanc-tabla 、 munkarend-kapu + munkarend-kapu-logic 、
lanc-nyom + lanc-nyom-logic 、gh-api 、projekt-config 、nyelv-cimkek 、uzenetek
2 つの GitHub Actions ワークフロー ( .github/workflows/ ): ci.yml 、 pr-jelzok.yml
11 人のエージェントのうち 3 人のエージェントがチャーターします: reka (コードレビュー担当者)、tibor (テストエンジニア)、
gergo (セキュリティ + プライバシー、拒否権付き) + 共有の _protokoll.md
45 の自動テスト、それ自体はグリーン — 独自のテストで検証された公開/カット
有料部分に依存しないスイート
npm 依存関係なし、Node.js ノードのみ: 組み込み (Node.js ≥ 22.6)
含まれないもの: 他の 8 つのエージェント チャーター (sara、bence、marci、izsak、zsofi、petra、devops、
columbo)、7 つのスラッシュ コマンド、バックログのインポート ワークフロー、インストーラー スクリプト、
バイリンガル テンプレート ツリー、ドキュメント スケルトン (ADR、CONTRIBUTING、ワークフロー ドキュメント)。
.github/scripts/ 、 .github/workflows/ 、および .claude/agents/ の内容を
独自のリポジトリを作成し、リポジトリのルートにagent-kit.config.jsonを追加します。
{
"projekt" : " プロジェクトの名前 " ,
"repo" : " your-github-username/your-repo " ,
"nyelv" : " 英語 " ,
"teszt" : " npm テスト "
}
手動の手順が 1 つ必要です。有料パッケージでは、インストーラー スクリプトが
エージェントチャーター内のプロジェクトの名前と説明 - 無料

パッケージにはインストーラーが含まれていないため、
reka.md 、 tibor.md 、および gergo.md の <WRITE YOUR PROJECT NAME HERE> に入力し、
<1 ～ 2 文を書く...> 箇所に手書きでマークを付けます。 _protokoll.md はインストーラーに同梱されています
「初心者」のデフォルト — 合わない場合は、これも手動で書き換えてください。
既知の制限: reka.md と _protokoll.md は CLAUDE.md と
_hierarchia.md — これらのファイルは有料パッケージにのみ存在します。無料インストールの場合は、次の参照を参照してください
空のままにしておきます (エージェントは単にそのようなファイルを見つけられません)。これはエラーを引き起こすのではなく、単に意味するだけです
エスカレーション パスや広範な規約ドキュメントは存在しません。
ブランチ feat/42-example-feature 上のデモ リポジトリで munkarend-kapu.mjs を実行します。
Agent-kit.config.json は存在しますが、GitHub トークンはありません (不明な SP → L に切り上げられます):
ワークフロー ゲート — セッション開始プロトコル (サンプル プロジェクト)
ステップ 0 (必須、何かに触れる前): チケットを特定し、そのサイズを計算します。 SP → サイズ: 1-2 = S、3-5 = M、8+ = L;不明/欠落している SP → L に切り上げられます。実装する前にサイズを声に出して言います。
支店名からのチケット: #42。
サイズを明確にする必要があります (使用可能な SP がありません) — 切り上げて L として扱います。
サイズ L の必須チェーン — PR トレース: tibor (テスト) → reka (レビュー) → zsofi (ドキュメント) → columbo (オーケストレーション監査) · チケット トレース: sara (分析) → bence (アーキテクチャ)。
厳格なルール: サイズ M/L では、ディスパッチャーは実装しません。レイヤーの割り当てとコードを izsak にルーティングし、チェーンを調整するだけです。
すべての変更後に必須: `npm test` — 緑色のみをコミットします。
(上記のサイズ L チェーンには、 zsofi 、 columbo 、 sara 、 bence もリストされています。これらのエージェントは、
有料パッケージ;ここでのチェーン エンジンは、チーム全体が必要とするものを単に示しているだけです)。
エージェントなしの PR に対するチェーン トレース ボットのコメントを示すライブ デモ PR

まったく痕跡がありません:
horilla98/デモエージェントキット#2
11 人のエージェント チーム全体、13 ステップのライフサイクル、スラッシュ コマンド、インストーラー、
バイリンガル テンプレート ツリーはここから入手できます: https://buy.polar.sh/polar_cl_HvLQ28C9lIlP8cjdU9Gx6Bked8w5yIhPF5ASg1lt2S7
連絡先: horilla.dniel97@gmail.com
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to horilla98/agent-kit development by creating an account on GitHub.

GitHub - horilla98/agent-kit · GitHub
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
horilla98
/
agent-kit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .claude/ agents .claude/ agents .github .github LICENSE LICENSE README.md README.md View all files Repository files navigation
Agent-kit (free core) — a size-based workflow gate for Claude Code
This is the free, open-source core of agent-kit : the engine behind a size-based mandatory
review chain. It reminds you of the ticket's size and its required chain steps at the start of
every Claude Code session, then tracks on the PR — via a bot comment — which chain members have
actually left a trace.
This is not the full agent-kit. The full package (11 agent charters, 7 slash commands, the
full 13-step development lifecycle, a bilingual template tree, an installer) is paid — see below.
What's in here (verified against the agent-kit v2.0 source)
8 scripts ( .github/scripts/ ): lanc-tabla , munkarend-kapu + munkarend-kapu-logic ,
lanc-nyom + lanc-nyom-logic , gh-api , projekt-config , nyelv-cimkek , uzenetek
2 GitHub Actions workflows ( .github/workflows/ ): ci.yml , pr-jelzok.yml
3 agent charters out of the full 11: reka (code reviewer), tibor (test engineer),
gergo (security + privacy, with veto power) + their shared _protokoll.md
45 automated tests , green on their own — the publikus/ cut verified with its own test
suite, with no dependency on the paid part
Zero npm dependencies, only Node.js node: built-ins (Node.js ≥ 22.6)
Not included: the other 8 agent charters (sara, bence, marci, izsak, zsofi, petra, devops,
columbo), the 7 slash commands, the Backlog import workflow, the installer script, the
bilingual template tree, the doc skeletons (ADR, CONTRIBUTING, workflow doc).
Copy the contents of .github/scripts/ , .github/workflows/ , and .claude/agents/ into your
own repo, then add an agent-kit.config.json at the repo root:
{
"projekt" : " Your project's name " ,
"repo" : " your-github-username/your-repo " ,
"nyelv" : " english " ,
"teszt" : " npm test "
}
One manual step is still needed: in the paid package, an installer script fills in the
project's name and description inside the agent charters — the free package has no installer, so
in reka.md , tibor.md , and gergo.md fill in the <WRITE YOUR PROJECT NAME HERE> and
<WRITE 1-2 SENTENCES...> marked spots by hand. _protokoll.md ships with the installer's
"beginner" default — rewrite that by hand too if it doesn't fit you.
Known limitation: reka.md and _protokoll.md reference a CLAUDE.md and a
_hierarchia.md — these files only exist in the paid package. On a free install these references
stay empty (the agent simply won't find such a file); this doesn't cause an error, it just means
the escalation paths and the broader convention doc aren't there.
Running munkarend-kapu.mjs in a demo repo, on branch feat/42-example-feature , with
agent-kit.config.json present, no GitHub token (unknown SP → rounded up to L):
Workflow gate — session-start protocol (Example Project)
STEP 0 (mandatory, BEFORE you touch anything): identify the ticket and compute its size. SP → size: 1-2 = S, 3-5 = M, 8+ = L; unknown/missing SP → rounded up to L. State the size out loud before you implement.
Ticket from the branch name: #42.
Size MUST BE CLARIFIED (no usable SP available) — treat it as L, rounded up.
Mandatory chain for size L — PR trace: tibor (test) → reka (review) → zsofi (docs) → columbo (orchestration audit) · ticket trace: sara (analysis) → bence (architecture).
Hard rule: at size M/L the dispatcher does NOT implement — it routes the layer assignment and the code to izsak, and only coordinates the chain.
Mandatory after every change: `npm test` — only commit green.
(The size-L chain above also lists zsofi , columbo , sara , bence — these agents live in the
paid package; here the chain engine simply shows what the full team would require.)
Live demo PR showing the chain-trace bot comment on a PR with no agent traces at all:
horilla98/demo-agent-kit#2
The full 11-agent team, the 13-step lifecycle, the slash commands, the installer, and the
bilingual template tree are available here: https://buy.polar.sh/polar_cl_HvLQ28C9lIlP8cjdU9Gx6Bked8w5yIhPF5ASg1lt2S7
Contact: horilla.dniel97@gmail.com
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
