---
source: "https://github.com/Gneiss-Group/Kessa"
hn_url: "https://news.ycombinator.com/item?id=49260224"
title: "Show HN: Kessa, a verifiable delegation and attestation system for AI agents"
article_title: "GitHub - Gneiss-Group/Kessa: Attestation and enforcement for delegated agent authority, with an independent offline verifier. · GitHub"
author: "raphtheb"
captured_at: "2026-08-11T16:48:11Z"
capture_tool: "hn-digest"
hn_id: 49260224
score: 1
comments: 0
posted_at: "2026-08-11T15:53:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Kessa, a verifiable delegation and attestation system for AI agents

- HN: [49260224](https://news.ycombinator.com/item?id=49260224)
- Source: [github.com](https://github.com/Gneiss-Group/Kessa)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T15:53:05Z

## Translation

タイトル: Show HN: Kessa、AI エージェント向けの検証可能な委任および認証システム
記事のタイトル: GitHub - Gneiss-Group/Kessa: 独立したオフライン検証者による、委任されたエージェント権限の認証と執行。 · GitHub
説明: 独立したオフライン検証者による、委任されたエージェント権限の認証と執行。 - 片麻岩グループ/ケッサ

記事本文:
GitHub - Gneiss-Group/Kessa: 独立したオフライン検証者による、委任されたエージェント権限の認証と執行。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
片麻岩グループ
/
ケッサ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
121 コミット 121 コミット .claude .claude .github .github ライセンス ライセンス Auditsink Auditsink build bu

ild cmd cmd docker docker docs docs 例 例 内部 内部パフォーマンス perf pkg/ タイプ pkg/ タイプ スクリプト スクリプト testdata testdata .dockerignore .dockerignore .gitignore .gitignore .gitleaks.toml .gitleaks.toml CHANGELOG.md CHANGELOG.md CLA-CORPORATE.md CLA-CORPORATE.md CLA.md CLA.md CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE LICENSING.md LICENSING.md Makefile Makefile NOTICE.md NOTICE.md PLUGIN_LICENSING.md PLUGIN_LICENSING.md README.md README.md REUSE.toml REUSE.toml SECURITY.md SECURITY.md UPCOMING.md UPCOMING.md go.mod go.mod すべてのファイルを表示 リポジトリ ファイルのナビゲーション
委任されたエージェント権限の証明と執行。人間の代理人は、
組織、組織からエージェント、エージェントからサブエージェント。各ホップが狭くなる
権限を拡大することはなく、結果として生じるすべてのアクションはログに記録されます。
独立したオフライン検証者が再チェックできる改ざん防止チェーン。
公開された DID 文書以外に私たちのものは何も信頼しません。
同じコマンドですが、実行間で 1 バイトが変更されました。ベリファイアは、改ざんされたエントリそのもので失敗し、ファイルのみからすべての判定を再導出します。 GIF がどのように構築および再生成されるか。
ジャンプ先: これで解決する内容 · 試してみる · ステータス · 標準 · 既知の制限
人が AI エージェントに仕事を渡し、そのエージェントがその一部を AI エージェントに渡す場合
もう一つは、そのチェーンを伝う権威は通常目に見えず、
通常は減っていません。 Kessa はチェーンを明確にし、狭くし、証明可能にします
事後。
エージェントは、タスクが必要とするよりもはるかに多くの権限を継承します。よくあるパターン
エージェントにユーザーのトークンのコピーを渡すことにより、要約ジョブが保持されます。
お金を動かす力。 Kessa は、各ホップに厳密に狭い新しいホップを発行します
資格情報、および自身の作成者の範囲を広げようとするホップ

次の場合に拒否されます
それは使用されるときではなく、鋳造されるのです。
エージェントが行動した時点で何をすることが許されていたのかは誰にもわかりません。
インシデント後にそれを再構築するということは、通常、アプリケーション ログを信頼することを意味します。
Kessa は、委任チェーン、ポリシー、および署名された承認を記録します。
行動に伴う証拠。
決定を行ったシステムは、その記録も書き込みます。つまり
間違っている最大の理由がある政党。 Kessa の検証者は、すべての
判決文を読むのではなく、証拠に基づいて判決を下す
ポイントが書き留められているため、手を抜いたプロキシはエクスポートに失敗します。
通常、検証とはベンダーを信頼することを意味します。 kessa verifyは別ですが、
ファイルに対してオフラインで実行される、寛容にライセンスされたバイナリ。 「いいえ」と表示されます
私たちのサービスであり、そのトラストルートは公開鍵のディレクトリです。
独自に取得します。
取り消しは、すでに委任された権限には及ばない。ケッサは毎回チェックします
アクション時に署名済みステータスリストを照合するため、資格情報を取り消します。
中間チェーンは、それに依存する結果的なアクションを停止します。
記録の改ざんは検出できません。監査ログはハッシュチェーン化されています
署名されており、署名には丸太の長さと先端が含まれているため、エントリは
事後に編集したり、並べ替えたり、静かに削除したりすることはできません。
ケッサは内容ではなく権威を統治します。アクションが許可されるかどうかを決定します
続行します。エージェントの発言を判断したり、承認されたデータを検査したりすることはありません
アクションタッチ。
3 つのバイナリが証拠を生成し、1 つがそれを再チェックして、どれも信頼しません。
3つ。委任、執行、および権限の図を含む完全なウォークスルー
検証段階に加えて、クリーンな評決が何を証明するのかについての正確な記述、
Kessa の仕組み にあります。
前提条件: Go 1.26 以降 ( go.mod のバージョン。
古いツールチェーンは 1 をダウンロードしようとします

.26、それができない場合は失敗します）、さらに make と
バッシュ。他にインストールするものはありません: Kessa にはサードパーティの Go がありません
依存関係があるため、デモをビルドし、ネットワークをオフにして実行します。の
コンテナー イメージにはさらに、実行中の Docker デーモンが必要です。
来歴チェックには GitHub CLI が必要です。
デモを作成 # 全体のストーリー: 7 つのシナリオすべて、エンドツーエンド、決定的
または、自分でピースを操作することもできます。
make build # すべてのバイナリを ./bin にビルドします (ベリファイアは ./bin/kessa です)
# 1. チェーンを発行し、公開アーティファクトを公開します。
./bin/kessa-issuer 発行 \
--spec の例/issuer/spec.json --keystore の例/issuer/keystore.json \
--root ./public --out ./private/chain.json
# 2. オフラインで、何も実行されていない状態で、それらに対するエクスポートを確認します。
./bin/kessa 検証 \
--export testdata/audit_export_v2.golden.json \
--dids ./public \
--status " https://localhost/orgs/acme/status.json=./public/localhost/orgs/acme/status.json "
# 3. 中間チェーン認証情報を取り消します。同じエクスポートでは検証が停止します。
./bin/kessa-issuer revoke --spec 例/issuer/spec.json \
--keystore 例/issuer/keystore.json --root ./public --index 42
終了 0 = すべてのエントリが持ち込まれた証拠に対して検証される、1 = FAIL または
整合性のみの v1 ダウングレード (クリーン パスではない)、2 = 使用量/IO エラー。
何も開始されず、何もダイヤルされません。検証者の唯一の入力はファイルです。
--fetch-dids は、パブリック Did:Web ドキュメントの HTTPS 解決を有効にします。
kessa はネットワーク アクセスを行うことができますが、デフォルトではオフになっています。
Kessa はエンドツーエンドで動作し、積極的に開発中です。 7つのシナリオ実行開始
終了 (デモの作成) し、複数の敵対的セキュリティ レビュー ラウンドが終了します
すべての発見事項が修正されました。すべてのラウンドは自動実行 AI レッドチーム パスであり、
第三者監査;外部監査はまだ委託されていません。それは
標準ライブラリのみ、

サードパーティへの依存関係はありません。
中心的な主張は 4 つあり、その 4 つすべてが構築されています。証明可能な減衰、
証明可能な監査可能性、独立した検証者、および完全な 7 つのシナリオのデモ。
まだ実稼働環境に対応していません。鍵の処理は署名者の背後で実行されます
ソフトウェア キーストア (デモ、CI、
--software-key パス (その秘密キーはファイル内に平文で存在します)、および
抽出不可能な P-256 キーを保持する macOS Secure Enclave バックエンド。エンクレイブ
生成 → 永続化 → リロード → 署名 → 削除のループは実際に検証されます
hardware であり、コンパイルされた Go デーモンはまだプロファイルの下で実行されていないため、
残りのステップは梱包です。 Linux/TPM または Windows バックエンドはありません。何
各バックエンドが証明しているかどうかは、率直に次のように述べられています。
バックエンドへの署名 ;他の境界線は下にあります
既知の制限と未解決の質問は、
今後の.md 。
非表示ではなく、検証者の出力に表示されます。なんて潔い判決なんだ
証明するかどうかは、How Kessa で正確に述べられています。
作品。
評決は、提供された DID 文書に基づいて行われます。 --dids は
ルートを信頼します。すべての署名は、そのディレクトリから読み取られたキーと照合してチェックされます。
(または --fetch-dids を使用して HTTPS 経由でフェッチされる)、完全に捏造されたエクスポート
名前を付けた DID ドキュメントが一致するように作成されている場合に、クリーンであることを検証します。これは
意図的 (Kessa サービスにアンカーすると、設計全体が無効になります)
これは、PASS がこれらのキーと一致していることを示しており、決して本物ではないことを意味します。もし
輸出と DID 書類が同じ当事者からあなたに届いたことを確認しました
その当事者は自分自身に同意します。 DID ドキュメントを個別に取得します。
完全性には限界があります。封筒の署名にはエントリ数が含まれており、
ログチップなので、誰も事後に署名付きエクスポートを短縮することはできません ( R2-02 )。それはあります
プロではない

施行ポイントは決定したすべてのことをログに記録しました: 署名された短いログ
正直に言って、何かを記録することを拒否した代理人によって署名された短いログは、
ファイルだけでは区別がつきません。先端を固定する必要があるクロージング
強制ポイントが制御できない場所であり、Kessa はまだ制御していません。
失効は、発行者がチェック可能に設定した場合にのみチェック可能です。ホップ
ステータスリスト参照なしで作成されたものは永久に取り消すことができません。検証者
黙ってスキップするのではなく、エントリ ( LIMIT: ) ごとにそうするようになりました。
検証されたポリシーは正しいポリシーではありません。結果性が再導出される
ベアビットとして信頼されるのではなく、実行され、署名されたポリシーからのものであるため、
検証者は、許可がポリシーと一致していることを証明します。
ポイントを公開しました。政策が国民にとって正しいものであることを証明することはできない。
環境。実行されているポリシーを検査して判断します。
ステータスはアクション時のリストではなく、現在のリストと照合してチェックされます。
時間 ( S1 、遅延)。後で取り消しが反転した後の古いエクスポートの再検証
以前は正当だったエントリを PASS から FAIL まで。これは正直なところ、誤った失敗です。
エクスプロイトではありません: 現在のリストのセマンティクスでは、本当に悪い履歴を作成することはできません。
アクションパス。
拒否は独立して再導出可能ではありません。拒否はポリシーに起因する可能性があります。
これには証拠が含まれていないため、拒否されたエントリはそのハッシュが一致すると合格します。
署名とチェーンの証拠は無傷です。正しい否定は正確に証明されます
許可チェックの 1 つが失敗しました。拒否されたエントリに対してこれらのチェックを実行すると、
「正しく拒否された」と「検証者の失敗」を区別できないようにします。
ログには、帰責可能な決定のみが記録されます。これは意図的なものです。
実装上の事故ではなく、所有物です。エントリは、
所有証明が検証されたリクエストなので、すべてのエントリは
校長は誰ですか

明らかに鍵を握っていた。誰も結び付けることができないリクエスト:
検証しないチェーン、または検証しない所有証明: 何も生成しません。
決定とエントリーなし。それは拒否され、次のように監査シンクに報告されます。
代わりにテレメトリーを使用します。
これで本当の穴が塞がれました ( R5-06 )。チェーン検証が唯一のゲートだった時代
書き込みの前に、誰かにエクスポートを与えると、
彼らが監査していたログ。所有権が最初にチェックされるため、これらのエントリは次のようになります。
単に否定されるのではなく、不可能です。
固定的な特性。修正によって変更されることはなく、次のことを意図したものではありません。
エクスポートには、各資格情報とその発行者証明が含まれるため、資格情報を保持している人は誰でも
検証する委任チェーンを再導出することができます。それは意図的です、それは何ですか
検証者は、共有されていない公開鍵に対してオフラインでチェーンを再チェックできます。
それはこの製品の中心的な主張です。チェーンは発行を証明します。
公的である。それは決して所有を証明しませんでした。チェーンを消費するあらゆるものをデザインする
それに応じて、R5-06 は、1 つのパスがより多くのチェーンを確立したと想定したために発生しました。
それよりも。
これで閉じられないこと: エンドポイントにはまだ発信者認証がないため、
アクセスできる人は誰でも送信できますが、単に記録することはできません
何でも。非

[切り捨てられた]

## Original Extract

Attestation and enforcement for delegated agent authority, with an independent offline verifier. - Gneiss-Group/Kessa

GitHub - Gneiss-Group/Kessa: Attestation and enforcement for delegated agent authority, with an independent offline verifier. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
Gneiss-Group
/
Kessa
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
121 Commits 121 Commits .claude .claude .github .github LICENSES LICENSES auditsink auditsink build build cmd cmd docker docker docs docs examples examples internal internal perf perf pkg/ types pkg/ types scripts scripts testdata testdata .dockerignore .dockerignore .gitignore .gitignore .gitleaks.toml .gitleaks.toml CHANGELOG.md CHANGELOG.md CLA-CORPORATE.md CLA-CORPORATE.md CLA.md CLA.md CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LICENSING.md LICENSING.md Makefile Makefile NOTICE.md NOTICE.md PLUGIN_LICENSING.md PLUGIN_LICENSING.md README.md README.md REUSE.toml REUSE.toml SECURITY.md SECURITY.md UPCOMING.md UPCOMING.md go.mod go.mod View all files Repository files navigation
Attestation and enforcement for delegated agent authority. A human delegates to
an org, the org to an agent, the agent to a sub-agent. Each hop narrows
authority and never broadens it, and every consequential action is logged to a
tamper-evident chain that an independent, offline verifier can re-check while
trusting nothing of ours beyond public DID documents.
Same command, one byte changed between runs. The verifier fails at exactly the tampered entry and re-derives every verdict from the files alone. How the GIF is built and regenerated.
Jump to: What this solves · Try it · Status · Standards · Known limits
When a person hands work to an AI agent, and that agent hands part of it to
another, the authority travelling down that chain is normally invisible and
usually undiminished. Kessa makes the chain explicit, narrowing, and provable
after the fact.
Agents inherit far more authority than the task needs. The common pattern
is to hand the agent a copy of the user's token, so a summarization job holds
the power to move money. Kessa issues each hop a new, strictly narrower
credential, and a hop that tries to broaden its own authority is rejected when
it is minted, not when it is used.
Nobody can say what an agent was allowed to do at the time it acted.
Reconstructing that after an incident usually means trusting application logs.
Kessa records the delegation chain, the policy, and the approval as signed
evidence carried with the action.
The system that made the decision also writes the record of it. That is
the party with the most reason to be wrong. Kessa's verifier re-derives every
verdict from the evidence instead of reading the decision the enforcement
point wrote down, so a proxy that cut a corner produces an export that fails.
Verifying usually means trusting the vendor. kessa verify is a separate,
permissively licensed binary that runs offline against files. It reads no
service of ours, and its trust root is a directory of public keys you can
obtain independently.
Revocation does not reach authority already delegated. Kessa checks every
hop against a signed status list at action time, so revoking a credential
mid-chain stops the consequential actions that depend on it.
Tampering with the record is undetectable. The audit log is hash-chained
and signed, and the signature covers the log's length and tip, so entries
cannot be edited, reordered, or quietly dropped after the fact.
Kessa governs authority , not content. It decides whether an action may
proceed; it does not judge what an agent says or inspect the data an approved
action touches.
Three binaries produce evidence and one re-checks it, trusting none of the
three. The full walkthrough, with diagrams of the delegation, enforcement, and
verification stages, plus the precise statement of what a clean verdict proves,
is in How Kessa Works .
Prerequisites: Go 1.26 or newer (the version in go.mod ; an
older toolchain will try to download 1.26 and fail if it cannot), plus make and
bash . There is nothing else to install: Kessa has no third-party Go
dependencies, so make demo builds and runs with the network off. The
container images additionally need a running Docker daemon,
and the provenance check needs the GitHub CLI .
make demo # the whole story: all seven scenarios, end to end, deterministic
Or drive the pieces yourself:
make build # builds every binary into ./bin (the verifier is ./bin/kessa)
# 1. Issue a chain and publish the public artifacts.
./bin/kessa-issuer publish \
--spec examples/issuer/spec.json --keystore examples/issuer/keystore.json \
--root ./public --out ./private/chain.json
# 2. Verify an export against them, offline, nothing running.
./bin/kessa verify \
--export testdata/audit_export_v2.golden.json \
--dids ./public \
--status " https://localhost/orgs/acme/status.json=./public/localhost/orgs/acme/status.json "
# 3. Revoke a mid-chain credential; the same export stops verifying.
./bin/kessa-issuer revoke --spec examples/issuer/spec.json \
--keystore examples/issuer/keystore.json --root ./public --index 42
Exit 0 = every entry verified against carried evidence, 1 = a FAIL or an
integrity-only v1 downgrade (never a clean pass), 2 = usage/IO error.
Nothing is started, nothing is dialled: the verifier's only inputs are files.
--fetch-dids enables HTTPS resolution of public did:web documents, the only
network access kessa can make, and it is off by default.
Kessa works end to end and is under active development. Seven scenarios run start
to finish ( make demo ), and multiple adversarial security review rounds are closed
with every finding fixed. All rounds were self-run AI red-team passes, not a
third-party audit ; no external audit has been commissioned yet. It is
standard-library only, with no third-party dependencies.
Four things carry the central claim, and all four are built: provable attenuation,
provable auditability, the independent verifier, and the full seven-scenario demo.
It is not production-hardened yet. Key handling runs behind a Signer
seam with two backends: a software keystore (the demo, CI, and
--software-key path, whose private key exists in plaintext in the file) and a
macOS Secure Enclave backend holding a non-extractable P-256 key. The Enclave
generate → persist → reload → sign → delete loop is validated on real
hardware , and the compiled Go daemon has not yet run under a profile, so
packaging is the remaining step. There is no Linux/TPM or Windows backend. What
each backend does and does not prove is stated bluntly in
Signing backends ; the other boundaries are under
Known limits , and open questions are collected in
UPCOMING.md .
Surfaced in the verifier's output rather than hidden. What a clean verdict
proves, and does not, is stated precisely in How Kessa
Works .
The verdict is relative to the DID documents you supply. --dids is the
trust root. Every signature is checked against a key read from that directory
(or fetched over HTTPS with --fetch-dids ), so a wholly fabricated export
verifies clean when the DID documents it names are fabricated to match. This is
deliberate (anchoring to a Kessa service would defeat the entire design) but
it means a PASS says consistent with these keys , never genuine . If the
export and the DID documents reached you from the same party, you have confirmed
that party agrees with itself. Obtain the DID documents independently.
Completeness is bounded. The envelope signature covers the entry count and
log tip, so nobody can shorten a signed export after the fact ( R2-02 ). It does
not prove the enforcement point logged everything it decided: a short log signed
honestly and a short log signed by a proxy that declined to record something are
indistinguishable from the file alone. Closing that needs the tip anchored
somewhere the enforcement point does not control, which Kessa does not yet do.
Revocation is only checkable where the issuer made it checkable. A hop
minted with no status-list reference is permanently unrevocable. The verifier
now says so per entry ( LIMIT: ) instead of skipping it silently.
A verified policy is not a correct policy. Consequentiality is re-derived
from the carried, signed policy rather than trusted as a bare bit, so the
verifier proves the allows are consistent with the policy the enforcement
point published. It cannot prove that policy is the right one for the
environment. Inspect the carried policy to judge that.
Status is checked against the current list , not the list as of action
time ( S1 , deferred). Re-verifying an old export after a later revocation flips
previously-legitimate entries from PASS to FAIL. This is an honest false-FAIL,
not an exploit: current-list semantics cannot make a genuinely-bad historical
action pass .
Denials are not independently re-derivable. A denial can stem from policy,
which is not carried evidence, so a denied entry passes when its hash,
signature, and chain evidence are intact. Correct denial is proven by exactly
one of the allow-checks failing; running those checks on denied entries would
make "correctly denied" and "verifier failure" indistinguishable.
The log records only attributable decisions , and that is a deliberate
property rather than an accident of implementation. An entry exists only for a
request whose proof of possession verified, so every entry is bound to a
principal who demonstrably held the key. A request nobody can be tied to: a
chain that does not verify, or a possession proof that does not: produces no
decision and no entry; it is refused, and reported to the audit sink as
telemetry instead.
This closed a real hole ( R5-06 ). When chain verification was the only gate
before a write, giving someone an export gave them the ability to append to the
log they were auditing. Possession is now checked first, so those entries are
impossible rather than merely denied.
A standing characteristic, which the fix does not change and is not meant to:
an export carries each credential with its issuer proof, so anyone holding one
can re-derive a delegation chain that verifies. That is deliberate, it is what
lets the verifier re-check a chain offline against public keys with no shared
secret, and it is the product's central claim. A chain proves issuance , which
is public; it never proved possession . Design anything that consumes a chain
accordingly: R5-06 happened because one path assumed a chain established more
than it does.
What this does not close: the endpoint still has no caller authentication, so
anyone who can reach it may submit , they simply cannot make it record
anything. Non

[truncated]
