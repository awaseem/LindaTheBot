FROM iron/base
WORKDIR /app
# copy binary into image
COPY LindaTheBot /app/
ENTRYPOINT ["./LindaTheBot"]