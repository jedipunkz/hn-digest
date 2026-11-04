---
source: "https://github.com/Dr-PampersMuffin/continuum-verifier"
hn_url: "https://news.ycombinator.com/item?id=50368207"
title: "Show HN: Continuum Verifier – offline checks for AI data-provenance records"
article_title: "GitHub - Dr-PampersMuffin/continuum-verifier: Offline, self-contained verifier for portable provenance records (Continuum Evidence Capsule) · GitHub"
image: "https://opengraph.githubassets.com/74712ed37a18e31a445dcc901ab898347c159461237277a7659ffaa7c55ac50c/Dr-PampersMuffin/continuum-verifier"
author: "PampersMuffin"
captured_at: "2026-11-04T16:54:09Z"
capture_tool: "hn-digest"
hn_id: 50368207
score: 1
comments: 0
posted_at: "2026-11-04T16:44:58Z"
tags:
  - hacker-news
---

# Show HN: Continuum Verifier – offline checks for AI data-provenance records

- HN: [50368207](https://news.ycombinator.com/item?id=50368207)
- Source: [github.com](https://github.com/Dr-PampersMuffin/continuum-verifier)
- Score: 1
- Comments: 0
- Posted: 2026-11-04T16:44:58Z

## Translation

Title: Show HN: Continuum Verifier – offline checks for AI data-provenance records
Article title: GitHub - Dr-PampersMuffin/continuum-verifier: Offline, self-contained verifier for portable provenance records (Continuum Evidence Capsule) · GitHub
Description: Offline, self-contained verifier for portable provenance records (Continuum Evidence Capsule) - Dr-PampersMuffin/continuum-verifier

Article text:
GitHub - Dr-PampersMuffin/continuum-verifier: Offline, self-contained verifier for portable provenance records (Continuum Evidence Capsule) · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Dr-PampersMuffin
continuum-verifier Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Offline, self-contained verifier for portable provenance records (Continuum Evidence Capsule)
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
samples samples .gitignore .gitignore CAPSULE-SPEC-0.1.md CAPSULE-SPEC-0.1.md LICENSE LICENSE README.md README.md attack_capsule.py attack_capsule.py capsule_pack.py capsule_pack.py capsule_restamp.py capsule_restamp.py capsule_verify.py capsule_verify.py View all files Repository files navigation
An open, offline verifier for portable provenance records.
A capsule is one self-contained file that carries a content fingerprint, a
chain of signed records (who attested to what, and when), Merkle inclusion
proofs, the issuer's key, public-witness anchors, and a re-stamp trail. A
stranger with no internet, no access to the issuer, and only this repo can hand
the capsule to the verifier and get an answer. That is the point: the proof
travels, and it keeps working after the system that made it is gone.
The verifier imports nothing from any server or private engine. It reads
only the capsule. The test suite proves this by running the verifier in a bare
directory that holds nothing but the verifier and a capsule.
# verify a capsule (offline, no keys, no network)
python3 capsule_verify.py samples/sample_capsule_durable.json
# the self-asserted sample (no durable witness) returns SELF_CONSISTENT
python3 capsule_verify.py samples/sample_capsule_selfasserted.json
# run the adversarial suite: clean capsules + every tamper + the battery swap
python3 attack_capsule.py # 14/14 expected
# make your own capsule, then re-stamp it onto fresh crypto
python3 capsule_pack.py mycap.json --witness nist-beacon
python3 capsule_restamp.py mycap.json mycap.v2.json
Only dependency beyond the Python standard library is an Ed25519 verify
( pip install cryptography ). Everything else is stdlib.
VERIFIED — every check passes AND a durable public witness anchors the root.
SELF_CONSISTENT — internally self-verifying, but the witness is self-asserted; durability over time is unproven.
RESTAMP_NEEDED — self-verifying, but the newest crypto is past its trust horizon. The verifier is telling you the proof needs re-stamping before you trust it.
WITHHELD — a hash, signature, or inclusion check failed. No confident answer, by design.
FAILED — not parseable as a capsule.
What it proves, and what it does not
Proves: a verified record was attested by a named key, is unmodified, folds into
a single anchored root (content included, not as decoration), and carries no
crypto past its trust horizon. The format is designed so that proving when
something was attested is a first-class check, not an afterthought.
Does not prove: the truth of the underlying facts. A capsule attests
custody and integrity. Whether a claim inside it is true is a separate question.
Honest status (read before relying on it)
This is a reference implementation at v0.1 , published to be read and broken.
The public witnesses and re-stamp signatures in the samples are simulated
with local keys . A sample witness named nist-beacon is a stand-in.
VERIFIED here means "would verify against a durable witness," not "is welded
to one." Wiring a real durable witness (OpenTimestamps, a transparency log) is
the next step.
The re-stamp trail rotates the hash family (SHA-256 to SHA3-512) to show the
mechanism. Rotating the signature to a post-quantum scheme is not done yet.
Forward only: a capsule attests what was true when it was sealed, carried
forward. It cannot certify history that was never witnessed.
See CAPSULE-SPEC-0.1.md for the full, language-agnostic verification algorithm
and the complete list of boundaries.
The suite ships 10+ tampers and a stress test, all expected to return WITHHELD.
If you find an input that verifies and should not, that is the most useful issue
you can open. PRs and adversarial test cases welcome.
MIT. See LICENSE . (Can be relicensed to Apache-2.0 later if a standards body
prefers its patent grant.)
Offline, self-contained verifier for portable provenance records (Continuum Evidence Capsule)
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Offline, self-contained verifier for portable provenance records (Continuum Evidence Capsule) - Dr-PampersMuffin/continuum-verifier

GitHub - Dr-PampersMuffin/continuum-verifier: Offline, self-contained verifier for portable provenance records (Continuum Evidence Capsule) · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Dr-PampersMuffin
continuum-verifier Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Offline, self-contained verifier for portable provenance records (Continuum Evidence Capsule)
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
samples samples .gitignore .gitignore CAPSULE-SPEC-0.1.md CAPSULE-SPEC-0.1.md LICENSE LICENSE README.md README.md attack_capsule.py attack_capsule.py capsule_pack.py capsule_pack.py capsule_restamp.py capsule_restamp.py capsule_verify.py capsule_verify.py View all files Repository files navigation
An open, offline verifier for portable provenance records.
A capsule is one self-contained file that carries a content fingerprint, a
chain of signed records (who attested to what, and when), Merkle inclusion
proofs, the issuer's key, public-witness anchors, and a re-stamp trail. A
stranger with no internet, no access to the issuer, and only this repo can hand
the capsule to the verifier and get an answer. That is the point: the proof
travels, and it keeps working after the system that made it is gone.
The verifier imports nothing from any server or private engine. It reads
only the capsule. The test suite proves this by running the verifier in a bare
directory that holds nothing but the verifier and a capsule.
# verify a capsule (offline, no keys, no network)
python3 capsule_verify.py samples/sample_capsule_durable.json
# the self-asserted sample (no durable witness) returns SELF_CONSISTENT
python3 capsule_verify.py samples/sample_capsule_selfasserted.json
# run the adversarial suite: clean capsules + every tamper + the battery swap
python3 attack_capsule.py # 14/14 expected
# make your own capsule, then re-stamp it onto fresh crypto
python3 capsule_pack.py mycap.json --witness nist-beacon
python3 capsule_restamp.py mycap.json mycap.v2.json
Only dependency beyond the Python standard library is an Ed25519 verify
( pip install cryptography ). Everything else is stdlib.
VERIFIED — every check passes AND a durable public witness anchors the root.
SELF_CONSISTENT — internally self-verifying, but the witness is self-asserted; durability over time is unproven.
RESTAMP_NEEDED — self-verifying, but the newest crypto is past its trust horizon. The verifier is telling you the proof needs re-stamping before you trust it.
WITHHELD — a hash, signature, or inclusion check failed. No confident answer, by design.
FAILED — not parseable as a capsule.
What it proves, and what it does not
Proves: a verified record was attested by a named key, is unmodified, folds into
a single anchored root (content included, not as decoration), and carries no
crypto past its trust horizon. The format is designed so that proving when
something was attested is a first-class check, not an afterthought.
Does not prove: the truth of the underlying facts. A capsule attests
custody and integrity. Whether a claim inside it is true is a separate question.
Honest status (read before relying on it)
This is a reference implementation at v0.1 , published to be read and broken.
The public witnesses and re-stamp signatures in the samples are simulated
with local keys . A sample witness named nist-beacon is a stand-in.
VERIFIED here means "would verify against a durable witness," not "is welded
to one." Wiring a real durable witness (OpenTimestamps, a transparency log) is
the next step.
The re-stamp trail rotates the hash family (SHA-256 to SHA3-512) to show the
mechanism. Rotating the signature to a post-quantum scheme is not done yet.
Forward only: a capsule attests what was true when it was sealed, carried
forward. It cannot certify history that was never witnessed.
See CAPSULE-SPEC-0.1.md for the full, language-agnostic verification algorithm
and the complete list of boundaries.
The suite ships 10+ tampers and a stress test, all expected to return WITHHELD.
If you find an input that verifies and should not, that is the most useful issue
you can open. PRs and adversarial test cases welcome.
MIT. See LICENSE . (Can be relicensed to Apache-2.0 later if a standards body
prefers its patent grant.)
Offline, self-contained verifier for portable provenance records (Continuum Evidence Capsule)
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
