---
source: "https://github.com/solara-associates/vaid"
hn_url: "https://news.ycombinator.com/item?id=48909213"
title: "VAID – proof-of-possession signing for AI agents, no network call to verify"
article_title: "GitHub - solara-associates/vaid · GitHub"
author: "solara123"
captured_at: "2026-07-14T17:04:38Z"
capture_tool: "hn-digest"
hn_id: 48909213
score: 1
comments: 2
posted_at: "2026-07-14T16:22:56Z"
tags:
  - hacker-news
  - translated
---

# VAID – proof-of-possession signing for AI agents, no network call to verify

- HN: [48909213](https://news.ycombinator.com/item?id=48909213)
- Source: [github.com](https://github.com/solara-associates/vaid)
- Score: 1
- Comments: 2
- Posted: 2026-07-14T16:22:56Z

## Translation

タイトル: VAID – AI エージェントの所有証明署名、検証のためのネットワーク呼び出しは不要
記事タイトル: GitHub - Solara-associates/vaid · GitHub
説明: GitHub でアカウントを作成して、solara-associates/vaid の開発に貢献します。

記事本文:
GitHub - ソララアソシエイツ/vaid · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
ソララアソシエイツ
/
無駄な
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダー

とファイル
19 コミット 19 コミット .github/ workflows .github/ workflows crates crates python python scripts scripts .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
検証可能なエージェント アクション ID (VAID) のオープン標準レイヤー。
VAID は、自律エージェントが実行するアクションにバインドされたポータブル ID です。
このリポジトリは、VAID にバインドされたリクエストを正規化して署名する方法を定義します。
これらを生成および検証する参照 SDK (Rust および Python) が同梱されています。
署名。これは相互運用性契約です。これに従うすべてのクライアント
共有ランタイムなしで、準拠する検証ツールが受け入れるバイトを生成します。
間にネットワークサービスはありません。どちらのリファレンス SDK も同じフリーズされたものを再現します
バイトごとの適合ベクトル — これは、言語を超えた証明です。
具体的 (2 つの言語、1 つのベクトルを参照)。
バイトレベルの標準、2 つの言語によるリファレンス実装、委任付きのリファレンスミント、LangChain 統合、および完了レコード:
vaid-pop (Rust、crates/vaid-pop) は所有証明 (PoP) です。
原始的な。 1 つの正規化パスを定義します: RFC 8785 JSON Canonicalization
スキーム (JCS)、次に正規バイト上の SHA-256、次に純粋な Ed25519
32 バイトのダイジェスト上の署名。また、取得するリクエスト ペイロードも定義します。
signed およびペイロードがバインドする VAID ID タイプ。これはバイトレベルです
コードとして記述された仕様。
vaid-client (Rust、crates/vaid-client ) は、上に構築されたリファレンス SDK です。
その原始的なもの。鋳造されたVAID文書とホルダーキーを4つに変換します
署名付きヘッダー

リクエストは実行されますが、いずれも再実装されません。
正規化。これは Vaid-pop のみに依存します。
vaid-pop (Python、python/vaid-pop ) は Python 参照署名者です。
同じ PoP コントラクトの単一の Python 定義。 Rustを反映しています
正規化パスは正確に (RFC 8785 JCS → SHA-256 → pure Ed25519) であり、
同じ凍結ベクトルにロックされます。それは暗号化のみに依存しており、
rfc8785 、他には何もありません。
vaid-mint (Rust + Python、 crates/vaid-mint 、 python/vaid-mint ) はリファレンス ミントです。 VAID を発行し、軽減された委任 ( mint_child 、子の権限は常に親のサブセットです) をサポートし、その信頼モデルを明確に文書化します。 Rust クレートの README ( crates/vaid-mint ) は、現在の 0.1.2 状態、つまり検証時に適用される TTL とプラグ可能な RevocationCheck シームを反映しています。 Python パッケージ ( python/vaid-mint ) はまだ追いついていません。0.1.1 のままで、元の勧告のみの有効期限があり、継ぎ目がなく、保留中の問題 #1 です。耐久性のあるものとそうでないものについては、使用している実装に一致する README をお読みください。
vaid-langchain (Python、python/vaid-langchain) は、httpx.Auth アダプター経由で VAID コントラクトを使用してリクエストに署名する LangChain 統合です。
完了レコード ( vaid-pop 、complete_v1.json ベクトル) — エージェントが行ったと主張する内容についての自己報告の出所記録。現在の単一層保証: 自己申告のみであり、そのタイプ自体のドキュメントにその旨記載されています。
それが公開範囲全体です。サーバー、データベース、ランタイムはありません。
セルフホストを選択した場合は、造幣局を超えて立ち上がってください。 Rust クレートを Cargo プロジェクトに追加するか、Python パッケージを pip install して呼び出します。
開発者は VAID 識別子を作成し、それに対してリクエストに署名し、検証することができます。
その署名は、スタンドアロンで、これらの c のみを使用します

料金。
プリミティブを使用して直接署名および検証する
chrono :: Utc を使用します。
リング:: ランド:: SystemRandom を使用します。
リングを使用:: 署名:: { Ed25519KeyPair , KeyPair } ;
sha2 :: { ダイジェスト , Sha256 } ; を使用します。
vaid_pop :: VaidId を使用します。
vaid_pop :: request_auth :: RequestAuthPayload を使用します。
vaid_pop :: vaid_pop :: {sign_payload , verify_signed_pa​​yload } を使用します。
// ペイロードは body_sha256 をバインドするため、小文字の 16 進数の SHA-256 である必要があります。
// 正確なリクエスト本文のバイト。以下の SDK がこれを計算します。ここにそれが示されています
// 明示的にそのため、プリミティブな例では空の文字列ではなく実際の本体をバインドします。
fn hex_sha256 (バイト : & [ u8 ] ) -> 文字列 {
Sha256::ダイジェスト (バイト) 。イター ( ) 。マップ ( |b| 形式 ! ( "{b:02x}" ) ) 。集める（）
}
// 1. アクションの VAID 識別子を作成し、Ed25519 キーを保持します。
let vaid = VaidId :: new () ;
let rng = SystemRandom :: new ( ) ;
pkcs8 = Ed25519KeyPair::generate_pkcs8 (&rng) とします。アンラップ ( ) ;
let key = Ed25519KeyPair :: from_pkcs8 ( pkcs8 . as_ref ( ) ) 。アンラップ ( ) ;
// 2. この VAID が許可するリクエストを説明します。
let request_body = br#"{"タスク":"第 3 四半期レポートの要約"}"# ;
let ペイロード = RequestAuthPayload {
vaid_id : vaid 、
メソッド: "POST" 。 ( ) に、
パス: "/v1/agents/execute" 。 ( ) に、
body_sha256 : hex_sha256 ( request_body ) 、
テナント ID : "acme" 。 ( ) に、
タイムスタンプ : Utc :: now ( ) 、
client_nonce : "要求ごとの新鮮なノンス" 。の中へ （ ） 、
} ;
// 3. 署名: JCS、次に SHA-256、次にダイジェスト上の Ed25519。
let 署名 =sign_payload (& ペイロード , & キー ) ;
// 4. 所有者の公開鍵と照合して検証します。
let verifyed = verify_signed_pa​​yload (& payload , key . public_key() . as_ref() , & signed ) ;
主張してください！ (検証済み) ;
SDKを使用してリクエストヘッダーを生成する
HTTP リクエストを認証する一般的なケースでは、SDK

ミントを取ります
VAID ドキュメントとキーを入力し、添付する 4 つのヘッダーを返します。それはハッシュします
body を作成し、新しい nonce を生成し、現在のタイムスタンプをスタンプします。
リング :: 署名 :: Ed25519KeyPair を使用します。
vaid_client :: RequestSigner を使用します。
letsigner = RequestSigner :: from_vaid_json (vaid_document_json , key ) ? ;
ヘッダー = 署名者とします。 Sign_headers ( "POST" 、 "/v1/agents/execute" 、 request_body ) ? ;
// headers.into_pairs() は次の順序で結果を返します。
// x-synthera-vaid、x-synthera-timestamp、x-synthera-nonce、x-synthera-signature
ヘッダーの for (name, value)。 into_pairs ( ) {
リクエストをset_header (名前, 値) ;
}
このパスの実行可能なバージョンは次のとおりです。
crates/vaid-client/examples/emit_pop.rs 。
バイトが移植可能であることの証明
署名パスは凍結されたテスト ベクトルによって固定されます。
crates/vaid-client/tests/vectors/operator_pop_v1.json 。適合性テスト
そのベクターの正確な SHA-256 ダイジェストとその正確な Ed25519 署名を再現します。
固定入力。それが相互運用性の保証を具体化したものです。
同じベクトルにヒットする独立した実装は、これとバイト互換性があります。
1つ。
貨物試験
適合スイートとプリミティブ自体のラウンドトリップおよび改ざん拒否
テストは他に何も存在しない状態で実行されます。
凍結されたベクター crates/vaid-client/tests/vectors/operator_pop_v1.json は、
唯一の真実の情報源。 python/vaid-pop の Python 参照署名者
バイト同一のコピーを提供し、同じ SHA-256 ダイジェストを再現します。
同じ固定入力からの同じ Ed25519 署名 — から証明されています。
インストールされたパッケージ。リポジトリのチェックアウトは必要ありません。
cd Python/vaid-pop
pip インストール 。
vaid-pop-conformance # PASS = インストールされた署名者 == 凍結ベクトル、バイトごと
したがって、相互運用性の保証は仕様書に関する主張ではありません。それは 2 つのことです。
で

依存する実装は 2 つの言語で行われ、共有ランタイムはなく、
同じバイト。上記の Rust カーゴ テストと Python vaid-pop-conformance
同じベクトルに対してアサートします。リポジトリのポップ準拠 CI ジョブは、
いかなる発散でも失敗します。それが標準であり、証明されています。
このリポジトリは、標準、その参照署名者、参照ミント、LangChain 統合、および完了レコードです。非公開かつ商用のままのものが 2 つあります。
VAID に許可される内容を表現するためのポリシー言語。
運用環境でミントを実行するホストされた機関 - KMS ベースのカーネル キー、記録の監査、耐久性のあるハッシュ チェーン失効、およびポリシー/メッシュ/フェデレーション コントロール プレーン。
ここでの参照ミントは、委任と減衰の形状を証明しています。それはホストされる権威ではありません。取り消しは、「商用」として申請するのではなく、明確に名前を付ける価値のあるシームです。Rust クレート (vaid-mint 0.1.2+) には、プラグイン可能な RevocationCheck シームが同梱されています。これは追加的で、メモリ内デフォルトで、検証時に強制適用される VAID 有効期限 (TTL) (デフォルトは 1 時間) です。そのため、セルフホスティング者は、クレートにパッチを適用することなく、独自の取り消しバックエンドを接続できます。商用であり続けるのは、永続的な失効そのものです。クレートは継ぎ目を出荷するものであり、永続的で再起動しても存続するハッシュ チェーン ストアではありません。これは Rust クレートにのみ適用されます。PyPI vaid-mint パッケージは 0.1.1 のままで、元の勧告のみの有効期限があり、継ぎ目はありません (問題 #1 で追跡されています)。現在実稼働環境でこれを実行している場合、何が耐久性があり、何が耐久性がないのか、そしてそれを軽減する方法については、 crates/vaid-mint 独自の信頼モ​​デルのドキュメントを参照してください。
プロダクション コントロール プレーンは別個の商用製品であり、本製品には含まれていません。
リポジトリ。その製品は、ホスト型 VAID 機関を提供し、
アイデンティティを剥奪する、ポリ

各 VAID が何を実行できるかを決定する cy エンジン。
テナント間でアクションをルーティングするフェデレーション層、強制メッシュ
通話時にそれらの決定を適用し、記録を保持する監査を行います。
検証可能な歴史。ここにあるものを使用するために必要なものは何もありません。
ここに含まれています。このリポジトリはオープン スタンダードとして独立しています。
VAID は相互運用性契約であるため、貢献の基準は具体的です。
変更を行う場合は、両方の参照 SDK が凍結された適合ベクトルを再現し続ける必要があります。
バイトごとに。
CONTRIBUTING.md — 開発セットアップ (Rust + Python)、
適合バー、および規格に影響を与える変更を提案する方法。
SECURITY.md — 脆弱性を非公開で報告する
( info@solara.associates );彼らのために公開問題を開かないでください。
CODE_OF_CONDUCT.md — 貢献者規約 2.1。
Apache-2.0 — 「ライセンス」と「通知」を参照してください。
著作権 © 2026 ソララ.アソシエイツ。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to solara-associates/vaid development by creating an account on GitHub.

GitHub - solara-associates/vaid · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
solara-associates
/
vaid
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits .github/ workflows .github/ workflows crates crates python python scripts scripts .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
The open standard layer for verifiable agent-action identity (VAID).
A VAID is a portable identity bound to an action that an autonomous agent takes.
This repository defines how a VAID-bound request is canonicalized and signed, and
ships reference SDKs — in Rust and Python — that produce and verify those
signatures. It is the interoperability contract: any client that follows it
produces bytes that any conforming verifier accepts, with no shared runtime and
no network service in between. Both reference SDKs reproduce the same frozen
conformance vector byte-for-byte — that is the cross-language proof, made
concrete (see Two languages, one vector ).
The byte-level standard, reference implementations in two languages, a reference mint with delegation, a LangChain integration, and completion records:
vaid-pop (Rust, crates/vaid-pop ) is the proof-of-possession (PoP)
primitive. It defines one canonicalization path: RFC 8785 JSON Canonicalization
Scheme (JCS), then SHA-256 over the canonical bytes, then a pure Ed25519
signature over the 32-byte digest. It also defines the request payload that gets
signed and the VAID identity types that payload binds. This is the byte-level
specification, written as code.
vaid-client (Rust, crates/vaid-client ) is the reference SDK built on
that primitive. It turns a minted VAID document and a holder key into the four
signed headers a request carries, and it does not reimplement any of the
canonicalization. It depends only on vaid-pop .
vaid-pop (Python, python/vaid-pop ) is the Python reference signer — the
single Python definition of the same PoP contract. It mirrors the Rust
canonicalization path exactly (RFC 8785 JCS → SHA-256 → pure Ed25519) and is
locked to the same frozen vector. It depends only on cryptography and
rfc8785 , nothing else.
vaid-mint (Rust + Python, crates/vaid-mint , python/vaid-mint ) is the reference mint. It issues VAIDs, supports attenuated delegation ( mint_child , where a child's authority is always a subset of its parent's), and documents its trust model plainly. The Rust crate README ( crates/vaid-mint ) reflects the current 0.1.2 state — TTL enforced at verification and a pluggable RevocationCheck seam. The Python package ( python/vaid-mint ) has not yet caught up: it remains on 0.1.1 with the original advisory-only expiry and no seam, pending issue #1 . Read the README that matches the implementation you're using for what's durable and what isn't.
vaid-langchain (Python, python/vaid-langchain ) is a LangChain integration that signs requests using the VAID contract via an httpx.Auth adapter.
completion records ( vaid-pop , completion_v1.json vector) — a self-reported provenance record for what an agent claims it did. Single-tier assurance today: self-reported only, and the type's own documentation says so.
That is the entire open scope. There is no server, no database, and no runtime to
stand up beyond the mint if you choose to self-host it. You add the Rust crates to a Cargo project, or pip install the Python packages, and call them.
A developer can create a VAID identifier, sign a request against it, and verify
that signature, standalone, using only these crates.
Sign and verify directly with the primitive
use chrono :: Utc ;
use ring :: rand :: SystemRandom ;
use ring :: signature :: { Ed25519KeyPair , KeyPair } ;
use sha2 :: { Digest , Sha256 } ;
use vaid_pop :: VaidId ;
use vaid_pop :: request_auth :: RequestAuthPayload ;
use vaid_pop :: vaid_pop :: { sign_payload , verify_signed_payload } ;
// The payload binds body_sha256, so it must be the lowercase hex SHA-256 of the
// exact request body bytes. The SDK below computes this for you; here it is shown
// explicitly so the primitive example binds a real body, not an empty string.
fn hex_sha256 ( bytes : & [ u8 ] ) -> String {
Sha256 :: digest ( bytes ) . iter ( ) . map ( |b| format ! ( "{b:02x}" ) ) . collect ( )
}
// 1. Create a VAID identifier for the action, and hold an Ed25519 key.
let vaid = VaidId :: new ( ) ;
let rng = SystemRandom :: new ( ) ;
let pkcs8 = Ed25519KeyPair :: generate_pkcs8 ( & rng ) . unwrap ( ) ;
let key = Ed25519KeyPair :: from_pkcs8 ( pkcs8 . as_ref ( ) ) . unwrap ( ) ;
// 2. Describe the request this VAID is authorizing.
let request_body = br#"{"task":"summarize the Q3 report"}"# ;
let payload = RequestAuthPayload {
vaid_id : vaid ,
method : "POST" . into ( ) ,
path : "/v1/agents/execute" . into ( ) ,
body_sha256 : hex_sha256 ( request_body ) ,
tenant_id : "acme" . into ( ) ,
timestamp : Utc :: now ( ) ,
client_nonce : "a-fresh-per-request-nonce" . into ( ) ,
} ;
// 3. Sign: JCS, then SHA-256, then Ed25519 over the digest.
let signature = sign_payload ( & payload , & key ) ;
// 4. Verify against the holder's public key.
let verified = verify_signed_payload ( & payload , key . public_key ( ) . as_ref ( ) , & signature ) ;
assert ! ( verified ) ;
Produce request headers with the SDK
For the common case of authenticating an HTTP request, the SDK takes the minted
VAID document and your key and returns the four headers to attach. It hashes the
body, generates a fresh nonce, and stamps a current timestamp for you.
use ring :: signature :: Ed25519KeyPair ;
use vaid_client :: RequestSigner ;
let signer = RequestSigner :: from_vaid_json ( vaid_document_json , key ) ? ;
let headers = signer . sign_headers ( "POST" , "/v1/agents/execute" , request_body ) ? ;
// headers.into_pairs() yields, in order:
// x-synthera-vaid, x-synthera-timestamp, x-synthera-nonce, x-synthera-signature
for ( name , value ) in headers . into_pairs ( ) {
request . set_header ( name , value ) ;
}
A runnable version of this path is in
crates/vaid-client/examples/emit_pop.rs .
Proof that the bytes are portable
The signing path is pinned by a frozen test vector,
crates/vaid-client/tests/vectors/operator_pop_v1.json . The conformance test
reproduces that vector's exact SHA-256 digest and its exact Ed25519 signature from
the fixed inputs. That is the interoperability guarantee made concrete: an
independent implementation that hits the same vector is byte-compatible with this
one.
cargo test
The conformance suite and the primitive's own round-trip and tamper-rejection
tests run with nothing else present.
The frozen vector crates/vaid-client/tests/vectors/operator_pop_v1.json is the
single source of truth. The Python reference signer under python/vaid-pop
vendors a byte-identical copy of it and reproduces the same SHA-256 digest and
the same Ed25519 signature from the same fixed inputs — proven from the
installed package, with no repo checkout required:
cd python/vaid-pop
pip install .
vaid-pop-conformance # PASS = installed signer == frozen vector, byte-for-byte
So the interoperability guarantee is not a claim about a spec document — it is two
independent implementations, in two languages, with no shared runtime, hitting the
same bytes. The Rust cargo test above and the Python vaid-pop-conformance
assert against the same vector; the repo's pop-conformance CI job runs both and
fails on any divergence. That is the standard, proven.
This repository is the standard, its reference signer, a reference mint, a LangChain integration, and completion records. Two things remain closed and commercial:
The policy language for expressing what a VAID is permitted to do.
The hosted authority that runs a mint in production — KMS-backed kernel keys, an audit-of-record, durable hash-chained revocation, and a policy/mesh/federation control plane.
The reference mint here proves the shape of delegation and attenuation; it is not that hosted authority. Revocation is the seam worth naming plainly rather than filing under "commercial": the Rust crate ( vaid-mint 0.1.2+) now ships a pluggable RevocationCheck seam — additive, with an in-memory default, and with VAID expiry (TTL) hard-enforced at verification (1-hour default) — so a self-hoster can wire their own revocation backend without patching the crate. What stays commercial is durable revocation itself: the crate ships the seam, not a durable, restart-surviving hash-chained store. This applies to the Rust crate only — the PyPI vaid-mint package remains on 0.1.1 with the original advisory-only expiry and no seam, tracked in issue #1 . See crates/vaid-mint 's own trust-model documentation for exactly what's durable, what isn't, and how to mitigate it if you're running this in production today.
The production control plane is a separate commercial product and is not in this
repository. That product provides the hosted VAID authority that issues and
revokes identities, the policy engine that decides what each VAID may do, the
federation layer that routes action across tenants, the enforcement mesh that
applies those decisions at call time, and the audit-of-record that retains a
verifiable history. None of that is required to use what is here, and none of it
is included here. This repository stands on its own as the open standard.
VAID is an interoperability contract, so the bar for contributions is concrete:
any change must keep both reference SDKs reproducing the frozen conformance vector
byte-for-byte.
CONTRIBUTING.md — dev setup (Rust + Python), the
conformance bar, and how to propose standard-affecting changes.
SECURITY.md — report vulnerabilities privately
( info@solara.associates ); please don't open public issues for them.
CODE_OF_CONDUCT.md — Contributor Covenant 2.1.
Apache-2.0 — see LICENSE and NOTICE .
Copyright © 2026 solara.associates.
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
