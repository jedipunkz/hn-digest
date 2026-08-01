---
source: "https://cryptonews.net/news/security/33234551/"
hn_url: "https://news.ycombinator.com/item?id=49139030"
title: "Was AI Responsible for Finding the Coldcard Security Flaw?"
article_title: "Was AI Responsible for Finding the Coldcard Security Flaw?"
author: "paulpauper"
captured_at: "2026-08-01T22:44:45Z"
capture_tool: "hn-digest"
hn_id: 49139030
score: 1
comments: 1
posted_at: "2026-08-01T22:17:17Z"
tags:
  - hacker-news
  - translated
---

# Was AI Responsible for Finding the Coldcard Security Flaw?

- HN: [49139030](https://news.ycombinator.com/item?id=49139030)
- Source: [cryptonews.net](https://cryptonews.net/news/security/33234551/)
- Score: 1
- Comments: 1
- Posted: 2026-08-01T22:17:17Z

## Translation

タイトル: AI はコールドカードのセキュリティ欠陥を発見する責任がありましたか?
説明: 組織的な掃討により数百のウォレットが空になる この事件は、約 500 の単一署名ビットコインから約 594 ドル BTC (当時の価値で約 3,800 万ドル相当) が移動された後に公になりました。

記事本文:
AI はコールドカードのセキュリティ欠陥を発見する責任を負っていましたか?
最新ニュース
ビデオ
ビットコイン
DeFi
NFT
イーサリアム
アルトコイン
ブロックチェーン
マイニング
金融
メタバース
法的
セキュリティ
分析
交換
その他
ゲームファイ
ICO
ニュース
AI はコールドカードのセキュリティ欠陥を発見する責任を負っていましたか?
連携した掃除で何百もの財布が空になる
この事件は、7月30日に約594ドルBTC（当時3,800万ドル近く相当）が約500の単一署名ビットコインアドレスから移動した後に公になった。このニュースが最初に報道されたとき、Bitcoin.comニュースは、送金は約25分以内に行われ、共通の技術的弱点を持つウォレットを標的にしているようだと指摘した。
その後のブロックチェーンの見直しにより、盗難の可能性のある規模が拡大しました。研究者らは、約 41 分間に 1,082 ～ 1,196 のアドレスが影響を受けた可能性があると推定しています。その後、コールドカード スイープ ウォッチと呼ばれるカスタム ダッシュボードでは、合計が 1,128.4717 $BTC と表示され、ビットコインが 63,044 ドル付近で取引されていたとき、約 7,110 万ドルに相当します。
資金のほとんどは数百ビットコインを保持するアドレスに統合され、その大部分はほとんど固定されたままでした。影響を受けたアドレスは、1 つの重要な詳細によって関連付けられていました。それは、それらの回復シードがカナダの会社 Coinkite によって製造された Coldcard ハードウェア ウォレット上に作成されていたということです。
リカバリ シードは、暗号通貨ウォレットへのアクセスを制御する単語のリストです。そのシードを再構築または入手できる人は、通常、物理デバイスを所有せずにウォレットの資金を移動できます。
Coldcard が壊れたランダム性システムを発見
Coinkite は、Coldcard デバイスで生成された特定のシードが弱い可能性があるという緊急勧告を発行しました。 2021 年 3 月頃にリリースされたファームウェア バージョン 4.0.1 以降のバージョンを実行している Mk3 デバイスは、最も大きなリスクに直面しているデバイスの 1 つです。
毛皮

この分析により、Coinkite が緊急ファームウェア修正をリリースする前に、一部の Mk4、Mk5、および Q デバイスで作成されたシードに懸念が広がりました。 Tapsigner、Opendime、Satscard 製品は、異なるソフトウェアを使用しているため影響を受けなかったと報告されています。
この欠陥には、ランダム データの生成に使用されるプロセスが関係していました。安全なウォレットは高品質のランダム性に依存しているため、回復シードを推測することはできません。最も深刻な影響を受けた Mk3 デバイスでは、シードに意図された 128 ビットではなく、有効なランダム性が約 40 ビットしか含まれていなかった可能性があると研究者は推定しました。
その違いは重要です。適切に生成された 128 ビット シードを総当たりで推測することは事実上不可能であると考えられています。 40 ビット シードでは可能性が大幅に減少するため、十分な計算能力を持つ攻撃者は潜在的なシードをオフラインでテストし、結果のアドレスをパブリック ビットコイン ブロックチェーンと比較することができます。
一部の新しいデバイスでは、安全なハードウェアによって予測不可能なデータの層が追加されたため、約 72 ビットの有効なランダム性が提供されている可能性があります。そうなるとシードの再構築が難しくなりますが、それでも意図したよりもはるかに弱いです。
1 つの構成エラーは 5 年間存続
この問題は、同様のジョブを実行する 2 つのソフトウェア機能に関連するビルド時の構成ミスから始まりました。 1 つの機能はデバイスのハードウェアベースの真の乱数生成器を使用し、もう 1 つは MicroPython から継承した弱いソフトウェア プロセスに依存していました。
Coinkite は、MicroPython オプションを無効にすることを目的としていました。ただし、ソフトウェア チェックでは、構成ラベルの値がゼロに設定されているかどうかではなく、構成ラベルが定義されているかどうかだけが調べられました。その結果、完成したファームウェアは、より弱い機能を黙って選択する可能性があります。
2 つの関数の形式が一致していたため、

ソフトウェアは明らかなエラーを生成することなくコンパイルと実行を続けました。この間違いは 2021 年のソフトウェア移行の頃にコードに侵入し、公開されているファームウェアに 5 年以上残存していました。
現在、デバイスを更新しても、欠陥のあるソフトウェアの下で生成されたシードは強化されません。影響を受けるユーザーは、修正されたファームウェアまたは別の安全なデバイスを使用して完全に新しいシードを作成し、その新しいシードによって管理されるアドレスに資金を移動する必要があります。
シードの作成時に少なくとも 50 個の独立したサイコロの目を追加したユーザーは、弱点を回避するのに十分な追加のランダム性を提供している可能性があります。強力な BIP-39 パスフレーズも再構築をより困難にした可能性がありますが、複数の独立したデバイスからの署名を必要とするウォレットにより、侵害された 1 つのシードが単独で資金を移動することができなかった可能性があります。
CoinkiteはAIを指摘しているが、証拠は依然として欠けている
CoinkiteのCEO、Rodolfo Novak氏は公式に謝罪し、ファームウェアの障害については同社が全責任を負っていると述べた。同氏は、チームがソフトウェアの修正、技術レポート、影響を受けるユーザーへのサポートに取り組んでいると述べた。
Coinkite と Novak は、欠陥がどのように発見されたかについても驚くべき理論を展開しました。同社のファームウェアは何年も前から公開されていたため、誰かがAIを使って古いバージョンのコードを調べ、弱いランダム性のパスを特定した可能性があると同社は述べた。
「他のすべての開発者へ: 私たちは、これが新しい AI パラダイムの厳然たる現実であると信じています。AI 支援によるコード レビューは、業界で最も熟練した専門家をも上回るスピードで潜在的なバグを発見できるようになりました。」と Novak 氏は X に掲載された謝罪投稿の中で書きました。「ファームウェアがオープンソースであるか、以前に公開されていた場合は、すでに攻撃者と防御者の両方に読み取られていると想定してください。」
今日の最新の AI コーディング システムは、

大規模なソフトウェア リポジトリを分析し、構成設定、機能、セキュリティ上の前提条件の間の疑わしい関係を特定します。攻撃者は、そのようなシステムに対して、特に弱い乱数生成器、フォールバック関数、または暗号キーに影響を与えるエラーを探すよう要求する可能性があります。
その後、独立した研究者が、根本的なランダム性の問題が判明した後、AI モデルを使用して問題を特定または説明することを報告しました。これは、AI 支援によるコード分析がいかに利用しやすくなったかを示しましたが、元の攻撃者が AI を使用していたことは証明されませんでした。
Coinkite は、主要な AI モデルを使用した独自のレビューで盗難前の欠陥を発見できなかったことを認めました。この結果は、AI システムがすべての重大な欠陥を自動的に検出するわけではないことを示しています。それらのパフォーマンスは、受け取る指示、提供されるコードの量、人間のレビュー担当者が警告サインを理解するかどうかによって決まります。
批評家は人間の失敗が先だと言う
セキュリティ専門家の中には、人工知能 (AI) に重点を置きすぎると、基本的なエンジニアリングの失敗から注意が逸れてしまう危険があると主張する人もいます。多くの人は、構成ミスは既知の種類のソフトウェア エラーであり、シード生成を中心とした従来のコード レビュー、テスト手順、または監査があれば、何年も前に検出できた可能性があると考えています。
対立する意見は必ずしも両立しないわけではありません。人的ミスによって脆弱性が生じ、それが存続する一方、AI によって脆弱性の発見、理解、悪用のコストが削減された可能性があります。防御側は危険な弱点をすべて特定する必要がありますが、攻撃側は 1 つだけを特定する必要があります。
この事件は、オープンソースのセキュリティに関する前提にも疑問を投げかけます。公開コードにより、独立した専門家がソフトウェアを検査することができますが、可用性だけでは誰かが正しいセクションをレビューすることは保証されません。

微妙な欠陥を認識し、攻撃者が行動を起こす前に報告します。
Coldcard ユーザーにとって、当面の優先事項は、シードがいつどのように作成されたかを判断することです。影響を受けるシードを所有している人は、公式 Coinkite チャネルを通じて指示を確認し、修正されたファームウェアをインストールし、新しいシードを作成し、フィッシングの試みや偽のサポート メッセージに注意しながら慎重に資金を移動する必要があります。
長期的な問題は、ビットコインがどれだけ盗まれたか、捜査官が攻撃者を特定できるか、欠陥の発見にAIが決定的な役割を果たしたかどうかなどが中心となるだろう。ハードウェアウォレット企業は、エントロピーテストを強化し、ビルド構成を監査し、人間の専門家と敵対的AIツールの両方を使って古いコードを継続的に検査するというプレッシャーにも直面するだろう。

## Original Extract

A Coordinated Sweep Empties Hundreds of Wallets The incident became public after roughly 594 $BTC, valued near $38 million at the time, moved from about 500 single-signature bitcoin

Was AI Responsible for Finding the Coldcard Security Flaw?
Latest news
Video
Bitcoin
DeFi
NFT
Ethereum
Altcoins
Blockchain
Mining
Finance
Metaverse
Legal
Security
Analytics
Exchange
Other
GameFi
ICO
News
Was AI Responsible for Finding the Coldcard Security Flaw?
A Coordinated Sweep Empties Hundreds of Wallets
The incident became public after roughly 594 $BTC , valued near $38 million at the time, moved from about 500 single-signature bitcoin addresses on July 30. When the news first broke, Bitcoin.com News noted that the transfers occurred within approximately 25 minutes and appeared to target wallets with a shared technical weakness.
Later blockchain reviews expanded the possible scale of the theft. Researchers estimated that between 1,082 and 1,196 addresses may have been affected during a period of about 41 minutes. A custom dashboard called Coldcard Sweep Watch subsequently placed the total at 1,128.4717 $BTC , worth about $71.1 million when bitcoin traded near $63,044.
Most of the funds were consolidated into an address holding hundreds of bitcoin, where a large portion remained largely stationary. The affected addresses were linked by one important detail: Their recovery seeds had been created on Coldcard hardware wallets manufactured by Canadian company Coinkite.
A recovery seed is a list of words that controls access to a cryptocurrency wallet. Anyone who can reconstruct or obtain that seed can usually move the wallet’s funds without possessing the physical device.
Coldcard Finds a Broken Randomness System
Coinkite issued an urgent advisory warning that certain seeds generated on Coldcard devices could be weak. Mk3 devices running firmware version 4.0.1, released around March 2021, and later versions were among those facing the greatest risk.
Further analysis widened the concern to seeds created on some Mk4, Mk5, and Q devices before Coinkite released emergency firmware fixes. Tapsigner, Opendime, and Satscard products were reportedly not affected because they use different software.
The flaw involved the process used to generate random data. Secure wallets depend on high-quality randomness so that their recovery seeds cannot be guessed. On the most seriously affected Mk3 devices, researchers estimated that the seed may have contained only about 40 bits of effective randomness instead of the intended 128 bits.
That difference is critical. A properly generated 128-bit seed is considered practically impossible to guess through brute force. A 40-bit seed offers dramatically fewer possibilities, allowing an attacker with enough computing power to test potential seeds offline and compare the resulting addresses with the public Bitcoin blockchain.
Some newer devices may have provided approximately 72 bits of effective randomness because secure hardware added another layer of unpredictable data. That would make the seeds harder to reconstruct, though still much weaker than intended.
One Configuration Error Survived for 5 Years
The problem began with a build-time configuration mistake involving two software functions that performed similar jobs. One function used the device’s hardware-based true random number generator, while the other relied on a weaker software process inherited from MicroPython.
Coinkite intended to disable the MicroPython option. However, a software check looked only at whether a configuration label had been defined, rather than whether its value had been set to zero. As a result, the finished firmware could silently select the weaker function.
Because the two functions had matching formats, the software continued to compile and run without producing an obvious error. The mistake entered the code around a 2021 software migration and remained in publicly available firmware for more than five years.
Updating a device now does not strengthen a seed that was generated under the faulty software. Affected users must create a completely new seed using corrected firmware or another secure device, then move their funds to addresses controlled by that new seed.
Users who added at least 50 independent dice rolls while creating their seed may have supplied enough extra randomness to avoid the weakness. A strong BIP-39 passphrase could also have made reconstruction more difficult, while wallets requiring signatures from several independent devices may have prevented one compromised seed from moving funds alone.
Coinkite Points to AI, but Proof Remains Missing
Coinkite CEO Rodolfo Novak apologized publicly and said the company took full responsibility for the firmware failure. He said the team was working on fixed software, technical reports, and support for affected users.
Coinkite and Novak also advanced a striking theory about how the flaw was discovered. Because its firmware had been publicly available for years, the company said it believed someone may have used AI to examine older versions of the code and locate the weak randomness path.
“To every other developer: we believe this is a sober reality of the new AI paradigm. AI-assisted code review can now find latent bugs at a speed that is outpacing even the industry’s most seasoned experts,” Novak wrote in his apology post published on X. “If your firmware is open-source or has ever been public, assume it’s already being read by attackers and defenders alike.”
Today’s modern AI coding systems can process large software repositories and identify suspicious relationships between configuration settings, functions, and security assumptions. An attacker could ask such a system to look specifically for weak random number generators, fallback functions, or errors affecting cryptographic keys.
Independent researchers later reported using AI models to locate or explain the issue after the underlying randomness problem was known. That demonstrated how accessible AI-assisted code analysis has become, but it did not establish that the original attacker used AI.
Coinkite acknowledged that its own review using a leading AI model failed to uncover the flaw before the theft. That result shows that AI systems do not automatically find every serious defect. Their performance can depend on the instructions they receive, the amount of code supplied and whether a human reviewer understands the warning signs.
Critics Say the Human Failure Came First
Some security specialists argue that focusing too heavily on artificial intelligence (AI) risks distracting from a basic engineering failure. Many believe the configuration mistake was a known type of software error, and conventional code reviews, testing procedures, or audits centered on seed generation could have detected it years earlier.
The opposing views are not necessarily incompatible. A human error created the vulnerability and allowed it to persist, while AI may have lowered the cost of finding, understanding, or exploiting it. Defenders must identify every dangerous weakness, while an attacker needs to locate only one.
The incident also challenges assumptions about open-source security. Public code allows independent experts to inspect software, but availability alone does not guarantee that someone will review the correct section, recognize a subtle defect, and report it before an attacker acts.
For Coldcard users, the immediate priority is determining when and how their seed was created. Anyone with an affected seed must verify instructions through official Coinkite channels, install corrected firmware, create a new seed, and move funds carefully while watching for phishing attempts and fake support messages.
The longer-term questions will center on how much bitcoin was taken, whether investigators can identify the attacker, and whether AI played any decisive role in finding the flaw. Hardware wallet companies will also face pressure to strengthen entropy testing, audit build configurations, and continuously examine old code with both human experts and adversarial AI tools.
