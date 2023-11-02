# Lime

A fresh take on a clean inbox.

Lime is a command line tool that will classify emails in your inbox. The goal is
to find all the emails that you can delete. Emails you want to read are very
different from the emails you want to keep. Having a massive inbox hampers
searchability, making it hard to find what you need. This is especially true
with newsletters and marketing emails that clog up your inbox. Useful at the
moment but unnecessary to keep around.

While the ultimate goal of Lime is to classify emails into keep and remove, it
can be unnerving to be given a button that indiscriminately deletes large
amounts of emails. Therefore, Lime will apply a set of labels to all emails in
your inbox and delete matching criteria. There is a set of default labels, but
those can be changed via configuration or pulled in from your current inbox.

## Usage

```bash
go build
./lime
```
