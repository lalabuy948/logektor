echo "POST http://localhost:8080/track" | vegeta attack -body clientTest/testpayload.json -insecure -duration 10s -rate 1000 | vegeta report