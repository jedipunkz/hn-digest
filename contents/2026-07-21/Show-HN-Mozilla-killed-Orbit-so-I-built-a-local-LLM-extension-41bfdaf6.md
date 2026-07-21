---
source: "https://github.com/darshi1337/apogee"
hn_url: "https://news.ycombinator.com/item?id=48993935"
title: "Show HN: Mozilla killed Orbit, so I built a local-LLM extension"
article_title: "GitHub - darshi1337/apogee: Apogee is a browser extension that uses local large language models to summarize and answer questions about articles, videos, emails, PDFs, and more. It runs entirely on-device, via WebGPU or WebAssembly, or connects directly to your own local Ollama instance. Your conten\n[truncated]"
author: "darshi7331"
captured_at: "2026-07-21T16:16:46Z"
capture_tool: "hn-digest"
hn_id: 48993935
score: 1
comments: 0
posted_at: "2026-07-21T15:49:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Mozilla killed Orbit, so I built a local-LLM extension

- HN: [48993935](https://news.ycombinator.com/item?id=48993935)
- Source: [github.com](https://github.com/darshi1337/apogee)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T15:49:51Z

## Translation

タイトル: HN を表示: Mozilla が Orbit を廃止したので、ローカル LLM 拡張機能を構築しました
記事のタイトル: GitHub - darshi1337/apogee: Apogee は、ローカルの大規模言語モデルを使用して、記事、ビデオ、電子メール、PDF などに関する質問を要約し、回答するブラウザ拡張機能です。 WebGPU または WebAssembly を介して完全にオンデバイスで実行されるか、独自のローカル Ollama インスタンスに直接接続されます。あなたのコンテンツ
[切り捨てられた]
説明: Apogee は、ローカルの大規模言語モデルを使用して、記事、ビデオ、電子メール、PDF などに関する質問を要約し、回答するブラウザ拡張機能です。 WebGPU または WebAssembly を介して完全にオンデバイスで実行されるか、独自のローカル Ollama インスタンスに直接接続されます。コンテンツがマシンから離れることはありません。
[切り捨てられた]
HN テキスト: こんにちは、HN。 昨年の夏、Mozilla は Orbit (独自のページ サマライザー拡張機能) を廃止しました。そこで私はそれを再構築しました。完全にローカルで、WebLLM (Chrome/Edge)、WASM (Firefox)、またはより強力なモデルが必要な場合は独自の Ollama インスタンスで実行します。貢献は大歓迎です。

記事本文:
GitHub - darshi1337/apogee: Apogee は、ローカルの大規模言語モデルを使用して、記事、ビデオ、電子メール、PDF などに関する質問を要約して回答するブラウザ拡張機能です。 WebGPU または WebAssembly を介して完全にオンデバイスで実行されるか、独自のローカル Ollama インスタンスに直接接続されます。コンテンツがマシンから離れることはありません。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロ

広告をクリックしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ダルシ1337
/
遠地点
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
79 コミット 79 コミット .github .github apogee-extension apogee-extension .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
WebGPU、WebAssembly、または独自のローカル Ollama を利用したプライベートのブラウザ内 AI サマライザー。
Apogee は、記事、ビデオ、電子メールなどのための AI ブラウザ アシスタントです。
それは完全にブラウザ内で実行されます。WebGPU 経由で GPU (Chrome、Edge、
および他の Chromium ブラウザ)、または WebAssembly (Firefox) 経由で CPU 上で動作します。いいえ
バックエンド、API キー、クラウドはありません。拡張機能をインストールするだけです。
パワー ユーザー向けに、Apogee は 127.0.0.1 経由でローカルの Ollama インスタンスに直接接続し、より大規模なモデルを実行することもできます。
TL;DR : Apogee は、クラウドへの依存関係や API キーを一切使用せず、Chromium ブラウザーの WebGPU および Firefox の WebAssembly 上で完全にブラウザー内で実行される、オフラインファーストのプライベート AI アシスタントです。記事、YouTube ビデオ、PDF を要約し、ローカル検索を使用してそれらに関する質問に回答します。すべて完全なプライバシーが保たれます。パワー ユーザーはローカル Ollama モードに切り替えて、何も残さずに自分のマシン上で大規模なモデルを実行できます。 Apogee は、Mozilla の廃止された Orbit のようなクラウド依存ツールに代わる、完全にローカルでプライバシーを尊重した代替手段として設計されています。
インスピレーション: Orbit (Mozilla によって削除)
Apogee は Mozilla の中止された Orbit プロジェクトからインスピレーションを受けました

(Mozilla による Orbit のレビューを読んでください)。 Orbit はブラウザベースのページ要約を提供しようとしましたが、集中 API サーバー (Mistral 7B) に依存し、 store_result などのエンドポイントを使用してサーバー側で要約をキャッシュしました。
Apogee は、完全にローカルファーストであることで、Orbit のアーキテクチャとプライバシーの欠陥を修正します。
サーバー オーバーヘッドゼロ : Apogee は、リモート クラウド API を介してクエリをルーティングする代わりに、WebGPU を介して完全にオンデバイスでトークン化と推論を実行します。
データ漏洩なし : Apogee はページのコンテンツや生成された概要を外部エンドポイントに送信せず、データがマシンから流出することはありません。
企業の独立性 : Apogee にはサーバーへの依存関係や料金を支払うクラウド インフラストラクチャがないため、シャットダウンしたりサービスを終了したりすることはありません。
Apogee は、量子化された言語モデルをブラウザーで直接実行します。 Chromeでは、
Edge とそれが使用する他の Chromium ブラウザ
WebLLM : 上でモデルを実行します。
WebGPU経由のGPU。 Firefox では使用します
Transformers.js 、
WebAssembly を介して CPU 上で小型の ONNX モデルを実行し、優れたパフォーマンスを発揮します。
最新のCPU。どちらの場合も、最初の使用ではモデルの重み (およそ
モデルに応じて 270 MB ～ 2.2 GB）、ローカルにキャッシュされます。後
つまり、すべてがオフラインで実行されます。
より大きなモデルをお好みですか? Local Ollama モードに切り替えると拡張機能が会話します
127.0.0.1 を介して独自の Ollama インスタンスに直接接続します。別途必要はありません。
インストールまたは実行するバックエンド。
Ask は概要そのものよりもさらに踏み込みます。盲目的に切り捨てるのではなく、
最初の数千文字までの長いページ、Apogee はページを埋め込みます
ローカル (小規模なオンデバイス モデル、上記の LLM 重みと同じ信頼層)
あなたの質問に最も関連のある文章のみを使用して回答します。これ
長い記事、PDF、またはビデオの奥深くに埋もれている内容について質問することを意味します
トランスクリプトに適合するものだけでなく、依然として機能します。

切り取られたスライスを開きます。
(現在、取得は Chromium ブラウザで実行されています。Firefox は
今のところはプレーンな切り詰められたスライスです。)
ページ内のハイライトを使用すると、ソースと比較して概要を確認できます。クリックします。
任意の箇条書き (文/段落モードでは行) を押すと、Apogee が
元のページの一節は、同じものを使用することに基づいている可能性が高いです
デバイス上の検索 Ask が使用し、ライブでスクロールしてハイライト表示します。
ページ。記事全体を読み直さずに主張をスポットチェックするのに役立ちます。
現時点では Chromium のみ。上記の Ask の取得と同じ制約。
拡張機能をインストールします (下記を参照)。
「Apogee」アイコンをクリックし、「このページの要約」をクリックします。
初めて使用するとき、モデルは自動的にダウンロードされます。その後はインスタントです。
それでおしまい。バックエンドのインストールやターミナル コマンドは必要ありません。
より速く要約する方法: ページ上の任意の場所を右クリックして選択します。
このページを要約するか、キーボード ショートカット (デフォルトは Alt+Shift+U 、
chrome://extensions/shortcuts で再マッピング可能)。どちらも開かなくても動作します
ポップアップ自体。システム通知により、準備が完了したことがわかります。クリック
それを実行して (またはポップアップを開いて)、結果を確認します。ポップアップ中にポップアップを開くと、
まだ生成中なので、デフォルトではなく通常の読み込みビューが表示されます
ホームページ。
Apogee は、使いやすさと本来の機能のバランスをとるために 2 つの動作モードを提供します。
モデル
ダウンロードサイズ
最適な用途
クウェン 2.5 1.5B (デフォルト)
～900MB
多言語要約
SmolLM2 1.7B
～1GB
一般的なタスク
ラマ 3.2 1B
～700MB
軽量、高速
ファイ 3.5 ミニ
～2.2GB
より強力な推論
Transformers.js (WASM/CPU、Firefox のみ)
モデル
ダウンロードサイズ
最適な用途
SmolLM2 360M (デフォルト)
～270MB
最小/最速の簡単な要約
クウェン 2.5 0.5B
~480MB
多言語対応
ラマ 3.2 1B
～1.2GB
より強力な推論、CPU の処理速度が遅い
Transformers.js 経由で実行
(WASM バックエン上の ONNX モデル

d)。決してそうではないという理由で特別に選ばれました
これとは異なり、ワーカーを生成します (onnxruntime-web のプロキシ モードはハードコードされてオフになっています)。
WebGPU ベースの WebLLM (Firefox にはないオフスクリーン ドキュメントが必要) または
wllama (BLOB が必要です: -URL Worker。これは Chrome と Firefox の両方にあります)
拡張 CSP ブロック)。生成はシングルスレッドです (拡張ページはシングルスレッドではありません)
クロスオリジン分離されるため、 SharedArrayBuffer がなく、コンテキストの上限は次のとおりです。
CPU での遅延を適切に保つための 4096 トークン。最新の高速 CPU ではこれ
デフォルトの SmolLM2 360M モデルでもよく要約されています。古いものや
低電力ハードウェアの場合、生成が著しく遅くなることが予想されるため、考慮してください。
代わりに Local Ollama に切り替えます。
取得したモデルはすべて拡張機能の設定に自動的に表示されます
(詳細: ローカル Ollama モードを参照)。これらは
まだ何も取得していない場合は、単なる出発点にすぎません。
要約は、選択したモデルに合わせてチャンク化も行います。
モデル (例: llama3.1 、 qwen2.5 、 gemma3 ) のチャンクが大きくなり、チャンクが少なくなります
同じ固定チャンク サイズではなく、長いコンテンツを渡します。
モデルが実際に処理できるもの。
Apogee は 2 つのビルドを出荷します: Chromium ビルド ( dist/chrome 、Manifest V3 と
WebGPU 用のオフスクリーン ドキュメント) と Firefox ビルド ( dist/firefox 、いいえ
オフスクリーンドキュメント）。 Chromium ベースのものはすべて同じビルドを受け入れます。
MDN の WebGPU API ブラウザ互換性表を参照してください。
ブラウザごと/OS ごとの WebGPU バージョンを正確にサポートするため、これは急速に変化するターゲットです
ここにハードコードされた数値よりも優れた真実の情報源です。 WebGPU を備えた GPU
サポート (過去数年間のほとんどの GPU) が必要です。
特にブラウザ内 (WebLLM) モード。ローカル Ollama モードには GPU がありません
オラマ自身が必要とするものを超えた、独自の要件。
クロム、エッジ、ブレイブ、オペラ、ヴィヴァルディ、アーク、ディア
これらはすべて Chromium ベースであり、

同じ dist/chrome ビルドとロード
ステップ。拡張機能ページの URL のみがわずかに異なります ( chrome://extensions 、
edge://extensions 、brave://extensions 、dia://extensions/ など)。
パッケージ化された拡張機能 .zip を Releases からダウンロードします。
ダウンロードした .zip ファイルをマシン上で抽出/解凍します。
ブラウザの拡張機能ページを開きます (Chrome/Brave/Opera/Vivaldi では chrome://extensions、Edge ではedge://extensions、Dia では dia://extensions/)。
開発者モードを有効にします (右上のトグル)。
「解凍してロード」をクリックし、抽出したフォルダー (ZIP ファイル自体ではなく、manifest.json を含む) を選択します。
cd apogee-extension && npm install && npm run build
ブラウザの拡張機能ページに移動し、開発者モードを有効にします。
「Load unpacked」をクリックし、apogee-extension/dist/chrome フォルダーを選択します。
npm run build は dist/chrome と dist/firefox の両方を生成します。 Chromium ベースのブラウザには dist/chrome を使用し、Firefox には dist/firefox を使用します。 npm run build:chrome または npm run build:firefox を使用して単一のターゲットをビルドすることもできます。
Apogee は Mozilla Add-ons から直接インストールすることも、 Releases からパッケージをダウンロードすることもできます。
Firefox はそのままで動作します。ブラウザ内 AI は Transformers.js 経由で実行されます。
WebAssembly (デフォルトでは SmolLM2 360M、セットアップは必要ありません)。古いものや
低電力 CPU、生成が遅くなる可能性があるため、ローカル Ollama モードに切り替えてください。
より大きなモデルでより高速な結果を得るための設定。
Ollama を介してより大きなモデル (8B+) をローカルで実行したい場合は、Apogee が話します
HTTP 経由で直接アクセスします。個別のバックエンド サーバーをインストールする必要はありません
または走り続けます。
https://ollama.com からインストールし、必要なモデルをプルします。
ollam は gemma3:4b # と qwen3:8b、mistral:latest、llama3.1:8b をプルします
2. 拡張機能が Ollama (CORS) に到達できるようにします。
Ollama は、許可リストに登録されている一連のブラウザー発信のリクエストのみを受け入れます。
オリジンと chrome-extensi

on:// はデフォルトのリストにありません。セット
Ollama を起動する前の OLLAMA_ORIGINS:
OLLAMA_ORIGINS= " chrome-extension://<your-extension-id> " ollama サーブ
chrome://extensions で <your-extension-id> を見つけます (開発者モードがオン)。
Ollama が代わりにシステム サービスとして実行される場合 (例: パッケージ経由でインストールされる)
Linux 上のマネージャー)、それを渡すのではなく永続的なオーバーライドとして設定します。
コマンドライン:
sudo mkdir -p /etc/systemd/system/ollama.service.d
printf ' [サービス]\n環境="OLLAMA_ORIGINS=chrome-extension://<拡張子ID>"\n ' \
| sudo tee /etc/systemd/system/ollama.service.d/override.conf
sudo systemctl デーモン-リロード
sudo systemctl ollamを再起動します
セキュリティ上の注意: OLLAMA_ORIGINS=chrome-extension://* (すべての
extension) は開発中に便利ですが、他の拡張子を意味します
ブラウザにインストールされているものも、ローカルの Ollama インスタンスに到達する可能性があります。
API。クイック以外のものについては、特定の拡張機能 ID にスコープを設定します
ローカルテスト。
3. 拡張機能を Ollama に向けます
拡張機能を開いて [設定] に移動し、 [ローカル Ollama] を選択します。ホスト
フィールドのデフォルトは http://127.0.0.1:11434 (Ollama 独自のデフォルト ポート)、
Ollama が他の場所でリッスンするように設定している場合にのみ変更してください。
接続すると、

[切り捨てられた]

## Original Extract

Apogee is a browser extension that uses local large language models to summarize and answer questions about articles, videos, emails, PDFs, and more. It runs entirely on-device, via WebGPU or WebAssembly, or connects directly to your own local Ollama instance. Your content never leaves your machine.
[truncated]

Hi HN, Last summer, Mozilla killed Orbit (their own page-summarizer extension). So I rebuilt it: fully local, running on WebLLM (Chrome/Edge), WASM (Firefox), or your own Ollama instance if you want a more powerful model. Contributions are very welcome.

GitHub - darshi1337/apogee: Apogee is a browser extension that uses local large language models to summarize and answer questions about articles, videos, emails, PDFs, and more. It runs entirely on-device, via WebGPU or WebAssembly, or connects directly to your own local Ollama instance. Your content never leaves your machine. · GitHub
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
darshi1337
/
apogee
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
79 Commits 79 Commits .github .github apogee-extension apogee-extension .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Private, in-browser AI summarizer powered by WebGPU, WebAssembly, or your own local Ollama.
Apogee is an AI browser assistant for articles, videos, emails, and more.
It runs entirely in your browser : on your GPU via WebGPU (Chrome, Edge,
and other Chromium browsers) or on your CPU via WebAssembly (Firefox). No
backend, no API keys, no cloud. Just install the extension and go.
For power users, Apogee also connects directly to a local Ollama instance over 127.0.0.1 to run larger models.
TL;DR : Apogee is an offline-first, private AI assistant that runs entirely in your browser, on WebGPU in Chromium browsers and on WebAssembly in Firefox, with zero cloud dependencies or API keys. It summarizes articles, YouTube videos, and PDFs, and answers questions about them using local retrieval, all with complete privacy. Power users can switch to Local Ollama mode to run larger models on their own machine, still with nothing leaving it. Apogee is designed as a fully local, privacy-respecting alternative to cloud-dependent tools like Mozilla's discontinued Orbit.
Inspiration: Orbit (Killed by Mozilla)
Apogee was inspired by Mozilla's discontinued Orbit project (read the Review of Orbit by Mozilla ). Orbit attempted to provide browser-based page summarization, but it relied on centralized API servers (Mistral 7B) and cached summaries on the server side using endpoints like store_result .
Apogee fixes Orbit's architectural and privacy flaws by being fully local-first:
Zero Server Overhead : Instead of routing queries through remote cloud APIs, Apogee performs tokenization and inference completely on-device via WebGPU.
No Data Leaks : Apogee does not send page content or generated summaries to any external endpoint, your data never leaves your machine.
Corporate Independence : Because Apogee has no server dependencies or cloud infrastructure to pay for, it can never be shut down or sunset.
Apogee runs quantized language models directly in your browser. On Chrome,
Edge, and other Chromium browsers it uses
WebLLM , which executes models on your
GPU via WebGPU. On Firefox it uses
Transformers.js , which
runs smaller ONNX models on your CPU via WebAssembly and performs well on a
modern CPU. In both cases the first use downloads the model weights (roughly
270 MB to 2.2 GB depending on the model) and caches them locally. After
that, everything runs offline.
Prefer larger models? Switch to Local Ollama mode and the extension talks
directly to your own Ollama instance over 127.0.0.1 , with no separate
backend to install or run.
Ask goes further than the summary itself: instead of blindly truncating
long pages to the first few thousand characters, Apogee embeds the page
locally (a small on-device model, same trust tier as the LLM weights above)
and answers using only the passages most relevant to your question. This
means asking about something buried deep in a long article, PDF, or video
transcript still works, not just what fit in the opening truncated slice.
(Retrieval currently runs on Chromium browsers; Firefox falls back to the
plain truncated slice for now.)
Highlight-in-page lets you check the summary against the source: click
any bullet (or line, in Sentences/Paragraphs mode) and Apogee finds the
passage of the original page it's most likely grounded in using the same
on-device retrieval Ask uses, then scrolls to and highlights it in the live
page. Useful for spot-checking a claim without re-reading the whole article.
Chromium-only for now, same constraint as Ask's retrieval above.
Install the extension (see below).
Click the Apogee icon, then Summarize this page .
On first use, the model downloads automatically. After that it's instant.
That's it. No backend installation, no terminal commands.
Faster ways to summarize : right-click anywhere on a page and pick
Summarize this page , or use the keyboard shortcut (default Alt+Shift+U ,
remappable at chrome://extensions/shortcuts ). Either works without opening
the popup at all. A system notification lets you know when it's ready; click
it (or open the popup) to see the result. If you open the popup while it's
still generating, it shows the normal loading view instead of the default
Home page.
Apogee offers two modes of operation to balance ease-of-use and raw capabilities:
Model
Download Size
Best For
Qwen 2.5 1.5B (default)
~900 MB
Multilingual summarization
SmolLM2 1.7B
~1 GB
General tasks
Llama 3.2 1B
~700 MB
Lightweight, fast
Phi 3.5 Mini
~2.2 GB
Stronger reasoning
Transformers.js (WASM/CPU, Firefox only)
Model
Download Size
Best For
SmolLM2 360M (default)
~270 MB
Smallest/fastest, quick summaries
Qwen 2.5 0.5B
~480 MB
Multilingual
Llama 3.2 1B
~1.2 GB
Stronger reasoning, slower on CPU
Runs via Transformers.js
(ONNX models on the WASM backend). Chosen specifically because it never
spawns a Worker (onnxruntime-web's proxy mode is hardcoded off), unlike
WebGPU-based WebLLM (needs an offscreen document Firefox doesn't have) or
wllama (needs a blob: -URL Worker, which both Chrome's and Firefox's
extension CSP block). Generation is single-threaded (extension pages aren't
cross-origin-isolated, so no SharedArrayBuffer ) and context is capped at
4096 tokens to keep latency reasonable on CPU. On a modern/fast CPU this
still summarizes well with the default SmolLM2 360M model; on older or
low-power hardware, expect noticeably slower generation, and consider
switching to Local Ollama instead.
Any model you've pulled shows up automatically in the extension's settings
(see Advanced: Local Ollama Mode ). These are
just a starting point if you haven't pulled anything yet:
Summarization also adapts its chunking to the model you pick: larger-context
models (e.g. llama3.1 , qwen2.5 , gemma3 ) get bigger chunks and fewer
passes over long content, rather than the same fixed chunk size regardless
of what the model can actually handle.
Apogee ships two builds: a Chromium build ( dist/chrome , Manifest V3 with an
offscreen document for WebGPU) and a Firefox build ( dist/firefox , no
offscreen document). Anything Chromium-based accepts the same build.
See MDN's WebGPU API browser compatibility table
for exact per-browser/per-OS WebGPU version support, it's a fast-moving target
and a better source of truth than a number hardcoded here. A GPU with WebGPU
support (most GPUs from the last several years) is required for the
In-Browser (WebLLM) mode specifically. Local Ollama mode has no GPU
requirement of its own beyond whatever Ollama itself needs.
Chrome, Edge, Brave, Opera, Vivaldi, Arc, Dia
These are all Chromium-based and use the same dist/chrome build and load
steps; only the extensions-page URL differs slightly ( chrome://extensions ,
edge://extensions , brave://extensions , dia://extensions/ , etc.).
Download the packaged extension .zip from Releases .
Extract/unzip the downloaded .zip file on your machine.
Open your browser's extensions page ( chrome://extensions on Chrome/Brave/Opera/Vivaldi, edge://extensions on Edge, dia://extensions/ on Dia).
Enable Developer mode (toggle in the top-right).
Click Load unpacked and select the extracted folder (containing manifest.json , not the ZIP file itself).
cd apogee-extension && npm install && npm run build
Go to your browser's extensions page and enable Developer mode .
Click Load unpacked and select the apogee-extension/dist/chrome folder.
npm run build produces both dist/chrome and dist/firefox . Use dist/chrome for any Chromium-based browser and dist/firefox for Firefox. You can also build a single target with npm run build:chrome or npm run build:firefox .
You can install Apogee directly from Mozilla Add-ons or download the package from Releases .
Firefox works out of the box: in-browser AI runs via Transformers.js on
WebAssembly (SmolLM2 360M by default, no setup needed). On older or
low-power CPUs, generation can be slow, switch to Local Ollama mode in
settings for faster results with larger models.
If you prefer running larger models (8B+) locally through Ollama, Apogee talks
to it directly over HTTP ; there's no separate backend server to install
or keep running.
Install from https://ollama.com , then pull the models you want:
ollama pull gemma3:4b # and qwen3:8b, mistral:latest, llama3.1:8b
2. Allow the extension to reach Ollama (CORS)
Ollama only accepts browser-originated requests from an allow-listed set of
origins, and chrome-extension:// isn't in that default list. Set
OLLAMA_ORIGINS before starting Ollama:
OLLAMA_ORIGINS= " chrome-extension://<your-extension-id> " ollama serve
Find <your-extension-id> on chrome://extensions (with Developer mode on).
If Ollama runs as a system service instead (e.g. installed via a package
manager on Linux), set it as a persistent override rather than passing it on
the command line:
sudo mkdir -p /etc/systemd/system/ollama.service.d
printf ' [Service]\nEnvironment="OLLAMA_ORIGINS=chrome-extension://<your-extension-id>"\n ' \
| sudo tee /etc/systemd/system/ollama.service.d/override.conf
sudo systemctl daemon-reload
sudo systemctl restart ollama
Security note : OLLAMA_ORIGINS=chrome-extension://* (allowing every
extension) is convenient while developing, but it means any other extension
installed in your browser could also reach your local Ollama instance's
API. Scope it to your specific extension ID for anything beyond quick
local testing.
3. Point the extension at Ollama
Open the extension, go to Settings, and select Local Ollama . The host
field defaults to http://127.0.0.1:11434 (Ollama's own default port),
only change it if you've configured Ollama to listen elsewhere.
Once connected, the

[truncated]
