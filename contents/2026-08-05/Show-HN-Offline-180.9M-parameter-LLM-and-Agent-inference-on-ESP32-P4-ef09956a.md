---
source: "https://github.com/cyfrit/p-for-llm"
hn_url: "https://news.ycombinator.com/item?id=49183893"
title: "Show HN: Offline 180.9M-parameter LLM and Agent inference on ESP32-P4"
article_title: "GitHub - cyfrit/p-for-llm: Offline 180.9M-parameter LLM and Agent inference on ESP32-P4. · GitHub"
author: "Cyfrit"
captured_at: "2026-08-05T15:08:43Z"
capture_tool: "hn-digest"
hn_id: 49183893
score: 1
comments: 0
posted_at: "2026-08-05T14:56:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Offline 180.9M-parameter LLM and Agent inference on ESP32-P4

- HN: [49183893](https://news.ycombinator.com/item?id=49183893)
- Source: [github.com](https://github.com/cyfrit/p-for-llm)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T14:56:34Z

## Translation

タイトル: HN の表示: ESP32-P4 でのオフライン 1 億 8,090 万パラメータ LLM とエージェント推論
記事のタイトル: GitHub - cyfrit/p-for-llm: ESP32-P4 でのオフライン 1 億 8,090 万パラメーターの LLM とエージェント推論。 · GitHub
説明: ESP32-P4 でのオフライン 1 億 8,090 万パラメータ LLM およびエージェント推論。 - cyfrit/p-for-llm

記事本文:
GitHub - cyfrit/p-for-llm: ESP32-P4 でのオフライン 1 億 8,090 万パラメーター LLM とエージェント推論。 · GitHub
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
サイフリット
/
p-for-llm
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット docs docs runtime ランタイム トレーニング training .gitignore .gitignore ライセンス ライセンス README.md README.md README_ZH.md README_ZH.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
PFor は LLM ru です

ESP32-P4 で動作しますが、技術的には SLM と呼ぶ必要があります。これには、Instruct(ChatML) 機能と Agent 機能がありますが、どちらも非常に初期で非常に不安定です。 ESP32-P4 での推論速度は約 9 トークン/秒です。
私は、32 MB PSRAM と 16 MB フラッシュを搭載した WT9932P4-Tiny 開発ボード (淘宝網で 39.9 人民元、約 6 米ドルで購入しました。AliExpress では約 10 米ドルです) を使用しています。
確かに、USB 接続では「オフライン」という言葉がやや疑わしいように見えます。ボードのフラッシュは PSRAM にロードされた重みを保持できないため、起動時に USB 経由で転送されます。ホストは推論を実行しません。ウェイト用の SD カードを提供すると、ホストなしで完全に実行できます。
パラメータ
値
パラメータ
180,920,432 (~180.9M)
レイヤー
12
隠しサイズ
192
語彙
32,768
注意
Q ヘッド 6 個、KV ヘッド 2 個、ヘッド寸法 32
教育省
レイヤーごとに 29 人のエキスパート、トップ 1 ルーティング
エキスパートFFN
512
PLE寸法
176
コンテキスト
1,024
量子化
W1.58A8
ターナリ、Q8、FP16 ストレージが混在する 1 億 8,090 万パラメータ モデルは、フラッシュと PSRAM で約 44 MiB を占有します。
PFor は MoE + PLE アーキテクチャを使用してハードウェアを最大限に活用します。 MoE 重みは PSRAM に保存され、PLE 重みはフラッシュに保存され、XIP を使用してアクセスされます。キャッシュと作業メモリは DIRAM に保存されます。ランタイムは主に Espressif XespV アクセラレーションを使用します。開発中の XespV テストでは、最大 30 倍の高速化が示されました。
WT9932P4-Tiny は ESP32-P4 v1.3 を使用します。初期の P4 リビジョンは 360 MHz でのみ実行できます。新しい v3 チップを使用している場合は、400 MHz で実行でき、生成速度が向上する可能性があります。
ESP32-S3 (@slvDev の協力に感謝) と比較して、P4 は、新しいハードウェア アクセラレータや最大 32 MB PSRAM のサポートなど、はるかに強力なパフォーマンスを提供します。
語彙は Qwen3.5 から削減されています。クウェンに感謝します。
ユーザー
フランスの首都はどこですか

?
アシスタント
フランスの首都はパリです。
ユーザー
Python プログラムを作成します。
アシスタント
このタスクを実行する Python プログラムの例を次に示します。
「」パイソン
def print_and_print():
i の場合、enumerate(items) の項目:
印刷(i)
プリント()
「」
このプログラムは、「items」にまだ存在しない「items」の項目を含む「item_list」という項目の「items」という新しいリストを作成します。 `print()` 関数は、項目の `item` を `items` リストに表示するために使用されます。
ユーザー
米国とは何ですか?
アシスタント
米国は世界で 3 番目に大きな国で、人口の約 1.8% がいます。
エージェント
タスク: config/cache.ini で、cache_mode を Lazy から Eager に変更します。
エージェント: config/cache.ini のキャッシュモードを検索します。
ツール: OK
4: キャッシュモード = 遅延
エージェント: config/cache.ini 4 を置き換えますcache_mode = eager
ツール: OK config/cache.ini:4 を置き換えました
エージェント: 完了cache_mode を Eager に更新しました。
その出力は多くの場合ナンセンスに過ぎず、頻繁に脱線しますが、それでも指示に従い、ほ​​ぼ主題を維持するある程度の能力を示しています。そのエージェント機能は汎用的ではありません。固定されたツール プロンプトを使用してトレーニングされました。それでも、非常に不安定ではあるものの、タスクを理解し、ツールを駆動するある程度の能力を示しています。
git クローン https://github.com/cyfrit/p-for-llm.git
cd p-for-llm
最新リリースからこれらのファイルをダウンロードし、リポジトリ ルートに配置します。
ダウンロードを確認して esptool をインストールします。
sha256sum --SHA256SUMS をチェックします
python3 -m pip install esptool
ファームウェアとモデルをフラッシュします。
python3 ランタイム/ホスト/flash.py \
--ファームウェア pfor-esp32p4.zip \
--model pfor-180m.llmcraft \
--port <ポート>
ターミナルチャットを開始します。
python3 runtime/host/chat.py --port < PORT > --artifact pfor-180m.llmcraft
コマンド: /help 、 /clear 、 /reload 、 /exit 。
python3 ランタイム/host/agent_demo.py --p

ort < PORT > --artifact pfor-180m.llmcraft
リリースされたファームウェアは、360 MHz の v3 以前の P4 ハードウェアをターゲットとしています。新しい P4 リビジョンでは、一致する構成を使用してソースから再構築する必要がある場合があります。
PFor は、RTX 5060 Ti のみを使用して約 12B の生データでトレーニングされました。トレーニング データが不十分なため、PFor は非常に不安定であり、その機能は非常に制限されています。ただし、LLM の特性は依然として示されており、特定の形式で単純なエージェント タスクを完了できます。
現在のパフォーマンスに基づいて、さらにトレーニングを行うことで大幅に向上するはずです。 LFM2.5 などの小規模モデル アーキテクチャは、すでに優れた機能を示しています。完全にトレーニングすると、1 億 8,090 万パラメータのモデルの方がはるかに優れた結果が得られるはずです。特定の垂直ドメインに合わせて調整されたトレーニング データにより、現在のアーキテクチャは、オンデバイス命令分析、情報抽出、意図分類、および構造化されたコマンド ルーティングをすでにサポートできます。
PFor は MoE を使用するため、理論的には専門家が複数の MCU に分散される可能性があります。トップ 1 ルーティング、8 ビット アクティベーション、隠しサイズ 192、および 12 レイヤーでは、生成された各トークンには約 4,608 バイトのエキスパート アクティベーション転送が必要です。 9 トークン/秒の理論上の帯域幅は、プロトコル オーバーヘッドを除いて、両方向を合わせて約 40.5 KiB/秒になります。
ESP32-P4 でのオフライン 1 億 8,090 万パラメータ LLM およびエージェント推論。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Offline 180.9M-parameter LLM and Agent inference on ESP32-P4. - cyfrit/p-for-llm

GitHub - cyfrit/p-for-llm: Offline 180.9M-parameter LLM and Agent inference on ESP32-P4. · GitHub
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
cyfrit
/
p-for-llm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits docs docs runtime runtime training training .gitignore .gitignore LICENSE LICENSE README.md README.md README_ZH.md README_ZH.md View all files Repository files navigation
PFor is an LLM running on ESP32-P4, although technically it should be called an SLM. It has Instruct(ChatML) and Agent capabilities, despite both being extremely early and highly unstable. Its inference speed on ESP32-P4 is about 9 tokens/s.
I use the WT9932P4-Tiny development board (I bought mine on Taobao for CNY 39.9, about USD 6; it is around USD 10 on AliExpress), which has 32 MB PSRAM and 16 MB Flash.
The USB connection admittedly makes "offline" look somewhat questionable. The board's Flash cannot hold the weights loaded into PSRAM, so they are transferred over USB at startup; the host performs no inference. Give it an SD card for the weights, and it can run entirely without a host.
Parameter
Value
Parameters
180,920,432 (~180.9M)
Layers
12
Hidden size
192
Vocabulary
32,768
Attention
6 Q heads, 2 KV heads, head dim 32
MoE
29 experts per layer, Top-1 routing
Expert FFN
512
PLE dimension
176
Context
1,024
Quantization
W1.58A8
With mixed ternary, Q8, and FP16 storage, the 180.9M-parameter model occupies about 44 MiB across Flash and PSRAM.
PFor uses a MoE + PLE architecture to make full use of the hardware. MoE weights are stored in PSRAM, PLE weights are stored in Flash and accessed with XIP, and the cache and working memory are stored in DIRAM. The runtime mainly uses Espressif XespV acceleration. During development, XespV tests showed speedups of up to 30x.
The WT9932P4-Tiny uses ESP32-P4 v1.3. Early P4 revisions can only run at 360 MHz. If you have a newer v3 chip, it can run at 400 MHz and may provide higher generation speed.
Compared with ESP32-S3 (thanks to @slvDev for the work), P4 provides much stronger performance, including new hardware accelerators and support for up to 32 MB PSRAM.
The vocabulary is pruned from Qwen3.5. Thanks to Qwen.
user
What is the capital of France?
assistant
The capital of France is Paris.
user
Write a Python program.
assistant
Here's an example of a Python program that accomplishes this task:
```python
def print_and_print():
for i, item in enumerate(items):
print(i)
print()
```
This program will create a new list called `items` of items called `item_list` with items from `items` that are not already in the `items`. The `print()` function is used to display the item's `item` in the `items` list.
user
What is the United States?
assistant
The United States is the third largest country in the world, with around 1.8% of the country's population.
Agent
Task: In config/cache.ini, change cache_mode from lazy to eager.
Agent: Search config/cache.ini cache_mode
Tool: OK
4: cache_mode = lazy
Agent: Replace config/cache.ini 4 cache_mode = eager
Tool: OK replaced config/cache.ini:4
Agent: Finish Updated cache_mode to eager.
Although its output is often little more than nonsense and frequently goes off the rails, it still demonstrates some ability to follow instructions and stay roughly on topic. Its Agent capability is not general-purpose: it was trained with a fixed tool prompt. Even so, it shows some ability to understand tasks and drive tools, albeit very unstably.
git clone https://github.com/cyfrit/p-for-llm.git
cd p-for-llm
Download these files from the latest Release and place them in the repository root:
Verify the downloads and install esptool :
sha256sum --check SHA256SUMS
python3 -m pip install esptool
Flash the firmware and model:
python3 runtime/host/flash.py \
--firmware pfor-esp32p4.zip \
--model pfor-180m.llmcraft \
--port < PORT >
Start terminal chat:
python3 runtime/host/chat.py --port < PORT > --artifact pfor-180m.llmcraft
Commands: /help , /clear , /reload , /exit .
python3 runtime/host/agent_demo.py --port < PORT > --artifact pfor-180m.llmcraft
The released firmware targets pre-v3 P4 hardware at 360 MHz. Newer P4 revisions may require rebuilding from source with a matching configuration.
PFor was trained on approximately 12B of raw data using only RTX 5060 Ti. Due to insufficient training data, PFor is highly unstable and its capabilities are extremely limited. However, it still shows the characteristics of an LLM and can complete simple Agent tasks in specific formats.
Based on its current performance, more training should improve it considerably. Small-model architectures such as LFM2.5 have already shown good capabilities. With complete training, a 180.9M-parameter model should have much better results. With training data tailored to a specific vertical domain, the current architecture can already support on-device instruction analysis, information extraction, intent classification, and structured command routing.
Because PFor uses MoE, experts could theoretically be distributed across multiple MCUs. With Top-1 routing, 8-bit activations, hidden size 192, and 12 layers, each generated token requires about 4,608 bytes of expert activation transfer. At 9 tokens/s, the theoretical bandwidth is about 40.5 KiB/s in both directions combined, excluding protocol overhead.
Offline 180.9M-parameter LLM and Agent inference on ESP32-P4.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
