---
source: "https://github.com/umudhasanli/flare-redact"
hn_url: "https://news.ycombinator.com/item?id=49019331"
title: "Flare Redact: Stop secrets from reaching logs, telemetry, and LLMs"
article_title: "GitHub - umudhasanli/flare-redact: International secret & PII redaction for JS/TS — logs, prompts, HTTP, datasets. 20+ languages, checksum-validated national IDs, reversible LLM layer. Zero deps, ReDoS-safe. · GitHub"
author: "umudhasanli"
captured_at: "2026-07-23T11:00:49Z"
capture_tool: "hn-digest"
hn_id: 49019331
score: 1
comments: 0
posted_at: "2026-07-23T10:21:38Z"
tags:
  - hacker-news
  - translated
---

# Flare Redact: Stop secrets from reaching logs, telemetry, and LLMs

- HN: [49019331](https://news.ycombinator.com/item?id=49019331)
- Source: [github.com](https://github.com/umudhasanli/flare-redact)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T10:21:38Z

## Translation

タイトル: Flare Redact: シークレットがログ、テレメトリ、LLM に到達するのを阻止する
記事のタイトル: GitHub - umudhasanli/flare-redact: JS/TS の国際秘密と PII 編集 - ログ、プロンプト、HTTP、データセット。 20 以上の言語、チェックサム検証済みの国民 ID、可逆 LLM レイヤー。ゼロデプス、ReDoS セーフ。 · GitHub
説明: JS/TS の国際秘密および PII 編集 — ログ、プロンプト、HTTP、データセット。 20 以上の言語、チェックサム検証済みの国民 ID、可逆 LLM レイヤー。ゼロデプス、ReDoS セーフ。 - ウムダサンリ/フレアリダクト

記事本文:
GitHub - umudhasanli/flare-redact: JS/TS の国際秘密および PII 編集 - ログ、プロンプト、HTTP、データセット。 20 以上の言語、チェックサム検証済みの国民 ID、可逆 LLM レイヤー。ゼロデプス、ReDoS セーフ。 · GitHub
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
却下エール

RT
{{ メッセージ }}
ウムダサンリ
/
フレアリダクト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
23 コミット 23 コミット .github .github アセット アセット ベンチ ベンチ ビン ビン ドキュメント ドキュメントの例 例 src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
漏洩する前に、ログ、プロンプト、テキスト内の秘密と PII を非表示にします。
🌐 デフォルトでインターナショナル — 24 言語
🇬🇧 🇨🇳 🇮🇳 🇪🇸 🇸🇦 🇫🇷 🇵🇹 🇷🇺 🇯🇵 🇩🇪 🇰🇷 🇹🇷 🇮🇹 🇮🇷 🇵🇱 🇺🇦 🇳🇱 🇻🇳 🇮🇩 🇹🇭 🇬🇷 🇮🇱 🇦🇿 🇷🇴
漏洩したすべての秘密には同じ起源があります。つまり、誰かがオブジェクトを記録し、
パスワード、トークン、または API キーがその中にありました。コードは無害に見えました —
logger.info({ user }) — ただし、ユーザーはセッション トークンを保持しており、現在はそれが存在しています。
ログ アグリゲータ、エラー トラッカー、および 3 つのベンダーのシステムを永久に使用できます。
Flare-redact は、そのデータをラップする関数の 1 つです。内容を読み取って、
フィールド名だけではなく、誰かがフリーテキストに貼り付けた AWS キーも捕捉します。
note 、Authorization ヘッダーの JWT、スタック トレースのカード番号 — そして
それらをマスクして、デバッグ可能な状態を維持するのに十分なヒントを保持します。
'flare-redact' から { redact } をインポートします。
redact ('ユーザー alice@corp.com は 4242 4242 4242 4242、トークン ghp_ で支払いました' + 'a' .repeat (36) ) ;
// → 'ユーザー a***@*** は **** **** **** 4242、トークン ghp_*** で支払いました'
設定するものは何もありません

。維持するフィールド パスのリストがありません。ネイティブのビルドステップはありません。
同じ問題は現在、LLM 呼び出しという新しいアドレスを持っています。 OpenAI をラップするか、
人間のクライアントと検出されたシークレットはプロンプトから削除され、復元されます。
応答 — 参照は存続しますが、モデルはそれらの元の値を参照することはありません。
ジャンプしてください↓
LLM に到達する前にプロンプトを編集する
コンテキストおよびモデル支援型 PII
ステージング用にデータセットを匿名化する
シークレットが侵入するとビルドが失敗する
多言語の秘密の語彙とID
npm インストールフレアリダクト
ノード 20 以上では、ブラウザーとエッジ ランタイムでも実行され、依存関係はありません。
0.8からアップグレードしますか? 0.9 移行ノートを読む
保護されたトランスフォームまたはボールトの動作を変更する前に。
リポジトリのクローンを作成し、次の小さなアプリケーションをローカルで実行します。
最初の実行前に npm run build を実行し、サンプルの依存関係をインストールします。
文字列、配列、オブジェクトを再帰的に。渡したシェイプは、あなたが渡したシェイプです。
戻ってください。
'flare-redact' から { redact } をインポートします。
編集 ( {
ユーザー: 'bob@corp.com' 、
パスワード: 'hunter2' 、
トークン: [ 'ghp_' + 'b' .繰り返し (36) ]、
注: '私の aws キーは AKIAIOSFODNN7EXAMPLE です' 、
} ) ;
// →
// {
// ユーザー: 'b***@***',
// パスワード: '***', // 機密フィールド名
// トークン: ['ghp_***'],
// 注: '私の aws キーは AKIA***', // フリーテキスト内にあります
// }
LLM に到達する前にプロンプトを編集する
アプリはユーザー データを OpenAI または Anthropic に送信します。そのプロンプトのどこかに、
顧客の電子メール、API キー、またはカード番号がシステムから離れます。
クライアントを一度ラップすると、検出されたシークレットがプロンプトから削除され、元に戻されます。
返信で。モデルはそれらの元の値を決して参照しません。あなたのコードは
必要な参照。
'flare-redact/llm' から {wrapOpenAI} をインポートします。
const openai = WrapOpenAI (new OpenAI() ) ;
const res = オペレーションを待つ

ない。チャット 。完成品。作成 ( {
モデル: 'gpt-4o' 、
メッセージ : [ { 役割 : 'ユーザー' 、内容 : '請求書を alice@corp.com、カード 4242 4242 4242 4242 に電子メールで送信' } ] 、
} ) ;
アプリが送信 → 請求書を alice@corp.com、カード 4242 4242 4242 4242 に電子メールで送信します
モデルが見る → 請求書を [FR_EMAIL_7f2a…]、カード [FR_CREDIT_CARD_19be…] に電子メールで送信します。
あなたのアプリは alice@corp.com に送信されます。カード 4242 4242 4242 4242 は保存されませんでした。
WrapAnthropic は、システム プロンプトを含むmessages.create に対しても同じことを行います。
どちらもストリーミングを処理します。プレースホルダーは、一方が複数に分割されている場合でも復元されます。
塊。ボールトを保持したい場合は、redactPrompt(text) もあります。
あなた自身。
データについて推論する必要があるかどうかに応じてモードを選択してください
隠した後。
redact ( 'bob@corp.com' , { mode : 'mask' } ) ; // 'b***@***' (デフォルト)
redact ( 'bob@corp.com' , { mode : 'label' } ) ; // '[編集済み:電子メール]'
const protectedOptions = {transformSecret:プロセス。環境 。 FLARE_REDACT_SECRET } ;
redact ( 'bob@corp.com' , { ... protectedOptions , mode : 'hash' } ) ;
// 'email_3baf4d28d7c88317a…' — HMAC-SHA-256 フィンガープリント
redact ( 'bob@corp.com' , { ... protectedOptions , mode : 'pseudonym' } ) ;
// 'kqz@rwmp.dnu' — キー付き、決定論的、文字クラスを保持
redact ( 'bob@corp.com' , { ... protectedOptions , mode : 'surrogate' } ) ;
// 'user_93a78c61e204@example.invalid' — 型一貫性のある合成値
保護された決定論的モードには、transformSecret が必要です。彼らは決して黙ってはいない
ソルトされていない公開指紋に戻ります。ハッシュは相関関係に役立ちます。
pseudonym は元の文字の形状を保持し、surrogate は入力されたものを出力します
予約済みドメインの電子メールや Luhn で有効なカード番号などの合成値。
環境または相関ドメインごとに別のシークレットを使用します。
仮名はデリベです

フォーマットを保持する暗号化とはあまり説明されていません。
これは不可逆的な仮名化であり、NIST FF1 ではありません。古い fpe 名はそのまま残ります
互換性エイリアスとして使用されますが、非推奨です。
または、すべてを 1 つの固定文字列に置き換えます。
redact (ペイロード , { マスク : '█' } ) ;
redact(ペイロード,{マスク:({ディテクタ})=>`<${ディテクタ.id}>`});
可逆的な編集
オリジナルを返す必要がある場合 — 上記の LLM ケース、またはデータを第三者に渡す場合
信頼できないシステムを復元する場合は、保管庫を使用してください。それぞれの秘密を交換します
安定したプレースホルダーとして使用され、マッピングが記憶されます。
'flare-redact' から { createVault } をインポートします。
const vault = createVault() ;
const 安全 = ボールト。 redact ('カード 4242 4242 4242 4242 で bob@corp.com に請求') ;
// 'カード [FR_CREDIT_CARD_19be63…] に [FR_EMAIL_7f2ad4…] をチャージ'
金庫。復元 (安全) ;
// 'カード 4242 4242 4242 4242 で bob@corp.com に請求してください'
同じ値は 1 つの Vault 内で同じプレースホルダーを取得するため、参照は残ります。
往復。デフォルトのプレースホルダーには、グローバルの代わりに 96 個のランダム ビットが含まれます。
シーケンス番号。人間が判読できる [EMAIL_1] カウンタは引き続き次の方法で利用できます。
信頼できるローカル ワークフローの場合は、createVault({ placeholderStyle: 'readable' })。
マッピングは元のデータと同じくらい機密性が高くなります。永続化する前に暗号化します。
'flare-redact' からインポート { sealVault 、 openVault 、復元 } ;
const encrypted = await sealVault(vault, process.env.FLARE_REDACT_VAULT_PASSWORD);
fs を待ちます。 writeFile( 'session.vault.json' , JSON . stringify ( encrypted ) , { mode : 0o600 } ) ;
constエントリー=await openVault(暗号化、プロセス.env.FLARE_REDACT_VAULT_PASSWORD);
復元 (安全、新しいマップ (エントリ)) ;
密閉された保管庫では、新鮮な塩を含む PBKDF2-SHA-256 と新鮮な塩を含む AES-256-GCM を使用します。
ノンス。間違ったパスワードや変更されたファイルはフェールクローズされます。

CLI から、 --vault および --restore はデフォルトで暗号化されたファイルを使用します。パスワード
FLARE_REDACT_VAULT_PASSWORD (または によって指定された変数) から取得されます。
--vault-password-env ) なので、シェル履歴には表示されません。
export FLARE_REDACT_VAULT_PASSWORD= ' 本番環境でシークレット マネージャーを使用する '
Flare-redact --vault session.vault.json < input.txt > safety.txt
Flare-redact --restore session.vault.json <safe.txt>restored.txt
コンテキストおよびモデル支援型 PII
構造化された識別子と資格情報は、決定論的なルールによって最適に処理されます。
そしてチェックサムバリデーター。名前とアドレスにはコンテキストが必要なので、保守的な 3 つの
検出器はオプトインです:
const 結果 = スキャン (
'顧客名: アリス例;住所: 120 Cedar Street;生年月日: 1990-04-23' 、
{ 有効にする : [ 'コンテキスト' ] } ,
) ;
// 人物名、住所、生年月日
// 各結果には、リスク、信頼度、正確な感知範囲が含まれます
より広範な多言語フリーテキスト PII の場合、カップリングなしでローカル モデルを接続します
依存関係のないコアを 1 つの ML ランタイムに接続します。
const ポリシー = {
セマンティックプロバイダー: {
非同期検出 (テキスト) {
戻り [ {
検出器: 'person_model' 、ラベル: 'person' 、
理由: 「ローカルの多言語 NER の結果。」 、
開始 : 12 、終了 : 25 、信頼度 : 0.94 、リスク : '高' 、
} ] ;
} 、
} 、
minConfidence : 0.8 、
} ;
constsafe = await redactAsync(input,policy);
意味論的スパンと決定論的スパンは、同じオーバーラップ調停に入ります。リスクが高く、
通常のものではなく、優先度が高く、より検証された結果が優先されます。
式が最初に実行されることがあります。
独自のローカル モデルまたは API を介してチャット インターフェイスを構築している場合、
session はドロップイン層です。 1 つのセッションで 1 つのボールトが保持されるため、値が保持されます
すべてのターンで同じプレースホルダー: 途中でユーザーのメッセージをマスクします。
途中でモデルの応答を復元します。それはmです

odel に依存しない同期型。
に依存するのではなく、独自のターゲット ランタイムで npm run ベンチマークを実行します。
ハードウェアに依存しない遅延の主張。
'flare-redact' からインポート { createSession } ;
const session = createSession ( {enable : [ 'pii' ] } ) ;
// 途中で — モデルはプレースホルダーのみを参照します
const プロンプト = セッション。編集 (userMessage) ;
const Reply = myModel を待ちます。生成 (プロンプト) ;
// 終了時 — ユーザーには実際の値が表示されます
show (セッション.復元(応答)) ;
ストリーミング？プレースホルダーが複数のチャンクに分割されている場合でも、トークンごとに復元します。
const out = セッション。ストリーム （ ） ;
for await (modelStream の const チャンク) process (out .push (chunk . text)) ;
プロセス (out .flush() ) ;
session.redactMessages([{ role, content }]) はチャット配列全体を一度にマスクします。
そして session.reset() が新たな会話を開始します。検出された元の値はそのまま残ります
アプリが可逆参照を保持している間、ローカルに保存されます。
検出者は製品のコード名、プロジェクト名、社内用語を知ることはできません。
それで彼らにリストを渡します。用語は、指定した単語 (あらゆる言語、
最長一致が最初、単語境界セーフ）、一方向または可逆。
// 一方向、独自の置換テキストを使用
redact ( 'Falcon で Project Zeus を起動' , {
用語 : { 'プロジェクト ゼウス' : '[機密]' 、'ファルコン' : '[機密]' } 、
} ) ;
// → '[CLASSIFIED] で [CLASSIFIED] を起動'
// 可逆 — モデルに送信し、それを戻す

[切り捨てられた]

## Original Extract

International secret & PII redaction for JS/TS — logs, prompts, HTTP, datasets. 20+ languages, checksum-validated national IDs, reversible LLM layer. Zero deps, ReDoS-safe. - umudhasanli/flare-redact

GitHub - umudhasanli/flare-redact: International secret & PII redaction for JS/TS — logs, prompts, HTTP, datasets. 20+ languages, checksum-validated national IDs, reversible LLM layer. Zero deps, ReDoS-safe. · GitHub
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
umudhasanli
/
flare-redact
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
23 Commits 23 Commits .github .github assets assets bench bench bin bin docs docs examples examples src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Hide secrets & PII in logs, prompts, and text — before they leak.
🌐 International by default — 24 languages
🇬🇧 🇨🇳 🇮🇳 🇪🇸 🇸🇦 🇫🇷 🇵🇹 🇷🇺 🇯🇵 🇩🇪 🇰🇷 🇹🇷 🇮🇹 🇮🇷 🇵🇱 🇺🇦 🇳🇱 🇻🇳 🇮🇩 🇹🇭 🇬🇷 🇮🇱 🇦🇿 🇷🇴
Every leaked secret has the same origin story: someone logged an object, and a
password, token, or API key was sitting inside it. The code looked innocent —
logger.info({ user }) — but user carried a session token, and now it's in
your log aggregator, your error tracker, and three vendors' systems forever.
flare-redact is one function you wrap around that data. It reads the content ,
not just the field names, so it catches the AWS key someone pasted into a free-text
note , the JWT in an Authorization header, the card number in a stack trace — and
masks them, keeping just enough of a hint to stay debuggable.
import { redact } from 'flare-redact' ;
redact ( 'User alice@corp.com paid with 4242 4242 4242 4242, token ghp_' + 'a' . repeat ( 36 ) ) ;
// → 'User a***@*** paid with **** **** **** 4242, token ghp_***'
Nothing to configure. No list of field paths to maintain. No native build step.
The same problem now has a new address: your LLM calls. Wrap your OpenAI or
Anthropic client and detected secrets are stripped from prompts and restored in
the reply — the model never sees those original values, while references survive.
Jump to it ↓
Redact prompts before they reach an LLM
Contextual and model-assisted PII
Anonymize a dataset for staging
Fail a build when a secret sneaks in
Multilingual secret vocabulary and IDs
npm install flare-redact
Node 20+, and it runs in the browser and edge runtimes too — zero dependencies.
Upgrading from 0.8? Read the 0.9 migration notes
before changing protected transform or vault behavior.
Clone the repository and run these small applications locally:
Run npm run build and install an example's dependencies before its first run.
Strings, arrays, and objects, recursively. The shape you pass in is the shape you
get back.
import { redact } from 'flare-redact' ;
redact ( {
user : 'bob@corp.com' ,
password : 'hunter2' ,
tokens : [ 'ghp_' + 'b' . repeat ( 36 ) ] ,
note : 'my aws key is AKIAIOSFODNN7EXAMPLE' ,
} ) ;
// →
// {
// user: 'b***@***',
// password: '***', // sensitive field name
// tokens: ['ghp_***'],
// note: 'my aws key is AKIA***', // found inside free text
// }
Redact prompts before they reach an LLM
Your app sends user data to OpenAI or Anthropic. Somewhere in that prompt is a
customer's email, an API key, or a card number — and now it's left your systems.
Wrap the client once, and detected secrets are stripped from prompts and put back
in the reply. The model never sees those original values; your code keeps the
references it needs.
import { wrapOpenAI } from 'flare-redact/llm' ;
const openai = wrapOpenAI ( new OpenAI ( ) ) ;
const res = await openai . chat . completions . create ( {
model : 'gpt-4o' ,
messages : [ { role : 'user' , content : 'Email the invoice to alice@corp.com, card 4242 4242 4242 4242' } ] ,
} ) ;
your app sends → Email the invoice to alice@corp.com, card 4242 4242 4242 4242
the model sees → Email the invoice to [FR_EMAIL_7f2a…], card [FR_CREDIT_CARD_19be…]
your app gets → Sent to alice@corp.com. Card 4242 4242 4242 4242 wasn't stored.
wrapAnthropic does the same for messages.create , including the system prompt.
Both handle streaming — placeholders are restored even when one is split across
chunks. There's a redactPrompt(text) too if you'd rather hold the vault
yourself.
Pick a mode depending on whether you still need to reason about the data
after it's hidden.
redact ( 'bob@corp.com' , { mode : 'mask' } ) ; // 'b***@***' (default)
redact ( 'bob@corp.com' , { mode : 'label' } ) ; // '[REDACTED:email]'
const protectedOptions = { transformSecret : process . env . FLARE_REDACT_SECRET } ;
redact ( 'bob@corp.com' , { ... protectedOptions , mode : 'hash' } ) ;
// 'email_3baf4d28d7c88317a…' — HMAC-SHA-256 fingerprint
redact ( 'bob@corp.com' , { ... protectedOptions , mode : 'pseudonym' } ) ;
// 'kqz@rwmp.dnu' — keyed, deterministic, keeps character classes
redact ( 'bob@corp.com' , { ... protectedOptions , mode : 'surrogate' } ) ;
// 'user_93a78c61e204@example.invalid' — type-consistent synthetic value
Protected deterministic modes require transformSecret ; they never silently
fall back to a public unsalted fingerprint. hash is useful for correlation,
pseudonym retains the original character shape, and surrogate emits typed
synthetic values such as reserved-domain emails and Luhn-valid card numbers.
Use a separate secret per environment or correlation domain.
pseudonym is deliberately not described as format-preserving encryption.
It is non-reversible pseudonymization, not NIST FF1. The old fpe name remains
as a compatibility alias but is deprecated.
Or replace everything with one fixed string:
redact ( payload , { mask : '█' } ) ;
redact ( payload , { mask : ( { detector } ) => `< ${ detector . id } >` } ) ;
Reversible redaction
When you need the originals back — the LLM case above, or handing data to a
system you don't trust and getting it back — use a vault. It swaps each secret
for a stable placeholder and remembers the mapping.
import { createVault } from 'flare-redact' ;
const vault = createVault ( ) ;
const safe = vault . redact ( 'charge bob@corp.com on card 4242 4242 4242 4242' ) ;
// 'charge [FR_EMAIL_7f2ad4…] on card [FR_CREDIT_CARD_19be63…]'
vault . restore ( safe ) ;
// 'charge bob@corp.com on card 4242 4242 4242 4242'
The same value gets the same placeholder inside one vault, so references survive
the round trip. Default placeholders include 96 random bits instead of a global
sequence number. Human-readable [EMAIL_1] counters remain available through
createVault({ placeholderStyle: 'readable' }) for trusted local workflows.
The mapping is as sensitive as the original data. Encrypt it before persistence:
import { sealVault , openVault , restore } from 'flare-redact' ;
const encrypted = await sealVault ( vault , process . env . FLARE_REDACT_VAULT_PASSWORD ) ;
await fs . writeFile ( 'session.vault.json' , JSON . stringify ( encrypted ) , { mode : 0o600 } ) ;
const entries = await openVault ( encrypted , process . env . FLARE_REDACT_VAULT_PASSWORD ) ;
restore ( safe , new Map ( entries ) ) ;
Sealed vaults use PBKDF2-SHA-256 with a fresh salt and AES-256-GCM with a fresh
nonce. Wrong passwords and modified files fail closed.
From the CLI, --vault and --restore use encrypted files by default. Passwords
come from FLARE_REDACT_VAULT_PASSWORD (or the variable named by
--vault-password-env ) so they do not appear in shell history:
export FLARE_REDACT_VAULT_PASSWORD= ' use-a-secret-manager-in-production '
flare-redact --vault session.vault.json < input.txt > safe.txt
flare-redact --restore session.vault.json < safe.txt > restored.txt
Contextual and model-assisted PII
Structured identifiers and credentials are best handled by deterministic rules
and checksum validators. Names and addresses need context, so three conservative
detectors are opt-in:
const findings = scan (
'Customer name: Alice Example; address: 120 Cedar Street; DOB: 1990-04-23' ,
{ enable : [ 'contextual' ] } ,
) ;
// person_name, street_address, date_of_birth
// each finding includes risk, confidence, and the exact sensitive span
For broader multilingual free-text PII, connect a local model without coupling
the zero-dependency core to one ML runtime:
const policy = {
semanticProvider : {
async detect ( text ) {
return [ {
detector : 'person_model' , label : 'Person' ,
why : 'Local multilingual NER result.' ,
start : 12 , end : 25 , confidence : 0.94 , risk : 'high' ,
} ] ;
} ,
} ,
minConfidence : 0.8 ,
} ;
const safe = await redactAsync ( input , policy ) ;
Semantic and deterministic spans enter the same overlap arbitration. Higher-risk,
higher-priority, and better-validated findings win instead of whichever regular
expression happens to run first.
If you're building a chat interface — over your own local model or any API — a
session is the drop-in layer. One session holds one vault, so a value keeps
the same placeholder across every turn: mask the user's message on the way in,
restore the model's reply on the way out. It's model-agnostic and synchronous.
Run npm run benchmark on your own target runtime instead of relying on a
hardware-independent latency claim.
import { createSession } from 'flare-redact' ;
const session = createSession ( { enable : [ 'pii' ] } ) ;
// on the way in — the model only ever sees placeholders
const prompt = session . redact ( userMessage ) ;
const reply = await myModel . generate ( prompt ) ;
// on the way out — the user sees the real values back
show ( session . restore ( reply ) ) ;
Streaming? Restore token by token, even when a placeholder is split across chunks:
const out = session . stream ( ) ;
for await ( const chunk of modelStream ) process ( out . push ( chunk . text ) ) ;
process ( out . flush ( ) ) ;
session.redactMessages([{ role, content }]) masks a whole chat array at once,
and session.reset() starts a fresh conversation. Detected original values stay
local while your app keeps a reversible reference.
Detectors can't know your product codenames, project names, or internal jargon —
so hand them a list. terms catches exactly the words you name (any language,
longest match first, word-boundary safe), one-way or reversibly.
// one-way, with your own replacement text
redact ( 'Launch Project Zeus with Falcon' , {
terms : { 'Project Zeus' : '[CLASSIFIED]' , 'Falcon' : '[CLASSIFIED]' } ,
} ) ;
// → 'Launch [CLASSIFIED] with [CLASSIFIED]'
// reversible — send to a model, get it back

[truncated]
