---
source: "https://github.com/ayauho/alchimist"
hn_url: "https://news.ycombinator.com/item?id=48665009"
title: "Show HN: Δlchimist – Local-first AI persona engine for the browser (BYOK)"
article_title: "GitHub - ayauho/alchimist: Browser-native AI content workstation with persona alchemy, intelligent targeting, and recursive refinement. BYOK. · GitHub"
author: "Ayauho"
captured_at: "2026-06-24T20:33:29Z"
capture_tool: "hn-digest"
hn_id: 48665009
score: 1
comments: 0
posted_at: "2026-06-24T20:10:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Δlchimist – Local-first AI persona engine for the browser (BYOK)

- HN: [48665009](https://news.ycombinator.com/item?id=48665009)
- Source: [github.com](https://github.com/ayauho/alchimist)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T20:10:04Z

## Translation

タイトル: HN を表示: Δlchimist – ブラウザー用のローカルファースト AI ペルソナ エンジン (BYOK)
記事タイトル: GitHub - ayauho/alchimist: ペルソナ錬金術、インテリジェントなターゲティング、再帰的洗練を備えたブラウザネイティブ AI コンテンツ ワークステーション。ビヨク。 · GitHub
説明: ペルソナ錬金術、インテリジェントなターゲティング、再帰的洗練を備えたブラウザネイティブの AI コンテンツ ワークステーション。ビヨク。 - アヤウホ/錬金術師

記事本文:
GitHub - ayauho/alchimist: ペルソナ錬金術、インテリジェントなターゲティング、再帰的洗練を備えたブラウザネイティブ AI コンテンツ ワークステーション。ビヨク。 · GitHub
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
アヤウホ
/
錬金術師
公共
通知
通知設定を変更するにはサインインする必要があります

イングス
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット コンポーネント コンポーネント アイコン アイコン libs libs modules modules services services utils utils ライセンス ライセンス README.md README.mdアンカー.js アンカー.js app.js app.js 背景.js マニフェスト.json マニフェスト.json プライバシー.html プライバシー.html シェル.html シェル.html スタイル.css スタイル.css テイルウィンド.min.css tailwind.min.css すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ΔLCHIMIST — AI コンテンツ ワークステーション
ΔLCHIMIST は、高度な完全にクライアント側のブラウザ拡張機能であり、アクティブな Web ページを根拠のあるオーサリング コンテキストに変換します。コンテンツを読み取り、解析し、選択したペルソナを通じて反映して、目的のある出力を生成します。プロのクリエイター、ゴーストライター、研究者、データ オペレーター向けに設計された ΔLCHIMIST は、すべての生成サイクルを実際のライブ ページ コンテンツに固定することで、一般的な空白のプロンプトを回避する包括的なオーサリング環境として機能します。
独自の API キー (BYOK) を使用することにより、ワークスペースは完全にブラウザーの実行境界内で動作します。中間サーバーや外部データ ルーティングはなく、独自の API エンドポイントを超えて公開されることはありません。
⚗️ コアアーキテクチャと機能
1. コンテキストに基づいた変換
Sovereign Scraper Core: アクティブなドキュメントの DOM ツリーを再帰的に走査し、すべてのフレームにわたるシャドウ ルートと iframe に自動的に降りて、ライブ ページからクリーンで構造化されたテキストを抽出します。ノイズ要素 (スクリプト、スタイル、広告、分析ペイロード) は、コンテンツが生成プロンプトにパッケージ化される前にフィルター処理されます。
ゴースト選択の無効化: テキストが複数のフレームにわたって選択されている場合、スクレーパーは各フレームの選択をマスター フレーム全体と照合します。

文章。マスター フレームに対して検証できない選択内容は、プロンプト アセンブリの前に無効化され、ファントムまたはクロスオリジン コンテキスト ブリードを防ぎます。
トポロジの単純化: 生の DOM ツリーが線形化され、ボトムアップでコンパクトな構造化された表現に単純化され、冗長な単一子のネストが崩壊し、目に見えないノードまたは質量ゼロのノードが削除されるため、AI は生の HTML ノイズではなく、クリーンで高密度のコンテンツ構造を受け取ります。
2. 認知ペルソナシステムとペルソナ錬金術
永続的なプロファイル: ペルソナは、音声、推論構文、構造パターン、制限された語彙を管理する永続的な認知マトリックスとして機能します。
コンテキスト学習: 世代サイクルごとに、アクティブなペルソナは新鮮な洞察を抽出し、関連する概念を統合し、冗長性を排除して内部の知識フレームワークを進化させます。
ペルソナ錬金術マニホールド:
合成: 提供されたソース テキストからまったく新しい認知プロファイルを抽出します。
変異: 明確な文体または概念的な影響を注入することで、ペルソナの中核となるプロファイルを変換します。
交配: 2 つの個別のペルソナを、両方の親ベクトルから言語特性を継承するハイブリッド子マトリックスに結合します。
3. 多段階のコンテンツ戦略
コメント: 特定のページの議論に向けて、高度にターゲットを絞った、状況に応じた戦略的なインタラクションを策定します。
再投稿: 投稿または記事の核となるテキストを抽出し、引用符を適用し、多層的な分析視点を追加します。
リライト: 選択したペルソナの正確なスタイルの特徴を通じて既存のマテリアルを処理します。
プロモーション投稿: 独自の構造化プロファイル インテリジェンスをドメイン テーマと合成して、作成者が検証した投稿を草案します。
長編記事 (2 段階のパイプライン):
ステージ 1 (マテリアルの準備): 複数のセッションにわたってリソースの断片を収集します。

n 組み込みの研究仲間が次のステップについてアドバイスします。
ステージ 2 (生成): 収集した資料を、選択した物語モデルを中心に構成された、まとまった長文のエッセイに編集します。
単一出力: 高度に最適化された単一の結果に焦点を当てて生成します。
マトリックス モード: 出力の 3 つの独立したスタイル バリアントを同時に生成します。それぞれには、さらなる反復のための 3 つのターゲットを絞った改良提案が伴います。
Nexus スレッド: 長編出版用に、3 つの独立した (または半依存した) マルチパート シリアル ブロックの一貫したシーケンスを作成します。
5. 高度なリファインメントおよびプロトコルスイート
動的メトリクスの再調整: 各世代は、自動的に割り当てられた評価軸 (明瞭さ、特異性、権威性、感情の重みなど) にわたってスコア付けされます。各メトリクスの上/下コントロールを使用して、その軸に沿ってシフトされたターゲットの再生成をトリガーします。ページ コンテキストから完全にカスタムの評価基準を合成、突然変異、または交配するための Metrics Alchemy をサポートします。
命令と指示:
命令: すべての世代に優先度の高い永続的な必須の編集上の制約を適用します。これは、戦略やペルソナに関係なく常に尊重する必要がある、交渉の余地のないトーン ルール、禁止されたトピック、またはフォーマット要件に役立ちます。
Directive Citadel: 履歴のある世代ごとの自由形式の命令フィールド。命令とは異なり、ディレクティブは、明示的に要求されない限り、ペルソナの意図をオーバーライドすることなく単一の世代を形成するワンショットのリクエストです。
言語および書式設定プロトコル:
Void Source Auditor: 標準の AI 生成マーカーを最小限に抑え、テキストの自然さを最大限に高めます。
Cognitive Origin Auditor: 7 つの法医学的側面にわたって言語の差異をスキャンし、入力コンテキストが AI によって生成された確率を推定します。
Twitter ショート: 文字制限を強制する

プラットフォームのコンプライアンスのために。
テーマのタグ付けと構造の強化: プラットフォームに最適化されたハッシュタグ配列を生成し、太字の強調ハイライトを適用します。
カスタム署名: 静的なパーソナライズされた透かし文字列を生成境界でシームレスに挿入します。
6. リレーショナル インテリジェンスとソーシャル スケーリング
ピア ターゲティング マトリックス: 対象となる個人から構造化された公開データ プロファイルを収集し、カスタム インタラクション目標を添付します。生成される出力は、誰に向けて、どのような目的に設定されているかに応じて変化します。
キャラクターとアーキタイプ: 特化したドメインの視点 (歴史上の人物、顔の見えない専門家) を自分自身のプロフィールまたは対象となる同僚のプロフィールに直接注入し、生成コンテキストに集中的な推論の深さを追加します。
プロフェッショナル バンドル: ペルソナ、キャラクター、意図の最適化されたマトリクスを備えた、準備された運用フレームワークを瞬時に切り替えます。
多目的インジェスト: ファイルを生成レイヤー ( .pdf 、 .docx 、 .html 、 .txt 、 .md 、 .json ) に直接アップロードします。アップロードされたドキュメントは、AI が生成中にページ コンテンツとともにマイニングする補足コンテキストとして利用できます。
🛠️ インフラストラクチャとデータ アーキテクチャ
BYOK キーとモデルのローテーション: 自動リアルタイム フォールバック ルーティングにより、最大 5 つの同時 API キーをサポートします。プライマリ スロットが 429 クォータ枯渇制限に達した場合、モデル エンジンは、現在の実行状態を中断することなく、次の利用可能なスロットにシームレスにローテーションするか、階層間で安全にダウングレードします。
Workspace Treasury Snapshots: 完全なトランザクション バックアップ ユーティリティ。ユーザーは、運用環境全体 (プリセット、記事、キー、ペルソナ、命令、添付ファイル、カスタム メトリクス、ピア、インテント、文字、アーキタイプ、スキーム) を暗号化されたスナップショットにエクスポートできます。堅牢な安全ルールにより、クローなしで状態をマージまたはロールバックできます。

既存の構造物を破壊します。
ΔLCHIMIST は、Google AI Studio が提供する標準 API キーで直接実行されます。
Google AI スタジオに移動します。
[API キーの取得] を選択し、続いて [新しいキーの作成] を選択します。
このリポジトリをローカル パスに複製または抽出します。
Google Chrome を開き、 chrome://extensions/ に移動します。
右上にある開発者モードをオンに切り替えます。
左上にある「解凍してロード」をクリックし、この拡張機能を含むフォルダーのパスを選択します。
ツールバーの ΔLCHIMIST アクション バッジをタップし、 [設定] → [API とモデル] をクリックして、AI Studio キーを貼り付けます。
ΔLCHIMIST は、Chrome ウェブストアから直接入手することもできます。
ΔLCHIMISTをインストールする
🔒 プライバシーとセキュリティのブループリント
絶対的なクライアント側実行: ΔLCHIMIST には仲介サーバーがありません。すべてのデータは、Chrome の分離されたローカル ストレージ サンドボックス内に存在します。
ソブリン ネットワーク ブリッジ: ネットワーク クエリは、自己プロビジョニングされた認証情報を使用して、ブラウザ スレッドと直接 Gemini API エンドポイントの間で排他的にルーティングされます。
ペルソナ錬金術、インテリジェントなターゲティング、再帰的洗練を備えたブラウザネイティブの AI コンテンツ ワークステーション。ビヨク。
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

Browser-native AI content workstation with persona alchemy, intelligent targeting, and recursive refinement. BYOK. - ayauho/alchimist

GitHub - ayauho/alchimist: Browser-native AI content workstation with persona alchemy, intelligent targeting, and recursive refinement. BYOK. · GitHub
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
ayauho
/
alchimist
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits components components icons icons libs libs modules modules services services utils utils LICENSE LICENSE README.md README.md anchor.js anchor.js app.js app.js background.js background.js manifest.json manifest.json privacy.html privacy.html shell.html shell.html styles.css styles.css tailwind.min.css tailwind.min.css View all files Repository files navigation
ΔLCHIMIST — AI Content Workstation
ΔLCHIMIST is an advanced, fully client-side browser extension that turns any active web page into a grounded authoring context — reading, parsing, and reflecting its content back through your chosen persona to generate purposeful output. Designed for professional creators, ghostwriters, researchers, and data operators, ΔLCHIMIST serves as a comprehensive authoring environment that avoids generic blank-box prompting by anchoring every generation cycle to real, live page content.
By bringing your own API keys (BYOK), the workspace operates completely within your browser's execution boundaries — no intermediary servers, no external data routing, no exposure beyond your own API endpoint.
⚗️ Core Architecture & Features
1. Context-Grounded Transmutation
Sovereign Scraper Core: Recursively traverses the active document's DOM tree, automatically descending into Shadow Roots and iframes across all frames, to extract clean, structured text from the live page. Noise elements (scripts, styles, ads, analytics payloads) are filtered before the content is packaged into the generation prompt.
Ghost Selection Neutralization: When text is selected across multiple frames, the scraper cross-checks each frame's selection against the master frame's full text. Any selection that cannot be verified against the master frame is neutralized before prompt assembly, preventing phantom or cross-origin context bleed.
Topological Simplification: The raw DOM tree is linearized and bottom-up simplified into a compact, structured representation — collapsing redundant single-child nesting and purging invisible or zero-mass nodes — so the AI receives a clean, high-density content structure rather than raw HTML noise.
2. Cognitive Persona System & Persona Alchemy
Persistent Profiles: Personas act as persistent cognitive matrices governing voice, reasoning syntax, structural patterns, and restricted vocabularies.
Contextual Learning: With every generation cycle, the active persona extracts fresh insights, synthesizes related concepts, and prunes redundancies to evolve its internal knowledge framework.
Persona Alchemy Manifold:
Synthesize: Extract a completely new cognitive profile from any provided source text.
Mutate: Transform a persona's core profile by injecting a distinct stylistic or conceptual influence.
Crossbreed: Combine two discrete personas into a hybrid child matrix that inherits linguistic traits from both parent vectors.
3. Multi-Stage Content Strategies
Comment: Formulate highly targeted, contextually strategic interactions geared towards specific page arguments.
Repost: Pull the core text of a post or article, apply quotes, and append a multi-layered analytical perspective.
Rewrite: Process existing materials through your selected persona's exact stylistic signature.
Promotional Post: Synthesize your own structured profile intelligence with domain themes to draft author-validated posts.
Long-Form Article (Two-Stage Pipeline):
Stage 1 (Material Preparation): Gather resource fragments across multiple sessions while an embedded research companion advises on the next step.
Stage 2 (Generation): Compile gathered materials into a cohesive long-form essay structured around a chosen narrative model.
Single Output: Focuses the generation on a single highly-optimized result.
Matrix Mode: Generates three independent stylistic variants of the output simultaneously, each accompanied by three targeted refinement suggestions for further iteration.
Nexus Threads: Creates a coherent sequence of three independent (or semi-dependent), multi-part serial blocks for long-form publishing.
5. Advanced Refinement & Protocols Suite
Dynamic Metrics Recalibration: Each generation is scored across automatically assigned evaluation axes (e.g. Clarity, Specificity, Authority, Emotional Weight). Use the up/down controls on each metric to trigger targeted regenerations shifted along that axis. Supports Metrics Alchemy to synthesize, mutate, or crossbreed entirely custom evaluation criteria from page context.
Imperatives & Directives:
Imperatives: Persistent, mandatory editorial constraints injected into every generation at high priority — useful for non-negotiable tone rules, forbidden topics, or format requirements that must always be respected regardless of strategy or persona.
Directive Citadel: A per-generation freeform instruction field with history. Unlike imperatives, directives are one-shot requests that shape a single generation without overriding persona intent unless explicitly asked.
Linguistic & Formatting Protocols:
Void Source Auditor: Minimizes standard AI generation markers to maximize text naturalness.
Cognitive Origin Auditor: Scans language variance across seven forensic dimensions to estimate the probability that the input context is AI-generated.
Twitter Short: Forces character constraint for platform compliance.
Thematic Tagging & Structural Enhancements: Generates platform-optimized hashtag arrays and applies bold emphasis highlights.
Custom Signatures: Injects static personalized watermark strings seamlessly at the generation boundary.
6. Relational Intelligence & Social Scaling
Peer Targeting Matrix: Harvest structured public data profiles from target individuals and attach custom interaction goals. Generation outputs adapt to who you are addressing and your stated objective.
Characters & Archetypes: Inject specialized domain perspectives (historical figures, faceless experts) directly into your own profile or a target peer's profile to add focused reasoning depth to the generation context.
Professional Bundles: Instantly swap between prepared operational frameworks carrying optimized matrices of personas, characters, and intentions.
Multipurpose Ingestion: Upload files directly into the generation layer ( .pdf , .docx , .html , .txt , .md , .json ). Uploaded documents are available as supplementary context that the AI mines alongside page content during generation.
🛠️ Infrastructure & Data Architecture
BYOK Key & Model Rotation: Supports up to 5 concurrent API keys with automatic, real-time fallback routing. If a primary slot trips a 429 Quota Exhausted limit, the model engine seamlessly rotates to the next available slot or safely downgrades across tiers without disrupting current execution states.
Workspace Treasury Snapshots: Fully transactional backup utility. Allows users to export their entire operational environment (presets, articles, keys, personas, imperatives, attachments, custom metrics, peers, intents, characters, archetypes, and schemes) into an encrypted snapshot, with robust safety rules to merge or roll back states without clobbering existing structures.
ΔLCHIMIST runs directly on standard API keys provided by Google AI Studio .
Navigate to Google AI Studio .
Select Get API Key followed by Create New Key .
Clone or extract this repository into your local path.
Open Google Chrome and navigate to chrome://extensions/ .
Toggle Developer mode on at the top right.
Click Load unpacked at the top left and select the folder path containing this extension.
Tap the ΔLCHIMIST action badge in your toolbar, click Settings → API & Model , and paste your AI Studio key.
ΔLCHIMIST is also available directly from the Chrome Web Store:
Install ΔLCHIMIST
🔒 Privacy & Security Blueprint
Absolute Client-Side Execution: ΔLCHIMIST has no middleman servers. All data live within Chrome's isolated local storage sandbox.
Sovereign Network Bridges: Network queries route exclusively between your browser thread and the direct Gemini API endpoint using your self-provisioned credential.
Browser-native AI content workstation with persona alchemy, intelligent targeting, and recursive refinement. BYOK.
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
