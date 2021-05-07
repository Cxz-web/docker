FROM alpine
ADD shop /shop
ENTRYPOINT [ "/shop" ]
