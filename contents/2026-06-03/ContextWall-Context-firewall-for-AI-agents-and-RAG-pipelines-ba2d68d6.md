---
source: "https://contextwall.io/"
hn_url: "https://news.ycombinator.com/item?id=48377225"
title: "ContextWall – Context firewall for AI agents and RAG pipelines"
article_title: "ContextWall: Context Firewall for AI Agents"
author: "sumeshpk"
captured_at: "2026-06-03T00:35:01Z"
capture_tool: "hn-digest"
hn_id: 48377225
score: 2
comments: 0
posted_at: "2026-06-02T22:31:35Z"
tags:
  - hacker-news
  - translated
---

# ContextWall – Context firewall for AI agents and RAG pipelines

- HN: [48377225](https://news.ycombinator.com/item?id=48377225)
- Source: [contextwall.io](https://contextwall.io/)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T22:31:35Z

## Translation

タイトル: ContextWall – AI エージェントと RAG パイプライン用のコンテキスト ファイアウォール
記事のタイトル: ContextWall: AI エージェント用のコンテキスト ファイアウォール
説明: プロンプト インジェクション、ポイズニングされた RAG、資格情報の漏洩がモデルに到達する前に阻止します。 YAML でセキュリティ ポリシーを定義します。エージェントに対するコードの変更はありません。

記事本文:
ContextWall: AI エージェント用の Context ファイアウォール ContextWall の仕組み クイックスタート プライバシー コンプライアンス 価格 GitHub 早期アクセスの取得 無料の早期アクセス · Apache 2.0 オープン ソース AI エージェントは信頼できないコンテンツを読み取ります。
エージェントが取得したすべての Web 結果、ドキュメント、および API 応答は、スクリーンなしでモデルのコンテキスト ウィンドウに直接入力されます。 ContextWall は、最初にそれを傍受し、プロンプト インジェクションと資格情報の漏洩をブロックし、LLM がそれを認識する前にセキュリティ ポリシーを適用します。
セルフホストの無料コンテキストウォール: ライブ施行フィード -- 待機中 _ エージェントへのコード変更なし · インフラストラクチャ内で実行 · スクリーニング パスに LLM なし
理論上の脅威ではなく、実際の本番環境のインシデント エージェントは、読み取ったものすべてを信頼します
LLM には、ソースの信頼という概念が組み込まれていません。 Web 検索から取得したコンテンツとシステム プロンプトからのコンテンツは、両方ともコンテキスト ウィンドウ内に表示されると同一に見えます。攻撃者はこれを直接悪用します。
攻撃者が細工したメールを送信します。 Copilot はそれを読み取り、埋め込まれた命令をコマンドとして解釈し、サイレントに内部 SharePoint ファイルにアクセスして、攻撃者に送信します。ユーザーは何もクリックしません。
Copilot には、信頼できるシステム命令と信頼できない電子メールの内容を区別する方法がありませんでした。コンテキスト ウィンドウ内ではどちらも同じように見えました。
研究者たちは、5 つの敵対的な文書を数百万の知識ベースに埋め込みました。ユーザーが質問すると、モデルは偽のコンテンツを取得して確信のある事実として繰り返しました。脱獄もシステムプロンプトの変更も、モデルへのアクセスも必要ありませんでした。
RAG パイプラインは、関連性スコアによってドキュメントを取得し、それらをモデルに直接渡します。文書の出所や信頼できるかどうかについてのチェックはありませんでした。
どちらの攻撃も同じギャップを悪用しました。つまり、攻撃には信頼境界がありません。

テキストレイヤー。 ContextWall は、すべてのコンテキスト ソースに信頼層をタグ付けし、コンテンツがモデルに到達する前にポリシー ルールを適用することで、この問題を修正します。
ContextWall は、ガイドラインだけでなくセキュリティの保証を必要とする AI を本番環境に導入するチームのために構築されています。
Web、内部ドキュメント、サードパーティ API から取得する RAG パイプラインとエージェント システムを出荷しています。取得されたすべてのドキュメントは潜在的な攻撃ベクトルであり、エージェントには正規のソースと汚染されたソースを区別する方法がありません。
1 つの pip インストールまたは Docker イメージ - エージェント コードの変更は不要
プロンプトに入る前にすべての文書を検査します
LLM がインジェクションと認証情報の漏洩を検知する前にブロックします。
AI システムは既存の境界制御をバイパスします。エージェントは、発信通話を行い、信頼できないコンテンツを取り込み、幅広い権限で動作します。これらはすべて、従来の検出スタックの外側で行われます。
ソース、チーム、リポジトリごとに適用可能なポリシー ルール
リアルタイムの施行フィードと改ざん明示的な監査ログ
導入されたすべてのエージェントにわたるフリート全体の可視性
HIPAA、SOC 2、および FedRAMP の監査人は、PHI が AI エージェントのコンテキスト ウィンドウから漏洩しないようにする方法を尋ねています。必要なのは証拠ではなく、保証です。
すべての施行決定はコンプライアンス制御 ID にマッピングされます
オンデマンドで暗号化署名された監査エクスポート
文書化されたデータ常駐: コンテキストがインフラストラクチャから離れることはありません
ContextWall が何を止めるのか、そして何を止めないのか
コンテキスト層での検出。スクリーニング パスに LLM がありません。私たちは範囲について正直に話します。
直接命令オーバーライド L1 + L2 「以前の命令をすべて無視…」
Bidi およびゼロ幅難読化 L1 RTL は、取得されたテキストに隠された文字をオーバーライドします
スペース文字インジェクション L1「無視します」
意味的パラフレーズ挿入 L3 「あなたの割り当ては置き換えられました…」
認証情報の漏洩 L2 AWS キー、GitH

ub PAT、無記名トークン
コンテキスト L2 電子メール、信頼できない層のドキュメント内の SSN を介した PII の漏洩
ContextWall はコンテキスト ウィンドウに入る内容をフィルターします。モデルがクリーンな入力から何を生成するかを制御することはできません。
システム プロンプトが過剰な権限を付与する場合、ContextWall はその設計上の決定をオーバーライドできません。
モデルの重みやデータの微調整に対する攻撃は、推論の前に発生します。 ContextWall は推論時のみに動作します。
L3 ヒューリスティックは、既知の意味上の言い換えをキャッチします。十分に斬新な攻撃の場合、スコアがブロックしきい値を下回る可能性があります。そのしきい値はユーザーが設定します。
あなたが許可した許可されたアクセス
ポリシーがソースを許可し、モデルがそのデータを使用する場合、ContextWall は、より厳格なポリシーではなく、ポリシーを強制します。
誠実な範囲は誤った保証に勝ります。多層防御とは、ContextWall がモデル プロバイダーの安全フィルターの代わりではなく、並行して機能することを意味します。
ContextWall は、すべてのドキュメントがコンテキスト ウィンドウに入る前にインターセプトします。ここでまさに何が起こります。
エージェントが文書を要求します
Web 検索結果、内部ドキュメント、API 応答、またはユーザーのアップロードなど、あらゆる外部コンテンツ。
デーモンは、コンテキスト ウィンドウに入る前にドキュメントを受け取ります。 LLM はまだ何も見ていません。
3 3 つの検出レイヤーが順番に実行されます
400 がエージェントに返されました。文書が LLM に到達することはありません。改ざん防止監査ログに書き込まれるイベント。
ドキュメントが LLM API に転送されました。クリーン コンテキストが通常どおりプロンプトに入ります。
スクリーニング パスに LLM がありません。外部通話はありません。データはホスト上に残ります。
各コンテキスト ソースが何であるかを宣言します。 ContextWall は、その層に基づいて適切なレベルの精査を自動的に適用します。
コード リポジトリ、内部 Wiki
パブリック Web、ユーザーが送信した入力
最も安いものから最も徹底したものへの順序で適用されます。外部呼び出しや LLM 推論はありません。
未加工のバイトをスキャンして既知の obfus を探します

ヒント: 双方向制御文字、ゼロ幅文字、スペース文字キーワード (「i g n or e a l l」)。これらは人間の目には見えませんが、モデルによって読み取ることができます。
正規化されたテキストに対して正規表現パターンを実行します。インジェクション構文、公開された API キー (AWS、GitHub、Anthropic)、ベアラー トークン、電子メール、電話番号、SSN などの PII をキャッチします。
文言が明白なキーワードを避けている場合でも、指示のような意図があるかどうかを各メッセージにスコア付けします。正規表現を完全にバイパスする「以前の割り当ては置き換えられました」のような言い換えを捕捉します。
データがインフラストラクチャから離れることはありません コンテキストはそのまま残ります
ContextWall は、独自のインフラストラクチャ内でデーモンとして実行されます。プロンプト、ドキュメント、ファイルの内容はローカルで検査され、どこにも送信されることはありません。クラウド コントロール プレーンはカウントとスコアのみを受信し、コンテンツは受信しません。
インフラストラクチャ AI エージェント ContextWall デーモン LLM API 呼び出し すべてのスクリーニングがここで行われます。何も出てこない。
リクエスト数 違反タイプ 遅延 (ミリ秒) セッション数 送信されない
プロンプト コンテンツ ファイル コンテンツ ユーザー データ PII / PHI ContextWall クラウド フリート ダッシュボード ポリシー オーサリング コンプライアンス レポート カウントとスコアのみを表示します。決して満足しないでください。
インフラストラクチャに留まり、常にプロンプトコンテンツとユーザーメッセージ
取得したドキュメントとファイルの内容
モデルの応答と完了
個人を特定できる情報
保護された健康情報 (PHI)
コントロール プレーンに送信、メタデータのみ リクエスト数 (ブロック/許可)
検出された違反タイプ (例: "pii")
平均遅延 (ミリ秒)
ポリシーバージョンの確認
完全にオフラインの方がいいですか? control_plane.url を空のままにすると、ContextWall はクラウドに依存せずに完全にローカルで実行されます。
デーモンは pip を使用してインストールされ、AI SDK 呼び出しをローカルにプロキシします。クラウド ダッシュボードはオプションであり、

集約されたメタデータのみが表示され、コンテンツは表示されません。
プロンプトの前にドキュメントをスクリーニングする
# 1. デーモンをインストールして起動します (インフラストラクチャ内で実行します)。
pip インストールコンテキストウォール
ctxfw 開始 --config ctxfw.yaml
# 2. ctxfw.yaml でコンテキスト ソースを宣言する
情報源:
- ID: ウェブ検索
タイプ: ウェブ
trust_tier: 信頼されていない
- ID: 内部ドキュメント
タイプ: 合流
trust_tier: 内部
# 3. AI SDK をローカル デーモンに向ける
エクスポート ANTHROPIC_BASE_URL=http://localhost:8080/proxy/anthropic
import ANTHROPIC_API_KEY=sk-ant-your-key # 変更なし
# エージェントが行うすべての通話はローカルでスクリーニングされるようになりました。
# プロンプトと応答がマシンから離れることはありません。どの言語でも Anthropic および OpenAI SDK で動作します。デーモンはリクエストをプロキシし、コンテキストをスクリーニングし、プロンプトを保存したり送信したりすることなく、クリーンなコンテンツを実際の API に転送します。
すべては YAML で宣言されます。ソース、ルール、およびしきい値はすべてファイル内に存在し、リポジトリにコミットし、プル リクエストでレビューし、他のインフラストラクチャ構成と一緒にデプロイします。
ソースは起動時に設定で宣言され、API 呼び出しやセットアップ スクリプトは不要
4 層ポリシー: 個々のリポジトリのオーバーライドに至るまでのフリート全体のルール
HIPAA、SOC2、および FedRAMP のスターター ポリシー テンプレートが含まれています
ルールはファイル変更後 5 秒以内にリロードされ、再起動は必要ありません
すべてのルールを監査証拠のコンプライアンス コントロール ID にマッピングできます。
#ctxfw.yaml
# ここでコンテキスト ソースを宣言します。 ContextWall がそれらを登録します
# 起動時に自動的に実行されます。 API 呼び出しやスクリプトの実行は必要ありません。
情報源:
- ID: ブレイブウェブ検索
タイプ: ウェブ
trust_tier: 信頼されていない
- ID: 内部合流
タイプ: 合流
trust_tier: 内部
データ分類: 機密性の高い
- ID: fhir-api
タイプ: API
trust_tier: 規制済み
データ分類: ファイ
所有者: 臨床データチーム コンプライアンスを念頭に置いて設計 コンプライアンス対応範囲
コ

ntextWall は、コンプライアンスがアーキテクチャの特性であり、後から完成させるチェックリストではないように設計されています。 PHI、PII、および機密データは、公開される前にローカルで処理されます。
保護された健康情報はローカルで検査されます。サードパーティのサーバーを経由することはありません。
規制されたソース層により、PHI は承認された内部システム間でのみフローできるようになります
すべての施行に関する決定は、タイムスタンプ、ソース ID、および監査人のレビューのために結果とともに記録されます。
違反イベントはタイムスタンプ、ソース ID、および監査人のレビューのためにポリシー決定とともに記録されます。
レビュー担当者が検証できる監査証跡
すべてのコンテキスト スクリーニング イベントは、ソース ID、信頼層、決定、およびタイムスタンプとともに記録されます。
来歴チェーンは暗号的にリンクされています。記録は検出されなければ変更できない
フリート ダッシュボードへのロールベースのアクセス。生のコンテキストはどこにも保存されず、どこにもアクセスできません
ポリシー ルールはバージョン管理されています。変更により完全な監査証跡が残る
個人データは設計上あなたの管轄内に留まります
PII (電子メール アドレス、電話番号、名前) が検出され、モデルに到達する前に編集されます。
デーモンは、国境を越えたデータ転送を行わずに、独自のインフラストラクチャ内ですべてのデータを処理します。
コントロール プレーンは集計されたカウントのみを受信し、個人データは受信しません
設計によりデータの最小化をサポートします。モデルはタスクを完了するために必要な最小限のデータを認識します。
完全にエアギャップ。外部依存関係はありません。
デーモンは外部依存関係なしで完全にエアギャップ環境または VPC 環境内で実行されます。
完全にオフラインで展開する場合は、control_plane.url を空のままにしておきます。テレメトリはホストから送信されません。
ポリシーと構成は、ユーザーが制御するファイル内に存在します。クラウドには何も保存されません
コントロール プレーンに到達できない場合でも、ローカルで強制が継続されます。
すべての施行はインフラストラクチャ内で実行されます - オープンソース

、永久無料。クラウドは、フリートの可視性とコンプライアンス ツールを最上位に追加します。
永遠に。 GitHub からのセルフホスト。
3 層検出エンジン (構造、正規表現、ヒューリスティック)
ソースの信頼層 (内部/外部/非信頼/規制)
YAML のコードとしてのポリシー、5 秒でホットリロード
改ざんが明らかなマークル監査ログ
Python SDK - SafeAnthropic、SafeOpenAI
HIPAA、SOC2、FedRAMP 用の事前構築済みポリシー パック
Prometheus メトリクス + ライブ WebSocket 強制フィード
インフラストラクチャ内で完全に実行されます。外部通話はありません。
GitHub でのセルフホスト 早期アクセスでは無料 クラウド + セルフホスト型 無料
クレジットカードはありません。サインインして開始してください。
オープンソースのすべてに加えて:
フリート ダッシュボード - すべてのデーモン、ステータス、ブロック レート
ポリシー作成 UI - YAML 編集は不要
コンプライアンスレポートの生成 (SOC2、HIPAA、FedRAMP エクスポート)
暗号署名されたワンクリック監査証跡エクスポート
クラウドにアクセスできるかどうかに関係なく、適用はローカルに維持されます。プロンプトとドキュメントがインフラストラクチャから離れることはありません。
機能 オープンソースのクラウド強制 (ブロック、検出) 監査ログ (ローカル、改ざん明示) コードとしてのポリシー (YAML) Python SDK Prometheus メトリクス フリート ダッシュボード (マルチデーモン) ポリシー作成 UI コンプライアンス レポートのエクスポート 電子メール サポート

[切り捨てられた]

## Original Extract

Stop prompt injection, poisoned RAG, and credential leakage before they reach your model. Define security policy in YAML. No code changes to your agents.

ContextWall: Context Firewall for AI Agents ContextWall How it works Quickstart Privacy Compliance Pricing GitHub Get early access Free early access · Apache 2.0 open source Your AI agent reads untrusted content.
Every web result, document, and API response your agent retrieves goes straight into the model's context window - unscreened. ContextWall intercepts it first, blocks prompt injection and credential leaks, and enforces your security policy before the LLM ever sees it.
Self-host free contextwall: live enforcement feed -- waiting _ No code changes to your agents · Runs in your infrastructure · No LLM in the screening path
Real production incidents, not theoretical threats Your agent trusts everything it reads
LLMs have no built-in concept of source trust. Content retrieved from a web search and content from your system prompt look identical once they are both inside the context window. Attackers exploit this directly.
An attacker sends a crafted email. Copilot reads it, interprets embedded instructions as commands, silently accesses internal SharePoint files, and sends them to the attacker. The user never clicks anything.
Copilot had no way to distinguish a trusted system instruction from untrusted email content. Both looked the same inside the context window.
Researchers planted five adversarial documents into a knowledge base of millions. When users asked questions, the model retrieved and repeated the false content as confident fact, with no jailbreak, no system prompt change, and no model access needed.
The RAG pipeline retrieved documents by relevance score and passed them straight to the model. There was no check on where the document came from or whether it should be trusted.
Both attacks exploited the same gap: no trust boundary at the context layer. ContextWall fixes this by tagging every context source with a trust tier and applying your policy rules before content reaches the model.
ContextWall is built for teams shipping AI into production who need security guarantees - not just guidelines.
You're shipping RAG pipelines and agentic systems that pull from the web, internal docs, and third-party APIs. Every retrieved document is a potential attack vector - and your agent has no way to tell a legitimate source from a poisoned one.
One pip install or Docker image - no changes to your agent code
Screens every document before it enters the prompt
Blocks injections and credential leaks before the LLM sees them
AI systems bypass your existing perimeter controls. Agents make outbound calls, ingest untrusted content, and operate with broad permissions - all outside your traditional detection stack.
Enforceable policy rules per source, team, and repo
Real-time enforcement feed and tamper-evident audit log
Fleet-wide visibility across all deployed agents
HIPAA, SOC 2, and FedRAMP auditors are asking how PHI can't leak through an AI agent's context window. You need evidence - not assurances.
Every enforcement decision mapped to a compliance control ID
Cryptographically signed audit exports on demand
Documented data residency: context never leaves your infrastructure
What ContextWall stops - and what it doesn't
Detection at the context layer. No LLM in the screening path. We're honest about the scope.
Direct instruction override L1 + L2 "IGNORE ALL PREVIOUS INSTRUCTIONS…"
Bidi & zero-width obfuscation L1 RTL override chars hidden in retrieved text
Spaced-letter injection L1 "i g n o r e p r e v i o u s"
Semantic paraphrase injection L3 "Your assignment has been superseded…"
Credential leakage L2 AWS keys, GitHub PATs, bearer tokens
PII exfiltration via context L2 Emails, SSNs in untrusted-tier documents
ContextWall filters what enters the context window - it cannot control what the model generates from clean inputs.
If your system prompt grants excessive permissions, ContextWall cannot override that design decision.
Attacks on model weights or fine-tuning data happen before inference. ContextWall operates at inference time only.
L3 heuristics catch known semantic paraphrases. A sufficiently novel attack may score below the block threshold - you set that threshold.
Authorized access you've allowed
If your policy permits a source and the model uses that data, ContextWall enforces your policy - not a stricter one.
Honest scope beats false assurances. Defense in depth means ContextWall works alongside your model provider's safety filters, not instead of them.
ContextWall intercepts every document before it enters the context window. Here's exactly what happens.
Your agent requests a document
A web search result, internal doc, API response, or user upload - any external content.
The daemon receives the document before it enters the context window. The LLM hasn't seen anything yet.
3 Three detection layers run in sequence
400 returned to your agent. Document never reaches the LLM. Event written to the tamper-evident audit log.
Document forwarded to the LLM API. Clean context enters the prompt as normal.
No LLM in the screening path. No external calls. Your data stays on your host.
You declare what each context source is. ContextWall applies the right level of scrutiny automatically based on that tier.
Your code repos, internal wikis
Public web, user-submitted input
Applied in order from cheapest to most thorough. No external calls, no LLM inference.
Scans raw bytes for known obfuscation tricks: bidirectional control characters, zero-width characters, and spaced-letter keywords ("i g n o r e a l l"). These are invisible to the human eye but readable by the model.
Runs regex patterns against normalized text. Catches injection syntax, exposed API keys (AWS, GitHub, Anthropic), bearer tokens, and PII like emails, phone numbers, and SSNs.
Scores each message for instruction-like intent, even when the wording avoids obvious keywords. Catches paraphrases like "your previous assignment has been superseded" that bypass regex entirely.
Data never leaves your infrastructure Your context stays yours
ContextWall runs as a daemon inside your own infrastructure. Prompts, documents, and file contents are screened locally and never transmitted anywhere. The cloud control plane receives only counts and scores, never content.
Your Infrastructure Your AI agent ContextWall daemon Your LLM API calls All screening happens here. Nothing exits.
request counts violation types latency (ms) session count Never transmitted
prompt content file contents user data PII / PHI ContextWall Cloud Fleet dashboard Policy authoring Compliance reports Sees counts and scores only. Never content.
Stays in your infrastructure, always Prompt content and user messages
Retrieved documents and file contents
Model responses and completions
Personally identifiable information
Protected health information (PHI)
Sent to control plane, metadata only Request counts (blocked / allowed)
Violation types detected (e.g. "pii")
Average latency in milliseconds
Policy version acknowledgement
Prefer fully offline? Leave control_plane.url empty and ContextWall runs entirely local with no cloud dependency.
The daemon installs with pip and proxies your AI SDK calls locally. The cloud dashboard is optional and sees only aggregated metadata, never content.
Screen documents before the prompt
# 1. Install and start the daemon (runs in your infrastructure)
pip install contextwall
ctxfw start --config ctxfw.yaml
# 2. Declare your context sources in ctxfw.yaml
sources:
- id: web-search
type: web
trust_tier: untrusted
- id: internal-docs
type: confluence
trust_tier: internal
# 3. Point your AI SDK at the local daemon
export ANTHROPIC_BASE_URL=http://localhost:8080/proxy/anthropic
export ANTHROPIC_API_KEY=sk-ant-your-key # unchanged
# Every call your agent makes is now screened locally.
# Prompts and responses never leave your machine. Works with the Anthropic and OpenAI SDKs in any language. The daemon proxies requests, screens context, then forwards clean content to the real API without ever storing or transmitting your prompts.
Everything is declared in YAML. Sources, rules, and thresholds all live in a file you commit to your repo, review in a pull request, and deploy alongside your other infrastructure config.
Sources declared in config at startup, no API calls or setup scripts
Four-layer policy: fleet-wide rules down to individual repo overrides
Starter policy templates for HIPAA, SOC2, and FedRAMP included
Rules reload within 5 seconds of a file change, no restart needed
Every rule can map to a compliance control ID for audit evidence
# ctxfw.yaml
# Declare your context sources here. ContextWall registers them
# on startup automatically. No API calls, no scripts to run.
sources:
- id: brave-web-search
type: web
trust_tier: untrusted
- id: internal-confluence
type: confluence
trust_tier: internal
data_classification: sensitive
- id: fhir-api
type: api
trust_tier: regulated
data_classification: phi
owner: clinical-data-team Designed with compliance in mind Compliance coverage
ContextWall is designed so that compliance is a property of the architecture, not a checklist you complete afterwards. PHI, PII, and sensitive data are handled locally before they can ever be exposed.
Protected health information is screened locally. It never transits a third-party server.
Regulated source tier enforces that PHI can only flow between approved internal systems
Every enforcement decision is logged with a timestamp, source ID, and outcome for auditor review
Violation events are logged with timestamps, source IDs, and policy decisions for auditor review
Audit trail your reviewers can verify
Every context screening event is logged with source ID, trust tier, decision, and timestamp
Provenance chain is cryptographically linked; records cannot be altered without detection
Role-based access to the fleet dashboard; no raw context is stored or accessible anywhere
Policy rules are version-controlled; changes leave a full audit trail
Personal data stays in your jurisdiction by design
PII (email addresses, phone numbers, names) is detected and redacted before reaching the model
The daemon processes all data inside your own infrastructure, with no cross-border data transfer
Control plane receives only aggregated counts, not personal data
Supports data minimisation by design: the model sees the least data necessary to complete the task
Fully air-gapped. No external dependencies.
Daemon runs entirely within your air-gapped or VPC environment with no external dependencies
Leave control_plane.url empty for a fully offline deployment. No telemetry leaves the host.
Policy and configuration live in files you control; nothing is stored in the cloud
Enforcement continues locally even when the control plane is unreachable
All enforcement runs in your infrastructure - open source, free forever. The cloud adds fleet visibility and compliance tooling on top.
Forever. Self-host from GitHub.
Three-layer detection engine (structural, regex, heuristic)
Source trust tiers (internal / external / untrusted / regulated)
Policy-as-code in YAML, hot-reloaded in 5 s
Tamper-evident Merkle audit log
Python SDK - SafeAnthropic, SafeOpenAI
Pre-built policy packs for HIPAA, SOC2, FedRAMP
Prometheus metrics + live WebSocket enforcement feed
Runs entirely in your infrastructure. No external calls.
Self-host on GitHub Free in early access Cloud + Self-hosted Free
No credit card. Sign in to get started.
Everything in Open Source, plus:
Fleet dashboard - all daemons, status, block rates
Policy authoring UI - no YAML editing required
Compliance report generation (SOC2, HIPAA, FedRAMP exports)
One-click audit trail exports, cryptographically signed
Enforcement stays local whether or not the cloud is reachable. Prompts and documents never leave your infrastructure.
Feature Open Source Cloud Enforcement (blocking, detection) Audit log (local, tamper-evident) Policy-as-code (YAML) Python SDK Prometheus metrics Fleet dashboard (multi-daemon) Policy authoring UI Compliance report exports Email support The

[truncated]
