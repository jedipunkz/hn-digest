---
source: "https://github.com/Broyojo/llm_from_scratch"
hn_url: "https://news.ycombinator.com/item?id=48662756"
title: "LLM from Scratch: a small LLM running inside MIT's Scratch"
article_title: "GitHub - Broyojo/llm_from_scratch · GitHub"
author: "alexkranias"
captured_at: "2026-06-24T17:37:02Z"
capture_tool: "hn-digest"
hn_id: 48662756
score: 3
comments: 0
posted_at: "2026-06-24T17:03:04Z"
tags:
  - hacker-news
  - translated
---

# LLM from Scratch: a small LLM running inside MIT's Scratch

- HN: [48662756](https://news.ycombinator.com/item?id=48662756)
- Source: [github.com](https://github.com/Broyojo/llm_from_scratch)
- Score: 3
- Comments: 0
- Posted: 2026-06-24T17:03:04Z

## Translation

タイトル: LLM from Scratch: MIT の Scratch 内で実行される小さな LLM
記事タイトル: GitHub - Broyojo/llm_from_scratch · GitHub
説明: GitHub でアカウントを作成して、Broyojo/llm_from_scratch の開発に貢献します。

記事本文:
GitHub - Broyojo/llm_from_scratch · GitHub
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
ブロヨジョ
/
llm_from_scratch
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く Fo

フォルダとファイル
11 コミット 11 コミット アーティファクト アーティファクト docs docs llama2.c llama2.c llvm2scratch llvm2scratch patches patchesScratch_llama2Scratch_llama2 .gitignore .gitignore README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
llvm2scratch を使用して C 推論コードを Scratch ブロックにコンパイルすることにより、Scratch/TurboWarp 内で最小の llama2.c モデル ( stories260K ) を実行します。
すべてが機能している場合、スプライトは見慣れたオープニングの生成を開始します。
かつて、... (吹き出しにトークンごとにストリーミングされます)。
スクラッチ プロジェクト: https://scratch.mit.edu/projects/1277883263
このリポジトリは、再現性を高めるために 2 つのアップストリーム プロジェクトをツリー内でベンダーします。
llama2.c by Andrej Karpathy (MIT)。出典: llama2.c/ および llama2.c/LICENSE 。
Classfied3D (MIT) による llvm2scratch。出典: llvm2scratch/ および llvm2scratch/LICENSE 。
artifacts/ のモデル/トークナイザー アーティファクトは、llama2.c エコシステムから取得されます。
Scratch_llama2/build_stories260k_sprite3.py は次のようになります。
artifacts/stories260K.bin (最小の llama2.c チェックポイント)
artifacts/tok512.bin (トークナイザーの語彙)
重み行列を Q8_0 (グループ サイズ 4) に量子化し、4 つの signed int8 値を 1 つの u32 にパックします。
すべてを 1 つのスクラッチ リスト !stack にレイアウトします。
梱包重量 + グループごとのスケール
RoPE cos/sin テーブル (削減された SEQ_LEN 用)
ランタイムバッファ (x/xb/hb/q/att + KV キャッシュ)
これは、scratch_llama2/generated_layout.h を 1 から始まるインデックスのアドレスで !stack に書き込みます。
以下を使用して、scratch_llama2/llama2_scratch.c を LLVM IR (scratch_llama2/llama2_scratch.ll) にコンパイルします。
Clang --target=i386-none-elf (ポインタを 32 ビット整数として保持)
llvm2scratch を実行して LLVM IR を Scratch ブロックに変換し、.sprite3 および .sb3 出力をエクスポートします。
!!output (list) には、生成されたトークン ID が格納されます。
!!vocab(リスト)にはトークンピース(文字列)が格納されます。
!!text (変数) acumul

デコードされたテキストを食べる。スプライトはそれを継続的に言います。
!!resets (変数) は、コンパイラーがブロードキャストベースの「スタック リセット」をトリガーすると増分します (進行状況インジケーター + JS 呼び出しスタックのブローアップを回避します)。
!!status (変数) は、高レベルのステート マシン (パラメーターの編集... -> 実行中... -> 完了) を示します。
ui_* 変数を使用すると、TurboWarp/Scr​​atch UI からサンプリング/生成設定を調整できます。
uv (および Python >= 3.12。llvm2scratch にはそれが必要です)
# 使用可能な Python をまだ持っていない場合:
# uv Python インストール 3.12
#
# オプション: TurboWarp の安定性/パフォーマンスのためにスタックのリセット頻度を調整します。
# 低い = 安定性が高くなります (「最大呼び出しスタック サイズを超えました」に遭遇する可能性は低くなります) が、速度は遅くなります。
# 高い = 高速ですが、TurboWarp でクラッシュする可能性があります。
# MAX_BRANCH_RECURSION=200 がデフォルトです。
MAX_BRANCH_RECURSION=200 \\
# オプション: 生成するトークンの数 (上限)。デフォルトは 20 です。
# (<= SEQ_LEN、現在は 32 である必要があります。)
GEN_STEPS=20 \\
# llvm2scratch には Python >= 3.12 が必要です。 uv が古いシステム Python を選択するのを避けるために、`--python` を介して固定します。
uv run --python 3.12 --no-project --with-editable ./llvm2scratch pythonScratch_llama2/build_stories260k_sprite3.py
出力:
Scratch_llama2/stories260k_inference.sprite3 : スプライト、非表示のブロック (高速エディター/インポート)
Scratch_llama2/stories260k_inference_visible.sprite3 : スプライト、ブロックが表示されます (デバッグ)
Scratch_llama2/stories260k_inference_visible.sb3 : 可視スプライトのスタンドアロン プロジェクト ラッパー
Scratch_llama2/stories260k_inference_visible_scratch.sprite3 : スクラッチ互換のスプライト (TurboWarp 専用ブロックなし)
Scratch_llama2/stories260k_inference_visible_scratch.sb3 : Scratch 互換のスタンドアロン プロジェクト
Scratch_llama2/stories260k_inference_visible.sprite3 を TurboWarp にインポートします ([ファイル] -> [スプライトをアップロード] またはドラッグ アンド ドロップ)。
ui_* 変数を編集します ([変数] パネル)。
スペースを押します（またはクリックします）

スプライト）を開始します。
TurboWarpでscratch_llama2/stories260k_inference_visible.sb3を開きます(「ファイル」→「コンピュータからロード」)。
ステージ上のスライダー/モニターを使用してパラメータを編集します。
スペースを押して (またはスプライトをクリックして) 開始します。
!!ステータスの更新: パラメーターを編集... -> 実行中... -> 完了。
!!resets は定期的に増分します (長時間実行中の「まだ生きている」インジケーター)。
トークンが生成されると、スプライトはデコードされたテキストを吹き出し ( !!text ) にストリーミングします。
デバッグのために、生成されたトークン ID が !!output リストに追加されます。
ui_steps : 生成する最大トークン数 (<= 32)。
ui_temperture : 0 => 貪欲; >0 => サンプリング。
ui_top_k : 1 => 貪欲; >1 => トップ K サンプリング。
ui_top_p : (0, 1] の核カットオフ (無効にするには 1 を使用)。
ui_seed : 非ゼロ => 決定的; 0 => 開始時にランダムなシードを選択します。
ui_prompt_preset : 0 => BOS から開始; 1 => トークン接頭辞を強制します (デモ)。
Scratch_llama2/stories260k_inference_visible_scratch.sb3 (推奨)
Scratch は TurboWarp よりも大幅に遅く、TurboWarp のみの「ハッキングされたカウンター」ブロックをサポートしていません。
Scratch_llama2/llama2_scratch.c は推論専用であり、スクラッチの実行可能性を高めるために削減された SEQ_LEN を使用します。
llvm2scratch はここで販売されており、事前シード !stack といくつかの追加の IR パターンをサポートするようにパッチが適用されています。
公式 Scratch は、TurboWarp のハッキングされたカウンター オペコードをサポートしていません。 Scratch.mit.edu には *_scratch.* 出力を使用します。
注目すべき llvm2scratch パッチ (このプロジェクト用)
llama2_scratch.c を実行可能にする重要な変更は次のとおりです。
プリシードメモリ: エクスポート時に !stack を直接挿入することで、巨大な「イニシャライザ」スクリプトの生成をスキップします。
i8 ポインター算術修正: Clang はバイト オフセット (4/8/12/...) を使用して getelementptr i8 を発行しますが、「メモリ」はリスト インデックス化されています。 i8 GEP インデックスを 32 ビット セルにスケールバックします ( i8_gep_div=4 )。
スタックr

eset progress: オプションの !!resets カウンタを使用して、長時間実行中に VM がまだ動作していることを確認します (生成されたテキストの吹き出しを保持します)。
トークン ストリーミング: SB3_emit_token_dbl は、トークン ID を !!output に記録し、 !!vocab を通じてデコードし、 !!text に追加し、スプライトの吹き出しを継続的に更新します。
組み込みサポートの追加: Clang は llvm.umin/umax/smin/smax を発行できます。 llvm2scratch は、-O2 IR がコンパイルできるようにこれらを変換するようになりました。
@misc { andrews2026llm_from_scratch ,
著者 = { アンドリュース、デイビッド } 、
タイトル = { llm\_from\_scratch } 、
年 = { 2026 } 、
howpublished = { \\url{https://github.com/broyojo/llm_from_scratch} }
}
について
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

Contribute to Broyojo/llm_from_scratch development by creating an account on GitHub.

GitHub - Broyojo/llm_from_scratch · GitHub
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
Broyojo
/
llm_from_scratch
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits artifacts artifacts docs docs llama2.c llama2.c llvm2scratch llvm2scratch patches patches scratch_llama2 scratch_llama2 .gitignore .gitignore README.md README.md View all files Repository files navigation
Run the smallest llama2.c model ( stories260K ) inside Scratch/TurboWarp by compiling C inference code to Scratch blocks with llvm2scratch .
If everything is working, the sprite will start generating the familiar opening:
Once upon a time, ... (streamed into the speech bubble token-by-token).
Scratch project: https://scratch.mit.edu/projects/1277883263
This repo vendors two upstream projects in-tree for reproducibility:
llama2.c by Andrej Karpathy (MIT). Source: llama2.c/ and llama2.c/LICENSE .
llvm2scratch by Classfied3D (MIT). Source: llvm2scratch/ and llvm2scratch/LICENSE .
The model/tokenizer artifacts in artifacts/ come from the llama2.c ecosystem.
scratch_llama2/build_stories260k_sprite3.py reads:
artifacts/stories260K.bin (the smallest llama2.c checkpoint)
artifacts/tok512.bin (tokenizer vocabulary)
It quantizes the weight matrices to Q8_0 (group size 4) and packs 4 signed int8 values into one u32 .
It lays out everything into a single Scratch list !stack :
packed weights + per-group scales
RoPE cos/sin tables (for a reduced SEQ_LEN )
runtime buffers (x/xb/hb/q/att + KV cache)
It writes scratch_llama2/generated_layout.h with 1-indexed addresses into !stack .
It compiles scratch_llama2/llama2_scratch.c to LLVM IR ( scratch_llama2/llama2_scratch.ll ) using:
clang --target=i386-none-elf (keeps pointers as 32-bit ints)
It runs llvm2scratch to turn LLVM IR into Scratch blocks, then exports .sprite3 and .sb3 outputs.
!!output (list) stores generated token IDs.
!!vocab (list) stores token pieces (strings).
!!text (variable) accumulates decoded text; the sprite say s it continuously.
!!resets (variable) increments when the compiler triggers a broadcast-based “stack reset” (progress indicator + avoids JS call stack blowups).
!!status (variable) shows a high-level state machine ( Edit params... -> Running... -> Done. ).
ui_* variables let you adjust sampling/generation settings from TurboWarp/Scratch UI.
uv (and Python >= 3.12; llvm2scratch requires it)
# If you don't have a usable Python yet:
# uv python install 3.12
#
# Optional: tune stack reset frequency for TurboWarp stability/perf.
# Lower = more stable (less likely to hit "Maximum call stack size exceeded"), but slower.
# Higher = faster, but can crash in TurboWarp.
# MAX_BRANCH_RECURSION=200 is the default.
MAX_BRANCH_RECURSION=200 \\
# Optional: number of tokens to generate (upper bound). Defaults to 20.
# (Must be <= SEQ_LEN, currently 32.)
GEN_STEPS=20 \\
# llvm2scratch requires Python >= 3.12; pin via `--python` to avoid uv picking an older system Python.
uv run --python 3.12 --no-project --with-editable ./llvm2scratch python scratch_llama2/build_stories260k_sprite3.py
Outputs:
scratch_llama2/stories260k_inference.sprite3 : sprite, blocks hidden (fast editor/import)
scratch_llama2/stories260k_inference_visible.sprite3 : sprite, blocks visible (debug)
scratch_llama2/stories260k_inference_visible.sb3 : standalone project wrapper around the visible sprite
scratch_llama2/stories260k_inference_visible_scratch.sprite3 : Scratch-compatible sprite (no TurboWarp-only blocks)
scratch_llama2/stories260k_inference_visible_scratch.sb3 : Scratch-compatible standalone project
Import scratch_llama2/stories260k_inference_visible.sprite3 into TurboWarp ( File -> Upload sprite or drag/drop).
Edit ui_* variables (Variables panel).
Press space (or click the sprite) to start.
Open scratch_llama2/stories260k_inference_visible.sb3 in TurboWarp ( File -> Load from your computer ).
Use the sliders/monitors on the stage to edit params.
Press space (or click the sprite) to start.
!!status updates: Edit params... -> Running... -> Done.
!!resets increments periodically (a "still alive" indicator during long runs).
As tokens are generated, the sprite streams decoded text into its speech bubble ( !!text ).
For debugging, generated token IDs are appended to the !!output list.
ui_steps : max tokens to generate (<= 32).
ui_temperature : 0 => greedy; >0 => sampling.
ui_top_k : 1 => greedy; >1 => top-k sampling.
ui_top_p : nucleus cutoff in (0, 1] (use 1 to disable).
ui_seed : nonzero => deterministic; 0 => pick a random seed at start.
ui_prompt_preset : 0 => start from BOS; 1 => force the token prefix Once upon a time, (demo).
scratch_llama2/stories260k_inference_visible_scratch.sb3 (recommended)
Scratch is significantly slower than TurboWarp, and does not support TurboWarp-only “hacked counter” blocks.
scratch_llama2/llama2_scratch.c is inference-only and uses a reduced SEQ_LEN for Scratch feasibility.
llvm2scratch is vendored here and patched to support pre-seeding !stack and a few extra IR patterns.
Official Scratch does not support TurboWarp's hacked counter opcodes. Use the *_scratch.* outputs for scratch.mit.edu.
Notable llvm2scratch Patches (For This Project)
These are the key changes that made llama2_scratch.c viable:
Preseeded memory: skip generating huge “initializer” scripts by directly injecting !stack at export time.
i8 pointer arithmetic fix: clang emits getelementptr i8 using byte offsets (4/8/12/...), but our “memory” is list-indexed; we scale i8 GEP indices back into 32-bit cells ( i8_gep_div=4 ).
Stack reset progress: optional !!resets counter to confirm the VM is still working during long runs (we keep the speech bubble for generated text).
Token streaming: SB3_emit_token_dbl logs token IDs to !!output , decodes through !!vocab , appends into !!text , and continuously updates the sprite speech bubble.
Added intrinsic support: clang can emit llvm.umin/umax/smin/smax ; llvm2scratch now translates these so -O2 IR compiles.
@misc { andrews2026llm_from_scratch ,
author = { Andrews, David } ,
title = { llm\_from\_scratch } ,
year = { 2026 } ,
howpublished = { \\url{https://github.com/broyojo/llm_from_scratch} }
}
About
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
