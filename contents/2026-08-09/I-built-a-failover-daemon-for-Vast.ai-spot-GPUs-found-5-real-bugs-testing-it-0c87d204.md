---
source: "https://github.com/enplabs/spotwarp"
hn_url: "https://news.ycombinator.com/item?id=49228960"
title: "I built a failover daemon for Vast.ai spot GPUs, found 5 real bugs testing it"
article_title: "GitHub - enplabs/spotwarp: Zero-Downtime Spot GPU Failover Guard & Auto-Resumer · GitHub"
author: "choi5844"
captured_at: "2026-08-09T06:38:09Z"
capture_tool: "hn-digest"
hn_id: 49228960
score: 1
comments: 0
posted_at: "2026-08-09T06:36:00Z"
tags:
  - hacker-news
  - translated
---

# I built a failover daemon for Vast.ai spot GPUs, found 5 real bugs testing it

- HN: [49228960](https://news.ycombinator.com/item?id=49228960)
- Source: [github.com](https://github.com/enplabs/spotwarp)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T06:36:00Z

## Translation

タイトル: Vast.ai スポット GPU 用のフェイルオーバー デーモンを構築しましたが、それをテストしたところ 5 つの実際のバグが見つかりました
記事のタイトル: GitHub - enplabs/spotwarp: ゼロ ダウンタイム スポット GPU フェイルオーバー ガードと自動レジューマー · GitHub
説明: ゼロダウンタイムのスポット GPU フェイルオーバー ガードと自動レジューマ - enplabs/spotwarp

記事本文:
GitHub - enplabs/spotwarp: ゼロダウンタイムスポット GPU フェイルオーバーガードと自動レジューマー · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
enplabs
/
spotwarp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット .github/ workflows .github/ workflows static static templates templates .gitignore .gitignore Dockerfile Dockerfile PRODUCT_STRA

TEGY.md PRODUCT_STRATEGY.md README.md README.md TECHNICAL_ARTICLE.md TECHNICAL_ARTICLE.md app_gpuaction.py app_gpuaction.pyentrypoint.shentrypoint.sh gpu-action.conf gpu-action.conf gpu_action_cli.py gpu_action_cli.py runpod_connector.py runpod_connector.py setup.py setup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
⚡ SpotWarp: スポット GPU エビクションによってトレーニング実行が失われることはありません
スポット GPU エビクションの実際のコストは数分間のダウンタイムではなく、何時間ものトレーニングの進捗がそれとともに消えていきます。 SpotWarp は、バックグラウンドでワークスペースの継続的な自動バックアップを実行する軽量の 100% ローカル Python デーモンであるため、エビクションによって作業が犠牲になることはありません。 1 分未満のクロスクラウド フェイルオーバー (Vast.ai ⇄ RunPod) により、保護されたワークスペースが手動で回復できるようになりますが、実際にユーザーを救うのはバックアップです。
オンデマンドではなくスポット価格を安全に使用することで、GPU コンピューティング料金を最大 70% 節約できます。通常ならギャンブルになる立ち退きのリスクが、まさに SpotWarp によって取り除かれます。
🆚 スポット GPU エビクション: 標準と SpotWarp
特徴
標準スポットインスタンス
スポットワープ (v3.2) あり
あなたのデータ、立ち退き時
消えた。手動で保存されなかったものはインスタンスとともに失われます。
エビクションが行われる前にバックグラウンドで継続的にバックアップされるため、失うものは何もありません。
回復プロセス
コンソールへの手動ログイン、新しい GPU の検索、手動セットアップ。
100% 自動操縦。並行する候補者レーシングは、1 分以内に代替品をレンタルして確認します。
クラウドの在庫がない場合
フェイルオーバーは完全に失敗し、移行先がありません。
クロスクラウドブリッジ。フォールバックとして RunPod で自動的にレンタルし、スポット料金が最初に適用され、スポットが提供されない場合はオンデマンドでレンタルされます。
ブリッジクラウド料金を永久に支払う
該当なし
自動フェイルバック。元のクラウドのより安価な容量が戻ってくるのを監視します

そして自動的に元に戻ります。ブリッジ クラウドは決して永続的なホームではありません。
ワークロードの継続
エポック 0 からトレーニングを再開します。
自動再開。スクリプトは、バックアップが中断されたところから、SSH 経由で nohup 経由で実行を継続します。
セキュリティリスク
不安定なレンタルホストに S3/GitHub キーを配置する必要があります。
キー漏れゼロ。すべての API キーはローカル マシンに残ります。
💎 主要な商用機能
🛡️ 継続的自動バックアップ — 本当のセーフティ ネット : エビクションが検出された後だけでなく、インスタンスの実行中ずっとバックグラウンドでワークスペースのローカル マシンへの高速 rsync / scp 増分バックアップを実行します。これは実際に損失を防ぐ機能です。以下に挙げるその他のことはすべて、手動で迅速に回復できるようにするだけです。
🏁 並列候補レーシング : エビクションの際、SpotWarp は一度に 1 つずつ試すのではなく、複数の代替候補を同時にレンタルします。単一の遅いホストや停止したホストによってダウンタイムが数分増加することはもうありません。通常のフェイルオーバー: 1 分未満です。
🌐 クロスクラウド フォールバック (Vast.ai ⇄ RunPod) : エビクションの時点でプライマリ クラウドに一致する候補がない場合、SpotWarp は自動的に RunPod にブリッジします。最初はスポット料金設定で、RunPod にその GPU モデルのスポット容量がない場合はオンデマンドで再試行されます。そのため、ワークロードは完全に失敗するのではなく保護されたままになります。
↩️ 自動コスト最適化フェイルバック : ブリッジクラウドの置き換えが、より高いレートで無期限に実行され続けることはありません。 SpotWarp は元のクラウドをバックグラウンドでチェックし続け、より安価な一致する候補が再び現れた瞬間にワークロードを元に戻します。単一のライブ インスタンスでエンドツーエンドで検証されます。Vast.ai でレンタルされ、削除され、RunPod にブリッジされ、容量が戻ったら自動的に Vast.ai に戻ります。
💸CF

O-Approved GPU Savings : Vast.ai の安価なスポット価格を安全に活用します。 SpotWarp は、スポット インスタンスの価格で専用オンデマンド GPU の信頼性を提供します。
✅ 実際の接続性検証 : 代替ホストは、信頼される前に、実際の SSH ハンドシェイクによって到達可能であることが確認されます。完全に正常なホストでの偽陰性を報告できる Jupyter API ping のようなプロキシ信号ではありません。
🔒 ゼロトラスト セキュリティ (100% ローカル) : クラウド プロバイダーの API キー ( VAST_API_KEY 、 RUNPOD_API_KEY ) はローカル PC に残ります。レンタルされたコンテナーがクラウドの認証情報を参照することはありません。
📦 ゼロ構成 : リモート コンテナー内にデーモン、cron ジョブ、または同期ツールをインストールする必要はありません。
⚡ トレーニング自動再開 : 新しいコンテナーのバックグラウンドでトレーニング スクリプト ( --resume-cmd ) を自動的に再開し、復元され、バックアップされたチェックポイントを直接指定します。
Pip 経由で公式パッケージをインストールします。
pip スポットワープをインストールする
2. APIキーのエクスポートとローカル環境のセットアップ
Vast.ai API キーをローカル マシンに設定します。クロスクラウド ブリッジが必要な場合は、RunPod API キーも追加します (オプションですが推奨。これにより、Vast.ai の在庫が少ない日がダウンタイムになるのを防ぐことができます)。
# Linux/macOS の場合
import VAST_API_KEY= " your_vast_api_key "
import RUNPOD_API_KEY= " your_runpod_api_key " # オプション、クロスクラウド フォールバック + 自動フェイルバックを有効にします
# Windows の場合 (PowerShell)
$env :VAST_API_KEY= " your_vast_api_key "
$env :RUNPOD_API_KEY= " your_runpod_api_key "
3. ガードを開始します (自動同期とトレーニング再開付き)
ローカル PC でガード デーモンを実行します。ライセンス キーを指定し、トレーニングを再開する方法を定義します。
Spotwarp start --license-key YOUR_SPOTWARP_KEY --resume-cmd " python /workspace/train.py --resume "
⚙️ 仕組み (ワーピングサイクル)
【レンタル中】

GPU ホスト] [ローカル PC (クライアント)] [新しい GPU ホスト]
(アクティブなワークロード)
│ │ │
│ ──── (同期: rsync/scp デルタ) ──────> │ (キャッシュされたワークスペース) │
│ │ │
[🚨退去!] │ │
❌ ──── (5 秒以内に検出) ───> │ │
│ ────（第4レース候補者） ──────> │
│ 準備ができているプールの最も安いものが勝ちます │
│ ────（ワークスペースの復元） ──────> │
│ ──── (Nohup 再開コマンド) ──────> [ワークロードの実行]
（続きます！）
Vast.ai に候補者がいない場合:
│ ────（RunPod へのブリッジ） ────────> [一時的な交換]
│ ⋯ バックグラウンドで Vast.ai を監視し続けます ⋯
│ ─── (膨大な AI の容量が戻ってきます) ──> [より安価に移行してください]
エビクション検出 : SpotWarp は 5 秒ごとにクラウド API をポーリングします。エビクションが検出されると、すぐにフェイルオーバーがトリガーされます。
Parallel Candidate Racing : SpotWarp は、オファーを 1 つずつ試すのではなく、最大 4 つの一致する GPU 候補を同時にレンタルし、SSH 到達可能であることが証明された最も安価な候補を選択します。
クロスクラウド ブリッジ (必要な場合) : その時点で Vast.ai の候補が存在しない場合、SpotWarp は代わりに RunPod で自動的にレンタルします (最初はスポット料金設定、2 回目はオンデマンドでの試行)。そのため、ワークロードが完全に保護されないままになることはありません。
デルタ同期復元 : SpotWarp は、キャッシュされたワークスペース フォルダーを新しいコンテナーに転送し、実際の SSH 接続経由でアクセスできることを確認します。
Nohup ハンドオーバー : SpotWarp は SSH 経由で接続し、バックグラウンドで --resume-cmd をトリガーします。
自動フェイルバック : ステップ 3 が必要な場合、SpotWarp は静かにその GPU モデルの Vast.ai を再チェックし続けます。候補が再び表示されると、ワークロードが自動的に元に移行され、ブリッジ クラウド ホストが解放されます。手動による介入は必要ありません。

つまり、忘れられた高価なインスタンスが実行されたままになることはありません。
SpotWarp は MIT ライセンスに基づいて配布されます。コードはローカル コンピューターのユーザー空間で 100% ローカルに実行され、完全な透明性とコンプライアンスが保証されます。
ゼロダウンタイムのスポット GPU フェイルオーバー ガードと自動レジューマ
Readme アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Zero-Downtime Spot GPU Failover Guard & Auto-Resumer - enplabs/spotwarp

GitHub - enplabs/spotwarp: Zero-Downtime Spot GPU Failover Guard & Auto-Resumer · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
enplabs
/
spotwarp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits .github/ workflows .github/ workflows static static templates templates .gitignore .gitignore Dockerfile Dockerfile PRODUCT_STRATEGY.md PRODUCT_STRATEGY.md README.md README.md TECHNICAL_ARTICLE.md TECHNICAL_ARTICLE.md app_gpuaction.py app_gpuaction.py entrypoint.sh entrypoint.sh gpu-action.conf gpu-action.conf gpu_action_cli.py gpu_action_cli.py runpod_connector.py runpod_connector.py setup.py setup.py View all files Repository files navigation
⚡ SpotWarp: Never Lose a Training Run to a Spot GPU Eviction
The real cost of a Spot GPU eviction was never the few minutes of downtime — it's the hours of training progress that vanish with it. SpotWarp is a lightweight, 100% local Python daemon that runs continuous automatic backups of your workspace in the background, so an eviction never costs you your work. Sub-minute cross-cloud failover (Vast.ai ⇄ RunPod) is what turns that protected workspace into a hands-off recovery, but the backup is the part that actually saves you.
Save up to 70% on GPU compute bills by safely using Spot pricing instead of on-demand — the eviction risk that normally makes that a gamble is exactly what SpotWarp removes.
🆚 Spot GPU Eviction: Standard vs. SpotWarp
Feature
Standard Spot Instance
With SpotWarp (v3.2)
Your Data, on Eviction
Gone. Whatever wasn't manually saved is lost with the instance.
Continuously backed up in the background before the eviction ever happens — nothing to lose.
Recovery Process
Manual console log-in, search for a new GPU, manual setup.
100% Autopilot . Parallel candidate racing rents & verifies a replacement in under a minute.
If your cloud is out of stock
Failover fails outright — nothing to migrate to.
Cross-Cloud Bridge . Automatically rents on RunPod as a fallback, spot pricing first, on-demand if spot isn't offered.
Paying bridge-cloud rates forever
N/A
Auto-Failback . Watches for your original cloud's cheaper capacity to return and migrates you back automatically — the bridge cloud is never a permanent home.
Workload Continuation
Restart training from epoch 0.
Auto-Resume . Script continues running via nohup over SSH, from where the backup left off.
Security Risk
Requires placing S3/GitHub keys on unstable rented hosts.
Zero Key Leakage . All API keys remain on your local machine.
💎 Core Commercial Features
🛡️ Continuous Automatic Backup — the real safety net : Runs high-speed rsync / scp incremental backups of your workspace to your local machine in the background, the whole time your instance is running — not just triggered after an eviction is detected. This is the feature that actually prevents loss; everything else below just makes recovering from it fast and hands-off.
🏁 Parallel Candidate Racing : On eviction, SpotWarp rents several replacement candidates concurrently instead of trying them one at a time — a single slow or dead host no longer adds minutes to your downtime. Typical failover: well under a minute.
🌐 Cross-Cloud Fallback (Vast.ai ⇄ RunPod) : If your primary cloud has zero matching candidates at the moment of eviction, SpotWarp automatically bridges to RunPod — spot pricing first, retrying on-demand if RunPod has no spot capacity for that GPU model — so your workload stays protected instead of failing outright.
↩️ Automatic Cost-Optimizing Failback : A bridge-cloud replacement is never left running indefinitely at the higher rate. SpotWarp keeps checking your original cloud in the background and migrates the workload back the instant a cheaper matching candidate reappears — verified end-to-end on a single live instance: rented on Vast.ai, evicted, bridged to RunPod, then automatically migrated back to Vast.ai once capacity returned.
💸 CFO-Approved GPU Savings : Safely exploit cheap Spot pricing on Vast.ai. SpotWarp gives you the reliability of a Dedicated On-Demand GPU for the price of a Spot instance.
✅ Real Connectivity Verification : Replacement hosts are confirmed reachable via an actual SSH handshake before they're trusted — not a proxy signal like a Jupyter API ping that can report false negatives on a perfectly healthy host.
🔒 Zero-Trust Security (100% Local) : Your cloud provider API keys ( VAST_API_KEY , RUNPOD_API_KEY ) stay on your local PC. Rented containers never see your cloud credentials.
📦 Zero-Configuration : No need to install daemons, cron jobs, or synchronization tools inside the remote container.
⚡ Training Auto-Resumer : Automatically restarts your training scripts ( --resume-cmd ) in the background of the new container, pointing directly to your restored, backed-up checkpoints.
Install the official package via Pip:
pip install spotwarp
2. Export API Keys & Set Up Local Environment
Set your Vast.ai API key on your local machine. Add a RunPod API key too if you want the cross-cloud bridge (optional, but recommended — it's what keeps a bad-inventory day on Vast.ai from becoming downtime):
# On Linux/macOS
export VAST_API_KEY= " your_vast_api_key "
export RUNPOD_API_KEY= " your_runpod_api_key " # optional, enables cross-cloud fallback + auto-failback
# On Windows (PowerShell)
$env :VAST_API_KEY= " your_vast_api_key "
$env :RUNPOD_API_KEY= " your_runpod_api_key "
3. Start the Guard (With Auto-Sync & Training Resume)
Run the guard daemon on your local PC. Point it to your license key and define how your training should resume:
spotwarp start --license-key YOUR_SPOTWARP_KEY --resume-cmd " python /workspace/train.py --resume "
⚙️ How It Works (The Warping Cycle)
[Rented GPU Host] [Local PC (Client)] [New GPU Host]
(Active Workload)
│ │ │
│ ─── (Sync: rsync/scp delta) ──────> │ (Cached Workspace) │
│ │ │
[🚨 Evicted!] │ │
❌ ─── (Detected within 5s) ────────> │ │
│ ─── (Race 4 candidates) ──────────> │
│ cheapest of the ready pool wins │
│ ─── (Restore Workspace) ──────────> │
│ ─── (Nohup Resume Command) ──────> [Run Workload]
(Continuing!)
If Vast.ai has zero candidates:
│ ─── (Bridge to RunPod) ────────────> [Temp Replacement]
│ ⋯ keeps watching Vast.ai in the background ⋯
│ ─── (Vast.ai capacity returns) ──> [Migrate back, cheaper]
Eviction Detection : SpotWarp polls the cloud API every 5 seconds. If eviction is detected, it triggers failover immediately.
Parallel Candidate Racing : SpotWarp rents up to 4 matching-GPU candidates concurrently and picks the cheapest one that proves SSH-reachable, instead of trying offers one at a time.
Cross-Cloud Bridge (if needed) : If no Vast.ai candidate exists at that moment, SpotWarp automatically rents on RunPod instead — spot pricing first, on-demand as a second attempt — so the workload is never left completely unprotected.
Delta Sync Restoration : SpotWarp transfers the cached workspace folder to the new container and confirms it's reachable over a real SSH connection.
Nohup Handover : SpotWarp connects via SSH to trigger the --resume-cmd in the background.
Auto-Failback : If step 3 was needed, SpotWarp keeps quietly re-checking Vast.ai for that GPU model. The moment a candidate reappears, it migrates the workload back automatically and releases the bridge-cloud host — no manual intervention, no forgotten expensive instance left running.
SpotWarp is distributed under the MIT License. The code executes 100% locally in user space on your local computer, ensuring full transparency and compliance.
Zero-Downtime Spot GPU Failover Guard & Auto-Resumer
Readme Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
