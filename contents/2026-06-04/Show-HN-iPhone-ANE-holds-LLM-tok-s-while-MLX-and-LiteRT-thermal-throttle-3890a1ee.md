---
source: "https://github.com/john-rocky/apple-silicon-llm-bench"
hn_url: "https://news.ycombinator.com/item?id=48394262"
title: "Show HN: iPhone ANE holds LLM tok/s while MLX and LiteRT thermal-throttle"
article_title: "GitHub - john-rocky/apple-silicon-llm-bench: Neutral, reproducible benchmark for local LLMs on Apple Silicon (Mac · iPhone · iPad) — MLX, llama.cpp, CoreML, Apple Foundation Models · GitHub"
author: "mlboy"
captured_at: "2026-06-04T07:46:03Z"
capture_tool: "hn-digest"
hn_id: 48394262
score: 1
comments: 0
posted_at: "2026-06-04T05:17:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: iPhone ANE holds LLM tok/s while MLX and LiteRT thermal-throttle

- HN: [48394262](https://news.ycombinator.com/item?id=48394262)
- Source: [github.com](https://github.com/john-rocky/apple-silicon-llm-bench)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T05:17:10Z

## Translation

タイトル: Show HN: MLX と LiteRT サーマル スロットルの間、iPhone ANE は LLM tok/s を保持します
記事のタイトル: GitHub - john-rocky/apple-silicon-llm-bench: Apple Silicon (Mac · iPhone · iPad) 上のローカル LLM の中立的で再現可能なベンチマーク — MLX、llama.cpp、CoreML、Apple Foundation Models · GitHub
説明: Apple Silicon (Mac · iPhone · iPad) 上のローカル LLM の中立的で再現可能なベンチマーク — MLX、llama.cpp、CoreML、Apple Foundation モデル - john-rocky/apple-silicon-llm-bench

記事本文:
GitHub - john-rocky/apple-silicon-llm-bench: Apple Silicon (Mac · iPhone · iPad) 上のローカル LLM の中立的で再現可能なベンチマーク — MLX、llama.cpp、CoreML、Apple Foundation Models · GitHub
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
ジョン・ロッキー
/
アップルシリコンllmベンチ
公共
いいえ

通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ジョン-ロッキー/アップル-シリコン-llm-ベンチ
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
27 コミット 27 コミット .github .github apple/ YardstickCLI/ Sources apple/ YardstickCLI/ Sources デバイス デバイス docs/ charts docs/ charts ios ios 方法論 方法論 モデル モデル プロンプト プロンプト 結果 結果 ランタイム ランタイム スクリプト スクリプト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md DESIGN.md DESIGN.md ライセンス ライセンス Package.swift Package.swift README.md README.md RESULTS.md RESULTS.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Apple Silicon のオンデバイス LLM ベンチマーク — iPhone · iPad · Mac。
Apple Silicon 上でローカル LLM (そして、やがては ASR / TTS) を実行するための中立的で再現可能なベンチマーク。サーバー上の tok/s だけでなく、実際のデバイスの制約の下で、MLX Swift、llama.cpp、CoreML (swift-transformers)、LiteRT-LM、ExecuTorch、ANEMLL、および Apple 独自の Foundation Models を比較します。
リポジトリ: apple-silicon-llm-bench · CLI/ブランド: yardstick 。 ios-llm-benchmark として活動を開始 — iPhone は依然として主要なターゲットであり、現在では iPad や Mac と並んで測定されています。
📱 TL;DR — iPhone 17 Pro (A19 Pro)
電話機上での実際の LLM 推論 — サーバーなし、デバイス上で。 iPhone 17 Pro、4 ビット、ショートチャット (128 トークン)、3 回のコールド ランの中央値。勝てるランタイムはモデルに依存します。そして番狂わせは Gemma にあります。
デコード スループット — tok/s、高いほど優れています (🏆 = 勝者):
ピークメモリ — MB、低いほど優れています (🏆 = 勝者):
番狂わせ — Gemma 4 E2B: Google の LiteRT-LM (INT4-QAT、GPU、そのネイティブ .litertlm ) は、デコード時に MLX-Swift を上回り、使用するメモリが約 4.5 分の 1 (641 MB 対 2,900) です。専用のランタイムは独自の形式で優れています。
MLX-Swift が Qwen 3.5 2B デコードで勝利 — 61 対 39

トク/秒。 (LiteRT-LM には Qwen エントリがありません。そのカタログは Gemma のみです。)
CoreML / ANE はメモリのチャンピオンです — Qwen 3.5 2B はニューラル エンジンのチャンク MLKV 経由でわずか 241 MB (MLX の 1,279 MB よりも約 5 倍少ない) です — しかし、これはデコードが最も遅く (ANE はスループットをフットプリントと引き換えにします)、M4 Max と同じ話です。
ANE はデスクトップとほぼ同等です。CoreML Gemma 4 E2B は、iPhone では 33 tok/s を実行しますが、M4 Max では 32.5 tok/s を実行します。同じシリコン ファミリです。 GPU ランタイムは、実際のデバイス上の税金を支払います。M4 Max よりも最大 4 ～ 5 倍遅い (Qwen 3.5 2B → 61 tok/s vs 292)。
カウント: MLX / llama.cpp / LiteRT-LM は正確なトークナイザー トークンをレポートします ( getBenchmarkInfo 経由の LiteRT-LM)。 CoreML/ANE はストリーミングされたピース (≈ トークン) をカウントします。 LiteRT-LM は EOS まで実行されます (通話ごとの上限なし → 応答トークンは最大 458 トーク、他の予算は 128 です)。 decode tok/s はレートであるため、head-to-head が成立します。
完全に自動化され、devicectl ヘッドレス モード経由でサイドロードされます。電話機では何も入力せず、デスクトップ行と同じ方法です。
次に登場するのは、Apple Foundation モデル、さらに多くのモデル、さらに多くの iPhone / iPad です。 1 行で素晴らしい PR になります。
LiteRT-LM 行の測定方法: Metal GPU バックエンドで litert-community/gemma-4-E2B-it.litertlm (INT4-QAT) を実行する google-ai-edge/LiteRT-LM 0.12.0、ツリー内 MediaPipeRuntime アダプター経由 — 他のすべての行と同じヘッドレス ハーネス + プロンプト (3 回のコールド ラン、中央値)。トークン数と tok/s は LiteRT-LM 独自のベンチマーク カウンター ( Conversation.getBenchmarkInfo ) から取得されるため、推定ではなく正確です。 EOS まで生成されるため (API には呼び出しごとの出力上限がない)、そのトークン数は 128 トークンのバジェットではなくモデルの完全な応答になります。デコード tok/s はレートであり、同等のままです。メモリは正確なプロセス RSS です。 LiteRT-LM はローカル SwiftPM パッケージとして販売されています ( scripts/bootstrap.sh は GIT_LFS_SKIP_SMUDGE=1 でクローンを作成します。リリースされたパッケージは t

-all_load を介して SwiftPM の unsafe-flags ルールをリッピングします)。
CoreML/ANE 行の測定方法: Neural Engine 上の john-rocky/CoreML-LLM (computeUnits: .cpuAndNeuralEngine) — Gemma 4 E2B はチャンク化された .mlmodelc パス経由、Qwen 3.5 2B は Qwen35MLKVGenerator 経由 (チャンク化された MLKV、したがって 241 MB)。デコードはストリーミングされた部分 (≈ トークン) をカウントします。初回ロード ANE コンパイルではロード時間が長くなります (そして、ランタイムのスループットが最も低くなります。ANE はメモリと引き換えに速度を犠牲にします)。
デコード tok/s は見出し番号です。実行ごとの完全な監査 (プレフィル、TTFT、トークン間ジッター、メモリ) は RESULTS.md にあります。
⏱ バーストトーク/秒は話の半分にすぎません - 持続的なスロットル
上の表はコールドバースト速度です。同じモデルを継続的に実行すると、モデルが反転します。GPU ランタイム (MLX、LiteRT-LM) が加熱し、約 60 秒以内に 50% 以上のスロットルが発生しますが、ANE はほとんど動きません。ANE は電力の約半分を消費するため、ゆっくりと加熱し、SoC はスロットルしません。
2 つの独立した GPU ランタイムが同じように折りたたまれるのは、携帯電話の GPU の熱特性であり、ランタイムの癖ではありません。 MLX は ANE を下回ってしまいます。 LiteRT は速度を半分に落とした後、わずかなリードしか保っていません。 GPU がスプリントに勝ちます。 ANE がマラソンに勝利し、アプリの残りの部分のために GPU を解放します。
方法: 600 秒の連続発電、コールド (公称) スタート、アンプラグド、tg128;ローリング ウィンドウからのデコード レート。 results/raw/iphone17pro-*-energy-tg128.jsonl 内の生の JSONL ; scripts/throttle_chart.py で再描画します ( scripts/throttle_curve.py 経由のカーブテーブル)。 LiteRT-LM には出力トークンの上限がなく (呼び出しごとに長い)、実行は適切な温度で開始されます。 CoreML-LLM はスライディング ウィンドウ アテンション (境界付きコンテキスト) を使用します。これが、そのデコードがフラットなままである理由の 1 つです。
🖥 デスクトップリファレンス — Apple M4 Max
規模を拡大するために、ラップトップクラスのチップに同じハーネスを使用します。ここですべてを勝ち取るランタイムはありません — ea

ch は、スループット / メモリ / エネルギー / ストリーミング ボックスの別の部分を最適化します。
mlx-swift は、測定されたすべてのセルでデコード スループットを獲得しました (2026 年初頭のカーネル更新後は、llama.cpp と比較して 1.4 ～ 1.8 倍)。
Apple Foundation Models は、GPU ベースのランタイムよりもトークンあたりのエネルギー効率が 2 倍、CoreML/ANE よりも 4 倍優れています。
CoreML / ANE はピーク メモリ (チャンク化された MLKV) に勝りますが、J/トークンでは最も遅く、最悪です。
llama.cpp はスピードとエネルギーの点で中間に位置します。勝つ軸もなければ、大きく負ける軸もありません。
行を追加した後に再生成します: python scripts/generate_charts.py 。
📊 完全な数値 — Apple M4 Max、ショートチャット (128 トークン、デコードトークン/秒、中央値)
1 つのデバイス、4 つのランタイム、複数のモデル。デコード tok/s は主要な見出し番号です。完全なテーブル (プレフィル、TTFT、ピーク メモリ、実行ごとの監査証跡) は RESULTS.md にあります。結論を導き出す前に、「ヘッドラインの観察」セクションをお読みください。ランタイムのランキングはモデルのサイズに依存します。
クロスランタイム — 同じ論理モデル、異なるバックエンド (デコードトークン/秒、中央値)
論理モデル
パラメータ
n
mlx-swift (第 4 四半期)
ラマ.cpp (Q4_K_M)
coreml-llm
litert-lm (.litertlm)
クウェン 2.5 0.5B
0.5B
3
531.1
297.1
181.2(FP16)
該当なし
クウェン 3.5 0.8B
0.8B
3
421.1
201.1
58.2 (INT8)
該当なし
クウェン 3.5 2B
2B
3
291.9
149.7
35.0 (INT8)
該当なし
ジェマ 4 E2B
2B
3
185.4
119.2
32.5 (INT4 パレット化)
保留中
ジェマ4 E4B
4B
3
113.5
80.5
走らない
保留中
litert-lm 列: 保留中 = google-ai-edge/LiteRT-LM v0.12.0 に対して接続されたアダプター、M4 Max の実行はまだキャプチャされていません ( RESULTS.md / Yardstick_USER_RUNS.md を参照)。 n/a = LiteRT-LM のカタログは Gemma のみ ( .litertlm ) であるため、Qwen 行にはエントリがありません。参考までに、Google の E2B モデル カードは、iPhone 17 Pro GPU で 56.5 tok/s を報告しています。これは、M4 Max Yardstick の測定値ではなく、別のデバイスでのベンダーの数値です。
→ MLX-Swift がデコを獲得

上流の mlx-swift-lm が 2026 年初頭に Qwen + Gemma カーネル アップデートを出荷した後、すべてのセルで de (llama.cpp の 1.4 ～ 1.8 倍) (Qwen 行は、それらが到達する前にキャプチャされたスナップショットと比較して約 3 倍になりました)。古い「llama.cpp Metal は常に小規模モデルのデコードに勝つ」ルールは M4 Max では当てはまりません。引用する前に再測定してください。 CoreML / ANE は、以下に示す劇的なメモリ節約と引き換えに、すべてのセルで 3 つの中で最も低速です。
クロスランタイム — ピークメモリ (MB、中央値)
上記の decode-tok/s テーブルはメモリ側を隠しています。同じモデルで、代わりにピークのワーキングセットに注目します。
→ チャンク化された MLKV レイアウトが開始されると、「CoreML/ANE がメモリを獲得する」ということが当てはまります。0.5 B パラメータでは、MLX-Swift はまだ小さいです (413 MB 対 CoreML の 959 MB モノリシック FP16)。 0.8 B 以降、CoreML のチャンク化された MLKV パス (Qwen35MLKVGenerator : mmap 埋め込みサイドカー + オンデマンド ANE チャンク) は、プロセス RSS をほぼフラット (0.8 B で 206 MB、2 B で 215 MB) に保持しますが、MLX と llama.cpp はパラメータ数に比例してスケールします。
クロスランタイム — トークンあたりのエネルギー (Gemma 4 E2B、sustained-512、M4 Max)
誰も公開していない数値: 各バックエンドは生成されたトークンごとに何ジュールを消費しますか? scripts/measure_energy.py 経由でキャプチャされ、パワーメトリクス (システム全体、パッケージ電力 = CPU + GPU + ANE) を同時実行し、サンプル ウィンドウをベンチの報告されたアクティブ時間にクリップします。
ANE パスは、フル デコード時に GPU パスのパッケージ電力の約半分を消費します (12.7 W 対 ~24.7 W)。これと同じ電力差により、ANE がレートを維持しながら iPhone で GPU ランタイムが熱的にスロットルされます (上記の持続的なスロットルのセクションを参照)。
→ エネルギーのランキングは、decode-tok/s のランキングを逆転させます。 Apple FM は、トークンの生成速度が約半分であるにもかかわらず、GPU ベースのランタイムに比べてトークンあたりの効率が 2 倍優れています。 CoreML/ANE は瞬間的な Po が最も低い

wer (12.7 W) ですが、デコードが遅い (32 tok/s) ため、パッケージの電源投入時間がずっと長く続くため、4x Apple FM では最悪の J/tok になります。 MLX-Swift と llama.cpp は最も多くの W (GPU) を消費しますが、~0.24 J/tok で損益分岐点になるのに十分な速さでトークンを生成します。システム全体の測定にはアイドル状態のベースラインが含まれるため、4 つの数値すべてがトークンあたりのエネルギーをわずかに増加させます。これは、絶対的な属性ではなくランキングに役立ちます。 iPhone Energy は、代わりに 1 %-battery-step API を使用します (異なる方法論、同様のテーブル形状)。
llama.cpp (Q4_K_M GGUF、M4 Max、ショートチャット)
mlx-swift (Q4 / MLX、M4 Max、ショートチャット)
coreml-llm (CoreML / ANE、M4 Max、ショートチャット)
→ CoreML/ANE はスループットをメモリと引き換えにします。同じモデル サイズの MLX-Swift / llama.cpp よりもピーク ワーキング セットが 3 ～ 8 倍少なく、デコード tok/s の約半分です。 Qwen 3.5 0.8B / 2B の数値は、汎用の CoreMLLLM.load(from:) パスではなく、専用の Qwen35MLKVGenerator (ANE チャンク デコード、MLState の KV — CoreML-LLM v1.9.0 以降のパブリック API) から取得されます。
Apple Foundation モデル (システム、オンデバイス — 参照行)
Apple FM は単一のプリインストール モデルであるため、上記のオープンウェイト ランタイムと「論理モデル」行を共有できません。基準点として独自のラインを獲得します。これは、「システム モデルのみを使用する」という選択肢がある場合に達成すべき数値です。
注意事項 — 比較する前にお読みください。
トークンは推定 ( utf8.count / 4 ) されます。

[切り捨てられた]

## Original Extract

Neutral, reproducible benchmark for local LLMs on Apple Silicon (Mac · iPhone · iPad) — MLX, llama.cpp, CoreML, Apple Foundation Models - john-rocky/apple-silicon-llm-bench

GitHub - john-rocky/apple-silicon-llm-bench: Neutral, reproducible benchmark for local LLMs on Apple Silicon (Mac · iPhone · iPad) — MLX, llama.cpp, CoreML, Apple Foundation Models · GitHub
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
john-rocky
/
apple-silicon-llm-bench
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
john-rocky/apple-silicon-llm-bench
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits .github .github apple/ YardstickCLI/ Sources apple/ YardstickCLI/ Sources devices devices docs/ charts docs/ charts ios ios methodology methodology models models prompts prompts results results runtimes runtimes scripts scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md LICENSE LICENSE Package.swift Package.swift README.md README.md RESULTS.md RESULTS.md View all files Repository files navigation
On-device LLM benchmark for Apple Silicon — iPhone · iPad · Mac.
A neutral, reproducible benchmark for running local LLMs (and, in time, ASR / TTS) on Apple Silicon. Compares MLX Swift, llama.cpp, CoreML (swift-transformers), LiteRT-LM, ExecuTorch, ANEMLL — and Apple's own Foundation Models — under real device constraints, not just tok/s on a server.
Repo: apple-silicon-llm-bench · CLI/brand: yardstick . Started life as ios-llm-benchmark — iPhone is still the headline target, now measured alongside iPad and Mac.
📱 TL;DR — iPhone 17 Pro (A19 Pro)
Real LLM inference on a phone — on-device, no server. iPhone 17 Pro, 4-bit, short-chat (128 tokens), median of 3 cold runs. The winning runtime is model-dependent — and the upset is on Gemma.
Decode throughput — tok/s, higher is better (🏆 = winner):
Peak memory — MB, lower is better (🏆 = winner):
The upset — Gemma 4 E2B: Google's LiteRT-LM (INT4-QAT, GPU, its native .litertlm ) beats MLX-Swift on decode and uses ~4.5× less memory (641 MB vs 2,900). The purpose-built runtime wins on its own format.
MLX-Swift wins Qwen 3.5 2B decode — 61 vs 39 tok/s. (LiteRT-LM has no Qwen entry — its catalog is Gemma-only.)
CoreML / ANE is the memory champion — Qwen 3.5 2B in just 241 MB (~5× leaner than MLX's 1,279) via chunked-MLKV on the Neural Engine — but it's the slowest decode (ANE trades throughput for footprint), same story as on M4 Max.
ANE is near-parity with the desktop: CoreML Gemma 4 E2B does 33 tok/s on iPhone vs 32.5 on M4 Max — same silicon family. The GPU runtimes pay the real on-device tax: ~4–5× slower than M4 Max (Qwen 3.5 2B → 61 tok/s vs 292).
Counting: MLX / llama.cpp / LiteRT-LM report exact tokenizer tokens (LiteRT-LM via getBenchmarkInfo ); CoreML/ANE counts streamed pieces (≈ tokens). LiteRT-LM runs to EOS (no per-call cap → ~458-tok reply vs the others' 128 budget); decode tok/s is a rate, so the head-to-head holds.
Fully automated, side-loaded via devicectl headless mode — nothing typed on the phone, same methodology as the desktop rows.
Coming next: Apple Foundation Models, more models and more iPhones / iPads. One row is a great PR .
How the LiteRT-LM row was measured: google-ai-edge/LiteRT-LM 0.12.0 running litert-community/gemma-4-E2B-it.litertlm (INT4-QAT) on the Metal GPU backend, via the in-tree MediaPipeRuntime adapter — same headless harness + prompt as every other row (3 cold runs, median). Token counts and tok/s come from LiteRT-LM's own benchmark counters ( Conversation.getBenchmarkInfo ), so they're exact, not estimated. It generates to EOS (no per-call output cap in the API), so its token count is the model's full reply rather than the 128-token budget — decode tok/s is a rate and stays comparable; memory is exact process RSS. LiteRT-LM is vendored as a local SwiftPM package ( scripts/bootstrap.sh clones it with GIT_LFS_SKIP_SMUDGE=1 ; the released package trips SwiftPM's unsafe-flags rule via its -all_load ).
How the CoreML/ANE rows were measured: john-rocky/CoreML-LLM on the Neural Engine ( computeUnits: .cpuAndNeuralEngine ) — Gemma 4 E2B via the chunked .mlmodelc path, Qwen 3.5 2B via Qwen35MLKVGenerator (chunked MLKV, hence the 241 MB). Decode counts streamed pieces (≈ tokens); first-load ANE compilation makes its load time high (and it's the lowest-throughput runtime — the ANE trades speed for memory).
Decode tok/s is the headline number; the full per-run audit (prefill, TTFT, inter-token jitter, memory) lives in RESULTS.md .
⏱ Burst tok/s is only half the story — sustained throttling
The table above is cold-burst speed. Run the same model continuously and it flips: the GPU runtimes (MLX, LiteRT-LM) heat up and throttle 50%+ within ~60 s , while the ANE barely moves — it draws ~half the power, so it heats slowly and the SoC doesn't throttle it.
Two independent GPU runtimes collapsing the same way is a GPU-thermal property of the phone, not a runtime quirk. MLX ends up below the ANE; LiteRT keeps only a slim lead after shedding half its speed. The GPU wins the sprint; the ANE wins the marathon — and it frees the GPU for the rest of the app.
Method: 600 s continuous generation, cold ( nominal ) start, unplugged, tg128; decode rate from a rolling window. Raw JSONL in results/raw/iphone17pro-*-energy-tg128.jsonl ; redraw with scripts/throttle_chart.py (curves table via scripts/throttle_curve.py ). LiteRT-LM has no output-token cap (longer per-call) and that run started at fair thermal; CoreML-LLM uses sliding-window attention (bounded context), part of why its decode stays flat.
🖥 Desktop reference — Apple M4 Max
The same harness on a laptop-class chip, for scale. No runtime wins everything here — each optimises a different corner of the throughput / memory / energy / streaming box:
mlx-swift wins decode throughput on every cell measured (1.4×–1.8× over llama.cpp after early-2026 kernel updates).
Apple Foundation Models is 2× more energy-efficient per token than the GPU-backed runtimes, 4× more than CoreML/ANE.
CoreML / ANE wins peak memory (chunked MLKV) but is the slowest and the worst on J/token.
llama.cpp sits in the middle on speed and energy — no axis it wins, no axis it loses badly.
Regenerate after adding rows: python scripts/generate_charts.py .
📊 Full numbers — Apple M4 Max, short-chat (128 tokens, decode tok/s, median)
One device, four runtimes, multiple models. Decode tok/s is the primary headline number; the full table (prefill, TTFT, peak memory, per-run audit trail) lives in RESULTS.md . Read the Headline observations section before drawing conclusions — the runtime ranking is model-size-dependent .
Cross-runtime — same logical model, different backends (decode tok/s, median)
Logical model
Params
n
mlx-swift (Q4)
llama.cpp (Q4_K_M)
coreml-llm
litert-lm (.litertlm)
Qwen 2.5 0.5B
0.5 B
3
531.1
297.1
181.2 (FP16)
n/a
Qwen 3.5 0.8B
0.8 B
3
421.1
201.1
58.2 (INT8)
n/a
Qwen 3.5 2B
2 B
3
291.9
149.7
35.0 (INT8)
n/a
Gemma 4 E2B
2 B
3
185.4
119.2
32.5 (INT4 palettized)
pending
Gemma 4 E4B
4 B
3
113.5
80.5
not run
pending
litert-lm column: pending = adapter wired against google-ai-edge/LiteRT-LM v0.12.0, M4 Max run not yet captured (see RESULTS.md / Yardstick_USER_RUNS.md ). n/a = LiteRT-LM's catalog is Gemma-only ( .litertlm ), so the Qwen rows have no entry. For reference, Google's E2B model card reports 56.5 tok/s on iPhone 17 Pro GPU — a vendor figure on a different device, not an M4 Max Yardstick measurement.
→ MLX-Swift now wins decode on every cell — 1.4×–1.8× over llama.cpp — after upstream mlx-swift-lm shipped Qwen + Gemma kernel updates in early 2026 (the Qwen rows roughly tripled vs. the snapshot captured before those landed). The old "llama.cpp Metal always wins small-model decode" rule is no longer true on M4 Max; re-measure before quoting it. CoreML / ANE is the slowest of the three on every cell, in exchange for the dramatic memory savings shown below.
Cross-runtime — peak memory (MB, median)
The decode-tok/s table above hides the memory side. Same models, looking at peak working-set instead:
→ "CoreML/ANE wins memory" is true once the chunked MLKV layout kicks in. At 0.5 B params MLX-Swift is still smaller (413 MB vs CoreML's 959 MB monolithic FP16); from 0.8 B onward, CoreML's chunked MLKV path ( Qwen35MLKVGenerator : mmap'd embed sidecar + on-demand ANE chunks) holds the process RSS roughly flat — 206 MB at 0.8 B, 215 MB at 2 B — while MLX and llama.cpp scale linearly with parameter count.
Cross-runtime — energy per token (Gemma 4 E2B, sustained-512, M4 Max)
The number nobody else publishes: how many joules does each backend burn per generated token? Captured via scripts/measure_energy.py which co-runs powermetrics (whole-system, package power = CPU + GPU + ANE) and clips the sample window to the bench's reported active time.
The ANE path draws ~half the GPU path's package power at full decode (12.7 W vs ~24.7 W) — the same power gap that makes the GPU runtimes thermally throttle on iPhone while the ANE holds its rate (see the sustained-throttle section above).
→ Energy ranking inverts the decode-tok/s ranking. Apple FM is 2× more efficient per token than the GPU-backed runtimes despite producing tokens at ~half the rate. CoreML/ANE has the lowest instantaneous power (12.7 W) but is the worst J/tok at 4× Apple FM, because the slower decode (32 tok/s) keeps the package powered up much longer. MLX-Swift and llama.cpp draw the most W (GPU) but produce tokens fast enough to break even at ~0.24 J/tok. Whole-system measurement includes the idle baseline so all four numbers slightly inflate per-token energy — useful for ranking, not for absolute attribution. iPhone energy uses the 1 %-battery-step API instead (different methodology, similar table shape).
llama.cpp (Q4_K_M GGUF, M4 Max, short-chat)
mlx-swift (Q4 / MLX, M4 Max, short-chat)
coreml-llm (CoreML / ANE, M4 Max, short-chat)
→ CoreML/ANE trades throughput for memory: 3-8× less peak working set than MLX-Swift / llama.cpp at the same model size, at ~half the decode tok/s. The Qwen 3.5 0.8B / 2B numbers come from the dedicated Qwen35MLKVGenerator (ANE chunked decode, KV in MLState — public API since CoreML-LLM v1.9.0 ), not the generic CoreMLLLM.load(from:) path.
Apple Foundation Models (system, on-device — reference row)
Apple FM is a single pre-installed model, so it can't share a "logical model" row with the open-weight runtimes above. It earns its own line as a reference point — the number to beat when "just use the system model" is the alternative.
Caveats — read before comparing.
Tokens are estimated ( utf8.count / 4 ) because

[truncated]
