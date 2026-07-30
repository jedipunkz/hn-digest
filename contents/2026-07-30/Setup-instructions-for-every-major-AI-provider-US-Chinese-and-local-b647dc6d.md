---
source: "https://big-agi.com/docs/connect-models"
hn_url: "https://news.ycombinator.com/item?id=49116394"
title: "Setup instructions for every major AI provider – US, Chinese, and local"
article_title: "Connect Models · Big-AGI Docs"
author: "fredliu"
captured_at: "2026-07-30T22:00:15Z"
capture_tool: "hn-digest"
hn_id: 49116394
score: 1
comments: 1
posted_at: "2026-07-30T21:58:34Z"
tags:
  - hacker-news
  - translated
---

# Setup instructions for every major AI provider – US, Chinese, and local

- HN: [49116394](https://news.ycombinator.com/item?id=49116394)
- Source: [big-agi.com](https://big-agi.com/docs/connect-models)
- Score: 1
- Comments: 1
- Posted: 2026-07-30T21:58:34Z

## Translation

タイトル: すべての主要な AI プロバイダー (米国、中国、および地方) のセットアップ手順
記事のタイトル: モデルの接続 · Big-AGI ドキュメント
説明: AI サービスに接続すると、モデルが Big-AGI に配置されます。どこから開始するか、多くのモデルの 1 つのキー、アプリ内のすべてのサービス、各ジョブが使用するモデル、および独自のマシン上のモデルです。

記事本文:
ここからスタート オープン、無料、またはプロ ファーストチャット 接続モデル 22 プロバイダー Alibaba Cloud Anthropic Azure OpenAI AWS Bedrock Cerebras Cohere Deepseek Google Gemini Groq LM Studio LocalAI Mistral Moonshot AI NVIDIA NIM Ollama OpenAI OpenRouter Perplexity Sagana AI Together AI xAI Z.ai カスタム エンドポイント 機能 添付ファイルと出力 BEAM ダイレクト接続インスペクター ペルソナ 同期とデバイス 音声ガイド モデル設定コンテキストプライバシーとアカウント アカウントとプロ バックアップ、復元、インポート すべて削除 キー ストレージ プライバシー データ フロー 問題の修正 エラー メッセージ キーとアクセスのエラー チャットの欠落 開く: セルフホスト構成 オプション サービス セルフホストの問題 モデルの接続
AI サービスに接続すると、モデルが Big-AGI に配置されます。どこから開始するか、多くのモデルに対応する 1 つのキー、アプリ内のすべてのサービス、各ジョブが使用するモデル、および独自のマシン上のモデルです。
ここからスタート オープン、無料、またはプロ ファーストチャット 接続モデル 22 プロバイダー Alibaba Cloud Anthropic Azure OpenAI AWS Bedrock Cerebras Cohere Deepseek Google Gemini Groq LM Studio LocalAI Mistral Moonshot AI NVIDIA NIM Ollama OpenAI OpenRouter Perplexity Sagana AI Together AI xAI Z.ai カスタム エンドポイント 機能 添付ファイルと出力 BEAM ダイレクト接続インスペクター ペルソナ 同期とデバイス 音声ガイド モデル設定コンテキストおよび制限 プライバシーとアカウント アカウントとプロ バックアップ、復元、インポート すべてを削除 キー ストレージ プライバシー データ フロー 問題の修正 エラー メッセージ キーとアクセスのエラー チャットの欠落 開く: セルフホスト構成 オプション サービス セルフホストの問題 Big-AGI は空から始まります: AI サービスを接続すると、モデルで満たされます。 1 つのサービスの 1 つのキーで開始するには十分です。最小費用はなく、2 番目のサービスには追加費用がかかりません。重要なのはチャットのサブスクリプションではありません。AI サービスは消費者向けチャット製品を販売しています。

API は 2 つの残高を持つ 2 つのアカウントとして扱われます。したがって、ChatGPT Plus または Claude Pro サブスクリプションにはキーは含まれず、Anthropic はそれを直接述べています。キーはサービスの開発者コンソールから独自のバランスで取得されます。
最初に、アプリが提供する 4 つの OpenAI 、 Anthropic 、 Gemini または OpenRouter のいずれかを開きます。特定のベンダー独自のモデルが必要な場合は、そのベンダーに問い合わせてください。有効なキーを取得するまでのステップを最小限に抑えたい - OpenRouter。支払い方法を使用せずに試したい - Gemini には、送信内容に関する独自の条件による無料枠があります。
これら 4 つは、追加サービス リストの注目グループであり、初回実行セットアップで提供されるのと同じ 4 つです。他のすべては 1 回クリックするだけです。これは 3 つの手順でカバーされます。サービスのコンソールでキーを作成し、それをモデルに貼り付けます ( Ctrl + Shift + M )。すると、そのサービスのモデルがリストに表示されます。それが機能したことを示します (最初のチャット)。
アグリゲーターはまさにそれを行います。OpenRouter はそのリストに含まれるものです。つまり、1 つのアカウント、1 つの残高、主要ベンダーのモデルにわたるカタログです。 Big-AGI は、ワンクリックでハンドオフを実現します。パネル内の OpenRouter キーをリンクすると、OpenRouter に送信されて認証され、キーがインストールされた状態で戻ります。この取引は仲介者です。つまり、1 つの請求書、1 つのカタログ、およびデフォルトで厳選されたモデル ファミリにフィルタリングされたリストです。追加サービス リストのギフト アイコンは、一部のモデルを無料で公開するサービスをマークします。
OpenAI Anthropic Gemini OpenRouter クラウド
Alibaba Cloud AWS Bedrock Azure OpenAI Cerebras Cohere DeepSeek Groq Mistral Moonshot AI NVIDIA NIM Perplexity Sakena AI Together AI xAI Z.ai Local
オラマ LM スタジオ LocalAI
OpenAI API を使用する壁の外側のものはすべて接続されます (カスタム エンドポイント)。新しいサービスのサポートは、最初にオープンソース コードベースで記述され、プロモーション時にホスト サイトに届きます。
Big-AGI はジョブ ラスごとにモデルを割り当てます

サービスごとではなく、接続しているすべてのものにわたって:
読み取られた瞬間に自動解決され、手動で厳選されたピックが存在する場合に勝ちます。モデルのギアのボタンのデフォルト/モデルを使用してジョブをピン留めします。チャットの上のドロップダウンは、そのチャットのみを変更します。
Ollama 、LM Studio、および LocalAI は他のサービスと同様に接続します。Big-AGI にサーバーのアドレスを指定すると、そのモデルが同じリストに追加されます。キーもアカウントもトークンごとの請求もありません。
直接接続をオフにすると、Big-AGI 高速エッジ サーバーは指定されたアドレスを呼び出します。クラウド サーバーはコンピューター上のアドレスに到達できません。これをオンにすると、ブラウザはアドレスを直接呼び出します。そのため、パネルはローカル サービスにこのアドレスを推奨しています。
直接接続は、ブラウザに保存されている API キーと、ブラウザの直接呼び出し (CORS) を許可する AI サービスでのみ機能します。使用できない場合、リクエストは代わりに Big-AGI 高速エッジ サーバーを介してルーティングされます。標準のアップロード サイズと時間制限内で、すべてが機能します。
トークンごとの請求は不要です。代わりにハードウェアと電気代がかかります。
オフラインで作業し、材料をマシンから離れることはありません。
メーター制ではありません。1000 ターンのコストは 1 ターンのコストと同じです。
OpenAI 互換エンドポイントを接続する
© 2026 Token Fabrics · サンディエゴで情熱を持って構築

## Original Extract

Connecting an AI service is what puts models in Big-AGI: where to start, one key for many models, every service in the app, the model each job uses, and models on your own machines.

Start Here Open, Free, or Pro First chat Connect Models 22 providers Alibaba Cloud Anthropic Azure OpenAI AWS Bedrock Cerebras Cohere Deepseek Google Gemini Groq LM Studio LocalAI Mistral Moonshot AI NVIDIA NIM Ollama OpenAI OpenRouter Perplexity Sakana AI Together AI xAI Z.ai Custom endpoints Features Attachments & outputs BEAM Direct Connection Inspector Personas Sync & devices Voice Guides Model Settings Context & limits Privacy & Account Account & Pro Back up, restore & import Delete everything Key storage Privacy data flow Fix problems Error messages Key & access errors Missing chats Open: Self Host Configuration Optional services Self-hosted issues Connect Models
Connecting an AI service is what puts models in Big-AGI: where to start, one key for many models, every service in the app, the model each job uses, and models on your own machines.
Start Here Open, Free, or Pro First chat Connect Models 22 providers Alibaba Cloud Anthropic Azure OpenAI AWS Bedrock Cerebras Cohere Deepseek Google Gemini Groq LM Studio LocalAI Mistral Moonshot AI NVIDIA NIM Ollama OpenAI OpenRouter Perplexity Sakana AI Together AI xAI Z.ai Custom endpoints Features Attachments & outputs BEAM Direct Connection Inspector Personas Sync & devices Voice Guides Model Settings Context & limits Privacy & Account Account & Pro Back up, restore & import Delete everything Key storage Privacy data flow Fix problems Error messages Key & access errors Missing chats Open: Self Host Configuration Optional services Self-hosted issues Big-AGI starts empty: connecting an AI service is what fills it with models. One key from one service is enough to start: no minimum spend, and a second service costs nothing to add. A key is not a chat subscription: AI services sell their consumer chat product and their API as two accounts with two balances. A ChatGPT Plus or Claude Pro subscription therefore carries no key, and Anthropic states it directly . The key comes from the service's developer console, on its own balance.
Open one of the four the app offers first: OpenAI , Anthropic , Gemini or OpenRouter . Want a specific vendor's own models - go to that vendor. Want the fewest steps to a working key - OpenRouter. Want to try without a payment method - Gemini has a FREE tier, on its own terms about what you send.
Those four are the Featured group in the add-service list, and the same four the first-run setup offers. Everything else is one click further. Three steps cover it: create the key in the service's console, paste it into Models ( Ctrl + Shift + M ), and that service's models fill the list - the signal it worked ( your first chat ).
An aggregator does exactly that, and OpenRouter is the one in the list: one account, one balance, a catalogue spanning the major vendors' models. Big-AGI carries a one-click handoff for it - Link OpenRouter Key in that panel sends you to OpenRouter to authorise and returns with the key installed. The trade is a middleman: one bill, one catalogue, and a list filtered to curated model families by default. A gift icon in the add-service list marks services that publish some models at no charge.
OpenAI Anthropic Gemini OpenRouter Cloud
Alibaba Cloud AWS Bedrock Azure OpenAI Cerebras Cohere DeepSeek Groq Mistral Moonshot AI NVIDIA NIM Perplexity Sakana AI Together AI xAI Z.ai Local
Ollama LM Studio LocalAI
Anything outside the wall that speaks the OpenAI API connects too: custom endpoints . Support for a new service is written in the open-source codebase first and reaches the hosted site on promotion.
Big-AGI assigns a model per job rather than per service, across everything you have connected:
Auto resolves the moment it is read, and a hand-curated pick wins where one exists. Pin a job with the Default / Model for buttons in a model's gear; the dropdown above a chat changes that chat only.
Ollama , LM Studio and LocalAI connect like any other service: point Big-AGI at the server's address, and its models join the same list - no key, no account, no per-token bill.
With Direct Connection off, the Big-AGI fast edge servers call the address you gave; a cloud server cannot reach an address on your computer. With it on, your browser calls the address directly - which is why the panel recommends it for local services.
Direct Connection works only with your API key stored in the browser , and with an AI service that permits direct browser calls (CORS). Where it cannot be used, requests route through the Big-AGI fast edge servers instead - everything still works, within the standard upload size and time limits.
No per-token bill - hardware and electricity instead.
Work offline, with material that never leaves the machine.
Nothing metered: a thousand turns cost what one costs.
Connect any OpenAI-compatible endpoint
© 2026 Token Fabrics · Built with passion in San Diego
