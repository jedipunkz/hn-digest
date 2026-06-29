---
source: "https://github.com/k08200/klorn"
hn_url: "https://news.ycombinator.com/item?id=48719602"
title: "Show HN: Klorn–I built an email firewall because every AI inbox made mine louder"
article_title: "GitHub - k08200/klorn: Klorn turns scattered work signals into approval-ready decisions. · GitHub"
author: "k08200"
captured_at: "2026-06-29T15:08:17Z"
capture_tool: "hn-digest"
hn_id: 48719602
score: 1
comments: 0
posted_at: "2026-06-29T14:18:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Klorn–I built an email firewall because every AI inbox made mine louder

- HN: [48719602](https://news.ycombinator.com/item?id=48719602)
- Source: [github.com](https://github.com/k08200/klorn)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T14:18:44Z

## Translation

タイトル: 表示 HN: Klorn – AI 受信トレイの音が大きくなったので、電子メール ファイアウォールを構築しました
記事のタイトル: GitHub - k08200/klorn: Klorn は、散在する作業信号を承認可能な意思決定に変換します。 · GitHub
説明: Klorn は、散在する作業信号を承認可能な意思決定に変換します。 - k08200/クローン

記事本文:
GitHub - k08200/klorn: Klorn は、散在する作業信号を承認可能な意思決定に変換します。 · GitHub
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
k08200
/
クローン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチタグ

s ファイルコードに移動 その他のアクションメニューを開く フォルダーとファイル
1,018 コミット 1,018 コミット .agents .agents .github .github アプリ アプリ ドキュメント ドキュメント サンプル サンプル パッケージ パッケージ Web サイト Web サイト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore BACKLOG.md BACKLOG.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md Dockerfile.api Dockerfile.api Dockerfile.web Dockerfile.web ライセンス ライセンス POC.md POC.md README.md README.md biome.json biome.json docker-compose.yml docker-compose.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml render.yaml render.yaml vercel.json vercel.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
受信トレイに注目のファイアウォール。提案エンジンではありません。
他の AI 受信トレイ ツールはすべて、各メールの横に提案カード、「AI は返信する必要があると考えています」というバッジ、レビューを待っている下書きなどのサーフェスを追加します。受信トレイの音が静かになるのではなく、大きくなります。
クローンはその逆を行います。各受信メールには、そのメールを生成した正確なバイトにバインドされた、SILENT / QUEUE / PUSH / AUTO という 1 つの分類が適用されます。チャット面はありません。提案カードはありません。 60 ツールのエージェントはありません。出力は 1 つの決定であり、ほとんどの場合、その決定は「これを見る必要はありません」です。
コードの前に教義を読んでください。それが実際の製品です。
▶️ 60 秒のウォークスルー · 📖 エディション / オープンコア境界 · 📋 変更履歴
階層
それが何を意味するか
何が起こるか
サイレント
記録され、決してレンダリングされない
この行は、真実のフィードバックのために存在します。決して見ることはありません。 (マーケティング、領収書、参考)
キュー
自分のスケジュールに合わせて復習する
キューに表示されます。プッシュも通知もありません。これがデフォルトです。
プッシュ
邪魔する価値がある
通知が発行されます。オプションで電報または 1 回の電話。
オート
リバーシブル、ハンズフリー
クラス

今日修正されました。アクション側は決定論的なフロアの背後にあります (下)。
決定方法 - そしてなぜ安価なモデルがそれを実行するのか
LLM は層を選択しません。すべての電子メールについて、信頼性、送信者信頼性、可逆性、緊急性という 4 つの特徴が 0 から 1 の間でスコア付けされ、tier-policy.ts の決定的ルールによってこれら 4 つの数値が層にマップされます。モデルは認識します。ルールを読んで単体テストで決定します。ポリシーは、ループ内のモデルなしでも監査可能です。
この分割からは 2 つの結果が生じます。
安いモデルが勝つ。モデルが一貫して 4 つの信号を読み取るだけでよい場合、フロンティア モデルの推論の深さは必要ありません。コミットされた 50 件の電子メール ゲート セット ( eval/judge-eval-set.json ) では、gemini-2.5-flash のスコアは 88% で、緊急メールのリコールは 100% で、数分の 1 のコストで gpt-4o と gemini-2.5-pro (両方 82%) を上回っています。自分で実行します: pnpm eval:judge 。
安全にフェールオープンします。 LLM がダウンしているかレートが制限されている場合でも、キーワード フォールバックはモデル呼び出しを行わずに同じ 4 つの機能を生成するため、緊急メールは引き続き通過します。モデルは認識層であり、負荷を支える層ではありません。
すべての分類は content-hash-bound です。スコアラーが読み取った正確なバイト ( from 、 subject 、snippet 、 label ) は、決定時に sha256 処理され、行とともに保存されます。読み取りパスは再ハッシュし、不一致の場合は AttendeeHashMismatchError をスローするため、後のエンリッチメントで層をサイレントに無効にすることはできません (PR #468)。
send_email 、 Permanent_delete 、 forward_external という 3 つのアクションは、ユーザーが 1 回クリックするだけでは元に戻すことができません。これらは分類子の信頼性を利用しません。 /approve 時に作成され、ペイロード バイト (正規の受信者/件名/本文に対する sha256) を固定し、実行時に検証される ActionReceipt が必要です。ドリフトがスローされ、アクションは拒否されます ( PR #480 、 #481 、 doctrine )

。
これは強制的なものであり、希望的なものではありません。executeToolCall の中央ガードは、検証済みの受信なしで到着したフロア アクションではフェイル クローズされるため、自律パスでも送信、転送、またはハード削除の承認を回避することはできません。現在、send_email は有線で呼び出し可能なケースです。残りの 2 つは、ケースが着地するまでフェールクローズで保護されています。自律エージェント自体はデフォルトで SUGGEST モード (読み取り専用ツールと提案専用) に設定されており、明示的に AUTO にオプトインした場合にのみ変更機能を取得します。
3 つの記事でアーキテクチャを説明し、トレードオフと正直な利点を示します。
GPT-4o と安価なモデルを受信トレイで争わせました。 GPT-4oが失われました。 — モデルのベイクオフ
LLM が私の電子メールを分類することを信頼していません。だから許さないんです。 — 得点者 vs 決定者
自信があれば決断するのに十分です。それだけでは十分ではありません。 — 決定論的なフロア
まだ終わっていない。これは 1 人の実際のユーザー (私) による初期の PoC です。 ICP 保持測定は現在行われているものです。 CHANGELOG では、何が固いのか、何が縫われているのかについて正直に説明しています。
「受信トレイとチャットする」ものではありません。チャット面はありません。
マルチテナントクラウドではありません。現時点では自己ホストが唯一のパスです。
オープンソースに対して機能ゲート化されていません。 docs/EDITIONS.md には、クラウドの最上位に販売されるもの (マネージド ホスティング、検証済みの Gmail スコープ、チーム ワークスペース) がリストされています。ファイアウォールの原則とコードは、両方のエディションのリポジトリに残ります。
CI チェックが赤になるのはなぜですか?スコープバジェットは意図的に失敗します。これは、変更によってルート/ページ/スキーマのサーフェスが固定予算を超えて拡大すると作動する、自ら課したラチェットであり、静かに広がるのではなく、「はい、このスコープにはそれだけの価値があります」という意識的な認識を強制します。赤いスコープバジェットは仕様によるものであり、壊れたビルドではありません。他のすべてのチェック (lint、types、tests、build、security、eval) は緑色です。
ホストされたデモ (klorn.ai) を試す
ホストされたデモは Go で実行されます

CASA Tier 2 検証を保留している間、ogle OAuth テスト モードを使用します (Klorn は Gmail の制限された gmail.modify スコープを使用します)。セルフホスティングなしで試すには、まずテスト ユーザーとして追加する必要があります。 3 つのパス (速い順):
使用したい Google メールに関する問題を開きます: 新しい oauth-tester 問題 — 私たちはあなたを追加します、コメント「追加されました」、あなたはログインします。
同じ情報を記載して k0820086@gmail.com に電子メールを送信してください。
または、ゲーティングを完全にスキップしてセルフホストします。フル機能同等であり、独自の Google OAuth 認証情報を持ち込むため、検証は必要ありません。
Google は、このモードでのテスト ユーザー スロットの上限を 100 に設定しています。 CASA 検証が出荷されると (PoC 保持測定でゲートが設定されます)、OAuth 画面が運用画面に切り替わり、ゲートが解除されます。ここにたどり着くほとんどの人にとって、セルフホストが最も早い方法です。
Klorn の最初の画面はチャットや受信箱ではなく、意思決定のキューです。散在するシグナルが収集され、3 つの質問に答えるカードとして提示されます。それは、何に注目すべきか、なぜそれが重要なのか、そしてどのようなアクションが用意されているのかということです。
意思決定キュー — 保留中の承認、コミットメント台帳、今日のリスク
メール — 優先度、返信要フラグ、添付ファイルおよび候補シグナル
カレンダー — 会議の準備状況、競合、次の内容のコンテキスト
ブリーフィング - 主要なシグナルと推奨されるアクションの毎日の概要
設定 — Google 接続、通知、実行境界、モデルとデータの制御
アクション前の承認 — メールの送信、カレンダーの変更、または外部へのプッシュには、明確な確認手順が必要です。
証拠に基づいた自動化 — すべての提案は、シグナル、推論、段階的なアクションを示しています。
漸進的な信頼 — Klorn は観察と提案モードで開始し、フィードバックを通じてより多くの自律性を獲得します。
空の状態が製品です。接続する前であっても、次のステップは明らかです。
1つ

明確な信号 — Klorn という名前は、ゲルマン語の klar (明確な) と古英語のホーン (応答する価値のある信号) に由来しています。
レイヤー
スタック
ウェブ
Next.js 15, React 19, TypeScript, Tailwind CSS
API
Fastify, TypeScript, Prisma
DB
PostgreSQL
認証
JWT, bcrypt, Google OAuth
AI
OpenAI 互換 (ローカルファースト)、OpenRouter / Gemini フェイルオーバー
リアルタイム
WebSocket, Web Push
請求
Stripe
Monorepo
pnpm workspaces
packages/
api/ Fastify API、Prisma スキーマ、エージェント/ツール オーケストレーション
Web/ Next.js アプリ: 意思決定キュー、メール、カレンダー、ブリーフィング、設定
core/ shared utilities and CLI-facing primitives
docs/ doctrine, screenshots, operational notes
Local development
git clone https://github.com/k08200/klorn.git
cd klorn
pnpmインストール
Environment files
Klorn reads two env files in local dev.データベース コンテナーが開始される前に、両方が存在する必要があります。
1. Root .env — 必要な変数を postgres + API サービスに挿入するために docker-compose によって使用されます。これがないと、 docker compose up -d postgres は、必須変数 JWT_SECRET is missing a value で失敗します。
cp .env.example .env
TOKEN_ENCRYPTION_KEY の 32 バイトの Base64 キーを生成し、それを root .env に貼り付けます。
ノード -e " console.log(require('crypto').randomBytes(32).toString('base64')) "
2. API .env — Fastify サーバーの実際のランタイム環境。
cp packages/api/.env.example packages/api/.env
Open packages/api/.env and at minimum set:
DATABASE_URL= " postgresql://klorn:klorn-local-dev@localhost:5432/klorn "
OPENROUTER_API_KEY= " " # https://openrouter.ai/keys — 無料のキーが機能します (または、以下のように完全にローカルに移動します)
WEB_URL= " http://localhost:8001 "
PORT=8000
JWT_SECRET と TOKEN_ENCRYPTION_KEY は開発ではオプションです。サーバーは警告とともに安全でないデフォルトに戻ります。再起動後も同じ開発クッキー/トークンが必要な場合は、それらを設定します。
To sync m

独自の OAuth クライアントを使用すれば、アプリの所有者であり唯一のユーザーであるため、セルフホストに Google 認証や CASA は必要ありません。
Google Cloud Console → プロジェクトを作成（または選択）します。
「APIとサービス」→「ライブラリ」→「Gmail API」と「Google Calendar API」を有効にします。
OAuth 同意画面 → ユーザーの種類 外部 → 基本事項を入力 → [テスト ユーザー] で、ログインに使用する Google アカウントを追加します。 (未検証のアプリは、テスト ユーザー リストにあるアカウントに対してのみ機能します。これが 100 スロットの上限であり、セルフホストに検証手順がないのはこのためです。クライアント上のアカウントです。)
スコープ : gmail.modify とカレンダーを追加します (Klorn はメールを読み取り、階層ラベルを書き込みます。scope-justification.md を参照)。
認証情報 → 認証情報の作成 → OAuth クライアント ID → Web アプリケーション。承認されたリダイレクト URI を http://localhost:8000/api/auth/google/callback に設定します (API ポートと GOOGLE_REDIRECT_URI と一致します)。
クライアント ID とシークレットを package/api/.env にコピーします。
GOOGLE_CLIENT_ID= " ...apps.googleusercontent.com "
GOOGLE_CLIENT_SECRET= " ... "
GOOGLE_REDIRECT_URI= " http://localhost:8000/api/auth/google/callback "
ローカル LLM (電子メールをマシン上に保持)
Klorn は、OpenAI 互換のエンドポイントと通信します。ローカル サーバー (Ollama、LM Studio、vLLM、llama.cpp) をポイントすると、最初に電子メール分類がそれに対して実行されます (クラウド キーの場合)。

[切り捨てられた]

## Original Extract

Klorn turns scattered work signals into approval-ready decisions. - k08200/klorn

GitHub - k08200/klorn: Klorn turns scattered work signals into approval-ready decisions. · GitHub
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
k08200
/
klorn
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,018 Commits 1,018 Commits .agents .agents .github .github apps apps docs docs examples examples packages packages website website .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore BACKLOG.md BACKLOG.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile.api Dockerfile.api Dockerfile.web Dockerfile.web LICENSE LICENSE POC.md POC.md README.md README.md biome.json biome.json docker-compose.yml docker-compose.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml render.yaml render.yaml vercel.json vercel.json View all files Repository files navigation
An attention firewall for your inbox. Not a suggestion engine.
Every other AI inbox tool adds a surface — a suggestion card next to each email, a badge that says "AI thinks you should reply," a draft waiting for review. The inbox gets louder, not quieter.
Klorn does the opposite. Each inbound email gets exactly one classification — SILENT / QUEUE / PUSH / AUTO — bound to the exact bytes that produced it. No chat surface. No suggestion cards. No 60-tool agent. The output is a single decision, and most of the time that decision is "you don't need to see this."
Read the doctrine before the code — that's the actual product.
▶️ 60-second walkthrough · 📖 Editions / open-core boundary · 📋 CHANGELOG
Tier
What it means
What happens
SILENT
Recorded, never rendered
The row exists for ground-truth feedback; you never see it. (marketing, receipts, FYI)
QUEUE
Review on your own schedule
Visible in the queue. No push, no notification. This is the default.
PUSH
Worth interrupting you
A notification fires. Optionally Telegram or one phone call.
AUTO
Reversible, hands-off
Classified today; the action side sits behind a deterministic floor (below).
How it decides — and why a cheap model runs it
The LLM does not pick the tier. On every email it scores four features between 0 and 1 — confidence , senderTrust , reversibility , urgency — and a deterministic rule in tier-policy.ts maps those four numbers to a tier. The model perceives; a rule you can read and unit-test decides. The policy is auditable without the model in the loop.
Two consequences fall out of that split:
A cheap model wins. When the model only has to read four signals consistently, you don't need a frontier model's reasoning depth. On the committed 50-email gate set ( eval/judge-eval-set.json ), gemini-2.5-flash scores 88% with 100% recall on urgent mail — beating gpt-4o and gemini-2.5-pro (both 82%) at a fraction of the cost. Run it yourself: pnpm eval:judge .
It fails open, safely. If the LLM is down or rate-limited, a keyword fallback produces the same four features with zero model calls, so urgent mail still gets through. The model is the perception layer, never the load-bearing one.
Every classification is content-hash-bound : the exact bytes the scorer read ( from , subject , snippet , labels ) are sha256'd at decision time and stored with the row. The read path re-hashes and throws AttentionHashMismatchError on mismatch, so a later enrichment can't silently invalidate a tier ( PR #468 ).
Three actions can't be undone with one user click — send_email , permanent_delete , forward_external . These don't ride on classifier confidence. They require an ActionReceipt minted at /approve time that pins the payload bytes (a sha256 over the canonical recipient/subject/body), verified at execute time — any drift throws and the action is refused ( PR #480 , #481 , doctrine ).
It's enforced, not aspirational: a central guard in executeToolCall fails closed on any floor action that arrives without a verified receipt, so even the autonomous path can't side-step approval to send, forward, or hard-delete. Today send_email is the wired callable case; the other two are guarded fail-closed until their cases land. The autonomous agent itself defaults to SUGGEST mode — read-only tools plus propose-only — and only gets mutating power when you explicitly opt into AUTO.
Three writeups walk through the architecture, with the tradeoffs and the honest edges:
I let GPT-4o and a cheaper model fight over my inbox. GPT-4o lost. — the model bake-off
I don't trust the LLM to classify my email. So I don't let it. — feature-scorer vs. decider
Confidence is enough to decide. It's not enough to do. — the deterministic floor
Not finished. This is an early PoC with one real user (me); ICP retention measurement is what's happening now. The CHANGELOG is honest about what's solid vs. what's stitched.
Not a "chat with your inbox" thing. There is no chat surface.
Not multi-tenant cloud. Self-host is the only path right now.
Not feature-gated against open source. docs/EDITIONS.md lists what Cloud will sell on top (managed hosting, verified Gmail scope, team workspaces) — the firewall doctrine and code stay in the repo on both editions.
Why is a CI check red? Scope Budget fails on purpose . It's a self-imposed ratchet that trips when a change grows the route / page / schema surface past a fixed budget, forcing a conscious "yes, this scope is worth it" instead of silent sprawl. A red Scope Budget is by design, not a broken build — every other check (lint, types, tests, build, security, eval) is green.
Trying the hosted demo (klorn.ai)
The hosted demo runs in Google OAuth testing mode while we hold off on CASA Tier 2 verification (Klorn uses Gmail's restricted gmail.modify scope). To try it without self-hosting, you have to be added as a test user first. Three paths, fastest first:
Open an issue with the Google email you want to use: new oauth-tester issue — we add you, comment "added", you log in.
Email k0820086@gmail.com with the same info.
Or skip the gating entirely and self-host — full feature parity, you bring your own Google OAuth credentials, no verification needed.
Google caps test-user slots at 100 in this mode. Once CASA verification ships (gated on PoC retention measurement), the OAuth screen flips to production and the gating goes away. For most people landing here, self-host is the fastest way in.
Klorn's first screen is not a chat or an inbox — it's a decision queue. Scattered signals are collected and presented as cards that answer three questions: what to look at , why it matters , and what action is ready .
Decision queue — pending approvals, the commitment ledger, today's risks
Mail — priority, reply-needed flags, attachment and candidate signals
Calendar — meeting readiness, conflicts, context for what's next
Briefing — a daily summary of top signals and recommended actions
Settings — Google connections, notifications, execution boundaries, model and data controls
Approval before action — sending mail, changing the calendar, or pushing externally requires a clear confirmation step.
Evidence-based automation — every suggestion shows the signal, the reasoning, and the staged action.
Progressive trust — Klorn starts in observe-and-suggest mode and earns more autonomy through your feedback.
The empty state is the product — even before any connection, the next step should be obvious.
One clear signal — the name Klorn comes from the Germanic klar (clear) and the Old English horn (a signal worth answering).
Layer
Stack
Web
Next.js 15, React 19, TypeScript, Tailwind CSS
API
Fastify, TypeScript, Prisma
DB
PostgreSQL
Auth
JWT, bcrypt, Google OAuth
AI
OpenAI-compatible (local-first), OpenRouter / Gemini failover
Realtime
WebSocket, Web Push
Billing
Stripe
Monorepo
pnpm workspaces
packages/
api/ Fastify API, Prisma schema, agent/tool orchestration
web/ Next.js app: decision queue, mail, calendar, briefing, settings
core/ shared utilities and CLI-facing primitives
docs/ doctrine, screenshots, operational notes
Local development
git clone https://github.com/k08200/klorn.git
cd klorn
pnpm install
Environment files
Klorn reads two env files in local dev. Both need to exist before the database container will even start.
1. Root .env — used by docker-compose to interpolate required vars into the postgres + api services. Without it, docker compose up -d postgres fails with required variable JWT_SECRET is missing a value .
cp .env.example .env
Generate a 32-byte base64 key for TOKEN_ENCRYPTION_KEY and paste it into the root .env :
node -e " console.log(require('crypto').randomBytes(32).toString('base64')) "
2. API .env — the actual runtime env for the Fastify server.
cp packages/api/.env.example packages/api/.env
Open packages/api/.env and at minimum set:
DATABASE_URL= " postgresql://klorn:klorn-local-dev@localhost:5432/klorn "
OPENROUTER_API_KEY= " " # https://openrouter.ai/keys — a free key works (or go fully local, below)
WEB_URL= " http://localhost:8001 "
PORT=8000
JWT_SECRET and TOKEN_ENCRYPTION_KEY are optional in dev — the server falls back to insecure defaults with a warning. Set them if you want the same dev cookies/tokens across restarts.
To sync mail you bring your own OAuth client — no Google verification or CASA needed for self-host, since you stay the app's owner and sole user.
Google Cloud Console → create (or pick) a project.
APIs & Services → Library → enable Gmail API and Google Calendar API .
OAuth consent screen → User type External → fill the basics → under Test users , add the Google account you'll log in with. (Unverified apps only work for accounts on the test-user list — that's the 100-slot cap, and it's why self-host has no verification step: it's your account on your client.)
Scopes : add gmail.modify and calendar (Klorn reads mail and writes tier labels; see scope-justification.md ).
Credentials → Create credentials → OAuth client ID → Web application. Set the Authorized redirect URI to http://localhost:8000/api/auth/google/callback (match your API port and GOOGLE_REDIRECT_URI ).
Copy the client ID and secret into packages/api/.env :
GOOGLE_CLIENT_ID= " ...apps.googleusercontent.com "
GOOGLE_CLIENT_SECRET= " ... "
GOOGLE_REDIRECT_URI= " http://localhost:8000/api/auth/google/callback "
Local LLM (keep your email on your machine)
Klorn speaks to any OpenAI-compatible endpoint. Point it at a local server (Ollama, LM Studio, vLLM, llama.cpp) and email classification runs against it first — cloud keys, if

[truncated]
