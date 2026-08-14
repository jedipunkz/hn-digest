---
source: "https://trpevski.com/blog/openrouter-unified-llm-api-with-routing-and-fallbacks/"
hn_url: "https://news.ycombinator.com/item?id=49296142"
title: "OpenRouter: Unified LLM API with Routing and Fallbacks"
article_title: "OpenRouter: Unified LLM API with Routing and Fallbacks — Darko Trpevski"
author: "dzugumot"
captured_at: "2026-08-14T08:58:15Z"
capture_tool: "hn-digest"
hn_id: 49296142
score: 1
comments: 0
posted_at: "2026-08-14T08:46:01Z"
tags:
  - hacker-news
  - translated
---

# OpenRouter: Unified LLM API with Routing and Fallbacks

- HN: [49296142](https://news.ycombinator.com/item?id=49296142)
- Source: [trpevski.com](https://trpevski.com/blog/openrouter-unified-llm-api-with-routing-and-fallbacks/)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T08:46:01Z

## Translation

タイトル: OpenRouter: ルーティングとフォールバックを備えた統合 LLM API
記事のタイトル: OpenRouter: ルーティングとフォールバックを備えた統合 LLM API — Darko Trpevski
説明: コストが最適化されたルーティング、フォールバック チェーン、およびキーごとのセキュリティ制御を備えた、複数の LLM プロバイダーにわたる統合プロキシ層。

記事本文:
OpenRouter: ルーティングとフォールバックを備えた統合 LLM API — Darko Trpevski
OpenRouter: ルーティングとフォールバックを備えた統合 LLM API
コストが最適化されたルーティング、フォールバック チェーン、およびキーごとのセキュリティ制御を備えた、複数の LLM プロバイダーにわたる統合プロキシ レイヤー。
あなたは、複数の AI 機能を使用するアプリを構築しています。 OpenAI の Whisper を使用した音声からテキストへの変換。クロードとの文章推理。 Llama または Mistral のオープンソース モデルを使用したいくつかの実験的な機能。各プロバイダーには独自の API キーが必要です。それぞれに異なる API 形式、異なるエラー処理、異なるレート制限があります。
コードベースには、プロバイダーごとに個別のコード パスが含まれ始めます。 .env ファイルには 6 つのキーがあります。 CI/CD パイプラインは、OpenAI、Anthropic、Togetter AI のシークレットを個別に渡します。モバイル アプリには API キーが付属しています。誰かがそれを抽出します。使用しなかったサービスに対して 1,200 ドルの請求書が届きます。キーを取り消し、新しいキーを追加し、アプリのアップデートを配布します。
来月、Docker から別のキーが漏洩します。同じ問題です。
より良いアプローチが必要であることに気づきました。 API キーを 1 か所で管理します。任意のモデルを呼び出す 1 つの方法。すべての支出を 1 つのダッシュボードで確認できます。
それが Open Router が解決する本当の問題です。
Open Router は、アプリケーションと複数の LLM プロバイダー (OpenAI、Anthropic、Meta など) の間に位置する API プロキシ サービスです。プロバイダーごとに個別の API キーと統合を管理する代わりに、1 つの Open Router API キーを使用して、さまざまなプロバイダーにわたる 300 以上のモデルにアクセスします。
このサービスは、すべてのプロバイダーにわたって API インターフェイスを標準化するため、パラメーターや応答処理を変更せずに、同じコードで Claude、GPT-4、または Llama を呼び出すことができます。
複数の API キーと SDK : LLM を使用して構築する場合、多くの場合、複数のプロバイダーが必要になります。音声には OpenAI、推論にはクロード、実験にはラマ。それぞれに個別の SDK が必要です。

個別の資格情報、個別のエラー処理。
キーのセキュリティ : モバイル アプリ、.env ファイル、および Docker 構成の API キーは脆弱です。 OpenAI キーが盗まれると、数分で数千ドルの費用がかかる可能性があります。複数のキーを管理すると、攻撃対象領域が増大します。
プロバイダーのロックイン : 1 つのプロバイダーの API に基づいて構築すると、切り替えにエンジニアリング時間がかかります。 Open Router を使用すると、コードを変更せずにモデルを交換できます。
一貫性のないインターフェイス : プロバイダーが異なれば、異なる応答形式が返され、エラー コードも異なり、レート制限の動作も異なります。 Open Router はこれを正規化します。
統合された監視がない : 複数のプロバイダーを使用すると、支出と使用量が異なるダッシュボードに分散されます。 LLM コストの合計を確認するのは困難です。
ダイレクト プロバイダー API (OpenAI、Anthropic、Togetter AI など)
プロバイダーに直接電話する
長所: 遅延の最小化、ボリュームディスカウント、フルコントロール
短所: 複数の統合、フォールバックなし、複数のキーを管理する必要がある
複数のプロバイダーをラップする AWS サービス
長所: エンタープライズ機能、AWS 統合
短所: AWS のロックイン、レイテンシの高さ、直接よりも高価
LLM API を統合するオープンソース プロキシ
長所: 自己ホスト型、ベンダー依存なし、無料
短所: 自己ホスト型の負担、管理されたフォールオーバーなし、ダッシュボードなし
LLM API の上のフレームワーク層
長所: 豊富な抽象化、エージェントのサポート、コミュニティ
短所: プロキシではない (生の API キーが必要)、フレームワークによる遅延の追加、複雑なデプロイメント
複数のオープンソース モデルの統一インターフェイス
長所: オープンソースのみのワークロードに適しています
短所: オープンモデルに限定されており、GPT-4 や Claude のような独自モデルはありません。
さまざまなモデルの推論プラットフォーム
長所: 使いやすく、生産推論に適しています
短所: モデルの選択が限られているため、画像/ビデオ モデルに重点が置かれます。
1 つの API キー: 1 つの認証情報で 300 以上のモデルにアクセスします。実際のプロバイダー キーはバックエンドのみに残ります。
プロバイダーインデックス

pendence : コードを変更せずにモデルを切り替えます。フォールバック チェーンを自動的に追加します。
管理されたインフラストラクチャ: セルフホスティングの負担がありません。 Open Router は、スケーリング、稼働時間、レート制限を処理します。
リアルタイム監視 : 単一のダッシュボードに、モデル別の支出、API 使用パターン、コスト アラートが表示されます。
組み込みフォールバック : プライマリ プロバイダーがダウンしているか、レートが制限されている場合の自動モデル フォールバック。
標準化された応答: すべてのプロバイダーで同じ応答形式。プロバイダー固有のエラー処理は必要ありません。
簡単な統合 : エンドポイントとヘッダーを簡単に変更することで、既存のコードと連携します。
透明性のある価格設定 : 5.5% のマークアップ、隠れた料金や交渉は必要ありません。
コストのオーバーヘッド: すべてのリクエストで 5.5% のマークアップ。大規模になると、これはさらに増加し​​ます。
レイテンシー: ルーティング層により、リクエストごとに 100 ～ 200 ミリ秒の追加料金がかかります。
単一障害点: Open Router がダウンしている場合、すべての LLM 呼び出しは失敗します。
ボリュームディスカウントなし : プロバイダーと価格交渉できません。
制限されたレート制限 : レート制限はオープン ルーターのものであり、基礎となるプロバイダーのものではありません。
Open Router を使用しない直接 OpenAI API 呼び出し:
let apiKey = Bundle.main.infoDictionary?["OPENAI_API_KEY"] as?文字列
var request = URLRequest(url: URL(string: "https://api.openai.com/v1/audio/transcriptions")!)
request.setValue("Bearer \(apiKey)", forHTTPHeaderField: "認可")
let task = URLSession.shared.dataTask(with: request) { データ、応答、エラー
if let データ = データ {
let transscription = try JSONDecoder().decode(Transcription.self, from: data)
}
}
task.resume()
オープンルーターの場合:
let openRouterKey = Bundle.main.infoDictionary?["OPENROUTER_KEY"] として?文字列
var request = URLRequest(url: URL(string: "https://openrouter.ai/api/v1/audio/transcriptions")!)
request.setValue("Bearer \(openRouterKey)", forHTTPHeaderField: "認可")
request.setValue("https://yourdomain.com", forHTTPHeaderF)

フィールド: "HTTP リファラー")
let task = URLSession.shared.dataTask(with: request) { データ、応答、エラー
if let データ = データ {
let transscription = try JSONDecoder().decode(Transcription.self, from: data)
}
}
task.resume()
エンドポイント URL と API キーのみを変更します。応答形式は同じままです。
すべての LLM 呼び出しに対して 1 つの中央エンドポイントを作成します。
// サービス/llm-proxy.js
const router = require('express').Router();
const axios = require('axios');
const OPENROUTER_KEY = プロセス.env.OPENROUTER_API_KEY;
const APP_DOMAIN = プロセス.env.APP_DOMAIN;
// 転写エンドポイント
router.post('/transcribe', async (req, res) => {
const { audioUrl } = req.body;
{を試してください
const 応答 = await axios.post(
'https://openrouter.ai/api/v1/audio/transcriptions',
{ URL: audioUrl }、
{
ヘッダー: {
'認可': `ベアラ ${OPENROUTER_KEY}`,
「HTTP リファラー」: APP_DOMAIN
}
}
);
res.json(response.data);
} キャッチ (エラー) {
res.status(error.response?.status || 500).json({ error: error.message });
}
});
// チャット完了エンドポイント
router.post('/completions', async (req, res) => {
const { プロンプト、モデル = 'anthropic/claude-3.5-sonnet' } = req.body;
{を試してください
const 応答 = await axios.post(
'https://openrouter.ai/api/v1/chat/completions',
{
モデル: モデル、
メッセージ: [{ 役割: 'ユーザー'、コンテンツ: プロンプト }]、
最大トークン数: 1024
}、
{
ヘッダー: {
'認可': `ベアラ ${OPENROUTER_KEY}`,
「HTTP リファラー」: APP_DOMAIN
}
}
);
res.json(response.data);
} キャッチ (エラー) {
res.status(error.response?.status || 500).json({ error: error.message });
}
});
// 再試行を伴うフォールバック チェーン
router.post('/completions-with-fallback', async (req, res) => {
const {プロンプト} = req.body;
定数モデル = [
'人間/クロード-3.5-ソネット',
'人間/クロード-3-作品',
「openai/gpt-4-turbo」
];
const maxRetries = 3;
constbaseDelayMs = 1000;
for (モデルの定数モデル

s) {
for (試行 = 0;試行 < maxRetries;試行 ++) {
{を試してください
const 応答 = await axios.post(
'https://openrouter.ai/api/v1/chat/completions',
{
モデル: モデル、
メッセージ: [{ 役割: 'ユーザー'、コンテンツ: プロンプト }]、
最大トークン数: 1024
}、
{
ヘッダー: {
'認可': `ベアラ ${OPENROUTER_KEY}`,
「HTTP リファラー」: APP_DOMAIN
}
}
);
res.json(response.data) を返します。
} キャッチ (エラー) {
if (error.response?.status === 429 || error.response?.status === 503) {
const lateMs = BaseDelayMs * Math.pow(2, 試行);
新しい Promise(r => setTimeout(r,
[切り捨てられた]
1 つの Open Router キーを使用して .env を作成します。
OPENROUTER_API_KEY=your_key_here
APP_DOMAIN=https://yourdomain.com
実際のプロバイダー キーは AWS Secrets Manager などに保存され、バックエンド サービスからのみアクセスできます。
curl -X POST http://localhost:3000/llm/completions \
-H "コンテンツ タイプ: application/json" \
-d '{
"prompt": "量子コンピューティングについて説明します",
"モデル": "人間/クロード-3.5-ソネット"
}'
フォールバック チェーンを使用する場合:
curl -X POST http://localhost:3000/llm/completions-with-fallback \
-H "コンテンツ タイプ: application/json" \
-d '{
"prompt": "量子コンピューティングについて説明します"
}'
長所
集中キー管理 : 多数の API キーではなく 1 つの API キーがアプリケーションに表示されます。実際のプロバイダー キーはバックエンド インフラストラクチャに残ります。
単一制御点: すべての LLM 呼び出しは 1 つのエンドポイントを経由します。動作の監査、監視、変更が簡単です。
プロバイダーに依存しない : アプリのコードを変更せずにモデルを変更します。設定値を変更して、Claude から GPT-4 に切り替えます。
組み込みフォールバック: Open Router はモデル フォールバック チェーンをサポートします。 1 つのモデルが利用できない場合は、自動的に次のモデルが試行されます。
レート制限: Open Router はアカウントにレート制限を適用します。偶発的なコストの暴走を防ぎます。
リアルタイム監視: ダッシュボードには、モデルごと、1 日ごとの支出、API 使用パターンが表示されます。
より簡単な展開:

複数のプロバイダー キーの代わりに、CI/CD で管理する 1 セットの資格情報。
セルフホスティングなし: マネージド サービスで、インフラストラクチャの負担はありません。
5.5% のコストオーバーヘッド: Open Router は、プロバイダーの価格に 5.5% のマークアップを追加します。大規模な場合、これは重要です。
遅延ペナルティ: 追加のネットワーク ホップにより、リクエストごとに 100 ～ 200 ミリ秒が追加されます。最初のトークンのレイテンシーは、直接 API 呼び出しよりもわずかに遅くなります。
単一障害点: Open Router がダウンしているか速度が遅い場合、すべての LLM 呼び出しが影響を受けます。フォールバックとしてプロバイダーに直接アクセスすることはできません。
応答形式の一貫性 : プロバイダーが異なれば、返される応答形式もわずかに異なります。 Open Router はこれを正規化しますが、エッジケースが存在します。
レート制限の可視性 : レート制限は Open Router の制限であり、基礎となるプロバイダーの制限ではありません。制限は直接 API アクセスよりも低くなります。
ボリュームディスカウントの損失: Open Router を通じてプロバイダーとボリューム価格の交渉を行うことができません。常に標準料金とマークアップを支払います。
ベンダー ロックイン ライト: Open Router から切り替えるには、再び直接プロバイダー API を指すようにコードを変更する必要があります。
1 つのコードベースで複数の LLM プロバイダーを管理する
クライアント アプリには API アクセスが必要です (モバイル アプリ、ブラウザー)
重要なセキュリティと集中管理
フォールバック チェーンにより付加価値が得られます (プロバイダーの停止を回避)
応答形式の一貫性が重要です
ボリュームは 1 か月あたり 100 万リクエスト未満です
セルフホスティングを行わずに管理されたインフラストラクチャが必要
単一のプロバイダーを長期的に利用することを約束する
リクエスト量は 1 か月あたり 100 万件を超えています (コストが重要)
エンタープライズコンプライアンスと必要な機能
管理されたエンタープライズ エクスペリエンスに対する高いコストを受け入れる

## Original Extract

A unified proxy layer over multiple LLM providers with cost-optimised routing, fallback chains, and per-key security controls.

OpenRouter: Unified LLM API with Routing and Fallbacks — Darko Trpevski
OpenRouter: Unified LLM API with Routing and Fallbacks
A unified proxy layer over multiple LLM providers with cost-optimised routing, fallback chains, and per-key security controls.
You are building an app that uses multiple AI features. Voice-to-text with OpenAI's Whisper. Text reasoning with Claude. Some experimental features with open-source models from Llama or Mistral. Each provider needs its own API key. Each has different API format, different error handling, different rate limits.
Your codebase starts having separate code paths for each provider. Your .env file has six keys. Your CI/CD pipeline passes secrets for OpenAI, Anthropic, Together AI separately. Your mobile app ships an API key. Someone extracts it. You get a $1200 bill for a service you did not use. You revoke the key, add a new one, ship an app update.
Next month, a different key leaks from Docker. Same problem.
You realize you need a better approach. One place to manage API keys. One way to call any model. One dashboard to see all spending.
That is the real problem Open Router solves.
Open Router is an API proxy service that sits between your application and multiple LLM providers (OpenAI, Anthropic, Meta, etc.). Instead of managing separate API keys and integrations for each provider, you use one Open Router API key to access 300+ models across different providers.
The service standardizes the API interface across all providers, meaning the same code can call Claude, GPT-4, or Llama without changing parameters or response handling.
Multiple API keys and SDKs : When building with LLMs, you often need multiple providers. OpenAI for voice, Claude for reasoning, Llama for experiments. Each requires a separate SDK, separate credentials, separate error handling.
Key security : API keys in mobile apps, .env files, and Docker configs are vulnerable. A stolen OpenAI key can cost thousands in minutes. Managing multiple keys multiplies the attack surface.
Provider lock-in : Once you build on one provider's API, switching costs engineering time. Open Router lets you swap models without code changes.
Inconsistent interfaces : Different providers return different response formats, have different error codes, different rate limit behaviors. Open Router normalizes this.
No unified monitoring : When using multiple providers, spending and usage are scattered across different dashboards. Hard to see total LLM costs.
Direct Provider APIs (OpenAI, Anthropic, Together AI, etc.)
You call the provider directly
Pros: Lowest latency, volume discounts, full control
Cons: Multiple integrations, no fallback, multiple keys to manage
AWS service that wraps multiple providers
Pros: Enterprise features, AWS integration
Cons: AWS lock-in, higher latency, more expensive than direct
Open-source proxy that unifies LLM APIs
Pros: Self-hosted, no vendor dependency, free
Cons: Self-hosted burden, no managed fallover, no dashboard
Framework layer above LLM APIs
Pros: Rich abstractions, agent support, community
Cons: Not a proxy (still need raw API keys), adds latency through framework, complex deployment
Unified interface for multiple open-source models
Pros: Good for open-source only workloads
Cons: Limited to open models, no proprietary models like GPT-4 or Claude
Inference platform for various models
Pros: Easy to use, good for production inference
Cons: Limited model selection, focus on image/video models
One API key : Access 300+ models with single credentials. Real provider keys stay backend-only.
Provider independence : Switch models without code changes. Add fallback chains automatically.
Managed infrastructure : No self-hosting burden. Open Router handles scaling, uptime, rate limiting.
Real-time monitoring : Single dashboard shows spending by model, API usage patterns, cost alerts.
Built-in fallback : Automatic model fallback if primary provider is down or rate-limited.
Standardized responses : Same response format across all providers. No provider-specific error handling needed.
Easy integration : Works with existing code through simple endpoint and header changes.
Transparent pricing : 5.5% markup, no hidden fees or negotiation required.
Cost overhead : 5.5% markup on all requests. At scale, this adds up.
Latency : 100-200ms extra per request due to routing layer.
Single point of failure : If Open Router is down, all LLM calls fail.
No volume discounts : Cannot negotiate pricing with providers.
Limited rate limits : Rate limits are Open Router's, not the underlying provider's.
Direct OpenAI API call without Open Router:
let apiKey = Bundle.main.infoDictionary?["OPENAI_API_KEY"] as? String
var request = URLRequest(url: URL(string: "https://api.openai.com/v1/audio/transcriptions")!)
request.setValue("Bearer \(apiKey)", forHTTPHeaderField: "Authorization")
let task = URLSession.shared.dataTask(with: request) { data, response, error in
if let data = data {
let transcription = try JSONDecoder().decode(Transcription.self, from: data)
}
}
task.resume()
With Open Router:
let openRouterKey = Bundle.main.infoDictionary?["OPENROUTER_KEY"] as? String
var request = URLRequest(url: URL(string: "https://openrouter.ai/api/v1/audio/transcriptions")!)
request.setValue("Bearer \(openRouterKey)", forHTTPHeaderField: "Authorization")
request.setValue("https://yourdomain.com", forHTTPHeaderField: "HTTP-Referer")
let task = URLSession.shared.dataTask(with: request) { data, response, error in
if let data = data {
let transcription = try JSONDecoder().decode(Transcription.self, from: data)
}
}
task.resume()
Only change: endpoint URL and API key. Response format stays identical.
Create one central endpoint for all LLM calls:
// services/llm-proxy.js
const router = require('express').Router();
const axios = require('axios');
const OPENROUTER_KEY = process.env.OPENROUTER_API_KEY;
const APP_DOMAIN = process.env.APP_DOMAIN;
// Transcription endpoint
router.post('/transcribe', async (req, res) => {
const { audioUrl } = req.body;
try {
const response = await axios.post(
'https://openrouter.ai/api/v1/audio/transcriptions',
{ url: audioUrl },
{
headers: {
'Authorization': `Bearer ${OPENROUTER_KEY}`,
'HTTP-Referer': APP_DOMAIN
}
}
);
res.json(response.data);
} catch (error) {
res.status(error.response?.status || 500).json({ error: error.message });
}
});
// Chat completion endpoint
router.post('/completions', async (req, res) => {
const { prompt, model = 'anthropic/claude-3.5-sonnet' } = req.body;
try {
const response = await axios.post(
'https://openrouter.ai/api/v1/chat/completions',
{
model: model,
messages: [{ role: 'user', content: prompt }],
max_tokens: 1024
},
{
headers: {
'Authorization': `Bearer ${OPENROUTER_KEY}`,
'HTTP-Referer': APP_DOMAIN
}
}
);
res.json(response.data);
} catch (error) {
res.status(error.response?.status || 500).json({ error: error.message });
}
});
// Fallback chain with retries
router.post('/completions-with-fallback', async (req, res) => {
const { prompt } = req.body;
const models = [
'anthropic/claude-3.5-sonnet',
'anthropic/claude-3-opus',
'openai/gpt-4-turbo'
];
const maxRetries = 3;
const baseDelayMs = 1000;
for (const model of models) {
for (let attempt = 0; attempt < maxRetries; attempt++) {
try {
const response = await axios.post(
'https://openrouter.ai/api/v1/chat/completions',
{
model: model,
messages: [{ role: 'user', content: prompt }],
max_tokens: 1024
},
{
headers: {
'Authorization': `Bearer ${OPENROUTER_KEY}`,
'HTTP-Referer': APP_DOMAIN
}
}
);
return res.json(response.data);
} catch (error) {
if (error.response?.status === 429 || error.response?.status === 503) {
const delayMs = baseDelayMs * Math.pow(2, attempt);
await new Promise(r => setTimeout(r,
[truncated]
Create .env with one Open Router key:
OPENROUTER_API_KEY=your_key_here
APP_DOMAIN=https://yourdomain.com
Real provider keys stay in AWS Secrets Manager or similar, only accessible by backend services.
curl -X POST http://localhost:3000/llm/completions \
-H "Content-Type: application/json" \
-d '{
"prompt": "Explain quantum computing",
"model": "anthropic/claude-3.5-sonnet"
}'
With fallback chain:
curl -X POST http://localhost:3000/llm/completions-with-fallback \
-H "Content-Type: application/json" \
-d '{
"prompt": "Explain quantum computing"
}'
Pros
Centralized key management : One API key visible to applications instead of many. Real provider keys stay on backend infrastructure.
Single point of control : All LLM calls go through one endpoint. Easy to audit, monitor, and modify behavior.
Provider agnostic : Change models without changing app code. Switch from Claude to GPT-4 by changing a config value.
Built-in fallback : Open Router supports model fallback chains. If one model is unavailable, automatically try the next one.
Rate limiting : Open Router enforces rate limits on your account. Prevents accidental runaway costs.
Real-time monitoring : Dashboard shows spending per model, per day, API usage patterns.
Easier deployment : One set of credentials to manage in CI/CD instead of multiple provider keys.
No self-hosting : Managed service, no infrastructure burden.
5.5% cost overhead : Open Router adds 5.5% markup on top of provider pricing. At scale, this is significant.
Latency penalty : Extra network hop adds 100-200ms per request. First token latency is slightly slower than direct API calls.
Single point of failure : If Open Router is down or slow, all LLM calls are affected. No direct provider access as fallback.
Response format consistency : Different providers return slightly different response formats. Open Router normalizes this, but edge cases exist.
Rate limit visibility : Rate limits are Open Router's limits, not the underlying provider's. Limits are lower than direct API access.
Loss of volume discounts : Cannot negotiate volume pricing with providers through Open Router. Always pay standard rates plus markup.
Vendor lock-in lite : Switching away from Open Router requires code changes to point to direct provider APIs again.
Managing multiple LLM providers in one codebase
Client apps need API access (mobile apps, browser)
Key security and centralized management matter
Fallback chains add value (avoiding provider outages)
Response format consistency is important
Volume is under 1M requests per month
Want managed infrastructure without self-hosting
Committed to single provider for long-term
Volume is >1M requests per month (cost matters)
Enterprise compliance and features needed
Accept higher cost for managed enterprise experience
