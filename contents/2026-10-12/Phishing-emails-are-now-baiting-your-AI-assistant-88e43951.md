---
source: "https://www.malwarebytes.com/blog/ai/2026/10/phishing-emails-are-baiting-your-ai-assistant"
hn_url: "https://news.ycombinator.com/item?id=50061773"
title: "Phishing emails are now baiting your AI assistant"
article_title: "Phishing emails are baiting your AI assistant | Malwarebytes"
image: "https://www.malwarebytes.com/wp-content/uploads/sites/2/2026/10/ai-assistant-email.jpg"
author: "Apollo_R"
captured_at: "2026-10-12T20:08:23Z"
capture_tool: "hn-digest"
hn_id: 50061773
score: 2
comments: 0
posted_at: "2026-10-12T19:35:02Z"
tags:
  - hacker-news
---

# Phishing emails are now baiting your AI assistant

- HN: [50061773](https://news.ycombinator.com/item?id=50061773)
- Source: [www.malwarebytes.com](https://www.malwarebytes.com/blog/ai/2026/10/phishing-emails-are-baiting-your-ai-assistant)
- Score: 2
- Comments: 0
- Posted: 2026-10-12T19:35:02Z

## Translation

Title: Phishing emails are now baiting your AI assistant
Article title: Phishing emails are baiting your AI assistant | Malwarebytes
Description: If you use an AI assistant to summarize your inbox, hidden instructions could turn a phishing email you wisely ignored into your next to-do item.

Article text:
Skip to content
Search
Search Malwarebytes.com
Search for:
Sign In
Sign in
Activate subscription >
Don’t have an account?
Sign up >
Mobile security for iOS and Android
Looking for small business protection? Visit Teams
About Company About Malwarebytes
Modern Love in the Digital Age
How AI is reshaping trust, identity, and scams
Phishing emails are baiting your AI assistant
You ignored the suspicious email. You didn’t open the attachment, click the link, or reply to the sender.
Then you asked your AI assistant to summarize your inbox.
Could the message you wisely avoided find its way back onto your to-do list—only this time presented as an urgent recommendation from a tool you trust?
That’s the risk behind phishing emails containing hidden instructions aimed at AI assistants. Instead of relying only on persuading you, attackers are also trying to influence the software that reads your messages. Recent research from Barracuda describes emails designed to target both audiences at once.
For the human recipient, there was a password-protected attachment, with the password supplied in the email. For an AI assistant processing the message, there were hidden instructions intended to influence how it summarized or handled the content.
By including a prompt injection , attackers give themselves another way to reach you. The malicious instructions could try to make the assistant describe the email as legitimate or urgent, encouraging you to revisit it.
In another example, a hidden block instructed a summarizing AI to add a fake priority action: Change a vendor’s payment details. That could result in an employee sending money to an attacker’s account.
The attacker isn’t simply asking the AI to repeat a fraudulent request. They’re trying to turn that request into the assistant’s recommendation.
Prompt injection means planting instructions in material an AI system is supposed to read, hoping it will follow those instructions instead of treating them as part of the material. When the instructions arrive through an outside source, such as an email or webpage, it’s often called indirect prompt injection.
The malicious instructions don’t have to appear as ordinary, readable paragraphs. Barracuda identified concealment techniques including comments in the email’s underlying HTML code, text made invisible through styling, encoded content, and invisible “zero-width” characters. These can obscure instructions from a person viewing the message while leaving content available for an AI system to process. Whether that content actually reaches the model depends on how the particular tool extracts and prepares the email.
That means this problem cannot be solved simply by asking people to look more carefully at their emails. Part of the defense must sit inside the tools handling those messages.
Apollo Raines , developer of the AI code review service ShipItClean , explains:
“The real flaw is that the model can’t tell the difference between text it’s supposed to read and instructions it’s supposed to follow. Stripping hidden content before it reaches the model is worth doing, and every vendor should do it. But it’s an arms race. Each new encoding trick is a new hole.”
“People have learned to distrust the email. They haven’t learned to distrust their AI assistant. So, the defense has to live inside the tools. I’d go one step further: some of it has to live inside the model itself. Wrappers, filters, sandboxing and human sign-off on payments are all necessary.”
An assistant that summarizes messages could mislead you about which email deserves attention or what action it requires.
An assistant with additional permissions could create a more serious problem. Depending on its capabilities and safeguards, malicious instructions might try to make it carry out unauthorized actions or disclose sensitive information .
You don’t need to abandon AI tools that help with email. But you do need to consider what you trust them to do and what permissions you give them.
The familiar phishing advice still applies. When an assistant reads your inbox on your behalf, check its recommendations before acting:
Check the original message before acting on a recommendation involving money or sensitive information. If a summary tells you to pay an invoice, change account details, or share sensitive information, don’t treat the summary itself as verification.
Always verify payment changes through a separate, established channel. Contact the person or organization using details you already trust, rather than relying on information supplied in the suspicious message.
Keep your AI assistant’s permissions limited to what you need and, at work, what your organization allows. An assistant used to summarize emails should not need broad authority to carry out unrelated actions. Limiting access can reduce the consequences if it follows malicious instructions.
Require human approval for high-impact actions like payments and changes to vendor payment details. Approval should include checking the underlying request, not simply accepting the assistant’s recommendation.
From reporting threats to removing them.
Cybersecurity risks should never spread beyond a headline. Keep threats off your devices by downloading Malwarebytes today .
Malware Intelligence Researcher
Was a Microsoft MVP in consumer security for 12 years running. Can speak four languages. Smells of rich mahogany and leather-bound books.
A list of topics we covered in the week of October 5 to October 11 of 2026
Stolen ASOS data includes shopping searches as well as personal details, so customers should watch out for convincing phishing messages.
Cybercriminals compromised three country-code namespaces to get certificates that could help them impersonate trusted sites.
Shadow AI explained: The work shortcut that could leak your company’s secrets
Losing gamblers pushed to bet more by DraftKings’ AI, report says
Meta’s Muse sent a Facebook Marketplace buyer to a seller’s home
Sign up for our newsletter to get the latest cybersecurity news to your inbox
By submitting this form, you consent to Malwarebytes contacting you regarding products and services and using your personal data as described in our Terms of Service and Privacy Policy .
Malwarebytes - all-in-one cybersecurity protection always by your side.
Identity Recovery & Resolution
Business Endpoint Security Solutions
Managed Service Provider (MSP) Program
Strategic Business Partnerships
Penrose One
Penrose Dock, 1st Floor
Cork T23 KW81
Ireland
2445 Augustine Drive
Suite 550
Santa Clara, CA
USA, 95054
Malwarebytes vs. Microsoft Defender
Cybersecurity info you can’t live without
Want to stay informed on the latest news in cybersecurity? Sign up for our newsletter and learn how to protect your computer from threats.
By submitting this form, you consent to Malwarebytes contacting you regarding products and services and using your personal data as described in our Terms of Service and Privacy Policy .

## Original Extract

If you use an AI assistant to summarize your inbox, hidden instructions could turn a phishing email you wisely ignored into your next to-do item.

Skip to content
Search
Search Malwarebytes.com
Search for:
Sign In
Sign in
Activate subscription >
Don’t have an account?
Sign up >
Mobile security for iOS and Android
Looking for small business protection? Visit Teams
About Company About Malwarebytes
Modern Love in the Digital Age
How AI is reshaping trust, identity, and scams
Phishing emails are baiting your AI assistant
You ignored the suspicious email. You didn’t open the attachment, click the link, or reply to the sender.
Then you asked your AI assistant to summarize your inbox.
Could the message you wisely avoided find its way back onto your to-do list—only this time presented as an urgent recommendation from a tool you trust?
That’s the risk behind phishing emails containing hidden instructions aimed at AI assistants. Instead of relying only on persuading you, attackers are also trying to influence the software that reads your messages. Recent research from Barracuda describes emails designed to target both audiences at once.
For the human recipient, there was a password-protected attachment, with the password supplied in the email. For an AI assistant processing the message, there were hidden instructions intended to influence how it summarized or handled the content.
By including a prompt injection , attackers give themselves another way to reach you. The malicious instructions could try to make the assistant describe the email as legitimate or urgent, encouraging you to revisit it.
In another example, a hidden block instructed a summarizing AI to add a fake priority action: Change a vendor’s payment details. That could result in an employee sending money to an attacker’s account.
The attacker isn’t simply asking the AI to repeat a fraudulent request. They’re trying to turn that request into the assistant’s recommendation.
Prompt injection means planting instructions in material an AI system is supposed to read, hoping it will follow those instructions instead of treating them as part of the material. When the instructions arrive through an outside source, such as an email or webpage, it’s often called indirect prompt injection.
The malicious instructions don’t have to appear as ordinary, readable paragraphs. Barracuda identified concealment techniques including comments in the email’s underlying HTML code, text made invisible through styling, encoded content, and invisible “zero-width” characters. These can obscure instructions from a person viewing the message while leaving content available for an AI system to process. Whether that content actually reaches the model depends on how the particular tool extracts and prepares the email.
That means this problem cannot be solved simply by asking people to look more carefully at their emails. Part of the defense must sit inside the tools handling those messages.
Apollo Raines , developer of the AI code review service ShipItClean , explains:
“The real flaw is that the model can’t tell the difference between text it’s supposed to read and instructions it’s supposed to follow. Stripping hidden content before it reaches the model is worth doing, and every vendor should do it. But it’s an arms race. Each new encoding trick is a new hole.”
“People have learned to distrust the email. They haven’t learned to distrust their AI assistant. So, the defense has to live inside the tools. I’d go one step further: some of it has to live inside the model itself. Wrappers, filters, sandboxing and human sign-off on payments are all necessary.”
An assistant that summarizes messages could mislead you about which email deserves attention or what action it requires.
An assistant with additional permissions could create a more serious problem. Depending on its capabilities and safeguards, malicious instructions might try to make it carry out unauthorized actions or disclose sensitive information .
You don’t need to abandon AI tools that help with email. But you do need to consider what you trust them to do and what permissions you give them.
The familiar phishing advice still applies. When an assistant reads your inbox on your behalf, check its recommendations before acting:
Check the original message before acting on a recommendation involving money or sensitive information. If a summary tells you to pay an invoice, change account details, or share sensitive information, don’t treat the summary itself as verification.
Always verify payment changes through a separate, established channel. Contact the person or organization using details you already trust, rather than relying on information supplied in the suspicious message.
Keep your AI assistant’s permissions limited to what you need and, at work, what your organization allows. An assistant used to summarize emails should not need broad authority to carry out unrelated actions. Limiting access can reduce the consequences if it follows malicious instructions.
Require human approval for high-impact actions like payments and changes to vendor payment details. Approval should include checking the underlying request, not simply accepting the assistant’s recommendation.
From reporting threats to removing them.
Cybersecurity risks should never spread beyond a headline. Keep threats off your devices by downloading Malwarebytes today .
Malware Intelligence Researcher
Was a Microsoft MVP in consumer security for 12 years running. Can speak four languages. Smells of rich mahogany and leather-bound books.
A list of topics we covered in the week of October 5 to October 11 of 2026
Stolen ASOS data includes shopping searches as well as personal details, so customers should watch out for convincing phishing messages.
Cybercriminals compromised three country-code namespaces to get certificates that could help them impersonate trusted sites.
Shadow AI explained: The work shortcut that could leak your company’s secrets
Losing gamblers pushed to bet more by DraftKings’ AI, report says
Meta’s Muse sent a Facebook Marketplace buyer to a seller’s home
Sign up for our newsletter to get the latest cybersecurity news to your inbox
By submitting this form, you consent to Malwarebytes contacting you regarding products and services and using your personal data as described in our Terms of Service and Privacy Policy .
Malwarebytes - all-in-one cybersecurity protection always by your side.
Identity Recovery & Resolution
Business Endpoint Security Solutions
Managed Service Provider (MSP) Program
Strategic Business Partnerships
Penrose One
Penrose Dock, 1st Floor
Cork T23 KW81
Ireland
2445 Augustine Drive
Suite 550
Santa Clara, CA
USA, 95054
Malwarebytes vs. Microsoft Defender
Cybersecurity info you can’t live without
Want to stay informed on the latest news in cybersecurity? Sign up for our newsletter and learn how to protect your computer from threats.
By submitting this form, you consent to Malwarebytes contacting you regarding products and services and using your personal data as described in our Terms of Service and Privacy Policy .
