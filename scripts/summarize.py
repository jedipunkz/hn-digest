#!/usr/bin/env python3
"""
Select top HN articles by interest-weighted score, summarize with GitHub Models.
Outputs summaries/YYYY-MM-DD.json.
"""

import json
import os
import re
import sys
from datetime import datetime, timezone
from pathlib import Path

import yaml
from openai import OpenAI

TOP_N = 10

# Each category contributes at most once to the multiplier.
# Final score = hn_score * (1.0 + sum of matched bonuses)
INTERESTS = [
    {
        "name": "ai",
        "keywords": [
            "ai", "llm", "llms", "machine learning", "ml", "openai", "anthropic",
            "claude", "gpt", "gemini", "neural", "artificial intelligence", "deepmind",
            "embedding", "rag", "inference", "fine-tun", "foundation model",
        ],
        "bonus": 0.5,
    },
    {
        "name": "sre",
        "keywords": [
            "sre", "site reliability", "incident", "on-call", "on call", "observability",
            "monitoring", "alerting", "postmortem", "runbook", "pagerduty", "chaos engineering",
            "toil", "error budget",
        ],
        "bonus": 0.4,
    },
    {
        "name": "platform",
        "keywords": [
            "platform engineering", "platform", "kubernetes", "k8s", "terraform",
            "infrastructure", "devops", "dev ops", "ci/cd", "cloud run", "gcp",
            "google cloud", "gitops", "helm", "argo", "pulumi", "internal developer",
        ],
        "bonus": 0.3,
    },
]


def parse_frontmatter(text: str) -> dict:
    match = re.match(r"^---\n(.*?)\n---\n", text, re.DOTALL)
    if not match:
        return {}
    try:
        return yaml.safe_load(match.group(1)) or {}
    except yaml.YAMLError:
        return {}


def extract_translation(text: str) -> str:
    match = re.search(r"## Translation\n\n(.*?)(?=\n## |\Z)", text, re.DOTALL)
    return match.group(1).strip()[:2000] if match else ""


def interest_multiplier(title: str, translation: str) -> float:
    combined = (title + " " + translation).lower()
    multiplier = 1.0
    for interest in INTERESTS:
        for kw in interest["keywords"]:
            if kw in combined:
                multiplier += interest["bonus"]
                break  # count each category at most once
    return multiplier


def build_client() -> OpenAI:
    token = os.environ.get("GITHUB_TOKEN", "")
    if not token:
        sys.exit("GITHUB_TOKEN is not set")
    return OpenAI(
        base_url="https://models.inference.ai.azure.com",
        api_key=token,
    )


def summarize(client: OpenAI, title: str, translation: str) -> str:
    resp = client.chat.completions.create(
        model="gpt-4o-mini",
        messages=[
            {
                "role": "system",
                "content": (
                    "あなたは技術ニュースの要約者です。"
                    "Hacker Newsの記事を日本語で簡潔に要約してください。"
                ),
            },
            {
                "role": "user",
                "content": (
                    "以下のHacker News記事を2〜3文の日本語で要約してください。\n"
                    "技術的なポイントとHNコミュニティでの注目理由を含めてください。\n\n"
                    f"タイトル: {title}\n\n本文:\n{translation}"
                ),
            },
        ],
        max_tokens=300,
        temperature=0.3,
    )
    return resp.choices[0].message.content.strip()


def main() -> None:
    today = datetime.now(timezone.utc).strftime("%Y-%m-%d")
    contents_dir = Path("contents") / today

    if not contents_dir.exists():
        print(f"No contents for {today}, nothing to do.", file=sys.stderr)
        return

    articles = []
    for md_file in contents_dir.glob("*.md"):
        text = md_file.read_text(encoding="utf-8")
        fm = parse_frontmatter(text)
        hn_id = fm.get("hn_id")
        hn_score = fm.get("score", 0)
        if not hn_id or not isinstance(hn_score, int):
            continue
        title = fm.get("title") or ""
        translation = extract_translation(text)
        multiplier = interest_multiplier(title, translation)
        articles.append(
            {
                "hn_id": int(hn_id),
                "title": title,
                "hn_url": str(fm.get("hn_url") or ""),
                "source_url": str(fm.get("source") or ""),
                "score": hn_score,
                "final_score": int(hn_score * multiplier),
                "comments": int(fm.get("comments", 0)),
                "posted_at": str(fm.get("posted_at") or ""),
                "translation": translation,
            }
        )

    if not articles:
        print("No valid articles found.", file=sys.stderr)
        return

    articles.sort(key=lambda a: a["final_score"], reverse=True)
    top = articles[:TOP_N]

    client = build_client()
    results = []
    for rank, art in enumerate(top, 1):
        print(f"[{rank}/{TOP_N}] score={art['final_score']} ({art['score']}×{art['final_score']//max(art['score'],1):.1f}x) {art['title'][:60]}")
        summary = summarize(client, art["title"], art["translation"])
        results.append(
            {
                "rank": rank,
                "hn_id": art["hn_id"],
                "title": art["title"],
                "hn_url": art["hn_url"],
                "source_url": art["source_url"],
                "score": art["score"],
                "final_score": art["final_score"],
                "comments": art["comments"],
                "posted_at": art["posted_at"],
                "summary_ja": summary,
            }
        )

    out_dir = Path("summaries")
    out_dir.mkdir(exist_ok=True)
    out_path = out_dir / f"{today}.json"
    out_path.write_text(
        json.dumps(
            {
                "date": today,
                "generated_at": datetime.now(timezone.utc).isoformat(),
                "articles": results,
            },
            ensure_ascii=False,
            indent=2,
        ),
        encoding="utf-8",
    )
    print(f"Saved {len(results)} summaries to {out_path}")


if __name__ == "__main__":
    main()
