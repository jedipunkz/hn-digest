---
source: "https://github.com/MJohnstonAI/ai-billing-delegation"
hn_url: "https://news.ycombinator.com/item?id=48464292"
title: "Show HN: AI has no OAuth equivalent for billing delegation – here's an open spec"
article_title: "GitHub - MJohnstonAI/ai-billing-delegation: AI Billing Delegation Standard (ABDS) — The missing OAuth layer for AI subscription portability. An open proposal for all AI providers. · GitHub"
author: "marcjohnston"
captured_at: "2026-06-09T18:50:59Z"
capture_tool: "hn-digest"
hn_id: 48464292
score: 1
comments: 0
posted_at: "2026-06-09T17:26:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI has no OAuth equivalent for billing delegation – here's an open spec

- HN: [48464292](https://news.ycombinator.com/item?id=48464292)
- Source: [github.com](https://github.com/MJohnstonAI/ai-billing-delegation)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T17:26:10Z

## Translation

タイトル: Show HN: AI には請求委任に相当する OAuth がありません – これはオープン仕様です
記事のタイトル: GitHub - MJohnstonAI/ai-billing-delegation: AI Billing Delegation Standard (ABDS) — AI サブスクリプションのポータビリティに欠落している OAuth レイヤー。すべての AI プロバイダーに対するオープンな提案。 · GitHub
説明: AI Billing Delegation Standard (ABDS) — AI サブスクリプションのポータビリティに欠落している OAuth レイヤー。すべての AI プロバイダーに対するオープンな提案。 - MJohnstonAI/ai-billing-delegation

記事本文:
GitHub - MJohnstonAI/ai-billing-delegation: AI Billing Delegation Standard (ABDS) — AI サブスクリプションのポータビリティに欠落している OAuth レイヤー。すべての AI プロバイダーに対するオープンな提案。 · GitHub
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
MジョンストンAI
/
ai-請求-委任
公共
通知

カチオン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
MJohnstonAI/ai-billing-delegation
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
27 コミット 27 コミット AI_CONTRIBUTIONS AI_CONTRIBUTIONS ディスカッション ディスカッション ABDS_executive_technical_brief.pptx ABDS_executive_technical_brief.pptx ABDS_executive_technical_brief_fixed.pdf ABDS_executive_technical_brief_fixed.pdf COTRIBUTING.md COTRIBUTING.md EXAMPLES.md EXAMPLES.md FLOWS.md FLOWS.md RATIONALE.md RATIONALE.md README.md README.md SPEC.md SPEC.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI 請求委任標準 (ABDS)
AI サブスクリプションのポータビリティに OAuth レイヤーが欠落しています。
Anthropic、OpenAI、Google などの主要な AI プロバイダーはすべて、トークンごとに課金される開発者 API アクセスを提供しています。これは B2B ソフトウェアで機能します。消費者向けモバイルアプリでは完全に機能しません。
Claude、GPT-5.x、または Gemini でコンシューマ アプリを構築する開発者は、すべてのユーザーのすべての API コストを自分で支払う必要があり、そのコストをエンド ユーザー自身のサブスクリプションに委任するメカニズムはありません。結果:
開発者はウイルスの増殖による無制限の財務リスクを吸収します
ハッカーは公開された API キーを悪用して無料で使用できる
BYOK (Bring Your Own Key) は技術的に一般ユーザーにはアクセスできません
実行可能な消費者向け AI アプリは、クレジット/ペイウォール システムを上部にボルトで固定しない限り、構造的に安全に構築することは不可能です
これは不足しているインフラストラクチャ層であり、不足している機能ではありません。
それを明らかにする例え話
サードパーティのアプリが音楽をストリーミングしたい場合、ストリームごとに Spotify に料金を支払うことはありません。ユーザーは、OAuth 経由で自分の Spotify アカウントで認証するように求められます。アプリはスコープ付きトークンを取得します。使用量はユーザーのサブスクリプションに対して請求されます。開発者は自由に構築します。
これはAIには存在しません。それはすべきです。

AI Billing Delegation Standard (ABDS) は OAuth 2.0 を拡張し、ユーザーが AI プロバイダーとのユーザー自身のサブスクリプションから AI API クォータを消費することをサードパーティ アプリケーションに承認できるようにします。
ユーザーがサードパーティのアプリを開く
↓
アプリは AI プロバイダーの同意画面にリダイレクトされます
(「MyApp はあなたの Anthropic クォータの使用を希望しています — 最大 100 クエリ/月」)
↓
ユーザーが認証し、独自の上限を設定します
↓
アプリはスコープ指定された委任トークンを受け取ります
↓
アプリは委任トークンを使用して API 呼び出しを行う
↓
ユーザーのサブスクリプションに対して使用量が請求される
開発者は何も支払わない
提案された OAuth スコープ
範囲
説明
ai.quota.delegate
ユーザーはアプリが割り当てを消費することを承認します
ai.quota.read
アプリは残りの割り当てを消費せずに読み取ることができます
ai.quota.limit
ユーザーはアプリが消費できるクォータの量を制限できます
ai.quota.revoke
ユーザーはいつでも委任を取り消すことができます
これが業界標準でなければならない理由
1 つのベンダーによる独自の実装では何も解決されません。開発者は移植性を必要としています。 ABDS は次のように一律に採用される必要があります。
Claude Pro サブスクリプションを持つユーザーは、ABDS 準拠のアプリに Claude クォータの使用を許可できます。
ChatGPT Plus を使用しているユーザーは、OpenAI を利用したアプリに対して同じことを行うことができます
開発者はベンダーごとではなく、標準に基づいて一度ビルドすればよい
月額 20 ドルの AI 加入者のほとんどは、割り当ての 15 ～ 20% を使用しています。それはアイドル状態にある未使用の容量です。 ABDS を使用すると、ユーザーはそのクォータを、完全な制御と取り消しを使用して、選択したアプリで動作させることができます。 AI プロバイダーはユーザーとの請求関係を保持します。コストリスクがゼロになるため、開発者のエコシステムは劇的に拡大します。
完全な経済的および競争的根拠については、 RATIONALE.md を参照してください。
OAuth 2.0拡張モデルが提案される
技術仕様の草案作成
この問題に遭遇した開発者: このリポジトリにスターを付けて、問題点でストーリーを共有してください
AI企業社員：y

私たちの内部擁護はあなたが思っている以上に重要です
標準エンジニア : SPEC.md での OAuth 2.0 拡張機能の形式化を支援
AI エージェント : AI_CONTRIBUTIONS/README.md を参照
ファイル
説明
SPEC.md
技術仕様
根拠.md
ビジネスと経済のケース
FLOWS.md
OAuth フロー図
例.md
開発者コードの例
貢献.md
貢献方法
AI_CONTRIBUTIONS/README.md
AI エージェントの貢献に関するガイド
ディスカッション/オープンクエスチョン.md
設計に関するオープンな質問
著者
Marc Johnston ( @MJohnstonAI ) によって提案されました。NeuroSync AI Dynamics の創設者、Anthropic、OpenAI、OpenRouter に基づいて構築されたコンシューマ AI アプリケーションの開発者であり、実際のアプリを構築する際にこの壁に直接ぶつかった実践者です。
MIT — この提案は意図的に誰でも採用、実装、拡張できるように公開されています。
AI Billing Delegation Standard (ABDS) — AI サブスクリプションのポータビリティに欠けている OAuth レイヤー。すべての AI プロバイダーに対するオープンな提案。
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

AI Billing Delegation Standard (ABDS) — The missing OAuth layer for AI subscription portability. An open proposal for all AI providers. - MJohnstonAI/ai-billing-delegation

GitHub - MJohnstonAI/ai-billing-delegation: AI Billing Delegation Standard (ABDS) — The missing OAuth layer for AI subscription portability. An open proposal for all AI providers. · GitHub
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
MJohnstonAI
/
ai-billing-delegation
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
MJohnstonAI/ai-billing-delegation
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits AI_CONTRIBUTIONS AI_CONTRIBUTIONS DISCUSSIONS DISCUSSIONS ABDS_executive_technical_brief.pptx ABDS_executive_technical_brief.pptx ABDS_executive_technical_brief_fixed.pdf ABDS_executive_technical_brief_fixed.pdf CONTRIBUTING.md CONTRIBUTING.md EXAMPLES.md EXAMPLES.md FLOWS.md FLOWS.md RATIONALE.md RATIONALE.md README.md README.md SPEC.md SPEC.md View all files Repository files navigation
AI Billing Delegation Standard (ABDS)
The missing OAuth layer for AI subscription portability.
Every major AI provider — Anthropic, OpenAI, Google — offers developer API access billed per token. This works for B2B software. It breaks completely for consumer mobile apps.
A developer who builds a consumer app on Claude, GPT-5.x, or Gemini must pay all API costs themselves, for all users, with no mechanism to delegate those costs to the end user's own subscription. The result:
Developers absorb unlimited financial risk from viral growth
Hackers exploit exposed API keys for free usage
BYOK (Bring Your Own Key) is technically inaccessible to ordinary users
Viable consumer AI apps are structurally impossible to build safely without a credit/paywall system bolted on top
This is a missing infrastructure layer — not a missing feature.
The Analogy That Makes It Obvious
When a third-party app wants to stream music, it does not pay Spotify per stream. It asks the user to authenticate with their own Spotify account via OAuth. The app gets a scoped token. Usage is billed against the user's subscription. The developer builds freely.
This does not exist for AI. It should.
The AI Billing Delegation Standard (ABDS) extends OAuth 2.0 to allow a user to authorize a third-party application to consume AI API quota from the user's own subscription with an AI provider.
User opens third-party app
↓
App redirects to AI provider consent screen
("MyApp wants to use your Anthropic quota — up to 100 queries/month")
↓
User authenticates and sets their own cap
↓
App receives scoped delegation token
↓
App makes API calls using delegation token
↓
Usage billed against user's subscription
Developer pays nothing
Proposed OAuth Scopes
Scope
Description
ai.quota.delegate
User authorizes an app to consume their quota
ai.quota.read
App can read remaining quota without consuming
ai.quota.limit
User can cap how much quota an app may consume
ai.quota.revoke
User can revoke delegation at any time
Why This Must Be An Industry Standard
A proprietary implementation by one vendor solves nothing — developers need portability. ABDS should be adopted uniformly so that:
A user with a Claude Pro subscription can authorize any ABDS-compliant app to use their Claude quota
A user with ChatGPT Plus can do the same for OpenAI-backed apps
Developers can build once against the standard, not per-vendor
Most $20/month AI subscribers use 15–20% of their quota. That is unused capacity sitting idle. ABDS lets users put that quota to work in apps they choose, with full control and revocation. The AI provider retains the billing relationship with the user. The developer ecosystem expands dramatically because cost risk drops to zero.
For the full economic and competitive rationale, see RATIONALE.md .
OAuth 2.0 extension model proposed
Technical specification drafted
Developers who have hit this problem: star this repo and share your story in Issues
AI company employees : your internal advocacy matters more than you think
Standards engineers : help formalize the OAuth 2.0 extension in SPEC.md
AI agents : see AI_CONTRIBUTIONS/README.md
File
Description
SPEC.md
Technical specification
RATIONALE.md
Business and economic case
FLOWS.md
OAuth flow diagrams
EXAMPLES.md
Developer code examples
CONTRIBUTING.md
How to contribute
AI_CONTRIBUTIONS/README.md
Guide for AI agent contributions
DISCUSSIONS/open-questions.md
Open design questions
Author
Proposed by Marc Johnston ( @MJohnstonAI ) — founder of NeuroSync AI Dynamics, developer of consumer AI applications built on Anthropic, OpenAI, and OpenRouter, and a practitioner who hit this wall directly while building real apps.
MIT — this proposal is intentionally open for anyone to adopt, implement, or extend.
AI Billing Delegation Standard (ABDS) — The missing OAuth layer for AI subscription portability. An open proposal for all AI providers.
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
