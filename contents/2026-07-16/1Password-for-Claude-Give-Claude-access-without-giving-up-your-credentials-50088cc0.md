---
source: "https://1password.com/blog/1password-for-claude"
hn_url: "https://news.ycombinator.com/item?id=48936522"
title: "1Password for Claude: Give Claude access without giving up your credentials"
article_title: "1Password for Claude: Give Claude access without giving up your credentials | 1Password"
author: "terracatta"
captured_at: "2026-07-16T17:05:03Z"
capture_tool: "hn-digest"
hn_id: 48936522
score: 6
comments: 3
posted_at: "2026-07-16T16:13:20Z"
tags:
  - hacker-news
  - translated
---

# 1Password for Claude: Give Claude access without giving up your credentials

- HN: [48936522](https://news.ycombinator.com/item?id=48936522)
- Source: [1password.com](https://1password.com/blog/1password-for-claude)
- Score: 6
- Comments: 3
- Posted: 2026-07-16T16:13:20Z

## Translation

タイトル: クロードの 1Password: 資格情報を放棄せずにクロードにアクセス権を与える
記事のタイトル: クロードの 1Password: 資格情報を放棄せずにクロードにアクセスを許可する | 1パスワード
説明: 秘密の制御を放棄することなく、クロードにアクセス権を与えます。 1Password により、AI エージェントは認証情報を暗号化して安全に保ち、AI エージェントがアクセスできない状態に保ちながら、認証情報によるアクセスが必要なタスクを完了できるようになりました。

記事本文:
1Password for Claude: 資格情報を放棄せずにクロードにアクセスを許可します | 1Password メインコンテンツにスキップ 人間、マシン、AI エージェントの認証情報を保護します。 1Password に切り替えます。ブログに戻る クロード用 1Password: 資格情報を放棄せずにクロードにアクセスを許可します
ミッチェル・コーエンとホリア・クレア著
AI エージェントは、人々の思考を支援することから、ブラウザ、アプリ、アカウントで人々に代わって行動することに移行しつつあります。それによってセキュリティ モデルが変わります。エージェントがクリック、購入、更新、送信できるようになると、重要な問題は次のとおりです。エージェントはどのような ID に基づいて動作するのか、どのようなアクセス権を取得する必要があるのか​​、ということです。
クロードは、取引を比較したり、カートに商品を追加したり、アカウントの詳細を更新したり、購入を完了したりできます。しかし、ログイン ページに到達すると、トレードオフに直面することになります。エージェントにパスワードを渡しますか、それともやめて自分でタスクを実行しますか?私たちが築くべき未来も同様です。
これまで、エージェントが認証情報を公開せずに使用するための安全かつ簡単な方法はありませんでした。
Claude の 1Password: 資格情報を公開せずに資格情報にアクセス
1Password for Claude は、ゼロ暴露アーキテクチャに基づいて構築されています。Claude は、ログインとワンタイム パスコードを必要とするブラウザ タスクを完了できますが、認証情報がモデルやそのメモリに入力されることはありません。 1Password はシークレットの信頼できる情報源であり、アクセスは実行時にのみ許可されます。
クロードがサインインする必要がある場合、1Password はユーザーにどの資格情報が要求されているか、およびその理由を示します。ユーザーの同意による生体認証の承認後、1Password は認証情報をページに直接挿入します。クロードは、保管庫アイテム、パスワード、またはワンタイム コードを決して見ることはありません。アクセスの範囲は現在のタスクに限定され、タスクが完了すると終了します。自動入力後、1Password はページ上でシークレットが公開されていないことをチェックします。送信が失敗すると、入力された値がクリアされます

制御を戻す前に。
1Password の CTO であるナンシー ワン氏は、「人間だけでなくエージェント専用に構築された新しいセキュリティ モデルが必要です」と述べています。 「その答えは、エージェントにあなたの機密情報を渡すことではありません。それは、エージェントに認証情報を見せずに、ユーザーがエージェントに資格情報の使用許可を与えることです。クロードは、それがあなたのログインを使用したことを知っています。そのコンテキストではパスワードやワンタイム コードは必要ありません。この違いが、エージェントへの信頼の始まりであり、私たちが Anthropic で構築している基盤です。」
実際にはどうなるか
Audible クレジットの有効期限が近づいています。ログインしてストアに移動し、手動でクレジットを引き換える代わりに、クロードにウィッシュリストを確認して新しいタイトルを選択するように依頼します。クロードがサイトに移動し、クロードが保管庫の資格情報を使用することを承認すると、1Password がログインを提供し、オーディオブックがライブラリに追加されます。あなたはパスワードや TOTP を入力したことはありませんし、クロードもそのどちらも見ることはありません。
AI を推進する中小企業経営者向け
中小企業の経営者は、Claude に Stripe の収益概要を尋ねたり、異常なアクティビティにフラグを立てたりすることができます。クロードはダッシュボードを操作でき、ビジネス所有者はクロードが自分の Stripe ログイン詳細を使用することを承認し、1Password が資格情報とワンタイム コードを処理でき、ユーザーは MFA を経由したりシークレットを公開したりすることなく答えを取得できます。
これらは 2 つの例にすぎません。同じパターンが、Chrome のクロードがアクションを実行できるサイト全体で機能します。ログインが 1Password である場合、クロードはそれを使用できます。あなたが承認すると、1Password が認証情報を提供し、Claude がジョブを完了します。タスクが変わっても、アクセス モデルは同じままで、資格情報が 1Password から離れることはありません。
エージェント モード: エージェントがブラウザを制御するときにボールトを保護する
2 つ目の問題は、ブラウザベースのエージェントが動作するとどうなるかということです。

1Password がインストールされているブラウザを制御しますか?適切なガードレールがないと、エージェントが拡張機能自体と対話しようとする可能性があります。エージェントティック モードは、そのギャップを埋める方法です。
エージェントティック モードは、1Password ブラウザ拡張機能の新機能であり、すべてのユーザーにブラウザベースの AI エージェントの可視性と制御を提供します。互換性のある AI エージェントが引き継ぐと、1Password 拡張機能は自動的にロックダウンされます。インターフェイスは非表示になっており、エージェントは現在のタスクに対して明示的に承認されたログインとワンタイム コードのみを使用できます。金庫室の残りの部分は手の届かないところにあります。
エージェント モードは、統合が設定されていない場合や、現在のエージェント タスクに 1Password が必要でない場合でも機能します。クロード以外の追加エージェントもサポートします。対象となる企業の場合、新たに構成する必要はありません。仕事用資格情報として 1Password を使用している従業員は、自動的に同じ保護を受けます。AI エージェントからのすべての資格情報リクエストは表示され、明示的であり、承認が必要です。
AI エージェント用の信頼されたアクセス層
1Password for Claude は、OpenAI Codex を使用した 1Password 環境 MCP サーバーによる開発者の資格情報の保護から、Kiro 用の 1Password MCP サーバーまで、エコシステム全体の AI エージェント向けに構築しているアクセス レイヤーの一部にすぎません。エージェントがブラウザ、IDE、リポジトリ、ターミナル、または CI/CD ワークフローで動作しているかどうかに関係なく、原則は同じです。シークレットは実行時に発行され、タスクにスコープされ、 1Password によって管理される必要があります。
エージェントがより有能になるにつれて、エージェントは新しいクラスのアイデンティティになります。人間や機械と同じように、管理されたアクセスが必要です。 1Password for Claude は、そのモデルをブラウザベースの委任に適用します。クロードは、明示的なユーザー認証を使用して動作し、必要なときに必要なアクセスのみを取得できます。資格情報は暗号化されたままになります

pted、制御され、モデルのコンテキストの外にあります。
あなたの秘密を渡さずに、クロードにもっと任せる準備はできていますか?
1Password for Claude は、ビジネス、ファミリー、個人の各プランで Mac で利用できるようになりました。設定方法の詳細については、ドキュメントを参照してください。この統合を有効にするには、次のものが必要です。
1Password ブラウザ拡張機能
Chrome ブラウザ拡張機能のクロード
これまで自分の速度を低下させていたタスクをクロードに任せられるまであと少しです。
1Password マーケットプレイスにアクセスします
1Password を使い始めて、初日から Claude 統合を準備しましょう。
1Password は OpenAI の Codex の信頼できるアクセス層になりました
1Password + Kiro: AI を活用した開発のための信頼できるアクセス
サイバー保険とコンプライアンス
利用規約 Cookie ポリシー プライバシーの選択 カリフォルニア州消費者プライバシー法 (CCPA) オプトアウト アイコン プライバシー ポリシー アクセシビリティ サイトマップ © 2026 1Password.無断転載を禁じます。 4711 Yonge St、10階、トロント、オンタリオ、M2N 6K8、カナダ

## Original Extract

Give Claude access, without giving up control of your secrets. 1Password now enables AI agents to complete tasks requiring credentialed access, while keeping your credentials encrypted, secure, and inaccessible to AI agents.

1Password for Claude: Give Claude access without giving up your credentials | 1Password Skip to Main Content Secure credentials for humans, machines, and AI agents. Switch to 1Password. Back to blog 1Password for Claude: Give Claude access without giving up your credentials
by Mitchell Cohen and Horia Culea
AI agents are moving from helping people think to acting on their behalf in browsers, apps, and accounts. That changes the security model. Once an agent can click, buy, update, and submit for you, the key question becomes: what identity is it acting under, and what access should it get?
Claude can compare deals, add an item to your cart, update account details, or complete a purchase. But once it reaches a login page, you face a tradeoff. Do you give the agent your password, or stop and do the task yourself? Neither is the future we should build toward.
Until now, there hasn’t been a secure, easy way for agents to use credentials without exposing them.
1Password for Claude: Credential access without credential exposure
1Password for Claude is built on a zero-exposure architecture: Claude can complete browser tasks that require logins and one-time passcodes, but the credentials never enter the model or its memory. 1Password stays the source of truth for the secret, and access is granted only at runtime.
When Claude needs to sign in, 1Password shows the user which credential is being requested and why. After user-consented biometric approval, 1Password injects the credential directly into the page. Claude never sees the vault item, password, or one-time code. Access is scoped to the current task and ends when the task is complete. After autofill, 1Password checks that secrets were not exposed on the page. If submission fails, it clears the filled values before returning control.
"We need a new security model that is purpose-built for agents, not just humans,” said Nancy Wang, CTO of 1Password. “The answer isn't handing agents your secrets. It is to let a user give an agent permission to use a credential without letting the agent see it. Claude knows it used your login; it does not need the password or one-time code in its context. That distinction is where trust in agents starts and the foundation we're building with Anthropic."
What this looks like in practice
Your Audible credits are about to expire. Instead of logging in, navigating the store, and manually redeeming a credit, you ask Claude to review your wishlist and choose a new title. Claude navigates to the site, you provide approval for Claude to use the credential from your vault, 1Password provides the login, and the audiobook lands in your library. You never typed a password or TOTP, and Claude never sees either.
For AI-forward small business owners
A small business owner could ask Claude for a Stripe revenue summary or to flag any unusual activity. Claude can navigate the dashboard, the business owner approves Claude to use their Stripe login details, 1Password can handle the credential and one-time code, and the user gets the answer without going through MFA or exposing the secret.
These are just two examples. The same pattern works across the sites where Claude in Chrome can take action: if the login is in 1Password, Claude can use it. You approve, 1Password supplies the credential, and Claude finishes the job. Even when the task changes, the access model stays the same, and your credentials never leave 1Password.
Agentic Mode: Protecting the vault when an agent controls the browser
There’s a second problem: what happens when a browser-based agent takes control of a browser where 1Password is installed? Without proper guardrails, the agent could try to interact with the extension itself. Agentic Mode is how we close that gap.
Agentic Mode is a new feature in the 1Password browser extension that gives every user visibility and control over browser-based AI agents. When a compatible AI agent takes over, the 1Password extension automatically locks down. The interface is hidden, and the agent can only use the logins and one-time codes explicitly approved for the current task. The rest of the vault stays out of reach.
Agentic Mode works even if the integration is not set up and even if 1Password is not required for the current agentic task. It also supports additional agents beyond Claude. For qualifying enterprises, there is nothing new to configure. Employees using 1Password for work credentials automatically get the same protection: every credential request from an AI agent is visible, explicit, and requires authorization.
The trusted access layer for AI agents
1Password for Claude is just one part of the access layer we’re building for AI agents across the ecosystem: from securing developer credentials with the 1Password Environments MCP Server with OpenAI Codex to the 1Password MCP Server for Kiro . Whether the agent is working in a browser, IDE, repo, terminal, or CI/CD workflow, the principle is the same: secrets should be issued at runtime, scoped to the task, and governed from 1Password .
As agents become more capable, they become a new class of identity. They need governed access just like humans and machines do. 1Password for Claude applies that model to browser-based delegation: Claude can act with explicit user authorization and only gets the access it needs, when it needs it. The credential stays encrypted, controlled, and out of the model context.
Ready to let Claude handle more, without handing over your secrets?
1Password for Claude is available now for Mac, across business, family, and individual plans. For detailed instructions on how to set it up, read our documentation . To enable this integration, you'll need:
The 1Password browser extension
Claude in Chrome browser extension
You're a few steps away from letting Claude handle the tasks that used to slow you down.
Go to the 1Password Marketplace
Get started with 1Password and have the Claude integration ready from day one.
1Password is now a trusted access layer for OpenAI’s Codex
1Password + Kiro: Trusted Access for AI-Powered Development
Cyber Insurance and Compliance
Terms of Service Cookie Policy Your Privacy Choices California Consumer Privacy Act (CCPA) Opt-Out Icon Privacy Policy Accessibility Sitemap © 2026 1Password. All rights reserved. 4711 Yonge St, 10th Floor, Toronto Ontario, M2N 6K8, Canada
