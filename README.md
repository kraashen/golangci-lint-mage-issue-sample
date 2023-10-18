# golangci-mage-sample

## Reproducing

```bash
git clone git@github.com:kraashen/golangci-lint-mage-issue-sample.git
cd golangci-lint-mage-issue-sample
mage checkall
Running target: CheckAll
Running dependency: Lint

echo $?
0
 
# only Lint target is run, but not the following subsequent target
```

### Expected output

```bash
# expected output (with fix https://github.com/kraashen/golangci-lint/commit/34382d0b26cce750721ac05dffd941843c3d810c):

❯ mage checkall
Running target: CheckAll
Running dependency: Lint
Running dependency: VulnCheck
Scanning your code and 41 packages across 0 dependent modules for known vulnerabilities...

No vulnerabilities found.

Share feedback at https://go.dev/s/govulncheck-feedback.
```
