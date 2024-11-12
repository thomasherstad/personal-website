FROM debian:stable-slim

COPY personal-website /bin/personal-website

WORKDIR /templates
COPY . /templates

WORKDIR /static
COPY . /static

WORKDIR /assets
COPY . /assets

ENV PORT=8080

CMD ["/bin/personal-website"]