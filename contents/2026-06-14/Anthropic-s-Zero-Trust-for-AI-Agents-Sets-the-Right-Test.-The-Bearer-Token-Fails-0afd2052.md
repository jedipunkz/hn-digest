---
source: "https://blog.hello.coop/2026/06/anthropics-zero-trust-for-ai-agents-sets-the-right-test-the-bearer-token-fails-it/"
hn_url: "https://news.ycombinator.com/item?id=48522754"
title: "Anthropic's Zero Trust for AI Agents Sets the Right Test. The Bearer Token Fails"
article_title: "Anthropic’s Zero Trust for AI Agents Sets the Right Test. The Bearer Token Fails It. – Hellō Blog"
author: "mooreds"
captured_at: "2026-06-14T01:02:04Z"
capture_tool: "hn-digest"
hn_id: 48522754
score: 3
comments: 0
posted_at: "2026-06-14T00:03:33Z"
tags:
  - hacker-news
  - translated
---

# Anthropic's Zero Trust for AI Agents Sets the Right Test. The Bearer Token Fails

- HN: [48522754](https://news.ycombinator.com/item?id=48522754)
- Source: [blog.hello.coop](https://blog.hello.coop/2026/06/anthropics-zero-trust-for-ai-agents-sets-the-right-test-the-bearer-token-fails-it/)
- Score: 3
- Comments: 0
- Posted: 2026-06-14T00:03:33Z

## Translation

タイトル: Anthropic の AI エージェントに対するゼロトラストは正しいテストを設定します。 Bearer トークンが失敗する
記事のタイトル: Anthropic の AI エージェントに対するゼロトラストは正しいテストを設定します。 Bearer トークンは失敗します。 – こんにちはブログ

記事本文:
Anthropic の AI エージェント向けゼロトラストは適切なテストを設定します。 Bearer トークンは失敗します。
フレームワークが停止する場所と、エージェント ネイティブの認証基盤が開始される場所。
Anthropic の AI エージェント向けゼロ トラストは、詳細かつよく構築されたエンタープライズ セキュリティ フレームワークであり、全文を読む価値があります。エージェントはチャットボットではないという出発点を正しく理解しています。彼らは目標を解釈し、ツールを選択し、セッション全体でコンテキストを保持し、他のエージェントと調整して行動します。そして、その結果は明確に述べられています。従来のアクセス制御では、エージェントが正当に保持している権限を悪用するのを阻止することはできません。それが正しいフレームです。
私はこの枠組みを真剣に受け止め、推進していきたいと考えています。それが間違っているからではなく、独自のベースライン推奨を廃止する正確なテストが含まれているからです。
文書自体が設定するテスト
この文書の中で最も強力なアイデアは、設計テストです。ガイドには、コントロールを評価するとき、次の 1 つの質問をするよう記載されています。
これでは攻撃が不可能になるのでしょうか、それとも単に退屈なのでしょうか?
Anthropic は、追加のホップ、レート制限、非標準ポート、SMS コードなどの摩擦を制御し、無制限の忍耐力とほぼゼロの試行あたりのコストで敵に対して機能を低下させることを明確に示しています。生き残るコントロールには、暗号化された ID、漏洩できない資格情報、単に不便なだけのパスではなく存在しないネットワーク パスなどのパターンがあります。疑わしい場合は、機能を抑制するコントロールよりも、機能を削除するコントロールを好むと彼らは言います。
それが正しいテストです。正直に適用すれば、文書よりもさらに効果が伝わります。
認証情報: 有効期間の短いトークンは面倒ですが、不可能ではありません
認証情報のベースラインは、ID プロバイダーからの有効期間が短く、範囲が狭いトークンであり、

相互 TLS に接続し、次に最上位層のハードウェアにバインドされた資格情報に接続します。この文書では、モデル支援型攻撃者がコード内で最初に発見するもののうち、静的 API キーと共有サービス アカウントのパスワードはすでに侵害されているものとして扱うべきであると率直に述べています。
ただし、独自のテストを通じてベースラインを実行します。有効期間の短いベアラー トークンはスロットルです。盗まれた秘密が役立つ期間が短縮されます。秘密は削除されません。トークンは依然として無記名資格情報であり、それを保持している人は誰でもそれを使用できます。ドアを閉めないと寿命が短くなりコストが上がります。ドキュメント独自の基準からすれば、それは退屈ではありますが、不可能ではありません。
資格情報の盗難を退屈ではなく不可能にする制御は、盗む秘密鍵が存在しない制御です。つまり、エージェントから離れることのない秘密鍵に関連付けられた所有証明です。エージェントはリクエストに署名します。キーはエクスポートできません。資格情報は決して文字列ではないため、侵害されたホストは再利用可能な資格情報を生成しません。それは署名を作成する能力でした。 HTTP メッセージ署名はまさにこれであり、AAuth が構築される基盤です。
何が欠けているかに注目する価値があります。この文書では、DPoP、OAuth トークン交換、リッチ認証リクエスト、またはあらゆる種類の送信者制限付きトークンについては言及されていません。これは、Advanced 層でのみ抽出不可能な資格情報に到達しており、それを、ベアラー トークン自体がそれ以下のすべての層に残存する欠陥であるという認識としてではなく、高保証ショップ向けの PKI 強化として枠付けしています。彼らの原則は処方箋を超えたものを指します。 「不可能だが退屈ではない」という正直な最終目標は、持ち去られる可能性のある秘密の発行をやめるということです。
承認: 「エージェントはアクセス権を持っている」では詳細が不十分です
ここにはより大きなギャップがあり、資格情報がどのように保持されるかが問題ではありません。それは約です

何が認可されるのか。
このフレームワークでは、承認を制御ファミリーのはしごとして説明します。つまり、ロールベースのアクセス制御、属性ベースのアクセス制御、継続的承認の順に、静的ロール、動的昇格、ジャストインタイム許可のライフサイクルはしごが組み合わされています。これは、アクセス制御カテゴリの合理的な分類です。これは、特定の呼び出しを許可するメカニズムではありません。
このモデルでは、認可の単位はエージェントとツールです。電子メール エージェントは電子メールのアクセス許可を取得します。データベース ツールは読み取りはできますが、書き込みはできません。属性ベースの制御により、時刻やリスク スコアなどのコンテキストが追加されます。しかし、許可は依然として「このエージェントはこのツールを使用できる」というものであり、おそらく大まかな属性によって絞り込まれています。呼び出しの実際のパラメータは入力検証として個別に処理され、実行前に不正な形式または疑わしい引数を拒否します。検証は、不正な入力に対するフィルターです。これは、特定のパラメーターが承認されるものであるという承認の決定ではありません。
この区別がエージェントにとって重要な点です。 「エージェントが支払いツールにアクセスできる」ということは有用な許可ではありません。 「このエージェントは金曜日までに 1 回、この受取人に最大 3,000 ドルを移動できます」です。パラメータはサニタイズされるノイズではありません。これらは承認コンテキストであり、呼び出しが行われた時点でリクエストごとに承認コンテキストとして評価される必要があります。
これが AAuth の R3 が提供するものです。パラメータが何であるかに固有の制約された呼び出しであり、それらのパラメータは評価されるコンテキストです。承認は、エージェントとツールの間の永続的な関係ではなく、実際の操作に関連付けられます。ドキュメントは正しい本能に従ってジェスチャーを行います。上級レベルでは次のことが求められます。
セッションの開始ではなく、アクションごとに認可を評価します。
しかし、それはこれをリストします

は、意欲的なトップレベルの機能として評価されており、評価がどのように機能するか、評価が何に結びつくかについては決して述べられていません。 How が重要な部分です。
このフレームワークに反応している他の企業も、同様の粒度の点を指摘し、それを継続的な認証として枠組み化し、汎用ゲートウェイでは「フライトを検索し、限度額まで予約するが、支払い方法には決して触れない」といった制約を表現できないと指摘している。診断は正しいです。メカニズムは着地すべきところです。つまり、ゲートウェイにボルトで固定されたポリシー表現ではなく、パラメータを評価されたコンテキストとして、制約された呼び出し自体が承認されたアーティファクトとして機能します。
委任: 正確に診断され、その後過少処方される
脅威セクションでは、委任について鋭く考察しています。これは、マネージャー エージェントがフル アクセス コンテキストを、権限が少ないワーカーに渡す、スコープ外の権限継承、および、権限の低いエージェントが、もっともらしい指示を権限の高いエージェントに中継し、元の意図を確認せずにその命令を実行する、混乱した代理の増幅と名付けています。それがマルチエージェントリスクの本当の姿です。
ただし、規定された救済策は検出です。つまり、信頼境界を確立し、各ホップで ID を検証し、エージェント間の通信をログに記録して、異常な委任パターンにフラグを立てて確認する必要があります。ロギングは委任モデルではありません。後で権限が漏れたことがわかります。権限の漏洩を止めることはできません。
この文書は、これが実際にどのような問題を引き起こすのかについて率直に述べています。独自のサブエージェントは親の完全な権限を継承でき、その区別はスコープ内で強制されるのではなく、テレメトリで取得されます。したがって、以前にフラグが立てられた障害モードは設計によって存在し、可視化によって軽減されます。
スロットルではなく削除という答えは派生した権威です。あるエージェントが別のエージェントに委任すると、サブエージェントはより狭い権限を受け取ります。

親の場合、目的が限定され、期限が限定され、範囲がサブタスクに限定され、慣例によるものではなく、おそらくそうであることがわかります。移動するにつれて縮小する委任は構造的な修正です。検出は構造物に障害が発生した場合のバックストップであり、代替品ではありません。
証拠: 改ざん防止ログは許可されていません
証拠のはしごは、不変で整合性が検証された監査証跡と完全な来歴チェーンで頂点に達します。これは良好な観察性です。しかし、暗号化ログの整合性は、記録が変更されていないことを証明します。それは、その行為が、誰の権限の下、どのような制約の中で許可されたのかを証明するものではありません。完全に改ざんが明らかな不正なアクションのログであっても、依然として不正なアクションのログです。
認可を証明するのは認可トークンそのものです。これは、この特定のエージェントがこの特定の通話に対してこの特定のアクセスを許可されたことを示す検証可能な成果物です。ログは追加専用であるため信頼できるログ エントリではなく、独自の署名に基づいており、ログの保持者に関係なくチェックできるアーティファクトです。
ここは、アイデンティティに向けて会話を広げている人々と私が別れる場所でもあります。委任の評価と調整をリソースと実行の境界に押し出そうとする本能により、難しい問題が新しい場所に再現されます。リソースには委任の意図を評価するコンテキストがありません。リソースに完全な委任セマンティクスの実行を要求すると、強制を行わずに情報が漏洩します。リソースに必要なのはアンカーです。つまり、権限が存在し、そのアクションが原因であることの証明であり、認可ロジック自体ではありません。評価はコンテキストが存在する場所に留まります。結果は伝わります。認証情報保管アプローチは、ベアラー トークンは保持しますが、サーバー側で保持しますが、これもまた別のカテゴリに分類されます。つまり、秘密鍵をプロキシします。

それを排除するのではなく。
バインディング: エージェントは存在しません。
これらすべての根底には、フレームワークにはない概念があります。それはエージェントのアイデンティティを他のすべての基礎として始まりますが、アイデンティティが最初に来るのは正しいことです。ただし、エージェントの ID を開始条件として扱います。つまり、暗号化 ID を割り当て、証明書を発行し、ハードウェアを証明します。エージェントは単に存在し、すでに識別されており、そこからコントロールが外側に構築されます。
エージェントは突然存在するわけではありません。誰かがエージェントを立ち上がらせた。個人または組織がそれを作成し、自分の権限の一部をそれに委任しました。この行為は、エージェントをその権限を行使するプリンシパルに結びつけるものであり、チェーン全体の根幹となります。それが「誰の権限の下に」という答えのある質問になるのです。それをスキップすると、エージェントのアイデンティティは名前にすぎません。暗号的に偽造できず、誰にも接続されません。
フレームワークがこのバインディングを確立することはありません。事後的にプリンシパルを 2 か所 (どちらも決定の下流) で推論します。エージェントのセッションにはテレメトリ上のユーザー ID と組織 ID が含まれるため、ログの属性を特定できます。そして、混乱した代理のケースでは、高い権限を持つエージェントは、元のユーザーの意図を検証する必要があると指示されますが、そもそもその意図がどのようにリクエストにバインドされているかについてのメカニズムはありません。ログ内の帰属は拘束力ではありません。これはアクション後に記録された推測であり、ログを書いた人がエージェントを正しいプリンシパルにマッピングしたと信じています。
これが、「誰が、誰の権限の下で行動したか」に可観測性を追加することで答えることができない理由です。答えは根元で確立され、最後に再構築されるのではなく、引き継がれなければなりません。 AAuth では、バインディングは登録時に作成され、i に対して人間の存在が確認されます。

最初はエージェントと個人のつながりがあり、それ以降、エージェントの権限は構造的にプリンシパルに遡ります。すべての制約された呼び出し、すべての委任、すべてのトークンは、それを許可した当事者に解決されます。フレームワークのエージェントには ID があります。まだプリンシパルがいません。それが、何が行動したかを知ることと、誰がその責任を負うのかを知ることの違いです。
これはより優れた OAuth ではありません。プリミティブが少なくなります。
私たちの分野の反射神経は、私たちが持っている部分を改善することです。永続的な問題には有効期限の短いトークンが対応し、弱い委任にはより多くのロギングが対応し、不透明なアクションにはより優れた監査が対応します。基板のシフトはその逆の動きです。 API キー、ベアラー トークン、事前登録されたクライアント、リダイレクト ベースのブートストラップなどのプリミティブを調整するのではなく、削除します。それは、より効率的な内燃エンジンと電気ドライブトレインの違いです。より良いキャブレターではありませんが、キャブレターはありません。
アントロピックは会話を適切な場所に移し、それを決定するテストを書き留めました。ベアラー トークンはそのテストに合格しません。エージェント プラス ツール、ログによる委任、不変ログによる証明、またはその下にプリンシパルのないエージェント ID のレベルでの認可も行われません。プロトコル層でのテストに合格することが AAuth の目的です。
ファウン

[切り捨てられた]

## Original Extract

Anthropic’s Zero Trust for AI Agents Sets the Right Test. The Bearer Token Fails It.
Where the framework stops, and where an agent-native authorization substrate has to begin.
Anthropic’s Zero Trust for AI Agents is a detailed and well-constructed enterprise security framework, and it is worth reading in full. It gets the starting point right: agents are not chatbots. They interpret goals, choose tools, persist context across sessions, coordinate with other agents, and act. And it states the consequence plainly: traditional access controls will not stop an agent from misusing permissions it legitimately holds. That is the correct frame.
I want to take the framework seriously enough to push on it. Not because it is wrong, but because it contains the exact test that retires its own baseline recommendation.
The test the document sets for itself
The strongest idea in the document is a design test. When you evaluate any control, the guide says to ask one question:
does this make the attack impossible, or just tedious?
Anthropic is explicit that controls whose value is friction, such as extra hops, rate limits, non-standard ports, and SMS codes, degrade against an adversary with unlimited patience and near-zero per-attempt cost. The controls that survive share a pattern: cryptographic identity, credentials that cannot be exfiltrated, and network paths that do not exist rather than paths that are merely inconvenient. When in doubt, they say, prefer a control that removes a capability over one that throttles it.
That is the right test. Apply it honestly and it reaches further than the document does.
Credentials: short-lived tokens are tedious, not impossible
The credential baseline is short-lived, narrowly scoped tokens from an identity provider, hardening to mutual TLS and then to hardware-bound credentials at the top tier. The document is blunt that static API keys and shared service-account passwords should be treated as already compromised, among the first things a model-assisted attacker finds in your code.
But run the baseline through their own test. A short-lived bearer token is a throttle. It shrinks the window during which a stolen secret is useful. It does not remove the secret. The token is still a bearer credential: whoever holds it can use it. Shortening its life raises cost without closing the door. By the document’s own standard, that is tedious, not impossible.
The control that makes credential theft impossible rather than tedious is the one where there is no bearer secret to steal: proof of possession bound to a private key that never leaves the agent. The agent signs its requests. The key is non-exportable. A compromised host yields no reusable credential, because the credential was never a string. It was the ability to produce a signature. HTTP Message Signatures is exactly this, and it is the substrate AAuth is built on.
Worth noting what is absent. The document does not mention DPoP, OAuth Token Exchange, Rich Authorization Requests, or sender-constrained tokens of any kind. It reaches toward non-exfiltratable credentials only at the Advanced tier, and frames that as PKI hardening for high-assurance shops rather than as the recognition that the bearer token itself is the residual flaw at every tier below. Their principle points past their prescription. The honest endpoint of “impossible, not tedious” is to stop issuing secrets that can be carried away.
Authorization: “the agent has access” is not granular enough
Here is the larger gap, and it is not about how the credential is held. It is about what gets authorized.
The framework describes authorization as a ladder of control families: role-based access control, then attribute-based access control, then continuous authorization, paired with a lifecycle ladder of static roles, dynamic elevation, and just-in-time grants. This is a reasonable taxonomy of access-control categories. What it is not is a mechanism for authorizing a specific call.
In this model the unit of authorization is the agent and the tool. An email agent gets email permissions. A database tool gets read but not write. Attribute-based control adds context such as time of day or a risk score. But the grant is still “this agent may use this tool,” possibly narrowed by coarse attributes. The actual parameters of the call are handled separately, as input validation: reject malformed or suspicious arguments before execution. Validation is a filter against bad input. It is not an authorization decision in which the specific parameters are the thing being authorized.
That distinction is the whole game for agents. “The agent has access to the payments tool” is not a useful grant. “This agent may move at most $3,000, to this payee, once, before Friday” is. The parameters are not noise to be sanitized. They are the authorization context, and they have to be evaluated as such, per request, at the moment the call is made.
This is what AAuth’s R3 provides: constrained calls that are specific about what the parameters may be, where those parameters are the context that gets evaluated. Authorization is bound to the actual operation, not to a standing relationship between an agent and a tool. The document gestures at the right instinct: its Advanced tier asks you to
evaluate authorization at each action rather than session start.
But it lists this as an aspirational top-tier capability and never says how the evaluation works or what it binds to. The how is the part that matters.
Others responding to the framework have made the same granularity point, framing it as continuous authorization and noting that a generic gateway cannot express a constraint like “search flights and book up to a limit, but never touch payment methods.” The diagnosis is right. The mechanism is where it has to land: the constrained call itself as the authorized artifact, with the parameters as the evaluated context, rather than policy expressiveness bolted onto a gateway.
Delegation: diagnosed precisely, then under-prescribed
The threat section is sharp on delegation. It names unscoped privilege inheritance, where a manager agent passes its full access context to a worker that should have less, and the confused-deputy amplification, where a low-privilege agent relays plausible instructions to a high-privilege one that executes them without checking the original intent. That is the real shape of multi-agent risk.
The prescribed remedy, though, is detection: establish trust boundaries, verify identity at each hop, and log inter-agent communications to flag unusual delegation patterns for review. Logging is not a delegation model. It tells you afterward that authority leaked. It does not stop authority from leaking.
The document is candid about where this bites in practice. Its own sub-agents can inherit up to the parent’s full permissions, with the distinction captured in telemetry rather than enforced in scope. So the failure mode flagged earlier is present by design, mitigated by visibility.
The removal-not-throttle answer is derived authority. When one agent delegates to another, the sub-agent receives authority that is narrower than the parent’s: purpose-bound, time-bound, scoped to the subtask, and provably so rather than by convention. Delegation that narrows as it travels is the structural fix. Detection is the backstop for when structure fails, not a substitute for it.
Proof: tamper-evident logs are not authorization
The evidence ladder tops out at immutable, integrity-verified audit trails and full provenance chains. This is good observability. But cryptographic log integrity proves one thing: that the record was not altered. It does not prove that the action was authorized, under whose authority, within what constraint. A perfectly tamper-evident log of an unauthorized action is still a log of an unauthorized action.
What proves authorization is the authorization token itself: a verifiable artifact showing that this specific agent was granted this specific access, for this specific call. Not a log entry you trust because the log is append-only, but an artifact that stands on its own signature and can be checked independently of whoever holds the log.
This is also where I part company with some who are extending the conversation toward identity. The instinct to push delegation evaluation and reconciliation out to the resource and execution boundary recreates the hard problem in a new place. Resources do not have the context to evaluate delegation intent, and asking them to carry full delegation semantics leaks information without buying enforcement. What a resource needs is an anchor: proof that authority existed and that the action is attributable, not the authorization logic itself. The evaluation stays where the context lives. The outcome travels. Credential-vaulting approaches, which keep the bearer token but hold it server-side, sit in a different category again: they proxy the secret rather than eliminate it.
Binding: the agent does not pop into existence
Underneath all of this sits a concept the framework does not have. It opens with agent identity as the foundation for everything else, and it is right that identity comes first. But it treats the agent’s identity as a starting condition: assign a cryptographic identifier, issue a certificate, attest the hardware. The agent simply exists, already identified, and the controls build outward from there.
Agents do not pop into existence. Someone stood the agent up. A person or an organization created it and delegated some slice of their own authority to it. That act, binding an agent to the principal whose authority it exercises, is the root of the entire chain. It is what makes “under whose authority” a question with an answer. Skip it and agent identity is only a name: cryptographically unforgeable, and connected to no one.
The framework never establishes this binding. It infers the principal after the fact, in two places, both downstream of any decision. The agent’s session carries a user identifier and an organization identifier on its telemetry, so logs can be attributed. And in the confused-deputy case, a high-privilege agent is told it should verify the original user’s intent, with no mechanism for how that intent was bound to the request in the first place. Attribution in a log is not a binding. It is a guess recorded after the action, trusting that whoever wrote the log mapped the agent to the right principal.
This is why “who acted, under whose authority” cannot be answered by adding observability. The answer has to be established at the root and carried forward, not reconstructed at the end. In AAuth the binding is created at enrollment, with confirmed human presence for the initial agent-to-person link, and from then on the agent’s authority traces to a principal by construction. Every constrained call, every delegation, every token resolves back to a party who granted it. The framework’s agent has an identity. It does not yet have a principal. That is the difference between knowing what acted and knowing who is accountable for it.
This is not a better OAuth. It is fewer primitives.
The reflex in our field is to improve the parts we have: persistent problems met with shorter-lived tokens, weak delegation met with more logging, opaque actions met with better audit. The substrate shift is the opposite move. It deletes primitives rather than tuning them: API keys, bearer tokens, pre-registered clients, redirect-based bootstrapping. It is the difference between a more efficient combustion engine and an electric drivetrain. Not a better carburetor, but no carburetor.
Anthropic moved the conversation to the right ground and wrote down the test that decides it. The bearer token does not pass that test. Neither does authorization at the level of agent-plus-tool, delegation-by-logging, proof-by-immutable-log, or an agent identity with no principal beneath it. Passing the test at the protocol layer is what AAuth is for.
Foun

[truncated]
