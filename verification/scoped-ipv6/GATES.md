# Verification gates

```json
{
  "base": "4deac5da0714abbc2037a90067ea82c033a3f266",
  "finalCommit": "42d7b17d99c0678ac22fe44076a4c94769964fc6",
  "tree": "667325f667030f93e51f9aa3848525929e7937ab",
  "gates": [
    {
      "gate": "Public HTTP regression",
      "status": "PASS on final candidate; expected master FAIL",
      "evidence": "module-zip/{windows,linux}-{candidate,baseline}-test.log"
    },
    {
      "gate": "IP/Scheme and strict parsing semantics",
      "status": "PASS both control and optimized trees",
      "evidence": "optimization/*related-tests.log and semantics-*"
    },
    {
      "gate": "Quality tools",
      "status": "PASS local and final hosted Checks",
      "evidence": "optimization/optimized-*.log; hosted-ci/DELIVERY.json"
    },
    {
      "gate": "Native Windows full suite",
      "status": "ENVIRONMENT FAIL on both compared trees; not concealed",
      "evidence": "native-environment/*full.log"
    },
    {
      "gate": "Final original hosted 9-matrix race/coverage",
      "status": "PASS all nine; three packages each",
      "evidence": "hosted-ci/DELIVERY.json and ten executed job logs"
    },
    {
      "gate": "No-replace final module ZIP consumption",
      "status": "PASS Windows and Linux race; exact ZIP/cache bytes verified",
      "evidence": "module-zip/DELIVERY.json and archive-zip-proof.json"
    },
    {
      "gate": "Hosted benchmark",
      "status": "SKIPPED by original push condition, not a pass",
      "evidence": "hosted-ci/DELIVERY.json"
    },
    {
      "gate": "Local benchmarks",
      "status": "COMPLETED; scoped allocation improvement and explicit noise limitations",
      "evidence": "optimization/benchmark-* and advisory-full-benchmark/*"
    }
  ],
  "scopeLimits": [
    "RemoteAddr is a documented fixture, not an OS link-local route",
    "Local pseudo-versions are not published upstream releases",
    "Own-fork CI is separate from upstream approval/merge/release"
  ]
}
```
