---
source: "https://github.com/jianzhichun/permafrost"
hn_url: "https://news.ycombinator.com/item?id=48471898"
title: "Permafrost – freeze Claude Code's prompt prefix, cut your DeepSeek bill 64%"
article_title: "GitHub - jianzhichun/permafrost: Freeze Claude Code's prompt prefix so DeepSeek's automatic cache always hits — alignment proxy + coalescing + keepalive, installable as a CC plugin. Measured 64% cheaper on real Claude Code traffic. · GitHub"
author: "jianzhichun"
captured_at: "2026-06-10T07:37:02Z"
capture_tool: "hn-digest"
hn_id: 48471898
score: 3
comments: 2
posted_at: "2026-06-10T05:41:00Z"
tags:
  - hacker-news
  - translated
---

# Permafrost – freeze Claude Code's prompt prefix, cut your DeepSeek bill 64%

- HN: [48471898](https://news.ycombinator.com/item?id=48471898)
- Source: [github.com](https://github.com/jianzhichun/permafrost)
- Score: 3
- Comments: 2
- Posted: 2026-06-10T05:41:00Z

## Translation

タイトル: Permafrost – クロード コードのプロンプト プレフィックスを凍結し、DeepSeek の請求額を 64% 削減します
記事のタイトル: GitHub - jianzhichun/permafrost: クロード コードのプロンプト プレフィックスをフリーズして、DeepSeek の自動キャッシュが常にヒットするようにします — アライメント プロキシ + 合体 + キープアライブ、CC プラグインとしてインストール可能。実際のクロード コードのトラフィックでは 64% 安く測定されました。 · GitHub
説明: クロード コードのプロンプト プレフィックスをフリーズして、DeepSeek の自動キャッシュが常にヒットするようにします (アライメント プロキシ + 合体 + キープアライブ、CC プラグインとしてインストール可能)。実際のクロード コードのトラフィックでは 64% 安く測定されました。 - 建志春/永久凍土

記事本文:
GitHub - jianzhichun/permafrost: クロード コードのプロンプト プレフィックスをフリーズして、DeepSeek の自動キャッシュが常にヒットするようにします — アライメント プロキシ + 合体 + キープアライブ、CC プラグインとしてインストール可能。実際のクロード コードのトラフィックでは 64% 安く測定されました。 · GitHub
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
{{ メッセージ

}}
建志春
/
永久凍土
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
13 コミット 13 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows ベンチマーク ベンチマーク cli cli コマンド コマンド docs docs e2e e2e フック フック プロキシ プロキシ スクリプト スクリプト テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md settings.example.json settings.example.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿
| _ \| __| _ \ \/ | /_\ | __| _ \/ _ \/ __|_ _|
| _/| _|| / |\/| |/ _ \| _|| / (_) \__ \ | |
|_| |__|_|_\_| |_/_/ \_\_| |_|_\\___/|___/ |_|
プレフィックスを凍結する・紙幣を溶かす
DeepSeek の自動キャッシュが常にヒットするように、プロンプト プレフィックスをフリーズする Claude Code プラグイン。
キャッシュ安定パススルー プロキシ · 決定的なツール順序 · 揮発性コンテンツの再配置 · ライブ ヒット レート ステータスライン · ゼロ デプス
Claude API トラフィックの $100 ごとに、Permafrost + DeepSeek を通じて ~$3.20 で実行されます (同じリクエスト、ライブ測定)。 (永久凍土なしで $8.80 まで。)
DeepSeek で Claude Code をポイントすると、数セントでコーディング エージェントを手に入れることができます
— DeepSeek のキャッシュがヒットし続ける場合。通常はそうではありません。クロード・コードの再シャッフル
MCP サーバーが接続すると、ツール リストが表示され、今日の日付とライブ git ステータスがシステム プロンプトに焼き付けられ、cache_control マーカーがスライドして表示されます。
カスタム エンドポイント — ツールの延期を停止し、ツール セット全体を毎回再インライン化します。
回ります。これらのそれぞれはリクエストの先頭を変更し、DeepSeek の
キャッシュは、バイト 0 から一致するプレフィックスのみにヒットします。だからあなたは黙ってお金を支払います
~無料であるはずのトークンの全額 (約 50 倍) の価格を見逃しています。
永久凍土はClの間にある

aude Code と DeepSeek を書き換えて、
すべてのリクエストの関連バイトをキャッシュして、ツールとシステム アンカーを維持する
バイト同一のターンを繰り返す — その後、応答をそのままストリーミングして返します。
クロード・コード ──Anthropic /v1/messages──▶ Permafrost ──▶ DeepSeek /anthropic
(変更なし) 127.0.0.1:8787 フリーズ プレフィックス (変換なし —
+ レコードのヒット曲はどちらも人間性を語っています)
計算: 同じワークロード、4 つの方法
実際の Claude Code e2e 実行からの正確なトークン トラフィック (4 リクエスト: 41,728)
キャッシュヒット + 21,339 キャッシュミス入力トークン、591 出力トークン)、4 通りの価格設定
現在の公定レートでは:
† 推定ではなく実際に測定されたものです。 31 倍はきれいに因数分解されます: DeepSeek の 11 倍
キャッシュヒットを維持し続ける Permafrost からの価格 × 2.8 倍。背後にある単価
it (USD/100万トークン): キャッシュヒット入力 $0.30 vs $0.0028 (107×)、入力ミス
3.00 ドル対 0.14 ドル (21 回)、出力は 15 ドル対 0.28 ドル (54 回)。オーパス層について
( claude-opus* は v4-pro にマップされます) 同じワークロードは $0.169 → $0.0099、17 倍になります
安くなります。
きちんとしたクロスチェック: Claude Code 自体が、e2e リクエストの 1 つに次の価格を設定しました。
total_cost_usd: $0.0625 (独自のソネットレート会計);同じリクエスト
Permafrost → DeepSeek 経由では、0.0029 ドルが請求されました。
正直な注意点: (1) これはモデルではなく、同一のトラフィックの価格を比較します。
品質 — deepseek-v4-flash は Claude Sonnet 4.6 ではありません。取引があるかどうか
価値があるのはあなたの呼びかけです。 (2) API 従量制と API 従量制です。クロード・マックス
サブスクリプションの価格は異なります。 (3) 行 1→3 は、DeepSeek の価格設定です。
仕事。行 3→4 (-64%) は Permafrost が上に追加したものです。
DeepSeek のプレフィックス キャッシュの忠実なエミュレータに対するオフライン ベンチマーク
(バイト 0 からのプレフィックス、ブロック量子化)、12 ターンのクロード コード形式の再生
会話。オフ = アライメントなし。アグレッシブ = デフォルト。
同一のトラフィックで ≈48% 安くなります (dri による)

11回のリセットからキャッシュアンカーを維持
バスターごとの内訳 (ツール注文チャーンのみ、git/env のみ、両方) は次のとおりです。
ベンチマーク/RESULTS.md 。
python3 benchmarks/bench.py --turns 12 # 再現 (API キーなし)
python3 benchmarks/bench.py --real # + ライブ DeepSeek プローブ (キーが必要)
実際の api.deepseek.com/anthropic エンドポイントに対してライブ検証済み:
繰り返しのプレフィックスは hit=1536 miss=77 を返します — 2 番目のキャッシュ ヒットは 95% です
リクエストは同一であり、DeepSeek の自動キャッシュが Permafrost のキャッシュを提供していることを確認しています。
正規のリクエスト形状。 (ベンチマーク/RESULTS.md を参照してください。)
実際のクロード コード (合成リクエストではない) で検証: headless claude -p
4 ターンのバグ修正タスクで、孤立し、Permafrost → DeepSeek 経由でルーティング —
本物の CC の ~85 KB / ~21,000 トークンで 66% のキャッシュ ヒット率、64% のコスト削減
システムプロンプト + ツール。再現: e2e/run_claude_code.sh
(資金提供された DEEPSEEK_API_KEY が必要です)。このランニングが私たちに教えてくれたこと、それも含めて
DeepSeek はキャッシュする前に再レンダリングし、CC のリクエストごとの請求ノンスが発生します。
docs/e2e-findings.md 。
プロキシ経由でも検証: Permafrost 経由で送信された 4 つの実際のリクエスト
ツールの順序をシャッフルし、毎回 git ステータスを変更する - まさに
オフ モードで欠落するトラフィック — 1 つのフリーズされたアンカーに調整されました
( prefix_resets=0 )。プレフィックスがウォームになると、DeepSeek はそのプレフィックスの約 97% を提供しました
キャッシュから (単一のウォーム リクエスト: hit=2048 、 miss=55 → 97% )。代理人
ナイーブなクライアントのチャーンにもかかわらず、アンカーを安定に保ちました。
/v1/messages?beta=true クエリ文字列のパス。
「アンカーのリセット」 = ツール + システム プレフィックスがバイトを変更した回数
走り。リセットのたびに、DeepSeek は正規料金でプレフィックス全体を再読み取りするように強制されます。
永久凍土の仕事はそれを 0 に保つことです。
# 1 — インストール (どこにでもクローンを作成します。プロキシは stdlib のみの Python であり、pip は必要ありません)

）
git clone https://github.com/jianzhichun/permafrost && cd permafrost
# 2 — DeepSeek を指定して、Permafrost 経由でクロード コードを起動します。
エクスポート ANTHROPIC_API_KEY=sk-your-deepseek-key
./cli/permafrost Wrap # プロキシを開始し、環境を設定し、`claude` を実行します
永久凍土ラップは、 ANTHROPIC_BASE_URL + ENABLE_TOOL_SEARCH=true を設定します。
子クロードプロセスのみ - シェルに触れたり、
~/.claude/settings.json 。
永続的な方がいいですか？ env ブロックをコピーします
settings.example.json を ~/.claude/settings.json に、
permafrost up を実行してから、claude を通常どおり起動します。 (クロード コードはその環境を 1 回読み取ります
起動時に設定されるため、これらはクロードを開始する前に設定する必要があります。)
クロードコードプラグインとしてインストールする
このリポジトリは独自のプラグイン マーケットプレイスであり、Claude Code 内の 2 つのコマンドです。
/プラグイン マーケットプレイス追加 jianzhichun/permafrost
/プラグインのインストール permafrost@permafrost
(またはシェルから: claude plugin Marketplace add jianzhichun/permafrost && claude plugin install permafrost@permafrost )
これにより、/permafrost:* コマンド、SessionStart 自動開始フック、
そしてステータスラインスクリプト。プロキシ自体はプラグインの一部です — permafrost ラップ / 上記の設定ブロックはクロード コードを指します。後で更新してください
/plugin マーケットプレイスは永久凍土を更新します。
/v1/messages リクエストごとに、アライメント パイプライン
( proxy/permafrost_align.py ):
キャッシュ_コントロールを削除 — DeepSeek はマーカーを無視します。彼らの漂流
位置は純粋なプレフィックス ノイズです。
ツールを決定的に並べ替えます。そのため、遅延バインディングの MCP サーバーは再シャッフルできません。
プレフィックスの位置 0。
env ブロックをフリーズ + デルタを発行 (アグレッシブ モード) — を固定します
最初に見た環境/コンテキスト ブロック (cwd、プラットフォーム、今日の日付、
git status ) をキャッシュされたアンカーに挿入し、その後のターンでは、
尾部のラインが変化しました。変更されていない環境のコストはゼロです

1 ターンあたり ns;
変更されたものは、ブロック全体を再送信するのではなく、そのデルタのみがかかります
毎ターン。 (再配置にフォールバックするには、PERMAFROST_FREEZE_ENV=0 を設定します。
代わりにプレフィックスをブロック全体から削除します。)
コンパクト、UTF-8 に忠実、安定したバイトを正規にシリアル化します。
次に、DeepSeek のプロンプト_キャッシュ_ヒット_トークン / プロンプト_キャッシュ_ミス_トークンを読み取ります。
即座に応答し、ライブヒット率と節約額を追跡します。
コールド アンカー合体 (並列サブエージェントの場合)
DeepSeek はキャッシュを非同期で書き込むため、クロード コードがタスクをファンアウトすると、
サブエージェント - 1 つの新しいプレフィックスを共有する N 個のリクエストが一度に起動される - どれも実行できない
他の人がまだ書いているものを読んで、N 人全員がコールドミスの代償を支払います。
永久凍土はそれらを合体させます。目に見えないアンカー上の最初のリクエストは次のように通過します。
「リーダー」であり、キャッシュを温めます。同じアンカーのリクエストは、
リーダーの応答はストリーミングを開始し、その後、ウォーム プレフィックスを読み取るために解放されます。
単一のリクエストは決して遅延しません (単一のリクエストがリーダーです)。だけ
同時の同じアンカー バーストは、スタックしたリーダーに対してタイムアウト ガードを備えて待機します。
/permafrost:doctor は、保持しているフォロワーの数を報告します。で切り替えます
PERMAFROST_COALESCE=0 。
リリースのタイミングは、測定されたトレードオフです。DeepSeek のキャッシュ書き込みは非同期です。
(私たちのプローブ: 最初のバイトはまだミスしてから ~4 秒、ヒット後 ~6 秒)、フォロワーはしばらく待ちます。
リリース後の PERMAFROST_COALESCE_SETTLE_MS (デフォルト 2500) - デフォルト
それ自体はライブ検証されています。デフォルト設定での 3 リクエストのコールド バーストには、
両方のフォロワーが完全にヒットしました (同じシェイプで 0 を解決し、0 スコアを獲得)。最大のヒットを得るには
オッズセット PERMAFROST_COALESCE_RELEASE=completion — フォロワーは、
リーダーの応答が完全にストリーミングされると、リクエスト境界キャッシュ ユニットが
持続します。
ライブ検証済み: 実際の DeepSeek に対する 4 つの同時コールド アンカー リクエスト

から行きました
0% ヒット (オフ) — 4 つすべてミス、定価で 16,356 トークン — 73% ヒット
(on) 、1 人のリーダーがキャッシュをウォームし、3 人のフォロワーがそれを読み取ります。費用
レイテンシー: フォロワーはリーダーの最初のバイトを待ちます (さらにオプションの
DeepSeek のランドへの非同期書き込みの場合は PERMAFROST_COALESCE_SETTLE_MS)。
DeepSeek は未使用のキャッシュ エントリを削除するため、長い思考時間のギャップが生じる可能性があります
次のターンでは、会話プレフィックス全体のミス代償を全額支払うことになります。と
PERMAFROST_KEEPALIVE_S=<秒> に設定すると、プロキシは最後のリクエストを再実行します —
変更されていない、本文とヘッダー - アイドル時間が長くなった後、
ヒット価格でのプレフィックス全体 (ミスの最大 2%)。実際に起動するため、デフォルトではオフです
請求可能なリクエストは自律的に実行されます。永久凍土が暖かいと同じリプレイがトリガーされる
手動で。実際の CC リクエストのリプレイで 99.9% のヒットがライブ検証されました。の
「変更されていない、ヘッダーも」部分は負荷がかかります — DeepSeek のキャッシュ ID
クライアントのヘッダー フィンガープリントと、パラメータまたはパラメータが異なるリプレイが含まれます。
ヘッダーが明らかに欠落しています ( docs/e2e-findings.md を参照)。
セッションごとおよびリネージごとの統計 (アンカー差分付き)
/permafrost/stats クロード コードのセッション ID ごとに使用状況をバケット化し、アンカーを追跡します
リネージごとのチャーン (安定したシステム + ツールの祖先) — したがって、CC のツールレス プリフライト
通話はチャーン信号を汚染せず、1 つのプロキシを介した複数のセッション
セント

[切り捨てられた]

## Original Extract

Freeze Claude Code's prompt prefix so DeepSeek's automatic cache always hits — alignment proxy + coalescing + keepalive, installable as a CC plugin. Measured 64% cheaper on real Claude Code traffic. - jianzhichun/permafrost

GitHub - jianzhichun/permafrost: Freeze Claude Code's prompt prefix so DeepSeek's automatic cache always hits — alignment proxy + coalescing + keepalive, installable as a CC plugin. Measured 64% cheaper on real Claude Code traffic. · GitHub
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
jianzhichun
/
permafrost
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
13 Commits 13 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows benchmarks benchmarks cli cli commands commands docs docs e2e e2e hooks hooks proxy proxy scripts scripts tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md settings.example.json settings.example.json View all files Repository files navigation
___ ___ ___ __ __ _ ___ ___ ___ ___ _____
| _ \| __| _ \ \/ | /_\ | __| _ \/ _ \/ __|_ _|
| _/| _|| / |\/| |/ _ \| _|| / (_) \__ \ | |
|_| |___|_|_\_| |_/_/ \_\_| |_|_\\___/|___/ |_|
freeze the prefix · melt the bill
A Claude Code plugin that freezes your prompt prefix so DeepSeek's automatic cache always hits.
cache-stable passthrough proxy · deterministic tool order · volatile-content relocation · live hit-rate statusline · zero deps
Every $100 of Claude API traffic runs for ~$3.20 through Permafrost + DeepSeek — same requests, live-measured. (~$8.80 without Permafrost.)
Point Claude Code at DeepSeek and you get a coding agent for cents on the dollar
— if DeepSeek's cache keeps hitting. It usually doesn't. Claude Code reshuffles
its tool list when MCP servers connect, bakes today's date and a live git status into the system prompt, slides cache_control markers around, and — under
a custom endpoint — stops deferring tools and re-inlines the whole tool set every
turn. Every one of those changes the front of the request, and DeepSeek's
cache only hits on a prefix that matches from byte 0 . So you quietly pay the
full (≈50×) miss price on tokens that should have been ~free.
Permafrost sits between Claude Code and DeepSeek and rewrites the
cache-relevant bytes of every request so the tools + system anchor stays
byte-identical turn after turn — then streams the reply straight back.
Claude Code ──Anthropic /v1/messages──▶ Permafrost ──▶ DeepSeek /anthropic
(unchanged) 127.0.0.1:8787 freeze prefix (no translation —
+ record hits both speak Anthropic)
The math: same workload, four ways
The exact token traffic from our real Claude Code e2e run (4 requests: 41,728
cache-hit + 21,339 cache-miss input tokens, 591 output tokens), priced four ways
at current official rates:
† live-measured, not estimated. The 31× factors cleanly: 11× from DeepSeek's
pricing × 2.8× from Permafrost keeping the cache hitting. Unit prices behind
it (USD/1M tokens): cache-hit input $0.30 vs $0.0028 (107×), miss input
$3.00 vs $0.14 (21×), output $15 vs $0.28 (54×). On the Opus tier
( claude-opus* maps to v4-pro) the same workload goes $0.169 → $0.0099, 17×
cheaper .
A neat cross-check: Claude Code itself priced one of our e2e requests at
total_cost_usd: $0.0625 (its own Sonnet-rate accounting); the same request
through Permafrost → DeepSeek billed $0.0029 .
Honest caveats: (1) this compares price for identical traffic , not model
quality — deepseek-v4-flash is not Claude Sonnet 4.6; whether the trade is
worth it is your call. (2) It's API-metered vs. API-metered; Claude Max
subscriptions price differently. (3) Rows 1→3 are DeepSeek's pricing doing the
work; rows 3→4 (−64%) are what Permafrost adds on top.
Offline benchmark against a faithful emulator of DeepSeek's prefix cache
(prefix-from-byte-0, block-quantized), replaying a 12-turn Claude-Code-shaped
conversation. off = no alignment; aggressive = the default.
≈48% cheaper on identical traffic , by driving the cache anchor from 11 resets
to 0. The per-buster breakdown (tool-order churn alone, git/env alone, both) is in
benchmarks/RESULTS.md .
python3 benchmarks/bench.py --turns 12 # reproduce (no API key)
python3 benchmarks/bench.py --real # + live DeepSeek probe (needs key)
Live-validated against the real api.deepseek.com/anthropic endpoint: a
repeated prefix returns hit=1536 miss=77 — a 95% cache hit on the second
identical request, confirming DeepSeek's automatic cache serves Permafrost's
canonical request shape. (See benchmarks/RESULTS.md .)
Validated with real Claude Code (not synthetic requests): headless claude -p
on a 4-turn fix-the-bug task, isolated, routed through Permafrost → DeepSeek —
66% cache hit rate, 64% cost reduction , on ~85 KB / ~21 K tokens of genuine CC
system prompt + tools. Reproduce: e2e/run_claude_code.sh
(needs a funded DEEPSEEK_API_KEY ). What this run taught us — including that
DeepSeek re-renders before caching, and CC's per-request billing nonce — is in
docs/e2e-findings.md .
Validated through the proxy, too: four real requests sent through Permafrost
with shuffled tool order and changing git status each time — exactly the
traffic that misses in off mode — were aligned to one frozen anchor
( prefix_resets=0 ). Once the prefix was warm, DeepSeek served ~97% of it
from cache (a single warm request: hit=2048 , miss=55 → 97% ). The proxy
held the anchor stable despite the naive-client churn — including over the
/v1/messages?beta=true query-string path.
"Anchor resets" = how many times the tools+system prefix changed bytes across
the run. Each reset forces DeepSeek to re-read the whole prefix at full price.
Permafrost's job is to keep it at 0.
# 1 — install (clone anywhere; the proxy is stdlib-only Python, no pip needed)
git clone https://github.com/jianzhichun/permafrost && cd permafrost
# 2 — launch Claude Code through Permafrost, pointed at DeepSeek
export ANTHROPIC_API_KEY=sk-your-deepseek-key
./cli/permafrost wrap # starts the proxy, sets the env, execs `claude`
permafrost wrap sets ANTHROPIC_BASE_URL + ENABLE_TOOL_SEARCH=true for the
child claude process only — it never touches your shell or
~/.claude/settings.json .
Prefer it persistent? Copy the env block from
settings.example.json into ~/.claude/settings.json ,
run permafrost up , then start claude normally. (Claude Code reads its env once
at launch, so these must be set before claude starts.)
Install as a Claude Code plugin
This repo is its own plugin marketplace — two commands inside Claude Code:
/plugin marketplace add jianzhichun/permafrost
/plugin install permafrost@permafrost
(or from a shell: claude plugin marketplace add jianzhichun/permafrost && claude plugin install permafrost@permafrost )
That gives you the /permafrost:* commands, the SessionStart auto-start hook,
and the statusline script. The proxy itself is part of the plugin — permafrost wrap / the settings block above point Claude Code at it. Update later with
/plugin marketplace update permafrost .
On every /v1/messages request, the alignment pipeline
( proxy/permafrost_align.py ):
strips cache_control — DeepSeek ignores the markers; their drifting
positions are pure prefix noise.
sorts tools deterministically — so late-binding MCP servers can't reshuffle
position 0 of the prefix.
freezes the env block + emits deltas (aggressive mode) — pins the
first-seen environment/context block (cwd, platform, today's date,
git status ) into the cached anchor, then on later turns sends only the
lines that changed on the tail. An unchanged env costs zero tokens per turn;
a changed one costs only its delta — instead of re-sending the whole block
every turn. (Set PERMAFROST_FREEZE_ENV=0 to fall back to relocating the
whole block off the prefix instead.)
serializes canonically — compact, UTF-8-faithful, stable bytes.
It then reads DeepSeek's prompt_cache_hit_tokens / prompt_cache_miss_tokens
straight off the response and tracks your live hit rate and dollar savings.
Cold-anchor coalescing (for parallel subagents)
DeepSeek writes its cache asynchronously , so when Claude Code fans out Task
subagents — N requests sharing one brand-new prefix, fired at once — none can
read what the others are still writing, and all N pay the cold-miss price .
Permafrost coalesces them: the first request on an unseen anchor goes through as
the "leader" and warms the cache; same-anchor requests are held until the
leader's response starts streaming, then released to read the warm prefix.
Single requests are never delayed (a lone request is the leader); only
concurrent same-anchor bursts wait, with a timeout guard against a stuck leader.
/permafrost:doctor reports how many followers it held. Toggle with
PERMAFROST_COALESCE=0 .
Release timing is a measured trade-off: DeepSeek's cache write is asynchronous
(our probes: ~4s after first byte still misses, ~6s hits), so followers wait a
PERMAFROST_COALESCE_SETTLE_MS (default 2500) after release — the default
itself is live-validated: a 3-request cold burst under default settings had
both followers hit in full (settle 0 scored 0% on the same shape). For maximum hit
odds set PERMAFROST_COALESCE_RELEASE=completion — followers go only after the
leader's response fully streams, which is when the request-boundary cache unit
persists.
Live-validated: 4 concurrent cold-anchor requests to real DeepSeek went from
0% hit (off) — all four miss, 16,356 tokens at full price — to 73% hit
(on) , where one leader warmed the cache and three followers read it. The cost
is latency: followers wait for the leader's first byte (plus an optional
PERMAFROST_COALESCE_SETTLE_MS for DeepSeek's async write to land).
DeepSeek evicts cache entries that go unused, so a long think-time gap can leave
your next turn paying full miss price on the whole conversation prefix. With
PERMAFROST_KEEPALIVE_S=<seconds> set, the proxy replays the last request —
unchanged, body and headers — after that much idle time, re-reading the
whole prefix at hit price (~2% of a miss). Off by default because it fires real
billable requests autonomously. permafrost warm triggers the same replay
manually. Live-validated at 99.9% hit on a real CC request replay; the
"unchanged, headers too" part is load-bearing — DeepSeek's cache identity
includes the client's header fingerprint, and replays that differ in params or
headers measurably miss (see docs/e2e-findings.md ).
Per-session & per-lineage stats, with anchor diffs
/permafrost/stats buckets usage by Claude Code session id, and tracks anchor
churn per lineage (stable system+tools ancestry) — so CC's tool-less preflight
calls don't pollute the churn signal, and multiple sessions through one proxy
st

[truncated]
