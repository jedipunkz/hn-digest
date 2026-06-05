---
source: "https://github.com/MattJackson/busbarAI"
hn_url: "https://news.ycombinator.com/item?id=48416730"
title: "Show HN: Busbar – every LLM behind one URL, in a single Rust binary"
article_title: "GitHub - MattJackson/busbarAI: Point your existing SDK at one URL and reach every LLM vendor — with real failover, not a try/except. One static Rust binary. · GitHub"
author: "mattjackson86"
captured_at: "2026-06-05T21:37:32Z"
capture_tool: "hn-digest"
hn_id: 48416730
score: 1
comments: 0
posted_at: "2026-06-05T18:58:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Busbar – every LLM behind one URL, in a single Rust binary

- HN: [48416730](https://news.ycombinator.com/item?id=48416730)
- Source: [github.com](https://github.com/MattJackson/busbarAI)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T18:58:16Z

## Translation

タイトル: Show HN: Busbar – 単一の Rust バイナリ内の、1 つの URL の背後にあるすべての LLM
記事のタイトル: GitHub - MattJackson/busbarAI: 既存の SDK を 1 つの URL に指定し、try/excel ではなく実際のフェイルオーバーを使用して、すべての LLM ベンダーにアクセスします。 1 つの静的な Rust バイナリ。 · GitHub
説明: 既存の SDK を 1 つの URL に指定し、try/excel ではなく実際のフェイルオーバーを使用して、すべての LLM ベンダーにアクセスします。 1 つの静的な Rust バイナリ。 - マットジャクソン/busbarAI
HN テキスト: 最近、AI エンドポイント (ローカルで実行するものを含む) に関連する複数のプロジェクトに取り組んでいますが、複数のエンドポイント間で簡単に負荷分散する方法が必要であることがわかりました。時々、オンプレミスで読み込みを処理できず、その時のクレジットの状況に応じて、z.ai の使用量または Anthropic を増やす必要がありました。あることが別のことにつながり、私は最終的に Busbar: Rust で書かれた LLM ゲートウェイを書くことになりました (私は最近 Rust に興味があります)。既存の OpenAI/Anthropic/Gemini SDK を指定し、モデルをプール名に変更すると、その名前がベンダー間で負荷分散されるようになります。クライアントのコードは変更されず、それが起こったことさえ学習しません。私の中心的な考え方は「プロバイダーではなくプロトコル」です。私は、Anthropic、OpenAI、Gemini、Bedrock、Responses、Cohere の 6 つのプロトコルをロスレスで実装しています。プロバイダーは YAML の 3 行で定義し、主にプロバイダーが通信するプロトコルを指定します。クライアントはバスバーにプロトコルを送信し、バスバーはプロバイダーにプロトコルを送信します。 - 各プロトコルは、ストリーミングおよびバッファリングされたリクエストとレスポンスを両方向に変換します。同じプロトコルの呼び出しはそのまま通過します。クロスプロトコル呼び出しは、ぎこちなさを調整します (ある方言では必須のフィールドで、別の方言ではオプションにするフィールド)。 - 障害が誰のせいであるかを認識するサーキットブレーカー。本当に障害が発生しているバックエンドへのルーティングを停止しますが、ペナルティは発生しません

単純に大きすぎるリクエストに対してモデルを作成し (代わりに、より大きなコンテキスト モデルで再試行します)、呼び出し元が不正なリクエストを送信したときにバックエンドを非難しません。健全なモデルは、そのモデルのせいではない理由でローテーションから除外されることはありません。私が個人的に直面した問題はすべて、バスバーでは 1 回、10 回のアプリケーションでは 10 回解決したいと思っていました。 - 手作業で AWS 実装を実装したため、AWS SDK に依存していません。SigV4 と、Bedrock It 1.0.0-rc.2 用の最初から作成した AWS イベントストリーム フレーム デコーダ — 機能が完全で API が安定しており、1.0.0 より前にリリース候補の検証が進行中です。私は自分のプロジェクトでそれを使用しており、問題をうまく解決しています。ソロプロジェクト、AGPL-3.0。 AGPL の選択については議論の余地があります。 request-path コンポーネントにとってそれが重要であることはわかっています。特にエッジケースで翻訳に損失が残る可能性がある箇所についてのフィードバックは大歓迎です。貢献と会話を求めています!

記事本文:
GitHub - MattJackson/busbarAI: 既存の SDK を 1 つの URL に指定し、try/excel ではなく実際のフェイルオーバーですべての LLM ベンダーにアクセスします。 1 つの静的な Rust バイナリ。 · GitHub
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
マットジャクソン
/
バスバーAI
公共
通知
Cha にサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
145 コミット 145 コミット .github .github docs docs src src .editorconfig .editorconfig .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンスREADME.md README.md SECURITY.md SECURITY.md config.yaml config.yaml Providers.yaml Providers.yaml Rustfmt.toml Rustfmt.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
既存の SDK を 1 つの URL に指定し、 try/excel ではなく実際のフェイルオーバーを使用して、すべての LLM ベンダーにアクセスします。
あなたのコードはすでに OpenAI (または Anthropic、または Gemini) を話します。ベース URL を Busbar に変更すると、model: "fast" がプールになります。Claude、GPT、および Gemini が 1 つの名前の背後にあり、重みによって負荷分散され、ベンダーの機能が低下した場合にはリクエスト中のフェイルオーバーが行われます。アプリケーション コードは、そのようなことが起こったことを学習することはありません。
- クライアント = OpenAI(api_key=OPENAI_KEY)
+ クライアント = OpenAI(api_key=BUSBAR_TOKEN、base_url="http://busbar:8080")
client.chat.completions.create(
- モデル = "gpt-4o-mini",
+ model="fast", # プール: 80% クロード、20% GPT-4o-mini、フェイルオーバー時の Gemini
メッセージ=[{"役割": "ユーザー", "コンテンツ": "こんにちは!"}],
)
そのリクエストは OpenAI として残され、Anthropic によって処理され、OpenAI として戻ってきた可能性があります。双方向にロスレスで変換されます。 Anthropic が飛行中に 429 を返した場合、Busbar はクライアントがバイトを確認する前に別のプール メンバーに再ルーティングしたでしょう。 1 つのベンダーの悪い日が停電につながることはなくなります。
これは 1 つの静的な 7.4 MB バイナリです。Python サイドカー、ランタイム、インタープリター、依存関係ツリー、リクエスト パスでの GC 一時停止はありません。バインドして 15 ミリ秒以内にサービスを開始します。それをダウンロードし、2 つの YAML ファイルを指定して実行します。
名前は配電から: a bu

sbar は 1 つのフィードを取得し、ブレーカーが切断された多数の回路にわたってファンアウトします。つまり、1 つのエントリ ポイント、重み付けされた分散、回路ごとの保護が行われます。まさにこの形ですね。
モデルは依存関係ではなく構成値になります。 GPT から Claude への切り替え、または GPT 間でトラフィックを 80/20 に分割することは、 config.yaml の編集であり、コードの変更やデプロイではありません。アプリにはベンダーごとの SDK はありません。
フェイルオーバーは実際のサブシステムであり、例外ブロックではありません。プール全体にわたる重み付けされたスムーズなラウンドロビン、指数関数的なクールダウンとシングルフライトのハーフオープン回復を備えた (プール、レーン) ごとのサーキット ブレーカー、デッドラインおよびホップ キャップ付きのフェイルオーバー、セッション アフィニティ、およびより大きなコンテキスト メンバーへのオーバーサイズ要求のフェイルオーバー。上流の Retry-After を尊重します。最初のバイトがクライアントに到達するまでストリーミング応答を再ルーティングします。
翻訳はロスレスで、難しい部分は処理されます。同じプロトコルの呼び出しはそのまま通過します。cache_control、思考ブロック、引用、およびネイティブ使用量アカウンティングはすべて存続します。クロスプロトコル呼び出しは、方言の非対称性も調整するスーパーセット IR を通過します。Anthropic (必要な) にルーティングされた max_tokens のない OpenAI リクエストは、拒否ではなく設定されたデフォルトを取得し、呼び出し元が指定した値は常に保持されます。 f32 の往復は 0.7 を静かに 0.699999988 に変えるため、偶数の温度は f64 として伝えられます。
それは正しい失敗にペナルティを与えます。バックエンドは、アップストリームの障害 (5xx、過負荷、レート制限、請求/割り当て、認証) に対しては排除されますが、クライアント 4xx に対しては排除されません。発信者が不正な形式または大きすぎるリクエストを送信したために、正常なレーンがデッドとしてマークされることはありません。ほとんどのゲートウェイはこれを誤解しています。
リクエストのパスが退屈なのは、意図的です。 SSRF セーフ (宛先は精査されたカタログからのみ取得され、リクエスト データからは取得されません)、constant-ti

me トークンの比較、SHA-256 でハッシュされた仮想キー、制限されたリクエスト本文、完全にパラメーター化された SQL、ログには決して触れないシークレット。
すでにオプションを検討している場合:
2 つの SDK クライアントを超える手動の Try/Except と比較します。これにより、フェイルオーバーではなくフォールバックが得られます。ヘルス トラッキング、重み付け、ブレーカー、ストリーミング境界はなく、ベンダーを追加するたびに新しいブランチが作成されます。 Busbar では、ベンダーを 3 行の YAML で追加できます。
対 Python ベースのゲートウェイ (LiteLLM など) — Busbar は、インタープリター、virtualenv、ホット パス内の GC を持たない単一のネイティブ バイナリです。信頼性プリミティブ (SWRR、2 段階ブレーカー、コンテキスト長フェイルオーバー) は、アドオンではなくファーストクラスです。
vs. ホスト型ルーター (OpenRouter など) — Busbar はベンダー キーを使用してインフラストラクチャ内で実行されます。トラフィックやプロンプトについてはネットワークから何も出ず、プロバイダーに直接料金を支払います。
テーマ: プロバイダーではなくプロトコル。少数のワイヤ プロトコルをロスレスで実装すると、そのプロトコルを話すすべてのベンダーは単なるカタログ エントリ (名前、base_url 、およびそのキーを保持する環境変数) にすぎません。 6 つのプロトコルが実装されています。 42 の精査されたプロバイダーがカタログ エントリとして出荷され、OpenAI 互換のエンドポイント (独自のものを含む) は 3 行の YAML です。
ステータス: 1.0.0-rc.2 — 機能が完成しており、API が安定しています。 1.0.0 に先立ってリリース候補の検証が進行中です。 AGPL-3.0。
リリース ページからプラットフォーム (Linux、macOS、Windows - Intel および ARM) のリリースを取得するか、ソースからビルドします。
カーゴビルド --release # → ターゲット/リリース/バスバー
2. 設定する
Busbar は 2 つの YAML ファイルを読み取ります。 Provides.yaml は出荷されたカタログです (プロトコル、base_url 、プロバイダーごとのエラー マップ - ほとんど触ることはありません)。 config.yaml がデプロイメントです。キーは構成に書き込まれることはなく、キーを保持する環境変数の名前のみが書き込まれます。

${VAR} はロード時に展開され、参照変数が設定されていないと起動に失敗します。
聞く：「0.0.0.0:8080」
認証:
モード : トークン
client_tokens : ["${BUSBAR_CLIENT_TOKEN}"]
プロバイダー:
anthropic : { api_key_env: ANTHROPIC_KEY }
openai : { api_key_env: OPENAI_KEY }
モデル:
クロード・ソネット : { プロバイダー: anthropic、最大同時実行数: 20 }
gpt-4o-mini : { プロバイダー: openai、max_concurrent: 50 }
プール :
速い：
メンバー：
- { ターゲット: クロード・ソネット、重み: 8 }
- { ターゲット: gpt-4o-mini、重み: 2 }
3. 走る
import BUSBAR_CLIENT_TOKEN=changeme ANTHROPIC_KEY=sk-ant-... OPENAI_KEY=sk-...
BUSBAR_PROVIDERS=./providers.yaml BUSBAR_CONFIG=./config.yaml ./target/release/busbar
4. 電話してみよう
クライアントは、SDK とまったく同じように、プロトコル パスを自分で追加します。 Anthropic 形式のクライアントが 1 つのモデルをヒットする:
カール -s http://localhost:8080/claude-sonnet/v1/messages \
-H " 認可: Bearer $BUSBAR_CLIENT_TOKEN " -H " コンテンツタイプ: application/json " \
-d ' {"モデル":"無視","max_tokens":256,"メッセージ":[{"役割":"ユーザー","コンテンツ":"Hello!"}]} '
モデルがクロスプロトコル高速プールを選択する OpenAI 形式のクライアント:
カール -s http://localhost:8080/v1/chat/completions \
-H " 認可: Bearer $BUSBAR_CLIENT_TOKEN " -H " コンテンツタイプ: application/json " \
-d ' {"モデル":"高速","メッセージ":[{"役割":"ユーザー","コンテンツ":"Hello!"}]} '
バスバーは、リクエストのモデルを選択されたメンバーに書き換えて、プロバイダーの資格情報を挿入します。呼び出し元自身のモデルとキー フィールドは無視されます (呼び出し元のキーが上流に転送されるパススルー認証モードを除く)。
バスバーのスコープは、プロバイダー数ではなく、プロトコル数 (6) です。各プロトコルは、クライアントが話す形式、バックエンドが話す形式、またはその両方であるファーストクラスの入力および出力です。
ストリーミングは 6 つすべてにおいて最高級です: Gemini via :streamGenerat

eContent?alt=sse 、Bedrock では、バイナリの application/vnd.amazon.eventstream フレームをデコードし、呼び出し元のプロトコルとして再フレーム化し、残りは SSE 経由で行います。
ルート
ターゲット
POST /<名前>/v1/messages
人間の侵入。 <name> はモデル ( /claude-sonnet ) またはプール ( /fast ) です。
POST /<プロバイダ>/<モデル>/v1/messages
1 つのプロバイダー + モデルへのアドホックな直接ルート、プールは不要
POST /v1/chat/completions
OpenAI イングレス。ボディのモデルがモデルまたはプールを選択します
GET /stats · GET /healthz · GET /metrics
レーンごとのヘルス (JSON) · ライブネス · Prometheus
/admin/keys (POST/GET/DELETE)
仮想キー管理 (ガバナンスのみ)
クロスプロトコル変換は 3 つのターゲティング モードすべてに適用されます。入力プロトコルはルートによって固定され、選択されたレーンが他のものを話している場合、バスバーは IR を介して変換します。
特徴
何をするのか
プールと重み付け
レーン間のスムーズな重み付きラウンドロビン。同時実行の上限が 1 つの集合体にスタックされる
フェイルオーバー
プールごとの期限 + ホップ キャップ + メンバーの除外、直接ルート、アドホック ルート、プールされたルート全体に適用される
枯渇政策
拒否 / status_503 / minimum_bad / fallback_pool:<名前>
サーキットブレーカー
2段階の分類→処分。 error_rate または連続した旅行。指数関数的なクールダウン。シングルフライトのハーフオープン回復。名誉リトライアフター
セッションアフィニティ
メンバーが正常な状態を維持している間のヘッダーによるスティッキー ルーティング
コンテキスト長フェイルオーバー
サイズ超過のリクエストは、小さいレーンにペナルティを与えることなく、より大きいコンテキスト メンバーにフェイルオーバーします。
健康状態の調査
プロバイダーごとに、なし / デッド / アクティブなバックグラウンド プローブ
ガバナンス
仮想キー、許可されたプール ACL、トークン精度のバジェット、RPM/TPM 制限 - 組み込み SQLite で永続的、デフォルトでオフ
可観測性
Prometheus /metrics 、オプションの OTLP トレース、オプションのリクエストログ Webhook
フィールドごとの完全な設定リファレンス (デフォルトと動作済み)

例は docs/configuration.md にあります。アーキテクチャとプロバイダではないプロトコルの理論は docs/architecture.md と docs/roadmap.md にあります。
Busbar はリクエスト パス内に配置されるため、そこで目立たないように構築されています。呼び出し元によって制御されるアップストリームはありません (宛先は精査された Providers.yaml から取得され、リクエスト データから取得されることはありません。そのため SSRF 安全です)、クライアントと管理トークンの定時比較、SHA-256 ハッシュとしてのみ保存される仮想キー、32 MiB に制限されたリクエスト本文 (オープンリレー モードでも)、完全にパラメータ化されたガバナンスSQL、およびプロバイダーのキー/トークン/本体はログに記録されません。これらはテスト スイート (267 テスト) によって実行され、専用の強化パスの焦点でした。脆弱性を報告するには、 SECURITY.md を参照してください。
単一の Rust バイナリ、安定したツールチェーン (エディション 2021)。 CI は、Linux、macOS、および Windows 上でビルドおよびテストします。リリースには、x86_64 / aarch64 Linux、Intel/Apple-Silicon macOS、および x86_64 Windows が含まれています。
カーゴビルド --release
貨物試験
カーゴクリッピー --all-targets -- -D 警告
貢献とライセンス
貢献を歓迎します — COTRIBUTING.md を参照してください。 AGPL-3.0 以降のライセンス ( LICENSE )。通常、バスバーはネットワーク サービスとして実行されるため、AGPL の §13 ネットワーク使用条項が適用されます。変更されたバスバーを実行し、他のユーザーがそれにアクセスできるようにします。

[切り捨てられた]

## Original Extract

Point your existing SDK at one URL and reach every LLM vendor — with real failover, not a try/except. One static Rust binary. - MattJackson/busbarAI

I have been working on multiple projects lately involving AI endpoints (including some I run locally) and I found I needed a way to easily load balance across multiple. Sometimes my on-prem would not be able to handle to load and Id have to crank up the z.ai usage or Anthropic depending on where my credits were at the time. One thing led to another and I ended up writing Busbar: An LLM gateway, written in Rust (I have a thing for Rust lately). You point your existing OpenAI/Anthropic/Gemini SDK at it, change the model to a pool name, and that name now load-balances across the vendors. Your client code doesn't change and never learns it even happened. My central idea is "protocols, not providers". I implement six protocols - Anthropic, OpenAI, Gemini, Bedrock, Responses, Cohere - losslessly. You define a provider in three lines of YAML, mainly specifying the protocol that provider speaks. Your client speaks a protocol in to Busbar and Busbar speaks a protocol out to the provider. - Each protocol translates request and response, streamed and buffered, in both directions. Same-protocol calls pass through untouched; cross-protocol calls reconcile the awkwardness (a field one dialect requires and another makes optional). - A circuit breaker that knows whose fault a failure is. It stops routing to a backend that's genuinely failing, but it won't penalize a model for a request that was simply too big (it retries on a larger-context model instead), and it won't blame a backend when the caller sent a bad request. A healthy model never gets pulled from rotation for something that wasn't its fault. All issues I have personally faced and wanted to fix one time in busbar vs 10x in 10 applications. - Hand-rolled AWS implementations so I am not reliant on AWS SDK's: SigV4 and a from-scratch AWS eventstream frame decoder for Bedrock It's 1.0.0-rc.2 — feature-complete and API-stable, with release-candidate validation underway before 1.0.0. I have been using it on my projects and its solving my problems nicely. Solo project, AGPL-3.0. The AGPL choice is open to discussion; I know it matters for a request-path component. Feedback very welcome, particularly on where the translation might still be lossy in edge cases. Contribution and conversation desired!

GitHub - MattJackson/busbarAI: Point your existing SDK at one URL and reach every LLM vendor — with real failover, not a try/except. One static Rust binary. · GitHub
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
MattJackson
/
busbarAI
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
145 Commits 145 Commits .github .github docs docs src src .editorconfig .editorconfig .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md config.yaml config.yaml providers.yaml providers.yaml rustfmt.toml rustfmt.toml View all files Repository files navigation
Point your existing SDK at one URL and reach every LLM vendor — with real failover, not a try/except .
Your code already speaks OpenAI (or Anthropic, or Gemini). Change the base URL to Busbar, and model: "fast" becomes a pool — Claude, GPT, and Gemini behind one name, load-balanced by weight, with mid-request failover when a vendor degrades. Your application code never learns that any of it happened.
- client = OpenAI(api_key=OPENAI_KEY)
+ client = OpenAI(api_key=BUSBAR_TOKEN, base_url="http://busbar:8080")
client.chat.completions.create(
- model="gpt-4o-mini",
+ model="fast", # a pool: 80% Claude, 20% GPT-4o-mini, Gemini on failover
messages=[{"role": "user", "content": "Hello!"}],
)
That request left as OpenAI, may have been served by Anthropic, and came back as OpenAI — translated losslessly both ways. If Anthropic had returned a 429 mid-flight, Busbar would have rerouted to another pool member before your client saw a byte. One vendor's bad day stops being your outage.
It's one static 7.4 MB binary — no Python sidecar, no runtime, no interpreter, no dependency tree, no GC pauses in your request path. It binds and starts serving in under 15 ms . Download it, point two YAML files at it, run it.
The name is from electrical distribution: a busbar takes one feed and fans it out across many breakered circuits — one entry point, weighted distribution, per-circuit protection. That's exactly the shape of this thing.
The model becomes a config value, not a dependency. Switching from GPT to Claude — or splitting traffic 80/20 between them — is an edit to config.yaml , not a code change and a deploy. No per-vendor SDKs in your app.
Failover is a real subsystem, not an except block. Weighted smooth round-robin across a pool, a per-(pool, lane) circuit breaker with exponential cooldown and single-flight half-open recovery, deadline- and hop-capped failover, session affinity, and oversized-request failover to a larger-context member. It honors upstream Retry-After . It reroutes a streaming response right up until the first byte reaches your client.
Translation is lossless, and the hard parts are handled. Same-protocol calls pass through untouched — cache_control , thinking blocks, citations, and native usage accounting all survive. Cross-protocol calls go through a superset IR that also reconciles dialect asymmetries: an OpenAI request with no max_tokens routed to Anthropic (which requires one) gets a configured default instead of a rejection, and a caller-supplied value is always preserved. Even temperature is carried as f64 , because an f32 round-trip would quietly turn 0.7 into 0.699999988 .
It penalizes the right failures. A backend is ejected for upstream faults — 5xx, overload, rate-limit, billing/quota, auth — but never for a client 4xx . A healthy lane is never marked dead because a caller sent a malformed or oversized request. Most gateways get this wrong.
It's boring in your request path, on purpose. SSRF-safe (destinations come only from your vetted catalog, never from request data), constant-time token comparison, SHA-256-hashed virtual keys, bounded request bodies, fully parameterized SQL, and secrets that never touch the logs.
If you're already weighing the options:
vs. a hand-rolled try/except over two SDK clients — that gives you fallback, not failover: no health tracking, no weighting, no breaker, no streaming boundary, and a new branch every time you add a vendor. Busbar makes adding a vendor three lines of YAML.
vs. Python-based gateways (e.g. LiteLLM) — Busbar is a single native binary with no interpreter, no virtualenv, and no GC in the hot path. The reliability primitives (SWRR, the two-stage breaker, context-length failover) are first-class, not add-ons.
vs. hosted routers (e.g. OpenRouter) — Busbar runs in your infrastructure with your vendor keys. Nothing about your traffic or prompts leaves your network, and you pay your providers directly.
Thesis: protocols, not providers. Implement a handful of wire protocols losslessly, and every vendor that speaks one is just a catalog entry — a name, a base_url , and the env var holding its key. Six protocols are implemented; 42 vetted providers ship as catalog entries, and any OpenAI-compatible endpoint (including your own) is three lines of YAML.
Status: 1.0.0-rc.2 — feature-complete and API-stable. Release-candidate validation underway ahead of 1.0.0. AGPL-3.0.
Grab a release for your platform (Linux, macOS, Windows — Intel and ARM) from the releases page , or build from source:
cargo build --release # → target/release/busbar
2. Configure
Busbar reads two YAML files. providers.yaml is the shipped catalog (protocol, base_url , error map per provider — you rarely touch it). config.yaml is your deployment. Keys are never written into config — only the names of the env vars that hold them; ${VAR} is expanded at load time, and an unset referenced variable is a loud startup failure.
listen : " 0.0.0.0:8080 "
auth :
mode : token
client_tokens : ["${BUSBAR_CLIENT_TOKEN}"]
providers :
anthropic : { api_key_env: ANTHROPIC_KEY }
openai : { api_key_env: OPENAI_KEY }
models :
claude-sonnet : { provider: anthropic, max_concurrent: 20 }
gpt-4o-mini : { provider: openai, max_concurrent: 50 }
pools :
fast :
members :
- { target: claude-sonnet, weight: 8 }
- { target: gpt-4o-mini, weight: 2 }
3. Run
export BUSBAR_CLIENT_TOKEN=changeme ANTHROPIC_KEY=sk-ant-... OPENAI_KEY=sk-...
BUSBAR_PROVIDERS=./providers.yaml BUSBAR_CONFIG=./config.yaml ./target/release/busbar
4. Call it
Clients append the protocol path themselves, exactly as their SDK would. Anthropic-format client hitting one model:
curl -s http://localhost:8080/claude-sonnet/v1/messages \
-H " Authorization: Bearer $BUSBAR_CLIENT_TOKEN " -H " content-type: application/json " \
-d ' {"model":"ignored","max_tokens":256,"messages":[{"role":"user","content":"Hello!"}]} '
OpenAI-format client whose model selects the cross-protocol fast pool:
curl -s http://localhost:8080/v1/chat/completions \
-H " Authorization: Bearer $BUSBAR_CLIENT_TOKEN " -H " content-type: application/json " \
-d ' {"model":"fast","messages":[{"role":"user","content":"Hello!"}]} '
Busbar rewrites the request's model to the selected member and injects the provider credential; the caller's own model and key fields are ignored (except in passthrough auth mode, where the caller's key is forwarded upstream).
Busbar's scope is the protocol count (6) , not the provider count. Each protocol is a first-class ingress and egress — the format a client speaks, the format a backend speaks, or both.
Streaming is first-class for all six: Gemini via :streamGenerateContent?alt=sse , Bedrock by decoding the binary application/vnd.amazon.eventstream frames and re-framing them as the caller's protocol, the rest via SSE.
Route
Targets
POST /<name>/v1/messages
Anthropic ingress; <name> is a model ( /claude-sonnet ) or a pool ( /fast )
POST /<provider>/<model>/v1/messages
ad-hoc direct route to one provider+model, no pool needed
POST /v1/chat/completions
OpenAI ingress; the body's model selects the model or pool
GET /stats · GET /healthz · GET /metrics
per-lane health (JSON) · liveness · Prometheus
/admin/keys (POST/GET/DELETE)
virtual-key management (governance only)
Cross-protocol translation applies to all three targeting modes: the ingress protocol is fixed by the route, and if the chosen lane speaks something else, Busbar translates through the IR.
Feature
What it does
Pools & weighting
Smooth weighted round-robin across lanes; concurrency caps stack into one aggregate
Failover
Per-pool deadline + hop cap + member exclusions, applied across direct, ad-hoc, and pooled routes
Exhaustion policy
reject / status_503 / least_bad / fallback_pool:<name>
Circuit breaker
Two-stage classify → disposition; error_rate or consecutive trips; exponential cooldown; single-flight half-open recovery; honors Retry-After
Session affinity
Sticky-by-header routing while a member stays healthy
Context-length failover
Oversized request fails over to a larger-context member without penalizing the smaller lane
Health probing
none / dead / active background probes per provider
Governance
Virtual keys, allowed-pools ACLs, token-accurate budgets, RPM/TPM limits — durable in embedded SQLite, off by default
Observability
Prometheus /metrics , optional OTLP traces, optional request-log webhook
Full field-by-field config reference, with defaults and worked examples, lives in docs/configuration.md ; the architecture and the protocols-not-providers thesis are in docs/architecture.md and docs/roadmap.md .
Busbar sits in your request path, so it's built to be unremarkable there: no caller-controlled upstreams (destinations come from your vetted providers.yaml , never from request data, so it's SSRF-safe), constant-time comparison of client and admin tokens, virtual keys stored only as SHA-256 hashes, request bodies bounded at 32 MiB (even in open-relay mode), fully parameterized governance SQL, and provider keys / tokens / bodies kept out of the logs. These are exercised by the test suite (267 tests) and were the focus of a dedicated hardening pass. To report a vulnerability, see SECURITY.md .
Single Rust binary, stable toolchain (edition 2021). CI builds and tests on Linux, macOS, and Windows; releases ship x86_64 / aarch64 Linux, Intel/Apple-Silicon macOS, and x86_64 Windows.
cargo build --release
cargo test
cargo clippy --all-targets -- -D warnings
Contributing & license
Contributions welcome — see CONTRIBUTING.md . Licensed AGPL-3.0-or-later ( LICENSE ). Because Busbar typically runs as a network service, the AGPL's §13 network-use clause applies: run a modified Busbar and let others reach it ove

[truncated]
