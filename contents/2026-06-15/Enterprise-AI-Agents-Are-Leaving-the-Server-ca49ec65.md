---
source: "https://focused.io/lab/enterprise-ai-agents-are-leaving-the-server"
hn_url: "https://news.ycombinator.com/item?id=48536050"
title: "Enterprise AI Agents Are Leaving the Server"
article_title: "Enterprise AI Agents Are Leaving the Server | Focused Labs | Focused"
author: "mooreds"
captured_at: "2026-06-15T04:37:18Z"
capture_tool: "hn-digest"
hn_id: 48536050
score: 1
comments: 0
posted_at: "2026-06-15T02:56:39Z"
tags:
  - hacker-news
  - translated
---

# Enterprise AI Agents Are Leaving the Server

- HN: [48536050](https://news.ycombinator.com/item?id=48536050)
- Source: [focused.io](https://focused.io/lab/enterprise-ai-agents-are-leaving-the-server)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T02:56:39Z

## Translation

タイトル: Enterprise AI エージェントがサーバーから離れる
記事のタイトル: Enterprise AI エージェントがサーバーを離れる |集中的なラボ |集中した
説明: Enterprise AI エージェントはクライアント ランタイムを横断するようになり、アプリの状態、権限、承認、フロントエンドの可観測性によって安全に動作できるかどうかが決まります。

記事本文:
Enterprise AI エージェントがサーバーから離れる |集中的なラボ |集中した
ほとんどの AI プロジェクトは失敗します。あなたのものはその必要はありません。今すぐスポットを予約して、わずか 3 週間で本番環境に対応したエージェント ブループリントを入手してください 6 スポット ‍
‍ 利用可能 エージェント ブループリントに登録します
機能について
カスタム エージェント 信頼性の高い RAG カスタム ソフトウェア開発 Eval 駆動開発 可観測性 LangChain ケーススタディ 集中ラボ お問い合わせ
ブログ Enterprise AI エージェントがサーバーを離れる
AI エージェントの統合は、クライアント ランタイムを越えて行われるようになりました。アプリの状態、承認、アクセス許可、およびフロントエンド トレースによって、エンタープライズ エージェントが安全に動作できるかどうかが決まります。
エンタープライズ AI エージェントはサーバーの境界を離れています。
ブラウザー タブ、デスクトップ アプリケーション、グリッド上の行、ローカルに保存された下書き、クリップボード、デバイスのアクセス許可、承認フロー、その他の混乱の中でエージェントがユーザーの代わりに行動を開始するまでは、一見小さく見える境界線。その人の作業が必ずしもサーバー側の記録に変換されるわけではないため、サーバー専用のエージェント ツールは主要な統合モデルとしては不十分です。
バックエンドツールは製品の瞬間を確認できません
サーバー ツールは、アカウントの更新、ナレッジ ベースの検索、チケットの作成、または ERP ワークフローの呼び出しを行うことができます。これは、製品の意図が保存された事実になった「後の記録」です。
製品の瞬間が早く到着します。
ユーザーは、ワークフロー内で提案された一連のアクションから 3 つの箇条書きを選択します。セールス エンジニアが製品セットの価格を編集しており、各製品の割引に保存されていない変更を加えました。サポート担当者は、一連の顧客に関するインシデントのタイムラインを表示しています。プロダクト マネージャーは、分析のために顧客のコホートを選択しました。クライアントは、カーソルがどこにあるか、ユーザーが何を選択したか、ページ内のスクロール位置を知っています。

roduct、ユーザーがいる現在のルート、ワークフローの現在のステップの未保存のフォーム データ、現在のビューポートのサイズ、現在のブラウザーの権限状態、およびユーザーが実行した最後の UI アクション。サーバーは何も知らないか、レコードのセットの古いオブジェクト モデルを知っています。
このギャップこそが、LangChain のヘッドレス ツール用アーキテクチャが非常に重要である理由です。モデルにとって、このツールは、名前、説明、パラメーターのスキーマ、および結果を備えた通常のツールにすぎません。この重要な点は、ツールがクライアント上で実行されることです。
これにより、企業内での統合の焦点も大きく変わります。 CRM 統合がエージェント ランタイムに移行することについて書いたように、統合が安全かどうかは、ID、承認、再試行、冪等性、およびトレースによって決まります。そして今週説明したように、同じモデルがブラウザーにも浸透しつつあります。バックエンド ランタイムは引き続き、バックエンド サービスのエンタープライズ統合を行う場所です。ただし、Figma で選択したオブジェクト、CRM モーダルの未保存のフィールド、またはさらに簡単に言えば、ブラウザの許可プロンプトはすべてエージェントの実行パス内にあります。
クライアント ランタイムは実行サーフェスの一部になります。
フロントエンド ツールは UI 接着剤ではなくコントラクトです
遅延アプローチはサイド チャネルです。アプリケーションの状態をシリアル化し、それを古いバイナリ BLOB としてサーバーに送信し、モデルに応答を生成させ、その結果から UI にパッチを適用するようにアプリに依頼します。確かに、それは初めてうまくいきます。その後、シリアル化されたデータの形状が、コードの作成者にとってさえ明らかではない方法で変化し、モデルは古いコンテキストから動作し始め、現在の UI がユーザーのアクションによるものなのか、ツールの実行によるものなのか、それともアプリ チームが従う間にモデルが盲目的に推測したのかは誰もわかりません。

それ。
フロントエンド ツールはコントラクトを明示的にします。 AG-UI は、ツールを、名前、説明、パラメーターの JSON スキーマとともに実行時にエージェントに渡されるフロントエンド定義関数として記述します。フロントエンドは、引数の検証、呼び出し完了後のツールの呼び出し、会話履歴へのツールの結果の挿入を実装します。単純。
重要な部分は、エージェントに渡される機能に対するフロントエンドの制御です。各ツールについて、フロントエンドは、ユーザー権限、アプリケーション コンテキスト、状態 (AG-UI ツール) に基づいて、ツールをランタイムに追加するか削除するかを決定できます。
たとえば、引用編集者は、引用の対象となるレコードが編集可能で、その句が承認されたライブラリから選択され、ユーザーが引用を変更する権限を持っている場合にのみ、insertApprovedClause を許可することを決定する場合があります。一方、サポート コンソールでは、draftCustomerReply を自由に許可できますが、sendCustomerReply の承認が必要な場合があります。設計ツールでは、承認なしで SummarySelectedFrame を許可しても、replaceSelectedFrameCopy の承認が必要な場合があります。
イベント ストリームはツール、状態、承認、サブエージェント、エラー、可観測性の型付きハンドルを製品に提供するため、エージェント UI はランタイム インフラストラクチャであると以前に議論しました。クライアント実行ツールでは、その議論が理論的ではなくなります。製品 UI は、もはやエージェントの単なるシェルではありません。これは、エージェントがサーバーから安全に偽装できない実行可能機能を所有しています。
AG-UI はスケジュールどおりに表示されるプロトコル層です
MCP はエージェントにツールとデータへの標準インターフェイスを提供し、A2A はエージェントが他のエージェントと対話するための標準インターフェイスを提供します。 AG-UI は、エージェントとユーザー側のアプリケーションのインターフェイスをターゲットとしています。この空間では、イベント (プログラムされたものまたは人間がトリガーしたもの) と更新のストリーミングが行われます。

o UI だけでなく、マルチモーダル入力 (音声やインクなど)、共有状態、フロントエンド ツールの呼び出し、および人間参加型の割り込みもすべて UI によって処理される必要があります。これは、現在 AG-UI によって定義されている機能の範囲です。
ユーザー向けアプリケーションが実行時の事実 (現在誰が存在しているか) を判断できる点には、システム内に明確な境界があります。ユーザーが選択したもの。ツールの結果に影響を与える、ユーザーのワークステーション上でローカルに変更された内容。ユーザーのワークステーションで何を元に戻すことができるか。そして、サーバー上で特定の一連の副作用が発生する前に、ユーザーのワークステーション上で人間によるクリックが必要なもの。ツールが製品内のパンフレットウェアの統合から製品内の運用アクションに移行すると、エージェントが操作可能なインターフェイスが製品になります。
Microsoft のエージェント フレームワーク AG-UI 統合は同じ方向を向いています。そのドキュメントには、リアルタイム ストリーミング、セッションとスレッドの管理、状態の同期と共有、人間参加型の承認ワークフロー、カスタム UI と生成 UI、ツールの実行、Web およびモバイル クライアント向けのツール結果のストリーミングがリストされています。
デモでは、「承認済み」などのテキストをパネルに送信し、承認されたテキストが適切な場所に表示されるかどうかをチェックするプログラムを利用できます。運用グレードのエンタープライズ AI エージェントは、要求されたクライアント アクション、ユーザーの承認、実行中のデータ、およびアクションが実際に別の場所に送信されたかどうかを考慮する必要があります。
ビジュアルビルダーはこの境界を所有しません
OpenAI の AgentKit ページには、Agent Builder と Evals が 2026 年 11 月 30 日以降に終了すると記載されています ( OpenAI AgentKit update )。同じ更新により、チームはコードとして継続するワークフロー用のエージェント SDK と自然なワークスペース エージェントを参照できるようになります。

l 言語のプロンプト。ビジュアルビルダーは引き続き意図をスケッチできます。永続的なエージェントの統合は、アプリケーション所有のコードに戻り続けます。
キャンバスはワークフローをスケッチできます。アクティブなブラウザーの選択が依然としてツール呼び出しの引数と一致するかどうかを確認できません。アプリケーションがローカル許可ルールを与えない限り、ローカル許可ルールを所有することはできません。承認プロンプトがこれから発生する副作用を反映していることを証明することはできません。エンタープライズ AI エージェントの場合、永続的な境界はアプリケーション アーキテクチャです。つまり、型指定されたツール、スコープ指定された資格情報、状態の同期、レビュー可能な副作用、およびアクションに続くトレースです。
これが、AI エージェントのガバナンスが実行パスに従う理由です。 AI エージェントのガバナンスは、LangGraph、AG-UI、ヘッドレス ツール、SDK などのツールを使用して、AI エージェントの制御下で実行されるアプリケーションの実行パスに従います。これはサーバー パスに従わないため、サーバー側アプリケーションのガバナンスとは明らかに異なります。前と同様、AI エージェントのガバナンスを成功させる鍵は、どのアプリケーションでも同じです。アプリケーションとその AI は、AI の機能を定義し、アプリケーションによって操作される AI の実行時の事実をレビューできる製品チームが所有する必要があります。
クライアントのアクションは観察可能でなければなりません
ブラウザがエージェントの計画の一部を実行している場合、バックエンドのみのトレースは機能しません。これは、エージェントがクライアント ツールにコマンドを送信できることを意味します。その後、クライアント ツールはローカル状態を検証できます。ユーザーはアクションを承認できます。その後、ブラウザは外部 API を呼び出すことができます。バックエンドはアクションの結果を保存できます。これらのスパンが接続されたトレースを形成しない場合、インシデントのレビューはスクリーンショットに変わり、Slack メッセージは時系列の逆順に一度に 1 つずつ読み取られます。
Honeycomb ブログは最近、usi に関する記事を公開しました。

ブラウザーの OpenTelemetry (ブラウザーの OpenTelemetry 上の Honeycomb)。著者が指摘しているように、コードは予期せぬ環境（つまり、同時に予測できないユーザー入力の下）で実行されるため、フロントエンド コードのインストルメントは難しく厄介な問題です。この投稿では、ブラウザー インストルメンテーションがトレース コンテキストを後続のバックエンド リクエストにどのように伝播するかについて説明し、同じセッション内の異なるユーザーのフロントエンド コードによって生成されたトレースを相互に関連付ける方法としてセッション ID の使用について説明します。
Honeycomb のフロントエンド可観測性 GA ポストは、エンドツーエンドのユーザー フロー、高カーディナリティ データ、ユーザー インタラクション コンテキスト、カスタム属性、およびユーザーに影響を与える特定の動作のデバッグをプッシュします。エージェントをフロントエンドに追加すると、トレースには、クライアントで実行されたすべてのアクションのエージェント識別子、ツール呼び出し ID、承認決定、許可結果、状態バージョン、および受信 ID が含まれる必要があります。
クライアント上で実行されているツールから得られる良好な結果は、単なる「ok: true」以上のものです。これには、実行されたコマンド、ツールが読み取った状態、開かれた権限、アクションの承認者、行われた変更、元に戻せるアクション、およびトレース ID に関する情報が含まれる必要があります。
エージェントが実行する前にクライアント ランタイムを所有する
製造チェックリストは簡単です。
クライアント ツールをコードとして定義します。これは、コンポーネント内に埋め込まれたコールバック スタイルの関数ではなく、型付きのコントラクトを意味します。システム プロンプトのヒューリスティックではなく、ツールの権限ルールを使用します。クライアントが古いリクエストを拒否できるように、各ツール呼び出しに最新の状態バージョンを含めます。正確な副作用の説明とともに、製品ワークフローを通じて承認をルーティングします。クライアントが実行したすべてのアクションのレシートを記録します。ブラウザー、エージェント ランタイム、バックグラウンドにわたる実行パスをたどります。

kend サービス、および外部 API。ローカルまたはリモートの状態を変更するアクションの元に戻すパスを構築します。誰かがインターフェースを所有する必要があります。
エンタープライズ AI エージェントは、サーバー上でのみ作業が行われることはなかったので、サーバーから離れています。この作業は、アプリケーションの状態、ユーザーの意図、承認、副作用が交わる厄介な中間にあります。これが AI エージェントの統合の現状です。
すぐにご連絡いたします。それまでの間、当社のケーススタディをご覧ください。
/お問い合わせ より良いエージェントを一緒に構築しましょう
Focused でレガシーを最新化する
スイート 1100-C
イリノイ州シカゴ 60607
‍
‍ work@focused.io
‍ (708) 303-8088

## Original Extract

Enterprise AI agents now cross the client runtime, where app state, permissions, approvals, and frontend observability decide whether they can act safely.

Enterprise AI Agents Are Leaving the Server | Focused Labs | Focused
Most AI projects fail. Yours doesn’t have to. Reserve your spot today and get a production-ready Agent Blueprint in just 3 weeks 6 spots ‍
‍ available Register for Your Agent Blueprint
About Capabilities
Custom Agents Reliable RAG Custom Software Development Eval Driven Development Observability LangChain Case Studies Focused Lab Contact us
Blog Enterprise AI Agents Are Leaving the Server
AI agent integration now crosses the client runtime. App state, approvals, permissions, and frontend traces decide whether enterprise agents can act safely.
Enterprise AI agents are leaving the server boundary.
A boundary that looks deceptively small until the agent starts acting on behalf of a person inside a browser tab, a desktop application, a row on a grid, a locally saved draft, a clipboard, a device permission, an approval flow, and the rest of the mess. That person’s work does not always translate into a server-side record, so server-only agent tools are insufficient as the primary integration model.
Backend tools cannot see the product moment
A server tool can update an account, search a knowledge base, create a ticket, or call an ERP workflow. This is the “record after” the product has turned intent into a stored fact.
The product moment arrives earlier.
A user selects three bullets from a proposed set of actions in a workflow. A sales engineer is editing the pricing for a set of products and has made unsaved changes to the discount for each. A support rep is viewing an incident timeline of incidents for a set of customers. A product manager has selected a cohort of customers for analysis. The client knows where the cursor is, what the user has selected, the scroll position in the product, the current route the user is on, the unsaved form data for the current step in the workflow, the dimensions of the current viewport, the current browser permission state, and the last UI action that the user performed. The server knows nothing, or it knows a stale object model for a set of records.
That gap is why LangChain's architecture for headless tools is so important. To the model, the tool is just another normal tool with a name, description, schema for the parameters, and result. The significant aspect of this is that the tool is being executed on the client.
This also shifts the focus of integration in the enterprise significantly. As we wrote about CRM integration moving into the agent runtime , identity, approval, retry, idempotency and tracing decide whether the integration is safe. And as we laid out this week, that same model is now crossing over into the browser as well. The backend runtime is still the place to put enterprise integration for that backend service. But the selected object in Figma, unsaved field in a CRM modal, or even more simply, the browser permission prompt are all now in the agent’s execution path.
The client runtime becomes part of the execution surface.
Frontend tools are contracts, not UI glue
The lazy approach is the side channel: serialize application state, send that off to the server as a big ol’ binary blob, let the model generate a reply, then ask the app to patch the UI from the result. Sure, that works the first time. Then the shape of the serialized data changes in a way that is not obvious even to the author of the code, the model starts operating off stale context, and nobody knows whether the current UI came from a user action, a tool execution, or the model making a blind guess while the app team followed it.
Frontend tools make the contract explicit. AG-UI describes tools as frontend-defined functions passed to the agent at runtime with a name, description and a JSON Schema for the parameters. The frontend implements the argument validation, invocation of the tool after the call has completed, and insertion of the tool’s result into the conversation history. Simple.
The important part is the control the frontend has over the capabilities passed to the agent. For each tool, the frontend can decide whether it should be added or removed from the runtime based on user permissions, application context, and state ( AG-UI tools ).
A quote editor for example might decide to allow insertApprovedClause only when the record the quote is for is editable, the clause was chosen from the approved library and the user has permission to change quotes. A support console on the other hand might allow draftCustomerReply freely but require sendCustomerReply to be approved. A design tool might allow summarizeSelectedFrame without approval but require replaceSelectedFrameCopy to be approved.
We argued earlier that agent UI is runtime infrastructure because event streams give products typed handles for tools, state, approvals, subagents, errors, and observability. Client-executed tools make that argument less theoretical. A product UI is no longer merely a shell around an agent. It owns executable capabilities the agent cannot safely fake from the server.
AG-UI is the protocol layer showing up on schedule
MCP provides a standard interface to Tools and Data for Agents, A2A provides a standard interface for Agents to interact with other Agents. AG-UI is targeting the Agent-to-user-facing-application interface. In this space, events (programmed or human triggered) and the streaming of updates to the UI, as well as, multi-modal input (e.g., speech and ink), shared state, frontend tool calls, and human-in-the-loop interrupts all need to be dealt with by the UI. This is the scope of the functionality currently defined by AG-UI .
There’s a clear boundary in the system at the point where the user-facing application can determine the facts of runtime: who is currently present; what has the user selected; what has changed locally on the user’s workstation that will affect the tool results; what can be undone on the user’s workstation; and what, on the user’s workstation, requires a human click before a particular set of side effects can occur on the server. The agent-operable interface is the product once the tool moves from brochureware integration within the product to production action within the product.
Microsoft's Agent Framework AG-UI integration points in the same direction. Its documentation lists real-time streaming, session and thread management, state synchronization and sharing, human-in-the-loop approval workflows, custom and generative UIs, tool execution, and tool-result streaming for web and mobile clients.
Demos can rely on a program that sends out text, for example “Approved,” to a panel and then checks whether the approved text shows up in the right place. Production-grade enterprise AI agents have to account for the client action requested, the user's approval, the data under execution, and whether the action was actually sent somewhere else.
Visual builders will not own this boundary
OpenAI's AgentKit page now says that Agent Builder and Evals will wind down after November 30, 2026 ( OpenAI AgentKit update ). The same update points teams toward the Agents SDK for workflows that should continue as code and Workspace Agents for natural-language prompting. Visual builders can still sketch intent. Durable agent integration keeps returning to application-owned code.
A canvas can sketch a workflow. It cannot check whether the active browser selection still matches the tool call arguments. It cannot own a local permission rule unless the application gives it one. It cannot prove that an approval prompt reflected the side effect about to occur. For enterprise AI agents, the durable boundary is application architecture: typed tools, scoped credentials, state synchronization, reviewable side effects, and traces that follow the action.
This is why AI agent governance follows the execution path . Governance for AI agents, using tools such as LangGraph, AG-UI, headless tools and SDKs, follows the execution path of the application running under the control of the AI agent. It does not follow the server path, and thus is distinctly different from governance of server-side applications. As before, the key to successful governance of AI agents, is the same as for any application: the application and its AI, must be owned by a product team, who can define the capabilities of the AI, and review the runtime facts of the AI operated by the application.
Client actions have to be observable
Backend-only traces don’t work when the browser is executing part of the agent’s plan. That means the agent can send a command to a client tool. The client tool can then validate local state. The user can approve the action. The browser can then call an external API. And the backend can store the result of the action. If these spans do not form a connected trace, then incident review turns into screenshots and Slack messages read one at a time in reverse chronological order.
The Honeycomb blog recently published a write-up on using OpenTelemetry in the browser ( Honeycomb on OpenTelemetry in the browser ). As the author points out, instrumenting frontend code is a difficult, messy problem because the code runs in surprise environments (i.e. under simultaneous and unpredictable user input). The post describes how browser instrumentation can propagate trace context to subsequent backend requests, and discusses the use of session IDs as a way to correlate together traces generated by the frontend code of different users within the same session.
Honeycomb’s frontend observability GA post pushes end-to-end user flows, high-cardinality data, user interaction context, custom attributes, and debugging specific user-impacting behavior. Add an agent to the frontend and the trace has to carry agent identifiers, tool-call IDs, approval decisions, permission outcomes, state versions, and receipt IDs for every action executed on the client.
A good result from a tool running on the client is more than just “ok: true”. It needs to include information about the command that was executed, the state that the tool read, the permission that was opened, who approved the action, the change that was made, the actions that can be undone, and the trace id.
Own the client runtime before the agent does
The production checklist is straightforward.
Define client tools as code, which means typed contracts, not callback-style functions buried inside a component. Use the permission rules of the tool rather than heuristics in system prompts. Include the latest state version in each tool call so the client can reject stale requests. Route approvals through the product workflow with exact side-effect descriptions. Record a receipt for every client-executed action. Follow the execution path across browser, agent runtime, backend service, and external API. Build undo paths for actions that modify local or remote state. Someone has to own the interface.
Enterprise AI agents are leaving the server because the work was never only on the server. The work is in the messy middle where application state, user intent, approval, and side effects meet. This is where AI agent integration lives now.
We’ll be in touch soon. In the mean time check out our case studies.
/Contact Us Let's Build better Agents Together
Modernize your legacy with Focused
Suite 1100-C
Chicago, IL 60607
‍
‍ work@focused.io
‍ (708) 303-8088
