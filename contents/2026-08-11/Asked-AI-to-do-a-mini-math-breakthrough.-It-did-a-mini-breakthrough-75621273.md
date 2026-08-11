---
source: "https://github.com/learademacher/ai-refines-ai-zeta-bound"
hn_url: "https://news.ycombinator.com/item?id=49264616"
title: "Asked AI to do a mini math breakthrough. It did a mini breakthrough"
article_title: "GitHub - learademacher/ai-refines-ai-zeta-bound: AI Refines AI: a reproducible candidate improvement to the 67.25% zeta-zero bound · GitHub"
author: "learademacher"
captured_at: "2026-08-11T21:35:36Z"
capture_tool: "hn-digest"
hn_id: 49264616
score: 3
comments: 1
posted_at: "2026-08-11T21:17:08Z"
tags:
  - hacker-news
  - translated
---

# Asked AI to do a mini math breakthrough. It did a mini breakthrough

- HN: [49264616](https://news.ycombinator.com/item?id=49264616)
- Source: [github.com](https://github.com/learademacher/ai-refines-ai-zeta-bound)
- Score: 3
- Comments: 1
- Posted: 2026-08-11T21:17:08Z

## Translation

タイトル: AI にミニ数学のブレークスルーを依頼しました。小さな進歩を遂げました
記事タイトル: GitHub - learademacher/ai-refines-ai-zeta-bound: AI Refines AI: 67.25% ゼータゼロ境界への再現可能な候補改善 · GitHub
説明: AI Refines AI: 67.25% ゼータゼロ境界への再現可能な候補改善 - learademacher/ai-refines-ai-zeta-bound

記事本文:
GitHub - learademacher/ai-refines-ai-zeta-bound: AI Refines AI: 67.25% ゼータゼロ境界への再現可能な改善候補 · GitHub
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
リアデマッチャー
/
ai-refines-ai-zeta-bound
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github/ workflows .github/ workflows 証明書 証明書 ドキュメント ドキュメント 紙 紙 スクリプト スクリプト src/ zeta_simple_zeros

src/ zeta_simple_zeros テスト テスト .dockerignore .dockerignore .gitignore .gitignore .python-version .python-version AUDIT.md AUDIT.md CITATION.cff CITATION.cff Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile PROVENANCE.md PROVENANCE.md README.md README.md REPRODUCIBILITY.md REPRODUCIBILITY.md pyproject.toml pyproject.toml要件.ロック要件.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
67.25% のゼータゼロ境界に対する再現可能な改善候補。
1 つの AI システムが Anthropic の新しい定理 D を生成しました。2 つ目の AI システムが生成しました
研究草案では小さな強化が発見されました。このリポジトリは、
より強力な引数、その正確な区間演算検証器、およびクリーンルーム
唯一のコンピュータ支援補題の再現。
$$
\liminf_{T\to\infty}\frac{N_0^s(T,2T)}{N(T,2T)}
\ge 0.6730213619501665335\ldots。
$$
ここで $N(T,2T)$ は自明でないゼロを多重度でカウントしますが、
$N_0^s(T,2T)$ はクリティカル ライン上の単純なゼロをカウントします。
これは未レビューの候補の絞り込みです。リーマンの証明にはならない
仮説であり、分析結果を独立して置き換えるものではありません
人間の定理 D からインポートされました。有限区間証明書には、
正確に再現されています。より広範な専門家のレビューが依然として必要です。
論文(PDF) · LaTeXソース ·
技術監査・再現ガイド
強化されたソースはコミットで分離されます
0faebf367ffb780951845d014e2e3d9a4a04adb1 。
公開検証者から始めて、ターゲットを $191/50000$ に変更し、
45,840 グリッド セルの一致する圧力カットオフを導き出します。
CPython 3.12.3 および python-flint==0.8.0 で実行されるクリーンな x86_64 Linux
返されました:
検証済み=真
ターゲット=F6 >= 191/50000
グリッド=4000
精度ビット=128
初期ボックス=729
ノード=786215
剪定済み=393472
分割=392743
最大深さ=43
kernel_table_sha256=f79a9147ffc37692b84330e984

22fb94cbf601c51d4f9f8ae749982f23838571
Second_derivative_table_sha256=cc98102590dba6e1a982a5a4c9fcd755848a93483db67edd6634813a35e5e3f5
期待される正確なレポートは、次の場所にコミットされます。
証明書/seven-point.expected.json 。
CI はクリーンなチェックアウトから徹底的なベリファイアを再実行し、すべての結果を比較します。
決定論的なフィールド。
$$
H_{\mathrm{MT}}
=\frac32-\frac1{\sqrt2}\cot\frac1{\sqrt2}
=0.672500703679\ldots。
$$
7 ポイントの改良により、インポートされた分析入力を使用して、次のことが証明されます。
$$
F_6(g_1,\ldots,g_6)\ge\frac{191}{50000}
$$
すべての非負のギャップに対して。ブロックサイズ $m=267$ の場合、最終的な控除は次のようになります。
$$
\frac{13{,}350{,}000H_{\mathrm{MT}}-26{,}600}{13{,}300{,}149}
=0.6730213619501665335\ldots。
$$
行列の不等式、カーネルの正規化、7 点の組み合わせ論、
シフトブロックのピンチと最後の算術演算については、
監査。分析トレースの推定値、テールバウンド、および最適化
テスト ファミリーは、引用された Anthropic 論文とその Lean 4 の依存関係が残ります。
アーティファクト。
最も正確なルートは、固定されたコンテナー イメージを使用します。
docker build -t ai-refines-ai-zeta-bound 。
docker run --rm ai-refines-ai-zeta-bound
ローカル CPython 3.12 インストールの場合:
python3 -m venv .venv
ソース .venv/bin/activate
python -m pip install --require-hashes -rrequirements.lock
PYTHONPATH=src python -m Unittest Discover -s テスト -v
PYTHONPATH=src python scripts/verify_release.py
完全な検証には通常、数分かかります。未解決
ターミナルセルはハード障害です。プログラムはそれを検証済みとして扱うことはありません。
上流のアーティファクト: ainta/zeta-simple-zeros 、
コミット 040c5e899e658aed7b56a2a87f501798fe10761d 。
上流リポジトリは、研究草案が生成されたものであることを識別します。
GPT-5.6ソル。
人類財団: 研究論文 、
紙、
そしてリーン4アーティファクト。
このリリースではアップストリーム MI が維持されます

Tライセンスと歴史。参照
正確なチェーンの PROVENANCE.md。
AI によって生成された研究草案。コンピュータ支援の補題は独立して再現されました。
人間による査読は主張されていません。議論を改ざんしたり再現しようとしたりする
リポジトリが確立されたら、別のアーキテクチャ上の証明書も歓迎されます。
公開されました。
MIT、上流のアーティファクトから継承。
AI による AI の洗練: 67.25% のゼータゼロ境界に対する再現可能な候補の改善
Readme MIT ライセンス このリポジトリを引用する アクティビティのスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI Refines AI: a reproducible candidate improvement to the 67.25% zeta-zero bound - learademacher/ai-refines-ai-zeta-bound

GitHub - learademacher/ai-refines-ai-zeta-bound: AI Refines AI: a reproducible candidate improvement to the 67.25% zeta-zero bound · GitHub
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
learademacher
/
ai-refines-ai-zeta-bound
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github/ workflows .github/ workflows certificates certificates docs docs paper paper scripts scripts src/ zeta_simple_zeros src/ zeta_simple_zeros tests tests .dockerignore .dockerignore .gitignore .gitignore .python-version .python-version AUDIT.md AUDIT.md CITATION.cff CITATION.cff Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile PROVENANCE.md PROVENANCE.md README.md README.md REPRODUCIBILITY.md REPRODUCIBILITY.md pyproject.toml pyproject.toml requirements.lock requirements.lock View all files Repository files navigation
A reproducible candidate improvement to the 67.25% zeta-zero bound.
One AI system produced Anthropic's new Theorem D. A second AI-generated
research draft found a small strengthening. This repository packages the
stronger argument, its exact interval-arithmetic verifier, and a clean-room
reproduction of the only computer-assisted lemma.
$$
\liminf_{T\to\infty}\frac{N_0^s(T,2T)}{N(T,2T)}
\ge 0.6730213619501665335\ldots.
$$
Here $N(T,2T)$ counts nontrivial zeros with multiplicity, while
$N_0^s(T,2T)$ counts simple zeros on the critical line.
This is an unreviewed candidate refinement. It does not prove the Riemann
hypothesis and it does not independently replace the analytic results
imported from Anthropic's Theorem D. The finite interval certificate has
been reproduced exactly; broader expert review is still needed.
Paper (PDF) · LaTeX source ·
technical audit · reproduction guide
The strengthened source is isolated in commit
0faebf367ffb780951845d014e2e3d9a4a04adb1 .
Starting from the public verifier, it changes the target to $191/50000$ and
derives the matching pressure cutoff of 45,840 grid cells.
A clean x86_64 Linux run with CPython 3.12.3 and python-flint==0.8.0
returned:
verified=true
target=F6 >= 191/50000
grid=4000
precision_bits=128
initial_boxes=729
nodes=786215
pruned=393472
splits=392743
maximum_depth=43
kernel_table_sha256=f79a9147ffc37692b84330e98422fb94cbf601c51d4f9f8ae749982f23838571
second_derivative_table_sha256=cc98102590dba6e1a982a5a4c9fcd755848a93483db67edd6634813a35e5e3f5
The exact expected report is committed in
certificates/seven-point.expected.json .
CI reruns the exhaustive verifier from a clean checkout and compares every
deterministic field.
$$
H_{\mathrm{MT}}
=\frac32-\frac1{\sqrt2}\cot\frac1{\sqrt2}
=0.672500703679\ldots.
$$
The seven-point refinement proves, using those imported analytic inputs,
$$
F_6(g_1,\ldots,g_6)\ge\frac{191}{50000}
$$
for all nonnegative gaps. With block size $m=267$ , the final deduction is
$$
\frac{13{,}350{,}000H_{\mathrm{MT}}-26{,}600}{13{,}300{,}149}
=0.6730213619501665335\ldots.
$$
The matrix inequality, kernel normalization, seven-point combinatorics,
shifted-block pinching, and final arithmetic are covered in the
audit . The analytic trace estimates, tail bounds, and optimized
test family remain dependencies of the cited Anthropic paper and its Lean 4
artifact.
The most exact route uses the pinned container image:
docker build -t ai-refines-ai-zeta-bound .
docker run --rm ai-refines-ai-zeta-bound
For a local CPython 3.12 installation:
python3 -m venv .venv
source .venv/bin/activate
python -m pip install --require-hashes -r requirements.lock
PYTHONPATH=src python -m unittest discover -s tests -v
PYTHONPATH=src python scripts/verify_release.py
The exhaustive verification usually takes a few minutes. An unresolved
terminal cell is a hard failure; the program never treats it as verified.
Upstream artifact: ainta/zeta-simple-zeros ,
commit 040c5e899e658aed7b56a2a87f501798fe10761d .
The upstream repository identifies the research draft as generated by
GPT-5.6 Sol.
Anthropic foundation: research article ,
paper ,
and Lean 4 artifact .
This release preserves the upstream MIT license and history. See
PROVENANCE.md for the exact chain.
AI-generated research draft. Computer-assisted lemma independently reproduced.
No human peer review is claimed. Attempts to falsify the argument or reproduce
the certificate on another architecture are welcome once the repository is
made public.
MIT, inherited from the upstream artifact.
AI Refines AI: a reproducible candidate improvement to the 67.25% zeta-zero bound
Readme MIT license Cite this repository Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
