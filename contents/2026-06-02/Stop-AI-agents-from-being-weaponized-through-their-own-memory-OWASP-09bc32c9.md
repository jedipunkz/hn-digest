---
source: "https://www.helpnetsecurity.com/2026/06/01/owasp-agent-memory-guard/"
hn_url: "https://news.ycombinator.com/item?id=48360116"
title: "Stop AI agents from being weaponized through their own memory (OWASP)"
article_title: "OWASP Agent Memory Guard: Stop AI agents from being weaponized through their own memory - Help Net Security"
author: "vgudur297"
captured_at: "2026-06-02T00:35:05Z"
capture_tool: "hn-digest"
hn_id: 48360116
score: 2
comments: 0
posted_at: "2026-06-01T17:41:39Z"
tags:
  - hacker-news
  - translated
---

# Stop AI agents from being weaponized through their own memory (OWASP)

- HN: [48360116](https://news.ycombinator.com/item?id=48360116)
- Source: [www.helpnetsecurity.com](https://www.helpnetsecurity.com/2026/06/01/owasp-agent-memory-guard/)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T17:41:39Z

## Translation

タイトル: AI エージェントが自身の記憶を通じて兵器化されるのを阻止する (OWASP)
記事のタイトル: OWASP Agent Memory Guard: AI エージェントが自身の記憶を通じて武器化されるのを阻止する - Help Net Security
説明: AI エージェントのメモリ セキュリティにランタイム層が追加されました。OWASP の Agent Memory Guard はすべての書き込みをスクリーニングし、59 μs のレイテンシーで 92.5% のリコール率を達成しました。

記事本文:
OWASP Agent Memory Guard: AI エージェントが自身の記憶を通じて武器化されるのを阻止する - 公式ヘルプ Net Security
Help Net Security ニュースレター : 毎日および毎週のニュース、サイバーセキュリティの仕事、オープンソース プロジェクト、最新ニュース - ここから購読してください!
OWASP Agent Memory Guard: AI エージェントが自身の記憶を通じて武器化されるのを阻止します。
AI エージェントはセッション間でメモリを保持します。会話履歴、ベクトル ストア、スクラッチパッド、および RAG インデックスは実行間で保持され、そのストアに書き込まれたものはすべて、エージェントが後で読み取る特権入力になります。攻撃者がテキストを間違ったフィールドに埋め込むと、エージェントの指示を無効にしたり、ユーザー データを引き出したり、将来のツール呼び出しを操作したりすることができます。また、メモリが存在するため、その効果はセッションをまたいで存続します。
Agent Memory Guard は、エージェントとそのメモリ ストアの間に位置するオープンソースのランタイム防御層で、検出器のパイプラインと YAML ポリシーを通じてすべての読み取りと書き込みをスクリーニングします。このプロジェクトは、ASI06、メモリ ポイズニングの OWASP リファレンス実装であり、エージェント アプリケーションの OWASP トップ 10 の 1 つのエントリです。
ガードは 5 つのコア検出カテゴリを実行します。 SHA-256 ベースラインは、不変キーによる帯域外改ざんをフラグします。内蔵の検出器は、プロンプト インジェクション マーカー、秘密および PII の漏洩、保護されたキーの変更、サイズの異常を検出します。 YAML ポリシーは、各検出結果をアクション (許可、編集、隔離、ブロック) にマッピングします。すべての決定により構造化された SecurityEvent が生成され、オペレーターは特定時点のスナップショットによってメモリを既知の正常な状態にロールバックできます。ドロップイン チャット履歴クラスは LangChain をカバーし、ミドルウェア パッケージはモデル入力、モデル出力、およびツール出力をスクリーニングします。
このベンチマークでは、5 つの検出器を通じて 55 のテスト ケース (4 つのカテゴリにわたる 40 の攻撃ペイロードと 15 の無害なサンプル) が実行されます。リコールカム

受信率は 92.5%、精度は 100%、偽陽性率はゼロで、待ち時間の中央値は 59 マイクロ秒でした。プロンプトインジェクションと保護されたキーの改ざんはそれぞれ 100% のスコアを獲得しました。機密データの漏洩は 83% に達し、サイズ異常は 80% に達しました。混同行列には、37 件の真陽性、3 件の偽陰性、および 0 件の偽陽性が記録されます。
「欠落したペイロードは両方とも、長さが固定長の正規表現パターンをわずかに超える API トークンです」と、プロジェクト作成者で OWASP プロジェクト リーダーの Vaishnavi Gudur 氏が機密データ カテゴリについて Help Net Security に語った。 1 つは、検出器が 36 を予期する ghp_ プレフィックスの後に 37 文字を持つ GitHub 個人アクセス トークンであり、もう 1 つは、検出器が 35 を予期する AIza プレフィックスの後に 38 文字を持つ Google API キーでした。
漏洩検出機能は固定長の量指定子を使用します。これは精度を優先し、ランダムな英数字文字列の誤検出をカットする意図的な選択ですが、プロバイダーがトークン形式を拡張すると陳腐化するという代償を伴います。 3 番目のミスは、58,913 バイトにシリアル化されるネストされた JSON 構造で、64 KB のしきい値をわずかに下回っていました。キーの以前の値と比較して 10 倍に増加したかどうかを 2 回目のチェックで確認すると、本番環境でキーが検出されます。ベンチマークは、以前の状態を持たないフレッシュ ガードで各テストを実行します。 Gudur 氏は、高リコール正規表現バリアントと適応しきい値キャリブレーションが v0.3.0 で予定されていると述べました。
オープンソース コードと目に見える YAML ポリシーにより、攻撃者はルールを読み取ることができます。 「現在のルールベースの検出器は第 1 層です」とグドゥル氏は述べ、より高度な脅威モデルを備えたチームがオープンソース層の上に追加の検出を重ねる多層防御設計について説明しました。保護キーのチェックはキー パス上で動作するため、ルールを知っていればバイパスできず、SHA-256 の整合性により、変更された不変のオブジェクトに対して決定的な不一致が生成されます。

値。 Base64 によるエンコード、文字分割、またはホモグリフにより、マッチング前の正規化が欠けている検出器を回避できるため、機密データのマッチングはより危険にさらされます。
適応回避テストが計画されています。 AgentThreatBench は、inspect_evals フレームワークに統合され、公開されたルールの知識に基づいて構築された回避対応ペイロード セットを追加します。防御面では、v0.4.0 ではセマンティック機能に ML ベースの異常検出が追加され、v0.3.0 ではチームがオープン YAML から除外できるカスタム検出器用のプラグイン インターフェイスが追加されています。
「GitHub Copilot はボイラープレートとスキャフォールディングに使用されました」と Gudur 氏は述べ、テスト設定、CI/CD 構成、pyproject.toml ファイル、およびプロバイダーのドキュメント、README セクションとドキュメント文字列に照らして検証されたドラフト正規表現パターンを引用しました。
検出器パイプライン アーキテクチャ、ポリシー エンジンの分離、MemoryStore プロトコル、スナップショットとロールバック メカニズム、およびソースクラス来歴システムは、OWASP ASI06 脅威モデルに対して人間が設計したものです。 40 のベンチマーク ペイロードは手作業で厳選されました。 Gudur 氏は、知的貢献は攻撃対象領域の特定、防御の設計、精選された敵対的コーパスに対する検証にあると述べ、定型的な標準的な実践に Copilot を使用するよう呼びかけました。
OWASP Agent Memory Guard は、GitHub で無料で入手できます。
予算を気にしない 25 のオープンソース サイバーセキュリティ ツール
GitHub CISO がセキュリティ戦略とオープンソース コミュニティとの協力について語る
Help Net Security の広告なしの月刊ニュースレターを購読して、重要なオープンソース サイバーセキュリティ ツールに関する最新情報を入手してください。ここから購読してください！
Windows Netlogon RCE が悪用され、ドメイン コントローラーが危険にさらされる (CVE-2026-41089)
NIST は国家脆弱性データベースの管理にどのように失敗したか
ハック

カーは Palo Alto GlobalProtect VPN 認証バイパスを悪用しています (CVE-2026-0257)
ダウンロード: AWS 上の AI ワークロードの安全な基盤
ダウンロード: 侵入テスト配信の自動化ガイド
CIS ベンチマーク 2026 年 3 月の更新
Windows Netlogon RCE が悪用され、ドメイン コントローラーが危険にさらされる (CVE-2026-41089)
NIST は国家脆弱性データベースの管理にどのように失敗したか
ハッカーがパロアルト GlobalProtect VPN 認証バイパスを悪用している (CVE-2026-0257)
NVIDIA、物理 AI エージェント ツールの大規模なバッチをオープンソースに移行
企業の不意を突くデータ発見のギャップ

## Original Extract

AI agent memory security gets a runtime layer: OWASP's Agent Memory Guard screens every write, hitting 92.5% recall at 59µs latency.

OWASP Agent Memory Guard: Stop AI agents from being weaponized through their own memory - Help Net Security
Help Net Security newsletters : Daily and weekly news, cybersecurity jobs, open source projects, breaking news – subscribe here!
OWASP Agent Memory Guard: Stop AI agents from being weaponized through their own memory
AI agents keep memory across sessions. Conversation history, vector stores, scratchpads, and RAG indexes persist between runs, and anything written into that store becomes a privileged input the agent reads back later. An attacker who plants text in the wrong field can override an agent’s instructions, pull out user data, or steer future tool calls, and the effect survives across sessions because the memory does.
Agent Memory Guard is an open-source runtime defense layer that sits between an agent and its memory store, screening every read and write through a pipeline of detectors and a YAML policy. The project is the OWASP reference implementation for ASI06, Memory Poisoning, one entry in the OWASP Top 10 for Agentic Applications.
The guard runs five core detection categories. SHA-256 baselines flag out-of-band tampering with immutable keys. Built-in detectors look for prompt injection markers, secret and PII leakage, protected-key modifications, and size anomalies. A YAML policy maps each finding to an action: allow, redact, quarantine, or block. Every decision emits a structured SecurityEvent, and point-in-time snapshots let an operator roll memory back to a known-good state. A drop-in chat history class covers LangChain, and a middleware package screens model inputs, model outputs, and tool outputs.
The benchmark runs 55 test cases through five detectors: 40 attack payloads across four categories and 15 benign samples. Recall came in at 92.5%, precision at 100%, and the false positive rate at zero, with median latency of 59 microseconds. Prompt injection and protected-key tampering each scored 100%. Sensitive data leakage reached 83%, and size anomaly reached 80%. The confusion matrix records 37 true positives, three false negatives, and zero false positives.
“Both missed payloads are API tokens whose length slightly exceeds the fixed-length regex pattern,” Vaishnavi Gudur , the project creator and OWASP project leader, told Help Net Security about the sensitive-data category. One was a GitHub personal access token with 37 characters after the ghp_ prefix where the detector expects 36, and the other a Google API key with 38 characters after the AIza prefix where it expects 35.
The leakage detector uses fixed-length quantifiers, a deliberate choice that favors precision and cuts false positives on random alphanumeric strings, at the cost of going stale when providers extend their token formats. The third miss was a nested JSON structure serializing to 58,913 bytes, sitting just under the 64KB threshold. A second check for tenfold growth against a key’s prior value would catch it in production. The benchmark runs each test on a fresh guard with no prior state. Gudur said higher-recall regex variants and adaptive threshold calibration are slated for v0.3.0.
Open-source code and a visible YAML policy let an attacker read the rules. “The current rule-based detectors are a first layer,” Gudur said, describing a defense-in-depth design where teams with higher threat models layer additional detection on top of the open-source layer. Protected-key checks operate on the key path, so knowing the rule gives no bypass, and SHA-256 integrity produces a deterministic mismatch on any altered immutable value. Sensitive-data matching is more exposed, since encoding through base64, character splitting, or homoglyphs can dodge a detector that lacks normalization before matching.
Adaptive evasion testing is planned. AgentThreatBench, now merged into the inspect_evals framework, will add an evasion-aware payload set built with knowledge of the published rules. On defense, v0.4.0 adds ML-based anomaly detection on semantic features, and v0.3.0 adds a plugin interface for custom detectors that teams can keep out of the open YAML.
“ GitHub Copilot was used for boilerplate and scaffolding,” Gudur said, citing test setup, CI/CD configuration, and the pyproject.toml file, along with draft regex patterns that were then validated against provider documentation, and README sections and docstrings.
The detector pipeline architecture, the policy-engine separation, the MemoryStore protocol, the snapshot and rollback mechanism, and the source-class provenance system were human-designed against the OWASP ASI06 threat model. The 40 benchmark payloads were curated by hand. Gudur said the intellectual contribution lies in identifying the attack surface, designing the defense, and validating it against a curated adversarial corpus, and called using Copilot for boilerplate standard practice.
OWASP Agent Memory Guard is available for free on GitHub .
25 open-source cybersecurity tools that don’t care about your budget
GitHub CISO on security strategy and collaborating with the open-source community
Subscribe to the Help Net Security ad-free monthly newsletter to stay informed on the essential open-source cybersecurity tools. Subscribe here!
Windows Netlogon RCE exploited, domain controllers at risk (CVE-2026-41089)
How NIST fumbled management of the National Vulnerability Database
Hackers are exploiting Palo Alto GlobalProtect VPN authentication bypass (CVE-2026-0257)
Download: Secure Foundations for AI Workloads on AWS
Download: Automating Pentest Delivery Guide
CIS Benchmarks March 2026 Update
Windows Netlogon RCE exploited, domain controllers at risk (CVE-2026-41089)
How NIST fumbled management of the National Vulnerability Database
Hackers are exploiting Palo Alto GlobalProtect VPN authentication bypass (CVE-2026-0257)
NVIDIA goes open source with a big batch of physical AI agent tools
Data discovery gaps that catch enterprises off guard
