---
source: "https://github.com/agiwhitelist/tokdiet"
hn_url: "https://news.ycombinator.com/item?id=48563156"
title: "Tokdiet a local proxy that cuts LLM token spend ~70% without quality loss"
article_title: "GitHub - agiwhitelist/tokdiet: ccusage that shrinks the bill — without losing quality. Local proxy that meters tokens, compacts bloated LLM context, and proves quality didn't drop (−72% tokens, 0 regressions on real-model A/B). · GitHub"
author: "agiwhitelist"
captured_at: "2026-06-17T01:03:01Z"
capture_tool: "hn-digest"
hn_id: 48563156
score: 1
comments: 1
posted_at: "2026-06-16T22:31:49Z"
tags:
  - hacker-news
  - translated
---

# Tokdiet a local proxy that cuts LLM token spend ~70% without quality loss

- HN: [48563156](https://news.ycombinator.com/item?id=48563156)
- Source: [github.com](https://github.com/agiwhitelist/tokdiet)
- Score: 1
- Comments: 1
- Posted: 2026-06-16T22:31:49Z

## Translation

タイトル: Tokdiet は品質を損なうことなく LLM トークンの支出を最大 70% 削減するローカル プロキシです
記事のタイトル: GitHub - agiwhitelist/tokdiet: 品質を損なうことなく請求額を削減する取り組み。トークンを測定し、肥大化した LLM コンテキストを圧縮し、品質が低下していないことを証明するローカル プロキシ (トークンが -72%、実際のモデル A/B で回帰が 0)。 · GitHub
説明: 品質を損なうことなく、請求書を縮小するという特徴があります。トークンを測定し、肥大化した LLM コンテキストを圧縮し、品質が低下していないことを証明するローカル プロキシ (トークンが -72%、実際のモデル A/B で回帰が 0)。 - agiwhitelist/tokdiet

記事本文:
GitHub - agiwhitelist/tokdiet: 品質を損なうことなく請求額を削減する取り組み。トークンを測定し、肥大化した LLM コンテキストを圧縮し、品質が低下していないことを証明するローカル プロキシ (トークンが -72%、実際のモデル A/B で回帰が 0)。 · GitHub
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

アジホワイトリスト
/
トクダイエット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows アセット アセット ベンチ ベンチ コマンド コマンド docs docs フック フック スクリプト スクリプト src src テスト テスト .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md package-lock.json package-lock.json package.json package.json 価格.json 価格.json tokdiet.config.example.json tokdiet.config.example.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントは、同じファイル ダンプを 5 回送信するために料金を払っています。 tokdiet は、エージェントとモデル API の間に位置するローカル プロキシで、すべてのトークンを測定し、肥大化したコンテキストを削減し、答えが悪化していないことを証明します。
品質を損なうことなく、料金を削減できるという利点があります。
証拠（これが要点です）
すべての「コンテキスト オプティマイザー」はトークンをカットします。恐ろしい質問は、彼らが答えることができない質問です。
「コンテキストを切り取ったら、モデルは愚かになってしまうでしょうか?」
そこで、測定してみました。実際のモデル (MiniMax‑M3) 上の 6 カテゴリにわたる 66 タスクの A/B ベンチマーク。各タスクは 2 回実行されます (フルコンテキスト (ベースライン) と tokdiet (管理された) を介して) で、既知の答えに対して採点され、×3 を繰り返し、多数決でモデルのノイズをキャンセルします。
ベースライントークダイエット
入力トークン 507 万 → 146 万 -71%
品質 (66 タスク) 64/66 63/66 ≈ パリティ (95–97%)
───────────────────
198 ペアラン

s · LLM 判定 92% の類似性 · 2 番目のモデルで確認 (MiniMax-M2.5: −72%)
−71% トークン、ベースラインと同等の品質。本物のリクエスト、本物の採点 - 模擬ではありません。 〜 1 ～ 2 のタスク ギャップは、モデルの非決定性と、コンテキストの損失ではなく、モデルが秘密をエコーすることを拒否していることです。 tokdiet はやみくもに削除しないため、最も困難な「ジャンクに針が埋められた」ような敵対的なケースは通過します。コールド コンテキストを回復可能な方法でページングし、トピック上のすべてを保護します。自分で再現します:node bench/run.mjs (env に API キーが必要)。
請求書を示します
請求書をカットする
品質保持を証明する
目玉/費用、非難
✅
❌
❌
手動 /コンパクト 、手動剪定コンテキスト
❌
✅ (ブラインド)
❌
トクダイエット
✅
✅
✅ 測定 + 自動セーフモード
誰もが法案を見せるか、無視します。 tokdiet はそれをカットし、モデルが愚かになっていないことを証明し、そうなる可能性がある瞬間にカットを停止します。
# 1. プロキシ (およびライブ ダッシュボード) を開始します - インストールは必要ありません
npxトークダイエットスタート
# 2. エージェントに実際の API ではなくプロキシを指定する
エクスポート ANTHROPIC_BASE_URL=http://localhost:7787
エクスポート OPENAI_BASE_URL=http://localhost:7787/v1
次に、エージェント (クロード コード、カーソル、コーデックス、独自のスクリプト) を通常どおり実行します。トラフィックは tokdiet を介して流れ、計測されて圧縮され、重要なあらゆる方法で変更されずに上流に転送されます。
API キーは常に手元に残ります。 tokdiet は x-api-key / Authorization を読み取り、それらを上流に転送します。これらは SQLite に書き込まれたり、ログに書き込まれたりすることはありません。そして、それはフェールオープンです。ガバナー内でエラーが発生した場合、透過的なパススルーにフォールバックします。プロキシがリクエストを中断したり、独自の 5xx を表面化したりすることはありません。
デフォルトのポート: プロキシ 7787 、ダッシュボード 7878 。 --port / --dashboard-port でオーバーライドします。
tokdiet は独自のマーケットプレイス経由で Claude Code プラグインとして出荷されます。
/プラグイン

マーケットプレイスに agiwhitelist/tokdiet を追加
/plugin install tokdiet
プラグインが何を行うのか、そして何をしないのか。プラグインには軽量のプラグインが同梱されています
メータリングフックと /tokdiet コマンド。フックはツール呼び出しごとに実行されます
( PreToolUse + PostToolUse ) およびツールの I/O バイト サイズをログに記録します。
~/.tokdiet/tool-meter.log 。それ自体ではトークンを保存しません - プラグイン
クロード コード プロセスに ANTHROPIC_BASE_URL を設定できないため、ルーティングできません
圧縮プロキシを経由するトラフィック。
実際のトークンの節約はプロキシによって行われます。起動してクロード コードをポイントします
(これにより、トークンが約 −71% 削減されます):
npxトークダイエットスタート
import ANTHROPIC_BASE_URL=http://localhost:7787 # その後、このシェルからクロード コードを起動します
npx tokdiet レポートでいつでも従量制トークン、コスト、節約額を表示するか、実行します
これらの手順については、クロード コード内の /tokdiet を参照してください。
クロードコードで動作します (そしてそれについては慎重です)
Claude Code は主力のユースケースであり、単純な圧縮プロキシが直接踏み込む 2 つの地雷があります。 tokdiet は次の両方を処理します。
即時のキャッシュ。クロード コードは、キャッシュされたプレフィックスを cache_control でマークします。キャッシュされた入力コストは通常​​の最大 10% です。そのプレフィックスを書き換えるとキャッシュが無効になり、リクエストのコストが高くなる可能性があります。 tokdiet はキャッシュを認識します。cache_control ブレークポイントまたはその前のコンテンツには決して触れません。
拡張された思考。クロード・コードは、Anthropic が逐語的に返すことを要求する署名済みの思考ブロックを送信します。タッチするのは一瞬です 400 。 tokdiet は思考安全です。署名されたブロックや思考ブロックが表面化したり、変化したりすることはありません。
どちらも回帰テスト (tests/cc-compat.test.ts) でカバーされています。
正直さについてのメモ: ドル節約の話は、ペイパートークン API キー (MiniMax、Anthropic API、OpenAI など) に当てはまります。フラットな Claude サブスクリプションでは、トークンごとの料金を削減する必要がないため、価値は

お金ではなく、測定、予算、ライブ ダッシュボードがあります。
tokdiet はストリーミング リバース プロキシです。 SSE 応答は段階的にプロキシされるため (全体がバッファリングされることはありません)、エージェントのトークンは引き続きリアルタイムでストリーミングされます。
tokdiet (ローカルホスト:7787)
エージェント ─────────────────────────► モデル API
(クロードリクエスト ┌───────┐ ┌───────┐ ┌────┐ ┌────────┐ (人智 /
コード、─────► │インターセプター│─►│ メーター │─►│ 予算 │─►│ コンパクター │─► OpenAI /
カーソル、生キー └───────┘ └───────┘ └───┘ └──────┬──────┘ ジェミニ /
コーデックス、転送された検出カウント セッション/ │ dedup / elision / MiniMax)
…) プロバイダー、トークンの日/リポジトリ │ 要約の途中
体とコストの制限を守る ▼
バイト忠実 ┌───────────┐
対応 │ 品質ガード │
◄─────────────────────────┤shadow-eval + │
ストリームバック、トークン対トークン │ セーフモード │
┌─────────┐ └───────┬───────┘
│ストア(SQLite)│◄─────┘
│ + ダッシュボード │ テレメトリ、節約、劣化
━─────────┘
仮想メモリとしてのコンテキスト（アイデア）
ブラインドコンパクションは「削除して祈る」ことです。 tokdiet はコンテキストを仮想メモリのように扱います: ho

コンテンツ (最近のコンテンツ、ピン留めされたコンテンツ、現在の質問に関連するコンテンツ) は常駐します。コールド コンテンツ (古い、冗長) は、削除されず、回復可能なスタブとしてローカル ストアにページアウトされます。完全なブロックは ID をキーとして SQLite に保存されるため、モデルが実際に必要とするときにオンデマンドで監査したり (ロードマップ) ページングしたりすることができます。
仕組み
何をするのか
シャドウ評価
圧縮されたリクエストのサンプリングされた部分を圧縮されていないベースラインに対して再実行し、相違をスコア付けします (0 = 同一、100 = 無関係)。これは「品質が低下したか？」に答える測定です。
品質予算
許容可能な測定された劣化の上限 (qualityBudget.maxDegradationPct、デフォルトは 2%)。近づくと、コンパクターは最も安全な戦略に制限されます。
セーフモード
ローリング劣化がバジェットを超える場合、問題の戦略は (戦略ごとに) 無効になり、セーフモード イベントが発生します。品質が向上する前に節約が停止します。
圧縮戦略 (安全第一)
重複排除 — 損失なし。同じ大きなブロックが会話全体で再貼り付けされる場合は、最新のコピーをそのまま保持し、以前のコピーをポインタ マーカーに置き換えます。バイト同一のものだけでなく、ほぼ重複したファイル（数行を変更して再貼り付けしたファイル）でも機能します。
エリシオン — 回復可能。大量の古いツールの結果 (ファイル ダンプ、コマンド出力) をページ アウトし、プレビューと重要な行 (エラー、ID、KEY=VALUE 、URL、パス、番号) を保持し、リカバリのために全体を保存します。最近の結果、固定された結果、および質問に関連した結果はそのまま保持されます。
中間要約 (デフォルトではオフ) — 安価なモデルで履歴の途中を要約します。オプトインします (費用がかかります)。
tokdiet < コマンド > [フラグ] # エイリアス: td
コマンド
何をするのか
主要なフラグ
始める
プロキシ + ライブ ダッシュボードを実行する
--port 、 --dashboard-port 、 --no-dashboard 、 --config <

パス>
報告書
使用状況レポートを印刷 (またはエクスポート)
--since <日> 、 --json 、 --csv <ファイル> 、 --config <パス>
初期化
cwd 内の tokdiet.config.json を足場にする
--force
インストール-クロード-プラグイン
冪等のクロード コード メータリング フックをインストールする
--settings <パス>
構成
tokdiet init を実行して tokdiet.config.json を作成するか、 --config を使用してそれを渡します。すべてのフィールドはオプションであり、適切なデフォルトにマージされます。
アップストリーム オーバーライド (デフォルト以外の原点を指す — 例: MiniMax): TOKDIET_ANTHROPIC_UPSTREAM 、 TOKDIET_OPENAI_UPSTREAM 、 TOKDIET_GEMINI_UPSTREAM (従来の CTXGOV_*_UPSTREAM は、後方互換性のために引き続き読み取られます)。
プロキシを実行している状態で、http://localhost:7878 を開きます。これは、SSE 経由でライブ更新をストリーミングする単一の自己完結型ページです (ループバックのみ。コスト データがマシンから送信されることはありません)。
┌─ tokdiet ─------------------------------------------------------------------------------------ ● live · :7878 ─┐
│ │
│ SESSION クロードコード › 私のリポジトリ › MiniMax-M3 │
│ コンテキスト ███████████████████░░░░░░░░░░ 64% 128,402 / 200,000 トーク │
│ │
│ ┌── 今日 ───────┐ ┌── 保存（累計） ───────┐ │
│ │ 143 万トークンを送金 │ │ $12.40 ▂▃▅▆▇█ ↑ 1.07 ドル/h 節約 │ │
│ │ 364 万トークンを保存しました │ │ 360 万トークンがこのボックスから出ることはありません │ │
│ │ $0.43 を費やす │ │ 実際のトラフィックに -71.8% │ │
│ ━━━━━━━━━━━━━━━━━━━━━┘ │
│ │
│ QUALITY GUARD 測定による劣化 0.4% ┃▏▏▏▏▏▏▏▏░

░┃予算 2.0% │
│ ▂ ▂ 72 シャドウ評価セーフモード ● オン・OK │
│ │
│ 戦略リーダーボードは保存されたトークンを発射します Δ 品質 │
│ ▸ 重複除去 ███████████ 312 1.91M +0.0% │
│ ▸ 省略 ██████ 168 1.42M +0.6% │
│ ▸ MidSummarize · オフ · 0 — — │
│ │
│ BY TOOL クロードコード ██████████ $0.31 カーソル ███ $0.09 コーデックス ▍$03 │
━━━━━━━━━━━━━━━━━━━━━━┘
5 つのライブ画面: ライブ セッション 、 節約 、 品質 (劣化 + セーフモード ステータス)、 ツールおよびリポジトリ別 、 戦略リーダーボード - すべて SSE 経由でリアルタイムに更新されます。
節約効果を確認 - API キーは必要ありません
npm ビルド && ノード スクリプト/demo.mjs を実行
ループバックで疑似 Anthropic アップストリームを立ち上げ、その前で実際の tokdiet プロキシを開始し、パイプライン全体を通じて 1 つの現実的な肥大化したエージェント リクエストを送信します。

[切り捨てられた]

## Original Extract

ccusage that shrinks the bill — without losing quality. Local proxy that meters tokens, compacts bloated LLM context, and proves quality didn't drop (−72% tokens, 0 regressions on real-model A/B). - agiwhitelist/tokdiet

GitHub - agiwhitelist/tokdiet: ccusage that shrinks the bill — without losing quality. Local proxy that meters tokens, compacts bloated LLM context, and proves quality didn't drop (−72% tokens, 0 regressions on real-model A/B). · GitHub
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
agiwhitelist
/
tokdiet
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows assets assets bench bench commands commands docs docs hooks hooks scripts scripts src src tests tests .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json pricing.json pricing.json tokdiet.config.example.json tokdiet.config.example.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Your AI agent is paying to send the same file dump five times. tokdiet is a local proxy that sits between your agent and the model API, meters every token, puts your bloated context on a diet — and proves the answer didn't get worse.
ccusage that shrinks the bill — without losing quality.
The proof (this is the whole point)
Every "context optimizer" cuts tokens. The scary question is the one they can't answer:
"If I cut the context, does the model get dumber?"
So we measured it. A 66-task A/B benchmark across 6 categories on a real model (MiniMax‑M3), each task run twice — full context (baseline) vs through tokdiet (governed) — graded against the known answer , repeated ×3 and majority‑voted to cancel model noise:
baseline tokdiet
input tokens 5.07M → 1.46M −71%
quality (66 tasks) 64/66 63/66 ≈ parity (95–97%)
─────────────────────────────────────────────────────────
198 paired runs · LLM-judge 92% similarity · confirmed on a 2nd model (MiniMax-M2.5: −72%)
−71% tokens, quality on par with baseline. Real requests, real grading — not a mock. The ~1–2 task gap is model nondeterminism plus the model declining to echo a secret — not context loss; the hardest "needle buried in junk" adversarial cases pass, because tokdiet doesn't delete blindly — it pages cold context out recoverably and protects anything on‑topic. Reproduce it yourself: node bench/run.mjs (needs an API key in env).
shows your bill
cuts the bill
proves quality held
eyeballing /cost , ccusage
✅
❌
❌
manual /compact , hand-pruning context
❌
✅ (blind)
❌
tokdiet
✅
✅
✅ measured + auto safe-mode
Everyone shows the bill or cuts it blind. tokdiet is the one that cuts it and proves the model didn't get dumber — and stops cutting the moment it might.
# 1. Start the proxy (and live dashboard) — no install needed
npx tokdiet start
# 2. Point your agent at the proxy instead of the real API
export ANTHROPIC_BASE_URL=http://localhost:7787
export OPENAI_BASE_URL=http://localhost:7787/v1
Now run your agent (Claude Code, Cursor, Codex, your own script) as usual. Traffic flows through tokdiet , gets metered and compacted, and is forwarded upstream unchanged in every way that matters.
Your API key stays with you. tokdiet reads x-api-key / Authorization only to forward them upstream. They are never written to SQLite and never written to any log . And it's fail‑open : if anything inside the governor errors, it falls back to transparent passthrough — the proxy will never break your request or surface its own 5xx.
Default ports: proxy 7787 , dashboard 7878 . Override with --port / --dashboard-port .
tokdiet ships as a Claude Code plugin via its own marketplace:
/plugin marketplace add agiwhitelist/tokdiet
/plugin install tokdiet
What the plugin does — and what it doesn't. The plugin ships a lightweight
metering hook plus a /tokdiet command. The hook runs on every tool call
( PreToolUse + PostToolUse ) and logs tool I/O byte sizes to
~/.tokdiet/tool-meter.log . It does not save tokens by itself — a plugin
can't set ANTHROPIC_BASE_URL for the Claude Code process, so it can't route
your traffic through the compacting proxy.
The actual token savings come from the proxy . Start it and point Claude Code
at it (this is what gives you the ~−71% token reduction):
npx tokdiet start
export ANTHROPIC_BASE_URL=http://localhost:7787 # then launch Claude Code from this shell
View metered tokens, cost, and savings any time with npx tokdiet report , or run
/tokdiet inside Claude Code for these instructions.
Works with Claude Code (and it's careful about it)
Claude Code is the flagship use case, and it has two landmines a naive compacting proxy walks straight into. tokdiet handles both:
Prompt caching. Claude Code marks a cached prefix with cache_control ; cached input costs ~10% of normal. Rewriting that prefix invalidates the cache and can make a request cost more . tokdiet is cache‑aware — it never touches content at or before a cache_control breakpoint.
Extended thinking. Claude Code sends signed thinking blocks that Anthropic requires returned verbatim; touching one is an instant 400 . tokdiet is thinking‑safe — signed/thinking blocks are never surfaced or mutated.
Both are covered by regression tests ( tests/cc-compat.test.ts ).
A note on honesty: the dollar‑savings story applies to pay‑per‑token API keys (MiniMax, Anthropic API, OpenAI, …). On a flat Claude subscription there are no per‑token charges to cut, so the value there is metering, budgets, and the live dashboard — not dollars.
tokdiet is a streaming reverse proxy. SSE responses are proxied incrementally (never buffered whole), so your agent's tokens still stream in real time.
tokdiet (localhost:7787)
agent ─────────────────────────────────────────────────────────────► model API
(Claude request ┌───────────┐ ┌───────┐ ┌────────┐ ┌───────────┐ (Anthropic /
Code, ──────────► │interceptor│─►│ meter │─►│ budget │─►│ compactor │──► OpenAI /
Cursor, raw key └───────────┘ └───────┘ └────────┘ └─────┬─────┘ Gemini /
Codex, forwarded detect count session/ │ dedup / elision / MiniMax)
…) provider, tokens day / repo │ mid-summarize
keep body & cost limits ▼
byte-faithful ┌───────────────┐
response │ quality guard │
◄──────────────────────────────────────────────────────────┤ shadow-eval + │
streamed back, token-for-token │ safe-mode │
┌──────────────┐ └───────┬───────┘
│ store(SQLite)│◄──────────┘
│ + dashboard │ telemetry, savings, degradation
└──────────────┘
Context as virtual memory (the idea)
Blind compaction is "delete and pray." tokdiet treats your context like virtual memory : hot content (recent, pinned, relevant to the current question) stays resident; cold content (stale, redundant) is paged out to a local store as a recoverable stub — not deleted . The full block is kept in SQLite keyed by an id, so it can be audited and (roadmap) paged back in on demand when the model actually needs it.
Mechanism
What it does
Shadow‑eval
Re‑runs a sampled fraction of compacted requests against the un‑compacted baseline and scores the divergence (0 = identical, 100 = unrelated). This is the measurement that answers "did quality drop?"
Quality budget
A hard ceiling on acceptable measured degradation ( qualityBudget.maxDegradationPct , default 2% ). As you approach it, the compactor restricts itself to its safest strategies.
Safe‑mode
If rolling degradation exceeds the budget, the offending strategy is disabled (per‑strategy) and a safe-mode event fires. Savings stop before quality does.
Compaction strategies (safest‑first)
Dedup — loss‑free. When the same large block is re‑pasted across a conversation, keep the freshest copy verbatim and replace earlier copies with a pointer marker. Works on near‑duplicates too (a file re‑pasted with a few lines changed), not just byte‑identical ones.
Elision — recoverable. Page out the bulk of old tool results (file dumps, command output), keeping a preview plus the salient lines (errors, ids, KEY=VALUE , URLs, paths, numbers) and storing the full body for recovery. Recent, pinned, and question‑relevant results are kept intact.
Mid‑summarize (off by default) — summarize mid‑history with a cheap model. Opt‑in (it costs money).
tokdiet < command > [flags] # alias: td
Command
What it does
Key flags
start
Run the proxy + live dashboard
--port , --dashboard-port , --no-dashboard , --config <path>
report
Print a usage report (or export)
--since <days> , --json , --csv <file> , --config <path>
init
Scaffold tokdiet.config.json in the cwd
--force
install-claude-plugin
Install an idempotent Claude Code metering hook
--settings <path>
Configuration
Run tokdiet init to create tokdiet.config.json , or pass one with --config . All fields are optional and merge over sensible defaults.
Upstream overrides (point at a non‑default origin — e.g. MiniMax): TOKDIET_ANTHROPIC_UPSTREAM , TOKDIET_OPENAI_UPSTREAM , TOKDIET_GEMINI_UPSTREAM (legacy CTXGOV_*_UPSTREAM still read for back‑compat).
With the proxy running, open http://localhost:7878 — a single self‑contained page that streams live updates over SSE (loopback only; your cost data never leaves the machine):
┌─ tokdiet ───────────────────────────────────────── ● live · :7878 ─┐
│ │
│ SESSION claude-code › my-repo › MiniMax-M3 │
│ context ███████████████████░░░░░░░░░░ 64% 128,402 / 200,000 tok │
│ │
│ ┌── TODAY ────────────────┐ ┌── SAVED (cumulative) ─────────────┐ │
│ │ sent 1.43M tok │ │ $12.40 ▁▂▃▅▆▇█ ↑ saving $1.07/h │ │
│ │ saved 3.64M tok │ │ 3.6M tokens never left this box │ │
│ │ spend $0.43 │ │ −71.8% on real traffic │ │
│ └──────────────────────────┘ └───────────────────────────────────┘ │
│ │
│ QUALITY GUARD measured degradation 0.4% ┃▏▏▏▏▏▏▏▏░░┃ budget 2.0% │
│ ▁▁▂▁▁▁▂▁▁▁ 72 shadow-evals safe-mode ● ON · OK │
│ │
│ STRATEGY LEADERBOARD fires tokens saved Δ quality │
│ ▸ dedup ███████████ 312 1.91M +0.0% │
│ ▸ elision ██████ 168 1.42M +0.6% │
│ ▸ midSummarize · off · 0 — — │
│ │
│ BY TOOL claude-code ██████████ $0.31 cursor ███ $0.09 codex ▍$03 │
└───────────────────────────────────────────────────────────────────────┘
Five live screens: Live session , Savings , Quality (degradation + safe‑mode status), By tool & repo , and Strategy leaderboard — all updating in real time over SSE.
See the savings — no API key required
npm run build && node scripts/demo.mjs
Stands up a mock Anthropic upstream on loopback, starts the real tokdiet proxy in front of it, and sends one realistic bloated agent request through the whole pipeline —

[truncated]
