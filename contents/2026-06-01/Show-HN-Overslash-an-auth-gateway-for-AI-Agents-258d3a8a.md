---
source: "https://www.overslash.com/"
hn_url: "https://news.ycombinator.com/item?id=48344584"
title: "Show HN: Overslash – an auth gateway for AI Agents"
article_title: "Overslash — Actions & auth gateway for AI agents"
author: "arturogoosnargh"
captured_at: "2026-06-01T00:33:12Z"
capture_tool: "hn-digest"
hn_id: 48344584
score: 2
comments: 1
posted_at: "2026-05-31T10:34:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Overslash – an auth gateway for AI Agents

- HN: [48344584](https://news.ycombinator.com/item?id=48344584)
- Source: [www.overslash.com](https://www.overslash.com/)
- Score: 2
- Comments: 1
- Posted: 2026-05-31T10:34:23Z

## Translation

タイトル: Show HN: Overslash – AI エージェントの認証ゲートウェイ
記事のタイトル: Overslash — AI エージェントのアクションと認証ゲートウェイ
説明: エージェントとエージェントが接続するすべてのサービスの間に 1 つのゲートウェイ。 OAuth、シークレット、承認、監査を処理します。

記事本文:
Overs / ash ベータ版 機能 仕組み 統合 価格 オープンソース インストール JP ▾ ☀ GitHub ↗ オープンクラウド ↗ エージェントが行動します。オーバー/アッシュコントロール。
エージェントが接触するすべてのサービスへの 1 つのゲートウェイ。拒否するか、1 回だけ承認するか、すべてを承認するかを決定します。
エージェントと他のすべての間にゲートウェイが 1 つあります。
オーバースラッシュはエージェントと外界の間に位置します。シークレット、OAuth、MCP、権限チェーン、人間による承認、認証された HTTP 実行を処理します。エージェントは鍵を持っていません。息を止めないんですね。
すべてのエージェントは ID、親、および爆発範囲を取得します。サブエージェントは、ルールを絞り込むまで、親からルールを継承します。
エージェントが初めて許可を求めるときは、承認または拒否します。選択したスコープラダーで「許可して記憶」します。
クライアント ID、トークン、署名キーを 1 か所に保管。エージェントごとにローテーション、取り消し。資格情報がエージェントのコンテキスト ウィンドウに影響することはありません。
すべての電話、すべてのバブルアップ、すべての拒否。ストリーミング、検索可能、エクスポート可能。
クリックするだけで任意の MCP クライアントを登録できます。Claude Desktop は独自のものです。 Overslash は、ツール、資格情報、承認を仲介します。エージェントはただ尋ねます。
サブエージェントは、未知のスコープに到達すると、その親にバブルアップします。親は人間に話しかけます。拒否は第一級であり、記録されます。
エージェントは尋ねます。オーバースラッシュが仲介します。サービスが応答します。
MCP 対応エージェント (Claude Code、Overfolder、OpenClaw、または独自のエージェント) は、Overslash に 1 回接続します。 Overslash は資格情報を保持し、ルールを適用し、認証および監査されたリクエストを各サービスに渡します。
エージェント:ヘンリー エージェント:デプロイボット エージェント:リサーチボット ゲートウェイオーバー/アッシュ ルールをチェックします。資格情報を注入します。親エージェントや人間にとっては未知のバブル。すべての通話を録音します。認証済みリクエスト ↓ サービス github.com 通常の認証済みリクエストを受け取ります。 AGについては何も知らない

入口。 ID オーバースラッシュは、すべての発信通話に決定論的なアクターをスタンプします。ルール ルールはゲートウェイに存在します。プロンプトにはありません。コードにはありません。人間がワンクリックで編集可能。バブルアップ 不明なスコープはエージェント ツリーを上にエスカレートし、人間にまでエスカレートします。拒否は記録されます。監査 すべての通話、応答ステータス、バブル、承認が監査ログにストリーミングされます。統合 エージェントがすでに知っているサービス レジストリ。
エージェントが実際に触れるものの第一級のテンプレート。 URL を指定できるその他のものに対する汎用の http サービス。
自己ホスト型では無料。個人向けクラウド無料。チームに対して支払われます。
オーバースラッシュを使用する 3 つの方法。セルフホストではすべてが無料で提供されます。 Cloud 上の個人組織は永久に無料です。同僚を連れて行きたい場合は、1 シートあたり €3 の料金でチーム組織を作成します。
自分で実行してください。フル機能 - ゲート、ライセンス キー、テレメトリはありません。エラスティック ライセンス 2.0。
無制限のエージェント、アクション、統合
あなた専用のホスト型個人組織。いつでも無料、カードは必要ありません。
すべての統合を備えた個人組織
1 人のユーザー向けのすべてのクラウド機能
チーム組織を作成し、同僚を連れてきます。各席3ユーロ。あなた自身の個人組織は引き続き無料です。
プールされた使用量と従量制の超過料金
ゲートウェイ コアには Elastic 2.0 のライセンスが付与されています。サービス レジストリ (コミュニティが貢献する部分) は MIT です。テレメトリーはありません。電話はありません。 Rust と SvelteKit で書かれています。
# クローン、ビルド、実行 — ローカル開発
$ git clone https://github.com/overfolder/overslash
$ cd オーバースラッシュ
$メイクインストール
$オーバースラッシュウェブ
# http://localhost:7171 で準備完了
# 署名キーが生成されました · 監査データベースが初期化されました
# 登録リンクを MCP クライアントに貼り付けます。 インストール · 開始 これをエージェントに渡します。
以下のブロックを Claude、Cursor、Open Interpreter、または MCP 対応エージェントに貼り付けます。スキルに従います。enr

自分自身を Overslash アカウントの下で実行し、許可が表示されると許可を求めます。

## Original Extract

One gateway between your agent and every service it touches. OAuth, secrets, approvals, and audit — handled.

Overs / ash beta Features How it works Integrations Pricing Open source Install EN ▾ ☀ GitHub ↗ Open Cloud ↗ Your agent acts. Overs / ash controls.
One gateway to every service your agent touches. You decide: deny, approve once, or approve all.
One gateway between your agent and everything else.
Overslash sits between agents and the outside world. It handles secrets, OAuth, MCP, permission chains, human approvals, and authenticated HTTP execution. The agent doesn't hold keys. You don't hold your breath.
Every agent gets an identity, a parent, and a blast radius. Sub-agents inherit rules from their parent — until you narrow them.
The first time an agent wants a permission, you approve or deny. Allow & Remember at the scope ladder you choose.
One place for client IDs, tokens, signing keys. Rotate, revoke, per-agent. No credentials ever touch the agent's context window.
Every call, every bubble-up, every deny. Streaming, searchable, exportable.
Enroll any MCP client in a click — Claude Desktop, your own. Overslash brokers the tools, the credentials, and the approvals. The agent just asks.
A sub-agent bubbles up to its parent when it hits an unknown scope. Parents bubble to humans. Denials are first-class and recorded.
Agents ask. Overslash mediates. Services respond.
Any MCP-capable agent — Claude Code, Overfolder, OpenClaw, or your own — connects to Overslash once. Overslash holds the credentials, enforces the rules, and hands each service an authenticated, audited request.
agent:henry agent:deploy-bot agent:research-bot Gateway Overs / ash Checks rules. Injects credentials. Bubbles unknowns to parent agents or humans. Records every call. authed request ↓ Service github.com Receives a normal authenticated request. Knows nothing about the agent. Identity Overslash stamps every outbound call with a deterministic actor. Rules Rules live at the gateway. Not in prompts. Not in code. Editable by a human in one click. Bubble-up Unknown scopes escalate up the agent tree, then to a human. Denials are recorded. Audit Every call, response status, bubble, and approval is streamed to the Audit log. Integrations A service registry your agent already knows.
First-class templates for the things agents actually touch. And a generic http service for anything else you can point a URL at.
Free in self-hosted. Cloud free for individuals. Paid for teams.
Three ways to use Overslash. Self-hosted gives you everything for free. Your Personal org on Cloud is free, forever. When you want to bring colleagues, create a Team org at €3 per seat.
Run it yourself. Full features — no gating, no license keys, no telemetry. Elastic License 2.0.
Unlimited agents, actions, integrations
A hosted Personal org, just for you. Always free, no card required.
Personal org with all integrations
All Cloud features for one user
Create Team orgs and bring your colleagues. Every seat €3. Your own Personal org stays free.
Pooled usage with metered overages
The gateway core is licensed Elastic 2.0. The services registry — the part the community contributes to — is MIT. No telemetry. No phone-home. Written in Rust and SvelteKit.
# clone, build, run — local dev
$ git clone https://github.com/overfolder/overslash
$ cd overslash
$ make install
$ overslash web
# ready on http://localhost:7171
# signing key generated · audit db initialised
# paste the enrollment link into your MCP client Install · get started Give this to your agent.
Paste the block below into Claude, Cursor, Open Interpreter, or any MCP-capable agent. It will follow the skill, enroll itself under your Overslash account, and ask you for permissions as they come up.
