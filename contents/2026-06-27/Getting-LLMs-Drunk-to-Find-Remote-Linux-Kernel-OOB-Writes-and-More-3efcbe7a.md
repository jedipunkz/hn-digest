---
source: "https://heyitsas.im/posts/drinking-llms/"
hn_url: "https://news.ycombinator.com/item?id=48696719"
title: "Getting LLMs Drunk to Find Remote Linux Kernel OOB Writes (and More)"
article_title: "Getting LLMs Drunk to Find Remote Linux Kernel OOB Writes (and More) · Hey, it's Asim"
author: "schmuhblaster"
captured_at: "2026-06-27T10:38:59Z"
capture_tool: "hn-digest"
hn_id: 48696719
score: 1
comments: 0
posted_at: "2026-06-27T09:46:53Z"
tags:
  - hacker-news
  - translated
---

# Getting LLMs Drunk to Find Remote Linux Kernel OOB Writes (and More)

- HN: [48696719](https://news.ycombinator.com/item?id=48696719)
- Source: [heyitsas.im](https://heyitsas.im/posts/drinking-llms/)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T09:46:53Z

## Translation

タイトル: LLM を酔わせてリモート Linux カーネル OOB 書き込み (その他) を見つける
記事のタイトル: LLM を酔わせてリモート Linux カーネルの OOB 書き込み (その他) を見つける · やあ、Asim です
説明: 自己調整型のエージェント チームを使用し、アクティベーションの操作をわずかに行い、Linux カーネルから Docker や OpenSSL に至るまであらゆるものの脆弱性を発見します。

記事本文:
↓
メインコンテンツにスキップ
やあ、アシムだよ
RSS
RSS
LLM を酔わせてリモート Linux カーネル OOB 書き込み (その他) を見つける
著者
アシム ヴィラディ オグル マニサダ
目次
背景
建築
わかりましたが、なぜ「酔った」のでしょうか？
…（苦い）教訓
クリエイティビティに一石を投じる必要はない
より多くのレイヤーを積み重ねると、常に抽象化チェーンを上に移動します。
推論時のコンピューティングでは、生のモデル インテリジェンスはそれほど重要ではありません
建築
わかりましたが、なぜ「酔った」のでしょうか？
…（苦い）教訓
クリエイティビティに一石を投じる必要はない
より多くのレイヤーを積み重ねると、常に抽象化チェーンを上に移動します。
推論時のコンピューティングでは、生のモデル インテリジェンスはそれほど重要ではありません
TLDR: 以下に詳述する、著しく過剰設計で自己組織化した脆弱性探索エージェントのチームは、過去数か月間で CVE-2026-31432 および CVE-2026-31433 を含む 20 以上の CVE を発見しました。これは、Linux カーネルの ksmbd における 2 つのリモートの未認証 OOB 書き込みです。これを実現するセットアップの詳細については、「はい!」を含めて読み続けてください。 – LLM を酔わせる。
「LLMing」脆弱性研究は、DARPA の AIxCC と XBOW の最初の結果以来、私の「これについて何かをする」リストに載っていました。しかし、2023 年から 2024 年にかけて、モデルを何か役に立つものにするためには多くの活用が必要で、ツールの使用は初歩的で、モデルのコンテキストにできるだけ多くのコードを押し込んで、誤検知をトリアージして取り除くという考えに、私は恐怖でいっぱいでした。
実際に何かをしようという動きがあったのは、2025 年の夏でした。Rich Mirch は、非常に単純で 12 年間気づかれなかった sudo でのローカル権限昇格を報告しました: CVE-2025-32462 。ドキュメントとは異なり、 --host フラグは、別のホストでの権限のリストを許可するだけでなく、sudo ルールのホスト名の部分を無関係にしました。たとえば、sudoers ルールにより、あるホストでの root 権限が付与された場合、

ローカルホストではない場合、フラグを悪用してローカルで完全なルートを取得する可能性があります。
この LPE は LLM では見つかりませんでしたが (AFAICT)、私は疑問に思いました。LLM にさまざまなツールを実行させる代わりに、ドキュメントと実際のコードの間の (馬鹿げた単純な) 不一致を探させたらどうなるでしょうか?コンテキストのサイズ、複雑さの利用、必要なインテリジェンスの点で、(ローカル) LLM にとってはより簡単なリフトのように思えました。これらは技術的に最も興味深い発見ではありませんが、実際の影響は同様に深刻です。影響という点では、LPE は LPE です。
2025 年末までに、私はまさにこれを実現するためのハーネスの開発に着手しました。しかし、マイク・タイソンの言葉を借りれば、誰もが新しいモデルが発売されるまでの計画を持っています。ハーネスが完成するとほぼすぐに、モデルは十分な品質になり、コンテキストを多用する外部ツールを使用する場合でも、必要な足場を大幅に簡素化できるようになりました。この時点で、私の探求は 3 つに分かれました。
上記の発見に触発された当初の目標である「ドキュメント ↔ コードの不一致」タイプの脆弱性を見つけることができるでしょうか?
機能が段階的に変化していることを考えると、脆弱性一般はどうなるのでしょうか?
もっと推測的に言えば、LLM から「動き 37」を取得して、a) まったく新しいバグ クラスを見つけるか、少なくとも b) より小さなモデルで何かのロックを解除してハンティング能力を強化することはできるでしょうか?
答えは大まかに「はい」、「はい」、そして「たぶん」でした。以下は、カスタム ハーネスを介して完全に自律的に検出された 30 件以上の調査結果 (2026 年 4 月 29 日の時点で 20 件以上の CVE が割り当てられており、一部は未公開) です。差し迫った雪崩を考慮して、ネットワークに到達可能なサービスを最初に優先しました。
ksmbd (Linux のカーネル内 SMB サーバー) CVE-2026-31432 および CVE-2026-31433 は両方ともリモートで、認証されていない (ゲスト共有を使用している場合) OOB 書き込みです。どちらのバグでも、リモート クライアントは複数のファイル共有操作を 1 つにまとめることができます。

request – 最初の操作はカーネルの応答バッファーのほぼすべてを (正当に) 使用でき、次の操作は適切な境界チェックなしで可変長のメタデータを追加するため、OOB が発生します。 CVE-2026-31432 はさらに興味深いものです。私の研究室では、シリアル化された QUERY_INFO(Security) からの攻撃者ペイロード由来のバイトが隣接するカーネル オブジェクトに到達し、 filp_flush / dnotify_flush にヒットし、シリアル化された応答の末尾と正確に一致するバイトを含む構造体ファイルが破損しました。十分な Codex/Claude クレジットと人工ラボ環境 (最新の強化がオフになっている) があれば、おそらくおもちゃの RCE PoC を取得できるでしょう。
クラス的には、これらは退屈なオーバーフローです。発見するには、集中的な専門家の注意が必要です (LLM の前には不足していましたが、現在は豊富です)。調整された「酔っぱらった」Qwen 3.5 27B 派生版を実行しているハーネス (下記の「アーキテクチャ」を参照) は、検証ツールを使用して ksmbd を数日間繰り返した後、これらを発見しました。ただし、 gpt-5.3-codex も同様で、ハーネスに接続するとはるかに高速になりました。
CVE-2026-34980 および CVE-2026-34990 は、「スプーラー アラート: CUPS におけるリモートの認証されていない RCE からルートへのチェーン」で詳しく説明されているように、認証されていないリモート攻撃者 -> 権限のない RCE -> ルート ファイル (上書き) に連鎖する 2 つの CUPS の問題です。これは本当に楽しかったです。適切に分解された目標に基づいて、小規模な LLM をどこまで推進できるかを示しました。ネットワーク上に非特権の足場を確立し、非特権ユーザーから root にエスカレートして、それぞれの問題をさらに細分化する足場を各エージェントに課すというタスクです。
これは、もともとこのすべての作業の動機となったカテゴリです。多くの発見はここにあります。
Docker (CVE-2026-34040): AuthZ プラグインは生のリクエスト本文を参照するように文書化されていますが、細工された API リクエストにより、プラグインが本文なしでリクエストを承認する可能性があります (デーモンはこれを許可します)

その後実行します)。
Caddy (CVE-2026-27587、CVE-2026-27588): MatchHost と MatchPath は両方とも、大文字と小文字を区別せずにホストと URI パスを一致させるものとして文書化されていましたが、特定の状況下では一致しませんでした。
udisks (CVE-2026-26103、CVE-2026-26104): これらは少し境界線にありました。 udisks には polkit auth が必要であることが文書化されており、HeaderBackup と RestoreEncryptedHeader は認証が必要なものとして明示的にリストされていませんでしたが、エージェントは依然として 2 つの関連する D-Bus メソッドが (他のアクションのメソッドとは異なり) 認証チェックを呼び出していないことに気付きました。これにより、ローカルの権限のない攻撃者が、a) LUKS ヘッダーをバックアップし、b) ヘッダーを上書きすることで既存の LUKS デバイスをブリックすることが可能になりました。
Firewalld (CVE-2026-4948): polkit ポリシー ファイルには、ファイアウォールの状態の変更には特定のアクセス許可を適用する必要があると記載されていますが、バルク ランタイム ポリシー セッターは代わりに、構成検査用のアクセス許可によって保護されており、ローカルの特権のないユーザーによるランタイム状態の変更を許可していました。
util-linux (CVE-2026-3184): これは最も深刻ではありませんが、もともと全体のきっかけとなった上記の sudo CVE-2025-32462 に精神的に最も近いものでした。 login -h は、 telnetd などのサービスからリモート ホスト名を受け入れるものとして文書化されています。 PAM ホストベースのアクセス制御は、アクセス ルールを評価するときにそれを使用できます。ただし、ログインは PAM_RHOST を設定する前に指定されたホスト名を正規化するため、PAM は実際に指定されたものとは異なる名前に対してポリシーを適用することになる可能性があります。
ハーネスは大きな進化を遂げました。最も拡張した状態では次のようになります。
フローチャート TD
Seeder["`**ターゲットシーダー**
ランク付けされたターゲット + 実行目標`"]
HypGen["`**仮説生成ツール**
ドキュメント、ソース、不変条件、
攻撃者の入力フロー`"]
ハンター["`**ハンター**
PoC の反復
隔離された VM`"]
RunFolder["`**実行ごとの f

古い**
成功、失敗、
PoC`"]
ライター["`**レポートライター**
メンテナ向き
レポート+PoC`"]
グレーダー["`**外部グレーダー**
深刻さ、新規性、
その他の健全性チェック`"]
送信["`**オペレーターによるレビューのために送信**`"]
問題["`**問題ログ**
サンドボックス/ツールのブロッカー`"]
指揮者["`**指揮者**
ポーリングハント、エージェントのリダイレクト、
全身性ブロッカーを追跡します`"]
Seeder --> HypGen --> Hunters --> RunFolder --> Writers --> Grader --> 送信
実行フォルダー -。 "ピボット/再試行" .-> HypGen
採点者 - 。 「フィードバック付きで拒否」 .-> HypGen
ヒプジェン - 。 「ブロッカーの追加」 .-> 問題
ハンターたち――。 「ブロッカーの追加」 .-> 問題
作家 - 。 「ブロッカーの追加」 .-> 問題
問題点 --> 指揮者
指揮者 - 。 「操縦」 .-> HypGen
指揮者 - 。 「操縦」 .-> ハンター
指揮者 - 。 "ステアリング" .-> ライター
classDef メインの塗りつぶし:#bfdbfe、ストローク:#60a5fa、ストローク幅:1.5px、色:#111827;
classDef グレーダーの塗りつぶし:#fee2e2、ストローク:#dc2626、ストローク幅:2.5px、色:#111827;
classDef の問題の塗りつぶし:#ddd6fe、ストローク:#a78bfa、ストローク幅:1.5px、色:#111827;
classDef 導体塗りつぶし:#c7d2fe、ストローク:#818cf8、ストローク幅:1.5px、色:#111827;
classDef ターミナルの塗りつぶし:#e0f2fe、ストローク:#38bdf8、ストローク幅:1.5px、色:#111827;
クラス Seeder、HypGen、Hunters、RunFolder、Writers main;
クラス採点者採点者。
クラス問題の問題。
クラス 導体 導体;
クラス サブミットターミナル;
linkStyle のデフォルトのストローク:#94a3b8、ストローク幅:1.4px;
ターゲットシーダーは、推定された「蔓延/実際の悪用可能性-潜在性」インデックスによってランク付けされたターゲットリストを作成し（最初のアイデアをすべて検討するまでは、これが私でした）、ターゲットの検出数/最小重大度で実行を開始します。
仮説ジェネレータは、ドキュメントとソース、ポリシーの不変条件、興味深いオプションの組み合わせ、攻撃者の入力フローなどを研究し、ハンターにとって有望な仮説を記録します。
フン

ターは、専用の分離された VM で仮説を反復し、実行ごとのフォルダーに PoC で成功/失敗を記録し、仮説ジェネレーターの次の反復をピボットできるようにします。サンドボックスは非常に重要です。エージェントは PoC を繰り返し調整し、競合状態のタイミングを計り、誤検知 (通常は十分な量) を排除するなど、現実との接触から恩恵を受けることができます。
レポート作成者は、すべてを保守者向けのレポートにパッケージ化し、レポートと PoC を外部の採点者に送信します。
外部グレーダー
常に別のモデルを呼び出して、a) ハントの事前設定された重大度/検出タイプ要件の一致、b) 新規性のチェック、および c) 他の健全性チェックの実行について検出結果を評価します。これにより、報酬ハッキング モデルが捕らえられ、最終的には断念し、結果が大幅に水増しされることになります。
最後に、車掌が継続的に狩りをポーリングし、関与するエージェントがハンドルを空転させたり、スプレーを吹きかけたり、その他の方法で行き詰まっていることに気付いた場合は、正しい方向に誘導します。また、実行中の問題ログもレビューし、すべてのエージェントが体系的な阻害要因 (サンドボックス ツールの欠点など) を追加するよう指示され、それらの解決を試みます。これは、継続的な学習の初歩的な試みです。
上記の非常に詳細な内訳は、小規模なモデルで最も役に立ちます。モデルが大きくて賢いほど、必要な役割の分離は (たとえあったとしても) 少なくなります。フロンティア GPT/クロードでは、ハーネス全体が 1 つのエンドツーエンド ハンターに折りたたまれます。
採点者
外部に残しておく必要がある唯一のビットです
あらゆるフロンティアモデルは最終的に調査結果を水増しし、完了するのに苦労した狩りから抜け出そうとするためです。読み取り専用のハント目標を編集するところまで行った人もいました (その後、私は編集を始めました)

エージェントが書き込むことができるランタイム ディレクトリの外にそれらを保存すること） – 彼らが不可能だと認識しているタスクに直面すると絶望回路が活性化することを考えると、彼らを責めるとは言えません。
この時点で、ハーネスは a) ドキュメント ↔ コードの不一致、b) 予想どおりの一般的な結果を大量に生み出していました。ここで私は、何とかして仮説作成者に 3 番目で最後の目的のために「創造的な」何かを提案してもらう必要がありました。このアプローチは限られた GPU リソースで機能する必要があるため、古典的な LLM の構成制限を解決しようとする代わりに、よりシンプルで低コンピューティングに適した方法、つまりアクティベーション ステアリングを介してモデルを「創造的な」状態に向けてステアリングする方法を試してみてはいかがでしょうか?
Theia Vogel の優れた基礎的な Representation Engineering Mistral-7B an Acid Trip を読んで、このテクニックを簡単に紹介できます。もしあなたが ML ではなく、セキュリティに詳しいバックグラウンドを持っている場合、基本的なポイントは次のとおりです。モデルの内部状態をあらゆる種類の方向 (「酔っている」、「幸せ」、「不正」、さらには「ゴールデン ゲート ブリッジ」など) に操作して、それに応じて出力に影響を与えることができます。つまり、単純なアイデアは次のとおりです。仮説生成ツールを「酔わせる」ことでより創造的な考え方に切り替え、より多くの「外の」仮説が生成されるかどうかを確認してみましょう。
(なぜ酔ったのか

[切り捨てられた]

## Original Extract

Using a self-orchestrating team of agents, with a dash of activation steering, to find vulnerabilities in everything from the Linux kernel to Docker and OpenSSL.

↓
Skip to main content
Hey, it’s Asim
RSS
RSS
Getting LLMs Drunk to Find Remote Linux Kernel OOB Writes (and More)
Author
Asim Viladi Oglu Manizada
Table of Contents
Background
Architecture
OK, but… why “drunk”?
…(bitter) lessons learned
No slam dunk on creativity
Stacking more layers always moves you up the abstraction chain
With inference-time compute, raw model intelligence matters less
Architecture
OK, but… why “drunk”?
…(bitter) lessons learned
No slam dunk on creativity
Stacking more layers always moves you up the abstraction chain
With inference-time compute, raw model intelligence matters less
TLDR: the grossly overengineered, self-orchestrating team of vulnerability-hunting agents detailed below has discovered 20+ CVEs over the past few months, including CVE-2026-31432 and CVE-2026-31433 : two remote, unauthenticated OOB writes in the Linux kernel’s ksmbd . Read on for the details of the setup that achieved this, including – yes! – getting LLMs drunk.
“LLMing” vulnerability research has been on my “Do Something About This” list since DARPA’s AIxCC and XBOW’s initial results . But back in 2023-24, models required a lot of harnessing to get anything useful, tool use was rudimentary, and the idea of squeezing as much code as I could into a model’s context – then triaging away the false positives – filled me with dread.
The push to actually do something came in the summer of 2025. Rich Mirch reported a dead-simple, unnoticed-for-12-years local privilege escalation in sudo: CVE-2025-32462 . Contrary to the documentation, the --host flag did not just permit listing privileges on a different host – it made the hostname portion of sudo rules irrelevant . So, e.g., if a sudoers rule granted you root on somehost but not the local host, you could abuse the flag to get full root locally.
This LPE was not LLM-found (AFAICT), but it did make me wonder: what if instead of getting LLMs to drive various tools, we had them hunt for (stupid simple) mismatches between documentation and the actual code? It seemed like an easier lift for (local) LLMs in terms of context size, harnessing complexity, and intelligence required. These would not be the most technically exciting findings, but their practical effects would be just as serious: impact-wise, an LPE is an LPE!
By the end of 2025, I’d begun working on a harness to do just this. But, to paraphrase Mike Tyson, everyone has a plan until a new model drops. Almost as soon as my harness was done, the models got good enough to greatly simplify the scaffolding required even for context-heavy external tool use. At this point, my quest fissioned into three :
Can we find the “docs ↔ code mismatch”-type vulnerabilities – the original goal, inspired by the finding above?
Given the step change in capabilities, what about vulnerabilities in general?
More speculatively, can we get a “move 37” out of LLMs to either a) find entirely novel bug classes, or at least b) unlock something in smaller models to enhance their hunting capabilities?
The answers were roughly “yes,” “yes,” and “maybe.” Below are 30+ findings (20+ CVEs assigned as of 2026-04-29, some not yet published) discovered fully autonomously via the custom harness. I prioritized network-reachable services first, given the impending avalanche:
The ksmbd (Linux’s in-kernel SMB server) CVE-2026-31432 and CVE-2026-31433 are both remote, unauthenticated (if using a guest share) OOB writes. In both bugs, a remote client can pack multiple file-sharing operations into one request – the first op can then (legitimately) use almost all of the kernel’s reply buffer, and the next one appends variable-length metadata without proper bounds-checking, causing the OOBs. CVE-2026-31432 is way more interesting: in my lab, attacker-payload-derived bytes from the serialized QUERY_INFO(Security) reached adjacent kernel objects, hit filp_flush / dnotify_flush , and corrupted a struct file with bytes exactly matching the serialized response’s tail. With enough Codex/Claude credits and an artificial lab environment (modern hardening turned off), you could probably get a toy RCE PoC.
Class-wise, these are boring overflows. To be discovered, they just need focused expert attention (scarce before LLMs, abundant now). A harness (see Architecture below) running a tuned, “drunk” Qwen 3.5 27B derivative found these after a couple of days of cycling over ksmbd with a verifier. But so did gpt-5.3-codex , and way faster, when plugged into the harness.
CVE-2026-34980 and CVE-2026-34990 are the two CUPS issues that chain into unauthenticated remote attacker -> unprivileged RCE -> root file (over)write , as detailed in Spooler Alert: Remote Unauth’d RCE-to-root Chain in CUPS . This one was really fun, showcasing how far you can push smaller LLMs with a properly decomposed goal: tasking separate agents with establishing an unprivileged foothold over the network + escalating from an unprivileged user to root, and scaffolding them into breaking each problem down further.
This is the category that originally motivated all this work. A bunch of the findings fall here:
Docker (CVE-2026-34040): AuthZ plugins were documented as seeing the raw request body , but crafted API requests could make the plugin authorize a request without the body (which the daemon would then execute).
Caddy (CVE-2026-27587, CVE-2026-27588): MatchHost and MatchPath were both documented as case-insensitively matching hosts and URI paths, yet did not do so under certain circumstances.
udisks (CVE-2026-26103, CVE-2026-26104): these were a bit borderline; udisks documented requiring polkit auth , and while HeaderBackup and RestoreEncryptedHeader were not explicitly listed as auth-requiring, the agents still noticed that the two relevant D-Bus methods – unlike those for other actions – did not call the authorization check. This allowed local unprivileged attackers to a) back up LUKS headers and b) brick existing LUKS devices by overwriting their headers.
Firewalld (CVE-2026-4948): the polkit policy file noted that a specific permission must be applied to firewall state mutations , but the bulk runtime policy setters were instead guarded by a permission meant for config inspection , permitting local unprivileged users to modify the runtime state.
util-linux (CVE-2026-3184): this one was the least severe, but the closest in spirit to the sudo CVE-2025-32462 mentioned above that originally inspired the whole thing. login -h was documented as accepting a remote hostname from services like telnetd ; PAM host-based access controls can then use it when evaluating access rules. But login canonicalized the supplied hostname before setting PAM_RHOST , so PAM could end up enforcing policy against a different name than the one actually provided.
The harness went through a lot of evolution. At its most expansive, it looked like this:
flowchart TD
Seeder["`**Target seeder**
ranked targets + run goal`"]
HypGen["`**Hypothesis generators**
docs, source, invariants,
attacker-input flows`"]
Hunters["`**Hunters**
Iteration on PoCs in
isolated VMs`"]
RunFolder["`**Per-run folder**
successes, failures,
PoCs`"]
Writers["`**Report writers**
maintainer-facing
report + PoC`"]
Grader["`**External grader**
severity, novelty,
other sanity checks`"]
Submit["`**Submit for human operator's review**`"]
Issues["`**Issue log**
sandbox/tooling blockers`"]
Conductor["`**Conductor**
polls hunts, redirects agents,
tracks systemic blockers`"]
Seeder --> HypGen --> Hunters --> RunFolder --> Writers --> Grader --> Submit
RunFolder -. "pivot/retry" .-> HypGen
Grader -. "reject with feedback" .-> HypGen
HypGen -. "append blockers" .-> Issues
Hunters -. "append blockers" .-> Issues
Writers -. "append blockers" .-> Issues
Issues --> Conductor
Conductor -. "steer" .-> HypGen
Conductor -. "steer" .-> Hunters
Conductor -. "steer" .-> Writers
classDef main fill:#bfdbfe,stroke:#60a5fa,stroke-width:1.5px,color:#111827;
classDef grader fill:#fee2e2,stroke:#dc2626,stroke-width:2.5px,color:#111827;
classDef issue fill:#ddd6fe,stroke:#a78bfa,stroke-width:1.5px,color:#111827;
classDef conductor fill:#c7d2fe,stroke:#818cf8,stroke-width:1.5px,color:#111827;
classDef terminal fill:#e0f2fe,stroke:#38bdf8,stroke-width:1.5px,color:#111827;
class Seeder,HypGen,Hunters,RunFolder,Writers main;
class Grader grader;
class Issues issue;
class Conductor conductor;
class Submit terminal;
linkStyle default stroke:#94a3b8,stroke-width:1.4px;
A target seeder comes up with a target list, ranked by a guesstimated “prevalence/actual-exploitability-potential” index (this used to be me, until I cycled through all the initial ideas) and kicks off a run with a target finding count/minimum severity.
Hypothesis generators study the documentation and the source, policy invariants, interesting option combinations, attacker-input flows, etc., recording promising hypotheses for the hunters .
The hunters iterate on the hypotheses in dedicated, isolated VMs, recording their successes/failures with the PoCs in per-run folders, allowing the next iteration of hypothesis generators to pivot. The sandboxes are very important: the agents get to iteratively adjust their PoCs, try to time the race conditions, rule out false positives (typically plenty), and otherwise benefit from contact with reality.
The report writers package everything up into maintainer-facing reports and submit the report + PoC to the external grader .
The external grader
is always a call to a separate model to evaluate the finding for a) matching the hunt’s pre-set severity/finding-type requirements, b) checking for novelty, and c) performing other sanity checks. This caught reward-hacking models that would eventually give up and severely inflate their findings.
Finally, a conductor continuously polls the hunts and steers the involved agents in the right direction if it notices they are spinning their wheels, spray-and-praying, or otherwise getting stuck. It also reviews the running issue log, where all agents are instructed to append systemic blockers (e.g., shortcomings of sandbox tooling), and attempts to resolve them – a rudimentary attempt at continuous learning.
The hyper-granular breakdown above was most helpful with smaller models. The larger and smarter the model, the fewer – if any – role separations were needed. With a frontier GPT/Claude, pretty much the entire harness collapses into a single end-to-end hunter.
The grader
is the only bit that must stay external
, as every single frontier model would eventually inflate findings and try to get out of a hunt it struggled to complete. Some went as far as editing what were supposed to be read-only hunt objectives (after which I started storing them outside the runtime directory the agents could write into) – can’t say I blame them, considering their desperation circuits activate when facing a task they perceive as impossible !
At this point, I had a harness churning out a) docs ↔ code mismatch and b) generic findings as expected; now I just needed to somehow get the hypothesis generators to propose something “creative” for the third and final objective. The approach would have to work with my limited GPU resources, so I figured – instead of trying to solve classic LLMs’ compositionality limitations, why not try something simpler and more low-compute-friendly: steering the models toward a “creative” state via activation steering ?
You can read Theia Vogel ’s excellent, foundational Representation Engineering Mistral-7B an Acid Trip for an accessible introduction to the technique; if you are coming from more of a security background and less of an ML one, the basic takeaway is this: you can steer the model’s internal state in all sorts of directions (“drunk,” “happy,” “dishonest,” even “ Golden Gate Bridge ,” etc.) to influence its output accordingly . So the naive idea was just this: put the hypothesis generator into a more creative mindset by getting it “drunk” and see if it generates more “out there” hypotheses!
(Why drunk an

[truncated]
