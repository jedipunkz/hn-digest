---
source: "https://github.com/izeigerman/claude-thermos"
hn_url: "https://news.ycombinator.com/item?id=49024882"
title: "Show HN: Claude-thermos – keeps your Claude session warm for you"
article_title: "GitHub - izeigerman/claude-thermos: Keeps your Claude session warm for you · GitHub"
author: "s0ck_r4w"
captured_at: "2026-07-23T17:09:24Z"
capture_tool: "hn-digest"
hn_id: 49024882
score: 1
comments: 0
posted_at: "2026-07-23T17:06:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude-thermos – keeps your Claude session warm for you

- HN: [49024882](https://news.ycombinator.com/item?id=49024882)
- Source: [github.com](https://github.com/izeigerman/claude-thermos)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T17:06:22Z

## Translation

タイトル: 表示 HN: Claude-thermos – クロード セッションを暖かく保ちます
記事のタイトル: GitHub - izeigerman/claude-thermos: クロード セッションを暖かく保ちます · GitHub
説明: クロード セッションを温かく保ちます。 GitHub でアカウントを作成して、izeigerman/claude-thermos の開発に貢献してください。

記事本文:
GitHub - izeigerman/claude-thermos: クロード セッションを暖かく保ちます · GitHub
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
アイゲルマン
/
クロード・サーモス
公共
通知
通知設定を変更するにはサインインする必要があります
追加ナビ

ゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
40 コミット 40 コミット .github/ workflows .github/ workflows src/ claude_thermos src/ claude_thermos テスト テスト .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示リポジトリ ファイルのナビゲーション
クロード コード キャッシュを再構築するためにお金を払うのはやめましょう。メイン エージェントがサブエージェントを 5 分以上待機すると、そのプロンプト キャッシュはサイレントに期限切れになり、次のターンでは会話全体を低コストで読み戻すのではなく、書き込みレートで再エンコードします。多くのサブエージェントとの長時間のセッションでは、請求額の約 20% になります。 claude-thermos はキャッシュを暖かく保つので、その税金を支払う必要はありません。
通常とまったく同じように Claude コードを実行しますが、 uvx を使用した claude-thermos を使用します。
uvx claude-thermos # の代わりに: claude
uvx claude-thermos -p " バグを修正 " # 任意のクロード引数はそのまま通過します
Python 3.11+ と PATH 上の claude CLI が必要です。
それでおしまい。ウォーミングはバックグラウンドで自動的に実行されます。コマンドを変更せずに実行時にこれを無効にするには、 CLAUDE_WARMER_DISABLE=1 を設定します。
Claude Code のプロンプト キャッシュは 5 分間の TTL を使用します。キャッシュが存続している限り、会話履歴全体が全額で再送信されるのではなく、毎ターン、入力価格の 0.1 倍でキャッシュから提供されます。
同じプレフィックスに対するリクエストの間に 5 分以上経過すると、キャッシュの有効期限が切れます。そのギャップを引き起こす主な原因は、あなたが考えていることではありません。これは、5 分を超えて実行されるサブエージェントでブロックされているメイン エージェントです。サブエージェントには異なるシステム プロンプトとツール セットがあるため、そのリクエストには異なるキャッシュ プレフィックスが付けられ、メイン エージェントのリクエストは更新されません。サブエージェントが動作している間、メイン エージェントのキャッシュされた履歴はそのままの状態で古くなっています。パス

t 5 分で終わります。サブエージェントが戻ると、メイン エージェントはバイト同一の追加のみの履歴で再開し、キャッシュが欠落していることに気づき、1.25 倍の書き込み速度で完全な再エンコードを強制します。
その時点までに履歴が大きくなるため、再エンコードにコストがかかります。個々の崩壊で 200K から 500K のトークンが再書き込みされます。約 185 のローカル セッションにわたって測定されたこれらの再構築は、合計請求額の約 22% を占め、直前にすでにキャッシュされていたコンテンツの再エンコードに費やされた費用です。
claude-thermos は、小規模なローカル リバース プロキシの背後で Claude Code を起動します (ループバック ポートで ANTHROPIC_BASE_URL を指します。すべてのトラフィックは依然として実際の Anthropic API に送られます)。
観察する。プロキシは /v1/messages トラフィックを監視し、それをセッションとリネージにグループ化します。リネージは 1 つのキャッシュ プレフィックスであり、モデル + ツール セット + システム テキストをキーとします。最初のツールを持つ系統が主要なエージェントです。残りはサブエージェントです。
危険な窓を検知します。メイン リネージュがアイドル状態になり、サブエージェントがアクティブに実行されている場合、メイン プレフィックスが期限切れになる危険があります。
暖かい。 5 分の TTL 未満の間隔で、メイン エージェントの最後の実際のリクエストをウォーム リクエストとして再生します。キャッシュ可能なプレフィックスは同一ですが、max_tokens: 1 でストリーミングはありません。単一のトークンは捨てられます。重要なのは、キャッシュされたプレフィックス全体を読み取って更新するプレフィルです。ウォーム リクエストはプロキシを経由せずに API に直接送信されるため、実際のトラフィックを妨げることはありません。
結果。サブエージェントが終了しても、メイン エージェントのキャッシュはまだウォーム状態です。完全な書き換えではなく、安価な読み取り料金がかかります。
ウォームごとにキャッシュ読み取りコストがかかります (0.1 倍)。書き換えを防ぐたびに、はるかに大きなプレフィックスへの書き込みコスト (1.25 倍) が発生するため、取引は非常に有利になります。
~/.claude-thermos/logs/<セッションID>/
§── events.jsonl # 追加専用の構造化イベント ストリーム
━━

summary.json # セッション終了時に書き込まれるロールアップ合計
events.jsonl は、各リクエスト/レスポンスのトークン使用量と、すべてのウォーミング決定 (warm_fired、warm_result、cap_reached、resume_detected など) を記録します。 summary.json は、通常読むことになるロールアップです。
3 つのコストの数値はすべて、基本入力トークン単位です (トークン数はキャッシュ乗数によってすでに重み付けされています)。 net_ Savings をドルに変換するには、入力トークンあたりのモデルの価格を掛けます。
節約されたドル ≈ net_ Savings × (入力トークン価格)
たとえば、入力価格が 3 ドル / 100 万トークンの場合、純節約額 1_200_000 は、そのセッションで約 1_200_000 × 3 ドル / 1_000_000 = 3.60 ドル節約されます。
クロードセッションを温かく保ちます
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Keeps your Claude session warm for you. Contribute to izeigerman/claude-thermos development by creating an account on GitHub.

GitHub - izeigerman/claude-thermos: Keeps your Claude session warm for you · GitHub
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
izeigerman
/
claude-thermos
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
40 Commits 40 Commits .github/ workflows .github/ workflows src/ claude_thermos src/ claude_thermos tests tests .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Stop paying to rebuild your Claude Code cache. When your main agent waits on a subagent for more than 5 minutes, its prompt cache silently expires, and the next turn re-encodes your entire conversation at the write rate instead of reading it back cheap. On long sessions with many subagents that's roughly 20% of your bill. claude-thermos keeps the cache warm so you never pay that tax.
Run Claude Code exactly as you normally would, but through claude-thermos with uvx :
uvx claude-thermos # instead of: claude
uvx claude-thermos -p " fix the bug " # any claude args pass straight through
Requires Python 3.11+ and the claude CLI on your PATH .
That's it. Warming runs automatically in the background. To disable it for a run without changing the command, set CLAUDE_WARMER_DISABLE=1 .
Claude Code's prompt cache uses a 5-minute TTL . Every turn, your whole conversation history is served from cache at 0.1x the input price instead of being re-sent at full price, as long as the cache stays alive.
The cache expires if more than 5 minutes pass between requests on the same prefix. The dominant trigger for that gap is not you thinking. It's the main agent blocked on a subagent that runs longer than 5 minutes . A subagent has a different system prompt and tool set, so its requests have a different cache prefix and never refresh the main agent's. While the subagent works, the main agent's cached history ages untouched; past 5 minutes it's gone. When the subagent returns, the main agent resumes with a byte-identical, append-only history, and finds its cache missing, forcing a full re-encode at the 1.25x write rate.
By then the history is large, so the re-encode is expensive: individual collapses re-write 200K to 500K tokens. Measured across roughly 185 local sessions, these rebuilds accounted for about 22% of the total bill , money spent re-encoding content that was already cached moments earlier.
claude-thermos launches Claude Code behind a small local reverse proxy (it points ANTHROPIC_BASE_URL at a loopback port; all traffic still goes to the real Anthropic API).
Observe. The proxy watches /v1/messages traffic and groups it into sessions and lineages , a lineage being one cache prefix, keyed by model + tool set + system text. The first tool-bearing lineage is the main agent; the rest are subagents.
Detect the danger window. When the main lineage goes idle and a subagent is actively running, the main prefix is at risk of expiring.
Warm. On an interval under the 5-minute TTL, it replays the main agent's last real request as a warm request : identical cacheable prefix, but max_tokens: 1 and no streaming. The single token is thrown away; the point is the prefill, which reads and refreshes the full cached prefix. Warm requests go directly to the API, never through the proxy, so they can't disturb real traffic.
Result. When the subagent finishes, the main agent's cache is still warm. It pays a cheap read instead of a full rewrite.
Each warm costs a cache read (0.1x); each rewrite it prevents would have cost a write (1.25x) on a much larger prefix, so the trade is heavily in your favor.
~/.claude-thermos/logs/<session_id>/
├── events.jsonl # append-only structured event stream
└── summary.json # rollup totals, written when the session ends
events.jsonl records each request/response's token usage plus every warming decision ( warm_fired , warm_result , cap_reached , resume_detected , and so on). summary.json is the rollup you'll usually read:
All three cost figures are in base-input-token units (token counts already weighted by their cache multiplier). To turn net_savings into dollars, multiply it by your model's price per input token :
dollars saved ≈ net_savings × (input token price)
For example, at an input price of $3 / 1M tokens, a net_savings of 1_200_000 is about 1_200_000 × $3 / 1_000_000 = $3.60 saved that session.
Keeps your Claude session warm for you
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
