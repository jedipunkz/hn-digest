---
source: "https://ar-agents.ar"
hn_url: "https://news.ycombinator.com/item?id=48506865"
title: "Show HN: Open infrastructure for Argentina's AI-run companies (36 NPM packages)"
article_title: "ar-agents · open infrastructure for AI corporations in Argentina"
author: "naza0"
captured_at: "2026-06-12T18:46:31Z"
capture_tool: "hn-digest"
hn_id: 48506865
score: 1
comments: 0
posted_at: "2026-06-12T17:23:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open infrastructure for Argentina's AI-run companies (36 NPM packages)

- HN: [48506865](https://news.ycombinator.com/item?id=48506865)
- Source: [ar-agents.ar](https://ar-agents.ar)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T17:23:03Z

## Translation

タイトル: Show HN: アルゼンチンの AI 経営企業向けのオープン インフラストラクチャ (36 NPM パッケージ)
記事のタイトル: ar-agents · アルゼンチンの AI 企業向けオープン インフラストラクチャ
説明: アルゼンチンの AI 企業向けのオープンソース インフラストラクチャ (sociedades-IA)。 36 個の npm パッケージ、6 個の RFC、パブリック レジストリ、HMAC + Ed25519 監査ログ。 MIT + CC-BY-4.0。

記事本文:
ar-agents · アルゼンチンの AI 企業向けオープン インフラストラクチャ ar-agents の実装
ES JP オープン インフラストラクチャ · MIT + CC-BY-4.0
AI社会のためのオープンインフラ。
アルゼンチン製。
自動化社会は従業員ではなく AI エージェントによって運営されます。 ar-agents は、それを構築および運用するためのオープンソースです。
2026 年 6 月 1 日に上院に送付される一般会社法の暫定草案。まだ法律ではありません。
GitHub 開発者向け Vercel AI SDK 6 用に入力された 36 npm パッケージ。請求、CUIT の検証、WhatsApp の送信、請求書の発行、BO の監視。
HMAC + Ed25519 を使用したフォレンジック監査ログ。オペレーターにキーを要求せずに検証可能。 1 ページャーを印刷可能。
6 つの RFC を技術的に統合し、推奨テキストを参照により引用します。法律を書いている人にとって。
↳ エージェント Claude Sonnet 4.6 ツール呼び出し @ar-agents/* 36 個のパッケージ 235 個の型付きツール Vercel AI SDK 6 AFIP / ARCA Mercado Pago WhatsApp Biz 官報 IGJ + BCRA ↓ 各ツール呼び出しは二重署名監査ログ (HMAC + Ed25519) に書き込まれます npm のパッケージ 33 個 6 RFC (CC-BY-4.0) 221 種類のツール 5 つの実装 100/100 コンプライアンス モデル · オープンコア
標準は無料です。信頼こそがビジネスなのです。
コードはオープンで無料です。トラストは有料サービスです。自社でテストしてみました。
ソブリン統合 (AFIP、Mercado Pago、IGJ、WhatsApp、請求書) と会社を構成するウィザード。 MIT+CC-BY-4.0。
それぞれの決定は署名されたログ (HMAC + Ed25519) に残され、誰でも検証できます。芸術。 102は人間に責任を課します。このログはあなたの防御策です。
ar-agents は自動化された会社として設立され、独自の監査人を使用しています。試験は公開です。
ツールキット: スタックの各部分に 1 つのパッケージ
アルゼンチンのビジネスが必要とするスタックのすべての部分をパックとして

npm では年齢に依存しません。それらは相互に構成され、ロゴ @ar-agents/mercadopago を使用して構成されます。
CUIT/CUIL の検証 + AFIP/ARCA レジストリの検索 (モノトリビュート + VAT 条件による一貫性)。サブパス経由の WSAA SOAP。
OIDC of My Argentina (アルゼンチン政府の ID でログイン)。 PKCE、RS256 トークン ID 検証、JWKS キャッシュ、リフレッシュ、セッション終了。 Web Crypto のみ、Edge 上で実行されます。
検証オーケストレーター (WhatsApp OTP、電子メールのマジックリンク、Auth0、Magic.link、MP Identity)。 HMAC と信頼レベルを使用して署名された証明書を返します。
WhatsApp ビジネスクラウド API。 Webhook + HMAC 検証。 ARフォンノーマライザー。 ScopedTo モード: 出力ツールを単一の送信者にバインドします。
AFIP/ARCA 電子請求書 (WSFE)。請求書 A/B/C、NC/ND、FCE MiPyME。ラウンドトリップの前に、最も一般的な 10 個の拒否理由を検出するローカルのプリフライト検証ツール。
銀行/PSP ID + BCRA Debtors Central + 主要 BCRA 変数 (公式 USD、CER、UVA、準備金、BADLAR、インフレ) による CBU/CVU の検証。ツールは全部で11個。
構造化消防ホースとしての官報: 検索、セクションによるフィルター、ID による標準の取得、キーワード/CUIT/組織による購読。 ARエコシステムに欠けていた「法的監視のためのVercel」。
General Inspection of Justice (IGJ) オープンデータ: エンティティ、住所、当局、貸借対照表、議会の検索。パブリック CKAN を data.jus.gob.ar でラップします。データセットはサンプリングされており (リアルタイムではありません)、「coverageNote」は各結果とともに移動します。
アルゼンチンのデジタル署名 (法律 25,506 / ONTI): 検証のみを解析し、実際の署名には物理トークンが必要です。
Andreani (REST 完了)、OCA、Corre

またはアルゼンチン人。見積/作成/追跡/キャンセル。地方ヘルパー + CPA。
TAD (リモート手順) + GDE (電子文書管理): 確立された電子アドレスからの通知の取り込み、手順の追跡、IGJ 登録のプリフライト。 AI 協会の 4 番目の部分、RFC-001 § 3.4。まだ文書化されたパブリック API はありません。アダプタ パターン + スクレイピングによるデフォルト。
Mercado Libre マーケットプレイス用の SDK および独立したツール コレクション (アイテム、カテゴリ、質問、注文、パック、苦情、出荷、評判、プロモーション、Webhook)。 OAuth 結合、/myfeeds リプレイ、不可逆操作の HITL ゲート。コミュニティで構築されており、所属はありません。
@ar-agents/エージェント-コマース-ブリッジ
ChatGPT、Claude、Gemini、およびその他のエージェント コマース クライアントと MercadoLibre + MercadoPago の橋渡しをする、Agentic Commerce Protocol (ACP) のオープンソース マーチャント ファシリテーター。 ACP チェックアウト セッション、署名された Webhook、discovery .well-known/acp.json、AFIP A/B/C 請求書の自動発行。中南米初の代理店と商取引の架け橋。
Agent Payments Protocol (AP2) v0.2 の最初の忠実な TypeScript 実装、スキーマ、ES256 SD-JWT VC コマンド (チェックアウト + 支払い、オープン + クローズ)、8 つの制約エバリュエーター、署名済みチェックアウト/支払い領収書。エッジランタイム互換。 FIDO Alliance Agentic Auth WG Reference Python SDK と連携しています。
「ar-agents.ar/api/auto-incorporate」の依存関係のない TypeScript クライアントにより、外部エージェント (USA-LLC、ChatGPT、Claude、Gemini) が 1 回の呼び出しでアルゼンチンの AI 企業を自動的に組み込むことができます。生成されたファイル、Vercel デプロイ URL、および署名された監査ログ参照を返します。
13 個のパッケージとツールをバンドルする MCP サーバー。 Claude デスクトップ / カーソル / 任意の場所に 1 回インストール

えーMCPホスト。有効にするパッケージを環境変数から自動検出します。
構成例：コレクションアシスタント
whatsapp-hello.ar-agents.ar は、MP が単一エージェント内で ID、ID-Attest、および whatsapp で構成されていることを示しています。 ARCA に対して CUIT を検証し、検証 (WhatsApp OTP) で多額の料金を徴収し、MP サブスクリプションを作成して、WhatsApp 経由で応答します。
オープンソースコア・フラッグシップパッケージ
Vercel AI SDK 6 用に入力された 89 のツール。決定論的冪等性、不可逆操作のプログラムによる HITL、HMAC Webhook 検証。オープンソース コア (36 パッケージ) の中で最も成熟した部分であり、残りのツールキットによって参照されます。コアは無料です。有料トラスト層は監査人であり、私たちはそれをドッグフードしています。私たちは自らを自動化社会として構成しています。
OAuth マーケットプレイスの作成 / キャプチャ / 返金 Checkout Pro 注文管理
作成/取得/一時停止/再開/キャンセル・計画・保存されたカード
AR プロモーション カタログ、割り当て、3DS チャレンジの解決
ローカル QR · 物理的なポイント デバイス · 店舗 + POS
HMAC 検証、リプレイ ウィンドウ、重複排除、handle_webhook コンボ
InMemory + Vercel KV アダプター対応、プラグイン可能なインターフェース
サブパス経由の OpenTelemetry トレース、監査ログ アダプター、サーキット ブレーカー
デフォルトでの決定論的冪等性 · 8 つの不可逆操作でのプログラマティック HITL
pnpm 追加 @ar-agents/mercadopago ai zod
import { Experimental_Agent as Agent, stepCountIs } from "ai";
インポート {
メルカドパゴクライアント、
マーケットペイツール、
メモリ状態アダプター内、
"@ar-agents/mercadopago" から;
const mp = new MercadoPagoClient({
accessToken: process.env.MP_ACCESS_TOKEN!, // TEST- サンドボックスの場合、APP_USR- 製品の場合
});
const エージェント = 新しいエージェント({
モデル: "anthropic/claude-sonnet-4-6",
ツール:marketPayTools(mp, {

state: new InMemoryStateAdapter(), // prod の VercelKVStateAdapter と交換します
backUrl: "https://yoursite.com/subscription/done",
})、
stopWhen: stepCountIs(8)、
});
const { text } = await エージェント.generate({
プロンプト: "customer@example.com に対して $1000 ARS の月額サブスクリプションを作成します。",
});よくある質問
公式 mercadopago SDK は基本的な REST クライアントです。 ar-agents は、Vercel AI SDK 6 ツール スキーマ、決定論的冪等性キー (LLM 再試行に耐性)、5 分間のリプレイ ウィンドウによる HMAC Webhook 検証、8 つの不可逆操作でのプログラムによる HITL、Web Crypto による Edge Runtime サポート、Vercel KV アダプター、および OpenTelemetry インストルメンテーションを追加します。同じプロジェクトで両方のパッケージを使用できます。
はい。パッケージ全体は Web Crypto を使用し、node:crypto には依存しません。 Vercel Edge、Cloudflare Workers、Deno、および任意の V8 分離ランタイムで実行します。 Webhook、HMAC、冪等キーの検証には Web Crypto API を使用します。
8 つのツールはステータスを不可逆的に変更します (refund_payment、cancel_subscription、delete_customer_card など)。ツールキットは、コードが true を返すまで実行をブロックする requireconfirmation コールバックを受け入れます。これは、LLM への単なる指示ではなく、プログラムによるゲートです。
4 つのツール (create_payment、create_subscription、create_payment_preference、refund_payment) は、関連する入力の SHA-256 ハッシュからべき等キーを導出します。同じ入力 → 同じキー → MP がサーバー側で重複を排除します。ツールを再試行する LLM は、2 回請求するのではなく、既存のリソースを返します。
はい。 MITライセンス。有料枠、テレメトリ、使用制限はありません。公開されたすべての tarball には、来歴 npm SLSA v1 証明書が含まれています。
サイドカー パッケージはアルゼンチン スタックの残りの部分をカバーします: @ar-agents/identity (CUIT + AFIP/ARCA レジストリ)

monotributo + VAT)、@ar-agents/facturacion (WSFE 電子請求書)、@ar-agents/whatsapp (Business Cloud API)、@ar-agents/banking (CBU/CVU + BCRA Central de Deudores)、@ar-agents/shipping (Andreani/OCA/Correo Argentino)、@ar-agents/identity-attest (HMAC 検証オーケストレーター)。それぞれ独立して出版されています。
はい。 @ar-agents/mcp は、13 個のツール パッケージすべてを、Claude Desktop、Cursor、Codeium、Continue、Cline と互換性のある単一の MCP サーバーにバンドルします。 Glama および公式 MCP レジストリに io.github.ar-agents/mcp としてリストされています。

## Original Extract

Open-source infrastructure for Argentine AI corporations (sociedades-IA). 36 npm packages, 6 RFCs, a public registry, an HMAC + Ed25519 audit log. MIT + CC-BY-4.0.

ar-agents · open infrastructure for AI corporations in Argentina ar-agents Implementación
ES EN Infraestructura abierta · MIT + CC-BY-4.0
Infraestructura abierta para sociedades de IA.
Hecha en Argentina.
Una Sociedad Automatizada opera con agentes de IA, no con empleados. ar-agents es el código abierto para constituirla y operarla.
Anteproyecto de Ley General de Sociedades enviado al Senado el 1-jun-2026. Todavía no es ley.
GitHub Para developers 36 paquetes npm tipados para Vercel AI SDK 6. Cobrar, validar CUIT, mandar WhatsApp, emitir factura, monitorear el BO.
Audit log forense con HMAC + Ed25519, verificable sin pedir la clave al operador. 1-pager imprimible.
Síntesis técnica de los 6 RFCs con texto sugerido cite-by-reference. Para quien esté redactando la ley.
↳ agente Claude Sonnet 4.6 tool-calls @ar-agents/* 36 paquetes 235 tools tipadas Vercel AI SDK 6 AFIP / ARCA Mercado Pago WhatsApp Biz Boletín Oficial IGJ + BCRA ↓ cada tool call se escribe al audit log dual-sign (HMAC + Ed25519) 33 Paquetes en npm 6 RFCs (CC-BY-4.0) 221 Tools tipadas 5 Implementaciones 100/100 Conformidad El modelo · open-core
El estándar es gratis. La confianza es el negocio.
El código es abierto y gratis. La confianza es un servicio pago. Lo probamos con nuestra propia empresa.
Las integraciones soberanas (AFIP, Mercado Pago, IGJ, WhatsApp, factura) y el wizard que constituye la sociedad. MIT + CC-BY-4.0.
Cada decisión queda en un log firmado (HMAC + Ed25519) que cualquiera puede verificar. El art. 102 hace responsable al humano. Este log es su defensa.
ar-agents se constituyó como Sociedad Automatizada y usa su propio Auditor. La prueba es pública.
El toolkit: un paquete por pieza del stack
Cada pieza del stack que un negocio argentino necesita, como package independiente en npm. Componen entre sí y con el insignia, @ar-agents/mercadopago .
Validación de CUIT/CUIL + lookup en padrón AFIP/ARCA (constancia con monotributo + condición IVA). WSAA SOAP vía subpath.
OIDC de Mi Argentina (login con la identidad del gobierno argentino). PKCE, verificación RS256 del ID token, JWKS cacheado, refresh, end-session. Solo Web Crypto, corre en Edge.
Orquestrador de verificación (WhatsApp OTP, email magic-link, Auth0, Magic.link, MP Identity). Devuelve attestation firmada con HMAC y un trust level.
WhatsApp Business Cloud API. Webhook + verificación HMAC. Normalizador de teléfonos AR. Modo scopedTo: bindea las tools de salida a un solo sender.
Factura electrónica AFIP/ARCA (WSFE). Factura A/B/C, NC/ND, FCE MiPyMEs. Validador local pre-flight que atrapa los 10 motivos de rechazo más comunes antes del round-trip.
Validación de CBU/CVU con identificación de banco/PSP + Central de Deudores BCRA + Principales Variables BCRA (USD oficial, CER, UVA, reservas, BADLAR, inflación). 11 tools en total.
Boletín Oficial como firehose estructurado: búsqueda, filtro por sección, obtener norma por id, suscripciones por keyword/CUIT/organismo. El 'Vercel for legal monitoring' que faltaba en el ecosistema AR.
Inspección General de Justicia (IGJ) datos abiertos: búsqueda de entidades, domicilios / autoridades / balances / asambleas. Wrappea el CKAN público en datos.jus.gob.ar. Dataset es muestreo (no real-time), `coverageNote` viaja con cada resultado.
Firma Digital argentina (Ley 25.506 / ONTI): parsea certs X.509, verifica cadenas ancladas en AC-Raíz Argentina, verifica firmas CMS / PKCS#7 desligadas, extrae CUIT del subject del firmante. Sólo verificación, la firma real requiere token físico.
Andreani (REST completo), OCA, Correo Argentino. cotizar / crear / trackear / cancelar. Helpers de provincia + CPA.
TAD (Trámites a Distancia) + GDE (Gestión Documental Electrónica): ingestión de notificaciones del Domicilio Electrónico Constituido, tracking de trámites, pre-flight de inscripciones IGJ. La 4ta pieza de las sociedades-IA, RFC-001 § 3.4. Sin API pública documentada todavía; adapter pattern + defaults vía scrape.
SDK y tool collection independiente para el marketplace de Mercado Libre (items, categorías, preguntas, órdenes, packs, reclamos, shipments, reputación, promociones, webhooks). OAuth coalescing, replay de /myfeeds, gates HITL en operaciones irreversibles. Community-built, sin afiliación.
@ar-agents/agentic-commerce-bridge
Merchant facilitator open-source para el Agentic Commerce Protocol (ACP) que puentea ChatGPT, Claude, Gemini y otros clientes agentic-commerce con MercadoLibre + MercadoPago. Sesiones de checkout ACP, webhooks firmados, discovery .well-known/acp.json, emisión automática de Factura A/B/C AFIP. Primer bridge agentic-commerce de LATAM.
Primera implementación TypeScript fiel del Agent Payments Protocol (AP2) v0.2, schemas, mandatos ES256 SD-JWT VC (Checkout + Payment, open + closed), 8 evaluadores de constraints, recibos Checkout/Payment firmados. Edge-Runtime-compatible. Alineado con el SDK Python de referencia del FIDO Alliance Agentic Auth WG.
Cliente TypeScript zero-dependency para `ar-agents.ar/api/auto-incorporate`, permite que un agente externo (USA-LLC, ChatGPT, Claude, Gemini) auto-incorpore una sociedad-IA argentina en una sola llamada. Devuelve los archivos generados, la URL de deploy en Vercel y la referencia firmada del audit-log.
Servidor MCP que bundlea los 13 packages con tools. Una sola instalación en Claude Desktop / Cursor / cualquier host MCP. Auto-detecta qué packages habilitar a partir de env vars.
Ejemplo de composición: asistente de cobros
whatsapp-hello.ar-agents.ar muestra MP componiéndose con identity, identity-attest y whatsapp en un solo agente. Valida el CUIT contra ARCA, gatea cobros grandes con verificación (WhatsApp OTP), crea la suscripción de MP y responde por WhatsApp.
Núcleo open-source · paquete insignia
89 tools tipadas para Vercel AI SDK 6. Idempotencia determinística, HITL programático en operaciones irreversibles, verificación de webhook HMAC. La pieza más madura del núcleo open-source (36 paquetes), referenciada por el resto del toolkit. El núcleo es gratis; la capa de confianza paga es El Auditor , y la dogfoodeamos: nos constituimos a nosotros mismos como Sociedad Automatizada.
create / capture / refund · OAuth marketplace · Checkout Pro · Order Management
create / get / pause / resume / cancel · planes · tarjetas guardadas
Catálogo de promos AR · cuotas · resolución de challenge 3DS
QR en local · dispositivos Point físicos · Stores + POS
Verificación HMAC · ventana de replay · deduplicación · handle_webhook combo
Adapters InMemory + Vercel KV listos · interfaz pluggable
Trazas OpenTelemetry vía subpath · adapter de audit log · circuit breaker
idempotencia determinística por default · HITL programático en 8 ops irreversibles
pnpm add @ar-agents/mercadopago ai zod
import { Experimental_Agent as Agent, stepCountIs } from "ai";
import {
MercadoPagoClient,
mercadoPagoTools,
InMemoryStateAdapter,
} from "@ar-agents/mercadopago";
const mp = new MercadoPagoClient({
accessToken: process.env.MP_ACCESS_TOKEN!, // TEST- for sandbox, APP_USR- for prod
});
const agent = new Agent({
model: "anthropic/claude-sonnet-4-6",
tools: mercadoPagoTools(mp, {
state: new InMemoryStateAdapter(), // swap for VercelKVStateAdapter in prod
backUrl: "https://yoursite.com/subscription/done",
}),
stopWhen: stepCountIs(8),
});
const { text } = await agent.generate({
prompt: "Creá una subscription mensual de $1000 ARS para customer@example.com.",
}); Preguntas frecuentes
El SDK oficial mercadopago es un cliente REST básico. ar-agents agrega tool schemas del Vercel AI SDK 6, idempotency keys determinísticas (resistentes a retry de LLM), verificación de webhook HMAC con ventana de replay de 5 minutos, HITL programático en 8 operaciones irreversibles, soporte Edge Runtime via Web Crypto, adapters de Vercel KV, e instrumentación OpenTelemetry. Podés usar ambos packages en el mismo proyecto.
Sí. Todo el package usa Web Crypto, sin dependencia de node:crypto. Corre en Vercel Edge, Cloudflare Workers, Deno, y cualquier runtime V8-isolate. Verificación de webhooks, HMAC, e idempotency keys usan la Web Crypto API.
8 tools modifican estado irreversiblemente (refund_payment, cancel_subscription, delete_customer_card, etc). El toolkit acepta un callback requireConfirmation que bloquea la ejecución hasta que tu código devuelva true. Es un gate programático, no solo instrucciones al LLM.
Cuatro tools (create_payment, create_subscription, create_payment_preference, refund_payment) derivan la idempotency key de un hash SHA-256 de los inputs relevantes. Mismos inputs → misma key → MP deduplica server-side. Un LLM que reintenta un tool devuelve el recurso existente en vez de cobrar dos veces.
Sí. Licencia MIT. Sin tier pago, sin telemetría, sin límites de uso. Todos los tarballs publicados llevan attestation de provenance npm SLSA v1.
Packages sidecar cubren el resto del stack argentino: @ar-agents/identity (CUIT + padrón AFIP/ARCA con monotributo + IVA), @ar-agents/facturacion (factura electrónica WSFE), @ar-agents/whatsapp (Business Cloud API), @ar-agents/banking (CBU/CVU + BCRA Central de Deudores), @ar-agents/shipping (Andreani/OCA/Correo Argentino), @ar-agents/identity-attest (orquestador de verificación HMAC). Cada uno se publica independiente.
Sí. @ar-agents/mcp bundlea los 13 packages con tools en un único servidor MCP compatible con Claude Desktop, Cursor, Codeium, Continue, Cline. Listado en Glama y en el MCP Registry oficial como io.github.ar-agents/mcp.
