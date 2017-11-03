#!/bin/bash
vegeta attack -targets=poll-target.txt -body=poll-body.json -rate=1000 -duration=3s | tee results.bin | vegeta report
