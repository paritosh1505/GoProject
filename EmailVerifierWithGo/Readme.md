This is basic email verification where we are verfying emails with following checks

1. intally we will enter the domain name of the email which we want to verify

2. second check should be if MX record is present for the given domain. MX record is basically a type of domain name service which tells about for which email server it is registered

3. Third check should be checking the text record of email. This check also be done to verify the email it contian basic information such as SPF(sender policy framework) which gived the idea of which ip address or dmain or mail server is allowed to send the mail on behalf of domain. it is generally sued for email authentication. text record should be look like this. This help to prtoect from email spoofing also

v=spf1 ip4:50.242.118.232/29 ip4:198.21.2.190 ip4:149.72.166.252 include:\_spf.google.com include:mail.zendesk.com ~all

4. Fourth check is DMarch check it is protocol built on top of SPF. It help user to protect spam,phising and other email fraud. Output of DMARCH look le this

v=DMARC1; p=quarantine; rua=mailto:dmarc-reports@example.com; ruf=mailto:dmarc-failures@example.com; pct=100;

1. v=DMARC1
   Description: This specifies the version of the DMARC protocol being used.
2. p=quarantine
   quarantine means that emails failing DMARC checks should be treated as suspicious and should be placed in the recipient’s spam or junk folder, rather than being outright rejected.
3. pct=100
   pct=100 means the policy applies to 100% of the emails.
   If you set pct=50, for example, the policy would only apply to 50% of the emails, which can be useful for testing purposes before fully enforcing the policy.
4. rua=mailto:postmaster@hello.com
   Description: This tag specifies the email address where aggregate reports should be sent.
