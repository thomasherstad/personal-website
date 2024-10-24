FROM debian:stable-slim

COPY personal-website /bin/personal-website
COPY /static /bin/static

ENV PORT=8080

CMD ["/bin/personal-website"]