FROM debian:stable-slim

COPY personal-website /bin/personal-website
COPY /templates /templates
COPY /static /static
COPY /assets /assets

ENV PORT=8080

CMD ["/bin/personal-website"]