---
source: "https://ihatedevops.com/"
hn_url: "https://news.ycombinator.com/item?id=49075464"
title: "Ihatedevops.com – because DevOps separates vibe coder from engineer"
article_title: "ihatedevops — devops best practices for Claude Code"
author: "connorelkin"
captured_at: "2026-07-27T21:02:38Z"
capture_tool: "hn-digest"
hn_id: 49075464
score: 1
comments: 1
posted_at: "2026-07-27T21:00:37Z"
tags:
  - hacker-news
  - translated
---

# Ihatedevops.com – because DevOps separates vibe coder from engineer

- HN: [49075464](https://news.ycombinator.com/item?id=49075464)
- Source: [ihatedevops.com](https://ihatedevops.com/)
- Score: 1
- Comments: 1
- Posted: 2026-07-27T21:00:37Z

## Translation

タイトル: Ihatedevops.com – DevOps はバイブコーダーとエンジニアを分離するため
記事のタイトル: ihatedevops — クロード コードの devops ベスト プラクティス
説明: 10 のアトミックな DevOps ベスト プラクティス (安全な Chainguard イメージ、ピン留めされた依存関係、最小権限 CI) をすべてのセッションに静かに挿入するクロード コード スキル。

記事本文:
🖥️ ">
dev@mac — ~/dev — zsh
➜ ~/dev フーミ
実際、devops を憎んでいない人。 Devopsはとにかく難しいです。
➜ ~/dev cat about.txt
イハテデヴォップス
GitHub でスターを付ける
➜ ~/dev ./install.sh # 快適さのレベルを選択してください
クロードコード
コーデックス
コピー＆ペースト
ワンライナー
クロード プラグイン マーケットプレイス zenconnor/ihatedevops を追加 && クロード プラグイン インストール ihatedevops@ihatedevops
コピー
公式プラグイン — アップデートを取得します。すでにクロードセッションに入っていますか？同じこと: /plugin マーケットプレイスに zenconnor/ihatedevops を追加し、次に /plugin install ihatedevops@ihatedevops。
git clone -- Depth 1 https://github.com/zenconnor/ihatedevops /tmp/ihd && mkdir -p ~/.agents/skills && cp -R /tmp/ihd/skills/devops-best-practices ~/.agents/skills/
コピー
同じスキル、オープン エージェント スキル標準 — コーデックスは ~/.agents/skills/ と読み取ります。プロジェクトごと: 代わりに、リポジトリの .agents/skills/ に cp を追加します。
git clone -- Depth 1 https://github.com/zenconnor/ihatedevops /tmp/ihd && mkdir -p ~/.claude/skills && cp -R /tmp/ihd/skills/devops-best-practices ~/.claude/skills/
コピー
すべてのステップが読み取り可能: クローン、mkdir、コピー。スクリプトは実行されません。 1 つのプロジェクトの場合のみ、代わりに .claude/skills/ に cp します。
カール -fsSL https://raw.githubusercontent.com/zenconnor/ihatedevops/main/install.sh |しー
コピー
最速。 1 つのフォルダーを ~/.claude/skills/ にインストールします。他には何もインストールしません。必要に応じて、最初に install.sh をお読みください。
➜ ~/dev ls practices/
デフォルトの Chainguard 基本イメージ — cgr.dev/chainguard/*、無料、~0 CVE
マルチステージビルド、非rootユーザー - ビルドツールは出荷されません
すべてをピン留めする — ロックファイル、イメージダイジェスト、コミット SHA によるアクション
最小権限の CI トークン — 明示的な権限: 内容: 読み取り
コード、画像、ログに秘密はありません — 有効期間の長いキー上の OIDC
ヘルスチェック + ロールバック対応のデプロイ — /healthz、ワンステップで元に戻す
申し込む前に計画を立ててください – よく読んでください

差分。ライブインフラを手動で編集しないでください
レベル付きの構造化ログ — print("here 2") ではなく、JSON イベント
フェイルファスト、CI でハードキャッシュ — テスト前にビルド前に lint を実行
小規模、頻繁、元に戻せるリリース - 退屈なデプロイが目標

## Original Extract

A Claude Code skill that quietly injects 10 atomic DevOps best practices — secure Chainguard images, pinned dependencies, least-privilege CI — into every session.

🖥️ ">
dev@mac — ~/dev — zsh
➜ ~/dev whoami
someone who does not, in fact, hate devops. devops is just hard.
➜ ~/dev cat about.txt
ihatedevops
Star on GitHub
➜ ~/dev ./install.sh # pick your comfort level
claude code
codex
copy-paste
one-liner
claude plugin marketplace add zenconnor/ihatedevops && claude plugin install ihatedevops@ihatedevops
copy
the official plugin — gets updates. already inside a claude session? same thing: /plugin marketplace add zenconnor/ihatedevops, then /plugin install ihatedevops@ihatedevops.
git clone --depth 1 https://github.com/zenconnor/ihatedevops /tmp/ihd && mkdir -p ~/.agents/skills && cp -R /tmp/ihd/skills/devops-best-practices ~/.agents/skills/
copy
same skill, open Agent Skills standard — Codex reads ~/.agents/skills/. per-project: cp into the repo's .agents/skills/ instead.
git clone --depth 1 https://github.com/zenconnor/ihatedevops /tmp/ihd && mkdir -p ~/.claude/skills && cp -R /tmp/ihd/skills/devops-best-practices ~/.claude/skills/
copy
every step readable: clone, mkdir, copy. no script runs. for one project only, cp into its .claude/skills/ instead.
curl -fsSL https://raw.githubusercontent.com/zenconnor/ihatedevops/main/install.sh | sh
copy
fastest. installs one folder into ~/.claude/skills/ — nothing else, ever. read install.sh first if you like.
➜ ~/dev ls practices/
Chainguard base images by default — cgr.dev/chainguard/*, free, ~0 CVEs
Multi-stage builds, non-root user — build tools never ship
Pin everything — lockfiles, image digests, actions by commit SHA
Least-privilege CI tokens — explicit permissions: contents: read
No secrets in code, images, or logs — OIDC over long-lived keys
Health checks + rollback-ready deploys — /healthz, one-step revert
Plan before apply — read the diff; never hand-edit live infra
Structured logs with levels — JSON events, not print("here 2")
Fail fast, cache hard in CI — lint before build before test
Small, frequent, reversible releases — boring deploys are the goal
