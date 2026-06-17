---
source: "https://github.com/blitzcrieg1/sentinel-trader-research"
hn_url: "https://news.ycombinator.com/item?id=48572422"
title: "I built an AI crypto trading bot, then proved it had no edge"
article_title: "GitHub - blitzcrieg1/sentinel-trader-research: Crypto trading research: an LLM directional bot (no edge, proven) and a delta-neutral funding-carry strategy (real, capacity-limited edge). · GitHub"
author: "blitzcrieg"
captured_at: "2026-06-17T16:20:03Z"
capture_tool: "hn-digest"
hn_id: 48572422
score: 1
comments: 0
posted_at: "2026-06-17T16:05:20Z"
tags:
  - hacker-news
  - translated
---

# I built an AI crypto trading bot, then proved it had no edge

- HN: [48572422](https://news.ycombinator.com/item?id=48572422)
- Source: [github.com](https://github.com/blitzcrieg1/sentinel-trader-research)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T16:05:20Z

## Translation

タイトル: AI 暗号通貨取引ボットを構築し、エッジがないことを証明しました
記事のタイトル: GitHub - blitzcrieg1/sentinel-trader-research: 暗号取引の研究: LLM 指向性ボット (エッジなし、実証済み) とデルタニュートラルな資金キャリー戦略 (実際の、容量が制限されたエッジ)。 · GitHub
説明: 暗号取引の研究: LLM 指向性ボット (エッジなし、実証済み) とデルタ中立の資金キャリー戦略 (実際の、容量が制限されたエッジ)。 - blitzcrieg1/sentinel-trader-research

記事本文:
GitHub - blitzcrieg1/sentinel-trader-research: 暗号取引の研究: LLM 指向性ボット (エッジなし、実証済み) とデルタ中立の資金キャリー戦略 (実際の、容量が制限されたエッジ)。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
電撃戦1
/
センチネルトレーダーリサーチ
出版

c
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
blitzcrieg1/sentinel-trader-research
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット デプロイ デプロイ docs docs Sentinel Sentinel testing testing .env.example .env.example .gitignore .gitignore DISCLAIMER.md DISCLAIMER.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
仮想通貨取引研究プロジェクト — そして私が仮想通貨取引研究プロジェクトを行ったときに何が起こったかの正直な記録
本当のエッジをテストしました。
AI 駆動の指向性ボットとしてスタート: 決定論的な市場の特徴 →
LLM 取引提案 → マルチゲート リスク エンジン → ペーパー/ライブ執行。それから私は
ほとんどのトレーディングプロジェクトが無視していることをやりました: 私は徹底的に反証しようとしました
それはお金を稼いだということ。そうではありませんでした。そこで私はペア、stat-arb、などのテストを続けました。
相互交換 — そしてそれらも反証しました。生き残った戦略は 1 つあります。
デルタニュートラルな資金調達は、一貫性の高い永久アルトコインを継続します。
このリポジトリには、両方の半分が含まれています: 全方向性ボット (エッジがなく、
その理由をお見せします) とキャリー戦略 (実際の構造的な、
ただしエッジは小さい）。重要なのは方法論です — を参照してください。
docs/METHHODOLOGY.md 。
まず免責事項.md をお読みください。教育/研究のみ。そうではない
財務上のアドバイス。デフォルトでは紙。取引には全損のリスクがあります。公開されているとおり、
キャリーエッジはバックテストとペーパーテストが行われていますが、長期間にわたって使用できることは証明されていません
実績。
ほとんどの「エッジ」は、実際のコストを考慮したりテストしたりすると消えてしまう幻想です。
サンプル外。方向性ボットの機械信号は +0.035R グロスでしたが、
現実的なスリッページと手数料を考慮すると、−0.243R。 LLM 層 (大脳と双子座、
ゲートリプレイでテストした) は損益分岐点を超えませんでした。変更前のペア/stat-arb
ああ

verfit — サンプル内では良好に見えたすべてのペアが、サンプル外では反転しました。
取引所間での資金調達の差はすでに年間約 0.4% まで裁定されています。唯一の
正直な原価計算の下で維持できたのは、資金を回収することでした。
構造的に空売りが難しく、ロングが継続的に支払いを行う永久債
ショートパンツ。デルタニュートラル (ロングスポット + ショートパープ) で開催され、おおよその支払いになります
年間 10 ～ 20%、市場中立、容量上限は 10 万ドル程度です。
モジュール
役割
センチネルキャリー
生き残った戦略。スキャナー、シミュレーター、デルタニュートラルポジション、リスクパリティブック、マネージャー、粘り強さ、年中無休のランナー
センチネルコア
スケジューラ、パイプラインオーケストレーション、ポジション、日報（指向性ボット）
センチネルデータ
CCXT市場データ、決定論的特徴、履歴キャッシュ
センチネル.ai
LLM クライアント、プロンプト、リフレクション、オフライン データセット ツール
センチネルリスク
マルチゲートリスクエンジン、サイジング、キルスイッチ
センチネル.exec
ブローカーインターフェイス、ペーパー/MEXC実行
センチネル.バックテスト
バックテスト エンジン、LLM ゲート リプレイ、コスト スイープ — 反証ツール
センチネルストア
SQLite (WAL) の永続性
センチネル管理者
テレグラム管理ボット
キャリー戦略は、単体テスト済みの純粋なコアに意図的に分割されています。
(スコアリング、サイジング、アカウンティング、スケジューリング) および薄いネットワーク/時間レイヤー
(価格フェッチ、ループ)。そのため、テスト可能で再起動しても安全です。
git clone https://github.com/blitzcrieg1/sentinel-trader-research.git
cd センチネル トレーダー リサーチ
Python -m venv .venv
# Windows: .\.venv\Scripts\Activate.ps1
# Linux/macOS: ソース .venv/bin/activate
pip install -e 。
cp .env.example .env # デフォルトはペーパーセーフです
キャリー戦略を実行します (紙、交換キーは不要、公開データのみ)。
python -m Sentinel.carry.run --capital 10000 --state data/carry_book.json -v
MEXC永久証書をスキャンし、一貫した資金を提供してバスケットをキュレートします

cy 、開きます
シミュレートされたデルタニュートラルヘッジ、8 時間ごと (00/08/16 UTC) に資金が発生します。
再起動しても存続するようにブックを永続化します。 TELEGRAM_BOT_TOKEN を追加 +
8 時間ごとのレポートの TELEGRAM_ADMIN_CHAT_ID を .env に変換します (オプション)。
方向性ボット (ペーパー) を実行します。
python -m センチネル.メイン
LLM キー ( GEMINI_API_KEY ) が必要で、紙にはダミーの MEXC キーを使用します。それ
機能しますが、エッジがないだけです。それがポイントです。
pytest # コアをキャリー: スコアリング、サイジング、デルタニュートラル アカウンティング、永続性
ラフチェック。
マイピーセンチネル/
エッジを自分で検証する (ウォークフォワード、OOS)
一貫性を重視して選別されたバスケットに対する最も鋭い反対意見は、生存/先読みです。
偏見。 Sentinel/carry/walkforward.py はそれに直面します: バスケットのバランスを再調整するたびに
過去のみのデータに基づいて選択され、次の目に見えないウィンドウでスコアが付けられ、その後
95% CI にブートストラップし、ランダム選択ベースラインと比較しました。
python -m Sentinel.carry.walkforward --scan --train 600 --test 120 --top-n 6 -v
エッジが後から考えてのみ存在した場合、サンプル外の収量は崩壊してしまいます。
ベースラインと CI はゼロにまたがります。 no-look-ahead プロパティは単体テストされています。参照
docs/METHODOLOGY.md §5.5。
そして、市場の薄いエッジでは容量がすべてであるため、capacity.py
「〜10万ドルの上限」を純利回り対想定元本曲線（平方根）に変える
市場への影響 + 参加上限 + 感応度バンド):
python -m Sentinel.carry.capacity --symbol XMR_USDT -v
導入
Linux + systemd 。ユニットは、deploy/ — Sentinel-carry.service (
Strategy) と Sentinel-trader.service (指向性ボット)。どちらも正常に動作します
小さな常時接続ボックス (Raspberry Pi または安価なミニ PC)。キャリーブックが使用するのは、
アトミック書き込みと SQLite は WAL を使用するため、電力損失によって状態が破損することはありません。
このプロジェクトとは何なのか、そしてそうではないのか
これは、小売業が

クリプトエッジは存在しますが、
すべての結果を再現するコードを使用します。
これは、構造上の理由により機能する、市場中立的なキャリーハーベスターです。
仕事と、容量が限られたささやかなリターン。
それは少ない賭け金から金持ちになる方法ではありません。 ～15%の市場中立性
リターンは素晴らしいですが、15% という小さな数字です。小さいとき
資本を考えると、どの戦略のリターンよりも貯蓄率の方がはるかに重要です。
これは財務上のアドバイスやすぐに使えるお金の機械ではありません。参照
免責事項.md 。
関連する研究 - およびこれとの違い
ファンディングキャリーは新しいものではなく、このリポジトリはそれを発見したとは主張していません。
これは暗号通貨で最も複製されたアイデアの 1 つであり、その構造的な理由です。
作品は十分に文書化されています - BitMEX 研究
資金調達は ~92% の確率でプラスであることがわかりました。多くの実装があります。
実行/検出ボット — 例:
aoki-h-jp/funding-rate-arbitrage 、
アーボット、
HLデルタ。これらは資金調達を検出します
機会を捉え、デルタニュートラルな取引を実行します。それらは便利ですが、実際には
実行エンジン: なし、ウォークフォワード検証、アウトオブサンプルを出荷
順列/選択バイアス ガード、または容量曲線。
学術 — 例:レバレッジを活用したBTCキャリーの研究
(~16%/年、Sharpe 6.1)。厳密ですが、単一資産で理論的です。
正直な「AI ボット」の書き込み — 例:
Jiri Dolejs の LLM 取引ボット、
これは独立して、LLM 信号に関して §2 と同じ結論に達します。
このリポジトリが追加するのは戦略ではなく、その周りの規律です。
生き残れない戦略（方向性、
LLM、ペア、クロスエクスチェンジ）なので、キャリーは単独で表示されません。
実行ボットがスキップする検証層 — ウォークフォワード + ブートストラップ、
選択バイアスの順列テストと容量曲線はすべて実行可能です。
正直な制限 (方法論 §6 )、
ある部品

実証されていない、または運用上のみ軽減可能である。
エッジは混雑した部分です。厳しさと率直さがポイントだ。
紙で構築およびテストされています。ここでの最も困難かつ最も価値のある結果は、
自分の考えをデータで反証することを学んでいます。
暗号取引の研究: LLM 指向性ボット (エッジなし、実証済み) とデルタ中立の資金キャリー戦略 (実際の、容量が制限されたエッジ)。
読み込み中にエラーが発生しました。このページをリロードしてください。
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

Crypto trading research: an LLM directional bot (no edge, proven) and a delta-neutral funding-carry strategy (real, capacity-limited edge). - blitzcrieg1/sentinel-trader-research

GitHub - blitzcrieg1/sentinel-trader-research: Crypto trading research: an LLM directional bot (no edge, proven) and a delta-neutral funding-carry strategy (real, capacity-limited edge). · GitHub
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
blitzcrieg1
/
sentinel-trader-research
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
blitzcrieg1/sentinel-trader-research
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits deploy deploy docs docs sentinel sentinel tests tests .env.example .env.example .gitignore .gitignore DISCLAIMER.md DISCLAIMER.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
A crypto trading research project — and an honest record of what happened when I
tested it for a real edge.
It started as an AI-driven directional bot: deterministic market features →
LLM trade suggestions → a multi-gate risk engine → paper/live execution. Then I
did the thing most trading projects skip: I tried, rigorously, to disprove
that it made money. It didn't. So I kept testing — pairs, stat-arb,
cross-exchange — and disproved those too. One strategy survived: a
delta-neutral funding carry on high-consistency altcoin perpetuals.
This repo contains both halves: the full directional bot (which has no edge ,
and I can show you why) and the carry strategy (which has a real, structural,
but small edge). The whole point is the methodology — see
docs/METHODOLOGY.md .
Read DISCLAIMER.md first. Educational/research only. Not
financial advice. Paper by default. Trading risks total loss. As published,
the carry edge is backtested and paper-tested, not proven over a long live
track record.
Most "edges" are illusions that vanish once you account for real costs or test
out-of-sample. The directional bot's mechanical signal was +0.035R gross but
−0.243R after realistic slippage and fees . The LLM layer (Cerebras and Gemini,
tested on a gated replay) did not beat break-even. Pairs/stat-arb was
overfit — every pair that looked good in-sample reversed out-of-sample.
Cross-exchange funding differences were already arbitraged to ~0.4%/yr. The only
thing that held up under honest cost accounting was harvesting funding on
perpetuals that are structurally hard to short, where longs persistently pay
shorts. Held delta-neutral (long spot + short perp), that pays roughly
10–20%/yr, market-neutral , with a hard capacity ceiling around $100k .
Module
Role
sentinel.carry
The strategy that survived. Scanner, simulator, delta-neutral position, risk-parity book, manager, persistence, 24/7 runner
sentinel.core
Scheduler, pipeline orchestration, positions, daily report (directional bot)
sentinel.data
CCXT market data, deterministic features, historical cache
sentinel.ai
LLM client, prompts, reflection, offline dataset tooling
sentinel.risk
Multi-gate risk engine, sizing, kill switch
sentinel.exec
Broker interface, paper/MEXC execution
sentinel.backtest
Backtest engine, LLM gated replay , cost sweep — the disproof tools
sentinel.store
SQLite (WAL) persistence
sentinel.admin
Telegram admin bot
The carry strategy is deliberately split into a pure, unit-tested core
(scoring, sizing, accounting, scheduling) and a thin network/time layer
(price fetch, the loop). That's why it's testable and restart-safe.
git clone https://github.com/blitzcrieg1/sentinel-trader-research.git
cd sentinel-trader-research
python -m venv .venv
# Windows: .\.venv\Scripts\Activate.ps1
# Linux/macOS: source .venv/bin/activate
pip install -e .
cp .env.example .env # defaults are paper-safe
Run the carry strategy (paper, no exchange keys needed — public data only):
python -m sentinel.carry.run --capital 10000 --state data/carry_book.json -v
It scans MEXC perpetuals, curates a basket by funding consistency , opens
simulated delta-neutral hedges, accrues funding every 8h (00/08/16 UTC), and
persists the book so it survives restarts. Add TELEGRAM_BOT_TOKEN +
TELEGRAM_ADMIN_CHAT_ID to .env for 8-hourly reports (optional).
Run the directional bot (paper):
python -m sentinel.main
It needs an LLM key ( GEMINI_API_KEY ) and uses dummy MEXC keys for paper. It
works — it just doesn't have an edge. That's the point.
pytest # carry core: scoring, sizing, delta-neutral accounting, persistence
ruff check .
mypy sentinel/
Validate the edge yourself (walk-forward, OOS)
The sharpest objection to a consistency-screened basket is survivorship/look-ahead
bias. sentinel/carry/walkforward.py confronts it: at each rebalance the basket
is picked on past-only data, scored on the next, unseen window, then
bootstrapped into a 95% CI and compared to a random-selection baseline.
python -m sentinel.carry.walkforward --scan --train 600 --test 120 --top-n 6 -v
If the edge only existed in hindsight, the out-of-sample yield collapses to the
baseline and the CI straddles zero. The no-look-ahead property is unit-tested. See
docs/METHODOLOGY.md §5.5.
And because capacity is the whole story for a thin-market edge, capacity.py
turns the "~$100k ceiling" into a net-yield-vs-notional curve (square-root
market impact + participation cap + a sensitivity band):
python -m sentinel.carry.capacity --symbol XMR_USDT -v
Deployment
Linux + systemd . Units are in deploy/ — sentinel-carry.service (the
strategy) and sentinel-trader.service (the directional bot). Both run fine on
a small always-on box (a Raspberry Pi or a cheap mini PC). The carry book uses
atomic writes and SQLite uses WAL, so a power loss won't corrupt state.
What this project is — and isn't
It is an honest, end-to-end study of whether a retail crypto edge exists,
with the code to reproduce every finding.
It is a working, market-neutral carry harvester with a structural reason
to work and modest, capacity-limited returns.
It is not a way to get rich from a small stake. A ~15% market-neutral
return is excellent — but 15% of a small number is a small number. At small
capital, your savings rate matters far more than any strategy's return.
It is not financial advice or a turnkey money machine. See
DISCLAIMER.md .
Related work — and how this differs
Funding carry is not new , and this repo doesn't claim to have discovered it.
It's one of the most replicated ideas in crypto, and the structural reason it
works is well documented — a BitMEX study
found funding is positive ~92% of the time. There are many implementations:
Execution / detection bots — e.g.
aoki-h-jp/funding-rate-arbitrage ,
ARBOT ,
HL-Delta . These detect funding
opportunities and run the delta-neutral trade. They're useful — but they're
execution engines : none ship walk-forward validation, an out-of-sample
permutation/selection-bias guard, or a capacity curve.
Academic — e.g. a leveraged BTC carry study
(~16%/yr, Sharpe 6.1). Rigorous, but single-asset and theoretical.
Honest "AI bot" writeups — e.g.
Jiri Dolejs' LLM trading bot ,
which independently reaches the same conclusion about LLM signals that §2 does.
What this repo adds is not the strategy — it's the discipline around it :
A documented disproof of the strategies that don't survive (directional,
LLM, pairs, cross-exchange), so the carry isn't presented in a vacuum.
A validation layer the execution bots skip — walk-forward + bootstrap, a
permutation test for selection bias, and a capacity curve, all runnable.
Honest limitations ( METHODOLOGY §6 ), including the
parts that are unproven or only operationally mitigable.
The edge is the crowded part; the rigor and the candor are the point.
Built and tested in paper. The hardest and most valuable result here was
learning to disprove my own ideas with data.
Crypto trading research: an LLM directional bot (no edge, proven) and a delta-neutral funding-carry strategy (real, capacity-limited edge).
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
