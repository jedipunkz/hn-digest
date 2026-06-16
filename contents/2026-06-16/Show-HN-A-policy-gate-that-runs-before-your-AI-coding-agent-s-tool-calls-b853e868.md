---
source: "https://sigmashake.com"
hn_url: "https://news.ycombinator.com/item?id=48558502"
title: "Show HN: A policy gate that runs before your AI coding agent's tool calls"
article_title: "SigmaShake — AI Agent Guardrails"
author: "cavalrytactics"
captured_at: "2026-06-16T19:21:02Z"
capture_tool: "hn-digest"
hn_id: 48558502
score: 1
comments: 0
posted_at: "2026-06-16T17:12:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A policy gate that runs before your AI coding agent's tool calls

- HN: [48558502](https://news.ycombinator.com/item?id=48558502)
- Source: [sigmashake.com](https://sigmashake.com)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T17:12:02Z

## Translation

タイトル: HN の表示: AI コーディング エージェントのツール呼び出し前に実行されるポリシー ゲート
記事のタイトル: SigmaShake — AI エージェントのガードレール
説明: SigmaShake は、AI エージェント ガードレールです。破壊的なツール呼び出しをブロックし、宣言的ルールを強制し、実行前にすべてのエージェント アクションを監査する決定的な 2 ミリ秒未満のガードレールです。
HN テキスト: 業界で 10 年以上の経験を持つセキュリティ エンジニアとして、私はエージェントのコーディングに関して同じ問題に何度も遭遇しました。指示は保証ではありません。 `CLAUDE.md`、`AGENTS.md`、メモリ ファイル、MCP の説明、およびツールのドキュメントにガイダンスを入れました。私はエージェントに次のようなことを明確に伝えました。 - アーキテクチャに関する質問には、リポジトリを grep する代わりにコード グラフを使用します。
- 非推奨の API や安全でないコードは使用しないでください。
- 特定のタスクには特定のツールを優先します。エージェントは依然として、驚くほど頻繁にこれらの指示を無視します。リポジトリ全体を grep したり、非推奨の API を使用したり、より優れたツールが利用可能な場合でも低速のツールを選択したりする可能性があります。その結果、プロンプトとルールがさまざまな問題を解決することに気づきました。プロンプトはモデルの動作に対する確率的な影響です。ルールは強制メカニズムです。そこで私は、エージェントとそのツールの間に位置する SSG (SigmaShake Governance) を構築しました。 SSG は、モデルにポリシーを記憶するように要求するのではなく、すべてのツール呼び出しを実行前に評価します。たとえば、次のルールは、アーキテクチャ関連のリポジトリ検索を再帰的 grep からコード グラフにリダイレクトします。
ルールルートコードベースgrep-to-graph {
trueを有効にする
優先度 80
重大度の警告
カテゴリ ツールルーティング
強制検索
IF ツールが「Grep」と等しい場合
メッセージ「アーキテクチャ、関係、依存関係に関する質問は、コードグラフ ツールに送られます。」
SUBSTITUTE "グラフ化クエリ \"<検索していたもの>\""
}
``` エージェントがアーキテクチャ qu に対して grep を試みるとき

通話はリダイレクトされます。非推奨のコードを書き込もうとした場合、コンテンツがディスクに到達する前に書き込みがブロックされ、代替 API が提案されることがあります。設計上の選択肢は次のとおりです。 - ルールはプレーン テキストであり、git バージョン管理されています。
- 強制はローカルで実行されます。
- 同じルールが、Claude Code、Codex、Cursor、Gemini、および MCP ベースのエージェント全体で機能します。
- バイパスは許可されますが、記録されます。
- 目標は、敵対的なモデルをサンドボックス化することではありません。エージェントの日常的なミスや手抜きを防ぐためです。既存のコントロールの多くは、動作が早すぎるか遅すぎることがわかりました。 - プロンプト ファイルは動作に影響を与えますが、強制はしません。
- ツールのホワイトリストは、多くの場合、全か無かです。
- コミット前フックは、ファイルが既に書き込まれた後に問題を検出します。
- ハーネス固有の権限はリポジトリと一緒に移動しません。 (チームが同じハーネスを使用していない場合はどうなりますか?) ターミナルを使用していない場合、SigmaShake デスクトップは CLI を必要としない同じガバナンス ダッシュボードです。macOS、Windows、および Linux 用に無料で直接ダウンロードできます。 Mac App Store と Microsoft Store では、マネージド インストールの料金を一度だけ支払いたい (または単にプロジェクトをサポートしたい) 場合は、自動更新され、Gatekeeper と SmartScreen のプロンプトをスキップする有料のサンドボックス ビルドとしても提供されています。
https://apps.apple.com/us/app/sigmashake-desktop/id676990115...
https://apps.microsoft.com/detail/9N2CHV3STGS4 私は数か月間、公開でこれを構築してきました。
https://twitch.tv/sigmashake
https://youtube.com/@sigmashakeinc エージェントが一貫して無視する指示は何ですか?

記事本文:
メインコンテンツにスキップ
🚀 発売週 — コード SHAKEITUP を使用すると最初の月が 50% オフ →
シグマシェイク
会社名
会社名
について
価格設定
営業担当者へのお問い合わせ
サポート
信頼とセキュリティ
セキュリティポータル
トラストセンター
セキュリティ勧告
ステータス
製品
ツール
新しいデスクトップアプリ
新しい VS コード拡張機能
ルールハブ
ルールビルダー
MCPサーバー
音声ディクテーション
フリート — チーム用
試してみる
デスクトップをダウンロード
ライブダッシュボード
SSGを試してみる
ヒーローデモ
エージェント
クロード・コード
カーソル
ジェミニ CLI
VSコード
コーデックス
リソース
学ぶ
ドキュメント
はじめに
ルールの構文
建築
内容
ケーススタディ
コミュニティ
不和
たるみ
GitHub
価格設定
SSGを試してみる→
サインイン
AI エージェントのガードレール
AI コーディング エージェントは非決定的です。 SigmaShake は、ユーザーが許可したことだけを実行するようにします。
エージェントが使用できるツールを強制します。危険なコマンドや安全でないコマンドを実行前にブロックします。すべての動作を監査します。
NPM の毎週のダウンロード
1,500
ルールを話して存在させる
視聴: ルールを話す、文字起こし後 2 ミリ秒以内にルールがコンパイルされる
ハブの毎週のダウンロード
1
SigmaShake デスクトップをダウンロード →
全プラットフォーム ↓
無料 · Windows、macOS、Linux · ユーザーごとのインストール
AI コーディング エージェント向けの第 1 位のガードレール
エージェントを保護するものを選択することは重要な決定であるため、それを測定可能にしました。重要な質問は、実際の作業を中断することなく不正エージェントを実際に収容しているガードレールはどれですか?あらゆるアプローチを通じて、9 つのエージェント ハーネスにわたって 324 の攻撃タスクを再現しました。 SigmaShake が勝利しました。以下のスコアはその程度を示しています。
チームがすでに使用しているエージェントと連携します
クロード・コード
PreToolUse フックの統合
カーソル
MCPサーバーの統合
コーデックス
MCPサーバーの統合
VSCode コパイロット
MCPサーバーの統合
ジェミニ CLI
MCPサーバーの統合
反重力
MCPサーバーの統合
Pi コーディング エージェント
MCPサーバーの統合
クロード・コード
PreToolUse フックの統合
カーソル
MCPサーバー内部

分離
コーデックス
MCPサーバーの統合
VSCode コパイロット
MCPサーバーの統合
ジェミニ CLI
MCPサーバーの統合
反重力
MCPサーバーの統合
Pi コーディング エージェント
MCPサーバーの統合
[ ライブプレビュー ] showcase.sigmashake.com
フルライブダッシュボードを開く ↗
コマンド ブリッジ · グラフ · ルール · 監査 · 承認 — フルサイズで駆動可能、インストール不要。
2 つのインストール方法
自分に合ったインストールを選択してください。
同じエンジン、同じルール。デスクトップ アプリは、ターミナルを必要としない GUI でそれをラップします。 CLI は、すでに実行しているシェル、CI、またはフック チェーンに接続します。
みんなのために
SigmaShake デスクトップ
クリックしてインストールするデスクトップ アプリ。端末も管理者も、Electron サイズのフットプリントもありません。 AI エージェントが統治されている場合、トレイのシールドが緑色に変わります。
✓ 30 秒のインストール · ユーザーごと · 管理者/UAC/sudo なし
✓ Windows: ~96 MB NSIS インストーラー (C#/.NET + WebView2) · macOS: ~20 MB .dmg (Swift/AppKit) · Linux: ~12 MB .tar.gz (Go/Wails)
✓ 自動アップデート: R2 デルタ経由の Windows · Sparkle 経由の macOS · R2 デルタ経由の Linux
✓ 視覚的な承認キュー + システム トレイ コントロール
✓ 同じ ssg バイナリをラップします — 同一のエンジン
個人使用の場合は無料 · Windows 10 以降 · macOS 14 Sonoma 以降 · Ubuntu 22.04 以降 / Fedora 38 以降 / Pop!_OS
curl、npm、docker、または PowerShell を介した 1 行のインストール。これを Claude Code、Cursor、Gemini CLI、VS Code、CI、または任意の PreToolUse フック チェーンに接続します。
✓ 永続的な Unix ソケット経由のデーモン モードでの 0.39 ミリ秒 p50 (測定: ssg 0.29.156、n=500 ; p99 は 2 ミリ秒未満)
✓ ワンライナーインストール —curl / npm / docker / pwsh
✓ スクリプト可能、CI フレンドリー、決定論的な終了コード
✓ すぐに使用できる 8 個以上の AI エージェント アダプター
✓ 100,000 ルールまでのフラットスケーリング · ~43 MB RSS
インストールスクリプトを表示する
ドキュメントを読む ↗
クローズドソースのバイナリ · ローカルで実行 · 製品分析のオプトイン · 匿名の使用状況カウンター自動 (IT レポートを参照)

パケットを表示) · スターターは永久無料 · npm、docker、curl、または PowerShell 経由の 1 行インストール
編集者向けにも
VS Code、Cursor、および任意の Open VSX エディター用の SigmaShake SSG
保留中の承認、ルール、監査ログがサイドバーに表示されます。埋め込みダッシュボードパネル。 .rules 構文の強調表示とスニペット。無料。
→
AI の初心者
アクセルを踏む前にブレーキをかけてください。
AI エージェントは強力です。つまり、AI エージェントは事前に実際の損害を引き起こす可能性があることを意味します。
何が起こったのかわかります。 SSGは最初に装着するシートベルトです。
AI エージェントは、開発者のプロジェクト フォルダー全体を消去しました。
「一時ファイルをクリーンアップしています。」あってはならないパターンと一致しました。そこに
元に戻すことはできませんでした。
AI アシスタントが一晩で 400 ドルのクラウド請求書を使い果たした
「エラー時の再試行」中に有料 API を繰り返し呼び出します。誰も
請求書が届くまで気づきました。
AI は頼まれてもいないのに顧客リストにメールを送信しました。
それはタスクを「うまく」完了することでした。返事はすぐに届きました。
AI エージェント ガードレール — ライブ、監査済み、2 ミリ秒未満
決定論的なガードレールは、あらゆる AI コーディング エージェントからのデータ損失、機密漏洩、特権昇格を阻止します。ライブ ブロックを監視し、ルール、監査ログ、ベンチマークを詳しく調べます。
30秒。設定ゼロ。依存関係のないインストール — サードパーティのパッケージはプルされません。
チームが使用するすべてのエージェントをカバーします。
クロード・コード
カーソル
コーデックス
VSCode コパイロット
ジェミニ CLI
反重力
Pi コーディング エージェント
ポリシーは、Claude Code、Cursor、Codex、Gemini CLI、Copilot などに均一に適用されます。
すべての決定は決定的です。
ブロックを拒否します。 「ASK」プロンプトが表示されます。 FORCE は、すべて 2 ミリ秒未満で、より安全なパスに置き換えます。
シナリオを選択します。ターミナルは AI エージェントを再実行して試行します。SSG はコマンドが実行される前にインターセプトします。
優先度/重大度
競合の優先順位とアクティブな監査ログのボリュームを決定します。
生の実行

一致時に動的に適用されるアクション。
生のエージェント プロセス ペイロードをインターセプトする AST 条件。
ssg/gov · コマンドブリッジ
評価数/分
12,847
P50 レイテンシー
0.39ミリ秒
ブロック
42
レイテンシ、過去 60 秒
12:14:08 Bash を許可する · npm run build
12:14:09 ASK 編集 · src/auth.ts
0.39 ms デーモン p50 · ssg 0.29.156、n=500 · ダッシュボードの例
本物のエンジン。模擬テレメトリ。ログインもインストールも不要です。コマンド ブリッジ、グラフ、ルール、監査、承認などのすべてのパネルをブラウザーでフルサイズで表示します。
01 コマンド ブリッジ — ヒーロー メトリクス + ライブ テレメトリ フィード
02 グラフ — ルールの競合 + 依存関係のビュー
03 ルール — ホバードキュメントを備えた DSL エディター
04 監査 — クエリ可能な意思決定ログ
05 承認 — 保留中の保留決定キュー
AMD Threadripper 3990X · 128 スレッド
64 GB DDR4 · 評価ごとに ~43 MB RSS
Bun 1.3.11 · Linux x64
▶
CLIモード
評価ごとのプロセス生成
95.8ミリ秒
待ち時間の中央値 (p50)
p50
95.8ミリ秒
p95
139.3ミリ秒
p99
166.5ミリ秒
RSS の処理 ~43 MB
起動オーバーヘッド ~73 ミリ秒 (バンのスポーン)
決定 120 許可 60 ブロック 20 ログ
⚡
デーモンモード
永続的な Unix ソケット
0.39ミリ秒
待ち時間中央値 (p50) · ssg 0.29.156、n=500
～246×
CLIより速い
p50
0.39ミリ秒
p95
0.76ミリ秒
p99
1.42ミリ秒
反復回数 1,000
メモリ内のルール事前ロード、ファイル I/O なし
プロトコル Unix ソケット (TCP なし)
スケーリング: レイテンシとルール数
1,000 回の反復 + 階層ごとに 50 回のウォームアップ · Unix ソケット RTT が含まれる · 各階層でデーモン p50 サブミリ秒 · p99 は 2 ミリ秒未満にとどまる (差異 ±0.5 ミリ秒)
CLI 遅延は Bun プロセスの起動によって支配されます (約 73 ミリ秒)。ルール数は 500 ルールまではほとんど変わりません。 100K ルールでは、ルール ファイル I/O により CLI が最大 276 ミリ秒に達します。デーモン p50 は、0.1 ～ 0.4 ミリ秒の帯域で平坦なままです (ミリ秒未満の測定ノイズが支配的です。階層の順序は意味がありません)。エンジン p99 は、最大 100,000 ルールまでの各層で 20 μs 未満に留まります。エンドツーエンドデーモン p99 (U を含む)

nix ソケット RTT) は 2 ミリ秒未満にとどまります。
エンジニアリング チームに参加して、スタック内のすべてのエージェントにわたって、AI エージェントのインシデント リスクを決定論的に削減します。
SigmaShake vs プロンプトエンジニアリング vs サンドボックス実行
本当の違いはベンチマークの数値ではなく、どこでチェックが実行されるかです。 SigmaShake はローカルファーストであり、マシン上で実行する前にツール呼び出しをゲートします。 LLM 出力フィルター (Lakera、Guardrails AI、NVIDIA NeMo Guardrails) は、エージェントがすでに行動を決定した後に、モデルの応答を検査します。 Microsoft は、SSG と並行して評価する価値のある、ポリシー適用のための MIT ライセンスのエージェント ガバナンス ツールキットも無料で提供しています。以下の 2 つの表の列は、古いインフラストラクチャ レベルのアプローチ (DIY プロンプトとサンドボックス コンテナー) を示しています。これらの出力フィルターとツールキットのオプションは別個の補完的なレイヤーであり、スコア付きの列としては表示されません。
SHAKEDOWN — エージェント封じ込めベンチマーク
誰でも自分のガードレールは安全であると主張できます。 SHAKEDOWNがそれを証明しています。厳選されたものを再プレイします
破壊的、永続的、資格情報へのアクセス、防御回避、およびサプライ チェーン ツール呼び出しのコーパス - それぞれ
MITRE ATT&CK テクニックにマッピング — サポートされているすべてのエージェント ハーネスの下で SSG を介して、どのくらいのスコアを獲得するか
正当な開発者の作業を放置したままブロックします。完全にローカル、GPU、トークンコストなし、オープンな方法論。
コミュニティ ルール、インストール準備完了
TypeScript、Python、Rust、セキュリティ、Docker などの事前構築済みルールセット。
認定され、バージョン管理され、1 つのコマンドでインストール可能です。
無料で始めましょう。さらに必要な場合は、毎月または毎年アップグレードしてください。
月額プランおよび年間プランの 14 日間の無料トライアルでは、すぐに 192 ドルが請求されます。トライアルなし
$20/月 → $10/月 — チェックアウト時にコードを使用する
毎月 5,000 件のツール評価
コミュニティ ハブへのアクセス - パブリック ルールセットのみ
承認ダッシュボード (ローカルホスト)
ローカルでの CLI

— 評価、lint、フォーマット
AI エージェントに何が許可されているかを証明する必要があるチーム向け。
プライベート ルールセット - 自分のプライベート GitHub リポジトリからインストールするか、チームのハブにプライベートに公開します
クラウド監査同期 - 署名済み、セキュリティレビュー用にエクスポート可能
チームポリシーの共有と SSO / SAML
応答時間 SLA (SOW で定義された時間) による優先サポート
ssg および SigmaShake Desktop のライセンス版のソース コード アクセス — 購入または拘束力のある購入契約に署名した顧客が利用できます。そのバージョンのみを対象とします。アップデートは含まれていません。
ご要望に応じて ISO 27001 および SOC 2 準拠の成果物をご提供します
難しい質問に対する正直な答え
懐疑的なエンジニアが実際に尋ねる質問に、直接答えます。
はい、強制メカニズムはエージェント フックまたは MCP サーバー (Claude Code、Codex、Antigravity、Gemini などをサポート) を介してネイティブに統合されています。これは意図的なものです。 SigmaShake が手動のフック スクリプトに追加するものは、最初から作成する必要のないコミュニティ ルール ライブラリ、ハブからのルールを信頼できるようにするための Ed25519 署名付きバンドル、行ごとの署名付き監査ログ (改ざん証拠のために各ガバナンス イベントに個別に署名)、多くのマシンにわたるフリート全体のポリシー同期、および承認とプロファイリングのためのダッシュボード UI です。 1 台のマシン上で 1 つのルールのみが必要な場合は、生のスクリプトで問題ありません。 SigmaShake は、システム全体を必要とするチーム向けです。
シェルにアクセスできる動機のある攻撃者: はい。正直なエージェントが間違いを犯したり、オートメーションの設定が間違っていたり、若手開発者が誤って破壊的なコマンドを書いてしまった場合は、いいえ。 SigmaShake は 95% のケースに対するガードレールであり、正しいことをしようとしているものの制約をすべて把握していないエージェントによる偶発的な危害を防ぎます。これはサンドボックスではなく、サンドボックスであるとは主張しません。敵対的な分離のために、次のように構成します。

OS レベルのサンドボックス (Docker、seccomp、Apple Sandbox) - 脅威の表面のさまざまな部分に対処します。
これらの製品は LLM 出力をフィルタリングします。これらの製品はモデルが応答した後に実行され、生成されたテキストが安全かどうかをチェックします。 SigmaShake はエージェント ツールの呼び出しをゲートします。アクションが実行される前に実行され、エージェントが行おうとしていることが許可されているかどうかがチェックされます。脅威モデルは補完的なものであり、競合するものではありません。 LLM 出力フィルターは、エージェントの rm -rf の実行を停止しません。ツールコールゲートになります。
ハブ上のすべてのルールセットは、配布前にコンテンツがハッシュされ、Ed25519 で署名されます。バンドルはロード時に検証されます。改ざんされている場合は実行されません。ルールセットはプレーン DSL テキストであり、インストール前に読み取ることができます。ルールセットを取得する前に、ルールセットが何を行うかを正確に監査できます: ssg Hub Inspection <ruleset-id> 。
はい。 ssg は、必須のクラウド依存関係のないローカル バイナリです。ハブはオプション (コミュニティ ルールのダウンロード用)、フリート同期はオプション (Pro+)、監査エクスポートはオプション (Pro+) です。ローカル専用モード: npm 経由でインストールし、 ssg init を実行して、エージェント フックまたは MCP サーバーを構成します。評価中にネットワーク呼び出しはありません。
Claude Code の組み込みの許可プロンプトは良いスタートですが、これらはバイナリ (許可またはスキップ) であり、すべてのアクションに注意が必要です。 SigmaShake は、生のフックからは得られない 4 つの機能を追加します: (1) スクリプトを書かずにデフォルト — スターター プリセット sh

[切り捨てられた]

## Original Extract

SigmaShake is AI Agent Guardrails — deterministic sub-2ms guardrails that block destructive tool calls, enforce declarative rules, and audit every agent action before it runs.

As a Security Engineer with over 10+ years in industry, I kept running into the same problem with coding agents: Instructions are not guarantees. I put guidance in `CLAUDE.md`, `AGENTS.md`, memory files, MCP descriptions, and tool documentation. I explicitly told the agent things like: - Use the code graph for architecture questions instead of grepping the repository.
- Do not use deprecated APIs or Unsafe code.
- Prefer specific tools for specific tasks. The agent would still ignore those instructions surprisingly often. It would grep the entire repo, use deprecated APIs, or choose a slower tool even when a better one was available. That made me realize prompts and rules solve different problems. A prompt is a probabilistic influence on model behavior. A rule is an enforcement mechanism. So I built SSG (SigmaShake Governance), which sits between the agent and its tools. Instead of asking the model to remember a policy, SSG evaluates every tool call before it executes. For example, this rule redirects architecture-related repository searches away from recursive grep and toward a code graph: ```text
rule route-codebase-grep-to-graph {
enable true
priority 80
severity warning
CATEGORY tool-routing
FORCE search
IF tool EQUALS "Grep"
MESSAGE "Architecture, relationship, and dependency questions are routed to the code-graph tool."
SUBSTITUTE "graphify query \"<what you were searching for>\""
}
``` When the agent attempts a grep for an architecture question, the call is redirected. If it attempts to write deprecated code, the write can be blocked before the content reaches disk and the replacement API can be suggested. A few design choices: - Rules are plain text and git-versioned.
- Enforcement runs locally.
- The same rules work across Claude Code, Codex, Cursor, Gemini, and MCP-based agents.
- Bypasses are allowed, but recorded.
- The goal is not to sandbox a hostile model; it's to prevent routine agent mistakes and shortcuts. I found that many existing controls operate either too early or too late: - Prompt files influence behavior but don't enforce it.
- Tool allowlists are often all-or-nothing.
- Pre-commit hooks catch problems after files have already been written.
- Harness-specific permissions don't travel with the repository. (what if your team does not use the same harness?) If you do not live in the terminal, SigmaShake Desktop is the same governance dashboard with no CLI required: a free direct download for macOS, Windows, and Linux. The Mac App Store and Microsoft Store also carry it as a paid, sandboxed build that auto-updates and skips the Gatekeeper and SmartScreen prompts, if you would rather pay once for the managed install (or just want to support the project):
https://apps.apple.com/us/app/sigmashake-desktop/id676990115...
https://apps.microsoft.com/detail/9N2CHV3STGS4 I've been building this in public for a few months.
https://twitch.tv/sigmashake
https://youtube.com/@sigmashakeinc What instructions do your agents consistently ignore?

Skip to main content
🚀 Launch week — 50% off your first month with code SHAKEITUP →
SigmaShake
Company
Company
About
Pricing
Contact Sales
Support
Trust & Security
Security Portal
Trust Center
Security Advisories
Status
Product
Tools
New Desktop App
New VS Code Extension
Rules Hub
Rule Builder
MCP Server
Voice Dictation
Fleet — for teams
Try It
Download Desktop
Live Dashboard
Try SSG
Hero Demo
Agents
Claude Code
Cursor
Gemini CLI
VS Code
Codex
Resources
Learn
Docs
Getting Started
Rule Syntax
Architecture
Content
Case Studies
Community
Discord
Slack
GitHub
Pricing
Try SSG →
Sign In
AI Agent Guardrails
AI coding agents are non-deterministic. SigmaShake makes sure yours only does what you allow.
Force the tools your agent can use. Block dangerous, unsafe commands before they run. Audit everything it does.
NPM Weekly Downloads
1,500
speak a rule into existence
Watch: speak a rule, rule compiled in <2 ms after transcription
Hub Weekly Downloads
1
Download SigmaShake Desktop →
All platforms ↓
Free · Windows, macOS, Linux · Per-user install
The #1-ranked guardrail for AI coding agents
Choosing what guards your agents is a critical decision, so we made it measurable. The question that matters: which guardrail actually contains a rogue agent without breaking real work? We replayed 324 attack tasks across 9 agent harnesses through every approach. SigmaShake wins — and the score below shows by how much.
Works with the agents your team already uses
Claude Code
PreToolUse hook integration
Cursor
MCP Server Integration
Codex
MCP Server Integration
VSCode Copilot
MCP Server Integration
Gemini CLI
MCP Server Integration
Antigravity
MCP Server Integration
Pi Coding Agent
MCP Server Integration
Claude Code
PreToolUse hook integration
Cursor
MCP Server Integration
Codex
MCP Server Integration
VSCode Copilot
MCP Server Integration
Gemini CLI
MCP Server Integration
Antigravity
MCP Server Integration
Pi Coding Agent
MCP Server Integration
[ live preview ] showcase.sigmashake.com
Open the full live dashboard ↗
Command Bridge · Graph · Rules · Audit · Approvals — driveable at full size, no install.
Two ways to install
Pick the install that fits you.
Same engine, same rules. The Desktop app wraps it in a no-terminal-required GUI; the CLI plugs into whatever shell, CI, or hook chain you already run.
For everyone
SigmaShake Desktop
Click-to-install desktop app. No terminal, no admin, no Electron-sized footprint. The shield in your tray turns green when your AI agents are governed.
✓ 30-second install · per-user · no admin/UAC/sudo
✓ Windows: ~96 MB NSIS installer (C#/.NET + WebView2) · macOS: ~20 MB .dmg (Swift/AppKit) · Linux: ~12 MB .tar.gz (Go/Wails)
✓ Auto-updates: Windows via R2 delta · macOS via Sparkle · Linux via R2 delta
✓ Visual approval queue + system tray controls
✓ Wraps the same ssg binary — identical engine
Free for personal use · Windows 10+ · macOS 14 Sonoma+ · Ubuntu 22.04+ / Fedora 38+ / Pop!_OS
One-line install via curl, npm, docker, or PowerShell. Wire it into Claude Code, Cursor, Gemini CLI, VS Code, CI, or any PreToolUse hook chain.
✓ 0.39 ms p50 in daemon mode via persistent Unix-socket ( measured: ssg 0.29.156, n=500 ; p99 under 2 ms)
✓ One-liner install — curl / npm / docker / pwsh
✓ Scriptable, CI-friendly, deterministic exit codes
✓ 8+ AI-agent adapters out of the box
✓ Flat-scaling to 100 000 rules · ~43 MB RSS
View install script
Read the docs ↗
Closed-source binary · runs locally · product analytics opt-in · anonymous usage counter automatic (see IT review packet ) · Starter free forever · one-line install via npm, docker, curl, or PowerShell
Also for editors
SigmaShake SSG for VS Code , Cursor , and any Open VSX editor
Pending Approvals, Rules, and Audit Log right in the side bar. Embedded dashboard panel. .rules syntax highlighting and snippets. Free.
→
New to AI
Put the brakes on before you hit the gas.
AI agents are powerful — and that means they can cause real damage before
you realise what happened. SSG is the seatbelt you put on first.
An AI agent wiped a developer's entire project folder while
"cleaning up temporary files." It matched a pattern it shouldn't have. There
was no undo.
An AI assistant ran up a $400 cloud bill overnight by
repeatedly calling a paid API while trying to "retry on error." Nobody
noticed until the invoice arrived.
An AI sent emails to a customer list without being asked to.
It was completing a task "helpfully." The replies came in fast.
AI Agent Guardrails — Live, Audited, Under 2ms
Deterministic guardrails intercept data loss, secret leaks, and privilege escalation from any AI coding agent — watch a live block, then dig into the rule, the audit log, and the benchmarks.
30 seconds. Zero config. Zero-dependency install — no third-party packages pulled.
Covers every agent your team uses.
Claude Code
Cursor
Codex
VSCode Copilot
Gemini CLI
Antigravity
Pi Coding Agent
Policy applies uniformly across Claude Code, Cursor, Codex, Gemini CLI, Copilot, and more.
Every decision — deterministic.
DENY blocks. ASK prompts. FORCE substitutes with a safer path — all in under 2ms.
Pick a scenario. The terminal re-plays an AI agent trying it — SSG intercepts before the command runs.
priority / severity
Determines conflict precedence and the active audit log volume.
The raw execution action dynamically enforced upon matching.
The AST condition intercepting the raw agent process payload.
ssg/gov · command bridge
EVALS/MIN
12,847
P50 LATENCY
0.39 ms
BLOCKS
42
latency, last 60s
12:14:08 ALLOW Bash · npm run build
12:14:09 ASK Edit · src/auth.ts
0.39 ms daemon p50 · ssg 0.29.156, n=500 · illustrative dashboard
Real engine. Mocked telemetry. No login, no install. Drive every panel — Command Bridge, Graph, Rules, Audit, Approvals — at full size in your browser.
01 Command Bridge — hero metrics + live telemetry feed
02 Graph — rule conflict + dependency view
03 Rules — DSL editor with hover docs
04 Audit — queryable decision log
05 Approvals — pending hold-and-decide queue
AMD Threadripper 3990X · 128 threads
64 GB DDR4 · ~43 MB RSS per eval
Bun 1.3.11 · Linux x64
▶
CLI Mode
Per-eval process spawn
95.8 ms
median latency (p50)
p50
95.8ms
p95
139.3ms
p99
166.5ms
Process RSS ~43 MB
Startup overhead ~73 ms (Bun spawn)
Decisions 120 allow 60 block 20 log
⚡
Daemon Mode
Persistent Unix socket
0.39 ms
median latency (p50) · ssg 0.29.156, n=500
~246×
faster than CLI
p50
0.39ms
p95
0.76ms
p99
1.42ms
Iterations 1,000
Rules in memory Pre-loaded, no file I/O
Protocol Unix socket (no TCP)
Scaling: Latency vs Rule Count
1,000 iterations + 50 warmup per tier · Unix socket RTT included · daemon p50 sub-ms at every tier · p99 stays under 2 ms (variance ±0.5 ms)
CLI latency is dominated by Bun process startup (~73 ms) — rule count barely moves the needle up to 500 rules. At 100K rules, rule-file I/O pushes CLI to ~276 ms. Daemon p50 stays flat in the 0.1–0.4 ms band (sub-ms measurement noise dominates — the ordering of tiers is not meaningful). Engine p99 stays < 20 µs at every tier up to 100,000 rules; end-to-end daemon p99 (including Unix socket RTT) stays under 2 ms.
Join engineering teams reducing AI-agent incident risk — deterministically, across every agent in their stack.
SigmaShake vs prompt engineering vs sandboxed execution
The real difference isn’t a benchmark number — it’s where the check runs. SigmaShake is local-first and gates the tool call before it executes , on your machine. LLM output filters (Lakera, Guardrails AI, NVIDIA NeMo Guardrails) inspect the model’s response — after the agent has already decided to act. Microsoft also ships a free, MIT-licensed Agent Governance Toolkit for policy enforcement worth evaluating alongside SSG. The two table columns below show the older infrastructure-level approaches (DIY prompts and sandboxed containers); those output-filter and toolkit options are distinct, complementary layers, not shown as scored columns.
SHAKEDOWN — the agent-containment benchmark
Anyone can claim their guardrail is safe. SHAKEDOWN proves it. We replay a curated
corpus of destructive, persistence, credential-access, defense-evasion and supply-chain tool calls — each
mapped to a MITRE ATT&CK technique — through SSG under every supported agent harness, and score how much
it blocks while leaving legitimate developer work alone. Fully local, no GPU, no token cost, open methodology.
Community Rules, Ready to Install
Pre-built rulesets for TypeScript, Python, Rust, Security, Docker, and more.
Certified, version-controlled, and installable with one command.
Start free. Upgrade when you need more, month to month or annually.
14-day free trial on Monthly · Annual plan bills $192 immediately, no trial
$20/mo → $10/mo — Use code at checkout
5,000 Tool evaluations per month
Community Hub access — public rulesets only
Approval dashboard (localhost)
CLI locally — eval, lint, format
For teams that must prove what their AI agents are allowed to do.
Private rulesets — install from your own private GitHub repos or publish privately to the Hub for your team
Cloud audit sync — signed, exportable for security review
Team policy sharing & SSO / SAML
Priority support with response-time SLA (hours defined in your SOW)
Source code access for the licensed version of ssg & SigmaShake Desktop — available to customers who purchase or sign a binding purchase commitment. Scoped to that version only; updates are not included.
ISO 27001 & SOC 2 compliance artifacts on request
Honest Answers to Hard Questions
Questions skeptical engineers actually ask — answered directly.
Yes, the enforcement mechanism integrates natively via agent hooks or MCP servers (supporting Claude Code, Codex, Antigravity, Gemini, and more) — and that's intentional. What SigmaShake adds on top of a hand-rolled hook script: a community rule library you don't have to write from scratch, an Ed25519-signed bundle so you can trust rules from the Hub, a per-row signed audit log (every governance event individually signed for tamper-evidence), fleet-wide policy sync across many machines, and a dashboard UI for approvals and profiling. If you only need one rule on one machine, a raw script is fine. SigmaShake is for teams that want the whole system.
A motivated attacker with shell access: yes. An honest agent making a mistake, a misconfigured automation, or a junior dev who accidentally wrote a destructive command: no. SigmaShake is a guardrail for the 95% case — preventing accidental harm from agents that are trying to do the right thing but might not know all your constraints. It is not a sandbox and does not claim to be. For adversarial isolation, compose it with OS-level sandboxing (Docker, seccomp, Apple Sandbox) — they address different parts of the threat surface.
Those products filter LLM output — they run after the model responds, checking whether generated text is safe. SigmaShake gates agent tool calls — it runs before the action executes, checking whether the thing the agent is about to do is allowed. The threat models are complementary, not competing. An LLM output filter won't stop an agent from running rm -rf ; a tool-call gate will.
Every ruleset on the Hub is content-hashed and Ed25519-signed before distribution. The bundle is verified at load time — if tampered with, it won't run. Rulesets are plain DSL text, readable before install. You can audit exactly what a ruleset does before pulling it: ssg hub inspect <ruleset-id> .
Yes. ssg is a local binary with no mandatory cloud dependency. The Hub is optional (for downloading community rules), the fleet sync is optional (Pro+), and the audit export is optional (Pro+). Local-only mode: install via npm, run ssg init , and configure your agent hooks or MCP server. No network calls during evaluation.
Claude Code's built-in permission prompts are a good start, but they're binary (allow or skip) and require your attention for every action. SigmaShake adds four things you can't get from a raw hook: (1) Defaults without writing a script — the Starter preset sh

[truncated]
