FROM python:slim
COPY * /golem/
WORKDIR /golem/
ENTRYPOINT ["sh"]
