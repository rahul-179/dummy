#!/bin/bash
echo "cloning automation script ...."
TOTAL_TEST=0
TEST_PASS=0
WORKDIR="/opt/dkube"
#echo"deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg  add - && apt-get update -y && apt-get install google-cloud-sdk -y

#apt-getupdate && apt-get install git vim python-pip unzip wget -y &&  pip install selenium paramiko boto3 botocore oauth2client google-api-core  google-api-python-client google-auth google-auth-httplib2 google-cloud-core google-cloud-dataproc google-cloud-datastore google-cloud-storage google-resumable-media googleapis-common-protos grpcio && wget -N https://chromedriver.storage.googleapis.com/83.0.4103.39/chromedriver_linux64.zip -P /tmp && unzip /tmp/chromedriver_linux64.zip -d /usr/bin && chmod +x /usr/bin/chromedriver
#ls-l /usr/bin 

#echo$PATH 
#export"PATH=\$PATH:/usr/local/bin" >> $HOME/.bashrc
#source$HOME/.bashrc
#echo$PATH 
cd $WORKDIR
git clone https://rahul-179:malhotrarahul1234@github.com/oneconvergence/gpuaas.git -b automation_2.0
if [ -d "gpuaas" ]
then
   sed -i  "s/\"ip\" :.*/\"ip\" :$(cat $WORKDIR/config.json | jq '.ip'),/" gpuaas/testing/atf/ui/config/setup.json
   sed -i  "s/\"port\" :.*/\"port\" :$(cat $WORKDIR/config.json | jq '.port'),/" gpuaas/testing/atf/ui/config/setup.json
   sed -i  "s/\"total_nodes\" :.*/\"total_nodes\" :$(cat $WORKDIR/config.json | jq '.total_nodes'),/" gpuaas/testing/atf/ui/config/config.json
   sed -i  "s/\"pool_type\" :.*/\"pool_type\" :$(cat $WORKDIR/config.json | jq '.pool_type'),/" gpuaas/testing/atf/ui/config/config.json
   sed -i  "s/\"total_gpus\" :.*/\"total_gpus\" :$(cat $WORKDIR/config.json | jq '.total_gpus'),/" gpuaas/testing/atf/ui/config/config.json

 	cd gpuaas/testing/atf/ui/tests/suites
	 echo "Running tests ...."
	 python -u  ds_suite.py crud  --mode $(cat $WORKDIR/config.json | jq '.mode' | tr -d '"') --headless True >> result.txt 
	 #python ds_suite.py crud  --mode cpu  --headless True >> result.txt
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
	echo"Could not find test code, exiting ...."
	exit1
fi
