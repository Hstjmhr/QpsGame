# 复制node.ini 文件到
\cp ./gvb_server.ini /etc/supervisord.d

cd ../

./main -esload init/data/article_index_20231112.json
./main -esload init/data/full_text_index_20231112.json
supervisorctl reload