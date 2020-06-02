#!/bin/bash
echo "cloning automation script ...."
TOTAL_TEST=0
TEST_PASS=0
git clone https://rahul-179:malhotrarahul1234@github.com/oneconvergence/gpuaas.git -b $(cat config.json | jq '.git_branch' | tr -d '"')
if [ -d "gpuaas" ]
then
    sed -i  "s/\"ip\" :.*/\"ip\" :$(cat config.json | jq '.ip'),/" gpuaas/testing/atf/ui/config/setup.json
    sed -i  "s/\"port\" :.*/\"port\" :$(cat config.json | jq '.port'),/" gpuaas/testing/atf/ui/config/setup.json
    sed -i  "s/\"total_nodes\" :.*/\"total_nodes\" :$(cat config.json | jq '.total_nodes'),/" gpuaas/testing/atf/ui/config/config.json
    sed -i  "s/\"pool_type\" :.*/\"pool_type\" :$(cat config.json | jq '.pool_type'),/" gpuaas/testing/atf/ui/config/config.json
    sed -i  "s/\"total_gpus\" :.*/\"total_gpus\" :$(cat config.json | jq '.total_gpus'),/" gpuaas/testing/atf/ui/config/config.json

  	cd gpuaas/testing/atf/ui/tests/suites
	  echo "Running tests ...."
	  python ds_suite.py crud  --mode $(cat /atf-test/config.json | jq '.mode' | tr -d '"') --headless True >> result.txt
    cat result.txt
    TOTAL_TEST=$(cat result.txt | awk '{print$1}'| grep -c "F\|E\|ok")
    TEST_PASS=$(cat result.txt | awk '{print$1}'| grep -c "ok")
    echo "TOTAL_TEST:" $TOTAL_TEST
    echo "TEST_PASS:" $TEST_PASS
    if [ $TEST_PASS == $TOTAL_TEST ]
    then
       echo "All tests passed"
       exit 0
    else
       echo "$TEST_PASS tests passed out of $TOTAL_TEST"
       exit 1
    fi

else
	echo "Could not find test code, exiting ...."
	exit 1
fi
