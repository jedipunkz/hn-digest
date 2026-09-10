---
source: "https://promptsign.ai/"
hn_url: "https://news.ycombinator.com/item?id=49647883"
title: "Show HN: PromptSign – Sigstore signing and verification for AI instruction files"
article_title: "PromptSign.ai — Code signing for AI skills and agents"
image: "https://promptsign.ai/og.png"
author: "sergey_v"
captured_at: "2026-09-10T18:09:17Z"
capture_tool: "hn-digest"
hn_id: 49647883
score: 1
comments: 0
posted_at: "2026-09-10T18:03:06Z"
tags:
  - hacker-news
---

# Show HN: PromptSign – Sigstore signing and verification for AI instruction files

- HN: [49647883](https://news.ycombinator.com/item?id=49647883)
- Source: [promptsign.ai](https://promptsign.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-09-10T18:03:06Z

## Translation

Title: Show HN: PromptSign – Sigstore signing and verification for AI instruction files
Article title: PromptSign.ai — Code signing for AI skills and agents
Description: Know who published the skills, agents and CLAUDE.md / AGENTS.md files your AI runs — and that not one byte has changed. Sign and verify in your browser.
HN text: I built PromptSign after finding some useful-looking AI skills on the Internet and realizing I couldn't really know where they came from. In particular, I could not answer: 1) Who created the skill? 2) Is what I'm invoking right now really what I installed in the first place? 3) Is this update from the same source as the original? 4) What if I only wanted to install by reputation i.e. to whitelist certain skill publishers and ignore everyone else? AI instructions are markdown files, freely modifiable after install. The only check I could find was once at install against a digest in the marketplace manifest that the publisher controls. In short, there was no permanent AI skill "identity". How it works:
PromptSign signs AI instruction files with Sigstore keyless signing after the author authenticates via a GitHub, Google, or Microsoft account. Sigstore issues a short-lived certificate binding that identity to an ephemeral Ed25519 key (generated locally and never written to disk). PromptSign signs a manifest of hashes for every file in the skill's directory with that key. The manifest is signed as a Dead Simple Signing Envelope (DSSE) and sent to Rekor for public timestamping, because Sigstore certificates expire in minutes. The bundle stored alongside the skill (.promptsign directory) holds that envelope, the signing certificate, and the Rekor receipt. That bundle verifies offline, naming the publisher at any time. For question 4 a policy file can pin a skill name pattern to a required identity and issuer. Question 3 is trust on first use (on by default): the first signer seen for a name is remembered, and a later signature from anyone else is a hard failure even when the policy is otherwise only warning. I added Claude Code and Codex hooks to call my verifier to report any skill integrity and identity at session start and tool use time. A skill with a failing signature is blocked before use by default, which is the point of a signing tool. For OpenClaw an install policy blocks a tampered skill before it reaches disk. Hooks can fail on unsigned skills too, rather than just report them. It's an enforce rule in the policy file, but the default is warn. There is a 2-minute silent demo video on the site with a real terminal session, not a mockup, at https://promptsign.ai/posts/what-signing-proves . However, there is another angle: the website itself can sign and verify skills when you don't want to install anything. Signing needs network access for Sigstore login, certificate request to Fulcio, and Rekor log POST. Note that Fulcio and Rekor don't send CORS headers, so the browser flow relays through a narrow forwarder on promptsign.ai (implemented by "promptsign proxy" CLI command). It passes GET/POST/OPTIONS to allowlisted Sigstore hosts only. File contents still never leave the tab; what transits is the manifest: relative paths and hashes. Verification is fully offline thanks to the pinned Sigstore root in the site's JavaScript. And now I want to say four things: 1) Signed is not a safety verdict. A malicious skill that is signed still verifies (but we'll know who did it). 2) Unsigned is not malicious, because almost the entire ecosystem is unsigned today. 3) Content scanning is still needed to tell you what the skill does. 4) The Rekor log is public. The signer's email ends up in the certificate that is permanently stored by Rekor, so that email is permanently public. The skill's filenames are not public, because Rekor stores only the SHA-256 hash of the manifest, but not manifest itself. PromptSign is work in progress with some existing rough edges. For example, CLI binaries are not Authenticode-signed or notarized (though GitHub build-provenance attestations are there), so you'll have to bypass Windows or macOS gatekeepers to run them. Spec and code are Apache 2.0. I'd love to hear your critiques on the approach!

Article text:
here, so this
page makes no third-party requests at all. That is deliberate: /verify
claims to run entirely offline and invites readers to confirm it in
devtools, and any third-party host in that request list would undercut
the claim. Grepping the build output for a font CDN hostname should
return nothing, so please keep one out of this comment. -->
PromptSign Sign Verify Docs Directory Integrate GitHub ↗ PKI for AI instruction files Sign the prompts
your agents obey.
Know who published the skills, agents, and CLAUDE.md / AGENTS.md files your AI runs, and verify that neither they nor the scripts they ship have changed since they were signed.
If a publisher becomes untrustworthy, revoke their trust.
A signature proves who published this and that it hasn't changed . It does not prove that it's safe.
signed · sealed · verifiable offline What a signature gives you Four guarantees. One limit.
You know exactly who published a skill: a named identity, never a bare checkmark.
You know not one byte, including the scripts a skill ships, has changed since signing.
Keyless signatures are recorded in a public transparency log that cannot be quietly rewritten.
A compromised key or a bad actor can be revoked, so consumers stop trusting them.
A valid signature on a prompt-injection payload is still a prompt-injection payload. Signing is provenance, not safety.
Start in a terminal, or right here.
The command-line tool signs and verifies, and plugs into Claude Code & Codex hooks that automatically block any instructions that are unsigned or have been tampered with.
Prefer not to install anything? The in-browser tools do the same thing, and nothing in the files content you drop in leaves your computer.
Proof of origin you can check yourself, even offline.
PromptSign takes a tamper-evident fingerprint (a SHA-256 hash) of every file in the skill, including any scripts it ships. Change a single byte later and the fingerprint no longer matches.
02 Sign it under a verified identity
The publisher signs those fingerprints by authenticating with an account you already recognize, such as GitHub, Google, Microsoft, or a company account, so the identity on the signature is verified, not self-declared.
The signature travels with the skill in a small .promptsign/bundle.json file. To check it, PromptSign re-fingerprints every file and confirms the signature, no server or internet connection required.
{
"schema": "promptsign/bundle/v1",
"envelope": {
"payloadType": "application/vnd.promptsign.manifest+json",
"payload": "eyJzY2hlbWEiOiJwcm9tcHRzaWduL21hbmlmZXN0...",
"signatures": [
{ "keyid": "a1b2c3d4e5f6…", "sig": "MEUCIQD…" }
]
},
"signer": {
"identity": "alice@trailofbits.com",
"scheme": "keyless",
"issuer": "https://accounts.google.com",
"certChain": ["MIICzDCCAlKg…"]
},
"transparency": { "logIndex": 51236401 }
} Notice the identity — confirmed by a real login and named by every verifier. Never a bare green check.

## Original Extract

Know who published the skills, agents and CLAUDE.md / AGENTS.md files your AI runs — and that not one byte has changed. Sign and verify in your browser.

I built PromptSign after finding some useful-looking AI skills on the Internet and realizing I couldn't really know where they came from. In particular, I could not answer: 1) Who created the skill? 2) Is what I'm invoking right now really what I installed in the first place? 3) Is this update from the same source as the original? 4) What if I only wanted to install by reputation i.e. to whitelist certain skill publishers and ignore everyone else? AI instructions are markdown files, freely modifiable after install. The only check I could find was once at install against a digest in the marketplace manifest that the publisher controls. In short, there was no permanent AI skill "identity". How it works:
PromptSign signs AI instruction files with Sigstore keyless signing after the author authenticates via a GitHub, Google, or Microsoft account. Sigstore issues a short-lived certificate binding that identity to an ephemeral Ed25519 key (generated locally and never written to disk). PromptSign signs a manifest of hashes for every file in the skill's directory with that key. The manifest is signed as a Dead Simple Signing Envelope (DSSE) and sent to Rekor for public timestamping, because Sigstore certificates expire in minutes. The bundle stored alongside the skill (.promptsign directory) holds that envelope, the signing certificate, and the Rekor receipt. That bundle verifies offline, naming the publisher at any time. For question 4 a policy file can pin a skill name pattern to a required identity and issuer. Question 3 is trust on first use (on by default): the first signer seen for a name is remembered, and a later signature from anyone else is a hard failure even when the policy is otherwise only warning. I added Claude Code and Codex hooks to call my verifier to report any skill integrity and identity at session start and tool use time. A skill with a failing signature is blocked before use by default, which is the point of a signing tool. For OpenClaw an install policy blocks a tampered skill before it reaches disk. Hooks can fail on unsigned skills too, rather than just report them. It's an enforce rule in the policy file, but the default is warn. There is a 2-minute silent demo video on the site with a real terminal session, not a mockup, at https://promptsign.ai/posts/what-signing-proves . However, there is another angle: the website itself can sign and verify skills when you don't want to install anything. Signing needs network access for Sigstore login, certificate request to Fulcio, and Rekor log POST. Note that Fulcio and Rekor don't send CORS headers, so the browser flow relays through a narrow forwarder on promptsign.ai (implemented by "promptsign proxy" CLI command). It passes GET/POST/OPTIONS to allowlisted Sigstore hosts only. File contents still never leave the tab; what transits is the manifest: relative paths and hashes. Verification is fully offline thanks to the pinned Sigstore root in the site's JavaScript. And now I want to say four things: 1) Signed is not a safety verdict. A malicious skill that is signed still verifies (but we'll know who did it). 2) Unsigned is not malicious, because almost the entire ecosystem is unsigned today. 3) Content scanning is still needed to tell you what the skill does. 4) The Rekor log is public. The signer's email ends up in the certificate that is permanently stored by Rekor, so that email is permanently public. The skill's filenames are not public, because Rekor stores only the SHA-256 hash of the manifest, but not manifest itself. PromptSign is work in progress with some existing rough edges. For example, CLI binaries are not Authenticode-signed or notarized (though GitHub build-provenance attestations are there), so you'll have to bypass Windows or macOS gatekeepers to run them. Spec and code are Apache 2.0. I'd love to hear your critiques on the approach!

here, so this
page makes no third-party requests at all. That is deliberate: /verify
claims to run entirely offline and invites readers to confirm it in
devtools, and any third-party host in that request list would undercut
the claim. Grepping the build output for a font CDN hostname should
return nothing, so please keep one out of this comment. -->
PromptSign Sign Verify Docs Directory Integrate GitHub ↗ PKI for AI instruction files Sign the prompts
your agents obey.
Know who published the skills, agents, and CLAUDE.md / AGENTS.md files your AI runs, and verify that neither they nor the scripts they ship have changed since they were signed.
If a publisher becomes untrustworthy, revoke their trust.
A signature proves who published this and that it hasn't changed . It does not prove that it's safe.
signed · sealed · verifiable offline What a signature gives you Four guarantees. One limit.
You know exactly who published a skill: a named identity, never a bare checkmark.
You know not one byte, including the scripts a skill ships, has changed since signing.
Keyless signatures are recorded in a public transparency log that cannot be quietly rewritten.
A compromised key or a bad actor can be revoked, so consumers stop trusting them.
A valid signature on a prompt-injection payload is still a prompt-injection payload. Signing is provenance, not safety.
Start in a terminal, or right here.
The command-line tool signs and verifies, and plugs into Claude Code & Codex hooks that automatically block any instructions that are unsigned or have been tampered with.
Prefer not to install anything? The in-browser tools do the same thing, and nothing in the files content you drop in leaves your computer.
Proof of origin you can check yourself, even offline.
PromptSign takes a tamper-evident fingerprint (a SHA-256 hash) of every file in the skill, including any scripts it ships. Change a single byte later and the fingerprint no longer matches.
02 Sign it under a verified identity
The publisher signs those fingerprints by authenticating with an account you already recognize, such as GitHub, Google, Microsoft, or a company account, so the identity on the signature is verified, not self-declared.
The signature travels with the skill in a small .promptsign/bundle.json file. To check it, PromptSign re-fingerprints every file and confirms the signature, no server or internet connection required.
{
"schema": "promptsign/bundle/v1",
"envelope": {
"payloadType": "application/vnd.promptsign.manifest+json",
"payload": "eyJzY2hlbWEiOiJwcm9tcHRzaWduL21hbmlmZXN0...",
"signatures": [
{ "keyid": "a1b2c3d4e5f6…", "sig": "MEUCIQD…" }
]
},
"signer": {
"identity": "alice@trailofbits.com",
"scheme": "keyless",
"issuer": "https://accounts.google.com",
"certChain": ["MIICzDCCAlKg…"]
},
"transparency": { "logIndex": 51236401 }
} Notice the identity — confirmed by a real login and named by every verifier. Never a bare green check.
