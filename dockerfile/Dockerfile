FROM ocdr/dkubeadm:2.0.7.1-tekton
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg  add - && apt-get update -y && apt-get install google-cloud-sdk -y
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN dpkg -i google-chrome-stable_current_amd64.deb; apt-get -fy install
RUN apt update && apt install git vim python-pip unzip wget curl jq -y &&  pip install selenium paramiko boto3 botocore oauth2client google-api-core  google-api-python-client google-auth google-auth-httplib2 google-cloud-core google-cloud-dataproc google-cloud-datastore google-cloud-storage google-resumable-media googleapis-common-protos grpcio && wget -N https://chromedriver.storage.googleapis.com/83.0.4103.39/chromedriver_linux64.zip -P /tmp && unzip /tmp/chromedriver_linux64.zip -d /usr/local/bin && chmod +x /usr/local/bin/chromedriver
RUN mkdir -p /root/.dkube
COPY ./dkube.ini /root/.dkube/dkube.ini
COPY config.json /opt/dkube
COPY run.sh /usr/local/bin
RUN chmod +x /usr/local/bin/run.sh
