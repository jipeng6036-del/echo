# Echo scoped IPv6 XFF verification

This evidence binds upstream base `4deac5da0714abbc2037a90067ea82c033a3f266`, [issue #3179](https://github.com/labstack/echo/issues/3179), [source PR #3180](https://github.com/labstack/echo/pull/3180), and final commit `42d7b17d99c0678ac22fe44076a4c94769964fc6` / tree `667325f667030f93e51f9aa3848525929e7937ab`.

Bharat Kaurav (TheNetherWatcher) authored the retained source commits:

- `9e2add6c8837fc018cf9974bf4f34f86b50ebb3b`
- `f65d3af076ff144ebc4fa34314f8dd55791e14b3`
- `db18e4b57993a2f809a9400346bc9f1cb1c2cab8`

Their replay preserves authorship and `cherry-pick -x` provenance. This version adds eight public HTTP regression scenarios and avoids formatting a canonical IPv6 string after strict validation: both callers subsequently use `net.ParseIP`, so returning the validated original prefix avoids one string allocation. The source PR's four extractor cases are retained and are not counted as new tests.

## Results and limits

| Evidence | Observed result |
| --- | --- |
| Eight public HTTP scenarios | Master: 3 expected failures / 5 controls pass. Final commit as an actual no-replace module ZIP: all 8 pass on Windows and Linux; Linux uses race detection. |
| IP/Scheme tests after allocation refinement | Source-PR control and refined implementation each pass 143 test/subtest lines. Ten additional public API semantic cases agree for noncanonical IPv6, IPv4-mapped forms and malformed/empty-zone input. |
| Quality | vet, golint, staticcheck, govulncheck, formatting and diff checks pass. |
| Native Windows full suite | Two IPv6-loopback socket-permission failures occur on both master and the source-PR-plus-HTTP-tests tree. These native full runs precede the allocation refinement. They are not reported as passing. Final hosted Windows/Linux/macOS race suites below complete the full-suite gate. |
| Final hosted CI | All 9 original matrix jobs succeed: Ubuntu/macOS/Windows x Go 1.25.14/1.26.10/1.27.3. Each log checks out the exact final SHA and shows race/coverage results for all three packages. |
| Final hosted Checks | golint, staticcheck and govulncheck actually execute successfully against `./...`; no zero-package warning. govulncheck reports no vulnerabilities. |
| Hosted benchmark | SKIPPED by the unchanged pull_request-only condition on a push run. Not counted as success. |

[CI run](https://github.com/jipeng6036-del/echo/actions/runs/40993443277) and [Checks run](https://github.com/jipeng6036-del/echo/actions/runs/40993443248) use the original workflow bytes. These are contributor-fork results, not upstream approval, merge or release.

The HTTP consumer opens a real TCP connection, serves Echo routes, calls `Context.RealIP` / `Context.Scheme`, and reads the response. `RemoteAddr` is explicitly substituted server-side as a proxy fixture; no real OS link-local IPv6 route or production proxy deployment is claimed.

## Performance evidence

Against the unmodified source-PR implementation, eight 100ms samples with the same public API benchmark on Windows / Go 1.27.3 measured scoped XFF at 289.8 -> 224.2 ns (4 -> 3 allocations), and scoped Scheme at 187.9 -> 122.3 ns (1 -> 0 allocations). IPv4 XFF measured 141.1 -> 145.8 ns (+3.40%) with three allocations unchanged. This does not claim a universal speedup or zero allocations for all extractors.

The original whole-package benchmark showed substantial variation on unrelated router/JSON paths (overall timing geomean +13.21%). Its logs are retained as noisy advisory evidence and do not establish either a global speedup or a code-attributed regression. The focused comparison is in `optimization/benchmark-comparison.log`.

## Actual module ZIP consumption

Module path is `github.com/labstack/echo/v5`. Candidate pseudo-version is `v5.4.3-0.20261030023608-42d7b17d99c0`; baseline is `v5.4.3-0.20261029045006-4deac5da0714`. Both derive from ancestor `v5.4.2` and actual commit timestamps/hashes. These are **unpublished local build artifacts**, not published upstream versions. An isolated `file://` GOPROXY serves them; consumer go.mod files contain no replace directive. The same consumer source, go.mod and go.sum bytes were used for Windows and Linux.

Canonical archives use `git -c core.autocrlf=false -c core.eol=lf -c tar.umask=0022 archive --format=tar COMMIT`. All 159 candidate / 158 baseline paths, Git blob bytes and permission modes match the final TAR and ZIP. Official `golang.org/x/mod/zip.CheckZip` accepts both; actual downloaded ZIP hashes match, and every extracted module-cache file matches its Git bytes. Only compact manifests and proxy artifacts are included, not build/module caches. Two earlier invalid archives remain internal and are excluded: default TAR permissions were 0664, and Windows native text conversion introduced CRLF before explicit LF/umask options were added.

The ZIP includes the upstream fixed TLS test fixture `_fixture/certs/key.pem`. Its bytes and Git blob are identical in base and candidate; it is an existing public test key, not a newly introduced credential. The public package credential scan identifies this expected fixture explicitly.

## Reading and reproducing

`GATES.md` maps gates to evidence. `hosted-ci/EVIDENCE_EXCERPTS.json` provides line references into captured hosted logs. `RAW_TO_PUBLIC_MANIFEST.json` records source/target hashes and every redaction/transformation. Private workstation paths are replaced by `<workspace>` and `<user>`; those are explanatory placeholders, not literal commands. Local logs here are redacted review copies, not unchanged internal originals. Hosted runner logs are byte-preserved where the manifest says so.

To consume the included artifact, set GOPROXY to the absolute file URL of `modules/proxy` followed by `,https://proxy.golang.org`, set GONOSUMDB to `github.com/labstack/echo/v5`, and run `go test -v -count=1 ./...` in a consumer directory (use `-race` where supported). The baseline is intentionally expected to fail three scenarios. Command receipts record the actually executed runs; no new execution is implied by these reproduction instructions.
