FROM  ubuntu:latest

# apt-get ソースを更新し、gnupgをインストール
RUN apt-get update && apt-get install -my wget gnupg
    
# MongoDB公開GPG鍵を取り込み、MongoDBリストファイルを作成
RUN wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | apt-key add -
RUN echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-4.2.list

# apt-getソースを更新し、MongoDBをインストール
RUN apt-get update && apt-get install -y mongodb-org

RUN mkdir -p /data/db

EXPOSE 27017

# usr/bin/mongod を Docker 化アプリケーションのエントリ・ポイントに設定
ENTRYPOINT ["/usr/bin/mongod"]