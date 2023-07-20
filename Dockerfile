
FROM alpine

LABEL maintainer="imoowi"
ENV VERSION=1.0
WORKDIR /
VOLUME ["/apps/config"]
COPY bin/gitlabevent2wechatbot /
COPY configs /configs
# COPY logs /logs
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8


#EXPOSE 声明容器的服务端口（仅仅是声明）
EXPOSE 8000


#CMD 运行容器时执行的shell环境
CMD ["/gitlabevent2wechatbot","server","-c", "/configs/settings-local.yml"]                                                                                     