---
source: "https://github.com/commensa-ai/commensa-audit"
hn_url: "https://news.ycombinator.com/item?id=48517988"
title: "Show HN:I audited 162 agent-written PRs – 27% were the AI fixing itself"
article_title: "GitHub - commensa-ai/commensa-audit: What % of your AI engineering effort went to fixing the AI's own work? One-page rework report from git history. Read-only, local-first. · GitHub"
author: "aimattb"
captured_at: "2026-06-13T15:37:22Z"
capture_tool: "hn-digest"
hn_id: 48517988
score: 3
comments: 1
posted_at: "2026-06-13T15:01:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN:I audited 162 agent-written PRs – 27% were the AI fixing itself

- HN: [48517988](https://news.ycombinator.com/item?id=48517988)
- Source: [github.com](https://github.com/commensa-ai/commensa-audit)
- Score: 3
- Comments: 1
- Posted: 2026-06-13T15:01:46Z

## Translation

タイトル: Show HN: エージェントが作成した 162 件の PR を監査しました – 27% は AI が自ら修正していました
記事のタイトル: GitHub - commensa-ai/commensa-audit: AI エンジニアリングの取り組みの何％が AI 自体の動作の修正に費やされましたか? git 履歴からの 1 ページのリワーク レポート。読み取り専用、ローカルファースト。 · GitHub
説明: AI エンジニアリングの取り組みの何パーセントが AI 自体の動作の修正に費やされましたか? git 履歴からの 1 ページのリワーク レポート。読み取り専用、ローカルファースト。 - コメンサ-ai/コメンサ-監査

記事本文:
GitHub - commensa-ai/commensa-audit: AI エンジニアリングの取り組みの何パーセントが AI 自体の動作の修正に費やされましたか? git 履歴からの 1 ページのリワーク レポート。読み取り専用、ローカルファースト。 · GitHub
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
コメンサアイ
/
コメンサ監査
公共
通知
あなたはしなければなりません

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット commensa_audit commensa_audit 品質 品質リファレンス リファレンス レビュー レビュー スイープ スイープ テスト テスト .gitignore .gitignore BUILD_LOG.md BUILD_LOG.md LICENSE LICENSE PICKUP.md PICKUP.md README.md README.md SPEC.md SPEC.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エンジニアリングの取り組みの何％が AI 自体の動作の修正に費やされましたか?
commensa-audit は git 履歴からそれを答えます。 GitHub リポジトリを指定します。 1 ページのレポートを取得します。
リワーク税 — 以前の作業を修正した PR (および変更された行) の割合と純新規価値の比較
置き換えられた作業 — 後で出力が完全に置き換えられた PR (個別に表示 — 破棄 ≠ 修正)
放棄された試行 - PR はマージせずに閉じられました: 無駄なマージベースのメトリックは表示されません
チャーン クラスター — 相互に書き換える PR のチェーン (「ダーク モードを適切に実行するには 10 個の PR が必要でした」)
行の生存 — ウィンドウの終わりにマージされたコードがどのくらい残っているか
ホットスポット — リポジトリ全体の割合に対するモジュールごとのリワークシェア
エージェントがマークしたシェア — 「PR の少なくとも X% にはエージェント マーカーが含まれている」（共著のトレーラー、本文の署名） — 明示された下限であり、帰属を主張することはありません
私たちはそれが必要だったのでそれを構築しました。独自のエージェントで構築された製品は 13 日間で 162 個の PR を出荷しましたが、監査ではそのうち 27% が AI による自己修正であることが判明しました。
pip インストール コメンサ監査
commensa-audit --repo 所有者/名前 --token $GH_TOKEN
またはソースから直接:
pip install git+https://github.com/commensa-ai/commensa-audit
出力: report_<repo>.html (自己完結型、転送可能)、audit_<repo>.json (生の数値)、units.csv (PR ごとのデータ)。
デフォルトでは、

監査は最新の 500 の PR を対象としています。これは、巨大なリポジトリでの単純な実行が高速かつ制限されたままであるための安全上限です。切り詰められると、実行はそれを引き上げる方法を示す通知を出力します。 2 つのオプションのフラグはウィンドウを制御します (両方とも新しいものから順)。
commensa-audit --repo 所有者/名前 --since 2026-03-14 --max-prs 150
--since YYYY-MM-DD — この UTC 日付以降に作成された PR のみ
--max-prs N — N 個の最新 PR に上限を設定します (デフォルトは 500、上限なしの場合は --max-prs 0 を使用します)
どちらもページネーションを早期に停止するため、 --max-prs 150 では、リポジトリの履歴全体ではなく、最大 150 PR 相当の API 呼び出しがかかります。 PR が 500 未満のリポジトリでフラグを指定せずに実行すると、以前とまったく同じようにすべてが取得されます。
読み取り専用。 GET リクエストのみ。読み取りスコープを持つトークンで十分です。
地元第一主義。すべてが実行され、マシン上に残ります。テレメトリーやテレフォンホームはなく、ネットワークからは何も出ません。
検査可能。純粋な Python、stdlib + リクエスト + jinja2 。実行する前にすべての行を読んでください。
分類の仕組み (およびその正直な限界)
すべての PR は、透明なシグナル カスケードによって分類されます。つまり、明示的な修正タイトル/取り消し → 自己修正 (主に過去 N 日間に追加された行を元に戻す PR) → チャーン クラスター メンバーシップ → その他の生成。出力内のすべての分類には、発生した信号と人間が判読できる理由が含まれます。しきい値は 1 つの構成ブロック内に存在します。それらを調整し、 --reuse を使用してオフラインで再実行します。
既知の制限 (レポートのフッターにも表示されます): 分類はヒューリスティックです。スカッシュは属性をぼかしマージします。生存ウィンドウは、若いリポジトリが楽観的に読み取られることを意味します。エージェントがマークしたシェアは下限です。マーカーが存在しないことは、人間の作成者の証拠にはなりません。私たちは偽りの精度ではなく、私たち自身の確実性を評価します。それがこのプロジェクトの重要な点です。
エージェント時代のチームは、PR のマージ、ラインの出荷、速度などのアクティビティを測定します。非

これにより、進行状況とクリーンアップが区別されます。リワーク税は、修正されたモーションの割合であり、「この作業がどの程度適切に指示されたか?」を最もよく表す git のみの代用です。これですべてがわかるわけではありません (結果あたりのコストには、git が持っていないトークン データが必要です。これが次に構築するものです)。しかし、これは最も正直な最初の数字であり、無料です。
このツールはスナップショットです。 Commensa はトレンドラインです。チームおよびモジュールごとの継続的なリワーク測定、アラート、月次経営レポート、および Git では確認できないコスト面がエージェント ハーネスでキャプチャされます。最初の 25 社: 設立パートナー条件。
騒音ではなく耐久性を測定してください。
AI エンジニアリングの取り組みの何％が AI 自体の動作の修正に費やされましたか? git 履歴からの 1 ページのリワーク レポート。読み取り専用、ローカルファースト。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

What % of your AI engineering effort went to fixing the AI's own work? One-page rework report from git history. Read-only, local-first. - commensa-ai/commensa-audit

GitHub - commensa-ai/commensa-audit: What % of your AI engineering effort went to fixing the AI's own work? One-page rework report from git history. Read-only, local-first. · GitHub
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
commensa-ai
/
commensa-audit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits commensa_audit commensa_audit quality quality reference reference reviews reviews sweep sweep tests tests .gitignore .gitignore BUILD_LOG.md BUILD_LOG.md LICENSE LICENSE PICKUP.md PICKUP.md README.md README.md SPEC.md SPEC.md pyproject.toml pyproject.toml View all files Repository files navigation
What % of your AI engineering effort went to fixing your AI's own work?
commensa-audit answers that from your git history. Point it at a GitHub repo; get a one-page report:
Rework tax — share of PRs (and changed lines) that corrected earlier work, vs. net-new value
Superseded work — PRs whose output was entirely replaced later (shown separately — discarded ≠ correcting)
Abandoned attempts — PRs closed without merging: the waste merge-based metrics never see
Churn clusters — chains of PRs rewriting each other ("it took 10 PRs to get dark mode right")
Line survival — how much merged code is still alive at the end of the window
Hotspots — rework share by module, against the repo-wide rate
Agent-marked share — "at least X% of PRs carry agent markers" (Co-Authored-By trailers, body signatures) — a stated lower bound, never an attribution claim
We built it because we needed it: our own agent-built product shipped 162 PRs in 13 days, and the audit showed 27% of them were the AI correcting itself .
pip install commensa-audit
commensa-audit --repo owner/name --token $GH_TOKEN
Or straight from source:
pip install git+https://github.com/commensa-ai/commensa-audit
Output: report_<repo>.html (self-contained, forwardable), audit_<repo>.json (raw numbers), units.csv (per-PR data).
By default the audit covers the newest 500 PRs — a safety cap so a naive run on a huge repo stays fast and bounded. When it truncates, the run prints a notice telling you how to raise it. Two optional flags control the window (both newest-first):
commensa-audit --repo owner/name --since 2026-03-14 --max-prs 150
--since YYYY-MM-DD — only PRs created on/after this UTC date
--max-prs N — cap to the N newest PRs (default 500 ; use --max-prs 0 for no cap)
Both early-stop pagination, so --max-prs 150 costs ~150 PRs' worth of API calls, not the repo's entire history. Run with no flags on a repo under 500 PRs and you get everything, exactly as before.
Read-only. GET requests only; a token with read scope is sufficient.
Local-first. Everything runs and stays on your machine. No telemetry, no phone-home, nothing leaves your network.
Inspectable. Pure Python, stdlib + requests + jinja2 . Read every line before you run it.
How classification works (and its honest limits)
Every PR is classified by a transparent signal cascade — explicit corrective titles/reverts → self-correction (a PR predominantly undoing lines added in the prior N days) → churn-cluster membership → otherwise generative. Every classification in the output carries the signal that fired and a human-readable why. Thresholds live in one config block; tune them and re-run offline with --reuse .
Known limits (also printed in the report footer): classification is heuristic; squash merges blur attribution; survival windows mean young repos read optimistic; agent-marked share is a lower bound — absence of a marker is not evidence of human authorship. We grade our own certainty rather than fake precision — that's the whole point of the project.
Agent-era teams measure activity — PRs merged, lines shipped, velocity. None of that distinguishes progress from cleanup. The rework tax does: it's the share of motion that was correction, the closest git-only proxy for "how well was this work directed?" It won't tell you everything (cost-per-outcome needs token data git doesn't have — that's what we're building next ) — but it's the most honest first number, and it's free.
This tool is a snapshot. Commensa is the trendline: continuous rework measurement by team and module, alerts, monthly executive reports — and the cost side git can't see, captured at the agent harness. First 25 companies: founding-partner terms.
measure the durable work, not the noise.
What % of your AI engineering effort went to fixing the AI's own work? One-page rework report from git history. Read-only, local-first.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
