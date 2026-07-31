---
source: "https://github.com/orgs/modelcontextprotocol/discussions/824"
hn_url: "https://news.ycombinator.com/item?id=49121724"
title: "Show HN: Secure Bearer JWT and JWKS pattern for headless AI agents (MCP #824)"
article_title: "Security: Bearer JWT + JWKS for agent callers (no master secret in the LLM path) · modelcontextprotocol · Discussion #824 · GitHub"
author: "MawyxxY"
captured_at: "2026-07-31T11:20:16Z"
capture_tool: "hn-digest"
hn_id: 49121724
score: 1
comments: 0
posted_at: "2026-07-31T11:17:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Secure Bearer JWT and JWKS pattern for headless AI agents (MCP #824)

- HN: [49121724](https://news.ycombinator.com/item?id=49121724)
- Source: [github.com](https://github.com/orgs/modelcontextprotocol/discussions/824)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T11:17:59Z

## Translation

タイトル: HN の表示: ヘッドレス AI エージェントのセキュア ベアラー JWT および JWKS パターン (MCP #824)
記事のタイトル: セキュリティ: エージェント呼び出し元のベアラー JWT + JWKS (LLM パスにマスター シークレットなし) · モデルコンテキストプロトコル · ディスカッション #824 · GitHub
説明: セキュリティ: エージェント呼び出し元のベアラー JWT + JWKS (LLM パスにマスター シークレットなし)

記事本文:
セキュリティ: エージェント呼び出し元のベアラー JWT + JWKS (LLM パスにマスター シークレットなし) · モデルコンテキストプロトコル · ディスカッション #824 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
モデルコンテキストプロトコル
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
モデルコンテキストプロトコル
セキュリティ: ベアラー JWT + JWKS f

またはエージェントの呼び出し元 (LLM パスにマスター シークレットがない)
#824
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
この投稿は MCP コミュニティ ガイドラインに従っています
セキュリティ: エージェント呼び出し元のベアラー JWT + JWKS (LLM パスにマスター シークレットなし)
MCP ツールを呼び出すヘッドレス エージェントは通常、2 つのデフォルトのいずれかを取得します。どちらも深刻な脅威モデルに失敗します。
エージェント環境の静的 API キー — プロンプトインジェクション / ログリーク ⇒ フリート全体での共有爆発半径。
人間の OAuth — ブラウザも同意クリックも必要ありません。その後、チームはサービスユーザーのハッキングを出荷するか、認証をオフのままにします。
形状の修正: エージェントが、この MCP ホスト名にバインドされた短期間の RS256 JWT をミントする → 認可: ベアラー → リソース サーバーが JWKS ( aud / exp / domain ) 経由でローカルに検証する → sub によって認可する。有効期間の長い mint 資格情報はエージェント ホスト上に留まり、RS に到達することはありません。
これが、LIME がマシン専用 AS として実装するものです。JWT を発行し、JWKS を公開し、側で検証を続けます (JWKS キャッシュ後のインプロセス暗号化。ツールのホット パスには LIME RTT はありません)。
静的キー vs 人間の OAuth vs JWKS パスポート
エージェント ホスト: 不透明なミント シークレット (env / シークレット マネージャー) → MCP には送信されません
エージェント → LIME: mint RS256 JWT { aud=mcp、domain=<mcp-host>、TTL≈ minutes }
エージェント → MCP: 権限: ベアラー <jwt>
MCP RS: JWKS 検証 → サブによる承認
環境の後ろのゲート (ローカル デモの場合はオフ、本番環境の場合はオン)。最初の JWKS フェッチはネットワークです。ウォームキャッシュ検証はローカルです。
JWKS: https://lime.pics/api/v1/core/.well-known/jwks.json
AS メタデータ: https://lime.pics/.well-known/oauth-authorization-server
RS SDK: https://github.com/Mawyxx/lime-mcp-server-sdk ( TokenVerifier 、 Expected_domain )
静的キーがすでに脅威モデルと一致している場合は、無視してください。エージェントの認証が審査に入る場合、これは通常生き残るパターンです

そうだね。
1
投票するにはログインする必要があります
すべての反応
返信:
コメント0件
見出し
太字
イタリック体
引用
コード
リンク
番号付きリスト
順序なしリスト
タスクリスト
ファイルを添付する
言及
参考資料
メニュー
見出し
読み込み中にエラーが発生しました。このページをリロードしてください。
新しい保存済み返信を作成する
👍
1
親指を立てた絵文字で反応した
👎
1
サムダウンの絵文字で反応した
😄
1
笑いの絵文字で反応した
🎉
1
万歳の絵文字で反応しました
😕
1
混乱した絵文字で反応した
❤️
1
ハートの絵文字で反応しました
🚀
1
ロケットの絵文字で反応しました
👀
1
目の絵文字で反応した
フッター
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Security: Bearer JWT + JWKS for agent callers (no master secret in the LLM path)

Security: Bearer JWT + JWKS for agent callers (no master secret in the LLM path) · modelcontextprotocol · Discussion #824 · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
modelcontextprotocol
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Model Context Protocol
Security: Bearer JWT + JWKS for agent callers (no master secret in the LLM path)
#824
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
This post follows the MCP community guidelines
Security: Bearer JWT + JWKS for agent callers (no master secret in the LLM path)
Headless agents calling MCP tools usually get one of two defaults — both fail a serious threat model:
Static API key in the agent env — prompt-injection / log leak ⇒ shared blast radius across the fleet.
Human OAuth — no browser, no consent click; teams then ship a service-user hack or leave auth off.
Fix shape: agent mints a short-lived RS256 JWT bound to this MCP hostname → Authorization: Bearer → resource server verifies locally via JWKS ( aud / exp / domain ) → authorize by sub . Long-lived mint credential stays on the agent host and never hits the RS.
That’s what LIME implements as a machine-only AS: issue the JWT, publish JWKS, keep verify on your side (in-process crypto after JWKS cache; no LIME RTT on the tool hot path).
Static key vs human OAuth vs JWKS passport
Agent host: opaque mint secret (env / secret manager) → never sent to MCP
Agent → LIME: mint RS256 JWT { aud=mcp, domain=<mcp-host>, TTL≈minutes }
Agent → MCP: Authorization: Bearer <jwt>
MCP RS: JWKS verify → authorize by sub
Gate behind env ( off for local demos, on for prod). First JWKS fetch is network; warm-cache verify is local.
JWKS: https://lime.pics/api/v1/core/.well-known/jwks.json
AS metadata: https://lime.pics/.well-known/oauth-authorization-server
RS SDK: https://github.com/Mawyxx/lime-mcp-server-sdk ( TokenVerifier , expected_domain )
If static keys already match the threat model, ignore. If agent auth is heading into review, this is the pattern that usually survives it.
1
You must be logged in to vote
All reactions
Replies:
0 comments
Heading
Bold
Italic
Quote
Code
Link
Numbered list
Unordered list
Task list
Attach files
Mention
Reference
Menu
Heading
There was an error while loading. Please reload this page .
Create a new saved reply
👍
1
reacted with thumbs up emoji
👎
1
reacted with thumbs down emoji
😄
1
reacted with laugh emoji
🎉
1
reacted with hooray emoji
😕
1
reacted with confused emoji
❤️
1
reacted with heart emoji
🚀
1
reacted with rocket emoji
👀
1
reacted with eyes emoji
Footer
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
