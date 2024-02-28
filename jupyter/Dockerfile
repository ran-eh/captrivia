FROM jupyter/minimal-notebook
USER root
RUN apt-get update && apt-get install -y postgresql
USER 1000
WORKDIR /work
COPY ./requirements.txt /tmp/requirements.txt
RUN pip install -r /tmp/requirements.txt
COPY ./*ipynb .
