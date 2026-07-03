---
source: "https://github.com/crp4222/PrivAiTe"
hn_url: "https://news.ycombinator.com/item?id=48776021"
title: "PrivAiTe: Self-hosted proxy that redacts PII from LLM calls, incl. tool-calls"
article_title: "GitHub - crp4222/PrivAiTe: Drop-in self-hosted LLM proxy that reversibly redacts PII before OpenAI, ChatGPT and Ollama calls, including inside tool-call arguments and multimodal content. OpenAI-compatible, zero telemetry. · GitHub"
author: "crp4222"
captured_at: "2026-07-03T15:50:10Z"
capture_tool: "hn-digest"
hn_id: 48776021
score: 2
comments: 0
posted_at: "2026-07-03T15:12:39Z"
tags:
  - hacker-news
  - translated
---

# PrivAiTe: Self-hosted proxy that redacts PII from LLM calls, incl. tool-calls

- HN: [48776021](https://news.ycombinator.com/item?id=48776021)
- Source: [github.com](https://github.com/crp4222/PrivAiTe)
- Score: 2
- Comments: 0
- Posted: 2026-07-03T15:12:39Z

## Translation

タイトル: PrivAiTe: LLM 呼び出しから PII を編集する自己ホスト型プロキシ。ツール呼び出し
記事のタイトル: GitHub - crp4222/PrivAiTe: OpenAI、ChatGPT、Ollama 呼び出しの前に PII を可逆的に編集するドロップイン セルフホスト LLM プロキシ (ツール呼び出し内部の引数やマルチモーダル コンテンツを含む)。 OpenAI 互換、テレメトリなし。 · GitHub
説明: OpenAI、ChatGPT、および Ollama 呼び出しの前に、内部ツール呼び出し引数やマルチモーダル コンテンツを含む PII を可逆的に編集する、ドロップインの自己ホスト型 LLM プロキシ。 OpenAI 互換、テレメトリなし。 - crp4222/プライベート

記事本文:
GitHub - crp4222/PrivAiTe: OpenAI、ChatGPT、Ollama 呼び出しの前に PII を可逆的に編集するドロップイン セルフホスト LLM プロキシ (内部ツール呼び出し引数やマルチモーダル コンテンツを含む)。 OpenAI 互換、テレメトリなし。 · GitHub
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
crp4222
/

プライベート
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
82 コミット 82 コミット .github .github config config docs docs 例 例 統合 統合 プライベート プライベート テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md docker-compose.yml docker-compose.yml llms.txt llms.txt pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
PII がプロバイダーに到達する前に、内部ツール呼び出し引数やマルチモーダル コンテンツを含むテレメトリなしで PII を可逆的に編集する、ドロップインの自己ホスト型 LLM プロキシ。
LLM 通話に個人データが含まれないようにします。 PrivAiTe は、アプリとモデル プロバイダーの間に位置するローカル プロキシです。名前、電子メール、電話番号、カード、IBAN、秘密などを検索し、マシンから何かが離れる前にそれらを代用と交換し、実際の値を応答に戻します。これは、メッセージ テキスト、ツール呼び出し引数、およびマルチモーダル コンテンツにわたって行われ、ほとんどのツールはここで検索を停止します。マシン上で検出が実行されますが、何も応答されません。デフォルトでは、完全な ONNX スイートが実行されるため、簡単な正規表現エンティティだけでなく、シークレットとパスワードも捕捉します。 OpenAI 互換クライアントをそれに向けます。
「Je m'appelle Marie Dupont、メール marie@acme.com」と入力します。
LLM は次のように認識します:「Je m'appelle <PERSON_1>、電子メール <EMAIL_ADDRESS_1>」
LLM は次のように言います: 「<PERSON_1> さん、電子メール <EMAIL_ADDRESS_1> に投票してください。」
「Bonjour Marie Dupont、電子メール marie@acme.com est noté にご投票ください。」
これはローカルの仮名化であり、匿名化ではありません。

検出は保証ではなくベストエフォート型です。あなたは引き続きデータ管理者です。脅威モデルは、何を防御し、何を防御しないかを正確に説明します。
PrivAiTe は、一緒にまたは個別に実行できる 2 つの検出エンジンを使用します。
Presidio (Microsoft): 正規表現 + spaCy NER
デフォルトのエンジン。パターン マッチングと基本的な NER を通じて構造化 PII を処理します。
Presidio は高速 (リクエストあたり数十ミリ秒) で、コード、ニュース記事、技術文書での誤検知はほとんど発生しません。その代償として、spaCy が認識しない名前 (珍しい名前、文脈のない単一の単語の名前) が見逃され、シークレット/パスワードは検出されません。
OpenAI プライバシー フィルター: コンテキスト ML モデル
OpenAI のオープンソース PII モデル (1.5B パラメーター、50M アクティブ、Apache 2.0)。 ONNX ランタイム経由でローカルで実行します (~800MB、PyTorch は必要ありません)。
プライバシー フィルターは低速 (リクエストあたり約 400 ミリ秒) で、技術的な識別子にアカウント番号 (例: "CMD-2024-98765") としてフラグを付けることがあります。これは、Presidio と並行して 2 番目のパスとして実行されます。Presidio は正規表現ベースのエンティティを処理し、プライバシー フィルターはコンテキスト NER を処理します。
Presidio だけでは、spaCy が認識できない名前を見逃すため、秘密を検出できません。しかし、誤検知はほとんどありません。
プライバシー フィルターだけでは、クレジット/リスト形式の一部の名前が欠落し、IBAN/クレジット カード チェックサムの正規表現検証機能がありません。
両方が一緒になってお互いの盲点をカバーします。 Presidio は検証付きの構造化フォーマットを処理し、プライバシー フィルターはコンテキスト依存の PII を処理します。
onnx がデフォルトです。完全なスイートを実行し、シークレットやパスワードを含むすべてを検出します。 light は、従来の PII のみを気にする場合に適した、より高速な Presidio 専用のオプションで、誤検知がほとんどありません。
*独立した 120 ドキュメントの AI4Privacy ベンチマーク (スパン レベル、厳密なトークン レベル ~80% / ~58%) を思い出してください。レイテンシーは

ハードウェアに依存します。 「ベンチマーク」を参照してください。
** max は、AI4Privacy から独立したデータでトレーニングされた PII モデルである GLiNER を追加します。分布外のコーパスでは、より多くの誤検知とトーチ依存性を犠牲にして、デフォルトがすでによく一般化している場合、再現率が ~84% から ~89% に上昇します。それはオプトインです。 onnx はデフォルトのままです。数値: OOD クロスチェック ( OOD_COMPARISON.md )。
ピイ：
プリセット: " onnx " # デフォルト。秘密を含むあらゆるものを検出します。初回実行時にモデルをダウンロードします。
# プリセット: "light" # より高速、Presidio のみ、クラシック PII のみ。
# プリセット: "max" # onnx + GLiNER: 配布外リコールが高くなります。必要なもの: pip install 'privaite[gliner]'
フットガン: detecters.presidio.entities をライト パス上の短い許可リストに固定しないでください。検出をそれらのタイプのみに制限し、再現率をおよそ半分にします (最大 35%)。エンティティを未設定のままにしておきます。プロキシは、低リコール構成を検出した場合、起動時に警告をログに記録します。
デフォルトのインストールにはすでに onnxruntime とトークナイザーが含まれているため、onnx プリセットはそのまま使用できます。モデルは、プロキシの初回起動時にダウンロードされます。 2 つのオプションの追加機能として、トーチ: ml (標準/フル BERT プリセット) と gliner (最大プリセット) が追加されます。 max が選択されているが、privaite[gliner] がインストールされていない場合、プロキシは検出をサイレントに削減して実行するのではなく、インストール ヒントで起動に失敗します。
onnx を使用する場合 (デフォルト): 最大限のカバレッジが必要です。シークレット、パスワード、API キー、アカウント番号、珍しい名前。技術識別子に関する時折の誤検知を受け入れます。
ライトを使用する場合: 中断をゼロにして最速のパスを実現したい場合。コード、ニュース、ビジネス テキストはすべてそのまま通過します。明確に識別可能な PII (名前、電子メール、電話、カード、IBAN) のみが匿名化されます。
他に 2 つのプリセット ( standard 、 full ) が存在しますが、実際にはあまり役に立ちません。B が追加されます。

ERT NER。spaCy よりもあまり改善されず、PyTorch を引き込みます。
DE、EN、FR、IT にわたる、Hugging Face 上のオープン AI4Privacy pii-masking-200k データセットからの 120 件の実際のドキュメント (458 個の PII 項目、10 人の独立した監査エージェントによってラベル付けされ、データセット独自の機密マスクとクロスチェックされた) と、誤検知の 14 件のクリーンなドキュメントで測定されました。データセットは明示的なライセンスを宣言していないため、ベンチマーク リポジトリは派生ラベルのみをコミットし、オンデマンドでソース テキストを取得します。
2 つのリコール列: スパンは、正確な完全な文字列が消失したときに捕捉されたマルチトークン PII スパン (上限) をクレジットします。 strict では、スパンのすべてのトークンを削除する必要があります。 Presidio ベースラインは、一般的なフラットテキスト アプローチ (ほとんどのドロップイン PII プロキシの背後にあるエンジン) です。設計上、ツール呼び出しの引数やマルチモーダル コンテンツには触れません。これは PrivAiTe が埋めるギャップです。したがって、ツール呼び出しの保護は 0.6% に対して 100% となります。構造化されたコラムは、「すべての競合他社との比較」ではなく、「構造化を意識したアプローチとフラットテキストのアプローチ」として読んでください。
言語ごとおよびエンティティごとのテーブル、方法論、および再現方法: private-bench 。 Presidio、LLM Guard、LiteLLM PII マスキングとの機能の直接比較: docs/comparison.md 。
(初期の、部分的に合成された 61 個の文書からなる小規模に厳選されたセットのスコアは高く、ライトは約 97%、オンエックスは約 100% でした。これは、そのデータが簡単であるためです。上記の AI4Privacy の数値は、現実的で独立した数値です。)
デフォルトで検出されないもの
デフォルトの onnx プリセットは、プライバシー フィルター モデルを通じて個人のアドレス ( LOCATION として) と個人の URL ( URL として) を検出し、それらを置き換えます。デフォルトでオフのままなのは、これらのタイプに対する Presidio の広範な認識機能です。これらは重大な誤検知を引き起こすためです。
一般的な地名 (Presidio LOCATION 認識エンジン): 「Paris」または「London」

相続人自身のものは PII ではなく、spaCy は通常の単語 (「Kubernetes」、「Saturday」) を場所としてフラグを立てます。 onnx プリセットは、この認識機能をオフのままにし、代わりにモデルのコンテキスト認識アドレス検出に依存します。
Presidio URL 正規表現: .ge は有効な TLD であるため、logging.getLogger のようなコードと一致します。 onnx プリセットはそれをオフに保ち、モデルは依然として本物の個人 URL を捕捉します。
ライト プリセット (Presidio のみ) では、アドレスと URL は検出されません。シークレットとパスワードは、onnx プリセットによってのみ検出されます。 YAML 構成で任意の認識機能を有効にすることができます。
PrivAiTe はローカルの仮名化を実行しますが、匿名化は保証されません。検出はマシン上で実行されます。実際の ↔ プレースホルダーのマッピングは、リクエストの間のみメモリ内に存在し、その後削除されます。
保護対象: 生の PII を保存、トレーニング、またはログに記録する LLM プロバイダー。プロバイダーは、メッセージ コンテンツ、ツール呼び出し引数、およびマルチモーダル テキストにわたって、検出器がキャッチしたすべてのプレースホルダー ( <PERSON_1> , …) を受け取ります。
保護できないもの:
PII 検出器がミスしました。検出は統計的であり、100% であることはありません (ベンチマークを参照)。認識できない名前がプロバイダーに届きます。 onnx プリセットのリコールが最も優れています。出力は保証ではなくベストエフォートとして扱われます。
文脈からの再識別。名前を置き換えたとしても、周囲のテキストは識別可能なままです (「3 月に辞任した <ORG_1> の CEO」)。
侵害されたローカル マシン。マッピングと生のテキストはローカル メモリ内に存在します。これはローカルの攻撃者に対する防御ではありません。
セッション内のリクエストを相関付けるプロバイダー。
GDPR/HIPAA の場合: これを匿名化ではなく、仮名化 + 転送の最小化として扱います。不可逆的な削除が必要な場合は、メソッド: "placeholder" の代わりにメソッド: "redact" を使用してください。
PIIの侵入を防ぐ

LLM 通話の多くは混雑した空間であり、PrivAiTe が常に正しい選択であるとは限りません。 2026 年 6 月時点の各プロジェクトの公開ドキュメントに基づいています。
LiteLLM には Presidio ガードレールが組み込まれており、LiteLLM プロキシをすでに実行していて PII をインラインで処理したい場合は、自然な選択です (リクエストと応答のスクラビングに関して未解決のバグがいくつかあります)。
Microsoft PII Shield や LangChain のゲートウェイ編集など、マネージド/クラウド オプションも存在します。
PrivAiTe が異なる点は、メッセージ テキストだけでなく、ツール呼び出し引数とマルチモーダル コンテンツ内の PII を匿名化し (たとえば、LangChain のゲートウェイ ドキュメントでは、ツール呼び出し引数はスキャンされないことに注意してください)、応答内の元の値を復元し、再現可能なベンチマークを提供します。トラフィックがエージェント的またはマルチモーダルである場合、そのギャップがこのギャップが存在する理由です。
pip install -e 。
python -m spacy ダウンロード en_core_web_lg
python -m spacy ダウンロード fr_core_news_md
デフォルトの onnx プリセットは、プロキシの初回起動時にモデルをダウンロードします。モデルをダウンロードせずに、より軽量で高速なパスをご希望ですか?設定でプリセット「light」を設定します。
cp .env.example .env
cp config/privaite.example.yaml config/privaite.yaml
.env を API キーで編集し、config/privaite.yaml を LLM プロバイダーで編集します。
Python -m pr

[切り捨てられた]

## Original Extract

Drop-in self-hosted LLM proxy that reversibly redacts PII before OpenAI, ChatGPT and Ollama calls, including inside tool-call arguments and multimodal content. OpenAI-compatible, zero telemetry. - crp4222/PrivAiTe

GitHub - crp4222/PrivAiTe: Drop-in self-hosted LLM proxy that reversibly redacts PII before OpenAI, ChatGPT and Ollama calls, including inside tool-call arguments and multimodal content. OpenAI-compatible, zero telemetry. · GitHub
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
crp4222
/
PrivAiTe
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
82 Commits 82 Commits .github .github config config docs docs examples examples integrations integrations privaite privaite tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml llms.txt llms.txt pyproject.toml pyproject.toml View all files Repository files navigation
A drop-in, self-hosted LLM proxy that reversibly redacts PII before it reaches the provider, including inside tool-call arguments and multimodal content, with zero telemetry.
Keep personal data out of your LLM calls. PrivAiTe is a local proxy that sits between your app and the model provider. It finds names, emails, phone numbers, cards, IBANs, secrets and more, swaps them for stand-ins before anything leaves your machine, and puts the real values back in the reply. It does this across message text, tool-call arguments, and multimodal content , which is where most tools stop looking. Detection runs on your machine and nothing phones home. By default it runs the full ONNX suite, so it also catches secrets and passwords , not just the easy regex entities. Point any OpenAI-compatible client at it.
You type: "Je m'appelle Marie Dupont, email marie@acme.com"
LLM sees: "Je m'appelle <PERSON_1>, email <EMAIL_ADDRESS_1>"
LLM says: "Bonjour <PERSON_1>, votre email <EMAIL_ADDRESS_1> est noté."
You see: "Bonjour Marie Dupont, votre email marie@acme.com est noté."
This is local pseudonymization, not anonymization, and detection is best-effort rather than a guarantee. You remain the data controller. The Threat model spells out exactly what it protects against and what it does not.
PrivAiTe uses two detection engines that can run together or separately:
Presidio (Microsoft): regex + spaCy NER
The default engine. Handles structured PII through pattern matching and basic NER.
Presidio is fast (tens of ms/request) and produces very few false positives on code, news articles, and technical text. The tradeoff: it misses names that spaCy doesn't recognize (unusual names, single-word names without context) and doesn't detect secrets/passwords.
OpenAI Privacy Filter: contextual ML model
OpenAI's open-source PII model (1.5B params, 50M active, Apache 2.0). Runs locally via ONNX Runtime (~800MB, no PyTorch needed).
The Privacy Filter is slower (~400ms/request) and occasionally flags technical identifiers as account numbers (e.g., "CMD-2024-98765"). It runs as a second pass alongside Presidio, which handles the regex-based entities while the Privacy Filter handles contextual NER.
Presidio alone misses names that spaCy doesn't recognize, and can't detect secrets. But it has very few false positives.
Privacy Filter alone misses some names in credit/list formats, and doesn't have regex validators for IBAN/credit card checksums.
Both together cover each other's blind spots. Presidio handles structured formats with validation, the Privacy Filter handles context-dependent PII.
onnx is the default. It runs the full suite and detects everything, including secrets and passwords. light is a faster, Presidio-only option with very few false positives, for when you only care about classic PII.
*Recall on the independent 120-document AI4Privacy benchmark (span-level; strict token-level ~80% / ~58%). Latency is hardware-dependent. See Benchmark .
** max adds GLiNER, a PII model trained on data independent of AI4Privacy. On an out-of-distribution corpus it raises recall from ~84% to ~89% where the default already generalizes well, at the cost of more false positives and a torch dependency. It is opt-in; onnx stays the default. Numbers: the OOD cross-check ( OOD_COMPARISON.md ).
pii :
preset : " onnx " # Default. Detects everything including secrets. Downloads the model on first run.
# preset: "light" # Faster, Presidio-only, classic PII only.
# preset: "max" # onnx + GLiNER: higher out-of-distribution recall. Needs: pip install 'privaite[gliner]'
Footgun: do not pin detectors.presidio.entities to a short allowlist on the light path. It restricts detection to only those types and roughly halves recall (to ~35%). Leave entities unset; the proxy logs a warning at startup if it detects a low-recall configuration.
The default install already includes onnxruntime and the tokenizer, so the onnx preset works out of the box. The model is downloaded the first time the proxy starts. Two optional extras add torch: ml (the standard / full BERT presets) and gliner (the max preset). With max selected but privaite[gliner] not installed, the proxy fails to start with an install hint rather than running with detection silently reduced.
When to use onnx (default): You want maximum coverage. Secrets, passwords, API keys, account numbers, unusual names. Accept occasional false positives on technical identifiers.
When to use light : You want zero disruption and the fastest path. Code, news, business text all pass through untouched. Only clearly identifiable PII (names, emails, phones, cards, IBANs) is anonymized.
Two other presets exist ( standard , full ) but are less useful in practice: they add BERT NER, which does not improve much over spaCy and pulls in PyTorch.
Measured on 120 real documents from the open AI4Privacy pii-masking-200k dataset on Hugging Face (458 PII items, labeled by 10 independent auditor agents and cross-checked against the dataset's own sensitive mask) across DE, EN, FR, IT, plus 14 clean documents for false positives. The dataset declares no explicit license, so the benchmark repo commits only derived labels and fetches the source text on demand.
Two recall columns: span credits a multi-token PII span as caught when its exact full string disappears (an upper bound); strict requires every token of the span to be removed. The Presidio baseline is the common flat-text approach (the engine behind most drop-in PII proxies); by design it does not touch tool-call arguments or multimodal content, which is the gap PrivAiTe closes — hence 100% tool-call protection vs 0.6%. Read the structured columns as "structured-aware vs the flat-text approach", not "vs every competitor".
Per-language and per-entity tables, the methodology, and how to reproduce: privaite-bench . Head-to-head feature comparison with Presidio, LLM Guard, and LiteLLM PII masking: docs/comparison.md .
(An earlier, smaller curated set of 61 partly-synthetic documents scores higher — light ~97%, onnx ~100% — because that data is easier; the AI4Privacy figures above are the realistic, independent numbers.)
What's NOT detected by default
The default onnx preset does detect personal addresses (as LOCATION ) and personal URLs (as URL ) through the Privacy Filter model, and replaces them. What stays off by default are Presidio's broad recognizers for those types, because they cause heavy false positives:
Generic place names (the Presidio LOCATION recognizer): "Paris" or "London" on their own aren't PII, and spaCy flags ordinary words ("Kubernetes", "Saturday") as locations. The onnx preset keeps this recognizer off and relies on the model's context-aware address detection instead.
The Presidio URL regex: it matches code like logging.getLogger because .ge is a valid TLD. The onnx preset keeps it off, and the model still catches genuine personal URLs.
On the light preset (Presidio only), addresses and URLs are not detected. Secrets and passwords are detected only by the onnx preset. Any recognizer can be turned on in the YAML config.
PrivAiTe performs local pseudonymization , not guaranteed anonymization. Detection runs on your machine; the real ↔ placeholder mapping lives in memory only for the duration of a request and is dropped afterwards.
What it protects against: the LLM provider storing, training on, or logging your raw PII. The provider receives placeholders ( <PERSON_1> , …) for everything the detector catches, across message content, tool-call arguments, and multimodal text.
What it does NOT protect against:
PII the detector misses. Detection is statistical and never 100% (see the benchmark ). A name it doesn't recognize reaches the provider. The onnx preset has the best recall; treat the output as best-effort, not a guarantee.
Re-identification from context. Even with names replaced, the surrounding text can stay identifying ("the CEO of <ORG_1> who resigned in March").
A compromised local machine. The mapping and raw text live in local memory; this is not a defense against a local attacker.
The provider correlating requests within a session.
For GDPR/HIPAA: treat this as pseudonymization + transfer minimization, not anonymization. If you need irreversible removal, use method: "redact" instead of method: "placeholder" .
Keeping PII out of LLM calls is a crowded space, and PrivAiTe is not always the right pick. Based on each project's public docs as of June 2026:
LiteLLM has a built-in Presidio guardrail, the natural choice if you already run the LiteLLM proxy and want PII handling inline (there are a few open bugs around scrubbing requests and responses).
Managed/cloud options exist too, such as Microsoft PII Shield and LangChain's gateway redaction .
Where PrivAiTe differs: it anonymizes PII inside tool-call arguments and multimodal content , not just message text (LangChain's gateway docs, for instance, note that tool-call arguments are not scanned), it restores the original values in the response, and it ships a reproducible benchmark . If your traffic is agentic or multimodal, that gap is the reason this exists.
pip install -e .
python -m spacy download en_core_web_lg
python -m spacy download fr_core_news_md
The default onnx preset downloads its model the first time the proxy starts. Want the lighter, faster path with no model download? Set preset: "light" in your config.
cp .env.example .env
cp config/privaite.example.yaml config/privaite.yaml
Edit .env with your API keys and config/privaite.yaml with your LLM providers.
python -m pr

[truncated]
