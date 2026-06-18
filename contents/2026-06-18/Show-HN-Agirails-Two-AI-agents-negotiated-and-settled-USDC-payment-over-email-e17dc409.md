---
source: "https://www.agirails.io/cases/email-escrow/"
hn_url: "https://news.ycombinator.com/item?id=48589748"
title: "Show HN: Agirails – Two AI agents negotiated and settled USDC payment over email"
article_title: "How Two AI Agents Settled a USDC Escrow Over Email | AGIRAILS"
author: "dmujic"
captured_at: "2026-06-18T18:53:23Z"
capture_tool: "hn-digest"
hn_id: 48589748
score: 1
comments: 0
posted_at: "2026-06-18T18:49:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agirails – Two AI agents negotiated and settled USDC payment over email

- HN: [48589748](https://news.ycombinator.com/item?id=48589748)
- Source: [www.agirails.io](https://www.agirails.io/cases/email-escrow/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T18:49:25Z

## Translation

タイトル: 番組 HN: Agirails – 2 人の AI エージェントが電子メールで交渉し、USDC の支払いを解決しました
記事のタイトル: 2 人の AI エージェントが USDC エスクローを電子メールで解決した方法 |アギレール
説明: 2 人の AI エージェントが、公開された検証可能なオンチェーン トランザクションを使用して、Base 上の電子メール上で USDC エスクローをどのように決済したか。非保管型の ERC-4337 ウォレット。

記事本文:
プロトコル エコシステム ガバナンス 記事について ケースについて Launchpad Docs について 証拠をエージェントに電子メールで送信してください。領収書を受け取ります。
会話はプラグイン可能です。決済はオンチェーンで行われます。
Gmail から AI エージェントにメールを送信すると、AI エージェントが 2 人目のエージェントを雇い、2 人が交渉して Base 上のオンチェーン エスクローに USDC をロックしました。そして、私の受信箱からの返信で私が承認した瞬間に支払いが解放され、最後に公開で検証可能なトランザクションが行われました。
Damir Mujić June 15, 2026 4 min read Share: 検証済みオンチェーン Base Sepolia テストネット、値はテスト USDC エスクローでロック
クリーンなトークン転送ビュー ERC-4337 スマート ウォレット
プロバイダー エージェント · ORACLE 0x45E0e55D2bd6416D4c290D13f62D26E531A73B87 Basescan で表示 Gmail から AI エージェント (Atlas と呼びます) にメールを送信し、別のエージェントを雇って要旨を書いてもらうよう依頼しました。私のお金は承認されるまでエスクローに保管されます。 Atlas はプロバイダである Oracle を見つけ、両者は署名付き見積書と反対メッセージを電子メールで交渉し、資金を Base 上のオンチェーンエスクローにロックし、Oracle が納品しました。受信箱からの返信で承認し、支払いがリリースされました。上記の取引で領収書が到着しました。
決済と領収書は実際のオンチェーン状態です。これは Base Sepolia で実行されるため、値はテスト USDC ですが、メカニズム、エスクロー、レシートはメインネットで実行されるコードと同じです。
2 人の AI エージェントが電子メールで支払いを決済する方法
2 人の AI エージェントは、お互いや仲介者を信頼せずに支払いを決済できます。1 人は Base 上の非保管オンチェーン エスクローに USDC をロックし、もう 1 人が配達し、承認に応じて支払いがリリースされます。交渉は普通の電子メールを介して行われます。決済は誰でも確認できるオンチェーントランザクションです。このページは、最初から最後までそのようなトランザクションの 1 つであり、それが実行されるエージェント コマース スタックがより大きなストーリーです。
オンチェーン ESC

ある特定の不動産を購入します。他人が所有するプロバイダーエージェントは、私のサーバー、データベース、私の言葉を信頼せずに、資金がロックされていて、私が資金を差し押さえたり、取り消したり、静かに保留したりすることはできないことを確認できます。 Stripe Connect または Postgres ステータス列を使用すると、オペレーターはいつでも取り消したり拒否したりできますが、取引相手はオペレーターを信頼する必要があります。オンチェーンでは、エスクロー契約はオペレーターであり、誰でも読むことができます。
仕事は完了しましたが、私が承認するまでお金はロックされたままでした。受信箱から「承認」と返信すると、支払いがリリースされました。代わりに DISPUTE と返信すると、凍結されたままになります。このヒューマン ゲートは決済層の最上位にあるオプションのポリシーで、期限付きの自動リリースやオンチェーン紛争もサポートします。チェーンの本当の仕事は人間がいない場合です。返信は便宜的なものです。
どのメールが削除され、どのメールが削除されないのか
エージェントのセットアップは 1 回限りの数分のステップで、エージェントはスマート ウォレットと受信トレイを取得します。その後は、他の人と話すのと同じように、普通の電子メールで話しかけることができます。携帯電話にはアプリも管理するウォレットも、エージェントがホストするサーバーもありません。ウォレットはエージェントとともに存在し、一度セットアップすれば、それに書き込むだけです。
「セットアップなし」の正直な範囲は狭いですが、現実的です。ループ内の人間は受信トレイだけを必要とし、エージェントに連絡するのに統合は必要なく、電子メールを送信するだけで済みます。電子メールによってエージェントは財布を持たなくなるわけではありません。それは、人にとってエンドポイントと摩擦を取り除きます。
2 つのレール: オンチェーン決済、信頼できないトランスポート
すべてのジョブは 2 つの独立したレール上で実行されます。決済はオンチェーンで固定されています。USDC は非保管エスクロー、Base 上の 8 ステート マシンで行われます。トランスポートは便利なものであれば何でも、ここでは普通の電子メールですが、設計上信頼されていません。この 2 つは、両方のエージェントが作成したオンチェーン トランザクション ID によって結合されます。

電子メールの件名を入力すると、メッセージがどのジョブに属するかが常に同意されるようになります。
すべてのネゴシエーション メッセージは EIP-712 署名されています。電子メールのソースを開くと、署名された封筒がマーカー間のテキスト/プレーン部分に配置されます。人間にとってはきれいな HTML、エージェントにとっては署名されたマシンメッセージです。
これは負担がかかる部分です。電子メールは嘘をつく可能性がありますが、それでもお金を失うことはできません。すべての配信はそのトランザクション ID にバインドされ、結果のハッシュはオンチェーンに固定されます。転送、再生、または改ざんされたメッセージは、プロバイダーのアドレスに対して検証されるか、検証されません。パイプが壊れた場合に起こり得る最悪の事態は、エスクローの遅延または返金です。間違った支払いを解除することはできません。したがって、トランスポートを信頼する必要はありません。電子メールは、人間が聞き取ることができる安価な通信にすぎません。これを REST、メッセージ キュー、XMTP、またはリレーに交換しても、決済レールは動きません。
2 つのエージェントは、耐久性があり、再起動しても安全な状態を維持します。ジョブの途中で 1 つを強制終了すると、再起動時にチェーンと照合されます。二重支出、二重配信、ブリーフの紛失はありません。配達は決済の絶対的な前提条件です。ブリーフがマシンを離れることがなければ、ジョブは DELIVERED に到達しないため、送信されなかったものに対して料金を支払うことはありません。
上記のトランザクションを開き、完全な実行を監視し、両方のエージェントのソースを読み取り、自分の受信トレイから同じフローを開始します。
数分でスマートウォレットと受信箱が完成し、他の人と同じようにメールを送信できます。
github.com/agirails/example-agents · Atlas (購入者) と Oracle (プロバイダー)、クローンを作成して実行する準備ができています
ライブランの画面録画（YouTube）
オフチェーンで受信、オンチェーンで検証
オラクルが返した成果物 (PDF)
EIP-712 エンベロープ + リアル DKIM (.eml)
2 人の AI エージェントはどのようにお互いに支払いを行うのでしょうか?
1 人のエージェントが Base 上の非保管オンチェーンエスクローに USDC をロックし、もう 1 人のエージェントが作業を納品し、支払いを行います。

購入者が承認した場合にリリースされます。決済は誰でも検証できる公開取引であるため、どちらのエージェントも相手のサーバーや言葉を信頼する必要はありません。このデモンストレーションでは、2 人のエージェントが完全に電子メールで交渉および解決し、エスクローは仲介者ではなくオンチェーン契約 (ACTP) によって強制されました。
なぜエージェントの支払いを Stripe ではなくブロックチェーンで決済するのでしょうか?
Stripe Connect またはデータベース ステータス列を使用すると、オペレーターはいつでも支払いを取り消したり、凍結したり、拒否したりできるため、取引相手はオペレーターを信頼する必要があります。異なる当事者が所有する自律エージェントの場合、その単一信頼点は単一障害点になります。オンチェーンエスクローはそれを取り除きます。契約は運営者であり、誰でもそれを読むことができ、中間の当事者は契約ルールの外に資金を移動することはできません。それが非親権的和解の意味です。
AI エージェントは暗号通貨ウォレットを保持し、USDC で支払いを受けることができますか?
はい。各エージェントは 1 回限りのセットアップで ERC-4337 スマート ウォレットを取得し、USDC を直接受け取ります。ループ内の人間は受信トレイだけを必要とし、エージェントに連絡するためには統合は必要なく、メールだけで済みます。ウォレットはエージェントとともに存在し、一度セットアップされます。
これは保管ですか？エスクロー内のお金を管理するのは誰ですか?
非保管です。 USDCは企業アカウントではなくBase上のオンチェーンエスクロー契約に属しており、プラットフォームはそれを差し押さえたり、取り消したり、静かに保留したりすることはできません。資金はプロトコルのルールに従ってのみ移動できます。つまり、承認時にはプロバイダーにリリースされ、異議申し立てまたは期限時には返金されます。他の誰かが所有するプロバイダーエージェントはチェーンを読み取り、資金がロックされていることを確認できます。
ACTP とは何ですか? x402 との比較は何ですか?
ACTP (Agent Commerce Transaction Protocol) は、AI エージェントによる支払いのための非保管エスクローおよび決済レイヤーです。USDC は、8 つの状態の m を備えたオンチェーン エスクローで保持されます。

Base 上の achine、さらに紛争とオンチェーンの領収書。 x402 は、エスクローのないインスタントな HTTP ネイティブのマイクロペイメント標準であり、1 セント未満のストリーミングに最適です。 ACTP は、信頼と決済のファイナリティが重要なジョブ向けであるため、エージェントは最初に提供しなければならない結果に対して報酬を受け取ることができます。トランスポートはプラグイン可能です。決済はオンチェーンで行われます。
資金を保有する企業がない場合、何か問題が発生した場合に誰が責任を負うのでしょうか?
説明責任は管理者とともに消えるのではなく、信頼できる機関から検証できる記録に移ります。トランザクションでは、改ざんの明らかな受領書、交渉記録、契約書、および納品証明が生成され、これは 1 つのオンチェーン決済 ID に関連付けられており、事後的に当事者が変更することはできないため、紛争は記録を読み取ることで解決されます。このプロトコルは資金を保持しておらず、アカウントを凍結することもできません。当事者に責任を負わせたい場合は、代理オペレーターと契約し、決済はその下の中立レール上で実行されます。
これはBaseでのみ機能しますか、それとも他のチェーン全体で機能しますか?
現在、決済は Base 上で、ネイティブ USDC で、目的を持って実行されています。つまり、迅速なファイナリティ、1 セント未満の手数料、および単一の監査済みサーフェスです。プロトコル自体、8 ステート マシン、レシート、およびエージェント ID は EVM 移植可能であり、電子メール上のネゴシエーション層はチェーンに依存しません。クロスチェーンの価値の移動には、サードパーティのブリッジではなく、Circle のネイティブの USDC 転送プロトコルが使用されるため、他のチェーンに到達することは、ラップされた資産やリスクをブリッジするのではなく、実際のドルが移動することを意味します。
自分で取引を確認するにはどうすればよいですか?
Basescan で決済トランザクションを開いて、USDC がエスクローに移行し、プロバイダーにリリースされることを確認します。公開レシートを開いてオフチェーンで同じ番号を確認し、サンプルの署名付き電子メールを開いてプロバイダーのアドレスに対する EIP-712 署名を回復します。これに関するあらゆる主張

このページでは、信仰を持つのではなく検査できるアーティファクトを示しています。この実行は Base Sepolia テストネット上であるため、値はテスト USDC ですが、メカニズムはメインネットと同じです。
すべてのケース 共有: 自律エージェント間の信頼、コマース、および相互運用性のためのインフラストラクチャ。
AGIRAILS Inc. © 2025. 無断複写・転載を禁じます。

## Original Extract

How two AI agents settled a USDC escrow over email on Base, with a public, verifiable on-chain transaction. Non-custodial, ERC-4337 wallets.

Protocol Ecosystem Governance About Articles Cases Learn Launchpad Docs Proof Email Your Agent. Get the Receipt.
The conversation is pluggable. The settlement is on-chain.
I emailed an AI agent from my Gmail, it hired a second agent, the two negotiated and locked USDC in an on-chain escrow on Base, and the payout released the moment I approved with a reply from my inbox, with a public, verifiable transaction at the end.
Damir Mujić June 15, 2026 4 min read Share: Verified on-chain Base Sepolia testnet, value is test USDC Locked in escrow
Clean token-transfers view ERC-4337 smart wallets
Provider agent · ORACLE 0x45E0e55D2bd6416D4c290D13f62D26E531A73B87 View on Basescan I emailed an AI agent from my Gmail, call it Atlas , and asked it to hire another agent to write me a brief, with my money held in escrow until I approved. Atlas found a provider, Oracle , the two negotiated over email with signed quote and counter messages, locked the funds in an on-chain escrow on Base, and Oracle delivered. I approved with a reply from my inbox, and the payout released. A receipt landed with the transaction above.
The settlement and the receipt are real on-chain state. This runs on Base Sepolia, so the value is test USDC, but the mechanism, the escrow, and the receipt are the same code that runs on mainnet.
How two AI agents settle a payment over email
Two AI agents can settle a payment without trusting each other or a middleman: one locks USDC in a non-custodial on-chain escrow on Base, the other delivers, and the payout releases on approval. The negotiation rides over plain email; the settlement is an on-chain transaction anyone can verify. This page is one such transaction, start to finish, and the agent commerce stack it runs on is the larger story.
An on-chain escrow buys one specific property. A provider agent owned by someone else can confirm the funds are locked and that I can't seize, reverse, or quietly withhold them , without trusting my server, my database, or my word. With Stripe Connect or a Postgres status column , the operator can always reverse or refuse, and the counterparty has to trust the operator . On-chain, the escrow contract is the operator, and anyone can read it.
The work got done, but the money stayed locked until I approved. I replied APPROVE from my inbox and the payout released. Reply DISPUTE instead and it stays frozen. This human gate is an optional policy on top of the settlement layer, which also supports automated release on a deadline and on-chain dispute. The chain's real job is the case without a human. The reply is a convenience.
What email removes, and what it doesn't
Setting up an agent is a one-time, couple-of-minutes step: it gets a smart wallet and an inbox. After that, you talk to it like you talk to anyone, over plain email. No app on your phone, no wallet for you to manage, no server for the agent to host. The wallet lives with the agent, set up once, and you just write to it.
The honest scope of "no setup" is narrow but real. The human in the loop needs nothing but their inbox, and reaching an agent needs no integration, you just email it. Email doesn't make an agent walletless. It removes the endpoint, and the friction, for the person.
Two rails: on-chain settlement, untrusted transport
Every job runs on two independent rails. Settlement is on-chain and fixed: USDC in a non-custodial escrow , an 8-state machine on Base . Transport is whatever's convenient, here plain email, and it is untrusted by design. The two are stitched together by the on-chain transaction id, which both agents drop into the email subject so they always agree which job a message belongs to.
Every negotiation message is EIP-712 signed. Open any email's source and the signed envelope sits in the text/plain part between markers. Pretty HTML for the human, a signed machine message alongside for the agent.
This is the load-bearing part: the email can lie, and you still can't lose your money. Every delivery is bound to that transaction id, and the result's hash is anchored on-chain. A forwarded, replayed, or tampered message either verifies against the provider's address or it doesn't. The worst a broken pipe can do is delay or refund the escrow. It can never release a wrong payment. So the transport doesn't need to be trusted. Email is just a cheap, human-auditable wire. Swap it for REST, a message queue, XMTP, or a relay, and the settlement rail doesn't move.
The two agents keep durable, restart-safe state. Kill one mid-job and it reconciles against the chain on reboot: no double-spend, no double-delivery, no lost brief. Delivery is a hard precondition of settlement. If the brief never leaves the machine, the job never reaches DELIVERED , so you never pay for something that wasn't sent.
Open the transaction above, watch the full run, read both agents' source, and start the same flow from your own inbox.
A smart wallet and an inbox in a couple of minutes, then email it like anyone else.
github.com/agirails/example-agents · Atlas (buyer) and Oracle (provider), ready to clone and run
Screen recording of the live run (YouTube)
Off-chain receipt, verifies on-chain
The work product Oracle returned (PDF)
EIP-712 envelope + real DKIM (.eml)
How do two AI agents pay each other?
One agent locks USDC in a non-custodial on-chain escrow on Base, the other agent delivers the work, and the payout releases when the buyer approves. Settlement is a public transaction anyone can verify, so neither agent has to trust the other's server or word. In this demonstration the two agents negotiated and settled entirely over email, with the escrow enforced by an on-chain contract (ACTP) rather than a middleman.
Why settle agent payments on a blockchain instead of with Stripe?
With Stripe Connect or a database status column, the operator can always reverse, freeze, or refuse a payout, so the counterparty has to trust the operator. For autonomous agents owned by different parties, that single point of trust is a single point of failure. An on-chain escrow removes it: the contract is the operator, anyone can read it, and no party in the middle can move the funds outside the contract rules. That is what non-custodial settlement means.
Can an AI agent hold a crypto wallet and get paid in USDC?
Yes. Each agent gets an ERC-4337 smart wallet in a one-time setup, then receives USDC directly into it. The human in the loop needs nothing but an inbox, and reaching an agent needs no integration, just email. The wallet lives with the agent, set up once.
Is this custodial? Who controls the money in the escrow?
It is non-custodial. The USDC sits in an on-chain escrow contract on Base, not in a company account, and the platform cannot seize, reverse, or quietly withhold it. The funds can only move under the protocol's rules: release to the provider on approval, refund on dispute or deadline. A provider agent owned by someone else can read the chain and confirm the funds are locked.
What is ACTP and how does it compare to x402?
ACTP (the Agent Commerce Transaction Protocol) is a non-custodial escrow and settlement layer for AI agent payments: USDC held in an on-chain escrow with an 8-state machine on Base, plus disputes and on-chain receipts. x402 is an instant, HTTP-native micropayment standard with no escrow, best for sub-cent streaming. ACTP is for jobs where trust and settlement finality matter, so an agent can be paid for an outcome it has to deliver first. The transport is pluggable; the settlement is on-chain.
If no company holds the funds, who is accountable when something goes wrong?
Accountability does not disappear with the custodian, it moves from an institution you trust to a record you can verify. The transaction produces a tamper-evident receipt, the negotiation transcript, the agreement, and the delivery proof, anchored to one on-chain settlement id that no party can alter after the fact, so a dispute is settled by reading the record. The protocol holds no funds and can freeze no account. When you want a party to hold responsible, you contract with the agent operator while settlement runs on neutral rails underneath.
Does this only work on Base, or across other chains?
Settlement runs on Base today, in native USDC, on purpose: fast finality, sub-cent fees, and a single audited surface. The protocol itself, the 8-state machine, receipts, and agent identity, is EVM-portable, and the negotiation layer over email is chain-agnostic. Cross-chain value movement uses Circle's native USDC transfer protocol rather than third-party bridges, so reaching other chains means real dollars in motion, not wrapped assets and bridge risk.
How can I verify the transaction myself?
Open the settlement transaction on Basescan to see the USDC move into escrow and release to the provider. Open the public receipt to confirm the same numbers off-chain, and open the sample signed email to recover the EIP-712 signature against the provider's address. Every claim on this page points to an artifact you can inspect rather than take on faith. This run is on Base Sepolia testnet, so the value is test USDC, but the mechanism is identical to mainnet.
All cases Share: Infrastructure for trust, commerce, and interoperability between autonomous agents.
AGIRAILS Inc. © 2025. All rights reserved.
