---
source: "https://github.com/arrowassassin/kintsugi"
hn_url: "https://news.ycombinator.com/item?id=48558325"
title: "Show HN: Kintsugi – a local-first safety net for AI agents and humans"
article_title: "GitHub - arrowassassin/kintsugi: Local-first safety layer for AI coding agents. Warns before destructive commands run, makes them reversible, records what every agent did. · GitHub"
author: "arr0wassass1n"
captured_at: "2026-06-16T19:21:22Z"
capture_tool: "hn-digest"
hn_id: 48558325
score: 1
comments: 0
posted_at: "2026-06-16T17:00:29Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Kintsugi – a local-first safety net for AI agents and humans

- HN: [48558325](https://news.ycombinator.com/item?id=48558325)
- Source: [github.com](https://github.com/arrowassassin/kintsugi)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T17:00:29Z

## Translation

タイトル: 番組の HN: 金継ぎ – AI エージェントと人間のためのローカル ファーストのセーフティ ネット
記事タイトル: GitHub - arrowassassin/kintsugi: AI コーディング エージェントのためのローカルファーストの安全層。破壊的なコマンドが実行される前に警告し、コマンドを元に戻せるようにし、すべてのエージェントの行動を記録します。 · GitHub
説明: AI コーディング エージェントのためのローカルファーストの安全層。破壊的なコマンドが実行される前に警告し、コマンドを元に戻せるようにし、すべてのエージェントの行動を記録します。 - アローアサシン/金継ぎ
HN テキスト: AI コーディング エージェントがマシン上で実際のシェル コマンドを実行するようになりました — rm -rf、git
Push --force、DROP TABLE、dd、ディスクに直接書き込みます。ほとんどいつもそれです
いいよ。そうでない場合（幻覚経路、プロンプト注入）
指示、確信を持って間違った推測など）元に戻すことはできず、後でわかります。 Kintsugi はエージェントとシステムの間に位置します。危険なものを捕まえる
コマンドを実行する前に、それを 1 つの平易な文で説明し、
破壊的なアクションはスナップショットで元に戻せ、すべてのコマンドを書き込みます
すべてのエージェントは、あなたが所有する追加専用のハッシュチェーンされたログに実行されました。ローカルファースト:
クラウドもアカウントもなく、マシンからは何も出ません。 AIに限った話ではありません。パッシブな bash/zsh レコーダー (エージェントは関与しない) は、
ユーザーが実行するすべてのコマンドは、同じ改ざん防止ログとスナップショットに対して実行されます。
破壊的なものはジャストインタイムで実行されます。つまり、「金継ぎの取り消し」により DBA のロールバックが行われます。
fat-fingered rm -rf または上書きは、ロールバックと同じ方法で上書きされます。
エージェントの。管理対象ホストでは、管理者パスワードの背後にある設定を封印できます。
デーモン側でブルートフォースロックアウトが強制されるため、エージェントまたは通常のユーザー
静かに消すことはできません。私が最も重視した設計ルール: 壊滅的な事態を阻止するという決定
コマンドは人間が作成した決定論的なルールによって作成されます。決して LLM ではありません。地元の人
モデルはコマンドを説明し、あいまいな中間関数に注意を追加することしかできません。

e;
ルールベースのブロックのロックを解除したり、ダウングレードしたりすることはできません。したがって、ブロックは
予測可能であり、巧妙なプロンプトによって回避することはできません。重要であることが判明したいくつかの点: - テキストではなく、実際のシェル構造を解析します。 2 つのパス - 高速トークナイザーと
真の bash AST パーサー (ブラシパーサー、純粋な Rust) — そしてそれにはより多くの時間がかかります
慎重な判決。これは $(...) 内に隠されたコマンドをキャッチします。ヒアドキュメント、
サブシェルと、部分文字列スキャナーがウェーブスルーする if/for/while ブロック。
echo "$(rm -rf /)" がキャッチされます。
- 慎重になると失敗します。パーサーが完全には理解できない行が保持されます。
決して安全だとは考えられていません。ゴールデン コーパスによって強制されるハード インバリアントは次のとおりです。
壊滅的危険性が安全に分類されることはありません。
- ツールごとのプラグインとしてではなく、プロセス/PATH レイヤーで動作します。ネイティブ
クロード コード、カーソル、コーデックス、クウェン、ジェミニ、コパイロット、
OpenCode — 加えて $PATH シムとその他すべての MCP サーバー、
生の bash スクリプトまたは Makefile が含まれます。 `kintsugi init` で配線します。
一つのコマンド。この分野には多くのツールがあるため、保証については正直に言いたいと思います。
過剰販売。 Kintsugi はシートベルトであり、カーネル ファイアウォールではありません。フックは、
インターセプト層 — yolo/自動承認モードのエージェント、または
絶対パスでバイナリを呼び出し、それらをバイパスできます。まさにそれが理由です
ファイルシステムウォッチャーのバックストップ: 約束は「回復不可能なものは何もない」ではなく、
「警告なしに実行されるものは何もありません。」そして、管理者ロックはエージェントまたは非 root を無効にします。
user — root は停止しません。悪意のあるものではなく、間違いを防ぐものです
root と同じユーザープロセス。それに対して敵対的評価を実行しました: 0/176 の危険なコマンドが漏洩しました
MITRE ATT&CK + GTFOBins コーパス、140 万個のファズ入力にわたって「安全」に保護されます。
1 つの実際のヒープ DoS が表面化しましたが、現在は修正されており、それ以来クラッシュは発生しておらず、安全でないものはゼロです
ブロック。毎秒

図はコミットされたテストによって再現されます。 Rust、MIT、クロスプラットフォーム (macOS/Linux/Windows) です。インストールは 1 行で、
モデルなしでもすぐに機能します。オプションのローカル GGUF は、単に画像をシャープにするだけです。
平易な英語の要約:curl -fsSL https://github.com/arrowassassin/kintsugi/releases/latest/download/install.sh |しー
# または、ソースから:
貨物取り付け金継ぎ
これは初期リリースなので、モデルがどこにあるのか教えていただきたいです。
間違っています - 安全なコマンドに対する誤った警報と、さらに重要なことに、
すり抜けられる壊滅的な命令。について何でも喜んでお答えします
ルール エンジン、AST アプローチ、または脅威モデル。

記事本文:
GitHub - arrowassassin/kintsugi: AI コーディング エージェントのためのローカルファーストの安全層。破壊的なコマンドが実行される前に警告し、コマンドを元に戻せるようにし、すべてのエージェントの行動を記録します。 · GitHub
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
矢暗殺者
/
金継ぎ
公共
通知
あなたはきっとsでしょう

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
153 コミット 153 コミット .claude-plugin .claude-plugin .github .github crates crates docs docs パッケージング/ homebrew パッケージング/ homebrew プラグイン/ aegis プラグイン/ aegis スクリプト スクリプト サイト サイト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DECISIONS.md DECISIONS.md ライセンス ライセンス README.md README.md kintsugi-admin-recorder-design.md kintsugi-admin-recorder-design.md kintsugi-design-doc.md kintsugi-design-doc.md kintsugi-enterprise-roadmap.md kintsugi-enterprise-roadmap.md kintsugi-phase0-1-tasklist.md kintsugi-phase0-1-tasklist.md kintsugi-phase2-5-designdoc.md kintsugi-phase2-5-designdoc.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントとあなたのチームが、マシンを破壊することなく迅速に行動できるようにします。
ボックス上のすべてのコマンドに対するセーフティ ネット: AI エージェントが実行される前に捕らえられます。
損傷。人間 (DBA、オペレーター) は、改ざんが明らかな監査証跡を取得します。
ワンコマンドで元に戻すことができます。
AI エージェントはコンピューター上で実際のシェル コマンド ( rm -rf 、 git Push --force 、 DROP TABLE ) を実行し、ディスクに直接書き込みます。ほとんどいつもそれで大丈夫です。その人
そうでないときは、幻覚のような道、即時に注入される指示、自信に満ちた態度などです。
間違った推測 — 元に戻すことはできず、後でわかります。
金継ぎはシートベルトです。エージェントとシステムの間に位置し、
危険なコマンドを実行する前に、それを 1 つの平易な文で説明し、
破壊的な行為は元に戻すことができ、すべての改ざんが明らかな記録を保持します。
どのエージェントもそうでした。ローカルファースト: クラウドもアカウントも不要で、マシンからは何も残りません。
ウェブサイト: https://arrowassassin.github.io/kintsugi/ ・ドキュメント: docs/
本物の金継ぎを

出力、ループ: カードを保持 → タイムラインが拒否 → ライブ TUI。 (静的
アニメーションが再生されない場合は、docs/img/ 内のフレーム。)
セキュリティスパイン: 壊滅的なコマンドを保持/拒否する決定は、
決定論的なルール。 LLM ではありません。生のコマンドは常にそのまま表示されます。
モデルは説明するだけです。イベント ログは追加専用でハッシュ チェーンされており、
マシンからは何も残りません。完全なルールについては、CLAUDE.md を参照してください。
保証: 敵対的なストレス + 脆弱性の評価
( docs/security-assessment.md ) で測定します —
0 / 176 個の危険なコマンドが MITRE ATT&CK + GTFOB を介して Safe に漏洩
コーパス、クラッシュなしの 140 万個のファズ入力 (実際のヒープ DoS が見つかり修正されました)、
既知の CVE は 0 件、安全ではない CVE は 0 件あります。すべての数値は綿密なテストによって再現されています。
事後分析ではなく、間違いが起こる前に間違いを阻止します。決定論的
ルール エンジン (サイコロを振る LLM ではありません) が何が壊滅的なものであるかを決定します。
ブロックは予測可能であり、巧妙なプロンプトによって「排除」することはできません。
すべてのエージェントとシェルで動作します。 Claude のネイティブ プレツール フック
コード、カーソル、コーデックス、Qwen、Gemini、Copilot、OpenCode — または生の bash
$PATH シム経由のスクリプト: プロセス レベルの 1 つの安全層ではなく、
脆弱なツールごとのプラグイン。 kintsugi init はそれらをすべて 1 つのコマンドで接続します。
デフォルトではリバーシブルです。 Kintsugi は破壊的な操作の前にファイルのスナップショットを作成するため、
金継ぎを元に戻すと元に戻ります。正直な約束は取り戻せないものではない――
ファイルシステムのバックストップは、傍受をすり抜けた変更もキャッチします。
プライベートで監査可能。クラウドもテレメトリもアカウントもありません。すべてのコマンド
実行されたすべてのエージェントは、あなたが所有する追加専用のハッシュチェーンされたログに着陸します - 改ざん
過去のエントリーと検証中断あり。
叫ばなければならないまでは落ち着いてください。安全なコマンドは、安全なコマンドを通過します。
ミリ秒;金継ぎは 1 回目のみ中断します

実際にあなたを傷つける可能性のあるもの。
macOS、Linux、Windows 上で動作します。インストールは 1 つのコマンドです。それはすぐに機能します
モデルもなければ、 kintsugi init 以降のセットアップもありません。
金継ぎは何が危険かをどのように判断するか
ブロックの決定は決定的で LLM フリーです。人間が作成した固定ルールです。
決してモデルを推測することはありません (モデルは説明するだけで、追加することしかできません)
注意)。それが信頼できる理由は、コマンドをどのように読み取るかです。
テキストではなく実際のシェル構造を解析します。金次はパスを2本走らせてテイクする
より慎重な判断: 高速トークナイザーと真の bash AST パーサー
( ブラシパーサー 、純粋な Rust)。 AST
pass は部分文字列の一致ができないものを確認します — コマンドの中に隠されたコマンド
置換 $(…) /バッククォート、シェルに供給されるヒアドキュメント、サブシェル、
if / for / while ブロック。したがって、echo "$(rm -rf /)" と bash <<<'rm -rf /' は
振り抜かれたのではなく、捕らえられた。これが業界標準のアプローチです。
AST 分析は、静的アナライザー (ShellCheck) とシェル ツールが使用するものであり、
これにより、正規表現/部分文字列スキャナーの文書化された障害モードが回避されます。
慎重になると失敗します。パーサーが完全に理解できない行が保持される
(曖昧)、決して安全だとは考えられていません。解析の失敗は注意を高めるだけです。
破滅的な評決を決して引き下げることはできない。黄金によって強制される厳しいルール
テスト コーパスは、致命的として分類された安全性がゼロです。
壊滅的なカテゴリは、重要な損害の種類に対応付けられます - データ
破壊 ( rm -rf 、 DROP TABLE 、 git Push --force 、 terraform destroy )、
ディスク/デバイスの書き込み ( dd 、 mkfs )、およびシークレットの読み取り ( .env 、 ~/.ssh/… ) —
MITRE ATT&CK 分類法と GTFOB の精神「この無害なバイナリでできること」
害」のカタログ。
自分で試してみてください。何も実行されません。
金継ぎテスト " cd build && rm -rf ../dist " # ⛔ CATASTROPHIC (ルール: rm:recursive)
金継ぎテスト「git sta」

タス " # ✓ 安全
kintsugi test ' echo "$(git Push --force)" ' # ⛔ — 置換内でキャッチされました
金継ぎテストは、クラス、発火したルール、何が起こるか、そして
Kintsugi が行内で認識する正確なコマンド — 実行、ログ記録、または
何でも連絡すること。
すべてのビルド フェーズが実装されます (「
kintsugi-phase0-1-tasklist.md と
金継ぎ-phase2-5-designdoc.md ):
フェーズ 0 — レコーダー: エージェントに依存しないインターセプト ( $PATH shim + クロード コード
フック + kintsugi-exec MCP サーバー) すべてのコマンドを改ざん防止機能に記録します。
ハッシュチェーンされた SQLite ログ。
フェーズ 1 — ゲート: 危険なコマンドを保持する決定論的ルール エンジン
リポジトリごとの意思決定メモリと .kintsugi.toml ポリシーを備えたワンキー承認。
フェーズ 2 — 説明と採点: 親切な Tier 2 スコアラーが平易な英語でスコアを記入します
概要と曖昧な帯域のリスク スコア (デフォルトではヒューリスティック、実際の CPU)
--features llama の背後にある GGUF 推論)。エスカレーションのみ - 決してダウングレードしません
ルールの決定。
フェーズ 3 — 元に戻す: 破壊的な操作前のスナップショット (reflink CoW + コピー)
フォールバック）と金継ぎ元に戻す/金継ぎ元に戻す --session 。
フェーズ 4 — レコーダー UI: FS ウォッチャーのバックストップとライブ ラタトゥイ タイムライン
（金継ぎツイ）。
フェーズ 5 — 起動: パニック キルスイッチ (金継ぎパニック / 金継ぎ再開)、
kintsugi init ポリッシュ、およびクロスプラットフォームのリリース ワークフロー。
エンタープライズ: 管理者ロック、パッシブ レコーダー、およびコントロール ルーム TUI
kintsugi init は、デフォルトで個人の姿勢を設定します - ゲートに加えてのみ
元に戻すことができ、管理する必要はありません。共有ホストまたは実稼働ホストで、次を実行します。
管理対象ポスチャの kintsugi init --enterprise 。コントロールを追加します。
以下:
パスワードロック設定+「停止パスワード」。金継ぎ管理者提供
管理者パスワードの背後にある設定を封印します (argon2id 検証者 +
XChaCha20-ポリ1305、

ワンタイム回復キーを使用します)。ロックして停止すると、
Kintsugi のフック解除または無効化にはパスワードが必要です。デーモン側で強制されます
チャレンジ/レスポンス経由 (パスワードがソケットを通過することはありません)
ブルートフォース ロックアウト — AI エージェントや通常のユーザーは、静かにロックアウトをオフにすることはできません。
そして、ハンマースクリプトはロックアウトされます。設定 (録音、自動開始、
enforcement 、fail-closed 、require-password-to-stop ) は強化のみです
コントロール - 壊滅的なフロアを緩めることはできません - TUI から管理されます
ロック解除時の設定パネル。正直なスコープ: これにより、エージェント/非 root ユーザーが無効になります。
強制シャットダウンをログに記録された回復可能なイベントに変換しますが、そうではありません。
root を停止します (「脅威マトリックス」を参照)
kintsugi-admin-recorder-design.md )。
自動再起動ウォッチドッグ + フェールクローズ。 kintsugi サービスのインストールは、
systemd / launchd の下のデーモンが restart-always で実行されるため、kill / pkill が再起動されます。
数秒以内に完了します。ウォッチドッグを無効にすること自体がパスワードでゲートされます。と
フェイルクローズドセットでは、到達不能なデーモンは無防備に実行されるのではなくブロックされるため、
デーモンを強制終了してもゲートを開けることはできません。
パッシブセッション記録 + リカバリー (AI エージェントなし)。金継ぎレコードインストール
bash/zsh 実行前フックを出力します (または --write ~/.bashrc はフックを
冪等のマネージド ブロック) なので、人間が実行するすべてのコマンドが同じに到達します。
改ざんが明らかな機密監査ログ — DBA/オペレーターのコンプライアンス向け。ブロックします
何もせず、デーモンが再起動されるまでスプールし、コマンドラインのシークレットを編集します。
ハッシュ化。コマンドよりも先にフックが発火するため、金継ぎスナップショットが生成されます。
破壊的な人間のコマンドをジャストインタイムで実行できるため、金継ぎのアンドゥにより、
ユーザーの破壊的なファイルシステム コマンド — rm -rf 、破壊的な上書き、
不正な mv — エージェントのロールバックと同じ方法です。 (これはファイルシステムリカバリーです:
インダ

tabase DROP / TRUNCATE /DML はファイルではないため、データベースの PITR / を使用してください。
それらのバックアップ。インタラクティブな bash/zsh セッションのみがキャプチャされます。)
金継ぎレポートには、レビュー用の破壊的なコマンドがリストされています。
本物のコントロールルーム TUI。 kintsugi tui はアニメーション化されたブランド端末を開きます
アプリ: ライブログ、バイタル上のタブ付きタイムライン/監査/レコーダービュー
ストリップ、ワンキー承認/拒否/元に戻す、ロック時のパスワードログイン、およびアプリ内
設定パネル — すべてを 1 つの画面で管理します。
木箱
役割
金継ぎ芯
共有タイプ、ルール エンジン、ポリシー、デシジョン メモリ、ハッシュ チェーン イベント ログ
金継ぎデーモン
常駐プロセス: ローカル IPC サーバー + 決定ループ
金継ぎインターセプト
$PATH シム、Claude Code フック ブリッジ、および kintsugi-exec MCP サーバー
金継ぎ
金継ぎバイナリ: init 、 status 、 stop 、 log 、 tui 、 …
金継ぎモデル
Tier-2 スコアラー: デフォルトでヒューリスティック、--features llama の背後にある実際の GGUF
金継ぎツイ
ライブラタトゥイタイムライン
インストール
一行。チェックサム検証済みのバイナリをダウンロードします (または、ソースからビルドします)
あなたのプラットフォームには何もありません)、その後、エージェントと
オプションのローカル モデル — オプションのものはすべてスキップできます。
curl -fsSL https://github.com/arrowassassin/kintsugi/releases/latest/download/install.sh |しー
貨物のほうが好きですか?カーゴインストール金継ぎ — すべてのビナをインストールします

[切り捨てられた]

## Original Extract

Local-first safety layer for AI coding agents. Warns before destructive commands run, makes them reversible, records what every agent did. - arrowassassin/kintsugi

AI coding agents now run real shell commands on your machine — rm -rf, git
push --force, DROP TABLE, dd, writes straight to disk. Almost always that's
fine. The one time it isn't (a hallucinated path, a prompt-injected
instruction, a confident wrong guess) there's no undo and you find out after. Kintsugi sits between the agent and your system. It catches the dangerous
command before it runs, explains it in one plain sentence, makes
destructive actions reversible with a snapshot, and writes every command
every agent ran to an append-only, hash-chained log you own. Local-first:
no cloud, no account, nothing leaves the machine. It's not only for AI. A passive bash/zsh recorder (no agent involved) puts
every command a person runs on the same tamper-evident log and snapshots the
destructive ones just-in-time — so `kintsugi undo` rolls back a DBA's
fat-fingered rm -rf or clobbering overwrite the same way it rolls back an
agent's. On a managed host you can seal the settings behind an admin password,
enforced daemon-side with brute-force lockout, so an agent or a normal user
can't quietly turn it off. The design rule I cared most about: the decision to block a catastrophic
command is made by deterministic rules a human wrote — never an LLM. A local
model can only explain a command and add caution to the ambiguous middle;
it can never unlock or downgrade a rule-based block. So the block is
predictable and can't be talked out of by a clever prompt. A few things that turned out to matter: - It parses real shell structure, not text. Two passes — a fast tokenizer and
a true bash AST parser (brush-parser, pure Rust) — and it takes the more
cautious verdict. That catches commands hidden inside $(...), here-docs,
subshells, and if/for/while blocks, which substring scanners wave through.
echo "$(rm -rf /)" is caught.
- It fails toward caution. A line the parser can't fully understand is held,
never assumed safe. The hard invariant, enforced by a golden corpus, is
zero catastrophic-classified-as-safe.
- It works at the process/PATH layer, not as a per-tool plugin. Native
pre-tool hooks for Claude Code, Cursor, Codex, Qwen, Gemini, Copilot,
OpenCode — plus a $PATH shim and an MCP server for everything else,
including a raw bash script or a Makefile. `kintsugi init` wires them in
one command. I want to be honest about the guarantee, because a lot of tools in this space
oversell. Kintsugi is a seatbelt, not a kernel firewall. Hooks are an
interception layer — an agent in a yolo/auto-approve mode, or a process that
calls a binary by absolute path, can bypass them. That's exactly why there's
a filesystem-watcher backstop: the promise is "nothing is unrecoverable," NOT
"nothing runs un-warned." And the admin lock defeats an agent or a non-root
user — it does not stop root. It guards against mistakes, not a malicious
same-user process with root. I ran an adversarial assessment against it: 0/176 dangerous commands leaked
to "safe" across a MITRE ATT&CK + GTFOBins corpus, 1.4M fuzz inputs — which
surfaced one real heap-DoS, now fixed, and no crashes since — and zero unsafe
blocks. Every figure is reproduced by a committed test. It's Rust, MIT, cross-platform (macOS/Linux/Windows). Install is one line and
it works immediately with no model — an optional local GGUF just sharpens the
plain-English summaries: curl -fsSL https://github.com/arrowassassin/kintsugi/releases/latest/download/install.sh | sh
# or, from source:
cargo install kintsugi
This is an early release and I'd genuinely like to be told where the model is
wrong — both false alarms on safe commands and, more importantly, any
catastrophic command that slips through. Happy to answer anything about the
rule engine, the AST approach, or the threat model.

GitHub - arrowassassin/kintsugi: Local-first safety layer for AI coding agents. Warns before destructive commands run, makes them reversible, records what every agent did. · GitHub
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
arrowassassin
/
kintsugi
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
153 Commits 153 Commits .claude-plugin .claude-plugin .github .github crates crates docs docs packaging/ homebrew packaging/ homebrew plugin/ aegis plugin/ aegis scripts scripts site site .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DECISIONS.md DECISIONS.md LICENSE LICENSE README.md README.md kintsugi-admin-recorder-design.md kintsugi-admin-recorder-design.md kintsugi-design-doc.md kintsugi-design-doc.md kintsugi-enterprise-roadmap.md kintsugi-enterprise-roadmap.md kintsugi-phase0-1-tasklist.md kintsugi-phase0-1-tasklist.md kintsugi-phase2-5-designdoc.md kintsugi-phase2-5-designdoc.md View all files Repository files navigation
Let AI agents — and your team — move fast, without wrecking the machine.
A safety net for every command on your box: AI agents are caught before they do
damage; the humans (DBAs, operators) get a tamper-evident audit trail with
one-command undo.
AI agents now run real shell commands on your computer: rm -rf , git push --force , DROP TABLE , writes straight to disk. Almost always that's fine. The one
time it isn't — a hallucinated path, a prompt-injected instruction, a confident
wrong guess — there's no undo, and you find out after.
Kintsugi is the seatbelt. It sits between the agent and your system, catches the
dangerous command before it runs, explains it in one plain sentence, makes
destructive actions reversible , and keeps a tamper-evident record of everything
every agent did. Local-first: no cloud, no account, nothing leaves your machine.
Website: https://arrowassassin.github.io/kintsugi/ · Docs: docs/
Real Kintsugi output, looping: hold card → denied timeline → live TUI. (Static
frames in docs/img/ if the animation doesn't play.)
Security spine: the decision to hold/deny a catastrophic command is made by
deterministic rules, never an LLM . The raw command is always shown verbatim;
the model only explains. The event log is append-only and hash-chained, and
nothing leaves your machine. See CLAUDE.md for the full rules.
Assurance: an adversarial stress + vulnerability assessment
( docs/security-assessment.md ) measures it —
0 / 176 dangerous commands leak to Safe across a MITRE ATT&CK + GTFOBins
corpus, 1.4M fuzz inputs with no crash (a real heap-DoS was found and fixed),
0 known CVEs, 0 unsafe . Every figure is reproduced by a committed test.
It stops the mistake before it happens — not a post-mortem. A deterministic
rule engine (not an LLM rolling the dice) decides what's catastrophic, so the
block is predictable and can't be "talked out of" by a clever prompt.
Works with every agent — and your shell. Native pre-tool hooks for Claude
Code, Cursor, Codex, Qwen, Gemini, Copilot, and OpenCode — or a raw bash
script via the $PATH shim: one safety layer at the process level, not a
fragile per-tool plugin. kintsugi init wires them all in one command.
Reversible by default. Kintsugi snapshots files before a destructive op, so
kintsugi undo brings them back. The honest promise is nothing unrecoverable —
a filesystem backstop catches even changes that slipped past interception.
Private and auditable. No cloud, no telemetry, no account. Every command
every agent ran lands on an append-only, hash-chained log you own — tamper
with a past entry and verification breaks.
Calm until it must shout. Safe commands fly through in well under a
millisecond; Kintsugi only interrupts for the ones that can actually hurt you.
Runs on macOS, Linux, and Windows. Install is one command; it works immediately
with no model and no setup beyond kintsugi init .
How Kintsugi decides what's dangerous
The block decision is deterministic and LLM-free — fixed rules a human wrote,
never a model guessing (the model only ever explains and can only add
caution). What makes it trustworthy is how it reads a command:
It parses real shell structure, not text. Kintsugi runs two passes and takes
the more cautious verdict: a fast tokenizer, and a true bash AST parser
( brush-parser , pure-Rust). The AST
pass sees what substring matching can't — commands hidden inside command
substitution $(…) /backticks, here-documents fed to a shell, subshells, and
if / for / while blocks. So echo "$(rm -rf /)" and bash <<<'rm -rf /' are
caught, not waved through. This is the industry-standard approach: real
AST analysis is what static analyzers (ShellCheck) and shell tooling use, and
it avoids the documented failure mode of regex/substring scanners.
It fails toward caution. A line the parser can't fully understand is held
(AMBIGUOUS), never assumed safe. A parse failure can only add caution — it
can never downgrade a catastrophic verdict. The hard rule, enforced by a golden
test corpus, is zero catastrophic-classified-as-safe .
Catastrophic categories map to the kinds of damage that matter — data
destruction ( rm -rf , DROP TABLE , git push --force , terraform destroy ),
disk/device writes ( dd , mkfs ), and secret reads ( .env , ~/.ssh/… ) — in
the spirit of the MITRE ATT&CK taxonomy and GTFOBins "this benign binary can do
harm" catalog.
Try it yourself — it runs nothing:
kintsugi test " cd build && rm -rf ../dist " # ⛔ CATASTROPHIC (rule: rm:recursive)
kintsugi test " git status " # ✓ SAFE
kintsugi test ' echo "$(git push --force)" ' # ⛔ — caught inside the substitution
kintsugi test shows the class, the rule that fired, what would happen, and the
exact commands Kintsugi sees inside your line — without executing, logging, or
contacting anything.
All build phases are implemented (see
kintsugi-phase0-1-tasklist.md and
kintsugi-phase2-5-designdoc.md ):
Phase 0 — Recorder: agent-agnostic interception ( $PATH shim + Claude Code
hook + kintsugi-exec MCP server) recording every command to a tamper-evident,
hash-chained SQLite log.
Phase 1 — Gate: a deterministic rule engine that holds dangerous commands
for one-key approval, with per-repo decision memory and .kintsugi.toml policy.
Phase 2 — Explain + score: a warm Tier-2 scorer fills a plain-English
summary and a risk score for the ambiguous band (heuristic by default; real CPU
GGUF inference behind --features llama ). Escalation-only — it never downgrades
a rules decision.
Phase 3 — Undo: snapshots before destructive ops (reflink CoW + copy
fallback) and kintsugi undo / kintsugi undo --session .
Phase 4 — Recorder UI: an FS-watcher backstop and a live ratatui timeline
( kintsugi tui ).
Phase 5 — Launch: the panic kill-switch ( kintsugi panic / kintsugi resume ),
kintsugi init polish, and a cross-platform release workflow.
Enterprise: admin lock, passive recorder, and the control-room TUI
kintsugi init sets up the personal posture by default — just the gate plus
reversible undo, nothing to administer. On a shared or production host, run
kintsugi init --enterprise for the managed posture , which adds the controls
below:
Password-locked settings + "password to stop." kintsugi admin provision
seals the settings behind an admin password (argon2id verifier +
XChaCha20-Poly1305, with a one-time recovery key). Once locked, stopping,
unhooking, or disabling Kintsugi requires the password , enforced daemon-side
via a challenge-response (the password never crosses the socket) with
brute-force lockout — an AI agent or a normal user can't quietly turn it off,
and a hammering script gets locked out. The settings ( recording , autostart ,
enforcement , fail-closed , require-password-to-stop ) are tightening-only
controls — none can loosen the catastrophic floor — and are managed from the TUI
settings panel when unlocked. Honest scope: this defeats an agent/non-root user
and turns a forced shutdown into a logged, recoverable event — it does not
stop root (see the threat matrix in
kintsugi-admin-recorder-design.md ).
Auto-restart watchdog + fail-closed. kintsugi service install runs the
daemon under systemd / launchd with restart-always, so a kill / pkill relaunches
it within seconds; disabling the watchdog is itself password-gated. With
fail-closed set, an unreachable daemon blocks rather than runs unguarded — so
killing the daemon can't be used to open the gate.
Passive session recording + recoverer (no AI agent). kintsugi record install
prints a bash/zsh preexec hook (or --write ~/.bashrc installs it as an
idempotent, managed block) so every command a human runs lands on the same
tamper-evident, classified audit log — for DBA/operator compliance. It blocks
nothing, spools across daemon restarts, and redacts command-line secrets before
hashing. Because the hook fires before the command, Kintsugi snapshots
destructive human commands just-in-time , so kintsugi undo can roll back a
person's destructive filesystem command — rm -rf , a clobbering overwrite, a
bad mv — the same way it rolls back an agent's. (It's a filesystem recoverer:
in-database DROP / TRUNCATE /DML aren't files, so use your database's PITR /
backups for those; and only interactive bash/zsh sessions are captured.)
kintsugi report lists the destructive commands for review.
A real control-room TUI. kintsugi tui opens an animated, branded terminal
app: tabbed Timeline / Audit / Recorder views over the live log, a vitals
strip, one-key approve/deny/undo, a password login when locked, and an in-app
settings panel — everything managed from one screen.
crate
role
kintsugi-core
shared types, rule engine, policy, decision memory, hash-chained event log
kintsugi-daemon
resident process: local IPC server + decision loop
kintsugi-intercept
the $PATH shim, Claude Code hook bridge, and kintsugi-exec MCP server
kintsugi
the kintsugi binary: init , status , stop , log , tui , …
kintsugi-model
Tier-2 scorer: heuristic by default, real GGUF behind --features llama
kintsugi-tui
live ratatui timeline
Install
One line. It downloads the checksum-verified binaries (or builds from source if
your platform has none), then walks you through wiring your agents and an
optional local model — everything optional can be skipped:
curl -fsSL https://github.com/arrowassassin/kintsugi/releases/latest/download/install.sh | sh
Prefer Cargo? cargo install kintsugi — installs all bina

[truncated]
